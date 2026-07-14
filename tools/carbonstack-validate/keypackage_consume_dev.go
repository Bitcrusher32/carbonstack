package main

import "fmt"

func (r *Runner) KeyPackageConsumeDev() error {
	fmt.Println("========================================")
	fmt.Println("CarbonStack validation profile: keypackage-consume-dev")
	fmt.Println("========================================")
	fmt.Println("status: dev/pre-alpha KeyPackage consume/receipt profile")
	fmt.Println("scope: B5c queued KeyPackage envelope -> local durable receipt -> scoped ACK after persistence")
	fmt.Println("boundary: local delivery-consume/receipt only; no add-member, no Welcome lifecycle, no trust/candidate mutation, no production key distribution")
	fmt.Println()

	if err := r.CheckRequiredPaths(); err != nil {
		return err
	}
	if err := r.CheckLiveGitUmbrella("keypackage-consume-dev"); err != nil {
		return err
	}

	r.ArtifactScan("pre-keypackage-consume-dev")

	steps := []Step{
		{
			Name:    "Comms B5d consume command tests",
			Dir:     r.Comms,
			Command: "go",
			Args:    []string{"test", "./internal/app", "-run", "KeyPackageConsume", "-count=1", "-v"},
		},
		{
			Name:    "Comms relay/client adjacent tests",
			Dir:     r.Comms,
			Command: "go",
			Args:    []string{"test", "./internal/client", "./internal/relay", "-run", "RelaySpace|KeyPackage|Ack|Envelope", "-count=1"},
		},
	}

	for _, step := range steps {
		if err := r.RunStep(step); err != nil {
			return err
		}
	}

	fmt.Println()
	fmt.Println("keypackage-consume-dev profile result:")
	fmt.Println("  PASS: exact queued KeyPackage envelope selection is required")
	fmt.Println("  PASS: artifact and receipt manifest are persisted before scoped ACK")
	fmt.Println("  PASS: exact replay returns stable local receipt without duplicate network mutation")
	fmt.Println("  PASS: ACK failure leaves persisted unacked receipt for explicit retry")
	fmt.Println("  keypackage_receipt_persisted: true")
	fmt.Println("  ack_after_persist: true")
	fmt.Println("  add_member_run: false")
	fmt.Println("  welcome_submitted: false")
	fmt.Println("  trust_or_candidate_state_mutated: false")
	fmt.Println("  public_directory_mutated: false")
	fmt.Println("  boundary: dev/pre-alpha B5d consume/receipt proof; not B6 Welcome lifecycle, not identity verification, not production E2EE")

	r.ArtifactScan("post-keypackage-consume-dev")
	if r.CleanGenerated {
		r.CleanGeneratedArtifacts()
	}
	fmt.Println()
	fmt.Println("VALIDATION PASSED")
	return nil
}
