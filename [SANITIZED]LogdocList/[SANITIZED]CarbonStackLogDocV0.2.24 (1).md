# CarbonStack LogDoc v0.2.24PRIME

**Last updated:** 2026-05-23 18:40 -04:00  
**Owner/operator:** bitcrusher32  
**Assistant/model:** GPT-5.5 Thinking  
**Current phase:** Phase 2D OpenMLS sidecar public-bundle-export summary checkpoint: `docs/53-openmls-sidecar-public-bundle-export-plan-v0.md` and `docs/54-openmls-sidecar-public-bundle-export-result-v0.md` have landed; the OpenMLS sidecar can now create dev-only identity state, load/check that identity state with `identity-status`, and run `public-bundle-export --device-label <safe>` to load existing dev identity state, generate a real OpenMLS KeyPackage in memory, compute a sanitized `sha256:<hex>` KeyPackage reference, write `public-bundle-summary.json`, return sanitized JSON with `public_bundle_exported: true`, `public_bundle_available: true`, `key_package_created: true`, `key_package_artifact_written: false`, `provider_storage_written: false`, and `private_material_included: false`, and emit `provider.public_bundle.exported`. This is still summary-only public bundle export. It does not write a full serialized KeyPackage artifact, create conversations, add members, join groups, protect/open messages, wire OpenMLS into CarbonStackComms runtime, route through CarbonStackCypher, mutate trust-state storage, or claim production E2EE.  
**Version schema:** v[scope].[timeline] — this file is `v0.2.24PRIME`, the Phase 2D public-bundle-export summary checkpoint after dev-only identity creation and identity-status were followed by a narrow KeyPackage-generation/public-bundle-summary rung.



**PRIME merge note:** This LogDoc is rebased from the lean `v0.2.24` checkpoint and carries forward the older bloated `v0.2.21` operational memory for full continuity. The v0.2.24 sections are authoritative for the current state. The v0.2.21 material preserved later in this file is historical/carry-forward context; where it conflicts with v0.2.22–v0.2.24, treat the newer v0.2.24 checkpoint as current and the older text as provenance.

---

## 1. Project Goal

**Active goal:** Preserve CarbonStack state at the Phase 2D OpenMLS sidecar public-bundle-export summary checkpoint. Phase 2C remains closed for mainline work. Phase 2D now has a Rust↔Go sidecar boundary, common JSON envelopes, stabilized command-surface events, recognized `identity-create`, dev-only secret-bearing identity creation, test-protected `identity-status`, and now a summary-only `public-bundle-export` command that generates a real OpenMLS KeyPackage in memory and writes a sanitized public bundle summary without printing secret material.

The sidecar can now:

- report provider metadata with `provider-info`;
- create dev-only identity state with `identity-create --device-label <safe>`;
- write secret-bearing `signer.json` under ignored local dev state;
- write sanitized `identity-prep.json`, `identity-summary.json`, and `identity-state.json`;
- refuse duplicate identity creation with `identity_already_exists`;
- load existing identity state with `identity-status --device-label <safe>`;
- deserialize signer state without printing it;
- recompute and verify the public identity reference against the sanitized identity summary;
- generate a real OpenMLS KeyPackage in memory with `public-bundle-export --device-label <safe>`;
- compute a sanitized KeyPackage reference/hash;
- write `public-bundle-summary.json`;
- report typed provider events for identity created, identity loaded, identity missing, identity exists, and public bundle exported.

This is still dev-only and not production secure storage. It does not write a full serialized KeyPackage artifact. It does not create conversations, add members, join groups, protect/open messages, write OpenMLS provider storage, call Comms runtime flows, route through Cypher, or mutate trust-state storage.

**Component model:**

- **CarbonStackOS** — deliberately constrained Android-derived appliance OS. Deferred.
- **CarbonStackComms** — text-first messaging client; currently CLI-first with stub crypto, Phase 2A trust behavior, Phase 2B provider-boundary skeleton, Phase 2C OpenMLS feasibility/fixture work, and Phase 2D OpenMLS sidecar prototype. The sidecar now supports dev-only identity creation, identity status/load checks, and summary-only public bundle export, but it is not wired into user-facing messaging.
- **CarbonStackCypher** — hostile-server opaque-envelope relay; no MLS payload routing yet.
- **carbonstack** — canonical doctrine/specification repo, protocol docs, validation matrix, and local validation runner owner.

**North star:**

> Its goal is not to be a better smartphone. Its goal is to stop being a smartphone.

**Core doctrine:**

> Every feature is guilty until it proves it does not add unacceptable parser, network, sensor, identity, or filesystem authority.

---

## 2. Explicit Non-Goals / Out of Scope

- Do not start CarbonStackOS build work.
- Do not start Android implementation.
- Do not implement custom cryptography.
- Do not wire OpenMLS into `comms send`, `comms inbox`, Cypher routing, `trust.json`, or `trust-events.jsonl` yet.
- Do not claim production security, Signal-equivalent security, hostile-server proof, replay resistance, metadata privacy, or production E2EE.
- Do not treat dev-only `signer.json` as a secure vault.
- Do not print, paste, expose, or commit `signer.json`, MemoryStorage JSON, provider storage JSON, private keys, seeds, recovery material, or raw key bytes.
- Do not treat `public-bundle-summary.json` as final onboarding material.
- Do not claim full serialized KeyPackage artifact export exists.
- Do not claim another client can consume the public bundle yet.
- Do not implement conversation lifecycle, add-member, join, or message protect/open without another plan/checkpoint.
- Do not add groups as a user-facing feature yet, even though conversations are conceptually group-shaped.
- Do not permanently couple CarbonStackComms to AGPL/libsignal unless a future decision explicitly accepts that tradeoff.

---

## 3. Current State

| Classification | Items |
|---|---|
| **VALIDATED** | Phase 1 relay/client skeleton works locally. Cypher API tests pass. Comms package tests pass. Comms lifecycle smoke test passes. Phase 2A trust lifecycle passes. Provider skeleton package tests pass. Phase 2C OpenMLS scratch/fixture/trust mapping path is validated. Phase 2D sidecar provider-info/envelope/identity-create/identity-status/public-bundle-export summary paths are Go-tested. |
| **VALIDATED DOC DECISION** | MLS/OpenMLS planning docs exist through `docs/54-openmls-sidecar-public-bundle-export-result-v0.md`. `docs/53` planned public-bundle-export; `docs/54` records the implemented summary-only result. |
| **VALIDATED SIDECAR IDENTITY CREATION** | `identity-create --device-label <safe>` validates label, refuses overwrite, creates ignored per-device dev state, creates `SignatureKeyPair`, creates `BasicCredential`, creates `CredentialWithKey`, writes secret-bearing `signer.json`, writes sanitized identity JSON files, returns sanitized success JSON, emits `provider.identity.created`, and keeps `private_material_included: false`. |
| **VALIDATED SIDECAR IDENTITY STATUS** | `identity-status --device-label <safe>` validates label, requires existing identity state, loads/deserializes `signer.json`, reads sanitized summary/state files, derives public signer key, recomputes `sha256:<hex>` public identity reference, compares it to `identity-summary.json`, returns sanitized success JSON, emits `provider.identity.loaded`, and keeps `private_material_included: false`. |
| **VALIDATED PUBLIC BUNDLE SUMMARY EXPORT** | `public-bundle-export --device-label <safe>` validates label, requires existing identity state, loads identity safely, generates a real OpenMLS KeyPackage in memory, extracts public KeyPackage, computes a KeyPackage reference/hash, writes sanitized `public-bundle-summary.json`, returns sanitized success JSON, emits `provider.public_bundle.exported`, and keeps `private_material_included: false`. It does not write a full serialized KeyPackage artifact. |
| **VALIDATED GO-SIDE CONTRACT** | Go tests parse sidecar envelopes; assert identity created/loaded/missing events; assert duplicate refusal; assert public-bundle-export capability; assert `provider.public_bundle.exported`; assert no obvious secret material appears in stdout; assert `signer.json` exists but is not printed; assert identity-status returns the same public identity ref and key length as identity-create; assert public-bundle-export returns the same public identity ref, a `sha256:` KeyPackage reference, hash length 32, and no full KeyPackage artifact. |
| **PARTIAL** | Sidecar identity creation/status/public-bundle summary exists only as dev-local state. It is not production secure storage, not a vault, not hardware-backed, not final onboarding, and not connected to messaging. |
| **NOT VALIDATED** | Full serialized KeyPackage artifact export, consumable onboarding artifact, provider storage persistence, conversation create/add/join, message protect/open, Comms CLI integration, Cypher MLS routing, trust-state mutation, production vault, hostile-server harness, replay resistance, metadata privacy, Android, CarbonStackOS, CI. |
| **BLOCKED / NEXT** | The next safest step is either a doc-first full serialized KeyPackage artifact export plan or a careful conversation-create/add-member planning step after deciding whether full artifact export is required. Do not jump straight to message protect/open or runtime integration. |

---

## 4. Critical Paths

### Local repo paths

- `C:\▮▮`
- `C:\▮▮\carbonstack`
- `C:\▮▮\carbonstack-comms`
- `C:\▮▮\carbonstack-cypher`
- `C:\▮▮\carbonstack-os`

### Current repo heads at v0.2.24

- `carbonstack`: `e620e38 docs: record OpenMLS sidecar public-bundle-export result`
- `carbonstack-comms`: `f55b647 test: add OpenMLS sidecar public bundle summary export`
- `carbonstack-cypher`: `0bfd5af chore: remove tracked Cypher local state artifacts`
- `carbonstack-os`: `b537475 Add CarbonStackOS north star and initial appliance model`

### Canonical docs and runner

- `carbonstack/docs/36-provider-event-taxonomy-v0.md`
- `carbonstack/docs/37-provider-negative-fixture-result-v0.md`
- `carbonstack/docs/38-provider-trust-state-mapping-v0.md`
- `carbonstack/docs/39-phase2d-sidecar-command-surface-plan.md`
- `carbonstack/docs/40-openmls-sidecar-provider-info-result-v0.md`
- `carbonstack/docs/41-openmls-sidecar-json-envelope-plan-v0.md`
- `carbonstack/docs/42-openmls-sidecar-json-envelope-result-v0.md`
- `carbonstack/docs/43-provider-command-unsupported-event-result-v0.md`
- `carbonstack/docs/44-openmls-sidecar-identity-create-plan-v0.md`
- `carbonstack/docs/45-openmls-sidecar-identity-create-prep-result-v0.md`
- `carbonstack/docs/46-provider-command-invalid-not-implemented-events-result-v0.md`
- `carbonstack/docs/47-openmls-sidecar-identity-prep-state-result-v0.md`
- `carbonstack/docs/48-provider-identity-prep-state-events-result-v0.md`
- `carbonstack/docs/49-openmls-sidecar-real-identity-create-plan-v0.md`
- `carbonstack/docs/50-openmls-sidecar-real-identity-create-result-v0.md`
- `carbonstack/docs/51-openmls-sidecar-identity-status-plan-v0.md`
- `carbonstack/docs/52-openmls-sidecar-identity-status-result-v0.md`
- `carbonstack/docs/53-openmls-sidecar-public-bundle-export-plan-v0.md`
- `carbonstack/docs/54-openmls-sidecar-public-bundle-export-result-v0.md`
- `carbonstack/scripts/validate-local.ps1`
- `carbonstack/scripts/validate-phase1.ps1` compatibility wrapper

### CarbonStackComms implementation/test paths

- `carbonstack-comms/internal/protocol/provider_events.go`
- `carbonstack-comms/internal/protocol/provider_events_test.go`
- `carbonstack-comms/internal/protocol/provider_trust.go`
- `carbonstack-comms/internal/protocol/provider_trust_test.go`
- `carbonstack-comms/internal/protocol/openmls_sidecar_provider_info_test.go`
- `carbonstack-comms/internal/protocol/mls/research/openmls-sidecar/Cargo.toml`
- `carbonstack-comms/internal/protocol/mls/research/openmls-sidecar/Cargo.lock`
- `carbonstack-comms/internal/protocol/mls/research/openmls-sidecar/src/main.rs`
- `carbonstack-comms/internal/protocol/mls/research/openmls-sidecar/src/labels.rs`
- `carbonstack-comms/internal/protocol/mls/research/openmls-sidecar/src/state.rs`
- `carbonstack-comms/scripts/check-no-rust-artifacts.ps1`

### Sidecar dev-state path

Ignored local dev identity/public-bundle state lives under:

- `carbonstack-comms/internal/protocol/mls/research/openmls-sidecar/.carbonstack-openmls-sidecar-state/dev/devices/<device-label>/`

Expected files for `carbonstack-alice-device` after successful dev identity creation and public-bundle summary export:

- `identity-prep.json` — sanitized/non-secret.
- `identity-summary.json` — sanitized/non-secret.
- `identity-state.json` — sanitized/non-secret.
- `public-bundle-summary.json` — sanitized/dev public-bundle summary; not final onboarding material.
- `signer.json` — secret-bearing dev-only signer material; do not print, paste, or commit.

---

## 5. Phase Timeline Summary

### Phase 2C closure carried forward

Phase 2C remains closed for mainline work. It validated OpenMLS feasibility through scratch probes, same-process reload, MemoryStorage file persistence, sanitized fixtures, Go-side fixture parsing, event taxonomy, negative fixture mapping, provider trust-decision mapping, and fixture-backed trust-decision tests. Reopen only for targeted fixes.

### v0.2.21 identity prep/checkpoint events

The v0.2.21 checkpoint stabilized identity prep/checkpoint events:

- `provider.identity.prep_state_written`
- `provider.identity.exists`
- `checkpoint.failed`

These were represented in Go provider-event taxonomy/trust mapping, sidecar tests compared events against typed constants, and docs/48 recorded the result.

### v0.2.22 real dev-only identity creation

The v0.2.22 checkpoint planned and implemented real dev-only sidecar identity creation.

New canonical docs:

- `docs/49-openmls-sidecar-real-identity-create-plan-v0.md`
- `docs/50-openmls-sidecar-real-identity-create-result-v0.md`

New implementation:

- Added OpenMLS identity-material dependencies to `openmls-sidecar`.
- Added/used `openmls`, `openmls_basic_credential`, `openmls_rust_crypto`, `serde`, `serde_json`, `sha2`, and `hex`.
- Implemented `create_dev_identity(...)` in sidecar state logic.
- Implemented sanitized `identity-create` success envelope using `serde_json::json!`.
- Introduced `provider.identity.created`.

Validated:

- `identity-create` creates `signer.json`; user confirmed it exists and did not paste it.
- `identity-create` writes sanitized identity files.
- `identity-create` refuses duplicate overwrite.
- Go tests validate no-secret stdout and no KeyPackage/provider-storage/public-bundle claims.

### v0.2.23 identity-status / load-check

The v0.2.23 checkpoint planned and implemented identity-status/load-check before public-bundle export.

New canonical docs:

- `docs/51-openmls-sidecar-identity-status-plan-v0.md`
- `docs/52-openmls-sidecar-identity-status-result-v0.md`

New implementation:

- Added `identity-status --device-label <safe>` sidecar command.
- Added state helper(s) to load/dev-check existing identity state.
- Deserializes `signer.json` without printing it.
- Reads sanitized `identity-summary.json` and `identity-state.json`.
- Derives signer public key with `to_public_vec()`.
- Recomputes `sha256:<hex>` public identity reference.
- Confirms recomputed public identity ref matches `identity-summary.json`.
- Emits sanitized success envelope with `provider.identity.loaded`.
- Emits missing-state envelope with `identity_missing` / `provider.identity.missing`.
- Reuses `provider.secret.material.unavailable` for unloadable/invalid secret material.
- Adds owned read structs for deserializing summary/state JSON to avoid Serde lifetime issues from borrowed writer structs.

Validated:

- `cargo check` passed after import/lifetime fixes.
- `cargo test` passed.
- `identity-status` before create returns missing identity behavior.
- `identity-create` then `identity-status` works.
- `identity-status` returns same public identity ref/key length as identity-create.
- `signer.json` exists but was not included in pasted output.
- `go test ./internal/protocol` passed.
- `go test ./...` passed.
- `check-no-rust-artifacts.ps1` and `validate-local.ps1` were run before snapshot.
- `docs/51` and `docs/52` landed in `carbonstack`.

Important implementation lessons:

- Writer structs with borrowed `&str` fields are awkward for generic Serde deserialization. Use separate owned read structs (`String` fields) for loaded JSON.
- Trust mapping insertion should be checked with `Select-String`; `provider.identity.missing` initially fell through to default `append_history/debug_only` until an explicit case was inserted.
- `identity-status` must load existing state only. It must not create, repair, overwrite, or export.

### New at v0.2.24

The post-v0.2.23 block planned and implemented summary-only public-bundle-export.

New canonical docs:

- `docs/53-openmls-sidecar-public-bundle-export-plan-v0.md`
- `docs/54-openmls-sidecar-public-bundle-export-result-v0.md`

Current heads:

- `carbonstack` `e620e38 docs: record OpenMLS sidecar public-bundle-export result`
- `carbonstack-comms` `f55b647 test: add OpenMLS sidecar public bundle summary export`

New implementation:

- Added `public-bundle-export --device-label <safe>` sidecar command.
- Updated provider-info capability set so `public-bundle-export` is a capability rather than unsupported.
- Moved `public-bundle-export` out of unsupported command tests; unsupported tests now use still-unsupported commands such as `conversation-create`.
- Added sidecar state helper(s) to export a dev-only public-bundle summary.
- Reuses identity-status style identity loading rather than creating identity state.
- Deserializes `signer.json` locally without printing it.
- Recreates `BasicCredential` and `CredentialWithKey`.
- Uses `openmls_rust_crypto::OpenMlsRustCrypto` as the provider for KeyPackage generation.
- Builds a real OpenMLS `KeyPackageBundle` in memory using `KeyPackage::builder().build(...)`.
- Extracts the public `KeyPackage`.
- Computes a KeyPackage reference/hash using `key_package.hash_ref(provider.crypto())`.
- Writes `public-bundle-summary.json`.
- Returns sanitized JSON with `provider.public_bundle.exported`.
- Adds/uses `provider.public_bundle.exported` in Go event taxonomy and trust mapping.

Validated by user report and snapshot:

- `cargo check` passed after fixing provider type and hash-ref reference usage.
- `cargo test` passed.
- Manual Rust probes confirmed `public-bundle-export` behavior before and after identity creation.
- Safe file inspection confirmed `public-bundle-summary.json` exists and `signer.json` exists but was not pasted.
- `public-bundle-summary.json` included `key_package_created: true`, `key_package_ref: sha256:<hex>`, `key_package_hash_len: 32`, `key_package_artifact_written: false`, `public_bundle_available: true`, `provider_storage_written: false`, and `private_material_included: false`.
- `go test ./internal/protocol` passed after provider-info/test contract drift was fixed.
- `go test ./...` passed by user report.
- `check-no-rust-artifacts.ps1` and `validate-local.ps1` were run before snapshot.
- `docs/53` and `docs/54` landed in `carbonstack`.

Important implementation lessons:

- `KeyPackage::builder().build(...)` requires an `OpenMlsProvider`; `openmls_rust_crypto::OpenMlsRustCrypto` is the correct provider type in this dependency set.
- `RustCrypto` directly is only the crypto backend and does not implement the provider trait expected by KeyPackage builder.
- `provider.crypto()` already returns the correct crypto reference for `KeyPackage::hash_ref(...)`; using `&provider.crypto()` creates a reference-to-reference and fails.
- When moving a command from unsupported to supported, update provider-info capabilities, unsupported lists, unsupported-command tests, and fixed array/slice definitions together.
- Prefer `const UNSUPPORTED_COMMANDS: &[&str] = &[...]` so unsupported list length does not become stale when commands graduate.
- Broad regex/string patches can corrupt Go tests; reset only the broken file if needed and reapply narrow `@' ... '@` replacements.
- Summary-only public bundle export is a meaningful rung, but it is not final onboarding material.

---

## 6. Event and Trust State at v0.2.24

### Stabilized operational events

- `provider.command.unsupported`
- `provider.command.invalid`
- `provider.command.not_implemented`
- `provider.identity.prep_state_written`
- `provider.identity.created`
- `provider.identity.loaded`
- `provider.identity.missing`
- `provider.identity.exists`
- `provider.public_bundle.exported`
- `checkpoint.failed`
- `provider.secret.material.unavailable` carried forward for serious/unloadable secret material failures

### `provider.identity.created`

Meaning:

- Sidecar generated local dev identity material.
- Required local state files were written.
- stdout remained sanitized.
- No public bundle was exported yet at identity-create time.
- No provider storage was written.
- No conversation/message state exists.

Current classification:

- class: storage/checkpoint
- severity: info/notice in envelope context, non-trust-relevant
- trust relevant: false

Current trust-decision behavior:

- append history
- debug only
- do not block send/receive/open
- do not require reverify
- not user visible

### `provider.identity.loaded`

Meaning:

- Existing sidecar identity state was found.
- Secret-bearing signer material was loaded/deserialized locally without printing.
- Public identity reference was recomputed and matched the sanitized summary.
- No new identity material was generated.
- No public bundle was exported by identity-status.

Current classification:

- class: storage/checkpoint
- severity: info
- trust relevant: false

Current trust-decision behavior:

- append history
- debug only
- do not block send/receive/open
- do not require reverify
- not user visible

### `provider.identity.missing`

Meaning:

- An identity-dependent operation expected local identity state.
- Required identity files were missing.

Current classification:

- class: storage/checkpoint
- severity: warning
- trust relevant: false

Current trust-decision behavior:

- stop operation
- show recovery path
- append history
- block current identity-dependent outgoing/send operation
- do not require reverify by default
- not automatically cryptographic trust relevant

