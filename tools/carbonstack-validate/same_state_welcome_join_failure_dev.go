package main

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

type welcomeJoinFailureSavedEnvelope struct {
	CiphertextB64    string
	PayloadSHA256    string
	PayloadSizeBytes int
}

func (r *Runner) SameStateWelcomeJoinFailureDev() error {
	r.PrintHeader("same-state-welcome-join-failure-dev")

	fmt.Println("status: dev/pre-alpha same-state Welcome join failure validation profile")
	fmt.Println("scope: corrupt Welcome join no-ack/no-drain/no-state-poison plus restored Welcome recovery")
	fmt.Println("proof: Relay KeyPackage -> add-member -> corrupt Welcome join failure -> no ack/drain/state poison -> restored Welcome join/ack")
	fmt.Println("boundary: live umbrella only; not full; not release-snapshot; not package-root validation; not adversarial harness; not production/security proof")
	fmt.Println("relationship: failure-path companion to same-state-integrated-dev; validates v0.6.20 atomic Welcome join state-write fix")
	fmt.Println("nonclaims: not hostile-server safety, not metadata privacy, not production secure messaging, not verified identity, not mature messenger UX")
	fmt.Println()

	if err := r.CheckRequiredPaths(); err != nil {
		return err
	}
	if err := r.CheckLiveGitUmbrella("same-state-welcome-join-failure-dev"); err != nil {
		return err
	}

	fmt.Println()
	fmt.Println("== Toolchains ==")
	_ = r.ReportTool("go", "version")
	_ = r.ReportTool("rustc", "--version")
	_ = r.ReportTool("cargo", "--version")
	_ = r.ReportTool("bash", "--version")
	_ = r.ReportTool("sqlite3", "--version")

	tempRoot, err := os.MkdirTemp("", "carbonstack-same-state-welcome-join-failure-dev-*")
	if err != nil {
		return fmt.Errorf("create same-state Welcome join failure temp root: %w", err)
	}
	defer func() {
		if err := os.RemoveAll(tempRoot); err != nil {
			fmt.Printf("WARN: remove same-state Welcome join failure temp root %s: %v\n", tempRoot, err)
		}
	}()

	binPath := filepath.Join(tempRoot, "carbonstack-cypher-same-state-welcome-join-failure")
	uniqueID := relayOpenMLSRunID()
	runID := "welcome-join-failure-" + uniqueID

	fmt.Println()
	fmt.Println("== same-state Welcome join failure generated-state root ==")
	fmt.Println("temp_root:", tempRoot)
	fmt.Println("cypher_bin:", binPath)
	fmt.Println("run_id:", runID)
	fmt.Println("note: this profile proves corrupt Welcome no-ack/no-drain/no-state-poison and restored Welcome recovery")

	fmt.Println()
	fmt.Println("== build temporary Cypher binary ==")
	if err := runLocalCypherCommand(r.Cypher, "go", "build", "-o", binPath, "./cmd/cypher"); err != nil {
		return err
	}

	r.ArtifactScan("pre-same-state-welcome-join-failure-dev")

	if err := r.runSameStateWelcomeJoinFailureSubrun(binPath, tempRoot, relayOpenMLSJoinSubrun{
		Name:     "welcome-join-failure-no-state-poison",
		AckAfter: true,
		RunID:    runID,
	}); err != nil {
		return err
	}

	r.ArtifactScan("post-same-state-welcome-join-failure-dev")

	fmt.Println()
	fmt.Println("same-state-welcome-join-failure-dev profile result:")
	fmt.Println("  PASS: Relay onboarding reached KeyPackage -> add-member -> queued Welcome before Bob join")
	fmt.Println("  PASS: corrupt Welcome join failed without join/ack success markers")
	fmt.Println("  PASS: corrupt Welcome join did not ack")
	fmt.Println("  PASS: corrupt Welcome join did not drain Bob Relay inbox")
	fmt.Println("  PASS: corrupt Welcome join left no final or staging Bob conversation state")
	fmt.Println("  PASS: Bob conversation load-check failed after corrupt Welcome join")
	fmt.Println("  PASS: restored valid Welcome joined with the same conversation label")
	fmt.Println("  PASS: restored valid Welcome acked only after successful join")
	fmt.Println("  PASS: Bob conversation was reloadable after restored join")
	fmt.Println("  proof_level: same-state Welcome join failure no-ack/no-drain/no-state-poison proof")
	fmt.Println("  boundary: live-dev failure-path proof; not full, not release-snapshot, not adversarial harness, not identity verification, not production/security proof")
	fmt.Println("  relationship: same-state-integrated-dev remains positive-path; v0.6.20 sidecar atomic state fix is now live-validated")
	fmt.Println("  nonclaims: not hostile-server safety, not metadata privacy, not production secure messaging, not identity verification")

	return nil
}

