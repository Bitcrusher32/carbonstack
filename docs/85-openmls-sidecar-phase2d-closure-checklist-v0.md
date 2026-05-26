# OpenMLS Sidecar Phase 2D Closure Checklist v0

Status: Phase 2D polish / closure planning
Component: CarbonStackComms / OpenMLS sidecar
Phase: Phase 2D
Previous docs:
- docs/78-openmls-sidecar-multi-message-continuity-result-v0.md
- docs/80-openmls-sidecar-message-ordering-replay-api-recon-v0.md
- docs/84-openmls-sidecar-alice-device-scoped-state-result-v0.md

## 1. Purpose

This document defines what remains before Phase 2D OpenMLS sidecar research can be considered polished enough to pause and move into Cypher MLS artifact relay research.

v0.2.40 moved Alice creator conversation state to the same device-scoped path pattern as Bob joined/open state.

The follow-up stale-path sweep found one remaining operational old-path bug: `protect_dev_message` was reading Alice device-scoped provider storage but still writing protected message artifacts through the old global conversation message helpers.

That bug was fixed in `carbonstack-comms` before this closure checklist proceeded.

## 2. Current Phase 2D validated baseline

The sidecar currently validates:

- provider-info;
- identity-create;
- identity-status;
- public-bundle-export;
- public-bundle-export --write-artifact;
- conversation-create;
- conversation-load-check;
- conversation-add-member;
- conversation-join;
- message-protect;
- message-open;
- explicit message labels;
- two-message continuity;
- same-sender two-message out-of-order open behavior;
- duplicate/replay rejection;
- corrupt/truncated artifact rejection;
- device-scoped creator and joined conversation state.

The validated lifecycle is:

    Alice identity-create
    Bob identity-create
    Bob public-bundle-export --write-artifact
    Alice conversation-create
    Alice conversation-add-member
    Bob conversation-join
    Alice message-protect message-0001
    Alice message-protect message-0002
    Bob message-open message-0001/message-0002
    Bob out-of-order open test
    Bob duplicate/replay open rejection test
    Bob corrupt artifact rejection test

## 3. Current canonical device-scoped layout

Identity/public bundle:

    .carbonstack-openmls-sidecar-state/dev/devices/<device-label>/

Creator/joined conversation state:

    .carbonstack-openmls-sidecar-state/dev/devices/<device-label>/conversations/<conversation-label>/

Alice Welcome artifact:

    .carbonstack-openmls-sidecar-state/dev/devices/<alice-device>/conversations/<conversation-label>/welcome.bin

Alice protected message artifact:

    .carbonstack-openmls-sidecar-state/dev/devices/<alice-device>/conversations/<conversation-label>/messages/<message-label>/application-message.bin

Bob opened-message summary:

    .carbonstack-openmls-sidecar-state/dev/devices/<bob-device>/conversations/<conversation-label>/opened-messages/<message-label>/message-open-summary.json

The old Alice-global path is historical only:

    .carbonstack-openmls-sidecar-state/dev/conversations/<conversation-label>/

It should not be used by current implementation paths.

## 4. Stale path sweep result

The stale path sweep checked:

    internal/protocol/openmls_sidecar_provider_info_test.go
    internal/protocol/mls/research/openmls-sidecar/src/state.rs
    internal/protocol/mls/research/openmls-sidecar/src/main.rs

Results:

- Go sidecar tests had no hardcoded old `dev/conversations` assumptions.
- `main.rs` had no hardcoded old `dev/conversations` assumptions; it prints path hints from result structs.
- `state.rs` still had old global helper definitions.
- Most operational call sites were already device-scoped.
- One operational bug remained in `protect_dev_message`.

The stale block was:

    conversation_message_dir(conversation_label, message_label)
    conversation_message_artifact_path(conversation_label, message_label)
    conversation_message_manifest_path(conversation_label, message_label)
    conversation_message_protect_summary_path(conversation_label, message_label)

It was replaced with:

    device_conversation_message_dir(device_label, conversation_label, message_label)
    device_conversation_message_artifact_path(device_label, conversation_label, message_label)
    device_conversation_message_manifest_path(device_label, conversation_label, message_label)
    device_conversation_message_protect_summary_path(device_label, conversation_label, message_label)

This completes the intended v0.2.40 device-scoped artifact layout.

