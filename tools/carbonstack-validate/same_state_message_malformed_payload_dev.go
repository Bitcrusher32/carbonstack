package main

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type malformedPayloadCase struct {
	ID       string
	Name     string
	Mutation string
}

type malformedPayloadSavedEnvelope struct {
	CiphertextB64    string
	PayloadSHA256    string
	PayloadSizeBytes int
}

type malformedPayloadMutation struct {
	CiphertextB64    string
	PayloadSHA256    string
	PayloadSizeBytes int
}

func (r *Runner) SameStateMessageMalformedPayloadDev() error {
	r.PrintHeader("same-state-message-malformed-payload-dev")

	fmt.Println("status: dev/pre-alpha same-state malformed normal-message payload validation profile")
	fmt.Println("scope: malformed normal application-message payload no-open/no-ack/no-drain/no-state-mutation plus restored payload recovery")
	fmt.Println("proof: Relay KeyPackage -> add-member -> Welcome -> join -> message-send-dev -> mutate payload -> message-inbox-dev --ack failure -> restore -> correct open/ack")
	fmt.Println("boundary: live umbrella only; not full; not release-snapshot; not package-root validation; not adversarial harness; not production/security proof")
	fmt.Println("relationship: failure-path companion to same-state-integrated-dev and same-state-message-unsupported-dev")
	fmt.Println("nonclaims: not hostile-server safety, not metadata privacy, not production secure messaging, not verified identity, not replay/duplicate classification")
	fmt.Println()

	if err := r.CheckRequiredPaths(); err != nil {
		return err
	}
	if err := r.CheckLiveGitUmbrella("same-state-message-malformed-payload-dev"); err != nil {
		return err
	}

	fmt.Println()
	fmt.Println("== Toolchains ==")
	_ = r.ReportTool("go", "version")
	_ = r.ReportTool("rustc", "--version")
	_ = r.ReportTool("cargo", "--version")
	_ = r.ReportTool("bash", "--version")
	_ = r.ReportTool("sqlite3", "--version")

	tempRoot, err := os.MkdirTemp("", "carbonstack-same-state-message-malformed-payload-dev-*")
	if err != nil {
		return fmt.Errorf("create malformed-payload temp root: %w", err)
	}
	defer func() {
		if err := os.RemoveAll(tempRoot); err != nil {
			fmt.Printf("WARN: remove malformed-payload temp root %s: %v\n", tempRoot, err)
		}
	}()

	binPath := filepath.Join(tempRoot, "carbonstack-cypher-same-state-message-malformed-payload")
	runID := relayOpenMLSRunID()

	fmt.Println()
	fmt.Println("== same-state malformed payload generated-state root ==")
	fmt.Println("temp_root:", tempRoot)
	fmt.Println("cypher_bin:", binPath)
	fmt.Println("run_id:", runID)
	fmt.Println("note: this profile proves malformed normal-message payload no-open/no-ack/no-drain/no-provider-mutation/no-envelope-rewrite plus restored payload recovery")

	fmt.Println()
	fmt.Println("== build temporary Cypher binary ==")
	if err := runLocalCypherCommand(r.Cypher, "go", "build", "-o", binPath, "./cmd/cypher"); err != nil {
		return err
	}

	r.ArtifactScan("pre-same-state-message-malformed-payload-dev")

	cases := []malformedPayloadCase{
		{ID: "p01", Name: "invalid-base64-storage-shape", Mutation: "invalid_base64_storage_shape"},
		{ID: "p02", Name: "valid-base64-random-bytes", Mutation: "valid_base64_random_bytes"},
		{ID: "p03", Name: "valid-base64-truncated-original", Mutation: "valid_base64_truncated_original"},
		{ID: "p04", Name: "valid-base64-single-byte-flip", Mutation: "valid_base64_single_byte_flip"},
		{ID: "p05", Name: "valid-base64-empty-bytes", Mutation: "valid_base64_empty_bytes"},
		{ID: "p06", Name: "valid-base64-original-plus-junk", Mutation: "valid_base64_original_plus_junk"},
	}

	for _, tc := range cases {
		if err := r.runMalformedPayloadSubrun(binPath, tempRoot, runID, tc); err != nil {
			return err
		}
	}

	r.ArtifactScan("post-same-state-message-malformed-payload-dev")

	fmt.Println()
	fmt.Println("same-state-message-malformed-payload-dev profile result:")
	fmt.Println("  PASS: malformed payload cases tested: 6")
	fmt.Println("  PASS: invalid base64 storage shape did not open, ack, drain, rewrite envelope, or mutate provider state")
	fmt.Println("  PASS: valid base64 random bytes did not open, ack, drain, rewrite envelope, or mutate provider state")
	fmt.Println("  PASS: truncated original payload did not open, ack, drain, rewrite envelope, or mutate provider state")
	fmt.Println("  PASS: single-byte-flipped original payload did not open, ack, drain, rewrite envelope, or mutate provider state")
	fmt.Println("  PASS: empty payload did not open, ack, drain, rewrite envelope, or mutate provider state")
	fmt.Println("  PASS: original payload plus junk did not open, ack, drain, rewrite envelope, or mutate provider state")
	fmt.Println("  PASS: restored original payload opened and acked after every malformed failure")
	fmt.Println("  proof_level: same-state malformed normal-message payload no-open/no-ack/no-drain proof")
	fmt.Println("  boundary: live-dev failure-path proof; not full, not release-snapshot, not adversarial harness, not replay classification, not production/security proof")
	fmt.Println("  nonclaims: not hostile-server safety, not metadata privacy, not production secure messaging, not identity verification")

	return nil
}

