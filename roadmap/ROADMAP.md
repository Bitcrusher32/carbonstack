# CarbonStack Roadmap

This roadmap describes the current public direction of CarbonStack.

Older numbered docs preserve older plans and implementation history. Use this file, the top-level README, the latest docs index, and release-specific runbooks for current public-facing direction.

## Current state after v0.5.58-REFACTOR

CarbonStack has completed the v0.5.0 Runtime and Registry Validation Minor Epoch Pre-Release. That remains the current Gitea-source-of-truth public pre-release for package validation and release-facing state.

v0.5.1 through v0.5.58 are post-release mainline development checkpoints. They do not retag v0.5.0 and do not create a new public release package.

The major v0.5.x work after v0.5.0 has included:

    state/trust/vault/PQ preliminary recon;
    storage/trust/provider-state inventory;
    storage-domain model;
    trust-state model;
    provider-state linkage planning;
    provider-trust report helpers and exposure decision;
    provider trust-history append planning/helpers;
    candidate identity storage/review/update/mismatch helpers;
    candidate observation orchestration;
    reset/recovery/re-enrollment boundary;
    recovery classifier/history/orchestration helpers;
    Relay Space architecture planning;
    Relay Space schema/API substrate;
    Relay Space scoped envelope routes;
    Comms Relay Space client and artifact bridge helpers;
    KeyPackage / Welcome / add-member / join dev command scaffolding;
    no-ack and ACK_AFTER_JOIN smoke evidence;
    relay-openmls-join-dev positive-path validation profile;
    compact profile summary;
    negative helper hardening;
    live negative-path ownership matrix;
    Comms no-ack/failure coverage matrix;
    add-member sidecar-failure command test.

Current public release:

    v0.5.0 Runtime and Registry Validation Minor Epoch Pre-Release

Current mainline checkpoint:

    v0.5.58 targeted Comms add-member sidecar-failure test, followed by v0.5.58-REFACTOR planning for pre-v0.6.0 truth hygiene.

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
    go run . --profile dev-runtime-openmls --clean-generated
    go run . --profile dev-runtime-openmls-wrappers --clean-generated
    go run . --profile relay-openmls-join-dev --compact-summary

Important:

    validation passing does not prove production security;
    v0.5.58 does not create local-backbone;
    v0.5.58 does not create a public deployment surface;
    v0.5.58 does not add relay-openmls-join-dev to full or release-snapshot;
    v0.5.58 does not complete key storage, vault, public CLI/manual, or mature inbox UX.

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
    has identity-candidates.json storage and candidate review/update helpers;
    has mapped mismatch classifier and candidate/mismatch history helpers;
    has candidate observation orchestration;
    has recovery classifier/history/orchestration helpers;
    has Relay Space client wrappers and OpenMLS artifact bridge helpers;
    has KeyPackage/Welcome/add-member/join dev commands;
    has openmls-relay-join-dev with optional --ack-after-join;
    has focused failure-path tests for join/no-ack and add-member sidecar failure;
    does not have mature UX;
    does not have verified provider identity import;
    does not have mature verification ceremony;
    does not have broad provider live-flow;
    does not have public CLI/manual polish;
    does not have secure vault/key storage;
    does not have mature public inbox semantics;
    does not have local-backbone.

### CarbonStackCypher

CarbonStackCypher is the self-hostable relay/storage server.

Current Cypher status:

    dev/pre-alpha;
    has Go/SQLite relay skeleton;
    supports dev invite creation/claim;
    supports device registration/lookup;
    supports encrypted envelope submit/retrieve/ack;
    supports Relay Space schema/API substrate;
    supports Relay Space-scoped envelope submit/inbox/ack;
    supports scoped ack rejection for wrong Relay Space and wrong recipient;
    has local operator runbook docs;
    is routing/storage infrastructure only;
    is not identity authority;
    is not a trust root;
    is not plaintext authority.

### CarbonStackOS

CarbonStackOS remains long-range constrained Android-derived appliance OS research/direction.

It is not part of current runnable validation packages and should not be treated as near-term implementation.

## Local-backbone status

Local-backbone is not implemented.

Correct current phrasing:

    local-backbone is closer;
    full local-backbone is not ready;
    current work is Relay Space / OpenMLS validation substrate and failure-path hardening;
    relay-openmls-join-dev is a positive-path local/dev profile, not local-backbone.

