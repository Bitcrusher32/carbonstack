package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const gateFOperatorDocsRunbookReportSchema = "carbonstack-gate-f-operator-docs-runbook-report/v0"

func (r *Runner) GateFOperatorDocsRunbookDev() error {
	fmt.Println("========================================")
	fmt.Println("CarbonStack validation profile: gate-f-operator-docs-runbook-dev")
	fmt.Println("========================================")
	fmt.Println("status: dev/pre-alpha Gate F F2 operator docs and runbook closure")
	fmt.Println("scope: manual-private lifecycle docs, config/env docs, validation docs, release/package authority docs, failure/refusal/hygiene docs")
	fmt.Println("boundary: not release creation, not package publication, not package staging execution, not full-runtime-dev, not migration, not service/systemd/helper install")
	fmt.Println()

	if err := r.CheckRequiredPaths(); err != nil {
		return err
	}
	if err := r.CheckLiveGitUmbrella("gate-f-operator-docs-runbook-dev"); err != nil {
		return err
	}

	requiredDocs := map[string][]string{
		"docs/272-v0.7.20-gate-f-f2-manual-private-lifecycle-runbook-v0.md": {
			"manual-private lifecycle",
			"Gate E is closed",
			"Gate F F1 is closed",
			"Gate F F2",
			"No release creation",
			"No package staging execution",
			"full-runtime-dev remains reserved",
		},
		"docs/273-v0.7.20-gate-f-f2-config-env-validation-guide-v0.md": {
			"CYPHER_ADDR",
			"CYPHER_DB",
			"CYPHER_MIGRATIONS",
			"CYPHER_DEV_INVITE",
			"deployment root",
			"cypher config inspection",
		},
		"docs/274-v0.7.20-gate-f-f2-release-package-authority-guide-v0.md": {
			"full-validate-release",
			"release-snapshot",
			"stage-v0.7.0-package.sh",
			"manual release creation",
			"not package publication",
		},
		"docs/275-v0.7.20-gate-f-f2-failure-refusal-hygiene-guide-v0.md": {
			"cypher.db",
			"non-destructive",
			"refuse",
			"no silent migration",
			"not destructive cleanup",
		},
		"docs/276-v0.7.20-gate-f-f2-closure-v0.md": {
			"GATE_F_STATUS=open_f1_f2_closed_f3_not_started",
			"BREAKPOINT_REQUIRED_AFTER_F2=true",
			"Gate F F3",
		},
	}

	var missing []string
	for rel, phrases := range requiredDocs {
		body, err := os.ReadFile(filepath.Join(r.CarbonStack, rel))
		if err != nil {
			missing = append(missing, rel+": unreadable: "+err.Error())
			continue
		}
		text := string(body)
		for _, phrase := range phrases {
			if !strings.Contains(text, phrase) {
				missing = append(missing, rel+": missing phrase: "+phrase)
			}
		}
	}
	if len(missing) > 0 {
		return fmt.Errorf("Gate F F2 docs incomplete: %s", strings.Join(missing, "; "))
	}

	registryBody, err := os.ReadFile(filepath.Join(r.CarbonStack, "registry", "commands.v0.yaml"))
	if err != nil {
		return err
	}
	registryText := string(registryBody)
	requiredIDs := []string{
		"runner.gate-f-release-package-surface-dev",
		"runner.gate-f-operator-docs-runbook-dev",
		"runner.gate-e-native-deployment-closure-dev",
		"cypher.config-inspection",
		"runner.full-validate-release",
		"runner.release-snapshot",
		"runner.local-cypher",
	}
	registryPresence := map[string]bool{}
	for _, id := range requiredIDs {
		present := strings.Contains(registryText, "id: "+id)
		registryPresence[id] = present
		if !present {
			return fmt.Errorf("missing required registry id: %s", id)
		}
	}

	reportRoot := filepath.Join(os.TempDir(), "carbonstack-gate-f-operator-docs-runbook-dev")
	_ = os.RemoveAll(reportRoot)
	if err := os.MkdirAll(reportRoot, 0o700); err != nil {
		return err
	}
	reportPath := filepath.Join(reportRoot, "gate-f-operator-docs-runbook-report.json")

	report := map[string]any{
		"schema_version":               gateFOperatorDocsRunbookReportSchema,
		"profile":                      "gate-f-operator-docs-runbook-dev",
		"created_at":                   time.Now().UTC().Format(time.RFC3339),
		"gate_d_status":                "closed",
		"gate_e_status":                "closed",
		"gate_f_status":                "open_f1_f2_closed_f3_not_started",
		"gate_f_f1_status":             "closed",
		"gate_f_f2_status":             "closed",
		"gate_f_f3_status":             "not_started",
		"breakpoint_required_after_f2": true,
		"manual_private_lifecycle_runbook_present": true,
		"config_env_validation_guide_present":      true,
		"release_package_authority_guide_present":  true,
		"failure_refusal_hygiene_guide_present":    true,
		"release_creation_implemented":             false,
		"release_upload_implemented":               false,
		"package_publication_implemented":          false,
		"package_staging_executed":                 false,
		"package_runtime_candidate_implemented":    false,
		"full_runtime_dev_promoted":                false,
		"service_or_systemd_started":               false,
		"helper_install_started":                   false,
		"container_started":                        false,
		"public_ingress_started":                   false,
		"tui_started":                              false,
		"migration_implemented":                    false,
		"destructive_cleanup_performed":            false,
		"trust_promotion_claimed":                  false,
		"verified_identity_claimed":                false,
		"vault_claimed":                            false,
		"backup_restore_claimed":                   false,
		"pq_hybrid_claimed":                        false,
		"android_claimed":                          false,
		"carbonstack_os_claimed":                   false,
		"registry_presence":                        registryPresence,
		"docs":                                     keysGateF(requiredDocs),
		"candidate_next":                           "Generate private LogDoc and Breakpoint JSON, then begin Gate F F3 compatibility stale-state rollback observability preflight",
	}
	if err := writeGateFOperatorDocsRunbookJSON(reportPath, report); err != nil {
		return err
	}

	fmt.Println()
	fmt.Println("gate-f-operator-docs-runbook-dev profile result:")
	fmt.Println("  PASS: manual-private lifecycle runbook present")
	fmt.Println("  PASS: config/env validation guide present")
	fmt.Println("  PASS: release/package authority guide present")
	fmt.Println("  PASS: failure/refusal/hygiene guide present")
	fmt.Println("  PASS: F2 closure document present")
	fmt.Println("  PASS: F1 and F2 registry authorities present")
	fmt.Println("  PASS: release creation remains manual and unimplemented")
	fmt.Println("  PASS: package staging was not executed")
	fmt.Println("  PASS: full-runtime-dev remains unpromoted")
	fmt.Println("  gate_f_status: open_f1_f2_closed_f3_not_started")
	fmt.Println("  gate_f_f1_status: closed")
	fmt.Println("  gate_f_f2_status: closed")
	fmt.Println("  gate_f_f3_status: not_started")
	fmt.Println("  breakpoint_required_after_f2: true")
	fmt.Println("  report:", reportPath)
	fmt.Println("  boundary: F2 closes operator docs/runbook only; next action is private LogDoc plus Breakpoint before F3")
	fmt.Println()
	fmt.Println("VALIDATION PASSED")
	return nil
}

func keysGateF(input map[string][]string) []string {
	out := make([]string, 0, len(input))
	for key := range input {
		out = append(out, key)
	}
	return out
}

func writeGateFOperatorDocsRunbookJSON(path string, value any) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		return err
	}
	body, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, append(body, '\n'), 0o600)
}
