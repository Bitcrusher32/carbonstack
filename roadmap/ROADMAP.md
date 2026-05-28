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

### v0.2.63 â€” Option C completion for testing

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

### v0.2.65+ â€” Option B planning and implementation

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
