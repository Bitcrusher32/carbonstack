package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type replayClassificationCase struct {
	ID   string
	Name string
	Kind string
}

func (r *Runner) SameStateMessageReplayClassificationDev() error {
	r.PrintHeader("same-state-message-replay-classification-dev")

	fmt.Println("status: dev/pre-alpha same-state normal application-message replay classification profile")
	fmt.Println("scope: duplicate/replayed normal application-message envelope classification")
	fmt.Println("proof: Relay KeyPackage -> add-member -> Welcome -> join -> message-send-dev -> replay/duplicate case -> classify counters and output")
	fmt.Println("boundary: live umbrella only; not full; not release-snapshot; not package-root validation; not adversarial harness; not production/security proof")
	fmt.Println("relationship: failure-path companion after same-state-message-malformed-payload-dev")
	fmt.Println("nonclaims: not replay safety, not hostile-server safety, not metadata privacy, not production secure messaging, not verified identity")
	fmt.Println()

	if err := r.CheckRequiredPaths(); err != nil {
		return err
	}
	if err := r.CheckLiveGitUmbrella("same-state-message-replay-classification-dev"); err != nil {
		return err
	}

	fmt.Println()
	fmt.Println("== Toolchains ==")
	_ = r.ReportTool("go", "version")
	_ = r.ReportTool("rustc", "--version")
	_ = r.ReportTool("cargo", "--version")
	_ = r.ReportTool("bash", "--version")
	_ = r.ReportTool("sqlite3", "--version")

	tempRoot, err := os.MkdirTemp("", "carbonstack-same-state-message-replay-classification-dev-*")
	if err != nil {
		return fmt.Errorf("create replay-classification temp root: %w", err)
	}
	defer func() {
		if err := os.RemoveAll(tempRoot); err != nil {
			fmt.Printf("WARN: remove replay-classification temp root %s: %v\n", tempRoot, err)
		}
	}()

	binPath := filepath.Join(tempRoot, "carbonstack-cypher-same-state-message-replay-classification")
	runID := relayOpenMLSRunID()

	fmt.Println()
	fmt.Println("== same-state replay classification generated-state root ==")
	fmt.Println("temp_root:", tempRoot)
	fmt.Println("cypher_bin:", binPath)
	fmt.Println("run_id:", runID)
	fmt.Println("note: this profile classifies normal application-message duplicate/replay behavior; it does not claim replay safety")

	fmt.Println()
	fmt.Println("== build temporary Cypher binary ==")
	if err := runLocalCypherCommand(r.Cypher, "go", "build", "-o", binPath, "./cmd/cypher"); err != nil {
		return err
	}

	r.ArtifactScan("pre-same-state-message-replay-classification-dev")

	cases := []replayClassificationCase{
		{ID: "r01", Name: "duplicate-same-envelope-id-insert", Kind: "duplicate_same_envelope_id_insert"},
		{ID: "r02", Name: "same-envelope-after-ack-no-requeue", Kind: "same_envelope_after_ack_no_requeue"},
		{ID: "r03", Name: "same-envelope-manual-requeue-after-ack", Kind: "manual_requeue_same_envelope_after_ack"},
		{ID: "r04", Name: "duplicate-new-envelope-before-ack", Kind: "duplicate_new_envelope_before_ack"},
		{ID: "r05", Name: "duplicate-new-envelope-after-ack", Kind: "duplicate_new_envelope_after_ack"},
	}

	for _, tc := range cases {
		if err := r.runReplayClassificationSubrun(binPath, tempRoot, runID, tc); err != nil {
			return err
		}
	}

	r.ArtifactScan("post-same-state-message-replay-classification-dev")

	fmt.Println()
	fmt.Println("same-state-message-replay-classification-dev profile result:")
	fmt.Println("  PASS: replay classification cases tested: 5")
	fmt.Println("  PASS: duplicate same envelope_id insert is storage rejected")
	fmt.Println("  PASS: same envelope after ack is delivery-state suppressed")
	fmt.Println("  PASS: manually requeued same envelope after ack fails no-open/no-ack/no-drain")
	fmt.Println("  PASS: duplicate same ciphertext under new envelope ID before original ack fails no-open/no-ack/no-drain")
	fmt.Println("  PASS: duplicate same ciphertext under new envelope ID after original ack fails no-open/no-ack/no-drain")
	fmt.Println("  proof_level: same-state normal application-message duplicate/replay classification proof")
	fmt.Println("  boundary: live-dev classification proof; not full, not release-snapshot, not adversarial harness, not replay-safety claim, not production/security proof")
	fmt.Println("  nonclaims: not hostile-server safety, not metadata privacy, not production secure messaging, not identity verification")

	return nil
}

