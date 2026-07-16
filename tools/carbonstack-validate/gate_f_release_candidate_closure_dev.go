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

const gateFReleaseCandidateClosureReportSchema = "carbonstack-gate-f-release-candidate-closure-report/v0"
const gateFReleaseCandidateHandoffSchema = "carbonstack-v0.8.0-manual-release-handoff/v0"

type gateF7MatrixItem struct {
	Gate       string
	RegistryID string
	Claim      string
}

func (r *Runner) GateFReleaseCandidateClosureDev() error {
	fmt.Println("========================================")
	fmt.Println("CarbonStack validation profile: gate-f-release-candidate-closure-dev")
	fmt.Println("========================================")
	fmt.Println("status: dev/pre-alpha Gate F F7 release-candidate closure matrix and manual release handoff")
	fmt.Println("scope: v0.8.0 release-candidate claim matrix, nonclaim matrix, validation command list, manual handoff checklist, and release notes scaffold")
	fmt.Println("boundary: not release creation, not release upload, not package publication, not package staging execution, not full-runtime-dev promotion")
	fmt.Println()

	if err := r.CheckRequiredPaths(); err != nil {
		return err
	}
	if err := r.CheckLiveGitUmbrella("gate-f-release-candidate-closure-dev"); err != nil {
		return err
	}

	reportRoot := filepath.Join(os.TempDir(), "carbonstack-gate-f-release-candidate-closure-dev")
	_ = os.RemoveAll(reportRoot)
	if err := os.MkdirAll(reportRoot, 0o700); err != nil {
		return err
	}

	heads, err := gateF7RepoHeads(map[string]string{
		"carbonstack":        r.CarbonStack,
		"carbonstack-comms":  r.Comms,
		"carbonstack-cypher": r.Cypher,
	})
	if err != nil {
		return err
	}
	osRepo := filepath.Join(filepath.Dir(r.CarbonStack), "carbonstack-os")
	if st, err := os.Stat(osRepo); err == nil && st.IsDir() {
		if head, err := gateF7GitValue(osRepo, "rev-parse", "HEAD"); err == nil {
			heads["carbonstack-os"] = head
		}
	}

	if err := gateF7RequiredPaths([]string{
		filepath.Join(r.CarbonStack, "docs", "251-v0.7.10-gate-b-closure-v0.md"),
		filepath.Join(r.CarbonStack, "docs", "261-v0.7.16-gate-c-closure-v0.md"),
		filepath.Join(r.CarbonStack, "docs", "263-v0.7.17-gate-d-closure-v0.md"),
		filepath.Join(r.CarbonStack, "docs", "269-v0.7.19-gate-e-closure-v0.md"),
		filepath.Join(r.CarbonStack, "docs", "271-v0.7.20-gate-f-f1-closure-v0.md"),
		filepath.Join(r.CarbonStack, "docs", "276-v0.7.20-gate-f-f2-closure-v0.md"),
		filepath.Join(r.CarbonStack, "docs", "278-v0.7.21-gate-f-f3-closure-v0.md"),
		filepath.Join(r.CarbonStack, "docs", "280-v0.7.22-gate-f-f4-closure-v0.md"),
		filepath.Join(r.CarbonStack, "docs", "282-v0.7.23-gate-f-f5-closure-v0.md"),
		filepath.Join(r.CarbonStack, "docs", "284-v0.7.24-gate-f-f6-closure-v0.md"),
		filepath.Join(r.CarbonStack, "registry", "commands.v0.yaml"),
		filepath.Join(r.CarbonStack, "registry", "COMMAND_REFERENCE.v0.md"),
		filepath.Join(r.CarbonStack, "registry", "COMMAND_BOUNDARY_TABLE.v0.md"),
	}); err != nil {
		return err
	}

	matrix := gateF7RequiredMatrix()
	registryText, err := os.ReadFile(filepath.Join(r.CarbonStack, "registry", "commands.v0.yaml"))
	if err != nil {
		return err
	}
	referenceText, err := os.ReadFile(filepath.Join(r.CarbonStack, "registry", "COMMAND_REFERENCE.v0.md"))
	if err != nil {
		return err
	}
	boundaryText, err := os.ReadFile(filepath.Join(r.CarbonStack, "registry", "COMMAND_BOUNDARY_TABLE.v0.md"))
	if err != nil {
		return err
	}
	for _, item := range matrix {
		if !gateF7RegistryIDPresent(string(registryText), item.RegistryID) {
			return fmt.Errorf("required closure registry ID missing from registry: %s", item.RegistryID)
		}
		if !strings.Contains(string(referenceText), item.RegistryID) {
			return fmt.Errorf("required closure registry ID missing from command reference: %s", item.RegistryID)
		}
		if !strings.Contains(string(boundaryText), item.RegistryID) {
			return fmt.Errorf("required closure registry ID missing from boundary table: %s", item.RegistryID)
		}
	}

	validatorDir := filepath.Join(r.CarbonStack, "tools", "carbonstack-validate")
	f6Out, err := gateF7RunCommand("Gate F F6 package runtime candidate dependency", validatorDir, []string{"go", "run", ".", "--profile", "gate-f-package-runtime-candidate-dev", "--compact-summary"})
	if err != nil {
		return err
	}
	refOut, err := gateF7RunCommand("generated command reference current", r.CarbonStack, []string{"python3", "-B", "tools/registry/render-command-reference.py", "--check"})
	if err != nil {
		return err
	}
	lookupBytes := 0
	for _, item := range matrix {
		out, err := gateF7RunCommand("registry lookup "+item.RegistryID, validatorDir, []string{"go", "run", ".", "--profile", "registry-lookup", "--registry-id", item.RegistryID})
		if err != nil {
			return err
		}
		lookupBytes += len(out)
	}
	missingOut, err := gateF7RunCommand("missing nonclaims scan", validatorDir, []string{"go", "run", ".", "--profile", "registry-lookup", "--list", "--missing-nonclaims"})
	if err != nil {
		return err
	}

	handoff := gateF7ManualReleaseHandoff(heads)
	handoffPath := filepath.Join(reportRoot, "v0.8.0-manual-release-handoff.json")
	if err := gateF7WriteJSON(handoffPath, handoff); err != nil {
		return err
	}

	scaffoldPath := filepath.Join(reportRoot, "v0.8.0-release-notes-scaffold.md")
	if err := gateF7WriteBytes(scaffoldPath, []byte(gateF7ReleaseNotesScaffold(heads))); err != nil {
		return err
	}

	checklistPath := filepath.Join(reportRoot, "v0.8.0-manual-release-checklist.md")
	if err := gateF7WriteBytes(checklistPath, []byte(gateF7ManualReleaseChecklist())); err != nil {
		return err
	}

	artifacts, err := gateF7ArchiveHits(r.CarbonStack, r.Comms, r.Cypher, osRepo)
	if err != nil {
		return err
	}
	if len(artifacts) > 0 {
		return fmt.Errorf("repo-root public package/archive artifacts detected: %v", artifacts)
	}
	noPromotion, err := gateF7NoFullRuntimePromotion(r.CarbonStack)
	if err != nil {
		return err
	}

	reportPath := filepath.Join(reportRoot, "gate-f-release-candidate-closure-report.json")
	report := map[string]any{
		"schema_version":                         gateFReleaseCandidateClosureReportSchema,
		"profile":                                "gate-f-release-candidate-closure-dev",
		"created_at":                             time.Now().UTC().Format(time.RFC3339),
		"gate_f_status":                          "closed_v0_8_0_release_candidate_handoff_ready",
		"gate_f_f7_status":                       "closed",
		"v0_8_0_release_candidate_handoff_ready": true,
		"manual_release_handoff_path":            handoffPath,
		"release_notes_scaffold_path":            scaffoldPath,
		"manual_release_checklist_path":          checklistPath,
		"source_heads":                           heads,
		"closure_matrix":                         matrix,
		"claims":                                 gateF7Claims(),
		"nonclaims":                              gateF7Nonclaims(),
		"validation_command_list":                gateF7ValidationCommandList(),
		"known_deferrals":                        gateF7KnownDeferrals(),
		"operator_supplied_release_inputs": []string{
			"previous public release markdown",
			"previous release visual assets",
			"previous release attached-file component list",
			"target Gitea release link or release creation surface",
			"operator confirmation to create/upload public release artifacts manually",
		},
		"f6_dependency_stdout_bytes":                 len(f6Out),
		"reference_check_stdout_bytes":               len(refOut),
		"registry_lookup_stdout_bytes":               lookupBytes,
		"missing_nonclaims_stdout_bytes":             len(missingOut),
		"repo_root_archive_artifacts":                artifacts,
		"full_runtime_dev_promoted":                  !noPromotion,
		"release_created":                            false,
		"release_uploaded":                           false,
		"package_published":                          false,
		"package_staging_executed":                   false,
		"public_package_artifact_created":            false,
		"service_or_systemd_started":                 false,
		"helper_install_started":                     false,
		"container_started":                          false,
		"public_ingress_started":                     false,
		"tui_started":                                false,
		"migration_implemented":                      false,
		"repair_implemented":                         false,
		"destructive_cleanup_performed":              false,
		"state_relocation_performed":                 false,
		"verified_identity_claimed":                  false,
		"trust_promotion_claimed":                    false,
		"secure_enrollment_claimed":                  false,
		"cryptographic_identity_binding_implemented": false,
		"automatic_trust_promotion_allowed":          false,
		"vault_claimed":                              false,
		"backup_restore_claimed":                     false,
		"production_e2ee_claimed":                    false,
		"pq_hybrid_claimed":                          false,
		"android_claimed":                            false,
		"carbonstack_os_claimed":                     false,
		"next_action":                                "private v0.7.26 LogDoc and Breakpoint, then operator-controlled v0.8.0 release preparation using prior release assets",
	}
	if err := gateF7WriteJSON(reportPath, report); err != nil {
		return err
	}

	fmt.Println()
	fmt.Println("gate-f-release-candidate-closure-dev profile result:")
	fmt.Println("  PASS: v0.8.0 release-candidate closure matrix generated")
	fmt.Println("  PASS: Gate B/C/D/E and Gate F F1-F6 authority surfaces present")
	fmt.Println("  PASS: Gate F F6 package/runtime candidate dependency passed")
	fmt.Println("  PASS: generated command reference current")
	fmt.Println("  PASS: registry lookup and missing-nonclaims checks passed")
	fmt.Println("  PASS: manual release handoff JSON written")
	fmt.Println("  PASS: release notes scaffold written")
	fmt.Println("  PASS: manual release checklist written")
	fmt.Println("  PASS: no public package/archive artifacts in repo roots")
	fmt.Println("  PASS: full-runtime-dev remains unpromoted")
	fmt.Println("  report:", reportPath)
	fmt.Println("  manual_release_handoff:", handoffPath)
	fmt.Println("  release_notes_scaffold:", scaffoldPath)
	fmt.Println("  manual_release_checklist:", checklistPath)
	fmt.Println("  gate_f_status: closed_v0_8_0_release_candidate_handoff_ready")
	fmt.Println("  gate_f_f7_status: closed")
	fmt.Println("  v0_8_0_release_candidate_handoff_ready: true")
	fmt.Println("  release_created: false")
	fmt.Println("  release_uploaded: false")
	fmt.Println("  package_published: false")
	fmt.Println("  package_staging_executed: false")
	fmt.Println("  public_package_artifact_created: false")
	fmt.Println("  full_runtime_dev_promoted: false")
	fmt.Println("  boundary: F7 closes release-candidate matrix and manual handoff only; public release remains operator-controlled")
	fmt.Println()
	fmt.Println("VALIDATION PASSED")
	return nil
}

