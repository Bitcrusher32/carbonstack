package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

type localCypherServer struct {
	cmd    *exec.Cmd
	stdout *bytes.Buffer
	stderr *bytes.Buffer
}

func (r *Runner) LocalCypher() error {
	r.PrintHeader("local-cypher")

	if err := r.CheckRequiredPaths(); err != nil {
		return err
	}

	fmt.Println()
	fmt.Println("== Toolchains ==")
	fmt.Println("go path:", r.ReportTool("go", "env", "GOROOT"))
	fmt.Println("go version:", r.ReportTool("go", "version"))
	fmt.Println("sqlite3 version:", r.ReportTool("sqlite3", "--version"))

	r.ArtifactScan("pre-local-cypher")

	tempDir, err := os.MkdirTemp("", "carbonstack-local-cypher-*")
	if err != nil {
		return fmt.Errorf("create temp dir: %w", err)
	}
	defer func() {
		if err := os.RemoveAll(tempDir); err != nil {
			fmt.Printf("WARN: remove temp dir %s: %v\n", tempDir, err)
		}
	}()

	binPath := filepath.Join(tempDir, "carbonstack-cypher-local")
	dbPath := filepath.Join(tempDir, "cypher.db")
	migrationsDir := filepath.Join(r.Cypher, "migrations")

	port, err := reserveLoopbackPort()
	if err != nil {
		return err
	}

	addr := fmt.Sprintf("127.0.0.1:%d", port)
	baseURL := "http://" + addr
	devInvite := "local-cypher-dev-invite"

	fmt.Println()
	fmt.Println("== local-cypher environment ==")
	fmt.Println("temp_dir:", tempDir)
	fmt.Println("cypher_bin:", binPath)
	fmt.Println("cypher_db:", dbPath)
	fmt.Println("cypher_migrations:", migrationsDir)
	fmt.Println("cypher_addr:", addr)

	fmt.Println()
	fmt.Println("== build temporary Cypher binary ==")
	if err := runLocalCypherCommand(r.Cypher, "go", "build", "-o", binPath, "./cmd/cypher"); err != nil {
		return err
	}

	env := append(os.Environ(),
		"CYPHER_ADDR="+addr,
		"CYPHER_DB="+dbPath,
		"CYPHER_MIGRATIONS="+migrationsDir,
		"CYPHER_DEV_INVITE="+devInvite,
	)

	fmt.Println()
	fmt.Println("== start first Cypher process ==")
	first, err := startLocalCypherServer(binPath, r.Cypher, env)
	if err != nil {
		return err
	}

	ids := map[string]string{}

	firstErr := func() error {
		if err := waitForLocalCypherHealth(baseURL + "/v0/health"); err != nil {
			return err
		}
		fmt.Println("PASS: first health check")

		alice, err := localCypherPOST(baseURL+"/v0/invites/claim", map[string]any{
			"invite_code":  devInvite,
			"display_name": "alice-local-cypher",
		}, http.StatusCreated)
		if err != nil {
			return err
		}
		ids["alice_account_id"] = stringField(alice, "account_id")
		fmt.Println("PASS: claim seeded Alice invite")

		_, err = localCypherPOST(baseURL+"/v0/dev/invites", map[string]any{
			"invite_code": "bob-local-cypher-invite",
		}, http.StatusCreated)
		if err != nil {
			return err
		}
		fmt.Println("PASS: create Bob dev invite")

		bob, err := localCypherPOST(baseURL+"/v0/invites/claim", map[string]any{
			"invite_code":  "bob-local-cypher-invite",
			"display_name": "bob-local-cypher",
		}, http.StatusCreated)
		if err != nil {
			return err
		}
		ids["bob_account_id"] = stringField(bob, "account_id")
		fmt.Println("PASS: claim Bob invite")

		aliceDevice, err := localCypherPOST(baseURL+"/v0/devices/register", map[string]any{
			"account_id":           ids["alice_account_id"],
			"device_label":         "alice-local-cypher-device",
			"public_identity_key":  "stub-alice-public-identity-key",
			"public_prekey_bundle": "stub-alice-public-prekey-bundle",
		}, http.StatusCreated)
		if err != nil {
			return err
		}
		ids["alice_device_id"] = stringField(aliceDevice, "device_id")
		fmt.Println("PASS: register Alice device")

		bobDevice, err := localCypherPOST(baseURL+"/v0/devices/register", map[string]any{
			"account_id":           ids["bob_account_id"],
			"device_label":         "bob-local-cypher-device",
			"public_identity_key":  "stub-bob-public-identity-key",
			"public_prekey_bundle": "stub-bob-public-prekey-bundle",
		}, http.StatusCreated)
		if err != nil {
			return err
		}
		ids["bob_device_id"] = stringField(bobDevice, "device_id")
		fmt.Println("PASS: register Bob device")

		aliceDevices, err := localCypherGET(baseURL+"/v0/accounts/"+ids["alice_account_id"]+"/devices", http.StatusOK)
		if err != nil {
			return err
		}
		if len(arrayField(aliceDevices, "devices")) != 1 {
			return fmt.Errorf("list Alice devices: expected 1 device, got %d", len(arrayField(aliceDevices, "devices")))
		}
		fmt.Println("PASS: list Alice devices")

		payload := []byte("v0.3.30 local-cypher runner opaque mls payload")
		payloadB64 := base64.StdEncoding.EncodeToString(payload)
		payloadSum := sha256.Sum256(payload)
		wantSHA := hex.EncodeToString(payloadSum[:])

		envelope, err := localCypherPOST(baseURL+"/v0/envelopes", map[string]any{
			"sender_device_id":    ids["alice_device_id"],
			"recipient_device_id": ids["bob_device_id"],
			"content_type":        "carbonstack.mls.application-message.v0",
			"protocol_version":    "carbonstack-openmls-sidecar-v0",
			"ciphertext_b64":      payloadB64,
		}, http.StatusCreated)
		if err != nil {
			return err
		}
		ids["envelope_id"] = stringField(envelope, "envelope_id")
		fmt.Println("PASS: submit opaque OpenMLS application-message envelope")

		inbox, err := localCypherGET(baseURL+"/v0/devices/"+ids["bob_device_id"]+"/envelopes", http.StatusOK)
		if err != nil {
			return err
		}
		envelopes := arrayField(inbox, "envelopes")
		if len(envelopes) != 1 {
			return fmt.Errorf("retrieve Bob inbox before ack: expected 1 envelope, got %d", len(envelopes))
		}

		gotEnvelope, ok := envelopes[0].(map[string]any)
		if !ok {
			return fmt.Errorf("retrieve Bob inbox before ack: first envelope has unexpected JSON shape")
		}
		if stringField(gotEnvelope, "payload_sha256") != wantSHA {
			return fmt.Errorf("payload_sha256 mismatch: got %q want %q", stringField(gotEnvelope, "payload_sha256"), wantSHA)
		}
		if int64Field(gotEnvelope, "payload_size_bytes") != int64(len(payload)) {
			return fmt.Errorf("payload_size_bytes mismatch: got %d want %d", int64Field(gotEnvelope, "payload_size_bytes"), len(payload))
		}
		if stringField(gotEnvelope, "content_type") != "carbonstack.mls.application-message.v0" {
			return fmt.Errorf("content_type mismatch: %q", stringField(gotEnvelope, "content_type"))
		}
		if stringField(gotEnvelope, "protocol_version") != "carbonstack-openmls-sidecar-v0" {
			return fmt.Errorf("protocol_version mismatch: %q", stringField(gotEnvelope, "protocol_version"))
		}
		fmt.Println("PASS: retrieve Bob inbox and verify payload metadata")

		ack, err := localCypherPOST(baseURL+"/v0/envelopes/"+ids["envelope_id"]+"/ack", map[string]any{
			"recipient_device_id": ids["bob_device_id"],
		}, http.StatusOK)
		if err != nil {
			return err
		}
		if stringField(ack, "delivery_state") != "acknowledged" {
			return fmt.Errorf("ack delivery_state = %q, want acknowledged", stringField(ack, "delivery_state"))
		}
		fmt.Println("PASS: ack envelope")

		inboxAfterAck, err := localCypherGET(baseURL+"/v0/devices/"+ids["bob_device_id"]+"/envelopes", http.StatusOK)
		if err != nil {
			return err
		}
		if len(arrayField(inboxAfterAck, "envelopes")) != 0 {
			return fmt.Errorf("retrieve Bob inbox after ack: expected empty inbox, got %d envelopes", len(arrayField(inboxAfterAck, "envelopes")))
		}
		fmt.Println("PASS: retrieve Bob inbox after ack")

		return nil
	}()

	stopErr := first.stop("first")
	if firstErr != nil {
		return firstErr
	}
	if stopErr != nil {
		return stopErr
	}

	fmt.Println()
	fmt.Println("== restart Cypher against same temp DB ==")
	second, err := startLocalCypherServer(binPath, r.Cypher, env)
	if err != nil {
		return err
	}

	secondErr := func() error {
		if err := waitForLocalCypherHealth(baseURL + "/v0/health"); err != nil {
			return err
		}
		fmt.Println("PASS: restart health check")

		aliceDevices, err := localCypherGET(baseURL+"/v0/accounts/"+ids["alice_account_id"]+"/devices", http.StatusOK)
		if err != nil {
			return err
		}
		if len(arrayField(aliceDevices, "devices")) != 1 {
			return fmt.Errorf("persisted Alice device state after restart: expected 1 device, got %d", len(arrayField(aliceDevices, "devices")))
		}
		fmt.Println("PASS: persisted Alice device state after restart")

		bobInbox, err := localCypherGET(baseURL+"/v0/devices/"+ids["bob_device_id"]+"/envelopes", http.StatusOK)
		if err != nil {
			return err
		}
		if len(arrayField(bobInbox, "envelopes")) != 0 {
			return fmt.Errorf("acked Bob inbox after restart: expected empty inbox, got %d envelopes", len(arrayField(bobInbox, "envelopes")))
		}
		fmt.Println("PASS: acked Bob inbox remains empty after restart")

		return nil
	}()

	stopErr = second.stop("second")
	if secondErr != nil {
		return secondErr
	}
	if stopErr != nil {
		return stopErr
	}

	r.ArtifactScan("post-local-cypher")

	fmt.Println()
	fmt.Println("VALIDATION PASSED")
	return nil
}

