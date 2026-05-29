# OpenMLS Backbone Self-Test Harness Result v0

Status: implementation result
Component: CarbonStackComms + CarbonStack
Phase: v0.2.66 OpenMLS backbone self-test harness implementation

## 1. Summary

CarbonStackComms now has a public-facing local self-test wrapper:

    scripts/self-test-openmls-backbone.ps1

The wrapper delegates to the existing lower-level real-Cypher smoke harness:

    scripts/smoke-openmls-real-cypher-relay.ps1

The wrapper exists to give builders a clearer validation entrypoint without pretending the project is a finished messenger.

## 2. What it proves

The self-test validates the current known-good local backbone:

    CarbonStackComms OpenMLS sidecar
    + CarbonStackCypher real local server
    + opaque OpenMLS artifact envelope relay
    + payload metadata validation
    + consume-then-ack semantics

It proves a repeatable local experimental backbone lifecycle.

## 3. What it does not prove

It does not prove:

- production readiness;
- production E2EE;
- hostile-server safety;
- metadata privacy;
- secure local vault/storage;
- Android readiness;
- polished runtime Comms UX;
- external audit or certification.

## 4. Commands

Targeted self-test:

    cd C:\Users\udaiv\repos\carbonstack_umbrella\carbonstack-comms
    powershell -ExecutionPolicy Bypass -File .\scripts\self-test-openmls-backbone.ps1

Full validation:

    powershell -ExecutionPolicy Bypass -File .\scripts\self-test-openmls-backbone.ps1 -Full

Lower-level smoke harness remains available:

    powershell -ExecutionPolicy Bypass -File .\scripts\smoke-openmls-real-cypher-relay.ps1

## 5. Boundary

Execution remains in `carbonstack-comms`.

The main `carbonstack` repo remains the public release/front-door documentation surface.

No top-level `carbonstack` orchestrator was added.

No Comms runtime `send` / `inbox` OpenMLS UX was added.

No new Cypher route was added.

No production or certification claim was added.
