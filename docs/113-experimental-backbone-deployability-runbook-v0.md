# Experimental Backbone Deployability Runbook v0

Status: Deployability runbook / known-good validation
Component: CarbonStack front-door docs + CarbonStackComms + CarbonStackCypher
Phase: v0.2.61 deployability docs/runbook + known-good validation
Previous docs:
- docs/107-deployable-server-cli-smoke-test-recon-v0.md
- docs/108-real-cypher-server-openmls-relay-lifecycle-result-v0.md
- docs/109-openmls-real-cypher-relay-smoke-harness-result-v0.md
- docs/110-openmls-relay-ack-semantics-result-v0.md
- docs/111-openmls-relay-payload-metadata-plan-v0.md
- docs/112-openmls-relay-payload-metadata-result-v0.md

## 1. What am I looking at?

This is an experimental CarbonStack backbone runbook.

It is not a finished messenger.

It is not a production security product.

It is not Android-ready.

It is not externally audited or certified secure.

The current validated artifact is a local developer proof that CarbonStackComms can use OpenMLS sidecar artifacts and CarbonStackCypher envelope relay storage to move a complete OpenMLS conversation bootstrap/message lifecycle through a real local Cypher server.

In plain terms:

    Bob creates a KeyPackage.
    Cypher relays it to Alice.
    Alice creates a Welcome.
    Cypher relays it to Bob.
    Bob joins the conversation.
    Alice creates an encrypted application-message artifact.
    Cypher relays it to Bob.
    Bob opens it through the OpenMLS sidecar.
    The plaintext matches.
    Envelopes are acknowledged only after sidecar consume succeeds.
    Payload metadata is checked before artifact bytes are written locally.

This is the current experimental backbone.

## 2. Component repos

The current system spans these repos:

    carbonstack
        doctrine, specs, docs, release framing, runbooks, and public front-door material.

    carbonstack-comms
        text-first Comms client implementation, OpenMLS sidecar, relay helper, protocol tests, and smoke harness.

    carbonstack-cypher
        self-hostable relay/storage server skeleton, envelope API, SQLite schema, migrations, and server tests.

    carbonstack-os
        future constrained appliance OS north star; not part of the current runnable relay proof.

The planned v0.3.0 prerelease should be framed from the `carbonstack` repo, not from a component repo alone.

## 3. Current known-good local validation

From an umbrella checkout containing sibling repos:

    C:\Users\udaiv\repos\carbonstack_umbrella\carbonstack
    C:\Users\udaiv\repos\carbonstack_umbrella\carbonstack-comms
    C:\Users\udaiv\repos\carbonstack_umbrella\carbonstack-cypher
    C:\Users\udaiv\repos\carbonstack_umbrella\carbonstack-os

Run Cypher tests:

    cd C:\Users\udaiv\repos\carbonstack_umbrella\carbonstack-cypher
    go test ./... -count=1

Run Comms real-server smoke harness:

    cd C:\Users\udaiv\repos\carbonstack_umbrella\carbonstack-comms
    Get-Process cypher -ErrorAction SilentlyContinue | Stop-Process -Force
    powershell -ExecutionPolicy Bypass -File .\scripts\smoke-openmls-real-cypher-relay.ps1

Run broader Comms validation:

    powershell -ExecutionPolicy Bypass -File .\scripts\smoke-openmls-real-cypher-relay.ps1 -Full
    powershell -ExecutionPolicy Bypass -File .\scripts\check-no-rust-artifacts.ps1

Run CarbonStack docs validation:

    cd C:\Users\udaiv\repos\carbonstack_umbrella\carbonstack
    powershell -ExecutionPolicy Bypass -File .\scripts\validate-local.ps1

## 4. What the smoke harness does

The Comms smoke harness runs the real-server OpenMLS relay lifecycle proof.

