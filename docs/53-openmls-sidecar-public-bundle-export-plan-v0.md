# CarbonStack OpenMLS Sidecar Public-Bundle-Export Plan v0

## Status

Classification: PHASE 2D SIDECAR PUBLIC-BUNDLE-EXPORT PLAN / POST-IDENTITY-STATUS / PRE-CONVERSATION-LIFECYCLE

This document plans the next safe rung after dev-only identity creation and identity-status/load-check.

It does not implement public-bundle-export.

It does not generate KeyPackages yet.

It does not create conversations.

It does not protect/open messages.

It does not wire OpenMLS into CarbonStackComms user-facing commands.

It does not route MLS traffic through CarbonStackCypher.

It does not mutate trust-state storage.

It does not claim production E2EE.

## Source Context

Relevant docs:

- `docs/49-openmls-sidecar-real-identity-create-plan-v0.md`
- `docs/50-openmls-sidecar-real-identity-create-result-v0.md`
- `docs/51-openmls-sidecar-identity-status-plan-v0.md`
- `docs/52-openmls-sidecar-identity-status-result-v0.md`

Relevant Comms paths:

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

- `public-bundle-export --device-label <label>`

The command should export/share sanitized public setup material derived from an existing dev-only sidecar identity.

It must not:

- print signer/private material
- create a new identity
- overwrite identity state
- create a conversation
- protect/open messages
- wire into Comms runtime
- route through Cypher
- mutate trust-state storage

## Why This Comes After Identity-Status

`identity-create` proves that dev-only OpenMLS identity material can be generated.

`identity-status` proves that existing identity state can be loaded, deserialized, and checked against its sanitized public identity reference without printing `signer.json`.

`public-bundle-export` should come after those because public bundle generation depends on reusable local identity state.

The sidecar should not export shareable setup material until identity loading and summary matching are validated.

## Working Definition: Public Bundle

For this rung, a CarbonStack public bundle means:

- sanitized public setup material for one device identity
- safe to print as JSON
- safe to copy into future onboarding/invite flows
- not secret-bearing
- not a conversation
- not a message
- not proof of production E2EE

Candidate public bundle contents:

- device label
- provider name
- implementation name
- ciphersuite label
- public identity reference
- public signature key length
- credential type
- KeyPackage hash/reference if KeyPackage generation is implemented
- public bundle version
- warnings
- `private_material_included: false`

The bundle must not include:

- private signing key material
- signer JSON body
- provider storage JSON
- MemoryStorage JSON
- recovery material
- raw private key bytes
- raw seed material
- secret-bearing file contents

## KeyPackage Semantics

This is the sensitive design point.

OpenMLS public setup likely needs a KeyPackage-like object for another member to add this device to a group.

However, KeyPackages should not be casually treated as an infinitely reusable public profile blob.

For this first sidecar rung, the recommended behavior is:

- generate a dev-only KeyPackage/public setup artifact from existing identity state
- return only a sanitized summary/reference in stdout
- write any full serialized public setup artifact under ignored dev sidecar state only if needed
- explicitly label the result as dev-only
- avoid claiming reusable/production semantics
- avoid conversation creation

If OpenMLS requires provider storage mutation for KeyPackage generation, the implementation must report that explicitly.

If provider storage is not intentionally written, the result must keep:

- `provider_storage_written: false`

## Recommended Command Shape

Command:

- `public-bundle-export --device-label <label>`

Required:

- `--device-label`

The command should:

1. validate device label
2. require existing identity state
3. load signer/identity state using the same safe identity-status path
4. generate or derive public setup material
5. write any dev-only public bundle artifact if needed
6. return sanitized JSON only
7. emit a typed provider event
8. keep `private_material_included: false`

The command should not create missing identity state.

The command should not overwrite identity state.

The command should not create conversations.

## State Layout Candidate

Existing identity state lives under:

- `.carbonstack-openmls-sidecar-state/dev/devices/<device-label>/`

Existing files:

- `identity-prep.json`
- `identity-summary.json`
- `identity-state.json`
- `signer.json`

Candidate public bundle files:

- `public-bundle-summary.json`
- `public-bundle-state.json`
- `key-package.bin` or `key-package.json`, only if OpenMLS serialization path is clear and the artifact is public-only

Any full KeyPackage artifact must be treated as public setup material, not secret material, but still dev-state until semantics are mature.

## Success Behavior

Successful `public-bundle-export --device-label carbonstack-alice-device` should return:

- exit code: `0`
- `ok: true`
- command: `public-bundle-export`
- phase: `phase2d-public-bundle-export-dev`
- `device_label: carbonstack-alice-device`
- `identity_exists: true`
- `identity_loadable: true`
- `public_bundle_available: true`
- `public_bundle_exported: true`
- `key_package_created: true` if KeyPackage generation is implemented
- `provider_storage_written: true/false`, depending on actual OpenMLS behavior
- `private_material_included: false`
- event: `provider.public_bundle.exported` or `provider.public_bundle.created`

## Event Vocabulary

Recommended new event:

- `provider.public_bundle.exported`

Meaning:

