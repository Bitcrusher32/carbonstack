# Go Validation Runner Design v0

Status: design record
Phase: v0.3.10 validation substrate design
Previous checkpoint: docs/129-wsl-debian-quick-portability-test-v0.md

## 1. Purpose

This document defines the first design for a Go-based CarbonStack validation runner.

The runner is intended to replace the PowerShell umbrella validation role over time.

The runner should not replace repo-local tests.

The runner should call known-good repo-local functional tests and make validation portable across Windows and Debian-like Linux environments.

## 2. Background

CarbonStack v0.3.x has validated the experimental Cypher + Comms OpenMLS relay backbone through several increasingly public paths.

Current verified public release path:

    v0.3.4 testing bug hotfix
    Windows 11 / PowerShell
    attached release source snapshots
    no-Git source snapshot self-test fallback

Current WSL Debian quick-portability finding:

    Debian WSL can run the current core Comms/Cypher OpenMLS backbone tests after the Rust toolchain is corrected.

The v0.3.9 WSL Debian test found:

    Debian apt Rust 1.85.0 was too old for OpenMLS 0.8.1.
    rustup stable Rust/Cargo 1.96.0 worked.
    targeted OpenMLS real-Cypher lifecycle passed.
    full carbonstack-comms/internal/protocol passed.
    full carbonstack-comms package suite passed.
    full carbonstack-cypher package suite passed.

The remaining portability gap is mostly validation orchestration, not obvious Go/Rust core behavior.

## 3. Design decision

Create an umbrella Go validation runner under the `carbonstack` repo.

Initial location:

    carbonstack/tools/carbonstack-validate

Reason:

    carbonstack is the doctrine/spec/docs/release authority;
    repo-local tests already exist in carbonstack-comms and carbonstack-cypher;
    the umbrella should orchestrate known-good functional tests across repos;
    validation should converge across Windows and Debian/WSL instead of maintaining divergent PowerShell and bash authorities.

## 4. Non-goals

The first runner is not a deployment tool.

It must not:

    expose Cypher publicly;
    configure cloudflared;
    create systemd services;
    install packages;
    mutate production paths;
    replace repo-local tests;
    make security claims;
    hide failing command output;
    implement runtime Comms send/inbox UX;
    package releases;
    publish releases.

## 5. First command shape

Preferred developer command:

    go run ./tools/carbonstack-validate --profile core

Later command:

    go run ./tools/carbonstack-validate --profile full

Optional future built binary:

    carbonstack-validate --profile core
    carbonstack-validate --profile full

## 6. Initial profiles

### doctor

Purpose:

    report environment readiness without running heavy tests.

Should report:

    OS;
    architecture;
    Go version;
    Rust version;
    Cargo version;
    sqlite3 availability;
    current working directory;
    inferred umbrella root;
    expected sibling repo paths;
    whether required repos exist.

### core

Purpose:

    run the current core validation set.

Initial core profile should run:

    doctor;
    required sibling repo checks;
    carbonstack-cypher package tests;
    carbonstack-comms package tests;
    targeted OpenMLS real-Cypher lifecycle test;
    generated artifact scan.

Initial core profile should not perform release packaging or release asset verification.

### full

Purpose:

    call core first, then future broader validation profiles.

At v0.3.10, full is a design target.

It should be intentionally simple:

    full = core + future release/snapshot checks

Do not overload "full" with deployment or security meaning.

### future: release-snapshot

Purpose:

    validate extracted attached release source snapshots.

This should eventually replace the manual public release asset verification flow.

### future: openmls-backbone

Purpose:

    run only the targeted OpenMLS backbone lifecycle proof.

### future: cypher

Purpose:

    run Cypher-only package tests.

### future: comms

Purpose:

    run Comms-only package tests.

## 7. Expected repo layout

The runner assumes sibling repo layout.

When run from:

    carbonstack/

The umbrella root is:

    parent directory of carbonstack/

Expected siblings:

    carbonstack/
    carbonstack-comms/
    carbonstack-cypher/

This should work for both:

    [WIN-PWRSHL] C:\Users\udaiv\repos\carbonstack_umbrella\carbonstack
    [WSL-DEBIAN] ~/carbonstack-wsl/carbonstack

The runner should support an explicit root override later:

    --root <path>

But the first implementation can use inferred sibling layout.

## 8. Toolchain requirements

The runner should check:

    go
    rustc
    cargo
    sqlite3

Known WSL Debian toolchain finding:

    Debian apt rustc 1.85.0 is too old for OpenMLS 0.8.1.
    rustup stable rustc/cargo 1.96.0 worked.

