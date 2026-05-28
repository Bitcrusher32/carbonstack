# OpenMLS Relay Ack Semantics Result v0

Status: Integration proof result
Component: CarbonStackComms + CarbonStackCypher real server + OpenMLS sidecar
Phase: v0.2.58 ack semantics after successful sidecar consume
Previous docs:
- docs/108-real-cypher-server-openmls-relay-lifecycle-result-v0.md
- docs/109-openmls-real-cypher-relay-smoke-harness-result-v0.md

## 1. Summary

This checkpoint records the first safe acknowledgement boundary for OpenMLS relay envelopes.

The rule proven in this checkpoint is:

    do not acknowledge a Cypher envelope when it is merely downloaded;
    do not acknowledge a Cypher envelope when it is merely written to disk;
    acknowledge only after the recipient sidecar successfully consumes the artifact.

This prevents the relay path from marking an envelope handled before the local OpenMLS sidecar has actually accepted it.

## 2. Ack boundary

The validated ack boundaries are:

    KeyPackage envelope:
        ack only after Alice conversation-add-member successfully consumes the downloaded KeyPackage artifact.

    Welcome envelope:
        ack only after Bob conversation-join successfully consumes the downloaded Welcome artifact.

    application-message envelope:
        ack only after Bob message-open successfully consumes the downloaded application-message artifact and plaintext recovery is validated.

## 3. What changed

CarbonStackComms real-server relay lifecycle test now calls:

    CypherClient.AckEnvelope(...)

after successful sidecar consume for each relay artifact.

The test then verifies the relevant recipient inbox is empty after ack.

## 4. Validated behavior

The real-server lifecycle now validates:

    Bob KeyPackage relays to Alice.
    Alice consumes KeyPackage with conversation-add-member.
    Alice acks the KeyPackage envelope.
    Alice inbox becomes empty.

    Alice Welcome relays to Bob.
    Bob consumes Welcome with conversation-join.
    Bob acks the Welcome envelope.
    Bob inbox becomes empty.

    Alice application-message relays to Bob.
    Bob consumes application-message with message-open.
    Bob acks the application-message envelope.
    Bob inbox becomes empty.

## 5. Preserved boundaries

This proof does not:

    wire comms send
    wire comms inbox
    create polished OpenMLS user CLI UX
    parse MLS internals in Cypher
    parse MLS internals in relay helper
    mutate trust-state
    relay signer.json
    relay provider-storage.json
    add Cypher routes
    add Cypher migrations
    add payload hash/size metadata
    claim production readiness
    package a deployable release

## 6. Validation

Expected validation:

    go test -p 1 ./internal/protocol -run "TestOpenMLSSidecarFullLifecycleRelayThroughRealCypherServer" -count=1 -timeout 240s
    powershell -ExecutionPolicy Bypass -File .\scripts\smoke-openmls-real-cypher-relay.ps1
    powershell -ExecutionPolicy Bypass -File .\scripts\smoke-openmls-real-cypher-relay.ps1 -Full
    powershell -ExecutionPolicy Bypass -File .\scripts\check-no-rust-artifacts.ps1

## 7. Next rung

Next planned rung:

    v0.2.59 — payload metadata/hash/size planning or migration.

Goal:

    decide whether to keep using ciphertext_b64 as the only payload field for the experimental backbone,
    or add payload_sha256 / payload_size_bytes metadata for safer relay validation and deployability.

This should be scoped carefully because it may require a Cypher DB migration and API/test updates.
