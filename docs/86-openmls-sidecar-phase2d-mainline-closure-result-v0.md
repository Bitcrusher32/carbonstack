# OpenMLS Sidecar Phase 2D Mainline Closure Result v0

Status: Mainline research closure checkpoint
Component: CarbonStackComms / OpenMLS sidecar
Phase: Phase 2D
Previous docs:
- docs/84-openmls-sidecar-alice-device-scoped-state-result-v0.md
- docs/85-openmls-sidecar-phase2d-closure-checklist-v0.md

## 1. Summary

This document records the mainline closure point for Phase 2D OpenMLS sidecar research.

The sidecar now validates the full dev-local Alice/Bob OpenMLS lifecycle needed before Cypher artifact relay research:

    identity-create
    public-bundle-export
    KeyPackage artifact export
    conversation-create
    conversation-add-member
    Welcome export
    conversation-join
    message-protect
    message-open
    explicit message labels
    two-message continuity
    out-of-order same-sender two-message open
    duplicate/replay rejection
    corrupt/truncated artifact rejection
    wrong-device / wrong-conversation rejection
    bidirectional Alice <-> Bob message flow
    device-scoped creator and joined conversation state

This is not a production E2EE claim. It is a research-sidecar closure point.

## 2. Final Phase 2D closure additions

The final closure polish added:

- stale-path sweep after Alice device-scoped migration;
- fix for remaining `protect_dev_message` old global artifact path usage;
- wrong-device message-open negative behavior;
- wrong-conversation message-open negative behavior;
- bidirectional Bob-to-Alice message proof.

## 3. Device-scoped state model

Current canonical layout:

    .carbonstack-openmls-sidecar-state/dev/devices/<device-label>/conversations/<conversation-label>/

Alice creator state:

    dev/devices/carbonstack-alice-device/conversations/carbonstack-test-conversation/

Bob joined state:

    dev/devices/carbonstack-bob-device/conversations/carbonstack-test-conversation/

Alice protected message artifact:

    dev/devices/carbonstack-alice-device/conversations/carbonstack-test-conversation/messages/<message-label>/application-message.bin

Bob protected message artifact:

    dev/devices/carbonstack-bob-device/conversations/carbonstack-test-conversation/messages/<message-label>/application-message.bin

Opened-message summaries:

    dev/devices/<receiver-device>/conversations/<conversation-label>/opened-messages/<message-label>/message-open-summary.json

## 4. Wrong-target behavior

Wrong-device open:

    Eve has identity state but no joined conversation state.
    Eve attempts to open Alice's valid artifact under carbonstack-test-conversation.

Result:

    ok=false
    code=conversation_or_message_missing
    message="device conversation provider storage is missing"
    provider_event=provider.conversation.missing
    severity=warning
    trust_relevant=false
    exit code=3

Wrong-conversation open:

    Bob has joined carbonstack-test-conversation.
    Bob attempts to open Alice's valid artifact under carbonstack-wrong-conversation.

Result:

    ok=false
    code=conversation_or_message_missing
    message="device conversation provider storage is missing"
    provider_event=provider.conversation.missing
    severity=warning
    trust_relevant=false
    exit code=3

Interpretation:

    message-open fails before decrypt/process if the target receiver device/conversation provider storage does not exist.

This is useful for Cypher routing research because wrong routing metadata currently fails safely and does not expose plaintext or private material.

## 5. Bidirectional proof

Bidirectional proof validates:

    Alice protects alice-message-0001.
    Bob opens alice-message-0001.
    Bob protects bob-message-0001.
    Alice opens bob-message-0001.

Expected plaintexts:

    "hello bob from alice"
    "hello alice from bob"

Interpretation:

    Alice's creator state after add-member/merge can process Bob's private application message.
    Bob's joined state can protect an application message back to Alice.
    Both device-scoped sides can now send and receive in the dev-local two-member group.

## 6. Go tests at closure

