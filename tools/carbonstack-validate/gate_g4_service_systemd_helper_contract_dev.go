package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const gateG4ServiceSystemdHelperSchema = "carbonstack-gate-g4-service-systemd-helper-contract-report/v0"

type gateG4RootPolicy struct {
	Name   string `json:"name"`
	Status string `json:"status"`
	Rule   string `json:"rule"`
}

type gateG4LifecycleVerb struct {
	Verb   string `json:"verb"`
	Status string `json:"status"`
}

type gateG4CarriedFinding struct {
	CaseID      string `json:"case_id"`
	Severity    string `json:"severity"`
	Status      string `json:"status"`
	Disposition string `json:"disposition"`
}

type gateG4ServiceSystemdHelperReport struct {
	SchemaVersion     string                 `json:"schema_version"`
	Profile           string                 `json:"profile"`
	Status            string                 `json:"status"`
	G4Mode            string                 `json:"g4_mode"`
	ServiceModel      string                 `json:"service_model"`
	UserSystemdUnit   string                 `json:"user_systemd_unit"`
	SystemSystemdUnit string                 `json:"system_systemd_unit"`
	ServiceInstall    string                 `json:"service_install"`
	HelperCommands    string                 `json:"helper_commands"`
	PrivilegeBoundary string                 `json:"privilege_boundary"`
	RootPolicy        []gateG4RootPolicy     `json:"root_policy"`
	LifecycleVerbs    []gateG4LifecycleVerb  `json:"lifecycle_verbs"`
	PublicBindDefault string                 `json:"public_bind_default"`
	PublicIngress     string                 `json:"public_ingress_status"`
	Containers        string                 `json:"containers"`
	FullRuntimeDev    string                 `json:"full_runtime_dev"`
	SafeSummaryStdout bool                   `json:"safe_summary_stdout"`
	CarriedFindings   []gateG4CarriedFinding `json:"carried_findings"`
	Nonclaims         []string               `json:"nonclaims"`
	CreatedAt         string                 `json:"created_at"`
	Metadata          map[string]string      `json:"metadata"`
}

func (r *Runner) GateG4ServiceSystemdHelperContractDev() error {
	root, err := gateG4FindCarbonRootFromCWD()
	if err != nil {
		return err
	}
	report, err := buildGateG4ServiceSystemdHelperReport(root, time.Now().UTC())
	if err != nil {
		return err
	}
	if err := validateGateG4ServiceSystemdHelperRepo(root); err != nil {
		return err
	}

	fmt.Println("========================================")
	fmt.Println("CarbonStack validation profile: gate-g4-service-systemd-helper-contract-dev")
	fmt.Println("========================================")
	fmt.Println("status: Gate G4a service/systemd/helper contract validated")
	fmt.Println("schema_version:", report.SchemaVersion)
	fmt.Println("g4_mode:", report.G4Mode)
	fmt.Println("service_model:", report.ServiceModel)
	fmt.Println("user_systemd_unit:", report.UserSystemdUnit)
	fmt.Println("system_systemd_unit:", report.SystemSystemdUnit)
	fmt.Println("service_install:", report.ServiceInstall)
	fmt.Println("helper_commands:", report.HelperCommands)
	fmt.Println("privilege_boundary:", report.PrivilegeBoundary)
	fmt.Println("public_bind_default:", report.PublicBindDefault)
	fmt.Println("public_ingress_status:", report.PublicIngress)
	fmt.Println("containers:", report.Containers)
	fmt.Println("full_runtime_dev:", report.FullRuntimeDev)
	fmt.Println("safe_summary_stdout:", report.SafeSummaryStdout)
	fmt.Println("carried_high_finding: ADV-NATIVE-LOG-LEAKAGE-001")
	fmt.Println("next_allowed_subgate: G4b only after G4a closure, or G5 only after G4 closure and public-ingress reconfirmation")
	fmt.Println("boundary: contract-only local/private lifecycle evidence; no service install, no systemd execution, no public bind implementation, no public-ingress safety, no container readiness, no release mutation, no production security, and no full-runtime-dev promotion")
	fmt.Println()
	fmt.Println("VALIDATION PASSED")
	return nil
}

