# CarbonStack v0.4.11 OpenMLS Bootstrap Wrapper Recon

Status: recon and planning checkpoint
Scope: v0.4.x runtime Comms OpenMLS integration
Primary environment: WSL Debian
Generated: 2026-06-05 03:33:17 -0400

## 1. Purpose

This document records the v0.4.11 recon and planning checkpoint for possible dev-only Comms wrappers around OpenMLS sidecar bootstrap commands.

v0.4.10 confirmed that the manual `dev-runtime-openmls` runner profile is repeatable, cleanup behavior is understood, non-git package-like roots are refused, and `full` remains separate.

The next remaining runtime-integration seam is the direct sidecar bootstrap setup currently embedded in:

    carbonstack-comms/scripts/dev-openmls-runtime-smoke.sh

This document does not implement wrappers.

It defines the planning boundary for future dev-only wrappers.

## 2. Current repo heads before this docs checkpoint

    carbonstack        7384f43 docs: record dev OpenMLS runtime profile boundary
    carbonstack-comms  0a48ae1 refactor: split OpenMLS runtime command helpers
    carbonstack-cypher 9ab994c docs: point to local operator runbook
    carbonstack-os     1bbbe52 docs: clarify CarbonStackOS target direction

## 3. Recon inputs

The v0.4.11 scout inspected:

    v0.4.10 profile-boundary document
    v0.4.x roadmap section
    v0.5.x state/trust/vault planning boundary
    carbonstack-comms/scripts/dev-openmls-runtime-smoke.sh
    smoke-script sidecar command sequence
    carbonstack-comms README dev runtime sections
    carbonstack-comms internal/app command registry
    carbonstack-comms internal/app/openmls_runtime.go
    openmls-send-dev and openmls-inbox-dev tests
    OpenMLS sidecar README command list and boundary
    OpenMLS sidecar Rust command definitions and flags
    existing protocol lifecycle tests as contract source
    Comms state/trust/label surfaces

## 4. Current direct sidecar bootstrap sequence

The committed smoke script currently uses direct sidecar calls for bootstrap setup:

    identity-create --device-label <alice-label>
    identity-create --device-label <bob-label>
    public-bundle-export --device-label <bob-label> --write-artifact
    conversation-create --device-label <alice-label> --conversation-label <conversation-label>
    conversation-add-member --device-label <alice-label> --conversation-label <conversation-label> --member-keypackage <path>
    conversation-join --device-label <bob-label> --conversation-label <conversation-label> --welcome <path>

After this setup, the actual runtime proof target remains:

    openmls-send-dev -> Cypher -> openmls-inbox-dev --ack

Correct interpretation:

    direct sidecar bootstrap is acceptable for the current dev smoke proof
    direct sidecar bootstrap is not mature Comms runtime UX
    direct sidecar bootstrap is the next candidate for dev-only wrapper planning
    replacing it is an end goal, not an immediate implementation requirement

## 5. Supported sidecar commands relevant to wrapper planning

The promoted sidecar currently supports:

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

Current Comms dev runtime commands already wrap:

    message-protect through openmls-send-dev
    message-open through openmls-inbox-dev

Candidate future bootstrap wrappers cover:

    identity-create
    identity-status
    public-bundle-export
    conversation-create
    conversation-load-check
    conversation-add-member
    conversation-join

## 6. Recommended wrapper naming

Use explicit dev-only names:

    openmls-identity-create-dev
    openmls-identity-status-dev
    openmls-bundle-export-dev
    openmls-conversation-create-dev
    openmls-conversation-load-check-dev
    openmls-conversation-add-member-dev
    openmls-conversation-join-dev

Do not use:

    local-backbone
    backbone
    join-space
    relay-space
    production
    secure
    vault
    setup

Rationale:

    openmls-*-dev matches existing openmls-send-dev and openmls-inbox-dev
    dev suffix keeps the maturity boundary loud
    names describe sidecar-adjacent development behavior
    names avoid importing future Relay Space or local-backbone assumptions

## 7. Recommended implementation boundary

Future wrapper code should probably live in:

    carbonstack-comms/internal/app/openmls_bootstrap.go

Keep:

    command registration and usage in commands.go
    message protect/open runtime glue in openmls_runtime.go
    bootstrap wrapper glue in openmls_bootstrap.go

Rationale:

    openmls_runtime.go already owns send/open application-message runtime behavior
    bootstrap commands are setup/onboarding-adjacent, not message runtime
    separating them reduces command-glue sprawl
    future tests can isolate bootstrap behavior from send/inbox behavior

## 8. Command contract recommendation

Future wrappers should:

    remain dev-only
    default to the existing sidecar directory
    require explicit --sidecar-device-label where a device label is needed
    require explicit --conversation where a conversation label is needed
    accept explicit input artifact paths for add-member/join
    print stable human-readable key/value output
    preserve raw sidecar path hints where useful
    convert sidecar-relative artifact hints to absolute paths when helpful
    surface sidecar errors clearly
    avoid mutating Comms trust state
    avoid pretending to manage secure vault/provider state
    avoid replacing send/inbox
    avoid Relay Space terminology except as a future nonclaim

For now, do not auto-derive sidecar labels from Comms device labels.

Reason:

    Comms state/trust/vault semantics are not mature enough
    sidecar state is dev-local and explicit-label based
    automatic derivation could create false confidence about identity binding

## 9. Output shape recommendation

Prefer key/value output similar to existing dev commands.

