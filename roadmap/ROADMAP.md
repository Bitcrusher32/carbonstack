# Current roadmap status

Current near-term target:

    v0.3.20 runner-backed testing release is the current preferred public testing release.

Primary validation target:

    Debian / WSL Debian.

Secondary final explicit Windows dev/test validation for this phase:

    Windows 11 short-path validation with CARGO_BUILD_JOBS=1.

Platform direction after v0.3.20:

    mainline public dev/test releases migrate to Debian / WSL Debian first;
    Windows 11 remains available through the older v0.3.4 PowerShell testbed and the v0.3.20 short-path validation record;
    future Windows/BSD/Linux-family port work may be reconsidered after server/backbone maturity.

Next after the v0.3.20 runner-backed release:

    actual local-only backbone deployability planning and implementation;
    Debian/server-like layout first;
    no cloudflared/public ingress/systemd/homelab deployment claims until separately validated.

Current nonclaims:

    not production-ready;
    not production E2EE;
    not hostile-server safe;
    not metadata-private;
    not deployment-ready;
    not externally audited;
    not certified.


# CarbonStack Roadmap

This roadmap describes the current public direction of CarbonStack.

Older numbered docs may preserve older plans. Use this file and release-specific runbooks for current public-facing direction.

## Current state

CarbonStack is in an experimental backbone phase.

The current validated artifact is a local Cypher + Comms OpenMLS relay proof:

- OpenMLS KeyPackage relay;
- OpenMLS Welcome relay;
- OpenMLS application-message relay;
- real local Cypher server;
- Comms relay helper;
- payload metadata validation;
- consume-then-ack semantics;
- repeatable smoke harness.

This is not a finished messenger.

This is not production-certified.

This is not externally audited.

This is not Android-ready.

## Near-term path

### v0.2.61 - Runbook and docs archive framing

Complete.

The current runbook is:

    docs/113-experimental-backbone-deployability-runbook-v0.md

### v0.2.62 - Public-surface cleanup

Current target.

Scope:

- README cleanup;
- quickstart/runbook cleanup;
- component repo public-surface cleanup;
- stale public claims removal;
- direct, reader-facing language;
- strict security nonclaims.

### v0.2.63 - Known-good validation matrix

Goal:

- keep the current local harness path repeatable;
- tighten known-good validation;
- remove misleading helper names or stale test-only framing where necessary.

### v0.2.64 - Envelope lifecycle standardization

Goal:

- standardize envelope lifecycle language;
- clarify queued/acked states;
- document payload metadata;
- ensure schema/API docs match implementation.

### v0.2.65+ - Self-test, release hardening, packaging, and v0.3.0 freeze

Goal:

- plan a user-visible CLI/dev harness path;
- expose a repeatable self-test surface without pretending it is a production messenger.

### Pre-v0.3.0 release-hardening checkpoint

Goal:

- repository cleanup;
- release-facing README hardening;
- stale-claims sweep;
- known-good validation;
- security disclaimer hardening;
- clear component map;
- no production certification claims.

### v0.3.0 - Experimental backbone epoch

Goal:

- publish CarbonStack as an experimental runner-backed backbone test surface, while keeping actual server deployability as later v0.3.x work;
- use `carbonstack` as the release front door;
- link component repos;
- describe the concrete validated artifact as the Cypher + Comms OpenMLS relay backbone;
- keep the release clearly pre-alpha and non-certified.

## Long-term path

Later work may include:

- runtime Comms send/inbox OpenMLS integration;
- trust-state mapping from sidecar/provider events;
- hostile-server rollback/replay harnesses;
- metadata minimization design;
- secure local vault/storage;
- external review and audit preparation;
- Android/Pixel development;
- CarbonStackOS appliance prototyping.

Android and CarbonStackOS work are not the current near-term target.
## v0.3.0 release README

The consolidated v0.3.0 release README is planned at:

    docs/v0.3.0-minor-epoch-release.md

