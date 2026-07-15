package main

import "testing"

func TestGateFCodeHealthSourceHygieneSchema(t *testing.T) {
	if gateFCodeHealthSourceHygieneReportSchema != "carbonstack-gate-f-code-health-source-hygiene-report/v0" {
		t.Fatalf("unexpected schema: %s", gateFCodeHealthSourceHygieneReportSchema)
	}
}
