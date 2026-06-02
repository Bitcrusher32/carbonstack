# CarbonStack v0.3.27 Local Cypher Explicit-Env API Lifecycle Proof

Status: v0 local API lifecycle proof checkpoint
Scope: v0.3.x local-only backbone deployability line
Primary environment: WSL Debian
Generated: 2026-06-02 15:46:46 -0400

## 1. Purpose

This document records the v0.3.27 local Cypher explicit-env API lifecycle proof.

v0.3.25 recorded the first local-only Cypher operator runbook skeleton. v0.3.26 recorded the local operator config/data convention. v0.3.27 now proves that the current Cypher API lifecycle can run under that explicit local operator environment against a temporary local SQLite DB.

This is still not a production deployability claim.

## 2. Current Repo Heads

    carbonstack        6e14fe0 docs: record local operator config data convention
    carbonstack-comms  012c8bf scripts: support source snapshot artifact guard
    carbonstack-cypher 9ab994c docs: point to local operator runbook
    carbonstack-os     1bbbe52 docs: clarify CarbonStackOS target direction

## 3. Working Tree Status Before This Doc

    [carbonstack]
    M docs/README.md
 M roadmap/ROADMAP.md

    [carbonstack-comms]
    (clean)

    [carbonstack-cypher]
    (clean)

    [carbonstack-os]
    (clean)

## 4. Toolchain Snapshot

    git:     git version 2.47.3
    go:      go version go1.24.4 linux/amd64
    rustc:   rustc 1.96.0 (ac68faa20 2026-05-25)
    cargo:   cargo 1.96.0 (30a34c682 2026-05-25)
    sqlite3: 3.46.1 2024-08-13 09:16:08 c9c2ab54ba1f5f46360f1b4f35d849cd3f080e6fc2b6c60e91b16c63f69aalt1 (64-bit)

## 5. Proof Result

    PASS

Build exit code:

    0

Build stdout:

    (empty)

Build stderr:

    (empty)

Proof errors:

    (none)

## 6. Proof Environment

Temporary proof directory:

    /home/Bitcrusher32/.local/share/carbonstack/cypher-v0327-lifecycle-proof

Temporary DB path:

    /home/Bitcrusher32/.local/share/carbonstack/cypher-v0327-lifecycle-proof/cypher.db

Temporary binary path:

    /home/Bitcrusher32/.local/share/carbonstack/cypher-v0327-lifecycle-proof/carbonstack-cypher-lifecycle-proof

Runtime environment:

    CYPHER_ADDR=127.0.0.1:18081
    CYPHER_DB=/home/Bitcrusher32/.local/share/carbonstack/cypher-v0327-lifecycle-proof/cypher.db
    CYPHER_MIGRATIONS=/home/Bitcrusher32/repos/carbonstack_umbrella/carbonstack-cypher/migrations
    CYPHER_DEV_INVITE=dev-invite

The temporary DB and binary were deleted after the proof.

## 7. API Lifecycle Steps

### first health check

Status: ok

    "{\"api_version\":\"v0\",\"service\":\"carbonstack-cypher\",\"status\":\"ok\"}\n"

### claim seeded alice invite

Status: ok

    {
      "account_id": "03c08367-d3d1-4c1a-ae21-7167367e963d",
      "created_at": "2026-06-02T19:46:48Z"
    }

### create bob dev invite

Status: ok

    {
      "created_at": "2026-06-02T19:46:48Z",
      "invite_code": "bob-local-proof-invite",
      "invite_id": "a0d00c52-f5f1-4eb9-817c-b1b210a40489"
    }

### claim bob invite

Status: ok

    {
      "account_id": "75d031b4-dea5-4464-a117-e4b0c42e5e3e",
      "created_at": "2026-06-02T19:46:48Z"
    }

### register alice device

Status: ok

    {
      "account_id": "03c08367-d3d1-4c1a-ae21-7167367e963d",
      "created_at": "2026-06-02T19:46:48Z",
      "device_id": "42994436-3932-496f-a7a2-1f83717aaf43"
    }

### register bob device

Status: ok

    {
      "account_id": "75d031b4-dea5-4464-a117-e4b0c42e5e3e",
      "created_at": "2026-06-02T19:46:48Z",
      "device_id": "d8ab2f6e-e862-406e-b0eb-219bf9a3fcee"
    }

### list alice devices

Status: ok

    {
      "account_id": "03c08367-d3d1-4c1a-ae21-7167367e963d",
      "devices": [
        {
          "created_at": "2026-06-02T19:46:48Z",
          "device_id": "42994436-3932-496f-a7a2-1f83717aaf43",
          "device_label": "alice-local-proof-device",
          "public_identity_key": "stub-alice-public-identity-key",
          "public_prekey_bundle": "stub-alice-public-prekey-bundle"
        }
      ]
    }

### submit opaque MLS application-message envelope

Status: ok

    {
      "delivery_state": "queued",
      "envelope_id": "66979f56-c546-42c1-baa0-3b365cbf70fe",
      "payload_sha256": "b008cf0e5535db0deebdada8899bd0ba815c04b7bb399ddbef86025aca6fed59",
      "payload_size_bytes": 57,
      "server_received_at": "2026-06-02T19:46:48Z"
    }

### retrieve bob inbox before ack