Examples of likely output fields:

    command: openmls-identity-create-dev
    status: created
    sidecar_device_label: <label>
    warning: dev/pre-alpha OpenMLS bootstrap path; not production identity UX

    command: openmls-bundle-export-dev
    status: exported
    sidecar_device_label: <label>
    key_package_artifact_path_hint: <hint>
    key_package_artifact_path: <absolute-path-if-available>

    command: openmls-conversation-add-member-dev
    status: welcome_created
    sidecar_device_label: <label>
    sidecar_conversation_label: <label>
    welcome_artifact_path_hint: <hint>
    welcome_artifact_path: <absolute-path-if-available>

Avoid claiming:

    verified identity
    trusted member
    secure group
    production-ready conversation
    Relay Space membership

## 10. Testing recommendation

Future wrapper tests should follow the existing app-test injection style.

Existing pattern:

    openmls-send-dev uses runOpenMLSMessageProtectForCommand
    openmls-send-dev uses submitOpenMLSArtifactEnvelopeForCommand
    openmls-inbox-dev uses runOpenMLSMessageOpenForCommand
    openmls-inbox-dev uses inbox/ack/write injection seams

Bootstrap wrappers should use a similar injected sidecar runner seam, for example:

    runOpenMLSBootstrapForCommand
    or one command-specific seam per wrapper if simpler

Tests should cover:

    missing required flags
    sidecar command failure does not print success
    expected sidecar command arguments are passed
    path-hint handling
    dev/pre-alpha warning presence
    no trust-state mutation
    no send/inbox replacement

Do not rely on full Rust sidecar execution in app unit tests.

Real sidecar execution should stay in protocol tests, smoke script, or runner profiles.

## 11. Minimal implementation sequence

Do not implement all wrappers blindly.

Recommended order:

    v0.4.12 command contract doc for dev-only OpenMLS bootstrap wrappers
    v0.4.13 first minimal wrapper patch:
        openmls-identity-create-dev
        openmls-identity-status-dev
    v0.4.14 bundle/conversation create wrappers:
        openmls-bundle-export-dev
        openmls-conversation-create-dev
        openmls-conversation-load-check-dev
    v0.4.15 member/welcome wrappers:
        openmls-conversation-add-member-dev
        openmls-conversation-join-dev
    v0.4.16 smoke-script migration recon:
        decide whether scripts/dev-openmls-runtime-smoke.sh should use the wrappers

Alternative if scope needs compression:

    v0.4.12 contract
    v0.4.13 implement all wrappers only if the contract is simple and tests remain clean

Preferred path:

    staged implementation

Reason:

    every wrapper is guilty until it proves it does not add unacceptable state, identity, path, or claim confusion

## 12. Why not immediately replace the smoke script

Do not replace direct sidecar bootstrap inside the smoke proof yet.

Reasons:

    wrappers do not exist yet
    wrapper output contracts are not defined yet
    wrapper tests do not exist yet
    direct bootstrap currently works and is validated
    smoke proof is currently the strongest runtime anchor
    breaking the smoke script would weaken v0.4.x confidence
    replacement should happen only after wrappers prove stable

Better sequence:

    define wrapper contract
    implement wrappers with app tests
    validate direct sidecar smoke still passes
    add a wrapper-based smoke variant or migrate the existing smoke script
    validate dev-runtime-openmls still passes
    then reassess local-backbone readiness

## 13. Relationship to local-backbone

This recon does not justify local-backbone naming.

local-backbone remains reserved because:

    bootstrap is still direct sidecar setup today
    wrapper commands do not exist yet
    wrappers would still be dev-only even after implementation
    old send/inbox remain stub-era
    Relay Space join/onboarding is not implemented
    state/trust/vault semantics remain immature
    hostile-server harnesses remain future work
    release-package runtime validation remains unsupported

Correct phrase:

    dev-only OpenMLS bootstrap wrapper planning

Incorrect phrase:

    local-backbone implementation

## 14. Relationship to v0.5.x

Keep v0.5.x deferred.

Wrapper planning touches identity/state/provider concepts, but it should not prematurely solve:

    secure vault
    production provider storage
    trust-state UX
    backup/recovery
    PQ/hybrid ciphersuite migration
    hostile-server safety

v0.5.x remains the correct place for state/trust/vault/PQ readiness.

v0.4.x should only make the current dev runtime path less manually bootstrapped and more explicit.

## 15. Nonclaims preserved

This planning checkpoint does not claim:

    production readiness
    production E2EE
    hostile-server safety
    metadata privacy
    secure vault/storage
    Android readiness
    CarbonStackOS readiness
    public ingress safety
    real deployment safety
    external audit or certification
    PQ/hybrid security
    local-backbone
    Relay Space join/onboarding
    mature Comms send/inbox UX
    old send/inbox are OpenMLS-backed
    dev-runtime-openmls belongs in full
    release-package runtime validation

## 16. Suggested next rung

Recommended next rung:

    v0.4.12 dev-only OpenMLS bootstrap command contract

Focus:

    define exact flags
    define exact output fields
    define sidecar JSON parsing policy
    define path-hint / absolute-path policy
    define testing seams
    define implementation order
    define smoke-script migration conditions

Do not implement wrapper commands until that contract exists.

## 17. Summary

v0.4.11 confirms that dev-only OpenMLS bootstrap wrappers are a reasonable next planning target, but not an immediate blind implementation target.

The current smoke proof should remain direct-sidecar-bootstrap based until wrappers have contracts, tests, and stable behavior.

The safest next action is a command-contract doc, followed by staged implementation.
