# OpenMLS Sidecar Add-Member / Welcome Export Result v0

Status: Validated implementation checkpoint
Component: CarbonStackComms / OpenMLS sidecar
Phase: Phase 2D
Previous docs:
- docs/67-openmls-sidecar-add-member-welcome-skeleton-v0.md
- docs/68-openmls-sidecar-add-member-welcome-api-recon-v0.md
- docs/66-openmls-sidecar-dev-provider-group-persistence-result-v0.md

## 1. Summary

This checkpoint implements the first dev-local OpenMLS add-member / Welcome export path.

The new sidecar command is:

    conversation-add-member --device-label <creator-device> --conversation-label <conversation> --member-keypackage <path>

The validated flow is:

1. Alice identity exists.
2. Bob identity exists.
3. Bob exports a public KeyPackage artifact.
4. Alice creates a conversation.
5. Alice reload-checks the persisted group.
6. Alice adds Bob's KeyPackage to the group.
7. The sidecar exports a Welcome carrier artifact.
8. The sidecar merges the pending commit.
9. The sidecar saves mutated provider storage.
10. A later load-check confirms the group reloads with two members.

This checkpoint does not implement Bob-side join.

## 2. Implemented command

New command:

    conversation-add-member

Required arguments:

    --device-label <creator-device>
    --conversation-label <conversation>
    --member-keypackage <path>

Validated example:

    cargo run -- conversation-add-member --device-label carbonstack-alice-device --conversation-label carbonstack-test-conversation --member-keypackage .\.carbonstack-openmls-sidecar-state\dev\devices\carbonstack-bob-device\public-bundle.keypackage.bin

## 3. Implementation files

CarbonStackComms changed:

    internal/protocol/mls/research/openmls-sidecar/src/state.rs
    internal/protocol/mls/research/openmls-sidecar/src/main.rs
    internal/protocol/openmls_sidecar_provider_info_test.go

Key implementation additions:

    add_dev_conversation_member(...)
    ConversationAddMemberResult
    WelcomeManifest
    AddMemberSummary
    validate_member_keypackage_path(...)
    conversation_welcome_artifact_path(...)
    conversation_welcome_manifest_path(...)
    conversation_add_member_summary_path(...)

## 4. OpenMLS API path used

The implementation uses the API path proven during v0.2.32 recon:

    KeyPackageIn::tls_deserialize(...)
    KeyPackageIn::validate(provider.crypto(), ProtocolVersion::default())
    MlsGroup::load(...)
    MlsGroup::add_members(...)
    MlsMessageOut::body()
    MlsMessageOut::to_bytes()
    MlsGroup::merge_pending_commit(...)
    CarbonStackSidecarProvider::save_storage_to_path(...)

Important recon result:

    KeyPackage::tls_deserialize(...) does not exist.
    KeyPackageIn::tls_deserialize(...) exists.
    KeyPackageIn cannot be converted into KeyPackage with .into().
    KeyPackageIn must be validated with validate(...) to produce KeyPackage.

The Welcome output is the second MlsMessageOut returned by add_members. The implementation verifies that this message body is:

    MlsMessageBodyOut::Welcome(_)

Then it serializes the outer MlsMessageOut with:

    welcome_message.to_bytes()

This writes a sendable Welcome carrier artifact rather than attempting to use the test-gated into_welcome() helper.

## 5. Operation sequence

The implemented sequence is:

1. parse device label;
2. parse conversation label;
3. parse member-keypackage path;
4. validate labels;
5. validate member-keypackage path;
6. load creator identity status;
7. load creator signer;
8. require conversation-summary.json;
9. require provider-storage.json;
10. create CarbonStackSidecarProvider;
11. load MemoryStorage from provider-storage.json;
12. derive deterministic dev GroupId from the conversation label;
13. call MlsGroup::load(provider.storage(), &group_id);
14. read Bob public-bundle.keypackage.bin;
15. deserialize it as KeyPackageIn;
16. validate it into KeyPackage;
17. record member_count_before and epoch_before;
18. call group.add_members(&provider, &signer, &[member_key_package]);
19. verify the second returned MlsMessageOut is a Welcome;
20. serialize the Welcome carrier with to_bytes();
21. write welcome.bin;
22. call group.merge_pending_commit(&provider);
23. record member_count_after and epoch_after;
24. save mutated provider-storage.json;
25. write welcome-manifest.json;
26. write add-member-summary.json;
27. return sanitized stdout.

## 6. Generated files

Under:

    .carbonstack-openmls-sidecar-state/dev/conversations/<conversation-label>/

The command writes:

    welcome.bin
    welcome-manifest.json
    add-member-summary.json

Existing files retained/updated:

    conversation-summary.json
    provider-storage.json

Important handling rules:

