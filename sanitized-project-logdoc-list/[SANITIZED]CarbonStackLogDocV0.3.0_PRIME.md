[META NOTE: REDACTIONS REPRESENTED AS ▓▓ CHARACTERS]

# CarbonStack LogDoc v0.3.0[PRIME]

**Last updated:** 2026-05-29 -04:00  
**Owner/operator:** bitcrusher32  
**Assistant/model:** GPT-5.5 Thinking  
**Current phase:** Phase 2D / Phase 2E-prep **v0.3.0[PRIME] experimental backbone release checkpoint**. `carbonstack` is tagged `v0.3.0` at `59245fc docs: freeze v0.3.0 release packaging plan`, after `ced826a`, `40670f2`, `6755460`, `af11446`, and `1affcf9`; `carbonstack-comms` remains at `6052f52 docs: harden OpenMLS self-test wording`; `carbonstack-cypher` remains at `13312c6 docs: standardize Cypher envelope semantics`; `carbonstack-os` remains at `b537475 Add CarbonStackOS north star and initial CarbonStack repository structure`. v0.3.0[PRIME] records the actual release/tag/asset publication step after the v0.2.70 packaging freeze. The public release lives at `https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases` under the title `CarbonStack v0.3.0 experimental backbone epoch`. The intended release package includes clean source snapshots for `carbonstack`, `carbonstack-comms`, and `carbonstack-cypher`, plus a release manifest, validation freeze note, checksums, `LICENSE`, and `v0.3.0-minor-epoch-release.md`; `carbonstack-os` remains future context and is not included as a runnable proof component. The Gitea default `Source Code (ZIP)` / `Source Code (TAR.GZ)` archives are auto-generated archives of the tagged `carbonstack` repo only; they are not the full intended multi-repo v0.3.0 package. The release remains pre-alpha / experimental and proves a repeatable local Cypher + Comms OpenMLS relay backbone lifecycle, not production readiness, production E2EE, hostile-server safety, metadata privacy, Android readiness, secure vault/storage, external audit, or certification. The `[PRIME]` flag denotes this uncompressed LogDoc handoff; do not treat it as a compressed summary.
**Version schema:** v[scope].[timeline] — this file is `v0.3.0[PRIME]`, the uncompressed v0.3.0 experimental backbone release checkpoint after v0.2.70 release packaging freeze. The `[PRIME]` suffix is a continuity flag meaning this LogDoc intentionally preserves full baseline context for compression in a later work session.


---

## 1. Project Goal

**Active goal:** Preserve CarbonStack state after the actual v0.3.0 experimental backbone release. Phase 2D OpenMLS research remains mainline-closed and validated; v0.2.43-v0.2.46 completed the sidecar maintainability promotion/split/cleanup ladder; v0.2.47-v0.2.48 recovered and widened the Cypher envelope relay path for OpenMLS artifacts; v0.2.49-v0.2.51 established `carbonstack-comms/internal/relay` and the Comms/Cypher bridge helpers; v0.2.52-v0.2.54 proved application-message, KeyPackage+Welcome, and then the stitched full lifecycle through a Cypher-compatible test server; v0.2.55 planned deployability; v0.2.56 proved the full lifecycle against a real locally-started `carbonstack-cypher` process; v0.2.57 wrapped that proof in a repeatable dev smoke harness; v0.2.58 proved safe consume-then-ack boundaries; v0.2.59 planned decoded-payload metadata; v0.2.60 implemented that metadata across Cypher and Comms; v0.2.61 added the main-repo experimental backbone runbook and docs archive framing; v0.2.62 cleaned the public-facing CarbonStack/Comms/Cypher surfaces; v0.2.63 recorded the known-good validation matrix; v0.2.64 standardized envelope lifecycle semantics, idempotent ack, and schema/API wording; v0.2.65 planned the OpenMLS backbone self-test harness public surface; v0.2.66 implemented that wrapper in `carbonstack-comms`; v0.2.67 recorded the pre-v0.3.0 release-hardening recon; v0.2.68 implemented the consolidated `docs/v0.3.0-minor-epoch-release.md` release README plus supporting public-surface cleanup; v0.2.69 closed the final stale public wording defects found by the stale-claims scout; v0.2.70 froze the v0.3.0 release package plan in `docs/119-v0.3.0-release-packaging-freeze-v0.md`; and v0.3.0 now tags/releases the experimental backbone epoch under the `carbonstack` repo release surface. Execution still remains in `carbonstack-comms`; `carbonstack` remains the public release/front-door documentation repo and tag/release surface; `carbonstack-cypher` remains the relay/server implementation surface. The next safest work is post-v0.3.0 release inspection and backbone maturation planning: verify assets remain downloadable, optionally record or fix any LICENSE-in-snapshot discrepancy in the next release process, then begin v0.3.x backbone maturation/portability work without retroactively inflating v0.3.0 claims.

CarbonStack’s broader project intent remains:

- **CarbonStack** is the secure communications appliance stack and shared doctrine.
- **CarbonStackOS** is the future constrained Android-derived appliance OS.
- **CarbonStackComms** is the text-first encrypted messaging client.
- **CarbonStackCypher** is the self-hostable hostile-server relay/storage stack.
- The `carbonstack` repo remains the doctrine/spec/docs authority; implementation repos must not contradict it without revising the spec.

**North star:**

> Its goal is not to be a better smartphone. Its goal is to stop being a smartphone.

**Core doctrine:**

> Every feature is guilty until it proves it does not add unacceptable parser, network, sensor, identity, or filesystem authority.

**Current Phase 2D sidecar capabilities after v0.2.38:**

- `provider-info` reports sanitized provider metadata using structured `serde_json::json!` and shared command arrays instead of fragile raw JSON.
- `identity-create --device-label <safe>` creates dev-only OpenMLS identity/signing state.
- `identity-status --device-label <safe>` loads existing identity state and verifies public identity ref.
- `public-bundle-export --device-label <safe>` generates a real OpenMLS KeyPackage and writes sanitized public-bundle summary metadata.
- `public-bundle-export --device-label <safe> --write-artifact` serializes the public KeyPackage artifact under ignored dev state as `public-bundle.keypackage.bin`, writes manifest/summary metadata, and saves device provider storage because the private KeyPackage bundle is required for later Welcome consumption.
- `conversation-create --device-label <safe> --conversation-label <safe>` creates a dev-local one-member OpenMLS group/conversation, writes sanitized `conversation-summary.json`, writes dev-local `provider-storage.json`, and proves group reloadability.
- `conversation-load-check --device-label <safe> --conversation-label <safe>` loads saved dev provider storage in a later sidecar command and proves `MlsGroup::load(provider.storage(), &group_id)` succeeds.
- `conversation-add-member --device-label <creator-device> --conversation-label <conversation> --member-keypackage <path>` loads the creator's persisted group, consumes a member public KeyPackage artifact, calls `add_members`, writes a Welcome carrier artifact, merges the pending commit, saves mutated provider storage, and preserves reloadability.
- `conversation-join --device-label <joining-device> --conversation-label <conversation> --welcome <path>` consumes the existing Welcome carrier artifact, stages and joins Bob into the group, saves device-scoped joined provider storage, and proves the joined group is reloadable.
- `message-protect --device-label <sender-device> --conversation-label <conversation> --plaintext <text>` loads Alice’s persisted group, protects bounded dev plaintext, writes `application-message.bin`, saves provider storage, and proves reloadability.
- `message-open --device-label <receiver-device> --conversation-label <conversation> --message <path>` loads Bob’s joined group, opens the protected MLS application message, returns bounded dev plaintext, saves provider storage, and proves reloadability.
- `state-checkpoint` and `state-load-check` remain unsupported.









**New at v0.3.0[PRIME]:**

- `carbonstack` is tagged `v0.3.0` at `59245fc docs: freeze v0.3.0 release packaging plan`.
- The public release page is:
  - `https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases`.
- Release title:
  - `CarbonStack v0.3.0 experimental backbone epoch`.
- Release status:
  - pre-alpha / experimental.
- Primary artifact:
  - Cypher + Comms OpenMLS relay backbone.
- `carbonstack-comms` remains at `6052f52 docs: harden OpenMLS self-test wording`.
- `carbonstack-cypher` remains at `13312c6 docs: standardize Cypher envelope semantics`.
- `carbonstack-os` remains at `b537475 Add CarbonStackOS north star and initial CarbonStack repository structure`.
- `carbonstack-os` remains future appliance OS/north-star context and is not part of the runnable v0.3.0 proof package.
- v0.3.0[PRIME] is the actual release/tag/asset publication checkpoint. It follows the v0.2.70 release packaging freeze. It is not new protocol/runtime work.
- The `[PRIME]` suffix is an operator continuity flag meaning this LogDoc is intentionally uncompressed. Do not compress continuity or baseline details inside this file; compression should happen at the next work session.
- The `v0.3.0` annotated tag was created and pushed from the `carbonstack` repo:
  - `git tag -a v0.3.0 -m "CarbonStack v0.3.0 experimental backbone epoch"`;
  - `git push origin v0.3.0`.
- Final tag snapshot showed:
  - `59245fc (HEAD -> main, tag: v0.3.0, origin/main, origin/HEAD) docs: freeze v0.3.0 release packaging plan`;
  - `v0.3.0` points at `carbonstack` HEAD.
- Gitea initially showed a tag/release page with only default generated downloads:
  - `Source Code (ZIP)`;
  - `Source Code (TAR.GZ)`.
- Important release blunder/continuity note:
  - the default Gitea source archives are only auto-generated archives of the `carbonstack` repo at the `v0.3.0` tag;
  - they are not the intended multi-repo v0.3.0 release package;
  - this was recognized after seeing the Gitea release UI.
- The intended release package assets were then attached manually to the release page.
- Final release downloads visible on the release page include:
  - `carbonstack-comms-v0.3.0-source-snapshot.zip`;
  - `carbonstack-cypher-v0.3.0-source-snapshot.zip`;
  - `carbonstack-v0.3.0-checksums.txt`;
  - `carbonstack-v0.3.0-release-manifest.json`;
  - `carbonstack-v0.3.0-source-snapshot.zip`;
  - `carbonstack-v0.3.0-validation-freeze.md`;
  - `LICENSE`;
  - `v0.3.0-minor-epoch-release.md`.
- The release body now states:
  - Status: pre-alpha / experimental;
  - Primary artifact: Cypher + Comms OpenMLS relay backbone;
  - the release packages the current known-good experimental backbone proof;
  - start at `docs/v0.3.0-minor-epoch-release.md`;
  - use attached v0.3.0 source snapshots and manifest for the intended release package rather than Gitea default `Source Code` archives;
  - boundary/nonclaims remain explicit.
- The release body lists the current known-good experimental backbone proof:
  - CarbonStackComms OpenMLS sidecar;
  - CarbonStackCypher real local relay server;
  - opaque OpenMLS artifact envelope relay;
  - payload metadata validation;
  - consume-then-ack semantics;
  - OpenMLS backbone self-test harness.
- The release body boundary says:
  - this release proves a repeatable local experimental backbone lifecycle;
  - it does not prove production readiness;
  - it does not prove production E2EE;
  - it does not prove hostile-server safety;
  - it does not prove metadata privacy;
  - it does not prove Android readiness;
  - it does not prove secure vault/storage;
  - it does not prove external audit or certification.
- The release body adds a welcome/boundary line:
  - this release is designed for developers, enthusiasts, and testers;
  - iteration, testing, and development are welcome, private or public.
- Release asset staging path used:
  - `C:\▓▓\repos\carbonstack_umbrella\release-staging\v0.3.0\assets`.
- Release assets generated locally before upload:
  - `carbonstack-v0.3.0-source-snapshot.zip`;
  - `carbonstack-comms-v0.3.0-source-snapshot.zip`;
  - `carbonstack-cypher-v0.3.0-source-snapshot.zip`;
  - `carbonstack-v0.3.0-release-manifest.json`;
  - `carbonstack-v0.3.0-validation-freeze.md`;
  - `carbonstack-v0.3.0-checksums.txt`.
- Local asset listing before release upload:
  - `carbonstack-comms-v0.3.0-source-snapshot.zip` length `170266`;
  - `carbonstack-cypher-v0.3.0-source-snapshot.zip` length `23142`;
  - `carbonstack-v0.3.0-checksums.txt` length `549`;
  - `carbonstack-v0.3.0-release-manifest.json` length `4553`;
  - `carbonstack-v0.3.0-source-snapshot.zip` length `383514`;
  - `carbonstack-v0.3.0-validation-freeze.md` length `1395`.
- Checksums generated and recorded:
  - `904E0911F20EC8147706F9BA5F87977C054D5B577F73C355142EABF374078855  carbonstack-v0.3.0-source-snapshot.zip`;
  - `CC09E29DD2124BAD0D44B1231B3D440A07099E9BE86161C361FEDB052678398B  carbonstack-comms-v0.3.0-source-snapshot.zip`;
  - `5FFBAB3CC86030450D6BCAC62D8A209EB166284F6433CB846232D86D46C304F6  carbonstack-cypher-v0.3.0-source-snapshot.zip`;
  - `99355E724C63C5A0E4946F07938D5A5929C543C684454C38F2AFB2AE0A9A1807  carbonstack-v0.3.0-release-manifest.json`;
  - `79054C995FCA15ED676408E0DAA8282F6B579EC4A53527A919785ACA89D10FFB  carbonstack-v0.3.0-validation-freeze.md`.
- The release page also includes a standalone `LICENSE` asset.
- Important license continuity note:
  - the `LICENSE` asset was added to the release surface after the snapshot/package discussion;
  - if the source snapshot ZIPs were not regenerated after adding LICENSE files to repos, those ZIPs may not contain the standalone license internally;
  - this is acceptable for this v0.3.0 release if the release page carries the license beside the assets;
  - future releases should commit/include `LICENSE` in each relevant repo before generating source snapshot ZIPs.
- Another important packaging continuity note:
  - the attached `v0.3.0-minor-epoch-release.md` file is a helpful direct release asset, while the canonical copy also remains inside the `carbonstack` source snapshot.
- The release package intentionally uses source snapshots rather than hand-curated partial runnable folders.
- Reason preserved:
  - the proof/self-test crosses Go packages, Rust sidecar source, scripts, test helpers, migrations, internal relay/client code, and docs;
  - hand-curated runnable subsets risk omitting required files.
- The release package intentionally includes:
  - `carbonstack` source snapshot for public release surface/docs/runbooks/validation matrix/lifecycle semantics;
  - `carbonstack-comms` source snapshot for OpenMLS sidecar bridge, relay helpers, real-Cypher lifecycle tests, and OpenMLS backbone self-test wrapper;
  - `carbonstack-cypher` source snapshot for real local relay/server, envelope API, OpenMLS content types, payload metadata, queued-only inbox, idempotent ack, and SQLite migrations.
- The release package intentionally does not include:
  - `.git/`;
  - `target/`;
  - `.carbonstack-openmls-sidecar-state/`;
  - `provider-storage.json`;
  - `signer.json`;
  - SQLite runtime DBs;
  - `.go-cache/`;
  - `.go-tmp/`;
  - Go build caches;
  - temp Go test binaries;
  - generated OpenMLS private/dev state;
  - local machine temp directories;
  - editor caches;
  - OS metadata files.
- No generated OpenMLS private/dev state is intentionally included in release assets.
- No `provider-storage.json` or `signer.json` should be included.
- No quarantined Go test binaries should be restored or shipped.
- Final repo/tag snapshot after tag push:
  - `carbonstack`: `59245fc (HEAD -> main, tag: v0.3.0, origin/main, origin/HEAD) docs: freeze v0.3.0 release packaging plan`;
  - `carbonstack-comms`: `6052f52 (HEAD -> main, origin/main, origin/HEAD) docs: harden OpenMLS self-test wording`;
  - `carbonstack-cypher`: `13312c6 (HEAD -> main, origin/main, origin/HEAD) docs: standardize Cypher envelope semantics`;
  - `carbonstack-os`: `b537475 (HEAD -> main, origin/main, origin/HEAD) Add CarbonStackOS north star and initial CarbonStack repository structure`.
- The final release asset list visible in the screenshot differs slightly from the initial local list because it includes manually added:
  - `LICENSE`;
  - `v0.3.0-minor-epoch-release.md`.
- No additional code or protocol behavior was added during the v0.3.0 release step.
- No new Cypher route was added.
- No `/v0/artifacts` API was added.
- No top-level `carbonstack` orchestrator was added.
- No OpenMLS runtime `cmd/comms send` / `inbox` UX was added.
- No Android work was added.
- No vault work was added.
- No hostile-server-complete proof was added.
- No production E2EE claim was added.
- No metadata-privacy claim was added.
- No external-audit or certified-secure claim was added.
- This is the first public CarbonStack release/tag under the `carbonstack` repo release surface.
- v0.3.0 closes the pre-v0.3.0 ladder from v0.2.55 through v0.2.70:
  - deployable server/CLI smoke-test planning;
  - real Cypher server relay proof;
  - repeatable local smoke harness;
  - consume-then-ack semantics;
  - payload metadata;
  - deployability/runbook/docs archive;
  - public surface cleanup;
  - lifecycle/schema standardization;
  - OpenMLS backbone self-test wrapper;
  - v0.3.0 release README;
  - stale-claims cleanup;
  - release packaging freeze;
  - actual release/tag/assets.
- Immediate next safest work after v0.3.0:
  - inspect the published release page/assets once more from a clean browser/session;
  - optionally download and checksum the attached assets to verify the public uploads match local `carbonstack-v0.3.0-checksums.txt`;
  - update any repo LICENSE placement in a future cleanup if needed;
  - begin post-v0.3.0 preflight rather than adding features directly.
- Recommended next epoch:
  - v0.3.x backbone maturation and portability.
- Likely post-v0.3.0 direction:
  - reduce Windows-only assumptions;
  - improve self-test/runbook UX;
  - validate release snapshots from a clean extraction;
  - clarify Cypher deployment configuration;
  - start planning runtime Comms OpenMLS `send` / `inbox` integration only after the backbone release is verified;
  - plan secure local vault/storage later;
  - plan hostile-server rollback/replay/metadata-abuse harnesses later;
  - keep Android/Pixel/CarbonStackOS work deferred until the core backbone is mature.
- The current release remains experimental and pre-alpha; do not retroactively describe it as production-ready.


**New at v0.2.70:**

- `carbonstack` head is now `59245fc docs: freeze v0.3.0 release packaging plan`.
- `carbonstack-comms` remains at `6052f52 docs: harden OpenMLS self-test wording`.
- `carbonstack-cypher` remains at `13312c6 docs: standardize Cypher envelope semantics`.
- `carbonstack-os` remains at `b537475 Add CarbonStackOS north star and initial CarbonStack repository structure`.
- v0.2.70 is the final v0.3.0 release packaging freeze checkpoint. It is not the v0.3.0 tag/release checkpoint and not a new protocol/runtime behavior checkpoint.
- `carbonstack/docs/119-v0.3.0-release-packaging-freeze-v0.md` now records the release packaging freeze / pre-tag checklist.
- `docs/119` freezes the public release name:
  - `CarbonStack v0.3.0 experimental backbone epoch`.
- `docs/119` freezes the release class:
  - pre-alpha / experimental.
- `docs/119` freezes the primary artifact:
  - Cypher + Comms OpenMLS relay backbone.
- `docs/119` points the public release README to:
  - `docs/v0.3.0-minor-epoch-release.md`.
- `docs/v0.3.0-minor-epoch-release.md` now links to `docs/119-v0.3.0-release-packaging-freeze-v0.md`.
- `carbonstack/docs/README.md` now links to `docs/119` as the current v0.3.0 packaging freeze.
- The release package is intentionally source-snapshot based rather than hand-curated into partial runnable folders.
- The reason for source snapshots is explicit:
  - the self-test crosses Go packages;
  - Rust sidecar source;
  - scripts;
  - test helpers;
  - migrations;
  - internal relay/client code;
  - docs;
  - hand-curated subsets risk omitting required files.
- The v0.3.0 release should include clean source snapshots for:
  - `carbonstack`;
  - `carbonstack-comms`;
  - `carbonstack-cypher`.
- The v0.3.0 release should not include `carbonstack-os` as a runnable proof component.
- `carbonstack-os` may be referenced only as future north-star context.
- Recommended release asset names are now frozen:
  - `carbonstack-v0.3.0-source-snapshot.zip`;
  - `carbonstack-comms-v0.3.0-source-snapshot.zip`;
  - `carbonstack-cypher-v0.3.0-source-snapshot.zip`;
  - `carbonstack-v0.3.0-release-manifest.json`;
  - `carbonstack-v0.3.0-validation-freeze.md`;
  - `carbonstack-v0.3.0-checksums.txt`.
- Critical material inside the `carbonstack` snapshot is recorded:
  - `README.md`;
  - `docs/README.md`;
  - `docs/v0.3.0-minor-epoch-release.md`;
  - `docs/113-experimental-backbone-deployability-runbook-v0.md`;
  - `docs/114-known-good-validation-matrix-v0.md`;
  - `docs/115-envelope-lifecycle-semantics-v0.md`;
  - `docs/117-openmls-backbone-self-test-harness-result-v0.md`;
  - `docs/118-v0.3.0-release-hardening-recon-v0.md`;
  - `docs/119-v0.3.0-release-packaging-freeze-v0.md`;
  - `roadmap/ROADMAP.md`;
  - `scripts/validate-local.ps1`.
- Critical material inside the `carbonstack-comms` snapshot is recorded:
  - `README.md`;
  - `scripts/self-test-openmls-backbone.ps1`;
  - `scripts/smoke-openmls-real-cypher-relay.ps1`;
  - `scripts/check-no-rust-artifacts.ps1`;
  - `scripts/README.md`;
  - `cmd/comms/`;
  - `internal/client/`;
  - `internal/relay/`;
  - `internal/protocol/`;
  - `internal/protocol/mls/openmls-sidecar/`;
  - `go.mod`;
  - `go.sum`.
- `docs/119` explicitly says test files under `internal/protocol` and `internal/relay` are part of the proof surface.
- `docs/119` preserves the runtime boundary:
  - current `cmd/comms send` / `inbox` remains stub-era;
  - it must not be presented as the OpenMLS messenger UX.
- Critical material inside the `carbonstack-cypher` snapshot is recorded:
  - `README.md`;
  - `cmd/cypher/`;
  - `internal/config/`;
  - `internal/db/`;
  - `internal/httpapi/`;
  - `migrations/`;
  - `docs/02-envelope-model.md`;
  - `docs/03-api-surface.md`;
  - `docs/04-storage-model.md`;
  - `docs/07-data-model-v0.md`;
  - `docs/08-api-contract-v0.md`;
  - `go.mod`;
  - `go.sum`.
- `docs/119` records Cypher’s release role:
  - local relay/server process;
  - HTTP envelope API;
  - OpenMLS artifact content-type acceptance;
  - payload metadata;
  - queued-only inbox semantics;
  - idempotent same-recipient ack;
  - SQLite migrations.
- `docs/119` freezes the material that must not be included in release assets:
  - `.git/`;
  - `target/`;
  - `.carbonstack-openmls-sidecar-state/`;
  - `provider-storage.json`;
  - `signer.json`;
  - `*.db`;
  - `*.db-shm`;
  - `*.db-wal`;
  - `.go-cache/`;
  - `.go-tmp/`;
  - Go build caches;
  - temp Go test binaries;
  - generated OpenMLS private/dev state;
  - local machine temp directories;
  - editor caches;
  - OS metadata files.
- `docs/119` explicitly says not to restore or ship quarantined Go test binaries; let Go rebuild them during validation.
- Validation freeze commands are recorded for:
  - `go test ./... -count=1` in `carbonstack-cypher`;
  - `scripts/self-test-openmls-backbone.ps1` in `carbonstack-comms`;
  - `scripts/self-test-openmls-backbone.ps1 -Full` in `carbonstack-comms`;
  - `scripts/check-no-rust-artifacts.ps1` in `carbonstack-comms`;
  - `scripts/validate-local.ps1` in `carbonstack`.
