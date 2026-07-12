package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type sameStateInboxResponse struct {
	Envelopes []json.RawMessage `json:"envelopes"`
}

func (r *Runner) SameStateIntegratedDev() error {
	r.PrintHeader("same-state-integrated-dev")

	fmt.Println("status: dev/pre-alpha same-state integrated proof profile")
	fmt.Println("scope: Relay onboarding plus normal message send/open/ack in one coherent temp universe")
	fmt.Println("proof: KeyPackage -> add-member -> Welcome -> join -> message-send-dev -> message-inbox-dev --ack")
	fmt.Println("boundary: live umbrella only; not full; not release-snapshot; not package-root validation; not production/security proof")
	fmt.Println("relationship: stronger than integrated-runtime-dev sequential composition; does not replace integrated-runtime-dev yet")
	fmt.Println("nonclaims: not local-backbone, not production secure messaging, not identity verification, not hostile-server safety, not metadata privacy, not mature messenger UX")
	fmt.Println()

	if err := r.CheckRequiredPaths(); err != nil {
		return err
	}
	if err := r.CheckLiveGitUmbrella("same-state-integrated-dev"); err != nil {
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

	tempRoot, err := os.MkdirTemp("", "carbonstack-same-state-integrated-dev-*")
	if err != nil {
		return fmt.Errorf("create same-state temp root: %w", err)
	}
	defer func() {
		if err := os.RemoveAll(tempRoot); err != nil {
			fmt.Printf("WARN: remove same-state temp root %s: %v\n", tempRoot, err)
		}
	}()

	binPath := filepath.Join(tempRoot, "carbonstack-cypher-same-state")
	runID := "same-state-" + relayOpenMLSRunID()
	messageLabel := "same-state-message-" + runID
	messageText := "CarbonStack same-state integrated proof payload " + runID

	fmt.Println()
	fmt.Println("== same-state generated-state root ==")
	fmt.Println("temp_root:", tempRoot)
	fmt.Println("cypher_bin:", binPath)
	fmt.Println("run_id:", runID)
	fmt.Println("message_label:", messageLabel)
	fmt.Println("note: one temp universe owns Cypher DB, Alice/Bob Comms states, Relay Space, sidecar labels, and normal message proof")

	fmt.Println()
	fmt.Println("== build temporary Cypher binary ==")
	if err := runLocalCypherCommand(r.Cypher, "go", "build", "-o", binPath, "./cmd/cypher"); err != nil {
		return err
	}

	r.ArtifactScan("pre-same-state-integrated-dev")

	result, err := r.runSameStateIntegratedSubrun(binPath, tempRoot, relayOpenMLSJoinSubrun{
		Name:     "same-state",
		AckAfter: true,
		RunID:    runID,
	}, messageLabel, messageText)
	if err != nil {
		return err
	}

	r.ArtifactScan("post-same-state-integrated-dev")

	fmt.Println()
	fmt.Println("same-state-integrated-dev profile result:")
	fmt.Println("  PASS: Relay onboarding completed through KeyPackage -> add-member -> Welcome -> join")
	fmt.Println("  PASS: Bob joined same conversation with group_reloadable evidence")
	fmt.Println("  PASS: normal message sent with message-send-dev using same Alice state/device/conversation")
	fmt.Println("  PASS: normal message opened and acked with message-inbox-dev using same Bob state/device/conversation")
	fmt.Println("  PASS: Bob inbox empty after Welcome ack plus normal message ack")
	fmt.Println("  proof_level: Level 4 same conversation normal-message proof")
	fmt.Println("  relay_space_id:", result.RelaySpaceID)
	fmt.Println("  envelopes:", result.EnvelopeCount)
	fmt.Println("  envelope_acks:", result.AckCount)
	fmt.Println("  keypackage_delivery_state:", result.KeyPackageDeliveryState)
	fmt.Println("  welcome_delivery_state:", result.WelcomeDeliveryState)
	fmt.Println("  boundary: live-dev positive-path same-state proof; not full, not release-snapshot, not package-root validation, not production/security proof")
	fmt.Println("  relationship: integrated-runtime-dev remains the sequential composition profile")
	fmt.Println("  nonclaims: not local-backbone, not production secure messaging, not identity verification, not hostile-server safety, not metadata privacy")

	return nil
}

func (r *Runner) runSameStateIntegratedSubrun(binPath string, tempRoot string, sub relayOpenMLSJoinSubrun, messageLabel string, messageText string) (relayOpenMLSJoinSubrunResult, error) {
	fmt.Println()
	fmt.Println("========================================")
	fmt.Println("Same-state integrated subrun:", sub.Name)
	fmt.Println("========================================")

	sub.TempDir = filepath.Join(tempRoot, sub.Name)
	if err := os.Mkdir(sub.TempDir, 0700); err != nil {
		return relayOpenMLSJoinSubrunResult{}, fmt.Errorf("create same-state subrun temp dir %s: %w", sub.TempDir, err)
	}

	port, err := reserveLoopbackPort()
	if err != nil {
		return relayOpenMLSJoinSubrunResult{}, err
	}

	addr := fmt.Sprintf("127.0.0.1:%d", port)
	sub.BaseURL = "http://" + addr
	sub.DBPath = filepath.Join(sub.TempDir, "cypher.db")
	sub.AliceState = filepath.Join(sub.TempDir, "alice-state.json")
	sub.BobState = filepath.Join(sub.TempDir, "bob-state.json")
	sub.RelaySpace = "same-state-integrated-dev-" + sub.RunID
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
	fmt.Println("message_label:", messageLabel)

	if err := r.refuseExistingSidecarDevice(sub.AliceSidecarLabel); err != nil {
		return relayOpenMLSJoinSubrunResult{}, err
	}
	if err := r.refuseExistingSidecarDevice(sub.BobSidecarLabel); err != nil {
		return relayOpenMLSJoinSubrunResult{}, err
	}

	env := append(os.Environ(),
		"CYPHER_ADDR="+addr,
		"CYPHER_DB="+sub.DBPath,
		"CYPHER_MIGRATIONS="+filepath.Join(r.Cypher, "migrations"),
		"CYPHER_DEV_INVITE=dev-invite",
	)

	server, err := startLocalCypherServer(binPath, r.Cypher, env)
	if err != nil {
		return relayOpenMLSJoinSubrunResult{}, err
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

		if err := assertRelayOpenMLSTrustCandidateAbsent("before same-state relay join", sub.AliceState, sub.BobState); err != nil {
			return err
		}

		if err := r.runRelayOpenMLSSmokeScript(&sub); err != nil {
			return err
		}

		if err := r.assertSameStateBobConversationReloadable(&sub); err != nil {
			return err
		}

		if err := assertRelayOpenMLSTrustCandidateAbsent("after same-state relay join", sub.AliceState, sub.BobState); err != nil {
			return err
		}

		if err := r.runSameStateMessageSendOpenAck(&sub, messageLabel, messageText); err != nil {
			return err
		}

		if err := assertRelayOpenMLSTrustCandidateAbsent("after same-state normal message open", sub.AliceState, sub.BobState); err != nil {
			return err
		}

		if err := assertSameStateIntegratedDBState(&sub); err != nil {
			return err
		}

		return nil
	}()

	stopErr := server.stop(sub.Name)
	if serverErr != nil {
		return relayOpenMLSJoinSubrunResult{}, serverErr
	}
	if stopErr != nil {
		return relayOpenMLSJoinSubrunResult{}, stopErr
	}

	result, err := collectRelayOpenMLSSubrunResult(&sub)
	if err != nil {
		return relayOpenMLSJoinSubrunResult{}, err
	}

	if result.EnvelopeCount != 3 {
		return relayOpenMLSJoinSubrunResult{}, fmt.Errorf("same-state envelope count = %d, want 3", result.EnvelopeCount)
	}
	if result.AckCount != 2 {
		return relayOpenMLSJoinSubrunResult{}, fmt.Errorf("same-state ack count = %d, want 2", result.AckCount)
	}

	fmt.Println("PASS: same-state integrated subrun", sub.Name)
	return result, nil
}

func (r *Runner) assertSameStateBobConversationReloadable(sub *relayOpenMLSJoinSubrun) error {
	fmt.Println()
	fmt.Println("== Same-state Bob conversation reload check ==")

	output, err := runRelayOpenMLSCommand(
		"comms openmls-conversation-load-check-dev Bob after Relay join",
		r.Comms,
		[]string{"RUST_BACKTRACE=1"},
		"go",
		"run", "./cmd/comms",
		"openmls-conversation-load-check-dev",
		"--sidecar-device-label", sub.BobSidecarLabel,
		"--conversation", sub.BobConversationLabel,
	)
	if err != nil {
		return err
	}
	if !strings.Contains(output, "group_reloadable: true") {
		return fmt.Errorf("Bob conversation load-check missing group_reloadable: true")
	}
	fmt.Println("PASS: Bob conversation reloadable after Relay join")
	return nil
}

func (r *Runner) runSameStateMessageSendOpenAck(sub *relayOpenMLSJoinSubrun, messageLabel string, messageText string) error {
	fmt.Println()
	fmt.Println("== Same-state normal message send ==")

	sendOutput, err := runRelayOpenMLSCommand(
		"comms message-send-dev same-state",
		r.Comms,
		[]string{"RUST_BACKTRACE=1"},
		"go",
		"run", "./cmd/comms",
		"message-send-dev",
		"--relay-space", sub.RelaySpace,
		"--state", sub.AliceState,
		"--to-device", sub.BobDeviceID,
		"--sidecar-device-label", sub.AliceSidecarLabel,
		"--conversation", sub.AliceConversationLabel,
		"--message-label", messageLabel,
		"--message", messageText,
	)
	if err != nil {
		return err
	}

	for _, needle := range []string{
		"message sent",
		"command: message-send-dev",
		"implementation_path: openmls-send-dev",
		"backend: OpenMLS sidecar + Cypher Relay Space-scoped application-message envelope",
		"status: sent",
		"recipient_device_id: " + sub.BobDeviceID,
		"conversation: " + sub.AliceConversationLabel,
		"message_label: " + messageLabel,
		"payload_sha256:",
	} {
		if !strings.Contains(sendOutput, needle) {
			return fmt.Errorf("message-send-dev same-state output missing required evidence line %q", needle)
		}
	}

	fmt.Println()
	fmt.Println("== Same-state normal message open and ack ==")

	inboxOutput, err := runRelayOpenMLSCommand(
		"comms message-inbox-dev same-state",
		r.Comms,
		[]string{"RUST_BACKTRACE=1"},
		"go",
		"run", "./cmd/comms",
		"message-inbox-dev",
		"--relay-space", sub.RelaySpace,
		"--state", sub.BobState,
		"--sidecar-device-label", sub.BobSidecarLabel,
		"--conversation", sub.BobConversationLabel,
		"--message-label", messageLabel,
		"--ack",
	)
	if err != nil {
		return err
	}

	for _, needle := range []string{
		"message inbox",
		"command: message-inbox-dev",
		"implementation_path: openmls-inbox-dev",
		"backend: OpenMLS sidecar + Cypher Relay Space-scoped application-message envelope",
		"device_id: " + sub.BobDeviceID,
		"ack_requested: true",
		"message opened",
		"ack_delivery_state: acknowledged",
		"from_device: " + sub.AliceDeviceID,
		"conversation: " + sub.BobConversationLabel,
		"message_label: " + messageLabel,
		"plaintext_utf8: " + messageText,
		"acked: true",
		"opened_envelopes: 1",
		"unsupported_envelopes: 0",
		"open_failures: 0",
		"ack_failures: 0",
	} {
		if !strings.Contains(inboxOutput, needle) {
			return fmt.Errorf("message-inbox-dev same-state output missing required evidence line %q", needle)
		}
	}

	count, err := sameStateDeviceInboxCount(sub.BaseURL, sub.RelaySpace, sub.BobDeviceID)
	if err != nil {
		return err
	}
	if count != 0 {
		return fmt.Errorf("Bob inbox count after same-state ack = %d, want 0", count)
	}
	fmt.Println("PASS: Bob inbox empty after normal message ack")
	return nil
}

func sameStateDeviceInboxCount(baseURL string, relaySpace string, deviceID string) (int, error) {
	resp, err := http.Get(baseURL + "/v0/relay-spaces/" + relaySpace + "/devices/" + deviceID + "/envelopes")
	if err != nil {
		return 0, fmt.Errorf("fetch device inbox after same-state ack: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("fetch device inbox after same-state ack status = %d, want %d", resp.StatusCode, http.StatusOK)
	}

	var body sameStateInboxResponse
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return 0, fmt.Errorf("decode device inbox after same-state ack: %w", err)
	}
	return len(body.Envelopes), nil
}

func assertSameStateIntegratedDBState(sub *relayOpenMLSJoinSubrun) error {
	fmt.Println()
	fmt.Println("== Same-state integrated DB assertions ==")

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
	if err := expectRelayOpenMLSDBCount(sub.DBPath, "SELECT COUNT(*) FROM envelopes;", 3, "envelopes"); err != nil {
		return err
	}
	if err := expectRelayOpenMLSDBCount(sub.DBPath, "SELECT COUNT(*) FROM envelope_acks;", 2, "envelope_acks"); err != nil {
		return err
	}
	if err := expectRelayOpenMLSDBCount(sub.DBPath, "SELECT COUNT(*) FROM envelopes WHERE content_type = 'carbonstack.mls.keypackage.v0' AND delivery_state = 'queued';", 1, "queued KeyPackage envelopes"); err != nil {
		return err
	}
	if err := expectRelayOpenMLSDBCount(sub.DBPath, "SELECT COUNT(*) FROM envelopes WHERE content_type = 'carbonstack.mls.welcome.v0' AND delivery_state = 'acknowledged';", 1, "acknowledged Welcome envelopes"); err != nil {
		return err
	}
	if err := expectRelayOpenMLSDBCount(sub.DBPath, "SELECT COUNT(*) FROM envelope_acks ea JOIN envelopes e ON ea.envelope_id = e.envelope_id WHERE e.content_type = 'carbonstack.mls.welcome.v0' AND e.delivery_state = 'acknowledged';", 1, "Welcome ack rows"); err != nil {
		return err
	}

	fmt.Println("PASS: DB assertions for same-state integrated proof")
	return nil
}
