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

The command surface did not change. The sidecar still exposes the same Phase 2D commands, but creator-side path hints now point into the device-owned conversation tree.

## 2. Why this changed

Before this checkpoint, Alice creator state and Bob joined state used different layout models:

    Alice:
      dev/conversations/<conversation-label>/

    Bob:
      dev/devices/<device-label>/conversations/<conversation-label>/

That asymmetry was acceptable while proving OpenMLS mechanics, but it became technical debt before Cypher routing and Comms runtime integration.

The sidecar now treats conversation state as device-owned on both creator and joiner paths.

## 3. Changed behavior

The command surface did not change.

Still supported:

    provider-info
    identity-create
    identity-status
    public-bundle-export
    conversation-create
    conversation-load-check
    conversation-add-member
    conversation-join
    message-protect
    message-open

Changed path hints:

    conversation-create now writes Alice creator state under:
      dev/devices/<alice-device>/conversations/<conversation-label>/

    conversation-load-check now reads the device-scoped conversation state for the supplied device label.

    conversation-add-member now reads Alice creator provider storage from the device-scoped conversation path and writes Welcome/add-member artifacts under:
      dev/devices/<alice-device>/conversations/<conversation-label>/

    message-protect now reads Alice creator provider storage from the device-scoped conversation path and writes protected message artifacts under:
      dev/devices/<alice-device>/conversations/<conversation-label>/messages/<message-label>/

Bob joined/open state remains device-scoped as before.

## 4. New path pattern

Alice creator state:

    .carbonstack-openmls-sidecar-state/dev/devices/<alice-device>/conversations/<conversation-label>/conversation-summary.json
    .carbonstack-openmls-sidecar-state/dev/devices/<alice-device>/conversations/<conversation-label>/provider-storage.json

Alice add-member / Welcome state:

    .carbonstack-openmls-sidecar-state/dev/devices/<alice-device>/conversations/<conversation-label>/welcome.bin
    .carbonstack-openmls-sidecar-state/dev/devices/<alice-device>/conversations/<conversation-label>/welcome-manifest.json
    .carbonstack-openmls-sidecar-state/dev/devices/<alice-device>/conversations/<conversation-label>/add-member-summary.json

Alice message-protect state:

    .carbonstack-openmls-sidecar-state/dev/devices/<alice-device>/conversations/<conversation-label>/messages/<message-label>/application-message.bin
    .carbonstack-openmls-sidecar-state/dev/devices/<alice-device>/conversations/<conversation-label>/messages/<message-label>/message-manifest.json
    .carbonstack-openmls-sidecar-state/dev/devices/<alice-device>/conversations/<conversation-label>/messages/<message-label>/message-protect-summary.json

Bob message-open state remains:

    .carbonstack-openmls-sidecar-state/dev/devices/<bob-device>/conversations/<conversation-label>/opened-messages/<message-label>/message-open-summary.json

## 5. Implementation notes

Changed in carbonstack-comms:

    internal/protocol/mls/research/openmls-sidecar/src/state.rs
    internal/protocol/openmls_sidecar_provider_info_test.go

New / expanded helper surface:

    device_conversation_welcome_artifact_path
    device_conversation_welcome_manifest_path
    device_conversation_add_member_summary_path
    device_conversation_messages_dir
    device_conversation_message_dir
    device_conversation_message_artifact_path
    device_conversation_message_manifest_path
    device_conversation_message_protect_summary_path

Updated state functions:

    create_dev_conversation
    load_dev_conversation_status
    add_dev_conversation_member
    protect_dev_message

Already device-scoped and mostly unchanged:

    join_dev_conversation
    open_dev_message

## 6. Test impact

The primary Go test impact was a stale hardcoded old global path in:

    TestOpenMLSSidecarConversationCreate

The test previously expected:

    dev/conversations/carbonstack-test-conversation/conversation-summary.json
    dev/conversations/carbonstack-test-conversation/provider-storage.json

The test now follows the sidecar envelope path hints instead:

    conversation_state_path_hint
    conversation_summary_path_hint
    provider_storage_path_hint

The stale local `stateDir` variable caused a Go compile failure after the first test patch. It was removed because summary/provider file assertions through envelope hints are sufficient.

## 7. Validation

Validation target for this checkpoint:

    cargo check
    cargo test
    go test ./internal/protocol
    go test ./...
    scripts/check-no-rust-artifacts.ps1

The important regression behaviors preserved:

- provider-info still works;
- conversation-create returns Alice device-scoped state path;
- conversation-load-check reads Alice device-scoped state;
- conversation-add-member returns Alice device-scoped Welcome path;
- conversation-join consumes the new Welcome path through envelope path hints;
- message-protect returns Alice device-scoped message artifact path;
- message-open consumes the new message artifact path through envelope path hints;
- two-message continuity still passes;
- out-of-order two-message delivery still passes;
- duplicate/replay open still fails with SecretReuseError;
- corrupt/truncated artifact still fails with message_artifact_invalid / provider.message.invalid.

## 8. Blunders and repair notes

### 8.1 Exact helper block patch failed

The first helper insertion patch looked for an exact `device_conversation_join_summary_path` block. It failed because `cargo fmt` had compacted the helper body into a one-line `.join("join-summary.json")`.

Fix:

    use a regex/line-marker insertion around the helper instead of exact block replacement.

Lesson:

    Do not assume helper body formatting. Use line markers or regex with whitespace flexibility.

### 8.2 Old hardcoded Go path survived initial replacement

The first Go test patch looked for a single `filepath.Join(..., "conversation-summary.json")` expression, but the test built the path in two steps:

    stateDir := filepath.Join(...)
    conversationSummaryPath := filepath.Join(stateDir, "conversation-summary.json")
    providerStoragePath := filepath.Join(stateDir, "provider-storage.json")

Fix:

    replace the local block directly or use line-based patching.

Lesson:

    Inspect the function body before writing exact replacement blocks.

### 8.3 Unused stateDir variable

After replacing the path construction, `stateDir` became unused and Go failed compilation.

Fix:

    remove the local `stateDir` variable and rely on file assertions through envelope path hints.

Lesson:

    Go test cleanup must account for unused locals immediately after patching.

## 9. Allowed claims

Allowed:

- Alice creator conversation state is now device-scoped in the dev OpenMLS sidecar.
- Bob joined/open conversation state remains device-scoped.
- The sidecar now has a consistent device-owned conversation state layout for both creator and joiner paths.
- The command surface did not change.
- Existing OpenMLS sidecar behavior remains the regression target after the path cleanup.

Not allowed:

- old global dev-state compatibility exists;
- a migration path exists;
- Cypher routes MLS payloads;
- Comms runtime send/inbox uses OpenMLS;
- generated message IDs exist;
- production secure storage exists;
- production E2EE exists.

## 10. Next recommended checkpoint

Next recommended checkpoint:

    Cypher MLS artifact routing design docs/recon

Reason:

The OpenMLS sidecar now has:

- device-scoped Alice/Bob conversation state;
- add-member / Welcome export;
- Welcome consume / join;
- message protect/open;
- explicit message labels;
- two-message continuity;
- out-of-order / duplicate / corrupt behavior tests.

The next design question is how Cypher should carry opaque Welcome and application-message artifacts without plaintext or provider storage.
