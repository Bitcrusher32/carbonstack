# OpenMLS Sidecar Conversation-Join Result v0

Status: Validated implementation checkpoint
Component: CarbonStackComms / OpenMLS sidecar
Phase: Phase 2D
Previous docs:
- docs/69-openmls-sidecar-add-member-welcome-export-result-v0.md
- docs/70-openmls-sidecar-conversation-join-skeleton-v0.md
- docs/71-openmls-sidecar-conversation-join-api-recon-v0.md

## 1. Summary

This checkpoint implements the first dev-local OpenMLS conversation-join path.

The new sidecar command is:

    conversation-join --device-label <joining-device> --conversation-label <conversation> --welcome <path>

The validated flow is:

1. Alice identity exists.
2. Bob identity exists.
3. Bob exports a public KeyPackage artifact.
4. Bob's public-bundle export also saves Bob device provider storage.
5. Alice creates a conversation.
6. Alice adds Bob's KeyPackage to the conversation.
7. Alice exports a Welcome carrier artifact as welcome.bin.
8. Bob consumes welcome.bin.
9. Bob stages and joins the group.
10. Bob saves joined device-scoped conversation provider storage.
11. Bob's joined group reload proof succeeds.

This checkpoint completes the first dev-local MLS membership onboarding lifecycle:

    create -> add-member -> Welcome export -> join

This checkpoint does not implement message protect/open.

## 2. Implemented command

New command:

    conversation-join

Required arguments:

    --device-label <joining-device>
    --conversation-label <conversation>
    --welcome <path>

Validated example:

    cargo run -- conversation-join --device-label carbonstack-bob-device --conversation-label carbonstack-test-conversation --welcome .\.carbonstack-openmls-sidecar-state\dev\conversations\carbonstack-test-conversation\welcome.bin

## 3. Implementation files

CarbonStackComms changed:

    internal/protocol/mls/research/openmls-sidecar/src/state.rs
    internal/protocol/mls/research/openmls-sidecar/src/main.rs
    internal/protocol/openmls_sidecar_provider_info_test.go

Key state.rs additions:

    device_provider_storage_path(...)
    device_conversations_dir(...)
    device_conversation_state_dir(...)
    device_conversation_summary_path(...)
    device_conversation_provider_storage_path(...)
    device_conversation_join_summary_path(...)
    ConversationJoinResult
    JoinedConversationSummary
    JoinSummary
    validate_welcome_artifact_path(...)
    join_dev_conversation(...)

Key main.rs additions:

    PHASE_CONVERSATION_JOIN
    parse_welcome_artifact_path(...)
    handle_conversation_join(...)
    print_conversation_join_success(...)
    print_conversation_join_missing_argument(...)
    print_conversation_join_invalid_label(...)
    print_conversation_join_failed(...)

Provider-info was updated so conversation-join is a capability, not an unsupported command.

## 4. Critical repair: Bob provider storage persistence

Initial manual join failed with:

    NoMatchingKeyPackage

Cause:

- Bob's public KeyPackage artifact existed as public-bundle.keypackage.bin.
- Bob's private KeyPackage bundle / HPKE private material existed only inside the temporary provider used during public-bundle-export.
- That provider storage was not saved.
- Later, conversation-join created a fresh provider with no matching private KeyPackage bundle.
- StagedWelcome::new_from_welcome could parse the Welcome, but could not find the matching KeyPackage private state.

Repair:

    public-bundle-export now saves Bob device provider storage after KeyPackage generation.

The new Bob device provider storage path is:

    .carbonstack-openmls-sidecar-state/dev/devices/<device-label>/provider-storage.json

public-bundle-export now reports:

    provider_storage_written=true

This is required for later Welcome consumption.

Important interpretation:

- The public KeyPackage artifact is not enough for Bob to join.
- Bob also needs the local private provider state generated alongside the KeyPackage.
- This private provider state is not printed.
- This storage remains dev-only and is not production secure vault storage.

## 5. Welcome consume / join API path

v0.2.33 recon predicted the correct path. v0.2.34 validated it.

