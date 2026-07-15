package main

import "testing"

func TestGateFReleasePackageSurfaceSchema(t *testing.T) {
	if gateFReleasePackageSurfaceReportSchema != "carbonstack-gate-f-release-package-surface-report/v0" {
		t.Fatalf("unexpected schema: %s", gateFReleasePackageSurfaceReportSchema)
	}
}