- welcome.bin is generated protocol artifact material and is not printed to stdout.
- provider-storage.json is generated dev OpenMLS provider/group state and is not printed to stdout.
- signer.json remains secret-bearing identity state and is not printed to stdout.
- add-member-summary.json and welcome-manifest.json are sanitized metadata.
- all generated sidecar state remains ignored by Git.

## 7. Success envelope

Successful conversation-add-member output reports:

    ok=true
    command=conversation-add-member
    phase=phase2d-conversation-add-member-dev
    private_material_included=false
    member_added=true
    welcome_artifact_written=true
    provider_storage_loaded=true
    provider_storage_written=true
    group_reloadable=true
    pending_commit_merged=true
    member_count_before=1
    member_count_after=2
    epoch_before=GroupEpoch(0)
    epoch_after=GroupEpoch(1)
    welcome_artifact_size_bytes>0

Manual validation observed:

    welcome_artifact_size_bytes=880
    member_count_before=1
    member_count_after=2
    epoch_before=GroupEpoch(0)
    epoch_after=GroupEpoch(1)

## 8. Provider-info changes

provider-info now lists conversation-add-member as supported.

Capabilities include:

    provider-info
    identity-create
    identity-status
    public-bundle-export
    conversation-create
    conversation-load-check
    conversation-add-member

Unsupported still includes:

    conversation-join
    message-protect
    message-open
    state-checkpoint
    state-load-check

The stale load-check warning that said conversation-add-member was not implemented was patched to reflect that add-member is now implemented for dev-local Welcome export.

## 9. Events

conversation-add-member success emits:

    provider.conversation.member_added
    provider.welcome.exported

Both are currently emitted as info events with trust_relevant=false in this dev-sidecar rung.

This checkpoint does not mutate trust.json or trust-events.jsonl.

## 10. Manual validation

Validated manual probe:

    Remove-Item -Recurse -Force .\.carbonstack-openmls-sidecar-state -ErrorAction SilentlyContinue

    cargo run -- identity-create --device-label carbonstack-alice-device

    cargo run -- identity-create --device-label carbonstack-bob-device

    cargo run -- public-bundle-export --device-label carbonstack-bob-device --write-artifact

    cargo run -- conversation-create --device-label carbonstack-alice-device --conversation-label carbonstack-test-conversation

    cargo run -- conversation-load-check --device-label carbonstack-alice-device --conversation-label carbonstack-test-conversation

    cargo run -- conversation-add-member --device-label carbonstack-alice-device --conversation-label carbonstack-test-conversation --member-keypackage .\.carbonstack-openmls-sidecar-state\dev\devices\carbonstack-bob-device\public-bundle.keypackage.bin

    cargo run -- conversation-load-check --device-label carbonstack-alice-device --conversation-label carbonstack-test-conversation

Manual proof:

- add-member returned ok=true;
- add-member wrote welcome.bin;
- add-member wrote welcome-manifest.json;
- add-member wrote add-member-summary.json;
- add-member reported provider_storage_loaded=true;
- add-member reported provider_storage_written=true;
- add-member reported pending_commit_merged=true;
- add-member reported member_count_before=1;
- add-member reported member_count_after=2;
- add-member reported epoch_before=GroupEpoch(0);
- add-member reported epoch_after=GroupEpoch(1);
- post-add conversation-load-check returned member_count=2;
- post-add conversation-load-check returned epoch=GroupEpoch(1);
- post-add conversation-load-check returned group_reloadable=true;
- no private material was printed.

Safe file listing showed:

    add-member-summary.json
    conversation-summary.json
    provider-storage.json
    welcome-manifest.json
    welcome.bin

## 11. Go contract tests

Go-side contract coverage was added for:

    TestOpenMLSSidecarConversationAddMemberWelcomeExport

The test validates:

- Alice identity-create;
- Bob identity-create;
- Bob public-bundle-export --write-artifact;
- Alice conversation-create;
- pre-add conversation-load-check reports member_count=1;
- conversation-add-member succeeds;
- member_added=true;
- welcome_artifact_written=true;
- provider_storage_loaded=true;
- provider_storage_written=true;
- group_reloadable=true;
- pending_commit_merged=true;
- member_count_before=1;
- member_count_after=2;
- Welcome artifact path is returned;
- Welcome manifest path is returned;
- add-member summary path is returned;
- Welcome artifact hash is returned;
- Welcome artifact size is positive;
- success emits two events;
- no secret material appears in stdout;
- welcome.bin exists but is not read;
- welcome-manifest.json exists;
- add-member-summary.json exists;
- post-add load-check reports member_count=2;
- duplicate add-member refuses overwrite with add_member_artifact_exists.

Provider-info test was updated so conversation-add-member is supported, while conversation-join remains unsupported.