The implemented Welcome consume path is:

    MlsMessageIn::tls_deserialize(&mut welcome_bytes.as_slice())
    mls_message_in.extract()
    MlsMessageBodyIn::Welcome(welcome)
    MlsGroupJoinConfig::builder().build()
    StagedWelcome::new_from_welcome(&provider, &join_config, welcome, None)
    staged_welcome.into_group(&provider)

The implementation consumes the v0.2.32 Welcome carrier artifact directly.

No change to v0.2.32 welcome.bin format was needed.

welcome.bin remains:

    MlsMessageOut::to_bytes()

from the Welcome carrier returned by:

    MlsGroup::add_members(...)

## 6. Operation sequence

The implemented conversation-join sequence is:

1. parse --device-label;
2. parse --conversation-label;
3. parse --welcome;
4. validate device label;
5. validate conversation label;
6. load Bob identity status;
7. validate Welcome artifact path;
8. reject forbidden paths such as signer.json, provider-storage.json, identity-state.json, public-bundle.keypackage.bin, public-bundle-manifest.json, public-bundle-summary.json, add-member-summary.json, and welcome-manifest.json;
9. create Bob device-scoped conversation state directory;
10. read welcome.bin bytes;
11. deserialize welcome.bin as MlsMessageIn;
12. extract the message body;
13. require MlsMessageBodyIn::Welcome(welcome);
14. create CarbonStackSidecarProvider;
15. load Bob device provider-storage.json;
16. build MlsGroupJoinConfig;
17. call StagedWelcome::new_from_welcome(..., None);
18. call into_group(...);
19. compute member count, epoch, group id ref, and group id len;
20. save Bob joined conversation provider-storage.json;
21. reload Bob joined group from saved provider storage;
22. write device-scoped conversation-summary.json;
23. write join-summary.json;
24. return sanitized stdout.

## 7. Bob-side state layout

Bob joined conversation state is device-scoped:

    .carbonstack-openmls-sidecar-state/dev/devices/<device-label>/conversations/<conversation-label>/

Validated Bob path:

    .carbonstack-openmls-sidecar-state/dev/devices/carbonstack-bob-device/conversations/carbonstack-test-conversation/

Files written:

    conversation-summary.json
    join-summary.json
    provider-storage.json

Bob device root also contains:

    provider-storage.json
    public-bundle.keypackage.bin
    public-bundle-summary.json
    public-bundle-manifest.json
    identity-prep.json
    identity-state.json
    identity-summary.json
    signer.json

Meaning of the two provider-storage files:

    dev/devices/<device-label>/provider-storage.json

stores Bob's device provider state after KeyPackage generation and is needed to consume the Welcome.

    dev/devices/<device-label>/conversations/<conversation-label>/provider-storage.json

stores Bob's joined group state after conversation-join.

Both are dev-only local state and are not production secure storage.

## 8. Success envelope

Successful conversation-join output reports:

    ok=true
    command=conversation-join
    phase=phase2d-conversation-join-dev
    private_material_included=false
    joined=true
    provider_storage_written=true
    provider_storage_loaded=true
    group_reloadable=true
    member_count=2
    epoch=GroupEpoch(1)
    group_id_ref=<same group ref as Alice/add-member>
    group_id_len=66
    conversation_state_path_hint=<device-scoped Bob conversation dir>
    conversation_summary_path_hint=<device-scoped conversation-summary.json>
    provider_storage_path_hint=<device-scoped joined provider-storage.json>
    join_summary_path_hint=<device-scoped join-summary.json>
    welcome_artifact_path_hint=<input welcome.bin>
    state_scope=dev-local-sidecar-state

Events:

    conversation.joined
    storage.saved

Warnings include:

    dev-only OpenMLS conversation join
    Welcome artifact was consumed locally but not printed
    provider storage is dev-only and not production secure vault storage
    message protect/open is not implemented

## 9. Manual validation

