# CarbonStack docs archive

This folder is a historical docs archive and continuity trail for CarbonStack.

It contains planning notes, recon notes, implementation result notes, process checkpoints, release-supporting docs, and other project continuity documents written over time.

## How to read this folder

The numbered docs are chronological.

Lower-numbered docs may be stale.

That is intentional.

Older docs preserve the reasoning, constraints, mistakes, pivots, implementation context, and decision history that led to later work. They should not be treated as the current release surface unless a newer release document explicitly points to them.

## Current-vs-historical rule

Use the newest relevant release docs, current-state docs, README files, and runbooks for current behavior.

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

## Release-specific docs

Release-specific docs will be added as the project matures.

A release-specific doc or runbook should say what is current for that release, what is proven, what is not proven, which component repos are involved, and which commands are known-good.

For example, a future v0.3.0 release surface may include docs such as:

    v0.3.0 experimental backbone release notes;
    v0.3.0 known-good runbook;
    v0.3.0 security status and nonclaims;
    v0.3.0 component map.

Those release docs should be treated as the public-facing source of truth for that release.

## Why this archive is preserved

CarbonStack is being developed through staged research, implementation, validation, and cleanup rungs.

Destroying or constantly rewriting the archive would erase important process context.

Instead, the project preserves old docs as historical artifacts and adds newer docs to clarify current state.

This means the docs folder is not a single polished manual.

It is a continuity archive plus release documentation surface.

## Security and maturity warning

Many docs describe experimental, pre-release, or dev-scaffold behavior.

CarbonStack is not production-certified.

CarbonStack has not received senior external audit or security certification.

Historical docs must not be used to imply production readiness, hostile-server safety, Android readiness, metadata privacy, or complete E2EE product status unless a current release document explicitly says so.
## Current v0.3.0 packaging freeze

The v0.3.0 release packaging freeze is recorded at:

    docs/119-v0.3.0-release-packaging-freeze-v0.md

Use it with:

    docs/v0.3.0-minor-epoch-release.md
## Post-v0.3.0 verification and governance

Current post-release verification and governance docs:

    docs/120-v0.3.0-post-release-verification-v0.md
    docs/121-security-claim-validation-policy-v0.md
    docs/122-logdoc-case-study-sanitization-plan-v0.md
## Sanitized project LogDoc list

A sanitized project LogDoc archive is available at:

    sanitized-project-logdoc-list/

This folder contains sanitized LogDoc material for workflow/case-study use. It is separate from the main numbered docs archive and should not be treated as the current release source of truth.
## Clean release snapshot self-test recon

The v0.3.3 clean release snapshot self-test recon is recorded at:

    docs/123-v0.3.0-clean-snapshot-self-test-recon-v0.md

This recon checks whether the v0.3.0 release source snapshots can run validation from extraction rather than from the active working repos.
