# OpenMLS Sidecar Conversation Create Persistence Repair Plan v0

Status: Planned
Component: CarbonStackComms / OpenMLS sidecar
Phase: Phase 2D
Related docs:
- docs/59-openmls-sidecar-conversation-create-plan-v0.md
- docs/60-openmls-sidecar-conversation-create-result-v0.md
- docs/61-openmls-sidecar-conversation-add-member-welcome-plan-v0.md
- docs/62-openmls-sidecar-add-member-welcome-api-recon-v0.md

## 1. Purpose

This document records a required repair before implementing conversation-add-member / Welcome export.

The v0.2.28 recon gate asked whether v0.2.27 conversation-create leaves reloadable OpenMLS group/provider state for a later command invocation.

The answer is not yet proven.

Current observed behavior:

- conversation-create creates an OpenMLS MlsGroup in-process;
- conversation-create writes sanitized conversation-summary.json;
- generated conversation state visibly contains only conversation-summary.json;
- no explicit persisted group/provider state artifact is visible;
- conversation-create currently reports provider_storage_written=true.

This is too strong. The next checkpoint must repair this before add-member.

## 2. Problem

conversation-add-member requires a usable mutable MlsGroup.

OpenMLS add_members requires:

    &mut MlsGroup
    provider
    signer
    &[KeyPackage]

A later sidecar command cannot safely call add_members unless it can load the conversation/group state created by conversation-create.

Current conversation-create does not prove that.

## 3. Required correction

Until reloadability is proven, conversation-create should not claim:

    provider_storage_written=true

Recommended near-term correction:

    provider_storage_written=false
    group_reloadable=false

or, if a reload probe is implemented and passes:

    provider_storage_written=true
    group_reloadable=true

The project must not keep a true provider-storage claim based only on an in-process MlsGroup object.

## 4. Planned implementation options

### Option A — honesty repair only

Change conversation-create to report:

    provider_storage_written=false

Add fields:

    group_reloadable=false

This is safest if explicit persistence is not yet implemented.

Result:

- add-member remains blocked;
- docs/result records that conversation-create is a local group creation summary, not reloadable provider storage;
- next checkpoint plans real provider/group persistence.

### Option B — implement reloadable dev provider/group persistence

Find and implement a dev-local storage path that allows:

    MlsGroup::load(provider.storage(), &group_id)

to succeed in a later command invocation.

This requires confirming:

- how OpenMlsRustCrypto stores state;
- whether MemoryStorage can be serialized;
- whether prior openmls-minimal persistence spike can be adapted;
- whether state should live under device path, conversation path, or both.

Result:

- add-member may proceed after reload-check passes;
- provider_storage_written=true becomes honest;
- group_reloadable=true can be tested.

### Recommendation

Start with a narrow reload-check probe.

If reload-check fails and persistence requires deeper storage work, do Option A first and record the blocker.

Do not implement add-member until reload-check passes.

## 5. Proposed command/test surface

Potential new command:

    conversation-status --conversation-label <safe>

or:

    conversation-load-check --device-label <safe> --conversation-label <safe>

Recommended for this repair:

    conversation-load-check --device-label <safe> --conversation-label <safe>

Purpose:

- validate device label;
- validate conversation label;
- require identity state;
- require conversation-summary.json;
- derive the same GroupId bytes from conversation label;
- attempt MlsGroup::load(provider.storage(), &group_id);
- report group_reloadable true/false;
- keep stdout sanitized.

If no persisted provider storage exists, expected result should be:

    group_reloadable=false

This is not a fatal security event; it is a dev persistence limitation.

## 6. Tests required

Go-side tests should cover:

- conversation-create success still writes conversation-summary.json;
- conversation-create no longer overclaims provider_storage_written unless reload is proven;
- conversation-load-check missing identity fails;
- conversation-load-check missing conversation fails;
- conversation-load-check after current conversation-create reports group_reloadable=false if persistence is absent;
- stdout contains no secret material;
- conversation-add-member remains unsupported.

If persistence is implemented:

- conversation-load-check after conversation-create reports group_reloadable=true;
- MlsGroup::load succeeds;
- epoch/member count match expected summary;
- generated provider/group state remains ignored and untracked.

## 7. Non-goals

This repair must not implement:

- conversation-add-member;
- Welcome export;
- conversation-join;
- message-protect;
- message-open;
- Comms runtime integration;
- Cypher routing;
- trust-state mutation;
- production secure vault;
- Android;
- CarbonStackOS.

## 8. Success criteria

The repair succeeds when:

- the provider_storage_written claim is honest;
- reloadability is explicitly reported or explicitly blocked;
- tests protect the claim;
- docs record the result;
- add-member remains blocked unless reloadability is proven.