func buildGateG4ServiceSystemdHelperReport(root string, createdAt time.Time) (gateG4ServiceSystemdHelperReport, error) {
	root = strings.TrimSpace(root)
	if root == "" {
		return gateG4ServiceSystemdHelperReport{}, errors.New("root is required")
	}
	report := gateG4ServiceSystemdHelperReport{
		SchemaVersion:     gateG4ServiceSystemdHelperSchema,
		Profile:           "gate-g4-service-systemd-helper-contract-dev",
		Status:            "g4a_contract_validated_service_helper_not_implemented",
		G4Mode:            "contract_only",
		ServiceModel:      "foreground_runner_plus_local_private_service_contract",
		UserSystemdUnit:   "described_not_installed",
		SystemSystemdUnit: "deferred",
		ServiceInstall:    "deferred",
		HelperCommands:    "modeled_not_executed",
		PrivilegeBoundary: "no_sudo_no_systemctl_no_install",
		RootPolicy: []gateG4RootPolicy{
			{Name: "CONFIG_ROOT", Status: "required_explicit", Rule: "unknown_or_ambiguous_refuse"},
			{Name: "STATE_ROOT", Status: "required_explicit", Rule: "unknown_or_ambiguous_refuse"},
			{Name: "RUNTIME_ROOT", Status: "required_explicit", Rule: "unknown_or_ambiguous_refuse"},
			{Name: "EVIDENCE_ROOT", Status: "required_explicit", Rule: "never_formal_git_repo"},
			{Name: "LOG_ROOT", Status: "required_explicit", Rule: "safe_summary_stdout_detailed_private_evidence"},
			{Name: "DB_PATH", Status: "required_explicit", Rule: "unknown_or_conflicting_refuse"},
			{Name: "MIGRATIONS_PATH", Status: "required_explicit", Rule: "unknown_or_conflicting_refuse"},
			{Name: "GENERATED_ARTIFACT_ROOT", Status: "required_explicit", Rule: "cleanup_only_profile_owned_same_run"},
		},
		LifecycleVerbs: []gateG4LifecycleVerb{
			{Verb: "check-config", Status: "modeled_not_executed"},
			{Verb: "dry-run", Status: "modeled_not_executed"},
			{Verb: "status", Status: "modeled_not_executed"},
			{Verb: "start", Status: "modeled_not_executed"},
			{Verb: "stop", Status: "modeled_not_executed"},
			{Verb: "restart", Status: "modeled_not_executed"},
			{Verb: "remove", Status: "blocked_except_future_profile_owned_cleanup_contract"},
		},
		PublicBindDefault: "refuse",
		PublicIngress:     "blocked_until_G5_reconfirmed",
		Containers:        "deferred_late_v0.9.x_or_v0.10.0",
		FullRuntimeDev:    "reserved_not_promoted",
		SafeSummaryStdout: true,
		CarriedFindings: []gateG4CarriedFinding{
			{CaseID: "ADV-NATIVE-LOG-LEAKAGE-001", Severity: "high", Status: "sensitive_output_review_required", Disposition: "carry_forward_into_G4_output_design"},
			{CaseID: "ADV-NATIVE-GENERATED-ARTIFACTS-001", Severity: "medium", Status: "artifact_hygiene_debt", Disposition: "carry_forward_known_artifact_review"},
		},
		Nonclaims: []string{
			"not service readiness", "not systemd readiness", "not helper install readiness", "not public-ingress safety", "not public deployment readiness", "not container readiness", "not deployment hardening certification", "not production security", "not production E2EE", "not verified identity", "not trust promotion", "not secure enrollment", "not malicious-relay safety", "not metadata privacy", "not external audit", "not external pen-test completion", "not PQ/hybrid security", "not Android readiness", "not CarbonStackOS implementation", "not TUI/public UX", "not full-runtime-dev promotion", "not zero-leakage public logging posture",
		},
		CreatedAt: createdAt.Format(time.RFC3339),
		Metadata:  map[string]string{"root": root, "g4_contract": "g4a_contract_only", "g3_anchor": "operator-lifecycle-dev"},
	}
	return report, nil
}

