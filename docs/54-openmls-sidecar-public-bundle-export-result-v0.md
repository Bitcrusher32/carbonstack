# CarbonStack OpenMLS Sidecar Public-Bundle-Export Result v0

## Status

Classification: PHASE 2D SIDECAR PUBLIC-BUNDLE-EXPORT RESULT / SUMMARY-ONLY KEYPACKAGE RUNG / PRE-CONVERSATION-LIFECYCLE

This document records the first OpenMLS sidecar public-bundle-export implementation.

This is not final onboarding material.

This is not production secure storage.

This does not write a full serialized KeyPackage artifact.

This does not create conversations.

This does not add members.

This does not join groups.

This does not protect/open messages.

This does not wire OpenMLS into CarbonStackComms user-facing commands.

This does not route MLS traffic through CarbonStackCypher.

This does not mutate trust-state storage.

This does not claim production E2EE.

## Source Context

Relevant docs:

- `docs/49-openmls-sidecar-real-identity-create-plan-v0.md`
- `docs/50-openmls-sidecar-real-identity-create-result-v0.md`
- `docs/51-openmls-sidecar-identity-status-plan-v0.md`
- `docs/52-openmls-sidecar-identity-status-result-v0.md`
- `docs/53-openmls-sidecar-public-bundle-export-plan-v0.md`

Relevant Comms paths:

- `carbonstack-comms/internal/protocol/mls/research/openmls-sidecar/src/main.rs`
- `carbonstack-comms/internal/protocol/mls/research/openmls-sidecar/src/state.rs`
- `carbonstack-comms/internal/protocol/mls/research/openmls-sidecar/src/labels.rs`
- `carbonstack-comms/internal/protocol/openmls_sidecar_provider_info_test.go`
- `carbonstack-comms/internal/protocol/provider_events.go`
- `carbonstack-comms/internal/protocol/provider_events_test.go`
- `carbonstack-comms/internal/protocol/provider_trust.go`
- `carbonstack-comms/internal/protocol/provider_trust_test.go`

## What Was Implemented

The OpenMLS sidecar now supports:

- `public-bundle-export --device-label <label>`

The command performs a narrow summary-only public bundle export rung.

Successful behavior:

- validates the device label
- requires existing identity state
- loads existing dev identity state
- deserializes `signer.json` locally without printing it
- reuses the existing identity-status load/check path
- recreates `BasicCredential`
- recreates `CredentialWithKey`
- builds a real OpenMLS `KeyPackageBundle` in memory
- extracts the public `KeyPackage`
- computes a KeyPackage reference/hash
- writes sanitized `public-bundle-summary.json`
- returns sanitized JSON only
- emits `provider.public_bundle.exported`
- reports `private_material_included: false`

The command does not write a full serialized KeyPackage artifact.

## Implemented Command

Command:

- `public-bundle-export --device-label carbonstack-alice-device`

Current success phase:

- `phase2d-public-bundle-export-dev`

Current success event:

- `provider.public_bundle.exported`

## State Files

For a device label such as:

- `carbonstack-alice-device`

The sidecar uses:

- `carbonstack-comms/internal/protocol/mls/research/openmls-sidecar/.carbonstack-openmls-sidecar-state/dev/devices/carbonstack-alice-device/`

Expected files after identity creation and public bundle summary export:

- `identity-prep.json`
- `identity-summary.json`
- `identity-state.json`
- `signer.json`
- `public-bundle-summary.json`

## Secret Boundary

### Secret-bearing file

`signer.json` remains secret-bearing dev-only local identity material.

Rules:

- it must not be printed
- it must not be pasted into chat
- it must not be committed
- it must remain under ignored sidecar dev state
- it is not production secure storage

### Sanitized public-bundle summary

`public-bundle-summary.json` is sanitized dev-state summary output.

It includes:

- device label
- ciphersuite
- credential type
- public identity reference
- public signature key length
- KeyPackage creation flag
- KeyPackage reference/hash
- KeyPackage hash length
- whether a full KeyPackage artifact was written
- public bundle availability flag
- provider storage flag
- private material flag
- warning

It must not include:

- signer JSON body
- private signing key material
- MemoryStorage JSON
- provider storage JSON
- recovery material
- raw private key bytes
- raw seed material
- secret-bearing file contents

