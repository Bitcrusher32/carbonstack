package main

import "testing"

func TestGateDRuntimeAggregateDevProfileName(t *testing.T) {
	r := Runner{Profile: "gate-d-runtime-aggregate-dev"}
	if r.Profile != "gate-d-runtime-aggregate-dev" {
		t.Fatalf("profile = %q", r.Profile)
	}
}
