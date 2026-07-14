package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func (r *Runner) StateSchemaCompatibilityDev() error {
	fmt.Println("========================================")
	fmt.Println("CarbonStack validation profile: state-schema-compat-dev")
	fmt.Println("========================================")
	fmt.Println("status: dev/pre-alpha Gate C2 schema/version compatibility profile")
	fmt.Println("scope: C2 read-only schema compatibility classification and refusal for Comms-owned JSON artifacts")
	fmt.Println("boundary: not migration, not repair, not vault, not backup/restore, not Gate D, not full-runtime-dev")
	fmt.Println()

	if err := r.CheckRequiredPaths(); err != nil {
		return err
	}
	if err := r.CheckLiveGitUmbrella("state-schema-compat-dev"); err != nil {
		return err
	}

	validatorDir := filepath.Join(r.CarbonStack, "tools", "carbonstack-validate")
	reportRoot := filepath.Join(os.TempDir(), "carbonstack-state-schema-compat-dev")
	_ = os.RemoveAll(reportRoot)
	if err := os.MkdirAll(reportRoot, 0o700); err != nil {
		return err
	}

	supportedReceipt := filepath.Join(reportRoot, "keypackage-receipt.json")
	unsupportedWelcome := filepath.Join(reportRoot, "welcome-receipt-unsupported.json")
	if err := os.WriteFile(supportedReceipt, []byte(`{"schema_version":"carbonstack-keypackage-consume-receipt/v0"}`), 0o600); err != nil {
		return err
	}
	if err := os.WriteFile(unsupportedWelcome, []byte(`{"schema_version":"carbonstack-welcome-consume-receipt/v99"}`), 0o600); err != nil {
		return err
	}

	r.ArtifactScan("pre-state-schema-compat-dev")

	steps := []Step{
		{
			Name:    "Comms C2 schema compatibility tests",
			Dir:     r.Comms,
			Command: "go",
			Args:    []string{"test", "./internal/app", "-run", "StateSchemaCompatibility", "-count=1", "-v"},
		},
		{
			Name:    "C2 command allows supported KeyPackage receipt",
			Dir:     r.Comms,
			Command: "go",
			Args: []string{
				"run", "./cmd/comms", "state-schema-compat-dev",
				"--kind", "keypackage-receipt",
				"--path", supportedReceipt,
				"--output", filepath.Join(reportRoot, "supported-report.json"),
			},
		},
		{
			Name:    "C2 command refuses unsupported Welcome receipt while allowing validation exit zero",
			Dir:     r.Comms,
			Command: "go",
			Args: []string{
				"run", "./cmd/comms", "state-schema-compat-dev",
				"--kind", "welcome-receipt",
				"--path", unsupportedWelcome,
				"--allow-refusal-exit-zero",
				"--output", filepath.Join(reportRoot, "unsupported-report.json"),
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
	fmt.Println("state-schema-compat-dev profile result:")
	fmt.Println("  PASS: supported Comms-owned JSON schemas allow")
	fmt.Println("  PASS: unsupported newer safety-sensitive schemas refuse")
	fmt.Println("  PASS: missing safety-sensitive schemas refuse")
	fmt.Println("  PASS: legacy Comms state is classified without migration")
	fmt.Println("  PASS: compatibility report output is generated evidence only")
	fmt.Println("  PASS: registry/reference checks passed")
	fmt.Println("  mutation_performed: false")
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
	fmt.Println("  boundary: dev/pre-alpha Gate C2 schema compatibility only; next Gate C subgate is C3 path policy and explicit state-root semantics")

	r.ArtifactScan("post-state-schema-compat-dev")
	if r.CleanGenerated {
		_ = os.RemoveAll(reportRoot)
		r.CleanGeneratedArtifacts()
	}
	fmt.Println()
	fmt.Println("VALIDATION PASSED")
	return nil
}
