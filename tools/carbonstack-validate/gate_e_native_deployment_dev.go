package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"
)

const gateENativeDeploymentReportSchema = "carbonstack-gate-e-native-deployment-report/v0"
const gateEDeploymentContextCandidateSchema = "carbonstack-gate-e-deployment-context-candidate/v0"

type gateECypherProcess struct {
	cmd *exec.Cmd
	log *os.File
}

func (r *Runner) GateENativeDeploymentDev() error {
	fmt.Println("========================================")
	fmt.Println("CarbonStack validation profile: gate-e-native-deployment-dev")
	fmt.Println("========================================")
	fmt.Println("status: dev/pre-alpha Gate E E1 manual private native deployment profile")
	fmt.Println("scope: explicit-env, loopback-only, deployment-root fixture; Cypher start/stop/restart; Gate C state-policy inspection")
	fmt.Println("boundary: not service/systemd, not helper install, not container, not public ingress, not TUI, not full-runtime-dev, not verified identity, not trust promotion")
	fmt.Println()

	if err := r.CheckRequiredPaths(); err != nil {
		return err
	}
	if err := r.CheckLiveGitUmbrella("gate-e-native-deployment-dev"); err != nil {
		return err
	}

	validatorDir := filepath.Join(r.CarbonStack, "tools", "carbonstack-validate")
	reportRoot := filepath.Join(os.TempDir(), "carbonstack-gate-e-native-deployment-dev")
	_ = os.RemoveAll(reportRoot)

	deployRoot := filepath.Join(reportRoot, "deploy-root")
	binRoot := filepath.Join(deployRoot, "bin")
	configRoot := filepath.Join(deployRoot, "config")
	commsStateRoot := filepath.Join(deployRoot, "state", "comms")
	cypherStateRoot := filepath.Join(deployRoot, "state", "cypher")
	sidecarStateRoot := filepath.Join(deployRoot, "state", "sidecar")
	logRoot := filepath.Join(deployRoot, "logs")
	evidenceRoot := filepath.Join(deployRoot, "evidence")
	tempRoot := filepath.Join(deployRoot, "tmp")

	for _, dir := range []string{binRoot, configRoot, commsStateRoot, cypherStateRoot, sidecarStateRoot, logRoot, evidenceRoot, tempRoot} {
		if err := os.MkdirAll(dir, 0o700); err != nil {
			return err
		}
	}

	cypherBin := filepath.Join(binRoot, "carbonstack-cypher-gate-e")
	cypherDB := filepath.Join(cypherStateRoot, "cypher.db")
	cypherMigrations := filepath.Join(r.Cypher, "migrations")
	cypherInvite := "gate-e-native-deployment-dev-invite"
	cypherAddr, cypherURL, err := gateELoopbackAddr()
	if err != nil {
		return err
	}
	restartAddr, restartURL, err := gateELoopbackAddr()
	if err != nil {
		return err
	}

	commsStatePath := filepath.Join(commsStateRoot, "state.json")
	envExamplePath := filepath.Join(configRoot, "gate-e.env.example")
	contextPath := filepath.Join(configRoot, "gate-e.deployment-context.json")
	runtimeReportPath := filepath.Join(evidenceRoot, "gate-e-native-deployment-report.json")
	c1InventoryReport := filepath.Join(evidenceRoot, "c1-inventory.json")
	c2C1CompatReport := filepath.Join(evidenceRoot, "c2-c1-compat.json")
	c2C3CompatReport := filepath.Join(evidenceRoot, "c2-c3-compat.json")
	c2C4CompatReport := filepath.Join(evidenceRoot, "c2-c4-compat.json")
	c3PathPolicyReport := filepath.Join(evidenceRoot, "c3-path-policy.json")
	c4WritePolicyReport := filepath.Join(evidenceRoot, "c4-write-policy.json")

	envExample := fmt.Sprintf(`# CarbonStack Gate E E1 deployment context candidate.
# Generated evidence only; not a production config and not a service install.

CARBONSTACK_DEPLOY_ROOT=%s
CARBONSTACK_COMMS_STATE=%s
CARBONSTACK_COMMS_STATE_ROOT=%s
CARBONSTACK_EVIDENCE_ROOT=%s
CARBONSTACK_LOG_ROOT=%s
CARBONSTACK_TEMP_ROOT=%s
CARBONSTACK_SIDECAR_STATE_ROOT=%s

CYPHER_ADDR=%s
CYPHER_DB=%s
CYPHER_MIGRATIONS=%s
CYPHER_DEV_INVITE=%s
`, deployRoot, commsStatePath, commsStateRoot, evidenceRoot, logRoot, tempRoot, sidecarStateRoot, cypherAddr, cypherDB, cypherMigrations, cypherInvite)

	if err := os.WriteFile(envExamplePath, []byte(envExample), 0o600); err != nil {
		return err
	}
	if err := writeGateEJSON(commsStatePath, map[string]any{
		"server_url": cypherURL,
		"account_id": "acct-gate-e-native-deployment-dev",
		"device_id":  "dev-gate-e-native-deployment-dev",
	}); err != nil {
		return err
	}

	contextDoc := map[string]any{
		"schema_version":                     gateEDeploymentContextCandidateSchema,
		"profile_name":                       "gate-e-native-deployment-dev",
		"deployment_model":                   "manual_private_explicit_env_first",
		"deploy_root":                        deployRoot,
		"bin_root":                           binRoot,
		"config_root":                        configRoot,
		"env_example_path":                   envExamplePath,
		"comms_state_path":                   commsStatePath,
		"comms_state_root":                   commsStateRoot,
		"cypher_state_root":                  cypherStateRoot,
		"cypher_db":                          cypherDB,
		"cypher_migrations":                  cypherMigrations,
		"cypher_addr":                        cypherAddr,
		"cypher_url":                         cypherURL,
		"sidecar_state_root":                 sidecarStateRoot,
		"log_root":                           logRoot,
		"evidence_root":                      evidenceRoot,
		"temp_root":                          tempRoot,
		"manual_private_deployment_first":    true,
		"semi_persistent_deployment_started": false,
		"service_or_systemd_started":         false,
		"helper_install_started":             false,
		"container_started":                  false,
		"public_ingress_started":             false,
		"tui_started":                        false,
		"full_runtime_dev_promoted":          false,
		"verified_identity_claimed":          false,
		"trust_promotion_claimed":            false,
		"vault_claimed":                      false,
		"backup_restore_claimed":             false,
		"production_e2ee_claimed":            false,
		"pq_hybrid_claimed":                  false,
		"android_claimed":                    false,
		"carbonstack_os_claimed":             false,
		"cypher_help_probe_used":             false,
		"cypher_inspection_posture":          "server_entrypoint; inspect by source/config docs or intentional explicit-env supervised startup only",
	}
	if err := writeGateEJSON(contextPath, contextDoc); err != nil {
		return err
	}

	r.ArtifactScan("pre-gate-e-native-deployment-dev")

	if err := r.RunStep(Step{
		Name:    "Build Cypher into explicit Gate E deployment bin root",
		Dir:     r.Cypher,
		Command: "go",
		Args:    []string{"build", "-o", cypherBin, "./cmd/cypher"},
	}); err != nil {
		return err
	}

	firstLog := filepath.Join(logRoot, "cypher-first-start.log")
	firstProc, err := startGateECypher(r.Cypher, cypherBin, cypherAddr, cypherDB, cypherMigrations, cypherInvite, firstLog)
	if err != nil {
		return err
	}
	if err := waitForGateEHealth(cypherURL+"/v0/health", 30*time.Second); err != nil {
		_ = stopGateECypher(firstProc)
		return err
	}
	if err := stopGateECypher(firstProc); err != nil {
		return err
	}
	fmt.Println("PASS: Cypher first explicit-env health check")
	fmt.Println("cypher_first_log:", firstLog)

	if _, err := os.Stat(cypherDB); err != nil {
		return fmt.Errorf("cypher DB missing after first explicit-env start: %w", err)
	}
	fmt.Println("PASS: Cypher DB created in explicit deployment state root")
	fmt.Println("cypher_db:", cypherDB)

	if _, err := exec.LookPath("sqlite3"); err == nil {
		if err := r.RunStep(Step{
			Name:    "Inspect Cypher DB tables",
			Dir:     r.Cypher,
			Command: "sqlite3",
			Args:    []string{cypherDB, ".tables"},
		}); err != nil {
			return err
		}
		if err := r.RunStep(Step{
			Name:    "Inspect Cypher schema_migrations",
			Dir:     r.Cypher,
			Command: "sqlite3",
			Args:    []string{cypherDB, "select migration_name, sha256 from schema_migrations order by migration_name;"},
		}); err != nil {
			return err
		}
	} else {
		fmt.Println("sqlite3 not available; DB migration query skipped")
	}

	restartLog := filepath.Join(logRoot, "cypher-restart.log")
	restartProc, err := startGateECypher(r.Cypher, cypherBin, restartAddr, cypherDB, cypherMigrations, cypherInvite, restartLog)
	if err != nil {
		return err
	}
	if err := waitForGateEHealth(restartURL+"/v0/health", 30*time.Second); err != nil {
		_ = stopGateECypher(restartProc)
		return err
	}
	if err := stopGateECypher(restartProc); err != nil {
		return err
	}
	fmt.Println("PASS: Cypher restart explicit-env health check against same DB")
	fmt.Println("cypher_restart_log:", restartLog)

	steps := []Step{
		{
			Name:    "C1 inventory over Gate E deployment roots",
			Dir:     r.Comms,
			Command: "go",
			Args: []string{
				"run", "./cmd/comms", "state-substrate-inventory-dev",
				"--state", commsStatePath,
				"--state-root", commsStateRoot,
				"--sidecar-dir", filepath.Join(r.Comms, "internal", "protocol", "mls", "openmls-sidecar"),
				"--cypher-db", cypherDB,
				"--evidence-root", evidenceRoot,
				"--output", c1InventoryReport,
			},
		},
		{
			Name:    "C3 path policy over Gate E deployment roots",
			Dir:     r.Comms,
			Command: "go",
			Args: []string{
				"run", "./cmd/comms", "state-path-policy-dev",
				"--state", commsStatePath,
				"--state-root", commsStateRoot,
				"--sidecar-dir", filepath.Join(r.Comms, "internal", "protocol", "mls", "openmls-sidecar"),
				"--cypher-db", cypherDB,
				"--validator-temp-root", tempRoot,
				"--evidence-root", evidenceRoot,
				"--output", c3PathPolicyReport,
			},
		},
		{
			Name:    "C4 write policy over Gate E deployment roots",
			Dir:     r.Comms,
			Command: "go",
			Args: []string{
				"run", "./cmd/comms", "state-write-policy-dev",
				"--state-root", commsStateRoot,
				"--sidecar-dir", filepath.Join(r.Comms, "internal", "protocol", "mls", "openmls-sidecar"),
				"--cypher-db", cypherDB,
				"--validator-temp-root", tempRoot,
				"--evidence-root", evidenceRoot,
				"--output", c4WritePolicyReport,
			},
		},
		{
			Name:    "C2 validates Gate E C1 inventory report",
			Dir:     r.Comms,
			Command: "go",
			Args: []string{
				"run", "./cmd/comms", "state-schema-compat-dev",
				"--kind", "state-substrate-inventory",
				"--path", c1InventoryReport,
				"--output", c2C1CompatReport,
			},
		},
		{
			Name:    "C2 validates Gate E C3 path-policy report",
			Dir:     r.Comms,
			Command: "go",
			Args: []string{
				"run", "./cmd/comms", "state-schema-compat-dev",
				"--kind", "path-policy-report",
				"--path", c3PathPolicyReport,
				"--output", c2C3CompatReport,
			},
		},
		{
			Name:    "C2 validates Gate E C4 write-policy report",
			Dir:     r.Comms,
			Command: "go",
			Args: []string{
				"run", "./cmd/comms", "state-schema-compat-dev",
				"--kind", "write-policy-report",
				"--path", c4WritePolicyReport,
				"--output", c2C4CompatReport,
			},
		},
		{
			Name:    "Generated command reference current",
			Dir:     r.CarbonStack,
			Command: "python3",
			Args:    []string{"tools/registry/render-command-reference.py", "--check"},
		},
		{
			Name:    "Registry missing nonclaims remains zero",
			Dir:     validatorDir,
			Command: "go",
			Args:    []string{"run", ".", "--profile", "registry-lookup", "--list", "--missing-nonclaims"},
		},
		{
			Name:    "Gate D runtime aggregate authority present",
			Dir:     validatorDir,
			Command: "go",
			Args:    []string{"run", ".", "--profile", "registry-lookup", "--registry-id", "runner.gate-d-runtime-aggregate-dev"},
		},
		{
			Name:    "Gate E native deployment authority present",
			Dir:     validatorDir,
			Command: "go",
			Args:    []string{"run", ".", "--profile", "registry-lookup", "--registry-id", "runner.gate-e-native-deployment-dev"},
		},
	}

	for _, step := range steps {
		if err := r.RunStep(step); err != nil {
			return err
		}
	}

	report := map[string]any{
		"schema_version":                             gateENativeDeploymentReportSchema,
		"profile":                                    "gate-e-native-deployment-dev",
		"created_at":                                 time.Now().UTC().Format(time.RFC3339),
		"gate_d_status":                              "closed",
		"gate_e_status":                              "open_e1_closed_e2_blocked",
		"gate_e_e1_status":                           "closed",
		"gate_e_e2_status":                           "not_started",
		"deployment_model":                           "manual_private_explicit_env_first",
		"deployment_context_schema":                  gateEDeploymentContextCandidateSchema,
		"deployment_context_path":                    contextPath,
		"deploy_root":                                deployRoot,
		"bin_root":                                   binRoot,
		"config_root":                                configRoot,
		"env_example_path":                           envExamplePath,
		"comms_state_path":                           commsStatePath,
		"comms_state_root":                           commsStateRoot,
		"cypher_db":                                  cypherDB,
		"cypher_migrations":                          cypherMigrations,
		"cypher_first_addr":                          cypherAddr,
		"cypher_restart_addr":                        restartAddr,
		"log_root":                                   logRoot,
		"evidence_root":                              evidenceRoot,
		"temp_root":                                  tempRoot,
		"manual_private_deployment_first":            true,
		"semi_persistent_deployment_started":         false,
		"service_or_systemd_started":                 false,
		"helper_install_started":                     false,
		"container_started":                          false,
		"public_ingress_started":                     false,
		"tui_started":                                false,
		"full_runtime_dev_promoted":                  false,
		"full_runtime_dev_registry_name_added":       false,
		"verified_identity_claimed":                  false,
		"trust_promotion_claimed":                    false,
		"vault_claimed":                              false,
		"backup_restore_claimed":                     false,
		"production_e2ee_claimed":                    false,
		"pq_hybrid_claimed":                          false,
		"android_claimed":                            false,
		"carbonstack_os_claimed":                     false,
		"cypher_help_probe_used":                     false,
		"cypher_server_entrypoint_hazard_classified": true,
		"closure_evidence": map[string]any{
			"explicit_env_file_generated":                  true,
			"cypher_built_into_deployment_bin_root":        true,
			"cypher_started_with_explicit_loopback_addr":   true,
			"cypher_started_with_explicit_db_path":         true,
			"cypher_started_with_explicit_migrations_path": true,
			"cypher_health_check_passed":                   true,
			"cypher_restart_against_same_db_passed":        true,
			"gate_c_state_policy_reports_passed":           true,
			"gate_c_report_compatibility_passed":           true,
			"registry_reference_current":                   true,
			"missing_nonclaims_zero":                       true,
			"gate_d_authority_present_without_promotion":   true,
		},
		"reports": map[string]any{
			"c1_inventory":    c1InventoryReport,
			"c2_c1_compat":    c2C1CompatReport,
			"c2_c3_compat":    c2C3CompatReport,
			"c2_c4_compat":    c2C4CompatReport,
			"c3_path_policy":  c3PathPolicyReport,
			"c4_write_policy": c4WritePolicyReport,
		},
	}
	if err := writeGateEJSON(runtimeReportPath, report); err != nil {
		return err
	}

	fmt.Println()
	fmt.Println("gate-e-native-deployment-dev profile result:")
	fmt.Println("  PASS: explicit deployment-root fixture generated")
	fmt.Println("  PASS: env example records explicit manual private deployment context")
	fmt.Println("  PASS: Cypher built into deployment bin root")
	fmt.Println("  PASS: Cypher first start used explicit loopback addr, DB path, migrations path, and invite")
	fmt.Println("  PASS: Cypher health check passed")
	fmt.Println("  PASS: Cypher stopped and restarted against the same explicit DB")
	fmt.Println("  PASS: C1/C2/C3/C4 state-policy surfaces passed over deployment roots")
	fmt.Println("  PASS: registry/reference checks passed")
	fmt.Println("  PASS: missing nonclaims remain zero")
	fmt.Println("  gate_d_status: closed")
	fmt.Println("  gate_e_status: open_e1_closed_e2_blocked")
	fmt.Println("  gate_e_e1_status: closed")
	fmt.Println("  gate_e_e2_status: not_started")
	fmt.Println("  runtime_report:", runtimeReportPath)
	fmt.Println("  deployment_context:", contextPath)
	fmt.Println("  cypher_help_probe_used: false")
	fmt.Println("  manual_private_deployment_first: true")
	fmt.Println("  semi_persistent_deployment_started: false")
	fmt.Println("  service_or_systemd_started: false")
	fmt.Println("  helper_install_started: false")
	fmt.Println("  container_started: false")
	fmt.Println("  public_ingress_started: false")
	fmt.Println("  tui_started: false")
	fmt.Println("  full_runtime_dev_promoted: false")
	fmt.Println("  verified_identity_claimed: false")
	fmt.Println("  trust_promotion_claimed: false")
	fmt.Println("  vault_claimed: false")
	fmt.Println("  backup_restore_claimed: false")
	fmt.Println("  production_e2ee_claimed: false")
	fmt.Println("  pq_hybrid_claimed: false")
	fmt.Println("  android_claimed: false")
	fmt.Println("  carbonstack_os_claimed: false")
	fmt.Println("  boundary: Gate E E1 manual private native deployment profile only; further Gate E subgates require breakpoint first")

	r.ArtifactScan("post-gate-e-native-deployment-dev")
	if r.CleanGenerated {
		_ = os.RemoveAll(reportRoot)
		r.CleanGeneratedArtifacts()
	}
	fmt.Println()
	fmt.Println("VALIDATION PASSED")
	return nil
}

