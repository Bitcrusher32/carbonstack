# Cypher Opaque MLS Artifact Relay Plan v0

Status: Plan / recon entrypoint
Component: CarbonStackCypher + CarbonStackComms OpenMLS sidecar
Phase: v0.2.47 Cypher minimal opaque MLS artifact relay recon
Previous docs:
- docs/91-openmls-sidecar-artifact-ownership-map-v0.md
- docs/92-openmls-sidecar-command-schema-matrix-v0.md
- docs/96-openmls-sidecar-readme-current-state-cleanup-result-v0.md

## 1. Purpose

This document defines the first Cypher relay recon target after the OpenMLS sidecar maintainability cleanup.

The goal is to design a minimal relay for opaque MLS artifacts already produced and consumed by the promoted OpenMLS sidecar.

This is recon/planning only. It does not implement Cypher routes, database schema, Comms runtime integration, or sidecar changes.

## 2. Current starting point

The promoted OpenMLS sidecar is active at:

    carbonstack-comms/internal/protocol/mls/openmls-sidecar

The frozen Phase 2D research reference remains at:

    carbonstack-comms/internal/protocol/mls/research/openmls-sidecar

The sidecar can already produce and consume:

    public-bundle.keypackage.bin
    welcome.bin
    application-message.bin

The sidecar can already validate:

- identity lifecycle;
- KeyPackage artifact export;
- conversation-create/load-check;
- add-member / Welcome export;
- join / Welcome consume;
- message-protect;
- message-open;
- two sequential messages;
- same-sender out-of-order open;
- duplicate/replay rejection;
- corrupt/truncated artifact rejection;
- wrong-device rejection;
- wrong-conversation rejection;
- bidirectional Alice/Bob message flow.

## 3. Relay design principle

Cypher should relay opaque MLS artifacts.

Cypher should not parse MLS internals.

Cypher should not understand plaintext.

Cypher should not receive signer/provider storage.

Cypher should not mutate trust-state.

Cypher should behave like a hostile-server-compatible mailbox for opaque artifacts, not as a trusted cryptographic participant.

## 4. Relay candidate artifacts

### KeyPackage artifact

Sidecar artifact:

    public-bundle.keypackage.bin

Producer:

    public-bundle-export --write-artifact

Consumer:

    conversation-add-member --member-keypackage <path>

Relay role:

    recipient/device publishes or makes available a public onboarding artifact.

Sensitivity:

    public protocol artifact, but still generated dev artifact and not final UX.

First Cypher treatment:

    allow opaque upload/download with metadata.

### Welcome artifact

Sidecar artifact:

    welcome.bin

Producer:

    conversation-add-member

Consumer:

    conversation-join --welcome <path>

Relay role:

    creator sends invite/onboarding artifact to invited device.

Sensitivity:

    opaque MLS Welcome carrier artifact. Do not parse or print raw bytes.

First Cypher treatment:

    allow opaque sender-to-recipient delivery.

### Application message artifact

Sidecar artifact:

    application-message.bin

Producer:

    message-protect

Consumer:

    message-open --message <path>

Relay role:

    sender sends protected MLS application message to recipient.

Sensitivity:

    encrypted/protected application message artifact. Do not parse or print raw bytes.

First Cypher treatment:

    allow opaque sender-to-recipient delivery.

## 5. Never-relay artifacts

Cypher must never accept or store:

    signer.json
    provider-storage.json
    raw MemoryStorage JSON
    raw group state
    raw signer/private key material
    plaintext
    message-open summaries containing plaintext
    local trust-state private material

These remain local-only.

## 6. Minimal metadata concept

A minimal relay record may need metadata like:

    artifact_id
    artifact_type
    conversation_ref
    sender_device_ref
    recipient_device_ref
    message_label or artifact_label
    sha256
    size_bytes
    created_at
    delivery_state

Do not design final identity semantics yet.

For v0.2.47, these refs can remain dev labels or dev-safe placeholders as long as the doc explicitly marks them as not final identity design.

## 7. First relay capabilities to recon

The recon should answer:

1. What artifact types should Cypher accept first?
2. What metadata is required to route each artifact?
3. What metadata must be avoided to reduce leakage?
4. Should the first relay be push, pull, or mailbox-style?
5. How should artifact integrity be represented?
6. How should upload/download avoid path traversal and arbitrary file write hazards?
7. What validation belongs in Cypher versus Comms/sidecar?
8. How should duplicate artifact IDs or labels be handled?
9. What are the minimum tests before implementation?
10. What does Cypher need to know before Android/Pixel 4a work later?

## 8. Likely first implementation shape

The likely first implementation should be a minimal mailbox-style relay:

    POST artifact metadata + bytes
    GET/list pending artifact metadata for recipient
    GET/download artifact bytes by artifact id
    mark/ack artifact as received

This does not imply final production API.

## 9. Non-goals

Do not implement in this recon rung:

- production account identity;
- final device identity model;
- hostile-server proof;
- metadata privacy solution;
- push notifications;
- Android app integration;
- Comms runtime integration;
- trust-state mutation;
- secure production storage;
- MLS parsing in Cypher;
- cryptographic verification inside Cypher;
- generated sidecar message IDs;
- group membership evolution beyond the validated sidecar artifacts.

## 10. Definition of done for v0.2.47 recon

This recon is complete when docs answer:

- artifact types;
- artifact ownership;
- relay semantics;
- metadata boundaries;
- non-relay sensitive state;
- proposed minimal API shape;
- proposed storage shape;
- test plan;
- implementation sequence for the next rung.

No code implementation is required for this plan doc.