func (r *Runner) runMalformedPayloadSubrun(binPath string, tempRoot string, runID string, tc malformedPayloadCase) error {
	fmt.Println()
	fmt.Println("========================================")
	fmt.Println("Same-state malformed payload subrun:", tc.ID, tc.Name)
	fmt.Println("========================================")

	shortID := malformedPayloadShortID(runID + "-" + tc.ID)

	sub := relayOpenMLSJoinSubrun{
		Name:     "malformed-payload-" + tc.ID + "-" + tc.Name,
		AckAfter: true,
		RunID:    "malformed-payload-" + tc.ID + "-" + shortID,
	}
	sub.TempDir = filepath.Join(tempRoot, sub.Name)
	if err := os.Mkdir(sub.TempDir, 0700); err != nil {
		return fmt.Errorf("create malformed-payload temp dir %s: %w", sub.TempDir, err)
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
	sub.RelaySpace = "same-state-malformed-payload-dev-" + tc.ID + "-" + shortID
	sub.AliceSidecarLabel = "cs-mp-" + tc.ID + "-a-" + shortID
	sub.BobSidecarLabel = "cs-mp-" + tc.ID + "-b-" + shortID
	sub.AliceConversationLabel = "cs-mp-" + tc.ID + "-cv-" + shortID
	sub.BobConversationLabel = sub.AliceConversationLabel

	messageLabel := "mp-" + tc.ID + "-" + shortID
	messageText := "CarbonStack malformed payload profile " + tc.ID

	fmt.Println("subrun_temp_dir:", sub.TempDir)
	fmt.Println("cypher_addr:", addr)
	fmt.Println("cypher_db:", sub.DBPath)
	fmt.Println("relay_space_id:", sub.RelaySpace)
	fmt.Println("alice_sidecar_label:", sub.AliceSidecarLabel)
	fmt.Println("bob_sidecar_label:", sub.BobSidecarLabel)
	fmt.Println("conversation_label:", sub.AliceConversationLabel)
	fmt.Println("message_label:", messageLabel)
	fmt.Println("mutation:", tc.Mutation)

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
		if err := assertRelayOpenMLSTrustCandidateAbsent("before malformed payload proof", sub.AliceState, sub.BobState); err != nil {
			return err
		}

		if err := r.runMalformedPayloadJoin(&sub); err != nil {
			return err
		}

		if err := r.runMalformedPayloadSendMutateOpenRestore(&sub, tc, messageLabel, messageText); err != nil {
			return err
		}

		if err := assertRelayOpenMLSTrustCandidateAbsent("after malformed payload proof", sub.AliceState, sub.BobState); err != nil {
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

	fmt.Println("PASS: same-state malformed payload subrun", tc.ID, tc.Name)
	return nil
}

func (r *Runner) runMalformedPayloadJoin(sub *relayOpenMLSJoinSubrun) error {
	fmt.Println()
	fmt.Println("== Relay onboarding through restored same-state conversation join ==")

	keyPackageOutput, keyPackageRC, err := runWelcomeJoinFailureCommand(
		r.Comms,
		[]string{"RUST_BACKTRACE=1"},
		"go", "run", "./cmd/comms",
		"openmls-relay-keypackage-submit-dev",
		"--state", sub.BobState,
		"--relay-space", sub.RelaySpace,
		"--to-device", sub.AliceDeviceID,
		"--sidecar-device-label", sub.BobSidecarLabel,
	)
	fmt.Print(keyPackageOutput)
	if err != nil || keyPackageRC != 0 {
		return fmt.Errorf("KeyPackage submit failed rc=%d err=%v", keyPackageRC, err)
	}
	for _, needle := range []string{
		"status: sent",
		"command: openmls-relay-keypackage-submit-dev",
		"recipient_device_id: " + sub.AliceDeviceID,
		"sidecar_device_label: " + sub.BobSidecarLabel,
	} {
		if !strings.Contains(keyPackageOutput, needle) {
			return fmt.Errorf("KeyPackage submit output missing %q", needle)
		}
	}

	addMemberOutput, addMemberRC, err := runWelcomeJoinFailureCommand(
		r.Comms,
		[]string{"RUST_BACKTRACE=1"},
		"go", "run", "./cmd/comms",
		"openmls-relay-add-member-dev",
		"--state", sub.AliceState,
		"--relay-space", sub.RelaySpace,
		"--sidecar-device-label", sub.AliceSidecarLabel,
		"--conversation", sub.AliceConversationLabel,
		"--welcome-to-device", sub.BobDeviceID,
	)
	fmt.Print(addMemberOutput)
	if err != nil || addMemberRC != 0 {
		return fmt.Errorf("add-member failed rc=%d err=%v", addMemberRC, err)
	}
	for _, needle := range []string{
		"status: welcome_created_and_sent",
		"command: openmls-relay-add-member-dev",
		"welcome_recipient_device_id: " + sub.BobDeviceID,
		"sidecar_device_label: " + sub.AliceSidecarLabel,
		"sidecar_conversation_label: " + sub.AliceConversationLabel,
		"welcome_acked: false",
		"group_reloadable: true",
	} {
		if !strings.Contains(addMemberOutput, needle) {
			return fmt.Errorf("add-member output missing %q", needle)
		}
	}

	joinOutput, joinRC, err := runWelcomeJoinFailureCommand(
		r.Comms,
		[]string{"RUST_BACKTRACE=1"},
		"go", "run", "./cmd/comms",
		"openmls-relay-join-dev",
		"--state", sub.BobState,
		"--relay-space", sub.RelaySpace,
		"--sidecar-device-label", sub.BobSidecarLabel,
		"--conversation", sub.BobConversationLabel,
		"--ack-after-join",
	)
	fmt.Print(joinOutput)
	if err != nil || joinRC != 0 {
		return fmt.Errorf("Relay join failed rc=%d err=%v", joinRC, err)
	}
	for _, needle := range []string{
		"status: joined",
		"joined: true",
		"welcome_acked: true",
		"ack_delivery_state: acknowledged",
		"group_reloadable: true",
		"sidecar_conversation_label: " + sub.BobConversationLabel,
	} {
		if !strings.Contains(joinOutput, needle) {
			return fmt.Errorf("join output missing %q", needle)
		}
	}

	loadOutput, loadRC, err := runWelcomeJoinFailureCommand(
		r.Comms,
		[]string{"RUST_BACKTRACE=1"},
		"go", "run", "./cmd/comms",
		"openmls-conversation-load-check-dev",
		"--sidecar-device-label", sub.BobSidecarLabel,
		"--conversation", sub.BobConversationLabel,
	)
	fmt.Print(loadOutput)
	if err != nil || loadRC != 0 {
		return fmt.Errorf("Bob conversation load-check failed after join rc=%d err=%v", loadRC, err)
	}
	if !strings.Contains(loadOutput, "status: loaded") || !strings.Contains(loadOutput, "group_reloadable: true") {
		return fmt.Errorf("Bob load-check output missing loaded/reloadable evidence")
	}

	return nil
}

func (r *Runner) runMalformedPayloadSendMutateOpenRestore(sub *relayOpenMLSJoinSubrun, tc malformedPayloadCase, messageLabel string, messageText string) error {
	fmt.Println()
	fmt.Println("== Send normal message, mutate payload, prove no-open/no-ack/no-drain, then restore ==")

	sendOutput, sendRC, err := runWelcomeJoinFailureCommand(
		r.Comms,
		[]string{"RUST_BACKTRACE=1"},
		"go", "run", "./cmd/comms",
		"message-send-dev",
		"--state", sub.AliceState,
		"--to-device", sub.BobDeviceID,
		"--sidecar-device-label", sub.AliceSidecarLabel,
		"--conversation", sub.AliceConversationLabel,
		"--message-label", messageLabel,
		"--message", messageText,
	)
	fmt.Print(sendOutput)
	if err != nil || sendRC != 0 {
		return fmt.Errorf("message-send-dev failed rc=%d err=%v", sendRC, err)
	}
	if !strings.Contains(sendOutput, "message sent") {
		return fmt.Errorf("message-send-dev output missing message sent")
	}

	appEnvelopeID, err := sqlite3QueryOne(sub.DBPath, fmt.Sprintf(
		"SELECT envelope_id FROM envelopes WHERE recipient_device_id = %s AND content_type = 'carbonstack.mls.application-message.v0' AND delivery_state = 'queued' ORDER BY rowid DESC LIMIT 1;",
		sqlite3Quote(sub.BobDeviceID),
	))
	if err != nil {
		return err
	}
	if appEnvelopeID == "" {
		return fmt.Errorf("could not locate Bob application-message envelope")
	}
	fmt.Println("app_envelope_id:", appEnvelopeID)

	original, err := saveMalformedPayloadOriginalEnvelope(sub.DBPath, appEnvelopeID)
	if err != nil {
		return err
	}

	acksBefore, err := sqlite3Count(sub.DBPath, fmt.Sprintf(
		"SELECT COUNT(*) FROM envelope_acks WHERE envelope_id = %s;",
		sqlite3Quote(appEnvelopeID),
	))
	if err != nil {
		return err
	}
	bobInboxBefore, err := malformedPayloadBobAppInboxCount(sub.DBPath, sub.BobDeviceID)
	if err != nil {
		return err
	}
	if acksBefore != 0 {
		return fmt.Errorf("application message already acked before mutation: %d", acksBefore)
	}
	if bobInboxBefore != 1 {
		return fmt.Errorf("Bob app inbox before mutation = %d, want 1", bobInboxBefore)
	}

	providerPath := malformedPayloadConversationProviderPath(r.Comms, sub.BobSidecarLabel, sub.BobConversationLabel)
	providerStateBefore := malformedPayloadFileState(providerPath)
	envelopeStateBefore := malformedPayloadEnvelopeFingerprint(original.PayloadSHA256, original.PayloadSizeBytes)

	mutated, err := buildMalformedPayloadMutation(tc.Mutation, original)
	if err != nil {
		return err
	}
	if err := restoreMalformedPayloadEnvelope(sub.DBPath, appEnvelopeID, malformedPayloadSavedEnvelope(mutated)); err != nil {
		return err
	}

	envelopeAfterMutation, err := malformedPayloadEnvelopeState(sub.DBPath, appEnvelopeID)
	if err != nil {
		return err
	}

	fmt.Println("mutation_kind:", tc.Mutation)
	fmt.Println("original_payload:", envelopeStateBefore)
	fmt.Println("mutated_payload:", malformedPayloadEnvelopeFingerprint(mutated.PayloadSHA256, mutated.PayloadSizeBytes))
	fmt.Println("provider_state_before:", providerStateBefore)
	fmt.Println("envelope_after_mutation:", malformedPayloadEnvelopeFingerprint(envelopeAfterMutation.PayloadSHA256, envelopeAfterMutation.PayloadSizeBytes))
	fmt.Println("acks_before_bad_open:", acksBefore)
	fmt.Println("bob_app_inbox_before_bad_open:", bobInboxBefore)

	badOutput, badRC, badErr := runWelcomeJoinFailureCommand(
		r.Comms,
		[]string{"RUST_BACKTRACE=1"},
		"go", "run", "./cmd/comms",
		"message-inbox-dev",
		"--state", sub.BobState,
		"--sidecar-device-label", sub.BobSidecarLabel,
		"--conversation", sub.BobConversationLabel,
		"--message-label", messageLabel,
		"--ack",
	)
	fmt.Print(badOutput)
	fmt.Println("bad_open_rc:", badRC)
	if badErr != nil && badRC < 0 {
		return fmt.Errorf("message-inbox-dev malformed payload invocation failed non-process error rc=%d err=%v", badRC, badErr)
	}

	if strings.Contains(badOutput, "message opened") {
		return fmt.Errorf("malformed payload unexpectedly opened")
	}
	for _, forbidden := range []string{"acked: true", "ack_delivery_state: acknowledged"} {
		if strings.Contains(badOutput, forbidden) {
			return fmt.Errorf("malformed payload output contained forbidden ack marker %q", forbidden)
		}
	}
	for _, needle := range []string{
		"message open failed",
		"acked: false",
		"opened_envelopes: 0",
		"open_failures: 1",
		"ack_failures: 0",
	} {
		if !strings.Contains(badOutput, needle) {
			return fmt.Errorf("malformed payload failure output missing %q", needle)
		}
	}

	acksAfterBad, err := sqlite3Count(sub.DBPath, fmt.Sprintf(
		"SELECT COUNT(*) FROM envelope_acks WHERE envelope_id = %s;",
		sqlite3Quote(appEnvelopeID),
	))
	if err != nil {
		return err
	}
	bobInboxAfterBad, err := malformedPayloadBobAppInboxCount(sub.DBPath, sub.BobDeviceID)
	if err != nil {
		return err
	}
	providerStateAfterBad := malformedPayloadFileState(providerPath)
	envelopeAfterBad, err := malformedPayloadEnvelopeState(sub.DBPath, appEnvelopeID)
	if err != nil {
		return err
	}

	fmt.Println("acks_after_bad_open:", acksAfterBad)
	fmt.Println("bob_app_inbox_after_bad_open:", bobInboxAfterBad)
	fmt.Println("provider_state_after_bad_open:", providerStateAfterBad)
	fmt.Println("envelope_after_bad_open:", malformedPayloadEnvelopeFingerprint(envelopeAfterBad.PayloadSHA256, envelopeAfterBad.PayloadSizeBytes))

	if acksAfterBad != acksBefore {
		return fmt.Errorf("malformed payload changed ack count: before=%d after=%d", acksBefore, acksAfterBad)
	}
	if bobInboxAfterBad != bobInboxBefore {
		return fmt.Errorf("malformed payload changed Bob app inbox count: before=%d after=%d", bobInboxBefore, bobInboxAfterBad)
	}
	if providerStateAfterBad != providerStateBefore {
		return fmt.Errorf("malformed payload failed open mutated provider state: before=%s after=%s", providerStateBefore, providerStateAfterBad)
	}
	if malformedPayloadEnvelopeFingerprint(envelopeAfterBad.PayloadSHA256, envelopeAfterBad.PayloadSizeBytes) != malformedPayloadEnvelopeFingerprint(envelopeAfterMutation.PayloadSHA256, envelopeAfterMutation.PayloadSizeBytes) {
		return fmt.Errorf("malformed payload failed open rewrote envelope state")
	}

	fmt.Println("PASS: malformed payload did not open, ack, drain, mutate provider state, or rewrite envelope")

	if err := restoreMalformedPayloadEnvelope(sub.DBPath, appEnvelopeID, original); err != nil {
		return err
	}

	recoveryOutput, recoveryRC, recoveryErr := runWelcomeJoinFailureCommand(
		r.Comms,
		[]string{"RUST_BACKTRACE=1"},
		"go", "run", "./cmd/comms",
		"message-inbox-dev",
		"--state", sub.BobState,
		"--sidecar-device-label", sub.BobSidecarLabel,
		"--conversation", sub.BobConversationLabel,
		"--message-label", messageLabel,
		"--ack",
	)
	fmt.Print(recoveryOutput)
	fmt.Println("recovery_open_rc:", recoveryRC)
	if recoveryErr != nil || recoveryRC != 0 {
		return fmt.Errorf("restored payload open failed rc=%d err=%v", recoveryRC, recoveryErr)
	}
	for _, needle := range []string{
		"message opened",
		"ack_delivery_state: acknowledged",
		"acked: true",
		"opened_envelopes: 1",
		"open_failures: 0",
		"ack_failures: 0",
		"plaintext_utf8: " + messageText,
	} {
		if !strings.Contains(recoveryOutput, needle) {
			return fmt.Errorf("restored payload output missing %q", needle)
		}
	}

	acksAfterRecovery, err := sqlite3Count(sub.DBPath, fmt.Sprintf(
		"SELECT COUNT(*) FROM envelope_acks WHERE envelope_id = %s;",
		sqlite3Quote(appEnvelopeID),
	))
	if err != nil {
		return err
	}
	bobInboxAfterRecovery, err := malformedPayloadBobAppInboxCount(sub.DBPath, sub.BobDeviceID)
	if err != nil {
		return err
	}
	if acksAfterRecovery != acksBefore+1 {
		return fmt.Errorf("restored payload ack count = %d, want %d", acksAfterRecovery, acksBefore+1)
	}
	if bobInboxAfterRecovery != 0 {
		return fmt.Errorf("Bob app inbox after restored payload open = %d, want 0", bobInboxAfterRecovery)
	}

	fmt.Println("PASS: restored payload recovery opened and acked")
	return nil
}

func malformedPayloadShortID(value string) string {
	sum := sha256.Sum256([]byte(value))
	return fmt.Sprintf("%x", sum[:])[:12]
}

func malformedPayloadConversationProviderPath(commsRoot string, deviceLabel string, conversationLabel string) string {
	return filepath.Join(
		commsRoot,
		"internal",
		"protocol",
		"mls",
		"openmls-sidecar",
		".carbonstack-openmls-sidecar-state",
		"dev",
		"devices",
		deviceLabel,
		"conversations",
		conversationLabel,
		"provider-storage.json",
	)
}

func malformedPayloadFileState(path string) string {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return "absent"
		}
		return "stat_error:" + err.Error()
	}
	if info.IsDir() {
		return "dir"
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return "read_error:" + err.Error()
	}
	sum := sha256.Sum256(data)
	return fmt.Sprintf("file:%d:%x", len(data), sum[:])
}

