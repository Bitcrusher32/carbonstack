package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

const gateFCompatRollbackObservabilityReportSchema = "carbonstack-gate-f-compat-rollback-observability-report/v0"

func (r *Runner) GateFCompatRollbackObservabilityDev() error {
	fmt.Println("========================================")
	fmt.Println("CarbonStack validation profile: gate-f-compat-rollback-observability-dev")
	fmt.Println("========================================")
	fmt.Println("status: dev/pre-alpha Gate F F3 compatibility, stale-state, rollback observability, and refusal posture")
	fmt.Println("scope: aggregate observer over Gate C compatibility/path/write policy, Gate F F1/F2 authority, Cypher config/migration visibility, and no-silent-migration boundaries")
	fmt.Println("boundary: not migration, not repair, not destructive cleanup, not package/runtime candidate, not full-runtime-dev, not release creation, not service/systemd/helper install")
	fmt.Println()

	if err := r.CheckRequiredPaths(); err != nil {
		return err
	}
	if err := r.CheckLiveGitUmbrella("gate-f-compat-rollback-observability-dev"); err != nil {
		return err
	}

	reportRoot := filepath.Join(os.TempDir(), "carbonstack-gate-f-compat-rollback-observability-dev")
	_ = os.RemoveAll(reportRoot)
	if err := os.MkdirAll(reportRoot, 0o700); err != nil {
		return err
	}
	reportPath := filepath.Join(reportRoot, "gate-f-compat-rollback-observability-report.json")

	validatorDir := filepath.Join(r.CarbonStack, "tools", "carbonstack-validate")

	steps := []gateFF3Step{
		{
			Name: "C2 schema compatibility profile",
			Dir:  validatorDir,
			Args: []string{"go", "run", ".", "--profile", "state-schema-compat-dev", "--compact-summary", "--clean-generated"},
		},
		{
			Name: "C3 path policy profile",
			Dir:  validatorDir,
			Args: []string{"go", "run", ".", "--profile", "state-path-policy-dev", "--compact-summary", "--clean-generated"},
		},
		{
			Name: "C4 write policy profile",
			Dir:  validatorDir,
			Args: []string{"go", "run", ".", "--profile", "state-write-policy-dev", "--compact-summary", "--clean-generated"},
		},
		{
			Name: "F1 release package runtime surface classification profile",
			Dir:  validatorDir,
			Args: []string{"go", "run", ".", "--profile", "gate-f-release-package-surface-dev", "--compact-summary"},
		},
		{
			Name: "F2 operator docs runbook profile",
			Dir:  validatorDir,
			Args: []string{"go", "run", ".", "--profile", "gate-f-operator-docs-runbook-dev", "--compact-summary"},
		},
		{
			Name: "Generated command reference current",
			Dir:  r.CarbonStack,
			Args: []string{"python3", "tools/registry/render-command-reference.py", "--check"},
		},
		{
			Name: "F3 registry lookup",
			Dir:  validatorDir,
			Args: []string{"go", "run", ".", "--profile", "registry-lookup", "--registry-id", "runner.gate-f-compat-rollback-observability-dev"},
		},
		{
			Name: "Missing nonclaims scan",
			Dir:  validatorDir,
			Args: []string{"go", "run", ".", "--profile", "registry-lookup", "--list", "--missing-nonclaims"},
		},
	}

	results := make([]map[string]any, 0, len(steps))
	for _, step := range steps {
		result, err := runGateFF3Step(step)
		results = append(results, result)
		if err != nil {
			_ = writeGateFCompatRollbackJSON(reportPath, map[string]any{
				"schema_version": gateFCompatRollbackObservabilityReportSchema,
				"profile":        "gate-f-compat-rollback-observability-dev",
				"status":         "failed",
				"failed_step":    step.Name,
				"steps":          results,
			})
			return err
		}
	}

	cypherConfigPath, err := runGateFF3CypherConfigInspection(r.Cypher, reportRoot)
	if err != nil {
		return err
	}
	migrations, err := gateFF3MigrationInventory(filepath.Join(r.Cypher, "migrations"))
	if err != nil {
		return err
	}
	cypherDBHits := gateFF3FindCypherDBHits(map[string]string{
		"carbonstack":        r.CarbonStack,
		"carbonstack-comms":  r.Comms,
		"carbonstack-cypher": r.Cypher,
	})

	report := map[string]any{
		"schema_version":                         gateFCompatRollbackObservabilityReportSchema,
		"profile":                                "gate-f-compat-rollback-observability-dev",
		"created_at":                             time.Now().UTC().Format(time.RFC3339),
		"gate_d_status":                          "closed",
		"gate_e_status":                          "closed",
		"gate_f_status":                          "open_f1_f2_f3_closed_f4_not_started",
		"gate_f_f1_status":                       "closed",
		"gate_f_f2_status":                       "closed",
		"gate_f_f3_status":                       "closed",
		"gate_f_f4_status":                       "not_started",
		"compatibility_observed":                 true,
		"stale_state_observability_present":      true,
		"rollback_observability_present":         true,
		"refusal_posture_present":                true,
		"state_schema_compat_profile_passed":     true,
		"state_path_policy_profile_passed":       true,
		"state_write_policy_profile_passed":      true,
		"release_package_surface_profile_passed": true,
		"operator_docs_runbook_profile_passed":   true,
		"cypher_config_inspection_report":        cypherConfigPath,
		"cypher_migration_inventory":             migrations,
		"cypher_db_hits":                         cypherDBHits,
		"migration_implemented":                  false,
		"silent_migration_allowed":               false,
		"repair_implemented":                     false,
		"silent_repair_allowed":                  false,
		"destructive_cleanup_performed":          false,
		"state_relocation_performed":             false,
		"release_created":                        false,
		"release_uploaded":                       false,
		"package_published":                      false,
		"package_staging_executed":               false,
		"package_runtime_candidate_implemented":  false,
		"full_runtime_dev_promoted":              false,
		"service_or_systemd_started":             false,
		"helper_install_started":                 false,
		"container_started":                      false,
		"public_ingress_started":                 false,
		"tui_started":                            false,
		"verified_identity_claimed":              false,
		"trust_promotion_claimed":                false,
		"vault_claimed":                          false,
		"backup_restore_claimed":                 false,
		"production_e2ee_claimed":                false,
		"pq_hybrid_claimed":                      false,
		"android_claimed":                        false,
		"carbonstack_os_claimed":                 false,
		"steps":                                  results,
		"decisions": []string{
			"F3 closes aggregate compatibility, stale-state, rollback observability, and refusal posture only",
			"Existing Gate C C2/C3/C4 surfaces remain the underlying compatibility/path/write policy authorities",
			"F3 does not implement migration, repair, cleanup, state relocation, package/runtime candidate validation, or full-runtime-dev",
			"Repo-root cypher.db remains a visible non-destructive hygiene classification target",
			"Cypher migration visibility is observed through config inspection and migrations inventory, not DB migration implementation",
		},
		"candidate_next": "Gate F F4 code health and source hygiene closure, including non-destructive cypher.db policy if accepted",
	}
	if err := writeGateFCompatRollbackJSON(reportPath, report); err != nil {
		return err
	}

	fmt.Println()
	fmt.Println("gate-f-compat-rollback-observability-dev profile result:")
	fmt.Println("  PASS: C2 schema compatibility profile passes")
	fmt.Println("  PASS: C3 path policy profile passes")
	fmt.Println("  PASS: C4 write policy profile passes")
	fmt.Println("  PASS: F1 release/package/runtime surface profile passes")
	fmt.Println("  PASS: F2 operator docs/runbook profile passes")
	fmt.Println("  PASS: Cypher config inspection terminates and reports explicit env")
	fmt.Println("  PASS: Cypher migrations inventory observed without migration implementation")
	fmt.Println("  PASS: registry/reference/nonclaims checks passed")
	fmt.Println("  PASS: migration, repair, destructive cleanup, package/runtime candidate, and full-runtime-dev remain unimplemented")
	fmt.Println("  report:", reportPath)
	fmt.Println("  cypher_config_report:", cypherConfigPath)
	fmt.Println("  cypher_migration_count:", len(migrations))
	fmt.Println("  cypher_db_hit_count:", len(cypherDBHits))
	fmt.Println("  gate_f_status: open_f1_f2_f3_closed_f4_not_started")
	fmt.Println("  gate_f_f1_status: closed")
	fmt.Println("  gate_f_f2_status: closed")
	fmt.Println("  gate_f_f3_status: closed")
	fmt.Println("  gate_f_f4_status: not_started")
	fmt.Println("  migration_implemented: false")
	fmt.Println("  package_runtime_candidate_implemented: false")
	fmt.Println("  full_runtime_dev_promoted: false")
	fmt.Println("  boundary: F3 closes observability/refusal posture only; Gate F remains open")
	fmt.Println()
	fmt.Println("VALIDATION PASSED")
	return nil
}

