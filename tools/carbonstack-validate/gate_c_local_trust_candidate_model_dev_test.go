package main

import "testing"

func TestGateCLocalTrustCandidateSchemas(t *testing.T) {
	if gateCLocalTrustCandidateReportSchema != "carbonstack-gate-c-local-trust-candidate-model-report/v0" {
		t.Fatalf("unexpected Gate C report schema: %s", gateCLocalTrustCandidateReportSchema)
	}
	if gateCLocalTrustCandidateSchema != "carbonstack-local-trust-candidate/v0" {
		t.Fatalf("unexpected Gate C candidate schema: %s", gateCLocalTrustCandidateSchema)
	}
}

func TestGateCSeedCases(t *testing.T) {
	if len(gateCAdversarialSeedCases) != 7 {
		t.Fatalf("unexpected Gate C seed case count: %d", len(gateCAdversarialSeedCases))
	}
	required := "ADV-TRUST-RELAY-MEMBERSHIP-NOT-TRUST-001"
	found := false
	for _, id := range gateCAdversarialSeedCases {
		if id == required {
			found = true
		}
	}
	if !found {
		t.Fatalf("missing carried-forward seed case %s", required)
	}
}

func TestGateCNonclaimsRemainFalse(t *testing.T) {
	for key, value := range gateCLocalTrustCandidateNonclaims() {
		if value {
			t.Fatalf("Gate C nonclaim %s unexpectedly true", key)
		}
	}
}
