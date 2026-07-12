[META NOTE: REDACTIONS REPRESENTED AS ▓▓ CHARACTERS]

# CarbonStack LogDoc v0.7.0PRIME

**Last updated:** 2026-07-12 local session  
**Owner/operator:** bitcrusher32  
**Assistant/model:** GPT-5.6 Thinking  
**Current phase:** Phase 2F closure -> **CarbonStack v0.7.0 Cumulative Pre-Alpha Engineering Boundary Pre-Release is public, its intended multi-repo package and companion assets were independently downloaded and validated, stale rehearsal-state wording in the first public attachment set was identified and corrected without changing source heads or tag history, the corrected asset set was rebuilt and revalidated from fresh extraction, the corrected assets were uploaded, direct public endpoint hashes matched, Gate 0 is closed, and the full v0.6.x OpenMLS message-flow integration / boundary-hardening ledger is formally carried forward below as the historical development record.**  
**Current public release:** `v0.7.0 Cumulative Pre-Alpha Engineering Boundary Pre-Release`  
**Current mainline checkpoint:** `v0.7.0PRIME Release / Gate 0 Closure Boundary`  
**Release URL:** https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.7.0  
**Active planning authority:** `CarbonStack_Long_Term_Roadmap_v0.8.0_EVERGREEN`  
**Target working title:** `v0.8.0 Operational Spine Maturation Pre-Release`  
**Update source:** the complete `CarbonStackLogDocV0.6.33` carried forward below; the public v0.7.0 release and remote tag; the final v0.6.34 release-candidate audit state represented by CarbonStack head `fdd99ba`; the uploaded v0.7.0 repository snapshots; the v0.7.0 Gate 0 public-download verification run; the corrected public asset build and validation run; the final direct public endpoint hash probe; the v0.7.0 mechanical-internals inspection; and the active v0.8.0 EVERGREEN planning authority.  
**Update purpose:** Preserve a true PRIME closure ledger rather than a compressed summary: retain the complete v0.6.x working chronology, decisions, blunders, validation ladders, docs, commits, and claim boundaries; add the v0.7.0 release cut, post-release verification, asset correction, public endpoint verification, current mechanical model, release-critical gaps, v0.6.x closure decision, and v0.7.x safe-resume direction; and establish a stable source document from which a separately sanitized publication copy can be derived.

**Version schema:** `v[scope].[timeline]`. This file is `v0.7.0PRIME`, the release-boundary PRIME closure ledger after the public v0.7.0 release, Gate 0 verification, corrected public asset replacement, and formal handoff from the v0.6.x engineering runway into v0.7.x operational-spine integration.

**PRIME status note:** This unsanitized file is a private development continuity ledger. It can contain local usernames, hostnames, paths, timestamps, exact operational failures, and other details that must not automatically enter a public repository or release package.

**LogDoc V2 usage note:** The Markdown LogDoc is intentionally a large human continuity ledger. It preserves the historical development path and post-release closure evidence. A lean JSON Breakpoint should capture only the current state, next action, nonclaims, and critical heads; it should not duplicate this document.

**Authority note:** Current sections in this v0.7.0PRIME header supersede older current-state metadata, public-release status, next-step ordering, and safe-resume instructions where they differ. The carried-forward v0.6.33 and earlier ledger remains authoritative historical provenance for the work performed at each checkpoint.

---

## 0. Executive Summary

CarbonStack v0.7.0 is complete as a **Cumulative Pre-Alpha Engineering Boundary Pre-Release**, and Gate 0 is closed.

This PRIME is not a replacement summary for the v0.6.x development ledger. It is the next ledger layer. The entire v0.6.33 document, including its carried-forward v0.6.32 and earlier history, is retained below under **Appendix A**.

The release closes the v0.6.x engineering runway around:

1. normal OpenMLS application-message path convergence;
2. opinionated `message-send-dev` and `message-inbox-dev` wrapper surfaces;
3. lower-level direct `openmls-*` proof surfaces;
4. explicit legacy stub demotion;
5. Relay onboarding artifact separation;
6. Relay Space schema/API/scoped-envelope substrate;
7. deterministic integrated and same-state validation evidence;
8. deterministic normal-message failure classification;
9. Welcome join partial-state safety;
10. sender-metadata and provider-output claim tightening;
11. state/security/no-silent modeling;
12. bounded non-mutating `state-audit-dev`;
13. PQ/hybrid placement modeling without implementation;
14. release/runtime validation naming separation;
15. explicit `full-validate-release`;
16. adversarial-harness contract and evidence-matrix design;
17. command registry and generated-reference discipline;
18. public/manual operator-surface cleanup;
19. v0.7.0-specific package staging and rehearsal;
20. release-candidate package-surface audit;
21. public release cut;
22. independent public-download verification;
23. correction of stale rehearsal-state wording in the first public attachment set;
24. fresh-extraction validation of the corrected attachment set;
25. corrected public endpoint hash verification.

Current frozen release heads:

    carbonstack
      fdd99ba00b83ec31b3c98c18650d414527effde8
      docs: audit v0.7.0 release candidate surface

    carbonstack-comms
      5a61646439df085d2d648520fba26aef23624012
      fix: align state audit output with boundary model

    carbonstack-cypher
      d18a564044b7a4dcffe5c906f7cda0a8d016f65e
      chore: restore validated Go module floor

    carbonstack-os
      1bbbe52020d623b81796694e5057c1d080ede3ea
      docs: clarify CarbonStackOS target direction
      excluded from runnable package

Current command-reference count:

    entries=89

Current public package includes:

    carbonstack
    carbonstack-comms
    carbonstack-cypher
    release metadata

Current public package excludes:

    carbonstack-os

Primary public validation:

    cd <package-root>/carbonstack/tools/carbonstack-validate
    go run . --profile verify-checksums --root <package-root>
    go run . --profile full-validate-release --root <package-root> --clean-generated

Compatibility validation:

    go run . --profile full --root <package-root> --clean-generated

Current profile semantics:

    full-validate-release = release-snapshot + local-cypher
    full = release-snapshot + local-cypher

`release-snapshot` already invokes `core`, so neither aggregate invokes `core` twice.

CarbonStack v0.7.0 remains pre-alpha and experimental. It is not a finished messenger, production E2EE system, hostile-server-safe service, authenticated deployment, verified-identity system, vault, Android client, or CarbonStackOS appliance.

---

## 1. v0.6.34 Release-Candidate Audit and v0.7.0 Release Cut

The carried-forward v0.6.33 ledger ended with a successful v0.7.0-style package rehearsal and identified the next required work as an adversarial release-candidate/package-surface audit before tag and upload.

That final pre-release lane completed at CarbonStack head:

    fdd99ba00b83ec31b3c98c18650d414527effde8
    docs: audit v0.7.0 release candidate surface

This moved CarbonStack beyond the v0.6.33 rehearsal head:

    0f9b8d7 release: add v0.7.0 package rehearsal scripts

The release audit preserved:

- the frozen Comms, Cypher, and OS heads;
- the 89-entry generated command reference;
- the package shape;
- `full-validate-release` as the preferred release-facing profile;
- `full` as compatibility wording;
- CarbonStackOS exclusion;
- the multi-repo attachment warning;
- strict nonclaims.

The v0.7.0 tag was created remotely at:

    fdd99ba00b83ec31b3c98c18650d414527effde8

The release was published as:

    CarbonStack v0.7.0
    Cumulative Pre-Alpha Engineering Boundary Pre-Release

The first release attachment set correctly froze the source heads and passed the intended package validation ladder, but retained wording generated for the rehearsal phase. That mismatch was not discovered until the independent post-release Gate 0 verification.

---

## 2. Gate 0 Public Release Verification

### 2.1 Purpose

Gate 0 was the mandatory closure bridge between:

    v0.6.x engineering runway
        ->
    v0.7.0 public release boundary
        ->
    v0.7.x operational-spine integration

Its purpose was to prevent new implementation from beginning on top of an unverified public release or an ambiguous continuity state.

Required checks:

- local repository identity and tracked cleanliness;
- remote tag identity;
- public attachment download;
- outer asset checksum verification;
- public manifest and release-text review;
- safe fresh extraction;
- internal package checksum verification;
- external/internal release metadata comparison;
- exact packaged-source comparison against local release heads;
- `verify-checksums`;
- `full-validate-release --clean-generated`;
- final forbidden-artifact scan;
- correction of any release-integrity defect;
- public endpoint verification after correction;
- PRIME and Breakpoint generation;
- formal v0.6.x closure.

### 2.2 Verification environment

Gate 0 ran in the current WSL working environment:

    local user: Bitcrusher32
    host: ▓▓
    platform: WSL2 / Debian-family linux/amd64

Observed toolchains:

    Python 3.13.5
    Git 2.47.3
    Go 1.24.4
    Rust 1.96.0
    Cargo 1.96.0
    SQLite 3.46.1

Working umbrella:

    /home/Bitcrusher32/repos/carbonstack_umbrella

Initial evidence root:

    /home/Bitcrusher32/carbonstack-gate0-evidence/v0.7.0-gate0-20260712-011026

Corrected asset build root:

    /home/Bitcrusher32/carbonstack-gate0-evidence/v0.7.0-asset-correction-20260712-011719

Windows export path used for corrected attachment upload:

    C:\Users\▓▓\Downloads\CarbonStack-v0.7.0-corrected-assets

These details are operationally useful in the private PRIME but are redacted in the sanitized copy.

### 2.3 Local repository identity

The Gate 0 preflight confirmed:

    carbonstack
      HEAD fdd99ba00b83ec31b3c98c18650d414527effde8
      tracked clean
      no non-ignored untracked files

    carbonstack-comms
      HEAD 5a61646439df085d2d648520fba26aef23624012
      tracked clean
      expected ignored OpenMLS state and target roots only

    carbonstack-cypher
      HEAD d18a564044b7a4dcffe5c906f7cda0a8d016f65e
      tracked clean
      expected ignored cypher.db only

    carbonstack-os
      HEAD 1bbbe52020d623b81796694e5057c1d080ede3ea
      tracked clean

The local CarbonStack clone did not initially resolve `v0.7.0` because the tag had not been fetched locally.

Remote tag query confirmed:

    refs/tags/v0.7.0
    fdd99ba00b83ec31b3c98c18650d414527effde8

This was a local Git metadata synchronization issue, not a release-integrity failure.

### 2.4 Initial public attachment download

The first public attachment set downloaded successfully.

Initial outer asset checksum verification passed for all seven covered assets.

Initial public package archive:

    carbonstack-v0.7.0-cumulative-pre-alpha-engineering-boundary-pre-release.tgz

Initial archive SHA-256:

    264a17b0aeb230c85c89672e215df78144db12a0673e7ec579b7f062c1a552dd

The package safely extracted into the expected shape:

    carbonstack-v0.7.0-package/
      carbonstack/
      carbonstack-comms/
      carbonstack-cypher/
      release/

CarbonStackOS was absent as intended.

### 2.5 Exact source comparison

The packaged source snapshots exactly matched local tracked files at the frozen release heads:

    carbonstack
      tracked files: 283
      missing: 0
      different: 0
      extra: 0

    carbonstack-comms
      tracked files: 172
      missing: 0
      different: 0
      extra: 0

    carbonstack-cypher
      tracked files: 32
      missing: 0
      different: 0
      extra: 0

This confirmed that the intended package carried the expected source snapshots even though it did not include `.git` metadata.

### 2.6 Package validation

The project's own checksum verifier passed:

    checksum verification passed: 492 file(s)

`verify-checksums` result:

    VALIDATION PASSED
    return code 0

`full-validate-release --clean-generated` result:

    VALIDATION PASSED
    return code 0

The release ladder validated:

- package layout;
- release metadata presence;
- registry and generated-reference presence;
- strict pre-test artifact hygiene;
- package checksums;
- runner doctor;
- OpenMLS real-Cypher lifecycle;
- CarbonStackComms tests;
- CarbonStackCypher tests;
- local Cypher lifecycle;
- local Cypher restart against the same temporary DB;
- expected cleanup of generated OpenMLS sidecar roots;
- final artifact hygiene.

### 2.7 Supplemental verifier defect: checksum paths containing spaces

The first supplemental Python verifier parsed checksum lines using a filename pattern that stopped at spaces.

This corrupted valid paths under:

    carbonstack/sanitized-project-logdoc-list/

into repeated apparent directory entries.

False failures included:

- one missing checksum target;
- several duplicate checksum entries.

The project's own checksum implementation correctly verified all 492 files, including names containing spaces.

The corrected asset builder used an exact checksum format and exact-path parser.

Lesson:

> Supplemental verification code must not assume release filenames contain no spaces. When the project verifier and byte-identical checksum file both pass, a weaker supplemental parser must not overrule them without evidence of an actual hash mismatch.

---

## 3. Public Release-Surface Defect

### 3.1 Genuine mismatch

The first public manifest still described:

    package_role: staged rehearsal package

and:

    public_release_status:
      pre-release candidate rehearsal;
      not tagged unless explicitly released

The first public release notes still said:

    Status: staged rehearsal notes.
    These notes are package-rehearsal material
    until an explicit public release is cut.

They also referred to:

    candidate framing should cover

and:

    if this rehearsal becomes a public release

The first validation freeze said:

    This staged package was assembled

These statements were accurate during rehearsal but became false once the release was publicly cut.

### 3.2 Classification

This was:

    a public release metadata integrity defect

It was not:

- a source mismatch;
- a checksum failure;
- a package-layout failure;
- a code/test failure;
- a tag mismatch;
- a component-head mismatch;
- a security regression.

The correct fix was to replace the public attachment set around the same frozen source heads.

Retagging, rewriting commit history, or changing component source was neither required nor justified.

---

## 4. Corrected Public Asset Build

### 4.1 Correction scope

The correction:

- consumed the already downloaded and technically validated public package;
- verified the expected frozen heads;
- rewrote only release metadata and companion text;
- regenerated internal package checksums;
- rebuilt the package archive;
- copied finalized companion assets;
- regenerated outer asset checksums;
- fresh-extracted the corrected archive;
- reverified internal checksums;
- reran project validation;
- rescanned final artifacts.

No repository was modified.

No release tag changed.

No source head changed.

### 4.2 Corrected metadata posture

The corrected manifest now describes the package as:

    Gitea source-of-truth public cumulative pre-alpha
    engineering boundary pre-release package

The corrected release status describes it as:

    attached intended multi-repo validation package
    for the public CarbonStack v0.7.0 pre-release

The corrected manifest also records:

- the reason for correction;
- the fact that correction scope was release metadata and companion text;
- the fact that included source heads remained unchanged.

### 4.3 Corrected release notes

The corrected notes now:

- identify the release as public;
- describe the cumulative v0.6.x engineering boundary;
- preserve `full-validate-release` semantics;
- list exact included and excluded heads;
- describe the actual v0.7.0 boundary;
- preserve hard nonclaims;
- warn against Gitea-generated single-repo source archives.

### 4.4 Corrected validation run

Corrected asset builder result:

    ASSET CORRECTION RESULT: PASSED

Internal checksum entries verified:

    492

Outer checksum-covered assets verified:

    7

Fresh corrected package validation:

    verify-checksums passed
    full-validate-release passed
    final artifact scan passed
    exit code 0

### 4.5 Final corrected public asset hashes

    LICENSE
      bytes: 1077
      sha256: 0d4c1ef6bcb6bb2cee1f1ba9356411a7548b1716035ee390044110f2eee82a01

    carbonstack-v0.7.0-asset-checksums.txt
      bytes: 719
      sha256: 0d02c5304ccba084edcd4487b8ceb26dca6825fb3423215ee292788937e45d04

    carbonstack-v0.7.0-cumulative-pre-alpha-engineering-boundary-pre-release.tgz
      bytes: 1356396
      sha256: da02c66752a78ae13eb5ec668dadafaf320863f209c909f580e6e71ef3f46aff

    carbonstack-v0.7.0-package-checksums.txt
      bytes: 63299
      sha256: 6f207329f181bb0ad451558cba2a63c29b89ebf3efae1e64912dc9ef68351425

    carbonstack-v0.7.0-release-manifest.json
      bytes: 2624
      sha256: 2bab1adfdb1999d2e0de1395af7ab6e0e2343b93c39abdb3afd046ccdc7eb7c4

    carbonstack-v0.7.0-validation-freeze.md
      bytes: 1196
      sha256: 3cdf22afcc0e72c2f43271b4308e53e84f80846f7af13f20b8e40ebc51e21bb6

    v0.7.0-release-notes.md
      bytes: 3284
      sha256: 45b628a473728f5449c834cdffd1f0601920f954f9dd62c8e26adf8bd257ef3b

    v0.7.0-testing-runbook.md
      bytes: 1780
      sha256: 5ad7fe5df55f7767bc8ccc42cd3962c3d545de59c66a0056d9e74c46cb0a81e4

---

## 5. Corrected Public Upload and Endpoint Verification

The corrected eight assets were uploaded to the existing v0.7.0 release.

The tag and source heads were left unchanged.

The Gitea release UI showed:

- corrected file sizes;
- recent attachment timestamps;
- the same eight intended asset names;
- the release still marked as a pre-release.

Corrected UI size indicators included:

    release manifest: 2.6 KiB
    validation freeze: 1.2 KiB
    release notes: 3.2 KiB
    testing runbook: 1.7 KiB

A direct WSL public endpoint probe returned exact corrected hashes for:

    carbonstack-v0.7.0-asset-checksums.txt
      0d02c5304ccba084edcd4487b8ceb26dca6825fb3423215ee292788937e45d04

    carbonstack-v0.7.0-release-manifest.json
      2bab1adfdb1999d2e0de1395af7ab6e0e2343b93c39abdb3afd046ccdc7eb7c4

    v0.7.0-release-notes.md
      45b628a473728f5449c834cdffd1f0601920f954f9dd62c8e26adf8bd257ef3b

A separate retrieval path temporarily continued to return stale old bodies and old attachment sizes.

This was resolved as:

    cache-path divergence

rather than:

    failed upload

Authority for closure:

- Gitea UI showed corrected sizes and fresh upload times;
- direct public endpoint hash probe matched corrected objects;
- corrected asset build had already passed fresh-extraction validation;
- source heads and tag were unchanged.

---

## 6. Gate 0 Closure Decision

Gate 0 status:

    CLOSED

Technical package verification:

    PASSED

Public release metadata correction:

    PASSED

Corrected public endpoint probe:

    PASSED

Source/tag identity:

    PASSED

v0.6.x closure:

    FORMALLY CLOSED ENOUGH FOR ARCHIVAL CONTINUITY

No v0.7.x tracked-source implementation occurred during Gate 0.

The next work belongs to the v0.7.x operational-spine integration epoch.

---

## 7. Current Mechanical Model at v0.7.0

A deep static inspection of the four uploaded repository snapshots was performed before Gate 0 implementation work.

The inspection was code-grounded across:

- command dispatch;
- state helpers;
- trust/candidate/recovery helpers;
- Cypher client;
- Relay bridge;
- OpenMLS sidecar;
- Cypher DB and HTTP routes;
- validation profiles;
- package scripts;
- registry and generated references;
- tests;
- public/operator documentation.

### 7.1 Repository roles

#### carbonstack

Primary responsibilities:

- doctrine;
- threat model;
- model/decision docs;
- command registry;
- generated command reference;
- command-boundary table;
- validation runner;
- release/package scripts;
- release metadata;
- roadmap;
- public claim authority.

It is not the main runtime.

#### carbonstack-comms

Primary responsibilities:

- operator CLI;
- local application state;
- state paths and state audit;
- local trust state;
- candidate and recovery model surfaces;
- Cypher client;
- Relay/OpenMLS orchestration;
- Go-to-Rust sidecar invocation;
- OpenMLS identity/group/message lifecycle;
- recommended message wrappers;
- lower-level proof commands.

It contains most runtime complexity.

#### carbonstack-cypher

Primary responsibilities:

- account/device records;
- Relay Space records;
- opaque envelope queue;
- payload metadata;
- delivery state;
- acknowledgement state;
- SQLite migrations;
- HTTP/API surface.

It deliberately does not parse MLS plaintext or become cryptographic group authority.

#### carbonstack-os

Current state:

- doctrine;
- North Star;
- future appliance constraints;
- target-device direction.

It does not contain a runnable CarbonStackOS build.

### 7.2 Actual authority topology

    CarbonStackComms
      operator intent
      local policy
      orchestration

    OpenMLS sidecar
      signer
      KeyPackage
      Welcome
      group state
      message protection/opening
      cryptographic conversation authority

    CarbonStackCypher
      account/device coordination
      Relay Space records
      opaque envelope routing
      delivery and acknowledgement authority

These authorities must remain explicit.

### 7.3 Identity domains

Three identity domains currently exist:

1. Cypher account/device registration identity material.
2. Comms local trust/candidate identity records.
3. OpenMLS signer and credential identity.

They are correlated operationally through labels and IDs but are not cryptographically bound.

A locally `verified` Cypher trust record does not prove that the corresponding OpenMLS signer protected a message.

This is a release-critical Gate A/C design issue, not a v0.7.0 release claim.

### 7.4 Normal message flow

Recommended dev/pre-alpha send:

    message-send-dev
      -> load Comms state
      -> evaluate local recipient trust
      -> invoke sidecar message-protect
      -> read generated MLS artifact
      -> submit artifact to Cypher unscoped envelope API

Recommended dev/pre-alpha inbox:

    message-inbox-dev
      -> query unscoped device inbox
      -> select OpenMLS application-message artifact
      -> write temporary artifact
      -> invoke sidecar message-open
      -> print plaintext
      -> optionally ACK after successful open

Current positive boundaries:

- ACK can be delayed until successful open;
- unsupported content types are classified;
- sender device metadata is labelled unverified;
- revoked/compromised recipients block;
- strict mode can block unknown/unverified/changed recipients;
- direct OpenMLS commands remain lower-level proof surfaces;
- legacy stub commands require explicit opt-in.

Current limitations:

- wrapper orchestration is duplicated;
- normal messages are not Relay Space scoped;
- plaintext is displayed through a dev CLI;
- temporary artifact cleanup is not yet a mature lifecycle;
- identity binding remains absent.

### 7.5 Relay onboarding flow

Current development proof:

    create identity
      -> export KeyPackage
      -> submit/retrieve scoped KeyPackage
      -> add member
      -> generate Welcome
      -> submit/retrieve scoped Welcome
      -> staged join
      -> ACK after successful join

Welcome join has meaningful partial-state safety.

Current missing lifecycle semantics include:

- complete artifact identity;
- freshness;
- duplicate policy;
- invite claim;
- membership enforcement;
- recovery-required states;
- Cypher/MLS membership mismatch policy.

---

## 8. Release-Critical Mechanical Gaps

### 8.1 Identity-domain binding

Cypher device identity, Comms trust state, and OpenMLS signer identity are separate.

The v0.7.x workflow contract must either:

- define explicit binding/enrollment; or
- preserve explicit unbound classification and refusal/display behavior.

### 8.2 Unauthenticated Cypher control/data plane

Current Cypher HTTP operations trust caller-supplied account, device, Relay Space, sender, and recipient IDs.

This is acceptable only as an explicitly local dev substrate.

It blocks:

- public ingress;
- hostile-network deployment claims;
- authenticated-device claims;
- secure enrollment claims.

### 8.3 Relay membership not enforced on scoped delivery

Relay Space membership is stored but is not currently required for scoped envelope submission.

Membership is coordination state, not routing authorization.

### 8.4 Unscoped inbox can cross Relay scope

The ordinary device inbox query filters recipient and delivery state but does not require `relay_space_id IS NULL`.

A Relay-scoped envelope can appear through the unscoped inbox.

This is a concrete Gate B boundary defect.

### 8.5 Normal messages remain unscoped

Relay onboarding is scoped.

Normal OpenMLS application messages use the unscoped envelope route.

The preferred operational model must decide whether:

- every normal conversation belongs to a Relay Space;
- unscoped direct messaging remains a separate supported mode; or
- unscoped messaging becomes legacy/dev-only.

### 8.6 Workflow duplication

`message-send-dev` / `message-inbox-dev` and lower-level `openmls-*` commands duplicate orchestration.

A reusable noninteractive workflow engine does not yet own:

- operator context;
- state reads/writes;
- mutation boundaries;
- lifecycle transitions;
- refusal semantics.

### 8.7 State atomicity and concurrency

Comms JSON, trust history, and OpenMLS provider/group state rely heavily on direct writes.

Current risks include:

- lost update;
- store/history disagreement;
- partial output;
- concurrent command collision;
- crash consistency ambiguity.

Welcome join is better staged than some other stateful operations, especially add-member.

### 8.8 Development-oriented state paths

OpenMLS state remains source-tree-relative.

Operations invoke the sidecar through Cargo.

This is development plumbing, not an installed runtime layout.

### 8.9 Development-only Cypher defaults

Current defaults include broad bind behavior, relative database paths, and development invite behavior.

A future native deployment profile needs explicit safe private defaults.

### 8.10 Registry incompleteness for Relay routes

The deterministic registry contains the original Cypher API surfaces but does not classify all newer Relay Space HTTP routes.

The registry should be updated only after Gate B decides route status and lifecycle meaning.

---

## 9. Validation Boundary at v0.7.0

Release/package validation:

    doctor
    core
    local-cypher
    release-snapshot
    full
    full-validate-release
    verify-checksums

Dev/runtime validation:

    dev-runtime-openmls
    dev-runtime-openmls-wrappers
    integrated-runtime-dev
    relay-openmls-join-dev

Same-state and failure evidence:

    same-state-integrated-dev
    same-state-message-failure-dev
    same-state-message-unsupported-dev
    same-state-message-recipient-failure-dev
    same-state-message-malformed-payload-dev
    same-state-message-replay-classification-dev
    same-state-welcome-join-failure-dev

State classification:

    state-audit-dev

Important:

    full-runtime-dev does not exist

Important:

    no adversarial aggregate profile exists

Release validation, runtime validation, state inspection, same-state evidence, failure evidence, and future adversarial evidence remain separate.

---

## 10. Adversarial Testing Boundary

Negative-path engineering continues during v0.7.x.

Examples:

- wrong-state refusal;
- malformed input;
- unsupported-state handling;
- stale/duplicate onboarding artifacts;
- membership mismatch;
- partial failure;
- state compatibility refusal;
- no-silent mutation;
- concurrent operation;
- restart/resume;
- rollback/stale-state observability.

This is not the same as the full adversarial campaign.

The full campaign is planned for post-v0.9.0 during v0.9.x Debian deployment work and may include:

- extreme cases;
- broad edge-case campaigns;
- guided penetration attempts;
- adversarial code review;
- deployment/network/state manipulation;
- malicious-relay scenarios;
- hostile-environment exercises;
- operator misuse and recovery attacks.

v0.7.x must not claim that campaign has occurred.

---

## 11. Active Roadmap and Gate Sequence

Active planning authority:

    CarbonStack_Long_Term_Roadmap_v0.8.0_EVERGREEN

Active epoch:

    v0.7.x operational-spine integration

Target:

    v0.8.0 Operational Spine Maturation Pre-Release

### Gate A — Workflow contract

Define:

- preferred operator lifecycle;
- command ownership;
- explicit operator context;
- capability domains;
- state reads/writes;
- mutation boundaries;
- refusal semantics;
- direct versus Relay-scoped message posture;
- identity-domain posture.

### Gate B — Relay lifecycle

Define and enforce:

- Relay addressability;
- membership semantics;
- invite claiming;
- KeyPackage identity and freshness;
- Welcome identity and freshness;
- duplicate behavior;
- normal-message scoping;
- acknowledgement discipline;
- Cypher-membership versus MLS-membership mismatch handling.

### Gate C — State enforcement

Define and implement:

- state authority;
- sensitivity;
- schema;
- boundary versions;
- canonical paths;
- locking;
- atomic replacement;
- compatibility checks;
- unsupported/newer-state refusal;
- no-silent migration;
- no-silent repair;
- no-silent trust promotion;
- restart/resume semantics.

### Gate D — Runtime aggregate

Only after stable leaf workflows exist, decide whether one coherent `full-runtime-dev`-like lifecycle is justified.

Do not aggregate every dev profile.

### Gate E — Native private Debian deployment

Establish one repeatable private native model with explicit:

- bind behavior;
- config;
- state;
- logs;
- ownership;
- startup/shutdown;
- restart/resume;
- diagnostics;
- failure behavior.

This does not imply public ingress.

### Gate F — v0.8.0 release maturity

Require enough:

- workflow coherence;
- Relay lifecycle enforcement;
- state compatibility/refusal;
- restart/resume evidence;
- coherent runtime validation;
- native private operation;
- operator documentation;
- package/runtime rehearsal;
- bounded public claims;

to justify the target pre-release.

---

## 12. Immediate Next Work

The next checkpoint is:

    no-tracked-source-mutation mechanical authority inventory

It should reconcile:

- all Comms commands;
- all Cypher HTTP routes;
- all validation profiles;
- all state read/write sets;
- all identity domains;
- direct versus Relay-scoped paths;
- mutation points;
- refusal behavior;
- tests;
- registry coverage;
- operator documentation;
- current release/public docs.

After the inventory:

    conduct structured Gate A design Q/A

Then:

    write and accept the workflow/context contract

Only then:

    select one narrow implementation leaf

Do not begin with:

- TUI;
- `full-runtime-dev`;
- containers;
- public ingress;
- vault implementation;
- backup/restore implementation;
- PQ migration;
- Android;
- CarbonStackOS.

---

## 13. Hard Nonclaims

CarbonStack v0.7.0 and this PRIME do not prove:

- production readiness;
- production secure messaging;
- production E2EE;
- hostile-server safety;
- malicious-relay safety;
- metadata privacy;
- metadata authenticity;
- sender authenticity;
- verified identity;
- secure enrollment;
- authenticated Cypher APIs;
- secure vault or key storage;
- encryption at rest;
- production backup or restore;
- rollback safety;
- broad replay safety;
- complete KeyPackage freshness handling;
- complete Welcome freshness handling;
- PQ or hybrid security;
- quantum-safe messaging;
- native deployment readiness;
- systemd readiness;
- cloudflared readiness;
- public ingress safety;
- real homelab validation;
- mature messenger UX;
- Android readiness;
- CarbonStackOS readiness;
- external audit;
- certification;
- general-public usability.

---

## 14. Current Safe Resume Point

Current public release:

    v0.7.0 Cumulative Pre-Alpha Engineering Boundary Pre-Release

Current frozen heads:

    carbonstack        fdd99ba00b83ec31b3c98c18650d414527effde8
    carbonstack-comms  5a61646439df085d2d648520fba26aef23624012
    carbonstack-cypher d18a564044b7a4dcffe5c906f7cda0a8d016f65e
    carbonstack-os     1bbbe52020d623b81796694e5057c1d080ede3ea

Command reference:

    entries=89

Gate 0:

    CLOSED

v0.6.x:

    formally closed enough for archival continuity

Active epoch:

    v0.7.x operational-spine integration

Next safest action:

    perform a no-tracked-source-mutation mechanical authority inventory
    then define and accept the Gate A workflow/context contract

Safe to pause:

    yes

---

## 15. Canonical Continuity Sentence

CarbonStack v0.7.0PRIME records the public v0.7.0 Cumulative Pre-Alpha Engineering Boundary Pre-Release at CarbonStack commit `fdd99ba00b83ec31b3c98c18650d414527effde8`, packaging CarbonStackComms at `5a61646439df085d2d648520fba26aef23624012` and CarbonStackCypher at `d18a564044b7a4dcffe5c906f7cda0a8d016f65e` while excluding CarbonStackOS at `1bbbe52020d623b81796694e5057c1d080ede3ea`; it carries forward the complete v0.6.x working ledger, records the v0.7.0 release-candidate audit and release cut, independently validates the public multi-repo package, distinguishes verifier defects from a real stale-rehearsal-metadata defect, corrects the public attachment set without changing source heads or tag history, revalidates 492 package checksums plus `verify-checksums`, `full-validate-release`, OpenMLS real-Cypher lifecycle, Comms/Cypher tests, local Cypher restart behavior and artifact hygiene, closes Gate 0 and the v0.6.x runway, and establishes v0.7.x mechanical authority inventory plus Gate A workflow-contract work as the next safe path toward the v0.8.0 Operational Spine Maturation Pre-Release while all production-security, hostile-server, verified-identity, vault, deployment, PQ, Android, CarbonStackOS, audit and certification claims remain explicitly out of scope.

---

# Appendix A — Complete Carried-Forward CarbonStack LogDoc v0.6.33

The complete prior development ledger follows verbatim.

Current v0.7.0PRIME sections above supersede only current-state metadata, current public-release status, current repo-head summary where later commits exist, next-step ordering, Gate 0 status, and safe-resume instructions where they differ.

All v0.6.33 and earlier chronology, implementation notes, decisions, validation evidence, blunders, historical heads, docs, and claim boundaries remain preserved as development provenance.

# CarbonStack LogDoc v0.6.33

**Last updated:** 2026-07-02 local session  
**Owner/operator:** bitcrusher32  
**Assistant/model:** GPT-5.5 Thinking  
**Current phase:** Phase 2F -> **v0.6.33 release-boundary / package-root rehearsal checkpoint complete; v0.6.33A completed release-boundary/package-root preflight; v0.6.33B committed the release-boundary rehearsal plan doc; v0.6.33C committed v0.7.0-specific staging/rehearsal scripts, registry rows, generated command reference refresh, and runner wording cleanup; v0.6.33D successfully rehearsed the v0.7.0 cumulative pre-alpha engineering boundary package from a fresh extraction, including asset checksum verification, package checksum verification, `full-validate-release`, cleanup, and final artifact scan.**  
**Current public release:** `v0.6.0 State/UX Boundary Validation Pre-Release`  
**Current mainline checkpoint:** `v0.6.33 Release-Boundary / Package-Root Rehearsal Checkpoint`  
**Release URL:** https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.6.0  
**Update source:** `CarbonStackLogDocV0.6.32`, `CarbonStack_BreakpointV0.6.32`, v0.6.33A release-boundary/package-root preflight, v0.6.33B release-boundary plan-doc checkpoint log, v0.6.33C v0.7.0 script/registry/runner repair and commit log, v0.6.33D fresh package rehearsal log, the v0.6.0 public release shape, and the active `CarbonStack_Long_Term_Roadmap_v0.6.26_EVERGREEN` planning authority.  
**Update purpose:** Preserve the v0.6.33 checkpoint after release-boundary/package-root planning and v0.7.0-style rehearsal work; record the v0.6.33A/B/C/D timeline, v0.7.0 scripts/registry/reference changes, rehearsal evidence, staged assets, blunders and continuity lessons, current repo heads, critical path/function updates, future release-candidate audit ordering, and nonclaims; preserve the full prior Markdown ledger; and keep the JSON breakpoint as a lean current-state handoff.

**Version schema:** `v[scope].[timeline]`. This file is `v0.6.33`, the working branch continuity ledger after v0.7.0 release-boundary/package-root rehearsal planning, package-helper implementation, and successful fresh-extraction package rehearsal in the v0.6.x OpenMLS message-flow integration and boundary-hardening epoch.

**LogDoc V2 usage note:** The Markdown LogDoc is intentionally a dev continuity ledger. It preserves historical v0.6.x timeline data, blunders, validation ladders, command-boundary decisions, and future path constraints. The JSON breakpoint remains a lean current-state handoff and should not try to duplicate the full Markdown timeline.

**Privacy / history-rewrite note:** Exact raw shell logs may contain local usernames, hostnames, email addresses, and old pre-rewrite author identities. This LogDoc avoids repeating local host-derived identity strings except where operationally necessary. Public commit identity remains normalized on public `main`/tag histories. Pre-v0.6.13 commit hashes from older LogDocs and historical docs are now historical **pre-rewrite** references and must not be treated as current Gitea `main`/tag SHAs.

## 0. Executive Summary

v0.6.33 is complete as a **release-boundary / package-root rehearsal checkpoint**.

It did eight bounded things:

1. **Ran v0.6.33A release-boundary/package-root preflight**
   - v0.6.33A was a no-mutation preflight over:
     - current repo heads and cleanliness;
     - generated command reference;
     - runner tests;
     - registry missing-nonclaims;
     - runner doctor;
     - required v0.6.27-v0.6.32 docs;
     - Comms/Cypher soft tests;
     - OpenMLS sidecar `cargo test`;
     - runner help and release-validation profile discoverability;
     - release-supported registry surfaces;
     - release/package-related docs and script inventory;
     - release-snapshot / checksum runner code;
     - package-root shape and source hygiene expectations;
     - toolchain-floor release-prep wording.
   - It confirmed the v0.6.32 baseline:
     - `carbonstack` at `f9d4ac4`;
     - `carbonstack-comms` at `5a61646`;
     - `carbonstack-cypher` at `d18a564`;
     - `carbonstack-os` at `1bbbe52`;
     - `COMMAND_REFERENCE.v0.md` current with `entries=87`;
     - runner tests pass;
     - registry missing-nonclaims scan passes;
     - runner doctor passes.
   - It found the existing v0.6.0 release scripts and release shape were strong prior art, not direct v0.7.0 authority.
   - It did not mutate repos, scripts, registry, generated reference, or release artifacts.

2. **Translated the v0.6.0 release shape into a v0.7.0 planning model**
   - The v0.6.0 public release shape established the desired release surface pattern:
     - intended multi-repo validation package archive;
     - release manifest;
     - package checksums;
     - asset checksums;
     - validation freeze;
     - testing runbook;
     - release notes;
     - LICENSE;
     - explicit warning that Gitea default Source Code ZIP/TAR.GZ archives are not the intended multi-repo validation package.
   - v0.6.33 preserved that release-shape pattern but updated the boundary:
     - v0.7.0 candidate title: `Cumulative Pre-Alpha Engineering Boundary Pre-Release`;
     - primary validation wording: `full-validate-release`;
     - `full` preserved as compatibility wording;
     - `carbonstack-os` remains excluded from the runnable package;
     - v0.7.0 scripts should be forked from v0.6.0 scripts, not parameterized yet.

3. **Committed v0.6.33B release-boundary/package-root rehearsal plan**
   - Added:
     - `docs/231-v0.6.33-release-boundary-package-root-rehearsal-plan-v0.md`
   - Updated:
     - `docs/README.md`
   - Commit:
     - `83fefb8 docs: plan release boundary rehearsal`
   - The plan selected:
     - fork-first v0.7.0 package scripts;
     - `full-validate-release` as primary release-validation wording;
     - `full` as compatibility wording through v0.7.0;
     - `carbonstack`, `carbonstack-comms`, and `carbonstack-cypher` as the runnable package set;
     - `carbonstack-os` excluded from the runnable package;
     - release assets modeled after v0.6.0 but updated to v0.7.0 candidate naming;
     - toolchain-floor wording captured in release metadata/runbook instead of silently changing `go.mod` floors.
   - The first scripted push failed auth; the manual retry succeeded.
   - `COMMAND_REFERENCE.v0.md` remained `entries=87` at the v0.6.33B planning checkpoint.
   - No package rehearsal happened during v0.6.33B.

4. **Created v0.7.0-specific package staging/rehearsal scripts**
   - Added:
     - `scripts/stage-v0.7.0-package.sh`
     - `scripts/rehearse-v0.7.0-package.sh`
   - Script strategy:
     - v0.6.0 scripts remain historical prior art;
     - v0.7.0 scripts are explicit/forked and committed before use as evidence.
   - The staging script:
     - requires clean tracked `carbonstack`, `carbonstack-comms`, and `carbonstack-cypher`;
     - archives tracked HEAD snapshots;
     - excludes `carbonstack-os`;
     - writes `release/manifest.json`;
     - writes release notes, testing runbook, validation freeze, and LICENSE;
     - writes and verifies package checksums;
     - creates the v0.7.0 package archive;
     - copies staged release-candidate assets to the stage root;
     - writes asset checksums.
   - The rehearsal script:
     - calls the staging script;
     - verifies staged release asset checksums;
     - extracts the archive into a fresh rehearsal root;
     - verifies package checksums from the fresh extraction;
     - runs `full-validate-release` from the fresh extraction with `--clean-generated`;
     - performs final forbidden-artifact checks.

5. **Updated registry/reference and runner wording**
   - Updated:
     - `registry/commands.v0.yaml`
     - `registry/COMMAND_REFERENCE.v0.md`
     - `tools/carbonstack-validate/main.go`
   - Added registry entries:
     - `carbonstack.script.stage-v0.7.0-package`
     - `carbonstack.script.rehearse-v0.7.0-package`
   - Regenerated command reference:
     - before v0.6.33C: `entries=87`;
     - after v0.6.33C: `entries=89`.
   - Runner wording now includes `full-validate-release` in:
     - help;
     - unknown-profile expected-profile output;
     - profile-output wording for the shared `full` / `full-validate-release` branch.
   - No new runner validation profile was added.
   - `full-validate-release` remains exact release-validation alias behavior:
     - `release-snapshot`
     - `local-cypher`

6. **Committed v0.6.33C script/registry/reference checkpoint before rehearsal**
   - Added:
     - `docs/232-v0.6.33-v0.7.0-package-scripts-and-rehearsal-prep-v0.md`
     - `scripts/stage-v0.7.0-package.sh`
     - `scripts/rehearse-v0.7.0-package.sh`
   - Updated:
     - `docs/README.md`
     - `registry/commands.v0.yaml`
     - `registry/COMMAND_REFERENCE.v0.md`
     - `tools/carbonstack-validate/main.go`
   - Commit:
     - `0f9b8d7 release: add v0.7.0 package rehearsal scripts`
   - Commit summary:
     - `7 files changed, 754 insertions(+), 5 deletions(-)`
     - `create mode 100644 docs/232-v0.6.33-v0.7.0-package-scripts-and-rehearsal-prep-v0.md`
     - `create mode 100755 scripts/rehearse-v0.7.0-package.sh`
     - `create mode 100755 scripts/stage-v0.7.0-package.sh`
   - Post-commit validation passed:
     - generated reference deterministic check with `entries=89`;
     - runner `go test ./...`;
     - registry missing-nonclaims scan.
   - Push succeeded:
     - `83fefb8..0f9b8d7 main -> main`.

7. **Ran v0.6.33D actual v0.7.0-style package rehearsal from fresh extraction**
   - Rehearsal command:
     - `bash scripts/rehearse-v0.7.0-package.sh`
   - Stage root:
     - `/tmp/carbonstack-v0.7.0-stage`
   - Fresh extraction root:
     - `/tmp/carbonstack-v0.7.0-rehearsal/extract/carbonstack-v0.7.0-package`
   - Archive:
     - `/tmp/carbonstack-v0.7.0-stage/carbonstack-v0.7.0-cumulative-pre-alpha-engineering-boundary-pre-release.tgz`
   - Stage included:
     - `carbonstack`
     - `carbonstack-comms`
     - `carbonstack-cypher`
   - Stage excluded:
     - `carbonstack-os`
   - Package checksums:
     - `checksum entries: 491`
   - Staged release-candidate assets:
     - `carbonstack-v0.7.0-cumulative-pre-alpha-engineering-boundary-pre-release.tgz`
     - `carbonstack-v0.7.0-release-manifest.json`
     - `carbonstack-v0.7.0-package-checksums.txt`
     - `carbonstack-v0.7.0-asset-checksums.txt`
     - `carbonstack-v0.7.0-validation-freeze.md`
     - `v0.7.0-testing-runbook.md`
     - `v0.7.0-release-notes.md`
     - `LICENSE`
   - Asset checksum verification passed.
   - Package checksum verification passed from fresh extraction.
   - `full-validate-release` passed from fresh extraction with `--clean-generated`.
   - Cleanup removed expected OpenMLS sidecar generated state/target inside the throwaway extraction.
   - Final package artifact scan passed:
     - required release files present;
     - registry README / commands / generated command reference present;
     - `COMMAND_BOUNDARY_TABLE.v0.md` present as optional file;
     - no `.git`;
     - no Rust `target`;
     - no sidecar generated state;
     - no `cypher.db`;
     - no SQLite WAL/SHM;
     - no provider storage;
     - no signer file;
     - no raw private LogDoc;
     - no raw private breakpoint.
   - Final result:
     - `PACKAGE REHEARSAL PASSED`.

8. **Preserved release boundary**
   - v0.6.33 is **not**:
     - a public release;
     - v0.7.0 tag;
     - release asset upload;
     - public page;
     - production secure messaging;
     - production E2EE;
     - deployment validation;
     - hostile-server safety;
     - adversarial harness implementation.
   - It is a successful release-candidate package rehearsal checkpoint.
   - The staged release-candidate assets live under `/tmp` and are not tracked repo files.
   - The clean working tree after rehearsal is expected:
     - the committed script/registry changes are already pushed;
     - the rehearsal output was intentionally outside the repo.

Current commits:

    carbonstack        0f9b8d7 release: add v0.7.0 package rehearsal scripts
    carbonstack-comms  5a61646 fix: align state audit output with boundary model
    carbonstack-cypher d18a564 chore: restore validated Go module floor
    carbonstack-os     1bbbe52 docs: clarify CarbonStackOS target direction

Current public release remains:

    v0.6.0 State/UX Boundary Validation Pre-Release

This is **not** a public product release.

### 0.1 Carried-forward v0.6.32 baseline summary

v0.6.32 is complete as a **public surface / command registry / manual-CLI docs inspection and selected discoverability cleanup checkpoint**.

It did seven bounded things:

1. Ran v0.6.32A public/operator-surface recon before release-boundary rehearsal.
2. Validated Comms and Cypher test posture during docs/release-readiness recon.
3. Identified public/manual docs and CLI discoverability gaps.
4. Patched runner help discoverability for `full-validate-release`.
5. Patched runner/front/docs README surfaces and added `docs/230-v0.6.32-public-surface-manual-docs-refresh-v0.md`.
6. Preserved registry/reference/profile boundaries.
7. Validated and pushed the checkpoint at `f9d4ac4`.

v0.6.32 preserved the key boundary:

    public/manual discoverability improved.
    What the system proves did not expand.
    No release rehearsal, release asset creation, deployment validation, production security claim, or hostile-server claim occurred.

## 1. Current Project Goal

**Active goal:** Continue from the successful v0.6.33 release-boundary/package-root rehearsal checkpoint into **v0.6.34A adversarial release-candidate/package-surface audit**, not immediate public release cutting.

After v0.6.33, the project has:

- the v0.6.0 public pre-release still live on Gitea;
- v0.7.0-specific staging/rehearsal scripts committed and pushed;
- `docs/231-v0.6.33-release-boundary-package-root-rehearsal-plan-v0.md`;
- `docs/232-v0.6.33-v0.7.0-package-scripts-and-rehearsal-prep-v0.md`;
- generated `registry/COMMAND_REFERENCE.v0.md`;
- generated command reference count now **89**;
- registry rows for v0.7.0 stage/rehearse package scripts;
- `full-validate-release` as the explicit release/package-root validation name mapped exactly to current `full`;
- `full-validate-release` visible in runner `--help` and unknown-profile expected output;
- runner wording that names the active profile in the shared `full` / `full-validate-release` branch;
- v0.7.0 rehearsal package assets staged under `/tmp`;
- successful package checksum verification from staged package and fresh extraction;
- successful asset checksum verification;
- successful `full-validate-release` from fresh extraction;
- successful final forbidden-artifact scan;
- a clean aligned `carbonstack` tree after rehearsal;
- expected ignored local generated/dev artifacts in Comms and Cypher only.

CarbonStack's broader intent remains:

- **CarbonStack** is the secure-communications appliance-stack project and shared doctrine/release authority.
- **CarbonStackComms** is the text-first messaging client and current OpenMLS sidecar / Relay Space integration surface.
- **CarbonStackCypher** is the self-hostable relay/storage server/API stack.
- **CarbonStackOS** is future constrained Android-derived appliance OS work and remains out of current runnable validation packages.
- The `carbonstack` Gitea repo remains the source of truth for release state, release assets, docs, registry, command-boundary classification, and public claims.
- GitHub mirrors, if present, are secondary push mirrors for discoverability/redundancy only.

**North star:**

> Its goal is not to be a better smartphone. Its goal is to stop being a smartphone.

**Core doctrine:**

> Every feature is guilty until it proves it does not add unacceptable parser, network, sensor, identity, or filesystem authority.

**v0.6.x engineering doctrine after v0.6.33:**

> The v0.7.0-style package rehearsal has passed from a committed, fresh-extracted, checksum-verified package root. The next work should adversarially audit the staged release-candidate package, release metadata, release notes, testing runbook, asset list, Gitea release description shape, public/manual claims, and nonclaims before any tag, upload, or public page. Do not treat the passed rehearsal as permission to cut v0.7.0 without release-surface audit.

## 2. Current Public Release and Mainline State

Current public release remains:

    https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.6.0

Release title:

    CarbonStack v0.6.0 State/UX Boundary Validation Pre-Release

Release tag:

    v0.6.0

Current post-rewrite tag target:

    70d4318 v0.6.0

Pre-rewrite v0.6.0 tag target for historical context only:

    7e676cd6ad / 7e676cd scripts: add v0.6.0 package rehearsal

Current mainline checkpoint:

    v0.6.33 Release-Boundary / Package-Root Rehearsal Checkpoint

Current CarbonStack head:

    0f9b8d7 release: add v0.7.0 package rehearsal scripts

Current CarbonStackComms head:

    5a61646 fix: align state audit output with boundary model

v0.6.33 is not a new public release and does not change the public release title. It is a post-v0.6.0 release-boundary/package-root rehearsal checkpoint for a candidate v0.7.0 cumulative pre-alpha engineering boundary.

Important release-page warning still applies:

    Gitea automatically provides Source Code ZIP/TAR.GZ archives for the carbonstack repo at a tag.
    These are not the intended runnable multi-repo validation package.
    Use the attached package, manifest, checksums, validation freeze, release notes, testing runbook, and LICENSE for the intended multi-repo validation package if/when v0.7.0 is cut.

v0.6.33 did not tag v0.7.0, upload release assets, or create a public release page.

New v0.6.33 docs/checkpoints:

    docs/231-v0.6.33-release-boundary-package-root-rehearsal-plan-v0.md
    docs/232-v0.6.33-v0.7.0-package-scripts-and-rehearsal-prep-v0.md

New v0.6.33 scripts:

    scripts/stage-v0.7.0-package.sh
    scripts/rehearse-v0.7.0-package.sh

Generated command reference count is now:

    entries=89

## 3. Current Repo Heads After v0.6.33

Final pushed heads:

    carbonstack        0f9b8d7 release: add v0.7.0 package rehearsal scripts
    carbonstack-comms  5a61646 fix: align state audit output with boundary model
    carbonstack-cypher d18a564 chore: restore validated Go module floor
    carbonstack-os     1bbbe52 docs: clarify CarbonStackOS target direction

Recent pushed history:

    carbonstack:
        0f9b8d7 release: add v0.7.0 package rehearsal scripts
        83fefb8 docs: plan release boundary rehearsal
        f9d4ac4 docs: refresh public manual surfaces
        3b9370a docs: define adversarial harness contract
        06f2c34 feat: add explicit release validation profile alias
        d31537c docs: define validation profile naming boundary
        44e781c docs: record state audit boundary alignment

    carbonstack-comms:
        5a61646 fix: align state audit output with boundary model
        2d13c45 fix: clarify load-check provider metadata output
        61b2e51 fix: label message sender metadata as unverified
        b42aa09 fix: make Welcome join state writes atomic
        53779d5 test: cover message wrapper residual edge cases
        462f796 chore: group Comms CLI help by command boundary
        020c07f test: align wrapper smoke with message wrappers
        9391029 feat: add opinionated OpenMLS message wrappers

    carbonstack-cypher:
        d18a564 chore: restore validated Go module floor
        52e3673 docs: make Cypher surface evergreen
        3589dbe docs: define local Cypher validation state contract
        15760c6 docs: record provider OpenMLS join boundary
        09afa0b feat: add Relay Space scoped envelope routes
        384affd feat: add Relay Space HTTP API routes
        cc6e589 feat: add Relay Space DB helpers
        6798d11 feat: add Relay Space schema substrate

    carbonstack-os:
        1bbbe52 docs: clarify CarbonStackOS target direction
        f984cdf Upload files to "/"
        c764790 chore: fix readme formatting
        953b006 chore: fix readme formatting
        b537475 Add CarbonStackOS north star and initial appliance model
        dab2792 Initial CarbonStack repository structure

Working tree note after v0.6.33D:

    carbonstack was clean/aligned with origin/main:
        nothing to commit, working tree clean
    carbonstack-comms was clean/aligned except ignored local artifacts:
        !! internal/protocol/mls/openmls-sidecar/.carbonstack-openmls-sidecar-state/
        !! internal/protocol/mls/openmls-sidecar/target/
    carbonstack-cypher retained the known ignored local DB:
        !! cypher.db
    carbonstack-os was clean/aligned.

The clean `carbonstack` tree after rehearsal is expected because all package/rehearsal output was produced under `/tmp`, not tracked into the repo.

## 4. Validation / Evidence State After v0.6.33

v0.6.33A preflight validation passed:

    generated reference deterministic check with entries=87;
    runner go test ./...;
    registry missing-nonclaims scan;
    runner doctor;
    required v0.6.27-v0.6.32 docs check;
    Comms go test ./... soft probe;
    Cypher go test ./... soft probe;
    OpenMLS sidecar cargo test soft probe;
    runner help/profile discoverability checks;
    registry release-supported list;
    registry front-readme-only list;
    release/package file inventory;
    release runner source inspection;
    package-shape inventory;
    release-boundary seed plan generation;
    no tracked mutation guard.

v0.6.33B docs/model validation passed:

    generated reference deterministic check with entries=87;
    runner go test ./...;
    registry missing-nonclaims scan;
    runner doctor;
    required v0.6.27-v0.6.32 docs check;
    v0.6.0 script syntax checks;
    docs/231 creation;
    docs README update;
    diff/check validation;
    registry/reference/script/runner guard;
    commit at 83fefb8;
    manual push succeeded after auth failure.

v0.6.33C implementation validation passed:

    v0.7.0 script syntax checks;
    registry row indentation repair;
    generated command reference regeneration;
    generated reference deterministic check with entries=89;
    runner go test ./...;
    registry missing-nonclaims scan;
    registry lookup for stage-v0.7.0-package;
    registry lookup for rehearse-v0.7.0-package;
    runner help includes full-validate-release;
    unknown-profile output includes full-validate-release;
    commit at 0f9b8d7;
    post-commit generated reference check with entries=89;
    post-commit runner go test ./...;
    post-commit registry missing-nonclaims scan;
    push rc=0.

v0.6.33D package rehearsal validation passed:

    stage helper ran from committed script;
    tracked cleanliness checked for carbonstack, comms, cypher;
    carbonstack-os explicitly excluded from runnable package;
    tracked HEAD snapshots archived into package root;
    package release metadata written;
    package checksums written;
    checksum entries = 491;
    package checksums verified in the staged package;
    package archive created;
    release-candidate assets copied to stage root;
    asset checksums written;
    asset checksums verified;
    archive extracted to fresh rehearsal root;
    package checksums verified from fresh extraction;
    full-validate-release ran from fresh extraction with --clean-generated;
    release-snapshot/core/local-cypher path passed through the release-validation ladder;
    known generated OpenMLS sidecar state/target removed by cleanup;
    final package artifact scan passed;
    final carbonstack working tree was clean.

Generated command reference is now:

    entries=89

New package/rehearsal output was under `/tmp`:

    /tmp/carbonstack-v0.7.0-stage
    /tmp/carbonstack-v0.7.0-rehearsal

Do not overclaim that these staged files are public release assets until they are intentionally uploaded to a Gitea release.

## 5. v0.6.33 Model Content Summary

### 5.1 New release-boundary plan doc

New doc:

    docs/231-v0.6.33-release-boundary-package-root-rehearsal-plan-v0.md

Core model:

    v0.7.0 should use a v0.6.0-like release surface but not blindly reuse v0.6.0 scripts or stale release text.
    The safest strategy is fork-first v0.7.0 package scripts.
    `full-validate-release` is primary release-validation wording.
    `full` remains compatibility wording through v0.7.0.
    Package shape remains carbonstack + comms + cypher + release metadata.
    carbonstack-os remains excluded from the runnable package.

### 5.2 New package script prep doc

New doc:

    docs/232-v0.6.33-v0.7.0-package-scripts-and-rehearsal-prep-v0.md

Core model:

    v0.7.0 staging/rehearsal helpers exist and are registry-classified before their output is treated as evidence.
    The scripts are release-candidate rehearsal helpers, not release publishers.

### 5.3 New script surfaces

New scripts:

    scripts/stage-v0.7.0-package.sh
    scripts/rehearse-v0.7.0-package.sh

New registry IDs:

    carbonstack.script.stage-v0.7.0-package
    carbonstack.script.rehearse-v0.7.0-package

Both are:

    audience: dev
    maturity: experimental
    include_in_front_readme: false

Their nonclaims include:

    does not cut release;
    does not upload assets;
    does not include carbonstack-os;
    not production security;
    not deployment;
    not local-backbone;
    not hostile-server safety.

### 5.4 Generated command reference

Before v0.6.33C:

    entries=87

After v0.6.33C:

    entries=89

Reason:

    two new package-helper script registry entries.

### 5.5 Release-candidate asset shape

The rehearsal staged the following release-candidate asset names:

    carbonstack-v0.7.0-cumulative-pre-alpha-engineering-boundary-pre-release.tgz
    carbonstack-v0.7.0-release-manifest.json
    carbonstack-v0.7.0-package-checksums.txt
    carbonstack-v0.7.0-asset-checksums.txt
    carbonstack-v0.7.0-validation-freeze.md
    v0.7.0-testing-runbook.md
    v0.7.0-release-notes.md
    LICENSE

These are staged rehearsal outputs, not uploaded release assets yet.

### 5.6 Toolchain-floor posture after rehearsal

v0.6.33 did not change Go module floors.

Observed toolchain posture remains:

    runner go.mod: go 1.24
    comms go.mod: go 1.26.3
    cypher go.mod: go 1.26.3
    local WSL validation Go: go1.24.4 linux/amd64
    Rust/Cargo: 1.96.0
    SQLite: 3.46.1

The v0.7.0 staging script records observed toolchain versions in release metadata / validation freeze. This is acceptable for rehearsal but should be audited before release notes are finalized.

### 5.7 What v0.6.33 does not change

v0.6.33 does not add:

    release tag;
    release asset upload;
    public page;
    production secure messaging;
    production E2EE;
    hostile-server safety;
    deployment readiness;
    adversarial harness implementation;
    full-runtime-dev;
    state-audit mutation;
    PQ/hybrid implementation;
    backup/restore implementation;
    CarbonStackOS readiness.

## 6. v0.6.33 Blunders / Continuity Notes

### 6.1 Large Python paste corruption caused a real script parse failure

During the first v0.6.33B attempt, terminal display truncation was initially suspected. The actual failure was real file-content corruption in `/tmp/carbonstack-v0633b-release-boundary-plan-doc.py`, producing an unterminated Python string.

Impact:

    the Python script died at parse time;
    no repo mutation occurred;
    v0.6.33B had to be retried with a smaller Bash recovery script.

Lesson:

    Do not assume every weird pasteback is only terminal display truncation.
    When an interpreter reports a syntax error from the temp file, treat it as real file-content corruption.
    Prefer smaller Bash blocks or simpler scripts for huge doc writes.

### 6.2 v0.6.33B scripted push failed auth, manual push succeeded

The v0.6.33B script committed:

    83fefb8 docs: plan release boundary rehearsal

The scripted push failed authentication. Manual retry succeeded:

    f9d4ac4..83fefb8 main -> main

Impact:

    remote aligned after manual push;
    no repo corruption.

Lesson:

    First push auth failure remains a recurring non-blocker.
    Always check final remote alignment before proceeding to evidence-producing steps.

### 6.3 v0.6.33C first run stopped on the `entries=89` guard

The initial v0.6.33C attempt added two v0.7.0 registry rows, but they were inserted at top-level YAML indentation rather than inside the registry list.

Observed:

    new rows existed in `registry/commands.v0.yaml`;
    renderer still emitted `entries=87`;
    `grep -q "entries=89"` failed;
    `set -e` stopped before commit, push, or rehearsal.

Impact:

    no invalid registry/reference state was committed;
    no package rehearsal was run from malformed registry input.

Lesson:

    The generated command reference entry-count guard worked exactly as intended.
    Future registry patches should use the existing row indentation and insertion marker rather than appending top-level YAML fragments.

### 6.4 v0.6.33C repair stopped silently after partial-state check because `bash -n` failed

The next repair script seemed to stop after `git status`, but the silent failing command was actually a `bash -n` check against the generated stage script.

Diagnostic found:

    scripts/stage-v0.7.0-package.sh: line 389: unexpected EOF while looking for matching `"`

Cause:

    the stage script contained Markdown-heavy heredocs with unquoted expansion and backtick-like release text risk.
    The safer fix was to rewrite the stage script with simpler release text and avoid Markdown backticks inside expansion heredocs.

Impact:

    `scripts/rehearse-v0.7.0-package.sh` was syntactically fine;
    only the stage script needed repair;
    no commit had happened yet.

Lesson:

    Always echo script syntax checks when `set -e` is active.
    Avoid Markdown-heavy unquoted heredocs inside scripts that write release notes/runbooks.
    For release scripts, `bash -n` is mandatory before commit.

### 6.5 Clean tree after rehearsal is expected

After v0.6.33D, `git status` showed:

    Your branch is up to date with 'origin/main'.
    nothing to commit, working tree clean

This is correct.

Impact:

    committed v0.6.33C changes are already pushed;
    rehearsal outputs live under `/tmp`;
    no release artifact was accidentally tracked.

Lesson:

    Clean-tree after rehearsal is not evidence loss.
    It is the desired result for a staged/fresh-extraction rehearsal flow.

### 6.6 Terminal glue/paste noise persisted but final evidence was coherent

The raw logs still contained shell-tail/glue around heredoc invocations, but final source-of-truth checks were clean:

    committed source exists;
    generated command reference check passed;
    registry lookups passed;
    scripts passed `bash -n`;
    v0.6.33C commit and push succeeded;
    v0.6.33D package rehearsal passed;
    final repo tree was clean.

Lesson:

    Continue treating final committed source, deterministic generated-reference check, registry lookup, script syntax, package rehearsal output, and final clean-tree status as authoritative over transient terminal glue.

## 7. Critical Path / Function Updates

New critical docs:

    carbonstack/docs/231-v0.6.33-release-boundary-package-root-rehearsal-plan-v0.md
    carbonstack/docs/232-v0.6.33-v0.7.0-package-scripts-and-rehearsal-prep-v0.md

Updated docs index:

    carbonstack/docs/README.md

New package helper scripts:

    carbonstack/scripts/stage-v0.7.0-package.sh
    carbonstack/scripts/rehearse-v0.7.0-package.sh

Updated registry/reference paths:

    carbonstack/registry/commands.v0.yaml
    carbonstack/registry/COMMAND_REFERENCE.v0.md

Generated command reference:

    entries=89

New registry IDs:

    carbonstack.script.stage-v0.7.0-package
    carbonstack.script.rehearse-v0.7.0-package

Updated runner path:

    carbonstack/tools/carbonstack-validate/main.go

Critical docs still active:

    carbonstack/docs/223-v0.6.27-vault-security-nosilent-state-model-v0.md
    carbonstack/docs/224-v0.6.28-pq-hybrid-placement-model-v0.md
    carbonstack/docs/225-v0.6.29-bounded-state-substrate-mechanics-v0.md
    carbonstack/docs/226-v0.6.29-state-audit-dev-model-alignment-v0.md
    carbonstack/docs/227-v0.6.30-validation-profile-naming-boundary-v0.md
    carbonstack/docs/228-v0.6.30-full-validate-release-alias-v0.md
    carbonstack/docs/229-v0.6.31-adversarial-harness-contract-evidence-matrix-v0.md
    carbonstack/docs/230-v0.6.32-public-surface-manual-docs-refresh-v0.md
    carbonstack/docs/231-v0.6.33-release-boundary-package-root-rehearsal-plan-v0.md
    carbonstack/docs/232-v0.6.33-v0.7.0-package-scripts-and-rehearsal-prep-v0.md
    CarbonStack_Long_Term_Roadmap_v0.6.26_EVERGREEN.pdf

Critical package commands:

    cd carbonstack
    bash scripts/stage-v0.7.0-package.sh
    bash scripts/rehearse-v0.7.0-package.sh

Critical validation commands:

    cd carbonstack/tools/carbonstack-validate
    go run . --profile verify-checksums --root <package-root>
    go run . --profile full-validate-release --root <package-root> --clean-generated
    go run . --profile full --root <package-root> --clean-generated

Critical semantics:

    full-validate-release remains exact alias behavior to full.
    full remains compatibility wording.
    release-snapshot remains package-root validation.
    full-validate-release is preferred release-facing wording.
    full-runtime-dev still does not exist.
    same-state profiles are still not added to full/full-validate-release/release-snapshot.
    state-audit-dev remains outside full/release-snapshot.
    no adversarial profile exists.
    no hostile-server safety is claimed.

Important v0.6.33 behavior:

    v0.7.0 scripts exist and are committed.
    v0.7.0 scripts are registry-classified.
    COMMAND_REFERENCE entries=89.
    release-candidate package rehearsal passed from fresh extraction.
    release-candidate artifacts are staged under /tmp.
    no release tag or upload happened.

## 8. Future To-Do / Critical Path Forward

### 8.1 Immediate next safest rung

Recommended next checkpoint:

    v0.6.34A = adversarial release-candidate review / package surface audit

Purpose:

    inspect the staged v0.7.0 release-candidate package, release metadata, release notes, testing runbook, asset list, package contents, Gitea release page text, public/manual claims, and nonclaims before deciding whether to cut v0.7.0.

The audit should answer:

    Are the staged asset names correct?
    Are package and asset checksums suitable for upload?
    Does manifest schema/content match the v0.7.0 boundary?
    Does validation-freeze capture exact heads/toolchains without overclaiming?
    Do release notes describe actual v0.6.x closure rather than stale v0.6.0/v0.5.x runway?
    Does testing runbook use full-validate-release as primary wording?
    Does the release description clearly warn against Gitea default Source Code archives?
    Are nonclaims complete?
    Is carbonstack-os exclusion clear?
    Is toolchain-floor wording honest?
    Should there be a downloaded-asset / RC simulation before tag/upload?
    Are there any release-blocking stale strings in staged assets?
    Are there any raw private handoff artifacts or generated local artifacts in package contents?

Do not cut a release tag before this audit.

### 8.2 Near-term sequence

Recommended order:

    v0.6.34A adversarial release-candidate/package surface audit
    v0.6.34B release-candidate text/asset polish or downloaded-asset simulation, if audit finds issues
    v0.6.34C final release cut prep, if clean
    v0.7.0 cumulative pre-alpha engineering boundary release, only after explicit approval

Bias:

    audit the staged package before release upload.

Reason:

    v0.6.33 proved the package flow can pass.
    It did not prove the release page text, public claims, and asset presentation are finalized.

### 8.3 Deferred until later

Do not do next by default:

    production security wording;
    release tag without audit;
    release asset upload without audit;
    public WordPress/project page;
    adversarial runner profile;
    adversarial mega-profile;
    adversarial test implementation;
    full-runtime-dev implementation;
    adding state-audit-dev to a runtime aggregate;
    adding same-state profiles to full/full-validate-release/release-snapshot;
    full deprecation;
    production vault implementation;
    encryption/secure storage implementation;
    backup/restore implementation;
    PQ/hybrid implementation or experiments until later model-permitted epoch;
    Welcome/KeyPackage replay implementation;
    deployment guide claiming safety;
    Android/CarbonStackOS runnable package inclusion.

### 8.4 Future adversarial implementation posture

Unchanged from v0.6.31:

    first adversarial implementation should be leaf-first, not aggregate-first.
    no aggregate adversarial profile should exist until multiple leaf profiles stabilize.
    future cases must follow docs/229 contract fields.

## 9. Claim Boundaries After v0.6.33

v0.6.33 is **not**:

    a public release;
    v0.7.0 tag;
    release asset upload;
    public release page;
    production secure messaging;
    production E2EE;
    hostile-server safety;
    malicious relay safety;
    server equivocation detection;
    replay safety;
    Welcome replay handling;
    KeyPackage replay handling;
    metadata privacy;
    metadata authenticity;
    sender authenticity;
    verified identity;
    secure enrollment;
    production vault/key storage;
    secure storage;
    encryption-at-rest;
    production backup/restore;
    state rollback safety;
    PQ or hybrid security;
    quantum-safe messaging;
    OpenMLS-native PQ/hybrid support;
    ciphersuite migration;
    algorithm migration;
    Android/CarbonStackOS readiness;
    deployment readiness;
    audit or certification;
    general-public usability;
    `full-runtime-dev`;
    adversarial harness implementation.

v0.6.33 is:

    release-boundary/package-root rehearsal planning;
    v0.7.0-specific package-helper implementation;
    registry/reference classification of new package helpers;
    successful fresh-extraction v0.7.0-style package rehearsal;
    staged release-candidate asset generation under /tmp;
    evidence that the current package flow can pass validation.

`docs/231` is:

    a release-boundary/package-root rehearsal plan.

`docs/232` is:

    a package-helper implementation/rehearsal-prep note.

The v0.7.0 staged package output is:

    release-candidate rehearsal output only.
    It is not public release material until explicitly uploaded with a tag/release.

## 10. Continuity Anchor

Current safe resume point:

    CarbonStack v0.6.33 is complete.
    CarbonStack head is 0f9b8d7.
    CarbonStackComms head is 5a61646.
    CarbonStackCypher head is d18a564.
    CarbonStackOS head is 1bbbe52.
    COMMAND_REFERENCE entries = 89.
    docs/231 and docs/232 exist.
    scripts/stage-v0.7.0-package.sh exists and passed syntax/rehearsal.
    scripts/rehearse-v0.7.0-package.sh exists and passed syntax/rehearsal.
    v0.7.0 package rehearsal passed from fresh extraction.
    full-validate-release passed from fresh extraction.
    asset checksums verified.
    package checksums verified.
    final package artifact scan passed.
    final carbonstack tree is clean/aligned with origin/main.
    current public release remains v0.6.0.
    no v0.7.0 tag or asset upload has happened.

Staged rehearsal artifacts from the successful run:

    /tmp/carbonstack-v0.7.0-stage/carbonstack-v0.7.0-cumulative-pre-alpha-engineering-boundary-pre-release.tgz
    /tmp/carbonstack-v0.7.0-stage/carbonstack-v0.7.0-release-manifest.json
    /tmp/carbonstack-v0.7.0-stage/carbonstack-v0.7.0-package-checksums.txt
    /tmp/carbonstack-v0.7.0-stage/carbonstack-v0.7.0-asset-checksums.txt
    /tmp/carbonstack-v0.7.0-stage/carbonstack-v0.7.0-validation-freeze.md
    /tmp/carbonstack-v0.7.0-stage/v0.7.0-testing-runbook.md
    /tmp/carbonstack-v0.7.0-stage/v0.7.0-release-notes.md
    /tmp/carbonstack-v0.7.0-stage/LICENSE

A future assistant should not proceed directly to release tag/upload. The next checkpoint should audit the staged package/release-candidate surface against v0.6.0 release-shape expectations and current v0.7.0 nonclaims.

---

## Appendix A. Carried-forward CarbonStack LogDoc v0.6.32

The following is the full prior v0.6.32 ledger carried forward for continuity. Current v0.6.33 sections above supersede the current-state metadata, repo heads, generated command reference count, next-step ordering, package rehearsal status, and checkpoint summary where they differ.

# CarbonStack LogDoc v0.6.32

**Last updated:** 2026-07-02 local session  
**Owner/operator:** bitcrusher32  
**Assistant/model:** GPT-5.5 Thinking  
**Current phase:** Phase 2F -> **v0.6.32 public surface / command registry / manual-CLI docs inspection and selected discoverability cleanup checkpoint complete; v0.6.32A completed a no-mutation recon over public/operator-facing surfaces, command registry, generated command reference, runner help, README/manual docs, state/recovery/no-silent docs, deployment warnings, and toolchain-floor clarity before release-boundary rehearsal; v0.6.32B committed a narrow public/manual docs + CLI discoverability cleanup that made `full-validate-release` visible in runner help, added a first-class runner README section, added current operator-surface pointers to the front README, added `docs/230-v0.6.32-public-surface-manual-docs-refresh-v0.md`, and updated docs README without changing registry semantics, generated command reference count, release validation behavior, adding `full-runtime-dev`, implementing adversarial cases, or making deployment/security/hostile-server claims.**  
**Current public release:** `v0.6.0 State/UX Boundary Validation Pre-Release`  
**Current mainline checkpoint:** `v0.6.32 Public Surface / Manual Docs Refresh Checkpoint`  
**Release URL:** https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.6.0  
**Update source:** `CarbonStackLogDocV0.6.31`, `CarbonStack_BreakpointV0.6.31`, v0.6.32A public-surface/manual-CLI docs recon log, v0.6.32B public/manual docs + CLI discoverability cleanup log, and the active `CarbonStack_Long_Term_Roadmap_v0.6.26_EVERGREEN` planning authority.  
**Update purpose:** Preserve the v0.6.32 checkpoint after the v0.6.31 adversarial harness contract made it safe to inspect and refresh public/operator-facing surfaces; record the v0.6.32A/B timeline, selected docs/help cleanup, validation evidence, paste-corruption continuity, repo heads, critical path/function updates, future release-boundary/package-root rehearsal ordering, and nonclaims; preserve the full prior Markdown ledger; and keep the JSON breakpoint as a lean current-state handoff.

**Version schema:** `v[scope].[timeline]`. This file is `v0.6.32`, the working branch continuity ledger after public surface / command registry / manual-CLI docs inspection and selected discoverability cleanup in the v0.6.x OpenMLS message-flow integration and boundary-hardening epoch.

**LogDoc V2 usage note:** The Markdown LogDoc is intentionally a dev continuity ledger. It preserves historical v0.6.x timeline data, blunders, validation ladders, command-boundary decisions, and future path constraints. The JSON breakpoint remains a lean current-state handoff and should not try to duplicate the full Markdown timeline.

**Privacy / history-rewrite note:** Exact raw shell logs may contain local usernames, hostnames, email addresses, and old pre-rewrite author identities. This LogDoc avoids repeating local host-derived identity strings except where operationally necessary. Public commit identity remains normalized on public `main`/tag histories. Pre-v0.6.13 commit hashes from older LogDocs and historical docs are now historical **pre-rewrite** references and must not be treated as current Gitea `main`/tag SHAs.

## 0. Executive Summary

v0.6.32 is complete as a **public surface / command registry / manual-CLI docs inspection and selected discoverability cleanup checkpoint**.

It did seven bounded things:

1. **Ran v0.6.32A public/operator-surface recon before release-boundary rehearsal**
   - v0.6.32A was a no-mutation recon over:
     - front README;
     - docs README;
     - command registry;
     - generated command reference;
     - runner README/help;
     - normal-message docs;
     - Relay onboarding docs;
     - validation docs;
     - state-audit / recovery / no-silent docs;
     - deployment-warning docs;
     - public-surface claim wording;
     - toolchain floor consistency.
   - The recon confirmed the v0.6.31 baseline:
     - `carbonstack` at `3b9370a`;
     - `carbonstack-comms` at `5a61646`;
     - `carbonstack-cypher` at `d18a564`;
     - `carbonstack-os` at `1bbbe52`;
     - `COMMAND_REFERENCE.v0.md` current with `entries=87`;
     - runner tests pass;
     - registry missing-nonclaims scan passes;
     - runner doctor passes.
   - v0.6.32A did not mutate docs, registry, generated reference, runner code, or any repo.

2. **Validated Comms and Cypher test posture during docs/release-readiness recon**
   - Soft probe validation passed:
     - CarbonStackComms `go test ./... -count=1`;
     - CarbonStackCypher `go test ./... -count=1`.
   - This does not make v0.6.32 a release rehearsal.
   - It confirms the source tree remains coherent enough to proceed toward release-boundary/package-root rehearsal planning.

3. **Identified public/manual docs and CLI discoverability gaps**
   - v0.6.32A found:
     - `full-validate-release` existed in the runner switch and registry, but was absent from runner `--help`;
     - the registry front-README candidate list had advanced beyond the front README;
     - runner README contained a v0.6.30 alias note but lacked a first-class `full-validate-release` section;
     - docs README referenced recent numbered docs but did not clearly expose the normal-message wrapper surface;
     - toolchain-floor clarity was needed before release-boundary/package-root rehearsal.
   - The recon explicitly showed:
     - `full_validate_release_in_switch=True`;
     - `full_validate_release_in_help=False`;
     - `full_validate_release_in_main_profile_help=False`.
   - Toolchain floor inspection showed:
     - runner `go.mod`: `go 1.24`;
     - Comms `go.mod`: `go 1.26.3`;
     - Cypher `go.mod`: `go 1.26.3`;
     - OpenMLS sidecar Rust crate: edition `2024`, OpenMLS `0.8.1`.
   - Decision:
     - do selected docs/help cleanup before release-boundary rehearsal;
     - do not change `go.mod` floors in this checkpoint.

4. **Patched runner help discoverability**
   - Updated:
     - `tools/carbonstack-validate/main.go`
   - Runner `--help` now lists:
     - `full-validate-release`
   - The alias remains behaviorally unchanged:
     - `full-validate-release` still maps to the same branch as `full`;
     - no release-validation semantics changed.

5. **Patched runner/front/docs README surfaces and added docs/230**
   - Updated:
     - `README.md`
     - `docs/README.md`
     - `tools/carbonstack-validate/README.md`
   - Added:
     - `docs/230-v0.6.32-public-surface-manual-docs-refresh-v0.md`
   - CarbonStack commit:
     - `f9d4ac4 docs: refresh public manual surfaces`
   - Commit summary:
     - `5 files changed, 253 insertions(+), 1 deletion(-)`
     - `create mode 100644 docs/230-v0.6.32-public-surface-manual-docs-refresh-v0.md`
   - The runner README now has a first-class `full-validate-release` section and a toolchain floor note before v0.7.0 rehearsal.
   - The front README now points to current operator/developer surfaces:
     - `full-validate-release`;
     - `full`;
     - `release-snapshot`;
     - `message-send-dev`;
     - `message-inbox-dev`;
     - `state-audit-dev`;
     - `registry/COMMAND_REFERENCE.v0.md`;
     - `docs/229-v0.6.31-adversarial-harness-contract-evidence-matrix-v0.md`.
   - Docs README now references docs/230 and current wrapper/state/adversarial-contract pointers.

6. **Preserved registry/reference/profile boundaries**
   - v0.6.32B did not change:
     - `registry/commands.v0.yaml`;
     - `registry/COMMAND_REFERENCE.v0.md`;
     - generated command reference count;
     - registry semantics;
     - `full` / `full-validate-release` / `release-snapshot` behavior.
   - Generated command reference remains:
     - `entries=87`.
   - v0.6.32B did not add:
     - runner profile;
     - registry row;
     - `full-runtime-dev`;
     - adversarial implementation;
     - state-audit mutation;
     - release-boundary/package-root rehearsal;
     - deployment/security/hostile-server claim.

7. **Validated and pushed the checkpoint**
   - Baseline validation passed:
     - generated reference deterministic check with `entries=87`;
     - runner `go test ./...`;
     - registry missing-nonclaims scan;
     - runner doctor;
     - required v0.6.27-v0.6.31 docs check.
   - Post-patch validation passed:
     - CarbonStack `git diff --check`;
     - generated reference deterministic check unchanged with `entries=87`;
     - runner `go test ./...`;
     - registry missing-nonclaims scan;
     - runner help now includes `full-validate-release`;
     - all public/manual marker checks passed;
     - registry/reference guard passed.
   - Commit validation passed:
     - cached diff check;
     - generated reference deterministic check post-commit with `entries=87`;
     - runner `go test ./...` post-commit;
     - registry missing-nonclaims scan post-commit.
   - Push succeeded without the prior recurring auth failure:
     - `3b9370a..f9d4ac4 main -> main`.
   - Final repo snapshot showed all repos aligned with `origin/main`; Comms and Cypher retained only expected ignored generated/dev artifacts.

Current commits:

    carbonstack        f9d4ac4 docs: refresh public manual surfaces
    carbonstack-comms  5a61646 fix: align state audit output with boundary model
    carbonstack-cypher d18a564 chore: restore validated Go module floor
    carbonstack-os     1bbbe52 docs: clarify CarbonStackOS target direction

Current public release remains:

    v0.6.0 State/UX Boundary Validation Pre-Release

This is **not** a public product release.

### 0.1 Carried-forward v0.6.31 baseline summary

v0.6.31 is complete as an **adversarial harness contract / evidence-matrix docs checkpoint**.

It did six bounded things:

1. Ran v0.6.31A deep recon before writing the contract.
2. Mapped adversarial case-family vocabulary.
3. Mapped manual/CLI docs gaps without folding them into v0.6.31B.
4. Added:
   - `docs/229-v0.6.31-adversarial-harness-contract-evidence-matrix-v0.md`
5. Preserved the release/runtime/profile boundary.
6. Validated and pushed the checkpoint.

v0.6.31 preserved the key boundary:

    existing same-state and failure-hardening profiles are evidence seeds.
    They are not adversarial harness coverage by default.
    Future adversarial work must not silently enter full, full-validate-release, release-snapshot, or full-runtime-dev.

## 1. Current Project Goal

**Active goal:** Continue the v0.6.x OpenMLS message-flow integration and boundary-hardening epoch from the v0.6.32 public/manual docs refresh checkpoint toward v0.7.0 release-boundary/package-root rehearsal **preflight and planning**, not immediate release packaging.

After v0.6.32, the project has:

- the v0.6.0 public pre-release still live on Gitea;
- generated `registry/COMMAND_REFERENCE.v0.md`;
- deterministic generated-reference freshness enforced by runner tests;
- registry entry count still **87**;
- `full-validate-release` as the explicit release/package-root validation name mapped exactly to current `full`;
- `full-validate-release` visible in runner `--help`;
- runner README first-class documentation for `full-validate-release`;
- a front README current operator-surface section;
- `docs/230-v0.6.32-public-surface-manual-docs-refresh-v0.md`;
- docs README pointers to docs/230 and current wrapper/state/adversarial-contract surfaces;
- legacy stub commands gated behind explicit opt-in;
- opinionated `message-send-dev` / `message-inbox-dev` wrappers as the recommended dev/pre-alpha normal-message path;
- direct `openmls-*` commands retained as lower-level implementation/proof surfaces;
- Relay onboarding boundaries separated from ordinary message inbox behavior;
- deterministic normal-message failure/classification profiles;
- Welcome join partial-state safety and `same-state-welcome-join-failure-dev`;
- v0.6.27 state/security/no-silent model;
- v0.6.28 PQ/hybrid placement model;
- v0.6.29 bounded state-substrate / `state-audit-dev` output model alignment;
- v0.6.30 validation naming boundary and explicit release-validation alias;
- v0.6.31 adversarial harness contract/evidence matrix;
- v0.6.32 public/manual discoverability cleanup.

CarbonStack's broader intent remains:

- **CarbonStack** is the secure-communications appliance-stack project and shared doctrine/release authority.
- **CarbonStackComms** is the text-first messaging client and current OpenMLS sidecar / Relay Space integration surface.
- **CarbonStackCypher** is the self-hostable relay/storage server/API stack.
- **CarbonStackOS** is future constrained Android-derived appliance OS work and remains out of current runnable validation packages.
- The `carbonstack` Gitea repo remains the source of truth for release state, release assets, docs, registry, command-boundary classification, and public claims.
- GitHub mirrors, if present, are secondary push mirrors for discoverability/redundancy only.

**North star:**

> Its goal is not to be a better smartphone. Its goal is to stop being a smartphone.

**Core doctrine:**

> Every feature is guilty until it proves it does not add unacceptable parser, network, sensor, identity, or filesystem authority.

**v0.6.x engineering doctrine after v0.6.32:**

> The public/manual surface now points at the current validation, wrapper, state-audit, and adversarial-contract surfaces without changing what those surfaces prove. The next work should preflight the v0.7.0 release-boundary/package-root rehearsal plan, including package shape, included repo heads, source hygiene, runbook/manifest/checksum material, toolchain-floor wording, claim-boundary review, and explicit nonclaims. Do not jump directly to release tagging or asset upload.

## 2. Current Public Release and Mainline State

Current public release remains:

    https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.6.0

Release title:

    CarbonStack v0.6.0 State/UX Boundary Validation Pre-Release

Release tag:

    v0.6.0

Current post-rewrite tag target:

    70d4318 v0.6.0

Pre-rewrite v0.6.0 tag target for historical context only:

    7e676cd6ad / 7e676cd scripts: add v0.6.0 package rehearsal

Current mainline checkpoint:

    v0.6.32 Public Surface / Manual Docs Refresh Checkpoint

Current CarbonStack head:

    f9d4ac4 docs: refresh public manual surfaces

Current CarbonStackComms head:

    5a61646 fix: align state audit output with boundary model

v0.6.32 is not a new public release and does not change the release title. It is a post-release public/manual docs and CLI discoverability checkpoint.

Important release-page warning retained:

    Gitea automatically provides Source Code ZIP/TAR.GZ archives for the carbonstack repo at the tag.
    These are not the intended runnable multi-repo validation package.
    Use the attached v0.6.0 package, manifest, checksums, validation freeze, release notes, testing runbook, and LICENSE.

v0.6.32 did not rebuild v0.6.0 release assets and did not change the v0.6.0 release tag.

New v0.6.32 docs/checkpoint:

    docs/230-v0.6.32-public-surface-manual-docs-refresh-v0.md

No new v0.6.32 registry/profile surface was added.

Generated command reference count remains:

    entries=87

## 3. Current Repo Heads After v0.6.32

Final pushed heads:

    carbonstack        f9d4ac4 docs: refresh public manual surfaces
    carbonstack-comms  5a61646 fix: align state audit output with boundary model
    carbonstack-cypher d18a564 chore: restore validated Go module floor
    carbonstack-os     1bbbe52 docs: clarify CarbonStackOS target direction

Recent pushed history:

    carbonstack:
        f9d4ac4 docs: refresh public manual surfaces
        3b9370a docs: define adversarial harness contract
        06f2c34 feat: add explicit release validation profile alias
        d31537c docs: define validation profile naming boundary
        44e781c docs: record state audit boundary alignment
        2f0be42 docs: define bounded state substrate mechanics
        4a66681 docs: define PQ hybrid placement model

    carbonstack-comms:
        5a61646 fix: align state audit output with boundary model
        2d13c45 fix: clarify load-check provider metadata output
        61b2e51 fix: label message sender metadata as unverified
        b42aa09 fix: make Welcome join state writes atomic
        53779d5 test: cover message wrapper residual edge cases
        462f796 chore: group Comms CLI help by command boundary
        020c07f test: align wrapper smoke with message wrappers
        9391029 feat: add opinionated OpenMLS message wrappers

    carbonstack-cypher:
        d18a564 chore: restore validated Go module floor
        52e3673 docs: make Cypher surface evergreen
        3589dbe docs: define local Cypher validation state contract
        15760c6 docs: record provider OpenMLS join boundary
        09afa0b feat: add Relay Space scoped envelope routes
        384affd feat: add Relay Space HTTP API routes
        cc6e589 feat: add Relay Space DB helpers
        6798d11 feat: add Relay Space schema substrate

    carbonstack-os:
        1bbbe52 docs: clarify CarbonStackOS target direction
        f984cdf Upload files to "/"
        c764790 chore: fix readme formatting
        953b006 chore: fix readme formatting
        b537475 Add CarbonStackOS north star and initial appliance model
        dab2792 Initial CarbonStack repository structure

Working tree note:

    carbonstack was clean/aligned with origin/main after v0.6.32B.
    carbonstack-comms was clean/aligned except ignored local artifacts:
        !! internal/protocol/mls/openmls-sidecar/.carbonstack-openmls-sidecar-state/
        !! internal/protocol/mls/openmls-sidecar/target/
    carbonstack-cypher retained the known ignored local DB:
        !! cypher.db
    carbonstack-os was clean/aligned.

## 4. Validation / Evidence State After v0.6.32

v0.6.32A recon validation passed:

    generated reference deterministic check with entries=87;
    runner go test ./...;
    registry missing-nonclaims scan;
    runner doctor;
    required v0.6.27-v0.6.31 docs check;
    toolchain floor inspection;
    Comms go test ./... soft probe;
    Cypher go test ./... soft probe;
    runner help / alias discoverability probe;
    registry lookups for public/front surfaces;
    registry front README candidates list;
    registry release-supported list;
    public surface wording scan;
    manual CLI docs category scan;
    front README marker checks;
    runner README marker checks;
    docs README marker checks;
    no tracked mutation guard.

v0.6.32B docs/help cleanup validation passed:

    preflight repo snapshot;
    generated reference deterministic check with entries=87;
    runner go test ./...;
    registry missing-nonclaims scan;
    runner doctor;
    required v0.6.27-v0.6.31 docs check;
    runner main.go help discoverability patch;
    runner README section and toolchain note patch;
    front README operator section patch;
    docs/230 creation;
    docs README update;
    gofmt;
    carbonstack diff --check;
    generated reference deterministic check unchanged with entries=87;
    runner go test ./... after patch;
    registry missing-nonclaims scan after patch;
    runner help full-validate-release visibility guard;
    public/manual marker checks;
    registry/reference guard;
    cached diff --check;
    cached diff --stat;
    git commit;
    post-commit generated reference deterministic check with entries=87;
    post-commit runner go test ./...;
    post-commit registry missing-nonclaims scan;
    push rc=0;
    final repo snapshot.

Generated command reference remains:

    entries=87

No registry entry was added.

No runner profile was added.

No `full`, `full-validate-release`, or `release-snapshot` behavior changed.

This is acceptable for a public/manual docs + CLI discoverability checkpoint.

Do not overclaim that v0.6.32 performs release-boundary rehearsal, release packaging, deployment validation, or production-security validation.

## 5. v0.6.32 Model Content Summary

### 5.1 New public/manual docs refresh note

New doc:

    docs/230-v0.6.32-public-surface-manual-docs-refresh-v0.md

Core model:

    v0.6.32B is a targeted discoverability cleanup.
    It makes the existing validation/manual surface easier to find before release rehearsal.
    It does not change what CarbonStack proves.

### 5.2 Runner help discoverability

Before v0.6.32B:

    runner switch accepted full-validate-release.
    registry lookup found runner.full-validate-release.
    runner --help did not list full-validate-release.

After v0.6.32B:

    runner --help lists full-validate-release.
    full-validate-release remains exact alias to full.

### 5.3 Current operator-surface pointers now visible

The front README now points to:

    full-validate-release
    full
    release-snapshot
    message-send-dev
    message-inbox-dev
    state-audit-dev
    registry/COMMAND_REFERENCE.v0.md
    docs/229-v0.6.31-adversarial-harness-contract-evidence-matrix-v0.md

### 5.4 Toolchain-floor note

v0.6.32 records the current toolchain-floor ambiguity before release rehearsal:

    runner declares go 1.24;
    Comms declares go 1.26.3;
    Cypher declares go 1.26.3;
    local WSL validation passed with Go 1.24.4;
    release-facing docs must state expected toolchain behavior clearly before v0.7.0 package-root rehearsal.

No `go.mod` files were changed.

### 5.5 What v0.6.32 does not change

v0.6.32 does not add:

    release-boundary rehearsal;
    release package;
    release tag;
    runner profile;
    registry entry;
    generated command reference entry;
    full-runtime-dev;
    adversarial test implementation;
    state-audit mutation;
    Go module floor mutation;
    deployment validation;
    production security claim;
    hostile-server safety claim.

## 6. v0.6.32 Blunders / Continuity Notes

### 6.1 Paste corruption/noise appeared again but validation protected the run

Both v0.6.32A and v0.6.32B logs showed small paste/composer corruption/noise near the temporary Python script invocation. Example pattern:

    python3 /tmp/...py...r(root

Impact:

    no tracked corruption;
    scripts executed;
    validation passed;
    v0.6.32A had no mutation;
    v0.6.32B committed and pushed successfully.

Lesson:

    Continue using paste-safe Python scripts, but continue ignoring the first few corrupted shell echo lines if the script body executed and validation/final snapshots are clean.

### 6.2 v0.6.32B push succeeded normally

Unlike v0.6.31B, the v0.6.32B scripted push succeeded:

    3b9370a..f9d4ac4 main -> main

Impact:

    remote is aligned at `f9d4ac4`.

Lesson:

    First-push auth failure remains a known possible recurring issue, but it did not occur in this checkpoint.

### 6.3 Front README operator section was appended

The v0.6.32B script attempted to insert the operator section before `## Current status`; the marker was not found in the exact expected form, so the section was appended.

Impact:

    marker checks passed;
    front README now contains the needed current operator-surface pointers;
    no functional issue.

Lesson:

    Future README layout cleanup can decide whether to move the section, but there is no need to reopen v0.6.32 solely for placement.

### 6.4 Toolchain floor was documented, not changed

The v0.6.32A recon found:

    runner go.mod = go 1.24;
    Comms go.mod = go 1.26.3;
    Cypher go.mod = go 1.26.3;
    local WSL Go = 1.24.4;
    Comms/Cypher soft tests passed.

v0.6.32B documented this as a release-prep issue but did not change module floors.

Impact:

    release-boundary/package-root rehearsal must explicitly decide or document toolchain behavior.

Lesson:

    Do not silently normalize `go.mod` floors in a docs refresh checkpoint. Handle it during release-boundary/package-root rehearsal preflight or plan.

## 7. Critical Path / Function Updates

New critical doc:

    carbonstack/docs/230-v0.6.32-public-surface-manual-docs-refresh-v0.md

Updated public/operator docs:

    carbonstack/README.md
    carbonstack/docs/README.md
    carbonstack/tools/carbonstack-validate/README.md
    carbonstack/tools/carbonstack-validate/main.go

Critical docs still active:

    carbonstack/docs/223-v0.6.27-vault-security-nosilent-state-model-v0.md
    carbonstack/docs/224-v0.6.28-pq-hybrid-placement-model-v0.md
    carbonstack/docs/225-v0.6.29-bounded-state-substrate-mechanics-v0.md
    carbonstack/docs/226-v0.6.29-state-audit-dev-model-alignment-v0.md
    carbonstack/docs/227-v0.6.30-validation-profile-naming-boundary-v0.md
    carbonstack/docs/228-v0.6.30-full-validate-release-alias-v0.md
    carbonstack/docs/229-v0.6.31-adversarial-harness-contract-evidence-matrix-v0.md
    carbonstack/docs/230-v0.6.32-public-surface-manual-docs-refresh-v0.md
    CarbonStack_Long_Term_Roadmap_v0.6.26_EVERGREEN.pdf

Critical registry/reference paths unchanged:

    carbonstack/registry/commands.v0.yaml
    carbonstack/registry/COMMAND_REFERENCE.v0.md

Critical runner paths changed only for help/docs:

    carbonstack/tools/carbonstack-validate/main.go
    carbonstack/tools/carbonstack-validate/README.md

Critical command/profile surfaces unchanged in semantics:

    cd carbonstack/tools/carbonstack-validate
    go run . --profile full
    go run . --profile full-validate-release
    go run . --profile release-snapshot
    go run . --profile integrated-runtime-dev
    go run . --profile same-state-integrated-dev
    go run . --profile same-state-message-failure-dev
    go run . --profile same-state-message-unsupported-dev
    go run . --profile same-state-message-recipient-failure-dev
    go run . --profile same-state-message-malformed-payload-dev
    go run . --profile same-state-message-replay-classification-dev
    go run . --profile same-state-welcome-join-failure-dev

Critical state-audit surface unchanged:

    cd carbonstack-comms
    go run ./cmd/comms state-audit-dev
    go run ./cmd/comms state-audit-dev --format json

Important behavior after v0.6.32:

    runner --help lists full-validate-release.
    full-validate-release remains exact alias to full.
    front README now points to current operator surfaces.
    docs/230 exists.
    registry is unchanged.
    generated command reference remains entries=87.
    no release rehearsal happened.
    no hostile-server safety is claimed.
    full-runtime-dev remains future-only.

## 8. Future To-Do / Critical Path Forward

### 8.1 Immediate next safest rung

Recommended next checkpoint:

    v0.6.33A = release-boundary/package-root rehearsal preflight

Purpose:

    inspect what a v0.7.0 cumulative pre-alpha engineering release boundary would require before actually rehearsing or packaging it.

The preflight should answer:

    What should the v0.7.0 release boundary include?
    Which repos are included?
    What exact heads are expected?
    What package-root shape is required?
    What files must be excluded?
    What release notes/runbook/checksum/manifest material is needed?
    What validation ladder should be used?
    Is full-validate-release now the preferred release validation command?
    Is toolchain floor clear enough?
    Are README/docs/registry/release claims aligned?
    What remains explicitly nonclaimed?

Do not jump straight to release tagging, asset upload, or package publication.

### 8.2 Near-term sequence

Recommended order:

    v0.6.33A release-boundary/package-root rehearsal preflight
    v0.6.33B release-boundary/package rehearsal plan docs checkpoint
    v0.6.33C or v0.6.34 actual package-root rehearsal script/checkpoint, if plan is clean
    v0.7.0 cumulative pre-alpha engineering release boundary

Bias:

    plan doc first, then rehearsal.

Reason:

    The project is close enough to release-prep that package shape should not be improvised directly.

### 8.3 Deferred until later

Do not do next by default:

    release tag;
    release asset upload;
    public WordPress/project page;
    adversarial runner profile;
    adversarial mega-profile;
    adversarial test implementation;
    full-runtime-dev implementation;
    adding state-audit-dev to a runtime aggregate;
    adding same-state profiles to full/full-validate-release/release-snapshot;
    full deprecation;
    full compatibility warning;
    production vault implementation;
    encryption/secure storage implementation;
    backup/restore implementation;
    PQ/hybrid implementation or experiments until later model-permitted epoch;
    Welcome/KeyPackage replay implementation;
    deployment guide claiming safety;
    production security wording.

### 8.4 Future adversarial implementation posture

Unchanged from v0.6.31:

    first adversarial implementation should be leaf-first, not aggregate-first.
    no aggregate adversarial profile should exist until multiple leaf profiles stabilize.
    future cases must follow docs/229 contract fields.

## 9. Claim Boundaries After v0.6.32

v0.6.32 is **not**:

    a public release;
    v0.7.0;
    release-boundary rehearsal;
    package-root rehearsal;
    release asset creation;
    release tag;
    adversarial harness implementation;
    hostile-server safety;
    malicious relay safety;
    server equivocation detection;
    replay safety;
    Welcome replay handling;
    KeyPackage replay handling;
    metadata privacy;
    metadata authenticity;
    sender authenticity;
    verified identity;
    secure enrollment;
    production secure messaging;
    production E2EE;
    production vault/key storage;
    secure storage;
    encryption-at-rest;
    production backup/restore;
    state rollback safety;
    PQ or hybrid security;
    quantum-safe messaging;
    OpenMLS-native PQ/hybrid support;
    ciphersuite migration;
    algorithm migration;
    Android/CarbonStackOS readiness;
    deployment readiness;
    audit or certification;
    package publication;
    `full-runtime-dev`.

`docs/230` is only:

    a selected public/manual docs and CLI discoverability refresh note.

It is not:

    a release rehearsal plan;
    a release runbook;
    a release package manifest;
    a test runner;
    an adversarial profile;
    hostile-server proof;
    security certification;
    v0.7.0 release-boundary rehearsal.

## 10. Continuity Anchor

Current safe resume point:

    CarbonStack v0.6.32 is complete.
    CarbonStack head is f9d4ac4.
    CarbonStackComms head is 5a61646.
    CarbonStackCypher head is d18a564.
    CarbonStackOS head is 1bbbe52.
    COMMAND_REFERENCE entries = 87.
    full-validate-release exists, maps to full, and appears in runner --help.
    front README points to current operator surfaces.
    docs/230 exists and records the public/manual docs refresh.
    docs/229 still defines the adversarial harness contract/evidence matrix.
    full-runtime-dev does not exist.
    state-audit-dev remains outside full/release-snapshot.
    no adversarial profile exists.
    no hostile-server safety is claimed.
    Next safest work is v0.6.33A release-boundary/package-root rehearsal preflight.

A future assistant should not proceed directly to release tag/package publication. The next checkpoint should inspect and model the v0.7.0 release boundary/package-root rehearsal requirements first.

---

## Appendix A. Carried-forward CarbonStack LogDoc v0.6.31

The following is the full prior v0.6.31 ledger carried forward for continuity. Current v0.6.32 sections above supersede the current-state metadata, repo heads, next-step ordering, and checkpoint summary where they differ.

# CarbonStack LogDoc v0.6.31

**Last updated:** 2026-07-02 local session  
**Owner/operator:** bitcrusher32  
**Assistant/model:** GPT-5.5 Thinking  
**Current phase:** Phase 2F -> **v0.6.31 adversarial harness contract / evidence-matrix checkpoint complete; v0.6.31A completed deep recon over project/toolchain/docs/registry/profile surfaces, and v0.6.31B committed `docs/229-v0.6.31-adversarial-harness-contract-evidence-matrix-v0.md` as a docs-only contract that separates deterministic same-state evidence seeds from future adversarial harness coverage without adding a runner profile, adversarial mega-profile, `full-runtime-dev`, registry/generated-reference mutation, `full` / `full-validate-release` / `release-snapshot` behavior changes, state-audit mutation, PQ, backup/restore, deployment validation, or hostile-server/security claims.**  
**Current public release:** `v0.6.0 State/UX Boundary Validation Pre-Release`  
**Current mainline checkpoint:** `v0.6.31 Adversarial Harness Contract / Evidence Matrix Docs Checkpoint`  
**Release URL:** https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.6.0  
**Update source:** `CarbonStackLogDocV0.6.30`, `CarbonStack_BreakpointV0.6.30`, v0.6.31A deep recon log, v0.6.31B adversarial harness contract docs checkpoint log, and the active `CarbonStack_Long_Term_Roadmap_v0.6.26_EVERGREEN` planning authority.  
**Update purpose:** Preserve the v0.6.31 checkpoint after the v0.6.30 validation naming split made it safe to design adversarial harness contract/evidence boundaries; record the v0.6.31A/B timeline, adversarial contract model, validation evidence, recurring push-auth continuity, repo heads, critical path/function updates, future public-surface/manual-docs inspection ordering, and nonclaims; preserve the full prior Markdown ledger; and keep the JSON breakpoint as a lean current-state handoff.

**Version schema:** `v[scope].[timeline]`. This file is `v0.6.31`, the working branch continuity ledger after the adversarial harness contract / evidence-matrix docs checkpoint in the v0.6.x OpenMLS message-flow integration and boundary-hardening epoch.

**LogDoc V2 usage note:** The Markdown LogDoc is intentionally a dev continuity ledger. It preserves historical v0.6.x timeline data, blunders, validation ladders, command-boundary decisions, and future path constraints. The JSON breakpoint remains a lean current-state handoff and should not try to duplicate the full Markdown timeline.

**Privacy / history-rewrite note:** Exact raw shell logs may contain local usernames, hostnames, email addresses, and old pre-rewrite author identities. This LogDoc avoids repeating local host-derived identity strings except where operationally necessary. Public commit identity remains normalized on public `main`/tag histories. Pre-v0.6.13 commit hashes from older LogDocs and historical docs are now historical **pre-rewrite** references and must not be treated as current Gitea `main`/tag SHAs.

## 0. Executive Summary

v0.6.31 is complete as an **adversarial harness contract / evidence-matrix docs checkpoint**.

It did six bounded things:

1. **Ran v0.6.31A deep recon before writing the contract**
   - v0.6.31A was a no-mutation recon over:
     - repo state;
     - toolchain availability;
     - registry/profile surfaces;
     - release/runtime/state validation boundaries;
     - manual/CLI docs readiness;
     - adversarial-harness vocabulary;
     - deterministic failure-profile evidence seeds.
   - The recon confirmed the v0.6.30 baseline:
     - `carbonstack` at `06f2c34`;
     - `carbonstack-comms` at `5a61646`;
     - `carbonstack-cypher` at `d18a564`;
     - `carbonstack-os` at `1bbbe52`;
     - `COMMAND_REFERENCE.v0.md` current with `entries=87`;
     - runner tests pass;
     - registry missing-nonclaims scan passes;
     - runner doctor passes.
   - Toolchain inventory:
     - Go `1.24.4`;
     - Rust `1.96.0`;
     - Cargo `1.96.0`;
     - SQLite `3.46.1`;
     - Python `3.13.5`;
     - Git `2.47.3`.
   - v0.6.31A did not mutate docs, registry, generated reference, runner code, or any repo.

2. **Mapped adversarial case-family vocabulary**
   - v0.6.31A found all expected adversarial/planning case families present somewhere in docs/code/planning vocabulary:
     - payload mutation;
     - unsupported envelope/content;
     - wrong conversation;
     - wrong recipient/device;
     - replay/duplicate;
     - Welcome/KeyPackage;
     - metadata lies;
     - server equivocation;
     - drop/delay/reorder;
     - routing/membership mutation;
     - state rollback;
     - downgrade/algorithm abuse;
     - ack/drain invariants.
   - Important interpretation:
     - presence of vocabulary is not implementation;
     - deterministic profiles are evidence seeds, not adversarial harness coverage;
     - hostile-server safety is still not claimed.

3. **Mapped manual/CLI docs gaps without folding them into v0.6.31B**
   - v0.6.31A confirmed the roadmap's manual/CLI docs warning is still relevant.
   - Command reference exists, but before v0.7.0 release-boundary/rehearsal the project still needs a public-surface/manual-docs inspection covering:
     - normal-message operator docs;
     - Relay onboarding docs;
     - validation profile docs;
     - state-audit / state-recovery / no-silent docs;
     - deployment-warning docs;
     - public surface wording;
     - front README / command registry / generated reference inspection.
   - Decision:
     - do not block adversarial contract docs on writing all manual docs;
     - use v0.6.31B to define the contract spine;
     - do public surface + command registry + manual/CLI docs inspection next, before release-boundary rehearsal.

4. **Added the v0.6.31 adversarial harness contract / evidence matrix**
   - Added:
     - `docs/229-v0.6.31-adversarial-harness-contract-evidence-matrix-v0.md`
   - Updated:
     - `docs/README.md`
   - CarbonStack commit:
     - `3b9370a docs: define adversarial harness contract`
   - Commit summary:
     - `2 files changed, 699 insertions(+)`
     - `create mode 100644 docs/229-v0.6.31-adversarial-harness-contract-evidence-matrix-v0.md`
   - The contract defines:
     - current validation boundaries;
     - deterministic evidence seeds;
     - required future adversarial case fields;
     - evidence matrix case families;
     - release-validation exclusion rules;
     - live-dev aggregation exclusion rules;
     - registry posture;
     - evidence storage posture;
     - candidate future profile names;
     - manual/CLI docs relationship;
     - required invariants for future implementation;
     - explicit nonclaims.

5. **Preserved the release/runtime/profile boundary**
   - v0.6.31B did not add:
     - runner profile;
     - adversarial mega-profile;
     - registry entry;
     - generated command reference change;
     - `full-runtime-dev`;
     - state-audit mutation;
     - PQ;
     - backup/restore;
     - deployment/security/hostile-server claim.
   - v0.6.31B did not change:
     - `full`;
     - `full-validate-release`;
     - `release-snapshot`;
     - `integrated-runtime-dev`;
     - same-state profiles;
     - `state-audit-dev`;
     - package-root validation semantics.
   - Generated command reference remains:
     - `entries=87`.

6. **Validated and pushed the checkpoint**
   - Baseline validation passed:
     - generated reference deterministic check with `entries=87`;
     - runner `go test ./...`;
     - registry missing-nonclaims scan;
     - runner doctor;
     - required v0.6.27-v0.6.30 docs check.
   - Post-doc validation passed:
     - CarbonStack `git diff --check`;
     - generated reference deterministic check unchanged with `entries=87`;
     - runner `go test ./...`;
     - registry missing-nonclaims scan;
     - registry/reference/runner guard.
   - Post-commit validation passed:
     - generated reference deterministic check with `entries=87`;
     - runner `go test ./...`;
     - registry missing-nonclaims scan.
   - Scripted push hit the recurring Gitea authentication failure:
     - `RESULT carbonstack push rc=128`
   - Manual push succeeded:
     - `06f2c34..3b9370a main -> main`.
   - No final post-manual-push `git status` was printed, but the manual push output shows the remote advanced to the v0.6.31B commit.

Current commits:

    carbonstack        3b9370a docs: define adversarial harness contract
    carbonstack-comms  5a61646 fix: align state audit output with boundary model
    carbonstack-cypher d18a564 chore: restore validated Go module floor
    carbonstack-os     1bbbe52 docs: clarify CarbonStackOS target direction

Current public release remains:

    v0.6.0 State/UX Boundary Validation Pre-Release

This is **not** a public product release.

### 0.1 Carried-forward v0.6.30 baseline summary

v0.6.30 is complete as a **validation profile naming boundary / explicit release-validation alias checkpoint**.

It did seven bounded things:

1. Ran v0.6.30A validation profile naming preflight.
2. Added:
   - `docs/227-v0.6.30-validation-profile-naming-boundary-v0.md`
3. Implemented:
   - `full-validate-release` as an exact alias to current `full` behavior.
4. Added:
   - `docs/228-v0.6.30-full-validate-release-alias-v0.md`
5. Added registry row:
   - `runner.full-validate-release`
6. Regenerated `registry/COMMAND_REFERENCE.v0.md` from `entries=86` to `entries=87`.
7. Preserved:
   - no `full-runtime-dev`;
   - no live-dev aggregation;
   - no same-state inclusion in `full`;
   - no `state-audit-dev` inclusion in `full`;
   - no `release-snapshot` behavior change;
   - no adversarial harness behavior;
   - no deployment/security claim.

v0.6.30 preserved the key boundary:

    full / full-validate-release = release-package validation
    release-snapshot = package-root validation
    integrated-runtime-dev / same-state = live-umbrella dev validation
    state-audit-dev = non-mutating Comms CLI proto-substrate inventory

## 1. Current Project Goal

**Active goal:** Continue the v0.6.x OpenMLS message-flow integration and boundary-hardening epoch from the v0.6.31 adversarial harness contract/evidence-matrix checkpoint toward public surface + command registry + manual/CLI docs inspection, selected docs refresh, v0.7.0 release-boundary/package-root rehearsal planning, and later runtime/adversarial implementation only after explicit model gates.

After v0.6.31, the project has:

- the v0.6.0 public pre-release still live on Gitea;
- generated `registry/COMMAND_REFERENCE.v0.md`;
- deterministic generated-reference freshness enforced by runner tests;
- registry entry count still **87**;
- `full-validate-release` as the explicit release/package-root validation name mapped exactly to current `full`;
- legacy stub commands gated behind explicit opt-in;
- opinionated `message-send-dev` / `message-inbox-dev` wrappers as the recommended dev/pre-alpha normal-message path;
- direct `openmls-*` commands retained as lower-level implementation/proof surfaces;
- Relay onboarding boundaries separated from ordinary message inbox behavior;
- `same-state-integrated-dev` proving Level 4 same-conversation positive-path continuity;
- deterministic normal-message failure/classification profiles;
- Welcome join partial-state safety and `same-state-welcome-join-failure-dev`;
- v0.6.27 state/security/no-silent model;
- v0.6.28 PQ/hybrid placement model;
- v0.6.29 bounded state-substrate / `state-audit-dev` output model alignment;
- v0.6.30 validation naming boundary and explicit release-validation alias;
- v0.6.31 adversarial harness contract/evidence matrix defining future case fields and claim boundaries.

CarbonStack's broader intent remains:

- **CarbonStack** is the secure-communications appliance-stack project and shared doctrine/release authority.
- **CarbonStackComms** is the text-first messaging client and current OpenMLS sidecar / Relay Space integration surface.
- **CarbonStackCypher** is the self-hostable relay/storage server/API stack.
- **CarbonStackOS** is future constrained Android-derived appliance OS work and remains out of current runnable validation packages.
- The `carbonstack` Gitea repo remains the source of truth for release state, release assets, docs, registry, command-boundary classification, and public claims.
- GitHub mirrors, if present, are secondary push mirrors for discoverability/redundancy only.

**North star:**

> Its goal is not to be a better smartphone. Its goal is to stop being a smartphone.

**Core doctrine:**

> Every feature is guilty until it proves it does not add unacceptable parser, network, sensor, identity, or filesystem authority.

**v0.6.x engineering doctrine after v0.6.31:**

> The adversarial harness now has a contract before an implementation. Existing same-state and failure-hardening profiles are evidence seeds, not adversarial coverage. Future cases must declare threat actor, setup state, mutation/action, expected invariant, allowed outputs, forbidden outputs, evidence files, claim granted, nonclaims, and deferred work before they are promoted. Next work should inspect the public surface, command registry, generated reference, and manual/CLI docs gaps before v0.7.0 release-boundary rehearsal.

## 2. Current Public Release and Mainline State

Current public release remains:

    https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.6.0

Release title:

    CarbonStack v0.6.0 State/UX Boundary Validation Pre-Release

Release tag:

    v0.6.0

Current post-rewrite tag target:

    70d4318 v0.6.0

Pre-rewrite v0.6.0 tag target for historical context only:

    7e676cd6ad / 7e676cd scripts: add v0.6.0 package rehearsal

Current mainline checkpoint:

    v0.6.31 Adversarial Harness Contract / Evidence Matrix Docs Checkpoint

Current CarbonStack head:

    3b9370a docs: define adversarial harness contract

Current CarbonStackComms head:

    5a61646 fix: align state audit output with boundary model

v0.6.31 is not a new public release and does not change the release title. It is a post-release adversarial contract / docs-model checkpoint.

Important release-page warning retained:

    Gitea automatically provides Source Code ZIP/TAR.GZ archives for the carbonstack repo at the tag.
    These are not the intended runnable multi-repo validation package.
    Use the attached v0.6.0 package, manifest, checksums, validation freeze, release notes, testing runbook, and LICENSE.

v0.6.31 did not rebuild v0.6.0 release assets and did not change the v0.6.0 release tag.

New v0.6.31 docs/checkpoint:

    docs/229-v0.6.31-adversarial-harness-contract-evidence-matrix-v0.md

No new v0.6.31 registry/profile surface was added.

Generated command reference count remains:

    entries=87

## 3. Current Repo Heads After v0.6.31

Final pushed heads:

    carbonstack        3b9370a docs: define adversarial harness contract
    carbonstack-comms  5a61646 fix: align state audit output with boundary model
    carbonstack-cypher d18a564 chore: restore validated Go module floor
    carbonstack-os     1bbbe52 docs: clarify CarbonStackOS target direction

Recent pushed history:

    carbonstack:
        3b9370a docs: define adversarial harness contract
        06f2c34 feat: add explicit release validation profile alias
        d31537c docs: define validation profile naming boundary
        44e781c docs: record state audit boundary alignment
        2f0be42 docs: define bounded state substrate mechanics
        4a66681 docs: define PQ hybrid placement model

    carbonstack-comms:
        5a61646 fix: align state audit output with boundary model
        2d13c45 fix: clarify load-check provider metadata output
        61b2e51 fix: label message sender metadata as unverified
        b42aa09 fix: make Welcome join state writes atomic
        53779d5 test: cover message wrapper residual edge cases
        462f796 chore: group Comms CLI help by command boundary
        020c07f test: align wrapper smoke with message wrappers
        9391029 feat: add opinionated OpenMLS message wrappers

    carbonstack-cypher:
        d18a564 chore: restore validated Go module floor
        52e3673 docs: make Cypher surface evergreen
        3589dbe docs: define local Cypher validation state contract
        15760c6 docs: record provider OpenMLS join boundary
        09afa0b feat: add Relay Space scoped envelope routes
        384affd feat: add Relay Space HTTP API routes
        cc6e589 feat: add Relay Space DB helpers
        6798d11 feat: add Relay Space schema substrate

    carbonstack-os:
        1bbbe52 docs: clarify CarbonStackOS target direction
        f984cdf Upload files to "/"
        c764790 chore: fix readme formatting
        953b006 chore: fix readme formatting
        b537475 Add CarbonStackOS north star and initial appliance model
        dab2792 Initial CarbonStack repository structure

Working tree note:

    carbonstack was clean before v0.6.31B and clean after the docs commit except the scripted final snapshot happened before manual push and showed `main` ahead of `origin/main` by one commit. Manual push then succeeded, advancing origin/main to `3b9370a`; no post-manual-push `git status` was printed.
    carbonstack-comms was clean/aligned except ignored local artifacts:
        !! internal/protocol/mls/openmls-sidecar/.carbonstack-openmls-sidecar-state/
        !! internal/protocol/mls/openmls-sidecar/target/
    carbonstack-cypher retained the known ignored local DB:
        !! cypher.db
    carbonstack-os was clean/aligned.

## 4. Validation / Evidence State After v0.6.31

v0.6.31A deep recon validation passed:

    generated reference deterministic check with entries=87;
    runner go test ./...;
    registry missing-nonclaims scan;
    runner doctor;
    toolchain inventory;
    required v0.6.27-v0.6.30 docs check;
    registry lookups for release/runtime/state surfaces;
    registry list for dev_only surfaces;
    registry list for release_supported surfaces;
    runner help/profile lookup probes;
    docs inventory;
    manual/CLI docs readiness scan;
    adversarial harness terms scan;
    deterministic failure/adversarial-adjacent code scan;
    claim-boundary/nonclaim vocabulary scan;
    adversarial contract seed extraction;
    docs gap seed table;
    post-recon mutation guard.

v0.6.31B docs/model validation passed:

    preflight repo snapshot;
    generated reference deterministic check with entries=87;
    runner go test ./...;
    registry missing-nonclaims scan;
    runner doctor;
    required v0.6.27-v0.6.30 docs check;
    docs/229 creation;
    docs/README update;
    markdown normalization;
    carbonstack diff --check;
    generated reference deterministic check unchanged with entries=87;
    runner go test ./... after docs patch;
    registry missing-nonclaims scan after docs patch;
    registry/reference/runner guard;
    cached diff --check;
    cached diff --stat;
    git commit;
    post-commit generated reference deterministic check with entries=87;
    post-commit runner go test ./...;
    post-commit registry missing-nonclaims scan;
    manual push after first auth failure.

Generated command reference remains:

    entries=87

No registry entry was added.

No runner profile was added.

No `full`, `full-validate-release`, or `release-snapshot` behavior changed.

This is acceptable for a docs-only adversarial harness contract checkpoint.

Do not overclaim that v0.6.31 implements adversarial testing or hostile-server safety.

## 5. v0.6.31 Model Content Summary

### 5.1 New adversarial harness contract doc

New doc:

    docs/229-v0.6.31-adversarial-harness-contract-evidence-matrix-v0.md

Core model:

    Existing same-state and failure-hardening profiles are evidence seeds.
    They are not adversarial harness coverage by default.
    Future cases must declare threat actor, setup state, mutation/action, expected invariant, allowed outputs, forbidden outputs, evidence files, claim granted, nonclaims, and deferred work.
    Future adversarial work must not silently enter full, full-validate-release, release-snapshot, or full-runtime-dev.
    The next public-surface/manual-docs refresh should use this contract as a claim-boundary spine before v0.7.0 release-boundary rehearsal.

### 5.2 Required future adversarial case fields

Every future adversarial case should declare:

    case_id
    case_family
    title
    status
    introduced_in
    owner_surface
    profile_or_command_surface
    threat_actor
    attacker_capability
    trusted_components
    untrusted_components
    out_of_scope_capabilities
    setup_state
    required_profiles
    required_repos
    required_artifacts
    preconditions
    state_authority_domains
    mutation_or_action
    mutation_location
    mutation_timing
    mutation_mechanism
    expected_invariant
    allowed_outputs
    forbidden_outputs
    ack_drain_expectation
    state_mutation_expectation
    trust_mutation_expectation
    cypher_mutation_expectation
    operator_warning_expectation
    evidence_files
    stdout_markers
    stderr_markers
    db_assertions
    state_assertions
    artifact_assertions
    negative_assertions
    cleanup_expectation
    claim_granted
    claim_not_granted
    explicit_nonclaims
    deferred_work

### 5.3 Evidence matrix planning families

v0.6.31 defines planning rows for:

    normal-message mutation families:
        payload mutation
        unsupported envelope/content
        wrong conversation
        wrong recipient/device
        duplicate/replay

    onboarding artifact families:
        corrupt Welcome
        stale Welcome
        stale KeyPackage
        wrong group onboarding artifact

    Relay/server behavior families:
        drop
        delay
        reorder
        duplicate
        equivocation
        metadata lies

    state and recovery families:
        state rollback
        backup rollback
        silent trust promotion
        Cypher routing/membership mutation
        generated state cleanup

    algorithm / PQ / downgrade families:
        protocol version mismatch
        algorithm downgrade
        PQ/hybrid downgrade
        ciphersuite migration abuse

### 5.4 Existing deterministic evidence seeds

v0.6.31 classifies these as evidence seeds:

    same-state-integrated-dev:
        positive same-state Relay onboarding + normal message proof.

    same-state-message-failure-dev:
        wrong-conversation no-ack/no-drain seed.

    same-state-message-unsupported-dev:
        unsupported content type no-ack/no-drain seed.

    same-state-message-recipient-failure-dev:
        wrong recipient/device no false open/no ack/no drain seed.

    same-state-message-malformed-payload-dev:
        malformed payload no false open/no ack/no drain/no provider mutation/no envelope rewrite seed.

    same-state-message-replay-classification-dev:
        duplicate/replay classification seed, not replay safety.

    same-state-welcome-join-failure-dev:
        corrupt Welcome failure seed, not Welcome/KeyPackage replay safety.

### 5.5 Placement rules

v0.6.31 placement rules:

    Do not add future adversarial cases to full.
    Do not add future adversarial cases to full-validate-release.
    Do not add future adversarial cases to release-snapshot.
    Do not automatically create or join full-runtime-dev.
    Registry presence is classification, not promotion.
    Evidence storage starts as local/dev artifact unless explicitly elevated.
    Evidence must not contain raw secrets.
    Evidence must not be treated as a production audit log.

### 5.6 Candidate future profile names

Candidate-only names, not implemented:

    adversarial-message-payload-mutation-dev
    adversarial-message-metadata-mutation-dev
    adversarial-message-replay-dev
    adversarial-welcome-artifact-dev
    adversarial-keypackage-artifact-dev
    adversarial-relay-metadata-dev
    adversarial-state-rollback-dev
    adversarial-contract-check-dev

Aggregate profile warning:

    Do not add an aggregate adversarial profile first.
    A future aggregate might be considered only after multiple leaf profiles are stable and their claim boundaries are clear.

### 5.7 What v0.6.31 does not change

v0.6.31 does not add:

    implementation;
    runner profile;
    adversarial mega-profile;
    adversarial test execution;
    evidence schema/helper implementation;
    registry entry;
    generated command reference change;
    full-runtime-dev;
    state-audit mutation;
    PQ;
    backup/restore;
    deployment validation;
    production security claim;
    hostile-server safety claim.

## 6. v0.6.31 Blunders / Continuity Notes

### 6.1 Paste corruption still visible but did not break the run

The composer/pasteback showed an odd leading fragment around the temporary Python script text:

    python3 /tmp/carbonstack-v0631b-adversarial-contract-docs.pyCLI docs inspection before release-boundary rehearsal.r(root

Despite that visible paste corruption/noise, the Python script itself executed and completed the checkpoint.

Impact:

    no tracked corruption;
    validation passed;
    commit succeeded;
    manual push succeeded.

Lesson:

    Continue using paste-safe Python scripts for complex operations, but still inspect the beginning of pasted logs for shell-text corruption. Do not trust only the first few lines; rely on validation and final snapshots.

### 6.2 Scripted push auth failure recurred

The scripted CarbonStack push failed:

    RESULT carbonstack push rc=128

Manual retry succeeded:

    06f2c34..3b9370a main -> main

Impact:

    no repo issue;
    final remote advanced to the v0.6.31B commit;
    no post-manual-push status was printed, so use the manual push output as the remote-update evidence.

Lesson:

    Continue treating first-push auth failures as retryable when local validation, commit state, and manual push output are clear.

### 6.3 v0.6.31A confirmed docs refresh belongs after the contract

The roadmap had a manual/CLI docs refresh item before release-boundary/rehearsal. v0.6.31A confirmed the gaps are real, but the agreed sequence is:

    contract first;
    public surface + command registry + manual/CLI docs inspection next;
    selected docs refresh after inspection;
    then release-boundary/package-root rehearsal.

Impact:

    v0.6.31B did not try to solve all manual docs;
    v0.6.32 should not jump straight to release-boundary rehearsal.

Lesson:

    Do not collapse docs gap closure into adversarial harness contract docs. Use the contract as a claim-boundary spine for the docs refresh.

### 6.4 v0.6.31B is intentionally docs-only

v0.6.31B added a 699-line docs checkpoint.

It intentionally did not add:

    adversarial runner profile;
    registry row;
    evidence schema implementation;
    adversarial helper;
    generated command reference mutation;
    release validation mutation.

Impact:

    generated command reference remains entries=87;
    no profile sprawl occurred;
    contract is now available before implementation.

Lesson:

    This is the right posture before v0.7.0 release-boundary work. Future adversarial implementation should happen as leaf profiles only after the public/manual docs and release-boundary inspections are grounded.

## 7. Critical Path / Function Updates

New critical doc:

    carbonstack/docs/229-v0.6.31-adversarial-harness-contract-evidence-matrix-v0.md

Updated doc index:

    carbonstack/docs/README.md

Critical docs still active:

    carbonstack/docs/223-v0.6.27-vault-security-nosilent-state-model-v0.md
    carbonstack/docs/224-v0.6.28-pq-hybrid-placement-model-v0.md
    carbonstack/docs/225-v0.6.29-bounded-state-substrate-mechanics-v0.md
    carbonstack/docs/226-v0.6.29-state-audit-dev-model-alignment-v0.md
    carbonstack/docs/227-v0.6.30-validation-profile-naming-boundary-v0.md
    carbonstack/docs/228-v0.6.30-full-validate-release-alias-v0.md
    carbonstack/docs/229-v0.6.31-adversarial-harness-contract-evidence-matrix-v0.md
    CarbonStack_Long_Term_Roadmap_v0.6.26_EVERGREEN.pdf

Critical registry/reference paths unchanged:

    carbonstack/registry/commands.v0.yaml
    carbonstack/registry/COMMAND_REFERENCE.v0.md

Critical runner paths unchanged:

    carbonstack/tools/carbonstack-validate/main.go
    carbonstack/tools/carbonstack-validate/README.md

Critical command/profile surfaces unchanged:

    cd carbonstack/tools/carbonstack-validate
    go run . --profile full
    go run . --profile full-validate-release
    go run . --profile release-snapshot
    go run . --profile integrated-runtime-dev
    go run . --profile same-state-integrated-dev
    go run . --profile same-state-message-failure-dev
    go run . --profile same-state-message-unsupported-dev
    go run . --profile same-state-message-recipient-failure-dev
    go run . --profile same-state-message-malformed-payload-dev
    go run . --profile same-state-message-replay-classification-dev
    go run . --profile same-state-welcome-join-failure-dev

Critical state-audit surface unchanged:

    cd carbonstack-comms
    go run ./cmd/comms state-audit-dev
    go run ./cmd/comms state-audit-dev --format json

Important behavior after v0.6.31:

    adversarial harness contract exists.
    existing same-state profiles are evidence seeds, not adversarial harness coverage.
    no adversarial profile exists.
    no hostile-server safety is claimed.
    full/full-validate-release/release-snapshot remain release/package-root validation surfaces only.
    full-runtime-dev remains future-only.
    generated command reference entry count remains 87.

## 8. Future To-Do / Critical Path Forward

### 8.1 Immediate next safest rung

Recommended next checkpoint:

    v0.6.32A = public surface + command registry + manual/CLI docs inspection

Purpose:

    inspect public-facing and operator-facing surfaces before v0.7.0 release-boundary/package-root rehearsal.

The inspection should cover:

    front README;
    docs README;
    command registry;
    generated command reference;
    runner README/help;
    release package README/runbook references;
    normal-message operator docs;
    Relay onboarding docs;
    validation-profile docs;
    state-audit / state-recovery / no-silent docs;
    deployment-warning docs;
    public page / WordPress readiness;
    claim-boundary wording;
    command audience/maturity/include_in_front_readme accuracy.

Do not jump straight to release-boundary rehearsal before this inspection.

### 8.2 Near-term sequence

Recommended order:

    v0.6.32A public surface + command registry + manual/CLI docs inspection
    v0.6.32B selected docs refresh / docs gap closure if scope is small enough
    v0.6.33 v0.7.0 release-boundary/package-root rehearsal plan
    v0.7.0 cumulative pre-alpha engineering release boundary

Optional but not immediate:

    v0.6.31C tiny non-executing evidence schema/helper only if a clear need appears.
    Default is to defer implementation until after public/manual docs inspection.

### 8.3 Deferred until later

Do not do next by default:

    adversarial runner profile;
    adversarial mega-profile;
    adversarial test implementation;
    full-runtime-dev implementation;
    adding state-audit-dev to a runtime aggregate;
    adding same-state profiles to full/full-validate-release/release-snapshot;
    full deprecation;
    full compatibility warning;
    production vault implementation;
    encryption/secure storage implementation;
    backup/restore implementation;
    PQ/hybrid implementation or experiments until later model-permitted epoch;
    Welcome/KeyPackage replay implementation;
    v0.7.0 package-root rehearsal before docs surface inspection;
    public WordPress project page before v0.7.0 framing.

### 8.4 Future adversarial implementation posture

First adversarial implementation should be leaf-first, not aggregate-first.

Before implementation, define:

    exact case_id;
    case_family;
    threat_actor;
    setup_state;
    mutation/action;
    expected invariant;
    allowed outputs;
    forbidden outputs;
    ack/drain expectation;
    state/trust/Cypher mutation expectation;
    evidence files;
    cleanup behavior;
    narrow claim;
    explicit nonclaims;
    registry entry posture;
    generated reference behavior.

Candidate first implementation families after docs/release-surface work:

    adversarial-message-payload-mutation-dev
    adversarial-message-replay-dev
    adversarial-welcome-artifact-dev
    adversarial-relay-metadata-dev
    adversarial-state-rollback-dev

No aggregate adversarial profile should exist until multiple leaf profiles stabilize.

## 9. Claim Boundaries After v0.6.31

v0.6.31 is **not**:

    a public release;
    v0.7.0;
    adversarial harness implementation;
    hostile-server safety;
    malicious relay safety;
    server equivocation detection;
    replay safety;
    Welcome replay handling;
    KeyPackage replay handling;
    metadata privacy;
    metadata authenticity;
    sender authenticity;
    verified identity;
    secure enrollment;
    production secure messaging;
    production E2EE;
    production vault/key storage;
    secure storage;
    encryption-at-rest;
    production backup/restore;
    state rollback safety;
    PQ or hybrid security;
    quantum-safe messaging;
    OpenMLS-native PQ/hybrid support;
    ciphersuite migration;
    algorithm migration;
    Android/CarbonStackOS readiness;
    deployment readiness;
    audit or certification;
    release-package inclusion;
    package-root validation change;
    package publication;
    `full-runtime-dev`.

`docs/229` is only:

    the adversarial harness contract / evidence-matrix vocabulary and claim-boundary model.

It is not:

    a test runner;
    an adversarial profile;
    hostile-server proof;
    security certification;
    public/manual docs refresh;
    v0.7.0 release-boundary rehearsal.

## 10. Continuity Anchor

Current safe resume point:

    CarbonStack v0.6.31 is complete.
    CarbonStack head is 3b9370a.
    CarbonStackComms head is 5a61646.
    CarbonStackCypher head is d18a564.
    CarbonStackOS head is 1bbbe52.
    COMMAND_REFERENCE entries = 87.
    full-validate-release exists and maps to full.
    full-runtime-dev does not exist.
    state-audit-dev remains outside full/release-snapshot.
    docs/229 exists and defines the adversarial harness contract/evidence matrix.
    no adversarial profile exists.
    no hostile-server safety is claimed.
    Next safest work is public surface + command registry + manual/CLI docs inspection.

A future assistant should not reopen the adversarial harness contract before the public/manual docs inspection unless a concrete gap appears. The current contract is sufficient to proceed into v0.6.32A public-surface and command-registry/docs inspection.

---

## Appendix A. Carried-forward CarbonStack LogDoc v0.6.30

The following is the full prior v0.6.30 ledger carried forward for continuity. Current v0.6.31 sections above supersede the current-state metadata, repo heads, next-step ordering, and checkpoint summary where they differ.

# CarbonStack LogDoc v0.6.30

**Last updated:** 2026-07-02 local session  
**Owner/operator:** bitcrusher32  
**Assistant/model:** GPT-5.5 Thinking  
**Current phase:** Phase 2F -> **v0.6.30 validation profile naming boundary and explicit release-validation alias checkpoint complete; v0.6.30A completed validation-profile naming preflight, v0.6.30B committed `docs/227-v0.6.30-validation-profile-naming-boundary-v0.md`, and v0.6.30C added `full-validate-release` as an exact alias to current `full` release-package validation behavior without adding `full-runtime-dev`, live-dev aggregation, state-audit inclusion in `full`, `release-snapshot` changes, adversarial harness behavior, deployment validation, or production-security claims.**  
**Current public release:** `v0.6.0 State/UX Boundary Validation Pre-Release`  
**Current mainline checkpoint:** `v0.6.30 Validation Profile Naming / full-validate-release Alias Checkpoint`  
**Release URL:** https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.6.0  
**Update source:** `CarbonStackLogDocV0.6.29`, `CarbonStack_BreakpointV0.6.29`, v0.6.30A validation profile naming preflight, v0.6.30B validation-profile naming docs checkpoint log, v0.6.30C full-validate-release alias implementation/repair log, and the active `CarbonStack_Long_Term_Roadmap_v0.6.26_EVERGREEN` planning authority.  
**Update purpose:** Preserve the v0.6.30 validation naming checkpoint after the v0.6.29 bounded state-substrate work created pressure to separate release-package validation from live-dev runtime aggregation; record the v0.6.30A/B/C timeline, exact alias semantics, registry/generated-reference changes, validation evidence, failed first v0.6.30C registry insertion and repair, repo heads, critical path/function updates, future work ordering, and nonclaims; preserve the full prior Markdown ledger; and keep the JSON breakpoint as a lean current-state handoff.  

**Version schema:** `v[scope].[timeline]`. This file is `v0.6.30`, the working branch continuity ledger after the validation profile naming / explicit release-validation alias checkpoint in the v0.6.x OpenMLS message-flow integration and boundary-hardening epoch.

**LogDoc V2 usage note:** The Markdown LogDoc is intentionally a dev continuity ledger. It preserves historical v0.6.x timeline data, blunders, validation ladders, command-boundary decisions, and future path constraints. The JSON breakpoint remains a lean current-state handoff and should not try to duplicate the full Markdown timeline.

**Privacy / history-rewrite note:** Exact raw shell logs may contain local usernames, hostnames, email addresses, and old pre-rewrite author identities. This LogDoc avoids repeating local host-derived identity strings except where operationally necessary. Public commit identity remains normalized on public `main`/tag histories. Pre-v0.6.13 commit hashes from older LogDocs and historical docs are now historical **pre-rewrite** references and must not be treated as current Gitea `main`/tag SHAs.

## 0. Executive Summary

v0.6.30 is complete as a **validation profile naming boundary / explicit release-validation alias checkpoint**.

It did seven bounded things:

1. **Ran v0.6.30A validation profile naming preflight**
   - v0.6.30A inspected the current validation/profile surface after v0.6.29 made `state-audit-dev` the first non-mutating state-boundary/proto-substrate surface.
   - The preflight confirmed the existing separation:
     - `full` is release-package validation, not live-dev runtime aggregation;
     - `release-snapshot` is package-root/layout/checksum/core validation;
     - `integrated-runtime-dev` is live-umbrella runtime-dev composition and explicitly excluded from `full` / `release-snapshot`;
     - same-state profiles are live-umbrella dev/pre-alpha evidence, not package-root validation;
     - `state-audit-dev` is a future `full-runtime-dev` candidate only, not a `full` / `release-snapshot` input.
   - Decision:
     - model first;
     - then optionally implement only `full-validate-release` as an exact alias to current `full`;
     - defer `full-runtime-dev` implementation because its member set defines a new aggregation surface.

2. **Added the v0.6.30 validation profile naming boundary model**
   - Added:
     - `docs/227-v0.6.30-validation-profile-naming-boundary-v0.md`
   - Updated:
     - `docs/README.md`
   - CarbonStack commit:
     - `d31537c docs: define validation profile naming boundary`
   - Commit summary:
     - `2 files changed, 535 insertions(+)`
     - `create mode 100644 docs/227-v0.6.30-validation-profile-naming-boundary-v0.md`
   - The model defined:
     - `full` remains the compatibility name for existing release-package validation through v0.7.0;
     - `full-validate-release` is the preferred explicit future name for that same release/package-root validation meaning;
     - `release-snapshot` remains package-root/layout/checksum/core validation;
     - `full-runtime-dev` remains future-only;
     - `integrated-runtime-dev` remains the highest current live-dev aggregate;
     - same-state profiles remain live-umbrella dev evidence;
     - `state-audit-dev` remains a future `full-runtime-dev` candidate only;
     - no live-dev profile is added to `full` or `release-snapshot` by this checkpoint.
   - Baseline and post-commit validation passed:
     - generated reference deterministic check;
     - runner `go test ./...`;
     - registry missing-nonclaims scan;
     - runner doctor;
     - registry evidence lookups for key validation surfaces;
     - CarbonStack `git diff --check`;
     - registry/reference/runner guard.
   - First scripted push hit the recurring Gitea authentication failure:
     - `RESULT carbonstack push rc=128`
   - Manual push succeeded:
     - `44e781c..d31537c main -> main`.

3. **Started v0.6.30C narrow alias implementation**
   - Goal:
     - add `full-validate-release` as an exact alias to current `full` behavior.
   - Intended boundary:
     - no `full-runtime-dev`;
     - no live-dev profile promotion;
     - no same-state profile inclusion in `full`;
     - no `state-audit-dev` inclusion in `full`;
     - no `release-snapshot` behavior change;
     - no package-root validation semantic change;
     - no compatibility warning from `full` yet;
     - no deprecation of `full` yet;
     - no hostile-server, deployment, or production-security claim.
   - Initial patch changed:
     - `tools/carbonstack-validate/main.go`;
     - `tools/carbonstack-validate/README.md`;
     - `registry/commands.v0.yaml`;
     - `docs/228-v0.6.30-full-validate-release-alias-v0.md`;
     - `docs/README.md`;
     - regenerated `registry/COMMAND_REFERENCE.v0.md`.

4. **Hit and repaired the v0.6.30C registry YAML insertion blunder**
   - The first v0.6.30C script inserted the new registry row as a malformed top-level YAML item:
     - `- id: runner.full-validate-release`
   - The existing runner test caught it:
     - `TestCommandRegistryHasNoTopLevelEntryIDs`
     - failure line: registry entry was malformed top-level YAML and expected two-space indentation under `entries:`.
   - This was a script insertion bug, not a concept failure.
   - The repair script:
     - indented the malformed `runner.full-validate-release` block under `entries:`;
     - verified no top-level registry entry remained;
     - verified `runner.full-runtime-dev` was not added;
     - verified `full-validate-release` mapped to the same branch as `full`;
     - regenerated the generated command reference.

5. **Completed v0.6.30C full-validate-release alias implementation**
   - CarbonStack commit:
     - `06f2c34 feat: add explicit release validation profile alias`
   - Commit summary:
     - `6 files changed, 161 insertions(+), 3 deletions(-)`
     - `create mode 100644 docs/228-v0.6.30-full-validate-release-alias-v0.md`
   - Changed files:
     - `tools/carbonstack-validate/main.go`
     - `tools/carbonstack-validate/README.md`
     - `registry/commands.v0.yaml`
     - `registry/COMMAND_REFERENCE.v0.md`
     - `docs/228-v0.6.30-full-validate-release-alias-v0.md`
     - `docs/README.md`
   - New accepted runner profile:
     - `go run . --profile full-validate-release`
   - Initial behavior:
     - exact alias to current `full`;
     - current release validation composition remains `release-snapshot` + `local-cypher`.
   - Registry row:
     - `runner.full-validate-release`;
     - audience: `public`;
     - maturity: `release_supported`;
     - validation surface: `release package validation ladder`;
     - related profiles: `full`, `release-snapshot`, `local-cypher`;
     - nonclaims include:
       - not deployment;
       - not local-backbone;
       - not runtime Comms UX;
       - not production security proof;
       - not `full-runtime-dev`;
       - not live-dev aggregation;
       - not adversarial harness;
       - not package publisher.

6. **Updated generated command reference**
   - Before v0.6.30C:
     - `entries=86`
   - After repaired v0.6.30C:
     - `entries=87`
   - Generated reference checks passed:
     - `python3 tools/registry/render-command-reference.py --check`
   - The entry-count increase is expected because this is the first v0.6.30 registry/profile addition.

7. **Preserved validation and claim boundaries**
   - Validation after repair passed:
     - registry indentation guard;
     - exact alias branch guard;
     - no `full-runtime-dev` registry guard;
     - generated reference deterministic check with `entries=87`;
     - generated reference entry-count guard;
     - CarbonStack `git diff --check`;
     - runner `go test ./...`;
     - registry lookup for `runner.full`;
     - registry lookup for `runner.full-validate-release`;
     - registry missing-nonclaims scan;
     - cached diff check;
     - post-commit generated reference check;
     - post-commit runner tests;
     - post-commit registry lookup for `runner.full-validate-release`;
     - push rc=0.
   - Final CarbonStack snapshot:
     - clean/aligned with `origin/main`;
     - head `06f2c34`.
   - Existing other-repo heads remain unchanged:
     - CarbonStackComms `5a61646`;
     - CarbonStackCypher `d18a564`;
     - CarbonStackOS `1bbbe52`.
   - Known ignored local artifacts remain acceptable:
     - Comms sidecar generated state and `target/`;
     - Cypher `cypher.db`.

Current commits:

    carbonstack        06f2c34 feat: add explicit release validation profile alias
    carbonstack-comms  5a61646 fix: align state audit output with boundary model
    carbonstack-cypher d18a564 chore: restore validated Go module floor
    carbonstack-os     1bbbe52 docs: clarify CarbonStackOS target direction

Current public release remains:

    v0.6.0 State/UX Boundary Validation Pre-Release

This is **not** a public product release.

### 0.1 Carried-forward v0.6.29 baseline summary

v0.6.29 is complete as a **bounded state substrate / state-audit-dev model-alignment checkpoint**.

It did six bounded things:

1. Ran v0.6.29A bounded state substrate mechanics preflight using the accepted v0.6.27 state/security/no-silent model and v0.6.28 PQ/hybrid placement model as controlling inputs.
2. Discovered that a new first substrate command was not needed because `comms.state-audit-dev` already existed as a registered non-mutating state-domain inventory surface.
3. Added:
   - `docs/225-v0.6.29-bounded-state-substrate-mechanics-v0.md`
4. Ran targeted `state-audit-dev` inspection and decided to patch the existing command rather than adding `state-boundary-check-dev`, a wrapper, or `--json`.
5. Aligned `state-audit-dev` text and JSON output with the bounded state-substrate model by adding report-level and domain-level model fields.
6. Added:
   - `docs/226-v0.6.29-state-audit-dev-model-alignment-v0.md`

v0.6.29 preserved the key boundary:

    state-audit-dev is the current proto-substrate / non-mutating inventory precursor.
    It is not a vault.
    It is not secure storage.
    It does not backup/restore.
    It does not mutate trust.
    It does not mutate Cypher.
    It reserves PQ tag language as not implemented.
    It does not join full or release-snapshot.

## 1. Current Project Goal

**Active goal:** Continue the v0.6.x OpenMLS message-flow integration and boundary-hardening epoch from the v0.6.30 validation profile naming / explicit release-validation alias checkpoint toward adversarial harness contract/evidence matrix design, v0.7.0 release-boundary/package-root planning, and later full-runtime-dev aggregation only after explicit member-set modeling.

After v0.6.30, the project has:

- the v0.6.0 public pre-release still live on Gitea;
- generated `registry/COMMAND_REFERENCE.v0.md`;
- deterministic generated-reference freshness enforced by runner tests;
- registry entry count now **87** after adding `runner.full-validate-release`;
- legacy stub commands gated behind explicit opt-in;
- opinionated `message-send-dev` / `message-inbox-dev` wrappers as the recommended dev/pre-alpha normal-message path;
- direct `openmls-*` commands retained as lower-level implementation/proof surfaces;
- Relay onboarding boundaries separated from ordinary message inbox behavior;
- `same-state-integrated-dev`, proving Level 4 same conversation positive-path continuity;
- deterministic normal-message failure/classification profiles;
- Welcome join partial-state safety and `same-state-welcome-join-failure-dev`;
- v0.6.27 state/security/no-silent model;
- v0.6.28 PQ/hybrid placement model;
- v0.6.29 bounded state-substrate / `state-audit-dev` output model alignment;
- v0.6.30 validation naming boundary;
- `full-validate-release` as the explicit release/package-root validation name mapped exactly to current `full`.

CarbonStack's broader intent remains:

- **CarbonStack** is the secure-communications appliance-stack project and shared doctrine/release authority.
- **CarbonStackComms** is the text-first messaging client and current OpenMLS sidecar / Relay Space integration surface.
- **CarbonStackCypher** is the self-hostable relay/storage server/API stack.
- **CarbonStackOS** is future constrained Android-derived appliance OS work and remains out of current runnable validation packages.
- The `carbonstack` Gitea repo remains the source of truth for release state, release assets, docs, registry, command-boundary classification, and public claims.
- GitHub mirrors, if present, are secondary push mirrors for discoverability/redundancy only.

**North star:**

> Its goal is not to be a better smartphone. Its goal is to stop being a smartphone.

**Core doctrine:**

> Every feature is guilty until it proves it does not add unacceptable parser, network, sensor, identity, or filesystem authority.

**v0.6.x engineering doctrine after v0.6.30:**

> The release-validation and live-dev runtime-validation surfaces are now separated by name and registry boundary. `full-validate-release` is the preferred explicit release-validation name. `full` remains the compatibility name. `full-runtime-dev` remains future-only until its member set and semantics are deliberately modeled. Next work should design the adversarial harness contract/evidence matrix without claiming hostile-server safety.

## 2. Current Public Release and Mainline State

Current public release remains:

    https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.6.0

Release title:

    CarbonStack v0.6.0 State/UX Boundary Validation Pre-Release

Release tag:

    v0.6.0

Current post-rewrite tag target:

    70d4318 v0.6.0

Pre-rewrite v0.6.0 tag target for historical context only:

    7e676cd6ad / 7e676cd scripts: add v0.6.0 package rehearsal

Current mainline checkpoint:

    v0.6.30 Validation Profile Naming / full-validate-release Alias Checkpoint

Current CarbonStack head:

    06f2c34 feat: add explicit release validation profile alias

Current CarbonStackComms head:

    5a61646 fix: align state audit output with boundary model

v0.6.30 is not a new public release and does not change the release title. It is a post-release validation naming / command-boundary checkpoint.

Important release-page warning retained:

    Gitea automatically provides Source Code ZIP/TAR.GZ archives for the carbonstack repo at the tag.
    These are not the intended runnable multi-repo validation package.
    Use the attached v0.6.0 package, manifest, checksums, validation freeze, release notes, testing runbook, and LICENSE.

v0.6.30 did not rebuild v0.6.0 release assets and did not change the v0.6.0 release tag.

New v0.6.30 docs/checkpoints:

    docs/227-v0.6.30-validation-profile-naming-boundary-v0.md
    docs/228-v0.6.30-full-validate-release-alias-v0.md

New v0.6.30 registry/profile surface:

    runner.full-validate-release
    go run . --profile full-validate-release

New generated command reference count:

    entries=87

## 3. Current Repo Heads After v0.6.30

Final pushed heads:

    carbonstack        06f2c34 feat: add explicit release validation profile alias
    carbonstack-comms  5a61646 fix: align state audit output with boundary model
    carbonstack-cypher d18a564 chore: restore validated Go module floor
    carbonstack-os     1bbbe52 docs: clarify CarbonStackOS target direction

Recent pushed history:

    carbonstack:
        06f2c34 feat: add explicit release validation profile alias
        d31537c docs: define validation profile naming boundary
        44e781c docs: record state audit boundary alignment
        2f0be42 docs: define bounded state substrate mechanics
        4a66681 docs: define PQ hybrid placement model

    carbonstack-comms:
        5a61646 fix: align state audit output with boundary model
        2d13c45 fix: clarify load-check provider metadata output
        61b2e51 fix: label message sender metadata as unverified
        b42aa09 fix: make Welcome join state writes atomic

    carbonstack-cypher:
        d18a564 chore: restore validated Go module floor
        52e3673 docs: make Cypher surface evergreen
        3589dbe docs: define local Cypher validation state contract

    carbonstack-os:
        1bbbe52 docs: clarify CarbonStackOS target direction
        f984cdf Upload files to "/"
        c764790 chore: fix readme formatting

Working tree note:

    carbonstack was clean/aligned in final snapshot.
    carbonstack-comms was previously clean/aligned except ignored local artifacts:
        !! internal/protocol/mls/openmls-sidecar/.carbonstack-openmls-sidecar-state/
        !! internal/protocol/mls/openmls-sidecar/target/
    carbonstack-cypher retained the known ignored local DB:
        !! cypher.db
    carbonstack-os was clean/aligned.

## 4. Validation / Evidence State After v0.6.30

v0.6.30B docs/model validation passed:

    generated reference deterministic check with entries=86;
    runner go test ./...;
    registry missing-nonclaims scan;
    runner doctor;
    registry evidence lookups for:
        runner.full;
        runner.release-snapshot;
        runner.integrated-runtime-dev;
        runner.dev-runtime-openmls;
        runner.dev-runtime-openmls-wrappers;
        runner.relay-openmls-join-dev;
        runner.local-cypher;
        runner.core;
        comms.state-audit-dev;
    carbonstack diff --check;
    generated reference deterministic check unchanged;
    runner go test ./... after docs patch;
    registry/reference/runner guard;
    cached diff --check;
    post-commit generated reference check;
    post-commit runner tests;
    manual push after first auth failure.

v0.6.30C initial implementation attempt failed correctly:

    runner go test ./... failed;
    TestCommandRegistryHasNoTopLevelEntryIDs caught a malformed top-level registry row:
        - id: runner.full-validate-release
    expected two-space indentation under entries:
        "  - id: runner.full-validate-release"

v0.6.30C repair validation passed:

    registry indentation guard;
    no full-runtime-dev registry guard;
    exact alias branch guard;
    render command reference;
    generated reference deterministic check with entries=87;
    generated reference entry count guard;
    carbonstack diff --check;
    runner go test ./...;
    registry lookup runner.full;
    registry lookup runner.full-validate-release;
    registry missing-nonclaims scan;
    cached diff --check;
    post-commit generated reference check;
    post-commit runner go test ./...;
    post-commit registry lookup full-validate-release;
    push rc=0.

Generated command reference now:

    entries=87

This is expected and intentional.

## 5. v0.6.30 Model / Implementation Content Summary

### 5.1 Validation profile naming boundary

New doc:

    docs/227-v0.6.30-validation-profile-naming-boundary-v0.md

Core decisions:

    full remains the compatibility name for existing release-package validation through v0.7.0.
    full-validate-release is the preferred explicit future name for that same release/package-root validation meaning.
    release-snapshot remains package-root/layout/checksum/core validation.
    full-runtime-dev is future-only for now.
    integrated-runtime-dev remains the highest current live-dev aggregate.
    same-state profiles remain live-umbrella dev evidence.
    state-audit-dev is a future full-runtime-dev candidate only.
    no live-dev profile is added to full or release-snapshot by this checkpoint.
    no runner implementation happened in v0.6.30B.

### 5.2 full-validate-release alias implementation

New doc:

    docs/228-v0.6.30-full-validate-release-alias-v0.md

Changed files:

    carbonstack/tools/carbonstack-validate/main.go
    carbonstack/tools/carbonstack-validate/README.md
    carbonstack/registry/commands.v0.yaml
    carbonstack/registry/COMMAND_REFERENCE.v0.md
    carbonstack/docs/228-v0.6.30-full-validate-release-alias-v0.md
    carbonstack/docs/README.md

New command/profile:

    go run . --profile full-validate-release

Semantics:

    exact alias to current full behavior

Current `full` / `full-validate-release` composition:

    release-snapshot
    local-cypher

`full` remains available.

`full` remains the compatibility name for release-package validation through v0.7.0.

`full-validate-release` is now the preferred explicit release-validation name.

### 5.3 What v0.6.30 does not change

v0.6.30 does not add:

    full-runtime-dev;
    live-dev runtime aggregation;
    same-state profile inclusion in full;
    state-audit-dev inclusion in full;
    release-snapshot behavior changes;
    package-root validation semantic changes;
    full deprecation;
    full compatibility warning;
    adversarial harness implementation;
    hostile-server safety;
    deployment validation;
    production secure messaging;
    production E2EE;
    production security proof.

## 6. v0.6.30 Blunders / Continuity Notes

### 6.1 First v0.6.30C registry insertion produced malformed YAML

The initial v0.6.30C alias script inserted `runner.full-validate-release` as a top-level YAML list item instead of as a two-space-indented `entries:` child.

Bad form:

    - id: runner.full-validate-release

Correct form:

      - id: runner.full-validate-release

The runner test `TestCommandRegistryHasNoTopLevelEntryIDs` caught this and failed the run.

Impact:

    no commit happened before the failure;
    working tree was dirty but recoverable;
    no reset was needed;
    repair script corrected the indentation and continued validation.

Lesson:

    When mutating `registry/commands.v0.yaml`, never insert registry rows using a dedented block unless the script deliberately preserves the existing `entries:` indentation. Keep the registry format test as mandatory.

### 6.2 Generated command reference count changed from 86 to 87

During the failed first alias attempt, the renderer still reported `entries=86`, because the malformed row was not properly under `entries:`.

After repair, the generated reference correctly reported:

    entries=87

Impact:

    the entry-count change is expected after adding `runner.full-validate-release`;
    future checks should expect 87 until another registry row is deliberately added.

Lesson:

    Renderer entry count is useful but not sufficient. The registry YAML format guard and runner tests are also required.

### 6.3 v0.6.30B first push auth failure

The v0.6.30B scripted push failed with the recurring Gitea authentication issue.

Manual push succeeded:

    44e781c..d31537c main -> main

Impact:

    no repo issue;
    final remote was updated;
    branch aligned with origin/main after manual retry.

Lesson:

    Continue treating first-push auth rc=128/fatal auth as retryable when local commit state and final push snapshot are clear.

### 6.4 Paste corruption remains a workflow hazard

Earlier v0.6.30 script attempts showed command-paste corruption from long shell blocks.

Impact:

    Python script files are safer than large nested Bash paste blocks;
    future scripts should prefer paste-safe Python or very small shell fragments.

Lesson:

    Avoid long Bash blocks with nested heredocs and shell functions in the ChatGPT textbox. Use Python scripts written to `/tmp/...py` for complex recon/patch workflows.

### 6.5 `full` was not deprecated and does not warn yet

v0.6.30C intentionally did not add a compatibility warning to `full`.

Impact:

    `full` continues to work as before;
    `full-validate-release` exists as the preferred explicit name;
    a future warning can be added later after enough continuity exists.

Lesson:

    Avoid changing operator-facing behavior more than needed in the same patch that introduces a safer name.

## 7. Critical Path / Function Updates

New critical docs:

    carbonstack/docs/227-v0.6.30-validation-profile-naming-boundary-v0.md
    carbonstack/docs/228-v0.6.30-full-validate-release-alias-v0.md

Updated doc index:

    carbonstack/docs/README.md

Updated registry/reference paths:

    carbonstack/registry/commands.v0.yaml
    carbonstack/registry/COMMAND_REFERENCE.v0.md

Updated runner paths:

    carbonstack/tools/carbonstack-validate/main.go
    carbonstack/tools/carbonstack-validate/README.md

New critical command/profile surface:

    cd carbonstack/tools/carbonstack-validate
    go run . --profile full-validate-release

Existing compatibility command retained:

    cd carbonstack/tools/carbonstack-validate
    go run . --profile full

Important behavior after v0.6.30:

    full-validate-release maps to the same branch as full.
    full remains available.
    release-snapshot behavior is unchanged.
    full-runtime-dev is not added.
    live-dev profiles are not added to full/release-snapshot.
    same-state profiles are not added to full/release-snapshot.
    state-audit-dev is not added to full/release-snapshot.
    generated command reference entry count is now 87.

Existing critical docs still active:

    carbonstack/docs/223-v0.6.27-vault-security-nosilent-state-model-v0.md
    carbonstack/docs/224-v0.6.28-pq-hybrid-placement-model-v0.md
    carbonstack/docs/225-v0.6.29-bounded-state-substrate-mechanics-v0.md
    carbonstack/docs/226-v0.6.29-state-audit-dev-model-alignment-v0.md
    carbonstack/docs/227-v0.6.30-validation-profile-naming-boundary-v0.md
    carbonstack/docs/228-v0.6.30-full-validate-release-alias-v0.md
    CarbonStack_Long_Term_Roadmap_v0.6.26_EVERGREEN.pdf

## 8. Future To-Do / Critical Path Forward

### 8.1 Immediate next safest rung

Recommended next checkpoint:

    v0.6.31A = adversarial harness contract/evidence matrix preflight

Purpose:

    Design the harness contract without claiming hostile-server safety.

The preflight should define:

    case families;
    threat actors;
    setup state;
    mutation/action;
    expected invariant;
    allowed outputs;
    forbidden outputs;
    claim granted;
    registry/profile placement;
    evidence storage;
    what is deferred.

Do not implement a mega-profile first.

Do not add hostile-server language.

Do not add adversarial claims to `full`, `full-validate-release`, or `release-snapshot`.

### 8.2 Near-term sequence

Recommended order:

    v0.6.31A adversarial harness contract/evidence matrix preflight
    v0.6.31B adversarial harness contract docs checkpoint
    optional v0.6.31C if a tiny non-executing evidence schema/helper is justified
    v0.6.32 v0.7.0 release-boundary/package-root rehearsal plan
    v0.7.0 cumulative pre-alpha engineering release

### 8.3 Deferred until later

Do not do next by default:

    full-runtime-dev implementation;
    full deprecation;
    full compatibility warning;
    adding state-audit-dev to an aggregate;
    adding same-state profiles to full;
    adding live-dev runtime profiles to release validation;
    production vault;
    backup/restore implementation;
    PQ/hybrid implementation;
    Welcome/KeyPackage replay implementation;
    public WordPress/project page before v0.7.0;
    Android/CarbonStackOS readiness work.

### 8.4 Future full-runtime-dev posture

`full-runtime-dev` is still useful, but not immediate.

Before implementation, define:

    exact member list;
    run order;
    artifact cleanup behavior;
    state-audit inclusion;
    failure semantics;
    time cost;
    registry row;
    generated reference behavior;
    nonclaims;
    difference from full-validate-release.

Candidate future inputs remain:

    integrated-runtime-dev;
    same-state-integrated-dev;
    selected same-state failure profiles;
    state-audit-dev.

## 9. Claim Boundaries After v0.6.30

v0.6.30 is **not**:

    a public release;
    v0.7.0;
    production secure messaging;
    production E2EE;
    hostile-server safety;
    adversarial harness validation;
    replay safety;
    Welcome replay handling;
    KeyPackage replay handling;
    metadata privacy;
    metadata authenticity;
    sender authenticity;
    verified sender identity;
    secure enrollment;
    production vault/key storage;
    secure storage;
    encryption-at-rest;
    production backup/restore;
    PQ or hybrid security;
    quantum-safe messaging;
    OpenMLS-native PQ/hybrid support;
    ciphersuite migration;
    algorithm migration;
    Android/CarbonStackOS readiness;
    deployability readiness;
    audit or certification;
    local-backbone;
    mature messenger UX;
    general-public UX;
    package publisher;
    `full-runtime-dev`.

`full-validate-release` is only:

    the preferred explicit name for current release-package validation behavior.

It is not:

    runtime-dev aggregation;
    deployment validation;
    production security proof;
    adversarial harness proof;
    live-dev profile bundle.

## 10. Continuity Anchor

Current safe resume point:

    CarbonStack v0.6.30 is complete.
    CarbonStack head is 06f2c34.
    CarbonStackComms head is 5a61646.
    CarbonStackCypher head is d18a564.
    CarbonStackOS head is 1bbbe52.
    COMMAND_REFERENCE entries = 87.
    full-validate-release exists and maps to full.
    full-runtime-dev does not exist.
    state-audit-dev remains outside full/release-snapshot.
    Next safest work is adversarial harness contract/evidence matrix preflight.

A future assistant should not reopen `full` naming unless a concrete issue appears. The naming boundary has enough evidence to proceed into adversarial harness contract design.

---

## Appendix A. Carried-forward CarbonStack LogDoc v0.6.29

The following is the full prior v0.6.29 ledger carried forward for continuity. Current v0.6.30 sections above supersede the current-state metadata, repo heads, command-reference count, next-step ordering, and checkpoint summary where they differ.

# CarbonStack LogDoc v0.6.29

**Last updated:** 2026-06-21 local session  
**Owner/operator:** bitcrusher32  
**Assistant/model:** GPT-5.5 Thinking  
**Current phase:** Phase 2F -> **v0.6.29 bounded state substrate / state-audit-dev model-alignment checkpoint complete; v0.6.29A completed bounded substrate mechanics preflight, v0.6.29B committed `docs/225-v0.6.29-bounded-state-substrate-mechanics-v0.md`, and v0.6.29C aligned existing `state-audit-dev` output with the bounded state substrate model without adding a new command, wrapper, registry entry, `full`, `release-snapshot`, vault, backup/restore, or PQ behavior.**  
**Current public release:** `v0.6.0 State/UX Boundary Validation Pre-Release`  
**Current mainline checkpoint:** `v0.6.29 Bounded State Substrate / State-Audit Model Alignment Checkpoint`  
**Release URL:** https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.6.0  
**Update source:** `CarbonStackLogDocV0.6.28`, `CarbonStack_BreakpointV0.6.28`, v0.6.29A bounded state substrate preflight log, v0.6.29B bounded substrate mechanics docs checkpoint log, v0.6.29C state-audit-dev inspection/model-alignment logs, and the active `CarbonStack_Long_Term_Roadmap_v0.6.26_EVERGREEN` planning authority.  
**Update purpose:** Preserve the v0.6.29 bounded state substrate checkpoint after the v0.6.27 state/no-silent model and v0.6.28 PQ/hybrid placement model were used to define the first non-mutating state-boundary substrate surface; record the v0.6.29A/B/C timeline, state-audit-dev output alignment, validation evidence, blunders/continuity notes, repo heads, critical path/function updates, future work ordering, and nonclaims; preserve the full prior Markdown ledger; and keep the JSON breakpoint as a lean current-state handoff.  

**Version schema:** `v[scope].[timeline]`. This file is `v0.6.29`, the working branch continuity ledger after the bounded state substrate / state-audit-dev model-alignment checkpoint in the v0.6.x OpenMLS message-flow integration and boundary-hardening epoch.

**LogDoc V2 usage note:** The Markdown LogDoc is intentionally a dev continuity ledger. It preserves historical v0.6.x timeline data, blunders, validation ladders, command-boundary decisions, and future path constraints. The JSON breakpoint remains a lean current-state handoff and should not try to duplicate the full Markdown timeline.

**Privacy / history-rewrite note:** Exact raw shell logs may contain local usernames, hostnames, email addresses, and old pre-rewrite author identities. This LogDoc avoids repeating local host-derived identity strings except where operationally necessary. Public commit identity remains normalized on public `main`/tag histories. Pre-v0.6.13 commit hashes from older LogDocs and historical docs are now historical **pre-rewrite** references and must not be treated as current Gitea `main`/tag SHAs.

## 0. Executive Summary

v0.6.29 is complete as a **bounded state substrate / state-audit-dev model-alignment checkpoint**.

It did six bounded things:

1. **Ran v0.6.29A bounded state substrate mechanics preflight**
   - v0.6.29A used the accepted v0.6.27 state/security/no-silent model and the v0.6.28 PQ/hybrid placement model as controlling inputs.
   - The preflight explicitly preserved the non-goal boundary:
     - no implementation in the recon script;
     - no encryption;
     - no secure storage;
     - no production vault;
     - no backup/restore implementation;
     - no PQ implementation;
     - no adversarial harness implementation;
     - no registry mutation;
     - no generated command reference mutation;
     - no `full` / `release-snapshot` mutation;
     - no public release;
     - no secure-use claim.
   - It confirmed the initial substrate posture:
     - prefer `state-*` naming for command/output where possible;
     - reserve `vault-substrate-dev` mostly for docs/internal lane language;
     - first implementation, if simple enough, should be inventory/check only;
     - Cypher DB should be inventory-only at first and not managed/restored/reconciled;
     - generated sidecar roots may be inspected with hard dev/test labels;
     - PQ tags should be reserved conceptually but not implemented;
     - `full` / `full-validate-release` / `full-runtime-dev` planning should come after substrate mechanics and before adversarial harness contract work.
   - Preflight validation passed:
     - generated reference deterministic check;
     - runner `go test ./...`;
     - registry missing-nonclaims scan;
     - runner doctor;
     - sidecar `cargo test --quiet` warning-free;
     - Comms `go test ./...`;
     - Cypher `go test ./...`.
   - Focused smoke profiles passed:
     - `same-state-integrated-dev`;
     - `same-state-message-malformed-payload-dev`;
     - `same-state-message-replay-classification-dev`.

2. **Discovered that a new first substrate command was not needed**
   - v0.6.29A found an existing registered surface:
     - `comms.state-audit-dev`
     - command: `go run ./cmd/comms state-audit-dev`
   - Existing registry description already framed it as:
     - non-mutating local state-domain inventory;
     - Comms, trust/candidate state, OpenMLS sidecar generated state, build output, and local Cypher DB boundary reporting;
     - not vault encryption;
     - not production key storage;
     - not recovery;
     - not deletion or cleanup;
     - not trust verification;
     - not local-backbone;
     - not production secure messaging.
   - This changed the implementation direction:
     - do **not** create `state-boundary-check-dev` by default;
     - keep `state-audit-dev` as the implementation surface;
     - use docs to define the broader bounded substrate mechanics;
     - patch `state-audit-dev` only if inspection showed concrete model-alignment gaps.

3. **Added the v0.6.29 bounded state substrate mechanics model**
   - Added:
     - `docs/225-v0.6.29-bounded-state-substrate-mechanics-v0.md`
   - Updated:
     - `docs/README.md`
   - CarbonStack commit:
     - `2f0be42 docs: define bounded state substrate mechanics`
   - Commit summary:
     - `2 files changed, 752 insertions(+)`
     - `create mode 100644 docs/225-v0.6.29-bounded-state-substrate-mechanics-v0.md`
   - The model defined:
     - the first substrate is not a vault;
     - the first substrate is not secure storage;
     - the first substrate is non-mutating inventory/checking;
     - existing `state-audit-dev` is the current proto-substrate surface;
     - `state-*` naming is preferred for public/dev-facing command/output wording;
     - `vault-substrate-dev` remains internal/docs wording;
     - Cypher is inventory-only at first;
     - generated sidecar roots are inspectable only with hard dev/test labels;
     - PQ/migration/downgrade tags are reserved as future requirements, not implemented;
     - registry mutation is deferred unless code changes require it;
     - `full` naming comes after substrate mechanics;
     - adversarial harness contract comes after `full` naming.
   - Scripted push initially hit recurring Gitea authentication failure:
     - `RESULT carbonstack push rc=128`
   - Manual push succeeded:
     - `4a66681..2f0be42 main -> main`.

4. **Ran v0.6.29C targeted `state-audit-dev` inspection**
   - The scout confirmed current `state-audit-dev` behavior:
     - text output existed;
     - machine-readable output existed through `--format json`;
     - unknown format rejection existed;
     - raw state contents were not printed;
     - mutation remained disallowed;
     - domain reporting already covered:
       - Comms state;
       - trust store;
       - trust history;
       - candidate store;
       - sidecar generated state;
       - sidecar build output;
       - local Cypher DB.
   - The scout also showed the command did not yet explicitly print the full v0.6.29 model vocabulary:
     - `state_boundary_model_version`;
     - `state_boundary_role`;
     - `proto_substrate`;
     - `pq_tags_reserved_not_implemented`;
     - `authority_class`;
     - `sensitivity_class`;
     - `no_silent_rule`;
     - `boundary_warning`;
     - `cypher_inventory_only`.
   - Decision:
     - patch the existing command;
     - do not add a new command;
     - do not add `--json`;
     - preserve `--format json`.

5. **Aligned `state-audit-dev` output with the bounded substrate model**
   - CarbonStackComms commit:
     - `5a61646 fix: align state audit output with boundary model`
   - Changed files:
     - `internal/state/state_audit.go`
     - `internal/app/state_audit_dev.go`
     - `internal/app/state_audit_dev_test.go`
   - Commit summary:
     - `3 files changed, 158 insertions(+), 18 deletions(-)`
   - `state-audit-dev` now prints report-level fields:
     - `state_boundary_model_version: v0.6.29-state-boundary-v0`
     - `state_boundary_role: proto_substrate_inventory_check`
     - `proto_substrate: true`
     - `pq_tags_reserved_not_implemented: true`
   - Domain output now includes:
     - `authority_class`;
     - `sensitivity_class`;
     - `no_silent_rule`;
     - `boundary_warning`;
     - `cypher_inventory_only`.
   - The machine-readable path remains:
     - `--format json`
   - No `--json` alias was added.
   - The patch kept:
     - `mutation_allowed: false`;
     - `raw_secret_contents_printed: false`;
     - no raw local state leakage;
     - no trust mutation;
     - no Cypher mutation;
     - no backup/restore behavior;
     - no PQ behavior.

6. **Added the v0.6.29C result note and preserved boundaries**
   - Added:
     - `docs/226-v0.6.29-state-audit-dev-model-alignment-v0.md`
   - Updated:
     - `docs/README.md`
   - CarbonStack commit:
     - `44e781c docs: record state audit boundary alignment`
   - Commit summary:
     - `2 files changed, 91 insertions(+)`
     - `create mode 100644 docs/226-v0.6.29-state-audit-dev-model-alignment-v0.md`
   - Validation passed:
     - focused Comms state/app tests;
     - state-audit-dev text output model-field assertions;
     - state-audit-dev JSON output parse/model-field assertions;
     - full Comms `go test ./...`;
     - CarbonStack `git diff --check`;
     - generated reference deterministic check;
     - runner `go test ./...`;
     - registry missing-nonclaims scan;
     - registry/reference/runner guard.
   - Push succeeded for:
     - `carbonstack-comms`: `2d13c45..5a61646 main -> main`;
     - `carbonstack`: `2f0be42..44e781c main -> main`.
   - Final working-tree snapshot:
     - CarbonStack clean and aligned with `origin/main`;
     - CarbonStackComms clean/aligned except ignored local sidecar generated state and build output;
     - CarbonStackCypher clean/aligned except known ignored `cypher.db`;
     - CarbonStackOS clean/aligned.

Current commits:

    carbonstack        44e781c docs: record state audit boundary alignment
    carbonstack-comms  5a61646 fix: align state audit output with boundary model
    carbonstack-cypher d18a564 chore: restore validated Go module floor
    carbonstack-os     1bbbe52 docs: clarify CarbonStackOS target direction

Current public release remains:

    v0.6.0 State/UX Boundary Validation Pre-Release

This is **not** a public product release.

### 0.1 Carried-forward v0.6.28 baseline summary

v0.6.28 is complete as a **docs/model checkpoint** for PQ/hybrid placement.

It did five bounded things:

1. Folded state-model review and PQ/hybrid placement into v0.6.28.
2. Added:
   - `docs/224-v0.6.28-pq-hybrid-placement-model-v0.md`
3. Defined the PQ/hybrid placement posture:
   - no mainline PQ implementation in v0.6.x;
   - default future posture is new-conversation-only or new-protocol-epoch;
   - no silent upgrade, downgrade, or mixed mode;
   - outer-envelope / transport experiments are allowed later only as nonclaiming experiments;
   - OpenMLS-native / ciphersuite-layer support remains the honest long-term direction;
   - credential/signature PQ is deferred behind stronger identity/trust/enrollment modeling;
   - first PQ experiments belong in isolated branch/workdir contexts, likely late v0.7.x;
   - mature-ish implementation is deferred to early v0.8.x after experiments and model maturity.
4. Preserved validation and mutation boundaries:
   - `COMMAND_REFERENCE.v0.md` remained current with `entries=86`;
   - no Comms/Cypher/OS runtime code changed;
   - no registry row was added;
   - `full` and `release-snapshot` remained unchanged.
5. Updated downstream ordering toward bounded substrate mechanics, full naming, adversarial harness contract, and v0.7.0 release-boundary planning.

Current public release remains:

    v0.6.0 State/UX Boundary Validation Pre-Release

This is **not** a public product release.

## 1. Current Project Goal

**Active goal:** Continue the v0.6.x OpenMLS message-flow integration and boundary-hardening epoch from the v0.6.29 bounded state substrate / state-audit-dev model-alignment checkpoint toward `full` / `full-validate-release` / `full-runtime-dev` naming cleanup, adversarial harness contract/evidence matrix design, and v0.7.0 release-boundary planning.

After v0.6.29, the project has:

- the v0.6.0 public pre-release still live on Gitea;
- attached v0.6.0 release assets already revalidated after the v0.6.13 tag rewrite;
- public `main`/tag identity hygiene verified after the history rewrite;
- generated `registry/COMMAND_REFERENCE.v0.md`;
- deterministic generated-reference freshness enforced by runner tests;
- registry entry count remains at 86;
- legacy stub commands gated behind explicit opt-in;
- opinionated `message-send-dev` / `message-inbox-dev` wrappers as the recommended dev/pre-alpha normal-message path;
- direct `openmls-*` commands retained as lower-level implementation/proof surfaces;
- Relay onboarding boundaries separated from ordinary message inbox behavior;
- `same-state-integrated-dev`, proving Level 4 same conversation positive-path continuity:
    - Relay onboarding -> joined conversation -> normal message send/open/ack;
- deterministic normal-message failure/classification profiles:
    - wrong conversation;
    - unsupported content type;
    - wrong recipient/device/sidecar;
    - malformed payload/ciphertext;
    - normal application-message replay/duplicate classification;
- Welcome join partial-state safety and `same-state-welcome-join-failure-dev`;
- v0.6.22 state-authority classification:
    - conversation-level provider storage and final conversation directory are runtime authority;
    - summary files are metadata/evidence for existing message-open;
    - device-level provider storage is identity/bootstrap substrate for existing message-open purposes;
- v0.6.25 normal-message envelope metadata classification:
    - payload hash/size and protocol support gates are enforced/rejected before successful open/ack;
    - recipient/delivery-state metadata acts as routing/bookkeeping selection;
    - sender metadata / `from_device` is unverified relay envelope metadata;
- v0.6.26 sender metadata output warning cleanup:
    - `from_device_unverified`;
    - `sender_identity_verified: false`;
    - explicit warning that `from_device` is relay envelope metadata, not verified identity;
- v0.6.26 load-check output cleanup:
    - `provider_reloadable`;
    - `summary_metadata_present`;
    - `provider_storage_present`;
    - metadata-missing output that distinguishes load-check failure from message-open authority;
- v0.6.26 EVERGREEN roadmap refresh:
    - remaining v0.6.x modeling and release-boundary decisions;
    - v0.7.x deployability and state-security integration;
    - v0.8.x maturation/code-health/release-candidate hardening;
    - v0.9.x full adversarial validation across real operational surfaces;
- v0.6.27 vault/security + no-silent state model:
    - authority and sensitivity categories;
    - no-silent mutation laws;
    - backup/recovery posture;
    - migration posture;
    - bounded substrate eligibility;
    - PQ/hybrid dependencies;
    - adversarial harness dependencies;
- v0.6.28 PQ/hybrid placement model:
    - no mainline PQ implementation in v0.6.x;
    - default future posture is new-conversation-only / new-protocol-epoch;
    - no silent upgrade/downgrade/mixed mode;
    - outer-envelope work is nonclaiming experiment only;
    - OpenMLS-native/ciphersuite layer is the honest long-term direction;
    - credential/signature PQ is deferred behind identity/trust modeling;
    - late-v0.7.x experiments and early-v0.8.x mature-ish implementation posture are defined;
- v0.6.29 bounded state substrate mechanics:
    - `state-audit-dev` is the current proto-substrate / inventory precursor;
    - first substrate is non-mutating inventory/checking, not a vault;
    - `state-*` naming is preferred for command/output;
    - `vault-substrate-dev` is docs/internal wording only;
    - Cypher is inventory-only at first;
    - generated sidecar roots require hard dev/test labels;
    - PQ tags are reserved but not implemented;
    - output/model alignment is implemented in `state-audit-dev`.

CarbonStack's broader intent remains:

- **CarbonStack** is the secure-communications appliance-stack project and shared doctrine/release authority.
- **CarbonStackComms** is the text-first messaging client and current OpenMLS sidecar / Relay Space integration surface.
- **CarbonStackCypher** is the self-hostable relay/storage server/API stack.
- **CarbonStackOS** is future constrained Android-derived appliance OS work and remains out of current runnable validation packages.
- The `carbonstack` Gitea repo remains the source of truth for release state, release assets, docs, registry, command-boundary classification, and public claims.
- GitHub mirrors, if present, are secondary push mirrors for discoverability/redundancy only.

**North star:**

> Its goal is not to be a better smartphone. Its goal is to stop being a smartphone.

**Core doctrine:**

> Every feature is guilty until it proves it does not add unacceptable parser, network, sensor, identity, or filesystem authority.

**v0.6.x engineering doctrine after v0.6.29:**

> The normal OpenMLS message-flow spine and deterministic normal-message failure-hardening lane are mostly complete for dev/pre-alpha boundaries. The state/security/no-silent model exists, PQ/hybrid placement exists, and the first non-mutating state-boundary inventory surface now speaks the model language. Next work should cleanly separate release-package validation from live-dev runtime aggregation before adversarial harness contract/evidence-matrix work.

## 2. Current Public Release and Mainline State

Current public release remains:

    https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.6.0

Release title:

    CarbonStack v0.6.0 State/UX Boundary Validation Pre-Release

Release tag:

    v0.6.0

Current post-rewrite tag target:

    70d4318 v0.6.0

Pre-rewrite v0.6.0 tag target for historical context only:

    7e676cd6ad / 7e676cd scripts: add v0.6.0 package rehearsal

Current mainline checkpoint:

    v0.6.29 Bounded State Substrate / State-Audit Model Alignment Checkpoint

Current CarbonStack head:

    44e781c docs: record state audit boundary alignment

Current CarbonStackComms head:

    5a61646 fix: align state audit output with boundary model

v0.6.29 is not a new public release and does not change the release title. It is a post-release model/implementation checkpoint.

Important release-page warning retained:

    Gitea automatically provides Source Code ZIP/TAR.GZ archives for the carbonstack repo at the tag.
    These are not the intended runnable multi-repo validation package.
    Use the attached v0.6.0 package, manifest, checksums, validation freeze, release notes, testing runbook, and LICENSE.

v0.6.29 did not rebuild v0.6.0 release assets and did not change package-root release validation.

New v0.6.29 docs/checkpoints:

    docs/225-v0.6.29-bounded-state-substrate-mechanics-v0.md
    docs/226-v0.6.29-state-audit-dev-model-alignment-v0.md

New v0.6.29 implementation-relevant surface:

    carbonstack-comms/internal/state/state_audit.go
    carbonstack-comms/internal/app/state_audit_dev.go
    carbonstack-comms/internal/app/state_audit_dev_test.go

It remains excluded from:

    full
    release-snapshot
    release package validation
    package-root validation
    adversarial relay harness claims
    hostile-server safety claims
    identity verification claims
    vault/key-storage safety claims
    production backup/restore claims
    PQ/hybrid claims
    production secure-use claims

## 3. Current Repo Heads After v0.6.29

Final pushed heads:

    carbonstack        44e781c docs: record state audit boundary alignment
    carbonstack-comms  5a61646 fix: align state audit output with boundary model
    carbonstack-cypher d18a564 chore: restore validated Go module floor
    carbonstack-os     1bbbe52 docs: clarify CarbonStackOS target direction

Recent pushed history:

    carbonstack:
        44e781c docs: record state audit boundary alignment
        2f0be42 docs: define bounded state substrate mechanics
        4a66681 docs: define PQ hybrid placement model
        52295f4 docs: append readme for creation of 223-v0.6.27-vault-security-nosilent-state-model-v0.md
        e0027a2 docs: created 223-v0.6.27-vault-security-nosilent-state-model-v0.md
        468432b docs: record load-check provider summary output

    carbonstack-comms:
        5a61646 fix: align state audit output with boundary model
        2d13c45 fix: clarify load-check provider metadata output
        61b2e51 fix: label message sender metadata as unverified
        b42aa09 fix: make Welcome join state writes atomic
        53779d5 test: cover message wrapper residual edge cases
        462f796 chore: group Comms CLI help by command boundary

    carbonstack-cypher:
        d18a564 chore: restore validated Go module floor
        52e3673 docs: make Cypher surface evergreen
        3589dbe docs: define local Cypher validation state contract
        15760c6 docs: record provider OpenMLS join boundary
        09afa0b feat: add Relay Space scoped envelope routes
        384affd feat: add Relay Space HTTP API routes

    carbonstack-os:
        1bbbe52 docs: clarify CarbonStackOS target direction
        f984cdf Upload files to "/"
        c764790 chore: fix readme formatting
        953b006 chore: fix readme formatting
        b537475 Add CarbonStackOS north star and initial appliance model
        dab2792 Initial CarbonStack repository structure

Working tree note:

    carbonstack was clean/aligned in final snapshot.
    carbonstack-comms was clean/aligned except ignored local artifacts:
        !! internal/protocol/mls/openmls-sidecar/.carbonstack-openmls-sidecar-state/
        !! internal/protocol/mls/openmls-sidecar/target/
    carbonstack-cypher retained the known ignored local DB:
        !! cypher.db
    carbonstack-os was clean/aligned.

## 4. Validation / Evidence State After v0.6.29

v0.6.29A preflight validation passed:

    generated reference deterministic check;
    runner go test ./...;
    registry missing-nonclaims scan;
    runner doctor;
    sidecar cargo test --quiet warning-free;
    Comms go test ./...;
    Cypher go test ./...;
    same-state-integrated-dev;
    same-state-message-malformed-payload-dev;
    same-state-message-replay-classification-dev.

v0.6.29B docs/model validation passed:

    generated reference deterministic check;
    runner go test ./...;
    registry missing-nonclaims scan;
    runner doctor;
    carbonstack diff --check;
    generated reference deterministic check unchanged after docs patch;
    runner go test ./... after docs patch;
    registry/reference/runner guard;
    cached diff --check;
    post-commit generated reference check;
    post-commit runner tests.

v0.6.29C implementation/docs validation passed:

    focused Comms state/app tests before patch;
    focused Comms state/app tests after patch;
    state-audit-dev text output model-field assertions;
    state-audit-dev JSON output parse/model-field assertions;
    Comms go test ./...;
    CarbonStack git diff --check;
    generated reference deterministic check;
    runner go test ./...;
    registry missing-nonclaims scan;
    registry/reference/runner guard;
    Comms cached diff --check;
    CarbonStack cached diff --check;
    post-commit Comms go test ./...;
    post-commit generated reference deterministic check;
    post-commit runner go test ./...;
    Comms push rc=0;
    CarbonStack push rc=0.

Generated command reference remained:

    entries=86

No registry entry was added.

No `full` or `release-snapshot` behavior changed.

No release package validation surface changed.

This is acceptable for a model-aligned non-mutating state-audit implementation checkpoint.

Do not overclaim that v0.6.29 implements a production vault, secure storage, backup/restore, PQ, or hostile-server safety.

## 5. v0.6.29 Model / Implementation Content Summary

### 5.1 Bounded state substrate model

New doc:

    docs/225-v0.6.29-bounded-state-substrate-mechanics-v0.md

Core model:

    The first substrate is not a vault.
    The first substrate is not secure storage.
    The first substrate is non-mutating inventory/checking.
    Existing state-audit-dev is the current proto-substrate surface.
    Do not create a duplicate command unless later inspection proves it is needed.
    Cypher is inventory-only at first.
    Sidecar dev state can be inspected, but must be labeled dev/pre-alpha and secret-bearing where appropriate.
    PQ/migration/downgrade tags are reserved as future requirements, not implemented.
    full naming comes after substrate mechanics.
    adversarial harness contract comes after full naming.

Included v0 surfaces:

    Comms local state JSON;
    OpenMLS sidecar device root;
    signer.json;
    conversation directories;
    provider-storage.json;
    conversation-summary.json;
    join-summary.json;
    trust.json;
    trust-events.jsonl;
    identity-candidates.json;
    sidecar .carbonstack-openmls-sidecar-state;
    sidecar target/;
    runner temp/generated roots when explicitly scoped.

Excluded or deferred v0 surfaces:

    Cypher DB backup/restore/reconciliation;
    release assets/checksums;
    backups;
    PQ state implementation;
    Android/CarbonStackOS;
    production deployment state.

### 5.2 state-audit-dev output/model alignment

New doc:

    docs/226-v0.6.29-state-audit-dev-model-alignment-v0.md

Changed files:

    carbonstack-comms/internal/state/state_audit.go
    carbonstack-comms/internal/app/state_audit_dev.go
    carbonstack-comms/internal/app/state_audit_dev_test.go

Existing command retained:

    go run ./cmd/comms state-audit-dev

Machine-readable output retained:

    --format json

No new command added.

No alias added.

No `--json` alias added.

Report-level output now includes:

    state_boundary_model_version: v0.6.29-state-boundary-v0
    state_boundary_role: proto_substrate_inventory_check
    proto_substrate: true
    pq_tags_reserved_not_implemented: true

Domain-level output now includes:

    authority_class
    sensitivity_class
    no_silent_rule
    boundary_warning
    cypher_inventory_only

Example domain classifications added by the model-alignment patch:

    comms_state:
        authority_class: runtime_authority
        sensitivity_class: safety_sensitive_possibly_privacy_sensitive
        no_silent_rule: no_silent_replacement_import_or_restore
        boundary_warning: state_boundary_inventory_only_not_backup_restore

    trust_store:
        authority_class: safety_sensitive_future_runtime_authority
        sensitivity_class: safety_sensitive_privacy_sensitive
        no_silent_rule: no_silent_trust_promotion
        boundary_warning: trust_state_inventory_only_not_verification

    trust_history:
        authority_class: metadata_evidence
        sensitivity_class: safety_sensitive_metadata_evidence
        no_silent_rule: no_silent_deletion_or_rewrite
        boundary_warning: trust_history_evidence_inventory_only_not_secure_audit_log

    candidate_store:
        authority_class: safety_sensitive_future_runtime_authority
        sensitivity_class: safety_sensitive_privacy_sensitive
        no_silent_rule: no_silent_verified_import
        boundary_warning: candidate_state_inventory_only_not_verified_identity

    sidecar_generated_state:
        authority_class: dev_runtime_authority_container
        sensitivity_class: secret_bearing_safety_sensitive_dev_scope
        no_silent_rule: no_silent_runtime_authority_regeneration
        boundary_warning: dev_generated_state_not_production_vault

    sidecar_build_output:
        authority_class: generated_disposable
        sensitivity_class: generated_build_output
        no_silent_rule: generated_cleanup_only_when_explicitly_scoped
        boundary_warning: generated_build_output_not_release_material

    local_cypher_db:
        authority_class: server_side_coordination_authority_inventory_only
        sensitivity_class: privacy_sensitive_safety_sensitive_maybe_secret_bearing
        no_silent_rule: inventory_only_no_silent_routing_membership_ack_mutation
        boundary_warning: cypher_inventory_only_not_comms_vault_not_restore
        cypher_inventory_only: true

### 5.3 What v0.6.29 does not change

v0.6.29 does not add:

    new command;
    wrapper;
    alias;
    --json flag;
    registry entry;
    generated command reference change;
    full behavior;
    release-snapshot behavior;
    package-root validation behavior;
    encryption;
    secure storage;
    production vault;
    backup;
    restore;
    migration;
    PQ implementation;
    trust mutation;
    identity verification;
    Cypher mutation;
    adversarial harness implementation.

## 6. v0.6.29 Blunders / Continuity Notes

### 6.1 v0.6.29A changed the expected implementation path

The original mental model was:

    recon/model -> docs/deeper recon if needed -> code

with possible first names like:

    state-boundary-check-dev
    state-boundary-inventory-dev

v0.6.29A found that `state-audit-dev` already existed and was already registered as a non-mutating state-domain inventory surface.

Impact:

    positive correction;
    avoided new command sprawl;
    avoided unnecessary registry mutation;
    kept the first implementation surface narrow.

Lesson:

    Always search existing registry and implementation surfaces before creating new command names.

### 6.2 v0.6.29B scripted push auth failure

The v0.6.29B scripted CarbonStack push failed with the recurring Gitea authentication issue.

Manual retry succeeded:

    4a66681..2f0be42 main -> main

Impact:

    no project issue;
    final remote was updated;
    branch aligned with origin/main after manual retry.

Lesson:

    Continue treating first-push auth rc=128/fatal auth as retryable when local commit state and final push snapshot are clear.

### 6.3 Local ignored generated artifacts remained visible

Final v0.6.29C snapshot retained ignored local artifacts:

    carbonstack-comms/internal/protocol/mls/openmls-sidecar/.carbonstack-openmls-sidecar-state/
    carbonstack-comms/internal/protocol/mls/openmls-sidecar/target/
    carbonstack-cypher/cypher.db

Impact:

    acceptable local dev/test state;
    no tracked dirty state;
    these artifacts are explicitly within the model's generated/dev-local/inventory-only lanes.

Lesson:

    Future `full`/release validation and package-root rehearsals must keep strict dirty tree / artifact hygiene separate from live-dev profile state.

### 6.4 `--format json` retained; no `--json` added

The v0.6.29C inspection considered JSON output shape.

Decision:

    preserve existing --format json;
    do not add --json.

Impact:

    avoids flag-surface churn;
    avoids registry/reference updates for alias behavior;
    keeps machine-readable output stable.

Lesson:

    Add aliases only when there is a user/operator need, not just because a shorthand is convenient.

### 6.5 Validation scope caveat

v0.6.29C patched Comms output/model alignment and docs.

It did not run the full same-state profile ladder after the output patch.

This is acceptable because:

    the changed behavior is `state-audit-dev` output/report structure;
    focused Comms tests passed;
    full Comms tests passed;
    generated reference and runner tests passed;
    no message flow, Cypher API, OpenMLS sidecar protocol, registry, full, or release-snapshot behavior changed.

Do not claim new message-flow evidence from v0.6.29C beyond the validations actually run.

## 7. Critical Path / Function Updates

New critical docs:

    carbonstack/docs/225-v0.6.29-bounded-state-substrate-mechanics-v0.md
    carbonstack/docs/226-v0.6.29-state-audit-dev-model-alignment-v0.md

Updated doc index:

    carbonstack/docs/README.md

New/updated critical implementation files:

    carbonstack-comms/internal/state/state_audit.go
    carbonstack-comms/internal/app/state_audit_dev.go
    carbonstack-comms/internal/app/state_audit_dev_test.go

Existing critical command surface:

    go run ./cmd/comms state-audit-dev
    go run ./cmd/comms state-audit-dev --format json

Important behavior after v0.6.29:

    state-audit-dev is the current proto-substrate / non-mutating inventory precursor.
    It is not a vault.
    It is not secure storage.
    It does not backup/restore.
    It does not mutate trust.
    It does not mutate Cypher.
    It reserves PQ tag language as not implemented.
    It does not join full or release-snapshot.

Existing critical docs still active:

    carbonstack/docs/223-v0.6.27-vault-security-nosilent-state-model-v0.md
    carbonstack/docs/224-v0.6.28-pq-hybrid-placement-model-v0.md
    carbonstack/docs/221-v0.6.26-sender-metadata-output-warning-v0.md
    carbonstack/docs/222-v0.6.26-loadcheck-provider-summary-output-v0.md
    carbonstack/docs/220-v0.6.25-normal-message-envelope-metadata-classification-v0.md
    CarbonStack_Long_Term_Roadmap_v0.6.26_EVERGREEN.pdf

Critical files unchanged in v0.6.29:

    carbonstack/registry/commands.v0.yaml
    carbonstack/registry/COMMAND_REFERENCE.v0.md
    carbonstack/tools/carbonstack-validate/*.go
    carbonstack-comms/internal/app/message_wrappers_dev.go
    carbonstack-comms/internal/app/openmls_runtime.go
    carbonstack-comms/internal/app/openmls_bootstrap.go
    carbonstack-comms/internal/protocol/mls/openmls-sidecar/src/state.rs
    carbonstack-cypher/internal/httpapi/api.go
    carbonstack-cypher/internal/db/relay_spaces.go

`COMMAND_REFERENCE.v0.md` remains:

    entries=86

No registry change occurred in v0.6.29.

## 8. Next Work / Future To-Do

Immediate next recommended rung:

    v0.6.30A = full / full-validate-release / full-runtime-dev naming plan

Purpose:

    separate package-root release validation from live-dev runtime aggregation;
    avoid silently expanding the historical full profile;
    define whether full remains compatibility alias;
    decide when/if full-validate-release and full-runtime-dev become real profiles;
    decide which future surfaces can be aggregated in runtime-dev without contaminating release-package validation.

Likely contract to model:

    full = compatibility alias for release-package validation
    full-validate-release = explicit package-root/release validation
    full-runtime-dev = live-umbrella runtime-dev aggregation

Likely near-term sequence:

    v0.6.30A = full/full-validate-release/full-runtime-dev naming recon/model
    v0.6.30B = docs checkpoint for validation naming plan
    later = implementation only if the plan is low-risk and consensus is clear
    later = adversarial harness contract/evidence matrix
    later = v0.7.0 release-boundary/package-root rehearsal plan

Do not do next:

    create state-boundary-check-dev alias without need;
    add --json alias without need;
    add state-audit-dev to full;
    add live-dev profiles to release-snapshot;
    implement adversarial harness before the validation naming contract is clear;
    implement backup/restore;
    implement vault/encryption;
    implement PQ/hybrid;
    claim secure-use readiness.

Open questions carried forward:

    Should full remain only an alias, or should it be deprecated after full-validate-release exists?
    Should full-validate-release be introduced as a new explicit profile before v0.7.0?
    Should full-runtime-dev exist as one profile or several staged live-dev profiles?
    Should state-audit-dev be included in full-runtime-dev later?
    Should same-state-integrated-dev become part of full-runtime-dev later?
    Should full-runtime-dev require --clean-generated by default?
    Should release-snapshot stay untouched until v0.7.0 package-root rehearsal?
    How should docs distinguish release-package validation, package-root validation, live-dev runtime validation, and source-tree operator checks?
    How should adversarial harness evidence reference validation profiles without implying broad hostile-server safety?

## 9. Current Nonclaims After v0.6.29

v0.6.29 does **not** claim:

    public release;
    public product readiness;
    production secure messaging;
    production E2EE;
    hostile-server safety;
    adversarial relay harness coverage;
    replay safety;
    Welcome replay handling;
    KeyPackage replay handling;
    server equivocation detection;
    metadata privacy;
    metadata authenticity;
    sender authenticity;
    verified sender identity;
    secure enrollment;
    production vault/key storage;
    secure storage;
    encryption-at-rest;
    production backup/restore;
    PQ or hybrid security;
    quantum-safe messaging;
    OpenMLS-native PQ/hybrid support;
    ciphersuite migration;
    algorithm migration;
    Android readiness;
    CarbonStackOS readiness;
    deployment readiness;
    audit;
    certification;
    release-package validation;
    package-root validation;
    inclusion in full;
    inclusion in release-snapshot;
    new command-surface stability;
    state-audit-dev as mature UX;
    state-audit-dev as recovery tooling;
    state-audit-dev as trust verification;
    state-audit-dev as Cypher repair/reconciliation.

## 10. Prior v0.6.28 LogDoc Preserved Below

The remainder of this file preserves the v0.6.28 baseline for continuity.

---

# CarbonStack LogDoc v0.6.28

**Last updated:** 2026-06-21 local session  
**Owner/operator:** bitcrusher32  
**Assistant/model:** GPT-5.5 Thinking  
**Current phase:** Phase 2F -> **v0.6.28 PQ/hybrid placement model docs checkpoint complete; v0.6.28A state-model review + PQ placement recon accepted the v0.6.27 state/security/no-silent model as controlling input, and v0.6.28B committed `docs/224-v0.6.28-pq-hybrid-placement-model-v0.md` as the model gate for future PQ/hybrid experiments and mature implementation attempts.**  
**Current public release:** `v0.6.0 State/UX Boundary Validation Pre-Release`  
**Current mainline checkpoint:** `v0.6.28 PQ/Hybrid Placement Model Docs Checkpoint`  
**Release URL:** https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.6.0  
**Update source:** `CarbonStackLogDocV0.6.27`, `CarbonStack_BreakpointV0.6.27`, v0.6.28A state-model review + PQ/hybrid placement recon, v0.6.28B PQ/hybrid placement model docs checkpoint log, and the active `CarbonStack_Long_Term_Roadmap_v0.6.26_EVERGREEN` planning authority.  
**Update purpose:** Preserve the v0.6.28 PQ/hybrid placement model checkpoint after the v0.6.27 state/no-silent model was accepted as the controlling state input; record the new PQ/hybrid placement posture, model decisions, validation evidence, push/retry behavior, repo heads, critical-path updates, downstream future-work ordering, blunders/continuity notes, and nonclaims; preserve the full prior Markdown ledger; and keep the JSON breakpoint as a lean current-state handoff.  

**Version schema:** `v[scope].[timeline]`. This file is `v0.6.28`, the working branch continuity ledger after the PQ/hybrid placement model checkpoint in the v0.6.x OpenMLS message-flow integration and boundary-hardening epoch.

**LogDoc V2 usage note:** The Markdown LogDoc is intentionally a dev continuity ledger. It preserves historical v0.6.x timeline data, blunders, validation ladders, command-boundary decisions, and future path constraints. The JSON breakpoint remains a lean current-state handoff and should not try to duplicate the full Markdown timeline.

**Privacy / history-rewrite note:** Exact raw shell logs may contain local usernames, hostnames, email addresses, and old pre-rewrite author identities. This LogDoc avoids repeating local host-derived identity strings except where operationally necessary. Public commit identity remains normalized on public `main`/tag histories. Pre-v0.6.13 commit hashes from older LogDocs and historical docs are now historical **pre-rewrite** references and must not be treated as current Gitea `main`/tag SHAs.

## 0. Executive Summary

v0.6.28 is complete as a **docs/model checkpoint** for PQ/hybrid placement.

It did five bounded things:

1. **Folded state-model review and PQ/hybrid placement into v0.6.28**
   - v0.6.27 was already cut as the vault/security + no-silent state model checkpoint.
   - The previously planned `v0.6.27C` state-model review and `v0.6.28A` PQ/hybrid placement recon were intentionally folded into the v0.6.28 line.
   - v0.6.28A accepted `docs/223-v0.6.27-vault-security-nosilent-state-model-v0.md` as the controlling state/security/no-silent model for the PQ/hybrid placement rung.
   - v0.6.28A confirmed that PQ/hybrid must remain model-only in v0.6.x.
   - v0.6.28A confirmed the core dependency chain:
     - state tags;
     - migration epoch;
     - backup/restore semantics;
     - compatibility refusal;
     - no-silent downgrade/upgrade.
   - The recon warning remained central:
     - an outer wrapper may improve transport metadata/story, but does **not** make the OpenMLS group itself PQ-secure.

2. **Added the v0.6.28 PQ/hybrid placement model**
   - Added:
     - `docs/224-v0.6.28-pq-hybrid-placement-model-v0.md`
   - Updated:
     - `docs/README.md`
   - CarbonStack commit:
     - `4a66681 docs: define PQ hybrid placement model`
   - Commit summary:
     - `2 files changed, 783 insertions(+)`
     - `create mode 100644 docs/224-v0.6.28-pq-hybrid-placement-model-v0.md`
   - Initial scripted push hit the recurring Gitea authentication failure:
     - `RESULT carbonstack push rc=128`
   - Manual push succeeded:
     - `52295f4..4a66681 main -> main`
   - Final CarbonStack status:
     - branch up to date with `origin/main`;
     - latest head `4a66681`.

3. **Defined the PQ/hybrid placement posture**
   - PQ/hybrid is **not** added to mainline in v0.6.x.
   - Default future posture is **new-conversation-only** or **new-protocol-epoch**.
   - Existing conversations must not silently become PQ/hybrid.
   - PQ-tagged or upgraded state must not silently downgrade.
   - Mixed mode is refused by default unless a future model explicitly defines it.
   - Outer-envelope / transport experiments are allowed later only as nonclaiming experiments.
   - OpenMLS-native / ciphersuite-layer support remains the honest long-term direction for protocol-level PQ/hybrid semantics.
   - Credential/signature PQ is deferred behind stronger identity/trust/enrollment modeling.
   - First PQ experiments belong in isolated branch/workdir contexts, likely late v0.7.x.
   - A more mature implementation attempt is deferred to early v0.8.x after experiments and model maturity.

4. **Preserved validation and mutation boundaries**
   - Baseline and post-patch checks passed:
     - generated reference deterministic check;
     - runner `go test ./...`;
     - registry missing-nonclaims scan;
     - runner doctor;
     - sidecar `cargo test --quiet` warning-free;
     - `carbonstack diff --check`;
     - cached diff check;
     - post-commit generated reference check;
     - post-commit runner tests.
   - `COMMAND_REFERENCE.v0.md` remained current with:
     - `entries=86`
   - Guard confirmed no unexpected mutation to:
     - `registry/commands.v0.yaml`;
     - `registry/COMMAND_REFERENCE.v0.md`;
     - `tools/carbonstack-validate`.
   - No Comms/Cypher/OS code changed.
   - No registry row was added.
   - `full` and `release-snapshot` remain unchanged.
   - No PQ implementation was added.
   - No OpenMLS ciphersuite changed.
   - No crypto code changed.
   - No vault/substrate code changed.
   - No adversarial harness code changed.
   - No quantum-safe claim was made.

5. **Updated downstream ordering**
   - v0.6.28 completes the PQ/hybrid placement model that was required before PQ experiments or substrate mechanics can honestly proceed.
   - The next likely rung becomes bounded substrate/stub mechanics preflight.
   - Future ordering:
     - bounded substrate/stub mechanics preflight;
     - adversarial harness contract/evidence matrix design;
     - `full` / `full-validate-release` / `full-runtime-dev` naming plan;
     - v0.7.0 release-boundary/package-root rehearsal plan;
     - late-v0.7.x isolated PQ/identity/outer-envelope experiments where models permit;
     - early-v0.8.x mature-ish PQ/deployability implementation attempt only after experimental findings and model maturity.

Current commits:

    carbonstack        4a66681 docs: define PQ hybrid placement model
    carbonstack-comms  2d13c45 fix: clarify load-check provider metadata output
    carbonstack-cypher d18a564 chore: restore validated Go module floor
    carbonstack-os     1bbbe52 docs: clarify CarbonStackOS target direction

Current public release remains:

    v0.6.0 State/UX Boundary Validation Pre-Release

This is **not** a public product release.

### 0.1 Carried-forward v0.6.27 baseline summary

v0.6.27 is complete as a **docs/model checkpoint** for the vault/security + no-silent state model.

It did five bounded things:

1. Promoted the v0.6.27A broad recon into `docs/223-v0.6.27-vault-security-nosilent-state-model-v0.md`.
2. Defined state vocabulary:
   - runtime authority;
   - server-side coordination authority;
   - metadata/evidence;
   - secret-bearing state;
   - privacy-sensitive state;
   - safety-sensitive but not secret;
   - generated/disposable state;
   - public/documentation state.
3. Defined critical decisions:
   - `trust.json` and `identity-candidates.json` are safety-sensitive now and future runtime authority once wired into send/receive/enrollment policy.
   - Cypher DB is server-side coordination authority and delivery/routing authority, not plaintext authority or verified identity authority.
   - `conversation-summary.json` is metadata/evidence, not runtime authority for existing message-open.
   - Welcome/KeyPackage replay should be documented/modelled now but implemented later after onboarding lifecycle and Addressable Relay Space mechanics mature.
   - PQ/hybrid needed a full placement model later, now satisfied by v0.6.28.
4. Defined no-silent laws:
   - no silent runtime-authority regeneration;
   - no silent runtime-authority replacement;
   - no silent verified trust promotion;
   - no silent identity continuity after loss/reinstall/restore;
   - no silent backup restore over newer state;
   - no silent migration;
   - no silent algorithm upgrade or downgrade;
   - no silent deletion of trust/security-relevant state;
   - no ack/drain after failed open or failed join;
   - no silent Relay Space membership or routing mutation.
5. Preserved the hard claim boundary:
   - no vault;
   - no encryption;
   - no production backup/restore;
   - no PQ/hybrid;
   - no adversarial harness;
   - no identity verification;
   - no production-security claim.

Current public release remains:

    v0.6.0 State/UX Boundary Validation Pre-Release

This is **not** a public product release.

## 1. Current Project Goal

**Active goal:** Continue the v0.6.x OpenMLS message-flow integration and boundary-hardening epoch from the v0.6.28 PQ/hybrid placement model checkpoint toward bounded substrate/stub mechanics, adversarial harness contract/evidence matrix design, validation/naming cleanup, and v0.7.0 release-boundary planning.

After v0.6.28, the project has:

- the v0.6.0 public pre-release still live on Gitea;
- attached v0.6.0 release assets already revalidated after the v0.6.13 tag rewrite;
- public `main`/tag identity hygiene verified after the history rewrite;
- generated `registry/COMMAND_REFERENCE.v0.md`;
- deterministic generated-reference freshness enforced by runner tests;
- registry entry count remains at 86;
- legacy stub commands gated behind explicit opt-in;
- opinionated `message-send-dev` / `message-inbox-dev` wrappers as the recommended dev/pre-alpha normal-message path;
- direct `openmls-*` commands retained as lower-level implementation/proof surfaces;
- Relay onboarding boundaries separated from ordinary message inbox behavior;
- `same-state-integrated-dev`, proving Level 4 same conversation positive-path continuity:
    - Relay onboarding -> joined conversation -> normal message send/open/ack;
- deterministic normal-message failure/classification profiles:
    - wrong conversation;
    - unsupported content type;
    - wrong recipient/device/sidecar;
    - malformed payload/ciphertext;
    - normal application-message replay/duplicate classification;
- Welcome join partial-state safety and `same-state-welcome-join-failure-dev`;
- v0.6.22 state-authority classification:
    - conversation-level provider storage and final conversation directory are runtime authority;
    - summary files are metadata/evidence for existing message-open;
    - device-level provider storage is identity/bootstrap substrate for existing message-open purposes;
- v0.6.25 normal-message envelope metadata classification:
    - payload hash/size and protocol support gates are enforced/rejected before successful open/ack;
    - recipient/delivery-state metadata acts as routing/bookkeeping selection;
    - sender metadata / `from_device` is unverified relay envelope metadata;
- v0.6.26 sender metadata output warning cleanup:
    - `from_device_unverified`;
    - `sender_identity_verified: false`;
    - explicit warning that `from_device` is relay envelope metadata, not verified identity;
- v0.6.26 load-check output cleanup:
    - `provider_reloadable`;
    - `summary_metadata_present`;
    - `provider_storage_present`;
    - metadata-missing output that distinguishes load-check failure from message-open authority;
- v0.6.26 EVERGREEN roadmap refresh:
    - remaining v0.6.x modeling and release-boundary decisions;
    - v0.7.x deployability and state-security integration;
    - v0.8.x maturation/code-health/release-candidate hardening;
    - v0.9.x full adversarial validation across real operational surfaces;
- v0.6.27 vault/security + no-silent state model:
    - authority and sensitivity categories;
    - no-silent mutation laws;
    - backup/recovery posture;
    - migration posture;
    - bounded substrate eligibility;
    - PQ/hybrid dependencies;
    - adversarial harness dependencies;
- v0.6.28 PQ/hybrid placement model:
    - no mainline PQ implementation in v0.6.x;
    - default future posture is new-conversation-only / new-protocol-epoch;
    - no silent upgrade/downgrade/mixed mode;
    - outer-envelope work is nonclaiming experiment only;
    - OpenMLS-native/ciphersuite layer is the honest long-term direction;
    - credential/signature PQ is deferred behind identity/trust modeling;
    - late-v0.7.x experiments and early-v0.8.x mature-ish implementation direction are scoped.

CarbonStack's broader intent remains:

- **CarbonStack** is the secure-communications appliance-stack project and shared doctrine/release authority.
- **CarbonStackComms** is the text-first messaging client and current OpenMLS sidecar / Relay Space integration surface.
- **CarbonStackCypher** is the self-hostable relay/storage server/API stack.
- **CarbonStackOS** is future constrained Android-derived appliance OS work and remains out of current runnable validation packages.
- The `carbonstack` Gitea repo remains the source of truth for release state, release assets, docs, registry, command-boundary classification, and public claims.
- GitHub mirrors, if present, are secondary push mirrors for discoverability/redundancy only.

**North star:**

> Its goal is not to be a better smartphone. Its goal is to stop being a smartphone.

**Core doctrine:**

> Every feature is guilty until it proves it does not add unacceptable parser, network, sensor, identity, or filesystem authority.

**v0.6.x engineering doctrine after v0.6.28:**

> The normal OpenMLS message-flow spine and deterministic normal-message failure-hardening lane are mostly complete for dev/pre-alpha boundaries. The state/security/no-silent model and the PQ/hybrid placement model now exist as controlling documents for future bounded substrate and PQ work. Do not implement vault/substrate, PQ/hybrid, or adversarial harness code before the corresponding model/contract and validation posture exist.

## 2. Current Public Release and Mainline State

Current public release remains:

    https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.6.0

Release title:

    CarbonStack v0.6.0 State/UX Boundary Validation Pre-Release

Release tag:

    v0.6.0

Current post-rewrite tag target:

    70d4318 v0.6.0

Pre-rewrite v0.6.0 tag target for historical context only:

    7e676cd6ad / 7e676cd scripts: add v0.6.0 package rehearsal

Current mainline checkpoint:

    v0.6.28 PQ/Hybrid Placement Model Docs Checkpoint

Current CarbonStack head:

    4a66681 docs: define PQ hybrid placement model

v0.6.28 is not a new public release and does not change the release title. It is a post-release docs/model checkpoint.

Important release-page warning retained:

    Gitea automatically provides Source Code ZIP/TAR.GZ archives for the carbonstack repo at the tag.
    These are not the intended runnable multi-repo validation package.
    Use the attached v0.6.0 package, manifest, checksums, validation freeze, release notes, testing runbook, and LICENSE.

v0.6.28 did not rebuild v0.6.0 release assets and did not change package-root release validation.

The new v0.6.28 docs/model checkpoint is:

    docs/224-v0.6.28-pq-hybrid-placement-model-v0.md

It remains excluded from:

    full
    release-snapshot
    release package validation
    package-root validation
    adversarial relay harness claims
    replay-safety claims
    hostile-server safety claims
    identity verification claims
    vault/PQ claims
    production backup/restore claims
    quantum-safe messaging claims
    OpenMLS-native PQ/hybrid claims

## 3. Current Repo Heads After v0.6.28

Final pushed heads:

    carbonstack        4a66681 docs: define PQ hybrid placement model
    carbonstack-comms  2d13c45 fix: clarify load-check provider metadata output
    carbonstack-cypher d18a564 chore: restore validated Go module floor
    carbonstack-os     1bbbe52 docs: clarify CarbonStackOS target direction

Recent pushed history:

    carbonstack:
        4a66681 docs: define PQ hybrid placement model
        52295f4 docs: append readme for creation of 223-v0.6.27-vault-security-nosilent-state-model-v0.md
        e0027a2 docs: created 223-v0.6.27-vault-security-nosilent-state-model-v0.md
        468432b docs: record load-check provider summary output
        499ab6a docs: record sender metadata output warning
        b7f89ea docs: classify normal message envelope metadata
        1c421e2 test: add same-state replay classification profile
        943e0e4 test: add same-state malformed payload profile

    carbonstack-comms:
        2d13c45 fix: clarify load-check provider metadata output
        61b2e51 fix: label message sender metadata as unverified
        b42aa09 fix: make Welcome join state writes atomic
        53779d5 test: cover message wrapper residual edge cases
        462f796 chore: group Comms CLI help by command boundary
        020c07f test: align wrapper smoke with message wrappers
        9391029 feat: add opinionated OpenMLS message wrappers
        e413ca0 test: cover OpenMLS message flow command edge cases

    carbonstack-cypher:
        d18a564 chore: restore validated Go module floor
        52e3673 docs: make Cypher surface evergreen
        3589dbe docs: define local Cypher validation state contract
        15760c6 docs: record provider OpenMLS join boundary
        09afa0b feat: add Relay Space scoped envelope routes
        384affd feat: add Relay Space HTTP API routes
        cc6e589 feat: add Relay Space DB helpers
        6798d11 feat: add Relay Space schema substrate

    carbonstack-os:
        1bbbe52 docs: clarify CarbonStackOS target direction
        f984cdf Upload files to "/"
        c764790 chore: fix readme formatting

Working tree note:

    Manual post-push snapshot showed CarbonStack at 4a66681 and aligned with origin/main.
    carbonstack-comms retained ignored sidecar build output:
        !! internal/protocol/mls/openmls-sidecar/target/
    carbonstack-cypher retained the known ignored local DB:
        !! cypher.db
    carbonstack-os was clean.

## 4. Validation / Evidence State After v0.6.28

v0.6.28 is a docs/model commit, but the rung included a lightweight validation ladder before and after the docs patch.

Pre-patch / baseline validation passed:

    generated reference deterministic check:
        OK: generated reference is current: registry/COMMAND_REFERENCE.v0.md
        entries=86

    runner go test ./...:
        ok

    registry missing-nonclaims scan:
        VALIDATION PASSED

    runner doctor:
        VALIDATION PASSED

    sidecar cargo test warning-free:
        20 passed; 0 failed

The controlling state model was verified to exist:

    docs/223-v0.6.27-vault-security-nosilent-state-model-v0.md

Post-doc validation passed:

    carbonstack diff --check:
        rc=0

    generated reference deterministic check unchanged:
        entries=86
        rc=0

    runner go test ./...:
        rc=0

    registry missing-nonclaims scan:
        VALIDATION PASSED

    runner doctor:
        VALIDATION PASSED

Guard passed:

    registry/reference/runner guard rc=0

Commit validation passed:

    carbonstack cached diff --check:
        rc=0

    commit:
        4a66681 docs: define PQ hybrid placement model

Post-commit validation passed:

    generated reference deterministic check post-commit:
        entries=86
        rc=0

    runner go test ./... post-commit:
        rc=0

Push:

    scripted push failed with the recurring authentication issue:
        RESULT carbonstack push rc=128

    manual push succeeded:
        52295f4..4a66681 main -> main

This is acceptable for a docs/model checkpoint.

Do not overclaim that v0.6.28 added or validated runtime PQ behavior.

## 5. v0.6.28 Model Content Summary

The committed PQ/hybrid placement model defines:

### 5.1 Relationship to state model

`docs/223-v0.6.27-vault-security-nosilent-state-model-v0.md` remains the controlling model for:

    runtime authority;
    server-side coordination authority;
    metadata/evidence;
    secret-bearing state;
    privacy-sensitive state;
    safety-sensitive but not secret state;
    generated/disposable state;
    public/documentation state;
    no-silent mutation rules;
    backup/recovery posture;
    migration posture;
    bounded substrate eligibility.

v0.6.28 adds algorithm/PQ pressure to these rules without weakening them.

### 5.2 Threat claims and nonclaims

The model separates:

    harvest-now-decrypt-later resistance;
    OpenMLS group secret PQ protection;
    credential/signature PQ protection;
    outer transport protection;
    malicious relay resistance;
    server compromise resistance;
    local disk compromise resistance;
    endpoint compromise resistance;
    migration/downgrade safety.

Current posture for all security-significant PQ/hybrid claims remains:

    model only;
    no security claim.

### 5.3 Placement options

The model classifies five options:

#### Option A: OpenMLS-native / ciphersuite layer

Includes:

    MLS ciphersuite;
    HPKE/KEM layer;
    OpenMLS provider behavior;
    credential layer;
    signature layer.

Current posture:

    honest long-term protocol-level direction;
    not first mainline implementation path;
    model only until upstream support and migration realities are clearer.

#### Option B: outer envelope / transport layer

Includes:

    new content type;
    new protocol version;
    hybrid envelope wrapper;
    outer payload protection;
    algorithm metadata around a ciphertext artifact.

Current posture:

    acceptable future experiment only with hard nonclaims;
    must not be described as OpenMLS-level PQ security.

#### Option C: new-conversation-only / new protocol epoch

Current posture:

    recommended default future model.

Reason:

    strongest no-silent posture;
    avoids mixed old/new ambiguity;
    avoids fake in-place migration;
    makes backup/restore and compatibility refusal easier to reason about.

#### Option D: identity credential / signature layer

Current posture:

    future dependency;
    not near-term mainline implementation;
    possible late-v0.7.x experiment only after identity/trust boundaries are clearer.

#### Option E: isolated branch/workdir experiment

Current posture:

    acceptable in late v0.7.x after the placement model;
    not a v0.6.x mainline implementation;
    not a security claim.

### 5.4 Rejected shortcuts

The model rejects:

    “just add PQ”;
    treating outer wrapper as quantum-safe messaging;
    silent in-place upgrade;
    silent downgrade;
    credential/signature PQ before identity model;
    schema-first PQ.

### 5.5 Default posture

The model states:

    PQ/hybrid is model-only in v0.6.x.
    PQ/hybrid should default to new-conversation-only or new-protocol-epoch in future implementation.
    Existing conversations stay old-mode unless an explicit future migration design exists.
    No state silently upgrades.
    No state silently downgrades.
    No backup silently overwrites newer or different algorithm state.
    Outer-envelope experiments are allowed only as nonclaiming experiments.
    OpenMLS-native / ciphersuite-layer support is the honest long-term direction for protocol-level PQ/hybrid.
    Credential/signature PQ waits for future identity/trust modeling.
    First PQ experiments belong in isolated branch/workdir contexts during late v0.7.x.
    A more mature implementation attempt belongs near early v0.8.x after experiments and model review.

### 5.6 Required future state tags

The model identifies required future tag families:

    conversation algorithm mode;
    protocol epoch;
    PQ/hybrid mode;
    OpenMLS ciphersuite identifier;
    outer envelope algorithm identifier;
    credential/signature algorithm family;
    migration state;
    migration ID;
    migration timestamp;
    backup compatibility version;
    downgrade refusal marker;
    client capability marker;
    creation mode.

These tags must affect:

    refusal behavior;
    compatibility checks;
    backup/restore;
    operator warnings.

### 5.7 Compatibility/refusal posture

The model defines a future default compatibility/refusal matrix:

    classical conversation + classical-capable client:
        allow under existing nonclaims;

    classical conversation + future PQ-capable client:
        allow as legacy/classical mode with explicit label;

    PQ/hybrid-tagged conversation + old-only client:
        refuse;

    PQ/hybrid-tagged conversation with missing algorithm tag:
        refuse or recovery-required;

    classical backup restored over PQ-tagged local state:
        refuse or explicit recovery-required;

    PQ-tagged backup into old-only client:
        refuse;

    mixed member capability without model:
        refuse;

    outer-wrapper experimental conversation opened as OpenMLS-native PQ:
        refuse or warn;

    credential/signature PQ metadata without trust model:
        treat as unverified/future dependency.

### 5.8 Future-epoch placement

v0.6.x owns:

    placement model;
    state/PQ/recovery/compatibility decisions;
    no mainline PQ implementation;
    no quantum-safe claim.

Late v0.7.x may own:

    isolated Option E experiments;
    possible Option D identity/credential exploration if trust model is ready;
    outer-envelope experiments with hard nonclaims;
    PQ state-tagging experiments;
    compatibility/refusal experiments.

Early v0.8.x may own:

    more mature implementation attempt;
    PQ/deployability maturation;
    migration/compatibility tests;
    operator warnings;
    algorithm/version display;
    package-root runtime validation candidates.

v0.9.x owns:

    adversarial validation across real operational surfaces;
    PQ compatibility abuse cases;
    migration downgrade abuse cases;
    hostile-server and deployment-surface tests;
    case-specific evidence matrices.

## 6. v0.6.28 Blunders / Continuity Notes

### 6.1 Shell paste-tail corruption at top of log

The pasted terminal log showed a corrupted fragment near the top of the script body:

    echo "REPORT_PATH=$report"t"n attempt is deferred to early v0.8.x after experiments/model maturity.carbonstack-validate,

This appears to be paste-tail/textbox corruption rather than script logic corruption.

Evidence:

    the script proceeded into the expected v0.6.28B checkpoint;
    baseline validation ran;
    docs/224 was created;
    README was updated;
    diff checks passed;
    cached diff check passed;
    commit succeeded;
    post-commit validation passed;
    final checkpoint meaning printed.

Impact:

    no repository damage;
    no failed command attributable to the corrupted visual fragment.

Carry-forward rule:

    continue trusting explicit RESULT markers, commit output, and final repo snapshots over visually corrupted pasted shell lead-in/tail fragments.

### 6.2 Recurring Gitea push authentication failure

The scripted push failed again:

    RESULT carbonstack push rc=128

Manual push succeeded:

    52295f4..4a66681 main -> main

Impact:

    no project issue;
    remote `main` updated;
    final manual snapshot showed CarbonStack aligned with origin/main.

Carry-forward rule:

    scripted push failures are not evidence of implementation failure when validation, commit state, and manual push snapshot are clean.

### 6.3 Docs/model checkpoint with light validation only

v0.6.28 was docs/model only.

It ran useful lightweight validation:

    generated reference check;
    runner tests;
    registry nonclaims scan;
    doctor;
    sidecar cargo test;
    diff checks;
    post-commit generated reference check;
    post-commit runner tests.

It did not run the full same-state profile ladder.

This is acceptable because:

    no runtime code changed;
    no registry changed;
    no runner changed;
    no Comms/Cypher/OS source changed;
    no profile semantics changed.

Future code-bearing, registry-bearing, or runner-bearing rungs must run the heavier ladder.

### 6.4 v0.6.28 intentionally froze implementation scope

The model explicitly rejects implementation shortcuts:

    no mainline PQ in v0.6.x;
    no schema-first PQ;
    no silent in-place migration;
    no silent downgrade;
    no outer-wrapper-equals-quantum-safe language;
    no credential/signature PQ before identity/trust modeling.

This should be preserved because future PQ work is attractive but easy to overclaim.

## 7. Critical Path / Function Updates

New critical doc:

    carbonstack/docs/224-v0.6.28-pq-hybrid-placement-model-v0.md

Updated doc index:

    carbonstack/docs/README.md

Existing controlling state model still active:

    carbonstack/docs/223-v0.6.27-vault-security-nosilent-state-model-v0.md

Existing critical docs still active:

    docs/217-v0.6.22-sidecar-state-authority-classification-v0.md
    docs/220-v0.6.25-normal-message-envelope-metadata-classification-v0.md
    docs/221-v0.6.26-sender-metadata-output-warning-v0.md
    docs/222-v0.6.26-loadcheck-provider-summary-output-v0.md
    CarbonStack_Long_Term_Roadmap_v0.6.26_EVERGREEN.pdf

Critical implementation files unchanged in v0.6.28:

    carbonstack-comms/internal/app/message_wrappers_dev.go
    carbonstack-comms/internal/app/openmls_runtime.go
    carbonstack-comms/internal/app/openmls_bootstrap.go
    carbonstack-comms/internal/protocol/mls/openmls-sidecar/src/state.rs
    carbonstack-cypher/internal/httpapi/api.go
    carbonstack-cypher/internal/db/relay_spaces.go
    carbonstack/tools/carbonstack-validate/*.go
    carbonstack/registry/commands.v0.yaml
    carbonstack/registry/COMMAND_REFERENCE.v0.md

`COMMAND_REFERENCE.v0.md` remains:

    entries=86

No registry change occurred in v0.6.28.

## 8. Next Work / Future To-Do

Immediate next recommended rung:

    v0.6.29A = bounded substrate/stub mechanics preflight

Purpose:

    inspect the v0.6.27 state/security/no-silent model and v0.6.28 PQ/hybrid placement model;
    decide what a non-production bounded state substrate/stub may wrap first;
    decide naming so it does not imply production secure vault behavior;
    decide whether it should be inventory-only, refusal-gate-only, or minimal wrapper;
    decide whether Cypher DB remains inventory-only for now;
    identify first validation profile or docs checkpoint for no-silent substrate behavior.

Likely follow-up rungs:

    later = adversarial harness contract/evidence matrix design;
    later = full/full-validate-release/full-runtime-dev naming plan;
    later = v0.7.0 release-boundary/package-root rehearsal plan;
    late v0.7.x = isolated PQ/identity/outer-envelope experiments where models permit;
    early v0.8.x = more mature PQ/deployability implementation attempt after experiments/model maturity.

Do not do next:

    implement production vault;
    implement encryption;
    implement OpenMLS-native PQ/hybrid ciphersuites;
    implement adversarial harness logic;
    mutate registry for docs-only model unless adding explicit command surfaces later;
    change message-open behavior;
    change ack/drain behavior;
    claim identity verification;
    claim hostile-server safety;
    claim vault/key-storage safety;
    claim production backup/restore;
    claim quantum-safe messaging;
    merge PQ experiment work into mainline without model gates.

Open questions carried forward:

    What exact future name should the bounded substrate use?
    Should a future state inventory command exist?
    Should metadata/evidence checks be separated from runtime load checks?
    What is the minimum backup manifest format needed before v0.7.x deployability?
    Which state surfaces should be included in the first bounded substrate/stub?
    Should future substrate checks refuse dirty generated state in repo-local sidecar roots?
    Should cypher.db in repo root remain ignored/test-local only, or should a future deployment model force explicit data dirs?
    How should Relay Space membership state be backed up, restored, or reconciled?
    How should stale backup detection work before cryptographic signatures or production vault exist?
    What validation profile should first prove no-silent substrate behavior?
    What exact algorithm/state tags are required before the first late-v0.7.x PQ experiment?
    Which PQ experiment should happen first: outer-envelope/transport or identity/credential exploration?
    What operator-visible wording prevents PQ experiments from becoming security claims?

## 9. Current Nonclaims After v0.6.28

v0.6.28 does **not** claim:

    public release;
    public product readiness;
    production secure messaging;
    production E2EE;
    hostile-server safety;
    adversarial relay harness coverage;
    replay safety;
    Welcome replay handling;
    KeyPackage replay handling;
    server equivocation detection;
    metadata privacy;
    metadata authenticity;
    sender authenticity;
    verified sender identity;
    secure enrollment;
    production vault/key storage;
    production backup/restore;
    PQ implementation;
    hybrid implementation;
    quantum-safe messaging;
    OpenMLS-native PQ/hybrid support;
    ciphersuite migration;
    algorithm migration;
    Android readiness;
    CarbonStackOS readiness;
    deployment readiness;
    audit;
    certification;
    release-package validation;
    package-root validation;
    inclusion in full;
    inclusion in release-snapshot.

## 10. Prior v0.6.27 LogDoc Preserved Below

The remainder of this file preserves the v0.6.27 baseline for continuity.

---

# CarbonStack LogDoc v0.6.27

**Last updated:** 2026-06-20 local session  
**Owner/operator:** bitcrusher32  
**Assistant/model:** GPT-5.5 Thinking  
**Current phase:** Phase 2F -> **v0.6.27 vault/security + no-silent state model docs checkpoint complete; broad v0.6.27A recon was promoted into `docs/223-v0.6.27-vault-security-nosilent-state-model-v0.md`; v0.6.26 roadmap/output-semantics cleanup remains the active immediate baseline, and the next work is review/consensus on this state model before PQ placement or bounded substrate mechanics.**  
**Current public release:** `v0.6.0 State/UX Boundary Validation Pre-Release`  
**Current mainline checkpoint:** `v0.6.27 Vault/Security No-Silent State Model Docs Checkpoint`  
**Release URL:** https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.6.0  
**Update source:** `CarbonStackLogDocV0.6.26`, `CarbonStack_BreakpointV0.6.26`, v0.6.27A vault/security/no-silent broad recon, v0.6.27B state-model Markdown, v0.6.27 commit/push snapshot, and the active `CarbonStack_Long_Term_Roadmap_v0.6.26_EVERGREEN` planning authority.  
**Update purpose:** Preserve the v0.6.27 docs/model checkpoint after the broad vault/security/no-silent recon was converted into a committed state model; record the model categories, no-silent rules, backup/recovery/migration posture, bounded substrate eligibility, and PQ/adversarial dependencies; record the typo/push blunders; update repo heads, critical paths, future work, and nonclaims; preserve the full v0.6.x timeline ledger; and keep the JSON breakpoint as a lean current-state handoff.  

**Version schema:** `v[scope].[timeline]`. This file is `v0.6.27`, the working branch continuity ledger after the vault/security + no-silent state model checkpoint in the v0.6.x OpenMLS message-flow integration and boundary-hardening epoch.

**LogDoc V2 usage note:** The Markdown LogDoc is intentionally a dev continuity ledger. It preserves historical v0.6.x timeline data, blunders, validation ladders, command-boundary decisions, and future path constraints. The JSON breakpoint remains a lean current-state handoff and should not try to duplicate the full Markdown timeline.

**Privacy / history-rewrite note:** Exact raw shell logs may contain local usernames, hostnames, email addresses, and old pre-rewrite author identities. This LogDoc avoids repeating local host-derived identity strings except where operationally necessary. Public commit identity remains normalized on public `main`/tag histories. Pre-v0.6.13 commit hashes from older LogDocs and historical docs are now historical **pre-rewrite** references and must not be treated as current Gitea `main`/tag SHAs.

## 0. Executive Summary

v0.6.27 is complete as a **docs/model checkpoint** for the vault/security + no-silent state model.

It did five bounded things:

1. **Promoted v0.6.27A recon into a committed model document**
   - v0.6.27A performed a broad state/security/no-silent recon across:
     - CarbonStack;
     - CarbonStackComms;
     - CarbonStackCypher;
     - CarbonStackOS.
   - The recon was intentionally broad and high-breadth.
   - It inventoried:
     - state/path/file-write surfaces;
     - secret/privacy-sensitive surfaces;
     - no-silent operation surfaces;
     - trust/candidate/identity surfaces;
     - OpenMLS sidecar provider state surfaces;
     - Cypher DB/schema/migration/envelope state surfaces;
     - runner/profile artifact and cleanup surfaces;
     - docs/model precedents.
   - The recon preserved the non-goal boundary:
     - no vault implementation;
     - no encryption implementation;
     - no secure storage claim;
     - no PQ implementation;
     - no adversarial harness implementation;
     - no registry mutation;
     - no generated command reference mutation;
     - no `full`/`release-snapshot` mutation;
     - no public release;
     - no production-security claim.

2. **Added the v0.6.27 state/security/no-silent model**
   - Added:
     - `docs/223-v0.6.27-vault-security-nosilent-state-model-v0.md`
   - Updated:
     - `docs/README.md`
   - CarbonStack commits:
     - `e0027a2 docs: created 223-v0.6.27-vault-security-nosilent-state-model-v0.md`
     - `52295f4 docs: append readme for creation of 223-v0.6.27-vault-security-nosilent-state-model-v0.md`
   - Manual push initially hit the recurring Gitea authentication failure, then succeeded:
     - `468432b..52295f4 main -> main`
   - Final CarbonStack status:
     - branch up to date with `origin/main`;
     - working tree clean.

3. **Defined the model that future state/vault work must obey**
   - The model explicitly says the state problem is not simply:
     - `secret vs not secret`
   - The correct model is:
     - authority class;
     - sensitivity class;
     - allowed mutation behavior;
     - no-silent rule;
     - recovery posture.
   - Defined state vocabulary:
     - runtime authority;
     - server-side coordination authority;
     - metadata/evidence;
     - secret-bearing state;
     - privacy-sensitive state;
     - safety-sensitive but not secret;
     - generated/disposable state;
     - public/documentation state.
   - Important decisions:
     - `trust.json` and `identity-candidates.json` are **safety-sensitive now** and **future runtime authority once wired into send/receive/enrollment policy**.
     - Cypher DB is **server-side coordination authority** and **delivery/routing authority**, but **not message plaintext authority** and **not verified identity authority**.
     - `conversation-summary.json` remains **metadata/evidence**, not runtime authority for existing message-open.
     - Welcome/KeyPackage replay should be documented/modelled now, but implementation should wait until onboarding lifecycle and Addressable Relay Space mechanics are mature enough.
     - PQ/hybrid appears in dependency notes only; full PQ/hybrid placement model comes later.

4. **Recorded no-silent laws and future implementation boundaries**
   - The model defines the hard no-silent rule families:
     - no silent runtime-authority regeneration;
     - no silent runtime-authority replacement;
     - no silent verified trust promotion;
     - no silent identity continuity after loss/reinstall/restore;
     - no silent backup restore over newer state;
     - no silent migration;
     - no silent algorithm upgrade or downgrade;
     - no silent deletion of trust/security-relevant state;
     - no ack/drain after failed open or failed join;
     - no silent Relay Space membership or routing mutation.
   - The model also defines:
     - recovery/degraded-state posture;
     - backup model v0;
     - migration model v0;
     - bounded substrate eligibility;
     - PQ/hybrid dependencies;
     - adversarial harness dependencies;
     - Welcome/KeyPackage placement note;
     - open implementation questions.
   - The model explicitly warns that a future bounded substrate can help enforce boundaries, but is **not a production vault**.

5. **Preserved claim boundaries and next-rung direction**
   - v0.6.27 is not a public release.
   - Current public release remains `v0.6.0`.
   - No Comms/Cypher/OS runtime code changed.
   - No registry entry was added.
   - `COMMAND_REFERENCE.v0.md` remains at `entries=86`.
   - `full` and `release-snapshot` remain unchanged.
   - No message-open behavior changed.
   - No ack/drain behavior changed.
   - No provider-state behavior changed.
   - No identity verification was implemented.
   - No vault/PQ/adversarial harness implementation was added.
   - The next safest work is:
     - v0.6.27C review/consensus pass on the state model;
     - then PQ/hybrid placement model recon or draft;
     - then bounded substrate/stub mechanics preflight.
   - Do not implement a bounded substrate before the model is reviewed.
   - Do not implement PQ before PQ placement is modeled.
   - Do not implement adversarial harness logic before the harness contract/evidence matrix is written.

Current commits:

    carbonstack        52295f4 docs: append readme for creation of 223-v0.6.27-vault-security-nosilent-state-model-v0.md
    carbonstack-comms  2d13c45 fix: clarify load-check provider metadata output
    carbonstack-cypher d18a564 chore: restore validated Go module floor
    carbonstack-os     1bbbe52 docs: clarify CarbonStackOS target direction

Current public release remains:

    v0.6.0 State/UX Boundary Validation Pre-Release

This is **not** a public product release.

### 0.1 Carried-forward v0.6.26 baseline summary

v0.6.26 is complete as a roadmap-refresh plus output-semantics cleanup checkpoint.

It did six bounded things:

1. **Generated the v0.6.26 EVERGREEN roadmap refresh**
   - The v0.6.26 roadmap became the active forward-planning authority after the v0.6.26 adversarial Q/A decision matrix.
   - It explicitly keeps CarbonStack in pre-alpha / experimental planning mode.
   - It preserves the key position that v0.7.0 is a cumulative pre-alpha engineering release boundary, not a production or secure-use release.
   - It defines the relative lane roadmap:
     - remaining v0.6.x decisions/modeling;
     - v0.7.x deployability and state-security integration;
     - v0.8.x PQ/deployability maturation, migration/compatibility, code-health optimization, operator docs, and release-candidate hardening;
     - v0.9.x full self-pen-test / adversarial validation across real operational surfaces.
   - Exact future breakpoint numbers remain LogDoc-level decisions, not hardline evergreen-roadmap promises.

2. **Ran v0.6.26B output-semantics scout**
   - The scout confirmed two release-boundary clarity gaps:
     - normal message inbox output still printed `from_device` without an explicit unverified sender metadata warning;
     - `openmls-conversation-load-check-dev` still conflated provider reloadability and summary metadata presence in output.
   - Scout validation passed:
     - generated command reference check stayed current with `entries=86`;
     - runner `go test ./...` passed;
     - registry missing-nonclaims scan passed;
     - runner doctor passed;
     - sidecar `cargo test --quiet` passed warning-free;
     - Comms `go test ./...` passed;
     - post-scout replay and malformed-payload profiles passed.
   - Scout result summary:
     - sender grep: 723 lines, 24 `from_device` hits;
     - load-check grep: 1801 lines, 105 `conversation-summary` hits, 200 `group_reloadable` hits;
     - tests grep: 223 lines;
     - live probe showed `from_device` present but no `sender_identity_verified`, `from_device_unverified`, or `sender_device_id_unverified`.

3. **Landed v0.6.26C sender metadata output warning cleanup**
   - Comms commit:
     - `61b2e51 fix: label message sender metadata as unverified`
   - CarbonStack docs commit:
     - `499ab6a docs: record sender metadata output warning`
   - Added:
     - `docs/221-v0.6.26-sender-metadata-output-warning-v0.md`
   - Updated:
     - `docs/README.md`
   - Normal message inbox outputs now preserve the legacy field:
     - `from_device: <sender_device_id>`
   - They also add the explicit trust-boundary fields:
     - `from_device_unverified: <sender_device_id>`
     - `sender_identity_verified: false`
     - `warning: from_device is relay envelope metadata, not verified identity`
   - This applies to normal application-message inbox/open outputs:
     - `message-inbox-dev`
     - `openmls-inbox-dev`
   - It does not redefine Relay onboarding artifact sender fields.

4. **Landed v0.6.26D load-check provider/summary output cleanup**
   - Comms commit:
     - `2d13c45 fix: clarify load-check provider metadata output`
   - CarbonStack docs commit:
     - `468432b docs: record load-check provider summary output`
   - Added:
     - `docs/222-v0.6.26-loadcheck-provider-summary-output-v0.md`
   - Updated:
     - `docs/README.md`
   - Successful `openmls-conversation-load-check-dev` output now distinguishes:
     - `group_reloadable: true`
     - `provider_reloadable: true`
     - `summary_metadata_present: true`
     - `provider_storage_present: true`
     - `summary_metadata_path: <path>`
     - `provider_storage_path: <path>`
     - `summary_metadata_warning: none`
   - Missing summary metadata now emits an explicit metadata-missing block before returning failure:
     - `status: metadata_missing`
     - `provider_reloadable: not_evaluated_by_load_check`
     - `summary_metadata_present: false`
     - `provider_storage_present: true/false`
     - `summary_metadata_warning: conversation-summary metadata is missing; conversation-load-check-dev is stricter than message-open and cannot confirm reloadability without summary metadata`
   - This keeps the v0.6.22/v0.6.24 model intact:
     - provider storage remains runtime authority for existing message-open;
     - summary metadata remains metadata/evidence;
     - load-check can be stricter than message-open, but now says so clearly.

5. **Recovered from two v0.6.26D implementation failures**
   - The first v0.6.26D script failed mid-run because the patch incorrectly skipped success-path output insertion after seeing `provider_reloadable` in the error path.
   - The first recovery inserted success-path fields but printed literal `\n` sequences, collapsed several output fields onto one line, and overpatched add-member/join tests that should not expect load-check-only fields.
   - The clean reapply reset only the partial v0.6.26D files, preserved v0.6.26C, reapplied a narrower patch, patched only the load-check unit test, and succeeded.
   - Final clean live probe passed:
     - `baseline_output_has_provider_summary_fields: yes`
     - `missing_summary_output_has_metadata_missing_fields: yes`
     - `missing_summary_rc: 1`

6. **Preserved claim boundaries and next-rung direction**
   - v0.6.26 is not a public release.
   - Current public release remains `v0.6.0`.
   - No registry entry was added.
   - `COMMAND_REFERENCE.v0.md` remained at `entries=86`.
   - `full` and `release-snapshot` remain unchanged.
   - No message-open behavior changed.
   - No ack/drain behavior changed.
   - No provider-state behavior changed.
   - No identity verification was implemented.
   - No vault/PQ/adversarial harness work was implemented.
   - The next safest mainline work is the v0.6.x model lane:
     - vault/security model;
     - no-silent regeneration/replacement/import/delete/migration/restore rules;
     - backup/recovery model;
     - bounded vault substrate/stub decision;
     - PQ/hybrid placement model;
     - adversarial harness contract/evidence matrix;
     - v0.7.0 release/package-root rehearsal plan.

Current commits:

    carbonstack        468432b docs: record load-check provider summary output
    carbonstack-comms  2d13c45 fix: clarify load-check provider metadata output
    carbonstack-cypher d18a564 chore: restore validated Go module floor
    carbonstack-os     1bbbe52 docs: clarify CarbonStackOS target direction

Current public release remains:

    v0.6.0 State/UX Boundary Validation Pre-Release

This is **not** a public product release.


## 1. Current Project Goal

**Active goal:** Continue the v0.6.x OpenMLS message-flow integration and boundary-hardening epoch from the v0.6.27 vault/security + no-silent state model checkpoint toward model review, PQ/hybrid placement modeling, bounded substrate/stub mechanics, adversarial harness contract/evidence matrix, and v0.7.0 release-boundary planning.

After v0.6.27, the project has:

- the v0.6.0 public pre-release still live on Gitea;
- attached v0.6.0 release assets already revalidated after the v0.6.13 tag rewrite;
- public `main`/tag identity hygiene verified after the history rewrite;
- generated `registry/COMMAND_REFERENCE.v0.md`;
- deterministic generated-reference freshness enforced by runner tests;
- registry entry count remains at 86;
- legacy stub commands gated behind explicit opt-in;
- opinionated `message-send-dev` / `message-inbox-dev` wrappers as the recommended dev/pre-alpha normal-message path;
- direct `openmls-*` commands retained as lower-level implementation/proof surfaces;
- Relay onboarding boundaries separated from ordinary message inbox behavior;
- `same-state-integrated-dev`, proving Level 4 same conversation positive-path continuity:
    - Relay onboarding -> joined conversation -> normal message send/open/ack;
- deterministic normal-message failure/classification profiles:
    - wrong conversation;
    - unsupported content type;
    - wrong recipient/device/sidecar;
    - malformed payload/ciphertext;
    - normal application-message replay/duplicate classification;
- Welcome join partial-state safety and `same-state-welcome-join-failure-dev`;
- v0.6.22 state-authority classification:
    - conversation-level provider storage and final conversation directory are runtime authority;
    - summary files are metadata/evidence for existing message-open;
    - device-level provider storage is identity/bootstrap substrate for existing message-open purposes;
- v0.6.25 normal-message envelope metadata classification:
    - payload hash/size and protocol support gates are enforced/rejected before successful open/ack;
    - recipient/delivery-state metadata acts as routing/bookkeeping selection;
    - sender metadata / `from_device` is unverified relay envelope metadata;
- v0.6.26 sender metadata output warning cleanup:
    - `from_device_unverified`;
    - `sender_identity_verified: false`;
    - explicit warning that `from_device` is relay envelope metadata, not verified identity;
- v0.6.26 load-check output cleanup:
    - `provider_reloadable`;
    - `summary_metadata_present`;
    - `provider_storage_present`;
    - metadata-missing output that distinguishes load-check failure from message-open authority;
- v0.6.26 EVERGREEN roadmap refresh:
    - remaining v0.6.x modeling and release-boundary decisions;
    - v0.7.x deployability and state-security integration;
    - v0.8.x maturation/code-health/release-candidate hardening;
    - v0.9.x full adversarial validation across real operational surfaces;
- v0.6.27 vault/security + no-silent state model:
    - authority and sensitivity categories;
    - no-silent mutation laws;
    - backup/recovery posture;
    - migration posture;
    - bounded substrate eligibility;
    - PQ/hybrid dependencies;
    - adversarial harness dependencies.

CarbonStack's broader intent remains:

- **CarbonStack** is the secure-communications appliance-stack project and shared doctrine/release authority.
- **CarbonStackComms** is the text-first messaging client and current OpenMLS sidecar / Relay Space integration surface.
- **CarbonStackCypher** is the self-hostable relay/storage server/API stack.
- **CarbonStackOS** is future constrained Android-derived appliance OS work and remains out of current runnable validation packages.
- The `carbonstack` Gitea repo remains the source of truth for release state, release assets, docs, registry, command-boundary classification, and public claims.
- GitHub mirrors, if present, are secondary push mirrors for discoverability/redundancy only.

**North star:**

> Its goal is not to be a better smartphone. Its goal is to stop being a smartphone.

**Core doctrine:**

> Every feature is guilty until it proves it does not add unacceptable parser, network, sensor, identity, or filesystem authority.

**v0.6.x engineering doctrine after v0.6.27:**

> The normal OpenMLS message-flow spine and deterministic normal-message failure-hardening lane are mostly complete for dev/pre-alpha boundaries. The state/security/no-silent model now exists as the controlling document for future bounded substrate and recovery work. Next work should review the model, then define PQ/hybrid placement and bounded substrate mechanics. Do not implement vault, PQ, or adversarial harness code before the corresponding model/contract exists.

## 2. Current Public Release and Mainline State

Current public release remains:

    https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.6.0

Release title:

    CarbonStack v0.6.0 State/UX Boundary Validation Pre-Release

Release tag:

    v0.6.0

Current post-rewrite tag target:

    70d4318 v0.6.0

Pre-rewrite v0.6.0 tag target for historical context only:

    7e676cd6ad / 7e676cd scripts: add v0.6.0 package rehearsal

Current mainline checkpoint:

    v0.6.27 Vault/Security No-Silent State Model Docs Checkpoint

Current CarbonStack head:

    52295f4 docs: append readme for creation of 223-v0.6.27-vault-security-nosilent-state-model-v0.md

v0.6.27 is not a new public release and does not change the release title. It is a post-release docs/model checkpoint.

Important release-page warning retained:

    Gitea automatically provides Source Code ZIP/TAR.GZ archives for the carbonstack repo at the tag.
    These are not the intended runnable multi-repo validation package.
    Use the attached v0.6.0 package, manifest, checksums, validation freeze, release notes, testing runbook, and LICENSE.

v0.6.27 did not rebuild v0.6.0 release assets and did not change package-root release validation.

The new v0.6.27 docs/model checkpoint is:

    docs/223-v0.6.27-vault-security-nosilent-state-model-v0.md

It remains excluded from:

    full
    release-snapshot
    release package validation
    package-root validation
    adversarial relay harness claims
    replay-safety claims
    hostile-server safety claims
    identity verification claims
    vault/PQ claims
    production backup/restore claims

## 3. Current Repo Heads After v0.6.27

Final pushed heads:

    carbonstack        52295f4 docs: append readme for creation of 223-v0.6.27-vault-security-nosilent-state-model-v0.md
    carbonstack-comms  2d13c45 fix: clarify load-check provider metadata output
    carbonstack-cypher d18a564 chore: restore validated Go module floor
    carbonstack-os     1bbbe52 docs: clarify CarbonStackOS target direction

Recent pushed history:

    carbonstack:
        52295f4 docs: append readme for creation of 223-v0.6.27-vault-security-nosilent-state-model-v0.md
        e0027a2 docs: created 223-v0.6.27-vault-security-nosilent-state-model-v0.md
        468432b docs: record load-check provider summary output
        499ab6a docs: record sender metadata output warning
        b7f89ea docs: classify normal message envelope metadata
        1c421e2 test: add same-state replay classification profile
        943e0e4 test: add same-state malformed payload profile
        8dd7db5 docs: classify sidecar state authority

    carbonstack-comms:
        2d13c45 fix: clarify load-check provider metadata output
        61b2e51 fix: label message sender metadata as unverified
        b42aa09 fix: make Welcome join state writes atomic
        53779d5 test: cover message wrapper residual edge cases
        462f796 chore: group Comms CLI help by command boundary
        020c07f test: align wrapper smoke with message wrappers
        9391029 feat: add opinionated OpenMLS message wrappers
        e413ca0 test: cover OpenMLS message flow command edge cases

    carbonstack-cypher:
        d18a564 chore: restore validated Go module floor
        52e3673 docs: make Cypher surface evergreen
        3589dbe docs: define local Cypher validation state contract
        15760c6 docs: record provider OpenMLS join boundary
        09afa0b feat: add Relay Space scoped envelope routes
        384affd feat: add Relay Space HTTP API routes
        cc6e589 feat: add Relay Space DB helpers
        6798d11 feat: add Relay Space schema substrate

    carbonstack-os:
        1bbbe52 docs: clarify CarbonStackOS target direction
        f984cdf Upload files to "/"
        c764790 chore: fix readme formatting
        953b006 chore: fix readme formatting
        b537475 Add CarbonStackOS north star and initial appliance model
        dab2792 Initial CarbonStack repository structure

Working tree note:

    carbonstack, carbonstack-comms, and carbonstack-os were clean in the final snapshot.
    carbonstack-cypher retained the known ignored local DB:
        !! cypher.db

## 4. Validation / Evidence State After v0.6.27

v0.6.27 itself was a docs/model commit and did not include a new automated validation ladder in the pasted commit log.

Evidence available for acceptance:

- v0.6.27A broad recon baseline validation passed before the model was written:
    - generated reference deterministic check passed;
    - `registry/COMMAND_REFERENCE.v0.md` stayed current with `entries=86`;
    - runner `go test ./...` passed;
    - registry missing-nonclaims scan passed;
    - runner doctor passed;
    - sidecar `cargo test --quiet` passed warning-free;
    - Comms `go test ./...` passed;
    - Cypher `go test ./...` passed;
    - focused smoke profiles passed:
        - `same-state-integrated-dev`;
        - `same-state-message-malformed-payload-dev`;
        - `same-state-message-replay-classification-dev`.
- v0.6.27 commit/push snapshot shows:
    - doc created and committed;
    - README updated and committed;
    - first `git push` failed with recurring authentication issue;
    - second `git push` succeeded;
    - final CarbonStack working tree clean;
    - final repo heads clean except ignored `carbonstack-cypher/cypher.db`.

This is acceptable for a docs/model checkpoint.

Do not overclaim that v0.6.27 added or validated runtime behavior.

## 5. v0.6.27 Model Content Summary

The committed state model defines:

### 5.1 Model vocabulary

- runtime authority;
- server-side coordination authority;
- metadata/evidence;
- secret-bearing state;
- privacy-sensitive state;
- safety-sensitive but not secret;
- generated/disposable state;
- public/documentation state.

### 5.2 Major state inventory decisions

- Comms local state JSON:
    - runtime authority for account/device/server binding;
    - safety-sensitive and possibly privacy-sensitive;
    - no silent replacement/import/restore over newer state.
- OpenMLS sidecar device root:
    - runtime authority container in dev/pre-alpha;
    - secret-bearing and safety-sensitive;
    - no silent regeneration for existing identity/conversation.
- `signer.json`:
    - cryptographic identity/signing authority;
    - secret-bearing;
    - no silent replacement or identity-continuity claim.
- Conversation-level `provider-storage.json`:
    - runtime authority for existing conversation message-open;
    - secret-bearing and safety-sensitive;
    - no silent regeneration/fallback for existing conversation.
- `conversation-summary.json`:
    - metadata/evidence;
    - not runtime authority for existing message-open;
    - no silent promotion to runtime authority.
- Relay envelope metadata:
    - server-side delivery/routing metadata;
    - sender metadata remains unverified;
    - no silent verified-identity claim.
- Cypher DB:
    - server-side coordination authority and delivery/routing authority;
    - not plaintext authority;
    - not verified identity authority.
- `trust.json` and `identity-candidates.json`:
    - safety-sensitive now;
    - future runtime authority once wired into send/receive/enrollment policy;
    - no silent trust promotion or verified import.
- Runner temp roots / build artifacts:
    - generated/disposable only inside scoped dev/build/profile contexts;
    - must not mutate user state.

### 5.3 No-silent laws

The model establishes:

- no silent runtime-authority regeneration;
- no silent runtime-authority replacement;
- no silent verified trust promotion;
- no silent identity continuity after loss/reinstall/restore;
- no silent backup restore over newer state;
- no silent migration;
- no silent algorithm upgrade or downgrade;
- no silent deletion of trust/security-relevant state;
- no ack/drain after failed open or failed join;
- no silent Relay Space membership or routing mutation.

### 5.4 Recovery and degraded-state posture

The model records expected future behavior for:

- missing summary metadata;
- missing conversation provider storage;
- corrupt provider storage;
- missing signer state;
- corrupt Comms local state;
- missing/corrupt trust store;
- stale backup;
- partial restore;
- wrong backup restored;
- Cypher DB loss;
- Cypher DB stale restore;
- Relay Space mismatch;
- migration failure;
- PQ state mismatch;
- generated dev state collision.

### 5.5 Bounded substrate eligibility

The model allows future bounded substrate responsibilities such as:

- state path classification;
- state category labels;
- central path policy;
- refusal behavior;
- warning output;
- non-security-claiming integrity metadata;
- no-silent mutation gates;
- import/export boundaries;
- dev-only substrate tests.

The model forbids default future substrate claims such as:

- encryption;
- secure storage;
- production key protection;
- silent secret regeneration;
- silent backup restore;
- silent migration;
- silent trust promotion;
- silent identity import;
- silent algorithm/PQ upgrade;
- hiding provider-storage failure;
- hiding Cypher DB mismatch;
- hiding Relay Space membership mismatch.

### 5.6 PQ/adversarial dependencies

The model does not define full PQ/hybrid placement.

It only records that future PQ/hybrid work must define:

- attachment layer;
- algorithm version tags;
- conversation state tags;
- identity credential/signature implications;
- provider/signer implications;
- migration epoch;
- compatibility mode;
- recovery impact;
- backup/restore impact;
- no-silent downgrade;
- no-silent upgrade;
- claim boundary.

The model also records future adversarial harness dependencies:

- normal message replay;
- Welcome replay;
- KeyPackage replay;
- server equivocation;
- drop/delay/reorder;
- metadata lies;
- backup rollback;
- state rollback;
- stale provider storage;
- corrupt summary metadata;
- corrupt trust store;
- Relay Space membership mismatch;
- migration downgrade;
- PQ compatibility abuse.

## 6. v0.6.27 Blunders / Continuity Notes

### 6.1 `git add` path typo

The first attempt to stage the model doc used:

    git add 23-v0.6.27-vault-security-nosilent-state-model-v0.md

This failed because the real file was:

    docs/223-v0.6.27-vault-security-nosilent-state-model-v0.md

The correct file was then staged and committed.

Impact:

    harmless staging typo;
    no commit pollution;
    no repo mutation before correction.

Lesson:

    The docs counter is now high enough that missing one leading digit is easy.
    Future scripts or copyable staging snippets should include exact file paths.

### 6.2 Split docs commit

The model doc and README update were committed separately:

    e0027a2 docs: created 223-v0.6.27-vault-security-nosilent-state-model-v0.md
    52295f4 docs: append readme for creation of 223-v0.6.27-vault-security-nosilent-state-model-v0.md

This is acceptable, but future docs-only checkpoints can stage doc + README together unless there is a reason to split.

### 6.3 Recurring push auth failure

The first `git push` failed with the recurring Gitea authentication issue.

A second `git push` succeeded.

Impact:

    no project issue;
    remote main updated from 468432b to 52295f4;
    final branch up to date with origin/main.

Lesson:

    Continue treating first-push auth rc=128/fatal auth as a retryable transport/auth quirk when local commits and final push snapshot are clear.

### 6.4 Validation caveat for docs-only checkpoint

The pasted v0.6.27 commit log does not show a post-doc validation ladder.

This is acceptable for a docs-only checkpoint because v0.6.27A recon already ran baseline validation and the final repo snapshot is clean.

However, do not claim v0.6.27 runtime validation beyond the recon baseline.

Future code-bearing or registry-bearing rungs must run the normal validation ladder.

## 7. Critical Path / Function Updates

New critical doc:

    carbonstack/docs/223-v0.6.27-vault-security-nosilent-state-model-v0.md

Updated doc index:

    carbonstack/docs/README.md

New controlling model for future state/vault work:

    docs/223-v0.6.27-vault-security-nosilent-state-model-v0.md

Existing critical docs still active:

    docs/217-v0.6.22-sidecar-state-authority-classification-v0.md
    docs/220-v0.6.25-normal-message-envelope-metadata-classification-v0.md
    docs/221-v0.6.26-sender-metadata-output-warning-v0.md
    docs/222-v0.6.26-loadcheck-provider-summary-output-v0.md
    CarbonStack_Long_Term_Roadmap_v0.6.26_EVERGREEN.pdf

Critical implementation files unchanged in v0.6.27:

    carbonstack-comms/internal/app/message_wrappers_dev.go
    carbonstack-comms/internal/app/openmls_runtime.go
    carbonstack-comms/internal/app/openmls_bootstrap.go
    carbonstack-comms/internal/protocol/mls/openmls-sidecar/src/state.rs
    carbonstack-cypher/internal/httpapi/api.go
    carbonstack-cypher/internal/db/relay_spaces.go
    carbonstack/tools/carbonstack-validate/*.go
    carbonstack/registry/commands.v0.yaml
    carbonstack/registry/COMMAND_REFERENCE.v0.md

`COMMAND_REFERENCE.v0.md` remains:

    entries=86

No registry change occurred in v0.6.27.

## 8. Next Work / Future To-Do

Immediate next recommended rung:

    v0.6.27C = review/consensus pass on the state model

Purpose:

    inspect `docs/223-v0.6.27-vault-security-nosilent-state-model-v0.md`;
    challenge authority/sensitivity categories;
    verify no-silent laws are neither too weak nor overbroad;
    decide whether the model is sufficient to gate PQ placement and bounded substrate mechanics.

Likely follow-up rungs:

    v0.6.28A = PQ/hybrid placement model recon or draft
    v0.6.29A = bounded substrate/stub mechanics preflight
    later = adversarial harness contract/evidence matrix design
    later = full/full-validate-release/full-runtime-dev naming plan
    later = v0.7.0 release-boundary/package-root rehearsal plan

Do not do next:

    implement production vault;
    implement encryption;
    implement PQ/hybrid ciphersuites;
    implement adversarial harness logic;
    mutate registry for the docs-only model;
    change message-open behavior;
    change ack/drain behavior;
    claim identity verification;
    claim hostile-server safety;
    claim vault/key-storage safety;
    claim production backup/restore;
    pull Welcome/KeyPackage replay implementation into v0.6.x before Relay Space/onboarding lifecycle modeling is ready.

Open questions carried forward:

    What exact future name should the bounded substrate use?
    Should a future state inventory command exist?
    Should metadata/evidence checks be separated from runtime load checks?
    What is the minimum backup manifest format needed before v0.7.x deployability?
    Which state surfaces should be included in the first bounded substrate/stub?
    Should future substrate checks refuse dirty generated state in repo-local sidecar roots?
    Should cypher.db in repo root remain ignored/test-local only, or should a future deployment model force explicit data dirs?
    How should Relay Space membership state be backed up, restored, or reconciled?
    How should stale backup detection work before cryptographic signatures or production vault exist?
    What validation profile should first prove no-silent substrate behavior?

## 9. Current Nonclaims After v0.6.27

v0.6.27 does **not** claim:

    public release;
    public product readiness;
    production secure messaging;
    production E2EE;
    hostile-server safety;
    adversarial relay harness coverage;
    replay safety;
    Welcome replay handling;
    KeyPackage replay handling;
    server equivocation detection;
    metadata privacy;
    metadata authenticity;
    sender authenticity;
    verified sender identity;
    secure enrollment;
    production vault/key storage;
    production backup/restore;
    PQ or hybrid security;
    Android readiness;
    CarbonStackOS readiness;
    deployment readiness;
    audit;
    certification;
    release-package validation;
    package-root validation;
    inclusion in full;
    inclusion in release-snapshot.

## 10. Prior v0.6.26 LogDoc Preserved Below

The remainder of this file preserves the v0.6.26 baseline for continuity.

---

# CarbonStack LogDoc v0.6.26

**Last updated:** 2026-06-20 local session  
**Owner/operator:** bitcrusher32  
**Assistant/model:** GPT-5.5 Thinking  
**Current phase:** Phase 2F -> **v0.6.26 roadmap refresh and output-semantics cleanup checkpoint complete; v0.6.26 EVERGREEN roadmap generated, sender metadata output explicitly marked unverified, and load-check output now distinguishes provider reloadability from summary metadata presence.**  
**Current public release:** `v0.6.0 State/UX Boundary Validation Pre-Release`  
**Current mainline checkpoint:** `v0.6.26 Roadmap + Output Semantics Cleanup Checkpoint`  
**Release URL:** https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.6.0  
**Update source:** `CarbonStackLogDocV0.6.25`, `CarbonStack_BreakpointV0.6.25`, `CarbonStack_Long_Term_Roadmap_v0.6.26_EVERGREEN`, v0.6.26 adversarial Q/A decision matrix, v0.6.26B output-semantics scout, v0.6.26C sender metadata output cleanup log, v0.6.26D load-check provider/summary cleanup logs, and final pushed repo snapshot.  
**Update purpose:** Preserve the v0.6.26 Q/A roadmap refresh and output-semantics cleanup checkpoint; record the sender metadata warning/output fix; record the load-check provider-vs-summary output fix and recovery blunders; update repo heads, critical paths, validation ladder, and future work toward vault/security modeling, no-silent/backup/recovery rules, bounded vault substrate/stub work, PQ/hybrid placement modeling, secured vault/PQ implementation planning, and adversarial harness contract/evidence matrix design; preserve the full v0.6.x timeline ledger; and keep the JSON breakpoint as a lean current-state handoff.  

**Version schema:** `v[scope].[timeline]`. This file is `v0.6.26`, the working branch continuity ledger after the v0.6.26 roadmap refresh and output-semantics cleanup checkpoint in the v0.6.x OpenMLS message-flow integration and boundary-hardening epoch.

**LogDoc V2 usage note:** The Markdown LogDoc is intentionally a dev continuity ledger. It preserves historical v0.6.x timeline data, blunders, validation ladders, command-boundary decisions, and future path constraints. The JSON breakpoint remains a lean current-state handoff and should not try to duplicate the full Markdown timeline.

**Privacy / history-rewrite note:** Exact raw shell logs may contain local usernames, hostnames, email addresses, and old pre-rewrite author identities. This LogDoc avoids repeating local host-derived identity strings except where operationally necessary. Public commit identity remains normalized on public `main`/tag histories. Pre-v0.6.13 commit hashes from older LogDocs and historical docs are now historical **pre-rewrite** references and must not be treated as current Gitea `main`/tag SHAs.

## 0. Executive Summary

v0.6.26 is complete as a roadmap-refresh plus output-semantics cleanup checkpoint.

It did six bounded things:

1. **Generated the v0.6.26 EVERGREEN roadmap refresh**
   - The v0.6.26 roadmap became the active forward-planning authority after the v0.6.26 adversarial Q/A decision matrix.
   - It explicitly keeps CarbonStack in pre-alpha / experimental planning mode.
   - It preserves the key position that v0.7.0 is a cumulative pre-alpha engineering release boundary, not a production or secure-use release.
   - It defines the relative lane roadmap:
     - remaining v0.6.x decisions/modeling;
     - v0.7.x deployability and state-security integration;
     - v0.8.x PQ/deployability maturation, migration/compatibility, code-health optimization, operator docs, and release-candidate hardening;
     - v0.9.x full self-pen-test / adversarial validation across real operational surfaces.
   - Exact future breakpoint numbers remain LogDoc-level decisions, not hardline evergreen-roadmap promises.

2. **Ran v0.6.26B output-semantics scout**
   - The scout confirmed two release-boundary clarity gaps:
     - normal message inbox output still printed `from_device` without an explicit unverified sender metadata warning;
     - `openmls-conversation-load-check-dev` still conflated provider reloadability and summary metadata presence in output.
   - Scout validation passed:
     - generated command reference check stayed current with `entries=86`;
     - runner `go test ./...` passed;
     - registry missing-nonclaims scan passed;
     - runner doctor passed;
     - sidecar `cargo test --quiet` passed warning-free;
     - Comms `go test ./...` passed;
     - post-scout replay and malformed-payload profiles passed.
   - Scout result summary:
     - sender grep: 723 lines, 24 `from_device` hits;
     - load-check grep: 1801 lines, 105 `conversation-summary` hits, 200 `group_reloadable` hits;
     - tests grep: 223 lines;
     - live probe showed `from_device` present but no `sender_identity_verified`, `from_device_unverified`, or `sender_device_id_unverified`.

3. **Landed v0.6.26C sender metadata output warning cleanup**
   - Comms commit:
     - `61b2e51 fix: label message sender metadata as unverified`
   - CarbonStack docs commit:
     - `499ab6a docs: record sender metadata output warning`
   - Added:
     - `docs/221-v0.6.26-sender-metadata-output-warning-v0.md`
   - Updated:
     - `docs/README.md`
   - Normal message inbox outputs now preserve the legacy field:
     - `from_device: <sender_device_id>`
   - They also add the explicit trust-boundary fields:
     - `from_device_unverified: <sender_device_id>`
     - `sender_identity_verified: false`
     - `warning: from_device is relay envelope metadata, not verified identity`
   - This applies to normal application-message inbox/open outputs:
     - `message-inbox-dev`
     - `openmls-inbox-dev`
   - It does not redefine Relay onboarding artifact sender fields.

4. **Landed v0.6.26D load-check provider/summary output cleanup**
   - Comms commit:
     - `2d13c45 fix: clarify load-check provider metadata output`
   - CarbonStack docs commit:
     - `468432b docs: record load-check provider summary output`
   - Added:
     - `docs/222-v0.6.26-loadcheck-provider-summary-output-v0.md`
   - Updated:
     - `docs/README.md`
   - Successful `openmls-conversation-load-check-dev` output now distinguishes:
     - `group_reloadable: true`
     - `provider_reloadable: true`
     - `summary_metadata_present: true`
     - `provider_storage_present: true`
     - `summary_metadata_path: <path>`
     - `provider_storage_path: <path>`
     - `summary_metadata_warning: none`
   - Missing summary metadata now emits an explicit metadata-missing block before returning failure:
     - `status: metadata_missing`
     - `provider_reloadable: not_evaluated_by_load_check`
     - `summary_metadata_present: false`
     - `provider_storage_present: true/false`
     - `summary_metadata_warning: conversation-summary metadata is missing; conversation-load-check-dev is stricter than message-open and cannot confirm reloadability without summary metadata`
   - This keeps the v0.6.22/v0.6.24 model intact:
     - provider storage remains runtime authority for existing message-open;
     - summary metadata remains metadata/evidence;
     - load-check can be stricter than message-open, but now says so clearly.

5. **Recovered from two v0.6.26D implementation failures**
   - The first v0.6.26D script failed mid-run because the patch incorrectly skipped success-path output insertion after seeing `provider_reloadable` in the error path.
   - The first recovery inserted success-path fields but printed literal `\n` sequences, collapsed several output fields onto one line, and overpatched add-member/join tests that should not expect load-check-only fields.
   - The clean reapply reset only the partial v0.6.26D files, preserved v0.6.26C, reapplied a narrower patch, patched only the load-check unit test, and succeeded.
   - Final clean live probe passed:
     - `baseline_output_has_provider_summary_fields: yes`
     - `missing_summary_output_has_metadata_missing_fields: yes`
     - `missing_summary_rc: 1`

6. **Preserved claim boundaries and next-rung direction**
   - v0.6.26 is not a public release.
   - Current public release remains `v0.6.0`.
   - No registry entry was added.
   - `COMMAND_REFERENCE.v0.md` remained at `entries=86`.
   - `full` and `release-snapshot` remain unchanged.
   - No message-open behavior changed.
   - No ack/drain behavior changed.
   - No provider-state behavior changed.
   - No identity verification was implemented.
   - No vault/PQ/adversarial harness work was implemented.
   - The next safest mainline work is the v0.6.x model lane:
     - vault/security model;
     - no-silent regeneration/replacement/import/delete/migration/restore rules;
     - backup/recovery model;
     - bounded vault substrate/stub decision;
     - PQ/hybrid placement model;
     - adversarial harness contract/evidence matrix;
     - v0.7.0 release/package-root rehearsal plan.

Current commits:

    carbonstack        468432b docs: record load-check provider summary output
    carbonstack-comms  2d13c45 fix: clarify load-check provider metadata output
    carbonstack-cypher d18a564 chore: restore validated Go module floor
    carbonstack-os     1bbbe52 docs: clarify CarbonStackOS target direction

Current public release remains:

    v0.6.0 State/UX Boundary Validation Pre-Release

This is **not** a public product release.

### 0.1 Carried-forward v0.6.25 baseline summary

v0.6.25 is complete as a docs/model checkpoint for normal application-message envelope metadata classification.

It did six bounded things:

1. **Ran malformed typed/envelope metadata recon**
   - v0.6.25A tested seven normal application-message envelope metadata mutations after valid same-state Relay onboarding and `message-send-dev`.
   - Mutation families:
     - `payload_sha256` mismatch while ciphertext remained valid;
     - `payload_size_bytes` mismatch while ciphertext remained valid;
     - invalid `protocol_version` while ciphertext remained valid;
     - wrong `sender_device_id` while ciphertext remained valid;
     - empty `sender_device_id` while ciphertext remained valid;
     - wrong `recipient_device_id` while ciphertext remained valid;
     - invalid `delivery_state` while ciphertext remained valid.
   - The final recon summary reported:
     - `cases_total: 7`;
     - `classification_count detected_or_rejected_no_ack_no_drain_recovery_ok: 3`;
     - `classification_count metadata_not_runtime_enforced_open_ack_drain: 2`;
     - `classification_count metadata_routing_suppressed_no_open_no_ack_recovery_ok: 1`;
     - `classification_count delivery_state_suppressed_no_open_no_ack_recovery_ok: 1`;
     - `opened_or_acked_cases: 2`;
     - `suppressed_or_rejected_cases: 5`;
     - `bookkeeping_concern_cases: 0`;
     - `unsafe_or_unexpected_state_mutation_cases: 0`;
     - `provider_changed_by_attempt_cases: 2`.

2. **Classified enforced metadata and support gates**
   - `payload_sha256` mismatch failed before open/ack/drain and recovered after restore.
   - `payload_size_bytes` mismatch failed before open/ack/drain and recovered after restore.
   - invalid `protocol_version` was skipped/unsupported without open/ack/drain and recovered after restore.
   - These are currently enforced/rejected before successful normal message-open.

3. **Classified routing and delivery-state suppression**
   - wrong `recipient_device_id` caused Bob's normal inbox query to see zero queued envelopes; restore allowed correct open/ack recovery.
   - invalid `delivery_state` caused the normal inbox query to see zero queued envelopes; restore allowed correct open/ack recovery.
   - These are fetch/routing/bookkeeping selectors, not cryptographic proof of intended recipient or hostile-server safety.

4. **Classified sender metadata as unverified relay metadata**
   - wrong `sender_device_id` and empty `sender_device_id` both opened and acked.
   - Output `from_device` reflected the mutated envelope metadata.
   - This is not an identity-verification regression because identity verification is already a nonclaim.
   - It is now explicitly documented as a trust/display boundary:
     - `from_device` in current `message-inbox-dev` output is relay envelope sender metadata;
     - it must not be treated as verified identity.

5. **Landed docs/model checkpoint only**
   - Added:
     - `docs/220-v0.6.25-normal-message-envelope-metadata-classification-v0.md`
   - Updated:
     - `docs/README.md`
   - No runner profile was added.
   - No registry entry was added.
   - `COMMAND_REFERENCE.v0.md` remained at:
     - `entries=86`.
   - No Comms/Cypher/OS source changed.
   - `full` and `release-snapshot` remained unchanged.
   - Commit:
     - `b7f89ea docs: classify normal message envelope metadata`.

6. **Preserved claim boundaries and next-rung direction**
   - v0.6.25 is not a public release.
   - Current public release remains `v0.6.0`.
   - This checkpoint does not claim sender authenticity, identity verification, metadata authenticity, hostile-server safety, replay safety, adversarial harness coverage, vault/key-storage safety, PQ/hybrid security, or production secure messaging.
   - Recommended next sequence:
     - adversarial QA for evergreen roadmap refresh;
     - vault/security model plus no-silent/backup/recovery rules;
     - bounded vault substrate/stub only after model;
     - PQ/hybrid placement model;
     - secured vault/PQ implementation planning;
     - adversarial harness.
   - Welcome/KeyPackage malformed/replay remains deferred as onboarding-artifact adversarial/failure work.

Current commits:

    carbonstack        b7f89ea docs: classify normal message envelope metadata
    carbonstack-comms  b42aa09 fix: make Welcome join state writes atomic
    carbonstack-cypher d18a564 chore: restore validated Go module floor
    carbonstack-os     1bbbe52 docs: clarify CarbonStackOS target direction

Current public release remains:

    v0.6.0 State/UX Boundary Validation Pre-Release

This is **not** a public product release.

### 0.2 Carried-forward v0.6.24 baseline summary

v0.6.24 is complete as a same-state normal application-message replay classification profile checkpoint plus load-check semantic docs checkpoint.

It did seven bounded things:

1. **Ran combined v0.6.24A recon**
   - Reproduced the known `conversation-load-check-dev` / `message-inbox-dev` semantic mismatch around missing `conversation-summary.json`.
   - Classified normal application-message duplicate/replay behavior after valid same-state Relay onboarding and message send.
   - Recon reported:
     - `LOADCHECK_SEMANTIC_RECON_RESULT cases_total: 1`;
     - `REPLAY_CLASSIFICATION_RECON_RESULT cases_total: 5`;
     - `unsafe_cases: 0`;
     - `bookkeeping_concern_cases: 0`;
     - `tolerated_or_unknown_cases: 0`;
     - `detected_or_rejected_cases: 3`;
     - `storage_rejected_cases: 1`;
     - `delivery_state_suppressed_cases: 1`.

2. **Recorded the load-check summary-metadata semantic mismatch**
   - Missing `conversation-summary.json` caused `conversation-load-check-dev` to fail.
   - `message-inbox-dev --ack` still opened and acked a valid queued message when conversation-level provider storage remained valid.
   - Classification:
     - `load_check_stricter_than_message_open_summary_metadata_missing`
   - Current model remains:
     - conversation provider storage is runtime authority;
     - `conversation-summary.json` is metadata/evidence for existing message-open;
     - this is command-semantics/documentation cleanup, not a no-ack/no-drain runtime bug.

3. **Classified normal application-message duplicate/replay behavior**
   - Same envelope ID duplicate insert:
     - storage/DB uniqueness rejected the duplicate;
     - original message recovery open/ack succeeded;
     - classification: `storage_rejected_duplicate_same_envelope_id_original_recovery_ok`.
   - Same envelope after ack without requeue:
     - delivery state suppressed re-open;
     - classification: `delivery_state_suppressed_same_envelope_after_ack`.
   - Same envelope manually requeued after ack:
     - second open failed with OpenMLS `SecretReuseError` family;
     - no second ack;
     - no drain;
     - classification: `detected_or_rejected_no_ack_no_drain_manual_requeue_same_envelope`.
   - Duplicate same ciphertext under a new envelope ID before original ack:
     - first valid envelope opened/acked;
     - duplicate failed open with `SecretReuseError` family;
     - no second ack;
     - no drain;
     - classification: `detected_or_rejected_no_ack_no_drain_duplicate_payload_new_envelope_before_ack`.
   - Duplicate same ciphertext under a new envelope ID after original ack:
     - duplicate failed open with `SecretReuseError` family;
     - no extra ack;
     - no drain;
     - classification: `detected_or_rejected_no_ack_no_drain_duplicate_payload_new_envelope_after_ack`.

4. **Promoted recon into a committed live runner profile**
   - Added:
     - `same-state-message-replay-classification-dev`
     - `runner.same-state-message-replay-classification-dev`
   - Added:
     - `tools/carbonstack-validate/same_state_message_replay_classification_dev.go`
     - `docs/219-v0.6.24-loadcheck-replay-classification-v0.md`
   - Updated:
     - `tools/carbonstack-validate/main.go`
     - `tools/carbonstack-validate/README.md`
     - `registry/commands.v0.yaml`
     - `registry/COMMAND_REFERENCE.v0.md`
     - `docs/README.md`

5. **Recovered from one runner implementation blunder**
   - Initial v0.6.24B implementation failed during `runner go test ./...` because `same_state_message_replay_classification_dev.go` referenced non-existent helper `runCommandCombined`.
   - Recovery added a local `replayRunPythonScript` helper using `os/exec` and replaced both bad calls.
   - This was a runner porting/implementation bug, not a replay runtime failure.

6. **Landed one profile/registry/docs commit**
   - Commit:
     - `b7f89ea docs: classify normal message envelope metadata`
   - `COMMAND_REFERENCE.v0.md` regenerated to:
     - `entries=86`
   - Scripted push succeeded:
     - `943e0e4..1c421e2 main -> main`

7. **Preserved claim boundaries and next-rung direction**
   - No Comms/Cypher/OS source changed.
   - No runtime replay patch was needed.
   - `full` and `release-snapshot` remain unchanged.
   - This is deterministic replay classification, not replay safety.
   - Current public release remains `v0.6.0`.
   - Recommended next sequence:
     - remaining malformed typed cases;
     - adversarial QA for evergreen roadmap refresh;
     - vault/security model plus no-silent/backup/recovery rules;
     - bounded vault substrate/stub;
     - PQ/hybrid placement model;
     - secured vault/PQ implementation planning;
     - adversarial harness.

Current commits:

    carbonstack        b7f89ea docs: classify normal message envelope metadata
    carbonstack-comms  b42aa09 fix: make Welcome join state writes atomic
    carbonstack-cypher d18a564 chore: restore validated Go module floor
    carbonstack-os     1bbbe52 docs: clarify CarbonStackOS target direction

Current public release remains:

    v0.6.0 State/UX Boundary Validation Pre-Release

This is **not** a public product release.


### 0.3 Carried-forward v0.6.23 baseline summary

v0.6.23 completed the same-state malformed normal-message payload profile checkpoint.

It added:

    same-state-message-malformed-payload-dev
    runner.same-state-message-malformed-payload-dev

It proved six malformed normal application-message payload mutation families fail without false open, ack, inbox drain, provider-state mutation, or envelope rewrite, and that restoring the original payload allows correct open/ack recovery.

v0.6.23 landed:

    943e0e4 test: add same-state malformed payload profile

`COMMAND_REFERENCE.v0.md` moved to:

    entries=85

The initial profile implementation failed because the app-envelope SQL lookup incorrectly required `relay_space_id=sub.RelaySpace`. Recovery removed that predicate and used recipient device + normal application-message content type + queued delivery state. That blunder remains important because normal application-message envelopes are not always Relay-Space-scoped the same way onboarding artifacts are.

### 0.4 Carried-forward v0.6.22 baseline summary

v0.6.22 is complete as a docs/model checkpoint.

It did five bounded things:

1. **Classified stale sidecar/conversation/provider state authority**
   - v0.6.22A/B recon tested missing, invalid, truncated, and moved sidecar-generated state around an already joined same-state conversation.
   - The final classification recon reported:
     - `cases_total: 12`;
     - `critical_runtime_bypass_cases: 0`;
     - `cases_with_target_rewrite_or_regeneration: 0`.
   - The key conclusion is:
     - conversation-level `provider-storage.json` is runtime authority for existing conversation message-open;
     - the final conversation directory is the runtime authority container;
     - `conversation-summary.json` is metadata/evidence for existing message-open, but `conversation-load-check-dev` currently treats a missing summary as failure;
     - `join-summary.json` is metadata/evidence;
     - device-level `provider-storage.json` is identity/bootstrap substrate for existing message-open purposes, not existing-conversation runtime message-open authority.

2. **Confirmed the stale-state safety boundary that matters most**
   - Runtime-authority mutations blocked message-open/ack/drain:
     - deleting conversation provider storage;
     - invalid conversation provider storage;
     - truncated conversation provider storage;
     - moved/missing final conversation directory.
   - No tested load-check or message-inbox path silently regenerated or rewrote the mutated target.
   - This means v0.6.22 did **not** reveal a critical runtime-authority bypass.

3. **Explained the load-check/message-open semantic mismatch**
   - Deleting `conversation-summary.json` makes `conversation-load-check-dev` fail with missing summary.
   - `message-inbox-dev --ack` can still open and ack a valid queued message if conversation-level provider storage remains usable.
   - Current model:
     - this is a command semantics/documentation issue, not a no-ack/no-drain bug;
     - future cleanup should clarify whether load-check means provider reloadability, metadata presence, or both.

4. **Landed one docs/model commit**
   - Added:
     - `docs/217-v0.6.22-sidecar-state-authority-classification-v0.md`
   - Updated:
     - `docs/README.md`
   - Commit:
     - `8dd7db5 docs: classify sidecar state authority`
   - No Comms/Cypher/OS source changed.
   - No registry entry was added.
   - No runner profile was added.
   - `COMMAND_REFERENCE.v0.md` remained at 84 entries.
   - `full` and `release-snapshot` remained unchanged.

5. **Preserved the next-rung direction**
   - The recommended next rung is malformed normal-message payload recon unless load-check output normalization is intentionally pulled forward first.
   - The planned sequence remains:
     - malformed normal-message payload recon;
     - replay/duplicate classification;
     - longer adversarial QA;
     - roadmap refresh;
     - remaining v0.6.x decisions;
     - v0.7.0 minor-epoch release definition;
     - expanded v0.7.x+ evergreen goals.

Current commits:

    carbonstack        943e0e4 test: add same-state malformed payload profile
    carbonstack-comms  b42aa09 fix: make Welcome join state writes atomic
    carbonstack-cypher d18a564 chore: restore validated Go module floor
    carbonstack-os     1bbbe52 docs: clarify CarbonStackOS target direction

Current public release remains:

    v0.6.0 State/UX Boundary Validation Pre-Release

This is **not** a public product release.

### 0.5 Carried-forward v0.6.21 baseline summary

v0.6.21 completed the same-state Welcome join failure profile checkpoint.

It added:

    same-state-welcome-join-failure-dev
    runner.same-state-welcome-join-failure-dev

It proved the v0.6.20 atomic Welcome join state-write fix in a live same-state runner profile:

    corrupt Welcome join fails
    -> no Welcome ack
    -> no Relay inbox drain
    -> no final/staging Bob conversation state poison
    -> restored valid Welcome joins and acks successfully with the same conversation label

v0.6.22 leaves the v0.6.21 runner/profile/registry surface unchanged.

### 0.6 Carried-forward v0.6.20 baseline summary

v0.6.20 completed the Welcome join partial-state safety fix.

It did not add a runner profile. It fixed the defect that made a corrupt Welcome join leave durable final-path Bob conversation state.

Core v0.6.20 sequence:

    corrupt Welcome join
    -> no ack
    -> no Relay inbox drain
    -> but restored valid Welcome failed with conversation_already_exists
    -> deep recon localized final-path Bob conversation state poisoning
    -> sidecar join_dev_conversation staged state and promoted only after successful materialization/reload-check
    -> warning recovery removed Rust warning that poisoned sidecar JSON parsing
    -> Comms and existing same-state profiles passed

v0.6.21 builds directly on this by proving the fixed behavior live and adding the leaf runner profile.

### 0.7 Carried-forward v0.6.19 baseline summary

v0.6.19 completed the third same-state normal-message failure-path profile.

It added:

    same-state-message-recipient-failure-dev
    runner.same-state-message-recipient-failure-dev

It proved wrong recipient/device/sidecar no-false-success/no-ack/no-drain behavior:

    Relay join -> normal message send -> wrong recipient/device/sidecar attempts
    -> no false message-open success
    -> no ack
    -> no Bob inbox drain
    -> correct Bob open/ack still succeeds

v0.6.21 leaves this profile unchanged.

### 0.8 Carried-forward v0.6.18 baseline summary

v0.6.18 completed the second same-state failure-path profile.

It added:

    same-state-message-unsupported-dev
    runner.same-state-message-unsupported-dev

It proved unsupported normal application-message content type no-ack/no-drain behavior:

    Relay join -> normal message send -> mutate normal message content_type
    -> message-inbox-dev --ack
    -> unsupported skip
    -> no ack
    -> no inbox drain
    -> restore valid content_type
    -> correct open/ack still succeeds

v0.6.21 leaves this profile unchanged.

### 0.9 Carried-forward v0.6.17 baseline summary

v0.6.17 completed the first same-state failure-path profile.

It added:

    same-state-message-failure-dev
    runner.same-state-message-failure-dev

It proved wrong-conversation normal-message open no-ack/no-drain behavior:

    Relay join -> normal message send -> wrong-conversation message-inbox-dev --ack
    -> message-open failure
    -> no ack
    -> no inbox drain
    -> correct conversation open/ack still succeeds

v0.6.21 leaves this profile unchanged.

### 0.10 Carried-forward v0.6.16 baseline summary

v0.6.16 completed the positive-path same-state integrated profile.

It added:

    same-state-integrated-dev
    runner.same-state-integrated-dev

It proved Level 4 same-conversation positive-path continuity:

    Relay KeyPackage -> add-member -> Welcome -> join -> message-send-dev -> message-inbox-dev --ack

inside one coherent live-dev temp universe with one Cypher server/DB, Alice/Bob Comms states, device IDs, sidecar labels, conversation label, and Bob inbox empty after ack.

v0.6.21 still does not promote this into `full` or `release-snapshot`.

### 0.11 Carried-forward v0.6.15 baseline summary

v0.6.15 completed the docs/planning breakpoint that made v0.6.16 safe.

It recorded that current `integrated-runtime-dev` was honest sequential composition, same-state proof should be proven before naming a profile, the target proof level should be Level 4 same conversation proof, and `full` / `release-snapshot` must remain untouched.

### 0.12 Carried-forward v0.6.14 baseline summary

v0.6.14 remains the post-history-rewrite sanity and generated-reference freshness guard checkpoint.

It verified public main/tag identity hygiene, v0.6.0 release asset sanity, attached release package validation, generated command-reference determinism, registry/reference consistency, and private all-ref bundle freeze before deleting local backup branches.

## 1. Current Project Goal

**Active goal:** Continue the v0.6.x OpenMLS message-flow integration and boundary-hardening epoch from the v0.6.26 roadmap/output-semantics checkpoint into the state/security modeling lane defined by the v0.6.26 EVERGREEN roadmap.

After v0.6.26, the project has:

- the v0.6.0 public pre-release still live on Gitea;
- attached v0.6.0 release assets already revalidated after the v0.6.13 tag rewrite;
- public `main`/tag identity hygiene verified after the history rewrite;
- generated `registry/COMMAND_REFERENCE.v0.md`;
- deterministic generated-reference freshness enforced by runner tests;
- registry entry count remains at 86;
- legacy stub commands gated behind explicit opt-in;
- opinionated `message-send-dev` / `message-inbox-dev` wrappers as the recommended dev/pre-alpha normal-message path;
- direct `openmls-*` commands retained as lower-level implementation/proof surfaces;
- Relay onboarding boundaries separated from ordinary message inbox behavior;
- `same-state-integrated-dev`, proving Level 4 same conversation positive-path continuity:
    - Relay onboarding -> joined conversation -> normal message send/open/ack;
- deterministic normal-message failure/classification profiles:
    - wrong conversation;
    - unsupported content type;
    - wrong recipient/device/sidecar;
    - malformed payload/ciphertext;
    - normal application-message replay/duplicate classification;
- Welcome join partial-state safety and `same-state-welcome-join-failure-dev`;
- v0.6.22 state-authority classification:
    - conversation-level provider storage and final conversation directory are runtime authority;
    - summary files are metadata/evidence for existing message-open;
    - device-level provider storage is identity/bootstrap substrate for existing message-open purposes;
- v0.6.25 normal-message envelope metadata classification:
    - payload hash/size and protocol support gates are enforced/rejected before successful open/ack;
    - recipient/delivery-state metadata acts as routing/bookkeeping selection;
    - sender metadata / `from_device` is unverified relay envelope metadata;
- v0.6.26 sender metadata output warning cleanup:
    - `from_device_unverified`;
    - `sender_identity_verified: false`;
    - explicit warning that `from_device` is relay envelope metadata, not verified identity;
- v0.6.26 load-check output cleanup:
    - `provider_reloadable`;
    - `summary_metadata_present`;
    - `provider_storage_present`;
    - metadata-missing output that distinguishes load-check failure from message-open authority;
- v0.6.26 EVERGREEN roadmap refresh:
    - remaining v0.6.x modeling and release-boundary decisions;
    - v0.7.x deployability and state-security integration;
    - v0.8.x maturation/code-health/release-candidate hardening;
    - v0.9.x full adversarial validation across real operational surfaces.

CarbonStack's broader intent remains:

- **CarbonStack** is the secure-communications appliance-stack project and shared doctrine/release authority.
- **CarbonStackComms** is the text-first messaging client and current OpenMLS sidecar / Relay Space integration surface.
- **CarbonStackCypher** is the self-hostable relay/storage server/API stack.
- **CarbonStackOS** is future constrained Android-derived appliance OS work and remains out of current runnable validation packages.
- The `carbonstack` Gitea repo remains the source of truth for release state, release assets, docs, registry, command-boundary classification, and public claims.
- GitHub mirrors, if present, are secondary push mirrors for discoverability/redundancy only.

**North star:**

> Its goal is not to be a better smartphone. Its goal is to stop being a smartphone.

**Core doctrine:**

> Every feature is guilty until it proves it does not add unacceptable parser, network, sensor, identity, or filesystem authority.

**v0.6.x engineering doctrine after v0.6.26:**

> The normal OpenMLS message-flow spine and deterministic normal-message failure-hardening lane are mostly complete for dev/pre-alpha boundaries. The next honest work is state/security modeling: vault/security model, no-silent/backup/recovery rules, bounded vault substrate/stub decision, PQ/hybrid placement model, and adversarial harness contract/evidence matrix design. Output clarity improvements do not grant production security, identity verification, hostile-server safety, replay safety, vault/key-storage safety, PQ security, or public-user readiness.

## 2. Current Public Release and Mainline State

Current public release remains:

    https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.6.0

Release title:

    CarbonStack v0.6.0 State/UX Boundary Validation Pre-Release

Release tag:

    v0.6.0

Current post-rewrite tag target:

    70d4318 v0.6.0

Pre-rewrite v0.6.0 tag target for historical context only:

    7e676cd6ad / 7e676cd scripts: add v0.6.0 package rehearsal

Current mainline checkpoint:

    v0.6.26 Roadmap + Output Semantics Cleanup Checkpoint

Current CarbonStack head:

    468432b docs: record load-check provider summary output

Current CarbonStackComms head:

    2d13c45 fix: clarify load-check provider metadata output

v0.6.26 is not a new public release and does not change the release title. It is a post-release roadmap/docs/output-semantics checkpoint.

Important release-page warning retained:

    Gitea automatically provides Source Code ZIP/TAR.GZ archives for the carbonstack repo at the tag.
    These are not the intended runnable multi-repo validation package.
    Use the attached v0.6.0 package, manifest, checksums, validation freeze, release notes, testing runbook, and LICENSE.

v0.6.26 did not rebuild v0.6.0 release assets and did not change package-root release validation.

The new v0.6.26 roadmap/docs are:

    CarbonStack_Long_Term_Roadmap_v0.6.26_EVERGREEN.pdf
    docs/221-v0.6.26-sender-metadata-output-warning-v0.md
    docs/222-v0.6.26-loadcheck-provider-summary-output-v0.md

They remain excluded from:

    full
    release-snapshot
    release package validation
    package-root validation
    adversarial relay harness claims
    replay-safety claims
    hostile-server safety claims
    identity verification claims
    vault/PQ claims

## 3. Current Repo Heads After v0.6.26

Final pushed heads:

    carbonstack        468432b docs: record load-check provider summary output
    carbonstack-comms  2d13c45 fix: clarify load-check provider metadata output
    carbonstack-cypher d18a564 chore: restore validated Go module floor
    carbonstack-os     1bbbe52 docs: clarify CarbonStackOS target direction

Recent pushed history:

    carbonstack:
        468432b docs: record load-check provider summary output
        499ab6a docs: record sender metadata output warning
        b7f89ea docs: classify normal message envelope metadata
        1c421e2 test: add same-state replay classification profile
        943e0e4 test: add same-state malformed payload profile
        8dd7db5 docs: classify sidecar state authority
        3876fc8 test: add same-state Welcome join failure profile
        e1f6b8c docs: record Welcome join partial-state fix

    carbonstack-comms:
        2d13c45 fix: clarify load-check provider metadata output
        61b2e51 fix: label message sender metadata as unverified
        b42aa09 fix: make Welcome join state writes atomic
        53779d5 test: cover message wrapper residual edge cases
        462f796 chore: group Comms CLI help by command boundary
        020c07f test: align wrapper smoke with message wrappers
        9391029 feat: add opinionated OpenMLS message wrappers
        e413ca0 test: cover OpenMLS message flow command edge cases

    carbonstack-cypher:
        d18a564 chore: restore validated Go module floor
        52e3673 docs: make Cypher surface evergreen
        3589dbe docs: define local Cypher validation state contract
        15760c6 docs: record provider OpenMLS join boundary
        09afa0b feat: add Relay Space scoped envelope routes
        384affd feat: add Relay Space HTTP API routes
        cc6e589 feat: add Relay Space DB helpers
        6798d11 feat: add Relay Space schema substrate

    carbonstack-os:
        1bbbe52 docs: clarify CarbonStackOS target direction
        f984cdf Upload files to "/"
        c764790 chore: fix readme formatting
        953b006 chore: fix readme formatting
        b537475 Add CarbonStackOS north star and initial appliance model
        dab2792 Initial CarbonStack repository structure

Working tree note:

    carbonstack and carbonstack-os were clean in the final snapshot.
    carbonstack-comms retained known ignored generated sidecar roots:
        !! internal/protocol/mls/openmls-sidecar/.carbonstack-openmls-sidecar-state/
        !! internal/protocol/mls/openmls-sidecar/target/
    carbonstack-cypher retained the known ignored local artifact:
        !! cypher.db

## 4. Preserved v0.6.x Timeline Ledger

This section preserves the v0.6.x historical chain. Exact commit hashes before v0.6.13 are partially invalidated by the author-history rewrite; use the post-rewrite heads and tag values above for current Gitea references.

### 4.1 v0.6.0 — State/UX Boundary Validation Pre-Release

v0.6.0 cut the current public pre-release.

Current post-rewrite tag:

    v0.6.0 -> 70d4318

Pre-rewrite tag reference:

    7e676cd6ad / 7e676cd scripts: add v0.6.0 package rehearsal

v0.6.14 recon confirmed:

    Gitea release API reachable.
    v0.6.0 remains a prerelease.
    asset_count=8.
    all attached assets downloaded.
    attached asset checksum file verified all 8 assets.
    downloaded attached package passed verify-checksums.
    downloaded attached package passed full --clean-generated.

Important v0.6.0 facts remain:

    carbonstack-os is deliberately excluded from the runnable package.
    relay-openmls-join-dev remains manual/dev-only and excluded from full/release-snapshot.
    full remains a release-package validation ladder, not mature integrated runtime validation.
    registry presence is not promotion.
    attached v0.6.0 release assets are the intended runnable multi-repo validation surface.
    Gitea auto source archives are not the intended runnable multi-repo package.

v0.6.0 did not claim production secure messaging, hostile-server safety, metadata privacy, verified identity, secure enrollment, production vault/key storage, PQ/hybrid security, Android readiness, CarbonStackOS readiness, audit, certification, local-backbone, deployment readiness, or mature messenger UX.

### 4.2 v0.6.1 — Message-surface preflight

v0.6.1 recorded post-release message-surface direction.

Preserved decision:

    openmls-send-dev and openmls-inbox-dev were the normalized immediate message-flow baseline at that time.
    legacy send/inbox/ack remained warning-gated stubs.
    ack should not happen on fetch.
    ack should happen only after successful message-open and explicit --ack.
    Relay KeyPackage/Welcome/add-member/join remained separate onboarding/artifact primitives.

### 4.3 v0.6.2 — Legacy stub demotion boundary

v0.6.2 required explicit opt-in for legacy stub-era message commands.

Preserved decision:

    send/inbox/ack require --allow-legacy-stub.
    OpenMLS inbox ack and Relay-scoped ack paths remain unaffected.
    Legacy stubs are continuity surfaces, not mature messaging UX.

Blunder carried forward:

    The first patch failed safely.
    A second patch malformed YAML and the runner caught it.
    The repaired patch validated cleanly.

### 4.4 v0.6.3 — Normalized OpenMLS message-flow contract

v0.6.3 documented the normalized OpenMLS message flow contract.

Preserved decision:

    openmls-send-dev / openmls-inbox-dev were still the canonical dev-only direct surface at that time.
    message wrappers were deferred until v0.6.5.
    Relay KeyPackage/Welcome/add-member/join remained onboarding/artifact primitives.
    Ack rule was locked: no ack on fetch or failed open; explicit ack only after successful open.

### 4.5 v0.6.4 — Command-level OpenMLS edge tests

v0.6.4 added command-level test coverage for direct OpenMLS runtime behavior.

Preserved tested edges included:

    submit failure after protect;
    empty artifact handling;
    stable output;
    empty inbox behavior;
    artifact write failure;
    ack failure after open;
    limit/generated label behavior.

Scope:

    Tests only.
    No runtime implementation change.
    No registry/docs change.

### 4.6 v0.6.5 — Opinionated normal-message wrappers

v0.6.5 added wrapper commands for the recommended dev/pre-alpha normal-message path.

Commands added:

    message-send-dev
    message-inbox-dev

Preserved wrapper interpretation:

    message-* wraps the OpenMLS sidecar + Cypher application-message path.
    message-* is dev/pre-alpha public-test UX candidate surface.
    openmls-* remains lower-level implementation/proof.
    Relay onboarding remains separate.

### 4.7 v0.6.6 — Wrapper smoke/profile alignment

v0.6.6 aligned the wrapper smoke profile and registry wording.

Preserved proof strings:

    dev-runtime-openmls-wrappers:
        openmls-*-dev bootstrap wrappers -> message-send-dev -> Cypher -> message-inbox-dev --ack

    dev-runtime-openmls:
        openmls-send-dev -> Cypher -> openmls-inbox-dev --ack

Preserved decision:

    dev-runtime-openmls-wrappers remains separate from full and release-snapshot.
    direct openmls remains the baseline direct proof.
    wrapper profile is not production messaging UX.

### 4.8 v0.6.7 — Operator surface compression + registry lookup

v0.6.7 improved operator clarity.

Preserved changes:

    Comms help grouped commands by boundary.
    legacy send/inbox/ack warnings point toward message-* wrappers.
    openmls-* is lower-level direct proof / transition candidate, not legacy.
    registry-lookup can inspect by registry ID or literal command.
    lifecycle_status vocabulary added:
        recommended_dev_wrapper
        lower_level_direct_proof_transition_candidate
        legacy_stub_explicit_opt_in

### 4.9 v0.6.8 — Residual wrapper tests + registry lookup enrichment + public asset verification

v0.6.8 closed residual wrapper behavior tests and verified the public v0.6.0 Gitea release asset surface.

Preserved validation meaning:

    attached v0.6.0 multi-repo package and companion assets were downloaded;
    checksums validated;
    fresh extracted package passed verify-checksums and full --clean-generated;
    Gitea auto source archives remained explicitly not the runnable release package.

v0.6.14 repeated the critical public release asset/package sanity after the v0.6.13 tag rewrite.

### 4.10 v0.6.9 — Relay onboarding boundary contract

v0.6.9 classified Relay onboarding as an artifact/onboarding bridge, not ordinary normal-message UX.

Preserved boundaries:

    KeyPackage and Welcome envelopes are onboarding artifacts.
    Relay onboarding commands are not message-inbox-dev.
    Ack must not happen merely because an artifact was fetched or written.
    Welcome ack is only eligible after successful join/reloadable state and explicit flag.

### 4.11 v0.6.10 — Integrated runtime policy + profile

v0.6.10 added integrated-runtime policy and the live-dev profile:

    integrated-runtime-dev

Current interpretation:

    integrated-runtime-dev = relay-openmls-join-dev -> dev-runtime-openmls-wrappers

Boundary:

    live umbrella only;
    not full;
    not release-snapshot;
    not package-root validation;
    not production security;
    not same-state/same-conversation proof yet.

### 4.12 v0.6.11 — Command reference generation policy

v0.6.11 documented command reference generation policy.

Preserved decisions:

    command reference generation should be registry-backed;
    generated Markdown command reference can exist before man pages;
    Unix man pages are deferred until command surface maturity improves;
    COMMAND_BOUNDARY_TABLE.v0.md remains hand-maintained;
    registry hardening must precede generated reference output.

### 4.13 v0.6.12 — Registry metadata hardening

v0.6.12 hardened registry metadata for generated-reference readiness.

Post-rewrite current commit:

    7b6446b registry: harden command reference metadata

Pre-rewrite commit for history context only:

    7b9fda7 registry: harden command reference metadata

Major effects:

    registry-lookup --missing-nonclaims reports matches: 0;
    sidecar/API/script validation-surface residuals closed;
    runner optional flag metadata added;
    runner.dev-runtime-openmls and runner.relay-openmls-join-dev lifecycle statuses added;
    no COMMAND_REFERENCE.v0.md yet;
    no man pages.

### 4.14 v0.6.13 — Generated command reference + author-history hygiene

v0.6.13 generated and committed the dev/operator command reference and cleaned public Git history.

Post-rewrite current commit:

    c5bc95c registry: add generated command reference

Pre-rewrite commit for history context only:

    a535208 registry: add generated command reference

Files changed:

    docs/README.md
    registry/README.md
    registry/COMMAND_REFERENCE.v0.md
    tools/registry/render-command-reference.py

Functional change:

    Adds generated Markdown command reference.

Non-change:

    no runtime behavior change;
    no command/profile behavior change;
    no man pages;
    no release retag as new release;
    no COMMAND_BOUNDARY_TABLE.v0.md replacement.

Generator facts:

    Input: registry/commands.v0.yaml
    Output: registry/COMMAND_REFERENCE.v0.md
    Renderer: tools/registry/render-command-reference.py
    Deterministic check: python3 tools/registry/render-command-reference.py --check
    Entry count: 79

Author-history hygiene:

    Public `main` and `carbonstack` tags were rewritten to normalized public identity.
    `carbonstack-comms` and `carbonstack-cypher` main histories were rewritten.
    `carbonstack-os` required no localdomain rewrite.
    Remote backup branches were deleted after validation.
    Local bundle backups remained.

### 4.15 v0.6.14 — Post-history-rewrite sanity + generated-reference freshness guard

v0.6.14 is the no-runtime-change sanity checkpoint after v0.6.13.

Current commit:

    77ad2d7 docs: record post-rewrite sanity checkpoint

Files changed:

    docs/209-v0.6.14-post-history-rewrite-sanity-v0.md
    docs/README.md
    registry/README.md
    tools/carbonstack-validate/README.md
    tools/carbonstack-validate/generated_command_reference_test.go

Functional effect:

    Adds a runner test that checks generated command reference freshness.

Non-change:

    no Comms runtime change;
    no Cypher runtime change;
    no CarbonStackOS change;
    no release asset rebuild;
    no release retag as a new public release;
    no man pages;
    no generated replacement for COMMAND_BOUNDARY_TABLE.v0.md.

Private safety effect:

    Created ~/CarbonStack_Umbrella_v0.6.12FREEZE.tar.gz before deleting local-only backup branches.
    Deleted local-only backup-before-author-rewrite* branches after bundle verification.

### 4.16 v0.6.15 — Same-state integrated proof plan

v0.6.15 is a docs/planning checkpoint that converts the same-state recon into an implementation target for v0.6.16.

Current commit:

    c5434b4 docs: plan same-state integrated proof

Files changed:

    docs/210-v0.6.15-same-state-integrated-proof-plan-v0.md
    docs/README.md

Functional effect:

    none at runtime.

Planning effect:

    documents that current integrated-runtime-dev is sequential composition;
    records same-state proof as likely feasible;
    chooses a temporary runner-side probe before naming/committing a profile;
    targets Level 4 same conversation proof if feasible;
    keeps full and release-snapshot untouched;
    defers vault/PQ/adversarial/man-page work.

Non-change:

    no Comms runtime change;
    no Cypher runtime change;
    no CarbonStackOS change;
    no registry entry;
    no COMMAND_REFERENCE.v0.md regeneration;
    no release asset rebuild;
    no release retag as a new public release;
    no man pages;
    no generated replacement for COMMAND_BOUNDARY_TABLE.v0.md.

Recon decision preserved:

    Current integrated-runtime-dev remains honest sequential composition.
    Same-state proof should be attempted as a temporary runner-side proof insertion in v0.6.16.
    If the temporary probe passes cleanly, then add a named dev-only profile such as same-state-integrated-dev with registry metadata, docs, generated reference update, and strict nonclaims.

---

### 4.17 v0.6.16 — Same-state integrated dev profile

v0.6.16 converted the v0.6.15 same-state plan and v0.6.16A temporary probe into a committed live-dev runner profile.

Commit:

    carbonstack 4f1fec3 feat: add same-state integrated dev profile

New profile:

    same-state-integrated-dev

Registry ID:

    runner.same-state-integrated-dev

Command:

    cd carbonstack/tools/carbonstack-validate
    go run . --profile same-state-integrated-dev --root <umbrella root> --clean-generated

Key additions:

    tools/carbonstack-validate/same_state_integrated_dev.go
    docs/211-v0.6.16-same-state-integrated-dev-profile-v0.md
    registry entry runner.same-state-integrated-dev
    regenerated registry/COMMAND_REFERENCE.v0.md
    README updates for docs and runner profile surface

Main behavior:

    The profile proves Relay onboarding and normal message send/open/ack inside one coherent live-dev temp universe.

State continuity proven:

    one runner-owned temp root;
    one temporary Cypher binary;
    one Cypher server;
    one Cypher DB;
    same Alice/Bob Comms state files;
    same Alice/Bob account IDs;
    same Alice/Bob device IDs;
    same Alice/Bob sidecar device labels;
    same Relay Space;
    same conversation label;
    Bob joins the same OpenMLS conversation through Relay onboarding;
    Alice sends a normal message through message-send-dev;
    Bob opens and acks through message-inbox-dev --ack;
    Bob inbox is empty after ack.

Proof level:

    Level 4 same conversation normal-message proof.

Important DB assertions:

    accounts = 2
    devices = 2
    relay_spaces = 1
    relay_space_members = 2
    envelopes = 3
    envelope_acks = 2
    queued KeyPackage envelopes = 1
    acknowledged Welcome envelopes = 1
    Welcome ack rows = 1

Envelope interpretation:

    The KeyPackage remains queued because KeyPackage artifact fetch/write is not ack eligibility.
    The Welcome is acknowledged only after successful join because ACK_AFTER_JOIN is enabled for this profile.
    The normal application message is acknowledged only after successful message-open and explicit --ack.

Boundary:

    The new profile is live-umbrella dev/pre-alpha validation only.
    It is not `full`.
    It is not `release-snapshot`.
    It is not package-root validation.
    It is not release-package validation.
    It is not production secure messaging.
    It is not hostile-server safety.
    It is not metadata privacy.
    It is not identity verification.
    It is not secure enrollment.
    It is not mature messenger UX.
    It is positive-path same-state proof only.

Relationship to `integrated-runtime-dev`:

    `integrated-runtime-dev` remains the sequential composition profile:
        relay-openmls-join-dev -> dev-runtime-openmls-wrappers

    `same-state-integrated-dev` is the stronger same-conversation proof profile.
    It does not silently replace or mutate `integrated-runtime-dev`.

Implementation blunder:

    The first implementation script inserted the new registry entry with malformed YAML indentation around `runner.registry-lookup`.
    `TestCommandRegistryHasNoTopLevelEntryIDs` caught the issue before commit.
    Recovery rewrote the registry row with proper two-space indentation under `entries:` and continued cleanly.

### 4.18 v0.6.17 — Same-state message failure dev profile

v0.6.17 converted the first same-state normal-message failure invariant into a committed live-dev runner profile.

Commit:

    carbonstack 03cd00a test: add same-state message failure profile

New profile:

    same-state-message-failure-dev

Registry ID:

    runner.same-state-message-failure-dev

Command:

    cd carbonstack/tools/carbonstack-validate
    go run . --profile same-state-message-failure-dev --root <umbrella root> --clean-generated

Key additions:

    tools/carbonstack-validate/same_state_message_failure_dev.go
    docs/212-v0.6.17-same-state-message-failure-dev-profile-v0.md
    registry entry runner.same-state-message-failure-dev
    regenerated registry/COMMAND_REFERENCE.v0.md
    README updates for docs and runner profile surface

Main behavior:

    The profile proves wrong-conversation normal-message open failure does not ack or drain the inbox after a same-state Relay join.

State continuity reused:

    one runner-owned temp root;
    one temporary Cypher binary;
    one Cypher server;
    one Cypher DB;
    same Alice/Bob Comms state files;
    same Alice/Bob account IDs;
    same Alice/Bob device IDs;
    same Alice/Bob sidecar device labels;
    same Relay Space;
    same correct conversation label;
    one deliberately wrong conversation label for the negative open attempt.

Proof sequence:

    Relay KeyPackage -> add-member -> Welcome -> join;
    Alice sends normal message with message-send-dev;
    Bob tries message-inbox-dev --ack with wrong conversation label;
    wrong-conversation open reports message-open failure;
    wrong-conversation open reports acked: false;
    wrong-conversation open leaves ack count unchanged;
    wrong-conversation open leaves Bob inbox count unchanged;
    Bob then opens the same message with the correct conversation label;
    correct open reports message opened and acked: true;
    Bob inbox is empty after the correct open/ack.

Expected wrong-conversation evidence:

    message open failed
    stage: message_open
    acked: false
    opened_envelopes: 0
    open_failures: 1
    ack_failures: 0

Expected counter evidence:

    acks_after_bad_open == acks_after_join
    bob_inbox_count_after_bad_open == bob_inbox_count_before_bad_open
    acks_after_correct_open == acks_after_join + 1
    bob_inbox_count_after_correct_open == 0

Boundary:

    The new profile is live-umbrella dev/pre-alpha validation only.
    It is not `full`.
    It is not `release-snapshot`.
    It is not package-root validation.
    It is not release-package validation.
    It is not an adversarial relay harness.
    It is not hostile-server safety.
    It is not metadata privacy.
    It is not production secure messaging.
    It is not production E2EE.
    It is not identity verification.
    It is not mature messenger UX.
    It currently covers wrong-conversation message-open failure only.

Relationship to other profiles:

    `same-state-integrated-dev` remains the positive-path same-conversation proof profile.
    `same-state-message-failure-dev` is the first failure-path hardening companion profile.
    `integrated-runtime-dev` remains the sequential composition profile.

Implementation blunder:

    The first implementation generated an overlong message label:
        same-state-failure-message-same-state-failure-<id>

    The sidecar rejected it:
        invalid_message_label: message label is too long

    Recovery split the long run ID from the message label and shortened the label to:
        fail-msg-<id>

    The recovered profile passed and was committed.

### 4.19 v0.6.18 — Same-state message unsupported dev profile

v0.6.18 converted the second same-state normal-message failure invariant into a committed live-dev runner profile.

Commit:

    carbonstack 1f8086b test: add same-state unsupported message profile

New profile:

    same-state-message-unsupported-dev

Registry ID:

    runner.same-state-message-unsupported-dev

Command:

    cd carbonstack/tools/carbonstack-validate
    go run . --profile same-state-message-unsupported-dev --root <umbrella root> --clean-generated

Key additions:

    tools/carbonstack-validate/same_state_message_unsupported_dev.go
    docs/213-v0.6.18-same-state-message-unsupported-dev-profile-v0.md
    registry entry runner.same-state-message-unsupported-dev
    regenerated registry/COMMAND_REFERENCE.v0.md
    README updates for docs and runner profile surface

Main behavior:

    The profile proves unsupported normal application-message content type does not ack or drain the inbox after a same-state Relay join.

State continuity reused:

    one runner-owned temp root;
    one temporary Cypher binary;
    one Cypher server;
    one Cypher DB;
    same Alice/Bob Comms state files;
    same Alice/Bob account IDs;
    same Alice/Bob device IDs;
    same Alice/Bob sidecar device labels;
    same Relay Space;
    same correct conversation label;
    one deliberately unsupported normal-message content type.

Proof sequence:

    Relay KeyPackage -> add-member -> Welcome -> join;
    Alice sends normal message with message-send-dev;
    runner locates the normal application-message envelope for Bob;
    runner mutates only that envelope's content_type from carbonstack.mls.application-message.v0 to carbonstack.test.unsupported-normal-message.v0;
    Bob runs message-inbox-dev --ack with the correct conversation and message label;
    unsupported envelope is skipped;
    unsupported attempt does not ack;
    unsupported attempt does not drain Bob's inbox;
    runner restores original content type;
    Bob opens the same message with the correct conversation label;
    correct open reports message opened and acked: true;
    Bob inbox is empty after the correct open/ack.

Expected unsupported evidence:

    message skipped
    reason: unsupported_envelope
    opened_envelopes: 0
    unsupported_envelopes: 1
    open_failures: 0
    ack_failures: 0

Expected counter evidence:

    acks_after_unsupported_open == acks_before_unsupported_open
    bob_inbox_count_after_unsupported_open == bob_inbox_count_before_unsupported_open
    acks_after_correct_open == acks_before_unsupported_open + 1
    bob_inbox_count_after_correct_open == 0

Boundary:

    The new profile is live-umbrella dev/pre-alpha validation only.
    It is not `full`.
    It is not `release-snapshot`.
    It is not package-root validation.
    It is not release-package validation.
    It is not an adversarial relay harness.
    It is not hostile-server safety.
    It is not metadata privacy.
    It is not production secure messaging.
    It is not production E2EE.
    It is not identity verification.
    It is not mature messenger UX.
    It currently covers unsupported normal application-message content_type only.

Relationship to other profiles:

    `same-state-integrated-dev` remains the positive-path same-conversation proof profile.
    `same-state-message-failure-dev` remains the wrong-conversation no-ack/no-drain profile.
    `same-state-message-unsupported-dev` is the unsupported normal-message no-ack/no-drain profile.
    `integrated-runtime-dev` remains the sequential composition profile.

Implementation blunder state:

    No source recovery was required in v0.6.18.
    The only operational repeat was the known scripted push authentication failure.
    Manual push succeeded and origin/main now matches 1f8086b.

### 4.20 v0.6.19 — Same-state recipient/device failure dev profile

v0.6.19 added the third same-state failure-path hardening profile:

    same-state-message-recipient-failure-dev

Registry ID:

    runner.same-state-message-recipient-failure-dev

Command:

    go run . --profile same-state-message-recipient-failure-dev --root <umbrella root> --clean-generated

Purpose:

    prove wrong recipient/device/sidecar attempts do not falsely open, ack, or drain Bob's inbox after same-state Relay join.

The profile covers:

    Case A: Alice state + Alice sidecar attempts message-inbox-dev --ack.
    Case B: Bob state + Alice sidecar attempts message-inbox-dev --ack.
    Case C: Bob state + missing sidecar attempts message-inbox-dev --ack.
    Recovery: Bob state + Bob sidecar opens and acks the same queued message.

Shared invariant:

    no false message-open success;
    no ack;
    no Bob inbox drain;
    correct Bob open/ack still succeeds.

v0.6.19 did not change:

    same-state-integrated-dev;
    same-state-message-failure-dev;
    same-state-message-unsupported-dev;
    integrated-runtime-dev;
    full;
    release-snapshot;
    Comms/Cypher/OS source;
    release assets;
    public release state.

Commit:

    197ea9f test: add same-state recipient failure profile

Generated command reference:

    entries=83

### 4.21 v0.6.20 — Welcome join partial-state safety patch

v0.6.20 started as failed Welcome/join no-ack hardening and became a correctness/state-safety patch.

v0.6.20A found:

    corrupt Welcome join with --ack-after-join
    -> failed join
    -> no Welcome ack
    -> no Bob Relay inbox drain
    -> restored valid Welcome blocked by conversation_already_exists

v0.6.20B deep recon found:

    failed corrupt Welcome join wrote durable final Bob conversation state;
    Bob's final conversation directory poisoned restored join;
    manual quarantine of that final directory allowed restored valid Welcome join and ack;
    therefore the fix target was sidecar conversation-join atomicity, not Cypher ack ordering or OpenMLS version docs.

v0.6.20C implementation patched the sidecar:

    join_dev_conversation stages joined state under a staging conversation directory;
    provider storage and summaries are written under staging;
    staged provider storage is reload-checked;
    staging is promoted to final conversation state only after success;
    failed materialization removes staging;
    corrupt Welcome no longer leaves final Bob conversation state in regression coverage.

Committed changes:

    carbonstack-comms:
        b42aa09 fix: make Welcome join state writes atomic

    carbonstack:
        e1f6b8c docs: record Welcome join partial-state fix

Files added/changed:

    carbonstack-comms/internal/protocol/mls/openmls-sidecar/src/state.rs
    carbonstack-comms/internal/protocol/openmls_sidecar_conversation_test.go
    carbonstack/docs/215-v0.6.20-welcome-join-partial-state-safety-v0.md
    carbonstack/docs/README.md

Important v0.6.20 nonclaim:

    No same-state-welcome-join-failure-dev profile exists yet.
    v0.6.20 fixed and documented the state-safety defect first.
    A later profile should prove corrupt Welcome no-ack/no-drain/no-final-state-poisoning in live same-state runner context.


### 4.22 v0.6.21 — Same-state Welcome join failure dev profile

v0.6.21 completed the profile checkpoint that v0.6.20 intentionally deferred.

v0.6.21A recon proved the fixed Welcome join failure behavior in a live Relay same-state context:

    corrupt Welcome join
    -> openmls-relay-join-dev --ack-after-join fails with welcome_invalid
    -> no join/ack success markers
    -> ack count unchanged
    -> Welcome ack rows unchanged
    -> Bob Relay inbox remains queued
    -> no final Bob conversation state
    -> no staging Bob conversation state
    -> Bob conversation load-check fails as expected
    -> restore original valid Welcome
    -> Bob joins with same conversation label
    -> Welcome ack occurs only after successful join
    -> Bob Relay inbox drains
    -> Bob conversation reload-check passes

Key v0.6.21A evidence:

    bad_join_rc: 1
    good_join_rc: 0
    bob_load_after_bad_rc: 1
    acks_before_bad_join: 0
    acks_after_bad_join: 0
    acks_after_good_join: 1
    welcome_acks_before_bad_join: 0
    welcome_acks_after_bad_join: 0
    welcome_acks_after_good_join: 1
    bob_relay_inbox_before_bad_join: 1
    bob_relay_inbox_after_bad_join: 1
    bob_relay_inbox_after_good_join: 0
    final_state_after_bad_join: absent
    staging_state_after_bad_join: absent
    final_state_after_restored_join: present
    restored_join_recovered_without_manual_cleanup: yes

v0.6.21B attempted to add the profile and initially failed because the profile asserted that `openmls-relay-add-member-dev` output must contain literal `content_type`. That assertion was over-strict because add-member output for the Welcome path does not currently print `content_type`; it prints Welcome-specific evidence instead.

Recovery removed only the invalid `content_type` assertion and kept the actual Welcome envelope/payload/delivery/ack evidence. The profile then passed.

v0.6.21 added:

    same-state-welcome-join-failure-dev
    runner.same-state-welcome-join-failure-dev
    docs/216-v0.6.21-same-state-welcome-join-failure-dev-profile-v0.md
    tools/carbonstack-validate/same_state_welcome_join_failure_dev.go

v0.6.21 updated:

    tools/carbonstack-validate/main.go
    tools/carbonstack-validate/README.md
    docs/README.md
    registry/commands.v0.yaml
    registry/COMMAND_REFERENCE.v0.md

Generated command reference moved to:

    entries=84

Commit:

    3876fc8 test: add same-state Welcome join failure profile

Push note:

    scripted push failed with rc=128;
    manual push succeeded: e1f6b8c..3876fc8 main -> main.

Boundary:

    live-umbrella dev profile only;
    not full;
    not release-snapshot;
    not package-root validation;
    not release-package validation;
    not adversarial harness;
    not hostile-server safety;
    not identity verification;
    not production secure messaging.

### 4.23 v0.6.22 — Sidecar state authority classification checkpoint

v0.6.22 classified stale sidecar/conversation/provider state behavior without adding a runtime patch or committed runner profile.

Commit:

    carbonstack 8dd7db5 docs: classify sidecar state authority

Files changed:

    docs/217-v0.6.22-sidecar-state-authority-classification-v0.md
    docs/README.md

Functional effect:

    none at runtime.

Registry/generated-reference effect:

    no registry entry added;
    no runner profile added;
    COMMAND_REFERENCE.v0.md remains at 84 entries;
    generated reference check remains current.

Recon sequence:

    v0.6.22A stale-state recon attempted to mutate sidecar-generated state around an already joined conversation.
    Initial attempt failed before stale-state semantics because generated message labels were too long.
    A short-label recovery recon reached the stale-state cases and found a split:
        summary files and device-level provider state did not block message-open/ack;
        conversation provider storage and missing conversation directory blocked open/ack/drain.
    v0.6.22B then performed state-authority classification recon over 12 cases.
    v0.6.22C documented the classification in the CarbonStack docs surface.

v0.6.22B classification result:

    cases_total: 12
    classification_count critical_state_blocks_open_ack_drain: 4
    classification_count device_provider_not_runtime_message_open_authority: 2
    classification_count join_summary_not_runtime_authority: 3
    classification_count load_check_stricter_than_message_open: 1
    classification_count summary_not_runtime_authority: 2
    authority_guess_count identity_or_bootstrap_substrate_not_conversation_runtime: 2
    authority_guess_count metadata_or_evidence_not_runtime_authority: 6
    authority_guess_count runtime_authority: 4
    cases_with_target_rewrite_or_regeneration: 0
    critical_runtime_bypass_cases: 0

Classified runtime authority:

    conversation-level provider-storage.json
    final conversation directory

Classified metadata/evidence or non-runtime existing-message-open substrate:

    conversation-summary.json
    join-summary.json
    device-level provider-storage.json

Important semantic mismatch:

    deleting conversation-summary.json makes conversation-load-check-dev fail;
    message-inbox-dev --ack can still open and ack a valid queued message when conversation provider-storage.json remains valid.

Current interpretation:

    conversation-summary.json is not cryptographic/runtime authority for existing message-open;
    load-check is currently stricter than message-open because it checks metadata presence as well as provider reloadability;
    future cleanup should split or clarify provider reloadability vs metadata state checks.

Nonclaims:

    not a public release;
    not a runtime patch;
    not a stale-state runner profile;
    not full;
    not release-snapshot;
    not package-root validation;
    not adversarial harness;
    not hostile-server safety;
    not vault/key-storage safety;
    not production secure messaging;
    not stale-state model closure across every possible file/failure mode.

Push note:

    scripted push failed again with the known auth quirk:
        RESULT push rc=128

    manual push succeeded:
        3876fc8..8dd7db5 main -> main


### 4.24 v0.6.23 — Same-state malformed normal-message payload profile

v0.6.23 promoted malformed normal application-message payload behavior into a live same-state runner profile.

It added:

    same-state-message-malformed-payload-dev
    runner.same-state-message-malformed-payload-dev
    docs/218-v0.6.23-same-state-message-malformed-payload-dev-profile-v0.md

It proved six payload/ciphertext mutation families fail without false open, ack, inbox drain, provider mutation, or envelope rewrite, and that restoring the original payload recovers open/ack.

Commit:

    943e0e4 test: add same-state malformed payload profile

### 4.25 v0.6.24 — Same-state normal application-message replay classification profile

v0.6.24 promoted normal application-message duplicate/replay classification into a live same-state runner profile and documented the load-check summary metadata mismatch.

It added:

    same-state-message-replay-classification-dev
    runner.same-state-message-replay-classification-dev
    docs/219-v0.6.24-loadcheck-replay-classification-v0.md

It classified:

    duplicate same envelope_id insert -> storage rejected;
    same envelope after ack -> delivery-state suppressed;
    same envelope manually requeued after ack -> detected/rejected no-open/no-ack/no-drain;
    duplicate same ciphertext under new envelope ID before original ack -> detected/rejected no-open/no-ack/no-drain;
    duplicate same ciphertext under new envelope ID after original ack -> detected/rejected no-open/no-ack/no-drain.

Commit:

    1c421e2 test: add same-state replay classification profile

### 4.26 v0.6.25 — Normal application-message envelope metadata classification docs checkpoint

v0.6.25 classified remaining normal application-message typed/envelope metadata behavior without adding a runner profile.

It added:

    docs/220-v0.6.25-normal-message-envelope-metadata-classification-v0.md

It classified:

    payload_sha256 mismatch -> detected/rejected no-ack/no-drain, recovery OK;
    payload_size_bytes mismatch -> detected/rejected no-ack/no-drain, recovery OK;
    invalid protocol_version -> unsupported/skipped no-ack/no-drain, recovery OK;
    wrong recipient_device_id -> routing suppression no-open/no-ack, recovery OK;
    invalid delivery_state -> delivery-state suppression no-open/no-ack, recovery OK;
    wrong/empty sender_device_id -> opens and acks as unverified relay metadata.

The key model update is:

    from_device in current message-inbox-dev output is relay envelope sender metadata, not verified identity.

Commit:

    b7f89ea docs: classify normal message envelope metadata



### 4.27 v0.6.26 — Roadmap refresh and output semantics cleanup checkpoint

v0.6.26 completed the long-form adversarial Q/A preflight and converted it into a refreshed EVERGREEN roadmap plus two output-semantics cleanup commits.

Roadmap artifact:

    CarbonStack_Long_Term_Roadmap_v0.6.26_EVERGREEN.pdf

Main decisions:

    remaining v0.6.x owns roadmap refresh, load-check/sender metadata cleanup, vault/security modeling, no-silent/backup/recovery rules, bounded vault substrate/stub decision, PQ/hybrid placement model, adversarial harness contract/evidence matrix design, and v0.7.0 release-boundary/package-root rehearsal planning;
    v0.7.x owns deployability plus state-security integration, Addressable Relay Space mechanics, rudimentary CLI dev app for joining/using Relay Spaces, full-validate-release/full-runtime-dev split, dev CLI/client lifecycle, local state/vault substrate integration, preferred/generic deployment modeling, and isolated PQ pilots only after v0.6 model completion;
    v0.8.x owns PQ/deployability maturation, migration/compatibility tests, optimization/code-health hardening, package-root runtime validation candidates, operator docs, and release-candidate hardening;
    v0.9.x owns full self-pen-test/adversarial validation across real operational surfaces.

v0.6.26B scout:

    confirmed sender metadata output needed explicit unverified labeling;
    confirmed load-check provider-vs-summary output needed cleanup;
    preserved baseline validation and did not patch code.

v0.6.26C sender metadata output warning cleanup:

    carbonstack-comms  61b2e51 fix: label message sender metadata as unverified
    carbonstack        499ab6a docs: record sender metadata output warning

New normal message inbox output:

    from_device: <legacy sender_device_id>
    from_device_unverified: <sender_device_id>
    sender_identity_verified: false
    warning: from_device is relay envelope metadata, not verified identity

v0.6.26D load-check provider/summary output cleanup:

    carbonstack-comms  2d13c45 fix: clarify load-check provider metadata output
    carbonstack        468432b docs: record load-check provider summary output

New successful load-check output:

    provider_reloadable: true
    summary_metadata_present: true
    provider_storage_present: true
    summary_metadata_warning: none

New missing-summary output:

    status: metadata_missing
    provider_reloadable: not_evaluated_by_load_check
    summary_metadata_present: false
    provider_storage_present: true/false
    summary_metadata_warning: conversation-summary metadata is missing; conversation-load-check-dev is stricter than message-open and cannot confirm reloadability without summary metadata

Boundaries:

    no registry change;
    COMMAND_REFERENCE.v0.md remains entries=86;
    no full/release-snapshot change;
    no message-open behavior change;
    no ack/drain behavior change;
    no provider-state behavior change;
    no identity verification implementation;
    no vault/PQ/adversarial harness implementation;
    no public release.

## 5. v0.6.14-v0.6.26 Validation Ledger

### 5.1 Hard sanity recon

Repo/rewrite sanity:

    carbonstack        main aligned with origin/main at c5bc95c during recon
    carbonstack-comms  main aligned with origin/main at 53779d5
    carbonstack-cypher main aligned with origin/main at d18a564
    carbonstack-os     main aligned with origin/main at 1bbbe52

Public main/tag identity hygiene:

    carbonstack        localdomain_count=0
    carbonstack-comms  localdomain_count=0
    carbonstack-cypher localdomain_count=0
    carbonstack-os     localdomain_count=0

Remote backup refs:

    none found

Release/tag sanity:

    local v0.6.0 tag  = 70d4318
    remote v0.6.0 tag = 70d4318
    current main      = c5bc95c during recon, later 77ad2d7 after v0.6.14 commit

Generated-reference sanity:

    RESULT generated reference deterministic check rc=0
    RESULT generated reference entry count check rc=0
    RESULT generated reference boundary header check rc=0
    RESULT generated reference script grouping check rc=0
    RESULT generated reference internal/api sections check rc=0
    RESULT registry/reference exact ID consistency rc=0
    RESULT no man pages guard rc=0

Registry/reference counts:

    registry_ids=79
    reference_ids=79
    unique_registry_ids=79
    unique_reference_ids=79
    missing_from_reference=[]
    extra_in_reference=[]
    duplicate_registry_ids=[]
    duplicate_reference_ids=[]

Release asset sanity:

    RESULT curl Gitea release API v0.6.0 rc=0
    asset_count=8
    RESULT download release assets rc=0
    asset checksum file verified all attached assets
    downloaded public package verify-checksums passed
    downloaded public package full --clean-generated passed

Runner validation:

    registry-lookup --missing-nonclaims passed
    registry lookup spot checks passed
    runner go test ./... passed
    runner doctor passed
    integrated-runtime-dev --clean-generated passed

### 5.2 v0.6.14 patch validation

Patch validation passed:

    RESULT carbonstack diff --check rc=0
    RESULT generated reference deterministic check rc=0
    RESULT runner go test ./... rc=0
    RESULT registry missing nonclaims scan rc=0
    RESULT runner doctor rc=0
    RESULT integrated-runtime-dev sanity rc=0
    RESULT no man pages guard rc=0
    RESULT cached diff --check rc=0

Commit:

    77ad2d7 docs: record post-rewrite sanity checkpoint

Commit diff summary:

    5 files changed, 242 insertions(+)
    create mode 100644 docs/209-v0.6.14-post-history-rewrite-sanity-v0.md
    create mode 100644 tools/carbonstack-validate/generated_command_reference_test.go

Post-commit validation passed:

    RESULT generated reference deterministic check post-commit rc=0
    RESULT runner go test ./... post-commit rc=0
    RESULT registry missing nonclaims scan post-commit rc=0
    RESULT runner doctor post-commit rc=0

Push:

    Scripted push failed authentication.
    Manual retry from carbonstack succeeded:
        c5bc95c..77ad2d7 main -> main

Final manual snapshot:

    carbonstack        77ad2d7 (HEAD -> main, origin/main, origin/HEAD)
    carbonstack-comms  53779d5 (HEAD -> main, origin/main, origin/HEAD)
    carbonstack-cypher d18a564 (HEAD -> main, origin/main, origin/HEAD)
    carbonstack-os     1bbbe52 (HEAD -> main, origin/main, origin/HEAD)

### 5.3 v0.6.15 same-state recon validation

Recon baseline validation passed:

    RESULT generated reference deterministic check rc=0
    RESULT runner go test ./... rc=0
    RESULT registry missing nonclaims scan rc=0
    RESULT runner doctor rc=0

Recon confirmed relevant registry entries exist and preserve strict nonclaims for:

    runner.integrated-runtime-dev
    runner.relay-openmls-join-dev
    runner.dev-runtime-openmls-wrappers
    runner.dev-runtime-openmls
    comms.message-send-dev
    comms.message-inbox-dev
    Relay KeyPackage/Welcome/add-member/join commands
    sidecar conversation and message primitives
    Cypher envelope submit/fetch/ack APIs

Important recon result:

    integrated-runtime-dev remains sequential composition.
    relay-openmls-join-dev already proves a strong Relay onboarding state story.
    dev-runtime-openmls-wrappers remains independent wrapper smoke and currently resets/owns its own state.
    same-state proof likely needs runner-side wiring rather than direct reuse of wrapper smoke script.

### 5.4 v0.6.15 docs/planning patch validation

Patch validation passed:

    RESULT carbonstack diff --check rc=0
    RESULT generated reference deterministic check rc=0
    RESULT runner go test ./... rc=0
    RESULT registry missing nonclaims scan rc=0
    RESULT runner doctor rc=0
    RESULT no runtime/registry/generated-reference drift guard rc=0
    RESULT cached diff --check rc=0

Commit:

    c5434b4 docs: plan same-state integrated proof

Commit diff summary:

    2 files changed, 369 insertions(+)
    create mode 100644 docs/210-v0.6.15-same-state-integrated-proof-plan-v0.md

Post-commit validation passed:

    RESULT generated reference deterministic check post-commit rc=0
    RESULT runner go test ./... post-commit rc=0
    RESULT registry missing nonclaims scan post-commit rc=0
    RESULT runner doctor post-commit rc=0

Push:

    Scripted push failed authentication.
    Manual retry from carbonstack succeeded:
        77ad2d7..c5434b4 main -> main

Final manual snapshot:

    carbonstack        c5434b4 (HEAD -> main, origin/main, origin/HEAD)
    carbonstack-comms  53779d5 (HEAD -> main, origin/main, origin/HEAD)
    carbonstack-cypher d18a564 (HEAD -> main, origin/main, origin/HEAD)
    carbonstack-os     1bbbe52 (HEAD -> main, origin/main, origin/HEAD)

---

### 5.5 v0.6.16A temporary same-state proof probe

The temporary v0.6.16A probe passed before the named profile was committed.

Probe target:

    Relay onboarding plus normal message send/open/ack inside one coherent state story.

Probe proved:

    one runner-owned temp workdir;
    one Cypher server;
    one Cypher DB;
    same Alice/Bob Comms state files;
    same Alice/Bob Cypher device IDs;
    same Alice/Bob sidecar labels;
    same conversation label;
    Relay KeyPackage -> add-member -> Welcome -> join;
    message-send-dev -> message-inbox-dev --ack using joined state;
    plaintext matched expected probe payload;
    Bob inbox count after normal ack was 0.

Probe DB counters:

    accounts|2
    devices|2
    relay_spaces|1
    relay_space_members|2
    envelopes|3
    envelope_acks|2

Probe interpretation:

    Same-state proof was no longer merely theoretical.
    The safe next rung became a committed runner profile rather than changing `integrated-runtime-dev` in place.

### 5.6 v0.6.16B committed profile implementation and recovery

Initial implementation attempt:

    added tools/carbonstack-validate/same_state_integrated_dev.go;
    updated main.go profile help/switch/error text;
    added registry entry for runner.same-state-integrated-dev;
    added docs/211-v0.6.16-same-state-integrated-dev-profile-v0.md;
    updated docs/README.md;
    updated tools/carbonstack-validate/README.md;
    regenerated registry/COMMAND_REFERENCE.v0.md.

Initial failure:

    runner go test ./... failed in TestCommandRegistryHasNoTopLevelEntryIDs.

Failure text:

    registry entry at line 1747 is malformed top-level YAML;
    expected two-space indentation under entries:
        "- id: runner.registry-lookup"

Cause:

    The registry insertion matched `- id: runner.registry-lookup` without preserving the leading two-space indentation.
    The new row existed, but the YAML shape was invalid.

Recovery:

    removed the malformed same-state row;
    normalized `runner.registry-lookup` back under `entries:`;
    reinserted `runner.same-state-integrated-dev` with correct two-space entry indentation;
    regenerated `COMMAND_REFERENCE.v0.md`.

Post-recovery registry facts:

    `runner.same-state-integrated-dev` appears under `entries:`.
    `registry/COMMAND_REFERENCE.v0.md` regenerated to 80 entries.
    `registry lookup --registry-id runner.same-state-integrated-dev` passed.
    `registry lookup --missing-nonclaims` returned matches 0.

Validation after recovery passed:

    RESULT carbonstack diff --check rc=0
    RESULT runner go test ./... rc=0
    RESULT registry missing nonclaims scan rc=0
    RESULT registry lookup same-state profile rc=0
    RESULT runner doctor rc=0
    RESULT same-state-integrated-dev --clean-generated rc=0
    RESULT generated reference deterministic check post-profile rc=0
    RESULT runner go test ./... post-profile rc=0
    RESULT registry missing nonclaims scan post-profile rc=0
    RESULT runner doctor post-profile rc=0
    RESULT full/release-snapshot/integrated-runtime-dev guard rc=0
    RESULT cached diff --check rc=0

Commit:

    4f1fec3 feat: add same-state integrated dev profile

Commit diff summary:

    7 files changed, 663 insertions(+), 4 deletions(-)
    create mode 100644 docs/211-v0.6.16-same-state-integrated-dev-profile-v0.md
    create mode 100644 tools/carbonstack-validate/same_state_integrated_dev.go

Post-commit validation passed:

    RESULT generated reference deterministic check post-commit rc=0
    RESULT runner go test ./... post-commit rc=0
    RESULT registry lookup same-state post-commit rc=0
    RESULT same-state-integrated-dev post-commit rc=0

Push:

    Scripted push failed authentication earlier during the implementation script.
    Manual retry from carbonstack succeeded:
        c5434b4..4f1fec3 main -> main

Final manual snapshot:

    carbonstack        4f1fec3 (HEAD -> main, origin/main, origin/HEAD)
    carbonstack-comms  53779d5 (HEAD -> main, origin/main, origin/HEAD)
    carbonstack-cypher d18a564 (HEAD -> main, origin/main, origin/HEAD)
    carbonstack-os     1bbbe52 (HEAD -> main, origin/main, origin/HEAD)

Final status:

    carbonstack main and origin/main aligned at 4f1fec3.
    carbonstack-comms, carbonstack-cypher, and carbonstack-os remained unchanged.
    `cypher.db` may remain as an ignored local dev artifact depending on local checkout state.

### 5.7 v0.6.17A same-state normal-message failure recon

Recon baseline validation passed:

    RESULT generated reference deterministic check rc=0
    entries=80
    RESULT runner go test ./... rc=0
    RESULT registry missing nonclaims scan rc=0
    RESULT runner doctor rc=0
    RESULT same-state-integrated-dev baseline rc=0

Registry lookup passed for:

    runner.same-state-integrated-dev
    runner.integrated-runtime-dev
    runner.relay-openmls-join-dev
    comms.message-send-dev
    comms.message-inbox-dev
    comms.openmls-send-dev
    comms.openmls-inbox-dev
    sidecar.message-protect
    sidecar.message-open
    cypher.api.envelopes-submit
    cypher.api.device-envelopes
    cypher.api.envelope-ack

Exploratory negative probe result:

    failure_case: wrong conversation label on message-inbox-dev --ack
    wrong_conversation_rc: 0
    acks_after_bad_open == acks_after_join
    bob_inbox_count_after_bad_open == bob_inbox_count_before_bad_open
    correct conversation open after failed attempt succeeded
    Bob inbox empty after correct open/ack

Important interpretation:

    Wrong-conversation open did not hard-fail the process, but it also did not ack and did not drain the inbox.
    The correct patch target was therefore a validation profile locking the no-ack/no-drain invariant, not runtime surgery to force a nonzero rc.

### 5.8 v0.6.17B committed failure profile implementation and recovery

Initial implementation attempt:

    added tools/carbonstack-validate/same_state_message_failure_dev.go;
    updated main.go profile help/switch/error text;
    added registry entry for runner.same-state-message-failure-dev;
    added docs/212-v0.6.17-same-state-message-failure-dev-profile-v0.md;
    updated docs/README.md;
    updated tools/carbonstack-validate/README.md;
    regenerated registry/COMMAND_REFERENCE.v0.md to 81 entries.

Initial validation before profile run passed:

    RESULT carbonstack diff --check rc=0
    RESULT runner go test ./... rc=0
    RESULT registry missing nonclaims scan rc=0
    RESULT registry lookup same-state message failure profile rc=0
    RESULT runner doctor rc=0

Initial failure:

    same-state-message-failure-dev failed during message-send-dev setup.

Failure text:

    error: OpenMLS sidecar message-protect failed: invalid_message_label: message label is too long

Cause:

    The generated message label duplicated the long run prefix:
        same-state-failure-message-same-state-failure-<id>

Recovery:

    introduced a separate uniqueID;
    kept runID as:
        same-state-failure-<id>
    shortened messageLabel to:
        fail-msg-<id>

Recovered validation passed:

    RESULT carbonstack diff --check rc=0
    RESULT runner go test ./... rc=0
    RESULT registry missing nonclaims scan rc=0
    RESULT registry lookup same-state message failure profile rc=0
    RESULT runner doctor rc=0
    RESULT same-state-message-failure-dev --clean-generated rc=0
    RESULT generated reference deterministic check post-profile rc=0
    RESULT runner go test ./... post-profile rc=0
    RESULT same-state-integrated-dev still passes post-profile rc=0
    RESULT registry missing nonclaims scan post-profile rc=0
    RESULT runner doctor post-profile rc=0
    RESULT protected profile guard rc=0
    RESULT cached diff --check rc=0

Commit:

    03cd00a test: add same-state message failure profile

Commit diff summary:

    7 files changed, 627 insertions(+), 4 deletions(-)
    create mode 100644 docs/212-v0.6.17-same-state-message-failure-dev-profile-v0.md
    create mode 100644 tools/carbonstack-validate/same_state_message_failure_dev.go

Post-commit validation passed:

    RESULT generated reference deterministic check post-commit rc=0
    RESULT runner go test ./... post-commit rc=0
    RESULT registry lookup same-state message failure post-commit rc=0
    RESULT same-state-message-failure-dev post-commit rc=0

Push:

    Scripted push failed authentication.
    Manual retry from carbonstack succeeded:
        4f1fec3..03cd00a main -> main

Final manual snapshot:

    carbonstack        03cd00a (HEAD -> main, origin/main, origin/HEAD)
    carbonstack-comms  53779d5 (HEAD -> main, origin/main, origin/HEAD)
    carbonstack-cypher d18a564 (HEAD -> main, origin/main, origin/HEAD)
    carbonstack-os     1bbbe52 (HEAD -> main, origin/main, origin/HEAD)

Final status:

    carbonstack main and origin/main aligned at 03cd00a.
    carbonstack-comms, carbonstack-cypher, and carbonstack-os remained unchanged.
    `cypher.db` may remain as an ignored local dev artifact depending on local checkout state.

### 5.9 v0.6.18A/B unsupported normal-message profile implementation

v0.6.18A recon baseline validation passed:

    RESULT generated reference deterministic check rc=0
    entries=81
    RESULT runner go test ./... rc=0
    RESULT registry missing nonclaims scan rc=0
    RESULT runner doctor rc=0
    RESULT same-state-integrated-dev baseline rc=0
    RESULT same-state-message-failure-dev baseline rc=0

v0.6.18A exploratory unsupported-content-type probe result:

    failure_case: unsupported normal application-message content_type on message-inbox-dev --ack
    original_content_type: carbonstack.mls.application-message.v0
    unsupported_content_type: carbonstack.test.unsupported-normal-message.v0
    message skipped
    reason: unsupported_envelope
    opened_envelopes: 0
    unsupported_envelopes: 1
    open_failures: 0
    ack_failures: 0
    acks_before_unsupported_open: 1
    acks_after_unsupported_open: 1
    acks_after_correct_open: 2
    bob_inbox_count_before_unsupported_open: 1
    bob_inbox_count_after_unsupported_open: 1
    bob_inbox_count_after_correct_open: 0

Important interpretation:

    Unsupported normal-message content_type did not ack and did not drain Bob inbox.
    Restoring the original content_type allowed the same valid message to open and ack.
    The correct patch target was a narrow validation profile, not runtime surgery.

v0.6.18B implementation added:

    tools/carbonstack-validate/same_state_message_unsupported_dev.go
    docs/213-v0.6.18-same-state-message-unsupported-dev-profile-v0.md
    registry entry runner.same-state-message-unsupported-dev
    main.go profile registration
    runner README update
    docs README update
    regenerated registry/COMMAND_REFERENCE.v0.md

Pre-profile validation passed:

    RESULT carbonstack diff --check rc=0
    RESULT runner go test ./... rc=0
    RESULT registry missing nonclaims scan rc=0
    RESULT registry lookup same-state unsupported profile rc=0
    RESULT runner doctor rc=0

Profile validation passed:

    RESULT same-state-message-unsupported-dev --clean-generated rc=0

Post-profile validation passed:

    RESULT generated reference deterministic check post-profile rc=0
    RESULT runner go test ./... post-profile rc=0
    RESULT same-state-integrated-dev still passes post-profile rc=0
    RESULT same-state-message-failure-dev still passes post-profile rc=0
    RESULT registry missing nonclaims scan post-profile rc=0
    RESULT runner doctor post-profile rc=0
    RESULT protected profile guard rc=0

Protected profile guard checked that these were not mutated:

    tools/carbonstack-validate/integrated_runtime_dev.go
    tools/carbonstack-validate/release_snapshot.go
    tools/carbonstack-validate/same_state_integrated_dev.go
    tools/carbonstack-validate/same_state_message_failure_dev.go

Commit:

    1f8086b test: add same-state unsupported message profile

Commit diff summary:

    7 files changed, 689 insertions(+), 4 deletions(-)
    create mode 100644 docs/213-v0.6.18-same-state-message-unsupported-dev-profile-v0.md
    create mode 100644 tools/carbonstack-validate/same_state_message_unsupported_dev.go

Post-commit validation passed:

    RESULT generated reference deterministic check post-commit rc=0
    RESULT runner go test ./... post-commit rc=0
    RESULT registry lookup same-state unsupported post-commit rc=0
    RESULT same-state-message-unsupported-dev post-commit rc=0

Generated reference:

    entries=82

Push:

    Scripted push failed authentication:
        RESULT push rc=128

    Manual retry from carbonstack succeeded:
        03cd00a..1f8086b main -> main

Final manual snapshot:

    carbonstack        1f8086b (HEAD -> main, origin/main, origin/HEAD)
    carbonstack-comms  53779d5 (HEAD -> main, origin/main, origin/HEAD)
    carbonstack-cypher d18a564 (HEAD -> main, origin/main, origin/HEAD)
    carbonstack-os     1bbbe52 (HEAD -> main, origin/main, origin/HEAD)

### 5.10 v0.6.19A/B recipient/device failure profile implementation

v0.6.19A recon baseline validation passed:

    RESULT generated reference deterministic check rc=0
    entries=82
    RESULT runner go test ./... rc=0
    RESULT registry missing nonclaims scan rc=0
    RESULT runner doctor rc=0
    RESULT same-state-integrated-dev baseline rc=0
    RESULT same-state-message-failure-dev baseline rc=0
    RESULT same-state-message-unsupported-dev baseline rc=0

v0.6.19A exploratory wrong recipient/device probe result:

    failure_case_family: wrong recipient/device/sidecar no false success
    Case A: Alice state + Alice sidecar did not falsely open/ack/drain Bob message.
    Case B: Bob state + Alice sidecar did not falsely open/ack/drain Bob message.
    Case C: Bob state + missing sidecar did not falsely open/ack/drain Bob message.
    Correct Bob path still opened and acked afterward.
    Ack count remained unchanged through wrong attempts.
    Bob inbox count remained unchanged through wrong attempts.
    Bob inbox was empty after correct open/ack.

Important interpretation:

    Wrong recipient/device/sidecar attempts may fail or skip at different layers.
    The shared invariant is no false success, no ack, no Bob inbox drain, and recovery still works.
    The correct patch target was a narrow validation profile, not runtime surgery.

v0.6.19B implementation added:

    tools/carbonstack-validate/same_state_message_recipient_failure_dev.go
    docs/214-v0.6.19-same-state-message-recipient-failure-dev-profile-v0.md
    registry entry runner.same-state-message-recipient-failure-dev
    main.go profile registration
    runner README update
    docs README update
    regenerated registry/COMMAND_REFERENCE.v0.md

Pre-profile validation passed:

    RESULT carbonstack diff --check rc=0
    RESULT runner go test ./... rc=0
    RESULT registry missing nonclaims scan rc=0
    RESULT registry lookup same-state recipient failure profile rc=0
    RESULT runner doctor rc=0

Profile validation passed:

    RESULT same-state-message-recipient-failure-dev --clean-generated rc=0

Profile result summary:

    PASS: Relay onboarding completed through KeyPackage -> add-member -> Welcome -> join
    PASS: normal message sent to Bob after same-state join
    PASS: Alice state + Alice sidecar did not falsely open/ack/drain Bob message
    PASS: Bob state + Alice sidecar did not falsely open/ack/drain Bob message
    PASS: Bob state + missing sidecar did not falsely open/ack/drain Bob message
    PASS: correct Bob sidecar open/ack still succeeded after wrong recipient/device attempts
    proof_level: same-state wrong recipient/device no false-success/no-ack/no-drain proof
    envelopes: 3
    envelope_acks: 2
    keypackage_delivery_state: queued
    welcome_delivery_state: acknowledged

Post-profile validation passed:

    RESULT generated reference deterministic check post-profile rc=0
    RESULT runner go test ./... post-profile rc=0
    RESULT same-state-integrated-dev still passes post-profile rc=0
    RESULT same-state-message-failure-dev still passes post-profile rc=0
    RESULT same-state-message-unsupported-dev still passes post-profile rc=0
    RESULT registry missing nonclaims scan post-profile rc=0
    RESULT runner doctor post-profile rc=0
    RESULT protected profile guard rc=0

Protected profile guard checked that these were not mutated:

    tools/carbonstack-validate/integrated_runtime_dev.go
    tools/carbonstack-validate/release_snapshot.go
    tools/carbonstack-validate/same_state_integrated_dev.go
    tools/carbonstack-validate/same_state_message_failure_dev.go
    tools/carbonstack-validate/same_state_message_unsupported_dev.go

Commit:

    197ea9f test: add same-state recipient failure profile

Commit diff summary:

    7 files changed, 674 insertions(+), 4 deletions(-)
    create mode 100644 docs/214-v0.6.19-same-state-message-recipient-failure-dev-profile-v0.md
    create mode 100644 tools/carbonstack-validate/same_state_message_recipient_failure_dev.go

Post-commit validation passed:

    RESULT generated reference deterministic check post-commit rc=0
    RESULT runner go test ./... post-commit rc=0
    RESULT registry lookup same-state recipient failure post-commit rc=0
    RESULT same-state-message-recipient-failure-dev post-commit rc=0

Generated reference:

    entries=83

Push:

    Scripted push failed authentication:
        RESULT push rc=128

    Manual retry from carbonstack succeeded:
        1f8086b..197ea9f main -> main

Final manual snapshot:

    carbonstack        197ea9f (HEAD -> main, origin/main, origin/HEAD)
    carbonstack-comms  53779d5 (HEAD -> main, origin/main, origin/HEAD)
    carbonstack-cypher d18a564 (HEAD -> main, origin/main, origin/HEAD)
    carbonstack-os     1bbbe52 (HEAD -> main, origin/main, origin/HEAD)

### 5.11 v0.6.20A/B/C Welcome join state-safety recon, patch, warning recovery

v0.6.20A Welcome/join failure no-ack recon:

    corrupt Welcome join with --ack-after-join failed as expected;
    no status: joined;
    no welcome_acked: true;
    no ack_delivery_state: acknowledged;
    total ack count unchanged;
    Welcome ack rows unchanged;
    Bob Relay inbox remained queued.

But v0.6.20A also found:

    restored valid Welcome join failed with conversation_already_exists.

v0.6.20B deep recon:

    confirmed the failure was final Bob sidecar conversation state poisoning;
    confirmed manual quarantine of the final conversation directory recovered restored valid join;
    recommended sidecar_conversation_join_atomic_staging_promote as fix target.

v0.6.20C implementation first pass:

    added remove_dir_all_if_exists helper;
    delayed final/staging dir creation until after initial preflight;
    added staged materialization and atomic promote;
    appended invalid-Welcome regression test;
    added v0.6.20 docs note;
    sidecar cargo test passed but emitted an unused-variable warning;
    targeted protocol tests passed;
    full Comms tests passed;
    CarbonStack generated-reference check remained current at entries=83;
    runner tests passed;
    registry missing-nonclaims scan passed;
    runner doctor passed.

First-pass blunder:

    Rust emitted:
        warning: unused variable: `reloaded_member_count`

    Runner-driven sidecar JSON parsing then failed because warning text appeared before JSON:
        invalid character 'w' looking for beginning of value

v0.6.20C recovery:

    renamed the unused binding to _reloaded_member_count;
    required sidecar cargo test to be warning-free;
    sidecar cargo test passed warning-free;
    targeted protocol test passed;
    Comms go test ./... passed;
    Comms diff --check passed;
    CarbonStack diff --check passed;
    generated reference check passed with entries=83;
    runner go test ./... passed;
    registry missing-nonclaims scan passed;
    runner doctor passed;
    same-state-integrated-dev passed;
    same-state-message-failure-dev passed;
    same-state-message-unsupported-dev passed;
    same-state-message-recipient-failure-dev passed.

Committed:

    b42aa09 fix: make Welcome join state writes atomic
        2 files changed, 241 insertions(+), 61 deletions(-)

    e1f6b8c docs: record Welcome join partial-state fix
        2 files changed, 66 insertions(+)
        created docs/215-v0.6.20-welcome-join-partial-state-safety-v0.md

Post-commit validation:

    post-commit Comms go test ./... passed;
    post-commit generated reference check passed;
    post-commit runner go test ./... passed;
    post-commit same-state-integrated-dev passed.

Push:

    scripted carbonstack-comms push failed with the recurring auth quirk;
    scripted carbonstack push succeeded;
    manual carbonstack push reported everything up-to-date;
    manual carbonstack-comms push succeeded.


### 5.12 v0.6.21A/B Welcome join failure profile implementation and assertion recovery

v0.6.21A recon validation passed:

    generated reference deterministic check: passed at 83 entries
    runner go test ./...: passed
    registry missing nonclaims scan: passed with matches 0
    runner doctor: passed
    sidecar cargo test warning-free: passed
    comms go test ./...: passed
    fixed-behavior probe: passed
    same-state-integrated-dev: passed after recon
    same-state-message-failure-dev: passed after recon
    same-state-message-unsupported-dev: passed after recon
    same-state-message-recipient-failure-dev: passed after recon

v0.6.21A fixed-behavior probe result:

    corrupt Welcome join failed with bad_join_rc: 1
    corrupt Welcome join did not ack
    corrupt Welcome join did not drain Bob Relay inbox
    corrupt Welcome join left no final/staging Bob conversation state
    restored Welcome join succeeded with good_join_rc: 0
    restored Welcome acked after successful join
    Bob conversation was reloadable after restored join

Initial v0.6.21B implementation validation:

    baseline generated reference check: passed at 83 entries
    runner go test ./...: passed
    registry missing nonclaims scan: passed
    runner doctor: passed
    sidecar cargo test warning-free: passed
    comms go test ./...: passed
    profile files/docs/registry added
    generated command reference regenerated to 84 entries
    diff --check: passed
    runner go test ./...: passed
    registry missing nonclaims scan: passed
    registry lookup runner.same-state-welcome-join-failure-dev: passed
    runner doctor: passed
    same-state-welcome-join-failure-dev failed due to profile assertion bug

Failure:

    VALIDATION FAILED: add-member output missing "content_type"

Cause:

    profile asserted `content_type` should appear in add-member output;
    add-member Welcome path does not print literal `content_type`;
    it does print Welcome-specific envelope, payload, delivery, and ack evidence.

Recovery validation:

    removed over-strict `content_type` assertion
    generated reference deterministic check remained current at 84 entries
    runner go test ./... passed
    registry missing nonclaims scan passed
    registry lookup runner.same-state-welcome-join-failure-dev passed
    runner doctor passed
    same-state-welcome-join-failure-dev --clean-generated passed
    same-state-integrated-dev passed
    same-state-message-failure-dev passed
    same-state-message-unsupported-dev passed
    same-state-message-recipient-failure-dev passed
    registry missing nonclaims scan post-profile passed
    runner doctor post-profile passed
    protected profile guard passed
    cached diff --check passed
    commit succeeded
    post-commit generated reference check passed
    post-commit runner go test ./... passed
    post-commit registry lookup passed
    post-commit same-state-welcome-join-failure-dev passed

Commit:

    3876fc8 test: add same-state Welcome join failure profile

Final push:

    scripted push rc=128
    manual git push succeeded

Final pushed heads:

    carbonstack        943e0e4 test: add same-state malformed payload profile
    carbonstack-comms  b42aa09 fix: make Welcome join state writes atomic
    carbonstack-cypher d18a564 chore: restore validated Go module floor
    carbonstack-os     1bbbe52 docs: clarify CarbonStackOS target direction

### 5.13 v0.6.22A/B/C stale-state authority recon and docs checkpoint

v0.6.22A initial recon:

    Baseline validation passed:
        generated reference deterministic check passed;
        runner go test ./... passed;
        registry missing-nonclaims scan passed;
        runner doctor passed;
        sidecar cargo test warning-free passed;
        Comms go test ./... passed;
        existing same-state profiles were exercised before the probe.
    The first stale-state probe failed before semantic state testing.
    Cause:
        generated message_label included the long case name and exceeded the sidecar label limit.
    Classification:
        probe hygiene failure, not a stale-state behavior finding.

v0.6.22A short-label recovery recon:

    Replaced long human-readable labels with short per-case IDs.
    Reached the stale-state cases.
    Initial conservative result:
        cases_total: 10;
        cases_with_open_ack_or_drain: 6.
    Cases that opened/acked/drained:
        delete/corrupt conversation-summary.json;
        delete/corrupt join-summary.json;
        delete/corrupt device-level provider-storage.json.
    Cases that did not open/ack/drain:
        delete/corrupt/truncate conversation provider-storage.json;
        moved conversation directory.
    Interpretation:
        this was not automatically a bug;
        it exposed the need to classify state authority.

v0.6.22B state-authority classification recon:

    cases_total: 12
    critical_runtime_bypass_cases: 0
    cases_with_target_rewrite_or_regeneration: 0

    Runtime-authority cases blocked open/ack/drain:
        s07 delete_conversation_provider_storage;
        s08 invalid_conversation_provider_storage_json;
        s09 truncate_conversation_provider_storage;
        s10 move_conversation_dir_as_stale_backup.

    Metadata/evidence or non-runtime existing-message-open cases opened/acked:
        s01 delete_conversation_summary;
        s02 invalid_conversation_summary_json;
        s03 valid_wrong_conversation_summary_json;
        s04 delete_join_summary;
        s05 invalid_join_summary_json;
        s06 valid_wrong_join_summary_json;
        s11 delete_device_provider_storage;
        s12 invalid_device_provider_storage_json.

    Important:
        no target changed by load-check;
        no target changed by message-inbox;
        therefore no tested silent regeneration/repair occurred.

v0.6.22C docs/model checkpoint:

    Added docs/217-v0.6.22-sidecar-state-authority-classification-v0.md.
    Updated docs/README.md.
    Committed:
        8dd7db5 docs: classify sidecar state authority.
    Validation passed:
        diff --check;
        generated reference check remained entries=84;
        runner go test ./...;
        registry missing-nonclaims scan;
        runner doctor;
        post-commit generated-reference check;
        post-commit runner go test ./...;
        post-commit same-state-welcome-join-failure-dev.
    Scripted push failed with rc=128, manual push succeeded.



### 5.14 v0.6.26 roadmap/output-semantics scout, cleanup, and recovery

v0.6.26 validation happened across four linked sub-rungs:

**v0.6.26 roadmap refresh**

    Long-form adversarial Q/A completed.
    CarbonStack_Long_Term_Roadmap_v0.6.26_EVERGREEN generated.
    Roadmap kept pre-alpha status and relative-lane planning mode.

**v0.6.26B output semantics scout**

    generated reference deterministic check: pass, entries=86
    runner go test ./...: pass
    registry missing-nonclaims scan: pass
    runner doctor: pass
    sidecar cargo test --quiet: pass, warning-free
    Comms go test ./...: pass
    live output semantics probe: pass
    same-state-message-replay-classification-dev: pass
    same-state-message-malformed-payload-dev: pass

Scout result:

    from_device present without explicit unverified identity fields;
    load-check output needed provider-vs-summary separation.

**v0.6.26C sender metadata cleanup**

Patch result:

    message_wrappers_dev.go patched;
    openmls_runtime.go patched;
    message_wrappers_dev_test.go patched;
    docs/221-v0.6.26-sender-metadata-output-warning-v0.md added;
    docs/README.md updated.

Validation:

    Comms diff --check: pass
    CarbonStack diff --check: pass
    generated reference deterministic check: pass, entries=86
    Comms go test ./...: pass
    runner go test ./...: pass
    registry missing-nonclaims scan: pass
    runner doctor: pass
    same-state-integrated-dev: pass
    same-state-message-failure-dev: pass
    same-state-message-unsupported-dev: pass
    same-state-message-malformed-payload-dev: pass
    same-state-message-replay-classification-dev: pass
    same-state-message-recipient-failure-dev: pass
    same-state-welcome-join-failure-dev: pass
    no registry/reference/runner/full/release-snapshot mutation: pass
    post-commit generated reference check: pass
    post-commit Comms tests: pass
    post-commit runner tests: pass
    post-commit replay and malformed-payload profiles: pass

Push:

    carbonstack push succeeded to 499ab6a.
    carbonstack-comms push initially failed with the recurring auth rc=128; the later v0.6.26D push advanced origin from b42aa09 to 2d13c45, carrying both 61b2e51 and 2d13c45.

**v0.6.26D load-check provider/summary cleanup**

Initial implementation failed mid-run:

    success-path provider/summary output was not actually inserted because the script falsely treated the error-path provider_reloadable string as evidence the success path was patched.

First recovery failed validation:

    it inserted fields with literal \n sequences, collapsing output lines;
    it overpatched add-member and join tests with load-check-only expectations;
    Comms go test ./... failed in internal/app.

Clean reapply succeeded:

    reset only the partial v0.6.26D files;
    preserved committed v0.6.26C;
    reinserted load-check success-path fields cleanly;
    patched only the load-check unit test;
    re-added docs/222 and README entry;
    live clean probe passed:
        baseline_output_has_provider_summary_fields: yes
        missing_summary_output_has_metadata_missing_fields: yes
        missing_summary_rc: 1

Validation:

    live load-check output probe clean: pass
    Comms diff --check: pass
    CarbonStack diff --check: pass
    generated reference deterministic check: pass, entries=86
    Comms go test ./...: pass
    runner go test ./...: pass
    registry missing-nonclaims scan: pass
    runner doctor: pass
    same-state-integrated-dev: pass
    same-state-message-malformed-payload-dev: pass
    same-state-message-replay-classification-dev: pass
    same-state-welcome-join-failure-dev: pass
    registry/reference/runner guard: pass
    post-commit generated reference check: pass
    post-commit Comms tests: pass
    post-commit runner tests: pass

Push:

    carbonstack-comms pushed b42aa09..2d13c45.
    carbonstack pushed 499ab6a..468432b.
    push rc=0 for both repos.

## 6. Current Behavior After v0.6.22

### 6.0 v0.6.21 behavior-change summary

v0.6.21 does not change Comms/Cypher/OS runtime source. It changes CarbonStack validation/docs/registry state by adding the live same-state Welcome join failure profile.

New behavior surface:

    same-state-welcome-join-failure-dev

It proves the v0.6.20 sidecar atomic state-write fix in the same-state Relay onboarding context:

    corrupt Welcome join
    -> fails without success markers
    -> no ack
    -> no Relay inbox drain
    -> no final/staging Bob conversation state
    -> Bob conversation load-check fails
    -> restored valid Welcome joins with the same conversation label
    -> Welcome is acked only after successful join
    -> Bob Relay inbox drains
    -> Bob conversation load-check succeeds

Generated command reference state:

    entries=84

New registry row:

    runner.same-state-welcome-join-failure-dev

Existing v0.6.20 behavior remains in force:

    sidecar conversation-join materializes state under staging and promotes only after successful materialization/reload-check.

Still unchanged:

    full
    release-snapshot
    integrated-runtime-dev
    release-package validation
    package-root validation
    Comms/Cypher/OS runtime source
    public release assets

### 6.1 Sidecar atomic staging/promotion

Patch target:

    carbonstack-comms/internal/protocol/mls/openmls-sidecar/src/state.rs

Main affected function:

    join_dev_conversation

Expected materialization flow:

    check final conversation path absent;
    create/use staging conversation path;
    write staged provider-storage.json;
    reload staged provider storage;
    write staged conversation-summary.json;
    write staged join-summary.json;
    promote staging directory to final conversation directory only after success;
    remove staging on materialization error.

### 6.2 Regression test coverage

Test added:

    TestOpenMLSSidecarConversationJoinInvalidWelcomeDoesNotPoisonFinalState

The regression test proves:

    corrupt Welcome -> conversation-join fails;
    corrupt Welcome -> no final Bob conversation state exists;
    restored valid Welcome -> conversation-join succeeds with same label;
    restored valid Welcome -> joined true;
    restored valid Welcome -> group_reloadable true;
    final conversation-summary/provider-storage/join-summary exist after success;
    sidecar stdout does not leak secret material.

### 6.3 v0.6.20 docs note

New CarbonStack doc:

    docs/215-v0.6.20-welcome-join-partial-state-safety-v0.md

This doc records:

    the v0.6.20A/v0.6.20B finding;
    why the fix target was sidecar state atomicity;
    the staging/promote behavior;
    the validation meaning;
    the nonclaims.

### 6.4 Existing same-state profiles retained

No existing same-state profile was renamed, broadened, or folded into `full`.

Retained profiles:

    same-state-integrated-dev
        positive-path same-state proof.

    same-state-message-failure-dev
        wrong-conversation no-ack/no-drain proof.

    same-state-message-unsupported-dev
        unsupported normal-message content-type no-ack/no-drain proof.

    same-state-message-recipient-failure-dev
        wrong recipient/device/sidecar no-false-success/no-ack/no-drain proof.

### 6.5 Generated command reference state

The generated command reference remains:

    registry/COMMAND_REFERENCE.v0.md
    entries=83

v0.6.20 did not add a registry row and did not change the command reference count.

Freshness guard remains:

    cd carbonstack
    python3 tools/registry/render-command-reference.py --check

Runner test guard remains:

    cd carbonstack/tools/carbonstack-validate
    go test ./...

### 6.6 Still missing after v0.6.21

A live same-state runner profile for Welcome/join failure has not been added yet.

Next profile target should prove:

    corrupt Welcome join with --ack-after-join
    -> no ack
    -> no Relay inbox drain
    -> no final Bob conversation state poisoning
    -> restored valid Welcome join/ack succeeds

Candidate future profile name:

    same-state-welcome-join-failure-dev

or:

    relay-join-failure-noack-dev

Recommendation remains:

    same-state-welcome-join-failure-dev

because the profile would use the same coherent temp universe but target Relay onboarding artifact semantics.

## 7. Message Flow Contract

### 7.1 Current normal-message path

Recommended dev/pre-alpha normal-message wrapper commands remain:

    message-send-dev
    message-inbox-dev

Lower-level direct OpenMLS commands remain:

    openmls-send-dev
    openmls-inbox-dev

Normal-message wrapper proof profiles now include:

    dev-runtime-openmls-wrappers
    same-state-integrated-dev
    same-state-message-failure-dev
    same-state-message-unsupported-dev
    same-state-message-recipient-failure-dev

### 7.2 Lower-level direct OpenMLS path

Direct lower-level commands remain useful for implementation visibility and debugging:

    openmls-send-dev
    openmls-inbox-dev

They are not the preferred human-facing dev/pre-alpha normal-message surface after wrapper introduction.

### 7.3 Relay onboarding/artifact path

Relay onboarding/artifact commands remain separate from normal message inbox behavior:

    openmls-relay-keypackage-submit-dev
    openmls-relay-keypackage-inbox-dev
    openmls-relay-add-member-dev
    openmls-relay-welcome-submit-dev
    openmls-relay-welcome-inbox-dev
    openmls-relay-join-dev

KeyPackage and Welcome are onboarding artifacts, not ordinary application messages.

### 7.4 Integrated runtime profiles

Existing sequential integrated profile:

    integrated-runtime-dev

Meaning:

    relay-openmls-join-dev -> dev-runtime-openmls-wrappers

Boundary:

    sequential composition;
    live umbrella only;
    not full;
    not release-snapshot;
    not package-root validation;
    not same-state proof.

Positive same-state profile:

    same-state-integrated-dev

Meaning:

    KeyPackage -> add-member -> Welcome -> join -> message-send-dev -> message-inbox-dev --ack

Boundary:

    live-dev positive-path same-state proof;
    not full;
    not release-snapshot;
    not release-package validation;
    not package-root validation;
    not adversarial proof.

### 7.5 Same-state wrong-conversation profile

Profile:

    same-state-message-failure-dev

Meaning:

    wrong-conversation message-inbox-dev --ack
    -> message-open failure
    -> no ack
    -> no Bob inbox drain
    -> correct conversation open/ack still works

Boundary:

    wrong-conversation-only.

### 7.6 Same-state unsupported normal-message profile

Profile:

    same-state-message-unsupported-dev

Meaning:

    unsupported normal application-message content_type
    -> unsupported skip
    -> no ack
    -> no Bob inbox drain
    -> restore valid content_type
    -> correct open/ack still works

Boundary:

    unsupported normal-message content-type only.

### 7.7 Same-state recipient/device failure profile

Profile:

    same-state-message-recipient-failure-dev

Meaning:

    wrong recipient/device/sidecar attempts
    -> no false message-open success
    -> no ack
    -> no Bob inbox drain
    -> correct Bob open/ack still works

Boundary:

    covers Alice state + Alice sidecar, Bob state + Alice sidecar, and Bob state + missing sidecar in the same-state proof universe.
    Does not prove identity verification, malicious relay resistance, or hostile-server safety.

### 7.8 Welcome join state-safety contract

v0.6.20 adds a sidecar state-safety contract for failed Welcome join:

    failed/invalid Welcome join must not create durable final joined conversation state.

Ack doctrine and state doctrine are separate:

    Ack doctrine:
        no Welcome ack unless join succeeds.

    State doctrine:
        no final Bob joined conversation path unless join state successfully materializes and reload-checks.

The v0.6.20 patch addresses the state doctrine.

A later runner profile should combine both:

    corrupt Welcome fails;
    no ack;
    no Relay inbox drain;
    no final state poisoning;
    restored Welcome succeeds and acks.

### 7.9 Same-state Welcome join failure profile

`same-state-welcome-join-failure-dev` is the live-umbrella corrupt-Welcome onboarding failure proof profile.

It proves:

    KeyPackage -> add-member -> queued Welcome
    corrupt queued Welcome payload
    -> Bob openmls-relay-join-dev --ack-after-join fails
    -> no Welcome ack
    -> no Bob Relay inbox drain
    -> no final/staging Bob conversation state
    -> Bob conversation load-check fails
    restore valid Welcome payload
    -> Bob openmls-relay-join-dev --ack-after-join succeeds
    -> Welcome is acked
    -> Bob Relay inbox drains
    -> Bob conversation load-check succeeds

Registry ID:

    runner.same-state-welcome-join-failure-dev

Command:

    go run . --profile same-state-welcome-join-failure-dev --root <umbrella root> --clean-generated

Nonclaims:

    not full;
    not release-snapshot;
    not release-package validation;
    not package-root validation;
    not adversarial relay harness;
    not hostile-server safety;
    not identity verification;
    not production secure messaging;
    not stale provider state modeling;
    not vault/PQ work.

### 7.10 Ack doctrine

Application-message ack rule:

    no ack on fetch;
    no ack for unsupported envelope;
    no ack on artifact-write failure;
    no ack on sidecar message-open failure;
    no ack on wrong-conversation/group open failure;
    no ack on wrong recipient/device/sidecar false path;
    ack only after successful sidecar message-open;
    ack only when --ack is explicit.

Relay onboarding ack rule:

    KeyPackage fetch/write alone is not ack eligibility.
    add-member success alone does not ack KeyPackage.
    Welcome fetch/write alone is not ack eligibility.
    Welcome join success/reloadable state may ack only under explicit flag.
    failed join must not ack.

v0.6.19 strengthens the application-message side by proving wrong recipient/device/sidecar attempts do not falsely open, ack, or drain.

## 8. Critical Paths

### 8.1 Repo roots

Umbrella root:

    ~/repos/carbonstack_umbrella

Repos:

    ~/repos/carbonstack_umbrella/carbonstack
    ~/repos/carbonstack_umbrella/carbonstack-comms
    ~/repos/carbonstack_umbrella/carbonstack-cypher
    ~/repos/carbonstack_umbrella/carbonstack-os

### 8.0 Current v0.6.25 paths

Current v0.6.25 docs/model checkpoint paths:

    carbonstack/docs/220-v0.6.25-normal-message-envelope-metadata-classification-v0.md
    carbonstack/docs/README.md

Current v0.6.24 replay classification profile paths retained:

    carbonstack/tools/carbonstack-validate/same_state_message_replay_classification_dev.go
    carbonstack/docs/219-v0.6.24-loadcheck-replay-classification-v0.md
    carbonstack/registry/commands.v0.yaml
    carbonstack/registry/COMMAND_REFERENCE.v0.md

Current command reference state:

    entries=86
    v0.6.25 did not add a registry entry or runner profile.

### 8.2 New v0.6.21 paths

New v0.6.21 runner profile:

    carbonstack/tools/carbonstack-validate/same_state_welcome_join_failure_dev.go

New v0.6.21 docs:

    carbonstack/docs/216-v0.6.21-same-state-welcome-join-failure-dev-profile-v0.md

Updated v0.6.21 registry/reference paths:

    carbonstack/registry/commands.v0.yaml#runner.same-state-welcome-join-failure-dev
    carbonstack/registry/COMMAND_REFERENCE.v0.md

Updated v0.6.21 runner/docs surfaces:

    carbonstack/tools/carbonstack-validate/main.go
    carbonstack/tools/carbonstack-validate/README.md
    carbonstack/docs/README.md

### 8.3 New v0.6.20 paths

    carbonstack/docs/215-v0.6.20-welcome-join-partial-state-safety-v0.md
    carbonstack-comms/internal/protocol/mls/openmls-sidecar/src/state.rs
    carbonstack-comms/internal/protocol/openmls_sidecar_conversation_test.go

New v0.6.20 commits:

    carbonstack:
        e1f6b8c docs: record Welcome join partial-state fix

    carbonstack-comms:
        b42aa09 fix: make Welcome join state writes atomic

No new v0.6.20 registry entry exists yet.

Future candidate profile path if/when added:

    carbonstack/tools/carbonstack-validate/same_state_welcome_join_failure_dev.go
    carbonstack/docs/216-v0.6.21-same-state-welcome-join-failure-dev-profile-v0.md

### 8.4 v0.6.19 paths

New profile implementation:

    carbonstack/tools/carbonstack-validate/same_state_message_recipient_failure_dev.go

New profile doc:

    carbonstack/docs/214-v0.6.19-same-state-message-recipient-failure-dev-profile-v0.md

New registry row:

    carbonstack/registry/commands.v0.yaml#runner.same-state-message-recipient-failure-dev

Generated command reference updated:

    carbonstack/registry/COMMAND_REFERENCE.v0.md

Runner README updated:

    carbonstack/tools/carbonstack-validate/README.md

Docs README updated:

    carbonstack/docs/README.md

### 8.5 v0.6.18 paths

Unsupported content-type profile implementation:

    carbonstack/tools/carbonstack-validate/same_state_message_unsupported_dev.go

Profile doc:

    carbonstack/docs/213-v0.6.18-same-state-message-unsupported-dev-profile-v0.md

Registry row:

    carbonstack/registry/commands.v0.yaml#runner.same-state-message-unsupported-dev

### 8.6 v0.6.17 paths

Wrong-conversation failure profile implementation:

    carbonstack/tools/carbonstack-validate/same_state_message_failure_dev.go

Profile doc:

    carbonstack/docs/212-v0.6.17-same-state-message-failure-dev-profile-v0.md

Registry row:

    carbonstack/registry/commands.v0.yaml#runner.same-state-message-failure-dev

### 8.7 v0.6.16 same-state positive paths

Same-state positive profile implementation:

    carbonstack/tools/carbonstack-validate/same_state_integrated_dev.go

Profile doc:

    carbonstack/docs/211-v0.6.16-same-state-integrated-dev-profile-v0.md

Registry row:

    carbonstack/registry/commands.v0.yaml#runner.same-state-integrated-dev

### 8.8 Carried-forward registry/generated docs

Command registry:

    carbonstack/registry/commands.v0.yaml

Generated command reference:

    carbonstack/registry/COMMAND_REFERENCE.v0.md

Command reference renderer:

    carbonstack/tools/registry/render-command-reference.py

Generated reference freshness test:

    carbonstack/tools/carbonstack-validate/generated_command_reference_test.go

Command boundary table:

    carbonstack/registry/COMMAND_BOUNDARY_TABLE.v0.md

Registry README:

    carbonstack/registry/README.md

### 8.9 Policy docs

Post-history sanity doc:

    carbonstack/docs/209-v0.6.14-post-history-rewrite-sanity-v0.md

Same-state plan doc:

    carbonstack/docs/210-v0.6.15-same-state-integrated-proof-plan-v0.md

Integrated runtime policy doc:

    carbonstack/docs/206-v0.6.10-integrated-runtime-policy-v0.md

Relay onboarding boundary contract:

    carbonstack/docs/205-v0.6.9-relay-onboarding-boundary-contract-v0.md

Command reference generation policy:

    carbonstack/docs/207-v0.6.11-command-reference-generation-policy-v0.md

Registry metadata hardening doc:

    carbonstack/docs/208-v0.6.12-registry-metadata-hardening-v0.md

### 8.10 Runner paths

Runner root:

    carbonstack/tools/carbonstack-validate

Important runner files:

    main.go
    registry_lookup.go
    integrated_runtime_dev.go
    relay_openmls_join_dev.go
    dev_runtime_openmls.go
    dev_runtime_openmls_wrappers.go
    same_state_integrated_dev.go
    same_state_message_failure_dev.go
    same_state_message_unsupported_dev.go
    same_state_message_recipient_failure_dev.go
    release_snapshot.go
    command_registry_test.go
    generated_command_reference_test.go

### 8.11 Comms paths

Comms root:

    carbonstack-comms

OpenMLS sidecar:

    carbonstack-comms/internal/protocol/mls/openmls-sidecar

OpenMLS sidecar generated dev state root:

    carbonstack-comms/internal/protocol/mls/openmls-sidecar/.carbonstack-openmls-sidecar-state

OpenMLS sidecar build artifact root:

    carbonstack-comms/internal/protocol/mls/openmls-sidecar/target

Message wrapper implementation/tests:

    carbonstack-comms/internal/app/message_wrappers_dev.go
    carbonstack-comms/internal/app/message_wrappers_dev_test.go

OpenMLS runtime implementation/tests:

    carbonstack-comms/internal/app/openmls_runtime.go
    carbonstack-comms/internal/app/openmls_runtime_test.go

Relay OpenMLS implementation/tests:

    carbonstack-comms/internal/app/openmls_relay_dev.go
    carbonstack-comms/internal/app/openmls_relay_dev_test.go

Relay smoke script:

    carbonstack-comms/scripts/openmls-relay-narrow-join-smoke-dev.sh

### 8.12 Cypher paths

Cypher root:

    carbonstack-cypher

Cypher command:

    carbonstack-cypher/cmd/cypher

Cypher migrations:

    carbonstack-cypher/migrations

Ignored local dev DB:

    carbonstack-cypher/cypher.db

### 8.13 Private/history paths

Private freeze tarball:

    ~/CarbonStack_Umbrella_v0.6.12FREEZE.tar.gz

Private freeze tarball SHA-256:

    ~/CarbonStack_Umbrella_v0.6.12FREEZE.tar.gz.sha256

Freeze SHA-256:

    9ce446d1a4c33389a97bb4870b945b38d5704527951c04fdcdaff0c1f10a65da

Boundary:

    private rollback/history safety artifact only; not release material.

## 9. Validation Commands Worth Carrying Forward

v0.6.20 state-safety / warning-clean validation:

    cd ~/repos/carbonstack_umbrella/carbonstack-comms/internal/protocol/mls/openmls-sidecar
    cargo test --quiet

    cd ~/repos/carbonstack_umbrella/carbonstack-comms
    go test ./internal/protocol -run 'TestOpenMLSSidecarConversationJoinWelcomeConsume|TestOpenMLSSidecarConversationJoinInvalidWelcomeDoesNotPoisonFinalState' -count=1
    go test ./... -count=1

    cd ~/repos/carbonstack_umbrella/carbonstack
    python3 tools/registry/render-command-reference.py --check

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go test ./... -count=1
    go run . --profile registry-lookup --root ~/repos/carbonstack_umbrella --list --missing-nonclaims
    go run . --profile doctor --root ~/repos/carbonstack_umbrella

Existing same-state profile preservation checks:

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go run . --profile same-state-integrated-dev --root ~/repos/carbonstack_umbrella --clean-generated
    go run . --profile same-state-message-failure-dev --root ~/repos/carbonstack_umbrella --clean-generated
    go run . --profile same-state-message-unsupported-dev --root ~/repos/carbonstack_umbrella --clean-generated
    go run . --profile same-state-message-recipient-failure-dev --root ~/repos/carbonstack_umbrella --clean-generated

Carried-forward pre-v0.6.20 commands:

Generated command reference freshness:

    cd ~/repos/carbonstack_umbrella/carbonstack
    python3 tools/registry/render-command-reference.py --check

Runner tests:

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go test ./... -count=1

Registry missing nonclaims:

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go run . --profile registry-lookup --root ~/repos/carbonstack_umbrella --list --missing-nonclaims

Registry lookup for new v0.6.19 profile:

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go run . --profile registry-lookup --root ~/repos/carbonstack_umbrella --registry-id runner.same-state-message-recipient-failure-dev

Doctor:

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go run . --profile doctor --root ~/repos/carbonstack_umbrella

Positive same-state profile:

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go run . --profile same-state-integrated-dev --root ~/repos/carbonstack_umbrella --clean-generated

Wrong-conversation failure profile:

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go run . --profile same-state-message-failure-dev --root ~/repos/carbonstack_umbrella --clean-generated

Unsupported content-type failure profile:

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go run . --profile same-state-message-unsupported-dev --root ~/repos/carbonstack_umbrella --clean-generated

Recipient/device failure profile:

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go run . --profile same-state-message-recipient-failure-dev --root ~/repos/carbonstack_umbrella --clean-generated

Release-package validation remains package-root only:

    go run . --profile verify-checksums --root <fresh extracted package root>
    go run . --profile full --root <fresh extracted package root> --clean-generated

Do not conflate live-umbrella same-state profiles with release-package `full`.


New v0.6.21 profile:

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go run . --profile same-state-welcome-join-failure-dev --root ~/repos/carbonstack_umbrella --clean-generated

Registry lookup:

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go run . --profile registry-lookup --root ~/repos/carbonstack_umbrella --registry-id runner.same-state-welcome-join-failure-dev

## 10. Known Residuals and Future To-Do

### 10.1 Immediate next safest action

Proceed to **adversarial QA for evergreen roadmap refresh**.

Reason:

    The current same-state normal-message failure-hardening lane is now substantially covered.
    v0.6.17 through v0.6.25 added or documented:
      wrong-conversation failure;
      unsupported content-type skip;
      malformed payload rejection/recovery;
      wrong recipient/device/sidecar failure;
      Welcome join failure no-ack/no-drain/no-state-poison;
      stale state authority classification;
      normal application-message replay classification;
      envelope metadata classification.
    The next useful step is not another narrow profile by default.
    The next useful step is to interrogate the roadmap and decide what remains in v0.6.x before vault/PQ/adversarial harness work.

Target questions for the next rung:

    Which v0.6.x goals are complete, mostly complete, active, deferred, or reframed?
    Does the load-check/provider-vs-summary semantic mismatch need code cleanup before v0.7.0?
    Should sender metadata warning/output be patched before identity/trust modeling?
    Which Welcome/KeyPackage malformed/replay cases should be v0.6.x vs adversarial-harness work?
    What exact model must exist before bounded vault substrate/stub work?
    What is the honest v0.7.0 minor-epoch release definition?

Boundary:

    QA/modeling first;
    no vault implementation yet;
    no PQ implementation yet;
    no adversarial harness implementation yet;
    no hostile-server safety claim.

### 10.2 Near-term follow-up after v0.6.25

Likely near-term follow-ups after adversarial QA:

    roadmap refresh;
    load-check output/semantic cleanup if deemed release-critical;
    sender metadata/from_device warning or docs hardening if deemed release-critical;
    vault/security model with no-silent/backup/recovery rules;
    bounded vault substrate/stub only after the model exists.

### 10.3 Still deferred

Still deferred after v0.6.25:

    no adversarial flag yet;
    no adversarial relay harness yet;
    no Welcome/KeyPackage replay profile yet;
    no server equivocation model yet;
    no inclusion of same-state profiles in full/release-snapshot yet;
    no full-runtime-dev / full-validate-release split yet;
    no vault implementation yet;
    no PQ/hybrid implementation yet;
    no public release asset rebuild yet;
    no Unix man-page generation yet.

### 10.4 Later roadmap order

Preferred larger order after v0.6.25:

    adversarial QA for evergreen roadmap refresh;
    vault/security model plus no-silent/backup/recovery rules;
    bounded vault substrate/stub;
    PQ/hybrid placement model;
    secured vault/PQ implementation planning;
    adversarial harness;
    later full-runtime-dev / full-validate-release semantic split;
    later package-root rehearsal and release-readiness checks;
    v0.7.0 cumulative minor-epoch release.

### 10.5 History rewrite residuals

Same as before:

    pre-v0.6.13 hashes in older docs are historical;
    public main/tag history after the rewrite is authoritative;
    avoid reusing old pre-rewrite commit hashes for current mainline claims.


## 11. Critical Blunders / Lessons to Carry Forward

### 11.1 Dry-run generator grouping bug

Initial v0.6.13 dry-run had:

    risk_issue_count=7

Cause:

    script rows were classified after generic Comms/legacy rules.

Fix:

    classify `kind: script` / `*.script.*` before generic legacy and Comms state/trust grouping.

Lesson:

    Renderer grouping order is a policy decision, not just formatting.
    Script/helper rows must not be promoted by accident.

### 11.2 Shell paste-tail corruption remains recurring

Several pasted shell blocks displayed tail corruption or smashed endings. The reliable signals remain:

    explicit RESULT markers;
    final git status;
    final log heads;
    final validation output;
    final localdomain counts;
    final remote branch cleanup output;
    final manual snapshots.

v0.6.14 repeated the pattern: the script finished the important commit/validation work, but the terminal log had paste tail corruption and an interrupted final explanatory heredoc.

### 11.3 Git identity/history cleanup remained multi-stage

Carried from v0.6.13:

    git-filter-repo was not initially installed;
    mailmap rewrite did not clean all public identities;
    commit-callback rewrite worked;
    pipefail + zero grep match aborted one intermediate script;
    tag force-with-lease rejected rewritten tags as stale;
    explicit per-tag force push fixed the tag update;
    remote backup branches had to be deleted because they preserved old public identities.

v0.6.14 follow-through:

    created a private all-ref bundle freeze before deleting local-only backup branches.

### 11.4 Package checksum context trap

During v0.6.14 release asset sanity, running package checksums from the asset directory produced missing-file errors.

Interpretation:

    carbonstack-v0.6.0-package-checksums.txt is package-root-relative and must be checked from the extracted package root or through the runner.

Resolution:

    extracted attached package passed verify-checksums and full --clean-generated, so release package sanity passed.

Lesson:

    Use runner package-root validation as authoritative release package validation.

### 11.5 Scripted push auth retry

The v0.6.14 scripted push failed authentication:

    Authentication failed for carbonstack.git

Manual retry succeeded:

    c5bc95c..77ad2d7 main -> main

Lesson:

    Preserve manual push retry snapshots when Gitea credentials prompt/caching causes scripted push failure.

### 11.6 Manual final status with unset ROOT failed

After the manual push, an attempted final block using `$ROOT` outside the original script context failed because `ROOT` was unset.

Resolution:

    The explicit final loop using `cd ~/repos/carbonstack_umbrella/$repo` showed the authoritative final heads.

Lesson:

    After interrupted scripts, use self-contained final snapshot commands that re-declare ROOT or use explicit paths.

### 11.7 v0.6.15 docs-only checkpoint correctly avoided profile-name drift

The v0.6.15 decision could have jumped straight into adding `same-state-integrated-dev`.

Instead, it recorded a plan first.

Lesson:

    same-state is a claim boundary, not just a stronger smoke test.
    The implementation must prove state continuity before a named profile is committed.
    Use temporary runner-side probe code first; commit a profile only after evidence passes.

### 11.8 Scripted push auth retry repeated

The v0.6.15 scripted push failed authentication, matching prior Gitea credential-cache friction.

Manual retry succeeded:

    77ad2d7..c5434b4 main -> main

Lesson:

    Keep scripted validation and manual push retry snapshots separate.
    Treat final explicit `git log` / `git status` snapshots as authoritative after push auth interruptions.

---

### 11.9 v0.6.16 registry YAML indentation failure

The first v0.6.16B implementation script inserted the new registry row but malformed the top-level YAML indentation around `runner.registry-lookup`.

Failure:

    TestCommandRegistryHasNoTopLevelEntryIDs failed.

Root cause:

    The insertion matched `- id: runner.registry-lookup` without preserving the required two-space indentation under `entries:`.

Fix:

    Remove the malformed `runner.same-state-integrated-dev` row.
    Normalize `runner.registry-lookup` to `  - id: runner.registry-lookup`.
    Reinsert `runner.same-state-integrated-dev` with two-space entry indentation.
    Regenerate `COMMAND_REFERENCE.v0.md`.

Lesson:

    Registry YAML shape is a validation boundary.
    Always inspect surrounding indentation before automated insertion.
    The registry format tests are doing useful work and should stay strict.

### 11.10 v0.6.16 shell spam / interrupted terminal context

The v0.6.16 pasted log contains a large run of repeated `^C` interruptions and shell-context noise between the failed partial implementation and the recovery run.

Interpretation:

    Not an engineering failure by itself.
    The authoritative evidence is the explicit RESULT markers, commit output, post-commit validation, manual push, and final repo snapshot.

Lesson:

    Continue treating shell paste-tail corruption and terminal-context spam as non-authoritative unless it affects actual command output or repo state.

### 11.11 v0.6.16 positive-path proof is not adversarial proof

The new profile is strong for state continuity, but it is still positive-path.

It does not prove:

    malformed Relay payload handling;
    duplicate/replay resistance;
    wrong-group rejection;
    wrong-recipient rejection;
    hostile-server safety;
    metadata privacy.

Lesson:

    Same-state is a necessary integration milestone, not the final security story.

### 11.12 v0.6.17 message-label length failure

The first v0.6.17B implementation generated a message label by concatenating the whole run ID twice:

    same-state-failure-message-same-state-failure-<id>

The OpenMLS sidecar rejected it:

    invalid_message_label: message label is too long

This was not an ack/doctrine failure. It was a runner-generated label hygiene failure.

Recovery:

    separate uniqueID from runID;
    keep runID long for human-readable state names;
    use short messageLabel:
        fail-msg-<id>

Lesson:

    Runner-generated labels must respect sidecar label constraints. Human-readable proof IDs and protocol-facing labels should not automatically share the same length budget.

### 11.13 v0.6.17 wrong-conversation open returned rc=0 but still preserved invariants

The v0.6.17A recon showed that wrong-conversation `message-inbox-dev --ack` did not hard-fail the command, but reported failure in structured output:

    message open failed
    stage: message_open
    acked: false
    opened_envelopes: 0
    open_failures: 1
    ack_failures: 0

Important lesson:

    Do not equate process rc with semantic receive/open success. For current dev wrappers, the critical invariant is no ack/no drain unless message-open succeeds.

The v0.6.17B profile locks the semantic invariant rather than forcing a runtime rc behavior change.

### 11.14 v0.6.17 scripted push auth retry repeated

Scripted push failed with Gitea authentication again.

Manual retry succeeded:

    4f1fec3..03cd00a main -> main

This remains an expected Gitea/server/operator quirk, not a project correctness issue, as long as the final manual snapshot shows `origin/main` aligned.

### 11.15 v0.6.18 clean source implementation still had recurring push auth failure

v0.6.18 source/profile implementation was clean:

    no registry indentation recovery;
    no message-label recovery;
    no runtime surgery;
    no protected profile mutation.

The recurring scripted push auth failure still happened:

    RESULT push rc=128

Manual push succeeded.

Lesson:

    Treat scripted push failure as a known Gitea/server/auth quirk unless final manual snapshot shows origin/main failed to advance. Do not cut a LogDoc until manual push confirms origin/main alignment.

### 11.16 v0.6.18 content_type mutation is a test harness behavior, not a user workflow

The unsupported profile mutates the temporary Cypher DB directly to classify unsupported normal-message behavior.

Lesson:

    This is valid as a live-dev validation harness, not as user-facing behavior. It must not be described as a supported CLI workflow or hostile-server proof. It only proves the current inbox/open/ack path does not ack or drain when the envelope content_type is unsupported.

### 11.17 v0.6.19 profile implementation was clean, but profile-surface sprawl pressure increased

v0.6.19 did not require a source recovery patch.

However, the third narrow same-state failure profile makes an architectural pressure visible:

    same-state-integrated-dev;
    same-state-message-failure-dev;
    same-state-message-unsupported-dev;
    same-state-message-recipient-failure-dev.

This is acceptable for locking invariants now.

Lesson:

    Focused profiles are audit-friendly while hardening a young subsystem.
    After one or two more failure profiles, decide whether to add a failure-suite/meta-profile.
    Do not add an adversarial flag by accident before the adversarial harness contract exists.

### 11.18 v0.6.19 scripted push auth retry repeated

The implementation script again failed at scripted `git push` with Gitea authentication failure:

    RESULT push rc=128

Manual push succeeded:

    1f8086b..197ea9f main -> main

Lesson:

    Treat scripted push failure as a known Gitea/auth quirk.
    Do not consider a breakpoint complete until manual push and final snapshot show `HEAD`, `origin/main`, and `origin/HEAD` aligned.

### 11.19 v0.6.19 pre-test artifact scan showed expected local generated residue

The profile run saw known OpenMLS sidecar generated roots and ignored `cypher.db` during artifact scans.

This was not a validation failure.

Lesson:

    Keep using `--clean-generated`.
    Distinguish known generated roots and ignored local DB residue from tracked-source drift.
    Future package-root validation must remain stricter than live-umbrella dev profile runs.

### 11.20 v0.6.20 recon correctly refused a premature profile

The v0.6.20A rung was initially framed as a likely Welcome/join failure no-ack profile.

The recon found the no-ack/no-drain invariant was good, but also found restored valid join was poisoned by failed corrupt-Welcome state.

Lesson:

    If a failure-path recon exposes partial state poisoning, stop profile-add work and fix state safety first.

Do not encode broken behavior into a profile just because the ack counter stayed correct.

### 11.21 v0.6.20 deep recon localized the fix target correctly

The v0.6.20B deep recon avoided an OpenMLS-version rabbit hole.

Manual quarantine of Bob's final conversation directory recovered restored valid join.

Lesson:

    When final-path state quarantine recovers behavior, the likely fix target is local state materialization/atomicity, not external protocol docs or relay ack ordering.

### 11.22 v0.6.20 Rust warning poisoned sidecar JSON parsing

The first atomic-state patch compiled and passed sidecar tests, but emitted:

    warning: unused variable: `reloaded_member_count`

That warning text appeared before sidecar JSON during runner-driven Go wrapper execution and caused:

    invalid character 'w' looking for beginning of value

Lesson:

    In sidecar subprocess flows where Go parses JSON, warnings are not harmless operationally.
    Sidecar validation should be warning-free before runner/profile validation is trusted.

### 11.23 v0.6.20 scripted push auth quirk repeated, this time in carbonstack-comms

The scripted push failed for `carbonstack-comms` with the recurring auth failure, while `carbonstack` pushed successfully.

Manual push later succeeded and final heads were clean.

Lesson:

    Continue to treat scripted push failure as a known auth quirk only after manual push and final origin/head snapshot confirm it.


### 11.24 v0.6.21 profile assertion overfit to a non-existent output line

Initial v0.6.21B failed because the new profile asserted that `openmls-relay-add-member-dev` output contains literal `content_type`.

The Welcome add-member path does not currently print that line. It prints the better evidence for this profile:

    welcome_envelope_id
    welcome_delivery_state
    welcome_payload_sha256
    welcome_payload_size_bytes
    welcome_acked: false
    group_reloadable: true

Lesson:

    assertions should match the command's actual boundary contract, not a guessed shared output shape from adjacent commands.

The fix was narrow:

    remove the `content_type` assertion only;
    keep actual Welcome evidence assertions;
    rerun profile and regression ladder.

### 11.25 v0.6.21 recovery confirmed that the failure was profile-code-only

The v0.6.21B failure did not require:

    Comms source changes;
    Cypher source changes;
    OS source changes;
    OpenMLS version-documentation recon;
    sidecar atomicity changes;
    registry rollback.

This matters because v0.6.20's sidecar atomic state fix remained valid. The recovery only corrected a CarbonStack runner assertion.

### 11.26 v0.6.21 scripted push auth quirk repeated

The scripted push failed again with:

    RESULT push rc=128

Manual push succeeded afterward:

    e1f6b8c..3876fc8 main -> main

Carry-forward rule:

    scripted push failures are not evidence of implementation failure when all validation and local commit state are clean;
    manual push retry remains acceptable and should be recorded in the LogDoc.

### 11.27 v0.6.25 correctly avoided premature sender-metadata profile

v0.6.25A showed wrong/empty `sender_device_id` opens and acks because current message-open treats sender metadata as relay envelope metadata rather than verified identity.

The correct response was docs/model classification only, not a runner profile.

Why:

    encoding "wrong sender metadata opens and acks" as a permanent passing runner invariant would be premature before identity/trust modeling;
    patching it immediately would also be premature because the system has not defined verified sender identity semantics yet.

### 11.28 v0.6.25 scripted push auth quirk repeated

The v0.6.25 docs commit succeeded locally, but scripted `git push` failed with rc=128.

Manual push succeeded:

    1c421e2..b7f89ea main -> main

This matches the recurring authentication-shell behavior seen in earlier checkpoints.


## 12. Claim Boundaries / Do Not Claim

### 12.0 v0.6.21 additional claim boundaries

The existence of `same-state-welcome-join-failure-dev` means only:

    in a live dev same-state Relay setup,
    a corrupt queued Welcome fails to join,
    does not ack,
    does not drain Bob's Relay inbox,
    leaves no final/staging Bob conversation state,
    and a restored valid Welcome can join and ack.

It does not mean:

    hostile-server safety;
    adversarial relay safety;
    malicious relay resistance;
    identity verification;
    secure enrollment;
    production secure messaging;
    production E2EE;
    metadata privacy;
    stale provider-state safety;
    replay/duplicate/drop/delay safety;
    vault/key-storage safety;
    PQ/hybrid security;
    general-public readiness;
    release-package validation;
    package-root validation.

The profile is a leaf profile. It may later be called by `full-runtime-dev` or another contextual `full-*` evaluator only after explicit aggregation policy exists.

### 12.1 v0.6.20 additional claim boundaries

Do not claim that v0.6.20 provides:

    production secure messaging;
    hostile-server safety;
    metadata privacy;
    identity verification;
    secure enrollment;
    secure vault/key storage;
    PQ or hybrid security;
    adversarial relay harness coverage;
    release-package validation;
    package-root validation;
    full/release-snapshot inclusion;
    mature messenger UX.

Do not claim that v0.6.20 added:

    same-state-welcome-join-failure-dev;
    a registry row;
    a generated command reference entry;
    a public release;
    release assets;
    man pages.

Precise v0.6.20 claim:

    The OpenMLS sidecar join path now has regression-covered atomic staging/promote behavior so corrupt/invalid Welcome join does not leave final Bob conversation state, and restored valid Welcome can join under the same label in protocol test coverage.

Live same-state Relay proof is still a future profile/recon step.

### 12.2 v0.6.19 carried-forward claim boundaries

v0.6.19 allows the narrow statement:

    same-state-message-recipient-failure-dev proves that Alice state + Alice sidecar, Bob state + Alice sidecar, and Bob state + missing sidecar attempts do not falsely open, ack, or drain Bob's queued message after same-state Relay join, and that correct Bob open/ack still succeeds afterward.

Do not inflate this into:

    identity verification;
    malicious relay resistance;
    hostile-server safety;
    metadata privacy;
    production secure messaging;
    production E2EE;
    secure enrollment;
    mature recipient verification;
    mature trust UX.

The profile only proves a deterministic no-false-success/no-ack/no-drain invariant under the dev same-state runner setup.

### 12.3 v0.6.18 carried-forward claim boundaries

v0.6.18 allows the narrow statement:

    same-state-message-unsupported-dev proves unsupported normal application-message content_type no-ack/no-drain behavior.

Do not inflate this into malformed payload coverage, replay/duplicate coverage, adversarial relay safety, or hostile-server safety.

### 12.4 v0.6.17 carried-forward claim boundaries

v0.6.17 allows the narrow statement:

    same-state-message-failure-dev proves wrong-conversation message-open no-ack/no-drain behavior.

Do not inflate this into all message-open failures, all wrong group behavior, or adversarial safety.

### 12.5 v0.6.16 carried-forward claim boundaries

v0.6.16 allows the narrow statement:

    same-state-integrated-dev proves a Level 4 same-conversation positive-path normal-message proof after Relay onboarding in a live-dev temp universe.

Do not inflate this into production secure messaging, full, release-snapshot, release-package validation, package-root validation, hostile-server safety, metadata privacy, identity verification, or secure enrollment.

### 12.6 Carried-forward full/release boundaries

Do not claim:

    full includes same-state runtime profiles;
    release-snapshot includes same-state runtime profiles;
    release package validates same-state profiles;
    v0.6.0 release assets include v0.6.16-v0.6.19 profiles.

Current `full` remains release-package validation.

Current same-state profiles are live-umbrella dev validation.

### 12.7 Carried-forward global claim boundaries

Do not claim:

    production secure messaging;
    production E2EE;
    hostile-server safety;
    metadata privacy;
    verified identity;
    secure enrollment;
    secure vault/key storage;
    PQ/hybrid security;
    Android readiness;
    CarbonStackOS readiness;
    audit/certification;
    local-backbone readiness;
    deployment readiness;
    mature messenger UX;
    general-public usability.

Do not treat:

    KeyPackage/Welcome artifacts as ordinary message inbox items;
    registry presence as promotion;
    generated command reference as man pages;
    same-state positive proof as adversarial proof;
    no-ack/no-drain profiles as hostile-server self-pen-test coverage.

## 13. Lean Breakpoint Summary

Current checkpoint:

    v0.6.25 Normal Message Envelope Metadata Classification Docs Checkpoint

Current commits:

    carbonstack        b7f89ea docs: classify normal message envelope metadata
    carbonstack-comms  b42aa09 fix: make Welcome join state writes atomic
    carbonstack-cypher d18a564 chore: restore validated Go module floor
    carbonstack-os     1bbbe52 docs: clarify CarbonStackOS target direction

New v0.6.25 doc:

    carbonstack/docs/220-v0.6.25-normal-message-envelope-metadata-classification-v0.md

Updated v0.6.25 docs surface:

    carbonstack/docs/README.md

Generated command reference:

    entries=86
    unchanged in v0.6.25

No new v0.6.25 runner profile or registry ID was added.

New classification locked:

    payload_sha256 mismatch
    -> detected/rejected before open/ack/drain
    -> recovery after restore OK

    payload_size_bytes mismatch
    -> detected/rejected before open/ack/drain
    -> recovery after restore OK

    invalid protocol_version
    -> unsupported/skipped
    -> no open, no ack, no drain
    -> recovery after restore OK

    wrong recipient_device_id
    -> Bob fetch suppressed by routing metadata
    -> recovery after restore OK

    invalid delivery_state
    -> fetch suppressed by delivery-state bookkeeping
    -> recovery after restore OK

    wrong or empty sender_device_id
    -> message opens and acks
    -> from_device reflects mutated relay envelope metadata
    -> classified as unverified relay metadata, not verified identity

Existing active profiles retained:

    same-state-integrated-dev
    same-state-message-failure-dev
    same-state-message-unsupported-dev
    same-state-message-malformed-payload-dev
    same-state-message-replay-classification-dev
    same-state-message-recipient-failure-dev
    same-state-welcome-join-failure-dev

Protected boundaries retained:

    full unchanged
    release-snapshot unchanged
    no public release rebuild
    no replay-safety claim
    no identity verification claim
    no hostile-server safety claim
    no adversarial harness claim
    no vault/PQ claim

Next safest action:

    adversarial QA for evergreen roadmap refresh.


## 14. v0.6.22 Continuity Overlay

This overlay is intentionally additive. It preserves the full v0.6.21 ledger above and records the v0.6.22 state-authority checkpoint without compressing away prior history.

### 14.1 Current v0.6.22 state

Current mainline checkpoint:

    v0.6.23 Same-State Malformed Normal-Message Payload Profile

Current pushed heads:

    carbonstack        943e0e4 test: add same-state malformed payload profile
    carbonstack-comms  b42aa09 fix: make Welcome join state writes atomic
    carbonstack-cypher d18a564 chore: restore validated Go module floor
    carbonstack-os     1bbbe52 docs: clarify CarbonStackOS target direction

Current public release remains:

    v0.6.0 State/UX Boundary Validation Pre-Release

Working-tree note:

    carbonstack, carbonstack-comms, and carbonstack-os were clean/aligned with origin/main after manual push.
    carbonstack-cypher retained the known ignored local artifact:
        !! cypher.db

### 14.2 v0.6.22 doc/model additions

New doc:

    carbonstack/docs/217-v0.6.22-sidecar-state-authority-classification-v0.md

Docs index updated:

    carbonstack/docs/README.md

No new runner profile.

No new registry ID.

No command-reference count change.

No Comms/Cypher/OS source change.

### 14.3 State authority model after v0.6.22

Runtime authority for existing conversation message-open:

    conversation-level provider-storage.json
    final conversation directory

Metadata/evidence or non-runtime existing-message-open substrate:

    conversation-summary.json
    join-summary.json
    device-level provider-storage.json

Detailed classification:

    conversation provider-storage missing/invalid/truncated:
        blocks message-open;
        no ack;
        no Bob inbox drain;
        runtime authority.

    final conversation directory moved/missing:
        blocks message-open;
        no ack;
        no Bob inbox drain;
        runtime authority container.

    conversation-summary.json missing:
        conversation-load-check-dev fails;
        message-inbox-dev --ack can still open/ack if provider storage remains valid;
        classification: load-check stricter than message-open.

    conversation-summary.json invalid or valid-but-wrong:
        message-open/ack succeeds if provider storage remains valid;
        classification: metadata/evidence, not existing-message-open authority.

    join-summary.json missing/invalid/valid-but-wrong:
        message-open/ack succeeds if provider storage remains valid;
        classification: metadata/evidence.

    device-level provider-storage.json missing/invalid:
        existing conversation message-open/ack succeeds if conversation provider storage remains valid;
        classification: identity/bootstrap substrate, not existing-conversation runtime message-open authority.

### 14.4 Critical implication for future code

Do not blindly patch `message-inbox-dev` to block on summary metadata.

Patch later only if future recon finds:

    runtime provider storage bypass;
    final conversation directory bypass;
    silent regeneration/replacement of state;
    ack on failed open;
    inbox drain on failed open;
    false message-open success.

For current v0.6.22 semantics, likely cleanup is one of:

    clarify/rename conversation-load-check-dev;
    split provider reloadability from metadata presence/integrity checks;
    add warning output when metadata is missing/malformed but provider storage remains usable;
    add a dedicated metadata-state-check command.

### 14.5 v0.6.22 blunders / lessons

1. **Long label failure repeated in recon form**
   - Initial stale-state recon failed before mutation semantics because message labels included long case names.
   - Lesson:
       probe labels must separate human-readable case names from sidecar/CLI-bounded machine labels.

2. **Conservative probe output looked scarier than the model warranted**
   - v0.6.22A reported cases with open/ack/drain after summary/device-provider mutation.
   - That was useful, but not sufficient for a bug call.
   - v0.6.22B correctly reframed the result as an authority-classification question.

3. **Tree/hash change heuristics can overstate silent regeneration**
   - The first recon's `tree_changed` field included normal message-open side effects such as message-open summary output.
   - v0.6.22B narrowed the measurement to target-state changes and reported no target rewrite/regeneration.

4. **Scripted push auth quirk repeated**
   - Scripted push failed with rc=128.
   - Manual push succeeded.
   - Continue treating final manual snapshot as authoritative when script push fails.

### 14.6 Updated near-term path

Immediate next recommended rung:

    malformed normal-message payload recon

Expected invariant:

    corrupt encrypted application-message payload
    -> no false message-open
    -> no ack
    -> no Bob inbox drain
    -> clear open-failure reporting
    -> recovery behavior defined if original payload is restored

Alternative short cleanup before malformed payload, if deliberately chosen:

    normalize conversation-load-check output semantics around provider reloadability vs metadata presence.

The next planned sequence remains:

    malformed normal-message payload profile complete
    -> replay/duplicate classification
    -> longer adversarial QA
    -> roadmap refresh
    -> remaining v0.6.x rungs
    -> v0.7.0 minor-epoch release definition
    -> expanded v0.7.x+ evergreen goals

### 14.7 Updated nonclaims after v0.6.22

Do not claim:

    stale-state model closure across all state files/failure modes;
    runtime patch for stale state;
    committed stale-state runner profile;
    adversarial relay harness;
    hostile-server safety;
    metadata privacy;
    production secure messaging;
    production E2EE;
    verified identity;
    secure enrollment;
    secure vault/key storage;
    PQ or hybrid security;
    Android/CarbonStackOS readiness;
    release-package inclusion;
    full or release-snapshot coverage.

## 15. v0.6.23 Malformed Normal-Message Payload Profile Checkpoint

### 15.1 Checkpoint summary

v0.6.23 promoted malformed normal-message payload behavior from recon evidence into a committed live same-state runner profile.

Added profile:

    same-state-message-malformed-payload-dev

Added registry ID:

    runner.same-state-message-malformed-payload-dev

New CarbonStack commit:

    943e0e4 test: add same-state malformed payload profile

Command reference state:

    entries=85

No Comms, Cypher, or OS source changed.

### 15.2 v0.6.23A recon result

v0.6.23A tested normal application-message payload corruption after a valid same-state Relay join and `message-send-dev`.

The tested payload mutation families were:

    p01 invalid ciphertext_b64 storage shape
    p02 valid base64 random bytes
    p03 valid base64 truncated original OpenMLS message bytes
    p04 valid base64 single-byte-flipped original OpenMLS message bytes
    p05 valid base64 empty bytes
    p06 valid base64 original bytes with appended junk

The recon summary was:

    cases_total: 6
    classification_count clear_failure_no_ack_no_drain_recovery_ok: 6
    cases_with_open_ack_or_drain: 0
    cases_with_provider_mutation_on_failed_open: 0
    cases_with_envelope_mutation_on_failed_open: 0
    cases_with_recovery_needing_review: 0

Every malformed case produced one open failure, did not ack, did not drain, did not mutate provider state, did not rewrite the envelope, and recovered after restoring the original payload.

### 15.3 v0.6.23B profile implementation

The committed profile proves the same core invariant in runner form:

    malformed normal application-message payloads do not falsely open;
    malformed normal application-message payloads do not ack;
    malformed normal application-message payloads do not drain Bob's queued normal-message inbox;
    malformed normal application-message payloads do not mutate conversation provider storage;
    malformed normal application-message payloads do not rewrite the envelope;
    restored original payload recovery opens and acks.

Files added or updated:

    tools/carbonstack-validate/same_state_message_malformed_payload_dev.go
    tools/carbonstack-validate/main.go
    tools/carbonstack-validate/README.md
    registry/commands.v0.yaml
    registry/COMMAND_REFERENCE.v0.md
    docs/218-v0.6.23-same-state-message-malformed-payload-dev-profile-v0.md
    docs/README.md

The profile is a live-umbrella dev/pre-alpha proof. It is not `full`, not `release-snapshot`, not release-package validation, and not an adversarial harness.

### 15.4 v0.6.23 implementation blunder and recovery

Initial v0.6.23B implementation failed inside the new profile after it had already performed Relay onboarding and normal `message-send-dev`.

Failure:

    VALIDATION FAILED: could not locate Bob application-message envelope

Root cause:

    the profile attempted to locate the application-message envelope with:
        recipient_device_id = Bob
        relay_space_id = sub.RelaySpace
        content_type = carbonstack.mls.application-message.v0

But normal `message-send-dev` application-message envelopes are not guaranteed to carry the same Relay Space column shape as onboarding artifacts.

Recovery:

    remove the relay_space_id predicate;
    locate Bob's queued application-message envelope by:
        recipient_device_id = Bob
        content_type = carbonstack.mls.application-message.v0
        delivery_state = queued

This matched the v0.6.23A recon lookup shape and preserved the intended invariant.

Classification:

    runner/profile lookup bug;
    not a malformed-payload safety failure;
    not a Comms runtime bug;
    not a Cypher runtime bug;
    not an OpenMLS sidecar bug.

### 15.5 v0.6.23 validation ladder

Validation passed after recovery:

    generated reference deterministic check;
    runner go test ./...;
    registry missing-nonclaims scan;
    registry lookup for runner.same-state-message-malformed-payload-dev;
    runner doctor;
    same-state-message-malformed-payload-dev --clean-generated;
    generated reference deterministic check post-profile;
    runner go test ./... post-profile;
    same-state-integrated-dev;
    same-state-message-failure-dev;
    same-state-message-unsupported-dev;
    same-state-message-recipient-failure-dev;
    same-state-welcome-join-failure-dev;
    registry missing-nonclaims scan post-profile;
    cached diff --check;
    post-commit generated reference check;
    post-commit runner go test ./...;
    post-commit registry lookup;
    post-commit same-state-message-malformed-payload-dev.

The new profile's own result included:

    PASS: malformed payload cases tested: 6
    PASS: invalid base64 storage shape did not open, ack, drain, rewrite envelope, or mutate provider state
    PASS: valid base64 random bytes did not open, ack, drain, rewrite envelope, or mutate provider state
    PASS: truncated original payload did not open, ack, drain, rewrite envelope, or mutate provider state
    PASS: single-byte-flipped original payload did not open, ack, drain, rewrite envelope, or mutate provider state
    PASS: empty payload did not open, ack, drain, rewrite envelope, or mutate provider state
    PASS: original payload plus junk did not open, ack, drain, rewrite envelope, or mutate provider state
    PASS: restored original payload opened and acked after every malformed failure

### 15.6 Push state

Scripted push failed with the known auth quirk:

    RESULT push rc=128

Manual push succeeded:

    8dd7db5..943e0e4 main -> main

Final pushed heads:

    carbonstack        943e0e4 test: add same-state malformed payload profile
    carbonstack-comms  b42aa09 fix: make Welcome join state writes atomic
    carbonstack-cypher d18a564 chore: restore validated Go module floor
    carbonstack-os     1bbbe52 docs: clarify CarbonStackOS target direction

`carbonstack-cypher/cypher.db` remains a local ignored artifact noted by status output.

### 15.7 Critical path/function updates

New runner profile:

    carbonstack/tools/carbonstack-validate/same_state_message_malformed_payload_dev.go

New registry entry:

    runner.same-state-message-malformed-payload-dev

New docs:

    carbonstack/docs/218-v0.6.23-same-state-message-malformed-payload-dev-profile-v0.md

Generated command reference:

    carbonstack/registry/COMMAND_REFERENCE.v0.md
    entries=85

Important distinction:

    same-state-message-unsupported-dev:
        mutates normal application-message content_type.

    same-state-message-malformed-payload-dev:
        keeps normal application-message content_type valid and mutates payload/ciphertext storage.

    same-state-message-failure-dev:
        wrong-conversation failure.

    same-state-message-recipient-failure-dev:
        wrong recipient/device/sidecar failure.

    same-state-welcome-join-failure-dev:
        corrupt Welcome join failure and recovery.

### 15.8 Near-future work

Recommended next rung:

    duplicate/replay classification recon

Initial replay/duplicate classification should not claim full replay safety. It should classify behavior using a deliberately boring vocabulary:

    detected
    rejected
    tolerated
    unknown
    deferred

Likely recon surfaces:

    duplicate queued application-message envelope with same payload;
    duplicate envelope ID insertion if DB constraints permit or reject it;
    duplicate payload under new envelope ID;
    replay after successful ack;
    replay before ack;
    stale replay after provider state advanced;
    Welcome replay / duplicate join may be separate from normal-message replay.

After duplicate/replay classification:

    longer adversarial QA;
    roadmap refresh;
    remaining v0.6.x decisions;
    v0.7.0 minor-epoch release definition;
    expanded v0.7.x+ evergreen goals.

### 15.9 Updated nonclaims after v0.6.23

Do not claim:

    public release readiness;
    full validation;
    release-snapshot validation;
    release-package validation;
    package-root validation;
    adversarial relay harness coverage;
    hostile-server safety;
    metadata privacy;
    production secure messaging;
    production E2EE;
    verified identity;
    secure enrollment;
    replay or duplicate protection;
    malformed payload model closure across all possible parser/state cases;
    secure vault/key storage;
    PQ or hybrid security;
    Android/CarbonStackOS readiness;
    local-backbone/deployment readiness;
    mature messenger UX;
    general-public usability.

## 16. v0.6.24 Load-Check Semantics and Normal Application-Message Replay Classification Checkpoint

### 16.1 Checkpoint summary

v0.6.24 promoted normal application-message duplicate/replay behavior from recon evidence into a committed live same-state runner profile, while also recording the load-check/provider-vs-summary semantic mismatch as docs/model state rather than a runtime patch.

Added profile:

    same-state-message-replay-classification-dev

Added registry ID:

    runner.same-state-message-replay-classification-dev

New CarbonStack commit:

    1c421e2 test: add same-state replay classification profile

Command reference state:

    entries=86

No Comms, Cypher, or OS source changed.

### 16.2 v0.6.24A load-check semantic recon result

v0.6.24A reproduced the known v0.6.22 semantics around missing `conversation-summary.json`:

    LOADCHECK_SEMANTIC_RECON_RESULT
    cases_total: 1
    l01 missing-summary-loadcheck-vs-message-open:
        load_rc=1
        load_loaded=no
        opened=yes
        acked=yes
        acks=0->1
        inbox=1->0
        classification=load_check_stricter_than_message_open_summary_metadata_missing

Meaning:

    conversation-load-check-dev is stricter than message-open when summary metadata is missing.
    message-inbox-dev can still open/ack a valid queued message when provider-storage.json remains valid.
    This remains a command-semantics/documentation issue, not a no-ack/no-drain runtime bug.

Do not make `conversation-summary.json` runtime authority by accident. v0.6.22 classified it as metadata/evidence for existing message-open.

### 16.3 v0.6.24A normal application-message replay recon result

v0.6.24A tested five normal application-message duplicate/replay cases:

    r01 duplicate same envelope_id insert attempt
    r02 same envelope after ack without requeue
    r03 same envelope manually requeued after ack
    r04 duplicate same ciphertext under new envelope_id before original ack
    r05 duplicate same ciphertext under new envelope_id after original ack

Recon summary:

    REPLAY_CLASSIFICATION_RECON_RESULT
    cases_total: 5
    classification_count delivery_state_suppressed_same_envelope_after_ack: 1
    classification_count detected_or_rejected_no_ack_no_drain_duplicate_payload_new_envelope_after_ack: 1
    classification_count detected_or_rejected_no_ack_no_drain_duplicate_payload_new_envelope_before_ack: 1
    classification_count detected_or_rejected_no_ack_no_drain_manual_requeue_same_envelope: 1
    classification_count storage_rejected_duplicate_same_envelope_id_original_recovery_ok: 1
    unsafe_cases: 0
    bookkeeping_concern_cases: 0
    tolerated_or_unknown_cases: 0
    detected_or_rejected_cases: 3
    storage_rejected_cases: 1
    delivery_state_suppressed_cases: 1

Important caveat:

    r01 provider_changed_by_replay=yes was normal provider-state evolution from opening the original valid message after duplicate same-ID insert was rejected.
    It was not a replay-bypass signal.

### 16.4 v0.6.24B profile implementation result

v0.6.24B added the live runner profile:

    same-state-message-replay-classification-dev

The profile proves the narrow classifications:

    duplicate same envelope_id insert is storage rejected;
    same envelope after ack is delivery-state suppressed;
    manually requeued same envelope after ack fails no-open/no-ack/no-drain;
    duplicate same ciphertext under new envelope ID before original ack fails no-open/no-ack/no-drain;
    duplicate same ciphertext under new envelope ID after original ack fails no-open/no-ack/no-drain.

The profile intentionally says classification, not replay-safety.

### 16.5 Validation ladder

Observed validation after recovery included:

    generated reference deterministic check: entries=86
    runner go test ./...: passed
    registry missing nonclaims scan: passed
    registry lookup runner.same-state-message-replay-classification-dev: passed
    runner doctor: passed
    sidecar cargo test warning-free: passed
    comms go test ./...: passed
    same-state-message-replay-classification-dev --clean-generated: passed

Regression ladder preserved existing profiles:

    same-state-integrated-dev still passes
    same-state-message-failure-dev still passes
    same-state-message-unsupported-dev still passes
    same-state-message-malformed-payload-dev still passes
    same-state-message-recipient-failure-dev still passes
    same-state-welcome-join-failure-dev still passes

Post-commit validation passed:

    generated reference deterministic check post-commit: entries=86
    runner go test ./... post-commit: passed
    registry lookup replay classification post-commit: passed
    same-state-message-replay-classification-dev post-commit: passed

### 16.6 Blunder and recovery

Initial v0.6.24B failed during `runner go test ./...` with:

    undefined: runCommandCombined

Cause:

    same_state_message_replay_classification_dev.go referenced a helper that did not exist in the runner package.

Recovery:

    add os/exec import;
    add local replayRunPythonScript helper;
    replace both bad runCommandCombined calls;
    rerun validation/profile/commit flow.

Meaning:

    runner implementation/porting bug only;
    not replay runtime failure;
    not Comms/Cypher/OpenMLS failure.

### 16.7 Push state

Scripted push succeeded:

    943e0e4..1c421e2 main -> main
    RESULT push rc=0

Final pushed heads:

    carbonstack        1c421e2 test: add same-state replay classification profile
    carbonstack-comms  b42aa09 fix: make Welcome join state writes atomic
    carbonstack-cypher d18a564 chore: restore validated Go module floor
    carbonstack-os     1bbbe52 docs: clarify CarbonStackOS target direction

`carbonstack-cypher/cypher.db` remains a local ignored artifact noted by status output.

### 16.8 Critical path/function updates

New runner profile:

    carbonstack/tools/carbonstack-validate/same_state_message_replay_classification_dev.go

New registry entry:

    runner.same-state-message-replay-classification-dev

New docs:

    carbonstack/docs/219-v0.6.24-loadcheck-replay-classification-v0.md

Generated command reference:

    carbonstack/registry/COMMAND_REFERENCE.v0.md
    entries=86

Important profile distinction:

    same-state-message-malformed-payload-dev:
        malformed payload/ciphertext storage no-open/no-ack/no-drain proof.

    same-state-message-replay-classification-dev:
        normal application-message duplicate/replay classification proof.

    same-state-message-unsupported-dev:
        unsupported content-type no-ack/no-drain proof.

    same-state-message-failure-dev:
        wrong-conversation failure.

    same-state-message-recipient-failure-dev:
        wrong recipient/device/sidecar failure.

    same-state-welcome-join-failure-dev:
        corrupt Welcome join failure and recovery.

### 16.9 Near-future work

Immediate next recommended rung:

    remaining malformed typed cases

This should cover remaining typed malformed cases not already covered by:

    unsupported content_type;
    malformed payload/ciphertext storage;
    wrong conversation;
    wrong recipient/device/sidecar;
    duplicate/replay classification.

Candidate remaining malformed classes to consider:

    envelope metadata mismatch;
    hash/size lying or missing metadata;
    missing/empty ciphertext field shape beyond the current payload cases;
    malformed typed onboarding artifacts if deliberately scoped later;
    malformed Welcome/KeyPackage replay should remain separate from normal application-message replay.

Corrected larger order agreed by operator:

    replay classification
    -> remaining malformed typed cases
    -> adversarial QA for evergreen roadmap refresh
    -> vault/security model + no-silent/backup/recovery rules
    -> bounded vault substrate/stub
    -> PQ/hybrid placement model
    -> secured vault/PQ implementation planning
    -> adversarial harness

### 16.10 Updated nonclaims after v0.6.24

Do not claim:

    public release readiness;
    full validation;
    release-snapshot validation;
    release-package validation;
    package-root validation;
    replay safety;
    adversarial relay harness coverage;
    hostile-server safety;
    server equivocation detection;
    Welcome replay handling;
    KeyPackage replay handling;
    network drop/delay handling;
    metadata privacy;
    production secure messaging;
    production E2EE;
    verified identity;
    secure enrollment;
    secure vault/key storage;
    PQ or hybrid security;
    Android/CarbonStackOS readiness;
    local-backbone/deployment readiness;
    mature messenger UX;
    general-public usability.

## 17. v0.6.25 Normal Message Envelope Metadata Classification Checkpoint

### 17.1 Checkpoint summary

v0.6.25 is a docs/model checkpoint, not a source/runtime/profile checkpoint.

It records how the current dev/pre-alpha normal message-open path treats selected envelope metadata fields:

    payload_sha256;
    payload_size_bytes;
    protocol_version;
    sender_device_id;
    recipient_device_id;
    delivery_state.

The checkpoint deliberately does not add a runner profile because two sender metadata cases open and ack under current behavior. That behavior is important to document, but it should not be promoted into a stable profile invariant before identity/trust modeling.

### 17.2 v0.6.25A recon result

v0.6.25A tested seven normal application-message envelope metadata mutations.

Final result:

    cases_total: 7
    classification_count detected_or_rejected_no_ack_no_drain_recovery_ok: 3
    classification_count metadata_not_runtime_enforced_open_ack_drain: 2
    classification_count metadata_routing_suppressed_no_open_no_ack_recovery_ok: 1
    classification_count delivery_state_suppressed_no_open_no_ack_recovery_ok: 1
    opened_or_acked_cases: 2
    suppressed_or_rejected_cases: 5
    bookkeeping_concern_cases: 0
    unsafe_or_unexpected_state_mutation_cases: 0
    provider_changed_by_attempt_cases: 2

Case result summary:

    m01 payload_sha256_mismatch:
      no open;
      no ack;
      no drain;
      provider unchanged;
      restore/recovery OK.

    m02 payload_size_bytes_mismatch:
      no open;
      no ack;
      no drain;
      provider unchanged;
      restore/recovery OK.

    m03 protocol_version_invalid:
      skipped/unsupported;
      no open;
      no ack;
      no drain;
      provider unchanged;
      restore/recovery OK.

    m04 sender_device_id_wrong:
      opened;
      acked;
      from_device reflected mutated sender metadata;
      provider changed as part of normal successful OpenMLS open.

    m05 sender_device_id_empty:
      opened;
      acked;
      from_device was empty;
      provider changed as part of normal successful OpenMLS open.

    m06 recipient_device_id_wrong:
      Bob inbox query saw zero queued envelopes;
      no open;
      no ack;
      provider unchanged;
      restore/recovery OK.

    m07 delivery_state_invalid:
      Bob inbox query saw zero queued envelopes;
      no open;
      no ack;
      provider unchanged;
      restore/recovery OK.

### 17.3 v0.6.25B docs/model implementation result

v0.6.25B added:

    docs/220-v0.6.25-normal-message-envelope-metadata-classification-v0.md

and updated:

    docs/README.md

It did not change:

    carbonstack-comms source;
    carbonstack-cypher source;
    carbonstack-os source;
    runner profiles;
    registry/commands.v0.yaml;
    registry/COMMAND_REFERENCE.v0.md;
    full;
    release-snapshot;
    release assets.

Commit:

    b7f89ea docs: classify normal message envelope metadata

### 17.4 Validation ladder

v0.6.25B validation passed:

    carbonstack diff --check
    generated reference deterministic check
    runner go test ./...
    registry missing nonclaims scan
    runner doctor
    same-state-message-replay-classification-dev
    same-state-message-malformed-payload-dev
    no registry/reference/runner/source patch guard
    cached diff --check
    post-commit generated reference deterministic check
    post-commit runner go test ./...
    post-commit same-state-message-replay-classification-dev

The generated command reference remained current at:

    entries=86

### 17.5 Push state

Scripted push failed with the recurring auth quirk:

    RESULT push rc=128

Manual push succeeded:

    1c421e2..b7f89ea main -> main

Final pushed heads:

    carbonstack        b7f89ea docs: classify normal message envelope metadata
    carbonstack-comms  b42aa09 fix: make Welcome join state writes atomic
    carbonstack-cypher d18a564 chore: restore validated Go module floor
    carbonstack-os     1bbbe52 docs: clarify CarbonStackOS target direction

`carbonstack-cypher` still retained the known ignored local artifact:

    !! cypher.db

### 17.6 Critical path/function updates

New v0.6.25 doc path:

    carbonstack/docs/220-v0.6.25-normal-message-envelope-metadata-classification-v0.md

Updated docs index:

    carbonstack/docs/README.md

Relevant existing function/surface paths:

    carbonstack-comms/internal/app/message_wrappers_dev.go
    carbonstack-comms/internal/app/openmls_runtime.go
    carbonstack-cypher/internal/httpapi/api.go
    carbonstack-cypher/migrations/001_init.sql
    carbonstack-cypher/migrations/002_envelope_payload_metadata.sql

Relevant live profiles retained:

    carbonstack/tools/carbonstack-validate/same_state_message_replay_classification_dev.go
    carbonstack/tools/carbonstack-validate/same_state_message_malformed_payload_dev.go
    carbonstack/tools/carbonstack-validate/same_state_message_recipient_failure_dev.go
    carbonstack/tools/carbonstack-validate/same_state_welcome_join_failure_dev.go

### 17.7 Critical model update

Current normal-message envelope metadata categories after v0.6.25:

Runtime-enforced before successful message-open:

    payload_sha256;
    payload_size_bytes;
    protocol_version support gate.

Fetch/routing/bookkeeping selectors:

    recipient_device_id;
    delivery_state.

Unverified relay metadata:

    sender_device_id;
    from_device output derived from sender_device_id.

Hard warning:

    from_device in current message-inbox-dev output is relay envelope sender metadata.
    It must not be presented as verified identity.

### 17.8 Near-future work

Immediate next safest action:

    adversarial QA for evergreen roadmap refresh.

Questions to resolve there:

    Which current v0.6.x hardening goals are complete vs still active?
    Should load-check/provider-vs-summary cleanup happen before v0.7.0?
    Should from_device output be renamed or warning-hardened before identity/trust modeling?
    Which Welcome/KeyPackage malformed/replay cases are v0.6.x, v0.7.x, or adversarial-harness work?
    What exact vault/security model is required before a bounded vault substrate/stub?
    What exact v0.7.0 release claim boundary is honest?

### 17.9 Updated nonclaims after v0.6.25

v0.6.25 does not claim:

    public release;
    full;
    release-snapshot;
    release-package validation;
    package-root validation;
    identity verification;
    sender authenticity;
    metadata authenticity;
    hostile-server safety;
    adversarial relay harness coverage;
    replay safety;
    server equivocation detection;
    Welcome replay handling;
    KeyPackage replay handling;
    metadata privacy;
    production secure messaging;
    production E2EE;
    secure enrollment;
    vault/key-storage safety;
    PQ or hybrid security;
    Android/CarbonStackOS readiness;
    audit/certification;
    local-backbone readiness;
    deployment readiness;
    general-public usability.

## 18. v0.6.26 Roadmap and Output Semantics Cleanup Checkpoint

### 18.1 Checkpoint summary

v0.6.26 closed the post-v0.6.25 adversarial Q/A preflight and immediate output-semantics cleanup lane.

The checkpoint is intentionally narrow:

    roadmap refresh;
    sender metadata output warning cleanup;
    load-check provider/summary output cleanup.

It is not:

    a public release;
    full;
    release-snapshot;
    package-root validation;
    vault/security modeling;
    PQ/hybrid implementation;
    adversarial harness implementation;
    identity verification;
    metadata authenticity;
    hostile-server safety.

### 18.2 v0.6.26 roadmap authority update

The active forward roadmap is now:

    CarbonStack_Long_Term_Roadmap_v0.6.26_EVERGREEN

The previous v0.6.17 roadmap remains historical context, but v0.6.26 supersedes it for forward planning.

Current accepted lane map:

    v0.6.x remaining:
        load-check and sender metadata cleanup complete;
        vault/security model;
        no-silent/backup/recovery rules;
        bounded vault substrate/stub decision;
        PQ/hybrid placement model;
        secured vault/PQ implementation planning boundaries;
        adversarial harness contract/evidence matrix design;
        v0.7.0 release boundary and package-root rehearsal plan.

    v0.7.x:
        deployability + state-security integration;
        CarbonStack Addressable Relay Space mechanics;
        rudimentary CLI dev app for joining/using Relay Spaces;
        full-validate-release / full-runtime-dev split;
        dev CLI/client lifecycle;
        local state/vault substrate integration;
        preferred + generic deployment modeling;
        PQ implementation pilots only in isolated branch/workdir and only if v0.6 model is complete.

    v0.8.x:
        PQ/deployability maturation;
        migration/compatibility tests;
        optimization/code-health hardening;
        package-root runtime validation candidates;
        operator docs;
        release-candidate hardening.

    v0.9.x:
        full self-pen-test / adversarial validation across real operational surfaces.

### 18.3 v0.6.26B output-semantics scout result

The scout produced the decision to split output cleanup into two implementation sub-rungs:

    v0.6.26C sender metadata / from_device unverified output cleanup;
    v0.6.26D load-check provider-vs-summary output cleanup.

Scout facts:

    sender_grep_lines: 723
    sender_grep_from_device_hits: 24
    loadcheck_grep_lines: 1801
    loadcheck_grep_conversation_summary_hits: 105
    loadcheck_grep_group_reloadable_hits: 200
    tests_grep_lines: 223
    live probe saw from_device but no sender_identity_verified/from_device_unverified fields.

This justified doing sender metadata first and load-check second.

### 18.4 v0.6.26C sender metadata output cleanup

Files changed in `carbonstack-comms`:

    internal/app/message_wrappers_dev.go
    internal/app/openmls_runtime.go
    internal/app/message_wrappers_dev_test.go

Files changed in `carbonstack`:

    docs/221-v0.6.26-sender-metadata-output-warning-v0.md
    docs/README.md

Output contract after v0.6.26C:

    from_device: <legacy sender_device_id>
    from_device_unverified: <same value>
    sender_identity_verified: false
    warning: from_device is relay envelope metadata, not verified identity

Important interpretation:

    from_device remains for compatibility.
    from_device_unverified is the clearer contract.
    sender_identity_verified is always false in this current path.
    This does not implement identity verification.

### 18.5 v0.6.26D load-check provider/summary output cleanup

Files changed in `carbonstack-comms`:

    internal/app/openmls_bootstrap.go
    internal/app/openmls_bootstrap_dev_test.go

Files changed in `carbonstack`:

    docs/222-v0.6.26-loadcheck-provider-summary-output-v0.md
    docs/README.md

Successful load-check output now includes:

    provider_reloadable: true
    summary_metadata_present: true
    provider_storage_present: true
    summary_metadata_path: <path>
    provider_storage_path: <path>
    summary_metadata_warning: none

Missing summary metadata output now includes:

    status: metadata_missing
    group_reloadable: false
    provider_reloadable: not_evaluated_by_load_check
    summary_metadata_present: false
    provider_storage_present: true/false
    summary_metadata_path: <path>
    provider_storage_path: <path>
    summary_metadata_warning: conversation-summary metadata is missing; conversation-load-check-dev is stricter than message-open and cannot confirm reloadability without summary metadata

Important interpretation:

    provider_reloadable means the load-check path successfully asked the sidecar to reload the conversation using current load-check semantics.
    summary_metadata_present means expected sidecar summary metadata is present for the load-check surface.
    provider_storage_present means expected provider storage is present for the load-check surface.
    provider_reloadable: not_evaluated_by_load_check means the command did not prove reloadability.
    This does not alter message-inbox-dev behavior.

### 18.6 v0.6.26 blunders / lessons

1. **False positive success-path patch detection**
   - The first v0.6.26D script searched the entire function for `provider_reloadable`.
   - Because the error path already contained that string, it skipped success-path insertion.
   - The live probe caught the missing success-path fields before commit.

2. **Literal newline escape output**
   - The first recovery emitted literal `\n` sequences inside output strings.
   - This collapsed multiple intended output fields onto one line.
   - Lesson: when patching Go `fmt.Printf` strings through Python, inspect actual rendered output, not just diff structure.

3. **Overpatched unit tests**
   - The first recovery patched tests for add-member and join to expect load-check-only fields.
   - That was semantically wrong because `provider_reloadable` belongs to `openmls-conversation-load-check-dev`, not all commands that print `group_reloadable`.
   - The clean recovery reset and patched only the load-check test.

4. **Unit test vs live filesystem mismatch**
   - The first recovery derived metadata presence from the filesystem during a mocked unit test.
   - The mocked sidecar reported reloadable state without real summary/provider files, causing false `summary_metadata_present: false` in test output.
   - The clean patch treats successful sidecar load-check as summary/provider presence for the load-check output contract and uses filesystem probes only for the error/missing-summary branch.

5. **Push/auth behavior remains a known operational nuisance**
   - v0.6.26C Comms push initially failed with the recurring rc=128 auth quirk.
   - v0.6.26D final push succeeded and carried both Comms commits to origin.

### 18.7 Critical path/function updates

New or updated docs:

    carbonstack/docs/221-v0.6.26-sender-metadata-output-warning-v0.md
    carbonstack/docs/222-v0.6.26-loadcheck-provider-summary-output-v0.md
    carbonstack/docs/README.md
    CarbonStack_Long_Term_Roadmap_v0.6.26_EVERGREEN.pdf

Updated Comms output functions:

    carbonstack-comms/internal/app/message_wrappers_dev.go
        message-inbox-dev output now labels from_device as unverified relay envelope metadata.

    carbonstack-comms/internal/app/openmls_runtime.go
        openmls-inbox-dev output now labels from_device as unverified relay envelope metadata.

    carbonstack-comms/internal/app/openmls_bootstrap.go
        openmls-conversation-load-check-dev output now separates provider reloadability and summary metadata presence.

Updated Comms tests:

    carbonstack-comms/internal/app/message_wrappers_dev_test.go
    carbonstack-comms/internal/app/openmls_bootstrap_dev_test.go

Reference paths retained:

    carbonstack/docs/220-v0.6.25-normal-message-envelope-metadata-classification-v0.md
    carbonstack/docs/219-v0.6.24-loadcheck-replay-classification-v0.md
    carbonstack/docs/217-v0.6.22-sidecar-state-authority-classification-v0.md
    carbonstack/tools/carbonstack-validate
    carbonstack/registry/commands.v0.yaml
    carbonstack/registry/COMMAND_REFERENCE.v0.md

### 18.8 Near-future work after v0.6.26

Immediate next safe mainline work:

    vault/security model;
    no-silent regeneration/replacement/import/delete/migration/restore rules;
    backup/recovery model.

Then:

    bounded vault substrate/stub decision and optional tightly scoped stub;
    PQ/hybrid placement model;
    secured vault/PQ implementation planning boundaries;
    adversarial harness contract/evidence matrix;
    v0.7.0 release boundary and package-root rehearsal plan.

Do not jump directly to implementation-heavy vault/PQ/adversarial code before the model.

### 18.9 Updated nonclaims after v0.6.26

v0.6.26 does not claim:

    public release;
    full;
    release-snapshot;
    release-package validation;
    package-root validation;
    production secure messaging;
    production E2EE;
    hostile-server safety;
    adversarial relay harness coverage;
    replay safety;
    server equivocation detection;
    Welcome replay handling;
    KeyPackage replay handling;
    metadata privacy;
    metadata authenticity;
    sender authenticity;
    identity verification;
    secure enrollment;
    vault/key-storage safety;
    PQ or hybrid security;
    Android/CarbonStackOS readiness;
    deployability readiness;
    mature user-facing client UX.

The new output fields are clarity improvements only. They do not expand security claims.
