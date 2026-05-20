# Updates

CarbonStackOS updates must be signed, verified, and rollback-protected.

The update system must not depend on a general browser or broad web stack.

## Required Properties

OS updates should be:

- signed
- verified before installation
- rollback-protected
- delivered through a minimal trusted updater
- auditable
- documented
- recoverable where possible

## Disallowed Update Sources

The device should not accept:

- unsigned OS images
- older vulnerable rollback images
- user-modified system partitions
- untrusted update channels
- arbitrary downloaded installers
- browser-mediated update flows
- third-party app-store updates

## Release Artifacts

For project credibility, releases should aim to provide:

- signed release manifests
- published hashes
- clear version numbers
- security changelogs
- independent build instructions
- reproducible or independently buildable releases where practical

## Update Channel

The update channel should be:

- narrow
- signed
- explicit
- independent of browser engines
- independent of general-purpose app stores

## Malicious Update Risk

A malicious signed update remains a catastrophic risk.

CarbonStack must treat signing infrastructure as high-value security infrastructure.

Recommended controls:

- offline signing where practical
- hardware-backed signing keys where practical
- documented key ceremony
- signed release notes
- public verification instructions
- emergency revocation process
- build environment isolation

## Rollback

Rollback to vulnerable builds should be blocked where platform support allows.

If rollback is required for recovery, it should be:

- explicit
- documented
- visible
- limited
- treated as a security event

## Principle

An update mechanism is remote code execution by design.

It must be narrower, more auditable, and more controlled than ordinary software installation.
