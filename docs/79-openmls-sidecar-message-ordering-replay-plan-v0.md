# OpenMLS Sidecar Message Ordering / Replay Plan v0

Status: Plan / pre-recon
Component: CarbonStackComms / OpenMLS sidecar
Phase: Phase 2D
Previous docs:
- docs/75-openmls-sidecar-message-protect-open-result-v0.md
- docs/76-openmls-sidecar-multi-message-continuity-plan-v0.md
- docs/77-openmls-sidecar-multi-message-api-recon-v0.md
- docs/78-openmls-sidecar-multi-message-continuity-result-v0.md

## 1. Purpose

v0.2.38 validated two sequential Alice-to-Bob application messages across separate sidecar invocations:

    create -> add-member -> Welcome export -> join -> protect/open message-0001 -> protect/open message-0002

The next question is not whether ordered delivery works. That is validated.

The next question is:

    What happens when delivery is out of order, duplicated, replayed, or stale?

This matters before CarbonStackCypher routing because an eventually-connected relay can naturally deliver:

- duplicate artifacts;
- artifacts retried by clients;
- artifacts fetched out of order;
- stale artifacts;
- missing earlier artifacts;
- maliciously replayed artifacts.

This document plans recon only. It does not implement new behavior.

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

The current validated ordered flow is:

1. Alice identity-create.
2. Bob identity-create.
3. Bob public-bundle-export --write-artifact.
4. Alice conversation-create.
5. Alice conversation-add-member using Bob KeyPackage.
6. Bob conversation-join using Alice Welcome.
7. Alice message-protect --message-label message-0001 --plaintext "hello bob 1".
8. Bob message-open --message-label message-0001.
9. Bob recovers "hello bob 1".
10. Alice message-protect --message-label message-0002 --plaintext "hello bob 2".
11. Bob message-open --message-label message-0002.
12. Bob recovers "hello bob 2".

Both protect and open save provider storage and prove group reloadability.

## 3. Why ordering/replay recon is next

Before routing MLS application payloads through Cypher, CarbonStack needs to know what the sidecar does when artifacts arrive in non-happy-path patterns.

Important cases:

1. Bob opens message-0002 before message-0001.
2. Bob opens message-0001 twice.
3. Bob opens message-0002 twice.
4. Bob opens message-0001 after already opening message-0002.
5. Bob opens a corrupted application-message artifact.
6. Bob opens an application message with the wrong conversation label.
7. Bob opens an application message with the wrong device label.
8. Alice protects message-0002 before Bob has opened message-0001.
9. Alice protects message-0001 and message-0002 without Bob opening either, then Bob opens in order.
10. Alice protects message-0001 and message-0002 without Bob opening either, then Bob opens out of order.

Some of these may work and some may fail. The goal is not to force behavior yet. The goal is to record exact OpenMLS/sidecar behavior.

## 4. Desired recon questions

The recon must answer:

1. Does OpenMLS allow Bob to process application messages out of order?
2. If out-of-order succeeds, does OpenMLS internally handle skipped message secrets?
3. If out-of-order fails, what exact error appears?
4. Does a failed out-of-order open mutate Bob provider storage?
5. Does opening message-0001 after a failed message-0002 attempt still work?
6. Does opening message-0001 twice fail?
7. If duplicate-open fails, what exact error appears?
8. Does duplicate-open mutate Bob provider storage?
9. Does opening message-0002 after message-0001 then opening message-0001 again fail deterministically?
10. Does corrupted application-message.bin fail before provider storage mutation?
11. Are errors distinguishable enough for CarbonStack event taxonomy?
12. Should CarbonStack classify these as warning, block, corruption, or replay?
13. Does sidecar currently save provider storage after failed open? It should not unless explicitly necessary and understood.
14. Should message-open write failure summaries, or only success summaries?
15. Should Cypher preserve ordering metadata separately from MLS artifacts?

## 5. Proposed manual probes

### Probe A: ordered baseline

