# CarbonStack v0.3.34 full Profile Release Validation Ladder

Status: v0 runner semantics / release-checklist checkpoint
Scope: late v0.3.x -> v0.4.0 release runway
Primary environment: WSL Debian
Generated: 2026-06-03 09:05:41 -0400

## 1. Purpose

This document records the v0.3.34 release validation ladder decision.

v0.3.33 cleaned the public release surface for the pre-v0.4.0 runway. v0.3.34 updates the Go runner so `full` is no longer a stale alias for `core`.

## 2. Current Repo Heads

    carbonstack        ec9b02f runner: make full profile run release validation ladder
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

## 5. Recon Finding

The v0.3.34 recon found:

    `full` still aliased `core`.
    `release-snapshot` already performs package checks, strict pre-test artifact scan, checksum verification, and then calls `core`.
    `local-cypher` is implemented separately and validates the Cypher-only local lifecycle.
    The release directory in the live working repo is not a release package root.
    v0.4.0 needs a release validation ladder, not new feature work.

## 6. Implementation Decision

`full` now runs:

    release-snapshot
    local-cypher

`release-snapshot` already calls `core`, so `full` does not call `core` a second time.

Recommended v0.4.0 release-package validation command:

    go run . --profile full --root /path/to/release-package-root --clean-generated

## 7. Scope Boundary

`full` is a validation ladder.

It is not:

    a deployment command
    local-backbone
    runtime Comms UX
    public ingress
    systemd/cloudflared
    production deployability
    security certification

## 8. Why Not release-snapshot -> local-cypher -> core

That sequence would duplicate `core`, because `release-snapshot` already calls `core`.

Correct sequence:

    full = release-snapshot + local-cypher

Meaning:

    release-snapshot handles package/checksum/core validation.
    local-cypher adds Cypher-only local lifecycle validation.
    --clean-generated can remove known generated roots after the whole successful profile.

## 9. Validation Results

Command:

    go test ./... -count=1

Exit code:

    0

Output:

    ?   	git.bitcrusher32.win/bitcrusher32/carbonstack/tools/carbonstack-validate	[no test files]

Error output:

    (empty)

Command:

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
go version: go version go1.24.4 linux/amd64
sqlite3 path: /usr/bin/sqlite3
sqlite3 version: 3.46.1 2024-08-13 09:16:08 c9c2ab54ba1f5f46360f1b4f35d849cd3f080e6fc2b6c60e91b16c63f69aalt1 (64-bit)
== pre-local-cypher artifact scan ==
pre-local-cypher artifact scan: no generated/private/build artifact hits

== local-cypher environment ==
temp_dir: /tmp/carbonstack-local-cypher-117666479
cypher_bin: /tmp/carbonstack-local-cypher-117666479/carbonstack-cypher-local
cypher_db: /tmp/carbonstack-local-cypher-117666479/cypher.db
cypher_migrations: /home/Bitcrusher32/repos/carbonstack_umbrella/carbonstack-cypher/migrations
cypher_addr: 127.0.0.1:33137

== build temporary Cypher binary ==
DIR:  /home/Bitcrusher32/repos/carbonstack_umbrella/carbonstack-cypher
CMD:  go build -o /tmp/carbonstack-local-cypher-117666479/carbonstack-cypher-local ./cmd/cypher

== start first Cypher process ==
health response: {"api_version":"v0","service":"carbonstack-cypher","status":"ok"}

PASS: first health check
PASS: claim seeded Alice invite
PASS: create Bob dev invite
PASS: claim Bob invite
PASS: register Alice device
PASS: register Bob device
PASS: list Alice devices
PASS: reject invalid stub-text/OpenMLS protocol pairing
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

Command:

    go run . --profile doctor

Exit code:

    0

Output:

    ========================================
CarbonStack validation profile: doctor
========================================
os:             linux
arch:           amd64
start_dir:      /home/Bitcrusher32/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
umbrella_root:  /home/Bitcrusher32/repos/carbonstack_umbrella
carbonstack:    /home/Bitcrusher32/repos/carbonstack_umbrella/carbonstack
comms:          /home/Bitcrusher32/repos/carbonstack_umbrella/carbonstack-comms
cypher:         /home/Bitcrusher32/repos/carbonstack_umbrella/carbonstack-cypher

