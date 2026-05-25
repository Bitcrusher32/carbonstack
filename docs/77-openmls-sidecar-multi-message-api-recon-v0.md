# OpenMLS Sidecar Multi-Message API Recon v0

Status: API recon / pre-implementation result
Component: CarbonStackComms / OpenMLS sidecar
Phase: Phase 2D
Previous docs:
- docs/75-openmls-sidecar-message-protect-open-result-v0.md
- docs/76-openmls-sidecar-multi-message-continuity-plan-v0.md

## 1. Purpose

This document records targeted reconnaissance for the next OpenMLS sidecar checkpoint:

    multi-message continuity

v0.2.36 proved a one-message application-message lifecycle:

    create -> add-member -> Welcome export -> join -> protect -> open

The validated message proof was:

    Alice protects "hello bob"
    Bob opens the protected artifact
    Bob recovers "hello bob"

This recon asks whether the sidecar can safely extend that to two sequential messages across separate sidecar invocations, while preserving provider storage continuity.

This document is docs/recon only. It does not implement multi-message support.

## 2. Current implementation baseline

Current v0.2.36 message-protect writes a fixed message label:

    message-0001

Current protected artifact path:

    .carbonstack-openmls-sidecar-state/dev/conversations/<conversation-label>/messages/message-0001/application-message.bin

Current Bob open summary path:

    .carbonstack-openmls-sidecar-state/dev/devices/<device-label>/conversations/<conversation-label>/opened-messages/message-0001/message-open-summary.json

Current sidecar state layout remains asymmetric:

Alice/global conversation state:

    .carbonstack-openmls-sidecar-state/dev/conversations/<conversation-label>/

Bob/device-scoped joined state:

    .carbonstack-openmls-sidecar-state/dev/devices/<device-label>/conversations/<conversation-label>/

This asymmetry is intentionally preserved for the first multi-message proof.

## 3. OpenMLS create_message recon

OpenMLS 0.8.1 application message creation lives in:

    src/group/mls_group/application.rs

The relevant API remains:

    MlsGroup::create_message(
        &mut self,
        provider,
        signer,
        message: &[u8],
    ) -> Result<MlsMessageOut, CreateMessageError>

The implementation path observed in recon:

1. checks the group is active;
2. rejects operation if pending proposals exist;
3. creates authenticated application content;
4. encrypts the authenticated content;
5. resets AAD;
6. returns the ciphertext as an `MlsMessageOut::from_private_message(...)`.

Relevant consequence:

    message-protect is state-mutating because create_message takes &mut self and encrypts through the group's message-secret machinery.

Therefore the v0.2.36 behavior remains correct:

    save Alice provider storage after every successful message-protect

The first multi-message implementation should not weaken this behavior.

## 4. OpenMLS process_message / unprotect_message recon

OpenMLS 0.8.1 message processing lives in:

    src/group/mls_group/processing.rs

The relevant public API remains:

    MlsGroup::process_message(
        &mut self,
        provider,
        message: impl Into<ProtocolMessage>,
    ) -> Result<ProcessedMessage, ProcessMessageError<Provider::StorageError>>

Recon showed this path:

    process_message(...)
        -> unprotect_message(...)
        -> process_unverified_message(...)

Inside `unprotect_message`, OpenMLS:

1. converts the input into a `ProtocolMessage`;
2. determines whether the message is a private message;
3. decrypts the message;
4. parses the decrypted message;
5. writes message secrets back to provider storage when the secret tree was modified.

Important observed behavior:

    will_modify_secret_tree = matches!(message, ProtocolMessage::PrivateMessage(_))

and when true, OpenMLS writes the message secrets to provider storage.

Relevant consequence:

    message-open is state-mutating for private application messages.

Therefore the v0.2.36 behavior remains correct:

    save Bob provider storage after every successful message-open

The first multi-message implementation should not weaken this behavior.

## 5. Existing scratch continuity evidence

Earlier scratch work already proved two sequential Alice-to-Bob OpenMLS application messages inside one process.

That scratch proof established:

- two sequential application messages can work at OpenMLS level;
- create_message required mutable Alice group state;
- process_message required mutable Bob group state;
- outbound and inbound provider operations must be treated as state-mutating and persistence-relevant.

