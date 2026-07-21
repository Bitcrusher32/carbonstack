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
- `219-v0.6.24-loadcheck-replay-classification-v0.md` — load-check semantics and normal application-message replay classification profile boundary.
- `220-v0.6.25-normal-message-envelope-metadata-classification-v0.md` — normal application-message envelope metadata enforcement and sender-metadata nonclaim classification.
- `221-v0.6.26-sender-metadata-output-warning-v0.md` — normal message inbox output warning that `from_device` is unverified relay envelope metadata, not verified identity.
- `222-v0.6.26-loadcheck-provider-summary-output-v0.md` — load-check output cleanup distinguishing provider reloadability from summary metadata presence.
- `223-v0.6.27-vault-security-nosilent-state-model-v0.md` — vault/security state model defining authority categories, sensitivity classes, no-silent rules, backup/recovery implications, bounded substrate eligibility, and PQ/adversarial dependencies.
- `224-v0.6.28-pq-hybrid-placement-model-v0.md` — PQ/hybrid placement model defining layer options, default new-conversation/protocol-epoch posture, state tags, migration/no-silent rules, backup/recovery implications, compatibility/refusal behavior, Cypher/Relay Space implications, and experiment gates.
- `225-v0.6.29-bounded-state-substrate-mechanics-v0.md` — bounded state substrate mechanics model defining state-audit-dev as the current non-mutating proto-substrate, state-boundary naming, included/excluded surfaces, Cypher inventory-only posture, generated/dev-root treatment, PQ tag reservation, and future validation gates.
- `226-v0.6.29-state-audit-dev-model-alignment-v0.md` — records the narrow v0.6.29C state-audit-dev output/model alignment patch adding explicit state-boundary model, authority/sensitivity, no-silent, Cypher inventory-only, and PQ-reserved-not-implemented labels without adding vault, backup/restore, PQ, or command-surface expansion.
- `227-v0.6.30-validation-profile-naming-boundary-v0.md` — validation profile naming boundary model defining `full` as compatibility release-package validation, `full-validate-release` as the preferred explicit future release-validation name, `release-snapshot` as package-root validation, `full-runtime-dev` as future-only live-dev aggregation, and live-dev/state-audit boundaries before adversarial harness work.
- `228-v0.6.30-full-validate-release-alias-v0.md` — records the narrow v0.6.30C implementation of `full-validate-release` as an exact alias to current `full` release-package validation behavior without adding `full-runtime-dev`, live-dev aggregation, adversarial harness checks, deployment validation, or production-security claims.
- `229-v0.6.31-adversarial-harness-contract-evidence-matrix-v0.md` — defines the adversarial harness contract and evidence matrix vocabulary, separates deterministic same-state evidence seeds from future adversarial harness coverage, preserves `full`/`full-validate-release`/`release-snapshot` boundaries, and stages manual/CLI docs refresh before v0.7.0 release-boundary rehearsal.
- `230-v0.6.32-public-surface-manual-docs-refresh-v0.md` — selected public/manual docs and CLI discoverability refresh after v0.6.32A, including `full-validate-release` help visibility, normal-message wrapper pointers (`message-send-dev` / `message-inbox-dev`), state-audit/non-recovery boundaries, adversarial-contract pointer, and toolchain-floor release-prep note.
- `231-v0.6.33-release-boundary-package-root-rehearsal-plan-v0.md` — defines the v0.7.0 cumulative pre-alpha engineering release-boundary/package-root rehearsal plan, treats v0.6.0 scripts and release shape as prior-art templates, selects fork-first v0.7.0 script strategy, makes `full-validate-release` the primary release-validation wording, and preserves nonclaims before actual package rehearsal.
- `232-v0.6.33-v0.7.0-package-scripts-and-rehearsal-prep-v0.md` — records the v0.7.0 package staging/rehearsal script fork, `full-validate-release` release-validation wording, registry classification for the new package helpers, generated command reference refresh, and nonclaims before actual rehearsal evidence.
- `233-v0.6.34-release-candidate-package-surface-audit-v0.md` — records the v0.7.0 release-candidate package surface audit, staged/fresh-extracted package checks, final manual Gitea release asset folder preparation, and hard nonclaims before tag/asset upload.
- 234-v0.7.1-gate-a-operational-workflow-contract-v0.md — accepted Gate A operational workflow contract covering the preferred lifecycle, explicit workspace context, workflow ownership, future full-operational-spine-dev classification, documented and callable leaf surfaces, Relay-scoped normal messaging, identity-domain separation, mutation classes, structured partial-completion results, acknowledgement rules, Relay membership authority, Cypher/MLS mismatch handling, state/concurrency policy, and validation boundaries.
- 235-v0.7.1-gate-a-closure-hardening-v0.md — formal Gate A closure authority defining explicit client, relay-admin, and dev capability domains; preferred-lifecycle stage ownership and mutation mapping; minimum stage-specific refusal behavior; aggregate-to-leaf relationships; evidence requirements; and Gate A closure assertions.
- `236-v0.7.2-gate-b-b3-relay-scoped-normal-message-path-v0.md` — current authority for explicit Relay Space context, active routing membership prerequisites, scoped ordinary-message send/inbox/ack, unsupported-envelope visibility, and B3b pre-closure posture.

