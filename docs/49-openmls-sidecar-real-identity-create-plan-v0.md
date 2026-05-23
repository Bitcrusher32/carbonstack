# CarbonStack OpenMLS Sidecar Real Identity-Create Plan v0

## Status

Classification: PHASE 2D SIDECAR REAL IDENTITY-CREATE PLAN / PRE-SECRET-GENERATION

This document plans the first real secret-generating OpenMLS sidecar identity command.

It does not implement real identity creation.

It does not generate OpenMLS signing keys.

It does not generate OpenMLS credentials.

It does not generate KeyPackages.

It does not export public bundles.

It does not create conversations.

It does not protect/open messages.

It does not wire OpenMLS into CarbonStackComms user-facing commands.

It does not route MLS traffic through CarbonStackCypher.

It does not mutate trust-state storage.

It does not claim production E2EE.

## Source Context

Relevant docs:

- `docs/42-openmls-sidecar-json-envelope-result-v0.md`
- `docs/43-provider-command-unsupported-event-result-v0.md`
- `docs/44-openmls-sidecar-identity-create-plan-v0.md`
- `docs/45-openmls-sidecar-identity-create-prep-result-v0.md`
- `docs/46-provider-command-invalid-not-implemented-events-result-v0.md`
- `docs/47-openmls-sidecar-identity-prep-state-result-v0.md`
- `docs/48-provider-identity-prep-state-events-result-v0.md`

Relevant Comms paths:

- `carbonstack-comms/internal/protocol/mls/research/openmls-sidecar`
- `carbonstack-comms/internal/protocol/mls/research/openmls-sidecar/src/main.rs`
- `carbonstack-comms/internal/protocol/mls/research/openmls-sidecar/src/labels.rs`
- `carbonstack-comms/internal/protocol/mls/research/openmls-sidecar/src/state.rs`
- `carbonstack-comms/internal/protocol/openmls_sidecar_provider_info_test.go`
- `carbonstack-comms/internal/protocol/provider_events.go`
- `carbonstack-comms/internal/protocol/provider_trust.go`

## Purpose

The next implementation rung should turn `identity-create --device-label <safe-label>` from non-secret prep-state creation into minimal dev-only OpenMLS identity material generation.

This is the first sidecar step that may generate secret material.

The goal is to define the boundary before implementation.

## Current Starting Point

Current `identity-create --device-label <safe-label>` behavior:

- validates the device label
- creates ignored dev-only per-device state directory
- writes non-secret `identity-prep.json`
- returns `ok: true`
- reports `identity_created: false`
- reports `state_written: true`
- reports `provider_storage_written: false`
- reports `private_material_included: false`
- refuses duplicate prep state

Current event vocabulary:

- `provider.identity.prep_state_written`
- `provider.identity.exists`
- `checkpoint.failed`

Current missing work:

- no OpenMLS signing key generation
- no OpenMLS credential generation
- no signer persistence
- no provider storage persistence
- no public bundle export
- no KeyPackage export
- no Comms/Cypher/trust integration

## Definition of Identity Created

For this phase, `identity_created: true` must mean:

- stable local device identity material exists on disk
- the identity is associated with one validated device label
- future sidecar commands can load the identity in a later process
- the sidecar can use the identity for future public bundle / KeyPackage work

It must not mean:

- public bundle exists
- KeyPackage exists
- conversation state exists
- messages can be protected/opened
- Comms CLI uses OpenMLS
- Cypher routes MLS payloads
- trust-state storage consumes provider events
- production secure storage exists

## Command Shape

Initial command:

- `identity-create --device-label <label>`

Still required:

- `--device-label`

Still rejected:

- empty labels
- `.`
- `..`
- labels with path separators
- labels with spaces
- labels with unsupported punctuation
- overlong labels

No overwrite flag should be added yet.

Overwrite remains refused by default.

## State Layout v0

Current non-secret prep state path:

- `.carbonstack-openmls-sidecar-state/dev/devices/<device-label>/identity-prep.json`

Candidate real identity state files:

- `.carbonstack-openmls-sidecar-state/dev/devices/<device-label>/identity-summary.json`
- `.carbonstack-openmls-sidecar-state/dev/devices/<device-label>/signer.json`
- `.carbonstack-openmls-sidecar-state/dev/devices/<device-label>/identity-state.json`

Interpretation:

- `identity-summary.json`: non-secret summary suitable for tests and debugging.
- `signer.json`: secret-bearing signer/private identity material; ignored by Git; never printed.
- `identity-state.json`: metadata about the dev identity state; should avoid secret material unless OpenMLS requires otherwise.

If OpenMLS APIs require a different layout, the implementation should document that in the result doc.

## stdout Safety Rule

The sidecar stdout must never print:

- private signing key material
- signer JSON body
- MemoryStorage JSON
- provider storage JSON
- recovery material
- raw private key bytes
- raw seed material
- secret-bearing file contents

The sidecar stdout may print:

- device label
- identity created true/false
- state written true/false
- provider storage written true/false
- public identity reference
- public key hash/fingerprint/reference
- state path hint
- warnings
- provider events
- `private_material_included: false`

## Success Envelope v0

Target success shape:

{
  "ok": true,
  "command": "identity-create",
  "provider": "openmls",
  "implementation": "carbonstack-openmls-sidecar",
  "mode": "experimental-sidecar",
  "phase": "phase2d-identity-create-dev",
  "data": {
    "device_label": "carbonstack-alice-device",
    "identity_created": true,
    "state_written": true,
    "state_scope": "dev-local-sidecar-state",
    "state_path_hint": ".carbonstack-openmls-sidecar-state/dev/devices/carbonstack-alice-device",
    "identity_summary_path_hint": ".carbonstack-openmls-sidecar-state/dev/devices/carbonstack-alice-device/identity-summary.json",
    "public_identity_ref": "dev-only-public-ref-or-fingerprint",
    "provider_storage_written": false,
    "public_bundle_available": false
  },
  "events": [
    {
      "event": "provider.identity.created",
      "severity": "notice",
      "trust_relevant": false
    }
  ],
  "warnings": [
    "dev-only identity material; not production secure storage",
    "private material was written locally but not printed",
    "OpenMLS is not wired into CarbonStackComms"
  ],
  "private_material_included": false
}

## Event Vocabulary

New event candidate:

- `provider.identity.created`

Meaning:

- a validated device label was accepted
- local dev identity material was generated
- required local state was written
- stdout remained sanitized
- no public bundle was exported
- no conversation state was created

Initial classification intent:

- class: storage/checkpoint or lifecycle
- severity: notice/info
- trust relevant: false

Recommended mapping:

- append history
- debug only
- no send/receive/open blocking
- no reverify
- not user visible

Existing events retained:

- `provider.identity.exists`
- `checkpoint.failed`
- `provider.command.invalid`
- `provider.command.unsupported`

## Duplicate / Refused Overwrite Behavior

If identity state already exists for a device label:

- refuse overwrite
- return `ok: false`
- use error code `identity_already_exists` or preserve current `identity_prep_state_already_exists` until renamed deliberately
- emit `provider.identity.exists`
- do not print private material
- do not replace signer/private state
- do not mutate existing identity state

Recommendation:

- if real identity files exist, use `identity_already_exists`
- if only prep manifest exists, either upgrade carefully or refuse and require cleanup
- do not silently convert prep state into real identity state unless tests explicitly validate the upgrade path

## KeyPackage / Public Bundle Boundary

Do not generate/export KeyPackages in `identity-create` unless implementation inspection proves that it is unavoidable.

Preferred command split:

- `identity-create`: creates stable local device identity/signing material.
- `public-bundle-export`: creates/exports public setup material later.

Reason:

- public bundles and KeyPackages may have one-time or lifecycle-sensitive semantics
- identity creation should stay separate from shareable/public setup generation
- later Cypher onboarding will need clear distinction between local identity and exported setup material

## Provider Storage Boundary

Do not claim provider storage is written unless OpenMLS provider storage is actually written intentionally.

For first real identity-create, `provider_storage_written` may remain false if the sidecar writes signer/identity material only.

If OpenMLS requires MemoryStorage/provider storage for identity creation, then:

- write only under ignored dev state
- never print provider storage JSON
- set `provider_storage_written: true`
- document exact files in result doc
- add tests proving no provider storage content appears in stdout

## Test Plan

Rust-side tests should validate:

- label validation still works
- identity-create still refuses invalid labels
- identity state path construction is stable
- duplicate state refuses overwrite
- generated summary does not include private material
- stdout helper reports `private_material_included: false`

Go-side tests should validate:

- successful `identity-create --device-label <safe>` returns `ok: true`
- `identity_created: true`
- `state_written: true`
- `private_material_included: false`
- duplicate identity-create refuses overwrite
- private material is not present in stdout
- expected state files exist
- secret-bearing files are not printed
- public bundle is not claimed available unless implemented
- provider storage is not claimed unless actually written

## Pre-Implementation Inspection Required

Before implementation, inspect current Rust/OpenMLS APIs and dependencies for:

- available signer/keypair creation APIs
- signer serialization APIs
- credential creation APIs
- whether signer can be persisted as JSON safely for dev
- whether provider storage is needed for identity-only state
- whether earlier scratch code has reusable signer persistence helpers
- whether KeyPackage creation should remain deferred

Do not implement by guessing.

## Implementation Order

Recommended implementation order:

1. Inspect sidecar dependencies and scratch signer persistence helpers.
2. Decide exact minimal identity material type.
3. Add `provider.identity.created` event to Go taxonomy/trust mapping if needed before code or alongside implementation.
4. Implement state file path helpers for real identity state.
5. Implement real identity generation.
6. Write secret-bearing state only under ignored dev state.
7. Write non-secret identity summary.
8. Return sanitized success envelope.
9. Add duplicate/refused overwrite tests.
10. Add no-secret stdout tests.
11. Add result doc.

## Security Boundary

This next implementation will still be dev-only.

It will not provide:

- production secure storage
- encrypted local vault
- hardware-backed keys
- recovery flow
- revocation
- public bundle export
- conversation lifecycle
- message protect/open
- Comms CLI integration
- Cypher routing
- trust-store integration
- production E2EE

## Allowed Claims After Plan

Allowed:

- CarbonStack has a plan for minimal dev-only real sidecar identity creation.
- The plan separates identity creation from public bundle export.
- The plan requires sanitized stdout and ignored local dev state.
- The plan requires overwrite refusal and no-secret output tests.

## Not Allowed Claims After Plan

Not allowed:

- identity creation is implemented.
- OpenMLS signer/credential generation exists in sidecar.
- KeyPackage generation exists in sidecar.
- public bundle export exists.
- conversation creation exists.
- message protect/open exists.
- production storage exists.
- Comms CLI consumes sidecar identity state.
- Cypher routes MLS payloads.
- trust-state storage consumes provider events.
- production E2EE exists.