It:

    builds a temporary carbonstack-cypher test binary;
    starts a real local Cypher server;
    uses a temp SQLite database;
    waits for /v0/health;
    uses the Comms Cypher client;
    registers/uses Cypher devices;
    runs OpenMLS KeyPackage -> Welcome -> application-message artifact relay;
    validates payload metadata;
    consumes artifacts through the OpenMLS sidecar;
    acknowledges envelopes only after successful sidecar consume;
    checks that generated Rust/OpenMLS artifacts are not accidentally committed.

## 5. Current relay content types

Cypher currently accepts these OpenMLS relay content types:

    carbonstack.mls.keypackage.v0
    carbonstack.mls.welcome.v0
    carbonstack.mls.application-message.v0

The current OpenMLS relay protocol version is:

    carbonstack-openmls-sidecar-v0

The existing stub content type remains:

    carbonstack.message.text.stub.v0

## 6. Current payload metadata

Cypher computes and stores:

    payload_size_bytes
    payload_sha256

Both describe decoded `ciphertext_b64` payload bytes.

Comms validates these metadata fields before writing downloaded artifact bytes to disk for sidecar consumption.

This metadata is useful for relay/debug/storage sanity.

It is not a cryptographic trust root.

The OpenMLS sidecar consume step remains the cryptographic validity gate.

## 7. Current ack semantics

The current safe ack rule is:

    do not ack on download;
    do not ack on artifact write;
    ack only after sidecar consume succeeds.

Current validated boundaries:

    KeyPackage:
        ack only after conversation-add-member consumes the downloaded KeyPackage.

    Welcome:
        ack only after conversation-join consumes the downloaded Welcome.

    application-message:
        ack only after message-open consumes the downloaded application-message and plaintext recovery is validated.

## 8. Generated state and cleanup warnings

Do not commit generated sidecar state.

Do not commit:

    .carbonstack-openmls-sidecar-state/
    target/
    cypher.db
    *.db
    signer.json
    provider-storage.json
    welcome.bin
    application-message.bin
    public-bundle.keypackage.bin
    raw OpenMLS group/provider state
    private keys
    trust-state private material

On Windows, stale test `cypher.exe` processes can hold temp SQLite files open.

Before rerunning real-server smoke tests, it is safe to inspect:

    Get-Process cypher -ErrorAction SilentlyContinue | Select-Object Id, ProcessName, Path

If stale test processes are present and no intentional Cypher server is running, stop them:

    Get-Process cypher -ErrorAction SilentlyContinue | Stop-Process -Force

## 9. What this currently proves

The current backbone proves:

    OpenMLS artifacts can be generated by the promoted sidecar.
    OpenMLS artifacts can be relayed through real Cypher envelope storage.
    Comms relay helper can submit/retrieve/write sidecar-compatible artifacts.
    Cypher can store OpenMLS artifact envelopes with payload metadata.
    Comms can validate payload metadata before writing artifacts.
    The recipient sidecar can consume relayed artifacts.
    Envelopes can be acknowledged after successful sidecar consume.
    The local dev smoke harness can repeat this path.

## 10. What this does not prove

This does not prove:

    production E2EE product readiness;
    hostile-server safety;
    metadata privacy;
    rollback/replay safety against a malicious server;
    secure local vault/storage;
    Android appliance readiness;
    polished Comms CLI UX;
    multi-user production operations;
    external audit or certification;
    resistance to all OpenMLS misuse cases;
    a stable public protocol.

## 11. Next planned rungs

Near-term:

    full-spectrum README recon;
    critical README updates only where stale or misleading;
    Option C test-harness path completion;
    inbox/ack/general semantics and schema standardization;
    Option B CLI/dev-harness planning and implementation;
    pre-v0.3.0 release hardening.

Future v0.3.0 framing:

    CarbonStack experimental backbone.

The concrete validated artifact should be described as:

    Cypher + Comms OpenMLS relay backbone.

The main public release surface should live in the `carbonstack` repo, with links to component repos.