- The Windows antivirus / Go temp-cache note is preserved:
  - preferred root remains `C:\▓▓\AppData\Local\Temp\go-build-carbonstack\`;
  - expected subpaths are `tmp` and `cache`;
  - do not exclude all `%TEMP%`;
  - do not commit Go cache/temp directories.
- `docs/119` recommends the release staging root:
  - `C:\▓▓\repos\carbonstack_umbrella\release-staging\v0.3.0`.
- `docs/119` records expected release manifest fields:
  - `release_name`;
  - `release_class`;
  - `release_readme`;
  - `included_components`;
  - `excluded_material`;
  - `validation_commands`;
  - `validation_result`;
  - `nonclaims`;
  - `generated_at_local_time`.
- Required nonclaims remain frozen:
  - no secure messenger claim;
  - no production E2EE claim;
  - no production readiness claim;
  - no hostile-server safety claim;
  - no metadata privacy claim;
  - no trustless-operation claim;
  - no drop-in replacement claim;
  - no finished-product claim;
  - no Android-readiness claim;
  - no external-audit claim;
  - no certification or audit-ready claim.
- Correct v0.3.0 framing remains:
  - experimental backbone;
  - known-good local proof;
  - OpenMLS backbone self-test;
  - Cypher + Comms OpenMLS relay backbone;
  - not production-certified;
  - not externally audited;
  - not a finished messenger;
  - not Android-ready.
- `docs/119` freezes the v0.3.0 release step:
  - run final validation;
  - confirm clean repo states;
  - create source snapshots;
  - create release manifest;
  - create checksums;
  - attach or publish release assets under the `carbonstack` repo release surface;
  - tag or mark the CarbonStack v0.3.0 experimental backbone epoch;
  - update LogDoc to v0.3.0.
- Important continuity/blunder notes:
  - This rung intentionally answered the release packaging question by choosing full clean source snapshots with a manifest over fragile hand-curated runnable folders.
  - Public release readers should not need to care about arbitrary branch heads as the source of truth; release assets should carry the known-good snapshots.
  - Exact heads remain important for LogDoc continuity and internal handoff.
  - The release package should not include `carbonstack-os` as a runnable proof component.
  - The Avast/Go-generated-test-exe problem remains local environment noise; use the narrow `go-build-carbonstack` temp/cache root and exception if it appears again.
- Validation passed for the docs-only packaging-freeze change through `carbonstack/scripts/validate-local.ps1`.
- No new protocol behavior was added.
- No new Cypher route was added.
- No `/v0/artifacts` API was added.
- No top-level `carbonstack` orchestrator was added.
- No OpenMLS runtime `cmd/comms send` / `inbox` UX was added.
- No Android, vault, hostile-server-complete, production E2EE, external-audit, or certified-secure claim was added.
- Next safest step:
  - v0.3.0 release/tag/asset packaging step.
- Expected v0.3.0 scope:
  - run final validation again from clean repo states;
  - generate clean source snapshot ZIPs for `carbonstack`, `carbonstack-comms`, and `carbonstack-cypher`;
  - generate `carbonstack-v0.3.0-release-manifest.json`;
  - generate `carbonstack-v0.3.0-validation-freeze.md`;
  - generate `carbonstack-v0.3.0-checksums.txt`;
  - attach/publish those assets under the `carbonstack` release surface;
  - tag or mark the v0.3.0 experimental backbone epoch;
  - update LogDoc to v0.3.0.
- Later:
  - post-v0.3.0 backbone maturation and portability before Android, production vault, or hostile-server-complete work.


**New at v0.2.69:**

- `carbonstack` head is now `ced826a docs: harden final v0.3.0 release wording`.
- `carbonstack-comms` head is now `6052f52 docs: harden OpenMLS self-test wording`.
- `carbonstack-cypher` remains at `13312c6 docs: standardize Cypher envelope semantics`.
- `carbonstack-os` remains at `b537475 Add CarbonStackOS north star and initial CarbonStack repository structure`.
- v0.2.69 is the final stale-claims / release-hardening cleanup checkpoint before the final validation freeze. It is not the v0.3.0 tag/release checkpoint and not a new protocol/runtime behavior checkpoint.
- The stale-claims scout found that most dangerous-term matches were safe explicit nonclaims:
  - `secure messenger`;
  - `production E2EE`;
  - `hostile-server safety`;
  - `metadata privacy`;
  - `metadata-private`;
  - `certified`;
  - `audit-ready`;
  - `Android-ready`;
  - `finished messenger`.
- Those matches remain valid when they are used to say what CarbonStack v0.3.0 does **not** prove.
- Actual release-surface defects were small and wording-focused:
  - stale public-facing `Option B CLI/dev-harness planning and implementation` in `docs/113`;
  - stale `Option C` and `Option B` phrasing in `docs/114`;
  - stale recon wording in `docs/118` saying the lower-level smoke harness still printed `Option C` language, even though v0.2.68 already fixed it;
  - `carbonstack-comms/README.md` saying “Run the current real-server OpenMLS relay smoke harness” while showing the self-test wrapper command;
  - `carbonstack-comms/internal/protocol/mls/openmls-sidecar/README.md` still pointing validation readers to the lower-level smoke harness instead of the public self-test wrapper.
- `carbonstack/docs/113-experimental-backbone-deployability-runbook-v0.md` now records:
  - OpenMLS backbone self-test harness implementation is complete at v0.2.66;
  - v0.3.0 release README implementation is complete at v0.2.68;
  - the remaining pre-v0.3.0 work is final stale-claims and validation sweep.
- `carbonstack/docs/114-known-good-validation-matrix-v0.md` now frames its phase as known-good OpenMLS backbone validation cleanup rather than `Option C` cleanup.
- `docs/114` now refers to OpenMLS backbone self-test harness work rather than `Option B CLI/dev-harness planning`.
- `carbonstack/docs/118-v0.3.0-release-hardening-recon-v0.md` now treats the lower-level smoke-script `Option C` wording as a previously found recon issue that was cleaned during v0.2.68.
- `carbonstack-comms/README.md` now says to run the current OpenMLS backbone self-test, matching the command:
  - `powershell -ExecutionPolicy Bypass -File .\scripts\self-test-openmls-backbone.ps1`.
- `carbonstack-comms/internal/protocol/mls/openmls-sidecar/README.md` now points validation readers to:
  - `powershell -ExecutionPolicy Bypass -File .\scripts\self-test-openmls-backbone.ps1`.
- The lower-level smoke harness remains available for implementation/debug use, but it is no longer the public-facing command surface.
- Validation passed after cleanup:
  - CarbonStack local validation passed;
  - Cypher validation passed through `carbonstack/scripts/validate-local.ps1`;
  - Comms package validation passed through `validate-local`;
  - trust lifecycle validation passed through `validate-local`;
  - no new code behavior was required.
- Important continuity/blunder notes:
  - The stale-claims scout was intentionally noisy; most matches were safe nonclaims and should not be removed.
  - `docs/116` and `docs/118` can still mention Option B/Option C as historical/recon language when explicitly saying those labels are not public release language.
  - The correct public language remains OpenMLS backbone self-test harness, known-good local backbone proof, and Cypher + Comms OpenMLS relay backbone.
  - The current v0.3.0 release README remains `carbonstack/docs/v0.3.0-minor-epoch-release.md`.
  - The Avast/Go-generated-test-exe problem remains local environment noise. The preferred mitigation remains the narrow `C:\▓▓\AppData\Local\Temp\go-build-carbonstack\` temp/cache root and exception, not a broad `%TEMP%` exclusion.
- No new protocol behavior was added.
- No new Cypher route was added.
- No `/v0/artifacts` API was added.
- No top-level `carbonstack` orchestrator was added.
- No OpenMLS runtime `cmd/comms send` / `inbox` UX was added.
- No Android, vault, hostile-server-complete, production E2EE, external-audit, or certified-secure claim was added.
- Next safest rung:
  - v0.2.70 final validation freeze and release packaging checklist.
- Expected v0.2.70 scope:
  - run known-good validation commands from clean repo states;
  - confirm no untracked local Go cache/temp artifacts or generated OpenMLS sidecar artifacts are present;
  - confirm `carbonstack/docs/v0.3.0-minor-epoch-release.md` is the release README entrypoint;
  - record release packaging assumptions for included repo snapshots/files;
  - prepare v0.3.0 experimental server-deployable CarbonStack backbone epoch.
- Later:
  - v0.3.0 experimental server-deployable CarbonStack backbone epoch.
  - post-v0.3.0 backbone maturation and portability before Android, production vault, or hostile-server-complete work.


**New at v0.2.68:**

- `carbonstack` head is now `40670f2 docs: add v0.3.0 release README`.
- `carbonstack-comms` head is now `c8ef6cf docs: point Comms README to backbone self-test`.
- `carbonstack-cypher` remains at `13312c6 docs: standardize Cypher envelope semantics`.
- `carbonstack-os` remains at `b537475 Add CarbonStackOS north star and initial CarbonStack repository structure`.
- v0.2.68 is the v0.3.0 release README implementation checkpoint. It is not the v0.3.0 tag/release checkpoint, not a new protocol/runtime behavior checkpoint, and not a production/security certification checkpoint.
- `carbonstack/docs/v0.3.0-minor-epoch-release.md` now exists as the consolidated v0.3.0 minor-epoch release README.
- The release README frames v0.3.0 as:
  - experimental communications backbone release;
  - pre-alpha / experimental;
  - primary artifact: Cypher + Comms OpenMLS relay backbone;
  - not a finished messenger;
  - not production-certified;
  - not externally audited;
  - not Android-ready;
  - not a complete secure communications product.
- The release README consolidates the release surface so readers do not need to assemble current truth from `docs/113` through `docs/118`.
- The release README explains the current known-good local backbone proof:
  - CarbonStackComms OpenMLS sidecar;
  - CarbonStackCypher real local relay server;
  - opaque OpenMLS artifact envelope relay;
  - payload metadata validation;
  - consume-then-ack semantics;
  - OpenMLS backbone self-test harness.
- The release README records the current proof path:
  - Bob exports an OpenMLS KeyPackage artifact;
  - Cypher relays the KeyPackage as an opaque envelope;
  - Alice retrieves the envelope payload;
  - Comms validates payload metadata before writing artifact bytes;
  - Alice consumes the KeyPackage through the OpenMLS sidecar;
  - Alice creates a Welcome artifact;
  - Cypher relays the Welcome as an opaque envelope;
  - Bob retrieves the envelope payload;
  - Comms validates payload metadata before writing artifact bytes;
  - Bob consumes the Welcome through the OpenMLS sidecar;
  - Alice protects an application-message artifact;
  - Cypher relays the application-message as an opaque envelope;
  - Bob retrieves the envelope payload;
  - Comms validates payload metadata before writing artifact bytes;
  - Bob consumes the application-message through the OpenMLS sidecar;
  - plaintext matches;
  - Comms acknowledges envelopes only after the relevant sidecar consume command succeeds.
- The release README uses `carbonstack-comms/scripts/self-test-openmls-backbone.ps1` as the public-facing self-test entrypoint.
- The release README keeps `scripts/smoke-openmls-real-cypher-relay.ps1` as the lower-level implementation/debug harness.
- The release README states the known-good validation environment is Windows + PowerShell + local Go/Rust toolchains.
- The release README explicitly does not claim cross-platform validation, Linux deployment validation, or production deployment readiness.
- The release README records the narrow local antivirus/Go workaround path:
  - `C:\▓▓\AppData\Local\Temp\go-build-carbonstack\`.
- The release README treats that path as a local validation workaround, not a release artifact.
- The release README explains the component repo roles:
  - `carbonstack` = public release surface, doctrine, runbooks, validation matrix, envelope lifecycle semantics, release README, docs archive;
  - `carbonstack-comms` = OpenMLS sidecar, Comms relay helpers, real-Cypher lifecycle tests, OpenMLS backbone self-test wrapper;
  - `carbonstack-cypher` = local relay/server component, envelope API, OpenMLS artifact content types, payload metadata, queued-only inbox semantics, idempotent same-recipient ack;
  - `carbonstack-os` = future appliance OS north-star repo, not part of the current runnable OpenMLS backbone proof.
- The release README intentionally does not make readers care about exact branch heads as the public source of truth; the v0.3.0 release should include the known-good repo snapshots/files needed for that specific release.
- The release README defines core terms at the release surface:
  - envelope;
  - opaque payload;
  - payload metadata;
  - sidecar consume;
  - ack.
- The release README explicitly says payload metadata is storage/transport sanity metadata, not a trust root and not proof of OpenMLS authenticity.
- The release README preserves the ack boundary:
  - in the current Comms proof, ack is sent only after sidecar consume succeeds;
  - Cypher records ack;
  - Cypher does not know sidecar consume state.
- The release README explicitly rejects production/security overclaims:
  - no production E2EE;
  - no production readiness;
  - no hostile-server safety;
  - no metadata privacy;
  - no secure local vault/storage;
  - no stable public protocol status;
  - no Android readiness;
  - no complete Comms runtime UX;
  - no multi-user production operations;
  - no external audit;
  - no certification.
- The release README says not to call v0.3.0:
  - a secure messenger;
  - production-ready;
  - hostile-server safe;
  - metadata-private;
  - certified;
  - audit-ready.
- `carbonstack/README.md` now points to `docs/v0.3.0-minor-epoch-release.md` as the v0.3.0 experimental backbone release README and consolidated release entrypoint.
- `carbonstack/roadmap/ROADMAP.md` was updated away from public-facing internal `Option B` / `Option C` terminology toward:
  - OpenMLS backbone self-test harness;
  - known-good local backbone proof;
  - v0.3.0 release README;
  - v0.3.0 experimental backbone epoch.
- `carbonstack-comms/README.md` now points the public self-test command to:
  - `powershell -ExecutionPolicy Bypass -File .\scripts\self-test-openmls-backbone.ps1`.
- `carbonstack-comms/README.md` now keeps the lower-level implementation harness as:
  - `powershell -ExecutionPolicy Bypass -File .\scripts\smoke-openmls-real-cypher-relay.ps1`.
- `carbonstack-comms/scripts/smoke-openmls-real-cypher-relay.ps1` no longer prints visible `Option C` wording; it now frames itself as the current known-good OpenMLS backbone proof.
- No `carbonstack-cypher` changes were required at v0.2.68.
- Validation expectation for this rung remains:
  - `powershell -ExecutionPolicy Bypass -File .\scripts\self-test-openmls-backbone.ps1` in `carbonstack-comms`;
  - `powershell -ExecutionPolicy Bypass -File .\scripts\self-test-openmls-backbone.ps1 -Full` in `carbonstack-comms`;
  - `powershell -ExecutionPolicy Bypass -File .\scripts\check-no-rust-artifacts.ps1` in `carbonstack-comms`;
  - `go test ./... -count=1` in `carbonstack-cypher`;
  - `powershell -ExecutionPolicy Bypass -File .\scripts\validate-local.ps1` in `carbonstack`, using the dedicated Go temp/cache workaround if Avast interferes.
- Important continuity/blunder notes:
  - The release README intentionally avoids making public readers care about exact component heads because the release artifact should include the specific known-good repo snapshots/files.
  - Exact heads still matter internally for LogDoc/breakpoint continuity.
  - The earlier odd `af11446 docs: update 133...` and duplicate Comms wrapper commits remain continuity notes only; current heads are authoritative.
  - The Avast/Go-generated-test-exe issue remains local environment noise; the preferred mitigation is still the narrow `go-build-carbonstack` temp/cache root and Avast exception, not broad `%TEMP%` exclusion.
- No new protocol behavior was added.
- No new Cypher route was added.
- No `/v0/artifacts` API was added.
- No top-level `carbonstack` orchestrator was added.
- No OpenMLS runtime `cmd/comms send` / `inbox` UX was added.
- No Android, vault, hostile-server-complete, production E2EE, external-audit, or certified-secure claim was added.
- Next safest rung:
  - v0.2.69 final stale-claims / validation / release-readiness sweep.
- Expected v0.2.69 scope:
  - scan public current surfaces for stale `Option B` / `Option C` release wording;
  - scan for overclaims like secure messenger, production E2EE, hostile-server safe, metadata-private, trustless, certified, audit-ready;
  - verify `carbonstack` release README and component READMEs point to the right self-test/runbook surfaces;
  - run known-good validation;
  - record final v0.3.0 epoch readiness without adding protocol/runtime features.
- Later:
  - v0.3.0 experimental server-deployable CarbonStack backbone epoch.
  - post-v0.3.0 should focus on backbone maturation and portability before Android, production vault, or hostile-server-complete work.



**New at v0.2.67:**

- `carbonstack` head is now `6755460 docs: record v0.3.0 release hardening recon`.
- `carbonstack-comms` remains at `7e745e5 scripts: add OpenMLS backbone self-test wrapper`.
- `carbonstack-cypher` remains at `13312c6 docs: standardize Cypher envelope semantics`.
- `carbonstack-os` remains at `b537475 Add CarbonStackOS north star and initial CarbonStack repository structure`.
- v0.2.67 is a pre-v0.3.0 release-hardening recon checkpoint. It is not the v0.3.0 release README implementation checkpoint, not a tag/release checkpoint, and not a new protocol/runtime behavior checkpoint.
- `carbonstack/docs/118-v0.3.0-release-hardening-recon-v0.md` now records the release-hardening scout before the v0.3.0 experimental backbone epoch.
- The next consolidated release artifact should be:
  - `carbonstack/docs/v0.3.0-minor-epoch-release.md`.
- That file should be referred to as the v0.3.0 release README in the CarbonStack repo release surface.
- The release README should answer:
  - what am I looking at;
  - what is currently proven;
  - how to run the self-test;
  - which repos are included;
  - which component heads are known-good;
  - what is not proven;
  - what security claims are explicitly not being made;
  - what builders should read next.
- `docs/118` records the current validated release candidate as:
  - CarbonStackComms OpenMLS sidecar;
  - CarbonStackCypher real local server;
  - opaque OpenMLS artifact envelope relay;
  - payload metadata validation;
  - consume-then-ack semantics;
  - OpenMLS backbone self-test wrapper.
- Current self-test public surface remains:
  - `carbonstack-comms/scripts/self-test-openmls-backbone.ps1`.
- Current lower-level implementation script remains:
  - `carbonstack-comms/scripts/smoke-openmls-real-cypher-relay.ps1`.
- `docs/118` records exact component heads at recon time:
  - `carbonstack` = `af11446` as the prior v0.2.66 state before the recon commit, with final v0.2.67 head now `6755460`;
  - `carbonstack-comms` = `7e745e5`;
  - `carbonstack-cypher` = `13312c6`;
  - `carbonstack-os` = `b537475`.
- `docs/118` states the commit subjects are not the release source of truth; the v0.3.0 release README should record exact component heads and describe what each component contributes.
- The release-hardening recon identifies files to update next:
  - `carbonstack/docs/v0.3.0-minor-epoch-release.md`;
  - `carbonstack/README.md`;
  - `carbonstack/roadmap/ROADMAP.md`;
  - `carbonstack-comms/README.md`;
  - `carbonstack-comms/scripts/smoke-openmls-real-cypher-relay.ps1`.
- Optional follow-up files if needed after diff review:
  - `carbonstack-comms/scripts/README.md`;
  - `carbonstack/docs/113-experimental-backbone-deployability-runbook-v0.md`;
  - `carbonstack/docs/114-known-good-validation-matrix-v0.md`;
  - `carbonstack/docs/117-openmls-backbone-self-test-harness-result-v0.md`.
- `docs/118` explicitly says the main `carbonstack/README.md` should remain short and front-door focused, while pointing to `docs/v0.3.0-minor-epoch-release.md`.
- `docs/118` records that `roadmap/ROADMAP.md` should stop using public-facing `Option B` / `Option C` terminology and should use:
  - OpenMLS backbone self-test harness;
  - pre-v0.3.0 release hardening;
  - v0.3.0 experimental backbone epoch.
- `docs/118` records that `carbonstack-comms/README.md` should point to the public-facing wrapper `scripts/self-test-openmls-backbone.ps1`, with `scripts/smoke-openmls-real-cypher-relay.ps1` kept as a lower-level implementation detail.
- `docs/118` records that the lower-level smoke script still printing `current known-good Option C backbone proof` is a release-surface issue because the wrapper delegates to that script and users still see its output. That wording should be changed to `current known-good OpenMLS backbone proof` or `current known-good local backbone proof`.
- Cypher requires no expected code change for the next release-surface implementation rung. Its current public docs are already aligned around opaque envelopes, no plaintext, no trust decisions, no production certification, no external audit, OpenMLS artifact content types, payload metadata, queued-only inbox, and ack behavior.
- Release nonclaims remain mandatory. The v0.3.0 release README must avoid and reject:
  - secure messenger;
  - production E2EE;
  - hostile-server safe;
  - metadata-private;
  - trustless;
  - drop-in replacement;
  - finished product;
  - certified;
  - audit-ready.
- Preferred release language remains:
  - experimental backbone;
  - known-good local proof;
  - OpenMLS backbone self-test;
  - Cypher + Comms OpenMLS relay backbone;
  - not production-certified;
  - not externally audited;
  - not a finished messenger;
  - not Android-ready.
- During v0.2.67 validation, Avast again blocked Go's generated `protocol.test.exe` under the default `%LOCALAPPDATA%\Temp\go-build...` path during `carbonstack/scripts/validate-local.ps1`.
- The stable workaround is now a dedicated Go build temp/cache root:
  - `C:\▓▓\AppData\Local\Temp\go-build-carbonstack\tmp`;
  - `C:\▓▓\AppData\Local\Temp\go-build-carbonstack\cache`.
- The recommended persistent Go env settings are:
  - `go env -w GOTMPDIR="$env:LOCALAPPDATA\Temp\go-build-carbonstack\tmp"`;
  - `go env -w GOCACHE="$env:LOCALAPPDATA\Temp\go-build-carbonstack\cache"`.
- The recommended Avast exception is narrow:
  - `C:\▓▓\AppData\Local\Temp\go-build-carbonstack\`.
- Do not exclude all of `%TEMP%`.
- Do not restore quarantined generated Go test binaries; let Go rebuild them inside the dedicated excluded folder.
- Validation then proceeded and the release-hardening recon was committed.
- Important continuity/blunder notes:
  - The local Avast/Go temp executable problem is local environment noise, not a CarbonStack logic failure.
  - A narrow temp/cache exception is now preferred over per-repo `.go-tmp` recreation when Avast interferes.
  - `docs/113` gained the link to `docs/118`; the attempted `docs/114` patch may have been a no-op depending on current text, so the final v0.2.67 commit should be treated as authority for what actually changed.
- No new Cypher route was added.
- No `/v0/artifacts` API was added.
- No top-level `carbonstack` orchestrator was added.
- No OpenMLS runtime `send` / `inbox` UX was added.
- No Android, vault, hostile-server-complete, production E2EE, external-audit, or certified-secure claim was added.
- Next safest rung:
  - v0.2.68 implement the consolidated v0.3.0 release README and supporting release-surface cleanup.
- Expected v0.2.68 scope:
  - add `docs/v0.3.0-minor-epoch-release.md`;
  - update `carbonstack/README.md`;
  - update `carbonstack/roadmap/ROADMAP.md`;
  - update `carbonstack-comms/README.md`;
  - update `carbonstack-comms/scripts/smoke-openmls-real-cypher-relay.ps1` to remove visible `Option C` wording;
  - validate with the known-good command set.
- Later:
  - final stale-claims/validation sweep;
  - v0.3.0 experimental server-deployable CarbonStack backbone epoch.


**New at v0.2.66:**

- `carbonstack` head is now `af11446 docs: update 133, experimental-backbone-deployability-runbook-v0`.
- `carbonstack` also landed `1affcf9 docs: record OpenMLS backbone self-test harness` after `f06ed27`.
- `carbonstack-comms` head is now `7e745e5 scripts: add OpenMLS backbone self-test wrapper`.
- `carbonstack-comms` also shows adjacent wrapper commit `83cafb1 scripts: add OpenMLS backbone self-test wrapper` under the current head. Treat `7e745e5` as the authoritative current repo head.
- `carbonstack-cypher` remains at `13312c6 docs: standardize Cypher envelope semantics`.
- `carbonstack-os` remains at `b537475 Add CarbonStackOS north star and initial CarbonStack repository structure`.
- v0.2.66 is the OpenMLS backbone self-test harness implementation checkpoint. It is not a runtime Comms OpenMLS `send` / `inbox` integration checkpoint and not a production release checkpoint.
- `carbonstack-comms/scripts/self-test-openmls-backbone.ps1` now exists as the public-facing local self-test wrapper.
- The self-test wrapper delegates execution to the existing lower-level real-Cypher smoke harness:
  - `scripts/smoke-openmls-real-cypher-relay.ps1`.
- The wrapper supports:
  - default targeted OpenMLS backbone self-test;
  - `-Full` pass-through for broader validation.
- Public-facing script language now frames the command as:
  - `CarbonStack OpenMLS backbone self-test harness`;
  - experimental local self-test only;
  - not a production deployment script;
  - not a finished messenger;
  - not production E2EE;
  - not hostile-server complete;
  - not metadata-private;
  - not Android-ready;
  - not externally audited or certified secure.
- The self-test validates the current known-good local backbone:
  - CarbonStackComms OpenMLS sidecar;
  - CarbonStackCypher real local server;
  - opaque OpenMLS artifact envelope relay;
  - payload metadata validation;
  - consume-then-ack semantics.
- `carbonstack-comms/scripts/README.md` now points to `scripts/self-test-openmls-backbone.ps1` as the current known-good entrypoint.
- The lower-level `scripts/smoke-openmls-real-cypher-relay.ps1` remains available as an implementation detail and lower-level smoke harness.
- No proof logic was duplicated into the wrapper. The wrapper delegates rather than reimplementing the proof.
- No OpenMLS relay path was wired into stub-era `cmd/comms send` / `inbox`.
- No top-level `carbonstack` orchestrator was added.
- `carbonstack/docs/117-openmls-backbone-self-test-harness-result-v0.md` now records the implementation result.
- `carbonstack/docs/116-openmls-backbone-self-test-harness-plan-v0.md` was updated from planning/recon wording toward implemented wrapper status/result wording.
- `carbonstack/docs/113-experimental-backbone-deployability-runbook-v0.md` and `docs/114-known-good-validation-matrix-v0.md` now point current validation commands at:
  - `powershell -ExecutionPolicy Bypass -File .\scripts\self-test-openmls-backbone.ps1`;
  - `powershell -ExecutionPolicy Bypass -File .\scripts\self-test-openmls-backbone.ps1 -Full`.
- `docs/117` records the boundary:
  - execution remains in `carbonstack-comms`;
  - `carbonstack` remains the public release/front-door documentation surface;
  - no production/security/Android/external-audit claim was added;
  - no new Cypher route was added;
  - no `/v0/artifacts` API was added;
  - no Comms runtime OpenMLS `send` / `inbox` UX was added.
- Important continuity/blunder notes:
  - `carbonstack-comms` shows two adjacent commits with the same wrapper commit subject; final head `7e745e5` is authoritative and should be treated as the actual state.
  - `carbonstack` shows `af11446 docs: update 133, experimental-backbone-deployability-runbook-v0`; despite the odd `133` in the commit subject, it is part of the v0.2.66 docs update chain after `1affcf9`.
  - The self-test wrapper intentionally stays small and calls the existing smoke harness rather than becoming a separate execution engine.
  - The Avast/Go temp-cache lesson still applies: `.go-cache/` and `.go-tmp/` are local artifacts and should remain ignored/uncommitted if used for antivirus-safe tests.
- Expected validation for this rung:
  - `powershell -ExecutionPolicy Bypass -File .\scripts\self-test-openmls-backbone.ps1`;
  - `powershell -ExecutionPolicy Bypass -File .\scripts\self-test-openmls-backbone.ps1 -Full`;
  - `powershell -ExecutionPolicy Bypass -File .\scripts\smoke-openmls-real-cypher-relay.ps1`;
  - `powershell -ExecutionPolicy Bypass -File .\scripts\check-no-rust-artifacts.ps1`;
  - `go test ./... -count=1` in `carbonstack-cypher`;
  - `powershell -ExecutionPolicy Bypass -File .\scripts\validate-local.ps1` in `carbonstack`.
- Current v0.2.66 result:
  - the public-facing self-test wrapper exists;
  - Comms remains the execution surface;
  - CarbonStack remains the release/front-door documentation surface;
  - Cypher remains unchanged from the v0.2.64 semantics checkpoint.
- Next safest rung:
  - pre-v0.3.0 release-hardening recon and cleanup planning.
- Pre-v0.3.0 should prepare:
  - v0.3.0-specific release docs;
  - component map;
  - known-good component heads;
  - validation command freeze;
  - final README/stale-claims sweep;
  - security status and nonclaims;
  - generated-artifact/cache hygiene;
  - clear statement that v0.3.0 is an experimental server-deployable backbone epoch, not a certified secure product.


**New at v0.2.65:**

- `carbonstack` head is now `f06ed27 docs: plan OpenMLS backbone self-test harness`.
- `carbonstack-comms` head is now `b21aee6 docs: plan Comms OpenMLS backbone self-test`.
- `carbonstack-cypher` remains at `13312c6 docs: standardize Cypher envelope semantics`.
- `carbonstack-os` remains at `b537475 Add CarbonStackOS north star and initial CarbonStack repository structure`.
- v0.2.65 is an OpenMLS backbone self-test harness planning checkpoint. It is not the self-test wrapper implementation checkpoint and not a runtime Comms CLI/OpenMLS `send` / `inbox` integration checkpoint.
- Public-facing language now favors:
  - `OpenMLS backbone self-test harness`;
  - `CarbonStack OpenMLS backbone self-test`;
  - `known-good local proof`;
  - `Cypher + Comms OpenMLS relay backbone`.
- The internal planning labels `Option B` / `Option C` remain useful for LogDoc continuity, but should not be used as public release-surface terminology.
- `carbonstack/docs/116-openmls-backbone-self-test-harness-plan-v0.md` now records:
  - purpose of the self-test harness plan;
  - current execution state in `carbonstack-comms`;
  - current runtime CLI boundary;
  - recommended next implementation as `scripts/self-test-openmls-backbone.ps1`;
  - non-goals;
  - public wording rules;
  - why a wrapper script should come before Go CLI/runtime integration;
  - expected v0.2.66 work;
  - local Go cache/temp note for antivirus-safe test runs.
- `docs/113-experimental-backbone-deployability-runbook-v0.md` and `docs/114-known-good-validation-matrix-v0.md` now point to `docs/116`.
- `carbonstack-comms/scripts/README.md` now replaces public “Option C validation path” wording with “current OpenMLS backbone self-test path.”
- `carbonstack-comms/scripts/README.md` now records a future self-test wrapper:
  - `scripts/self-test-openmls-backbone.ps1`;
  - wrapper should call the existing real-Cypher smoke harness instead of duplicating proof logic.
- `carbonstack-comms/.gitignore` now ignores:
  - `.go-cache/`;
  - `.go-tmp/`.
- The `.go-cache/` and `.go-tmp/` ignore entries exist because Avast previously interfered with Go-generated test executables under the normal temp path. These directories are local test/cache artifacts only and must not be committed.
- Recon confirmed the current execution surface:
  - `scripts/smoke-openmls-real-cypher-relay.ps1` remains the working lower-level proof path;
  - `scripts/smoke-openmls-real-cypher-relay.ps1 -Full` remains the broader validation path;
  - `scripts/check-no-rust-artifacts.ps1` remains the artifact guard;
  - `scripts/test-local-lifecycle.ps1` and `scripts/test-trust-lifecycle.ps1` remain older stub-era CLI/trust scaffolding, not the current OpenMLS backbone proof.
- Recon confirmed `cmd/comms` / `internal/app` are still stub-era for runtime `send`, `inbox`, and `ack`:
  - `send` still uses the mock/stub provider path;
  - `inbox` still decrypts stub plaintext;
  - `ack` is useful for the old stub lifecycle but not a polished OpenMLS runtime UX.
- Therefore, the next implementation should not wire OpenMLS relay into `internal/app send` / `inbox` yet.
- Recommended next implementation:
  - add a small Comms-hosted wrapper script, probably `scripts/self-test-openmls-backbone.ps1`;
  - support default targeted self-test and `-Full` pass-through;
  - keep the current smoke harness as the underlying implementation detail;
  - keep status/nonclaim text near the top;
  - avoid new protocol behavior.
- Validation during v0.2.65 recon:
  - initial validation failed because `GOTMPDIR` was still pointed at deleted `.go-tmp`;
  - after recreating `.go-tmp` / `.go-cache` and setting `GOTMPDIR` / `GOCACHE`, validation recovered;
  - `go test ./... -count=1` in `carbonstack-cypher` passed;
  - `go test -p 1 ./internal/relay -count=1` in `carbonstack-comms` passed;
  - targeted `internal/protocol` real/fake lifecycle tests passed;
  - `scripts/smoke-openmls-real-cypher-relay.ps1` passed;
  - `scripts/smoke-openmls-real-cypher-relay.ps1 -Full` passed;
  - `scripts/check-no-rust-artifacts.ps1` passed;
  - `carbonstack/scripts/validate-local.ps1` passed after the temp/cache workaround was restored.
- Important continuity/blunder notes:
  - deleting `.go-tmp` while leaving `GOTMPDIR` pointed at it caused Go to fail before tests could run;
  - this was local environment/cache configuration noise, not a CarbonStack code/test failure;
  - the fix was to recreate `.go-tmp` / `.go-cache`, set `GOTMPDIR` / `GOCACHE`, and later ignore those local directories;
  - public docs should not expose “Option B” as the release-facing term.
- No new Cypher route was added.
- No `/v0/artifacts` API was added.
- No OpenMLS runtime `send` / `inbox` UX was added.
- No top-level `carbonstack` orchestrator was added.
- No production/security/Android/external-audit claim was added.
- Next safest rung:
  - v0.2.66 implement the OpenMLS backbone self-test wrapper in `carbonstack-comms`.
- Expected v0.2.66 scope:
  - add `scripts/self-test-openmls-backbone.ps1`;
  - call `scripts/smoke-openmls-real-cypher-relay.ps1`;
  - support `-Full`;
  - update `carbonstack-comms/scripts/README.md`;
  - update `carbonstack/docs/113`, `docs/114`, and possibly `docs/116`;
  - validate using the existing known-good command set.
- Later:
  - pre-v0.3.0 release hardening;
  - v0.3.0 experimental server-deployable CarbonStack backbone epoch.


**New at v0.2.64:**

- `carbonstack` head is now `6223e1c docs: define envelope lifecycle semantics`.
- `carbonstack-comms` head is now `ad141f8 test: standardize Comms relay ack semantics`.
- `carbonstack-cypher` head is now `13312c6 docs: standardize Cypher envelope semantics`.
- `carbonstack-os` remains at `b537475 Add CarbonStackOS north star and initial CarbonStack repository structure`.
- v0.2.64 is an inbox/ack/general semantics and schema/API standardization checkpoint. It is not Option B CLI/dev-harness implementation and not a production release checkpoint.
- `carbonstack/docs/115-envelope-lifecycle-semantics-v0.md` now defines the current canonical envelope lifecycle language:
  - envelope;
  - opaque payload;
  - content type;
  - protocol version;
  - payload metadata;
  - inbox;
  - queued;
  - sidecar consume;
  - ack;
  - acknowledged;
  - what Cypher knows;
  - what Comms knows;
  - what the OpenMLS sidecar proves;
  - what payload metadata does not prove.
- `docs/113-experimental-backbone-deployability-runbook-v0.md` and `docs/114-known-good-validation-matrix-v0.md` now point to `docs/115`.
- Current safe lifecycle rule is now explicit:
  - download is not ack;
  - artifact write is not ack;
  - payload metadata validation is not ack;
  - sidecar consume success permits ack.
- Cypher ack semantics are now standardized:
  - `queued` means the envelope is available through the recipient inbox route;
  - `acknowledged` means Cypher accepted a recipient-device ack for the envelope;
  - inbox returns queued envelopes only;
  - ack is idempotent for the same recipient so retries after a lost HTTP response are safe;
  - wrong-recipient ack is rejected;
  - unknown-envelope ack is rejected;
  - missing-recipient ack is rejected.
- Important boundary: Cypher records ack. Comms decides when to ack. The OpenMLS sidecar proves consume success. Cypher does not know sidecar consume state.
- `carbonstack-cypher/internal/httpapi/api.go` now implements idempotent same-recipient ack behavior.
- `carbonstack-cypher/internal/httpapi/api_test.go` now covers:
  - idempotent same-recipient ack;
  - unknown-envelope ack rejection;
  - missing-recipient ack rejection;
  - existing wrong-recipient ack behavior remains preserved.
- Cypher public/current docs were standardized around current semantics:
  - `docs/02-envelope-model.md`;
  - `docs/03-api-surface.md`;
  - `docs/04-storage-model.md`;
  - `docs/07-data-model-v0.md`;
  - `docs/08-api-contract-v0.md`.
- Current payload metadata language is now consistent:
  - `payload_sha256` and `payload_size_bytes` describe decoded `ciphertext_b64` bytes;
  - metadata is relay/debug/storage sanity metadata;
  - metadata is not a trust root;
  - a malicious server can lie about server-returned metadata;
  - OpenMLS sidecar consume remains the cryptographic validity gate.
- Comms fake Cypher test servers were updated to better mirror real Cypher behavior:
  - `internal/relay/test_cypher_server_test.go`;
  - `internal/protocol/openmls_sidecar_test_cypher_server_test.go`.
- The fake servers now include payload metadata in submit/inbox responses, reject invalid base64, implement ack routes, and return queued envelopes only from inbox.
- `internal/protocol/openmls_sidecar_relay_test.go` now consume-then-acks the fake-server full lifecycle:
  - KeyPackage envelope ack after `conversation-add-member` succeeds;
  - Welcome envelope ack after `conversation-join` succeeds;
  - application-message envelope ack after `message-open` succeeds;
  - Bob inbox expectations now assume the Welcome is acked before the application-message is queued/read.
- Validation passed after Avast temporarily blocked Go’s generated `protocol.test.exe` under `%TEMP%`:
  - initial failures were local antivirus/Go temp executable noise, not project logic failures;
  - workaround used project-local `.go-tmp` and `.go-cache` via `GOTMPDIR` and `GOCACHE`;
  - targeted Comms protocol proof passed after cache relocation;
  - normal smoke harness passed;
  - `-Full` smoke harness passed;
  - artifact guard passed.
- `go test ./... -count=1` in `carbonstack-cypher` passed.
- `go test -p 1 ./internal/relay -count=1` in `carbonstack-comms` passed.
- `go test -p 1 ./internal/protocol -run "TestOpenMLSSidecarFullLifecycleRelayThroughCypherEnvelope|TestOpenMLSSidecarFullLifecycleRelayThroughRealCypherServer" -count=1 -timeout 300s` passed after local Go temp/cache relocation.
- `powershell -ExecutionPolicy Bypass -File .\scripts\smoke-openmls-real-cypher-relay.ps1` passed.
- `powershell -ExecutionPolicy Bypass -File .\scripts\smoke-openmls-real-cypher-relay.ps1 -Full` passed.
- `powershell -ExecutionPolicy Bypass -File .\scripts\check-no-rust-artifacts.ps1` passed.
- `powershell -ExecutionPolicy Bypass -File .\scripts\validate-local.ps1` was affected by the same Avast-generated-test-exe issue before the workaround, and later validation/commit proceeded after the affected Comms protocol path passed. Treat the antivirus block as local environment noise rather than project failure.
- Local `.go-cache/` and `.go-tmp/` directories were created for the Avast workaround and then removed before the breakpoint so the repo snapshot was clean.
- Next small hygiene task: if those directories are recreated for future tests, add `.go-cache/` and `.go-tmp/` to `carbonstack-comms/.gitignore` before the next breakpoint.
- Important continuity/blunder notes:
  - Avast falsely/heuristically flagged Go-generated `protocol.test.exe` under `%TEMP%`;
  - the issue cleared when using project-local Go temp/cache directories;
  - an earlier validation command was pasted as `validate-local.ps1cd ...`, causing a malformed PowerShell command; this was operator paste error, not project failure;
  - `.go-cache/` and `.go-tmp/` appeared as untracked after the workaround and were removed before checkpointing.
- No new public production/security/Android/external-audit claim was added.
- No new Cypher route was added.
- No `/v0/artifacts` API was added.
- No Comms runtime `send` / `inbox` OpenMLS UX was added.
- No top-level `carbonstack` orchestrator was added.
- Next safest rung:
  - v0.2.65+ Option B CLI/dev-harness planning and implementation.
- Before or during v0.2.65:
  - recreate and ignore `.go-cache/` and `.go-tmp/` if future Avast-safe test runs need local Go cache/temp directories;
  - keep Option B framed as a dev harness, not a finished messenger UX;
  - keep execution in `carbonstack-comms` unless a later pre-v0.3.0 hardening pass proves a top-level orchestrator can stay simple.


**New at v0.2.63:**

- `carbonstack` head is now `74bd964 docs: record known-good backbone validation matrix`.
- `carbonstack-comms` head is now `7152e0b docs: clarify Comms backbone smoke scripts`.
- `carbonstack-cypher` remains at `baac52d docs: clean public Cypher surface`.
- `carbonstack-os` remains at `b537475 Add CarbonStackOS north star and initial CarbonStack repository structure`.
- v0.2.63 is an Option C / known-good validation cleanup checkpoint, not a new protocol/runtime behavior checkpoint.
- The current execution entrypoint remains in `carbonstack-comms`:
  - `scripts/smoke-openmls-real-cypher-relay.ps1`.
- `carbonstack` remains the release/front-door documentation repo:
  - public explanation;
  - known-good validation matrix;
  - runbook;
  - component map;
  - security nonclaims.
- The smoke script wording was updated so it no longer frames itself only as the older v0.2.56 real-server proof.
- The smoke script now describes the current known-good Option C backbone proof:
  - real local Cypher server;
  - temp SQLite DB;
  - OpenMLS KeyPackage -> Welcome -> application-message relay;
  - payload metadata validation before artifact write;
  - ack only after successful sidecar consume;
  - final `message-open` plaintext recovery.
- The smoke script already had useful failure/process hygiene and was preserved:
  - `$ErrorActionPreference = "Stop"`;
  - `Invoke-NativeCommand` throws on nonzero native exit codes;
  - existing `cypher` processes are detected before test startup;
  - stale-process cleanup command is printed instead of silently killing arbitrary processes;
  - `-Full` runs broader relay/protocol/package validation;
  - generated Rust/OpenMLS artifact guard runs after targeted smoke and again after `-Full`.
- `carbonstack-comms/scripts/README.md` now records:
  - current known-good entrypoint;
  - broader `-Full` validation;
  - generated Rust/build artifact guard;
  - stale Cypher process warning;
  - older local/trust lifecycle scripts as earlier scaffolding, not the current OpenMLS + Cypher backbone proof.
- `carbonstack/docs/114-known-good-validation-matrix-v0.md` now records:
  - purpose and boundary of the current validation matrix;
  - current known-good proof shape;
  - primary validation commands for Cypher, Comms smoke, Comms full validation, and CarbonStack docs;
  - validation layers;
  - current execution boundary;
  - what the proof does and does not prove;
  - next rung: inbox/ack/general semantics and schema/API wording standardization.
- `carbonstack/docs/113-experimental-backbone-deployability-runbook-v0.md` now points to:
  - `docs/114-known-good-validation-matrix-v0.md`.
- `docs/113` also updates the near-term list:
  - public README/surface cleanup complete at v0.2.62;
  - known-good validation matrix complete at v0.2.63;
  - next: inbox/ack/general semantics and schema standardization;
  - then Option B CLI/dev-harness planning/implementation;
  - then pre-v0.3.0 release hardening.
- The current known-good proof remains:
  - Bob exports an OpenMLS KeyPackage;
  - Cypher relays it to Alice;
  - Alice consumes it and creates a Welcome;
  - Cypher relays the Welcome to Bob;
  - Bob consumes it and joins;
  - Alice protects an application-message;
  - Cypher relays it to Bob;
  - Bob validates payload metadata before artifact write;
  - Bob consumes the application-message through the sidecar;
  - plaintext matches;
  - envelopes are acked only after successful sidecar consume.
- No component runtime behavior changed in v0.2.63.
- No new Cypher route was added.
- No Comms runtime `send` / `inbox` OpenMLS UX was added.
- No top-level `carbonstack` orchestration script was added.
- No production/security/Android/external-audit claim was added.
- Important continuity/blunder note:
  - recon found the smoke harness was already structurally good enough for Option C;
  - the main mismatch was wording drift from the older v0.2.56 proof label;
  - the fix stayed small and avoided unnecessary new execution surfaces.
- Validation expected for this rung:
  - `go test ./... -count=1` in `carbonstack-cypher`;
  - `powershell -ExecutionPolicy Bypass -File .\scripts\smoke-openmls-real-cypher-relay.ps1` in `carbonstack-comms`;
  - `powershell -ExecutionPolicy Bypass -File .\scripts\smoke-openmls-real-cypher-relay.ps1 -Full` in `carbonstack-comms`;
  - `powershell -ExecutionPolicy Bypass -File .\scripts\check-no-rust-artifacts.ps1` in `carbonstack-comms`;
  - `powershell -ExecutionPolicy Bypass -File .\scripts\validate-local.ps1` in `carbonstack`.
- Next safest rung:
  - v0.2.64 inbox/ack/general semantics and schema/API wording standardization.
- Later:
  - Option B CLI/dev-harness planning and implementation;
  - pre-v0.3.0 release-hardening checkpoint;
  - v0.3.0 experimental server-deployable CarbonStack backbone epoch.


**New at v0.2.38:**


- `carbonstack` head is now `0135902 docs: record OpenMLS multi-message continuity result`.
- `carbonstack-comms` head is now `27dc33d feat: add OpenMLS sidecar explicit message labels`.
- `docs/78-openmls-sidecar-multi-message-continuity-result-v0.md` landed at `0135902`.
- `message-protect` now accepts `--message-label <safe-label>`.
- `message-open` now accepts `--message-label <safe-label>`.
- Omitted `--message-label` defaults to `message-0001`, preserving v0.2.36 one-message compatibility.
- `validate_message_label(...)` rejects unsafe filesystem labels and reserved/internal names.
- `protect_dev_message(...)` and `open_dev_message(...)` now accept message labels and use them for artifact/summary paths.
- Manual proof validated:
  - `message-0001` / `"hello bob 1"`;
  - `message-0002` / `"hello bob 2"`.
- Go contract proof validated `TestOpenMLSSidecarMessageProtectOpenTwoSequentialMessages`.
- The sidecar now validates `create -> add-member -> Welcome export -> join -> protect/open message-0001 -> protect/open message-0002`.
- Alice-global/Bob-device-scoped state asymmetry remains intentionally preserved.
- Out-of-order, replay, bidirectional messages, generated message IDs, Cypher routing, Comms runtime integration, and trust mutation remain not validated.


- `carbonstack` head is now `0135902 docs: record OpenMLS multi-message continuity result`.
- `carbonstack-comms` head is now `27dc33d feat: add OpenMLS sidecar explicit message labels`.
- `docs/76-openmls-sidecar-multi-message-continuity-plan-v0.md` landed at `680ce12`.
- `docs/77-openmls-sidecar-multi-message-api-recon-v0.md` landed at `f6266d2`.
- Provider-info raw hand-formatted JSON was replaced with structured JSON via `serde_json::json!`, `CAPABILITIES`, and `UNSUPPORTED_COMMANDS`.
- Multi-message continuity recon confirms:
  - `create_message` still returns `MlsMessageOut` and encrypts application data.
  - `create_message` checks active state, rejects pending proposals, creates authenticated application content, encrypts it, resets AAD, and returns a private-message carrier.
  - `process_message` routes through `unprotect_message`, decrypts private messages, parses them, and writes message secrets back to provider storage when private-message processing modifies the secret tree.
  - Save-after-protect and save-after-open remain correct and must not be weakened.
  - The next implementation should add explicit `--message-label` and prove two sequential messages.
- Broad OpenMLS crate recon (`3A`) was omitted because it was too large/noisy; targeted recon of `application.rs`, `processing.rs`, current sidecar `state.rs/main.rs`, and the Go contract test was sufficient.
- v0.2.38 added explicit message-label implementation and two-message continuity validation.
- v0.2.38 is a full implementation breakpoint and should be preserved for handoff continuity.


**New at v0.2.39:**

- `carbonstack` head is now `01bc4a8 docs: record OpenMLS corrupt message artifact behavior`.
- `carbonstack-comms` head is now `80685a7 test: cover OpenMLS sidecar message ordering replay cases`.
- `docs/79-openmls-sidecar-message-ordering-replay-plan-v0.md` landed at `01cd442`.
- `docs/80-openmls-sidecar-message-ordering-replay-api-recon-v0.md` landed at `9411aeb`.
- `docs/80` was updated at `01bc4a8` to record the corrected corrupt/truncated artifact behavior.
- Go contract tests now cover:
  - `TestOpenMLSSidecarMessageOpenOutOfOrderTwoMessageDelivery`;
  - `TestOpenMLSSidecarMessageOpenDuplicateRejected`;
  - `TestOpenMLSSidecarMessageOpenCorruptArtifactRejected`.
- Helper functions added to `openmls_sidecar_provider_info_test.go`:
  - `setupOpenMLSTwoMemberConversation(...)`;
  - `protectOpenMLSSidecarMessage(...)`;
  - `openOpenMLSSidecarMessage(...)`.
- Validated behavior:
  - protect `message-0001`, protect `message-0002`, open `message-0002`, then open `message-0001` succeeds for the current same-sender Alice-to-Bob dev-local flow;
  - duplicate/replay opening the same message artifact fails with `ValidationError(UnableToDecrypt(SecretTreeError(SecretReuseError)))`, surfaced as `message_open_failed` / `checkpoint.failed` / exit code `3`;
  - truncated/corrupt `application-message.bin` fails before message processing with `message_artifact_invalid` / `provider.message.invalid` / `EndOfStream` / exit code `3`.
- No Rust sidecar feature change was required for v0.2.39; this rung is primarily docs/recon + Go contract-test coverage.
- Remaining not validated: long skipped-message windows, multi-sender ordering, epoch-change ordering, membership-change ordering, wrong-device/wrong-conversation behavior, Comms runtime integration, Cypher routing, trust-state mutation, and Alice device-scoped state migration.


**New at v0.2.40:**

- `carbonstack` head is now `2d1d657 docs: record OpenMLS Alice device-scoped state result`.
- `carbonstack-comms` head is now `7ccf545 refactor: use device-scoped OpenMLS conversation state`.
- `docs/82-openmls-sidecar-alice-device-scoped-state-layout-plan-v0.md` planned the hard-cut dev-state layout cleanup.
- `docs/83-openmls-sidecar-alice-device-scoped-state-recon-v0.md` recorded the path/helper/call-site recon for the cleanup.
- `docs/84-openmls-sidecar-alice-device-scoped-state-result-v0.md` records the result. The final snapshot shows two adjacent `docs: record OpenMLS Alice device-scoped state result` commits; the current head `2d1d657` is the authority.
- `conversation-create`, `conversation-load-check`, `conversation-add-member`, and `message-protect` now use device-scoped creator conversation state.
- Alice/creator state moved from:
  - `.carbonstack-openmls-sidecar-state/dev/conversations/<conversation-label>/`
  to:
  - `.carbonstack-openmls-sidecar-state/dev/devices/<device-label>/conversations/<conversation-label>/`
- Bob `conversation-join` and `message-open` were already device-scoped and remain so.
- The command surface did not change. Path hints changed.
- The cleanup is a hard-cut dev-state break; no old global-path migration compatibility exists.
- Go test `TestOpenMLSSidecarConversationCreate` had a stale hardcoded old global path and was patched to follow envelope path hints.
- The stale local `stateDir` variable caused a Go compile failure and was removed.
- Core v0.2.39 behavior remains the validation target: two-message continuity, out-of-order two-message open, duplicate/replay rejection, and corrupt/truncated artifact rejection.
- Next safest rung: Cypher MLS artifact routing design docs/recon.


**New at v0.2.41:**

- `carbonstack` head is now `45851ef docs: record OpenMLS sidecar Phase 2D mainline closure`.
- `carbonstack-comms` head is now `65c202b test: cover OpenMLS sidecar bidirectional message flow`.
- `docs/85-openmls-sidecar-phase2d-closure-checklist-v0.md` landed at `7013406` and was updated at `f75b529` to record wrong-target behavior.
- `docs/86-openmls-sidecar-phase2d-mainline-closure-result-v0.md` landed at `45851ef`.
- Stale-path sweep after v0.2.40 found one remaining operational old-path bug: `protect_dev_message` read Alice device-scoped provider storage but still wrote protected message artifacts through old global conversation message helpers.
- `cf8ac20` fixed `protect_dev_message` to write message artifacts under `dev/devices/<sender-device>/conversations/<conversation>/messages/<message-label>/`.
- Go contract tests added:
  - `TestOpenMLSSidecarMessageOpenWrongDeviceRejected`;
  - `TestOpenMLSSidecarMessageOpenWrongConversationRejected`;
  - `TestOpenMLSSidecarMessageProtectOpenBidirectional`.
- Validated wrong-device behavior:
  - Eve has identity state but no joined conversation provider storage;
  - Eve attempts to open Alice's valid artifact under the real conversation;
  - sidecar returns `ok=false`, `conversation_or_message_missing`, `provider.conversation.missing`, warning severity, `trust_relevant=false`, exit code `3`.
- Validated wrong-conversation behavior:
  - Bob has joined `carbonstack-test-conversation`;
  - Bob attempts to open Alice's valid artifact under `carbonstack-wrong-conversation`;
  - sidecar returns `ok=false`, `conversation_or_message_missing`, `provider.conversation.missing`, warning severity, `trust_relevant=false`, exit code `3`.
- Validated bidirectional flow:
  - Alice protects `alice-message-0001`, Bob opens it and recovers `"hello bob from alice"`;
  - Bob protects `bob-message-0001`, Alice opens it and recovers `"hello alice from bob"`;
  - Bob's protected message artifact lives under Bob's device-scoped conversation tree;
  - Alice's creator state after add-member/merge can process Bob's private application message.
- Phase 2D is now mainline-closed for research continuity: complete enough to proceed to Cypher minimal opaque MLS artifact relay research.
- Remaining Phase 2D work is polish/future revisits, not a blocker for Cypher research:
  - remove dead old global helper definitions;
  - split the large Go sidecar test file;
  - refine replay/secret-reuse taxonomy;
  - add missing-artifact / invalid-label explicit tests if desired;
  - explore long skipped-message, multi-sender, and membership-change matrices.
- Local validation used low parallelism after a Windows BSOD during heavy test/build activity. Debugging showed a local `WIN32K_CRITICAL_FAILURE (0x164)` involving `dwm.exe` / `win32kbase.sys`, with Vanguard and Avast kernel drivers loaded. Treat this as local machine stability/environment noise, not a CarbonStack project failure.


**New at v0.2.42:**

- `carbonstack` head is now `314a3d5 docs: plan OpenMLS sidecar maintainability promotion`.
- `carbonstack-comms` remains at `65c202b test: cover OpenMLS sidecar bidirectional message flow`.
- `carbonstack-cypher` remains at `0bfd5af chore: remove tracked Cypher local state artifacts`.
- `carbonstack-os` remains at `b537475 Add CarbonStackOS north star and initial appliance model`.
- v0.2.42 is a docs-only maintainability / promotion planning checkpoint. No sidecar code, Go tests, Cypher code, or runtime integration changed.
- New docs landed:
  - `docs/87-openmls-sidecar-current-state-index-v0.md`;
  - `docs/88-openmls-sidecar-maintainability-promotion-plan-v0.md`;
  - `docs/89-openmls-sidecar-module-split-plan-v0.md`;
  - `docs/90-openmls-sidecar-test-suite-split-plan-v0.md`;
  - `docs/91-openmls-sidecar-artifact-ownership-map-v0.md`;
  - `docs/92-openmls-sidecar-command-schema-matrix-v0.md`.
- The plan explicitly preserves `internal/protocol/mls/research/openmls-sidecar` as the known-good Phase 2D research reference.
- The planned promoted sidecar path is `internal/protocol/mls/openmls-sidecar` unless a later implementation rung deliberately chooses another nearby path.
- The promotion principle is: promoted means maintained implementation scaffold, not production E2EE, not secure vault, not user release, not Android, and not Cypher routing.
- The planned cleanup ladder is now:
  - v0.2.43: copy/promote sidecar scaffold above research with behavior identical; COMPLETE at b44adbd/2c0f576;
  - v0.2.44: split promoted Rust sidecar modules with no behavior change; NEXT;
  - v0.2.45: split Go sidecar contract tests and point them at the promoted sidecar;
  - v0.2.46: update READMEs, current-state docs, known-good commands, stale warnings;
  - v0.2.47: begin Cypher minimal opaque MLS artifact relay recon. COMPLETE at 40ccdfa.
- The maintainability recon records the current largest hotspots:
  - `internal/protocol/openmls_sidecar_provider_info_test.go` is about 2,938 lines;
  - `internal/protocol/mls/research/openmls-sidecar/src/main.rs` is about 2,718 lines;
  - `internal/protocol/mls/research/openmls-sidecar/src/state.rs` is about 2,261 lines;
  - `internal/protocol/mls/research/openmls-sidecar/README.md` is stale and still describes an older provider-info-only state.
- v0.2.42 adds a command/schema matrix for current sidecar commands and failure shapes.
- v0.2.42 adds an artifact ownership map distinguishing future relay candidates (`public-bundle.keypackage.bin`, `welcome.bin`, `application-message.bin`) from local-only sensitive state (`signer.json`, `provider-storage.json`, raw MemoryStorage/group state).
- v0.2.42 explicitly defers Cypher routing until after the maintainability promotion and cleanup ladder.



**New at v0.2.43:**

- `carbonstack` head is now `2c0f576 docs: record OpenMLS sidecar promotion scaffold`.
- `carbonstack-comms` head is now `b44adbd refactor: promote OpenMLS sidecar scaffold`.
- `docs/93-openmls-sidecar-promotion-scaffold-result-v0.md` records the behavior-preserving scaffold promotion result.
- Promoted maintained sidecar scaffold now exists at:
  - `carbonstack-comms/internal/protocol/mls/openmls-sidecar`
- Known-good Phase 2D research reference remains intact at:
  - `carbonstack-comms/internal/protocol/mls/research/openmls-sidecar`
- Go sidecar contract tests now target the promoted scaffold via:
  - `openMLSSidecarDir = "mls/openmls-sidecar"`
  - `openMLSSidecarStateDir = "mls/openmls-sidecar/.carbonstack-openmls-sidecar-state"`
- The promoted scaffold initially carries the same Rust structure as the research sidecar:
  - `Cargo.toml`;
  - `Cargo.lock`;
  - `README.md` still present and still scheduled for v0.2.46 cleanup;
  - `src/main.rs`;
  - `src/state.rs`;
  - `src/provider.rs`;
  - `src/labels.rs`.
- Validation confirmed the promoted sidecar compiles and tests:
  - `cargo fmt`;
  - `cargo check`;
  - `cargo test`;
  - Go Phase 2D contract tests from `carbonstack-comms` using `go test -p 1`.
- Rust warnings remain for dead old global conversation helper definitions in `state.rs`. These are known residue and are intentionally deferred to v0.2.44 module/cleanup work.
- Behavior was not intentionally changed:
  - command names unchanged;
  - JSON envelope shape unchanged;
  - path-hint semantics unchanged;
  - event names unchanged;
  - exit codes unchanged;
  - OpenMLS lifecycle unchanged.
- Important blunder/repair:
  - The first local `refactor: promote OpenMLS sidecar scaffold` commit accidentally staged generated `.carbonstack-openmls-sidecar-state` dev artifacts under the promoted sidecar, including `signer.json`, `provider-storage.json`, `welcome.bin`, `application-message.bin`, public-bundle artifacts, and other generated state.
  - The issue was caught before `git push`.
  - The local commit was reset, generated state/build output was removed, ignore coverage was added for `internal/protocol/mls/**/.carbonstack-openmls-sidecar-state/` and `internal/protocol/mls/**/target/`, and the clean scaffold commit was recreated as `b44adbd`.
  - No generated secret-bearing sidecar state was pushed to the remote according to the final snapshot.
- Next rung:
  - v0.2.44 split promoted Rust modules without behavior changes.


**New at v0.2.44:**

- `carbonstack` head is now `2648f93 docs: record OpenMLS sidecar Rust module split`.
- `carbonstack-comms` head is now `9024447 refactor: split OpenMLS sidecar CLI parsers`.
- `docs/94-openmls-sidecar-rust-module-split-result-v0.md` records the behavior-preserving Rust module split result.
- v0.2.44 completed the first promoted-sidecar Rust maintainability split while preserving command behavior.
- Promoted maintained sidecar remains active at:
  - `carbonstack-comms/internal/protocol/mls/openmls-sidecar`
- Known-good Phase 2D research reference remains intact at:
  - `carbonstack-comms/internal/protocol/mls/research/openmls-sidecar`
- New/active promoted Rust modules after v0.2.44:
  - `src/cli.rs`;
  - `src/envelope.rs`;
  - `src/labels.rs`;
  - `src/main.rs`;
  - `src/paths.rs`;
  - `src/provider.rs`;
  - `src/schema.rs`;
  - `src/state.rs`.
- `paths.rs` now owns current device-scoped path construction:
  - device state root paths;
  - identity/public-bundle paths;
  - device-scoped conversation paths;
  - Welcome/add-member paths;
  - message artifact/manifest/protect/open summary paths.
- Old global `dev/conversations/<conversation-label>/` helper residue was removed from the promoted sidecar during the path split. The research sidecar preserves the historical reference.
- `schema.rs` now owns public result structs:
  - `IdentityCreateResult`;
  - `IdentityStatusResult`;
  - `PublicBundleExportResult`;
  - `ConversationCreateResult`;
  - `ConversationLoadCheckResult`;
  - `ConversationAddMemberResult`;
  - `ConversationJoinResult`;
  - `MessageProtectResult`;
  - `MessageOpenResult`.
- `state.rs` re-exports these result structs where needed so the split stayed low-risk.
- `envelope.rs` now owns:
  - provider-info rendering;
  - shared phase constants;
  - supported/unsupported command lists;
  - shared warning strings;
  - `json_escape`;
  - identity command printers;
  - public-bundle command printers;
  - unsupported-command printer.
- Conversation and message printers remain in `main.rs` intentionally; they can be moved later after Go test split or command-family splits.
- `cli.rs` now owns simple CLI parsing:
  - `parse_device_label`;
  - `parse_conversation_label`;
  - `parse_member_keypackage_path`;
  - `parse_welcome_artifact_path`;
  - `parse_plaintext`;
  - `parse_message_artifact_path`;
  - `parse_message_label`;
  - `parse_write_artifact_flag`.
- `main.rs` is still not fully minimal, but it is materially cleaner than the promoted scaffold:
  - dispatch and command handlers remain;
  - conversation/message envelope printers remain;
  - path helpers, public result schemas, selected envelope printers, provider-info constants, and CLI parsers are split out.
- Validation target remained behavior-preserving:
  - `cargo check`;
  - `cargo test`;
  - `cargo run -- provider-info`;
  - `go test -p 1 ./internal/protocol`;
  - `go test -p 1 ./...`;
  - `powershell -ExecutionPolicy Bypass -File .\scripts\check-no-rust-artifacts.ps1`.
- Important blunders/repair:
  - Broad regex deletion during the `paths.rs` split temporarily removed the `ConversationJoinResult` / `join_dev_conversation` section from promoted `state.rs`.
  - The missing join section was restored from the untouched research sidecar reference.
  - Leftover old global message helper residue caused duplicate/missing helper errors and was removed.
  - The first provider-envelope split attempt duplicated phase constants and provider-info helpers between `main.rs` and `envelope.rs`.
  - The final provider-envelope split preserved the existing provider-info JSON shape instead of silently changing schema.
  - During the CLI split, parser-removal regex left an orphaned tail (`index += 1; ... None`) in `main.rs`; it was removed before validation.
  - Some validation failures were working-directory or paste-concatenation issues (`provider-infocd`, Go tests run from the sidecar crate). These were environment/operator blunders, not sidecar behavior regressions.
  - A final docs commit attempt initially failed from the umbrella folder because the umbrella is not a Git repo; the docs commit was completed from `carbonstack`.
- No command names, JSON envelope schema, path-hint semantics, event names, exit codes, or OpenMLS lifecycle behavior were intentionally changed.
- Next rung:
  - v0.2.45 split Go sidecar contract tests and keep them targeting the promoted sidecar.
- Remaining cleanup ladder:
  - v0.2.45: split Go sidecar tests; COMPLETE at 6a3a24b/a062dc5;
  - v0.2.46: README/current-state/known-good command/stale-warning cleanup; COMPLETE at a371a96/687f6ad;
  - v0.2.47: begin Cypher minimal opaque MLS artifact relay recon. NEXT.


**New at v0.2.45:**

- `carbonstack` head is now `a062dc5 docs: record OpenMLS sidecar Go test split`.
- `carbonstack-comms` head is now `6a3a24b test: split OpenMLS sidecar message tests`.
- `docs/95-openmls-sidecar-go-test-suite-split-result-v0.md` records the behavior-preserving Go contract test suite split.
- v0.2.45 completed the OpenMLS sidecar Go test maintainability split while preserving coverage and keeping tests pointed at the promoted sidecar.
- The promoted maintained sidecar remains active at:
  - `carbonstack-comms/internal/protocol/mls/openmls-sidecar`
- The known-good Phase 2D research reference remains intact at:
  - `carbonstack-comms/internal/protocol/mls/research/openmls-sidecar`
- The Go sidecar contract tests are now split by ownership:
  - `internal/protocol/openmls_sidecar_helpers_test.go`;
  - `internal/protocol/openmls_sidecar_provider_info_test.go`;
  - `internal/protocol/openmls_sidecar_identity_test.go`;
  - `internal/protocol/openmls_sidecar_public_bundle_test.go`;
  - `internal/protocol/openmls_sidecar_conversation_test.go`;
  - `internal/protocol/openmls_sidecar_message_test.go`;
  - `internal/protocol/openmls_sidecar_message_negative_test.go`.
- `openmls_sidecar_helpers_test.go` owns shared runner/parsing/assertion/setup helpers, including sidecar invocation helpers, envelope parsing, file assertions, two-member Alice/Bob setup, message protect/open wrappers, and shared message success assertions.
- `openmls_sidecar_provider_info_test.go` owns provider-info and unsupported-command envelope tests.
- `openmls_sidecar_identity_test.go` owns identity-create / identity-status tests only.
- `openmls_sidecar_public_bundle_test.go` owns public-bundle summary/export/artifact tests.
- `openmls_sidecar_conversation_test.go` owns conversation-create, conversation-load-check, add-member/Welcome export, and conversation-join/Welcome consume tests.
- `openmls_sidecar_message_test.go` owns positive/success message tests: one-way protect/open, two sequential messages, bidirectional Alice/Bob message flow, and out-of-order same-sender two-message delivery.
- `openmls_sidecar_message_negative_test.go` owns wrong-device, wrong-conversation, duplicate/replay, and corrupt/truncated artifact negative tests.
- Behavior was not intentionally changed:
  - sidecar command names unchanged;
  - JSON envelope shape unchanged;
  - path-hint semantics unchanged;
  - event names unchanged;
  - exit codes unchanged;
  - OpenMLS lifecycle unchanged.
- Important blunders/repair:
  - The first provider-info extraction attempt was run before the expected identity target file existed. PowerShell continued after a failed `Get-Content`, causing an incomplete provider-info file and empty/null identity target candidate. The files were restored from Git and the extraction was redone with hard guards.
  - During the message test split, `openmls_sidecar_message_test.go` initially missed `path/filepath`, then retained unused `os` while needing `strings`. The import block was repaired with a direct regex replacement and validation continued.
  - These were test-split/operator scripting issues, not sidecar behavior regressions.
- Validation target remained behavior-preserving:
  - `go test -p 1 ./internal/protocol`;
  - `go test -p 1 ./...`;
  - `powershell -ExecutionPolicy Bypass -File .\scripts\check-no-rust-artifacts.ps1`.
- Next rung:
  - v0.2.46 README/current-state/known-good command/stale-warning cleanup.
- Remaining cleanup ladder:
  - v0.2.46: README/current-state/known-good command/stale-warning cleanup; COMPLETE at a371a96/687f6ad;
  - v0.2.47: begin Cypher minimal opaque MLS artifact relay recon. NEXT.


**New at v0.2.46:**

- `carbonstack` head is now `687f6ad docs: update OpenMLS sidecar current state cleanup`.
- `carbonstack-comms` head is now `a371a96 docs: update OpenMLS sidecar README and stale warnings`.
- `docs/96-openmls-sidecar-readme-current-state-cleanup-result-v0.md` records the README/current-state cleanup result.
- v0.2.46 completed the planned maintainability cleanup ladder after:
  - v0.2.43 promoted the sidecar scaffold outside the research path;
  - v0.2.44 split promoted Rust sidecar modules;
  - v0.2.45 split Go sidecar contract tests by ownership;
  - v0.2.46 updated README/current-state docs and stale warnings.
- The promoted maintained sidecar remains active at:
  - `carbonstack-comms/internal/protocol/mls/openmls-sidecar`
- The known-good Phase 2D research reference remains intact at:
  - `carbonstack-comms/internal/protocol/mls/research/openmls-sidecar`
- The promoted sidecar README now accurately records:
  - current supported commands;
  - unsupported `state-checkpoint` and `state-load-check`;
  - dev-local / experimental status;
  - non-production E2EE status;
  - generated sidecar state safety warnings;
  - device-scoped state layout;
  - split Go test ownership map;
  - validation commands.
- The research sidecar README now starts with a frozen-reference notice and points active work to the promoted sidecar.
- Stale runtime/help/warning text was corrected where it still implied implemented commands were unsupported.
- `docs/87-openmls-sidecar-current-state-index-v0.md` now records the v0.2.43-v0.2.46 maintainability ladder state.
- `docs/88-openmls-sidecar-maintainability-promotion-plan-v0.md` now marks the maintainability plan as landed through v0.2.45 and v0.2.46 cleanup.
- `docs/89-openmls-sidecar-module-split-plan-v0.md` now records the v0.2.44 Rust module split result note.
- `docs/90-openmls-sidecar-test-suite-split-plan-v0.md` now records the v0.2.45 Go test split result note.
- Behavior was not intentionally changed:
  - sidecar command names unchanged;
  - JSON envelope shape unchanged;
  - path-hint semantics unchanged;
  - event names unchanged;
  - exit codes unchanged;
  - OpenMLS lifecycle unchanged.
- Validation target remained behavior-preserving:
  - `cargo check`;
  - `cargo test`;
  - `cargo run -- provider-info`;
  - `go test -p 1 ./internal/protocol`;
  - `go test -p 1 ./...`;
  - `powershell -ExecutionPolicy Bypass -File .\scripts\check-no-rust-artifacts.ps1`;
  - `powershell -ExecutionPolicy Bypass -File .\scripts\validate-local.ps1`.
- Important continuity notes:
  - v0.2.46 was intentionally documentation/wording cleanup rather than feature work.
  - Historical docs were not rewritten wholesale; only current-state / plan docs were lightly updated.
  - The research sidecar remains preserved rather than modernized, with a warning that it is a frozen reference.
  - No generated sidecar state/build artifacts were intentionally touched or committed.
- Next rung:
  - v0.2.47 Cypher minimal opaque MLS artifact relay recon.
- Remaining future cleanup/polish:
  - optional Rust command-family splits (`errors.rs`, `identity.rs`, `public_bundle.rs`, `conversation.rs`, `message.rs`);
  - optional movement of remaining conversation/message printers out of `main.rs`;
  - future production storage/vault design;
  - future Comms runtime integration;
  - future Android/Pixel 4a testing much later.


**New at v0.2.47:**

- `carbonstack` head is now `40ccdfa docs: record Cypher OpenMLS envelope relay recon`.
- `carbonstack-comms` remains at `a371a96 docs: update OpenMLS sidecar README and stale warnings`.
- `carbonstack-cypher` remains at `0bfd5af chore: remove tracked Cypher local state artifacts`; no Cypher implementation code has landed yet.
- `docs/97-cypher-opaque-mls-artifact-relay-plan-v0.md` and `docs/98-cypher-opaque-mls-artifact-relay-recon-v0.md` landed under `b9f6fe0 docs: plan Cypher opaque MLS artifact relay`.
- `docs/99-cypher-openmls-envelope-relay-recon-v0.md` landed under `40ccdfa docs: record Cypher OpenMLS envelope relay recon`.
- v0.2.47 is a recon/design breakpoint, not an implementation checkpoint.
- Important recovered Cypher context:
  - CarbonStackCypher is already a Go + SQLite HTTP relay skeleton;
  - server startup is in `cmd/cypher/main.go`;
  - configuration is in `internal/config/config.go`;
  - SQLite/open/migration helpers are in `internal/db/db.go`;
  - HTTP routes and handlers are in `internal/httpapi/api.go`;
  - API tests are in `internal/httpapi/api_test.go`;
  - initial schema is in `migrations/001_init.sql`;
  - current routes include `GET /v0/health`, invite/device registration routes, `POST /v0/envelopes`, `GET /v0/devices/{device_id}/envelopes`, and `POST /v0/envelopes/{envelope_id}/ack`.
- Recon corrected the first implementation direction:
  - the original “opaque artifact relay” concept remains valid;
  - however, the repo already has an opaque envelope relay, so a parallel `/v0/artifacts` API is not the first safe move;
  - first implementation should reuse `/v0/envelopes` and widen accepted envelope content types/protocol compatibility.
- Existing `envelopes` storage already carries the right first-order routing fields:
  - `envelope_id`;
  - `sender_device_id`;
  - `recipient_device_id`;
  - `content_type`;
  - `protocol_version`;
  - `ciphertext_b64`;
  - `client_created_at`;
  - `server_received_at`;
  - `delivery_state`.
- For the first Cypher/OpenMLS link, `ciphertext_b64` is allowed to carry opaque artifact bytes as base64 despite the imperfect name.
- First accepted OpenMLS content types should be:
  - `carbonstack.mls.keypackage.v0`;
  - `carbonstack.mls.welcome.v0`;
  - `carbonstack.mls.application-message.v0`.
- Existing stub content type should remain accepted:
  - `carbonstack.message.text.stub.v0`.
- Recommended OpenMLS protocol version for the first relay scaffold:
  - `carbonstack-openmls-sidecar-v0`.
- Sidecar artifact mapping for Cypher relay:
  - `public-bundle.keypackage.bin` -> `carbonstack.mls.keypackage.v0`;
  - `welcome.bin` -> `carbonstack.mls.welcome.v0`;
  - `application-message.bin` -> `carbonstack.mls.application-message.v0`.
- Do not route or store:
  - `signer.json`;
  - `provider-storage.json`;
  - raw MemoryStorage JSON;
  - raw OpenMLS group state;
  - plaintext;
  - private keys;
  - trust-state private material.
- First implementation todo is now concrete:
  - add accepted content-type helper(s) in Cypher;
  - add protocol/content compatibility helper(s);
  - update `submitEnvelope` validation without changing route shape;
  - add tests for the three MLS content types;
  - add exact opaque byte roundtrip tests through submit/list;
  - keep existing stub envelope tests passing;
  - keep unknown content type, unsupported protocol version, invalid base64, and recipient-only ack rejection behavior.
- No database migration is required for the first implementation rung. Future migrations may add `payload_sha256` and `payload_size_bytes`, but v0.2.47 recon recommends test-side byte equality first.
- Comms/Cypher linking remains manual/test-scaffold next:
  - sidecar writes artifact bytes;
  - test/helper base64-encodes bytes;
  - Cypher stores via `/v0/envelopes`;
  - recipient lists/downloads envelope payload;
  - helper decodes bytes and writes a sidecar-compatible artifact file;
  - sidecar consumes it.
- Continuity/blunder note:
  - the initial recon framing over-weighted a new artifact API;
  - structural recon recovered that Cypher already has the correct mailbox primitive;
  - the user correctly flagged that this recovered context was checkpoint-worthy for linking Cypher + Comms, so v0.2.47 is preserved as a breakpoint before implementation.
- Next rung:
  - v0.2.48 implement Cypher OpenMLS envelope content-type widening and exact byte roundtrip tests.
- Later rung after Cypher accepts MLS artifact envelopes:
  - Comms-side helper/bridge proof that moves sidecar-produced `application-message.bin` through Cypher and back into `message-open`.


**New at v0.2.48:**

- `carbonstack` head is now `88a4e39 docs: record Cypher OpenMLS envelope content types`.
- `carbonstack-cypher` head is now `f491cea feat: accept OpenMLS envelope content types`.
- `carbonstack-comms` remains at `a371a96 docs: update OpenMLS sidecar README and stale warnings`.
- `docs/100-cypher-openmls-envelope-content-type-result-v0.md` records the first Cypher OpenMLS envelope implementation result.
- v0.2.48 is a narrow implementation checkpoint in `carbonstack-cypher`, not a Comms runtime integration checkpoint.
- Cypher now accepts OpenMLS sidecar artifact content types through the existing `/v0/envelopes` opaque envelope route:
  - `carbonstack.mls.keypackage.v0`;
  - `carbonstack.mls.welcome.v0`;
  - `carbonstack.mls.application-message.v0`.
- Existing stub envelope support remains preserved:
  - content type `carbonstack.message.text.stub.v0`;
  - protocol version `stub-v0`.
- OpenMLS artifact envelopes use the CarbonStack-specific protocol version:
  - `carbonstack-openmls-sidecar-v0`.
- The protocol-version name is deliberately CarbonStack-specific. It is not a claim of generic OpenMLS standard compatibility, production E2EE, or external security certification.
- The implementation reuses the existing mailbox flow:
  - `POST /v0/envelopes`;
  - `GET /v0/devices/{device_id}/envelopes`;
  - `POST /v0/envelopes/{envelope_id}/ack`.
- No new routes were added.
- No database migration was added.
- No `/v0/artifacts` API was added.
- No Comms runtime wiring was added.
- No MLS parsing was added inside Cypher.
- No plaintext handling was added inside Cypher.
- No signer/provider-storage relay was added.
- For this scaffold, existing `ciphertext_b64` is intentionally treated as “opaque envelope payload bytes encoded as base64.” This is semantically imperfect for KeyPackage/Welcome artifacts, but preserves the existing database/API shape and keeps the first relay proof small.
- Tests now prove exact opaque byte roundtrip for:
  - keypackage-like bytes;
  - welcome-like bytes;
  - application-message-like bytes.
- Tests also preserve/cover:
  - existing stub envelope lifecycle behavior;
  - unsupported content-type rejection;
  - unsupported protocol-version rejection;
  - OpenMLS content type with old `stub-v0` rejection;
  - stub content type with `carbonstack-openmls-sidecar-v0` rejection;
  - invalid base64 rejection.
- Important blunder/repair:
  - the first `api.go` patch only added constants/helpers but did not replace the old hardcoded `content_type == carbonstack.message.text.stub.v0` / `protocol_version == stub-v0` checks in `submitEnvelope`;
  - the first `api_test.go` helper replacement failed because it depended on exact formatting of `submitEnvelope`;
  - this was repaired with boundary/regex-based patches that replaced the hardcoded validation and added `submitEnvelopeWithContentType(...)` without depending on fragile whitespace;
  - this was an operator/patching issue, not an architectural change.
- Validation target:
  - `go test ./internal/httpapi -run "TestOpenMLSContentTypesRoundTripOpaqueBytes|TestOpenMLSContentTypeRejectsWrongProtocolVersion|TestStubContentTypeRejectsOpenMLSProtocolVersion"`;
  - `go test ./internal/httpapi`;
  - `go test ./...`;
  - CarbonStack docs validation via `scripts/validate-local.ps1`.
- Next rung:
  - v0.2.49 Comms-side relay bridge recon: map sidecar artifact path hints to Cypher envelope submit/list/decode/writeback flows before implementation.
- Later rung:
  - v0.2.50-ish end-to-end sidecar artifact relay proof: KeyPackage -> Welcome -> application-message through Cypher and back into sidecar consumers.

### Near-future server-deployable experimental backbone goalset

The project now has a near-future “server-deployable” goalset after the relay proof ladder, likely v0.2.51+.

This should be framed as:

- an **experimental server-deployable CarbonStack backbone**;
- CarbonStack minus Android app work;
- a deployable Cypher/OpenMLS-sidecar foundation that builders can inspect, run, and build on;
- immature/pre-release infrastructure, not certified secure infrastructure;
- not production E2EE;
- not externally audited;
- not suitable for strong security claims until experienced external senior developers/reviewers audit, harden, and fix the system.

The purpose is practical: a premade OpenMLS/Cypher relay backbone is much easier for others to deploy and experiment with than requiring every builder to sprint from zero. The documentation must stay honest about what is proven, what is only dev-scaffolded, and what remains unreviewed.


**New at v0.2.49:**

- `carbonstack` head is now `aef99d9 docs: record Comms OpenMLS Cypher relay bridge recon`.
- `carbonstack-cypher` remains at `f491cea feat: accept OpenMLS envelope content types`.
- `carbonstack-comms` remains at `a371a96 docs: update OpenMLS sidecar README and stale warnings`; no Comms implementation code has landed yet.
- `docs/101-comms-openmls-cypher-relay-bridge-recon-v0.md` records the Comms-side bridge recon and placement decision.
- v0.2.49 is a recon/design checkpoint, not a runtime integration checkpoint.
- Important recovered Comms context:
  - Comms already has `internal/client/cypher.go` with `CypherClient.SubmitEnvelope(...)`, `CypherClient.Inbox(...)`, and `CypherClient.AckEnvelope(...)`;
  - these client methods already carry `sender_device_id`, `recipient_device_id`, `content_type`, `protocol_version`, and `ciphertext_b64`;
  - `internal/app/commands.go` owns current stub-era CLI `send`, `inbox`, and `ack`;
  - `internal/protocol/openmls_sidecar_*_test.go` owns the promoted OpenMLS sidecar contract tests and helper facts;
  - docs already describe the historical Phase 1 Comms client flow of base64-encoding payloads, submitting envelopes to Cypher, retrieving queued envelopes, and acknowledging delivery.
- Recon placement decision:
  - first implementation should not go directly into `internal/app`, because that would prematurely wire OpenMLS artifact relay into user-facing `send` / `inbox`;
  - first implementation should not live only as sidecar test helper code, because the bridge is a real Comms/Cypher translation boundary that will later feed runtime integration;
  - recommended first implementation location is `carbonstack-comms/internal/relay`.
- Intended `internal/relay` responsibility:
  - map OpenMLS artifact kind to Cypher content type;
  - use `carbonstack-openmls-sidecar-v0`;
  - read sidecar artifact bytes from caller-provided safe path hints;
  - base64-encode artifact bytes;
  - submit through `client.CypherClient.SubmitEnvelope(...)`;
  - decode `client.EnvelopeRecord.CiphertextB64`;
  - write exact bytes to a caller-provided safe output path;
  - reject unsupported artifact kinds;
  - avoid raw byte logging.
- Suggested first relay constants:
  - `carbonstack.mls.keypackage.v0`;
  - `carbonstack.mls.welcome.v0`;
  - `carbonstack.mls.application-message.v0`;
  - `carbonstack-openmls-sidecar-v0`.
- Suggested first tests:
  - KeyPackage artifact kind maps to `carbonstack.mls.keypackage.v0`;
  - Welcome artifact kind maps to `carbonstack.mls.welcome.v0`;
  - application-message artifact kind maps to `carbonstack.mls.application-message.v0`;
  - arbitrary binary bytes roundtrip through relay helper encode/decode;
  - unsupported artifact kind is rejected;
  - output path behavior is explicit and safe;
  - no signer/provider-storage artifact kind exists.
- Explicitly deferred:
  - `comms send` integration;
  - `comms inbox` integration;
  - automatic sidecar invocation;
  - automatic ack after sidecar consume;
  - trust-state mutation;
  - local vault integration;
  - payload hash/size migration;
  - Android behavior;
  - production server-deployable packaging.
- Continuity/blunder note:
  - the user’s breakpoint request included a typo-like “v0.2.35 breakpoint handoff state” phrase, but the authoritative snapshot and attached files are v0.2.48 -> v0.2.49;
  - this checkpoint preserves v0.2.49 as the Comms bridge recon state.
- Next rung:
  - v0.2.50 Comms relay helper scaffold in `internal/relay`.
- Later rung:
  - v0.2.51 sidecar-produced application-message relay proof through Cypher and back into `message-open`.


**New at v0.2.50:**

- `carbonstack` head is now `ec52079 docs: record Comms OpenMLS artifact relay helper`.
- `carbonstack-comms` head is now `d06b2be feat: add OpenMLS artifact relay helpers`.
- `carbonstack-cypher` remains at `f491cea feat: accept OpenMLS envelope content types`.
- `docs/102-comms-openmls-artifact-relay-helper-result-v0.md` records the Comms relay helper scaffold result.
- v0.2.50 is a narrow Comms helper implementation checkpoint, not an end-to-end OpenMLS/Cypher runtime integration checkpoint.
- New Comms package:
  - `carbonstack-comms/internal/relay`.
- New helper files:
  - `internal/relay/openmls_artifacts.go`;
  - `internal/relay/openmls_artifacts_test.go`.
- The helper owns OpenMLS artifact kind constants:
  - `keypackage`;
  - `welcome`;
  - `application-message`.
- The helper maps artifact kinds to Cypher content types:
  - `keypackage` -> `carbonstack.mls.keypackage.v0`;
  - `welcome` -> `carbonstack.mls.welcome.v0`;
  - `application-message` -> `carbonstack.mls.application-message.v0`.
- The helper exposes the CarbonStack-specific OpenMLS relay protocol version:
  - `carbonstack-openmls-sidecar-v0`.
- The helper provides local/pure artifact payload functions:
  - read artifact bytes from disk;
  - write artifact bytes to disk;
  - base64 encode artifact bytes;
  - base64 decode envelope payload bytes;
  - read artifact -> base64 payload;
  - base64 payload -> artifact file.
- Tests validate:
  - KeyPackage content-type mapping;
  - Welcome content-type mapping;
  - application-message content-type mapping;
  - unsupported artifact kind rejection;
  - arbitrary binary artifact read -> base64 -> decode -> write roundtrip;
  - invalid base64 rejection;
  - directory read rejection.
- The helper intentionally does not yet:
  - invoke the OpenMLS sidecar;
  - call Cypher directly;
  - submit through `client.CypherClient.SubmitEnvelope`;
  - retrieve through `client.CypherClient.Inbox`;
  - acknowledge Cypher envelopes;
  - wire into `comms send` or `comms inbox`;
  - mutate trust-state;
  - parse MLS internals;
  - inspect plaintext;
  - route signer/provider storage.
- Safety boundary remains:
  - no helper artifact kind exists for `signer.json`;
  - no helper artifact kind exists for `provider-storage.json`;
  - raw MemoryStorage JSON, raw OpenMLS group state, plaintext, private keys, and trust-state private material remain local-only and non-relayable.
- Validation target:
  - `go test -p 1 ./internal/relay`;
  - `go test -p 1 ./...`;
  - `powershell -ExecutionPolicy Bypass -File .\scripts\check-no-rust-artifacts.ps1`;
  - CarbonStack docs validation via `scripts/validate-local.ps1`.
- Continuity/blunder note:
  - v0.2.50 stayed correctly narrow after the v0.2.49 placement decision;
  - the package is a helper scaffold only, avoiding premature CLI/runtime wiring;
  - no generated sidecar state, DB files, or artifact bytes should be committed as part of this rung.
- Next rung:
  - v0.2.51 sidecar-produced application-message relay proof through Cypher and back into `message-open`.
- Later rungs:
  - v0.2.52 KeyPackage + Welcome relay proof;
  - v0.2.53 full KeyPackage -> Welcome -> application-message relay lifecycle through Cypher;
  - v0.2.54+ deployable server + CLI smoke tests and docs;
  - v0.3.0 minor epoch for an experimental server-deployable CarbonStack backbone after the relay path is demonstrably testable.


**New at v0.2.51:**

- `carbonstack` head is now `3754844 docs: record Comms Cypher OpenMLS bridge helpers`.
- `carbonstack-comms` head is now `cb28f34 feat: add Cypher OpenMLS artifact bridge helpers`.
- `carbonstack-cypher` remains at `f491cea feat: accept OpenMLS envelope content types`.
- `docs/103-comms-cypher-openmls-application-artifact-bridge-helper-result-v0.md` records the Comms/Cypher/OpenMLS bridge helper result.
- v0.2.51 is still helper/test infrastructure, not polished runtime send/inbox integration and not the full sidecar-produced application-message relay proof.
- `carbonstack-comms/internal/relay` now includes bridge helpers beyond the v0.2.50 pure local artifact/base64 scaffold:
  - `SubmitOpenMLSArtifactEnvelope(...)`;
  - `WriteOpenMLSArtifactFromEnvelope(...)`.
- `SubmitOpenMLSArtifactEnvelope(...)`:
  - maps an OpenMLS artifact kind to the Cypher content type;
  - reads the local artifact file;
  - base64-encodes the artifact bytes;
  - uses `client.CypherClient.SubmitEnvelope(...)`;
  - sends `carbonstack-openmls-sidecar-v0`;
  - accepts caller-provided sender/recipient Cypher device IDs;
  - fills `client_created_at` if omitted.
- `WriteOpenMLSArtifactFromEnvelope(...)`:
  - decodes `client.EnvelopeRecord.CiphertextB64`;
  - writes the exact decoded bytes to a caller-provided local path;
  - creates the output parent directory as needed through existing helper behavior.
- Tests now validate:
  - writing exact decoded bytes from a Cypher envelope record to a local artifact file;
  - bridge request construction against a local `httptest` Cypher-like server;
  - correct `POST /v0/envelopes` method/path;
  - correct sender/recipient device IDs;
  - correct OpenMLS application-message content type;
  - correct `carbonstack-openmls-sidecar-v0` protocol version;
  - exact base64 payload matching the local artifact bytes;
  - unsupported artifact kinds such as `signer.json` are rejected before network submission.
- The helper still does not:
  - invoke the OpenMLS sidecar;
  - wire into `internal/app` send/inbox;
  - automatically ack Cypher envelopes;
  - mutate trust-state;
  - parse MLS internals;
  - inspect plaintext;
  - relay signer/provider storage.
- Important continuity/blunder note:
  - The first bridge-test draft attempted to read the HTTP request body through an invalid `r.Body.(*os.File)` assumption.
  - The test was replaced with a correct `io.ReadAll(r.Body)` + JSON decode implementation before validation and commit.
  - This was a test-drafting/operator issue, not an architectural change.
- Validation target:
  - `go test -p 1 ./internal/relay`;
  - `go test -p 1 ./...`;
  - `powershell -ExecutionPolicy Bypass -File .\scripts\check-no-rust-artifacts.ps1`;
  - CarbonStack docs validation via `scripts/validate-local.ps1`.
- Next rung:
  - v0.2.52 sidecar-produced application-message relay proof through Cypher and back into `message-open`.
- Later rungs:
  - KeyPackage + Welcome relay proof;
  - full KeyPackage -> Welcome -> application-message relay lifecycle;
  - deployable server + CLI smoke tests/docs;
  - pre-v0.3.0 release cleanup breakpoint;
  - v0.3.0 experimental server-deployable backbone epoch.



**New at v0.2.52:**

- `carbonstack` head is now `22fd7de docs: record OpenMLS application message relay proof`.
- `carbonstack-comms` head is now `caa6683 test: prove OpenMLS application message relay through Cypher envelope`.
- `carbonstack-cypher` remains at `f491cea feat: accept OpenMLS envelope content types`.
- `docs/104-openmls-application-message-relay-through-cypher-result-v0.md` records the first application-message relay proof result.
- v0.2.52 is the first dev/test proof that a **sidecar-produced** OpenMLS application message artifact can cross the Comms/Cypher relay boundary and still be opened by the recipient sidecar.
- The proof validates the path:
  - promoted OpenMLS sidecar sets up Alice/Bob two-member conversation;
  - Alice-side `message-protect` writes `application-message.bin`;
  - `internal/relay.SubmitOpenMLSArtifactEnvelope(...)` reads the artifact bytes, base64-encodes them, and submits a Cypher-compatible `/v0/envelopes` request;
  - a protocol-local Cypher-compatible `httptest` envelope server stores the queued envelope;
  - recipient Comms client retrieves the queued envelope via `Inbox(...)`;
  - `internal/relay.WriteOpenMLSArtifactFromEnvelope(...)` decodes `ciphertext_b64` and writes a downloaded `application-message.bin`;
  - Bob-side `message-open` consumes the downloaded artifact;
  - `assertMessageOpenSuccess(...)` confirms `plaintext_utf8` matches the original plaintext.
- New/active Comms test files for the proof:
  - `internal/protocol/openmls_sidecar_relay_test.go`;
  - `internal/protocol/openmls_sidecar_test_cypher_server_test.go`.
- `internal/relay/test_cypher_server_test.go` also exists as a relay-package test-server helper from the first draft/scaffold.
- The proof uses the existing content type and protocol contract:
  - `carbonstack.mls.application-message.v0`;
  - `carbonstack-openmls-sidecar-v0`.
- The proof intentionally does not:
  - wire `comms send`;
  - wire `comms inbox`;
  - automatically ack after sidecar consume;
  - parse MLS internals in Cypher;
  - parse MLS internals in `internal/relay`;
  - mutate trust-state;
  - relay signer/provider storage;
  - add Cypher routes;
  - add Cypher migrations;
  - claim production readiness.
- Important blunders/repair:
  - The first relay proof test draft assumed `setupOpenMLSTwoMemberConversation(...)` returned a custom setup struct with Alice/Bob labels, but the real helper returns an `openMLSSidecarEnvelope`.
  - The first test draft also called `protectOpenMLSSidecarMessage(...)` and `openOpenMLSSidecarMessage(...)` with non-existent expanded signatures.
  - The repair used the existing fixed helper contract: `setupOpenMLSTwoMemberConversation(t)`, `protectOpenMLSSidecarMessage(t, messageLabel, plaintext)`, and `openOpenMLSSidecarMessage(t, messageLabel, path)`.
  - The first plaintext assertion looked for a non-existent `Data.Plaintext` field.
  - The repair reused the existing `assertMessageOpenSuccess(...)` helper and the real `PlaintextUTF8`/`PlaintextLen` schema instead of inventing a new assertion path.
  - A transient Windows `go: unlinkat ... relay.test.exe` warning appeared after an `ok` relay test, likely from local AV/process handle timing; it was not the semantic test failure.
- Validation target:
  - `go test -p 1 ./internal/relay`;
  - `go test -p 1 ./internal/protocol -run "TestOpenMLSSidecarApplicationMessageRelayThroughCypherEnvelope"`;
  - `go test -p 1 ./internal/protocol`;
  - `go test -p 1 ./...`;
  - `powershell -ExecutionPolicy Bypass -File .\scripts\check-no-rust-artifacts.ps1`;
  - CarbonStack docs validation via `scripts/validate-local.ps1`.
- Next rung:
  - v0.2.53 KeyPackage + Welcome relay proof.
- Later rungs:
  - full KeyPackage -> Welcome -> application-message relay lifecycle;
  - deployable server + CLI smoke tests/docs;
  - pre-v0.3.0 release cleanup/modularity/stale-claims sweep;
  - v0.3.0 experimental server-deployable backbone epoch.


**New at v0.2.53:**

- `carbonstack` head is now `964a390 docs: record OpenMLS keypackage welcome relay proof`.
- `carbonstack-comms` head is now `6ace21d test: prove OpenMLS keypackage welcome relay through Cypher envelope`.
- `carbonstack-cypher` remains at `f491cea feat: accept OpenMLS envelope content types`.
- `docs/105-openmls-keypackage-welcome-relay-through-cypher-result-v0.md` records the KeyPackage + Welcome relay proof result.
- v0.2.53 proves the onboarding half of the Cypher/OpenMLS relay path as a dev/test proof.
- The proof validates the KeyPackage path:
  - Bob creates sidecar identity state;
  - Bob runs `public-bundle-export --write-artifact`;
  - sidecar writes `public-bundle.keypackage.bin`;
  - `internal/relay.SubmitOpenMLSArtifactEnvelope(...)` submits the KeyPackage artifact as `carbonstack.mls.keypackage.v0`;
  - the protocol-local Cypher-compatible `httptest` server stores the queued envelope for Alice;
  - Alice retrieves the KeyPackage envelope through `client.CypherClient.Inbox(...)`;
  - `internal/relay.WriteOpenMLSArtifactFromEnvelope(...)` writes a downloaded sidecar-compatible KeyPackage artifact;
  - Alice `conversation-add-member --member-keypackage <downloaded-keypackage>` consumes the relayed artifact.
- The proof validates the Welcome path:
  - Alice `conversation-add-member` writes `welcome.bin`;
  - `internal/relay.SubmitOpenMLSArtifactEnvelope(...)` submits the Welcome artifact as `carbonstack.mls.welcome.v0`;
  - the protocol-local Cypher-compatible server stores the queued envelope for Bob;
  - Bob retrieves the Welcome envelope through `client.CypherClient.Inbox(...)`;
  - `internal/relay.WriteOpenMLSArtifactFromEnvelope(...)` writes a downloaded sidecar-compatible Welcome artifact;
  - Bob `conversation-join --welcome <downloaded-welcome>` consumes the relayed artifact;
  - Bob joined group is reloadable and has `member_count = 2`.
- `internal/protocol/openmls_sidecar_test_cypher_server_test.go` was widened so the protocol-local test Cypher server accepts all current OpenMLS relay content types:
  - `carbonstack.mls.keypackage.v0`;
  - `carbonstack.mls.welcome.v0`;
  - `carbonstack.mls.application-message.v0`.
- The protocol-local test server now uses unique `test-envelope-000N` envelope IDs, making multi-envelope relay tests easier to inspect.
- The new proof intentionally does not use `setupOpenMLSTwoMemberConversation(t)`, because that helper consumes local KeyPackage/Welcome artifacts directly and would bypass the relay boundary being tested.
- Important blunders/repair:
  - The first content-type widening patch partially applied: it added `fmt` but did not replace the application-message-only content-type check or the hardcoded `test-envelope-0001` ID.
  - That caused a Go compile failure due to an unused `fmt` import.
  - The repair used direct regex replacement to widen the content-type switch and to generate unique envelope IDs from `len(tc.envelopes)+1`.
  - This was a patching/PowerShell exact-string issue, not a protocol or sidecar behavior issue.
- Validation target:
  - `go test -p 1 ./internal/protocol -run "TestOpenMLSSidecarKeyPackageWelcomeRelayThroughCypherEnvelope"`;
  - `go test -p 1 ./internal/protocol -run "TestOpenMLSSidecarApplicationMessageRelayThroughCypherEnvelope|TestOpenMLSSidecarKeyPackageWelcomeRelayThroughCypherEnvelope"`;
  - `go test -p 1 ./internal/relay`;
  - `go test -p 1 ./internal/protocol`;
  - `go test -p 1 ./...`;
  - `powershell -ExecutionPolicy Bypass -File .\scripts\check-no-rust-artifacts.ps1`;
  - CarbonStack docs validation via `scripts/validate-local.ps1`.
- Next rung:
  - v0.2.54 full KeyPackage -> Welcome -> application-message relay lifecycle proof.
- Later rungs:
  - deployable server + CLI smoke tests/docs;
  - pre-v0.3.0 release cleanup/modularity/stale-claims sweep;
  - v0.3.0 experimental server-deployable backbone epoch.


---


**New at v0.2.54:**

- `carbonstack` head is now `eb9956e docs: record full OpenMLS relay lifecycle proof`.
- `carbonstack-comms` head is now `e4cbf9b test: prove full OpenMLS relay lifecycle through Cypher envelope`.
- `carbonstack-cypher` remains at `f491cea feat: accept OpenMLS envelope content types`.
- `docs/106-full-openmls-relay-lifecycle-through-cypher-result-v0.md` records the full stitched lifecycle proof result.
- v0.2.54 proves the complete current OpenMLS artifact relay lifecycle in one dev/test path:
  - Bob creates identity state and exports `public-bundle.keypackage.bin`;
  - `internal/relay.SubmitOpenMLSArtifactEnvelope(...)` submits the KeyPackage artifact as `carbonstack.mls.keypackage.v0`;
  - Alice retrieves the KeyPackage envelope through `client.CypherClient.Inbox(...)`;
  - `internal/relay.WriteOpenMLSArtifactFromEnvelope(...)` writes the downloaded KeyPackage artifact;
  - Alice `conversation-add-member --member-keypackage <downloaded-keypackage>` consumes it and writes `welcome.bin`;
  - `internal/relay.SubmitOpenMLSArtifactEnvelope(...)` submits the Welcome artifact as `carbonstack.mls.welcome.v0`;
  - Bob retrieves the Welcome envelope through `client.CypherClient.Inbox(...)`;
  - `internal/relay.WriteOpenMLSArtifactFromEnvelope(...)` writes the downloaded Welcome artifact;
  - Bob `conversation-join --welcome <downloaded-welcome>` consumes it and joins the group;
  - Alice `message-protect` writes `application-message.bin`;
  - `internal/relay.SubmitOpenMLSArtifactEnvelope(...)` submits the application message artifact as `carbonstack.mls.application-message.v0`;
  - Bob retrieves the application-message envelope, writes a sidecar-compatible artifact, and `message-open` recovers the expected plaintext.
- The proof stitches the separately validated v0.2.52 application-message relay and v0.2.53 KeyPackage/Welcome relay into a single current lifecycle proof.
- The proof intentionally remains a protocol-level dev/test proof using a Cypher-compatible `httptest` envelope server, not a deployed Cypher server or polished Comms runtime UX.
- The Bob inbox in the full lifecycle test sees both the earlier Welcome envelope and later application-message envelope because automatic ack remains deliberately deferred; the test selects the application-message envelope by content type.
- No Cypher route, migration, or storage change was needed for v0.2.54.
- No sidecar command, JSON envelope schema, path-hint semantics, or OpenMLS lifecycle behavior was intentionally changed.
- Preserved boundaries:
  - no `comms send` / `comms inbox` OpenMLS UX wiring;
  - no automatic ack after sidecar consume;
  - no trust-state mutation;
  - no MLS parsing in Cypher or `internal/relay`;
  - no signer/provider-storage relay;
  - no production E2EE, metadata privacy, hostile-server proof, Android readiness, external audit, or certified secure claim;
  - no deployable server packaging yet.
- Validation target:
  - `go test -p 1 ./internal/protocol -run "TestOpenMLSSidecarFullLifecycleRelayThroughCypherEnvelope"`;
  - `go test -p 1 ./internal/protocol -run "TestOpenMLSSidecarApplicationMessageRelayThroughCypherEnvelope|TestOpenMLSSidecarKeyPackageWelcomeRelayThroughCypherEnvelope|TestOpenMLSSidecarFullLifecycleRelayThroughCypherEnvelope"`;
  - `go test -p 1 ./internal/relay`;
  - `go test -p 1 ./internal/protocol`;
  - `go test -p 1 ./...`;
  - `powershell -ExecutionPolicy Bypass -File .\scripts\check-no-rust-artifacts.ps1`;
  - CarbonStack docs validation via `scripts/validate-local.ps1`.
- Next rung:
  - v0.2.55 deployable server + CLI smoke-test recon/planning.
- Later rungs:
  - server/CLI deployability smoke tests and docs;
  - pre-v0.3.0 cleanup/modularity/stale-claims/release-hardening checkpoint;
  - v0.3.0 experimental server-deployable backbone epoch.


**New at v0.2.55:**

- `carbonstack` head is now `33d0502 docs: plan deployable server CLI smoke tests`.
- `carbonstack-comms` remains at `e4cbf9b test: prove full OpenMLS relay lifecycle through Cypher envelope`.
- `carbonstack-cypher` remains at `f491cea feat: accept OpenMLS envelope content types`.
- `carbonstack-os` remains at `b537475 Add CarbonStackOS north star and initial appliance model`.
- `docs/107-deployable-server-cli-smoke-test-recon-v0.md` records the deployable server + CLI smoke-test recon/planning result.
- v0.2.55 is a docs/recon breakpoint, not an implementation checkpoint.
- Recon confirmed the real `carbonstack-cypher` server startup path:
  - entry point: `carbonstack-cypher/cmd/cypher/main.go`;
  - config loader: `carbonstack-cypher/internal/config/config.go`;
  - DB/migration helper: `carbonstack-cypher/internal/db/db.go`;
  - HTTP routes: `carbonstack-cypher/internal/httpapi/api.go`;
  - server creates a `db.Store`, runs migrations, seeds the dev invite when configured, creates `httpapi.New(...)`, and listens with `http.ListenAndServe`.
- Recon confirmed the real Cypher server is env-configured:
  - `CYPHER_ADDR` default `:8080`;
  - `CYPHER_DB` default `cypher.db`;
  - `CYPHER_MIGRATIONS` default `migrations`;
  - `CYPHER_DEV_INVITE` default `dev-invite`.
- Recommended v0.2.56 smoke-test overrides:
  - `CYPHER_ADDR=127.0.0.1:<test-port>`;
  - `CYPHER_DB=<temp-dir>/cypher-smoke.db`;
  - `CYPHER_MIGRATIONS=<absolute-cypher-repo>/migrations`;
  - `CYPHER_DEV_INVITE=dev-invite`.
- Recon confirmed the real Cypher route family already matches the relay proof needs:
  - `GET /v0/health`;
  - `POST /v0/dev/invites`;
  - `POST /v0/invites/claim`;
  - `POST /v0/devices/register`;
  - `GET /v0/accounts/{account_id}/devices`;
  - `POST /v0/envelopes`;
  - `GET /v0/devices/{device_id}/envelopes`;
  - `POST /v0/envelopes/{envelope_id}/ack`.
- Recon confirmed `carbonstack-comms/internal/client/cypher.go` already exposes the required client primitives:
  - `client.New(serverURL)`;
  - `CreateDevInvite(...)`;
  - `ClaimInvite(...)`;
  - `RegisterDevice(...)`;
  - `ListDevices(...)`;
  - `SubmitEnvelope(...)`;
  - `Inbox(...)`;
  - `AckEnvelope(...)`.
- Recon confirmed `carbonstack-comms/internal/app/commands.go` is still stub-era CLI surface:
  - `send`, `inbox`, and `ack` use the existing mock/stub text envelope path;
  - OpenMLS relay must not be wired into user-facing CLI until the test/harness path is proven.
- Placement decision for v0.2.56:
  - keep the real-server OpenMLS lifecycle smoke proof under `carbonstack-comms/internal/protocol`;
  - reason: the proof still depends heavily on sidecar lifecycle helpers;
  - keep `internal/relay` generic and already-proven;
  - avoid `internal/app` until a deliberate runtime UX boundary exists.
- v0.2.56 target:
  - start a real `carbonstack-cypher` process on localhost;
  - use a temp SQLite DB and real migrations;
  - wait for `GET /v0/health`;
  - use `internal/client.CypherClient` against the real server;
  - claim/register Alice and Bob Cypher devices;
  - run the full OpenMLS sidecar artifact lifecycle through real `/v0/envelopes`;
  - kill/cleanup the server process and avoid generated DB/state artifacts.
- Agreed deployability ladder:
  - v0.2.56: real Cypher server + Comms client/relay smoke proof;
  - v0.2.57: CLI/dev harness for repeatable local relay lifecycle;
  - v0.2.58: ack semantics after successful sidecar consume;
  - v0.2.59: payload metadata/hash/size planning or migration;
  - v0.2.60: deployability docs/runbook + known-good validation;
  - v0.2.61: Option C ironed out and completed for testing;
  - v0.2.62: stale helper cleanup, inbox/ack/general semantics/schema standardization;
  - v0.2.63+: Option B planning + implementation;
  - v0.2.x: pre-v0.3.0 cleanup/release-hardening checkpoint;
  - v0.3.0: experimental server-deployable CarbonStack backbone epoch.
- Boundary preserved:
  - no deployed server proof has landed yet;
  - no CLI/dev harness has landed yet;
  - no automatic ack has landed yet;
  - no payload hash/size migration has landed yet;
  - no release packaging has landed yet;
  - no production E2EE, hostile-server proof, metadata privacy, Android readiness, external audit, or certified secure claim is made.
- Next rung:
  - v0.2.56 real Cypher server + Comms client/relay smoke proof.


**New at v0.2.56:**

- `carbonstack` head is now `2bd48f6 docs: record real Cypher OpenMLS relay proof`.
- `carbonstack-comms` head is now `73f29ce test: prove OpenMLS relay lifecycle through real Cypher server`.
- `carbonstack-cypher` remains at `f491cea feat: accept OpenMLS envelope content types`.
- `carbonstack-os` remains at `b537475 Add CarbonStackOS north star and initial appliance model`.
- `docs/108-real-cypher-server-openmls-relay-lifecycle-result-v0.md` records the real-server OpenMLS relay proof result.
- v0.2.56 is the first implementation checkpoint proving the full OpenMLS sidecar relay lifecycle against a real locally-started `carbonstack-cypher` process rather than the protocol-local Cypher-compatible `httptest` server used in v0.2.54.
- The new real-server proof lives in `carbonstack-comms/internal/protocol`, preserving the v0.2.55 placement decision because the flow still depends heavily on OpenMLS sidecar lifecycle helpers.
- New/active real-server smoke-test files:
  - `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_server_test.go`;
  - `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go`.
- The proof starts a real Cypher server using a test-built Cypher binary rather than `go run ./cmd/cypher`, because the first `go run` approach left orphaned Windows `cypher.exe` child processes and held the temp SQLite DB open during cleanup.
- The test server helper:
  - resolves the sibling `carbonstack-cypher` repo from the `carbonstack-comms` test path;
  - builds `./cmd/cypher` to a temp test binary;
  - allocates a localhost test port;
  - starts the Cypher binary with explicit env vars;
  - uses a temp SQLite DB;
  - points `CYPHER_MIGRATIONS` at the real migrations directory;
  - waits for `GET /v0/health`;
  - kills/cleans up the process after the test.
- Real-server smoke-test env shape:
  - `CYPHER_ADDR=127.0.0.1:<test-port>`;
  - `CYPHER_DB=<temp-dir>/cypher-real-server-smoke.db`;
  - `CYPHER_MIGRATIONS=<carbonstack-cypher>/migrations`;
  - `CYPHER_DEV_INVITE=dev-invite`.
- The proof uses `internal/client.CypherClient` against the real server for:
  - `ClaimInvite(...)`;
  - `CreateDevInvite(...)`;
  - `RegisterDevice(...)`;
  - `SubmitEnvelope(...)`;
  - `Inbox(...)`.
- The proof validates the full real-server lifecycle:
  - Alice and Bob claim/register Cypher devices against the real server;
  - Bob sidecar `public-bundle-export --write-artifact` writes `public-bundle.keypackage.bin`;
  - `internal/relay.SubmitOpenMLSArtifactEnvelope(...)` submits the KeyPackage artifact to real `/v0/envelopes` with `carbonstack.mls.keypackage.v0`;
  - Alice retrieves the KeyPackage from real `/v0/devices/{device_id}/envelopes`;
  - Alice sidecar `conversation-add-member` consumes the downloaded KeyPackage and writes `welcome.bin`;
  - `internal/relay.SubmitOpenMLSArtifactEnvelope(...)` submits the Welcome artifact to real `/v0/envelopes` with `carbonstack.mls.welcome.v0`;
  - Bob retrieves the Welcome from the real inbox and consumes it with `conversation-join`;
  - Alice sidecar `message-protect` writes `application-message.bin`;
  - `internal/relay.SubmitOpenMLSArtifactEnvelope(...)` submits the application-message artifact to real `/v0/envelopes` with `carbonstack.mls.application-message.v0`;
  - Bob retrieves the application-message envelope from the real inbox, writes a sidecar-compatible artifact, and `message-open` recovers matching plaintext.
- Important blunders/repair:
  - The first real-server helper used `go run ./cmd/cypher` with stdout/stderr attached through `strings.Builder` pipes. On Windows, the `go run` parent and compiled `cypher.exe` child did not cleanly terminate, causing long test timeouts inside cleanup / `cmd.Wait()`.
  - A second run confirmed orphaned `cypher` processes remained and held the temp SQLite DB open, causing `TempDir RemoveAll cleanup` to fail with “file being used by another process.”
  - The repair killed stale local `cypher` processes, changed the helper to build a temp Cypher binary first, run that binary directly, log to a temp file, and use process-tree cleanup on Windows.
  - This was a Windows process-management/test-harness issue, not a Cypher API, Comms relay, or OpenMLS lifecycle failure.
- Validation target:
  - `go test -p 1 ./internal/protocol -run "TestOpenMLSSidecarFullLifecycleRelayThroughRealCypherServer" -count=1 -timeout 240s`;
  - `go test -p 1 ./internal/protocol -run "TestOpenMLSSidecarFullLifecycleRelayThroughCypherEnvelope|TestOpenMLSSidecarFullLifecycleRelayThroughRealCypherServer" -count=1 -timeout 300s`;
  - `go test -p 1 ./internal/relay`;
  - `go test -p 1 ./internal/protocol -count=1 -timeout 360s`;
  - `go test -p 1 ./... -count=1 -timeout 360s`;
  - `powershell -ExecutionPolicy Bypass -File .\scripts\check-no-rust-artifacts.ps1`;
  - CarbonStack docs validation via `scripts/validate-local.ps1`.
- Preserved boundaries:
  - no polished `comms send` / `comms inbox` OpenMLS UX wiring;
  - no automatic ack after sidecar consume;
  - no payload hash/size migration;
  - no trust-state mutation;
  - no MLS parsing in Cypher or `internal/relay`;
  - no signer/provider-storage relay;
  - no release package;
  - no production E2EE, hostile-server proof, metadata privacy, Android readiness, external audit, or certified secure claim.
- Next rung:
  - v0.2.57 CLI/dev harness for repeatable local relay lifecycle.
- Later rungs:
  - v0.2.58 ack semantics after successful sidecar consume;
  - v0.2.59 payload metadata/hash/size planning or migration;
  - v0.2.60 deployability docs/runbook + known-good validation;
  - v0.2.61 Option C ironed out and completed for testing;
  - v0.2.62 stale helper cleanup, inbox/ack/general semantics/schema standardization;
  - v0.2.63+ Option B planning + implementation;
  - pre-v0.3.0 cleanup/release-hardening checkpoint;
  - v0.3.0 experimental server-deployable CarbonStack backbone epoch.



**New at v0.2.57:**

- `carbonstack` head is now `ead63f1 docs: record OpenMLS real Cypher relay smoke harness`.
- `carbonstack-comms` head is now `c91800d Cypher relay smoke test script addition`, after `bfadceb test: add OpenMLS real Cypher relay smoke harness`.
- `carbonstack-cypher` remains at `f491cea feat: accept OpenMLS envelope content types`.
- `carbonstack-os` remains at `b537475 Add CarbonStackOS north star and initial appliance model`.
- `docs/109-openmls-real-cypher-relay-smoke-harness-result-v0.md` records the v0.2.57 dev harness result.
- v0.2.57 turns the v0.2.56 real-server proof into a repeatable developer-facing smoke harness.
- New Comms harness path:
  - `carbonstack-comms/scripts/smoke-openmls-real-cypher-relay.ps1`.
- The harness runs the known-good real-server proof:
  - `TestOpenMLSSidecarFullLifecycleRelayThroughRealCypherServer`.
- The harness validates the real local Cypher server flow:
  - builds/runs the real `carbonstack-cypher` test binary through the Go test path;
  - starts a local real Cypher server;
  - uses a temp SQLite DB;
  - waits for `/v0/health`;
  - uses Comms `internal/client.CypherClient`;
  - uses Comms `internal/relay` helpers;
  - runs the OpenMLS KeyPackage -> Welcome -> application-message relay lifecycle;
  - verifies final sidecar `message-open` plaintext recovery.
- The harness runs the generated Rust/OpenMLS artifact guard:
  - `powershell -ExecutionPolicy Bypass -File .\scripts\check-no-rust-artifacts.ps1`.
- The harness supports a broader `-Full` mode that runs additional protocol/relay validation.
- The harness is deliberately developer-facing and experimental:
  - not polished Comms runtime UX;
  - not production E2EE;
  - not certified secure;
  - not externally audited;
  - not Android-ready;
  - not a release package.
- The harness refuses to proceed when existing `cypher` processes are detected before a run, to avoid stale process/SQLite lock confusion after the Windows orphan-process issue seen during v0.2.56.
- README/current-facing Comms notes were updated to document the harness and warn that normal `comms send` / `comms inbox` remain stub-era.
- Validation target:
  - `powershell -ExecutionPolicy Bypass -File .\scripts\smoke-openmls-real-cypher-relay.ps1`;
  - `powershell -ExecutionPolicy Bypass -File .\scripts\smoke-openmls-real-cypher-relay.ps1 -Full`;
  - CarbonStack docs validation via `scripts/validate-local.ps1`.
- Important continuity/blunder notes:
  - the final snapshot shows two Comms commits for this rung: `bfadceb` for the original harness addition and `c91800d` for a follow-up smoke-test script addition/update;
  - treat `c91800d` as the authoritative Comms head for v0.2.57;
  - no `carbonstack-cypher` implementation changes were required for v0.2.57.
- Next rung:
  - v0.2.58 ack semantics after successful sidecar consume.
- v0.2.58 should prove the first safe ack boundary:
  - write downloaded artifact;
  - sidecar successfully consumes it;
  - only then acknowledge the Cypher envelope;
  - keep this scoped to tests/harness before user-facing runtime UX.


**New at v0.2.58:**

- `carbonstack` head is now `d94c322 docs: record OpenMLS relay ack semantics`.
- `carbonstack-comms` head is now `e978e49 test: ack OpenMLS relay envelopes after sidecar consume`.
- `carbonstack-cypher` remains at `f491cea feat: accept OpenMLS envelope content types`.
- `carbonstack-os` remains at `b537475 Add CarbonStackOS north star and initial CarbonStack repository structure`.
- `docs/110-openmls-relay-ack-semantics-result-v0.md` records the v0.2.58 ack semantics result.
- v0.2.58 proves the first safe consume-then-ack boundary for relayed OpenMLS artifacts against the real local Cypher server proof path.
- The rule now validated is:
  - do not ack when an envelope is merely downloaded;
  - do not ack when artifact bytes are merely written to disk;
  - only ack after the recipient sidecar successfully consumes the artifact.
- The real-server lifecycle test now calls `CypherClient.AckEnvelope(...)` after successful sidecar consume for each relay artifact.
- Validated ack boundaries:
  - KeyPackage: Alice acks the KeyPackage envelope only after `conversation-add-member` successfully consumes the downloaded KeyPackage artifact and produces the Welcome artifact.
  - Welcome: Bob acks the Welcome envelope only after `conversation-join` successfully consumes the downloaded Welcome artifact and the joined group is reloadable.
  - application-message: Bob acks the application-message envelope only after `message-open` successfully consumes the downloaded `application-message.bin` and plaintext recovery is validated.
- After each ack, the test checks the relevant recipient inbox is empty:
  - Alice inbox empty after KeyPackage ack;
  - Bob inbox empty after Welcome ack;
  - Bob inbox empty after application-message ack.
- Because Welcome is now acked before application-message relay, Bob's later message inbox expectation changes from two queued envelopes to one queued application-message envelope.
- `scripts/smoke-openmls-real-cypher-relay.ps1` was also hardened so native command failures stop the harness correctly.
- The harness now wraps `go test` / artifact-guard native commands through `Invoke-NativeCommand` and throws on nonzero `$LASTEXITCODE`.
- Important continuity/blunder notes:
  - The first docs/110 commit landed in `carbonstack` before the corresponding Comms ack implementation landed, creating a temporary docs-ahead-of-implementation mismatch.
  - The first PowerShell `.Replace(...)` patch for `openmls_sidecar_real_cypher_relay_test.go` was a no-op because exact text/indentation did not match the `gofmt`ed file. `Select-String` showed no `AckEnvelope`, `git diff` was empty, and `git commit` reported nothing to commit.
  - A regex-based patch then inserted the ack calls, but the first compile failed because `CypherClient.AckEnvelope(...)` returns two values. The repair used `_, err := cypherClient.AckEnvelope(...)`.
  - The same failed validation exposed a harness bug: PowerShell `$ErrorActionPreference = "Stop"` does not automatically fail the script on native command nonzero exit codes. The script was repaired to check `$LASTEXITCODE`.
  - These were patching/harness issues, not a Cypher route, sidecar lifecycle, or relay architecture failure.
- Validation target:
  - `go test -p 1 ./internal/protocol -run "TestOpenMLSSidecarFullLifecycleRelayThroughRealCypherServer" -count=1 -timeout 240s`;
  - `powershell -ExecutionPolicy Bypass -File .\scripts\smoke-openmls-real-cypher-relay.ps1`;
  - `powershell -ExecutionPolicy Bypass -File .\scripts\smoke-openmls-real-cypher-relay.ps1 -Full`;
  - `powershell -ExecutionPolicy Bypass -File .\scripts\check-no-rust-artifacts.ps1`;
  - CarbonStack docs validation via `scripts/validate-local.ps1`.
- Preserved boundaries:
  - no polished `comms send` / `comms inbox` OpenMLS UX wiring;
  - no payload hash/size migration yet;
  - no trust-state mutation;
  - no MLS parsing in Cypher or `internal/relay`;
  - no signer/provider-storage relay;
  - no release package;
  - no production E2EE, hostile-server proof, metadata privacy, Android readiness, external audit, or certified secure claim.
- Next rung:
  - v0.2.59 payload metadata/hash/size planning or migration.
- v0.2.59 should decide whether to keep using `ciphertext_b64` as the only payload field for the experimental backbone or add payload metadata such as `payload_sha256` and `payload_size_bytes` through a scoped Cypher migration/API/test update.


**New at v0.2.59:**

- `carbonstack` head is now `5e371ee docs: plan OpenMLS relay payload metadata`.
- `carbonstack-comms` remains at `e978e49 test: ack OpenMLS relay envelopes after sidecar consume`.
- `carbonstack-cypher` remains at `f491cea feat: accept OpenMLS envelope content types`.
- `carbonstack-os` remains at `b537475 Add CarbonStackOS north star and initial CarbonStack repository structure`.
- `docs/111-openmls-relay-payload-metadata-plan-v0.md` records the v0.2.59 payload metadata/hash/size plan.
- v0.2.59 is intentionally a planning/design checkpoint, not a schema/API migration implementation checkpoint.
- Recon confirmed the current Cypher `envelopes` table has no payload metadata fields; it stores `ciphertext_b64` with routing/metadata fields such as `content_type`, `protocol_version`, timestamps, and `delivery_state`.
- Recon confirmed `submitEnvelope` currently validates base64 but discards decoded bytes after validation, then stores the original `ciphertext_b64` string.
- Recon confirmed inbox responses currently return queued envelopes with `ciphertext_b64` but no decoded payload hash/size metadata.
- Recon confirmed Comms `EnvelopeRecord` and `SubmitEnvelopeResponse` currently mirror this shape and do not yet expose `PayloadSHA256` or `PayloadSizeBytes`.
- Planned metadata fields:
  - `payload_size_bytes`: integer count of decoded payload bytes from `ciphertext_b64`;
  - `payload_sha256`: lowercase hex SHA-256 digest of decoded payload bytes from `ciphertext_b64`.
- Planned server-side rule:
  - compute both fields in Cypher during `submitEnvelope` after successful base64 decode;
  - store them with the envelope;
  - return them in submit response and inbox response;
  - do not trust client-provided metadata in the first implementation.
- Planned DB migration:
  - add a new migration such as `002_envelope_payload_metadata.sql`;
  - add nullable `payload_sha256 TEXT` and `payload_size_bytes INTEGER` to `envelopes`;
  - populate new rows from server-computed decoded payload metadata;
  - tolerate old dev rows as null unless a later backfill is deliberately added.
- Planned Comms behavior:
  - add payload metadata fields to client response structs;
  - validate metadata in relay helper before writing decoded artifact bytes;
  - reject mismatched size/hash before sidecar consume;
  - keep sidecar consume (`conversation-add-member`, `conversation-join`, `message-open`) as the cryptographic semantic gate.
- Security boundary:
  - payload metadata is relay/debug/storage sanity metadata;
  - it is not OpenMLS authenticity;
  - it is not hostile-server proof;
  - it is not production E2EE or a trust root.
- Important continuity note:
  - v0.2.59 deliberately chose planning first because the change crosses Cypher schema/API/tests, Comms client structs, Comms relay validation, real-server smoke proof expectations, and documentation claim wording.
  - The next implementation rung should update all affected seams together rather than silently changing one repo.
- Suggested next rung:
  - payload metadata implementation across `carbonstack-cypher` and `carbonstack-comms`, possibly shifting the previous v0.2.60 deployability-docs rung later.
- Preserved boundaries:
  - no DB migration landed at v0.2.59;
  - no API response shape changed at v0.2.59;
  - no Comms client struct changed at v0.2.59;
  - no relay helper metadata validation landed at v0.2.59;
  - no polished `comms send` / `comms inbox` OpenMLS UX wiring;
  - no release package;
  - no production E2EE, hostile-server proof, metadata privacy, Android readiness, external audit, or certified secure claim.



**New at v0.2.60:**

- `carbonstack` head is now `30da19c docs: record OpenMLS relay payload metadata`.
- `carbonstack-comms` head is now `d606a19 test: validate OpenMLS relay payload metadata`.
- `carbonstack-cypher` head is now `ab5e99f feat: add envelope payload metadata`.
- `carbonstack-os` remains at `b537475 Add CarbonStackOS north star and initial CarbonStack repository structure`.
- `docs/112-openmls-relay-payload-metadata-result-v0.md` records the v0.2.60 payload metadata implementation result.
- v0.2.60 implements the v0.2.59 plan rather than merely documenting it.
- Cypher now stores and returns metadata over decoded `ciphertext_b64` payload bytes:
  - `payload_size_bytes`: decoded payload byte length;
  - `payload_sha256`: lowercase hex SHA-256 digest over decoded payload bytes.
- New Cypher migration:
  - `carbonstack-cypher/migrations/002_envelope_payload_metadata.sql`;
  - adds nullable `payload_sha256 TEXT`;
  - adds nullable `payload_size_bytes INTEGER`.
- Cypher `submitEnvelope` now:
  - decodes `ciphertext_b64`;
  - rejects invalid base64 as before;
  - computes `payloadSizeBytes = len(decodedPayload)`;
  - computes `payloadSHA256 = sha256(decodedPayload)` encoded as lowercase hex;
  - stores both metadata fields with the envelope;
  - returns both metadata fields in the submit response.
- Cypher inbox response now returns payload metadata for queued envelopes.
- `carbonstack-cypher/internal/httpapi/api.go` and `api_test.go` were updated to cover the new API/storage shape.
- `carbonstack-comms/internal/client/cypher.go` now carries:
  - `PayloadSHA256 string`;
  - `PayloadSizeBytes int64`;
  on `SubmitEnvelopeResponse` and `EnvelopeRecord`.
- `carbonstack-comms/internal/relay/cypher_bridge.go` now validates metadata before writing downloaded sidecar artifact bytes:
  - base64 decode first;
  - reject unsupported OpenMLS artifact content type/protocol;
  - compare decoded length against `PayloadSizeBytes` when nonzero;
  - compare `sha256(decoded bytes)` against `PayloadSHA256` when present;
  - fail before writing artifact on mismatch;
  - create the parent output directory before `os.WriteFile(..., 0o600)`.
- `carbonstack-comms/internal/relay/cypher_bridge_test.go` now covers matching metadata and mismatch rejection:
  - valid application-message artifact write with metadata;
  - payload-size mismatch rejection;
  - payload-SHA256 mismatch rejection.
- The real-server OpenMLS relay proof now asserts metadata is present and correct for:
  - KeyPackage envelopes;
  - Welcome envelopes;
  - application-message envelopes.
- Ack semantics from v0.2.58 remain preserved:
  - KeyPackage ack only after `conversation-add-member`;
  - Welcome ack only after `conversation-join`;
  - application-message ack only after `message-open` and plaintext validation.
- Payload metadata remains a relay/debug/storage sanity feature:
  - it is not OpenMLS authenticity;
  - it is not hostile-server proof;
  - it is not production E2EE;
  - it is not a trust root;
  - OpenMLS sidecar consume remains the cryptographic semantic gate.
- Important blunders/repair during v0.2.60:
  - The first Cypher import patch added `crypto/sha256` and `encoding/hex` before the submit body actually used them, causing unused-import failures.
  - The first Cypher body patch produced an 11-column `INSERT` but initially still passed the old 9-value set; this was repaired by replacing the full `Exec(...)` block.
  - Cypher inbox `SELECT` was updated to return 11 columns, but `rows.Scan(...)` initially still scanned 9 destinations; the response struct and scan list were repaired together.
  - The first Comms client struct patch no-op'd because exact text matching missed the actual formatting; forced struct replacement was used.
  - The first Comms bridge rewrite referenced nonexistent `ArtifactKind` / `ArtifactKindForContentType` names; the bridge was corrected to preserve the repo's current `artifactKind string` and content-type helper style.
  - The stricter `WriteOpenMLSArtifactFromEnvelope(...)` initially broke the legacy test fixture because the fixture lacked `ContentType` and `ProtocolVersion`; the test was updated to include OpenMLS application-message metadata.
  - Direct `os.WriteFile(...)` initially regressed parent-directory creation behavior that had been provided by the older helper; `os.MkdirAll(filepath.Dir(outputPath), 0o700)` was added before writing.
  - Windows intermittently reported `unlinkat ... relay.test.exe ... being used by another process`; this was treated as local Windows/AV/file-lock noise after failed test runs, not a project semantic failure.
- Validation target at v0.2.60:
  - `carbonstack-cypher`: `go test ./internal/httpapi -count=1`;
  - `carbonstack-cypher`: `go test ./... -count=1`;
  - `carbonstack-comms`: `go test -p 1 ./internal/relay -count=1`;
  - `carbonstack-comms`: real-server OpenMLS relay proof;
  - `carbonstack-comms`: `scripts/smoke-openmls-real-cypher-relay.ps1`;
  - `carbonstack-comms`: `scripts/smoke-openmls-real-cypher-relay.ps1 -Full`;
  - `carbonstack-comms`: `scripts/check-no-rust-artifacts.ps1`;
  - `carbonstack`: `scripts/validate-local.ps1`.
- Suggested next rung:
  - deployability docs/runbook + known-good validation, probably under the main `carbonstack` repo as the front-door release/doctrine surface.
- v0.3.0 release-framing continuity:
  - preferred framing is a hybrid: **CarbonStack experimental backbone** at the top level, with the concrete validated artifact being the Cypher + Comms OpenMLS relay backbone;
  - the main `carbonstack` README/release surface should start with a nontechnical “what am I looking at?” head, then continue into a technical body/table of contents;
  - component repos (`carbonstack-comms`, `carbonstack-cypher`) remain focused on implementation updates, development, testing, and component-specific docs.



**New at v0.2.61:**

- `carbonstack` head is now `995924f Docs: Added readme for the docs folder`.
- `carbonstack` also landed `bc45be0 docs: add experimental backbone runbook`.
- `carbonstack-comms` remains at `d606a19 test: validate OpenMLS relay payload metadata`.
- `carbonstack-cypher` remains at `ab5e99f feat: add envelope payload metadata`.
- `carbonstack-os` remains at `b537475 Add CarbonStackOS north star and initial CarbonStack repository structure`.
- `docs/113-experimental-backbone-deployability-runbook-v0.md` records the first deployability runbook for the current experimental backbone.
- `docs/README.md` now explicitly frames `carbonstack/docs` as a historical docs archive and continuity trail, not a single polished always-current manual.
- The docs archive rule is now explicit:
  - numbered docs are chronological;
  - lower-numbered docs may be stale;
  - older docs preserve planning/recon/mistakes/pivots/context;
  - current behavior should be read from newest relevant release docs, current-state docs, READMEs, and runbooks;
  - release-specific docs should become the source of truth for a given release surface.
- v0.2.61 preserves the historical docs model instead of rewriting old process docs or pretending historical docs will not drift.
- The deployability runbook describes the current validated artifact in non-release-hype language:
  - an experimental CarbonStack backbone;
  - not a finished messenger;
  - not a production security product;
  - not Android-ready;
  - not externally audited or certified secure;
  - the concrete validated artifact is the Cypher + Comms OpenMLS relay backbone.
- The runbook records the current known-good local validation path:
  - `carbonstack-cypher`: `go test ./... -count=1`;
  - `carbonstack-comms`: stop stale `cypher` processes, run `scripts/smoke-openmls-real-cypher-relay.ps1`, run `-Full`, and run `scripts/check-no-rust-artifacts.ps1`;
  - `carbonstack`: `scripts/validate-local.ps1`.
- The runbook documents the current relay content types:
  - `carbonstack.mls.keypackage.v0`;
  - `carbonstack.mls.welcome.v0`;
  - `carbonstack.mls.application-message.v0`;
  - protocol version `carbonstack-openmls-sidecar-v0`.
- The runbook records current payload metadata behavior:
  - `payload_size_bytes`;
  - `payload_sha256`;
  - both describe decoded `ciphertext_b64` payload bytes;
  - Comms validates the metadata before writing sidecar artifact bytes;
  - this remains relay/debug/storage sanity metadata, not a cryptographic trust root.
- The runbook records current ack semantics:
  - do not ack on download;
  - do not ack on artifact write;
  - ack only after sidecar consume succeeds.
- The runbook records generated-state cleanup warnings and the Windows stale `cypher.exe` process check.
- No component implementation code changed in v0.2.61.
- No Comms runtime `send` / `inbox` integration was added.
- No production/security/Android/external-audit claim was added.
- Important continuity note:
  - the user explicitly wants runbook docs first;
  - then full-spectrum README/public-surface recon;
  - then critical README/quick-start/info-doc cleanup where stale or oddly framed;
  - future README edits should use the user's long writing samples and formal-style meta-analysis before rewriting public-facing language.
- Next rung:
  - full-spectrum README and public-surface recon across `carbonstack`, `carbonstack-comms`, `carbonstack-cypher`, and relevant quick-start/info docs;
  - only then apply critical wording/stale-claim cleanup to READMEs and first-contact docs.




**New at v0.2.62:**

- `carbonstack` head is now `020725e docs: clean public CarbonStack surface`.
- `carbonstack-comms` head is now `0318e11 docs: clean public Comms surface`.
- `carbonstack-cypher` head is now `baac52d docs: clean public Cypher surface`.
- `carbonstack-os` remains at `b537475 Add CarbonStackOS north star and initial CarbonStack repository structure`.
- v0.2.62 is a public-facing documentation cleanup checkpoint, not a runtime/protocol implementation checkpoint.
- The cleanup used the user's formal technical writing style extrapolated from the attached academic-adjacent writing samples:
  - define the system first;
  - define the failure/limit;
  - state the proof boundary;
  - state non-goals early;
  - avoid hype and soft marketing language;
  - avoid GPT/internal-assistant lexical framing;
  - address the external reader directly;
  - preserve direct, system-level, first-principles explanation.
- `carbonstack/README.md` is now the clear public front door:
  - starts with "CarbonStack is an experimental secure-communications backbone";
  - states that it is not a finished messenger, not production-certified, not externally audited, and not Android-ready;
  - names the current validated artifact as the local Cypher + Comms OpenMLS relay proof;
  - lists what is currently proven;
  - lists what is not proven;
  - maps the component repositories;
  - points to `docs/113-experimental-backbone-deployability-runbook-v0.md`;
  - keeps the v0.3.0 direction framed as an experimental backbone epoch.
- `carbonstack/docs/113-experimental-backbone-deployability-runbook-v0.md` was tightened into a current known-good runbook:
  - clearer "what am I looking at" entry;
  - stronger component repo roles;
  - direct local checkout assumptions;
  - known-good validation commands;
  - current relay contract;
  - payload metadata boundary;
  - consume-then-ack semantics;
  - generated-state and stale `cypher.exe` process warnings;
  - explicit proof and non-proof lists.
- `carbonstack/roadmap/ROADMAP.md` was updated to match the current ladder:
  - v0.2.61 deployability runbook complete;
  - v0.2.62 public surface cleanup current/complete;
  - v0.2.63 Option C completion for testing;
  - v0.2.64 inbox/ack/schema semantics cleanup;
  - v0.2.65+ Option B planning/implementation;
  - pre-v0.3.0 release-hardening checkpoint;
  - v0.3.0 experimental backbone epoch.
- `carbonstack/docs/14-carbonstack-full-specification.md` now starts with a current-status notice:
  - it is a historical canonical specification draft;
  - it preserves doctrine and long-range intent;
  - it should not be treated as the current release surface by itself;
  - current behavior starts from `README.md`, `docs/README.md`, and `docs/113`.
- `carbonstack-cypher/README.md` was rewritten as a component front door:
  - Cypher is the experimental relay/storage server;
  - it stores opaque envelopes;
  - it does not handle plaintext;
  - it does not decide trust;
  - it is not production-certified or externally audited;
  - it lists current implemented routes, content types, protocol versions, migrations, payload metadata, and validation commands.
- `carbonstack-cypher/docs/05-mvp-roadmap.md` now reflects current implemented relay/server scaffold state instead of older vague roadmap framing.
- `carbonstack-cypher/docs/07-data-model-v0.md` now describes the implemented development schema:
  - invite/account/device/envelope tables;
  - `ciphertext_b64`;
  - `payload_sha256`;
  - `payload_size_bytes`;
  - delivery state;
  - security boundary.
- `carbonstack-cypher/docs/08-api-contract-v0.md` now describes the implemented development API:
  - health;
  - invite claim;
  - device registration;
  - device lookup;
  - envelope submit;
  - inbox list;
  - ack;
  - submit/inbox payload metadata;
  - accepted OpenMLS content types and protocol version.
- `carbonstack-comms/README.md` was rewritten as a component front door:
  - Comms is the text-first communications client component;
  - it is not a finished messenger;
  - it explains the current OpenMLS sidecar + Cypher relay proof;
  - it lists the current validated relay path;
  - it gives smoke harness commands;
  - it points back to the main `carbonstack` runbook.
- `carbonstack-comms/internal/protocol/mls/README.md` now correctly states that the MLS area is no longer only a future placeholder:
  - active promoted sidecar lives at `internal/protocol/mls/openmls-sidecar`;
  - research remains historical;
  - current use is experimental OpenMLS sidecar artifact relay proofing;
  - nonclaims are explicit.
- `carbonstack-comms/internal/protocol/mls/openmls-sidecar/README.md` was updated from stale v0.2.47-era framing:
  - no longer says the sidecar is not wired into Cypher;
  - records its use in current Comms protocol tests and real-Cypher relay lifecycle proofs;
  - lists supported commands;
  - preserves generated-state warnings;
  - preserves non-production/non-audited/non-runtime-UX boundaries.
- `carbonstack-comms/docs/02-client-protocol-foundation.md` now clearly distinguishes earlier Signal-style Phase 1 exploration from the current OpenMLS sidecar relay direction.
- `carbonstack-comms/docs/04-phase1-client-lifecycle.md` now has a current-state notice:
  - old Phase 1 CLI lifecycle is historical scaffolding;
  - current proof is OpenMLS sidecar + Cypher relay, not polished runtime CLI messenger.
- `carbonstack-comms/docs/05-local-state-model-v0.md` now warns that the early development local-state model is not a secure production vault and that current generated sidecar state is sensitive dev material.
- `carbonstack-comms/docs/requirements.md` now separates intended requirements from current implementation and avoids implying that production local vault/security requirements are already met.
- `carbonstack-comms/docs/local-vault.md` now frames local vault as a future design requirement, not an implemented production storage model.
- Historical numbered CarbonStack archive docs were not rewritten wholesale.
- The cleanup intentionally avoided converting component repos into the main release narrative; clean release framing remains centered in `carbonstack`.
- No Cypher server behavior changed.
- No Comms runtime behavior changed.
- No OpenMLS sidecar command behavior changed.
- No new production/security/Android/external-audit claim was added.
- Important blunders/continuity notes:
  - the first public-surface identification pass correctly separated first-contact docs from historical archive docs;
  - the cleanup script rewrote multiple public-surface docs across repos rather than editing the numbered history archive;
  - `docs/14` was not fully rewritten, only guarded with a current-status notice to prevent stale canonical-spec misuse;
  - this rung deliberately favors public clarity and nonclaim hardening over new implementation.
- Next rung:
  - v0.2.63 Option C completion / known-good testing cleanup;
  - keep the current local harness path repeatable;
  - verify public README/runbook commands still match the current repo state;
  - then proceed toward inbox/ack/schema standardization and pre-v0.3.0 release hardening.


## 2. Explicit Non-Goals / Out of Scope

At v0.2.62, do not:

- treat payload metadata as a security proof or OpenMLS authenticity layer;
- claim payload metadata proves hostile-server safety, metadata privacy, production E2EE, or external certification;

- wire the `internal/relay` helper into `internal/app` send/inbox before the full lifecycle proof lands;
- treat `internal/relay` as a full Comms/Cypher runtime integration layer yet;
- add helper support for `signer.json`, `provider-storage.json`, raw MemoryStorage JSON, raw OpenMLS group state, plaintext, private keys, or trust-state private material;

- claim production E2EE, Signal-equivalent security, hostile-server proof, replay resistance, metadata privacy, or production-grade secure storage;
- wire OpenMLS into `comms send`, `comms inbox`, Cypher routing, `trust.json`, or `trust-events.jsonl`;
- claim Comms CLI uses this sidecar lifecycle;
- claim polished Comms runtime routes MLS application payloads through Cypher;
- mutate trust-state storage from sidecar events;
- treat the full KeyPackage -> Welcome -> application-message `httptest` lifecycle proof as a deployed server proof;
- claim production identity/device mapping is solved;
- promote relay helpers into user-facing CLI without a scoped runtime integration plan;
- implement autogenerated message IDs before explicit message-label proof;
- implement out-of-order buffering or replay tracking beyond what OpenMLS already enforces;
- treat dev-only `signer.json` as a secure vault;
- treat dev-only `provider-storage.json` as a secure vault or safe application data;
- print, paste, inspect casually, expose, or commit `signer.json`, `provider-storage.json`, provider storage JSON, MemoryStorage JSON, private keys, seeds, recovery material, raw key bytes, raw group state, raw Welcome bytes, raw KeyPackage bytes, or raw protected application-message bytes;
- treat `public-bundle.keypackage.bin`, `welcome.bin`, or `application-message.bin` as final CarbonStack UX formats;
- silently replace exported public artifacts, group state, Welcome artifacts, joined state, message artifacts, or conversation state;
- start Android or CarbonStackOS integration.
- treat sidecar promotion as productionization;
- delete or churn the known-good research sidecar during promoted module/test cleanup;
- split Go tests in the same commit as Rust module changes unless explicitly scoped and validated;
- begin Cypher routing before v0.2.45-v0.2.46 cleanup has either landed or been explicitly deferred;
- delete or churn the research sidecar during Rust module split;
- change command names, JSON envelopes, path-hint semantics, event names, or exit codes during v0.2.45 Go test splitting;
- commit `.carbonstack-openmls-sidecar-state/`, `target/`, `signer.json`, `provider-storage.json`, `welcome.bin`, `application-message.bin`, or other generated OpenMLS sidecar artifacts;
- add a parallel `/v0/artifacts` API before extending/evaluating the existing `/v0/envelopes` relay;
- parse MLS internals in Cypher;
- store OpenMLS signer/provider storage in Cypher;
- rename or migrate Cypher's `ciphertext_b64` field in the first relay scaffold unless explicitly scoped later;
- claim Cypher/OpenMLS integration is complete before exact byte roundtrip tests land;
- implement `internal/app` send/inbox OpenMLS runtime wiring before the `internal/relay` helper proof lands;
- place the first Comms bridge implementation directly in user-facing CLI commands;
- automatically acknowledge Cypher envelopes before sidecar consumption succeeds;
- treat the Comms relay bridge recon as polished UX or production integration;


---

## 3. Current State

| Classification | Items |
|---|---|
| **VALIDATED** | Phase 1 relay/client skeleton works locally. Phase 2A trust lifecycle remains test-protected. Phase 2C OpenMLS scratch/fixture/trust mapping path is closed and validated. Phase 2D sidecar provider-info/envelope/identity-create/identity-status/public-bundle-export summary/artifact/conversation-create/persistence/load-check/add-member/join/protect/open paths are Go-tested. |
| **VALIDATED PROVIDER-INFO CLEANUP** | `provider-info` no longer depends on fragile hand-authored raw JSON. `CAPABILITIES` and `UNSUPPORTED_COMMANDS` feed a `serde_json::json!` envelope. This prevents missing-comma command-list failures before further command churn. |
| **VALIDATED SIDECAR IDENTITY CREATION** | `identity-create --device-label <safe>` validates label, refuses overwrite, creates ignored per-device dev state, creates `SignatureKeyPair`, creates `BasicCredential`, creates `CredentialWithKey`, writes secret-bearing `signer.json`, writes sanitized identity JSON files, returns sanitized success JSON, emits `provider.identity.created`, and keeps `private_material_included=false`. |
| **VALIDATED SIDECAR IDENTITY STATUS** | `identity-status --device-label <safe>` validates label, requires existing identity state, loads/deserializes `signer.json`, reads sanitized summary/state files, derives public signer key, recomputes `sha256:<hex>` public identity reference, compares it to `identity-summary.json`, returns sanitized success JSON, emits `provider.identity.loaded`, and keeps `private_material_included=false`. |
| **VALIDATED PUBLIC BUNDLE EXPORT** | `public-bundle-export` generates a real OpenMLS KeyPackage, writes sanitized `public-bundle-summary.json`, and in artifact mode writes `public-bundle.keypackage.bin` and `public-bundle-manifest.json`. Since v0.2.34, it also saves device-root `provider-storage.json` so the private KeyPackage bundle is available for `StagedWelcome::new_from_welcome`. |
| **VALIDATED CONVERSATION-CREATE** | `conversation-create --device-label <safe> --conversation-label <safe>` validates device/conversation labels, requires identity state, creates a dev-local one-member OpenMLS group/conversation, writes sanitized `conversation-summary.json`, writes dev-local `provider-storage.json`, reports `member_count=1`, `epoch=GroupEpoch(0)`, `group_id_ref=sha256:<hex>`, refuses duplicates, emits `provider.conversation.created` / `provider.conversation.exists`, and keeps `private_material_included=false`. |
| **VALIDATED PROVIDER/GROUP PERSISTENCE** | `CarbonStackSidecarProvider` owns `RustCrypto + MemoryStorage`, implements `OpenMlsProvider`, saves `MemoryStorage` to `provider-storage.json`, reloads it into a fresh provider, and proves `MlsGroup::load` succeeds. |
| **VALIDATED CONVERSATION-LOAD-CHECK** | `conversation-load-check --device-label <safe> --conversation-label <safe>` validates labels, requires identity/conversation/provider-storage state, loads `provider-storage.json` into a fresh provider, calls `MlsGroup::load`, returns `provider_storage_loaded=true`, `group_reloadable=true`, member count, epoch, emits `conversation.loaded`, and keeps `private_material_included=false`. |
| **VALIDATED ADD-MEMBER / WELCOME EXPORT** | `conversation-add-member` consumes Bob's serialized public KeyPackage artifact, validates it into a `KeyPackage`, adds Bob to Alice's persisted group with `add_members`, exports the second returned `MlsMessageOut` as `welcome.bin`, merges the pending commit, saves mutated provider storage, writes sanitized Welcome/add-member metadata, and post-add load-check reports `member_count=2`. |
| **VALIDATED CONVERSATION-JOIN** | `conversation-join` consumes `welcome.bin`, deserializes it as `MlsMessageIn`, extracts `MlsMessageBodyIn::Welcome(welcome)`, loads Bob's device provider storage, calls `StagedWelcome::new_from_welcome`, calls `into_group`, saves Bob device-scoped joined provider storage, writes sanitized `conversation-summary.json` and `join-summary.json`, and proves reloadability. |
| **VALIDATED MESSAGE-PROTECT** | `message-protect` loads sender identity/signer and sender device-scoped conversation provider storage, loads Alice group, calls `MlsGroup::create_message`, serializes returned `MlsMessageOut` with `to_bytes()`, writes `application-message.bin`, saves sender provider storage, reload-proves Alice group, and writes sanitized `message-manifest.json` and `message-protect-summary.json`. |
| **VALIDATED MESSAGE-OPEN** | `message-open` loads receiver device-scoped conversation provider storage, loads Bob group, reads `application-message.bin`, deserializes as `MlsMessageIn`, converts to `ProtocolMessage`, calls `MlsGroup::process_message`, extracts `ProcessedMessageContent::ApplicationMessage(...).into_bytes()`, returns bounded dev `plaintext_utf8`, saves Bob provider storage, reload-proves Bob group, and writes sanitized `message-open-summary.json`. |
| **RECON COMPLETE FOR NEXT RUNG** | `docs/76` and `docs/77` define/confirm the next multi-message continuity implementation path: explicit `--message-label`, default `message-0001`, two sequential messages, and continued save-after-protect/open. |
| **PARTIAL** | Sidecar state remains dev-local. `public-bundle.keypackage.bin`, `welcome.bin`, and `application-message.bin` are real protocol artifacts but not final CarbonStack UX formats. Alice and Bob conversation state are now both device-scoped in the dev sidecar; the old global Alice path is retired. |
| **NOT VALIDATED** | Comms CLI integration, Cypher MLS routing, trust-state mutation, production vault, hostile-server harness, long skipped-message behavior, multi-sender ordering beyond Alice/Bob, epoch/membership-change ordering, metadata privacy, Android, CarbonStackOS, CI. |
| **NEXT** | v0.2.50 Comms relay helper scaffold in `internal/relay`: artifact kind/content-type mapping, artifact bytes -> Cypher envelope payload, Cypher envelope payload -> safe local artifact file, tests only before user-facing runtime wiring. |

---

## 4. Critical Paths

### Local repo paths

- `C:\▓▓\repos\carbonstack_umbrella`
- `C:\▓▓\repos\carbonstack_umbrella\carbonstack`
- `C:\▓▓\repos\carbonstack_umbrella\carbonstack-comms`
- `C:\▓▓\repos\carbonstack_umbrella\carbonstack-cypher`
- `C:\▓▓\repos\carbonstack_umbrella\carbonstack-os`

### Current repo heads at v0.2.50

- `carbonstack`: `ec52079 docs: record Comms OpenMLS artifact relay helper`
- `carbonstack-comms`: `d06b2be feat: add OpenMLS artifact relay helpers`
- `carbonstack-cypher`: `f491cea feat: accept OpenMLS envelope content types`
- `carbonstack-os`: `b537475 Add CarbonStackOS north star and initial appliance model`

### Snapshot cleanliness note

The v0.2.50 snapshot is clean across the four repo heads listed above. The v0.2.40 duplicate-adjacent commit-message note is historical only; the heads listed above are the authority.

### Canonical docs and runner

- `carbonstack/docs/00-project-charter.md`
- `carbonstack/docs/14-carbonstack-full-specification.md`
- `carbonstack/docs/15-protocol-foundation.md`
- `carbonstack/docs/36-provider-event-taxonomy-v0.md`
- `carbonstack/docs/38-provider-trust-state-mapping-v0.md`
- `carbonstack/docs/39-phase2d-sidecar-command-surface-plan.md`
- `carbonstack/docs/49-openmls-sidecar-real-identity-create-plan-v0.md`
- `carbonstack/docs/50-openmls-sidecar-real-identity-create-result-v0.md`
- `carbonstack/docs/51-openmls-sidecar-identity-status-plan-v0.md`
- `carbonstack/docs/52-openmls-sidecar-identity-status-result-v0.md`
- `carbonstack/docs/53-openmls-sidecar-public-bundle-export-plan-v0.md`
- `carbonstack/docs/54-openmls-sidecar-public-bundle-export-result-v0.md`
- `carbonstack/docs/55-openmls-sidecar-keypackage-artifact-export-plan-v0.md`
- `carbonstack/docs/56-openmls-sidecar-keypackage-artifact-export-result-v0.md`
- `carbonstack/docs/57-openmls-sidecar-conversation-lifecycle-plan-v0.md`
- `carbonstack/docs/58-openmls-sidecar-conversation-lifecycle-api-recon-v0.md`
- `carbonstack/docs/59-openmls-sidecar-conversation-create-plan-v0.md`
- `carbonstack/docs/60-openmls-sidecar-conversation-create-result-v0.md`
- `carbonstack/docs/61-openmls-sidecar-conversation-add-member-welcome-plan-v0.md`
- `carbonstack/docs/62-openmls-sidecar-add-member-welcome-api-recon-v0.md`
- `carbonstack/docs/63-openmls-sidecar-conversation-create-persistence-repair-plan-v0.md`
- `carbonstack/docs/64-openmls-sidecar-conversation-create-persistence-repair-result-v0.md`
- `carbonstack/docs/65-openmls-sidecar-dev-provider-group-persistence-plan-v0.md`
- `carbonstack/docs/66-openmls-sidecar-dev-provider-group-persistence-result-v0.md`
- `carbonstack/docs/67-openmls-sidecar-add-member-welcome-skeleton-v0.md`
- `carbonstack/docs/68-openmls-sidecar-add-member-welcome-api-recon-v0.md`
- `carbonstack/docs/69-openmls-sidecar-add-member-welcome-export-result-v0.md`
- `carbonstack/docs/70-openmls-sidecar-conversation-join-skeleton-v0.md`
- `carbonstack/docs/71-openmls-sidecar-conversation-join-api-recon-v0.md`
- `carbonstack/docs/72-openmls-sidecar-conversation-join-result-v0.md`
- `carbonstack/docs/73-openmls-sidecar-message-protect-open-skeleton-v0.md`
- `carbonstack/docs/74-openmls-sidecar-message-protect-open-api-recon-v0.md`
- `carbonstack/docs/75-openmls-sidecar-message-protect-open-result-v0.md`
- `carbonstack/docs/76-openmls-sidecar-multi-message-continuity-plan-v0.md`
- `carbonstack/docs/77-openmls-sidecar-multi-message-api-recon-v0.md`
- `carbonstack/docs/78-openmls-sidecar-multi-message-continuity-result-v0.md`
- `carbonstack/docs/79-openmls-sidecar-message-ordering-replay-plan-v0.md`
- `carbonstack/docs/80-openmls-sidecar-message-ordering-replay-api-recon-v0.md`
- `carbonstack/docs/82-openmls-sidecar-alice-device-scoped-state-layout-plan-v0.md`
- `carbonstack/docs/83-openmls-sidecar-alice-device-scoped-state-recon-v0.md`
- `carbonstack/docs/84-openmls-sidecar-alice-device-scoped-state-result-v0.md`
- `carbonstack/docs/85-openmls-sidecar-phase2d-closure-checklist-v0.md`
- `carbonstack/docs/86-openmls-sidecar-phase2d-mainline-closure-result-v0.md`
- `carbonstack/docs/87-openmls-sidecar-current-state-index-v0.md`
- `carbonstack/docs/88-openmls-sidecar-maintainability-promotion-plan-v0.md`
- `carbonstack/docs/89-openmls-sidecar-module-split-plan-v0.md`
- `carbonstack/docs/90-openmls-sidecar-test-suite-split-plan-v0.md`
- `carbonstack/docs/91-openmls-sidecar-artifact-ownership-map-v0.md`
- `carbonstack/docs/92-openmls-sidecar-command-schema-matrix-v0.md`
- `carbonstack/docs/93-openmls-sidecar-promotion-scaffold-result-v0.md`
- `carbonstack/docs/94-openmls-sidecar-rust-module-split-result-v0.md`
- `carbonstack/docs/95-openmls-sidecar-go-test-suite-split-result-v0.md`
- `carbonstack/docs/96-openmls-sidecar-readme-current-state-cleanup-result-v0.md`
- `carbonstack/docs/97-cypher-opaque-mls-artifact-relay-plan-v0.md`
- `carbonstack/docs/98-cypher-opaque-mls-artifact-relay-recon-v0.md`
- `carbonstack/docs/99-cypher-openmls-envelope-relay-recon-v0.md`
- `carbonstack/docs/100-cypher-openmls-envelope-content-type-result-v0.md
- `carbonstack/docs/101-comms-openmls-cypher-relay-bridge-recon-v0.md``
- `carbonstack-cypher/internal/httpapi/api.go`
- `carbonstack-cypher/internal/httpapi/api_test.go`
- `carbonstack-cypher/migrations/001_init.sql`
- Cypher OpenMLS content types:
  - `carbonstack.mls.keypackage.v0`
  - `carbonstack.mls.welcome.v0`
  - `carbonstack.mls.application-message.v0`
- Cypher OpenMLS protocol version: `carbonstack-openmls-sidecar-v0`

- `carbonstack/scripts/validate-local.ps1`

### CarbonStackComms implementation/test paths

- `carbonstack-comms/internal/protocol/provider_events.go`
- `carbonstack-comms/internal/protocol/provider_events_test.go`
- `carbonstack-comms/internal/protocol/provider_trust.go`
- `carbonstack-comms/internal/protocol/provider_trust_test.go`
- `carbonstack-comms/internal/protocol/openmls_sidecar_provider_info_test.go`
- `carbonstack-comms/internal/protocol/mls/research/openmls-sidecar/Cargo.toml`
- `carbonstack-comms/internal/protocol/mls/research/openmls-sidecar/Cargo.lock`
- `carbonstack-comms/internal/protocol/mls/research/openmls-sidecar/src/main.rs`
- `carbonstack-comms/internal/protocol/mls/research/openmls-sidecar/src/provider.rs`
- `carbonstack-comms/internal/protocol/mls/research/openmls-sidecar/src/labels.rs`
- `carbonstack-comms/internal/protocol/mls/research/openmls-sidecar/src/state.rs`
- promoted sidecar path: `carbonstack-comms/internal/protocol/mls/openmls-sidecar`
- promoted sidecar Rust modules after v0.2.44: `main.rs`, `cli.rs`, `envelope.rs`, `labels.rs`, `paths.rs`, `provider.rs`, `schema.rs`, `state.rs`; future optional splits remain `errors.rs`, `identity.rs`, `public_bundle.rs`, `conversation.rs`, `message.rs`
- split Go sidecar tests after v0.2.45:
  - `openmls_sidecar_helpers_test.go`
  - `openmls_sidecar_provider_info_test.go`
  - `openmls_sidecar_identity_test.go`
  - `openmls_sidecar_public_bundle_test.go`
  - `openmls_sidecar_conversation_test.go`
  - `openmls_sidecar_message_test.go`
  - `openmls_sidecar_message_negative_test.go`
- `carbonstack-comms/internal/protocol/mls/openmls-sidecar/README.md` — current promoted sidecar README after v0.2.46.
- `carbonstack-comms/internal/protocol/mls/research/openmls-sidecar/README.md` — frozen research-reference README after v0.2.46 notice.
- `carbonstack-comms/scripts/check-no-rust-artifacts.ps1`



### CarbonStackCypher implementation/test paths recovered at v0.2.47

- `carbonstack-cypher/cmd/cypher/main.go`
- `carbonstack-cypher/internal/config/config.go`
- `carbonstack-cypher/internal/db/db.go`
- `carbonstack-cypher/internal/httpapi/api.go`
- `carbonstack-cypher/internal/httpapi/api_test.go`
- `carbonstack-cypher/migrations/001_init.sql`
- `carbonstack-cypher/docs/07-data-model-v0.md`
- `carbonstack-cypher/docs/08-api-contract-v0.md`

Existing Cypher routes relevant to OpenMLS relay:

- `POST /v0/envelopes`
- `GET /v0/devices/{device_id}/envelopes`
- `POST /v0/envelopes/{envelope_id}/ack`

Existing first implementation target:

- widen `submitEnvelope` content-type/protocol-version validation;
- keep `envelopes.ciphertext_b64` as the opaque base64 payload carrier;
- add/adjust tests in `internal/httpapi/api_test.go` or adjacent focused test file.


### Sidecar dev-state paths after v0.2.40

Identity/public-bundle state remains:

- `carbonstack-comms/internal/protocol/mls/research/openmls-sidecar/.carbonstack-openmls-sidecar-state/dev/devices/<device-label>/`

Creator/Alice conversation state is now device-scoped:

- `carbonstack-comms/internal/protocol/mls/research/openmls-sidecar/.carbonstack-openmls-sidecar-state/dev/devices/<device-label>/conversations/<conversation-label>/`

Bob joined/open conversation state remains device-scoped:

- `carbonstack-comms/internal/protocol/mls/research/openmls-sidecar/.carbonstack-openmls-sidecar-state/dev/devices/<device-label>/conversations/<conversation-label>/`

Alice device-scoped Welcome/add-member state:

- `.carbonstack-openmls-sidecar-state/dev/devices/<alice-device>/conversations/<conversation-label>/welcome.bin`
- `.carbonstack-openmls-sidecar-state/dev/devices/<alice-device>/conversations/<conversation-label>/welcome-manifest.json`
- `.carbonstack-openmls-sidecar-state/dev/devices/<alice-device>/conversations/<conversation-label>/add-member-summary.json`

Alice device-scoped protected message state:

- `.carbonstack-openmls-sidecar-state/dev/devices/<alice-device>/conversations/<conversation-label>/messages/<message-label>/application-message.bin`
- `.carbonstack-openmls-sidecar-state/dev/devices/<alice-device>/conversations/<conversation-label>/messages/<message-label>/message-manifest.json`
- `.carbonstack-openmls-sidecar-state/dev/devices/<alice-device>/conversations/<conversation-label>/messages/<message-label>/message-protect-summary.json`

Bob device-scoped opened-message state:

- `.carbonstack-openmls-sidecar-state/dev/devices/<bob-device>/conversations/<conversation-label>/opened-messages/<message-label>/message-open-summary.json`

Old global creator state path is intentionally retired for new dev state:

- `.carbonstack-openmls-sidecar-state/dev/conversations/<conversation-label>/`

No compatibility or migration fallback exists. Reset `.carbonstack-openmls-sidecar-state/` for manual tests.

### Legacy sidecar dev-state paths before v0.2.40

Identity/public-bundle state:

- `carbonstack-comms/internal/protocol/mls/research/openmls-sidecar/.carbonstack-openmls-sidecar-state/dev/devices/<device-label>/`

Alice/global conversation state:

- `carbonstack-comms/internal/protocol/mls/research/openmls-sidecar/.carbonstack-openmls-sidecar-state/dev/conversations/<conversation-label>/`

Bob joined conversation state:

- `carbonstack-comms/internal/protocol/mls/research/openmls-sidecar/.carbonstack-openmls-sidecar-state/dev/devices/<device-label>/conversations/<conversation-label>/`

Current Alice v0 message artifact state:

- `.carbonstack-openmls-sidecar-state/dev/conversations/<conversation-label>/messages/message-0001/`
- `application-message.bin`
- `message-manifest.json`
- `message-protect-summary.json`

Planned v0.2.38 explicit-label Alice message artifact state:

- `.carbonstack-openmls-sidecar-state/dev/conversations/<conversation-label>/messages/<message-label>/`
- `application-message.bin`
- `message-manifest.json`
- `message-protect-summary.json`

Current Bob v0 opened-message state:

- `.carbonstack-openmls-sidecar-state/dev/devices/<device-label>/conversations/<conversation-label>/opened-messages/message-0001/`
- `message-open-summary.json`

Planned v0.2.38 explicit-label Bob opened-message state:

- `.carbonstack-openmls-sidecar-state/dev/devices/<device-label>/conversations/<conversation-label>/opened-messages/<message-label>/`
- `message-open-summary.json`

Important generated files:

- `identity-prep.json` — sanitized/non-secret.
- `identity-summary.json` — sanitized/non-secret.
- `identity-state.json` — sanitized/non-secret.
- `public-bundle-summary.json` — sanitized generated dev metadata.
- `public-bundle-manifest.json` — sanitized generated dev metadata.
- `public-bundle.keypackage.bin` — serialized public OpenMLS KeyPackage artifact; generated dev public artifact, not final onboarding format.
- device-root `provider-storage.json` — generated dev OpenMLS provider state containing the private KeyPackage bundle required for Welcome consumption; sensitive; do not inspect/paste/commit.
- Alice/global `conversation-summary.json` — sanitized generated dev conversation metadata.
- Alice/global `provider-storage.json` — generated dev OpenMLS provider/group state; sensitive; do not inspect/paste/commit.
- `welcome.bin` — generated OpenMLS Welcome carrier artifact serialized from the second `MlsMessageOut` returned by `add_members`; do not print/paste/commit; not final onboarding UX.
- `welcome-manifest.json` — sanitized generated dev metadata.
- `add-member-summary.json` — sanitized generated dev metadata.
- Bob joined `conversation-summary.json` — sanitized generated dev metadata.
- Bob joined `join-summary.json` — sanitized generated dev metadata.
- Bob joined `provider-storage.json` — generated dev OpenMLS joined group state; sensitive; do not inspect/paste/commit.
- `application-message.bin` — generated OpenMLS protected application-message artifact; do not print/paste/commit raw bytes.
- `message-manifest.json` — sanitized generated dev metadata.
- `message-protect-summary.json` — sanitized generated dev metadata.
- `message-open-summary.json` — sanitized generated dev metadata.
- `signer.json` — secret-bearing dev-only signer material; do not print, paste, inspect casually, or commit.

### Function/code surface after v0.2.38

Implemented and active:

- `validate_device_label(label)`
- `validate_conversation_label(label)`
- `create_dev_identity(device_label)`
- `load_dev_identity_status(device_label)`
- `export_dev_public_bundle_summary(device_label, write_artifact)`
- `create_dev_conversation(device_label, conversation_label)`
- `load_dev_conversation_status(device_label, conversation_label)`
- `add_dev_conversation_member(device_label, conversation_label, member_keypackage_path)`
- `join_dev_conversation(device_label, conversation_label, welcome_artifact_path)`
- `protect_dev_message(device_label, conversation_label, plaintext)`
- `open_dev_message(device_label, conversation_label, message_artifact_path)`
- `validate_plaintext_for_dev(plaintext)`
- `validate_message_artifact_path(path)`
- `conversation_state_dir(conversation_label)`
- `conversation_summary_path(conversation_label)`
- `conversation_provider_storage_path(conversation_label)`
- `conversation_welcome_artifact_path(conversation_label)`
- `conversation_welcome_manifest_path(conversation_label)`
- `conversation_add_member_summary_path(conversation_label)`
- `conversation_messages_dir(conversation_label)`
- `conversation_message_dir(conversation_label, message_label)`
- `conversation_message_artifact_path(conversation_label, message_label)`
- `conversation_message_manifest_path(conversation_label, message_label)`
- `conversation_message_protect_summary_path(conversation_label, message_label)`
- `device_provider_storage_path(device_label)`
- `device_conversations_dir(device_label)`
- `device_conversation_state_dir(device_label, conversation_label)`
- `device_conversation_summary_path(device_label, conversation_label)`
- `device_conversation_provider_storage_path(device_label, conversation_label)`
- `device_conversation_join_summary_path(device_label, conversation_label)`
- `device_conversation_message_open_summary_path(device_label, conversation_label, message_label)`
- `validate_member_keypackage_path(path)`
- `validate_welcome_artifact_path(path)`
- `parse_conversation_label(args)`
- `parse_member_keypackage_path(args)`
- `parse_welcome_artifact_path(args)`
- `parse_plaintext(args)`
- `parse_message_artifact_path(args)`
- `handle_conversation_create(args)`
- `handle_conversation_load_check(args)`
- `handle_conversation_add_member(args)`
- `handle_conversation_join(args)`
- `handle_message_protect(args)`
- `handle_message_open(args)`
- message-protect/open success/failure/missing/invalid-label printers
- `CarbonStackSidecarProvider::default()`
- `CarbonStackSidecarProvider::save_storage_to_path(path)`
- `CarbonStackSidecarProvider::load_storage_from_path(path)`
- `CAPABILITIES`
- `UNSUPPORTED_COMMANDS`
- structured `print_provider_info()`

Planned next functions/changes for v0.2.38:

- `validate_message_label(label)`
- `parse_message_label(args)`
- update `protect_dev_message` to accept `message_label`
- update `open_dev_message` to accept `message_label`
- default omitted `--message-label` to `message-0001`
- preserve one-message default test
- add two-message explicit label test
- `TestOpenMLSSidecarMessageOpenOutOfOrderTwoMessageDelivery`
- `TestOpenMLSSidecarMessageOpenDuplicateRejected`
- `TestOpenMLSSidecarMessageOpenCorruptArtifactRejected`
- `setupOpenMLSTwoMemberConversation(...)`
- `protectOpenMLSSidecarMessage(...)`
- `openOpenMLSSidecarMessage(...)`

Current sidecar command routing:

- supported:
  - `provider-info`
  - `identity-create`
  - `identity-status`
  - `public-bundle-export`
  - `conversation-create`
  - `conversation-load-check`
  - `conversation-add-member`
  - `conversation-join`
  - `message-protect`
  - `message-open`
- unsupported:
  - `state-checkpoint`
  - `state-load-check`

### API facts carried forward

Pinned dependency set:

- `openmls 0.8.1`
- `openmls_rust_crypto 0.5.1`
- `openmls_memory_storage 0.5.0`
- `openmls_traits 0.5.0`
- `tls_codec 0.4.2`

Provider/group persistence facts:

- `OpenMlsRustCrypto` contains private `key_store: MemoryStorage`, which made direct explicit save/load inconvenient.
- `openmls_memory_storage::MemoryStorage` exposes persistence helpers behind the `persistence` feature.
- `MemoryStorage::save_to_file(...)` and `MemoryStorage::load_from_file(...)` are usable for dev-local persistence.
- `CarbonStackSidecarProvider` owns `RustCrypto + MemoryStorage`, implements `OpenMlsProvider`, and explicitly saves/loads provider state.
- `MlsGroup::load(provider.storage(), &group_id)` works across sidecar command invocations when provider storage is loaded from `provider-storage.json`.

Add-member / Welcome facts:

- `KeyPackage::tls_deserialize(...)` does not exist.
- `KeyPackageIn::tls_deserialize(...)` exists.
- `KeyPackageIn` does not convert into `KeyPackage` with `.into()`.
- Correct path is `KeyPackageIn::tls_deserialize(...)` then `KeyPackageIn::validate(provider.crypto(), ProtocolVersion::default())`.
- `MlsGroup::add_members<Provider: OpenMlsProvider>(&mut self, provider, signer, &[KeyPackage])` is the membership API used.
- `add_members` returns `(MlsMessageOut, MlsMessageOut, Option<GroupInfo>)`.
- The first `MlsMessageOut` is the Commit.
- The second `MlsMessageOut` is the Welcome carrier.
- `MlsMessageOut::body()` is public and verifies `MlsMessageBodyOut::Welcome(_)`.
- `MlsMessageOut::to_bytes()` is public and calls `tls_serialize_detached()`.
- `MlsMessageOut::into_welcome()` exists only under `#[cfg(any(feature = "test-utils", test))]` and is not used.
- `welcome.bin` currently contains the outer Welcome carrier `MlsMessageOut` bytes, not a separately extracted inner `Welcome` value.
- `group.merge_pending_commit(&provider)` is called after Welcome export.
- Mutated provider storage is saved after merge.
- Post-add `conversation-load-check` reloads mutated storage and reports `member_count=2`.

