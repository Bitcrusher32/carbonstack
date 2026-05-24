# OpenMLS Sidecar KeyPackage Artifact Export Plan v0

Status: Planned
Component: CarbonStackComms / OpenMLS sidecar
Phase: Phase 2D
Previous docs:
- docs/53-openmls-sidecar-public-bundle-export-plan-v0.md
- docs/54-openmls-sidecar-public-bundle-export-result-v0.md

## 1. Purpose

This document defines the next narrow Phase 2D rung after summary-only public-bundle export.

The current OpenMLS sidecar can create dev-only identity state, load/check identity state, and run public-bundle-export to generate a real OpenMLS KeyPackage in memory, compute a sanitized KeyPackage reference, write public-bundle-summary.json, and emit provider.public_bundle.exported.

That current output is intentionally summary-only. It does not write a full serialized KeyPackage artifact. It is not final onboarding material. Another client cannot consume it yet.

This plan defines a controlled dev-only path for writing a full serialized public KeyPackage artifact to ignored local sidecar state, with sanitized stdout and tests.

## 2. Current validated baseline

The following remains true from v0.2.24:

- provider-info is supported.
- identity-create is supported for dev-only identity generation.
- identity-status is supported for dev-only identity load/check.
- public-bundle-export is supported for summary-only KeyPackage generation.
- public-bundle-export generates a real OpenMLS KeyPackage in memory.
- public-bundle-export computes a sanitized sha256 KeyPackage reference.
- public-bundle-export writes public-bundle-summary.json.
- public-bundle-export returns private_material_included=false.
- public-bundle-export returns key_package_artifact_written=false.
- public-bundle-export returns provider_storage_written=false.
- provider.public_bundle.exported is typed and mapped as public setup / info / non-trust-relevant.

## 3. Problem

Conversation lifecycle work needs real consumable public setup material.

The existing public-bundle-summary.json proves that KeyPackage generation happened, but it is not itself the serialized KeyPackage artifact. Building conversation create/add/join on top of summary-only output would risk accidentally designing around a placeholder rather than a real provider object.

Therefore, before conversation lifecycle work, CarbonStack should either:

1. implement full serialized KeyPackage artifact export; or
2. explicitly defer artifact export and document why conversation lifecycle can proceed without it.

Recommendation for this phase: implement full serialized KeyPackage artifact export first.

## 4. Non-goals

This rung must not:

- wire OpenMLS into comms send;
- wire OpenMLS into comms inbox;
- route MLS payloads through CarbonStackCypher;
- mutate trust.json;
- mutate trust-events.jsonl;
- create conversations;
- add members;
- join groups;
- protect messages;
- open messages;
- implement provider storage persistence;
- implement a production secure vault;
- claim production E2EE;
- claim hostile-server proof;
- claim replay resistance;
- claim metadata privacy;
- print or expose signer.json;
- print or expose private key material;
- treat dev signer.json as production secure storage;
- treat the exported artifact as final stable CarbonStack onboarding format.

## 5. Recommended command shape

Keep the existing summary command stable and add artifact behavior explicitly.

Preferred command shape:

    public-bundle-export --device-label <safe> --write-artifact

Alternative command shape:

    public-bundle-artifact-export --device-label <safe>

Recommendation: use the explicit flag on the existing public-bundle-export command.

Reasoning:

- It keeps summary-only behavior backward-compatible.
- It makes artifact writing opt-in and visible.
- It avoids multiplying commands too early.
- It preserves one conceptual operation: export public bundle material.

## 6. Artifact model

The sidecar should write two public files under the ignored per-device dev-state directory:

    public-bundle.keypackage.bin
    public-bundle-manifest.json

The binary artifact should contain the serialized public OpenMLS KeyPackage.

The manifest should contain sanitized metadata only.

Expected manifest fields:

    schema
    device_label
    public_identity_ref
    key_package_ref
    key_package_artifact
    key_package_artifact_sha256
    key_package_artifact_size_bytes
    private_material_included
    provider_storage_written
    warning

The manifest must not contain:

- signer private key material;
- seeds;
- raw private keys;
- provider storage;
- MemoryStorage JSON;
- recovery material;
- secret-bearing identity state;
- message content.

## 7. Stdout behavior

Stdout must remain sanitized.

The sidecar may return:

- success boolean;
- event name;
- device label;
- public identity reference;
- KeyPackage reference;
- artifact relative path;
- artifact sha256;
- artifact size;
- private_material_included=false;
- provider_storage_written=false;
- key_package_artifact_written=true.

The sidecar must not print:

- signer.json contents;
- private key bytes;
- raw secret material;
- provider storage JSON;
- MemoryStorage JSON;
- recovery material.

Recommendation: do not print the serialized KeyPackage bytes to stdout in this rung.

Reasoning:

- Public KeyPackage bytes are not secret by design, but stdout is commonly copied, logged, pasted, and test-captured.
- Path/hash output is safer and easier to reason about.
- The artifact can be inspected by hash and parsed in tests without making stdout a transport channel.

## 8. Serialization path to inspect

Implementation must inspect the correct OpenMLS serialization path before patching.

Expected direction:

- use the public KeyPackage object already created during public-bundle-export;
- serialize the public KeyPackage using OpenMLS-supported serialization / tls_codec path;
- write serialized bytes to public-bundle.keypackage.bin;
- compute sha256 over the exact artifact bytes;
- write public-bundle-manifest.json;
- return sanitized envelope fields.

Do not guess serialization APIs casually. Confirm against the currently pinned OpenMLS dependency set in the sidecar crate.

## 9. Duplicate and regeneration behavior

Default behavior should refuse to overwrite an existing artifact.

