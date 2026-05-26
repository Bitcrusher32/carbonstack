# OpenMLS Sidecar Go Test Suite Split Result v0

Status: Go contract test maintainability split result
Component: CarbonStackComms / OpenMLS sidecar
Phase: v0.2.45 maintainability promotion
Previous docs:
- docs/90-openmls-sidecar-test-suite-split-plan-v0.md
- docs/93-openmls-sidecar-promotion-scaffold-result-v0.md
- docs/94-openmls-sidecar-rust-module-split-result-v0.md

## 1. Summary

This checkpoint records the split of the OpenMLS sidecar Go contract tests.

The promoted OpenMLS sidecar remains at:

    carbonstack-comms/internal/protocol/mls/openmls-sidecar

The known-good Phase 2D research reference remains at:

    carbonstack-comms/internal/protocol/mls/research/openmls-sidecar

The goal of this rung was maintainability only. No sidecar command schema, OpenMLS behavior, JSON envelope behavior, path-hint semantics, event names, or exit-code behavior was intentionally changed.

## 2. What changed

The previous large sidecar contract test file was split into focused files:

    internal/protocol/openmls_sidecar_helpers_test.go
    internal/protocol/openmls_sidecar_provider_info_test.go
    internal/protocol/openmls_sidecar_identity_test.go
    internal/protocol/openmls_sidecar_public_bundle_test.go
    internal/protocol/openmls_sidecar_conversation_test.go
    internal/protocol/openmls_sidecar_message_test.go
    internal/protocol/openmls_sidecar_message_negative_test.go

The package remains:

    package protocol

Shared helpers remain available package-wide.

## 3. Helper split

`openmls_sidecar_helpers_test.go` now owns shared test infrastructure, including:

- sidecar runner helpers;
- JSON envelope parser;
- file assertion helpers;
- generated-state cleanup helpers;
- two-member Alice/Bob setup helper;
- message protect/open helper wrappers;
- shared message success assertions.

## 4. Provider-info split

`openmls_sidecar_provider_info_test.go` now owns:

- provider-info command test;
- unsupported-command envelope test.

## 5. Identity split

`openmls_sidecar_identity_test.go` now owns identity tests only, including:

- missing label;
- invalid label;
- identity create state write;
- overwrite refusal;
- missing identity status;
- existing identity status.

## 6. Public-bundle split

`openmls_sidecar_public_bundle_test.go` now owns public-bundle tests, including:

- missing identity behavior;
- summary export;
- KeyPackage artifact export.

## 7. Conversation split

`openmls_sidecar_conversation_test.go` now owns conversation lifecycle tests, including:

- conversation-create;
- conversation-load-check;
- conversation-add-member Welcome export;
- conversation-join Welcome consume.

## 8. Message split

`openmls_sidecar_message_test.go` now owns positive/success message tests, including:

- one-way message protect/open;
- two sequential messages;
- bidirectional Alice/Bob message flow;
- out-of-order same-sender two-message delivery.

`openmls_sidecar_message_negative_test.go` now owns negative message tests, including:

- wrong-device rejection;
- wrong-conversation rejection;
- duplicate/replay rejection;
- corrupt/truncated artifact rejection.

## 9. Behavior preservation

This rung was behavior-preserving.

Expected validation:

    go test -p 1 ./internal/protocol
    go test -p 1 ./...
    powershell -ExecutionPolicy Bypass -File .\scripts\check-no-rust-artifacts.ps1

The tests continue to target the promoted sidecar path:

    internal/protocol/mls/openmls-sidecar

## 10. Blunders / continuity notes

During this split, the first provider-info extraction attempt was run before the expected identity target file existed. PowerShell continued after a failed `Get-Content`, creating an incomplete provider-info file and an empty/null identity file candidate.

The issue was caught before commit/push. The damaged file was restored from Git, the accidental identity file was removed, and the split was redone with hard guards that stop if expected files do not exist or if extracted content is empty.

This is a useful continuity note: future PowerShell split scripts should fail hard before writing when expected source content is missing.

## 11. Allowed claims

Allowed:

- the OpenMLS sidecar Go contract test suite is split by ownership;
- test coverage is intended to be unchanged;
- shared helpers are centralized;
- tests continue targeting the promoted sidecar;
- the research sidecar remains intact.

Not allowed:

- README/stale-warning cleanup is complete;
- Cypher relay recon has started;
- Comms runtime uses OpenMLS;
- production E2EE exists;
- secure production storage exists;
- all Rust command-family modules are fully split.

## 12. Next rung

Next planned rung:

    v0.2.46 — README/current-state/known-good command/stale warning cleanup.

Goal:

    update stale sidecar README, current-state docs, known-good commands, and stale warning text so future Cypher relay recon starts from accurate project state.

After that:

    v0.2.47 — Cypher minimal opaque MLS artifact relay recon.
