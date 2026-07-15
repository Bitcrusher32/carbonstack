package main

import "testing"

func TestGateFOperatorDocsRunbookSchema(t *testing.T) {
	if gateFOperatorDocsRunbookReportSchema != "carbonstack-gate-f-operator-docs-runbook-report/v0" {
		t.Fatalf("unexpected schema: %s", gateFOperatorDocsRunbookReportSchema)
	}
}
