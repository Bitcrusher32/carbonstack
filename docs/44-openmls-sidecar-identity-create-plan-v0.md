# CarbonStack OpenMLS Sidecar Identity-Create Plan v0

## Status

Classification: PHASE 2D SIDECAR IDENTITY-CREATE PLAN / PRE-IMPLEMENTATION

This document plans the first secret-adjacent OpenMLS sidecar command:

- `identity-create`

It does not implement identity creation.

It does not implement public bundle export.

It does not implement conversation creation.

It does not implement message protect/open.

It does not wire OpenMLS into CarbonStackComms user-facing commands.

It does not route MLS traffic through CarbonStackCypher.

It does not mutate trust-state storage.

It does not claim production E2EE.

## Source Context

Relevant docs:

- `docs/36-provider-event-taxonomy-v0.md`
- `docs/38-provider-trust-state-mapping-v0.md`
- `docs/39-phase2d-sidecar-command-surface-plan.md`
- `docs/40-openmls-sidecar-provider-info-result-v0.md`
- `docs/41-openmls-sidecar-json-envelope-plan-v0.md`
- `docs/42-openmls-sidecar-json-envelope-result-v0.md`
- `docs/43-provider-command-unsupported-event-result-v0.md`

Relevant Comms paths:

- `carbonstack-comms/internal/protocol/mls/research/openmls-sidecar`
- `carbonstack-comms/internal/protocol/mls/research/openmls-sidecar/src/main.rs`
- `carbonstack-comms/internal/protocol/openmls_sidecar_provider_info_test.go`
- `carbonstack-comms/internal/protocol/provider_events.go`
- `carbonstack-comms/internal/protocol/provider_trust.go`

## Purpose

`identity-create` is the first sidecar command that may need to generate or persist provider-side identity material.

Because it is secret-adjacent, it must be planned before implementation.

The goal is to define:

- command input
- success envelope
- error envelope
- local state behavior
- no-secret stdout rules
- provider event surface
- validation expectations
- explicitly unsupported production claims

## Core Principle

`identity-create` may create secret material internally.

`identity-create` must not print secret material.

stdout must remain safe summary JSON.

## Proposed Command

```powershell
cargo run -- identity-create --device-label carbonstack-alice-device
```

Candidate arguments:

- `--device-label <label>`

Optional future arguments:

- `--profile <dev|test|local>`
- `--state-dir <path>`
- `--overwrite false`
- `--output-public-summary true`

For the first implementation, keep arguments minimal.

Recommended v0 command:

- require `--device-label`
- use a dev-only sidecar state directory
- refuse overwrite unless explicitly designed later

## State Directory v0

The sidecar needs a local state directory for generated provider material.

Candidate default under the sidecar crate:

- `.carbonstack-openmls-sidecar-state/dev`

Candidate per-device layout:

- `.carbonstack-openmls-sidecar-state/dev/devices/<device-label>/identity.json`
- `.carbonstack-openmls-sidecar-state/dev/devices/<device-label>/signer.json`
- `.carbonstack-openmls-sidecar-state/dev/devices/<device-label>/provider-storage.json`

This is dev-only.

It is not production vault design.

The directory must be ignored by Git.

The directory must never be committed.

## stdout Safety Rule

`identity-create` stdout must never include:

- private signing key material
- signer JSON
- MemoryStorage JSON
- provider storage JSON
- recovery secret
- raw private key bytes
- raw seed material
- full local filesystem secret contents

stdout may include:

- `ok`
- `command`
- provider metadata
- device label
- public credential summary
- public key or public-key hash/reference
- public bundle reference
- state checkpoint status
- provider events
- warnings
- `private_material_included: false`

## Success Envelope v0

Proposed success envelope:

```json
{
  "ok": true,
  "command": "identity-create",
  "provider": "openmls",
  "implementation": "carbonstack-openmls-sidecar",
  "mode": "experimental-sidecar",
  "phase": "phase2d-identity-create",
  "data": {
    "device_label": "carbonstack-alice-device",
    "identity_created": true,
    "public_bundle_available": true,
    "public_identity_ref": "dev-only-summary-ref",
    "state_written": true,
    "state_scope": "dev-local-sidecar-state",
    "state_path_hint": ".carbonstack-openmls-sidecar-state/dev/devices/carbonstack-alice-device",
    "overwrite": false
  },
  "events": [
    {
      "event": "provider.identity.created",
      "severity": "notice",
      "trust_relevant": false
    },
    {
      "event": "storage.saved",
      "severity": "notice",
      "trust_relevant": false
    }
  ],
  "warnings": [
    "dev-only state; not production secure storage",
    "private material was written locally but not printed",
    "OpenMLS is not wired into CarbonStackComms"
  ],
  "private_material_included": false
}
```

