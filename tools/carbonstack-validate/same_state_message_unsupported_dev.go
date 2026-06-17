package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const sameStateUnsupportedNormalMessageContentType = "carbonstack.test.unsupported-normal-message.v0"

func (r *Runner) SameStateMessageUnsupportedDev() error {
	r.PrintHeader("same-state-message-unsupported-dev")

	fmt.Println("status: dev/pre-alpha same-state unsupported normal-message validation profile")
	fmt.Println("scope: unsupported normal application-message content_type after same-state Relay join")
	fmt.Println("proof: Relay join -> normal message send -> mutate message content_type -> unsupported skip no-ack/no-drain -> restore -> correct open/ack")
	fmt.Println("boundary: live umbrella only; not full; not release-snapshot; not package-root validation; not adversarial harness; not production/security proof")
	fmt.Println("relationship: failure-path companion to same-state-integrated-dev; separate from wrong-conversation same-state-message-failure-dev")
	fmt.Println("nonclaims: not hostile-server safety, not metadata privacy, not production secure messaging, not verified identity, not mature messenger UX")
	fmt.Println()

	if err := r.CheckRequiredPaths(); err != nil {
		return err
	}
	if err := r.CheckLiveGitUmbrella("same-state-message-unsupported-dev"); err != nil {
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

	tempRoot, err := os.MkdirTemp("", "carbonstack-same-state-message-unsupported-dev-*")
	if err != nil {
		return fmt.Errorf("create same-state message-unsupported temp root: %w", err)
	}
	defer func() {
		if err := os.RemoveAll(tempRoot); err != nil {
			fmt.Printf("WARN: remove same-state message-unsupported temp root %s: %v\n", tempRoot, err)
		}
	}()

	binPath := filepath.Join(tempRoot, "carbonstack-cypher-same-state-message-unsupported")
	uniqueID := relayOpenMLSRunID()
	runID := "same-state-unsupported-" + uniqueID
	messageLabel := "unsup-msg-" + uniqueID
	messageText := "CarbonStack same-state unsupported message proof payload " + runID

	fmt.Println()
	fmt.Println("== same-state message-unsupported generated-state root ==")
	fmt.Println("temp_root:", tempRoot)
	fmt.Println("cypher_bin:", binPath)
	fmt.Println("run_id:", runID)
	fmt.Println("message_label:", messageLabel)
	fmt.Println("unsupported_content_type:", sameStateUnsupportedNormalMessageContentType)
	fmt.Println("note: this profile proves unsupported normal-message content_type does not ack or drain the inbox")

	fmt.Println()
	fmt.Println("== build temporary Cypher binary ==")
	if err := runLocalCypherCommand(r.Cypher, "go", "build", "-o", binPath, "./cmd/cypher"); err != nil {
		return err
	}

	r.ArtifactScan("pre-same-state-message-unsupported-dev")

	result, err := r.runSameStateMessageUnsupportedSubrun(binPath, tempRoot, relayOpenMLSJoinSubrun{
		Name:     "unsupported-content-type-noack",
		AckAfter: true,
		RunID:    runID,
	}, messageLabel, messageText)
	if err != nil {
		return err
	}

	r.ArtifactScan("post-same-state-message-unsupported-dev")

	fmt.Println()
	fmt.Println("same-state-message-unsupported-dev profile result:")
	fmt.Println("  PASS: Relay onboarding completed through KeyPackage -> add-member -> Welcome -> join")
	fmt.Println("  PASS: normal message sent after same-state join")
	fmt.Println("  PASS: normal application-message content_type mutation was classified as unsupported")
	fmt.Println("  PASS: unsupported normal-message open did not ack and did not drain Bob inbox")
	fmt.Println("  PASS: restoring original content_type allowed correct open/ack")
	fmt.Println("  proof_level: same-state unsupported normal-message no-ack/no-drain proof")
	fmt.Println("  relay_space_id:", result.RelaySpaceID)
	fmt.Println("  envelopes:", result.EnvelopeCount)
	fmt.Println("  envelope_acks:", result.AckCount)
	fmt.Println("  keypackage_delivery_state:", result.KeyPackageDeliveryState)
	fmt.Println("  welcome_delivery_state:", result.WelcomeDeliveryState)
	fmt.Println("  boundary: live-dev failure-path proof; not full, not release-snapshot, not adversarial harness, not production/security proof")
	fmt.Println("  relationship: same-state-integrated-dev remains positive-path; same-state-message-failure-dev remains wrong-conversation-only")
	fmt.Println("  nonclaims: not hostile-server safety, not metadata privacy, not production secure messaging, not identity verification")

	return nil
}

func (r *Runner) runSameStateMessageUnsupportedSubrun(binPath string, tempRoot string, sub relayOpenMLSJoinSubrun, messageLabel string, messageText string) (relayOpenMLSJoinSubrunResult, error) {
	fmt.Println()
	fmt.Println("========================================")
	fmt.Println("Same-state message unsupported subrun:", sub.Name)
	fmt.Println("========================================")

	sub.TempDir = filepath.Join(tempRoot, sub.Name)
	if err := os.Mkdir(sub.TempDir, 0700); err != nil {
		return relayOpenMLSJoinSubrunResult{}, fmt.Errorf("create same-state message-unsupported subrun temp dir %s: %w", sub.TempDir, err)
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
	sub.RelaySpace = "same-state-message-unsupported-dev-" + sub.RunID
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
	fmt.Println("unsupported_content_type:", sameStateUnsupportedNormalMessageContentType)

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

		if err := assertRelayOpenMLSTrustCandidateAbsent("before same-state message-unsupported relay join", sub.AliceState, sub.BobState); err != nil {
			return err
		}

		if err := r.runRelayOpenMLSSmokeScript(&sub); err != nil {
			return err
		}

		if err := r.assertSameStateBobConversationReloadable(&sub); err != nil {
			return err
		}

		if err := assertRelayOpenMLSTrustCandidateAbsent("after same-state message-unsupported relay join", sub.AliceState, sub.BobState); err != nil {
			return err
		}

		if err := r.runSameStateUnsupportedContentTypeNoAckThenCorrectOpen(&sub, messageLabel, messageText); err != nil {
			return err
		}

		if err := assertRelayOpenMLSTrustCandidateAbsent("after same-state message-unsupported normal message open", sub.AliceState, sub.BobState); err != nil {
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
		return relayOpenMLSJoinSubrunResult{}, fmt.Errorf("same-state message-unsupported envelope count = %d, want 3", result.EnvelopeCount)
	}
	if result.AckCount != 2 {
		return relayOpenMLSJoinSubrunResult{}, fmt.Errorf("same-state message-unsupported ack count = %d, want 2", result.AckCount)
	}

	fmt.Println("PASS: same-state message unsupported subrun", sub.Name)
	return result, nil
}

func (r *Runner) runSameStateUnsupportedContentTypeNoAckThenCorrectOpen(sub *relayOpenMLSJoinSubrun, messageLabel string, messageText string) error {
	fmt.Println()
	fmt.Println("== Same-state unsupported setup: normal message send ==")

	sendOutput, err := runRelayOpenMLSCommand(
		"comms message-send-dev same-state unsupported setup",
		r.Comms,
		[]string{"RUST_BACKTRACE=1"},
		"go",
		"run", "./cmd/comms",
		"message-send-dev",
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
		"status: sent",
		"recipient_device_id: " + sub.BobDeviceID,
		"conversation: " + sub.AliceConversationLabel,
		"message_label: " + messageLabel,
	} {
		if !strings.Contains(sendOutput, needle) {
			return fmt.Errorf("message-send-dev same-state unsupported setup output missing required evidence line %q", needle)
		}
	}

	appEnvelopeID, err := sameStateSQLiteScalarString(sub.DBPath, fmt.Sprintf(
		"SELECT envelope_id FROM envelopes WHERE recipient_device_id = '%s' AND content_type NOT IN ('carbonstack.mls.keypackage.v0', 'carbonstack.mls.welcome.v0') ORDER BY rowid DESC LIMIT 1;",
		sub.BobDeviceID,
	))
	if err != nil {
		return err
	}
	if appEnvelopeID == "" {
		return fmt.Errorf("could not locate normal application-message envelope for Bob")
	}

	originalContentType, err := sameStateSQLiteScalarString(sub.DBPath, fmt.Sprintf(
		"SELECT content_type FROM envelopes WHERE envelope_id = '%s';",
		appEnvelopeID,
	))
	if err != nil {
		return err
	}
	if originalContentType == "" {
		return fmt.Errorf("could not locate original content_type for envelope %s", appEnvelopeID)
	}
	if originalContentType != "carbonstack.mls.application-message.v0" {
		return fmt.Errorf("normal message content_type = %q, want carbonstack.mls.application-message.v0", originalContentType)
	}

	acksBeforeUnsupportedOpen, err := sameStateEnvelopeAckCount(sub.DBPath)
	if err != nil {
		return err
	}
	inboxBeforeUnsupportedOpen, err := sameStateDeviceInboxCount(sub.BaseURL, sub.BobDeviceID)
	if err != nil {
		return err
	}
	if inboxBeforeUnsupportedOpen != 1 {
		return fmt.Errorf("Bob inbox count before unsupported open = %d, want 1", inboxBeforeUnsupportedOpen)
	}

	fmt.Println("app_envelope_id:", appEnvelopeID)
	fmt.Println("original_content_type:", originalContentType)
	fmt.Println("unsupported_content_type:", sameStateUnsupportedNormalMessageContentType)
	fmt.Println("acks_before_unsupported_open:", acksBeforeUnsupportedOpen)
	fmt.Println("bob_inbox_count_before_unsupported_open:", inboxBeforeUnsupportedOpen)

	fmt.Println()
	fmt.Println("== Same-state unsupported proof: mutate content_type and assert no ack/no drain ==")

	if err := sameStateSQLiteExec(sub.DBPath, fmt.Sprintf(
		"UPDATE envelopes SET content_type = '%s' WHERE envelope_id = '%s';",
		sameStateUnsupportedNormalMessageContentType,
		appEnvelopeID,
	)); err != nil {
		return err
	}

	unsupportedOutput, err := runRelayOpenMLSCommand(
		"comms message-inbox-dev unsupported content_type no-ack proof",
		r.Comms,
		[]string{"RUST_BACKTRACE=1"},
		"go",
		"run", "./cmd/comms",
		"message-inbox-dev",
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
		"ack_requested: true",
		"message skipped",
		"reason: unsupported_envelope",
		"opened_envelopes: 0",
		"unsupported_envelopes: 1",
		"open_failures: 0",
		"ack_failures: 0",
	} {
		if !strings.Contains(unsupportedOutput, needle) {
			return fmt.Errorf("unsupported content_type message-inbox-dev output missing required evidence line %q", needle)
		}
	}

	for _, forbidden := range []string{
		"acked: true",
		"ack_delivery_state: acknowledged",
		"message opened",
	} {
		if strings.Contains(unsupportedOutput, forbidden) {
			return fmt.Errorf("unsupported content_type message-inbox-dev output contained forbidden evidence line %q", forbidden)
		}
	}

	acksAfterUnsupportedOpen, err := sameStateEnvelopeAckCount(sub.DBPath)
	if err != nil {
		return err
	}
	inboxAfterUnsupportedOpen, err := sameStateDeviceInboxCount(sub.BaseURL, sub.BobDeviceID)
	if err != nil {
		return err
	}

	fmt.Println("acks_after_unsupported_open:", acksAfterUnsupportedOpen)
	fmt.Println("bob_inbox_count_after_unsupported_open:", inboxAfterUnsupportedOpen)

	if acksAfterUnsupportedOpen != acksBeforeUnsupportedOpen {
		return fmt.Errorf("unsupported content_type open changed ack count: before=%d after=%d", acksBeforeUnsupportedOpen, acksAfterUnsupportedOpen)
	}
	if inboxAfterUnsupportedOpen != inboxBeforeUnsupportedOpen {
		return fmt.Errorf("unsupported content_type open changed Bob inbox count: before=%d after=%d", inboxBeforeUnsupportedOpen, inboxAfterUnsupportedOpen)
	}

	fmt.Println("PASS: unsupported content_type open did not ack and did not drain Bob inbox")

	fmt.Println()
	fmt.Println("== Same-state recovery proof: restore content_type and open/ack ==")

	if err := sameStateSQLiteExec(sub.DBPath, fmt.Sprintf(
		"UPDATE envelopes SET content_type = '%s' WHERE envelope_id = '%s';",
		originalContentType,
		appEnvelopeID,
	)); err != nil {
		return err
	}

	correctOutput, err := runRelayOpenMLSCommand(
		"comms message-inbox-dev correct conversation after unsupported content_type",
		r.Comms,
		[]string{"RUST_BACKTRACE=1"},
		"go",
		"run", "./cmd/comms",
		"message-inbox-dev",
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
		if !strings.Contains(correctOutput, needle) {
			return fmt.Errorf("correct-conversation message-inbox-dev output missing required evidence line %q", needle)
		}
	}

	acksAfterCorrectOpen, err := sameStateEnvelopeAckCount(sub.DBPath)
	if err != nil {
		return err
	}
	inboxAfterCorrectOpen, err := sameStateDeviceInboxCount(sub.BaseURL, sub.BobDeviceID)
	if err != nil {
		return err
	}

	fmt.Println("acks_after_correct_open:", acksAfterCorrectOpen)
	fmt.Println("bob_inbox_count_after_correct_open:", inboxAfterCorrectOpen)

	if acksAfterCorrectOpen != acksBeforeUnsupportedOpen+1 {
		return fmt.Errorf("correct open ack count = %d, want %d", acksAfterCorrectOpen, acksBeforeUnsupportedOpen+1)
	}
	if inboxAfterCorrectOpen != 0 {
		return fmt.Errorf("Bob inbox count after correct open/ack = %d, want 0", inboxAfterCorrectOpen)
	}

	fmt.Println("PASS: correct conversation open/ack succeeded after unsupported content_type attempt")
	return nil
}

func sameStateSQLiteScalarString(dbPath string, query string) (string, error) {
	cmd := exec.Command("sqlite3", dbPath, query)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("sqlite3 scalar query failed: %s: %w: %s", query, err, strings.TrimSpace(string(out)))
	}
	return strings.TrimSpace(string(out)), nil
}

func sameStateSQLiteExec(dbPath string, query string) error {
	cmd := exec.Command("sqlite3", dbPath, query)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("sqlite3 exec failed: %s: %w: %s", query, err, strings.TrimSpace(string(out)))
	}
	return nil
}
