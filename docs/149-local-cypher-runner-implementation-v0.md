# CarbonStack v0.3.30 local-cypher Runner Implementation

Status: v0 implementation checkpoint
Scope: v0.3.x local-only backbone deployability line
Primary environment: WSL Debian
Generated: 2026-06-03 07:19:13 -0400

## 1. Purpose

This document records the v0.3.30 implementation of the `local-cypher` validation profile in the Go runner.

v0.3.29 defined the contract. v0.3.30 implements the first positive lifecycle version of that contract.

This is still not `local-backbone`, not runtime Comms UX, not public ingress, and not a production deployment claim.

## 2. Current Repo Heads

    carbonstack        a98ee9f runner: add local-cypher validation profile
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

## 5. Implementation Summary

New/updated runner files:

    tools/carbonstack-validate/local_cypher.go
    tools/carbonstack-validate/main.go
    tools/carbonstack-validate/README.md

The new profile command is:

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go run . --profile local-cypher

The profile:

    checks required paths
    reports basic toolchain information
    runs a pre-local-cypher artifact scan
    builds a temporary Cypher binary
    creates a temporary isolated DB
    chooses a dynamic loopback port
    starts Cypher on 127.0.0.1 only
    runs the positive invite/device/envelope/ack lifecycle
    verifies payload_sha256
    verifies payload_size_bytes
    stops Cypher
    restarts Cypher against the same temporary DB
    verifies persisted device state
    verifies the acked inbox remains empty
    stops Cypher again
    removes temporary state
    runs a post-local-cypher artifact scan

## 6. Scope Boundary

`local-cypher` is Cypher-only.

It does not validate:

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

## 7. Validation Results

Command:

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go test ./... -count=1

Exit code:

    0

Output:

    ?   	git.bitcrusher32.win/bitcrusher32/carbonstack/tools/carbonstack-validate	[no test files]

Error output:

    (empty)

Command:

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go run . --profile local-cypher

Exit code:

    0

Output:

    ========================================
CarbonStack validation profile: local-cypher
========================================

== Required paths ==
OK:      /home/Bitcrusher32/repos/carbonstack_umbrella/carbonstack
OK:      /home/Bitcrusher32/repos/carbonstack_umbrella/carbonstack-comms
OK:      /home/Bitcrusher32/repos/carbonstack_umbrella/carbonstack-cypher
OK:      /home/Bitcrusher32/repos/carbonstack_umbrella/carbonstack-comms/go.mod
OK:      /home/Bitcrusher32/repos/carbonstack_umbrella/carbonstack-cypher/go.mod
OK:      /home/Bitcrusher32/repos/carbonstack_umbrella/carbonstack-comms/internal/protocol/mls/openmls-sidecar/Cargo.toml

== Toolchains ==
go path: /usr/lib/go-1.24/bin/go
go version: /usr/lib/go-1.24
go path: /usr/lib/go-1.24
go path: /usr/lib/go-1.24/bin/go
go version: go version go1.24.4 linux/amd64
go version: go version go1.24.4 linux/amd64
sqlite3 path: /usr/bin/sqlite3
sqlite3 version: 3.46.1 2024-08-13 09:16:08 c9c2ab54ba1f5f46360f1b4f35d849cd3f080e6fc2b6c60e91b16c63f69aalt1 (64-bit)
sqlite3 version: 3.46.1 2024-08-13 09:16:08 c9c2ab54ba1f5f46360f1b4f35d849cd3f080e6fc2b6c60e91b16c63f69aalt1 (64-bit)
== pre-local-cypher artifact scan ==
pre-local-cypher artifact scan: no generated/private/build artifact hits

== local-cypher environment ==
temp_dir: /tmp/carbonstack-local-cypher-1337898034
cypher_bin: /tmp/carbonstack-local-cypher-1337898034/carbonstack-cypher-local
cypher_db: /tmp/carbonstack-local-cypher-1337898034/cypher.db
cypher_migrations: /home/Bitcrusher32/repos/carbonstack_umbrella/carbonstack-cypher/migrations
cypher_addr: 127.0.0.1:45191

== build temporary Cypher binary ==
DIR:  /home/Bitcrusher32/repos/carbonstack_umbrella/carbonstack-cypher
CMD:  go build -o /tmp/carbonstack-local-cypher-1337898034/carbonstack-cypher-local ./cmd/cypher

== start first Cypher process ==
health response: {"api_version":"v0","service":"carbonstack-cypher","status":"ok"}

PASS: first health check
PASS: claim seeded Alice invite
PASS: create Bob dev invite
PASS: claim Bob invite
PASS: register Alice device
PASS: register Bob device
PASS: list Alice devices
PASS: submit opaque OpenMLS application-message envelope
PASS: retrieve Bob inbox and verify payload metadata
PASS: ack envelope
PASS: retrieve Bob inbox after ack
INFO: first Cypher process stopped with expected termination: signal: interrupt

== restart Cypher against same temp DB ==
health response: {"api_version":"v0","service":"carbonstack-cypher","status":"ok"}

PASS: restart health check
PASS: persisted Alice device state after restart
PASS: acked Bob inbox remains empty after restart
INFO: second Cypher process stopped with expected termination: signal: interrupt
== post-local-cypher artifact scan ==
post-local-cypher artifact scan: no generated/private/build artifact hits

VALIDATION PASSED

VALIDATION PASSED

Error output:

    (empty)

## 8. Current Nonclaims

v0.3.30 does not supersede v0.3.20 as the public runner-backed testing release.

v0.3.30 does not make CarbonStack production-deployable.

v0.3.30 does not implement `local-backbone`.

v0.3.30 does not wire runtime Comms UX.

v0.3.30 does not add negative-path validation yet.

## 9. Recommended Next Rung

Recommended next rung:

    v0.3.31 local-cypher profile proof/result hardening checkpoint

Possible focus:

    run local-cypher, doctor, and core after the implementation commit
    document final clean status
    decide whether negative-path validation should be v0.3.32
    keep helper tooling deferred

## 10. Summary

v0.3.30 implements the first `local-cypher` runner profile.

It turns the v0.3.29 contract into a runnable Cypher-only local lifecycle validation profile while preserving the claim boundary: this is not local-backbone, not production security, not public deployability, and not runtime Comms UX.