### `provider.public_bundle.exported`

Meaning:

- Existing identity state was loaded.
- A real OpenMLS KeyPackage was generated in memory.
- A sanitized public bundle summary was written.
- A KeyPackage reference/hash was returned.
- stdout remained sanitized.
- No private material was printed.
- No full serialized KeyPackage artifact was written.
- No conversation was created.
- No message was protected/opened.

Current classification:

- class: public/setup
- severity: info
- trust relevant: false

Current trust-decision behavior:

- append history
- debug only
- do not block send/receive/open
- do not require reverify
- not user visible

### Duplicate identity refusal

Duplicate identity creation uses:

- error code: `identity_already_exists`
- provider event: `provider.identity.exists`
- exit code: 3
- no overwrite
- no private material printed
- `state_written: false`

This remains critical because silent identity replacement would become a severe future trust hazard.

---

## 7. Known Good Commands

From `carbonstack`:

- `cd C:\▮▮\carbonstack; powershell -ExecutionPolicy Bypass -File .\scripts\validate-local.ps1`

From `carbonstack-comms`:

- `cd C:\▮▮\carbonstack-comms; powershell -ExecutionPolicy Bypass -File .\scripts\check-no-rust-artifacts.ps1`
- `cd C:\▮▮\carbonstack-comms; go test ./internal/protocol`
- `cd C:\▮▮\carbonstack-comms; go test ./...`
- `cd C:\▮▮\carbonstack-comms; go test ./internal/protocol -run TestOpenMLSSidecarIdentityCreateWritesDevIdentityState`
- `cd C:\▮▮\carbonstack-comms; go test ./internal/protocol -run TestOpenMLSSidecarIdentityCreateRefusesOverwrite`
- `cd C:\▮▮\carbonstack-comms; go test ./internal/protocol -run TestOpenMLSSidecarIdentityStatusMissing`
- `cd C:\▮▮\carbonstack-comms; go test ./internal/protocol -run TestOpenMLSSidecarIdentityStatusLoadsExistingIdentity`
- `cd C:\▮▮\carbonstack-comms; go test ./internal/protocol -run TestOpenMLSSidecarPublicBundleExportMissingIdentity`
- `cd C:\▮▮\carbonstack-comms; go test ./internal/protocol -run TestOpenMLSSidecarPublicBundleExportCreatesSummary`
- `cd C:\▮▮\carbonstack-comms; go test ./internal/protocol -run TestDescribeProviderEventIdentityLoaded`
- `cd C:\▮▮\carbonstack-comms; go test ./internal/protocol -run TestDescribeProviderEventIdentityMissing`
- `cd C:\▮▮\carbonstack-comms; go test ./internal/protocol -run TestDescribeProviderEventPublicBundleExported`
- `cd C:\▮▮\carbonstack-comms; go test ./internal/protocol -run TestDecideProviderTrustIdentityLoaded`
- `cd C:\▮▮\carbonstack-comms; go test ./internal/protocol -run TestDecideProviderTrustIdentityMissing`
- `cd C:\▮▮\carbonstack-comms; go test ./internal/protocol -run TestDecideProviderTrustPublicBundleExported`

From `openmls-sidecar`:

- `cd C:\▮▮\carbonstack-comms\internal\protocol\mls\research\openmls-sidecar; cargo check`
- `cd C:\▮▮\carbonstack-comms\internal\protocol\mls\research\openmls-sidecar; cargo test`
- `cd C:\▮▮\carbonstack-comms\internal\protocol\mls\research\openmls-sidecar; cargo run -- provider-info`
- `cd C:\▮▮\carbonstack-comms\internal\protocol\mls\research\openmls-sidecar; cargo run -- identity-create`
- `cd C:\▮▮\carbonstack-comms\internal\protocol\mls\research\openmls-sidecar; cargo run -- identity-create --device-label "../bad"`
- `cd C:\▮▮\carbonstack-comms\internal\protocol\mls\research\openmls-sidecar; Remove-Item -Recurse -Force .\.carbonstack-openmls-sidecar-state -ErrorAction SilentlyContinue; cargo run -- identity-status --device-label carbonstack-alice-device`
- `cd C:\▮▮\carbonstack-comms\internal\protocol\mls\research\openmls-sidecar; cargo run -- identity-create --device-label carbonstack-alice-device`
- `cd C:\▮▮\carbonstack-comms\internal\protocol\mls\research\openmls-sidecar; cargo run -- identity-status --device-label carbonstack-alice-device`
- `cd C:\▮▮\carbonstack-comms\internal\protocol\mls\research\openmls-sidecar; cargo run -- public-bundle-export --device-label carbonstack-alice-device`
- `cd C:\▮▮\carbonstack-comms\internal\protocol\mls\research\openmls-sidecar; cargo run -- public-bundle-export`
- `cd C:\▮▮\carbonstack-comms\internal\protocol\mls\research\openmls-sidecar; cargo run -- public-bundle-export --device-label "../bad"`

Safe inspection command:

- `cd C:\▮▮\carbonstack-comms\internal\protocol\mls\research\openmls-sidecar; Get-ChildItem .\.carbonstack-openmls-sidecar-state\dev\devices\carbonstack-alice-device | Select-Object Name, Length`
- `cd C:\▮▮\carbonstack-comms\internal\protocol\mls\research\openmls-sidecar; Get-Content .\.carbonstack-openmls-sidecar-state\dev\devices\carbonstack-alice-device\public-bundle-summary.json`

Do not inspect or paste `signer.json`.

---

## 8. Not Validated

- Full serialized KeyPackage artifact export.
- Final consumable public onboarding bundle.
- Whether future add-member/join can consume the current summary-only bundle.
- Provider storage persistence.
- Conversation creation/add/join.
- Message protect/open.
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

## 9. Hazards and Lessons

- `signer.json` is real secret-bearing dev state. Never print, paste, commit, or expose it.
- `.carbonstack-openmls-sidecar-state/` must remain ignored.
- `public-bundle-summary.json` is generated dev state. It is sanitized but still should not be committed as normal source.
- `identity-status` and `public-bundle-export` load signer material locally; stdout must remain sanitized.
- `identity-status` must not create, repair, overwrite, or export identity state.
- `public-bundle-export` must not create missing identity state.
- `public-bundle-export` currently writes summary only; it does not write full serialized KeyPackage artifact.
- `identity-create` must refuse overwrite by default.
- Keep public identity refs and KeyPackage refs as dev-only implementation references, not final UX/security claims.
- Separate borrowed writer structs from owned read structs in Rust/Serde.
- Use `OpenMlsRustCrypto` as the provider for KeyPackage building in this dependency set.
- Use `provider.crypto()` directly for `KeyPackage::hash_ref(...)`.
- Trust mapping cases should be inspected after insertion; missing cases can fall through to default non-blocking behavior.
- Provider-info/test contracts must be updated when command capability changes.
- Broad regex/string patches can corrupt files; when a single file becomes mangled, reset only that file and reapply narrow patches.
- Run Rust artifact guard before Rust-related commits.
- Keep `target/`, `.exe`, `.pdb`, signer JSON, provider state, MemoryStorage files, and generated sidecar state out of Git.

---

## 10. Allowed Claims

Allowed:

- Phase 2D sidecar bootstrap exists.
- OpenMLS sidecar provider-info command builds/runs and emits sanitized JSON.
- Go-side tests invoke and parse OpenMLS sidecar envelopes.
- `identity-create` generates dev-only OpenMLS identity/signing material.
- `identity-create` writes secret-bearing `signer.json` under ignored local dev state.
- `identity-create` writes sanitized summary/state/prep JSON.
- `identity-create` returns sanitized success JSON with `private_material_included: false`.
- `identity-create` refuses overwrite with `identity_already_exists`.
- `identity-status` loads existing dev-only sidecar identity state.
- `identity-status` deserializes `signer.json` without printing it.
- `identity-status` recomputes the public identity reference from signer public key and checks it against the summary.
- `public-bundle-export` loads existing dev-only sidecar identity state.
- `public-bundle-export` generates a real OpenMLS KeyPackage in memory.
- `public-bundle-export` computes and returns a sanitized KeyPackage reference/hash.
- `public-bundle-export` writes sanitized `public-bundle-summary.json`.
- `public-bundle-export` emits `provider.public_bundle.exported`.
- `provider.identity.created`, `provider.identity.loaded`, `provider.identity.missing`, and `provider.public_bundle.exported` are typed/classified/trust-mapped.
- Go tests validate identity creation, identity status success/missing behavior, public-bundle-export success/missing behavior, duplicate refusal, no-secret stdout, and no provider-storage/full-KeyPackage-artifact claims.
- No Comms/Cypher/trust runtime integration exists yet.

---

## 11. Not Allowed Claims

Not allowed:

- CarbonStack is production secure or Signal-equivalent.
- Local signer storage is a secure vault.
- Hardware-backed identity exists.
- OpenMLS is integrated into CarbonStackComms messaging.
- Cypher routes MLS payloads.
- Trust-state storage consumes sidecar provider events.
- Full serialized KeyPackage artifact export exists.
- Public bundle output is final onboarding material.
- Another client can consume the current public-bundle summary.
- Conversation creation/add/join exists.
- Message protect/open exists.
- Provider storage persistence is solved.
- Production storage, hostile-server security, replay resistance, metadata privacy, hardware-key flows, Android, or CarbonStackOS are implemented.

---

## 12. Next TODO

Recommended next checkpoint: **full KeyPackage artifact export planning**, unless a careful repo inspection proves conversation planning can proceed with summary-only bundle.

Suggested next doc/code flow:

1. Add a doc-first plan for serialized public KeyPackage artifact export.
2. Inspect the exact OpenMLS/tls_codec serialization path for `KeyPackage`.
3. Decide artifact format:
   - binary `.bin`
   - base64 JSON field
   - both
   - path-only stdout with file written under ignored dev state
4. Decide whether stdout may include serialized public artifact bytes, or only a path/hash.
5. Decide whether the artifact is one-time, dev-reusable, or regenerated per command.
6. Decide duplicate/regeneration behavior.
7. Decide whether provider storage is required or remains false.
8. Add tests that:
   - generate identity
   - run identity-status
   - export full public artifact
   - parse or at least verify artifact existence/hash
   - assert signer/private material is not printed
   - assert generated artifact is not confused with secret state
9. Record result doc.
10. Only then plan conversation-create/add-member/join.

Alternative next checkpoint if artifact export is deferred:

1. Add a conversation lifecycle planning doc.
2. Explicitly state that current public-bundle summary is not consumable onboarding.
3. Identify exact blocker(s) that require serialized KeyPackage artifact.
4. Do not implement conversation lifecycle until the blocker is resolved.

Do not jump directly to message protect/open, Comms runtime integration, Cypher routing, trust-store mutation, Android, or CarbonStackOS.


---

## 13. PRIME Rebase Continuity Index

This section explains how to read the merged artifact.

- **Authoritative current checkpoint:** Sections 1–12 above, derived from `CarbonStack LogDoc v0.2.24`.
- **Historical continuity payload:** Section 14 below, derived from `CarbonStack LogDoc v0.2.21`.
- **Conflict rule:** v0.2.24 supersedes v0.2.21 for current status, commands, event vocabulary, repo heads, and next TODOs.
- **No work-done context loss rule:** Older milestones, timeline entries, blunders, validation ladder, dangerous commands, and breakpoint notes are preserved below even when superseded.
- **Safety rule:** Mentions of `signer.json`, MemoryStorage JSON, provider storage JSON, private keys, and recovery material are path/risk references only. Do not print, paste, commit, or treat any secret-bearing artifact as public.

### Continuity compression map

| v0.2.21 area | PRIME handling |
|---|---|
| Early canonical docs `docs/16`–`docs/35` | Preserved in historical snapshot; current lean section keeps newer `docs/36`–`docs/54` focus. |
| OpenMLS minimal scratch research path | Preserved in historical snapshot; current state treats Phase 2C as closed. |
| Entries 1–72 active iteration notes | Preserved in historical snapshot for step-by-step provenance. |
| Blunders/lessons through v0.2.21 | Preserved in historical snapshot; v0.2.24 hazards add newer sidecar lessons. |
| v0.2.21 TODOs and breakpoint notes | Preserved as historical; v0.2.24 Section 12 is current next-action authority. |
| v0.2.22–v0.2.24 identity/status/public-bundle work | Preserved in current lean sections above; these supersede v0.2.21 prep-only status. |

---

## 14. Historical Continuity Payload from v0.2.21

The following snapshot is intentionally retained because it contains the detailed operational memory that the lean v0.2.24 file compressed away. It is not the current-state authority where newer sections above say otherwise.

# Historical Carry-Forward Snapshot: CarbonStack LogDoc v0.2.21

**Last updated:** 2026-05-23 16:20 -04:00  
**Owner/operator:** bitcrusher32  
**Assistant/model:** GPT-5.5 Thinking  
**Current phase:** Phase 2D identity prep state events checkpoint: `docs/48-provider-identity-prep-state-events-result-v0.md` records stabilization of `provider.identity.prep_state_written`, `provider.identity.exists`, and `checkpoint.failed`; Go provider-event taxonomy and provider-trust mapping now classify/map these state-skeleton events; sidecar tests compare prep-state/duplicate/checkpoint events against typed constants; `identity-create` still writes only dev-only non-secret prep state and does not generate OpenMLS signer/credential/KeyPackage/provider storage; Comms/Cypher/trust runtime integration remains intentionally absent.
**Version schema:** v[scope].[timeline] — this file is `v0.2.21`, the Phase 2D identity prep state events checkpoint after `provider.identity.prep_state_written`, `provider.identity.exists`, and `checkpoint.failed` were stabilized across taxonomy, trust mapping, sidecar tests, and docs. Multi-digit timeline terms are acceptable if they follow the same schema.

---

## 1. Project Goal

**Active goal:** Preserve CarbonStack state at the Phase 2D identity prep state events checkpoint. Phase 2C remains closed for mainline work. Phase 2D now has the Rust↔Go sidecar boundary, common JSON envelopes, stabilized command-surface events, recognized `identity-create`, the first dev-only state-write skeleton, and stabilized identity-prep state events. `identity-create --device-label <safe>` still validates the label, creates an ignored per-device state directory, writes a non-secret `identity-prep.json` manifest, and returns a success envelope with `ok: true`, `identity_created: false`, `state_written: true`, `provider_storage_written: false`, and `private_material_included: false`. Duplicate runs still refuse overwrite with `identity_prep_state_already_exists`, `provider.identity.exists`, and exit code 3. New at this checkpoint: `provider.identity.prep_state_written`, `provider.identity.exists`, and `checkpoint.failed` are represented in Go provider-event taxonomy/trust mapping and recorded in `docs/48`. This is still not real OpenMLS identity creation: no signing key, credential, KeyPackage, MemoryStorage, provider storage, or recovery material is generated or written. The sidecar still does not export bundles, create/join conversations, protect/open messages, route through Cypher, call Comms runtime flows, or touch trust-state storage.

CarbonStack now has a provider-neutral protocol seam, Phase 2A trust behavior, Phase 2B provider skeleton, Phase 2C MLS/OpenMLS planning docs through `docs/38`, a Rust-only OpenMLS scratch crate, validated one-message/two-message/same-process-reload/MemoryStorage-file-persistence/provider-fixture OpenMLS scratch probes plus Go-side fixture parsing/event classification and negative fixture parsing/classification, canonical docs summarizing the scratch result and provider-boundary implications, and a Rust artifact guard. This is still not a real CarbonStack provider, not wired into Comms/Cypher, and not production secure messaging.

**Component model:**

- **CarbonStackOS** — deliberately constrained Android-derived appliance OS. Deferred.
- **CarbonStackComms** — text-first messaging client; currently CLI-first with stub crypto, Phase 2A trust behavior, Phase 2B provider-boundary skeleton, and Phase 2C Rust-only OpenMLS scratch research path plus pure Go provider-contract/trust-decision mapping; Phase 2D sidecar command-surface planning is defined, `provider-info` exists, the sidecar emits common success/error JSON envelopes, command-surface events are stabilized, `identity-create` is recognized, the first dev-only non-secret prep-state write path exists, and identity prep state events are now typed/classified/trust-mapped. Runtime Comms/Cypher/trust integration remains absent.
- **CarbonStackCypher** — hostile-server opaque-envelope relay.
- **carbonstack** — canonical doctrine/specification repo, protocol docs, validation matrix, and local validation runner owner.

**North star:**

> Its goal is not to be a better smartphone. Its goal is to stop being a smartphone.

**Core doctrine:**

> Every feature is guilty until it proves it does not add unacceptable parser, network, sensor, identity, or filesystem authority.

---

## 2. Explicit Non-Goals / Out of Scope

- Do not start CarbonStackOS build work.
- Do not start Android implementation.
- Do not implement custom cryptography.
- Do not wire OpenMLS, mls-rs, libsignal, or any real provider into CarbonStackComms CLI behavior yet.
- Do not claim production security, Signal-equivalent security, hostile-server proof, replay resistance, metadata privacy, or production E2EE.
- Do not treat stub/base64 payloads, mock-provider payloads, fake/dev fingerprints, README-only MLS slots, or the dependency-probe crate as real encryption or real identity verification.
- Do not add browser/WebView, app store, arbitrary APK installation, rich previews, attachments, media ecosystems, or general file management.
- Do not add groups as a user-facing feature yet, even though conversations are conceptually group-shaped.
- Do not permanently couple CarbonStackComms to AGPL/libsignal unless a future decision explicitly accepts that tradeoff.
- Do not wire Rust scratch code into Go CLI, Cypher, `trust.json`, `trust-events.jsonl`, or production state.

---

## 3. Current State

| Classification | Items |
|---|---|
| **VALIDATED** | Phase 1 relay/client skeleton works locally. Cypher API tests pass. Comms package tests pass. Comms lifecycle smoke test passes. Phase 2A trust lifecycle passes. `internal/protocol` provider skeleton package tests pass. Local validation runner passes after OpenMLS dependency-probe work. |
| **VALIDATED DOC DECISION** | MLS/OpenMLS planning docs exist through `docs/48-provider-identity-prep-state-events-result-v0.md`, and the first Phase 2D Rust sidecar crate exists with a non-secret `provider-info` command plus a recognized `identity-create` command. Success, unsupported-command, missing-label, invalid-label, prep-state-write, and duplicate-overwrite-refusal JSON envelopes are invoked/parsed by Go tests. `provider.command.unsupported`, `provider.command.invalid`, and `provider.command.not_implemented` are typed warning/non-trust-relevant operational events; `provider.identity.prep_state_written`, `provider.identity.exists`, and `checkpoint.failed` are stabilized as typed provider events/trust decisions. Docs still define the authoritative architecture; the sidecar remains a runtime-boundary bootstrap. |
| **VALIDATED RESEARCH SLOT** | `carbonstack-comms/internal/protocol/mls/README.md` and `carbonstack-comms/internal/protocol/mls/research/openmls-minimal/README.md` reserve the future experimental MLS/OpenMLS path. |
| **VALIDATED OPENMLS PROBES** | `carbonstack-comms/internal/protocol/mls/research/openmls-minimal` now contains a Rust scratch crate. Dependency/build, credential/KeyPackage, group-add, Welcome join, application-message, two-message state-continuity, same-process provider-storage reload, MemoryStorage file-persistence, and sanitized provider-fixture probes have passed by user report. Latest user-reported checkpoint includes the negative provider fixture result after `9a3f80a test: classify provider fixture events` and `4dc350b docs: define provider event taxonomy`; exact v0.2.11 commit hashes were not visible in the prompt, but user reported the negative fixture code/doc block pushed and snapshot clean. |
| **PARTIAL** | Phase 2C now validates local scratch-level group creation, Bob join, one-message protect/open, two sequential Alice-to-Bob application messages, same-process `MlsGroup::load(provider.storage(), group_id)` reload, scratch-only MemoryStorage file save/load across fresh provider construction, and dev-only sanitized fixture/event/error summary output, Go-side fixture parsing tests, and provider event taxonomy/classification tests and negative fixture tests, and pure Go provider trust-decision mapping tests. Production storage, secure vault design, real provider mapping, Go/Comms integration, Cypher routing, trust-state integration, and hostile-server testing are not validated. |
| **NOT VALIDATED** | Production-safe storage, secure local vault, provider state export/import strategy, OpenMLS provider integration, Go/Rust integration, Cypher envelope routing with MLS payloads, trust-state integration, mls-rs integration, libsignal integration, final protocol selection, hostile-server harness, replay resistance, metadata privacy, hardware-key flows, Android, CarbonStackOS, CI. |
| **BLOCKED / NEXT** | The next step is implementing real dev-only OpenMLS identity material generation carefully, now that prep-state/write/duplicate/failure event semantics are stabilized. No code should touch Comms/Cypher/trust runtime integration, and stdout must remain sanitized/no-secret even when the sidecar begins generating local dev identity state. Direct Comms/Cypher wiring remains blocked until the sidecar contract is stable. Production security claims remain blocked on real protocol integration, review, local vault work, hardware-key work, and hostile-server validation. |

---

## 4. Critical Paths

### Local repo paths

- `C:\▮▮`
- `C:\▮▮\carbonstack`
- `C:\▮▮\carbonstack-comms`
- `C:\▮▮\carbonstack-cypher`
- `C:\▮▮\carbonstack-os`

### Canonical docs and runner

