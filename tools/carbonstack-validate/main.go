package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const (
	minRustMajor = 1
	minRustMinor = 96
)

type Runner struct {
	Profile      string
	StartDir     string
	UmbrellaRoot string
	CarbonStack  string
	Comms        string
	Cypher       string
}

type Step struct {
	Name    string
	Dir     string
	Command string
	Args    []string
	Env     []string
}

type ArtifactHit struct {
	Path string
	Kind string
}

func main() {
	profile := flag.String("profile", "doctor", "validation profile: doctor, core, full")
	rootOverride := flag.String("root", "", "optional umbrella root containing carbonstack, carbonstack-comms, carbonstack-cypher")
	flag.Parse()

	r, err := NewRunner(*profile, *rootOverride)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
		os.Exit(2)
	}

	var runErr error

	switch r.Profile {
	case "doctor":
		runErr = r.Doctor()
	case "core":
		runErr = r.Core()
	case "full":
		fmt.Println("profile full currently aliases core")
		runErr = r.Core()
	default:
		runErr = fmt.Errorf("unknown profile %q; expected doctor, core, or full", r.Profile)
	}

	if runErr != nil {
		fmt.Fprintf(os.Stderr, "\nVALIDATION FAILED: %v\n", runErr)
		os.Exit(1)
	}

	fmt.Println("\nVALIDATION PASSED")
}

func NewRunner(profile string, rootOverride string) (*Runner, error) {
	startDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	var umbrellaRoot string

	if rootOverride != "" {
		abs, err := filepath.Abs(rootOverride)
		if err != nil {
			return nil, err
		}
		umbrellaRoot = abs
	} else {
		absStart, err := filepath.Abs(startDir)
		if err != nil {
			return nil, err
		}

		found, err := inferUmbrellaRoot(absStart)
		if err != nil {
			return nil, err
		}
		umbrellaRoot = found
	}

	return &Runner{
		Profile:      profile,
		StartDir:     startDir,
		UmbrellaRoot: umbrellaRoot,
		CarbonStack:  filepath.Join(umbrellaRoot, "carbonstack"),
		Comms:        filepath.Join(umbrellaRoot, "carbonstack-comms"),
		Cypher:       filepath.Join(umbrellaRoot, "carbonstack-cypher"),
	}, nil
}

func inferUmbrellaRoot(start string) (string, error) {
	current := start

	for {
		if hasSiblingRepoLayout(current) {
			return current, nil
		}

		if filepath.Base(current) == "carbonstack" {
			parent := filepath.Dir(current)
			if hasSiblingRepoLayout(parent) {
				return parent, nil
			}
		}

		parent := filepath.Dir(current)
		if parent == current {
			break
		}
		current = parent
	}

	return "", fmt.Errorf("could not infer umbrella root from %s; pass --root explicitly", start)
}

func hasSiblingRepoLayout(root string) bool {
	required := []string{
		filepath.Join(root, "carbonstack"),
		filepath.Join(root, "carbonstack-comms"),
		filepath.Join(root, "carbonstack-cypher"),
	}

	for _, path := range required {
		info, err := os.Stat(path)
		if err != nil || !info.IsDir() {
			return false
		}
	}

	return true
}

func (r *Runner) Doctor() error {
	r.PrintHeader("doctor")

	fmt.Printf("os:             %s\n", runtime.GOOS)
	fmt.Printf("arch:           %s\n", runtime.GOARCH)
	fmt.Printf("start_dir:      %s\n", r.StartDir)
	fmt.Printf("umbrella_root:  %s\n", r.UmbrellaRoot)
	fmt.Printf("carbonstack:    %s\n", r.CarbonStack)
	fmt.Printf("comms:          %s\n", r.Comms)
	fmt.Printf("cypher:         %s\n", r.Cypher)

	if err := r.CheckRequiredPaths(); err != nil {
		return err
	}

	fmt.Println()
	fmt.Println("== Toolchains ==")

	goVersion := r.ReportTool("go", "version")
	rustVersion := r.ReportTool("rustc", "--version")
	_ = r.ReportTool("cargo", "--version")
	_ = r.ReportTool("sqlite3", "--version")

	fmt.Println()
	fmt.Println("== Rust floor note ==")
	fmt.Println("OpenMLS 0.8.1 failed under Debian apt rustc 1.85.0 during v0.3.9.")
	fmt.Println("rustup stable rustc/cargo 1.96.0 passed under WSL Debian during v0.3.9.")
	fmt.Println("This runner reports toolchain versions but does not install or mutate toolchains.")

	if goVersion == "" {
		fmt.Println("WARN: go was not found or did not report a version")
	}

	if rustVersion == "" {
		fmt.Println("WARN: rustc was not found or did not report a version")
	} else if !rustAtLeast(rustVersion, minRustMajor, minRustMinor) {
		fmt.Printf("WARN: rustc appears older than known-good floor %d.%d for current OpenMLS sidecar tests\n", minRustMajor, minRustMinor)
	}

	return nil
}

