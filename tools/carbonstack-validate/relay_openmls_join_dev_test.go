package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestFormatRelayOpenMLSCompactSummary(t *testing.T) {
	summary := formatRelayOpenMLSCompactSummary([]relayOpenMLSJoinSubrunResult{
		{
			Name:                    "no-ack",
			AckAfter:                false,
			RelaySpaceID:            "relay-noack",
			KeyPackageDeliveryState: "queued",
			WelcomeDeliveryState:    "queued",
			EnvelopeCount:           2,
			AckCount:                0,
			WelcomeAckRows:          0,
			TrustCandidateCheck:     "checked absent before and after subrun",
			TempDirLifecycle:        "runner-owned temp root removed after profile completion",
		},
		{
			Name:                    "ack-after-join",
			AckAfter:                true,
			RelaySpaceID:            "relay-ack",
			KeyPackageDeliveryState: "queued",
			WelcomeDeliveryState:    "acknowledged",
			EnvelopeCount:           2,
			AckCount:                1,
			WelcomeAckRows:          1,
			TrustCandidateCheck:     "checked absent before and after subrun",
			TempDirLifecycle:        "runner-owned temp root removed after profile completion",
		},
	})

	required := []string{
		"== compact relay-openmls-join-dev evidence summary ==",
		"profile: relay-openmls-join-dev",
		"status: PASS",
		"scope: positive-path local/dev Relay Space OpenMLS join validation",
		"subruns: 2",
		"subrun: no-ack",
		"  ack_mode: no-ack",
		"  relay_space_id: relay-noack",
		"  envelopes: 2",
		"  envelope_acks: 0",
		"  keypackage_delivery_state: queued",
		"  welcome_delivery_state: queued",
		"  welcome_ack_rows: 0",
		"subrun: ack-after-join",
		"  ack_mode: ACK_AFTER_JOIN",
		"  relay_space_id: relay-ack",
		"  envelope_acks: 1",
		"  welcome_delivery_state: acknowledged",
		"  welcome_ack_rows: 1",
		"trust_candidate_check: checked absent before and after subrun",
		"temp_dir_lifecycle: runner-owned temp root removed after profile completion",
		"nonclaims: not local-backbone; not production secure messaging; not identity verification; not hostile-server safety; not metadata privacy; not audit/certification",
	}

	for _, needle := range required {
		if !strings.Contains(summary, needle) {
			t.Fatalf("compact summary missing %q\nsummary:\n%s", needle, summary)
		}
	}
}

func TestCollectRelayOpenMLSSubrunResult(t *testing.T) {
	dbPath := filepath.Join(t.TempDir(), "cypher.db")
	createRelayOpenMLSTestDB(t, dbPath)

	sub := &relayOpenMLSJoinSubrun{
		Name:       "ack-after-join",
		AckAfter:   true,
		DBPath:     dbPath,
		RelaySpace: "relay-test",
	}

	result, err := collectRelayOpenMLSSubrunResult(sub)
	if err != nil {
		t.Fatalf("collect result: %v", err)
	}

	if result.Name != "ack-after-join" {
		t.Fatalf("Name = %q", result.Name)
	}
	if !result.AckAfter {
		t.Fatal("AckAfter = false, want true")
	}
	if result.RelaySpaceID != "relay-test" {
		t.Fatalf("RelaySpaceID = %q", result.RelaySpaceID)
	}
	if result.KeyPackageDeliveryState != "queued" {
		t.Fatalf("KeyPackageDeliveryState = %q", result.KeyPackageDeliveryState)
	}
	if result.WelcomeDeliveryState != "acknowledged" {
		t.Fatalf("WelcomeDeliveryState = %q", result.WelcomeDeliveryState)
	}
	if result.EnvelopeCount != 2 {
		t.Fatalf("EnvelopeCount = %d", result.EnvelopeCount)
	}
	if result.AckCount != 1 {
		t.Fatalf("AckCount = %d", result.AckCount)
	}
	if result.WelcomeAckRows != 1 {
		t.Fatalf("WelcomeAckRows = %d", result.WelcomeAckRows)
	}
}

func TestRefuseExistingSidecarDevice(t *testing.T) {
	root := t.TempDir()
	runner := &Runner{
		Comms: filepath.Join(root, "carbonstack-comms"),
	}

	label := "carbonstack-test-device"
	devicePath := filepath.Join(
		runner.Comms,
		"internal",
		"protocol",
		"mls",
		"openmls-sidecar",
		".carbonstack-openmls-sidecar-state",
		"dev",
		"devices",
		label,
	)

	if err := runner.refuseExistingSidecarDevice(label); err != nil {
		t.Fatalf("fresh label refusal check failed: %v", err)
	}

	if err := os.MkdirAll(devicePath, 0700); err != nil {
		t.Fatalf("create stale sidecar device path: %v", err)
	}

	err := runner.refuseExistingSidecarDevice(label)
	if err == nil {
		t.Fatal("expected stale sidecar device path refusal")
	}
	if !strings.Contains(err.Error(), "refuse stale sidecar device state") {
		t.Fatalf("unexpected refusal error: %v", err)
	}
	if !strings.Contains(err.Error(), label) {
		t.Fatalf("refusal error does not include label %q: %v", label, err)
	}
}

func createRelayOpenMLSTestDB(t *testing.T, dbPath string) {
	t.Helper()

	statements := []string{
		"CREATE TABLE envelopes (envelope_id TEXT PRIMARY KEY, content_type TEXT NOT NULL, delivery_state TEXT NOT NULL);",
		"CREATE TABLE envelope_acks (ack_id TEXT PRIMARY KEY, envelope_id TEXT NOT NULL);",
		"INSERT INTO envelopes (envelope_id, content_type, delivery_state) VALUES ('kp1', 'carbonstack.mls.keypackage.v0', 'queued');",
		"INSERT INTO envelopes (envelope_id, content_type, delivery_state) VALUES ('welcome1', 'carbonstack.mls.welcome.v0', 'acknowledged');",
		"INSERT INTO envelope_acks (ack_id, envelope_id) VALUES ('ack1', 'welcome1');",
	}

	for _, statement := range statements {
		if _, err := relayOpenMLSSQLiteScalar(dbPath, statement); err != nil {
			t.Fatalf("sqlite statement failed: %v\nstatement: %s", err, statement)
		}
	}
}