Status: ok

    {
      "device_id": "d8ab2f6e-e862-406e-b0eb-219bf9a3fcee",
      "envelopes": [
        {
          "ciphertext_b64": "djAuMy4yNyBsb2NhbCBvcGVyYXRvciBsaWZlY3ljbGUgcHJvb2Ygb3BhcXVlIG1scyBwYXlsb2Fk",
          "client_created_at": "",
          "content_type": "carbonstack.mls.application-message.v0",
          "delivery_state": "queued",
          "envelope_id": "66979f56-c546-42c1-baa0-3b365cbf70fe",
          "payload_sha256": "b008cf0e5535db0deebdada8899bd0ba815c04b7bb399ddbef86025aca6fed59",
          "payload_size_bytes": 57,
          "protocol_version": "carbonstack-openmls-sidecar-v0",
          "recipient_device_id": "d8ab2f6e-e862-406e-b0eb-219bf9a3fcee",
          "sender_device_id": "42994436-3932-496f-a7a2-1f83717aaf43",
          "server_received_at": "2026-06-02T19:46:48Z"
        }
      ]
    }

### ack envelope

Status: ok

    {
      "acknowledged_at": "2026-06-02T19:46:48Z",
      "delivery_state": "acknowledged",
      "envelope_id": "66979f56-c546-42c1-baa0-3b365cbf70fe"
    }

### retrieve bob inbox after ack

Status: ok

    {
      "device_id": "d8ab2f6e-e862-406e-b0eb-219bf9a3fcee",
      "envelopes": []
    }

### restart health check against same DB

Status: ok

    "{\"api_version\":\"v0\",\"service\":\"carbonstack-cypher\",\"status\":\"ok\"}\n"

### list persisted alice devices after restart

Status: ok

    {
      "account_id": "03c08367-d3d1-4c1a-ae21-7167367e963d",
      "devices": [
        {
          "created_at": "2026-06-02T19:46:48Z",
          "device_id": "42994436-3932-496f-a7a2-1f83717aaf43",
          "device_label": "alice-local-proof-device",
          "public_identity_key": "stub-alice-public-identity-key",
          "public_prekey_bundle": "stub-alice-public-prekey-bundle"
        }
      ]
    }

### verify acked inbox remains empty after restart

Status: ok

    {
      "device_id": "d8ab2f6e-e862-406e-b0eb-219bf9a3fcee",
      "envelopes": []
    }

## 8. Blunder Preserved: Wrong Protocol Pairing In First Attempt

The first v0.3.27 proof attempt used:

    content_type=carbonstack.message.text.stub.v0
    protocol_version=carbonstack-openmls-sidecar-v0

Cypher correctly rejected that envelope with:

    unsupported_protocol_version

That was a script blunder, not a Cypher failure. The corrected proof uses:

    content_type=carbonstack.mls.application-message.v0
    protocol_version=carbonstack-openmls-sidecar-v0

This matches the accepted OpenMLS envelope pattern and keeps the payload opaque.

## 9. Process Output

First process return code:

    -15

First process stdout:

    (empty)

First process stderr:

    2026/06/02 15:46:47 CarbonStackCypher listening on 127.0.0.1:18081
2026/06/02 15:46:47 database: /home/Bitcrusher32/.local/share/carbonstack/cypher-v0327-lifecycle-proof/cypher.db
2026/06/02 15:46:47 dev invite enabled: true

Second process return code:

    -15

Second process stdout:

    (empty)

Second process stderr:

    2026/06/02 15:46:48 CarbonStackCypher listening on 127.0.0.1:18081
2026/06/02 15:46:48 database: /home/Bitcrusher32/.local/share/carbonstack/cypher-v0327-lifecycle-proof/cypher.db
2026/06/02 15:46:48 dev invite enabled: true

## 10. What This Proves

This proof validates the current local-only Cypher API lifecycle under explicit local operator environment variables:

    health endpoint responds
    seeded dev invite can be claimed
    a second dev invite can be created and claimed
    devices can be registered
    account device listing works
    an opaque OpenMLS application-message envelope can be submitted
    recipient inbox retrieval works
    payload hash/size metadata is returned
    ack works
    acked envelope leaves the recipient inbox
    Cypher can restart against the same DB
    already-applied migrations do not block restart
    persisted device state remains visible after restart
    acked inbox remains empty after restart

## 11. What This Does Not Prove

This proof does not validate:

    production deployability
    production E2EE
    hostile-server safety
    metadata privacy
    public ingress
    LAN exposure
    systemd
    cloudflared
    real homelab deployment
    remote admin plane
    runtime Comms OpenMLS UX
    Android app
    CarbonStackOS
    external audit
    certification

## 12. Implementation Boundary

v0.3.27 does not change Cypher code.

It records that the current Cypher API can complete the local operator lifecycle under explicit environment variables. It does not add helper tooling, a config parser, a runner local-backbone profile, or a public deployment guide.

## 13. Recommended Next Rung

Recommended next rung:

    v0.3.28 local operator lifecycle proof review and helper/runner decision point

Possible next actions:

    decide whether a tiny manual helper script is justified
    decide whether local-backbone runner profile success criteria are concrete enough
    decide whether more API lifecycle proof should cover negative paths
    decide whether explicit 127.0.0.1 default should become a code change later
    keep public ingress/systemd/cloudflared deferred

## 14. Summary

v0.3.27 proves that the local-only Cypher operator convention can support a basic end-to-end Cypher API lifecycle using explicit env vars and a local SQLite DB, including restart against the same DB after schema_migrations landed.

This moves the v0.3.x local deployability line from health-check-only proof toward a real local API lifecycle proof, while still avoiding public deployment claims.
