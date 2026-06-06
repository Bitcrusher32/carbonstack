# CarbonStack Roadmap

This roadmap describes the current public direction of CarbonStack.

Older numbered docs preserve older plans and implementation history. Use this file, the top-level README, the latest docs index, and release-specific runbooks for current public-facing direction.

## Current state after v0.5.0

CarbonStack has completed the v0.5.0 Runtime and Registry Validation Minor Epoch Pre-Release. That is the current Gitea-source-of-truth public pre-release for package validation and release-facing state.

v0.5.0 packages accumulated v0.4.x runtime, runner, wrapper, command-registry, and package-validation work. It remains pre-alpha / experimental and is not a general-public usable release, not v1.0.0, not production secure, and not local-backbone.

Current public release:

    v0.5.0 Runtime and Registry Validation Minor Epoch Pre-Release

Current mainline checkpoint:

    v0.5.0 compressed minor-epoch baseline after the v0.5.0 release cut.
    v0.5.1 is the current post-release baseline / public-surface sanity rung.

Current mainline repo heads at the v0.5.0 release checkpoint:

    carbonstack        c6aa4e3 test: rehearse v0.5.0 package validation
    carbonstack-comms  cb4e59d test: add wrapper-based OpenMLS runtime smoke proof
    carbonstack-cypher 9ab994c docs: point to local operator runbook
    carbonstack-os     1bbbe52 docs: clarify CarbonStackOS target direction

Current public/source-of-truth policy:

    Gitea remains authoritative for releases, tags, attached release assets, and current project state.
    GitHub mirrors may exist for discoverability and redundancy, but they are push mirrors only.
    Official general-public usable releases belong later, such as the intended v1.0.0 major epoch line, unless policy changes.

## Current validated surfaces

Current release-package validation shape for v0.5.0:

    cd <package-root>/carbonstack/tools/carbonstack-validate
    go run . --profile verify-checksums --root <package-root>
    go run . --profile full --root <package-root> --clean-generated

Current live mainline development validation shape:

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go test ./... -count=1
    go run . --profile dev-runtime-openmls-wrappers --clean-generated
    go run . --profile dev-runtime-openmls --clean-generated
    go run . --profile local-cypher
    go run . --profile doctor
    go run . --profile core --clean-generated

Current runner split:

    full:
        release-package validation ladder
        release-snapshot followed by local-cypher
        not deployment
        not local-backbone
        not runtime Comms UX

    dev-runtime-openmls:
        manual live-umbrella direct-sidecar OpenMLS runtime smoke profile
        openmls-send-dev -> Cypher -> openmls-inbox-dev --ack
        not included in full

    dev-runtime-openmls-wrappers:
        manual live-umbrella wrapper-bootstrap OpenMLS runtime smoke profile
        openmls-*-dev bootstrap wrappers -> openmls-send-dev -> Cypher -> openmls-inbox-dev --ack
        separate maturity surface
        does not replace dev-runtime-openmls yet
        not included in full

Current command-surface hygiene:

    registry/commands.v0.yaml exists as a provisional cross-repo command registry.
    registry validation is covered by tools/carbonstack-validate/command_registry_test.go.
    The registry tracks runner profiles, Comms CLI commands, old stub-era send/inbox/ack, OpenMLS dev runtime commands, bootstrap wrappers, Comms smoke scripts, sidecar commands, Cypher surfaces, and legacy helper scripts.

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
    quantum-safe
    local-backbone
    mature messaging UX
    general-public usable software

Existing send/inbox remain stub-era.
The OpenMLS runtime commands are explicit dev/pre-alpha commands.
The wrapper bootstrap commands are not final Relay Space join UX.
Neither runtime profile is release-package validation yet.
Neither runtime profile is included in full.
The registry is command-surface hygiene, not a security proof.

## v0.5.x post-release baseline and next planning runway

v0.5.0 has now been cut as a minor epoch pre-release of accumulated v0.4.x runtime, runner, wrapper, registry, and validation work.

The immediate v0.5.x priority is not implementation expansion. It is:

    post-release public-surface sanity
    current-state doc cleanup
    state/trust/provider/vault inventory
    PQ/hybrid migration readiness planning
    claim-boundary preservation

It should not start immediate PQ implementation.

Compressed next rungs:

    v0.5.1 post-release preflight and public-surface baseline
    v0.5.2 state/trust/vault/PQ planning recon
    v0.5.3 storage/trust/provider-state inventory
    later: PQ/hybrid ciphersuite migration planning only after state/trust/provider boundaries are clear

### v0.4.22 — v0.5.0 checksum and fresh extraction rehearsal

Result:

    The v0.5.0 release-helper system now has:

        scripts/stage-v0.5.0-package.sh
        scripts/rehearse-v0.5.0-package.sh

    The rehearsal helper performs:

        stage package skeleton
        write release/checksums.txt
        verify checksums in staged package
        archive package
        fresh extract package
        verify checksums from fresh extraction
        run full from fresh extraction with --clean-generated