## Current Public Bundle Summary Shape

The current summary file is:

- `public-bundle-summary.json`

Current summary version:

- `public-bundle-summary/v0`

Important fields:

- `key_package_created: true`
- `key_package_ref: sha256:<hex>`
- `key_package_hash_len: 32`
- `key_package_artifact_written: false`
- `public_bundle_available: true`
- `provider_storage_written: false`
- `private_material_included: false`

## KeyPackage Boundary

This rung does generate a real OpenMLS KeyPackage in memory.

This rung does not write the full serialized KeyPackage artifact.

This means:

- KeyPackage generation path is proven.
- A sanitized KeyPackage reference/hash is available.
- The project still does not have a final consumable onboarding artifact.
- The output is not yet enough to implement production add-member/join flows.
- KeyPackage lifecycle semantics are not finalized.

This is intentionally conservative.

## Success Envelope

Successful `public-bundle-export --device-label carbonstack-alice-device` returns:

- `ok: true`
- `command: public-bundle-export`
- `phase: phase2d-public-bundle-export-dev`
- `identity_exists: true`
- `identity_loadable: true`
- `public_bundle_exported: true`
- `public_bundle_available: true`
- `key_package_created: true`
- `key_package_artifact_written: false`
- `provider_storage_written: false`
- `private_material_included: false`
- event: `provider.public_bundle.exported`

The envelope includes:

- device label
- state path hint
- public bundle summary path hint
- public identity reference
- public signature key length
- KeyPackage reference/hash
- KeyPackage hash length

The envelope does not include signer material.

## Missing Identity Behavior

`public-bundle-export --device-label <safe-label>` before identity creation returns:

- nonzero exit
- expected exit code: `3`
- `ok: false`
- error code: `identity_missing`
- provider event: `provider.identity.missing`
- severity: `warning`
- trust relevant: `false`
- `private_material_included: false`

The command does not auto-create identity state.

## Invalid Command Input Behavior

Missing device label:

- command: `public-bundle-export`
- error code: `missing_required_argument`
- provider event: `provider.command.invalid`
- severity: `warning`
- trust relevant: `false`
- private material included: `false`

Invalid device label:

- command: `public-bundle-export`
- error code: `invalid_device_label`
- provider event: `provider.command.invalid`
- severity: `warning`
- trust relevant: `false`
- private material included: `false`

## Unloadable Identity Behavior

If identity state exists but cannot be loaded, the sidecar emits:

- error code: `secret_material_unavailable`
- provider event: `provider.secret.material.unavailable`
- severity: `fatal`
- trust relevant: `true`
- private material included: `false`

This path is meant for malformed, unreadable, mismatched, or deserialization-failing local identity state.

It should not print the malformed secret-bearing file.

It should not auto-repair identity state.

## Provider Events

New event stabilized/used by this rung:

- `provider.public_bundle.exported`

Existing events reused:

- `provider.command.invalid`
- `provider.identity.missing`
- `provider.secret.material.unavailable`
- `checkpoint.failed`

### `provider.public_bundle.exported`

Meaning:

- existing identity state was loaded
- a real OpenMLS KeyPackage was generated in memory
- a sanitized public bundle summary was written
- a KeyPackage reference/hash was returned
- stdout remained sanitized
- no private material was printed
- no serialized KeyPackage artifact was written
- no conversation was created
- no message was protected/opened

Current classification:

- class: `public_setup`
- severity: `info`
- trust relevant: `false`

Current trust-decision behavior:

- append history
- debug only
- do not block send
- do not block receive
- do not block open
- do not require reverify
- not user visible

## Go-Side Validation

Go-side tests now validate:

