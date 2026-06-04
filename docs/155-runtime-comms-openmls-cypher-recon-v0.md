# CarbonStack v0.4.1 Runtime Comms/OpenMLS/Cypher Recon

Status: v0 runtime integration recon checkpoint
Scope: v0.4.x runtime Comms OpenMLS integration
Primary environment: WSL Debian
Generated: 2026-06-04 14:28:06 -0400

## 1. Purpose

This document records the first v0.4.x runtime integration recon after the v0.4.0 broad local deployability pre-release.

v0.4.0 proved the release-package validation ladder and local deployability package surface. v0.4.1 begins the next epoch: understanding how CarbonStackComms runtime commands can move from stub-era send/inbox behavior toward the already-validated OpenMLS/Cypher backbone.

This is a recon and planning document, not an implementation.

## 2. Current repo heads

    carbonstack        6ba7aab docs: make front readme evergreen
    carbonstack-comms  012c8bf scripts: support source snapshot artifact guard
    carbonstack-cypher 9ab994c docs: point to local operator runbook
    carbonstack-os     1bbbe52 docs: clarify CarbonStackOS target direction

## 3. Working tree status during recon

    [carbonstack]
    (clean)

    [carbonstack-comms]
    (clean)

    [carbonstack-cypher]
    (clean)

    [carbonstack-os]
    (clean)

## 4. Main finding

The user-facing Comms CLI already has send, inbox, and ack commands, but they are still stub-era commands.

Current CLI command surface includes:

    dev-create-invite
    claim-invite
    register-device
    list-devices
    verify-device
    trust-events
    mark-device-changed
    revoke-device
    send
    inbox
    ack

Current send path:

    evaluates trust policy
    encrypts user text through the mock/stub provider
    submits a Cypher envelope

Current inbox path:

    retrieves queued Cypher envelopes
    prints stub_plaintext by decrypting through the mock/stub provider
    does not call OpenMLS message-open

Therefore, Comms has a runtime CLI path, but it is not yet the OpenMLS/Cypher runtime path.

## 5. Existing pieces that are already usable

The recon found that the lower-level seams are stronger than the user-facing runtime commands.

Existing Cypher client methods in carbonstack-comms:

    CreateDevInvite
    ClaimInvite
    RegisterDevice
    ListDevices
    SubmitEnvelope
    Inbox
    AckEnvelope

Existing OpenMLS relay helper surface in carbonstack-comms:

    internal/relay.SubmitOpenMLSArtifactEnvelope
    internal/relay.WriteOpenMLSArtifactFromEnvelope

Existing OpenMLS sidecar command surface includes:

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

Existing Cypher API surface is sufficient for the first runtime proof:

    POST /v0/dev/invites
    POST /v0/invites/claim
    POST /v0/devices/register
    GET /v0/accounts/<account_id>/devices
    POST /v0/envelopes
    GET /v0/devices/<device_id>/envelopes
    POST /v0/envelopes/<envelope_id>/ack

## 6. Existing validation evidence

Existing tests already prove the OpenMLS artifact relay path through Cypher:

    sidecar message-protect writes an OpenMLS application-message artifact
    internal/relay submits the artifact bytes through Cypher as an opaque envelope
    recipient retrieves the envelope through Cypher inbox
    internal/relay writes the artifact back to a sidecar-compatible path
    sidecar message-open consumes the artifact
    plaintext matches

Existing real-Cypher test coverage includes a full KeyPackage, Welcome, and application-message relay lifecycle through a real local Cypher server.

The v0.4.1 recon baseline validation also passed:

    local-cypher
    doctor
    core --clean-generated

During core, Comms package tests and Cypher package tests passed. Known OpenMLS generated roots were removed through --clean-generated. Final clean status was clean across all four repos.

## 7. Key gap

The gap is not that OpenMLS/Cypher relay is missing.

The gap is that the user-facing Comms runtime CLI does not yet call the OpenMLS sidecar and relay helpers.

Current stub-era commands:

    send
    inbox
    ack

should not be silently replaced in one step.

## 8. Recommended implementation strategy

Recommended strategy:

    add parallel dev-only OpenMLS runtime commands first
    keep existing send/inbox/ack stable for now
    prove OpenMLS runtime send/inbox behavior under explicit dev command names
    only later decide whether the old send/inbox commands become aliases or are replaced

Candidate command names:

    openmls-send-dev
    openmls-inbox-dev

Alternative command names:

    send-openmls-dev
    inbox-openmls-dev

Rationale:

    avoids breaking old stub-era CLI behavior immediately
    makes the experimental path explicit
    avoids pretending this is mature UX
    creates a clear testable runtime seam
    keeps Android/GUI/UX work deferred

## 9. First runtime proof shape

First send-side proof should likely do:

    load local Comms state
    require registered sender device state
    require explicit recipient device ID
    require explicit sidecar device label
    require explicit sidecar conversation label
    call sidecar message-protect with plaintext
    obtain application-message artifact path
    submit artifact through internal/relay.SubmitOpenMLSArtifactEnvelope
    print envelope_id and metadata
    do not claim production send UX

First inbox-side proof should likely do:

    load local Comms state
    query Cypher inbox for current device ID
    filter or select application-message envelopes
    write artifact through internal/relay.WriteOpenMLSArtifactFromEnvelope
    call sidecar message-open
    print plaintext only after sidecar success
    ack only after sidecar consume succeeds
    clearly mark dev/pre-alpha behavior

## 10. Ack policy

Ack should remain consume-success gated.

For OpenMLS runtime inbox:

    do not ack before message-open succeeds
    ack only after sidecar consume succeeds
    if message-open fails, leave envelope queued or explicitly report unacked state

This preserves the current CarbonStack rule that the receiving side should only ack after successful sidecar consumption.

## 11. What not to do yet

Do not:

    rename this local-backbone yet
    implement Android or GUI work
    implement CarbonStackOS work
    add public ingress
    add systemd/cloudflared
    claim production E2EE
    claim hostile-server safety
    claim metadata privacy
    introduce PQ/hybrid ciphersuite work in this rung
    build a broad negative-path suite unless the runtime seam exposes a direct blocker
    replace existing send/inbox without an explicit migration decision

## 12. Suggested next rung

Recommended next rung:

    v0.4.2 runtime OpenMLS command contract

Focus:

    choose dev command names
    define CLI flags
    define state dependencies
    define sidecar invocation boundaries
    define output schema and warnings
    define ack-on-success behavior
    define tests before implementation

After that:

    v0.4.3 first dev-only OpenMLS send command
    v0.4.4 first dev-only OpenMLS inbox/open/ack command
    v0.4.5 runtime smoke proof
    v0.4.6 decide whether local-backbone is justified yet

## 13. Summary

v0.4.1 recon shows that CarbonStackComms has a stub-era runtime CLI and a separately validated OpenMLS/Cypher relay backbone.

The next work should join those carefully through explicit dev-only OpenMLS runtime commands, not by prematurely replacing send/inbox or claiming mature messaging UX.
