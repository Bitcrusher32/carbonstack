# OpenMLS Sidecar Conversation Create Persistence Repair Result v0

Status: Validated repair checkpoint
Component: CarbonStackComms / OpenMLS sidecar
Phase: Phase 2D
Plan doc:
- docs/63-openmls-sidecar-conversation-create-persistence-repair-plan-v0.md
Related docs:
- docs/60-openmls-sidecar-conversation-create-result-v0.md
- docs/61-openmls-sidecar-conversation-add-member-welcome-plan-v0.md
- docs/62-openmls-sidecar-add-member-welcome-api-recon-v0.md

## 1. Summary

This checkpoint repairs an overclaim in the v0.2.27 conversation-create result.

v0.2.27 correctly created a one-member OpenMLS group in-process and wrote sanitized conversation metadata, but it reported:

    provider_storage_written=true

v0.2.28 recon and v0.2.29 inspection showed that this was too strong.

OpenMlsRustCrypto uses an in-memory MemoryStorage by default. A new sidecar process gets a fresh provider/storage instance. The current conversation-create command writes only sanitized conversation-summary.json under the conversation directory and does not persist reloadable OpenMLS group/provider state.

This checkpoint corrects the claim.

## 2. Corrected behavior

conversation-create now reports:

    provider_storage_written=false
    group_reloadable=false

This appears in:

- conversation-create stdout;
- conversation-summary.json;
- Go-side contract tests.

The warning now explicitly states that the group is not reloadable across sidecar process invocations yet.

## 3. Why this repair was required

conversation-add-member requires a mutable MlsGroup.

OpenMLS MlsGroup::load expects persisted storage state including:

- PublicGroup;
- group epoch secrets;
- own leaf index;
- message secrets;
- resumption PSK store;
- join config;
- own leaf nodes;
- group state.

The current sidecar does not persist those fields.

Therefore, add-member / Welcome export must remain blocked until a future checkpoint implements real dev-local provider/group persistence or another honest reload strategy.

## 4. Validated manual behavior

Manual probe:

    Remove-Item -Recurse -Force .\.carbonstack-openmls-sidecar-state -ErrorAction SilentlyContinue
    cargo run -- identity-create --device-label carbonstack-alice-device
    cargo run -- conversation-create --device-label carbonstack-alice-device --conversation-label carbonstack-test-conversation
    Get-Content .\.carbonstack-openmls-sidecar-state\dev\conversations\carbonstack-test-conversation\conversation-summary.json

Validated result:

- identity-create succeeds;
- conversation-create succeeds;
- stdout reports provider_storage_written=false;
- stdout reports group_reloadable=false;
- conversation-summary.json reports provider_storage_written=false;
- conversation-summary.json reports group_reloadable=false;
- private_material_included=false;
- generated signer material remains local and unprinted.

## 5. Tests and guard

Validated:

- cargo check;
- go test ./internal/protocol -run TestOpenMLSSidecarConversationCreate;
- go test ./internal/protocol;
- go test ./...;
- scripts/check-no-rust-artifacts.ps1.

Rust artifact guard passed with no tracked Rust/build artifacts.

## 6. Files changed

CarbonStackComms changed:

    internal/protocol/mls/research/openmls-sidecar/src/main.rs
    internal/protocol/mls/research/openmls-sidecar/src/state.rs
    internal/protocol/openmls_sidecar_provider_info_test.go

CarbonStack docs changed:

    docs/63-openmls-sidecar-conversation-create-persistence-repair-plan-v0.md
    docs/64-openmls-sidecar-conversation-create-persistence-repair-result-v0.md

## 7. Allowed claims after this checkpoint

Allowed:

- conversation-create creates a dev-local one-member OpenMLS group in-process.
- conversation-create writes sanitized conversation-summary.json.
- conversation-create reports provider_storage_written=false.
- conversation-create reports group_reloadable=false.
- Go tests protect against reintroducing the provider-storage overclaim.
- add-member remains blocked on real provider/group persistence.

Not allowed:

- conversation-create persists reloadable OpenMLS group state.
- conversation-create can be used as a basis for add-members in a later sidecar command.
- provider storage persistence is solved.
- conversation-add-member exists.
- Welcome export exists.
- message protect/open exists.
- production E2EE exists.

## 8. Next recommended checkpoint

Next safest checkpoint:

    OpenMLS sidecar dev provider/group persistence plan + implementation

Possible doc:

    docs/65-openmls-sidecar-dev-provider-group-persistence-plan-v0.md

The next checkpoint should decide how to persist enough OpenMLS group/provider state so that:

    MlsGroup::load(provider.storage(), &group_id)

can succeed across sidecar invocations.

Only after that succeeds should CarbonStack return to:

    conversation-add-member / Welcome export implementation
