# CarbonStack v0.4.12 OpenMLS Bootstrap Command Contract

Status: command-contract checkpoint
Scope: v0.4.x runtime Comms OpenMLS integration
Primary environment: WSL Debian
Generated: 2026-06-05 03:56:48 -0400

## 1. Purpose

This document defines the dev-only command contract for future CarbonStackComms wrappers around OpenMLS sidecar bootstrap commands.

v0.4.11 concluded that wrappers are a reasonable next planning target, but not an immediate blind implementation target. This v0.4.12 contract defines exact command names, flags, output policy, sidecar JSON handling, path handling, test seams, implementation order, and nonclaims.

This document does not implement wrapper commands.

## 2. Current repo heads before this contract checkpoint

    carbonstack        8ee3ebb docs: plan dev OpenMLS bootstrap wrappers
    carbonstack-comms  0a48ae1 refactor: split OpenMLS runtime command helpers
    carbonstack-cypher 9ab994c docs: point to local operator runbook
    carbonstack-os     1bbbe52 docs: clarify CarbonStackOS target direction

## 3. Contract inputs

The v0.4.12 recon inspected:

    docs/161-openmls-bootstrap-wrapper-recon-v0.md
    roadmap v0.4.11/v0.4.12 area
    OpenMLS sidecar README
    OpenMLS sidecar src/main.rs command handlers
    OpenMLS sidecar JSON output fields
    OpenMLS sidecar src/paths.rs
    CarbonStackComms command registry
    carbonstack-comms/internal/app/openmls_runtime.go
    openmls-send-dev and openmls-inbox-dev tests
    smoke script bootstrap sequence
    protocol lifecycle tests
    state/trust non-mutation boundary

## 4. Current sidecar command set

The promoted OpenMLS sidecar supports:

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

Existing CarbonStackComms dev runtime wrappers already cover:

    message-protect through openmls-send-dev
    message-open through openmls-inbox-dev

This contract covers future wrappers for:

    identity-create
    identity-status
    public-bundle-export
    conversation-create
    conversation-load-check
    conversation-add-member
    conversation-join

## 5. Command names

Future commands should be named:

    openmls-identity-create-dev
    openmls-identity-status-dev
    openmls-bundle-export-dev
    openmls-conversation-create-dev
    openmls-conversation-load-check-dev
    openmls-conversation-add-member-dev
    openmls-conversation-join-dev

Naming rules:

    use openmls-*-dev
    keep dev explicit
    do not use local-backbone
    do not use Relay Space naming
    do not use secure/vault/production naming
    do not imply mature onboarding

## 6. Shared flags

All wrappers should support:

    --sidecar-dir <path>

Default:

    internal/protocol/mls/openmls-sidecar

Device-scoped wrappers should require:

    --sidecar-device-label <label>

Conversation-scoped wrappers should require:

    --conversation <label>

Path-consuming wrappers should require explicit artifact path flags.

Do not auto-derive sidecar labels from Comms state yet.

Rationale:

    sidecar state is dev-local and explicit-label based
    Comms state/trust/vault semantics are not mature enough
    automatic derivation could create false confidence about identity binding

## 7. Command-specific flag contract

### openmls-identity-create-dev

Required:

    --sidecar-device-label <label>

Optional:

    --sidecar-dir <path>

Sidecar call:

    cargo run --quiet -- identity-create --device-label <label>

Success output should include:

    command: openmls-identity-create-dev
    status: created
    sidecar_device_label: <label>
    sidecar_command: identity-create
    warning: dev/pre-alpha OpenMLS bootstrap path; not production identity UX

### openmls-identity-status-dev

Required:

    --sidecar-device-label <label>

Optional:

    --sidecar-dir <path>

Sidecar call:

    cargo run --quiet -- identity-status --device-label <label>

Success output should include:

    command: openmls-identity-status-dev
    status: loaded
    sidecar_device_label: <label>
    sidecar_command: identity-status
    identity_exists: true
    warning: dev/pre-alpha OpenMLS bootstrap path; not production identity UX

### openmls-bundle-export-dev

Required:

    --sidecar-device-label <label>

Optional:

    --sidecar-dir <path>
    --write-artifact

Sidecar call:

    cargo run --quiet -- public-bundle-export --device-label <label> [--write-artifact]

Success output should include:

    command: openmls-bundle-export-dev
    status: exported
    sidecar_device_label: <label>
    sidecar_command: public-bundle-export
    key_package_artifact_path_hint: <hint when present>
    key_package_artifact_path: <absolute path when hint is present>
    warning: dev/pre-alpha OpenMLS bootstrap path; not production identity UX

