package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestCommandRegistryHasNoTopLevelEntryIDs(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("getwd: %v", err)
	}

	registryPath := filepath.Join(wd, "..", "..", "registry", "commands.v0.yaml")
	raw, err := os.ReadFile(registryPath)
	if err != nil {
		t.Fatalf("read registry: %v", err)
	}

	for i, line := range strings.Split(string(raw), "\n") {
		if strings.HasPrefix(line, "- id: ") {
			t.Fatalf("registry entry at line %d is malformed top-level YAML; expected two-space indentation under entries: %q", i+1, line)
		}
	}
}