Unsupported command test was moved from conversation-add-member to conversation-join.

## 12. Validation commands

Validated during the rung:

    cargo check

    cargo run -- provider-info

    cargo run -- conversation-add-member --device-label carbonstack-alice-device --conversation-label carbonstack-test-conversation --member-keypackage .\.carbonstack-openmls-sidecar-state\dev\devices\carbonstack-bob-device\public-bundle.keypackage.bin

    cargo run -- conversation-load-check --device-label carbonstack-alice-device --conversation-label carbonstack-test-conversation

    go test ./internal/protocol -run "TestOpenMLSSidecarProviderInfoCommand|TestOpenMLSSidecarUnsupportedCommandEnvelope|TestOpenMLSSidecarConversationAddMemberWelcomeExport"

Final closeout validation should still include:

    go test ./internal/protocol
    go test ./...
    powershell -ExecutionPolicy Bypass -File .\scripts\check-no-rust-artifacts.ps1
    powershell -ExecutionPolicy Bypass -File .\scripts\validate-local.ps1

## 13. Blunders and lessons from this rung

### 13.1 KeyPackage deserialize path was not obvious

Initial assumption:

    KeyPackage::tls_deserialize(...)

was wrong.

Then:

    KeyPackageIn::tls_deserialize(...)
    key_package_in.into()

was also wrong.

Correct path:

    KeyPackageIn::tls_deserialize(...)
    KeyPackageIn::validate(provider.crypto(), ProtocolVersion::default())

Lesson:

OpenMLS separates inbound unverified KeyPackage material from verified KeyPackage values. Do not bypass validation or assume .into() exists.

### 13.2 Provider-info raw JSON remains fragile

provider-info still has manually formatted/raw JSON-like sections. Earlier rungs already showed this causes brittle patches. In this rung, conversation-add-member had to be moved carefully from unsupported to capabilities.

Lesson:

Future cleanup should convert provider-info output to structured serde_json::json! generation or central command capability arrays.

### 13.3 PowerShell patching can inject garbage when replacement strings are wrong

A bad replacement inserted:

    .\.git\assertStringPresent(...)

into the Go file, causing gofmt errors.

Lesson:

Use @' ... '@ blocks for patch text.
Inspect small windows after risky regex replacements.
Prefer line-based or exact visible block patches.
Run gofmt immediately after each patch.
If gofmt fails, inspect the exact line before continuing.

### 13.4 Exact string patches failed repeatedly due to gofmt tabs/spaces

Several $Text.Contains($Old) patches failed because the Go file had gofmt-managed tabs and indentation.

Lesson:

For Go files, regex patches anchored on semantic lines are safer than exact pasted blocks.
When possible, inspect the target range first and patch the smallest visible block.

### 13.5 Unsupported-command test had to move targets

conversation-add-member used to be an unsupported command. After implementation, the unsupported command test needed to move to conversation-join.

Lesson:

Whenever a command graduates from unsupported to supported, update both provider-info assertions and unsupported-command tests.

## 14. Allowed claims after this checkpoint

Allowed:

- CarbonStack OpenMLS sidecar can add a member to a persisted dev-local MLS group.
- conversation-add-member accepts a direct Bob KeyPackage artifact path.
- Bob's KeyPackage artifact is deserialized as KeyPackageIn and validated into KeyPackage.
- conversation-add-member calls MlsGroup::add_members.
- conversation-add-member exports a Welcome carrier artifact as welcome.bin.
- conversation-add-member merges the pending commit.
- conversation-add-member saves mutated provider-storage.json.
- post-add conversation-load-check reloads the group and reports member_count=2.
- provider-info lists conversation-add-member as supported.
- Go contract tests cover the add-member / Welcome export path.

Not allowed:

- Bob has joined the group.
- conversation-join exists.
- message-protect exists.
- message-open exists.
- Cypher routes MLS payloads.
- Comms runtime consumes MLS payloads.
- trust-state is updated from add-member events.
- provider-storage.json is production secure storage.
- welcome.bin is a production onboarding UX.
- production E2EE exists.
- Signal-equivalent security exists.

## 15. Next recommended checkpoint

Next checkpoint:

    OpenMLS sidecar conversation-join skeleton / API recon

Recommended future doc:

    docs/70-openmls-sidecar-conversation-join-skeleton-v0.md

The next checkpoint should define:

- Bob-side command surface;
- Welcome artifact input path;
- Bob identity load path;
- Welcome deserialize path;
- StagedWelcome::new_from_welcome path;
- into_group / new group creation path;
- provider storage save path for Bob;
- Bob-side conversation summary path;
- sanitized stdout;
- tests;
- non-goals.

Implementation should not begin until the join skeleton and API recon are written.
