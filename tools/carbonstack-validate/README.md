# CarbonStack Validate

Status: experimental validation runner
Phase: v0.3.12 runner hardening / docs integration

This is the Go-based umbrella validation runner for CarbonStack.

It is intended to replace shell-specific umbrella validation over time while still calling repo-local tests.

## Profiles

### doctor

Reports environment, inferred repo layout, required paths, executable paths, and toolchain versions.

    go run . --profile doctor

### core

Runs the current core validation path:

    doctor
    pre-test artifact scan
    targeted OpenMLS real-Cypher lifecycle test
    full carbonstack-comms package tests
    full carbonstack-cypher package tests
    post-test artifact scan

    go run . --profile core

### full

Currently aliases `core`.

    go run . --profile full

`full` should remain a simple alias until later v0.3.x work creates real release/deployability validation surfaces.

## Expected layout

The runner expects sibling repos:

    carbonstack/
    carbonstack-comms/
    carbonstack-cypher/

It can infer the umbrella root when launched from inside the `carbonstack` repo, including from:

    carbonstack/tools/carbonstack-validate

You can also pass an explicit umbrella root:

    go run . --profile core --root /path/to/carbonstack_umbrella

## Windows example

    cd C:\Users\udaiv\repos\carbonstack_umbrella\carbonstack\tools\carbonstack-validate
    go run . --profile doctor
    go run . --profile core

## WSL Debian example

    . "$HOME/.cargo/env"
    cd "$HOME/carbonstack-wsl/carbonstack/tools/carbonstack-validate"
    go run . --profile doctor
    go run . --profile core

## Rust toolchain note

OpenMLS 0.8.1 failed under Debian apt rustc 1.85.0 during v0.3.9.

rustup stable rustc/cargo 1.96.0 passed under WSL Debian during v0.3.9.

The runner reports Rust/Cargo paths and versions, but it does not install or mutate toolchains.

## Artifact scan behavior

Artifact scans are non-destructive.

Pre-test hits may indicate source/copy hygiene problems.

Post-test hits are expected only when they stay in known generated roots such as the OpenMLS sidecar `target/` and `.carbonstack-openmls-sidecar-state/`.

## Boundaries

This runner does not prove production readiness, production E2EE, hostile-server safety, metadata privacy, Debian deployability, systemd readiness, cloudflared readiness, audit, or certification.

It does not install dependencies, delete artifacts, package releases, publish releases, configure services, or deploy anything.
### release-snapshot

Validates a formal release-like package root before calling `core`.

    go run . --profile release-snapshot --root /path/to/release-package-root

The package root is expected to contain:

    carbonstack/
    carbonstack-comms/
    carbonstack-cypher/
    release/

The profile checks required repo files, release metadata, and fails if forbidden generated/private/build artifacts are present before tests.

After package checks pass, it calls `core`.

`release-snapshot` does not package, upload, deploy, clean, install dependencies, or make security claims.