func gateF7RequiredMatrix() []gateF7MatrixItem {
	return []gateF7MatrixItem{
		{"Gate B", "runner.gate-b-relay-lifecycle-closure-dev", "Relay lifecycle closure"},
		{"Gate C", "runner.gate-c-state-substrate-closure-dev", "State substrate closure"},
		{"Gate D", "runner.gate-d-runtime-aggregate-dev", "Runtime aggregate closure"},
		{"Gate E", "runner.gate-e-native-deployment-closure-dev", "Manual-private native deployment closure"},
		{"Gate F F1", "runner.gate-f-release-package-surface-dev", "Release/package/runtime surface classification"},
		{"Gate F F2", "runner.gate-f-operator-docs-runbook-dev", "Operator docs/runbook closure"},
		{"Gate F F3", "runner.gate-f-compat-rollback-observability-dev", "Compatibility and rollback observability closure"},
		{"Gate F F4", "runner.gate-f-code-health-source-hygiene-dev", "Code-health/source-hygiene closure"},
		{"Gate F F5", "runner.gate-f-basic-local-trust-posture-dev", "Basic local trust candidate posture closure"},
		{"Gate F F6", "runner.gate-f-package-runtime-candidate-dev", "Package/runtime candidate validation closure"},
		{"Release package", "runner.full-validate-release", "Release-package validation ladder authority"},
		{"Release snapshot", "runner.release-snapshot", "Release package layout/checksum validation authority"},
		{"Cypher config", "cypher.config-inspection", "Terminating explicit-env Cypher config inspection"},
	}
}

