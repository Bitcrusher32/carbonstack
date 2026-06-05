# CarbonStack v0.4.17 Wrapper OpenMLS Runtime Runner Profile

Status: runner/profile alignment checkpoint
Scope: v0.4.x runtime Comms OpenMLS validation
Primary environment: WSL Debian
Generated: 2026-06-05 05:50:46 -0400

## 1. Purpose

This document records the v0.4.17 runner/profile alignment checkpoint for the wrapper-based OpenMLS runtime smoke proof.

v0.4.16 added a parallel wrapper-based smoke script in CarbonStackComms:

    carbonstack-comms/scripts/dev-openmls-runtime-smoke-wrappers.sh

v0.4.17 promotes that script into a separate manual validation runner profile:

    dev-runtime-openmls-wrappers

This does not replace the original direct-sidecar smoke profile.

## 2. Current repo heads before this checkpoint

    carbonstack        0583683 docs: define dev OpenMLS bootstrap command contract
    carbonstack-comms  cb4e59d test: add wrapper-based OpenMLS runtime smoke proof
    carbonstack-cypher 9ab994c docs: point to local operator runbook
    carbonstack-os     1bbbe52 docs: clarify CarbonStackOS target direction

## 3. Profile added

New profile:

    dev-runtime-openmls-wrappers

Command:

    go run . --profile dev-runtime-openmls-wrappers --clean-generated

Proof shape:

    openmls-*-dev bootstrap wrappers -> openmls-send-dev -> Cypher -> openmls-inbox-dev --ack

Wrapped script:

    carbonstack-comms/scripts/dev-openmls-runtime-smoke-wrappers.sh

## 4. Existing profile preserved

Existing profile remains unchanged:

    dev-runtime-openmls

It continues to wrap:

    carbonstack-comms/scripts/dev-openmls-runtime-smoke.sh

Existing proof shape remains:

    direct sidecar bootstrap -> openmls-send-dev -> Cypher -> openmls-inbox-dev --ack

## 5. Why separate profiles

The wrapper smoke is newer than the original direct-sidecar smoke. Keeping it separate allows wrapper maturity work without weakening the known-good direct baseline.

`dev-runtime-openmls` remains the direct-smoke baseline.

`dev-runtime-openmls-wrappers` is the wrapper-smoke maturity surface.

They may be merged later only after wrapper behavior, reset semantics, runner semantics, and docs all justify that maturity step.

## 6. Boundaries

This checkpoint does not:

    replace the original smoke script
    replace dev-runtime-openmls
    add either runtime profile to full
    add release-package runtime validation
    introduce local-backbone naming
    implement Relay Space join UX
    make production security claims
    change old send/inbox behavior
    start v0.5.x state/trust/vault/PQ work

## 7. Runner behavior

The new profile follows the same broad guard pattern as dev-runtime-openmls:

    required path checks
    live git umbrella guard
    toolchain reporting
    pre-profile artifact scan
    wrapper smoke script execution
    post-profile artifact scan
    explicit nonclaim summary

It remains:

    manual-only
    live-umbrella-only
    not included in full
    not release-package validation

## 8. Why this is not local-backbone

`local-backbone` remains reserved because:

    old send/inbox remain stub-era
    runtime profiles are still dev/pre-alpha
    sidecar/provider state reset semantics are not mature
    trust/vault/provider state is not production-designed
    Relay Space join/onboarding does not exist
    hostile-server harnesses remain future work
    release-package runtime validation remains unsupported

Correct phrase:

    wrapper OpenMLS runtime validation profile

Incorrect phrase:

    local-backbone

## 9. Validation expectation

The v0.4.17 checkpoint should validate:

    go test ./... -count=1 in carbonstack/tools/carbonstack-validate
    go run . --profile dev-runtime-openmls-wrappers --clean-generated
    go run . --profile dev-runtime-openmls --clean-generated
    go run . --profile local-cypher
    go run . --profile doctor
    go run . --profile core --clean-generated

CarbonStackComms should also continue to pass:

    go test ./internal/app -count=1
    go test ./... -count=1 -timeout 600s

## 10. Suggested next rung

Recommended next rung after v0.4.17:

    pre-v0.5.x command registry recon/planning

Reason:

    command sprawl is now real
    v0.5.x will likely add state/trust/vault/PQ command surfaces
    a machine-readable command registry before v0.5.x can prevent README sprawl

## 11. Summary

v0.4.17 adds a separate manual runner profile for wrapper-based OpenMLS runtime validation while preserving the original direct-sidecar smoke profile as the baseline.

This is a maturity/alignment checkpoint, not a production/security claim.
