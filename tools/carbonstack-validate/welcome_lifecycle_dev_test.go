package main

import "testing"

func TestWelcomeLifecycleDevProfileName(t *testing.T) {
	r := Runner{Profile: "welcome-lifecycle-dev"}
	if r.Profile != "welcome-lifecycle-dev" {
		t.Fatalf("profile = %q", r.Profile)
	}
}
