package main

import "testing"

func TestStateSchemaCompatibilityDevProfileName(t *testing.T) {
	r := Runner{Profile: "state-schema-compat-dev"}
	if r.Profile != "state-schema-compat-dev" {
		t.Fatalf("profile = %q", r.Profile)
	}
}
