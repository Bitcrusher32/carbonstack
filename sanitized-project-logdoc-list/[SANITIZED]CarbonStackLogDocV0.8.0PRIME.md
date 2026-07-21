# CarbonStack LogDoc v0.8.0 PRIME

**Last updated:** 2026-07-16 local session, after operator-reported public posting of the v0.8.0 release  
**Owner/operator:** bitcrusher32  
**Assistant/model:** GPT-5.5 Thinking  
**Current phase:** **v0.8.0 Operational Spine Maturation Pre-Release is posted; v0.7.x operational-spine integration epoch is closed; next action is new evergreen roadmap generation.**  
**Current public release:** `v0.8.0 Operational Spine Maturation Pre-Release`  
**Previous public release:** `v0.7.0 Cumulative Pre-Alpha Engineering Boundary Pre-Release`  
**Current working checkpoint:** `v0.8.0 PRIME — Public Release Posted / v0.7.x Epoch Closed / Roadmap Handoff`  
**Prior planning authority:** `CarbonStack_Long_Term_Roadmap_v0.8.0_EVERGREEN`  
**Next planning authority:** **To be generated next**  
**Update source:** `CarbonStackLogDocV0.7.26.md`, operator-provided v0.8.0 release-posted status, v0.8.0 release asset-prep continuity, and accepted v0.8.0 release-body drafting corrections.  
**Update purpose:** Convert the v0.7.26 private Gate-F-closed release-candidate handoff into the final v0.8.0 PRIME continuity anchor after public release posting; preserve the complete v0.7.x development ledger, release-posting continuity, final public-release meaning, hard nonclaims, downstream effects, and safe handoff into the next evergreen roadmap.

**Process note:** This PRIME is generated outside WSL. It is a post-release continuity document, not an implementation script output and not a public release asset by itself.

**Release-posting note:** The operator reported that the v0.8.0 assets/tag draft were completed and the release was posted. The provided release index is:

```text
https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases
```

**Validation-depth note:** The v0.8.0 public release is a pre-alpha operational-spine maturation checkpoint. It closes the v0.7.x operational-spine integration runway and publicizes the bounded validation package. It does **not** claim production readiness, production E2EE, hostile-server safety, verified identity, trust promotion, public deployment readiness, containers, TUI, Android, CarbonStackOS implementation, PQ/hybrid security, external audit, or general-public usability.

**Authority order after this PRIME:** Public v0.8.0 release assets control public v0.8.0 claims. Current committed code and tests control current behavior. This PRIME controls private post-release continuity. The v0.8.0 EVERGREEN roadmap is now a completed release-era planning artifact and should be deprecated after the next roadmap is generated. The next evergreen roadmap will control future work after acceptance.

**Development-ledger note:** The full point-form development ledger remains a first-class LogDoc feature. It is retained below through v0.8.0 PRIME. Future Markdown LogDocs should keep the ledger compressed but meaningful. JSON Breakpoints, when generated, should remain lean/current-state oriented.

---

## 0. Current Rung Summary

The v0.8.0 release has been posted by the operator.

Current repository heads from the v0.7.26 release-candidate handoff remain the final source-head continuity anchor unless a later public release audit proves otherwise:

| Repository | Current full head | Commit |
|---|---:|---|
| `carbonstack` | `715fd76f7700e63f1de877b730c2ce6bb37225bf` | `715fd76 test: add Gate F release candidate closure profile` |
| `carbonstack-comms` | `32681e20784e560fdc8075bbc8a6597b742823c9` | `32681e2 feat: add Gate F basic local trust posture` |
| `carbonstack-cypher` | `59c732ef53e198f65c03cc4e4178a66521c26a4c` | `59c732e feat: add Cypher config inspection flags` |
| `carbonstack-os` | `1bbbe52020d623b81796694e5057c1d080ede3ea` | `1bbbe52 docs: clarify CarbonStackOS target direction` |

Public release-posting state after operator action:

```text
PUBLIC_V0_8_0_RELEASE_POSTED=true
V0_7_X_OPERATIONAL_SPINE_EPOCH_CLOSED=true
NEXT_REQUIRED_ACTION=generate_next_evergreen_roadmap
```

Current gate state:

- Gate 0: CLOSED
- Gate A: CLOSED
- Gate B: CLOSED
- Gate B1-B9: CLOSED
- Gate C: CLOSED
- Gate C C0-C5: CLOSED
- Gate D: CLOSED
- Gate E: CLOSED for manual-private native deployment
- Gate E E1/E2/E3: CLOSED
- Gate F: CLOSED as `closed_v0_8_0_release_candidate_handoff_ready`
- Gate F F0-F7: CLOSED
- Public v0.8.0 release: POSTED by operator
- Next evergreen roadmap: NOT GENERATED YET
- Later platform/client/adversarial lanes: NOT OPEN under a new accepted roadmap yet

Known baseline still accepted:

- `carbonstack-comms/internal/trust/identity_candidate_review.go:71`
  - self-assignment of `current.ObservedAt`;
  - accepted full-Comms-vet warning;
  - not a v0.8.0 release blocker;
  - must remain classified, not silently normalized outside explicit targeted work.
- Repeated `carbonstack-cypher/cypher.db` artifact/hygiene hit:
  - review-only/non-destructive;
  - final clean status passed before release asset prep;
  - intended operator DB location remains deployment root `state/cypher`.

---

## 1. Delta from v0.7.26

v0.7.26 stopped at:

- Gate 0/A/B/C/D/E/F closed;
- v0.8.0 release-candidate handoff ready;
- public v0.8.0 release not yet created;
- command reference entries=`140`;
- `carbonstack` at `715fd76f7700e63f1de877b730c2ce6bb37225bf`;
- `carbonstack-comms` at `32681e20784e560fdc8075bbc8a6597b742823c9`;
- `carbonstack-cypher` at `59c732ef53e198f65c03cc4e4178a66521c26a4c`;
- `carbonstack-os` at `1bbbe52020d623b81796694e5057c1d080ede3ea`;
- next action was operator-controlled v0.8.0 release preparation.

v0.8.0 PRIME current continuity state is:

- operator reported v0.8.0 release posted;
- release assets were prepared and uploaded manually;
- release body was accepted after correction toward the v0.7.0 template style;
- v0.7.x operational-spine integration epoch is closed;
- v0.8.0 EVERGREEN roadmap has reached its terminal release function;
- next action is to generate the next evergreen roadmap;
- no new implementation gate is open yet.

Important continuity chain:

1. F7 closed Gate F as release-candidate handoff ready.
2. Release asset prep started from the v0.7.26 source-head freeze.
3. Initial asset-prep script had a package-shape brittleness issue:
   - external versioned package-checksum name was used internally;
   - validator expected `release/checksums.txt`;
   - corrected in V2 by using canonical internal metadata paths.
4. V2 asset prep generated the intended v0.8.0 upload asset set and fresh-package validation path.
5. Public release body was drafted from the v0.7.0 template.
6. Operator rejected the first draft for excessive internalism and unrealistic priority labeling.
7. Final accepted draft adhered more closely to v0.7.0 headings/order/tone and used more pragmatic criticality labels.
8. Operator reported release posted.
9. PRIME continuity now supersedes v0.7.26 as the active post-release private anchor.

---

## 2. v0.8.0 Public Release Meaning

v0.8.0 is a Gitea-source-of-truth pre-alpha release.

It truthfully represents:

- a multi-repo operational-spine validation package;
- Debian / WSL Debian first validation;
- Relay lifecycle closure;
- state-substrate closure;
- runtime aggregate closure;
- manual-private native deployment posture closure;
- release/package/runtime maturity closure;
- basic local trust posture;
- package/runtime candidate validation;
- release-candidate closure matrix and manual handoff.

It does not represent:

- a finished messenger;
- production readiness;
- production E2EE;
- hostile-server safety;
- malicious-relay safety;
- verified identity;
- full trust promotion;
- secure enrollment;
- public deployment readiness;
- container readiness;
- TUI/public UX;
- Android;
- CarbonStackOS implementation;
- PQ/hybrid security;
- external audit/certification.

The attached package is the intended validation package. Default Gitea source archives remain non-authoritative for multi-repo validation.

---

## 3. Release Asset / Body Continuity

### 3.1 Uploaded release assets

Final v0.8.0 public release assets were prepared to match the v0.7.0-style release asset set:

```text
carbonstack-v0.8.0-asset-checksums.txt
carbonstack-v0.8.0-operational-spine-maturation-pre-release.tgz
carbonstack-v0.8.0-package-checksums.txt
carbonstack-v0.8.0-release-manifest.json
carbonstack-v0.8.0-validation-freeze.md
LICENSE
v0.8.0-release-notes.md
v0.8.0-testing-runbook.md
```

The release package uses canonical internal package metadata:

```text
release/manifest.json
release/checksums.txt
```

External uploaded companion assets use versioned filenames.

### 3.2 Release body continuity

The accepted public release body deliberately follows the v0.7.0 structure:

- title;
- status and release metadata;
- introductory release boundary paragraphs;
- Testing notes;
- Release assets;
- What changed;
- Validated;
- Boundary.

Final public-body style constraints:

- semi-public, not private LogDoc style;
- concise, not exhaustive;
- pragmatic criticality labels;
- no table where everything is high/release-critical;
- same general heading order and tone as v0.7.0;
- deep continuity remains in PRIME/LogDocs rather than public body.

---

## 4. Development Ledger Continuity Rule

The point-form development ledger is a **first-class LogDoc feature** and MUST be retained across future LogDoc updates.

Every future LogDoc update must preserve the development ledger as a point-form, per-breakpoint summation of important failures, implementations, decisions, evidence/validation posture, changed critical paths, and safe-handoff constraints. The ledger may be compressed and curated to avoid verbatim historical bloat, but it must not be deleted, replaced by only the most recent delta, or collapsed into a generic short summary unless the operator explicitly requests an archival reset.

The JSON Breakpoint remains lean and current-state oriented. The Markdown LogDoc carries continuity memory.

---

## 5. Full Point-Form LogDoc Timeline Continuity Ledger
**Posture:** This ledger replaces verbatim historical LogDoc appendices with grouped point-form continuity records. It preserves meaningful engineering memory without exponential context growth.

**Rule going forward:** Preserve decisions, failures, semantic contracts, changed surfaces, evidence roots, validation profiles, and safe-resume instructions. Do not embed full prior LogDocs unless the operator explicitly asks for archival reconstruction.

## v0.7.1 — Gate A Acceptance Package

- **Purpose**
  - Start the active v0.7.x operational-spine ledger after the public v0.7.0 release.
  - Accept the Gate A workflow contract before implementation.
  - Keep Gate 0/v0.7.0 release boundary stable and avoid replaying the full v0.6.x archive.
- **Decisions**
  - Future work must use explicit operator lifecycle, command-table-visible leaves, workspace context, state-path discipline, mutation classes, partial-completion reporting, and strict ACK discipline.
  - Relay membership is routing authority only, not trust or MLS membership.
- **Safe handoff**
  - Gate A package first; no Gate B mutation from memory.

## v0.7.2 — Gate A Closure / Gate B Entry

- **Purpose**
  - Close Gate A durably and open Gate B with controlled implementation posture.
- **Meaningful work**
  - Committed Gate A documents and hardened command/registry expectations.
  - Established that Gate B must proceed through narrow leaves rather than broad “full runtime” aggregation.
- **Continuity lesson**
  - Keep release boundary and private implementation boundary separate.

## v0.7.3 — Gate B3 Normal Message Path Closure

- **Purpose**
  - Close scoped normal-message routing and classification through B1-B3.
- **Meaningful work**
  - Closed unscoped inbox isolation, active Relay Space routing enforcement, and Relay Space-scoped normal-message wrappers.
  - Preserved unsupported-envelope visibility.
- **Issues / blunders**
  - Gate B3b churn hazard:
    - incomplete cross-callsite reconnaissance;
    - stale negative-path assumptions;
    - helper filters hiding unsupported scoped envelopes;
    - brittle mutation/status guards.
- **Lesson**
  - Mandatory consolidated recon before high-surface subgate mutation.

## v0.7.4 — Gate B4 Invite/Member Lifecycle Closure

- **Purpose**
  - Close bounded invite/member lifecycle and active/disabled/left routing semantics.
- **Meaningful work**
  - Invite claim, member state, restart persistence, and delivery authority became explicit.
  - Left-to-active requires explicit rejoin workflow.
- **Nonclaims**
  - B4 does not prove identity/trust, MLS membership, production routing, or deployment readiness.
- **Safe handoff**
  - B5 KeyPackage lifecycle next; consolidated recon required.

## v0.7.5 — Gate B5 Prerequisite / B5a

- **Purpose**
  - Establish scoped delivery authority and KeyPackage inspection boundaries.
- **Meaningful work**
  - Active Relay Space membership became prerequisite for scoped inbox and ACK.
  - B5a closed local KeyPackage inspection/ownership/lifetime/integrity evidence.
- **Decision**
  - `key_package_ref`, artifact hash, envelope ID, generation ID, request ID, and Relay destination identity are distinct.
- **Safe handoff**
  - B5b repeatable generation/rotation next.

## v0.7.6 — Gate B5b Repeatable Generation/Rotation

- **Purpose**
  - Close repeatable KeyPackage generation/rotation/inventory behavior.
- **Meaningful work**
  - Added persistent inventory, request idempotence, legacy adoption, explicit retirement, concurrency serialization, restart and manifest recovery.
- **Blunder avoided/fixed**
  - Fixed public bundle file handling was not enough; provider state had to be loaded/evolved safely.
- **Process preference**
  - Spend more time on recon/precomputation/shadow builds to reduce destructive churn.

## v0.7.7 — Gate B5c Relay KeyPackage Publication

- **Purpose**
  - Close deterministic Relay publication of selected active B5b KeyPackage generation.
- **Meaningful work**
  - Added persistent sender-scoped publication binding, exact replay, reuse conflict, identity conflict, and restart persistence.
  - Public KeyPackage directory was explicitly not added.
- **Recovery**
  - Recovered from context-window/churn failure using log-only WSL pasteback.
  - Patched stale `acknowledged_at` fixture assumption: ACK timestamps live in `envelope_acks`, not `envelopes`.