- existing identity state was loaded
- public setup material was generated/exported
- stdout remained sanitized
- no private material was printed
- no conversation was created
- no message was protected/opened

Recommended classification:

- class: `public_setup`
- severity: `info`
- trust relevant: `false`

Recommended trust decision:

- append history
- debug only
- no send/receive/open blocking
- no reverify
- not user visible by default

Existing events reused:

- `provider.command.invalid`
- `provider.identity.missing`
- `provider.secret.material.unavailable`
- `checkpoint.failed`

Possible future event:

- `provider.public_bundle.exists`

Use only if duplicate export/refusal semantics are implemented.

## Missing Identity Behavior

If identity state is missing:

- exit code: `3`
- `ok: false`
- error code: `identity_missing`
- provider event: `provider.identity.missing`
- private material included: `false`

The command must not auto-create identity state.

## Unloadable Identity Behavior

If signer/identity state cannot be loaded:

- exit code: `4`
- `ok: false`
- error code: `secret_material_unavailable`
- provider event: `provider.secret.material.unavailable`
- severity: `fatal`
- trust relevant: `true`
- private material included: `false`

The command must not print corrupt or malformed secret-bearing file contents.

## Duplicate / Regeneration Semantics

For the first implementation, choose one of two explicit behaviors:

### Option A: overwrite public bundle summary only

Allowed only if public bundle output is public-only and deterministic enough for this dev rung.

Pros:

- simpler iteration
- public bundle command can be rerun

Cons:

- may hide KeyPackage lifecycle semantics

### Option B: refuse duplicate public bundle by default

Safer if KeyPackage semantics are unclear.

Pros:

- avoids accidental KeyPackage reuse/regeneration confusion
- matches identity-create overwrite discipline

Cons:

- requires cleanup during tests

Recommendation:

- Use Option B if a serialized KeyPackage/public setup artifact is written.
- Use Option A only if the command returns a summary-only non-state-mutating public reference.

## stdout Safety Rule

stdout may include:

- device label
- public identity ref
- public signature key length
- public bundle ref/hash
- KeyPackage hash/reference
- ciphersuite label
- credential type
- state path hints
- public artifact path hints
- provider events
- warnings
- `private_material_included: false`

stdout must not include:

- signer JSON
- private signing key material
- MemoryStorage JSON
- provider storage JSON
- recovery material
- raw secret bytes
- raw seed material

## Test Plan

Rust-side tests should validate:

- command recognition
- missing label behavior
- invalid label behavior
- phase constant
- public bundle path helpers if added
- no state mutation in invalid paths

Go-side tests should validate:

- public-bundle-export before identity returns identity_missing
- identity-create followed by public-bundle-export succeeds
- public-bundle-export reports `private_material_included: false`
- public-bundle-export reports no conversation/message creation
- public-bundle-export emits typed public-bundle event
- public identity ref matches identity-create/status identity ref
- public bundle summary file exists if written
- no obvious secret tokens appear in stdout
- `signer.json` exists but is not printed
- provider storage claim matches actual implementation behavior

## Implementation Order

Recommended order:

1. Inspect OpenMLS KeyPackage creation APIs in the current sidecar dependency set.
2. Reuse identity-status loading path to load existing signer and summary.
3. Decide whether the first public bundle includes an actual KeyPackage or only a public identity/credential summary.
4. Add public bundle state/path helpers.
5. Implement sanitized success/error envelopes.
6. Add `provider.public_bundle.exported` to Go event taxonomy/trust mapping.
7. Add Go sidecar tests.
8. Add result doc.

## What Not To Do In This Rung

Do not:

- create conversations
- add members
- join groups
- protect/open messages
- route through Cypher
- wire into Comms CLI
- mutate trust-state storage
- claim production security
- print or commit signer material
- treat dev KeyPackages as production onboarding UX

## Security Boundary

Even after this plan lands, CarbonStack still will not have:

- production secure storage
- encrypted local vault
- hardware-backed keys
- recovery/revocation
- verified identity UX
- conversation lifecycle
- message protect/open
- Comms CLI integration
- Cypher MLS routing
- trust-store integration
- production E2EE

## Next Work After This Plan

After this plan lands:

- inspect current OpenMLS KeyPackage/public setup APIs
- implement the narrowest safe `public-bundle-export`
- record result doc
- checkpoint

After public-bundle-export:

- harden missing/unloadable/mismatch/bundle duplicate behavior
- then plan conversation-create/add-member/join

## Allowed Claims After Plan

Allowed:

- CarbonStack has a public-bundle-export plan.
- The plan places public-bundle-export after identity-create and identity-status.
- The plan treats KeyPackage/public setup semantics as sensitive.
- The plan requires sanitized stdout and no secret printing.
- The plan keeps Comms/Cypher/trust runtime integration out of scope.

## Not Allowed Claims After Plan

Not allowed:

- public-bundle-export is implemented.
- KeyPackage generation exists.
- public bundle semantics are finalized.
- conversation creation exists.
- message protect/open exists.
- Comms CLI consumes OpenMLS.
- Cypher routes MLS payloads.
- trust-state storage consumes sidecar events.
- production secure storage exists.
- production E2EE exists.
