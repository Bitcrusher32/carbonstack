# OpenMLS Sidecar Conversation Lifecycle API Recon v0

Status: Reconnaissance
Component: CarbonStackComms / OpenMLS sidecar
Phase: Phase 2D
Related plan:
- docs/57-openmls-sidecar-conversation-lifecycle-plan-v0.md

## 1. Purpose

This document records targeted API reconnaissance for the future OpenMLS sidecar conversation lifecycle rung.

The goal is to reduce implementation uncertainty before touching sidecar conversation code.

This is not an implementation result. No conversation-create, conversation-add-member, conversation-join, message-protect, message-open, Comms runtime wiring, Cypher routing, trust-state mutation, Android work, or CarbonStackOS work is implemented by this checkpoint.

## 2. Local dependency set

The sidecar currently uses the following pinned local dependency set:

- openmls 0.8.1
- openmls_rust_crypto 0.5.1
- openmls_traits 0.5.0
- tls_codec 0.4.2

This matters because future implementation must target these exact APIs, not generic OpenMLS memory or newer docs.

## 3. Source locations identified

Targeted local PowerShell inspection identified the relevant OpenMLS source locations.

### MlsGroup

Primary source:

    openmls-0.8.1/src/group/mls_group/mod.rs

Relevant surrounding files:

    openmls-0.8.1/src/group/mls_group/builder.rs
    openmls-0.8.1/src/group/mls_group/config.rs
    openmls-0.8.1/src/group/mls_group/creation.rs
    openmls-0.8.1/src/group/mls_group/membership.rs

### Welcome

Primary source candidates:

    openmls-0.8.1/src/messages/mod.rs
    openmls-0.8.1/src/group/mls_group/commit_builder.rs
    openmls-0.8.1/src/group/mls_group/tests_and_kats/kats/welcome.rs

### KeyPackage / KeyPackageIn

Primary source:

    openmls-0.8.1/src/key_packages/mod.rs
    openmls-0.8.1/src/key_packages/key_package_in.rs

### Join / add APIs

Targeted source locations:

    openmls-0.8.1/src/group/mls_group/creation.rs
      new_from_welcome

    openmls-0.8.1/src/group/mls_group/membership.rs
      add_members

## 4. Existing CarbonStack research code

The older `openmls-minimal` research crate remains the highest-value implementation reference.

Path:

    carbonstack-comms/internal/protocol/mls/research/openmls-minimal/src/main.rs

It already demonstrates a complete scratch-level lifecycle:

1. create Alice provider;
2. create Bob provider;
3. create Alice setup material;
4. create Bob setup material;
5. create Alice group;
6. add Bob from Bob KeyPackage;
7. extract Welcome from MlsMessageOut;
8. merge Alice pending commit;
9. stage Welcome for Bob;
10. turn staged Welcome into Bob group;
11. protect/open message one;
12. save/load MemoryStorage;
13. reload groups with MlsGroup::load;
14. protect/open message two.

This is still scratch research. It is not sidecar runtime code and does not establish production E2EE.

## 5. Group creation pattern

The scratch code uses:

    let create_config = MlsGroupCreateConfig::builder()
        .ciphersuite(ciphersuite)
        .use_ratchet_tree_extension(true)
        .build();

    let group_id = GroupId::from_slice(GROUP_ID_BYTES);

    let mut alice_group = MlsGroup::new_with_group_id(
        &alice_provider,
        &alice.signer,
        &create_config,
        group_id,
        alice.credential_with_key.clone(),
    )
    .expect("failed to create Alice MLS group");

Implication for sidecar:

- `conversation-create` should probably use `MlsGroup::new_with_group_id` at first.
- The future sidecar must define deterministic dev-local group ID bytes from a safe conversation label.
- The result must be described as dev-local conversation/group state, not production conversation state.

## 6. Add-member pattern

The scratch code uses:

    let (_commit, welcome_msg, group_info) = alice_group
        .add_members(&alice_provider, &alice.signer, &[bob.key_package.clone()])
        .expect("failed to add Bob to Alice MLS group");

Then it extracts Welcome:

    let welcome = match welcome_msg.body() {
        MlsMessageBodyOut::Welcome(welcome) => welcome.clone(),
        other => panic!("expected Welcome MlsMessageOut body, got: {:?}", other),
    };

Then Alice merges the pending commit:

    alice_group
        .merge_pending_commit(&alice_provider)
        .expect("failed to merge Alice pending commit after adding Bob");

Implication for sidecar:

- `conversation-add-member` probably consumes a serialized public KeyPackage artifact.
- That artifact must be deserialized back into the OpenMLS KeyPackage type expected by `add_members`.
- The command should write a serialized Welcome artifact and manifest.
- The command should merge Alice pending commit only after successful add-member flow.
- Duplicate Welcome output should refuse overwrite by default.

