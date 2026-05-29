# Known-Good Validation Matrix v0

Status: current validation matrix
Component: CarbonStack + CarbonStackComms + CarbonStackCypher
Phase: v0.2.63 Option C known-good validation cleanup

Related current docs:

- `docs/113-experimental-backbone-deployability-runbook-v0.md`
- `docs/115-envelope-lifecycle-semantics-v0.md`
- `docs/116-openmls-backbone-self-test-harness-plan-v0.md`

## 1. Purpose

This document defines the current known-good validation path for the CarbonStack experimental backbone.

It does not define a production deployment.

It does not define a stable public protocol.

It does not prove production security.

The current execution entrypoint remains in `carbonstack-comms`.

The public release explanation remains in `carbonstack`.

## 2. Current known-good proof

The current known-good proof is:

    CarbonStackComms OpenMLS sidecar
    + CarbonStackCypher real local server
    + opaque OpenMLS artifact envelope relay
    + payload metadata validation
    + consume-then-ack semantics

The proof validates:

1. Bob exports an OpenMLS KeyPackage.
2. Cypher relays the KeyPackage to Alice.
3. Alice consumes the KeyPackage and creates a Welcome.
4. Cypher relays the Welcome to Bob.
5. Bob consumes the Welcome and joins the conversation.
6. Alice protects an application-message.
7. Cypher relays the application-message to Bob.
8. Bob validates payload metadata before artifact write.
9. Bob consumes the application-message through the OpenMLS sidecar.
10. The plaintext matches.
11. Envelopes are acknowledged only after successful sidecar consume.

## 3. Primary validation commands

Run Cypher tests:

    cd C:\Users\udaiv\repos\carbonstack_umbrella\carbonstack-cypher
    go test ./... -count=1

Run the current Comms-hosted backbone smoke harness:

    cd C:\Users\udaiv\repos\carbonstack_umbrella\carbonstack-comms
    Get-Process cypher -ErrorAction SilentlyContinue | Stop-Process -Force
    powershell -ExecutionPolicy Bypass -File .\scripts\self-test-openmls-backbone.ps1

Run full Comms validation:

    powershell -ExecutionPolicy Bypass -File .\scripts\self-test-openmls-backbone.ps1 -Full
    powershell -ExecutionPolicy Bypass -File .\scripts\check-no-rust-artifacts.ps1

Run CarbonStack docs validation:

    cd C:\Users\udaiv\repos\carbonstack_umbrella\carbonstack
    powershell -ExecutionPolicy Bypass -File .\scripts\validate-local.ps1

## 4. Validation layers

### Cypher layer

Command:

    go test ./... -count=1

Validates:

- server package compilation;
- HTTP API tests;
- envelope lifecycle behavior;
- OpenMLS content-type acceptance;
- payload metadata implementation;
- SQLite migration compatibility.

### Comms smoke layer

Command:

    powershell -ExecutionPolicy Bypass -File .\scripts\self-test-openmls-backbone.ps1

Validates:

- real local Cypher server startup;
- temp SQLite database use;
- real migration use;
- OpenMLS KeyPackage relay;
- OpenMLS Welcome relay;
- OpenMLS application-message relay;
- payload metadata validation;
- message-open plaintext recovery;
- consume-then-ack behavior;
- generated Rust/OpenMLS artifact guard.

### Comms full layer

Command:

    powershell -ExecutionPolicy Bypass -File .\scripts\self-test-openmls-backbone.ps1 -Full

Validates the smoke layer plus broader protocol, relay, and Go package tests.

Use this before pushing protocol, relay, client, sidecar, or validation-script changes.

### CarbonStack docs layer

Command:

    powershell -ExecutionPolicy Bypass -File .\scripts\validate-local.ps1

Validates the main docs/release-surface repository.

## 5. Current execution boundary

Execution remains in `carbonstack-comms`.

`carbonstack` documents the release surface, component map, runbooks, and known-good validation path.

`carbonstack-cypher` supplies the known-good relay/server component.

Do not add a top-level `carbonstack` orchestrator unless the pre-v0.3.0 hardening pass shows it can remain simple.

## 6. What this proves

This proves a repeatable local experimental backbone lifecycle.

It proves that Comms and Cypher can agree on the current OpenMLS artifact relay path.

It proves that payload metadata and consume-then-ack behavior are included in the current known-good path.

## 7. What this does not prove

This does not prove:

- production E2EE;
- hostile-server safety;
- metadata privacy;
- secure local vault/storage;
- Android readiness;
- polished runtime Comms UX;
- stable public protocol status;
- external audit or certification.

## 8. Next rung

The next rung should clean up inbox, ack, general envelope semantics, and schema/API wording.

The goal is to make the relay lifecycle language consistent before Option B CLI/dev-harness planning and pre-v0.3.0 release hardening.



