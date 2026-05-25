# OpenMLS Sidecar Multi-Message Continuity Result v0

Status: Validated implementation checkpoint
Component: CarbonStackComms / OpenMLS sidecar
Phase: Phase 2D
Previous docs:
- docs/75-openmls-sidecar-message-protect-open-result-v0.md
- docs/76-openmls-sidecar-multi-message-continuity-plan-v0.md
- docs/77-openmls-sidecar-multi-message-api-recon-v0.md

## 1. Summary

This checkpoint extends the OpenMLS sidecar message-protect/message-open proof from one fixed message artifact to explicit message labels and two sequential application messages.

v0.2.36 validated:

    create -> add-member -> Welcome export -> join -> protect -> open

with one hardcoded message label:

    message-0001

This checkpoint validates:

    create -> add-member -> Welcome export -> join -> protect/open message-0001 -> protect/open message-0002

The validated plaintexts were:

    message-0001: "hello bob 1"
    message-0002: "hello bob 2"

Both messages were protected by Alice and opened by Bob across separate sidecar invocations.

## 2. New behavior

message-protect now accepts an explicit message label:

    message-protect --device-label <sender-device> --conversation-label <conversation> --message-label <message-label> --plaintext <text>

message-open now accepts an explicit message label:

    message-open --device-label <receiver-device> --conversation-label <conversation> --message-label <message-label> --message <path>

If `--message-label` is omitted, both commands default to:

    message-0001

This preserves v0.2.36 compatibility.

## 3. Message label validation

A new message label validator rejects unsafe labels.

Allowed shape:

    ASCII letters
    ASCII numbers
    hyphen
    underscore

Rejected:

    empty labels
    labels longer than 64 bytes
    labels starting with dot
    labels containing path separators
    labels containing characters outside the allowed ASCII set
    reserved/internal label names such as signer, provider-storage, welcome, application-message, con, nul, etc.

The validator currently lives in the sidecar state layer.

## 4. Artifact layout

Alice protected artifacts now use the explicit message label:

    .carbonstack-openmls-sidecar-state/dev/conversations/<conversation-label>/messages/<message-label>/application-message.bin

Associated metadata:

    message-manifest.json
    message-protect-summary.json

Bob open summaries now use the explicit message label:

    .carbonstack-openmls-sidecar-state/dev/devices/<device-label>/conversations/<conversation-label>/opened-messages/<message-label>/message-open-summary.json

## 5. Validated manual flow

Manual flow validated:

1. reset sidecar state;
2. Alice identity-create;
3. Bob identity-create;
4. Bob public-bundle-export --write-artifact;
5. Alice conversation-create;
6. Alice conversation-add-member using Bob KeyPackage;
7. Bob conversation-join using Alice Welcome;
8. Alice message-protect --message-label message-0001 --plaintext "hello bob 1";
9. Bob message-open --message-label message-0001;
10. Bob recovered plaintext_utf8="hello bob 1";
11. Alice message-protect --message-label message-0002 --plaintext "hello bob 2";
12. Bob message-open --message-label message-0002;
13. Bob recovered plaintext_utf8="hello bob 2".

The final corrected message-0002 open output reported:

    message_label=message-0002
    plaintext_utf8="hello bob 2"
    plaintext_len=11
    message_opened=true
    provider_storage_loaded=true
    provider_storage_written=true
    group_reloadable=true

## 6. Go contract test

Go contract coverage added:

    TestOpenMLSSidecarMessageProtectOpenTwoSequentialMessages

The test validates:

- Alice identity-create;
- Bob identity-create;
- Bob public-bundle-export --write-artifact;
- Alice conversation-create;
- Alice conversation-add-member;
- Bob conversation-join;
- message-0001 protect/open;
- message-0002 protect/open;
- message-0001 plaintext is "hello bob 1";
- message-0002 plaintext is "hello bob 2";
- message-0001 and message-0002 artifact paths differ;
- message artifact hashes differ;
- provider_storage_loaded=true;
- provider_storage_written=true;
- group_reloadable=true;
- member_count=2;
- group_id_ref remains consistent;
- duplicate message-0002 protect refuses overwrite with message_artifact_exists;
- stdout does not include forbidden secret material.

