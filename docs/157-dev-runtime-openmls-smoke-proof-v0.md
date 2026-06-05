# CarbonStack v0.4.5 Dev Runtime OpenMLS Smoke Proof

Status: recorded smoke proof / main-repo status alignment
Scope: v0.4.x runtime Comms OpenMLS integration
Primary environment: WSL Debian
Generated: 2026-06-04 22:30:14 -0400

## 1. Purpose

This document records the first dev runtime OpenMLS CLI smoke proof after the v0.4.3 and v0.4.4 implementation rungs.

v0.4.3 added the explicit dev-only send command in carbonstack-comms:

    openmls-send-dev

v0.4.4 added the explicit dev-only receive/open/optional-ack command in carbonstack-comms:

    openmls-inbox-dev

v0.4.5 proved the application-message runtime CLI path:

    openmls-send-dev -> Cypher -> openmls-inbox-dev --ack

This document records the result in the main carbonstack repo so that the project-level documentation and roadmap match the current implementation state.

## 2. Current repo heads at this alignment point

    carbonstack        a799764 docs: define runtime OpenMLS command contract
    carbonstack-comms  8e6e8b4 test: add dev OpenMLS runtime smoke proof
    carbonstack-cypher 9ab994c docs: point to local operator runbook
    carbonstack-os     1bbbe52 docs: clarify CarbonStackOS target direction

## 3. Working tree status during alignment

    [carbonstack]
    (clean)

    [carbonstack-comms]
    (clean)

    [carbonstack-cypher]
    (clean)

    [carbonstack-os]
    (clean)

## 4. What was proven

The committed carbonstack-comms smoke script:

    scripts/dev-openmls-runtime-smoke.sh

proves the current dev runtime application-message path:

    openmls-send-dev -> Cypher -> openmls-inbox-dev --ack

The smoke proof creates:

    a temporary local Cypher server
    a temporary Cypher database
    temporary Comms local state files
    dev-local OpenMLS sidecar identities
    a dev-local OpenMLS sidecar conversation
    a registered Alice device
    a registered Bob device

It then:

    sends plaintext through openmls-send-dev
    calls the OpenMLS sidecar message-protect path
    submits the resulting OpenMLS application-message artifact through Cypher
    retrieves the message through openmls-inbox-dev
    writes the OpenMLS application-message artifact from the envelope
    calls the OpenMLS sidecar message-open path
    prints plaintext only after sidecar open succeeds
    acks the envelope only after sidecar open succeeds
    verifies Bob's inbox is empty after ack

## 5. Observed successful output

The successful smoke run reported:

    PASS: dev runtime OpenMLS CLI smoke proof
    proof: openmls-send-dev -> Cypher -> openmls-inbox-dev --ack
    plaintext: hello bob through openmls runtime smoke
    boundary: dev/pre-alpha smoke proof; not local-backbone; not production messaging UX

A self-termination cleanup issue was observed during the first successful run. The proof had passed, but the script required manual interrupt because the temporary Cypher process cleanup path waited too long. The script was patched to send INT, then TERM, then KILL if needed. After the patch, the smoke script printed PASS and returned to the shell by itself.

## 6. Validation result

v0.4.5 validation passed with:

    scripts/dev-openmls-runtime-smoke.sh
    go test ./internal/app -count=1
    go test ./... -count=1 -timeout 600s
    go run . --profile local-cypher
    go run . --profile doctor
    go run . --profile core --clean-generated

The final state after v0.4.5 was clean across the four repos.

## 7. Why this matters

Before v0.4.5, CarbonStack had lower-level OpenMLS/Cypher lifecycle proofs and two explicit dev runtime commands, but it did not yet have a committed runtime CLI smoke proof using both commands together.

v0.4.5 moves the runtime integration state from:

    commands exist separately

to:

    the dev application-message CLI path is repeatably smoke-tested

This is an important v0.4.x maturity step.

## 8. Boundary and nonclaims

This smoke proof is still dev/pre-alpha.

It does not prove:

    production readiness
    production E2EE product readiness
    hostile-server safety
    metadata privacy
    secure local vault/storage
    Android readiness
    CarbonStackOS readiness
    public ingress safety
    systemd readiness
    cloudflared readiness
    real homelab deployment safety
    external audit or certification
    post-quantum or hybrid security
    polished messaging UX
    CarbonStack Relay Space mechanics
    mature onboarding
    local-backbone readiness

The smoke script still uses direct sidecar setup for identity, KeyPackage, Welcome, and conversation bootstrapping. The runtime CLI proof target is the application-message path, not the full future onboarding or Relay Space join model.

Existing stub-era commands remain:

    send
    inbox
    ack

Do not describe old send/inbox as OpenMLS-backed yet.

## 9. Current allowed claim

Allowed claim:

    CarbonStackComms now has a dev-only runtime smoke proof for the OpenMLS application-message path through Cypher using openmls-send-dev and openmls-inbox-dev --ack.

More specific allowed claim:

    The smoke script starts a temporary Cypher server, creates temporary local Comms state, bootstraps dev-local sidecar OpenMLS state, sends an OpenMLS application-message artifact through openmls-send-dev, retrieves and opens it through openmls-inbox-dev, verifies plaintext, acks after sidecar success, and verifies the recipient inbox is empty after ack.

## 10. Why this is not local-backbone yet

local-backbone remains reserved.

Reasons:

    setup still uses direct sidecar bootstrap commands
    KeyPackage and Welcome setup are not yet first-class Comms runtime flows
    state model remains dev-local and explicit-label based
    no production-safe provider storage exists
    no mature local vault exists
    no hostile-server harness exists
    no public deployability or operator path exists
    no runner profile wraps this proof yet
    no public release package includes this as a known-good release artifact yet

It is acceptable to call this:

    dev runtime OpenMLS smoke proof

It is not yet acceptable to call it:

    local-backbone

## 11. Suggested next rung

Recommended next rung:

    v0.4.6 pre-local-backbone assessment

Focus:

    decide whether the smoke proof should become a validation runner profile
    decide whether sidecar bootstrap must be wrapped before local-backbone naming is justified
    document current state requirements and reset semantics
    identify whether commands.go needs helper extraction before further runtime growth
    avoid changing public release framing until a new release is intentionally cut

## 12. Summary

v0.4.5 proves the first dev runtime OpenMLS CLI application-message path.

The project now has:

    openmls-send-dev
    openmls-inbox-dev
    a committed smoke script that proves the two commands work together through Cypher

The correct next step is careful assessment and documentation before naming this local-backbone or adding a broader runner profile.
