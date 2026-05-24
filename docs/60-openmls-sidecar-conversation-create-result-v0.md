# OpenMLS Sidecar Conversation Create Result v0

Status: Validated dev-sidecar checkpoint
Component: CarbonStackComms / OpenMLS sidecar
Phase: Phase 2D
Plan doc:
- docs/59-openmls-sidecar-conversation-create-plan-v0.md

## 1. Summary

This checkpoint implements and validates the first dev-only OpenMLS conversation lifecycle command:

    conversation-create --device-label <safe> --conversation-label <safe>

This command creates a local one-member OpenMLS group/conversation for an existing dev-only sidecar identity and writes sanitized conversation metadata.

This is not group messaging support. It does not add members, consume KeyPackage artifacts, write Welcome artifacts, join another device, protect messages, open messages, wire into CarbonStackComms runtime, route through CarbonStackCypher, or mutate trust-state storage.

## 2. Implemented command

New supported command:

    conversation-create --device-label <safe-device-label> --conversation-label <safe-conversation-label>

The command:

- validates the device label;
- validates the conversation label;
- requires existing dev-only identity state;
- loads signer material locally without printing it;
- recreates the local BasicCredential / CredentialWithKey path for the device identity;
- creates an OpenMLS provider;
- builds MlsGroupCreateConfig with the current Phase 2D ciphersuite;
- enables the ratchet tree extension;
- derives deterministic dev-local GroupId bytes from the conversation label;
- calls MlsGroup::new_with_group_id;
- writes sanitized conversation-summary.json;
- returns sanitized stdout;
- refuses duplicate conversation creation.

## 3. New conversation label validation

A separate conversation label validator was added.

Allowed shape:

- non-empty;
- not "." or "..";
- max length 96;
- ASCII alphanumeric, hyphen, and underscore only.

Rejected:

- path separators;
- spaces;
- dots;
- relative paths;
- Unicode/confusable labels;
- unsupported punctuation.

Rust tests cover accepted and rejected device/conversation labels.

## 4. Dev-state layout

Existing device state remains:

    .carbonstack-openmls-sidecar-state/dev/devices/<device-label>/

New conversation state lives under:

    .carbonstack-openmls-sidecar-state/dev/conversations/<conversation-label>/

Current generated conversation file:

    conversation-summary.json

The summary is sanitized metadata. It is generated dev state and must not be treated as production secure storage.

## 5. conversation-summary.json

The summary includes:

    summary_version
    conversation_label
    creator_device_label
    state_scope
    ciphersuite
    group_id_ref
    group_id_len
    member_count
    epoch
    conversation_created
    provider_storage_written
    private_material_included
    warning

The summary does not include:

- signer material;
- private keys;
- OpenMLS epoch secrets;
- raw provider storage;
- raw MlsGroup serialization;
- Welcome bytes;
- message content.

## 6. Provider storage semantics

conversation-create reports:

    provider_storage_written=true

Reason:

OpenMLS group creation is stateful and this sidecar rung now explicitly acknowledges that dev-local OpenMLS group/provider state exists conceptually for the created group.

Important limitation:

This is still dev-local provider/group state. It is not a production secure vault, not hardware-backed storage, and not a final persistence design.

## 7. Provider-info changes

conversation-create graduated from unsupported to supported.

Provider capabilities now include:

    provider-info
    identity-create
    identity-status
    public-bundle-export
    conversation-create

Unsupported commands still include:

    conversation-add-member
    conversation-join
    message-protect
    message-open
    state-checkpoint
    state-load-check

The unsupported-command test target moved from conversation-create to conversation-add-member.

## 8. Events and trust mapping

New sidecar event emitted on success:

    provider.conversation.created

New sidecar event emitted on duplicate refusal:

    provider.conversation.exists

Current classification:

    provider.conversation.created
      severity: info
      trust_relevant: false

    provider.conversation.exists
      severity: warning
      trust_relevant: false

Current trust behavior:

- append history;
- debug-only;
- not user visible;
- does not block send/open;
- does not require reverify.

This is intentionally conservative for the dev-sidecar rung. Future member-add/join events will need more careful trust treatment because membership changes are user/security meaningful.

## 9. Validated behavior

Manual probe validated:

    provider-info

