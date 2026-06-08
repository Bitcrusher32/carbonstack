# CarbonStack Roadmap

This roadmap describes the current public direction of CarbonStack.

Older numbered docs preserve older plans and implementation history. Use this file, the top-level README, the latest docs index, and release-specific runbooks for current public-facing direction.

## Current state after v0.5.35

CarbonStack has completed the v0.5.0 Runtime and Registry Validation Minor Epoch Pre-Release. That remains the current Gitea-source-of-truth public pre-release for package validation and release-facing state.

v0.5.1 through v0.5.35 are post-release mainline checkpoints. They do not retag v0.5.0 and do not create a new public release package.

The major v0.5.x work after v0.5.0 has been state/trust/provider/candidate/recovery/Relay Space planning and narrow helper implementation:

    state/trust/vault/PQ preliminary recon;
    storage/trust/provider-state inventory;
    storage-domain model;
    trust-state model;
    provider-state linkage plan;
    provider-trust decision report helper;
    provider-trust report contract and exposure decision;
    provider-originated trust-history append plan;
    provider trust-history draft helper;
    provider trust-event draft bridge;
    provider trust-event append helper;
    provider identity candidate / unverified import plan;
    mapped provider identity mismatch plan;
    Relay Space architecture decision record;
    local-backbone feasibility reassessment;
    candidate identity storage;
    mapped mismatch classifier;
    candidate review/update mechanics;
    candidate/mismatch history events;
    candidate observation orchestration;
    reset/recovery/re-enrollment boundary;
    pure recovery classifier;
    recovery-history append helpers;
    recovery orchestration;
    local-backbone blocker reassessment;
    Relay Space join/invite/member mechanics planning;
    provider live-flow boundary planning;
    validation profile boundary planning;
    local-backbone go/no-go reassessment;
    implementation-readiness cleanup.

Current public release:

    v0.5.0 Runtime and Registry Validation Minor Epoch Pre-Release

Current mainline checkpoint:

    v0.5.35 cleanup / implementation-readiness checkpoint.

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

    cd ~/repos/carbonstack_umbrella/carbonstack-comms
    go test ./... -count=1

    cd ~/repos/carbonstack_umbrella/carbonstack-cypher
    go test ./... -count=1

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go test ./... -count=1
    go run . --profile doctor

Additional existing profiles remain useful where relevant:

    go run . --profile core --clean-generated
    go run . --profile local-cypher
    go run . --profile dev-runtime-openmls-wrappers --clean-generated

Important:

    validation passing does not prove production security;
    v0.5.35 does not create local-backbone;
    v0.5.35 does not create a public deployment surface;
    v0.5.35 does not add a local-backbone validation profile.

## Current architecture truth

### CarbonStack

CarbonStack is the umbrella doctrine, roadmap, release, docs, registry, and validation surface.

It remains the public claim-boundary authority.

### CarbonStackComms

CarbonStackComms is the text-first messaging client and current OpenMLS sidecar integration surface.

Current Comms status:

    dev/pre-alpha;
    has local state and trust-state packages;
    has OpenMLS sidecar wrapper/runtime surfaces;
    has provider event/trust classification helpers;
    has provider-originated trust-history draft/event/append helper path;
    has identity-candidates.json storage;
    has mapped mismatch classifier;
    has candidate review/update mechanics;
    has candidate/mismatch history append helpers;
    has candidate observation orchestration;
    has recovery classifier/history/orchestration helpers;
    has Relay Space, provider live-flow, validation-profile, and local-backbone boundary docs;
    does not have mature UX;
    does not have verified provider identity import;
    does not have mature verification ceremony;
    does not have broad provider live-flow;
    does not have local-backbone.

### CarbonStackCypher

CarbonStackCypher is the self-hostable relay/storage server.

Current Cypher status:

    dev/pre-alpha;
    has Go/SQLite relay skeleton;
    supports dev invite creation/claim;
    supports device registration/lookup;
    supports encrypted envelope submit/retrieve/ack;
    has local operator runbook docs;
    has Relay Space boundary docs;
    has validation-profile and local-backbone boundary docs;
    does not implement Relay Space schema/API yet;
    does not implement Relay Space join yet;
    is not identity authority;
    is not a trust root.

### CarbonStackOS

CarbonStackOS remains long-range constrained Android-derived appliance OS research/direction.

It is not part of current runnable validation packages and should not be treated as near-term implementation.

## Local-backbone status

Local-backbone is not implemented.

v0.5.34 granted only a conditional GO for first narrow implementation planning.

Correct current phrasing:

    local-backbone is closer;
    full local-backbone is not ready;
    first narrow implementation target is Cypher Relay Space schema/API substrate;
    do not call the next implementation full local-backbone.

Local-backbone remains not profile/release-ready because:

    Relay Space schema/API implementation does not exist;
    Relay Space join/invite/member behavior is not implemented;
    Comms Relay Space client wrappers do not exist;
    provider live-flow remains deferred;
    validation profile claims would still overstate the system;
    CLI/registry exposure remains deferred;
    destructive reset/recovery/re-enrollment behavior is intentionally absent;
    mature verification ceremony, vault, PQ/hybrid migration, hostile-server harnesses, and operator-grade deployability remain later.

