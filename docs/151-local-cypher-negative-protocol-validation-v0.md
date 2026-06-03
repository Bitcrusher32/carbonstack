# CarbonStack v0.3.32 local-cypher Negative Protocol Validation

Status: v0 negative-path validation checkpoint
Scope: v0.3.x local-only backbone deployability line
Primary environment: WSL Debian
Generated: 2026-06-03 08:06:58 -0400

## 1. Purpose

This document records the v0.3.32 local-cypher negative protocol validation rung.

v0.3.30 implemented the first positive local-cypher runner profile. v0.3.31 polished the runner and added explicit generated-artifact cleanup with --clean-generated. v0.3.32 adds one historically grounded negative-path check to local-cypher.

This is still not local-backbone, not runtime Comms UX, not public ingress, and not production deployability.

## 2. Current Repo Heads

    carbonstack        7a123d7 runner: add local-cypher negative protocol validation
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

Updated files:

    tools/carbonstack-validate/local_cypher.go
    tools/carbonstack-validate/README.md

New runner behavior:

    local-cypher now verifies that the historically bad content/protocol pairing is rejected before the valid opaque OpenMLS application-message envelope is submitted.

Negative-path request:

    content_type=carbonstack.message.text.stub.v0
    protocol_version=carbonstack-openmls-sidecar-v0

Expected response:

    HTTP 400
    unsupported_protocol_version

This negative case comes from the v0.3.27 blunder where a stub text content type was incorrectly paired with the OpenMLS sidecar protocol version.

## 6. Why This Negative Path Matters

The negative check proves Cypher is not blindly accepting arbitrary content/protocol combinations during local-cypher validation.

This is a small but meaningful validation improvement because local-cypher now checks:

    positive lifecycle works
    accepted OpenMLS opaque envelope pair works
    one known-invalid protocol pairing is rejected

The accepted pair remains:

    content_type=carbonstack.mls.application-message.v0
    protocol_version=carbonstack-openmls-sidecar-v0

## 7. Scope Boundary

v0.3.32 adds one negative path only.

It does not add a broad adversarial test suite.

It does not yet cover:

    invalid base64 envelope
    ack unknown envelope
    ack with wrong recipient
    missing recipient_device_id on ack
    malformed account device path
    duplicate invite code
    already-claimed invite

These remain future validation maturity candidates.

## 8. Validation Results

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
temp_dir: /tmp/carbonstack-local-cypher-3137370852
cypher_bin: /tmp/carbonstack-local-cypher-3137370852/carbonstack-cypher-local
cypher_db: /tmp/carbonstack-local-cypher-3137370852/cypher.db
cypher_migrations: /home/Bitcrusher32/repos/carbonstack_umbrella/carbonstack-cypher/migrations
cypher_addr: 127.0.0.1:46377

== build temporary Cypher binary ==
DIR:  /home/Bitcrusher32/repos/carbonstack_umbrella/carbonstack-cypher
CMD:  go build -o /tmp/carbonstack-local-cypher-3137370852/carbonstack-cypher-local ./cmd/cypher

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
[stdout] --- PASS: TestOpenMLSSidecarFullLifecycleRelayThroughRealCypherServer (16.42s)
[stdout] PASS
[stdout] ok  	git.bitcrusher32.win/bitcrusher32/carbonstack-comms/internal/protocol	16.419s
PASS: OpenMLS real-Cypher lifecycle elapsed=17.127s

