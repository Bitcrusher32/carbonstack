package main

import "testing"

func TestGateAV08xResetSchema(t *testing.T) {
	if gateAV08xResetReportSchema != "carbonstack-gate-a-v08x-reset-report/v0" {
		t.Fatalf("unexpected Gate A schema: %s", gateAV08xResetReportSchema)
	}
}

func TestGateANonclaimsRemainFalse(t *testing.T) {
	for key, value := range gateANonclaims() {
		if value {
			t.Fatalf("Gate A nonclaim %s unexpectedly true", key)
		}
	}
}