## 5. Phase 2D closure checklist

### Required before Cypher routing research

- [x] Stale path sweep:
  - active Go tests no longer hardcode old Alice global `dev/conversations/<conversation-label>/` paths;
  - active sidecar operational call sites no longer use old global message artifact helpers;
  - `message-protect` now writes protected message artifacts under the sender device-scoped conversation path.

- [x] Wrong-device negative behavior:
  - opening a message with a device that has identity state but lacks joined conversation state fails cleanly before decrypt/process;
  - validated by `TestOpenMLSSidecarMessageOpenWrongDeviceRejected`;
  - result:
    - `ok=false`;
    - `code=conversation_or_message_missing`;
    - `message="device conversation provider storage is missing"`;
    - `provider_event=provider.conversation.missing`;
    - `severity=warning`;
    - `trust_relevant=false`;
    - exit code `3`.

- [x] Wrong-conversation negative behavior:
  - opening a valid artifact under a receiver conversation label with no joined provider storage fails cleanly before decrypt/process;
  - validated by `TestOpenMLSSidecarMessageOpenWrongConversationRejected`;
  - result:
    - `ok=false`;
    - `code=conversation_or_message_missing`;
    - `message="device conversation provider storage is missing"`;
    - `provider_event=provider.conversation.missing`;
    - `severity=warning`;
    - `trust_relevant=false`;
    - exit code `3`.

- [ ] Bidirectional proof:
  - after Alice adds Bob and Bob joins, Bob should protect a message and Alice should open it;
  - if current sidecar cannot support this yet, document why and what API/state change is needed.

- [ ] Phase 2D closure result doc:
  - summarize validated sidecar behavior;
  - list not-validated boundaries;
  - mark the sidecar research as ready for Cypher artifact relay design.

### Nice-to-have before Cypher, but not mandatory

- [ ] Split the giant Go sidecar test file into focused test files.
- [ ] Remove unused old global helper definitions if they are no longer called.
- [ ] Improve replay/secret-reuse event taxonomy.
- [ ] Add invalid message-label tests if missing.
- [ ] Add missing-artifact test coverage if missing.

## 6. Recommended next implementation/test rungs

### Completed during v0.2.41 polish: wrong-device / wrong-conversation negative tests

Purpose:

    Validate that Cypher routing metadata mistakes fail safely.

Validated tests:

    TestOpenMLSSidecarMessageOpenWrongDeviceRejected
    TestOpenMLSSidecarMessageOpenWrongConversationRejected

Both cases currently fail as missing receiver device/conversation provider storage before OpenMLS message processing.

### Next candidate: bidirectional message proof

Purpose:

    Validate that Bob can send back to Alice after joining.

Suggested test:

    TestOpenMLSSidecarMessageProtectOpenBidirectional

Possible question:

    Does Alice need a joined/creator receiver distinction, or does creator state after add-member already support opening Bob's private application message?

Do not assume. Probe and test.

### v0.2.44 candidate: Phase 2D closure result

Purpose:

    Declare sidecar research complete enough to support Cypher artifact relay research.

This is not productionization.

## 7. Boundary before Cypher

Cypher research may start after the required closure checklist is addressed or explicitly deferred.

Cypher v0 should be a minimal opaque artifact relay, not a full messaging product.

Cypher should route:

- Welcome artifacts;
- application-message artifacts;
- metadata needed for delivery;
- hashes/sizes;
- sender/recipient/conversation labels or future IDs;
- ordering/retry hints.

Cypher must not learn:

- plaintext;
- provider storage;
- signer material;
- private keys;
- group secrets;
- raw MemoryStorage;
- trust-state private material.

## 8. Future maturity rung after Cypher research

After Cypher routing research agrees with Comms/OpenMLS research, add a dedicated maturity and cleanup phase.

Working name:

    Phase 2E: Research-to-Implementation Promotion

This phase should move researched systems out of research directories into official repo implementation structures.

This is not a user release.

Expected work:

- define official internal package boundaries;
- move sidecar code out of `mls/research/openmls-sidecar`;
- define stable internal protocol interfaces;
- define artifact schemas;
- define official test harnesses;
- keep dev-only constraints explicit;
- decide what remains research-only;
- avoid Android/Pixel 4a work until a barebones Android dev app exists much later.

