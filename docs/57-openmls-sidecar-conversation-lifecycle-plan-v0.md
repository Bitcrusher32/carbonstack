# OpenMLS Sidecar Conversation Lifecycle Plan v0

Status: Planned
Component: CarbonStackComms / OpenMLS sidecar
Phase: Phase 2D
Previous docs:
- docs/55-openmls-sidecar-keypackage-artifact-export-plan-v0.md
- docs/56-openmls-sidecar-keypackage-artifact-export-result-v0.md

## 1. Purpose

This document defines the next planning rung after dev-only serialized OpenMLS KeyPackage artifact export.

The current sidecar can:

- create dev-only identity/signing material;
- load/check dev-only identity state;
- generate a real OpenMLS KeyPackage in memory;
- write a serialized public KeyPackage artifact under ignored dev state;
- write a sanitized public-bundle manifest;
- refuse duplicate KeyPackage artifact export;
- keep stdout sanitized;
- keep private_material_included=false;
- keep provider_storage_written=false.

This plan prepares the next dev-only conversation lifecycle rung.

The goal is not to wire OpenMLS into CarbonStackComms runtime messaging yet. The goal is to define a safe, narrow, sidecar-only model for:

- creating a local OpenMLS group/conversation;
- adding a member using a serialized KeyPackage artifact;
- writing a Welcome artifact;
- joining a conversation from a Welcome artifact;
- preserving sanitized outputs and no-secret stdout;
- defining tests before implementation.

## 2. Current baseline

Current validated command ladder:

    provider-info
    identity-create --device-label <safe>
    identity-status --device-label <safe>
    public-bundle-export --device-label <safe>
    public-bundle-export --device-label <safe> --write-artifact

Current validated artifact export behavior:

    public-bundle.keypackage.bin
    public-bundle-manifest.json
    public-bundle-summary.json

Current command properties:

- dev-only;
- local sidecar state only;
- ignored generated state;
- no Comms runtime wiring;
- no Cypher routing;
- no trust-state mutation;
- no production E2EE claim.

## 3. Problem

A serialized KeyPackage artifact is now available, but there is still no dev conversation lifecycle.

The sidecar cannot yet:

- create an OpenMLS group/conversation as a named dev-local object;
- consume another device's KeyPackage artifact;
- add a member to a group;
- produce a Welcome artifact;
- join a group from a Welcome artifact;
- persist or summarize group/provider state for later lifecycle steps;
- prove basic Alice/Bob lifecycle through sidecar commands.

Before implementation, CarbonStack needs a planning checkpoint that defines the command surface, artifact names, state layout, duplicate behavior, provider storage assumptions, and tests.

## 4. Non-goals

This rung must not:

- wire OpenMLS into comms send;
- wire OpenMLS into comms inbox;
- route MLS payloads through CarbonStackCypher;
- mutate trust.json;
- mutate trust-events.jsonl;
- implement message-protect;
- implement message-open;
- implement production secure storage;
- implement hardware-backed identity;
- implement Android integration;
- implement CarbonStackOS;
- claim production E2EE;
- claim hostile-server proof;
- claim replay resistance;
- claim metadata privacy;
- treat dev provider state as secure vault storage;
- print or expose signer.json;
- print or expose private key material;
- commit generated sidecar state.

## 5. Recommended command surface

Recommended future commands:

    conversation-create --device-label <safe> --conversation-label <safe>

    conversation-add-member --device-label <safe> --conversation-label <safe> --member-keypackage <path>

    conversation-join --device-label <safe> --welcome <path>

These commands should remain sidecar-only and dev-local.

They should not be invoked by CarbonStackComms user-facing send/inbox flows in this checkpoint.

## 6. Conversation naming

Use a separate safe label type for conversation labels.

Recommended validation rules should mirror device labels initially:

- non-empty;
- limited length;
- lowercase ASCII letters;
- digits;
- hyphen;
- underscore if already allowed;
- no dots;
- no path separators;
- no spaces;
- no Unicode confusables;
- no shell metacharacters.

Recommended example:

    carbonstack-test-conversation

Do not use arbitrary group names or user-facing titles yet.

## 7. Proposed dev-state layout

Under the sidecar dev state root:

    .carbonstack-openmls-sidecar-state/dev/

Current device state:

    devices/<device-label>/
      identity-prep.json
      identity-summary.json
      identity-state.json
      signer.json
      public-bundle-summary.json
      public-bundle-manifest.json
      public-bundle.keypackage.bin

Proposed conversation state:

    conversations/<conversation-label>/
      conversation-summary.json
      group-state.bin or provider-state/
      add-member-summary.json
      welcome.bin
      welcome-manifest.json

Open question:

- Whether OpenMLS provider storage should remain device-scoped, conversation-scoped, or both.

Recommendation:

- Start with explicit dev-local conversation state under conversations/<conversation-label>/.
- Preserve device identity under devices/<device-label>/.
- Do not design production storage yet.
- If OpenMLS provider storage must be written, label it clearly as dev provider state, not secure vault storage.

