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
