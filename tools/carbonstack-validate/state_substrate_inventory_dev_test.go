package main

import "testing"

func TestStateSubstrateInventoryDevProfileName(t *testing.T) {
	r := Runner{Profile: "state-substrate-inventory-dev"}
	if r.Profile != "state-substrate-inventory-dev" {
		t.Fatalf("profile = %q", r.Profile)
	}
}
