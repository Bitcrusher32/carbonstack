# Clean Working Snapshot Runner Validation v0

Status: validation record
Phase: v0.3.13 clean working snapshot runner validation
Previous checkpoint: docs/131-go-validation-runner-implementation-v0.md

## 1. Purpose

This document records the v0.3.13 clean working source snapshot validation.

The goal was to test whether a clean copied snapshot of the current working repos can run the Go validation runner outside the live Git working trees.

This is not a public release.

This is not release-snapshot profile validation.

This is not Debian deployability validation.

## 2. Snapshot source

A clean working snapshot archive was created from the current local repos:

    carbonstack
    carbonstack-comms
    carbonstack-cypher

The snapshot excluded:

    .git
    target
    .carbonstack-openmls-sidecar-state
    .go-cache
    .go-tmp
    provider-storage.json
    signer.json
    *.db
    *.db-shm
    *.db-wal
    *.exe
    *.test.exe
    Thumbs.db
    .DS_Store

Archive path:

    clean-snapshot-validation/v0.3.13/carbonstack-v0.3.13-clean-working-snapshot.tgz

## 3. WSL Debian validation

WSL clean snapshot root:

    ~/carbonstack-v0.3.13-clean

WSL preflight artifact scan:

    PASS

No forbidden generated/private/build artifacts were present before running the runner.

WSL toolchains:

    go version go1.24.4 linux/amd64
    rustc 1.96.0
    cargo 1.96.0
    sqlite3 3.46.1

WSL runner commands:

    go test ./...
    go run . --profile doctor
    go run . --profile core
    go run . --profile full

WSL result:

    PASS

The runner inferred the clean snapshot umbrella root correctly:

    ~/carbonstack-v0.3.13-clean

WSL core validation passed:

    targeted OpenMLS real-Cypher lifecycle;
    full carbonstack-comms package tests;
    full carbonstack-cypher package tests.

WSL post-run artifact scan found expected generated artifacts under the OpenMLS sidecar generated roots:

    carbonstack-comms/internal/protocol/mls/openmls-sidecar/target
    carbonstack-comms/internal/protocol/mls/openmls-sidecar/.carbonstack-openmls-sidecar-state

Nested provider-storage.json and signer.json files were also present under the generated sidecar state root after tests.

These are expected post-test artifacts.

## 4. Windows validation

Windows clean snapshot root:

    clean-snapshot-validation/v0.3.13/windows-clean

Windows preflight artifact scan:

    PASS

No forbidden generated/private/build artifacts were present before running the runner.

Windows toolchains:

    go version go1.26.3 windows/amd64
    rustc 1.95.0
    cargo 1.95.0
    sqlite3 unavailable in PATH

SQLite absence is currently a warning, not a failure.

The runner warned that Windows rustc 1.95.0 appears older than the WSL known-good Rust floor of 1.96.0 for the current OpenMLS sidecar tests.

This warning did not block validation because the tested Windows clean snapshot still passed.

Windows runner commands:

    go test ./...
    go run . --profile doctor
    go run . --profile core
    go run . --profile full

Windows result:

    PASS

The runner inferred the clean snapshot umbrella root correctly:

    clean-snapshot-validation/v0.3.13/windows-clean

Windows core validation passed:

    targeted OpenMLS real-Cypher lifecycle;
    full carbonstack-comms package tests;
    full carbonstack-cypher package tests.

Windows post-run artifact scan found expected generated artifacts under the OpenMLS sidecar generated roots:

    carbonstack-comms/internal/protocol/mls/openmls-sidecar/target
    carbonstack-comms/internal/protocol/mls/openmls-sidecar/.carbonstack-openmls-sidecar-state

On Windows, the sidecar target root includes many Cargo-generated .exe build artifacts.

These are expected only when nested under the known OpenMLS sidecar target root after tests.

## 5. Artifact interpretation correction

Earlier wording that expected only the OpenMLS sidecar target and state roots was correct but too compressed.

The more precise rule is root-scoped:

    provider-storage.json is acceptable only when nested under the known generated .carbonstack-openmls-sidecar-state test root after validation;
    signer.json is acceptable only when nested under the known generated .carbonstack-openmls-sidecar-state test root after validation;
    *.exe files are acceptable only when nested under the known OpenMLS sidecar target root after Windows Cargo validation;
    these artifacts are not acceptable if present before tests in clean snapshots;
    these artifacts are not acceptable if scattered outside known generated roots.

## 6. Interpretation

v0.3.13 validates that the current clean working snapshot can run the Go validation runner outside the live Git repos on both WSL Debian and Windows.

This is a major step toward runner-backed release validation.

It does not yet validate a formal release-snapshot profile.

It does not publish a release.

It does not prove Debian deployability.

## 7. Allowed claim

Allowed:

    A clean copied working snapshot passed the Go validation runner doctor/core/full path on WSL Debian and Windows at v0.3.13.

More specifically:

    WSL Debian clean snapshot validation passed with rustup stable Rust/Cargo 1.96.0.
    Windows clean snapshot validation passed with Rust/Cargo 1.95.0 despite a runner warning that it is below the current WSL known-good Rust floor.
    Pre-test artifact scans were clean.
    Post-test generated artifacts stayed under expected OpenMLS sidecar generated roots.

## 8. Nonclaims

This does not prove:

    public Linux release support;
    Debian deployability;
    systemd readiness;
    cloudflared readiness;
    real homelab validation;
    production readiness;
    production E2EE;
    hostile-server safety;
    metadata privacy;
    external audit;
    certification.

## 9. Next work

Recommended next work:

    design the release-snapshot validation profile;
    keep full as a core alias until real release/deployability surfaces exist;
    preserve WSL Debian as first validation environment;
    preserve Windows as second confirmation before breakpoints/releases;
    do not touch the real Debian homelab yet;
    do not publish a runner-backed release until release-snapshot validation exists.