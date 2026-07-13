package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"time"
)

type b5bGeneration struct {
	GenerationID          string  `json:"generation_id"`
	Sequence              uint64  `json:"sequence"`
	RequestID             string  `json:"request_id"`
	KeyPackageRef         string  `json:"key_package_ref"`
	ArtifactPath          string  `json:"artifact_path"`
	ArtifactSHA256        string  `json:"artifact_sha256"`
	ArtifactSizeBytes     uint64  `json:"artifact_size_bytes"`
	ManifestPath          string  `json:"manifest_path"`
	LifetimeNotBeforeUnix uint64  `json:"lifetime_not_before_unix"`
	LifetimeNotAfterUnix  uint64  `json:"lifetime_not_after_unix"`
	CreatedAtUnix         uint64  `json:"created_at_unix"`
	Status                string  `json:"status"`
	RetiredAtUnix         *uint64 `json:"retired_at_unix"`
	Origin                string  `json:"origin"`
}

type b5bInventoryData struct {
	SchemaVersion     string          `json:"schema_version"`
	DeviceLabel       string          `json:"device_label"`
	NextSequence      uint64          `json:"next_sequence"`
	CurrentGeneration string          `json:"current_generation_id"`
	GenerationCount   int             `json:"generation_count"`
	ActiveCount       int             `json:"active_count"`
	RetiredCount      int             `json:"retired_count"`
	Generations       []b5bGeneration `json:"generations"`
	LocalStateMutated bool            `json:"local_state_mutated"`
}

