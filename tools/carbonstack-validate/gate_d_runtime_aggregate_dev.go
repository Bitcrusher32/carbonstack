package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

const gateDRuntimeAggregateReportSchema = "carbonstack-gate-d-runtime-aggregate-report/v0"
const gateDRunSpaceContextCandidateSchema = "carbonstack-run-space-context-candidate/v0"

func (r *Runner) GateDRuntimeAggregateDev() error {
	fmt.Println("========================================")
	fmt.Println("CarbonStack validation profile: gate-d-runtime-aggregate-dev")
	fmt.Println("========================================")
	fmt.Println("status: dev/pre-alpha Gate D runtime aggregate profile")
	fmt.Println("scope: profile-only mechanical runtime aggregate over Gate B Relay lifecycle and Gate C state substrate")
	fmt.Println("boundary: not full-runtime-dev, not TUI, not Gate E deployment, not verified identity, not trust promotion, not production E2EE")
	fmt.Println()

	if err := r.CheckRequiredPaths(); err != nil {
		return err
	}
	if err := r.CheckLiveGitUmbrella("gate-d-runtime-aggregate-dev"); err != nil {
		return err
	}

	validatorDir := filepath.Join(r.CarbonStack, "tools", "carbonstack-validate")
	reportRoot := filepath.Join(os.TempDir(), "carbonstack-gate-d-runtime-aggregate-dev")
	_ = os.RemoveAll(reportRoot)
	if err := os.MkdirAll(filepath.Join(reportRoot, ".carbonstack-comms"), 0o700); err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Join(reportRoot, "evidence"), 0o700); err != nil {
		return err
	}

	statePath := filepath.Join(reportRoot, ".carbonstack-comms", "state.json")
	stateRoot := filepath.Join(reportRoot, ".carbonstack-comms")
	sidecarDir := filepath.Join(r.Comms, "internal", "protocol", "mls", "openmls-sidecar")
	cypherDB := filepath.Join(reportRoot, "cypher.db")
	evidenceRoot := filepath.Join(reportRoot, "evidence")

	runSpaceContextPath := filepath.Join(reportRoot, "run-space.context.json")
	runtimeAggregateReportPath := filepath.Join(reportRoot, "gate-d-runtime-aggregate-report.json")
	c1InventoryReport := filepath.Join(evidenceRoot, "c1-inventory.json")
	c2C1CompatReport := filepath.Join(evidenceRoot, "c2-c1-compat.json")
	c2C3CompatReport := filepath.Join(evidenceRoot, "c2-c3-compat.json")
	c2C4CompatReport := filepath.Join(evidenceRoot, "c2-c4-compat.json")
	c3PathPolicyReport := filepath.Join(evidenceRoot, "c3-path-policy.json")
	c4WritePolicyReport := filepath.Join(evidenceRoot, "c4-write-policy.json")

	runSpaceContext := map[string]any{
		"schema_version":                  gateDRunSpaceContextCandidateSchema,
		"profile_name":                    "gate-d-runtime-aggregate-dev",
		"run_space_root":                  reportRoot,
		"server_url":                      "http://127.0.0.1:1",
		"account_id":                      "acct-gate-d-runtime-aggregate-dev",
		"device_id":                       "dev-gate-d-runtime-aggregate-dev",
		"relay_space_id":                  "relay-space-gate-d-runtime-aggregate-dev",
		"state_path":                      statePath,
		"state_root":                      stateRoot,
		"sidecar_dir":                     sidecarDir,
		"cypher_db":                       cypherDB,
		"evidence_root":                   evidenceRoot,
		"authority":                       "explicit context candidate only; not hidden ambient authority",
		"tui_claimed":                     false,
		"deployment_claimed":              false,
		"full_runtime_dev_promoted":       false,
		"full_runtime_dev_registry_added": false,
	}
	if err := writeGateDJSON(runSpaceContextPath, runSpaceContext); err != nil {
		return err
	}
	if err := writeGateDJSON(statePath, map[string]any{
		"server_url": "http://127.0.0.1:1",
		"account_id": "acct-gate-d-runtime-aggregate-dev",
		"device_id":  "dev-gate-d-runtime-aggregate-dev",
	}); err != nil {
		return err
	}

	r.ArtifactScan("pre-gate-d-runtime-aggregate-dev")

	steps := []Step{
		{
			Name:    "Gate C closure profile",
			Dir:     validatorDir,
			Command: "go",
			Args:    []string{"run", ".", "--profile", "gate-c-state-substrate-closure-dev", "--compact-summary", "--clean-generated"},
		},
		{
			Name:    "C1 inventory over explicit Gate D run-space context",
			Dir:     r.Comms,
			Command: "go",
			Args: []string{
				"run", "./cmd/comms", "state-substrate-inventory-dev",
				"--state", statePath,
				"--state-root", stateRoot,
				"--sidecar-dir", sidecarDir,
				"--cypher-db", cypherDB,
				"--evidence-root", evidenceRoot,
				"--output", c1InventoryReport,
			},
		},
		{
			Name:    "C3 path policy over explicit Gate D run-space context",
			Dir:     r.Comms,
			Command: "go",
			Args: []string{
				"run", "./cmd/comms", "state-path-policy-dev",
				"--state", statePath,
				"--state-root", stateRoot,
				"--sidecar-dir", sidecarDir,
				"--cypher-db", cypherDB,
				"--validator-temp-root", reportRoot,
				"--evidence-root", evidenceRoot,
				"--output", c3PathPolicyReport,
			},
		},
		{
			Name:    "C4 write policy over explicit Gate D run-space context",
			Dir:     r.Comms,
			Command: "go",
			Args: []string{
				"run", "./cmd/comms", "state-write-policy-dev",
				"--state-root", stateRoot,
				"--sidecar-dir", sidecarDir,
				"--cypher-db", cypherDB,
				"--validator-temp-root", reportRoot,
				"--evidence-root", evidenceRoot,
				"--output", c4WritePolicyReport,
			},
		},
		{
			Name:    "C2 validates C1 inventory report",
			Dir:     r.Comms,
			Command: "go",
			Args: []string{
				"run", "./cmd/comms", "state-schema-compat-dev",
				"--kind", "state-substrate-inventory",
				"--path", c1InventoryReport,
				"--output", c2C1CompatReport,
			},
		},
		{
			Name:    "C2 validates C3 path-policy report",
			Dir:     r.Comms,
			Command: "go",
			Args: []string{
				"run", "./cmd/comms", "state-schema-compat-dev",
				"--kind", "path-policy-report",
				"--path", c3PathPolicyReport,
				"--output", c2C3CompatReport,
			},
		},
		{
			Name:    "C2 validates C4 write-policy report",
			Dir:     r.Comms,
			Command: "go",
			Args: []string{
				"run", "./cmd/comms", "state-schema-compat-dev",
				"--kind", "write-policy-report",
				"--path", c4WritePolicyReport,
				"--output", c2C4CompatReport,
			},
		},
		{
			Name:    "Workflow relay onboarding profile",
			Dir:     validatorDir,
			Command: "go",
			Args:    []string{"run", ".", "--profile", "workflow-relay-onboarding-dev", "--compact-summary", "--clean-generated"},
		},
		{
			Name:    "Relay Space member restart/resume inspection profile",
			Dir:     validatorDir,
			Command: "go",
			Args:    []string{"run", ".", "--profile", "relay-space-member-restart-dev", "--compact-summary", "--clean-generated"},
		},
		{
			Name:    "Same-state integrated KeyPackage/Welcome/message send/inbox proof",
			Dir:     validatorDir,
			Command: "go",
			Args:    []string{"run", ".", "--profile", "same-state-integrated-dev", "--compact-summary", "--clean-generated"},
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
			Name:    "Gate B closure authority present without rerun",
			Dir:     validatorDir,
			Command: "go",
			Args:    []string{"run", ".", "--profile", "registry-lookup", "--registry-id", "runner.gate-b-relay-lifecycle-closure-dev"},
		},
		{
			Name:    "Gate C closure authority present",
			Dir:     validatorDir,
			Command: "go",
			Args:    []string{"run", ".", "--profile", "registry-lookup", "--registry-id", "runner.gate-c-state-substrate-closure-dev"},
		},
		{
			Name:    "Workflow authority present",
			Dir:     validatorDir,
			Command: "go",
			Args:    []string{"run", ".", "--profile", "registry-lookup", "--registry-id", "runner.workflow-relay-onboarding-dev"},
		},
		{
			Name:    "Message send wrapper authority present",
			Dir:     validatorDir,
			Command: "go",
			Args:    []string{"run", ".", "--profile", "registry-lookup", "--registry-id", "comms.message-send-dev"},
		},
		{
			Name:    "Message inbox wrapper authority present",
			Dir:     validatorDir,
			Command: "go",
			Args:    []string{"run", ".", "--profile", "registry-lookup", "--registry-id", "comms.message-inbox-dev"},
		},
		{
			Name:    "Gate D runtime aggregate registry authority present",
			Dir:     validatorDir,
			Command: "go",
			Args:    []string{"run", ".", "--profile", "registry-lookup", "--registry-id", "runner.gate-d-runtime-aggregate-dev"},
		},
	}

	for _, step := range steps {
		if err := r.RunStep(step); err != nil {
			return err
		}
	}

	report := map[string]any{
		"schema_version":                       gateDRuntimeAggregateReportSchema,
		"profile":                              "gate-d-runtime-aggregate-dev",
		"created_at":                           time.Now().UTC().Format(time.RFC3339),
		"run_space_context_schema":             gateDRunSpaceContextCandidateSchema,
		"run_space_context_path":               runSpaceContextPath,
		"run_space_root":                       reportRoot,
		"state_path":                           statePath,
		"state_root":                           stateRoot,
		"sidecar_dir":                          sidecarDir,
		"cypher_db":                            cypherDB,
		"evidence_root":                        evidenceRoot,
		"gate_d_status":                        "closed",
		"gate_e_status":                        "not_started",
		"full_runtime_dev_promoted":            false,
		"full_runtime_dev_registry_name_added": false,
		"tui_started":                          false,
		"deployment_started":                   false,
		"verified_identity_claimed":            false,
		"trust_promotion_claimed":              false,
		"vault_claimed":                        false,
		"backup_restore_claimed":               false,
		"production_e2ee_claimed":              false,
		"pq_hybrid_claimed":                    false,
		"android_claimed":                      false,
		"carbonstack_os_claimed":               false,
		"reports": map[string]any{
			"c1_inventory":    c1InventoryReport,
			"c2_c1_compat":    c2C1CompatReport,
			"c2_c3_compat":    c2C3CompatReport,
			"c2_c4_compat":    c2C4CompatReport,
			"c3_path_policy":  c3PathPolicyReport,
			"c4_write_policy": c4WritePolicyReport,
		},
		"closure_evidence": map[string]any{
			"gate_c_closure_profile_passed":                  true,
			"explicit_run_space_context_inspected":           true,
			"state_substrate_preflight_passed":               true,
			"workflow_relay_onboarding_profile_passed":       true,
			"relay_space_restart_resume_profile_passed":      true,
			"same_state_integrated_message_lifecycle_passed": true,
			"minimal_message_send_inbox_included":            true,
			"registry_reference_current":                     true,
			"missing_nonclaims_zero":                         true,
			"gate_b_closure_authority_present_without_rerun": true,
		},
	}
	if err := writeGateDJSON(runtimeAggregateReportPath, report); err != nil {
		return err
	}

	fmt.Println()
	fmt.Println("gate-d-runtime-aggregate-dev profile result:")
	fmt.Println("  PASS: Gate C closure profile passes")
	fmt.Println("  PASS: explicit run-space/context candidate generated and inspected")
	fmt.Println("  PASS: C1/C2/C3/C4 state-substrate preflight passes over explicit run-space paths")
	fmt.Println("  PASS: workflow relay onboarding profile passes")
	fmt.Println("  PASS: Relay Space member restart/resume inspection profile passes")
	fmt.Println("  PASS: same-state integrated proof covers KeyPackage -> Welcome -> message-send-dev -> message-inbox-dev --ack")
	fmt.Println("  PASS: registry/reference checks passed")
	fmt.Println("  PASS: missing nonclaims remain zero")
	fmt.Println("  PASS: Gate B closure authority remains present without rerunning the Gate B closure ladder")
	fmt.Println("  gate_d_status: closed")
	fmt.Println("  gate_e_status: not_started")
	fmt.Println("  runtime_aggregate_report:", runtimeAggregateReportPath)
	fmt.Println("  run_space_context:", runSpaceContextPath)
	fmt.Println("  full_runtime_dev_promoted: false")
	fmt.Println("  full_runtime_dev_registry_name_added: false")
	fmt.Println("  tui_started: false")
	fmt.Println("  deployment_started: false")
	fmt.Println("  verified_identity_claimed: false")
	fmt.Println("  trust_promotion_claimed: false")
	fmt.Println("  vault_claimed: false")
	fmt.Println("  backup_restore_claimed: false")
	fmt.Println("  production_e2ee_claimed: false")
	fmt.Println("  pq_hybrid_claimed: false")
	fmt.Println("  android_claimed: false")
	fmt.Println("  carbonstack_os_claimed: false")
	fmt.Println("  boundary: dev/pre-alpha Gate D runtime aggregate only; full-runtime-dev remains reserved until a later explicit promotion decision")

	r.ArtifactScan("post-gate-d-runtime-aggregate-dev")
	if r.CleanGenerated {
		_ = os.RemoveAll(reportRoot)
		r.CleanGeneratedArtifacts()
	}
	fmt.Println()
	fmt.Println("VALIDATION PASSED")
	return nil
}

func writeGateDJSON(path string, value any) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		return err
	}
	body, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, append(body, '\n'), 0o600)
}