## 7. Join-from-Welcome pattern

The scratch code uses:

    let join_config = MlsGroupJoinConfig::builder()
        .use_ratchet_tree_extension(true)
        .build();

    let staged_welcome = StagedWelcome::new_from_welcome(
        &bob_provider,
        &join_config,
        welcome,
        None,
    )
    .expect("failed to stage Bob Welcome");

    let mut bob_group = staged_welcome
        .into_group(&bob_provider)
        .expect("failed to turn staged Welcome into Bob MLS group");

Implication for sidecar:

- `conversation-join` should consume a serialized Welcome artifact.
- It should produce sanitized joined-conversation summary output.
- It likely writes or mutates dev-local provider/group state for the joining device.
- It must not print Welcome bytes, provider storage, signer material, or group state.

## 8. Provider/group storage implications

The scratch code confirms that OpenMLS group operations are stateful.

Important observed patterns:

- `add_members` mutates Alice-side group state.
- Alice must call `merge_pending_commit`.
- `process_message` mutates Bob-side group state in the later message flow.
- `MlsGroup::load(provider.storage(), &group_id)` can reload groups from provider storage in the scratch persistence probe.
- MemoryStorage was useful for scratch persistence, but it is not a production vault.

Implication for sidecar:

- Conversation lifecycle probably cannot stay `provider_storage_written=false` forever.
- The next implementation must decide whether to introduce explicit dev-local provider storage.
- If provider storage is written, stdout and manifests must say `provider_storage_written=true`.
- That storage must be documented as dev-local provider state, not secure vault storage.

## 9. Current sidecar command state

Current sidecar supported commands:

    provider-info
    identity-create
    identity-status
    public-bundle-export

Current sidecar unsupported commands still include:

    conversation-create
    conversation-add-member
    conversation-join
    message-protect
    message-open
    state-checkpoint
    state-load-check

This is correct for the current checkpoint.

Conversation lifecycle remains planned only.

## 10. Implementation recommendation after recon

Do not implement the full create/add/join lifecycle in one jump unless API coupling forces it.

Recommended next implementation checkpoint:

    conversation-create only

Reason:

- It introduces conversation label validation.
- It introduces dev conversation state path helpers.
- It forces the provider/group storage decision.
- It avoids immediately combining KeyPackage consumption, Welcome export, join, and storage persistence into one risky patch.

Suggested next plan/result docs:

    docs/59-openmls-sidecar-conversation-create-plan-v0.md
    docs/60-openmls-sidecar-conversation-create-result-v0.md

Only after conversation-create validates should CarbonStack move to:

    conversation-add-member / Welcome export planning

## 11. Questions still open before implementation

- What exact line-range-confirmed signature should be used for `MlsGroup::new_with_group_id` in the pinned OpenMLS dependency?
- What exact line-range-confirmed signature should be used for `add_members`?
- What exact line-range-confirmed serialization path should be used for Welcome?
- What exact line-range-confirmed deserialization path should be used for KeyPackage artifacts?
- Should conversation labels produce deterministic GroupId bytes directly?
- Should conversation state be global under `conversations/<label>/` or nested by creator device?
- Should provider storage be stored per device, per conversation, or both?
- Should the next implementation expose only `conversation-create` or also a `conversation-status` check?

## 12. Recommended next PowerShell inspection

Use tiny, one-file line windows only. Avoid recursive context floods.

Recommended method:

    $Path = "<exact file>"
    $Start = <line - 20>
    $End = <line + 80>

    $i = 1
    Get-Content $Path | ForEach-Object {
        if ($i -ge $Start -and $i -le $End) {
            "{0,4}: {1}" -f $i, $_
        }
        $i++
    }

Recommended exact targets:

- `creation.rs` around `MlsGroup::new_with_group_id`
- `membership.rs` around `add_members`
- `creation.rs` around `new_from_welcome`
- `messages/mod.rs` around `Welcome`
- `key_packages/key_package_in.rs` around `KeyPackageIn`
- `key_packages/mod.rs` around `KeyPackage`

## 13. Success criteria for this recon checkpoint

This recon checkpoint succeeds when:

- docs/57 exists;
- this docs/58 recon note is committed;
- no code changes are made;
- no unsupported conversation command is moved to supported;
- no Comms/Cypher/trust runtime integration is started;
- the next implementation target is narrowed to conversation-create planning/implementation.

## 14. Non-goals preserved

This recon checkpoint does not implement:

- conversation-create;
- conversation-add-member;
- conversation-join;
- Welcome artifact writing;
- KeyPackage artifact consumption;
- message-protect;
- message-open;
- Comms runtime integration;
- Cypher routing;
- trust-state mutation;
- production vault storage;
- hostile-server proof;
- replay resistance;
- metadata privacy;
- Android;
- CarbonStackOS.
