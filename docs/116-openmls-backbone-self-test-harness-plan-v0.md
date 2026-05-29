# OpenMLS Backbone Self-Test Harness Plan v0

Status: planning / recon result
Component: CarbonStack + CarbonStackComms
Phase: v0.2.65 OpenMLS backbone self-test harness planning

## 1. Purpose

This document defines the next public-facing validation surface for the CarbonStack experimental backbone.

The internal planning label "Option B" is not public release language.

Public-facing docs should call this surface:

    OpenMLS backbone self-test harness

or:

    CarbonStack OpenMLS backbone self-test

The goal is to expose the current known-good backbone proof through a clearer builder-facing command surface without pretending it is a finished messenger.

## 2. Current execution state

The current working execution path lives in `carbonstack-comms`.

Current script:

    scripts/smoke-openmls-real-cypher-relay.ps1

Current broader validation:

    scripts/smoke-openmls-real-cypher-relay.ps1 -Full

Current artifact guard:

    scripts/check-no-rust-artifacts.ps1

The current smoke harness already proves the local OpenMLS sidecar + real Cypher server relay lifecycle.

It starts a real local Cypher server, uses a temp SQLite database, relays KeyPackage / Welcome / application-message artifacts, validates payload metadata before artifact write, acknowledges envelopes only after successful sidecar consume, and verifies final plaintext recovery.

## 3. Current runtime CLI boundary

The current `cmd/comms` runtime CLI remains stub-era for `send`, `inbox`, and `ack`.

Those commands are useful historical/local scaffolding.

They are not the current OpenMLS backbone proof.

The OpenMLS relay path should not be wired into user-facing `send` / `inbox` commands during this planning rung.

## 4. Recommended next implementation

The next implementation should add a small wrapper script in `carbonstack-comms`:

    scripts/self-test-openmls-backbone.ps1

The wrapper should call the existing known-good smoke harness instead of duplicating the proof logic.

Expected behavior:

    default:
        run the targeted OpenMLS backbone self-test

    -Full:
        pass through to the broader validation path

    output:
        use public-facing self-test language
        keep maturity warnings near the top
        clearly state what is proven
        clearly state what is not proven

The existing lower-level script may remain as the implementation detail:

    scripts/smoke-openmls-real-cypher-relay.ps1

## 5. Non-goals

The self-test harness must not become:

- a production deployment script;
- a user-facing encrypted chat CLI;
- a top-level `carbonstack` orchestrator;
- an Android test path;
- a secure vault proof;
- a hostile-server-complete proof;
- a production E2EE claim;
- a metadata-privacy claim;
- an external audit or certification claim.

## 6. Public wording

Use:

    OpenMLS backbone self-test harness
    experimental local backbone lifecycle
    known-good local proof
    Cypher + Comms OpenMLS relay backbone

Avoid public release wording such as:

    Option B
    Option C
    secure messenger
    production E2EE
    hostile-server safe
    metadata-private
    trustless
    certified
    audit-ready

## 7. Why a wrapper script first

A wrapper script is the smallest honest public surface.

It avoids touching stub-era `send` / `inbox` runtime commands.

It avoids creating a fake product UX.

It keeps execution in `carbonstack-comms`.

It lets `carbonstack` document the release surface without owning execution.

It can be replaced or promoted later if the CLI/dev harness matures.

## 8. Expected v0.2.66 work

The next implementation rung should:

1. add `scripts/self-test-openmls-backbone.ps1`;
2. make it call the current smoke harness;
3. support `-Full`;
4. print direct public-facing status/nonclaim text;
5. update `carbonstack-comms/scripts/README.md`;
6. update the main `carbonstack` runbook and validation matrix;
7. validate the same known-good command set.

## 9. Local Go cache note

On Windows, antivirus software may flag generated Go test executables or interfere with Go temp output.

A narrow local workaround is to use repo-local Go temp/cache directories:

    .go-tmp/
    .go-cache/

These directories are local build/test artifacts.

They must not be committed.

If used, they should be ignored by Git.
