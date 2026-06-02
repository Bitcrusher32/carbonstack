# CarbonStack v0.3.29 local-cypher Validation Contract Recon

Status: v0 contract / recon checkpoint
Scope: v0.3.x local-only backbone deployability line
Primary environment: WSL Debian
Generated: 2026-06-02 16:12:42 -0400

## 1. Purpose

This document records the v0.3.29 local-cypher validation contract recon.

v0.3.27 proved the local-only Cypher explicit-env API lifecycle manually. v0.3.28 decided not to add helper tooling or a runner profile yet, reserved local-backbone for later whole-stack validation, and named local-cypher as the future Cypher-only validation surface.

v0.3.29 now defines what a future local-cypher runner profile should own before implementation.

This is docs-only. It does not add a runner profile.

## 2. Current Repo Heads

    carbonstack        af33139 docs: record local operator helper runner decision recon
    carbonstack-comms  012c8bf scripts: support source snapshot artifact guard
    carbonstack-cypher 9ab994c docs: point to local operator runbook
    carbonstack-os     1bbbe52 docs: clarify CarbonStackOS target direction

## 3. Working Tree Status Before This Doc

    [carbonstack]
    (clean)

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

## 5. Scout Findings

The v0.3.29 read-only WSL Debian scout confirmed:

    runner files currently live under carbonstack/tools/carbonstack-validate
    runner implementation files are main.go, checksums.go, release_snapshot.go, README.md, and go.mod
    existing profiles are doctor, core, full, release-snapshot, write-checksums, and verify-checksums
    full currently aliases core
    release-snapshot calls core after package/checksum/strict pre-test checks
    core already runs pre-test artifact scan, targeted OpenMLS real-Cypher lifecycle, Comms package tests, Cypher package tests, and post-test artifact scan
    artifact scans are currently non-destructive outside strict release-snapshot pre-test behavior
    runner profile dispatch is currently centralized in main.go
    runner can infer an umbrella root or receive --root
    Cypher API tests already cover positive relay lifecycle and useful negative paths
    Cypher current config defaults remain CYPHER_ADDR=:8080, CYPHER_DB=cypher.db, CYPHER_MIGRATIONS=migrations, and CYPHER_DEV_INVITE=dev-invite
    local operator docs recommend explicit CYPHER_ADDR=127.0.0.1:8080 for local operator use

## 6. Candidate Profile Name

Future profile name:

    local-cypher

Intended command shape:

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go run . --profile local-cypher

Optional explicit-root form:

    go run . --profile local-cypher --root /path/to/carbonstack_umbrella

Do not call this profile local-backbone.

local-backbone remains reserved for a later whole-stack validation surface where CarbonStackComms, CarbonStackCypher, OpenMLS runtime wiring, and backbone-level lifecycle semantics are meaningfully integrated.

## 7. What local-cypher Should Validate

A future local-cypher profile should validate the Cypher-only local API lifecycle under an explicit local operator environment.

Required positive lifecycle:

    build a temporary Cypher binary
    create a temporary isolated data directory
    create a temporary isolated SQLite DB
    set CYPHER_ADDR to an available 127.0.0.1 port
    set CYPHER_DB to the temporary DB path
    set CYPHER_MIGRATIONS to the source-tree migrations path
    set CYPHER_DEV_INVITE to a known temporary dev invite
    start Cypher
    wait for GET /v0/health
    claim seeded invite
    create second dev invite
    claim second invite
    register two devices
    list account devices
    submit opaque OpenMLS application-message envelope
    retrieve recipient inbox
    verify payload_sha256
    verify payload_size_bytes
    ack envelope
    verify recipient inbox is empty after ack
    stop Cypher
    restart Cypher against the same temporary DB
    verify /v0/health after restart
    verify persisted device state remains visible
    verify acked recipient inbox remains empty after restart
    stop Cypher
    remove temporary binary and temporary DB directory

Required accepted envelope pair:

    content_type=carbonstack.mls.application-message.v0
    protocol_version=carbonstack-openmls-sidecar-v0

## 8. What local-cypher Should Not Validate Yet

local-cypher should not validate:

    runtime CarbonStackComms send/inbox UX
    OpenMLS user-facing message flow through Comms CLI
    CarbonStack Relay Space mechanics
    local-backbone behavior
    production E2EE
    hostile-server safety
    metadata privacy
    public ingress
    LAN exposure
    cloudflared
    systemd
    real homelab deployment
    Android app behavior
    CarbonStackOS
    external audit
    certification

local-cypher should remain a Cypher API lifecycle validation profile, not a CarbonStack secure messaging claim.

## 9. Temporary DB Contract

local-cypher should use a temporary isolated DB by default.

The profile must not mutate the documented manual operator DB path:

    $HOME/.local/share/carbonstack/cypher/cypher.db

Reason:

    validation should be repeatable
    validation should not depend on private local operator state
    validation should not destroy a user's manual operator DB
    validation should avoid stale state coupling
    validation should clean up after itself

Temporary DB validation is acceptable for dev/test and pre-v0.4.x release-adjacent validation.

Maturity warning:

    temporary DB validation is not enough for later public Comms or real-user testing maturity

Later maturity needs validation for:

    real persistent operator state
    backup/restore expectations
    upgrade semantics
    failure recovery
    operator mistakes
    user data migration boundaries