Boundary:

    No final release assets yet.
    No v0.5.0 tag yet.
    No upload.
    No local-backbone.
    No PQ/state/vault implementation.
    Runtime OpenMLS profiles remain live-umbrella-only and outside full.

Next rung:

    v0.4.23 final release notes, LogDoc sanitization/export, release asset generation, and v0.5.0 release cut prep.

### v0.4.21 — v0.5.0 package rehearsal plan and staging

Result:

    A repeatable package staging helper exists:

        scripts/stage-v0.5.0-package.sh

    It stages:

        carbonstack
        carbonstack-comms
        carbonstack-cypher
        release metadata skeleton

    It excludes:

        carbonstack-os

Boundary:

    No final checksums yet.
    No archive yet.
    No fresh extraction validation yet.
    No final release assets yet.
    No v0.5.0 tag yet.
    No local-backbone.
    No PQ/state/vault implementation.

Next rung:

    v0.4.22 checksum generation, fresh extraction validation, and release notes formulation using v0.4.0 continuity reference.

## v0.4.x completed runtime, wrapper, and registry rungs

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

### v0.4.12 — dev-only OpenMLS bootstrap command contract

Result:

    The command contract for future dev-only OpenMLS bootstrap wrappers is defined.
    Wrapper names should use the `openmls-*-dev` pattern.
    Sidecar labels remain explicit and are not derived from Comms state.
    Sidecar JSON should be normalized to stable Comms key/value output.
    Sidecar path hints should be preserved, with absolute artifact paths printed when applicable.
    Wrappers must not mutate Comms state/trust files.
    Future implementation should live in carbonstack-comms/internal/app/openmls_bootstrap.go.

Boundary:

    This is a docs/contract checkpoint.
    No wrapper commands are implemented yet.
    This is not local-backbone.
    This does not replace old send/inbox.
    This does not move dev-runtime-openmls into full.
    This does not start v0.5.x state/trust/vault/PQ work.

Recommended next rung:

    v0.4.13 implement identity create/status dev wrappers.

### v0.4.17 — wrapper OpenMLS runtime validation profile

Result:

    A separate manual runner profile exists for the wrapper-based smoke proof:

        dev-runtime-openmls-wrappers

    It wraps:

        carbonstack-comms/scripts/dev-openmls-runtime-smoke-wrappers.sh

    Proof shape:

        openmls-*-dev bootstrap wrappers -> openmls-send-dev -> Cypher -> openmls-inbox-dev --ack

Boundary:

    The original `dev-runtime-openmls` direct-smoke profile remains unchanged.
    `dev-runtime-openmls-wrappers` is separate while wrappers mature.
    Neither runtime profile is included in `full`.
    Neither profile is release-package validation yet.
    This is not local-backbone.
    This is not production/security proof.

Recommended next rung:

    pre-v0.5.x known-good command registry recon/planning.

### v0.4.18 — known-good command registry

Result:

    A provisional cross-repo command registry now exists:

        carbonstack/registry/commands.v0.yaml

    It tracks:

        runner profiles
        Comms CLI commands
        old stub-era send/inbox/ack commands
        OpenMLS dev runtime commands
        OpenMLS bootstrap wrappers
        Comms smoke scripts
        OpenMLS sidecar commands
        Cypher CLI/env/API surfaces
        legacy PowerShell helper scripts

Boundary:

    This is metadata and command-surface hygiene.
    It does not change runtime behavior.
    It is not local-backbone.
    It is not production/security proof.
    It does not add either runtime profile to full.
    It does not start v0.5.x PQ/state/vault work.

Recommended next rung:

    command registry validation / local-help planning.

### v0.4.19 — command registry validation and local help plan

Result:

    The provisional command registry is now validated by Go test coverage in:

        carbonstack/tools/carbonstack-validate/command_registry_test.go

    The test checks registry structure, duplicate IDs, required fields, source_path existence, runner profile coverage, Comms command coverage, Comms script coverage, sidecar command IDs, Cypher API IDs, legacy send/inbox/ack classification, direct-vs-wrapper OpenMLS profile separation, and local-backbone/source-of-truth boundaries.

Local help/manual planning decision:

    Use registry validation first.
    Add generated markdown command reference later.
    Avoid embedding registry-derived local help in Comms command code until command structure is ready.
    Avoid adding a runner profile for registry validation until the check proves stable.
    Keep front README minimal; registry is the complete command-surface inventory.

Boundary:

    No runtime behavior changes.
    No generated command reference yet.
    No local-backbone.
    No v0.5.x state/trust/vault/PQ implementation yet.

Carry-forward:

    After late v0.4.x grouped registry/help-planning work, refresh the long-term roadmap and prepare a v0.5.0 minor epoch release ladder with cleanup, deployability testing, release-style replication, and release hardening.

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