func reserveLoopbackPort() (int, error) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return 0, fmt.Errorf("reserve loopback port: %w", err)
	}
	defer listener.Close()

	tcpAddr, ok := listener.Addr().(*net.TCPAddr)
	if !ok {
		return 0, fmt.Errorf("reserve loopback port: unexpected addr type %T", listener.Addr())
	}

	return tcpAddr.Port, nil
}

func runLocalCypherCommand(dir string, command string, args ...string) error {
	fmt.Printf("DIR:  %s\n", dir)
	fmt.Printf("CMD:  %s %s\n", command, strings.Join(args, " "))

	cmd := exec.Command(command, args...)
	cmd.Dir = dir
	output, err := cmd.CombinedOutput()
	if len(output) > 0 {
		fmt.Print(string(output))
	}
	if err != nil {
		return fmt.Errorf("%s %s: %w", command, strings.Join(args, " "), err)
	}
	return nil
}

func startLocalCypherServer(binPath string, dir string, env []string) (*localCypherServer, error) {
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}

	cmd := exec.Command(binPath)
	cmd.Dir = dir
	cmd.Env = env
	cmd.Stdout = stdout
	cmd.Stderr = stderr

	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("start Cypher: %w", err)
	}

	return &localCypherServer{
		cmd:    cmd,
		stdout: stdout,
		stderr: stderr,
	}, nil
}

