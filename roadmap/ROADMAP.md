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

