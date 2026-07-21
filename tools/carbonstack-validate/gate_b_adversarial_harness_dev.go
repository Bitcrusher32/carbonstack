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

const gateBAdversarialHarnessReportSchema = "carbonstack-adversarial-harness-report/v0"
const gateBAdversarialCaseRegistrySchema = "carbonstack-adversarial-case-registry/v0"
const gateBSeedCaseID = "ADV-TRUST-RELAY-MEMBERSHIP-NOT-TRUST-001"

func (r *Runner) GateBAdversarialHarnessDev() error {
	fmt.Println("========================================")
	fmt.Println("CarbonStack validation profile: gate-b-adversarial-harness-dev")
	fmt.Println("========================================")
	fmt.Println("status: dev/pre-alpha v0.8.x Gate B adversarial harness contract seed")
	fmt.Println("scope: case registry, report schema, severity/disposition model, and one low-risk historical trust-boundary seed case")
	fmt.Println("boundary: not broad adversarial coverage, not hostile-server safety, not external pen-test, not production security")
	fmt.Println()

	required := []string{
		filepath.Join(r.CarbonStack, "registry", "adversarial_cases.v0.yaml"),
		filepath.Join(r.CarbonStack, "registry", "commands.v0.yaml"),
		filepath.Join(r.CarbonStack, "registry", "COMMAND_REFERENCE.v0.md"),
		filepath.Join(r.CarbonStack, "docs", "290-v0.8.2-gate-b-adversarial-harness-contract-v0.md"),
		filepath.Join(r.CarbonStack, "docs", "291-v0.8.2-gate-b-seed-case-registry-v0.md"),
		filepath.Join(r.CarbonStack, "tools", "carbonstack-validate", "gate_b_adversarial_harness_dev.go"),
		filepath.Join(r.CarbonStack, "tools", "carbonstack-validate", "gate_b_adversarial_harness_dev_test.go"),
	}
	for _, path := range required {
		if _, err := os.Stat(path); err != nil {
			return fmt.Errorf("Gate B required path missing: %s: %w", path, err)
		}
	}

	caseRegistryBytes, err := os.ReadFile(filepath.Join(r.CarbonStack, "registry", "adversarial_cases.v0.yaml"))
	if err != nil {
		return err
	}
	commandRegistryBytes, err := os.ReadFile(filepath.Join(r.CarbonStack, "registry", "commands.v0.yaml"))
	if err != nil {
		return err
	}
	commandReferenceBytes, err := os.ReadFile(filepath.Join(r.CarbonStack, "registry", "COMMAND_REFERENCE.v0.md"))
	if err != nil {
		return err
	}
	contractDocBytes, err := os.ReadFile(filepath.Join(r.CarbonStack, "docs", "290-v0.8.2-gate-b-adversarial-harness-contract-v0.md"))
	if err != nil {
		return err
	}
	seedDocBytes, err := os.ReadFile(filepath.Join(r.CarbonStack, "docs", "291-v0.8.2-gate-b-seed-case-registry-v0.md"))
	if err != nil {
		return err
	}

	caseRegistry := string(caseRegistryBytes)
	commandRegistry := string(commandRegistryBytes)
	commandReference := string(commandReferenceBytes)
	docs := string(contractDocBytes) + "\n" + string(seedDocBytes)

	if !strings.Contains(caseRegistry, "schema_version: "+gateBAdversarialCaseRegistrySchema) {
		return fmt.Errorf("adversarial case registry schema marker missing")
	}
	if !gateBRegistryIDPresent(commandRegistry, "runner.gate-b-adversarial-harness-dev") {
		return fmt.Errorf("Gate B runner registry ID missing from commands.v0.yaml")
	}
	if !strings.Contains(commandReference, "runner.gate-b-adversarial-harness-dev") {
		return fmt.Errorf("Gate B runner registry ID missing from command reference")
	}
	if !strings.Contains(commandReference, "Registry entry count: **142**") {
		return fmt.Errorf("command reference does not show expected 142 entries after Gate B registry addition")
	}

	requiredCaseMarkers := []string{
		"case_id: " + gateBSeedCaseID,
		"surface: local-trust-boundary",
		"case_family: trust-boundary-nonpromotion",
		"severity: high",
		"finding_disposition: seed_contract_only_not_executed",
		"attack_or_fault: Treat Relay membership as local trust or verified identity.",
		"Relay membership is not trust.",
		"MLS join is not trust.",
		"Provider observation is not trust.",
		"No verified identity claim.",
		"No automatic trust promotion.",
		"regression_profile: runner.gate-b-adversarial-harness-dev",
	}
	for _, marker := range requiredCaseMarkers {
		if !strings.Contains(caseRegistry, marker) {
			return fmt.Errorf("adversarial case registry missing marker %q", marker)
		}
	}

	requiredSeverityLabels := []string{
		"  - release-blocker",
		"  - high",
		"  - medium",
		"  - informational",
		"  - accepted-risk",
	}
	for _, marker := range requiredSeverityLabels {
		if !strings.Contains(caseRegistry, marker) {
			return fmt.Errorf("severity model missing marker %q", marker)
		}
	}

	requiredDispositionLabels := []string{
		"  - passed",
		"  - failed",
		"  - not_executed",
		"  - seed_contract_only_not_executed",
		"  - fixed",
		"  - deferred",
		"  - nonclaim_preserved",
		"  - accepted_risk",
		"  - not_applicable",
	}
	for _, marker := range requiredDispositionLabels {
		if !strings.Contains(caseRegistry, marker) {
			return fmt.Errorf("finding disposition model missing marker %q", marker)
		}
	}

	requiredDocMarkers := []string{
		"Gate B is not comprehensive adversarial coverage",
		"carbonstack-only",
		"stable semantic IDs",
		"finding disposition",
		"ADV-TRUST-RELAY-MEMBERSHIP-NOT-TRUST-001",
		"Relay membership must not become local trust",
		"Do not run broad integration ladders for the narrow Gate B framework seed",
		"use broad ladders when testing full CarbonStack integration or closure",
	}
	for _, marker := range requiredDocMarkers {
		if !strings.Contains(docs, marker) {
			return fmt.Errorf("Gate B docs missing marker %q", marker)
		}
	}

	nonclaims := gateBAdversarialHarnessNonclaims()
	for key, value := range nonclaims {
		if value {
			return fmt.Errorf("Gate B nonclaim unexpectedly true: %s", key)
		}
	}

	caseIDValid := regexp.MustCompile(`^ADV-[A-Z0-9]+(?:-[A-Z0-9]+)*-[0-9]{3}$`).MatchString(gateBSeedCaseID)
	if !caseIDValid {
		return fmt.Errorf("seed case ID does not match semantic ID policy")
	}

	reportRoot := filepath.Join(os.TempDir(), "carbonstack-gate-b-adversarial-harness-dev")
	_ = os.RemoveAll(reportRoot)
	if err := os.MkdirAll(reportRoot, 0o700); err != nil {
		return err
	}
	reportPath := filepath.Join(reportRoot, "gate-b-adversarial-harness-report.json")
	report := map[string]any{
		"schema_version":         gateBAdversarialHarnessReportSchema,
		"profile":                "gate-b-adversarial-harness-dev",
		"created_at":             time.Now().UTC().Format(time.RFC3339),
		"gate":                   "v0.8.x Gate B",
		"case_registry_schema":   gateBAdversarialCaseRegistrySchema,
		"case_registry_version":  "v0.8.2-gate-b-seed",
		"case_registry_path":     "carbonstack/registry/adversarial_cases.v0.yaml",
		"cases_total":            1,
		"cases_executed":         0,
		"cases_passed":           0,
		"cases_failed":           0,
		"release_blockers":       0,
		"high_findings":          0,
		"medium_findings":        0,
		"informational_findings": 0,
		"accepted_risks":         0,
		"seed_case_id":           gateBSeedCaseID,
		"seed_case_disposition":  "seed_contract_only_not_executed",
		"nonclaims_preserved":    true,
		"nonclaims":              nonclaims,
		"broad_ladder_policy":    "focused only for this narrow framework seed; broad ladders allowed for full CarbonStack integration or closure",
		"next_action":            "Review Gate B seed harness evidence, then update the active LogDoc outside WSL.",
	}
	if err := gateBAdversarialHarnessWriteJSON(reportPath, report); err != nil {
		return err
	}

	fmt.Println("gate-b-adversarial-harness-dev profile result:")
	fmt.Println("  PASS: adversarial case registry exists")
	fmt.Println("  PASS: report schema contract is represented")
	fmt.Println("  PASS: severity labels and separate finding dispositions exist")
	fmt.Println("  PASS: seed case ADV-TRUST-RELAY-MEMBERSHIP-NOT-TRUST-001 is present and nonclaim-bounded")
	fmt.Println("  PASS: registry/reference classification exists without broad adversarial coverage claim")
	fmt.Println("  report:", reportPath)
	return nil
}

