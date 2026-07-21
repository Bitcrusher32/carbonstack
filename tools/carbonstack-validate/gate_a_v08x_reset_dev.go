package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

const gateAV08xResetReportSchema = "carbonstack-gate-a-v08x-reset-report/v0"

func (r *Runner) GateAV08xResetDev() error {
	fmt.Println("========================================")
	fmt.Println("CarbonStack validation profile: gate-a-v08x-reset-dev")
	fmt.Println("========================================")
	fmt.Println("status: dev/pre-alpha v0.8.x Gate A reset and LogDocV2 workflow doctrine closure")
	fmt.Println("scope: external planning-authority acceptance, source-head freeze, active v0.8.x ledger reset contract, and Gate B block")
	fmt.Println("boundary: not adversarial harness implementation, not trust promotion, not full-runtime-dev promotion, not release creation")
	fmt.Println()

	required := []string{
		filepath.Join(r.CarbonStack, "README.md"),
		filepath.Join(r.CarbonStack, "registry", "commands.v0.yaml"),
		filepath.Join(r.CarbonStack, "registry", "COMMAND_REFERENCE.v0.md"),
		filepath.Join(r.CarbonStack, "registry", "COMMAND_BOUNDARY_TABLE.v0.md"),
		filepath.Join(r.CarbonStack, "docs", "README.md"),
		filepath.Join(r.CarbonStack, "docs", "288-v0.8.1-gate-a-post-v0.8.0-reset-workflow-doctrine-v0.md"),
		filepath.Join(r.CarbonStack, "docs", "289-v0.8.1-gate-a-closure-v0.md"),
		filepath.Join(r.CarbonStack, "sanitized-project-logdoc-list", "[SANITIZED]CarbonStackLogDocV0.8.0PRIME.md"),
	}
	for _, path := range required {
		if _, err := os.Stat(path); err != nil {
			return fmt.Errorf("Gate A required path missing: %s: %w", path, err)
		}
	}

	registryBytes, err := os.ReadFile(filepath.Join(r.CarbonStack, "registry", "commands.v0.yaml"))
	if err != nil {
		return err
	}
	referenceBytes, err := os.ReadFile(filepath.Join(r.CarbonStack, "registry", "COMMAND_REFERENCE.v0.md"))
	if err != nil {
		return err
	}
	boundaryBytes, err := os.ReadFile(filepath.Join(r.CarbonStack, "registry", "COMMAND_BOUNDARY_TABLE.v0.md"))
	if err != nil {
		return err
	}
	doc1Bytes, err := os.ReadFile(filepath.Join(r.CarbonStack, "docs", "288-v0.8.1-gate-a-post-v0.8.0-reset-workflow-doctrine-v0.md"))
	if err != nil {
		return err
	}
	doc2Bytes, err := os.ReadFile(filepath.Join(r.CarbonStack, "docs", "289-v0.8.1-gate-a-closure-v0.md"))
	if err != nil {
		return err
	}

	registryText := string(registryBytes)
	referenceText := string(referenceBytes)
	boundaryText := string(boundaryBytes)
	docText := string(doc1Bytes) + "\n" + string(doc2Bytes)

	if !gateARegistryIDPresent(registryText, "runner.gate-a-v08x-reset-dev") {
		return fmt.Errorf("Gate A registry ID missing")
	}
	if !strings.Contains(referenceText, "runner.gate-a-v08x-reset-dev") {
		return fmt.Errorf("Gate A command reference entry missing")
	}
	if !strings.Contains(referenceText, "Registry entry count: **141**") {
		return fmt.Errorf("command reference does not show expected 141 registry entries after Gate A addition")
	}
	if gateARegistryIDPresent(registryText, "runner.gate-f-basic-local-trust-dev") {
		return fmt.Errorf("stale shortened Gate F trust runner ID is present; expected runner.gate-f-basic-local-trust-posture-dev")
	}
	if !gateARegistryIDPresent(registryText, "runner.gate-f-basic-local-trust-posture-dev") {
		return fmt.Errorf("real Gate F basic local trust posture runner ID missing")
	}
	if strings.Contains(boundaryText, "- Registry entries rendered: 76") {
		return fmt.Errorf("stale boundary-table rendered-entry marker remains unclassified")
	}
	if !strings.Contains(boundaryText, "Historical boundary-table rendered entries: 76") {
		return fmt.Errorf("boundary table does not classify historical 76-entry marker")
	}

	for _, marker := range []string{
		"external accepted planning artifact",
		"not required to exist inside the WSL workspace",
		"CarbonStackLogDocV0.8.1.md outside WSL",
		"v0.8.x working development ledger",
		"Gate B remains blocked",
		"No WSL LogDoc/Breakpoint generation",
	} {
		if !strings.Contains(docText, marker) {
			return fmt.Errorf("Gate A docs missing required reset marker %q", marker)
		}
	}
	for key, value := range gateANonclaims() {
		if value {
			return fmt.Errorf("Gate A nonclaim unexpectedly true: %s", key)
		}
	}

	reportRoot := filepath.Join(os.TempDir(), "carbonstack-gate-a-v08x-reset-dev")
	_ = os.RemoveAll(reportRoot)
	if err := os.MkdirAll(reportRoot, 0o700); err != nil {
		return err
	}
	reportPath := filepath.Join(reportRoot, "gate-a-v08x-reset-report.json")
	report := map[string]any{
		"schema_version":                          gateAV08xResetReportSchema,
		"profile":                                 "gate-a-v08x-reset-dev",
		"created_at":                              time.Now().UTC().Format(time.RFC3339),
		"gate_a_status":                           "closed_v0_8_x_reset_workflow_doctrine_ready",
		"public_v0_8_0_release_head":              "715fd76f7700e63f1de877b730c2ce6bb37225bf",
		"active_planning_authority":               "CarbonStack_Long_Term_Roadmap_v0.9.0_EVERGREEN",
		"active_planning_authority_is_external":   true,
		"roadmap_file_required_in_wsl":            false,
		"archived_prime_continuity":               "CarbonStackLogDocV0.8.0PRIME",
		"v0_8_1_logdoc_required_after_gate_a":     true,
		"v0_7_x_ledger_archived_by_pointer":       true,
		"v0_8_x_working_ledger_template_required": true,
		"gate_b_blocked_until_contract":           true,
		"logdoc_generated_in_wsl":                 false,
		"breakpoint_generated_in_wsl":             false,
		"full_runtime_dev_promoted":               false,
		"nonclaims":                               gateANonclaims(),
		"next_action":                             "Generate CarbonStackLogDocV0.8.1 outside WSL, then begin Gate B adversarial harness contract preflight only after acceptance.",
	}
	if err := gateAWriteJSON(reportPath, report); err != nil {
		return err
	}

	fmt.Println()
	fmt.Println("gate-a-v08x-reset-dev profile result:")
	fmt.Println("  PASS: v0.8.x Gate A reset and LogDocV2 workflow doctrine closure validated")
	fmt.Println("  PASS: v0.9.0 EVERGREEN recorded as external accepted planning authority")
	fmt.Println("  PASS: v0.8.0 PRIME retained as archived v0.7.x continuity source")
	fmt.Println("  PASS: v0.8.1 LogDoc active-ledger reset required outside WSL")
	fmt.Println("  PASS: Gate B remains blocked pending adversarial harness contract")
	fmt.Println("  report:", reportPath)
	return nil
}

