package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type relayOpenMLSJoinSubrun struct {
	Name       string
	AckAfter   bool
	RunID      string
	TempDir    string
	BaseURL    string
	DBPath     string
	RelaySpace string

	AliceState string
	BobState   string

	AliceAccountID string
	BobAccountID   string
	AliceDeviceID  string
	BobDeviceID    string

	AliceSidecarLabel      string
	BobSidecarLabel        string
	AliceConversationLabel string
	BobConversationLabel   string
}

type relayOpenMLSCommsState struct {
	AccountID string `json:"account_id"`
	DeviceID  string `json:"device_id"`
}

func (r *Runner) RelayOpenMLSJoinDev() error {
	r.PrintHeader("relay-openmls-join-dev")

	fmt.Println("status: dev/pre-alpha positive-path validation profile")
	fmt.Println("scope: Relay Space OpenMLS KeyPackage -> add-member -> Welcome -> join, with no-ack and ACK_AFTER_JOIN subruns")
	fmt.Println("boundary: not local-backbone, not production messaging, not verified identity, not hostile-server safety, not metadata privacy")
	fmt.Println("trust/candidate boundary: does not mutate trust.json, trust-events.jsonl, or identity-candidates.json")
	fmt.Println("cleanup boundary: uses runner-owned temp roots and unique sidecar labels; no destructive cleanup of unknown state")

	if err := r.CheckRequiredPaths(); err != nil {
		return err
	}
	if err := r.CheckLiveGitUmbrella("relay-openmls-join-dev"); err != nil {
		return err
	}

	fmt.Println()
	fmt.Println("== Toolchains ==")
	_ = r.ReportTool("go", "version")
	_ = r.ReportTool("rustc", "--version")
	_ = r.ReportTool("cargo", "--version")
	_ = r.ReportTool("bash", "--version")
	_ = r.ReportTool("sqlite3", "--version")

	script := filepath.Join(r.Comms, "scripts", "openmls-relay-narrow-join-smoke-dev.sh")
	if info, err := os.Stat(script); err != nil || info.IsDir() {
		return fmt.Errorf("missing Relay OpenMLS narrow join smoke script: %s", script)
	}

	tempRoot, err := os.MkdirTemp("", "carbonstack-relay-openmls-join-dev-*")
	if err != nil {
		return fmt.Errorf("create relay-openmls temp root: %w", err)
	}
	defer func() {
		if err := os.RemoveAll(tempRoot); err != nil {
			fmt.Printf("WARN: remove relay-openmls temp root %s: %v\n", tempRoot, err)
		}
	}()

	binPath := filepath.Join(tempRoot, "carbonstack-cypher-relay-openmls")

	fmt.Println()
	fmt.Println("== relay-openmls generated-state root ==")
	fmt.Println("temp_root:", tempRoot)
	fmt.Println("cypher_bin:", binPath)
	fmt.Println("note: temp root is runner-owned and removed after profile completion")
	fmt.Println("note: sidecar labels are unique per run; profile refuses if matching sidecar device paths already exist")

	fmt.Println()
	fmt.Println("== build temporary Cypher binary ==")
	if err := runLocalCypherCommand(r.Cypher, "go", "build", "-o", binPath, "./cmd/cypher"); err != nil {
		return err
	}

	r.ArtifactScan("pre-relay-openmls-join-dev")

	runID := relayOpenMLSRunID()

	if err := r.runRelayOpenMLSJoinSubrun(binPath, tempRoot, relayOpenMLSJoinSubrun{
		Name:     "no-ack",
		AckAfter: false,
		RunID:    runID + "-noack",
	}); err != nil {
		return err
	}

	if err := r.runRelayOpenMLSJoinSubrun(binPath, tempRoot, relayOpenMLSJoinSubrun{
		Name:     "ack-after-join",
		AckAfter: true,
		RunID:    runID + "-ack",
	}); err != nil {
		return err
	}

	r.ArtifactScan("post-relay-openmls-join-dev")

	fmt.Println()
	fmt.Println("relay-openmls-join-dev profile result:")
	fmt.Println("  PASS: no-ack subrun completed with KeyPackage queued, Welcome queued, and zero envelope_acks")
	fmt.Println("  PASS: ACK_AFTER_JOIN subrun completed with KeyPackage queued, Welcome acknowledged, and one envelope_acks row")
	fmt.Println("  boundary: positive-path local/dev validation only")
	fmt.Println("  nonclaims: not local-backbone, not production secure messaging, not identity verification, not hostile-server safety, not metadata privacy")

	return nil
}