## 8. Proposed artifact model

### conversation-create

Should write:

    conversation-summary.json

Possible future state files:

    group-state.bin
    provider-state/

The summary should include sanitized metadata only:

    schema
    conversation_label
    creator_device_label
    group_id_ref
    ciphersuite
    member_count
    created_locally
    provider_storage_written
    private_material_included
    warning

### conversation-add-member

Should consume:

    public-bundle.keypackage.bin

Should write:

    add-member-summary.json
    welcome.bin
    welcome-manifest.json

The Welcome artifact should be treated as public/semi-public onboarding material but not dumped to stdout.

The manifest should include:

    schema
    conversation_label
    inviter_device_label
    member_keypackage_ref
    welcome_artifact
    welcome_artifact_sha256
    welcome_artifact_size_bytes
    provider_storage_written
    private_material_included
    warning

### conversation-join

Should consume:

    welcome.bin

Should write:

    joined-conversation-summary.json

Possibly writes provider/group state for the joining device.

The summary should include:

    schema
    joining_device_label
    welcome_artifact_sha256
    group_id_ref
    joined_locally
    provider_storage_written
    private_material_included
    warning

## 9. Stdout behavior

Stdout must remain sanitized.

Allowed stdout fields:

- command;
- ok;
- provider;
- implementation;
- mode;
- phase;
- device label;
- conversation label;
- path hints;
- hashes;
- artifact sizes;
- provider events;
- warnings;
- private_material_included=false.

Disallowed stdout content:

- signer.json contents;
- private keys;
- raw secret material;
- OpenMLS provider storage;
- raw group state;
- raw Welcome bytes;
- raw KeyPackage bytes;
- recovery material.

Recommendation:

- Never print serialized artifact bytes to stdout in this rung.
- Use path/hash/size only.

## 10. Provider storage question

Conversation lifecycle likely requires persistent OpenMLS state earlier than previous rungs.

Known current state:

- identity creation writes signer.json but not OpenMLS provider storage.
- public-bundle artifact export reports provider_storage_written=false.

Open question:

- Does conversation-create/add-member/join require writing provider storage for correctness across commands/processes?

Recommendation:

- Inspect the prior Phase 2C scratch docs and code before implementation.
- If provider storage is required, make it explicit and dev-only.
- Add fields such as provider_storage_written=true only if the command actually writes such state.
- Do not call provider storage secure vault storage.
- Do not design production vault storage in this rung.

## 11. Events

Potential new provider events:

    provider.conversation.created
    provider.conversation.exists
    provider.conversation.missing
    provider.conversation.member_added
    provider.conversation.joined
    provider.welcome.exported
    provider.welcome.missing
    provider.artifact.invalid

Initial expected classifications:

### provider.conversation.created

- class: conversation/setup
- severity: info
- trust relevant: false initially
- action: append history / debug only

### provider.conversation.exists

- class: conversation/setup
- severity: warning
- trust relevant: false initially
- action: stop operation / append history

### provider.conversation.missing

- class: conversation/setup
- severity: warning
- trust relevant: false initially
- action: stop operation / show recovery path / append history

### provider.conversation.member_added

- class: public/setup or membership/setup
- severity: info
- trust relevant: probably true later, but dev-sidecar-only for now
- action: append history; do not mutate trust storage yet

### provider.conversation.joined

- class: public/setup or membership/setup
- severity: info
- trust relevant: probably true later, but dev-sidecar-only for now
- action: append history; do not mutate trust storage yet

### provider.welcome.exported

- class: public/setup
- severity: info
- trust relevant: false or pending
- action: append history / debug only

### provider.artifact.invalid

- class: input/artifact
- severity: warning or error
- trust relevant: false by default, unless linked to tamper path later
- action: stop operation / append history

Important future concern:

Membership changes are trust-sensitive at the user/product level. Even if this dev-sidecar rung keeps them non-mutating and not user-visible, future trust-state planning must treat group membership changes carefully.

## 12. Duplicate and overwrite behavior

Default behavior should refuse silent overwrite.

Recommended refusal cases:

- conversation-create refuses if conversation-summary.json or group state already exists.
- conversation-add-member refuses if welcome.bin or welcome-manifest.json already exists for the same operation path.
- conversation-join refuses if local joined state already exists for the same conversation/device path.

Do not add --force in this rung unless absolutely necessary.

Reasoning:

- Silent conversation/group replacement is a future trust hazard.
- Silent Welcome replacement could hide membership or onboarding changes.
- CarbonStack doctrine prefers loud trust changes and explicit state transitions.

## 13. Tests required before implementation is complete

Go-side contract tests should cover:

### conversation-create

- missing device label fails;
- missing conversation label fails;
- invalid device label fails;
- invalid conversation label fails;
- missing identity fails;
- successful create writes sanitized conversation summary;
- duplicate create refuses overwrite;
- stdout contains no obvious secret material;
- no Comms/Cypher/trust runtime mutation.