- `carbonstack/docs/16-phase1-integration-plan.md`
- `carbonstack/docs/17-phase1-test-matrix.md`
- `carbonstack/docs/18-protocol-threat-requirements-v0.md`
- `carbonstack/docs/19-protocol-candidate-evaluation.md`
- `carbonstack/docs/20-identity-and-trust-state-v0.md`
- `carbonstack/docs/21-phase2-protocol-plan.md`
- `carbonstack/docs/22-protocol-feasibility-matrix.md`
- `carbonstack/docs/23-phase2a-trust-state-plan.md`
- `carbonstack/docs/24-protocol-provider-boundary-v0.md`
- `carbonstack/docs/25-mls-feasibility-spike-plan.md`
- `carbonstack/docs/26-mls-implementation-candidate-notes.md`
- `carbonstack/docs/27-openmls-minimal-example-plan.md`
- `carbonstack/docs/28-openmls-upstream-example-notes.md`
- `carbonstack/docs/29-openmls-scratch-result-v0.md`
- `carbonstack/docs/30-openmls-provider-boundary-implications.md`
- `carbonstack/docs/31-openmls-persistence-spike-plan.md`
- `carbonstack/docs/32-openmls-storage-decision-v0.md`
- `carbonstack/docs/33-openmls-memory-storage-persistence-result-v0.md`
- `carbonstack/docs/34-openmls-provider-fixture-contract-plan.md`
- `carbonstack/docs/35-openmls-provider-fixture-result-v0.md`
- `carbonstack/docs/36-provider-event-taxonomy-v0.md`
- `carbonstack/docs/37-provider-negative-fixture-result-v0.md`
- `carbonstack/docs/38-provider-trust-state-mapping-v0.md`
- `carbonstack/docs/39-phase2d-sidecar-command-surface-plan.md`
- `carbonstack/docs/40-openmls-sidecar-provider-info-result-v0.md`
- `carbonstack/docs/41-openmls-sidecar-json-envelope-plan-v0.md`
- `carbonstack/docs/42-openmls-sidecar-json-envelope-result-v0.md`
- `carbonstack/docs/43-provider-command-unsupported-event-result-v0.md`
- `carbonstack/docs/44-openmls-sidecar-identity-create-plan-v0.md`
- `carbonstack/docs/45-openmls-sidecar-identity-create-prep-result-v0.md`
- `carbonstack/docs/46-provider-command-invalid-not-implemented-events-result-v0.md`
- `carbonstack/docs/47-openmls-sidecar-identity-prep-state-result-v0.md`
- `carbonstack/docs/48-provider-identity-prep-state-events-result-v0.md`
- `carbonstack/scripts/validate-local.ps1`
- `carbonstack/scripts/validate-phase1.ps1` compatibility wrapper

### CarbonStackComms validated implementation/test paths

- `carbonstack-comms/internal/protocol/types.go`
- `carbonstack-comms/internal/protocol/mock.go`
- `carbonstack-comms/internal/protocol/mock_test.go`
- `carbonstack-comms/internal/protocol/mls/README.md`
- `carbonstack-comms/internal/protocol/mls/research/openmls-minimal/README.md`
- `carbonstack-comms/internal/protocol/mls/research/openmls-minimal/Cargo.toml`
- `carbonstack-comms/internal/protocol/mls/research/openmls-minimal/Cargo.lock`
- `carbonstack-comms/internal/protocol/mls/research/openmls-minimal/src/main.rs`
- `carbonstack-comms/internal/protocol/mls/research/openmls-minimal/fixtures/dev/provider-summary.json`
- `carbonstack-comms/internal/protocol/mls/research/openmls-minimal/fixtures/dev/provider-events.jsonl`
- `carbonstack-comms/internal/protocol/mls/research/openmls-minimal/fixtures/dev/invalid-signature-error.json`
- `carbonstack-comms/internal/protocol/mls/research/openmls-minimal/fixtures/dev/missing-storage-error.json`
- `carbonstack-comms/internal/protocol/mls/research/openmls-minimal/fixtures/dev/missing-signer-error.json`
- `carbonstack-comms/internal/protocol/mls/research/openmls-minimal/fixtures/dev/wrong-group-error.json`
- `carbonstack-comms/internal/protocol/mls/research/openmls-minimal/fixtures/dev/malformed-message-error.json`
- `carbonstack-comms/internal/protocol/openmls_fixture_test.go`
- `carbonstack-comms/internal/protocol/provider_events.go`
- `carbonstack-comms/internal/protocol/provider_events_test.go`
- `carbonstack-comms/internal/protocol/openmls_negative_fixture_test.go`
- `carbonstack-comms/internal/protocol/provider_trust.go`
- `carbonstack-comms/internal/protocol/provider_trust_test.go`
- `carbonstack-comms/internal/trust/trust.go`
- `carbonstack-comms/internal/trust/trust_test.go`
- `carbonstack-comms/scripts/test-local-lifecycle.ps1`
- `carbonstack-comms/scripts/test-trust-lifecycle.ps1`
- `carbonstack-comms/scripts/check-no-rust-artifacts.ps1`

### Rust/OpenMLS scratch path

- `carbonstack-comms/internal/protocol/mls/research/openmls-minimal`

### Rust/OpenMLS sidecar path

- `carbonstack-comms/internal/protocol/mls/research/openmls-sidecar`
- `carbonstack-comms/internal/protocol/mls/research/openmls-sidecar/Cargo.toml`
- `carbonstack-comms/internal/protocol/mls/research/openmls-sidecar/Cargo.lock`
- `carbonstack-comms/internal/protocol/mls/research/openmls-sidecar/README.md`
- `carbonstack-comms/internal/protocol/mls/research/openmls-sidecar/src/main.rs`
- `carbonstack-comms/internal/protocol/mls/research/openmls-sidecar/src/labels.rs`
- `carbonstack-comms/internal/protocol/mls/research/openmls-sidecar/src/state.rs`
- `carbonstack-comms/internal/protocol/openmls_sidecar_provider_info_test.go`

Current purpose:

- Phase 2D runtime-boundary bootstrap.
- Supports `cargo run -- provider-info`.
- Recognizes `cargo run -- identity-create --device-label <label>` and writes dev-only non-secret prep state for safe labels.
- Emits a common JSON success envelope for `provider-info`.
- Emits common JSON error envelopes for unsupported commands, invalid identity-create usage, invalid labels, and safe-label-not-implemented identity-create prep.
- Uses stable event candidate `provider.command.unsupported`, now typed/classified in Go taxonomy and mapped to non-blocking debug/history trust-decision behavior.
- Uses prep events `provider.command.invalid` and `provider.command.not_implemented`, now typed/classified in Go taxonomy and mapped to non-blocking debug/history trust-decision behavior.
- Writes `identity-prep.json` under the ignored sidecar dev-state path for safe-label `identity-create`, while still reporting `identity_created: false` and `provider_storage_written: false`.
- Uses stabilized identity prep state events `provider.identity.prep_state_written`, `provider.identity.exists`, and `checkpoint.failed` in Go taxonomy/trust mapping/tests.
- Does not generate secrets, write provider storage, call Comms CLI, route through Cypher, mutate trust-state storage, or handle user messages.


Current purpose:

- Rust-only dependency/build probe.
- No Go integration.
- No CarbonStackComms CLI integration.
- No Cypher integration.
- No trust-state integration.
- No production persistence.
- No security claims.

---

## 5. Protocol / MLS Direction

### Provider boundary consensus carried forward

CarbonStack proceeds with:

- **MLS-shaped provider-neutral architecture.**
- **Every conversation is conceptually group-shaped, including 1:1 conversations.**
- **Signal/libsignal remains a reference and fallback, not a mainline dependency now.**
- **Avoid AGPL dependencies in mainline unless absolutely necessary.**
- **Rust is acceptable inside protocol/provider modules if it serves the project.**
- **Future experimental MLS provider work lives under `carbonstack-comms/internal/protocol/mls`.**
- **OpenMLS is the first intended spike candidate; mls-rs remains a serious alternate.**

### New at v0.2.3

The Rust/OpenMLS scratch crate has advanced beyond dependency probing into the first Alice-side group-shape probe.

Latest visible Comms commits in the attached validation log:

- `carbonstack-comms` `1f6e36e test: probe OpenMLS group add flow`
- `carbonstack-comms` `d87c86b test: probe OpenMLS credential and KeyPackage APIs`

Validated by user report and attached local validation log:

- OpenMLS dependency/build path works.
- Basic credential/signature/KeyPackage APIs work.
- A pasted-local-docs workflow successfully grounded method selection.
- Alice and Bob setup material can be created.
- Alice can create an `MlsGroup`.
- Alice can add Bob from Bob's `KeyPackage`.
- OpenMLS produces add commit / Welcome material.
- Alice can merge the pending commit.
- Alice can inspect epoch/member count after add.
- Local CarbonStack validation still passes.

Still not implemented:

- Bob join from Welcome/StagedWelcome
- application message protect/open
- state export/import
- provider boundary mapping
- Go/Rust integration
- real secure messaging

---

### New at v0.2.5

The Rust/OpenMLS scratch crate now validates a complete local two-member message flow at scratch level:

- Alice/Bob setup material
- Alice group creation
- Bob add from KeyPackage
- Welcome extraction and Bob join
- Alice application message creation
- Bob message processing/opening
- plaintext match assertion

This materially strengthens OpenMLS feasibility, but it is still isolated research code. It does not mean OpenMLS is integrated into CarbonStackComms or selected as the final provider.

A major Git hygiene issue was also resolved: Cargo `target/` artifacts entered local unpushed history and caused a huge push attempt. The final clean recovery restored only source files from a backup branch and recommitted cleanly.

### New at v0.2.6

The post-application-message Phase 2C block added two canonical docs in `carbonstack`:

- `docs/29-openmls-scratch-result-v0.md`
- `docs/30-openmls-provider-boundary-implications.md`

The Comms repo also added a Rust artifact guard and cleaned tracked Cargo artifacts:

- `carbonstack-comms` `7cfb590 chore: remove tracked Rust artifacts and add guard`
- `carbonstack-comms/scripts/check-no-rust-artifacts.ps1`

The Rust/OpenMLS scratch crate then advanced to a two-message state-continuity probe:

- `carbonstack-comms` `8be5e36 test: probe OpenMLS two-message state continuity`

Validated by user report and repo snapshot:

- Two sequential Alice-to-Bob OpenMLS application messages work inside one process.
- Both messages are opened by Bob and plaintext matches.
- `create_message` required mutable Alice group state.
- `process_message` required mutable Bob group state.
- Design implication: outbound protect/send and inbound open/process are both state-mutating, persistence-relevant provider operations.
- This is still not disk persistence or process-restart recovery.


### New at v0.2.7

The post-v0.2.6 Phase 2C block completed the docs/hygiene refresh and moved from two-message in-process continuity into storage/reload decision work.

New/updated canonical docs in `carbonstack`:

- `docs/31-openmls-persistence-spike-plan.md`
- `docs/32-openmls-storage-decision-v0.md`

Current heads at the v0.2.7 checkpoint:

- `carbonstack` `74092be docs: record OpenMLS storage decision`
- `carbonstack-comms` `64eaebe chore: add direct OpenMLS traits dependency`
- `carbonstack-cypher` `0bfd5af chore: remove tracked Cypher local state artifacts`
- `carbonstack-os` `b537475 Add CarbonStackOS north star and initial appliance model`

New validated scratch result:

- Same-process provider-storage reload passed in `carbonstack-comms` at `5215d0d test: probe OpenMLS same-process storage reload`.
- The scratch crate loaded Alice and Bob groups from each device's provider storage using `MlsGroup::load(provider.storage(), group_id)`.
- Loaded groups preserved epoch/member count and could continue the conversation by sending/opening message two.

Storage decision:

- `OpenMlsRustCrypto` uses `RustCrypto + MemoryStorage`, but its `key_store` field is private and no mutable storage accessor was identified.
- `MemoryStorage` exposes feature-gated persistence methods: `save_to_file`, `save`, `load_from_file`, and `load`.
- Next spike is Option A-lite: create a CarbonStack scratch provider wrapper that owns `RustCrypto` and `MemoryStorage`, implements `OpenMlsProvider`, enables MemoryStorage persistence, and tests process-restart-shaped file save/load.
- This remains scratch-only and does not solve production storage, secure vault integration, or Comms/Cypher MLS integration.


### New at v0.2.8

The post-v0.2.7 Phase 2C block completed the Option A-lite MemoryStorage file-persistence spike and recorded it canonically.

New canonical doc in `carbonstack`:

- `docs/33-openmls-memory-storage-persistence-result-v0.md`

Current heads at the v0.2.8 checkpoint:

- `carbonstack` `71e2002 docs: record OpenMLS MemoryStorage persistence result`
- `carbonstack-comms` `58d7211 test: probe OpenMLS MemoryStorage file persistence`
- `carbonstack-cypher` `0bfd5af chore: remove tracked Cypher local state artifacts`
- `carbonstack-os` `b537475 Add CarbonStackOS north star and initial appliance model`

New validated scratch result:

- Option A-lite passed in `carbonstack-comms` at `58d7211 test: probe OpenMLS MemoryStorage file persistence`.
- The scratch crate now has `phase-a` and `phase-b` modes:
  - `cargo run -- phase-a`
  - `cargo run -- phase-b`
- Phase A creates Alice/Bob state, creates the group, adds Bob, has Bob join from Welcome, sends/opens message one, saves Alice signer, and saves Alice/Bob `MemoryStorage`.
- Phase B constructs fresh providers, loads Alice/Bob `MemoryStorage`, reloads Alice/Bob `MlsGroup` with `MlsGroup::load`, reloads Alice signer, sends message two, and has Bob process/open message two successfully.
- Bob-opened phase-B plaintext matched Alice plaintext after file save/load and fresh provider construction.

Important failure before success:

- Phase B initially failed with `ValidationError(InvalidSignature)` when it used a fresh Alice signer.
- This was useful protocol behavior: Bob correctly rejected a message signed with a different Alice signing key.
- Fix: persist Alice signer in Phase A and reload it in Phase B.
- Lesson: provider storage and signer identity persistence are both required for restart-continuity.

Storage and security interpretation:

- `MemoryStorage` file persistence is a scratch feasibility tool, not production storage.
- Alice signer JSON in OS temp is not secure storage and must not be treated as a final vault design.
- The result proves a process-restart-shaped scratch continuation path, not production E2EE or production key management.

Next direction:

- Do not wire OpenMLS into Comms/Cypher yet.
- Update provider-boundary docs with persistence lessons.
- Plan a fixture/provider-contract compatibility spike so Go-side provider types can learn Rust/OpenMLS outputs without binding Rust into the CLI yet.
- Continue keeping all real provider work isolated from user-facing messaging.


### New at v0.2.9

The post-v0.2.8 Phase 2C block completed provider-boundary persistence documentation refresh, defined a fixture/provider-contract plan, and generated the first sanitized OpenMLS provider fixture summaries.

New/updated canonical docs in `carbonstack`:

- `docs/30-openmls-provider-boundary-implications.md` gained v0.2.8 persistence lessons.
- `docs/34-openmls-provider-fixture-contract-plan.md` defines the fixture-first provider-contract path.
- `docs/35-openmls-provider-fixture-result-v0.md` records the successful fixture generation result.

Current heads at the v0.2.9 checkpoint:

- `carbonstack` `963ed7f docs: record OpenMLS provider fixture result`
- `carbonstack-comms` `b853460 test: add OpenMLS provider fixture summaries`
- `carbonstack-cypher` `0bfd5af chore: remove tracked Cypher local state artifacts`
- `carbonstack-os` `b537475 Add CarbonStackOS north star and initial appliance model`

New validated scratch result:

- The OpenMLS scratch crate now supports `cargo run -- fixtures`.
- Fixture mode generates dev-only sanitized summaries under `carbonstack-comms/internal/protocol/mls/research/openmls-minimal/fixtures/dev`.
- Generated fixtures include provider summary, Alice/Bob device summaries, public setup summaries, Welcome summary, conversation checkpoint summaries, message summaries, provider event JSONL, and the invalid-signature error fixture.
- Fixture outputs intentionally do not include signer JSON, MemoryStorage JSON, raw private key material, or provider storage.
- The fixture mode preserves concrete event names such as `provider.fixture.started`, `provider.public_bundle.created`, `conversation.created`, `conversation.welcome.created`, `conversation.member_added`, `conversation.welcome.staged`, `conversation.joined`, `message.protected`, `message.opened`, `conversation.loaded`, and `provider.fixture.completed`.
- `invalid-signature-error.json` preserves the known `ValidationError(InvalidSignature)` failure as candidate `provider.signature.invalid`, with suggested trust actions: block, warn, append trust event, and require reverify if identity changed.

Important architectural result:

- CarbonStack now has its first concrete bridge from Rust/OpenMLS scratch behavior to provider-contract shape without linking Rust into the Go CLI.
- The project remains pre-integration: no `comms send`/`inbox` OpenMLS wiring, no Cypher MLS routing, no production E2EE claim.

Next direction:

- Use the fixture result to design Go-side provider-contract review/tests and event/type-shape mapping.
- Map fixture event names into trust-state candidates before any sidecar/FFI/CLI integration.
- Continue negative-path fixture work where useful, but keep integration pressure controlled.



### New at v0.2.10

The post-v0.2.9 Phase 2C block made the provider fixture contract executable on the Go side and recorded the provider event taxonomy canonically.

New canonical doc in `carbonstack`:

- `docs/36-provider-event-taxonomy-v0.md`

Current heads at the v0.2.10 checkpoint:

- `carbonstack` `4dc350b docs: define provider event taxonomy`
- `carbonstack-comms` `9a3f80a test: classify provider fixture events`
- `carbonstack-cypher` `0bfd5af chore: remove tracked Cypher local state artifacts`
- `carbonstack-os` `b537475 Add CarbonStackOS north star and initial appliance model`

New validated Go-side provider-contract work:

- `carbonstack-comms` `893957f test: parse OpenMLS provider fixture summaries`
- `carbonstack-comms` `9a3f80a test: classify provider fixture events`

What is now validated:

- Go-side tests parse committed OpenMLS fixture summaries without invoking Rust.
- Go-side tests parse `provider-events.jsonl`.
- Go-side tests assert fixture/event shape for:
  - provider summary
  - Alice/Bob device summaries
  - invalid-signature error fixture
  - provider event stream
- Fixture path bug was found and corrected:
  - wrong path: `../internal/protocol/mls/research/openmls-minimal/fixtures/dev`
  - correct path from `internal/protocol` tests: `mls/research/openmls-minimal/fixtures/dev`
- Provider event taxonomy is now represented in Go code through event names, classes, severities, and trust-relevance classification.
- Fixture event names now map to known provider-event descriptors instead of remaining loose strings.
- `provider.signature.invalid` maps to trust/security class with security severity and trust relevance.
- Terminal/fatal provider errors are classified as fatal and trust/security relevant.
- Unknown future provider events intentionally map to warning/unknown and do not become trust-relevant automatically.

Provider taxonomy classes recorded:

- lifecycle
- public setup
- membership
- message
- storage/checkpoint
- trust/security
- terminal/fatal provider errors

Important architectural result:

- CarbonStack now has a concrete path from Rust/OpenMLS scratch behavior to Go provider-contract tests without runtime Rust coupling.
- This strengthens the fixture-first integration path.
- The project remains pre-integration: no `comms send`/`inbox` OpenMLS wiring, no Cypher MLS routing, no trust-state consumption, and no production E2EE claim.

Next direction:

- Add negative-path fixture expansion for missing storage, wrong group, malformed message, duplicate/replay-ish message, stale epoch, missing signer, and public bundle/identity replacement where practical.
- Map provider event descriptors to trust-state candidate behavior in docs first.
- Keep runtime integration pressure controlled until negative-path/error behavior is better understood.


### New at v0.2.11

The post-v0.2.10 Phase 2C block added negative provider fixtures and a canonical result doc.

New canonical doc in `carbonstack`:

- `docs/37-provider-negative-fixture-result-v0.md`

Current heads reported in the v0.2.12 checkpoint include the v0.2.11 work:

- `carbonstack` `2243cee docs: record provider negative fixture result`
- `carbonstack-comms` `d5174b7 test: add OpenMLS negative provider fixtures`

Validated negative fixture mappings:

- `missing-storage-error.json` maps to `storage.missing`, class `storage_checkpoint`, warning severity, not trust-relevant.
- `missing-signer-error.json` maps to `provider.secret.material.unavailable`, class `terminal_fatal`, fatal severity, trust-relevant.
- `wrong-group-error.json` maps to `provider.group.unrecoverable`, class `terminal_fatal`, fatal severity, trust-relevant.
- `malformed-message-error.json` maps to `provider.message.tamper.detected`, class `trust_security`, security severity, trust-relevant.

Important lesson:

- Go rejects BOM-prefixed JSON created by Windows PowerShell `Set-Content -Encoding UTF8`.
- Go-parsed fixtures should be written as UTF-8 without BOM.

### New at v0.2.12

The post-v0.2.11 Phase 2C block added provider-event-to-trust-action mapping docs and pure Go mapping code/tests.

New canonical doc in `carbonstack`:

- `docs/38-provider-trust-state-mapping-v0.md`

Current heads at the v0.2.12 checkpoint:

- `carbonstack` `1435143 docs: map provider events to trust actions`
- `carbonstack-comms` `23b98bc test: map provider events to trust decisions`
- `carbonstack-cypher` `0bfd5af chore: remove tracked Cypher local state artifacts`
- `carbonstack-os` `b537475 Add CarbonStackOS north star and initial appliance model`

New Go-side provider trust-mapping paths:

- `carbonstack-comms/internal/protocol/provider_trust.go`
- `carbonstack-comms/internal/protocol/provider_trust_test.go`

What is now validated:

