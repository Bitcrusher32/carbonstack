package main

import (
	"fmt"
	"path/filepath"
)

func (r *Runner) GateBRelayLifecycleClosureDev() error {
	fmt.Println("========================================")
	fmt.Println("CarbonStack validation profile: gate-b-relay-lifecycle-closure-dev")
	fmt.Println("========================================")
	fmt.Println("status: dev/pre-alpha Gate B Relay lifecycle closure profile")
	fmt.Println("scope: B1-B8 integration coherence and closure evidence")
	fmt.Println("boundary: closes Gate B only; not Gate C, not full-runtime-dev, not production E2EE, not verified identity, not trust promotion")
	fmt.Println()

	if err := r.CheckRequiredPaths(); err != nil {
		return err
	}
	if err := r.CheckLiveGitUmbrella("gate-b-relay-lifecycle-closure-dev"); err != nil {
		return err
	}

	validatorDir := filepath.Join(r.CarbonStack, "tools", "carbonstack-validate")

	r.ArtifactScan("pre-gate-b-relay-lifecycle-closure-dev")

	steps := []Step{
		{
			Name:    "Relay Space invite claim regression",
			Dir:     validatorDir,
			Command: "go",
			Args:    []string{"run", ".", "--profile", "relay-space-invite-claim-dev", "--compact-summary", "--clean-generated"},
		},
		{
			Name:    "Relay Space member state regression",
			Dir:     validatorDir,
			Command: "go",
			Args:    []string{"run", ".", "--profile", "relay-space-member-state-dev", "--compact-summary", "--clean-generated"},
		},
		{
			Name:    "Relay Space member restart regression",
			Dir:     validatorDir,
			Command: "go",
			Args:    []string{"run", ".", "--profile", "relay-space-member-restart-dev", "--compact-summary", "--clean-generated"},
		},
		{
			Name:    "Relay Space scoped delivery/ACK authority regression",
			Dir:     validatorDir,
			Command: "go",
			Args:    []string{"run", ".", "--profile", "relay-space-delivery-authority-dev", "--compact-summary", "--clean-generated"},
		},
		{
			Name:    "KeyPackage inspection regression",
			Dir:     validatorDir,
			Command: "go",
			Args:    []string{"run", ".", "--profile", "keypackage-inspect-dev", "--compact-summary", "--clean-generated"},
		},
		{
			Name:    "KeyPackage generation/rotation regression",
			Dir:     validatorDir,
			Command: "go",
			Args:    []string{"run", ".", "--profile", "keypackage-rotation-dev", "--compact-summary", "--clean-generated"},
		},
		{
			Name:    "KeyPackage publication regression",
			Dir:     validatorDir,
			Command: "go",
			Args:    []string{"run", ".", "--profile", "keypackage-publication-dev", "--compact-summary", "--clean-generated"},
		},
		{
			Name:    "KeyPackage consume/receipt regression",
			Dir:     validatorDir,
			Command: "go",
			Args:    []string{"run", ".", "--profile", "keypackage-consume-dev", "--compact-summary", "--clean-generated"},
		},
		{
			Name:    "Welcome lifecycle regression",
			Dir:     validatorDir,
			Command: "go",
			Args:    []string{"run", ".", "--profile", "welcome-lifecycle-dev", "--compact-summary", "--clean-generated"},
		},
		{
			Name:    "Cypher/MLS mismatch refusal regression",
			Dir:     validatorDir,
			Command: "go",
			Args:    []string{"run", ".", "--profile", "cypher-mls-mismatch-dev", "--compact-summary", "--clean-generated"},
		},
		{
			Name:    "Workflow relay onboarding evaluator regression",
			Dir:     validatorDir,
			Command: "go",
			Args:    []string{"run", ".", "--profile", "workflow-relay-onboarding-dev", "--compact-summary", "--clean-generated"},
		},
		{
			Name:    "Legacy lower-level Relay OpenMLS join regression",
			Dir:     validatorDir,
			Command: "go",
			Args:    []string{"run", ".", "--profile", "relay-openmls-join-dev", "--compact-summary", "--clean-generated"},
		},
		{
			Name:    "Same-state Welcome join failure regression",
			Dir:     validatorDir,
			Command: "go",
			Args:    []string{"run", ".", "--profile", "same-state-welcome-join-failure-dev", "--compact-summary", "--clean-generated"},
		},
		{
			Name:    "Same-state integrated normal-message regression",
			Dir:     validatorDir,
			Command: "go",
			Args:    []string{"run", ".", "--profile", "same-state-integrated-dev", "--compact-summary", "--clean-generated"},
		},
		{
			Name:    "Dev runtime OpenMLS wrappers regression",
			Dir:     validatorDir,
			Command: "go",
			Args:    []string{"run", ".", "--profile", "dev-runtime-openmls-wrappers", "--compact-summary", "--clean-generated"},
		},
		{
			Name:    "Generated command reference current",
			Dir:     r.CarbonStack,
			Command: "python3",
			Args:    []string{"tools/registry/render-command-reference.py", "--check"},
		},
		{
			Name:    "Registry missing nonclaims remains zero",
			Dir:     validatorDir,
			Command: "go",
			Args:    []string{"run", ".", "--profile", "registry-lookup", "--list", "--missing-nonclaims"},
		},
	}

	for _, step := range steps {
		if err := r.RunStep(step); err != nil {
			return err
		}
	}

	fmt.Println()
	fmt.Println("gate-b-relay-lifecycle-closure-dev profile result:")
	fmt.Println("  PASS: B1-B4 Relay Space routing/member lifecycle regressions passed")
	fmt.Println("  PASS: B5 KeyPackage inspection/generation/publication/consume lifecycle regressions passed")
	fmt.Println("  PASS: B6 Welcome consume/join/ACK-after-join lifecycle regression passed")
	fmt.Println("  PASS: B7 Cypher/MLS mismatch inspection/refusal regression passed")
	fmt.Println("  PASS: B8 workflow report/evaluator regression passed")
	fmt.Println("  PASS: lower-level Relay OpenMLS join and same-state normal-message regressions remain valid")
	fmt.Println("  PASS: registry/reference current and missing-nonclaims check passed")
	fmt.Println("  gate_b_closed: true")
	fmt.Println("  gate_c_started: false")
	fmt.Println("  full_runtime_dev_promoted: false")
	fmt.Println("  production_e2ee_claimed: false")
	fmt.Println("  verified_identity_claimed: false")
	fmt.Println("  trust_or_candidate_state_mutated: false")
	fmt.Println("  cypher_mls_reconciled: false")
	fmt.Println("  deployment_claimed: false")
	fmt.Println("  container_claimed: false")
	fmt.Println("  pq_claimed: false")
	fmt.Println("  vault_backup_restore_claimed: false")
	fmt.Println("  android_or_carbonstackos_claimed: false")
	fmt.Println("  boundary: dev/pre-alpha Gate B Relay lifecycle closure only; next gate is Gate C state enforcement")

	r.ArtifactScan("post-gate-b-relay-lifecycle-closure-dev")
	if r.CleanGenerated {
		r.CleanGeneratedArtifacts()
	}
	fmt.Println()
	fmt.Println("VALIDATION PASSED")
	return nil
}