func (r *Runner) runRelayOpenMLSJoinSubrun(binPath string, tempRoot string, sub relayOpenMLSJoinSubrun) error {
	fmt.Println()
	fmt.Println("========================================")
	fmt.Printf("Relay OpenMLS join subrun: %s\n", sub.Name)
	fmt.Println("========================================")

	sub.TempDir = filepath.Join(tempRoot, sub.Name)
	if err := os.Mkdir(sub.TempDir, 0700); err != nil {
		return fmt.Errorf("create subrun temp dir %s: %w", sub.TempDir, err)
	}

	port, err := reserveLoopbackPort()
	if err != nil {
		return err
	}

	addr := fmt.Sprintf("127.0.0.1:%d", port)
	sub.BaseURL = "http://" + addr
	sub.DBPath = filepath.Join(sub.TempDir, "cypher.db")
	sub.AliceState = filepath.Join(sub.TempDir, "alice-state.json")
	sub.BobState = filepath.Join(sub.TempDir, "bob-state.json")
	sub.RelaySpace = "relay-openmls-join-dev-" + sub.RunID
	sub.AliceSidecarLabel = "carbonstack-" + sub.RunID + "-alice-device"
	sub.BobSidecarLabel = "carbonstack-" + sub.RunID + "-bob-device"
	sub.AliceConversationLabel = "carbonstack-" + sub.RunID + "-conversation"
	sub.BobConversationLabel = sub.AliceConversationLabel

	fmt.Println("subrun_temp_dir:", sub.TempDir)
	fmt.Println("cypher_addr:", addr)
	fmt.Println("cypher_db:", sub.DBPath)
	fmt.Println("relay_space_id:", sub.RelaySpace)
	fmt.Println("alice_sidecar_label:", sub.AliceSidecarLabel)
	fmt.Println("bob_sidecar_label:", sub.BobSidecarLabel)
	fmt.Println("conversation_label:", sub.AliceConversationLabel)

	if err := r.refuseExistingSidecarDevice(sub.AliceSidecarLabel); err != nil {
		return err
	}
	if err := r.refuseExistingSidecarDevice(sub.BobSidecarLabel); err != nil {
		return err
	}

	env := append(os.Environ(),
		"CYPHER_ADDR="+addr,
		"CYPHER_DB="+sub.DBPath,
		"CYPHER_MIGRATIONS="+filepath.Join(r.Cypher, "migrations"),
		"CYPHER_DEV_INVITE=dev-invite",
	)

	server, err := startLocalCypherServer(binPath, r.Cypher, env)
	if err != nil {
		return err
	}

	serverErr := func() error {
		if err := waitForLocalCypherHealth(sub.BaseURL + "/v0/health"); err != nil {
			return err
		}
		fmt.Println("PASS: Cypher health check")

		if err := r.setupRelayOpenMLSCommsState(&sub); err != nil {
			return err
		}
		if err := r.setupRelayOpenMLSCypherState(&sub); err != nil {
			return err
		}
		if err := r.setupRelayOpenMLSSidecarState(&sub); err != nil {
			return err
		}

		if err := assertRelayOpenMLSTrustCandidateAbsent("before "+sub.Name, sub.AliceState, sub.BobState); err != nil {
			return err
		}

		if err := r.runRelayOpenMLSSmokeScript(&sub); err != nil {
			return err
		}

		if err := assertRelayOpenMLSTrustCandidateAbsent("after "+sub.Name, sub.AliceState, sub.BobState); err != nil {
			return err
		}

		if err := assertRelayOpenMLSDBState(&sub); err != nil {
			return err
		}

		return nil
	}()

	stopErr := server.stop(sub.Name)
	if serverErr != nil {
		return serverErr
	}
	if stopErr != nil {
		return stopErr
	}

	fmt.Printf("PASS: relay-openmls subrun %s\n", sub.Name)
	return nil
}

