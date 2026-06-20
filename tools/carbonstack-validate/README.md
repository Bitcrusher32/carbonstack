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




### registry-lookup

Prints a dev/operator registry entry by registry ID or literal command.

Examples:

    go run . --profile registry-lookup --registry-id comms.message-send-dev
    go run . --profile registry-lookup --command "go run ./cmd/comms message-send-dev"

The lookup output is a classification aid. It may print command text, audience, maturity, lifecycle status, source path, validation surface, related surfaces, and nonclaims.

Boundary:

    registry presence is not promotion;
    not generated public manual;
    not production UX;
    not security certification.

### integrated-runtime-dev

Runs the first dev/pre-alpha integrated runtime composition profile:

    relay-openmls-join-dev -> dev-runtime-openmls-wrappers

Meaning:

    Relay onboarding proof:
      KeyPackage -> add-member -> Welcome -> join

    then normal message wrapper proof:
      openmls-*-dev bootstrap wrappers -> message-send-dev -> Cypher -> message-inbox-dev --ack

Run:

    go run . --profile integrated-runtime-dev --root ~/repos/carbonstack_umbrella --clean-generated

Boundary:

    live umbrella only;
    not full;
    not release-snapshot;
    not package-root validation;
    not production secure messaging;
    not hostile-server safety;
    not local-backbone;
    not mature messenger UX.

This first implementation composes existing validated dev profiles in sequence. It does not yet claim a same-state/same-conversation package-root release proof. The individual profiles remain separately callable.

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

Registry lookup enrichment examples:

    go run . --profile registry-lookup --list
    go run . --profile registry-lookup --list --lifecycle-status recommended_dev_wrapper
    go run . --profile registry-lookup --list --audience dev --maturity dev_only
    go run . --profile registry-lookup --list --missing-nonclaims

These remain dev/operator classification helpers. They are not a generated public manual, production UX, or security certification.

Relay onboarding registry lookup examples:

    go run . --profile registry-lookup --list --lifecycle-status relay_onboarding_artifact_transport
    go run . --profile registry-lookup --list --lifecycle-status relay_onboarding_artifact_bridge
    go run . --profile registry-lookup --list --lifecycle-status relay_onboarding_join_transition_candidate

These classify Relay Space OpenMLS onboarding primitives. They are not wrapper commands, not normal message inbox UX, and not production enrollment or membership claims.

## v0.6.10 integrated runtime policy note

`integrated-runtime-dev` is the preferred future live-dev profile name for an in-series proof that chains Relay onboarding with the recommended `message-*` normal-message wrapper flow.

As of the Stage 2 implementation patch, the profile is implemented as a live-umbrella composition profile.

Policy:

- keep `full` as release-package validation;
- keep `release-snapshot` as package-root validation;
- keep `relay-openmls-join-dev`, `dev-runtime-openmls-wrappers`, and `dev-runtime-openmls` individually callable;
- do not make Relay onboarding artifacts ordinary `message-inbox-dev` messages;
- do not introduce `full-runtime-dev` until integrated runtime has matured and an explicit later decision is made.

## v0.6.11 command reference generation policy

A generated command reference is intentionally deferred.

The runner registry lookup profile can inspect registry rows, but it is not yet a generated public manual. Before adding a generated Markdown reference, the registry needs stronger metadata for risky dev/internal surfaces, especially flags, nonclaims, validation surfaces, lifecycle statuses, and root/package/live-umbrella boundaries.

`integrated-runtime-dev` remains a live-umbrella dev profile and must not be folded into `full` by documentation.

## v0.6.12 registry metadata hardening

The registry metadata hardening pass improves future command-reference readiness while keeping generation deferred.

Validation remains registry-oriented:

    go run . --profile registry-lookup --list --missing-nonclaims
    go run . --profile registry-lookup --registry-id runner.integrated-runtime-dev

`registry-lookup` is still an inspection tool, not a generated public manual.

## Generated command-reference freshness

The runner test suite includes a generated-reference freshness guard.

    go test ./...

The guard runs:

    python3 tools/registry/render-command-reference.py --check

from the `carbonstack` repo root. If `registry/commands.v0.yaml` changes without regenerating `registry/COMMAND_REFERENCE.v0.md`, the runner tests should fail.

Boundary:

- this is a docs/registry freshness guard;
- it is not runtime validation;
- it does not generate Unix man pages;
- it does not make the generated reference general-public UX documentation.

## same-state-integrated-dev

`same-state-integrated-dev` is a live-umbrella dev/pre-alpha validation profile.

    go run . --profile same-state-integrated-dev --root <umbrella root> --clean-generated

It proves Relay onboarding and normal message send/open/ack inside one coherent temp universe:

    KeyPackage -> add-member -> Welcome -> join -> message-send-dev -> message-inbox-dev --ack

Boundary:

- not `full`;
- not `release-snapshot`;
- not release-package validation;
- not package-root validation;
- not production secure messaging;
- not hostile-server safety;
- not metadata privacy;
- not identity verification;
- not secure enrollment;
- not mature messenger UX.