### conversation-add-member

- missing inviter identity fails;
- missing conversation fails;
- missing member KeyPackage artifact fails;
- invalid member KeyPackage artifact fails;
- successful add-member writes Welcome artifact;
- successful add-member writes Welcome manifest;
- stdout includes path/hash/size only;
- duplicate Welcome export refuses overwrite;
- stdout contains no obvious secret material.

### conversation-join

- missing joining identity fails;
- missing Welcome artifact fails;
- invalid Welcome artifact fails;
- successful join writes sanitized joined summary;
- duplicate join refuses overwrite;
- stdout contains no obvious secret material.

### Artifact/stability tests

- Welcome artifact exists and size > 0.
- Welcome manifest hash matches artifact bytes.
- KeyPackage artifact is consumed from file rather than stdout.
- Existing public-bundle artifact export tests remain valid.
- Existing summary-only public-bundle tests remain valid.

## 14. Manual probe shape after future implementation

Potential future manual probe:

    Remove-Item -Recurse -Force .\.carbonstack-openmls-sidecar-state -ErrorAction SilentlyContinue

    cargo run -- identity-create --device-label carbonstack-alice-device
    cargo run -- identity-create --device-label carbonstack-bob-device

    cargo run -- public-bundle-export --device-label carbonstack-bob-device --write-artifact

    cargo run -- conversation-create --device-label carbonstack-alice-device --conversation-label carbonstack-test-conversation

    cargo run -- conversation-add-member --device-label carbonstack-alice-device --conversation-label carbonstack-test-conversation --member-keypackage .carbonstack-openmls-sidecar-state\dev\devices\carbonstack-bob-device\public-bundle.keypackage.bin

    cargo run -- conversation-join --device-label carbonstack-bob-device --welcome .carbonstack-openmls-sidecar-state\dev\conversations\carbonstack-test-conversation\welcome.bin

Safe inspection only:

    Get-ChildItem .\.carbonstack-openmls-sidecar-state\dev\devices\carbonstack-alice-device | Select-Object Name, Length

    Get-ChildItem .\.carbonstack-openmls-sidecar-state\dev\devices\carbonstack-bob-device | Select-Object Name, Length

    Get-ChildItem .\.carbonstack-openmls-sidecar-state\dev\conversations\carbonstack-test-conversation | Select-Object Name, Length

Do not inspect signer.json, provider storage, or raw group state.

## 15. Implementation order recommendation

Recommended future implementation order:

1. Add conversation label validation.
2. Add state path helpers for conversations.
3. Add provider events/taxonomy docs for conversation lifecycle.
4. Add not-implemented command recognition for conversation-create/add-member/join if useful.
5. Implement conversation-create only.
6. Test conversation-create only.
7. Result doc for conversation-create.
8. Then plan add-member/Welcome export.
9. Then implement add-member/Welcome export.
10. Then plan join.
11. Then implement join.

Alternative:

- Implement all three commands in one larger rung only if OpenMLS API coupling makes splitting impossible.

Recommendation:

- Prefer smaller rungs unless API reality forces a combined lifecycle checkpoint.

## 16. Success criteria for this planning checkpoint

This planning checkpoint succeeds when:

- this doc is committed;
- no code changes are made;
- no runtime integration is started;
- next implementation target is clear;
- questions/blockers are explicitly listed.

## 17. Open questions before implementation

- What exact OpenMLS APIs should be used for group creation in the pinned dependency set?
- What exact OpenMLS APIs should consume a serialized KeyPackage artifact?
- What exact OpenMLS APIs should serialize and deserialize Welcome artifacts?
- What state must persist between conversation-create and conversation-add-member?
- What state must persist between conversation-join and later message protect/open?
- Is provider storage required now?
- Can group state be safely represented as dev-local files without designing production storage?
- Should conversation label be global or nested per creator device?
- Should Welcome artifacts live under conversation path or recipient device path?
- Should `provider.conversation.member_added` be marked trust-relevant now or deferred until trust-state storage planning?
- Should the next implementation checkpoint be conversation-create only, or full Alice-create/Bob-join lifecycle?

## 18. Recommended next result doc

After a future implementation checkpoint, write one of:

    docs/58-openmls-sidecar-conversation-create-result-v0.md

or, if the implementation necessarily covers create/add/join together:

    docs/58-openmls-sidecar-conversation-lifecycle-result-v0.md

Do not write a result doc until code/tests validate the actual behavior.

## 19. Next after planning

Recommended next technical checkpoint:

    Phase 2D conversation-create implementation

Only after conversation-create is validated should CarbonStack move to:

    conversation-add-member / Welcome export planning

Do not jump directly to message protect/open.
Do not jump directly to Comms runtime integration.
Do not jump directly to Cypher routing.
Do not jump directly to trust-state mutation.
