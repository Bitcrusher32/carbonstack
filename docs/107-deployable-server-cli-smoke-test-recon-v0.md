# Deployable Server + CLI Smoke-Test Recon v0

Status: Recon / implementation planning
Component: CarbonStackCypher + CarbonStackComms + OpenMLS sidecar
Phase: v0.2.55 deployable server + CLI smoke-test recon/planning
Previous docs:
- docs/104-openmls-application-message-relay-through-cypher-result-v0.md
- docs/105-openmls-keypackage-welcome-relay-through-cypher-result-v0.md
- docs/106-full-openmls-relay-lifecycle-through-cypher-result-v0.md

## 1. Purpose

This document records the transition plan from protocol-local Cypher-compatible httptest relay proofs to a real CarbonStackCypher server process and repeatable Comms-side smoke testing.

v0.2.54 proved the full OpenMLS artifact lifecycle through a Cypher-compatible envelope server:

    KeyPackage relay
    Welcome relay
    application-message relay
    message-open plaintext recovery

v0.2.55 decides how to move toward a server-deployable experimental backbone without prematurely claiming production readiness.

## 2. Current proven state

The following are already proven at dev/test level:

    OpenMLS sidecar can produce public-bundle.keypackage.bin.
    OpenMLS sidecar can consume a KeyPackage artifact with conversation-add-member.
    OpenMLS sidecar can produce welcome.bin.
    OpenMLS sidecar can consume a Welcome artifact with conversation-join.
    OpenMLS sidecar can produce application-message.bin.
    OpenMLS sidecar can consume an application-message artifact with message-open.
    Comms internal/relay can read artifact bytes and submit them as Cypher envelope payloads.
    Comms internal/relay can decode Cypher envelope payloads and write sidecar-compatible artifact files.
    Cypher accepts OpenMLS artifact content types through /v0/envelopes.
    The full KeyPackage -> Welcome -> application-message lifecycle works through a Cypher-compatible httptest server.

## 3. Cypher real-server startup recon

CarbonStackCypher has a real server entry point:

    cmd/cypher/main.go

Server configuration is environment-driven through:

    internal/config/config.go

Current config keys:

    CYPHER_ADDR
    CYPHER_DB
    CYPHER_MIGRATIONS
    CYPHER_DEV_INVITE

Defaults:

    CYPHER_ADDR=:8080
    CYPHER_DB=cypher.db
    CYPHER_MIGRATIONS=migrations
    CYPHER_DEV_INVITE=dev-invite

Startup behavior:

    load config
    open SQLite database
    run migrations
    seed dev invite if configured
    create httpapi.New(...)
    listen with http.ListenAndServe

For test/smoke harness use, the next rung should override:

    CYPHER_ADDR=127.0.0.1:<test-port>
    CYPHER_DB=<temp-dir>/cypher-smoke.db
    CYPHER_MIGRATIONS=<repo-root>/migrations
    CYPHER_DEV_INVITE=dev-invite

## 4. Cypher route/API recon

The real server exposes the same route family used by the current proofs:

    GET /v0/health
    POST /v0/dev/invites
    POST /v0/invites/claim
    POST /v0/devices/register
    GET /v0/accounts/{account_id}/devices
    POST /v0/envelopes
    GET /v0/devices/{device_id}/envelopes
    POST /v0/envelopes/{envelope_id}/ack

The OpenMLS content types accepted by Cypher are:

    carbonstack.mls.keypackage.v0
    carbonstack.mls.welcome.v0
    carbonstack.mls.application-message.v0

The OpenMLS protocol version is:

    carbonstack-openmls-sidecar-v0

## 5. Comms client recon

CarbonStackComms already has a client package:

    internal/client/cypher.go

It exposes:

    client.New(serverURL)
    CreateDevInvite(...)
    ClaimInvite(...)
    RegisterDevice(...)
    ListDevices(...)
    SubmitEnvelope(...)
    Inbox(...)
    AckEnvelope(...)

This is enough for a real-server smoke proof without user-facing CLI changes.

## 6. Comms CLI recon

Current Comms CLI commands still exist in:

    internal/app/commands.go

Current `send`, `inbox`, and `ack` commands are stub-era and use mock/stub text-envelope behavior.

They should not be wired to OpenMLS relay yet.

The next implementation should avoid polished runtime UX and use a test/harness path first.

## 7. Recommended v0.2.56 implementation shape

v0.2.56 should prove the full OpenMLS relay lifecycle against a real CarbonStackCypher server process.

Recommended location:

    carbonstack-comms/internal/protocol

Reason:

    the test still depends heavily on OpenMLS sidecar lifecycle helpers;
    internal/relay remains generic and already proven;
    internal/app should not be touched until runtime UX is intentionally designed.

High-level flow:

    start real carbonstack-cypher process on a local test port
    use a temp SQLite database
    use the real migrations directory
    wait for GET /v0/health
    create Comms client with client.New(realServerURL)
    claim/register Alice and Bob Cypher devices
    run OpenMLS sidecar KeyPackage -> Welcome -> application-message lifecycle
    submit/retrieve artifacts through the real Cypher server
    open the final message with sidecar message-open
    assert plaintext recovery

## 8. Generated artifact hygiene

The real-server smoke proof must not commit:

    cypher.db
    sqlite databases
    .carbonstack-openmls-sidecar-state/
    target/
    signer.json
    provider-storage.json
    welcome.bin
    application-message.bin
    public-bundle.keypackage.bin
    any generated OpenMLS state

The implementation should use:

    t.TempDir()
    explicit CYPHER_DB in temp dir
    existing rust-artifact guard
    existing git status suspicious-file checks

## 9. Option C -> Option B deployability ladder

The agreed near-term deployability path is C -> B.

Option C first:

    test harness proves real-server relay lifecycle
    docs mirror the future CLI/runbook
    correctness and repeatability come before UX

Option B later:

    CLI/dev harness becomes visible and self-testable
    builders can run a release-facing flow without reading test internals

Planned ladder:

    v0.2.56  Real Cypher server + Comms client/relay smoke proof
    v0.2.57  CLI/dev harness for repeatable local relay lifecycle
    v0.2.58  Ack semantics after successful sidecar consume
    v0.2.59  Payload metadata/hash/size planning or migration
    v0.2.60  Deployability docs/runbook + known-good validation
    v0.2.61  Option C ironed out and completed for testing
    v0.2.62  Stale helper cleanup, inbox/ack/general semantics/schema standardization
    v0.2.63+ Option B planning + implementation
    v0.2.x   Pre-v0.3.0 cleanup/release-hardening checkpoint
    v0.3.0   Experimental server-deployable CarbonStack backbone epoch

## 10. Explicit non-goals

v0.2.55 and v0.2.56 do not claim:

    production E2EE
    hostile-server proof
    metadata privacy completion
    Android readiness
    external audit
    certified secure status
    stable public protocol
    polished user runtime UX

## 11. Next implementation candidate

Next rung:

    v0.2.56 — real Cypher server + Comms client/relay smoke proof.

Implementation should add a test/harness that starts the real Cypher server process and runs the full OpenMLS relay lifecycle through it.

Do not wire `comms send` / `comms inbox` yet.
Do not add automatic ack yet.
Do not add payload metadata migration yet.
Do not package a release yet.
