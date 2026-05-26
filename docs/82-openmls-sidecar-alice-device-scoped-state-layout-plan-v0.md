# OpenMLS Sidecar Alice Device-Scoped State Layout Cleanup Plan v0

Status: Plan / pre-recon
Component: CarbonStackComms / OpenMLS sidecar
Phase: Phase 2D
Previous docs:
- docs/78-openmls-sidecar-multi-message-continuity-result-v0.md
- docs/79-openmls-sidecar-message-ordering-replay-plan-v0.md
- docs/80-openmls-sidecar-message-ordering-replay-api-recon-v0.md
- docs/81-openmls-sidecar-message-ordering-replay-test-result-v0.md

## 1. Purpose

v0.2.39 validated important message-delivery behavior for the OpenMLS sidecar:

- two-message ordered delivery;
- two-message same-sender out-of-order delivery;
- duplicate/replay rejection through SecretReuseError;
- corrupt/truncated artifact rejection through message_artifact_invalid / provider.message.invalid.

The sidecar is now mature enough to clean up a known state-layout asymmetry before Cypher routing or Comms runtime integration.

Current asymmetry:

    Alice:
      .carbonstack-openmls-sidecar-state/dev/conversations/<conversation-label>/

    Bob:
      .carbonstack-openmls-sidecar-state/dev/devices/<device-label>/conversations/<conversation-label>/

This document plans a hard-cut cleanup so Alice also uses device-scoped conversation state.

## 2. Decision

Use a hard-cut dev-state break.

Do not support both old and new Alice layouts.

Reason:

- `.carbonstack-openmls-sidecar-state/` is ignored generated dev state.
- There is no user data migration requirement.
- Supporting two layouts creates confusing path precedence.
- Future Cypher/Comms integration should not inherit transitional state logic.
- Tests can reset sidecar state and validate the new layout directly.

The cleanup target is:

    dev/devices/<device-label>/conversations/<conversation-label>/

for every device-owned conversation state path.

## 3. Current layout

### Alice global creator state

Currently, conversation-create writes Alice creator state to:

    .carbonstack-openmls-sidecar-state/dev/conversations/<conversation-label>/

Files:

    conversation-summary.json
    provider-storage.json

After add-member, Alice also writes:

    welcome.bin
    welcome-manifest.json
    add-member-summary.json

After message-protect, Alice writes:

    messages/<message-label>/application-message.bin
    messages/<message-label>/message-manifest.json
    messages/<message-label>/message-protect-summary.json

### Bob device-scoped joined state

conversation-join writes Bob state to:

    .carbonstack-openmls-sidecar-state/dev/devices/<device-label>/conversations/<conversation-label>/

Files:

    conversation-summary.json
    join-summary.json
    provider-storage.json
    opened-messages/<message-label>/message-open-summary.json

## 4. Target layout

All conversation state should be device-scoped.

### Alice creator state target

conversation-create should write Alice creator state to:

    .carbonstack-openmls-sidecar-state/dev/devices/<alice-device>/conversations/<conversation-label>/

Files:

    conversation-summary.json
    provider-storage.json

After add-member:

    welcome.bin
    welcome-manifest.json
    add-member-summary.json

After message-protect:

    messages/<message-label>/application-message.bin
    messages/<message-label>/message-manifest.json
    messages/<message-label>/message-protect-summary.json

### Bob joined state remains

conversation-join should continue writing Bob joined state to:

    .carbonstack-openmls-sidecar-state/dev/devices/<bob-device>/conversations/<conversation-label>/

Files:

    conversation-summary.json
    join-summary.json
    provider-storage.json
    opened-messages/<message-label>/message-open-summary.json

## 5. Affected commands

Commands that must change:

    conversation-create
    conversation-load-check
    conversation-add-member
    message-protect

Commands that likely do not need path changes:

    conversation-join
    message-open

Reason:

- `conversation-join` already writes joined state under `dev/devices/<device>/conversations/<conversation>/`.
- `message-open` already reads Bob device-scoped joined state.
- Alice-side creator operations still use global conversation paths.

## 6. Affected state.rs functions

Current global conversation path functions likely need replacement or refactor:

    conversation_state_dir(conversation_label)
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

Existing device-scoped functions already point in the desired direction:

    device_conversations_dir(device_label)
    device_conversation_state_dir(device_label, conversation_label)
    device_conversation_summary_path(device_label, conversation_label)
    device_conversation_provider_storage_path(device_label, conversation_label)
    device_conversation_join_summary_path(device_label, conversation_label)
    device_conversation_message_open_summary_path(device_label, conversation_label, message_label)

