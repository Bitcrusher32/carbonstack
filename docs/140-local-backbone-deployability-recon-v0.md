# CarbonStack v0.3.21 Local Backbone Deployability Recon

Status: v0 planning / recon checkpoint
Scope: CarbonStack v0.3.x experimental backbone
Primary environment: WSL Debian
Current release baseline: v0.3.20 runner-backed testing release

## 1. Purpose

This document records the v0.3.21 local-backbone deployability recon checkpoint.

v0.3.20 proved the public runner-backed testing release path for the experimental Cypher + Comms OpenMLS relay backbone. v0.3.21 begins the next line of work: making the backbone understandable as a local-only deployable operator profile before runtime Comms UX, public ingress, systemd, cloudflared, Android, CarbonStackOS, or production-security work.

The immediate goal is not to ship a public server. The goal is to define and inspect the surfaces needed for a deliberate local/server-like layout.

## 2. Current Validated Baseline

The WSL Debian sibling-repo working umbrella was validated from:

~/repos/carbonstack_umbrella

The Go validation runner passed:

* doctor
* core

Observed toolchain baseline:

* Go 1.24.4
* rustc 1.96.0 via rustup
* cargo 1.96.0 via rustup
* SQLite 3.46.1
* git 2.47.3

The core validation path passed:

* OpenMLS real-Cypher lifecycle test
* carbonstack-comms package tests
* carbonstack-cypher package tests
* pre-test generated/private/build artifact scan
* post-test artifact scan with only expected OpenMLS sidecar generated roots

This validates the current WSL Debian working umbrella as a suitable development base for v0.3.21 recon.

## 3. v0.3.21 Goal

v0.3.21 should define the first local-only deployability model for the Cypher + Comms OpenMLS backbone.

A developer/operator should eventually be able to run the current backbone in a deliberate local/server-like layout with documented:

* Cypher bind address
* Cypher port
* Cypher database path
* Cypher migration assumptions
* Comms-to-Cypher addressing seam
* start/stop lifecycle
* stdout/stderr logging expectations
* data cleanup boundaries
* what survives restart
* what remains intentionally ephemeral
* what remains unsafe or unproven

## 4. Explicit Non-Goals

v0.3.21 does not claim or implement:

* production deployability
* production E2EE
* hostile-server safety
* metadata privacy
* public ingress
* cloudflared routes
* systemd service
* real homelab deployment
* shared production paths
* runtime Comms send/inbox OpenMLS UX
* secure local vault
* Android app
* CarbonStackOS
* external audit
* certification

Runtime Comms OpenMLS send/inbox integration remains a v0.4.x direction.

## 5. Local-Only Operator Model

The initial deployability model should be local-only.

Recommended local Cypher bind address:

127.0.0.1:8080

Avoid for this phase:

* 0.0.0.0 bind
* LAN exposure
* public ingress
* cloudflared
* systemd
* persistent public service assumptions

The first useful operator model is:

* start Cypher locally
* point Comms/backbone validation at the local Cypher URL
* validate the OpenMLS relay lifecycle
* inspect logs/errors
* stop Cypher
* understand which state persists and which state is test-generated

## 6. Cypher Configuration Surface

Current Cypher deployability-relevant surfaces to inspect and document:

* address/bind configuration
* database path configuration
* migration directory configuration
* dev invite behavior
* server startup behavior
* server shutdown behavior
* stdout/stderr output
* SQLite file lifecycle
* envelope submit/retrieve/ack behavior
* payload hash/size metadata behavior
* OpenMLS content type handling

Current expected simple configuration shape:

CYPHER_ADDR=127.0.0.1:8080
CYPHER_DB=$HOME/.local/share/carbonstack/cypher/cypher.db
CYPHER_MIGRATIONS=./migrations
CYPHER_DEV_INVITE=dev-invite

These names reflect the current simple environment-variable style. A mature config file can be considered later if the operator model becomes too complex for explicit environment variables.

## 7. Database and Migration Position

CarbonStackCypher currently uses SQLite.

SQLite remains appropriate for v0.3.21 because the current target is a local-only single-node relay profile.

The Cypher database stores relay/server state, not plaintext message trust:

* dev invites
* accounts
* devices
* opaque encrypted envelopes
* envelope metadata
* delivery/ack state

The database must not be treated as a plaintext message store or a client trust oracle.

