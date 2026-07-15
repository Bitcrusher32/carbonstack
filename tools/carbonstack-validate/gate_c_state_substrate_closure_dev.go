package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func (r *Runner) GateCStateSubstrateClosureDev() error {
	fmt.Println("========================================")
	fmt.Println("CarbonStack validation profile: gate-c-state-substrate-closure-dev")
	fmt.Println("========================================")
	fmt.Println("status: dev/pre-alpha Gate C closure profile")
	fmt.Println("scope: Gate C C1-C4 integrated state-substrate closure")
	fmt.Println("boundary: not Gate D, not full-runtime-dev, not deployment, not vault, not backup/restore, not production E2EE")
	fmt.Println()

	if err := r.CheckRequiredPaths(); err != nil {
		return err
	}
	if err := r.CheckLiveGitUmbrella("gate-c-state-substrate-closure-dev"); err != nil {
		return err
	}

	validatorDir := filepath.Join(r.CarbonStack, "tools", "carbonstack-validate")
	reportRoot := filepath.Join(os.TempDir(), "carbonstack-gate-c-state-substrate-closure-dev")
	_ = os.RemoveAll(reportRoot)
	if err := os.MkdirAll(reportRoot, 0o700); err != nil {
		return err
	}

	r.ArtifactScan("pre-gate-c-state-substrate-closure-dev")

	pathPolicyFixture := filepath.Join(reportRoot, "path-policy-fixture.json")
	writePolicyFixture := filepath.Join(reportRoot, "write-policy-fixture.json")
	if err := os.WriteFile(pathPolicyFixture, []byte(`{"schema_version":"carbonstack-state-path-policy-report/v0"}`), 0o600); err != nil {
		return err
	}
	if err := os.WriteFile(writePolicyFixture, []byte(`{"schema_version":"carbonstack-state-write-policy-report/v0"}`), 0o600); err != nil {
		return err
	}

	steps := []Step{
		{
			Name:    "Comms Gate C C1-C4 focused tests",
			Dir:     r.Comms,
			Command: "go",
			Args:    []string{"test", "./internal/app", "-run", "StateSubstrateInventory|StateSchemaCompatibility|StatePathPolicy|StateWritePolicy", "-count=1", "-v"},
		},
		{
			Name:    "Validator package tests",
			Dir:     validatorDir,
			Command: "go",
			Args:    []string{"test", "./...", "-count=1"},
		},
		{
			Name:    "C1 state substrate inventory profile",
			Dir:     validatorDir,
			Command: "go",
			Args:    []string{"run", ".", "--profile", "state-substrate-inventory-dev", "--compact-summary", "--clean-generated"},
		},
		{
			Name:    "C2 schema compatibility profile",
			Dir:     validatorDir,
			Command: "go",
			Args:    []string{"run", ".", "--profile", "state-schema-compat-dev", "--compact-summary", "--clean-generated"},
		},
		{
			Name:    "C3 path policy profile",
			Dir:     validatorDir,
			Command: "go",
			Args:    []string{"run", ".", "--profile", "state-path-policy-dev", "--compact-summary", "--clean-generated"},
		},
		{
			Name:    "C4 write policy profile",
			Dir:     validatorDir,
			Command: "go",
			Args:    []string{"run", ".", "--profile", "state-write-policy-dev", "--compact-summary", "--clean-generated"},
		},
		{
			Name:    "C2 accepts C3 path-policy report schema through explicit kind",
			Dir:     r.Comms,
			Command: "go",
			Args: []string{
				"run", "./cmd/comms", "state-schema-compat-dev",
				"--kind", "path-policy-report",
				"--path", pathPolicyFixture,
				"--output", filepath.Join(reportRoot, "path-policy-compat-report.json"),
			},
		},
		{
			Name:    "C2 accepts C4 write-policy report schema through explicit kind",
			Dir:     r.Comms,
			Command: "go",
			Args: []string{
				"run", "./cmd/comms", "state-schema-compat-dev",
				"--kind", "write-policy-report",
				"--path", writePolicyFixture,
				"--output", filepath.Join(reportRoot, "write-policy-compat-report.json"),
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
		{
			Name:    "Gate B closure registry authority still present",
			Dir:     validatorDir,
			Command: "go",
			Args:    []string{"run", ".", "--profile", "registry-lookup", "--registry-id", "runner.gate-b-relay-lifecycle-closure-dev"},
		},
		{
			Name:    "C1 runner registry present",
			Dir:     validatorDir,
			Command: "go",
			Args:    []string{"run", ".", "--profile", "registry-lookup", "--registry-id", "runner.state-substrate-inventory-dev"},
		},
		{
			Name:    "C2 runner registry present",
			Dir:     validatorDir,
			Command: "go",
			Args:    []string{"run", ".", "--profile", "registry-lookup", "--registry-id", "runner.state-schema-compat-dev"},
		},
		{
			Name:    "C3 runner registry present",
			Dir:     validatorDir,
			Command: "go",
			Args:    []string{"run", ".", "--profile", "registry-lookup", "--registry-id", "runner.state-path-policy-dev"},
		},
		{
			Name:    "C4 runner registry present",
			Dir:     validatorDir,
			Command: "go",
			Args:    []string{"run", ".", "--profile", "registry-lookup", "--registry-id", "runner.state-write-policy-dev"},
		},
		{
			Name:    "C5 runner registry present",
			Dir:     validatorDir,
			Command: "go",
			Args:    []string{"run", ".", "--profile", "registry-lookup", "--registry-id", "runner.gate-c-state-substrate-closure-dev"},
		},
	}

	for _, step := range steps {
		if err := r.RunStep(step); err != nil {
			return err
		}
	}

	fmt.Println()
	fmt.Println("gate-c-state-substrate-closure-dev profile result:")
	fmt.Println("  PASS: C1 state-substrate inventory profile passes")
	fmt.Println("  PASS: C2 schema/version compatibility profile passes")
	fmt.Println("  PASS: C3 path policy profile passes")
	fmt.Println("  PASS: C4 write policy profile passes")
	fmt.Println("  PASS: C3 path-policy report schema is accepted by explicit C2 kind")
	fmt.Println("  PASS: C4 write-policy report schema is accepted by explicit C2 kind")
	fmt.Println("  PASS: registry/reference checks passed")
	fmt.Println("  PASS: missing nonclaims remain zero")
	fmt.Println("  PASS: Gate B closure authority remains present without rerunning the Gate B closure ladder")
	fmt.Println("  gate_c_status: closed")
	fmt.Println("  gate_d_status: not_started")
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
	fmt.Println("  boundary: dev/pre-alpha Gate C closure only; Gate D remains blocked until a fresh Gate D contract/recon is accepted")

	r.ArtifactScan("post-gate-c-state-substrate-closure-dev")
	if r.CleanGenerated {
		_ = os.RemoveAll(reportRoot)
		r.CleanGeneratedArtifacts()
	}
	fmt.Println()
	fmt.Println("VALIDATION PASSED")
	return nil
}
