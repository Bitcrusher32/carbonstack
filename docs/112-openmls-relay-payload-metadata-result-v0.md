# OpenMLS Relay Payload Metadata Result v0

Status: Implementation result
Component: CarbonStackCypher + CarbonStackComms relay helper + OpenMLS sidecar
Phase: v0.2.60 payload metadata implementation
Previous docs:
- docs/110-openmls-relay-ack-semantics-result-v0.md
- docs/111-openmls-relay-payload-metadata-plan-v0.md

## 1. Summary

This checkpoint records implementation of payload metadata for Cypher envelopes carrying OpenMLS sidecar artifacts.

The implemented metadata fields are:

    payload_size_bytes
    payload_sha256

Both fields describe decoded `ciphertext_b64` payload bytes.

This improves relay/debug/storage sanity checks for OpenMLS artifact transport.

It is not a production security claim and not a substitute for OpenMLS sidecar validation.

## 2. Cypher behavior

CarbonStackCypher now computes payload metadata server-side during envelope submission.

The server:

    decodes ciphertext_b64;
    rejects invalid base64 as before;
    computes payload_size_bytes from decoded payload length;
    computes payload_sha256 as lowercase SHA-256 hex over decoded payload bytes;
    stores metadata with the envelope;
    returns metadata in submit response;
    returns metadata in inbox response.

The envelope route remains:

    POST /v0/envelopes
    GET /v0/devices/{device_id}/envelopes
    POST /v0/envelopes/{envelope_id}/ack

No parallel `/v0/artifacts` route was added.

## 3. Comms behavior

CarbonStackComms client structs now carry payload metadata from Cypher responses.

The relay helper validates metadata before writing sidecar-compatible artifact bytes.

The validation rule is:

    decode CiphertextB64;
    compare decoded length against payload_size_bytes when present;
    compare SHA-256(decoded bytes) against payload_sha256 when present;
    fail before writing artifact if metadata mismatches;
    write artifact only after metadata validation succeeds.

## 4. OpenMLS relay behavior

The real-server OpenMLS relay proof now validates payload metadata for:

    KeyPackage envelopes
    Welcome envelopes
    application-message envelopes

The sidecar consume boundary remains:

    KeyPackage is consumed by conversation-add-member
    Welcome is consumed by conversation-join
    application-message is consumed by message-open

Ack remains consume-then-ack from v0.2.58.

## 5. Security boundary

Payload metadata is relay/storage/debug metadata.

It does not prove:

    production E2EE
    hostile-server safety
    metadata privacy
    OpenMLS semantic validity
    payload authenticity against a malicious server
    Android readiness
    external audit or certification

The cryptographic validation boundary remains OpenMLS sidecar processing.

## 6. Validation

Expected validation:

    carbonstack-cypher:
        go test ./internal/httpapi -count=1
        go test ./... -count=1

    carbonstack-comms:
        go test -p 1 ./internal/relay -count=1
        go test -p 1 ./internal/protocol -run "TestOpenMLSSidecarFullLifecycleRelayThroughRealCypherServer" -count=1 -timeout 240s
        powershell -ExecutionPolicy Bypass -File .\scripts\smoke-openmls-real-cypher-relay.ps1
        powershell -ExecutionPolicy Bypass -File .\scripts\smoke-openmls-real-cypher-relay.ps1 -Full
        powershell -ExecutionPolicy Bypass -File .\scripts\check-no-rust-artifacts.ps1

## 7. Next rung

Next planned rung:

    deployability docs/runbook + known-good validation.

Goal:

    document the local experimental backbone path from the carbonstack front-door repo,
    with clear prerelease framing and no production/certification claims.
