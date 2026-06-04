# CarbonStack v0.4.2 Runtime OpenMLS Command Contract

Status: v0 runtime command contract
Scope: v0.4.x runtime Comms OpenMLS integration
Primary environment: WSL Debian
Generated: 2026-06-04 15:27:11 -0400

## 1. Purpose

This document defines the first runtime OpenMLS command contract after the v0.4.1 recon.

v0.4.1 found that CarbonStackComms already has runtime `send`, `inbox`, and `ack` commands, but `send` and `inbox` are still stub-era. The OpenMLS/Cypher relay seam exists underneath through sidecar and relay helper tests.

v0.4.2 chooses an explicit dev-only command path before implementation.

This document is a contract and planning record. It does not implement the commands.

## 2. Current repo heads

    carbonstack        83d2f0f docs: record runtime OpenMLS Comms recon
    carbonstack-comms  012c8bf scripts: support source snapshot artifact guard
    carbonstack-cypher 9ab994c docs: point to local operator runbook
    carbonstack-os     1bbbe52 docs: clarify CarbonStackOS target direction

## 3. Working tree status during contract writing

    [carbonstack]
    (clean)

    [carbonstack-comms]
    (clean)

    [carbonstack-cypher]
    (clean)

    [carbonstack-os]
    (clean)

## 4. Command naming decision

Use explicit dev-only command names:

    openmls-send-dev
    openmls-inbox-dev

Reasoning:

    avoids silently replacing existing stub-era send/inbox
    makes the experimental path obvious
    keeps runtime OpenMLS work CLI-only
    avoids pretending UX is mature
    leaves Android, GUI, Relay Space UX, and polished user flows deferred

Existing commands remain:

    send
    inbox
    ack

for now.

## 5. openmls-send-dev contract

Purpose:

    Protect plaintext through the OpenMLS sidecar and submit the resulting application-message artifact through Cypher.

Candidate command shape:

    comms openmls-send-dev \
      --to-device <recipient-cypher-device-id> \
      --sidecar-device-label <sender-sidecar-device-label> \
      --conversation <sidecar-conversation-label> \
      --message <plaintext> \
      [--message-label <label>] \
      [--strict]

Required inputs:

    recipient Cypher device ID
    local Comms state with sender device ID
    Cypher server URL from existing config/state path
    sidecar device label
    sidecar conversation label
    plaintext message

Send-side behavior:

    load local Comms state
    evaluate trust policy for recipient device if applicable
    call OpenMLS sidecar message-protect
    require sidecar success before relay submit
    obtain application-message artifact path from sidecar output
    submit artifact using internal/relay.SubmitOpenMLSArtifactEnvelope
    print envelope ID and relay metadata
    do not print private sidecar state
    do not claim production-send success

Expected output fields:

    command: openmls-send-dev
    status: sent
    recipient_device_id
    sender_device_id
    content_type
    protocol_version
    envelope_id
    server_received_at
    sidecar_message_label
    sidecar_conversation_label
    warning: dev/pre-alpha OpenMLS runtime path

Failure behavior:

    if sidecar message-protect fails, do not submit to Cypher
    if relay submit fails, report failure and leave sidecar artifact local
    if strict trust blocks recipient, do not protect or submit
    if state is missing, fail loudly with required setup hint

## 6. openmls-inbox-dev contract

Purpose:

    Retrieve OpenMLS application-message envelopes from Cypher, write artifacts to a sidecar-compatible location, call message-open, and ack only after sidecar consume succeeds.

Candidate command shape:

    comms openmls-inbox-dev \
      --sidecar-device-label <recipient-sidecar-device-label> \
      --conversation <sidecar-conversation-label> \
      [--limit <n>] \
      [--ack]

Required inputs:

    local Comms state with recipient device ID
    Cypher server URL from existing config/state path
    sidecar device label
    sidecar conversation label

Inbox-side behavior:

    load local Comms state
    fetch inbox through Cypher client
    filter or select application-message envelopes
    validate content_type and protocol_version
    write artifact using internal/relay.WriteOpenMLSArtifactFromEnvelope
    call OpenMLS sidecar message-open
    print plaintext only after sidecar success
    if --ack is set, ack only after message-open succeeds
    if message-open fails, leave envelope unacked and report failure

Expected output fields:

    command: openmls-inbox-dev
    device_id
    queued_envelopes
    opened_envelopes
    unacked_envelopes
    envelope_id
    sender_device_id
    content_type
    protocol_version
    plaintext
    acked: true/false
    warning: dev/pre-alpha OpenMLS runtime path

Failure behavior:

    if envelope content/protocol pair is unsupported, skip or report unsupported
    if artifact write fails, do not ack
    if sidecar message-open fails, do not ack
    if ack fails after successful open, report opened_but_ack_failed

## 7. Ack policy

Ack must remain consume-success gated.

Rule:

    no ack before sidecar message-open success

Meaning:

    retrieval alone is not enough
    artifact write alone is not enough
    sidecar parse/open success is required
    only then may the client ack the envelope

This preserves the CarbonStack rule that messages are acknowledged only after successful receiving-side sidecar consumption.

## 8. State boundary

v0.4.2 does not solve the v0.5.x state model.

Allowed for v0.4.x dev commands:

    existing Comms local state
    sidecar dev-local provider/signer state
    explicit device labels
    explicit conversation labels
    temporary/dev artifacts
    clear warnings

Not solved yet:

    secure vault
    production-safe local encryption
    mature provider storage
    account portability
    backup/export/recovery
    trust-state UX
    PQ/hybrid ciphersuite migration
    hostile-server validation

## 9. Content/protocol pair

OpenMLS application-message envelopes should use:

    content_type=carbonstack.mls.application-message.v0
    protocol_version=carbonstack-openmls-sidecar-v0

The historical invalid pair must remain invalid:

    content_type=carbonstack.message.text.stub.v0
    protocol_version=carbonstack-openmls-sidecar-v0

## 10. Test expectations before implementation

v0.4.3 implementation should add or update tests that prove:

    openmls-send-dev rejects missing required flags
    openmls-send-dev does not submit if sidecar protect fails
    openmls-send-dev submits application-message artifact through Cypher on sidecar success
    openmls-inbox-dev does not ack before message-open success
    openmls-inbox-dev writes and opens application-message artifact
    openmls-inbox-dev acks only after successful message-open when --ack is set

Keep tests local/dev-focused.

## 11. What not to do yet

Do not:

    replace send/inbox
    rename this local-backbone
    create Android/GUI UX
    add public ingress
    add cloudflared/systemd
    add PQ/hybrid ciphersuite work in this rung
    claim production E2EE
    claim hostile-server safety
    claim metadata privacy
    claim polished messaging UX
    claim secure vault/storage

## 12. Recommended next rung

Recommended next rung:

    v0.4.3 first dev-only OpenMLS send command

Focus:

    implement openmls-send-dev in carbonstack-comms
    keep it explicit and dev-only
    call sidecar message-protect
    submit application-message artifact through Cypher relay helper
    add minimal tests
    update carbonstack-comms README/docs as needed

After that:

    v0.4.4 first dev-only OpenMLS inbox/open/ack command
    v0.4.5 runtime smoke proof
    v0.4.6 decide whether a runner profile or local-backbone naming is justified

## 13. Summary

v0.4.2 defines the runtime OpenMLS command contract.

The project should add explicit dev-only OpenMLS runtime commands first, then migrate or replace stub-era send/inbox only after the OpenMLS path is proven and documented.
