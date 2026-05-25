# OpenMLS Sidecar Message Protect/Open Skeleton v0

Status: Skeleton / pre-implementation contract
Component: CarbonStackComms / OpenMLS sidecar
Phase: Phase 2D
Previous docs:
- docs/69-openmls-sidecar-add-member-welcome-export-result-v0.md
- docs/70-openmls-sidecar-conversation-join-skeleton-v0.md
- docs/71-openmls-sidecar-conversation-join-api-recon-v0.md
- docs/72-openmls-sidecar-conversation-join-result-v0.md

## 1. Purpose

This document defines the next command surface after the v0.2.34 conversation-join checkpoint.

v0.2.34 validated the first dev-local MLS membership onboarding lifecycle:

    create -> add-member -> Welcome export -> join

The next lifecycle step is application message protection/opening.

This document is intentionally a skeleton and planning contract. It does not implement message-protect or message-open.

## 2. Current state before implementation

Validated sidecar flow:

1. Alice identity exists.
2. Bob identity exists.
3. Bob exports a public KeyPackage artifact.
4. Bob's public-bundle export saves Bob device provider storage.
5. Alice creates a conversation.
6. Alice adds Bob's KeyPackage.
7. Alice exports a Welcome carrier artifact as welcome.bin.
8. Bob consumes welcome.bin.
9. Bob joins the group.
10. Bob saves device-scoped joined conversation provider storage.
11. Bob joined group reload proof succeeds.
12. Alice/add-member and Bob/join report the same group_id_ref.

Known non-goals still in force:

- no production E2EE claim;
- no Comms runtime integration;
- no Cypher MLS routing;
- no trust-state mutation;
- no production secure vault;
- no Android/CarbonStackOS integration.

## 3. New command targets

The first message commands should be dev-sidecar commands only.

Candidate command A:

    message-protect --device-label <sender-device> --conversation-label <conversation> --plaintext <text>

Candidate command B:

    message-open --device-label <receiver-device> --conversation-label <conversation> --message <path>

Alternative names to consider before implementation:

    conversation-message-protect
    conversation-message-open

or:

    application-message-protect
    application-message-open

Initial recommendation:

    message-protect
    message-open

Reason:

- provider-info already lists message-protect/message-open as unsupported;
- the command names are short;
- they map directly to MLS operation intent;
- tests can move unsupported-command target to state-checkpoint after implementation.

## 4. Proposed v0 input style

For message-protect:

    --device-label <sender-device>
    --conversation-label <conversation>
    --plaintext <text>

For message-open:

    --device-label <receiver-device>
    --conversation-label <conversation>
    --message <path>

Deferred:

    --plaintext-file <path>
    --message-manifest <path>
    stdin plaintext
    binary stdout
    multi-message batch mode
    app-level conversation IDs
    Cypher envelope routing

## 5. Artifact format target

message-protect should write a protected MLS application message artifact under ignored dev state.

Candidate Alice output path:

    .carbonstack-openmls-sidecar-state/dev/conversations/<conversation-label>/messages/<message-id>/application-message.bin

or, if Alice is migrated to device-scoped state:

    .carbonstack-openmls-sidecar-state/dev/devices/<device-label>/conversations/<conversation-label>/messages/<message-id>/application-message.bin

Associated sanitized metadata:

    message-protect-summary.json
    message-manifest.json

message-open should consume:

    application-message.bin

and may write:

    message-open-summary.json

Do not print raw protected message bytes.
Do not print raw group state.
Do not print provider storage.
Plaintext output must be explicit and bounded for dev testing only.

## 6. Critical design issue: Alice state layout

Current Alice conversation state still uses the older global path:

    .carbonstack-openmls-sidecar-state/dev/conversations/<conversation-label>/

Bob joined state uses device-scoped path:

    .carbonstack-openmls-sidecar-state/dev/devices/<device-label>/conversations/<conversation-label>/

This is acceptable for v0.2.34 but becomes awkward for message-protect/open.

Potential options:

### Option A: Keep asymmetry for one message proof

Alice message-protect loads from the old global conversation path.
Bob message-open loads from the new device-scoped conversation path.

Pros:

- least migration risk;
- builds directly on validated v0.2.34 state.

Cons:

- weird model;
- future multi-device/member work gets confusing;
- Alice and Bob use different state path conventions.

### Option B: Migrate Alice conversation-create to device-scoped state before message-protect/open

New Alice path:

    dev/devices/<alice-device>/conversations/<conversation-label>/

Pros:

- cleaner long-term model;
- all participant conversation state becomes device-scoped;
- message-protect/open tests become more symmetric.

Cons:

- migration touches conversation-create, conversation-load-check, conversation-add-member, and existing tests;
- risks destabilizing the recently validated add-member/join lifecycle.

### Option C: Add device-scoped mirror for Alice without removing old path

conversation-create writes both:

    dev/conversations/<conversation-label>/
    dev/devices/<alice-device>/conversations/<conversation-label>/

Pros:

- allows new commands to use device-scoped state;
- preserves old tests initially.

Cons:

- duplicated provider storage;
- risk of divergence;
- not ideal as a long-term model.

Initial recommendation:

Use v0.2.35 recon to decide. Do not migrate in the skeleton doc. If implementation begins before migration, use Option A for the first one-way proof and document the asymmetry clearly.

## 7. Proposed operation sequence: message-protect

Possible sequence:

1. validate device label;
2. validate conversation label;
3. parse plaintext argument;
4. reject empty plaintext;
5. reject overly large plaintext;
6. locate sender conversation provider storage;
7. load sender provider storage;
8. load sender group;
9. call OpenMLS application message protection API;
10. write protected application message artifact;
11. save updated sender provider storage if the API advances state;
12. reload sender group if possible;
13. write message-protect summary/manifest;
14. return sanitized stdout with path/hash/size metadata.

Expected success fields:

    ok=true
    command=message-protect
    phase=phase2d-message-protect-dev
    private_material_included=false
    message_protected=true
    protected_message_written=true
    provider_storage_loaded=true
    provider_storage_written=<true if state advances>
    group_reloadable=true
    conversation_label
    device_label
    message_artifact_path_hint
    message_manifest_path_hint
    message_artifact_sha256
    message_artifact_size_bytes
    epoch_before
    epoch_after
    member_count
    group_id_ref

Events:

    message.protected
    storage.saved

## 8. Proposed operation sequence: message-open

Possible sequence:

1. validate device label;
2. validate conversation label;
3. parse message artifact path;
4. validate message artifact path;
5. locate receiver device-scoped conversation provider storage;
6. load receiver provider storage;
7. load receiver group;
8. deserialize protected application message artifact;
9. call OpenMLS application message open/process API;
10. extract application plaintext;
11. save updated receiver provider storage if the API advances state;
12. reload receiver group if possible;
13. write message-open summary;
14. return sanitized stdout.

Expected success fields:

    ok=true
    command=message-open
    phase=phase2d-message-open-dev
    private_material_included=false
    message_opened=true
    provider_storage_loaded=true
    provider_storage_written=<true if state advances>
    group_reloadable=true
    conversation_label
    device_label
    message_artifact_path_hint
    plaintext_utf8=<dev bounded plaintext only>
    plaintext_len
    epoch_before
    epoch_after
    member_count
    group_id_ref

Events:

    message.opened
    storage.saved

## 9. Security and stdout rules

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

- refs/hashes;
- sizes;
- path hints;
- member count;
- epoch string;
- bounded plaintext only for explicit dev message-open proof.

Plaintext in stdout is acceptable only for this dev proof command and must be documented as test-only.

## 10. Proposed failure classes

Shared:

    missing_required_argument
    invalid_device_label
    invalid_conversation_label
    identity_missing
    conversation_missing
    provider_storage_missing
    provider_group_unavailable
    provider_storage_load_failed
    provider_storage_save_failed

message-protect-specific:

    plaintext_missing
    plaintext_too_large
    message_protect_failed
    message_artifact_exists

message-open-specific:

    invalid_message_path
    message_artifact_missing
    message_artifact_invalid
    message_open_failed
    plaintext_not_utf8

## 11. Required API recon questions

v0.2.35 must answer:

1. What OpenMLS 0.8.1 API protects application messages?
2. What OpenMLS 0.8.1 API opens/processes application messages?
3. What exact message type is returned by protect?
4. What serialization method should be used for protected application message artifacts?
5. What deserialization method should message-open use?
6. Does protect advance sender group state immediately?
7. Does open/process advance receiver group state immediately?
8. Does either side need to merge/process pending proposals/commits for simple application messages?
9. What imports are needed?
10. Does the sender require signer access for application message protect?
11. Does message-open return plaintext bytes directly or through processed message content?
12. How should application messages be distinguished from commits/welcomes/proposals?
13. Can Bob open Alice's first message immediately after join with current group state?
14. Should Alice state be migrated to device-scoped before message implementation?
15. What exact test flow should prove one-way message delivery?

## 12. Required tests after implementation

Provider-info:

- message-protect supported after implementation;
- message-open supported after implementation;
- unsupported-command test moves to state-checkpoint or state-load-check.

message-protect:

- missing device label fails;
- missing conversation label fails;
- missing plaintext fails;
- invalid device label fails;
- invalid conversation label fails;
- missing conversation fails;
- success writes protected message artifact;
- success saves provider storage if required;
- success emits no private material;
- duplicate artifact refuses overwrite.

message-open:

- missing device label fails;
- missing conversation label fails;
- missing message path fails;
- invalid labels fail;
- missing joined conversation fails;
- invalid artifact fails;
- success opens protected message;
- success returns expected dev plaintext;
- success saves provider storage if required;
- success emits no private material.

Full one-way integration test:

1. remove state;
2. Alice identity-create;
3. Bob identity-create;
4. Bob public-bundle-export --write-artifact;
5. Alice conversation-create;
6. Alice conversation-add-member with Bob KeyPackage;
7. Bob conversation-join with welcome.bin;
8. Alice message-protect --plaintext "hello bob";
9. Bob message-open --message <artifact>;
10. assert plaintext is "hello bob";
11. assert no private material appears in stdout.

## 13. Result doc requirements

The implementation result doc must record:

- exact protect API;
- exact open/process API;
- exact protected artifact type and serialization format;
- exact sender state path;
- exact receiver state path;
- whether sender state advances;
- whether receiver state advances;
- whether Alice state remained global or migrated to device-scoped;
- exact tests run;
- allowed claims;
- not allowed claims;
- all blunders and API surprises.

## 14. Recommended next checkpoint

Next checkpoint after this skeleton:

    docs/74-openmls-sidecar-message-protect-open-api-recon-v0.md

Do not implement message-protect/message-open until the API path is documented.