Validated manual probe:

    Remove-Item -Recurse -Force .\.carbonstack-openmls-sidecar-state -ErrorAction SilentlyContinue

    cargo run -- identity-create --device-label carbonstack-alice-device

    cargo run -- identity-create --device-label carbonstack-bob-device

    cargo run -- public-bundle-export --device-label carbonstack-bob-device --write-artifact

    cargo run -- conversation-create --device-label carbonstack-alice-device --conversation-label carbonstack-test-conversation

    cargo run -- conversation-add-member --device-label carbonstack-alice-device --conversation-label carbonstack-test-conversation --member-keypackage .\.carbonstack-openmls-sidecar-state\dev\devices\carbonstack-bob-device\public-bundle.keypackage.bin

    cargo run -- conversation-join --device-label carbonstack-bob-device --conversation-label carbonstack-test-conversation --welcome .\.carbonstack-openmls-sidecar-state\dev\conversations\carbonstack-test-conversation\welcome.bin

Manual proof observed:

- public-bundle-export returned provider_storage_written=true;
- conversation-add-member returned welcome_artifact_written=true;
- conversation-add-member returned member_count_after=2;
- conversation-add-member returned epoch_after=GroupEpoch(1);
- conversation-join returned ok=true;
- conversation-join returned joined=true;
- conversation-join returned member_count=2;
- conversation-join returned epoch=GroupEpoch(1);
- conversation-join returned group_reloadable=true;
- conversation-join returned provider_storage_loaded=true;
- conversation-join returned provider_storage_written=true;
- conversation-join returned the same group_id_ref as Alice/add-member;
- conversation-join returned private_material_included=false.

Safe file listing confirmed:

    dev/devices/carbonstack-bob-device/
        conversations
        identity-prep.json
        identity-state.json
        identity-summary.json
        provider-storage.json
        public-bundle-manifest.json
        public-bundle-summary.json
        public-bundle.keypackage.bin
        signer.json

    dev/devices/carbonstack-bob-device/conversations/carbonstack-test-conversation/
        conversation-summary.json
        join-summary.json
        provider-storage.json

## 10. Go contract tests

Go test coverage was updated for:

    TestOpenMLSSidecarProviderInfoCommand
    TestOpenMLSSidecarUnsupportedCommandEnvelope
    TestOpenMLSSidecarPublicBundleExportCreatesSummary
    TestOpenMLSSidecarPublicBundleExportWritesArtifact
    TestOpenMLSSidecarConversationJoinWelcomeConsume

Provider-info changes:

- conversation-join is now expected in capabilities;
- conversation-join is no longer expected in unsupported;
- unsupported-command test moved to message-protect.

Public-bundle changes:

- public-bundle-export now reports provider_storage_written=true;
- public-bundle summary now reports provider_storage_written=true;
- public-bundle manifest now reports provider_storage_written=true;
- this is explicitly tied to storing the private KeyPackage provider state required for Welcome consumption.

Join contract test validates:

- Alice identity-create;
- Bob identity-create;
- Bob public-bundle-export --write-artifact;
- Bob provider_storage_written=true after public-bundle export;
- Alice conversation-create;
- Alice conversation-add-member;
- Welcome artifact path exists in add-member result;
- Bob conversation-join succeeds;
- joined=true;
- provider_storage_written=true;
- provider_storage_loaded=true;
- group_reloadable=true;
- member_count=2;
- epoch is present;
- group_id_ref is present;
- group_id_ref matches add-member group_id_ref;
- device-scoped conversation state directory exists;
- conversation-summary.json exists;
- provider-storage.json exists;
- join-summary.json exists;
- duplicate join refuses overwrite with conversation_already_exists;
- stdout does not contain private material.

## 11. Validation commands

Validated:

    go test ./internal/protocol -run "TestOpenMLSSidecarPublicBundleExportCreatesSummary|TestOpenMLSSidecarPublicBundleExportWritesArtifact|TestOpenMLSSidecarConversationJoinWelcomeConsume"

    go test ./internal/protocol

    go test ./...

    powershell -ExecutionPolicy Bypass -File .\scripts\check-no-rust-artifacts.ps1

Observed results:

    go test ./internal/protocol
    ok

    go test ./...
    ok

    check-no-rust-artifacts.ps1
    PASS: no tracked Rust/build artifacts found

## 12. Blunders and lessons from this rung

### 12.1 Welcome parsing worked, private KeyPackage state was missing

