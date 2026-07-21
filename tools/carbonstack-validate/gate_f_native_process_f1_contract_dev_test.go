package main

import "testing"

func TestGateFNativeProcessF1CaseIDs(t *testing.T) {
	if len(gateFNativeProcessF1CaseIDs) != 15 {
		t.Fatalf("expected 15 Gate F F1 cases, got %d", len(gateFNativeProcessF1CaseIDs))
	}
	required := map[string]bool{
		"ADV-NATIVE-CYPHER-CONFIG-ENV-MISUSE-001":     false,
		"ADV-NATIVE-LOG-LEAKAGE-001":                  false,
		"ADV-NATIVE-DEBIAN-HOST-EXPORT-SLOW-DISK-001": false,
		"ADV-NATIVE-CLI-TERMINATING-INSPECTION-001":   false,
	}
	for _, id := range gateFNativeProcessF1CaseIDs {
		if _, ok := required[id]; ok {
			required[id] = true
		}
	}
	for id, found := range required {
		if !found {
			t.Fatalf("missing required Gate F F1 case %s", id)
		}
	}
}

func TestGateFNativeProcessF1Schema(t *testing.T) {
	if gateFNativeProcessF1ReportSchema != "carbonstack-gate-f-native-process-f1-contract-report/v0" {
		t.Fatalf("unexpected schema %s", gateFNativeProcessF1ReportSchema)
	}
}
