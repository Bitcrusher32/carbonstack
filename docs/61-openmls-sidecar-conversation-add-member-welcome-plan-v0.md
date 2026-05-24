# OpenMLS Sidecar Conversation Add-Member / Welcome Export Plan v0

Status: Planned
Component: CarbonStackComms / OpenMLS sidecar
Phase: Phase 2D
Previous docs:
- docs/57-openmls-sidecar-conversation-lifecycle-plan-v0.md
- docs/58-openmls-sidecar-conversation-lifecycle-api-recon-v0.md
- docs/59-openmls-sidecar-conversation-create-plan-v0.md
- docs/60-openmls-sidecar-conversation-create-result-v0.md

## 1. Purpose

This document defines the next intended implementation rung after validated dev-sidecar conversation-create.

The next target is:

    conversation-add-member --device-label <creator-device> --conversation-label <safe-conversation> --member-keypackage <path>

The command should add a member to an existing dev-local OpenMLS conversation by consuming a serialized public KeyPackage artifact and writing a serialized Welcome artifact plus sanitized metadata.

This is a planning checkpoint only. It does not implement add-member, Welcome export, join, message protect/open, Comms runtime integration, Cypher routing, trust-state mutation, Android, or CarbonStackOS.

## 2. Current baseline

Validated before this plan:

    provider-info
    identity-create
    identity-status
    public-bundle-export
    public-bundle-export --write-artifact
    conversation-create

Current supported sidecar command state:

    provider-info
    identity-create
    identity-status
    public-bundle-export
    conversation-create

Current unsupported command state:

    conversation-add-member
    conversation-join
    message-protect
    message-open
    state-checkpoint
    state-load-check

v0.2.27 conversation-create validates:

- device label;
- conversation label;
- existing identity state;
- dev-local OpenMLS group creation;
- sanitized conversation-summary.json;
- duplicate refusal;
- provider.conversation.created event;
- provider.conversation.exists event;
- provider_storage_written=true;
- private_material_included=false.

## 3. Planned command shape

Recommended command:

    conversation-add-member --device-label <creator-device-label> --conversation-label <conversation-label> --member-keypackage <path>

Required arguments:

    --device-label
    --conversation-label
    --member-keypackage

Argument meanings:

- `--device-label` is the existing local device identity that owns/created the conversation.
- `--conversation-label` is the existing local dev conversation/group label.
- `--member-keypackage` is a path to another device's serialized public KeyPackage artifact, currently produced by:

      public-bundle-export --device-label <member-device> --write-artifact

Initial expected artifact name:

    public-bundle.keypackage.bin

## 4. Planned inputs

Required local state:

    .carbonstack-openmls-sidecar-state/dev/devices/<creator-device-label>/
      signer.json
      identity-summary.json
      identity-state.json

Required conversation state:

    .carbonstack-openmls-sidecar-state/dev/conversations/<conversation-label>/
      conversation-summary.json
      future provider/group state needed to reload MlsGroup

Required member artifact:

    public-bundle.keypackage.bin

Important blocker:

v0.2.27 conversation-create writes sanitized `conversation-summary.json`, but the next implementation must verify whether it persists enough OpenMLS group/provider state for a later command invocation to call `add_members`.

If not, the next code checkpoint must first improve conversation-create persistence before implementing add-member.

## 5. Planned outputs

Recommended output files under:

    .carbonstack-openmls-sidecar-state/dev/conversations/<conversation-label>/

Add-member / Welcome export should write:

    add-member-summary.json
    welcome.bin
    welcome-manifest.json

Potential future files, only if required:

    group-state.bin
    provider-state/
    commit-summary.json

Do not write raw private material, signer contents, or secret provider storage into stdout.

## 6. OpenMLS API basis

Recon confirms OpenMLS 0.8.1 has:

    MlsGroup::add_members(
        &mut self,
        provider,
        signer,
        &[KeyPackage],
    )

Prior scratch code confirms the rough pattern:

    let (_commit, welcome_msg, group_info) =
        alice_group.add_members(&alice_provider, &alice.signer, &[bob.key_package.clone()])?;

    let welcome = match welcome_msg.body() {
        MlsMessageBodyOut::Welcome(welcome) => welcome.clone(),
        other => panic!("expected Welcome MlsMessageOut body, got: {:?}", other),
    };

    alice_group.merge_pending_commit(&alice_provider)?;