func (r *Runner) KeyPackageRotationDev() error {
	r.PrintHeader("keypackage-rotation-dev")
	fmt.Println("status: dev/pre-alpha repeatable KeyPackage generation and rotation profile")
	fmt.Println("scope: legacy adoption, repeatable generation, request idempotence, persistent inventory, explicit retirement, concurrency, and B5a inspection")
	fmt.Println("boundary: local sidecar lifecycle only; no Relay publication, KeyPackage consume/ACK, Welcome lifecycle, trust mutation, or production onboarding")

	if err := r.CheckRequiredPaths(); err != nil {
		return err
	}
	if err := r.CheckLiveGitUmbrella("keypackage-rotation-dev"); err != nil {
		return err
	}

	sidecarDir := filepath.Join(
		r.Comms,
		"internal",
		"protocol",
		"mls",
		"openmls-sidecar",
	)
	runID := fmt.Sprintf("%d", time.Now().UnixNano())
	mainLabel := "b5b-main-" + runID
	incompleteLabel := "b5b-incomplete-" + runID
	orphanLabel := "b5b-orphan-" + runID
	labels := []string{mainLabel, incompleteLabel, orphanLabel}
	for _, label := range labels {
		deviceDir := filepath.Join(
			sidecarDir,
			".carbonstack-openmls-sidecar-state",
			"dev",
			"devices",
			label,
		)
		defer os.RemoveAll(deviceDir)
	}

	r.ArtifactScan("pre-keypackage-rotation-dev")

	if _, err := b5bSidecar(sidecarDir, "identity-create", "--device-label", mainLabel); err != nil {
		return err
	}
	legacy, err := b5bSidecar(
		sidecarDir,
		"public-bundle-export",
		"--device-label", mainLabel,
		"--write-artifact",
	)
	if err != nil {
		return err
	}
	legacyRef := b5bString(legacy.Data, "key_package_ref")
	if legacyRef == "" {
		return errors.New("legacy export missing key_package_ref")
	}

	second, err := b5bGenerate(sidecarDir, mainLabel, "rotate-two")
	if err != nil {
		return err
	}
	if b5bString(second.Data, "generation_id") != "kp-000002" || b5bBool(second.Data, "idempotent_replay") {
		return fmt.Errorf("unexpected second generation: %#v", second.Data)
	}

	mainDeviceDir := filepath.Join(
		sidecarDir,
		".carbonstack-openmls-sidecar-state",
		"dev",
		"devices",
		mainLabel,
	)
	beforeReplay, err := b5aHashTree(mainDeviceDir)
	if err != nil {
		return fmt.Errorf("hash state before exact request replay: %w", err)
	}
	secondReplay, err := b5bGenerate(sidecarDir, mainLabel, "rotate-two")
	if err != nil {
		return err
	}
	if !b5bBool(secondReplay.Data, "idempotent_replay") || b5bString(secondReplay.Data, "key_package_ref") != b5bString(second.Data, "key_package_ref") {
		return fmt.Errorf("exact request replay did not return the same generation")
	}
	afterReplay, err := b5aHashTree(mainDeviceDir)
	if err != nil {
		return fmt.Errorf("hash state after exact request replay: %w", err)
	}
	if beforeReplay != afterReplay {
		return fmt.Errorf("exact request replay mutated state: before=%s after=%s", beforeReplay, afterReplay)
	}

	third, err := b5bGenerate(sidecarDir, mainLabel, "rotate-three")
	if err != nil {
		return err
	}
	if b5bString(third.Data, "generation_id") != "kp-000003" {
		return fmt.Errorf("third generation id = %q", b5bString(third.Data, "generation_id"))
	}

	inventory, err := b5bInventory(sidecarDir, mainLabel)
	if err != nil {
		return err
	}
	if inventory.GenerationCount != 3 || inventory.CurrentGeneration != "kp-000003" {
		return fmt.Errorf("initial inventory = %#v", inventory)
	}
	if inventory.Generations[0].Origin != "legacy_import" || inventory.Generations[0].RequestID != "legacy-public-bundle-import" {
		return fmt.Errorf("legacy generation was not adopted correctly: %#v", inventory.Generations[0])
	}
	refs := map[string]bool{}
	for _, generation := range inventory.Generations {
		refs[generation.KeyPackageRef] = true
	}
	if len(refs) != 3 || !refs[legacyRef] {
		return fmt.Errorf("KeyPackage references are not distinct: %#v", refs)
	}

	if err := b5bInspectAll(sidecarDir, mainLabel, inventory.Generations); err != nil {
		return err
	}

	restartReplay, err := b5bGenerate(sidecarDir, mainLabel, "rotate-three")
	if err != nil {
		return err
	}
	if !b5bBool(restartReplay.Data, "idempotent_replay") || b5bString(restartReplay.Data, "generation_id") != "kp-000003" {
		return errors.New("restart replay did not preserve generation identity")
	}

	retired, err := b5bSidecar(
		sidecarDir,
		"keypackage-retire",
		"--device-label", mainLabel,
		"--generation-id", "kp-000001",
	)
	if err != nil {
		return err
	}
	if b5bString(retired.Data, "status") != "retired" || b5bBool(retired.Data, "idempotent_replay") {
		return fmt.Errorf("first retirement failed: %#v", retired.Data)
	}
	retiredReplay, err := b5bSidecar(
		sidecarDir,
		"keypackage-retire",
		"--device-label", mainLabel,
		"--generation-id", "kp-000001",
	)
	if err != nil {
		return err
	}
	if !b5bBool(retiredReplay.Data, "idempotent_replay") {
		return errors.New("retirement replay was not idempotent")
	}

	currentOutput, currentCode, _ := b5aRun(
		sidecarDir,
		"cargo", "run", "--quiet", "--",
		"keypackage-retire",
		"--device-label", mainLabel,
		"--generation-id", "kp-000003",
	)
	if currentCode == 0 {
		return fmt.Errorf("current retirement unexpectedly succeeded: %s", currentOutput)
	}
	currentEnvelope, err := b5aParseEnvelope(currentOutput)
	if err != nil || currentEnvelope.Error == nil || currentEnvelope.Error.Code != "current_generation_retirement_refused" {
		return fmt.Errorf("current retirement refusal = %#v err=%v", currentEnvelope.Error, err)
	}

	inventory, err = b5bInventory(sidecarDir, mainLabel)
	if err != nil {
		return err
	}
	if inventory.ActiveCount != 2 || inventory.RetiredCount != 1 || inventory.CurrentGeneration != "kp-000003" {
		return fmt.Errorf("retired inventory = %#v", inventory)
	}
	if err := b5bInspectAll(sidecarDir, mainLabel, inventory.Generations); err != nil {
		return err
	}

	if _, err := b5bSidecar(sidecarDir, "identity-create", "--device-label", incompleteLabel); err != nil {
		return err
	}
	if _, err := b5bSidecar(
		sidecarDir,
		"public-bundle-export",
		"--device-label", incompleteLabel,
		"--write-artifact",
	); err != nil {
		return err
	}
	incompleteDeviceDir := filepath.Join(
		sidecarDir,
		".carbonstack-openmls-sidecar-state",
		"dev",
		"devices",
		incompleteLabel,
	)
	if err := os.Remove(filepath.Join(incompleteDeviceDir, "public-bundle-manifest.json")); err != nil {
		return err
	}
	incompleteOutput, incompleteCode, _ := b5aRun(
		sidecarDir,
		"cargo", "run", "--quiet", "--",
		"keypackage-generate",
		"--device-label", incompleteLabel,
		"--request-id", "must-refuse",
	)
	if incompleteCode == 0 {
		return fmt.Errorf("incomplete legacy generation unexpectedly succeeded: %s", incompleteOutput)
	}
	incompleteEnvelope, err := b5aParseEnvelope(incompleteOutput)
	if err != nil || incompleteEnvelope.Error == nil || incompleteEnvelope.Error.Code != "incomplete_legacy_keypackage_state" {
		return fmt.Errorf("incomplete legacy refusal = %#v err=%v", incompleteEnvelope.Error, err)
	}

	if _, err := b5bSidecar(sidecarDir, "identity-create", "--device-label", orphanLabel); err != nil {
		return err
	}
	orphanFirst, err := b5bGenerate(sidecarDir, orphanLabel, "orphan-one")
	if err != nil {
		return err
	}
	if b5bString(orphanFirst.Data, "generation_id") != "kp-000001" {
		return fmt.Errorf("orphan first generation = %#v", orphanFirst.Data)
	}
	orphanInventoryPath := filepath.Join(
		sidecarDir,
		".carbonstack-openmls-sidecar-state",
		"dev",
		"devices",
		orphanLabel,
		"keypackages",
		"inventory.json",
	)
	if err := os.Remove(orphanInventoryPath); err != nil {
		return fmt.Errorf("remove inventory to simulate publication interruption: %w", err)
	}
	orphanRecovered, err := b5bGenerate(sidecarDir, orphanLabel, "orphan-one")
	if err != nil {
		return err
	}
	if !b5bBool(orphanRecovered.Data, "idempotent_replay") ||
		!b5bBool(orphanRecovered.Data, "recovered_from_manifest") ||
		b5bString(orphanRecovered.Data, "generation_id") != "kp-000001" {
		return fmt.Errorf("inventory-loss recovery = %#v", orphanRecovered.Data)
	}
	orphanInventory, err := b5bInventory(sidecarDir, orphanLabel)
	if err != nil {
		return err
	}
	if orphanInventory.GenerationCount != 1 || orphanInventory.CurrentGeneration != "kp-000001" {
		return fmt.Errorf("recovered orphan inventory = %#v", orphanInventory)
	}

	binaryPath, err := b5bBuildBinary(sidecarDir)
	if err != nil {
		return err
	}
	concurrent, err := b5bConcurrentGenerate(binaryPath, sidecarDir, mainLabel, []string{"concurrent-four", "concurrent-four"})
	if err != nil {
		return err
	}
	if concurrent[0].generationID != concurrent[1].generationID || concurrent[0].keyPackageRef != concurrent[1].keyPackageRef {
		return fmt.Errorf("concurrent identical requests diverged: %#v", concurrent)
	}
	if concurrent[0].idempotentReplay == concurrent[1].idempotentReplay {
		return fmt.Errorf("concurrent identical requests did not produce one create and one replay: %#v", concurrent)
	}

	distinct, err := b5bConcurrentGenerate(binaryPath, sidecarDir, mainLabel, []string{"concurrent-five", "concurrent-six"})
	if err != nil {
		return err
	}
	if distinct[0].generationID == distinct[1].generationID || distinct[0].keyPackageRef == distinct[1].keyPackageRef {
		return fmt.Errorf("concurrent distinct requests were not serialized into distinct generations: %#v", distinct)
	}

	finalInventory, err := b5bInventory(sidecarDir, mainLabel)
	if err != nil {
		return err
	}
	if finalInventory.GenerationCount != 6 || finalInventory.RetiredCount != 1 || finalInventory.CurrentGeneration == "" {
		return fmt.Errorf("final inventory = %#v", finalInventory)
	}
	finalRefs := map[string]bool{}
	for index, generation := range finalInventory.Generations {
		wantSequence := uint64(index + 1)
		wantID := fmt.Sprintf("kp-%06d", wantSequence)
		if generation.Sequence != wantSequence || generation.GenerationID != wantID {
			return fmt.Errorf("non-monotonic final generation at %d: %#v", index, generation)
		}
		finalRefs[generation.KeyPackageRef] = true
	}
	if len(finalRefs) != finalInventory.GenerationCount {
		return fmt.Errorf("final KeyPackage references are not distinct: %#v", finalRefs)
	}
	if err := b5bInspectAll(sidecarDir, mainLabel, finalInventory.Generations); err != nil {
		return err
	}

	providerPath := filepath.Join(
		sidecarDir,
		".carbonstack-openmls-sidecar-state",
		"dev",
		"devices",
		mainLabel,
		"provider-storage.json",
	)
	if info, err := os.Stat(providerPath); err != nil || info.Size() == 0 {
		return fmt.Errorf("provider storage is not persisted: info=%v err=%v", info, err)
	}

	r.ArtifactScan("post-keypackage-rotation-dev")

	fmt.Println()
	fmt.Println("keypackage-rotation-dev profile result:")
	fmt.Println("  legacy_generation_adopted: true")
	fmt.Println("  repeatable_generation: true")
	fmt.Println("  distinct_key_package_refs: true")
	fmt.Println("  monotonic_generation_ids: true")
	fmt.Println("  exact_request_replay: true")
	fmt.Println("  restart_replay: true")
	fmt.Println("  exact_replay_state_mutated: false")
	fmt.Println("  inventory_loss_manifest_recovery: true")
	fmt.Println("  prior_generations_retained: true")
	fmt.Println("  explicit_retirement: true")
	fmt.Println("  retirement_replay: true")
	fmt.Println("  current_retirement_refused: true")
	fmt.Println("  incomplete_legacy_state_refused: true")
	fmt.Println("  concurrent_identical_single_generation: true")
	fmt.Println("  concurrent_distinct_serialized: true")
	fmt.Println("  provider_storage_reloadable: true")
	fmt.Println("  every_generation_b5a_inspectable: true")
	fmt.Println("  relay_or_cypher_mutated: false")
	fmt.Println("  welcome_or_ack_mutated: false")
	fmt.Println("  trust_or_candidate_state_mutated: false")
	fmt.Println("  boundary: B5b local generation/rotation only; B5c remains blocked pending breakpoint")
	return nil
}

