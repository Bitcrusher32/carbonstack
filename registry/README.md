# CarbonStack Command Registry

Status: provisional v0
Owner repo: carbonstack
Boundary checkpoint: v0.5.68B public-surface refresh

This directory tracks known command, script, runner-profile, sidecar, API, and operator-facing surfaces across the CarbonStack repo family.

Current registry:

    registry/commands.v0.yaml


Current boundary table:

    registry/COMMAND_BOUNDARY_TABLE.v0.md

The command boundary table is the current registry-facing release-boundary artifact for command/profile/script/API classification. It is not an end-user manual and does not promote every listed command into public UX.

## Purpose

The registry exists to:

    keep command sprawl visible;
    avoid losing older operator/helper surfaces;
    distinguish public/dev/internal/legacy/future commands;
    preserve nonclaims beside commands;
    make future local help/manual text easier to generate;
    keep front-door READMEs from becoming command encyclopedias.

## Boundary

This registry is not a security claim.

This registry is not production certification.

This registry is not local-backbone.

This registry does not mean every command is user-facing.

This registry does not supersede release runbooks.

This registry does not itself create a public CLI/manual surface.

Registry presence is not command promotion.

Gitea remains source of truth.

## Current status after v0.5.68A

The registry is the current internal/current navigation and claim-boundary artifact for command surfaces.

The command boundary table renders the current registry-facing release-boundary classification:

    registry/COMMAND_BOUNDARY_TABLE.v0.md

The table classifies public/release-facing, source developer, dev-only OpenMLS, hidden/private validation, legacy/stub-era, internal sidecar/API, and future/unsupported surfaces.

Do not treat registry presence as public promotion.

As of v0.5.66, `relay-openmls-join-dev` has an explicit v0.6.0 release-surface decision: it remains manual/dev-only, excluded from `full`, excluded from `release-snapshot`, and is only a future full-profile candidate after repeated clean runs, package-root rehearsal, artifact behavior review, and an explicit later inclusion decision.

## Command-surface classes

The registry may include:

    public release validation surfaces;
    source developer validation surfaces;
    hidden/private validation surfaces;
    Comms dev command surfaces;
    legacy/stub-era surfaces;
    sidecar internal surfaces;
    Cypher API/operator surfaces;
    future placeholders.

These classes do not all have the same public meaning.

A command being listed in the registry only means:

    it exists or is intentionally tracked;
    it has a claim boundary;
    it should not be forgotten.

It does not mean:

    it is public CLI UX;
    it belongs in top-level README quick starts;
    it belongs in release runbooks;
    it belongs in full;
    it belongs in release-snapshot;
    it is production secure.

## Front-door README rule

Only commands marked `include_in_front_readme: true` should be considered for top-level/front-door docs.

`include_in_front_readme: true` is a candidate flag, not an automatic promotion rule.

`include_in_front_readme: false` means keep the command out of top-level public quick-start docs by default.

Dev/internal/legacy command surfaces should usually live in registry docs, component READMEs, implementation-specific docs, or maintainer-only notes.



## Flag metadata rule

Registry entries may include `required_flags`, `optional_flags`, and `environment` metadata when a command/profile/script has important operator-facing inputs.

The command boundary table renders that metadata where present. Missing flag metadata means the registry still needs enrichment; it does not mean the command has no flags.

A future command/helper lookup may expose per-command or per-flag help from this registry, but no lookup helper exists yet.


## Generated command reference rule

A generated command reference does not exist yet.

Before generating one, CarbonStack needs:

    stable registry schema;
    resolved audience/maturity vocabulary;
    explicit filtering rules;
    explicit nonclaim rendering;
    release-facing vs dev-only separation;
    legacy/stub-era warnings;
    future/unsupported entry suppression;
    hidden/private profile suppression or a separate maintainer-only section.

A generated command reference must not become a claim surface without preserving nonclaims.

## relay-openmls-join-dev rule

`relay-openmls-join-dev` remains:

    positive-path;
    live-umbrella/dev-oriented;
    hidden from the front README;
    outside full;
    outside release-snapshot;
    not local-backbone;
    not production secure messaging.

As of v0.5.66, the v0.6.0 decision is conservative exclusion: it remains manual/dev-only, excluded from full and release-snapshot, and is only a future full-profile candidate after repeated clean runs, package-root rehearsal, artifact behavior review, and an explicit later inclusion decision.

## Maintenance rule

When adding a new CLI command, runner profile, smoke script, sidecar command, release helper, or operator/API surface, update the registry in the same or next checkpoint.

## Validation status

The provisional v0 registry is validated by:

    carbonstack/tools/carbonstack-validate/command_registry_test.go

The check runs with:

    cd carbonstack/tools/carbonstack-validate
    go test ./... -count=1

The test validates registry structure, source_path existence, runner profile coverage, Comms command coverage, Comms script coverage, sidecar command IDs, Cypher API IDs, legacy send/inbox/ack classification, direct-vs-wrapper OpenMLS profile separation, and local-backbone/Gitea source-of-truth boundaries.

The check is intentionally a Go test for now, not a runner profile. It may become a runner profile later only if the registry policy stabilizes enough to justify public validation-profile status.

## Runtime send/inbox convergence rule

As of v0.5.65, the current canonical dev/pre-alpha OpenMLS application-message runtime paths are:

- `openmls-send-dev`
- `openmls-inbox-dev`

The older `send`, `inbox`, and `ack` commands remain legacy/stub-era surfaces. They are preserved for continuity and must not be described as mature OpenMLS-backed messaging UX.

