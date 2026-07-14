package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func (r *Runner) StateSubstrateInventoryDev() error {
	fmt.Println("========================================")
	fmt.Println("CarbonStack validation profile: state-substrate-inventory-dev")
	fmt.Println("========================================")
	fmt.Println("status: dev/pre-alpha Gate C1 state substrate inventory profile")
	fmt.Println("scope: C1 read-only-by-default state authority map and machine-readable inventory")
	fmt.Println("boundary: not schema enforcement, not migration, not repair, not vault, not backup/restore, not Gate D")
	fmt.Println()

	if err := r.CheckRequiredPaths(); err != nil {
		return err
	}
	if err := r.CheckLiveGitUmbrella("state-substrate-inventory-dev"); err != nil {
		return err
	}

	validatorDir := filepath.Join(r.CarbonStack, "tools", "carbonstack-validate")
	reportRoot := filepath.Join(os.TempDir(), "carbonstack-state-substrate-inventory-dev")
	reportPath := filepath.Join(reportRoot, "inventory.json")
	_ = os.RemoveAll(reportRoot)

	r.ArtifactScan("pre-state-substrate-inventory-dev")

	steps := []Step{
		{
			Name:    "Comms C1 state substrate inventory tests",
			Dir:     r.Comms,
			Command: "go",
			Args:    []string{"test", "./internal/app", "-run", "StateSubstrateInventory", "-count=1", "-v"},
		},
		{
			Name:    "C1 command smoke writes generated evidence report",
			Dir:     r.Comms,
			Command: "go",
			Args: []string{
				"run", "./cmd/comms", "state-substrate-inventory-dev",
				"--state", filepath.Join(reportRoot, ".carbonstack-comms", "state.json"),
				"--sidecar-dir", filepath.Join(r.Comms, "internal", "protocol", "mls", "openmls-sidecar"),
				"--output", reportPath,
			},
		},
		{
			Name:    "Gate B closure remains valid before Gate C mutation",
			Dir:     validatorDir,
			Command: "go",
			Args:    []string{"run", ".", "--profile", "gate-b-relay-lifecycle-closure-dev", "--compact-summary", "--clean-generated"},
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

	if _, err := os.Stat(reportPath); err != nil {
		return fmt.Errorf("expected C1 generated inventory report at %s: %w", reportPath, err)
	}

	fmt.Println()
	fmt.Println("state-substrate-inventory-dev profile result:")
	fmt.Println("  PASS: C1 inventory command is read-only by default")
	fmt.Println("  PASS: optional --output writes machine-readable generated evidence")
	fmt.Println("  PASS: explicit --state compatibility is preserved")
	fmt.Println("  PASS: canonical Comms-owned root is a policy anchor, not a brittle mandatory chokepoint")
	fmt.Println("  PASS: Comms-owned, sidecar-owned, Cypher-owned, validator/evidence-generated, legacy/unversioned, and unknown state classes are representable")
	fmt.Println("  PASS: supported receipt/workflow schemas are detected")
	fmt.Println("  PASS: unsupported schemas are classified for future C2 refusal")
	fmt.Println("  PASS: Gate B closure profile remains valid before deeper Gate C mutation")
	fmt.Println("  no_silent_repair: true")
	fmt.Println("  no_silent_migration: true")
	fmt.Println("  trust_or_candidate_state_mutated: false")
	fmt.Println("  verified_identity_claimed: false")
	fmt.Println("  cypher_mls_reconciled: false")
	fmt.Println("  vault_claimed: false")
	fmt.Println("  backup_restore_claimed: false")
	fmt.Println("  deployment_claimed: false")
	fmt.Println("  full_runtime_dev_promoted: false")
	fmt.Println("  gate_d_started: false")
	fmt.Println("  boundary: dev/pre-alpha Gate C1 inventory only; next Gate C subgate is C2 schema/version compatibility refusal")

	r.ArtifactScan("post-state-substrate-inventory-dev")
	if r.CleanGenerated {
		_ = os.RemoveAll(reportRoot)
		r.CleanGeneratedArtifacts()
	}
	fmt.Println()
	fmt.Println("VALIDATION PASSED")
	return nil
}