func (r *Runner) setupRelayOpenMLSCommsState(sub *relayOpenMLSJoinSubrun) error {
	fmt.Println()
	fmt.Println("== Comms state / account / device setup ==")

	commands := [][]string{
		{"init", "--state", sub.AliceState, "--server", sub.BaseURL},
		{"init", "--state", sub.BobState, "--server", sub.BaseURL},
		{"dev-create-invite", "--state", sub.AliceState, "--server", sub.BaseURL, "--invite", sub.RunID + "-alice"},
		{"claim-invite", "--state", sub.AliceState, "--server", sub.BaseURL, "--invite", sub.RunID + "-alice", "--name", "Alice " + sub.RunID},
		{"register-device", "--state", sub.AliceState, "--label", "alice-" + sub.RunID},
		{"dev-create-invite", "--state", sub.BobState, "--server", sub.BaseURL, "--invite", sub.RunID + "-bob"},
		{"claim-invite", "--state", sub.BobState, "--server", sub.BaseURL, "--invite", sub.RunID + "-bob", "--name", "Bob " + sub.RunID},
		{"register-device", "--state", sub.BobState, "--label", "bob-" + sub.RunID},
	}

	for _, args := range commands {
		full := append([]string{"run", "./cmd/comms"}, args...)
		if _, err := runRelayOpenMLSCommand("comms "+args[0], r.Comms, nil, "go", full...); err != nil {
			return err
		}
	}

	alice, err := readRelayOpenMLSCommsState(sub.AliceState)
	if err != nil {
		return err
	}
	bob, err := readRelayOpenMLSCommsState(sub.BobState)
	if err != nil {
		return err
	}

	if alice.AccountID == "" || alice.DeviceID == "" || bob.AccountID == "" || bob.DeviceID == "" {
		return fmt.Errorf("Comms setup produced empty account/device IDs")
	}

	sub.AliceAccountID = alice.AccountID
	sub.AliceDeviceID = alice.DeviceID
	sub.BobAccountID = bob.AccountID
	sub.BobDeviceID = bob.DeviceID

	fmt.Println("PASS: Comms Alice/Bob states created and populated")
	fmt.Println("alice_device_id:", sub.AliceDeviceID)
	fmt.Println("bob_device_id:", sub.BobDeviceID)

	return nil
}

func (r *Runner) setupRelayOpenMLSCypherState(sub *relayOpenMLSJoinSubrun) error {
	fmt.Println()
	fmt.Println("== Relay Space setup ==")

	if _, err := localCypherPOST(sub.BaseURL+"/v0/relay-spaces", map[string]any{
		"relay_space_id":        sub.RelaySpace,
		"display_label":         "relay-openmls-join-dev " + sub.Name,
		"created_by_account_id": sub.AliceAccountID,
		"created_by_device_id":  sub.AliceDeviceID,
	}, http.StatusCreated); err != nil {
		return err
	}
	fmt.Println("PASS: create Relay Space")

	members := []map[string]any{
		{
			"routing_member_id": "alice-" + sub.RunID,
			"account_id":        sub.AliceAccountID,
			"device_id":         sub.AliceDeviceID,
			"display_label":     "Alice " + sub.RunID,
			"state":             "active",
		},
		{
			"routing_member_id": "bob-" + sub.RunID,
			"account_id":        sub.BobAccountID,
			"device_id":         sub.BobDeviceID,
			"display_label":     "Bob " + sub.RunID,
			"state":             "active",
		},
	}

	for _, member := range members {
		if _, err := localCypherPOST(sub.BaseURL+"/v0/relay-spaces/"+sub.RelaySpace+"/members", member, http.StatusCreated); err != nil {
			return err
		}
	}

	fmt.Println("PASS: register Relay Space routing members")
	return nil
}

func (r *Runner) setupRelayOpenMLSSidecarState(sub *relayOpenMLSJoinSubrun) error {
	fmt.Println()
	fmt.Println("== OpenMLS sidecar identity / conversation setup ==")

	commands := [][]string{
		{"openmls-identity-create-dev", "--sidecar-device-label", sub.AliceSidecarLabel},
		{"openmls-identity-create-dev", "--sidecar-device-label", sub.BobSidecarLabel},
		{"openmls-identity-status-dev", "--sidecar-device-label", sub.AliceSidecarLabel},
		{"openmls-identity-status-dev", "--sidecar-device-label", sub.BobSidecarLabel},
		{"openmls-conversation-create-dev", "--sidecar-device-label", sub.AliceSidecarLabel, "--conversation", sub.AliceConversationLabel},
		{"openmls-conversation-load-check-dev", "--sidecar-device-label", sub.AliceSidecarLabel, "--conversation", sub.AliceConversationLabel},
	}

	for _, args := range commands {
		full := append([]string{"run", "./cmd/comms"}, args...)
		if _, err := runRelayOpenMLSCommand("comms "+args[0], r.Comms, []string{"RUST_BACKTRACE=1"}, "go", full...); err != nil {
			return err
		}
	}

	fmt.Println("PASS: sidecar identities and Alice conversation created")
	return nil
}

