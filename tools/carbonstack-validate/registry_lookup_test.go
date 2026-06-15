package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestParseRegistryLookupEntriesParsesFieldsAndLists(t *testing.T) {
	raw := `entries:
  - id: comms.message-send-dev
    command: go run ./cmd/comms message-send-dev
    repo: carbonstack-comms
    component: cmd/comms
    kind: cli
    audience: dev
    maturity: dev_only
    lifecycle_status: recommended_dev_wrapper
    source_path: carbonstack-comms/internal/app/message_wrappers_dev.go
    short_help: wrapper
    why_exists: test
    nonclaims:
      - dev/pre-alpha only
      - not production E2EE claim
    include_in_front_readme: true
`

	entries := parseRegistryLookupEntries(raw)
	if len(entries) != 1 {
		t.Fatalf("entries = %d, want 1", len(entries))
	}
	entry := entries[0]
	if entry.ID != "comms.message-send-dev" {
		t.Fatalf("id = %q", entry.ID)
	}
	if entry.Fields["lifecycle_status"] != "recommended_dev_wrapper" {
		t.Fatalf("lifecycle_status = %q", entry.Fields["lifecycle_status"])
	}
	if got := entry.Lists["nonclaims"]; len(got) != 2 || got[1] != "not production E2EE claim" {
		t.Fatalf("nonclaims = %#v", got)
	}
}

func TestPrintRegistryLookupEntryIncludesBoundary(t *testing.T) {
	entry := registryLookupEntry{
		ID: "comms.message-send-dev",
		Fields: map[string]string{
			"command":          "go run ./cmd/comms message-send-dev",
			"audience":         "dev",
			"maturity":         "dev_only",
			"lifecycle_status": "recommended_dev_wrapper",
			"short_help":       "wrapper",
			"why_exists":       "test",
		},
		Lists: map[string][]string{
			"nonclaims": {"not production E2EE claim"},
		},
	}

	output := captureStdout(func() {
		printRegistryLookupEntry(entry)
	})

	for _, want := range []string{
		"id: comms.message-send-dev",
		"lifecycle_status: recommended_dev_wrapper",
		"nonclaims:",
		"not production E2EE claim",
		"boundary: registry presence is classification, not promotion",
	} {
		if !strings.Contains(output, want) {
			t.Fatalf("lookup output missing %q\n%s", want, output)
		}
	}
}

func captureStdout(fn func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	fn()

	_ = w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)
	return buf.String()
}

func TestFilterRegistryLookupEntriesByLifecycleStatus(t *testing.T) {
	raw := `entries:
  - id: comms.message-send-dev
    command: go run ./cmd/comms message-send-dev
    audience: dev
    maturity: dev_only
    lifecycle_status: recommended_dev_wrapper
    kind: cli
    nonclaims:
      - not production E2EE claim
  - id: comms.openmls-send-dev
    command: go run ./cmd/comms openmls-send-dev
    audience: dev
    maturity: dev_only
    lifecycle_status: lower_level_direct_proof_transition_candidate
    kind: cli
    nonclaims:
      - not production E2EE claim
`
	entries := parseRegistryLookupEntries(raw)
	matches := filterRegistryLookupEntries(entries, registryLookupOptions{
		List:            true,
		LifecycleStatus: "recommended_dev_wrapper",
	})
	if len(matches) != 1 || matches[0].ID != "comms.message-send-dev" {
		t.Fatalf("matches = %#v", matches)
	}
}

func TestFilterRegistryLookupEntriesMissingNonclaims(t *testing.T) {
	raw := `entries:
  - id: with.nonclaims
    command: with
    audience: dev
    maturity: dev_only
    nonclaims:
      - not production E2EE claim
  - id: without.nonclaims
    command: without
    audience: dev
    maturity: dev_only
`
	entries := parseRegistryLookupEntries(raw)
	matches := filterRegistryLookupEntries(entries, registryLookupOptions{
		List:             true,
		MissingNonclaims: true,
	})
	if len(matches) != 1 || matches[0].ID != "without.nonclaims" {
		t.Fatalf("matches = %#v", matches)
	}
}

func TestPrintRegistryLookupListIncludesFiltersAndBoundary(t *testing.T) {
	entry := registryLookupEntry{
		ID: "comms.message-send-dev",
		Fields: map[string]string{
			"command":          "go run ./cmd/comms message-send-dev",
			"audience":         "dev",
			"maturity":         "dev_only",
			"lifecycle_status": "recommended_dev_wrapper",
			"short_help":       "wrapper",
		},
		Lists: map[string][]string{},
	}
	output := captureStdout(func() {
		printRegistryLookupList([]registryLookupEntry{entry}, registryLookupOptions{
			List:            true,
			LifecycleStatus: "recommended_dev_wrapper",
		})
	})
	for _, want := range []string{
		"registry entries",
		"matches: 1",
		"filter_lifecycle_status: recommended_dev_wrapper",
		"id: comms.message-send-dev",
		"boundary: registry presence is classification, not promotion",
	} {
		if !strings.Contains(output, want) {
			t.Fatalf("lookup list output missing %q\n%s", want, output)
		}
	}
}