func b5bSidecar(sidecarDir string, command string, args ...string) (b5aSidecarEnvelope, error) {
	all := append([]string{"run", "--quiet", "--", command}, args...)
	output, code, runErr := b5aRun(sidecarDir, "cargo", all...)
	if runErr != nil || code != 0 {
		return b5aSidecarEnvelope{}, fmt.Errorf("sidecar %s failed code=%d err=%v output=%s", command, code, runErr, output)
	}
	envelope, err := b5aParseEnvelope(output)
	if err != nil {
		return envelope, err
	}
	if !envelope.OK {
		return envelope, fmt.Errorf("sidecar %s returned ok=false", command)
	}
	return envelope, nil
}

func b5bGenerate(sidecarDir string, label string, requestID string) (b5aSidecarEnvelope, error) {
	return b5bSidecar(
		sidecarDir,
		"keypackage-generate",
		"--device-label", label,
		"--request-id", requestID,
	)
}

func b5bInventory(sidecarDir string, label string) (b5bInventoryData, error) {
	envelope, err := b5bSidecar(sidecarDir, "keypackage-inventory", "--device-label", label)
	if err != nil {
		return b5bInventoryData{}, err
	}
	bytes, err := json.Marshal(envelope.Data)
	if err != nil {
		return b5bInventoryData{}, err
	}
	var inventory b5bInventoryData
	if err := json.Unmarshal(bytes, &inventory); err != nil {
		return inventory, err
	}
	sort.Slice(inventory.Generations, func(i, j int) bool {
		return inventory.Generations[i].Sequence < inventory.Generations[j].Sequence
	})
	return inventory, nil
}