## 10. Process Lifecycle Contract

local-cypher should own the Cypher process lifecycle.

It should:

    build the temporary binary
    start the binary itself
    wait for health
    run lifecycle calls
    stop the process
    restart the process against the same temp DB
    stop the second process
    clean temporary files

It should not require a pre-existing manually started Cypher server.

Reason:

    runner-owned lifecycle is more repeatable
    pass/fail is less ambiguous
    cleanup can be enforced
    ports, DB paths, and env vars can be controlled
    validation should not depend on a user's manual terminal state

## 11. Port Contract

local-cypher should bind to loopback only.

Preferred future behavior:

    choose a temporary localhost port dynamically if practical

Acceptable first implementation:

    use a fixed high localhost port with clear error on conflict

Non-negotiable:

    do not bind 0.0.0.0
    do not bind LAN interfaces
    do not expose public ingress

## 12. Artifact and Cleanup Contract

local-cypher should leave no source-tree artifacts.

It should not commit or leave behind:

    temporary binary
    temporary SQLite DB
    temp proof directory
    generated sidecar state
    Rust target directory
    provider-storage.json
    signer.json
    .go-cache
    .go-tmp

Cleanup should happen on success and failure where possible.

The standard post-core OpenMLS generated roots may still appear when core is run separately:

    carbonstack-comms/internal/protocol/mls/openmls-sidecar/.carbonstack-openmls-sidecar-state
    carbonstack-comms/internal/protocol/mls/openmls-sidecar/target

local-cypher itself should avoid creating Comms/OpenMLS sidecar artifacts unless its scope expands later. If it remains Cypher-only, its artifact surface should stay smaller than core.

## 13. Pass/Fail Contract

local-cypher should fail if:

    required paths are missing
    Cypher build fails
    Cypher exits before health check
    health check never succeeds
    invite claim fails
    device registration fails
    envelope submit fails
    inbox retrieval fails
    payload_sha256 mismatch occurs
    payload_size_bytes mismatch occurs
    ack fails
    inbox is not empty after ack
    restart against the same DB fails
    persisted device state is missing after restart
    acked inbox is not empty after restart
    process cleanup fails badly enough to leave an active server
    temporary DB/binary cleanup fails badly enough to leave suspicious project artifacts

It may warn, rather than fail, if:

    cleanup removes temp state but process return code reflects expected SIGTERM
    temporary directory cleanup needs a second attempt but succeeds
    process stderr only contains expected startup logs

## 14. Negative Path Contract

local-cypher v1 may initially focus on the positive lifecycle.

However, future negative-path validation should include the v0.3.27 blunder as a deliberate test:

    content_type=carbonstack.message.text.stub.v0
    protocol_version=carbonstack-openmls-sidecar-v0

Expected result:

    HTTP 400
    error code unsupported_protocol_version

Other future candidate negative paths:

    invalid base64 envelope
    ack unknown envelope
    ack with wrong recipient
    missing recipient_device_id on ack
    malformed account device path
    duplicate invite code
    already-claimed invite

Recommendation:

    implement positive local-cypher profile first
    add negative-path coverage in a follow-up rung unless implementation is very small and obvious

## 15. Relationship To Existing Profiles

Existing profiles:

    doctor
    core
    full
    release-snapshot
    write-checksums
    verify-checksums

Proposed future relationship:

    doctor:
      environment/path/toolchain report

    local-cypher:
      doctor-like prerequisites plus Cypher-only local API lifecycle

    core:
      current OpenMLS real-Cypher lifecycle plus Comms package tests plus Cypher package tests

    full:
      should continue to alias core until there is a real reason to expand it

    release-snapshot:
      package/checksum/layout validation, then core

local-cypher should not replace core.

local-cypher should not become release-snapshot.

local-cypher should be a focused Cypher local lifecycle proof.

## 16. README / Help Contract If Implemented Later

If implemented, update:

    carbonstack/tools/carbonstack-validate/README.md
    carbonstack/tools/carbonstack-validate/main.go help text
    docs/README.md
    roadmap/ROADMAP.md

Help text should include local-cypher in the profile list only after implementation exists.

Do not document the profile as runnable before implementation.

## 17. Implementation Timing Decision

Do not implement local-cypher in v0.3.29.

Recommended next rung:

    v0.3.30 implement local-cypher runner profile

But only if this contract still looks correct after review.

If not, use v0.3.30 for contract correction before implementation.

## 18. Release Semantics

v0.3.29 does not supersede v0.3.20 as the public runner-backed testing release.

Current public release remains:

    v0.3.20

v0.3.29 is mainline local deployability groundwork.

It is not:

    production deployability
    public server release
    homelab guide
    Comms app UX
    Android release
    CarbonStackOS work
    audited security claim

## 19. Summary

v0.3.29 defines the future local-cypher validation contract.

The contract is:

    Cypher-only
    runner-owned process lifecycle
    temporary isolated DB
    explicit loopback bind
    positive invite/device/envelope/ack lifecycle
    restart against same DB
    cleanup required
    no local-backbone name
    no helper tooling
    no public deployment claim

This gives v0.3.30 a concrete implementation target without prematurely expanding CarbonStack's claims.
