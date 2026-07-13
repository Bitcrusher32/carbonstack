package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func (r *Runner) RelaySpaceMemberStateDev() error {
	r.PrintHeader("relay-space-member-state-dev")

	fmt.Println(
		"status: dev/pre-alpha Relay Space routing-member state profile",
	)
	fmt.Println(
		"scope: Comms operator command -> disable -> idempotent disable -> " +
			"reactivate -> leave -> idempotent leave -> rejoin refusal",
	)
	fmt.Println(
		"boundary: routing-state administration only; not authenticated " +
			"administration, production authorization, identity verification, " +
			"trust promotion, OpenMLS membership, member deletion, or rejoin",
	)
	fmt.Println(
		"state boundary: the command reads an explicit --state path for " +
			"server context and must not rewrite operator or member state files",
	)

	if err := r.CheckRequiredPaths(); err != nil {
		return err
	}
	if err := r.CheckLiveGitUmbrella(
		"relay-space-member-state-dev",
	); err != nil {
		return err
	}

	tempRoot, err := os.MkdirTemp(
		"",
		"carbonstack-relay-space-member-state-dev-*",
	)
	if err != nil {
		return fmt.Errorf("create member-state temp root: %w", err)
	}
	defer func() {
		if err := os.RemoveAll(tempRoot); err != nil {
			fmt.Printf(
				"WARN: remove member-state temp root %s: %v\n",
				tempRoot,
				err,
			)
		}
	}()

	binPath := filepath.Join(tempRoot, "carbonstack-cypher-member-state")
	dbPath := filepath.Join(tempRoot, "cypher.db")
	aliceStatePath := filepath.Join(tempRoot, "alice-state.json")
	bobStatePath := filepath.Join(tempRoot, "bob-state.json")

	port, err := reserveLoopbackPort()
	if err != nil {
		return err
	}
	baseURL := "http://127.0.0.1:" + strconv.Itoa(port)

	fmt.Println()
	fmt.Println("== member-state generated-state root ==")
	fmt.Println("temp_root:", tempRoot)
	fmt.Println("cypher_db:", dbPath)
	fmt.Println("cypher_url:", baseURL)

	if err := runLocalCypherCommand(
		r.Cypher,
		"go",
		"build",
		"-o",
		binPath,
		"./cmd/cypher",
	); err != nil {
		return err
	}

	env := append(
		os.Environ(),
		"CYPHER_ADDR=127.0.0.1:"+strconv.Itoa(port),
		"CYPHER_DB="+dbPath,
		"CYPHER_MIGRATIONS="+filepath.Join(r.Cypher, "migrations"),
	)

	server, err := startLocalCypherServer(binPath, r.Cypher, env)
	if err != nil {
		return err
	}
	serverRunning := true
	defer func() {
		if serverRunning {
			_ = server.stop("member-state")
		}
	}()

	if err := waitForLocalCypherHealth(baseURL + "/v0/health"); err != nil {
		return err
	}

	runID := strconv.FormatInt(time.Now().UnixNano(), 10)
	aliceInvite := "b4d2-alice-" + runID
	bobInvite := "b4d2-bob-" + runID

	for _, setup := range [][]string{
		{
			"init",
			"--state",
			aliceStatePath,
			"--server",
			baseURL,
		},
		{
			"dev-create-invite",
			"--state",
			aliceStatePath,
			"--server",
			baseURL,
			"--invite",
			aliceInvite,
		},
		{
			"claim-invite",
			"--state",
			aliceStatePath,
			"--server",
			baseURL,
			"--invite",
			aliceInvite,
			"--name",
			"Alice B4d2",
		},
		{
			"register-device",
			"--state",
			aliceStatePath,
			"--label",
			"alice-b4d2-" + runID,
		},
		{
			"init",
			"--state",
			bobStatePath,
			"--server",
			baseURL,
		},
		{
			"dev-create-invite",
			"--state",
			bobStatePath,
			"--server",
			baseURL,
			"--invite",
			bobInvite,
		},
		{
			"claim-invite",
			"--state",
			bobStatePath,
			"--server",
			baseURL,
			"--invite",
			bobInvite,
			"--name",
			"Bob B4d2",
		},
		{
			"register-device",
			"--state",
			bobStatePath,
			"--label",
			"bob-b4d2-" + runID,
		},
	} {
		if err := runRelaySpaceInviteClaimSetupCommand(
			r.Comms,
			setup...,
		); err != nil {
			return err
		}
	}

	aliceState, err := readRelaySpaceInviteClaimState(aliceStatePath)
	if err != nil {
		return err
	}
	bobState, err := readRelaySpaceInviteClaimState(bobStatePath)
	if err != nil {
		return err
	}

	relaySpaceID := "relay-space-b4d2-" + runID
	_, err = localCypherPOST(
		baseURL+"/v0/relay-spaces",
		map[string]any{
			"relay_space_id":        relaySpaceID,
			"display_label":         "B4d2 member-state space",
			"created_by_account_id": aliceState.AccountID,
			"created_by_device_id":  aliceState.DeviceID,
		},
		201,
	)
	if err != nil {
		return err
	}

	aliceMemberID := "alice-member-b4d2-" + runID
	bobMemberID := "bob-member-b4d2-" + runID

	_, err = localCypherPOST(
		baseURL+"/v0/relay-spaces/"+relaySpaceID+"/members",
		map[string]any{
			"routing_member_id": aliceMemberID,
			"account_id":        aliceState.AccountID,
			"device_id":         aliceState.DeviceID,
			"display_label":     "Alice B4d2 operator",
		},
		201,
	)
	if err != nil {
		return err
	}
	_, err = localCypherPOST(
		baseURL+"/v0/relay-spaces/"+relaySpaceID+"/members",
		map[string]any{
			"routing_member_id": bobMemberID,
			"account_id":        bobState.AccountID,
			"device_id":         bobState.DeviceID,
			"display_label":     "Bob B4d2 member",
		},
		201,
	)
	if err != nil {
		return err
	}
	fmt.Println("PASS: created Relay Space and two routing members")

	aliceBefore, err := os.ReadFile(aliceStatePath)
	if err != nil {
		return fmt.Errorf("read Alice state before: %w", err)
	}
	bobBefore, err := os.ReadFile(bobStatePath)
	if err != nil {
		return fmt.Errorf("read Bob state before: %w", err)
	}

	disableOutput, err := runRelaySpaceMemberStateCommand(
		r.Comms,
		aliceStatePath,
		relaySpaceID,
		bobMemberID,
		"disabled",
	)
	if err != nil {
		return err
	}
	for _, expected := range []string{
		"transition_classification: transitioned",
		"idempotent: false",
		"previous_state: active",
		"current_state: disabled",
		"relay_space_id: " + relaySpaceID,
		"routing_member_id: " + bobMemberID,
		"account_id: " + bobState.AccountID,
		"device_id: " + bobState.DeviceID,
		"local_state_mutated: false",
		"not authenticated administration",
		"not OpenMLS group membership",
	} {
		if !strings.Contains(disableOutput, expected) {
			return fmt.Errorf(
				"disable output missing %q:\n%s",
				expected,
				disableOutput,
			)
		}
	}
	fmt.Println("PASS: Comms command disabled Bob")

	status, body, err := relaySpaceMemberStatePOST(
		baseURL+"/v0/relay-spaces/"+relaySpaceID+"/envelopes",
		map[string]any{
			"sender_device_id":    aliceState.DeviceID,
			"recipient_device_id": bobState.DeviceID,
			"content_type":        "carbonstack.message.text.stub.v0",
			"protocol_version":    "stub-v0",
			"ciphertext_b64":      "AQID",
		},
	)
	if err != nil {
		return err
	}
	if status != http.StatusForbidden ||
		!strings.Contains(body, "recipient_not_relay_member") {
		return fmt.Errorf(
			"disabled recipient routing result = %d %s",
			status,
			body,
		)
	}
	fmt.Println("PASS: disabled Bob lost recipient routing authority")

	idempotentDisable, err := runRelaySpaceMemberStateCommand(
		r.Comms,
		aliceStatePath,
		relaySpaceID,
		bobMemberID,
		"disabled",
	)
	if err != nil {
		return err
	}
	for _, expected := range []string{
		"transition_classification: already_in_state",
		"idempotent: true",
		"previous_state: disabled",
		"current_state: disabled",
	} {
		if !strings.Contains(idempotentDisable, expected) {
			return fmt.Errorf(
				"idempotent disable output missing %q:\n%s",
				expected,
				idempotentDisable,
			)
		}
	}
	fmt.Println("PASS: repeated disable was idempotent")

	reactivateOutput, err := runRelaySpaceMemberStateCommand(
		r.Comms,
		aliceStatePath,
		relaySpaceID,
		bobMemberID,
		"active",
	)
	if err != nil {
		return err
	}
	for _, expected := range []string{
		"transition_classification: transitioned",
		"previous_state: disabled",
		"current_state: active",
		"\ndisabled_at: \n",
	} {
		if !strings.Contains(reactivateOutput, expected) {
			return fmt.Errorf(
				"reactivation output missing %q:\n%s",
				expected,
				reactivateOutput,
			)
		}
	}

	status, body, err = relaySpaceMemberStatePOST(
		baseURL+"/v0/relay-spaces/"+relaySpaceID+"/envelopes",
		map[string]any{
			"sender_device_id":    aliceState.DeviceID,
			"recipient_device_id": bobState.DeviceID,
			"content_type":        "carbonstack.message.text.stub.v0",
			"protocol_version":    "stub-v0",
			"ciphertext_b64":      "AQID",
		},
	)
	if err != nil {
		return err
	}
	if status != http.StatusCreated {
		return fmt.Errorf(
			"reactivated routing result = %d %s",
			status,
			body,
		)
	}
	fmt.Println("PASS: explicit reactivation restored routing authority")

	leaveOutput, err := runRelaySpaceMemberStateCommand(
		r.Comms,
		aliceStatePath,
		relaySpaceID,
		bobMemberID,
		"left",
	)
	if err != nil {
		return err
	}
	for _, expected := range []string{
		"transition_classification: transitioned",
		"previous_state: active",
		"current_state: left",
	} {
		if !strings.Contains(leaveOutput, expected) {
			return fmt.Errorf(
				"leave output missing %q:\n%s",
				expected,
				leaveOutput,
			)
		}
	}

	status, body, err = relaySpaceMemberStatePOST(
		baseURL+"/v0/relay-spaces/"+relaySpaceID+"/envelopes",
		map[string]any{
			"sender_device_id":    aliceState.DeviceID,
			"recipient_device_id": bobState.DeviceID,
			"content_type":        "carbonstack.message.text.stub.v0",
			"protocol_version":    "stub-v0",
			"ciphertext_b64":      "AQID",
		},
	)
	if err != nil {
		return err
	}
	if status != http.StatusForbidden ||
		!strings.Contains(body, "recipient_not_relay_member") {
		return fmt.Errorf(
			"left recipient routing result = %d %s",
			status,
			body,
		)
	}
	fmt.Println("PASS: left Bob lost routing authority")

	idempotentLeave, err := runRelaySpaceMemberStateCommand(
		r.Comms,
		aliceStatePath,
		relaySpaceID,
		bobMemberID,
		"left",
	)
	if err != nil {
		return err
	}
	if !strings.Contains(
		idempotentLeave,
		"transition_classification: already_in_state",
	) || !strings.Contains(idempotentLeave, "idempotent: true") {
		return fmt.Errorf(
			"idempotent leave output unexpected:\n%s",
			idempotentLeave,
		)
	}
	fmt.Println("PASS: repeated leave was idempotent")

	rejoinOutput, rejoinErr := runRelaySpaceMemberStateCommand(
		r.Comms,
		aliceStatePath,
		relaySpaceID,
		bobMemberID,
		"active",
	)
	if rejoinErr == nil {
		return fmt.Errorf(
			"left-to-active unexpectedly succeeded:\n%s",
			rejoinOutput,
		)
	}
	if !strings.Contains(
		rejoinOutput,
		"relay_space_member_rejoin_required",
	) && !strings.Contains(
		rejoinOutput,
		"explicit rejoin workflow",
	) {
		return fmt.Errorf(
			"rejoin refusal output lacked explicit classification:\n%s",
			rejoinOutput,
		)
	}
	fmt.Println("PASS: left-to-active refused explicit rejoin boundary")

	aliceAfter, err := os.ReadFile(aliceStatePath)
	if err != nil {
		return fmt.Errorf("read Alice state after: %w", err)
	}
	bobAfter, err := os.ReadFile(bobStatePath)
	if err != nil {
		return fmt.Errorf("read Bob state after: %w", err)
	}
	if !bytes.Equal(aliceBefore, aliceAfter) {
		return fmt.Errorf("operator command rewrote Alice state")
	}
	if !bytes.Equal(bobBefore, bobAfter) {
		return fmt.Errorf("operator command rewrote Bob state")
	}
	fmt.Println("PASS: operator and member local state remained unchanged")

	memberState, err := relaySpaceInviteClaimSQLiteScalar(
		dbPath,
		"SELECT state FROM relay_space_members WHERE routing_member_id = "+
			relaySpaceInviteClaimSQLQuote(bobMemberID)+";",
	)
	if err != nil {
		return err
	}
	if memberState != "left" {
		return fmt.Errorf("Bob final state = %q, want left", memberState)
	}

	disabledAt, err := relaySpaceInviteClaimSQLiteScalar(
		dbPath,
		"SELECT COALESCE(disabled_at, '') FROM relay_space_members "+
			"WHERE routing_member_id = "+
			relaySpaceInviteClaimSQLQuote(bobMemberID)+";",
	)
	if err != nil {
		return err
	}
	if disabledAt != "" {
		return fmt.Errorf(
			"Bob left disabled_at = %q, want empty",
			disabledAt,
		)
	}

	envelopeCount, err := relaySpaceInviteClaimSQLiteScalar(
		dbPath,
		"SELECT COUNT(*) FROM envelopes;",
	)
	if err != nil {
		return err
	}
	if envelopeCount != "1" {
		return fmt.Errorf(
			"successful envelope count = %q, want 1",
			envelopeCount,
		)
	}
	fmt.Println("PASS: final DB state is left with one successful envelope")

	for _, name := range []string{
		"trust.json",
		"trust-events.jsonl",
		"identity-candidates.json",
	} {
		path := filepath.Join(tempRoot, name)
		if _, err := os.Stat(path); err == nil {
			return fmt.Errorf(
				"unexpected trust/candidate mutation: %s",
				path,
			)
		} else if !os.IsNotExist(err) {
			return fmt.Errorf("inspect %s: %w", path, err)
		}
		fmt.Println("ABSENT:", path)
	}
	fmt.Println("PASS: member-state path did not create trust/candidate state")

	if err := server.stop("member-state"); err != nil {
		return err
	}
	serverRunning = false

	fmt.Println()
	fmt.Println("relay-space-member-state-dev profile result:")
	fmt.Println("  disable: transitioned")
	fmt.Println("  repeated_disable: already_in_state")
	fmt.Println("  reactivation: transitioned")
	fmt.Println("  leave: transitioned")
	fmt.Println("  repeated_leave: already_in_state")
	fmt.Println("  left_to_active: rejoin_required")
	fmt.Println("  disabled_routing_authorized: false")
	fmt.Println("  reactivated_routing_authorized: true")
	fmt.Println("  left_routing_authorized: false")
	fmt.Println("  operator_state_mutated: false")
	fmt.Println("  member_state_file_mutated: false")
	fmt.Println("  trust_or_candidate_state_mutated: false")
	fmt.Println("  final_member_state: left")
	fmt.Println("  final_disabled_at_present: false")
	fmt.Println("  successful_envelopes: 1")
	fmt.Println(
		"  boundary: dev routing-state operator proof only; no authenticated " +
			"administration, production authorization, trust, OpenMLS, " +
			"member deletion, or rejoin claim",
	)

	return nil
}

