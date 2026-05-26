# OpenMLS Sidecar Alice Device-Scoped State Layout Result v0

Status: Validated implementation checkpoint
Component: CarbonStackComms / OpenMLS sidecar
Phase: Phase 2D
Previous docs:
- docs/82-openmls-sidecar-alice-device-scoped-state-layout-plan-v0.md
- docs/83-openmls-sidecar-alice-device-scoped-state-recon-v0.md

## 1. Summary

This checkpoint hard-cuts the OpenMLS sidecar creator/Alice conversation state layout from a global conversation path to a device-scoped conversation path.

Before:

    .carbonstack-openmls-sidecar-state/dev/conversations/<conversation-label>/

After:

    .carbonstack-openmls-sidecar-state/dev/devices/<device-label>/conversations/<conversation-label>/

This is a dev-state break only. No migration compatibility was added.

## 2. Why this changed

Before this checkpoint, Alice creator state and Bob joined state used different layout models:

    Alice:
      dev/conversations/<conversation-label>/

    Bob:
      dev/devices/<device-label>/conversations/<conversation-label>/

That asymmetry was acceptable while proving OpenMLS mechanics, but it was becoming technical debt before Cypher routing and Comms runtime integration.

The sidecar now treats conversation state as device-owned on both creator and joiner paths.

## 3. Changed behavior

The command surface did not change.

Still supported:

    conversation-create
    conversation-load-check
    conversation-add-member
    conversation-join
    message-protect
    message-open

Changed path hints:

    conversation-create now writes Alice creator state under:
      dev/devices/<alice-device>/conversations/<conversation-label>/

    conversation-add-member now writes Alice Welcome artifacts under:
      dev/devices/<alice-device>/conversations/<conversation-label>/

    message-protect now writes Alice protected message artifacts under:
      dev/devices/<alice-device>/conversations/<conversation-label>/messages/<message-label>/

Bob joined/open state remains device-scoped as before.

## 4. New/used path pattern

Alice creator state:

    .carbonstack-openmls-sidecar-state/dev/devices/<alice-device>/conversations/<conversation-label>/conversation-summary.json
    .carbonstack-openmls-sidecar-state/dev/devices/<alice-device>/conversations/<conversation-label>/provider-storage.json

Alice add-member/Welcome state:

    .carbonstack-openmls-sidecar-state/dev/devices/<alice-device>/conversations/<conversation-label>/welcome.bin
    .carbonstack-openmls-sidecar-state/dev/devices/<alice-device>/conversations/<conversation-label>/welcome-manifest.json
    .carbonstack-openmls-sidecar-state/dev/devices/<alice-device>/conversations/<conversation-label>/add-member-summary.json

Alice message-protect state:

    .carbonstack-openmls-sidecar-state/dev/devices/<alice-device>/conversations/<conversation-label>/messages/<message-label>/application-message.bin
    .carbonstack-openmls-sidecar-state/dev/devices/<alice-device>/conversations/<conversation-label>/messages/<message-label>/message-manifest.json
    .carbonstack-openmls-sidecar-state/dev/devices/<alice-device>/conversations/<conversation-label>/messages/<message-label>/message-protect-summary.json

Bob message-open state:

    .carbonstack-openmls-sidecar-state/dev/devices/<bob-device>/conversations/<conversation-label>/opened-messages/<message-label>/message-open-summary.json

## 5. Implementation notes

Changed in carbonstack-comms:

    internal/protocol/mls/research/openmls-sidecar/src/state.rs

New/expanded helper surface:

    device_conversation_welcome_artifact_path
    device_conversation_welcome_manifest_path
    device_conversation_add_member_summary_path
    device_conversation_messages_dir
    device_conversation_message_dir
    device_conversation_message_artifact_path
    device_conversation_message_manifest_path
    device_conversation_message_protect_summary_path

Updated functions:

    create_dev_conversation
    load_dev_conversation_status
    add_dev_conversation_member
    protect_dev_message

Already device-scoped and mostly unchanged:

    join_dev_conversation
    open_dev_message

## 6. Validation

Validation should include:

    cargo check
    cargo test
    go test ./internal/protocol
    go test ./...
    scripts/check-no-rust-artifacts.ps1

The key regression behaviors to preserve are:

- provider-info still works;
- conversation-create returns Alice device-scoped state path;
- conversation-add-member returns Alice device-scoped Welcome path;
- conversation-join consumes the new Welcome path;
- message-protect returns Alice device-scoped message artifact path;
- message-open consumes the new message artifact path;
- two-message continuity still passes;
- out-of-order two-message delivery still passes;
- duplicate/replay open still fails with SecretReuseError;
- corrupt/truncated artifact still fails with message_artifact_invalid / provider.message.invalid.

## 7. Allowed claims

Allowed:

- Alice creator conversation state is now device-scoped in the dev sidecar.
- Bob joined conversation state remains device-scoped.
- The sidecar has a more consistent device-owned conversation state layout.
- The OpenMLS sidecar tests still validate the core Alice/Bob lifecycle after the path cleanup.

Not allowed:

- production migration exists;
- old global dev-state compatibility exists;
- Cypher routing exists;
- Comms runtime integration exists;
- generated message IDs exist;
- production secure storage exists.

## 8. Next recommended checkpoint

Next recommended checkpoint:

    Cypher MLS artifact routing design docs/recon

Reason:

The OpenMLS sidecar now has:

- device-scoped Alice/Bob conversation state;
- add-member/Welcome export;
- Welcome consume/join;
- message protect/open;
- two-message continuity;
- out-of-order/duplicate/corrupt behavior tests.

The next design question is how Cypher should carry opaque Welcome and application-message artifacts without plaintext or provider storage.
