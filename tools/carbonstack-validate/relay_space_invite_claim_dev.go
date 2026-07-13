package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type relaySpaceInviteClaimState struct {
	ServerURL   string `json:"server_url"`
	AccountID   string `json:"account_id"`
	DeviceID    string `json:"device_id"`
	DeviceLabel string `json:"device_label"`
}

func (r *Runner) RelaySpaceInviteClaimDev() error {
	r.PrintHeader("relay-space-invite-claim-dev")

	fmt.Println(
		"status: dev/pre-alpha Relay Space invite-claim lifecycle profile",
	)
	fmt.Println(
		"scope: existing account/device -> full-token claim -> " +
			"idempotent retry -> routing-member persistence",
	)
	fmt.Println(
		"boundary: routing and coordination authority only; not identity " +
			"verification, trust promotion, OpenMLS membership, secure " +
			"enrollment, deployment, or production UX",
	)
	fmt.Println(
		"state boundary: the Comms claim command reads an explicit --state " +
			"path and must not rewrite that local state",
	)

	if err := r.CheckRequiredPaths(); err != nil {
		return err
	}
	if err := r.CheckLiveGitUmbrella(
		"relay-space-invite-claim-dev",
	); err != nil {
		return err
	}

	tempRoot, err := os.MkdirTemp(
		"",
		"carbonstack-relay-space-invite-claim-dev-*",
	)
	if err != nil {
		return fmt.Errorf("create invite-claim temp root: %w", err)
	}
	defer func() {
		if err := os.RemoveAll(tempRoot); err != nil {
			fmt.Printf(
				"WARN: remove invite-claim temp root %s: %v\n",
				tempRoot,
				err,
			)
		}
	}()

	binPath := filepath.Join(tempRoot, "carbonstack-cypher-invite-claim")
	dbPath := filepath.Join(tempRoot, "cypher.db")
	aliceStatePath := filepath.Join(tempRoot, "alice-state.json")
	bobStatePath := filepath.Join(tempRoot, "bob-state.json")

	port, err := reserveLoopbackPort()
	if err != nil {
		return err
	}
	baseURL := "http://127.0.0.1:" + strconv.Itoa(port)

	fmt.Println()
	fmt.Println("== invite-claim generated-state root ==")
	fmt.Println("temp_root:", tempRoot)
	fmt.Println("cypher_db:", dbPath)
	fmt.Println("cypher_url:", baseURL)
	fmt.Println(
		"note: temp state is runner-owned and removed after profile completion",
	)

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
			_ = server.stop("invite-claim")
		}
	}()

	if err := waitForLocalCypherHealth(baseURL + "/v0/health"); err != nil {
		return err
	}

	runID := strconv.FormatInt(time.Now().UnixNano(), 10)
	aliceInvite := "b4c-alice-" + runID
	bobInvite := "b4c-bob-" + runID

	if err := runRelaySpaceInviteClaimSetupCommand(
		r.Comms,
		"init",
		"--state",
		aliceStatePath,
		"--server",
		baseURL,
	); err != nil {
		return err
	}
	if err := runRelaySpaceInviteClaimSetupCommand(
		r.Comms,
		"dev-create-invite",
		"--state",
		aliceStatePath,
		"--server",
		baseURL,
		"--invite",
		aliceInvite,
	); err != nil {
		return err
	}
	if err := runRelaySpaceInviteClaimSetupCommand(
		r.Comms,
		"claim-invite",
		"--state",
		aliceStatePath,
		"--server",
		baseURL,
		"--invite",
		aliceInvite,
		"--name",
		"Alice B4c",
	); err != nil {
		return err
	}
	if err := runRelaySpaceInviteClaimSetupCommand(
		r.Comms,
		"register-device",
		"--state",
		aliceStatePath,
		"--label",
		"alice-b4c-"+runID,
	); err != nil {
		return err
	}

	if err := runRelaySpaceInviteClaimSetupCommand(
		r.Comms,
		"init",
		"--state",
		bobStatePath,
		"--server",
		baseURL,
	); err != nil {
		return err
	}
	if err := runRelaySpaceInviteClaimSetupCommand(
		r.Comms,
		"dev-create-invite",
		"--state",
		bobStatePath,
		"--server",
		baseURL,
		"--invite",
		bobInvite,
	); err != nil {
		return err
	}
	if err := runRelaySpaceInviteClaimSetupCommand(
		r.Comms,
		"claim-invite",
		"--state",
		bobStatePath,
		"--server",
		baseURL,
		"--invite",
		bobInvite,
		"--name",
		"Bob B4c",
	); err != nil {
		return err
	}
	if err := runRelaySpaceInviteClaimSetupCommand(
		r.Comms,
		"register-device",
		"--state",
		bobStatePath,
		"--label",
		"bob-b4c-"+runID,
	); err != nil {
		return err
	}

	aliceState, err := readRelaySpaceInviteClaimState(aliceStatePath)
	if err != nil {
		return err
	}
	bobState, err := readRelaySpaceInviteClaimState(bobStatePath)
	if err != nil {
		return err
	}

	relaySpaceID := "relay-space-b4c-" + runID
	space, err := localCypherPOST(
		baseURL+"/v0/relay-spaces",
		map[string]any{
			"relay_space_id":        relaySpaceID,
			"display_label":         "B4c invite claim space",
			"created_by_account_id": aliceState.AccountID,
			"created_by_device_id":  aliceState.DeviceID,
		},
		201,
	)
	if err != nil {
		return err
	}
	if stringField(space, "relay_space_id") != relaySpaceID {
		return fmt.Errorf("created Relay Space ID mismatch")
	}
	fmt.Println("PASS: created Relay Space")

	creator, err := localCypherPOST(
		baseURL+"/v0/relay-spaces/"+relaySpaceID+"/members",
		map[string]any{
			"routing_member_id": "creator-b4c-" + runID,
			"account_id":        aliceState.AccountID,
			"device_id":         aliceState.DeviceID,
			"display_label":     "Alice B4c creator",
		},
		201,
	)
	if err != nil {
		return err
	}
	creatorMemberID := stringField(creator, "routing_member_id")
	if creatorMemberID == "" {
		return fmt.Errorf("creator routing_member_id missing")
	}
	fmt.Println("PASS: explicitly registered creator routing member")

	fullInviteToken := "full-relay-space-token-b4c-" + runID
	inviteID := "relay-invite-b4c-" + runID
	_, err = localCypherPOST(
		baseURL+"/v0/relay-spaces/"+relaySpaceID+"/invites",
		map[string]any{
			"relay_space_invite_id": inviteID,
			"invite_token":          fullInviteToken,
			"display_code":          "B4C-" + runID[len(runID)-6:],
			"word_code":             "b4c-full-token-only",
			"created_by_member_id":  creatorMemberID,
			"max_claims":            1,
			"note":                  "B4c Comms operator claim proof",
		},
		201,
	)
	if err != nil {
		return err
	}
	fmt.Println("PASS: created one-use Relay Space invite")

	stateBefore, err := os.ReadFile(bobStatePath)
	if err != nil {
		return fmt.Errorf("read Bob state before claim: %w", err)
	}

	createdOutput, err := runRelaySpaceInviteClaimCommand(
		r.Comms,
		"relay-space-invite-claim-dev",
		"--state",
		bobStatePath,
		"--invite-token",
		fullInviteToken,
		"--display-label",
		"Bob B4c routing member",
	)
	if err != nil {
		return err
	}

	for _, expected := range []string{
		"claim_classification: created",
		"idempotent: false",
		"claim_consumed: true",
		"relay_space_id: " + relaySpaceID,
		"account_id: " + bobState.AccountID,
		"device_id: " + bobState.DeviceID,
		"invite_state: claimed",
		"invite_claim_count: 1",
		"local_state_mutated: false",
		"routing and coordination authority only",
	} {
		if !strings.Contains(createdOutput, expected) {
			return fmt.Errorf(
				"created claim output missing %q:\n%s",
				expected,
				createdOutput,
			)
		}
	}
	fmt.Println("PASS: first Comms claim classified created")

	stateAfterCreated, err := os.ReadFile(bobStatePath)
	if err != nil {
		return fmt.Errorf("read Bob state after created claim: %w", err)
	}
	if !bytes.Equal(stateBefore, stateAfterCreated) {
		return fmt.Errorf("created claim rewrote Bob local state")
	}
	fmt.Println("PASS: first claim did not mutate Bob local state")

	retryOutput, err := runRelaySpaceInviteClaimCommand(
		r.Comms,
		"relay-space-invite-claim-dev",
		"--state",
		bobStatePath,
		"--invite-token",
		fullInviteToken,
	)
	if err != nil {
		return err
	}

	for _, expected := range []string{
		"claim_classification: already_active",
		"idempotent: true",
		"claim_consumed: false",
		"relay_space_id: " + relaySpaceID,
		"account_id: " + bobState.AccountID,
		"device_id: " + bobState.DeviceID,
		"invite_state: claimed",
		"invite_claim_count: 1",
		"local_state_mutated: false",
	} {
		if !strings.Contains(retryOutput, expected) {
			return fmt.Errorf(
				"idempotent retry output missing %q:\n%s",
				expected,
				retryOutput,
			)
		}
	}
	fmt.Println("PASS: retry classified already_active without claim use")

	stateAfterRetry, err := os.ReadFile(bobStatePath)
	if err != nil {
		return fmt.Errorf("read Bob state after retry: %w", err)
	}
	if !bytes.Equal(stateBefore, stateAfterRetry) {
		return fmt.Errorf("idempotent retry rewrote Bob local state")
	}
	fmt.Println("PASS: retry did not mutate Bob local state")

	inviteClaimCount, err := relaySpaceInviteClaimSQLiteScalar(
		dbPath,
		"SELECT claim_count FROM relay_space_invites "+
			"WHERE relay_space_invite_id = "+
			relaySpaceInviteClaimSQLQuote(inviteID)+";",
	)
	if err != nil {
		return err
	}
	if inviteClaimCount != "1" {
		return fmt.Errorf(
			"invite claim_count = %q, want 1",
			inviteClaimCount,
		)
	}

	inviteState, err := relaySpaceInviteClaimSQLiteScalar(
		dbPath,
		"SELECT state FROM relay_space_invites "+
			"WHERE relay_space_invite_id = "+
			relaySpaceInviteClaimSQLQuote(inviteID)+";",
	)
	if err != nil {
		return err
	}
	if inviteState != "claimed" {
		return fmt.Errorf("invite state = %q, want claimed", inviteState)
	}

	memberCount, err := relaySpaceInviteClaimSQLiteScalar(
		dbPath,
		"SELECT COUNT(*) FROM relay_space_members "+
			"WHERE relay_space_id = "+
			relaySpaceInviteClaimSQLQuote(relaySpaceID)+";",
	)
	if err != nil {
		return err
	}
	if memberCount != "2" {
		return fmt.Errorf(
			"Relay Space member count = %q, want 2",
			memberCount,
		)
	}
	fmt.Println(
		"PASS: DB has one consumed claim, claimed invite, " +
			"and creator plus claimant",
	)

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
	fmt.Println("PASS: claim path did not create trust/candidate state")

	if err := server.stop("invite-claim"); err != nil {
		return err
	}
	serverRunning = false

	fmt.Println()
	fmt.Println("relay-space-invite-claim-dev profile result:")
	fmt.Println("  first_claim: created")
	fmt.Println("  first_claim_consumed: true")
	fmt.Println("  retry_classification: already_active")
	fmt.Println("  retry_claim_consumed: false")
	fmt.Println("  local_state_mutated: false")
	fmt.Println("  trust_or_candidate_state_mutated: false")
	fmt.Println("  invite_claim_count: 1")
	fmt.Println("  invite_state: claimed")
	fmt.Println("  routing_member_count: 2")
	fmt.Println(
		"  boundary: routing coordination only; no identity, trust, " +
			"OpenMLS membership, deployment, or production claim",
	)

	return nil
}

