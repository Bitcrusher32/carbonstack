# CarbonStack docs

This directory is a living documentation archive.

Lower-numbered documents preserve historical design, research, validation, and release-process state. They may contain stale assumptions because they intentionally record the project at the time they were written.

For current public testing and release status, start with:

    top-level README.md
    roadmap/ROADMAP.md
    docs/139-runner-backed-testing-release-cleanup-v0.md

Current public testing release:

    v0.3.20 runner-backed testing release

Current mainline validation state:

    v0.3.32 local-cypher negative protocol validation

Current mainline validation commands:

    cd tools/carbonstack-validate
    go run . --profile local-cypher
    go run . --profile doctor
    go run . --profile core --clean-generated

Current v0.4.0 direction:

    broad local deployability pre-release
    milestone / research-and-development release
    not public-user-ready
    not application-use-ready
    not production secure
    not hostile-server-certified
    not a mature Comms release
    not a CarbonStackOS release

## Historical archive note

Most numbered docs are historical records, not the current release front door.

Do not assume an older numbered document supersedes the latest release/runbook docs unless the newer docs explicitly say so.

## How to read this folder

The numbered docs are chronological.

Lower-numbered docs may be stale.

That is intentional.

Older docs preserve the reasoning, constraints, mistakes, pivots, implementation context, and decision history that led to later work. They should not be treated as the current release surface unless a newer release document explicitly points to them.

## Current-vs-historical rule

Use the newest relevant release docs, current-state docs, README files, and runbooks for current behavior.

Use older numbered docs for:

    project history
    design rationale
    implementation context
    continuity preservation
    debugging old decisions
    understanding why a later result exists

Do not assume an old plan document still describes current behavior.

Do not assume an old result document still describes the latest release surface.

Do not rewrite old historical docs just because implementation moved forward.

## Current release and validation docs

- `docs/139-runner-backed-testing-release-cleanup-v0.md` — v0.3.20 public runner-backed testing release cleanup and final validation plan.
- `docs/148-local-cypher-validation-contract-v0.md` — v0.3.29 local-cypher validation contract.
- `docs/149-local-cypher-runner-implementation-v0.md` — v0.3.30 local-cypher runner implementation.
- `docs/150-local-cypher-polish-generated-cleanup-v0.md` — v0.3.31 local-cypher polish and explicit generated-artifact cleanup.
- `docs/151-local-cypher-negative-protocol-validation-v0.md` — v0.3.32 local-cypher negative protocol validation.
- `docs/152-pre-v0.4.0-release-surface-cleanup-v0.md` — v0.3.33 pre-v0.4.0 release-surface cleanup.
- `docs/153-full-profile-release-validation-ladder-v0.md` — v0.3.34 full profile release validation ladder; `full` now runs `release-snapshot` then `local-cypher`.
- `docs/154-v0.4.0-package-rehearsal-v0.md` — v0.3.36 throwaway v0.4.0-style package rehearsal; fresh extraction checksum verification and full validation passed under WSL Debian.
- `docs/155-runtime-comms-openmls-cypher-recon-v0.md` — v0.4.1 runtime Comms/OpenMLS/Cypher recon; records that send/inbox are still stub-era while OpenMLS/Cypher relay is validated through lower-level seams.
- `docs/156-runtime-openmls-command-contract-v0.md` — v0.4.2 runtime OpenMLS command contract; defines `openmls-send-dev` and `openmls-inbox-dev` before implementation.
- `docs/157-dev-runtime-openmls-smoke-proof-v0.md` — records the v0.4.5 dev runtime OpenMLS CLI smoke proof using openmls-send-dev -> Cypher -> openmls-inbox-dev --ack.
- `docs/158-pre-local-backbone-assessment-v0.md` — v0.4.7 pre-local-backbone assessment; keeps local-backbone reserved, recommends future `dev-runtime-openmls` profile naming, and plans helper extraction before more runtime growth.
- `docs/159-dev-runtime-openmls-runner-profile-v0.md` — v0.4.9 manual `dev-runtime-openmls` validation profile; wraps the Comms smoke proof while preserving not-local-backbone and not-production boundaries.
- `docs/160-dev-runtime-openmls-profile-boundary-v0.md` — v0.4.10 boundary check for the manual `dev-runtime-openmls` profile; records repeatability, cleanup behavior, full-profile separation, and non-git package-like root refusal.
- `docs/161-openmls-bootstrap-wrapper-recon-v0.md` — v0.4.11 recon/planning for dev-only OpenMLS bootstrap wrappers; preserves openmls-*-dev naming, keeps local-backbone reserved, and recommends a command-contract doc before implementation.
- `docs/162-openmls-bootstrap-command-contract-v0.md` — v0.4.12 command contract for dev-only OpenMLS bootstrap wrappers; defines flags, output shape, path handling, sidecar JSON parsing, testing seams, and staged implementation order.

## Security and maturity warning

Many docs describe experimental, pre-release, or dev-scaffold behavior.

CarbonStack is not production-certified.

CarbonStack has not received senior external audit or security certification.

Historical docs must not be used to imply production readiness, hostile-server safety, Android readiness, metadata privacy, or complete E2EE product status unless a current release document explicitly says so.
