package main

import (
	"encoding/json"
	"strings"
	"testing"
	"time"
)

func TestGateGOperatorLifecycleReportPreservesFrozenG2(t *testing.T) {
	report, err := buildGateGOperatorLifecycleReport("/tmp/carbonstack_umbrella", time.Unix(0, 0).UTC())
	if err != nil {
		t.Fatalf("build report: %v", err)
	}
	if report.SchemaVersion != gateGOperatorLifecycleSchema {
		t.Fatalf("schema = %q", report.SchemaVersion)
	}
	if report.Profile != "operator-lifecycle-dev" {
		t.Fatalf("profile = %q", report.Profile)
	}
	if !report.SafeSummaryStdout {
		t.Fatal("expected safe-summary stdout")
	}
	if report.PreferredAggregate != "operator-lifecycle-dev" {
		t.Fatalf("preferred aggregate = %q", report.PreferredAggregate)
	}
	if report.FullRuntimeDev != "reserved_not_promoted" {
		t.Fatalf("full-runtime-dev = %q", report.FullRuntimeDev)
	}
	if report.Containers != "deferred_late_v0.9.x_or_v0.10.0" {
		t.Fatalf("containers = %q", report.Containers)
	}
	if report.PublicIngressDefault != "disabled" {
		t.Fatalf("public ingress default = %q", report.PublicIngressDefault)
	}
	if report.ServiceSystemdHelperStatus != "conditional_g4_local_private_substrate" {
		t.Fatalf("service status = %q", report.ServiceSystemdHelperStatus)
	}
	if report.PublicIngressStatus != "conditional_g5_threat_model_prototype_substrate" {
		t.Fatalf("public ingress status = %q", report.PublicIngressStatus)
	}
}

func TestGateGOperatorLifecycleReportRejectsMissingRoot(t *testing.T) {
	if _, err := buildGateGOperatorLifecycleReport("   ", time.Unix(0, 0).UTC()); err == nil {
		t.Fatal("expected missing root error")
	}
}

func TestGateGOperatorLifecycleReportCarriesFindingsAndNonclaims(t *testing.T) {
	report, err := buildGateGOperatorLifecycleReport("/tmp/carbonstack_umbrella", time.Unix(0, 0).UTC())
	if err != nil {
		t.Fatalf("build report: %v", err)
	}
	encoded, err := json.Marshal(report)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	text := string(encoded)
	for _, marker := range []string{
		"ADV-NATIVE-LOG-LEAKAGE-001",
		"ADV-NATIVE-GENERATED-ARTIFACTS-001",
		"not public deployment readiness",
		"not public-ingress safety",
		"not container readiness",
		"not full-runtime-dev promotion",
		"not production security",
	} {
		if !strings.Contains(text, marker) {
			t.Fatalf("expected marker %q in report JSON: %s", marker, text)
		}
	}
	if report.ExitPolicy.ReleaseBlocker != "nonzero_in_all_aggregates" {
		t.Fatalf("release-blocker exit policy = %q", report.ExitPolicy.ReleaseBlocker)
	}
	if !strings.Contains(report.ExitPolicy.High, "visible_non_clean_status") {
		t.Fatalf("high exit policy = %q", report.ExitPolicy.High)
	}
}