## Current v0.7.4 Gate B4 invite/member lifecycle closure

Current Gate B4 closure authority:

    docs/237-v0.7.4-gate-b-b4-invite-member-lifecycle-closure-v0.md

This document records the implemented Relay Space invite/member lifecycle, including atomic invite claim, active/disabled/left routing-member states, explicit rejoin refusal, member inspection, restart persistence, command/registry surfaces, validation requirements, and hard nonclaims.

Gate B4 is closed. Gate B remains open. B5 KeyPackage lifecycle is next and requires a consolidated deep reconnaissance pass before mutation.

## Current v0.7.5 Gate B5 prerequisite scoped delivery authority

Current prerequisite authority:

    docs/238-v0.7.5-gate-b5-prerequisite-scoped-delivery-authority-v0.md

Scoped Relay Space inbox and scoped ACK now require current active recipient membership. Disabled and left members cannot fetch or acknowledge queued envelopes; queued state remains persisted, restart-safe, and available again only after an allowed explicit reactivation or future rejoin workflow.

This is routing/delivery authority only. B5 KeyPackage inspection, generation, publication, and consumption remain separate leaves.

## Current v0.7.5 Gate B5a KeyPackage inspection and ownership

Canonical decision/implementation contract:

    docs/239-v0.7.5-gate-b5a-keypackage-inspection-ownership-v0.md

B5a establishes a read-only KeyPackage inspection surface, validated
OpenMLS KeyPackage reference identity, lifetime metadata, artifact-integrity
metadata, and local sidecar ownership evidence.

It does not implement repeatable generation, rotation, Relay publication,
consumption, Welcome lifecycle, identity verification, or trust promotion.

B5b is next. A continuity breakpoint will be cut immediately after B5b closes,
before B5c begins.

## Current v0.7.5 Gate B5b repeatable KeyPackage generation and rotation

Canonical result and boundary:

    docs/240-v0.7.5-gate-b5b-keypackage-generation-rotation-v0.md

B5b provides persistent local KeyPackage generation, request idempotence,
legacy adoption, restart-stable inventory, explicit metadata-only retirement,
and device-local concurrency serialization.

It does not publish to Relay, consume or ACK KeyPackages, process Welcome,
or mutate trust/candidate state.

A private LogDoc and Breakpoint JSON are mandatory immediately after B5b
closure. B5c must not begin before that breakpoint is accepted.

## Current v0.7.6 Gate B5c idempotent KeyPackage publication

Canonical result and boundary:

    docs/241-v0.7.6-gate-b5c-keypackage-publication-v0.md

B5c publishes one explicitly selected B5b generation only after B5a inspection.

Cypher persists one sender-scoped publication binding so exact retries return the original envelope ID across response loss, acknowledgement, concurrency, and process restart. Cross-destination reuse and reference/payload disagreement are deterministic conflicts.

The older `openmls-relay-keypackage-submit-dev` path remains compatibility-only and does not provide B5c publication semantics.

B5c does not consume or ACK KeyPackages, run add-member, submit or join Welcome, or mutate identity/trust state. B5d remains next.

