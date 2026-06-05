# CarbonStack v0.4.9 Dev-runtime-openmls Runner Profile

Status: manual validation profile implementation
Scope: v0.4.x runtime Comms OpenMLS integration
Primary environment: WSL Debian
Generated: 2026-06-05 00:16:28 -0400

## 1. Purpose

This document records the first main `carbonstack` runner profile for the current dev runtime OpenMLS CLI smoke proof.

New manual profile:

    dev-runtime-openmls

The profile wraps the existing `carbonstack-comms` smoke script:

    carbonstack-comms/scripts/dev-openmls-runtime-smoke.sh

It validates the current dev/pre-alpha OpenMLS application-message CLI path:

    openmls-send-dev -> Cypher -> openmls-inbox-dev --ack

## 2. Current repo heads before this rung

    carbonstack        dc2d16c docs: assess pre-local-backbone runtime proof
    carbonstack-comms  0a48ae1 refactor: split OpenMLS runtime command helpers
    carbonstack-cypher 9ab994c docs: point to local operator runbook
    carbonstack-os     1bbbe52 docs: clarify CarbonStackOS target direction

## 3. What the profile does

The `dev-runtime-openmls` profile:

    checks required sibling repo paths
    enforces a live git umbrella checkout for now
    reports relevant toolchains
    runs a non-destructive pre-profile artifact scan
    calls `carbonstack-comms/scripts/dev-openmls-runtime-smoke.sh`
    runs a non-destructive post-profile artifact scan
    prints explicit claim boundaries

The smoke script itself:

    starts a temporary local Cypher server
    creates temporary Comms state files
    creates dev-local OpenMLS sidecar identities
    bootstraps a dev-local OpenMLS sidecar conversation
    sends through `openmls-send-dev`
    retrieves/opens through `openmls-inbox-dev --ack`
    verifies plaintext
    verifies ack after sidecar success
    verifies recipient inbox is empty after ack

## 4. How to run it

From the live umbrella checkout:

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go run . --profile dev-runtime-openmls --clean-generated

`--clean-generated` is recommended because the OpenMLS sidecar and Rust build can create known generated roots:

    carbonstack-comms/internal/protocol/mls/openmls-sidecar/.carbonstack-openmls-sidecar-state
    carbonstack-comms/internal/protocol/mls/openmls-sidecar/target

## 5. Boundary

This profile is manual-only for now.

It is not included in:

    full
    release-snapshot
    v0.4.0 release package validation

It does not prove:

    local-backbone
    mature messaging UX
    production readiness
    production E2EE
    hostile-server safety
    metadata privacy
    deployment readiness
    systemd/cloudflared readiness
    Android readiness
    CarbonStackOS readiness
    post-quantum or hybrid security
    external audit or certification

Existing `send` and `inbox` remain stub-era.

## 6. Why live umbrella only

The profile currently requires live git checkout markers for:

    carbonstack/.git
    carbonstack-comms/.git
    carbonstack-cypher/.git

This intentionally prevents treating the profile as a release-package validation surface before fresh-root behavior is tested.

A future release-root version may be added later after package behavior, generated-state cleanup, and claim boundaries are tested explicitly.

## 7. Why this is not local-backbone

`local-backbone` remains reserved because:

    sidecar identity creation is still direct dev setup
    KeyPackage export is still direct dev setup
    Welcome generation and conversation join are still direct dev setup
    Relay Space join/onboarding does not exist yet
    old send/inbox remain stub-era
    state is dev-local and explicit-label based
    provider/trust/vault model is not mature
    no hostile-server harness exists
    no public operator/deployment path exists

Correct label:

    dev-runtime-openmls

Not correct:

    local-backbone

## 8. Relationship to v0.4.8

v0.4.8 split OpenMLS runtime command helpers out of `commands.go` into a dedicated app file in `carbonstack-comms`.

That behavior-preserving cleanup reduced command-glue risk before promoting the smoke proof into a main runner profile.

## 9. Suggested next rung

Recommended next rung:

    v0.4.10 validation/profile polish or sidecar bootstrap wrapper recon

Possible focus:

    verify dev-runtime-openmls from a clean live checkout state
    inspect whether generated-root cleanup should happen before the profile starts
    decide whether sidecar bootstrap dev wrappers should come next
    keep full unchanged until dev-runtime-openmls maturity testing is complete
    keep local-backbone reserved

## 10. Summary

v0.4.9 adds the manual `dev-runtime-openmls` runner profile.

This improves repeatability for the dev runtime OpenMLS CLI smoke proof while preserving strict nonclaims.
