# CarbonStack Command Registry

Status: provisional v0
Owner repo: carbonstack

This directory tracks known command, script, runner-profile, sidecar, API, and operator-facing surfaces across the CarbonStack repo family.

Current registry:

    registry/commands.v0.yaml

Purpose:

    keep command sprawl visible
    avoid losing older operator/helper surfaces
    distinguish public/dev/internal/legacy/future commands
    preserve nonclaims beside commands
    make future local help/manual text easier to generate
    keep front-door READMEs from becoming command encyclopedias

Boundary:

    This registry is not a security claim.
    This registry is not production certification.
    This registry is not local-backbone.
    This registry does not mean every command is user-facing.
    This registry does not supersede release runbooks.
    Gitea remains source of truth.

Maintenance rule:

    When adding a new CLI command, runner profile, smoke script, sidecar command, release helper, or operator/API surface, update the registry in the same or next checkpoint.

Front-door README rule:

    Only commands marked `include_in_front_readme: true` should be considered for top-level/front-door docs.
    Dev/internal/legacy command surfaces should usually live in registry docs, component READMEs, or implementation-specific docs.

Validation status:

    The provisional v0 registry is validated by:

        carbonstack/tools/carbonstack-validate/command_registry_test.go

    The check runs with:

        cd carbonstack/tools/carbonstack-validate
        go test ./... -count=1

    The test validates registry structure, source_path existence, runner profile coverage, Comms command coverage, Comms script coverage, sidecar command IDs, Cypher API IDs, legacy send/inbox/ack classification, direct-vs-wrapper OpenMLS profile separation, and local-backbone/Gitea source-of-truth boundaries.

    The check is intentionally a Go test for now, not a runner profile. It may become a runner profile later only if the registry policy stabilizes enough to justify public validation-profile status.
