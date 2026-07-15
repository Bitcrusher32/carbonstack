package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"
)

const gateFReleasePackageSurfaceReportSchema = "carbonstack-gate-f-release-package-surface-report/v0"

func (r *Runner) GateFReleasePackageSurfaceDev() error {
	fmt.Println("========================================")
	fmt.Println("CarbonStack validation profile: gate-f-release-package-surface-dev")
	fmt.Println("========================================")
	fmt.Println("status: dev/pre-alpha Gate F F1 release/package/runtime surface classification")
	fmt.Println("scope: static classification of release validation, package staging, runtime candidate, helper, and hygiene surfaces")
	fmt.Println("boundary: not release creation, not package publication, not package staging execution, not full-runtime-dev, not migration, not service/systemd/helper install")
	fmt.Println()

	if err := r.CheckRequiredPaths(); err != nil {
		return err
	}
	if err := r.CheckLiveGitUmbrella("gate-f-release-package-surface-dev"); err != nil {
		return err
	}

	reportRoot := filepath.Join(os.TempDir(), "carbonstack-gate-f-release-package-surface-dev")
	_ = os.RemoveAll(reportRoot)
	if err := os.MkdirAll(reportRoot, 0o700); err != nil {
		return err
	}
	reportPath := filepath.Join(reportRoot, "gate-f-release-package-surface-report.json")

	repos := map[string]string{
		"carbonstack":        r.CarbonStack,
		"carbonstack-comms":  r.Comms,
		"carbonstack-cypher": r.Cypher,
	}
	osRepo := filepath.Join(filepath.Dir(r.CarbonStack), "carbonstack-os")
	if st, err := os.Stat(osRepo); err == nil && st.IsDir() {
		repos["carbonstack-os"] = osRepo
	}

	surfaces, warnings := collectGateFReleasePackageSurfaces(repos)

	registryText, err := os.ReadFile(filepath.Join(r.CarbonStack, "registry", "commands.v0.yaml"))
	if err != nil {
		return err
	}
	registry := string(registryText)
	requiredIDs := []string{
		"runner.full-validate-release",
		"runner.release-snapshot",
		"runner.local-cypher",
		"runner.integrated-runtime-dev",
		"runner.gate-e-native-deployment-dev",
		"cypher.config-inspection",
		"runner.gate-e-native-deployment-closure-dev",
	}
	registryPresence := map[string]bool{}
	for _, id := range requiredIDs {
		present := strings.Contains(registry, "id: "+id)
		registryPresence[id] = present
		if !present {
			warnings = append(warnings, "missing registry id: "+id)
		}
	}

	stageScripts := filterSurfacePaths(surfaces, "package_staging_script")
	rehearseScripts := filterSurfacePaths(surfaces, "package_rehearsal_script")
	releaseProfiles := filterSurfacePaths(surfaces, "release_validation_profile")
	runtimeProfiles := filterSurfacePaths(surfaces, "runtime_validation_profile")

	if len(stageScripts) == 0 {
		warnings = append(warnings, "no package staging scripts classified")
	}
	if len(releaseProfiles) == 0 {
		warnings = append(warnings, "no release validation profiles classified")
	}

	cypherDBHits := findCypherDBHits(repos)

	report := map[string]any{
		"schema_version":                     gateFReleasePackageSurfaceReportSchema,
		"profile":                            "gate-f-release-package-surface-dev",
		"created_at":                         time.Now().UTC().Format(time.RFC3339),
		"gate_d_status":                      "closed",
		"gate_e_status":                      "closed",
		"gate_f_status":                      "open_f1_closed_f2_not_started",
		"gate_f_f1_status":                   "closed",
		"gate_f_f2_status":                   "not_started",
		"breakpoint_deferred_until_after_f2": true,
		"manual_release_creation_planned":    true,
		"release_creation_implemented":       false,
		"release_upload_implemented":         false,
		"package_publication_implemented":    false,
		"package_staging_executed":           false,
		"full_runtime_dev_promoted":          false,
		"service_or_systemd_started":         false,
		"helper_install_started":             false,
		"container_started":                  false,
		"public_ingress_started":             false,
		"tui_started":                        false,
		"migration_implemented":              false,
		"destructive_cleanup_performed":      false,
		"trust_promotion_claimed":            false,
		"verified_identity_claimed":          false,
		"vault_claimed":                      false,
		"backup_restore_claimed":             false,
		"pq_hybrid_claimed":                  false,
		"android_claimed":                    false,
		"carbonstack_os_claimed":             false,
		"registry_presence":                  registryPresence,
		"classified_surface_count":           len(surfaces),
		"package_staging_scripts":            stageScripts,
		"package_rehearsal_scripts":          rehearseScripts,
		"release_validation_profiles":        releaseProfiles,
		"runtime_validation_profiles":        runtimeProfiles,
		"cypher_db_hits":                     cypherDBHits,
		"warnings":                           warnings,
		"surfaces":                           surfaces,
		"decisions": []string{
			"full-validate-release remains release-package validation, not release creation",
			"release-snapshot remains package-root layout/checksum/core validation, not package creation or upload",
			"stage-v0.7.0-package.sh is the strongest historical staging scaffold, but v0.8.0 staging is not implemented by F1",
			"v0.8.0 release creation remains manual operator flow using prior release scaffolds",
			"package/runtime candidate validation is deferred to a later Gate F subgate",
			"full-runtime-dev remains reserved until a coherent package/runtime candidate validates the preferred lifecycle",
			"repo-root cypher.db is hygiene classification only and must not be destructively cleaned in F1",
		},
		"candidate_next": "Gate F F2 operator docs and runbook closure before breakpoint",
	}
	if err := writeGateFReleasePackageSurfaceJSON(reportPath, report); err != nil {
		return err
	}

	for _, warning := range warnings {
		fmt.Println("warning:", warning)
	}

	fmt.Println()
	fmt.Println("gate-f-release-package-surface-dev profile result:")
	fmt.Println("  PASS: release/package/runtime/helper surfaces classified")
	fmt.Println("  PASS: release validation profiles remain separate from release creation")
	fmt.Println("  PASS: package staging scripts classified without execution")
	fmt.Println("  PASS: package/runtime candidate deferred")
	fmt.Println("  PASS: full-runtime-dev remains unpromoted")
	fmt.Println("  PASS: breakpoint deferred until after F2 by operator decision")
	fmt.Println("  report:", reportPath)
	fmt.Println("  classified_surface_count:", len(surfaces))
	fmt.Println("  package_staging_script_count:", len(stageScripts))
	fmt.Println("  package_rehearsal_script_count:", len(rehearseScripts))
	fmt.Println("  release_validation_profile_count:", len(releaseProfiles))
	fmt.Println("  runtime_validation_profile_count:", len(runtimeProfiles))
	fmt.Println("  cypher_db_hit_count:", len(cypherDBHits))
	fmt.Println("  gate_f_status: open_f1_closed_f2_not_started")
	fmt.Println("  gate_f_f1_status: closed")
	fmt.Println("  gate_f_f2_status: not_started")
	fmt.Println("  release_creation_implemented: false")
	fmt.Println("  package_publication_implemented: false")
	fmt.Println("  package_staging_executed: false")
	fmt.Println("  full_runtime_dev_promoted: false")
	fmt.Println("  migration_implemented: false")
	fmt.Println("  boundary: F1 classifies authority only; F2 operator docs/runbook should follow before breakpoint")
	fmt.Println()
	fmt.Println("VALIDATION PASSED")
	return nil
}

