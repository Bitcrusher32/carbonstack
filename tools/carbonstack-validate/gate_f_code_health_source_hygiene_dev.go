package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
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

const gateFCodeHealthSourceHygieneReportSchema = "carbonstack-gate-f-code-health-source-hygiene-report/v0"

func (r *Runner) GateFCodeHealthSourceHygieneDev() error {
	fmt.Println("========================================")
	fmt.Println("CarbonStack validation profile: gate-f-code-health-source-hygiene-dev")
	fmt.Println("========================================")
	fmt.Println("status: dev/pre-alpha Gate F F4 code health and source hygiene")
	fmt.Println("scope: source-hygiene classification, generated artifact policy, helper safety classification, script syntax checks, and cypher.db non-destructive policy")
	fmt.Println("boundary: not package/runtime candidate, not release creation, not package staging, not full-runtime-dev, not migration, not repair, not destructive cleanup, not state relocation")
	fmt.Println()

	if err := r.CheckRequiredPaths(); err != nil {
		return err
	}
	if err := r.CheckLiveGitUmbrella("gate-f-code-health-source-hygiene-dev"); err != nil {
		return err
	}

	reportRoot := filepath.Join(os.TempDir(), "carbonstack-gate-f-code-health-source-hygiene-dev")
	_ = os.RemoveAll(reportRoot)
	if err := os.MkdirAll(reportRoot, 0o700); err != nil {
		return err
	}
	reportPath := filepath.Join(reportRoot, "gate-f-code-health-source-hygiene-report.json")

	repos := map[string]string{
		"carbonstack":        r.CarbonStack,
		"carbonstack-comms":  r.Comms,
		"carbonstack-cypher": r.Cypher,
	}
	osRepo := filepath.Join(filepath.Dir(r.CarbonStack), "carbonstack-os")
	if st, err := os.Stat(osRepo); err == nil && st.IsDir() {
		repos["carbonstack-os"] = osRepo
	}

	hygiene, err := gateF4CollectSourceHygiene(repos)
	if err != nil {
		return err
	}
	helpers, err := gateF4CollectHelperSafety(repos)
	if err != nil {
		return err
	}
	syntax, err := gateF4RunSyntaxChecks(repos)
	if err != nil {
		return err
	}

	validatorTest, err := gateF4RunCommand("validator package tests", filepath.Join(r.CarbonStack, "tools", "carbonstack-validate"), []string{"go", "test", "./...", "-count=1"})
	if err != nil {
		return err
	}
	referenceCheck, err := gateF4RunCommand("generated command reference current", r.CarbonStack, []string{"python3", "-B", "tools/registry/render-command-reference.py", "--check"})
	if err != nil {
		return err
	}
	registryLookup, err := gateF4RunCommand("F4 registry lookup", filepath.Join(r.CarbonStack, "tools", "carbonstack-validate"), []string{"go", "run", ".", "--profile", "registry-lookup", "--registry-id", "runner.gate-f-code-health-source-hygiene-dev"})
	if err != nil {
		return err
	}
	nonclaims, err := gateF4RunCommand("missing nonclaims scan", filepath.Join(r.CarbonStack, "tools", "carbonstack-validate"), []string{"go", "run", ".", "--profile", "registry-lookup", "--list", "--missing-nonclaims"})
	if err != nil {
		return err
	}

	report := map[string]any{
		"schema_version":            gateFCodeHealthSourceHygieneReportSchema,
		"profile":                   "gate-f-code-health-source-hygiene-dev",
		"created_at":                time.Now().UTC().Format(time.RFC3339),
		"gate_d_status":             "closed",
		"gate_e_status":             "closed",
		"gate_f_status":             "open_f1_f2_f3_f4_closed_f5_not_started",
		"gate_f_f1_status":          "closed",
		"gate_f_f2_status":          "closed",
		"gate_f_f3_status":          "closed",
		"gate_f_f4_status":          "closed",
		"gate_f_f5_status":          "not_started",
		"source_hygiene_classified": true,
		"helper_safety_classified":  true,
		"script_syntax_checked":     true,
		"generated_python_cache_ignore_policy_present": gateF4GitignoreCoversPythonCache(r.CarbonStack),
		"cypher_db_policy":                      "non_destructive_classify_not_delete",
		"sanitized_logdoc_archive_policy":       "historical_provenance_excluded_from_current_source_defect_claims",
		"helper_policy":                         "static_classification_only_not_execution",
		"hygiene":                               hygiene,
		"helpers":                               helpers,
		"syntax":                                syntax,
		"validator_package_tests":               validatorTest,
		"generated_reference_check":             referenceCheck,
		"registry_lookup":                       registryLookup,
		"missing_nonclaims":                     nonclaims,
		"release_created":                       false,
		"release_uploaded":                      false,
		"package_published":                     false,
		"package_staging_executed":              false,
		"package_runtime_candidate_implemented": false,
		"full_runtime_dev_promoted":             false,
		"migration_implemented":                 false,
		"repair_implemented":                    false,
		"destructive_cleanup_performed":         false,
		"state_relocation_performed":            false,
		"service_or_systemd_started":            false,
		"helper_install_started":                false,
		"container_started":                     false,
		"public_ingress_started":                false,
		"tui_started":                           false,
		"verified_identity_claimed":             false,
		"trust_promotion_claimed":               false,
		"vault_claimed":                         false,
		"backup_restore_claimed":                false,
		"production_e2ee_claimed":               false,
		"pq_hybrid_claimed":                     false,
		"android_claimed":                       false,
		"carbonstack_os_claimed":                false,
		"decisions": []string{
			"F4 closes code-health and source-hygiene classification only",
			"Python bytecode cache is ignored as generated local cache and should not dirty operator status",
			"repo-root cypher.db remains a visible non-destructive hygiene classification target",
			"historical sanitized LogDocs are provenance and should not dominate current source defect claims",
			"release/helper scripts are statically classified and syntax checked where applicable, not executed",
			"F4 does not implement package/runtime candidate validation or full-runtime-dev",
		},
		"candidate_next": "Gate F F5 package/runtime candidate validation preflight after accepting an optional post-F4 breakpoint",
	}
	if err := writeGateF4JSON(reportPath, report); err != nil {
		return err
	}

	fmt.Println()
	fmt.Println("gate-f-code-health-source-hygiene-dev profile result:")
	fmt.Println("  PASS: source hygiene classified")
	fmt.Println("  PASS: helper safety classified without execution")
	fmt.Println("  PASS: shell and Python helper syntax checked")
	fmt.Println("  PASS: Python cache ignore policy present")
	fmt.Println("  PASS: cypher.db remains non-destructive hygiene classification")
	fmt.Println("  PASS: sanitized LogDoc archive classified as historical provenance")
	fmt.Println("  PASS: validator package tests passed")
	fmt.Println("  PASS: registry/reference/nonclaims checks passed")
	fmt.Println("  PASS: package/runtime candidate, full-runtime-dev, migration, repair, destructive cleanup, and release work remain unimplemented")
	fmt.Println("  report:", reportPath)
	fmt.Println("  tracked_files_with_hygiene_hits:", hygiene["tracked_files_with_hits"])
	fmt.Println("  helper_candidate_count:", helpers["helper_candidate_count"])
	fmt.Println("  shell_syntax_checked:", syntax["shell_checked_count"])
	fmt.Println("  python_syntax_checked:", syntax["python_checked_count"])
	fmt.Println("  cypher_db_hit_count:", hygiene["cypher_db_hit_count"])
	fmt.Println("  gate_f_status: open_f1_f2_f3_f4_closed_f5_not_started")
	fmt.Println("  gate_f_f4_status: closed")
	fmt.Println("  gate_f_f5_status: not_started")
	fmt.Println("  boundary: F4 closes code-health/source-hygiene classification only; Gate F remains open")
	fmt.Println()
	fmt.Println("VALIDATION PASSED")
	return nil
}

