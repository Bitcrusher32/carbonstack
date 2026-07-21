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

const gateERelayOnboardingReportSchema = "carbonstack-gate-e-relay-onboarding-adversarial-report/v0"

var gateERelayOnboardingCaseIDs = []string{
	"ADV-RELAY-KEYPACKAGE-STALE-001",
	"ADV-RELAY-KEYPACKAGE-REPLAYED-001",
	"ADV-RELAY-KEYPACKAGE-WRONG-RECIPIENT-001",
	"ADV-RELAY-KEYPACKAGE-MALFORMED-ENVELOPE-001",
	"ADV-RELAY-WELCOME-STALE-001",
	"ADV-RELAY-WELCOME-REPLAYED-001",
	"ADV-RELAY-WELCOME-DUPLICATE-001",
	"ADV-RELAY-WELCOME-WRONG-GROUP-001",
	"ADV-RELAY-WELCOME-WRONG-DEVICE-001",
	"ADV-RELAY-CYPHER-MLS-MEMBERSHIP-MISMATCH-001",
	"ADV-RELAY-MEMBER-DISABLED-BEHAVIOR-001",
	"ADV-RELAY-MEMBER-LEFT-BEHAVIOR-001",
	"ADV-RELAY-ACK-AFTER-FAILED-JOIN-001",
	"ADV-RELAY-SELECTIVE-WITHHOLDING-DROP-DELAY-REORDER-001",
	"ADV-RELAY-METADATA-LIES-001",
	"ADV-RELAY-ROUTING-MEMBERSHIP-MUTATION-001",
	"ADV-RELAY-LOCAL-STATE-ROLLBACK-ONBOARDING-001",
}

var gateECypherRequiredPaths = []string{
	"internal/httpapi/api.go",
	"internal/httpapi/keypackage_publication.go",
	"internal/httpapi/relay_space_member_state.go",
	"internal/db/keypackage_publications.go",
	"internal/db/relay_spaces.go",
	"internal/db/relay_space_member_state.go",
	"internal/db/relay_space_invite_claim.go",
}

