# Comms OpenMLS Cypher Relay Bridge Recon v0

Status: Recon
Component: CarbonStackComms + CarbonStackCypher + OpenMLS sidecar
Phase: v0.2.49 Comms-side relay bridge recon
Previous docs:
- docs/91-openmls-sidecar-artifact-ownership-map-v0.md
- docs/99-cypher-openmls-envelope-relay-recon-v0.md
- docs/100-cypher-openmls-envelope-content-type-result-v0.md

## 1. Purpose

This document defines the first CarbonStackComms bridge shape for moving OpenMLS sidecar artifacts through CarbonStackCypher.

The bridge should connect:

    sidecar artifact path hints
    Cypher /v0/envelopes opaque payloads
    sidecar consume commands

This is recon/design. It does not wire production `comms send` / `comms inbox`.

## 2. Current validated sides

The promoted OpenMLS sidecar can produce and consume:

    public-bundle.keypackage.bin
    welcome.bin
    application-message.bin

CarbonStackCypher can now accept OpenMLS artifact envelope content types:

    carbonstack.mls.keypackage.v0
    carbonstack.mls.welcome.v0
    carbonstack.mls.application-message.v0

with protocol version:

    carbonstack-openmls-sidecar-v0

through the existing envelope mailbox:

    POST /v0/envelopes
    GET /v0/devices/{device_id}/envelopes
    POST /v0/envelopes/{envelope_id}/ack

## 3. Bridge principle

Comms should act as the bridge.

Cypher should not invoke the sidecar.

Cypher should not parse MLS internals.

The sidecar should not know Cypher route semantics.

Comms should translate between:

    local sidecar artifact files
    Cypher opaque envelope payloads
    local sidecar input files

## 4. First bridge flow: KeyPackage

Producer side:

    Bob sidecar runs public-bundle-export --write-artifact.
    Comms reads the sidecar result JSON path hint for public-bundle.keypackage.bin.
    Comms reads artifact bytes.
    Comms base64-encodes bytes.
    Comms submits /v0/envelopes with content_type carbonstack.mls.keypackage.v0.

Consumer side:

    Alice Comms retrieves envelope from Cypher.
    Comms decodes ciphertext_b64.
    Comms writes bytes to a safe local artifact file.
    Comms passes that path to conversation-add-member --member-keypackage <path>.

## 5. Second bridge flow: Welcome

Producer side:

    Alice sidecar runs conversation-add-member.
    Comms reads the sidecar result JSON path hint for welcome.bin.
    Comms reads Welcome bytes.
    Comms base64-encodes bytes.
    Comms submits /v0/envelopes with content_type carbonstack.mls.welcome.v0.

Consumer side:

    Bob Comms retrieves envelope from Cypher.
    Comms decodes ciphertext_b64.
    Comms writes bytes to a safe local artifact file.
    Comms passes that path to conversation-join --welcome <path>.

## 6. Third bridge flow: application message

Producer side:

    Sender sidecar runs message-protect.
    Comms reads the sidecar result JSON path hint for application-message.bin.
    Comms reads protected message bytes.
    Comms base64-encodes bytes.
    Comms submits /v0/envelopes with content_type carbonstack.mls.application-message.v0.

Consumer side:

    Recipient Comms retrieves envelope from Cypher.
    Comms decodes ciphertext_b64.
    Comms writes bytes to a safe local artifact file.
    Comms passes that path to message-open --message <path>.

## 7. Suggested Comms helper boundaries

First helper package should be dev/test scoped unless implementation recon finds a better existing location.

Likely helper responsibilities:

    map sidecar artifact kind -> Cypher content_type
    read artifact bytes from safe sidecar path hint
    base64 encode artifact bytes
    build Cypher submit-envelope request
    parse Cypher inbox envelope response
    base64 decode envelope payload
    write artifact bytes to safe local file
    reject unsafe output paths
    reject unsupported content types
    avoid raw byte logging

Do not mix this with final `comms send` / `comms inbox` UX yet.

## 8. Suggested first tests

First tests should be integration-style but dev-scaffolded:

    sidecar produces application-message.bin
    helper reads artifact bytes
    helper constructs Cypher envelope request body
    helper decodes a Cypher envelope response body
    helper writes exact bytes back to a local artifact path
    sidecar opens the written artifact successfully

Then expand to:

    KeyPackage relay helper
    Welcome relay helper
    full Alice/Bob KeyPackage -> Welcome -> application-message artifact relay path

## 9. Safety boundaries

Comms bridge must not relay:

    signer.json
    provider-storage.json
    raw MemoryStorage JSON
    raw OpenMLS group state
    plaintext
    private keys
    trust-state private material

