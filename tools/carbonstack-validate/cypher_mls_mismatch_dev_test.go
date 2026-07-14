package main

import "testing"

func TestCypherMLSMismatchDevProfileName(t *testing.T) {
	r := Runner{Profile: "cypher-mls-mismatch-dev"}
	if r.Profile != "cypher-mls-mismatch-dev" {
		t.Fatalf("profile = %q", r.Profile)
	}
}
