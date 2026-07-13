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

func (r *Runner) RelaySpaceDeliveryAuthorityDev() error {
	r.PrintHeader("relay-space-delivery-authority-dev")

	fmt.Println(
		"status: dev/pre-alpha scoped delivery-authority profile",
	)
	fmt.Println(
		"scope: prove queued Relay Space envelopes remain persisted but " +
			"scoped inbox and scoped ACK require current active recipient " +
			"membership across disable, restart, reactivation, and leave",
	)
	fmt.Println(
		"boundary: routing/delivery authority only; not identity, trust, " +
			"OpenMLS membership, KeyPackage consumption, Welcome lifecycle, " +
			"authenticated administration, deletion, or production safety",
	)

	if err := r.CheckRequiredPaths(); err != nil {
		return err
	}
	if err := r.CheckLiveGitUmbrella(
		"relay-space-delivery-authority-dev",
	); err != nil {
		return err
	}

	tempRoot, err := os.MkdirTemp(
		"",
		"carbonstack-relay-space-delivery-authority-dev-*",
	)
	if err != nil {
		return fmt.Errorf("create delivery-authority temp root: %w", err)
	}
	defer func() {
		if err := os.RemoveAll(tempRoot); err != nil {
			fmt.Printf(
				"WARN: remove delivery-authority temp root %s: %v\n",
				tempRoot,
				err,
			)
		}
	}()

	binPath := filepath.Join(
		tempRoot,
		"carbonstack-cypher-delivery-authority",
	)
	dbPath := filepath.Join(tempRoot, "cypher.db")
	aliceStatePath := filepath.Join(tempRoot, "alice-state.json")
	bobStatePath := filepath.Join(tempRoot, "bob-state.json")

	port, err := reserveLoopbackPort()
	if err != nil {
		return err
	}
	baseURL := "http://127.0.0.1:" + strconv.Itoa(port)

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
			_ = server.stop("delivery-authority-start-failure")
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
			_ = server.stop("delivery-authority")
		}
	}()

	runID := strconv.FormatInt(time.Now().UnixNano(), 10)
	aliceInvite := "delivery-authority-alice-" + runID
	bobInvite := "delivery-authority-bob-" + runID

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
			"Alice delivery authority",
		},
		{
			"register-device",
			"--state",
			aliceStatePath,
			"--label",
			"alice-delivery-authority-" + runID,
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
			"Bob delivery authority",
		},
		{
			"register-device",
			"--state",
			bobStatePath,
			"--label",
			"bob-delivery-authority-" + runID,
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

	relaySpaceID := "relay-space-delivery-authority-" + runID
	aliceMemberID := "alice-delivery-authority-member-" + runID
	bobMemberID := "bob-delivery-authority-member-" + runID

	if _, err := localCypherPOST(
		baseURL+"/v0/relay-spaces",
		map[string]any{
			"relay_space_id":        relaySpaceID,
			"display_label":         "delivery authority",
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
			label: "Alice delivery authority",
		},
		{
			id:    bobMemberID,
			state: bobState,
			label: "Bob delivery authority",
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

	firstEnvelopeID, err := submitDeliveryAuthorityEnvelope(
		baseURL,
		relaySpaceID,
		aliceState.DeviceID,
		bobState.DeviceID,
		"queued before disable",
	)
	if err != nil {
		return err
	}

	if err := requireDeliveryAuthorityInbox(
		baseURL,
		relaySpaceID,
		bobState.DeviceID,
		http.StatusOK,
		firstEnvelopeID,
		"",
	); err != nil {
		return err
	}

	if _, err := runRelaySpaceMemberStateCommand(
		r.Comms,
		aliceStatePath,
		relaySpaceID,
		bobMemberID,
		"disabled",
	); err != nil {
		return err
	}

	if err := requireDeliveryAuthorityRefusal(
		baseURL,
		relaySpaceID,
		bobState.DeviceID,
		firstEnvelopeID,
	); err != nil {
		return err
	}
	if err := requireDeliveryAuthorityDBState(
		dbPath,
		firstEnvelopeID,
		"queued",
	); err != nil {
		return err
	}
	fmt.Println("PASS: disabled recipient cannot fetch or ack queued envelope")

	if err := server.stop("delivery-authority-disabled"); err != nil {
		return err
	}
	serverRunning = false

	server, err = start()
	if err != nil {
		return err
	}
	serverRunning = true

	if err := requireDeliveryAuthorityRefusal(
		baseURL,
		relaySpaceID,
		bobState.DeviceID,
		firstEnvelopeID,
	); err != nil {
		return err
	}
	if err := requireDeliveryAuthorityDBState(
		dbPath,
		firstEnvelopeID,
		"queued",
	); err != nil {
		return err
	}
	fmt.Println(
		"PASS: disabled refusal and queued state survive Cypher restart",
	)

	if _, err := runRelaySpaceMemberStateCommand(
		r.Comms,
		aliceStatePath,
		relaySpaceID,
		bobMemberID,
		"active",
	); err != nil {
		return err
	}

	if err := requireDeliveryAuthorityInbox(
		baseURL,
		relaySpaceID,
		bobState.DeviceID,
		http.StatusOK,
		firstEnvelopeID,
		"",
	); err != nil {
		return err
	}
	if err := requireDeliveryAuthorityAck(
		baseURL,
		relaySpaceID,
		firstEnvelopeID,
		bobState.DeviceID,
		http.StatusOK,
		"",
	); err != nil {
		return err
	}
	if err := requireDeliveryAuthorityDBState(
		dbPath,
		firstEnvelopeID,
		"acknowledged",
	); err != nil {
		return err
	}
	fmt.Println("PASS: reactivation restores fetch and ack authority")

	secondEnvelopeID, err := submitDeliveryAuthorityEnvelope(
		baseURL,
		relaySpaceID,
		aliceState.DeviceID,
		bobState.DeviceID,
		"queued before leave",
	)
	if err != nil {
		return err
	}

	if _, err := runRelaySpaceMemberStateCommand(
		r.Comms,
		aliceStatePath,
		relaySpaceID,
		bobMemberID,
		"left",
	); err != nil {
		return err
	}

	if err := requireDeliveryAuthorityRefusal(
		baseURL,
		relaySpaceID,
		bobState.DeviceID,
		secondEnvelopeID,
	); err != nil {
		return err
	}
	if err := requireDeliveryAuthorityDBState(
		dbPath,
		secondEnvelopeID,
		"queued",
	); err != nil {
		return err
	}
	fmt.Println("PASS: left recipient cannot fetch or ack queued envelope")

	if err := server.stop("delivery-authority-left"); err != nil {
		return err
	}
	serverRunning = false

	server, err = start()
	if err != nil {
		return err
	}
	serverRunning = true

	if err := requireDeliveryAuthorityRefusal(
		baseURL,
		relaySpaceID,
		bobState.DeviceID,
		secondEnvelopeID,
	); err != nil {
		return err
	}
	if err := requireDeliveryAuthorityDBState(
		dbPath,
		secondEnvelopeID,
		"queued",
	); err != nil {
		return err
	}
	fmt.Println("PASS: left refusal and queued state survive Cypher restart")

	aliceAfter, err := os.ReadFile(aliceStatePath)
	if err != nil {
		return fmt.Errorf("read Alice state after: %w", err)
	}
	bobAfter, err := os.ReadFile(bobStatePath)
	if err != nil {
		return fmt.Errorf("read Bob state after: %w", err)
	}
	if !bytes.Equal(aliceBefore, aliceAfter) {
		return fmt.Errorf("Alice state changed during delivery-authority profile")
	}
	if !bytes.Equal(bobBefore, bobAfter) {
		return fmt.Errorf("Bob state changed during delivery-authority profile")
	}

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

	if err := server.stop("delivery-authority-complete"); err != nil {
		return err
	}
	serverRunning = false

	fmt.Println()
	fmt.Println("relay-space-delivery-authority-dev profile result:")
	fmt.Println("  active_membership_required_for_scoped_inbox: true")
	fmt.Println("  active_membership_required_for_scoped_ack: true")
	fmt.Println("  disabled_envelope_remains_queued: true")
	fmt.Println("  disabled_refusal_survives_restart: true")
	fmt.Println("  reactivation_restores_fetch_and_ack: true")
	fmt.Println("  left_envelope_remains_queued: true")
	fmt.Println("  left_refusal_survives_restart: true")
	fmt.Println("  local_state_files_mutated: false")
	fmt.Println("  trust_or_candidate_state_mutated: false")
	fmt.Println(
		"  boundary: routing/delivery authority only; not KeyPackage " +
			"consumption, Welcome lifecycle, identity, trust, MLS membership, " +
			"authenticated administration, or production safety",
	)

	return nil
}

func submitDeliveryAuthorityEnvelope(
	baseURL string,
	relaySpaceID string,
	senderDeviceID string,
	recipientDeviceID string,
	payload string,
) (string, error) {
	body, err := localCypherPOST(
		baseURL+"/v0/relay-spaces/"+relaySpaceID+"/envelopes",
		map[string]any{
			"sender_device_id":    senderDeviceID,
			"recipient_device_id": recipientDeviceID,
			"content_type":        "carbonstack.message.text.stub.v0",
			"protocol_version":    "stub-v0",
			"ciphertext_b64":      "AQID",
			"client_created_at":   payload,
		},
		http.StatusCreated,
	)
	if err != nil {
		return "", err
	}

	envelopeID := strings.TrimSpace(stringField(body, "envelope_id"))
	if envelopeID == "" {
		return "", fmt.Errorf(
			"delivery-authority submit returned no envelope_id",
		)
	}
	return envelopeID, nil
}

func requireDeliveryAuthorityRefusal(
	baseURL string,
	relaySpaceID string,
	deviceID string,
	envelopeID string,
) error {
	if err := requireDeliveryAuthorityInbox(
		baseURL,
		relaySpaceID,
		deviceID,
		http.StatusForbidden,
		"",
		"recipient_not_relay_member",
	); err != nil {
		return err
	}
	return requireDeliveryAuthorityAck(
		baseURL,
		relaySpaceID,
		envelopeID,
		deviceID,
		http.StatusForbidden,
		"recipient_not_relay_member",
	)
}

func requireDeliveryAuthorityInbox(
	baseURL string,
	relaySpaceID string,
	deviceID string,
	expectedStatus int,
	expectedEnvelopeID string,
	expectedMarker string,
) error {
	resp, err := http.Get(
		baseURL +
			"/v0/relay-spaces/" +
			relaySpaceID +
			"/devices/" +
			deviceID +
			"/envelopes",
	)
	if err != nil {
		return fmt.Errorf("delivery-authority inbox: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read delivery-authority inbox: %w", err)
	}
	if resp.StatusCode != expectedStatus {
		return fmt.Errorf(
			"delivery-authority inbox status = %d, want %d: %s",
			resp.StatusCode,
			expectedStatus,
			string(body),
		)
	}
	if expectedMarker != "" &&
		!strings.Contains(string(body), expectedMarker) {
		return fmt.Errorf(
			"delivery-authority inbox lacks %q: %s",
			expectedMarker,
			string(body),
		)
	}
	if expectedEnvelopeID != "" {
		var response struct {
			Envelopes []struct {
				EnvelopeID string `json:"envelope_id"`
			} `json:"envelopes"`
		}
		if err := json.Unmarshal(body, &response); err != nil {
			return fmt.Errorf(
				"decode delivery-authority inbox: %w",
				err,
			)
		}
		if len(response.Envelopes) != 1 ||
			response.Envelopes[0].EnvelopeID != expectedEnvelopeID {
			return fmt.Errorf(
				"delivery-authority inbox = %+v, want one envelope %s",
				response.Envelopes,
				expectedEnvelopeID,
			)
		}
	}
	return nil
}

func requireDeliveryAuthorityAck(
	baseURL string,
	relaySpaceID string,
	envelopeID string,
	deviceID string,
	expectedStatus int,
	expectedMarker string,
) error {
	status, body, err := relaySpaceMemberStatePOST(
		baseURL+
			"/v0/relay-spaces/"+
			relaySpaceID+
			"/envelopes/"+
			envelopeID+
			"/ack",
		map[string]any{
			"recipient_device_id": deviceID,
		},
	)
	if err != nil {
		return err
	}
	if status != expectedStatus {
		return fmt.Errorf(
			"delivery-authority ack status = %d, want %d: %s",
			status,
			expectedStatus,
			body,
		)
	}
	if expectedMarker != "" && !strings.Contains(body, expectedMarker) {
		return fmt.Errorf(
			"delivery-authority ack lacks %q: %s",
			expectedMarker,
			body,
		)
	}
	return nil
}

func requireDeliveryAuthorityDBState(
	dbPath string,
	envelopeID string,
	expectedState string,
) error {
	state, err := relaySpaceInviteClaimSQLiteScalar(
		dbPath,
		"SELECT delivery_state FROM envelopes WHERE envelope_id = "+
			relaySpaceInviteClaimSQLQuote(envelopeID)+";",
	)
	if err != nil {
		return err
	}
	if state != expectedState {
		return fmt.Errorf(
			"envelope %s delivery_state = %q, want %q",
			envelopeID,
			state,
			expectedState,
		)
	}
	return nil
}