That document is the public release entrypoint for the experimental backbone epoch.
## Post-v0.3.0 minor-epoch direction

v0.3.x should focus on backbone maturation and release verification:

    release asset verification
    clean snapshot extraction checks
    self-test/runbook UX cleanup
    portability recon
    low-level deploy configuration
    backbone hardening

Future security/safety claims require a dedicated adversarial-validation or self-penetration-testing minor epoch before they are made.

That future epoch must document:

    what was tested
    methodology
    tested version/deployment context
    what the test validates
    what the test does not validate
    residual risk
    whether the claim is allowed, narrowed, deferred, or rejected

Self-testing, AI-assisted testing, and community attempts are not external audits unless a real scoped external audit occurs.
## Public identity after v0.3.0

CarbonStack's end goal is the secure-communications appliance stack.

The current verified release is narrower:

    v0.3.0 / v0.3.x experimental backbone

That backbone is the current public proof surface. It should not be described as the finished appliance stack, a production messenger, or a certified secure product.
## v0.3.x portability sequence

v0.3.x should finish Windows/public release basics before Linux portability work.

Current intended sequence:

    public tester runbook and self-test UX cleanup;
    Windows assumptions / portability recon;
    Debian homelab recon;
    low-level deploy config and operator model.

Debian target context:

    x86_64 Debian homelab;
    currently hosts Gitea and WordPress services;
    cloudflared is part of the current public ingress setup;
    future recon should inspect actual host setup before proposing deploy changes.

Debian/Linux portability is not yet validated.
## WSL Debian bridge before homelab validation

Before touching the real Debian homelab, v0.3.x should use WSL Debian as a quick portability triage layer.

The intended sequence is:

    Windows public release basics
    Windows assumptions / portability recon
    WSL Debian quick-portability triage
    real Debian homelab local-only validation
    low-level Cypher deploy config and operator model

WSL Debian is not a deployment proof.

It can identify Linux path/toolchain/script issues before the real homelab is used.

The real Debian homelab remains the meaningful future target for deployability recon because it has systemd, active services, cloudflared ingress, and existing operator constraints.
## v0.3.8 WSL Debian setup scout result

v0.3.8 established Debian under WSL as the quick-portability bridge before real Debian homelab validation.

Observed baseline:

    Debian 13 trixie under WSL2
    linux/amd64
    git available
    Go available
    Rust/Cargo available
    sqlite3 available

The first WSL test rung should use current working repo snapshots, run direct Go/Rust tests first, and avoid PowerShell wrappers until the Linux validation gap is understood.

Long-term direction remains a small Go validation runner so Windows and Linux release validation can converge.
## v0.3.9 WSL Debian quick-portability test result

v0.3.9 validated the current Comms/Cypher OpenMLS backbone core under Debian WSL after correcting the Rust toolchain.

Finding:

    Debian apt Rust 1.85.0 was too old for OpenMLS 0.8.1.
    rustup stable 1.96.0 worked.

Passed under WSL Debian after rustup stable:

    targeted OpenMLS real-Cypher lifecycle test;
    full carbonstack-comms/internal/protocol package;
    full carbonstack-comms package suite;
    full carbonstack-cypher package suite.

The remaining portability gap is mostly validation orchestration, not obvious Go/Rust core behavior.

Next direction:

    design a Go-based cross-platform full-test wrapper;
    avoid letting separate PowerShell and bash validation flows become long-term tech debt;
    keep real Debian homelab testing deferred until deployability / IRC-style setup work requires it.
## v0.3.10 Go validation runner design

v0.3.10 records the decision to converge validation around an umbrella Go runner.

Initial intended location:

    carbonstack/tools/carbonstack-validate

Reason:

    repo-local tests already exist;
    the umbrella should orchestrate functional tests across repos;
    Windows and WSL Debian validation should converge;
    PowerShell and bash should become convenience wrappers, not validation authorities.