Join facts validated in v0.2.34:

- `MlsMessageIn::tls_deserialize(&mut welcome_bytes.as_slice())` works for the v0.2.32 `welcome.bin`.
- `.extract()` yields `MlsMessageBodyIn::Welcome(welcome)`.
- `StagedWelcome::new_from_welcome(&provider, &join_config, welcome, None)` works with the current add-member output.
- `StagedWelcome::into_group(&provider)` returns Bob's joined `MlsGroup`.
- `ratchet_tree=None` worked for the current flow; no ratchet tree artifact was required in v0.2.34.
- Bob's provider storage must be loaded from device-root `provider-storage.json` before `new_from_welcome`.
- Bob's joined group can be saved and reloaded from device-scoped joined `provider-storage.json`.
- Bob join reports `member_count=2` and `epoch=GroupEpoch(1)`.
- Bob join group_id_ref matches Alice/add-member group_id_ref.

Message protect/open facts validated in v0.2.36 and reconfirmed in v0.2.38:

- `MlsGroup::create_message(&provider, &signer, plaintext.as_bytes())` works for Alice after add-member/commit merge.
- The returned protected application message is an `MlsMessageOut`.
- `MlsMessageOut::to_bytes()` is the protected artifact serialization method.
- `application-message.bin` stores the protected `MlsMessageOut` bytes.
- `MlsMessageIn::tls_deserialize_exact_bytes(...)` works for `application-message.bin`.
- `try_into_protocol_message()` works for converting deserialized message to processable protocol message.
- `MlsGroup::process_message(&provider, protocol_message)` works for receiver joined/creator group where valid.
- `ProcessedMessageContent::ApplicationMessage(application_message)` is the expected processed content type.
- `application_message.into_bytes()` returns the original plaintext bytes.
- UTF-8 decoding produced `plaintext_utf8="hello bob"`.
- `create_message` and `process_message` are persistence-relevant.
- Provider storage is saved after both protect and open.
- Both message operations reload-prove group state.
- Visible epoch stayed at `GroupEpoch(1)` for the one-message proof, but this must not be generalized to all future message flows.
- OpenMLS `create_message` rejects pending proposals and returns a private message carrier after encryption.
- OpenMLS `process_message` / `unprotect_message` persists message secrets after private-message processing modifies the secret tree.

