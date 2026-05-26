# OpenMLS Sidecar Module Split Plan v0

Status: Rust sidecar modularization plan
Component: CarbonStackComms / OpenMLS sidecar
Phase: v0.2.42 planning for v0.2.44 implementation
Previous docs:
- docs/86-openmls-sidecar-phase2d-mainline-closure-result-v0.md
- docs/88-openmls-sidecar-maintainability-promotion-plan-v0.md

## 1. Purpose

This document plans the Rust sidecar module split after the sidecar is promoted out of the research path.

The current research sidecar has validated behavior, but its file shape is too large for the next phase.

Current large files:

    src/main.rs
      about 2,718 lines

    src/state.rs
      about 2,261 lines

The promoted implementation should keep behavior but separate responsibilities.

## 2. Current responsibilities

main.rs currently mixes:

- command dispatch;
- argument parsing;
- provider-info JSON;
- success envelope printing;
- failure envelope printing;
- command constants;
- warning strings;
- unsupported command handling;
- Rust-side tests/help text.

state.rs currently mixes:

- identity operations;
- public bundle operations;
- conversation operations;
- message operations;
- path helpers;
- result structs;
- validation helpers;
- file IO;
- OpenMLS adapter logic.

This is understandable from a research sprint but should not be carried into the maintained implementation.

## 3. Proposed promoted module layout

Recommended promoted sidecar path:

    internal/protocol/mls/openmls-sidecar

Recommended src layout:

    src/main.rs
    src/cli.rs
    src/envelope.rs
    src/errors.rs
    src/labels.rs
    src/paths.rs
    src/provider.rs
    src/identity.rs
    src/public_bundle.rs
    src/conversation.rs
    src/message.rs
    src/schema.rs

## 4. Module ownership

main.rs:

- process entrypoint;
- command dispatch;
- no large JSON literals;
- no OpenMLS operation logic;
- no path construction.

cli.rs:

- parse device label;
- parse conversation label;
- parse member KeyPackage path;
- parse Welcome path;
- parse message path;
- parse plaintext;
- parse message label;
- future argument shape validation.

envelope.rs:

- provider-info response;
- success response serialization;
- error response serialization;
- warning list;
- capability list;
- unsupported command list;
- common envelope fields.

errors.rs:

- error codes;
- exit code mapping;
- provider event mapping;
- trust relevance flags;
- reusable error constructors if useful.

labels.rs:

- validate_device_label;
- validate_conversation_label;
- validate_message_label;
- reserved name checks;
- filesystem-safe label rules.

paths.rs:

- device root paths;
- identity paths;
- public bundle paths;
- device-scoped conversation paths;
- Welcome paths;
- message artifact paths;
- opened-message summary paths.

provider.rs:

- CarbonStackSidecarProvider;
- OpenMlsProvider implementation;
- MemoryStorage save/load.

identity.rs:

- create_dev_identity;
- load_dev_identity_status;
- identity result structs if kept local.

public_bundle.rs:

- export_dev_public_bundle_summary;
- public bundle result structs;
- KeyPackage artifact/manifest logic.

conversation.rs:

- create_dev_conversation;
- load_dev_conversation_status;
- add_dev_conversation_member;
- join_dev_conversation;
- conversation result structs.

message.rs:

- protect_dev_message;
- open_dev_message;
- plaintext validation;
- message artifact validation;
- message result structs.

schema.rs:

- shared result structs if keeping them out of command modules is cleaner.

## 5. Dead helper cleanup

The promoted implementation should not carry old global creator conversation helpers unless they are still required.

Likely deletion candidates after v0.2.41:

- conversation_state_dir;
- conversation_summary_path;
- conversation_provider_storage_path;
- conversation_welcome_artifact_path;
- conversation_welcome_manifest_path;
- conversation_add_member_summary_path;
- conversation_messages_dir;
- conversation_message_dir;
- conversation_message_artifact_path;
- conversation_message_manifest_path;
- conversation_message_protect_summary_path.

The research sidecar preserves historical context. The promoted sidecar should express only the current device-scoped model.

## 6. Stale warning cleanup

The promoted sidecar should not keep old warnings that imply implemented commands are unimplemented.

Keep warnings like:

- OpenMLS is not wired into CarbonStackComms;
- Cypher does not route MLS payloads;
- trust-state storage does not consume provider events;
- dev-only signer/provider storage is not a secure vault.

Remove/update warnings that say current implemented commands are not implemented.

## 7. Split order

Recommended order for v0.2.44:

1. Create paths.rs and move path helpers.
2. Create schema.rs or local result structs per module.
3. Move identity operations.
4. Move public bundle operations.
5. Move conversation operations.
6. Move message operations.
7. Move CLI parsing.
8. Move envelope printing.
9. Reduce main.rs to dispatcher.
10. Remove dead old global helpers.
11. Run full Phase 2D contract suite.

## 8. Validation target

After the module split:

- cargo check passes for promoted sidecar;
- cargo test passes for promoted sidecar;
- Go Phase 2D closure tests pass against promoted sidecar;
- command output JSON remains equivalent;
- no research sidecar behavior is altered.

## 9. Non-goals

Do not use the module split to:

- add Cypher routing;
- alter command schemas;
- introduce generated message IDs;
- change provider storage format;
- change OpenMLS lifecycle semantics;
- productionize storage;
- add trust-state mutation.
