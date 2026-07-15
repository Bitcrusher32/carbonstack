package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

const gateENativeDeploymentClosureReportSchema = "carbonstack-gate-e-native-deployment-closure-report/v0"

func (r *Runner) GateENativeDeploymentClosureDev() error {
	fmt.Println("========================================")
	fmt.Println("CarbonStack validation profile: gate-e-native-deployment-closure-dev")
	fmt.Println("========================================")
	fmt.Println("status: dev/pre-alpha Gate E closure profile")
	fmt.Println("scope: Gate E E1 manual-private deployment profile plus E2 Cypher terminating config inspection plus registry/nonclaim closure")
	fmt.Println("boundary: not service/systemd, not helper install, not container, not public ingress, not TUI, not full-runtime-dev, not verified identity, not trust promotion")
	fmt.Println()

	if err := r.CheckRequiredPaths(); err != nil {
		return err
	}
	if err := r.CheckLiveGitUmbrella("gate-e-native-deployment-closure-dev"); err != nil {
		return err
	}

	validatorDir := filepath.Join(r.CarbonStack, "tools", "carbonstack-validate")
	reportRoot := filepath.Join(os.TempDir(), "carbonstack-gate-e-native-deployment-closure-dev")
	_ = os.RemoveAll(reportRoot)
	configRoot := filepath.Join(reportRoot, "config")
	stateRoot := filepath.Join(reportRoot, "state")
	migrationsRoot := filepath.Join(reportRoot, "migrations")
	evidenceRoot := filepath.Join(reportRoot, "evidence")

	for _, dir := range []string{configRoot, stateRoot, migrationsRoot, evidenceRoot} {
		if err := os.MkdirAll(dir, 0o700); err != nil {
			return err
		}
	}
	if err := os.WriteFile(filepath.Join(migrationsRoot, "001_test.sql"), []byte("select 1;\n"), 0o600); err != nil {
		return err
	}

	cypherDB := filepath.Join(stateRoot, "cypher.db")
	printConfigReport := filepath.Join(evidenceRoot, "cypher-print-config.json")
	closureReportPath := filepath.Join(evidenceRoot, "gate-e-native-deployment-closure-report.json")

	r.ArtifactScan("pre-gate-e-native-deployment-closure-dev")

	steps := []Step{
		{
			Name:    "Gate E E1 manual-private native deployment profile passes",
			Dir:     validatorDir,
			Command: "go",
			Args:    []string{"run", ".", "--profile", "gate-e-native-deployment-dev", "--compact-summary", "--clean-generated"},
		},
		{
			Name:    "Gate E E1 registry authority present",
			Dir:     validatorDir,
			Command: "go",
			Args:    []string{"run", ".", "--profile", "registry-lookup", "--registry-id", "runner.gate-e-native-deployment-dev"},
		},
		{
			Name:    "Gate E E2 Cypher config-inspection registry authority present",
			Dir:     validatorDir,
			Command: "go",
			Args:    []string{"run", ".", "--profile", "registry-lookup", "--registry-id", "cypher.config-inspection"},
		},
	}

	for _, step := range steps {
		if err := r.RunStep(step); err != nil {
			return err
		}
	}

	env := append(os.Environ(),
		"CYPHER_ADDR=127.0.0.1:19090",
		"CYPHER_DB="+cypherDB,
		"CYPHER_MIGRATIONS="+migrationsRoot,
		"CYPHER_DEV_INVITE=gate-e-closure-dev",
	)

	printCmd := exec.Command("go", "run", "./cmd/cypher", "--print-config")
	printCmd.Dir = r.Cypher
	printCmd.Env = env
	printOut, err := printCmd.CombinedOutput()
	fmt.Println("cypher --print-config output:")
	fmt.Print(string(printOut))
	if err != nil {
		return fmt.Errorf("cypher --print-config failed: %w", err)
	}
	if err := os.WriteFile(printConfigReport, printOut, 0o600); err != nil {
		return err
	}

	var parsed map[string]any
	if err := json.Unmarshal(printOut, &parsed); err != nil {
		return fmt.Errorf("cypher --print-config did not emit JSON: %w", err)
	}
	if parsed["schema_version"] != "carbonstack-cypher-config-inspection/v0" {
		return fmt.Errorf("unexpected cypher config schema: %v", parsed["schema_version"])
	}
	if parsed["starts_server"] != false {
		return fmt.Errorf("cypher --print-config reported starts_server=%v", parsed["starts_server"])
	}
	if parsed["terminating_inspection"] != true {
		return fmt.Errorf("cypher --print-config reported terminating_inspection=%v", parsed["terminating_inspection"])
	}
	if parsed["db_path_source"] != "env" {
		return fmt.Errorf("cypher --print-config did not use env DB source: %v", parsed["db_path_source"])
	}
	fmt.Println("PASS: Cypher --print-config terminates and emits expected inspection schema")

	checkCmd := exec.Command("go", "run", "./cmd/cypher", "--check-config")
	checkCmd.Dir = r.Cypher
	checkCmd.Env = env
	checkOut, err := checkCmd.CombinedOutput()
	fmt.Println("cypher --check-config output:")
	fmt.Print(string(checkOut))
	if err != nil {
		return fmt.Errorf("cypher --check-config failed: %w", err)
	}
	fmt.Println("PASS: Cypher --check-config terminates and validates explicit Gate E style env")

	postSteps := []Step{
		{
			Name:    "Generated command reference current",
			Dir:     r.CarbonStack,
			Command: "python3",
			Args:    []string{"tools/registry/render-command-reference.py", "--check"},
		},
		{
			Name:    "Gate E closure registry authority present",
			Dir:     validatorDir,
			Command: "go",
			Args:    []string{"run", ".", "--profile", "registry-lookup", "--registry-id", "runner.gate-e-native-deployment-closure-dev"},
		},
		{
			Name:    "Registry missing nonclaims remains zero",
			Dir:     validatorDir,
			Command: "go",
			Args:    []string{"run", ".", "--profile", "registry-lookup", "--list", "--missing-nonclaims"},
		},
	}
	for _, step := range postSteps {
		if err := r.RunStep(step); err != nil {
			return err
		}
	}

	report := map[string]any{
		"schema_version":                     gateENativeDeploymentClosureReportSchema,
		"profile":                            "gate-e-native-deployment-closure-dev",
		"created_at":                         time.Now().UTC().Format(time.RFC3339),
		"gate_d_status":                      "closed",
		"gate_e_status":                      "closed",
		"gate_e_e1_status":                   "closed",
		"gate_e_e2_status":                   "closed",
		"gate_e_e3_status":                   "closed",
		"gate_f_status":                      "not_started",
		"manual_private_deployment_first":    true,
		"semi_persistent_deployment_started": false,
		"service_or_systemd_started":         false,
		"helper_install_started":             false,
		"container_started":                  false,
		"public_ingress_started":             false,
		"tui_started":                        false,
		"full_runtime_dev_promoted":          false,
		"verified_identity_claimed":          false,
		"trust_promotion_claimed":            false,
		"vault_claimed":                      false,
		"backup_restore_claimed":             false,
		"production_e2ee_claimed":            false,
		"pq_hybrid_claimed":                  false,
		"android_claimed":                    false,
		"carbonstack_os_claimed":             false,
		"closure_evidence": map[string]any{
			"gate_e_e1_profile_passed":                      true,
			"cypher_print_config_terminates":                true,
			"cypher_check_config_terminates":                true,
			"cypher_config_inspection_schema_supported":     true,
			"gate_e_docs_registry_reference_current":        true,
			"missing_nonclaims_zero":                        true,
			"manual_private_deployment_model_documented":    true,
			"service_systemd_helper_explicitly_not_started": true,
		},
		"reports": map[string]any{
			"cypher_print_config": printConfigReport,
		},
	}
	if err := writeGateEClosureJSON(closureReportPath, report); err != nil {
		return err
	}

	fmt.Println()
	fmt.Println("gate-e-native-deployment-closure-dev profile result:")
	fmt.Println("  PASS: Gate E E1 manual-private native deployment profile passes")
	fmt.Println("  PASS: E2 Cypher --print-config terminates and emits inspection schema")
	fmt.Println("  PASS: E2 Cypher --check-config terminates and validates explicit env")
	fmt.Println("  PASS: registry/reference/nonclaim checks passed")
	fmt.Println("  gate_d_status: closed")
	fmt.Println("  gate_e_status: closed")
	fmt.Println("  gate_e_e1_status: closed")
	fmt.Println("  gate_e_e2_status: closed")
	fmt.Println("  gate_e_e3_status: closed")
	fmt.Println("  gate_f_status: not_started")
	fmt.Println("  closure_report:", closureReportPath)
	fmt.Println("  manual_private_deployment_first: true")
	fmt.Println("  semi_persistent_deployment_started: false")
	fmt.Println("  service_or_systemd_started: false")
	fmt.Println("  helper_install_started: false")
	fmt.Println("  container_started: false")
	fmt.Println("  public_ingress_started: false")
	fmt.Println("  tui_started: false")
	fmt.Println("  full_runtime_dev_promoted: false")
	fmt.Println("  verified_identity_claimed: false")
	fmt.Println("  trust_promotion_claimed: false")
	fmt.Println("  production_e2ee_claimed: false")
	fmt.Println("  boundary: Gate E closed for manual-private native deployment only; Gate F requires fresh preflight/contract")

	r.ArtifactScan("post-gate-e-native-deployment-closure-dev")
	if r.CleanGenerated {
		_ = os.RemoveAll(reportRoot)
		r.CleanGeneratedArtifacts()
	}
	fmt.Println()
	fmt.Println("VALIDATION PASSED")
	return nil
}

func writeGateEClosureJSON(path string, value any) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		return err
	}
	body, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, append(body, '\n'), 0o600)
}