The key OpenMLS sidecar tests include:

    TestOpenMLSSidecarConversationCreate
    TestOpenMLSSidecarConversationLoadCheck
    TestOpenMLSSidecarConversationAddMemberWelcomeExport
    TestOpenMLSSidecarConversationJoinWelcomeConsume
    TestOpenMLSSidecarMessageProtectOpenOneWay
    TestOpenMLSSidecarMessageProtectOpenTwoSequentialMessages
    TestOpenMLSSidecarMessageOpenOutOfOrderTwoMessageDelivery
    TestOpenMLSSidecarMessageOpenDuplicateRejected
    TestOpenMLSSidecarMessageOpenCorruptArtifactRejected
    TestOpenMLSSidecarMessageOpenWrongDeviceRejected
    TestOpenMLSSidecarMessageOpenWrongConversationRejected
    TestOpenMLSSidecarMessageProtectOpenBidirectional

## 7. Validation

Validation completed for the Phase 2D closure path should include:

    go test -p 1 ./internal/protocol
    go test -p 1 ./...
    powershell -ExecutionPolicy Bypass -File .\scripts\check-no-rust-artifacts.ps1

Low parallelism is recommended temporarily after the local Windows BSOD during heavy validation. That BSOD was a local machine stability event, not a CarbonStack test failure.

## 8. Allowed claims

Allowed:

- Phase 2D OpenMLS sidecar mainline research is complete enough to support Cypher artifact relay research.
- The sidecar validates a full dev-local two-device Alice/Bob OpenMLS lifecycle.
- Device-scoped creator and joined conversation state are validated.
- Both Alice and Bob can protect/open messages in the same two-member group.
- Wrong receiver device/conversation metadata fails safely before message processing.
- Duplicate/replay open is rejected by OpenMLS secret reuse behavior.
- Corrupt/truncated artifacts are rejected before message processing.

Not allowed:

- production E2EE exists;
- Signal-equivalent security exists;
- Cypher routes MLS artifacts;
- Comms runtime uses OpenMLS;
- trust-state is mutated from sidecar events;
- production vault/secure storage exists;
- Android/Pixel 4a target is involved;
- sidecar code has been promoted out of research;
- generated message IDs exist;
- multi-member/multi-sender stress behavior is complete;
- long skipped-message windows are validated.

## 9. Known remaining polish

Phase 2D can be revisited later for polish, including:

- remove dead old global helper definitions;
- split the large Go sidecar test file;
- improve replay/secret-reuse event taxonomy;
- add missing-artifact explicit tests if desired;
- add invalid message-label negative tests if desired;
- explore long skipped-message windows;
- explore multi-sender and membership-change matrices.

These are not blockers for Cypher minimal relay research.

## 10. Local machine note

During closure validation, the local Windows machine BSODed during heavy test/build activity. Debugging showed:

- bugcheck `0x164` / `WIN32K_CRITICAL_FAILURE`;
- active process `dwm.exe`;
- failing module `win32kbase.sys`;
- Vanguard and Avast kernel drivers loaded;
- repeated Vanguard `vgm.exe` user-mode crashes in Reliability Monitor;
- SFC repaired Windows files.

This is treated as local environment instability, not project-state corruption. Git commits protected the repo state. Continue heavy validation with `go test -p 1` until local security-driver/OS stability is addressed.

## 11. Next phase direction

Next mainline direction:

    Cypher minimal opaque MLS artifact relay research.

Cypher v0 should route:

- Welcome artifacts;
- application-message artifacts;
- delivery metadata;
- hashes/sizes;
- sender/recipient/conversation identifiers;
- ordering/retry hints.

Cypher v0 must not learn:

- plaintext;
- provider storage;
- signer material;
- private keys;
- group secrets;
- raw MemoryStorage;
- trust-state private material.

After Cypher routing research agrees with Comms/OpenMLS research, begin a maturity/cleanup rung:

    Phase 2E: Research-to-Implementation Promotion

That future rung should move researched systems out of research directories into official implementation structures. It is not a user release.
