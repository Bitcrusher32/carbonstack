# OpenMLS Sidecar Add-Member / Welcome Export Skeleton v0

Status: Skeleton / pre-implementation contract
Component: CarbonStackComms / OpenMLS sidecar
Phase: Phase 2D
Previous docs:
- docs/61-openmls-sidecar-conversation-add-member-welcome-plan-v0.md
- docs/62-openmls-sidecar-add-member-welcome-api-recon-v0.md
- docs/65-openmls-sidecar-dev-provider-group-persistence-plan-v0.md
- docs/66-openmls-sidecar-dev-provider-group-persistence-result-v0.md

## 1. Purpose

This document defines the next command surface after v0.2.30.

v0.2.30 validated dev-local provider/group persistence:

- conversation-create writes provider-storage.json;
- conversation-create reports provider_storage_written=true;
- conversation-create reports group_reloadable=true;
- conversation-load-check reloads the group in a later sidecar invocation;
- conversation-load-check reports provider_storage_loaded=true and group_reloadable=true.

That removes the v0.2.29 blocker. The next lifecycle step is to add another member to a persisted OpenMLS group and export a Welcome artifact.

This document is intentionally a skeleton and API contract. It does not implement add-member.

## 2. New command target

Proposed command:

    conversation-add-member --device-label <creator-device> --conversation-label <conversation> --member-keypackage <path>

Meaning:

- creator-device is the existing local device that owns the current persisted group state;
- conversation is the existing local conversation label;
- member-keypackage is a path to a serialized public OpenMLS KeyPackage artifact for the new member.

Initial v0 target:

    conversation-add-member --device-label carbonstack-alice-device --conversation-label carbonstack-test-conversation --member-keypackage <bob-public-bundle.keypackage.bin>

Deferred input style:

    --member-public-bundle-manifest <path>

The manifest-driven flow is deferred. The first implementation should accept the direct KeyPackage artifact path because it matches the current v0.2.25 artifact exporter and is easier to test.

## 3. Required preconditions

Before conversation-add-member can succeed:

1. creator identity state exists;
2. creator signer state is loadable;
3. conversation-summary.json exists;
4. provider-storage.json exists;
5. MlsGroup::load succeeds for the deterministic conversation GroupId;
6. member-keypackage path exists;
7. member-keypackage is not a directory;
8. member-keypackage size is sane;
9. member-keypackage deserializes as a public OpenMLS KeyPackage or KeyPackageIn path chosen by implementation;
10. no output artifact for the same add-member attempt is being overwritten silently.

## 4. Non-goals

This checkpoint and first implementation must not:

- implement Bob-side conversation-join;
- implement message-protect;
- implement message-open;
- wire OpenMLS into Comms runtime send/inbox;
- route MLS payloads through Cypher;
- mutate trust.json or trust-events.jsonl;
- implement production secure vault storage;
- implement hardware-backed identity;
- implement Android or CarbonStackOS;
- claim production E2EE or Signal-equivalent security.

## 5. Proposed operation sequence

Recommended sequence:

1. validate device label;
2. validate conversation label;
3. validate member-keypackage path;
4. reject unsafe path targets such as signer.json, provider-storage.json, raw state files, or directories;
5. require existing creator identity state;
6. load creator signer state;
7. require existing conversation summary;
8. require existing provider-storage.json;
9. create fresh CarbonStackSidecarProvider;
10. load MemoryStorage from provider-storage.json;
11. derive deterministic dev GroupId from conversation label;
12. load MlsGroup with MlsGroup::load(provider.storage(), &group_id);
13. read member KeyPackage artifact bytes;
14. deserialize member KeyPackage;
15. call group.add_members(&provider, &signer, &[member_key_package]);
16. extract Welcome from returned message/commit output;
17. write welcome.bin;
18. write welcome-manifest.json;
19. merge_pending_commit(&provider);
20. save mutated provider-storage.json back to disk;
21. write add-member-summary.json;
22. return sanitized stdout.

