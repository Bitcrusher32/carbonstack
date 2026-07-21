package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const gateFNativeProcessF1ReportSchema = "carbonstack-gate-f-native-process-f1-contract-report/v0"

var gateFNativeProcessF1CaseIDs = []string{
	"ADV-NATIVE-CYPHER-CONFIG-ENV-MISUSE-001",
	"ADV-NATIVE-DB-PATH-CONFUSION-001",
	"ADV-NATIVE-MIGRATION-PATH-CONFUSION-001",
	"ADV-NATIVE-STATE-ROOT-CONFUSION-001",
	"ADV-NATIVE-PACKAGE-VS-RUNTIME-ROOT-CONFUSION-001",
	"ADV-NATIVE-RESTART-SHUTDOWN-BEHAVIOR-001",
	"ADV-NATIVE-STALE-PROCESS-001",
	"ADV-NATIVE-LOG-LEAKAGE-001",
	"ADV-NATIVE-FILE-PERMISSIONS-001",
	"ADV-NATIVE-GENERATED-ARTIFACTS-001",
	"ADV-NATIVE-PORT-CONFIG-COLLISION-001",
	"ADV-NATIVE-LOCAL-CLEANUP-HAZARDS-001",
	"ADV-NATIVE-FOREGROUND-PROCESS-FAILURE-PARTIAL-STATE-001",
	"ADV-NATIVE-DEBIAN-HOST-EXPORT-SLOW-DISK-001",
	"ADV-NATIVE-CLI-TERMINATING-INSPECTION-001",
}

var gateFNativeProcessF1RequiredPaths = []string{
	"docs/298-v0.8.6-gate-f-native-deployment-process-contract-v0.md",
	"docs/299-v0.8.6-gate-f-service-systemd-helper-threat-model-v0.md",
	"docs/300-v0.8.6-gate-f-native-process-case-matrix-v0.md",
	"docs/301-v0.8.6-gate-f-subgate-plan-v0.md",
	"docs/README.md",
	"registry/adversarial_cases.v0.yaml",
	"registry/commands.v0.yaml",
	"registry/COMMAND_REFERENCE.v0.md",
	"tools/carbonstack-validate/gate_f_native_process_f1_contract_dev.go",
	"tools/carbonstack-validate/gate_f_native_process_f1_contract_dev_test.go",
}

func (r *Runner) GateFNativeProcessF1ContractDev() error {
	fmt.Println("========================================")
	fmt.Println("CarbonStack validation profile: gate-f-native-process-f1-contract-dev")
	fmt.Println("========================================")
	fmt.Println("status: dev/pre-alpha v0.8.x Gate F F1 contract/model validation")
	fmt.Println("scope: native deployment/process adversarial contract, subgate split, case matrix, service/systemd/helper threat model")
	fmt.Println("boundary: F1 is carbonstack-only; F2 owns component probes; F3/F4 own full validation and real Debian host validation")
	fmt.Println()

	for _, rel := range gateFNativeProcessF1RequiredPaths {
		if _, err := os.Stat(filepath.Join(r.CarbonStack, rel)); err != nil {
			return fmt.Errorf("Gate F F1 required path missing: %s: %w", rel, err)
		}
	}

	registry, err := os.ReadFile(filepath.Join(r.CarbonStack, "registry", "commands.v0.yaml"))
	if err != nil {
		return err
	}
	reference, err := os.ReadFile(filepath.Join(r.CarbonStack, "registry", "COMMAND_REFERENCE.v0.md"))
	if err != nil {
		return err
	}
	cases, err := os.ReadFile(filepath.Join(r.CarbonStack, "registry", "adversarial_cases.v0.yaml"))
	if err != nil {
		return err
	}

	if !regexp.MustCompile(`(?m)^\s*-\s+id:\s*runner.gate-f-native-process-f1-contract-dev\s*$`).Match(registry) {
		return fmt.Errorf("Gate F F1 registry ID missing")
	}
	if !strings.Contains(string(reference), "runner.gate-f-native-process-f1-contract-dev") {
		return fmt.Errorf("Gate F F1 registry ID missing from command reference")
	}
	if !strings.Contains(string(reference), "Registry entry count: **146**") {
		return fmt.Errorf("command reference does not show expected 146 entries")
	}

	for _, caseID := range gateFNativeProcessF1CaseIDs {
		if !strings.Contains(string(cases), caseID) {
			return fmt.Errorf("Gate F case missing from adversarial registry: %s", caseID)
		}
		if !regexp.MustCompile(`^ADV-NATIVE-[A-Z0-9]+(?:-[A-Z0-9]+)*-[0-9]{3}$`).MatchString(caseID) {
			return fmt.Errorf("Gate F case ID violates policy: %s", caseID)
		}
	}

	docBlob := ""
	for _, rel := range []string{
		"docs/298-v0.8.6-gate-f-native-deployment-process-contract-v0.md",
		"docs/299-v0.8.6-gate-f-service-systemd-helper-threat-model-v0.md",
		"docs/300-v0.8.6-gate-f-native-process-case-matrix-v0.md",
		"docs/301-v0.8.6-gate-f-subgate-plan-v0.md",
	} {
		data, err := os.ReadFile(filepath.Join(r.CarbonStack, rel))
		if err != nil {
			return err
		}
		docBlob += strings.ToLower(string(data))
	}

	for _, marker := range []string{
		"f1 is carbonstack-only",
		"f2",
		"component runtime/process probes",
		"f3/f4",
		"real debian host",
		"disposable run root",
		"service/systemd/helper threat model",
		"not service readiness",
		"not public ingress",
		"not container readiness",
		"full local validation",
	} {
		if !strings.Contains(docBlob, marker) {
			return fmt.Errorf("Gate F F1 docs missing marker %q", marker)
		}
	}

	report := map[string]any{
		"schema_version": gateFNativeProcessF1ReportSchema,
		"profile":        "gate-f-native-process-f1-contract-dev",
		"case_count":     len(gateFNativeProcessF1CaseIDs),
		"subgates": []string{
			"F1 carbonstack-only contract/model",
			"F2 carbonstack+comms+cypher component probes as justified",
			"F3/F4 full local and real Debian host validation",
		},
		"nonclaims_preserved": true,
	}
	data, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(data))
	fmt.Println("VALIDATION PASSED")
	return nil
}