The runner should print Rust/Cargo versions before OpenMLS-related tests.

The runner should not install toolchains.

## 9. Command execution behavior

The runner should be boring and explicit.

For each command, it should print:

    step name;
    working directory;
    command;
    start time;
    streamed stdout/stderr;
    pass/fail;
    elapsed time.

On first failure, the default behavior should be:

    stop;
    print failing step;
    return nonzero exit code.

Later optional behavior:

    --keep-going

## 10. Initial command set

The initial core profile should call:

From carbonstack-cypher:

    go test ./... -count=1

From carbonstack-comms:

    go test ./... -count=1 -timeout 600s

From carbonstack-comms:

    go test ./internal/protocol -run TestOpenMLSSidecarFullLifecycleRelayThroughRealCypherServer -count=1 -timeout 300s -v

The targeted OpenMLS real-Cypher lifecycle test may be redundant with the full Comms package suite, but it should stay as a separately visible functional proof because it is the backbone test public docs refer to.

## 11. Artifact scan behavior

The runner should scan for generated/private/build artifacts.

Pre-test scan should identify unexpected shipped/copied artifacts.

Post-test scan should report expected generated artifacts.

Known generated categories:

    target/
    .carbonstack-openmls-sidecar-state/
    provider-storage.json
    signer.json
    *.db
    *.db-shm
    *.db-wal
    *.exe
    *.test.exe
    .go-cache/
    .go-tmp/

The runner should distinguish:

    pre-test artifact hits:
        possible hygiene problem

    post-test artifact hits:
        expected if generated by validation and located in known roots

Do not make artifact scanning destructive in the first implementation.

No cleanup behavior in v0.3.10/v0.3.11.

## 12. Windows behavior

Windows remains the current public release platform boundary for v0.3.4.

The runner should eventually work under Windows.

The runner should reduce reliance on:

    carbonstack/scripts/validate-local.ps1
    carbonstack-comms/scripts/self-test-openmls-backbone.ps1

PowerShell scripts may remain as convenience wrappers.

PowerShell should stop being the validation authority after the Go runner is validated across Windows and WSL Debian.

## 13. WSL Debian behavior

WSL Debian is now the preferred candidate for fast core-development validation.

Reasons:

    WSL Debian passed current core Comms/Cypher OpenMLS tests after rustup stable;
    tests are faster;
    no Avast/Windows Go test executable interference was observed;
    Debian-like filesystem/toolchain behavior is closer to future CLI deployment targets.

Do not switch final authority yet.

The authority should move to the Go runner, not to bash or shell-specific WSL behavior.

## 14. Authority transfer criteria

WSL Debian can become the preferred core-validation environment when:

    Go runner core profile passes on WSL Debian;
    Go runner core profile passes on Windows;
    toolchain floors are documented;
    generated artifact behavior is stable;
    runner output is clear enough to replace validate-local.ps1 for core validation.

Windows can stop being the primary validation authority when:

    release packaging and public testing use the Go runner;
    Windows and WSL Debian run equivalent validation profiles;
    release docs no longer depend on PowerShell-specific semantics.

Until then:

    [WIN-PWRSHL] remains current public release compatibility / legacy validation path.
    [WSL-DEBIAN] is the preferred fast core-dev/test candidate.
    [GO-RUNNER] is the future validation authority.

## 15. Security and claim boundaries

The runner must not create new security claims.

Passing the runner does not prove:

    production readiness;
    production E2EE;
    hostile-server safety;
    metadata privacy;
    Android readiness;
    secure vault/storage;
    Linux public release support;
    Debian deployability;
    systemd readiness;
    cloudflared readiness;
    external audit;
    certification.

Every validation profile must have a stated scope and non-scope.

## 16. Implementation plan

Recommended v0.3.11 implementation scope:

    create tools/carbonstack-validate/main.go;
    implement --profile doctor;
    implement --profile core;
    infer sibling repo layout;
    print toolchain versions;
    run Cypher tests;
    run Comms tests;
    run targeted OpenMLS real-Cypher lifecycle test;
    perform non-destructive artifact scan;
    return nonzero on failure.

Do not implement release-snapshot profile yet.

Do not implement cleanup yet.

Do not implement deployment behavior.

## 17. Future release direction

Once the Go runner exists and passes on both Windows and WSL Debian, future releases can move toward a near-universal Windows + Debian/WSL validation story.

A future public release should include:

    toolchain floor requirements;
    Go runner instructions;
    Windows validation result;
    WSL Debian validation result;
    explicit nonclaim boundary.

Do not make a Linux/Debian release claim until runner-backed release validation exists.
