# CarbonStack OpenMLS Sidecar Identity-Status Result v0

## Status

Classification: PHASE 2D SIDECAR IDENTITY-STATUS RESULT / POST-DEV-IDENTITY-CREATE / PRE-PUBLIC-BUNDLE

This document records the first OpenMLS sidecar identity-status/load-check command.

This is not production secure storage.

This does not generate new OpenMLS signing keys.

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

- `docs/48-provider-identity-prep-state-events-result-v0.md`
- `docs/49-openmls-sidecar-real-identity-create-plan-v0.md`
- `docs/50-openmls-sidecar-real-identity-create-result-v0.md`
- `docs/51-openmls-sidecar-identity-status-plan-v0.md`

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

- `identity-status --device-label <label>`

The command loads and summarizes existing dev-only identity state without creating new identity material.

Successful identity-status behavior:

- validates the device label
- requires existing sidecar identity state
- loads/deserializes secret-bearing `signer.json`
- reads `identity-summary.json`
- reads `identity-state.json`
- derives the signer public key
- recomputes the `sha256:<hex>` public identity reference
- compares the recomputed public identity reference against the stored summary reference
- emits sanitized JSON only
- reports `identity_exists: true`
- reports `identity_loadable: true`
- reports `identity_created: true`
- reports `provider_storage_written: false`
- reports `public_bundle_available: false`
- reports `private_material_included: false`
- emits `provider.identity.loaded`

The command does not print `signer.json`.

## State Files Read

For a device label such as:

- `carbonstack-alice-device`

The command checks under:

- `carbonstack-comms/internal/protocol/mls/research/openmls-sidecar/.carbonstack-openmls-sidecar-state/dev/devices/carbonstack-alice-device/`

Expected files:

- `identity-prep.json`
- `identity-summary.json`
- `identity-state.json`
- `signer.json`

`signer.json` is secret-bearing and must not be printed, pasted, or committed.

## Success Envelope

Successful `identity-status --device-label carbonstack-alice-device` returns:

- `ok: true`
- `command: identity-status`
- `phase: phase2d-identity-status-dev`
- `identity_exists: true`
- `identity_loadable: true`
- `identity_created: true`
- `provider_storage_written: false`
- `public_bundle_available: false`
- `private_material_included: false`
- event: `provider.identity.loaded`

The success envelope includes:

- device label
- state path hint
- identity summary path hint
- identity state path hint
- signer path hint
- public identity reference
- public signature key length

The signer path hint is allowed.

The signer file contents are not allowed in stdout.

## Public Identity Reference Check

`identity-status` recomputes the public identity reference from the loaded signer public key.

Current shape:

- `sha256:<hex digest>`

The recomputed reference must match the reference stored in `identity-summary.json`.

This proves more than file existence:

- the signer can be deserialized
- the public key can be derived
- the sanitized summary corresponds to the signer
- the identity can be loaded in a later sidecar command without printing private material

## Missing Identity Behavior

`identity-status --device-label <safe-label>` before identity creation returns:

- nonzero exit
- expected exit code: `3`
- `ok: false`
- error code: `identity_missing`
- provider event: `provider.identity.missing`
- severity: `warning`
- trust relevant: `false`
- `identity_exists: false`
- `identity_loadable: false`
- `private_material_included: false`

Missing identity is operationally blocking for identity-dependent commands, but it is not automatically evidence of cryptographic compromise.

## Invalid Command Input Behavior

Missing device label:

- command: `identity-status`
- error code: `missing_required_argument`
- provider event: `provider.command.invalid`
- severity: `warning`
- trust relevant: `false`
- private material included: `false`

Invalid device label:

- command: `identity-status`
- error code: `invalid_device_label`
- provider event: `provider.command.invalid`
- severity: `warning`
- trust relevant: `false`
- private material included: `false`

## Unloadable / Corrupt Identity Behavior

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

New events stabilized/used by this rung:

- `provider.identity.loaded`
- `provider.identity.missing`

Existing event reused:

- `provider.secret.material.unavailable`

### `provider.identity.loaded`

Meaning:

- existing sidecar identity state was found
- secret-bearing signer material was loaded/deserialized locally
- public identity reference was recomputed
- summary state matched signer-derived public identity reference
- no private material was printed
- no new identity material was generated
- no public bundle was exported

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

