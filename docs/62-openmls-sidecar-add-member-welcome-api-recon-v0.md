# OpenMLS Sidecar Add-Member / Welcome API Recon and Code-State v0

Status: Reconnaissance / code-state checkpoint
Component: CarbonStackComms / OpenMLS sidecar
Phase: Phase 2D
Related plan:
- docs/61-openmls-sidecar-conversation-add-member-welcome-plan-v0.md

## 1. Purpose

This document records the current deployed sidecar code state and targeted API reconnaissance before implementing conversation-add-member / Welcome export.

This is intentionally a planning/recon checkpoint.

No add-member implementation is included here.

## 2. Repo state at recon start

Observed repo heads:

- carbonstack: 57a4789 docs: record OpenMLS conversation create result
- carbonstack-comms: 84fd5c1 feat: add OpenMLS sidecar conversation create
- carbonstack-cypher: 0bfd5af chore: remove tracked Cypher local state artifacts
- carbonstack-os: b537475 Add CarbonStackOS north star and initial appliance model

## 3. Current sidecar supported/unsupported state

Current sidecar supported commands:

    provider-info
    identity-create
    identity-status
    public-bundle-export
    conversation-create

Current sidecar unsupported commands:

    conversation-add-member
    conversation-join
    message-protect
    message-open
    state-checkpoint
    state-load-check

This is correct before add-member implementation.

## 4. Current sidecar command routing

Current `main.rs` routes:

    provider-info
    identity-create
    identity-status
    public-bundle-export
    conversation-create

`conversation-add-member` is not routed. It falls through to unsupported command handling.

Current `WARNINGS` still state:

- OpenMLS is not wired into CarbonStackComms.
- Cypher does not route MLS payloads.
- trust-state storage does not consume provider events.
- identity-create writes dev-only secret-bearing signer state but never prints private material.

## 5. Current conversation-create behavior

Current `conversation-create`:

    conversation-create --device-label <safe> --conversation-label <safe>

It:

- parses device label;
- parses conversation label;
- validates device label;
- validates conversation label;
- calls `create_dev_conversation`;
- emits sanitized JSON;
- writes conversation-summary.json;
- reports provider_storage_written=true;
- emits provider.conversation.created on success;
- emits provider.conversation.exists on duplicate refusal;
- does not add members;
- does not write Welcome artifacts;
- does not join Bob;
- does not protect/open messages.

## 6. Current label validation

Current `labels.rs` provides:

    validate_device_label(label)
    validate_conversation_label(label)

Both delegate to shared safe-label validation.

Current allowed label characters:

- ASCII alphanumeric;
- hyphen;
- underscore.

Rejected:

- empty labels;
- "." and "..";
- length over 96;
- slashes;
- backslashes;
- spaces;
- dots;
- unsupported punctuation;
- Unicode/non-ASCII characters.

## 7. Current Go contract state

Current Go sidecar test envelope data includes conversation fields:

    conversation_label
    conversation_created
    conversation_state_path_hint
    conversation_summary_path_hint
    ciphersuite
    group_id_ref
    group_id_len
    member_count
    epoch

Current provider-info tests assert:

- provider-info supported;
- identity-create supported;
- identity-status supported;
- public-bundle-export supported;
- conversation-create supported;
- conversation-add-member unsupported;
- conversation-join unsupported;
- message-protect unsupported;
- message-open unsupported.

Current unsupported-command test uses:

    conversation-add-member

This is correct before add-member implementation, but must be changed when add-member graduates.

## 8. Current event/trust state

Current sidecar success event for conversation-create:

    provider.conversation.created

Current duplicate event:

    provider.conversation.exists

Current trust mapping treats conversation-create and duplicate refusal as non-trust-relevant dev-sidecar setup/history/debug events.

Future add-member and Welcome events should be added carefully because membership changes are product-level trust-sensitive.

## 9. OpenMLS add_members API recon

Targeted local inspection confirms OpenMLS 0.8.1 exposes:

    pub fn add_members<Provider: OpenMlsProvider>(
        &mut self,
        provider: &Provider,
        signer: &impl Signer,
        key_packages: &[KeyPackage],
    ) -> Result<...>

Important consequences:

- add-member requires mutable MlsGroup state.
- add-member requires the creator device signer.
- add-member requires the incoming member public KeyPackage as `KeyPackage`.
- add-member likely mutates local group/provider state.
- v0.2.29 cannot be implemented honestly unless the sidecar can reload or persist the creator's MlsGroup state.

## 10. Welcome API recon