This v0.2.37 recon extends that lesson into the sidecar persistence path.

The new question is no longer:

    Can OpenMLS do two messages in one process?

The new question is:

    Can CarbonStack's sidecar persist Alice/Bob provider state across separate command invocations for two messages?

## 6. Message label decision

The current hardcoded `message-0001` must be replaced or parameterized before multi-message implementation.

Recommended command extension:

    message-protect --device-label <sender-device> --conversation-label <conversation> --message-label <message-label> --plaintext <text>

Keep backward compatibility:

    message-protect --device-label <sender-device> --conversation-label <conversation> --plaintext <text>

If `--message-label` is omitted, default to:

    message-0001

This preserves the existing one-message test.

Recommended message-open extension:

    message-open --device-label <receiver-device> --conversation-label <conversation> --message <path> --message-label <message-label>

If `--message-label` is omitted, default to:

    message-0001

Rationale:

- message-open already has an authoritative `--message <path>` input;
- however, Bob's open summary path currently uses the internal message label;
- adding `--message-label` keeps Bob's summary paths deterministic for tests.

## 7. Message label validation decision

Add a dedicated safe message label validator.

Suggested function:

    validate_message_label(label: &str) -> Result<(), String>

Initial allowed shape:

    ASCII alphanumeric
    hyphen
    underscore

Initial constraints:

    non-empty
    <= 64 bytes
    must not start with "."
    must not contain path separators
    must not contain whitespace
    must not be reserved internal file stem

Allowed examples:

    message-0001
    message-0002
    alice-0001
    test_msg_1

Rejected examples:

    ../message-0001
    message/0001
    message 0001
    .secret
    signer
    provider-storage
    welcome
    public-bundle
    con
    nul

Implementation can live in `state.rs` initially, but longer-term this belongs near `labels.rs`.

## 8. Artifact path decision

With explicit labels, Alice protected artifacts should write to:

    .carbonstack-openmls-sidecar-state/dev/conversations/<conversation-label>/messages/<message-label>/application-message.bin

Metadata:

    message-manifest.json
    message-protect-summary.json

Bob open summaries should write to:

    .carbonstack-openmls-sidecar-state/dev/devices/<device-label>/conversations/<conversation-label>/opened-messages/<message-label>/message-open-summary.json

Do not generate dynamic IDs in this checkpoint.

Rationale:

- explicit labels make Go tests deterministic;
- duplicate refusal is easy to test;
- later runtime integration can replace labels with message IDs or content-addressed envelope IDs.

## 9. Ordering behavior

The source recon confirms that private message processing modifies secret tree state and writes message secrets back to provider storage.

This strongly implies:

    Bob should process messages in delivery order for the first validated proof.

Do not make strong claims yet about OpenMLS out-of-order buffering or skipped-message tolerance until directly tested.

The implementation should include the happy path first:

    protect message-0001
    open message-0001
    protect message-0002
    open message-0002

Optional negative behavior test, if easy:

    protect message-0001
    protect message-0002
    try open message-0002 before message-0001

But this should be treated as exploratory and documented exactly. If it fails, preserve the error. If it succeeds, do not generalize beyond the specific OpenMLS 0.8.1 behavior observed.

## 10. Recommended v0.2.38 implementation scope

Implement only:

1. explicit `--message-label` parsing for message-protect;
2. optional explicit `--message-label` parsing for message-open;
3. message label validation;
4. replace hardcoded `message-0001` with supplied/default label;
5. Go test for the existing one-message default path;
6. Go test for explicit two-message sequential path.

Do not implement:

- autogenerated message IDs;
- out-of-order buffering;
- replay tracking beyond OpenMLS errors;
- Alice state migration;
- Cypher routing;
- Comms runtime send/inbox integration;
- trust-state mutation;
- production storage.

## 11. Proposed implementation functions / changes

In `state.rs`:

    validate_message_label(...)
    protect_dev_message_with_label(...)
    open_dev_message_with_label(...)

or modify existing functions to accept `message_label`.