The first conversation-join manual probe reached OpenMLS staging but failed with:

    NoMatchingKeyPackage

This was not a Welcome carrier parsing failure.

Lesson:

OpenMLS join requires local provider storage containing the private KeyPackage bundle generated alongside Bob's public KeyPackage. The public KeyPackage bytes alone are insufficient.

### 12.2 public-bundle-export now has stateful consequences

Before v0.2.34, public-bundle-export was treated as a public artifact export that did not write provider storage.

After v0.2.34, this had to change:

    provider_storage_written=true

because the exported public KeyPackage and private local provider state are a pair.

Lesson:

KeyPackage export is not purely public metadata. It creates local private provider state needed for later Welcome consumption.

### 12.3 Device-scoped join state avoided Alice/Bob collision

Bob join writes under:

    dev/devices/<device-label>/conversations/<conversation-label>/

instead of the older Alice global path:

    dev/conversations/<conversation-label>/

Lesson:

Two-party tests need device-scoped joined state to avoid overwriting Alice's provider-storage.json. Longer-term, Alice conversation-create should probably migrate to the same layout, but that was intentionally not done in this rung.

### 12.4 Go test patching was fragile

Several PowerShell exact-block replacements failed due to tabs, spacing, and local struct pollution from prior patches.

Final working approach:

- use exact visible line-level replacement when possible;
- when exact blocks fail, line-walk by function scope;
- keep identity-create and identity-status storage expectations unchanged;
- only flip public-bundle storage expectations.

Lesson:

For this file, large exact-block patches are fragile. Future changes should either:
- use small line-walk patches;
- or use manual file editing;
- or split the Go test file into smaller focused test files.

### 12.5 Public-bundle local summary structs are bloated

Earlier broad patches caused public-bundle local summary structs to include fields from later envelope types.

This is not fatal because Go JSON ignores absent fields, but it increases noise.

Lesson:

A future cleanup checkpoint should slim local summary structs and maybe split tests by command.

## 13. Allowed claims after this checkpoint

Allowed:

- CarbonStack OpenMLS sidecar can create a dev-local MLS conversation.
- Alice can add Bob's public KeyPackage artifact to a persisted dev-local MLS group.
- Alice can export a Welcome carrier artifact.
- Bob can consume the Welcome carrier artifact.
- Bob can join the dev-local MLS group.
- Bob writes device-scoped joined conversation state.
- Bob's joined group is reloadable from saved provider storage.
- Bob and Alice/add-member outputs report the same group_id_ref.
- public-bundle-export now persists Bob device provider storage because the private KeyPackage bundle is needed for Welcome consumption.
- The first dev-local MLS membership onboarding lifecycle exists:
  create -> add-member -> Welcome export -> join.

Not allowed:

- message-protect exists.
- message-open exists.
- plaintext application messages are encrypted/decrypted.
- Cypher routes MLS payloads.
- Comms runtime consumes MLS payloads.
- trust-state is updated from join events.
- provider storage is production secure storage.
- CarbonStack has production E2EE.
- Android integration exists.
- CarbonStackOS integration exists.
- Signal-equivalent security exists.

## 14. Next recommended checkpoint

Next checkpoint should be docs/recon only:

    OpenMLS sidecar message-protect/message-open skeleton and API recon

Recommended future docs:

    docs/73-openmls-sidecar-message-protect-open-skeleton-v0.md
    docs/74-openmls-sidecar-message-protect-open-api-recon-v0.md

Questions to answer before implementation:

1. How does OpenMLS 0.8.1 protect application messages?
2. What exact API produces an MLS application message from plaintext?
3. What exact API opens an MLS application message?
4. How do both Alice and Bob provider storages advance after protect/open?
5. Does Bob's device-scoped load-check need a new command?
6. Should Alice conversation-create migrate to device-scoped state before message tests?
7. What artifact format should protected messages use?
8. How do we keep stdout private-material-safe?
9. What test flow proves one-way message delivery?
10. What test flow proves state continuity over multiple messages?

Recommended v0.2.35 boundary:

    docs + API recon only

Do not implement message-protect/message-open until the API path is documented.