`message-send-dev` and `message-inbox-dev` are now dev/pre-alpha opinionated wrapper commands after the v0.6.5 implementation checkpoint; they materially improve public-testing IO while preserving implementation truth. For v0.6.0, the release posture is documented separation:

- legacy send/inbox/ack remain warning-only;
- OpenMLS dev runtime commands remain explicit dev proof surfaces;
- mature product messaging UX remains future work.

## Dev/operator lookup

As of v0.6.7, the validation runner can inspect registry entries directly:

    go run . --profile registry-lookup --registry-id comms.message-send-dev
    go run . --profile registry-lookup --command "go run ./cmd/comms message-send-dev"

Lookup output is classification only. Registry presence is not promotion, not a generated public manual, not production UX, and not security certification.


Registry lookup enrichment examples:

    go run . --profile registry-lookup --list
    go run . --profile registry-lookup --list --lifecycle-status recommended_dev_wrapper
    go run . --profile registry-lookup --list --audience dev --maturity dev_only
    go run . --profile registry-lookup --list --missing-nonclaims

These remain dev/operator classification helpers. They are not a generated public manual, production UX, or security certification.

## v0.6.9 Relay onboarding lifecycle-status vocabulary

Relay onboarding commands remain dev/pre-alpha artifact/onboarding surfaces, not ordinary application-message UX.

Lifecycle statuses added in v0.6.9:

- `relay_onboarding_artifact_transport`
  - `comms.openmls-relay-keypackage-submit-dev`
  - `comms.openmls-relay-keypackage-inbox-dev`
  - `comms.openmls-relay-welcome-submit-dev`
  - `comms.openmls-relay-welcome-inbox-dev`
- `relay_onboarding_artifact_bridge`
  - `comms.openmls-relay-add-member-dev`
- `relay_onboarding_join_transition_candidate`
  - `comms.openmls-relay-join-dev`

These values are classification aids for operators and future generated docs. They do not promote Relay onboarding commands to production UX, hostile-server safety, identity verification, local-backbone readiness, `full`, or release-snapshot scope.

## v0.6.10 integrated runtime policy note

The planned integrated runtime profile name is `integrated-runtime-dev`.

It is now registered as `runner.integrated-runtime-dev` after the Stage 2 implementation patch.

When implemented, it should classify a live-umbrella-only dev profile that chains Relay onboarding plus `message-send-dev` / `message-inbox-dev --ack` while preserving individual primitive command surfaces and strict nonclaims.

Do not treat this planned profile as part of `full`, `release-snapshot`, production UX, local-backbone, hostile-server safety, or public release validation.

## v0.6.11 command reference generation policy

`registry/commands.v0.yaml` is the intended source for a future generated Markdown command reference, but the reference is intentionally deferred until registry metadata is hardened.

Current policy:

- do not generate `COMMAND_REFERENCE.v0.md` yet;
- do not generate Unix man pages yet;
- keep `COMMAND_BOUNDARY_TABLE.v0.md` hand-maintained for now;
- harden missing nonclaims, validation surfaces, flag metadata, lifecycle statuses, and internal/API/script boundary wording before generation;
- preserve the distinction between `full`, `release-snapshot`, and `integrated-runtime-dev`.

See `docs/207-v0.6.11-command-reference-generation-policy-v0.md`.

## v0.6.12 registry metadata hardening

v0.6.12 hardens registry metadata for future command-reference generation.

The goal is to reduce missing nonclaims and validation-surface opacity before generating a Markdown command reference. This does not create `COMMAND_REFERENCE.v0.md`, does not create man pages, and does not replace `COMMAND_BOUNDARY_TABLE.v0.md`.

The registry remains a classification and claim-boundary artifact. Registry presence is not command promotion.

## v0.6.13 generated command reference

`registry/COMMAND_REFERENCE.v0.md` is the generated dev/operator command reference.

Source:

    registry/commands.v0.yaml

Renderer:

    tools/registry/render-command-reference.py

Refresh/check:

    python3 tools/registry/render-command-reference.py
    python3 tools/registry/render-command-reference.py --check

Boundary:

- registry presence is classification, not promotion;
- the generated reference is dev/operator-facing, not general-public UX documentation;
- it is not a man-page set;
- it does not replace `COMMAND_BOUNDARY_TABLE.v0.md`;
- it does not change `full`, `release-snapshot`, or `integrated-runtime-dev`.

## v0.6.14 generated-reference freshness guard

After v0.6.14, the generated command reference is checked by the runner test suite:

    cd tools/carbonstack-validate
    go test ./...

The test runs the renderer in `--check` mode and fails when `registry/COMMAND_REFERENCE.v0.md` is stale relative to `registry/commands.v0.yaml`.

Boundary:

- `registry/COMMAND_REFERENCE.v0.md` remains generated dev/operator documentation;
- `registry/COMMAND_BOUNDARY_TABLE.v0.md` remains hand-maintained;
- no Unix man pages are generated by this guard.
## Relay Space-scoped normal-message boundary

The recommended dev/pre-alpha ordinary-message wrappers now require explicit Relay Space context:

    go run ./cmd/comms message-send-dev --relay-space <relay-space-id> ...
    go run ./cmd/comms message-inbox-dev --relay-space <relay-space-id> ... --ack

The sender and recipient must have active, non-disabled device routing membership in the selected Relay Space before scoped submission. Membership is routing authority only; it is not identity verification, trust acceptance, MLS authorization, or onboarding.

`message-inbox-dev` fetches the whole scoped device inbox so unsupported content/protocol pairs remain visible for explicit no-ack classification. `--ack` is scoped and is attempted only after successful OpenMLS message-open.

Current authority: `docs/236-v0.7.2-gate-b-b3-relay-scoped-normal-message-path-v0.md`.
