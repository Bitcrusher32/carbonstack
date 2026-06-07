# CarbonStack Roadmap

This roadmap describes the current public direction of CarbonStack.

Older numbered docs preserve older plans and implementation history. Use this file, the top-level README, the latest docs index, and release-specific runbooks for current public-facing direction.

## Current state after v0.5.17

CarbonStack has completed the v0.5.0 Runtime and Registry Validation Minor Epoch Pre-Release. That remains the current Gitea-source-of-truth public pre-release for package validation and release-facing state.

v0.5.1 through v0.5.17 are post-release mainline checkpoints. They do not retag v0.5.0 and do not create a new public release package.

The major v0.5.x work after v0.5.0 has been state/trust/provider/Relay Space planning and a few narrow helper spikes:

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
    local-backbone feasibility reassessment.

Current public release:

    v0.5.0 Runtime and Registry Validation Minor Epoch Pre-Release

Current mainline checkpoint:

    v0.5.17 local-backbone feasibility reassessment and roadmap refresh.

Current mainline repo heads at the v0.5.16 checkpoint before v0.5.17 roadmap patch:

    carbonstack        9f4e982 docs: define v0.5.16 Relay Space architecture
    carbonstack-comms  94e0ff2 docs: record Comms Relay Space boundary
    carbonstack-cypher 24166b0 docs: record Relay Space boundary
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

    Validation passing does not prove production security.
    v0.5.17 does not create local-backbone.
    v0.5.17 does not create a public deployment surface.

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
    has candidate identity and mapped mismatch plans;
    does not have mature UX;
    does not have provider identity import;
    does not have candidate identity storage;
    does not have mapped mismatch implementation;
    does not have mature verification ceremony;
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
    does not implement Relay Space schema/join yet;
    is not identity authority;
    is not a trust root.

### CarbonStackOS

CarbonStackOS remains long-range constrained Android-derived appliance OS research/direction.

It is not part of current runnable validation packages and should not be treated as near-term implementation.

## Local-backbone status

Local-backbone is not ready to implement.

It is now architecture-discussable because the v0.5.x planning arc clarified state, trust, provider, candidate identity, mapped mismatch, and Relay Space boundaries.

It is not profile/release-ready because:

    Relay Space join/invite/member behavior is not implemented;
    candidate identity storage does not exist;
    mapped mismatch behavior is not implemented;
    provider event draft/append helpers are not wired into live app flows;
    Comms send/inbox remain dev/stub-era split surfaces;
    reset/recovery/re-enrollment behavior is not mature;
    validation profile claims would overstate the current system.

Future local-backbone work must remain claim-careful.

## Relay Space status

Relay Space is now defined as routing/conversation infrastructure.

Relay Space is not identity authority.

Relay admin is not a trust root.

Cypher delivery is not trust.

Server membership claims are not enough for client trust.

Client trust history and verification remain local/client-owned.

Future Relay Space work must preserve:

    hostile-server assumption;
    local trust-store authority;
    candidate identity review boundary;
    mapped mismatch/reverify boundary;
    ack-after-open boundary;
    no plaintext server authority.

## Provider/trust status

Implemented helper path:

    ProviderTrustReport;
    ProviderTrustHistoryDraft;
    ProviderTrustEventDraft;
    trust.ProviderEventAppendDraft;
    trust.BuildProviderEvent;
    trust.AppendProviderEvent.

Current boundary:

    provider-originated append helper exists;
    runtime provider events are not wired into live append flow yet;
    provider identity import does not exist;
    candidate identity storage does not exist;
    mapped mismatch mutation does not exist;
    trust.json mutation from provider observation does not exist.

Core rule:

    provider-observed identity material is not trust.

## Near-term roadmap

### v0.5.18 — roadmap-refresh aftermath / implementation-priority decision

Decide which post-planning implementation spike should happen first.

Candidate directions:

    candidate identity draft/storage spike;
    mapped mismatch classifier spike;
    provider-event live-flow append integration spike;
    Relay Space join/invite/member planning;
    validation profile boundary planning.

Default bias:

    do not create local-backbone next.
    choose one narrow implementation spike that reduces ambiguity without overclaiming.

### v0.5.19+ — targeted implementation or planning spikes

Possible sequence:

    candidate identity draft/storage model;
    mapped mismatch classifier;
    safe provider event live-flow wiring;
    Relay Space join/invite/member plan;
    local-backbone profile boundary plan;
    reset/recovery/re-enrollment plan.

Exact ordering should be chosen after v0.5.18.

## Medium-term roadmap

### Vault boundary design

Vault work should come after state/trust/provider/candidate/mismatch behavior is clearer.

Do not design vault as a generic encrypted folder.

Vault must be domain-aware:

    identity-bearing state;
    trust-bearing state;
    group/conversation-bearing state;
    message/plaintext state;
    relay-staging state;
    recovery/export state;
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
    local-backbone readiness;
    mature messenger UX;
    general-public usable software;
    Android readiness;
    CarbonStackOS readiness;
    vault security;
    PQ/hybrid security;
    quantum-safe messaging;
    external audit or certification.

v0.5.17 specifically does not claim:

    local-backbone implementation;
    Relay Space implementation;
    candidate identity implementation;
    mapped mismatch implementation;
    trust.json mutation from provider observation;
    provider identity import;
    ack behavior changes.
