# OpenMLS Sidecar Current State Index v0

Status: Current-state index / orientation doc
Component: CarbonStackComms / OpenMLS sidecar
Phase: Post-Phase 2D mainline closure, pre-Cypher relay research
Previous docs:
- docs/84-openmls-sidecar-alice-device-scoped-state-result-v0.md
- docs/85-openmls-sidecar-phase2d-closure-checklist-v0.md
- docs/86-openmls-sidecar-phase2d-mainline-closure-result-v0.md

## 1. Purpose

This document is the current-state map for the OpenMLS sidecar after Phase 2D mainline closure.

Older docs remain valuable historical records. They should not be rewritten wholesale. However, many older docs describe states that were true at the time but are no longer current, especially old Alice-global paths under:

    .carbonstack-openmls-sidecar-state/dev/conversations/<conversation-label>/

The current state is device-scoped.

Use this document as the first read before continuing sidecar, Cypher, or Comms/OpenMLS work.

## 2. Current repo truth

Current Phase 2D mainline closure is recorded in:

    docs/86-openmls-sidecar-phase2d-mainline-closure-result-v0.md

The current OpenMLS sidecar implementation is still in the research path:

    carbonstack-comms/internal/protocol/mls/research/openmls-sidecar

The planned maintainability promotion will keep that research sidecar as a known-good reference and create a promoted maintained sidecar above the research path.

## 3. Current validated sidecar lifecycle

The sidecar validates:

- provider-info;
- identity-create;
- identity-status;
- public-bundle-export;
- public-bundle-export with KeyPackage artifact;
- conversation-create;
- conversation-load-check;
- conversation-add-member;
- Welcome export;
- conversation-join;
- Welcome consume;
- message-protect;
- message-open;
- explicit message labels;
- two-message continuity;
- same-sender two-message out-of-order open;
- duplicate/replay open rejection;
- corrupt/truncated message artifact rejection;
- wrong-device message-open rejection;
- wrong-conversation message-open rejection;
- bidirectional Alice to Bob and Bob to Alice message flow;
- device-scoped creator and joined conversation state.

## 4. Current canonical state layout

Device root:

    .carbonstack-openmls-sidecar-state/dev/devices/<device-label>/

Device-scoped conversation state:

    .carbonstack-openmls-sidecar-state/dev/devices/<device-label>/conversations/<conversation-label>/

Welcome artifact:

    .carbonstack-openmls-sidecar-state/dev/devices/<alice-device>/conversations/<conversation-label>/welcome.bin

Protected message artifact:

    .carbonstack-openmls-sidecar-state/dev/devices/<sender-device>/conversations/<conversation-label>/messages/<message-label>/application-message.bin

Opened-message summary:

    .carbonstack-openmls-sidecar-state/dev/devices/<receiver-device>/conversations/<conversation-label>/opened-messages/<message-label>/message-open-summary.json

The old global creator path is historical only:

    .carbonstack-openmls-sidecar-state/dev/conversations/<conversation-label>/

Current implementation should not rely on that path.

## 5. Current command surface

Supported:

- provider-info;
- identity-create;
- identity-status;
- public-bundle-export;
- conversation-create;
- conversation-load-check;
- conversation-add-member;
- conversation-join;
- message-protect;
- message-open.

Still unsupported:

- state-checkpoint;
- state-load-check.

## 6. Current repo maintainability issue

The sidecar research proof is validated, but the implementation shape is research-sprint shaped.

Known large files in carbonstack-comms:

    internal/protocol/openmls_sidecar_provider_info_test.go
      about 2,938 lines

    internal/protocol/mls/research/openmls-sidecar/src/main.rs
      about 2,718 lines

    internal/protocol/mls/research/openmls-sidecar/src/state.rs
      about 2,261 lines

The sidecar README is stale and still describes an earlier provider-info-only state.

This is not a failure. It is normal research sediment after a successful proof. The next rung should preserve behavior while improving maintainability.

## 7. Current non-goals

Do not claim:

- production E2EE;
- Signal-equivalent security;
- hostile-server proof;
- metadata privacy;
- production vault/secure storage;
- Comms runtime OpenMLS integration;
- Cypher MLS routing;
- Android / Pixel 4a validation;
- generated message IDs;
- sidecar promotion out of research already completed.

## 8. Next planned rungs

v0.2.42:

    Maintainability promotion plan/recon docs only.

v0.2.43:

    Copy/promote sidecar scaffold above research path while preserving behavior.

v0.2.44:

    Split promoted Rust sidecar modules.

v0.2.45:

    Split Go sidecar contract tests and point them at the promoted sidecar.

v0.2.46:

    README/current-state/known-good command cleanup.

v0.2.47:

    Begin Cypher minimal opaque MLS artifact relay recon.

## v0.2.46 current-state update

The maintainability promotion ladder has now landed through v0.2.45:

- v0.2.43 promoted the sidecar scaffold to `carbonstack-comms/internal/protocol/mls/openmls-sidecar`;
- v0.2.44 split promoted Rust modules into `cli.rs`, `envelope.rs`, `paths.rs`, and `schema.rs` while preserving behavior;
- v0.2.45 split the Go sidecar contract tests by ownership;
- v0.2.46 updates READMEs, current-state docs, known-good command references, and stale warning strings.

The active sidecar path is:

    carbonstack-comms/internal/protocol/mls/openmls-sidecar

The research sidecar remains a frozen Phase 2D reference:

    carbonstack-comms/internal/protocol/mls/research/openmls-sidecar

The active Go sidecar contract tests are now split across:

    internal/protocol/openmls_sidecar_helpers_test.go
    internal/protocol/openmls_sidecar_provider_info_test.go
    internal/protocol/openmls_sidecar_identity_test.go
    internal/protocol/openmls_sidecar_public_bundle_test.go
    internal/protocol/openmls_sidecar_conversation_test.go
    internal/protocol/openmls_sidecar_message_test.go
    internal/protocol/openmls_sidecar_message_negative_test.go

Next planned rung:

    v0.2.47 — Cypher minimal opaque MLS artifact relay recon.

