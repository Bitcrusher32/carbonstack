# OpenMLS Sidecar Alice Device-Scoped State Layout Recon v0

Status: API/code recon / pre-implementation result
Component: CarbonStackComms / OpenMLS sidecar
Phase: Phase 2D
Previous docs:
- docs/78-openmls-sidecar-multi-message-continuity-result-v0.md
- docs/79-openmls-sidecar-message-ordering-replay-plan-v0.md
- docs/80-openmls-sidecar-message-ordering-replay-api-recon-v0.md
- docs/81-openmls-sidecar-message-ordering-replay-test-result-v0.md
- docs/82-openmls-sidecar-alice-device-scoped-state-layout-plan-v0.md

## 1. Purpose

This document records code reconnaissance for the Alice device-scoped state-layout cleanup.

v0.2.39 validated ordering/replay/corrupt-artifact behavior for the OpenMLS sidecar.

The next cleanup target is the known state-layout asymmetry:

    Alice/current creator state:
      .carbonstack-openmls-sidecar-state/dev/conversations/<conversation-label>/

    Bob joined/open state:
      .carbonstack-openmls-sidecar-state/dev/devices/<device-label>/conversations/<conversation-label>/

This recon confirms where the old global Alice path is used and how to hard-cut it to device-scoped state.

## 2. Decision carried forward

Use a hard-cut dev-state break.

Do not support migration compatibility with the old global path.

Rationale:

- `.carbonstack-openmls-sidecar-state/` is ignored generated dev state.
- There is no user-data migration requirement.
- Supporting two layouts creates confusing path precedence.
- Future Cypher and Comms runtime integration should not inherit transitional path logic.
- Existing tests already reset sidecar state.

## 3. Current global helper set

The current global conversation helpers are defined in `state.rs`.

Global root helper:

    conversation_state_dir(conversation_label)

Current root:

    .carbonstack-openmls-sidecar-state/dev/conversations/<conversation-label>/

Derived global helpers:

    conversation_summary_path(conversation_label)
    conversation_provider_storage_path(conversation_label)
    conversation_welcome_artifact_path(conversation_label)
    conversation_welcome_manifest_path(conversation_label)
    conversation_add_member_summary_path(conversation_label)
    conversation_messages_dir(conversation_label)
    conversation_message_dir(conversation_label, message_label)
    conversation_message_artifact_path(conversation_label, message_label)
    conversation_message_manifest_path(conversation_label, message_label)
    conversation_message_protect_summary_path(conversation_label, message_label)

These helpers currently drive Alice creator state, Alice add-member/Welcome export state, and Alice message-protect artifact state.

## 4. Existing device-scoped helper set

The existing device-scoped helpers are also defined in `state.rs`.

Device-scoped root:

    device_conversation_state_dir(device_label, conversation_label)

Current root:

    .carbonstack-openmls-sidecar-state/dev/devices/<device-label>/conversations/<conversation-label>/

Existing derived helpers:

    device_conversation_summary_path(device_label, conversation_label)
    device_conversation_provider_storage_path(device_label, conversation_label)
    device_conversation_join_summary_path(device_label, conversation_label)
    device_conversation_message_open_summary_path(device_label, conversation_label, message_label)

These already support Bob's joined conversation state and Bob's message-open summary state.

## 5. Functions that must change

The recon identified four Alice-side functions that still use the global helper set.

### 5.1 create_dev_conversation

Current behavior:

    conversation_state_dir(conversation_label)
    conversation_summary_path(conversation_label)
    conversation_provider_storage_path(conversation_label)

This writes Alice creator state to:

    dev/conversations/<conversation-label>/

Target behavior:

    device_conversation_state_dir(device_label, conversation_label)
    device_conversation_summary_path(device_label, conversation_label)
    device_conversation_provider_storage_path(device_label, conversation_label)

New state path:

    dev/devices/<device-label>/conversations/<conversation-label>/

### 5.2 load_dev_conversation_status

Current behavior:

    conversation_state_dir(conversation_label)
    conversation_summary_path(conversation_label)
    conversation_provider_storage_path(conversation_label)

Target behavior:

    device_conversation_state_dir(device_label, conversation_label)
    device_conversation_summary_path(device_label, conversation_label)
    device_conversation_provider_storage_path(device_label, conversation_label)

This keeps command surface unchanged:

    conversation-load-check --device-label <device> --conversation-label <conversation>

but makes `--device-label` meaningful for locating the conversation state.

### 5.3 add_dev_conversation_member

Current behavior:

    conversation_state_dir(conversation_label)
    conversation_summary_path(conversation_label)
    conversation_provider_storage_path(conversation_label)
    conversation_welcome_artifact_path(conversation_label)
    conversation_welcome_manifest_path(conversation_label)
    conversation_add_member_summary_path(conversation_label)