Preferred minimal change:

    protect_dev_message(device_label, conversation_label, message_label, plaintext)

    open_dev_message(device_label, conversation_label, message_label, message_artifact_path)

Then route defaulting in `main.rs`:

    let message_label = parse_message_label(args).unwrap_or("message-0001");

In `main.rs`:

    parse_message_label(args)

In Go test struct:

    MessageLabel already exists from v0.2.36 envelope data.

In Go tests:

    TestOpenMLSSidecarMessageProtectOpenOneWay remains valid using default message-0001.

Add:

    TestOpenMLSSidecarMessageProtectOpenTwoSequentialMessages

## 12. Proposed two-message Go contract test

High-level test flow:

1. remove sidecar state;
2. Alice identity-create;
3. Bob identity-create;
4. Bob public-bundle-export --write-artifact;
5. Alice conversation-create;
6. Alice conversation-add-member;
7. Bob conversation-join;
8. Alice message-protect --message-label message-0001 --plaintext "hello bob 1";
9. Bob message-open --message-label message-0001 --message <message-0001 artifact>;
10. assert plaintext_utf8 == "hello bob 1";
11. Alice message-protect --message-label message-0002 --plaintext "hello bob 2";
12. Bob message-open --message-label message-0002 --message <message-0002 artifact>;
13. assert plaintext_utf8 == "hello bob 2";
14. assert both messages have distinct artifact paths;
15. assert both messages have non-empty sha256;
16. assert both message opens return matching group_id_ref;
17. assert provider_storage_loaded/written/reloadable are true for all protect/open steps;
18. assert no private material appears in stdout;
19. assert duplicate `message-0002` protect refuses overwrite.

Expected result:

    create -> add-member -> Welcome export -> join -> protect/open #1 -> protect/open #2

## 13. State persistence expectations

Expected across sequential messages:

Alice:

- message-0001 protect loads Alice provider storage;
- message-0001 protect saves Alice provider storage;
- message-0002 protect loads the updated Alice provider storage;
- message-0002 protect saves Alice provider storage again.

Bob:

- message-0001 open loads Bob joined provider storage;
- message-0001 open saves Bob joined provider storage;
- message-0002 open loads the updated Bob joined provider storage;
- message-0002 open saves Bob joined provider storage again.

This should prove process-restart-like continuity because each sidecar command is a separate invocation under the Go test harness.

## 14. Expected fields

message-protect should continue returning:

    message_label
    message_artifact_path_hint
    message_manifest_path_hint
    message_protect_summary_path_hint
    message_artifact_sha256
    message_artifact_size_bytes
    provider_storage_loaded
    provider_storage_written
    group_reloadable
    message_protected
    protected_message_written

message-open should continue returning:

    message_label
    message_artifact_path_hint
    message_open_summary_path_hint
    plaintext_utf8
    plaintext_len
    provider_storage_loaded
    provider_storage_written
    group_reloadable
    message_opened

New expected behavior:

    message_label reflects the supplied --message-label when present.

## 15. Blunders / lessons to preserve

- Broad recon over the whole OpenMLS crate is too noisy. `3A` was omitted because it produced too much output.
- High-value files for this rung are:
  - `src/group/mls_group/application.rs`
  - `src/group/mls_group/processing.rs`
  - current sidecar `state.rs`
  - current sidecar `main.rs`
  - current Go sidecar contract test.
- Source recon confirms that processing private messages writes message secrets back to provider storage; do not remove save-after-open.
- Source recon confirms `create_message` resets AAD and returns a private message carrier; keep artifact handling as `MlsMessageOut::to_bytes()`.
- Prior scratch work already proved two-message OpenMLS feasibility in-process, but sidecar persistence is still the thing to validate.
- Keep the `[recon + doc] -> [implement + patch]` workflow.

## 16. Final recommendation

Proceed to v0.2.38 implementation.

Implement explicit message labels and two sequential Alice-to-Bob messages.

Do not migrate Alice state yet.

Do not route through Cypher yet.

Do not wire Comms runtime yet.

The next validated claim should be:

    The OpenMLS sidecar can protect/open two sequential application messages across separate sidecar invocations while preserving Alice and Bob provider state.