`integrated-runtime-dev` remains the sequential composition profile. `same-state-integrated-dev` is the stronger same-conversation proof profile.

## same-state-message-failure-dev

`same-state-message-failure-dev` is a live-umbrella dev/pre-alpha validation profile.

    go run . --profile same-state-message-failure-dev --root <umbrella root> --clean-generated

It proves the first same-state normal-message failure rule:

    wrong-conversation message-inbox-dev --ack must not ack and must not drain the inbox.

After the failed wrong-conversation open attempt, the profile opens the same queued message with the correct conversation and asserts the normal explicit ack path still works.

Boundary:

- not `full`;
- not `release-snapshot`;
- not release-package validation;
- not package-root validation;
- not adversarial relay harness;
- not hostile-server safety;
- not metadata privacy;
- not production secure messaging;
- not production E2EE;
- currently covers wrong-conversation message-open failure only.

`same-state-integrated-dev` remains the positive-path same-conversation proof profile.

## same-state-message-unsupported-dev

`same-state-message-unsupported-dev` is a live-umbrella dev/pre-alpha validation profile.

    go run . --profile same-state-message-unsupported-dev --root <umbrella root> --clean-generated

It proves a same-state normal-message unsupported-envelope rule:

    unsupported normal application-message content_type must not ack and must not drain the inbox.

The profile mutates only the normal application-message envelope content type inside a temporary Cypher DB, runs `message-inbox-dev --ack`, verifies unsupported skip/no-ack/no-drain behavior, restores the original content type, and verifies normal open/ack recovery.

Boundary:

- not `full`;
- not `release-snapshot`;
- not release-package validation;
- not package-root validation;
- not adversarial relay harness;
- not hostile-server safety;
- not metadata privacy;
- not production secure messaging;
- not production E2EE;
- currently covers unsupported normal application-message content_type only.

`same-state-integrated-dev` remains the positive-path same-conversation proof profile. `same-state-message-failure-dev` remains the wrong-conversation no-ack/no-drain profile.

## same-state-message-recipient-failure-dev

`same-state-message-recipient-failure-dev` is a live-umbrella dev/pre-alpha validation profile.

    go run . --profile same-state-message-recipient-failure-dev --root <umbrella root> --clean-generated

It proves a same-state normal-message wrong recipient/device/sidecar rule:

    wrong recipient/device/sidecar attempts must not falsely open, ack, or drain Bob's inbox.

The profile sends a normal message to Bob after same-state Relay join, then checks Alice state + Alice sidecar, Bob state + Alice sidecar, and Bob state + missing sidecar attempts before proving the correct Bob open/ack path still works.

Boundary:

- not `full`;
- not `release-snapshot`;
- not release-package validation;
- not package-root validation;
- not adversarial relay harness;
- not hostile-server safety;
- not metadata privacy;
- not production secure messaging;
- not production E2EE;
- not identity verification;
- currently covers wrong recipient/device/sidecar no-false-success only.

`same-state-integrated-dev` remains the positive-path same-conversation proof profile. Existing failure profiles remain narrow and separate.

## same-state-welcome-join-failure-dev

`same-state-welcome-join-failure-dev` is a live-umbrella dev/pre-alpha validation profile.

    go run . --profile same-state-welcome-join-failure-dev --root <umbrella root> --clean-generated

It proves a fixed Relay onboarding failure invariant:

    corrupt Welcome join fails;
    corrupt Welcome join does not ack;
    corrupt Welcome join does not drain Bob's Relay inbox;
    corrupt Welcome join leaves no final/staging Bob conversation state;
    restored valid Welcome joins and acks with the same conversation label.

Boundary:

- not `full`;
- not `release-snapshot`;
- not release-package validation;
- not package-root validation;
- not adversarial relay harness;
- not hostile-server safety;
- not metadata privacy;
- not production secure messaging;
- not production E2EE;
- not identity verification;
- not stale provider state modeling;
- currently covers corrupt Welcome join failure plus restored-Welcome recovery only.

This is a leaf profile. It can be called by a future `full-runtime-dev` / contextual `full-*` evaluator after explicit aggregation policy exists.

## same-state-message-malformed-payload-dev

`same-state-message-malformed-payload-dev` is a live-umbrella dev/pre-alpha validation profile.

    go run . --profile same-state-message-malformed-payload-dev --root <umbrella root> --clean-generated

It proves malformed normal application-message payloads do not falsely open, ack, drain Bob's queued normal-message inbox, mutate conversation provider storage, or rewrite the envelope. It also proves restored-payload recovery for each malformed case.

Boundary:

- not `full`;
- not `release-snapshot`;
- not release-package validation;
- not package-root validation;
- not adversarial relay harness;
- not hostile-server safety;
- not metadata privacy;
- not production secure messaging;
- not production E2EE;
- not identity verification;
- not replay or duplicate classification.

This is a deterministic failure-hardening profile. It may become an input to a later adversarial harness, but it is not one today.
