package main

import "testing"

func TestKeyPackageConsumeDevProfileName(t *testing.T) {
	r := Runner{Profile: "keypackage-consume-dev"}
	if r.Profile != "keypackage-consume-dev" {
		t.Fatalf("profile = %q", r.Profile)
	}
}