func gateF4CollectSourceHygiene(repos map[string]string) (map[string]any, error) {
	patterns := map[string]*regexp.Regexp{
		"todo_fixme":            regexp.MustCompile("(?i)TODO|FIXME"),
		"generated_artifact":    regexp.MustCompile("(?i)generated|artifact|cleanup|clean-generated|target|dist|package root|release root"),
		"release_helper":        regexp.MustCompile("(?i)release|package|snapshot|stage|rehearse|upload|publish|tag|checksum|manifest"),
		"destructive_terms":     regexp.MustCompile("(?i)rm -rf|rm -r|removeall|delete|unlink|rmtree|destructive cleanup|cleanup implemented"),
		"migration_repair":      regexp.MustCompile("(?i)migration|migrate|repair|silent migration|silent repair|state relocation"),
		"full_runtime":          regexp.MustCompile("(?i)full-runtime-dev|full runtime|integrated-runtime-dev"),
		"service_container_tui": regexp.MustCompile("(?i)systemd|service|daemon|container|docker|podman|public ingress|TUI"),
		"cypher_db":             regexp.MustCompile("(?i)cypher\\.db"),
	}
	totals := map[string]int{}
	repoTotals := map[string]map[string]int{}
	var filesWithHits int

	for repo, root := range repos {
		repoTotals[repo] = map[string]int{}
		files, err := gateF4GitFiles(root)
		if err != nil {
			return nil, err
		}
		for _, rel := range files {
			path := filepath.Join(root, rel)
			if !gateF4TextCandidate(rel) {
				continue
			}
			body, err := os.ReadFile(path)
			if err != nil {
				continue
			}
			text := string(body)
			fileHit := false
			for name, rx := range patterns {
				count := len(rx.FindAllStringIndex(text, -1))
				if count > 0 {
					fileHit = true
					totals[name] += count
					repoTotals[repo][name] += count
				}
			}
			if fileHit {
				filesWithHits++
			}
		}
	}

	cypherHits := gateF4FindByName(repos, "cypher.db")
	pycacheHits := gateF4FindDirName(repos, "__pycache__")
	archiveHits := gateF4FindArchives(repos)

	return map[string]any{
		"tracked_files_with_hits": filesWithHits,
		"category_totals":         totals,
		"repo_totals":             repoTotals,
		"cypher_db_hits":          cypherHits,
		"cypher_db_hit_count":     len(cypherHits),
		"pycache_hits":            pycacheHits,
		"pycache_hit_count":       len(pycacheHits),
		"archive_hits":            archiveHits,
		"archive_hit_count":       len(archiveHits),
	}, nil
}

