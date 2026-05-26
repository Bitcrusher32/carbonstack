# OpenMLS Sidecar Command Schema Matrix v0

Status: Command/schema matrix
Component: CarbonStackComms / OpenMLS sidecar
Phase: Post-Phase 2D closure, pre-Cypher relay recon
Previous docs:
- docs/86-openmls-sidecar-phase2d-mainline-closure-result-v0.md
- docs/91-openmls-sidecar-artifact-ownership-map-v0.md

## 1. Purpose

This document records the current OpenMLS sidecar command surface after Phase 2D mainline closure.

It is intended to support:

- sidecar promotion;
- module split;
- test-suite split;
- README cleanup;
- future Cypher relay design.

## 2. Matrix

| Command | Required args | Reads | Writes | Success event | Failure/event examples | Validation |
|---|---|---|---|---|---|---|
| provider-info | none | none | stdout JSON | provider.info.reported | unsupported command for unknown names | provider-info contract tests |
| identity-create | --device-label | none | signer.json, identity metadata | provider.identity.created | provider.identity.exists, provider.identity.invalid | identity tests |
| identity-status | --device-label | signer.json, identity metadata | stdout JSON | provider.identity.loaded | provider.identity.missing | identity tests |
| public-bundle-export | --device-label, optional --write-artifact | signer/provider identity state | summary; optional KeyPackage artifact, manifest, device-root provider-storage | provider.public_bundle.exported | provider.identity.missing, artifact exists | public bundle tests |
| conversation-create | --device-label, --conversation-label | device identity/signer state | device-scoped conversation-summary, provider-storage | provider.conversation.created | provider.conversation.exists | conversation tests |
| conversation-load-check | --device-label, --conversation-label | device-scoped conversation provider-storage | stdout JSON | provider.conversation.loaded | provider.conversation.missing | conversation tests |
| conversation-add-member | --device-label, --conversation-label, --member-keypackage | creator conversation provider-storage, member KeyPackage artifact | welcome.bin, welcome-manifest, add-member-summary, updated provider-storage | provider.conversation.member_added | provider.member_keypackage.invalid | conversation tests |
| conversation-join | --device-label, --conversation-label, --welcome | device-root provider-storage, Welcome artifact | device-scoped joined conversation provider-storage, summary, join-summary | provider.conversation.joined | provider.welcome.invalid, provider.conversation.missing | join tests |
| message-protect | --device-label, --conversation-label, --plaintext, optional --message-label | sender device-scoped conversation provider-storage | application-message.bin, message-manifest, message-protect-summary, updated provider-storage | provider.message.protected | provider.message.exists, provider.conversation.missing | message tests |
| message-open | --device-label, --conversation-label, --message, optional --message-label | receiver device-scoped conversation provider-storage, application-message.bin | message-open-summary, updated provider-storage, stdout plaintext proof | provider.message.opened | provider.message.invalid, provider.conversation.missing, checkpoint.failed | message and negative tests |

## 3. Current supported command list

Supported:

- provider-info;
- identity-create;
- identity-status;
- public-bundle-export;
- conversation-create;
- conversation-load-check;
- conversation-add-member;
- conversation-join;
- message-protect;
- message-open.

Unsupported:

- state-checkpoint;
- state-load-check.

## 4. Current notable failure shapes

Wrong-device message-open:

    code=conversation_or_message_missing
    provider_event=provider.conversation.missing
    severity=warning
    trust_relevant=false
    exit code=3

Wrong-conversation message-open:

    code=conversation_or_message_missing
    provider_event=provider.conversation.missing
    severity=warning
    trust_relevant=false
    exit code=3

Duplicate/replay message-open:

    code=message_open_failed
    provider_event=checkpoint.failed
    severity=warning
    trust_relevant=false
    exit code=3

Corrupt/truncated message artifact:

    code=message_artifact_invalid
    provider_event=provider.message.invalid
    severity=warning
    trust_relevant=false
    exit code=3

## 5. Trust relevance

Current sidecar events remain dev-provider events and do not mutate trust-state storage.

Most current events are not trust-relevant.

Future trust-state integration must be explicitly designed and tested. Do not infer trust mutation from current event names.

## 6. Cypher relevance

Cypher relay research should focus on the artifact-producing and artifact-consuming commands:

- public-bundle-export;
- conversation-add-member;
- conversation-join;
- message-protect;
- message-open.

Primary relay artifacts:

- public-bundle.keypackage.bin;
- welcome.bin;
- application-message.bin.

Local-only sensitive state:

- signer.json;
- provider-storage.json.

## 7. Non-goals

This matrix does not define:

- final production API;
- final wire format;
- final mobile UX;
- final server schema;
- secure vault implementation;
- generated message IDs;
- trust-state mutation;
- Cypher relay routes.

It is a current Phase 2D sidecar contract map.
