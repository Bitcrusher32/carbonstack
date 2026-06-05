# CarbonStack v0.4.10 Dev-runtime-openmls Profile Boundary Check

Status: boundary/polish documentation checkpoint
Scope: v0.4.x runtime Comms OpenMLS integration
Primary environment: WSL Debian
Generated: 2026-06-05 00:30:12 -0400

## 1. Purpose

This document records the v0.4.10 boundary check for the manual `dev-runtime-openmls` validation profile added in v0.4.9.

This rung does not change runner behavior.

It records that the profile boundary is behaving as intended:

    repeatable from the live umbrella checkout
    manual-only
    live-git-umbrella-only
    not included in `full`
    not a release-package validation surface
    not local-backbone
    not production messaging UX

## 2. Current repo heads before this docs checkpoint

    carbonstack        8eeadb2 feat: add dev OpenMLS runtime validation profile
    carbonstack-comms  0a48ae1 refactor: split OpenMLS runtime command helpers
    carbonstack-cypher 9ab994c docs: point to local operator runbook
    carbonstack-os     1bbbe52 docs: clarify CarbonStackOS target direction

## 3. What was checked

The v0.4.10 scout inspected and validated:

    current dev-runtime-openmls runner implementation
    runner dispatch around dev-runtime-openmls and full
    runner README profile documentation
    docs/159-dev-runtime-openmls-runner-profile-v0.md
    docs index and roadmap references
    direct live dev-runtime-openmls repeatability
    behavior with and without --clean-generated
    generated-root cleanup behavior
    full profile separation
    non-git package-like root refusal
    final baseline validation

## 4. Repeatability result

The manual profile passed from the live umbrella checkout:

    go run . --profile dev-runtime-openmls

It also passed again with explicit cleanup:

    go run . --profile dev-runtime-openmls --clean-generated

The profile continued to validate:

    openmls-send-dev -> Cypher -> openmls-inbox-dev --ack

The smoke proof verified:

    plaintext matched
    message was opened
    ack happened after sidecar open success
    recipient inbox was empty after ack

## 5. Generated artifact behavior

Running `dev-runtime-openmls` without `--clean-generated` produced only known OpenMLS sidecar generated roots:

    carbonstack-comms/internal/protocol/mls/openmls-sidecar/.carbonstack-openmls-sidecar-state
    carbonstack-comms/internal/protocol/mls/openmls-sidecar/target

This is expected after sidecar smoke tests.

`core --clean-generated` and `dev-runtime-openmls --clean-generated` both removed known generated roots.

Current policy remains:

    artifact scans are non-destructive by default
    generated roots are visible unless cleanup is explicitly requested
    --clean-generated removes known generated/build roots only
    known generated roots are acceptable after smoke/tests if they stay inside recognized paths

## 6. Full profile boundary

The runner dispatch still keeps `full` separate.

`full` currently runs:

    release-snapshot
    local-cypher

It does not call:

    dev-runtime-openmls

This is intentional.

During the v0.4.10 scout, running `full` from the live umbrella failed during release-snapshot metadata checks because the live umbrella did not contain the expected release metadata directory.

That is not a dev-runtime-openmls failure.

Correct interpretation:

    full remains a release-package validation ladder
    full should be run from a fresh extracted or throwaway staged release package root
    dev-runtime-openmls remains a live-umbrella manual profile
    the two profiles intentionally serve different validation surfaces

## 7. Non-git package-like root refusal

The profile was tested against a copied package-like root with `.git` directories removed.

Expected result:

    failure

Observed result:

    dev-runtime-openmls refused to run
    missing live checkout markers were reported for carbonstack, carbonstack-comms, and carbonstack-cypher
    the command exited nonzero
    the scout printed PASS for the refusal check

This confirms the live-git-umbrella guard is working as intended.

## 8. Current validated profile boundary

`dev-runtime-openmls` currently means:

    manual live-umbrella dev/pre-alpha smoke validation
    OpenMLS application-message runtime CLI proof through Comms and Cypher
    wrapper around carbonstack-comms/scripts/dev-openmls-runtime-smoke.sh
    explicit not-local-backbone profile
    explicit not-production/security profile
    not included in full
    not release-package validation yet

It does not mean:

    local-backbone
    release validation
    production messaging UX
    production E2EE
    hostile-server safety
    metadata privacy
    deployment readiness
    Android readiness
    CarbonStackOS readiness
    post-quantum or hybrid security
    audit/certification

## 9. Why no code change in v0.4.10

The profile behaved as intended.

No code patch is needed because:

    repeatability passed
    cleanup behavior matched expectations
    full remained separate
    non-git package-like root refusal worked
    final validation passed
    final repo status was clean

This checkpoint exists to record the evidence and clarify the profile boundary.

## 10. Suggested next rung

Recommended next rung:

    v0.4.11 sidecar bootstrap wrapper recon

Focus:

    inspect whether direct sidecar setup steps in scripts/dev-openmls-runtime-smoke.sh should become explicit dev-only Comms commands
    preserve dev-only naming
    avoid replacing send/inbox yet
    avoid local-backbone naming
    avoid putting dev-runtime-openmls into full
    keep v0.5.x state/trust/vault/PQ work deferred

Candidate wrapper areas:

    OpenMLS identity create/status
    public bundle export
    conversation create
    add member / Welcome generation
    conversation join
    conversation load-check

## 11. Summary

v0.4.10 confirms the v0.4.9 `dev-runtime-openmls` runner profile boundary is stable enough to preserve.

The profile is repeatable, cleanup behavior is understood, non-git package-like roots are refused, and `full` remains separate.

The next safest work is sidecar bootstrap wrapper recon, not local-backbone promotion.
