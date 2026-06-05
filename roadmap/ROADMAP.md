# CarbonStack Roadmap

This roadmap describes the current public direction of CarbonStack.

Older numbered docs preserve older plans and implementation history. Use this file, the top-level README, the latest docs index, and release-specific runbooks for current public-facing direction.

## Current state after v0.3.32

CarbonStack is in the late v0.3.x local deployability validation line.

Current public testing release:

    v0.3.20 runner-backed testing release

Current mainline validation state:

    v0.3.32 local-cypher negative protocol validation

Mainline WSL Debian validation now includes:

    go run . --profile local-cypher
    go run . --profile doctor
    go run . --profile core --clean-generated

`local-cypher` is Cypher-only. It validates local Cypher lifecycle behavior with temporary state, dynamic loopback bind, invite/device/envelope/ack lifecycle, restart against the same DB, cleanup, and one negative protocol pairing.

`local-cypher` is not `local-backbone`.

`local-backbone` remains reserved for later whole-stack validation after Comms runtime UX is actually wired through the backbone.

## Current nonclaims

CarbonStack is not currently:

    production-ready
    production E2EE
    hostile-server safe
    metadata-private
    deployment-ready
    systemd-ready
    cloudflared-ready
    Android-ready
    CarbonStackOS-ready
    externally audited
    certified

## Late v0.3.x runway

### v0.3.33 — pre-v0.4.0 release-surface cleanup

Goal:

    refresh README/docs/roadmap public wording
    remove stale v0.2.x/v0.3.0-era current-state language from public surfaces
    surface v0.3.32 as the current mainline validation state
    preserve v0.3.20 as current public testing release
    frame v0.4.0 as broad local deployability pre-release
    avoid new feature work

### v0.3.34 — full profile release validation ladder

Goal:

    make `full` the v0.4.0 release-package validation ladder
    confirm release-snapshot/checksum/package expectations
    define full as release-snapshot plus local-cypher
    avoid duplicated core execution because release-snapshot already calls core
    keep release claims narrow


Recommended v0.4.0 validation command shape:

    go run . --profile full --root /path/to/release-package-root --clean-generated

### v0.3.35 / v0.3.36 — v0.4.0 package rehearsal

Goal:

    stage a throwaway v0.4.0-style package root
    include carbonstack, carbonstack-comms, carbonstack-cypher, and release metadata
    exclude carbonstack-os from the runnable package
    generate release checksums
    archive and fresh-extract the package
    verify checksums from the fresh extraction
    run full from the fresh extraction with --clean-generated
    document package rehearsal result

Result:

    v0.3.36 package rehearsal passed under WSL Debian.

### v0.3.37 — release candidate wording and asset plan

Goal:

    draft v0.4.0 release notes using v0.3.20 release presentation as continuity template
    define attached asset expectations
    define validation instructions
    define known-good toolchains
    define nonclaims
    decide whether another RC validation pass is needed

### v0.3.36+ — pre-v0.4.0 release cut, if clean

Goal:

    cut v0.4.0 only if release-surface cleanup, package validation, checksums, release instructions, and nonclaims are coherent

## v0.4.x current rung

### v0.4.1 — runtime Comms/OpenMLS/Cypher recon

Result:

    Current Comms CLI send/inbox/ack exist, but send/inbox are still stub-era.
    The OpenMLS/Cypher relay seam is already validated through internal relay/protocol tests.
    The next implementation should use explicit dev-only OpenMLS runtime commands before replacing send/inbox.

Recommended next rung:

    v0.4.2 runtime OpenMLS command contract

### v0.4.2 — runtime OpenMLS command contract

Goal:

    Define explicit dev-only runtime OpenMLS commands before implementation.

Decision:

    Use openmls-send-dev and openmls-inbox-dev first.
    Do not silently replace existing send/inbox yet.
    Ack remains consume-success gated.
    Keep local-backbone reserved until whole-path runtime validation deserves it.

Recommended next rung:

    v0.4.3 first dev-only OpenMLS send command

### v0.4.5 — dev runtime OpenMLS smoke proof

Result:

    carbonstack-comms now includes scripts/dev-openmls-runtime-smoke.sh.
    The script proves openmls-send-dev -> Cypher -> openmls-inbox-dev --ack for the OpenMLS application-message path.
    It verifies plaintext and confirms the recipient inbox is empty after ack.

Boundary:

    This is a dev/pre-alpha smoke proof.
    This is not local-backbone.
    This is not production messaging UX.
    Sidecar KeyPackage/Welcome/bootstrap setup remains direct dev setup.
    Existing send/inbox remain stub-era.

Recommended next rung:

    v0.4.7 pre-local-backbone assessment and validation-profile decision.

### v0.4.7 — pre-local-backbone assessment

Result:

    local-backbone remains reserved.
    dev-runtime-openmls is the preferred future runner-profile name if the smoke proof is promoted.
    The v0.4.5 smoke script should remain a carbonstack-comms dev helper until runner boundaries, generated-state cleanup, and release-root behavior are documented and tested.
    commands.go/helper extraction should be considered before adding more OpenMLS runtime wrappers.

Boundary:

    This is an assessment and planning checkpoint.
    No runner profile is added yet.
    Existing send/inbox remain stub-era.
    Direct sidecar KeyPackage/Welcome/bootstrap setup remains dev setup.
    This is not production messaging UX.