Future local-backbone work must remain claim-careful.

## Relay Space status

Relay Space is routing/conversation infrastructure and a vector to OpenMLS join.

Relay Space is not identity authority.

OpenMLS/provider membership is cryptographic group participation, not local verification.

Local Comms verification remains actual trust/auth/presence.

Relay admin is not a trust root.

Cypher delivery is not trust.

Server membership claims are not enough for client trust.

Client trust history, candidate review, recovery decisions, and verification remain local/client-owned.

Future Relay Space work must preserve:

    hostile-server assumption;
    local trust-store authority;
    candidate identity review boundary;
    mapped mismatch/reverify boundary;
    recovery classification boundary;
    ack-after-open boundary;
    no plaintext server authority.

## Provider/trust status

Implemented helper path:

    ProviderTrustReport;
    ProviderTrustHistoryDraft;
    ProviderTrustEventDraft;
    trust.ProviderEventAppendDraft;
    trust.BuildProviderEvent;
    trust.AppendProviderEvent;
    identity candidate storage;
    mapped mismatch classifier;
    candidate review/update;
    candidate/mismatch history events;
    candidate observation orchestration;
    recovery classifier;
    recovery-history append;
    recovery orchestration.

Current boundary:

    provider-originated append helper exists;
    runtime provider events are not broadly wired into live append flow yet;
    verified provider identity import does not exist;
    trust.json mutation from provider observation does not exist;
    Relay Space membership must not mutate local trust.

Core rule:

    provider-observed identity material is not trust.

## Near-term roadmap

### v0.5.36+ — first narrow Cypher Relay Space schema/API substrate

Before implementation:

    user must resend latest LogDoc as direct continuity reference;
    assistant must scout relevant planning docs;
    first-rung scope must be confirmed before code.

Preferred first implementation target:

    Cypher Relay Space schema/API substrate.

Possible narrow split:

    relay_spaces schema;
    relay_space_invites schema;
    relay_space_members schema;
    relay_space_id-scoped envelope support or migration path;
    create/get/list Relay Space APIs;
    invite create/claim APIs;
    routing member registration/list APIs;
    routing-only tests;
    no-trust-authority tests.

Do not combine all of these into one rung if a schema-only first rung is safer.

### Later v0.5.x after Cypher substrate

Possible later sequence:

    Comms Relay Space client wrapper;
    Comms no-trust-mutation tests;
    candidate handoff only when identity material is explicit;
    provider/OpenMLS join wiring;
    ack-boundary tests;
    validation profile implementation only after concrete substrate exists;
    CLI/registry only if the dev surface is honest and useful.

## Medium-term roadmap

### Vault boundary design

Vault work should come after state/trust/provider/candidate/mismatch/recovery/Relay Space behavior is clearer.

Do not design vault as a generic encrypted folder.

Vault must be domain-aware:

    identity-bearing state;
    trust-bearing state;
    group/conversation-bearing state;
    message/plaintext state;
    relay-staging state;
    recovery/export state;
    Relay Space routing state;
    revocation/compromise state.

### PQ/hybrid readiness

PQ/hybrid work should remain deferred until state/trust/provider migration boundaries are clearer.

PQ/hybrid is not a one-line ciphersuite flip.

It affects:

    KeyPackage size and compatibility;
    Welcome size and compatibility;
    OpenMLS group/provider state;
    protocol_version and content_type mapping;
    trust display;
    group reinit/rekey policy;
    runner artifact-size assumptions;
    release claim boundaries.

Do not claim quantum-safe messaging from one successful test.

### Hostile-server harnesses

Hostile-server and abuse-resistance validation remain later work.

Future harnesses should test:

    replay;
    stale epoch;
    server equivocation;
    malicious identity material delivery;
    delayed or hidden messages;
    malicious membership claims;
    ack/order edge cases.

### Deployability and operations hardening

Deployability/ops hardening remains later than local architecture correctness.

Future work should include:

    operator runbooks;
    migration policy;
    backup/restore policy;
    retention/pruning policy;
    server admin boundary docs;
    realistic threat-model docs.

## Long-range roadmap

Later epochs remain:

    v0.6.x hostile-server and abuse-resistance harnesses;
    v0.7.x deployability and operations hardening;
    v0.8.x documented self-pentest / adversarial validation;
    v0.9.x claim-boundary review;
    v0.10.x Android backend/app exploration;
    v1.x.x public app/server major epoch only if justified by maturity.

CarbonStackOS remains beyond those near/mid-term milestones.

## Hard nonclaims

CarbonStack currently does not claim:

    production readiness;
    production E2EE;
    hostile-server safety;
    metadata privacy;
    local-backbone implementation;
    local-backbone validation profile;
    mature messenger UX;
    general-public usable software;
    Android readiness;
    CarbonStackOS readiness;
    vault security;
    PQ/hybrid security;
    quantum-safe messaging;
    external audit or certification.

v0.5.35 specifically does not claim:

    Relay Space schema/API implementation;
    provider live-flow implementation;
    validation profile implementation;
    CLI/registry update;
    trust.json mutation from provider observation;
    provider identity verified import;
    ack behavior changes;
    destructive reset/recovery/re-enrollment support.
