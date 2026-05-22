# CarbonStack OpenMLS Storage Decision v0

## Status

Classification: PHASE 2C STORAGE DECISION / SCRATCH-ONLY

This document records the storage-path decision after the OpenMLS same-process reload probe.

It does not implement production storage.

It does not wire OpenMLS into CarbonStackComms.

It does not claim process restart recovery works yet.

## Current Proven Baseline

CarbonStack has validated:

- OpenMLS dependency/build probe
- credential and KeyPackage creation
- Alice MLS group creation
- Bob add from KeyPackage
- Bob join from Welcome
- Alice-to-Bob application message protect/open
- two-message state continuity inside one process
- same-process `MlsGroup::load(provider.storage(), group_id)` reload for Alice and Bob
- second message after same-process group reload

## Storage Search Findings

OpenMLS `MlsGroup::load(storage, group_id)` can load group state from provider storage.

`OpenMlsProvider` exposes:

- `storage()`
- `crypto()`
- `rand()`

`openmls_rust_crypto::OpenMlsRustCrypto` uses:

- `RustCrypto`
- `MemoryStorage`

But its internal storage field is private.

This means CarbonStack cannot cleanly load file state into the stock `OpenMlsRustCrypto` instance through a public mutable storage accessor.

## MemoryStorage Persistence Finding

`openmls_memory_storage::MemoryStorage` has a feature-gated persistence module.

The persistence module provides public methods on `MemoryStorage`:

- `save_to_file`
- `save(user_name)`
- `load_from_file`
- `load(user_name)`

The implementation serializes the in-memory key/value map to JSON using base64-encoded keys and values.

This is acceptable for scratch feasibility testing.

This is not acceptable as a final production storage model by itself.

## Decision

Proceed with Option A-lite:

- enable/use `openmls_memory_storage` persistence in the scratch crate
- define a CarbonStack-only scratch provider wrapper
- the scratch provider owns `RustCrypto` and `MemoryStorage`
- implement `OpenMlsProvider` for the scratch provider
- use `MemoryStorage::save` / `MemoryStorage::load` for a process-restart-shaped persistence probe

## Why This Path

This path tests real file save/load behavior without prematurely designing:

- custom production storage
- Rust sidecar APIs
- Go/Rust FFI
- secure vault integration
- CarbonStackComms provider integration

It is the smallest next rung after same-process reload.

## Explicit Non-Decision

This does not decide:

- final OpenMLS storage architecture
- final provider sidecar design
- final secure vault model
- final Comms integration model
- final OS integration model

## Next Spike

Create a process-restart-shaped scratch experiment.

Proposed modes:

- `phase-a`: create Alice/Bob group state, send/open message one, save Alice/Bob provider storage files.
- `phase-b`: create fresh providers, load Alice/Bob provider storage files, reload groups with `MlsGroup::load`, send/open message two.

Success means:

- provider storage survives through file save/load
- loaded Alice/Bob groups can continue the conversation after reload
- plaintext for message two matches after reload

Failure means:

- document the blocker
- likely choose fixture-based compatibility spike before sidecar/custom storage

## Allowed Claims

Allowed:

- CarbonStack found a plausible scratch persistence path through `MemoryStorage`.
- CarbonStack will test MemoryStorage file persistence before Go/Rust compatibility.
- The stock `OpenMlsRustCrypto` wrapper does not expose a convenient mutable storage load path.

## Not Allowed Claims

Not allowed:

- production storage is solved
- process restart recovery is already validated
- secure vault storage is solved
- OpenMLS is integrated into CarbonStackComms
- CarbonStack has production E2EE
- hostile-server security is solved