The implementation must not rely on memory alone. Before coding, inspect exact local line windows for:

- `MlsGroup::add_members`;
- `MlsMessageBodyOut::Welcome`;
- `Welcome` TLS serialization;
- `KeyPackageIn` deserialization and validation;
- conversion from incoming KeyPackage material to `KeyPackage`;
- group/provider state load and save.

## 7. KeyPackage consumption plan

The input artifact is expected to be a serialized public OpenMLS KeyPackage:

    public-bundle.keypackage.bin

Possible deserialization path:

1. read bytes from `--member-keypackage`;
2. TLS-deserialize into `KeyPackageIn` or another OpenMLS input type;
3. validate/convert into `KeyPackage`;
4. compute a sanitized sha256 artifact hash;
5. compute or verify OpenMLS KeyPackage ref;
6. pass `&[KeyPackage]` into `add_members`.

Open issue:

The exact `KeyPackageIn` validation / conversion path must be confirmed before implementation.

Do not accept arbitrary JSON wrappers or stdin bytes in this rung. Use an explicit file path only.

## 8. Welcome export plan

After successful `add_members`, extract Welcome from the returned message wrapper.

Expected scratch-derived pattern:

    match welcome_msg.body() {
        MlsMessageBodyOut::Welcome(welcome) => welcome.clone(),
        other => fail,
    }

Then serialize Welcome using the local `tls_codec` serialization method confirmed for the pinned dependency set.

Recommended output:

    welcome.bin

Recommended manifest:

    welcome-manifest.json

Manifest fields:

    manifest_version
    conversation_label
    inviter_device_label
    member_keypackage_path
    member_keypackage_sha256
    member_keypackage_ref
    welcome_artifact
    welcome_artifact_sha256
    welcome_artifact_size_bytes
    provider_storage_written
    private_material_included
    warning

## 9. Add-member summary plan

Recommended summary file:

    add-member-summary.json

Recommended fields:

    summary_version
    conversation_label
    inviter_device_label
    member_keypackage_ref
    member_keypackage_sha256
    welcome_exported
    welcome_artifact
    welcome_artifact_sha256
    welcome_artifact_size_bytes
    member_count_after_add
    epoch_after_add
    pending_commit_merged
    provider_storage_written
    private_material_included
    warning

## 10. Duplicate and overwrite behavior

Default behavior must refuse silent overwrite.

conversation-add-member should fail if any of these already exist:

    welcome.bin
    welcome-manifest.json
    add-member-summary.json

No `--force` in this rung.

Reason:

Welcome replacement and membership changes are trust-sensitive. Silent replacement would conflict with CarbonStack's loud trust-change doctrine.

## 11. Provider storage and group persistence

This is the main implementation risk.

conversation-add-member requires a usable mutable `MlsGroup`.

The implementation must answer:

- Can the group created in v0.2.27 be reloaded from provider storage?
- Is the current `OpenMlsRustCrypto::default()` provider enough across process invocations?
- Does conversation-create currently persist enough group state?
- If not, should v0.2.29 patch conversation-create persistence first?
- Where should dev-local provider/group state live?
- What can be summarized safely without exposing secrets?

If provider/group storage is written or mutated, stdout and summaries must report:

    provider_storage_written=true

But this remains dev-local provider state, not production vault storage.

## 12. Planned stdout behavior

Success stdout may include:

    command
    ok
    provider
    implementation
    phase
    device_label
    conversation_label
    member_keypackage_path_hint
    member_keypackage_sha256
    member_keypackage_ref
    welcome_artifact_path_hint
    welcome_artifact_sha256
    welcome_artifact_size_bytes
    member_count_after_add
    epoch_after_add
    provider_storage_written
    private_material_included=false

Success stdout must not include:

    signer.json
    private keys
    raw provider storage
    raw group state
    raw Welcome bytes
    raw KeyPackage bytes
    message plaintext
    message ciphertext
    recovery material

## 13. Planned events

Likely new provider events:

    provider.conversation.member_added
    provider.welcome.exported
    provider.artifact.invalid
    provider.conversation.missing

Initial classifications:

### provider.conversation.member_added

