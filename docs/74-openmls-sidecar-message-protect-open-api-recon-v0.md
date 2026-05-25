# OpenMLS Sidecar Message Protect/Open API Recon v0

Status: API recon / pre-implementation result
Component: CarbonStackComms / OpenMLS sidecar
Phase: Phase 2D
Previous docs:
- docs/72-openmls-sidecar-conversation-join-result-v0.md
- docs/73-openmls-sidecar-message-protect-open-skeleton-v0.md

## 1. Purpose

This document records targeted OpenMLS API reconnaissance for the next implementation checkpoint:

    message-protect --device-label <sender-device> --conversation-label <conversation> --plaintext <text>

    message-open --device-label <receiver-device> --conversation-label <conversation> --message <path>

This checkpoint is docs/recon only. It does not implement message-protect or message-open.

## 2. Current project state before implementation

v0.2.34 validated the first dev-local MLS membership onboarding lifecycle:

    create -> add-member -> Welcome export -> join

Validated current sidecar flow:

1. Alice identity-create.
2. Bob identity-create.
3. Bob public-bundle-export --write-artifact.
4. Bob public-bundle-export saves device provider-storage.json because Bob later needs the private KeyPackage bundle for Welcome consumption.
5. Alice conversation-create.
6. Alice conversation-add-member with Bob public KeyPackage artifact.
7. Alice Welcome carrier export as welcome.bin.
8. Bob conversation-join with welcome.bin.
9. Bob joined group state is written under device-scoped conversation state.
10. Bob joined group reload proof succeeds.
11. Alice/add-member and Bob/join report the same group_id_ref.

message-protect and message-open are not implemented yet.

## 3. API recon summary

OpenMLS 0.8.1 exposes the application-message protect path as:

    MlsGroup::create_message(
        &mut self,
        provider: &Provider,
        signer: &impl Signer,
        message: &[u8],
    ) -> Result<MlsMessageOut, CreateMessageError>

The protected output is an MlsMessageOut.

OpenMLS 0.8.1 exposes the application-message open/process path as:

    MlsGroup::process_message(
        &mut self,
        provider: &Provider,
        message: impl Into<ProtocolMessage>,
    ) -> Result<ProcessedMessage, ProcessMessageError<Provider::StorageError>>

The expected application plaintext path is:

    MlsMessageIn::tls_deserialize_exact_bytes(message_bytes)
    try_into_protocol_message()
    group.process_message(provider, protocol_message)
    processed_message.into_content()
    ProcessedMessageContent::ApplicationMessage(application_message)
    application_message.into_bytes()

Existing scratch history already validated this pattern at least in the Rust scratch/minimal path.

## 4. Protect path details

The protect operation should load sender group state, call create_message, serialize the returned MlsMessageOut, write a protected application message artifact, then save sender provider storage.

Planned v0 implementation shape:

    let message_out = group.create_message(&provider, &signer, plaintext.as_bytes())?;

    let message_bytes = message_out.to_bytes()?;

    write application-message.bin

    provider.save_storage_to_path(...)?;

Important:

create_message takes &mut self.

That means message-protect must be treated as state-mutating, not read-only.

Even if the first implementation does not observe visible epoch changes, the provider/group state must be saved after message-protect unless OpenMLS docs/source prove it is unnecessary. The safer v0 behavior is to save sender provider storage after protect.

## 5. Open path details

The open operation should load receiver group state, read the protected application message artifact, deserialize it as MlsMessageIn, convert it to ProtocolMessage, process it, extract plaintext, then save receiver provider storage.

Planned v0 implementation shape:

    let message_in = MlsMessageIn::tls_deserialize_exact_bytes(&message_bytes)?;

    let protocol_message = message_in.try_into_protocol_message()?;

    let processed_message = group.process_message(&provider, protocol_message)?;

    let plaintext = match processed_message.into_content() {
        ProcessedMessageContent::ApplicationMessage(application_message) => {
            application_message.into_bytes()
        }
        other => return Err(...)
    };

    provider.save_storage_to_path(...)?;

Important:

process_message takes &mut self.

That means message-open must be treated as state-mutating, not read-only.

## 6. Serialization decision

Protected application message artifact should initially store:

    MlsMessageOut::to_bytes()

The receiving side should consume it as:

    MlsMessageIn::tls_deserialize_exact_bytes(...)

or the closest available OpenMLS/tls_codec equivalent if exact_bytes is not available in the sidecar import context.