func TestFilterRegistryLookupEntriesByRelayOnboardingLifecycleStatus(t *testing.T) {
	raw := `entries:
  - id: comms.openmls-relay-keypackage-submit-dev
    command: go run ./cmd/comms openmls-relay-keypackage-submit-dev
    audience: dev
    maturity: dev_only
    lifecycle_status: relay_onboarding_artifact_transport
    kind: cli
    nonclaims:
      - not production key distribution UX
  - id: comms.openmls-relay-keypackage-inbox-dev
    command: go run ./cmd/comms openmls-relay-keypackage-inbox-dev
    audience: dev
    maturity: dev_only
    lifecycle_status: relay_onboarding_artifact_transport
    kind: cli
    nonclaims:
      - no ack
  - id: comms.openmls-relay-welcome-submit-dev
    command: go run ./cmd/comms openmls-relay-welcome-submit-dev
    audience: dev
    maturity: dev_only
    lifecycle_status: relay_onboarding_artifact_transport
    kind: cli
    nonclaims:
      - not production membership UX
  - id: comms.openmls-relay-welcome-inbox-dev
    command: go run ./cmd/comms openmls-relay-welcome-inbox-dev
    audience: dev
    maturity: dev_only
    lifecycle_status: relay_onboarding_artifact_transport
    kind: cli
    nonclaims:
      - no conversation-join
  - id: comms.openmls-relay-add-member-dev
    command: go run ./cmd/comms openmls-relay-add-member-dev
    audience: dev
    maturity: dev_only
    lifecycle_status: relay_onboarding_artifact_bridge
    kind: cli
    nonclaims:
      - does not ack KeyPackage or Welcome
  - id: comms.openmls-relay-join-dev
    command: go run ./cmd/comms openmls-relay-join-dev
    audience: dev
    maturity: dev_only
    lifecycle_status: relay_onboarding_join_transition_candidate
    kind: cli
    nonclaims:
      - not hostile-server safety
`
	entries := parseRegistryLookupEntries(raw)

	transport := filterRegistryLookupEntries(entries, registryLookupOptions{
		List:            true,
		LifecycleStatus: "relay_onboarding_artifact_transport",
	})
	if len(transport) != 4 {
		t.Fatalf("transport matches = %d, want 4: %#v", len(transport), transport)
	}

	bridge := filterRegistryLookupEntries(entries, registryLookupOptions{
		List:            true,
		LifecycleStatus: "relay_onboarding_artifact_bridge",
	})
	if len(bridge) != 1 || bridge[0].ID != "comms.openmls-relay-add-member-dev" {
		t.Fatalf("bridge matches = %#v", bridge)
	}

	join := filterRegistryLookupEntries(entries, registryLookupOptions{
		List:            true,
		LifecycleStatus: "relay_onboarding_join_transition_candidate",
	})
	if len(join) != 1 || join[0].ID != "comms.openmls-relay-join-dev" {
		t.Fatalf("join matches = %#v", join)
	}
}

func TestPrintRegistryLookupListIncludesRelayOnboardingLifecycleStatus(t *testing.T) {
	entry := registryLookupEntry{
		ID: "comms.openmls-relay-join-dev",
		Fields: map[string]string{
			"command":          "go run ./cmd/comms openmls-relay-join-dev",
			"audience":         "dev",
			"maturity":         "dev_only",
			"lifecycle_status": "relay_onboarding_join_transition_candidate",
			"short_help":       "Consume a Relay Space Welcome, run sidecar conversation-join, and optionally ack only after join succeeds.",
		},
		Lists: map[string][]string{},
	}

	output := captureStdout(func() {
		printRegistryLookupList([]registryLookupEntry{entry}, registryLookupOptions{
			List:            true,
			LifecycleStatus: "relay_onboarding_join_transition_candidate",
		})
	})

	for _, want := range []string{
		"registry entries",
		"matches: 1",
		"filter_lifecycle_status: relay_onboarding_join_transition_candidate",
		"id: comms.openmls-relay-join-dev",
		"lifecycle_status: relay_onboarding_join_transition_candidate",
		"boundary: registry presence is classification, not promotion",
	} {
		if !strings.Contains(output, want) {
			t.Fatalf("relay lifecycle lookup output missing %q\n%s", want, output)
		}
	}
}
