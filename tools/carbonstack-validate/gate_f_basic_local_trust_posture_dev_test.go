package main

import "testing"

func TestGateFBasicLocalTrustPostureSchema(t *testing.T) {
	if gateFBasicLocalTrustPostureReportSchema != "carbonstack-gate-f-basic-local-trust-posture-report/v0" {
		t.Fatalf("unexpected schema: %s", gateFBasicLocalTrustPostureReportSchema)
	}
}
