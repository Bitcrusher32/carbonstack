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

### Explicit generated-artifact cleanup

By default, artifact scans are non-destructive. This keeps validation honest: generated/private/build artifacts remain visible after tests and must not be mistaken for source files.

When desired, run with explicit cleanup:

    go run . --profile core --clean-generated

`--clean-generated` only removes known generated/build artifact roots currently recognized by the runner:

    carbonstack-comms/internal/protocol/mls/openmls-sidecar/.carbonstack-openmls-sidecar-state
    carbonstack-comms/internal/protocol/mls/openmls-sidecar/target

It does not delete manual local operator DBs, does not touch `$HOME/.local/share/carbonstack/cypher/cypher.db`, does not clean arbitrary untracked files, and does not replace artifact scans.

### local-cypher

Runs the local-only Cypher API lifecycle validation contract:

    required path checks
    pre-local-cypher artifact scan
    temporary Cypher binary build
    temporary isolated SQLite DB
    explicit loopback bind on 127.0.0.1
    invite/device/envelope/ack lifecycle
    restart against the same temporary DB
    persisted state checks after restart
    post-local-cypher artifact scan
    temporary state cleanup

    go run . --profile local-cypher

`local-cypher` is Cypher-only. It is not `local-backbone`, not runtime Comms UX, not public ingress, not systemd/cloudflared, and not a production deployment or security claim.

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
## release-snapshot run-order warning

`release-snapshot` must be run from a fresh extracted or throwaway staged package root.

Do not validate the package source root that will later be archived or published.

A successful `release-snapshot` run calls `core`, and `core` generates OpenMLS sidecar state and Rust build artifacts. If that same package root is archived afterward, the archive will contain forbidden generated/private/build artifacts and should fail strict pre-test validation later.

Correct pattern:

    create clean package source root
    archive it without running release-snapshot inside it
    extract archive into a throwaway validation root
    run release-snapshot from the throwaway extraction
    discard or preserve the throwaway extraction only as validation evidence

Do not run `release-snapshot` twice in the same extraction unless you intentionally expect the second run to fail strict pre-test artifact scanning.
## Release checksum helper profiles

### write-checksums

Writes real SHA-256 checksums for a clean release package root:

    go run . --profile write-checksums --root /path/to/package-root

The checksum file is written to:

    release/checksums.txt

The helper excludes generated/private/build artifacts and excludes `release/checksums.txt` itself.

### verify-checksums

Verifies `release/checksums.txt` against the release package root:

    go run . --profile verify-checksums --root /path/to/package-root

### release-snapshot relationship

`release-snapshot` now verifies real checksums before calling `core`.

The expected flow is:

    create clean package source root
    write release metadata
    run write-checksums against the package source root
    run only non-generating sanity checks
    archive the package source root
    validate from a fresh extraction with release-snapshot

Do not run `release-snapshot` against the package source root intended for archive/publish.