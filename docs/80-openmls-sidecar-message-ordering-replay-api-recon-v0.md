# OpenMLS Sidecar Message Ordering / Replay API Recon v0

Status: API recon / manual probe result
Component: CarbonStackComms / OpenMLS sidecar
Phase: Phase 2D
Previous docs:
- docs/75-openmls-sidecar-message-protect-open-result-v0.md
- docs/76-openmls-sidecar-multi-message-continuity-plan-v0.md
- docs/77-openmls-sidecar-multi-message-api-recon-v0.md
- docs/78-openmls-sidecar-multi-message-continuity-result-v0.md
- docs/79-openmls-sidecar-message-ordering-replay-plan-v0.md

## 1. Purpose

This document records v0.2.39 ordering/replay reconnaissance for the OpenMLS sidecar.

v0.2.38 validated the happy path:

    create -> add-member -> Welcome export -> join -> protect/open message-0001 -> protect/open message-0002

The next question was:

    What happens when Bob receives messages out of order, duplicated, replayed, or corrupted?

This document records source recon and manual probes. It does not implement new behavior.

## 2. Current validated baseline

The current sidecar supports:

    provider-info
    identity-create
    identity-status
    public-bundle-export
    conversation-create
    conversation-load-check
    conversation-add-member
    conversation-join
    message-protect
    message-open

The v0.2.38 validated ordered path was:

    message-0001 -> "hello bob 1"
    message-0002 -> "hello bob 2"

Both messages are protected by Alice and opened by Bob across separate sidecar invocations.

Both protect and open save provider storage and prove group reloadability.

## 3. Source recon summary

Targeted OpenMLS 0.8.1 source recon focused on:

    src/group/mls_group/processing.rs

Important observed path:

    MlsGroup::process_message(...)
        -> unprotect_message(...)
        -> decrypt_message(...)
        -> public_group.parse_message(...)
        -> process_unverified_message(...)

Relevant observed behavior:

    will_modify_secret_tree = matches!(message, ProtocolMessage::PrivateMessage(_))

When a private message is processed, OpenMLS writes message secrets back to provider storage:

    provider.storage().write_message_secrets(self.group_id(), &self.message_secrets_store)

Implication:

    message-open remains persistence-relevant.
    save-after-open must stay.
    duplicate/replay behavior can be enforced by the secret tree / sender ratchet state.

This recon reinforces the existing v0.2.36/v0.2.38 discipline:

    save provider storage after successful message-protect
    save provider storage after successful message-open

## 4. Sidecar failure mapping recon

Current sidecar message-open failure mapping is broad.

Relevant mappings:

- missing conversation/message file:
  - code: conversation_or_message_missing
  - provider_event: provider.conversation.missing
  - severity: warning
  - trust_relevant: false
  - exit code: 3

- invalid message path:
  - code: invalid_message_path
  - provider_event: provider.command.invalid
  - severity: warning
  - trust_relevant: false
  - exit code: 2

- invalid artifact/deserialization:
  - code: message_artifact_invalid
  - provider_event: provider.message.invalid
  - severity: warning
  - trust_relevant: false
  - exit code: 3

- generic process/decrypt failure:
  - code: message_open_failed
  - provider_event: checkpoint.failed
  - severity: warning
  - trust_relevant: false
  - exit code: 3

The duplicate-open probe currently lands in the generic process/decrypt failure bucket:

    message_open_failed
    checkpoint.failed

This is behaviorally correct enough for dev-sidecar, but future CarbonStack event taxonomy should likely split replay/duplicate errors into a more specific provider event.

## 5. Probe B: protect both, open in order

### Flow

Fresh state:

1. Alice identity-create.
2. Bob identity-create.
3. Bob public-bundle-export --write-artifact.
4. Alice conversation-create.
5. Alice conversation-add-member.
6. Bob conversation-join.
7. Alice protects message-0001 with plaintext "hello bob 1".
8. Alice protects message-0002 with plaintext "hello bob 2".
9. Bob opens message-0001.
10. Bob opens message-0002.

### Result

Both opens succeeded.

message-0001 open returned:

    ok=true
    command=message-open
    message_label=message-0001
    plaintext_utf8="hello bob 1"
    plaintext_len=11
    message_opened=true
    provider_storage_loaded=true
    provider_storage_written=true
    group_reloadable=true

