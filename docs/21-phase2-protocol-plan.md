# CarbonStack Phase 2 Protocol Plan

## Status

Classification: DRAFT / PHASE 2 PLANNING

This document defines the transition from Phase 1 lifecycle scaffolding to Phase 2 protocol research and integration.

Phase 2 must not begin by writing custom cryptography.

## Current Baseline

Phase 1 has validated:

- CarbonStackCypher local relay lifecycle
- CarbonStackComms CLI lifecycle
- invite/account/device/envelope flow
- package tests
- API tests
- lifecycle smoke test
- canonical cross-repo validation runner

Phase 1 has not validated:

- real encryption
- sender authentication
- forward secrecy
- post-compromise recovery
- replay resistance
- hostile-server proof
- hardware-key identity
- local secure vault
- Android client
- CarbonStackOS

## Phase 2 Goal

Phase 2 should replace the Phase 1 MockCryptoProvider with a protocol-shaped implementation path while preserving the validated client/server lifecycle.

Phase 2 should answer:

- What is a CarbonStack device identity?
- What is a CarbonStack trust record?
- How does one device establish a session with another?
- How are key changes detected?
- How are revoked devices handled?
- What protocol library or standard should be used first?
- What tests must pass before claims change?

## Recommended Phase 2 Order

### Step 1: Identity and Trust State

Define local identity, device records, trust states, key-change behavior, and revocation semantics.

Status: started in docs/20-identity-and-trust-state-v0.md

### Step 2: Protocol Candidate Feasibility

Evaluate the practical integration path for:

- Signal Protocol / libsignal
- MLS
- Noise
- libsodium or NaCl-style primitives

Evaluation must include:

- license
- supported languages
- bindings
- testability
- maintenance state
- fit for CLI first
- fit for Android later
- fit for hostile-server assumptions
- fit for future groups

### Step 3: Protocol Harness Design

Before implementation, define a local protocol harness.

Required tests should include:

- Alice and Bob identity creation
- first session establishment
- message encrypt/decrypt
- wrong-recipient decrypt failure
- sender identity verification
- replay detection or replay warning
- key-change detection
- revoked device behavior
- persistence and reload
- server cannot decrypt envelope contents

### Step 4: MockCryptoProvider Replacement Boundary

Define an interface boundary that allows swapping:

- MockCryptoProvider
- future Signal-style provider
- future MLS provider
- future transport provider

The Comms CLI should not care which provider is used beyond explicit command mode and warnings.

### Step 5: 1:1 Protocol Spike

Implement a non-production 1:1 protocol spike only after the trust model and harness are defined.

The spike should not change allowed security claims until tests and review justify it.

## Candidate Direction

Current provisional direction:

- 1:1 messaging: investigate Signal Protocol or libsignal first
- future groups: investigate MLS
- transfer or appliance channels: investigate Noise
- primitives: use libsodium only for narrow utilities or non-production experiments

## Phase 2 Deliverables

Minimum docs:

- identity and trust state model
- protocol feasibility matrix
- protocol harness plan
- MockCryptoProvider replacement plan
- updated threat model
- updated allowed and not-allowed claims

Minimum code, later:

- protocol package skeleton
- test vectors or deterministic fixtures
- provider interface
- negative tests
- local state migration from Phase 1 development state

## Phase 2 Non-Goals

Do not include:

- Android app
- CarbonStackOS
- group messaging implementation
- attachments
- media
- production authentication
- production deployment
- hardware-key-required release mode
- security claims beyond what tests validate

## Exit Criteria

Phase 2 planning is complete when the project has:

- clear identity/trust model
- selected first protocol integration candidate or feasibility spike
- documented rejected paths
- protocol test harness design
- updated test matrix
- no custom cryptography dependency
- no inflated claims

Phase 2 implementation should begin only after those planning artifacts are committed and the Phase 1 validation runner still passes.

## Next Action

Complete protocol-candidate feasibility research and map each candidate against the identity/trust requirements.

Recommended next file:

- docs/22-protocol-feasibility-matrix.md