Targeted local inspection identifies `Welcome` in:

    openmls-0.8.1/src/messages/mod.rs

The prior scratch code confirms Welcome is extracted from the message wrapper:

    let welcome = match welcome_msg.body() {
        MlsMessageBodyOut::Welcome(welcome) => welcome.clone(),
        other => panic!("expected Welcome MlsMessageOut body, got: {:?}", other),
    };

Important correction from prior scratch history:

- Do not serialize the whole `MlsMessageOut` and try to deserialize it directly as `Welcome`.
- The Welcome body must be extracted from the wrapper first.

Implementation still needs exact line-window confirmation for Welcome TLS serialization.

## 11. KeyPackage input recon

Targeted local inspection identifies:

    openmls-0.8.1/src/key_packages/mod.rs
    openmls-0.8.1/src/key_packages/key_package_in.rs

Relevant types:

    KeyPackage
    KeyPackageIn

Current known facts:

- v0.2.25 writes a serialized public KeyPackage artifact as `public-bundle.keypackage.bin`.
- `MlsGroup::add_members` expects `KeyPackage`, not `KeyPackageBundle`.
- Prior scratch code used `key_package_bundle.key_package().clone()` before passing to add_members.
- Incoming artifact deserialization likely involves `KeyPackageIn` and validation/conversion to `KeyPackage`.

Open issue:

The exact `KeyPackageIn` TLS-deserialize / validate / convert path must be inspected before implementation.

## 12. Prior scratch pattern

The prior `openmls-minimal` research crate demonstrates the expected local flow:

    Alice identity/setup
    Bob identity/setup
    Alice MlsGroup::new_with_group_id
    Bob KeyPackage creation
    Alice add_members(&alice_provider, &alice.signer, &[bob.key_package.clone()])
    extract MlsMessageBodyOut::Welcome
    Alice merge_pending_commit
    Bob StagedWelcome::new_from_welcome
    Bob into_group
    later create_message / process_message

Only the Alice add-member / Welcome extraction portion is in scope for the next implementation.

Bob join remains later.

## 13. Current implementation blocker

The largest risk is state continuity.

v0.2.27 conversation-create writes sanitized `conversation-summary.json`, but add-member requires a usable mutable MlsGroup.

The next implementation must answer:

    Can conversation-create's group state be loaded in a later process?

If not, the next implementation must repair conversation-create persistence before or as part of add-member.

Do not fake add-member by recreating a fresh group with the same label unless the resulting state is explicitly documented as invalid for real lifecycle continuity.

## 14. Implementation-time inspections still required

Before code, inspect tiny line windows for:

- `MlsGroup::add_members` return type and exact returned tuple shape.
- `MlsMessageBodyOut` enum and `Welcome` variant.
- `Welcome` TLS serialization method.
- `KeyPackageIn` TLS deserialization.
- `KeyPackageIn` validation/conversion into `KeyPackage`.
- `MlsGroup::load` / storage API behavior.
- `OpenMlsRustCrypto` provider storage persistence behavior.

Avoid broad recursive terminal floods.

## 15. Planned docs/code path after this recon

Next implementation plan:

    conversation-add-member / Welcome export implementation

Possible implementation files:

    carbonstack-comms/internal/protocol/mls/research/openmls-sidecar/src/main.rs
    carbonstack-comms/internal/protocol/mls/research/openmls-sidecar/src/state.rs
    carbonstack-comms/internal/protocol/openmls_sidecar_provider_info_test.go
    carbonstack-comms/internal/protocol/provider_events.go
    carbonstack-comms/internal/protocol/provider_events_test.go
    carbonstack-comms/internal/protocol/provider_trust.go
    carbonstack-comms/internal/protocol/provider_trust_test.go

Possible result doc:

    docs/63-openmls-sidecar-conversation-add-member-welcome-result-v0.md

If add-member reveals missing persistence:

    docs/63-openmls-sidecar-conversation-create-persistence-repair-plan-v0.md

or a clearly named repair result doc if repair is implemented first.

## 16. Non-goals preserved

This checkpoint does not implement:

- conversation-add-member;
- Welcome export;
- conversation-join;
- message-protect;
- message-open;
- Comms runtime integration;
- Cypher routing;
- trust-state mutation;
- production storage;
- secure vault;
- hostile-server proof;
- replay resistance;
- metadata privacy;
- Android;
- CarbonStackOS.

## 17. Recommendation

Proceed to implementation only after deciding whether v0.2.27 conversation-create persists enough group/provider state.

If the answer is no, do not force add-member.

First repair persistence, then add-member.
