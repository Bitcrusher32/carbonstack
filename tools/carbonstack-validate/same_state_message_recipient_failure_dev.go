package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func (r *Runner) SameStateMessageRecipientFailureDev() error {
	r.PrintHeader("same-state-message-recipient-failure-dev")

	fmt.Println("status: dev/pre-alpha same-state recipient/device failure validation profile")
	fmt.Println("scope: wrong recipient/device/sidecar attempts after same-state Relay join")
	fmt.Println("proof: Relay join -> normal message send to Bob -> Alice inbox / wrong sidecar / missing sidecar no false-success/no-ack/no-drain -> correct Bob open/ack")
	fmt.Println("boundary: live umbrella only; not full; not release-snapshot; not package-root validation; not adversarial harness; not identity verification; not production/security proof")
	fmt.Println("relationship: failure-path companion to same-state-integrated-dev; separate from wrong-conversation and unsupported content-type profiles")
	fmt.Println("nonclaims: not hostile-server safety, not metadata privacy, not production secure messaging, not verified identity, not mature messenger UX")
	fmt.Println()

	if err := r.CheckRequiredPaths(); err != nil {
		return err
	}
	if err := r.CheckLiveGitUmbrella("same-state-message-recipient-failure-dev"); err != nil {
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

	tempRoot, err := os.MkdirTemp("", "carbonstack-same-state-message-recipient-failure-dev-*")
	if err != nil {
		return fmt.Errorf("create same-state recipient-failure temp root: %w", err)
	}
	defer func() {
		if err := os.RemoveAll(tempRoot); err != nil {
			fmt.Printf("WARN: remove same-state recipient-failure temp root %s: %v\n", tempRoot, err)
		}
	}()

	binPath := filepath.Join(tempRoot, "carbonstack-cypher-same-state-recipient-failure")
	uniqueID := relayOpenMLSRunID()
	runID := "same-state-recipient-failure-" + uniqueID
	messageLabel := "rcpt-msg-" + uniqueID
	messageText := "CarbonStack same-state recipient failure proof payload " + runID

	fmt.Println()
	fmt.Println("== same-state recipient-failure generated-state root ==")
	fmt.Println("temp_root:", tempRoot)
	fmt.Println("cypher_bin:", binPath)
	fmt.Println("run_id:", runID)
	fmt.Println("message_label:", messageLabel)
	fmt.Println("note: this profile proves wrong recipient/device/sidecar attempts do not falsely open, ack, or drain Bob's inbox")

	fmt.Println()
	fmt.Println("== build temporary Cypher binary ==")
	if err := runLocalCypherCommand(r.Cypher, "go", "build", "-o", binPath, "./cmd/cypher"); err != nil {
		return err
	}

	r.ArtifactScan("pre-same-state-message-recipient-failure-dev")

	result, err := r.runSameStateMessageRecipientFailureSubrun(binPath, tempRoot, relayOpenMLSJoinSubrun{
		Name:     "recipient-device-no-false-success",
		AckAfter: true,
		RunID:    runID,
	}, messageLabel, messageText)
	if err != nil {
		return err
	}

	r.ArtifactScan("post-same-state-message-recipient-failure-dev")

	fmt.Println()
	fmt.Println("same-state-message-recipient-failure-dev profile result:")
	fmt.Println("  PASS: Relay onboarding completed through KeyPackage -> add-member -> Welcome -> join")
	fmt.Println("  PASS: normal message sent to Bob after same-state join")
	fmt.Println("  PASS: Alice state + Alice sidecar did not falsely open/ack/drain Bob message")
	fmt.Println("  PASS: Bob state + Alice sidecar did not falsely open/ack/drain Bob message")
	fmt.Println("  PASS: Bob state + missing sidecar did not falsely open/ack/drain Bob message")
	fmt.Println("  PASS: correct Bob sidecar open/ack still succeeded after wrong recipient/device attempts")
	fmt.Println("  proof_level: same-state wrong recipient/device no false-success/no-ack/no-drain proof")
	fmt.Println("  relay_space_id:", result.RelaySpaceID)
	fmt.Println("  envelopes:", result.EnvelopeCount)
	fmt.Println("  envelope_acks:", result.AckCount)
	fmt.Println("  keypackage_delivery_state:", result.KeyPackageDeliveryState)
	fmt.Println("  welcome_delivery_state:", result.WelcomeDeliveryState)
	fmt.Println("  boundary: live-dev failure-path proof; not full, not release-snapshot, not adversarial harness, not identity verification, not production/security proof")
	fmt.Println("  relationship: same-state-integrated-dev remains positive-path; previous failure profiles remain narrow")
	fmt.Println("  nonclaims: not hostile-server safety, not metadata privacy, not production secure messaging, not identity verification")

	return nil
}

func (r *Runner) runSameStateMessageRecipientFailureSubrun(binPath string, tempRoot string, sub relayOpenMLSJoinSubrun, messageLabel string, messageText string) (relayOpenMLSJoinSubrunResult, error) {
	fmt.Println()
	fmt.Println("========================================")
	fmt.Println("Same-state message recipient-failure subrun:", sub.Name)
	fmt.Println("========================================")

	sub.TempDir = filepath.Join(tempRoot, sub.Name)
	if err := os.Mkdir(sub.TempDir, 0700); err != nil {
		return relayOpenMLSJoinSubrunResult{}, fmt.Errorf("create same-state recipient-failure subrun temp dir %s: %w", sub.TempDir, err)
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
	sub.RelaySpace = "same-state-message-recipient-failure-dev-" + sub.RunID
	sub.AliceSidecarLabel = "carbonstack-" + sub.RunID + "-alice-device"
	sub.BobSidecarLabel = "carbonstack-" + sub.RunID + "-bob-device"
	sub.AliceConversationLabel = "carbonstack-" + sub.RunID + "-conversation"
	sub.BobConversationLabel = sub.AliceConversationLabel

	missingSidecarLabel := "carbonstack-" + sub.RunID + "-missing-device"

	fmt.Println("subrun_temp_dir:", sub.TempDir)
	fmt.Println("cypher_addr:", addr)
	fmt.Println("cypher_db:", sub.DBPath)
	fmt.Println("relay_space_id:", sub.RelaySpace)
	fmt.Println("alice_sidecar_label:", sub.AliceSidecarLabel)
	fmt.Println("bob_sidecar_label:", sub.BobSidecarLabel)
	fmt.Println("missing_sidecar_label:", missingSidecarLabel)
	fmt.Println("conversation_label:", sub.AliceConversationLabel)
	fmt.Println("message_label:", messageLabel)

	if err := r.refuseExistingSidecarDevice(sub.AliceSidecarLabel); err != nil {
		return relayOpenMLSJoinSubrunResult{}, err
	}
	if err := r.refuseExistingSidecarDevice(sub.BobSidecarLabel); err != nil {
		return relayOpenMLSJoinSubrunResult{}, err
	}
	if err := r.refuseExistingSidecarDevice(missingSidecarLabel); err != nil {
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

		if err := assertRelayOpenMLSTrustCandidateAbsent("before same-state recipient-failure relay join", sub.AliceState, sub.BobState); err != nil {
			return err
		}

		if err := r.runRelayOpenMLSSmokeScript(&sub); err != nil {
			return err
		}

		if err := r.assertSameStateBobConversationReloadable(&sub); err != nil {
			return err
		}

		if err := assertRelayOpenMLSTrustCandidateAbsent("after same-state recipient-failure relay join", sub.AliceState, sub.BobState); err != nil {
			return err
		}

		if err := r.runSameStateRecipientFailureNoAckThenCorrectOpen(&sub, missingSidecarLabel, messageLabel, messageText); err != nil {
			return err
		}

		if err := assertRelayOpenMLSTrustCandidateAbsent("after same-state recipient-failure normal message open", sub.AliceState, sub.BobState); err != nil {
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
		return relayOpenMLSJoinSubrunResult{}, fmt.Errorf("same-state recipient-failure envelope count = %d, want 3", result.EnvelopeCount)
	}
	if result.AckCount != 2 {
		return relayOpenMLSJoinSubrunResult{}, fmt.Errorf("same-state recipient-failure ack count = %d, want 2", result.AckCount)
	}

	fmt.Println("PASS: same-state message recipient-failure subrun", sub.Name)
	return result, nil
}

func (r *Runner) runSameStateRecipientFailureNoAckThenCorrectOpen(sub *relayOpenMLSJoinSubrun, missingSidecarLabel string, messageLabel string, messageText string) error {
	fmt.Println()
	fmt.Println("== Same-state recipient-failure setup: normal message send ==")

	sendOutput, err := runRelayOpenMLSCommand(
		"comms message-send-dev same-state recipient-failure setup",
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
			return fmt.Errorf("message-send-dev recipient-failure setup output missing required evidence line %q", needle)
		}
	}

	acksBeforeWrongAttempts, err := sameStateEnvelopeAckCount(sub.DBPath)
	if err != nil {
		return err
	}
	bobInboxBeforeWrongAttempts, err := sameStateDeviceInboxCount(sub.BaseURL, sub.BobDeviceID)
	if err != nil {
		return err
	}
	aliceInboxBeforeWrongAttempts, err := sameStateDeviceInboxCount(sub.BaseURL, sub.AliceDeviceID)
	if err != nil {
		return err
	}

	if bobInboxBeforeWrongAttempts != 1 {
		return fmt.Errorf("Bob inbox before wrong recipient/device attempts = %d, want 1", bobInboxBeforeWrongAttempts)
	}

	fmt.Println("acks_before_wrong_attempts:", acksBeforeWrongAttempts)
	fmt.Println("bob_inbox_count_before_wrong_attempts:", bobInboxBeforeWrongAttempts)
	fmt.Println("alice_inbox_count_before_wrong_attempts:", aliceInboxBeforeWrongAttempts)

	cases := []struct {
		Name        string
		StatePath   string
		Sidecar     string
		Description string
	}{
		{
			Name:        "case_a_alice_state_alice_sidecar",
			StatePath:   sub.AliceState,
			Sidecar:     sub.AliceSidecarLabel,
			Description: "Alice state + Alice sidecar must not false-open Bob's message",
		},
		{
			Name:        "case_b_bob_state_alice_sidecar",
			StatePath:   sub.BobState,
			Sidecar:     sub.AliceSidecarLabel,
			Description: "Bob state + Alice sidecar must not false-open Bob's message",
		},
		{
			Name:        "case_c_bob_state_missing_sidecar",
			StatePath:   sub.BobState,
			Sidecar:     missingSidecarLabel,
			Description: "Bob state + missing sidecar must not false-open Bob's message",
		},
	}

	for _, tc := range cases {
		fmt.Println()
		fmt.Println("== Same-state recipient-failure proof:", tc.Description, "==")

		output, err := runRelayOpenMLSCommand(
			"comms message-inbox-dev "+tc.Name,
			r.Comms,
			[]string{"RUST_BACKTRACE=1"},
			"go",
			"run", "./cmd/comms",
			"message-inbox-dev",
			"--state", tc.StatePath,
			"--sidecar-device-label", tc.Sidecar,
			"--conversation", sub.BobConversationLabel,
			"--message-label", messageLabel,
			"--ack",
		)
		if err != nil {
			return err
		}

		if err := assertSameStateRecipientFailureNoFalseSuccess(output, tc.Name); err != nil {
			return err
		}

		acksAfterWrongAttempt, err := sameStateEnvelopeAckCount(sub.DBPath)
		if err != nil {
			return err
		}
		bobInboxAfterWrongAttempt, err := sameStateDeviceInboxCount(sub.BaseURL, sub.BobDeviceID)
		if err != nil {
			return err
		}

		fmt.Println(tc.Name+"_acks_after:", acksAfterWrongAttempt)
		fmt.Println(tc.Name+"_bob_inbox_after:", bobInboxAfterWrongAttempt)

		if acksAfterWrongAttempt != acksBeforeWrongAttempts {
			return fmt.Errorf("%s changed ack count: before=%d after=%d", tc.Name, acksBeforeWrongAttempts, acksAfterWrongAttempt)
		}
		if bobInboxAfterWrongAttempt != bobInboxBeforeWrongAttempts {
			return fmt.Errorf("%s changed Bob inbox count: before=%d after=%d", tc.Name, bobInboxBeforeWrongAttempts, bobInboxAfterWrongAttempt)
		}

		fmt.Println("PASS:", tc.Name, "did not false-open, ack, or drain Bob inbox")
	}

	fmt.Println()
	fmt.Println("== Same-state recovery proof: correct Bob open and ack after recipient/device failures ==")

	correctOutput, err := runRelayOpenMLSCommand(
		"comms message-inbox-dev correct Bob after recipient/device failures",
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
			return fmt.Errorf("correct Bob message-inbox-dev output missing required evidence line %q", needle)
		}
	}

	acksAfterCorrectOpen, err := sameStateEnvelopeAckCount(sub.DBPath)
	if err != nil {
		return err
	}
	bobInboxAfterCorrectOpen, err := sameStateDeviceInboxCount(sub.BaseURL, sub.BobDeviceID)
	if err != nil {
		return err
	}

	fmt.Println("acks_after_correct_open:", acksAfterCorrectOpen)
	fmt.Println("bob_inbox_count_after_correct_open:", bobInboxAfterCorrectOpen)

	if acksAfterCorrectOpen != acksBeforeWrongAttempts+1 {
		return fmt.Errorf("correct open ack count = %d, want %d", acksAfterCorrectOpen, acksBeforeWrongAttempts+1)
	}
	if bobInboxAfterCorrectOpen != 0 {
		return fmt.Errorf("Bob inbox count after correct open/ack = %d, want 0", bobInboxAfterCorrectOpen)
	}

	fmt.Println("PASS: correct Bob open/ack succeeded after recipient/device failure attempts")
	return nil
}

func assertSameStateRecipientFailureNoFalseSuccess(output string, label string) error {
	for _, required := range []string{
		"message inbox",
		"command: message-inbox-dev",
		"implementation_path: openmls-inbox-dev",
		"ack_requested: true",
		"opened_envelopes: 0",
		"ack_failures: 0",
	} {
		if !strings.Contains(output, required) {
			return fmt.Errorf("%s output missing required no-false-success evidence line %q", label, required)
		}
	}

	for _, forbidden := range []string{
		"acked: true",
		"ack_delivery_state: acknowledged",
		"message opened",
	} {
		if strings.Contains(output, forbidden) {
			return fmt.Errorf("%s output contained forbidden false-success evidence line %q", label, forbidden)
		}
	}

	if !strings.Contains(output, "unsupported_envelopes: 1") && !strings.Contains(output, "open_failures: 1") {
		return fmt.Errorf("%s output did not show either unsupported skip or message-open failure", label)
	}

	return nil
}