func gateARegistryIDPresent(registry string, id string) bool {
	rx := regexp.MustCompile("(?m)^\\s*-\\s+id:\\s*" + regexp.QuoteMeta(id) + "\\s*$")
	return rx.MatchString(registry)
}

func gateANonclaims() map[string]bool {
	return map[string]bool{
		"adversarial_harness_implemented":        false,
		"verified_identity_claimed":              false,
		"full_trust_promotion_claimed":           false,
		"secure_enrollment_claimed":              false,
		"cryptographic_identity_binding_claimed": false,
		"vault_backup_restore_claimed":           false,
		"migration_repair_implemented":           false,
		"full_runtime_dev_promoted":              false,
		"tui_started":                            false,
		"service_systemd_helper_started":         false,
		"public_ingress_started":                 false,
		"containers_started":                     false,
		"pq_hybrid_implemented":                  false,
		"android_started":                        false,
		"carbonstack_os_implemented":             false,
		"production_readiness_claimed":           false,
		"production_e2ee_claimed":                false,
		"external_pen_test_claimed":              false,
		"external_audit_claimed":                 false,
	}
}

func gateAWriteJSON(path string, value any) error {
	data, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		return err
	}
	data = append(data, '\n')
	return os.WriteFile(path, data, 0o600)
}
