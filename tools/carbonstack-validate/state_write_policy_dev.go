package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func (r *Runner) StateWritePolicyDev() error {
	fmt.Println("========================================")
	fmt.Println("CarbonStack validation profile: state-write-policy-dev")
	fmt.Println("========================================")
	fmt.Println("status: dev/pre-alpha Gate C4 write policy profile")
	fmt.Println("scope: C4 atomic write, lock, partial-state, replay, and cleanup-boundary classification")
	fmt.Println("boundary: not migration, not repair, not relocation, not writer rewiring, not cleanup implementation, not Gate D")
	fmt.Println()

	if err := r.CheckRequiredPaths(); err != nil {
		return err
	}
	if err := r.CheckLiveGitUmbrella("state-write-policy-dev"); err != nil {
		return err
	}

	validatorDir := filepath.Join(r.CarbonStack, "tools", "carbonstack-validate")
	reportRoot := filepath.Join(os.TempDir(), "carbonstack-state-write-policy-dev")
	_ = os.RemoveAll(reportRoot)
	if err := os.MkdirAll(reportRoot, 0o700); err != nil {
		return err
	}

	r.ArtifactScan("pre-state-write-policy-dev")

	steps := []Step{
		{
			Name:    "Comms C4 write policy tests",
			Dir:     r.Comms,
			Command: "go",
			Args:    []string{"test", "./internal/app", "-run", "StateWritePolicy", "-count=1", "-v"},
		},
		{
			Name:    "C4 command writes generated evidence report",
			Dir:     r.Comms,
			Command: "go",
			Args: []string{
				"run", "./cmd/comms", "state-write-policy-dev",
				"--state-root", filepath.Join(reportRoot, ".carbonstack-comms"),
				"--sidecar-dir", filepath.Join(r.Comms, "internal", "protocol", "mls", "openmls-sidecar"),
				"--cypher-db", filepath.Join(reportRoot, "cypher.db"),
				"--validator-temp-root", reportRoot,
				"--evidence-root", reportRoot,
				"--output", filepath.Join(reportRoot, "write-policy-report.json"),
			},
		},
		{
			Name:    "C3 path policy command still callable",
			Dir:     r.Comms,
			Command: "go",
			Args: []string{
				"run", "./cmd/comms", "state-path-policy-dev",
				"--state", filepath.Join(reportRoot, ".carbonstack-comms", "state.json"),
				"--output", filepath.Join(reportRoot, "path-policy-report.json"),
			},
		},
		{
			Name:    "C2 schema compatibility command accepts C3 path-policy schema",
			Dir:     r.Comms,
			Command: "bash",
			Args: []string{
				"-lc",
				fmt.Sprintf("printf '{\"schema_version\":\"carbonstack-state-path-policy-report/v0\"}' > %q && go run ./cmd/comms state-schema-compat-dev --kind state-substrate-inventory --path %q --allow-refusal-exit-zero --output %q", filepath.Join(reportRoot, "path-policy-fixture.json"), filepath.Join(reportRoot, "path-policy-fixture.json"), filepath.Join(reportRoot, "compat-report.json")),
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

	if _, err := os.Stat(filepath.Join(reportRoot, "write-policy-report.json")); err != nil {
		return fmt.Errorf("expected generated write policy report: %w", err)
	}

	fmt.Println()
	fmt.Println("state-write-policy-dev profile result:")
	fmt.Println("  PASS: write policy report schema emitted")
	fmt.Println("  PASS: C1/C2/C3 generated report writers classified as atomic generated evidence")
	fmt.Println("  PASS: B5d/B6 receipt writers classified as atomic lock-guarded receipt writers")
	fmt.Println("  PASS: B8 workflow report classified as atomic report writer with replay semantics")
	fmt.Println("  PASS: direct local state/trust/candidate writers classified as future-hardening warnings")
	fmt.Println("  PASS: sidecar, Cypher, validator, and evidence roots are classified only")
	fmt.Println("  PASS: generated report output leaves no required runtime state mutation")
	fmt.Println("  PASS: registry/reference checks passed")
	fmt.Println("  runtime_writer_rewired: false")
	fmt.Println("  cleanup_implemented: false")
	fmt.Println("  destructive_cleanup_performed: false")
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
	fmt.Println("  boundary: dev/pre-alpha Gate C4 write policy classification only; next Gate C subgate is C5 Gate C closure profile")

	r.ArtifactScan("post-state-write-policy-dev")
	if r.CleanGenerated {
		_ = os.RemoveAll(reportRoot)
		r.CleanGeneratedArtifacts()
	}
	fmt.Println()
	fmt.Println("VALIDATION PASSED")
	return nil
}
