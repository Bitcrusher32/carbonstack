# Stack Architecture

CarbonStack is divided into three main implementation projects and one shared specification layer.

## carbonstack

Base documentation, protocol notes, threat model, roadmap, and security policy.

## carbonstack-comms

Messaging client.

Responsibilities:

- user identity
- local secure vault
- text normalization
- message encryption and decryption
- trust state display
- QR or hardware-key verification
- key-change warnings
- duress-aware behavior where supported by CarbonStackOS

## carbonstack-cypher

Server and relay.

Responsibilities:

- encrypted envelope storage
- routing
- rate limiting
- group delivery support
- revocation propagation
- small-community self-hosting

Non-responsibilities:

- plaintext access
- trusted identity replacement
- trusted group membership mutation
- trusted message history authority

## carbonstack-os

Appliance OS.

Responsibilities:

- constrained runtime
- immutable base posture
- signed updates
- interface gating
- parser compartmentalization
- duress state machine
- device assurance enforcement