This mirrors the existing scratch pattern:

    MlsMessageOut serialized
    MlsMessageIn deserialized
    MlsMessageIn converted to ProtocolMessage
    process_message called

Do not invent a custom CarbonStack binary format in this rung.

Initial artifact name:

    application-message.bin

Initial sanitized metadata:

    message-manifest.json
    message-protect-summary.json
    message-open-summary.json

## 7. State mutation and persistence decision

Both outbound protect and inbound open should be treated as persistence-relevant.

Reason:

- create_message takes &mut self.
- process_message takes &mut self.
- Prior scratch continuity recorded that create_message required mutable Alice group state.
- Prior scratch continuity recorded that process_message required mutable Bob group state.
- Future multi-message continuity depends on saving state after outbound and inbound message operations.

Therefore v0 implementation should:

- load sender provider storage before message-protect;
- save sender provider storage after message-protect;
- reload-prove sender group if practical;
- load receiver provider storage before message-open;
- save receiver provider storage after message-open;
- reload-prove receiver group if practical.

## 8. State layout issue

Current state layout after v0.2.34:

Alice conversation-create/add-member path:

    .carbonstack-openmls-sidecar-state/dev/conversations/<conversation-label>/

Bob conversation-join path:

    .carbonstack-openmls-sidecar-state/dev/devices/<device-label>/conversations/<conversation-label>/

This asymmetry is known.

For the first message proof, there are two viable implementation strategies.

### Option A: Keep current asymmetry for one proof rung

message-protect loads Alice from:

    dev/conversations/<conversation-label>/provider-storage.json

message-open loads Bob from:

    dev/devices/<bob-device>/conversations/<conversation-label>/provider-storage.json

Pros:

- least disruptive;
- uses validated v0.2.34 paths;
- avoids migration before message API is proven.

Cons:

- awkward model;
- Alice and Bob use different path conventions;
- future multi-device logic will need cleanup.

### Option B: Migrate Alice to device-scoped conversation state first

conversation-create writes Alice under:

    dev/devices/<alice-device>/conversations/<conversation-label>/

Pros:

- cleaner model;
- protect/open symmetric;
- future conversation commands become device-centered.

Cons:

- touches conversation-create, conversation-load-check, conversation-add-member, tests, and docs;
- risks destabilizing the newly validated membership lifecycle.

Recommended implementation boundary:

Use Option A for the first message-protect/open proof unless implementation becomes too ugly. Document the asymmetry loudly. Migrate Alice to device-scoped state in a later cleanup checkpoint after message APIs are proven.

## 9. Proposed command behavior

### message-protect

Proposed command:

    message-protect --device-label <sender-device> --conversation-label <conversation> --plaintext <text>

Initial sender:

    carbonstack-alice-device

Initial sender state source:

    dev/conversations/<conversation-label>/provider-storage.json

Expected output artifact:

    dev/conversations/<conversation-label>/messages/<message-id>/application-message.bin

or, if generated message IDs are deferred:

    dev/conversations/<conversation-label>/application-message.bin

Recommended v0 path:

    dev/conversations/<conversation-label>/messages/message-0001/application-message.bin

Expected success fields:

    ok=true
    command=message-protect
    phase=phase2d-message-protect-dev
    private_material_included=false
    message_protected=true
    protected_message_written=true
    provider_storage_loaded=true
    provider_storage_written=true
    group_reloadable=true
    device_label
    conversation_label
    group_id_ref
    member_count
    epoch_before
    epoch_after
    message_artifact_path_hint
    message_manifest_path_hint
    message_artifact_sha256
    message_artifact_size_bytes
    state_scope=dev-local-sidecar-state

Events:

    message.protected
    storage.saved

### message-open

Proposed command:

    message-open --device-label <receiver-device> --conversation-label <conversation> --message <path>

Initial receiver:

    carbonstack-bob-device

Initial receiver state source:

    dev/devices/<receiver-device>/conversations/<conversation-label>/provider-storage.json

Expected input artifact:

    application-message.bin

Expected output metadata:

    message-open-summary.json

Expected success fields:

    ok=true
    command=message-open
    phase=phase2d-message-open-dev
    private_material_included=false
    message_opened=true
    provider_storage_loaded=true
    provider_storage_written=true
    group_reloadable=true
    device_label
    conversation_label
    group_id_ref
    member_count
    epoch_before
    epoch_after
    message_artifact_path_hint
    plaintext_utf8
    plaintext_len
    state_scope=dev-local-sidecar-state

