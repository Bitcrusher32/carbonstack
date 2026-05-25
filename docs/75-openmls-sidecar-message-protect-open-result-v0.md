# OpenMLS Sidecar Message Protect/Open Result v0

Status: Validated implementation checkpoint
Component: CarbonStackComms / OpenMLS sidecar
Phase: Phase 2D
Previous docs:
- docs/72-openmls-sidecar-conversation-join-result-v0.md
- docs/73-openmls-sidecar-message-protect-open-skeleton-v0.md
- docs/74-openmls-sidecar-message-protect-open-api-recon-v0.md

## 1. Summary

This checkpoint implements the first dev-local OpenMLS application message protect/open path.

The new sidecar commands are:

    message-protect --device-label <sender-device> --conversation-label <conversation> --plaintext <text>

    message-open --device-label <receiver-device> --conversation-label <conversation> --message <path>

The validated flow is:

1. Alice identity exists.
2. Bob identity exists.
3. Bob exports a public KeyPackage artifact and local private provider state.
4. Alice creates a conversation.
5. Alice adds Bob's KeyPackage.
6. Alice exports a Welcome carrier artifact.
7. Bob consumes the Welcome and joins.
8. Alice protects plaintext "hello bob".
9. Alice writes a protected application-message artifact.
10. Bob opens the protected application-message artifact.
11. Bob recovers plaintext "hello bob".
12. Both sides save provider storage after message operations.
13. Both sides prove group reloadability.

This extends the validated dev-local MLS lifecycle to:

    create -> add-member -> Welcome export -> join -> protect -> open

This checkpoint does not implement Comms runtime integration, Cypher routing, trust-state mutation, production storage, or production E2EE.

## 2. Implemented commands

### message-protect

    message-protect --device-label <sender-device> --conversation-label <conversation> --plaintext <text>

Validated example:

    cargo run -- message-protect --device-label carbonstack-alice-device --conversation-label carbonstack-test-conversation --plaintext "hello bob"

### message-open

    message-open --device-label <receiver-device> --conversation-label <conversation> --message <path>

Validated example:

    cargo run -- message-open --device-label carbonstack-bob-device --conversation-label carbonstack-test-conversation --message .\.carbonstack-openmls-sidecar-state\dev\conversations\carbonstack-test-conversation\messages\message-0001\application-message.bin

## 3. Implementation files

CarbonStackComms changed:

    internal/protocol/mls/research/openmls-sidecar/src/state.rs
    internal/protocol/mls/research/openmls-sidecar/src/main.rs
    internal/protocol/openmls_sidecar_provider_info_test.go

Key state.rs additions:

    conversation_messages_dir(...)
    conversation_message_dir(...)
    conversation_message_artifact_path(...)
    conversation_message_manifest_path(...)
    conversation_message_protect_summary_path(...)
    device_conversation_message_open_summary_path(...)
    MessageProtectResult
    MessageOpenResult
    MessageManifest
    MessageProtectSummary
    MessageOpenSummary
    validate_plaintext_for_dev(...)
    validate_message_artifact_path(...)
    protect_dev_message(...)
    open_dev_message(...)

Key main.rs additions:

    PHASE_MESSAGE_PROTECT
    PHASE_MESSAGE_OPEN
    parse_plaintext(...)
    parse_message_artifact_path(...)
    handle_message_protect(...)
    handle_message_open(...)
    print_message_protect_success(...)
    print_message_open_success(...)
    shared message command failure printers

Provider-info was updated so message-protect and message-open are capabilities. Unsupported commands are now state-checkpoint and state-load-check.

## 4. Protect API path

The implemented protect path uses:

    MlsGroup::create_message(&provider, &signer, plaintext.as_bytes())

This returns:

    MlsMessageOut

The artifact is serialized with:

    MlsMessageOut::to_bytes()

The protected bytes are written as:

    application-message.bin

Initial v0 message path:

    .carbonstack-openmls-sidecar-state/dev/conversations/<conversation-label>/messages/message-0001/application-message.bin

Alice state source remains the v0.2.34 global conversation path:

    .carbonstack-openmls-sidecar-state/dev/conversations/<conversation-label>/provider-storage.json

This preserves the existing validated flow and intentionally defers Alice device-scoped state migration.

## 5. Open API path

The implemented open path reads the protected artifact bytes and uses:

    MlsMessageIn::tls_deserialize_exact_bytes(...)
    try_into_protocol_message()
    MlsGroup::process_message(&provider, protocol_message)
    processed_message.into_content()
    ProcessedMessageContent::ApplicationMessage(application_message)
    application_message.into_bytes()

The plaintext is decoded as UTF-8 for the dev proof and returned as bounded stdout field:

    plaintext_utf8

Bob state source is the device-scoped joined conversation path:

    .carbonstack-openmls-sidecar-state/dev/devices/<device-label>/conversations/<conversation-label>/provider-storage.json

## 6. State mutation and persistence

Both message-protect and message-open are treated as state-mutating.

Reason:

- create_message takes &mut self.
- process_message takes &mut self.
- future multi-message continuity requires persistence after each side's operation.

message-protect does:

1. load Alice provider storage;
2. load Alice group;
3. create application message;
4. write application-message.bin;
5. save Alice provider storage;
6. reload Alice group from saved storage.

message-open does:

1. load Bob joined provider storage;
2. load Bob joined group;
3. deserialize/process application message;
4. recover plaintext bytes;
5. save Bob provider storage;
6. reload Bob group from saved storage.

Both report:

    provider_storage_loaded=true
    provider_storage_written=true
    group_reloadable=true

## 7. State layout

This checkpoint intentionally keeps the v0.2.34 asymmetry:

Alice/global path:

    .carbonstack-openmls-sidecar-state/dev/conversations/<conversation-label>/

Bob/device-scoped joined path:

    .carbonstack-openmls-sidecar-state/dev/devices/<device-label>/conversations/<conversation-label>/

Alice message artifact path:

    .carbonstack-openmls-sidecar-state/dev/conversations/<conversation-label>/messages/message-0001/

Files:

    application-message.bin
    message-manifest.json
    message-protect-summary.json

Bob open summary path:

    .carbonstack-openmls-sidecar-state/dev/devices/<device-label>/conversations/<conversation-label>/opened-messages/message-0001/message-open-summary.json

This is acceptable for the first message proof, but future cleanup should consider migrating Alice conversation-create to device-scoped state.

## 8. Success envelopes

### message-protect success

Validated fields include:

    ok=true
    command=message-protect
    phase=phase2d-message-protect-dev
    private_material_included=false
    message_protected=true
    protected_message_written=true
    provider_storage_loaded=true
    provider_storage_written=true
    group_reloadable=true
    member_count=2
    epoch_before=GroupEpoch(1)
    epoch_after=GroupEpoch(1)
    group_id_ref=<same as add-member/join>
    message_artifact_path_hint=<...application-message.bin>
    message_artifact_sha256=sha256:<hex>
    message_artifact_size_bytes=205

Events:

    message.protected
    storage.saved

### message-open success

Validated fields include:

    ok=true
    command=message-open
    phase=phase2d-message-open-dev
    private_material_included=false
    message_opened=true
    plaintext_utf8="hello bob"
    plaintext_len=9
    provider_storage_loaded=true
    provider_storage_written=true
    group_reloadable=true
    member_count=2
    epoch_before=GroupEpoch(1)
    epoch_after=GroupEpoch(1)
    group_id_ref=<same as add-member/join/protect>
    message_open_summary_path_hint=<...message-open-summary.json>

Events:

    message.opened
    storage.saved

## 9. Manual validation

Validated manual flow:

    Remove-Item -Recurse -Force .\.carbonstack-openmls-sidecar-state -ErrorAction SilentlyContinue

    cargo run -- identity-create --device-label carbonstack-alice-device

    cargo run -- identity-create --device-label carbonstack-bob-device

    cargo run -- public-bundle-export --device-label carbonstack-bob-device --write-artifact

    cargo run -- conversation-create --device-label carbonstack-alice-device --conversation-label carbonstack-test-conversation

    cargo run -- conversation-add-member --device-label carbonstack-alice-device --conversation-label carbonstack-test-conversation --member-keypackage .\.carbonstack-openmls-sidecar-state\dev\devices\carbonstack-bob-device\public-bundle.keypackage.bin

    cargo run -- conversation-join --device-label carbonstack-bob-device --conversation-label carbonstack-test-conversation --welcome .\.carbonstack-openmls-sidecar-state\dev\conversations\carbonstack-test-conversation\welcome.bin

    cargo run -- message-protect --device-label carbonstack-alice-device --conversation-label carbonstack-test-conversation --plaintext "hello bob"

    cargo run -- message-open --device-label carbonstack-bob-device --conversation-label carbonstack-test-conversation --message .\.carbonstack-openmls-sidecar-state\dev\conversations\carbonstack-test-conversation\messages\message-0001\application-message.bin

Manual proof observed:

- message-protect returned ok=true.
- message-protect returned message_protected=true.
- message-protect returned protected_message_written=true.
- message-protect returned provider_storage_loaded=true.
- message-protect returned provider_storage_written=true.
- message-protect returned group_reloadable=true.
- message-protect wrote application-message.bin.
- message-open returned ok=true.
- message-open returned message_opened=true.
- message-open returned plaintext_utf8="hello bob".
- message-open returned plaintext_len=9.
- message-open returned provider_storage_loaded=true.
- message-open returned provider_storage_written=true.
- message-open returned group_reloadable=true.
- message-open wrote message-open-summary.json.

Safe file listing confirmed Alice message path:

    application-message.bin
    message-manifest.json
    message-protect-summary.json

Safe file listing confirmed Bob opened-message path:

    message-open-summary.json

## 10. Go contract tests

Go test coverage was updated for:

    TestOpenMLSSidecarProviderInfoCommand
    TestOpenMLSSidecarUnsupportedCommandEnvelope
    TestOpenMLSSidecarMessageProtectOpenOneWay