Initial profiles:

    doctor
    core
    full

Implementation should start small:

    environment/toolchain report;
    sibling repo path checks;
    Cypher package tests;
    Comms package tests;
    targeted OpenMLS real-Cypher lifecycle test;
    non-destructive artifact scan.

WSL Debian may become the preferred fast core-validation environment after the Go runner passes on both WSL Debian and Windows.
## v0.3.12 Go validation runner hardening

v0.3.12 hardens the Go validation runner and records the first runner validation path across Windows and WSL Debian.

Current runner location:

    carbonstack/tools/carbonstack-validate

Current profiles:

    doctor
    core
    full

`full` currently aliases `core`.

The runner now acts as the future umbrella validation authority candidate. PowerShell and bash should become thin wrappers over the runner rather than independent validation authorities.
## v0.3.13 clean working snapshot runner validation

v0.3.13 validates that a clean copied working snapshot can run the Go validation runner outside the live Git working repos.

Result:

    WSL Debian clean snapshot runner validation passed.
    Windows clean snapshot runner validation passed.
    Pre-test artifact scans were clean.
    Post-test artifacts stayed in expected OpenMLS sidecar generated roots.

Next direction:

    design release-snapshot validation;
    then implement release-snapshot validation;
    then update public testing docs around runner-backed validation;
    then consider a future runner-backed public testing release;
    then proceed to actual backbone deployability work inside late v0.3.x.
## v0.3.14 release-snapshot profile design

v0.3.14 defines the future Go runner `release-snapshot` profile.

Intended behavior:

    validate formal release package layout;
    check required repo folders and metadata;
    reject forbidden generated/private/build artifacts before tests;
    call core validation;
    report post-test artifacts;
    preserve explicit nonclaims.

This remains design-only. Implementation comes later.

The v0.3.x sequence after this remains:

    implement release-snapshot profile;
    validate a staged release-like package;
    update runner-backed public testing docs;
    rehearse runner-backed release staging;
    verify uploaded/downloaded release assets;
    optionally cut a runner-backed public testing release;
    then proceed into actual backbone deployability planning and local-only implementation;
    later, after deployability is ironed out, plan deployable server-system mechanics such as IRC-like CarbonStack standard server systems and server/admin workflows under the hostile-endpoint assumption.
## v0.3.15 release-snapshot profile implementation

v0.3.15 implements the Go runner `release-snapshot` profile.

The profile validates package layout and strict pre-test artifact hygiene before calling `core`.

Next direction:

    validate staged release-like package behavior;
    refine manifest/checksum expectations;
    update runner-backed public testing docs after validation stabilizes;
    then rehearse runner-backed release staging and uploaded/downloaded asset verification;
    then move toward late-v0.3.x actual backbone deployability work.
## v0.3.16 release-snapshot validation hardening

v0.3.16 validates the release-snapshot profile against a staged release-like package from fresh extractions.

Important rule:

    do not validate the package source root that will later be archived or published.

Correct flow:

    create clean package source root;
    archive it before running release-snapshot inside it;
    validate from fresh throwaway extraction;
    use WSL Debian first;
    use Windows second.

Next direction:

    refine runner-backed public testing docs;
    define real checksum/manifest semantics;
    rehearse runner-backed release staging;
    verify uploaded/downloaded release assets;
    only then consider an optional public runner-backed testing release.
## v0.3.17 runner-backed public testing docs

v0.3.17 documents the future runner-backed public testing path.

This does not supersede the current public v0.3.4 Windows/PowerShell testing release.

The documented future path is:

    release package with carbonstack, carbonstack-comms, carbonstack-cypher, and release metadata;
    fresh extraction;
    WSL Debian validation first;
    Windows validation second;
    Go runner release-snapshot profile;
    strict pre-test artifact scan;
    core validation;
    post-test artifact interpretation.

