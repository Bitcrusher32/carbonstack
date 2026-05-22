# CarbonStack OpenMLS Scratch Result v0

## Status

Classification: PHASE 2C RESEARCH RESULT / SCRATCH ONLY

This document records the first successful local OpenMLS scratch sequence for CarbonStack.

It does not claim production encryption.

It does not claim CarbonStackComms uses OpenMLS for real messaging.

It does not select OpenMLS as the final protocol provider.

## Purpose

CarbonStack needed to answer a basic feasibility question:

Can a real MLS implementation support CarbonStack's long-term conversation model?

The answer at this stage is:

Yes, at scratch level, OpenMLS can create a two-member group, add a member, process a Welcome, and protect/open one application message.

This is a meaningful feasibility signal, but not production integration.

## Current Repos At v0.2.5

Known current heads from the v0.2.5 checkpoint:

- carbonstack: `5ec9c31 docs: add OpenMLS upstream example notes`
- carbonstack-comms: `d75bf46 test: probe OpenMLS application message flow`
- carbonstack-cypher: `0bfd5af chore: remove tracked Cypher local state artifacts`
- carbonstack-os: `b537475 Add CarbonStackOS north star and initial appliance model`

## Scratch Location

Current Rust-only OpenMLS scratch path:

- `carbonstack-comms/internal/protocol/mls/research/openmls-minimal`

Expected committed source files:

- `Cargo.toml`
- `Cargo.lock`
- `README.md`
- `src/main.rs`

Do not commit:

- `target/`
- generated rustdoc
- `.fingerprint/`
- `debug/`
- `release/`
- `.exe`
- `.pdb`
- `.o`

## Dependency Set

The scratch crate currently uses:

- `openmls`
- `openmls_rust_crypto`
- `openmls_basic_credential`
- `tls_codec`

## Validated Scratch Ladder

### 1. Dependency / build probe

Validated:

- Rust installed.
- Cargo works.
- OpenMLS-related crates resolve locally.
- Scratch crate builds on the local Windows/MSVC setup.
- Scratch crate can run without Comms/Cypher integration.

### 2. Credential / KeyPackage probe

Validated:

- basic credential creation
- signature key creation
- KeyPackageBundle creation
- public KeyPackage extraction
- KeyPackage hash reference inspection

Important lesson:

- `KeyPackage::builder().build(...)` returned a `KeyPackageBundle`.
- `MlsGroup::add_members(...)` needs a public `KeyPackage`.
- Working extraction:

  `key_package_bundle.key_package().clone()`

### 3. Group-add probe

Validated:

- Alice setup material exists.
- Bob setup material exists.
- Alice can create an `MlsGroup`.
- Alice can add Bob using Bob's public KeyPackage.
- OpenMLS emits commit and Welcome material.
- Alice can merge the pending commit.
- Alice can inspect member count and epoch.

### 4. Welcome join probe

Validated:

- Welcome material can be extracted from the `MlsMessageOut` body.
- Bob can stage the Welcome.
- Bob can turn staged Welcome into his own `MlsGroup`.
- Alice and Bob both see a two-member group.
- Alice and Bob must use separate provider/storage instances.

Important failed path:

- Serializing full `MlsMessageOut` bytes and deserializing them directly as raw `Welcome` failed with `TrailingData`.

Working lesson:

- `MlsMessageOut` is a wrapper.
- Extract Welcome by inspecting `welcome_msg.body()` and matching `MlsMessageBodyOut::Welcome(welcome)`.

Important storage lesson:

- One shared `OpenMlsRustCrypto::default()` provider/storage for Alice and Bob caused `GroupAlreadyExists`.
- Alice and Bob need separate provider/storage instances.
- Design implication: OpenMLS provider storage is device-local protocol state.

### 5. Application message probe

Validated:

- Alice can create an MLS application message.
- The `MlsMessageOut` application message can be serialized.
- It can be deserialized as `MlsMessageIn`.
- It can be converted into `ProtocolMessage`.
- Bob can process the message.
- `ProcessedMessageContent::ApplicationMessage(...)` can be extracted.
- `ApplicationMessage::into_bytes()` returns opened plaintext.
- Bob-opened plaintext matches Alice plaintext.

Important lesson:

- Processing an incoming application message mutates Bob's group/provider state.
- `bob_group` had to be mutable.
- Design implication: provider `open/process` is not a pure read/decrypt operation. It may update local state and must be persisted later.

## Current Feasibility Conclusion

OpenMLS is viable enough to continue Phase 2C.

It has now proven the minimum local shape CarbonStack needed:

- group-shaped conversations
- two-member 1:1-as-group model
- explicit setup material
- Welcome-based join
- application-message protect/open
- local epoch/member inspection

However, this is still only a local scratch experiment.

## What This Does Not Prove

This does not prove:

- production E2EE
- CarbonStackComms OpenMLS integration
- Cypher delivery integration
- hostile-server security
- metadata privacy
- replay resistance
- persisted provider state
- process restart recovery
- hardware-key identity
- real trust UX
- multi-device support
- revocation
- group membership policy enforcement
- Android viability
- CarbonStackOS viability
- final protocol selection

## Provider Boundary Implications

The current provider boundary likely needs to represent:

- provider identity
- public setup material / KeyPackage equivalent
- join material / Welcome equivalent
- device-local provider storage
- conversation/group state
- epoch
- member list or member summary
- provider event stream
- state-mutating message processing
- provider persistence checkpoints

The key shift:

The provider is not only an encrypt/decrypt utility. It is a stateful protocol engine.

## Trust Boundary Implications

CarbonStack should continue to own:

- user-visible trust state
- verification UX
- warnings
- blocking policy
- reverify policy
- revocation semantics
- trust history

OpenMLS/provider should own:

- cryptographic group state
- setup material
- Welcome/join processing
- message protection/opening
- provider-specific state transitions
- provider errors/events

CarbonStack must not blindly trust provider events as product policy.

## Git / Hygiene Lessons

Cargo build artifacts accidentally entered local Git history during the application-message probe.

This caused a large push attempt.

The recovery path used:

- backup branch
- reset to `origin/main`
- restore only clean source files from backup branch
- recommit only intended files

Permanent lesson:

Commit only source and lock files from the scratch crate.

Do not commit generated Cargo output.

Required ignore posture:

- `target/`
- `**/target/`
- `.fingerprint/`
- `debug/`
- `release/`

## Safe Commands

Run OpenMLS scratch directly:

cd C:\Users\udaiv\repos\carbonstack_umbrella\carbonstack-comms\internal\protocol\mls\research\openmls-minimal
cargo check
cargo run

Run local validation from `carbonstack`:

cd C:\Users\udaiv\repos\carbonstack_umbrella\carbonstack
powershell -ExecutionPolicy Bypass -File .\scripts\validate-local.ps1

Check staged files before any Rust commit:

cd C:\Users\udaiv\repos\carbonstack_umbrella\carbonstack-comms
git diff --cached --name-only

Check the latest commit before push:

git show --stat --oneline HEAD

The latest commit must not show:

- `target/`
- `.exe`
- `.pdb`
- `.o`
- `.fingerprint`

## Allowed Claims

Allowed:

- CarbonStack has a Rust-only OpenMLS scratch crate.
- The scratch crate can create a local two-member OpenMLS group.
- The scratch crate can add Bob using Bob's KeyPackage.
- The scratch crate can stage Bob from Welcome.
- The scratch crate can locally protect/open one application message.
- OpenMLS remains the first intended provider candidate.

## Not Allowed Claims

Not allowed:

- CarbonStackComms uses OpenMLS for real messaging.
- CarbonStack has production E2EE.
- CarbonStack is Signal-equivalent.
- OpenMLS has been selected as final provider.
- Cypher carries real OpenMLS traffic.
- Hostile-server security is solved.
- Replay resistance is tested.
- Metadata privacy is solved.
- Provider persistence is solved.
- Hardware-key identity is solved.
- Android/OS work is implemented.

## Next Work

Next immediate doc:

- `docs/30-openmls-provider-boundary-implications.md`

Next code experiment after docs/hygiene:

- provider-state persistence / restart simulation

The persistence/restart experiment should answer:

- Can Alice/Bob provider/group state survive beyond one in-memory run?
- What must be persisted after group creation?
- What must be persisted after Welcome join?
- What must be persisted after processing an application message?
- Does OpenMLS expose a clean storage/export shape for CarbonStack?
- Does CarbonStack need a sidecar/provider-state database layer?