### openmls-conversation-create-dev

Required:

    --sidecar-device-label <label>
    --conversation <label>

Optional:

    --sidecar-dir <path>

Sidecar call:

    cargo run --quiet -- conversation-create --device-label <label> --conversation-label <label>

Success output should include:

    command: openmls-conversation-create-dev
    status: created
    sidecar_device_label: <label>
    sidecar_conversation_label: <label>
    sidecar_command: conversation-create
    conversation_state_path_hint: <hint when present>
    conversation_summary_path_hint: <hint when present>
    provider_storage_path_hint: <hint when present>
    warning: dev/pre-alpha OpenMLS bootstrap path; not production conversation UX

### openmls-conversation-load-check-dev

Required:

    --sidecar-device-label <label>
    --conversation <label>

Optional:

    --sidecar-dir <path>

Sidecar call:

    cargo run --quiet -- conversation-load-check --device-label <label> --conversation-label <label>

Success output should include:

    command: openmls-conversation-load-check-dev
    status: loaded
    sidecar_device_label: <label>
    sidecar_conversation_label: <label>
    sidecar_command: conversation-load-check
    group_reloadable: true
    warning: dev/pre-alpha OpenMLS bootstrap path; not production conversation UX

### openmls-conversation-add-member-dev

Required:

    --sidecar-device-label <label>
    --conversation <label>
    --member-keypackage <path>

Optional:

    --sidecar-dir <path>

Sidecar call:

    cargo run --quiet -- conversation-add-member --device-label <label> --conversation-label <label> --member-keypackage <path>

Success output should include:

    command: openmls-conversation-add-member-dev
    status: welcome_created
    sidecar_device_label: <label>
    sidecar_conversation_label: <label>
    sidecar_command: conversation-add-member
    member_keypackage_path_hint: <path/hint>
    welcome_artifact_path_hint: <hint when present>
    welcome_artifact_path: <absolute path when hint is present>
    warning: dev/pre-alpha OpenMLS bootstrap path; not production membership UX

### openmls-conversation-join-dev

Required:

    --sidecar-device-label <label>
    --conversation <label>
    --welcome <path>

Optional:

    --sidecar-dir <path>

Sidecar call:

    cargo run --quiet -- conversation-join --device-label <label> --conversation-label <label> --welcome <path>

Success output should include:

    command: openmls-conversation-join-dev
    status: joined
    sidecar_device_label: <label>
    sidecar_conversation_label: <label>
    sidecar_command: conversation-join
    welcome_artifact_path_hint: <path/hint>
    warning: dev/pre-alpha OpenMLS bootstrap path; not production membership UX

## 8. Sidecar JSON parsing policy

Wrapper commands should parse sidecar JSON.

They should not simply pass through raw JSON as their primary output.

Reasons:

    existing Comms dev commands use stable key/value output
    smoke scripts can parse key/value output more easily
    wrapper output should be consistent across commands
    sidecar JSON remains an internal provider envelope, not the user-facing Comms CLI contract

However, wrappers should preserve enough sidecar data to debug:

    sidecar_command
    path hints
    key status fields
    sidecar error code/message on failure

If parsing fails:

    return an error
    include the sidecar command name
    do not print success
    do not mutate Comms state/trust

## 9. Path handling policy

Sidecar path hints should be preserved exactly when present.

If a path hint points to a sidecar-generated artifact and is relative, wrappers should also print an absolute path resolved against:

    --sidecar-dir

Example:

    key_package_artifact_path_hint: .carbonstack-openmls-sidecar-state/dev/devices/bob/public-bundle.keypackage.bin
    key_package_artifact_path: internal/protocol/mls/openmls-sidecar/.carbonstack-openmls-sidecar-state/dev/devices/bob/public-bundle.keypackage.bin

Path policy:

    keep *_path_hint for sidecar-relative or sidecar-provided values
    add *_path for resolved absolute or caller-useful paths when applicable
    do not rewrite input artifact paths unless needed for sidecar invocation
    do not print private material
    do not claim vault safety

## 10. Error policy

Missing required flags should return normal Comms CLI errors before invoking sidecar.

Sidecar ok=false envelopes should become Go errors.

The error should include:

    sidecar command
    sidecar error code if present
    sidecar message if present

Wrappers should not print success after sidecar failure.

Suggested shape:

    OpenMLS sidecar <command> failed: <code>: <message>

If sidecar exits nonzero and output cannot be parsed:

    run OpenMLS sidecar <command>: <exec error>

## 11. State/trust policy

