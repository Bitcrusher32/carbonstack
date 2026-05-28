# Full OpenMLS Relay Lifecycle Through Cypher Result v0

Status: Integration proof result
Component: CarbonStackComms + CarbonStackCypher envelope contract + OpenMLS sidecar
Phase: v0.2.54 full KeyPackage -> Welcome -> application-message relay lifecycle proof
Previous docs:
- docs/104-openmls-application-message-relay-through-cypher-result-v0.md
- docs/105-openmls-keypackage-welcome-relay-through-cypher-result-v0.md

## 1. Summary

This checkpoint records the first full dev/test OpenMLS relay lifecycle proof through the Comms relay helper boundary and a Cypher-compatible envelope server.

The proof validates the whole current relay chain in one test:

    Bob exports public-bundle.keypackage.bin
    KeyPackage relays through Cypher-compatible /v0/envelopes to Alice
    Alice writes the downloaded KeyPackage artifact
    Alice conversation-add-member consumes it and writes welcome.bin
    Welcome relays through Cypher-compatible /v0/envelopes to Bob
    Bob writes the downloaded Welcome artifact
    Bob conversation-join consumes it
    Alice protects an application message
    application-message.bin relays through Cypher-compatible /v0/envelopes to Bob
    Bob writes the downloaded application-message artifact
    Bob message-open consumes it
    plaintext matches

This is a dev/test proof.

It is not polished Comms runtime UX.

It is not production E2EE.

It is not a deployed server + CLI release.

## 2. What changed

CarbonStackComms added a protocol-level full lifecycle relay test.

The test uses:

    internal/relay.SubmitOpenMLSArtifactEnvelope
    internal/relay.WriteOpenMLSArtifactFromEnvelope
    internal/client.CypherClient
    promoted OpenMLS sidecar public-bundle-export
    promoted OpenMLS sidecar conversation-create
    promoted OpenMLS sidecar conversation-add-member
    promoted OpenMLS sidecar conversation-join
    promoted OpenMLS sidecar message-protect
    promoted OpenMLS sidecar message-open
    a protocol-local Cypher-compatible httptest envelope server

## 3. Validated lifecycle

Validated lifecycle:

    KeyPackage relay
    Welcome relay
    application-message relay
    message-open plaintext recovery

This stitches together the separately proven v0.2.52 and v0.2.53 relay paths into one complete OpenMLS sidecar artifact lifecycle.

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
    package a deployable server release

## 5. Validation

Expected validation:

    go test -p 1 ./internal/protocol -run "TestOpenMLSSidecarFullLifecycleRelayThroughCypherEnvelope"
    go test -p 1 ./internal/protocol -run "TestOpenMLSSidecarApplicationMessageRelayThroughCypherEnvelope|TestOpenMLSSidecarKeyPackageWelcomeRelayThroughCypherEnvelope|TestOpenMLSSidecarFullLifecycleRelayThroughCypherEnvelope"
    go test -p 1 ./internal/relay
    go test -p 1 ./internal/protocol
    go test -p 1 ./...
    powershell -ExecutionPolicy Bypass -File .\scripts\check-no-rust-artifacts.ps1

## 6. Next rung

Next planned rung:

    v0.2.55 — deployable server + CLI smoke-test recon/planning.

Goal:

    determine the smallest honest server-deployable CarbonStack backbone proof:
    Cypher server can run repeatably
    Comms CLI or test harness can exercise the relay lifecycle
    docs clearly frame the result as experimental/pre-release
    no Android dependency
    no production security certification claim

A pre-v0.3.0 cleanup/release-hardening checkpoint should still occur before v0.3.0.