== Required paths ==
OK:      /home/Bitcrusher32/repos/carbonstack_umbrella/carbonstack
OK:      /home/Bitcrusher32/repos/carbonstack_umbrella/carbonstack-comms
OK:      /home/Bitcrusher32/repos/carbonstack_umbrella/carbonstack-cypher
OK:      /home/Bitcrusher32/repos/carbonstack_umbrella/carbonstack-comms/go.mod
OK:      /home/Bitcrusher32/repos/carbonstack_umbrella/carbonstack-cypher/go.mod
OK:      /home/Bitcrusher32/repos/carbonstack_umbrella/carbonstack-comms/internal/protocol/mls/openmls-sidecar/Cargo.toml

== Toolchains ==
go path: /usr/lib/go-1.24/bin/go
go version: go version go1.24.4 linux/amd64
rustc path: /home/Bitcrusher32/.cargo/bin/rustc
rustc version: rustc 1.96.0 (ac68faa20 2026-05-25)
cargo path: /home/Bitcrusher32/.cargo/bin/cargo
cargo version: cargo 1.96.0 (30a34c682 2026-05-25)
sqlite3 path: /usr/bin/sqlite3
sqlite3 version: 3.46.1 2024-08-13 09:16:08 c9c2ab54ba1f5f46360f1b4f35d849cd3f080e6fc2b6c60e91b16c63f69aalt1 (64-bit)

== Rust floor note ==
OpenMLS 0.8.1 failed under Debian apt rustc 1.85.0 during v0.3.9.
rustup stable rustc/cargo 1.96.0 passed under WSL Debian during v0.3.9.
This runner reports toolchain versions but does not install or mutate toolchains.

VALIDATION PASSED

Error output:

    (empty)

Command:

    go run . --profile core --clean-generated

Exit code:

    0

Output:

    ========================================
CarbonStack validation profile: core
========================================
========================================
CarbonStack validation profile: doctor
========================================
os:             linux
arch:           amd64
start_dir:      /home/Bitcrusher32/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
umbrella_root:  /home/Bitcrusher32/repos/carbonstack_umbrella
carbonstack:    /home/Bitcrusher32/repos/carbonstack_umbrella/carbonstack
comms:          /home/Bitcrusher32/repos/carbonstack_umbrella/carbonstack-comms
cypher:         /home/Bitcrusher32/repos/carbonstack_umbrella/carbonstack-cypher

== Required paths ==
OK:      /home/Bitcrusher32/repos/carbonstack_umbrella/carbonstack
OK:      /home/Bitcrusher32/repos/carbonstack_umbrella/carbonstack-comms
OK:      /home/Bitcrusher32/repos/carbonstack_umbrella/carbonstack-cypher
OK:      /home/Bitcrusher32/repos/carbonstack_umbrella/carbonstack-comms/go.mod
OK:      /home/Bitcrusher32/repos/carbonstack_umbrella/carbonstack-cypher/go.mod
OK:      /home/Bitcrusher32/repos/carbonstack_umbrella/carbonstack-comms/internal/protocol/mls/openmls-sidecar/Cargo.toml

== Toolchains ==
go path: /usr/lib/go-1.24/bin/go
go version: go version go1.24.4 linux/amd64
rustc path: /home/Bitcrusher32/.cargo/bin/rustc
rustc version: rustc 1.96.0 (ac68faa20 2026-05-25)
cargo path: /home/Bitcrusher32/.cargo/bin/cargo
cargo version: cargo 1.96.0 (30a34c682 2026-05-25)
sqlite3 path: /usr/bin/sqlite3
sqlite3 version: 3.46.1 2024-08-13 09:16:08 c9c2ab54ba1f5f46360f1b4f35d849cd3f080e6fc2b6c60e91b16c63f69aalt1 (64-bit)

== Rust floor note ==
OpenMLS 0.8.1 failed under Debian apt rustc 1.85.0 during v0.3.9.
rustup stable rustc/cargo 1.96.0 passed under WSL Debian during v0.3.9.
This runner reports toolchain versions but does not install or mutate toolchains.

== pre-test artifact scan ==
pre-test artifact scan: no generated/private/build artifact hits