New or refactored functions likely needed:

    device_conversation_welcome_artifact_path(device_label, conversation_label)
    device_conversation_welcome_manifest_path(device_label, conversation_label)
    device_conversation_add_member_summary_path(device_label, conversation_label)
    device_conversation_messages_dir(device_label, conversation_label)
    device_conversation_message_dir(device_label, conversation_label, message_label)
    device_conversation_message_artifact_path(device_label, conversation_label, message_label)
    device_conversation_message_manifest_path(device_label, conversation_label, message_label)
    device_conversation_message_protect_summary_path(device_label, conversation_label, message_label)

Alternative:

    delete global conversation_* functions after replacing all callers, except if they are kept temporarily only for helper composition during the implementation patch.

## 7. Affected main.rs / CLI behavior

Command surface should not change.

Still use:

    conversation-create --device-label <device> --conversation-label <conversation>
    conversation-load-check --device-label <device> --conversation-label <conversation>
    conversation-add-member --device-label <device> --conversation-label <conversation> --member-keypackage <path>
    conversation-join --device-label <device> --conversation-label <conversation> --welcome <path>
    message-protect --device-label <device> --conversation-label <conversation> --message-label <label> --plaintext <text>
    message-open --device-label <device> --conversation-label <conversation> --message-label <label> --message <path>

Only path hints change.

## 8. Expected artifact path changes

### Before

Alice Welcome:

    .carbonstack-openmls-sidecar-state/dev/conversations/carbonstack-test-conversation/welcome.bin

Alice message artifact:

    .carbonstack-openmls-sidecar-state/dev/conversations/carbonstack-test-conversation/messages/message-0001/application-message.bin

### After

Alice Welcome:

    .carbonstack-openmls-sidecar-state/dev/devices/carbonstack-alice-device/conversations/carbonstack-test-conversation/welcome.bin

Alice message artifact:

    .carbonstack-openmls-sidecar-state/dev/devices/carbonstack-alice-device/conversations/carbonstack-test-conversation/messages/message-0001/application-message.bin

Bob joined/open state remains:

    .carbonstack-openmls-sidecar-state/dev/devices/carbonstack-bob-device/conversations/carbonstack-test-conversation/

## 9. Go test impact

The existing tests likely rely on path hints returned by envelopes more than hardcoded paths, but some direct expectations may include old global paths.

Search targets:

    dev\\conversations
    dev/conversations
    conversation_state_path_hint
    welcome_artifact_path_hint
    message_artifact_path_hint
    messages\\message-0001
    messages/message-0001

Expected test updates:

- `conversation-create` should expect `conversation_state_path_hint` under Alice device.
- `conversation-add-member` should return Welcome path under Alice device.
- `conversation-join` should consume `addMemberEnvelope.Data.WelcomeArtifactPathHint`, so it should mostly survive.
- `message-protect` should return message artifact path under Alice device.
- `message-open` should consume `messageProtectEnvelope.Data.MessageArtifactPathHint`, so it should mostly survive.
- Any direct hardcoded path checks need update.

## 10. Manual proof target after implementation

After implementation, the same v0.2.39 behavior should remain valid:

1. Alice identity-create.
2. Bob identity-create.
3. Bob public-bundle-export --write-artifact.
4. Alice conversation-create.
5. Alice conversation-add-member.
6. Bob conversation-join using Alice device-scoped welcome path.
7. Alice message-protect message-0001.
8. Alice message-protect message-0002.
9. Bob opens message-0002 first.
10. Bob opens message-0001 afterward.
11. Duplicate-open still fails with SecretReuseError.
12. Corrupt artifact still fails with message_artifact_invalid / provider.message.invalid.

## 11. Security and artifact rules

Do not print, paste, commit, or inspect casually:

- signer.json;
- provider-storage.json;
- raw MemoryStorage JSON;
- raw group state;
- raw KeyPackage bytes;
- raw Welcome bytes;
- raw application-message bytes.

Allowed:

- path hints;
- file existence;
- file sizes;
- sanitized summary/manifests;
- hashes;
- stdout envelopes.

## 12. Risks

Main risks:

- stale global path helper remains in one caller;
- tests pass by following path hints but docs still describe old paths;
- conversation-load-check breaks if it still reads global state;
- add-member reads new provider storage but writes Welcome to old global path;
- message-protect reads new state but writes messages to old global path;
- duplicate artifact refusal may accidentally check wrong directory;
- cleanup may cause false negatives in tests if old dev state is not reset.

Mitigation:

- hard reset `.carbonstack-openmls-sidecar-state/` in tests;
- update all path hints in one implementation commit;
- run the full OpenMLS sidecar test subset;
- run full `go test ./internal/protocol`;
- run `go test ./...`;
- run Rust artifact guard.

## 13. Recommended next recon doc

Next document:

    docs/83-openmls-sidecar-alice-device-scoped-state-recon-v0.md

It should record:

- exact current callers of global `conversation_*` path helpers;
- exact tests with old global assumptions;
- recommended helper replacement map;
- implementation order;
- expected validation commands.
