# OpenMLS Sidecar Add-Member / Welcome API Recon v0

Status: API recon / pre-implementation result
Component: CarbonStackComms / OpenMLS sidecar
Phase: Phase 2D
Previous docs:
- docs/61-openmls-sidecar-conversation-add-member-welcome-plan-v0.md
- docs/62-openmls-sidecar-add-member-welcome-api-recon-v0.md
- docs/67-openmls-sidecar-add-member-welcome-skeleton-v0.md

## 1. Purpose

This document records the targeted OpenMLS API reconnaissance for the next implementation checkpoint:

    conversation-add-member --device-label <creator-device> --conversation-label <conversation> --member-keypackage <path>

The goal is to prepare v0.2.32 implementation without starting implementation in v0.2.31.

## 2. Current project state before implementation

v0.2.30 validated dev-local provider/group persistence:

- conversation-create writes provider-storage.json;
- conversation-create reports provider_storage_written=true;
- conversation-create reports group_reloadable=true;
- conversation-load-check reloads the group in a later sidecar invocation;
- provider-info lists conversation-load-check;
- conversation-add-member remains unsupported.

v0.2.31 docs/67 defines the add-member / Welcome export skeleton but does not implement it.

## 3. add_members API shape

OpenMLS 0.8.1 exposes:

    pub fn add_members<Provider: OpenMlsProvider>(
        &mut self,
        provider: &Provider,
        signer: &impl Signer,
        key_packages: &[KeyPackage],
    ) -> Result<
        (MlsMessageOut, MlsMessageOut, Option<GroupInfo>),
        AddMembersError<Provider::StorageError>,
    >

Confirmed behavior:

- add_members requires a mutable MlsGroup;
- it requires a provider;
- it requires a signer;
- it requires one or more KeyPackage values;
- it errors on empty key package input;
- it returns a triple:
  - first MlsMessageOut is the Commit;
  - second MlsMessageOut is the Welcome carrier;
  - third is optional GroupInfo.

Internal flow observed:

    commit_builder()
      .propose_adds(key_packages.iter().cloned())
      .force_self_update(force_self_update)
      .load_psks(provider.storage())?
      .build(provider.rand(), provider.crypto(), signer, |_| true)?
      .stage_commit(provider)?

Then:

    let welcome: MlsMessageOut = bundle.to_welcome_msg().ok_or(...)?;
    let (commit, _, group_info) = bundle.into_contents();
    self.reset_aad();
    Ok((commit, welcome, group_info))

## 4. Welcome carrier and serialization

OpenMLS has:

    MlsMessageOut

with:

    pub fn to_bytes(&self) -> Result<Vec<u8>, MlsMessageError>
    pub fn body(&self) -> &MlsMessageBodyOut

The to_bytes method calls:

    tls_serialize_detached()

The recon also confirmed:

    MlsMessageBodyOut::Welcome(Welcome)

Important implementation note:

The convenient owned extraction helper:

    into_welcome()

exists only behind:

    #[cfg(any(feature = "test-utils", test))]

Therefore the production sidecar should not rely on into_welcome unless test-utils is deliberately enabled and documented, which is not recommended for this rung.

Recommended v0.2.32 default:

    write welcome.bin from welcome_message.to_bytes()

where welcome_message is the second MlsMessageOut returned by add_members.

Rationale:

- MlsMessageOut is explicitly intended to be serialized and sent;
- to_bytes is public;
- this avoids needing private-field access or test-utils-only extraction;
- Bob-side join can later decide whether to deserialize MlsMessageOut and extract the Welcome, or whether a future checkpoint should export inner Welcome bytes separately.

Potential future alternative:

    write welcome.inner.bin from inner Welcome TLS bytes

But this should not be the first implementation unless the exact public extraction path is confirmed.

## 5. Message body facts

MlsMessageOut contains:

    version: ProtocolVersion
    body: MlsMessageBodyOut

MlsMessageBodyOut includes:

    PublicMessage(PublicMessage)
    PrivateMessage(PrivateMessage)
    Welcome(Welcome)
    GroupInfo(GroupInfo)
    KeyPackage(KeyPackage)

The body() accessor returns a borrowed body reference. This is useful for validating that the second returned message is actually a Welcome without consuming it.

Recommended implementation check:

    match welcome_message.body() {
        MlsMessageBodyOut::Welcome(_) => ok,
        _ => error welcome_export_failed,
    }

Then serialize the outer message:

    welcome_message.to_bytes()

## 6. KeyPackage input

add_members expects:

    &[KeyPackage]

The sidecar already writes Bob’s public KeyPackage artifact in v0.2.25:

    public-bundle.keypackage.bin

The exact deserialization line still needs implementation-time proof, but the likely path is tls_codec deserialization into KeyPackage or an inbound KeyPackage wrapper that converts into KeyPackage.

Implementation must not assume extension alone is valid. Content parsing is authoritative.

Minimum path validation remains:

