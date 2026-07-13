package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestB5cProfileSourcePreservesPublicationBoundary(t *testing.T) {
	body, err := os.ReadFile(
		filepath.Join(".", "keypackage_publication_dev.go"),
	)
	if err != nil {
		t.Fatal(err)
	}
	text := string(body)
	for _, marker := range []string{
		"keypackage-publication-dev",
		"keypackage-generate",
		"keypackage-inspect",
		"KeyPackagePublication",
		"keypackage_acked: false",
		"welcome_submitted: false",
		"trust_or_candidate_state_mutated: false",
	} {
		if !strings.Contains(text, marker) {
			t.Fatalf("profile source missing %q", marker)
		}
	}
	for _, forbidden := range []string{
		"conversation-add-member",
		"conversation-join",
		"ackRelaySpaceEnvelope",
	} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("profile source contains forbidden B5d/B6 marker %q", forbidden)
		}
	}
}
