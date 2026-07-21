package main

import "testing"

func TestGateFNativeProcessF2CaseIDs(t *testing.T) {
	if len(gateFNativeProcessF2CaseIDs) != 15 {
		t.Fatalf("expected 15 Gate F F2 cases, got %d", len(gateFNativeProcessF2CaseIDs))
	}
}

func TestGateFNativeProcessF2Schema(t *testing.T) {
	if gateFNativeProcessF2ReportSchema != "carbonstack-gate-f-native-process-f2-component-probes-report/v0" {
		t.Fatalf("unexpected schema %s", gateFNativeProcessF2ReportSchema)
	}
}
