package main

import "testing"

func TestGateDStateRecoverySchemas(t *testing.T) {
	if gateDStateRecoveryReportSchema != "carbonstack-gate-d-state-recovery-vault-backup-report/v0" {
		t.Fatalf("unexpected Gate D report schema: %s", gateDStateRecoveryReportSchema)
	}
	if gateDStateRecoveryModelSchema != "carbonstack-state-recovery-vault-backup-model/v0" {
		t.Fatalf("unexpected Gate D model schema: %s", gateDStateRecoveryModelSchema)
	}
}

func TestGateDSeedCases(t *testing.T) {
	if len(gateDAdversarialSeedCases) != 6 {
		t.Fatalf("unexpected Gate D seed case count: %d", len(gateDAdversarialSeedCases))
	}
	required := "ADV-STATE-RESTORE-MUST-NOT-IMPORT-TRUST-001"
	found := false
	for _, id := range gateDAdversarialSeedCases {
		if id == required {
			found = true
		}
	}
	if !found {
		t.Fatalf("missing Gate D seed case %s", required)
	}
}

func TestGateDNonclaimsRemainFalse(t *testing.T) {
	for key, value := range gateDStateRecoveryNonclaims() {
		if value {
			t.Fatalf("Gate D nonclaim %s unexpectedly true", key)
		}
	}
}