func readRelaySpaceInviteClaimState(
	path string,
) (relaySpaceInviteClaimState, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return relaySpaceInviteClaimState{}, fmt.Errorf(
			"read state %s: %w",
			path,
			err,
		)
	}

	var state relaySpaceInviteClaimState
	if err := json.Unmarshal(raw, &state); err != nil {
		return relaySpaceInviteClaimState{}, fmt.Errorf(
			"decode state %s: %w",
			path,
			err,
		)
	}

	if state.ServerURL == "" ||
		state.AccountID == "" ||
		state.DeviceID == "" {
		return relaySpaceInviteClaimState{}, fmt.Errorf(
			"state %s lacks server/account/device context",
			path,
		)
	}

	return state, nil
}

func runRelaySpaceInviteClaimSetupCommand(
	commsDir string,
	args ...string,
) error {
	fmt.Println()
	fmt.Println(
		"COMMS SETUP:",
		"go run ./cmd/comms "+strings.Join(args, " "),
	)

	cmd := exec.Command("go", append([]string{"run", "./cmd/comms"}, args...)...)
	cmd.Dir = commsDir
	output, err := cmd.CombinedOutput()
	if len(output) > 0 {
		fmt.Print(string(output))
	}
	if err != nil {
		return fmt.Errorf(
			"Comms setup command failed: %w",
			err,
		)
	}

	return nil
}

func runRelaySpaceInviteClaimCommand(
	commsDir string,
	args ...string,
) (string, error) {
	fmt.Println()
	fmt.Println(
		"COMMS CLAIM:",
		"go run ./cmd/comms "+strings.Join(args, " "),
	)

	cmd := exec.Command("go", append([]string{"run", "./cmd/comms"}, args...)...)
	cmd.Dir = commsDir
	output, err := cmd.CombinedOutput()
	fmt.Print(string(output))
	if err != nil {
		return string(output), fmt.Errorf(
			"Comms invite-claim command failed: %w",
			err,
		)
	}

	return string(output), nil
}

func relaySpaceInviteClaimSQLiteScalar(
	dbPath string,
	query string,
) (string, error) {
	cmd := exec.Command("sqlite3", dbPath, query)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf(
			"sqlite scalar query failed: %w: %s",
			err,
			string(output),
		)
	}

	return strings.TrimSpace(string(output)), nil
}

func relaySpaceInviteClaimSQLQuote(value string) string {
	return "'" + strings.ReplaceAll(value, "'", "''") + "'"
}
