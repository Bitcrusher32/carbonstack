# CarbonStack OpenMLS Sidecar Identity-Status Plan v0

## Status

Classification: PHASE 2D SIDECAR IDENTITY-STATUS PLAN / POST-DEV-IDENTITY-CREATE / PRE-PUBLIC-BUNDLE

This document plans the next safe rung after dev-only OpenMLS sidecar identity creation.

It does not implement identity-status.

It does not generate new OpenMLS signing keys.

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

- `docs/48-provider-identity-prep-state-events-result-v0.md`
- `docs/49-openmls-sidecar-real-identity-create-plan-v0.md`
- `docs/50-openmls-sidecar-real-identity-create-result-v0.md`

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

## Purpose

The next implementation rung should add a public-facing sidecar command:

- `identity-status --device-label <label>`

The command should answer:

- does local sidecar identity state exist?
- can the secret-bearing signer state be loaded/deserialized?
- does the public identity reference still match the signer-derived public key?
- what safe summary fields are available?

It must not:

- print signer material
- generate new identity material
- export public bundles
- generate KeyPackages
- create conversations
- protect/open messages
- touch Comms CLI runtime
- route through Cypher
- mutate trust-state storage

## Why This Rung Comes Before Public Bundle Export

The sidecar now writes real dev-only secret-bearing identity material.

Before using that material to create public bundles or KeyPackages, the project should prove that the identity can be loaded safely in a later command.

This protects the next steps from assuming that files merely existing means the identity is usable.

A load/status check is also useful for future diagnostics, corruption handling, and user/developer support.

## Command Shape

Command:

- `identity-status --device-label <label>`

Required argument:

- `--device-label`

Input validation should reuse the existing device-label validator.

No state mutation should occur during successful status checks.

The command should not create missing state.

The command should not repair broken state.

The command should not overwrite anything.

## Success Behavior

For an existing valid identity:

- exit code: `0`
- `ok: true`
- command: `identity-status`
- phase: `phase2d-identity-status-dev`
- `identity_exists: true`
- `identity_loadable: true`
- `identity_created: true`
- `state_scope: dev-local-sidecar-state`
- `provider_storage_written: false`
- `public_bundle_available: false`
- `private_material_included: false`
- return `public_identity_ref`
- return `public_signature_key_len`
- emit `provider.identity.loaded`

The command should read/verify:

- `signer.json`
- `identity-summary.json`
- `identity-state.json`
- `identity-prep.json`

The command should not print any secret-bearing file contents.

## Preferred Verification Behavior

Best target behavior:

1. Load `signer.json` as `SignatureKeyPair`.
2. Derive public signing key with `to_public_vec()`.
3. Compute `sha256:<hex>` public identity reference from that public key.
4. Read `identity-summary.json`.
5. Compare computed public identity ref against summary public identity ref.
6. Return success only if the signer loads and the summary matches.

This proves that:

- signer state is deserializable
- the sanitized summary corresponds to the actual signer state
- the sidecar can reuse local identity state in a later process

If this exact check becomes difficult due to OpenMLS API details, the implementation may fall back to signer deserialization plus summary parsing, but the preferred target is recompute-and-compare.

## Missing Identity Behavior

For missing identity state:

- exit code: nonzero
- recommended exit code: `3`
- `ok: false`
- error code: `identity_missing`
- provider event: `provider.identity.missing`
- severity: `warning`
- trust relevant: `false`
- private material included: `false`

Missing identity is operationally blocking for commands that require local identity, but it is not automatically a cryptographic compromise.

Recommended trust behavior:

- append history
- debug only or show recovery path
- no reverify
- not automatically user-visible as a trust/security alarm

## Corrupt / Unloadable Identity Behavior

For corrupt, malformed, unreadable, or deserialization-failing secret material:

- exit code: nonzero
- recommended exit code: `4`
- `ok: false`
- error code: `identity_unloadable` or `secret_material_unavailable`
- provider event: `provider.secret.material.unavailable`
- severity: `fatal`
- trust relevant: `true`
- private material included: `false`

Rationale:

- If a signer exists but cannot be loaded, the local identity state is not usable.
- This should stop identity-dependent operations.
- It should not print the malformed signer contents.
- It should not attempt automatic repair.

Existing negative fixture/trust mapping already treats `provider.secret.material.unavailable` as serious/fatal local state behavior.

## Summary Mismatch Behavior

If signer-derived public identity ref does not match `identity-summary.json`:

- exit code: nonzero
- recommended exit code: `4`
- `ok: false`
- error code: `identity_summary_mismatch`
- provider event candidate: `provider.identity.mismatch`
- severity: `security` or `fatal`
- trust relevant: `true`
- private material included: `false`

This may be deferred if the first implementation does not include mismatch testing yet.

If implemented, it should be treated as serious because the safe summary no longer corresponds to the secret identity state.

## Event Vocabulary

New event candidates:

- `provider.identity.loaded`
- `provider.identity.missing`
- `provider.identity.mismatch`

Existing event reused:

- `provider.secret.material.unavailable`

### `provider.identity.loaded`

Meaning:

- identity-status loaded existing local identity state
- signer material was deserialized without printing it
- safe summary data was returned
- no new identity material was generated
- no public bundle was exported

Recommended classification:

- class: `storage/checkpoint`
- severity: `info`
- trust relevant: `false`

Recommended trust decision:

- append history
- debug only
- no send/receive/open blocking
- no reverify
- not user-visible