- **Lesson**
  - Recover from shadow evidence with targeted fixes; avoid hash-heavy context flooding.

## v0.7.8 — Gate B5d/B5e KeyPackage Consume and Gate B5 Closure

- **Purpose**
  - Close local KeyPackage consume receipt and Gate B5 lifecycle profile/docs/registry.
- **Meaningful work**
  - Added `openmls-relay-keypackage-consume-dev`.
  - ACK happens only after local KeyPackage artifact/receipt persistence.
  - Exact replay returns `already_consumed` at runtime while durable receipt history remains stable.
- **Recovery**
  - Corrected replay-test expectation so durable history is not rewritten.
  - Corrected validator API mismatch: `RequirePaths` vs `CheckRequiredPaths`.
- **Safe handoff**
  - B6 Welcome lifecycle next; do not treat KeyPackage receipt as add-member, Welcome, trust, identity, or production distribution.

## v0.7.9 — Gate B6 and Gate B7 Closure

- **B6 purpose**
  - Close Welcome consume/join/ACK-after-join.
- **B6 meaningful work**
  - Added `openmls-relay-welcome-consume-dev` and `welcome-lifecycle-dev`.
  - Welcome artifact is persisted before sidecar `conversation-join`.
  - Successful join evidence is persisted before ACK.
  - Failed join leaves unacked receipt evidence.
- **B7 purpose**
  - Close Cypher/MLS mismatch inspection/refusal.
- **B7 meaningful work**
  - Added `openmls-cypher-mls-mismatch-inspect-dev` and `cypher-mls-mismatch-dev`.
  - Classified aligned, active/absent, inactive/present, inactive-with-receipts, wrong-relay-space, local-device, sidecar-label, conversation-label, unsupported/stale/ambiguous, and incomplete local consume/join states.
  - Unsafe states refuse.
- **Process lesson**
  - B7 correctly produced no LogDoc/Breakpoint in WSL and stopped at repo commits/evidence log.
- **Safe handoff**
  - B8 workflow engine next; B7 is a refusal leaf, not an orchestration layer.

## v0.7.10 — Gate B8 Workflow Report/Evaluator Closure

- **Purpose**
  - Close reusable noninteractive workflow report/evaluator over B4/B5/B6/B7 leaf evidence.
- **Meaningful work**
  - Added `workflow-relay-onboarding-dev` command.
  - Added `workflow-relay-onboarding-dev` validator profile.
  - Added durable workflow report and replay behavior.
  - Added stage-level output preserving:
    - ready local state;
    - B7 mismatch result;
    - KeyPackage receipt result;
    - Welcome receipt result;
    - workflow result.
- **Closed semantics**
  - `workflow_ready` requires aligned B7 state plus persisted/ACKed KeyPackage and Welcome receipts.
  - B7 refusal blocks workflow progression.
  - Missing/incomplete onboarding state is refused as partial rather than hidden.
  - Leaf boundaries remain visible in stage output.
- **Nonclaims**
  - Not B9 Gate B closure.
  - Not `full-runtime-dev`.
  - Not production E2EE.
  - Not trust promotion, verified identity, silent repair/rejoin, or Cypher/MLS reconciliation.
- **Safe handoff**
  - B9 Gate B integration/closure is next.
  - Start B9 with deep recon and contract freeze.
  - Do not open Gate C until B9 is closed or explicitly replanned.

---

## v0.7.11 — Gate B9 Integration Closure / Gate B Closed / Gate C Blocked

- Closed B9 as a `carbonstack`-only validator/docs/registry closure rung.
- Added `gate-b-relay-lifecycle-closure-dev` as the Gate B closure profile.
- Added docs:
  - `docs/250-v0.7.10-gate-b9-integration-closure-v0.md`
  - `docs/251-v0.7.10-gate-b-closure-v0.md`
- Added validator sources:
  - `tools/carbonstack-validate/gate_b_relay_lifecycle_closure_dev.go`
  - `tools/carbonstack-validate/gate_b_relay_lifecycle_closure_dev_test.go`
- Updated:
  - `tools/carbonstack-validate/main.go`
  - `docs/README.md`
  - `registry/commands.v0.yaml`
  - `registry/COMMAND_REFERENCE.v0.md`
  - `registry/COMMAND_BOUNDARY_TABLE.v0.md`
- Advanced command reference from `entries=117` to `entries=118`.
- Proved the Gate B closure ladder:
  - B1-B4 Relay Space routing/member lifecycle regressions;
  - B5 KeyPackage inspection/generation/publication/consume regressions;
  - B6 Welcome consume/join/ACK-after-join regression;
  - B7 Cypher/MLS mismatch inspection/refusal regression;
  - B8 workflow report/evaluator regression;
  - lower-level Relay OpenMLS join and same-state normal-message regressions;
  - registry/reference freshness and missing-nonclaims zero.
- Closed Gate B as a coherent dev/pre-alpha Relay lifecycle integration lane.
- Did not start Gate C.
- Did not mutate `carbonstack-comms`, `carbonstack-cypher`, or `carbonstack-os`.
- Did not create a new Comms command.
- Did not promote `full-runtime-dev`.
- Did not claim production E2EE, verified identity, trust promotion, secure enrollment, hostile-server safety, deployment readiness, containers, PQ/hybrid support, vault/backup/restore, Android, or CarbonStackOS implementation.
- Noted the recurring first-push authentication failure and successful retry as continuity noise, not a project blocker.
- Noted the shadow `carbonstack-cypher/cypher.db` artifact scan hit remained inside the shadow evidence path and is non-destructive review evidence, not a semantic closure issue.

## v0.7.12 — Gate C Recon / State Substrate Contract Freeze / C1 Blocked

- **Purpose**
  - Preserve the completed Gate C state-substrate deep recon and freeze the Gate C contract before implementation.
  - Establish that Gate C should proceed by subgates rather than one broad mutation pass.
  - Capture the accepted canonical Comms-owned state-root posture with compatibility caveats.
- **Evidence**
  - Gate C recon ran with no tracked-source mutation, no shadow mutation, and no WSL LogDoc/Breakpoint generation.
  - Recon confirmed post-Gate-B heads remained clean across all four repositories.
  - Recon summary ended with `RECON_COMMAND_FAILURE_COUNT=0`.
  - Continuity files were missing inside WSL by design; downloadable continuity artifacts remain generated outside WSL.
- **Contract frozen**
  - Gate C defines and enforces the first bounded CarbonStack state substrate.
  - Gate C is not a vault, backup/restore system, secure storage layer, deployment system, TUI, full-runtime aggregate, trust system, verified identity layer, or production E2EE claim.
  - Gate C makes safety-sensitive state explicit, inspectable, versioned where appropriate, and refusal-driven when uncertain.
- **Canonical state-root posture**
  - A canonical Comms-owned state root is accepted as a policy anchor, not as a brittle mandatory chokepoint.
  - Default operator behavior should remain intuitive, likely centered on `.carbonstack-comms/` and the existing explicit `--state` path model.
  - Existing `--state` compatibility must remain valid.
  - Roots derived beside `--state` remain supported where already established, especially KeyPackage receipts, Welcome receipts, and workflow reports.
  - New Gate C surfaces may introduce explicit `--state-root` only if it improves clarity and does not silently override `--state`.
  - Deep recon and read-only inspection must continue to work without requiring a perfect canonical layout.
- **Subgate plan**
  - C1: state substrate inventory and authority map.
  - C2: schema/version metadata and compatibility refusal for Comms-owned JSON state/report/receipt files.
  - C3: path policy and explicit state-root semantics.
  - C4: atomic write, lock, partial-state, and replay classification policy.
  - C5: Gate C closure profile proving no-silent repair/migration and machine-readable inspection.
- **Nonclaims**
  - Not production security.
  - Not verified identity.
  - Not trust promotion.
  - Not secure enrollment.
  - Not hostile-server safety.
  - Not vault security.
  - Not backup/restore safety.
  - Not deployment readiness.
  - Not `full-runtime-dev`.
  - Not Gate D/E/F.
  - Not PQ/hybrid.
  - Not Android or CarbonStackOS.
- **Safe handoff**
  - Implement C1 first as a read-only state inventory/authority-map surface.
  - Do not mutate sidecar cryptographic provider internals, Cypher DB schema, trust promotion, verified identity, backup/restore, deployment, TUI, PQ, Android, CarbonStackOS, or `full-runtime-dev` during C1.

---


## v0.7.13 — Gate C1 State Substrate Inventory / C2 Blocked

- **Purpose**
  - Close Gate C1 as the first executable state-substrate inventory and authority-map surface.
  - Move from v0.7.12 contract freeze into actual Gate C implementation while keeping the first subgate read-only-by-default.
- **Meaningful work**
  - Added Comms command:
    - `state-substrate-inventory-dev`
    - registry ID `comms.state-substrate-inventory-dev`
    - report schema `carbonstack-state-substrate-inventory/v0`
  - Added validator profile:
    - `state-substrate-inventory-dev`
    - registry ID `runner.state-substrate-inventory-dev`
  - Added docs:
    - `docs/252-v0.7.12-gate-c1-state-substrate-inventory-v0.md`
    - `docs/253-v0.7.12-gate-c1-closure-v0.md`
  - Updated command registry, generated command reference, boundary table, validator dispatch, and docs README.
  - Advanced command reference from `entries=118` to `entries=120`.
- **Closed semantics**
  - C1 inventory is read-only by default.
  - Optional `--output` writes a machine-readable generated evidence report, not runtime state.
  - Explicit `--state` compatibility is preserved.
  - Canonical Comms-owned state root is a policy anchor, not a brittle mandatory chokepoint.
  - The report can classify Comms-owned, sidecar-owned, Cypher-owned, validator-generated, evidence-generated, legacy/unversioned, unsupported, and unknown state surfaces.
  - Supported receipt/workflow schemas are detected.
  - Unsupported schemas are classified for future C2 refusal; C1 itself does not enforce C2 behavior.
- **Critical report coverage**
  - Comms state root and state file.
  - Trust records, trust events, and identity candidates.
  - KeyPackage receipt roots and discovered `receipt.json` manifests.
  - Welcome receipt roots and discovered `receipt.json` manifests.
  - Workflow report roots and discovered `workflow-report.json` manifests.
  - Sidecar project directory and sidecar generated state root.
  - Optional Cypher DB path.
  - Optional evidence root.
- **Validation**
  - Shadow and live render/check advanced command reference to `entries=120`.
  - C1 tests passed repeatedly.
  - C1 adjacent app/state/trust tests passed.
  - Validator tests passed.
  - C1 profile passed.
  - Gate B closure regression passed.
  - Registry lookups for both C1 IDs passed.
  - Missing-nonclaims check passed.
  - Final clean status was confirmed for all four repositories.
- **Process/blunder note**
  - The implementation script took about 37 minutes because it overvalidated an inventory-only subgate:
    - `StateSubstrateInventory` tests ran with `-count=30` in shadow and live;
    - the C1 profile itself invoked the full Gate B closure ladder;
    - the script also ran explicit Gate B closure and additional regression profiles.
  - This was finite and successful, not an infinite loop, but it was excessive for C1.
  - Future scripts should reserve full repeated closure ladders for closure gates, high-risk semantic changes, or explicit operator request.
  - For narrow inventory/schema/path subgates, default to targeted tests plus one bounded regression profile at most.
- **Nonclaims**
  - Not C2 schema enforcement.
  - Not migration.
  - Not silent repair.
  - Not cleanup/deletion.
  - Not vault or backup/restore.
  - Not trust promotion.
  - Not verified identity.
  - Not Cypher/MLS reconciliation.
  - Not deployment.
  - Not `full-runtime-dev`.
  - Not Gate D runtime aggregate.
  - Not production E2EE.
- **Safe handoff**
  - C2 is next and must focus on schema/version metadata and compatibility refusal for Comms-owned JSON state/report/receipt files.
  - Do not broaden C2 into migration, cleanup, vault/backup, deployment, TUI, Gate D, or full-runtime aggregation.

## v0.7.14 — Gate C2 State Schema Compatibility Closure / C3 Blocked

- **Purpose**
  - Close Gate C2 as the first explicit schema/version compatibility and refusal surface for Comms-owned JSON state/report/receipt artifacts.
  - Preserve the C1 inventory posture while adding a narrow compatibility classifier rather than broad runtime migration.
- **Meaningful work**
  - Added `state-schema-compat-dev` Comms command.
  - Added `state-schema-compat-dev` validator profile.
  - Added report schema `carbonstack-state-schema-compatibility-report/v0`.
  - Added docs:
    - `docs/254-v0.7.13-gate-c2-state-schema-compatibility-v0.md`
    - `docs/255-v0.7.13-gate-c2-closure-v0.md`
  - Added registry IDs:
    - `comms.state-schema-compat-dev`
    - `runner.state-schema-compat-dev`
  - Advanced command reference from `entries=120` to `entries=122`.
- **Closed semantics**
  - Supported Comms-owned JSON schemas allow:
    - `carbonstack-state-substrate-inventory/v0`
    - `carbonstack-keypackage-consume-receipt/v0`
    - `carbonstack-welcome-consume-receipt/v0`
    - `carbonstack-workflow-relay-onboarding-report/v0`
    - `carbonstack-cypher-mls-mismatch-report/v0`
  - Unsupported newer safety-sensitive schemas refuse.
  - Missing schemas on safety-sensitive receipt/report artifacts refuse.
  - Invalid JSON refuses.
  - Legacy/unversioned local Comms state is classified without migration.
  - Optional `--output` writes generated evidence only.
- **Validation**
  - Shadow and live gofmt, render, diff checks passed.
  - Shadow and live C2 tests passed with `-count=1`.
  - Shadow and live C2/C1 focused tests passed.
  - Shadow and live validator tests passed.
  - Shadow and live `state-schema-compat-dev` profile passed.
  - Shadow and live registry lookups for both C2 IDs passed.
  - Shadow and live missing-nonclaims checks passed.
  - Light live smoke proved C2 can read a C1 inventory schema and C1 remains callable.
  - Final clean status was confirmed for all four repositories.
