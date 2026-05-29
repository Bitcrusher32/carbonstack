# Experimental Backbone Deployability Runbook v0

Status: current known-good runbook
Component: CarbonStack + CarbonStackComms + CarbonStackCypher
Phase: v0.2.61 public deployability runbook

## 1. What am I looking at?

This is the current CarbonStack experimental backbone runbook.

It is not a finished messenger.

It is not a production security product.

It is not Android-ready.

It is not externally audited or certified secure.

The current validated artifact is a local developer proof that CarbonStackComms can use OpenMLS sidecar artifacts and CarbonStackCypher envelope relay storage to move a complete OpenMLS bootstrap/message lifecycle through a real local Cypher server.

In plain terms:

1. Bob creates a KeyPackage.
2. Cypher relays it to Alice.
3. Alice creates a Welcome.
4. Cypher relays it to Bob.
5. Bob joins the conversation.
6. Alice creates an encrypted application-message artifact.
7. Cypher relays it to Bob.
8. Bob opens it through the OpenMLS sidecar.
9. The plaintext matches.
10. Envelopes are acknowledged only after sidecar consume succeeds.
11. Payload metadata is checked before artifact bytes are written locally.

This is the current experimental backbone.

## 2. Component repositories

CarbonStack currently spans these repositories:

- `carbonstack`
  - doctrine, specs, docs, release framing, runbooks, and public front-door material.

- `carbonstack-comms`
  - text-first Comms client implementation, OpenMLS sidecar, relay helper, protocol tests, and smoke harness.

- `carbonstack-cypher`
  - self-hostable relay/storage server skeleton, envelope API, SQLite schema, migrations, and server tests.

- `carbonstack-os`
  - future constrained appliance OS north star. It is not part of the current runnable relay proof.

The planned v0.3.0 prerelease should be framed from the `carbonstack` repo. Component repos should remain implementation and development surfaces.

## 3. Local checkout shape

The known-good local commands assume sibling repositories under one umbrella directory:

    C:\Users\udaiv\repos\carbonstack_umbrella\carbonstack
    C:\Users\udaiv\repos\carbonstack_umbrella\carbonstack-comms
    C:\Users\udaiv\repos\carbonstack_umbrella\carbonstack-cypher
    C:\Users\udaiv\repos\carbonstack_umbrella\carbonstack-os

## 4. Known-good validation commands

Run Cypher tests:

    cd C:\Users\udaiv\repos\carbonstack_umbrella\carbonstack-cypher
    go test ./... -count=1

Run the Comms real-server smoke harness:

    cd C:\Users\udaiv\repos\carbonstack_umbrella\carbonstack-comms
    Get-Process cypher -ErrorAction SilentlyContinue | Stop-Process -Force
    powershell -ExecutionPolicy Bypass -File .\scripts\smoke-openmls-real-cypher-relay.ps1

Run broader Comms validation:

    powershell -ExecutionPolicy Bypass -File .\scripts\smoke-openmls-real-cypher-relay.ps1 -Full
    powershell -ExecutionPolicy Bypass -File .\scripts\check-no-rust-artifacts.ps1

Run CarbonStack docs validation:

    cd C:\Users\udaiv\repos\carbonstack_umbrella\carbonstack
    powershell -ExecutionPolicy Bypass -File .\scripts\validate-local.ps1

## 5. What the smoke harness does

The Comms smoke harness runs the real-server OpenMLS relay lifecycle proof.

It:

- builds a temporary `carbonstack-cypher` test server path;
- starts a real local Cypher server;
- uses a temp SQLite database;
- waits for `/v0/health`;
- uses the Comms Cypher client;
- registers or uses Cypher devices;
- runs OpenMLS KeyPackage -> Welcome -> application-message artifact relay;
- validates payload metadata;
- consumes artifacts through the OpenMLS sidecar;
- acknowledges envelopes only after successful sidecar consume;
- checks that generated Rust/OpenMLS artifacts are not accidentally committed.

## 6. Current relay contract

Cypher currently accepts these OpenMLS relay content types:

    carbonstack.mls.keypackage.v0
    carbonstack.mls.welcome.v0
    carbonstack.mls.application-message.v0

The current OpenMLS relay protocol version is:

    carbonstack-openmls-sidecar-v0

The existing stub content type remains:

    carbonstack.message.text.stub.v0

## 7. Payload metadata

Cypher computes and stores:

    payload_size_bytes
    payload_sha256

Both fields describe decoded `ciphertext_b64` payload bytes.

Comms validates these metadata fields before writing downloaded artifact bytes to disk for sidecar consumption.

This metadata is useful for relay, debugging, and storage sanity.

It is not a cryptographic trust root.

A malicious server can lie about server-returned metadata. The OpenMLS sidecar consume step remains the cryptographic validity gate.

## 8. Ack semantics

The current ack rule is:

    do not ack on download;
    do not ack on artifact write;
    ack only after sidecar consume succeeds.

Validated boundaries:

- KeyPackage:
  - ack only after `conversation-add-member` consumes the downloaded KeyPackage.

- Welcome:
  - ack only after `conversation-join` consumes the downloaded Welcome.

- application-message:
  - ack only after `message-open` consumes the downloaded application-message and plaintext recovery is validated.

## 9. Generated-state warnings

Do not commit generated sidecar state.

Do not commit:

- `.carbonstack-openmls-sidecar-state/`
- `target/`
- `cypher.db`
- `*.db`
- `signer.json`
- `provider-storage.json`
- `welcome.bin`
- `application-message.bin`
- `public-bundle.keypackage.bin`
- raw OpenMLS group/provider state
- private keys
- trust-state private material

On Windows, stale test `cypher.exe` processes can hold temp SQLite files open.

Inspect stale processes:

    Get-Process cypher -ErrorAction SilentlyContinue | Select-Object Id, ProcessName, Path

Stop stale test processes when no intentional Cypher server is running:

    Get-Process cypher -ErrorAction SilentlyContinue | Stop-Process -Force

## 10. What this currently proves

The current backbone proves:

- OpenMLS artifacts can be generated by the promoted sidecar.
- OpenMLS artifacts can be relayed through real Cypher envelope storage.
- Comms relay helper can submit, retrieve, and write sidecar-compatible artifacts.
- Cypher can store OpenMLS artifact envelopes with payload metadata.
- Comms can validate payload metadata before writing artifacts.
- The recipient sidecar can consume relayed artifacts.
- Envelopes can be acknowledged after successful sidecar consume.
- The local dev smoke harness can repeat this path.

## 11. What this does not prove

This does not prove:

- production E2EE;
- hostile-server safety;
- metadata privacy;
- rollback/replay safety against a malicious server;
- secure local vault/storage;
- Android appliance readiness;
- polished Comms CLI UX;
- multi-user production operations;
- external audit or certification;
- a stable public protocol.

## 12. Next planned rungs

Near-term:

- public README/surface cleanup: complete at v0.2.62;
- known-good validation matrix: complete at v0.2.63;
- inbox/ack/general semantics and schema standardization;
- Option B CLI/dev-harness planning and implementation;
- pre-v0.3.0 release hardening.

Future v0.3.0 framing:

    CarbonStack experimental backbone.

Concrete validated artifact:

    Cypher + Comms OpenMLS relay backbone.

