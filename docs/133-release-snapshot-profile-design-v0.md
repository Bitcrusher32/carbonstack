# Release Snapshot Profile Design v0

Status: design record
Phase: v0.3.14 release-snapshot validation design
Previous checkpoint: docs/132-clean-working-snapshot-runner-validation-v0.md

## 1. Purpose

This document defines the first design for a `release-snapshot` profile in the CarbonStack Go validation runner.

The goal is to move from clean working snapshot validation toward formal runner-backed release package validation.

The profile should validate a release-like package layout before calling the existing core backbone validation path.

This is design-only.

This does not implement the profile.

This does not publish a release.

This does not prove deployability.

## 2. Background

v0.3.13 validated that a clean copied working snapshot can run the Go validation runner outside the live Git working repos.

That validation passed on:

    WSL Debian first;
    Windows second.

The runner currently has:

    doctor;
    core;
    full.

`full` currently aliases `core`.

The next step is to design a release-snapshot validation profile so future releases can be validated by a clear runner-backed process rather than manual command choreography.

## 3. Design decision

Add a future runner profile:

    release-snapshot

Intended command shape:

    go run . --profile release-snapshot --root <release-package-root>

The profile should:

    run release package checks;
    verify expected release layout;
    verify required docs and metadata;
    verify forbidden artifacts are absent before tests;
    then call the existing core validation path;
    report post-test artifacts;
    return nonzero on failure.

The profile should layer checks before validation, similar to the current model:

    doctor -> core
    release-snapshot checks -> core

## 4. Formal release package root

The release-snapshot profile should be designed for a formal package root, not only the current live sibling repo layout.

Preferred future release package root shape:

    <release-root>/
      carbonstack/
      carbonstack-comms/
      carbonstack-cypher/
      release/
        manifest.json
        checksums.txt
        validation-freeze.md
        testing-runbook.md
        release-notes.md
        LICENSE

This formal package root can evolve, but the profile should not assume the user's personal machine paths.

It should infer or receive the root through:

    --root <release-package-root>

Compatibility note:

    The profile may initially support the existing sibling repo layout with release metadata in the carbonstack repo, but the design target should be a formal package root.

## 5. Required repo folders

Required folders:

    carbonstack/
    carbonstack-comms/
    carbonstack-cypher/

Optional for current v0.3.x release-snapshot validation:

    carbonstack-os/

Reason:

    carbonstack-os is part of the long-term project, but not part of the current runnable v0.3.x backbone proof.

## 6. Required repo files

Required carbonstack files:

    carbonstack/README.md
    carbonstack/LICENSE
    carbonstack/docs/README.md
    carbonstack/roadmap/ROADMAP.md
    carbonstack/tools/carbonstack-validate/go.mod
    carbonstack/tools/carbonstack-validate/main.go
    carbonstack/tools/carbonstack-validate/README.md

Required carbonstack-comms files:

    carbonstack-comms/README.md
    carbonstack-comms/LICENSE
    carbonstack-comms/go.mod
    carbonstack-comms/internal/protocol/mls/openmls-sidecar/Cargo.toml

Required carbonstack-cypher files:

    carbonstack-cypher/README.md
    carbonstack-cypher/LICENSE
    carbonstack-cypher/go.mod
    carbonstack-cypher/go.sum
    carbonstack-cypher/migrations/001_init.sql

Optional / include-if-present:

    carbonstack-comms/go.sum

Reason:

    carbonstack-comms/go.sum has not existed in current project state and must not be treated as required until dependency state actually creates it.

## 7. Required release metadata files

Preferred formal release metadata files:

    release/manifest.json
    release/checksums.txt
    release/validation-freeze.md
    release/testing-runbook.md
    release/release-notes.md
    release/LICENSE

If the release package is still using attached assets rather than an internal `release/` folder, the profile should eventually support explicit file paths or a known extracted-assets layout.

For first implementation, it is acceptable to check only a minimal release metadata set:

    manifest.json
    checksums.txt
    testing-runbook.md or validation-freeze.md

The design should keep this flexible until the release packaging format is finalized.

## 8. Required docs semantics

The release package should include a human-readable validation entry point.

It should explain:

    what this release is;
    what platform/toolchains were validated;
    how to run the Go runner;
    what passing validation proves;
    what passing validation does not prove;
    where generated artifacts may appear after tests;
    what artifacts must not be present before tests.

It must not imply:

    production readiness;
    production E2EE;
    hostile-server safety;
    metadata privacy;
    Debian deployability;
    systemd readiness;
    cloudflared readiness;
    external audit;
    certification.

## 9. Pre-test artifact policy

Before running core validation, release-snapshot must scan for forbidden generated/private/build artifacts.

Forbidden before tests:

    .git/
    target/
    .carbonstack-openmls-sidecar-state/
    .go-cache/
    .go-tmp/
    provider-storage.json
    signer.json
    *.db
    *.db-shm
    *.db-wal
    *.exe
    *.test.exe
    Thumbs.db
    .DS_Store

Pre-test hits should be failures for release-snapshot validation unless explicitly allowlisted.

This is stricter than the current core runner scan, which reports hits but does not fail.

Reason:

    A release package should not ship generated/private/build artifacts.

