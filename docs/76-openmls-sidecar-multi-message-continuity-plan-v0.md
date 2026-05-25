# OpenMLS Sidecar Multi-Message Continuity Plan v0

Status: Plan / pre-recon
Component: CarbonStackComms / OpenMLS sidecar
Phase: Phase 2D
Previous docs:
- docs/72-openmls-sidecar-conversation-join-result-v0.md
- docs/73-openmls-sidecar-message-protect-open-skeleton-v0.md
- docs/74-openmls-sidecar-message-protect-open-api-recon-v0.md
- docs/75-openmls-sidecar-message-protect-open-result-v0.md

## 1. Purpose

v0.2.36 proved the first dev-local OpenMLS application message lifecycle:

    create -> add-member -> Welcome export -> join -> protect -> open

The validated one-message proof was:

    Alice protects "hello bob"
    Bob opens the protected artifact
    Bob recovers "hello bob"

This document plans the next continuity checkpoint before implementation.

The next question is not whether one protected message can be opened. That is validated.

The next question is:

    Can the sidecar preserve MLS message state across multiple protected/opened messages and separate sidecar invocations?

This document is intentionally planning-only. It does not implement multi-message continuity.

## 2. Current validated baseline

The sidecar currently supports:

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

The current validated message proof uses:

    message-protect --device-label carbonstack-alice-device --conversation-label carbonstack-test-conversation --plaintext "hello bob"

and writes:

    .carbonstack-openmls-sidecar-state/dev/conversations/carbonstack-test-conversation/messages/message-0001/application-message.bin

Then Bob runs:

    message-open --device-label carbonstack-bob-device --conversation-label carbonstack-test-conversation --message <application-message.bin>

and recovers:

    plaintext_utf8="hello bob"

## 3. Known limitation

The v0.2.36 implementation hardcodes:

    message-0001

This is acceptable for the first proof, but it blocks multi-message testing.

Before implementing multi-message continuity, the sidecar needs a controlled way to select a message artifact label.

## 4. Proposed command extension

Add explicit message label support.

Proposed command:

    message-protect --device-label <sender-device> --conversation-label <conversation> --message-label <safe-message-label> --plaintext <text>

Current command remains valid as a compatibility/default path:

    message-protect --device-label <sender-device> --conversation-label <conversation> --plaintext <text>

Default behavior may remain:

    message_label = message-0001

Proposed message-open command may optionally accept `--message-label` for summary path selection, but it does not strictly need it because the message artifact path is already explicit.

Potential command:

    message-open --device-label <receiver-device> --conversation-label <conversation> --message <path> --message-label <safe-message-label>

If omitted, message-open may infer or default to:

    message-0001

Recommendation for implementation:

- Add `--message-label` to message-protect first.
- Add `--message-label` to message-open only if summary path correctness requires it.
- Keep `--message <path>` as the authoritative input for message-open.

## 5. Message label validation

Message labels should be safe filesystem labels.

Initial validation should allow:

    a-z
    A-Z
    0-9
    -
    _

Initial validation should reject:

    empty labels
    labels longer than 64 bytes
    path separators
    dot segments
    spaces
    shell metacharacters
    Unicode normalization surprises
    labels beginning with `.`
    Windows reserved names if practical

Suggested examples:

Allowed:

    message-0001
    message-0002
    alice-0001
    bob-0001
    test-msg-1

Rejected:

    ../message-0001
    message/0001
    message 0001
    .secret
    provider-storage
    signer
    welcome
    con
    nul

## 6. Artifact layout target

With explicit message labels, Alice protected artifacts should write to:

    .carbonstack-openmls-sidecar-state/dev/conversations/<conversation-label>/messages/<message-label>/application-message.bin

Metadata:

    message-manifest.json
    message-protect-summary.json

Bob open summaries should write to:

    .carbonstack-openmls-sidecar-state/dev/devices/<device-label>/conversations/<conversation-label>/opened-messages/<message-label>/message-open-summary.json

This preserves the existing v0.2.36 Alice-global / Bob-device-scoped asymmetry for one more proof rung.

## 7. Multi-message proof target

