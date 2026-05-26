# OpenMLS Sidecar README and Current-State Cleanup Result v0

Status: README/current-state cleanup result
Component: CarbonStack / CarbonStackComms OpenMLS sidecar
Phase: v0.2.46 maintainability cleanup
Previous docs:
- docs/87-openmls-sidecar-current-state-index-v0.md
- docs/88-openmls-sidecar-maintainability-promotion-plan-v0.md
- docs/89-openmls-sidecar-module-split-plan-v0.md
- docs/90-openmls-sidecar-test-suite-split-plan-v0.md
- docs/94-openmls-sidecar-rust-module-split-result-v0.md
- docs/95-openmls-sidecar-go-test-suite-split-result-v0.md

## 1. Summary

This checkpoint records the README/current-state cleanup after the OpenMLS sidecar promotion, Rust module split, and Go test split.

The active promoted sidecar remains:

    carbonstack-comms/internal/protocol/mls/openmls-sidecar

The frozen Phase 2D research reference remains:

    carbonstack-comms/internal/protocol/mls/research/openmls-sidecar

No sidecar behavior change was intended.

## 2. What changed in CarbonStackComms

The promoted sidecar README was updated to reflect the current state:

- supported commands are listed accurately;
- unsupported commands are limited to `state-checkpoint` and `state-load-check`;
- the sidecar is described as dev-local / experimental, not production E2EE;
- generated signer/provider state is explicitly marked as non-committable;
- the current device-scoped state layout is documented;
- the split Go test ownership map is documented;
- validation commands now point at the promoted sidecar.

The research sidecar README now begins with a frozen-reference notice pointing to the promoted sidecar.

Stale runtime/help text was corrected where it still implied implemented commands were unsupported.

## 3. What changed in CarbonStack docs

Current-state and planning docs were lightly updated:

- docs/87 now records the v0.2.43-v0.2.46 maintainability ladder state;
- docs/88 now marks the maintainability plan as landed through v0.2.45 and v0.2.46 active cleanup;
- docs/89 now records the v0.2.44 Rust module split result note;
- docs/90 now records the v0.2.45 Go test split result note;
- docs/96 records this cleanup checkpoint.

Historical docs were not rewritten wholesale.

## 4. Current active files

Promoted Rust sidecar modules:

    src/cli.rs
    src/envelope.rs
    src/labels.rs
    src/main.rs
    src/paths.rs
    src/provider.rs
    src/schema.rs
    src/state.rs

Go sidecar contract test split:

    openmls_sidecar_helpers_test.go
    openmls_sidecar_provider_info_test.go
    openmls_sidecar_identity_test.go
    openmls_sidecar_public_bundle_test.go
    openmls_sidecar_conversation_test.go
    openmls_sidecar_message_test.go
    openmls_sidecar_message_negative_test.go

## 5. Validation

Expected validation for this checkpoint:

    cargo check
    cargo test
    cargo run -- provider-info
    go test -p 1 ./internal/protocol
    go test -p 1 ./...
    powershell -ExecutionPolicy Bypass -File .\scripts\check-no-rust-artifacts.ps1
    powershell -ExecutionPolicy Bypass -File .\scripts\validate-local.ps1

## 6. Allowed claims

Allowed:

- the promoted sidecar README is current;
- the research README clearly identifies itself as a frozen reference;
- current-state docs now reflect the completed promotion/module/test split ladder;
- stale warnings that claimed implemented commands were unimplemented were corrected;
- v0.2.47 Cypher relay recon can start from a cleaner sidecar/test/doc baseline.

Not allowed:

- production E2EE exists;
- Cypher routes MLS artifacts;
- Comms runtime uses OpenMLS;
- production secure vault storage exists;
- Android/Pixel 4a validation exists;
- all optional Rust command-family splits are complete.

## 7. Next rung

Next planned rung:

    v0.2.47 — Cypher minimal opaque MLS artifact relay recon.

Goal:

    design minimal relay behavior for opaque MLS artifacts such as KeyPackage, Welcome, and application-message artifacts without exposing signer/provider storage or plaintext.