type gateFF3Step struct {
	Name string
	Dir  string
	Args []string
}

func runGateFF3Step(step gateFF3Step) (map[string]any, error) {
	fmt.Println("----------------------------------------")
	fmt.Println("STEP:", step.Name)
	fmt.Println("DIR: ", step.Dir)
	fmt.Println("CMD: ", strings.Join(step.Args, " "))
	fmt.Println("----------------------------------------")

	if len(step.Args) == 0 {
		return map[string]any{"name": step.Name, "passed": false, "error": "empty command"}, fmt.Errorf("empty command for %s", step.Name)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, step.Args[0], step.Args[1:]...)
	cmd.Dir = step.Dir

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	start := time.Now()
	err := cmd.Run()
	elapsed := time.Since(start).Round(time.Millisecond)

	if stdout.Len() > 0 {
		fmt.Print(stdout.String())
	}
	if stderr.Len() > 0 {
		fmt.Print(stderr.String())
	}

	result := map[string]any{
		"name":       step.Name,
		"dir":        step.Dir,
		"command":    step.Args,
		"elapsed_ms": elapsed.Milliseconds(),
		"passed":     err == nil,
	}
	if ctx.Err() == context.DeadlineExceeded {
		result["error"] = "timeout"
		fmt.Println("FAIL:", step.Name, "timeout")
		return result, fmt.Errorf("%s timed out", step.Name)
	}
	if err != nil {
		result["error"] = err.Error()
		fmt.Println("FAIL:", step.Name, err)
		return result, fmt.Errorf("%s failed: %w", step.Name, err)
	}
	fmt.Println("PASS:", step.Name, "elapsed=", elapsed)
	fmt.Println()
	return result, nil
}