- **Process note**
  - C2 used the intended proportionate validation posture:
    - no `-count=30`;
    - no repeated Gate B closure ladder;
    - no C1 profile rerun;
    - no `full-runtime-dev`;
    - no deployment/PQ/vault work.
  - This correctly applied the v0.7.13 lesson from the 37-minute C1 run.
- **Nonclaims**
  - Not migration.
  - Not silent repair.
  - Not state relocation.
  - Not deletion/cleanup.
  - Not vault or backup/restore.
  - Not trust promotion.
  - Not verified identity.
  - Not Cypher/MLS reconciliation.
  - Not deployment.
  - Not `full-runtime-dev`.
  - Not Gate D runtime aggregate.
  - Not production E2EE.
- **Safe handoff**
  - C3 is next and should focus on path policy and explicit state-root semantics.
  - Do not broaden C3 into migration, C4 atomicity, Gate D aggregation, deployment, TUI, vault/backup, trust, identity, PQ, Android, or CarbonStackOS work.


---

## v0.7.15 — Gate C3 Path Policy Closure / C4 Blocked

- **Purpose**
  - Close Gate C3 as the path-policy and explicit state-root semantics subgate.
  - Convert the C3 recon contract into an executable, machine-readable, dev/pre-alpha path-policy surface.
  - Preserve intuitive canonical-root posture without turning `.carbonstack-comms` into a brittle mandatory chokepoint.
- **Meaningful work**
  - Added `state-path-policy-dev` in `carbonstack-comms`.
  - Added validator profile `state-path-policy-dev` in `carbonstack`.
  - Added report schema `carbonstack-state-path-policy-report/v0`.
  - Added docs:
    - `docs/256-v0.7.14-gate-c3-path-policy-v0.md`
    - `docs/257-v0.7.14-gate-c3-closure-v0.md`
  - Added registry IDs:
    - `comms.state-path-policy-dev`
    - `runner.state-path-policy-dev`
  - Command reference moved to `entries=124`.
- **Closed semantics**
  - Explicit `--state` compatibility is preserved.
  - `.carbonstack-comms` is preferred canonical policy, not mandatory layout enforcement.
  - Explicit `--state-root` mismatch is classified rather than refused.
  - Parent traversal is refused as unsafe path policy.
  - Sidecar, Cypher, validator, and evidence roots are classified only.
  - Optional `--output` is generated evidence only.
- **Validation**
  - Shadow and live C3 tests passed with `-count=1`.
  - Shadow and live C3/C2/C1 focused tests passed.
  - Shadow and live validator tests passed.
  - Shadow and live `state-path-policy-dev` profiles passed.
  - Generated reference checks passed with `entries=124`.
  - Registry lookups for both C3 IDs passed.
  - Missing-nonclaims checks passed.
  - Light live smoke proved C2 and C1 remain callable and C3 explicit-root classification works.
  - Final clean status was confirmed for all four repositories.
- **Process notes**
  - C3 used the corrected proportionate validation posture:
    - no `-count=30`;
    - no repeated Gate B closure ladder;
    - no broad regression ladder;
    - no deployment/PQ/vault work.
  - This is the preferred posture for narrow Gate C subgates.
  - The recurring `carbonstack-cypher/cypher.db` artifact-scan hit remains review-only when it appears as a known local/generated database and final git clean status is still clean.
- **Known limitations / future handoff**
  - C3 does not implement C4 atomicity/lock closure.
  - C3 does not migrate, repair, relocate, delete, clean, or normalize state.
  - C3 refuses parent traversal; deeper symlink/filesystem hardening remains a later C4/C5 concern unless a concrete defect appears.
- **Nonclaims**
  - Not C4 atomicity or lock closure.
  - Not C5 Gate C closure.
  - Not migration.
  - Not silent repair.
  - Not state relocation.
  - Not cleanup implementation.
  - Not vault or backup/restore.
  - Not trust promotion.
  - Not verified identity.
  - Not Cypher/MLS reconciliation.
  - Not deployment.
  - Not `full-runtime-dev`.
  - Not Gate D runtime aggregate.
  - Not production E2EE.
- **Safe handoff**
  - C4 is next and should focus on atomic write, lock, partial-state, and replay/recovery classification policy.
  - Do not broaden C4 into Gate D aggregation, deployment, TUI, vault/backup, trust, identity, PQ, Android, or CarbonStackOS work.

## v0.7.16 — Gate C4 Write Policy Closure / C5 Blocked Handoff

- **Purpose**
  - Close Gate C4 as a write-policy, atomicity, lock, partial-state, replay, and cleanup-boundary classification subgate.
  - Preserve a bounded, classify-first posture instead of rewriting every runtime writer before Gate C closure.
- **Meaningful work**
  - Added `state-write-policy-dev` in CarbonStackComms.
  - Added `state-write-policy-dev` validator profile in CarbonStack.
  - Added report schema `carbonstack-state-write-policy-report/v0`.
  - Added docs:
    - `docs/258-v0.7.15-gate-c4-write-policy-v0.md`
    - `docs/259-v0.7.15-gate-c4-closure-v0.md`
  - Added registry IDs:
    - `comms.state-write-policy-dev`
    - `runner.state-write-policy-dev`
  - Command reference advanced to `entries=126`.
- **Closed classifications**
  - C1/C2/C3 generated report writers are classified as atomic generated evidence writers using temporary-file then rename discipline.
  - B5d KeyPackage consume receipts are classified as atomic lock-guarded receipt writers with persisted-before-ACK and exact `already_consumed` replay semantics.
  - B6 Welcome consume receipts are classified as atomic lock-guarded receipt writers with Welcome-persisted-before-join, joined-evidence-before-ACK, failed-join-unacked, and exact `already_joined` replay semantics.
  - B8 workflow reports are classified as atomic report writers with stage-preserving partial workflow classification and exact `already_reported` replay semantics.
  - Local Comms state, trust state, and identity-candidate state are classified as direct-write / append-current-behavior surfaces and recorded as future-hardening warnings.
  - Sidecar, Cypher, validator, and evidence roots remain classified only.
- **Validation**
  - Shadow and live C4 tests passed at `-count=1`.
  - Focused C4/C3/C2/C1 Comms tests passed.
  - Validator tests passed.
  - `state-write-policy-dev` profile passed.
  - Registry lookups passed for C4 command/profile IDs.
  - Missing nonclaims remained zero.
  - Generated command reference check passed at `entries=126`.
  - Light C4/C3/C2/C1 smoke passed.
  - Live temp-residue check found no `.tmp` or `.lock` residue in the C4 smoke root.
  - Final clean state passed for all four repositories.
- **Process / blunders**
  - Correctly avoided `-count=30` and repeated Gate B closure ladders.
  - Correctly kept C4 classify-first and did not perform runtime writer rewiring.
  - `carbonstack-comms` push hit the recurring first-attempt authentication failure and succeeded on retry.
  - The C4 profile step label saying C2 “accepts C3 path-policy schema” is semantically misleading: the command actually classified/refused `carbonstack-state-path-policy-report/v0` as unsupported under `--kind state-substrate-inventory`, then exited zero because `--allow-refusal-exit-zero` was used. Treat this as a wording hazard and future C5 to-do, not a C4 closure failure.
  - Do not claim C2 supports C3/C4 report schemas until a targeted update explicitly adds those schema kinds or the C5 closure profile validates them by another explicit route.
- **Nonclaims**
  - C4 does not claim C5 Gate C closure, universal runtime writer rewiring, migration, repair, state relocation, cleanup implementation, destructive cleanup, vault/backup safety, trust promotion, verified identity, Cypher/MLS reconciliation, deployment, full-runtime-dev, Gate D, production E2EE, PQ/hybrid migration, Android, or CarbonStackOS implementation.
- **Safe handoff**
  - C5 is next. It should close Gate C by proving C1-C4 surfaces together:
    - inventory/authority map;
    - schema/refusal compatibility;
    - path policy;
    - write/partial/replay policy;
    - registry/reference freshness;
    - no missing nonclaims;
    - no silent repair/migration/relocation/cleanup/trust promotion.
  - C5 must not broaden into Gate D aggregation, full-runtime-dev, deployment, TUI, vault/backup, PQ, Android, or CarbonStackOS.

---

## v0.7.17 — Gate C5 State Substrate Closure / Gate C Closed / Gate D Blocked

- **Purpose**
  - Close Gate C as the bounded state-substrate enforcement lane.
  - Convert C1-C4 from isolated classification surfaces into one accepted Gate C closure profile.
  - Preserve the hard boundary that Gate D runtime aggregate work is not started by Gate C closure.
- **Meaningful work**
  - Added `gate-c-state-substrate-closure-dev` validator profile.
  - Added registry ID `runner.gate-c-state-substrate-closure-dev`.
  - Added closure docs:
    - `docs/260-v0.7.16-gate-c5-state-substrate-closure-v0.md`
    - `docs/261-v0.7.16-gate-c-closure-v0.md`
  - Expanded C2 schema compatibility to explicitly support:
    - `path-policy-report` -> `carbonstack-state-path-policy-report/v0`
    - `write-policy-report` -> `carbonstack-state-write-policy-report/v0`
  - Corrected the C4 wording/route mismatch where a validator step label implied C2 accepted the C3 path-policy schema while the command had actually refused it under the wrong `state-substrate-inventory` kind with `--allow-refusal-exit-zero`.
  - Proved C1/C2/C3/C4 profiles together inside C5 without rerunning the full Gate B closure ladder.
  - Preserved `GATE_D_STATUS=not_started`.
- **Validation**
  - Shadow and live `StateSchemaCompatibility` tests passed.
  - Shadow and live Gate C focused C1-C4 app tests passed.
  - Validator package tests passed.
  - C4 profile passed after explicit C2 kind correction.
  - C5 closure profile passed.
  - Generated command reference current at `entries=127`.
  - Registry lookup for `runner.gate-c-state-substrate-closure-dev` passed.
  - Missing nonclaims remained zero.
  - Final clean state confirmed across all four repositories.
- **Blunders / process correction**
  - First C5 script failed before live mutation because it used a brittle exact multi-line anchor against `state_schema_compat_dev.go`.
  - That first failed script also continued to `gofmt` after Python generation failed, producing a second-order missing-file failure.
  - Retry fixed both hazards:
    - robust C2 schema-kind insertion rather than exact multi-line anchoring;
    - explicit generation return-code guard before formatting/validation.
  - This is now a recorded brittleness lesson: for closure scripts with generated files, generation must hard-stop before any subsequent path-dependent step.
- **Nonclaims preserved**
  - Gate C closure is dev/pre-alpha state-substrate closure only.
  - It is not Gate D runtime aggregate.
  - It is not `full-runtime-dev`.
  - It is not migration, silent repair, state relocation, runtime writer rewiring, cleanup implementation, destructive cleanup, vault security, backup/restore safety, trust promotion, verified identity, Cypher/MLS reconciliation, deployment readiness, production E2EE, PQ/hybrid migration, Android, or CarbonStackOS implementation.
- **Safe handoff**
  - Gate C is closed.
  - Gate D remains blocked until a fresh Gate D recon and contract freeze is accepted.
  - Do not begin runtime aggregate work from memory or by stretching the C5 closure profile.

## v0.7.18 — Gate D Runtime Aggregate Closure / Gate D Closed / Gate E Blocked

- **Purpose**
  - Close Gate D as the first profile-only mechanical runtime aggregate after Gate B and Gate C closure.
  - Prove one coherent dev/pre-alpha runtime aggregate without prematurely promoting `full-runtime-dev`.
  - Move the safe handoff from Gate D recon/contract freeze to Gate E recon/contract freeze.
- **Preflight / decision layer**
  - Accepted that `full-runtime-dev` should remain reserved until material runtime evidence justifies the name.
  - Accepted `gate-d-runtime-aggregate-dev` as the first Gate D surface.
  - Accepted a mechanical lifecycle proof as registry-visible runbook/testing evidence.
  - Accepted minimal normal-message send/inbox evidence as part of runtime aggregate closure.
  - Accepted that Gate D must not absorb Gate E deployment semantics.
  - Accepted direct state/trust/candidate writer warnings as loud non-blocking warnings unless concrete recon evidence shows a runtime-aggregate hazard.
  - Accepted JSON authority plus concise human summary.
  - Accepted that verified identity/trust promotion must be ironed out during v0.8.x rather than silently absorbed by Gate D.
  - Deferred TUI to mid/late v0.8.x; CLI/run-space/context remains scaffolding, not a UI mutation surface.
- **Gate D recon attempt 1**
  - Read-only and no tracked mutation.
  - Clean repo orientation was confirmed.
  - Oversized static scan froze/ran excessively for roughly 25 minutes.
  - Operator interrupted it safely.
  - Lesson: recon scripts must be bounded, capped, and hostile to huge full-repo term dumps.
- **Gate D recon retry1**
  - Read-only and bounded.
  - Froze the first Gate D contract candidate:
    - profile `gate-d-runtime-aggregate-dev`;
    - registry ID `runner.gate-d-runtime-aggregate-dev`;
    - report schema `carbonstack-gate-d-runtime-aggregate-report/v0`;
    - explicit run-space/context candidate;
    - profile-only first implementation;
    - no `full-runtime-dev`, no TUI, no Gate E, no verified identity/trust promotion.
  - Completed with `RECON_COMMAND_FAILURE_COUNT=0`.
  - Accepted as sufficient evidence to proceed to implementation.
- **Gate D implementation attempt 1**
  - Shadow-first and generation-hard-stopped.
  - Failed safely during shadow generation before live apply/commit/push.
  - Failure: `could not find validator dispatch insertion point for Gate D`.
  - Cause: brittle assumption that `tools/carbonstack-validate/main.go` contained a literal `case "full":` dispatch insertion point.
  - Lesson: dispatch insertion should anchor on nearby known semantic cases or use syntax-aware insertion, not assumed future-profile blocks.