## 10. Post-test artifact policy

After validation, generated artifacts are expected.

Allowed post-test generated roots:

    carbonstack-comms/internal/protocol/mls/openmls-sidecar/target
    carbonstack-comms/internal/protocol/mls/openmls-sidecar/.carbonstack-openmls-sidecar-state

Allowed only under `.carbonstack-openmls-sidecar-state`:

    provider-storage.json
    signer.json

Allowed only under `target` on Windows:

    *.exe

Not allowed after tests:

    generated/private/build artifacts outside known generated roots;
    database files in repo roots;
    provider/signer files outside sidecar generated state;
    executable/test artifacts outside sidecar target or normal external tool caches.

## 11. Toolchain policy

release-snapshot should call doctor or equivalent environment reporting before running core.

It should report:

    OS;
    architecture;
    Go path/version;
    Rust path/version;
    Cargo path/version;
    sqlite3 path/version if available;
    release package root;
    runner path;
    inferred repo paths.

Known Rust findings:

    Debian apt rustc 1.85.0 failed with OpenMLS 0.8.1.
    WSL Debian rustup stable rustc/cargo 1.96.0 passed.
    Windows rustc/cargo 1.95.0 passed current tests despite the runner warning that it is below the WSL known-good 1.96.0 floor.

The first release-snapshot implementation should keep Rust version below 1.96 as a warning, not a hard failure, unless a known-bad version is detected or tests fail.

Known bad:

    rustc 1.85.0 for current OpenMLS 0.8.1 path.

Known good in current testing:

    WSL Debian rustc 1.96.0
    Windows rustc 1.95.0

## 12. Validation flow

Intended release-snapshot profile flow:

    print profile header;
    run package root detection;
    run required folder checks;
    run required file checks;
    run release metadata checks;
    run toolchain/doctor checks;
    run strict pre-test artifact scan;
    call core validation;
    run post-test artifact scan;
    summarize result;
    print allowed claim and nonclaims;
    return nonzero on failure.

## 13. Relationship to core and full

`core` should remain focused on current backbone validation:

    doctor;
    pre-test artifact scan;
    OpenMLS real-Cypher lifecycle;
    Comms package tests;
    Cypher package tests;
    post-test artifact scan.

`release-snapshot` should perform release/package checks and then call `core`.

`full` should remain an alias for `core` until there are real release/deployability surfaces to include.

Later, after release-snapshot is implemented and stable, `full` may become:

    core + release-snapshot-like checks

only if the command context makes that honest.

## 14. Output expectations

The profile should produce output useful for a tester or release verifier.

It should clearly show:

    package root;
    profile name;
    required files checked;
    toolchains;
    pre-test artifact result;
    core validation result;
    post-test artifact result;
    final pass/fail.

It should not hide subprocess output.

It should stream test output like the current runner.

## 15. Failure behavior

Default behavior:

    stop on first failure;
    print failing step;
    return nonzero.

Future optional behavior:

    --keep-going

Do not implement `--keep-going` before the basic profile is stable.

## 16. Non-goals

release-snapshot must not:

    install dependencies;
    delete artifacts;
    clean build outputs;
    package releases;
    upload releases;
    create checksums;
    mutate release contents;
    configure systemd;
    configure cloudflared;
    bind public services;
    touch the real Debian homelab;
    run production paths;
    make security claims.

## 17. What passing release-snapshot should validate

Passing release-snapshot should validate only:

    the release package has the expected layout;
    required source/docs/metadata files are present;
    forbidden generated/private/build artifacts were not present before tests;
    the current core backbone validation passes from the release package;
    post-test artifacts stay in expected generated roots.

## 18. What passing release-snapshot must not validate

Passing release-snapshot must not validate:

    production readiness;
    production E2EE;
    hostile-server safety;
    metadata privacy;
    Android readiness;
    secure vault/storage;
    Debian deployability;
    systemd readiness;
    cloudflared readiness;
    public ingress;
    real homelab deployment;
    external audit;
    certification.

## 19. Future implementation plan

Recommended v0.3.15 implementation scope:

    add --profile release-snapshot;
    reuse existing root inference / --root;
    add required folder/file checks;
    add strict pre-test artifact failure behavior;
    call core after package checks;
    preserve non-destructive post-test scan;
    keep output explicit and boring.

Recommended after v0.3.15:

    validate release-snapshot against a clean staged release-like package;
    update public testing docs only after local release-snapshot validation passes;
    do not publish a runner-backed release until uploaded/downloaded assets are verified.

## 20. Future deployability note

After runner-backed validation, release-snapshot validation, public testing docs, and any runner-backed public testing release are stable, v0.3.x should continue into actual backbone deployability planning and local-only implementation.

That deployability work should happen before v0.4.x runtime Comms UX work.

Later, after local deployability is ironed out, CarbonStack should plan deployable server-system mechanics:

    IRC-like standard server systems;
    server join and operator workflows;
    server/admin responsibilities;
    how those responsibilities interact with the doctrine that every start/endpoint is hostile;
    what a server can enforce without becoming a trust root;
    what clients must verify independently.

This later server-system mechanics work should not happen before basic local deployability exists.