func (r *Runner) runRelayOpenMLSSmokeScript(sub *relayOpenMLSJoinSubrun) error {
	fmt.Println()
	fmt.Println("== Relay OpenMLS smoke script ==")

	script := filepath.Join(r.Comms, "scripts", "openmls-relay-narrow-join-smoke-dev.sh")
	logPath := filepath.Join(sub.TempDir, "smoke-"+sub.Name+".log")

	ackValue := "0"
	if sub.AckAfter {
		ackValue = "1"
	}

	env := []string{
		"RUST_BACKTRACE=1",
		"COMMS_DIR=" + r.Comms,
		"CYPHER_URL=" + sub.BaseURL,
		"ALICE_STATE=" + sub.AliceState,
		"BOB_STATE=" + sub.BobState,
		"RELAY_SPACE_ID=" + sub.RelaySpace,
		"ALICE_DEVICE_ID=" + sub.AliceDeviceID,
		"BOB_DEVICE_ID=" + sub.BobDeviceID,
		"ALICE_SIDECAR_LABEL=" + sub.AliceSidecarLabel,
		"BOB_SIDECAR_LABEL=" + sub.BobSidecarLabel,
		"ALICE_CONVERSATION_LABEL=" + sub.AliceConversationLabel,
		"BOB_CONVERSATION_LABEL=" + sub.BobConversationLabel,
		"ACK_AFTER_JOIN=" + ackValue,
	}

	output, err := runRelayOpenMLSCommand("relay smoke "+sub.Name, r.Comms, env, "bash", script)
	if writeErr := os.WriteFile(logPath, []byte(output), 0600); writeErr != nil {
		return fmt.Errorf("write smoke log %s: %w", logPath, writeErr)
	}
	fmt.Println("smoke_log:", logPath)
	if err != nil {
		return err
	}

	required := []string{
		"command: openmls-relay-keypackage-submit-dev",
		"status: sent",
		"content_type: carbonstack.mls.keypackage.v0",
		"command: openmls-relay-add-member-dev",
		"status: welcome_created_and_sent",
		"sidecar_command: conversation-add-member",
		"keypackage_acked: false",
		"welcome_acked: false",
		"command: openmls-relay-join-dev",
		"status: joined",
		"sidecar_command: conversation-join",
		"joined: true",
	}

	if sub.AckAfter {
		required = append(required,
			"ack_requested: true",
			"welcome_acked: true",
			"ack_delivery_state: acknowledged",
			"acknowledged_at:",
		)
	} else {
		required = append(required,
			"ack_requested: false",
			"welcome_acked: false",
		)
	}

	for _, needle := range required {
		if !strings.Contains(output, needle) {
			return fmt.Errorf("smoke output for %s missing required evidence line %q", sub.Name, needle)
		}
	}

	fmt.Println("PASS: smoke output assertions")
	return nil
}

func (r *Runner) refuseExistingSidecarDevice(label string) error {
	path := filepath.Join(r.Comms, "internal", "protocol", "mls", "openmls-sidecar", ".carbonstack-openmls-sidecar-state", "dev", "devices", label)
	if _, err := os.Stat(path); err == nil {
		return fmt.Errorf("refuse stale sidecar device state for label %q at %s", label, path)
	} else if !os.IsNotExist(err) {
		return fmt.Errorf("inspect sidecar device state for label %q at %s: %w", label, path, err)
	}
	fmt.Println("PASS: no pre-existing sidecar device state for", label)
	return nil
}

func readRelayOpenMLSCommsState(path string) (relayOpenMLSCommsState, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return relayOpenMLSCommsState{}, err
	}
	var state relayOpenMLSCommsState
	if err := json.Unmarshal(raw, &state); err != nil {
		return relayOpenMLSCommsState{}, err
	}
	return state, nil
}

func runRelayOpenMLSCommand(name string, dir string, env []string, command string, args ...string) (string, error) {
	fmt.Println()
	fmt.Println("----------------------------------------")
	fmt.Printf("STEP: %s\n", name)
	fmt.Printf("DIR:  %s\n", dir)
	fmt.Printf("CMD:  %s %s\n", command, strings.Join(args, " "))
	if len(env) > 0 {
		fmt.Printf("ENV:  %s\n", strings.Join(env, " "))
	}
	fmt.Println("----------------------------------------")

	cmd := exec.Command(command, args...)
	cmd.Dir = dir
	cmd.Env = append(os.Environ(), env...)

	output, err := cmd.CombinedOutput()
	text := string(output)
	if text != "" {
		fmt.Print(text)
	}
	if err != nil {
		return text, fmt.Errorf("%s failed: %w", name, err)
	}
	fmt.Println("PASS:", name)
	return text, nil
}

