# OpenMLS Sidecar Maintainability Promotion Plan v0

Status: Maintainability plan / docs-only recon
Component: CarbonStackComms / OpenMLS sidecar
Phase: v0.2.42 maintainability promotion planning
Previous docs:
- docs/86-openmls-sidecar-phase2d-mainline-closure-result-v0.md
- docs/87-openmls-sidecar-current-state-index-v0.md

## 1. Purpose

This document defines the maintainability promotion rung after Phase 2D mainline closure.

The OpenMLS sidecar research path has done its job. It proved the dev-local OpenMLS lifecycle and the critical message behavior needed before Cypher relay research.

The next goal is to preserve that known-good research sidecar while creating a cleaner maintained implementation scaffold above the research path.

## 2. Core decision

Keep the research sidecar intact as a known-good reference.

Create a promoted maintained sidecar higher in the tree.

Current research path:

    carbonstack-comms/internal/protocol/mls/research/openmls-sidecar

Recommended promoted path:

    carbonstack-comms/internal/protocol/mls/openmls-sidecar

Alternative acceptable path:

    carbonstack-comms/internal/protocol/mls/sidecar/openmls

The first promoted scaffold should be behavior-identical.

## 3. Research vs promoted meaning

Research sidecar:

- known-good Phase 2D reference;
- preserves sprint history;
- should not be churned during maintainability refactors;
- can be used for comparison if behavior changes accidentally.

Promoted sidecar:

- maintained dev implementation scaffold;
- still not production secure;
- still not a user release;
- still not Android/Pixel 4a;
- still not wired into Comms runtime;
- but no longer treated as disposable research.

Promoted does not mean production.

## 4. Why promote now

Phase 2D proved:

- identity lifecycle;
- public KeyPackage export;
- conversation lifecycle;
- add-member/Welcome export;
- join/Welcome consume;
- message protect/open;
- explicit message labels;
- duplicate/replay/corrupt/wrong-target behavior;
- bidirectional Alice/Bob flow;
- device-scoped creator and joined state.

However, the current implementation is hard to maintain:

- main.rs is too large;
- state.rs is too large;
- Go sidecar contract tests are too large;
- sidecar README is stale;
- old helper residue remains;
- current behavior is spread across many historical docs.

A maintainability rung should clean the shape before Cypher relay research begins.

## 5. Required preservation rules

During promotion and modularization:

- preserve command names;
- preserve required arguments;
- preserve JSON envelope shape;
- preserve event names and phases unless explicitly versioned;
- preserve exit codes;
- preserve path-hint semantics;
- preserve dev-state security boundaries;
- preserve all Phase 2D tests;
- do not add Cypher routing;
- do not add Comms runtime integration;
- do not change OpenMLS protocol behavior intentionally.

## 6. Promotion equivalence target

The promoted sidecar must pass the same Phase 2D closure contract suite as the research sidecar.

Equivalent behavior means:

- same lifecycle success path;
- same wrong-device and wrong-conversation failure shape;
- same duplicate/replay failure shape;
- same corrupt artifact failure shape;
- same bidirectional message proof;
- same device-scoped artifact path model.

One-time equivalence proof is required before the promoted sidecar becomes the active test target.

## 7. Rung sequence

v0.2.42:

    docs-only promotion plan/recon.

v0.2.43:

    copy research sidecar to promoted location;
    update build/test harness only enough to target promoted path;
    no behavior changes.

v0.2.44:

    split promoted Rust modules;
    remove dead old global helper residue;
    clean stale warnings;
    keep command/schema behavior identical.

v0.2.45:

    split Go sidecar tests;
    centralize helpers;
    preserve coverage.

v0.2.46:

    update READMEs;
    update current-state docs;
    update known-good command examples;
    clearly mark historical docs as historical.

v0.2.47:

    begin Cypher minimal opaque MLS artifact relay recon.

## 8. Non-goals

Do not use this rung to:

- implement Cypher relay;
- move toward Android;
- create production storage;
- mutate trust-state;
- introduce generated message IDs;
- redesign MLS state;
- expand group membership matrix;
- add secure vault abstractions;
- rewrite historical docs wholesale.

## 9. Definition of done for v0.2.42

v0.2.42 is complete when the following docs exist:

- current state index;
- promotion plan;
- Rust sidecar module split plan;
- Go sidecar test split plan;
- artifact ownership map;
- command/schema matrix.

No code changes are required for v0.2.42.

## v0.2.46 status update

The planned maintainability ladder has landed through v0.2.45:

- v0.2.43 promoted the sidecar scaffold outside the research path;
- v0.2.44 split the promoted Rust sidecar modules;
- v0.2.45 split the Go sidecar contract tests;
- v0.2.46 updates README/current-state/stale-warning documentation.

The next planned rung is v0.2.47 Cypher minimal opaque MLS artifact relay recon.