func runRelaySpaceMemberStateCommand(
	commsDir string,
	statePath string,
	relaySpaceID string,
	routingMemberID string,
	targetState string,
) (string, error) {
	args := []string{
		"run",
		"./cmd/comms",
		"relay-space-member-state-dev",
		"--state",
		statePath,
		"--relay-space-id",
		relaySpaceID,
		"--routing-member-id",
		routingMemberID,
		"--target-state",
		targetState,
	}

	fmt.Println()
	fmt.Println("COMMS MEMBER STATE: go", strings.Join(args, " "))

	cmd := exec.Command("go", args...)
	cmd.Dir = commsDir
	output, err := cmd.CombinedOutput()
	fmt.Print(string(output))
	if err != nil {
		return string(output), fmt.Errorf(
			"Comms member-state command failed: %w",
			err,
		)
	}

	return string(output), nil
}

func relaySpaceMemberStatePOST(
	url string,
	payload map[string]any,
) (int, string, error) {
	raw, err := json.Marshal(payload)
	if err != nil {
		return 0, "", fmt.Errorf("encode POST payload: %w", err)
	}

	req, err := http.NewRequest(
		http.MethodPost,
		url,
		bytes.NewReader(raw),
	)
	if err != nil {
		return 0, "", fmt.Errorf("create POST request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, "", fmt.Errorf("perform POST request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, "", fmt.Errorf("read POST response: %w", err)
	}

	return resp.StatusCode, string(body), nil
}
