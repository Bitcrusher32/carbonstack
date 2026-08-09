package main

import (
	"encoding/json"
	"strings"
	"testing"
	"time"
)

func TestGateG4ServiceSystemdHelperReportPreservesContract(t *testing.T) {
	report, err := buildGateG4ServiceSystemdHelperReport("/tmp/carbonstack", time.Unix(0, 0).UTC())
	if err != nil {
		t.Fatal(err)
	}
	checks := map[string]string{
		"schema":             report.SchemaVersion,
		"profile":            report.Profile,
		"mode":               report.G4Mode,
		"service_model":      report.ServiceModel,
		"user_systemd":       report.UserSystemdUnit,
		"system_systemd":     report.SystemSystemdUnit,
		"service_install":    report.ServiceInstall,
		"helper_commands":    report.HelperCommands,
		"privilege_boundary": report.PrivilegeBoundary,
		"public_bind":        report.PublicBindDefault,
		"public_ingress":     report.PublicIngress,
		"containers":         report.Containers,
		"full_runtime":       report.FullRuntimeDev,
	}
	expected := map[string]string{
		"schema":             "carbonstack-gate-g4-service-systemd-helper-contract-report/v0",
		"profile":            "gate-g4-service-systemd-helper-contract-dev",
		"mode":               "contract_only",
		"service_model":      "foreground_runner_plus_local_private_service_contract",
		"user_systemd":       "described_not_installed",
		"system_systemd":     "deferred",
		"service_install":    "deferred",
		"helper_commands":    "modeled_not_executed",
		"privilege_boundary": "no_sudo_no_systemctl_no_install",
		"public_bind":        "refuse",
		"public_ingress":     "blocked_until_G5_reconfirmed",
		"containers":         "deferred_late_v0.9.x_or_v0.10.0",
		"full_runtime":       "reserved_not_promoted",
	}
	for key, want := range expected {
		if checks[key] != want {
			t.Fatalf("%s = %q, want %q", key, checks[key], want)
		}
	}
	if !report.SafeSummaryStdout {
		t.Fatalf("safe summary stdout must be true")
	}
}

func TestGateG4ServiceSystemdHelperNonclaims(t *testing.T) {
	report, err := buildGateG4ServiceSystemdHelperReport("/tmp/carbonstack", time.Unix(0, 0).UTC())
	if err != nil {
		t.Fatal(err)
	}
	body, err := json.Marshal(report)
	if err != nil {
		t.Fatal(err)
	}
	text := string(body)
	for _, marker := range []string{"not service readiness", "not systemd readiness", "not helper install readiness", "not public-ingress safety", "not public deployment readiness", "not container readiness", "not full-runtime-dev promotion", "not zero-leakage public logging posture"} {
		if !strings.Contains(text, marker) {
			t.Fatalf("expected marker %q in report JSON: %s", marker, text)
		}
	}
	if !strings.Contains(text, "ADV-NATIVE-LOG-LEAKAGE-001") {
		t.Fatalf("carried log-leakage finding missing: %s", text)
	}
}

func TestGateG4LifecycleVerbsAreModeledNotHostActions(t *testing.T) {
	report, err := buildGateG4ServiceSystemdHelperReport("/tmp/carbonstack", time.Unix(0, 0).UTC())
	if err != nil {
		t.Fatal(err)
	}
	verbs := map[string]string{}
	for _, verb := range report.LifecycleVerbs {
		verbs[verb.Verb] = verb.Status
	}
	for _, verb := range []string{"check-config", "dry-run", "status", "start", "stop", "restart"} {
		if verbs[verb] != "modeled_not_executed" {
			t.Fatalf("verb %s status = %q", verb, verbs[verb])
		}
	}
	if !strings.Contains(verbs["remove"], "blocked") {
		t.Fatalf("remove must remain blocked, got %q", verbs["remove"])
	}
}
