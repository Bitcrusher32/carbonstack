package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func (r *Runner) DevRuntimeOpenMLSWrappers() error {
	r.PrintHeader("dev-runtime-openmls-wrappers")

	fmt.Println("status: dev/pre-alpha manual validation profile")
	fmt.Println("scope: OpenMLS wrapper-bootstrap runtime CLI path through Comms and Cypher")
	fmt.Println("proof: openmls-*-dev bootstrap wrappers -> message-send-dev -> Cypher -> message-inbox-dev --ack")
	fmt.Println("boundary: not local-backbone, not mature messaging UX, not deployment, and not a production/security claim")
	fmt.Println("release-package status: live umbrella checkout only for now; not included in full")
	fmt.Println("relationship: separate maturity profile; does not replace dev-runtime-openmls yet")

	if err := r.CheckRequiredPaths(); err != nil {
		return err
	}

	if err := r.CheckLiveGitUmbrella("dev-runtime-openmls-wrappers"); err != nil {
		return err
	}

	script := filepath.Join(r.Comms, "scripts", "dev-openmls-runtime-smoke-wrappers.sh")
	if info, err := os.Stat(script); err != nil || info.IsDir() {
		return fmt.Errorf("missing wrapper dev OpenMLS runtime smoke script: %s", script)
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
	r.ArtifactScan("pre-dev-runtime-openmls-wrappers")

	step := Step{
		Name:    "carbonstack-comms wrapper dev OpenMLS runtime smoke proof",
		Dir:     r.Comms,
		Command: "bash",
		Args:    []string{script},
		Env:     []string{"RUST_BACKTRACE=1"},
	}

	if err := r.RunStep(step); err != nil {
		return err
	}

	fmt.Println()
	r.ArtifactScan("post-dev-runtime-openmls-wrappers")

	fmt.Println()
	fmt.Println("dev-runtime-openmls-wrappers profile result:")
	fmt.Println("  PASS: wrapper-based dev runtime OpenMLS CLI smoke proof completed")
	fmt.Println("  proof: openmls-*-dev bootstrap wrappers -> message-send-dev -> Cypher -> message-inbox-dev --ack")
	fmt.Println("  boundary: not local-backbone, not production messaging UX, not deployment, not security certification")
	fmt.Println("  relationship: separate message-wrapper smoke maturity profile; does not replace dev-runtime-openmls yet")
	fmt.Println("  note: run with --clean-generated to remove known OpenMLS sidecar generated roots after success")

	return nil
}
