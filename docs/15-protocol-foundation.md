# Protocol Foundation

CarbonStack MUST NOT invent cryptography casually.

CarbonStack defines an application architecture, trust model, hostile-server model, text policy, device policy, and appliance posture around mature cryptographic protocol foundations.

The project should avoid designing new ratchets, group key agreement schemes, or encryption constructions unless no mature foundation satisfies the requirement and the resulting design receives serious review.

## Phase 1 Foundation

The Phase 1 messaging target is one-to-one asynchronous messaging.

CarbonStackComms SHOULD investigate a Signal-style foundation for this phase:

- asynchronous session setup
- prekey-based delivery support
- Double Ratchet-style message key evolution
- forward secrecy
- post-compromise recovery properties
- replay resistance
- loud key-change behavior

Candidate foundations:

- Signal Protocol / libsignal
- X3DH or PQXDH-style initial key agreement
- Double Ratchet-style session encryption

The exact implementation remains future work.

The project MUST evaluate:

- licensing
- language bindings
- Android viability
- testability
- auditability
- build reproducibility
- hardware-key identity integration
- compatibility with strict CarbonStack text and envelope rules

## Phase 2+ Group Foundation

Future group messaging SHOULD investigate MLS as the preferred group foundation.

MLS is a better conceptual fit for auditable group epochs, membership changes, and scalable group keying than inventing a custom group protocol.

CarbonStack group work SHOULD preserve:

- explicit membership changes
- auditable group epochs
- visible device additions
- visible device removals
- revocation propagation
- no silent server-side group mutation
- no silent key replacement
- no server-authored group truth

## Noise Framework

Noise MAY be investigated for:

- narrow custom transport handshakes
- local device-to-device channels
- appliance-to-relay bootstrap channels
- constrained file transfer channels

Noise SHOULD NOT be treated as a complete messaging protocol by itself.

CarbonStack would still need to define:

- identity
- storage
- replay behavior
- group state
- trust UX
- revocation
- recovery
- hostile-server behavior

## CarbonStack-Owned Layer

Even when using mature crypto foundations, CarbonStack still owns:

- strict text normalization before encryption
- message envelope format
- server storage model
- device identity model
- hardware-key enrollment ceremony
- QR verification ceremony
- safety-number and key-change UX
- local vault lifecycle
- group membership visibility
- revocation propagation
- hostile-server failure behavior
- appliance OS integration

## Requirement Language

MUST:

- use mature reviewed cryptographic foundations where available
- document protocol choices
- document limitations
- include test vectors where possible
- distinguish confidentiality, authenticity, metadata, and availability claims

MUST NOT:

- claim Signal-equivalent security without review
- invent custom crypto casually
- hide trust changes
- allow the server to silently replace identities
- allow the server to silently add group members
- treat TLS as a replacement for E2EE

SHOULD:

- start with 1:1 messaging
- keep group-aware schema concepts
- preserve future MLS compatibility
- keep cryptographic code isolated and testable
- use hostile-server test harnesses

MAY:

- use Noise for narrow transport or transfer cases
- support multiple protocol backends in experimental builds
- add MLS groups after the 1:1 model is stable

## Core Principle

CarbonStack should innovate around constrained systems, hostile-server deployment, trust visibility, parser minimization, and appliance posture.

It should not innovate by casually inventing cryptography.