Target behavior:

    device_conversation_state_dir(device_label, conversation_label)
    device_conversation_summary_path(device_label, conversation_label)
    device_conversation_provider_storage_path(device_label, conversation_label)
    device_conversation_welcome_artifact_path(device_label, conversation_label)
    device_conversation_welcome_manifest_path(device_label, conversation_label)
    device_conversation_add_member_summary_path(device_label, conversation_label)

This moves Alice's Welcome artifact from:

    dev/conversations/<conversation-label>/welcome.bin

to:

    dev/devices/<alice-device>/conversations/<conversation-label>/welcome.bin

### 5.4 protect_dev_message

Current behavior:

    conversation_state_dir(conversation_label)
    conversation_provider_storage_path(conversation_label)
    conversation_message_dir(conversation_label, message_label)
    conversation_message_artifact_path(conversation_label, message_label)
    conversation_message_manifest_path(conversation_label, message_label)
    conversation_message_protect_summary_path(conversation_label, message_label)

Target behavior:

    device_conversation_state_dir(device_label, conversation_label)
    device_conversation_provider_storage_path(device_label, conversation_label)
    device_conversation_message_dir(device_label, conversation_label, message_label)
    device_conversation_message_artifact_path(device_label, conversation_label, message_label)
    device_conversation_message_manifest_path(device_label, conversation_label, message_label)
    device_conversation_message_protect_summary_path(device_label, conversation_label, message_label)

This moves Alice's message artifacts from:

    dev/conversations/<conversation-label>/messages/<message-label>/application-message.bin

to:

    dev/devices/<alice-device>/conversations/<conversation-label>/messages/<message-label>/application-message.bin

## 6. Functions that should remain mostly unchanged

### 6.1 join_dev_conversation

Already uses:

    device_conversation_state_dir(device_label, conversation_label)
    device_conversation_summary_path(device_label, conversation_label)
    device_conversation_provider_storage_path(device_label, conversation_label)
    device_conversation_join_summary_path(device_label, conversation_label)

No path migration needed.

However, manual commands and tests must consume the new Alice device-scoped Welcome path returned by add-member.

### 6.2 open_dev_message

Already uses:

    device_conversation_state_dir(device_label, conversation_label)
    device_conversation_provider_storage_path(device_label, conversation_label)
    device_conversation_message_open_summary_path(device_label, conversation_label, message_label)

No receiver-side path migration needed.

However, tests must pass the new Alice device-scoped message artifact path returned by message-protect.

## 7. New helper functions recommended

Add the following helpers:

    device_conversation_welcome_artifact_path(device_label, conversation_label)
    device_conversation_welcome_manifest_path(device_label, conversation_label)
    device_conversation_add_member_summary_path(device_label, conversation_label)
    device_conversation_messages_dir(device_label, conversation_label)
    device_conversation_message_dir(device_label, conversation_label, message_label)
    device_conversation_message_artifact_path(device_label, conversation_label, message_label)
    device_conversation_message_manifest_path(device_label, conversation_label, message_label)
    device_conversation_message_protect_summary_path(device_label, conversation_label, message_label)

These should compose from:

    device_conversation_state_dir(device_label, conversation_label)

## 8. Deprecated helper functions after implementation

After implementation, these global helpers should either be deleted or left unused temporarily until dead-code cleanup:

    conversation_state_dir
    conversation_summary_path
    conversation_provider_storage_path
    conversation_welcome_artifact_path
    conversation_welcome_manifest_path
    conversation_add_member_summary_path
    conversation_messages_dir
    conversation_message_dir
    conversation_message_artifact_path
    conversation_message_manifest_path
    conversation_message_protect_summary_path

Recommendation:

- For the first implementation patch, do not force-delete all old helpers if that causes noisy Rust warnings or patch risk.
- Prefer replacing all call sites first.
- Then run `cargo check` and inspect dead-code warnings.
- Delete unused global helpers in the same commit only if straightforward.
- Do not keep compatibility reads from old global paths.

## 9. main.rs impact

`main.rs` mostly does not need semantic routing changes.

It prints path hints from result structs:

    conversation_state_path_hint
    conversation_summary_path_hint
    provider_storage_path_hint
    welcome_artifact_path_hint
    message_artifact_path_hint
    message_manifest_path_hint
    message_protect_summary_path_hint
    message_open_summary_path_hint

Therefore, once `state.rs` result paths change, output envelopes should naturally reflect the new layout.

Command routing remains unchanged:

    conversation-create
    conversation-load-check
    conversation-add-member
    conversation-join
    message-protect
    message-open

No new CLI flags are needed.

## 10. Go test impact

The recon indicates many tests use path hints from envelopes rather than hardcoded old paths.