- Pure Go `DecideProviderTrust(...)` maps provider event names/descriptors to candidate trust decisions.
- Candidate actions include append history, debug-only, warn user, block send, block receive, block open, quarantine message, require reverify, mark identity changed, show recovery path, stop operation, and fatal local state.
- Happy-path fixture events map to append-history behavior and do not block send/receive/open.
- `provider.signature.invalid` maps to block-open, warn-user, append-history, require-reverify, user-visible, and history-relevant behavior.
- `storage.missing` maps to stop-operation and show-recovery-path behavior and blocks send without automatically blocking open.
- `provider.secret.material.unavailable` maps to fatal-local-state, block-send, show-recovery-path, user-visible, and history-relevant behavior.
- `provider.message.tamper.detected` maps to block-open, quarantine-message, warn-user, append-history, user-visible, and history-relevant behavior.
- `provider.group.unrecoverable` maps to fatal-local-state, stop-operation, show-recovery-path, and blocks send/receive/open.
- Unknown future provider events map to append-history/debug-only and do not automatically block operations or require reverify.

Important architectural boundary:

- This is policy-shaping and pure mapping only.
- It does not mutate `trust.json`.
- It does not write `trust-events.jsonl`.
- It does not change CLI behavior.
- It does not wire OpenMLS into Comms/Cypher.
- It does not make production security claims.


### New at v0.2.13

The post-v0.2.12 block closed Phase 2C for mainline work and created the Phase 2D setup point.

New canonical doc in `carbonstack`:

- `docs/39-phase2d-sidecar-command-surface-plan.md`

Current heads at the v0.2.13 checkpoint:

- `carbonstack` `ddf883c docs: define Phase 2D sidecar command surface`
- `carbonstack-comms` `7d8d366 test: map OpenMLS fixtures to trust decisions`
- `carbonstack-cypher` `0bfd5af chore: remove tracked Cypher local state artifacts`
- `carbonstack-os` `b537475 Add CarbonStackOS north star and initial appliance model`

New Go-side fixture-backed trust-decision path:

- `carbonstack-comms/internal/protocol/openmls_trust_decision_fixture_test.go`

What is now validated:

- Negative fixture JSON now connects directly to `DecideProviderTrust(...)` through Go tests.
- `invalid-signature-error.json` now maps to `provider.signature.invalid` and validates block-open, warn-user, append-history, require-reverify, user-visible, and history-relevant behavior.
- `missing-storage-error.json` validates stop-operation/show-recovery-path and send-blocking behavior without automatically blocking open.
- `missing-signer-error.json` validates fatal-local-state, block-send, show-recovery-path, user-visible, and history-relevant behavior.
- `wrong-group-error.json` validates fatal-local-state, stop-operation, show-recovery-path, and blocking send/receive/open.
- `malformed-message-error.json` validates block-open, quarantine-message, warn-user, append-history, user-visible, and history-relevant behavior.

Phase 2D setup decision:

- Rust sidecar is the first intended runtime-provider integration path.
- FFI is deferred.
- Pure Go MLS rewrite is deferred unless OpenMLS becomes a blocker.
- Sidecar work must begin boringly, likely with `provider-info`, and must not wire into `comms send`, `comms inbox`, Cypher routing, or trust-store mutation yet.

Phase 2C closure meaning:

- Phase 2C is closed for mainline work.
- The project has enough provider feasibility, fixture, negative-path, event taxonomy, and trust-decision scaffolding to transition to Phase 2D planning/prototyping.
- This still does not validate production storage, secure vault design, hostile-server security, replay resistance, metadata privacy, or production E2EE.



### New at v0.2.14

The first Phase 2D sidecar bootstrap landed after Phase 2C closure.

New CarbonStackComms sidecar path:

- `carbonstack-comms/internal/protocol/mls/research/openmls-sidecar`

Current heads at the v0.2.14 checkpoint:

- `carbonstack`: `ddf883c docs: define Phase 2D sidecar command surface`
- `carbonstack-comms`: `01f2f9a test: add OpenMLS sidecar provider-info command`
- `carbonstack-cypher`: `0bfd5af chore: remove tracked Cypher local state artifacts`
- `carbonstack-os`: `b537475 Add CarbonStackOS north star and initial appliance model`

New sidecar files:

- `carbonstack-comms/internal/protocol/mls/research/openmls-sidecar/Cargo.toml`
- `carbonstack-comms/internal/protocol/mls/research/openmls-sidecar/Cargo.lock`
- `carbonstack-comms/internal/protocol/mls/research/openmls-sidecar/README.md`
- `carbonstack-comms/internal/protocol/mls/research/openmls-sidecar/src/main.rs`

Validated by user report:

- `cargo run -- provider-info` builds and runs in the sidecar crate.
- The command emits JSON only.
- JSON reports `provider: openmls`.
- JSON reports `implementation: carbonstack-openmls-sidecar`.
- JSON reports `mode: experimental-sidecar`.
- JSON reports `phase: phase2d-provider-info`.
- JSON reports `capabilities: [provider-info]`.
- JSON lists unsupported commands: `identity-create`, `public-bundle-export`, `conversation-create`, `conversation-add-member`, `conversation-join`, `message-protect`, `message-open`, `state-checkpoint`, and `state-load-check`.
- JSON reports `security_level: experimental; not production E2EE`.
- JSON reports `private_material_included: false`.
- JSON warnings explicitly state that OpenMLS is not wired into CarbonStackComms, Cypher does not route MLS payloads, trust-state storage does not consume provider events, and no secret-bearing sidecar commands are implemented.

Why this matters:

- Phase 2D has started without collapsing into premature Comms/Cypher integration.
- The runtime provider boundary now has a concrete executable foothold.
- The sidecar is deliberately boring and non-secret, which is the correct first command for a security-sensitive provider boundary.

Important boundary:

- This does not implement a real OpenMLS provider.
- This does not create identities.
- This does not create or join conversations.
- This does not protect/open messages.
- This does not save/load provider state.
- This does not integrate with `comms send`, `comms inbox`, Cypher routing, `trust.json`, or `trust-events.jsonl`.
- This does not validate production storage, hostile-server security, replay resistance, metadata privacy, or production E2EE.

Next direction:

- Add a tiny Go or script-level smoke test that invokes/parses `provider-info` JSON output.
- Define/record sidecar JSON response/error envelope expectations before adding secret-bearing commands.
- Keep the next Phase 2D rung narrow: provider-info test protection first, then schema/error envelope, then only later `identity-create` or public-bundle work.


### New at v0.2.15

The second Phase 2D sidecar bootstrap rung landed: Go-side provider-info parsing and the canonical result doc.

New canonical doc in `carbonstack`:

- `docs/40-openmls-sidecar-provider-info-result-v0.md`

New CarbonStackComms Go-side sidecar test path:

- `carbonstack-comms/internal/protocol/openmls_sidecar_provider_info_test.go`

Current heads at the v0.2.15 checkpoint:

- `carbonstack`: `dfb35b1 docs: record OpenMLS sidecar provider-info result`
- `carbonstack-comms`: `b53c2b4 test: parse OpenMLS sidecar provider-info output`
- `carbonstack-cypher`: `0bfd5af chore: remove tracked Cypher local state artifacts`
- `carbonstack-os`: `b537475 Add CarbonStackOS north star and initial appliance model`

Validated by user report and clean snapshot:

- Go-side `internal/protocol` tests invoke the Rust sidecar with `cargo run --quiet -- provider-info`.
- The test parses the JSON output into a provider-info struct.
- The test asserts `provider = openmls`.
- The test asserts `implementation = carbonstack-openmls-sidecar`.
- The test asserts `mode = experimental-sidecar`.
- The test asserts `phase = phase2d-provider-info`.
- The test asserts `private_material_included = false`.
- The test asserts `provider-info` appears in capabilities.
- The test asserts all secret-bearing/state-mutating commands remain listed as unsupported: `identity-create`, `public-bundle-export`, `conversation-create`, `conversation-add-member`, `conversation-join`, `message-protect`, `message-open`, `state-checkpoint`, and `state-load-check`.
- The test asserts warnings are present.
- `docs/40-openmls-sidecar-provider-info-result-v0.md` records this as the first Go-tested OpenMLS sidecar command boundary.

Why this matters:

- Phase 2D now has a test-protected Rust↔Go command boundary.
- The runtime-provider sidecar is still deliberately non-secret and non-state-mutating.
- Go can invoke and parse the sidecar without wiring it into Comms CLI behavior.
- This is the correct bridge before adding response/error envelopes or secret-bearing provider commands.

Important boundary:

- This does not implement identity creation.
- This does not implement public bundle export.
- This does not implement conversation creation/join/add.
- This does not implement message protect/open.
- This does not implement state checkpoint/load.
- This does not route MLS payloads through Cypher.
- This does not mutate `trust.json` or `trust-events.jsonl`.
- This does not validate production storage, hostile-server security, replay resistance, metadata privacy, or production E2EE.

Next direction:

- Define sidecar response and error envelope expectations before any secret-bearing command.
- Consider adding an unsupported-command error-envelope check after the envelope shape is documented.
- Do not jump directly to `identity-create`, `message-protect`, or `message-open`.


### New at v0.2.16

The third Phase 2D sidecar rung landed: JSON success/error envelopes for the experimental OpenMLS sidecar.

New canonical docs in `carbonstack`:

- `docs/41-openmls-sidecar-json-envelope-plan-v0.md`
- `docs/42-openmls-sidecar-json-envelope-result-v0.md`

Current heads at the v0.2.16 checkpoint:

- `carbonstack`: `a38cfb7 docs: record OpenMLS sidecar JSON envelope result`
- `carbonstack-comms`: `f6f3a3d test: envelope OpenMLS sidecar provider-info output`
- `carbonstack-cypher`: `0bfd5af chore: remove tracked Cypher local state artifacts`
- `carbonstack-os`: `b537475 Add CarbonStackOS north star and initial appliance model`

New/updated sidecar paths:

- `carbonstack-comms/internal/protocol/mls/research/openmls-sidecar/src/main.rs`
- `carbonstack-comms/internal/protocol/openmls_sidecar_provider_info_test.go`

Validated by user report and clean snapshot:

- `cargo run -- provider-info` still builds/runs in the sidecar crate.
- `provider-info` now emits a success envelope with `ok: true`.
- Success envelope includes top-level `command`, `provider`, `implementation`, `mode`, `phase`, `data`, `events`, `warnings`, and `private_material_included` fields.
- `provider-info` command-specific fields moved under `data`:
  - `data.capabilities`
  - `data.unsupported`
  - `data.security_level`
- `private_material_included` remains `false`.
- The unsupported command path is now machine-readable:
  - `cargo run -- identity-create` prints a JSON error envelope.
  - The process exits with code `2`.
  - The error envelope contains `ok: false`, `error.code: unsupported_command`, `error.provider_event: provider.command.unsupported`, `error.severity: warning`, and `error.trust_relevant: false`.
  - The envelope also includes an event entry for `provider.command.unsupported`.
- The Go-side sidecar test now validates both:
  - `provider-info` success envelope.
  - `identity-create` unsupported-command error envelope.

Why this matters:

- The sidecar no longer has ad-hoc success-only JSON.
- Go now has a stable parse target for both success and failure.
- Unsupported/future commands no longer collapse into unstructured stderr/process failure.
- The sidecar command contract is safer to extend because future commands should follow the same envelope shape.

Important boundary:

- This does not implement `identity-create`.
- This does not implement public bundle export.
- This does not implement conversation creation/add/join.
- This does not implement message protect/open.
- This does not implement state checkpoint/load.
- This does not wire OpenMLS into Comms runtime commands.
- This does not route MLS payloads through Cypher.
- This does not mutate `trust.json` or `trust-events.jsonl`.
- This does not validate production storage, hostile-server security, replay resistance, metadata privacy, or production E2EE.

New event vocabulary candidate:

- `provider.command.unsupported`

Current meaning:

- The sidecar received a command that is intentionally unsupported/future at this phase.
- This is an operational warning, not a cryptographic trust failure.
- Current severity is `warning`.
- Current trust relevance is `false`.

Next direction:

- Plan `identity-create` carefully before implementation.
- Define command input, output envelope, public summary shape, local state path, private material rules, provider events, checkpoint behavior, and error cases.
- Only after the plan lands, implement a minimal dev-only identity-create sidecar command.
- Do not jump directly to message protect/open.



### New at v0.2.17

The fourth Phase 2D sidecar rung landed: `provider.command.unsupported` was stabilized across Go provider event taxonomy, provider trust-decision mapping, sidecar envelope tests, and canonical docs.

New canonical doc in `carbonstack`:

- `docs/43-provider-command-unsupported-event-result-v0.md`

Current heads at the v0.2.17 checkpoint:

- `carbonstack`: `803f05d docs: record unsupported sidecar command event`
- `carbonstack-comms`: `5abd4a2 test: classify unsupported sidecar command event`
- `carbonstack-cypher`: `0bfd5af chore: remove tracked Cypher local state artifacts`
- `carbonstack-os`: `b537475 Add CarbonStackOS north star and initial appliance model`

New/updated Comms paths:

- `carbonstack-comms/internal/protocol/provider_events.go`
- `carbonstack-comms/internal/protocol/provider_events_test.go`
- `carbonstack-comms/internal/protocol/provider_trust.go`
- `carbonstack-comms/internal/protocol/provider_trust_test.go`
- `carbonstack-comms/internal/protocol/openmls_sidecar_provider_info_test.go`

Validated by user report and clean snapshot:

- `ProviderEventCommandUnsupported` now exists as `provider.command.unsupported`.
- `DescribeProviderEvent(ProviderEventCommandUnsupported)` maps to:
  - class: `lifecycle`
  - severity: `warning`
  - trust relevant: `false`
- `DecideProviderTrust(ProviderEventCommandUnsupported)` maps to:
  - `append_history`
  - `debug_only`
  - no send/receive/open blocking
  - no reverify requirement
  - not user-visible
  - history-relevant for developer/audit continuity
- The sidecar unsupported-command envelope test now compares against `string(ProviderEventCommandUnsupported)` rather than a loose string literal.
- The previously observed failing state, where `ProviderEventCommandUnsupported` was accidentally captured by the debug lifecycle branch or fell through to unknown, was corrected by adding an explicit switch case.
- Go tests passed after the explicit warning-severity switch case was added.
- The local validation/snapshot after docs and code pushes was clean.

Why this matters:

- Unsupported sidecar commands are now structured operational events, not loose strings.
- Unsupported commands remain machine-readable without being misclassified as cryptographic trust/security failures.
- Future Phase 2D commands can rely on the sidecar envelope path while keeping unsupported/future command behavior non-alarming and non-secret.
- This closes the immediate post-envelope taxonomy gap from v0.2.16.

Important boundary:

- This does not implement `identity-create`.
- This does not implement public bundle export.
- This does not implement conversation creation/add/join.
- This does not implement message protect/open.
- This does not implement state checkpoint/load.
- This does not wire OpenMLS into Comms runtime commands.
- This does not route MLS payloads through Cypher.
- This does not mutate `trust.json` or `trust-events.jsonl`.
- This does not validate production storage, hostile-server security, replay resistance, metadata privacy, or production E2EE.

Next direction:

- Plan `identity-create` carefully before implementation.
- The identity-create plan should define command input, output envelope, public summary shape, local state write behavior, storage path expectations, private material rules, provider events, checkpoint behavior, error cases, and validation flow.
- Do not jump directly to message protect/open.


### New at v0.2.18

The fifth Phase 2D sidecar rung landed: `identity-create` was planned, then recognized by the sidecar for argument-validation-only prep behavior without generating secrets or writing state.

New canonical docs in `carbonstack`:

- `docs/44-openmls-sidecar-identity-create-plan-v0.md`
- `docs/45-openmls-sidecar-identity-create-prep-result-v0.md`

Current heads at the v0.2.18 checkpoint:

- `carbonstack`: `6b77f3e docs: record OpenMLS sidecar identity-create prep result`
- `carbonstack-comms`: `2750a45 test: prep OpenMLS sidecar identity-create validation`
- `carbonstack-cypher`: `0bfd5af chore: remove tracked Cypher local state artifacts`
- `carbonstack-os`: `b537475 Add CarbonStackOS north star and initial appliance model`

New/updated Comms paths:

- `carbonstack-comms/.gitignore`
- `carbonstack-comms/internal/protocol/mls/research/openmls-sidecar/src/main.rs`
- `carbonstack-comms/internal/protocol/mls/research/openmls-sidecar/src/labels.rs`
- `carbonstack-comms/internal/protocol/openmls_sidecar_provider_info_test.go`

Validated by user report and clean snapshot:

- `provider-info` now lists `identity-create` under `data.capabilities`.
- `identity-create` is no longer listed under unsupported commands.
- `provider-info` includes a warning that `identity-create` is recognized for argument validation only and does not generate secrets yet.
- The sidecar local dev-state directory is ignored through `.gitignore`:
  - `internal/protocol/mls/research/openmls-sidecar/.carbonstack-openmls-sidecar-state/`
- `labels.rs` validates device labels before future state-path use.
- Label validation accepts ASCII letters, numbers, dash, and underscore.
- Label validation rejects empty labels, `.`, `..`, path separators, spaces, unsupported punctuation, and labels longer than 96 characters.
- `cargo run -- identity-create` returns a structured JSON error envelope:
  - exit code: `2`
  - error code: `missing_required_argument`
  - provider event: `provider.command.invalid`
  - severity: `warning`
  - trust relevant: `false`
  - private material included: `false`
- `cargo run -- identity-create --device-label "../bad"` returns a structured JSON error envelope:
  - exit code: `2`
  - error code: `invalid_device_label`
  - provider event: `provider.command.invalid`
  - severity: `warning`
  - trust relevant: `false`
  - private material included: `false`
  - data echoes only the submitted label.
- `cargo run -- identity-create --device-label carbonstack-alice-device` returns a structured JSON error envelope:
  - exit code: `3`
  - error code: `not_implemented`
  - provider event: `provider.command.not_implemented`
  - severity: `warning`
  - trust relevant: `false`
  - private material included: `false`
  - `data.identity_created: false`
  - `data.state_written: false`
- `cargo run -- public-bundle-export` remains an unsupported command:
  - exit code: `2`
  - error code: `unsupported_command`
  - provider event: `provider.command.unsupported`
- Go-side tests now cover provider-info, unsupported public-bundle-export, identity-create missing-label, identity-create invalid-label, and identity-create safe-label-not-implemented behavior.
- All covered prep paths keep `private_material_included: false`.

Why this matters:

- `identity-create` moved from generic unsupported command into a recognized command surface without prematurely touching secrets.
- Future per-device state work has path-traversal guardrails before state write exists.
- The sidecar now has a safer ladder toward real identity generation: plan, validate args, validate labels, return no-secret/no-state envelopes, then only later generate dev-only identity material.
- This keeps the project from jumping directly from command recognition to cryptographic state mutation.

Important boundary:

- This does not implement OpenMLS signing key generation.
- This does not implement OpenMLS credential generation.
- This does not create KeyPackages.
- This does not write provider storage, signer JSON, MemoryStorage JSON, or recovery material.
- This does not export public bundles.
- This does not implement conversation creation/add/join.
- This does not implement message protect/open.
- This does not wire OpenMLS into Comms runtime commands.
- This does not route MLS payloads through Cypher.
- This does not mutate `trust.json` or `trust-events.jsonl`.
- This does not validate production storage, hostile-server security, replay resistance, metadata privacy, or production E2EE.

New event vocabulary candidates:

- `provider.command.invalid`
- `provider.command.not_implemented`

Current meaning:

- `provider.command.invalid`: recognized command failed syntax/argument validation.
- `provider.command.not_implemented`: recognized command passed validation but intentionally does not perform the operation yet.

Current classification intent:

- severity: `warning`
- trust relevant: `false`
- operational/developer-facing
- not a cryptographic failure

Next direction:

- Stabilize `provider.command.invalid` and `provider.command.not_implemented` in Go provider-event taxonomy and provider trust-decision mapping if they are intended to remain stable.
- Then proceed to the minimal real dev-only `identity-create` state write path.
- The real implementation must still print only sanitized public summaries and keep Comms CLI, Cypher routing, and trust-store mutation out of scope.
- Do not jump directly to message protect/open.



### New at v0.2.19

The sixth Phase 2D sidecar rung landed: identity-create prep command events were stabilized across Go provider event taxonomy, provider trust-decision mapping, sidecar envelope tests, and canonical docs.

New canonical doc in `carbonstack`:

- `docs/46-provider-command-invalid-not-implemented-events-result-v0.md`

Current heads at the v0.2.19 checkpoint:

- `carbonstack`: `74fafc6 docs: record identity-create prep command events`
- `carbonstack-comms`: `120cefc test: classify identity-create prep command events`
- `carbonstack-cypher`: `0bfd5af chore: remove tracked Cypher local state artifacts`
- `carbonstack-os`: `b537475 Add CarbonStackOS north star and initial appliance model`

New/updated Comms paths:

- `carbonstack-comms/internal/protocol/provider_events.go`
- `carbonstack-comms/internal/protocol/provider_events_test.go`
- `carbonstack-comms/internal/protocol/provider_trust.go`
- `carbonstack-comms/internal/protocol/provider_trust_test.go`
- `carbonstack-comms/internal/protocol/openmls_sidecar_provider_info_test.go`

Validated by user report and clean snapshot:

- `ProviderEventCommandInvalid` now exists as `provider.command.invalid`.
- `ProviderEventCommandNotImplemented` now exists as `provider.command.not_implemented`.
- `DescribeProviderEvent(ProviderEventCommandInvalid)` maps to:
  - class: `lifecycle`
  - severity: `warning`
  - trust relevant: `false`
- `DescribeProviderEvent(ProviderEventCommandNotImplemented)` maps to:
  - class: `lifecycle`
  - severity: `warning`
  - trust relevant: `false`
