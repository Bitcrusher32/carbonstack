# OpenMLS Sidecar Dev Provider/Group Persistence Plan v0

Status: Planned
Component: CarbonStackComms / OpenMLS sidecar
Phase: Phase 2D
Previous docs:
- docs/63-openmls-sidecar-conversation-create-persistence-repair-plan-v0.md
- docs/64-openmls-sidecar-conversation-create-persistence-repair-result-v0.md

## 1. Purpose

This document defines the next implementation rung after the v0.2.29 persistence honesty repair.

v0.2.29 corrected conversation-create to report:

    provider_storage_written=false
    group_reloadable=false

That was correct because the current sidecar uses OpenMlsRustCrypto::default(), which owns an in-memory MemoryStorage. A fresh sidecar command receives a fresh empty provider/storage instance.

This checkpoint will attempt to make dev-local OpenMLS provider/group persistence work across sidecar invocations.

## 2. Goal

The goal is to make this true:

    conversation-create creates a one-member OpenMLS group and persists enough dev-local provider/group state that a later command can load it with MlsGroup::load.

New proof command:

    conversation-load-check --device-label <safe> --conversation-label <safe>

Expected after success:

    conversation-create:
      provider_storage_written=true
      group_reloadable=true

    conversation-load-check:
      group_reloadable=true
      member_count=1
      epoch=GroupEpoch(0)
      private_material_included=false

## 3. Key recon findings

OpenMlsRustCrypto is built from:

    RustCrypto
    MemoryStorage

But its MemoryStorage field is private. This makes OpenMlsRustCrypto::default() inconvenient for explicit sidecar save/load.

The earlier OpenMLS scratch crate already validated an Option A-lite persistence path:

- define a CarbonStack-owned provider wrapper;
- make it own RustCrypto + MemoryStorage;
- implement OpenMlsProvider;
- call MemoryStorage.save(...) after group state changes;
- call MemoryStorage.load(...) in a later fresh provider;
- call MlsGroup::load(provider.storage(), &group_id);
- continue the conversation after reload.

Important scratch lesson:

Provider/group storage persistence is not enough by itself. Signer identity persistence is also required. v0.2.22+ already persists signer.json as dev-only sidecar state.

## 4. Implementation approach

Add a sidecar-local provider module or helper:

    CarbonStackSidecarProvider {
      crypto: RustCrypto,
      key_store: MemoryStorage,
    }

It should implement OpenMlsProvider:

    type CryptoProvider = RustCrypto
    type RandProvider = RustCrypto
    type StorageProvider = MemoryStorage

It should expose dev-only helper methods:

    load_storage(name_or_path)
    save_storage(name_or_path)

Exact API names depend on openmls_memory_storage persistence methods.

Recommended file:

    carbonstack-comms/internal/protocol/mls/research/openmls-sidecar/src/provider.rs

Or, if keeping small for this rung:

    state.rs-local helper struct first, refactor later.

## 5. State layout

Existing conversation state:

    .carbonstack-openmls-sidecar-state/dev/conversations/<conversation-label>/
      conversation-summary.json

New expected generated dev persistence file:

    provider-storage.json

or:

    memory-storage.json

Recommended path:

    .carbonstack-openmls-sidecar-state/dev/conversations/<conversation-label>/provider-storage.json

The exact filename should be boring and explicit.

Important:

This is generated dev state. It is not production secure storage, not encrypted vault storage, and not hardware-backed storage.

## 6. Conversation group ID

Use the same deterministic dev group ID as v0.2.27/v0.2.29:

    carbonstack-openmls-dev-conversation:<conversation-label>

Then:

    GroupId::from_slice(group_id_bytes)

This must stay stable or load-check will fail.

## 7. conversation-create behavior after this rung

conversation-create should:

1. validate device label;
2. validate conversation label;
3. require existing identity state;
4. load signer state locally without printing it;
5. create CarbonStackSidecarProvider;
6. create OpenMLS group with MlsGroup::new_with_group_id;
7. save provider MemoryStorage to conversation provider-storage file;
8. immediately attempt MlsGroup::load from the saved/reloaded provider if practical;
9. write conversation-summary.json;
10. report honest persistence fields.

