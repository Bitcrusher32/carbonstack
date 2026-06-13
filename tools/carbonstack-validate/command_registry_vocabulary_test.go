package main

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

func TestCommandRegistryAudienceAndMaturityVocabulary(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("getwd: %v", err)
	}

	registryPath := filepath.Join(wd, "..", "..", "registry", "commands.v0.yaml")
	raw, err := os.ReadFile(registryPath)
	if err != nil {
		t.Fatalf("read registry: %v", err)
	}

	allowedAudience := map[string]bool{
		"public":   true,
		"dev":      true,
		"internal": true,
		"legacy":   true,
	}

	allowedMaturity := map[string]bool{
		"release_supported": true,
		"experimental":      true,
		"dev_only":          true,
		"legacy":            true,
		"internal":          true,
		"future":            true,
	}

	entryRE := regexp.MustCompile(`(?m)^  - id:\s*(\S+)`)
	audienceRE := regexp.MustCompile(`(?m)^    audience:\s*(\S+)`)
	maturityRE := regexp.MustCompile(`(?m)^    maturity:\s*(\S+)`)

	parts := strings.Split(string(raw), "\n  - id: ")
	for i, part := range parts {
		if i == 0 {
			continue
		}
		block := "  - id: " + part

		idMatch := entryRE.FindStringSubmatch(block)
		if idMatch == nil {
			continue
		}
		id := idMatch[1]

		audienceMatch := audienceRE.FindStringSubmatch(block)
		if audienceMatch == nil {
			t.Fatalf("%s missing audience", id)
		}
		audience := audienceMatch[1]
		if !allowedAudience[audience] {
			t.Fatalf("%s has unsupported audience %q", id, audience)
		}

		maturityMatch := maturityRE.FindStringSubmatch(block)
		if maturityMatch == nil {
			t.Fatalf("%s missing maturity", id)
		}
		maturity := maturityMatch[1]
		if !allowedMaturity[maturity] {
			t.Fatalf("%s has unsupported maturity %q", id, maturity)
		}
	}
}
