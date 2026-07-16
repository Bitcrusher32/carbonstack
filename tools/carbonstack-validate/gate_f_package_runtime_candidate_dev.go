package main

import (
	"bytes"
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

const gateFPackageRuntimeCandidateReportSchema = "carbonstack-gate-f-package-runtime-candidate-report/v0"
const gateFPackageRuntimeCandidateManifestSchema = "carbonstack-package-runtime-candidate-manifest/v0"

func (r *Runner) GateFPackageRuntimeCandidateDev() error {
	fmt.Println("========================================")
	fmt.Println("CarbonStack validation profile: gate-f-package-runtime-candidate-dev")
	fmt.Println("========================================")
	fmt.Println("status: dev/pre-alpha Gate F F6 package/runtime candidate validation")
	fmt.Println("scope: disposable package/runtime candidate root shape, root separation, release boundary, F5 trust posture dependency, and nonclaim preservation")
	fmt.Println("boundary: not release creation, not release upload, not package publication, not legacy package staging execution, not full-runtime-dev promotion")
	fmt.Println()

	if err := r.CheckRequiredPaths(); err != nil {
		return err
	}
	if err := r.CheckLiveGitUmbrella("gate-f-package-runtime-candidate-dev"); err != nil {
		return err
	}

	reportRoot := filepath.Join(os.TempDir(), "carbonstack-gate-f-package-runtime-candidate-dev")
	_ = os.RemoveAll(reportRoot)
	if err := os.MkdirAll(reportRoot, 0o700); err != nil {
		return err
	}

	candidateRoot := filepath.Join(reportRoot, "candidate-root")
	packageRoot := filepath.Join(candidateRoot, "package-root")
	runtimeRoot := filepath.Join(candidateRoot, "runtime-root")
	releaseRoot := filepath.Join(candidateRoot, "release-artifact-root")
	evidenceRoot := filepath.Join(candidateRoot, "evidence-root")
	reportDir := filepath.Join(candidateRoot, "reports")
	localStateRoot := filepath.Join(candidateRoot, "local-state")

	requiredDirs := []string{
		packageRoot,
		filepath.Join(packageRoot, "manifest"),
		filepath.Join(packageRoot, "sources"),
		filepath.Join(packageRoot, "docs"),
		filepath.Join(packageRoot, "registry"),
		runtimeRoot,
		filepath.Join(runtimeRoot, "config"),
		filepath.Join(runtimeRoot, "state", "cypher"),
		filepath.Join(runtimeRoot, "logs"),
		filepath.Join(runtimeRoot, "trust"),
		releaseRoot,
		evidenceRoot,
		reportDir,
		localStateRoot,
	}
	for _, dir := range requiredDirs {
		if err := os.MkdirAll(dir, 0o700); err != nil {
			return err
		}
	}

	heads, err := gateF6RepoHeads(map[string]string{
		"carbonstack":        r.CarbonStack,
		"carbonstack-comms":  r.Comms,
		"carbonstack-cypher": r.Cypher,
	})
	if err != nil {
		return err
	}

	osRepo := filepath.Join(filepath.Dir(r.CarbonStack), "carbonstack-os")
	if st, err := os.Stat(osRepo); err == nil && st.IsDir() {
		if head, err := gateF6GitValue(osRepo, "rev-parse", "HEAD"); err == nil {
			heads["carbonstack-os"] = head
		}
	}

	requiredPaths := []string{
		filepath.Join(r.CarbonStack, "registry", "commands.v0.yaml"),
		filepath.Join(r.CarbonStack, "registry", "COMMAND_REFERENCE.v0.md"),
		filepath.Join(r.CarbonStack, "docs", "282-v0.7.23-gate-f-f5-closure-v0.md"),
		filepath.Join(r.Comms, "docs", "32-gate-f-f5-basic-local-trust-closure-v0.md"),
		filepath.Join(r.Cypher, "cmd", "cypher"),
		filepath.Join(r.Cypher, "migrations"),
	}
	if err := gateF6RequirePaths(requiredPaths); err != nil {
		return err
	}

	stageScripts, err := gateF6StageRehearsePriorArt(map[string]string{
		"carbonstack":        r.CarbonStack,
		"carbonstack-comms":  r.Comms,
		"carbonstack-cypher": r.Cypher,
	})
	if err != nil {
		return err
	}

	trustPosturePath := filepath.Join(reportDir, "f5-basic-local-trust-posture.json")
	trustOut, err := gateF6RunCommand("F5 basic local trust posture dependency", r.Comms, []string{
		"go", "run", "./cmd/comms", "basic-local-trust-posture-dev",
		"--state", localStateRoot,
		"--subject-label", "package-runtime-candidate-probe",
		"--cypher-account", "acct-package-runtime-probe",
		"--cypher-device", "device-package-runtime-probe",
		"--comms-fingerprint", "fp-package-runtime-probe",
		"--openmls-device-label", "sidecar-package-runtime-probe",
		"--openmls-keypackage-ref", "kp-package-runtime-probe",
		"--relay-space", "relay-package-runtime-probe",
		"--report", trustPosturePath,
	})
	if err != nil {
		return err
	}
	trustReport, err := gateF6ReadJSON(trustPosturePath)
	if err != nil {
		return err
	}
	if trustReport["schema_version"] != "carbonstack-basic-local-trust-posture/v0" {
		return fmt.Errorf("unexpected F5 trust posture schema: %v", trustReport["schema_version"])
	}
	trustClaims, ok := trustReport["claims"].(map[string]any)
	if !ok {
		return fmt.Errorf("F5 trust posture claims object missing")
	}
	if trustClaims["verified_identity"] != false || trustClaims["trust_promotion"] != false || trustClaims["automatic_trust_promotion"] != false || trustClaims["cryptographic_binding_across_cypher_comms_openmls"] != false {
		return fmt.Errorf("F5 trust posture dependency made forbidden identity/trust claim")
	}

	cypherConfigPath := filepath.Join(reportDir, "cypher-config.json")
	if err := gateF6CypherConfigInspection(r.Cypher, filepath.Join(localStateRoot, "cypher"), cypherConfigPath); err != nil {
		return err
	}

	manifestPath := filepath.Join(packageRoot, "manifest", "package-runtime-candidate-manifest.json")
	manifest := map[string]any{
		"schema_version":          gateFPackageRuntimeCandidateManifestSchema,
		"created_at":              time.Now().UTC().Format(time.RFC3339),
		"candidate_root":          candidateRoot,
		"package_root":            packageRoot,
		"runtime_deployment_root": runtimeRoot,
		"release_artifact_root":   releaseRoot,
		"evidence_root":           evidenceRoot,
		"report_root":             reportDir,
		"local_state_root":        localStateRoot,
		"source_repositories":     heads,
		"required_profiles": []string{
			"runner.gate-f-release-package-surface-dev",
			"runner.gate-f-operator-docs-runbook-dev",
			"runner.gate-f-compat-rollback-observability-dev",
			"runner.gate-f-code-health-source-hygiene-dev",
			"runner.gate-f-basic-local-trust-posture-dev",
			"runner.gate-f-package-runtime-candidate-dev",
		},
		"root_separation_policy":                                "package root, release artifact root, runtime deployment root, evidence root, report root, and local state root must be distinct",
		"legacy_stage_rehearse_scripts_classified_not_executed": true,
		"release_created":                                       false,
		"release_uploaded":                                      false,
		"package_published":                                     false,
		"package_staging_executed":                              false,
		"full_runtime_dev_promoted":                             false,
		"service_or_systemd_started":                            false,
		"helper_install_started":                                false,
		"container_started":                                     false,
		"public_ingress_started":                                false,
		"tui_started":                                           false,
		"verified_identity_claimed":                             false,
		"trust_promotion_claimed":                               false,
		"secure_enrollment_claimed":                             false,
		"cryptographic_identity_binding_implemented":            false,
	}
	if err := gateF6WriteJSON(manifestPath, manifest); err != nil {
		return err
	}

	if err := gateF6ValidateRootSeparation([]string{candidateRoot, packageRoot, runtimeRoot, releaseRoot, evidenceRoot, reportDir, localStateRoot}); err != nil {
		return err
	}

	archives, err := gateF6ArchiveHits(candidateRoot)
	if err != nil {
		return err
	}
	if len(archives) > 0 {
		return fmt.Errorf("candidate root unexpectedly contains archive artifacts: %v", archives)
	}

	noPromotion, err := gateF6NoFullRuntimePromotion(r.CarbonStack)
	if err != nil {
		return err
	}

	validatorDir := filepath.Join(r.CarbonStack, "tools", "carbonstack-validate")
	referenceCheck, err := gateF6RunCommand("generated command reference current", r.CarbonStack, []string{"python3", "-B", "tools/registry/render-command-reference.py", "--check"})
	if err != nil {
		return err
	}
	registryLookup, err := gateF6RunCommand("F6 registry lookup", validatorDir, []string{"go", "run", ".", "--profile", "registry-lookup", "--registry-id", "runner.gate-f-package-runtime-candidate-dev"})
	if err != nil {
		return err
	}
	nonclaims, err := gateF6RunCommand("missing nonclaims scan", validatorDir, []string{"go", "run", ".", "--profile", "registry-lookup", "--list", "--missing-nonclaims"})
	if err != nil {
		return err
	}

	reportPath := filepath.Join(reportRoot, "gate-f-package-runtime-candidate-report.json")
	report := map[string]any{
		"schema_version":   gateFPackageRuntimeCandidateReportSchema,
		"profile":          "gate-f-package-runtime-candidate-dev",
		"created_at":       time.Now().UTC().Format(time.RFC3339),
		"gate_f_status":    "open_f1_f2_f3_f4_f5_f6_closed_f7_not_started",
		"gate_f_f6_status": "closed",
		"gate_f_f7_status": "not_started",
		"package_runtime_candidate_validation_implemented":      true,
		"disposable_candidate_root_validated":                   true,
		"candidate_root":                                        candidateRoot,
		"manifest_path":                                         manifestPath,
		"f5_basic_local_trust_posture_report":                   trustPosturePath,
		"f5_trust_posture_stdout_bytes":                         len(trustOut),
		"cypher_config_report":                                  cypherConfigPath,
		"stage_rehearse_prior_art":                              stageScripts,
		"legacy_stage_rehearse_scripts_classified_not_executed": true,
		"root_separation_validated":                             true,
		"candidate_archive_artifacts":                           archives,
		"no_full_runtime_promotion":                             noPromotion,
		"generated_reference_check_stdout_bytes":                len(referenceCheck),
		"registry_lookup_stdout_bytes":                          len(registryLookup),
		"missing_nonclaims_stdout_bytes":                        len(nonclaims),
		"release_created":                                       false,
		"release_uploaded":                                      false,
		"package_published":                                     false,
		"package_staging_executed":                              false,
		"public_package_artifact_created":                       false,
		"full_runtime_dev_promoted":                             false,
		"service_or_systemd_started":                            false,
		"helper_install_started":                                false,
		"container_started":                                     false,
		"public_ingress_started":                                false,
		"tui_started":                                           false,
		"migration_implemented":                                 false,
		"repair_implemented":                                    false,
		"destructive_cleanup_performed":                         false,
		"state_relocation_performed":                            false,
		"verified_identity_claimed":                             false,
		"trust_promotion_claimed":                               false,
		"secure_enrollment_claimed":                             false,
		"cryptographic_identity_binding_implemented":            false,
		"automatic_trust_promotion_allowed":                     false,
		"vault_claimed":                                         false,
		"backup_restore_claimed":                                false,
		"production_e2ee_claimed":                               false,
		"pq_hybrid_claimed":                                     false,
		"android_claimed":                                       false,
		"carbonstack_os_claimed":                                false,
		"decisions": []string{
			"F6 validates disposable package/runtime candidate root shape only",
			"legacy stage and rehearse scripts are prior art and were not executed",
			"package root, release artifact root, runtime deployment root, evidence root, report root, and local state root are separated",
			"F5 basic local trust posture is included as a nonclaim-aware dependency, not verified identity",
			"release creation, upload, publication, package staging, and full-runtime-dev promotion remain false",
			"F7 should be v0.8.0 release-candidate closure matrix and manual release handoff",
		},
	}
	if err := gateF6WriteJSON(reportPath, report); err != nil {
		return err
	}

	fmt.Println()
	fmt.Println("gate-f-package-runtime-candidate-dev profile result:")
	fmt.Println("  PASS: disposable package/runtime candidate root validated")
	fmt.Println("  PASS: package/release/runtime/evidence/report/local-state roots separated")
	fmt.Println("  PASS: candidate manifest written")
	fmt.Println("  PASS: F5 basic local trust posture dependency preserved nonclaims")
	fmt.Println("  PASS: Cypher config inspection terminates with explicit env")
	fmt.Println("  PASS: legacy stage/rehearse scripts classified but not executed")
	fmt.Println("  PASS: no public package archive artifacts created")
	fmt.Println("  PASS: full-runtime-dev remains unpromoted")
	fmt.Println("  PASS: registry/reference/nonclaim checks passed")
	fmt.Println("  report:", reportPath)
	fmt.Println("  manifest:", manifestPath)
	fmt.Println("  candidate_root:", candidateRoot)
	fmt.Println("  stage_rehearse_script_count:", len(stageScripts))
	fmt.Println("  gate_f_status: open_f1_f2_f3_f4_f5_f6_closed_f7_not_started")
	fmt.Println("  gate_f_f6_status: closed")
	fmt.Println("  gate_f_f7_status: not_started")
	fmt.Println("  package_runtime_candidate_validation_implemented: true")
	fmt.Println("  release_created: false")
	fmt.Println("  release_uploaded: false")
	fmt.Println("  package_published: false")
	fmt.Println("  package_staging_executed: false")
	fmt.Println("  full_runtime_dev_promoted: false")
	fmt.Println("  boundary: F6 closes validation of disposable package/runtime candidate shape only; Gate F remains open")
	fmt.Println()
	fmt.Println("VALIDATION PASSED")
	return nil
}

func gateF6RepoHeads(repos map[string]string) (map[string]string, error) {
	out := map[string]string{}
	for name, dir := range repos {
		head, err := gateF6GitValue(dir, "rev-parse", "HEAD")
		if err != nil {
			return nil, err
		}
		out[name] = head
	}
	return out, nil
}

func gateF6GitValue(dir string, args ...string) (string, error) {
	cmd := exec.Command("git", args...)
	cmd.Dir = dir
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}

func gateF6RequirePaths(paths []string) error {
	for _, path := range paths {
		if _, err := os.Stat(path); err != nil {
			return fmt.Errorf("required path missing: %s: %w", path, err)
		}
	}
	return nil
}

func gateF6StageRehearsePriorArt(repos map[string]string) ([]map[string]any, error) {
	nameRx := regexp.MustCompile("(?i)(stage|rehearse|package|release|snapshot|manifest|checksum|runtime)")
	dangerRx := regexp.MustCompile("(?i)(git push|git tag|curl|scp|rsync|rm -rf|rm -r|docker|podman|systemctl|sudo|gpg|openssl|upload|release create|gh release|tea release|forgejo|gitea|tar |zip |sha256sum|checksum|manifest)")
	var rows []map[string]any
	for repo, root := range repos {
		files, err := gateF6GitFiles(root)
		if err != nil {
			return nil, err
		}
		for _, rel := range files {
			if !nameRx.MatchString(rel) || !gateF6TextCandidate(rel) {
				continue
			}
			body, err := os.ReadFile(filepath.Join(root, rel))
			if err != nil {
				continue
			}
			text := string(body)
			rows = append(rows, map[string]any{
				"repo":                                repo,
				"path":                                rel,
				"line_count":                          len(strings.Split(text, "\n")),
				"has_release_action_or_archive_terms": dangerRx.MatchString(text),
				"classification":                      "prior_art_static_only_not_executed",
			})
		}
	}
	sort.Slice(rows, func(i, j int) bool {
		return fmt.Sprint(rows[i]["repo"], rows[i]["path"]) < fmt.Sprint(rows[j]["repo"], rows[j]["path"])
	})
	return rows, nil
}

func gateF6RunCommand(name string, dir string, args []string) ([]byte, error) {
	fmt.Println("----------------------------------------")
	fmt.Println("STEP:", name)
	fmt.Println("DIR: ", dir)
	fmt.Println("CMD: ", strings.Join(args, " "))
	fmt.Println("----------------------------------------")
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Dir = dir
	cmd.Env = append(os.Environ(), "PYTHONDONTWRITEBYTECODE=1")
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if stdout.Len() > 0 {
		fmt.Print(stdout.String())
	}
	if stderr.Len() > 0 {
		fmt.Print(stderr.String())
	}
	if err != nil {
		return append(stdout.Bytes(), stderr.Bytes()...), fmt.Errorf("%s failed: %w", name, err)
	}
	fmt.Println("PASS:", name)
	fmt.Println()
	return append(stdout.Bytes(), stderr.Bytes()...), nil
}

func gateF6CypherConfigInspection(cypherDir string, stateRoot string, reportPath string) error {
	migrations := filepath.Join(stateRoot, "migrations")
	state := filepath.Join(stateRoot, "state")
	if err := os.MkdirAll(migrations, 0o700); err != nil {
		return err
	}
	if err := os.MkdirAll(state, 0o700); err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(migrations, "001_test.sql"), []byte("select 1;\n"), 0o600); err != nil {
		return err
	}
	env := append(os.Environ(),
		"CYPHER_ADDR=127.0.0.1:19980",
		"CYPHER_DB="+filepath.Join(state, "cypher.db"),
		"CYPHER_MIGRATIONS="+migrations,
		"CYPHER_DEV_INVITE=gate-f-f6",
	)
	cmd := exec.Command("go", "run", "./cmd/cypher", "--print-config")
	cmd.Dir = cypherDir
	cmd.Env = env
	out, err := cmd.Output()
	if err != nil {
		return err
	}
	var decoded map[string]any
	if err := json.Unmarshal(out, &decoded); err != nil {
		return err
	}
	if decoded["schema_version"] != "carbonstack-cypher-config-inspection/v0" {
		return fmt.Errorf("unexpected Cypher config schema: %v", decoded["schema_version"])
	}
	if decoded["starts_server"] != false || decoded["terminating_inspection"] != true || decoded["db_path_source"] != "env" {
		return fmt.Errorf("Cypher config inspection did not preserve terminating explicit-env posture")
	}
	return gateF6WriteBytes(reportPath, out)
}

