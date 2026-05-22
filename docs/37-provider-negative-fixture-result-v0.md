# CarbonStack Provider Negative Fixture Result v0

## Status

Classification: PHASE 2C PROVIDER-CONTRACT NEGATIVE-FIXTURE RESULT / PRE-INTEGRATION

This document records the first summary-only negative provider fixture set.

It does not wire OpenMLS into CarbonStackComms.

It does not route MLS traffic through CarbonStackCypher.

It does not implement trust-state integration.

It does not claim production E2EE.

## Source Context

Relevant docs:

- `docs/35-openmls-provider-fixture-result-v0.md`
- `docs/36-provider-event-taxonomy-v0.md`

Relevant Comms paths:

- `carbonstack-comms/internal/protocol/openmls_negative_fixture_test.go`
- `carbonstack-comms/internal/protocol/provider_events.go`
- `carbonstack-comms/internal/protocol/provider_events_test.go`
- `carbonstack-comms/internal/protocol/mls/research/openmls-minimal/fixtures/dev`

## What Was Added

Four summary-only negative fixtures were added to the OpenMLS dev fixture set:

- `missing-storage-error.json`
- `missing-signer-error.json`
- `wrong-group-error.json`
- `malformed-message-error.json`

Go-side tests parse these files and compare their provider-event candidates against the current provider event taxonomy.

## Negative Fixture Mappings

### Missing storage

Fixture:

- `missing-storage-error.json`

Candidate provider event:

- `storage.missing`

Class:

- `storage_checkpoint`

Severity:

- `warning`

Trust relevance:

- `false`

Suggested behavior:

- stop operation
- do not send
- show recovery path
- do not mutate provider state further

### Missing signer

Fixture:

- `missing-signer-error.json`

Candidate provider event:

- `provider.secret.material.unavailable`

Class:

- `terminal_fatal`

Severity:

- `fatal`

Trust relevance:

- `true`

Suggested behavior:

- block send
- preserve local evidence
- show recovery path
- require explicit identity recovery flow

### Wrong group ID

Fixture:

- `wrong-group-error.json`

Candidate provider event:

- `provider.group.unrecoverable`

Class:

- `terminal_fatal`

Severity:

- `fatal`

Trust relevance:

- `true`

Suggested behavior:

- stop operation
- do not send
- do not open message
- surface provider state mismatch

### Malformed provider message

Fixture:

- `malformed-message-error.json`

Candidate provider event:

- `provider.message.tamper.detected`

Class:

- `trust_security`

Severity:

- `security`

Trust relevance:

- `true`

Suggested behavior:

- block message
- append trust event
- warn user
- retain sanitized diagnostic summary

## Why This Matters

The v0.2.10 provider event taxonomy gave CarbonStack names and classifications.

These negative fixtures make the taxonomy executable against concrete JSON fixture examples.

This keeps safety and failure semantics ahead of runtime integration.

## Important Encoding Lesson

The first PowerShell write used `Set-Content -Encoding UTF8`, which wrote JSON files with a UTF-8 BOM on the user's Windows PowerShell environment.

Go's `encoding/json` rejected those files with:

- `invalid character 'ï' looking for beginning of value`

Fix:

- rewrite JSON files as UTF-8 without BOM using `[System.Text.UTF8Encoding]::new($false)` and `[System.IO.File]::WriteAllText(...)`.

Future rule:

- For JSON fixture files parsed by Go, write UTF-8 without BOM.

## Current Validation Meaning

Current validation means:

- summary-only negative fixtures exist
- Go can parse the negative fixtures
- provider event candidates match the current Go taxonomy
- event class/severity/trust relevance are checked by tests

It does not mean:

- OpenMLS runtime emitted these exact failures
- Comms handles these errors at runtime
- trust-state consumes provider errors
- Cypher carries provider messages
- production storage/recovery is solved

## Fixture Safety Rules

Negative fixtures must remain summary-only.

Do not include:

- private keys
- signer JSON
- MemoryStorage JSON
- provider storage
- raw MLS secret material
- real recovery material
- generated build artifacts

## Next Recommended Work

Recommended next step:

- define trust-state mapping behavior for provider events/errors, still as docs or pure mapping tests before runtime integration.

Candidate next doc:

- `docs/38-provider-trust-state-mapping-v0.md`

Candidate mapping targets:

- `provider.signature.invalid`
- `provider.message.tamper.detected`
- `provider.secret.material.unavailable`
- `provider.group.unrecoverable`
- `storage.missing`
- `checkpoint.failed`

Keep this pre-integration.

Do not wire OpenMLS into `comms send` / `comms inbox` yet.

## Allowed Claims

Allowed:

- CarbonStack has summary-only negative provider fixtures.
- Go-side provider tests parse and classify negative fixture examples.
- Negative fixture work strengthens the provider-contract layer before runtime integration.

## Not Allowed Claims

Not allowed:

- OpenMLS runtime error integration is implemented.
- CarbonStackComms handles provider errors in CLI flows.
- trust-state consumes provider events.
- Cypher carries MLS traffic.
- production E2EE exists.
- hostile-server security is solved.
