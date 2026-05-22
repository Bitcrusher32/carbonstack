# CarbonStack MLS Implementation Candidate Notes

## Status

Classification: DRAFT / PHASE 2C RESEARCH

This document records candidate MLS implementations for a future CarbonStack MLS feasibility spike.

It does not select a final provider.

It does not authorize production cryptography claims.

## Current Baseline

CarbonStack has:

- MLS-shaped provider-neutral architecture
- provider-neutral protocol skeleton
- Phase 2A trust-state scaffold
- local validation runner
- no real cryptographic provider

## Goal

Identify the first MLS implementation candidate to investigate for a local-only feasibility spike.

The spike should prove whether MLS can fit CarbonStack's provider boundary and trust model.

## Candidate A: OpenMLS

### Summary

OpenMLS is a Rust implementation of Messaging Layer Security as specified in RFC 9420.

Its documentation describes a high-level API for creating and managing MLS groups. It also supports interchangeable cryptographic provider, key store, and random number generator components.

### Why It Fits

OpenMLS aligns well with CarbonStack's direction:

- Rust implementation
- MLS/RFC 9420 oriented
- group-first model
- public documentation/book exists
- conceptually maps to CarbonStack's provider boundary
- can likely support local-only experimental provider work

### Questions

- What is the exact current crate layout?
- What is the simplest two-member group example?
- What local persistence APIs are required?
- What build requirements exist on Windows?
- How painful is Go/Rust integration later?
- Should the first spike be Rust-only before any Go bridge?

### Initial Fit

First spike candidate: STRONG

Integration risk: MEDIUM

Main concern: Rust/Go boundary and local state persistence complexity.

## Candidate B: mls-rs

### Summary

mls-rs is an AWS Labs MLS implementation.

Its public docs advertise an easy-to-use client interface for managing multiple MLS identities and groups, RFC 9420 conformance, WASM support, and configurable storage including in-memory and SQLite options.

### Why It Might Fit

mls-rs may be attractive because CarbonStack needs:

- multiple identities eventually
- group/conversation state
- local testability
- storage/persistence hooks
- possibly SQLite-aligned persistence ideas

### Questions

- Is the API simpler than OpenMLS for a two-member local spike?
- Is the project maintained enough for this use?
- What license applies?
- Are examples clearer than OpenMLS for CarbonStack's target flow?
- Does the storage abstraction map cleanly to CarbonStack provider state?

### Initial Fit

First spike candidate: POSSIBLE

Integration risk: MEDIUM

Main concern: project fit and API clarity compared with OpenMLS.

## Candidate C: rmls / other Rust MLS libraries

### Summary

Other Rust MLS implementations exist.

They may be useful for comparison but should not be first target unless OpenMLS and mls-rs are unsuitable.

### Initial Fit

First spike candidate: LOW

Use as fallback/reference only.

## Decision Criteria

The first spike candidate should be selected based on:

- RFC 9420 alignment
- maintained implementation
- clear local examples
- two-member group feasibility
- ability to protect/open application messages
- ability to serialize/restore group state
- inspectable epoch/membership state
- license compatibility
- Windows build feasibility
- eventual fit with Go CLI/server glue
- minimal custom cryptography risk

## Recommended First Spike Path

Recommended order:

1. Try OpenMLS first.
2. Keep mls-rs as an alternate if OpenMLS proves too awkward.
3. Keep other implementations as fallback/reference only.

## Spike Shape

First spike should be Rust-local and experimental.

Do not wire it into CarbonStackComms CLI yet.

Suggested path:

- carbonstack-comms/internal/protocol/mls
- include README or doc.go marking the package experimental
- optionally create a tiny Rust scratch subfolder only after dependency shape is understood

The first executable proof should show:

- Alice identity
- Bob identity
- two-member conversation/group
- protected text message
- opened text message
- visible epoch or state version
- visible membership
- state export/import if practical

## Integration Shape Options

### Option D: In-project experimental provider path

Preferred for now.

Path:

- carbonstack-comms/internal/protocol/mls

Benefits:

- keeps the work close to the provider interface
- makes future integration easier
- documents intent inside the project
- allows mock and MLS providers to share vocabulary

Risks:

- may become messy if Rust/Go integration gets complicated
- must be clearly marked experimental

### Sidecar fallback

If direct integration is messy, use a Rust helper binary or protocol lab approach.

Benefits:

- avoids unsafe FFI early
- keeps Go CLI simple
- isolates Rust build complexity

Risks:

- adds process/serialization boundary
- may become awkward later

## Current Recommendation

Proceed with an in-project experimental MLS provider slot, but do not add OpenMLS yet.

Add:

- carbonstack-comms/internal/protocol/mls/README.md

Then research or prototype the smallest Rust-local OpenMLS example.

## Allowed Claims

Allowed:

- CarbonStack is evaluating MLS implementation candidates.
- OpenMLS is the first intended spike candidate.
- mls-rs is a serious alternate candidate.
- No MLS implementation has been integrated yet.

Not allowed:

- CarbonStack uses MLS.
- CarbonStack has real encryption.
- CarbonStack has selected a final MLS implementation.
- CarbonStack has production E2EE.