The existing one-message test remains valid because the default message label remains message-0001.

## 7. Implementation files

Changed in carbonstack-comms:

    internal/protocol/mls/research/openmls-sidecar/src/state.rs
    internal/protocol/mls/research/openmls-sidecar/src/main.rs
    internal/protocol/openmls_sidecar_provider_info_test.go

state.rs updates:

    validate_message_label(...)
    protect_dev_message(..., message_label, plaintext)
    open_dev_message(..., message_label, message_artifact_path)

main.rs updates:

    parse_message_label(...)
    message-protect validates and passes message_label
    message-open validates and passes message_label
    both commands default to message-0001 when omitted

Go test updates:

    MessageLabel envelope data field
    TestOpenMLSSidecarMessageProtectOpenTwoSequentialMessages
    assertMessageProtectSuccess(...)
    assertMessageOpenSuccess(...)

## 8. Blunders and lessons

### 8.1 Function signature mismatch

After changing protect_dev_message to require message_label, main.rs still called the old three-argument form.

Fix:

    parse --message-label
    default to message-0001
    pass message_label into protect_dev_message

Lesson:

When state function signatures change, patch CLI handlers immediately.

### 8.2 Missing import

main.rs used validate_message_label before importing it from state.rs.

Fix:

    import validate_message_label

Lesson:

The compiler suggestion was correct; apply direct import fixes before more structural patching.

### 8.3 Stale message-open internal default

open_dev_message accepted message_label but then shadowed it internally:

    let message_label = "message-0001";

This caused message-open to decrypt message-0002 correctly but write/report message-0001.

Fix:

    remove the stale local shadowing line

Lesson:

Manual output matters. Correct plaintext alone was not enough; metadata and summary paths also had to match.

### 8.4 Go helper type typo

New Go helper signatures used sidecarEnvelope, but the real type is openMLSSidecarEnvelope.

Fix:

    replace sidecarEnvelope with openMLSSidecarEnvelope

Lesson:

Reuse existing local type names from parseSidecarEnvelope/assertProviderEnvelopeBase.

### 8.5 Duplicate Go struct fields

The first envelope data patch duplicated message fields that already existed from v0.2.36.

Fix:

    keep the first occurrence and remove duplicate field lines

Lesson:

Before adding fields to the large test envelope struct, inspect the top of file. Most v0.2.36 message fields already existed; only MessageLabel was missing.

## 9. Allowed claims

Allowed:

- message-protect supports explicit safe message labels.
- message-open supports explicit safe message labels.
- message-protect/message-open preserve backward-compatible default message-0001.
- Alice can protect message-0001 and message-0002 across separate sidecar invocations.
- Bob can open message-0001 and message-0002 across separate sidecar invocations.
- Bob recovered "hello bob 1" and "hello bob 2".
- The sidecar now validates two sequential Alice-to-Bob application messages.
- Duplicate message labels refuse overwrite.

Not allowed:

- out-of-order delivery behavior is proven.
- replay behavior is proven.
- bidirectional messaging is proven.
- generated message IDs exist.
- Alice state has been migrated to device-scoped layout.
- Cypher routes MLS payloads.
- Comms runtime send/inbox uses OpenMLS.
- trust-state consumes protect/open events.
- production E2EE exists.

## 10. Next recommended checkpoint

Next safest checkpoint should be docs/recon, not immediate runtime integration.

Recommended options:

1. Out-of-order / replay behavior recon:
   - Bob tries to open message-0002 before message-0001.
   - Bob tries to open message-0001 twice.
   - Record exact OpenMLS behavior.

2. Alice state layout cleanup plan:
   - Migrate Alice conversation state from global path to device-scoped path.
   - Avoid doing this casually because many commands/tests depend on the current global Alice path.

3. Cypher routing design:
   - Only after ordering/replay and state layout are understood.
   - Route Welcome and application-message artifacts as opaque payloads.

Recommendation:

    Do out-of-order/replay recon first, then state layout cleanup, then Cypher routing.
