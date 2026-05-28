# OpenMLS KeyPackage and Welcome Relay Through Cypher Result v0

Status: Integration proof result
Component: CarbonStackComms + CarbonStackCypher envelope contract + OpenMLS sidecar
Phase: v0.2.53 KeyPackage + Welcome relay proof
Previous docs:
- docs/100-cypher-openmls-envelope-content-type-result-v0.md
- docs/101-comms-openmls-cypher-relay-bridge-recon-v0.md
- docs/102-comms-openmls-artifact-relay-helper-result-v0.md
- docs/103-comms-cypher-openmls-application-artifact-bridge-helper-result-v0.md
- docs/104-openmls-application-message-relay-through-cypher-result-v0.md

## 1. Summary

This checkpoint records the first KeyPackage + Welcome relay proof across the Comms relay helper boundary, a Cypher-compatible envelope server, and the promoted OpenMLS sidecar.

The proof validates:

    Bob public-bundle-export --write-artifact writes public-bundle.keypackage.bin
    internal/relay submits the KeyPackage artifact bytes as a Cypher OpenMLS envelope
    Alice inbox retrieves the opaque KeyPackage envelope payload
    internal/relay writes the KeyPackage payload back to a local artifact file
    Alice conversation-add-member consumes the relayed KeyPackage artifact
    Alice conversation-add-member writes welcome.bin
    internal/relay submits the Welcome artifact bytes as a Cypher OpenMLS envelope
    Bob inbox retrieves the opaque Welcome envelope payload
    internal/relay writes the Welcome payload back to a local artifact file
    Bob conversation-join consumes the relayed Welcome artifact
    Bob joined group is reloadable

This is a dev/test proof.

It is not polished Comms runtime UX.

It is not production E2EE.

It is not a complete all-in-one KeyPackage -> Welcome -> application-message relay lifecycle test yet.

## 2. What changed

CarbonStackComms added a protocol-level test proving the KeyPackage and Welcome relay path.

The protocol-local Cypher-compatible httptest server now accepts all current OpenMLS relay content types:

    carbonstack.mls.keypackage.v0
    carbonstack.mls.welcome.v0
    carbonstack.mls.application-message.v0

The test uses:

    internal/relay.SubmitOpenMLSArtifactEnvelope
    internal/relay.WriteOpenMLSArtifactFromEnvelope
    internal/client.CypherClient
    promoted OpenMLS sidecar public-bundle-export
    promoted OpenMLS sidecar conversation-create
    promoted OpenMLS sidecar conversation-add-member
    promoted OpenMLS sidecar conversation-join
    a protocol-local Cypher-compatible httptest envelope server

## 3. Validated KeyPackage flow

The validated KeyPackage relay flow is:

    Bob creates identity state
    Bob exports public-bundle.keypackage.bin
    relay helper reads KeyPackage artifact bytes
    relay helper submits /v0/envelopes request with carbonstack.mls.keypackage.v0
    test Cypher server stores queued envelope for Alice
    Alice inbox retrieves queued KeyPackage envelope
    relay helper decodes ciphertext_b64
    relay helper writes downloaded public-bundle.keypackage.bin
    Alice conversation-add-member consumes the downloaded KeyPackage artifact

## 4. Validated Welcome flow

The validated Welcome relay flow is:

    Alice conversation-add-member writes welcome.bin
    relay helper reads Welcome artifact bytes
    relay helper submits /v0/envelopes request with carbonstack.mls.welcome.v0
    test Cypher server stores queued envelope for Bob
    Bob inbox retrieves queued Welcome envelope
    relay helper decodes ciphertext_b64
    relay helper writes downloaded welcome.bin
    Bob conversation-join consumes the downloaded Welcome artifact
    Bob joined group is reloadable

## 5. Preserved boundaries

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

## 6. Validation

Expected validation:

    go test -p 1 ./internal/protocol -run "TestOpenMLSSidecarKeyPackageWelcomeRelayThroughCypherEnvelope"
    go test -p 1 ./internal/protocol -run "TestOpenMLSSidecarApplicationMessageRelayThroughCypherEnvelope|TestOpenMLSSidecarKeyPackageWelcomeRelayThroughCypherEnvelope"
    go test -p 1 ./internal/relay
    go test -p 1 ./internal/protocol
    go test -p 1 ./...
    powershell -ExecutionPolicy Bypass -File .\scripts\check-no-rust-artifacts.ps1

## 7. Next rung

Next planned rung:

    v0.2.54 — full KeyPackage -> Welcome -> application-message relay lifecycle proof.

Goal:

    prove the complete sidecar lifecycle in one Cypher envelope path:
    KeyPackage relay
    Welcome relay
    application-message relay
    message-open plaintext recovery

This should remain a dev/test proof, not polished runtime UX.
