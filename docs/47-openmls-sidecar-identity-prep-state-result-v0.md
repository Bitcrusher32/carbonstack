# CarbonStack OpenMLS Sidecar Identity Prep State Result v0

## Status

Classification: PHASE 2D SIDECAR IDENTITY PREP STATE RESULT / PRE-SECRET-GENERATION

This document records the first dev-only state-writing step for the OpenMLS sidecar `identity-create` command.

It does not implement real identity creation.

It does not generate OpenMLS signing keys.

It does not generate OpenMLS credentials.

It does not generate KeyPackages.

It does not write provider storage.

It does not export public bundles.

It does not create conversations.

It does not protect/open messages.

It does not wire OpenMLS into CarbonStackComms user-facing commands.

It does not route MLS traffic through CarbonStackCypher.

It does not mutate trust-state storage.

It does not claim production E2EE.

## Source Context

Relevant docs:

- `docs/41-openmls-sidecar-json-envelope-plan-v0.md`
- `docs/42-openmls-sidecar-json-envelope-result-v0.md`
- `docs/43-provider-command-unsupported-event-result-v0.md`
- `docs/44-openmls-sidecar-identity-create-plan-v0.md`
- `docs/45-openmls-sidecar-identity-create-prep-result-v0.md`
- `docs/46-provider-command-invalid-not-implemented-events-result-v0.md`

Relevant Comms paths:

- `carbonstack-comms/internal/protocol/mls/research/openmls-sidecar/src/main.rs`
- `carbonstack-comms/internal/protocol/mls/research/openmls-sidecar/src/labels.rs`
- `carbonstack-comms/internal/protocol/mls/research/openmls-sidecar/src/state.rs`
- `carbonstack-comms/internal/protocol/openmls_sidecar_provider_info_test.go`
- `carbonstack-comms/.gitignore`

## What Was Added

The sidecar `identity-create --device-label <label>` path now creates dev-only prep state after label validation.

Successful safe-label behavior now:

- validates the device label
- creates an ignored per-device state directory
- writes a non-secret `identity-prep.json` manifest
- returns a JSON success envelope
- reports `identity_created: false`
- reports `state_written: true`
- reports `provider_storage_written: false`
- reports `private_material_included: false`

This is a state skeleton, not identity creation.

## State Path

The dev-only prep manifest is written under:

- `carbonstack-comms/internal/protocol/mls/research/openmls-sidecar/.carbonstack-openmls-sidecar-state/dev/devices/<device-label>/identity-prep.json`

The state root is ignored by Git:

- `internal/protocol/mls/research/openmls-sidecar/.carbonstack-openmls-sidecar-state/`

## Manifest Shape

The current non-secret prep manifest contains:

- `manifest_version`
- `device_label`
- `state_scope`
- `identity_created`
- `provider_storage_written`
- `private_material_included`
- `warning`

Required current values:

- `manifest_version: identity-prep/v0`
- `identity_created: false`
- `provider_storage_written: false`
- `private_material_included: false`

The manifest must not contain:

- private signing keys
- signer JSON
- MemoryStorage JSON
- provider storage JSON
- recovery secrets
- raw private key bytes
- raw seed material
- OpenMLS private material

## Validated Command Behavior

### Provider info

Command:

- `cargo run -- provider-info`

Expected behavior:

- exits `0`
- returns `ok: true`
- includes `provider-info`
- includes `identity-create`
- reports `private_material_included: false`
- warns that identity-create writes dev-only non-secret prep state but does not generate OpenMLS secrets yet

### Missing device label

Command:

- `cargo run -- identity-create`

Expected behavior:

- exits `2`
- returns `ok: false`
- error code: `missing_required_argument`
- provider event: `provider.command.invalid`
- severity: `warning`
- trust relevant: `false`
- private material included: `false`

### Invalid device label

Command:

- `cargo run -- identity-create --device-label "../bad"`

Expected behavior:

- exits `2`
- returns `ok: false`
- error code: `invalid_device_label`
- provider event: `provider.command.invalid`
- severity: `warning`
- trust relevant: `false`
- private material included: `false`
- echoes only the submitted label in `data.device_label`

### Safe label first run

Command:

- `cargo run -- identity-create --device-label carbonstack-alice-device`

Expected behavior:

- exits `0`
- returns `ok: true`
- event: `provider.identity.prep_state_written`
- severity: `notice`
- trust relevant: `false`
- `data.device_label: carbonstack-alice-device`
- `data.identity_created: false`
- `data.state_written: true`
- `data.provider_storage_written: false`
- `private_material_included: false`
- writes `identity-prep.json`