func malformedPayloadEnvelopeFingerprint(sha string, size int) string {
	return fmt.Sprintf("%s:%d", sha, size)
}

func malformedPayloadBobAppInboxCount(dbPath string, bobDeviceID string) (int, error) {
	return sqlite3Count(dbPath, fmt.Sprintf(
		"SELECT COUNT(*) FROM envelopes WHERE recipient_device_id = %s AND content_type = 'carbonstack.mls.application-message.v0' AND delivery_state = 'queued';",
		sqlite3Quote(bobDeviceID),
	))
}

func malformedPayloadEnvelopeState(dbPath string, envelopeID string) (malformedPayloadSavedEnvelope, error) {
	return saveMalformedPayloadOriginalEnvelope(dbPath, envelopeID)
}

func saveMalformedPayloadOriginalEnvelope(dbPath string, envelopeID string) (malformedPayloadSavedEnvelope, error) {
	row, err := sqlite3QueryOne(dbPath, fmt.Sprintf(
		"SELECT ciphertext_b64 || char(31) || payload_sha256 || char(31) || payload_size_bytes FROM envelopes WHERE envelope_id = %s;",
		sqlite3Quote(envelopeID),
	))
	if err != nil {
		return malformedPayloadSavedEnvelope{}, err
	}

	parts := strings.Split(row, "\x1f")
	if len(parts) != 3 {
		return malformedPayloadSavedEnvelope{}, fmt.Errorf("unexpected saved application-message row shape: %q", row)
	}

	size, err := strconv.Atoi(strings.TrimSpace(parts[2]))
	if err != nil {
		return malformedPayloadSavedEnvelope{}, fmt.Errorf("parse saved application-message payload size %q: %w", parts[2], err)
	}

	return malformedPayloadSavedEnvelope{
		CiphertextB64:    parts[0],
		PayloadSHA256:    parts[1],
		PayloadSizeBytes: size,
	}, nil
}

