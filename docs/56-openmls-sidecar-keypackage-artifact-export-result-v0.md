# OpenMLS Sidecar KeyPackage Artifact Export Result v0

Status: Validated dev-sidecar checkpoint
Component: CarbonStackComms / OpenMLS sidecar
Phase: Phase 2D
Plan doc:
- docs/55-openmls-sidecar-keypackage-artifact-export-plan-v0.md

## 1. Summary

This checkpoint implements and validates dev-only full serialized OpenMLS KeyPackage artifact export for the CarbonStack OpenMLS sidecar.

Before this checkpoint, public-bundle-export generated a real OpenMLS KeyPackage in memory, computed a sanitized KeyPackage reference, and wrote public-bundle-summary.json only.

After this checkpoint, public-bundle-export can optionally write a serialized public KeyPackage artifact and sanitized manifest when called with:

    public-bundle-export --device-label <safe-label> --write-artifact

This is still dev-only provider-boundary work. It is not production E2EE, not final CarbonStack onboarding material, not Comms runtime integration, and not Cypher routing.

## 2. Implemented command behavior

Existing summary-only behavior remains:

    public-bundle-export --device-label <safe-label>

Summary-only mode:

- loads existing dev-only identity state;
- generates a real OpenMLS KeyPackage in memory;
- computes key_package_ref as sha256:<hex>;
- writes public-bundle-summary.json;
- returns key_package_artifact_written=false;
- returns public_bundle_manifest_written=false;
- returns private_material_included=false;
- returns provider_storage_written=false.

New artifact behavior:

    public-bundle-export --device-label <safe-label> --write-artifact

Artifact mode:

- loads existing dev-only identity state;
- generates a real OpenMLS KeyPackage in memory;
- serializes the public KeyPackage using tls_codec serialization;
- writes public-bundle.keypackage.bin under ignored dev state;
- computes sha256 over the exact artifact bytes;
- writes public-bundle-manifest.json;
- updates public-bundle-summary.json with artifact metadata;
- returns key_package_artifact_written=true;
- returns public_bundle_manifest_written=true;
- returns key_package_artifact_sha256=sha256:<hex>;
- returns key_package_artifact_size_bytes > 0;
- returns private_material_included=false;
- returns provider_storage_written=false.

## 3. Files written in artifact mode

Under:

    .carbonstack-openmls-sidecar-state/dev/devices/<device-label>/

Artifact mode writes:

    public-bundle.keypackage.bin
    public-bundle-manifest.json
    public-bundle-summary.json

The existing identity files remain:

    identity-prep.json
    identity-summary.json
    identity-state.json
    signer.json

Important:

    signer.json is secret-bearing dev material.
    signer.json must not be printed, pasted, inspected casually, committed, or treated as production secure storage.

## 4. Manifest fields

The artifact manifest uses:

    manifest_version
    device_label
    state_scope
    ciphersuite
    credential_type
    public_identity_ref
    public_signature_key_len
    key_package_ref
    key_package_hash_len
    key_package_artifact
    key_package_artifact_sha256
    key_package_artifact_size_bytes
    provider_storage_written
    private_material_included
    warning

The manifest does not include private key material.

## 5. Duplicate behavior

If public-bundle.keypackage.bin or public-bundle-manifest.json already exists, artifact export refuses overwrite.

Current duplicate behavior returns a public-bundle-export failure envelope with:

    code: public_bundle_export_failed
    provider_event: checkpoint.failed
    exit code: 4

This is acceptable for this rung. A future checkpoint may add a more specific event such as:

    provider.public_bundle.exists

Do not add force-regeneration until lifecycle semantics decide whether KeyPackages are one-time, reusable-dev, or regenerated per enrollment.

## 6. Test coverage added

Go-side contract tests now cover artifact export behavior.

The tests validate:

- identity-create before artifact export;
- public-bundle-export --write-artifact success;
- key_package_artifact_written=true;
- public_bundle_manifest_written=true;
- key_package_artifact_sha256 has sha256: prefix;
- key_package_artifact_size_bytes > 0;
- artifact file exists;
- manifest file exists;
- summary file exists;
- manifest metadata matches envelope metadata;
- summary metadata matches envelope metadata;
- artifact file size matches envelope size;
- provider_storage_written=false;
- private_material_included=false;
- provider.public_bundle.exported remains non-trust-relevant;
- duplicate artifact export fails;
- duplicate failure stdout does not include obvious secret material.

Existing summary-only tests remain and continue to assert:

- key_package_artifact_written=false;
- public bundle summary exists;
- no provider storage claim;
- no private material claim.

## 7. Validation status

Validated locally:

    cargo check
    cargo test
    go test ./internal/protocol
    go test ./...
    scripts/check-no-rust-artifacts.ps1

Manual probe validated:

    identity-create --device-label carbonstack-alice-device
    public-bundle-export --device-label carbonstack-alice-device --write-artifact
    duplicate public-bundle-export --write-artifact refusal
    invalid device label rejection

Artifact guard passed:

    PASS: no tracked Rust/build artifacts found

## 8. Still not validated / not implemented

This checkpoint does not implement:

- final CarbonStack onboarding bundle format;
- conversation-create;
- conversation-add-member;
- conversation-join;
- Welcome artifact handling;
- message-protect;
- message-open;
- Comms CLI runtime use of sidecar identity state;
- Cypher routing of MLS payloads;
- trust.json mutation;
- trust-events.jsonl mutation;
- provider storage persistence;
- secure vault storage;
- hardware-backed identity;
- Android app integration;
- CarbonStackOS work;
- hostile-server proof;
- replay resistance;
- metadata privacy;
- production E2EE.

## 9. Allowed claims after this checkpoint

Allowed:

- The OpenMLS sidecar can create dev-only identity/signing material.
- The OpenMLS sidecar can load existing dev-only identity state.
- The OpenMLS sidecar can generate a real OpenMLS KeyPackage in memory.
- The OpenMLS sidecar can compute a sanitized KeyPackage reference.
- The OpenMLS sidecar can write a serialized public KeyPackage artifact under ignored dev state.
- The OpenMLS sidecar can write a sanitized public-bundle manifest.
- The artifact export path keeps stdout sanitized and reports private_material_included=false.
- The artifact export path keeps provider_storage_written=false.
- Duplicate artifact export refuses overwrite.
- Go-side contract tests cover the artifact export behavior.

Not allowed:

- CarbonStack is production secure.
- CarbonStack is Signal-equivalent.
- Full onboarding is solved.
- Conversation lifecycle exists.
- Message encryption/decryption exists in runtime.
- Cypher routes MLS payloads.
- Trust storage consumes sidecar output.
- signer.json is secure vault storage.
- The current artifact is final CarbonStack public bundle format.

## 10. Recommended next checkpoint

Next safest checkpoint:

    Phase 2D conversation lifecycle planning

Recommended next doc:

    docs/57-openmls-sidecar-conversation-lifecycle-plan-v0.md

That plan should decide:

- whether the serialized KeyPackage artifact is consumed directly or wrapped;
- how add-member / Welcome artifacts are represented;
- whether provider storage persistence becomes required;
- how conversation state is stored in dev mode;
- how provider events map into trust-state history;
- how to avoid premature Comms/Cypher runtime wiring.

Do not jump directly to message protect/open.
Do not jump directly to Cypher routing.
Do not jump directly to trust-store mutation.
