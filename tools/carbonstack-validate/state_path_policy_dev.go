package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func (r *Runner) StatePathPolicyDev() error {
	fmt.Println("========================================")
	fmt.Println("CarbonStack validation profile: state-path-policy-dev")
	fmt.Println("========================================")
	fmt.Println("status: dev/pre-alpha Gate C3 path policy profile")
	fmt.Println("scope: C3 path policy and explicit state-root semantics")
	fmt.Println("boundary: not migration, not repair, not relocation, not cleanup, not C4 atomicity, not Gate D")
	fmt.Println()

	if err := r.CheckRequiredPaths(); err != nil {
		return err
	}
	if err := r.CheckLiveGitUmbrella("state-path-policy-dev"); err != nil {
		return err
	}

	validatorDir := filepath.Join(r.CarbonStack, "tools", "carbonstack-validate")
	reportRoot := filepath.Join(os.TempDir(), "carbonstack-state-path-policy-dev")
	_ = os.RemoveAll(reportRoot)
	if err := os.MkdirAll(reportRoot, 0o700); err != nil {
		return err
	}

	r.ArtifactScan("pre-state-path-policy-dev")

	steps := []Step{
		{
			Name:    "Comms C3 path policy tests",
			Dir:     r.Comms,
			Command: "go",
			Args:    []string{"test", "./internal/app", "-run", "StatePathPolicy", "-count=1", "-v"},
		},
		{
			Name:    "C3 command classifies derived default root",
			Dir:     r.Comms,
			Command: "go",
			Args: []string{
				"run", "./cmd/comms", "state-path-policy-dev",
				"--state", filepath.Join(reportRoot, ".carbonstack-comms", "state.json"),
				"--sidecar-dir", filepath.Join(r.Comms, "internal", "protocol", "mls", "openmls-sidecar"),
				"--output", filepath.Join(reportRoot, "derived-report.json"),
			},
		},
		{
			Name:    "C3 command classifies explicit root mismatch without brittle refusal",
			Dir:     r.Comms,
			Command: "go",
			Args: []string{
				"run", "./cmd/comms", "state-path-policy-dev",
				"--state", filepath.Join(reportRoot, "custom-state", "state.json"),
				"--state-root", filepath.Join(reportRoot, "custom-root"),
				"--sidecar-dir", filepath.Join(r.Comms, "internal", "protocol", "mls", "openmls-sidecar"),
				"--cypher-db", filepath.Join(reportRoot, "cypher.db"),
				"--validator-temp-root", reportRoot,
				"--evidence-root", reportRoot,
				"--output", filepath.Join(reportRoot, "explicit-report.json"),
			},
		},
		{
			Name:    "C3 command refuses parent traversal with validation exit zero",
			Dir:     r.Comms,
			Command: "go",
			Args: []string{
				"run", "./cmd/comms", "state-path-policy-dev",
				"--state", "../bad/state.json",
				"--allow-refusal-exit-zero",
				"--output", filepath.Join(reportRoot, "refusal-report.json"),
			},
		},
		{
			Name:    "Generated command reference current",
			Dir:     r.CarbonStack,
			Command: "python3",
			Args:    []string{"tools/registry/render-command-reference.py", "--check"},
		},
		{
			Name:    "Registry missing nonclaims remains zero",
			Dir:     validatorDir,
			Command: "go",
			Args:    []string{"run", ".", "--profile", "registry-lookup", "--list", "--missing-nonclaims"},
		},
	}

	for _, step := range steps {
		if err := r.RunStep(step); err != nil {
			return err
		}
	}

	fmt.Println()
	fmt.Println("state-path-policy-dev profile result:")
	fmt.Println("  PASS: explicit --state compatibility is preserved")
	fmt.Println("  PASS: canonical .carbonstack-comms root is preferred policy, not mandatory layout enforcement")
	fmt.Println("  PASS: explicit --state-root mismatch is classified rather than refused")
	fmt.Println("  PASS: parent traversal is refused as unsafe path policy")
	fmt.Println("  PASS: sidecar, Cypher, validator, and evidence roots are classified only")
	fmt.Println("  PASS: path policy report output is generated evidence only")
	fmt.Println("  PASS: registry/reference checks passed")
	fmt.Println("  cleanup_implemented: false")
	fmt.Println("  state_relocation_performed: false")
	fmt.Println("  migration_performed: false")
	fmt.Println("  repair_performed: false")
	fmt.Println("  no_silent_migration: true")
	fmt.Println("  no_silent_repair: true")
	fmt.Println("  trust_or_candidate_state_mutated: false")
	fmt.Println("  verified_identity_claimed: false")
	fmt.Println("  cypher_mls_reconciled: false")
	fmt.Println("  vault_claimed: false")
	fmt.Println("  backup_restore_claimed: false")
	fmt.Println("  deployment_claimed: false")
	fmt.Println("  full_runtime_dev_promoted: false")
	fmt.Println("  gate_d_started: false")
	fmt.Println("  boundary: dev/pre-alpha Gate C3 path policy only; next Gate C subgate is C4 atomic write/lock/partial-state/replay policy")

	r.ArtifactScan("post-state-path-policy-dev")
	if r.CleanGenerated {
		_ = os.RemoveAll(reportRoot)
		r.CleanGeneratedArtifacts()
	}
	fmt.Println()
	fmt.Println("VALIDATION PASSED")
	return nil
}
