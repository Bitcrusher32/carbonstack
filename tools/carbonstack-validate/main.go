package main

import (
"bufio"
"errors"
"flag"
"fmt"
"os"
"os/exec"
"path/filepath"
"runtime"
"strings"
"time"
)

type Runner struct {
Profile      string
StartDir     string
UmbrellaRoot string
CarbonStack  string
Comms        string
Cypher       string
Failed       bool
}

type Step struct {
Name    string
Dir     string
Command string
Args    []string
Env     []string
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

_ = r.RunStepAllowFailure(Step{Name: "go version", Dir: r.CarbonStack, Command: "go", Args: []string{"version"}})
_ = r.RunStepAllowFailure(Step{Name: "rustc version", Dir: r.CarbonStack, Command: "rustc", Args: []string{"--version"}})
_ = r.RunStepAllowFailure(Step{Name: "cargo version", Dir: r.CarbonStack, Command: "cargo", Args: []string{"--version"}})
_ = r.RunStepAllowFailure(Step{Name: "sqlite3 version", Dir: r.CarbonStack, Command: "sqlite3", Args: []string{"--version"}})

fmt.Println()
fmt.Println("== Rust floor note ==")
fmt.Println("OpenMLS 0.8.1 failed under Debian apt rustc 1.85.0 during v0.3.9.")
fmt.Println("rustup stable rustc/cargo 1.96.0 passed under WSL Debian during v0.3.9.")
fmt.Println("This runner reports toolchain versions but does not install or mutate toolchains.")

return nil
}

func (r *Runner) Core() error {
r.PrintHeader("core")

if err := r.Doctor(); err != nil {
return err
}

fmt.Println()
fmt.Println("== Pre-test artifact scan ==")
r.ArtifactScan()

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
fmt.Println("== Post-test artifact scan ==")
r.ArtifactScan()

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

func (r *Runner) RunStepAllowFailure(step Step) error {
r.PrintStep(step)

cmd := exec.Command(step.Command, step.Args...)
cmd.Dir = step.Dir
cmd.Env = append(os.Environ(), step.Env...)

out, err := cmd.CombinedOutput()
fmt.Print(string(out))

if err != nil {
fmt.Printf("WARN: %s failed: %v\n", step.Name, err)
return err
}

return nil
}

func stream(prefix string, pipe interface {
Read([]byte) (int, error)
}, done chan<- struct{}) {
defer func() { done <- struct{}{} }()

scanner := bufio.NewScanner(pipe)
for scanner.Scan() {
fmt.Printf("[%s] %s\n", prefix, scanner.Text())
}
}

func (r *Runner) ArtifactScan() {
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

var hits []string

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
hits = append(hits, path)
if d.IsDir() {
return filepath.SkipDir
}
return nil
}
}

if !d.IsDir() {
for _, suffix := range suffixes {
if strings.HasSuffix(name, suffix) {
hits = append(hits, path)
return nil
}
}
}

return nil
})
}

if len(hits) == 0 {
fmt.Println("artifact scan: no generated/private/build artifact hits")
return
}

fmt.Println("artifact scan hits:")
for _, hit := range hits {
fmt.Printf("  %s\n", hit)
}
fmt.Println("artifact scan is non-destructive; interpret pre-test hits differently from post-test hits")
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

var _ = errors.New