func gateF7Claims() map[string]bool {
	return map[string]bool{
		"gate_b_closed":                               true,
		"gate_c_closed":                               true,
		"gate_d_closed":                               true,
		"gate_e_closed":                               true,
		"gate_f_f1_closed":                            true,
		"gate_f_f2_closed":                            true,
		"gate_f_f3_closed":                            true,
		"gate_f_f4_closed":                            true,
		"gate_f_f5_closed":                            true,
		"gate_f_f6_closed":                            true,
		"gate_f_f7_closed":                            true,
		"v0_8_0_release_candidate_handoff_ready":      true,
		"manual_private_lifecycle_documented":         true,
		"package_runtime_candidate_validation_exists": true,
		"basic_local_trust_posture_exists":            true,
	}
}

func gateF7Nonclaims() map[string]bool {
	return map[string]bool{
		"release_created":                            false,
		"release_uploaded":                           false,
		"package_published":                          false,
		"package_staging_executed":                   false,
		"public_package_artifact_created":            false,
		"full_runtime_dev_promoted":                  false,
		"service_or_systemd_started":                 false,
		"helper_install_started":                     false,
		"container_started":                          false,
		"public_ingress_started":                     false,
		"tui_started":                                false,
		"migration_implemented":                      false,
		"repair_implemented":                         false,
		"destructive_cleanup_performed":              false,
		"state_relocation_performed":                 false,
		"verified_identity_claimed":                  false,
		"trust_promotion_claimed":                    false,
		"secure_enrollment_claimed":                  false,
		"cryptographic_identity_binding_implemented": false,
		"automatic_trust_promotion_allowed":          false,
		"vault_claimed":                              false,
		"backup_restore_claimed":                     false,
		"production_e2ee_claimed":                    false,
		"pq_hybrid_claimed":                          false,
		"android_claimed":                            false,
		"carbonstack_os_claimed":                     false,
	}
}