Bootstrap wrappers must not mutate:

    Comms state.json
    trust.json
    trust-events.jsonl

Bootstrap wrappers may mutate sidecar dev-local state because that is the purpose of sidecar bootstrap commands.

Bootstrap wrappers should not require Comms ready state unless a later command genuinely needs Cypher/Comms identity.

Current wrapper contract:

    no state.RequireReadyDevice
    no trust.EvaluateSend
    no trust mutation
    no Cypher calls

This keeps wrappers explicitly sidecar bootstrap/dev utilities, not mature identity UX.

## 12. Testing policy

Future app unit tests should not run Rust sidecar commands directly.

Tests should use injected seams, following the existing style used by openmls-send-dev and openmls-inbox-dev.

Preferred seam:

    var runOpenMLSBootstrapSidecarForCommand = runOpenMLSBootstrapSidecar

Suggested helper signature:

    func runOpenMLSBootstrapSidecar(sidecarDir string, sidecarCommand string, args ...string) (openMLSSidecarBootstrapEnvelope, error)

Tests should cover:

    missing required flags
    sidecar command arguments
    sidecar failure prevents success
    path-hint normalization
    dev warning presence where output is captured
    no Comms state/trust mutation
    no old send/inbox replacement

Real Rust sidecar execution should remain in:

    protocol tests
    smoke script
    dev-runtime-openmls profile

## 13. Implementation file boundary

Future wrapper code should live in:

    carbonstack-comms/internal/app/openmls_bootstrap.go

Future wrapper tests should live in:

    carbonstack-comms/internal/app/openmls_bootstrap_dev_test.go

Keep:

    command registration and usage in internal/app/commands.go
    message runtime helpers in internal/app/openmls_runtime.go
    bootstrap wrapper helpers in internal/app/openmls_bootstrap.go

## 14. Implementation order

Preferred staged implementation:

    v0.4.13:
        openmls-identity-create-dev
        openmls-identity-status-dev

    v0.4.14:
        openmls-bundle-export-dev
        openmls-conversation-create-dev
        openmls-conversation-load-check-dev

    v0.4.15:
        openmls-conversation-add-member-dev
        openmls-conversation-join-dev

    v0.4.16:
        smoke-script migration recon
        decide whether direct sidecar bootstrap should be replaced with wrapper calls

Do not migrate the smoke script before wrapper tests and direct smoke proof both pass.

## 15. Relationship to smoke script

Current smoke proof remains authoritative:

    carbonstack-comms/scripts/dev-openmls-runtime-smoke.sh

It should remain direct-sidecar-bootstrap based until wrappers exist and pass tests.

Later migration options:

    keep direct-sidecar smoke and add wrapper-based smoke variant
    or replace direct bootstrap with wrappers only after wrapper coverage is strong

Do not weaken the existing smoke proof.

## 16. Relationship to local-backbone

This contract does not justify local-backbone naming.

local-backbone remains reserved because:

    wrappers are not implemented yet
    wrappers will be dev-only when implemented
    old send/inbox remain stub-era
    Relay Space join/onboarding is not implemented
    state/trust/vault semantics remain immature
    hostile-server harnesses remain future work
    release-package runtime validation remains unsupported

## 17. Nonclaims preserved

This contract does not claim:

    wrapper implementation exists
    local-backbone
    Relay Space join/onboarding
    mature send/inbox UX
    old send/inbox are OpenMLS-backed
    dev-runtime-openmls belongs in full
    release-package runtime validation
    production readiness
    production E2EE
    hostile-server safety
    metadata privacy
    secure vault/storage
    PQ/hybrid security
    Android readiness
    CarbonStackOS readiness
    external audit or certification

## 18. Suggested next rung

Recommended next rung:

    v0.4.13 implement identity create/status dev wrappers

Scope:

    carbonstack-comms only
    add internal/app/openmls_bootstrap.go
    add internal/app/openmls_bootstrap_dev_test.go
    register openmls-identity-create-dev
    register openmls-identity-status-dev
    update carbonstack-comms README only if implementation lands
    validate app tests, Comms tests, dev-runtime-openmls, local-cypher, doctor, and core

Do not implement all bootstrap wrappers at once unless the first patch proves the shape is trivial and safe.

## 19. Summary

v0.4.12 defines a strict dev-only OpenMLS bootstrap wrapper command contract.

The contract keeps sidecar labels explicit, normalizes sidecar JSON to stable Comms key/value output, preserves path hints, avoids Comms trust/state mutation, and stages implementation.

The next safest patch is the smallest identity-wrapper subset.
