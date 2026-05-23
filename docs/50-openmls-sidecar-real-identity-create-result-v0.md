# CarbonStack OpenMLS Sidecar Real Identity-Create Result v0

## Status

Classification: PHASE 2D SIDECAR REAL IDENTITY-CREATE RESULT / DEV-ONLY SECRET-GENERATION

This document records the first dev-only real OpenMLS identity material generation path in the OpenMLS sidecar.

This is not production identity creation.

This is not production secure storage.

This does not generate KeyPackages.

This does not export public bundles.

This does not create conversations.

This does not protect/open messages.

This does not wire OpenMLS into CarbonStackComms user-facing commands.

This does not route MLS traffic through CarbonStackCypher.

This does not mutate trust-state storage.

This does not claim production E2EE.

## Source Context

Relevant docs:

- `docs/44-openmls-sidecar-identity-create-plan-v0.md`
- `docs/45-openmls-sidecar-identity-create-prep-result-v0.md`
- `docs/46-provider-command-invalid-not-implemented-events-result-v0.md`
- `docs/47-openmls-sidecar-identity-prep-state-result-v0.md`
- `docs/48-provider-identity-prep-state-events-result-v0.md`
- `docs/49-openmls-sidecar-real-identity-create-plan-v0.md`

Relevant Comms paths:

- `carbonstack-comms/internal/protocol/mls/research/openmls-sidecar/Cargo.toml`
- `carbonstack-comms/internal/protocol/mls/research/openmls-sidecar/Cargo.lock`
- `carbonstack-comms/internal/protocol/mls/research/openmls-sidecar/src/main.rs`
- `carbonstack-comms/internal/protocol/mls/research/openmls-sidecar/src/state.rs`
- `carbonstack-comms/internal/protocol/mls/research/openmls-sidecar/src/labels.rs`
- `carbonstack-comms/internal/protocol/openmls_sidecar_provider_info_test.go`
- `carbonstack-comms/internal/protocol/provider_events.go`
- `carbonstack-comms/internal/protocol/provider_events_test.go`
- `carbonstack-comms/internal/protocol/provider_trust.go`
- `carbonstack-comms/internal/protocol/provider_trust_test.go`

## What Was Implemented

The sidecar `identity-create --device-label <safe-label>` path now creates dev-only real OpenMLS identity material.

Successful safe-label behavior now:

- validates the device label
- refuses overwrite if identity/prep state already exists
- creates ignored per-device dev state
- creates an OpenMLS `SignatureKeyPair`
- creates a `BasicCredential`
- creates a `CredentialWithKey`
- writes secret-bearing `signer.json`
- writes non-secret `identity-summary.json`
- writes non-secret `identity-state.json`
- writes updated `identity-prep.json`
- returns a sanitized JSON success envelope
- reports `identity_created: true`
- reports `state_written: true`
- reports `provider_storage_written: false`
- reports `public_bundle_available: false`
- reports `private_material_included: false`

This is the first Phase 2D sidecar rung that generates secret-bearing local dev state.

## Dependency Change

The OpenMLS sidecar now directly depends on the OpenMLS identity material stack required for this rung.

Added or used sidecar dependencies include:

- `openmls`
- `openmls_basic_credential`
- `openmls_rust_crypto`
- `serde`
- `serde_json`
- `sha2`
- `hex`

The sidecar still does not implement message protection, group operations, public bundle export, or provider storage persistence.

## State Files

For a device label such as:

- `carbonstack-alice-device`

The sidecar writes under:

- `carbonstack-comms/internal/protocol/mls/research/openmls-sidecar/.carbonstack-openmls-sidecar-state/dev/devices/carbonstack-alice-device/`

Expected files:

- `identity-prep.json`
- `identity-summary.json`
- `identity-state.json`
- `signer.json`

## Secret Boundary

### Secret-bearing file

`signer.json` is secret-bearing dev-only local identity material.

Rules:

- it must not be printed
- it must not be pasted into chat
- it must not be committed
- it must remain under the ignored sidecar dev-state path
- it is not production secure storage

### Non-secret/sanitized files

The following files are intended as safe summaries for testing/debugging:

- `identity-prep.json`
- `identity-summary.json`
- `identity-state.json`

They must still avoid private key material, signer body, provider storage body, MemoryStorage JSON, recovery material, raw key bytes, and seed material.

## Success Envelope

Successful `identity-create --device-label carbonstack-alice-device` now returns:

- `ok: true`
- `command: identity-create`
- `phase: phase2d-identity-create-dev`
- `identity_created: true`
- `state_written: true`
- `provider_storage_written: false`
- `public_bundle_available: false`
- `private_material_included: false`
- event: `provider.identity.created`

The success envelope includes path hints and a public identity reference, but it does not print `signer.json`.

## Public Identity Reference

The implementation derives a sanitized public identity reference from the public signature key.

Current shape:

- `sha256:<hex digest>`

This is a dev-only reference and should not be mistaken for a final production fingerprint UX.

## Provider Events

This rung stabilizes/uses:

- `provider.identity.created`
- `provider.identity.exists`
- `provider.command.invalid`
- `provider.command.unsupported`
- `checkpoint.failed`

