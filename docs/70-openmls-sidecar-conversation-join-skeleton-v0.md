# OpenMLS Sidecar Conversation-Join Skeleton v0

Status: Skeleton / pre-implementation contract
Component: CarbonStackComms / OpenMLS sidecar
Phase: Phase 2D
Previous docs:
- docs/67-openmls-sidecar-add-member-welcome-skeleton-v0.md
- docs/68-openmls-sidecar-add-member-welcome-api-recon-v0.md
- docs/69-openmls-sidecar-add-member-welcome-export-result-v0.md

## 1. Purpose

This document defines the next command surface after v0.2.32.

v0.2.32 validated:

- Alice can create a persisted dev-local OpenMLS group;
- Alice can add Bob using Bob's public KeyPackage artifact;
- Alice can export a Welcome carrier artifact as welcome.bin;
- Alice can merge the pending commit;
- Alice can save mutated provider-storage.json;
- a later load-check proves Alice's group reloads with member_count=2.

The next lifecycle step is Bob-side join.

This document is intentionally a skeleton and planning contract. It does not implement conversation-join.

## 2. New command target

Proposed command:

    conversation-join --device-label <joining-device> --conversation-label <conversation> --welcome <path>

Initial v0 target:

    conversation-join --device-label carbonstack-bob-device --conversation-label carbonstack-test-conversation --welcome <welcome.bin>

Meaning:

- joining-device is the local Bob device that already has sidecar identity/signing state;
- conversation-label is the desired local label for Bob's joined group state;
- welcome is the Welcome carrier artifact exported by Alice's conversation-add-member command.

Deferred input style:

    --welcome-manifest <path>

The manifest-driven flow is deferred. The first implementation should accept a direct Welcome artifact path because v0.2.32 writes welcome.bin directly and this keeps v0.2.34 bounded.

## 3. Required preconditions

Before conversation-join can succeed:

1. Bob identity state exists;
2. Bob signer state is loadable;
3. Welcome artifact path exists;
4. Welcome artifact path points to a file;
5. Welcome artifact size is sane;
6. Welcome artifact deserializes into the OpenMLS type required for join;
7. Bob provider storage can be initialized/saved;
8. Bob joined group can be created from the Welcome;
9. Bob joined group can be saved to provider-storage.json;
10. Bob joined group can be reloaded by a later load-check.

## 4. Critical design issue: state layout

Current conversation state layout is:

    .carbonstack-openmls-sidecar-state/dev/conversations/<conversation-label>/

This was enough for Alice-only group state. It becomes risky for Bob join because Alice and Bob run under the same dev sidecar root in local tests.

Potential collision:

    Alice provider storage:
    dev/conversations/carbonstack-test-conversation/provider-storage.json

    Bob provider storage:
    dev/conversations/carbonstack-test-conversation/provider-storage.json

That means Bob join could overwrite Alice's provider-storage.json in local dev tests.

Recommended v0.2.33 decision:

For joined-side implementation, use device-scoped conversation state.

Recommended future layout:

    .carbonstack-openmls-sidecar-state/dev/devices/<device-label>/conversations/<conversation-label>/

Device-scoped files:

    conversation-summary.json
    provider-storage.json

For v0.2.34, Bob join should write Bob's joined group state under:

    dev/devices/carbonstack-bob-device/conversations/carbonstack-test-conversation/

Alice's existing v0.2.32 state can remain under the old conversation-global path for now, but this creates a transition state.

Longer-term cleanup should migrate Alice conversation-create to the same device-scoped layout.

## 5. Proposed operation sequence

Recommended sequence:

1. validate device label;
2. validate conversation label;
3. parse --welcome path;
4. validate Welcome path;
5. reject unsafe path targets such as signer.json, provider-storage.json, raw state files, or directories;
6. require existing Bob identity state;
7. load Bob signer state if needed by join path;
8. create fresh CarbonStackSidecarProvider for Bob;
9. read welcome.bin bytes;
10. deserialize Welcome carrier;
11. extract or obtain inner Welcome as required by StagedWelcome::new_from_welcome;
12. build join config;
13. call StagedWelcome::new_from_welcome(...);
14. call into_group(...);
15. save Bob provider-storage.json;
16. immediately reload Bob group from saved provider storage if possible;
17. write Bob joined conversation-summary.json;
18. return sanitized stdout.

## 6. Welcome artifact question

v0.2.32 currently writes welcome.bin as:

    second MlsMessageOut returned by add_members
    serialized with MlsMessageOut::to_bytes()

v0.2.33 recon must answer:

1. Can conversation-join deserialize welcome.bin directly as an MlsMessageIn or MlsMessageOut carrier?
2. Is there a public conversion path from carrier to inner Welcome?
3. Does StagedWelcome::new_from_welcome require inner Welcome?
4. If it requires inner Welcome, can the inner Welcome be extracted without test-utils-only helpers?
5. If not, should v0.2.32/v0.2.34 change the artifact to inner Welcome bytes instead?

Do not implement join until this is answered.

