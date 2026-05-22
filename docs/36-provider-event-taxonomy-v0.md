# CarbonStack Provider Event Taxonomy v0

## Status

Classification: PHASE 2C PROVIDER-CONTRACT TAXONOMY / PRE-INTEGRATION

This document classifies provider events before OpenMLS is wired into CarbonStackComms or CarbonStackCypher.

It is based on the OpenMLS fixture-contract work and the current provider-neutral boundary.

It does not implement provider integration.

It does not route MLS payloads through Cypher.

It does not claim production E2EE.

## Source Context

Relevant docs:

- `docs/30-openmls-provider-boundary-implications.md`
- `docs/33-openmls-memory-storage-persistence-result-v0.md`
- `docs/34-openmls-provider-fixture-contract-plan.md`
- `docs/35-openmls-provider-fixture-result-v0.md`

Relevant fixture path:

- `carbonstack-comms/internal/protocol/mls/research/openmls-minimal/fixtures/dev`

Relevant Go-side fixture parser:

- `carbonstack-comms/internal/protocol/openmls_fixture_test.go`

## Purpose

Provider events are the bridge between cryptographic provider behavior and CarbonStack's own application policy.

CarbonStack should not treat provider events as raw logs only.

Provider events may drive:

- trust history
- user warnings
- blocked sends
- blocked receives
- reverify prompts
- storage checkpoint decisions
- hostile-server detection candidates
- local recovery flows

## Taxonomy Overview

Provider events are grouped into seven classes:

1. lifecycle
2. public setup
3. membership
4. message
5. storage/checkpoint
6. trust/security
7. terminal/fatal provider errors

## 1. Lifecycle Events

Lifecycle events describe provider startup, fixture generation, initialization, and normal provider-mode transitions.

Current fixture examples:

- `provider.fixture.started`
- `provider.fixture.completed`

Future likely examples:

- `provider.initialized`
- `provider.loaded`
- `provider.closed`
- `provider.version.detected`
- `provider.capabilities.detected`
- `provider.mode.changed`

Expected CarbonStack behavior:

- Usually append to debug/provider history.
- Usually not user-visible.
- May be included in diagnostic output.
- Should not by itself change trust state.

Risks:

- Provider version/capability drift may matter later.
- Provider mode changes from dev/mock/fixture to real should be explicit and auditable.

## 2. Public Setup Events

Public setup events describe creation, publication, or consumption of public provider material.

Current fixture examples:

- `provider.public_bundle.created`

OpenMLS-shaped examples:

- KeyPackage created
- public setup bundle generated
- public bundle hash/reference created
- public setup material consumed

Future likely examples:

- `provider.public_bundle.created`
- `provider.public_bundle.exported`
- `provider.public_bundle.imported`
- `provider.public_bundle.published`
- `provider.public_bundle.consumed`
- `provider.public_bundle.replaced`

Expected CarbonStack behavior:

- Record in provider history.
- Expose summary in debug/inspection commands.
- Do not treat as verified identity by default.
- Feed trust-state only after verification policy exists.

Trust implications:

A public bundle is not the same thing as trusted identity.

A public setup event says material exists.

It does not say the user verified it.

## 3. Membership Events

Membership events describe conversation/group changes.

Current fixture examples:

- `conversation.created`
- `conversation.member_added`
- `conversation.welcome.created`
- `conversation.welcome.staged`
- `conversation.joined`

Future likely examples:

- `conversation.created`
- `conversation.loaded`
- `conversation.member_added`
- `conversation.member_removed`
- `conversation.member_updated`
- `conversation.commit.created`
- `conversation.commit.merged`
- `conversation.welcome.created`
- `conversation.welcome.received`
- `conversation.welcome.staged`
- `conversation.joined`
- `conversation.epoch.changed`

Expected CarbonStack behavior:

- Record in conversation/provider history.
- Surface membership changes where user trust expectations require it.
- Treat unexpected membership changes as suspicious until policy is defined.
- Never let the server silently become identity truth.

Trust implications:

Membership events are high-signal.

A member add/remove/update can affect who can read future messages.

For 1:1 conversations, group-shaped membership changes should still be treated as important because every CarbonStack conversation is conceptually group-shaped.

## 4. Message Events

Message events describe protect/open operations.

Current fixture examples:

- `message.protected`
- `message.opened`

Future likely examples:

- `message.protect.started`
- `message.protected`
- `message.open.started`
- `message.opened`
- `message.rejected`
- `message.duplicate.detected`
- `message.stale_epoch.detected`
- `message.malformed.detected`
- `message.unreadable`

Expected CarbonStack behavior:

- `message.protected` may require provider-state checkpoint after send.
- `message.opened` may require provider-state checkpoint after receive.
- `message.rejected` may need trust/security mapping depending on cause.
- Message events should not expose plaintext to Cypher.
- Message events should not leak secret material.

Persistence implications:

Prior OpenMLS probes showed:

- `create_message` mutates Alice/provider state.
- `process_message` mutates Bob/provider state.

Therefore:

- outbound protection is persistence-relevant
- inbound opening is persistence-relevant
- both may require checkpoint events

## 5. Storage / Checkpoint Events

Storage/checkpoint events describe provider state save/load requirements and restart continuity.

Current fixture-adjacent examples:

- `conversation.loaded`

Current scratch-probe concepts:

- Alice/Bob `MemoryStorage` saved
- Alice/Bob `MemoryStorage` loaded
- Alice signer saved
- Alice signer loaded
- Alice/Bob `MlsGroup` loaded with `MlsGroup::load`