func gateBRegistryIDPresent(registry string, id string) bool {
	rx := regexp.MustCompile("(?m)^\\s*-\\s+id:\\s*" + regexp.QuoteMeta(id) + "\\s*$")
	return rx.MatchString(registry)
}

func gateBAdversarialHarnessNonclaims() map[string]bool {
	return map[string]bool{
		"comprehensive_adversarial_coverage_claimed": false,
		"hostile_server_safety_claimed":              false,
		"malicious_relay_safety_claimed":             false,
		"verified_identity_claimed":                  false,
		"full_trust_promotion_claimed":               false,
		"secure_enrollment_claimed":                  false,
		"cryptographic_identity_binding_claimed":     false,
		"production_security_claimed":                false,
		"production_e2ee_claimed":                    false,
		"external_pen_test_completion_claimed":       false,
		"external_audit_claimed":                     false,
		"public_deployment_readiness_claimed":        false,
		"container_readiness_claimed":                false,
		"tui_readiness_claimed":                      false,
		"android_or_os_readiness_claimed":            false,
		"pq_hybrid_security_claimed":                 false,
	}
}

func gateBAdversarialHarnessWriteJSON(path string, value any) error {
	data, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		return err
	}
	data = append(data, '\n')
	return os.WriteFile(path, data, 0o600)
}