- `DecideProviderTrust(...)` for both prep command events maps to:
  - `append_history`
  - `debug_only`
  - no send/receive/open blocking
  - no reverify requirement
  - not user-visible
  - history-relevant for developer/audit continuity
- Sidecar Go tests now compare:
  - missing-label `identity-create` against `string(ProviderEventCommandInvalid)`
  - invalid-label `identity-create` against `string(ProviderEventCommandInvalid)`
  - safe-label not-implemented `identity-create` against `string(ProviderEventCommandNotImplemented)`
- The earlier copy-box/string-replacement breakage inserted constants into the switch area and caused `provider_events.go` syntax errors. It was repaired by removing misplaced constant lines, reinserting constants only in the const block, adding explicit switch cases, running `gofmt`, and re-running Go tests.
- Clean snapshot confirms the code/doc pushes landed.

Why this matters:

- `identity-create` prep behavior is now structured as typed operational sidecar events rather than loose strings.
- Invalid command usage and recognized-but-not-implemented command behavior remain machine-readable without being misclassified as cryptographic trust/security failures.
- The project can now proceed toward a real dev-only identity-create state write path with the command-prep error vocabulary already stabilized.
- This closes the immediate v0.2.18 gap before any code generates signing/provider state.

Important boundary:

- This does not implement real OpenMLS identity generation.
- This does not write signer JSON, MemoryStorage JSON, provider storage, or recovery material.
- This does not export public bundles.
- This does not implement conversation creation/add/join.
- This does not implement message protect/open.
- This does not wire OpenMLS into Comms runtime commands.
- This does not route MLS payloads through Cypher.
- This does not mutate `trust.json` or `trust-events.jsonl`.
- This does not validate production storage, hostile-server security, replay resistance, metadata privacy, or production E2EE.

Next direction:

- Implement the minimal dev-only `identity-create` state write path carefully.
- The implementation should generate only local dev identity material, write only under the ignored sidecar dev-state path, refuse overwrite by default, return sanitized JSON only, include `private_material_included: false`, and keep Comms CLI, Cypher routing, and trust-store mutation out of scope.
- Do not jump directly to public bundle export, conversation lifecycle, or message protect/open.



### New at v0.2.20

The seventh Phase 2D sidecar rung landed: `identity-create` now writes dev-only non-secret prep state and refuses duplicate prep state overwrites.

New canonical doc in `carbonstack`:

- `docs/47-openmls-sidecar-identity-prep-state-result-v0.md`

Current heads at the v0.2.20 checkpoint:

- `carbonstack`: `f502c1a docs: record OpenMLS sidecar identity prep state result`
- `carbonstack-comms`: `f41e98f test: add OpenMLS sidecar identity prep state skeleton`
- `carbonstack-cypher`: `0bfd5af chore: remove tracked Cypher local state artifacts`
- `carbonstack-os`: `b537475 Add CarbonStackOS north star and initial appliance model`

New/updated Comms paths:

- `carbonstack-comms/internal/protocol/mls/research/openmls-sidecar/src/main.rs`
- `carbonstack-comms/internal/protocol/mls/research/openmls-sidecar/src/state.rs`
- `carbonstack-comms/internal/protocol/openmls_sidecar_provider_info_test.go`

Validated by user report and clean snapshot:

- `cargo run -- provider-info` still exits `0`, returns `ok: true`, lists `provider-info` and `identity-create`, and reports `private_material_included: false`.
- `provider-info` now warns that `identity-create` writes dev-only non-secret prep state but does not generate OpenMLS secrets yet.
- `cargo run -- identity-create` still exits `2` with `missing_required_argument`, `provider.command.invalid`, warning severity, and `private_material_included: false`.
- `cargo run -- identity-create --device-label "../bad"` still exits `2` with `invalid_device_label`, `provider.command.invalid`, warning severity, and `private_material_included: false`.
- First run of `cargo run -- identity-create --device-label carbonstack-alice-device` exits `0` with:
  - `ok: true`
  - event `provider.identity.prep_state_written`
  - severity `notice`
  - trust relevant `false`
  - `identity_created: false`
  - `state_written: true`
  - `provider_storage_written: false`
  - `private_material_included: false`
- The sidecar writes a non-secret manifest at:
  - `.carbonstack-openmls-sidecar-state/dev/devices/carbonstack-alice-device/identity-prep.json`
- The manifest uses `manifest_version: identity-prep/v0` and reports:
  - `identity_created: false`
  - `provider_storage_written: false`
  - `private_material_included: false`
- Second run of the same safe-label command exits `3` and refuses overwrite with:
  - `identity_prep_state_already_exists`
  - event `provider.identity.exists`
  - `state_written: false`
  - `private_material_included: false`
- `cargo run -- public-bundle-export` remains unsupported with `provider.command.unsupported` and exit code `2`.
- Go-side tests validate provider-info, unsupported command, missing-label, invalid-label, prep-state success, duplicate-overwrite refusal, manifest JSON validity, and no-secret/no-provider-storage claims.
- The stale `not_implemented` safe-label path and stale Rust phase test expectation were cleaned up after the state-skeleton patch.

Why this matters:

- This is the first actual sidecar state-write rung.
- The sidecar can now create an ignored per-device state directory and write a non-secret manifest while preserving no-secret stdout behavior.
- Duplicate state creation is explicitly refused, preventing accidental silent overwrite behavior before real identity material exists.
- The project now has state-path and overwrite guardrails before introducing real OpenMLS signing/credential/provider storage material.

Important boundary:

- This does not implement real OpenMLS signing key generation.
- This does not implement OpenMLS credential generation.
- This does not create KeyPackages.
- This does not write signer JSON, MemoryStorage JSON, provider storage, or recovery material.
- This does not export public bundles.
- This does not implement conversation creation/add/join.
- This does not implement message protect/open.
- This does not wire OpenMLS into Comms runtime commands.
- This does not route MLS payloads through Cypher.
- This does not mutate `trust.json` or `trust-events.jsonl`.
- This does not validate production storage, hostile-server security, replay resistance, metadata privacy, or production E2EE.

New event vocabulary candidates:

- `provider.identity.prep_state_written`
- `provider.identity.exists`
- `checkpoint.failed`

Current meaning:

- `provider.identity.prep_state_written`: non-secret dev-only prep manifest was written.
- `provider.identity.exists`: existing identity/prep state was found and overwrite was refused.
- `checkpoint.failed`: intended state/checkpoint write failed.

Current classification intent:

- `provider.identity.prep_state_written`: lifecycle/storage-adjacent notice, trust relevant false.
- `provider.identity.exists`: lifecycle warning, trust relevant false.
- `checkpoint.failed`: storage/checkpoint warning, operationally important, not automatically cryptographic trust failure at this rung.

Next direction:

- Stabilize the new identity prep state events in Go provider-event taxonomy and provider trust-decision mapping.
- Then implement minimal dev-only real `identity-create` material generation with sanitized stdout only.
- Keep Comms CLI, Cypher routing, and trust-store mutation out of scope.
- Do not jump directly to public-bundle-export, conversation lifecycle, or message protect/open.



### New at v0.2.21

The eighth Phase 2D sidecar rung landed: identity prep state events were stabilized across Go provider event taxonomy, provider trust-decision mapping, sidecar tests, and canonical docs.

New canonical doc in `carbonstack`:

- `docs/48-provider-identity-prep-state-events-result-v0.md`

Current heads at the v0.2.21 checkpoint:

- `carbonstack`: `37803ee docs: record identity prep state events`
- `carbonstack-comms`: `ee60289 test: classify identity prep state events`
- `carbonstack-cypher`: `0bfd5af chore: remove tracked Cypher local state artifacts`
- `carbonstack-os`: `b537475 Add CarbonStackOS north star and initial appliance model`

New/updated Comms paths:

- `carbonstack-comms/internal/protocol/provider_events.go`
- `carbonstack-comms/internal/protocol/provider_events_test.go`
- `carbonstack-comms/internal/protocol/provider_trust.go`
- `carbonstack-comms/internal/protocol/provider_trust_test.go`
- `carbonstack-comms/internal/protocol/openmls_sidecar_provider_info_test.go`

Validated by user report and clean snapshot:

- `ProviderEventIdentityPrepStateWritten` now exists as `provider.identity.prep_state_written`.
- `ProviderEventIdentityExists` now exists as `provider.identity.exists`.
- `ProviderEventCheckpointFailed` is confirmed/stabilized as `checkpoint.failed` for this sidecar state-write path.
- `DescribeProviderEvent(ProviderEventIdentityPrepStateWritten)` maps to:
  - class: `storage/checkpoint`
  - severity: `info`
  - trust relevant: `false`
- `DescribeProviderEvent(ProviderEventIdentityExists)` maps to:
  - class: `storage/checkpoint`
  - severity: `warning`
  - trust relevant: `false`
- `DescribeProviderEvent(ProviderEventCheckpointFailed)` remains storage/checkpoint, warning, and non-trust-relevant.
- `DecideProviderTrust(ProviderEventIdentityPrepStateWritten)` maps to append-history/debug-only behavior and does not block send/receive/open, require reverify, or become user-visible.
- `DecideProviderTrust(ProviderEventIdentityExists)` maps to append-history/debug-only behavior and does not block send/receive/open, require reverify, or become user-visible.
- `DecideProviderTrust(ProviderEventCheckpointFailed)` maps to stop-operation/show-recovery-path behavior, blocks current outgoing/send mutation, is history-relevant, and does not require reverify by default.
- Sidecar Go tests compare prep-state and duplicate-state events against typed constants rather than loose strings.
- `checkpoint.failed` already existed in the older storage/checkpoint descriptor group; the final fix preserved the existing grouped handling instead of duplicating the switch case.
- The `checkpoint.failed` trust test was narrowed to avoid prematurely deciding all receive/open behavior in this rung.
- Go tests, hygiene, docs, and the final snapshot were reported clean.

Why this matters:

- Phase 2D has crossed into controlled local state writes, so event semantics now matter more.
- Prep-state write, duplicate-state refusal, and checkpoint failure now have typed/provider-mapped meanings before real OpenMLS identity material is generated.
- This prevents state-write behavior from remaining loose-string-only or being confused with cryptographic trust failures.
- The project now has a safer ladder: command surface, validation, non-secret prep state, stabilized state events, then real dev-only identity generation.

Important boundary:

- This does not implement real OpenMLS signing key generation.
- This does not implement OpenMLS credential generation.
- This does not create KeyPackages.
- This does not write signer JSON, MemoryStorage JSON, provider storage, or recovery material.
- This does not export public bundles.
- This does not implement conversation creation/add/join.
- This does not implement message protect/open.
- This does not wire OpenMLS into Comms runtime commands.
- This does not route MLS payloads through Cypher.
- This does not mutate `trust.json` or `trust-events.jsonl`.
- This does not validate production storage, hostile-server security, replay resistance, metadata privacy, or production E2EE.

Next direction:

- Implement minimal dev-only real `identity-create` material generation.
- The first real implementation should generate only the smallest viable OpenMLS identity/signing material needed for later public bundle work.
- State must remain under the ignored sidecar dev-state directory.
- Duplicate identity/prep state must refuse overwrite by default.
- stdout must remain sanitized with `private_material_included: false`.
- Keep Comms CLI, Cypher routing, and trust-store mutation out of scope.
- Do not jump directly to `public-bundle-export`, conversation lifecycle, message protect/open, or Comms runtime integration.


## 6. Current Phase Boundary

### Phase 1

Stable and test-protected local relay/client skeleton. Validates shape, lifecycle, and regression behavior only.

### Phase 2A

Validated trust-state scaffold:

- fake/dev fingerprints
- manual verification
- `trust-list`
- trust history
- strict/dev send policy
- key-change simulation
- reverify after changed identity
- revocation simulation
- local validation runner coverage

Still not real security.

### Phase 2B

Validated provider-boundary scaffold:

- provider-boundary doc
- provider-neutral code skeleton
- mock provider tests
- no real provider dependency in Go

### Phase 2C

Closed for mainline work as a local-only OpenMLS feasibility and provider-contract research phase:

- MLS feasibility spike plan exists.
- MLS candidate notes exist.
- OpenMLS minimal example plan exists.
- OpenMLS upstream example notes exist.
- OpenMLS scratch experiment directory exists.
- Rust/OpenMLS dependency probe crate exists.
- Credential/KeyPackage API probe exists.
- Alice-side group creation/add probe exists.
- Bob join from Welcome/StagedWelcome probe exists.
- MLS application message protect/open probe exists at Rust-only scratch level.
- OpenMLS scratch result doc exists.
- OpenMLS provider-boundary implications doc exists.
- Rust artifact guard exists in carbonstack-comms.
- Two-message in-process state-continuity probe exists at Rust-only scratch level.
- Scratch-only MemoryStorage file persistence across fresh provider construction is validated.
- Dev-only sanitized OpenMLS provider fixture summaries/events/errors are validated.
- Go-side fixture parsing and provider event classification tests are validated.
- Provider event taxonomy exists in canonical docs and Go code.
- Summary-only negative provider fixtures and Go classification tests are validated.
- Negative fixture result exists in `docs/37-provider-negative-fixture-result-v0.md`.
- Provider trust-state mapping exists in `docs/38-provider-trust-state-mapping-v0.md`.
- Provider event → candidate trust-decision mapping exists in canonical docs and pure Go tests.
- Mapping is still pure/pre-integration: no trust-store mutation, no CLI behavior, no runtime OpenMLS wiring.
- Fixture-backed trust-decision tests now connect positive/negative fixture examples to expected trust decisions.
- Phase 2D sidecar command-surface plan exists in `docs/39-phase2d-sidecar-command-surface-plan.md`.
- Production storage/secure vault design, real provider mapping, Go/Rust integration, and provider integration remain unvalidated.

### Later phases

- **Phase 2D:** active current phase: real provider integration path has begun with a deliberately boring Rust sidecar. `provider-info` builds/runs and emits a success envelope, unsupported commands emit JSON error envelopes, command-surface events are stabilized, and `identity-create` now writes dev-only non-secret prep state for safe device labels. Prep-state success and duplicate-overwrite refusal are Go-tested. Next work should stabilize `provider.identity.prep_state_written`, `provider.identity.exists`, and checkpoint/state-write failure behavior before real OpenMLS identity material generation. No Comms/Cypher/trust runtime integration yet.
- **Phase 2E:** hostile-server/protocol hardening.
- **Phase 2F:** security posture lift: vault, hardware-key flows, stricter release mode, Android/OS groundwork later.

---

## 7. Timeline / Active Iteration Notes

### Entries 1–48 — Through v0.2.1

**Compression note:** This section intentionally compresses CarbonStack LogDoc v0.2.1 timeline entries 1–48. It preserves the relevant handoff state while avoiding repeated forensic replay.

**Compressed continuity summary:**

- Repo family initialized under `C:\▮▮` and pushed remotely.
- Early Gitea web-edit divergence established fetch + pull rebase + push as the safe sync pattern.
- Base doctrine, architecture, and canonical spec docs were created, including the north star and “feature is guilty” doctrine.
- CarbonStackCypher was scaffolded as a Go/SQLite/HTTP JSON relay and manually validated for health, invite/account/device flows, and envelope submit/retrieve/ack lifecycle.
- CarbonStackComms CLI skeleton was implemented after a wrong-repo commit blunder into Cypher was corrected through normal cleanup commits.
- CLI-driven Alice/Bob lifecycle, Comms package tests, Cypher API tests, lifecycle scripts, validation matrix, and local validation runner were added.
- Generated DB/local state hygiene was corrected and ignore rules hardened.
- Phase 2 protocol foundation docs were added through feasibility matrix.
- Phase 2A trust-state scaffold was implemented and validated: fake/dev fingerprints, trust store, trust events, manual verification, strict/dev send policy, simulated key change, reverify, revocation, package tests, and lifecycle script.
- PowerShell expected-failure handling was fixed for negative-path tests.
- Validation runner was renamed to `validate-local.ps1`.
- Provider-boundary consensus selected MLS-shaped provider-neutral architecture, with Signal/libsignal as reference/fallback and no AGPL dependency in mainline unless necessary.
- `docs/24-protocol-provider-boundary-v0.md` was added.
- `carbonstack-comms/internal/protocol` was added with provider-neutral interface/types, mock provider, and tests.
- MLS/OpenMLS feasibility docs were added: `docs/25`, `docs/26`, `docs/27`, and `docs/28`.
- Experimental MLS provider path was reserved at `carbonstack-comms/internal/protocol/mls`.
- Rust-only OpenMLS scratch experiment path was reserved at `carbonstack-comms/internal/protocol/mls/research/openmls-minimal`.
- v0.2.1 clean heads were:
  - `carbonstack`: `5ec9c31 docs: add OpenMLS upstream example notes`
  - `carbonstack-comms`: `8ecfda8 docs: reserve OpenMLS minimal scratch experiment`
  - `carbonstack-cypher`: `0bfd5af chore: remove tracked Cypher local state artifacts`
  - `carbonstack-os`: `b537475 Add CarbonStackOS north star and initial appliance model`

### Entry 49 — 2026-05-21 — Rust installed/verified

**Action:** User installed Rust without VS Code and confirmed the setup worked.

**Known verification commands:**

- `rustc --version`
- `cargo --version`
- `rustup --version`

**Intent:** Prepare for Rust-only OpenMLS scratch work under the reserved research path.

**Classification:** VALIDATED TOOLCHAIN PREREQUISITE by user report.

### Entry 50 — 2026-05-21 — OpenMLS dependency probe added

**Action:** Added first Rust scratch crate under:

- `carbonstack-comms/internal/protocol/mls/research/openmls-minimal`

**Commit:** `carbonstack-comms` `7980ace test: add OpenMLS dependency probe`

**Expected/provided contents:**

- `Cargo.toml`
- `Cargo.lock`
- `src/main.rs`
- updated `README.md`

**Purpose:**

- Verify Rust/Cargo works inside the project tree.
- Verify OpenMLS-related dependency probe can be represented in the reserved research path.
- Keep the experiment isolated from Go CLI, Cypher, and trust-state behavior.
- Establish the first real Phase 2C code artifact without claiming MLS messaging works.

**Classification:** VALIDATED DEPENDENCY/BUILD PROBE by user report and local validation.

### Entry 51 — 2026-05-21 — Local validation after dependency probe

**Observed from user-provided validation summary:**

- `carbonstack`: `5ec9c31 docs: add OpenMLS upstream example notes`
- `carbonstack-comms`: `7980ace test: add OpenMLS dependency probe`
- `carbonstack-cypher`: `0bfd5af chore: remove tracked Cypher local state artifacts`
- `carbonstack-os`: `b537475 Add CarbonStackOS north star and initial appliance model`

**Validation status:**

- User reported local validation passed after the dependency probe commit.
- The pasted validation summary begins at repo status and shows the new `carbonstack-comms` head at `7980ace`.
- No changes were reported for `carbonstack-cypher` or `carbonstack-os`.

**Classification:** VALIDATED v0.2.2 BREAKPOINT.

---


### Entry 52 — 2026-05-21 — OpenMLS credential and KeyPackage API probe

**Action:** Advanced the Rust scratch crate from dependency/build probe into a first OpenMLS API probe.

**Commit visible in later history:** `carbonstack-comms` `d87c86b test: probe OpenMLS credential and KeyPackage APIs`.

**Purpose:**

- Validate the exact local OpenMLS dependency versions:
  - `openmls = "0.8.1"`
  - `openmls_basic_credential = "0.5.0"`
  - `openmls_rust_crypto = "0.5.1"`
- Validate `BasicCredential`, `SignatureKeyPair`, and `KeyPackage` setup material creation.
- Keep all work Rust-only and isolated from Go CLI/Cypher/trust state.
- Preserve narrow claims: setup material works; messaging does not.

**Blunder / correction:**

- Initial dependency add failed because a PowerShell command was accidentally glued together:
  - `cargo add openmls_basic_credential@0.5.0Get-Content .\Cargo.toml`
- Corrected by running `cargo add openmls_basic_credential@0.5.0` alone.
- Lesson: issue Cargo commands one line at a time; verify `Cargo.toml` after dependency edits.

**Classification:** VALIDATED API PROBE by user report.

### Entry 53 — 2026-05-21 — Pasted local docs became source-of-truth workflow

**Action:** User compiled key local `cargo doc` / OpenMLS method details into a paste, which was then used as the working source of truth for the group-add probe.

**Source material included:**

- `MlsGroup::new_with_group_id`
- `MlsGroup::add_members`
- `MlsGroup::merge_pending_commit`
- `MlsGroup::members`
- `MlsGroup::epoch`
- `StagedWelcome::new_from_welcome`
- `KeyPackage` / `KeyPackageBundle` behavior
- `StagedWelcome` and Welcome-stage API notes

**Why this mattered:**

The OpenMLS docs were large/noisy, but pasted local method signatures were precise enough to avoid blind API guessing. This workflow caught that:

- `add_members` expects `KeyPackage`, not `KeyPackageBundle`.
- `KeyPackage::builder().build(...)` returned `KeyPackageBundle` locally.
- The actual public `KeyPackage` must be extracted from the bundle before passing into `add_members`.

**New safe workflow:**

When upstream Rust docs are large, use:

1. `cargo doc --open -p openmls
    cargo doc --open -p openmls_memory_storage`
2. search the local docs for only the relevant structs/methods
3. paste the exact method signatures into chat
4. patch against the pasted local docs, not memory or stale snippets
5. stop on compiler errors and adapt to the local API

**Classification:** VALIDATED WORKFLOW / IMPORTANT LESSON.

### Entry 54 — 2026-05-21 — OpenMLS group-add probe validated