Provider-info changes:

- message-protect is now expected in capabilities.
- message-open is now expected in capabilities.
- message-protect/message-open are no longer expected in unsupported.
- unsupported-command test moved to state-checkpoint.

One-way test validates:

1. Alice identity-create;
2. Bob identity-create;
3. Bob public-bundle-export --write-artifact;
4. Bob provider_storage_written=true after public-bundle export;
5. Alice conversation-create;
6. Alice conversation-add-member;
7. Bob conversation-join;
8. Alice message-protect --plaintext "hello bob";
9. message-protect returns expected fields;
10. protected message artifact files exist;
11. Bob message-open consumes the artifact;
12. message-open returns plaintext_utf8="hello bob";
13. message-open returns plaintext_len=9;
14. message-open reports provider_storage_loaded/written/reloadable;
15. duplicate message-protect refuses overwrite with message_artifact_exists;
16. stdout does not contain forbidden secret material.

## 11. Validation commands

Validated during implementation:

    cargo check

    cargo run -- provider-info

    cargo run -- message-protect --device-label carbonstack-alice-device --conversation-label carbonstack-test-conversation --plaintext "hello bob"

    cargo run -- message-open --device-label carbonstack-bob-device --conversation-label carbonstack-test-conversation --message .\.carbonstack-openmls-sidecar-state\dev\conversations\carbonstack-test-conversation\messages\message-0001\application-message.bin

    go test ./internal/protocol -run "TestOpenMLSSidecarProviderInfoCommand|TestOpenMLSSidecarUnsupportedCommandEnvelope"

    go test ./internal/protocol -run "TestOpenMLSSidecarProviderInfoCommand|TestOpenMLSSidecarUnsupportedCommandEnvelope|TestOpenMLSSidecarMessageProtectOpenOneWay"

Before final commit, also run:

    go test ./internal/protocol
    go test ./...
    powershell -ExecutionPolicy Bypass -File .\scripts\check-no-rust-artifacts.ps1

## 12. Blunders and lessons from this rung

### 12.1 Provider-info malformed JSON

The first provider-info patch inserted message-protect/message-open into capabilities but left a missing comma after conversation-join.

Failure:

    invalid character '"' after array element

Repair:

    add comma after "conversation-join"
    clean blank unsupported list area

Lesson:

provider-info raw JSON string construction remains fragile. It should eventually be rewritten using serde_json::json! or structured arrays.

### 12.2 Unsupported-command test had to move again

Once message-protect became real, the unsupported-command test could no longer use message-protect.

Failure:

    message-protect returned missing_required_argument

Repair:

    unsupported-command test moved to state-checkpoint.

Lesson:

Every time a command moves from unsupported to supported, update:

- provider-info capabilities;
- provider-info unsupported list;
- Go provider-info expectations;
- unsupported-command test target.

### 12.3 State layout remains asymmetric

Alice still uses:

    dev/conversations/<conversation-label>/

Bob uses:

    dev/devices/<device-label>/conversations/<conversation-label>/

This was intentionally preserved to minimize migration risk.

Lesson:

The message proof is real, but future work should decide whether to migrate Alice to device-scoped state before multi-message, multi-device, or runtime integration.

### 12.4 Message operation persistence is necessary

Both create_message and process_message require mutable group state.

Lesson:

Treat message-protect/message-open as persistence-relevant operations. Saving provider storage after both operations is the correct conservative v0 behavior.

### 12.5 Plaintext stdout is test-only

message-open returns plaintext_utf8 for the dev proof.

Lesson:

This is acceptable for sidecar validation but must not be framed as production UX or final Comms message handling.

## 13. Allowed claims after this checkpoint

Allowed:

- The sidecar can create a dev-local OpenMLS group.
- Alice can add Bob using Bob's KeyPackage.
- Alice can export a Welcome.
- Bob can consume the Welcome and join.
- Alice can protect a plaintext application message.
- Bob can open the protected application message.
- Bob recovered plaintext "hello bob" from Alice's protected MLS artifact.
- Both message-protect and message-open save provider storage.
- Both message-protect and message-open reload-prove group state.
- The dev-local MLS lifecycle now exists through:
  create -> add-member -> Welcome export -> join -> protect -> open.

Not allowed:

- Comms CLI uses this flow.
- Cypher routes MLS application payloads.
- trust-state mutates from message events.
- production E2EE exists.
- storage is production secure.
- metadata privacy/replay resistance/hostile-server security are proven.
- Android or CarbonStackOS integration exists.
- multi-message continuity is proven.

## 14. Next recommended checkpoint

Next docs/recon checkpoint should decide between:

Option A:

    multi-message continuity proof

Option B:

    Alice device-scoped state migration

Option C:

    sidecar command cleanup / provider-info structured JSON migration

Recommended next safest checkpoint:

    docs/recon for multi-message continuity + state layout cleanup

Do not jump to Cypher/Comms runtime integration yet.
