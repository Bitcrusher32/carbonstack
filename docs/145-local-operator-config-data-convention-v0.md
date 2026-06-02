# CarbonStack v0.3.26 Local Operator Config/Data Convention

Status: v0 convention checkpoint
Scope: v0.3.x local-only backbone deployability line
Primary environment: WSL Debian
Generated: 2026-06-02 15:31:00 -0400

## 1. Purpose

This document records the v0.3.26 local operator config/data convention checkpoint.

v0.3.25 recorded the first local-only Cypher operator runbook skeleton and health-check proof. v0.3.26 refines the file/path/config convention that should be used before helper scripts, runner local-backbone profiles, public ingress, systemd, cloudflared, or real homelab work.

This is docs/convention work only. It does not change Cypher code defaults.

## 2. Current Repo Heads

    carbonstack        a5c6351 docs: update readme, roadmap, and historical doc
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

## 5. Convention Summary

Recommended local-only operator address:

    127.0.0.1:8080

Recommended local operator data directory:

    $HOME/.local/share/carbonstack/cypher

Recommended local operator DB path:

    $HOME/.local/share/carbonstack/cypher/cypher.db

Recommended local operator config directory:

    $HOME/.config/carbonstack/cypher

Recommended source-tree migrations path during WSL Debian development:

    $HOME/repos/carbonstack_umbrella/carbonstack-cypher/migrations

Recommended current config style:

    explicit environment variables

Not recommended yet:

    systemd unit files
    cloudflared routes
    public bind addresses
    LAN bind addresses
    global install paths
    remote operator plane
    config file parser
    helper/supervisor command
    local-backbone runner profile

## 6. Why Separate Source State From Operator State

Source tree state and operator state should stay separate.

Source tree:

    code
    tests
    migrations
    docs
    validation runner inputs

Operator state:

    local SQLite DB
    local operator config notes/env
    temporary logs if added later
    local reset target

The source tree should be safe to inspect, commit, package, and validate without accidentally carrying live operator state.

The local DB should not live at:

    carbonstack-cypher/cypher.db

Preferred local-only DB path:

    $HOME/.local/share/carbonstack/cypher/cypher.db

This keeps the DB outside Git repos while remaining easy to find and reset.

## 7. Recommended Environment Variables

Recommended local operator environment:

    CYPHER_ADDR=127.0.0.1:8080
    CYPHER_DB=$HOME/.local/share/carbonstack/cypher/cypher.db
    CYPHER_MIGRATIONS=$HOME/repos/carbonstack_umbrella/carbonstack-cypher/migrations
    CYPHER_DEV_INVITE=dev-invite

Important:

    CYPHER_ADDR should be explicit.
    CYPHER_DB should be explicit.
    CYPHER_MIGRATIONS should be explicit.
    CYPHER_DEV_INVITE should be explicit.

Do not rely on ambiguous defaults for local operator experiments.

## 8. Current Startup Shape

From WSL Debian:

    cd ~/repos/carbonstack_umbrella/carbonstack-cypher

    mkdir -p ~/.local/share/carbonstack/cypher
    mkdir -p ~/.config/carbonstack/cypher

    CYPHER_ADDR=127.0.0.1:8080 \
    CYPHER_DB=$HOME/.local/share/carbonstack/cypher/cypher.db \
    CYPHER_MIGRATIONS=$HOME/repos/carbonstack_umbrella/carbonstack-cypher/migrations \
    CYPHER_DEV_INVITE=dev-invite \
    go run ./cmd/cypher

In another terminal:

    curl -i http://127.0.0.1:8080/v0/health

## 9. Reset Convention

Stop Cypher first.

Then remove the experimental local DB:

    rm -f ~/.local/share/carbonstack/cypher/cypher.db

This is intentionally blunt.

Do not add reset tooling yet. The operator must understand what state is being destroyed.

## 10. Migration Convention

v0.3.24 implemented schema_migrations.

Current migration policy:

    migration files are applied in sorted filename order
    applied migrations are recorded by migration_name, sha256, and applied_at
    already-applied matching migrations are skipped
    already-applied migration checksum drift hard-fails
    applied migration files should be treated as immutable
    fix forward with a new migration file

This makes local operator persistence less fragile, but does not make the DB production-grade.

## 11. Config File Decision

Do not add a config parser yet.

Reason:

    env vars are already supported
    local operator semantics are still maturing
    adding TOML/YAML/JSON config introduces parser and UX surface
    the v0.3.x goal is local-only deployability groundwork, not mature packaging

Future config files may be reconsidered after:

    local operator lifecycle is validated
    reset/backup/upgrade semantics are clearer
    runtime Comms integration is closer
    CarbonStack Relay Space mechanics are less abstract

## 12. Helper Tool Decision

Do not add helper tooling yet.

Future helper shape remains plausible:

    carbonstack local start
    carbonstack local health
    carbonstack local stop
    carbonstack local reset

But it is too early because start/stop/reset success contracts are not stable enough.

Helper tooling should follow a validated manual convention, not invent one.

## 13. Runner Profile Decision

Do not add a local-backbone runner profile yet.

The Go runner remains the validation authority candidate, but a local-backbone profile needs a precise contract.

Missing pieces before a runner profile:

    stable local operator startup convention
    stable health check expectation
    stable reset semantics
    stable generated artifact boundaries
    decision on whether runner starts/stops Cypher or targets an already-running process
    decision on DB lifecycle during validation

## 14. Non-Goals

v0.3.26 does not implement or validate:

    production deployability
    production E2EE
    hostile-server safety
    metadata privacy
    systemd
    cloudflared
    public ingress
    LAN exposure
    real homelab deployment
    remote admin plane
    runtime Comms OpenMLS UX
    Android app
    CarbonStackOS
    external audit
    certification

## 15. Recommended Next Rung

Recommended next rung:

    v0.3.27 local Cypher explicit-env API lifecycle proof

Suggested proof:

    start Cypher with explicit local operator env
    verify /v0/health
    create dev invite
    claim invite
    register device
    submit envelope
    retrieve envelope
    ack envelope
    stop Cypher
    restart Cypher against the same DB
    verify persisted state where appropriate

This should still be local-only and pre-alpha.

## 16. Summary

v0.3.26 records the local operator config/data convention.

The convention is:

    explicit loopback bind
    explicit DB path outside source tree
    explicit migrations path
    explicit dev invite
    blunt reset command
    no config parser yet
    no helper command yet
    no runner local-backbone profile yet

This keeps the project moving toward local deployability without overclaiming maturity.