- **Gate D implementation retry1 meaningful work**
  - Added validator profile:
    - `gate-d-runtime-aggregate-dev`
    - registry ID `runner.gate-d-runtime-aggregate-dev`
    - report schema `carbonstack-gate-d-runtime-aggregate-report/v0`
    - run-space context candidate schema `carbonstack-run-space-context-candidate/v0`
  - Added validator source/test:
    - `carbonstack/tools/carbonstack-validate/gate_d_runtime_aggregate_dev.go`
    - `carbonstack/tools/carbonstack-validate/gate_d_runtime_aggregate_dev_test.go`
  - Updated validator dispatch/help:
    - `carbonstack/tools/carbonstack-validate/main.go`
  - Added docs:
    - `carbonstack/docs/262-v0.7.17-gate-d-runtime-aggregate-v0.md`
    - `carbonstack/docs/263-v0.7.17-gate-d-closure-v0.md`
  - Updated docs README, command registry, generated command reference, and command boundary table.
  - Advanced command reference to `entries=128`.
  - Committed only `carbonstack`:
    - `7daf0c6eefa771ed1c8531a15761b8e797f56914`
    - `test: add Gate D runtime aggregate profile`
  - `carbonstack-comms`, `carbonstack-cypher`, and `carbonstack-os` remained unchanged.
- **Gate D closed proof shape**
  - `gate-d-runtime-aggregate-dev` proves:
    - Gate C closure profile passes;
    - explicit run-space/context candidate is generated and inspected;
    - C1/C2/C3/C4 state-substrate preflight passes over explicit run-space paths;
    - workflow relay onboarding profile passes;
    - Relay Space member restart/resume inspection profile passes;
    - same-state integrated proof covers KeyPackage -> Welcome -> `message-send-dev` -> `message-inbox-dev --ack`;
    - registry/reference checks pass;
    - missing nonclaims remain zero;
    - Gate B closure authority remains present without rerunning the full Gate B closure ladder as an implementation side-effect.
- **Validation**
  - Shadow and live gofmt passed.
  - Shadow and live command reference render passed at `entries=128`.
  - Shadow and live diff checks passed across all four repos.
  - Shadow and live validator tests passed.
  - Shadow and live `gate-d-runtime-aggregate-dev` profile passed.
  - Shadow and live generated reference checks passed.
  - Shadow and live registry lookup for `runner.gate-d-runtime-aggregate-dev` passed.
  - Shadow and live missing-nonclaims checks passed.
  - Light adjacent non-promotion check passed.
  - Final clean status passed for all four repos.
- **Process / blunders**
  - Initial Gate D recon was too broad and produced excessive output; retry1 narrowed and bounded it.
  - Initial Gate D implementation used a brittle validator dispatch anchor; retry1 anchored after the known Gate C closure profile case.
  - Retry1 preserved `set -Eeuo pipefail`, generation hard-stop, shadow-first mutation, no per-file hash guards, and no WSL LogDoc/Breakpoint generation.
  - First push for `carbonstack` hit the recurring auth failure and succeeded on retry.
  - Recurring `carbonstack-cypher/cypher.db` artifact scan remained review-only while final clean status passed.
- **Nonclaims preserved**
  - Gate D closure is dev/pre-alpha runtime aggregate proof only.
  - It does not promote `full-runtime-dev`.
  - It does not start Gate E deployment.
  - It does not start TUI/public UX.
  - It does not claim production readiness, production E2EE, hostile-server safety, verified identity, trust promotion, vault security, backup/restore safety, public ingress, PQ/hybrid support, Android, or CarbonStackOS implementation.
- **Safe handoff**
  - Gate D is closed.
  - Gate E is next but blocked until fresh Gate E recon and contract freeze.
  - Gate E should focus on one preferred private single-host Debian-family deployment model.
  - Do not start Gate E implementation, `full-runtime-dev` promotion, TUI, containers/public ingress, migration, silent repair, destructive cleanup, trust promotion, verified identity, PQ/hybrid, Android, CarbonStackOS, or v0.8.0 release packaging from this handoff.

---

## v0.7.19 — Gate E E1 Manual-Private Native Deployment / E1 Closed / E2 Blocked

- **Purpose**
  - Close the first Gate E native-deployment rung without jumping to a service model, systemd, helper install, TUI, containers, public ingress, or `full-runtime-dev`.
  - Convert the Gate E recon contract into a bounded `carbonstack` validator profile proving a manual-private explicit-env native deployment fixture.
  - Cut a breakpoint before any further Gate E subgates.
- **Meaningful work**
  - Added validator profile `gate-e-native-deployment-dev`.
  - Added registry ID `runner.gate-e-native-deployment-dev`.
  - Added report schema `carbonstack-gate-e-native-deployment-report/v0` and deployment context candidate schema `carbonstack-gate-e-deployment-context-candidate/v0` inside the profile.
  - Added docs:
    - `docs/264-v0.7.18-gate-e-e1-native-deployment-v0.md`
    - `docs/265-v0.7.18-gate-e-e1-closure-v0.md`
  - Updated `docs/README.md`, `registry/commands.v0.yaml`, `registry/COMMAND_REFERENCE.v0.md`, `registry/COMMAND_BOUNDARY_TABLE.v0.md`, and validator dispatch in `tools/carbonstack-validate/main.go`.
  - Added `tools/carbonstack-validate/gate_e_native_deployment_dev.go` and `tools/carbonstack-validate/gate_e_native_deployment_dev_test.go`.
- **Proof shape**
  - Creates an explicit deployment-root fixture under generated temp evidence:
    - `bin/`
    - `config/`
    - `state/comms/`
    - `state/cypher/`
    - `state/sidecar/`
    - `logs/`
    - `evidence/`
    - `tmp/`
  - Writes a generated `gate-e.env.example` and `gate-e.deployment-context.json`.
  - Builds Cypher into the deployment-owned `bin/` root.
  - Starts Cypher only with explicit `CYPHER_ADDR`, `CYPHER_DB`, `CYPHER_MIGRATIONS`, and `CYPHER_DEV_INVITE`.
  - Health-checks Cypher over loopback.
  - Stops Cypher and restarts it against the same explicit deployment-root DB.
  - Runs C1/C2/C3/C4 state-policy surfaces over the deployment roots.
  - Verifies registry/reference freshness, missing nonclaims, Gate D authority, and Gate E E1 authority.
- **Validation posture**
  - Shadow and live gofmt passed.
  - Shadow and live command-reference render passed with `entries=129`.
  - Shadow and live diff checks passed across all four repos.
  - Shadow and live validator `go test ./... -count=1` passed.
  - Shadow and live `gate-e-native-deployment-dev --compact-summary --clean-generated` passed.
  - Shadow and live registry lookup for `runner.gate-e-native-deployment-dev` passed.
  - Shadow and live missing-nonclaims scan passed.
  - Light adjacent check confirmed no active `runner.full-runtime-dev` registry/dispatch promotion.
  - Light adjacent check confirmed no active service/systemd/container/TUI files were added by E1.
  - Final repo clean checks passed for all four repos.
- **Blunders / hazards preserved**
  - The first Gate E recon script attempted `go run ./cmd/cypher --help`, which started a blocking Cypher server because `cmd/cypher` is a server entrypoint, not a safe terminating CLI help surface.
  - Retry1 classified this as a real Gate E operator-surface hazard, not merely a script typo.
  - E1 therefore records `CYPHER_HELP_PROBE_USED=false` and keeps `CYPHER_SERVER_ENTRYPOINT_HAZARD_CLASSIFIED=true`.
  - Future semi-persistent service work should consider a terminating `--help`, `--print-config`, or `--check-config` surface before service/systemd work.
  - The recurring `carbonstack-cypher/cypher.db` artifact-scan hit remains review-only and non-destructive while final clean status passes; E1 proves future runtime DB generation belongs under declared deployment roots.
- **Nonclaims**
  - Gate E final closure is not claimed.
  - Semi-persistent deployment is not started.
  - Service/systemd/helper install is not started.
  - Public ingress, containers, TUI, `full-runtime-dev`, verified identity, trust promotion, vault/backup, production E2EE, PQ/hybrid, Android, CarbonStackOS, and v0.8.0 release readiness remain out of scope.
- **Safe handoff**
  - Gate E status is `open_e1_closed_e2_blocked`.
  - `GATE_E_E1_STATUS=closed`.
  - `GATE_E_E2_STATUS=not_started`.
  - Do not begin further Gate E subgates until the v0.7.19 LogDoc and Breakpoint are accepted.


---

## v0.7.20 — Gate E E2/E3 Cypher Config Inspection and Gate E Closure

- **Purpose**
  - Consolidate the post-v0.7.19 Gate E continuation into a durable breakpoint after completing E2 and E3.
  - Close the manual-private native deployment line before any Gate F preflight or additional native-deployment expansion.
- **E2 purpose**
  - Fix the Gate E recon hazard where `cmd/cypher --help` started the blocking server instead of terminating.
  - Add terminating Cypher operator inspection before runbook/helper/service decisions.
- **E2 meaningful work**
  - Added `--help`, `--print-config`, and `--check-config` behavior to `carbonstack-cypher/cmd/cypher`.
  - Added config-inspection JSON schema `carbonstack-cypher-config-inspection/v0`.
  - `--print-config` reports effective address, DB path, migrations path, env/default sources, dev invite enabled state without leaking the invite value, server-entrypoint posture, and `starts_server=false` / `terminating_inspection=true`.
  - `--check-config` validates address syntax, DB parent accessibility, migrations directory readability, and presence of SQL migrations.
  - Added registry ID `cypher.config-inspection` and docs `266/267`.
  - `carbonstack-cypher` advanced to `59c732ef53e198f65c03cc4e4178a66521c26a4c`.
  - `carbonstack` advanced to `d2acfafde32fa3f4b5ddd8652a3f9603d674c529` for E2 docs/registry.
  - Command reference entries became `130`.
- **E2 validation**
  - Shadow and live Cypher cmd tests passed.
  - Shadow and live full Cypher package tests passed.
  - Shadow and live `go run ./cmd/cypher --help` terminated without printing `CarbonStackCypher listening`.
  - Shadow and live `--print-config` emitted valid `carbonstack-cypher-config-inspection/v0` JSON and exited.
  - Shadow and live `--check-config` returned `config ok` for explicit Gate-E-style env and exited.
  - Registry lookup for `cypher.config-inspection` passed.
  - Missing nonclaims remained zero.
  - No Cypher process residue remained.
- **E3 purpose**
  - Add a Gate E closure profile and closure docs after E1 and E2, with breakpoint deferred by operator decision until E3 completion.
  - Close Gate E only for the manual-private native deployment model.
- **E3 meaningful work**
  - Added `gate-e-native-deployment-closure-dev` validator profile and test.
  - Added registry ID `runner.gate-e-native-deployment-closure-dev`.
  - Added docs `268-v0.7.19-gate-e-e3-native-deployment-closure-v0.md` and `269-v0.7.19-gate-e-closure-v0.md`.
  - The closure profile reruns Gate E E1 manual-private native deployment validation, verifies E1 and E2 registry authorities, verifies terminating `--print-config` and `--check-config`, and validates registry/reference/nonclaims.
  - `carbonstack` advanced to `52de931c8a5ec135fd26d01c4a370d37a8f7edfe`.
  - Command reference entries became `131`.
- **E3 validation**
  - Shadow and live gofmt passed.
  - Shadow and live command-reference render/check passed with `entries=131`.
  - Shadow and live validator tests passed.
  - Shadow and live `gate-e-native-deployment-closure-dev` passed.
  - Registry lookup for `runner.gate-e-native-deployment-closure-dev` passed.
  - Missing nonclaims remained zero.
  - Light adjacent check confirmed no active `full-runtime-dev` registry/dispatch promotion.
  - No Cypher process residue remained.
  - Final clean state passed for all four repositories.
- **Gate E closure meaning**
  - Gate E is closed for manual-private native deployment only.
  - Closure means the project has a dev/pre-alpha single-host native deployment proof model with explicit deployment root, explicit env, explicit loopback Cypher, explicit Cypher DB/migrations roots, explicit Comms state root, evidence/log/temp roots, Cypher start/stop/restart proof, terminating config inspection, Gate C state-policy inspection, and registry/nonclaim validation.
- **Nonclaims preserved**
  - No semi-persistent service/systemd/helper install.
  - No public ingress.
  - No containers.
  - No TUI.
  - No `full-runtime-dev` promotion.
  - No verified identity or trust promotion.
  - No vault/backup/restore claim.
  - No production E2EE/security claim.
  - No PQ/hybrid, Android, or CarbonStackOS implementation.
  - No v0.8.0 release readiness by Gate E alone.
- **Process/blunders continuity**
  - The code-box backtick issue caused the E3 script to need regeneration without grave-accent backticks; future long scripts should avoid backtick-heavy Markdown/codebox content when the user flags copy/paste fragility.
  - Operator intentionally deferred breakpoint until after E3; this is accepted for this narrow continuation because E3 was carbonstack-only closure over E1/E2, but future high-surface subgates should still checkpoint at accepted boundaries.
  - Repeated `carbonstack-cypher/cypher.db` artifact-scan hit remained review-only/non-destructive while final clean state passed; keep tracking as hygiene, not as a blocker.
  - E2 resolved the earlier `cypher --help` blocking-server hazard by making help/config inspection terminating.
- **Safe handoff**
  - Gate F is not started.
  - Next safe action is Gate F preflight/contract freeze, not implementation from memory.
  - Do not reopen Gate E expansion, semi-persistent service work, packaging, release readiness, or `full-runtime-dev` without fresh scope decision.
## v0.7.21 — Gate F F0/F1/F2 Release-Package Surface and Operator Docs / F2 Closed / F3 Blocked

- **Purpose**
  - Preserve the first Gate F checkpoint after Gate E closure.
  - Record Gate F0 recon, Gate F1 release/package/runtime surface classification, and Gate F2 operator docs/runbook closure.
  - Cut the requested breakpoint after F2, before any Gate F F3 implementation.
