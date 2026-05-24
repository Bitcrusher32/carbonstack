# OpenMLS Sidecar Conversation Create Plan v0

Status: Planned
Component: CarbonStackComms / OpenMLS sidecar
Phase: Phase 2D
Previous docs:
- docs/57-openmls-sidecar-conversation-lifecycle-plan-v0.md
- docs/58-openmls-sidecar-conversation-lifecycle-api-recon-v0.md

## 1. Purpose

This document defines the next narrow implementation rung for the OpenMLS sidecar:

    conversation-create --device-label <safe> --conversation-label <safe>

The goal is to create a dev-local OpenMLS group/conversation for one existing local identity, write sanitized conversation metadata, and preserve strict boundaries.

This is not group messaging support. It is not add-member. It is not Welcome handling. It is not message encryption/decryption in CarbonStack runtime.

## 2. Current baseline

Validated sidecar commands before this rung:

    provider-info
    identity-create
    identity-status
    public-bundle-export
    public-bundle-export --write-artifact

Current unsupported commands include:

    conversation-create
    conversation-add-member
    conversation-join
    message-protect
    message-open

After this rung, only conversation-create should graduate from unsupported to supported.

conversation-add-member, conversation-join, message-protect, and message-open must remain unsupported.

## 3. API reconnaissance basis

Targeted local inspection confirmed the pinned OpenMLS 0.8.1 API includes:

    MlsGroup::new_with_group_id(
        provider,
        signer,
        mls_group_create_config,
        group_id,
        credential_with_key,
    )

The old openmls-minimal research crate already demonstrated a scratch flow using:

    MlsGroupCreateConfig::builder()
        .ciphersuite(ciphersuite)
        .use_ratchet_tree_extension(true)
        .build();

    let group_id = GroupId::from_slice(GROUP_ID_BYTES);

    let mut alice_group = MlsGroup::new_with_group_id(
        &alice_provider,
        &alice.signer,
        &create_config,
        group_id,
        alice.credential_with_key.clone(),
    )?;

The current implementation should adapt only the group creation portion, not the add-member, Welcome, join, protect, or open message portions.

## 4. Command shape

Supported command after this rung:

    conversation-create --device-label <safe-device-label> --conversation-label <safe-conversation-label>

Required arguments:

    --device-label
    --conversation-label

Missing either argument should fail with a sanitized command-invalid envelope and exit code 2.

## 5. Label validation

Device labels should continue to use the existing device label validation path.

Conversation labels should receive their own validator, or a shared safe-label validator with distinct error text.

Initial conversation label rules:

- non-empty;
- not "." or "..";
- max length 96;
- ASCII alphanumeric, hyphen, and underscore only;
- no spaces;
- no slashes;
- no backslashes;
- no dots;
- no Unicode;
- no shell metacharacters.

Recommended examples:

    carbonstack-test-conversation
    test_conversation_01

Rejected examples:

    ../test
    .\test
    test conversation
    test.conversation
    test/conversation
    test\conversation

## 6. Dev-state layout

Existing device state:

    .carbonstack-openmls-sidecar-state/dev/devices/<device-label>/

New conversation state:

    .carbonstack-openmls-sidecar-state/dev/conversations/<conversation-label>/

Expected sanitized file:

    conversation-summary.json

Potential future provider/group state files are not finalized in this plan. If OpenMLS provider storage writes are necessary, they must remain ignored generated dev state and must not be described as secure vault storage.

## 7. Group ID derivation

Use deterministic dev-local group ID bytes derived from the conversation label.

Recommended initial pattern:

    carbonstack-openmls-dev-conversation:<conversation-label>

Then:

    GroupId::from_slice(group_id_bytes)

This is not a production group ID design. It is a deterministic dev-sidecar convenience so the group can be reidentified in tests.

The summary should include a sanitized hash/reference for the group ID, not raw secret material.

## 8. Conversation creation behavior

conversation-create should:

1. parse required arguments;
2. validate device label;
3. validate conversation label;
4. require existing identity state;
5. load signer state locally without printing it;
6. recreate the local CredentialWithKey for the identity;
7. create an OpenMLS provider;
8. build MlsGroupCreateConfig with the current ciphersuite and ratchet tree extension enabled;
9. derive a dev-local GroupId from the conversation label;
10. call MlsGroup::new_with_group_id;
11. write sanitized conversation-summary.json;
12. return sanitized stdout;
13. refuse duplicate conversation creation.

## 9. Duplicate behavior

conversation-create must refuse overwrite if any of these already exist:

    conversation-summary.json

If later implementation adds explicit group/provider state files under the conversation path, those must also trigger duplicate refusal.

No --force flag in this rung.

## 10. Provider storage semantics

OpenMLS group creation is stateful. The implementation must be honest about provider storage.

If MlsGroup::new_with_group_id writes group state into the OpenMLS provider storage, return:

    provider_storage_written=true

If the implementation only creates a group object and does not persist provider storage beyond the command, document that clearly and return:

    provider_storage_written=false

Recommendation:

    Prefer explicit dev-local provider storage if needed for future MlsGroup::load, but do not design production storage in this rung.

Important:

    Dev provider storage is not secure vault storage.

## 11. Sanitized stdout fields