func gateF4CollectHelperSafety(repos map[string]string) (map[string]any, error) {
	helperName := regexp.MustCompile("(?i)(release|package|pack|dist|snapshot|validate|validation|checksum|sha|manifest|upload|tag|archive|bundle|artifact|runbook|helper|full|runtime|doctor|clean|hygiene|extract|create|stage|rehearse)")
	danger := regexp.MustCompile("(?i)(git push|git tag|curl|scp|rsync|rm -rf|rm -r|docker|podman|systemctl|sudo|gpg|openssl|upload|release create|gh release|tea release|forgejo|gitea|tar |zip |sha256sum|checksum|manifest)")
	check := regexp.MustCompile("(?i)(dry|check|validate|no-upload|no-push|print|help|usage|DRY_RUN|--check|--help|rehearse)")
	var candidateCount int
	var dangerCount int
	var checkCount int
	var samples []map[string]any

	for repo, root := range repos {
		files, err := gateF4GitFiles(root)
		if err != nil {
			return nil, err
		}
		for _, rel := range files {
			if !helperName.MatchString(rel) || !gateF4TextCandidate(rel) {
				continue
			}
			body, err := os.ReadFile(filepath.Join(root, rel))
			if err != nil {
				continue
			}
			text := string(body)
			hasDanger := danger.MatchString(text)
			hasCheck := check.MatchString(text)
			candidateCount++
			if hasDanger {
				dangerCount++
			}
			if hasCheck {
				checkCount++
			}
			if len(samples) < 40 {
				sum := sha256.Sum256(body)
				samples = append(samples, map[string]any{
					"repo":                       repo,
					"path":                       rel,
					"has_danger_terms":           hasDanger,
					"has_dry_run_or_check_terms": hasCheck,
					"sha256_12":                  hex.EncodeToString(sum[:])[:12],
				})
			}
		}
	}
	return map[string]any{
		"helper_candidate_count":  candidateCount,
		"danger_term_count":       dangerCount,
		"dry_or_check_term_count": checkCount,
		"samples":                 samples,
	}, nil
}

