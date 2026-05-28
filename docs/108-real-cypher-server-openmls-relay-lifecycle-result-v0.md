# Real Cypher Server OpenMLS Relay Lifecycle Result v0

Status: Integration proof result
Component: CarbonStackComms + CarbonStackCypher real server + OpenMLS sidecar
Phase: v0.2.56 real Cypher server + Comms client/relay smoke proof
Previous docs:
- docs/106-full-openmls-relay-lifecycle-through-cypher-result-v0.md
- docs/107-deployable-server-cli-smoke-test-recon-v0.md

## 1. Summary

This checkpoint records the first OpenMLS relay lifecycle proof through a real CarbonStackCypher server process.

v0.2.54 proved the full lifecycle through a protocol-local Cypher-compatible httptest server.

v0.2.56 moves that proof to a real carbonstack-cypher process started locally with explicit test configuration.

The proof validates:

    real carbonstack-cypher server startup
    temp SQLite database use
    real migrations use
    /v0/health readiness
    Comms client.New(realServerURL)
    real invite claim/device registration routes
    real /v0/envelopes submission
    real /v0/devices/{device_id}/envelopes retrieval
    KeyPackage relay
    Welcome relay
    application-message relay
    sidecar message-open plaintext recovery

This is still a dev/test smoke proof.

It is not polished Comms runtime UX.

It is not a release package.

It is not production E2EE.

## 2. What changed

CarbonStackComms added a protocol-level real-server smoke test.

The test starts a real CarbonStackCypher process using:

    go run ./cmd/cypher

with test-scoped environment:

    CYPHER_ADDR=127.0.0.1:<test-port>
    CYPHER_DB=<temp-dir>/cypher-real-server-smoke.db
    CYPHER_MIGRATIONS=<carbonstack-cypher>/migrations
    CYPHER_DEV_INVITE=dev-invite

The test waits for:

    GET /v0/health

Then it uses:

    internal/client.CypherClient
    internal/relay.SubmitOpenMLSArtifactEnvelope
    internal/relay.WriteOpenMLSArtifactFromEnvelope
    promoted OpenMLS sidecar helpers

to run the full relay lifecycle against the real server.

## 3. Validated lifecycle

Validated real-server lifecycle:

    Alice and Bob claim/register Cypher devices against real server
    Bob sidecar exports public-bundle.keypackage.bin
    KeyPackage artifact submits to real /v0/envelopes
    Alice retrieves KeyPackage from real inbox
    Alice sidecar consumes downloaded KeyPackage with conversation-add-member
    Alice sidecar writes welcome.bin
    Welcome artifact submits to real /v0/envelopes
    Bob retrieves Welcome from real inbox
    Bob sidecar consumes downloaded Welcome with conversation-join
    Alice sidecar writes application-message.bin
    application-message artifact submits to real /v0/envelopes
    Bob retrieves application-message from real inbox
    Bob sidecar consumes downloaded application-message with message-open
    plaintext matches

## 4. Preserved boundaries

This proof does not:

    wire comms send
    wire comms inbox
    create user-facing OpenMLS CLI UX
    automatically acknowledge after sidecar consume
    parse MLS internals in Cypher
    parse MLS internals in relay helper
    mutate trust-state
    relay signer.json
    relay provider-storage.json
    add Cypher routes
    add Cypher migrations
    claim production readiness
    package a deployable release

## 5. Validation

Expected validation:

    go test -p 1 ./internal/protocol -run "TestOpenMLSSidecarFullLifecycleRelayThroughRealCypherServer" -count=1 -timeout 180s
    go test -p 1 ./internal/protocol -run "TestOpenMLSSidecarFullLifecycleRelayThroughCypherEnvelope|TestOpenMLSSidecarFullLifecycleRelayThroughRealCypherServer" -count=1 -timeout 240s
    go test -p 1 ./internal/relay
    go test -p 1 ./internal/protocol -count=1 -timeout 300s
    go test -p 1 ./... -count=1 -timeout 300s
    powershell -ExecutionPolicy Bypass -File .\scripts\check-no-rust-artifacts.ps1

## 6. Next rung

Next planned rung:

    v0.2.57 — CLI/dev harness for repeatable local relay lifecycle.

Goal:

    turn the real-server proof into a repeatable dev harness shape that can be run intentionally,
    while still avoiding polished production UX claims.

The harness should remain experimental and should make generated DB/state/artifact hygiene obvious.
