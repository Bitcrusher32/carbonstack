# CarbonStack Phase 2A Trust-State Plan

## Status

Classification: ACTIVE / PHASE 2A

This document defines the first Phase 2A implementation slice: CarbonStackComms trust-state scaffolding with fake/dev fingerprints.

This is not real cryptography. It is application-level trust machinery that will later constrain real protocol integration.

## Goal

Implement CarbonStack's trust semantics before replacing MockCryptoProvider.

Phase 2A proves:

- devices can have local trust states
- devices can expose fake/dev fingerprints
- users can manually verify a device
- trust history records verification
- strict send blocks unsafe trust states
- development send can warn and allow changed devices
- revoked devices block even in development mode

## Why This Comes Before Real Crypto

Cryptographic libraries can provide keys, sessions, signatures, and message security.

They do not automatically define CarbonStack's trust UX.

CarbonStack must own:

- verified/unverified/changed/revoked states
- key-change response policy
- trust history
- send policy
- local conflict handling
- device revocation semantics

## Implemented Phase 2A Components

In `carbonstack-comms`:

- `internal/trust/trust.go`
- `internal/trust/trust_test.go`
- `scripts/test-trust-lifecycle.ps1`

CLI commands:

- `fingerprint`
- `verify-device`
- `trust-history`
- `simulate-key-change`
- `revoke-device`
- `send --strict`

Local development trust files:

- `trust.json`
- `trust-events.jsonl`

## Trust States

Current states:

- `unknown`
- `unverified`
- `verified`
- `changed`
- `revoked`
- `compromised`

## Send Policy

Development mode:

- unknown device: warn and allow
- unverified device: warn and allow
- changed device: warn and allow
- revoked device: block
- compromised device: block
- verified device: allow

Strict mode:

- unknown device: block
- unverified device: block
- changed device: block until reverified
- revoked device: block
- compromised device: block
- verified device: allow

## Validation Script

Phase 2A trust lifecycle script:

- `carbonstack-comms/scripts/test-trust-lifecycle.ps1`

It validates:

- Alice/Bob state initialization
- invite creation
- invite claim
- device registration
- Bob fingerprint extraction
- Alice verifies Bob
- trust history records `device_verified`
- strict send succeeds after verification
- simulated key change marks Bob as changed
- strict send blocks changed device
- development send warns but allows changed device
- revocation marks Bob as revoked
- revoked device blocks even in development mode
- final trust history records verification, key change, and revocation events

## Current Security Claim

Allowed:

- CarbonStackComms has local trust-state scaffolding.
- CarbonStackComms can simulate verification, key change, and revocation behavior.
- CarbonStackComms can enforce development and strict send policy against trust state.

Not allowed:

- CarbonStack has real cryptographic identity.
- CarbonStack has real encryption.
- CarbonStack has real key agreement.
- CarbonStack has hostile-server tamper detection.
- CarbonStack has secure local vault storage.
- CarbonStack is production secure.

## Relationship To Protocol Work

Phase 2A creates the application-level trust machinery that future providers must satisfy.

Future providers may include:

- Signal-like provider for 1:1
- MLS provider for groups
- Noise provider for transfer/provisioning channels

Provider code should expose cryptographic facts.

CarbonStack application logic should own trust decisions.

## Next Work

Immediate next:

- add Phase 2A trust lifecycle script to validation workflow
- run full validation
- checkpoint after script and docs pass

Later:

- define protocol provider interface
- map trust records to real protocol identity keys
- research libsignal integration feasibility
- research MLS for future groups
- define hostile-server tamper tests