- [242-v0.7.7 Gate B5d KeyPackage Consume Receipt](242-v0.7.7-gate-b5d-keypackage-consume-receipt-v0.md)
- [243-v0.7.8 Gate B5e KeyPackage Lifecycle Closure](243-v0.7.8-gate-b5e-keypackage-lifecycle-closure-v0.md)

- [244-v0.7.8 Gate B6 Welcome Lifecycle](244-v0.7.8-gate-b6-welcome-lifecycle-v0.md)
- [245-v0.7.9 Gate B6 Closure](245-v0.7.9-gate-b6-closure-v0.md)

- [246-v0.7.8 Gate B7 Cypher MLS Mismatch](246-v0.7.8-gate-b7-cypher-mls-mismatch-v0.md)
- [247-v0.7.8 Gate B7 Closure](247-v0.7.8-gate-b7-closure-v0.md)

- [248-v0.7.9 Gate B8 Workflow Engine](248-v0.7.9-gate-b8-workflow-engine-v0.md)
- [249-v0.7.9 Gate B8 Closure](249-v0.7.9-gate-b8-closure-v0.md)

- [250-v0.7.10 Gate B9 Integration Closure](250-v0.7.10-gate-b9-integration-closure-v0.md)
- [251-v0.7.10 Gate B Closure](251-v0.7.10-gate-b-closure-v0.md)

- [252-v0.7.12 Gate C1 State Substrate Inventory](252-v0.7.12-gate-c1-state-substrate-inventory-v0.md)
- [253-v0.7.12 Gate C1 Closure](253-v0.7.12-gate-c1-closure-v0.md)

- [254-v0.7.13 Gate C2 State Schema Compatibility](254-v0.7.13-gate-c2-state-schema-compatibility-v0.md)
- [255-v0.7.13 Gate C2 Closure](255-v0.7.13-gate-c2-closure-v0.md)

- [256-v0.7.14 Gate C3 Path Policy](256-v0.7.14-gate-c3-path-policy-v0.md)
- [257-v0.7.14 Gate C3 Closure](257-v0.7.14-gate-c3-closure-v0.md)

- [258-v0.7.15 Gate C4 Write Policy](258-v0.7.15-gate-c4-write-policy-v0.md)
- [259-v0.7.15 Gate C4 Closure](259-v0.7.15-gate-c4-closure-v0.md)

- [260-v0.7.16 Gate C5 State Substrate Closure](260-v0.7.16-gate-c5-state-substrate-closure-v0.md)
- [261-v0.7.16 Gate C Closure](261-v0.7.16-gate-c-closure-v0.md)

- [262-v0.7.17 Gate D Runtime Aggregate](262-v0.7.17-gate-d-runtime-aggregate-v0.md)
- [263-v0.7.17 Gate D Closure](263-v0.7.17-gate-d-closure-v0.md)

- [264-v0.7.18 Gate E E1 Native Deployment](264-v0.7.18-gate-e-e1-native-deployment-v0.md)
- [265-v0.7.18 Gate E E1 Closure](265-v0.7.18-gate-e-e1-closure-v0.md)

- [266-v0.7.19 Gate E E2 Cypher Config Inspection](266-v0.7.19-gate-e-e2-cypher-config-inspection-v0.md)
- [267-v0.7.19 Gate E E2 Closure](267-v0.7.19-gate-e-e2-closure-v0.md)

- [268-v0.7.19 Gate E E3 Native Deployment Closure](268-v0.7.19-gate-e-e3-native-deployment-closure-v0.md)
- [269-v0.7.19 Gate E Closure](269-v0.7.19-gate-e-closure-v0.md)

