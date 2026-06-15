# CarbonStack Validate

<!-- CS-V061-CURRENT-STATUS:BEGIN -->
## Current status note

Current public release: **CarbonStack v0.6.0 State/UX Boundary Validation Pre-Release**.

Current v0.6.x focus: **OpenMLS message-flow integration and boundary hardening**. The next work should normalize the OpenMLS application-message path, keep Relay Space artifact/onboarding paths separate from normal message inboxes, and preserve strict nonclaims around production security, hostile-server safety, metadata privacy, vault/key storage, Android, CarbonStackOS, and mature messenger UX.

Historical docs may preserve older release anchors as provenance. Treat the latest compressed LogDoc, current release assets, current registry/table, and current repo behavior as authority for active work.
<!-- CS-V061-CURRENT-STATUS:END -->


Status: experimental validation runner
Phase: v0.5.x validation runner / post-v0.5.0 baseline surface

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

Negative-path coverage currently includes the historically preserved v0.3.27 blunder case:

    content_type=carbonstack.message.text.stub.v0
    protocol_version=carbonstack-openmls-sidecar-v0

Expected result:

    HTTP 400
    unsupported_protocol_version

`local-cypher` is Cypher-only. It is not `local-backbone`, not runtime Comms UX, not public ingress, not systemd/cloudflared, and not a production deployment or security claim.

### dev-runtime-openmls

Runs the current dev/pre-alpha OpenMLS application-message runtime CLI smoke proof:

    openmls-send-dev -> Cypher -> openmls-inbox-dev --ack

Command:

    go run . --profile dev-runtime-openmls --clean-generated

This profile is manual-only and live-umbrella-only for now. It requires sibling git checkouts for:

    carbonstack/
    carbonstack-comms/
    carbonstack-cypher/

It wraps:

    carbonstack-comms/scripts/dev-openmls-runtime-smoke.sh

Boundary:

    not local-backbone
    not mature messaging UX
    not deployment
    not release-package validation yet
    not production/security proof
    not included in full

`--clean-generated` is recommended after successful runs because the smoke proof can leave known OpenMLS sidecar generated roots.

This profile intentionally refuses non-git package-like roots for now. Use it from the live umbrella checkout, not from release package roots. `full` remains the release-package validation ladder and does not include `dev-runtime-openmls`.

### dev-runtime-openmls-wrappers

Runs the separate dev/pre-alpha OpenMLS wrapper-bootstrap runtime CLI smoke proof:

    openmls-*-dev bootstrap wrappers -> message-send-dev -> Cypher -> message-inbox-dev --ack

Command:

    go run . --profile dev-runtime-openmls-wrappers --clean-generated

This profile is manual-only and live-umbrella-only for now. It requires sibling git checkouts for:

    carbonstack/
    carbonstack-comms/
    carbonstack-cypher/

It wraps:

    carbonstack-comms/scripts/dev-openmls-runtime-smoke-wrappers.sh

Boundary:

    not local-backbone
    not mature messaging UX
    not deployment
    not release-package validation yet
    not production/security proof
    not included in full
    does not replace dev-runtime-openmls yet

The existing `dev-runtime-openmls` profile remains the direct-sidecar smoke baseline. This wrapper profile exists as a separate maturity surface while bootstrap wrappers and message wrappers continue to harden.

`--clean-generated` is recommended after successful runs because the smoke proof can leave known OpenMLS sidecar generated roots.

This profile intentionally refuses non-git package-like roots for now. Use it from the live umbrella checkout, not from release package roots. `full` remains the release-package validation ladder and does not include `dev-runtime-openmls-wrappers`.



### relay-openmls-join-dev

Runs the current positive-path Relay Space OpenMLS join development validation profile:

    KeyPackage -> add-member -> Welcome -> join

Command:

    go run . --profile relay-openmls-join-dev --compact-summary

This profile is live-umbrella/dev-oriented for now.

It validates two positive-path subruns:

    no-ack;
    ACK_AFTER_JOIN.

Boundary:

    not local-backbone;
    not production secure messaging;
    not identity verification;
    not hostile-server safety;
    not metadata privacy;
    not audit or certification;
    not included in full;
    not included in release-snapshot;
    not front-door public CLI.

`--compact-summary` is console output/evidence convenience only. It does not preserve generated artifacts and does not change validation scope.

As of v0.5.66, relay-openmls-join-dev has already been deliberately kept outside full/release-snapshot for v0.6.0. Registry presence is not promotion.

v0.5.66 release-profile decision:

    relay-openmls-join-dev remains a manual/dev live-umbrella validation profile for v0.6.0.
    It is not included in full.
    It is not included in release-snapshot.
    It is not a package-root release validation profile yet.
    It is a future full-profile candidate only after repeated clean runs, package-root rehearsal, artifact behavior review, and an explicit release-profile inclusion decision.

Current rationale:

    The profile is useful and bounded: it uses runner-owned temp roots, a temp Cypher DB, unique sidecar labels, no-ack and ACK_AFTER_JOIN subruns, trust/candidate absence checks, DB assertions, and compact evidence output.
    However, it still deliberately checks for a live git umbrella and depends on OpenMLS sidecar generated state behavior.
    Therefore it should remain outside full/release-snapshot for v0.6.0.


### full

Runs the current release-package validation ladder:

    release-snapshot
    local-cypher

`release-snapshot` already calls `core`, so `full` does not call `core` a second time.

Recommended release-package command:

    go run . --profile full --root /path/to/release-package-root --clean-generated

`full` is intended for fresh extracted or throwaway staged release package roots. It is not a deployment command, not `local-backbone`, not runtime Comms UX, and not a production/security claim.

## Expected layout

The runner expects sibling repos:

    carbonstack/
    carbonstack-comms/
    carbonstack-cypher/

It can infer the umbrella root when launched from inside the `carbonstack` repo, including from:

    carbonstack/tools/carbonstack-validate

You can also pass an explicit umbrella root:

    go run . --profile core --root /path/to/carbonstack_umbrella

## Windows note

Windows validation is not the current mainline release-prep target. Prefer Debian / WSL Debian for current runner-backed validation unless a release runbook explicitly says otherwise.

## WSL Debian example

    . "$HOME/.cargo/env"
    cd "$HOME/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate"
    go test ./... -count=1
    go run . --profile doctor
    go run . --profile core --clean-generated
    go run . --profile local-cypher

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
## v0.5.0 release validation recommendation

For the v0.5.0 Runtime and Registry Validation Minor Epoch Pre-Release, use the release-attached runbook and assets. The intended package validation shape is:

    go run . --profile verify-checksums --root /path/to/release-package-root
    go run . --profile full --root /path/to/release-package-root --clean-generated

This verifies the release package checksums, runs the release package/checksum/core validation path through `release-snapshot`, then runs the Cypher-only `local-cypher` lifecycle validation. It remains a validation ladder, not a deployment command.

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