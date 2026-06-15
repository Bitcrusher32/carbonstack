package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func (r *Runner) RunIntegratedRuntimeDev(cleanGenerated bool) error {
	r.PrintHeader("integrated-runtime-dev")
	fmt.Println("status: dev/pre-alpha integrated runtime composition profile")
	fmt.Println("scope: Relay onboarding proof followed by recommended message-wrapper runtime proof")
	fmt.Println("proof: relay-openmls-join-dev -> dev-runtime-openmls-wrappers")
	fmt.Println("boundary: live umbrella only; not full; not release-snapshot; not package-root validation; not production/security proof")
	fmt.Println("relationship: composes existing validated dev profiles without replacing individual primitive surfaces")
	fmt.Println("nonclaims: not local-backbone, not production secure messaging, not identity verification, not hostile-server safety, not metadata privacy, not mature messenger UX")
	fmt.Println()

	if err := r.CheckRequiredPaths(); err != nil {
		return err
	}
	if err := r.CheckLiveGitUmbrella("integrated-runtime-dev"); err != nil {
		return err
	}

	r.ArtifactScan("pre-integrated-runtime-dev")

	if err := r.runIntegratedRuntimeSubprofile("relay-openmls-join-dev", cleanGenerated); err != nil {
		return err
	}
	if err := r.runIntegratedRuntimeSubprofile("dev-runtime-openmls-wrappers", cleanGenerated); err != nil {
		return err
	}

	r.ArtifactScan("post-integrated-runtime-dev")

	fmt.Println("integrated-runtime-dev profile result:")
	fmt.Println("  PASS: Relay onboarding proof completed through relay-openmls-join-dev")
	fmt.Println("  PASS: normal message wrapper proof completed through dev-runtime-openmls-wrappers")
	fmt.Println("  proof: relay-openmls-join-dev -> dev-runtime-openmls-wrappers")
	fmt.Println("  boundary: live-dev composition profile; not full, not release-snapshot, not package-root validation, not production/security proof")
	fmt.Println("  relationship: individual Relay onboarding and message wrapper profiles remain separately callable")
	fmt.Println("  note: this first integrated profile composes existing profile proofs in sequence; it does not yet claim a same-state/same-conversation package-root release proof")

	if cleanGenerated {
		if err := r.CleanGeneratedArtifacts(); err != nil {
			return err
		}
	}

	return nil
}

func (r *Runner) runIntegratedRuntimeSubprofile(profile string, cleanGenerated bool) error {
	args := []string{
		"--profile", profile,
		"--root", r.UmbrellaRoot,
	}
	if cleanGenerated {
		args = append(args, "--clean-generated")
	}

	fmt.Println()
	fmt.Println("========================================")
	fmt.Println("Integrated runtime subprofile:", profile)
	fmt.Println("========================================")
	fmt.Println("command:", os.Args[0], strings.Join(args, " "))

	cmd := exec.Command(os.Args[0], args...)
	cmd.Dir = filepath.Join(r.UmbrellaRoot, "carbonstack", "tools", "carbonstack-validate")
	cmd.Env = os.Environ()
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("integrated-runtime-dev subprofile %s failed: %w", profile, err)
	}

	fmt.Println("PASS: integrated runtime subprofile", profile)
	return nil
}
