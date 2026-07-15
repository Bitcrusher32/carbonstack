package main

import "testing"

func TestGateCStateSubstrateClosureDevProfileName(t *testing.T) {
	r := Runner{Profile: "gate-c-state-substrate-closure-dev"}
	if r.Profile != "gate-c-state-substrate-closure-dev" {
		t.Fatalf("profile = %q", r.Profile)
	}
}