---

## 5. Phase Timeline Summary

### Phase 2C closure carried forward

Phase 2C remains closed for mainline work. It validated OpenMLS feasibility through scratch probes, same-process reload, MemoryStorage file persistence, sanitized fixtures, Go-side fixture parsing, event taxonomy, negative fixture mapping, provider trust-decision mapping, and fixture-backed trust-decision tests. Reopen only for targeted fixes.

### v0.2.22 through v0.2.36 carried forward

The earlier Phase 2D ladder remains as recorded:

1. v0.2.22 real dev-only identity creation.
2. v0.2.23 identity-status / load-check.
3. v0.2.24 summary-only public-bundle export.
4. v0.2.25 serialized KeyPackage artifact export.
5. v0.2.26 conversation lifecycle plan + API recon.
6. v0.2.27 conversation-create implementation.
7. v0.2.28 add-member / Welcome plan + recon.
8. v0.2.29 persistence honesty repair.
9. v0.2.30 provider/group persistence.
10. v0.2.31 add-member / Welcome skeleton + API recon.
11. v0.2.32 conversation-add-member / Welcome export implementation.
12. v0.2.33 conversation-join skeleton + API recon.
13. v0.2.34 conversation-join implementation.
14. v0.2.34 join command wiring cleanup at `2f4b69e`.
15. v0.2.35 message-protect/open skeleton + API recon.
16. v0.2.36 message-protect/open implementation.

