# CarbonStack v0.3.28 Local Operator Helper and Runner Decision Recon

Status: v0 decision / recon checkpoint
Scope: v0.3.x local-only backbone deployability line
Primary environment: WSL Debian
Generated: 2026-06-02 16:00:24 -0400

## 1. Purpose

This document records the v0.3.28 decision/recon checkpoint after the v0.3.27 local Cypher explicit-env API lifecycle proof.

v0.3.27 proved that the current Cypher API lifecycle can run under the documented explicit local operator environment, including invite claim, device registration, opaque envelope submit/retrieve/ack, restart against the same DB, and persisted state checks.

v0.3.28 does not add tooling. It decides what should not be automated yet, what should remain manual, what names should be reserved, and what validation semantics should be used later.

This is intentionally docs-only.

## 2. Current Repo Heads

    carbonstack        e516fc7 docs: record local Cypher API lifecycle proof
    carbonstack-comms  012c8bf scripts: support source snapshot artifact guard
    carbonstack-cypher 9ab994c docs: point to local operator runbook
    carbonstack-os     1bbbe52 docs: clarify CarbonStackOS target direction

## 3. Working Tree Status Before This Doc

    [carbonstack]
    (clean)

    [carbonstack-comms]
    (clean)

    [carbonstack-cypher]
    (clean)

    [carbonstack-os]
    (clean)

## 4. Toolchain Snapshot

    git:     git version 2.47.3
    go:      go version go1.24.4 linux/amd64
    rustc:   rustc 1.96.0 (ac68faa20 2026-05-25)
    cargo:   cargo 1.96.0 (30a34c682 2026-05-25)
    sqlite3: 3.46.1 2024-08-13 09:16:08 c9c2ab54ba1f5f46360f1b4f35d849cd3f080e6fc2b6c60e91b16c63f69aalt1 (64-bit)

## 5. Decision Summary

v0.3.28 decisions:

    Do not add helper tooling yet.
    Do not add a local-cypher runner profile yet.
    Do not add a local-backbone runner profile yet.
    Use temporary isolated DBs for dev/test validation.
    Do not use the documented manual operator DB path for automated validation by default.
    Reserve local-backbone for later whole-stack validation, not current Cypher-only lifecycle checks.
    Use local-cypher as the future candidate name for Cypher-only lifecycle validation, if/when added.
    Keep public ingress, systemd, cloudflared, and homelab paths deferred.
    Keep runtime Comms OpenMLS UX deferred to v0.4.x.

## 6. Why No Helper Yet

A tiny helper could eventually be useful.

Possible future helper shape:

    carbonstack local start
    carbonstack local health
    carbonstack local stop
    carbonstack local reset

But helper tooling should not be added yet.

Reasons:

    v0.3.27 proved a lifecycle manually, but did not define a mature operator UX.
    helper commands would freeze start/stop/reset semantics too early.
    reset behavior is intentionally blunt and should remain user-visible.
    helper tooling risks hiding important experimental state boundaries.
    config parser/helper decisions should follow a stable manual convention.

Current decision:

    no helper in v0.3.28
    no helper before the lifecycle contract is clearer
    prefer docs and explicit commands for now

## 7. Why No Runner Profile Yet

A future runner profile is plausible, but premature.

Current runner profiles already validate important project surfaces:

    doctor
    core
    full
    release-snapshot
    write-checksums
    verify-checksums

A local Cypher lifecycle profile would need a precise success/failure contract.

Open questions before adding such a profile:

    Should the runner build and start Cypher itself?
    Should the runner target an already-running Cypher process?
    Should the runner own cleanup?
    Should it always use a temporary DB?
    Should it ever touch the documented manual operator DB path?
    Should negative API paths be included?
    Should the profile be Cypher-only or include Comms?
    Should it be release-blocking or informational?

Current decision:

    no local-cypher runner profile in v0.3.28
    no local-backbone runner profile in v0.3.28
    revisit after at least one more lifecycle proof/review rung

## 8. Naming Decision: local-cypher vs local-backbone

Future candidate name for Cypher-only validation:

    local-cypher

Reason:

    v0.3.27 proved a Cypher API lifecycle only.
    It did not validate runtime Comms UX.
    It did not validate end-to-end user messaging.
    It did not validate a full CarbonStack backbone.
    It did not include CarbonStack Relay Space mechanics.

Reserved future name:

    local-backbone

local-backbone should be reserved for a top-level runner/wrapper that validates the actual CarbonStack backbone when Comms/Cypher/OpenMLS are wired meaningfully enough to deserve that name.

