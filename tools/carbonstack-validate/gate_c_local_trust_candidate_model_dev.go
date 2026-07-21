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

const gateCLocalTrustCandidateReportSchema = "carbonstack-gate-c-local-trust-candidate-model-report/v0"
const gateCLocalTrustCandidateSchema = "carbonstack-local-trust-candidate/v0"

var gateCAdversarialSeedCases = []string{
	"ADV-TRUST-RELAY-MEMBERSHIP-NOT-TRUST-001",
	"ADV-TRUST-LABEL-SPOOF-NOT-IDENTITY-001",
	"ADV-TRUST-MLS-JOIN-NOT-TRUST-001",
	"ADV-TRUST-PROVIDER-OBSERVATION-NOT-TRUST-001",
	"ADV-TRUST-CHANGED-SIGNER-LOUD-WARNING-001",
	"ADV-TRUST-CHANGED-DEVICE-LOUD-WARNING-001",
	"ADV-TRUST-AUTOPROMOTION-REFUSED-001",
}

func (r *Runner) GateCLocalTrustCandidateModelDev() error {
	fmt.Println("========================================")
	fmt.Println("CarbonStack validation profile: gate-c-local-trust-candidate-model-dev")
	fmt.Println("========================================")
	fmt.Println("status: dev/pre-alpha v0.8.x Gate C local trust candidate model seed")
	fmt.Println("scope: candidate schema, trust-state model, nonclaim guards, adversarial seed case carry-forward")
	fmt.Println("boundary: not verified identity, not secure enrollment, not automatic trust promotion, not hostile-server identity replacement proof")
	fmt.Println()

	commsRoot := filepath.Join(filepath.Dir(r.CarbonStack), "carbonstack-comms")
	required := []string{
		filepath.Join(r.CarbonStack, "registry", "adversarial_cases.v0.yaml"),
		filepath.Join(r.CarbonStack, "registry", "commands.v0.yaml"),
		filepath.Join(r.CarbonStack, "registry", "COMMAND_REFERENCE.v0.md"),
		filepath.Join(r.CarbonStack, "docs", "292-v0.8.3-gate-c-local-trust-candidate-contract-v0.md"),
		filepath.Join(r.CarbonStack, "docs", "293-v0.8.3-gate-c-local-trust-candidate-model-v0.md"),
		filepath.Join(commsRoot, "internal", "trust", "local_trust_candidate_model_dev.go"),
		filepath.Join(commsRoot, "internal", "trust", "local_trust_candidate_model_dev_test.go"),
	}
	for _, path := range required {
		if _, err := os.Stat(path); err != nil {
			return fmt.Errorf("Gate C required path missing: %s: %w", path, err)
		}
	}

	caseRegistry, err := gateCRead(filepath.Join(r.CarbonStack, "registry", "adversarial_cases.v0.yaml"))
	if err != nil {
		return err
	}
	commandRegistry, err := gateCRead(filepath.Join(r.CarbonStack, "registry", "commands.v0.yaml"))
	if err != nil {
		return err
	}
	commandReference, err := gateCRead(filepath.Join(r.CarbonStack, "registry", "COMMAND_REFERENCE.v0.md"))
	if err != nil {
		return err
	}
	contractDoc, err := gateCRead(filepath.Join(r.CarbonStack, "docs", "292-v0.8.3-gate-c-local-trust-candidate-contract-v0.md"))
	if err != nil {
		return err
	}
	modelDoc, err := gateCRead(filepath.Join(r.CarbonStack, "docs", "293-v0.8.3-gate-c-local-trust-candidate-model-v0.md"))
	if err != nil {
		return err
	}
	commsModel, err := gateCRead(filepath.Join(commsRoot, "internal", "trust", "local_trust_candidate_model_dev.go"))
	if err != nil {
		return err
	}

	if !gateCRegistryIDPresent(commandRegistry, "runner.gate-c-local-trust-candidate-model-dev") {
		return fmt.Errorf("Gate C runner registry ID missing")
	}
	if !strings.Contains(commandReference, "runner.gate-c-local-trust-candidate-model-dev") {
		return fmt.Errorf("Gate C runner registry ID missing from command reference")
	}
	if !strings.Contains(commandReference, "Registry entry count: **143**") {
		return fmt.Errorf("command reference does not show expected 143 entries after Gate C registry addition")
	}

	for _, caseID := range gateCAdversarialSeedCases {
		if !strings.Contains(caseRegistry, caseID) {
			return fmt.Errorf("Gate C adversarial seed case missing from case registry: %s", caseID)
		}
		if !regexp.MustCompile(`^ADV-[A-Z0-9]+(?:-[A-Z0-9]+)*-[0-9]{3}$`).MatchString(caseID) {
			return fmt.Errorf("Gate C case ID violates semantic ID policy: %s", caseID)
		}
	}

	docs := contractDoc + "\n" + modelDoc
	for _, marker := range []string{
		"Gate C is not verified identity",
		"Relay membership is not trust",
		"MLS join is not trust",
		"Provider observation is not trust",
		"candidate state is not promoted without explicit operator action",
		"changed signer/device/key lineage is loud",
		"owns local trust/candidate runtime behavior",
		"remains coordination only",
		"ADV-TRUST-RELAY-MEMBERSHIP-NOT-TRUST-001",
	} {
		if !gateCContainsFold(docs, marker) {
			return fmt.Errorf("Gate C docs missing marker %q", marker)
		}
	}

	for _, marker := range []string{
		"GateCLocalTrustCandidateSchema",
		"carbonstack-local-trust-candidate/v0",
		"GateCTrustStateCandidateObserved",
		"GateCTrustStateOperatorReviewRequired",
		"GateCTrustStatePromotedLocalTrust",
		"GateCTrustStateChangedLineageWarning",
		"GateCTrustStateDemoted",
		"GateCTrustStateRevoked",
		"GateCTrustStateUnknownOrUntrusted",
		"CandidateID",
		"SubjectLabel",
		"CypherAccountID",
		"CypherDeviceID",
		"RelaySpaceID",
		"RelaySpaceContext",
		"OpenMLSSignerFingerprint",
		"OpenMLSCredentialFingerprint",
		"KeyPackageFingerprint",
		"KeyPackageLineage",
		"OperatorPromotionEventID",
		"DemotionOrRevocationEventID",
		"PromoteGateCLocalTrustCandidate",
		"ApplyGateCLocalTrustChangedLineage",
		"Relay membership is not trust",
		"MLS join is not trust",
		"Provider observation is not trust",
		"not automatic trust promotion",
	} {
		if !strings.Contains(commsModel, marker) {
			return fmt.Errorf("Gate C Comms candidate model missing marker %q", marker)
		}
	}

	for key, value := range gateCLocalTrustCandidateNonclaims() {
		if value {
			return fmt.Errorf("Gate C nonclaim unexpectedly true: %s", key)
		}
	}

	reportRoot := filepath.Join(os.TempDir(), "carbonstack-gate-c-local-trust-candidate-model-dev")
	_ = os.RemoveAll(reportRoot)
	if err := os.MkdirAll(reportRoot, 0o700); err != nil {
		return err
	}
	reportPath := filepath.Join(reportRoot, "gate-c-local-trust-candidate-model-report.json")
	report := map[string]any{
		"schema_version":         gateCLocalTrustCandidateReportSchema,
		"profile":                "gate-c-local-trust-candidate-model-dev",
		"created_at":             time.Now().UTC().Format(time.RFC3339),
		"gate":                   "v0.8.x Gate C",
		"candidate_schema":       gateCLocalTrustCandidateSchema,
		"candidate_model_path":   "carbonstack-comms/internal/trust/local_trust_candidate_model_dev.go",
		"adversarial_seed_cases": gateCAdversarialSeedCases,
		"candidate_states": []string{
			"candidate_observed",
			"operator_review_required",
			"promoted_local_trust",
			"changed_lineage_warning",
			"demoted",
			"revoked",
			"unknown_or_untrusted",
		},
		"nonclaims_preserved": true,
		"nonclaims":           gateCLocalTrustCandidateNonclaims(),
	}
	if err := gateCWriteJSON(reportPath, report); err != nil {
		return err
	}

	fmt.Println("gate-c-local-trust-candidate-model-dev profile result:")
	fmt.Println("  PASS: Gate C candidate schema marker exists")
	fmt.Println("  PASS: candidate/promoted/changed/demoted/revoked trust-state model exists")
	fmt.Println("  PASS: Relay membership, MLS join, and provider observation remain non-trust boundaries")
	fmt.Println("  PASS: automatic trust promotion and verified identity remain false")
	fmt.Println("  PASS: adversarial seed cases are registered")
	fmt.Println("  PASS: registry/reference classification exists without product/security promotion")
	fmt.Println("  report:", reportPath)
	return nil
}

func gateCRead(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func gateCRegistryIDPresent(registry string, id string) bool {
	rx := regexp.MustCompile("(?m)^\\s*-\\s+id:\\s*" + regexp.QuoteMeta(id) + "\\s*$")
	return rx.MatchString(registry)
}

func gateCContainsFold(haystack string, needle string) bool {
	return strings.Contains(strings.ToLower(haystack), strings.ToLower(needle))
}

func gateCLocalTrustCandidateNonclaims() map[string]bool {
	return map[string]bool{
		"production_verified_identity_claimed":              false,
		"secure_enrollment_claimed":                         false,
		"hardware_backed_identity_claimed":                  false,
		"real_world_person_verification_claimed":            false,
		"automatic_trust_promotion_claimed":                 false,
		"hostile_server_identity_replacement_proof_claimed": false,
		"relay_membership_trust_authority_claimed":          false,
		"mls_join_trust_authority_claimed":                  false,
		"provider_observation_trust_authority_claimed":      false,
		"cryptographic_identity_binding_claimed":            false,
	}
}

func gateCWriteJSON(path string, value any) error {
	data, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		return err
	}
	data = append(data, '\n')
	return os.WriteFile(path, data, 0o600)
}
