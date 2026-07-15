package main

import "testing"

func TestGateENativeDeploymentDevProfileName(t *testing.T) {
	r := Runner{Profile: "gate-e-native-deployment-dev"}
	if r.Profile != "gate-e-native-deployment-dev" {
		t.Fatalf("profile = %q", r.Profile)
	}
}