Expected if reload proof succeeds:

    provider_storage_written=true
    group_reloadable=true

If save succeeds but load proof fails:

    provider_storage_written=true
    group_reloadable=false

If save is not implemented:

    provider_storage_written=false
    group_reloadable=false

Do not overclaim.

## 8. New conversation-load-check command

Command:

    conversation-load-check --device-label <safe> --conversation-label <safe>

Behavior:

1. validate device label;
2. validate conversation label;
3. require existing identity state;
4. require conversation-summary.json;
5. require provider-storage file;
6. create fresh CarbonStackSidecarProvider;
7. load provider MemoryStorage from file;
8. derive deterministic GroupId from conversation label;
9. call MlsGroup::load(provider.storage(), &group_id);
10. return sanitized result.

Success stdout should include:

    conversation_label
    device_label
    provider_storage_loaded=true
    group_reloadable=true
    member_count
    epoch
    group_id_ref
    private_material_included=false

Failure stdout should distinguish:

    identity_missing
    conversation_missing
    provider_storage_missing
    provider_storage_unloadable
    group_missing
    group_load_failed

## 9. Provider-info changes

After implementation, provider-info should list:

    conversation-load-check

as supported.

The unsupported list should still include:

    conversation-add-member
    conversation-join
    message-protect
    message-open
    state-checkpoint
    state-load-check

conversation-add-member remains unsupported until persistence is proven.

## 10. Event/trust mapping

Potential new events:

    provider.conversation.loaded
    provider.conversation.missing
    provider.storage.loaded
    provider.storage.unavailable

Initial trust classification:

- dev-sidecar only;
- append history / debug-only;
- no trust-state mutation;
- not user visible;
- not production security meaningful.

However:

provider.storage.unavailable should stop the current operation.

## 11. Tests required

Go-side tests should cover:

- provider-info lists conversation-load-check as supported;
- conversation-add-member remains unsupported;
- conversation-load-check missing device label fails;
- conversation-load-check missing conversation label fails;
- invalid labels fail;
- missing identity fails;
- missing conversation fails;
- successful flow:
  - identity-create Alice;
  - conversation-create Alice;
  - conversation-load-check Alice;
- successful load-check reports group_reloadable=true;
- load-check reports member_count=1;
- load-check reports epoch;
- stdout contains no obvious secret material;
- generated provider storage is not printed.

Rust-side tests may cover provider wrapper save/load if feasible.

## 12. Manual probe after implementation

From sidecar crate:

    Remove-Item -Recurse -Force .\.carbonstack-openmls-sidecar-state -ErrorAction SilentlyContinue

    cargo run -- identity-create --device-label carbonstack-alice-device

    cargo run -- conversation-create --device-label carbonstack-alice-device --conversation-label carbonstack-test-conversation

    cargo run -- conversation-load-check --device-label carbonstack-alice-device --conversation-label carbonstack-test-conversation

Safe inspection:

    Get-ChildItem .\.carbonstack-openmls-sidecar-state\dev\conversations\carbonstack-test-conversation | Select-Object Name, Length

    Get-Content .\.carbonstack-openmls-sidecar-state\dev\conversations\carbonstack-test-conversation\conversation-summary.json

Do not inspect:

    signer.json
    provider-storage.json
    memory-storage.json
    raw group state
    raw key material

## 13. Non-goals

This rung must not implement:

- conversation-add-member;
- Welcome export;
- conversation-join;
- message-protect;
- message-open;
- Comms runtime integration;
- Cypher routing;
- trust-state mutation;
- production secure vault;
- hardware-backed identity;
- Android;
- CarbonStackOS.

## 14. Success criteria

This checkpoint succeeds when:

- docs/65 lands before code;
- sidecar provider/group persistence is implemented or honestly blocked;
- conversation-load-check exists if implementation proceeds;
- MlsGroup::load succeeds across a fresh provider invocation if persistence works;
- conversation-create updates provider_storage_written/group_reloadable honestly;
- tests protect the claim;
- artifact guard passes;
- result doc records behavior and limitations.

## 15. Next after success

Only after provider/group persistence and conversation-load-check validate should CarbonStack return to:

    conversation-add-member / Welcome export skeleton doc

Then:

    conversation-add-member implementation
