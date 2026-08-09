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

const gateGOperatorLifecycleSchema = "carbonstack-gate-g-operator-lifecycle-report/v0"

type gateGSubgatePlan struct {
	ID          string `json:"id"`
	Status      string `json:"status"`
	Description string `json:"description"`
}

type gateGExitPolicy struct {
	ReleaseBlocker string `json:"release_blocker"`
	High           string `json:"high"`
	Medium         string `json:"medium"`
	Informational  string `json:"informational"`
}

type gateGCarriedFinding struct {
	CaseID      string `json:"case_id"`
	Severity    string `json:"severity"`
	Status      string `json:"status"`
	Disposition string `json:"disposition"`
}

type gateGOperatorLifecycleReport struct {
	SchemaVersion              string                `json:"schema_version"`
	Profile                    string                `json:"profile"`
	Status                     string                `json:"status"`
	SafeSummaryStdout          bool                  `json:"safe_summary_stdout"`
	DetailedEvidenceRootPolicy string                `json:"detailed_evidence_root_policy"`
	PreferredAggregate         string                `json:"preferred_aggregate"`
	FullRuntimeDev             string                `json:"full_runtime_dev"`
	Containers                 string                `json:"containers"`
	PublicIngressDefault       string                `json:"public_ingress_default"`
	ServiceSystemdHelperStatus string                `json:"service_systemd_helper_status"`
	PublicIngressStatus        string                `json:"public_ingress_status"`
	Subgates                   []gateGSubgatePlan    `json:"subgates"`
	ExitPolicy                 gateGExitPolicy       `json:"exit_policy"`
	CarriedFindings            []gateGCarriedFinding `json:"carried_findings"`
	Nonclaims                  []string              `json:"nonclaims"`
	CreatedAt                  string                `json:"created_at"`
	Metadata                   map[string]string     `json:"metadata"`
}

func (r *Runner) GateGOperatorLifecycleDev() error {
	root, err := gateGFindCarbonRootFromCWD()
	if err != nil {
		return err
	}
	report, err := buildGateGOperatorLifecycleReport(root, time.Now().UTC())
	if err != nil {
		return err
	}
	if err := validateGateGOperatorLifecycleRepo(root); err != nil {
		return err
	}

	fmt.Println("========================================")
	fmt.Println("CarbonStack validation profile: operator-lifecycle-dev")
	fmt.Println("========================================")
	fmt.Println("status: Gate G3 operator lifecycle contract validated")
	fmt.Println("schema_version:", report.SchemaVersion)
	fmt.Println("preferred_aggregate:", report.PreferredAggregate)
	fmt.Println("full_runtime_dev:", report.FullRuntimeDev)
	fmt.Println("containers:", report.Containers)
	fmt.Println("public_ingress_default:", report.PublicIngressDefault)
	fmt.Println("service_systemd_helper_status:", report.ServiceSystemdHelperStatus)
	fmt.Println("public_ingress_status:", report.PublicIngressStatus)
	fmt.Println("safe_summary_stdout:", report.SafeSummaryStdout)
	fmt.Println("carried_high_finding: ADV-NATIVE-LOG-LEAKAGE-001")
	fmt.Println("next_allowed_subgate: G4 only after G3 closure; G5 only after G4 closure/reconfirmation")
	fmt.Println("boundary: operator lifecycle classification, not service readiness, public deployment readiness, public-ingress safety, container readiness, release mutation, production security, or full-runtime-dev promotion")
	fmt.Println()
	fmt.Println("VALIDATION PASSED")
	return nil
}