func collectGateFReleasePackageSurfaces(repos map[string]string) ([]map[string]any, []string) {
	var surfaces []map[string]any
	var warnings []string

	for repo, root := range repos {
		files, err := gitTrackedGateFFiles(root)
		if err != nil {
			warnings = append(warnings, repo+": git ls-files failed: "+err.Error())
			continue
		}
		for _, rel := range files {
			kind, role, include := classifyGateFFile(repo, rel)
			if !include {
				continue
			}
			surface := map[string]any{
				"repo":      repo,
				"path":      rel,
				"kind":      kind,
				"role":      role,
				"f1_action": "classify_only",
				"executed":  false,
				"mutated":   false,
			}
			if rel == "scripts/stage-v0.7.0-package.sh" {
				surface["f1_note"] = "latest historical package staging scaffold; candidate input for later v0.8.0 staging work"
			}
			if rel == "tools/carbonstack-validate/release_snapshot.go" {
				surface["f1_note"] = "release package layout/checksum/core validation profile implementation"
			}
			if rel == "tools/carbonstack-validate/main.go" {
				surface["f1_note"] = "contains release profile dispatch and full/full-validate-release naming boundary"
			}
			surfaces = append(surfaces, surface)
		}
	}

	sort.Slice(surfaces, func(i, j int) bool {
		ai := surfaces[i]["repo"].(string) + "/" + surfaces[i]["path"].(string)
		aj := surfaces[j]["repo"].(string) + "/" + surfaces[j]["path"].(string)
		return ai < aj
	})
	return surfaces, warnings
}

