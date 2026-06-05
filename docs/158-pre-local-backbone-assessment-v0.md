# CarbonStack v0.4.7 Pre-local-backbone Assessment

Status: assessment and planning checkpoint
Scope: v0.4.x runtime Comms OpenMLS integration
Primary environment: WSL Debian
Generated: 2026-06-04 22:54:58 -0400

## 1. Purpose

This document records the v0.4.7 pre-local-backbone assessment after the v0.4.5 dev runtime OpenMLS smoke proof and v0.4.6 main-repo documentation alignment.

The immediate question is whether the committed carbonstack-comms smoke script should be promoted into a main carbonstack validation runner profile, whether it deserves local-backbone naming, and what cleanup/state/extraction work should happen before further runtime expansion.

Assessment result:

    Do not promote to local-backbone yet.
    Do not add a runner profile in this rung.
    Keep the smoke script as a comms-local dev helper for now.
    Plan a future dev-runtime-openmls profile only after boundary and cleanup semantics are explicit.
    Plan helper extraction before adding more runtime command wrappers.

## 2. Current repo heads

    carbonstack        3cd77a8 docs: record dev OpenMLS runtime smoke proof
    carbonstack-comms  8e6e8b4 test: add dev OpenMLS runtime smoke proof
    carbonstack-cypher 9ab994c docs: point to local operator runbook
    carbonstack-os     1bbbe52 docs: clarify CarbonStackOS target direction

## 3. Working tree status during assessment

    [carbonstack]
    (clean)

    [carbonstack-comms]
    (clean)

    [carbonstack-cypher]
    (clean)

    [carbonstack-os]
    (clean)

## 4. Assessment inputs

The v0.4.7 recon inspected:

    main carbonstack README claim/nonclaim surfaces
    docs/README.md current release and validation docs index
    docs/157-dev-runtime-openmls-smoke-proof-v0.md
    roadmap/ROADMAP.md v0.4.x section
    carbonstack validation runner architecture
    tools/carbonstack-validate profile dispatch
    tools/carbonstack-validate README profile boundaries
    local-cypher runner structure
    release-snapshot runner structure
    carbonstack-comms dev OpenMLS runtime smoke script
    carbonstack-comms README smoke/runtime command sections
    generated artifact ignore/cleanup policy
    carbonstack-comms commands.go size/function index
    app command tests
    state/trust path surfaces
    OpenMLS sidecar state/reset surfaces
    Cypher runtime/temp DB surfaces
    live smoke script behavior
    live validation baseline

The recon ended with smoke and validation passing, and the final repo status was clean.

## 5. Current proof state

Current proven runtime smoke path:

    openmls-send-dev -> Cypher -> openmls-inbox-dev --ack

The committed smoke script:

    carbonstack-comms/scripts/dev-openmls-runtime-smoke.sh

proves:

    temporary local Cypher server starts on dynamic loopback
    temporary Comms state files are created
    Alice/Bob Cypher devices are created
    dev-local OpenMLS sidecar identities are created
    dev-local OpenMLS sidecar conversation is bootstrapped
    openmls-send-dev submits an OpenMLS application-message artifact through Cypher
    openmls-inbox-dev retrieves, writes, sidecar-opens, prints, and acks the message
    ack happens after sidecar message-open success
    recipient inbox is empty after ack
    script self-terminates cleanly after cleanup hardening

## 6. Why this is still not local-backbone

The smoke proof is real and useful, but local-backbone naming is still premature.

Reasons:

    sidecar identity creation is still direct setup
    KeyPackage export is still direct sidecar setup
    Welcome generation and conversation join are still direct sidecar setup
    Relay Space join/onboarding does not exist yet
    old send/inbox commands remain stub-era
    state is dev-local and explicit-label based
    provider/trust/vault model is not mature
    generated sidecar state/reset semantics are not fully documented as a runner contract
    no hostile-server harness exists
    no public operator/deployment path exists
    no runner profile currently wraps the proof
    no release package currently claims this proof as a release validation surface

Correct label:

    dev runtime OpenMLS smoke proof

Not yet correct:

    local-backbone
    deployment profile
    production messaging UX
    production E2EE proof
    hostile-server proof

## 7. Runner-profile assessment

A future runner profile is plausible, but should not be added yet.

Recommended future profile name:

    dev-runtime-openmls

Avoid:

    local-backbone
    backbone
    production
    deploy
    server
    release-runtime

Reasoning:

    dev-runtime-openmls names what is actually proven
    it avoids implying deployment or production maturity
    it distinguishes this proof from local-cypher
    it distinguishes runtime Comms CLI proof from lower-level protocol tests

Current runner profiles are claim-careful:

    doctor reports toolchain/environment readiness
    core runs doctor, OpenMLS real-Cypher lifecycle, comms tests, cypher tests, and artifact scans
    local-cypher validates Cypher-only lifecycle and explicitly does not mean local-backbone
    full runs release-snapshot then local-cypher and is not deployment/runtime UX

A future dev-runtime-openmls profile should follow the same nonclaim discipline.

## 8. Why not add dev-runtime-openmls immediately

Do not add the profile in this rung because:

    the smoke proof has only just been documented at the main repo level
    the smoke script lives in carbonstack-comms and assumes sibling repo layout
    release-package behavior has not been tested with the smoke script as a formal runner step
    generated sidecar state cleanup is understood but not yet documented as a profile contract
    local-vs-release-root behavior needs explicit test/recon
    command output is still dev/pre-alpha
    KeyPackage/Welcome bootstrap is still direct sidecar setup
    commands.go extraction should be considered before more runtime growth

