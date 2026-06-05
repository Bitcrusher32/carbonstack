# CarbonStack v0.4.19 Command Registry Validation and Local Help Plan

Status: registry validation and local-help planning checkpoint
Scope: late v0.4.x command-surface hygiene before v0.5.x
Primary environment: WSL Debian
Generated: 2026-06-05 07:41:59 -0400

## 1. Purpose

v0.4.18 introduced the first provisional cross-repo command registry:

    registry/commands.v0.yaml

v0.4.19 validates that registry against live command surfaces and records a local help/manual plan.

This rung is intentionally boring:

    no command behavior changes
    no runner profile behavior changes
    no local-backbone naming
    no generated command reference yet
    no Comms embedded help metadata yet
    no v0.5.x state/trust/vault/PQ implementation yet

## 2. Current repo heads before this checkpoint

    carbonstack        a4162d7 docs: add known-good command registry
    carbonstack-comms  cb4e59d test: add wrapper-based OpenMLS runtime smoke proof
    carbonstack-cypher 9ab994c docs: point to local operator runbook
    carbonstack-os     1bbbe52 docs: clarify CarbonStackOS target direction

## 3. Recon result

The v0.4.19 recon found:

    registry_entries: 62
    registry_command_fields: 62
    duplicate_ids: none
    missing_required_markers: none

Field presence:

    short_help: 62
    why_exists: 62
    include_in_front_readme: 62
    maturity: 62
    audience: 62
    kind: 62
    source_path: 62

Runner coverage:

    registry_missing_runner_profiles: none
    registry_extra_runner_ids: none

Comms coverage:

    registry_missing_comms_commands: none
    registry_extra_comms_ids: none

Script coverage:

    all actual Comms script files were represented except scripts/README.md.
    scripts/README.md is documentation, not a command/script surface, so it is intentionally not registered.

## 4. Validation added

v0.4.19 adds a Go test:

    carbonstack/tools/carbonstack-validate/command_registry_test.go

This test validates, when run from a full umbrella/release-package layout:

    registry has entries
    registry IDs are unique
    required fields exist for each entry
    source_path values resolve to existing files/directories
    runner profile dispatch is covered
    Comms command dispatch is covered
    Comms scripts are covered, excluding scripts/README.md
    sidecar command IDs are present
    Cypher server/API IDs are present
    legacy send/inbox/ack remain marked legacy
    direct and wrapper OpenMLS runtime profiles remain distinct
    local-backbone and Gitea source-of-truth boundaries remain present
    include_in_front_readme count stays small enough to avoid README sprawl

The test skips only when a sibling umbrella layout is unavailable.

## 5. Why Go test, not runner profile yet

Registry validation should be accounted for in normal development and deploy/release thinking.

For now, the right first step is Go test coverage inside:

    carbonstack/tools/carbonstack-validate

This means the check runs with:

    go test ./... -count=1

It should not become a runner profile yet because:

    the registry is still v0
    the policy may evolve quickly
    runner profile names are public-facing validation surface
    a simple test is enough to catch registry drift for now

A future runner profile can be considered only after this check proves stable.

## 6. Local help/manual plan

Preferred ladder:

    B first: registry validation test/check
    C later: generated markdown command reference from registry
    D later: Comms embedded help/manual strings, only after command structure is ready
    E later: runner profile for registry validation, only if useful and stable

Avoid for now:

    generating local command help from the registry
    moving command text into commands.go prematurely
    replacing component READMEs with generated docs
    expanding the top README into a command encyclopedia

## 7. README vs registry vs generated command reference boundary

Front README:

    purpose, current release, source of truth, where to find releases, minimal validation entrypoint
    only commands marked include_in_front_readme: true should be candidates

Registry:

    complete command/profile/script/API inventory
    maturity, audience, why_exists, source_path, validation_surface, nonclaims, replacement status
    source of truth for command-surface hygiene

Generated command reference, later:

    human-readable view derived from registry
    grouped by repo/component/audience/maturity
    should not replace release runbooks
    should not become a claim surface without nonclaims preserved

Component READMEs:

    local component workflow
    important examples
    boundary warnings
    not necessarily complete cross-stack inventory

## 8. Downstream effect

v0.4.19 completes the grouped registry/help-planning goalset enough to prepare for a roadmap refresh after the grouped work lands.

The next roadmap refresh should acknowledge that v0.4.15 through v0.4.19 are complete:

    add-member/join wrappers
    wrapper-based smoke
    wrapper runtime profile
    provisional command registry
    registry validation/local-help plan

## 9. v0.5.0 minor epoch release carry-forward

After late v0.4.x finishes, CarbonStack should cut a v0.5.0 minor epoch release using the accumulated v0.4.x work.

That v0.5.0 release should have dedicated rungs for:

    cleanup
    deployability testing
    release-style replication
    release package rehearsal
    release hardening
    public wording / claim-boundary cleanup
    final LogDoc and breakpoint export
    release asset generation
    release notes using the established continuity style

This should happen before major new v0.5.x state/trust/vault/PQ implementation begins, unless a specific reason emerges to merge the release and implementation work.

## 10. Nonclaims

v0.4.19 does not prove:

    production readiness
    production E2EE
    hostile-server safety
    metadata privacy
    secure vault/storage
    PQ/hybrid security
    local-backbone
    Relay Space join/onboarding
    release-package runtime validation
    mature messaging UX

## 11. Suggested next rung

Recommended next rung after v0.4.19:

    refresh the long-term roadmap after the grouped registry/help-planning goalset

Then proceed toward late-v0.4.x cleanup and the eventual v0.5.0 minor epoch release ladder.