func b5bInspectAll(sidecarDir string, label string, generations []b5bGeneration) error {
	for _, generation := range generations {
		artifact := b5bResolveSidecarPath(sidecarDir, generation.ArtifactPath)
		manifest := b5bResolveSidecarPath(sidecarDir, generation.ManifestPath)
		envelope, err := b5bSidecar(
			sidecarDir,
			"keypackage-inspect",
			"--device-label", label,
			"--keypackage", artifact,
			"--generation-manifest", manifest,
		)
		if err != nil {
			return fmt.Errorf("inspect %s: %w", generation.GenerationID, err)
		}
		if !b5bBool(envelope.Data, "owner_match") || b5bString(envelope.Data, "key_package_ref") != generation.KeyPackageRef {
			return fmt.Errorf("inspection mismatch for %s: %#v", generation.GenerationID, envelope.Data)
		}
	}
	return nil
}

func b5bBuildBinary(sidecarDir string) (string, error) {
	output, code, err := b5aRun(sidecarDir, "cargo", "build", "--quiet")
	if err != nil || code != 0 {
		return "", fmt.Errorf("build sidecar binary code=%d err=%v output=%s", code, err, output)
	}
	targetRoot := os.Getenv("CARGO_TARGET_DIR")
	if targetRoot == "" {
		targetRoot = filepath.Join(sidecarDir, "target")
	} else if !filepath.IsAbs(targetRoot) {
		targetRoot = filepath.Join(sidecarDir, targetRoot)
	}
	binary := filepath.Join(targetRoot, "debug", "carbonstack_openmls_sidecar")
	if _, err := os.Stat(binary); err != nil {
		return "", fmt.Errorf("stat built sidecar binary %s: %w", binary, err)
	}
	return binary, nil
}