- **Gate F0 deep recon**
  - Ran read-only with no tracked-source mutation, no shadow mutation, no commit, no push, no release creation, no package publication, no `full-runtime-dev` promotion, no service/systemd/helper install, no TUI/container/public ingress, no migration, no destructive cleanup, and no LogDoc/Breakpoint generation.
  - Reconned release validation, release creation, package, runtime, and helper surfaces.
  - Source scan covered 454 files.
  - Source-scan category totals included:
    - `release_package`: 7024;
    - `operator_docs`: 3071;
    - `trust_identity`: 2271;
    - `migration_compat`: 1996;
    - `forbidden_scope`: 1683;
    - `code_health`: 1612;
    - `runtime_candidate`: 482;
    - `gate_f`: 471;
    - `pq`: 469.
  - Helper classification found 167 candidates.
  - Important F0 top surfaces included:
    - `carbonstack/registry/COMMAND_REFERENCE.v0.md`;
    - `carbonstack/registry/commands.v0.yaml`;
    - `carbonstack/registry/COMMAND_BOUNDARY_TABLE.v0.md`;
    - `carbonstack/scripts/stage-v0.7.0-package.sh`;
    - `carbonstack/docs/231-v0.6.33-release-boundary-package-root-rehearsal-plan-v0.md`;
    - `carbonstack/docs/227-v0.6.30-validation-profile-naming-boundary-v0.md`;
    - `carbonstack/tools/carbonstack-validate/README.md`;
    - `carbonstack/tools/carbonstack-validate/gate_e_native_deployment_dev.go`;
    - `carbonstack-comms/internal/app/state_substrate_inventory_dev.go`.
  - F0 confirmed `carbonstack/package`, `carbonstack/release`, and `carbonstack/dist` directories are not current package/runtime roots.
  - F0 concluded Gate F should not begin by releasing v0.8.0; it should first classify surfaces, close operator docs, then proceed to compatibility/refusal/rollback and package/runtime candidate work.
- **F0 script blunders / recovery**
  - First F0 script completed useful heavy recon, then failed while writing final compact JSON because `RECON_JSON` was not exported to embedded Python.
  - Retry script reused the completed F0 evidence and completed useful recovery checks, then failed due generated Python newline escaping in a JSON write line.
  - Finalizer correctly avoided repeating the heavy recon, summarized the already-complete artifacts, and confirmed no deep recon repetition was needed.
  - These were assistant script-emission bugs, not project semantic failures.
- **F1 release/package/runtime surface classification**
  - Added profile:
    - `gate-f-release-package-surface-dev`.
  - Added registry ID:
    - `runner.gate-f-release-package-surface-dev`.
  - Added schema:
    - `carbonstack-gate-f-release-package-surface-report/v0`.
  - Added docs:
    - `docs/270-v0.7.20-gate-f-f1-release-package-surface-v0.md`;
    - `docs/271-v0.7.20-gate-f-f1-closure-v0.md`.
  - Advanced generated command reference to `entries=132`.
  - Classified 113 surfaces:
    - 3 package staging scripts;
    - 3 package rehearsal scripts;
    - 1 release validation profile;
    - 2 runtime validation profiles;
    - 1 `cypher.db` hygiene hit.
  - Classified `full-validate-release` as release-package validation, not release creation.
  - Classified `release-snapshot` as package-root layout/checksum/core validation, not package creation or upload.
  - Classified `local-cypher` as release validation component, not deployment.
  - Classified `stage-v0.7.0-package.sh` as strongest historical package staging scaffold, not current v0.8.0 staging.
  - Kept manual v0.8.0 release creation as an operator process.
  - Did not execute package staging.
  - Did not publish a package.
  - Did not promote `full-runtime-dev`.
  - Did not implement migration.
  - Did not start service/systemd/helper/container/public ingress/TUI.
  - Committed `carbonstack`:
    - `4c99391598b478cd1f1d1d39d1af3f2110cf3821`
    - `test: add Gate F release package surface profile`
  - First push failed authentication and immediate retry succeeded.
- **F2 operator docs/runbook closure**
  - Added profile:
    - `gate-f-operator-docs-runbook-dev`.
  - Added registry ID:
    - `runner.gate-f-operator-docs-runbook-dev`.
  - Added schema:
    - `carbonstack-gate-f-operator-docs-runbook-report/v0`.
  - Added docs:
    - `docs/272-v0.7.20-gate-f-f2-manual-private-lifecycle-runbook-v0.md`;
    - `docs/273-v0.7.20-gate-f-f2-config-env-validation-guide-v0.md`;
    - `docs/274-v0.7.20-gate-f-f2-release-package-authority-guide-v0.md`;
    - `docs/275-v0.7.20-gate-f-f2-failure-refusal-hygiene-guide-v0.md`;
    - `docs/276-v0.7.20-gate-f-f2-closure-v0.md`.
  - Advanced generated command reference to `entries=133`.
  - F2 profile validated:
    - manual-private lifecycle runbook present;
    - config/env validation guide present;
    - release/package authority guide present;
    - failure/refusal/hygiene guide present;
    - F2 closure document present;
    - F1 and F2 registry authorities present;
    - release creation remains manual and unimplemented;
    - package staging was not executed;
    - `full-runtime-dev` remains unpromoted.
  - F2 closure state:
    - `GATE_F_STATUS=open_f1_f2_closed_f3_not_started`;
    - `GATE_F_F1_STATUS=closed`;
    - `GATE_F_F2_STATUS=closed`;
    - `GATE_F_F3_STATUS=not_started`;
    - `BREAKPOINT_REQUIRED_AFTER_F2=true`.
  - Committed `carbonstack`:
    - `5b2421387f4824b45a594a8993a1cd7d34bfa649`
    - `docs: add Gate F operator runbooks`
  - Push succeeded on first attempt.
- **Validation evidence**
  - F1 and F2 both passed shadow gofmt, render, diff checks across all four repos, validator tests, profile execution, generated-reference check, registry lookup, missing-nonclaims check, and adjacent boundary checks.
  - F1 and F2 both ended with final clean state across all four repositories.
  - F2 ended with no package/release artifacts in repo roots and no Cypher process residue.
- **Nonclaims preserved**
  - Not v0.8.0 release readiness.
  - Not release creation or upload.
  - Not package publication.
  - Not package staging execution.
  - Not package/runtime candidate.
  - Not `full-runtime-dev`.
  - Not migration implementation.
  - Not service/systemd/helper install.
  - Not public ingress.
  - Not container readiness.
  - Not TUI.
  - Not production readiness.
  - Not production E2EE.
  - Not verified identity.
  - Not trust promotion.
  - Not vault security.
  - Not backup/restore.
  - Not PQ/hybrid.
  - Not Android.
  - Not CarbonStackOS.
- **Safe handoff**
  - Generate this private LogDoc and lean Breakpoint before Gate F F3.
  - Gate F F3 should target compatibility, stale-state, rollback observability, and refusal posture.
  - Migration implementation, package/runtime candidate validation, and `full-runtime-dev` promotion remain unauthorized until explicitly contracted.
## v0.7.22 — Gate F F3 Compatibility/Rollback Observability Closure / F4 Blocked

- **Purpose**
  - Preserve the Gate F F3 closure checkpoint after F1 release/package/runtime surface classification and F2 operator docs/runbook closure.
  - Close a bounded aggregate observability/refusal posture for compatibility, stale-state, rollback/downgrade, Cypher migration visibility, and hygiene classification.
  - Keep Gate F open and block F4 until this LogDoc/Breakpoint handoff is accepted.
- **Preflight basis**
  - F3 preflight was read-only and completed with `PREFLIGHT_COMMAND_FAILURE_COUNT=0`.
  - Confirmed clean repositories, command reference current at `entries=133`, missing nonclaims zero, and correct F3 scope.
  - Targeted scan covered 43 files and identified the right F3 authority cluster:
    - C2 schema compatibility;
    - C3 path policy;
    - C4 write policy;
    - F1 release/package/runtime surface classification;
    - F2 operator docs/runbook authority;
    - Cypher config inspection and migration visibility;
    - recurring `cypher.db` hygiene evidence.
  - Preflight contract accepted F3 as observability/refusal only, not migration or package/runtime candidate work.
- **F3 implementation**
  - Added validator profile:
    - `gate-f-compat-rollback-observability-dev`.
  - Added registry ID:
    - `runner.gate-f-compat-rollback-observability-dev`.
  - Added schema:
    - `carbonstack-gate-f-compat-rollback-observability-report/v0`.
  - Added docs:
    - `docs/277-v0.7.21-gate-f-f3-compat-rollback-observability-v0.md`;
    - `docs/278-v0.7.21-gate-f-f3-closure-v0.md`.
  - Advanced generated command reference to `entries=134`.
  - Mutated only `carbonstack`.
  - Did not mutate `carbonstack-comms`, `carbonstack-cypher`, or `carbonstack-os`.
- **F3 profile behavior**
  - Aggregates and revalidates:
    - `state-schema-compat-dev`;
    - `state-path-policy-dev`;
    - `state-write-policy-dev`;
    - `gate-f-release-package-surface-dev`;
    - `gate-f-operator-docs-runbook-dev`;
    - generated command reference freshness;
    - F3 registry lookup;
    - missing-nonclaims scan.
  - Runs Cypher config inspection with explicit temporary env and requires:
    - schema `carbonstack-cypher-config-inspection/v0`;
    - `starts_server=false`;
    - `terminating_inspection=true`;
    - DB path source `env`.
  - Inventories Cypher migration SQL files without applying migrations.
  - Classifies `cypher.db` hits without deleting or moving them.
- **Validation evidence**
  - Shadow gofmt, command-reference render, diff checks across all four repos, validator tests, F3 profile, generated-reference check, F3 registry lookup, and missing-nonclaims scan all passed.
  - Live gofmt, command-reference render, diff checks across all four repos, validator tests, F3 profile, generated-reference check, F3 registry lookup, and missing-nonclaims scan all passed.
  - F3 profile showed:
    - C2/C3/C4 profiles pass;
    - F1/F2 profiles pass;
    - Cypher config inspection terminates and reports explicit env;
    - Cypher migration inventory observed without migration implementation;
    - registry/reference/nonclaim checks pass;
    - migration, repair, destructive cleanup, package/runtime candidate, and `full-runtime-dev` remain unimplemented;
    - `cypher_migration_count=4`;
    - `cypher_db_hit_count=1`.
  - Adjacent boundary checks passed:
    - no active `full-runtime-dev` registry/dispatch promotion;
    - no package/release archive artifacts created in repo roots;
    - no Cypher process residue after F3 profile.
  - Final clean state passed for all four repositories.
- **Commit**
  - `carbonstack` advanced to:
    - `696ed532b8754feea20582a747803e8f1540be6a`
    - `test: add Gate F compatibility observability profile`
  - Commit included 9 changed files, 633 insertions, 3 deletions.
  - Added new files:
    - `docs/277-v0.7.21-gate-f-f3-compat-rollback-observability-v0.md`;
    - `docs/278-v0.7.21-gate-f-f3-closure-v0.md`;
    - `tools/carbonstack-validate/gate_f_compat_rollback_observability_dev.go`;
    - `tools/carbonstack-validate/gate_f_compat_rollback_observability_dev_test.go`.
  - First push failed authentication, immediate retry succeeded.
- **Nonclaims preserved**
  - Not v0.8.0 release readiness.
  - Not release creation or upload.
  - Not package publication.
  - Not package staging execution.
  - Not package/runtime candidate.
  - Not `full-runtime-dev`.
  - Not migration implementation.
  - Not silent migration.
  - Not repair implementation.
  - Not silent repair.
  - Not destructive cleanup.
  - Not state relocation.
  - Not service/systemd/helper install.
  - Not public ingress.
  - Not containers.
  - Not TUI.
  - Not production readiness.
  - Not production E2EE.
  - Not verified identity.
  - Not trust promotion.
  - Not vault security.
  - Not backup/restore.
  - Not PQ/hybrid.
  - Not Android.
  - Not CarbonStackOS.
- **Blunders / continuity**
  - No new implementation blunder in F3 itself.
  - Routine infrastructure issue repeated: first push failed authentication and retry succeeded.
  - The F3 profile intentionally reran C2/C3/C4/F1/F2. This is acceptable for an aggregate closure profile, but should not become the default for narrower F4 work.
  - The recurring `cypher.db` hit remained visible and non-destructive.
- **Safe handoff**
  - Gate F is open with F1/F2/F3 closed.
  - Gate F F4 is not started.
  - Next likely subgate: F4 code-health/source-hygiene closure, including a non-destructive `cypher.db` policy if accepted.
  - Do not start package/runtime candidate, migration, `full-runtime-dev`, release creation, service/systemd/helper, TUI, container/public ingress, verified identity/trust, vault/backup, PQ, Android, or OS work without a fresh explicit contract.
## v0.7.23 — Gate F F4 Code-Health / Source-Hygiene Closure / F5 Blocked

- **Purpose**
  - Preserve the Gate F F4 closure checkpoint after F3 compatibility/rollback observability.
  - Close a bounded code-health and source-hygiene classification layer before package/runtime candidate work.
  - Keep Gate F open and block F5 until this LogDoc/Breakpoint handoff is accepted.
- **F4 preflight basis**
  - F4 preflight was read-only and completed with `PREFLIGHT_COMMAND_FAILURE_COUNT=0`.
  - Confirmed clean starting repositories, command reference current at `entries=134`, missing nonclaims zero, and correct F4 scope.
  - Targeted code-health/source-hygiene scan found 565 tracked files with hits across scanned categories.
  - Preflight category totals included:
    - `release_helper`: 14775;
    - `generated_artifact`: 8882;
    - `panic_fatal_exit`: 2606;
    - `migration_repair`: 1706;
    - `service_container_tui`: 629;
    - `unsafe_or_hack`: 442;
    - `full_runtime`: 373;
    - `destructive_terms`: 357;
    - `cypher_db`: 257;
    - `todo_fixme`: 10.
  - Preflight correctly showed historical sanitized LogDocs dominate keyword scans and should be classified as historical provenance rather than current source defects.
  - Preflight found a generated local Python bytecode cache after registry rendering:
    - `carbonstack/tools/registry/__pycache__/render-command-reference.cpython-313.pyc`.
  - That generated cache was classified as a source-hygiene issue and not a project semantic failure.