Result:

- conversation-create appears as supported;
- conversation-add-member remains unsupported;
- conversation-join remains unsupported.

Manual probe validated missing identity:

    conversation-create --device-label carbonstack-alice-device --conversation-label carbonstack-test-conversation

Result:

- fails with identity_missing;
- emits provider.identity.missing;
- exits nonzero;
- private_material_included=false.

Manual probe validated success after identity-create:

    identity-create --device-label carbonstack-alice-device
    conversation-create --device-label carbonstack-alice-device --conversation-label carbonstack-test-conversation

Result:

- succeeds;
- conversation_created=true;
- member_count=1;
- epoch=GroupEpoch(0);
- group_id_ref=sha256:<hex>;
- provider_storage_written=true;
- private_material_included=false;
- writes conversation-summary.json.

Manual probe validated duplicate refusal:

    conversation-create --device-label carbonstack-alice-device --conversation-label carbonstack-test-conversation

Result:

- fails with conversation_already_exists;
- emits provider.conversation.exists;
- does not overwrite state;
- private_material_included=false.

Manual probe validated invalid conversation label:

    conversation-create --device-label carbonstack-alice-device --conversation-label ../bad

Result:

- fails with invalid_conversation_label;
- emits provider.command.invalid;
- private_material_included=false.

## 10. Test coverage added / updated

Go-side tests now cover:

- provider-info lists conversation-create as supported;
- provider-info keeps conversation-add-member and conversation-join unsupported;
- unsupported-command envelope test now uses conversation-add-member;
- conversation-create missing identity failure;
- identity-create prerequisite;
- conversation-create success;
- success envelope fields;
- provider.conversation.created event;
- provider_storage_written=true;
- private_material_included=false;
- conversation-summary.json exists;
- summary metadata matches stdout;
- duplicate conversation-create refusal;
- provider.conversation.exists event;
- invalid conversation label failure;
- no obvious secret material in stdout.

Provider event/trust mapping was updated for:

    provider.conversation.created
    provider.conversation.exists

Rust label tests cover conversation label validation.

## 11. Still not implemented

This checkpoint does not implement:

- conversation-add-member;
- KeyPackage artifact consumption;
- Welcome artifact export;
- conversation-join;
- message-protect;
- message-open;
- Comms runtime integration;
- Cypher MLS routing;
- trust-state storage mutation;
- production provider storage;
- secure vault design;
- hardware-backed identity;
- hostile-server harness;
- replay resistance;
- metadata privacy;
- Android;
- CarbonStackOS.

## 12. Allowed claims after this checkpoint

Allowed:

- The sidecar can create a dev-local one-member OpenMLS group/conversation for an existing dev identity.
- conversation-create validates device and conversation labels.
- conversation-create writes sanitized conversation-summary.json.
- conversation-create reports member_count=1 and epoch metadata.
- conversation-create refuses duplicate state.
- provider.conversation.created and provider.conversation.exists are typed/classified/trust-mapped.
- Go-side tests validate conversation-create success and negative paths.
- No Comms/Cypher/trust runtime integration exists.

Not allowed:

- CarbonStack supports user-facing group chat.
- CarbonStack can add members.
- CarbonStack can generate or consume Welcome artifacts.
- CarbonStack can protect/open runtime messages.
- CarbonStack is production secure or Signal-equivalent.
- Dev provider storage is a secure vault.
- Conversation state is production persistence.

## 13. Recommended next checkpoint

Next safest checkpoint:

    conversation-add-member / Welcome export planning

Recommended next doc:

    docs/61-openmls-sidecar-conversation-add-member-welcome-plan-v0.md

That plan should decide:

- how Bob's public-bundle.keypackage.bin is deserialized;
- whether raw KeyPackage artifacts are consumed directly or wrapped;
- where welcome.bin is written;
- welcome-manifest.json fields;
- duplicate Welcome export behavior;
- provider storage implications after add_members and merge_pending_commit;
- event/trust mapping for membership changes;
- tests for missing/invalid KeyPackage artifacts;
- no-secret stdout boundaries.

Do not jump directly to message-protect/message-open.
Do not jump directly to Comms runtime integration.
Do not jump directly to Cypher routing.
Do not jump directly to trust-state mutation.