Expected success:

    protect/open message-0001
    protect/open message-0002

This is already validated by v0.2.38 and should remain the baseline.

### Probe B: protect both, open in order

Flow:

    protect message-0001
    protect message-0002
    open message-0001
    open message-0002

Question:

    Does Bob need to process message-0001 before Alice creates message-0002?

Expected guess:

    This likely succeeds, but must be validated.

### Probe C: protect both, open out of order

Flow:

    protect message-0001
    protect message-0002
    open message-0002
    open message-0001

Question:

    Does OpenMLS support skipped-message secrets/out-of-order private message processing?

Expected guess:

    Unknown. Do not assume.

### Probe D: duplicate-open

Flow:

    protect message-0001
    open message-0001
    open message-0001 again

Question:

    Does OpenMLS reject replay/duplicate processing?

Expected guess:

    Likely fails, but exact error matters.

### Probe E: wrong conversation label

Flow:

    protect message-0001 in conversation A
    Bob attempts open with conversation label B

Question:

    Does sidecar fail at provider group load / missing conversation state before OpenMLS processing?

Expected guess:

    Should fail as missing device conversation provider storage.

### Probe F: corrupted artifact

Flow:

    protect message-0001
    copy application-message.bin to a temp file
    modify one byte or truncate
    open corrupted file

Question:

    Does deserialization fail, decrypt fail, or process fail?

Expected guess:

    Unknown. Must classify exact behavior.

## 6. State mutation safety during recon

During probes, never inspect or paste:

- signer.json;
- provider-storage.json;
- raw MemoryStorage;
- raw group state;
- raw KeyPackage bytes;
- raw Welcome bytes;
- raw application-message bytes.

Allowed safe inspection:

- stdout envelopes;
- sanitized summary JSON if needed;
- file existence;
- file sizes;
- path hints;
- errors;
- hashes.

Do not cat application-message.bin.

For corruption tests, use PowerShell byte manipulation without printing raw bytes.

## 7. Expected implementation boundary after recon

The next implementation after recon may add Go tests for deterministic behavior.

Possible tests:

- `TestOpenMLSSidecarMessageOpenDuplicateRejected`
- `TestOpenMLSSidecarMessageOpenOutOfOrderBehavior`
- `TestOpenMLSSidecarMessageOpenCorruptArtifactRejected`

But do not add tests until the behavior is observed and documented.

## 8. CarbonStack implications

Cypher likely needs delivery metadata beyond opaque MLS bytes:

- conversation ID;
- sender device;
- receiver device or recipient group scope;
- message label or future message ID;
- sequence / ordering hint;
- artifact hash;
- artifact size;
- created timestamp;
- delivery status;
- retry status.

But Cypher must not need plaintext or private MLS state.

The server should remain opaque to:

- plaintext;
- provider storage;
- sender ratchet secrets;
- receiver ratchet state;
- private keys;
- group secrets.

## 9. Known current constraints

Current sidecar constraints:

- Alice state remains global:
  - `dev/conversations/<conversation-label>/`
- Bob state remains device-scoped:
  - `dev/devices/<device-label>/conversations/<conversation-label>/`
- message labels are explicit safe filesystem labels;
- message-0001 remains the default;
- no autogenerated message IDs exist;
- no Cypher routing exists;
- no Comms runtime integration exists;
- no trust-state mutation exists.

## 10. Result doc requirements

The recon result should record:

- exact OpenMLS source paths inspected;
- exact manual probe commands;
- exact stdout success/failure envelopes;
- whether failed opens mutate provider storage;
- whether duplicate opens fail;
- whether out-of-order opens succeed or fail;
- whether later in-order recovery works after out-of-order failure;
- CarbonStack event classification recommendation;
- whether next step should be tests, state cleanup, or Cypher design.

## 11. Recommended next document

Next document:

    docs/80-openmls-sidecar-message-ordering-replay-api-recon-v0.md

It should record both source recon and manual probe observations.