## Proposed New Provider Events

Candidate new events:

- `provider.identity.created`
- `provider.identity.exists`
- `provider.identity.create.failed`

Candidate use:

- `provider.identity.created`: identity material was created and state was written
- `provider.identity.exists`: identity already exists and command refused to overwrite
- `provider.identity.create.failed`: identity creation failed before safe completion

These should be added to taxonomy only when implementation begins.

## Error Cases v0

### Missing device label

Command:

```powershell
cargo run -- identity-create
```

Candidate error:

- `missing_required_argument`

Candidate provider event:

- `provider.command.invalid`

Severity:

- `warning`

Trust relevant:

- `false`

Exit code:

- `2`

### Identity already exists

Candidate error:

- `identity_already_exists`

Candidate provider event:

- `provider.identity.exists`

Severity:

- `warning`

Trust relevant:

- `false`

Exit code:

- `3`

Required behavior:

- do not overwrite existing secret material
- do not print existing secret material
- show state path hint only

### State write failure

Candidate error:

- `state_write_failed`

Candidate provider event:

- `checkpoint.failed`

Severity:

- `warning`

Trust relevant:

- maybe false initially; escalate only with runtime policy

Exit code:

- `4`

Required behavior:

- do not claim identity creation succeeded
- do not print secret material
- return machine-readable error envelope

## State Mutation Rule

`identity-create` is state-mutating.

It must clearly report:

- whether state was written
- where state was intended to be written
- whether checkpoint/save succeeded
- whether overwrite occurred

First implementation should refuse overwrite.

## Git Hygiene Rule

Add ignore rules before or alongside implementation.

Candidate ignore entry in `carbonstack-comms/.gitignore`:

```text
internal/protocol/mls/research/openmls-sidecar/.carbonstack-openmls-sidecar-state/
```

Continue running:

```powershell
powershell -ExecutionPolicy Bypass -File .\scripts\check-no-rust-artifacts.ps1
```

before commits.

## Test Plan v0

Rust-side tests should cover:

- provider constants still correct
- unsupported commands still return expected behavior
- state path builder sanitizes or rejects invalid labels if implemented
- identity-create output contains `private_material_included: false`

Go-side tests should cover:

- `identity-create` success envelope in a temp/dev state directory
- repeated `identity-create` refuses overwrite
- missing `--device-label` returns error envelope
- no stdout field contains obvious private-material keys
- unsupported-command behavior remains unchanged

## Label Safety

Device labels should be constrained.

Recommended first validation:

- allow letters
- allow numbers
- allow dash
- allow underscore
- reject path separators
- reject empty string
- reject `.` and `..`

Do not allow arbitrary path traversal through `--device-label`.

## Implementation Order

Recommended implementation order:

1. Add sidecar state ignore rule.
2. Add sidecar argument parsing for `identity-create --device-label`.
3. Add label validation.
4. Add JSON envelope for missing-label error.
5. Add dev-only state path creation.
6. Add placeholder identity summary only if needed.
7. Then wire real OpenMLS identity material generation.
8. Add Go-side tests after Rust command behavior is stable.

A placeholder-only identity command is acceptable only if clearly labeled as non-OpenMLS.

Better first real implementation should create actual OpenMLS credential/signature material but print only safe summaries.

## Critical Boundary

Even after `identity-create` exists:

- OpenMLS still will not be wired into `comms send`
- Cypher still will not route MLS payloads
- trust-state storage still will not consume provider events
- production storage still will not be solved
- production E2EE still will not exist

## Allowed Claims

Allowed:

- CarbonStack has an identity-create sidecar command plan.
- The plan defines no-secret stdout rules.
- The plan requires local state and overwrite behavior to be explicit.
- The plan keeps identity-create pre-Comms and pre-Cypher.

## Not Allowed Claims

Not allowed:

- identity-create is implemented.
- public bundle export is implemented.
- conversation creation is implemented.
- message protect/open is implemented.
- production storage exists.
- production E2EE exists.
- trust-state consumes sidecar identity events.
- Comms CLI uses OpenMLS identity state.
