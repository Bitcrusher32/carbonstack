package main

import (
	"fmt"
	"path/filepath"
)

func (r *Runner) WelcomeLifecycleDev() error {
	fmt.Println("========================================")
	fmt.Println("CarbonStack validation profile: welcome-lifecycle-dev")
	fmt.Println("========================================")
	fmt.Println("status: dev/pre-alpha Welcome lifecycle profile")
	fmt.Println("scope: B6 Welcome persistence -> sidecar join -> ACK-after-join")
	fmt.Println("boundary: Welcome lifecycle only; no verified identity, no trust promotion, no B7 Cypher/MLS reconciliation")
	fmt.Println()

	if err := r.CheckRequiredPaths(); err != nil {
		return err
	}
	if err := r.CheckLiveGitUmbrella("welcome-lifecycle-dev"); err != nil {
		return err
	}

	validatorDir := filepath.Join(r.CarbonStack, "tools", "carbonstack-validate")

	r.ArtifactScan("pre-welcome-lifecycle-dev")

	steps := []Step{
		{
			Name:    "Comms B6 Welcome consume command tests",
			Dir:     r.Comms,
			Command: "go",
			Args:    []string{"test", "./internal/app", "-run", "WelcomeConsume", "-count=1", "-v"},
		},
		{
			Name:    "Relay OpenMLS join positive path profile",
			Dir:     validatorDir,
			Command: "go",
			Args:    []string{"run", ".", "--profile", "relay-openmls-join-dev", "--compact-summary", "--clean-generated"},
		},
		{
			Name:    "Same-state corrupt Welcome join failure profile",
			Dir:     validatorDir,
			Command: "go",
			Args:    []string{"run", ".", "--profile", "same-state-welcome-join-failure-dev", "--compact-summary", "--clean-generated"},
		},
		{
			Name:    "B5d KeyPackage receipt nonclaim regression",
			Dir:     validatorDir,
			Command: "go",
			Args:    []string{"run", ".", "--profile", "keypackage-consume-dev", "--compact-summary", "--clean-generated"},
		},
	}

	for _, step := range steps {
		if err := r.RunStep(step); err != nil {
			return err
		}
	}

	fmt.Println()
	fmt.Println("welcome-lifecycle-dev profile result:")
	fmt.Println("  PASS: Welcome artifact is persisted before sidecar conversation-join")
	fmt.Println("  PASS: ACK is sent only after persisted local join evidence")
	fmt.Println("  PASS: corrupt Welcome join failure does not ACK or poison final state")
	fmt.Println("  PASS: existing relay add-member/join positive path remains valid")
	fmt.Println("  PASS: B5d KeyPackage receipt remains not-add-member and not-Welcome-lifecycle")
	fmt.Println("  local_welcome_persisted: true")
	fmt.Println("  joined_before_ack: true")
	fmt.Println("  ack_after_join: true")
	fmt.Println("  no_ack_after_failed_join: true")
	fmt.Println("  trust_or_candidate_state_mutated: false")
	fmt.Println("  verified_identity_claimed: false")
	fmt.Println("  cypher_mls_reconciled: false")
	fmt.Println("  public_directory_mutated: false")
	fmt.Println("  boundary: dev/pre-alpha B6 Welcome lifecycle proof; not B7 reconciliation, not trust promotion, not production E2EE")

	r.ArtifactScan("post-welcome-lifecycle-dev")
	if r.CleanGenerated {
		r.CleanGeneratedArtifacts()
	}
	fmt.Println()
	fmt.Println("VALIDATION PASSED")
	return nil
}