### v0.2.38 provider-info cleanup

Landed:

- `27dc33d feat: add OpenMLS sidecar explicit message labels`

Problem:

- Provider-info raw string JSON had repeatedly become a command-list patch hazard.
- v0.2.36 specifically broke provider-info with a missing comma after `"conversation-join"`.
- More command churn was expected for multi-message labels and later state/runtime commands.

Cleanup:

- Introduced/used `CAPABILITIES`.
- Preserved `UNSUPPORTED_COMMANDS`.
- Rewrote `print_provider_info()` using `serde_json::json!`.
- Preserved supported commands:
  - `provider-info`
  - `identity-create`
  - `identity-status`
  - `public-bundle-export`
  - `conversation-create`
  - `conversation-load-check`
  - `conversation-add-member`
  - `conversation-join`
  - `message-protect`
  - `message-open`
- Preserved unsupported commands:
  - `state-checkpoint`
  - `state-load-check`
- Patched stale internal Rust test expectations so `message-protect` and `message-open` are not considered unsupported.

Outcome:

- Provider-info command list is now structurally generated instead of comma-sensitive raw JSON.

### v0.2.38 multi-message continuity plan + API recon

Landed:

- `680ce12 docs: plan OpenMLS multi-message continuity`
- `0135902 docs: record OpenMLS multi-message continuity result`