Comms bridge must not casually log:

    raw KeyPackage bytes
    raw Welcome bytes
    raw application-message bytes
    plaintext

Comms bridge may log:

    content_type
    protocol_version
    size in bytes
    SHA-256 reference later if implemented
    local path hints only when not secret-bearing and useful for dev diagnostics

## 10. Open questions

- Should bridge helpers live under internal/protocol, internal/client, or a new internal/relay package?
- Should the first bridge tests use an in-process httptest Cypher server or static Cypher response fixtures?
- Should downloaded artifact files live under sidecar dev state or a Comms-owned relay staging directory?
- Should ack happen only after sidecar command succeeds?
- Should the bridge track which Cypher envelope maps to which sidecar message_label?
- Should artifact staging use envelope_id as filename, message_label as filename, or generated safe local names?
- Should the first bridge proof start with application-message only, then KeyPackage/Welcome, or all three?

## 11. Near-future server-deployable experimental backbone goalset

After the Cypher + Comms sidecar artifact bridge is proven, the project can target a future experimental server-deployable backbone.

This should mean:

    deployable CarbonStackCypher server
    OpenMLS sidecar artifact relay contract
    docs for builders to deploy and experiment
    no Android app dependency
    no claim of production certification

This should not mean:

    certified secure
    externally audited
    production E2EE
    final hostile-server proof
    complete metadata privacy
    finished CarbonStack appliance

The future pre-release should be clearly labeled immature/experimental unless external senior reviewers audit and harden the implementation.

## 12. Next implementation candidate

After recon, the next candidate implementation is a narrow Comms dev/test relay helper:

    sidecar artifact path hint -> Cypher envelope request
    Cypher envelope response -> local sidecar-compatible artifact file

No polished UX.
No final runtime wiring.
No trust mutation.

## 13. Recon result / placement decision

Repository recon found that CarbonStackComms already has a usable Cypher client layer:

    internal/client/cypher.go

This client already exposes:

    SubmitEnvelope(...)
    Inbox(...)
    AckEnvelope(...)

and already carries:

    sender_device_id
    recipient_device_id
    content_type
    protocol_version
    ciphertext_b64

The CLI/runtime command layer already exists at:

    internal/app/commands.go

with current stub-era commands:

    send
    inbox
    ack

The OpenMLS sidecar contract tests and helpers live under:

    internal/protocol/openmls_sidecar_*_test.go

The first bridge implementation should not go directly into `internal/app`, because that would prematurely wire OpenMLS artifact relay into the user-facing send/inbox flow.

The first bridge implementation should also not live only as sidecar test helper code, because the bridge is a real Comms/Cypher translation boundary that will later feed runtime integration.

Recommended first implementation location:

    internal/relay

Reason:

    internal/client owns raw Cypher HTTP transport.
    internal/protocol owns provider/sidecar protocol facts and tests.
    internal/app owns user-facing CLI commands.
    internal/relay can own the translation between sidecar artifact files and Cypher opaque envelope payloads without claiming polished runtime UX.

## 14. Recommended first implementation scope

The next implementation rung should add a narrow dev/test relay helper package:

    carbonstack-comms/internal/relay

Initial responsibilities:

    map OpenMLS artifact kind to Cypher content_type
    use protocol version carbonstack-openmls-sidecar-v0
    read artifact bytes from a caller-provided safe path
    base64 encode artifact bytes
    submit through client.CypherClient.SubmitEnvelope
    decode client.EnvelopeRecord.CiphertextB64
    write exact bytes to a caller-provided safe output path
    reject unsupported artifact kinds
    avoid raw byte logging

Suggested constants:

    carbonstack.mls.keypackage.v0
    carbonstack.mls.welcome.v0
    carbonstack.mls.application-message.v0
    carbonstack-openmls-sidecar-v0

Suggested first tests:

    keypackage artifact bytes -> envelope request content_type mapping
    welcome artifact bytes -> envelope request content_type mapping
    application-message artifact bytes -> envelope request content_type mapping
    arbitrary binary bytes roundtrip through relay helper encode/decode
    unsupported artifact kind rejected
    output path parent directory is created or explicit error is returned
    no signer.json/provider-storage.json artifact kind exists

## 15. Explicit defer list

Do not implement yet:

    comms send integration
    comms inbox integration
    automatic sidecar invocation
    automatic ack after sidecar consume
    trust-state mutation
    local vault integration
    payload_sha256 / payload_size_bytes migration
    Android app behavior
    production server-deployable packaging

The next code rung is a bridge helper proof, not a user-facing runtime integration.