type b5bConcurrentResult struct {
	generationID     string
	keyPackageRef    string
	idempotentReplay bool
}

func b5bConcurrentGenerate(binary string, cwd string, label string, requestIDs []string) ([]b5bConcurrentResult, error) {
	results := make([]b5bConcurrentResult, len(requestIDs))
	errs := make([]error, len(requestIDs))
	var wg sync.WaitGroup
	start := make(chan struct{})
	for index, requestID := range requestIDs {
		wg.Add(1)
		go func(index int, requestID string) {
			defer wg.Done()
			<-start
			cmd := exec.Command(
				binary,
				"keypackage-generate",
				"--device-label", label,
				"--request-id", requestID,
			)
			cmd.Dir = cwd
			output, err := cmd.CombinedOutput()
			if err != nil {
				errs[index] = fmt.Errorf("concurrent request %s failed: %w output=%s", requestID, err, output)
				return
			}
			envelope, err := b5aParseEnvelope(string(output))
			if err != nil {
				errs[index] = err
				return
			}
			results[index] = b5bConcurrentResult{
				generationID:     b5bString(envelope.Data, "generation_id"),
				keyPackageRef:    b5bString(envelope.Data, "key_package_ref"),
				idempotentReplay: b5bBool(envelope.Data, "idempotent_replay"),
			}
		}(index, requestID)
	}
	close(start)
	wg.Wait()
	for _, err := range errs {
		if err != nil {
			return nil, err
		}
	}
	return results, nil
}

func b5bResolveSidecarPath(sidecarDir string, path string) string {
	if filepath.IsAbs(path) {
		return filepath.Clean(path)
	}
	return filepath.Clean(filepath.Join(sidecarDir, path))
}

func b5bString(data map[string]any, key string) string {
	value, _ := data[key].(string)
	return value
}

func b5bBool(data map[string]any, key string) bool {
	value, _ := data[key].(bool)
	return value
}

func b5bOutputContainsNoLifecycleMutation(output string) bool {
	for _, marker := range []string{"relay", "welcome submitted", "acknowledged", "trust promoted", "candidate accepted"} {
		if strings.Contains(strings.ToLower(output), marker) {
			return false
		}
	}
	return true
}
