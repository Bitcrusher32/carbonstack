package main

import (
	"fmt"
	"path/filepath"
)

func (r *Runner) WorkflowRelayOnboardingDev() error {
	fmt.Println("========================================")
	fmt.Println("CarbonStack validation profile: workflow-relay-onboarding-dev")
	fmt.Println("========================================")
	fmt.Println("status: dev/pre-alpha reusable workflow report/evaluator profile")
	fmt.Println("scope: B8 explicit staged relay onboarding workflow evaluation")
	fmt.Println("boundary: no silent repair, no silent rejoin, no trust promotion, no verified identity, no B9 Gate B closure")
	fmt.Println()

	if err := r.CheckRequiredPaths(); err != nil {
		return err
	}
	if err := r.CheckLiveGitUmbrella("workflow-relay-onboarding-dev"); err != nil {
		return err
	}

	validatorDir := filepath.Join(r.CarbonStack, "tools", "carbonstack-validate")

	r.ArtifactScan("pre-workflow-relay-onboarding-dev")

	steps := []Step{
		{
			Name:    "Comms B8 workflow command tests",
			Dir:     r.Comms,
			Command: "go",
			Args:    []string{"test", "./internal/app", "-run", "WorkflowRelayOnboarding", "-count=1", "-v"},
		},
		{
			Name:    "B7 mismatch refusal regression",
			Dir:     validatorDir,
			Command: "go",
			Args:    []string{"run", ".", "--profile", "cypher-mls-mismatch-dev", "--compact-summary", "--clean-generated"},
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
	}

	for _, step := range steps {
		if err := r.RunStep(step); err != nil {
			return err
		}
	}

	fmt.Println()
	fmt.Println("workflow-relay-onboarding-dev profile result:")
	fmt.Println("  PASS: workflow_ready requires aligned B7 state plus persisted/acked KeyPackage and Welcome receipts")
	fmt.Println("  PASS: B7 refusal blocks workflow progression")
	fmt.Println("  PASS: partial onboarding state is refused rather than hidden")
	fmt.Println("  PASS: workflow report is durable and replayable")
	fmt.Println("  PASS: B5/B6/B7 leaf boundaries remain visible in stage output")
	fmt.Println("  no_silent_repair: true")
	fmt.Println("  no_silent_rejoin: true")
	fmt.Println("  leaf_boundaries_preserved: true")
	fmt.Println("  trust_or_candidate_state_mutated: false")
	fmt.Println("  verified_identity_claimed: false")
	fmt.Println("  cypher_mls_reconciled: false")
	fmt.Println("  b9_gate_b_closure_claimed: false")
	fmt.Println("  boundary: dev/pre-alpha B8 workflow report/evaluator; not B9 Gate B closure, not full-runtime-dev, not production E2EE")

	r.ArtifactScan("post-workflow-relay-onboarding-dev")
	if r.CleanGenerated {
		r.CleanGeneratedArtifacts()
	}
	fmt.Println()
	fmt.Println("VALIDATION PASSED")
	return nil
}