message-0002 open returned:

    ok=true
    command=message-open
    message_label=message-0002
    plaintext_utf8="hello bob 2"
    plaintext_len=11
    message_opened=true
    provider_storage_loaded=true
    provider_storage_written=true
    group_reloadable=true

### Conclusion

Bob does not need to open message-0001 before Alice can protect message-0002.

Ordered batch delivery works:

    protect message-0001
    protect message-0002
    open message-0001
    open message-0002

## 6. Probe C: protect both, open out of order

### Flow

Fresh state:

1. Alice identity-create.
2. Bob identity-create.
3. Bob public-bundle-export --write-artifact.
4. Alice conversation-create.
5. Alice conversation-add-member.
6. Bob conversation-join.
7. Alice protects message-0001 with plaintext "hello bob 1".
8. Alice protects message-0002 with plaintext "hello bob 2".
9. Bob opens message-0002 first.
10. Bob opens message-0001 afterward.

### Result

Both opens succeeded.

message-0002 opened first and returned:

    ok=true
    message_label=message-0002
    plaintext_utf8="hello bob 2"
    plaintext_len=11
    message_opened=true
    provider_storage_loaded=true
    provider_storage_written=true
    group_reloadable=true

message-0001 opened afterward and returned:

    ok=true
    message_label=message-0001
    plaintext_utf8="hello bob 1"
    plaintext_len=11
    message_opened=true
    provider_storage_loaded=true
    provider_storage_written=true
    group_reloadable=true

### Conclusion

For this exact OpenMLS 0.8.1 dev-sidecar Alice-to-Bob private application-message flow, out-of-order open succeeded.

This suggests OpenMLS can handle the skipped message secret case for this small sequence:

    protect message-0001
    protect message-0002
    open message-0002
    open message-0001

Do not overgeneralize this to all epochs, all sender ratchet configurations, all group sizes, or arbitrary long gaps. This is validated only for the current two-message dev-local proof.

## 7. Probe D: duplicate-open / replay

### Flow

Fresh state:

1. Alice identity-create.
2. Bob identity-create.
3. Bob public-bundle-export --write-artifact.
4. Alice conversation-create.
5. Alice conversation-add-member.
6. Bob conversation-join.
7. Alice protects message-0001 with plaintext "hello bob 1".
8. Bob opens message-0001.
9. Bob tries to open the same message-0001 artifact again under label message-0001-duplicate.

### Result

First open succeeded:

    ok=true
    message_label=message-0001
    plaintext_utf8="hello bob 1"
    message_opened=true

Duplicate open failed.

Failure envelope:

    ok=false
    command=message-open
    code=message_open_failed
    message="message open failed: ValidationError(UnableToDecrypt(SecretTreeError(SecretReuseError)))"
    provider_event=checkpoint.failed
    severity=warning
    trust_relevant=false
    private_material_included=false

Process exit:

    exit code 3

### Conclusion

Duplicate-open/replay of the same private application message is rejected in the current sidecar flow.

The underlying OpenMLS error was:

    SecretTreeError(SecretReuseError)

This is strong evidence that duplicate open consumes or invalidates the relevant message secret after first use.

CarbonStack should eventually classify this more specifically than generic checkpoint failure.

Recommended future event:

    provider.message.replay_detected

or:

    provider.message.secret_reuse_detected

Potential severity:

    warning or block

Potential trust relevance:

    likely true once integrated into product trust UX, because replay/duplicate delivery can indicate either benign retry or adversarial replay.

Do not make it trust-relevant in the sidecar yet without a broader event taxonomy decision.

## 8. Probe F: corrupted artifact

### Initial failed attempt

The first corruption probe was inconclusive because PowerShell failed to resolve the good artifact path before byte mutation. No corrupt artifact was created, so the sidecar only observed a missing file.

That initial failure produced:

    code=conversation_or_message_missing
    provider_event=provider.conversation.missing

This was missing-file behavior, not corrupted-artifact behavior.

### Corrected flow

