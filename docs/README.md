# CarbonStack docs

This directory is a living documentation archive.

Lower-numbered documents preserve historical design, research, validation, and release-process state. They may contain stale assumptions because they intentionally record the project at the time they were written.

For current public testing and release status, start with:

    docs/139-runner-backed-testing-release-cleanup-v0.md

Current release direction:

    v0.3.20 runner-backed testing release candidate
    Debian / WSL Debian as the primary validation target
    Windows 11 short-path validation as the final explicit Windows dev/test validation for this phase
    v0.3.4 remains available as the older Windows 11 / PowerShell testbed

After v0.3.20, mainline public dev/test releases are expected to migrate to Debian / WSL Debian first.

Current validation boundary:

    checksum verification
    release-snapshot validation
    OpenMLS real-Cypher lifecycle
    carbonstack-comms package tests
    carbonstack-cypher package tests

Current nonclaims:

    not production-ready
    not production E2EE
    not hostile-server safe
    not metadata-private
    not deployment-ready
    not systemd-ready
    not cloudflared-ready
    not externally audited
    not certified

## Historical archive note

Most numbered docs are historical records, not the current release front door.

Do not assume an older numbered document supersedes the latest release/runbook docs unless the newer docs explicitly say so.

## How to read this folder

The numbered docs are chronological.

Lower-numbered docs may be stale.

That is intentional.

Older docs preserve the reasoning, constraints, mistakes, pivots, implementation context, and decision history that led to later work. They should not be treated as the current release surface unless a newer release document explicitly points to them.

## Current-vs-historical rule

Use the newest relevant release docs, current-state docs, README files, and runbooks for current behavior.

Use older numbered docs for:

    project history;
    design rationale;
    implementation context;
    continuity preservation;
    debugging old decisions;
    understanding why a later result exists.

Do not assume an old plan document still describes current behavior.

Do not assume an old result document still describes the latest release surface.

Do not rewrite old historical docs just because implementation moved forward.

## Release-specific docs

Release-specific docs will be added as the project matures.

A release-specific doc or runbook should say what is current for that release, what is proven, what is not proven, which component repos are involved, and which commands are known-good.

For example, a future v0.3.0 release surface may include docs such as:

    v0.3.0 experimental backbone release notes;
    v0.3.0 known-good runbook;
    v0.3.0 security status and nonclaims;
    v0.3.0 component map.

Those release docs should be treated as the public-facing source of truth for that release.

## Why this archive is preserved

CarbonStack is being developed through staged research, implementation, validation, and cleanup rungs.

Destroying or constantly rewriting the archive would erase important process context.

Instead, the project preserves old docs as historical artifacts and adds newer docs to clarify current state.

This means the docs folder is not a single polished manual.

It is a continuity archive plus release documentation surface.

## Security and maturity warning

Many docs describe experimental, pre-release, or dev-scaffold behavior.

CarbonStack is not production-certified.

CarbonStack has not received senior external audit or security certification.

Historical docs must not be used to imply production readiness, hostile-server safety, Android readiness, metadata privacy, or complete E2EE product status unless a current release document explicitly says so.
## Current v0.3.0 packaging freeze

The v0.3.0 release packaging freeze is recorded at:

    docs/119-v0.3.0-release-packaging-freeze-v0.md

Use it with:

    docs/v0.3.0-minor-epoch-release.md
## Post-v0.3.0 verification and governance

Current post-release verification and governance docs:

    docs/120-v0.3.0-post-release-verification-v0.md
    docs/121-security-claim-validation-policy-v0.md
    docs/122-logdoc-case-study-sanitization-plan-v0.md
## Sanitized project LogDoc list

A sanitized project LogDoc archive is available at:

    sanitized-project-logdoc-list/

This folder contains sanitized LogDoc material for workflow/case-study use. It is separate from the main numbered docs archive and should not be treated as the current release source of truth.
## Clean release snapshot self-test recon

The v0.3.3 clean release snapshot self-test recon is recorded at:

    docs/123-v0.3.0-clean-snapshot-self-test-recon-v0.md

This recon checks whether the v0.3.0 release source snapshots can run validation from extraction rather than from the active working repos.
## Release snapshot self-test fix

The v0.3.4 source-snapshot self-test fix is recorded at:

    docs/124-v0.3.4-release-snapshot-self-test-fix-v0.md

This follows the v0.3.3 clean snapshot recon and records the fix for source snapshots that do not include Git metadata.
## Public v0.3.4 release asset verification

