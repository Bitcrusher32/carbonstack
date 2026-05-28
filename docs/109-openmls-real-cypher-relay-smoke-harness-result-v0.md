# OpenMLS Real Cypher Relay Smoke Harness Result v0

Status: Dev harness result
Component: CarbonStackComms + CarbonStackCypher real server + OpenMLS sidecar
Phase: v0.2.57 CLI/dev harness for repeatable local relay lifecycle
Previous docs:
- docs/107-deployable-server-cli-smoke-test-recon-v0.md
- docs/108-real-cypher-server-openmls-relay-lifecycle-result-v0.md

## 1. Summary

This checkpoint records the first repeatable dev smoke harness for the OpenMLS real-Cypher relay lifecycle.

v0.2.56 proved the full OpenMLS relay lifecycle through a real CarbonStackCypher server process from a Go integration test.

v0.2.57 adds a script-level harness that intentionally runs that proof as a repeatable local smoke test.

The harness is developer-facing.

It is not polished Comms runtime UX.

It is not production E2EE.

It is not externally audited or certified secure.

## 2. What changed

CarbonStackComms now includes:

    scripts/smoke-openmls-real-cypher-relay.ps1

The script runs:

    TestOpenMLSSidecarFullLifecycleRelayThroughRealCypherServer

and then runs the generated Rust/OpenMLS artifact guard.

With `-Full`, the script also runs broader protocol/relay validation.

## 3. Validated harness behavior

The harness validates the existing real-server proof path:

    builds a temp carbonstack-cypher test binary
    starts real CarbonStackCypher on localhost
    uses temp SQLite DB
    waits for /v0/health
    uses Comms client/relay helpers
    runs KeyPackage -> Welcome -> application-message relay lifecycle
    verifies final sidecar message-open plaintext recovery
    checks for generated Rust/OpenMLS artifacts

## 4. Safety behavior

The harness refuses to continue if existing `cypher` processes are detected before the run.

This is intended to avoid false-positive test results and avoid stale orphan processes holding SQLite DB files open on Windows.

## 5. Preserved boundaries

This harness does not:

    wire comms send
    wire comms inbox
    create polished OpenMLS user CLI UX
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

## 6. Validation

Expected validation:

    powershell -ExecutionPolicy Bypass -File .\scripts\smoke-openmls-real-cypher-relay.ps1
    powershell -ExecutionPolicy Bypass -File .\scripts\smoke-openmls-real-cypher-relay.ps1 -Full

## 7. Next rung

Next planned rung:

    v0.2.58 — ack semantics after successful sidecar consume.

Goal:

    decide and prove the first safe ack boundary:
    acknowledge Cypher envelopes only after the recipient writes the artifact and the sidecar successfully consumes it.

This should remain test/harness scoped before user-facing runtime UX.
