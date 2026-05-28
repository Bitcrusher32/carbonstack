# OpenMLS Application Message Relay Through Cypher Result v0

Status: Integration proof result
Component: CarbonStackComms + CarbonStackCypher envelope contract + OpenMLS sidecar
Phase: v0.2.52 application-message relay proof
Previous docs:
- docs/100-cypher-openmls-envelope-content-type-result-v0.md
- docs/101-comms-openmls-cypher-relay-bridge-recon-v0.md
- docs/102-comms-openmls-artifact-relay-helper-result-v0.md
- docs/103-comms-cypher-openmls-application-artifact-bridge-helper-result-v0.md

## 1. Summary

This checkpoint records the first end-to-end application-message artifact relay proof across the Comms relay helper boundary, a Cypher-compatible envelope server, and the promoted OpenMLS sidecar.

The proof validates:

    sidecar message-protect writes application-message.bin
    internal/relay submits the artifact bytes as a Cypher OpenMLS envelope
    recipient inbox retrieves the opaque envelope payload
    internal/relay writes the payload back to a local application-message.bin
    sidecar message-open consumes the written artifact
    plaintext matches

This is a dev/test proof.

It is not polished Comms runtime UX.

It is not production E2EE.

It is not a full KeyPackage/Welcome/application-message lifecycle through Cypher yet.

## 2. What changed

CarbonStackComms added a protocol-level test proving the application-message relay path.

The test uses:

    internal/relay.SubmitOpenMLSArtifactEnvelope
    internal/relay.WriteOpenMLSArtifactFromEnvelope
    internal/client.CypherClient
    promoted OpenMLS sidecar message-protect/message-open helpers
    a protocol-local Cypher-compatible httptest envelope server

## 3. Validated flow

The validated flow is:

    setup Alice/Bob OpenMLS sidecar conversation
    Alice protects plaintext with message-protect
    sidecar writes application-message.bin
    relay helper reads artifact bytes
    relay helper submits /v0/envelopes request
    test Cypher server stores queued envelope
    Bob inbox retrieves queued envelope
    relay helper decodes ciphertext_b64
    relay helper writes downloaded application-message.bin
    Bob opens the downloaded artifact with message-open
    recovered plaintext equals original plaintext

## 4. Preserved boundaries

This proof does not:

    wire comms send
    wire comms inbox
    automatically acknowledge after sidecar consume
    parse MLS internals in Cypher
    parse MLS internals in relay helper
    mutate trust-state
    relay signer.json
    relay provider-storage.json
    add Cypher routes
    add Cypher migrations
    claim production readiness

## 5. Validation

Expected validation:

    go test -p 1 ./internal/relay
    go test -p 1 ./internal/protocol -run "TestOpenMLSSidecarApplicationMessageRelayThroughCypherEnvelope"
    go test -p 1 ./internal/protocol
    go test -p 1 ./...
    powershell -ExecutionPolicy Bypass -File .\scripts\check-no-rust-artifacts.ps1

## 6. Next rung

Next planned rung:

    v0.2.53 — KeyPackage + Welcome relay proof.

Goal:

    Bob exports public-bundle.keypackage.bin
    relay sends KeyPackage artifact through Cypher envelope
    Alice retrieves and writes KeyPackage artifact
    Alice conversation-add-member consumes KeyPackage and writes welcome.bin
    relay sends Welcome artifact through Cypher envelope
    Bob retrieves and writes Welcome artifact
    Bob conversation-join consumes Welcome

After that, the full KeyPackage -> Welcome -> application-message relay lifecycle can be stitched together.