func (s *localCypherServer) stop(label string) error {
	if s == nil || s.cmd == nil || s.cmd.Process == nil {
		return nil
	}

	done := make(chan error, 1)

	if s.cmd.ProcessState == nil {
		if err := s.cmd.Process.Signal(os.Interrupt); err != nil {
			_ = s.cmd.Process.Kill()
		}
	}

	go func() {
		done <- s.cmd.Wait()
	}()

	select {
	case err := <-done:
		if err != nil {
			if strings.Contains(err.Error(), "interrupt") || strings.Contains(err.Error(), "killed") || strings.Contains(err.Error(), "signal") {
				fmt.Printf("INFO: %s Cypher process stopped with expected termination: %v\n", label, err)
			} else {
				return fmt.Errorf("wait for %s Cypher process: %w\nstdout:\n%s\nstderr:\n%s", label, err, s.stdout.String(), s.stderr.String())
			}
		}
		return nil

	case <-time.After(5 * time.Second):
		_ = s.cmd.Process.Kill()
		err := <-done
		if err != nil && !strings.Contains(err.Error(), "killed") && !strings.Contains(err.Error(), "signal") {
			return fmt.Errorf("kill %s Cypher process: %w\nstdout:\n%s\nstderr:\n%s", label, err, s.stdout.String(), s.stderr.String())
		}
		fmt.Printf("INFO: %s Cypher process killed after graceful stop timeout\n", label)
		return nil
	}
}