Important ordering:

    add_members
    extract Welcome
    write Welcome artifact / manifest
    merge_pending_commit
    save mutated provider storage
    write sanitized summary

The exact order may be adjusted if OpenMLS requires it, but the result doc must record the final order.

## 6. Expected generated files

Under:

    .carbonstack-openmls-sidecar-state/dev/conversations/<conversation-label>/

Expected new files:

    welcome.bin
    welcome-manifest.json
    add-member-summary.json

Existing files:

    conversation-summary.json
    provider-storage.json

Rules:

- welcome.bin is generated protocol artifact material and must not be printed to stdout;
- welcome-manifest.json is sanitized metadata;
- add-member-summary.json is sanitized metadata;
- provider-storage.json is generated dev OpenMLS provider/group state and must not be printed, inspected casually, committed, or treated as production secure storage;
- signer.json remains secret-bearing identity state and must not be printed, inspected casually, or committed.

## 7. Proposed stdout fields

Success envelope:

    ok=true
    command=conversation-add-member
    provider=openmls
    implementation=carbonstack-openmls-sidecar
    mode=experimental-sidecar
    phase=phase2d-conversation-add-member-dev
    private_material_included=false

Data fields:

    device_label
    conversation_label
    conversation_state_path_hint
    provider_storage_path_hint
    member_keypackage_path_hint
    welcome_artifact_path_hint
    welcome_manifest_path_hint
    add_member_summary_path_hint
    provider_storage_loaded=true
    provider_storage_written=true
    group_reloadable=true
    member_added=true
    pending_commit_merged=true
    member_count_before
    member_count_after
    epoch_before
    epoch_after
    group_id_ref
    welcome_artifact_written=true
    welcome_artifact_sha256
    welcome_artifact_size_bytes
    state_scope=dev-local-sidecar-state

Events:

    conversation.member_added
    conversation.welcome.created
    storage.saved

Trust:

- dev-sidecar metadata only for now;
- membership changes are product-level trust sensitive later;
- do not mutate trust state in this rung.

## 8. Proposed failure classes

Missing argument:

    missing_required_argument

Invalid labels:

    invalid_device_label
    invalid_conversation_label

Invalid artifact path:

    invalid_member_keypackage_path

Missing input state:

    identity_missing
    conversation_missing
    provider_storage_missing
    member_keypackage_missing

Unavailable/corrupt state:

    provider_storage_unavailable
    provider_group_unavailable
    secret_material_unavailable
    member_keypackage_invalid
    welcome_export_failed
    provider_storage_save_failed

Duplicate/refuse overwrite:

    add_member_artifact_exists

Generic failure:

    conversation_add_member_failed

## 9. Duplicate behavior

Initial implementation should refuse overwrite if any of these already exist:

    welcome.bin
    welcome-manifest.json
    add-member-summary.json

This is intentionally conservative. Regeneration or per-member/per-attempt directories can be designed later.

Potential later layout:

    members/<member-ref>/welcome.bin
    members/<member-ref>/welcome-manifest.json
    members/<member-ref>/add-member-summary.json

But v0 should stay simple unless implementation proves name collisions are immediately problematic.

## 10. KeyPackage artifact validation

The first implementation should accept:

    --member-keypackage <path>

Minimum validation:

- path argument exists;
- path is not empty;
- path points to a file;
- path size is non-zero;
- path size is below a conservative maximum;
- path basename is not signer.json;
- path basename is not provider-storage.json;
- path basename is not conversation-summary.json;
- path basename is not identity-state.json;
- path basename is not identity-summary.json;
- path basename is not public-bundle-summary.json;
- path basename is not public-bundle-manifest.json unless manifest mode is deliberately added later;
- content deserializes through the chosen OpenMLS/tls_codec path.

Do not trust extension alone.

## 11. API recon questions

Implementation must answer:

1. Should the sidecar deserialize directly to KeyPackage, KeyPackageIn, or another inbound type?
2. Which tls_codec trait/method is needed for KeyPackage artifact deserialization?
3. Does add_members require KeyPackage or KeyPackageIn conversion?
4. What exact return type does MlsGroup::add_members produce in openmls 0.8.1?
5. Where exactly is Welcome carried in the returned message/body?
6. Which type should be serialized for welcome.bin?
7. Which tls_codec method serializes Welcome in this dependency set?
8. Must merge_pending_commit happen before or after writing welcome.bin?
9. After merge_pending_commit and provider storage save, does conversation-load-check report updated epoch/member count?
10. Can a duplicate add-member be detected at this stage, or only artifact overwrite?

## 12. Required tests for implementation

Go-side contract tests should cover:

- provider-info lists conversation-add-member as supported after implementation;
- conversation-add-member missing device label fails;
- conversation-add-member missing conversation label fails;
- conversation-add-member missing member-keypackage fails;
- invalid labels fail;
- missing identity fails;
- missing conversation fails;
- missing provider storage fails;
- missing member KeyPackage path fails;
- invalid member KeyPackage artifact fails;
- success flow:
  - remove state;
  - identity-create Alice;
  - identity-create Bob;
  - public-bundle-export Bob --write-artifact;
  - conversation-create Alice;
  - conversation-load-check Alice;
  - conversation-add-member Alice with Bob KeyPackage artifact;
- success asserts:
  - ok=true;
  - member_added=true;
  - welcome_artifact_written=true;
  - provider_storage_written=true;
  - pending_commit_merged=true;
  - member_count_before=1;
  - member_count_after=2;
  - welcome.bin exists and non-empty;
  - welcome-manifest.json exists and is sanitized;
  - add-member-summary.json exists and is sanitized;
  - provider-storage.json exists but is not read;
  - stdout contains no private material.
- duplicate add-member artifact attempt refuses overwrite.

Rust-side tests may cover helper/path validation if feasible.

## 13. Manual probe target after implementation

Expected future probe:

    Remove-Item -Recurse -Force .\.carbonstack-openmls-sidecar-state -ErrorAction SilentlyContinue

    cargo run -- identity-create --device-label carbonstack-alice-device

    cargo run -- identity-create --device-label carbonstack-bob-device

    cargo run -- public-bundle-export --device-label carbonstack-bob-device --write-artifact

    cargo run -- conversation-create --device-label carbonstack-alice-device --conversation-label carbonstack-test-conversation

    cargo run -- conversation-load-check --device-label carbonstack-alice-device --conversation-label carbonstack-test-conversation

    cargo run -- conversation-add-member --device-label carbonstack-alice-device --conversation-label carbonstack-test-conversation --member-keypackage .\.carbonstack-openmls-sidecar-state\dev\devices\carbonstack-bob-device\public-bundle.keypackage.bin

Safe inspection only:

    Get-ChildItem .\.carbonstack-openmls-sidecar-state\dev\conversations\carbonstack-test-conversation | Select-Object Name, Length

Do not inspect:

    signer.json
    provider-storage.json
    welcome.bin
    raw KeyPackage bytes
    raw group state
    raw private material

## 14. Result doc requirements

The result doc must explicitly record:

- exact command surface implemented;
- exact artifact paths written;
- exact OpenMLS types used for KeyPackage deserialization;
- exact Welcome extraction path;
- exact Welcome serialization method;
- whether pending commit was merged;
- whether provider storage was saved after membership mutation;
- updated conversation-load-check behavior after add-member;
- tests run;
- artifact guard result;
- allowed claims;
- not allowed claims.

## 15. Next after implementation

After conversation-add-member / Welcome export succeeds:

    conversation-join skeleton/recon

Then:

    conversation-join implementation

Then:

    message-protect/message-open planning

Do not combine add-member and join into one checkpoint.
