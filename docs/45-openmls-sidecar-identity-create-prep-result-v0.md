# CarbonStack OpenMLS Sidecar Identity-Create Prep Result v0

## Status

Classification: PHASE 2D SIDECAR IDENTITY-CREATE PREP RESULT / PRE-SECRET-GENERATION

This document records the first implementation-prep step for the OpenMLS sidecar `identity-create` command.

It does not implement identity creation.

It does not generate OpenMLS identity material.

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

- `docs/39-phase2d-sidecar-command-surface-plan.md`
- `docs/41-openmls-sidecar-json-envelope-plan-v0.md`
- `docs/42-openmls-sidecar-json-envelope-result-v0.md`
- `docs/43-provider-command-unsupported-event-result-v0.md`
- `docs/44-openmls-sidecar-identity-create-plan-v0.md`

Relevant Comms paths:

- `carbonstack-comms/internal/protocol/mls/research/openmls-sidecar/src/main.rs`
- `carbonstack-comms/internal/protocol/mls/research/openmls-sidecar/src/labels.rs`
- `carbonstack-comms/internal/protocol/openmls_sidecar_provider_info_test.go`
- `carbonstack-comms/.gitignore`

## What Was Added

The sidecar now recognizes:

- `identity-create`

The command accepts and validates:

- `--device-label <label>`

The command remains pre-implementation:

- safe labels return a structured `not_implemented` envelope
- no identity material is generated
- no provider storage is written
- no private material is printed

## Git Hygiene

The sidecar local dev state path is ignored:

- `internal/protocol/mls/research/openmls-sidecar/.carbonstack-openmls-sidecar-state/`

This prepares for future dev-only state work while keeping generated local provider material out of Git.

## Label Validation

Device label validation now exists in the sidecar.

Current accepted label shape:

- ASCII letters
- ASCII numbers
- dash
- underscore
- max length 96

Current rejected labels include:

- empty string
- `.`
- `..`
- labels containing path separators
- labels containing spaces
- labels containing unsupported punctuation
- overlong labels

Purpose:

- prevent path traversal through `--device-label`
- keep future per-device state paths predictable
- make identity-create preflight behavior testable before secret generation exists

## Provider Info Change

`provider-info` now reports `identity-create` as a recognized capability.

`identity-create` is no longer listed as unsupported.

It is still not implemented for secret generation.

Current `provider-info` warning includes:

- `identity-create is recognized for argument validation only and does not generate secrets yet`

## Validated Command Behavior

### Provider info

Command:

- `cargo run -- provider-info`

Expected behavior:

- exits `0`
- returns `ok: true`
- includes `provider-info`
- includes `identity-create`
- lists remaining unsupported commands
- reports `private_material_included: false`

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

### Safe label, not implemented yet

Command:

- `cargo run -- identity-create --device-label carbonstack-alice-device`

Expected behavior:

- exits `3`
- returns `ok: false`
- error code: `not_implemented`
- provider event: `provider.command.not_implemented`
- severity: `warning`
- trust relevant: `false`
- private material included: `false`
- `data.identity_created: false`
- `data.state_written: false`

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
- unsupported-command envelope using `public-bundle-export`
- `identity-create` missing-label envelope
- `identity-create` invalid-label envelope
- `identity-create --device-label carbonstack-alice-device` not-implemented envelope
- no private material in these paths
- no identity creation during prep
- no state write during prep

## Event Vocabulary Candidates

This prep result uses two additional event candidates:

- `provider.command.invalid`
- `provider.command.not_implemented`

Current meaning:

- `provider.command.invalid`: command syntax/argument validation failed
- `provider.command.not_implemented`: command is recognized but intentionally not implemented yet

Current classification intent:

- severity: `warning`
- trust relevant: `false`
- operational/developer-facing
- not a cryptographic failure

These should be stabilized in Go provider-event taxonomy before or during the next implementation rung if they are intended to remain stable.

## Why This Matters

This moves `identity-create` from a generic unsupported command into a recognized sidecar command surface without generating secrets prematurely.

The project now has:

- argument validation
- label safety
- structured error envelopes
- no-secret stdout behavior
- no-state-write behavior
- test coverage before secret-bearing implementation

This keeps the sidecar integration ladder controlled.

## Security Boundary

This prep step is not production identity creation.

It does not:

- create signing keys
- create OpenMLS credentials
- create KeyPackages
- write provider storage
- write signer JSON
- write MemoryStorage JSON
- export public bundles
- mutate trust state
- affect Comms CLI behavior
- affect Cypher routing

## Next Recommended Work

Recommended next step:

- stabilize `provider.command.invalid` and `provider.command.not_implemented` in provider-event taxonomy and trust-decision mapping.

Then:

- plan or implement the minimal real `identity-create` state write path.

The real implementation should still be dev-only and should print only sanitized public summaries.

Do not jump directly to message protect/open.

## Allowed Claims

Allowed:

- `identity-create` is now recognized by the sidecar for argument validation.
- Device label validation exists.
- Go-side tests validate missing-label, invalid-label, and safe-label not-implemented envelopes.
- No identity material is generated yet.
- No provider state is written yet.

## Not Allowed Claims

Not allowed:

- identity creation is implemented.
- OpenMLS signer/credential generation exists in the sidecar.
- public bundle export exists.
- conversation creation exists.
- message protect/open exists.
- provider state persistence is solved.
- production storage exists.
- Comms CLI consumes sidecar identity state.
- Cypher routes MLS payloads.
- production E2EE exists.
