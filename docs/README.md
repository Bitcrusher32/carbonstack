# CarbonStack docs

<!-- CS-V061-CURRENT-STATUS:BEGIN -->
## Current status note

Current public release: **CarbonStack v0.6.0 State/UX Boundary Validation Pre-Release**.

Current v0.6.x focus: **OpenMLS message-flow integration and boundary hardening**. The next work should normalize the OpenMLS application-message path, keep Relay Space artifact/onboarding paths separate from normal message inboxes, and preserve strict nonclaims around production security, hostile-server safety, metadata privacy, vault/key storage, Android, CarbonStackOS, and mature messenger UX.

Historical docs may preserve older release anchors as provenance. Treat the latest compressed LogDoc, current release assets, current registry/table, and current repo behavior as authority for active work.
<!-- CS-V061-CURRENT-STATUS:END -->


This directory is a living documentation archive.

Lower-numbered documents preserve historical design, research, validation, implementation, release-process, and failure-recovery state. They may contain stale assumptions because they intentionally record the project at the time they were written.

For current public testing and release status, start with:

    top-level README.md
    roadmap/ROADMAP.md
    registry/COMMAND_BOUNDARY_TABLE.v0.md
    current Gitea release page and attached release runbook

Current public release:

    v0.5.0 Runtime and Registry Validation Minor Epoch Pre-Release

Current mainline state:

    Mainline development has moved beyond the v0.5.0 public release package.
    The current working direction after v0.5.68A is a pre-v0.6.0 public-surface, artifact-hygiene, and package-rehearsal arc.

    v0.5.49-v0.5.68A completed a concentrated Relay Space, OpenMLS, state-boundary, command-boundary, and registry-surface hardening sequence:
        no-ack live evidence;
        ACK_AFTER_JOIN live evidence;
        relay-openmls-join-dev positive-path validation profile;
        compact summary output;
        negative helper / empty-scalar hardening;
        live negative-path ownership matrix;
        Comms no-ack/failure coverage matrix;
        add-member sidecar-failure command test.

    This is still not local-backbone.
    This is still not production secure messaging.
    This is still not hostile-server safety.
    This is still not metadata privacy.
    This is still not verified identity.
    This is still not secure vault/key storage.
    This is still not mature public send/inbox UX.

Current release-package validation shape remains release-specific. For the v0.5.0 release package, use the release-attached runbook and assets. The general release-package pattern is:

    go run . --profile verify-checksums --root <package-root>
    go run . --profile full --root <package-root> --clean-generated

Current live mainline development validation shape:

    cd ~/repos/carbonstack_umbrella/carbonstack-comms
    go test ./... -count=1

    cd ~/repos/carbonstack_umbrella/carbonstack-cypher
    go test ./... -count=1

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go test ./... -count=1
    go run . --profile doctor

Additional profile runs may be useful during development, but they are not automatically public release-package validation:

    go run . --profile core --clean-generated
    go run . --profile local-cypher
    go run . --profile dev-runtime-openmls --clean-generated
    go run . --profile dev-runtime-openmls-wrappers --clean-generated
    go run . --profile relay-openmls-join-dev --compact-summary

## Historical archive note

Most numbered docs are historical records, not the current release front door.

Do not assume an older numbered document supersedes the latest release/runbook docs unless a newer document explicitly says so.

## How to read this folder

The numbered docs are chronological.

Lower-numbered docs may be stale.

That is intentional.

Older docs preserve the reasoning, constraints, mistakes, pivots, implementation context, and decision history that led to later work. They should not be treated as the current release surface unless a newer release document explicitly points to them.

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

## Current-vs-historical rule

Use the newest relevant release docs, current-state docs, README files, and runbooks for current behavior.

Current behavior should be inferred from:

    latest release page and release runbook for public package validation;
    newest LogDoc / breakpoint handoff for active mainline state;
    newest roadmap refresh for forward planning;
    newest README surfaces for public/project orientation;
    newest numbered docs for a specific subsystem or result.

## Recent current-state docs

