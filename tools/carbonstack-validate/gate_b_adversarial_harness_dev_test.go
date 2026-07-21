package main

import "testing"

func TestGateBAdversarialHarnessSchema(t *testing.T) {
	if gateBAdversarialHarnessReportSchema != "carbonstack-adversarial-harness-report/v0" {
		t.Fatalf("unexpected Gate B report schema: %s", gateBAdversarialHarnessReportSchema)
	}
	if gateBAdversarialCaseRegistrySchema != "carbonstack-adversarial-case-registry/v0" {
		t.Fatalf("unexpected Gate B case registry schema: %s", gateBAdversarialCaseRegistrySchema)
	}
}

func TestGateBSeedCaseID(t *testing.T) {
	if gateBSeedCaseID != "ADV-TRUST-RELAY-MEMBERSHIP-NOT-TRUST-001" {
		t.Fatalf("unexpected seed case ID: %s", gateBSeedCaseID)
	}
}

func TestGateBAdversarialHarnessNonclaimsRemainFalse(t *testing.T) {
	for key, value := range gateBAdversarialHarnessNonclaims() {
		if value {
			t.Fatalf("Gate B nonclaim %s unexpectedly true", key)
		}
	}
}
