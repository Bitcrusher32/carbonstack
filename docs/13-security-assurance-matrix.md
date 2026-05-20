# Security Assurance Matrix

The Security Assurance Matrix tracks how CarbonStack claims are tested, limited, and validated.

CarbonStack should avoid vague security claims.

Every major claim should map to:

- intended protection
- mechanism
- test method
- failure mode
- known limitation
- project phase

## Assurance Categories

### Protocol Assurance

Questions:

- Can the server read plaintext?
- Can the server silently replace identity keys?
- Can the server silently add group members?
- Can messages be replayed?
- Can group state be rolled back?
- Are key changes visible?

Tests:

- hostile-server simulation
- replay tests
- rollback tests
- identity-replacement tests
- group-membership mutation tests
- client warning tests

### Parser Assurance

Questions:

- Are invalid byte sequences rejected?
- Are dangerous Unicode controls blocked or marked?
- Are unsupported files rejected?
- Are accepted files canonicalized?
- Can parser compromise reach message keys?

Tests:

- Unicode rejection corpus
- malformed TXT/MD corpus
- malformed WAV/FLAC corpus
- quarantine validation tests
- compartment permission tests

### Device Assurance

Questions:

- Is boot state verified?
- Is the bootloader locked?
- Is the OS signed?
- Is rollback blocked?
- Is the app running on an approved device profile?
- Are unsafe device states refused?

Tests:

- unlocked bootloader test
- stale patch test
- developer-options test
- USB-debugging test
- root-detection test
- app-signature mismatch test
- device-profile mismatch test

### Local Vault Assurance

Questions:

- Is the vault cryptographically locked?
- Are keys evicted on lock?
- Are keys evicted on duress?
- Can notes or media compartments access message keys?
- Can backups expose plaintext?

Tests:

- session timeout tests
- lockdown tests
- memory/key lifecycle tests
- compartment access tests
- backup export tests

### Server Assurance

Questions:

- Does CarbonStackCypher store only encrypted envelopes?
- Are admin actions logged?
- Can admin access plaintext?
- Are registrations invite-only?
- Can revocation propagate?
- Can the server forge trust events?

Tests:

- database inspection tests
- admin permission tests
- audit log tests
- invite-flow tests
- revocation-flow tests
- malicious relay tests

### OS Appliance Assurance

Questions:

- Is the system free of browser engines?
- Is WebView absent or unusable by normal apps?
- Is APK sideloading blocked?
- Is ADB disabled in production?
- Is USB data disabled by default?
- Is Bluetooth restricted?
- Are updates signed?

Tests:

- installed package audit
- forbidden component audit
- USB mode tests
- Bluetooth profile tests
- update signature tests
- rollback tests
- production-build policy tests

## Example Matrix Row Format

| Claim | Mechanism | Test | Failure Mode | Phase |
|---|---|---|---|---|
| Server cannot read messages | client-side E2EE | inspect database and relay logs | client leaks plaintext before encryption | Phase 1 |
| Key changes are loud | trust-state UI | forced key replacement test | user ignores warning | Phase 1 |
| FLAC parser cannot access message vault | compartment separation | permission and IPC test | OS sandbox escape | Phase 5 |
| Lockdown evicts vault keys | privileged state machine | trigger lockdown and attempt vault read | compromised kernel | Phase 4 |

## Principle

A security claim without a test is a wish.

A security claim without a limitation is marketing.

CarbonStack should be engineering, not marketing.