Current rule:

    Do not call a Cypher-only lifecycle proof local-backbone.

## 9. Temporary DB Policy

Dev/test validation should use temporary isolated DBs by default.

This applies to:

    local lifecycle proof scripts
    future local-cypher runner profile, if added
    release-adjacent dev/test validation
    CI-like validation paths

Reason:

    avoids mutating a user's manual operator DB
    makes runs repeatable
    makes cleanup straightforward
    avoids stale state coupling
    avoids accidental reliance on local private state
    fits source/test separation established in v0.3.26

Important maturity boundary:

    temporary DB validation is acceptable for dev/test and pre-v0.4.x release-adjacent validation.

It should not be treated as the final model for public Comms use or real user testing.

Later, when CarbonStack matures toward public Comms and real-user testing, validation must also address real persistent operator/user state, backup/restore expectations, upgrade semantics, and failure recovery. Temporary DBs should not be the only evidence used at that stage.

## 10. Manual Operator DB vs Validation DB

Manual local operator path from v0.3.26:

    $HOME/.local/share/carbonstack/cypher/cypher.db

Automated proof/validation DB policy:

    use a temporary isolated DB path
    remove it after proof
    do not mutate the manual operator DB by default

This keeps the manual runbook and validation harness from stepping on each other.

## 11. Negative Path Decision

v0.3.27 accidentally produced a useful negative path:

    content_type=carbonstack.message.text.stub.v0
    protocol_version=carbonstack-openmls-sidecar-v0

Cypher correctly rejected that pair with:

    unsupported_protocol_version

That blunder should not be ignored.

Future local-cypher validation may include a deliberate negative-path section, but not yet.

Candidate negative paths:

    invalid content/protocol pairing
    invalid base64 envelope
    ack unknown envelope
    ack with wrong recipient
    malformed device lookup
    duplicate invite code
    already-claimed invite

Current decision:

    do not add negative path runner coverage in v0.3.28
    record it as future validation maturity work

## 12. 127.0.0.1 Default Decision

Current code default still allows:

    CYPHER_ADDR=:8080

Current local operator docs recommend:

    CYPHER_ADDR=127.0.0.1:8080

Decision:

    do not change the code default in v0.3.28

Reason:

    default bind change is real behavior, not docs cleanup
    it should be its own focused code/config rung if chosen
    current docs already require explicit local operator env
    v0.3.x should avoid mixing docs/recon with unrelated behavior changes

Future possible rung:

    change default bind to 127.0.0.1:8080
    update config tests
    update docs/README claims
    validate doctor/core

## 13. Release Semantics

v0.3.28 does not supersede v0.3.20 as the public runner-backed testing release.

Current public release remains:

    v0.3.20

v0.3.28 is mainline local deployability groundwork.

It is not:

    production deployability
    public server release
    homelab guide
    Comms app UX
    Android release
    CarbonStackOS work
    audited security claim

## 14. Recommended Next Rungs

Recommended near-term sequence:

    v0.3.29:
      local-cypher runner profile design doc or one more lifecycle proof review

    v0.3.30:
      implement local-cypher profile only if contract is concrete

    v0.3.31+:
      consider tiny helper script only if manual lifecycle remains stable

Alternative if the project wants to stay docs-only longer:

    v0.3.29:
      local-cypher validation contract doc

    v0.3.30:
      negative-path validation recon

    v0.3.31:
      local-cypher runner implementation

Avoid next:

    public ingress
    cloudflared
    systemd
    real homelab deployment
    local-backbone profile
    runtime Comms send/inbox UX
    Android app
    CarbonStackOS
    CarbonStack Relay Space implementation

## 15. Current Nonclaims

v0.3.28 does not validate:

    production deployability
    production E2EE
    hostile-server safety
    metadata privacy
    public ingress
    LAN exposure
    systemd
    cloudflared
    real homelab deployment
    remote admin plane
    runtime Comms OpenMLS UX
    Android app
    CarbonStackOS
    external audit
    certification

## 16. Summary

v0.3.28 is a pause-and-decide checkpoint after the successful v0.3.27 local Cypher API lifecycle proof.

It decides:

    no helper yet
    no runner profile yet
    future Cypher-only validation should be called local-cypher, not local-backbone
    temporary DBs are correct for dev/test validation
    local-backbone is reserved for a real whole-backbone validation surface
    public/homelab/systemd/cloudflared work remains deferred

This keeps CarbonStack's local deployability line evidence-led without prematurely hardening the wrong abstraction.
