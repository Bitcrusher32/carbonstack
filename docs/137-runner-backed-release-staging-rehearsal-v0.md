# Runner-Backed Release Staging Rehearsal v0

Status: staging rehearsal / validation record
Phase: v0.3.18 runner-backed release staging rehearsal
Previous checkpoint: docs/136-runner-backed-public-testing-path-v0.md

## 1. Purpose

This document records v0.3.18 runner-backed release staging rehearsal.

The goal was to build a local release-candidate rehearsal package with a formal package layout, real Go-generated checksum coverage, release metadata, runner-backed validation instructions, and explicit nonclaims.

This is not a public release.

This is not uploaded/downloaded release asset verification.

This is not deployability validation.

## 2. Package shape

The v0.3.18 local release-candidate rehearsal package used:

    <package-root>/
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

## 3. carbonstack-os status

carbonstack-os was recorded in the manifest as related but not included.

Reason:

    carbonstack-os is part of the long-term appliance stack direction but is not part of the current runnable v0.3.x backbone validation package.

## 4. Checksum implementation

v0.3.18 adds Go runner checksum helper profiles:

    write-checksums
    verify-checksums

The checksum file is:

    release/checksums.txt

The checksum helper excludes generated/private/build artifacts, carrier archives, and release/checksums.txt itself.

The release-snapshot profile now verifies real checksums before calling core validation.

## 5. Validation lessons

Paths with spaces in checksum entries required parsing checksum lines by fixed 64-character SHA-256 prefix rather than strings.Fields.

Carrier archives must not be placed inside the package root. Use a downloads/ and package/ split for validation extractions.

Archive suffixes are defensively skipped by checksum collection:

    .tgz
    .tar.gz
    .tar
    .zip

## 6. Run-order rule

Correct flow:

    create clean package source root;
    write release metadata;
    run write-checksums against the package source root;
    run verify-checksums against the package source root;
    run only non-generating sanity checks;
    archive the package source root;
    validate from fresh throwaway extraction.

Do not run release-snapshot against the package source root intended for archive/publish.

## 7. WSL Debian validation

Validation order:

    WSL Debian first.

Expected and observed behavior:

    fresh package extraction used;
    carrier archive kept outside package root;
    pre-run forbidden artifact scan was clean;
    package-root top-level file scan was clean;
    verify-checksums passed;
    release layout checks passed;
    release metadata checks passed;
    strict pre-test artifact scan passed;
    release checksum verification passed;
    core validation passed;
    post-test generated artifacts stayed under known OpenMLS sidecar generated roots.

## 8. Windows validation

Validation order:

    Windows second.

Expected and observed behavior:

    fresh package extraction used;
    carrier archive kept outside package root;
    pre-run forbidden artifact scan was clean;
    package-root top-level file scan was clean;
    verify-checksums passed;
    release layout checks passed;
    release metadata checks passed;
    strict pre-test artifact scan passed;
    release checksum verification passed;
    core validation passed;
    post-test generated artifacts stayed under known OpenMLS sidecar generated roots.

## 9. Allowed claim

Allowed:

    A local runner-backed release-candidate rehearsal package with real Go-generated checksums passed verification and release-snapshot validation from fresh extractions on WSL Debian and Windows.

## 10. Nonclaims

This does not prove:

    public release readiness;
    uploaded/downloaded asset integrity;
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

## 11. Next work

Recommended next work:

    prepare uploaded/downloaded release asset verification workflow;
    verify a future release candidate as a public user would;
    keep v0.3.4 as the current public testing release until a runner-backed release is actually published and verified;
    do not begin deployability until runner-backed release readiness is clean.