Future likely examples:

- `storage.save.required`
- `storage.saved`
- `storage.load.started`
- `storage.loaded`
- `storage.missing`
- `storage.corrupt`
- `storage.version.unsupported`
- `checkpoint.required`
- `checkpoint.completed`
- `checkpoint.failed`
- `identity.signer.saved`
- `identity.signer.loaded`
- `identity.signer.missing`

Expected CarbonStack behavior:

- Treat checkpoint events as operationally important.
- Never ignore failed persistence after provider state mutation.
- Distinguish missing storage from invalid cryptographic material.
- Avoid production claims until storage is backed by a secure vault or equivalent design.

Security interpretation:

Storage events can be security-relevant.

Losing signer state or group state can cause identity failures, unreadable messages, or unsafe recovery flows.

## 6. Trust / Security Events

Trust/security events describe provider-observed conditions that should influence CarbonStack trust state or user warnings.

Current fixture result:

- `invalid-signature-error.json` maps the observed OpenMLS error `ValidationError(InvalidSignature)` to candidate event `provider.signature.invalid`.

Future likely examples:

- `provider.signature.invalid`
- `provider.identity.changed`
- `provider.identity.untrusted`
- `provider.identity.reverify.required`
- `provider.message.tamper.detected`
- `provider.epoch.stale`
- `provider.replay.detected`
- `provider.membership.unexpected`
- `provider.public_bundle.unexpected`
- `provider.storage.rollback.detected`
- `provider.commit.invalid`

Expected CarbonStack behavior:

- Append to trust history.
- Warn user when appropriate.
- Block send/open when policy requires.
- Require reverify when identity continuity is broken.
- Preserve enough detail for audit/debug without exposing secrets.

Known policy candidate:

`provider.signature.invalid` should likely map to:

- block the message
- warn the user
- append trust event
- require reverify if identity changed

Important lesson:

The v0.2.8 phase-B fresh-signer failure was good behavior.

Bob rejected a message signed with a different Alice key.

CarbonStack should eventually make this failure understandable instead of collapsing it into a generic message error.

## 7. Terminal / Fatal Provider Errors

Terminal/fatal provider errors describe conditions where the provider cannot safely continue normal operation.

Future likely examples:

- `provider.fatal`
- `provider.storage.unrecoverable`
- `provider.state.inconsistent`
- `provider.group.unrecoverable`
- `provider.unsupported.version`
- `provider.crypto.backend.failed`
- `provider.secret.material.unavailable`
- `provider.invariant.violation`

Expected CarbonStack behavior:

- Stop the affected operation.
- Avoid retry loops that mutate state further.
- Preserve local evidence/debug context.
- Avoid sending new messages from an uncertain state.
- Require explicit recovery flow where appropriate.

User-facing behavior:

Fatal provider errors should not be hidden.

They should produce a clear operational failure and a safe recovery path.

## Event Severity Levels

Suggested severity levels:

- `debug`
- `info`
- `notice`
- `warning`
- `security`
- `fatal`

Suggested default mapping:

| Event class | Default severity |
|---|---|
| lifecycle | debug/info |
| public setup | info/notice |
| membership | notice/warning |
| message | info/warning |
| storage/checkpoint | notice/warning/fatal |
| trust/security | security |
| terminal/fatal provider errors | fatal |

## Event Handling Policy v0

### Append-only history

Provider events should generally be append-only in local logs/history.

Do not silently overwrite security-relevant history.

### Server is not authority

Cypher should not become provider-event authority.

The client/provider should derive and verify provider events locally.

### No secret leakage

Provider events should not include:

- private keys
- signing keys
- raw provider storage
- plaintext except deliberate dev-only fixtures
- raw secret-bearing MLS material
- recovery secrets

### Typed errors

Provider errors should be typed enough to distinguish:

- invalid signature
- missing storage
- corrupt storage
- malformed message
- stale epoch
- duplicate/replay candidate
- unsupported provider version
- fatal invariant failure

## Current Fixture Event Coverage

Current fixture event stream covers:

- provider lifecycle start/end
- public bundle creation
- conversation creation
- Welcome creation/staging
- member add
- join
- message protect/open
- same-process conversation load

Current fixture error coverage includes:

- invalid signature after signer mismatch

Missing fixture coverage:

- missing storage
- corrupt storage
- wrong group ID
- malformed message
- duplicate/replay-ish message
- stale epoch
- unexpected membership change
- public bundle replacement
- fatal provider state failure

## Next Implementation Targets

Recommended immediate implementation targets:

1. Keep Go-side fixture parser/tests passing.
2. Add explicit Go event-name constants or provider event structs.
3. Add fixture review tests for invalid-signature mapping.
4. Add negative-path fixture generation in Rust scratch mode.
5. Map trust/security events to trust-state candidate behavior.
6. Only after that, consider a Rust sidecar command prototype.

## Allowed Claims

Allowed:

- CarbonStack has a provider event taxonomy draft.
- Current OpenMLS fixtures exercise lifecycle, setup, membership, message, load, and invalid-signature concepts.
- The taxonomy is pre-integration and policy-shaping.

## Not Allowed Claims

Not allowed:

- Provider event integration is implemented.
- Trust-state consumes OpenMLS provider events.
- Cypher carries MLS payloads.
- OpenMLS is wired into Comms send/inbox.
- The taxonomy is final.
- Production E2EE exists.