Likely stable because of path-hint chaining:

- conversation-join consumes:
  - `addMemberEnvelope.Data.WelcomeArtifactPathHint`
- message-open consumes:
  - `messageProtectEnvelope.Data.MessageArtifactPathHint`
- corrupt-artifact test builds from:
  - `message1ProtectEnvelope.Data.MessageArtifactPathHint`

This is good. The tests should mostly follow the new path automatically.

However, inspect/update tests that directly assert old global path substrings.

Search terms:

    dev\\conversations
    dev/conversations
    conversation_state_path_hint
    WelcomeArtifactPathHint
    MessageArtifactPathHint
    MessageProtectSummaryPathHint
    messages\\message-
    messages/message-

Expected updates:

- Conversation-create path hints should contain:
  - `dev\devices\carbonstack-alice-device\conversations\carbonstack-test-conversation`
- Add-member Welcome path hints should contain:
  - `dev\devices\carbonstack-alice-device\conversations\carbonstack-test-conversation\welcome.bin`
- Message-protect artifact path hints should contain:
  - `dev\devices\carbonstack-alice-device\conversations\carbonstack-test-conversation\messages\<message-label>\application-message.bin`
- Bob open summary paths remain:
  - `dev\devices\carbonstack-bob-device\conversations\carbonstack-test-conversation\opened-messages\<message-label>\message-open-summary.json`

## 11. Implementation order recommendation

Implement in this order:

1. Add new device-scoped Alice-side helper functions.
2. Change `create_dev_conversation` to use device-scoped helpers.
3. Change `load_dev_conversation_status` to use device-scoped helpers.
4. Change `add_dev_conversation_member` to use device-scoped helpers.
5. Change `protect_dev_message` to use device-scoped helpers.
6. Run `cargo fmt`.
7. Run `cargo check`.
8. Patch any dead-code or missing-helper errors.
9. Run targeted manual or Go tests.
10. Update Go path expectations if needed.
11. Run full OpenMLS sidecar test subset.
12. Run full `go test ./internal/protocol`.
13. Run `go test ./...`.
14. Run Rust artifact guard.

## 12. Validation commands after implementation

From sidecar crate:

    cargo check
    cargo test

From carbonstack-comms:

    go test ./internal/protocol -run "TestOpenMLSSidecarProviderInfoCommand|TestOpenMLSSidecarUnsupportedCommandEnvelope|TestOpenMLSSidecarMessageProtectOpenTwoSequentialMessages|TestOpenMLSSidecarMessageOpenOutOfOrderTwoMessageDelivery|TestOpenMLSSidecarMessageOpenDuplicateRejected|TestOpenMLSSidecarMessageOpenCorruptArtifactRejected"

    go test ./internal/protocol
    go test ./...

    powershell -ExecutionPolicy Bypass -File .\scripts\check-no-rust-artifacts.ps1

## 13. Manual proof target

Manual proof should show new Alice paths.

Expected after implementation:

    conversation-create:
      conversation_state_path_hint = .carbonstack-openmls-sidecar-state\dev\devices\carbonstack-alice-device\conversations\carbonstack-test-conversation

    conversation-add-member:
      welcome_artifact_path_hint = .carbonstack-openmls-sidecar-state\dev\devices\carbonstack-alice-device\conversations\carbonstack-test-conversation\welcome.bin

    message-protect:
      message_artifact_path_hint = .carbonstack-openmls-sidecar-state\dev\devices\carbonstack-alice-device\conversations\carbonstack-test-conversation\messages\message-0001\application-message.bin

    message-open:
      message_open_summary_path_hint = .carbonstack-openmls-sidecar-state\dev\devices\carbonstack-bob-device\conversations\carbonstack-test-conversation\opened-messages\message-0001\message-open-summary.json

## 14. Risks and expected blunders

Expected possible failures:

- `conversation-load-check` may still read old global state.
- `conversation-add-member` may read device-scoped provider storage but write Welcome to old global path.
- `message-protect` may read device-scoped provider storage but write message artifact to old global path.
- Go tests may pass behaviorally but fail hardcoded path assertions.
- Dead-code warnings may appear for old global helpers.
- Manual commands copied from old docs may use old Welcome/message paths and fail.

Mitigation:

- Prefer envelope path hints over hardcoded manual paths.
- Update docs after implementation.
- Hard reset `.carbonstack-openmls-sidecar-state/` for manual tests.
- Do not support fallback to old global path.

## 15. Result doc target after implementation

After implementation, create:

    docs/84-openmls-sidecar-alice-device-scoped-state-result-v0.md

It should record:

- new device-scoped Alice path;
- changed functions;
- changed artifacts;
- test results;
- blunders;
- allowed/non-allowed claims;
- next recommended Cypher routing design step.