func gateELoopbackAddr() (string, string, error) {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return "", "", err
	}
	defer ln.Close()
	addr := ln.Addr().(*net.TCPAddr)
	value := "127.0.0.1:" + strconv.Itoa(addr.Port)
	return value, "http://" + value, nil
}

func startGateECypher(dir, bin, addr, dbPath, migrationsDir, invite, logPath string) (*gateECypherProcess, error) {
	if err := os.MkdirAll(filepath.Dir(logPath), 0o700); err != nil {
		return nil, err
	}
	logFile, err := os.Create(logPath)
	if err != nil {
		return nil, err
	}
	cmd := exec.Command(bin)
	cmd.Dir = dir
	cmd.Env = append(os.Environ(),
		"CYPHER_ADDR="+addr,
		"CYPHER_DB="+dbPath,
		"CYPHER_MIGRATIONS="+migrationsDir,
		"CYPHER_DEV_INVITE="+invite,
	)
	cmd.Stdout = logFile
	cmd.Stderr = logFile
	if err := cmd.Start(); err != nil {
		_ = logFile.Close()
		return nil, err
	}
	return &gateECypherProcess{cmd: cmd, log: logFile}, nil
}

func stopGateECypher(proc *gateECypherProcess) error {
	if proc == nil || proc.cmd == nil {
		return nil
	}
	if proc.cmd.Process != nil {
		_ = proc.cmd.Process.Kill()
	}
	err := proc.cmd.Wait()
	if proc.log != nil {
		_ = proc.log.Close()
	}
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok && !exitErr.Success() {
			return nil
		}
	}
	return nil
}