func gitTrackedGateFFiles(root string) ([]string, error) {
	cmd := exec.Command("git", "ls-files")
	cmd.Dir = root
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	var files []string
	for _, line := range strings.Split(string(out), "\n") {
		line = strings.TrimSpace(line)
		if line != "" {
			files = append(files, line)
		}
	}
	return files, nil
}

func classifyGateFFile(repo string, rel string) (string, string, bool) {
	lower := strings.ToLower(rel)

	if repo == "carbonstack" && strings.HasPrefix(lower, "scripts/stage-v") && strings.HasSuffix(lower, "-package.sh") {
		return "script", "package_staging_script", true
	}
	if repo == "carbonstack" && strings.HasPrefix(lower, "scripts/rehearse-v") && strings.HasSuffix(lower, "-package.sh") {
		return "script", "package_rehearsal_script", true
	}
	if repo == "carbonstack" && rel == "tools/carbonstack-validate/release_snapshot.go" {
		return "runner_profile_source", "release_validation_profile", true
	}
	if repo == "carbonstack" && rel == "tools/carbonstack-validate/checksums.go" {
		return "runner_profile_source", "release_checksum_helper", true
	}
	if repo == "carbonstack" && rel == "tools/carbonstack-validate/local_cypher.go" {
		return "runner_profile_source", "release_validation_component", true
	}
	if repo == "carbonstack" && rel == "tools/carbonstack-validate/integrated_runtime_dev.go" {
		return "runner_profile_source", "runtime_validation_profile", true
	}
	if repo == "carbonstack" && rel == "tools/carbonstack-validate/gate_d_runtime_aggregate_dev.go" {
		return "runner_profile_source", "runtime_validation_profile", true
	}
	if repo == "carbonstack" && rel == "tools/carbonstack-validate/gate_e_native_deployment_dev.go" {
		return "runner_profile_source", "deployment_validation_profile", true
	}
	if repo == "carbonstack" && rel == "tools/carbonstack-validate/gate_e_native_deployment_closure_dev.go" {
		return "runner_profile_source", "deployment_closure_profile", true
	}
	if repo == "carbonstack" && rel == "tools/carbonstack-validate/main.go" {
		return "runner_dispatch_source", "release_runtime_profile_dispatch", true
	}
	if repo == "carbonstack" && rel == "registry/commands.v0.yaml" {
		return "registry", "command_authority_source", true
	}
	if repo == "carbonstack" && rel == "registry/COMMAND_REFERENCE.v0.md" {
		return "registry", "generated_command_reference", true
	}
	if repo == "carbonstack" && rel == "registry/COMMAND_BOUNDARY_TABLE.v0.md" {
		return "registry", "command_boundary_table", true
	}

	docMatch := regexp.MustCompile(`(?i)(release|package|snapshot|runtime|deploy|gate-f|v0\.8|validation|runbook|compat|rollback|migration)`)
	if strings.HasPrefix(lower, "docs/") && docMatch.MatchString(rel) {
		return "docs", "historical_or_current_gate_f_context", true
	}

	if strings.Contains(lower, "release") || strings.Contains(lower, "package") || strings.Contains(lower, "runtime") || strings.Contains(lower, "validation") {
		if strings.HasSuffix(lower, ".go") || strings.HasSuffix(lower, ".sh") || strings.HasSuffix(lower, ".py") || strings.HasSuffix(lower, ".ps1") || strings.HasSuffix(lower, ".md") {
			return "source_or_doc", "gate_f_related_context", true
		}
	}

	return "", "", false
}

func filterSurfacePaths(surfaces []map[string]any, role string) []string {
	var out []string
	for _, surface := range surfaces {
		if surface["role"] == role {
			out = append(out, surface["repo"].(string)+"/"+surface["path"].(string))
		}
	}
	sort.Strings(out)
	return out
}

func findCypherDBHits(repos map[string]string) []string {
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
	sort.Strings(hits)
	return hits
}

func writeGateFReleasePackageSurfaceJSON(path string, value any) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		return err
	}
	body, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, append(body, '\n'), 0o600)
}