func gateF6ValidateRootSeparation(paths []string) error {
	seen := map[string]bool{}
	for _, path := range paths {
		clean, err := filepath.Abs(path)
		if err != nil {
			return err
		}
		if seen[clean] {
			return fmt.Errorf("duplicate root path: %s", clean)
		}
		seen[clean] = true
	}
	return nil
}

func gateF6ArchiveHits(root string) ([]string, error) {
	var hits []string
	err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return nil
		}
		name := strings.ToLower(d.Name())
		if strings.HasSuffix(name, ".tgz") || strings.HasSuffix(name, ".tar.gz") || strings.HasSuffix(name, ".zip") {
			rel, _ := filepath.Rel(root, path)
			hits = append(hits, rel)
		}
		return nil
	})
	sort.Strings(hits)
	return hits, err
}

func gateF6NoFullRuntimePromotion(carbonRoot string) (bool, error) {
	registry, err := os.ReadFile(filepath.Join(carbonRoot, "registry", "commands.v0.yaml"))
	if err != nil {
		return false, err
	}
	main, err := os.ReadFile(filepath.Join(carbonRoot, "tools", "carbonstack-validate", "main.go"))
	if err != nil {
		return false, err
	}
	registryText := string(registry)
	mainText := string(main)
	activeID := regexp.MustCompile("(?m)^\\s*-\\s+id:\\s*runner\\.full-runtime-dev\\s*$")
	if activeID.MatchString(registryText) {
		return false, fmt.Errorf("active runner.full-runtime-dev registry id detected")
	}
	if strings.Contains(mainText, "case \"full-runtime-dev\":") {
		return false, fmt.Errorf("full-runtime-dev dispatch detected")
	}
	if strings.Contains(registryText, "--profile full-runtime-dev") {
		return false, fmt.Errorf("full-runtime-dev registry command detected")
	}
	return true, nil
}

func gateF6GitFiles(root string) ([]string, error) {
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

func gateF6TextCandidate(rel string) bool {
	ext := strings.ToLower(filepath.Ext(rel))
	switch ext {
	case ".go", ".md", ".sh", ".py", ".yaml", ".yml", ".json", ".toml", ".sql", ".txt":
		return true
	default:
		return false
	}
}

func gateF6ReadJSON(path string) (map[string]any, error) {
	body, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var out map[string]any
	if err := json.Unmarshal(body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

func gateF6WriteJSON(path string, value any) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		return err
	}
	body, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		return err
	}
	return gateF6WriteBytes(path, append(body, '\n'))
}

func gateF6WriteBytes(path string, body []byte) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		return err
	}
	return os.WriteFile(path, body, 0o600)
}