The v0.3.5 public v0.3.4 release asset verification is recorded at:

    docs/125-v0.3.4-public-release-asset-verification-v0.md

This verifies the uploaded v0.3.4 Gitea release assets, not just local release staging.
## Public tester runbook

The v0.3.4 public tester runbook is recorded at:

    docs/126-v0.3.4-public-tester-runbook-v0.md

This is the version-specific tester runbook for the v0.3.4 experimental backbone release. v0.3.4 remains available as the older Windows 11 / PowerShell testbed; for current status, start from the top-level README and current release/runbook docs.
## Windows reliance and Debian homelab recon

The v0.3.7 Windows reliance and Debian homelab recon is recorded at:

    docs/127-windows-reliance-and-debian-homelab-recon-v0.md

This records the current Windows/PowerShell validation boundary, Debian homelab target context, and the planned WSL Debian quick-portability bridge before real Debian validation.
## WSL Debian quick-portability setup scout

The v0.3.8 WSL Debian setup scout is recorded at:

    docs/128-wsl-debian-quick-portability-setup-scout-v0.md

This records the Debian WSL setup and toolchain baseline before running Linux-side quick-portability tests.
## WSL Debian quick-portability test

The v0.3.9 WSL Debian quick-portability test is recorded at:

    docs/129-wsl-debian-quick-portability-test-v0.md

This records the first WSL Debian core validation pass, including the Rust toolchain floor finding and the passing Comms/Cypher OpenMLS backbone tests after rustup stable was used.
## Go validation runner design

The v0.3.10 Go validation runner design is recorded at:

    docs/130-go-validation-runner-design-v0.md

This defines the umbrella Go runner direction for converging Windows and WSL Debian validation without maintaining separate PowerShell and bash validation authorities.
## Go validation runner implementation

The Go validation runner implementation and first cross-platform validation record is recorded at:

    docs/131-go-validation-runner-implementation-v0.md

This records the `tools/carbonstack-validate` runner, doctor/core/full profiles, and validation under both Windows PowerShell and WSL Debian.
## Clean working snapshot runner validation

The v0.3.13 clean working snapshot runner validation is recorded at:

    docs/132-clean-working-snapshot-runner-validation-v0.md

This records clean copied working snapshot validation through the Go runner on both WSL Debian and Windows.
## Release-snapshot profile design

The v0.3.14 release-snapshot profile design is recorded at:

    docs/133-release-snapshot-profile-design-v0.md

This defines the intended Go runner profile for validating formal release-like packages before implementing the profile.
## Release-snapshot profile implementation

The v0.3.15 release-snapshot profile implementation is recorded at:

    docs/134-release-snapshot-profile-implementation-v0.md

This records the first Go runner `release-snapshot` profile, package checks, strict pre-test artifact policy, and validation path.
## Release-snapshot validation hardening

The v0.3.16 release-snapshot validation hardening record is recorded at:

    docs/135-release-snapshot-validation-hardening-v0.md

This records staged release-like package validation from fresh extractions on WSL Debian and Windows, plus the run-order rule that release-snapshot must not be run inside the package source root that will later be archived.
## Runner-backed public testing path

The v0.3.17 runner-backed public testing path draft is recorded at:

    docs/136-runner-backed-public-testing-path-v0.md

This documents the Go runner release-snapshot public testing path while preserving v0.3.4 as the older Windows 11 / PowerShell testbed until the runner-backed release surface is cut and published.
## Runner-backed release staging rehearsal

The v0.3.18 runner-backed release staging rehearsal is recorded at:

    docs/137-runner-backed-release-staging-rehearsal-v0.md

This records a local release-candidate rehearsal package with real Go-generated checksums and release-snapshot validation from fresh extractions on WSL Debian and Windows.
## Local downloaded-asset simulation

The v0.3.19 local downloaded-asset simulation and public release readiness plan is recorded at:

    docs/138-local-downloaded-asset-simulation-v0.md

This records a local public-user-style asset flow using published-assets, downloads, and fresh package extraction roots before touching the Gitea release surface.
## Runner-backed testing release cleanup

The v0.3.20 runner-backed testing release cleanup and final validation plan is recorded at:

    docs/139-runner-backed-testing-release-cleanup-v0.md

This prepares the final release surface for the first runner-backed testing release, with Debian / WSL Debian as the primary validation target and Windows 11 short-path validation as the final explicit Windows dev/test validation for this phase.