Recent docs that are especially relevant to current v0.5.x state include:

    docs/168-v0.5.2-state-trust-vault-pq-preliminary-recon-v0.md
    docs/169-v0.5.3-storage-trust-provider-state-inventory-v0.md
    docs/170-v0.5.4-storage-domain-model-v0.md
    docs/171-v0.5.5-trust-state-model-v0.md
    docs/172-v0.5.6-provider-state-linkage-plan-v0.md
    docs/173-v0.5.8-provider-trust-report-contract-v0.md
    docs/174-v0.5.9-provider-trust-report-exposure-decision-v0.md
    docs/175-v0.5.10-provider-originated-trust-history-append-plan-v0.md
    docs/176-v0.5.14-provider-identity-candidate-import-plan-v0.md
    docs/177-v0.5.15-mapped-provider-identity-mismatch-plan-v0.md
    docs/178-v0.5.16-relay-space-architecture-decision-v0.md
    docs/179-v0.5.17-local-backbone-feasibility-reassessment-v0.md
    docs/180-v0.5.18-implementation-priority-decision-v0.md
    docs/181-v0.5.21-candidate-review-update-priority-decision-v0.md
    docs/182-v0.5.25-reset-recovery-reenrollment-decision-v0.md
    docs/183-v0.5.27-post-recovery-classifier-priority-decision-v0.md
    docs/184-v0.5.30-local-backbone-blocker-reassessment-v0.md
    docs/185-v0.5.31-relay-space-join-invite-member-planning-v0.md
    docs/186-v0.5.32-provider-live-flow-boundary-v0.md
    docs/187-v0.5.33-validation-profile-boundary-v0.md
    docs/188-v0.5.34-local-backbone-go-no-go-v0.md
    docs/189-v0.5.35-implementation-readiness-checkpoint-v0.md
    docs/190-v0.5.38-cypher-relay-space-api-result-v0.md
    docs/191-v0.5.39-relay-space-scoped-envelope-routes-result-v0.md
    docs/192-v0.5.42-roadmap-refresh-anchor-v0.md
    docs/193-v0.5.48-narrow-join-smoke-proof-runbook-v0.md
    docs/194-v0.5.50-ack-after-join-smoke-evidence-v0.md
    docs/195-v0.5.50-ack-after-join-smoke-evidence-and-validation-profile-preflight-v0.md
    docs/196-v0.5.51-validation-profile-design-contract-v0.md
    docs/197-v0.5.56-live-negative-path-design-matrix-v0.md

Some files in this list may not exist in older checkouts or release packages. Treat the latest mainline checkout and latest LogDoc as authoritative for active development state.

## Security and maturity warning

Many docs describe experimental, pre-release, or dev-scaffold behavior.

CarbonStack is not production-certified.

CarbonStack has not received senior external audit or security certification.

Historical docs must not be used to imply production readiness, hostile-server safety, Android readiness, metadata privacy, complete E2EE product status, or verified identity unless a current release document explicitly says so.
- `200-v0.5.65-runtime-send-inbox-convergence-decision-v0.md` — Runtime send/inbox convergence decision: legacy warning-only surfaces; `openmls-send-dev` / `openmls-inbox-dev` are current canonical dev OpenMLS runtime message paths.
- `204-v0.6.5-message-wrapper-implementation-v0.md` — documents the v0.6.5 dev/pre-alpha message wrapper implementation boundary.

- `206-v0.6.10-integrated-runtime-policy-v0.md` — defines the planned `integrated-runtime-dev` policy boundary before implementation.

- `207-v0.6.11-command-reference-generation-policy-v0.md` — records why command reference generation is deferred behind registry hardening.

- `208-v0.6.12-registry-metadata-hardening-v0.md` — records the registry metadata hardening pass before generated command-reference work.
- `../registry/COMMAND_REFERENCE.v0.md` — generated dev/operator command reference from the command registry.
- `209-v0.6.14-post-history-rewrite-sanity-v0.md` — post-history-rewrite sanity checkpoint, release/tag sanity, generated-reference validation, and local backup freeze cleanup.
- `210-v0.6.15-same-state-integrated-proof-plan-v0.md` — same-state integrated proof recon decision and v0.6.16 implementation/probe plan.
- `211-v0.6.16-same-state-integrated-dev-profile-v0.md` — committed same-state integrated dev profile proof boundary and nonclaims.
- `212-v0.6.17-same-state-message-failure-dev-profile-v0.md` — first same-state normal-message failure profile and no-ack/no-drain boundary.
- `213-v0.6.18-same-state-message-unsupported-dev-profile-v0.md` — same-state unsupported normal-message content-type no-ack/no-drain profile boundary.
- `214-v0.6.19-same-state-message-recipient-failure-dev-profile-v0.md` — same-state wrong recipient/device/sidecar no-false-success profile boundary.
- `215-v0.6.20-welcome-join-partial-state-safety-v0.md` — failed Welcome join sidecar partial-state safety patch boundary.
- `216-v0.6.21-same-state-welcome-join-failure-dev-profile-v0.md` — same-state corrupt Welcome join no-ack/no-drain/no-state-poison profile boundary.
- `217-v0.6.22-sidecar-state-authority-classification-v0.md` — sidecar conversation/provider state authority classification after stale-state recon.
- `218-v0.6.23-same-state-message-malformed-payload-dev-profile-v0.md` — same-state malformed normal-message payload no-open/no-ack/no-drain profile boundary.
