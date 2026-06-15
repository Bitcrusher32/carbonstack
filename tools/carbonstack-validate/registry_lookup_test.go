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
