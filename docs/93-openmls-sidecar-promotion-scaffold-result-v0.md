# OpenMLS Sidecar Promotion Scaffold Result v0

Status: Behavior-preserving promotion scaffold
Component: CarbonStackComms / OpenMLS sidecar
Phase: v0.2.43 maintainability promotion
Previous docs:
- docs/87-openmls-sidecar-current-state-index-v0.md
- docs/88-openmls-sidecar-maintainability-promotion-plan-v0.md
- docs/89-openmls-sidecar-module-split-plan-v0.md
- docs/90-openmls-sidecar-test-suite-split-plan-v0.md
- docs/91-openmls-sidecar-artifact-ownership-map-v0.md
- docs/92-openmls-sidecar-command-schema-matrix-v0.md

## 1. Summary

This checkpoint creates the promoted OpenMLS sidecar scaffold while preserving the Phase 2D research sidecar as a known-good reference.

Research reference remains at:

    carbonstack-comms/internal/protocol/mls/research/openmls-sidecar

Promoted maintained scaffold now exists at:

    carbonstack-comms/internal/protocol/mls/openmls-sidecar

The promoted sidecar is behavior-identical at this checkpoint. No command schema, OpenMLS lifecycle behavior, artifact shape, or event semantics were intentionally changed.

## 2. Why this changed

Phase 2D mainline research closed with a validated OpenMLS sidecar, but the sidecar still lived under `mls/research`.

Before splitting large Rust files or Go tests, the project needed a promoted implementation scaffold above the research path.

This avoids destructively refactoring the known-good research proof.

## 3. What changed

In carbonstack-comms:

- copied the known-good research sidecar into `internal/protocol/mls/openmls-sidecar`;
- left `internal/protocol/mls/research/openmls-sidecar` intact;
- pointed Go sidecar contract tests at the promoted sidecar path;
- validated the promoted sidecar with the Phase 2D contract suite.

## 4. What did not change

No behavior change was intended.

The command surface remains:

    provider-info
    identity-create
    identity-status
    public-bundle-export
    conversation-create
    conversation-load-check
    conversation-add-member
    conversation-join
    message-protect
    message-open

Unsupported commands remain:

    state-checkpoint
    state-load-check

The sidecar remains dev-local and not production secure.

## 5. Validation target

Validation for this checkpoint:

    cargo check
    cargo test
    cargo run -- provider-info
    go test -p 1 ./internal/protocol
    go test -p 1 ./...
    scripts/check-no-rust-artifacts.ps1

The important claim is equivalence:

    the promoted sidecar passes the same Phase 2D closure tests previously used against the research sidecar.

## 6. Allowed claims

Allowed:

- a promoted maintained OpenMLS sidecar scaffold exists outside the research path;
- the research sidecar remains as a known-good Phase 2D reference;
- Go sidecar tests now target the promoted scaffold;
- behavior is intended to be equivalent to the Phase 2D closure sidecar.

Not allowed:

- productionization is complete;
- the Rust modules have been split;
- the Go tests have been split;
- README/current-state cleanup is complete;
- Cypher routes MLS artifacts;
- Comms runtime uses OpenMLS;
- secure production storage exists.

## 7. Next rung

Next planned rung:

    v0.2.44 — split promoted Rust sidecar modules.

Goal:

    reduce main.rs and state.rs by separating CLI, envelopes, paths, identity, public bundle, conversation, and message logic while preserving behavior.
