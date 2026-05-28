# Comms Cypher OpenMLS Application Artifact Bridge Helper Result v0

Status: Implementation result
Component: CarbonStackComms
Phase: v0.2.51 application-message relay proof preparation
Previous docs:
- docs/101-comms-openmls-cypher-relay-bridge-recon-v0.md
- docs/102-comms-openmls-artifact-relay-helper-result-v0.md

## 1. Summary

This checkpoint records the next Comms relay helper step after the initial OpenMLS artifact helper scaffold.

The relay package now has bridge helpers for submitting a local OpenMLS artifact file through the existing Cypher envelope client and writing a Cypher envelope payload back to a sidecar-compatible artifact file.

This is still helper/test infrastructure.

It does not wire polished `comms send` or `comms inbox`.

It does not automatically invoke the OpenMLS sidecar.

It does not automatically acknowledge Cypher envelopes after sidecar consumption.

## 2. What changed

The Comms relay package now supports:

    local artifact path -> base64 payload -> client.CypherClient.SubmitEnvelope
    client.EnvelopeRecord.CiphertextB64 -> decoded bytes -> local artifact path

The helper uses the existing content type/protocol contract:

    carbonstack.mls.keypackage.v0
    carbonstack.mls.welcome.v0
    carbonstack.mls.application-message.v0
    carbonstack-openmls-sidecar-v0

## 3. Validation

Tests prove:

    WriteOpenMLSArtifactFromEnvelope writes exact decoded bytes to disk.
    SubmitOpenMLSArtifactEnvelope builds the expected Cypher /v0/envelopes request.
    SubmitOpenMLSArtifactEnvelope uses the correct sender/recipient device IDs.
    SubmitOpenMLSArtifactEnvelope uses the correct OpenMLS content_type.
    SubmitOpenMLSArtifactEnvelope uses carbonstack-openmls-sidecar-v0.
    SubmitOpenMLSArtifactEnvelope rejects unsupported artifact kinds such as signer.json.

Expected validation:

    go test -p 1 ./internal/relay
    go test -p 1 ./...
    powershell -ExecutionPolicy Bypass -File .\scripts\check-no-rust-artifacts.ps1

## 4. Safety boundary

The helper still must not relay:

    signer.json
    provider-storage.json
    raw MemoryStorage JSON
    raw OpenMLS group state
    plaintext
    private keys
    trust-state private material

## 5. Next rung

Next planned rung:

    v0.2.52 — sidecar-produced application-message relay proof through Cypher and back into message-open.

Goal:

    sidecar message-protect writes application-message.bin
    relay helper submits artifact bytes through Cypher
    recipient retrieves envelope
    relay helper writes application-message.bin to a local sidecar-compatible path
    sidecar message-open consumes the written artifact
    plaintext matches

This should remain a dev/test proof, not polished runtime UX.