----------------------------------------
STEP: carbonstack-comms package tests
DIR:  /home/Bitcrusher32/repos/carbonstack_umbrella/carbonstack-comms
CMD:  go test ./... -count=1 -timeout 600s
ENV:  RUST_BACKTRACE=1
----------------------------------------
[stdout] ?   	git.bitcrusher32.win/bitcrusher32/carbonstack-comms/cmd/comms	[no test files]
[stdout] ?   	git.bitcrusher32.win/bitcrusher32/carbonstack-comms/internal/app	[no test files]
[stdout] ok  	git.bitcrusher32.win/bitcrusher32/carbonstack-comms/internal/client	0.008s
[stdout] ok  	git.bitcrusher32.win/bitcrusher32/carbonstack-comms/internal/crypto	0.002s
[stdout] ok  	git.bitcrusher32.win/bitcrusher32/carbonstack-comms/internal/protocol	11.142s
[stdout] ok  	git.bitcrusher32.win/bitcrusher32/carbonstack-comms/internal/relay	0.007s
[stdout] ok  	git.bitcrusher32.win/bitcrusher32/carbonstack-comms/internal/state	0.002s
[stdout] ok  	git.bitcrusher32.win/bitcrusher32/carbonstack-comms/internal/trust	0.003s
PASS: carbonstack-comms package tests elapsed=11.733s

----------------------------------------
STEP: carbonstack-cypher package tests
DIR:  /home/Bitcrusher32/repos/carbonstack_umbrella/carbonstack-cypher
CMD:  go test ./... -count=1
----------------------------------------
[stdout] ?   	git.bitcrusher32.win/bitcrusher32/carbonstack-cypher/cmd/cypher	[no test files]
[stdout] ?   	git.bitcrusher32.win/bitcrusher32/carbonstack-cypher/internal/config	[no test files]
[stdout] ok  	git.bitcrusher32.win/bitcrusher32/carbonstack-cypher/internal/db	0.009s
[stdout] ok  	git.bitcrusher32.win/bitcrusher32/carbonstack-cypher/internal/httpapi	0.057s
PASS: carbonstack-cypher package tests elapsed=778ms

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

Generated artifact state after core --clean-generated:

    ABSENT: /home/Bitcrusher32/repos/carbonstack_umbrella/carbonstack-comms/internal/protocol/mls/openmls-sidecar/.carbonstack-openmls-sidecar-state
ABSENT: /home/Bitcrusher32/repos/carbonstack_umbrella/carbonstack-comms/internal/protocol/mls/openmls-sidecar/target
ABSENT: /home/Bitcrusher32/repos/carbonstack_umbrella/carbonstack-comms/internal/protocol/mls/openmls-sidecar/provider-storage.json
ABSENT: /home/Bitcrusher32/repos/carbonstack_umbrella/carbonstack-comms/internal/protocol/mls/openmls-sidecar/signer.json

## 9. Current Nonclaims

v0.3.32 does not supersede v0.3.20 as the current public runner-backed testing release.

v0.3.32 does not make CarbonStack production-deployable.

v0.3.32 does not implement local-backbone.

v0.3.32 does not wire runtime Comms UX.

v0.3.32 does not prove hostile-server safety, metadata privacy, production E2EE, Android readiness, audit, or certification.

## 10. Release Framing Note

The eventual v0.4.0 release should be framed as a broad local deployability pre-release / milestone / research-and-development release.

It should clearly state:

    not intended for public-user use
    not intended for application use
    not production secure
    not hostile-server-certified
    not a mature Comms release
    not a CarbonStackOS release

The concrete validated artifact should remain the WSL Debian runner-backed validation surface, including local-cypher/core.

## 11. Recommended Next Rung

Recommended next rung:

    v0.3.33 local-cypher negative-path result review and release-surface readiness preflight

Possible focus:

    review whether this single negative path is enough for pre-v0.4.0
    decide whether more negative paths are needed now or later
    begin pre-v0.4.0 README/docs/roadmap claim cleanup planning
    keep local-backbone deferred
    keep runtime Comms UX deferred to v0.4.x

## 12. Summary

v0.3.32 adds the first negative-path validation to local-cypher.

The runner now verifies the historical stub-text/OpenMLS protocol mismatch fails with unsupported_protocol_version before continuing through the positive local-cypher lifecycle.

This strengthens the local-cypher validation surface without expanding the claim boundary.