### `provider.identity.missing`

Meaning:

- identity-dependent command expected local identity state
- required identity files were missing

Current classification:

- class: `storage/checkpoint`
- severity: `warning`
- trust relevant: `false`

Current trust-decision behavior:

- stop operation
- show recovery path
- append history
- block current identity-dependent outgoing/send operation
- do not require reverify by default
- not automatically cryptographic trust relevant

## Go-Side Validation

Go-side tests now validate:

- `ProviderEventIdentityLoaded` exists
- `ProviderEventIdentityMissing` exists
- `provider.identity.loaded` is classified as storage/checkpoint, info, non-trust-relevant
- `provider.identity.missing` is classified as storage/checkpoint, warning, non-trust-relevant
- `provider.identity.loaded` maps to append-history/debug-only behavior
- `provider.identity.missing` maps to stop-operation/show-recovery-path/append-history behavior
- identity-status before create returns `identity_missing`
- identity-status after identity-create succeeds
- identity-status reports `identity_exists: true`
- identity-status reports `identity_loadable: true`
- identity-status reports `identity_created: true`
- identity-status returns same public identity ref as identity-create
- identity-status returns same public signature key length as identity-create
- identity-status does not print obvious secret-related tokens
- identity-status does not claim provider storage or public bundle availability

## Rust-Side Validation

Rust-side validation includes:

- sidecar compiles
- sidecar tests pass
- `identity-status` command is recognized
- `identity-status` missing-label path works
- `identity-status` invalid-label path works
- `identity-status` missing-state path works
- `identity-status` after identity-create works
- `signer.json` exists but was not printed
- sanitized `identity-summary.json` and `identity-state.json` remain safe to inspect

## Validated Commands

Validated locally:

- `cargo check`
- `cargo test`
- `cargo run -- identity-status --device-label carbonstack-alice-device`
- `cargo run -- identity-create --device-label carbonstack-alice-device`
- `cargo run -- identity-status`
- `cargo run -- identity-status --device-label "../bad"`
- `cargo run -- public-bundle-export`
- `go test ./internal/protocol`
- `go test ./...`

## Important Implementation Lessons

### Keep writer structs and reader structs separate

The first implementation attempt tried to deserialize borrowed `IdentitySummary<'a>` / `IdentityState<'a>` structs.

That triggered Rust/Serde lifetime problems.

The stable approach uses:

- borrowed structs for writing sanitized JSON
- owned read structs with `String` fields for deserialization

### Status should load, not create

`identity-status` does not create missing state.

This is important because status/check commands should not mutate identity state as a side effect.

### Missing identity is operational, not cryptographic compromise

`provider.identity.missing` blocks the current identity-dependent operation and shows recovery path, but is not automatically trust-relevant.

### Unloadable secret material remains serious

`provider.secret.material.unavailable` remains the correct serious/fatal path for signer deserialization or secret-material availability failures.

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

Recommended next major step:

- plan `public-bundle-export`

The public-bundle-export plan should define:

- whether bundle means KeyPackage, credential summary, or both
- whether KeyPackage generation mutates provider storage
- whether KeyPackages are one-time, reusable, or dev-only reusable in this prototype
- whether output is printed, written to disk, or both
- event vocabulary
- no-secret stdout rules
- duplicate/regeneration semantics
- tests for no private material

Do not proceed directly to conversation lifecycle, message protect/open, Comms runtime integration, Cypher routing, or trust-store mutation.

## Allowed Claims

Allowed:

- `identity-status` is implemented.
- `identity-status` can load existing dev-only sidecar identity state.
- `identity-status` deserializes `signer.json` without printing it.
- `identity-status` recomputes the public identity reference from signer public key.
- `identity-status` verifies the sanitized summary matches the signer-derived reference.
- `provider.identity.loaded` and `provider.identity.missing` are typed/classified/trust-mapped.
- Go tests validate identity-status success and missing-state behavior.
- No KeyPackage, public bundle, provider storage, conversation, message, Comms, Cypher, or trust-store integration exists yet.

## Not Allowed Claims

Not allowed:

- production secure identity storage exists.
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
