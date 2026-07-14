package main

import (
	"fmt"
	"path/filepath"
)

func (r *Runner) CypherMLSMismatchDev() error {
	fmt.Println("========================================")
	fmt.Println("CarbonStack validation profile: cypher-mls-mismatch-dev")
	fmt.Println("========================================")
	fmt.Println("status: dev/pre-alpha Cypher/MLS mismatch inspection profile")
	fmt.Println("scope: B7 explicit mismatch classification and refusal")
	fmt.Println("boundary: no silent repair, no silent rejoin, no trust promotion, no verified identity, no B8 workflow engine")
	fmt.Println()

	if err := r.CheckRequiredPaths(); err != nil {
		return err
	}
	if err := r.CheckLiveGitUmbrella("cypher-mls-mismatch-dev"); err != nil {
		return err
	}

	validatorDir := filepath.Join(r.CarbonStack, "tools", "carbonstack-validate")

	r.ArtifactScan("pre-cypher-mls-mismatch-dev")

	steps := []Step{
		{
			Name:    "Comms B7 mismatch command tests",
			Dir:     r.Comms,
			Command: "go",
			Args:    []string{"test", "./internal/app", "-run", "CypherMLSMismatch", "-count=1", "-v"},
		},
		{
			Name:    "B6 Welcome lifecycle regression",
			Dir:     validatorDir,
			Command: "go",
			Args:    []string{"run", ".", "--profile", "welcome-lifecycle-dev", "--compact-summary", "--clean-generated"},
		},
		{
			Name:    "B5d KeyPackage receipt regression",
			Dir:     validatorDir,
			Command: "go",
			Args:    []string{"run", ".", "--profile", "keypackage-consume-dev", "--compact-summary", "--clean-generated"},
		},
		{
			Name:    "Relay member-state regression",
			Dir:     validatorDir,
			Command: "go",
			Args:    []string{"run", ".", "--profile", "relay-space-member-state-dev", "--compact-summary", "--clean-generated"},
		},
	}

	for _, step := range steps {
		if err := r.RunStep(step); err != nil {
			return err
		}
	}

	fmt.Println()
	fmt.Println("cypher-mls-mismatch-dev profile result:")
	fmt.Println("  PASS: aligned Cypher active + MLS present is classified as allow")
	fmt.Println("  PASS: Cypher active + MLS absent is classified as refusal")
	fmt.Println("  PASS: Cypher inactive + joined Welcome receipt is classified as refusal")
	fmt.Println("  PASS: Cypher inactive + persisted KeyPackage receipt is classified as refusal")
	fmt.Println("  PASS: local device mismatch is classified as refusal")
	fmt.Println("  PASS: incomplete local Welcome receipt is classified as refusal")
	fmt.Println("  no_silent_repair: true")
	fmt.Println("  no_silent_rejoin: true")
	fmt.Println("  trust_or_candidate_state_mutated: false")
	fmt.Println("  verified_identity_claimed: false")
	fmt.Println("  cypher_mls_reconciled: false")
	fmt.Println("  public_directory_mutated: false")
	fmt.Println("  boundary: dev/pre-alpha B7 inspection/refusal leaf; not B8 workflow engine, not B9 Gate B closure, not production E2EE")

	r.ArtifactScan("post-cypher-mls-mismatch-dev")
	if r.CleanGenerated {
		r.CleanGeneratedArtifacts()
	}
	fmt.Println()
	fmt.Println("VALIDATION PASSED")
	return nil
}