func assertRelayOpenMLSTrustCandidateAbsent(label string, statePaths ...string) error {
	fmt.Println()
	fmt.Println("== Trust/candidate absence check:", label, "==")

	files := []string{"trust.json", "trust-events.jsonl", "identity-candidates.json"}
	for _, statePath := range statePaths {
		stateDir := filepath.Dir(statePath)
		for _, name := range files {
			path := filepath.Join(stateDir, name)
			if _, err := os.Stat(path); err == nil {
				return fmt.Errorf("%s: unexpected trust/candidate file exists: %s", label, path)
			} else if !os.IsNotExist(err) {
				return fmt.Errorf("%s: inspect trust/candidate file %s: %w", label, path, err)
			}
			fmt.Println("ABSENT:", path)
		}
	}

	fmt.Println("PASS: checked trust/candidate files absent")
	return nil
}

func assertRelayOpenMLSDBState(sub *relayOpenMLSJoinSubrun) error {
	fmt.Println()
	fmt.Println("== Relay OpenMLS DB assertions ==")

	if err := expectRelayOpenMLSDBCount(sub.DBPath, "SELECT COUNT(*) FROM accounts;", 2, "accounts"); err != nil {
		return err
	}
	if err := expectRelayOpenMLSDBCount(sub.DBPath, "SELECT COUNT(*) FROM devices;", 2, "devices"); err != nil {
		return err
	}
	if err := expectRelayOpenMLSDBCount(sub.DBPath, "SELECT COUNT(*) FROM relay_spaces;", 1, "relay_spaces"); err != nil {
		return err
	}
	if err := expectRelayOpenMLSDBCount(sub.DBPath, "SELECT COUNT(*) FROM relay_space_members;", 2, "relay_space_members"); err != nil {
		return err
	}
	if err := expectRelayOpenMLSDBCount(sub.DBPath, "SELECT COUNT(*) FROM envelopes;", 2, "envelopes"); err != nil {
		return err
	}

	wantAcks := 0
	wantWelcomeState := "queued"
	if sub.AckAfter {
		wantAcks = 1
		wantWelcomeState = "acknowledged"
	}

	if err := expectRelayOpenMLSDBCount(sub.DBPath, "SELECT COUNT(*) FROM envelope_acks;", wantAcks, "envelope_acks"); err != nil {
		return err
	}
	if err := expectRelayOpenMLSDBCount(sub.DBPath, "SELECT COUNT(*) FROM envelopes WHERE content_type = 'carbonstack.mls.keypackage.v0' AND delivery_state = 'queued';", 1, "queued KeyPackage envelopes"); err != nil {
		return err
	}
	if err := expectRelayOpenMLSDBCount(sub.DBPath, "SELECT COUNT(*) FROM envelopes WHERE content_type = 'carbonstack.mls.welcome.v0' AND delivery_state = '"+wantWelcomeState+"';", 1, "Welcome envelopes with expected state"); err != nil {
		return err
	}

	if sub.AckAfter {
		if err := expectRelayOpenMLSDBCount(sub.DBPath, "SELECT COUNT(*) FROM envelope_acks ea JOIN envelopes e ON ea.envelope_id = e.envelope_id WHERE e.content_type = 'carbonstack.mls.welcome.v0' AND e.delivery_state = 'acknowledged';", 1, "Welcome ack rows"); err != nil {
			return err
		}
	}

	fmt.Println("PASS: DB assertions for", sub.Name)
	return nil
}

func expectRelayOpenMLSDBCount(dbPath string, query string, want int, label string) error {
	gotText, err := relayOpenMLSSQLiteScalar(dbPath, query)
	if err != nil {
		return err
	}
	got, err := strconv.Atoi(strings.TrimSpace(gotText))
	if err != nil {
		return fmt.Errorf("parse sqlite count for %s from %q: %w", label, gotText, err)
	}
	if got != want {
		return fmt.Errorf("%s count = %d, want %d", label, got, want)
	}
	fmt.Printf("PASS: %s count = %d\n", label, got)
	return nil
}

func relayOpenMLSSQLiteScalar(dbPath string, query string) (string, error) {
	cmd := exec.Command("sqlite3", dbPath, query)
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("sqlite3 query failed: %w\nquery: %s\nstderr: %s", err, query, stderr.String())
	}

	return strings.TrimSpace(stdout.String()), nil
}

func relayOpenMLSRunID() string {
	raw := fmt.Sprintf("%d-%d", time.Now().UnixNano(), os.Getpid())
	raw = strings.ReplaceAll(raw, "-", "")
	if len(raw) > 18 {
		raw = raw[len(raw)-18:]
	}
	return "v0552" + raw
}