**Action:** Added and validated a Rust-only group-add probe.

**Commit:** `carbonstack-comms` `1f6e36e test: probe OpenMLS group add flow`

**Validated in scratch crate by user report:**

- Alice setup material exists.
- Bob setup material exists.
- Alice can create a local `MlsGroup`.
- Alice can add Bob from Bob's `KeyPackage`.
- OpenMLS emits add commit and Welcome material.
- Alice can merge the pending commit.
- Alice sees two members after add.
- Alice epoch/member inspection works enough for the current probe.

**Important compiler/API finding:**

The first group-add patch failed because `KeyPackage::builder().build(...)` produced a `KeyPackageBundle`, while `add_members` expected `KeyPackage`.

**Fix:**

Extract the public KeyPackage from the bundle:

    let key_package_bundle = KeyPackage::builder()
        .build(ciphersuite, provider, &signer, credential_with_key.clone())
        .expect("failed to build OpenMLS KeyPackageBundle");

    let key_package = key_package_bundle.key_package().clone();

**Validation after commit:**

Attached local validation log shows:

- `carbonstack`: `5ec9c31 docs: add OpenMLS upstream example notes`
- `carbonstack-comms`: `1f6e36e test: probe OpenMLS group add flow`
- `carbonstack-cypher`: `0bfd5af chore: remove tracked Cypher local state artifacts`
- `carbonstack-os`: `b537475 Add CarbonStackOS north star and initial appliance model`
- Cypher tests passed.
- Comms Go package tests passed.
- Local two-client CLI lifecycle passed.
- Phase 2A trust lifecycle passed.
- Final output: `PASS: CarbonStack local validation passed`.

**Note:** The attached validation log shows `1f6e36e (HEAD -> main)` without `origin/main` annotation for `carbonstack-comms`. If not already pushed separately, run the standard sync/hygiene flow before continuing.

**Classification:** VALIDATED v0.2.3 BREAKPOINT.


### Entry 55 — 2026-05-21 — OpenMLS Welcome join probe validated

**Action:** Advanced the Rust-only scratch crate from Alice-side group-add into a successful Bob join-from-Welcome probe.

**Commit:** `carbonstack-comms` `6ecf9e3 test: probe OpenMLS Welcome join flow`

**Validated in scratch crate by user report and attached validation log:**

- Alice setup material exists.
- Bob setup material exists.
- Alice can create a local `MlsGroup`.
- Alice can add Bob from Bob's `KeyPackage`.
- OpenMLS emits add commit and Welcome material.
- Welcome extraction works from `MlsMessageOut` by matching `welcome_msg.body()` against `MlsMessageBodyOut::Welcome(welcome)`.
- Alice can merge the pending commit.
- Bob can stage the Welcome through `StagedWelcome::new_from_welcome`.
- Bob can turn the staged Welcome into a local `MlsGroup`.
- Alice and Bob both see a two-member group.
- Epoch/member inspection works for both sides at scratch level.

**Important failed path / correction:**

- Attempting to serialize the full `MlsMessageOut` with `to_bytes()` and deserialize it directly as `Welcome` failed with `TrailingData`.
- Correct lesson: `MlsMessageOut` is a wrapper; the Welcome body should be extracted through `welcome_msg.body()` and the `MlsMessageBodyOut::Welcome` variant.

**Important storage/device-local lesson:**

- Initial Bob staging failed with `GroupAlreadyExists`.
- Root cause: Alice and Bob were using the same `OpenMlsRustCrypto::default()` provider/storage instance.
- Fix: use separate provider/storage instances:
  - `alice_provider = OpenMlsRustCrypto::default()`
  - `bob_provider = OpenMlsRustCrypto::default()`
- Design implication: provider storage is device-local protocol state. CarbonStack's future provider boundary must treat provider identity/state/conversation storage as per-device local state, not shared process-global state.

**Validation after commit:**

Attached local validation log shows:

- `carbonstack`: `5ec9c31 docs: add OpenMLS upstream example notes`
- `carbonstack-comms`: `6ecf9e3 test: probe OpenMLS Welcome join flow`
- `carbonstack-cypher`: `0bfd5af chore: remove tracked Cypher local state artifacts`
- `carbonstack-os`: `b537475 Add CarbonStackOS north star and initial appliance model`
- Cypher tests passed.
- Comms Go package tests passed.
- Local two-client CLI lifecycle passed.
- Phase 2A trust lifecycle passed.
- Final output: `PASS: CarbonStack local validation passed`.

**Note:** The attached validation log shows `6ecf9e3 (HEAD -> main)` without `origin/main` annotation for `carbonstack-comms`. If not already pushed separately, run the standard sync/hygiene flow before continuing.

**Classification:** VALIDATED v0.2.4 BREAKPOINT.


### Entry 56 — 2026-05-21 — OpenMLS application-message probe validated

**Action:** Advanced the Rust-only scratch crate from Welcome-join into the first successful local application-message protect/open probe.

**Commit:** `carbonstack-comms` `d75bf46 test: probe OpenMLS application message flow`

**Validated by user report and repo summary:**

- Alice and Bob still use separate `OpenMlsRustCrypto::default()` provider/storage instances.
- Alice creates the local MLS group.
- Alice adds Bob from Bob's `KeyPackage`.
- Welcome extraction still works through `MlsMessageBodyOut::Welcome`.
- Bob joins from `StagedWelcome` into his own `MlsGroup`.
- Alice creates an application message with `create_message`.
- Alice's `MlsMessageOut` is serialized.
- Bob deserializes into `MlsMessageIn`, converts to `ProtocolMessage`, and calls `process_message`.
- Bob extracts `ProcessedMessageContent::ApplicationMessage(...)` and calls `ApplicationMessage::into_bytes()`.
- Bob-opened plaintext matches Alice plaintext.

**Important API/state lesson:**

- `process_message` mutates Bob's group state, so `bob_group` had to be declared `mut`.
- Design implication: provider `open/process` should be treated as a state-mutating operation. Future provider persistence cannot model decrypt/open as a pure read-only call.

**Validation summary after push:**

- `carbonstack`: `5ec9c31 docs: add OpenMLS upstream example notes`
- `carbonstack-comms`: `d75bf46 test: probe OpenMLS application message flow`
- `carbonstack-cypher`: `0bfd5af chore: remove tracked Cypher local state artifacts`
- `carbonstack-os`: `b537475 Add CarbonStackOS north star and initial appliance model`

**Classification:** VALIDATED v0.2.5 APPLICATION-MESSAGE SCRATCH PROBE.

### Entry 57 — 2026-05-21 — Rust target artifact Git recovery ordeal

**Action:** Recovered from an accidental local history/staging problem where Rust/Cargo `target/` artifacts entered the local commit path and made the push attempt huge.

**Observed problem:**

- Push size ballooned to roughly 175 MiB.
- `git show --stat` showed `.exe`, `.pdb`, object files, `.fingerprint`, and many `target/` artifacts.
- After cleanup attempts, Git still showed many tracked-deletion lines under `internal/protocol/mls/research/openmls-minimal/target/...`.

**Recovery path that worked:**

- Created a safety branch: `git branch backup/pre-target-cleanup`.
- Verified `origin/main` did not already contain the `target/` tree with `git ls-tree -r origin/main -- internal/protocol/mls/research/openmls-minimal/target`.
- Reset back to `origin/main`.
- Restored only clean source files from the backup branch using `git restore --source backup/pre-target-cleanup -- <path>`.
- Recommitted only `.gitignore`, `Cargo.toml`, `Cargo.lock`, `README.md`, and `src/main.rs`.
- Final clean commit became `d75bf46 test: probe OpenMLS application message flow`.

**Lessons:**

- Rust `target/` must never be committed.
- `.gitignore` does not remove already tracked files.
- If build artifacts enter unpushed local history, rewrite the local commit before pushing.
- `git reset -- <path>` uses `--` to separate paths from revisions; without it, Git may report an ambiguous argument.
- When in doubt, create a backup branch before destructive recovery.
- Source recovery from a backup branch is safer than trying to manually delete hundreds of build artifacts from the index.

**Classification:** VALIDATED HYGIENE/RECOVERY LESSON.


### Entry 58 — 2026-05-22 — OpenMLS scratch result doc added

**Action:** Added a canonical OpenMLS scratch result document in the `carbonstack` repo.

**Commit:** `carbonstack` `6f10d2b docs: record OpenMLS scratch result`

**Path:**

- `carbonstack/docs/29-openmls-scratch-result-v0.md`

**Purpose:**

- Preserve the full scratch ladder:
  - dependency/build probe
  - credential/KeyPackage probe
  - group-add probe
  - Welcome join probe
  - application-message protect/open probe
- Clarify allowed/not-allowed claims.
- Preserve the Git/Cargo artifact hygiene lessons.
- Convert the last major scratch work into durable project knowledge rather than chat-scroll knowledge.

**Classification:** VALIDATED DOC CHECKPOINT by user report and repo snapshot.

### Entry 59 — 2026-05-22 — OpenMLS provider-boundary implications doc added

**Action:** Added a provider-boundary implications document in the `carbonstack` repo.

**Commit:** `carbonstack` `440131f docs: record OpenMLS provider boundary implications`

**Path:**

- `carbonstack/docs/30-openmls-provider-boundary-implications.md`
- `carbonstack/docs/31-openmls-persistence-spike-plan.md`
- `carbonstack/docs/32-openmls-storage-decision-v0.md`

**Purpose:**

- Translate the OpenMLS scratch result into architecture requirements.
- Record that the provider is not a stateless encrypt/decrypt utility; it is a stateful protocol engine.
- Identify likely boundary concepts:
  - provider identity
  - public setup material / KeyPackage equivalent
  - join material / Welcome equivalent
  - conversation/group state
  - provider-local storage
  - epoch
  - member summary
  - provider events
  - persistence hooks
  - state-updated markers
- Keep OpenMLS out of Comms/Cypher until provider state and persistence are understood.

**Classification:** VALIDATED DOC CHECKPOINT by user report and repo snapshot.

### Entry 60 — 2026-05-22 — Rust artifact guard and target cleanup

**Action:** Added a Rust/Cargo artifact hygiene guard in `carbonstack-comms` and removed tracked Rust build artifacts.

**Commit:** `carbonstack-comms` `7cfb590 chore: remove tracked Rust artifacts and add guard`

**New helper script:**

- `carbonstack-comms/scripts/check-no-rust-artifacts.ps1`

**Validated by user report:**

- The guard initially failed because tracked files still existed under:
  - `internal/protocol/mls/research/openmls-minimal/target/...`
- Cleanup removed those files from Git tracking.
- The guard then passed and the cleanup commit landed.

**Lessons:**

- `.gitignore` prevents new tracking but does not remove already tracked artifacts.
- `git rm -r --cached --ignore-unmatch -- internal/protocol/mls/research/openmls-minimal/target` is the correct cleanup path for tracked Cargo output.
- A one-time cleanup commit showing many deleted `target/` files is acceptable.
- Future Rust scratch commits should show no `target/`, `.fingerprint`, `.exe`, `.pdb`, `.o`, `.rlib`, or `.rmeta`.

**Classification:** VALIDATED HYGIENE GUARD.

### Entry 61 — 2026-05-22 — OpenMLS two-message state-continuity probe validated

**Action:** Advanced the Rust-only scratch crate from a one-message application probe into a two-message state-continuity probe.

**Commit:** `carbonstack-comms` `8be5e36 test: probe OpenMLS two-message state continuity`

**Validated by user report and repo snapshot:**

- Alice/Bob setup material still works.
- Alice creates the MLS group.
- Alice adds Bob from Bob's `KeyPackage`.
- Welcome extraction and Bob `StagedWelcome` join still work.
- Alice sends message one.
- Bob processes/opens message one and plaintext matches.
- Alice sends message two in the same process.
- Bob processes/opens message two and plaintext matches.
- In-memory Alice/Bob group/provider state remains usable across sequential messages inside one process.

**Important compiler/API lesson:**

- `create_message` required mutable Alice group state.
- Earlier `process_message` already required mutable Bob group state.
- Design implication: both outbound protect/send and inbound open/process are state-mutating provider operations.
- Future persistence design must checkpoint after outbound messages as well as inbound messages.

**Important limitation:**

This does not validate real disk persistence or process-restart recovery. It only validates in-process state continuity across sequential messages.

**Repo snapshot after cleanup and push:**

- `carbonstack`: `440131f docs: record OpenMLS provider boundary implications`
- `carbonstack-comms`: `8be5e36 test: probe OpenMLS two-message state continuity`
- `carbonstack-cypher`: `0bfd5af chore: remove tracked Cypher local state artifacts`
- `carbonstack-os`: `b537475 Add CarbonStackOS north star and initial appliance model`

**Classification:** VALIDATED v0.2.6 TWO-MESSAGE STATE-CONTINUITY SCRATCH PROBE.


### Entry 62 — 2026-05-22 — v0.2.6 docs/hygiene refresh completed

**Action:** Ran the post-v0.2.6 docs and hygiene consistency pass across the CarbonStack umbrella.

**Resulting commits:**

- `carbonstack` `2f5ffe5 docs: refresh OpenMLS scratch docs for v0.2.6`
- `carbonstack-comms` `39bf60b docs: refresh OpenMLS scratch README for state continuity`

**Purpose:**

- Bring OpenMLS scratch docs in line with the two-message state-continuity reality.
- Correct stale v0.2.5 references.
- Record that `create_message` and `process_message` are both state-mutating/persistence-relevant.
- Keep not-allowed claims explicit: no Comms/Cypher OpenMLS integration, no production E2EE, no disk persistence yet.

**Classification:** VALIDATED DOC/HYGIENE REFRESH by user report and clean snapshot.

### Entry 63 — 2026-05-22 — OpenMLS persistence spike plan and same-process reload result

**Action:** Added the persistence spike plan and then recorded the first persistence-rung result.

**Commits:**

- `carbonstack` `2e87eeb docs: define OpenMLS persistence spike plan`
- `carbonstack` `c5ce14c docs: record OpenMLS same-process reload result`

**Validated by user report:**

- `MlsGroup::load(storage, group_id)` exists and returns `Result<Option<MlsGroup>, Storage::Error>`.
- Alice and Bob groups can be loaded from each device's provider storage inside the same process.
- Loaded Alice/Bob groups preserve epoch and member count.
- Loaded Alice can create message two.
- Loaded Bob can process/open message two.
- Bob-opened plaintext for message two matches Alice plaintext.

**Important boundary lesson:**

Same-process provider-storage reload is stronger than keeping the same mutable variables alive, but it is still not disk persistence or process restart recovery.

**Classification:** VALIDATED SAME-PROCESS PROVIDER-STORAGE RELOAD RUNG.

### Entry 64 — 2026-05-22 — Direct OpenMLS traits dependency added after reload probe

**Action:** Added a direct `openmls_traits` dependency after the reload probe needed the `OpenMlsProvider` trait in scope for `provider.storage()`.

**Commit:**

- `carbonstack-comms` `64eaebe chore: add direct OpenMLS traits dependency`

**Lesson:**

If a trait method is used directly in the scratch crate, keep the trait crate as an explicit dependency rather than relying on transitive availability.

**Classification:** VALIDATED DEPENDENCY HYGIENE.

### Entry 65 — 2026-05-22 — OpenMLS storage API inspection and storage decision

**Action:** Inspected local `openmls_memory_storage` and `openmls_rust_crypto` source/docs to decide the next storage spike.

**Commit:**

- `carbonstack` `74092be docs: record OpenMLS storage decision`

**Findings:**

- `OpenMlsRustCrypto` is built from `RustCrypto` and `MemoryStorage`.
- `OpenMlsRustCrypto` implements `OpenMlsProvider`, exposing `storage()`, `crypto()`, and `rand()`.
- Its `key_store: MemoryStorage` field is private; no clean mutable storage accessor was identified in the visible source.
- `MemoryStorage` has a feature-gated persistence module with public `save_to_file`, `save`, `load_from_file`, and `load` methods.
- `MemoryStorage` persistence serializes the in-memory key/value map as JSON with base64-encoded keys and values.

**Decision:**

Proceed with Option A-lite: create a CarbonStack-only scratch provider wrapper that owns `RustCrypto` and `MemoryStorage`, implements `OpenMlsProvider`, enables `openmls_memory_storage` persistence, and tests file save/load in a process-restart-shaped scratch experiment.

**Explicit non-decision:**

This does not solve production provider storage, secure vault integration, Go/Rust sidecar design, or CarbonStackComms MLS integration.

**Classification:** VALIDATED v0.2.7 STORAGE DECISION BREAKPOINT.


### Entry 66 — 2026-05-22 — OpenMLS MemoryStorage file-persistence spike validated

**Action:** Implemented and validated the Option A-lite scratch persistence spike in the Rust-only OpenMLS scratch crate.

**Commit:**

- `carbonstack-comms` `58d7211 test: probe OpenMLS MemoryStorage file persistence`

**Implementation shape:**

- Added/used `openmls_memory_storage` persistence support.
- Added a CarbonStack-only scratch provider wrapper owning:
  - `RustCrypto`
  - `MemoryStorage`
- Implemented `OpenMlsProvider` for the scratch provider.
- Added two run modes:
  - `cargo run -- phase-a`
  - `cargo run -- phase-b`

**Validated phase-a by user report:**

- Alice/Bob setup material still works.
- Alice creates the MLS group.
- Alice adds Bob from Bob's `KeyPackage`.
- Welcome extraction and Bob `StagedWelcome` join still work.
- Alice sends phase-A message one.
- Bob processes/opens phase-A message one and plaintext matches.
- Alice signer is saved to OS temp JSON.
- Alice `MemoryStorage` is saved through OpenMLS memory-storage persistence helpers.
- Bob `MemoryStorage` is saved through OpenMLS memory-storage persistence helpers.

**Validated phase-b by user report:**

- Fresh Alice/Bob providers are created.
- Alice `MemoryStorage` is loaded from file.
- Bob `MemoryStorage` is loaded from file.
- Alice group is reloaded from loaded provider storage with `MlsGroup::load`.
- Bob group is reloaded from loaded provider storage with `MlsGroup::load`.
- Loaded Alice/Bob groups preserve epoch/member count.
- Alice signer is reloaded from temp JSON.
- Alice sends phase-B message two after reload.
- Bob processes/opens phase-B message two and plaintext matches.
- Phase B succeeds after fresh provider construction and file-load boundary.

**Important failure before success:**

- Phase B initially failed with `ValidationError(InvalidSignature)` when it created a fresh Alice signer.
- This proved Bob correctly rejects a message signed with a different Alice key.
- It also proved signer identity persistence is required in addition to provider/group storage persistence.
- Persisting/reloading Alice signer resolved the failure at scratch level.

**Canonical result doc:**

- `carbonstack` `71e2002 docs: record OpenMLS MemoryStorage persistence result`
- `carbonstack/docs/33-openmls-memory-storage-persistence-result-v0.md`

**Important limitations:**

- This is still scratch-only.
- Temp JSON signer persistence is not secure.
- `MemoryStorage` JSON/base64 file persistence is not a final vault design.
- OpenMLS is still not wired into Comms/Cypher.
- No production E2EE, hostile-server security, replay resistance, or metadata privacy claim is allowed.

**Classification:** VALIDATED v0.2.8 MEMORYSTORAGE FILE-PERSISTENCE SCRATCH PROBE.

### Entry 67 — 2026-05-22 — Provider-boundary persistence lessons and fixture plan added

**Action:** Updated canonical provider-boundary docs with v0.2.8 persistence lessons and added the fixture/provider-contract plan.

**Commits:**

- `carbonstack` `322f5ca docs: add OpenMLS persistence lessons to provider boundary`
- `carbonstack` `b70aa02 docs: define OpenMLS provider fixture contract plan`

**Purpose:**

- Record restart-aware provider-boundary requirements after the MemoryStorage persistence spike.
- Make signer identity persistence, provider-local storage, group reload, checkpoint-after-send, checkpoint-after-receive, and invalid-signature mapping first-class provider concepts.
- Define the fixture-first path so Rust/OpenMLS can emit sanitized provider-contract summaries before Go/Rust integration.

**Plan added:**

- `carbonstack/docs/34-openmls-provider-fixture-contract-plan.md`

**Decision:**

Proceed fixture-first rather than direct CLI wiring. This allows Go-side provider types/tests to learn concrete provider output shapes without binding Rust OpenMLS into the CLI yet.

**Classification:** VALIDATED DOC/ARCHITECTURE STEP.

### Entry 68 — 2026-05-22 — OpenMLS provider fixture summaries validated

**Action:** Added `cargo run -- fixtures` mode to the Rust-only OpenMLS scratch crate and committed sanitized dev-only fixture outputs.

**Commit:**

- `carbonstack-comms` `b853460 test: add OpenMLS provider fixture summaries`

**Canonical result doc:**

- `carbonstack` `963ed7f docs: record OpenMLS provider fixture result`
- `carbonstack/docs/35-openmls-provider-fixture-result-v0.md`

**Validated by user report:**

- Fixture mode succeeded.
- Generated fixture directory: `carbonstack-comms/internal/protocol/mls/research/openmls-minimal/fixtures/dev`.
- Generated files are small, summary-only JSON/JSONL files.
- `provider-summary.json` explicitly states fixtures are dev-only sanitized summaries and not production crypto integration.
- `provider-events.jsonl` contains provider/conversation/message lifecycle events.
- `invalid-signature-error.json` maps the known signer mismatch failure to candidate `provider.signature.invalid`.
- No MemoryStorage JSON, signer JSON, raw private keys, or provider storage were intentionally written into the fixture directory.

**Fixture files generated:**

