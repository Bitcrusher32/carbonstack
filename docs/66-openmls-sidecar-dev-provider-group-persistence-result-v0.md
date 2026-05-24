# OpenMLS Sidecar Dev Provider/Group Persistence Result v0

Status: Validated implementation checkpoint
Component: CarbonStackComms / OpenMLS sidecar
Phase: Phase 2D
Plan doc:
- docs/65-openmls-sidecar-dev-provider-group-persistence-plan-v0.md
Previous docs:
- docs/63-openmls-sidecar-conversation-create-persistence-repair-plan-v0.md
- docs/64-openmls-sidecar-conversation-create-persistence-repair-result-v0.md

## 1. Summary

This checkpoint implements dev-local OpenMLS provider/group persistence for the OpenMLS sidecar.

v0.2.29 correctly repaired an overclaim by changing conversation-create to report:

    provider_storage_written=false
    group_reloadable=false

That was correct at the time because the sidecar used OpenMlsRustCrypto::default(), whose MemoryStorage was process-local and not reloadable across sidecar command invocations.

v0.2.30 adds a CarbonStack-owned sidecar provider wrapper and dev-local MemoryStorage persistence. conversation-create now writes provider-storage.json under the conversation state directory and proves that MlsGroup::load can reload the created group from a fresh provider instance.

This unblocks future add-member planning, but add-member is still not implemented in this checkpoint.

## 2. Implementation

CarbonStackComms changed:

    internal/protocol/mls/research/openmls-sidecar/Cargo.toml
    internal/protocol/mls/research/openmls-sidecar/Cargo.lock
    internal/protocol/mls/research/openmls-sidecar/src/provider.rs
    internal/protocol/mls/research/openmls-sidecar/src/main.rs
    internal/protocol/mls/research/openmls-sidecar/src/state.rs
    internal/protocol/openmls_sidecar_provider_info_test.go

New sidecar provider wrapper:

    CarbonStackSidecarProvider

It owns:

    RustCrypto
    MemoryStorage

It implements:

    OpenMlsProvider

It exposes dev-only helper methods:

    save_storage_to_path(...)
    load_storage_from_path(...)

The sidecar now depends directly on:

    openmls_memory_storage = { version = "0.5.0", features = ["persistence"] }
    openmls_traits = "0.5.0"

## 3. Generated dev state

conversation-create now writes:

    .carbonstack-openmls-sidecar-state/dev/conversations/<conversation-label>/conversation-summary.json
    .carbonstack-openmls-sidecar-state/dev/conversations/<conversation-label>/provider-storage.json

Important:

provider-storage.json is generated dev OpenMLS provider/group state. It is not production secure vault storage. It should not be printed, inspected, committed, or treated as safe application data.

## 4. Corrected conversation-create behavior

conversation-create now reports:

    provider_storage_written=true
    group_reloadable=true
    provider_storage_path_hint=<conversation state>/provider-storage.json

conversation-summary.json now reports:

    provider_storage_written=true
    group_reloadable=true
    provider_storage_file="provider-storage.json"

The command still reports:

    private_material_included=false

and does not print signer material or provider storage contents.

## 5. New command

New sidecar command:

    conversation-load-check --device-label <safe> --conversation-label <safe>

Purpose:

- validate identity state;
- validate conversation state;
- require provider-storage.json;
- create a fresh CarbonStackSidecarProvider;
- load MemoryStorage from provider-storage.json;
- derive the deterministic dev GroupId;
- call MlsGroup::load(provider.storage(), &group_id);
- return sanitized proof fields.

Successful output includes:

    provider_storage_loaded=true
    group_reloadable=true
    member_count=1
    epoch=GroupEpoch(0)
    event=conversation.loaded
    private_material_included=false

## 6. Provider-info changes

provider-info now lists conversation-load-check as supported.

Unsupported still includes:

    conversation-add-member
    conversation-join
    message-protect
    message-open
    state-checkpoint
    state-load-check

This is intentional. Provider/group persistence is now proven, but add-member / Welcome export has not been implemented yet.

## 7. Events

Rust sidecar conversation-load-check success emits:

    conversation.loaded

This matches the existing Go-side provider event constant:

    ProviderEventConversationLoaded

The load-check event is not trust-relevant in this dev-sidecar rung.

## 8. Validated manual behavior

Manual probe:

    Remove-Item -Recurse -Force .\.carbonstack-openmls-sidecar-state -ErrorAction SilentlyContinue
    cargo run -- identity-create --device-label carbonstack-alice-device
    cargo run -- conversation-create --device-label carbonstack-alice-device --conversation-label carbonstack-test-conversation
    cargo run -- conversation-load-check --device-label carbonstack-alice-device --conversation-label carbonstack-test-conversation

Validated:

- identity-create succeeds;
- conversation-create succeeds;
- conversation-create writes provider-storage.json;
- conversation-create reports provider_storage_written=true;
- conversation-create reports group_reloadable=true;
- conversation-load-check succeeds in a later command invocation;
- conversation-load-check reports provider_storage_loaded=true;
- conversation-load-check reports group_reloadable=true;
- conversation-load-check reports member_count=1;
- conversation-load-check reports epoch=GroupEpoch(0);
- no private material is printed.

## 9. Tests and guard

Validated:

    go test ./internal/protocol -run "TestOpenMLSSidecarProviderInfoCommand|TestOpenMLSSidecarConversationCreate|TestOpenMLSSidecarConversationLoadCheck"
    go test ./internal/protocol
    go test ./...
    powershell -ExecutionPolicy Bypass -File .\scripts\check-no-rust-artifacts.ps1

Rust artifact guard passed.

## 10. Allowed claims after this checkpoint

Allowed:

- CarbonStack OpenMLS sidecar now has dev-local provider/group persistence for one-member conversation-create.
- conversation-create writes provider-storage.json.
- conversation-create proves group reloadability by saving/loading provider storage and calling MlsGroup::load.
- conversation-load-check independently reloads the group in a later sidecar invocation.
- provider-info lists conversation-load-check as supported.
- Go-side contract tests protect conversation-create persistence claims and conversation-load-check behavior.

Not allowed:

- This is production secure storage.
- This is a hardware-backed secure vault.
- provider-storage.json is safe to inspect, print, commit, or expose.
- conversation-add-member exists.
- Welcome export exists.
- conversation-join exists.
- message protect/open exists.
- Comms runtime consumes this state.
- Cypher routes MLS payloads.
- production E2EE exists.

## 11. Next recommended checkpoint

Next checkpoint:

    OpenMLS sidecar add-member / Welcome export skeleton doc

Recommended future doc:

    docs/67-openmls-sidecar-add-member-welcome-skeleton-v0.md

The next checkpoint should define:

- command surface for conversation-add-member;
- Bob KeyPackage input path;
- Alice group load path;
- add_members call;
- Welcome artifact export path;
- merge_pending_commit behavior;
- provider storage save after membership mutation;
- sanitized stdout fields;
- tests;
- non-goals.

Only after that skeleton is approved should implementation begin.
