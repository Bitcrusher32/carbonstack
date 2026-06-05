# CarbonStack v0.4.18 Known-Good Command Registry Recon

Status: registry implementation checkpoint
Scope: cross-repo command/profile/script/API surface inventory
Primary environment: WSL Debian
Generated: 2026-06-05 07:14:11 -0400

## 1. Purpose

v0.4.18 creates the first provisional known-good command registry for CarbonStack.

The registry lives in the main `carbonstack` repo because it spans:

    carbonstack
    carbonstack-comms
    carbonstack-cypher
    carbonstack-os placeholders / future surfaces
    release validation
    runner profiles
    smoke scripts
    Comms CLI commands
    OpenMLS sidecar commands
    Cypher CLI/API/operator surfaces
    legacy PowerShell helpers

## 2. Current repo heads before this checkpoint

    carbonstack        a501442 feat: add wrapper OpenMLS runtime validation profile
    carbonstack-comms  cb4e59d test: add wrapper-based OpenMLS runtime smoke proof
    carbonstack-cypher 9ab994c docs: point to local operator runbook
    carbonstack-os     1bbbe52 docs: clarify CarbonStackOS target direction

## 3. Files added

    registry/README.md
    registry/commands.v0.yaml

## 4. Registry purpose

The registry is meant to prevent command sprawl from becoming invisible before v0.5.x adds more state/trust/vault/PQ surfaces.

It tracks:

    command
    repo
    component
    kind
    audience
    maturity
    introduced_in where known
    source_path
    validation_surface
    short_help
    why_exists
    example where useful
    nonclaims
    related profiles/scripts
    replacement/deprecation status
    include_in_front_readme

## 5. Surfaces covered in v0

The initial registry covers:

    carbonstack validation runner profiles:
      doctor
      core
      local-cypher
      dev-runtime-openmls
      dev-runtime-openmls-wrappers
      full
      release-snapshot
      write-checksums
      verify-checksums

    carbonstack-comms CLI:
      init
      dev-create-invite
      claim-invite
      register-device
      list-devices
      fingerprint
      verify-device
      trust-history
      trust-list
      simulate-key-change
      revoke-device
      send
      inbox
      ack
      openmls-send-dev
      openmls-inbox-dev
      openmls-identity-create-dev
      openmls-identity-status-dev
      openmls-bundle-export-dev
      openmls-conversation-create-dev
      openmls-conversation-load-check-dev
      openmls-conversation-add-member-dev
      openmls-conversation-join-dev

    carbonstack-comms scripts:
      dev-openmls-runtime-smoke.sh
      dev-openmls-runtime-smoke-wrappers.sh
      self-test-openmls-backbone.ps1
      smoke-openmls-real-cypher-relay.ps1
      check-no-rust-artifacts.ps1
      test-local-lifecycle.ps1
      test-trust-lifecycle.ps1

    OpenMLS sidecar commands:
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
      state-checkpoint as unsupported/future
      state-load-check as unsupported/future

    carbonstack-cypher:
      go run ./cmd/cypher
      key env vars
      health/invite/device/envelope/ack API surfaces

    carbonstack legacy scripts:
      scripts/validate-local.ps1
      scripts/validate-phase1.ps1

## 6. Key policy decisions

Registry lives in:

    carbonstack/registry/commands.v0.yaml

It should not be buried in `carbonstack-comms`, because it is a cross-stack and release-independent map.

Old stub-era `send`, `inbox`, and `ack` are included as legacy/stub-era surfaces so they cannot be confused with the newer OpenMLS runtime commands.

The direct and wrapper OpenMLS runtime profiles remain separate:

    dev-runtime-openmls
    dev-runtime-openmls-wrappers

## 7. What this does not do

v0.4.18 does not:

    change command behavior
    change runner behavior
    add local-backbone
    add runtime profiles to full
    replace direct smoke with wrapper smoke
    remove old send/inbox
    start v0.5.x state/trust/vault/PQ work
    move command docs out of front-door READMEs yet
    generate local help text automatically yet

## 8. Why this matters before v0.5.x

v0.5.x is expected to increase command/profile/API surface area around:

    local storage
    provider state
    trust state
    vault design
    PQ/hybrid migration readiness
    migration/rekey/recovery concepts

Without a registry, the project risks README sprawl and semantic drift.

## 9. Suggested next rung

Recommended next rung:

    v0.4.19 command registry validation / local-help planning

Candidate work:

    verify registry entries against current dispatch/help output
    decide whether registry can generate component command references
    decide whether `commands.go` should gain short explainer strings beside command registration
    decide whether to add a lightweight registry validation script
    keep front-door README changes minimal

## 10. Summary

v0.4.18 introduces a provisional cross-repo command registry as command-surface hygiene before v0.5.x.

This is a documentation/metadata maturity checkpoint, not a new runtime or security capability.
