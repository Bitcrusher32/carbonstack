# CarbonStack Phase 1 Test Matrix

## Status

Classification: ACTIVE / DEVELOPMENT

This document tracks Phase 1 validation coverage for the CarbonStack relay/client skeleton.

Phase 1 tests validate implementation shape only. They do not validate production security, real encryption, metadata privacy, endpoint compromise resistance, or Signal-equivalent behavior.

## Test Scope

Phase 1 covers:

- CarbonStackCypher local relay behavior
- CarbonStackComms CLI behavior
- CLI-to-Cypher local lifecycle
- stub/base64 MockCryptoProvider behavior
- local plaintext development state behavior

Phase 1 does not cover:

- real E2EE cryptography
- hardware-key enrollment
- local secure vault
- Android implementation
- CarbonStackOS
- production auth
- replay resistance
- metadata privacy
- hostile-server proof

## Validation Matrix

| Area | Test / Validation | Location | Status |
|---|---|---|---|
| Cypher health endpoint | `GET /v0/health` | `carbonstack-cypher/internal/httpapi/api_test.go` | VALIDATED |
| Cypher dev invite creation | `POST /v0/dev/invites` | `carbonstack-cypher/internal/httpapi/api_test.go` | VALIDATED |
| Cypher invite claim | `POST /v0/invites/claim` | `carbonstack-cypher/internal/httpapi/api_test.go` | VALIDATED |
| Cypher device registration | `POST /v0/devices/register` | `carbonstack-cypher/internal/httpapi/api_test.go` | VALIDATED |
| Cypher device lookup | `GET /v0/accounts/{account_id}/devices` | `carbonstack-cypher/internal/httpapi/api_test.go` | VALIDATED |
| Cypher envelope submit | `POST /v0/envelopes` | `carbonstack-cypher/internal/httpapi/api_test.go` | VALIDATED |
| Cypher envelope retrieval | `GET /v0/devices/{device_id}/envelopes` | `carbonstack-cypher/internal/httpapi/api_test.go` | VALIDATED |
| Cypher envelope acknowledgement | `POST /v0/envelopes/{envelope_id}/ack` | `carbonstack-cypher/internal/httpapi/api_test.go` | VALIDATED |
| Cypher wrong-recipient ack rejection | reject mismatched recipient ack | `carbonstack-cypher/internal/httpapi/api_test.go` | VALIDATED |
| Cypher invalid base64 rejection | reject malformed ciphertext_b64 | `carbonstack-cypher/internal/httpapi/api_test.go` | VALIDATED |
| Comms mock crypto round trip | base64 stub encode/decode | `carbonstack-comms/internal/crypto/mock_test.go` | VALIDATED |
| Comms state save/load | local JSON development state | `carbonstack-comms/internal/state/state_test.go` | VALIDATED |
| Comms Cypher client wrapper | HTTP client methods against test server | `carbonstack-comms/internal/client/cypher_test.go` | VALIDATED |
| CLI lifecycle smoke test | Alice/Bob lifecycle through CLI | `carbonstack-comms/scripts/test-local-lifecycle.ps1` | VALIDATED |
| Real cryptography | mature protocol integration | future work | NOT VALIDATED |
| Hostile-server harness | malicious relay behavior tests | future work | NOT VALIDATED |
| Android client | Android app lifecycle | future work | NOT VALIDATED |
| CarbonStackOS | appliance OS build/test | future work | DEFERRED |

## Required Commands

Run Cypher tests:

```powershell
cd C:\Users\udaiv\repos\carbonstack_umbrella\carbonstack-cypher
go test ./...
Run Comms package tests:

cd C:\Users\udaiv\repos\carbonstack_umbrella\carbonstack-comms
go test ./...

Run Comms CLI lifecycle smoke test:

cd C:\Users\udaiv\repos\carbonstack_umbrella\carbonstack-comms
powershell -ExecutionPolicy Bypass -File .\scripts\test-local-lifecycle.ps1

The lifecycle smoke test expects CarbonStackCypher to already be running:

cd C:\Users\udaiv\repos\carbonstack_umbrella\carbonstack-cypher
go run .\cmd\cypher
Allowed Claims After Passing Phase 1 Tests

Allowed:

CarbonStackCypher has automated API tests for the Phase 1 local relay lifecycle.
CarbonStackComms has package tests for state, mock crypto, and Cypher client wrappers.
CarbonStackComms has a PowerShell lifecycle smoke test for the CLI-driven Alice/Bob flow.
The Phase 1 relay/client skeleton is test-protected at a local development level.

Not allowed:

CarbonStack is production secure.
CarbonStack is audited.
CarbonStack is Signal-equivalent.
Stub/base64 payloads provide encryption.
Metadata privacy is solved.
Production authentication is solved.
Hostile-server security is validated.
Android or CarbonStackOS are implemented.
Next Test Work

Recommended next additions:

Comms CLI command-level tests.
Cypher DB-layer tests for migrations and invite seeding.
Negative-path tests for missing fields and unsupported content/protocol versions.
Test-mode server launch script that uses a temporary DB.
Hostile-server behavior tests after real protocol selection.