- severity: info or notice
- trust_relevant: false for dev-sidecar rung, but product-level trust-sensitive later
- action: append history / debug only for now

### provider.welcome.exported

- severity: info
- trust_relevant: false for this rung
- action: append history / debug only

### provider.artifact.invalid

- severity: warning
- trust_relevant: false initially
- action: stop operation / append history

### provider.conversation.missing

- severity: warning
- trust_relevant: false initially
- action: stop operation / append history

Important:

Future user-facing membership events must become loud and trust-visible. This rung must not silently normalize membership changes as boring forever.

## 14. Required tests

Go-side tests should cover:

- provider-info still lists conversation-add-member unsupported until implementation starts;
- after implementation, provider-info lists conversation-add-member supported;
- missing device label fails;
- missing conversation label fails;
- missing member-keypackage path fails;
- invalid device label fails;
- invalid conversation label fails;
- missing creator identity fails;
- missing conversation fails;
- missing KeyPackage artifact fails;
- invalid KeyPackage artifact fails;
- successful Alice/Bob setup:
  - identity-create Alice;
  - identity-create Bob;
  - public-bundle-export Bob --write-artifact;
  - conversation-create Alice;
  - conversation-add-member Alice using Bob artifact;
- welcome.bin exists;
- welcome-manifest.json exists;
- add-member-summary.json exists;
- manifest hash/size match actual welcome.bin;
- stdout reports private_material_included=false;
- stdout does not include obvious secret material;
- duplicate add-member / Welcome export refuses overwrite;
- conversation-join remains unsupported;
- message-protect/open remain unsupported.

Rust-side tests may cover:

- member-keypackage path parsing if helper is added;
- safe artifact path constraints if path validation is implemented;
- no extension-trust behavior if artifact type is validated by parse, not filename.

## 15. Manual probe shape after future implementation

Potential manual probe:

    Remove-Item -Recurse -Force .\.carbonstack-openmls-sidecar-state -ErrorAction SilentlyContinue

    cargo run -- identity-create --device-label carbonstack-alice-device
    cargo run -- identity-create --device-label carbonstack-bob-device

    cargo run -- public-bundle-export --device-label carbonstack-bob-device --write-artifact

    cargo run -- conversation-create --device-label carbonstack-alice-device --conversation-label carbonstack-test-conversation

    cargo run -- conversation-add-member --device-label carbonstack-alice-device --conversation-label carbonstack-test-conversation --member-keypackage .carbonstack-openmls-sidecar-state\dev\devices\carbonstack-bob-device\public-bundle.keypackage.bin

Safe inspection:

    Get-ChildItem .\.carbonstack-openmls-sidecar-state\dev\conversations\carbonstack-test-conversation | Select-Object Name, Length
    Get-Content .\.carbonstack-openmls-sidecar-state\dev\conversations\carbonstack-test-conversation\welcome-manifest.json
    Get-Content .\.carbonstack-openmls-sidecar-state\dev\conversations\carbonstack-test-conversation\add-member-summary.json

Do not inspect:

    signer.json
    raw provider storage
    raw group state
    raw Welcome bytes

## 16. Non-goals

This rung must not implement:

- Bob join;
- StagedWelcome use in sidecar runtime;
- message protect;
- message open;
- Comms send/inbox integration;
- Cypher MLS payload routing;
- trust.json mutation;
- trust-events.jsonl mutation;
- production storage;
- secure vault;
- hostile-server proof;
- replay resistance;
- metadata privacy;
- Android;
- CarbonStackOS.

## 17. Success criteria for future implementation

The implementation checkpoint succeeds only when:

- planning doc exists before code;
- exact APIs are inspected before patching;
- add-member succeeds for Alice/Bob dev flow;
- Welcome artifact is written;
- Welcome manifest is written;
- duplicate overwrite is refused;
- no secret material appears in stdout;
- Go-side tests validate behavior;
- artifact guard passes;
- result doc records behavior and limitations;
- conversation-join remains blocked until its own plan.

## 18. Recommended next checkpoint

After this planning/recon checkpoint, the next safest implementation checkpoint is:

    conversation-add-member / Welcome export implementation

However, if implementation inspection proves v0.2.27 conversation-create does not persist enough group/provider state, the next checkpoint should instead be:

    conversation-create persistence repair

Do not force add-member through bad state assumptions.
