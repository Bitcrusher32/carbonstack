package main

import "testing"

func TestStateWritePolicyDevProfileName(t *testing.T) {
	r := Runner{Profile: "state-write-policy-dev"}
	if r.Profile != "state-write-policy-dev" {
		t.Fatalf("profile = %q", r.Profile)
	}
}
