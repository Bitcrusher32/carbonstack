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
	Profile        string
	CleanGenerated bool
	CompactSummary bool
	StartDir       string
	UmbrellaRoot   string
	CarbonStack    string
	Comms          string
	Cypher         string
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
	profile := flag.String("profile", "doctor", "validation profile: doctor, core, local-cypher, dev-runtime-openmls, dev-runtime-openmls-wrappers, integrated-runtime-dev, same-state-integrated-dev, same-state-message-failure-dev, same-state-message-unsupported-dev, same-state-message-malformed-payload-dev, same-state-message-replay-classification-dev, same-state-message-recipient-failure-dev, same-state-welcome-join-failure-dev, registry-lookup, relay-openmls-join-dev, relay-space-invite-claim-dev, relay-space-member-state-dev, relay-space-member-restart-dev, relay-space-delivery-authority-dev, keypackage-inspect-dev, keypackage-rotation-dev, keypackage-publication-dev, keypackage-consume-dev, welcome-lifecycle-dev, cypher-mls-mismatch-dev, workflow-relay-onboarding-dev, gate-b-relay-lifecycle-closure-dev, state-substrate-inventory-dev, state-schema-compat-dev, state-path-policy-dev, state-write-policy-dev, gate-c-state-substrate-closure-dev, gate-d-runtime-aggregate-dev, gate-e-native-deployment-dev, gate-e-native-deployment-closure-dev, gate-f-release-package-surface-dev, gate-f-operator-docs-runbook-dev, gate-f-compat-rollback-observability-dev, gate-f-code-health-source-hygiene-dev, gate-f-basic-local-trust-posture-dev, gate-f-package-runtime-candidate-dev, gate-f-release-candidate-closure-dev, full, full-validate-release, release-snapshot, write-checksums, verify-checksums")
	rootOverride := flag.String("root", "", "optional umbrella root containing carbonstack, carbonstack-comms, carbonstack-cypher")
	registryID := flag.String("registry-id", "", "registry id to inspect when --profile registry-lookup is used")
	registryCommand := flag.String("command", "", "literal registry command to inspect when --profile registry-lookup is used")
	registryList := flag.Bool("list", false, "list registry entries when --profile registry-lookup is used")
	registryAudience := flag.String("audience", "", "filter registry lookup list by audience")
	registryMaturity := flag.String("maturity", "", "filter registry lookup list by maturity")
	registryLifecycleStatus := flag.String("lifecycle-status", "", "filter registry lookup list by lifecycle_status")
	registryKind := flag.String("kind", "", "filter registry lookup list by kind")
	registryFrontReadmeOnly := flag.Bool("front-readme-only", false, "filter registry lookup list to front README candidates")
	registryMissingNonclaims := flag.Bool("missing-nonclaims", false, "filter registry lookup list to entries without nonclaims")
	cleanGenerated := flag.Bool("clean-generated", false, "after a successful profile run, remove known generated/build artifacts such as OpenMLS sidecar target/state roots")
	compactSummary := flag.Bool("compact-summary", false, "print a compact profile evidence summary for supported validation profiles")
	flag.Parse()

	r, err := NewRunner(*profile, *rootOverride)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
		os.Exit(2)
	}
	r.CleanGenerated = *cleanGenerated
	r.CompactSummary = *compactSummary

	var runErr error

	switch r.Profile {
	case "doctor":
		runErr = r.Doctor()
	case "core":
		runErr = r.Core()
	case "local-cypher":
		runErr = r.LocalCypher()
	case "dev-runtime-openmls":
		runErr = r.DevRuntimeOpenMLS()
	case "dev-runtime-openmls-wrappers":
		runErr = r.DevRuntimeOpenMLSWrappers()
	case "integrated-runtime-dev":
		runErr = r.RunIntegratedRuntimeDev(*cleanGenerated)
	case "same-state-integrated-dev":
		runErr = r.SameStateIntegratedDev()
	case "same-state-message-failure-dev":
		runErr = r.SameStateMessageFailureDev()
	case "same-state-message-unsupported-dev":
		runErr = r.SameStateMessageUnsupportedDev()
	case "same-state-message-malformed-payload-dev":
		runErr = r.SameStateMessageMalformedPayloadDev()
	case "same-state-message-replay-classification-dev":
		runErr = r.SameStateMessageReplayClassificationDev()
	case "same-state-message-recipient-failure-dev":
		runErr = r.SameStateMessageRecipientFailureDev()
	case "same-state-welcome-join-failure-dev":
		runErr = r.SameStateWelcomeJoinFailureDev()
	case "registry-lookup":
		runErr = r.RegistryLookupWithOptions(registryLookupOptions{
			RegistryID:       *registryID,
			LiteralCommand:   *registryCommand,
			List:             *registryList,
			Audience:         *registryAudience,
			Maturity:         *registryMaturity,
			LifecycleStatus:  *registryLifecycleStatus,
			Kind:             *registryKind,
			FrontReadmeOnly:  *registryFrontReadmeOnly,
			MissingNonclaims: *registryMissingNonclaims,
		})
	case "relay-openmls-join-dev":
		runErr = r.RelayOpenMLSJoinDev()
	case "relay-space-invite-claim-dev":
		runErr = r.RelaySpaceInviteClaimDev()
	case "relay-space-member-state-dev":
		runErr = r.RelaySpaceMemberStateDev()
	case "relay-space-member-restart-dev":
		runErr = r.RelaySpaceMemberRestartDev()
	case "relay-space-delivery-authority-dev":
		runErr = r.RelaySpaceDeliveryAuthorityDev()
	case "keypackage-inspect-dev":
		runErr = r.KeyPackageInspectDev()
	case "keypackage-rotation-dev":
		runErr = r.KeyPackageRotationDev()
	case "keypackage-publication-dev":
		runErr = r.KeyPackagePublicationDev()
	case "keypackage-consume-dev":
		runErr = r.KeyPackageConsumeDev()
	case "welcome-lifecycle-dev":
		runErr = r.WelcomeLifecycleDev()
	case "cypher-mls-mismatch-dev":
		runErr = r.CypherMLSMismatchDev()
	case "workflow-relay-onboarding-dev":
		runErr = r.WorkflowRelayOnboardingDev()
	case "gate-b-relay-lifecycle-closure-dev":
		runErr = r.GateBRelayLifecycleClosureDev()
	case "state-substrate-inventory-dev":
		runErr = r.StateSubstrateInventoryDev()
	case "state-schema-compat-dev":
		runErr = r.StateSchemaCompatibilityDev()
	case "state-path-policy-dev":
		runErr = r.StatePathPolicyDev()
	case "state-write-policy-dev":
		runErr = r.StateWritePolicyDev()
	case "gate-c-state-substrate-closure-dev":
		runErr = r.GateCStateSubstrateClosureDev()
	case "gate-d-runtime-aggregate-dev":
		runErr = r.GateDRuntimeAggregateDev()
	case "gate-e-native-deployment-dev":
		runErr = r.GateENativeDeploymentDev()
	case "gate-e-native-deployment-closure-dev":
		runErr = r.GateENativeDeploymentClosureDev()
	case "gate-f-release-package-surface-dev":
		runErr = r.GateFReleasePackageSurfaceDev()
	case "gate-f-operator-docs-runbook-dev":
		runErr = r.GateFOperatorDocsRunbookDev()
	case "gate-f-compat-rollback-observability-dev":
		runErr = r.GateFCompatRollbackObservabilityDev()
	case "gate-f-code-health-source-hygiene-dev":
		runErr = r.GateFCodeHealthSourceHygieneDev()
	case "gate-f-basic-local-trust-posture-dev":
		runErr = r.GateFBasicLocalTrustPostureDev()
	case "gate-f-package-runtime-candidate-dev":
		runErr = r.GateFPackageRuntimeCandidateDev()
	case "gate-f-release-candidate-closure-dev":
		runErr = r.GateFReleaseCandidateClosureDev()
	case "full", "full-validate-release":
		fmt.Printf("profile %s runs release-snapshot, then local-cypher\n", r.Profile)
		fmt.Println("release-snapshot already calls core; full/full-validate-release do not call core a second time")
		if err := r.ReleaseSnapshot(); err != nil {
			runErr = err
		} else {
			runErr = r.LocalCypher()
		}
	case "release-snapshot":
		runErr = r.ReleaseSnapshot()
	case "write-checksums":
		runErr = r.WriteReleaseChecksums()
	case "verify-checksums":
		runErr = r.VerifyReleaseChecksums()
	default:
		runErr = fmt.Errorf("unknown profile %q; expected doctor, core, local-cypher, dev-runtime-openmls, dev-runtime-openmls-wrappers, integrated-runtime-dev, same-state-integrated-dev, same-state-message-failure-dev, same-state-message-unsupported-dev, same-state-message-malformed-payload-dev, same-state-message-replay-classification-dev, same-state-message-recipient-failure-dev, same-state-welcome-join-failure-dev, registry-lookup, relay-openmls-join-dev, relay-space-invite-claim-dev, relay-space-member-state-dev, relay-space-member-restart-dev, relay-space-delivery-authority-dev, keypackage-inspect-dev, keypackage-rotation-dev, keypackage-publication-dev, keypackage-consume-dev, welcome-lifecycle-dev, full, full-validate-release, release-snapshot, write-checksums, welcome-lifecycle-dev, or verify-checksums", r.Profile)
	}

	if runErr != nil {
		fmt.Fprintf(os.Stderr, "\nVALIDATION FAILED: %v\n", runErr)
		os.Exit(1)
	}

	if r.CleanGenerated {
		if err := r.CleanGeneratedArtifacts(); err != nil {
			fmt.Fprintf(os.Stderr, "\nVALIDATION FAILED: %v\n", err)
			os.Exit(1)
		}
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

func (r *Runner) CleanGeneratedArtifacts() error {
	fmt.Println()
	fmt.Println("== clean generated artifacts ==")
	fmt.Println("cleanup mode: explicit --clean-generated")
	fmt.Println("cleanup scope: known generated/build artifact roots only")

	paths := []string{
		filepath.Join(r.Comms, "internal", "protocol", "mls", "openmls-sidecar", ".carbonstack-openmls-sidecar-state"),
		filepath.Join(r.Comms, "internal", "protocol", "mls", "openmls-sidecar", "target"),
	}

	for _, path := range paths {
		if _, err := os.Stat(path); err != nil {
			if os.IsNotExist(err) {
				fmt.Println("SKIP absent:", path)
				continue
			}
			return fmt.Errorf("inspect generated artifact path %s: %w", path, err)
		}

		fmt.Println("REMOVE:", path)
		if err := os.RemoveAll(path); err != nil {
			return fmt.Errorf("remove generated artifact path %s: %w", path, err)
		}
	}

	fmt.Println("generated artifact cleanup complete")
	return nil
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