- `provider-summary.json`
- `alice-device-summary.json`
- `bob-device-summary.json`
- `alice-public-setup-summary.json`
- `bob-public-setup-summary.json`
- `conversation-before-message.json`
- `welcome-summary.json`
- `conversation-after-message-one.json`
- `conversation-after-reload.json`
- `message-one-summary.json`
- `message-two-summary.json`
- `conversation-after-message-two.json`
- `invalid-signature-error.json`
- `provider-events.jsonl`

**Important interpretation:**

This is the first concrete bridge from OpenMLS scratch behavior to CarbonStack provider-contract shape. It is still not provider integration, not CLI integration, not Cypher routing, and not production E2EE.

**Classification:** VALIDATED v0.2.9 PROVIDER-FIXTURE CONTRACT RUNG.



### Entry 69 — 2026-05-22 — Go-side OpenMLS fixture parser and provider event taxonomy validated

**Action:** Converted the v0.2.9 provider-fixture result into executable Go-side provider-contract tests and canonical provider event taxonomy.

**Commits:**

- `carbonstack-comms` `893957f test: parse OpenMLS provider fixture summaries`
- `carbonstack-comms` `9a3f80a test: classify provider fixture events`
- `carbonstack` `4dc350b docs: define provider event taxonomy`

**New canonical doc:**

- `carbonstack/docs/36-provider-event-taxonomy-v0.md`

**New/updated Go-side paths:**

- `carbonstack-comms/internal/protocol/openmls_fixture_test.go`
- `carbonstack-comms/internal/protocol/provider_events.go`
- `carbonstack-comms/internal/protocol/provider_events_test.go`

**Validated by user report and checkpoint summary:**

- Go tests parse committed OpenMLS provider fixture JSON/JSONL without invoking Rust.
- Provider summary fixture is checked for version, provider name, mode, implementation, group ID label, and warnings.
- Alice/Bob device summary fixtures are checked for role, labels, KeyPackage hash reference length, and no private material.
- Invalid-signature fixture is checked for `ValidationError(InvalidSignature)`, candidate mapping `provider.signature.invalid`, suggested trust actions, and no private material.
- Provider event JSONL is parsed and required event names are asserted.
- Event counts for public bundles and message protect/open events are checked.
- Provider event names/classes/severities are now represented in Go.
- Fixture event names classify successfully through `DescribeProviderEvent`.
- Invalid signature maps to trust/security class, security severity, and trust relevance.
- Fatal provider invariant failures map to terminal/fatal class, fatal severity, and trust relevance.
- Unknown future provider events map to unknown/warning and are not automatically trust-relevant.

**Important path bug and fix:**

- Initial Go fixture tests failed because the fixture path assumed a repo-root working directory.
- During `go test ./internal/protocol`, the package working directory is `carbonstack-comms/internal/protocol`.
- Correct relative fixture directory is:
  - `mls/research/openmls-minimal/fixtures/dev`
- This was fixed before commit.

**Taxonomy classes now carried forward:**

- lifecycle
- public setup
- membership
- message
- storage/checkpoint
- trust/security
- terminal/fatal provider errors

**Architectural interpretation:**

This is still not OpenMLS integration.

It is an executable contract bridge: Rust/OpenMLS scratch emits dev-only sanitized fixtures; Go-side provider code can parse/classify those fixtures and reason about event names without runtime Rust coupling.

**Classification:** VALIDATED v0.2.10 PROVIDER EVENT TAXONOMY / GO FIXTURE CONTRACT RUNG.



### Entry 70 — 2026-05-22 — Negative provider fixtures and result doc validated

**Action:** Added the first summary-only negative provider fixture set, Go-side negative fixture parsing/classification tests, and a canonical negative-fixture result doc.

**Commits / heads:**

- `carbonstack-comms`: user reported the negative fixture code block pushed after `9a3f80a test: classify provider fixture events`; exact commit hash was not visible in the prompt.
- `carbonstack`: user reported `docs/37-provider-negative-fixture-result-v0.md` committed/pushed after `4dc350b docs: define provider event taxonomy`; exact commit hash was not visible in the prompt.

**New canonical doc:**

- `carbonstack/docs/37-provider-negative-fixture-result-v0.md`

**New/updated Comms paths:**

- `carbonstack-comms/internal/protocol/openmls_negative_fixture_test.go`
- `carbonstack-comms/internal/protocol/mls/research/openmls-minimal/fixtures/dev/missing-storage-error.json`
- `carbonstack-comms/internal/protocol/mls/research/openmls-minimal/fixtures/dev/missing-signer-error.json`
- `carbonstack-comms/internal/protocol/mls/research/openmls-minimal/fixtures/dev/wrong-group-error.json`
- `carbonstack-comms/internal/protocol/mls/research/openmls-minimal/fixtures/dev/malformed-message-error.json`

**Negative fixture mappings validated by Go tests:**

- `missing-storage-error.json`
  - candidate event: `storage.missing`
  - class: `storage_checkpoint`
  - severity: `warning`
  - trust relevant: `false`
  - behavior candidate: stop operation, do not send, show recovery path, do not mutate provider state further

- `missing-signer-error.json`
  - candidate event: `provider.secret.material.unavailable`
  - class: `terminal_fatal`
  - severity: `fatal`
  - trust relevant: `true`
  - behavior candidate: block send, preserve local evidence, show recovery path, require explicit identity recovery flow

- `wrong-group-error.json`
  - candidate event: `provider.group.unrecoverable`
  - class: `terminal_fatal`
  - severity: `fatal`
  - trust relevant: `true`
  - behavior candidate: stop operation, do not send, do not open message, surface provider state mismatch

- `malformed-message-error.json`
  - candidate event: `provider.message.tamper.detected`
  - class: `trust_security`
  - severity: `security`
  - trust relevant: `true`
  - behavior candidate: block message, append trust event, warn user, retain sanitized diagnostic summary

**Important encoding failure before success:**

The first negative fixture JSON files failed Go parsing with:

    invalid character 'ï' looking for beginning of value

Cause:

- Windows PowerShell `Set-Content -Encoding UTF8` wrote UTF-8 with BOM in the user's environment.
- Go `encoding/json` rejected the BOM at the start of the JSON files.

Fix:

- Rewrite the fixture JSON files as UTF-8 without BOM using `[System.Text.UTF8Encoding]::new($false)` and `[System.IO.File]::WriteAllText(...)`.

Future rule:

- For JSON fixtures parsed by Go, use UTF-8 without BOM.
- Avoid relying on `Set-Content -Encoding UTF8` for Go-parsed JSON in Windows PowerShell.

**Architectural interpretation:**

This keeps CarbonStack on the safety/rigor path:

- fixtures first
- Go parser/tests second
- taxonomy and negative-case behavior third
- trust-state mapping next
- runtime OpenMLS sidecar/CLI/Cypher integration later

**Still not validated:**

- OpenMLS runtime emission of these exact negative errors
- Comms CLI handling of provider errors
- trust-state consumption of provider events/errors
- Cypher carrying MLS/provider payloads
- production storage or secure vault behavior
- hostile-server proof, replay resistance, or metadata privacy

**Classification:** VALIDATED v0.2.11 NEGATIVE PROVIDER FIXTURE / GO CLASSIFICATION RUNG.


### Entry 71 — 2026-05-22 — Provider trust-decision mapping docs and pure Go tests validated

**Action:** Added the first provider-event-to-trust-action mapping doc and a pure Go mapping layer from provider events to candidate trust decisions.

**Commits / heads:**

- `carbonstack` `1435143 docs: map provider events to trust actions`
- `carbonstack-comms` `23b98bc test: map provider events to trust decisions`
- `carbonstack-cypher` `0bfd5af chore: remove tracked Cypher local state artifacts`
- `carbonstack-os` `b537475 Add CarbonStackOS north star and initial appliance model`

**New canonical doc:**

- `carbonstack/docs/38-provider-trust-state-mapping-v0.md`

**New Comms paths:**

- `carbonstack-comms/internal/protocol/provider_trust.go`
- `carbonstack-comms/internal/protocol/provider_trust_test.go`

**Mapping shape:**

`ProviderTrustDecision` is a pure pre-integration policy result. It includes:

- provider event name
- provider event descriptor
- candidate actions
- send/receive/open blocking booleans
- reverify requirement boolean
- user-visible boolean
- history-relevant boolean

**Candidate action vocabulary validated in code:**

- `none`
- `append_history`
- `debug_only`
- `warn_user`
- `block_send`
- `block_receive`
- `block_open`
- `quarantine_message`
- `require_reverify`
- `mark_identity_changed`
- `show_recovery_path`
- `stop_operation`
- `fatal_local_state`

**Validated by Go tests:**

- Happy-path fixture events map to append-history behavior and do not block operation.
- Invalid signature maps to block-open, warn-user, append-history, require-reverify, user-visible, and history-relevant behavior.
- Missing storage maps to stop-operation and show-recovery-path behavior; it blocks send but does not automatically block open in the pure mapping.
- Missing signer maps to fatal-local-state, block-send, show-recovery-path, user-visible, and history-relevant behavior.
- Malformed/tamper-detected message maps to block-open, quarantine-message, warn-user, append-history, user-visible, and history-relevant behavior.
- Wrong/unrecoverable group maps to fatal-local-state, stop-operation, show-recovery-path, and blocks send/receive/open.
- Unknown future provider events map to append-history/debug-only and do not automatically block operations or require reverify.

**Critical interpretation:**

This is not trust-state integration yet.

The mapping is deliberately pure. It does not:

- mutate `trust.json`
- append to `trust-events.jsonl`
- affect CLI send/inbox behavior
- call Rust/OpenMLS runtime code
- route through Cypher
- prove production E2EE

**Next safest direction:**

Do a preflight decision between:

1. one more doc/code refinement for provider trust decisions, such as fixture-backed tests that negative fixtures imply the same trust decisions; or
2. start planning the first Rust sidecar command surface while keeping runtime integration out of user-facing CLI flows.

**Classification:** VALIDATED v0.2.12 PROVIDER TRUST-DECISION PURE MAPPING RUNG.


### Entry 72 — 2026-05-22 — Phase 2C closure and Phase 2D sidecar surface plan

**Action:** Closed Phase 2C for mainline work after adding fixture-backed trust-decision tests and the Phase 2D sidecar command-surface planning doc.

**Commits / heads:**

- `carbonstack` `ddf883c docs: define Phase 2D sidecar command surface`
- `carbonstack-comms` `7d8d366 test: map OpenMLS fixtures to trust decisions`
- `carbonstack-cypher` `0bfd5af chore: remove tracked Cypher local state artifacts`
- `carbonstack-os` `b537475 Add CarbonStackOS north star and initial appliance model`

**New canonical doc:**

- `carbonstack/docs/39-phase2d-sidecar-command-surface-plan.md`
- `carbonstack/docs/40-openmls-sidecar-provider-info-result-v0.md`
- `carbonstack/docs/41-openmls-sidecar-json-envelope-plan-v0.md`
- `carbonstack/docs/42-openmls-sidecar-json-envelope-result-v0.md`

**New Comms test path:**

- `carbonstack-comms/internal/protocol/openmls_trust_decision_fixture_test.go`

**Validated by Go tests:**

- Negative fixture JSON now directly maps into `ProviderEventName` and `DecideProviderTrust(...)`.
- The known invalid-signature fixture maps to `provider.signature.invalid` and validates block-open, warn-user, append-history, require-reverify, user-visible, and history-relevant behavior.
- Missing storage validates stop-operation/show-recovery-path and send-blocking behavior.
- Missing signer validates fatal-local-state, block-send, show-recovery-path, user-visible, and history-relevant behavior.
- Wrong/unrecoverable group validates fatal-local-state, stop-operation, show-recovery-path, and blocking send/receive/open.
- Malformed provider message validates block-open, quarantine-message, warn-user, append-history, user-visible, and history-relevant behavior.

**Phase 2D command-surface decision:**

- Start runtime-provider integration as a Rust sidecar command prototype.
- Do not start with FFI.
- Do not wire sidecar output into user-facing Comms CLI commands yet.
- Do not route MLS payloads through Cypher yet.
- Do not mutate trust-state storage from provider events yet.

**Initial intended sidecar surface:**

- `provider-info`
- `identity-create`
- `public-bundle-export`
- `conversation-create`
- `conversation-add-member`
- `conversation-join`
- `message-protect`
- `message-open`
- `state-checkpoint`
- `state-load-check`

**Next safest implementation step:**

Create the first boring sidecar prototype, likely under an experimental MLS sidecar path, and implement only `provider-info` first. It should output JSON, build locally, avoid secrets, avoid storage mutation, avoid Comms/Cypher runtime integration, and be callable from a script/test.

**Classification:** VALIDATED v0.2.13 PHASE 2C CLOSURE / PHASE 2D SETUP BREAKPOINT.


## 8. Blunders / Lessons Learned

### 8.1 Gitea Web Edit Divergence

Remote web edits can cause local/remote divergence. Use fetch + pull rebase + push.

### 8.2 Copy-Box / Heredoc Formatting Hazard

Avoid backtick/fenced-code formatting in copy boxes for this user. Prefer `$Content = @' ... '@` patch style, smaller blocks, or downloadable files.

### 8.3 Wrong-Repo Commit

CarbonStackComms CLI was once committed into CarbonStackCypher. Always verify repo path before file generation and commits.

### 8.4 Test Before Growing Surface

Once a flow works, protect it before adding features.

### 8.5 Local DB / State Tracking Hazard

Generated DBs and local CLI state must remain outside source history.

### 8.6 Protocol Gravity

Do not jump into libsignal/MLS code before CarbonStack trust semantics and provider boundary are stable.

### 8.7 PowerShell Expected-Failure Harness Issue

Negative-path tests need harness logic that can distinguish expected failure from script failure.

### 8.8 Runner Path Gotcha

`validate-local.ps1` and wrapper live inside `carbonstack`; run them from the `carbonstack` repo.

### 8.9 Provider Skeleton Does Not Equal Security

The mock provider and provider-neutral interfaces are architecture scaffolding only.

### 8.10 OpenMLS Scratch Probes Do Not Equal MLS Integration

The Rust scratch crate now validates dependency/build, setup material, and Alice-side group-add behavior. It still does not mean CarbonStack uses MLS for messaging, has real E2EE, or has selected OpenMLS permanently.

### 8.11 Keep Rust Experiments Isolated

The first Rust/OpenMLS work must remain local-only and detached from Go CLI, Cypher, trust state, and production state until the minimal flow is understood.


### 8.13 Welcome Extraction Must Respect Message Wrappers

`MlsMessageOut::to_bytes()` serializes the full MLS message wrapper. Deserializing those bytes directly as raw `Welcome` caused `TrailingData`. The working scratch approach is to inspect `welcome_msg.body()` and match the `MlsMessageBodyOut::Welcome(welcome)` variant.

### 8.14 OpenMLS Provider Storage Is Device-Local State

Alice and Bob cannot share the same `OpenMlsRustCrypto` provider/storage in a two-device scratch model. Using one shared provider caused `GroupAlreadyExists` when Bob staged the Welcome because Alice's group already existed in that storage. Future CarbonStack provider design should treat provider storage as local device protocol state.


### 8.15 Application Message Processing Mutates Provider State

`process_message` required `bob_group` to be mutable. This means future provider design should treat message opening/processing as state-mutating: secrets, ratchets, epoch-related state, pending proposals, or replay-related state may change during processing. Do not model provider open/decrypt as a pure read-only helper.

### 8.16 Rust Target Artifacts Must Stay Out of Git History

Cargo `target/` artifacts accidentally entered local history/staging and caused a huge push attempt. The final safe rule is: commit only source files and lockfiles for the scratch crate, never `target/`, `.fingerprint`, `.exe`, `.pdb`, `.o`, generated docs, or build cache. If artifacts enter unpushed local history, create a backup branch, reset to `origin/main`, restore only clean source files from backup, and recommit cleanly.

### 8.17 Git Path Recovery Syntax Matters

When Git reports an ambiguous argument while resetting/unstaging a path, use `--` before the path, for example `git reset -- internal/protocol/mls/research/openmls-minimal/target`. The `--` tells Git that the remainder is a file path, not a revision.

### 8.18 Backup Branch Before Dangerous Cleanup

Before hard resets or local-history cleanup, create a backup branch such as `git branch backup/pre-target-cleanup`. This preserved the application-message work and allowed clean source restoration after the bad target-artifact commit path was discarded.

### 8.12 Pasted Local Rust Docs Are a Valid Source-of-Truth Workflow

For large Rust APIs like OpenMLS, local generated docs can be too broad to navigate live. A good workflow is to compile exact method signatures into a paste, then patch against those local signatures. This worked surprisingly well for the group-add probe and should be repeated for Welcome/StagedWelcome and message-protection work.

### 8.19 Docs Before Integration Paid Off

After the application-message scratch probe, the project deliberately paused to add `docs/29-openmls-scratch-result-v0.md` and `docs/30-openmls-provider-boundary-implications.md`. This preserved reasoning before integration pressure could distort the architecture.

### 8.20 Rust Artifact Guard Is Now a Required Pre-Push Check

`carbonstack-comms/scripts/check-no-rust-artifacts.ps1` exists to detect tracked Cargo/build artifacts. Run it before Rust scratch commits and before push hygiene checks.

### 8.21 `.gitignore` Is Not Retroactive

The guard found tracked `target/` artifacts even after ignore rules existed. Already tracked artifacts require `git rm -r --cached --ignore-unmatch -- <path>` and an intentional cleanup commit.

### 8.22 `create_message` Also Mutates Provider State

The two-message probe showed `create_message` required mutable Alice group state. Provider design must treat outbound message creation as state-mutating and persistence-relevant, not just inbound processing.

### 8.23 Two-Message State Continuity Is Not Disk Persistence

Two sequential messages inside one process are a useful continuity signal, but they do not prove provider state can survive process restart or be exported/imported safely. Real storage/reload remains future work.


### 8.24 Same-Process Provider-Storage Reload Is a Distinct Rung

`MlsGroup::load(provider.storage(), group_id)` worked for Alice and Bob inside one process. This proves OpenMLS provider storage contains usable group state at scratch level, but it still does not prove process restart recovery.

### 8.25 Direct Trait Dependencies Are Better Than Hidden Transitive Luck

The same-process reload probe needed `OpenMlsProvider` in scope to use `provider.storage()`. The scratch crate now carries a direct `openmls_traits` dependency so the API dependency is explicit.

### 8.26 Stock `OpenMlsRustCrypto` Is Not the Right File-Load Wrapper

`OpenMlsRustCrypto` contains `RustCrypto` and `MemoryStorage`, but its `key_store` field is private and no mutable storage accessor was identified. For the next scratch spike, a CarbonStack-owned scratch provider wrapper is cleaner than fighting private internals.

### 8.27 MemoryStorage Persistence Is Scratch-Useful, Not Production Storage

`MemoryStorage` exposes feature-gated file persistence methods that serialize the key/value map to JSON/base64. This is acceptable for process-restart-shaped feasibility testing, but it is not a final secure storage or vault design.


### 8.28 Restart-Continuity Needs Signer Persistence, Not Just Group Storage

The first phase-B attempt loaded Alice/Bob group state successfully but failed with `ValidationError(InvalidSignature)` because it used a fresh Alice signer. This is good behavior: Bob rejected a message signed by the wrong key. Future provider design must persist signing identity or integrate it with a secure vault/hardware-key model.

### 8.29 MemoryStorage File Persistence Is a Feasibility Win, Not a Vault

The scratch crate can save/load `MemoryStorage` files and continue messaging after fresh provider construction. However, the file format is still scratch-level JSON/base64 storage, and Alice signer JSON in OS temp is intentionally not production-safe. Treat this as proof of feasibility, not as a security design.

### 8.30 Fixture Compatibility Should Precede Direct CLI Wiring

After v0.2.8, the best next path is provider-contract/fixture shaping rather than wiring OpenMLS directly into `comms send`/`inbox`. Let Rust/OpenMLS emit stable fixture-shaped concepts first, then let Go-side provider docs/types/tests learn those shapes without committing to sidecar/FFI prematurely.


### 8.31 Fixture-First Integration Avoids Premature Coupling

The v0.2.9 fixture mode gives Go-side provider-boundary work concrete shapes to consume before Rust is linked into the CLI. This keeps integration pressure low and avoids sidecar/FFI decisions before provider events, errors, and persistence semantics are stable.

### 8.32 Sanitized Fixtures Must Stay Summary-Only

Fixture outputs may include public summaries, lengths, event names, and error shapes. They must not include MemoryStorage JSON, signer JSON, raw private keys, provider storage, or production secrets. The current fixture set follows that rule by design.

### 8.33 Invalid Signature Is a Trust Event Candidate

The known `ValidationError(InvalidSignature)` path is now preserved as a fixture. It should eventually map to a CarbonStack trust/security event, not a generic message failure. Likely actions include block, warn, append trust history, and require reverify if identity changed.



### 8.34 Go Test Working Directory Matters

The first Go-side fixture parser tests failed because the fixture path assumed repo-root execution. For `go test ./internal/protocol`, the working directory is the package directory. The correct fixture path from that package is `mls/research/openmls-minimal/fixtures/dev`.

### 8.35 Provider Event Taxonomy Should Be Executable

The taxonomy doc alone is not enough. `provider_events.go` and `provider_events_test.go` make event names/classes/severities executable and test-protected before trust-state or CLI integration begins.

### 8.36 Unknown Provider Events Should Not Become Trust-Relevant Automatically

Unknown future provider events are mapped to `unknown` with warning severity, but not automatically trust-relevant. This avoids silently escalating unreviewed provider behavior into user-facing trust policy while still making drift visible.


---


### 8.34 Negative Fixtures Make Failure Semantics Executable