Docs/76 planned:

- explicit `--message-label`;
- `message-0001` default compatibility;
- two sequential messages:
  - `message-0001` / `"hello bob 1"`;
  - `message-0002` / `"hello bob 2"`;
- preserve Alice-global/Bob-device-scoped asymmetry for one more proof rung;
- avoid autogenerated IDs, Cypher routing, Comms runtime, trust-state mutation, and Alice state migration during the first multi-message proof.

Docs/77 recon confirmed:

- `MlsGroup::create_message`:
  - checks active group;
  - rejects pending proposals;
  - creates authenticated application content;
  - encrypts;
  - resets AAD;
  - returns `MlsMessageOut::from_private_message(...)`.
- `MlsGroup::process_message`:
  - calls `unprotect_message`;
  - private messages modify the secret tree;
  - OpenMLS writes message secrets back to provider storage after decrypt;
  - flows into `process_unverified_message`.
- Save-after-protect and save-after-open are required for multi-message continuity.
- Prior scratch work already proved two sequential Alice-to-Bob OpenMLS messages in-process, but sidecar persistence across separate invocations remains unvalidated.

Outcome:

- v0.2.38 should implement explicit message labels and a two-message sequential proof.
- Do not migrate Alice state yet.
- Do not start Cypher/Comms runtime yet.


