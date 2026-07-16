package main

import "testing"

func TestGateFReleaseCandidateClosureSchemas(t *testing.T) {
	if gateFReleaseCandidateClosureReportSchema != "carbonstack-gate-f-release-candidate-closure-report/v0" {
		t.Fatalf("unexpected report schema: %s", gateFReleaseCandidateClosureReportSchema)
	}
	if gateFReleaseCandidateHandoffSchema != "carbonstack-v0.8.0-manual-release-handoff/v0" {
		t.Fatalf("unexpected handoff schema: %s", gateFReleaseCandidateHandoffSchema)
	}
}

func TestGateF7NonclaimsRemainFalse(t *testing.T) {
	nonclaims := gateF7Nonclaims()
	for key, value := range nonclaims {
		if value {
			t.Fatalf("nonclaim %s unexpectedly true", key)
		}
	}
}