- **F4 implementation**
  - Added targeted generated Python cache cleanup before the clean guard.
  - Deleted only the generated pycache file created by F4 preflight and removed the empty pycache directory if possible.
  - Added `.gitignore` coverage for Python generated caches:
    - `__pycache__/`;
    - `*.pyc`;
    - `*.pyo`.
  - Added validator profile:
    - `gate-f-code-health-source-hygiene-dev`.
  - Added registry ID:
    - `runner.gate-f-code-health-source-hygiene-dev`.
  - Added schema:
    - `carbonstack-gate-f-code-health-source-hygiene-report/v0`.
  - Added docs:
    - `docs/279-v0.7.22-gate-f-f4-code-health-source-hygiene-v0.md`;
    - `docs/280-v0.7.22-gate-f-f4-closure-v0.md`.
  - Advanced generated command reference to `entries=135`.
  - Mutated only `carbonstack`.
  - Did not mutate `carbonstack-comms`, `carbonstack-cypher`, or `carbonstack-os`.
- **F4 profile behavior**
  - Classifies source hygiene across tracked text/source/docs/config/script files.
  - Classifies helper/release/package surfaces statically without executing release actions.
  - Syntax-checks shell helpers with `bash -n`.
  - Syntax-checks Python helpers with Python bytecode disabled.
  - Confirms Python cache ignore policy is present.
  - Classifies repo-root `cypher.db` as non-destructive hygiene evidence.
  - Classifies `sanitized-project-logdoc-list` as historical provenance rather than current source defect evidence.
  - Runs validator package tests.
  - Confirms command-reference freshness.
  - Confirms F4 registry lookup.
  - Confirms missing nonclaims zero.
- **Validation evidence**
  - Shadow gofmt, command-reference render, diff checks across all four repos, validator tests, F4 profile, generated-reference check, F4 registry lookup, and missing-nonclaims scan all passed.
  - Live gofmt, command-reference render, diff checks across all four repos, validator tests, F4 profile, generated-reference check, F4 registry lookup, and missing-nonclaims scan all passed.
  - F4 profile reported:
    - `tracked_files_with_hygiene_hits=564`;
    - `helper_candidate_count=179`;
    - `shell_syntax_checked=9`;
    - `python_syntax_checked=1`;
    - `cypher_db_hit_count=1`;
    - `gate_f_status=open_f1_f2_f3_f4_closed_f5_not_started`;
    - `gate_f_f4_status=closed`;
    - `gate_f_f5_status=not_started`.
  - Adjacent boundary checks passed:
    - no active `full-runtime-dev` registry/dispatch promotion;
    - no package/release archive artifacts created in repo roots;
    - no Cypher process residue after F4 profile.
  - Final clean state passed for all four repositories.
- **Commit**
  - `carbonstack` advanced to:
    - `53b5ce56ab7eacf91762a3892b0d41bdbe3cc8ae`
    - `test: add Gate F source hygiene profile`
  - Commit included 10 changed files, 738 insertions, 3 deletions.
  - Added new files:
    - `docs/279-v0.7.22-gate-f-f4-code-health-source-hygiene-v0.md`;
    - `docs/280-v0.7.22-gate-f-f4-closure-v0.md`;
    - `tools/carbonstack-validate/gate_f_code_health_source_hygiene_dev.go`;
    - `tools/carbonstack-validate/gate_f_code_health_source_hygiene_dev_test.go`.
  - Push succeeded on first attempt.
- **Nonclaims preserved**
  - Not v0.8.0 release readiness.
  - Not release creation or upload.
  - Not package publication.
  - Not package staging execution.
  - Not package/runtime candidate.
  - Not `full-runtime-dev`.
  - Not migration implementation.
  - Not repair implementation.
  - Not destructive cleanup of state.
  - Not state relocation.
  - Not service/systemd/helper install.
  - Not public ingress.
  - Not containers.
  - Not TUI.
  - Not production readiness.
  - Not production E2EE.
  - Not verified identity.
  - Not trust promotion.
  - Not vault security.
  - Not backup/restore.
  - Not PQ/hybrid.
  - Not Android.
  - Not CarbonStackOS.
- **Blunders / continuity**
  - F4 preflight left an untracked Python bytecode cache in the repo. This was a tool/script side effect and a real hygiene issue, not a source semantic defect.
  - F4 implementation recovered correctly by targeted generated-cache cleanup before the clean guard and then adding a `.gitignore` policy to prevent recurrence.
  - This cleanup was not state cleanup and did not touch `cypher.db`, evidence roots, release artifacts, or runtime state.
  - F4 stayed appropriately narrower than F3: it did not rerun the full C2/C3/C4/F1/F2/F3 aggregate ladder.
  - The recurring `cypher.db` hit remained visible and non-destructive.
- **Safe handoff**
  - Gate F is open with F1/F2/F3/F4 closed.
  - Gate F F5 is not started.
  - Next likely subgate: F5 package/runtime candidate validation preflight.
  - Do not start package/runtime candidate implementation, release creation/upload, package publication/staging, `full-runtime-dev`, migration, repair, destructive cleanup, state relocation, service/systemd/helper, TUI, container/public ingress, verified identity/trust, vault/backup, PQ, Android, or OS work without a fresh explicit contract.
## v0.7.24 — Gate F F5 Basic Local Trust Candidate Posture Closure / F6 Blocked

- **Purpose**
  - Preserve the Gate F F5 closure checkpoint after the deliberate operator-approved scope pivot from package/runtime candidate validation to a smaller pre-release identity/trust posture subgate.
  - Close a **basic local manual trust candidate posture** before package/runtime candidate work.
  - Keep the release claim honest:
    - basic local trust posture exists;
    - manual local acceptance evidence path exists;
    - verified identity remains a nonclaim;
    - full trust promotion remains a nonclaim;
    - secure enrollment remains a nonclaim;
    - cryptographic binding across Cypher, Comms, and OpenMLS remains a nonclaim;
    - automatic trust promotion remains forbidden.
  - Keep Gate F open and block F6 until this LogDoc/Breakpoint handoff is accepted.
- **F5 preflight basis**
  - F5 preflight was read-only and completed with `PREFLIGHT_COMMAND_FAILURE_COUNT=0`.
  - Confirmed clean starting repositories:
    - `carbonstack` at `53b5ce56ab7eacf91762a3892b0d41bdbe3cc8ae`;
    - `carbonstack-comms` at `25d8876a20dc2e545f87c400e51c9842dc0e3f31`;
    - `carbonstack-cypher` at `59c732ef53e198f65c03cc4e4178a66521c26a4c`;
    - `carbonstack-os` at `1bbbe52020d623b81796694e5057c1d080ede3ea`.
  - Confirmed command reference current at `entries=135` and missing nonclaims zero.
  - Confirmed existing F1/F2/F3/F4 registry IDs and Gate C state registry IDs remained healthy.
  - Trust/identity scan confirmed existing project posture:
    - Relay membership is routing authority only;
    - KeyPackage/Welcome lifecycle surfaces repeatedly nonclaim identity verification and trust promotion;
    - current message wrappers already carry strict trust/fingerprint language but do not close verified identity;
    - Comms, Cypher, OpenMLS, and Relay identity/material domains remain separate.
  - Identity-domain inventory split the posture into:
    - Cypher account/device domain;
    - Comms local trust/candidate domain;
    - OpenMLS signer/KeyPackage domain;
    - Relay membership domain.
  - Preflight contract accepted F5 as basic local trust candidate posture only, not verified identity or full trust promotion.
- **F5 implementation**
  - Mutated `carbonstack-comms` and `carbonstack` only.
  - Did not mutate `carbonstack-cypher` or `carbonstack-os`.
  - Added Comms dev commands:
    - `basic-local-trust-posture-dev`;
    - `basic-local-trust-accept-dev`.
  - Added Comms schemas:
    - `carbonstack-basic-local-trust-posture/v0`;
    - `carbonstack-basic-local-trust-acceptance-event/v0`;
    - `carbonstack-basic-local-trust-acceptance-command-result/v0`.
  - Added CarbonStack validator profile:
    - `gate-f-basic-local-trust-posture-dev`.
  - Added CarbonStack report schema:
    - `carbonstack-gate-f-basic-local-trust-posture-report/v0`.
  - Added registry IDs:
    - `comms.basic-local-trust-posture-dev`;
    - `comms.basic-local-trust-accept-dev`;
    - `runner.gate-f-basic-local-trust-posture-dev`.
  - Added Comms docs:
    - `carbonstack-comms/docs/31-gate-f-f5-basic-local-trust-posture-v0.md`;
    - `carbonstack-comms/docs/32-gate-f-f5-basic-local-trust-closure-v0.md`.
  - Added CarbonStack docs:
    - `carbonstack/docs/281-v0.7.23-gate-f-f5-basic-local-trust-posture-v0.md`;
    - `carbonstack/docs/282-v0.7.23-gate-f-f5-closure-v0.md`.
  - Advanced generated command reference to `entries=138`.
- **F5 behavior**
  - `basic-local-trust-posture-dev` produces a JSON posture report over four identity domains:
    - Cypher account/device identity: coordination and routing identity only;
    - Comms local trust/candidate fingerprint: local operator policy evidence only;
    - OpenMLS sidecar device label / KeyPackage ref: cryptographic group material, not verified real-world identity;
    - Relay membership: routing/membership coordination only, never trust promotion.
  - `basic-local-trust-accept-dev` writes an explicit local manual trust candidate acceptance event.
  - Acceptance requires:
    - `--accept-candidate`;
    - `--reason`;
    - subject label;
    - Cypher account;
    - Cypher device;
    - Comms fingerprint.
  - Acceptance may include:
    - OpenMLS sidecar device label;
    - OpenMLS KeyPackage reference;
    - Relay Space;
    - source report path.
  - Missing `--accept-candidate` is a hard failure.
  - Acceptance event records nonclaims and does not silently promote trust.
- **Validation evidence**
  - Shadow gofmt, generated-reference render, diff checks across all four repos, Comms focused tests, Comms package tests, CarbonStack validator tests, F5 profile, generated-reference check, registry lookups for all F5 IDs, and missing-nonclaims scan all passed.
  - Live gofmt, generated-reference render, diff checks across all four repos, Comms focused tests, Comms package tests, CarbonStack validator tests, F5 profile, generated-reference check, registry lookups for all F5 IDs, and missing-nonclaims scan all passed.
  - F5 profile proved:
    - posture report generated;
    - missing `--accept-candidate` fails as expected;
    - manual local acceptance event writes to evidence state;
    - verified identity remains false;
    - trust promotion remains false;
    - cryptographic identity binding remains false;
    - automatic trust promotion remains false;
    - registry/reference/nonclaim checks pass.
  - Adjacent boundary checks passed:
    - no active `full-runtime-dev` registry/dispatch promotion;
    - no package/release archive artifacts created in repo roots;
    - no Cypher process residue after F5 profile.
  - Final clean state passed for all four repositories.
- **Commits**
  - `carbonstack-comms` advanced to:
    - `32681e20784e560fdc8075bbc8a6597b742823c9`
    - `feat: add Gate F basic local trust posture`
  - `carbonstack` advanced to:
    - `d053249b78d84ec948b0ffc71f1fd530580142e8`
    - `test: add Gate F basic local trust profile`
  - `carbonstack-cypher` remained:
    - `59c732ef53e198f65c03cc4e4178a66521c26a4c`
  - `carbonstack-os` remained:
    - `1bbbe52020d623b81796694e5057c1d080ede3ea`
  - `carbonstack-comms` commit included 6 changed files, 515 insertions.
  - `carbonstack` commit included 9 changed files, 634 insertions, 4 deletions.
  - `carbonstack-comms` first push failed authentication and immediate retry succeeded.
  - `carbonstack` push succeeded on first attempt.
- **Nonclaims preserved**
  - Not verified identity.
  - Not full trust promotion.
  - Not secure enrollment.
  - Not server-hostile identity replacement proof.
  - Not real-world person verification.
  - Not cryptographic binding across Cypher, Comms, and OpenMLS identities.
  - Not automatic trust promotion.
  - Not trust from Relay membership.
  - Not trust from successful Welcome or MLS join.
  - Not production E2EE.
  - Not v0.8.0 release readiness.
  - Not release creation or upload.
  - Not package publication.
  - Not package staging execution.
  - Not package/runtime candidate.
  - Not `full-runtime-dev`.
  - Not migration implementation.
  - Not repair implementation.
  - Not destructive cleanup.
  - Not state relocation.
  - Not service/systemd/helper install.
  - Not public ingress.
  - Not containers.
  - Not TUI.
  - Not vault security.
  - Not backup/restore.
  - Not PQ/hybrid.
  - Not Android.
  - Not CarbonStackOS.
- **Blunders / continuity**
  - F5 implementation itself had no semantic failure.
  - Recurring infrastructure noise repeated: `carbonstack-comms` first push failed authentication and immediate retry succeeded.
  - The F5 sequencing intentionally displaced the previous v0.7.23 “F5 package/runtime candidate” plan. Package/runtime candidate validation is now F6.
  - F5 avoided the dangerous overclaim:
    - “basic local trust posture exists” is now allowed;
    - “verified identity exists” remains false;
    - “trust promotion exists” remains false.
  - F5 correctly did not touch Cypher schema, sidecar provider internals, Android/OS, PQ, deployment helpers, or packaging.
- **Safe handoff**
  - Gate F is open with F1/F2/F3/F4/F5 closed.
  - Gate F F6 is not started.
  - Next likely subgate: F6 package/runtime candidate validation preflight.
  - Do not start package/runtime candidate implementation, release creation/upload, package publication/staging, `full-runtime-dev`, migration, repair, destructive cleanup, state relocation, service/systemd/helper, TUI, container/public ingress, verified identity/full trust promotion, secure enrollment, cryptographic identity binding, vault/backup, PQ, Android, or OS work without a fresh explicit contract.
## v0.7.25 — Gate F F6 Package/Runtime Candidate Validation Closure / F7 Blocked

- **Purpose**
  - Preserve the Gate F F6 closure checkpoint after F5 basic local trust candidate posture.
  - Close a bounded **package/runtime candidate validation** layer before v0.8.0 release-candidate closure.
  - Prove a disposable candidate root shape and root-separation policy without creating, uploading, publishing, staging, or promoting a public/runtime release surface.
  - Keep Gate F open and block F7 until this LogDoc/Breakpoint handoff is accepted.