If public-bundle.keypackage.bin or public-bundle-manifest.json already exists, the command should fail with a typed event such as:

    provider.public_bundle.exists

or another explicitly defined event if a better name is chosen.

A future force flag may be added later:

    public-bundle-export --device-label <safe> --write-artifact --force

Do not add force behavior in this rung unless needed.

Reasoning:

- Public onboarding material should not be silently replaced.
- Loud replacement semantics align with CarbonStack trust doctrine.
- Future lifecycle work can decide whether KeyPackages are one-time, reusable-dev, or regenerated per enrollment.

## 10. Provider storage semantics

For this rung, provider_storage_written should remain false unless the implementation truly writes OpenMLS provider storage.

This rung is artifact export, not provider storage persistence.

If serialization or KeyPackage generation internally touches provider storage in a meaningful way, document that explicitly before changing the field.

## 11. Event taxonomy

Existing event:

    provider.public_bundle.exported

may continue to represent successful export.

If artifact writing is distinguished from summary export, add a more specific event:

    provider.public_bundle.artifact_exported

Recommendation:

- keep provider.public_bundle.exported for compatibility only if tests and docs clearly distinguish key_package_artifact_written=true;
- otherwise add provider.public_bundle.artifact_exported for clarity.

If duplicate artifact refusal is implemented, add:

    provider.public_bundle.exists

Expected trust mapping:

- artifact exported: public setup / info / non-trust-relevant / append history / debug-only;
- artifact already exists: public setup / warning or recovery depending on final semantics;
- missing identity: stop operation / show recovery path / append history;
- secret material unavailable: fatal / trust-relevant / stop operation.

## 12. Required tests

Go-side contract tests should cover:

- identity creation before export;
- artifact export success;
- missing identity failure;
- unsafe device label rejection;
- duplicate artifact refusal;
- stdout contains no obvious private material;
- signer.json exists but is not printed;
- artifact file exists;
- manifest file exists;
- manifest hash matches artifact hash;
- artifact size is non-zero;
- key_package_artifact_written=true on artifact export;
- private_material_included=false;
- provider_storage_written=false unless actually changed;
- existing summary-only public-bundle-export behavior still works without --write-artifact.

Rust-side tests may cover:

- label validation;
- manifest construction;
- duplicate refusal;
- serialization helper behavior, if factored.

## 13. Manual probe after implementation

From the sidecar crate:

    Remove-Item -Recurse -Force .\.carbonstack-openmls-sidecar-state -ErrorAction SilentlyContinue

    cargo test

    cargo run -- provider-info

    cargo run -- identity-create --device-label carbonstack-alice-device

    cargo run -- identity-status --device-label carbonstack-alice-device

    cargo run -- public-bundle-export --device-label carbonstack-alice-device

    cargo run -- public-bundle-export --device-label carbonstack-alice-device --write-artifact

    cargo run -- public-bundle-export --device-label carbonstack-alice-device --write-artifact

    cargo run -- public-bundle-export --device-label "../bad" --write-artifact

Safe inspection only:

    Get-ChildItem .\.carbonstack-openmls-sidecar-state\dev\devices\carbonstack-alice-device | Select-Object Name, Length

    Get-Content .\.carbonstack-openmls-sidecar-state\dev\devices\carbonstack-alice-device\identity-summary.json

    Get-Content .\.carbonstack-openmls-sidecar-state\dev\devices\carbonstack-alice-device\identity-state.json

    Get-Content .\.carbonstack-openmls-sidecar-state\dev\devices\carbonstack-alice-device\public-bundle-summary.json

    Get-Content .\.carbonstack-openmls-sidecar-state\dev\devices\carbonstack-alice-device\public-bundle-manifest.json

Do not inspect, paste, print, or commit signer.json.

## 14. Artifact guard expectations

Generated files must stay ignored and untracked:

- .carbonstack-openmls-sidecar-state/
- signer.json
- public-bundle-summary.json
- public-bundle.keypackage.bin
- public-bundle-manifest.json
- target/
- .exe
- .pdb
- provider storage JSON
- MemoryStorage JSON
- raw key material

The existing Rust artifact guard should be updated if necessary to catch the new generated files.

## 15. Success criteria

This rung is successful when:

- the plan is committed;
- implementation writes a full serialized public KeyPackage artifact under ignored dev state;
- implementation writes a sanitized manifest;
- stdout remains sanitized;
- tests prove artifact existence and hash consistency;
- duplicate behavior is explicit and tested;
- generated files remain untracked;
- docs record the result;
- no Comms/Cypher/trust runtime integration is added;
- no production security claim is made.

## 16. Failure criteria

This rung fails if:

- private material is printed or committed;
- signer.json is exposed;
- serialized artifact bytes are dumped to stdout without an explicit decision;
- generated dev-state files become tracked;
- summary-only output is described as final onboarding material;
- conversation lifecycle is implemented in the same rung;
- message protect/open is implemented in the same rung;
- Cypher routing is modified in the same rung;
- trust-state storage is mutated in the same rung;
- production E2EE is claimed.

## 17. Recommended next result doc

After implementation and validation, write:

    docs/56-openmls-sidecar-keypackage-artifact-export-result-v0.md

That result doc should record:

- exact command shape implemented;
- exact artifact names and paths;
- exact manifest fields;
- test results;
- artifact guard result;
- whether provider storage remains false;
- whether duplicate export is refused;
- explicit remaining blockers before conversation lifecycle.

## 18. Next after this rung

Only after this rung passes:

1. plan conversation-create/add-member/join;
2. define Welcome artifact handling;
3. define provider storage persistence needs;
4. define how provider events map into trust-state storage;
5. then consider message protect/open.

Do not jump directly from artifact export to runtime messaging integration.
