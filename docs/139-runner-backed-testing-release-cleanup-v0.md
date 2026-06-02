# Runner-Backed Testing Release Cleanup and Final Validation Plan v0

Status: release-readiness cleanup / final validation plan
Phase: v0.3.20 runner-backed testing release cleanup
Previous checkpoint: docs/138-local-downloaded-asset-simulation-v0.md

## 1. Purpose

This document records the v0.3.20 release-readiness cleanup plan for the first runner-backed CarbonStack testing release.

The goal is to clean the release surface, prepare final release text/runbook material, rebuild a final release-candidate package, and validate that package before publishing.

If final validation is clean, v0.3.20 may become the public runner-backed testing release.

If final validation exposes new release-surface or platform problems, v0.3.20 remains a release-readiness checkpoint and public release moves to a later version.

## 2. Current public release status

Current recommended public testing release before v0.3.20:

    v0.3.4

v0.3.4 remains the older Windows 11 / PowerShell testbed release.

v0.3.20 is intended to prepare and, if clean, publish the first runner-backed testing release.

## 3. Platform status for v0.3.20

Primary validated target:

    Debian 13 / WSL2 Debian, linux/amd64

Secondary explicit dev/test validation:

    Windows 11, windows/amd64

Windows validation requirements:

    use a short extraction root such as C:\cs-v0320
    set CARGO_BUILD_JOBS=1 for Rust/OpenMLS validation

Windows migration note:

    v0.3.20 is expected to be the final release in this phase where Windows dev/test validation is explicitly provided as part of the runner-backed release surface.

After v0.3.20, CarbonStack mainline public dev/test releases should migrate to Debian / WSL Debian first.

This does not mean Windows is permanently unsupported.

Later Linux-family, BSD, or Windows ports may be reconsidered after the server/backbone stack is mature.

## 4. Release package shape

The v0.3.20 runner-backed testing package should use:

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

carbonstack-os remains related but not included in the current runnable v0.3.x backbone validation package.

## 5. Required runner profiles

The validation surface remains:

    verify-checksums
    release-snapshot

A separate release-verify profile is deferred until after v0.3.21+ maturity, when deployable backbone and IRC-style join/admin mechanics create a larger public release surface.

## 6. Final validation order

Preferred final validation order:

    WSL Debian first
    Windows short-path second

WSL Debian should be treated as the primary validation path.

Windows should be treated as secondary final explicit dev/test validation for this phase.

## 7. Final WSL Debian command shape

From a fresh package extraction:

    . "$HOME/.cargo/env"
    cd <package-root>/carbonstack/tools/carbonstack-validate
    go test ./...
    go run . --profile verify-checksums --root <package-root>
    go run . --profile release-snapshot --root <package-root>

## 8. Final Windows command shape

Use a short root:

    C:\cs-v0320\package

Set:

    CARGO_BUILD_JOBS=1

From a fresh package extraction:

    cd C:\cs-v0320\package\carbonstack\tools\carbonstack-validate
    go test ./...
    go run . --profile verify-checksums --root C:\cs-v0320\package
    go run . --profile release-snapshot --root C:\cs-v0320\package

Do not use deeply nested release-verification paths for final Windows Rust/OpenMLS validation.

## 9. Release-surface cleanup checklist

Before publishing, review:

    carbonstack README front door;
    docs/README.md;
    roadmap/ROADMAP.md;
    release/manifest.json;
    release/release-notes.md;
    release/testing-runbook.md;
    release/validation-freeze.md;
    release/checksums.txt;
    public release page text;
    allowed claims;
    nonclaims;
    platform target wording;
    Windows migration note;
    v0.3.4 legacy Windows testbed wording.

## 10. Allowed release claim if final validation passes

Allowed:

    CarbonStack v0.3.20 is a pre-alpha / experimental runner-backed testing release for the current Cypher + Comms OpenMLS relay backbone.

    The v0.3.20 release package passed checksum verification and release-snapshot validation from fresh package extractions on Debian / WSL Debian and Windows 11 short-path validation.

## 11. Nonclaims

This release does not prove:

    production readiness;
    production E2EE;
    hostile-server safety;
    metadata privacy;
    Android readiness;
    secure vault/storage;
    public ingress safety;
    Debian deployability;
    systemd readiness;
    cloudflared readiness;
    real homelab validation;
    external audit;
    certification.

## 12. Next work after public runner-backed release

After v0.3.20 release work is clean, the next v0.3.x work should move toward actual local-only backbone deployability:

    local server-like layout;
    Cypher bind/config/data-dir policy;
    Comms-to-Cypher config;
    start/stop lifecycle;
    logs/errors;
    runner validation for local deployment profile.

Still out of scope initially:

    cloudflared;
    public ingress;
    systemd service;
    real homelab deployment;
    shared production service paths;
    MariaDB;
    production claims.