Events:

    message.opened
    storage.saved

## 10. Error/failure facts from recon

Relevant OpenMLS error surface includes:

- pending proposals can block create_message;
- application messages must be encrypted;
- non-member application messages fail;
- external application messages are not permitted in the ordinary internal member path;
- message decryption can fail;
- messages from too-old epochs can fail.

Initial sidecar failure classes should include:

    missing_required_argument
    invalid_device_label
    invalid_conversation_label
    identity_missing
    conversation_missing
    provider_storage_missing
    provider_storage_load_failed
    provider_group_unavailable
    plaintext_missing
    plaintext_too_large
    message_artifact_exists
    message_artifact_missing
    message_artifact_invalid
    message_protect_failed
    message_open_failed
    unexpected_processed_message_type
    plaintext_not_utf8

## 11. Security/stdout rules

Never print:

- signer.json;
- provider-storage.json;
- raw provider storage;
- raw group state;
- raw protected message bytes;
- raw Welcome bytes;
- raw KeyPackage bytes;
- private key material;
- key store contents;
- MemoryStorage JSON.

Allowed to print:

- path hints;
- hashes;
- sizes;
- member counts;
- epoch strings;
- group refs;
- bounded dev plaintext from message-open only.

Plaintext in message-open stdout is allowed only for the explicit dev proof command. It must not be described as final product UX.

## 12. Required implementation tests

Provider-info:

- message-protect supported after implementation;
- message-open supported after implementation;
- unsupported-command test moves to state-checkpoint or state-load-check.

message-protect tests:

- missing --device-label fails;
- missing --conversation-label fails;
- missing --plaintext fails;
- invalid device label fails;
- invalid conversation label fails;
- missing conversation fails;
- success writes protected application message artifact;
- success writes message manifest/summary;
- success reports provider_storage_loaded=true;
- success reports provider_storage_written=true;
- success reports group_reloadable=true;
- stdout contains no private material;
- duplicate message artifact refuses overwrite.

message-open tests:

- missing --device-label fails;
- missing --conversation-label fails;
- missing --message fails;
- invalid labels fail;
- missing joined conversation fails;
- missing message artifact fails;
- invalid message artifact fails;
- success opens protected message;
- success returns expected dev plaintext;
- success reports provider_storage_loaded=true;
- success reports provider_storage_written=true;
- success reports group_reloadable=true;
- stdout contains no private material beyond explicit bounded dev plaintext.

One-way lifecycle test:

1. remove state;
2. Alice identity-create;
3. Bob identity-create;
4. Bob public-bundle-export --write-artifact;
5. Alice conversation-create;
6. Alice conversation-add-member with Bob KeyPackage artifact;
7. Bob conversation-join with welcome.bin;
8. Alice message-protect --plaintext "hello bob";
9. Bob message-open --message <application-message.bin>;
10. assert plaintext_utf8 is "hello bob";
11. assert no forbidden secret tokens appear in stdout.

## 13. Implementation boundary recommendation

v0.2.36 can implement either:

Option 1:

    message-protect only

Option 2:

    message-protect and message-open together

Recommendation:

Implement protect/open together only if the first implementation stays narrow and tests remain focused. The useful proof is end-to-end plaintext recovery by Bob; implementing protect alone is less valuable.

Do not implement:

- Comms runtime integration;
- Cypher routing;
- trust-state mutation;
- multi-message continuity;
- device-scoped Alice migration;
- production secure vault;
- Android;
- CarbonStackOS.

## 14. Open questions for implementation

v0.2.36 still must verify in code:

1. exact imports for ProcessedMessageContent;
2. exact availability of MlsMessageIn::tls_deserialize_exact_bytes in the sidecar crate context;
3. whether MlsMessageOut::to_bytes is the cleanest serialization method for application messages;
4. whether message-protect should use old Alice global state for v0 or migrate first;
5. whether create_message changes epoch or only secret/ratchet state;
6. whether process_message changes epoch or only secret/ratchet state;
7. whether a one-message proof survives separate sidecar invocations with saved provider storage;
8. how duplicate message IDs should be generated/refused;
9. whether plaintext should be UTF-8 only in v0 or raw bytes with a text-only test path.

## 15. Next checkpoint

Next recommended checkpoint:

    v0.2.36 message-protect/open implementation

First implementation target:

    Alice protects "hello bob"
    Bob opens it
    Bob plaintext equals "hello bob"
    both sides save provider storage
    both sides reload-prove group state