Local-backbone remains not profile/release-ready because:

    public CLI/manual surface is not defined;
    key storage/vault is not complete;
    inbox semantics are fragmented across dev surfaces;
    provider live-flow remains deferred;
    trust/candidate mechanics are not mature public UX;
    relay-openmls-join-dev is not included in full or release-snapshot;
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
    ack-after-open / ack-after-join boundaries;
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

## Key storage / vault status

Key storage is not complete.

Current sidecar/provider state remains dev-local generated state and must not be represented as a production vault.

Future storage/vault work must define:

    what state is source-owned;
    what state is generated dev state;
    what state is local app state;
    what state is trust state;
    what state is relay staging;
    what state is validation-only;
    what state must never silently regenerate;
    what state belongs in a future encrypted vault;
    backup/restore behavior;
    re-enrollment behavior;
    compromise/revocation behavior.

## Inbox semantics status

Inbox semantics are not complete.

Current inbox-like surfaces include:

    legacy/stub-era send/inbox;
    openmls-inbox-dev for application-message artifacts and optional ack-after-open;
    openmls-relay-keypackage-inbox-dev without ack;
    openmls-relay-welcome-inbox-dev without ack;
    openmls-relay-join-dev with optional Welcome ack after join success.

These are dev/pre-alpha surfaces, not one mature public inbox model.

Future work must define public inbox semantics before claiming mature messaging UX.

## Public CLI / manual status

Public CLI/manual surface is not complete.

The command registry exists as a navigation and claim-boundary artifact, but there is not yet a generated command reference or mature public CLI manual.

Future public/manual work must distinguish:

    public release validation commands;
    source developer commands;
    dev-only OpenMLS commands;
    hidden/private validation profiles;
    legacy/stub-era commands;
    future commands not yet implemented.

## Near-term refactor roadmap

### v0.5.59-REFACTOR-1 — surface truth / evergreen docs cleanup

Refresh current public-facing source surfaces so they do not point readers to stale v0.5.35/v0.5.48 assumptions.

Likely files:

    carbonstack/README.md
    carbonstack/docs/README.md
    carbonstack/roadmap/ROADMAP.md
    carbonstack-comms/README.md
    carbonstack-cypher/README.md
    carbonstack-comms/docs/README.md
    carbonstack-cypher/docs/README.md

### v0.5.60-REFACTOR-2 — registry/manual boundary

Define what the command registry is and is not.

Clarify which command surfaces are public release validation, developer-only, hidden/private, legacy, and future.

Do not expose relay-openmls-join-dev in front-door release docs yet.

### v0.5.61-REFACTOR-3 — hygiene / portability checkpoint

Resolve or document:

    repo-local carbonstack-cypher/cypher.db;
    generated sidecar state policy;
    Go toolchain version portability;
    local committer metadata warning;
    WSL Git Credential Manager warning;
    pre-release generated-artifact policy.

### v0.5.62-REFACTOR-4 — v0.6.0 inclusion decision

Decide whether relay-openmls-join-dev is ready for v0.6.0 release-package validation.

If yes, design inclusion deliberately.

If no, keep it live-mainline/dev-only and document why.

### Later v0.5.x / pre-v0.6.0

Remaining targeted tests may still be useful:

    add-member KeyPackage artifact write failure;
    add-member Welcome submit failure after sidecar success.

But cleanup/truth-hygiene is now the priority before more feature growth.

## Do not do next

Do not:

    call current mainline full local-backbone;
    add relay-openmls-join-dev to front README;
    add relay-openmls-join-dev to full or release-snapshot without a separate decision;
    promote dev-only OpenMLS commands as mature public CLI;
    claim production secure messaging;
    claim metadata privacy;
    claim hostile-server safety;
    claim verified identity;
    claim secure enrollment;
    claim mature provider live-flow;
    claim mature send/inbox UX;
    claim secure vault/key storage;
    mutate trust.json from Relay Space routing membership or provider/OpenMLS observation alone;
    mutate identity-candidates.json into verified trust from provider/OpenMLS observation alone;
    mark provider-observed identity as verified;
    silently replace known key material;
    implement destructive reset/recovery/re-enrollment helpers without a new explicit plan;
    delete unknown sidecar state;
    delete non-profile provider storage;
    delete non-profile Cypher DBs without an explicit hygiene decision;
    kill arbitrary user Cypher processes;
    implement vault encryption yet;
    implement PQ/hybrid ciphersuites yet;
    commit raw LogDoc or breakpoint files to Gitea.
