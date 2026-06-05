package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func (r *Runner) DevRuntimeOpenMLS() error {
	r.PrintHeader("dev-runtime-openmls")

	fmt.Println("status: dev/pre-alpha manual validation profile")
	fmt.Println("scope: OpenMLS application-message runtime CLI path through Comms and Cypher")
	fmt.Println("proof: openmls-send-dev -> Cypher -> openmls-inbox-dev --ack")
	fmt.Println("boundary: not local-backbone, not mature messaging UX, not deployment, and not a production/security claim")
	fmt.Println("release-package status: live umbrella checkout only for now; not included in full")

	if err := r.CheckRequiredPaths(); err != nil {
		return err
	}

	if err := r.CheckLiveGitUmbrella("dev-runtime-openmls"); err != nil {
		return err
	}

	script := filepath.Join(r.Comms, "scripts", "dev-openmls-runtime-smoke.sh")
	if info, err := os.Stat(script); err != nil || info.IsDir() {
		return fmt.Errorf("missing dev OpenMLS runtime smoke script: %s", script)
	}

	fmt.Println()
	fmt.Println("== Toolchains ==")
	_ = r.ReportTool("go", "version")
	_ = r.ReportTool("rustc", "--version")
	_ = r.ReportTool("cargo", "--version")
	_ = r.ReportTool("curl", "--version")
	_ = r.ReportTool("python3", "--version")
	_ = r.ReportTool("bash", "--version")

	fmt.Println()
	r.ArtifactScan("pre-dev-runtime-openmls")

	step := Step{
		Name:    "carbonstack-comms dev OpenMLS runtime smoke proof",
		Dir:     r.Comms,
		Command: "bash",
		Args:    []string{script},
		Env:     []string{"RUST_BACKTRACE=1"},
	}

	if err := r.RunStep(step); err != nil {
		return err
	}

	fmt.Println()
	r.ArtifactScan("post-dev-runtime-openmls")

	fmt.Println()
	fmt.Println("dev-runtime-openmls profile result:")
	fmt.Println("  PASS: dev runtime OpenMLS CLI smoke proof completed")
	fmt.Println("  proof: openmls-send-dev -> Cypher -> openmls-inbox-dev --ack")
	fmt.Println("  boundary: not local-backbone, not production messaging UX, not deployment, not security certification")
	fmt.Println("  note: run with --clean-generated to remove known OpenMLS sidecar generated roots after success")

	return nil
}

func (r *Runner) CheckLiveGitUmbrella(profileName string) error {
	fmt.Println()
	fmt.Println("== Live umbrella checkout guard ==")

	required := []string{
		filepath.Join(r.CarbonStack, ".git"),
		filepath.Join(r.Comms, ".git"),
		filepath.Join(r.Cypher, ".git"),
	}

	var missing []string
	for _, path := range required {
		if _, err := os.Stat(path); err != nil {
			fmt.Printf("MISSING LIVE CHECKOUT MARKER: %s\n", path)
			missing = append(missing, path)
		} else {
			fmt.Printf("OK LIVE CHECKOUT MARKER: %s\n", path)
		}
	}

	if len(missing) > 0 {
		return fmt.Errorf("%s is currently live-umbrella only; use a git checkout with sibling carbonstack, carbonstack-comms, and carbonstack-cypher repos", profileName)
	}

	fmt.Println("live umbrella guard passed")
	return nil
}