func (r *Runner) Core() error {
	r.PrintHeader("core")

	if err := r.Doctor(); err != nil {
		return err
	}

	fmt.Println()
	r.ArtifactScan("pre-test")

	steps := []Step{
		{
			Name:    "OpenMLS real-Cypher lifecycle",
			Dir:     r.Comms,
			Command: "go",
			Args: []string{
				"test",
				"./internal/protocol",
				"-run",
				"TestOpenMLSSidecarFullLifecycleRelayThroughRealCypherServer",
				"-count=1",
				"-timeout",
				"300s",
				"-v",
			},
			Env: []string{"RUST_BACKTRACE=1"},
		},
		{
			Name:    "carbonstack-comms package tests",
			Dir:     r.Comms,
			Command: "go",
			Args: []string{
				"test",
				"./...",
				"-count=1",
				"-timeout",
				"600s",
			},
			Env: []string{"RUST_BACKTRACE=1"},
		},
		{
			Name:    "carbonstack-cypher package tests",
			Dir:     r.Cypher,
			Command: "go",
			Args: []string{
				"test",
				"./...",
				"-count=1",
			},
		},
	}

	for _, step := range steps {
		if err := r.RunStep(step); err != nil {
			return err
		}
	}

	fmt.Println()
	r.ArtifactScan("post-test")

	return nil
}

func (r *Runner) CheckRequiredPaths() error {
	fmt.Println()
	fmt.Println("== Required paths ==")

	required := []string{
		r.CarbonStack,
		r.Comms,
		r.Cypher,
		filepath.Join(r.Comms, "go.mod"),
		filepath.Join(r.Cypher, "go.mod"),
		filepath.Join(r.Comms, "internal", "protocol", "mls", "openmls-sidecar", "Cargo.toml"),
	}

	var missing []string

	for _, path := range required {
		if _, err := os.Stat(path); err != nil {
			fmt.Printf("MISSING: %s\n", path)
			missing = append(missing, path)
		} else {
			fmt.Printf("OK:      %s\n", path)
		}
	}

	if len(missing) > 0 {
		return fmt.Errorf("missing required paths: %d", len(missing))
	}

	return nil
}

func (r *Runner) ReportTool(command string, args ...string) string {
	path, err := exec.LookPath(command)
	if err != nil {
		fmt.Printf("%s path: WARN not found in PATH\n", command)
		fmt.Printf("%s version: WARN unavailable\n", command)
		return ""
	}

	fmt.Printf("%s path: %s\n", command, path)

	cmd := exec.Command(command, args...)
	cmd.Dir = r.CarbonStack
	out, err := cmd.CombinedOutput()
	text := strings.TrimSpace(string(out))

	if text != "" {
		fmt.Printf("%s version: %s\n", command, text)
	}

	if err != nil {
		fmt.Printf("%s version: WARN command failed: %v\n", command, err)
		return ""
	}

	return text
}

