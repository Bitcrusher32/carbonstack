# CarbonStack OpenMLS MemoryStorage Persistence Result v0

## Status

Classification: PHASE 2C STORAGE/PERSISTENCE RESULT / SCRATCH-ONLY

This document records the first successful OpenMLS process-restart-shaped persistence probe.

It does not claim production storage.

It does not claim secure key storage.

It does not wire OpenMLS into CarbonStackComms.

It does not claim production E2EE.

## Source Context

Relevant prior docs:

- `docs/31-openmls-persistence-spike-plan.md`
- `docs/32-openmls-storage-decision-v0.md`

Relevant scratch path:

- `carbonstack-comms/internal/protocol/mls/research/openmls-minimal`

Relevant commit:

- `carbonstack-comms` `58d7211 test: probe OpenMLS MemoryStorage file persistence`

## What Changed

The OpenMLS scratch crate now has a two-phase run mode:

- `cargo run -- phase-a`
- `cargo run -- phase-b`

Phase A creates local Alice/Bob MLS state, sends and opens one application message, and saves provider/signing state.

Phase B starts with fresh providers, loads saved state, reloads Alice/Bob groups, sends a second application message, and validates Bob can open it.

## Why This Matters

Previous v0.2.6/v0.2.7 probes proved:

- application message protect/open works in one process
- two sequential messages work in one process
- `MlsGroup::load(provider.storage(), group_id)` works inside the same process
- OpenMLS provider operations are stateful

This result goes further.

It proves that scratch-level OpenMLS state can survive a file save/load boundary and continue messaging after fresh provider construction.

## Implementation Shape

The scratch crate now uses a CarbonStack-only scratch provider wrapper.

Reason:

- stock `OpenMlsRustCrypto` uses `MemoryStorage`
- stock `OpenMlsRustCrypto` keeps its `key_store` field private
- `MemoryStorage` exposes feature-gated persistence helpers
- a scratch wrapper can own `RustCrypto + MemoryStorage` directly

The wrapper exists only for research.

It is not a production provider design.

## Validated Phase A

Phase A validated:

- Alice provider and Bob provider are separate.
- Alice setup material is created.
- Bob setup material is created.
- Alice creates an MLS group.
- Alice adds Bob from Bob's KeyPackage.
- Welcome is extracted from `MlsMessageOut`.
- Alice merges pending commit.
- Bob stages Welcome.
- Bob joins into his own MlsGroup.
- Alice sends phase-A message one.
- Bob processes/opens phase-A message one.
- Bob plaintext matches Alice plaintext.
- Alice signer is saved to OS temp JSON.
- Alice MemoryStorage is saved using OpenMLS memory-storage persistence helpers.
- Bob MemoryStorage is saved using OpenMLS memory-storage persistence helpers.

## Validated Phase B

Phase B validated:

- Fresh Alice/Bob providers can be constructed.
- Alice MemoryStorage can be loaded from file.
- Bob MemoryStorage can be loaded from file.
- Alice group can be reloaded with `MlsGroup::load`.
- Bob group can be reloaded with `MlsGroup::load`.
- Loaded Alice group preserves epoch/member state.
- Loaded Bob group preserves epoch/member state.
- Alice signer can be reloaded from temp JSON.
- Loaded Alice group can create phase-B message two.
- Loaded Bob group can process/open phase-B message two.
- Bob plaintext matches Alice plaintext after reload.

## Important Failure Before Success

Phase B initially failed with:

- `ValidationError(InvalidSignature)`

Cause:

- Phase B created a fresh Alice signer.

Meaning:

- Bob correctly rejected a message signed by a different Alice signing key.
- This was good protocol behavior.
- Signer identity persistence is required for restart-continuity.

Fix:

- Save Alice signer in Phase A.
- Reload Alice signer in Phase B.
- Use the reloaded signer for message two.

## Storage Finding

`openmls_memory_storage::MemoryStorage` has feature-gated persistence helpers.

The scratch crate now enables MemoryStorage persistence and uses the helper path to save/load storage.

The storage is still scratch-level.

It serializes the in-memory key/value map and is not a final secure vault model.

## Current Feasibility Conclusion

OpenMLS remains viable as the first CarbonStack MLS provider candidate.

The project has now validated:

- local setup material
- group creation
- add member
- Welcome join
- application message protect/open
- two-message state continuity
- same-process group reload
- process-restart-shaped MemoryStorage file persistence
- signer reload needed for continued valid messages

This is a major Phase 2C feasibility unlock.

## What This Still Does Not Prove

This does not prove:

- production storage
- secure key storage
- encrypted local vault design
- hardware-key storage
- real CarbonStack provider adapter
- Go/Rust compatibility
- sidecar design
- CarbonStackComms CLI integration
- CarbonStackCypher MLS routing
- trust-state integration
- hostile-server security
- replay resistance
- metadata privacy
- production E2EE

## Provider Boundary Implications

The provider boundary likely needs explicit concepts for:

- provider-local storage
- storage load/save/checkpoint
- group reload by conversation/group ID
- signer identity persistence
- provider state checkpoint after outbound message creation
- provider state checkpoint after inbound message processing
- restart/recovery error classes
- storage corruption/missing-state error classes

The provider cannot be modeled as a stateless helper.

## Trust Implications

The invalid-signature failure is useful.

It shows that if signing identity changes after restart, Bob rejects the message.

CarbonStack should eventually surface this kind of failure through:

- trust-state warnings
- identity-change events
- reverify requirements
- blocked sends/receives where appropriate

This is not implemented yet.

## Safe Commands

Run the scratch persistence probe:

cd C:\Users\udaiv\repos\carbonstack_umbrella\carbonstack-comms\internal\protocol\mls\research\openmls-minimal
cargo check
cargo run -- phase-a
cargo run -- phase-b

Run the Rust artifact guard:

cd C:\Users\udaiv\repos\carbonstack_umbrella\carbonstack-comms
powershell -ExecutionPolicy Bypass -File .\scripts\check-no-rust-artifacts.ps1

Run full local validation:

cd C:\Users\udaiv\repos\carbonstack_umbrella\carbonstack
powershell -ExecutionPolicy Bypass -File .\scripts\validate-local.ps1

## Files/Artifacts Warning

The scratch persistence probe writes test storage/signing files under the OS temp directory.

Do not commit generated provider-state or signer files.

Do not treat temp JSON signer persistence as production-safe.

## Recommended Next Work

Next work should be documentation and compatibility shaping, not immediate Comms integration.

Recommended order:

1. Update provider-boundary docs with persistence lessons.
2. Add a fixture compatibility spike plan.
3. Decide what Rust should emit as provider-contract fixtures.
4. Generate sanitized fixtures if useful.
5. Teach Go-side provider boundary docs/tests to consume fixture-shaped concepts.
6. Only later decide between sidecar, FFI, or another integration model.

## Allowed Claims

Allowed:

- CarbonStack's Rust-only OpenMLS scratch crate can persist/load MemoryStorage state at scratch level.
- Fresh phase-B providers can load saved MemoryStorage files and reload Alice/Bob MlsGroups.
- A reloaded Alice signer is needed for Bob to validate post-reload messages.
- The scratch crate can send/open a second message after file save/load and fresh provider construction.

## Not Allowed Claims

Not allowed:

- CarbonStack has production E2EE.
- CarbonStack has production storage.
- signer JSON persistence is secure.
- OpenMLS is integrated into CarbonStackComms messaging.
- Cypher routes MLS payloads.
- hostile-server security is solved.
- replay resistance is solved.
- final provider architecture is selected.
