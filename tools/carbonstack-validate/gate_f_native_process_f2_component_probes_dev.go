package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const gateFNativeProcessF2ReportSchema = "carbonstack-gate-f-native-process-f2-component-probes-report/v0"

var gateFNativeProcessF2CaseIDs = []string{
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

func (r *Runner) GateFNativeProcessF2ComponentProbesDev() error {
	fmt.Println("========================================")
	fmt.Println("CarbonStack validation profile: gate-f-native-process-f2-component-probes-dev")
	fmt.Println("========================================")
	fmt.Println("status: dev/pre-alpha v0.8.x Gate F F2 component probe validation")
	fmt.Println("scope: carbonstack + comms + cypher component runtime/process probe substrate")
	fmt.Println("boundary: F2 is not deep WSL process testing or real Debian laptop validation")
	fmt.Println()

	commsRoot := filepath.Join(filepath.Dir(r.CarbonStack), "carbonstack-comms")
	cypherRoot := filepath.Join(filepath.Dir(r.CarbonStack), "carbonstack-cypher")

	required := []string{
		filepath.Join(r.CarbonStack, "docs", "302-v0.8.6-gate-f-f2-component-runtime-process-probes-contract-v0.md"),
		filepath.Join(r.CarbonStack, "docs", "303-v0.8.6-gate-f-f2-component-probe-report-model-v0.md"),
		filepath.Join(r.CarbonStack, "registry", "commands.v0.yaml"),
		filepath.Join(r.CarbonStack, "registry", "COMMAND_REFERENCE.v0.md"),
		filepath.Join(commsRoot, "internal", "adversarial", "gate_f_native_process_probe_dev.go"),
		filepath.Join(commsRoot, "internal", "adversarial", "gate_f_native_process_probe_dev_test.go"),
		filepath.Join(cypherRoot, "internal", "adversarial", "gate_f_native_process_probe_dev.go"),
		filepath.Join(cypherRoot, "internal", "adversarial", "gate_f_native_process_probe_dev_test.go"),
	}
	for _, path := range required {
		if _, err := os.Stat(path); err != nil {
			return fmt.Errorf("Gate F F2 required path missing: %s: %w", path, err)
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
	commsProbe, err := os.ReadFile(filepath.Join(commsRoot, "internal", "adversarial", "gate_f_native_process_probe_dev.go"))
	if err != nil {
		return err
	}
	cypherProbe, err := os.ReadFile(filepath.Join(cypherRoot, "internal", "adversarial", "gate_f_native_process_probe_dev.go"))
	if err != nil {
		return err
	}

	if !regexp.MustCompile(`(?m)^\s*-\s+id:\s*runner.gate-f-native-process-f2-component-probes-dev\s*$`).Match(registry) {
		return fmt.Errorf("Gate F F2 registry ID missing")
	}
	if !strings.Contains(string(reference), "runner.gate-f-native-process-f2-component-probes-dev") {
		return fmt.Errorf("Gate F F2 registry ID missing from command reference")
	}
	if !strings.Contains(string(reference), "Registry entry count: **147**") {
		return fmt.Errorf("command reference does not show expected 147 entries")
	}

	for _, caseID := range gateFNativeProcessF2CaseIDs {
		if !strings.Contains(string(cases), caseID) {
			return fmt.Errorf("Gate F case missing from adversarial registry: %s", caseID)
		}
	}

	for _, marker := range []string{
		"GateFCommsNativeProcessProbeSchema",
		"GateFClassifyCommsArtifactPath",
		"GateFClassifyCommsLogOutput",
		"GateFClassifyCommsCLIHelp",
	} {
		if !strings.Contains(string(commsProbe), marker) {
			return fmt.Errorf("Comms F2 probe missing marker %q", marker)
		}
	}
	for _, marker := range []string{
		"GateFCypherNativeProcessProbeSchema",
		"GateFClassifyCypherDBPath",
		"GateFClassifyCypherPortCollision",
		"GateFRedactCypherConfig",
		"GateFClassifyCypherLogOutput",
	} {
		if !strings.Contains(string(cypherProbe), marker) {
			return fmt.Errorf("Cypher F2 probe missing marker %q", marker)
		}
	}

	report := map[string]any{
		"schema_version":                       gateFNativeProcessF2ReportSchema,
		"profile":                              "gate-f-native-process-f2-component-probes-dev",
		"case_count":                           len(gateFNativeProcessF2CaseIDs),
		"comms_probe_schema":                   "carbonstack-comms-gate-f-native-process-probe/v0",
		"cypher_probe_schema":                  "carbonstack-cypher-gate-f-native-process-probe/v0",
		"f2_component_probe_substrate_present": true,
		"deep_wsl_testing_deferred":            true,
		"real_debian_laptop_testing_deferred":  true,
		"nonclaims_preserved":                  true,
	}
	data, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(data))
	fmt.Println("VALIDATION PASSED")
	return nil
}