func validateGateG4ServiceSystemdHelperRepo(root string) error {
	carbonRoot := gateG4CarbonRoot(root)
	required := map[string][]string{
		filepath.Join(carbonRoot, "docs", "306-v0.8.10-gate-g-g4-service-systemd-helper-contract-v0.md"):     {"carbonstack-gate-g4-service-systemd-helper-contract/v0", "G4_MODE=contract_only", "SERVICE_MODEL=foreground_runner_plus_local_private_service_contract", "SUDO_ALLOWED=false", "SYSTEMCTL_EXECUTION_ALLOWED=false", "PUBLIC_BIND_DEFAULT=refuse", "PUBLIC_INGRESS_IMPLEMENTED=false", "CONTAINERS=deferred_late_v0.9.x_or_v0.10.0", "FULL_RUNTIME_DEV=reserved_not_promoted"},
		filepath.Join(carbonRoot, "docs", "307-v0.8.10-gate-g-g4-service-systemd-helper-report-model-v0.md"): {"carbonstack-gate-g4-service-systemd-helper-contract-report/v0", "profile=gate-g4-service-systemd-helper-contract-dev", "g4_mode=contract_only", "service_model=foreground_runner_plus_local_private_service_contract", "public_bind_default=refuse", "public_ingress_status=blocked_until_G5_reconfirmed", "full_runtime_dev=reserved_not_promoted"},
		filepath.Join(carbonRoot, "registry", "commands.v0.yaml"):                                            {"runner.gate-g4-service-systemd-helper-contract-dev", "go run . --profile gate-g4-service-systemd-helper-contract-dev", "not service readiness", "not systemd readiness", "not helper install readiness", "not public-ingress safety", "not container readiness", "not full-runtime-dev promotion"},
	}
	for path, markers := range required {
		body, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("read required Gate G4 artifact %s: %w", path, err)
		}
		text := string(body)
		for _, marker := range markers {
			if !strings.Contains(text, marker) {
				return fmt.Errorf("required marker %q missing from %s", marker, path)
			}
		}
	}
	report, err := buildGateG4ServiceSystemdHelperReport(root, time.Unix(0, 0).UTC())
	if err != nil {
		return err
	}
	if !report.SafeSummaryStdout || report.G4Mode != "contract_only" || report.ServiceInstall != "deferred" || report.PrivilegeBoundary != "no_sudo_no_systemctl_no_install" || report.PublicBindDefault != "refuse" || report.PublicIngress != "blocked_until_G5_reconfirmed" || report.Containers != "deferred_late_v0.9.x_or_v0.10.0" || report.FullRuntimeDev != "reserved_not_promoted" {
		return errors.New("G4a frozen contract values are not preserved")
	}
	for _, rel := range []string{"carbonstack.service", "carbonstack-cypher.service", "Dockerfile", "docker-compose.yml", "docker-compose.yaml", "compose.yml", "compose.yaml"} {
		if _, err := os.Stat(filepath.Join(carbonRoot, rel)); err == nil {
			return fmt.Errorf("G4a contract-only profile must not add deploy/container artifact %s", rel)
		}
	}
	_, err = json.Marshal(report)
	return err
}

func gateG4CarbonRoot(root string) string {
	root = strings.TrimSpace(root)
	if root == "" {
		return root
	}
	if _, err := os.Stat(filepath.Join(root, "registry", "commands.v0.yaml")); err == nil {
		return root
	}
	return filepath.Join(root, "carbonstack")
}

func gateG4FindCarbonRootFromCWD() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	for current := cwd; ; current = filepath.Dir(current) {
		if _, err := os.Stat(filepath.Join(current, "registry", "commands.v0.yaml")); err == nil {
			return current, nil
		}
		if filepath.Base(current) == "carbonstack_umbrella" {
			candidate := filepath.Join(current, "carbonstack")
			if _, err := os.Stat(filepath.Join(candidate, "registry", "commands.v0.yaml")); err == nil {
				return candidate, nil
			}
		}
		parent := filepath.Dir(current)
		if parent == current {
			return "", fmt.Errorf("could not locate carbonstack root from cwd %s", cwd)
		}
	}
}
