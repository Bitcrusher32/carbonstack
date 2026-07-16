package main

import "testing"

func TestGateFPackageRuntimeCandidateSchemas(t *testing.T) {
	if gateFPackageRuntimeCandidateReportSchema != "carbonstack-gate-f-package-runtime-candidate-report/v0" {
		t.Fatalf("unexpected report schema: %s", gateFPackageRuntimeCandidateReportSchema)
	}
	if gateFPackageRuntimeCandidateManifestSchema != "carbonstack-package-runtime-candidate-manifest/v0" {
		t.Fatalf("unexpected manifest schema: %s", gateFPackageRuntimeCandidateManifestSchema)
	}
}