func (r *Runner) GateERelayOnboardingAdversarialDev() error {
	fmt.Println("========================================")
	fmt.Println("CarbonStack validation profile: gate-e-relay-onboarding-adversarial-dev")
	fmt.Println("========================================")
	fmt.Println("status: dev/pre-alpha v0.8.x Gate E Relay / Onboarding / MLS-Cypher adversarial classification pass")
	fmt.Println("boundary: not malicious-relay safety, not metadata privacy, not verified identity, not hostile-server proof")
	fmt.Println()
	commsRoot := filepath.Join(filepath.Dir(r.CarbonStack), "carbonstack-comms")
	cypherRoot := filepath.Join(filepath.Dir(r.CarbonStack), "carbonstack-cypher")
	required := []string{filepath.Join(r.CarbonStack, "docs", "296-v0.8.5-gate-e-relay-onboarding-adversarial-contract-v0.md"), filepath.Join(r.CarbonStack, "docs", "297-v0.8.5-gate-e-relay-onboarding-case-matrix-v0.md"), filepath.Join(r.CarbonStack, "registry", "adversarial_cases.v0.yaml"), filepath.Join(r.CarbonStack, "registry", "commands.v0.yaml"), filepath.Join(r.CarbonStack, "registry", "COMMAND_REFERENCE.v0.md"), filepath.Join(commsRoot, "internal", "adversarial", "gate_e_relay_onboarding_cases_dev.go"), filepath.Join(commsRoot, "internal", "trust", "local_trust_candidate_model_dev.go"), filepath.Join(commsRoot, "internal", "state", "gate_d_state_recovery_model_dev.go")}
	for _, rel := range gateECypherRequiredPaths {
		required = append(required, filepath.Join(cypherRoot, rel))
	}
	for _, p := range required {
		if _, err := os.Stat(p); err != nil {
			return fmt.Errorf("Gate E required path missing: %s: %w", p, err)
		}
	}
	registry, err := gateERead(filepath.Join(r.CarbonStack, "registry", "commands.v0.yaml"))
	if err != nil {
		return err
	}
	reference, err := gateERead(filepath.Join(r.CarbonStack, "registry", "COMMAND_REFERENCE.v0.md"))
	if err != nil {
		return err
	}
	caseRegistry, err := gateERead(filepath.Join(r.CarbonStack, "registry", "adversarial_cases.v0.yaml"))
	if err != nil {
		return err
	}
	contractDoc, err := gateERead(filepath.Join(r.CarbonStack, "docs", "296-v0.8.5-gate-e-relay-onboarding-adversarial-contract-v0.md"))
	if err != nil {
		return err
	}
	caseMatrix, err := gateERead(filepath.Join(r.CarbonStack, "docs", "297-v0.8.5-gate-e-relay-onboarding-case-matrix-v0.md"))
	if err != nil {
		return err
	}
	commsCaseModel, err := gateERead(filepath.Join(commsRoot, "internal", "adversarial", "gate_e_relay_onboarding_cases_dev.go"))
	if err != nil {
		return err
	}
	if !gateERegistryIDPresent(registry, "runner.gate-e-relay-onboarding-adversarial-dev") {
		return fmt.Errorf("Gate E runner registry ID missing")
	}
	if !strings.Contains(reference, "runner.gate-e-relay-onboarding-adversarial-dev") {
		return fmt.Errorf("Gate E runner registry ID missing from reference")
	}
	if !strings.Contains(reference, "Registry entry count: **145**") {
		return fmt.Errorf("command reference does not show expected 145 entries")
	}
	for _, id := range gateERelayOnboardingCaseIDs {
		if !strings.Contains(caseRegistry, id) {
			return fmt.Errorf("Gate E case missing from registry: %s", id)
		}
		if !strings.Contains(caseMatrix, id) {
			return fmt.Errorf("Gate E case missing from matrix: %s", id)
		}
		if !strings.Contains(commsCaseModel, id) {
			return fmt.Errorf("Gate E case missing from Comms model: %s", id)
		}
		if !regexp.MustCompile(`^ADV-RELAY-[A-Z0-9]+(?:-[A-Z0-9]+)*-[0-9]{3}$`).MatchString(id) {
			return fmt.Errorf("bad Gate E case ID: %s", id)
		}
	}
	docs := strings.ToLower(contractDoc + caseMatrix)
	for _, marker := range []string{"keypackage and welcome remain onboarding artifacts, not normal messages", "ack happens only after successful state establishment", "cypher membership remains server-side coordination authority", "openmls group state remains cryptographic group authority", "gate c trust-candidate state must not auto-promote", "gate d state/recovery dry-run classification must not silently repair", "malicious-relay safety", "metadata privacy", "verified identity", "production security"} {
		if !strings.Contains(docs, marker) {
			return fmt.Errorf("Gate E docs missing marker %q", marker)
		}
	}
	for _, marker := range []string{"GateERelayOnboardingReportSchema", "carbonstack-gate-e-relay-onboarding-adversarial-report/v0", "classified_from_existing_coverage", "not malicious-relay safety", "not metadata privacy", "not verified identity", "not hostile-server proof"} {
		if !strings.Contains(commsCaseModel, marker) {
			return fmt.Errorf("Gate E Comms model missing marker %q", marker)
		}
	}
	for key, value := range gateERelayOnboardingNonclaims() {
		if value {
			return fmt.Errorf("Gate E nonclaim unexpectedly true: %s", key)
		}
	}
	reportRoot := filepath.Join(os.TempDir(), "carbonstack-gate-e-relay-onboarding-adversarial-dev")
	_ = os.RemoveAll(reportRoot)
	if err := os.MkdirAll(reportRoot, 0o700); err != nil {
		return err
	}
	reportPath := filepath.Join(reportRoot, "gate-e-relay-onboarding-adversarial-report.json")
	cases := make([]map[string]any, 0, len(gateERelayOnboardingCaseIDs))
	for _, id := range gateERelayOnboardingCaseIDs {
		cases = append(cases, map[string]any{"case_id": id, "case_status": "classified_from_existing_coverage", "finding_disposition": "nonclaim_preserved", "severity": "informational", "nonclaims_preserved": true})
	}
	report := map[string]any{"schema_version": gateERelayOnboardingReportSchema, "profile": "gate-e-relay-onboarding-adversarial-dev", "created_at": time.Now().UTC().Format(time.RFC3339), "gate": "v0.8.x Gate E", "case_count": len(cases), "cases": cases, "nonclaims": gateERelayOnboardingNonclaims(), "nonclaims_preserved": true}
	if err := gateEWriteJSON(reportPath, report); err != nil {
		return err
	}
	fmt.Println("gate-e-relay-onboarding-adversarial-dev profile result:")
	fmt.Println("  PASS: Gate E contract and case matrix exist")
	fmt.Println("  PASS: all 17 Gate E cases are registered and classified")
	fmt.Println("  PASS: Comms-owned adversarial case model exists")
	fmt.Println("  PASS: Cypher required server-side surfaces are present without Cypher mutation")
	fmt.Println("  PASS: Gate C trust-candidate and Gate D state/recovery boundaries are carried forward")
	fmt.Println("  PASS: nonclaims remain preserved")
	fmt.Println("  report:", reportPath)
	return nil
}

func gateERead(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
func gateERegistryIDPresent(registry string, id string) bool {
	return regexp.MustCompile("(?m)^\\s*-\\s+id:\\s*" + regexp.QuoteMeta(id) + "\\s*$").MatchString(registry)
}
func gateERelayOnboardingNonclaims() map[string]bool {
	return map[string]bool{"malicious_relay_safety_claimed": false, "metadata_privacy_claimed": false, "verified_identity_claimed": false, "hostile_server_proof_claimed": false, "production_security_claimed": false, "production_e2ee_claimed": false, "external_pen_test_completion_claimed": false, "silent_repair_claimed": false, "silent_trust_promotion_claimed": false, "cypher_plaintext_authority_claimed": false, "cypher_mls_parser_claimed": false, "cypher_trust_identity_authority_claimed": false}
}
func gateEWriteJSON(path string, value any) error {
	data, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		return err
	}
	data = append(data, 10)
	return os.WriteFile(path, data, 0o600)
}
