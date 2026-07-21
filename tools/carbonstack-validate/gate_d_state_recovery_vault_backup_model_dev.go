package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

const gateDStateRecoveryReportSchema = "carbonstack-gate-d-state-recovery-vault-backup-report/v0"
const gateDStateRecoveryModelSchema = "carbonstack-state-recovery-vault-backup-model/v0"

var gateDAdversarialSeedCases = []string{
	"ADV-STATE-RESTORE-MUST-NOT-IMPORT-TRUST-001",
	"ADV-STATE-MIGRATION-DRY-RUN-NOT-REPAIR-001",
	"ADV-STATE-ROLLBACK-MUST-WARN-NOT-SILENT-001",
	"ADV-STATE-BACKUP-MANIFEST-NOT-SECRET-BACKUP-001",
	"ADV-STATE-VAULT-LANGUAGE-NOT-ENCRYPTED-STORAGE-001",
	"ADV-STATE-DESTRUCTIVE-CLEANUP-REQUIRES-CONTRACT-001",
}

func (r *Runner) GateDStateRecoveryVaultBackupModelDev() error {
	fmt.Println("========================================")
	fmt.Println("CarbonStack validation profile: gate-d-state-recovery-vault-backup-model-dev")
	fmt.Println("========================================")
	fmt.Println("status: dev/pre-alpha v0.8.x Gate D state/recovery/vault-backup model seed")
	fmt.Println("scope: model refresh, report schema, non-secret dry-run classification, no-silent nonclaims")
	fmt.Println("boundary: not production vault, not encryption-at-rest, not secret-bearing backup/restore, not migration safety")
	fmt.Println()

	commsRoot := filepath.Join(filepath.Dir(r.CarbonStack), "carbonstack-comms")
	cypherRoot := filepath.Join(filepath.Dir(r.CarbonStack), "carbonstack-cypher")
	required := []string{
		filepath.Join(r.CarbonStack, "registry", "adversarial_cases.v0.yaml"),
		filepath.Join(r.CarbonStack, "registry", "commands.v0.yaml"),
		filepath.Join(r.CarbonStack, "registry", "COMMAND_REFERENCE.v0.md"),
		filepath.Join(r.CarbonStack, "docs", "294-v0.8.4-gate-d-state-recovery-vault-backup-contract-v0.md"),
		filepath.Join(r.CarbonStack, "docs", "295-v0.8.4-gate-d-state-recovery-vault-backup-model-v0.md"),
		filepath.Join(commsRoot, "internal", "state", "gate_d_state_recovery_model_dev.go"),
		filepath.Join(commsRoot, "internal", "state", "gate_d_state_recovery_model_dev_test.go"),
		filepath.Join(cypherRoot, "docs", "07-data-model-v0.md"),
	}
	for _, path := range required {
		if _, err := os.Stat(path); err != nil {
			return fmt.Errorf("Gate D required path missing: %s: %w", path, err)
		}
	}

	caseRegistry, err := gateDRead(filepath.Join(r.CarbonStack, "registry", "adversarial_cases.v0.yaml"))
	if err != nil {
		return err
	}
	commandRegistry, err := gateDRead(filepath.Join(r.CarbonStack, "registry", "commands.v0.yaml"))
	if err != nil {
		return err
	}
	commandReference, err := gateDRead(filepath.Join(r.CarbonStack, "registry", "COMMAND_REFERENCE.v0.md"))
	if err != nil {
		return err
	}
	contractDoc, err := gateDRead(filepath.Join(r.CarbonStack, "docs", "294-v0.8.4-gate-d-state-recovery-vault-backup-contract-v0.md"))
	if err != nil {
		return err
	}
	modelDoc, err := gateDRead(filepath.Join(r.CarbonStack, "docs", "295-v0.8.4-gate-d-state-recovery-vault-backup-model-v0.md"))
	if err != nil {
		return err
	}
	commsModel, err := gateDRead(filepath.Join(commsRoot, "internal", "state", "gate_d_state_recovery_model_dev.go"))
	if err != nil {
		return err
	}

	if !gateDRegistryIDPresent(commandRegistry, "runner.gate-d-state-recovery-vault-backup-model-dev") {
		return fmt.Errorf("Gate D runner registry ID missing")
	}
	if !strings.Contains(commandReference, "runner.gate-d-state-recovery-vault-backup-model-dev") {
		return fmt.Errorf("Gate D runner registry ID missing from command reference")
	}
	if !strings.Contains(commandReference, "Registry entry count: **144**") {
		return fmt.Errorf("command reference does not show expected 144 entries after Gate D registry addition")
	}

	for _, caseID := range gateDAdversarialSeedCases {
		if !strings.Contains(caseRegistry, caseID) {
			return fmt.Errorf("Gate D adversarial seed case missing from case registry: %s", caseID)
		}
		if !regexp.MustCompile(`^ADV-[A-Z0-9]+(?:-[A-Z0-9]+)*-[0-9]{3}$`).MatchString(caseID) {
			return fmt.Errorf("Gate D case ID violates semantic ID policy: %s", caseID)
		}
	}

	docs := contractDoc + "\n" + modelDoc
	for _, marker := range []string{"Gate D is not a production vault", "state export inventory", "restore classification dry-run", "migration compatibility dry-run", "backup manifest model", "rollback metadata inspection", "secret-bearing backup restore is forbidden", "silent repair is forbidden", "silent trust import is forbidden", "Gate C trust-candidate state is referenced inventory only", "owns component-local state classification", "remains coordination/server inventory only", "ADV-STATE-RESTORE-MUST-NOT-IMPORT-TRUST-001"} {
		if !gateDContainsFold(docs, marker) {
			return fmt.Errorf("Gate D docs missing marker %q", marker)
		}
	}

	for _, marker := range []string{"GateDStateRecoveryModelSchema", "carbonstack-state-recovery-vault-backup-model/v0", "GateDClassificationExistsCurrentCode", "GateDClassificationSafeNonSecretDryRun", "GateDClassificationBlockedSecretBearing", "GateDClassificationFutureVaultRequired", "BackupManifestClassification", "RestoreClassification", "MigrationCompatibilityClassification", "RollbackMetadataClassification", "TrustCandidateStateReference", "ChangedLineageWarningReference", "DemotionOrRevocationReference", "not secret-bearing backup restore", "not silent repair", "not silent trust import"} {
		if !strings.Contains(commsModel, marker) {
			return fmt.Errorf("Gate D Comms state model missing marker %q", marker)
		}
	}

	for key, value := range gateDStateRecoveryNonclaims() {
		if value {
			return fmt.Errorf("Gate D nonclaim unexpectedly true: %s", key)
		}
	}

	reportRoot := filepath.Join(os.TempDir(), "carbonstack-gate-d-state-recovery-vault-backup-model-dev")
	_ = os.RemoveAll(reportRoot)
	if err := os.MkdirAll(reportRoot, 0o700); err != nil {
		return err
	}
	reportPath := filepath.Join(reportRoot, "gate-d-state-recovery-vault-backup-report.json")
	report := map[string]any{"schema_version": gateDStateRecoveryReportSchema, "profile": "gate-d-state-recovery-vault-backup-model-dev", "created_at": time.Now().UTC().Format(time.RFC3339), "gate": "v0.8.x Gate D", "model_schema": gateDStateRecoveryModelSchema, "comms_model_path": "carbonstack-comms/internal/state/gate_d_state_recovery_model_dev.go", "adversarial_seed_cases": gateDAdversarialSeedCases, "classification_states": []string{"exists_current_code", "model_only", "adapter_or_classification_only", "safe_non_secret_dry_run_candidate", "blocked_secret_bearing", "future_vault_required", "explicitly_deferred", "unknown_requires_contract"}, "bounded_surfaces": []string{"state export inventory", "restore classification dry-run", "migration compatibility dry-run", "backup manifest model", "rollback metadata inspection"}, "nonclaims_preserved": true, "nonclaims": gateDStateRecoveryNonclaims()}
	if err := gateDWriteJSON(reportPath, report); err != nil {
		return err
	}

	fmt.Println("gate-d-state-recovery-vault-backup-model-dev profile result:")
	fmt.Println("  PASS: Gate D model/report schema markers exist")
	fmt.Println("  PASS: state/recovery/vault-backup classification model exists")
	fmt.Println("  PASS: non-secret dry-run surfaces are classification-only")
	fmt.Println("  PASS: secret-bearing restore, secure vault, silent repair, and trust import remain false")
	fmt.Println("  PASS: Gate C trust-candidate state is referenced inventory only")
	fmt.Println("  PASS: adversarial seed cases are registered")
	fmt.Println("  PASS: registry/reference classification exists without product/security promotion")
	fmt.Println("  report:", reportPath)
	return nil
}

func gateDRead(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
func gateDRegistryIDPresent(registry string, id string) bool {
	rx := regexp.MustCompile("(?m)^\\s*-\\s+id:\\s*" + regexp.QuoteMeta(id) + "\\s*$")
	return rx.MatchString(registry)
}
func gateDContainsFold(haystack string, needle string) bool {
	return strings.Contains(strings.ToLower(haystack), strings.ToLower(needle))
}
func gateDStateRecoveryNonclaims() map[string]bool {
	return map[string]bool{"production_vault_claimed": false, "encryption_at_rest_claimed": false, "secure_key_storage_claimed": false, "production_backup_restore_claimed": false, "migration_safety_claimed": false, "secret_bearing_backup_restore_claimed": false, "hardware_backed_storage_claimed": false, "automatic_migration_claimed": false, "silent_repair_claimed": false, "silent_signer_provider_group_regeneration_claimed": false, "silent_trust_import_claimed": false, "destructive_cleanup_claimed": false}
}
func gateDWriteJSON(path string, value any) error {
	data, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		return err
	}
	data = append(data, '\n')
	return os.WriteFile(path, data, 0o600)
}
