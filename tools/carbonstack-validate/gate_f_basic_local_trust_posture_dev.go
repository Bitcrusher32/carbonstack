package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

const gateFBasicLocalTrustPostureReportSchema = "carbonstack-gate-f-basic-local-trust-posture-report/v0"

func (r *Runner) GateFBasicLocalTrustPostureDev() error {
	fmt.Println("========================================")
	fmt.Println("CarbonStack validation profile: gate-f-basic-local-trust-posture-dev")
	fmt.Println("========================================")
	fmt.Println("status: dev/pre-alpha Gate F F5 basic local trust candidate posture")
	fmt.Println("scope: local manual candidate acceptance posture, identity-domain separation, no automatic trust promotion, loud nonclaims")
	fmt.Println("boundary: not verified identity, not full trust promotion, not secure enrollment, not cryptographic binding, not package/runtime candidate, not release creation")
	fmt.Println()

	if err := r.CheckRequiredPaths(); err != nil {
		return err
	}
	if err := r.CheckLiveGitUmbrella("gate-f-basic-local-trust-posture-dev"); err != nil {
		return err
	}

	reportRoot := filepath.Join(os.TempDir(), "carbonstack-gate-f-basic-local-trust-posture-dev")
	_ = os.RemoveAll(reportRoot)
	if err := os.MkdirAll(reportRoot, 0o700); err != nil {
		return err
	}

	commsDir := r.Comms
	stateRoot := filepath.Join(reportRoot, "comms-state")
	eventRoot := filepath.Join(reportRoot, "trust-events")
	posturePath := filepath.Join(reportRoot, "basic-local-trust-posture.json")
	acceptSourceReport := posturePath

	postureOut, err := runGateF5Command(commsDir, []string{
		"go", "run", "./cmd/comms", "basic-local-trust-posture-dev",
		"--state", stateRoot,
		"--subject-label", "alice",
		"--cypher-account", "acct-alice",
		"--cypher-device", "device-alice",
		"--comms-fingerprint", "fp-alice",
		"--openmls-device-label", "alice-sidecar",
		"--openmls-keypackage-ref", "kp-alice",
		"--relay-space", "relay-alpha",
		"--report", posturePath,
	})
	if err != nil {
		return err
	}
	posture, err := readGateF5JSON(posturePath)
	if err != nil {
		return err
	}
	if posture["schema_version"] != "carbonstack-basic-local-trust-posture/v0" {
		return fmt.Errorf("unexpected posture schema: %v", posture["schema_version"])
	}
	if posture["ready_for_manual_local_acceptance"] != true {
		return fmt.Errorf("posture did not become ready for manual local acceptance")
	}
	claims := posture["claims"].(map[string]any)
	if claims["verified_identity"] != false || claims["trust_promotion"] != false || claims["automatic_trust_promotion"] != false || claims["cryptographic_binding_across_cypher_comms_openmls"] != false {
		return fmt.Errorf("posture report made forbidden identity/trust claim")
	}

	if err := expectGateF5CommandFailure(commsDir, []string{
		"go", "run", "./cmd/comms", "basic-local-trust-accept-dev",
		"--state", stateRoot,
		"--event-root", eventRoot,
		"--subject-label", "alice",
		"--cypher-account", "acct-alice",
		"--cypher-device", "device-alice",
		"--comms-fingerprint", "fp-alice",
		"--reason", "manual out-of-band comparison",
	}); err != nil {
		return err
	}

	acceptOut, err := runGateF5Command(commsDir, []string{
		"go", "run", "./cmd/comms", "basic-local-trust-accept-dev",
		"--state", stateRoot,
		"--event-root", eventRoot,
		"--subject-label", "alice",
		"--cypher-account", "acct-alice",
		"--cypher-device", "device-alice",
		"--comms-fingerprint", "fp-alice",
		"--openmls-device-label", "alice-sidecar",
		"--openmls-keypackage-ref", "kp-alice",
		"--relay-space", "relay-alpha",
		"--reason", "manual out-of-band comparison",
		"--source-report", acceptSourceReport,
		"--accept-candidate",
	})
	if err != nil {
		return err
	}
	acceptResult, err := parseGateF5JSONBytes(acceptOut)
	if err != nil {
		return err
	}
	if acceptResult["schema_version"] != "carbonstack-basic-local-trust-acceptance-command-result/v0" {
		return fmt.Errorf("unexpected acceptance result schema: %v", acceptResult["schema_version"])
	}
	eventPath, _ := acceptResult["event_path"].(string)
	if eventPath == "" {
		return fmt.Errorf("acceptance result missing event_path")
	}
	event, ok := acceptResult["event"].(map[string]any)
	if !ok {
		return fmt.Errorf("acceptance result missing event object")
	}
	if event["schema_version"] != "carbonstack-basic-local-trust-acceptance-event/v0" {
		return fmt.Errorf("unexpected acceptance event schema: %v", event["schema_version"])
	}
	eventClaims := event["claims"].(map[string]any)
	if eventClaims["verified_identity"] != false || eventClaims["trust_promotion"] != false || eventClaims["automatic_trust_promotion"] != false || eventClaims["cryptographic_binding_across_cypher_comms_openmls"] != false {
		return fmt.Errorf("acceptance event made forbidden identity/trust claim")
	}
	if _, err := os.Stat(eventPath); err != nil {
		return fmt.Errorf("acceptance event path not written: %w", err)
	}

	validatorDir := filepath.Join(r.CarbonStack, "tools", "carbonstack-validate")
	registryLookup, err := runGateF5Command(validatorDir, []string{"go", "run", ".", "--profile", "registry-lookup", "--registry-id", "runner.gate-f-basic-local-trust-posture-dev"})
	if err != nil {
		return err
	}
	nonclaims, err := runGateF5Command(validatorDir, []string{"go", "run", ".", "--profile", "registry-lookup", "--list", "--missing-nonclaims"})
	if err != nil {
		return err
	}

	report := map[string]any{
		"schema_version":   gateFBasicLocalTrustPostureReportSchema,
		"profile":          "gate-f-basic-local-trust-posture-dev",
		"created_at":       time.Now().UTC().Format(time.RFC3339),
		"gate_f_status":    "open_f1_f2_f3_f4_f5_closed_f6_not_started",
		"gate_f_f5_status": "closed",
		"gate_f_f6_status": "not_started",
		"basic_local_trust_candidate_posture_present": true,
		"manual_local_acceptance_event_written":       true,
		"posture_report":                              posturePath,
		"acceptance_event":                            eventPath,
		"posture_stdout_bytes":                        len(postureOut),
		"accept_stdout_bytes":                         len(acceptOut),
		"registry_lookup_stdout_bytes":                len(registryLookup),
		"missing_nonclaims_stdout_bytes":              len(nonclaims),
		"verified_identity_claimed":                   false,
		"trust_promotion_claimed":                     false,
		"secure_enrollment_claimed":                   false,
		"cryptographic_identity_binding_implemented":  false,
		"automatic_trust_promotion_allowed":           false,
		"relay_membership_promotes_trust":             false,
		"welcome_or_mls_join_promotes_trust":          false,
		"release_created":                             false,
		"release_uploaded":                            false,
		"package_published":                           false,
		"package_staging_executed":                    false,
		"package_runtime_candidate_implemented":       false,
		"full_runtime_dev_promoted":                   false,
		"migration_implemented":                       false,
		"repair_implemented":                          false,
		"destructive_cleanup_performed":               false,
		"state_relocation_performed":                  false,
		"service_or_systemd_started":                  false,
		"helper_install_started":                      false,
		"container_started":                           false,
		"public_ingress_started":                      false,
		"tui_started":                                 false,
		"vault_claimed":                               false,
		"backup_restore_claimed":                      false,
		"production_e2ee_claimed":                     false,
		"pq_hybrid_claimed":                           false,
		"android_claimed":                             false,
		"carbonstack_os_claimed":                      false,
		"decisions": []string{
			"F5 closes basic local manual trust candidate posture only",
			"verified identity remains a nonclaim",
			"full trust promotion remains a nonclaim",
			"Relay membership and successful Welcome or MLS join do not promote trust",
			"Cypher, Comms, and OpenMLS identity domains remain correlated by local evidence but not cryptographically unified",
			"package/runtime candidate validation moves to the next Gate F subgate",
		},
	}
	reportPath := filepath.Join(reportRoot, "gate-f-basic-local-trust-posture-report.json")
	if err := writeGateF5JSON(reportPath, report); err != nil {
		return err
	}

	fmt.Println()
	fmt.Println("gate-f-basic-local-trust-posture-dev profile result:")
	fmt.Println("  PASS: basic local trust posture report generated")
	fmt.Println("  PASS: manual local acceptance requires explicit accept-candidate flag")
	fmt.Println("  PASS: manual local acceptance event written")
	fmt.Println("  PASS: verified identity remains false")
	fmt.Println("  PASS: trust promotion remains false")
	fmt.Println("  PASS: cryptographic identity binding remains false")
	fmt.Println("  PASS: automatic trust promotion remains false")
	fmt.Println("  PASS: registry/reference/nonclaim checks passed")
	fmt.Println("  report:", reportPath)
	fmt.Println("  posture_report:", posturePath)
	fmt.Println("  acceptance_event:", eventPath)
	fmt.Println("  gate_f_status: open_f1_f2_f3_f4_f5_closed_f6_not_started")
	fmt.Println("  gate_f_f5_status: closed")
	fmt.Println("  gate_f_f6_status: not_started")
	fmt.Println("  verified_identity_claimed: false")
	fmt.Println("  trust_promotion_claimed: false")
	fmt.Println("  automatic_trust_promotion_allowed: false")
	fmt.Println("  package_runtime_candidate_implemented: false")
	fmt.Println("  full_runtime_dev_promoted: false")
	fmt.Println("  boundary: F5 closes basic local trust candidate posture only; Gate F remains open")
	fmt.Println()
	fmt.Println("VALIDATION PASSED")
	return nil
}

func runGateF5Command(dir string, args []string) ([]byte, error) {
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Dir = dir
	out, err := cmd.CombinedOutput()
	if len(out) > 0 {
		fmt.Print(string(out))
	}
	if err != nil {
		return out, fmt.Errorf("%s failed: %w", args[0], err)
	}
	return out, nil
}

func expectGateF5CommandFailure(dir string, args []string) error {
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Dir = dir
	out, err := cmd.CombinedOutput()
	if len(out) > 0 {
		fmt.Print(string(out))
	}
	if err == nil {
		return fmt.Errorf("expected command failure but command succeeded: %v", args)
	}
	fmt.Println("expected failure observed:", err)
	return nil
}

func readGateF5JSON(path string) (map[string]any, error) {
	body, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return parseGateF5JSONBytes(body)
}

func parseGateF5JSONBytes(body []byte) (map[string]any, error) {
	var out map[string]any
	if err := json.Unmarshal(body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

func writeGateF5JSON(path string, value any) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		return err
	}
	body, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, append(body, '\n'), 0o600)
}
