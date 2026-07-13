package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func (r *Runner) RelaySpaceMemberRestartDev() error {
	r.PrintHeader("relay-space-member-restart-dev")

	fmt.Println(
		"status: dev/pre-alpha Relay Space member inspection and restart profile",
	)
	fmt.Println(
		"scope: persist disabled and left routing-member states across " +
			"Cypher restarts, inspect them through the member-list route, " +
			"and re-prove routing and rejoin semantics",
	)
	fmt.Println(
		"boundary: restart persistence and inspection proof only; not " +
			"authenticated administration, production authorization, " +
			"identity verification, trust promotion, OpenMLS membership, " +
			"member deletion, backup/restore, rollback safety, or rejoin",
	)

	if err := r.CheckRequiredPaths(); err != nil {
		return err
	}
	if err := r.CheckLiveGitUmbrella(
		"relay-space-member-restart-dev",
	); err != nil {
		return err
	}

	tempRoot, err := os.MkdirTemp(
		"",
		"carbonstack-relay-space-member-restart-dev-*",
	)
	if err != nil {
		return fmt.Errorf("create member-restart temp root: %w", err)
	}
	defer func() {
		if err := os.RemoveAll(tempRoot); err != nil {
			fmt.Printf(
				"WARN: remove member-restart temp root %s: %v\n",
				tempRoot,
				err,
			)
		}
	}()

	binPath := filepath.Join(tempRoot, "carbonstack-cypher-member-restart")
	dbPath := filepath.Join(tempRoot, "cypher.db")
	aliceStatePath := filepath.Join(tempRoot, "alice-state.json")
	bobStatePath := filepath.Join(tempRoot, "bob-state.json")

	port, err := reserveLoopbackPort()
	if err != nil {
		return err
	}
	baseURL := "http://127.0.0.1:" + strconv.Itoa(port)

	fmt.Println()
	fmt.Println("== member-restart generated-state root ==")
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

	start := func() (*localCypherServer, error) {
		server, err := startLocalCypherServer(binPath, r.Cypher, env)
		if err != nil {
			return nil, err
		}
		if err := waitForLocalCypherHealth(
			baseURL + "/v0/health",
		); err != nil {
			_ = server.stop("member-restart-start-failure")
			return nil, err
		}
		return server, nil
	}

	server, err := start()
	if err != nil {
		return err
	}
	serverRunning := true
	defer func() {
		if serverRunning {
			_ = server.stop("member-restart")
		}
	}()

	runID := strconv.FormatInt(time.Now().UnixNano(), 10)
	aliceInvite := "b4e-alice-" + runID
	bobInvite := "b4e-bob-" + runID

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
			"Alice B4e",
		},
		{
			"register-device",
			"--state",
			aliceStatePath,
			"--label",
			"alice-b4e-" + runID,
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
			"Bob B4e",
		},
		{
			"register-device",
			"--state",
			bobStatePath,
			"--label",
			"bob-b4e-" + runID,
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

	aliceBefore, err := os.ReadFile(aliceStatePath)
	if err != nil {
		return fmt.Errorf("read Alice state before: %w", err)
	}
	bobBefore, err := os.ReadFile(bobStatePath)
	if err != nil {
		return fmt.Errorf("read Bob state before: %w", err)
	}

	relaySpaceID := "relay-space-b4e-" + runID
	aliceMemberID := "alice-member-b4e-" + runID
	bobMemberID := "bob-member-b4e-" + runID

	if _, err := localCypherPOST(
		baseURL+"/v0/relay-spaces",
		map[string]any{
			"relay_space_id":        relaySpaceID,
			"display_label":         "B4e restart inspection space",
			"created_by_account_id": aliceState.AccountID,
			"created_by_device_id":  aliceState.DeviceID,
		},
		http.StatusCreated,
	); err != nil {
		return err
	}

	for _, member := range []struct {
		id    string
		state relaySpaceInviteClaimState
		label string
	}{
		{
			id:    aliceMemberID,
			state: aliceState,
			label: "Alice B4e operator",
		},
		{
			id:    bobMemberID,
			state: bobState,
			label: "Bob B4e member",
		},
	} {
		if _, err := localCypherPOST(
			baseURL+"/v0/relay-spaces/"+relaySpaceID+"/members",
			map[string]any{
				"routing_member_id": member.id,
				"account_id":        member.state.AccountID,
				"device_id":         member.state.DeviceID,
				"display_label":     member.label,
			},
			http.StatusCreated,
		); err != nil {
			return err
		}
	}
	fmt.Println("PASS: created Relay Space and two routing members")

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
		"current_state: disabled",
		"routing_member_id: " + bobMemberID,
		"device_id: " + bobState.DeviceID,
	} {
		if !strings.Contains(disableOutput, expected) {
			return fmt.Errorf(
				"disable output missing %q:\n%s",
				expected,
				disableOutput,
			)
		}
	}

	if err := requireRelaySpaceRoutingResult(
		baseURL,
		relaySpaceID,
		aliceState.DeviceID,
		bobState.DeviceID,
		http.StatusForbidden,
		"recipient_not_relay_member",
	); err != nil {
		return err
	}
	fmt.Println("PASS: disabled routing refused before restart")

	if err := server.stop("member-restart-disabled"); err != nil {
		return err
	}
	serverRunning = false

	server, err = start()
	if err != nil {
		return err
	}
	serverRunning = true

	disabledMember, err := inspectRelaySpaceMemberAfterRestart(
		baseURL,
		relaySpaceID,
		bobMemberID,
	)
	if err != nil {
		return err
	}
	if disabledMember.State != "disabled" {
		return fmt.Errorf(
			"disabled member state after restart = %q",
			disabledMember.State,
		)
	}
	if disabledMember.DisabledAt == "" {
		return fmt.Errorf(
			"disabled member lost disabled_at after restart",
		)
	}
	fmt.Println(
		"PASS: member inspection preserved disabled state and disabled_at",
	)

	if err := requireRelaySpaceRoutingResult(
		baseURL,
		relaySpaceID,
		aliceState.DeviceID,
		bobState.DeviceID,
		http.StatusForbidden,
		"recipient_not_relay_member",
	); err != nil {
		return err
	}
	fmt.Println("PASS: disabled routing refused after restart")

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
	if !strings.Contains(reactivateOutput, "current_state: active") ||
		!strings.Contains(reactivateOutput, "\ndisabled_at: \n") {
		return fmt.Errorf(
			"reactivation output did not prove active with empty disabled_at:\n%s",
			reactivateOutput,
		)
	}

	if err := requireRelaySpaceRoutingResult(
		baseURL,
		relaySpaceID,
		aliceState.DeviceID,
		bobState.DeviceID,
		http.StatusCreated,
		"",
	); err != nil {
		return err
	}
	fmt.Println("PASS: reactivation restored routing authority")

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
	if !strings.Contains(leaveOutput, "current_state: left") {
		return fmt.Errorf("leave output unexpected:\n%s", leaveOutput)
	}

	if err := server.stop("member-restart-left"); err != nil {
		return err
	}
	serverRunning = false

	server, err = start()
	if err != nil {
		return err
	}
	serverRunning = true

	leftMember, err := inspectRelaySpaceMemberAfterRestart(
		baseURL,
		relaySpaceID,
		bobMemberID,
	)
	if err != nil {
		return err
	}
	if leftMember.State != "left" {
		return fmt.Errorf(
			"left member state after restart = %q",
			leftMember.State,
		)
	}
	if leftMember.DisabledAt != "" {
		return fmt.Errorf(
			"left member disabled_at after restart = %q",
			leftMember.DisabledAt,
		)
	}
	fmt.Println(
		"PASS: member inspection preserved left state without disabled_at",
	)

	if err := requireRelaySpaceRoutingResult(
		baseURL,
		relaySpaceID,
		aliceState.DeviceID,
		bobState.DeviceID,
		http.StatusForbidden,
		"recipient_not_relay_member",
	); err != nil {
		return err
	}
	fmt.Println("PASS: left routing remained refused after restart")

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
			"left-to-active refusal lacked explicit classification:\n%s",
			rejoinOutput,
		)
	}
	fmt.Println(
		"PASS: left-to-active remained explicit rejoin refusal after restart",
	)

	aliceAfter, err := os.ReadFile(aliceStatePath)
	if err != nil {
		return fmt.Errorf("read Alice state after: %w", err)
	}
	bobAfter, err := os.ReadFile(bobStatePath)
	if err != nil {
		return fmt.Errorf("read Bob state after: %w", err)
	}
	if !bytes.Equal(aliceBefore, aliceAfter) {
		return fmt.Errorf("Alice state changed during restart profile")
	}
	if !bytes.Equal(bobBefore, bobAfter) {
		return fmt.Errorf("Bob state changed during restart profile")
	}
	fmt.Println("PASS: local state files remained byte-identical")

	finalState, err := relaySpaceInviteClaimSQLiteScalar(
		dbPath,
		"SELECT state FROM relay_space_members WHERE routing_member_id = "+
			relaySpaceInviteClaimSQLQuote(bobMemberID)+";",
	)
	if err != nil {
		return err
	}
	if finalState != "left" {
		return fmt.Errorf("final Bob state = %q, want left", finalState)
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
	fmt.Println("PASS: final database state and envelope count are correct")

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
	}
	fmt.Println("PASS: no trust or identity-candidate state was created")

	if err := server.stop("member-restart-complete"); err != nil {
		return err
	}
	serverRunning = false

	fmt.Println()
	fmt.Println("relay-space-member-restart-dev profile result:")
	fmt.Println("  disabled_state_survived_restart: true")
	fmt.Println("  disabled_at_survived_restart: true")
	fmt.Println("  disabled_routing_refused_after_restart: true")
	fmt.Println("  reactivation_restored_routing: true")
	fmt.Println("  left_state_survived_restart: true")
	fmt.Println("  left_disabled_at_absent_after_restart: true")
	fmt.Println("  left_routing_refused_after_restart: true")
	fmt.Println("  left_to_active_rejoin_required: true")
	fmt.Println("  member_list_inspection_proven: true")
	fmt.Println("  local_state_files_mutated: false")
	fmt.Println("  trust_or_candidate_state_mutated: false")
	fmt.Println("  successful_envelopes: 1")
	fmt.Println(
		"  boundary: dev restart/inspection proof only; no authenticated " +
			"administration, production authorization, backup/restore, " +
			"rollback safety, identity/trust promotion, OpenMLS membership, " +
			"member deletion, or rejoin claim",
	)

	return nil
}