func waitForGateEHealth(url string, timeout time.Duration) error {
	client := &http.Client{Timeout: 2 * time.Second}
	deadline := time.Now().Add(timeout)
	var lastErr error

	for time.Now().Before(deadline) {
		resp, err := client.Get(url)
		if err != nil {
			lastErr = err
			time.Sleep(500 * time.Millisecond)
			continue
		}
		body, readErr := io.ReadAll(resp.Body)
		_ = resp.Body.Close()
		if readErr != nil {
			lastErr = readErr
			time.Sleep(500 * time.Millisecond)
			continue
		}
		if resp.StatusCode != http.StatusOK {
			lastErr = fmt.Errorf("health status=%d body=%s", resp.StatusCode, string(body))
			time.Sleep(500 * time.Millisecond)
			continue
		}
		var parsed map[string]any
		if err := json.Unmarshal(body, &parsed); err != nil {
			lastErr = fmt.Errorf("health invalid json: %w body=%s", err, string(body))
			time.Sleep(500 * time.Millisecond)
			continue
		}
		if parsed["status"] == "ok" {
			fmt.Println("health_status:", resp.StatusCode)
			fmt.Println("health_body:", string(bytes.TrimSpace(body)))
			return nil
		}
		lastErr = fmt.Errorf("health status field not ok: %s", string(body))
		time.Sleep(500 * time.Millisecond)
	}

	return fmt.Errorf("timed out waiting for Gate E Cypher health at %s: %v", url, lastErr)
}

func writeGateEJSON(path string, value any) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		return err
	}
	body, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, append(body, '\n'), 0o600)
}