- `ProviderEventPublicBundleExported` exists.
- `provider.public_bundle.exported` is classified as public setup, info, non-trust-relevant.
- `provider.public_bundle.exported` maps to append-history/debug-only behavior.
- provider-info now lists `public-bundle-export` as a capability.
- unsupported-command tests no longer use `public-bundle-export`.
- unsupported-command tests use a still-unsupported command such as `conversation-create`.
- public-bundle-export before identity returns `identity_missing`.
- identity-create followed by public-bundle-export succeeds.
- public-bundle-export reports `identity_exists: true`.
- public-bundle-export reports `identity_loadable: true`.
- public-bundle-export reports `public_bundle_exported: true`.
- public-bundle-export reports `public_bundle_available: true`.
- public-bundle-export reports `key_package_created: true`.
- public-bundle-export reports `key_package_artifact_written: false`.
- public-bundle-export reports `provider_storage_written: false`.
- public-bundle-export reports `private_material_included: false`.
- public-bundle-export returns a `sha256:` KeyPackage reference.
- public-bundle-export KeyPackage hash length is 32.
- `public-bundle-summary.json` exists.
- `signer.json` exists but is not printed.
- no obvious secret-related tokens appear in stdout.

## Rust-Side Validation

Rust-side validation includes:

- sidecar compiles
- sidecar tests pass
- `public-bundle-export` command is recognized
- missing-label path works
- invalid-label path works
- missing-identity path works
- identity-create then public-bundle-export works
- `public-bundle-summary.json` is written
- `signer.json` exists but was not printed

## Important Implementation Lessons

### Use `OpenMlsRustCrypto` as provider

`KeyPackage::builder().build(...)` requires an `OpenMlsProvider`.

The correct provider type in this dependency set is:

- `openmls_rust_crypto::OpenMlsRustCrypto`

Using `RustCrypto` directly fails because it is only the crypto backend, not the full provider.

### Use `provider.crypto()` directly for KeyPackage hash

`KeyPackage::hash_ref(...)` expects a crypto provider reference.

`provider.crypto()` already returns the correct reference.

Using `&provider.crypto()` creates a reference-to-reference and fails.

### Unsupported list must shrink as commands become real

When `public-bundle-export` became recognized, tests needed to stop treating it as unsupported.

Provider-info now treats it as a capability.

Unsupported-command tests should use a command that is still unsupported, such as:

- `conversation-create`

### Array lengths should be inferred for unsupported command lists

When moving commands from unsupported to supported, fixed array lengths can become stale.

Use a slice form:

- `const UNSUPPORTED_COMMANDS: &[&str] = &[...]`

## Security Boundary

This result is still dev-only.

It does not provide:

- production secure storage
- encrypted local vault
- hardware-backed keys
- key recovery
- revocation
- identity verification UX
- full serialized public KeyPackage artifact
- final onboarding bundle
- conversation lifecycle
- message protect/open
- Comms CLI integration
- Cypher routing
- trust-store mutation
- production E2EE

## Next Recommended Work

Recommended next major step:

- plan serialized public KeyPackage artifact export, or plan conversation-create only after deciding whether full artifact serialization is required first.

Safer next rung:

- inspect and plan full public KeyPackage artifact serialization/export.

Likely future doc:

- `docs/55-openmls-sidecar-keypackage-artifact-export-plan-v0.md`

The next plan should decide:

- exact serialization API
- whether artifact is `.bin`, `.json`, base64, or both
- whether artifact is written to dev state
- whether stdout includes artifact bytes or only a path/hash
- whether the artifact is one-time or dev-reusable
- duplicate/regeneration behavior
- how future conversation-add-member will consume it

Do not proceed directly to conversation lifecycle until the project knows whether the summary-only public bundle is enough or whether a serialized KeyPackage artifact must exist.

## Allowed Claims

Allowed:

- `public-bundle-export` is implemented.
- It loads existing dev-only sidecar identity state.
- It generates a real OpenMLS KeyPackage in memory.
- It computes and returns a sanitized KeyPackage reference/hash.
- It writes sanitized `public-bundle-summary.json`.
- It emits `provider.public_bundle.exported`.
- It keeps `private_material_included: false`.
- It does not write a full serialized KeyPackage artifact.
- It does not create conversations, protect/open messages, route through Cypher, or mutate trust-state storage.

## Not Allowed Claims

Not allowed:

- production secure identity storage exists.
- local signer storage is a secure vault.
- full serialized KeyPackage artifact export exists.
- public bundle output is final onboarding material.
- conversation creation exists.
- add-member/join exists.
- message protect/open exists.
- provider storage persistence is solved.
- Comms CLI consumes OpenMLS.
- Cypher routes MLS payloads.
- trust-state storage consumes sidecar events.
- hostile-server security is validated.
- replay resistance is validated.
- metadata privacy is validated.
- production E2EE exists.