### `provider.identity.missing`

Meaning:

- identity-status did not find required local identity state

Recommended classification:

- class: `storage/checkpoint`
- severity: `warning`
- trust relevant: `false`

Recommended trust decision:

- append history
- debug only or show recovery path
- block identity-dependent operation
- no reverify
- not a cryptographic compromise by itself

### `provider.identity.mismatch`

Meaning:

- summary state does not match signer-derived identity data

Recommended classification:

- class: `trust/security` or `terminal/fatal`
- severity: `security` or `fatal`
- trust relevant: `true`

Recommended trust decision:

- stop operation
- warn user or developer
- show recovery path
- require investigation
- do not auto-repair

This event can be deferred until mismatch behavior is implemented.

## Success Envelope Shape

Target success shape:

{
  "ok": true,
  "command": "identity-status",
  "provider": "openmls",
  "implementation": "carbonstack-openmls-sidecar",
  "mode": "experimental-sidecar",
  "phase": "phase2d-identity-status-dev",
  "data": {
    "device_label": "carbonstack-alice-device",
    "identity_exists": true,
    "identity_loadable": true,
    "identity_created": true,
    "state_scope": "dev-local-sidecar-state",
    "state_path_hint": ".carbonstack-openmls-sidecar-state/dev/devices/carbonstack-alice-device",
    "identity_summary_path_hint": ".carbonstack-openmls-sidecar-state/dev/devices/carbonstack-alice-device/identity-summary.json",
    "identity_state_path_hint": ".carbonstack-openmls-sidecar-state/dev/devices/carbonstack-alice-device/identity-state.json",
    "signer_path_hint": ".carbonstack-openmls-sidecar-state/dev/devices/carbonstack-alice-device/signer.json",
    "public_identity_ref": "sha256:<hex>",
    "public_signature_key_len": 32,
    "provider_storage_written": false,
    "public_bundle_available": false
  },
  "events": [
    {
      "event": "provider.identity.loaded",
      "severity": "info",
      "trust_relevant": false
    }
  ],
  "warnings": [
    "dev-only identity status; not production secure storage",
    "private material was loaded locally but not printed",
    "OpenMLS is not wired into CarbonStackComms",
    "public bundle export is not implemented"
  ],
  "private_material_included": false
}

## Secret Boundary

The command may load:

- `signer.json`

The command must not print:

- `signer.json` contents
- private signing key material
- signer JSON body
- MemoryStorage JSON
- provider storage JSON
- recovery material
- raw private key bytes
- raw seed material
- secret-bearing file contents

The command may print:

- device label
- public identity ref
- public signature key length
- state path hints
- summary path hints
- identity exists/loadable booleans
- provider storage/public bundle booleans
- warnings
- provider events
- `private_material_included: false`

## Test Plan

Rust-side tests should validate:

- identity-status argument parsing
- missing device label behavior
- invalid device label behavior
- phase constant
- identity state helper paths
- signer load helper if exposed as a testable function

Go-side tests should validate:

- successful identity-create followed by identity-status
- identity-status returns `ok: true`
- identity-status reports `identity_exists: true`
- identity-status reports `identity_loadable: true`
- identity-status reports `private_material_included: false`
- identity-status returns `provider.identity.loaded`
- identity-status public identity ref matches identity-create public identity ref
- identity-status does not print obvious secret-related tokens
- missing identity-status returns `identity_missing` / `provider.identity.missing`
- invalid label returns `provider.command.invalid`
- corrupt/unloadable signer behavior if practical to simulate safely

## Implementation Order

Recommended order:

1. Add `identity-status` to sidecar command recognition.
2. Add state helper structs/functions for loading identity summaries and signer state.
3. Implement signer deserialization without printing signer body.
4. Recompute public identity ref from signer public key if available.
5. Compare recomputed ref with summary ref.
6. Emit sanitized success envelope.
7. Add missing identity error path.
8. Add unloadable signer error path.
9. Add Go provider-event taxonomy/trust mapping for new events.
10. Add Go sidecar tests.
11. Add result doc.

## What Not To Do In This Rung

Do not:

- generate a new signer
- overwrite identity state
- export a public bundle
- generate a KeyPackage
- create or join a conversation
- protect/open messages
- write provider storage
- wire into Comms runtime
- route through Cypher
- mutate trust-state storage
- print secret material
- implement production storage

## Security Boundary

This is still dev-only.

Even after identity-status is implemented, CarbonStack still will not have:

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

## Next Work After This Plan

After this plan lands, implement identity-status.

After identity-status result is validated, the next likely major step is:

- `public-bundle-export` plan

Do not proceed to public-bundle export until identity-status proves that existing local identity state can be loaded and summarized safely.

## Allowed Claims After Plan

Allowed:

- CarbonStack has a plan for identity-status.
- The plan requires loading existing dev identity state without printing secrets.
- The plan places identity-status before public-bundle export.
- The plan defines event candidates for loaded, missing, and mismatched identity state.
- The plan preserves no Comms/Cypher/trust runtime integration.

## Not Allowed Claims After Plan

Not allowed:

- identity-status is implemented.
- identity load/check behavior is validated.
- KeyPackage generation exists.
- public bundle export exists.
- conversation creation exists.
- message protect/open exists.
- production storage exists.
- Comms CLI consumes sidecar identity state.
- Cypher routes MLS payloads.
- trust-state storage consumes sidecar events.
- production E2EE exists.