- **F6 preflight basis**
  - F6 preflight was read-only and completed with `PREFLIGHT_COMMAND_FAILURE_COUNT=0`.
  - Confirmed clean repositories and command reference current at `entries=138`.
  - Confirmed missing nonclaims zero.
  - Confirmed F1/F2/F3/F4/F5 registry surfaces remained healthy.
  - Package/runtime scan found a dense historical package/release surface and validated that F6 should treat stage/rehearse scripts as prior art rather than execute them.
  - Runtime-root inventory defined the intended separation between:
    - candidate package root;
    - release artifact root;
    - manual-private runtime deployment root;
    - evidence root;
    - generated report root;
    - local runtime state root.
  - Preflight contract accepted F6 as package/runtime candidate validation only, not release creation, package publication, legacy package staging, or `full-runtime-dev`.
- **F6 implementation**
  - Mutated `carbonstack` only.
  - Did not mutate `carbonstack-comms`, `carbonstack-cypher`, or `carbonstack-os`.
  - Added validator profile:
    - `gate-f-package-runtime-candidate-dev`.
  - Added registry ID:
    - `runner.gate-f-package-runtime-candidate-dev`.
  - Added report schema:
    - `carbonstack-gate-f-package-runtime-candidate-report/v0`.
  - Added package candidate manifest schema:
    - `carbonstack-package-runtime-candidate-manifest/v0`.
  - Added docs:
    - `docs/283-v0.7.24-gate-f-f6-package-runtime-candidate-v0.md`;
    - `docs/284-v0.7.24-gate-f-f6-closure-v0.md`.
  - Advanced generated command reference to `entries=139`.
- **F6 behavior**
  - Builds and validates a disposable candidate root under temp/evidence-style space.
  - Separates:
    - candidate root;
    - package root;
    - runtime deployment root;
    - release artifact root;
    - evidence root;
    - report root;
    - local state root.
  - Writes a candidate manifest:
    - `/tmp/carbonstack-gate-f-package-runtime-candidate-dev/candidate-root/package-root/manifest/package-runtime-candidate-manifest.json`.
  - Writes a profile report:
    - `/tmp/carbonstack-gate-f-package-runtime-candidate-dev/gate-f-package-runtime-candidate-report.json`.
  - Invokes the F5 basic local trust posture command as a nonclaim-aware dependency and verifies:
    - verified identity remains false;
    - trust promotion remains false;
    - automatic trust promotion remains false;
    - cryptographic identity binding remains false.
  - Performs Cypher config inspection with explicit environment and terminating inspection.
  - Classifies legacy stage/rehearse scripts as prior art, not executed behavior.
  - Confirms candidate root contains no public archive package artifacts.
  - Confirms active `full-runtime-dev` promotion remains absent.
- **Validation evidence**
  - Shadow gofmt, command-reference render, diff checks across all four repos, validator tests, F6 profile, generated-reference check, F6 registry lookup, and missing-nonclaims scan all passed.
  - Live gofmt, command-reference render, diff checks across all four repos, validator tests, F6 profile, generated-reference check, F6 registry lookup, and missing-nonclaims scan all passed.
  - F6 profile reported:
    - disposable package/runtime candidate root validated;
    - package/release/runtime/evidence/report/local-state roots separated;
    - candidate manifest written;
    - F5 trust posture dependency preserved nonclaims;
    - Cypher config inspection terminates with explicit env;
    - legacy stage/rehearse scripts classified but not executed;
    - no public package archive artifacts created;
    - `full-runtime-dev` remains unpromoted;
    - registry/reference/nonclaim checks passed;
    - `stage_rehearse_script_count=84`;
    - `gate_f_status=open_f1_f2_f3_f4_f5_f6_closed_f7_not_started`.
  - Adjacent boundary checks passed:
    - no active `full-runtime-dev` registry/dispatch promotion;
    - no package/release archive artifacts created in repo roots;
    - no Cypher process residue after F6 profile.
  - Final clean state passed for all four repositories.
- **Commit**
  - `carbonstack` advanced to:
    - `2f76f9b25f00804e72070d70f22e09b2f7c4fbaa`
    - `test: add Gate F package runtime candidate profile`
  - Commit included 9 changed files, 784 insertions, 3 deletions.
  - Added new files:
    - `docs/283-v0.7.24-gate-f-f6-package-runtime-candidate-v0.md`;
    - `docs/284-v0.7.24-gate-f-f6-closure-v0.md`;
    - `tools/carbonstack-validate/gate_f_package_runtime_candidate_dev.go`;
    - `tools/carbonstack-validate/gate_f_package_runtime_candidate_dev_test.go`.
  - Push succeeded on first attempt.
  - `carbonstack-comms`, `carbonstack-cypher`, and `carbonstack-os` remained unchanged.
- **Nonclaims preserved**
  - Not v0.8.0 release readiness by itself.
  - Not release creation.
  - Not release upload.
  - Not package publication.
  - Not package staging execution.
  - Not public package artifact creation.
  - Not `full-runtime-dev`.
  - Not service/systemd.
  - Not helper install.
  - Not public ingress.
  - Not container readiness.
  - Not TUI.
  - Not migration implementation.
  - Not repair implementation.
  - Not destructive cleanup.
  - Not state relocation.
  - Not verified identity.
  - Not full trust promotion.
  - Not secure enrollment.
  - Not cryptographic identity binding.
  - Not vault security.
  - Not backup restore.
  - Not production E2EE.
  - Not PQ/hybrid.
  - Not Android.
  - Not CarbonStackOS.
- **Blunders / continuity**
  - F6 implementation itself had no semantic or script failure.
  - Push succeeded on first attempt; no auth retry noise in F6.
  - F6 successfully avoided executing legacy package staging scripts, which is important because recon found many historical release/package surfaces and danger terms.
  - F6 did not promote `full-runtime-dev`; this remains correct until a later explicit contract proves a coherent runtime aggregate.
  - F6 used F5 trust posture only as a nonclaim-aware dependency and did not accidentally upgrade basic local trust posture into verified identity or trust promotion.
  - F6 did not create a public archive/package artifact in repo roots.
- **Safe handoff**
  - Gate F is open with F1/F2/F3/F4/F5/F6 closed.
  - Gate F F7 is not started.
  - Next likely subgate: F7 v0.8.0 release-candidate closure matrix and manual release handoff.
  - Do not start release creation/upload, package publication/staging, `full-runtime-dev`, migration, repair, destructive cleanup, state relocation, service/systemd/helper, TUI, container/public ingress, verified identity/full trust promotion, secure enrollment, cryptographic identity binding, vault/backup, PQ, Android, or OS work without a fresh explicit contract.
## v0.7.26 — Gate F F7 Release-Candidate Closure Matrix / Manual Release Handoff Ready

- **Purpose**
  - Preserve the Gate F F7 closure checkpoint after F6 package/runtime candidate validation.
  - Close the bounded **v0.8.0 release-candidate closure matrix and manual release handoff** layer.
  - Move Gate F from open to **closed as v0.8.0 release-candidate handoff ready**, without performing the public release action.
  - Preserve a safe breakpoint before operator-controlled v0.8.0 release preparation.
- **F7 preflight basis**
  - F7 preflight was read-only and completed with `PREFLIGHT_COMMAND_FAILURE_COUNT=0`.
  - Confirmed clean repositories and command reference current at `entries=139`.
  - Confirmed missing nonclaims zero.
  - Confirmed the closure matrix had every required authority present in registry, command reference, and boundary table:
    - Gate B closure profile;
    - Gate C closure profile;
    - Gate D runtime aggregate profile;
    - Gate E native deployment closure profile;
    - Gate F F1 release/package/runtime surface profile;
    - Gate F F2 operator docs/runbook profile;
    - Gate F F3 compatibility/rollback observability profile;
    - Gate F F4 code-health/source-hygiene profile;
    - Gate F F5 basic local trust posture profile;
    - Gate F F6 package/runtime candidate validation profile;
    - release package validation profile;
    - release snapshot validation profile;
    - Cypher config inspection.
  - Release-surface inventory found a large historical release/handoff surface, including prior release ledgers and registry/reference files.
  - F7 preflight correctly classified historical release-package/action terms as evidence/prior-art to be summarized into handoff/checklist, not executed.
  - F7 contract accepted a `carbonstack`-only implementation unless recon found a bounded need elsewhere; no such need was found.
- **F7 implementation**
  - Mutated `carbonstack` only.
  - Did not mutate `carbonstack-comms`, `carbonstack-cypher`, or `carbonstack-os`.
  - Added validator profile:
    - `gate-f-release-candidate-closure-dev`.
  - Added registry ID:
    - `runner.gate-f-release-candidate-closure-dev`.
  - Added report schema:
    - `carbonstack-gate-f-release-candidate-closure-report/v0`.
  - Added manual handoff schema:
    - `carbonstack-v0.8.0-manual-release-handoff/v0`.
  - Added docs:
    - `docs/285-v0.7.25-gate-f-f7-release-candidate-closure-v0.md`;
    - `docs/286-v0.7.25-gate-f-f7-manual-release-handoff-v0.md`;
    - `docs/287-v0.7.25-gate-f-f7-closure-v0.md`.
  - Advanced generated command reference to `entries=140`.
- **F7 behavior**
  - Generates and validates a release-candidate closure report under temp validation output:
    - `/tmp/carbonstack-gate-f-release-candidate-closure-dev/gate-f-release-candidate-closure-report.json`.
  - Generates manual release handoff JSON:
    - `/tmp/carbonstack-gate-f-release-candidate-closure-dev/v0.8.0-manual-release-handoff.json`.
  - Generates release notes scaffold:
    - `/tmp/carbonstack-gate-f-release-candidate-closure-dev/v0.8.0-release-notes-scaffold.md`.
  - Generates manual release checklist:
    - `/tmp/carbonstack-gate-f-release-candidate-closure-dev/v0.8.0-manual-release-checklist.md`.
  - Requires closure matrix coverage across Gate B, Gate C, Gate D, Gate E, and Gate F F1-F6.
  - Re-runs F6 package/runtime candidate dependency during F7 validation.
  - Verifies F5 basic local trust posture remains nonclaim-aware through the F6 dependency.
  - Confirms generated command reference is current.
  - Runs registry lookups for every required closure authority.
  - Confirms missing nonclaims zero.
  - Confirms no public package/archive artifacts in repo roots.
  - Confirms `full-runtime-dev` remains unpromoted.
- **Validation evidence**
  - Shadow gofmt, command-reference render, diff checks across all four repos, validator tests, F7 profile, generated-reference check, F7 registry lookup, and missing-nonclaims scan all passed.
  - Live gofmt, command-reference render, diff checks across all four repos, validator tests, F7 profile, generated-reference check, F7 registry lookup, and missing-nonclaims scan all passed.
  - F7 profile reported:
    - v0.8.0 release-candidate closure matrix generated;
    - Gate B/C/D/E and Gate F F1-F6 authority surfaces present;
    - F6 package/runtime candidate dependency passed;
    - generated command reference current;
    - registry lookup and missing-nonclaims checks passed;
    - manual release handoff JSON written;
    - release notes scaffold written;
    - manual release checklist written;
    - no public package/archive artifacts in repo roots;
    - `full-runtime-dev` remains unpromoted;
    - `gate_f_status=closed_v0_8_0_release_candidate_handoff_ready`;
    - `v0_8_0_release_candidate_handoff_ready=true`.
  - Adjacent boundary checks passed:
    - no active `full-runtime-dev` registry/dispatch promotion;
    - no package/release archive artifacts created in repo roots;
    - no Cypher process residue after F7 profile.
  - Final clean state passed for all four repositories.
- **Commit**
  - `carbonstack` advanced to:
    - `715fd76f7700e63f1de877b730c2ce6bb37225bf`
    - `test: add Gate F release candidate closure profile`
  - Commit included 10 changed files, 847 insertions, 3 deletions.
  - Added new files:
    - `docs/285-v0.7.25-gate-f-f7-release-candidate-closure-v0.md`;
    - `docs/286-v0.7.25-gate-f-f7-manual-release-handoff-v0.md`;
    - `docs/287-v0.7.25-gate-f-f7-closure-v0.md`;
    - `tools/carbonstack-validate/gate_f_release_candidate_closure_dev.go`;
    - `tools/carbonstack-validate/gate_f_release_candidate_closure_dev_test.go`.
  - First push failed authentication and immediate retry succeeded.
  - `carbonstack-comms`, `carbonstack-cypher`, and `carbonstack-os` remained unchanged.
- **Nonclaims preserved**
  - Not release creation.
  - Not release upload.
  - Not package publication.
  - Not package staging execution.
  - Not public package artifact creation.
  - Not `full-runtime-dev`.
  - Not service/systemd.
  - Not helper install.
  - Not public ingress.
  - Not container readiness.
  - Not TUI.
  - Not migration implementation.
  - Not repair implementation.
  - Not destructive cleanup.
  - Not state relocation.
  - Not verified identity.
  - Not full trust promotion.
  - Not secure enrollment.
  - Not cryptographic identity binding.
  - Not vault security.
  - Not backup restore.
  - Not production E2EE.
  - Not PQ/hybrid.
  - Not Android.
  - Not CarbonStackOS.
- **Blunders / continuity**
  - F7 implementation itself had no semantic failure.
  - Script-generation blunder before operator run:
    - the first generated artifact draft accidentally contained backticks in Go struct tags/raw string literal syntax;
    - the local no-backtick assertion failed before any user-side WSL execution;
    - the artifact was corrected and regenerated with backtick count zero before delivery;
    - no repository or WSL state was affected.
  - Recurring infrastructure noise repeated:
    - `carbonstack` first push failed authentication;
    - immediate retry succeeded.
  - F7 kept release handoff separate from release action, preventing the overclaim that v0.8.0 was already created/uploaded.
  - F7 used F6/F5 dependencies as validation inputs without promoting package staging, `full-runtime-dev`, verified identity, or trust promotion.