This is not a blocker against future promotion. It is a sequencing boundary.

## 9. Recommended future runner contract

If added later, dev-runtime-openmls should:

    call carbonstack-comms/scripts/dev-openmls-runtime-smoke.sh
    require sibling carbonstack, carbonstack-comms, and carbonstack-cypher layout
    print explicit boundary warnings
    run artifact scan before and after if integrated into carbonstack-validate
    clean known OpenMLS generated roots only when --clean-generated is set
    fail if the smoke script does not self-exit
    fail if plaintext does not match
    fail if ack does not happen after message-open success
    fail if recipient inbox is not empty after ack
    not imply deployment readiness
    not imply local-backbone
    not imply old send/inbox are OpenMLS-backed

Suggested runner wording:

    dev-runtime-openmls validates the current dev/pre-alpha OpenMLS application-message CLI path through Comms and Cypher.
    It is not local-backbone, not mature messaging UX, not deployment, and not a production/security claim.

## 10. State/reset assessment

Current smoke state behavior:

    uses mktemp workspace
    builds temporary Cypher binary under temp workspace
    uses temporary Cypher DB under temp workspace
    creates temporary Alice/Bob Comms state under temp workspace
    removes prior OpenMLS sidecar dev state before running
    creates sidecar generated state under the sidecar tree
    leaves known OpenMLS sidecar target/state roots possible until runner cleanup
    robustly stops temporary Cypher process with INT -> TERM -> KILL fallback

Current cleanup stance:

    acceptable for dev smoke
    acceptable when followed by core --clean-generated
    not yet strong enough to call release-grade profile behavior without explicit runner/docs contract

Needed before profile promotion:

    document generated paths
    document cleanup expectations
    confirm behavior from a fresh extracted release package
    confirm smoke script does not create tracked/untracked repo dirt beyond known generated roots
    confirm --clean-generated removes expected known generated roots afterward

## 11. commands.go extraction assessment

The current command implementation remains workable, but extraction pressure is real.

Reasons:

    existing stub-era commands remain in commands.go
    openmls-send-dev added sidecar protect parsing/glue
    openmls-inbox-dev added sidecar open parsing/glue
    future KeyPackage/Welcome/bootstrap wrappers would add more command-specific sidecar glue
    app tests now rely on package-level injection seams
    commands.go may become a runtime integration blob if more sidecar workflows are added directly

Do not extract blindly in this rung.

Recommended next code-structure step:

    inspect function boundaries
    choose a small extraction target
    probably move OpenMLS runtime command glue into a dedicated internal/app/openmls_runtime.go or internal/app/openmls_commands.go file
    keep CLI registration in commands.go
    keep tests near app command behavior
    do not change behavior during extraction

Best future sequence:

    plan extraction
    do behavior-preserving extraction
    validate
    then add any KeyPackage/Welcome/bootstrap command wrappers

## 12. KeyPackage/Welcome/bootstrap assessment

The smoke script currently uses direct sidecar setup for:

    identity-create
    public-bundle-export
    conversation-create
    conversation-add-member
    conversation-join

This is acceptable for a dev smoke proof.

It is not enough for mature runtime Comms UX.

Before local-backbone naming becomes reasonable, consider whether CarbonStackComms needs explicit dev CLI wrappers for:

    OpenMLS identity create/status
    public bundle export
    conversation create
    add member / produce Welcome
    conversation join from Welcome
    conversation load-check

These can remain dev-only.

They should not be confused with final Relay Space join UX.

## 13. Roadmap correction

The v0.4.5 smoke proof doc said the recommended next rung was v0.4.6 pre-local-backbone assessment.

Since v0.4.6 was used for main-repo smoke-proof docs/status alignment, the pre-local-backbone assessment is now v0.4.7.

This document records that correction.

## 14. Recommended next rung

Recommended next rung:

    v0.4.8 behavior-preserving command/helper extraction recon

Focus:

    inspect internal/app command boundaries
    choose a low-risk extraction plan
    avoid changing runtime behavior
    keep openmls-send-dev and openmls-inbox-dev outputs stable
    preserve tests
    validate smoke and core afterward

Alternative if prioritizing runner work:

    v0.4.8 dev-runtime-openmls runner-profile recon from fresh package-like root

Preferred order:

    extraction recon first, then runner-profile implementation later.

Rationale:

    more runtime work and future bootstrap wrappers will become easier if OpenMLS command glue is organized before profile promotion.

## 15. Nonclaims preserved

This assessment does not claim:

    local-backbone exists
    send/inbox are OpenMLS-backed
    mature Comms runtime UX exists
    Relay Space join exists
    production readiness
    production E2EE product readiness
    hostile-server safety
    metadata privacy
    secure vault/storage
    Android readiness
    CarbonStackOS readiness
    public ingress safety
    systemd/cloudflared readiness
    real homelab deployment safety
    external audit or certification
    post-quantum or hybrid security

## 16. Summary

v0.4.7 concludes that CarbonStack should keep the v0.4.5 smoke script as a comms-local dev proof for now.

A future dev-runtime-openmls runner profile is plausible, but local-backbone naming remains premature.

The next safest work is to assess and likely perform behavior-preserving OpenMLS command/helper extraction before adding more runtime wrappers or promoting the smoke proof into the main validation runner.
