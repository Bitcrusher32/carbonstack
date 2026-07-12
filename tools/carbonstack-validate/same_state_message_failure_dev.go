package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func (r *Runner) SameStateMessageFailureDev() error {
	r.PrintHeader("same-state-message-failure-dev")

	fmt.Println("status: dev/pre-alpha same-state normal-message failure validation profile")
	fmt.Println("scope: wrong-conversation message-inbox-dev --ack failure after same-state Relay join")
	fmt.Println("proof: Relay join -> normal message send -> wrong-conversation open failure no-ack/no-drain -> correct open/ack still succeeds")
	fmt.Println("boundary: live umbrella only; not full; not release-snapshot; not package-root validation; not adversarial harness; not production/security proof")
	fmt.Println("relationship: failure-path companion to same-state-integrated-dev; does not mutate the positive-path profile")
	fmt.Println("nonclaims: not hostile-server safety, not metadata privacy, not production secure messaging, not verified identity, not mature messenger UX")
	fmt.Println()

	if err := r.CheckRequiredPaths(); err != nil {
		return err
	}
	if err := r.CheckLiveGitUmbrella("same-state-message-failure-dev"); err != nil {
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

	tempRoot, err := os.MkdirTemp("", "carbonstack-same-state-message-failure-dev-*")
	if err != nil {
		return fmt.Errorf("create same-state message-failure temp root: %w", err)
	}
	defer func() {
		if err := os.RemoveAll(tempRoot); err != nil {
			fmt.Printf("WARN: remove same-state message-failure temp root %s: %v\n", tempRoot, err)
		}
	}()

	binPath := filepath.Join(tempRoot, "carbonstack-cypher-same-state-message-failure")
	uniqueID := relayOpenMLSRunID()
	runID := "same-state-failure-" + uniqueID
	messageLabel := "fail-msg-" + uniqueID
	messageText := "CarbonStack same-state message failure proof payload " + runID

	fmt.Println()
	fmt.Println("== same-state message-failure generated-state root ==")
	fmt.Println("temp_root:", tempRoot)
	fmt.Println("cypher_bin:", binPath)
	fmt.Println("run_id:", runID)
	fmt.Println("message_label:", messageLabel)
	fmt.Println("note: this profile proves wrong-conversation message-open failure does not ack or drain the inbox")

	fmt.Println()
	fmt.Println("== build temporary Cypher binary ==")
	if err := runLocalCypherCommand(r.Cypher, "go", "build", "-o", binPath, "./cmd/cypher"); err != nil {
		return err
	}

	r.ArtifactScan("pre-same-state-message-failure-dev")

	result, err := r.runSameStateMessageFailureSubrun(binPath, tempRoot, relayOpenMLSJoinSubrun{
		Name:     "wrong-conversation-noack",
		AckAfter: true,
		RunID:    runID,
	}, messageLabel, messageText)
	if err != nil {
		return err
	}

	r.ArtifactScan("post-same-state-message-failure-dev")

	fmt.Println()
	fmt.Println("same-state-message-failure-dev profile result:")
	fmt.Println("  PASS: Relay onboarding completed through KeyPackage -> add-member -> Welcome -> join")
	fmt.Println("  PASS: normal message sent after same-state join")
	fmt.Println("  PASS: wrong-conversation message-inbox-dev --ack reported message-open failure")
	fmt.Println("  PASS: wrong-conversation open did not ack and did not drain Bob inbox")
	fmt.Println("  PASS: correct conversation open/ack still succeeded after failed attempt")
	fmt.Println("  proof_level: same-state normal-message failure no-ack/no-drain proof")
	fmt.Println("  relay_space_id:", result.RelaySpaceID)
	fmt.Println("  envelopes:", result.EnvelopeCount)
	fmt.Println("  envelope_acks:", result.AckCount)
	fmt.Println("  keypackage_delivery_state:", result.KeyPackageDeliveryState)
	fmt.Println("  welcome_delivery_state:", result.WelcomeDeliveryState)
	fmt.Println("  boundary: live-dev failure-path proof; not full, not release-snapshot, not adversarial harness, not production/security proof")
	fmt.Println("  relationship: same-state-integrated-dev remains the positive-path same-conversation proof profile")
	fmt.Println("  nonclaims: not hostile-server safety, not metadata privacy, not production secure messaging, not identity verification")

	return nil
}

func (r *Runner) runSameStateMessageFailureSubrun(binPath string, tempRoot string, sub relayOpenMLSJoinSubrun, messageLabel string, messageText string) (relayOpenMLSJoinSubrunResult, error) {
	fmt.Println()
	fmt.Println("========================================")
	fmt.Println("Same-state message failure subrun:", sub.Name)
	fmt.Println("========================================")

	sub.TempDir = filepath.Join(tempRoot, sub.Name)
	if err := os.Mkdir(sub.TempDir, 0700); err != nil {
		return relayOpenMLSJoinSubrunResult{}, fmt.Errorf("create same-state message-failure subrun temp dir %s: %w", sub.TempDir, err)
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
	sub.RelaySpace = "same-state-message-failure-dev-" + sub.RunID
	sub.AliceSidecarLabel = "carbonstack-" + sub.RunID + "-alice-device"
	sub.BobSidecarLabel = "carbonstack-" + sub.RunID + "-bob-device"
	sub.AliceConversationLabel = "carbonstack-" + sub.RunID + "-conversation"
	sub.BobConversationLabel = sub.AliceConversationLabel

	wrongConversationLabel := "carbonstack-" + sub.RunID + "-wrong-conversation"

	fmt.Println("subrun_temp_dir:", sub.TempDir)
	fmt.Println("cypher_addr:", addr)
	fmt.Println("cypher_db:", sub.DBPath)
	fmt.Println("relay_space_id:", sub.RelaySpace)
	fmt.Println("alice_sidecar_label:", sub.AliceSidecarLabel)
	fmt.Println("bob_sidecar_label:", sub.BobSidecarLabel)
	fmt.Println("conversation_label:", sub.AliceConversationLabel)
	fmt.Println("wrong_conversation_label:", wrongConversationLabel)
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

		if err := assertRelayOpenMLSTrustCandidateAbsent("before same-state message-failure relay join", sub.AliceState, sub.BobState); err != nil {
			return err
		}

		if err := r.runRelayOpenMLSSmokeScript(&sub); err != nil {
			return err
		}

		if err := r.assertSameStateBobConversationReloadable(&sub); err != nil {
			return err
		}

		if err := assertRelayOpenMLSTrustCandidateAbsent("after same-state message-failure relay join", sub.AliceState, sub.BobState); err != nil {
			return err
		}

		if err := r.runSameStateWrongConversationNoAckThenCorrectOpen(&sub, wrongConversationLabel, messageLabel, messageText); err != nil {
			return err
		}

		if err := assertRelayOpenMLSTrustCandidateAbsent("after same-state message-failure normal message open", sub.AliceState, sub.BobState); err != nil {
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
		return relayOpenMLSJoinSubrunResult{}, fmt.Errorf("same-state message-failure envelope count = %d, want 3", result.EnvelopeCount)
	}
	if result.AckCount != 2 {
		return relayOpenMLSJoinSubrunResult{}, fmt.Errorf("same-state message-failure ack count = %d, want 2", result.AckCount)
	}

	fmt.Println("PASS: same-state message failure subrun", sub.Name)
	return result, nil
}

func (r *Runner) runSameStateWrongConversationNoAckThenCorrectOpen(sub *relayOpenMLSJoinSubrun, wrongConversationLabel string, messageLabel string, messageText string) error {
	fmt.Println()
	fmt.Println("== Same-state failure setup: normal message send ==")

	sendOutput, err := runRelayOpenMLSCommand(
		"comms message-send-dev same-state message-failure",
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
		"status: sent",
		"recipient_device_id: " + sub.BobDeviceID,
		"conversation: " + sub.AliceConversationLabel,
		"message_label: " + messageLabel,
	} {
		if !strings.Contains(sendOutput, needle) {
			return fmt.Errorf("message-send-dev same-state failure setup output missing required evidence line %q", needle)
		}
	}

	acksAfterJoin, err := sameStateEnvelopeAckCount(sub.DBPath)
	if err != nil {
		return err
	}
	inboxBeforeBadOpen, err := sameStateDeviceInboxCount(sub.BaseURL, sub.RelaySpace, sub.BobDeviceID)
	if err != nil {
		return err
	}
	if inboxBeforeBadOpen != 1 {
		return fmt.Errorf("Bob inbox count before wrong-conversation open = %d, want 1", inboxBeforeBadOpen)
	}

	fmt.Println("acks_after_join:", acksAfterJoin)
	fmt.Println("bob_inbox_count_before_bad_open:", inboxBeforeBadOpen)

	fmt.Println()
	fmt.Println("== Same-state failure proof: wrong conversation open must not ack ==")

	wrongOutput, err := runRelayOpenMLSCommand(
		"comms message-inbox-dev wrong-conversation no-ack proof",
		r.Comms,
		[]string{"RUST_BACKTRACE=1"},
		"go",
		"run", "./cmd/comms",
		"message-inbox-dev",
		"--relay-space", sub.RelaySpace,
		"--state", sub.BobState,
		"--sidecar-device-label", sub.BobSidecarLabel,
		"--conversation", wrongConversationLabel,
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
		"message open failed",
		"stage: message_open",
		"acked: false",
		"opened_envelopes: 0",
		"open_failures: 1",
		"ack_failures: 0",
	} {
		if !strings.Contains(wrongOutput, needle) {
			return fmt.Errorf("wrong-conversation message-inbox-dev output missing required evidence line %q", needle)
		}
	}

	for _, forbidden := range []string{
		"acked: true",
		"ack_delivery_state: acknowledged",
		"message opened",
	} {
		if strings.Contains(wrongOutput, forbidden) {
			return fmt.Errorf("wrong-conversation message-inbox-dev output contained forbidden evidence line %q", forbidden)
		}
	}

	acksAfterBadOpen, err := sameStateEnvelopeAckCount(sub.DBPath)
	if err != nil {
		return err
	}
	inboxAfterBadOpen, err := sameStateDeviceInboxCount(sub.BaseURL, sub.RelaySpace, sub.BobDeviceID)
	if err != nil {
		return err
	}

	fmt.Println("acks_after_bad_open:", acksAfterBadOpen)
	fmt.Println("bob_inbox_count_after_bad_open:", inboxAfterBadOpen)

	if acksAfterBadOpen != acksAfterJoin {
		return fmt.Errorf("wrong-conversation open changed ack count: before=%d after=%d", acksAfterJoin, acksAfterBadOpen)
	}
	if inboxAfterBadOpen != inboxBeforeBadOpen {
		return fmt.Errorf("wrong-conversation open changed Bob inbox count: before=%d after=%d", inboxBeforeBadOpen, inboxAfterBadOpen)
	}

	fmt.Println("PASS: wrong-conversation open did not ack and did not drain Bob inbox")

	fmt.Println()
	fmt.Println("== Same-state recovery proof: correct conversation still opens and acks ==")

	correctOutput, err := runRelayOpenMLSCommand(
		"comms message-inbox-dev correct conversation after failure",
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
	inboxAfterCorrectOpen, err := sameStateDeviceInboxCount(sub.BaseURL, sub.RelaySpace, sub.BobDeviceID)
	if err != nil {
		return err
	}

	fmt.Println("acks_after_correct_open:", acksAfterCorrectOpen)
	fmt.Println("bob_inbox_count_after_correct_open:", inboxAfterCorrectOpen)

	if acksAfterCorrectOpen != acksAfterJoin+1 {
		return fmt.Errorf("correct open ack count = %d, want %d", acksAfterCorrectOpen, acksAfterJoin+1)
	}
	if inboxAfterCorrectOpen != 0 {
		return fmt.Errorf("Bob inbox count after correct open/ack = %d, want 0", inboxAfterCorrectOpen)
	}

	fmt.Println("PASS: correct conversation open/ack succeeded after failed wrong-conversation attempt")
	return nil
}

func sameStateEnvelopeAckCount(dbPath string) (int, error) {
	return relayOpenMLSSQLiteCount(dbPath, "SELECT COUNT(*) FROM envelope_acks;")
}
