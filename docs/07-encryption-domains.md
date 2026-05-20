# Encryption Domains

CarbonStackOS should separate encrypted state into distinct domains.

The system should avoid one global user data bucket where unrelated components can access each other's sensitive state.

## Suggested Domains

### OS Base

Purpose:

- verified operating system base
- read-only from ordinary user perspective
- signed
- rollback-protected
- immutable in production use

Contains:

- system components
- privileged state machine
- bundled apps
- update verifier
- fixed policy configuration

### User Profile

Purpose:

- local non-sensitive user settings
- optional notes
- optional approved local media
- device preferences

Contains:

- display settings
- local UI configuration
- non-sensitive notes if configured
- approved music library if configured

### Secure Messaging Vault

Purpose:

- protect CarbonStackComms state

Contains:

- message database
- identity keys
- group state
- contact trust records
- revocation state
- safety-number history
- local delivery state

This vault must be cryptographically locked.

It must not merely be hidden by UI.

### Quarantine / Import Vault

Purpose:

- hold inbound files before validation

Contains:

- raw inbound files
- temporary parser output
- validation logs
- canonicalized output awaiting approval

This vault must not have access to message keys.

### Recovery Vault

Purpose:

- store minimum state required for recovery ceremony

Contains:

- recovery policy
- wrapped recovery metadata
- hardware-key enrollment references
- duress state recovery material where configured

This vault should not contain message plaintext.

## Lockdown Behavior

Lockdown should:

- evict sensitive keys from memory
- unmount or cryptographically lock vaults where possible
- deny app access to secure domains
- disable transfer endpoints
- disable media indexing
- make quarantine inaccessible

Lockdown should not simply hide files.

## Domain Separation Principle

A compromise in one compartment should not automatically expose unrelated domains.

Examples:

- a FLAC parser bug should not expose message keys
- a Markdown parser bug should not expose contact trust records
- a transfer-service bug should not expose the secure messaging vault
- a notes bug should not enable OS update tampering