----------------------------------------
STEP: OpenMLS real-Cypher lifecycle
DIR:  /home/Bitcrusher32/repos/carbonstack_umbrella/carbonstack-comms
CMD:  go test ./internal/protocol -run TestOpenMLSSidecarFullLifecycleRelayThroughRealCypherServer -count=1 -timeout 300s -v
ENV:  RUST_BACKTRACE=1
----------------------------------------
[stdout] === RUN   TestOpenMLSSidecarFullLifecycleRelayThroughRealCypherServer
[stdout] --- PASS: TestOpenMLSSidecarFullLifecycleRelayThroughRealCypherServer (15.59s)
[stdout] PASS
[stdout] ok  	git.bitcrusher32.win/bitcrusher32/carbonstack-comms/internal/protocol	15.592s
PASS: OpenMLS real-Cypher lifecycle elapsed=16.395s

----------------------------------------
STEP: carbonstack-comms package tests
DIR:  /home/Bitcrusher32/repos/carbonstack_umbrella/carbonstack-comms
CMD:  go test ./... -count=1 -timeout 600s
ENV:  RUST_BACKTRACE=1
----------------------------------------
[stdout] ?   	git.bitcrusher32.win/bitcrusher32/carbonstack-comms/cmd/comms	[no test files]
[stdout] ?   	git.bitcrusher32.win/bitcrusher32/carbonstack-comms/internal/app	[no test files]
[stdout] ok  	git.bitcrusher32.win/bitcrusher32/carbonstack-comms/internal/client	0.009s
[stdout] ok  	git.bitcrusher32.win/bitcrusher32/carbonstack-comms/internal/crypto	0.003s
[stdout] ok  	git.bitcrusher32.win/bitcrusher32/carbonstack-comms/internal/protocol	11.087s
[stdout] ok  	git.bitcrusher32.win/bitcrusher32/carbonstack-comms/internal/relay	0.006s
[stdout] ok  	git.bitcrusher32.win/bitcrusher32/carbonstack-comms/internal/state	0.003s
[stdout] ok  	git.bitcrusher32.win/bitcrusher32/carbonstack-comms/internal/trust	0.003s
PASS: carbonstack-comms package tests elapsed=11.594s

----------------------------------------
STEP: carbonstack-cypher package tests
DIR:  /home/Bitcrusher32/repos/carbonstack_umbrella/carbonstack-cypher
CMD:  go test ./... -count=1
----------------------------------------
[stdout] ?   	git.bitcrusher32.win/bitcrusher32/carbonstack-cypher/cmd/cypher	[no test files]
[stdout] ?   	git.bitcrusher32.win/bitcrusher32/carbonstack-cypher/internal/config	[no test files]
[stdout] ok  	git.bitcrusher32.win/bitcrusher32/carbonstack-cypher/internal/db	0.009s
[stdout] ok  	git.bitcrusher32.win/bitcrusher32/carbonstack-cypher/internal/httpapi	0.058s
PASS: carbonstack-cypher package tests elapsed=992ms

== post-test artifact scan ==
post-test artifact scan hits:
  [known-openmls-sidecar-generated-root] /home/Bitcrusher32/repos/carbonstack_umbrella/carbonstack-comms/internal/protocol/mls/openmls-sidecar/.carbonstack-openmls-sidecar-state
  [known-openmls-sidecar-generated-root] /home/Bitcrusher32/repos/carbonstack_umbrella/carbonstack-comms/internal/protocol/mls/openmls-sidecar/target
artifact scan is non-destructive
pre-test hits are potential source/copy hygiene issues
post-test hits are expected only when they stay in known generated roots

== clean generated artifacts ==
cleanup mode: explicit --clean-generated
cleanup scope: known generated/build artifact roots only
REMOVE: /home/Bitcrusher32/repos/carbonstack_umbrella/carbonstack-comms/internal/protocol/mls/openmls-sidecar/.carbonstack-openmls-sidecar-state
REMOVE: /home/Bitcrusher32/repos/carbonstack_umbrella/carbonstack-comms/internal/protocol/mls/openmls-sidecar/target
generated artifact cleanup complete

VALIDATION PASSED

Error output:

    (empty)

## 10. Recommended Next Rung

Recommended next rung:

    v0.3.35 release package rehearsal

Focus:

    create/stage a release-like package root
    generate release metadata/checksums
    validate from a fresh extraction
    run go run . --profile full --root <package-root> --clean-generated
    document exact package validation results
    prepare v0.4.0 release notes if clean

## 11. Summary

v0.3.34 makes `full` the intended v0.4.0 release-package validation ladder.

The runner now reflects the project reality: `core` remains core implementation validation, `local-cypher` remains Cypher-only lifecycle validation, `release-snapshot` remains package/checksum/core validation, and `full` ties release-snapshot plus local-cypher together for release-package validation.