func waitForLocalCypherHealth(url string) error {
	deadline := time.Now().Add(20 * time.Second)
	var lastErr error

	for time.Now().Before(deadline) {
		resp, err := http.Get(url)
		if err == nil {
			body, readErr := io.ReadAll(resp.Body)
			closeErr := resp.Body.Close()
			if readErr != nil {
				return fmt.Errorf("read health response: %w", readErr)
			}
			if closeErr != nil {
				return fmt.Errorf("close health response: %w", closeErr)
			}
			if resp.StatusCode == http.StatusOK {
				fmt.Printf("health response: %s\n", string(body))
				return nil
			}
			lastErr = fmt.Errorf("health status %d: %s", resp.StatusCode, string(body))
		} else {
			lastErr = err
		}
		time.Sleep(250 * time.Millisecond)
	}

	return fmt.Errorf("wait for health %s: %w", url, lastErr)
}

func localCypherGET(url string, expectedStatus int) (map[string]any, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	return localCypherDo(req, expectedStatus)
}

func localCypherPOST(url string, body map[string]any, expectedStatus int) (map[string]any, error) {
	payload, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	return localCypherDo(req, expectedStatus)
}

func localCypherDo(req *http.Request, expectedStatus int) (map[string]any, error) {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != expectedStatus {
		return nil, fmt.Errorf("%s %s: expected HTTP %d, got %d: %s", req.Method, req.URL.String(), expectedStatus, resp.StatusCode, string(raw))
	}

	if len(bytes.TrimSpace(raw)) == 0 {
		return map[string]any{}, nil
	}

	var parsed map[string]any
	if err := json.Unmarshal(raw, &parsed); err != nil {
		return nil, fmt.Errorf("%s %s: decode JSON response: %w: %s", req.Method, req.URL.String(), err, string(raw))
	}

	return parsed, nil
}

func stringField(m map[string]any, key string) string {
	value, _ := m[key].(string)
	return value
}

func int64Field(m map[string]any, key string) int64 {
	switch value := m[key].(type) {
	case float64:
		return int64(value)
	case int64:
		return value
	case int:
		return int64(value)
	default:
		return 0
	}
}

func arrayField(m map[string]any, key string) []any {
	value, _ := m[key].([]any)
	return value
}