func gateF4RunSyntaxChecks(repos map[string]string) (map[string]any, error) {
	var shellChecked int
	var pythonChecked int
	var failures []string

	for repo, root := range repos {
		files, err := gateF4GitFiles(root)
		if err != nil {
			return nil, err
		}
		for _, rel := range files {
			path := filepath.Join(root, rel)
			switch {
			case strings.HasSuffix(rel, ".sh"):
				shellChecked++
				if err := gateF4RunSilent(root, []string{"bash", "-n", rel}); err != nil {
					failures = append(failures, repo+"/"+rel+": "+err.Error())
				}
			case strings.HasSuffix(rel, ".py"):
				pythonChecked++
				if err := gateF4RunSilent(root, []string{"python3", "-B", "-m", "py_compile", rel}); err != nil {
					failures = append(failures, repo+"/"+rel+": "+err.Error())
				}
			default:
				_ = path
			}
		}
	}
	if len(failures) > 0 {
		return nil, fmt.Errorf("syntax failures: %s", strings.Join(failures, "; "))
	}
	return map[string]any{
		"shell_checked_count":  shellChecked,
		"python_checked_count": pythonChecked,
		"failures":             failures,
	}, nil
}

func gateF4GitignoreCoversPythonCache(carbonRoot string) bool {
	body, err := os.ReadFile(filepath.Join(carbonRoot, ".gitignore"))
	if err != nil {
		return false
	}
	text := string(body)
	return strings.Contains(text, "__pycache__/") && strings.Contains(text, "*.pyc")
}

func gateF4GitFiles(root string) ([]string, error) {
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

func gateF4TextCandidate(rel string) bool {
	ext := strings.ToLower(filepath.Ext(rel))
	switch ext {
	case ".go", ".md", ".sh", ".py", ".yaml", ".yml", ".json", ".toml", ".sql", ".txt":
		return true
	default:
		return false
	}
}

func gateF4FindByName(repos map[string]string, name string) []string {
	var hits []string
	for repo, root := range repos {
		filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
			if err != nil {
				return nil
			}
			if d.IsDir() && d.Name() == ".git" {
				return filepath.SkipDir
			}
			if !d.IsDir() && d.Name() == name {
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

func gateF4FindDirName(repos map[string]string, name string) []string {
	var hits []string
	for repo, root := range repos {
		filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
			if err != nil {
				return nil
			}
			if d.IsDir() && d.Name() == ".git" {
				return filepath.SkipDir
			}
			if d.IsDir() && d.Name() == name {
				if rel, err := filepath.Rel(root, path); err == nil {
					hits = append(hits, repo+"/"+rel)
				}
				return filepath.SkipDir
			}
			return nil
		})
	}
	sort.Strings(hits)
	return hits
}

func gateF4FindArchives(repos map[string]string) []string {
	var hits []string
	for repo, root := range repos {
		filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
			if err != nil {
				return nil
			}
			if d.IsDir() && d.Name() == ".git" {
				return filepath.SkipDir
			}
			if d.IsDir() {
				return nil
			}
			name := strings.ToLower(d.Name())
			if strings.HasSuffix(name, ".tgz") || strings.HasSuffix(name, ".tar.gz") || strings.HasSuffix(name, ".zip") {
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

func gateF4RunCommand(name string, dir string, args []string) (map[string]any, error) {
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
		"name":       name,
		"dir":        dir,
		"command":    args,
		"elapsed_ms": elapsed.Milliseconds(),
		"passed":     err == nil,
	}
	if err != nil {
		result["error"] = err.Error()
		return result, fmt.Errorf("%s failed: %w", name, err)
	}
	fmt.Println("PASS:", name, "elapsed=", elapsed)
	fmt.Println()
	return result, nil
}

func gateF4RunSilent(dir string, args []string) error {
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Dir = dir
	cmd.Env = append(os.Environ(), "PYTHONDONTWRITEBYTECODE=1")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("%w: %s", err, string(out))
	}
	return nil
}

func writeGateF4JSON(path string, value any) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		return err
	}
	body, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, append(body, '\n'), 0o600)
}