type relaySpaceMemberRestartInspection struct {
	RoutingMemberID string
	RelaySpaceID    string
	AccountID       string
	DeviceID        string
	State           string
	DisabledAt      string
}

func inspectRelaySpaceMemberAfterRestart(
	baseURL string,
	relaySpaceID string,
	routingMemberID string,
) (relaySpaceMemberRestartInspection, error) {
	resp, err := http.Get(
		baseURL + "/v0/relay-spaces/" + relaySpaceID + "/members",
	)
	if err != nil {
		return relaySpaceMemberRestartInspection{}, fmt.Errorf(
			"inspect Relay Space members: %w",
			err,
		)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return relaySpaceMemberRestartInspection{}, fmt.Errorf(
			"read Relay Space member inspection: %w",
			err,
		)
	}
	if resp.StatusCode != http.StatusOK {
		return relaySpaceMemberRestartInspection{}, fmt.Errorf(
			"member inspection status %d: %s",
			resp.StatusCode,
			string(body),
		)
	}

	var decoded any
	if err := json.Unmarshal(body, &decoded); err != nil {
		return relaySpaceMemberRestartInspection{}, fmt.Errorf(
			"decode member inspection: %w",
			err,
		)
	}

	member := findRelaySpaceMemberInspection(decoded, routingMemberID)
	if member == nil {
		return relaySpaceMemberRestartInspection{}, fmt.Errorf(
			"routing member %s absent from inspection response",
			routingMemberID,
		)
	}

	return relaySpaceMemberRestartInspection{
		RoutingMemberID: stringMapValue(member, "routing_member_id"),
		RelaySpaceID:    stringMapValue(member, "relay_space_id"),
		AccountID:       stringMapValue(member, "account_id"),
		DeviceID:        stringMapValue(member, "device_id"),
		State:           stringMapValue(member, "state"),
		DisabledAt:      stringMapValue(member, "disabled_at"),
	}, nil
}

