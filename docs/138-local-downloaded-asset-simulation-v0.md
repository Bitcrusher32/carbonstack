# Local Downloaded-Asset Simulation and Public Release Readiness Plan v0

Status: local verification simulation / release-readiness planning
Phase: v0.3.19 local downloaded-asset simulation
Previous checkpoint: docs/137-runner-backed-release-staging-rehearsal-v0.md

## 1. Purpose

This document records v0.3.19 local downloaded-asset simulation and public release readiness planning.

The goal was to simulate the public-user asset flow locally before touching the Gitea release surface.

This is intentionally conservative.

This is not a public release.

This is not a Gitea upload/download verification.

This is not deployability validation.

## 2. Why local simulation first

Hotfix releases like v0.3.4 are sometimes unavoidable, but the goal is to avoid hotfix flooding.

Before publishing a runner-backed public release, CarbonStack should prove the release asset flow locally, then only move to the public release surface when package structure, checksums, validation, docs, and platform wording are stable.

## 3. Simulated release flow

The v0.3.19 local simulation used this structure:

    release-verification/v0.3.19-local-download-simulation/
      published-assets/
      downloads/
      extracted/
        package/

Interpretation:

    published-assets/ simulates the release-page asset store.
    downloads/ simulates what a tester receives.
    extracted/package/ is the fresh package root used for validation.

The carrier archive is kept outside the package root.

The package root must contain only:

    carbonstack/
    carbonstack-comms/
    carbonstack-cypher/
    release/

## 4. Validation commands

From the fresh package root's runner module:

    cd carbonstack/tools/carbonstack-validate
    go test ./...
    go run . --profile verify-checksums --root <package-root>
    go run . --profile release-snapshot --root <package-root>

## 5. Windows validation

Windows validation used a fresh local extracted package root. The final known-good Windows pass used a short extraction root to avoid long-path/toolchain/file-lock instability observed in the first Windows attempt.

Expected and observed behavior:

    package-root top-level directory shape was correct;
    package-root top-level file scan was clean;
    pre-run forbidden artifact scan was clean;
    verify-checksums passed;
    release-snapshot passed;
    core validation passed;
    post-test generated artifacts stayed under known OpenMLS sidecar generated roots.

## 5a. Windows short-path validation lesson

Initial Windows validation from a long nested release-verification path reached checksum verification and release-snapshot pre-core checks, but failed during Rust/MSVC linking with LNK1104 while creating build_script_build executables under the OpenMLS sidecar target tree.

This appeared to be a Windows toolchain/path/file-lock/environment issue rather than a CarbonStack package or checksum failure.

A short-path Windows validation root was then used:

    C:\cs-v0319-win2\package

With:

    CARGO_BUILD_JOBS=1

The short-path Windows validation passed:

    verify-checksums;
    release-snapshot layout checks;
    release metadata checks;
    strict pre-test artifact scan;
    release checksum verification;
    OpenMLS real-Cypher lifecycle;
    carbonstack-comms package tests;
    carbonstack-cypher package tests.

Windows known-good validation should use a short extraction root until the path/file-lock behavior is better characterized.
## 6. WSL Debian validation

WSL Debian validation used a fresh local extracted package root.

Expected and observed behavior:

    package-root top-level directory shape was correct;
    package-root top-level file scan was clean;
    pre-run forbidden artifact scan was clean;
    verify-checksums passed;
    release-snapshot passed;
    core validation passed;
    post-test generated artifacts stayed under known OpenMLS sidecar generated roots.

## 7. Current public release policy

v0.3.4 remains the current recommended public testing release until a runner-backed public release is fully prepared and verified.

The next public runner-backed release should not be rushed.

Before public release, complete:

    local downloaded-asset simulation;
    Windows known-good validation;
    WSL Debian known-good validation;
    release metadata review;
    checksum/manifest review;
    docs cleanup;
    release page text review;
    public README/front-door consistency check;
    explicit claim/nonclaim review.

## 8. Version policy

v0.3.19 is a mainline local simulation/planning checkpoint.

The eventual public runner-backed release can be v0.3.19 or v0.3.20 depending on which version boundary is cleaner.

Use the cleaner public release version, not the most compressed one.

## 9. release-verify deferral

A future separate `release-verify` runner profile may make sense later, but not yet.

It should be deferred until after v0.3.21+ maturity, when IRC-style join/admin mechanics and the deployable backbone have enough mature code/docs to justify a dedicated public release-verification profile.

For now:

    verify-checksums
    release-snapshot

are enough.

## 10. carbonstack-os status

carbonstack-os remains related but not included in current runnable v0.3.x package validation.

Once individual platform repos mature into real service-stack components, separate Cypher, Comms, and OS releases can be considered.

That is too early now.

## 11. Allowed claim

Allowed:

    A local downloaded-asset simulation passed checksum verification and release-snapshot validation from fresh package extractions on Windows and WSL Debian.

## 12. Nonclaims

This does not prove:

    public release readiness;
    Gitea upload/download integrity;
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

## 13. Next work

Recommended next work:

    perform full public-release readiness docs cleanup/check;
    prepare final release page text draft;
    prepare final in-release testing runbook;
    rerun WSL Debian and Windows known-good platform tests from fresh package extractions;
    only then decide whether public runner-backed release should be v0.3.19 or v0.3.20.