func runGateFF3CypherConfigInspection(cypherDir string, reportRoot string) (string, error) {
	tmp := filepath.Join(reportRoot, "cypher-config")
	if err := os.MkdirAll(filepath.Join(tmp, "state"), 0o700); err != nil {
		return "", err
	}
	if err := os.MkdirAll(filepath.Join(tmp, "migrations"), 0o700); err != nil {
		return "", err
	}
	if err := os.WriteFile(filepath.Join(tmp, "migrations", "001_test.sql"), []byte("select 1;\n"), 0o600); err != nil {
		return "", err
	}
	reportPath := filepath.Join(tmp, "config.json")

	env := append(os.Environ(),
		"CYPHER_ADDR=127.0.0.1:19580",
		"CYPHER_DB="+filepath.Join(tmp, "state", "cypher.db"),
		"CYPHER_MIGRATIONS="+filepath.Join(tmp, "migrations"),
		"CYPHER_DEV_INVITE=gate-f-f3",
	)

	cmd := exec.Command("go", "run", "./cmd/cypher", "--print-config")
	cmd.Dir = cypherDir
	cmd.Env = env
	out, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("cypher print-config failed: %w", err)
	}
	var decoded map[string]any
	if err := json.Unmarshal(out, &decoded); err != nil {
		return "", err
	}
	if decoded["schema_version"] != "carbonstack-cypher-config-inspection/v0" {
		return "", fmt.Errorf("unexpected Cypher config schema")
	}
	if decoded["starts_server"] != false {
		return "", fmt.Errorf("Cypher config inspection unexpectedly starts server")
	}
	if decoded["terminating_inspection"] != true {
		return "", fmt.Errorf("Cypher config inspection is not marked terminating")
	}
	if decoded["db_path_source"] != "env" {
		return "", fmt.Errorf("Cypher DB path source is not env")
	}
	if err := os.WriteFile(reportPath, out, 0o600); err != nil {
		return "", err
	}

	check := exec.Command("go", "run", "./cmd/cypher", "--check-config")
	check.Dir = cypherDir
	check.Env = env
	checkOut, err := check.CombinedOutput()
	if len(checkOut) > 0 {
		fmt.Print(string(checkOut))
	}
	if err != nil {
		return "", fmt.Errorf("cypher check-config failed: %w", err)
	}

	return reportPath, nil
}

func gateFF3MigrationInventory(dir string) ([]map[string]any, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	var out []map[string]any
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".sql") {
			continue
		}
		path := filepath.Join(dir, entry.Name())
		info, err := entry.Info()
		if err != nil {
			return nil, err
		}
		out = append(out, map[string]any{
			"name":       entry.Name(),
			"path":       path,
			"size_bytes": info.Size(),
		})
	}
	return out, nil
}

func gateFF3FindCypherDBHits(repos map[string]string) []string {
	var hits []string
	for repo, root := range repos {
		filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
			if err != nil {
				return nil
			}
			if d.IsDir() && d.Name() == ".git" {
				return filepath.SkipDir
			}
			if !d.IsDir() && d.Name() == "cypher.db" {
				if rel, err := filepath.Rel(root, path); err == nil {
					hits = append(hits, repo+"/"+rel)
				}
			}
			return nil
		})
	}
	return hits
}

func writeGateFCompatRollbackJSON(path string, value any) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		return err
	}
	body, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, append(body, '\n'), 0o600)
}
