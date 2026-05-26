# OpenMLS Sidecar Rust Module Split Result v0

Status: Rust maintainability split result
Component: CarbonStackComms / OpenMLS sidecar
Phase: v0.2.44 maintainability promotion
Previous docs:
- docs/87-openmls-sidecar-current-state-index-v0.md
- docs/88-openmls-sidecar-maintainability-promotion-plan-v0.md
- docs/89-openmls-sidecar-module-split-plan-v0.md
- docs/93-openmls-sidecar-promotion-scaffold-result-v0.md

## 1. Summary

This checkpoint records the first promoted OpenMLS sidecar Rust module split.

The promoted sidecar remains at:

    carbonstack-comms/internal/protocol/mls/openmls-sidecar

The Phase 2D research reference remains intact at:

    carbonstack-comms/internal/protocol/mls/research/openmls-sidecar

The goal of this rung was maintainability only. No command schema, JSON envelope semantics, MLS lifecycle behavior, path-hint meaning, or event/exit-code behavior was intentionally changed.

## 2. What changed

The promoted sidecar gained focused Rust modules:

    src/paths.rs
    src/schema.rs
    src/envelope.rs
    src/cli.rs

The existing files remain:

    src/main.rs
    src/state.rs
    src/provider.rs
    src/labels.rs

## 3. paths.rs split

`paths.rs` now owns current device-scoped path construction.

This includes:

- device state root paths;
- identity paths;
- public bundle paths;
- device-scoped conversation state paths;
- Welcome artifact and manifest paths;
- message artifact and manifest paths;
- opened-message summary paths.

Old global `dev/conversations/<conversation-label>/` helper residue was removed from the promoted sidecar.

The research sidecar remains available as the historical reference.

## 4. schema.rs split

`schema.rs` now owns public result structs used across command handlers and envelope rendering.

Moved result structs include:

- IdentityCreateResult;
- IdentityStatusResult;
- PublicBundleExportResult;
- ConversationCreateResult;
- ConversationLoadCheckResult;
- ConversationAddMemberResult;
- ConversationJoinResult;
- MessageProtectResult;
- MessageOpenResult.

`state.rs` continues to re-export these as needed so the split remained low-risk.

## 5. envelope.rs split

`envelope.rs` now owns:

- provider-info rendering;
- shared phase constants;
- supported command list;
- unsupported command list;
- shared warning strings;
- JSON escaping helper;
- identity command printers;
- public-bundle command printers;
- unsupported-command printer.

Some conversation and message printers remain in `main.rs` for now. That is intentional. They can be moved in a later polish rung after test split or command-family splits.

## 6. cli.rs split

`cli.rs` now owns simple CLI argument parsing helpers:

- parse_device_label;
- parse_conversation_label;
- parse_member_keypackage_path;
- parse_welcome_artifact_path;
- parse_plaintext;
- parse_message_artifact_path;
- parse_message_label;
- parse_write_artifact_flag;
- DEFAULT_MESSAGE_LABEL;
- parse_path_buf.

This keeps `main.rs` focused more on dispatch and command handling.

## 7. Behavior preservation

This rung was behavior-preserving.

Expected validation:

    cargo check
    cargo test
    cargo run -- provider-info
    go test -p 1 ./internal/protocol
    go test -p 1 ./...
    powershell -ExecutionPolicy Bypass -File .\scripts\check-no-rust-artifacts.ps1

The Go sidecar contract tests continue to target the promoted sidecar path.

## 8. Blunders / continuity notes

During the Rust split work, a few cleanup hazards occurred and were corrected:

- broad regex deletion during paths split temporarily removed the join section from `state.rs`;
- the missing join result/function were restored from the untouched research sidecar reference;
- leftover old global helper residue caused duplicate/missing path helper errors and was removed;
- provider envelope split initially duplicated phase constants and provider-info helpers between `main.rs` and `envelope.rs`;
- the final provider envelope split preserved the existing provider-info JSON shape rather than silently changing schema.

These are exactly why the research sidecar remains frozen as a known-good reference during promotion.

## 9. Allowed claims

Allowed:

- the promoted OpenMLS sidecar now has a cleaner Rust module structure;
- path helpers, result schemas, selected envelope printers, and CLI parsers have been split out;
- behavior is intended to remain equivalent to the v0.2.43 promoted scaffold;
- the research sidecar remains untouched as the Phase 2D reference.

Not allowed:

- full sidecar modularization is complete;
- all envelope printers have been moved;
- command-family logic has been split into identity/public_bundle/conversation/message modules;
- Go tests have been split;
- README/current-state cleanup is complete;
- Cypher routes MLS artifacts;
- production E2EE exists;
- Comms runtime uses OpenMLS.

## 10. Remaining maintainability ladder

Next rung:

    v0.2.45 — split Go sidecar contract tests.

Goal:

    preserve coverage while moving the giant Go sidecar test file into focused files with shared helpers.

Then:

    v0.2.46 — docs/readme cleanup.

Goal:

    update stale sidecar README, current-state docs, known-good commands, and stale warning text.

Then:

    v0.2.47 — Cypher relay recon begins.

Goal:

    design minimal opaque MLS artifact relay using the validated OpenMLS sidecar artifact map.
