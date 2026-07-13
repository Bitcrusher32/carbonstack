package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

type b5aSidecarEnvelope struct {
	OK      bool           `json:"ok"`
	Command string         `json:"command"`
	Data    map[string]any `json:"data"`
	Error   *struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

func (r *Runner) KeyPackageInspectDev() error {
	r.PrintHeader("keypackage-inspect-dev")
	fmt.Println("status: dev/pre-alpha KeyPackage inspection and ownership profile")
	fmt.Println("scope: generate two local KeyPackages, inspect one read-only, prove local ownership evidence, refuse cross-owner and tampered artifacts")
	fmt.Println("boundary: inspection/identity/lifetime/ownership foundation only; not repeatable rotation, Relay publication, consumption, Welcome lifecycle, trust promotion, account binding, or production safety")

	if err := r.CheckRequiredPaths(); err != nil {
		return err
	}
	if err := r.CheckLiveGitUmbrella("keypackage-inspect-dev"); err != nil {
		return err
	}

	sidecarDir := filepath.Join(
		r.Comms,
		"internal",
		"protocol",
		"mls",
		"openmls-sidecar",
	)
	tempRoot, err := os.MkdirTemp("", "carbonstack-keypackage-inspect-dev-*")
	if err != nil {
		return fmt.Errorf("create KeyPackage inspection temp root: %w", err)
	}
	defer os.RemoveAll(tempRoot)

	runID := fmt.Sprintf("%d", time.Now().UnixNano())
	alice := "b5a-alice-" + runID
	bob := "b5a-bob-" + runID

	for _, label := range []string{alice, bob} {
		deviceDir := filepath.Join(
			sidecarDir,
			".carbonstack-openmls-sidecar-state",
			"dev",
			"devices",
			label,
		)
		defer os.RemoveAll(deviceDir)
	}

	r.ArtifactScan("pre-keypackage-inspect-dev")

	aliceArtifact, err := b5aCreateAndExport(sidecarDir, alice)
	if err != nil {
		return err
	}
	bobArtifact, err := b5aCreateAndExport(sidecarDir, bob)
	if err != nil {
		return err
	}

	aliceDeviceDir := filepath.Join(
		sidecarDir,
		".carbonstack-openmls-sidecar-state",
		"dev",
		"devices",
		alice,
	)
	before, err := b5aHashTree(aliceDeviceDir)
	if err != nil {
		return fmt.Errorf("hash Alice sidecar state before inspection: %w", err)
	}

	ownOutput, ownCode, err := b5aRun(
		sidecarDir,
		"cargo",
		"run",
		"--quiet",
		"--",
		"keypackage-inspect",
		"--device-label",
		alice,
		"--keypackage",
		aliceArtifact,
	)
	if err != nil || ownCode != 0 {
		return fmt.Errorf(
			"own KeyPackage inspection failed code=%d err=%v output=%s",
			ownCode,
			err,
			ownOutput,
		)
	}
	ownEnvelope, err := b5aParseEnvelope(ownOutput)
	if err != nil {
		return fmt.Errorf("parse own KeyPackage inspection: %w", err)
	}
	if !ownEnvelope.OK {
		return errors.New("own KeyPackage inspection returned ok=false")
	}
	for key, want := range map[string]bool{
		"valid_at_inspection_time":  true,
		"openmls_validation_passed": true,
		"owner_match":               true,
		"local_state_mutated":       false,
	} {
		got, ok := ownEnvelope.Data[key].(bool)
		if !ok || got != want {
			return fmt.Errorf(
				"own inspection %s = %#v, want %t",
				key,
				ownEnvelope.Data[key],
				want,
			)
		}
	}
	for _, key := range []string{
		"key_package_ref",
		"key_package_artifact_sha256",
		"lifetime_not_before_unix",
		"lifetime_not_after_unix",
		"owner_evidence",
		"identity_binding",
	} {
		if _, ok := ownEnvelope.Data[key]; !ok {
			return fmt.Errorf("own inspection missing %s", key)
		}
	}

	wrapperOutput, wrapperCode, wrapperErr := b5aRun(
		r.Comms,
		"go",
		"run",
		"./cmd/comms",
		"openmls-keypackage-inspect-dev",
		"--sidecar-dir",
		sidecarDir,
		"--sidecar-device-label",
		alice,
		"--keypackage",
		aliceArtifact,
	)
	if wrapperErr != nil || wrapperCode != 0 {
		return fmt.Errorf(
			"Comms KeyPackage inspection wrapper failed code=%d err=%v output=%s",
			wrapperCode,
			wrapperErr,
			wrapperOutput,
		)
	}
	for _, marker := range []string{
		"command: openmls-keypackage-inspect-dev",
		"status: inspected",
		"openmls_validation_passed: true",
		"owner_match: true",
		"local_state_mutated: false",
		"identity_binding: local-sidecar-device-label-only",
	} {
		if !strings.Contains(wrapperOutput, marker) {
			return fmt.Errorf(
				"wrapper output missing %q:\n%s",
				marker,
				wrapperOutput,
			)
		}
	}

	crossOutput, crossCode, _ := b5aRun(
		sidecarDir,
		"cargo",
		"run",
		"--quiet",
		"--",
		"keypackage-inspect",
		"--device-label",
		alice,
		"--keypackage",
		bobArtifact,
	)
	if crossCode == 0 {
		return fmt.Errorf(
			"cross-owner KeyPackage inspection unexpectedly succeeded: %s",
			crossOutput,
		)
	}
	crossEnvelope, err := b5aParseEnvelope(crossOutput)
	if err != nil {
		return fmt.Errorf("parse cross-owner refusal: %w", err)
	}
	if crossEnvelope.Error == nil ||
		crossEnvelope.Error.Code != "keypackage_owner_mismatch" {
		return fmt.Errorf(
			"cross-owner refusal code = %#v, want keypackage_owner_mismatch",
			crossEnvelope.Error,
		)
	}

	tamperedPath := filepath.Join(tempRoot, "tampered.keypackage.bin")
	artifact, err := os.ReadFile(aliceArtifact)
	if err != nil {
		return err
	}
	if len(artifact) == 0 {
		return errors.New("Alice KeyPackage artifact is empty")
	}
	artifact[len(artifact)/2] ^= 0xff
	if err := os.WriteFile(tamperedPath, artifact, 0o600); err != nil {
		return err
	}

	tamperedOutput, tamperedCode, _ := b5aRun(
		sidecarDir,
		"cargo",
		"run",
		"--quiet",
		"--",
		"keypackage-inspect",
		"--device-label",
		alice,
		"--keypackage",
		tamperedPath,
	)
	if tamperedCode == 0 {
		return fmt.Errorf(
			"tampered KeyPackage inspection unexpectedly succeeded: %s",
			tamperedOutput,
		)
	}
	tamperedEnvelope, err := b5aParseEnvelope(tamperedOutput)
	if err != nil {
		return fmt.Errorf("parse tampered refusal: %w", err)
	}
	if tamperedEnvelope.Error == nil {
		return errors.New("tampered refusal has no error object")
	}
	allowedTamperCodes := map[string]bool{
		"keypackage_deserialize_failed": true,
		"keypackage_trailing_data":      true,
		"keypackage_validation_failed":  true,
		"keypackage_owner_mismatch":     true,
	}
	if !allowedTamperCodes[tamperedEnvelope.Error.Code] {
		return fmt.Errorf(
			"unexpected tampered refusal code %q",
			tamperedEnvelope.Error.Code,
		)
	}

	after, err := b5aHashTree(aliceDeviceDir)
	if err != nil {
		return fmt.Errorf("hash Alice sidecar state after inspection: %w", err)
	}
	if before != after {
		return fmt.Errorf(
			"KeyPackage inspection mutated Alice sidecar state: before=%s after=%s",
			before,
			after,
		)
	}

	r.ArtifactScan("post-keypackage-inspect-dev")

	fmt.Println()
	fmt.Println("keypackage-inspect-dev profile result:")
	fmt.Println("  keypackage_ref_is_lifecycle_identity: true")
	fmt.Println("  artifact_sha256_is_transport_integrity_only: true")
	fmt.Println("  openmls_validation_passed: true")
	fmt.Println("  lifetime_metadata_exposed: true")
	fmt.Println("  local_owner_evidence_matched: true")
	fmt.Println("  cross_owner_refused: true")
	fmt.Println("  tampered_artifact_refused: true")
	fmt.Println("  inspection_state_mutated: false")
	fmt.Println("  trust_or_candidate_state_mutated: false")
	fmt.Println("  identity_binding: local-sidecar-device-label-only")
	fmt.Println("  boundary: B5a inspection/identity/lifetime/ownership foundation; B5b rotation remains next")
	return nil
}

func b5aCreateAndExport(sidecarDir string, label string) (string, error) {
	identityOutput, identityCode, identityErr := b5aRun(
		sidecarDir,
		"cargo",
		"run",
		"--quiet",
		"--",
		"identity-create",
		"--device-label",
		label,
	)
	if identityErr != nil || identityCode != 0 {
		return "", fmt.Errorf(
			"identity-create %s failed code=%d err=%v output=%s",
			label,
			identityCode,
			identityErr,
			identityOutput,
		)
	}

	exportOutput, exportCode, exportErr := b5aRun(
		sidecarDir,
		"cargo",
		"run",
		"--quiet",
		"--",
		"public-bundle-export",
		"--device-label",
		label,
		"--write-artifact",
	)
	if exportErr != nil || exportCode != 0 {
		return "", fmt.Errorf(
			"public-bundle-export %s failed code=%d err=%v output=%s",
			label,
			exportCode,
			exportErr,
			exportOutput,
		)
	}
	envelope, err := b5aParseEnvelope(exportOutput)
	if err != nil {
		return "", err
	}
	hint, ok := envelope.Data["key_package_artifact_path_hint"].(string)
	if !ok || hint == "" {
		return "", fmt.Errorf(
			"public-bundle-export %s missing KeyPackage path hint",
			label,
		)
	}
	path := hint
	if !filepath.IsAbs(path) {
		path = filepath.Join(sidecarDir, path)
	}
	path = filepath.Clean(path)
	if _, err := os.Stat(path); err != nil {
		return "", fmt.Errorf("stat exported KeyPackage %s: %w", path, err)
	}
	return path, nil
}

func b5aRun(
	cwd string,
	command string,
	args ...string,
) (string, int, error) {
	cmd := exec.Command(command, args...)
	cmd.Dir = cwd
	output, err := cmd.CombinedOutput()
	if err == nil {
		return string(output), 0, nil
	}
	var exitError *exec.ExitError
	if errors.As(err, &exitError) {
		return string(output), exitError.ExitCode(), err
	}
	return string(output), -1, err
}

func b5aParseEnvelope(output string) (b5aSidecarEnvelope, error) {
	var envelope b5aSidecarEnvelope
	if err := json.Unmarshal([]byte(output), &envelope); err != nil {
		return envelope, fmt.Errorf(
			"parse sidecar JSON: %w: %s",
			err,
			output,
		)
	}
	return envelope, nil
}

func b5aHashTree(root string) (string, error) {
	var paths []string
	err := filepath.WalkDir(root, func(
		path string,
		entry fs.DirEntry,
		err error,
	) error {
		if err != nil {
			return err
		}
		if entry.IsDir() {
			return nil
		}
		relative, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		paths = append(paths, relative)
		return nil
	})
	if err != nil {
		return "", err
	}
	sort.Strings(paths)

	hash := sha256.New()
	for _, relative := range paths {
		path := filepath.Join(root, relative)
		bytes, err := os.ReadFile(path)
		if err != nil {
			return "", err
		}
		hash.Write([]byte(relative))
		hash.Write([]byte{0})
		hash.Write(bytes)
		hash.Write([]byte{0})
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}