- **Safe handoff**
  - Gate F is closed as `closed_v0_8_0_release_candidate_handoff_ready`.
  - v0.8.0 release is still not created/uploaded/published.
  - Next likely action after accepting this breakpoint:
    - operator-controlled v0.8.0 release preparation using prior release assets.
  - Required operator inputs remain:
    - previous release markdown scaffold;
    - previous release visual assets;
    - previous attached-file component list;
    - target Gitea release link or release creation surface;
    - explicit confirmation to create and upload public release assets.
  - After public v0.8.0 release:
    - capture public release URL;
    - generate post-release PRIME continuity artifact;
    - depreciate/deprecate the v0.8.0 EVERGREEN roadmap;
    - start v0.9.x adversarial roadmap Q/A.
  - Do not start public release creation/upload, package publication/staging, `full-runtime-dev`, migration, repair, destructive cleanup, state relocation, service/systemd/helper, TUI, container/public ingress, verified identity/full trust promotion, secure enrollment, cryptographic identity binding, vault/backup, PQ, Android, or OS work without explicit operator confirmation.
## v0.8.0 PRIME — Public Release Posted / v0.7.x Epoch Closed / Evergreen Deprecation Pending

- **Purpose**
  - Preserve the final post-release continuity state after the operator reported the v0.8.0 public pre-release was posted on Gitea.
  - Convert the v0.7.26 private release-candidate handoff into a v0.8.0 PRIME public-release continuity anchor.
  - Mark the v0.7.x operational-spine integration epoch as closed.
  - Prepare the handoff into the next roadmap-generation step.
- **Operator-reported release state**
  - The v0.8.0 release assets and tag draft were completed and manually uploaded.
  - The release body was drafted from the v0.7.0 public release template, then corrected to more closely match the v0.7.0 structure and a semi-public release tone.
  - The operator accepted the corrected release body and reported: “release posted.”
  - Provided release index:
    - `https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases`
- **Release asset preparation continuity**
  - Initial release asset prep script failed after staging because the package-internal checksum path did not match the validator’s canonical expectation.
  - Root cause:
    - script wrote the internal checksum file using an external versioned asset-style name;
    - validator expected canonical package-internal `release/checksums.txt`.
  - Correction:
    - v2 release asset prep wrote package-internal `release/manifest.json`;
    - generated package-internal `release/checksums.txt` using the project validator;
    - copied internal metadata/checksums into external versioned upload asset names;
    - fresh extraction validation used `verify-checksums` and `full-validate-release`;
    - F7 release-candidate closure remained live prepackage evidence, not a fresh-package validation command.
  - Uploaded asset shape matched the prior v0.7.0 family:
    - package archive;
    - release manifest;
    - package checksums;
    - asset checksums;
    - validation freeze;
    - testing runbook;
    - release notes asset;
    - LICENSE.
- **Release body drafting continuity**
  - First draft over-explained and drifted too far from the v0.7.0 public-template structure.
  - Operator corrected:
    - not every table row can be high priority;
    - public release text should be pragmatic and semi-public;
    - use the v0.7.0 headings/order/shape more strictly.
  - Final accepted body used:
    - same heading order as v0.7.0;
    - same public pre-alpha tone;
    - shorter/pragmatic “What changed” and “Validated” tables;
    - realistic criticality labels;
    - strong but less private/internal wording;
    - explicit nonclaims and source-archive warning retained.
- **Public release meaning**
  - v0.8.0 is a Gitea-source-of-truth pre-release.
  - The attached v0.8.0 package is the intended multi-repo validation package.
  - Default Gitea source archives remain non-authoritative for validation.
  - Public release publication completes the v0.7.x operational-spine implementation epoch.
- **What v0.8.0 truthfully closes**
  - Gate 0 release bridge.
  - Gate A operational workflow contract.
  - Gate B Relay lifecycle integration.
  - Gate C bounded state-substrate enforcement.
  - Gate D bounded runtime aggregate.
  - Gate E manual-private native deployment posture.
  - Gate F release/package/runtime maturity runway:
    - F1 release/package/runtime surface classification;
    - F2 operator docs/runbook closure;
    - F3 compatibility/refusal/rollback observability;
    - F4 code-health/source-hygiene classification;
    - F5 basic local trust candidate posture;
    - F6 package/runtime candidate validation;
    - F7 release-candidate closure matrix/manual handoff.
- **What remains explicitly nonclaimed**
  - Finished messenger.
  - Production readiness.
  - Production E2EE.
  - Hostile-server safety.
  - Malicious-relay safety.
  - Metadata privacy/authenticity proof.
  - Verified identity.
  - Full trust promotion.
  - Secure enrollment.
  - Cryptographic identity binding across Cypher, Comms, and OpenMLS.
  - Automatic trust promotion.
  - Vault/backup/restore.
  - Migration/repair/destructive cleanup/state relocation.
  - `full-runtime-dev`.
  - Service/systemd/helper install.
  - Public ingress.
  - Containers.
  - TUI/public UX.
  - Android.
  - CarbonStackOS implementation.
  - PQ/hybrid security.
  - Adversarial harness coverage.
  - External audit/certification.
  - General-public usability.
- **Important continuity lessons**
  - The v0.8.0 release should be remembered as a disciplined bounded checkpoint, not as a product-readiness release.
  - Template adherence matters for semi-public release text:
    - public release bodies should be concise and release-facing;
    - deep private continuity belongs in PRIME/LogDocs and future roadmap docs, not necessarily the public release page.
  - Package shape matters:
    - internal release metadata paths must remain canonical;
    - external asset filenames may be versioned, but validator expectations should not be broken.
- **Next required action**
  - Generate the next evergreen roadmap after this PRIME is accepted.
  - Deprecate the v0.8.0 EVERGREEN as the active planning anchor after it has served its release purpose.
  - Start a fresh v0.9.x-plus adversarial roadmap Q/A before new implementation.
- **Safe handoff**
  - Do not immediately start coding a new epoch from residual v0.7.x momentum.
  - Next work should be roadmap generation and adversarial planning.
  - Do not start v0.9.x implementation, verified identity, vault/backup, containers, public ingress, TUI, Android, CarbonStackOS, PQ/hybrid, or `full-runtime-dev` without an accepted new roadmap/contract.



## 6. Current Critical Path / Function Map

### 6.1 Completed v0.7.x -> v0.8.0 critical path

```text
v0.7.0 public release / Gate 0 bridge
-> Gate A workflow contract
-> Gate B Relay lifecycle integration
-> Gate C state substrate
-> Gate D runtime aggregate
-> Gate E manual-private native deployment
-> Gate F release/package/runtime maturity runway
   -> F1 release/package/runtime surface classification
   -> F2 operator docs/runbook closure
   -> F3 compatibility / stale-state / rollback observability / refusal posture
   -> F4 code health / source hygiene
   -> F5 basic local trust candidate posture
   -> F6 package/runtime candidate validation
   -> F7 release-candidate closure matrix and manual release handoff
-> v0.8.0 release asset prep
-> v0.8.0 public release posted
-> v0.8.0 PRIME
-> next evergreen roadmap generation
```

### 6.2 Current state by repo

`carbonstack` carries the validation, registry, docs, and release-handoff surface for the v0.8.0 pre-release:

- Gate B closure profile;
- Gate C closure profile;
- Gate D runtime aggregate profile;
- Gate E native deployment and closure profiles;
- Gate F F1-F7 profiles;
- release/package/runtime candidate validation;
- release-candidate closure/manual handoff profile;
- generated command reference at `entries=140`;
- release-validation ladder `verify-checksums`, `release-snapshot`, `full-validate-release`, `full`, and `local-cypher`.

`carbonstack-comms` carries:

- Relay lifecycle work;
- KeyPackage and Welcome lifecycle work;
- workflow/onboarding reports;
- state-substrate C1-C4 commands;
- basic local trust posture and local acceptance-event commands;
- Comms-side docs for the F5 trust posture.

`carbonstack-cypher` carries:

- server entrypoint and local relay/coordination behavior;
- migrations and local DB behavior;
- invite/device/envelope APIs;
- KeyPackage publication/lookup support;
- terminating `--help`, `--print-config`, and `--check-config` inspection behavior.

`carbonstack-os` remains:

- future constrained-appliance OS doctrine and target direction;
- not part of the runnable v0.8.0 package;
- not implemented as a product/appliance release.

### 6.3 Current package/runtime posture

Allowed to say:

- package/runtime candidate validation exists;
- v0.8.0 public release package was posted by operator;
- attached package is the intended multi-repo validation package;
- release package validation path is `verify-checksums` then `full-validate-release`;
- package-internal metadata uses `release/manifest.json` and `release/checksums.txt`;
- default Gitea source archives are not the intended validation package.

Not allowed to say:

- public release tooling autonomously created/uploaded/published the release;
- package publication/staging is now a general release system;
- `full-runtime-dev` is promoted;
- service/systemd/helper install exists;
- public ingress/container/TUI exists.

### 6.4 Current identity/trust posture

Allowed to say:

- basic local trust posture exists;
- manual local trust candidate acceptance event path exists;
- identity-domain separation is explicit;
- local acceptance requires explicit operator action;
- Relay membership does not promote trust;
- Welcome/MLS join does not promote trust.

Not allowed to say:

- verified identity exists;
- full trust promotion exists;
- secure enrollment exists;
- server-hostile identity replacement proof exists;
- real-world person verification exists;
- Cypher, Comms, and OpenMLS identity domains are cryptographically unified;
- production E2EE is proven.

### 6.5 Must-not-promote surfaces after v0.8.0

These remain explicitly unpromoted:

- verified identity;
- full trust promotion;
- secure enrollment;
- cryptographic identity binding across Cypher/Comms/OpenMLS;
- automatic trust promotion;
- `full-runtime-dev`;
- migration implementation;
- silent migration;
- repair implementation;
- silent repair;
- destructive cleanup of state;
- state relocation;
- service/systemd/helper install;
- public ingress;
- containers;
- TUI;
- vault/backup/restore;
- PQ/hybrid;
- Android;
- CarbonStackOS implementation;
- adversarial harness coverage;
- production readiness;
- public/general usability.

---

## 7. Future Work / To-Do

### Immediate next step

Generate the next evergreen roadmap.

Do not start implementation from v0.7.x momentum before the next roadmap is accepted.

### Next roadmap should decide

- v0.9.x adversarial/native-platform readiness order.
- Whether the next major work is:
  - adversarial test harness design/implementation;
  - native deployment hardening;
  - migration/repair/backup/vault split;
  - verified identity/trust promotion roadmap;
  - `full-runtime-dev` promotion decision;
  - TUI over stable noninteractive surfaces;
  - container/public ingress promotion after native hardening;
  - Android/CarbonStackOS lane sequencing.
- How to deprecate or archive the v0.8.0 EVERGREEN roadmap.
- Which v0.8.0 docs/profiles remain active authority and which become historical release-era references.
- How to preserve release-package validation discipline without letting it become deployment claims.

### Do not start without explicit next-roadmap contract

- v0.9.x implementation;
- adversarial harness implementation;
- verified identity/trust promotion;
- secure enrollment;
- cryptographic identity binding;
- vault/backup/restore;
- migration/repair/destructive cleanup/state relocation;
- `full-runtime-dev`;
- TUI/public UX;
- service/systemd/helper install;
- public ingress;
- containers;
- Android;
- CarbonStackOS;
- PQ/hybrid implementation.

---

## 8. Safe Resume State

Current safe resume:

```text
carbonstack
  715fd76f7700e63f1de877b730c2ce6bb37225bf
  test: add Gate F release candidate closure profile

carbonstack-comms
  32681e20784e560fdc8075bbc8a6597b742823c9
  feat: add Gate F basic local trust posture

carbonstack-cypher
  59c732ef53e198f65c03cc4e4178a66521c26a4c
  feat: add Cypher config inspection flags

carbonstack-os
  1bbbe52020d623b81796694e5057c1d080ede3ea
  docs: clarify CarbonStackOS target direction
```

Current gate status:

```text
Gate 0: CLOSED
Gate A: CLOSED
Gate B: CLOSED
Gate C: CLOSED
Gate D: CLOSED
Gate E: CLOSED
Gate F: CLOSED
v0.8.0 release: POSTED
Next evergreen roadmap: NOT GENERATED
```

Current registry/reference status:

```text
command_reference_entries=140
generated_reference_current=true
missing_nonclaims=0
```

Next safest action:

```text
Generate the next evergreen roadmap after accepting this PRIME.
```

Canonical resume sentence:

```text
CarbonStack v0.8.0 PRIME private continuity: the v0.8.0 Operational Spine Maturation Pre-Release has been operator-posted on Gitea after v0.7.26 closed Gate 0, Gate A, Gate B, Gate C, Gate D, Gate E, and Gate F; current source heads are carbonstack 715fd76f7700e63f1de877b730c2ce6bb37225bf, carbonstack-comms 32681e20784e560fdc8075bbc8a6597b742823c9, carbonstack-cypher 59c732ef53e198f65c03cc4e4178a66521c26a4c, carbonstack-os 1bbbe52020d623b81796694e5057c1d080ede3ea; command reference entries=140; the public release package is a Debian/WSL Debian multi-repo validation package with verify-checksums and full-validate-release as the preferred fresh-extraction path; v0.8.0 closes the v0.7.x operational-spine integration epoch but does not claim production readiness, production E2EE, hostile-server safety, verified identity, full trust promotion, secure enrollment, cryptographic identity binding, full-runtime-dev, migration/repair, vault/backup, public ingress, containers, TUI, PQ/hybrid, Android, or CarbonStackOS implementation; next action is to generate the next evergreen roadmap and not start v0.9.x implementation from memory.
```

---

## 9. Artifact Source Integrity

Primary source artifacts used for this PRIME:

- `CarbonStackLogDocV0.7.26.md`
- operator report that v0.8.0 release was posted:
  - `https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases`
- v0.8.0 release asset-prep continuity:
  - initial package-shape failure around missing canonical `release/checksums.txt`;
  - corrected V2 package asset prep using canonical internal metadata;
  - manual upload of final v0.8.0 assets.
- v0.8.0 release-body drafting continuity:
  - first draft rejected as too private/internal and over-prioritized;
  - corrected draft accepted with stricter v0.7.0 template adherence.

Generated continuity artifact:

- `CarbonStackLogDocV0.8.0PRIME.md`

Generation posture:

- PRIME generated outside WSL.
- Markdown LogDoc preserves the full point-form ledger through v0.8.0 PRIME.
- No JSON Breakpoint was requested in this step.
- Next artifact should be the next evergreen roadmap after PRIME acceptance.