Next direction:

    define real manifest/checksum semantics;
    rehearse runner-backed release staging;
    verify uploaded/downloaded release assets;
    only then consider an optional public runner-backed testing release.
## v0.3.18 runner-backed release staging rehearsal

v0.3.18 rehearses a local runner-backed release-candidate package.

It adds Go runner checksum helper profiles:

    write-checksums
    verify-checksums

The release-snapshot profile now verifies real checksums before calling core validation.

The local rehearsal package includes:

    carbonstack
    carbonstack-comms
    carbonstack-cypher
    release metadata
    real checksum coverage
    validation runbook
    explicit nonclaims

carbonstack-os is recorded as related but not included because it is future appliance-stack work, not part of the current runnable v0.3.x backbone validation package.

Next direction:

    uploaded/downloaded release asset verification;
    then optional public runner-backed testing release only if verification is clean;
    then late-v0.3.x actual backbone deployability planning and local-only implementation.
## v0.3.19 local downloaded-asset simulation

v0.3.19 locally simulates the public downloaded-asset flow before touching the Gitea release surface.

The local simulation shape is:

    published-assets/
    downloads/
    extracted/package/

The package root must contain only:

    carbonstack
    carbonstack-comms
    carbonstack-cypher
    release

The validation path remains:

    verify-checksums
    release-snapshot

A separate release-verify profile is deferred until after v0.3.21+ maturity, when deployable backbone and IRC-style join/admin mechanics are more mature.

Next direction:

    full public-release docs cleanup/check;
    final release page text draft;
    final in-release testing runbook;
    WSL Debian and Windows known-good platform tests from fresh package extractions;
    then choose the cleaner public release boundary, v0.3.19 or v0.3.20.
### v0.3.19 Windows validation caveat

The v0.3.19 local downloaded-asset simulation passed on WSL Debian and Windows.

Windows initially failed from a long nested release-verification path during Rust/MSVC linking with LNK1104 while creating OpenMLS sidecar build-script executables.

The known-good Windows pass used:

    C:\cs-v0319-win2\package
    CARGO_BUILD_JOBS=1

Windows known-good validation should use a short extraction root until this path/toolchain/file-lock behavior is better characterized.
## v0.3.20 runner-backed testing release cleanup

v0.3.20 is the first public runner-backed testing release.

Primary validation target:

    Debian / WSL Debian

Final explicit Windows dev/test validation for this phase:

    Windows 11
    short extraction root
    CARGO_BUILD_JOBS=1

v0.3.20 is expected to be the final release in this phase where Windows dev/test validation is explicitly provided as part of the runner-backed release surface.

After v0.3.20, mainline public dev/test releases should migrate to Debian / WSL Debian first.

v0.3.4 remains available as the older Windows 11 / PowerShell testbed.

Next direction after the runner-backed release is clean:

    actual local-only backbone deployability planning and implementation;
    Debian/server-like layout first;
    Windows no longer mainline for deployability unless a later mature port effort justifies it.

### v0.3.22 — Cypher local operator surface recon

v0.3.22 records a docs/recon-only Cypher local operator surface pass for the WSL Debian-first local-only deployability line. It introduces the CarbonStack Relay Space terminology, inspects Cypher config/database/migration/API surfaces, records Comms-to-Cypher addressing seams, preserves runtime Comms OpenMLS UX as v0.4.x work, and identifies persistent SQLite migration behavior as the highest-priority deployability risk before a stronger local operator profile.

### v0.3.23 — Cypher migration persistence recon

v0.3.23 records a proof-first recon of CarbonStackCypher repeated SQLite migration behavior. A temporary test ran the current migration path twice against the same SQLite database, then removed the test file. The result confirms that fresh DB migration remains the known-good path while persistent local DB migration/restart behavior needs `schema_migrations`, idempotent guards, or explicit wipe-only experimental documentation before stronger local operator claims.