## 7. Proposed generated files

Under recommended Bob device-scoped path:

    .carbonstack-openmls-sidecar-state/dev/devices/<device-label>/conversations/<conversation-label>/

Expected files:

    conversation-summary.json
    provider-storage.json

Potential additional metadata:

    join-summary.json
    welcome-consume-summary.json

Recommended v0:

    conversation-summary.json
    provider-storage.json

Optional:

    join-summary.json

If join-summary.json is added, keep it sanitized.

## 8. Proposed stdout fields

Success envelope:

    ok=true
    command=conversation-join
    provider=openmls
    implementation=carbonstack-openmls-sidecar
    mode=experimental-sidecar
    phase=phase2d-conversation-join-dev
    private_material_included=false

Data fields:

    device_label
    conversation_label
    conversation_state_path_hint
    conversation_summary_path_hint
    provider_storage_path_hint
    welcome_artifact_path_hint
    provider_storage_written=true
    group_reloadable=true
    joined=true
    member_count
    epoch
    group_id_ref
    group_id_len
    state_scope=dev-local-sidecar-state

Events:

    conversation.joined
    storage.saved

Trust:

- dev-sidecar metadata only for now;
- joining a group is product-level trust sensitive later;
- do not mutate trust state in this rung.

## 9. Proposed failure classes

Missing argument:

    missing_required_argument

Invalid labels:

    invalid_device_label
    invalid_conversation_label

Invalid Welcome path:

    invalid_welcome_path

Missing input state:

    identity_missing
    welcome_missing

Unavailable/corrupt state:

    secret_material_unavailable
    welcome_invalid
    conversation_join_failed
    provider_storage_save_failed
    provider_group_unavailable

Duplicate/refuse overwrite:

    conversation_already_exists

Generic failure:

    conversation_join_failed

## 10. Duplicate behavior

Initial implementation should refuse overwrite if Bob's device-scoped joined conversation state already exists:

    dev/devices/<device-label>/conversations/<conversation-label>/

Do not silently overwrite Bob's joined group state.

## 11. Required tests for implementation

Go-side contract tests should cover:

- provider-info lists conversation-join as supported after implementation;
- missing --device-label fails;
- missing --conversation-label fails;
- missing --welcome fails;
- invalid device label fails;
- invalid conversation label fails;
- missing identity fails;
- missing Welcome file fails;
- invalid Welcome artifact fails;
- duplicate join refuses overwrite;
- success flow:
  - remove state;
  - identity-create Alice;
  - identity-create Bob;
  - public-bundle-export Bob --write-artifact;
  - conversation-create Alice;
  - conversation-add-member Alice with Bob KeyPackage artifact;
  - conversation-join Bob with welcome.bin;
- success asserts:
  - ok=true;
  - joined=true;
  - provider_storage_written=true;
  - group_reloadable=true;
  - member_count=2;
  - provider-storage.json exists under Bob device-scoped conversation state;
  - conversation-summary.json exists under Bob device-scoped conversation state;
  - stdout contains no private material.

## 12. Manual probe target after implementation

Expected future probe:

    Remove-Item -Recurse -Force .\.carbonstack-openmls-sidecar-state -ErrorAction SilentlyContinue

    cargo run -- identity-create --device-label carbonstack-alice-device

    cargo run -- identity-create --device-label carbonstack-bob-device

    cargo run -- public-bundle-export --device-label carbonstack-bob-device --write-artifact

    cargo run -- conversation-create --device-label carbonstack-alice-device --conversation-label carbonstack-test-conversation

    cargo run -- conversation-add-member --device-label carbonstack-alice-device --conversation-label carbonstack-test-conversation --member-keypackage .\.carbonstack-openmls-sidecar-state\dev\devices\carbonstack-bob-device\public-bundle.keypackage.bin

    cargo run -- conversation-join --device-label carbonstack-bob-device --conversation-label carbonstack-test-conversation --welcome .\.carbonstack-openmls-sidecar-state\dev\conversations\carbonstack-test-conversation\welcome.bin

Safe inspection only:

    Get-ChildItem .\.carbonstack-openmls-sidecar-state\dev\devices\carbonstack-bob-device\conversations\carbonstack-test-conversation | Select-Object Name, Length

Do not inspect:

    signer.json
    provider-storage.json
    welcome.bin
    raw KeyPackage bytes
    raw group state
    raw private material

## 13. Result doc requirements

The result doc must explicitly record:

- exact command surface implemented;
- exact Welcome artifact input type consumed;
- exact Welcome deserialization path;
- whether welcome.bin remained outer MlsMessageOut bytes or changed;
- exact StagedWelcome::new_from_welcome path;
- exact into_group path;
- exact Bob-side state layout;
- exact artifact paths written;
- tests run;
- artifact guard result;
- allowed claims;
- not allowed claims.

## 14. Next after join

After conversation-join succeeds:

    message-protect/message-open skeleton/recon

Then:

    message-protect implementation
    message-open implementation

Do not combine join and message protect/open into one checkpoint.