The v0.2.11 negative fixture set converts likely provider failure cases into parsed/tested JSON examples. Missing storage, missing signer, wrong group ID, and malformed message now map to candidate provider events, classes, severities, trust relevance, and suggested behavior before runtime integration.

### 8.35 Go-Parsed JSON Fixtures Must Be UTF-8 Without BOM

PowerShell `Set-Content -Encoding UTF8` wrote JSON with a UTF-8 BOM in the user's Windows environment, causing Go `encoding/json` to fail with `invalid character 'ï' looking for beginning of value`. For future Go-parsed fixtures, write JSON with `[System.Text.UTF8Encoding]::new($false)` and `[System.IO.File]::WriteAllText(...)`.



### 8.23 Pure Trust Mapping Before Trust-State Mutation

`ProviderTrustDecision` intentionally maps provider events to candidate actions without mutating `trust.json`, appending trust events, or changing CLI behavior. This preserves safety: provider policy can be tested and reviewed before runtime integration.

### 8.24 Unknown Provider Events Should Not Auto-Escalate

Unknown future provider events should remain visible and history/debug relevant, but should not automatically become trust-relevant, block operations, or require reverify without explicit review.


## 9. Risk Model Updates

### Safe assumptions

- `carbonstack` owns cross-project doctrine and local validation.
- CLI-first remains the correct Comms path.
- Go + SQLite + HTTP JSON remains the first Cypher path.
- Phase 2A trust behavior exists at development-scaffold level.
- Provider boundary is documented and has a code skeleton.
- Phase 2C MLS/OpenMLS feasibility path is documented and path-reserved.
- Rust toolchain is usable enough to run scratch probes by user report.
- OpenMLS dependency/build probe exists under the reserved research path.
- OpenMLS credential/KeyPackage API probe exists.
- OpenMLS Alice-side group-add probe exists.
- OpenMLS Bob join-from-Welcome probe exists.
- OpenMLS application-message protect/open probe exists at Rust-only scratch level.
- OpenMLS two-message in-process state-continuity probe exists at Rust-only scratch level.
- OpenMLS scratch result and provider-boundary implications docs exist.
- Rust artifact guard exists and should be used before scratch commits.
- Pasted local docs/method signatures are a safe workflow for grounding Rust API patches.
- OpenMLS provider/storage should be modeled as device-local state.
- Scratch-only MemoryStorage file persistence and signer reload are validated at feasibility level.
- Dev-only sanitized OpenMLS provider fixtures are validated at provider-contract planning level.
- Summary-only negative provider fixtures are validated at provider-contract planning level.
- Go-side fixture parsing and provider event classification are validated at provider-contract planning level.
- Provider event taxonomy is documented and represented in Go types/tests.
- Real CarbonStack cryptographic integration remains future work.

### Unsafe assumptions

- That stub/base64 payloads are encryption.
- That mock-provider payloads are encryption.
- That fake/dev fingerprints are real identity verification.
- That README-only MLS slots mean OpenMLS is integrated.
- That the application-message scratch probe means CarbonStackComms uses MLS for messaging.
- That the application-message or two-message scratch probe means production E2EE, hostile-server security, replay resistance, disk persistence, process-restart recovery, or trust-state integration.
- That passing local tests means hostile-server security.
- That provider interface code implies final protocol selection.
- That OpenMLS will definitely fit.
- That mls-rs is rejected.
- That AGPL/libsignal implications are solved.
- That the project is safe for real secrets.
- That MemoryStorage JSON/base64 files or temp signer JSON are production-safe storage.
- That v0.2.8 persistence result means Comms/Cypher use MLS.
- That v0.2.9/v0.2.11 fixtures are production provider behavior or secure messaging.
- That v0.2.10 provider event classification means trust-state consumes provider events.
- That event taxonomy is final or sufficient for hostile-server security.

---

## 10. Validation Ladder

| Rung | Description | Status |
|---|---|---|
| 1 | Repo family initialized and pushed | VALIDATED |
| 2 | Base doctrine and canonical spec captured | VALIDATED |
| 3 | Phase 1 integration/data/API/client docs captured | VALIDATED |
| 4 | CarbonStackCypher Go/SQLite scaffolding committed | VALIDATED |
| 5 | CarbonStackCypher local server starts | VALIDATED |
| 6 | Invite/account/device/envelope lifecycle works | VALIDATED |
| 7 | CarbonStackComms CLI-driven two-client lifecycle works | VALIDATED |
| 8 | Cypher API lifecycle tests pass | VALIDATED |
| 9 | Comms package tests pass | VALIDATED |
| 10 | Local validation runner passes | VALIDATED |
| 11 | Phase 1 hygiene cleanup complete | VALIDATED |
| 12 | Phase 2 protocol docs through feasibility matrix | VALIDATED |
| 13 | Phase 2A trust-state scaffold implemented | VALIDATED |
| 14 | Phase 2A trust lifecycle script passes | VALIDATED |
| 15 | `trust-list` and reverify lifecycle covered | VALIDATED |
| 16 | Protocol provider boundary documented | VALIDATED |
| 17 | Provider-neutral protocol skeleton | VALIDATED |
| 18 | Mock provider package tests | VALIDATED |
| 19 | MLS feasibility spike plan | VALIDATED DOC |
| 20 | MLS implementation candidate notes | VALIDATED DOC |
| 21 | Experimental MLS provider slot reserved | VALIDATED DOC |
| 22 | OpenMLS minimal example plan | VALIDATED DOC |
| 23 | OpenMLS upstream example notes | VALIDATED DOC |
| 24 | OpenMLS minimal scratch experiment path reserved | VALIDATED DOC |
| 25 | Rust toolchain installed/usable | VALIDATED BY USER REPORT |
| 26 | OpenMLS dependency probe crate added | VALIDATED BUILD/DEPENDENCY PROBE |
| 27 | OpenMLS credential/KeyPackage API probe | VALIDATED API PROBE |
| 28 | OpenMLS Alice group creation + add Bob from KeyPackage | VALIDATED GROUP-ADD PROBE |
| 29 | Bob join from Welcome/StagedWelcome | VALIDATED WELCOME-JOIN PROBE |
| 30 | MLS application message protect/open | VALIDATED SCRATCH PROBE |
| 31 | OpenMLS scratch result doc | VALIDATED DOC |
| 32 | OpenMLS provider-boundary implications doc | VALIDATED DOC |
| 33 | Rust artifact guard / tracked target cleanup | VALIDATED HYGIENE |
| 34 | OpenMLS two-message in-process state continuity | VALIDATED SCRATCH PROBE |
| 35 | OpenMLS same-process provider-storage reload | VALIDATED SCRATCH PROBE |
| 36 | OpenMLS MemoryStorage file-persistence/process-restart-shaped spike | VALIDATED SCRATCH PROBE |
| 37 | Production-safe storage / secure vault design | NOT VALIDATED |
| 38 | Provider-contract / fixture plan | VALIDATED DOC |
| 39 | Dev-only OpenMLS provider fixture summaries | VALIDATED SCRATCH/FIXTURE RUNG |
| 40 | Go-side provider-contract fixture review/tests | VALIDATED GO TEST RUNG |
| 41 | Provider event taxonomy doc | VALIDATED DOC |
| 42 | Provider event names/classes/severities in Go | VALIDATED GO TEST RUNG |
| 43 | Trust-state mapping from provider events | NOT VALIDATED |
| 44 | Negative-path provider fixtures | NOT VALIDATED |
| 45 | Real cryptographic protocol integration | NOT VALIDATED |
| 46 | Hostile-server test harness | NOT VALIDATED |
| 47 | CarbonStackOS appliance prototype | DEFERRED |

---

## 11. Next Version TODO

Next target: **Phase 2D sidecar provider-info prototype**, still without wiring OpenMLS into Comms/Cypher runtime behavior.

### Immediate TODO candidates

1. Create a new experimental sidecar location, likely one of:
   - `carbonstack-comms/internal/protocol/mls/research/openmls-sidecar`
   - `carbonstack-comms/internal/protocol/mls/sidecar/openmls-provider`
2. Implement only a boring `provider-info` command first.
   - output JSON
   - include provider name/mode/capabilities
   - no secrets
   - no provider storage mutation
   - no Comms CLI integration
   - no Cypher routing
3. Add a script/test that can call the sidecar and parse its JSON output.
4. Keep the existing scratch crate and fixture tests intact.
5. Continue hygiene:
   - run `check-no-rust-artifacts.ps1` before Rust commits
   - do not commit temp signer/provider-state JSON files
   - keep generated Cargo artifacts out of Git
   - keep fixture files summary-only until maturity requires protected payload fixtures
6. After `provider-info`, consider the next boring sidecar command: `identity-create` or a fixture-equivalent command, but only after the first command contract is test-protected.

### Do not do next

- Do not wire OpenMLS into `comms send` or `comms inbox`.
- Do not route MLS payloads through Cypher yet.
- Do not integrate libsignal.
- Do not implement MLS from scratch.
- Do not implement production app-level crypto.
- Do not start Android.
- Do not start OS work.
- Do not add user-facing groups, attachments, or media.
- Do not claim production security.

## 12. Open Blockers / TODO

- Welcome extraction path validated: use `welcome_msg.body()` and match `MlsMessageBodyOut::Welcome(welcome)` rather than deserializing full `MlsMessageOut` bytes as raw `Welcome`.
- Scratch-only MemoryStorage file-persistence is validated; production-safe storage/secure vault design is not.
- Provider boundary mapping from OpenMLS is documented at implications/taxonomy level and partially represented in Go event descriptors, but it is not wired into real provider runtime behavior.
- Provider skeleton is not wired into CLI send/inbox path.
- AGPL/libsignal consequences remain unresolved, though libsignal is not in mainline.
- Negative-path provider fixtures and hostile-server tamper tests do not exist.
- Local secure vault does not exist.
- Hardware-key flow does not exist.
- Android client does not exist.
- CarbonStackOS is deferred.
- CI pipeline does not exist.

---

## 13. Known-Good Commands

Status check:

    cd C:\▮▮
    Get-ChildItem -Directory | ForEach-Object { Write-Host ""; Write-Host "=== $($_.Name) ==="; git -C $_.FullName status --short; git -C $_.FullName log --oneline -3 }

Run local validation:

    cd C:\▮▮\carbonstack
    powershell -ExecutionPolicy Bypass -File .\scripts\validate-local.ps1

Run Cypher tests directly:

    cd C:\▮▮\carbonstack-cypher
    go test ./...

Run Comms tests directly:

    cd C:\▮▮\carbonstack-comms
    go test ./...

Run OpenMLS scratch crate directly:

    cd C:\▮▮\carbonstack-comms\internal\protocol\mls\research\openmls-minimal
    cargo check
    cargo run
    cargo run -- phase-a
    cargo run -- phase-b

Generate local OpenMLS docs for pasted-docs workflow:

    cd C:\▮▮\carbonstack-comms\internal\protocol\mls\research\openmls-minimal
    cargo doc --open -p openmls
    cargo doc --open -p openmls_memory_storage

Use the pasted-docs workflow:

    Search local docs for the exact struct/method.
    Copy only the relevant signatures and summaries into a paste.
    Patch against those local signatures.
    Stop on compiler errors.

Standard safe sync flow:

    git fetch origin
    git pull --rebase origin main
    git push

Rust checks:

    rustc --version
    cargo --version
    rustup --version

Git hygiene for Rust scratch crate:

    cd C:\▮▮\carbonstack-comms
    git diff --cached --name-only
    git show --stat --oneline HEAD
    git ls-tree -r origin/main -- internal/protocol/mls/research/openmls-minimal/target

Recovery pattern if Cargo `target/` enters local unpushed history:

    git branch backup/pre-target-cleanup
    git reset --hard origin/main
    git restore --source backup/pre-target-cleanup -- .gitignore
    git restore --source backup/pre-target-cleanup -- internal/protocol/mls/research/openmls-minimal/Cargo.toml
    git restore --source backup/pre-target-cleanup -- internal/protocol/mls/research/openmls-minimal/Cargo.lock
    git restore --source backup/pre-target-cleanup -- internal/protocol/mls/research/openmls-minimal/README.md
    git restore --source backup/pre-target-cleanup -- internal/protocol/mls/research/openmls-minimal/src/main.rs


Run Rust artifact guard:

    cd C:\▮▮\carbonstack-comms
    powershell -ExecutionPolicy Bypass -File .\scripts\check-no-rust-artifacts.ps1

Remove tracked Rust target artifacts if guard fails:

    cd C:\▮▮\carbonstack-comms
    git rm -r --cached --ignore-unmatch -- internal/protocol/mls/research/openmls-minimal/target

---

## 14. Dangerous Commands

Avoid unless intentionally recovering with backups:

    git push --force
    git reset --hard origin/main
    git clean -fdx

Never commit:

- private keys
- signing keys
- `.env` / `.env.*`
- secrets
- hardware-key recovery material
- unredacted personal secrets
- build signing material
- private device credentials
- local production databases
- real recovery material
- local CLI state files unless intentionally sanitized as fixtures
- real private MLS key material
- generated secret/provider state from MLS experiments unless intentionally sanitized as test fixtures

---

## 15. Allowed Claims

After v0.2.5, allowed claims:

- CarbonStack has a coherent doctrine and specification foundation.
- Phase 1 relay/client skeleton works locally.
- CarbonStackCypher has passing API lifecycle tests.
- CarbonStackComms has passing package and lifecycle tests.
- CarbonStackComms has Phase 2A trust-state scaffolding.
- Local validation runner passes.
- CarbonStack has documented an MLS-shaped, provider-neutral protocol boundary.
- CarbonStackComms has a provider-neutral protocol skeleton and mock provider package tests.
- CarbonStack has an MLS feasibility spike plan.
- CarbonStack has OpenMLS implementation candidate notes and upstream example notes.
- CarbonStackComms has reserved an experimental MLS provider slot and OpenMLS minimal scratch experiment path.
- OpenMLS is the first intended spike candidate.
- mls-rs remains a serious alternate candidate.
- CarbonStackComms has a Rust-only OpenMLS dependency/build probe crate.
- CarbonStackComms can create OpenMLS credential/KeyPackage setup material locally in the scratch crate.
- CarbonStackComms can create an Alice MLS group and add Bob from a KeyPackage locally in the scratch crate.
- CarbonStackComms can complete a Rust-only Bob join-from-Welcome flow in the scratch crate.
- CarbonStackComms can complete a Rust-only Alice-to-Bob OpenMLS application-message protect/open flow in the scratch crate.
- CarbonStackComms can complete two sequential Alice-to-Bob OpenMLS application messages inside one process in the scratch crate.
- CarbonStack has canonical OpenMLS scratch-result and provider-boundary implications docs.
- CarbonStackComms has a Rust artifact guard script.
- No final cryptographic protocol has been selected or integrated.

## 16. Not Allowed Claims

After v0.2.5, do not claim:

- CarbonStack is production secure.
- CarbonStack is audited.
- CarbonStack is Signal-equivalent.
- Stub/base64, mock-provider payloads, README-only MLS slots, dependency-probe code, or fake/dev fingerprints provide real encryption or real identity verification.
- MLS/OpenMLS has been integrated into CarbonStackComms messaging.
- mls-rs has been integrated.
- libsignal has been integrated.
- A final protocol has been selected.
- OpenMLS application-message scratch work means CarbonStackComms messaging uses MLS.
- OpenMLS scratch group-add, Welcome-join, application-message, or two-message state-continuity work means production E2EE.
- OpenMLS two-message state continuity proves disk persistence or process restart recovery.
- Metadata privacy, replay resistance, or hostile-server security is solved.
- Hardware-key flows, Android, or CarbonStackOS are implemented.
- A CI pipeline exists.
- Provider skeleton equals real E2EE.

---

## 17. Breakpoint

### Breakpoint: 2026-05-22 — OpenMLS Docs, Hygiene Guard, and Two-Message State Continuity v0.2.6

**Current state:**

CarbonStack has advanced from the v0.2.5 application-message probe into a documented and guarded Phase 2C OpenMLS scratch state. The `carbonstack` repo now contains canonical scratch-result and provider-boundary implications docs. The `carbonstack-comms` repo now contains a Rust artifact guard and has removed tracked Cargo `target/` artifacts. The OpenMLS scratch crate can send two sequential Alice-to-Bob application messages inside one process and open both with matching plaintext.

This is still not provider integration, not disk persistence, not process-restart recovery, not Cypher routing, not trust-state integration, and not production security.

**Validated:**

- `docs/29-openmls-scratch-result-v0.md` added at `carbonstack` `6f10d2b`.
- `docs/30-openmls-provider-boundary-implications.md` added at `carbonstack` `440131f`.
- Rust artifact guard added at `carbonstack-comms` `7cfb590`.
- Tracked Cargo `target/` artifacts removed from `carbonstack-comms`.
- OpenMLS two-message state-continuity probe committed at `carbonstack-comms` `8be5e36`.
- Two sequential Alice-to-Bob OpenMLS application messages work inside one process.
- `create_message` required mutable Alice group state.
- `process_message` required mutable Bob group state.
- Provider outbound and inbound operations must both be considered state-mutating and persistence-relevant.
- Repo heads in user snapshot:
  - `carbonstack`: `440131f docs: record OpenMLS provider boundary implications`
  - `carbonstack-comms`: `8be5e36 test: probe OpenMLS two-message state continuity`
  - `carbonstack-cypher`: `0bfd5af chore: remove tracked Cypher local state artifacts`
  - `carbonstack-os`: `b537475 Add CarbonStackOS north star and initial appliance model`

**Not validated:**

- real disk persistence
- process restart recovery
- provider state export/import
- provider mapping from OpenMLS to CarbonStack code
- Go/Rust integration
- CarbonStackComms CLI integration
- CarbonStackCypher routing of MLS payloads
- trust-state integration
- real crypto/security properties in CarbonStackComms
- final protocol selection
- production auth
- metadata privacy
- replay resistance
- hostile-server test harness
- hardware-key flows
- local secure vault
- Android client
- OS/device work
- CI pipeline

**Next safest action:**

Run a full repo docs + hygiene review, then design the real OpenMLS provider storage/export/reload investigation. Do not wire OpenMLS into Comms/Cypher yet.

**Do NOT do next:**

Do not wire OpenMLS into `comms send` or `comms inbox`. Do not integrate libsignal. Do not implement MLS from scratch. Do not implement production app-level crypto. Do not start Android or OS. Do not add user-facing groups, attachments, or media. Do not claim production security.

**Safe to pause here:** Yes

---


## v0.2.16 Next ToDos

Immediate next action:

- Decide whether `provider.command.unsupported` is now a stable provider event and, if yes, add it to the Go provider event taxonomy/mapping.

Recommended next safest sequence:

1. Update `carbonstack-comms/internal/protocol/provider_events.go` with `ProviderEventCommandUnsupported` if stable.
2. Add tests that classify `provider.command.unsupported` as warning/non-trust-relevant operational behavior.
3. Update provider trust mapping only if needed so unsupported commands remain non-blocking/non-reverify and do not become trust failures.
4. Add a short canonical doc/update if the event vocabulary changes materially.
5. Plan `identity-create` before implementation, including storage location, output envelope, public-only fields, no secret stdout, and checkpoint/event behavior.
6. Keep Comms CLI, Cypher routing, and trust-store mutation out of scope until sidecar command contracts are stable and tested.

Do not do next:

- Do not wire OpenMLS into `comms send` / `comms inbox`.
- Do not route MLS payloads through Cypher.
- Do not mutate trust-state storage from provider events.
- Do not add production vault/key storage claims.
- Do not jump directly to `message-protect` / `message-open`.
- Do not treat sidecar envelope parsing as production provider integration.
- Do not start Android/OS work.

## Appendix A: Versioning Schema

Use `v[MAJOR].[MINOR]` where:

- **MAJOR** = project scope / identity change.
- **MINOR** = timeline continuity within that scope.

This file uses `v0.2.16` because Phase 2D now has a minimally test-protected sidecar success/error envelope: `provider-info` returns an `ok: true` JSON envelope, unsupported commands return an `ok: false` JSON error envelope plus exit code 2, and both paths remain non-secret and detached from Comms/Cypher/trust runtime integration.

---

## Appendix B: When to Update This LogDoc

Update after any of these:

- full repo docs + hygiene review completed
- OpenMLS provider persistence/export/reload investigation succeeds/fails
- sidecar command surface, JSON envelopes, or provider-info tests change
- OpenMLS scratch result doc added
- provider-boundary mapping changes
- OpenMLS scratch result docs added
- group creation/add probe changes
- MLS message protect/open attempted
- provider-boundary changes
- provider skeleton gets wired into CLI
- protocol feasibility decision changes
- security claim validated or invalidated
- blocker appears
- validation pass or failure
- OS/device work begins
- pause/handoff point reached


---

---

## v0.2.21 Breakpoint Notes

**Current safe-to-pause state:** stable after `docs/48` and the identity prep state event taxonomy/trust-mapping rung landed.

**Next safest action:** implement minimal dev-only real `identity-create` material generation under the ignored sidecar dev-state path, while keeping stdout sanitized and no Comms/Cypher/trust-store integration.

**Then:** harden real identity duplicate/refused-overwrite behavior, then plan `public-bundle-export`.

**Do not do next:**

- Do not implement `public-bundle-export` before real identity state generation and overwrite refusal are stable.
- Do not jump to conversation lifecycle or message protect/open.
- Do not wire the sidecar into `comms send` or `comms inbox`.
- Do not route MLS payloads through Cypher.
- Do not mutate trust-state storage from provider events.
- Do not print or commit signer JSON, MemoryStorage JSON, provider storage, private keys, recovery material, target artifacts, `.exe`, or `.pdb`.

