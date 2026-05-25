# OpenMLS Sidecar Conversation-Join API Recon v0

Status: API recon / pre-implementation result
Component: CarbonStackComms / OpenMLS sidecar
Phase: Phase 2D
Previous docs:
- docs/69-openmls-sidecar-add-member-welcome-export-result-v0.md
- docs/70-openmls-sidecar-conversation-join-skeleton-v0.md

## 1. Purpose

This document records targeted OpenMLS API reconnaissance for the next implementation checkpoint:

    conversation-join --device-label <joining-device> --conversation-label <conversation> --welcome <path>

The goal is to prepare Bob-side join implementation without starting implementation in this checkpoint.

## 2. Current project state before implementation

v0.2.32 validated:

- Alice identity-create;
- Bob identity-create;
- Bob public-bundle-export --write-artifact;
- Alice conversation-create;
- Alice conversation-add-member using Bob public-bundle.keypackage.bin;
- Alice Welcome carrier export as welcome.bin;
- Alice merge_pending_commit;
- Alice provider-storage.json save after membership mutation;
- later Alice conversation-load-check reports member_count=2 and group_reloadable=true.

v0.2.33 docs/70 defines the conversation-join skeleton.

conversation-join is not implemented yet.

## 3. Command target

Planned command:

    conversation-join --device-label <joining-device> --conversation-label <conversation> --welcome <path>

Initial v0 target:

    conversation-join --device-label carbonstack-bob-device --conversation-label carbonstack-test-conversation --welcome .\.carbonstack-openmls-sidecar-state\dev\conversations\carbonstack-test-conversation\welcome.bin

Meaning:

- joining-device is Bob's local device label;
- conversation-label is the desired local label for Bob's joined conversation state;
- welcome is Alice's exported Welcome carrier artifact from conversation-add-member.

Manifest-driven input remains deferred.

## 4. Welcome artifact type decision

v0.2.32 writes welcome.bin from:

    welcome_message.to_bytes()

where welcome_message is the second MlsMessageOut returned by:

    group.add_members(&provider, &signer, &[member_key_package])

v0.2.33 recon confirms this remains a usable artifact strategy.

The OpenMLS example path is:

    let mls_message_in = MlsMessageIn::tls_deserialize(&mut serialized_welcome.as_slice())?;
    let welcome = match mls_message_in.extract() {
        MlsMessageBodyIn::Welcome(welcome) => welcome,
        _ => unreachable!("Unexpected message type."),
    };

Therefore, conversation-join should initially consume the v0.2.32 outer Welcome carrier artifact as:

    MlsMessageIn::tls_deserialize(...)
    mls_message_in.extract()
    MlsMessageBodyIn::Welcome(welcome)

Then pass the inner Welcome into StagedWelcome::new_from_welcome.

No v0.2.32 artifact rewrite is required at this point.

## 5. Join API shape

Recon confirms:

    StagedWelcome::new_from_welcome<Provider: OpenMlsProvider>(
        provider: &Provider,
        mls_group_config: &MlsGroupJoinConfig,
        welcome: Welcome,
        ratchet_tree: Option<RatchetTreeIn>,
    ) -> Result<StagedWelcome, WelcomeError<Provider::StorageError>>

Recon also confirms:

    StagedWelcome::into_group<Provider: OpenMlsProvider>(
        self,
        provider: &Provider,
    ) -> Result<MlsGroup, WelcomeError<Provider::StorageError>>

Recommended v0.2.34 path:

    let mls_message_in = MlsMessageIn::tls_deserialize(&mut welcome_bytes.as_slice())?;

    let welcome = match mls_message_in.extract() {
        MlsMessageBodyIn::Welcome(welcome) => welcome,
        _ => return Err(...),
    };

    let join_config = MlsGroupJoinConfig::builder().build();

    let staged_welcome = StagedWelcome::new_from_welcome(
        &bob_provider,
        &join_config,
        welcome,
        None,
    )?;

    let bob_group = staged_welcome.into_group(&bob_provider)?;

Exact join config construction should be verified during implementation because the sidecar currently uses MlsGroupCreateConfig in conversation-create, while join uses MlsGroupJoinConfig.

## 6. Ratchet tree question

StagedWelcome::new_from_welcome requires:

    ratchet_tree: Option<RatchetTreeIn>

For the v0.2.32 add-member flow, Alice group creation used the existing conversation-create config. If the group uses the ratchet tree extension, Welcome may contain sufficient information for Bob to join with:

    None

If join fails with a ratchet-tree-related error, v0.2.34 may need to export/import a ratchet tree artifact or adjust the conversation-create/add-member config.

Recommended v0.2.34 first attempt:

    ratchet_tree=None

Record the exact result in the implementation result doc.

## 7. Bob-side state layout decision

The old Alice-side conversation state layout is:

    .carbonstack-openmls-sidecar-state/dev/conversations/<conversation-label>/

For Bob join, this risks state collision in local tests because Alice and Bob share the same sidecar root.

Recommended v0.2.34 Bob-side join layout:

    .carbonstack-openmls-sidecar-state/dev/devices/<device-label>/conversations/<conversation-label>/