- [270-v0.7.20 Gate F F1 Release Package Runtime Surface](270-v0.7.20-gate-f-f1-release-package-surface-v0.md)
- [271-v0.7.20 Gate F F1 Closure](271-v0.7.20-gate-f-f1-closure-v0.md)
- [272-v0.7.20 Gate F F2 Manual Private Lifecycle Runbook](272-v0.7.20-gate-f-f2-manual-private-lifecycle-runbook-v0.md)
- [273-v0.7.20 Gate F F2 Config Env Validation Guide](273-v0.7.20-gate-f-f2-config-env-validation-guide-v0.md)
- [274-v0.7.20 Gate F F2 Release Package Authority Guide](274-v0.7.20-gate-f-f2-release-package-authority-guide-v0.md)
- [275-v0.7.20 Gate F F2 Failure Refusal Hygiene Guide](275-v0.7.20-gate-f-f2-failure-refusal-hygiene-guide-v0.md)
- [276-v0.7.20 Gate F F2 Closure](276-v0.7.20-gate-f-f2-closure-v0.md)
- [277-v0.7.21 Gate F F3 Compatibility Rollback Observability](277-v0.7.21-gate-f-f3-compat-rollback-observability-v0.md)
- [278-v0.7.21 Gate F F3 Closure](278-v0.7.21-gate-f-f3-closure-v0.md)
- [279-v0.7.22 Gate F F4 Code Health Source Hygiene](279-v0.7.22-gate-f-f4-code-health-source-hygiene-v0.md)
- [280-v0.7.22 Gate F F4 Closure](280-v0.7.22-gate-f-f4-closure-v0.md)
- [281-v0.7.23 Gate F F5 Basic Local Trust Posture](281-v0.7.23-gate-f-f5-basic-local-trust-posture-v0.md)
- [282-v0.7.23 Gate F F5 Closure](282-v0.7.23-gate-f-f5-closure-v0.md)
- [283-v0.7.24 Gate F F6 Package Runtime Candidate Validation](283-v0.7.24-gate-f-f6-package-runtime-candidate-v0.md)
- [284-v0.7.24 Gate F F6 Closure](284-v0.7.24-gate-f-f6-closure-v0.md)
- [285-v0.7.25 Gate F F7 Release Candidate Closure Matrix](285-v0.7.25-gate-f-f7-release-candidate-closure-v0.md)
- [286-v0.7.25 Gate F F7 Manual Release Handoff](286-v0.7.25-gate-f-f7-manual-release-handoff-v0.md)
- [287-v0.7.25 Gate F F7 Closure](287-v0.7.25-gate-f-f7-closure-v0.md)
- [288-v0.8.1 Gate A Post-v0.8.0 Reset Workflow Doctrine](288-v0.8.1-gate-a-post-v0.8.0-reset-workflow-doctrine-v0.md)
- [289-v0.8.1 Gate A Closure](289-v0.8.1-gate-a-closure-v0.md)
- [290-v0.8.2 Gate B Adversarial Harness Contract](290-v0.8.2-gate-b-adversarial-harness-contract-v0.md)
- [291-v0.8.2 Gate B Seed Case Registry](291-v0.8.2-gate-b-seed-case-registry-v0.md)
- [292-v0.8.3 Gate C Local Trust Candidate Contract](292-v0.8.3-gate-c-local-trust-candidate-contract-v0.md)
- [293-v0.8.3 Gate C Local Trust Candidate Model](293-v0.8.3-gate-c-local-trust-candidate-model-v0.md)
- [294-v0.8.4 Gate D State Recovery Vault-Backup Contract](294-v0.8.4-gate-d-state-recovery-vault-backup-contract-v0.md)
- [295-v0.8.4 Gate D State Recovery Vault-Backup Model](295-v0.8.4-gate-d-state-recovery-vault-backup-model-v0.md)
- [296-v0.8.5 Gate E Relay Onboarding Adversarial Contract](296-v0.8.5-gate-e-relay-onboarding-adversarial-contract-v0.md)
- [297-v0.8.5 Gate E Relay Onboarding Case Matrix](297-v0.8.5-gate-e-relay-onboarding-case-matrix-v0.md)
- [298-v0.8.6 Gate F Native Deployment Process Contract](298-v0.8.6-gate-f-native-deployment-process-contract-v0.md)
- [299-v0.8.6 Gate F Service systemd Helper Threat Model](299-v0.8.6-gate-f-service-systemd-helper-threat-model-v0.md)
- [300-v0.8.6 Gate F Native Process Case Matrix](300-v0.8.6-gate-f-native-process-case-matrix-v0.md)
- [301-v0.8.6 Gate F Subgate Plan](301-v0.8.6-gate-f-subgate-plan-v0.md)
- [302-v0.8.6 Gate F F2 Component Runtime Process Probes Contract](302-v0.8.6-gate-f-f2-component-runtime-process-probes-contract-v0.md)
- [303-v0.8.6 Gate F F2 Component Probe Report Model](303-v0.8.6-gate-f-f2-component-probe-report-model-v0.md)
