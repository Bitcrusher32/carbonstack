package main

import "testing"

func TestGateERelayOnboardingCaseIDsFrozen(t *testing.T) {
	if len(gateERelayOnboardingCaseIDs) != 17 {
		t.Fatalf("expected 17 Gate E cases, got %d", len(gateERelayOnboardingCaseIDs))
	}
	required := map[string]bool{"ADV-RELAY-KEYPACKAGE-STALE-001": false, "ADV-RELAY-WELCOME-WRONG-DEVICE-001": false, "ADV-RELAY-CYPHER-MLS-MEMBERSHIP-MISMATCH-001": false, "ADV-RELAY-ACK-AFTER-FAILED-JOIN-001": false, "ADV-RELAY-LOCAL-STATE-ROLLBACK-ONBOARDING-001": false}
	for _, id := range gateERelayOnboardingCaseIDs {
		if _, ok := required[id]; ok {
			required[id] = true
		}
	}
	for id, found := range required {
		if !found {
			t.Fatalf("missing Gate E case %s", id)
		}
	}
}
func TestGateERelayOnboardingNonclaimsRemainFalse(t *testing.T) {
	for key, value := range gateERelayOnboardingNonclaims() {
		if value {
			t.Fatalf("Gate E nonclaim %s unexpectedly true", key)
		}
	}
}
func TestGateERelayOnboardingSchema(t *testing.T) {
	if gateERelayOnboardingReportSchema != "carbonstack-gate-e-relay-onboarding-adversarial-report/v0" {
		t.Fatalf("unexpected schema: %s", gateERelayOnboardingReportSchema)
	}
}