func findRelaySpaceMemberInspection(
	value any,
	routingMemberID string,
) map[string]any {
	switch typed := value.(type) {
	case map[string]any:
		if stringMapValue(typed, "routing_member_id") == routingMemberID {
			return typed
		}
		for _, nested := range typed {
			if found := findRelaySpaceMemberInspection(
				nested,
				routingMemberID,
			); found != nil {
				return found
			}
		}
	case []any:
		for _, nested := range typed {
			if found := findRelaySpaceMemberInspection(
				nested,
				routingMemberID,
			); found != nil {
				return found
			}
		}
	}
	return nil
}

func stringMapValue(value map[string]any, key string) string {
	raw, ok := value[key]
	if !ok || raw == nil {
		return ""
	}
	text, _ := raw.(string)
	return text
}

func requireRelaySpaceRoutingResult(
	baseURL string,
	relaySpaceID string,
	senderDeviceID string,
	recipientDeviceID string,
	expectedStatus int,
	expectedMarker string,
) error {
	status, body, err := relaySpaceMemberStatePOST(
		baseURL+"/v0/relay-spaces/"+relaySpaceID+"/envelopes",
		map[string]any{
			"sender_device_id":    senderDeviceID,
			"recipient_device_id": recipientDeviceID,
			"content_type":        "carbonstack.message.text.stub.v0",
			"protocol_version":    "stub-v0",
			"ciphertext_b64":      "AQID",
		},
	)
	if err != nil {
		return err
	}
	if status != expectedStatus {
		return fmt.Errorf(
			"routing status = %d, want %d: %s",
			status,
			expectedStatus,
			body,
		)
	}
	if expectedMarker != "" && !strings.Contains(body, expectedMarker) {
		return fmt.Errorf(
			"routing response lacks %q: %s",
			expectedMarker,
			body,
		)
	}
	return nil
}
