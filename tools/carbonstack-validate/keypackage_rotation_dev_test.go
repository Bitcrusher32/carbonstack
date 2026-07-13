package main

import "testing"

func TestB5bResolveSidecarPath(t *testing.T) {
	got := b5bResolveSidecarPath("/tmp/sidecar", ".state/device/keypackage.bin")
	want := "/tmp/sidecar/.state/device/keypackage.bin"
	if got != want {
		t.Fatalf("got %q want %q", got, want)
	}
}

func TestB5bOutputContainsNoLifecycleMutation(t *testing.T) {
	if !b5bOutputContainsNoLifecycleMutation("local generation only") {
		t.Fatal("safe output rejected")
	}
	if b5bOutputContainsNoLifecycleMutation("trust promoted") {
		t.Fatal("trust mutation marker accepted")
	}
}
