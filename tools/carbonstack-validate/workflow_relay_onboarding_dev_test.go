package main

import "testing"

func TestWorkflowRelayOnboardingDevProfileName(t *testing.T) {
	r := Runner{Profile: "workflow-relay-onboarding-dev"}
	if r.Profile != "workflow-relay-onboarding-dev" {
		t.Fatalf("profile = %q", r.Profile)
	}
}
