package main

import "testing"

func TestGateFCompatRollbackObservabilitySchema(t *testing.T) {
	if gateFCompatRollbackObservabilityReportSchema != "carbonstack-gate-f-compat-rollback-observability-report/v0" {
		t.Fatalf("unexpected schema: %s", gateFCompatRollbackObservabilityReportSchema)
	}
}