Persistent local DB behavior is experimental until migration behavior is hardened.

Known deployability concern:

* fresh DB behavior is the known-good path
* persistent DB reopen/upgrade needs explicit validation
* migration tracking or idempotent migration guards are likely needed before stronger persistence claims

Allowed v0.3.21 stance:

Persistent local DB state is experimental. Schema reset may be required between versions. PostgreSQL or another OSS database may be considered later after the Cypher data model and operator model mature.

## 8. Comms-to-Cypher Addressing Surface

The first deployability seam should be simple.

Expected future/semi-formal seam:

CARBONSTACK_CYPHER_URL=http://127.0.0.1:8080

A future CLI flag may also make sense:

--cypher-url http://127.0.0.1:8080

For v0.3.21, do not force final UX. The goal is to identify and document where Comms/backbone validation points to Cypher. Mature user-facing runtime wiring belongs to v0.4.x.

## 9. OpenMLS Sidecar Deployability Surface

The OpenMLS sidecar currently remains dev/pre-alpha infrastructure.

Deployability-relevant surfaces to inspect:

* sidecar command surface
* generated state root behavior
* provider-storage writes
* signer/private material boundaries
* KeyPackage artifact behavior
* Welcome artifact behavior
* application-message artifact behavior
* message labels
* duplicate/replay rejection behavior
* wrong-device/wrong-conversation rejection behavior
* generated artifact cleanup expectations

Secret-bearing generated state must not be committed, pasted, normalized as safe, or treated as a secure production vault.

## 10. Runner and Validation Surface

The Go runner remains the preferred validation authority candidate.

Current relevant profiles:

* doctor
* core
* full
* verify-checksums
* release-snapshot

For v0.3.21, do not create a deployability profile unless the local operator model is clear enough to validate safely.

Possible future profile:

local-backbone

Potential purpose:

* verify local operator layout
* start or target local Cypher
* run backbone validation against local config
* check generated artifact boundaries
* report nonclaims

This should not be rushed before the deployability model is documented.

## 11. Future Relay-Space / IRC-Like Mechanics

The later server model should not be planned before local deployability is grounded.

Current direction for later planning:

One addressable Cypher relay space maps to one server/conversation context.

This is intentionally simpler than a many-room general-purpose server model.

A future CarbonStack relay-space model may include:

* one Cypher relay instance per addressable relay space
* multiple isolated Cypher instances under one install
* one data directory per relay space
* one config per relay space
* one operator/admin boundary per relay space
* one clear join/invite context per relay space

This increases deployment overhead, but may reduce attack surface, metadata coupling, and cross-space administrative complexity.

Do not implement this in v0.3.21.

## 12. Future Invite / Join Categories

Current dev invites are bootstrap plumbing only.

Future invite categories may include:

* relay access invite
* conversation membership invite
* device enrollment invite
* operator/admin invite
* QR verification ceremony
* hardware-key-backed enrollment or recovery ceremony

These should remain separate concepts. Do not collapse them into the current dev invite mechanism.

## 13. Admin Direction

Long-term admin actions may be app-mediated and signed by an operator/admin identity or hardware-backed key.

Future direction:

* operator actions issued from a CarbonStack client/app
* Cypher verifies admin authority without becoming a message trust root
* admin authority remains separate from message trust
* admin actions are auditable
* server cannot read message plaintext

v0.3.21 should not implement a remote admin plane, web dashboard, or public operator workflow.

## 14. Recommended Next Implementation Rungs

Recommended order after this recon checkpoint:

1. Cypher deployability recon result doc
2. Cypher migration/persistence hardening plan
3. local-only operator runbook skeleton
4. simple local Cypher config/data directory convention
5. optional runner support only after the operator model is stable

Highest-value technical risk:

Cypher persistent DB migration behavior.

Before claiming a persistent local operator profile, Cypher should either support migration tracking or clearly document wipe-only experimental DB behavior.

## 15. Current Claim Boundary

Allowed claim:

CarbonStack v0.3.21 begins local-only deployability recon for the already released v0.3.x experimental Cypher + Comms OpenMLS backbone, using WSL Debian as the primary development/validation environment.

Not allowed:

v0.3.21 does not make CarbonStack deployable in production, systemd-ready, cloudflared-ready, hostile-server safe, metadata-private, Android-ready, audited, certified, or production-secure.