func gateF7ValidationCommandList() []string {
	return []string{
		"go run . --profile gate-f-release-candidate-closure-dev --compact-summary",
		"go run . --profile gate-f-package-runtime-candidate-dev --compact-summary",
		"go run . --profile gate-f-basic-local-trust-posture-dev --compact-summary",
		"go run . --profile registry-lookup --list --missing-nonclaims",
		"env PYTHONDONTWRITEBYTECODE=1 python3 -B tools/registry/render-command-reference.py --check",
	}
}

func gateF7KnownDeferrals() []string {
	return []string{
		"public release creation and upload remain operator-controlled",
		"package publication and package staging execution remain unperformed",
		"full-runtime-dev promotion remains deferred",
		"service/systemd/helper install remains deferred",
		"container/public ingress/TUI remain deferred",
		"migration, repair, destructive cleanup, and state relocation remain deferred",
		"verified identity, full trust promotion, secure enrollment, and cryptographic identity binding remain deferred",
		"vault, backup/restore, production E2EE, PQ/hybrid, Android, and CarbonStackOS remain deferred",
	}
}

func gateF7ManualReleaseHandoff(heads map[string]string) map[string]any {
	return map[string]any{
		"schema_version":          gateFReleaseCandidateHandoffSchema,
		"created_at":              time.Now().UTC().Format(time.RFC3339),
		"release":                 "v0.8.0",
		"release_title_candidate": "CarbonStack v0.8.0 Operational Spine Maturation Pre-Release",
		"source_heads":            heads,
		"operator_inputs_required": []string{
			"prior release markdown scaffold",
			"prior release visual assets",
			"prior release attached-file component list",
			"Gitea release target/link",
			"operator confirmation to create and upload public release assets",
		},
		"manual_steps": []string{
			"review private v0.7.26 LogDoc and Breakpoint after F7",
			"prepare release notes from scaffold",
			"verify source heads and command reference entries",
			"run final validation command list",
			"manually create the v0.8.0 pre-release on Gitea",
			"manually attach expected artifacts and visuals",
			"capture public release URL",
			"generate post-release PRIME continuity artifact",
			"deprecate the v0.8.0 EVERGREEN roadmap",
			"start v0.9.x adversarial roadmap Q/A",
		},
		"nonclaims": gateF7Nonclaims(),
	}
}

