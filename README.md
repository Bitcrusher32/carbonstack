# CarbonStack

CarbonStack is an experimental secure-communications appliance-stack project.

It is not a finished messenger. It is not production-certified. It is not externally audited. It is not Android-ready. Do not use CarbonStack for operationally sensitive communications unless you have read the relevant release notes, runbooks, and limitations, and you understand the current release's boundaries.

The canonical project home is the self-hosted Gitea repository:

    https://git.bitcrusher32.win/bitcrusher32/carbonstack

GitHub mirrors may exist for discoverability and redundancy, but Gitea remains the source of truth for releases, tags, attached release assets, and current project state.

## Start here

For runnable or known-good artifacts, start from the Releases page:

    https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases

Use the attached release assets and the release-specific testing runbook for the release you are testing.

Do not treat Gitea's default Source Code ZIP/TAR.GZ downloads as the intended multi-repo validation package unless a release explicitly says so. The default archives are generated from the carbonstack repository alone. CarbonStack release packages may include multiple repositories plus release metadata.

## What this repository is

The carbonstack repository is the public front door for the CarbonStack project.

It contains:

- project doctrine;
- public release framing;
- roadmap material;
- validation and testing docs;
- historical design and implementation records;
- the CarbonStack validation runner;
- release/package validation support.

Implementation work is split across component repositories:

- carbonstack: doctrine, docs, release framing, runbooks, validation runner, and public front-door material.
- carbonstack-comms: text-first Comms client, OpenMLS sidecar, relay helper, tests, and smoke harness.
- carbonstack-cypher: relay/storage server, envelope API, SQLite schema, migrations, and server tests.
- carbonstack-os: future constrained appliance OS concept; not part of the current runnable validation package unless a release explicitly says otherwise.

Related repositories:

    https://git.bitcrusher32.win/bitcrusher32/carbonstack-comms
    https://git.bitcrusher32.win/bitcrusher32/carbonstack-cypher
    https://git.bitcrusher32.win/bitcrusher32/carbonstack-os

## Current maturity

CarbonStack is pre-alpha / experimental.

It currently focuses on validating pieces of a secure-communications backbone, especially the relationship between:

- CarbonStackComms;
- CarbonStackCypher;
- OpenMLS sidecar artifacts;
- local validation runners;
- release-package verification.

The exact claims, tested platforms, validation commands, package shape, and known-good artifacts are release-specific. Check the latest release notes and attached testing runbook before testing.

## What has been demonstrated so far

Across the v0.3.x and v0.4.x line, CarbonStack has demonstrated experimental validation coverage for:

- OpenMLS KeyPackage relay through Cypher;
- OpenMLS Welcome relay through Cypher;
- OpenMLS application-message relay through Cypher;
- consume-then-ack behavior;
- payload metadata checks before downloaded artifact bytes are written locally;
- Cypher schema migration idempotence through schema_migrations;
- local Cypher operator config/data conventions;
- Cypher-only local lifecycle validation through local-cypher;
- restart/persistence checks against a temporary DB;
- rejection of a historical invalid stub-text/OpenMLS protocol pairing;
- release package layout checks;
- package checksum verification;
- fresh-extraction validation;
- explicit cleanup of known OpenMLS sidecar generated roots.

This is useful project evidence. It is not a production security proof.

## What is not proven

CarbonStack does not currently prove:

- production readiness;
- production E2EE product readiness;
- hostile-server safety;
- metadata privacy;
- secure local vault/storage;
- Android appliance readiness;
- CarbonStackOS readiness;
- polished Comms runtime send/inbox UX;
- rollback/replay safety against a malicious server;
- public ingress safety;
- systemd/cloudflared deployment readiness;
- real homelab deployment safety;
- external audit or certification;
- a stable public protocol.

Do not treat this repository as a finished secure messenger.

## Testing and validation

Release-specific instructions live with each release.

The general pattern for current release packages is:

    download the attached multi-repo release package
    download the attached checksums and runbook
    extract to a fresh package root
    follow the release-specific testing runbook

Current runner profiles may include:

    verify-checksums
    release-snapshot
    full
    local-cypher
    core
    doctor

The meaning of these profiles can change as the project matures. Use the testing runbook attached to the release you are validating.

For development from a live checkout, see:

    tools/carbonstack-validate/README.md
    docs/README.md
    roadmap/ROADMAP.md

## CarbonStack Relay Space terminology

Use:

    CarbonStack Relay Space

for the future addressable server/conversation space concept.

Avoid importing IRC moderation/culture assumptions into CarbonStack. "IRC-like" may be used only as a historical analogy when explaining earlier design thinking.

## Docs archive

The docs folder is a chronological archive and release documentation surface.

Start here:

- docs/README.md

Older numbered docs may be stale. They preserve design history, failures, pivots, and implementation context. Current release pages, release runbooks, and newer docs define current behavior for a release.

## Core doctrine

Every feature is guilty until it proves it does not add unacceptable parser, network, sensor, identity, or filesystem authority.

CarbonStack prioritizes restricted surfaces, explicit trust changes, hostile-server assumptions, disposable parsers, and minimal ambient authority.

---

License: MIT.
See the repository's LICENSE file for more information.
