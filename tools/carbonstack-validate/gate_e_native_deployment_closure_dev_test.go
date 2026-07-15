package main

import "testing"

func TestGateENativeDeploymentClosureDevProfileName(t *testing.T) {
	r := Runner{Profile: "gate-e-native-deployment-closure-dev"}
	if r.Profile != "gate-e-native-deployment-closure-dev" {
		t.Fatalf("profile = %q", r.Profile)
	}
}
