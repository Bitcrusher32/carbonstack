# CarbonStack v0.3.25 Local Cypher Operator Runbook Skeleton

Status: v0 runbook skeleton / local-only proof checkpoint
Scope: v0.3.x local-only backbone deployability line
Primary environment: WSL Debian
Generated: 2026-06-02 15:19:49 -0400

## 1. Purpose

This document records the first local Cypher operator runbook skeleton after the v0.3.24 schema migration tracking implementation.

v0.3.21 defined the local-only deployability model. v0.3.22 inspected Cypher local operator surfaces. v0.3.23 proved the old repeated-migration hazard. v0.3.24 implemented schema_migrations tracking. v0.3.25 now records the first simple local operator shape: explicit loopback bind, explicit SQLite DB path, explicit migrations path, explicit dev invite, health check, and reset boundary.

This is still not a production deployability claim.

## 2. Current Repo Heads

    carbonstack        76f0258 docs: record Cypher schema migrations implementation
    carbonstack-comms  012c8bf scripts: support source snapshot artifact guard
    carbonstack-cypher 42f838f db: track applied migrations
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

## 5. Local-Only Operator Model

Recommended local-only Cypher address:

    127.0.0.1:8080

Recommended experimental local operator DB path:

    $HOME/.local/share/carbonstack/cypher/cypher.db

Recommended explicit migrations path during source-tree development:

    $HOME/repos/carbonstack_umbrella/carbonstack-cypher/migrations

Recommended explicit dev invite for local bootstrap testing:

    dev-invite

The important rule is that local operator startup should not rely on ambiguous defaults. The operator should provide address, DB path, migrations path, and bootstrap invite explicitly.

## 6. Example Local Cypher Startup

From WSL Debian:

    cd ~/repos/carbonstack_umbrella/carbonstack-cypher

    mkdir -p ~/.local/share/carbonstack/cypher

    CYPHER_ADDR=127.0.0.1:8080 \
    CYPHER_DB=$HOME/.local/share/carbonstack/cypher/cypher.db \
    CYPHER_MIGRATIONS=$HOME/repos/carbonstack_umbrella/carbonstack-cypher/migrations \
    CYPHER_DEV_INVITE=dev-invite \
    go run ./cmd/cypher

This intentionally uses 127.0.0.1:8080 rather than the current code default :8080.

## 7. Health Check

In another terminal:

    curl -i http://127.0.0.1:8080/v0/health

Expected result:

    HTTP success response from the Cypher health endpoint.

Exact response shape may be refined later. The purpose of this rung is only to prove basic local process start and local HTTP health reachability.

## 8. Reset / Cleanup

Stop the Cypher process with Ctrl-C.

To reset the experimental local operator DB:

    rm -f ~/.local/share/carbonstack/cypher/cypher.db

This is intentionally blunt. Do not add reset tooling yet.

## 9. v0.3.25 Health-Check Proof

A temporary proof built and launched Cypher as a local binary with an isolated v0.3.25 health-check DB path.

Temporary binary path:

    /home/Bitcrusher32/.local/share/carbonstack/cypher-v0325-healthcheck/carbonstack-cypher-healthcheck

Temporary DB path:

    /home/Bitcrusher32/.local/share/carbonstack/cypher-v0325-healthcheck/cypher.db

Proof environment:

    CYPHER_ADDR=127.0.0.1:18080
    CYPHER_DB=/home/Bitcrusher32/.local/share/carbonstack/cypher-v0325-healthcheck/cypher.db
    CYPHER_MIGRATIONS=/home/Bitcrusher32/repos/carbonstack_umbrella/carbonstack-cypher/migrations
    CYPHER_DEV_INVITE=v0.3.25-healthcheck-dev-invite

Health-check URL:

    http://127.0.0.1:18080/v0/health

Observed health status:

    200

Observed health body:

    {"api_version":"v0","service":"carbonstack-cypher","status":"ok"}


Observed health error, if any:

    (none)

Process return code after termination:

    -15

Process output:

    (no process output captured)

The temporary proof DB and binary were deleted after the proof.

## 10. Current Persistence Boundary

v0.3.24 added migration tracking, which removes the known repeated-migration blocker from v0.3.23.

Current safe claim:

    Cypher has a basic schema_migrations ledger that supports repeated migration against the same DB when migration files are unchanged.

Still not safe:

    production-grade persistence
    production DB operations
    concurrent/multi-operator guarantees
    backup/restore guarantees
    production upgrade policy

## 11. Current Non-Goals

This runbook skeleton does not implement or validate:

    systemd
    cloudflared
    public ingress
    LAN exposure
    real homelab deployment
    production deployability
    production E2EE
    hostile-server safety
    metadata privacy
    runtime Comms OpenMLS UX
    Android app
    CarbonStackOS
    external audit
    certification

## 12. Next Rungs

Recommended next rungs:

    v0.3.26:
      local operator config/data path convention refinement
      decide whether to add a tiny source-tree helper script

    v0.3.27:
      local Cypher process + basic API lifecycle proof using explicit local operator env

    v0.3.28+:
      consider local-backbone runner profile only after start/stop/reset semantics are stable

## 13. Summary

v0.3.25 records the first local-only Cypher operator runbook skeleton.

The project now has:

    v0.3.21 local-only deployability model
    v0.3.22 Cypher local operator surface recon
    v0.3.23 repeated-migration hazard proof
    v0.3.24 schema_migrations implementation
    v0.3.25 local Cypher operator runbook skeleton and health-check proof

This is still pre-alpha infrastructure work, not deployability certification.