The corrected probe used resolved absolute paths:

    $Root = (Get-Location).Path
    $Good = Join-Path $Root ".carbonstack-openmls-sidecar-state\dev\conversations\carbonstack-test-conversation\messages\message-0001\application-message.bin"
    $Bad = Join-Path $Root ".carbonstack-openmls-sidecar-state\dev\conversations\carbonstack-test-conversation\messages\message-0001\corrupt-application-message.bin"

The probe verified:

    Test-Path $Good = True
    Test-Path $Bad = True

Then it truncated the valid artifact without printing raw bytes and attempted:

    message-open --message-label corrupt-message-0001 --message <absolute corrupt artifact path>

### Corrected result

The corrupt artifact was rejected before message processing.

Failure envelope:

    ok=false
    command=message-open
    code=message_artifact_invalid
    message="message artifact deserialization failed: EndOfStream"
    provider_event=provider.message.invalid
    severity=warning
    trust_relevant=false
    private_material_included=false
    exit code 3

### Conclusion

Truncated/corrupt `application-message.bin` is rejected as invalid provider message material.

For this exact truncation case, the sidecar maps the failure to:

    message_artifact_invalid
    provider.message.invalid

This is the desired broad behavior for corrupted serialized MLS artifacts.

Future tests can assert this path without inspecting raw artifact bytes.

## 10. Ordering/replay conclusions

Validated:

- Ordered batch delivery works:
  - protect message-0001
  - protect message-0002
  - open message-0001
  - open message-0002

- Out-of-order two-message delivery works in this exact current flow:
  - protect message-0001
  - protect message-0002
  - open message-0002
  - open message-0001

- Duplicate/replay of the same already-opened message fails:
  - first open succeeds
  - second open fails with SecretReuseError

Validated:

- corrupted/truncated existing artifact behavior:
  - rejected with message_artifact_invalid;
  - provider_event=provider.message.invalid;
  - underlying deserialization failure was EndOfStream.

Not validated:

- wrong conversation label;
- wrong device label;
- long skipped-message windows;
- behavior across epoch changes;
- behavior with more than two messages;
- behavior with multiple senders;
- behavior after membership changes;
- behavior after failed out-of-order opens, because current out-of-order probe succeeded;
- whether failed duplicate-open mutates provider storage before error;
- whether sidecar should write failure summaries.

## 11. CarbonStack implications

### Cypher routing

Cypher can likely deliver opaque MLS application-message artifacts without strict ordering for the simplest two-message same-sender case, because Bob successfully opened message-0002 before message-0001.

However, CarbonStack should still preserve ordering metadata.

Cypher should carry or store:

- conversation label / future conversation ID;
- sender device;
- recipient scope;
- message label / future message ID;
- artifact hash;
- artifact size;
- created timestamp;
- delivery timestamp;
- retry count or duplicate marker;
- optional sequence hint.

The server must not need:

- plaintext;
- private provider storage;
- signer material;
- sender ratchet state;
- receiver ratchet state;
- group secrets.

### Comms runtime

Comms should treat duplicate-open / replay errors as meaningful protocol outcomes, not just generic failures.

Suggested future classification:

- duplicate/replayed message:
  - event: provider.message.replay_detected or provider.message.secret_reuse_detected
  - default severity: warning/block
  - trust relevance: likely true after taxonomy review

- missing artifact:
  - event: provider.message.missing or provider.artifact.missing
  - severity: warning
  - trust relevance: probably false by default

- corrupted artifact:
  - not validated yet
  - likely event: provider.message.invalid or provider.message.authentication_failed

### Trust state

Do not mutate trust state yet.

But later, replay/secret-reuse should likely feed a user-visible warning or delivery integrity log.

## 12. Recommended next checkpoint

The corrected corrupt-artifact probe is now validated.

Next decide whether v0.2.40 should implement tests for:

    TestOpenMLSSidecarMessageOpenOutOfOrderTwoMessageDelivery
    TestOpenMLSSidecarMessageOpenDuplicateRejected
    TestOpenMLSSidecarMessageOpenCorruptArtifactRejected

If corrupted-artifact behavior remains messy, split it:

    v0.2.40: ordering + duplicate tests
    v0.2.41: corrupted artifact / invalid artifact classification cleanup

Do not route through Cypher yet.

Do not migrate Alice state yet.

Do not wire Comms runtime yet.


