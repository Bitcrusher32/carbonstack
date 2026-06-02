# Release Snapshot Validation Hardening v0

Status: validation / hardening record
Phase: v0.3.16 staged release-like package validation
Previous checkpoint: docs/134-release-snapshot-profile-implementation-v0.md

## 1. Purpose

This document records v0.3.16 release-snapshot validation hardening.

The goal was to validate a staged release-like package through the implemented Go runner `release-snapshot` profile, while preserving the lessons from v0.3.15 around stale tarballs and package-root contamination.

This is not a public release.

This is not uploaded/downloaded release asset verification.

This is not deployability validation.

## 2. Current runner behavior

The Go runner now includes:

    doctor
    core
    full
    release-snapshot

`full` still aliases `core`.

`release-snapshot` performs release/package checks and strict pre-test artifact scanning before calling `core`.

## 3. Staged package root shape

The v0.3.16 staged package used this shape:

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

`carbonstack-os` is still not part of the current v0.3.x runnable backbone validation package.

## 4. Run-order rule

Do not run `release-snapshot` inside the package source root before archiving or publishing.

A successful `release-snapshot` run calls `core`.

`core` generates OpenMLS sidecar state and Rust build artifacts.

If the same package root is archived afterward, the archive will contain forbidden generated/private/build artifacts and should fail strict pre-test validation later.

Correct pattern:

    create clean package source root;
    perform only non-generating sanity checks;
    archive the clean package source root;
    extract into a fresh throwaway validation root;
    run release-snapshot from the throwaway extraction;
    discard or preserve the throwaway extraction only as validation evidence.

## 5. Pre-archive sanity checks

Allowed pre-archive checks:

    confirm release-snapshot code exists;
    run go test ./... for the runner module;
    scan for forbidden artifacts.

Do not run:

    go run . --profile release-snapshot --root <package-source-root>

against the package source root that will be archived.

## 6. Checksum policy

The v0.3.16 staged package uses placeholder checksum text only to validate release-snapshot mechanics.

This is acceptable for local profile hardening.

It is not acceptable for a public runner-backed release.

Future runner-backed release staging must define real checksum generation and verification before uploaded/downloaded release asset verification.

## 7. WSL Debian validation

Validation order:

    WSL Debian first.

Expected and observed behavior:

    fresh extraction used;
    pre-run forbidden artifact scan was clean;
    release layout checks passed;
    release metadata checks passed;
    strict pre-test artifact scan passed;
    core validation passed;
    post-test generated artifacts stayed under known OpenMLS sidecar generated roots.

WSL Debian remains the preferred first validation environment.

## 8. Windows validation

Validation order:

    Windows second.

Expected and observed behavior:

    fresh extraction used;
    pre-run forbidden artifact scan was clean;
    release layout checks passed;
    release metadata checks passed;
    strict pre-test artifact scan passed;
    core validation passed;
    post-test generated artifacts stayed under known OpenMLS sidecar generated roots.

Windows remains the second confirmation environment before breakpoints/releases.

## 9. Interpretation

v0.3.16 validates the release-snapshot profile against a staged release-like package from fresh extractions on both WSL Debian and Windows.

This is stronger than v0.3.15's first implementation probe because it explicitly preserves the run-order rule and treats the package source root as non-generating.

This still does not publish a public release.

This still does not verify uploaded/downloaded release assets.

This still does not prove deployability.

## 10. Allowed claim

Allowed:

    A staged local release-like package passed the Go runner release-snapshot profile from fresh extractions on WSL Debian and Windows.

More specifically:

    required package layout and metadata checks passed;
    strict pre-test artifact scans passed;
    core validation passed;
    post-test artifacts stayed under known OpenMLS sidecar generated roots.

## 11. Nonclaims

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

## 12. Next work

Recommended next work:

    refine runner-backed public testing docs and runbook path;
    keep v0.3.4 PowerShell release path historical/current until superseded;
    design real checksum/manifest semantics before a public runner-backed release;
    rehearse runner-backed release staging with real checksums before uploaded/downloaded verification;
    do not publish a new public release yet unless release staging and uploaded/downloaded verification are clean.