### v0.2.39 ordering / replay / corrupt artifact recon + tests

Landed in `carbonstack`:

- `01cd442 docs: plan OpenMLS message ordering replay recon`
- `9411aeb docs: record OpenMLS message ordering replay recon`
- `01bc4a8 docs: record OpenMLS corrupt message artifact behavior`

Landed in `carbonstack-comms`:

- `80685a7 test: cover OpenMLS sidecar message ordering replay cases`

Docs/79 planned the next safety checkpoint after v0.2.38:

- open two messages in order after both were protected;
- open two messages out of order;
- attempt duplicate/replay open;
- attempt corrupt/truncated artifact open;
- record exact OpenMLS/sidecar errors before Cypher routing.

Docs/80 source recon confirmed:

- `MlsGroup::process_message(...)` routes into `unprotect_message(...)`;
- private `ProtocolMessage::PrivateMessage(_)` sets `will_modify_secret_tree`;
- OpenMLS writes message secrets back to provider storage after private-message decrypt;
- duplicate/replay behavior can be enforced by the secret tree / sender ratchet state;
- sidecar error taxonomy currently maps replay/secret-reuse to generic `message_open_failed` / `checkpoint.failed`.

Manual probes validated:

- ordered batch delivery succeeds: protect `message-0001`, protect `message-0002`, open `message-0001`, open `message-0002`;
- out-of-order two-message delivery succeeds for the exact current same-sender Alice-to-Bob flow: protect `message-0001`, protect `message-0002`, open `message-0002`, then open `message-0001`;
- duplicate/replay open fails with `SecretReuseError`, surfaced as `message_open_failed`, `checkpoint.failed`, warning, exit code `3`;
- initial corrupt-artifact probe was inconclusive because the corrupt file was not created and the sidecar only observed a missing file;
- corrected corrupt-artifact probe used resolved paths, verified both good/bad files existed, truncated without printing bytes, and produced `message_artifact_invalid`, `provider.message.invalid`, `EndOfStream`, warning, exit code `3`.

Go tests added in `openmls_sidecar_provider_info_test.go`:

- `TestOpenMLSSidecarMessageOpenOutOfOrderTwoMessageDelivery`
- `TestOpenMLSSidecarMessageOpenDuplicateRejected`
- `TestOpenMLSSidecarMessageOpenCorruptArtifactRejected`

Outcome:

- v0.2.39 is a contract-test checkpoint, not a sidecar feature checkpoint.
- The simple MLS artifact delivery behavior is much better understood before Cypher routing.
- Next safest rung is Alice state layout cleanup plan/recon, not immediate runtime integration.


### v0.2.40 Alice device-scoped state layout cleanup

Landed in `carbonstack`:

- `9e182a0 Add docs/82-openmls-sidecar-alice-device-scoped-state-layout-plan-v0.md`
- `f58db2d docs: record OpenMLS Alice device-scoped state result`
- `2d1d657 docs: record OpenMLS Alice device-scoped state result`

Landed in `carbonstack-comms`:

- `66c9c8a refactor: use device-scoped OpenMLS conversation state`
- `7ccf545 refactor: use device-scoped OpenMLS conversation state`

Result:

- Alice/creator conversation state no longer uses the old global `dev/conversations/<conversation-label>/` tree.
- Creator-side operations now use `dev/devices/<device-label>/conversations/<conversation-label>/`.
- Affected operations:
  - `conversation-create`;
  - `conversation-load-check`;
  - `conversation-add-member`;
  - `message-protect`.
- Already device-scoped operations remain aligned:
  - `conversation-join`;
  - `message-open`.
- The command surface did not change.
- The cleanup is a hard-cut dev-state break with no migration compatibility.
- Test fallout was limited to stale old-path assumptions in `TestOpenMLSSidecarConversationCreate`.

Outcome:

- The OpenMLS sidecar now has a more coherent device-owned state layout before Cypher routing design.
- Next safe rung is Cypher MLS artifact routing design docs/recon.


### v0.2.41 Phase 2D mainline closure

Landed in `carbonstack`:

- `7013406 docs: define OpenMLS sidecar Phase 2D closure checklist`
- `f75b529 docs: record OpenMLS wrong-target message open behavior`
- `45851ef docs: record OpenMLS sidecar Phase 2D mainline closure`

Landed in `carbonstack-comms`:

- `cf8ac20 fix: write OpenMLS protected messages under device-scoped state`
- `bb5bd47 test: cover OpenMLS sidecar wrong-target message open cases`
- `65c202b test: cover OpenMLS sidecar bidirectional message flow`

Result:

- v0.2.41 closes the mainline Phase 2D OpenMLS sidecar research path.
- Stale-path sweep found and fixed one remaining operational old-path issue in `protect_dev_message`.
- `message-protect` now writes protected message artifacts under the sender device-scoped conversation path.
- Wrong-device and wrong-conversation `message-open` behavior is now contract-tested.
- Bidirectional Alice/Bob message flow is now contract-tested.
- The sidecar now validates the dev-local lifecycle needed before Cypher artifact relay research:
  - identity-create;
  - public-bundle export and serialized KeyPackage artifact;
  - conversation-create;
  - add-member and Welcome export;
  - conversation-join;
  - message-protect/open;
  - explicit message labels;
  - two-message continuity;
  - out-of-order same-sender two-message open;
  - duplicate/replay rejection;
  - corrupt/truncated artifact rejection;
  - wrong-target rejection;
  - bidirectional Alice↔Bob flow;
  - device-scoped creator/joiner state.
- This is not a production E2EE claim, not Comms runtime integration, not Cypher routing, and not research-to-implementation promotion.
- The next mainline direction is Cypher minimal opaque MLS artifact relay research.
- During validation, the local Windows machine BSODed under heavy build/test activity. The dump pointed to `WIN32K_CRITICAL_FAILURE (0x164)` in the `dwm.exe` / `win32kbase.sys` path with Vanguard and Avast drivers loaded. Use low-parallelism `go test -p 1` until local environment stability is addressed. This is recorded as a local-environment case study, not a CarbonStack test failure.



### v0.2.42 maintainability promotion plan/recon

Landed in `carbonstack`:

- `314a3d5 docs: plan OpenMLS sidecar maintainability promotion`

Docs added:

- `docs/87-openmls-sidecar-current-state-index-v0.md`
- `docs/88-openmls-sidecar-maintainability-promotion-plan-v0.md`
- `docs/89-openmls-sidecar-module-split-plan-v0.md`
- `docs/90-openmls-sidecar-test-suite-split-plan-v0.md`
- `docs/91-openmls-sidecar-artifact-ownership-map-v0.md`
- `docs/92-openmls-sidecar-command-schema-matrix-v0.md`

Purpose:

- Stop treating the next step as vague cleanup.
- Preserve the research sidecar as the known-good Phase 2D reference.
- Plan a promoted maintained sidecar scaffold above the research path.
- Plan Rust module split, Go test split, README/current-state cleanup, and Cypher relay recon sequencing.

Important planned future flow:

1. v0.2.43 — copy/promote sidecar scaffold:
   - create promoted location above research;
   - keep behavior identical;
   - prove equivalence against the Phase 2D closure test suite.
2. v0.2.44 — split Rust modules:
   - modularize promoted sidecar;
   - remove dead old global helper residue;
   - keep command/schema behavior unchanged.
3. v0.2.45 — split Go tests:
   - point tests at the promoted sidecar;
   - split helper/provider/identity/public-bundle/conversation/message/negative tests;
   - preserve coverage and test names where practical.
4. v0.2.46 — docs/README cleanup:
   - update sidecar README;
   - update current-state docs;
   - update known-good commands to device-scoped paths or envelope path hints;
   - remove stale warnings.
5. v0.2.47 — Cypher minimal opaque MLS artifact relay recon:
   - design Welcome and application-message relay;
   - use artifact ownership map and command/schema matrix;
   - avoid plaintext/provider-storage leakage.

Outcome:

- v0.2.42 is a stable docs-only planning checkpoint.
- It deliberately delays Cypher until maintainability work gives the sidecar a cleaner maintained structure.
- No code behavior changed.


---


### v0.2.43 promoted OpenMLS sidecar scaffold

Landed in `carbonstack`:

- `2c0f576 docs: record OpenMLS sidecar promotion scaffold`

Landed in `carbonstack-comms`:

- `b44adbd refactor: promote OpenMLS sidecar scaffold`

Result:

- The known-good Phase 2D research sidecar was copied from:
  - `internal/protocol/mls/research/openmls-sidecar`
  to:
  - `internal/protocol/mls/openmls-sidecar`
- The research sidecar remains intact as a reference.
- The promoted sidecar is now the active Go contract-test target.
- The promoted scaffold is intentionally behavior-identical at this checkpoint.
- The command surface, JSON envelopes, path hints, provider events, exit codes, and OpenMLS lifecycle were not intentionally changed.
- Direct Rust validation of the promoted scaffold passed (`cargo fmt`, `cargo check`, `cargo test`).
- Go Phase 2D contract validation passed against the promoted scaffold.
- Rust dead-code warnings remain for unused old global conversation helpers. These are scheduled for cleanup during the promoted sidecar Rust module split.

Blunder/repair continuity:

- A first local scaffold-promotion commit accidentally staged generated OpenMLS dev state under the new promoted sidecar path.
- Included generated files covered secret-bearing and sensitive dev state patterns such as `signer.json`, `provider-storage.json`, `welcome.bin`, `application-message.bin`, public-bundle artifacts, and message artifacts.
- The bad commit was caught before push and was not allowed to leave the machine.
- The local commit was reset, generated state/build output was deleted, ignore coverage was added for promoted/future sidecar generated state and build output, and the clean commit was recreated.
- This is now a permanent lesson: after any sidecar path copy or test run, verify staged files do not include `.carbonstack-openmls-sidecar-state/`, `target/`, signer/provider storage, or raw MLS artifacts before committing.

Outcome:

- v0.2.43 completes the behavior-preserving scaffold promotion rung.
- Next safest rung is v0.2.44: split promoted Rust modules without behavior change.



## 6. Successes and Blunders Preserved for Continuity



### New successes from v0.2.41

- Phase 2D mainline research was closed cleanly after stale-path sweep, wrong-target tests, and bidirectional proof.
- `protect_dev_message` now writes message artifacts under the sender device-scoped conversation tree, completing the intended v0.2.40 state-layout change.
- Wrong-device `message-open` is contract-tested and fails safely before decrypt/process when receiver device/conversation provider storage is missing.
- Wrong-conversation `message-open` is contract-tested and fails safely with the same missing provider-storage shape.
- Bidirectional Alice↔Bob messaging is contract-tested:
  - Alice can send to Bob;
  - Bob can send back to Alice;
  - Alice's creator state after add-member/merge can process Bob's private application message.
- Docs/85 and docs/86 define and close the mainline Phase 2D closure boundary.
- The next phase is clearly Cypher minimal opaque MLS artifact relay research, not Android or runtime promotion.

### New blunders / repair lessons from v0.2.41

- The v0.2.40 state-layout cleanup initially missed one operational call-site cluster: `protect_dev_message` still used old global message artifact helpers while reading device-scoped provider storage. Stale-path sweeps must search both read paths and write paths.
- Manual probes confirmed wrong-target behavior before assertions were locked into Go tests. This avoided guessing error codes/events.
- Heavy local validation triggered a Windows BSOD on the operator machine. Debugging suggested a local Windows/DWM/win32k/security-driver stability issue, not repo corruption. Use `go test -p 1` for heavy validation on this machine until the local environment is stabilized.
- Treat repeated Vanguard/Avast/driver signals as local-environment context only; do not encode them as project assumptions beyond the timeline note.



### New successes from v0.2.42

- The post-Phase 2D maintainability problem is now explicitly scoped instead of left as vague cleanup.
- `docs/87` provides a current-state index so future branches can distinguish current truth from historical docs.
- `docs/88` defines research-vs-promoted sidecar semantics and preserves the research sidecar as known-good reference.
- `docs/89` defines the Rust module split target and assigns ownership to `cli`, `envelope`, `errors`, `paths`, `identity`, `public_bundle`, `conversation`, and `message` modules.
- `docs/90` defines the Go sidecar test split and keeps existing behavior/coverage as the preservation target.
- `docs/91` maps artifacts by producer, consumer, sensitivity, and relay eligibility before Cypher routing begins.
- `docs/92` records the current sidecar command/schema matrix and important failure shapes.
- The planned cleanup ladder v0.2.43-v0.2.47 is now explicit and can be followed without disrupting Phase 2D continuity.

### New blunders / repair lessons from v0.2.42

- No implementation blunder occurred because v0.2.42 was docs-only.
- The main process lesson is that maintainability work should be planned as a series of narrow equivalence-preserving rungs, not a single “make everything clean” mega-refactor.
- The research sidecar should not be deleted or heavily churned during the promotion scaffold; it is the known-good reference while the promoted implementation becomes the maintained target.
- Historical docs should not be rewritten wholesale just because some paths are stale; add current-state indexes and clear current/historical boundaries instead.

### New successes from v0.2.40

- Alice/creator state layout asymmetry was resolved before starting Cypher routing.
- `conversation-create`, `conversation-load-check`, `conversation-add-member`, and `message-protect` now use device-scoped paths.
- Bob join/open paths were already device-scoped and remained stable.
- The command surface stayed unchanged; path hints carry the new layout.
- The hard-cut dev-state break avoided legacy path fallback complexity.
- Go tests were patched to follow envelope path hints rather than hardcoded global paths.

### New blunders / repair lessons from v0.2.40

- The initial helper insertion patch looked for an exact `device_conversation_join_summary_path` block and failed because `cargo fmt` compacted the helper body. Use line markers or whitespace-flexible regex for helper insertion.
- The first Go path patch searched for single `filepath.Join(..., "conversation-summary.json")` expressions, but the test built paths through `stateDir`. Inspect the function body before exact block replacements.
- `TestOpenMLSSidecarConversationCreate` kept checking the retired global path until the local path-construction block was replaced.
- A temporary `stateDir` variable became unused after path hint refactor and caused a Go compile failure. Either assert the directory or remove the local immediately.
- Duplicate-looking commit messages appeared in both `carbonstack` and `carbonstack-comms`; the repo heads from the final snapshot remain the authority.

### Successes

- v0.2.39 converted ordering/replay/corrupt-artifact recon into Go contract tests without new Rust feature work.
- Out-of-order two-message delivery is now regression-tested for the exact current same-sender dev-local flow.
- Duplicate/replay open is now regression-tested and fails through OpenMLS `SecretReuseError`.
- Truncated/corrupt message artifact handling is now regression-tested as `message_artifact_invalid` / `provider.message.invalid` / `EndOfStream`.
- The project now has a clearer boundary between Cypher delivery metadata concerns and OpenMLS secret-tree/replay behavior.

- v0.2.38 cleaned up the known provider-info JSON hazard before it caused more command-list failures.
- The `[recon + doc] -> [implement + patch]` workflow continued successfully.
- Multi-message continuity was planned before implementation.
- The recon narrowed to targeted OpenMLS files instead of wrestling huge output:
  - `src/group/mls_group/application.rs`
  - `src/group/mls_group/processing.rs`
  - current sidecar `state.rs`
  - current sidecar `main.rs`
  - current Go sidecar contract test.
- OpenMLS source recon confirmed the existing v0.2.36 persistence discipline.
- The next implementation can stay narrow: message labels + two sequential messages.
- Alice state migration was deliberately deferred rather than mixed into the multi-message proof.

### Blunders / repair lessons carried forward from v0.2.30–v0.2.36

- Provider-info in `main.rs` was raw/hand-formatted and fragile; v0.2.38 cleaned it up.
- Exact block patches often fail after `gofmt`; use line windows, narrow regex, or line-walking by function scope.
- When a command becomes real, update route, provider-info capability/unsupported lists, Go provider-info assertions, unsupported-command tests, and Rust internal tests together.
- `KeyPackage::tls_deserialize` is wrong; `KeyPackageIn::tls_deserialize` + `.validate(...)` is correct.
- `MlsMessageOut::into_welcome()` is test-gated and not valid for sidecar production code.
- `welcome.bin` contains outer MLS message carrier bytes, not inner `Welcome` bytes.
- First `conversation-join` manual probe failed with `NoMatchingKeyPackage`; the fix was Bob device provider storage persistence, not Welcome artifact format changes.
- Public KeyPackage artifact bytes alone are insufficient for join; Bob also needs local private provider state generated alongside the KeyPackage.
- `ConversationStatePathHint` for join is a directory, not a file.
- The broad 3C recon search was too large; use targeted source files/line ranges once function names are known.
- Inspect dirty diffs before checkout/revert; the v0.2.34 dirty diff was essential join command wiring.
- The first v0.2.36 state implementation assumed nonexistent `load_signer(...)`; use existing `load_dev_identity_status` + `read_json_file(status.signer_path)`.
- The first v0.2.36 protect/open implementation tried to read raw `group_id` from summaries; use deterministic `carbonstack-openmls-dev-conversation:{conversation_label}` group ID.

### New blunders / repair lessons from v0.2.39

- The first corrupt-artifact probe was invalid because PowerShell failed to resolve the good artifact path, so the corrupt file was never created; the observed sidecar error was only missing-file behavior.
- Corrected corrupt-artifact probing must verify both `$Good` and `$Bad` exist before invoking `message-open`.
- The Go corrupt-artifact test initially passed an absolute `badPath` and received `conversation_or_message_missing`; passing a relative path hint from the sidecar crate root matches the normal sidecar artifact style better.
- Repeated patch attempts left `badPathHint` referenced before definition; inspect the target range after patching and before re-running tests.
- One PowerShell patch corrupted source with `.\.git\t.Fatalf(...)`; this reinforced that here-strings must not accidentally interpolate shell/path fragments into source.
- Use `@' ... '@` / `@" ... "@` style cautiously and prefer function-scoped or line-scoped edits when exact matching fails repeatedly.
- `Select-String` output wraps long Go lines; do not assume the actual source line is split just because the terminal display wraps it.

### New blunders / repair lessons from v0.2.38

- Cleanup patches must account for `gofmt` collapsing simple arrays into one-line constants. The initial cleanup patch looked for a multi-line `UNSUPPORTED_COMMANDS` block, but the actual file had a single-line block:
  - `const UNSUPPORTED_COMMANDS: &[&str] = &["state-checkpoint", "state-load-check"];`
- The stale internal Rust unit test still asserted that `message-protect` and `message-open` were unsupported even after v0.2.36. The cleanup had to patch this too.
- Broad OpenMLS recon over all group files (`3A`) produced too much output and was omitted. That is acceptable when targeted high-value recon files answer the question.
- Multi-message ordering/out-of-order behavior is not proven yet. Source hints imply private message processing mutates secret tree state; do not infer full out-of-order/replay behavior without tests.
- Fixed `message-0001` remains valid for default compatibility but is now a known implementation blocker.

### Permanent process lesson

The workflow pattern is now trusted for CarbonStack:

> **[recon + doc] -> [implement + patch]**

Meaning:

1. Define the command/API surface in docs.
2. Recon exact upstream APIs and state effects before implementation.
3. Land docs in `carbonstack`.
4. Implement narrowly in `carbonstack-comms`.
5. Patch compile/runtime/test failures as they appear.
6. Validate manually and through Go tests.
7. Record result docs.
8. Cut LogDoc + lean JSON when the state is worth preserving.

This pattern worked well for:
- v0.2.31 -> v0.2.32 add-member / Welcome export.
- v0.2.33 -> v0.2.34 conversation-join.
- v0.2.35 -> v0.2.36 message-protect/open.
- v0.2.38 cleanup + multi-message recon.