The basic happy path should be:

1. reset sidecar state;
2. Alice identity-create;
3. Bob identity-create;
4. Bob public-bundle-export --write-artifact;
5. Alice conversation-create;
6. Alice conversation-add-member with Bob KeyPackage;
7. Bob conversation-join with welcome.bin;
8. Alice message-protect --message-label message-0001 --plaintext "hello bob 1";
9. Bob message-open --message-label message-0001 --message <message-0001 artifact>;
10. Bob recovers "hello bob 1";
11. Alice message-protect --message-label message-0002 --plaintext "hello bob 2";
12. Bob message-open --message-label message-0002 --message <message-0002 artifact>;
13. Bob recovers "hello bob 2";
14. both sides report provider_storage_loaded=true;
15. both sides report provider_storage_written=true;
16. both sides report group_reloadable=true;
17. no forbidden private material appears in stdout.

## 8. Ordering / skipped-message recon questions

The recon must answer:

1. Does OpenMLS require Bob to process application messages in sender order?
2. What happens if Bob opens message-0002 before message-0001?
3. Does OpenMLS buffer skipped application messages?
4. Does OpenMLS reject out-of-order messages?
5. Does processing message-0001 advance receiver state such that message-0002 can then open?
6. Can Alice create message-0002 before Bob opens message-0001?
7. Does sender state advance after message-0001 protect in a way required for message-0002 protect?
8. Does message-protect need to persist after every message to support future messages?
9. Does message-open need to persist after every message to support future messages?
10. Are stale or replayed application messages rejected by process_message?

## 9. State layout question

The current layout is still asymmetric:

Alice/global state:

    dev/conversations/<conversation-label>/

Bob/device-scoped joined state:

    dev/devices/<device-label>/conversations/<conversation-label>/

For v0.2.37/v0.2.38, this plan recommends preserving the asymmetry temporarily.

Reason:

- The one-message proof works.
- Multi-message continuity is the next unknown.
- Migrating Alice state before testing continuity may mix two risks.

However, the recon should record whether multi-message continuity makes the asymmetry painful.

If so, a later cleanup checkpoint should migrate Alice to:

    dev/devices/<alice-device>/conversations/<conversation-label>/

Do not perform the migration during the first multi-message proof unless absolutely necessary.

## 10. Provider-info cleanup carry-forward

Before this plan, provider-info raw hardcoded JSON was cleaned up into structured JSON.

This reduces risk before adding more command fields.

Future command changes should update:

- CAPABILITIES;
- UNSUPPORTED_COMMANDS;
- provider-info Go expectations;
- unsupported-command test target if needed.

## 11. Security/stdout rules

Never print:

- signer.json;
- provider-storage.json;
- raw provider storage;
- raw group state;
- raw Welcome bytes;
- raw KeyPackage bytes;
- raw application-message bytes;
- private keys;
- MemoryStorage JSON.

Allowed to print:

- path hints;
- hashes;
- sizes;
- epochs;
- member counts;
- group refs;
- bounded dev plaintext from message-open only.

Plaintext stdout remains test-only and must not be described as final product UX.

## 12. Required implementation tests later

When implementation begins, add tests for:

- message-protect default message-label still works if kept;
- message-protect explicit message-0001 writes expected path;
- message-protect explicit message-0002 writes expected path;
- duplicate explicit message label refuses overwrite;
- invalid message labels fail;
- message-open message-0001 recovers expected plaintext;
- message-open message-0002 recovers expected plaintext;
- message-open invalid artifact fails;
- out-of-order open behavior is documented by a test if deterministic;
- replay/open duplicate behavior is documented by a test if deterministic;
- no forbidden private material in stdout.

## 13. Recommended next document

Next document:

    docs/77-openmls-sidecar-multi-message-api-recon-v0.md

It should record:

- exact OpenMLS source locations for application message ordering/ratchet behavior;
- whether out-of-order application messages are supported or rejected;
- expected behavior for replay/stale messages;
- whether sidecar implementation should test out-of-order behavior now or defer it;
- final recommendation for v0.2.38 implementation scope.
