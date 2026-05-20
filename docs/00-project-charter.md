# CarbonStack Project Charter

CarbonStack is a secure communications appliance stack designed for small trusted groups.

It is not a general-purpose smartphone platform. It is not a consumer convenience platform. It is not a Signal clone.

CarbonStack exists to reduce endpoint, server, parser, and identity attack surface by combining:

1. a constrained appliance OS,
2. a text-first E2EE client,
3. a hostile-server relay,
4. strict parser and file policies,
5. explicit hardware-backed trust ceremonies.

## Stack Components

### CarbonStackOS

The hardened Android-derived appliance operating system.

Target reference device:

- Pixel 9a
- AOSP / GrapheneOS-derived base
- locked bootloader
- verified boot
- signed OS images only
- no cellular dependency
- no general-purpose app ecosystem

### CarbonStackComms

The deployable text-first messaging client.

Properties:

- text only by default
- no rich previews
- no inline attachments
- no browser rendering
- no hidden linkification
- loud trust changes
- hardware-key-backed identity where possible

### CarbonStackCypher

The deployable relay/server stack.

Properties:

- stores encrypted envelopes
- routes messages
- enforces basic abuse controls
- supports small private groups
- cannot read plaintext
- cannot silently replace identities
- cannot silently add group members

## Core Doctrine

Every feature is guilty until it proves it does not add unacceptable parser, network, sensor, identity, or filesystem authority.

## Non-Claim

CarbonStack does not promise perfect security.

Encrypted messaging cannot protect plaintext once it is displayed on a compromised endpoint.

CarbonStack reduces endpoint compromise likelihood by refusing to be a general-purpose endpoint.