func (r *Runner) RunStep(step Step) error {
	r.PrintStep(step)

	cmd := exec.Command(step.Command, step.Args...)
	cmd.Dir = step.Dir
	cmd.Env = append(os.Environ(), step.Env...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	start := time.Now()

	if err := cmd.Start(); err != nil {
		return err
	}

	done := make(chan struct{})

	go stream("stdout", stdout, done)
	go stream("stderr", stderr, done)

	err = cmd.Wait()

	<-done
	<-done

	elapsed := time.Since(start)

	if err != nil {
		fmt.Printf("FAIL: %s elapsed=%s\n", step.Name, elapsed.Round(time.Millisecond))
		return fmt.Errorf("%s failed: %w", step.Name, err)
	}

	fmt.Printf("PASS: %s elapsed=%s\n", step.Name, elapsed.Round(time.Millisecond))
	return nil
}

func stream(prefix string, pipe io.Reader, done chan<- struct{}) {
	defer func() { done <- struct{}{} }()

	scanner := bufio.NewScanner(pipe)
	for scanner.Scan() {
		fmt.Printf("[%s] %s\n", prefix, scanner.Text())
	}
}

func (r *Runner) ArtifactScan(phase string) {
	fmt.Printf("== %s artifact scan ==\n", phase)

	patterns := []string{
		"target",
		".carbonstack-openmls-sidecar-state",
		"provider-storage.json",
		"signer.json",
		".go-cache",
		".go-tmp",
	}

	suffixes := []string{
		".db",
		".db-shm",
		".db-wal",
		".exe",
		".test.exe",
	}

	roots := []string{
		r.CarbonStack,
		r.Comms,
		r.Cypher,
	}

	var hits []ArtifactHit

	for _, root := range roots {
		_ = filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
			if err != nil {
				return nil
			}

			name := d.Name()

			if name == ".git" {
				if d.IsDir() {
					return filepath.SkipDir
				}
				return nil
			}

			for _, pattern := range patterns {
				if name == pattern {
					hits = append(hits, ArtifactHit{Path: path, Kind: pattern})
					if d.IsDir() {
						return filepath.SkipDir
					}
					return nil
				}
			}

			if !d.IsDir() {
				for _, suffix := range suffixes {
					if strings.HasSuffix(name, suffix) {
						hits = append(hits, ArtifactHit{Path: path, Kind: suffix})
						return nil
					}
				}
			}

			return nil
		})
	}

	if len(hits) == 0 {
		fmt.Printf("%s artifact scan: no generated/private/build artifact hits\n", phase)
		return
	}

	fmt.Printf("%s artifact scan hits:\n", phase)
	for _, hit := range hits {
		class := classifyArtifactPath(hit.Path)
		fmt.Printf("  [%s] %s\n", class, hit.Path)
	}

	fmt.Println("artifact scan is non-destructive")
	fmt.Println("pre-test hits are potential source/copy hygiene issues")
	fmt.Println("post-test hits are expected only when they stay in known generated roots")
}

func classifyArtifactPath(path string) string {
	normalized := filepath.ToSlash(path)

	if strings.Contains(normalized, "internal/protocol/mls/openmls-sidecar/target") ||
		strings.Contains(normalized, "internal/protocol/mls/openmls-sidecar/.carbonstack-openmls-sidecar-state") {
		return "known-openmls-sidecar-generated-root"
	}

	if strings.Contains(normalized, "internal/protocol/mls/research/") {
		return "research-generated-root"
	}

	if strings.Contains(normalized, ".go-cache") || strings.Contains(normalized, ".go-tmp") {
		return "local-go-cache-root"
	}

	return "review"
}

func rustAtLeast(versionText string, major int, minor int) bool {
	fields := strings.Fields(versionText)
	if len(fields) < 2 {
		return false
	}

	version := fields[1]
	parts := strings.Split(version, ".")
	if len(parts) < 2 {
		return false
	}

	gotMajor, err := strconv.Atoi(parts[0])
	if err != nil {
		return false
	}

	gotMinor, err := strconv.Atoi(parts[1])
	if err != nil {
		return false
	}

	if gotMajor != major {
		return gotMajor > major
	}

	return gotMinor >= minor
}

func (r *Runner) PrintHeader(name string) {
	fmt.Println("========================================")
	fmt.Printf("CarbonStack validation profile: %s\n", name)
	fmt.Println("========================================")
}

func (r *Runner) PrintStep(step Step) {
	fmt.Println()
	fmt.Println("----------------------------------------")
	fmt.Printf("STEP: %s\n", step.Name)
	fmt.Printf("DIR:  %s\n", step.Dir)
	fmt.Printf("CMD:  %s %s\n", step.Command, strings.Join(step.Args, " "))
	if len(step.Env) > 0 {
		fmt.Printf("ENV:  %s\n", strings.Join(step.Env, " "))
	}
	fmt.Println("----------------------------------------")
}

var _ = bytes.Compare
