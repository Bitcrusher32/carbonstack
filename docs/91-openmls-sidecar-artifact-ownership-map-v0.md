# OpenMLS Sidecar Artifact Ownership Map v0

Status: Artifact ownership map / Cypher bridge prep
Component: CarbonStackComms / OpenMLS sidecar
Phase: Post-Phase 2D closure, pre-Cypher relay recon
Previous docs:
- docs/86-openmls-sidecar-phase2d-mainline-closure-result-v0.md
- docs/88-openmls-sidecar-maintainability-promotion-plan-v0.md

## 1. Purpose

This document maps the current OpenMLS sidecar artifacts by owner, producer, consumer, sensitivity, and relay eligibility.

This is a bridge document for Cypher minimal opaque MLS artifact relay research.

## 2. Principles

Cypher may eventually route opaque public/protocol artifacts.

Cypher must not receive or store local private state.

Local-only sensitive files must never be printed, pasted, committed, or relayed.

## 3. Artifact map

### signer.json

Path pattern:

    dev/devices/<device-label>/signer.json

Producer:

    identity-create

Consumer:

    local sidecar commands for the same device

Sensitivity:

    secret-bearing dev signer material

Relay eligibility:

    never relay

Commit eligibility:

    never commit

Notes:

    Dev-only. Not production vault.

### provider-storage.json at device root

Path pattern:

    dev/devices/<device-label>/provider-storage.json

Producer:

    public-bundle-export --write-artifact

Consumer:

    conversation-join for the joining device

Sensitivity:

    sensitive OpenMLS provider storage containing private KeyPackage bundle state

Relay eligibility:

    never relay

Commit eligibility:

    never commit

Notes:

    Required for Welcome consumption after KeyPackage export.

### identity summary/state files

Path pattern:

    dev/devices/<device-label>/identity-summary.json
    dev/devices/<device-label>/identity-state.json
    dev/devices/<device-label>/identity-prep.json

Producer:

    identity-create

Consumer:

    identity-status and diagnostics

Sensitivity:

    sanitized dev metadata

Relay eligibility:

    no, local metadata only

Commit eligibility:

    generated dev state; do not commit

### public-bundle.keypackage.bin

Path pattern:

    dev/devices/<device-label>/public-bundle.keypackage.bin

Producer:

    public-bundle-export --write-artifact

Consumer:

    conversation-add-member by creator device

Sensitivity:

    public OpenMLS KeyPackage artifact

Relay eligibility:

    yes, future onboarding/relay candidate

Commit eligibility:

    generated dev artifact; do not commit

Notes:

    Public does not mean final UX format.

### public-bundle-manifest.json

Path pattern:

    dev/devices/<device-label>/public-bundle-manifest.json

Producer:

    public-bundle-export --write-artifact

Consumer:

    diagnostics / tests

Sensitivity:

    sanitized metadata

Relay eligibility:

    maybe metadata reference only

Commit eligibility:

    generated dev state; do not commit

### conversation provider-storage.json

Path pattern:

    dev/devices/<device-label>/conversations/<conversation-label>/provider-storage.json

Producer:

    conversation-create
    conversation-add-member
    conversation-join
    message-protect
    message-open

Consumer:

    local sidecar commands for that device/conversation

Sensitivity:

    sensitive OpenMLS group/provider state

Relay eligibility:

    never relay

Commit eligibility:

    never commit

Notes:

    This is the main local group state. Cypher must not learn it.

### conversation-summary.json

Path pattern:

    dev/devices/<device-label>/conversations/<conversation-label>/conversation-summary.json

Producer:

    conversation-create
    conversation-join

Consumer:

    diagnostics / tests

Sensitivity:

    sanitized dev metadata

Relay eligibility:

    no, local metadata only for now

Commit eligibility:

    generated dev state; do not commit

### welcome.bin

Path pattern:

    dev/devices/<creator-device>/conversations/<conversation-label>/welcome.bin

Producer:

    conversation-add-member

Consumer:

    conversation-join by invited device

Sensitivity:

    OpenMLS Welcome carrier artifact

Relay eligibility:

    yes, future Cypher opaque artifact candidate

Commit eligibility:

    generated dev artifact; do not commit

Notes:

    Contains outer MlsMessageOut carrier bytes, not final CarbonStack onboarding UX.

### welcome-manifest.json

Path pattern:

    dev/devices/<creator-device>/conversations/<conversation-label>/welcome-manifest.json

Producer:

    conversation-add-member

Consumer:

    diagnostics / tests

Sensitivity:

    sanitized metadata

Relay eligibility:

    maybe metadata reference only

Commit eligibility:

    generated dev state; do not commit

### add-member-summary.json

Path pattern:

    dev/devices/<creator-device>/conversations/<conversation-label>/add-member-summary.json

Producer:

    conversation-add-member

Consumer:

    diagnostics / tests

Sensitivity:

    sanitized metadata

Relay eligibility:

    no, local metadata only for now

Commit eligibility:

    generated dev state; do not commit

### application-message.bin

Path pattern:

    dev/devices/<sender-device>/conversations/<conversation-label>/messages/<message-label>/application-message.bin

Producer:

    message-protect

Consumer:

    message-open by receiver device

Sensitivity:

    encrypted/protected MLS application message artifact

Relay eligibility:

    yes, future Cypher opaque artifact candidate

Commit eligibility:

    generated dev artifact; do not commit

Notes:

    Do not print raw bytes. Cypher may store/route opaque bytes later.

### message-manifest.json

Path pattern:

    dev/devices/<sender-device>/conversations/<conversation-label>/messages/<message-label>/message-manifest.json

Producer:

    message-protect

Consumer:

    diagnostics / tests

Sensitivity:

    sanitized metadata

Relay eligibility:

    maybe metadata reference only

Commit eligibility:

    generated dev state; do not commit

### message-protect-summary.json

Path pattern:

    dev/devices/<sender-device>/conversations/<conversation-label>/messages/<message-label>/message-protect-summary.json

Producer:

    message-protect

Consumer:

    diagnostics / tests

Sensitivity:

    sanitized metadata

Relay eligibility:

    no, local metadata only for now

Commit eligibility:

    generated dev state; do not commit

### message-open-summary.json

Path pattern:

    dev/devices/<receiver-device>/conversations/<conversation-label>/opened-messages/<message-label>/message-open-summary.json

Producer:

    message-open

Consumer:

    diagnostics / tests

Sensitivity:

    sanitized metadata, but may reference plaintext status

Relay eligibility:

    no, receiver-local metadata only

Commit eligibility:

    generated dev state; do not commit

## 4. Future Cypher relay candidates

Candidate relay artifacts:

- public-bundle.keypackage.bin;
- welcome.bin;
- application-message.bin.

Candidate relay metadata:

- artifact type;
- artifact hash;
- artifact size;
- sender device identifier or future device ref;
- recipient device identifier or future device ref;
- conversation identifier or future conversation ref;
- message label or future generated message ID;
- created time;
- delivery status;
- retry/ack hints.

## 5. Never-relay artifacts

Never relay:

- signer.json;
- provider-storage.json;
- raw MemoryStorage JSON;
- private keys;
- local group state;
- trust-state private material;
- plaintext;
- message-open summaries containing plaintext-bearing fields.

## 6. Next use

Cypher relay recon should consume this map before designing storage tables or HTTP routes.