func buildGateGOperatorLifecycleReport(root string, createdAt time.Time) (gateGOperatorLifecycleReport, error) {
	root = strings.TrimSpace(root)
	if root == "" {
		return gateGOperatorLifecycleReport{}, errors.New("root is required")
	}
	report := gateGOperatorLifecycleReport{
		SchemaVersion:              gateGOperatorLifecycleSchema,
		Profile:                    "operator-lifecycle-dev",
		Status:                     "g3_operator_lifecycle_contract_validated",
		SafeSummaryStdout:          true,
		DetailedEvidenceRootPolicy: "outer_scripts_write_detailed_evidence_under_carbonstack_v08x_evidence",
		PreferredAggregate:         "operator-lifecycle-dev",
		FullRuntimeDev:             "reserved_not_promoted",
		Containers:                 "deferred_late_v0.9.x_or_v0.10.0",
		PublicIngressDefault:       "disabled",
		ServiceSystemdHelperStatus: "conditional_g4_local_private_substrate",
		PublicIngressStatus:        "conditional_g5_threat_model_prototype_substrate",
		Subgates: []gateGSubgatePlan{
			{ID: "G3", Status: "active_current", Description: "operator lifecycle engine mutation"},
			{ID: "G4", Status: "conditional_after_g3", Description: "local/private service/systemd/helper substrate"},
			{ID: "G5", Status: "conditional_after_g4", Description: "public-ingress threat-model/prototype substrate"},
			{ID: "H", Status: "blocked_until_gate_g_surfaces_settle", Description: "integrated native adversarial campaign"},
			{ID: "I", Status: "blocked_until_operator_deploy_shape_known", Description: "PQ/hybrid research posture"},
			{ID: "J", Status: "blocked_until_gate_a_i_evidence", Description: "v0.9.0 release-candidate closure"},
		},
		ExitPolicy: gateGExitPolicy{
			ReleaseBlocker: "nonzero_in_all_aggregates",
			High:           "visible_non_clean_status_in_dev_profiles_nonzero_in_closure_release_profiles_unless_accepted_or_deferred",
			Medium:         "classified_in_report_not_auto_fail_unless_closure_profile_says_so",
			Informational:  "report_only",
		},
		CarriedFindings: []gateGCarriedFinding{
			{CaseID: "ADV-NATIVE-LOG-LEAKAGE-001", Severity: "high", Status: "sensitive_output_review_required", Disposition: "carry_forward_log_hygiene_debt"},
			{CaseID: "ADV-NATIVE-GENERATED-ARTIFACTS-001", Severity: "medium", Status: "artifact_hygiene_debt", Disposition: "carry_forward_known_artifact_review"},
		},
		Nonclaims: []string{
			"not production security",
			"not production E2EE",
			"not verified identity",
			"not trust promotion",
			"not secure enrollment",
			"not malicious-relay safety",
			"not metadata privacy",
			"not service readiness",
			"not systemd readiness",
			"not helper install readiness",
			"not public-ingress safety",
			"not public deployment readiness",
			"not container readiness",
			"not deployment hardening certification",
			"not release/package mutation",
			"not external audit",
			"not external pen-test completion",
			"not PQ/hybrid security",
			"not Android readiness",
			"not CarbonStackOS implementation",
			"not TUI/public UX",
			"not full-runtime-dev promotion",
		},
		CreatedAt: createdAt.Format(time.RFC3339),
		Metadata: map[string]string{
			"root":        root,
			"g2_contract": "frozen_v0.8.8",
		},
	}
	return report, nil
}

func validateGateGOperatorLifecycleRepo(root string) error {
	carbonRoot := gateGCarbonRoot(root)
	required := map[string][]string{
		filepath.Join(carbonRoot, "docs", "304-v0.8.8-gate-g-operator-lifecycle-contract-v0.md"): {
			"carbonstack-gate-g-operator-lifecycle-contract/v0",
			"PREFERRED_AGGREGATE=operator-lifecycle-dev",
			"CONTAINERS=deferred_late_v0.9.x_or_v0.10.0",
			"PUBLIC_INGRESS_DEFAULT=disabled",
			"PUBLIC_INGRESS_GATE_H_REQUIREMENT=mandatory_if_G5_exists",
		},
		filepath.Join(carbonRoot, "docs", "305-v0.8.8-gate-g-operator-lifecycle-report-model-v0.md"): {
			"carbonstack-gate-g-operator-lifecycle-report/v0",
			"preferred_aggregate=operator-lifecycle-dev",
			"full_runtime_dev=reserved_not_promoted",
			"containers=deferred_late_v0.9.x_or_v0.10.0",
			"public_ingress_default=disabled",
		},
		filepath.Join(carbonRoot, "registry", "commands.v0.yaml"): {
			"runner.operator-lifecycle-dev",
			"go run . --profile operator-lifecycle-dev",
			"not public deployment readiness",
			"not full-runtime-dev promotion",
		},
	}

	for path, markers := range required {
		body, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("read required Gate G operator lifecycle artifact %s: %w", path, err)
		}
		text := string(body)
		for _, marker := range markers {
			if !strings.Contains(text, marker) {
				return fmt.Errorf("required marker %q missing from %s", marker, path)
			}
		}
	}

	report, err := buildGateGOperatorLifecycleReport(root, time.Unix(0, 0).UTC())
	if err != nil {
		return err
	}
	if !report.SafeSummaryStdout {
		return errors.New("operator lifecycle report must default to safe-summary stdout")
	}
	if report.FullRuntimeDev != "reserved_not_promoted" {
		return errors.New("full-runtime-dev must remain reserved/not promoted")
	}
	if report.PublicIngressDefault != "disabled" {
		return errors.New("public ingress must default disabled")
	}
	if report.Containers != "deferred_late_v0.9.x_or_v0.10.0" {
		return errors.New("containers must remain deferred")
	}
	_, err = json.Marshal(report)
	return err
}

func gateGCarbonRoot(root string) string {
	root = strings.TrimSpace(root)
	if root == "" {
		return root
	}
	if _, err := os.Stat(filepath.Join(root, "registry", "commands.v0.yaml")); err == nil {
		return root
	}
	return filepath.Join(root, "carbonstack")
}

func gateGFindCarbonRootFromCWD() (string, error) {
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