func restoreMalformedPayloadEnvelope(dbPath string, envelopeID string, saved malformedPayloadSavedEnvelope) error {
	return sqlite3Exec(dbPath, fmt.Sprintf(
		"UPDATE envelopes SET ciphertext_b64 = %s, payload_sha256 = %s, payload_size_bytes = %d WHERE envelope_id = %s;",
		sqlite3Quote(saved.CiphertextB64),
		sqlite3Quote(saved.PayloadSHA256),
		saved.PayloadSizeBytes,
		sqlite3Quote(envelopeID),
	))
}

func buildMalformedPayloadMutation(mutation string, original malformedPayloadSavedEnvelope) (malformedPayloadMutation, error) {
	originalBytes, err := base64.StdEncoding.DecodeString(original.CiphertextB64)
	if err != nil {
		return malformedPayloadMutation{}, fmt.Errorf("decode original ciphertext_b64: %w", err)
	}

	var stored string
	var rawForHash []byte

	switch mutation {
	case "invalid_base64_storage_shape":
		stored = "not-valid-base64-%%%v0623b"
		rawForHash = []byte(stored)
	case "valid_base64_random_bytes":
		sum := sha256.Sum256([]byte("carbonstack-v0623b-random-bytes"))
		rawForHash = append(sum[:], 0x00, 0xff, 0x7f)
		stored = base64.StdEncoding.EncodeToString(rawForHash)
	case "valid_base64_truncated_original":
		n := len(originalBytes) / 3
		if n < 1 {
			n = 1
		}
		if n > len(originalBytes) {
			n = len(originalBytes)
		}
		rawForHash = originalBytes[:n]
		stored = base64.StdEncoding.EncodeToString(rawForHash)
	case "valid_base64_single_byte_flip":
		if len(originalBytes) == 0 {
			rawForHash = []byte{0x00}
		} else {
			mutated := append([]byte(nil), originalBytes...)
			idx := len(mutated) / 2
			mutated[idx] ^= 0x80
			rawForHash = mutated
		}
		stored = base64.StdEncoding.EncodeToString(rawForHash)
	case "valid_base64_empty_bytes":
		rawForHash = []byte{}
		stored = base64.StdEncoding.EncodeToString(rawForHash)
	case "valid_base64_original_plus_junk":
		rawForHash = append(append([]byte(nil), originalBytes...), []byte("\x00CARBONSTACK-V0623B-JUNK\xff")...)
		stored = base64.StdEncoding.EncodeToString(rawForHash)
	default:
		return malformedPayloadMutation{}, fmt.Errorf("unknown malformed payload mutation %q", mutation)
	}

	sum := sha256.Sum256(rawForHash)
	return malformedPayloadMutation{
		CiphertextB64:    stored,
		PayloadSHA256:    fmt.Sprintf("%x", sum[:]),
		PayloadSizeBytes: len(rawForHash),
	}, nil
}
