package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestB5aHashTreeIsContentDeterministic(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(
		filepath.Join(root, "a.txt"),
		[]byte("alpha"),
		0o600,
	); err != nil {
		t.Fatal(err)
	}
	first, err := b5aHashTree(root)
	if err != nil {
		t.Fatal(err)
	}
	second, err := b5aHashTree(root)
	if err != nil {
		t.Fatal(err)
	}
	if first != second {
		t.Fatalf("hash changed without mutation: %s != %s", first, second)
	}
	if err := os.WriteFile(
		filepath.Join(root, "a.txt"),
		[]byte("beta"),
		0o600,
	); err != nil {
		t.Fatal(err)
	}
	third, err := b5aHashTree(root)
	if err != nil {
		t.Fatal(err)
	}
	if third == first {
		t.Fatal("hash did not change after content mutation")
	}
}