It should remain the default for future protocol-critical CarbonStack rungs.

---

## 7. Event and Trust State at v0.2.38

### Stabilized operational events carried forward

- `provider.command.unsupported`
- `provider.command.invalid`
- `provider.command.not_implemented`
- `provider.identity.prep_state_written`
- `provider.identity.created`
- `provider.identity.loaded`
- `provider.identity.missing`
- `provider.identity.exists`
- `provider.public_bundle.exported`
- `provider.conversation.created`
- `provider.conversation.exists`
- `conversation.loaded`
- `checkpoint.failed`
- `provider.secret.material.unavailable`

### Add-member / Welcome events emitted since v0.2.32

`conversation-add-member` success emits:

- `provider.conversation.member_added`
- `provider.welcome.exported`

Both are currently `severity=info` and `trust_relevant=false` in the dev-sidecar rung.

### Join events emitted since v0.2.34

`conversation-join` success emits:

- `conversation.joined`
- `storage.saved`

Both are currently emitted as dev-sidecar events and do not mutate trust-state storage.

### Message events emitted since v0.2.36

`message-protect` success emits:

- `message.protected`
- `storage.saved`

`message-open` success emits:

- `message.opened`
- `storage.saved`

Do not mutate trust state from these events yet. Future product integration may need distinct user-visible semantics for membership changes, message-processing errors, replay/epoch issues, and identity/trust drift.

### Current unsupported command state

At v0.2.38, the following remain unsupported:

- `state-checkpoint`
- `state-load-check`

This is expected and correct.

---

## 8. Known Good Commands and Workflow

### Permanent CarbonStack workflow

Use this as a trusted process pattern for protocol-critical rungs:

```text
[recon + doc] -> [implement + patch]
```

Expanded:

```text
1. Scope the next small proof.
2. Write skeleton/plan docs in carbonstack.
3. Recon exact upstream/library APIs.
4. Record API recon in carbonstack.
5. Implement only the scoped proof in carbonstack-comms.
6. Patch compile failures first.
7. Patch manual/runtime failures second.
8. Patch Go contract tests third.
9. Run targeted tests, full tests, artifact guard.
10. Record result docs in carbonstack.
11. Validate docs.
12. Snapshot all repos.
13. Cut LogDoc + lean JSON when the checkpoint is worth preserving.
```

Why this is trusted:

- It prevents speculative implementation against unknown OpenMLS APIs.
- It catches false persistence claims before building on them.
- It preserves blunders and exact API facts.
- It keeps doctrine/spec/docs ahead of code.
- It keeps implementation rungs narrow enough to debug.

### From `carbonstack`

- `cd C:\▓▓\repos\carbonstack_umbrella\carbonstack; powershell -ExecutionPolicy Bypass -File .\scripts\validate-local.ps1`

### From `carbonstack-comms`

- `cd C:\▓▓\repos\carbonstack_umbrella\carbonstack-comms; powershell -ExecutionPolicy Bypass -File .\scripts\check-no-rust-artifacts.ps1`
- `cd C:\▓▓\repos\carbonstack_umbrella\carbonstack-comms; go test ./internal/protocol`
- `cd C:\▓▓\repos\carbonstack_umbrella\carbonstack-comms; go test ./...`
- `cd C:\▓▓\repos\carbonstack_umbrella\carbonstack-comms; go test ./internal/protocol -run "TestOpenMLSSidecarProviderInfoCommand|TestOpenMLSSidecarUnsupportedCommandEnvelope|TestOpenMLSSidecarMessageProtectOpenOneWay"`

### From `openmls-sidecar`

- `cd C:\▓▓\repos\carbonstack_umbrella\carbonstack-comms\internal\protocol\mls\research\openmls-sidecar; cargo check`
- `cd C:\▓▓\repos\carbonstack_umbrella\carbonstack-comms\internal\protocol\mls\research\openmls-sidecar; cargo test`
- `cd C:\▓▓\repos\carbonstack_umbrella\carbonstack-comms\internal\protocol\mls\research\openmls-sidecar; cargo run -- provider-info`
- `cd C:\▓▓\repos\carbonstack_umbrella\carbonstack-comms\internal\protocol\mls\research\openmls-sidecar; cargo run -- identity-create --device-label carbonstack-alice-device`
- `cd C:\▓▓\repos\carbonstack_umbrella\carbonstack-comms\internal\protocol\mls\research\openmls-sidecar; cargo run -- identity-create --device-label carbonstack-bob-device`
- `cd C:\▓▓\repos\carbonstack_umbrella\carbonstack-comms\internal\protocol\mls\research\openmls-sidecar; cargo run -- public-bundle-export --device-label carbonstack-bob-device --write-artifact`
- `cd C:\▓▓\repos\carbonstack_umbrella\carbonstack-comms\internal\protocol\mls\research\openmls-sidecar; cargo run -- conversation-create --device-label carbonstack-alice-device --conversation-label carbonstack-test-conversation`
- `cd C:\▓▓\repos\carbonstack_umbrella\carbonstack-comms\internal\protocol\mls\research\openmls-sidecar; cargo run -- conversation-add-member --device-label carbonstack-alice-device --conversation-label carbonstack-test-conversation --member-keypackage .\.carbonstack-openmls-sidecar-state\dev\devices\carbonstack-bob-device\public-bundle.keypackage.bin`
- `cd C:\▓▓\repos\carbonstack_umbrella\carbonstack-comms\internal\protocol\mls\research\openmls-sidecar; cargo run -- conversation-join --device-label carbonstack-bob-device --conversation-label carbonstack-test-conversation --welcome .\.carbonstack-openmls-sidecar-state\dev\conversations\carbonstack-test-conversation\welcome.bin`
- `cd C:\▓▓\repos\carbonstack_umbrella\carbonstack-comms\internal\protocol\mls\research\openmls-sidecar; cargo run -- message-protect --device-label carbonstack-alice-device --conversation-label carbonstack-test-conversation --plaintext "hello bob"`
- `cd C:\▓▓\repos\carbonstack_umbrella\carbonstack-comms\internal\protocol\mls\research\openmls-sidecar; cargo run -- message-open --device-label carbonstack-bob-device --conversation-label carbonstack-test-conversation --message .\.carbonstack-openmls-sidecar-state\dev\conversations\carbonstack-test-conversation\messages\message-0001\application-message.bin`

Expected `message-protect` stdout:

- `"message_protected": true`
- `"protected_message_written": true`
- `"provider_storage_loaded": true`
- `"provider_storage_written": true`
- `"group_reloadable": true`
- `"private_material_included": false`

Expected `message-open` stdout:

- `"message_opened": true`
- `"plaintext_utf8": "hello bob"`
- `"plaintext_len": 9`
- `"provider_storage_loaded": true`
- `"provider_storage_written": true`
- `"group_reloadable": true`
- `"private_material_included": false`

Safe inspection only:

- list generated state file names and lengths;
- inspect sanitized summaries/manifests if needed;
- do not inspect/paste `signer.json`, `provider-storage.json`, raw provider storage, raw group state, raw Welcome bytes, raw KeyPackage bytes, raw application message bytes, or raw private material.

### Planned v0.2.38 manual proof shape

After implementation adds `--message-label`:

```text
message-protect --device-label carbonstack-alice-device --conversation-label carbonstack-test-conversation --message-label message-0001 --plaintext "hello bob 1"
message-open --device-label carbonstack-bob-device --conversation-label carbonstack-test-conversation --message-label message-0001 --message <message-0001 application-message.bin>

message-protect --device-label carbonstack-alice-device --conversation-label carbonstack-test-conversation --message-label message-0002 --plaintext "hello bob 2"
message-open --device-label carbonstack-bob-device --conversation-label carbonstack-test-conversation --message-label message-0002 --message <message-0002 application-message.bin>
```

Expected:

- Bob recovers `"hello bob 1"`.
- Bob recovers `"hello bob 2"`.
- Both message labels have distinct artifact paths.
- Both protect/open operations save provider storage.
- Both protect/open operations reload-prove group state.
- No forbidden private material appears in stdout.

---

## 9. Not Validated

- Message ID generation beyond fixed/default `message-0001`.
- Bidirectional message flow.
- Multiple members beyond Alice/Bob.
- Bob device-scoped load-check command.
- Alice conversation-create migration to device-scoped state.
- Runtime OpenMLS provider mapping into `internal/protocol` send/open paths.
- CarbonStackComms CLI OpenMLS integration.
- CarbonStackCypher routing of MLS application payloads.
- Trust-state storage integration with provider events.
- User-facing warning/block/reverify UX from `ProviderTrustDecision`.
- Production secure vault / encrypted local storage.
- Hardware-key enrollment/recovery/revocation.
- Replay resistance, metadata privacy, hostile-server security harness.
- Android client, CarbonStackOS, CI.

---

## 10. Hazards and Lessons

- `signer.json` is real secret-bearing dev state. Never print, paste, inspect casually, commit, or expose it.
- Device-root `provider-storage.json` after public-bundle export contains private KeyPackage bundle provider state and is required for Welcome consumption. Never print, paste, inspect casually, commit, or expose it.
- Joined conversation `provider-storage.json` contains Bob's joined group state. Never print, paste, inspect casually, commit, or expose it.
- Alice/global `provider-storage.json` contains Alice's group state and now Alice message ratchet/secret state after protect. Never print, paste, inspect casually, commit, or expose it.
- Bob joined `provider-storage.json` contains Bob's group state and now Bob message-processing state after open. Never print, paste, inspect casually, commit, or expose it.
- `application-message.bin` is real protected OpenMLS protocol material. Never print, paste, inspect casually, commit, or expose raw bytes.
- `welcome.bin` is real generated Welcome carrier protocol material. Never print, paste, inspect casually, commit, or expose it.
- `public-bundle.keypackage.bin` is public protocol material but generated dev state. It is not final CarbonStack onboarding format.
- `.carbonstack-openmls-sidecar-state/` must remain ignored.
- Message summaries/manifests must remain sanitized metadata only.
- Message-open plaintext stdout is a bounded dev proof only, not final UX.
- Alice state is still global; Bob state is device-scoped. Keep this asymmetry explicit until a migration checkpoint.
- Fixed/default `message-0001` is acceptable for backward compatibility but not a final message ID strategy.
- Save provider storage after both protect and open even when visible epoch does not change.
- Source recon confirms private-message processing can write message secrets back to provider storage.
- Add-member, join, protect, and open remain dev-sidecar events and do not mutate trust state yet.
- Provider-info now uses structured JSON, but future command moves still require synchronized updates to constants/tests.
- Broad regex/string patches can corrupt files; inspect target ranges, patch narrowly, run formatter immediately.
- PowerShell wildcard/regex matching is not safe for arbitrary source-code markers. Use `.Contains()`, `-SimpleMatch`, or escaped regex for markers containing parentheses/brackets.
- Use `$Old = @' ... '@` / `$New = @' ... '@` blocks for PowerShell patching when possible, but switch to line-walk patches when block matching keeps failing.
- Run Rust artifact guard before Rust-related commits.
- Keep `target/`, `.exe`, `.pdb`, signer JSON, provider state, MemoryStorage files, raw group state, Welcome artifacts, protected message artifacts, and generated sidecar state out of Git.

---

## 11. Allowed Claims

Allowed:

- Phase 2D sidecar bootstrap exists.
- OpenMLS sidecar provider-info command builds/runs and emits sanitized JSON.
- Provider-info command list is now structured JSON instead of fragile raw JSON.
- Go-side tests invoke and parse OpenMLS sidecar envelopes.
- `identity-create` generates dev-only OpenMLS identity/signing material.
- `identity-status` loads existing dev-only sidecar identity state and verifies public identity ref.
- `public-bundle-export` summary/artifact mode generates a real OpenMLS KeyPackage and saves private device provider state needed for later Welcome consumption.
- `conversation-create` creates a dev-local one-member OpenMLS conversation/group.
- `conversation-load-check` loads dev-local provider storage and proves `MlsGroup::load` works across sidecar command invocations.
- `conversation-add-member` consumes Bob's public KeyPackage artifact, adds Bob to Alice’s group, exports a Welcome carrier artifact, merges pending commit, and saves provider storage.
- `conversation-join` consumes the existing `welcome.bin`, stages the Welcome, joins Bob, saves Bob joined group state, and proves reloadability.
- Bob joined group state is device-scoped.
- The dev-local MLS membership onboarding lifecycle exists:
  - `create -> add-member -> Welcome export -> join`.
- `message-protect` is implemented for a bounded dev-local one-message proof.
- `message-open` is implemented for a bounded dev-local one-message proof.
- Alice can protect plaintext `"hello bob"`.
- Bob can open Alice’s protected artifact and recover `"hello bob"`.
- The dev-local MLS application-message lifecycle exists:
  - `create -> add-member -> Welcome export -> join -> protect -> open`.
- Both protect and open save provider storage and reload-prove group state.
- Multi-message continuity has been planned and API-reconned, not implemented.

---

## 12. Not Allowed Claims

Not allowed:

- CarbonStack is production secure or Signal-equivalent.
- Local signer storage is a secure vault.
- Local `provider-storage.json` is a secure vault.
- Hardware-backed identity exists.
- The serialized KeyPackage artifact is final CarbonStack onboarding format.
- `welcome.bin` is final CarbonStack onboarding UX.
- `application-message.bin` is final CarbonStack message format.
- Comms CLI consumes sidecar identity/conversation/message state.
- Cypher routes MLS payloads.
- Trust-state storage consumes add-member, join, protect, or open events.
- Production E2EE exists.
- Add-member/join/protect/open means user-facing product UX exists.
- Device-scoped Bob join/open means all state layout has been migrated.
- The one-message proof proves multi-message continuity.
- The one-message proof proves hostile-server security.
- v0.2.38 recon proves out-of-order or replay semantics.

---

## 13. Next TODO

Next safest rung: **v0.2.63 Option C completion / known-good testing cleanup**.

Recommended immediate work:

- run a clean validation pass after the public-surface rewrite:
  - `carbonstack/scripts/validate-local.ps1`;
  - `carbonstack-cypher go test ./... -count=1`;
  - `carbonstack-comms scripts/smoke-openmls-real-cypher-relay.ps1`;
  - `carbonstack-comms scripts/smoke-openmls-real-cypher-relay.ps1 -Full`;
  - `carbonstack-comms scripts/check-no-rust-artifacts.ps1`;
- verify public README/runbook commands still match the current repo state;
- verify that component READMEs point back to the main `carbonstack` front door instead of carrying the whole release narrative;
- preserve the public-surface tone:
  - direct;
  - definition-first;
  - proof-boundary-first;
  - non-hype;
  - non-GPT-ish;
  - explicit about nonclaims;
- keep historical archive docs intact unless a current README directly points to them as authoritative;
- begin Option C completion for testing:
  - repeatable local harness path;
  - known-good validation surface;
  - failure cleanup notes;
  - generated-state/process hygiene;
  - no production or audit claims.

After that:

- v0.2.64: inbox/ack/general semantics and schema standardization;
- v0.2.65+: Option B planning/implementation;
- pre-v0.3.0: release hardening, stale-claim sweep, repo cleanup, security disclaimer hardening;
- v0.3.0: experimental server-deployable CarbonStack backbone epoch.

Do not jump straight to runtime Comms send/inbox UX or Android work before the harness/semantics/release-hardening rungs are complete.

## 14. Historical Continuity Carry-Forward

This section compresses prior continuity payloads. Current-state authority is Sections 1–13 above and the lean JSON breakpoint file.

Phase 1 relay/client skeleton, Phase 2A trust-state scaffold, Phase 2B provider boundary skeleton, and Phase 2C OpenMLS feasibility remain carried forward as previously recorded. Phase 2D now has a progressively validated sidecar ladder:

1. provider-info
2. identity-create
3. identity-status
4. public-bundle-export summary
5. public-bundle-export artifact + private provider state persistence
6. conversation-create
7. provider/group persistence
8. conversation-load-check
9. conversation-add-member / Welcome export
10. conversation-join / Welcome consume
11. message-protect/message-open skeleton and API recon
12. message-protect/message-open implementation
13. provider-info structured JSON cleanup
14. multi-message continuity plan + API recon

The current validated Phase 2D proof is:

```text
create -> add-member -> Welcome export -> join -> protect -> open
```

The next planned Phase 2D proof is:

```text
create -> add-member -> Welcome export -> join -> protect/open #1 -> protect/open #2
```

The next implementation should remain narrow: explicit labels plus two sequential messages, no runtime integration.


---

## 14. v0.2.38 Implementation Checkpoint Addendum

This addendum is authoritative for the v0.2.38 implementation delta over v0.2.37.

### 14.1 Snapshot heads

- `carbonstack`: `01bc4a8 docs: record OpenMLS corrupt message artifact behavior`
- `carbonstack-comms`: `80685a7 test: cover OpenMLS sidecar message ordering replay cases`
- `carbonstack-cypher`: `0bfd5af chore: remove tracked Cypher local state artifacts`
- `carbonstack-os`: `b537475 Add CarbonStackOS north star and initial appliance model`

### 14.2 What changed

v0.2.38 implemented the v0.2.37 recon target:

```text
create -> add-member -> Welcome export -> join -> protect/open message-0001 -> protect/open message-0002
```

The new command behavior is:

```text
message-protect --device-label <sender-device> --conversation-label <conversation> --message-label <message-label> --plaintext <text>

message-open --device-label <receiver-device> --conversation-label <conversation> --message-label <message-label> --message <path>
```

Compatibility behavior remains:

```text
message-protect ... --plaintext "hello bob"
message-open ... --message <message-0001 path>
```

When omitted, `--message-label` defaults to:

```text
message-0001
```

### 14.3 Validated manual proof

The validated manual sequence was:

```text
identity-create Alice
identity-create Bob
public-bundle-export Bob --write-artifact
conversation-create Alice
conversation-add-member Alice + Bob KeyPackage
conversation-join Bob + Alice welcome.bin
message-protect Alice --message-label message-0001 --plaintext "hello bob 1"
message-open Bob --message-label message-0001 --message <message-0001 artifact>
message-protect Alice --message-label message-0002 --plaintext "hello bob 2"
message-open Bob --message-label message-0002 --message <message-0002 artifact>
```

The final corrected `message-open --message-label message-0002` output reported:

```text
message_label = message-0002
message_open_summary_path_hint = ...opened-messages\message-0002\message-open-summary.json
plaintext_utf8 = "hello bob 2"
plaintext_len = 11
message_opened = true
provider_storage_loaded = true
provider_storage_written = true
group_reloadable = true
```

This proves two sequential Alice-to-Bob application messages across separate sidecar invocations.

### 14.4 Implementation files changed

In `carbonstack-comms`:

```text
internal/protocol/mls/research/openmls-sidecar/src/main.rs
internal/protocol/mls/research/openmls-sidecar/src/state.rs
internal/protocol/openmls_sidecar_provider_info_test.go
```

In `carbonstack`:

```text
docs/78-openmls-sidecar-multi-message-continuity-result-v0.md
```

### 14.5 Critical function/API updates

`state.rs` now includes / updates:

```text
validate_message_label(label)
protect_dev_message(device_label, conversation_label, message_label, plaintext)
open_dev_message(device_label, conversation_label, message_label, message_artifact_path)
conversation_message_dir(conversation_label, message_label)
conversation_message_artifact_path(conversation_label, message_label)
conversation_message_manifest_path(conversation_label, message_label)
conversation_message_protect_summary_path(conversation_label, message_label)
device_conversation_message_open_summary_path(device_label, conversation_label, message_label)
```

`main.rs` now includes / updates:

```text
parse_message_label(args)
handle_message_protect(args) -> default/validate/pass message_label
handle_message_open(args) -> default/validate/pass message_label
validate_message_label import from state.rs
```

`openmls_sidecar_provider_info_test.go` now includes / updates:

```text
MessageLabel field in envelope data
TestOpenMLSSidecarMessageProtectOpenTwoSequentialMessages
assertMessageProtectSuccess(...)
assertMessageOpenSuccess(...)
```

### 14.6 Message label validator

The message label validator rejects unsafe labels.

Allowed:

```text
ASCII letters
ASCII numbers
hyphen
underscore
```

Rejected:

```text
empty labels
labels longer than 64 bytes
labels beginning with "."
path separators
unsafe characters
reserved/internal names such as signer, provider-storage, welcome, public-bundle, application-message, con, nul
```

### 14.7 Artifact layout after v0.2.38

Alice protected message artifacts:

```text
.carbonstack-openmls-sidecar-state/dev/conversations/<conversation-label>/messages/<message-label>/application-message.bin
.carbonstack-openmls-sidecar-state/dev/conversations/<conversation-label>/messages/<message-label>/message-manifest.json
.carbonstack-openmls-sidecar-state/dev/conversations/<conversation-label>/messages/<message-label>/message-protect-summary.json
```

Bob open summaries:

```text
.carbonstack-openmls-sidecar-state/dev/devices/<device-label>/conversations/<conversation-label>/opened-messages/<message-label>/message-open-summary.json
```

### 14.8 Go contract validation

New validated test:

```text
TestOpenMLSSidecarMessageProtectOpenTwoSequentialMessages
```

It validates:

- Alice/Bob identity creation.
- Bob public-bundle export with provider storage persistence.
- Alice conversation creation.
- Alice add-member / Welcome export.
- Bob Welcome consume / join.
- `message-0001` protect/open with `"hello bob 1"`.
- `message-0002` protect/open with `"hello bob 2"`.
- message artifact paths differ.
- message artifact hashes differ.
- `provider_storage_loaded=true`.
- `provider_storage_written=true`.
- `group_reloadable=true`.
- `member_count=2`.
- `group_id_ref` remains consistent across add-member/join/protect/open.
- duplicate `message-0002` protect refuses overwrite with `message_artifact_exists`.
- stdout does not include forbidden secret material.

The pre-existing one-message test remains valid because default `message-0001` behavior remains.

### 14.9 Validation commands run for this checkpoint

The v0.2.38 closure used the established validation set:

```text
go test ./internal/protocol -run "TestOpenMLSSidecarProviderInfoCommand|TestOpenMLSSidecarUnsupportedCommandEnvelope|TestOpenMLSSidecarMessageProtectOpenOneWay|TestOpenMLSSidecarMessageProtectOpenTwoSequentialMessages"
go test ./internal/protocol
# temporary local-stability fallback after v0.2.41 BSOD:
go test -p 1 ./internal/protocol
go test ./...
powershell -ExecutionPolicy Bypass -File .\scripts\check-no-rust-artifacts.ps1
powershell -ExecutionPolicy Bypass -File .\scripts\validate-local.ps1
```

### 14.10 Blunders and repairs from v0.2.38

#### Function signature mismatch

After changing `protect_dev_message` to require `message_label`, `main.rs` still called the old three-argument form.

Repair:

```text
parse --message-label
default to message-0001
validate label
pass message_label into protect_dev_message
```

#### Missing import

`main.rs` used `validate_message_label` before importing it.

Repair:

```text
import validate_message_label from state.rs
```

#### Stale open default shadowing

`open_dev_message` accepted `message_label` but then shadowed it internally:

```text
let message_label = "message-0001";
```

This caused `message-open --message-label message-0002` to decrypt `"hello bob 2"` correctly but report/write `message-0001`.

Repair:

```text
remove stale local shadowing line
use the function parameter throughout
```

Lesson:

Correct plaintext alone is not enough; metadata and summary paths must also match the explicit message label.

#### Go helper type typo

New helpers used a nonexistent type:

```text
sidecarEnvelope
```

Actual project type:

```text
openMLSSidecarEnvelope
```

Repair:

```text
assertMessageProtectSuccess(... envelope openMLSSidecarEnvelope ...)
assertMessageOpenSuccess(... envelope openMLSSidecarEnvelope ...)
```

#### Duplicate Go struct fields

The envelope struct already had most v0.2.36 message fields. The patch duplicated them and only truly needed `MessageLabel`.

Repair:

```text
keep first message field block
remove duplicate repeated field lines
```

Lesson:

Inspect the top struct before adding fields. Use line-based cleanup when exact duplicate-block matching fails.

### 14.11 Allowed claims after v0.2.38

Allowed:

- `message-protect` supports explicit safe message labels.
- `message-open` supports explicit safe message labels.
- Omitting `--message-label` preserves `message-0001` compatibility.
- Alice can protect two sequential labeled application messages.
- Bob can open two sequential labeled application messages.
- Bob recovered `"hello bob 1"` and `"hello bob 2"`.
- Duplicate message-label protect refuses overwrite.
- The sidecar validates multi-message continuity for two ordered Alice-to-Bob messages across separate invocations.

Not allowed:

- out-of-order delivery behavior is proven;
- replay/duplicate-open behavior is proven;
- bidirectional messaging is proven;
- generated message IDs exist;
- Alice state has been migrated to device-scoped layout;
- Cypher routes MLS payloads;
- Comms runtime send/inbox uses OpenMLS;
- trust-state consumes protect/open events;
- production E2EE exists.

### 14.12 Next safest action

Recommended v0.2.39 direction:

```text
docs/recon first:
  out-of-order open behavior
  duplicate-open / replay behavior
  skipped-message behavior
  state-layout pressure
```

Suggested proof questions:

- What happens if Bob opens `message-0002` before `message-0001`?
- What happens if Bob opens `message-0001` twice?
- Does OpenMLS reject stale/replayed application messages at this layer?
- Do we need explicit app-level replay metadata before Cypher routing?
- Should Alice global state be migrated to device-scoped before routing?

Do not jump directly to Cypher/Comms runtime integration until ordering/replay behavior and state layout are documented.

---

## 14. v0.2.39 Breakpoint Addendum

This breakpoint supersedes v0.2.38 for the current handoff state.

### Current repo heads

- `carbonstack`: `01bc4a8 docs: record OpenMLS corrupt message artifact behavior`
- `carbonstack-comms`: `80685a7 test: cover OpenMLS sidecar message ordering replay cases`
- `carbonstack-cypher`: `0bfd5af chore: remove tracked Cypher local state artifacts`
- `carbonstack-os`: `b537475 Add CarbonStackOS north star and initial appliance model`

### Validated now

- v0.2.38 explicit message labels and two-message continuity remain validated.
- Ordered batch delivery remains validated.
- Same-sender two-message out-of-order open is validated and regression-tested for the current dev-local flow.
- Duplicate/replay open is rejected with `SecretReuseError` and regression-tested.
- Truncated/corrupt application-message artifact is rejected with `message_artifact_invalid`, `provider.message.invalid`, and `EndOfStream`, and is regression-tested.
- No new Rust sidecar behavior was required for v0.2.39.

### Next safest action

Plan/recon Alice state layout cleanup:

- current Alice state: `dev/conversations/<conversation-label>/`;
- current Bob joined state: `dev/devices/<device-label>/conversations/<conversation-label>/`;
- likely target: `dev/devices/<alice-device>/conversations/<conversation-label>/` for Alice as well.

Do this as docs/recon first. Do not combine state migration with Cypher routing or Comms runtime integration.



---

## 15. v0.2.40 Breakpoint Addendum

This addendum preserves the v0.2.40 state-layout cleanup as a historical checkpoint.

### Current repo heads at v0.2.40

- `carbonstack`: `2d1d657 docs: record OpenMLS Alice device-scoped state result`
- `carbonstack-comms`: `7ccf545 refactor: use device-scoped OpenMLS conversation state`
- `carbonstack-cypher`: `0bfd5af chore: remove tracked Cypher local state artifacts`
- `carbonstack-os`: `b537475 Add CarbonStackOS north star and initial appliance model`

### Validated at v0.2.40

- Alice/creator conversation state moved from `dev/conversations/<conversation-label>/` to `dev/devices/<device-label>/conversations/<conversation-label>/`.
- `conversation-create`, `conversation-load-check`, `conversation-add-member`, and `message-protect` were intended to use device-scoped creator state.
- `conversation-join` and `message-open` were already device-scoped.
- Command surface stayed unchanged; path hints changed.
- No old global-path migration compatibility exists.

### Important correction after v0.2.40

v0.2.41 stale-path sweep discovered that v0.2.40 had missed one write path: `protect_dev_message` still wrote protected message artifacts through old global message helpers. This was fixed at `cf8ac20`.

---

## 16. v0.2.41 Breakpoint Addendum

This breakpoint supersedes v0.2.40 for the current handoff state.

### Current repo heads

- `carbonstack`: `a062dc5 docs: record OpenMLS sidecar Go test split`
- `carbonstack-comms`: `6a3a24b test: split OpenMLS sidecar message tests`
- `carbonstack-cypher`: `0bfd5af chore: remove tracked Cypher local state artifacts`
- `carbonstack-os`: `b537475 Add CarbonStackOS north star and initial appliance model`

### Validated now

- Phase 2D OpenMLS sidecar mainline research is closed enough to proceed to Cypher relay research.
- Device-scoped Alice/Bob conversation state is validated.
- `protect_dev_message` writes sender message artifacts under the sender device-scoped conversation path.
- Wrong-device message-open is regression-tested:
  - missing receiver device/conversation provider storage;
  - `conversation_or_message_missing`;
  - `provider.conversation.missing`;
  - warning severity;
  - trust relevant false;
  - exit code `3`.
- Wrong-conversation message-open is regression-tested with the same safe missing-provider-storage shape.
- Bidirectional Alice↔Bob message flow is regression-tested:
  - Alice protects and Bob opens `alice-message-0001`;
  - Bob protects and Alice opens `bob-message-0001`;
  - expected plaintexts are `"hello bob from alice"` and `"hello alice from bob"`.
- The large OpenMLS sidecar Go test file now includes:
  - `TestOpenMLSSidecarMessageOpenWrongDeviceRejected`;
  - `TestOpenMLSSidecarMessageOpenWrongConversationRejected`;
  - `TestOpenMLSSidecarMessageProtectOpenBidirectional`.

### Local BSOD case-study note

During v0.2.41 validation, the local Windows machine BSODed during heavy build/test activity. The minidump analysis indicated a local `WIN32K_CRITICAL_FAILURE (0x164)` involving `dwm.exe` / `win32kbase.sys`; Vanguard and Avast kernel drivers were loaded, and Reliability Monitor showed repeated Vanguard user-mode crashes nearby. The project state was protected by git commits, and this is recorded only as local environment instability. Use `go test -p 1` for broad validation on this machine until local OS/security-driver stability is addressed.

### Allowed claims at v0.2.41

Allowed:

- Phase 2D mainline OpenMLS sidecar research is complete enough to support Cypher minimal opaque MLS artifact relay research.
- The dev-local sidecar validates full Alice/Bob two-device lifecycle through bidirectional message flow.
- Wrong-target receiver metadata fails safely before message processing when provider storage is missing.
- Duplicate/replay open and corrupt/truncated artifact behavior are contract-tested.
- The sidecar remains research/dev-only and is not a production secure-storage or E2EE release.

Not allowed:

- Cypher routes MLS artifacts;
- Comms runtime send/inbox uses OpenMLS;
- trust-state consumes sidecar events;
- production E2EE exists;
- Android/Pixel 4a testing has begun;
- sidecar code has been promoted out of research paths;
- generated message IDs exist;
- multi-sender, long skipped-message, or membership-change matrices are complete.

### Next safest action

Start Cypher minimal opaque MLS artifact relay research:

- route Welcome artifacts;
- route application-message artifacts;
- define delivery metadata;
- define hashes/sizes and sender/recipient/conversation identifiers;
- preserve server blindness to plaintext, provider storage, signer material, group secrets, raw MemoryStorage, and trust-state private material.

After Cypher research agrees with Comms/OpenMLS research, plan `Phase 2E: Research-to-Implementation Promotion` to move researched systems out of research directories into official implementation structures. This is not a user release.
