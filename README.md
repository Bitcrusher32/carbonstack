# CarbonStack

CarbonStack is an experimental secure-communications appliance-stack project.

At this stage, it is not a finished messenger. It is not production-certified. It is not externally audited. It is not Android-ready.

The current validated artifact is the v0.3.x experimental backbone: a local Cypher + Comms OpenMLS relay proof. CarbonStackComms generates OpenMLS sidecar artifacts. CarbonStackCypher relays them as opaque envelopes. The receiving side writes and consumes the artifacts through the OpenMLS sidecar.

This proves a local experimental relay lifecycle. It does not prove production security.

_Related repositories: [carbonstack-comms](https://git.bitcrusher32.win/bitcrusher32/carbonstack-comms) / [carbonstack-cypher](https://git.bitcrusher32.win/bitcrusher32/carbonstack-cypher) / [carbonstack-os](https://git.bitcrusher32.win/bitcrusher32/carbonstack-os)_

## What is currently proven

The current local proof validates this path:

1. Bob creates an OpenMLS KeyPackage artifact.
2. Cypher relays the KeyPackage to Alice.
3. Alice consumes the KeyPackage and creates a Welcome artifact.
4. Cypher relays the Welcome to Bob.
5. Bob consumes the Welcome and joins the conversation.
6. Alice creates an OpenMLS application-message artifact.
7. Cypher relays the application-message to Bob.
8. Bob consumes the application-message through the OpenMLS sidecar.
9. The plaintext matches.
10. Envelopes are acknowledged only after sidecar consume succeeds.
11. Payload metadata is checked before downloaded artifact bytes are written locally.

This is the current experimental backbone.

## What is not proven

CarbonStack does not currently prove:

- production E2EE product readiness;
- hostile-server safety;
- metadata privacy;
- secure local vault/storage;
- Android appliance readiness;
- polished Comms runtime send/inbox integration;
- rollback/replay safety against a malicious server;
- external audit or certification;
- a stable public protocol.

Do not treat this repository as a finished secure messenger.

## Repositories

CarbonStack is split across component repositories.

- `carbonstack`: doctrine, docs, release framing, runbooks, and public front-door material.
- `carbonstack-comms`: text-first Comms client, OpenMLS sidecar, relay helper, tests, and smoke harness.
- `carbonstack-cypher`: relay/storage server, envelope API, SQLite schema, migrations, and server tests.
- `carbonstack-os`: future constrained appliance OS concept; not part of the current runnable relay proof.

The public release surface starts here, in `carbonstack`.

Component repositories carry implementation details and development tests.

## Known-good local proof

The current runbook is:

- `docs/113-experimental-backbone-deployability-runbook-v0.md`

That document explains the current local validation path, component roles, smoke harness, payload metadata, ack semantics, generated-state warnings, and security nonclaims.

## Docs archive

The docs folder is a chronological archive and release documentation surface.

Start here:

- `docs/README.md`

Older numbered docs may be stale. They preserve design history, failures, pivots, and implementation context. Current release/runbook docs define current behavior for a release.

## Core doctrine

Every feature is guilty until it proves it does not add unacceptable parser, network, sensor, identity, or filesystem authority.

CarbonStack prioritizes restricted surfaces, explicit trust changes, hostile-server assumptions, disposable parsers, and minimal ambient authority.

## Current direction

The near-term target is v0.3.x backbone maturation: verify the released v0.3.0 artifacts, improve self-test/runbook UX, reduce Windows-only assumptions, clarify low-level deploy configuration, and harden the backbone before runtime Comms UX work.

The current v0.3.x release surface should be framed as:

- CarbonStack experimental backbone;
- concrete validated artifact: Cypher + Comms OpenMLS relay backbone;
- pre-alpha / experimental;
- not certified secure;
- not externally audited;
- not a finished messenger product.
## v0.3.0 release surface

For the v0.3.0 experimental backbone release README, start here:

    docs/v0.3.0-minor-epoch-release.md

That document is the consolidated release entrypoint for the current Cypher + Comms OpenMLS relay backbone.

---

License: MIT.
See the repository's LICENSE file for more information.