### `provider.identity.created`

Meaning:

- the sidecar generated local dev identity material
- required local state files were written
- stdout remained sanitized
- no public bundle was exported
- no provider storage was written
- no message/conversation state exists

Current classification:

- class: `storage/checkpoint`
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

## Duplicate / Overwrite Refusal

Duplicate safe-label `identity-create` now refuses overwrite.

Expected duplicate behavior:

- exits with code `3`
- returns `ok: false`
- error code: `identity_already_exists`
- provider event: `provider.identity.exists`
- severity: `warning`
- trust relevant: `false`
- reports `identity_created: false`
- reports `state_written: false`
- does not print private material
- does not replace existing identity state

This is critical. Silent identity replacement would be a serious future trust hazard.

## Go-Side Validation

Go-side tests now validate:

- `provider-info` envelope still works
- unsupported command envelope still works
- missing-label `identity-create` still emits `provider.command.invalid`
- invalid-label `identity-create` still emits `provider.command.invalid`
- safe-label `identity-create` exits successfully
- safe-label `identity-create` reports `identity_created: true`
- safe-label `identity-create` reports `state_written: true`
- safe-label `identity-create` reports `provider_storage_written: false`
- safe-label `identity-create` reports `public_bundle_available: false`
- safe-label `identity-create` reports `private_material_included: false`
- safe-label `identity-create` emits `provider.identity.created`
- output includes a `sha256:` public identity reference
- expected state files exist
- `signer.json` exists but is not printed
- summary/state JSON files parse and preserve no-secret claims
- duplicate identity-create refuses overwrite with `identity_already_exists`
- duplicate refusal emits `provider.identity.exists`
- stdout is checked for obvious secret-related tokens

## Rust-Side Validation

Rust-side checks validate:

- sidecar compiles
- sidecar unit tests pass
- label validation remains intact
- state path helper tests pass
- expected identity phase is `phase2d-identity-create-dev`

## Validated Commands

Validated locally:

- `cargo check`
- `cargo test`
- `cargo run -- provider-info`
- `cargo run -- identity-create`
- `cargo run -- identity-create --device-label "../bad"`
- `cargo run -- identity-create --device-label carbonstack-alice-device`
- duplicate `cargo run -- identity-create --device-label carbonstack-alice-device`
- `cargo run -- public-bundle-export`
- `go test ./internal/protocol`
- `go test ./...`

## Important Implementation Lessons

### Old prep state blocks real identity creation

The real identity-create path correctly refuses overwrite if old prep state already exists.

This means old `.carbonstack-openmls-sidecar-state` must be removed during local tests when intentionally re-running the first-create path.

This is expected and correct.

### stdout JSON should avoid hand-written format braces where possible

The success envelope was safer to implement with `serde_json::json!` and `serde_json::to_string_pretty(...)` than manual `println!(r#"..."#)` formatting, because Rust format strings treat braces specially.

### signer persistence is real now

Earlier scratch work proved signer persistence mattered for valid signatures after reload.

This sidecar rung now writes `signer.json`, but it does not yet prove load/reuse behavior.

A future rung should validate identity loading without printing signer contents.

## Security Boundary

This result is dev-only.

It does not provide:

- production secure storage
- encrypted local vault
- hardware-backed keys
- key recovery
- revocation
- identity verification UX
- public bundle export
- KeyPackage generation/export
- conversation lifecycle
- message protect/open
- Comms CLI integration
- Cypher routing
- trust-store mutation
- production E2EE

## Next Recommended Work

Recommended next step:

- add a load/status check for existing identity state, or plan public-bundle-export only after proving the identity can be loaded safely.

Preferred next rung:

- `identity-status --device-label <label>` or `identity-load-check --device-label <label>`

Purpose:

- verify sidecar can detect/load existing identity state
- verify `signer.json` can be deserialized
- return sanitized summary only
- do not print signer body
- do not generate KeyPackages yet

Alternative next rung:

- plan `public-bundle-export`

But the safer order is:

1. identity load/status check
2. public-bundle-export plan
3. public-bundle-export implementation
4. conversation lifecycle planning

## Allowed Claims

Allowed:

- `identity-create` now generates dev-only OpenMLS identity/signing material.
- `identity-create` writes secret-bearing `signer.json` under ignored local dev state.
- `identity-create` writes sanitized summary/state JSON.
- `identity-create` returns a sanitized success envelope.
- `identity-create` does not print private material.
- `identity-create` refuses overwrite.
- Go tests validate real dev identity state creation and duplicate refusal.
- No KeyPackage, public bundle, provider storage, conversation, message, Comms, Cypher, or trust-store integration exists yet.

## Not Allowed Claims

Not allowed:

- identity creation is production secure.
- local signer storage is a secure vault.
- hardware-backed identity exists.
- KeyPackage generation exists in the sidecar.
- public bundle export exists.
- conversation creation exists.
- message protect/open exists.
- provider storage persistence is solved.
- Comms CLI consumes sidecar identity state.
- Cypher routes MLS payloads.
- trust-state storage consumes sidecar events.
- hostile-server security is validated.
- replay resistance is validated.
- metadata privacy is validated.
- production E2EE exists.
