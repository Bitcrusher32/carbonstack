# CarbonStack Phase 2D Provider Sidecar Command Surface Plan v0

## Status

Classification: PHASE 2D PLANNING / PRE-IMPLEMENTATION

This document defines the first intended runtime-provider integration shape after Phase 2C.

It does not implement the sidecar.

It does not wire OpenMLS into CarbonStackComms.

It does not route MLS traffic through CarbonStackCypher.

It does not claim production E2EE.

## Source Context

Phase 2C validated:

- OpenMLS dependency/build path
- credential and KeyPackage creation
- Alice group creation
- Bob add from KeyPackage
- Welcome join
- application message protect/open
- two-message state continuity
- same-process provider-storage reload
- MemoryStorage file persistence across fresh provider construction
- signer persistence requirement
- sanitized positive provider fixtures
- sanitized negative provider fixtures
- Go-side fixture parsing
- provider event taxonomy
- pure provider-event to trust-decision mapping
- fixture-backed trust-decision tests

The next phase should not jump directly to production integration.

The next phase should create a narrow local sidecar command surface.

## Decision

Use a Rust sidecar command prototype as the first runtime integration path.

Do not use FFI first.

Do not rewrite the provider in Go first.

Do not wire sidecar behavior into user-facing Comms CLI commands until the sidecar command contract is stable.

## Why Sidecar First

Sidecar-first is preferred because it:

- keeps Go/Rust coupling low
- keeps OpenMLS isolated
- makes debugging easier
- allows command-level contract tests
- can be replaced later if the architecture changes
- avoids early FFI complexity
- fits the fixture-first provider-contract work already completed

## Sidecar Design Constraints

The sidecar must be:

- local-only at first
- explicit about input/output JSON
- deterministic enough for tests where possible
- clear about state mutation
- clear about checkpoint requirements
- hostile to silent failure
- free of production-security claims

The sidecar must not:

- expose private keys in normal output
- write secrets into the repo
- depend on Cypher
- depend on Comms CLI state initially
- silently mutate state without an event/checkpoint result
- hide provider trust/security failures as generic errors

## Proposed Command Surface v0

### `provider-info`

Purpose:

- Report provider name, version, supported modes, and capabilities.

Candidate output:

- provider name
- provider implementation
- ciphersuite support summary
- storage mode summary
- warning that this is experimental

### `identity-create`

Purpose:

- Create local provider identity material for a device.

Candidate output:

- device label
- public bundle summary
- provider events
- storage checkpoint requirement

Must not output:

- raw private signing key
- raw provider storage
- recovery secret

### `public-bundle-export`

Purpose:

- Export public setup material for another device/user to consume.

Candidate output:

- public bundle envelope
- hash/reference summary
- provider events

### `conversation-create`

Purpose:

- Create a local MLS conversation/group as the initiating device.

Candidate output:

- local conversation ID
- provider group ID
- epoch summary
- member summary
- provider events
- checkpoint result

### `conversation-add-member`

Purpose:

- Add a member from public setup material.

Candidate output:

- Welcome/join material envelope
- commit summary
- epoch summary
- member summary
- provider events
- checkpoint result

### `conversation-join`

Purpose:

- Join from Welcome/join material.

Candidate output:

- local conversation ID
- provider group ID
- epoch summary
- member summary
- provider events
- checkpoint result

### `message-protect`

Purpose:

- Convert plaintext into provider-protected message material.

Candidate input:

- local conversation ID
- plaintext
- optional associated data / envelope metadata

Candidate output:

- protected provider message envelope
- epoch summary
- provider events
- checkpoint result

Important:

- this is state-mutating
- checkpoint-after-send is required

### `message-open`

Purpose:

- Process/open provider-protected message material.

Candidate input:

- local conversation ID
- protected provider message envelope

Candidate output:

- plaintext or blocked result
- sender summary
- epoch summary
- provider events
- trust/security decision candidate
- checkpoint result

Important:

- this is state-mutating
- checkpoint-after-receive is required

### `state-checkpoint`

Purpose:

- Explicitly save/checkpoint provider state if not automatically handled.

Candidate output:

- checkpoint status
- state version/summary
- provider events

### `state-load-check`

Purpose:

- Validate that provider state can be loaded for a conversation/device.

Candidate output:

- load status
- group/conversation summary
- epoch/member summary
- provider events
- typed error if missing/corrupt/unrecoverable

## Error Surface v0

Sidecar errors should be typed.

Required early error mappings:

- `storage.missing`
- `storage.corrupt`
- `checkpoint.failed`
- `provider.signature.invalid`
- `provider.message.tamper.detected`
- `provider.group.unrecoverable`
- `provider.secret.material.unavailable`
- `provider.state.inconsistent`
- `provider.invariant.violation`

Errors should include:

- event name
- class
- severity
- trust relevance
- candidate trust decision
- sanitized detail string
- no secret material

## Event Surface v0

Sidecar success and failure outputs should include provider events.

Events should align with:

- `docs/36-provider-event-taxonomy-v0.md`
- `docs/38-provider-trust-state-mapping-v0.md`
- `carbonstack-comms/internal/protocol/provider_events.go`
- `carbonstack-comms/internal/protocol/provider_trust.go`

## State Model v0

Sidecar state must be device-local.

Lessons carried forward:

- Alice and Bob cannot share provider storage.
- `create_message` is state-mutating.
- `process_message` is state-mutating.
- signer identity persistence is required.
- provider storage persistence is required.
- checkpoint failures must not be ignored.

## Integration Boundary v0

Initial sidecar work should be called by tests or dev scripts, not by user-facing Comms commands.

Recommended order:

1. Implement sidecar `provider-info`.
2. Implement sidecar fixture-equivalent command if useful.
3. Implement sidecar identity/public-bundle commands.
4. Implement sidecar local Alice/Bob create/join/protect/open smoke command.
5. Only then decide how Go Comms should invoke it.
6. Only later wire into `comms send` / `comms inbox`.
7. Only later route provider envelopes through Cypher.

## Phase 2C Closure Criteria

Phase 2C is considered mature enough to transition if:

- OpenMLS scratch feasibility is validated.
- provider fixtures exist and are Go-parsed.
- negative fixtures exist and are Go-parsed.
- provider event taxonomy exists in docs and Go tests.
- provider-trust decision mapping exists in docs and Go tests.
- sidecar command-surface plan exists.
- no production security claims are made.

## Recommended Phase 2D First Step

Add a new experimental sidecar directory, likely under:

- `carbonstack-comms/internal/protocol/mls/sidecar/openmls-provider`

or keep it near the scratch crate until stable:

- `carbonstack-comms/internal/protocol/mls/research/openmls-sidecar`

Initial command:

- `provider-info`

This should be intentionally boring.

Success criteria:

- builds locally
- outputs JSON
- has no access to secrets
- is callable from a script/test
- does not affect Comms CLI behavior

## Allowed Claims

Allowed:

- CarbonStack has a Phase 2D sidecar command-surface plan.
- Rust sidecar is the first intended runtime-provider integration path.
- FFI is deferred.
- Comms/Cypher integration remains intentionally deferred.

## Not Allowed Claims

Not allowed:

- sidecar is implemented
- OpenMLS is wired into Comms
- Cypher routes MLS payloads
- production E2EE exists
- production storage is solved
- trust-state integration is complete
