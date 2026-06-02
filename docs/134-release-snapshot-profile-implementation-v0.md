# Release Snapshot Profile Implementation v0

Status: implementation / validation record
Phase: v0.3.15 release-snapshot profile implementation
Previous checkpoint: docs/133-release-snapshot-profile-design-v0.md

## 1. Purpose

This document records the first implementation of the Go runner `release-snapshot` profile.

The profile validates a formal release-like package root before calling the existing `core` validation path.

## 2. Implemented profile

    release-snapshot

Intended command shape:

    go run . --profile release-snapshot --root <release-package-root>

## 3. Current package layout

The profile expects a package root containing:

    carbonstack/
    carbonstack-comms/
    carbonstack-cypher/
    release/

## 4. Current checks

The profile checks:

    required repo folders;
    required repo files;
    minimal release metadata;
    strict forbidden pre-test artifacts;
    then calls core validation.

## 5. Required release metadata

Current minimal release metadata checks:

    release/manifest.json
    release/checksums.txt
    release/validation-freeze.md OR release/testing-runbook.md

Optional:

    release/release-notes.md
    release/LICENSE

## 6. Strict pre-test artifact policy

Unlike `core`, `release-snapshot` treats forbidden generated/private/build artifacts before tests as a failure.

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

## 7. Validation

A local formal package-root validation probe was created at:

    release-snapshot-validation/v0.3.15

Validation order:

    WSL Debian first;
    Windows second.

Expected result:

    release layout checks pass;
    strict pre-test artifact scan passes;
    core validation passes;
    post-test generated artifacts stay under known OpenMLS sidecar generated roots.

## 8. Important run-order note

Do not run `release-snapshot` twice in the same extracted root.

The first successful run generates OpenMLS sidecar state and Rust target artifacts.

A second run in the same root should fail strict pre-test artifact scanning because release-snapshot intentionally treats pre-existing generated artifacts as release package contamination.

Use a fresh extraction for each release-snapshot validation run.

## 9. Nonclaims

Passing `release-snapshot` does not prove:

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

## 10. Next work

Recommended next work:

    validate a staged release-like package more formally;
    refine manifest/checksum expectations;
    update public testing docs only after release-snapshot behavior is stable;
    do not publish a runner-backed public release until uploaded/downloaded assets are verified.