Recommended next rung:

    v0.4.8 behavior-preserving OpenMLS command/helper extraction recon, or dev-runtime-openmls runner-profile recon if runner promotion becomes the priority.

### v0.4.9 — manual dev-runtime-openmls runner profile

Result:

    carbonstack/tools/carbonstack-validate now includes a manual `dev-runtime-openmls` profile.
    The profile wraps carbonstack-comms/scripts/dev-openmls-runtime-smoke.sh.
    It validates openmls-send-dev -> Cypher -> openmls-inbox-dev --ack from a live umbrella checkout.
    It is not included in `full`.

Boundary:

    This is a dev/pre-alpha manual validation profile.
    This is not local-backbone.
    This is not production messaging UX.
    This is not release-package validation yet.
    Existing send/inbox remain stub-era.

Recommended next rung:

    v0.4.10 validation/profile polish or sidecar bootstrap wrapper recon after the v0.4.9 breakpoint.

### v0.4.10 — dev-runtime-openmls profile boundary check

Result:

    The manual `dev-runtime-openmls` profile was repeatability-checked from the live umbrella checkout.
    Running without `--clean-generated` left only known OpenMLS sidecar generated roots.
    Running with `--clean-generated` cleaned known generated roots.
    A non-git package-like root was refused as expected.
    `full` remains separate and does not call `dev-runtime-openmls`.

Boundary:

    This is a documentation/evidence checkpoint.
    No runner behavior changed.
    This is not local-backbone.
    This is not release-package validation.
    This is not production messaging UX.
    Existing send/inbox remain stub-era.

Recommended next rung:

    v0.4.11 sidecar bootstrap wrapper recon.

### v0.4.11 — OpenMLS bootstrap wrapper recon

Result:

    Direct sidecar bootstrap inside scripts/dev-openmls-runtime-smoke.sh was inspected.
    Candidate dev-only wrapper names should follow the `openmls-*-dev` pattern.
    Wrapper planning should stay separate from local-backbone and Relay Space naming.
    Future wrapper code should likely live in carbonstack-comms/internal/app/openmls_bootstrap.go.
    The smoke script should not be migrated until wrapper contracts and tests exist.

Boundary:

    This is recon/planning only.
    No wrapper commands are implemented yet.
    This is not local-backbone.
    This does not replace old send/inbox.
    This does not move dev-runtime-openmls into full.
    This does not start v0.5.x state/trust/vault/PQ work.

Recommended next rung:

    v0.4.12 dev-only OpenMLS bootstrap command contract.

## v0.4.x — runtime Comms OpenMLS integration

Goal:

    make runtime Comms paths deliberately use the OpenMLS/Cypher backbone.

This means:

    send/inbox stops being stub-era
    Comms uses the OpenMLS sidecar/backbone path deliberately
    Cypher integration becomes part of actual client flow
    local developer UX becomes more coherent
    still pre-alpha/dev UX
    no production security claims

Early v0.4.x should avoid:

    public ingress
    production deployment
    Android app
    CarbonStackOS
    broad hostile-server claims
    secure vault claims

## v0.5.x — local storage, trust state, provider state, and vault design

Goal:

    solve the local state model before stronger user-readiness claims.

Topics:

    device identity state
    provider storage
    trust state
    vault boundaries
    local encryption
    backup/export policy
    compromise assumptions
    recovery assumptions
    state migration

No claim of production-safe vault until adversarial validation exists.

## v0.6.x — hostile-server and abuse-resistance harnesses

Goal:

    test server-hostile behavior deliberately.

Topics:

    replay
    rollback
    dropped messages
    delayed messages
    malformed envelopes
    metadata abuse
    server equivocation
    ack manipulation
    denial-of-service surfaces

## v0.7.x — deployability and operations hardening

Goal:

    harden the deployable/operator story after runtime and state mechanics exist.

Topics:

    operator docs
    service lifecycle
    data paths
    logs
    backup
    update/migration
    failure recovery
    server administration boundaries
    local-only vs public deployment modes

## v0.8.x — documented self-pentest / adversarial validation

Goal:

    run documented adversarial validation before making security claims.

Every record must include:

    what was tested
    methodology
    environment
    what it validates
    what it does not validate
    limitations
    follow-up fixes

This still does not replace external audit.

## v0.9.x — claim-boundary review

Goal:

    review what claims are actually allowed after accumulated validation.

Output should include:

    allowed claims
    forbidden claims
    conditional claims
    validation references
    required future audits/reviews

## v0.10.x+ — Android, app UX, and CarbonStackOS later

Android app and CarbonStackOS remain later.

Approximate direction:

    v0.10.x Android backend research + Cypher integration
    v0.11.x Android frontend / UX / CarbonStack Relay Space join model
    v0.12.x pre-major hardening
    v1.0.x deployable server + app stack major epoch
    v1.x.x CarbonStackOS / appliance OS research later

## CarbonStack Relay Space terminology

Use:

    CarbonStack Relay Space

instead of "IRC-like server" except when explaining the historical analogy.

This avoids importing IRC moderation/culture assumptions into the CarbonStack model.

## Governing principle

Every major claim must trail validation.

Do not do:

    build feature -> claim safety

Do:

    build feature -> test it -> document scope and non-scope -> then make limited claim

CarbonStack should remain evidence-led, not vibes-led.

Current package rehearsal record:

    docs/154-v0.4.0-package-rehearsal-v0.md
