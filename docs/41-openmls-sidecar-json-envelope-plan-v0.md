# CarbonStack OpenMLS Sidecar JSON Envelope Plan v0

## Status

Classification: PHASE 2D SIDECAR ENVELOPE PLAN / PRE-SECRET-COMMAND

This document defines the JSON response/error envelope shape for the experimental OpenMLS sidecar.

It does not implement identity creation.

It does not implement public bundle export.

It does not implement conversation creation.

It does not implement message protect/open.

It does not wire OpenMLS into CarbonStackComms user-facing commands.

It does not route MLS traffic through CarbonStackCypher.

It does not mutate trust-state storage.

It does not claim production E2EE.

## Source Context

Relevant docs:

- `docs/39-phase2d-sidecar-command-surface-plan.md`
- `docs/40-openmls-sidecar-provider-info-result-v0.md`

Relevant sidecar path:

- `carbonstack-comms/internal/protocol/mls/research/openmls-sidecar`

Relevant Go-side test:

- `carbonstack-comms/internal/protocol/openmls_sidecar_provider_info_test.go`

## Purpose

Before adding secret-bearing or state-mutating sidecar commands, every sidecar command should return one predictable JSON envelope.

This prevents future command outputs from drifting into ad-hoc JSON.

It also gives Go-side tests a stable parse target before Comms CLI or Cypher integration begins.

## Design Principles

The sidecar envelope must be:

- JSON-only
- explicit about success/failure
- explicit about the command being handled
- explicit about provider identity
- explicit about warnings
- explicit about provider events
- explicit about whether private material is included
- usable from Go tests
- safe to log at summary level unless a command is explicitly marked secret-bearing

The sidecar envelope must not:

- print private keys
- print signer material
- print provider storage
- print MemoryStorage JSON
- print recovery secrets
- silently mutate state without event/checkpoint fields
- hide provider trust/security failures as generic errors
- use different shapes for every command

## Success Envelope v0

Every successful sidecar command should return:

```json
{
  "ok": true,
  "command": "provider-info",
  "provider": "openmls",
  "implementation": "carbonstack-openmls-sidecar",
  "mode": "experimental-sidecar",
  "phase": "phase2d-provider-info",
  "data": {},
  "events": [],
  "warnings": [],
  "private_material_included": false
}
```

Required top-level fields:

- `ok`
- `command`
- `provider`
- `implementation`
- `mode`
- `phase`
- `data`
- `events`
- `warnings`
- `private_material_included`

## Error Envelope v0

Every sidecar command failure should return JSON too.

Proposed shape:

```json
{
  "ok": false,
  "command": "identity-create",
  "provider": "openmls",
  "implementation": "carbonstack-openmls-sidecar",
  "mode": "experimental-sidecar",
  "phase": "phase2d-provider-info",
  "error": {
    "code": "unsupported_command",
    "message": "unsupported command: identity-create",
    "provider_event": "provider.command.unsupported",
    "severity": "warning",
    "trust_relevant": false
  },
  "events": [
    {
      "event": "provider.command.unsupported",
      "severity": "warning",
      "trust_relevant": false
    }
  ],
  "warnings": [],
  "private_material_included": false
}
```

Required error object fields:

- `code`
- `message`
- `provider_event`
- `severity`
- `trust_relevant`

## Provider Info Data Shape v0

After envelope migration, `provider-info` should place command-specific information under `data`.

Proposed `data` shape:

```json
{
  "capabilities": [
    "provider-info"
  ],
  "unsupported": [
    "identity-create",
    "public-bundle-export",
    "conversation-create",
    "conversation-add-member",
    "conversation-join",
    "message-protect",
    "message-open",
    "state-checkpoint",
    "state-load-check"
  ],
  "security_level": "experimental; not production E2EE"
}
```

Top-level warnings should preserve the current boundary warnings:

- OpenMLS is not wired into CarbonStackComms.
- Cypher does not route MLS payloads.
- trust-state storage does not consume provider events.
- no secret-bearing sidecar commands are implemented.

## Unsupported Command Behavior v0

Unsupported commands should:

- return JSON error envelope
- set `ok` to `false`
- set `error.code` to `unsupported_command`
- set `error.provider_event` to `provider.command.unsupported`
- set severity to `warning`
- set trust relevance to `false`
- set `private_material_included` to `false`
- use a nonzero exit code

Recommended exit code:

- `2` for unsupported command / bad usage

## Provider Events in Sidecar Envelopes

Sidecar envelopes should eventually align with the provider event taxonomy.

Current event taxonomy source:

- `docs/36-provider-event-taxonomy-v0.md`
- `carbonstack-comms/internal/protocol/provider_events.go`

New event candidate:

- `provider.command.unsupported`

Classification candidate:

- class: lifecycle or unknown/operational
- severity: warning
- trust relevant: false

This event exists to make unsupported-command failures machine-readable without implying a cryptographic trust failure.

## Trust Mapping Boundary

The sidecar can include event candidates.

The sidecar must not directly mutate trust state.

CarbonStackComms may later map provider events to trust decisions, but that remains separate from the sidecar itself.

## Command Migration Order

Recommended next implementation order:

1. migrate `provider-info` success output into envelope shape
2. make unsupported commands return JSON error envelope
3. update Go-side provider-info test
4. add Go-side unsupported-command error test
5. document result
6. only then consider `identity-create`

## Allowed Claims

Allowed:

- CarbonStack has a planned JSON envelope shape for the experimental OpenMLS sidecar.
- The envelope plan is intended to stabilize sidecar success/error parsing before secret-bearing commands.
- Unsupported commands should be machine-readable JSON failures.

## Not Allowed Claims

Not allowed:

- sidecar envelope is implemented
- identity creation exists
- public bundle export exists
- message protect/open exists
- Comms CLI consumes sidecar envelopes
- Cypher routes MLS payloads
- trust-state consumes sidecar events
- production E2EE exists
