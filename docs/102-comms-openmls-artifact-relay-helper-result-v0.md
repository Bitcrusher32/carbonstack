# Comms OpenMLS Artifact Relay Helper Result v0

Status: Implementation result
Component: CarbonStackComms
Phase: v0.2.50 Comms relay helper scaffold
Previous docs:
- docs/101-comms-openmls-cypher-relay-bridge-recon-v0.md

## 1. Summary

This checkpoint records the first CarbonStackComms relay helper scaffold for OpenMLS sidecar artifacts.

The implementation adds a narrow helper package:

    carbonstack-comms/internal/relay

The helper translates between local OpenMLS sidecar artifact files and Cypher-compatible opaque envelope payloads.

This is not polished runtime integration.

This is not `comms send` / `comms inbox` OpenMLS wiring.

This is not a full end-to-end Cypher/OpenMLS proof yet.

## 2. What changed

The Comms repo now has helper code for:

    OpenMLS artifact kind constants
    OpenMLS artifact kind -> Cypher content_type mapping
    CarbonStack OpenMLS sidecar protocol version constant
    reading artifact bytes from disk
    writing artifact bytes to disk
    base64 encoding artifact bytes
    base64 decoding envelope payload bytes
    read artifact -> base64 payload
    base64 payload -> artifact file

## 3. Content types

The helper aligns with Cypher v0.2.48 content types:

    carbonstack.mls.keypackage.v0
    carbonstack.mls.welcome.v0
    carbonstack.mls.application-message.v0

Protocol version:

    carbonstack-openmls-sidecar-v0

## 4. Current scope

This scaffold is deliberately local/pure.

It does not yet:

    invoke the OpenMLS sidecar
    call Cypher directly
    wire into Comms send/inbox
    acknowledge Cypher envelopes
    mutate trust-state
    parse MLS internals
    inspect plaintext
    route signer/provider storage
    claim production E2EE

## 5. Validation

The helper tests validate:

    keypackage content-type mapping
    welcome content-type mapping
    application-message content-type mapping
    unsupported artifact kind rejection
    arbitrary binary artifact read -> base64 -> decode -> write roundtrip
    invalid base64 rejection
    directory read rejection

Expected validation:

    go test -p 1 ./internal/relay
    go test -p 1 ./...
    powershell -ExecutionPolicy Bypass -File .\scripts\check-no-rust-artifacts.ps1

## 6. Safety boundary

The helper must not introduce relay kinds for:

    signer.json
    provider-storage.json
    raw MemoryStorage JSON
    raw OpenMLS group state
    plaintext
    private keys
    trust-state private material

Those remain local-only.

## 7. Next rung

Next planned rung:

    v0.2.51 — sidecar-produced application-message relay proof.

Goal:

    sidecar message-protect writes application-message.bin
    relay helper encodes artifact bytes
    Comms client submits bytes to Cypher /v0/envelopes
    recipient Comms client retrieves envelope
    relay helper decodes/writes application-message.bin
    sidecar message-open consumes written artifact
    plaintext matches

This should still be a dev/test proof, not polished CLI UX.

## 8. Later path

After application-message relay proof:

    v0.2.52 — KeyPackage + Welcome relay proof
    v0.2.53 — full KeyPackage -> Welcome -> application-message relay lifecycle
    v0.2.54+ — server/CLI deployability smoke tests and docs
    v0.3.0 — minor epoch for experimental server-deployable CarbonStack backbone

The v0.3.0 target remains experimental/pre-release and must not claim certified security, production E2EE, or external audit.
