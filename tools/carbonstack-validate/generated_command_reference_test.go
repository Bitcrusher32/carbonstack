package main

import (
	"os/exec"
	"path/filepath"
	"testing"
)

func TestGeneratedCommandReferenceIsCurrent(t *testing.T) {
	cwd, err := filepath.Abs(".")
	if err != nil {
		t.Fatalf("resolve cwd: %v", err)
	}

	carbonstackRoot := filepath.Clean(filepath.Join(cwd, "..", ".."))
	cmd := exec.Command("python3", "tools/registry/render-command-reference.py", "--check")
	cmd.Dir = carbonstackRoot

	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("generated command reference is stale or renderer failed: %v\n%s", err, string(output))
	}
}
