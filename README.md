# CarbonStack

CarbonStack is an experimental secure-communications appliance-stack project.

At this stage, it is not a finished messenger. It is not production-certified. It is not externally audited. It is not Android-ready.

The current public testing release remains:

    v0.3.20 runner-backed testing release

The current mainline validation state has moved beyond v0.3.20. Mainline v0.3.32 now includes a WSL Debian Go-runner validation surface with:

    local-cypher
    doctor
    core --clean-generated

`local-cypher` validates a Cypher-only local lifecycle. It builds a temporary Cypher binary, uses a temporary SQLite DB, binds to loopback, claims invites, registers devices, submits/retrieves/acks an opaque OpenMLS envelope, verifies restart/persistence behavior, checks one negative protocol pairing, and leaves no source-tree artifacts.

This proves a local experimental validation path. It does not prove production security.

_Related repositories: [carbonstack-comms](https://git.bitcrusher32.win/bitcrusher32/carbonstack-comms) / [carbonstack-cypher](https://git.bitcrusher32.win/bitcrusher32/carbonstack-cypher) / [carbonstack-os](https://git.bitcrusher32.win/bitcrusher32/carbonstack-os)_

## What is currently proven

The current validated backbone path includes:

1. OpenMLS KeyPackage relay through Cypher.
2. OpenMLS Welcome relay through Cypher.
3. OpenMLS application-message relay through Cypher.
4. Consume-then-ack behavior.
5. Payload metadata checks before downloaded artifact bytes are written locally.
6. Cypher schema migration idempotence through `schema_migrations`.
7. Explicit local Cypher operator config/data convention.
8. `local-cypher` runner-owned positive lifecycle validation.
9. `local-cypher` restart against the same temporary DB.
10. `local-cypher` rejection of the historical invalid stub-text/OpenMLS protocol pairing.
11. `core --clean-generated` validation with explicit generated-artifact cleanup.

This is the current experimental validation surface.

## What is not proven

CarbonStack does not currently prove:

- production E2EE product readiness;
- hostile-server safety;
- metadata privacy;
- secure local vault/storage;
- Android appliance readiness;
- polished Comms runtime send/inbox integration;
- rollback/replay safety against a malicious server;
- public ingress safety;
- systemd/cloudflared deployment readiness;
- external audit or certification;
- a stable public protocol.

Do not treat this repository as a finished secure messenger.

## Repositories

CarbonStack is split across component repositories.

- `carbonstack`: doctrine, docs, release framing, runbooks, validation runner, and public front-door material.
- `carbonstack-comms`: text-first Comms client, OpenMLS sidecar, relay helper, tests, and smoke harness.
- `carbonstack-cypher`: relay/storage server, envelope API, SQLite schema, migrations, and server tests.
- `carbonstack-os`: future constrained appliance OS concept; not part of the current runnable validation surface.

The public release surface starts here, in `carbonstack`.

Component repositories carry implementation details and development tests.

## Current testing path

For public testing or development, start from the latest CarbonStack release:

    https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases

Use the attached release assets and version-specific runbook for that release.

Do not use Gitea's default Source Code ZIP/TAR.GZ archives as the multi-repo test package unless a release explicitly says otherwise.

Current preferred public testing release:

    v0.3.20 runner-backed testing release

v0.3.4 remains available as the older Windows 11 / PowerShell testbed.

Current mainline dev/test validation target:

    Debian / WSL Debian, linux/amd64

Current mainline validation commands:

    cd tools/carbonstack-validate
    go run . --profile local-cypher
    go run . --profile doctor
    go run . --profile core --clean-generated

Windows 11 validation is retained as historical/secondary validation for v0.3.20, but after v0.3.20 CarbonStack mainline public dev/test releases are expected to continue Debian / WSL Debian first. Later Windows/BSD/Linux-family port work may be reconsidered after the server/backbone stack is more mature.

## v0.4.0 release direction

The next minor epoch release should be framed as:

    CarbonStack v0.4.0 broad local deployability pre-release

Intended meaning:

    milestone / research-and-development release
    WSL Debian runner-backed local validation surface
    local-cypher/core validation as the concrete artifact
    transition point before runtime Comms OpenMLS UX work

Not intended meaning:

    public-user-ready messenger
    application-use-ready messenger
    production secure stack
    hostile-server-certified system
    mature Comms release
    Android release
    CarbonStackOS release

## CarbonStack Relay Space terminology

Use:

    CarbonStack Relay Space

for the future addressable server/conversation space concept.

Avoid importing IRC culture/moderation assumptions. "IRC-like" may be used only as a historical analogy when explaining the earlier mental model.

## Docs archive

The docs folder is a chronological archive and release documentation surface.

Start here:

- `docs/README.md`

Older numbered docs may be stale. They preserve design history, failures, pivots, and implementation context. Current release/runbook docs define current behavior for a release.

## Core doctrine

Every feature is guilty until it proves it does not add unacceptable parser, network, sensor, identity, or filesystem authority.

CarbonStack prioritizes restricted surfaces, explicit trust changes, hostile-server assumptions, disposable parsers, and minimal ambient authority.

---

License: MIT.
See the repository's LICENSE file for more information.