Success stdout may include:

    device_label
    conversation_label
    conversation_created
    state_scope
    conversation_state_path_hint
    conversation_summary_path_hint
    ciphersuite
    group_id_ref
    member_count
    epoch
    provider_storage_written
    private_material_included=false

Success stdout must not include:

    signer.json contents
    private key material
    raw group epoch secrets
    raw MlsGroup storage bytes
    provider storage JSON
    message content
    Welcome bytes
    KeyPackage bytes

## 12. conversation-summary.json fields

Recommended summary fields:

    summary_version
    conversation_label
    creator_device_label
    state_scope
    ciphersuite
    group_id_ref
    group_id_len
    member_count
    epoch
    conversation_created
    provider_storage_written
    private_material_included
    warning

The warning should say:

    dev-only OpenMLS conversation state; not production messaging or secure vault storage

## 13. Events

Add a new provider event only if the command is implemented:

    provider.conversation.created

Initial classification recommendation:

    class: conversation/setup
    severity: info
    trust_relevant: false for this dev-sidecar rung
    action: append history / debug only

Rationale:

Conversation creation is not yet user-facing group membership. Future member add/join events will become more trust-sensitive. Do not overclaim trust integration now.

Potential duplicate event:

    provider.conversation.exists

This may be added if useful. Alternatively duplicate failure may initially use checkpoint.failed, matching current duplicate public-bundle artifact behavior.

## 14. Provider-info updates

When conversation-create graduates to supported:

- add conversation-create to capabilities;
- remove conversation-create from unsupported;
- leave conversation-add-member unsupported;
- leave conversation-join unsupported;
- leave message-protect unsupported;
- leave message-open unsupported.

The existing unsupported-command Go test currently uses conversation-create. After this rung, change that test target to:

    conversation-add-member

or another still-unsupported command.

## 15. Required tests

Rust-side tests:

- conversation label accepts safe labels;
- conversation label rejects empty label;
- conversation label rejects dot and dot-dot;
- conversation label rejects path separators, dots, spaces, and unsafe characters;
- conversation label rejects overlong labels.

Go-side sidecar tests:

- provider-info lists conversation-create as supported;
- provider-info no longer lists conversation-create as unsupported;
- provider-info still lists conversation-add-member and conversation-join as unsupported;
- unsupported-command test uses conversation-add-member;
- conversation-create missing device label fails;
- conversation-create missing conversation label fails;
- conversation-create invalid device label fails;
- conversation-create invalid conversation label fails;
- conversation-create missing identity fails;
- conversation-create success after identity-create;
- success envelope reports conversation_created=true;
- success envelope reports private_material_included=false;
- success envelope reports member_count=1;
- success envelope includes group_id_ref;
- conversation-summary.json exists;
- conversation-summary.json metadata matches stdout;
- duplicate conversation-create refuses overwrite;
- stdout contains no obvious secret material.

## 16. Manual probe after future implementation

From sidecar crate:

    Remove-Item -Recurse -Force .\.carbonstack-openmls-sidecar-state -ErrorAction SilentlyContinue

    cargo run -- provider-info

    cargo run -- conversation-create --device-label carbonstack-alice-device --conversation-label carbonstack-test-conversation

Expected before identity exists:

    identity_missing or equivalent sanitized failure

Then:

    cargo run -- identity-create --device-label carbonstack-alice-device

    cargo run -- conversation-create --device-label carbonstack-alice-device --conversation-label carbonstack-test-conversation

Expected success:

    conversation_created=true
    member_count=1
    private_material_included=false

Duplicate:

    cargo run -- conversation-create --device-label carbonstack-alice-device --conversation-label carbonstack-test-conversation

Expected failure/refusal.

Invalid label:

    cargo run -- conversation-create --device-label carbonstack-alice-device --conversation-label "../bad"

Expected invalid conversation label.

Safe inspection:

    Get-ChildItem .\.carbonstack-openmls-sidecar-state\dev\conversations\carbonstack-test-conversation | Select-Object Name, Length

    Get-Content .\.carbonstack-openmls-sidecar-state\dev\conversations\carbonstack-test-conversation\conversation-summary.json

Do not inspect signer.json or raw provider/group state.

## 17. Non-goals

This rung must not:

- implement conversation-add-member;
- consume KeyPackage artifacts;
- write Welcome artifacts;
- implement conversation-join;
- implement message-protect;
- implement message-open;
- wire OpenMLS into comms send;
- wire OpenMLS into comms inbox;
- route MLS payloads through Cypher;
- mutate trust.json;
- mutate trust-events.jsonl;
- implement production secure vault storage;
- claim production E2EE;
- claim group messaging support;
- claim Signal-equivalent behavior;
- start Android or CarbonStackOS.

## 18. Success criteria

This rung succeeds when:

- docs/59 is committed before code;
- conversation-create is implemented narrowly;
- conversation-add-member/join/message commands remain unsupported;
- tests cover success and negative paths;
- generated dev state remains ignored;
- artifact guard passes;
- docs/60 result records behavior;
- repos validate and push cleanly;
- LogDoc/breakpoint v0.2.27 records the checkpoint.

## 19. Recommended next after this rung

After conversation-create validates, the next safe checkpoint should be:

    conversation-add-member / Welcome export planning

Do not jump directly to message-protect/message-open.