### Safe label duplicate run

Command:

- `cargo run -- identity-create --device-label carbonstack-alice-device`

Expected behavior when prep state already exists:

- exits `3`
- returns `ok: false`
- error code: `identity_prep_state_already_exists`
- provider event: `provider.identity.exists`
- severity: `warning`
- trust relevant: `false`
- `data.identity_created: false`
- `data.state_written: false`
- private material included: `false`
- does not overwrite existing prep state

### Other unsupported command

Command:

- `cargo run -- public-bundle-export`

Expected behavior:

- exits `2`
- returns `ok: false`
- error code: `unsupported_command`
- provider event: `provider.command.unsupported`
- severity: `warning`
- trust relevant: `false`
- private material included: `false`

## Go-Side Validation

Go-side tests now cover:

- `provider-info` success envelope
- unsupported command envelope
- identity-create missing-label envelope
- identity-create invalid-label envelope
- identity-create safe-label success envelope
- duplicate identity-create overwrite refusal envelope
- prep manifest creation
- prep manifest JSON validity
- prep manifest non-secret claims
- no identity creation during prep state skeleton
- no provider storage write during prep state skeleton
- `private_material_included: false`

## Event Vocabulary Candidates

This result introduces or uses additional provider event candidates:

- `provider.identity.prep_state_written`
- `provider.identity.exists`
- `checkpoint.failed`

Current meaning:

- `provider.identity.prep_state_written`: non-secret dev-only prep manifest was written
- `provider.identity.exists`: sidecar refused to overwrite existing identity/prep state
- `checkpoint.failed`: sidecar failed to write intended state/checkpoint

Current classification intent:

- `provider.identity.prep_state_written`: notice, non-trust-relevant
- `provider.identity.exists`: warning, non-trust-relevant
- `checkpoint.failed`: warning, operationally important, not automatically cryptographic trust failure at this rung

These events should be stabilized in Go provider-event taxonomy/trust mapping before real OpenMLS identity material generation if they are intended to remain stable.

## Why This Matters

This is the first Phase 2D sidecar state-write rung.

It proves that the sidecar can:

- validate labels
- create an ignored per-device state directory
- write a non-secret manifest
- return a structured success envelope
- refuse overwrite
- preserve no-secret stdout behavior

It still avoids the risky jump into cryptographic identity generation.

## Security Boundary

This result is not production identity creation.

It does not:

- create OpenMLS signing keys
- create OpenMLS credentials
- create KeyPackages
- write signer JSON
- write MemoryStorage JSON
- write provider storage
- export public bundles
- create conversations
- protect/open messages
- mutate `trust.json`
- mutate `trust-events.jsonl`
- affect Comms CLI behavior
- affect Cypher routing

## Current Known Cleanup Notes

The state skeleton replaced the older safe-label `not_implemented` path.

Any stale dead-code helper for the old not-implemented response should be removed from the sidecar if still present.

Rust tests should expect:

- `phase2d-identity-create-state-skeleton`

not:

- `phase2d-identity-create-prep`

## Next Recommended Work

Recommended next step:

- stabilize `provider.identity.prep_state_written`, `provider.identity.exists`, and any checkpoint/state-write event behavior in Go taxonomy/trust mapping.

Then:

- implement minimal dev-only real identity material generation only after event semantics are stable.

The real identity-create implementation should:

- write dev-only state intentionally
- refuse overwrite by default
- print only sanitized public summary
- keep `private_material_included: false`
- avoid Comms CLI integration
- avoid Cypher routing
- avoid trust-store mutation

Do not jump directly to public bundle export, conversation lifecycle, or message protect/open.

## Allowed Claims

Allowed:

- `identity-create` can now write dev-only non-secret prep state.
- The prep manifest is ignored by Git.
- The prep manifest does not claim identity creation.
- The prep manifest does not contain provider storage.
- Duplicate prep creation is refused.
- Go-side tests validate prep state write and overwrite refusal.

## Not Allowed Claims

Not allowed:

- identity creation is implemented.
- OpenMLS signer/credential generation exists in the sidecar.
- KeyPackage generation exists in the sidecar.
- public bundle export exists.
- conversation creation exists.
- message protect/open exists.
- provider state persistence is solved.
- production storage exists.
- Comms CLI consumes sidecar identity state.
- Cypher routes MLS payloads.
- production E2EE exists.