- path exists;
- path is a file;
- size is non-zero;
- size is under a conservative max;
- basename is not signer.json;
- basename is not provider-storage.json;
- basename is not conversation-summary.json;
- basename is not identity-state.json;
- basename is not identity-summary.json;
- path is not a generated secret-bearing sidecar state file.

## 7. Recommended v0.2.32 operation sequence

Recommended implementation sequence:

1. validate device label;
2. validate conversation label;
3. parse and validate --member-keypackage path;
4. reject obvious secret/state paths;
5. require existing creator identity state;
6. load creator signer state;
7. require conversation-summary.json;
8. require provider-storage.json;
9. create CarbonStackSidecarProvider;
10. load provider storage from provider-storage.json;
11. derive deterministic dev GroupId;
12. load MlsGroup with MlsGroup::load(provider.storage(), &group_id);
13. read Bob member KeyPackage artifact bytes;
14. deserialize Bob KeyPackage;
15. record member_count_before and epoch_before;
16. call group.add_members(&provider, &signer, &[member_key_package]);
17. verify the second returned MlsMessageOut has MlsMessageBodyOut::Welcome(_);
18. serialize that second MlsMessageOut with to_bytes();
19. write welcome.bin;
20. write welcome-manifest.json;
21. call group.merge_pending_commit(&provider);
22. record member_count_after and epoch_after;
23. save mutated provider-storage.json;
24. optionally reload-check immediately;
25. write add-member-summary.json;
26. return sanitized stdout.

## 8. Open implementation questions

The next implementation still must prove:

1. exact KeyPackage artifact deserialization call;
2. exact import paths required for KeyPackage deserialization;
3. exact error type conversions for deserialization failure;
4. whether add_members accepts the deserialized KeyPackage directly or needs conversion from KeyPackageIn;
5. whether provider storage save after merge_pending_commit is enough for conversation-load-check to report member_count=2;
6. whether Welcome outer MlsMessageOut bytes are sufficient for the later join checkpoint;
7. whether optional GroupInfo should be written to a manifest field only or ignored for v0.

## 9. Planned output files

Under:

    .carbonstack-openmls-sidecar-state/dev/conversations/<conversation-label>/

Expected v0.2.32 generated files:

    welcome.bin
    welcome-manifest.json
    add-member-summary.json

Existing files:

    conversation-summary.json
    provider-storage.json

Rules:

- welcome.bin should not be printed to stdout;
- provider-storage.json should not be printed or inspected casually;
- signer.json should not be printed or inspected casually;
- add-member-summary.json and welcome-manifest.json should be sanitized metadata;
- all generated state must remain ignored by Git.

## 10. Proposed stdout success claims

conversation-add-member success may claim:

    member_added=true
    welcome_artifact_written=true
    provider_storage_written=true
    pending_commit_merged=true
    group_reloadable=true
    private_material_included=false

It may report:

    member_count_before=1
    member_count_after=2
    epoch_before=<value>
    epoch_after=<value>
    welcome_artifact_sha256=sha256:<hex>
    welcome_artifact_size_bytes=<n>

It must not claim:

- Bob has joined;
- conversation-join exists;
- message protect/open exists;
- production secure storage exists;
- production E2EE exists;
- trust-state has been mutated.

## 11. Proposed tests for v0.2.32

Go contract tests should cover:

- provider-info lists conversation-add-member as supported after implementation;
- conversation-add-member missing --device-label fails;
- missing --conversation-label fails;
- missing --member-keypackage fails;
- invalid labels fail;
- missing identity fails;
- missing conversation fails;
- missing provider storage fails;
- missing member KeyPackage file fails;
- invalid member KeyPackage file fails;
- successful flow:
  - remove state;
  - identity-create Alice;
  - identity-create Bob;
  - public-bundle-export Bob --write-artifact;
  - conversation-create Alice;
  - conversation-load-check Alice;
  - conversation-add-member Alice using Bob public-bundle.keypackage.bin;
- success asserts:
  - ok=true;
  - member_added=true;
  - welcome_artifact_written=true;
  - provider_storage_written=true;
  - pending_commit_merged=true;
  - member_count_before=1;
  - member_count_after=2;
  - welcome.bin exists and is non-empty;
  - welcome-manifest.json exists and is sanitized;
  - add-member-summary.json exists and is sanitized;
  - provider-storage.json exists but is not read;
  - stdout contains no private material.
- duplicate artifact attempt refuses overwrite.

## 12. Recommended implementation boundary

v0.2.32 should implement only:

    conversation-add-member + Welcome export

It should not implement:

- Bob-side conversation-join;
- message-protect;
- message-open;
- Comms runtime integration;
- Cypher routing;
- trust-state mutation;
- production vault;
- Android;
- CarbonStackOS.

## 13. Next checkpoint after add-member

After v0.2.32 succeeds:

    conversation-join skeleton/recon

Then:

    conversation-join implementation

Only after join works should message-protect/message-open be planned.