Expected Bob-side files:

    conversation-summary.json
    provider-storage.json
    join-summary.json

This avoids overwriting Alice's global conversation state from v0.2.32.

Longer-term cleanup should migrate Alice conversation-create to the same device-scoped layout, but that migration should not be combined with the first join implementation unless necessary.

## 8. Recommended v0.2.34 operation sequence

Recommended implementation sequence:

1. validate device label;
2. validate conversation label;
3. parse --welcome;
4. validate Welcome path;
5. reject forbidden paths such as signer.json, provider-storage.json, identity-state.json, and directories;
6. require Bob identity state;
7. create fresh CarbonStackSidecarProvider for Bob;
8. read welcome.bin bytes;
9. deserialize as MlsMessageIn;
10. call extract();
11. require MlsMessageBodyIn::Welcome(welcome);
12. construct MlsGroupJoinConfig;
13. call StagedWelcome::new_from_welcome(&bob_provider, &join_config, welcome, None);
14. call staged_welcome.into_group(&bob_provider);
15. save Bob provider storage to device-scoped provider-storage.json;
16. reload Bob group from saved provider storage if possible;
17. write device-scoped conversation-summary.json;
18. write join-summary.json;
19. return sanitized stdout.

## 9. Proposed generated files

Recommended Bob-side path:

    .carbonstack-openmls-sidecar-state/dev/devices/<device-label>/conversations/<conversation-label>/

Files:

    conversation-summary.json
    provider-storage.json
    join-summary.json

Rules:

- provider-storage.json is generated dev OpenMLS provider/group state and must not be printed, pasted, inspected casually, committed, or treated as production secure storage.
- join-summary.json and conversation-summary.json must be sanitized metadata.
- welcome.bin must not be printed or committed.

## 10. Proposed success fields

conversation-join success should report:

    ok=true
    command=conversation-join
    phase=phase2d-conversation-join-dev
    private_material_included=false
    joined=true
    provider_storage_written=true
    group_reloadable=true
    provider_storage_loaded=true
    member_count=2
    epoch=GroupEpoch(1)
    group_id_ref=sha256:<hex>
    group_id_len=<n>
    conversation_state_path_hint=<device-scoped path>
    conversation_summary_path_hint=<device-scoped summary>
    provider_storage_path_hint=<device-scoped provider-storage.json>
    join_summary_path_hint=<device-scoped join-summary.json>
    welcome_artifact_path_hint=<input path>
    state_scope=dev-local-sidecar-state

Potential events:

    conversation.joined
    storage.saved

## 11. Proposed failure classes

Missing argument:

    missing_required_argument

Invalid labels:

    invalid_device_label
    invalid_conversation_label

Invalid Welcome path:

    invalid_welcome_path

Missing input:

    identity_missing
    welcome_missing

Invalid/corrupt input:

    welcome_invalid
    welcome_unexpected_message_type

Join failure:

    conversation_join_failed
    provider_group_unavailable
    provider_storage_save_failed

Duplicate state:

    conversation_already_exists

## 12. Required tests for v0.2.34 implementation

Go contract tests should cover:

- provider-info lists conversation-join as supported after implementation;
- unsupported-command test moves from conversation-join to message-protect or message-open;
- missing --device-label fails;
- missing --conversation-label fails;
- missing --welcome fails;
- invalid device label fails;
- invalid conversation label fails;
- missing identity fails;
- missing Welcome file fails;
- invalid Welcome artifact fails;
- duplicate join refuses overwrite;
- success flow:
  - remove state;
  - identity-create Alice;
  - identity-create Bob;
  - public-bundle-export Bob --write-artifact;
  - conversation-create Alice;
  - conversation-add-member Alice with Bob KeyPackage artifact;
  - conversation-join Bob with Alice welcome.bin;
- success asserts:
  - ok=true;
  - command=conversation-join;
  - joined=true;
  - provider_storage_written=true;
  - group_reloadable=true;
  - member_count=2;
  - device-scoped provider-storage.json exists;
  - device-scoped conversation-summary.json exists;
  - join-summary.json exists if implemented;
  - stdout contains no private material.

## 13. Open implementation questions

v0.2.34 still must prove:

1. exact imports for MlsMessageIn and MlsMessageBodyIn;
2. exact join config construction syntax;
3. whether StagedWelcome::new_from_welcome(..., None) works with the current add-member output;
4. whether Bob provider storage can save/load the joined group immediately;
5. whether Bob load-check should reuse existing conversation-load-check or require a device-scoped load-check variant;
6. whether group_id_ref after join matches Alice's group_id_ref;
7. whether device-scoped conversation state should become the default immediately or only for join.

## 14. Recommended implementation boundary

v0.2.34 should implement only:

    conversation-join

It should not implement:

- message-protect;
- message-open;
- Comms runtime send/inbox integration;
- Cypher routing;
- trust-state mutation;
- production secure vault;
- Android;
- CarbonStackOS.

## 15. Next checkpoint after join

After conversation-join succeeds:

    message-protect/message-open skeleton and API recon

Then implementation should likely split:

    message-protect artifact export
    message-open artifact consume/decrypt
    two-message continuity

Do not combine join with message protect/open.