func (r *Runner) runSameStateWelcomeJoinFailureSubrun(binPath string, tempRoot string, sub relayOpenMLSJoinSubrun) error {
	fmt.Println()
	fmt.Println("========================================")
	fmt.Println("Same-state Welcome join failure subrun:", sub.Name)
	fmt.Println("========================================")

	sub.TempDir = filepath.Join(tempRoot, sub.Name)
	if err := os.Mkdir(sub.TempDir, 0700); err != nil {
		return fmt.Errorf("create same-state Welcome join failure subrun temp dir %s: %w", sub.TempDir, err)
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
	sub.RelaySpace = "same-state-welcome-join-failure-dev-" + sub.RunID
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

		if err := assertRelayOpenMLSTrustCandidateAbsent("before same-state Welcome join failure proof", sub.AliceState, sub.BobState); err != nil {
			return err
		}

		if err := r.runWelcomeJoinFailureManualKeyPackageAndAddMember(&sub); err != nil {
			return err
		}

		if err := r.runWelcomeJoinFailureCorruptThenRestore(&sub); err != nil {
			return err
		}

		if err := assertRelayOpenMLSTrustCandidateAbsent("after same-state Welcome join failure proof", sub.AliceState, sub.BobState); err != nil {
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

	fmt.Println("PASS: same-state Welcome join failure subrun", sub.Name)
	return nil
}

func (r *Runner) runWelcomeJoinFailureManualKeyPackageAndAddMember(sub *relayOpenMLSJoinSubrun) error {
	fmt.Println()
	fmt.Println("== Manual Relay onboarding through add-member, stopping before Bob join ==")

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
		"content_type: carbonstack.mls.keypackage.v0",
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

	if err := assertWelcomeJoinFailureNoFinalOrStagingState(r.Comms, sub.BobSidecarLabel, sub.BobConversationLabel, "before Bob join"); err != nil {
		return err
	}

	return nil
}

func (r *Runner) runWelcomeJoinFailureCorruptThenRestore(sub *relayOpenMLSJoinSubrun) error {
	fmt.Println()
	fmt.Println("== Locate, save, and corrupt queued Welcome ==")

	welcomeEnvelopeID, err := sqlite3QueryOne(sub.DBPath, fmt.Sprintf(
		"SELECT envelope_id FROM envelopes WHERE recipient_device_id = %s AND relay_space_id = %s AND content_type = 'carbonstack.mls.welcome.v0' ORDER BY rowid DESC LIMIT 1;",
		sqlite3Quote(sub.BobDeviceID),
		sqlite3Quote(sub.RelaySpace),
	))
	if err != nil {
		return err
	}
	if welcomeEnvelopeID == "" {
		return fmt.Errorf("could not locate queued Welcome envelope for Bob")
	}
	fmt.Println("welcome_envelope_id:", welcomeEnvelopeID)

	original, err := saveWelcomeJoinFailureOriginalEnvelope(sub.DBPath, welcomeEnvelopeID)
	if err != nil {
		return err
	}

	badPayload := []byte("not-a-valid-openmls-welcome-v0621b")
	badB64 := base64.StdEncoding.EncodeToString(badPayload)
	badHash := sha256.Sum256(badPayload)
	badSHA := fmt.Sprintf("%x", badHash[:])

	if err := sqlite3Exec(sub.DBPath, fmt.Sprintf(
		"UPDATE envelopes SET ciphertext_b64 = %s, payload_sha256 = %s, payload_size_bytes = %d WHERE envelope_id = %s;",
		sqlite3Quote(badB64),
		sqlite3Quote(badSHA),
		len(badPayload),
		sqlite3Quote(welcomeEnvelopeID),
	)); err != nil {
		return err
	}

	acksBeforeBadJoin, err := sqlite3Count(sub.DBPath, "SELECT COUNT(*) FROM envelope_acks;")
	if err != nil {
		return err
	}
	welcomeAcksBeforeBadJoin, err := sqlite3Count(sub.DBPath, fmt.Sprintf(
		"SELECT COUNT(*) FROM envelope_acks WHERE envelope_id = %s;",
		sqlite3Quote(welcomeEnvelopeID),
	))
	if err != nil {
		return err
	}
	bobRelayInboxBeforeBadJoin, err := sqlite3Count(sub.DBPath, fmt.Sprintf(
		"SELECT COUNT(*) FROM envelopes WHERE relay_space_id = %s AND recipient_device_id = %s AND delivery_state = 'queued';",
		sqlite3Quote(sub.RelaySpace),
		sqlite3Quote(sub.BobDeviceID),
	))
	if err != nil {
		return err
	}

	fmt.Println("acks_before_bad_join:", acksBeforeBadJoin)
	fmt.Println("welcome_acks_before_bad_join:", welcomeAcksBeforeBadJoin)
	fmt.Println("bob_relay_inbox_before_bad_join:", bobRelayInboxBeforeBadJoin)

	if welcomeAcksBeforeBadJoin != 0 {
		return fmt.Errorf("Welcome already has ack rows before corrupt join: %d", welcomeAcksBeforeBadJoin)
	}
	if bobRelayInboxBeforeBadJoin != 1 {
		return fmt.Errorf("Bob Relay inbox before corrupt join = %d, want 1", bobRelayInboxBeforeBadJoin)
	}

	fmt.Println()
	fmt.Println("== Case A: corrupt Welcome join must fail without ack/drain/state poison ==")

	badJoinOutput, badJoinRC, badJoinErr := runWelcomeJoinFailureCommand(
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
	fmt.Print(badJoinOutput)
	fmt.Println("bad_join_rc:", badJoinRC)
	if badJoinErr == nil || badJoinRC == 0 {
		return fmt.Errorf("corrupt Welcome join unexpectedly succeeded")
	}

	if !strings.Contains(badJoinOutput, "welcome_invalid") {
		return fmt.Errorf("corrupt Welcome join output missing welcome_invalid")
	}
	if err := assertWelcomeJoinFailureNoSuccessMarkers(badJoinOutput, "corrupt Welcome join"); err != nil {
		return err
	}

	acksAfterBadJoin, err := sqlite3Count(sub.DBPath, "SELECT COUNT(*) FROM envelope_acks;")
	if err != nil {
		return err
	}
	welcomeAcksAfterBadJoin, err := sqlite3Count(sub.DBPath, fmt.Sprintf(
		"SELECT COUNT(*) FROM envelope_acks WHERE envelope_id = %s;",
		sqlite3Quote(welcomeEnvelopeID),
	))
	if err != nil {
		return err
	}
	bobRelayInboxAfterBadJoin, err := sqlite3Count(sub.DBPath, fmt.Sprintf(
		"SELECT COUNT(*) FROM envelopes WHERE relay_space_id = %s AND recipient_device_id = %s AND delivery_state = 'queued';",
		sqlite3Quote(sub.RelaySpace),
		sqlite3Quote(sub.BobDeviceID),
	))
	if err != nil {
		return err
	}

	fmt.Println("acks_after_bad_join:", acksAfterBadJoin)
	fmt.Println("welcome_acks_after_bad_join:", welcomeAcksAfterBadJoin)
	fmt.Println("bob_relay_inbox_after_bad_join:", bobRelayInboxAfterBadJoin)

	if acksAfterBadJoin != acksBeforeBadJoin {
		return fmt.Errorf("corrupt Welcome join changed total ack count: before=%d after=%d", acksBeforeBadJoin, acksAfterBadJoin)
	}
	if welcomeAcksAfterBadJoin != welcomeAcksBeforeBadJoin {
		return fmt.Errorf("corrupt Welcome join changed Welcome ack count: before=%d after=%d", welcomeAcksBeforeBadJoin, welcomeAcksAfterBadJoin)
	}
	if bobRelayInboxAfterBadJoin != bobRelayInboxBeforeBadJoin {
		return fmt.Errorf("corrupt Welcome join changed Bob Relay inbox count: before=%d after=%d", bobRelayInboxBeforeBadJoin, bobRelayInboxAfterBadJoin)
	}

	if err := assertWelcomeJoinFailureNoFinalOrStagingState(r.Comms, sub.BobSidecarLabel, sub.BobConversationLabel, "corrupt Welcome join"); err != nil {
		return err
	}

	loadOutput, loadRC, loadErr := runWelcomeJoinFailureCommand(
		r.Comms,
		[]string{"RUST_BACKTRACE=1"},
		"go", "run", "./cmd/comms",
		"openmls-conversation-load-check-dev",
		"--sidecar-device-label", sub.BobSidecarLabel,
		"--conversation", sub.BobConversationLabel,
	)
	fmt.Print(loadOutput)
	fmt.Println("bob_load_after_bad_rc:", loadRC)
	if loadErr == nil || loadRC == 0 {
		return fmt.Errorf("Bob conversation load-check unexpectedly succeeded after corrupt Welcome join")
	}

	fmt.Println()
	fmt.Println("== Restore valid Welcome and prove recovery ==")

	if err := restoreWelcomeJoinFailureOriginalEnvelope(sub.DBPath, welcomeEnvelopeID, original); err != nil {
		return err
	}

	goodJoinOutput, goodJoinRC, goodJoinErr := runWelcomeJoinFailureCommand(
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
	fmt.Print(goodJoinOutput)
	fmt.Println("good_join_rc:", goodJoinRC)
	if goodJoinErr != nil || goodJoinRC != 0 {
		return fmt.Errorf("restored Welcome join failed rc=%d err=%v", goodJoinRC, goodJoinErr)
	}

	for _, needle := range []string{
		"status: joined",
		"joined: true",
		"welcome_acked: true",
		"ack_delivery_state: acknowledged",
		"group_reloadable: true",
		"sidecar_conversation_label: " + sub.BobConversationLabel,
	} {
		if !strings.Contains(goodJoinOutput, needle) {
			return fmt.Errorf("restored Welcome join output missing %q", needle)
		}
	}

	acksAfterGoodJoin, err := sqlite3Count(sub.DBPath, "SELECT COUNT(*) FROM envelope_acks;")
	if err != nil {
		return err
	}
	welcomeAcksAfterGoodJoin, err := sqlite3Count(sub.DBPath, fmt.Sprintf(
		"SELECT COUNT(*) FROM envelope_acks WHERE envelope_id = %s;",
		sqlite3Quote(welcomeEnvelopeID),
	))
	if err != nil {
		return err
	}
	bobRelayInboxAfterGoodJoin, err := sqlite3Count(sub.DBPath, fmt.Sprintf(
		"SELECT COUNT(*) FROM envelopes WHERE relay_space_id = %s AND recipient_device_id = %s AND delivery_state = 'queued';",
		sqlite3Quote(sub.RelaySpace),
		sqlite3Quote(sub.BobDeviceID),
	))
	if err != nil {
		return err
	}

	fmt.Println("acks_after_good_join:", acksAfterGoodJoin)
	fmt.Println("welcome_acks_after_good_join:", welcomeAcksAfterGoodJoin)
	fmt.Println("bob_relay_inbox_after_good_join:", bobRelayInboxAfterGoodJoin)

	if acksAfterGoodJoin != acksBeforeBadJoin+1 {
		return fmt.Errorf("restored Welcome join ack count = %d, want %d", acksAfterGoodJoin, acksBeforeBadJoin+1)
	}
	if welcomeAcksAfterGoodJoin != 1 {
		return fmt.Errorf("Welcome ack rows after restored join = %d, want 1", welcomeAcksAfterGoodJoin)
	}
	if bobRelayInboxAfterGoodJoin != 0 {
		return fmt.Errorf("Bob Relay inbox after restored join = %d, want 0", bobRelayInboxAfterGoodJoin)
	}

	if err := assertWelcomeJoinFailureFinalStatePresent(r.Comms, sub.BobSidecarLabel, sub.BobConversationLabel); err != nil {
		return err
	}

	reloadOutput, reloadRC, reloadErr := runWelcomeJoinFailureCommand(
		r.Comms,
		[]string{"RUST_BACKTRACE=1"},
		"go", "run", "./cmd/comms",
		"openmls-conversation-load-check-dev",
		"--sidecar-device-label", sub.BobSidecarLabel,
		"--conversation", sub.BobConversationLabel,
	)
	fmt.Print(reloadOutput)
	if reloadErr != nil || reloadRC != 0 {
		return fmt.Errorf("Bob conversation load-check after restored join failed rc=%d err=%v", reloadRC, reloadErr)
	}
	if !strings.Contains(reloadOutput, "status: loaded") || !strings.Contains(reloadOutput, "group_reloadable: true") {
		return fmt.Errorf("Bob conversation load-check after restored join missing loaded/reloadable evidence")
	}

	fmt.Println()
	fmt.Println("fixed_welcome_join_failure_result:")
	fmt.Println("  bad_join_rc:", badJoinRC)
	fmt.Println("  good_join_rc:", goodJoinRC)
	fmt.Println("  bob_load_after_bad_rc:", loadRC)
	fmt.Println("  final_state_after_bad_join: absent")
	fmt.Println("  staging_state_after_bad_join: absent")
	fmt.Println("  final_state_after_restored_join: present")
	fmt.Println("  restored_join_recovered_without_manual_cleanup: yes")
	fmt.Println("  corrupt_welcome_no_ack: true")
	fmt.Println("  corrupt_welcome_no_drain: true")
	fmt.Println("  restored_welcome_ack_after_join: true")

	return nil
}

func runWelcomeJoinFailureCommand(dir string, env []string, args ...string) (string, int, error) {
	if len(args) == 0 {
		return "", -1, fmt.Errorf("missing command")
	}

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Dir = dir
	cmd.Env = append(os.Environ(), env...)

	outputBytes, err := cmd.CombinedOutput()
	output := string(outputBytes)

	if err == nil {
		return output, 0, nil
	}

	if exitErr, ok := err.(*exec.ExitError); ok {
		return output, exitErr.ExitCode(), err
	}

	return output, -1, err
}

func assertWelcomeJoinFailureNoSuccessMarkers(output string, label string) error {
	for _, forbidden := range []string{
		"status: joined",
		"joined: true",
		"welcome_acked: true",
		"ack_delivery_state: acknowledged",
	} {
		if strings.Contains(output, forbidden) {
			return fmt.Errorf("%s output contained forbidden success marker %q", label, forbidden)
		}
	}
	return nil
}

func welcomeJoinFailureConversationDir(commsRoot string, deviceLabel string, conversationLabel string) string {
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
	)
}

func assertWelcomeJoinFailureNoFinalOrStagingState(commsRoot string, deviceLabel string, conversationLabel string, label string) error {
	finalDir := welcomeJoinFailureConversationDir(commsRoot, deviceLabel, conversationLabel)
	if _, err := os.Stat(finalDir); err == nil {
		return fmt.Errorf("final Bob conversation state exists after %s: %s", label, finalDir)
	} else if !os.IsNotExist(err) {
		return fmt.Errorf("stat final Bob conversation state after %s: %w", label, err)
	}

	parent := filepath.Dir(finalDir)
	matches, err := filepath.Glob(filepath.Join(parent, ".join-staging-"+conversationLabel+"-*"))
	if err != nil {
		return fmt.Errorf("glob staging state after %s: %w", label, err)
	}
	if len(matches) > 0 {
		return fmt.Errorf("staging Bob conversation state exists after %s: %v", label, matches)
	}

	fmt.Println("PASS: no final/staging Bob conversation state after", label)
	return nil
}

func assertWelcomeJoinFailureFinalStatePresent(commsRoot string, deviceLabel string, conversationLabel string) error {
	finalDir := welcomeJoinFailureConversationDir(commsRoot, deviceLabel, conversationLabel)
	info, err := os.Stat(finalDir)
	if err != nil {
		return fmt.Errorf("final Bob conversation state missing after restored join: %w", err)
	}
	if !info.IsDir() {
		return fmt.Errorf("final Bob conversation state is not a directory: %s", finalDir)
	}

	for _, file := range []string{
		"conversation-summary.json",
		"join-summary.json",
		"provider-storage.json",
	} {
		path := filepath.Join(finalDir, file)
		info, err := os.Stat(path)
		if err != nil {
			return fmt.Errorf("joined-state file missing after restored join %s: %w", path, err)
		}
		if info.IsDir() {
			return fmt.Errorf("joined-state path is directory, want file: %s", path)
		}
	}

	fmt.Println("PASS: final Bob conversation state present after restored valid join")
	return nil
}

func sqlite3Quote(value string) string {
	return "'" + strings.ReplaceAll(value, "'", "''") + "'"
}

func sqlite3QueryOne(dbPath string, query string) (string, error) {
	cmd := exec.Command("sqlite3", dbPath, query)
	outputBytes, err := cmd.CombinedOutput()
	output := strings.TrimSpace(string(outputBytes))
	if err != nil {
		return "", fmt.Errorf("sqlite3 query failed: %w\nquery: %s\noutput: %s", err, query, output)
	}
	return output, nil
}

func sqlite3Exec(dbPath string, query string) error {
	cmd := exec.Command("sqlite3", dbPath, query)
	outputBytes, err := cmd.CombinedOutput()
	output := strings.TrimSpace(string(outputBytes))
	if err != nil {
		return fmt.Errorf("sqlite3 exec failed: %w\nquery: %s\noutput: %s", err, query, output)
	}
	return nil
}

func sqlite3Count(dbPath string, query string) (int, error) {
	value, err := sqlite3QueryOne(dbPath, query)
	if err != nil {
		return 0, err
	}
	n, err := strconv.Atoi(strings.TrimSpace(value))
	if err != nil {
		return 0, fmt.Errorf("parse sqlite3 count %q: %w", value, err)
	}
	return n, nil
}

func saveWelcomeJoinFailureOriginalEnvelope(dbPath string, envelopeID string) (welcomeJoinFailureSavedEnvelope, error) {
	row, err := sqlite3QueryOne(dbPath, fmt.Sprintf(
		"SELECT ciphertext_b64 || char(31) || payload_sha256 || char(31) || payload_size_bytes FROM envelopes WHERE envelope_id = %s;",
		sqlite3Quote(envelopeID),
	))
	if err != nil {
		return welcomeJoinFailureSavedEnvelope{}, err
	}

	parts := strings.Split(row, "\x1f")
	if len(parts) != 3 {
		return welcomeJoinFailureSavedEnvelope{}, fmt.Errorf("unexpected saved Welcome row shape: %q", row)
	}

	size, err := strconv.Atoi(strings.TrimSpace(parts[2]))
	if err != nil {
		return welcomeJoinFailureSavedEnvelope{}, fmt.Errorf("parse saved Welcome payload size %q: %w", parts[2], err)
	}

	return welcomeJoinFailureSavedEnvelope{
		CiphertextB64:    parts[0],
		PayloadSHA256:    parts[1],
		PayloadSizeBytes: size,
	}, nil
}

func restoreWelcomeJoinFailureOriginalEnvelope(dbPath string, envelopeID string, saved welcomeJoinFailureSavedEnvelope) error {
	return sqlite3Exec(dbPath, fmt.Sprintf(
		"UPDATE envelopes SET ciphertext_b64 = %s, payload_sha256 = %s, payload_size_bytes = %d WHERE envelope_id = %s;",
		sqlite3Quote(saved.CiphertextB64),
		sqlite3Quote(saved.PayloadSHA256),
		saved.PayloadSizeBytes,
		sqlite3Quote(envelopeID),
	))
}
