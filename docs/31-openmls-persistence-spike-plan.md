# CarbonStack OpenMLS Persistence Spike Plan

## Status

Classification: PHASE 2C STORAGE/PERSISTENCE DESIGN / PRE-INTEGRATION

This document defines the next OpenMLS research spike after the v0.2.6 two-message state-continuity probe.

It does not implement provider persistence.

It does not wire OpenMLS into CarbonStackComms.

It does not claim process restart recovery works.

## Context

CarbonStack has validated a Rust-only OpenMLS scratch ladder:

- dependency/build probe
- credential and KeyPackage probe
- Alice group creation
- Bob add from KeyPackage
- Bob join from Welcome
- Alice-to-Bob application message protect/open
- two-message state-continuity inside one process

The latest scratch probe proves in-memory group/provider state remains usable across sequential messages inside one process.

It does not prove disk persistence or process restart recovery.

## Purpose

The next technical question is:

Can OpenMLS provider/group state be persisted, scoped, restored, and used safely enough for CarbonStack's provider boundary?

CarbonStack needs this answer before any Go/Rust compatibility spike or Comms integration.

## Why Persistence Matters

OpenMLS scratch work showed that provider operations are stateful.

Important lessons:

- Alice and Bob need separate provider/storage instances.
- Shared provider/storage caused `GroupAlreadyExists`.
- `create_message` required mutable Alice group state.
- `process_message` required mutable Bob group state.
- Both outbound protection and inbound processing may mutate local provider/group state.

Therefore CarbonStack must eventually persist provider state after both sends and receives.

## Non-Goals

Do not:

- wire OpenMLS into `comms send`
- wire OpenMLS into `comms inbox`
- route OpenMLS payloads through CarbonStackCypher yet
- implement secure vault storage
- implement hardware-key storage
- implement production key management
- claim production E2EE
- claim replay resistance
- claim hostile-server security
- implement custom cryptography
- hide persistence gaps behind vague abstractions

## Key Questions

### Provider storage

- What storage backend does `OpenMlsRustCrypto::default()` use?
- Is it memory-only?
- Is there a file-backed or SQLite-backed provider available?
- Is `openmls_memory_storage` only suitable for scratch tests?
- Can the provider storage be replaced cleanly?

### Group reload

- Can an `MlsGroup` be loaded from provider storage by group ID?
- What API loads it?
- What must be supplied to load it?
- Does loading require configuration material beyond group ID?
- Does loading require serialized group bytes?
- Is the group state stored automatically after operations?

### Device-local state

- How should Alice provider storage and Bob provider storage be separated?
- Can storage be scoped by device?
- Can storage be scoped by account/device identity?
- Can storage be moved into a future secure vault?

### Send/receive checkpointing

- What state changes after `create_message`?
- What state changes after `process_message`?
- What state changes after joining from Welcome?
- What state changes after merging a pending commit?
- When must CarbonStack force a persistence checkpoint?

### Export/import

- Does OpenMLS expose safe serialization for group state?
- Does the provider expose a storage export path?
- Is direct export discouraged?
- Should CarbonStack treat provider state as opaque database-backed storage instead?

### Error behavior

- What happens if state is missing?
- What happens if group state exists twice?
- What happens if Bob processes a message after losing local state?
- What happens if Alice sends after losing local state?
- What happens if stale provider state processes a newer message?

## Proposed Spike Rungs

### Rung 1: local docs/API inspection

Goal:

Identify the relevant OpenMLS 0.8.1 storage and group-load APIs using local docs and source search.

Look for:

- `OpenMlsRustCrypto`
- `OpenMlsProvider`
- `StorageProvider`
- `openmls_memory_storage`
- `MlsGroup::load`
- `load_group`
- `group_id`
- `store`
- `delete`
- `write`
- `read`
- `MlsGroupState`
- `StoredMlsGroup`
- `ProviderStorage`

Output:

- pasted local method signatures
- notes on possible persistence API
- no code changes unless the API is clear

### Rung 2: same-run load/reload probe

Goal:

Determine whether a group can be reloaded from the same provider storage during one process.

Possible flow:

1. Create Alice provider.
2. Create Bob provider.
3. Create Alice group.
4. Add Bob.
5. Bob joins.
6. Send/open message one.
7. Drop or stop using direct group handles if possible.
8. Reload group state from provider storage if API supports it.
9. Send/open message two from reloaded state.

Expected outcome:

- either reload works
- or API limits are documented clearly

### Rung 3: file-backed storage investigation

Goal:

Determine whether OpenMLS has a practical disk-backed storage option.

Possible paths:

- built-in provider storage backend
- SQLite provider feature
- custom storage provider
- sidecar-managed storage
- opaque provider database

Output:

- storage backend recommendation
- integration risk notes

### Rung 4: process restart simulation

Goal:

If a disk-backed or export/import path is found, test state across a real process restart.

Possible flow:

1. Run phase A: create state and write artifacts/storage.
2. Exit.
3. Run phase B: reload state.
4. Send/open next message.
5. Validate plaintext match.

This rung should not be attempted until Rung 1 and Rung 2 are understood.

## Expected Outputs

The persistence spike should produce one of these results:

### Outcome A: simple provider reload works

CarbonStack can likely map provider state into a future local provider storage layer with manageable complexity.

### Outcome B: storage works but requires custom provider/storage implementation

CarbonStack should design a provider-owned storage adapter before CLI integration.

### Outcome C: memory-only scratch is easy but persistence is nontrivial

Continue with fixtures/sidecar research before integration.

### Outcome D: OpenMLS persistence model is too awkward for current project phase

Re-evaluate mls-rs or defer provider integration.


## Current Result: Same-Process Provider Storage Reload

The first persistence spike rung has passed in the Rust-only scratch crate.

Validated:

- Alice and Bob still use separate provider/storage instances.
- Alice sends message one.
- Bob opens message one.
- Alice group is loaded from Alice provider storage using `MlsGroup::load`.
- Bob group is loaded from Bob provider storage using `MlsGroup::load`.
- Loaded Alice and Bob groups preserve epoch and member count.
- Loaded Alice group can create message two.
- Loaded Bob group can process/open message two.
- Bob-opened plaintext for message two matches Alice plaintext.

Important conclusion:

- OpenMLS provider storage contains usable group state inside the same process.
- CarbonStack needs a provider load/reload concept.
- This is stronger than merely keeping mutable group variables alive.

Still not validated:

- disk-backed storage
- process restart recovery
- portable state export/import
- secure vault mapping
- custom CarbonStack provider storage adapter
- Go/Rust compatibility
- Comms/Cypher integration

Next rung:

- identify real disk-backed provider storage or a practical storage adapter strategy.

## Compatibility Spike Dependency

The Go/Rust compatibility spike should not begin until this persistence spike answers at least:

- how provider state is stored
- how provider state is reloaded
- when state must be checkpointed
- whether Rust should own a provider database
- whether Go should call a sidecar or consume fixtures first

## Recommended Next Workflow

1. Inspect local OpenMLS docs/source for storage APIs.
2. Paste exact relevant signatures.
3. Patch only after API shape is known.
4. Stop on compiler errors.
5. Record storage lessons before integration decisions.

## Allowed Claims

Allowed:

- CarbonStack has validated OpenMLS in-memory two-message state continuity.
- CarbonStack has identified provider persistence as the next blocker.
- CarbonStack recognizes OpenMLS provider operations as stateful.

## Not Allowed Claims

Not allowed:

- OpenMLS provider persistence works.
- process restart recovery works.
- provider storage/export is solved.
- OpenMLS is integrated into CarbonStackComms.
- CarbonStack has production E2EE.
- hostile-server security is solved.