func gateF7ReleaseNotesScaffold(heads map[string]string) string {
	var b strings.Builder
	b.WriteString("# CarbonStack v0.8.0 Operational Spine Maturation Pre-Release\n\n")
	b.WriteString("Status: release notes scaffold only. Public release remains operator-controlled.\n\n")
	b.WriteString("## Source heads\n\n")
	keys := make([]string, 0, len(heads))
	for key := range heads {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		b.WriteString("- ")
		b.WriteString(key)
		b.WriteString(": ")
		b.WriteString(heads[key])
		b.WriteString("\n")
	}
	b.WriteString("\n## Claims\n\n")
	b.WriteString("- Gate B Relay lifecycle closed.\n")
	b.WriteString("- Gate C state substrate closed.\n")
	b.WriteString("- Gate D runtime aggregate closed.\n")
	b.WriteString("- Gate E manual-private native deployment closed.\n")
	b.WriteString("- Gate F release/package/runtime, operator docs, compatibility, source hygiene, basic local trust posture, package/runtime candidate validation, and release-candidate handoff closed.\n")
	b.WriteString("\n## Nonclaims\n\n")
	for key, value := range gateF7Nonclaims() {
		b.WriteString("- ")
		b.WriteString(key)
		b.WriteString(": ")
		if value {
			b.WriteString("true")
		} else {
			b.WriteString("false")
		}
		b.WriteString("\n")
	}
	b.WriteString("\n## Manual release steps\n\n")
	b.WriteString("Use prior release markdown, visual assets, attached-file component list, and operator-controlled Gitea release flow.\n")
	return b.String()
}

func gateF7ManualReleaseChecklist() string {
	return "# CarbonStack v0.8.0 Manual Release Checklist\n\n" +
		"Status: checklist scaffold only. Not release creation.\n\n" +
		"1. Review final F7 log.\n" +
		"2. Generate and accept the post-F7 private LogDoc and Breakpoint.\n" +
		"3. Prepare release notes from the scaffold.\n" +
		"4. Verify source heads.\n" +
		"5. Verify command reference entries.\n" +
		"6. Run final validation commands.\n" +
		"7. Confirm nonclaims.\n" +
		"8. Manually create the v0.8.0 pre-release on Gitea.\n" +
		"9. Manually attach expected artifacts and visuals.\n" +
		"10. Capture public release URL.\n" +
		"11. Generate post-release PRIME continuity artifact.\n" +
		"12. Deprecate the v0.8.0 EVERGREEN roadmap.\n" +
		"13. Start v0.9.x adversarial roadmap Q/A.\n"
}

func gateF7RepoHeads(repos map[string]string) (map[string]string, error) {
	out := map[string]string{}
	for name, dir := range repos {
		head, err := gateF7GitValue(dir, "rev-parse", "HEAD")
		if err != nil {
			return nil, err
		}
		out[name] = head
	}
	return out, nil
}

func gateF7GitValue(dir string, args ...string) (string, error) {
	cmd := exec.Command("git", args...)
	cmd.Dir = dir
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}

func gateF7RequiredPaths(paths []string) error {
	for _, path := range paths {
		if _, err := os.Stat(path); err != nil {
			return fmt.Errorf("required path missing: %s: %w", path, err)
		}
	}
	return nil
}

func gateF7RegistryIDPresent(registry string, id string) bool {
	rx := regexp.MustCompile("(?m)^\\s*-\\s+id:\\s*" + regexp.QuoteMeta(id) + "\\s*$")
	return rx.MatchString(registry)
}

func gateF7RunCommand(name string, dir string, args []string) ([]byte, error) {
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

func gateF7NoFullRuntimePromotion(carbonRoot string) (bool, error) {
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

func gateF7ArchiveHits(paths ...string) ([]string, error) {
	var hits []string
	for _, root := range paths {
		if root == "" {
			continue
		}
		if st, err := os.Stat(root); err != nil || !st.IsDir() {
			continue
		}
		err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
			if err != nil || d.IsDir() {
				return nil
			}
			rel, _ := filepath.Rel(root, path)
			if strings.HasPrefix(rel, ".git"+string(os.PathSeparator)) {
				return nil
			}
			depth := strings.Count(rel, string(os.PathSeparator))
			if depth > 1 {
				return nil
			}
			name := strings.ToLower(d.Name())
			if strings.HasSuffix(name, ".tgz") || strings.HasSuffix(name, ".tar.gz") || strings.HasSuffix(name, ".zip") {
				hits = append(hits, filepath.Base(root)+"/"+rel)
			}
			return nil
		})
		if err != nil {
			return nil, err
		}
	}
	sort.Strings(hits)
	return hits, nil
}

func gateF7WriteJSON(path string, value any) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		return err
	}
	body, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		return err
	}
	return gateF7WriteBytes(path, append(body, '\n'))
}

func gateF7WriteBytes(path string, body []byte) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		return err
	}
	return os.WriteFile(path, body, 0o600)
}