func (r *Runner) runReplayClassificationSubrun(binPath string, tempRoot string, runID string, tc replayClassificationCase) error {
	fmt.Println()
	fmt.Println("========================================")
	fmt.Println("Same-state replay classification subrun:", tc.ID, tc.Name)
	fmt.Println("========================================")

	shortID := malformedPayloadShortID(runID + "-replay-" + tc.ID)

	sub := relayOpenMLSJoinSubrun{
		Name:     "replay-classification-" + tc.ID + "-" + tc.Name,
		AckAfter: true,
		RunID:    "replay-classification-" + tc.ID + "-" + shortID,
	}
	sub.TempDir = filepath.Join(tempRoot, sub.Name)
	if err := os.Mkdir(sub.TempDir, 0700); err != nil {
		return fmt.Errorf("create replay-classification temp dir %s: %w", sub.TempDir, err)
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
	sub.RelaySpace = "same-state-replay-classification-dev-" + tc.ID + "-" + shortID
	sub.AliceSidecarLabel = "cs-rp-" + tc.ID + "-a-" + shortID
	sub.BobSidecarLabel = "cs-rp-" + tc.ID + "-b-" + shortID
	sub.AliceConversationLabel = "cs-rp-" + tc.ID + "-cv-" + shortID
	sub.BobConversationLabel = sub.AliceConversationLabel

	messageLabel := "rp-" + tc.ID + "-" + shortID
	messageText := "CarbonStack replay classification profile " + tc.ID

	fmt.Println("subrun_temp_dir:", sub.TempDir)
	fmt.Println("cypher_addr:", addr)
	fmt.Println("cypher_db:", sub.DBPath)
	fmt.Println("relay_space_id:", sub.RelaySpace)
	fmt.Println("alice_sidecar_label:", sub.AliceSidecarLabel)
	fmt.Println("bob_sidecar_label:", sub.BobSidecarLabel)
	fmt.Println("conversation_label:", sub.AliceConversationLabel)
	fmt.Println("message_label:", messageLabel)
	fmt.Println("replay_case:", tc.Kind)

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
		if err := assertRelayOpenMLSTrustCandidateAbsent("before replay classification proof", sub.AliceState, sub.BobState); err != nil {
			return err
		}

		if err := r.runMalformedPayloadJoin(&sub); err != nil {
			return err
		}

		if err := r.runReplayClassificationCase(&sub, tc, messageLabel, messageText); err != nil {
			return err
		}

		if err := assertRelayOpenMLSTrustCandidateAbsent("after replay classification proof", sub.AliceState, sub.BobState); err != nil {
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

	fmt.Println("PASS: same-state replay classification subrun", tc.ID, tc.Name)
	return nil
}

func (r *Runner) runReplayClassificationCase(sub *relayOpenMLSJoinSubrun, tc replayClassificationCase, messageLabel string, messageText string) error {
	fmt.Println()
	fmt.Println("== Send normal message and classify replay/duplicate behavior ==")

	appEnvelopeID, err := r.replaySendAndLocateAppEnvelope(sub, messageLabel, messageText)
	if err != nil {
		return err
	}
	fmt.Println("app_envelope_id:", appEnvelopeID)

	switch tc.Kind {
	case "duplicate_same_envelope_id_insert":
		return r.replayCaseDuplicateSameEnvelopeID(sub, appEnvelopeID, messageLabel, messageText)
	case "same_envelope_after_ack_no_requeue":
		return r.replayCaseSameEnvelopeAfterAckNoRequeue(sub, appEnvelopeID, messageLabel, messageText)
	case "manual_requeue_same_envelope_after_ack":
		return r.replayCaseManualRequeueSameEnvelopeAfterAck(sub, appEnvelopeID, messageLabel, messageText)
	case "duplicate_new_envelope_before_ack":
		return r.replayCaseDuplicateNewEnvelopeBeforeAck(sub, appEnvelopeID, messageLabel, messageText)
	case "duplicate_new_envelope_after_ack":
		return r.replayCaseDuplicateNewEnvelopeAfterAck(sub, appEnvelopeID, messageLabel, messageText)
	default:
		return fmt.Errorf("unknown replay classification case kind %q", tc.Kind)
	}
}

func (r *Runner) replaySendAndLocateAppEnvelope(sub *relayOpenMLSJoinSubrun, messageLabel string, messageText string) (string, error) {
	sendOutput, sendRC, err := runWelcomeJoinFailureCommand(
		r.Comms,
		[]string{"RUST_BACKTRACE=1"},
		"go", "run", "./cmd/comms",
		"message-send-dev",
		"--relay-space", sub.RelaySpace,
		"--state", sub.AliceState,
		"--to-device", sub.BobDeviceID,
		"--sidecar-device-label", sub.AliceSidecarLabel,
		"--conversation", sub.AliceConversationLabel,
		"--message-label", messageLabel,
		"--message", messageText,
	)
	fmt.Print(sendOutput)
	if err != nil || sendRC != 0 {
		return "", fmt.Errorf("message-send-dev failed rc=%d err=%v", sendRC, err)
	}
	if !strings.Contains(sendOutput, "message sent") {
		return "", fmt.Errorf("message-send-dev output missing message sent")
	}

	appEnvelopeID, err := sqlite3QueryOne(sub.DBPath, fmt.Sprintf(
		"SELECT envelope_id FROM envelopes WHERE recipient_device_id = %s AND content_type = 'carbonstack.mls.application-message.v0' AND delivery_state = 'queued' ORDER BY rowid ASC LIMIT 1;",
		sqlite3Quote(sub.BobDeviceID),
	))
	if err != nil {
		return "", err
	}
	if appEnvelopeID == "" {
		return "", fmt.Errorf("could not locate Bob application-message envelope")
	}

	return appEnvelopeID, nil
}

func (r *Runner) replayOpenInbox(sub *relayOpenMLSJoinSubrun, label string, messageLabel string) (string, int, error) {
	fmt.Println()
	fmt.Println("== message-inbox-dev --ack:", label, "==")
	output, rc, err := runWelcomeJoinFailureCommand(
		r.Comms,
		[]string{"RUST_BACKTRACE=1"},
		"go", "run", "./cmd/comms",
		"message-inbox-dev",
		"--relay-space", sub.RelaySpace,
		"--state", sub.BobState,
		"--sidecar-device-label", sub.BobSidecarLabel,
		"--conversation", sub.BobConversationLabel,
		"--message-label", messageLabel,
		"--ack",
	)
	fmt.Print(output)
	fmt.Println(label+"_rc:", rc)
	if err != nil && rc < 0 {
		return output, rc, fmt.Errorf("%s message-inbox-dev non-process error rc=%d err=%v", label, rc, err)
	}
	return output, rc, nil
}

func (r *Runner) replayAssertOpenedAcked(output string, label string, messageText string) error {
	for _, needle := range []string{
		"message opened",
		"ack_delivery_state: acknowledged",
		"acked: true",
		"opened_envelopes: 1",
		"open_failures: 0",
		"ack_failures: 0",
		"plaintext_utf8: " + messageText,
	} {
		if !strings.Contains(output, needle) {
			return fmt.Errorf("%s output missing %q", label, needle)
		}
	}
	return nil
}

func (r *Runner) replayAssertRejectedNoAck(output string, label string) error {
	for _, needle := range []string{
		"message open failed",
		"acked: false",
		"opened_envelopes: 0",
		"open_failures: 1",
		"ack_failures: 0",
	} {
		if !strings.Contains(output, needle) {
			return fmt.Errorf("%s output missing %q", label, needle)
		}
	}
	for _, forbidden := range []string{
		"message opened",
		"acked: true",
		"ack_delivery_state: acknowledged",
	} {
		if strings.Contains(output, forbidden) {
			return fmt.Errorf("%s output contained forbidden marker %q", label, forbidden)
		}
	}
	return nil
}

func (r *Runner) replayAssertEmptySuppressed(output string, label string) error {
	for _, needle := range []string{
		"queued_envelopes: 0",
		"opened_envelopes: 0",
		"open_failures: 0",
		"ack_failures: 0",
	} {
		if !strings.Contains(output, needle) {
			return fmt.Errorf("%s output missing %q", label, needle)
		}
	}
	for _, forbidden := range []string{
		"message opened",
		"acked: true",
		"ack_delivery_state: acknowledged",
	} {
		if strings.Contains(output, forbidden) {
			return fmt.Errorf("%s output contained forbidden marker %q", label, forbidden)
		}
	}
	return nil
}

func (r *Runner) replayCaseDuplicateSameEnvelopeID(sub *relayOpenMLSJoinSubrun, envelopeID string, messageLabel string, messageText string) error {
	acksBefore, err := replayBobAppAckCount(sub.DBPath, sub.BobDeviceID)
	if err != nil {
		return err
	}
	inboxBefore, err := replayBobQueuedAppCount(sub.DBPath, sub.BobDeviceID)
	if err != nil {
		return err
	}

	insertOK, insertErr, err := replayTryDuplicateSameEnvelopeID(sub.DBPath, envelopeID)
	if err != nil {
		return err
	}
	fmt.Println("duplicate_same_envelope_id_insert_ok:", insertOK)
	fmt.Println("duplicate_same_envelope_id_insert_error:", insertErr)

	if insertOK {
		return fmt.Errorf("duplicate same envelope_id insert unexpectedly succeeded")
	}
	if !strings.Contains(insertErr, "UNIQUE constraint failed") {
		return fmt.Errorf("duplicate same envelope_id insert did not report expected uniqueness failure: %s", insertErr)
	}

	output, _, err := r.replayOpenInbox(sub, "open-original-after-duplicate-id-insert", messageLabel)
	if err != nil {
		return err
	}
	if err := r.replayAssertOpenedAcked(output, "open-original-after-duplicate-id-insert", messageText); err != nil {
		return err
	}

	acksAfter, err := replayBobAppAckCount(sub.DBPath, sub.BobDeviceID)
	if err != nil {
		return err
	}
	inboxAfter, err := replayBobQueuedAppCount(sub.DBPath, sub.BobDeviceID)
	if err != nil {
		return err
	}

	if acksAfter != acksBefore+1 {
		return fmt.Errorf("acks after original recovery = %d, want %d", acksAfter, acksBefore+1)
	}
	if inboxBefore != 1 || inboxAfter != 0 {
		return fmt.Errorf("inbox before/after original recovery = %d/%d, want 1/0", inboxBefore, inboxAfter)
	}

	fmt.Println("classification: storage_rejected_duplicate_same_envelope_id_original_recovery_ok")
	return nil
}

func (r *Runner) replayCaseSameEnvelopeAfterAckNoRequeue(sub *relayOpenMLSJoinSubrun, envelopeID string, messageLabel string, messageText string) error {
	acksBefore, err := replayBobAppAckCount(sub.DBPath, sub.BobDeviceID)
	if err != nil {
		return err
	}
	inboxBefore, err := replayBobQueuedAppCount(sub.DBPath, sub.BobDeviceID)
	if err != nil {
		return err
	}

	first, _, err := r.replayOpenInbox(sub, "first-open-original", messageLabel)
	if err != nil {
		return err
	}
	if err := r.replayAssertOpenedAcked(first, "first-open-original", messageText); err != nil {
		return err
	}

	acksMid, err := replayBobAppAckCount(sub.DBPath, sub.BobDeviceID)
	if err != nil {
		return err
	}
	inboxMid, err := replayBobQueuedAppCount(sub.DBPath, sub.BobDeviceID)
	if err != nil {
		return err
	}

	second, _, err := r.replayOpenInbox(sub, "second-open-same-envelope-after-ack", messageLabel)
	if err != nil {
		return err
	}
	if err := r.replayAssertEmptySuppressed(second, "second-open-same-envelope-after-ack"); err != nil {
		return err
	}

	acksAfter, err := replayBobAppAckCount(sub.DBPath, sub.BobDeviceID)
	if err != nil {
		return err
	}
	inboxAfter, err := replayBobQueuedAppCount(sub.DBPath, sub.BobDeviceID)
	if err != nil {
		return err
	}

	if acksMid != acksBefore+1 || acksAfter != acksMid {
		return fmt.Errorf("ack counts before/mid/after = %d/%d/%d, want 0/+1/unchanged", acksBefore, acksMid, acksAfter)
	}
	if inboxBefore != 1 || inboxMid != 0 || inboxAfter != 0 {
		return fmt.Errorf("inbox before/mid/after = %d/%d/%d, want 1/0/0", inboxBefore, inboxMid, inboxAfter)
	}

	fmt.Println("classification: delivery_state_suppressed_same_envelope_after_ack")
	return nil
}

func (r *Runner) replayCaseManualRequeueSameEnvelopeAfterAck(sub *relayOpenMLSJoinSubrun, envelopeID string, messageLabel string, messageText string) error {
	acksBefore, err := replayBobAppAckCount(sub.DBPath, sub.BobDeviceID)
	if err != nil {
		return err
	}
	inboxBefore, err := replayBobQueuedAppCount(sub.DBPath, sub.BobDeviceID)
	if err != nil {
		return err
	}

	first, _, err := r.replayOpenInbox(sub, "first-open-original", messageLabel)
	if err != nil {
		return err
	}
	if err := r.replayAssertOpenedAcked(first, "first-open-original", messageText); err != nil {
		return err
	}

	acksMid, err := replayBobAppAckCount(sub.DBPath, sub.BobDeviceID)
	if err != nil {
		return err
	}
	if err := sqlite3Exec(sub.DBPath, fmt.Sprintf("UPDATE envelopes SET delivery_state = 'queued' WHERE envelope_id = %s;", sqlite3Quote(envelopeID))); err != nil {
		return err
	}
	inboxRequeued, err := replayBobQueuedAppCount(sub.DBPath, sub.BobDeviceID)
	if err != nil {
		return err
	}
	if inboxRequeued != 1 {
		return fmt.Errorf("manual requeue inbox count = %d, want 1", inboxRequeued)
	}

	second, _, err := r.replayOpenInbox(sub, "second-open-requeued-same-envelope", messageLabel)
	if err != nil {
		return err
	}
	if err := r.replayAssertRejectedNoAck(second, "second-open-requeued-same-envelope"); err != nil {
		return err
	}

	acksAfter, err := replayBobAppAckCount(sub.DBPath, sub.BobDeviceID)
	if err != nil {
		return err
	}
	inboxAfter, err := replayBobQueuedAppCount(sub.DBPath, sub.BobDeviceID)
	if err != nil {
		return err
	}

	if acksMid != acksBefore+1 || acksAfter != acksMid {
		return fmt.Errorf("ack counts before/mid/after = %d/%d/%d, want 0/+1/unchanged", acksBefore, acksMid, acksAfter)
	}
	if inboxBefore != 1 || inboxAfter != 1 {
		return fmt.Errorf("inbox before/after rejected replay = %d/%d, want 1/1", inboxBefore, inboxAfter)
	}

	fmt.Println("classification: detected_or_rejected_no_ack_no_drain_manual_requeue_same_envelope")
	return nil
}

func (r *Runner) replayCaseDuplicateNewEnvelopeBeforeAck(sub *relayOpenMLSJoinSubrun, envelopeID string, messageLabel string, messageText string) error {
	duplicateID, err := replayDuplicateEnvelopeWithNewID(sub.DBPath, envelopeID)
	if err != nil {
		return err
	}
	fmt.Println("duplicate_envelope_id:", duplicateID)

	acksBefore, err := replayBobAppAckCount(sub.DBPath, sub.BobDeviceID)
	if err != nil {
		return err
	}
	inboxBefore, err := replayBobQueuedAppCount(sub.DBPath, sub.BobDeviceID)
	if err != nil {
		return err
	}
	if inboxBefore != 2 {
		return fmt.Errorf("duplicate-new-before-ack inbox before = %d, want 2", inboxBefore)
	}

	first, _, err := r.replayOpenInbox(sub, "first-open-with-duplicate-present", messageLabel)
	if err != nil {
		return err
	}
	if err := r.replayAssertOpenedAcked(first, "first-open-with-duplicate-present", messageText); err != nil {
		return err
	}

	acksMid, err := replayBobAppAckCount(sub.DBPath, sub.BobDeviceID)
	if err != nil {
		return err
	}
	inboxMid, err := replayBobQueuedAppCount(sub.DBPath, sub.BobDeviceID)
	if err != nil {
		return err
	}

	second, _, err := r.replayOpenInbox(sub, "second-open-duplicate-new-envelope", messageLabel)
	if err != nil {
		return err
	}
	if err := r.replayAssertRejectedNoAck(second, "second-open-duplicate-new-envelope"); err != nil {
		return err
	}

	acksAfter, err := replayBobAppAckCount(sub.DBPath, sub.BobDeviceID)
	if err != nil {
		return err
	}
	inboxAfter, err := replayBobQueuedAppCount(sub.DBPath, sub.BobDeviceID)
	if err != nil {
		return err
	}

	if acksMid != acksBefore+1 || acksAfter != acksMid {
		return fmt.Errorf("ack counts before/mid/after = %d/%d/%d, want 0/+1/unchanged", acksBefore, acksMid, acksAfter)
	}
	if inboxMid != 1 || inboxAfter != 1 {
		return fmt.Errorf("inbox mid/after duplicate replay = %d/%d, want 1/1", inboxMid, inboxAfter)
	}

	fmt.Println("classification: detected_or_rejected_no_ack_no_drain_duplicate_payload_new_envelope_before_ack")
	return nil
}

func (r *Runner) replayCaseDuplicateNewEnvelopeAfterAck(sub *relayOpenMLSJoinSubrun, envelopeID string, messageLabel string, messageText string) error {
	acksBefore, err := replayBobAppAckCount(sub.DBPath, sub.BobDeviceID)
	if err != nil {
		return err
	}
	inboxBefore, err := replayBobQueuedAppCount(sub.DBPath, sub.BobDeviceID)
	if err != nil {
		return err
	}

	first, _, err := r.replayOpenInbox(sub, "first-open-original", messageLabel)
	if err != nil {
		return err
	}
	if err := r.replayAssertOpenedAcked(first, "first-open-original", messageText); err != nil {
		return err
	}

	acksMid, err := replayBobAppAckCount(sub.DBPath, sub.BobDeviceID)
	if err != nil {
		return err
	}
	duplicateID, err := replayDuplicateEnvelopeWithNewID(sub.DBPath, envelopeID)
	if err != nil {
		return err
	}
	fmt.Println("duplicate_envelope_id:", duplicateID)

	inboxAfterInsert, err := replayBobQueuedAppCount(sub.DBPath, sub.BobDeviceID)
	if err != nil {
		return err
	}
	if inboxAfterInsert != 1 {
		return fmt.Errorf("duplicate-new-after-ack inbox after insert = %d, want 1", inboxAfterInsert)
	}

	second, _, err := r.replayOpenInbox(sub, "second-open-duplicate-new-envelope-after-ack", messageLabel)
	if err != nil {
		return err
	}
	if err := r.replayAssertRejectedNoAck(second, "second-open-duplicate-new-envelope-after-ack"); err != nil {
		return err
	}

	acksAfter, err := replayBobAppAckCount(sub.DBPath, sub.BobDeviceID)
	if err != nil {
		return err
	}
	inboxAfter, err := replayBobQueuedAppCount(sub.DBPath, sub.BobDeviceID)
	if err != nil {
		return err
	}

	if acksMid != acksBefore+1 || acksAfter != acksMid {
		return fmt.Errorf("ack counts before/mid/after = %d/%d/%d, want 0/+1/unchanged", acksBefore, acksMid, acksAfter)
	}
	if inboxBefore != 1 || inboxAfter != 1 {
		return fmt.Errorf("inbox before/after duplicate replay = %d/%d, want 1/1", inboxBefore, inboxAfter)
	}

	fmt.Println("classification: detected_or_rejected_no_ack_no_drain_duplicate_payload_new_envelope_after_ack")
	return nil
}

func replayBobQueuedAppCount(dbPath string, bobDeviceID string) (int, error) {
	return sqlite3Count(dbPath, fmt.Sprintf(
		"SELECT COUNT(*) FROM envelopes WHERE recipient_device_id = %s AND content_type = 'carbonstack.mls.application-message.v0' AND delivery_state = 'queued';",
		sqlite3Quote(bobDeviceID),
	))
}

func replayBobAppAckCount(dbPath string, bobDeviceID string) (int, error) {
	return sqlite3Count(dbPath, fmt.Sprintf(
		"SELECT COUNT(*) FROM envelope_acks WHERE envelope_id IN (SELECT envelope_id FROM envelopes WHERE recipient_device_id = %s AND content_type = 'carbonstack.mls.application-message.v0');",
		sqlite3Quote(bobDeviceID),
	))
}

func replayTryDuplicateSameEnvelopeID(dbPath string, envelopeID string) (bool, string, error) {
	script := fmt.Sprintf(`
import sqlite3
db = %q
envelope_id = %q
con = sqlite3.connect(db)
try:
    con.execute("INSERT INTO envelopes SELECT * FROM envelopes WHERE envelope_id = ?", (envelope_id,))
    con.commit()
    print("insert_ok=yes")
    print("insert_error=")
except Exception as e:
    print("insert_ok=no")
    print("insert_error=" + repr(e))
finally:
    con.close()
`, dbPath, envelopeID)

	output, err := replayRunPythonScript(script)
	if err != nil {
		return false, output, fmt.Errorf("duplicate same envelope_id python helper failed: %w\n%s", err, output)
	}

	insertOK := strings.Contains(output, "insert_ok=yes")
	insertErr := ""
	for _, line := range strings.Split(output, "\n") {
		if strings.HasPrefix(line, "insert_error=") {
			insertErr = strings.TrimPrefix(line, "insert_error=")
		}
	}

	return insertOK, insertErr, nil
}

func replayDuplicateEnvelopeWithNewID(dbPath string, envelopeID string) (string, error) {
	script := fmt.Sprintf(`
import datetime
import sqlite3
import uuid

db = %q
envelope_id = %q
new_id = str(uuid.uuid4())
con = sqlite3.connect(db)
cur = con.cursor()
cols = [r[1] for r in cur.execute("PRAGMA table_info(envelopes)").fetchall()]
row = cur.execute("SELECT " + ",".join(cols) + " FROM envelopes WHERE envelope_id = ?", (envelope_id,)).fetchone()
if row is None:
    raise SystemExit("missing envelope " + envelope_id)
data = dict(zip(cols, row))
data["envelope_id"] = new_id
if "delivery_state" in data:
    data["delivery_state"] = "queued"
now = datetime.datetime.now(datetime.UTC).replace(microsecond=0).isoformat().replace("+00:00", "Z")
for k in ("created_at", "updated_at", "server_received_at"):
    if k in data:
        data[k] = now
for k in ("acknowledged_at", "delivered_at"):
    if k in data:
        data[k] = None
placeholders = ",".join(["?"] * len(cols))
cur.execute("INSERT INTO envelopes (" + ",".join(cols) + ") VALUES (" + placeholders + ")", [data[c] for c in cols])
con.commit()
con.close()
print(new_id)
`, dbPath, envelopeID)

	output, err := replayRunPythonScript(script)
	if err != nil {
		return "", fmt.Errorf("duplicate envelope with new ID failed: %w\n%s", err, output)
	}

	return strings.TrimSpace(output), nil
}

func replayRunPythonScript(script string) (string, error) {
	cmd := exec.Command("python3", "-c", script)
	out, err := cmd.CombinedOutput()
	return string(out), err
}
