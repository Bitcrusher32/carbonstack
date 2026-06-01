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

### v0.2.61 â€” Deployability runbook

Complete.

The current runbook is:

    docs/113-experimental-backbone-deployability-runbook-v0.md

### v0.2.62 â€” Public surface cleanup

Current target.

Scope:

- README cleanup;
- quickstart/runbook cleanup;
- component repo public-surface cleanup;
- stale public claims removal;
- direct, reader-facing language;
- strict security nonclaims.

### v0.2.63 â€” known-good local backbone proof completion for testing

Goal:

- keep the current local harness path repeatable;
- tighten known-good validation;
- remove misleading helper names or stale test-only framing where necessary.

### v0.2.64 â€” Inbox, ack, and schema semantics cleanup

Goal:

- standardize envelope lifecycle language;
- clarify queued/acked states;
- document payload metadata;
- ensure schema/API docs match implementation.

### v0.2.65+ â€” OpenMLS backbone self-test and dev-harness planning

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

### v0.3.0 â€” Experimental backbone epoch

Goal:

- publish CarbonStack as an experimental server-deployable backbone;
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
