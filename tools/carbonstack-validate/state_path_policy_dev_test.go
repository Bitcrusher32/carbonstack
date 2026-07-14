package main

import "testing"

func TestStatePathPolicyDevProfileName(t *testing.T) {
	r := Runner{Profile: "state-path-policy-dev"}
	if r.Profile != "state-path-policy-dev" {
		t.Fatalf("profile = %q", r.Profile)
	}
}
