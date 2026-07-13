package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func (r *Runner) KeyPackagePublicationDev() error {
	r.PrintHeader("keypackage-publication-dev")
	fmt.Println(
		"status: dev/pre-alpha idempotent KeyPackage publication profile",
	)
	fmt.Println(
		"scope: B5b generation -> B5a inspection -> dedicated Cypher publication authority -> exact replay, conflict, concurrency, and restart proofs",
	)
	fmt.Println(
		"boundary: publication only; no KeyPackage consume or ACK, add-member, Welcome lifecycle, trust/candidate mutation, public directory, or production key distribution",
	)

	if err := r.CheckRequiredPaths(); err != nil {
		return err
	}
	if err := r.CheckLiveGitUmbrella(
		"keypackage-publication-dev",
	); err != nil {
		return err
	}
	r.ArtifactScan("pre-keypackage-publication-dev")

	label := fmt.Sprintf("b5c-publication-%d", time.Now().UnixNano())
	deviceDir := filepath.Join(
		r.Comms,
		"internal",
		"protocol",
		"mls",
		"openmls-sidecar",
		".carbonstack-openmls-sidecar-state",
		"dev",
		"devices",
		label,
	)
	defer os.RemoveAll(deviceDir)

	sidecarDir := filepath.Join(
		r.Comms,
		"internal",
		"protocol",
		"mls",
		"openmls-sidecar",
	)
	if output, code, err := b5aRun(
		sidecarDir,
		"cargo", "run", "--quiet", "--",
		"identity-create",
		"--device-label", label,
	); err != nil || code != 0 {
		return fmt.Errorf(
			"create B5c sidecar identity code=%d err=%v output=%s",
			code, err, output,
		)
	}

	generateOutput, generateCode, generateErr := b5aRun(
		sidecarDir,
		"cargo", "run", "--quiet", "--",
		"keypackage-generate",
		"--device-label", label,
		"--request-id", "publication-generation",
	)
	if generateErr != nil || generateCode != 0 {
		return fmt.Errorf(
			"generate B5c KeyPackage code=%d err=%v output=%s",
			generateCode,
			generateErr,
			generateOutput,
		)
	}
	generated, err := b5aParseEnvelope(generateOutput)
	if err != nil {
		return fmt.Errorf("parse B5c generation: %w", err)
	}
	artifactPath := b5bString(generated.Data, "artifact_path")
	manifestPath := b5bString(generated.Data, "manifest_path")
	keyPackageRef := b5bString(generated.Data, "key_package_ref")
	if artifactPath == "" || manifestPath == "" || keyPackageRef == "" {
		return fmt.Errorf("B5c generation output incomplete: %#v", generated.Data)
	}

	inspectOutput, inspectCode, inspectErr := b5aRun(
		sidecarDir,
		"cargo", "run", "--quiet", "--",
		"keypackage-inspect",
		"--device-label", label,
		"--keypackage", artifactPath,
		"--generation-manifest", manifestPath,
	)
	if inspectErr != nil || inspectCode != 0 {
		return fmt.Errorf(
			"inspect B5c KeyPackage code=%d err=%v output=%s",
			inspectCode,
			inspectErr,
			inspectOutput,
		)
	}
	inspected, err := b5aParseEnvelope(inspectOutput)
	if err != nil {
		return fmt.Errorf("parse B5c inspection: %w", err)
	}
	if !b5bBool(inspected.Data, "valid_at_inspection_time") ||
		!b5bBool(inspected.Data, "openmls_validation_passed") ||
		!b5bBool(inspected.Data, "owner_match") ||
		b5bBool(inspected.Data, "local_state_mutated") ||
		b5bString(inspected.Data, "key_package_ref") != keyPackageRef {
		return fmt.Errorf("B5c inspection contract failed: %#v", inspected.Data)
	}

	if err := b5cRunProfileCommand(
		r.Cypher,
		"go", "test", "./internal/db", "./internal/httpapi",
		"-run", "KeyPackagePublication",
		"-count=1",
	); err != nil {
		return err
	}
	if err := b5cRunProfileCommand(
		r.Comms,
		"go", "test",
		"./internal/client",
		"./internal/relay",
		"./internal/app",
		"-run", "KeyPackagePublication",
		"-count=1",
	); err != nil {
		return err
	}

	fmt.Println()
	fmt.Println("keypackage-publication-dev profile result:")
	fmt.Println(
		"  PASS: one existing B5b generation was validated through B5a inspection",
	)
	fmt.Println(
		"  PASS: dedicated Cypher publication tests proved creation, exact replay, stable envelope identity, destination reuse conflict, and identity conflict",
	)
	fmt.Println(
		"  PASS: real SQLite tests proved concurrent serialization, migration persistence, restart replay, and replay after acknowledgement without requeue",
	)
	fmt.Println(
		"  PASS: Comms tests proved explicit generation selection, pre-publication inspection, exact artifact submission, and retired-generation refusal before network mutation",
	)
	fmt.Println("  keypackage_ref:", keyPackageRef)
	fmt.Println("  keypackage_acked: false")
	fmt.Println("  add_member_run: false")
	fmt.Println("  welcome_submitted: false")
	fmt.Println("  trust_or_candidate_state_mutated: false")
	fmt.Println(
		"  boundary: dev/pre-alpha B5c publication proof; not B5d consume/ACK, not B6 Welcome lifecycle, not identity verification, not production E2EE",
	)

	r.ArtifactScan("post-keypackage-publication-dev")
	return nil
}

func b5cRunProfileCommand(cwd string, command ...string) error {
	if len(command) == 0 {
		return fmt.Errorf("B5c profile command is empty")
	}
	fmt.Println()
	fmt.Println("== B5c profile command ==")
	fmt.Println("cwd:", cwd)
	fmt.Println("run:", strings.Join(command, " "))
	process := exec.Command(command[0], command[1:]...)
	process.Dir = cwd
	output, err := process.CombinedOutput()
	fmt.Print(string(output))
	if err != nil {
		return fmt.Errorf(
			"B5c profile command failed: %s: %w",
			strings.Join(command, " "),
			err,
		)
	}
	return nil
}
