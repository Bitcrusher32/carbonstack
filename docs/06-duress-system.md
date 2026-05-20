# Duress System

CarbonStackOS includes configurable duress behavior.

A duress PIN does not "encrypt the device." The device should already be encrypted.

A duress PIN changes system state.

It can:

- evict keys
- lock vaults
- disable interfaces
- suppress notifications
- hide sensitive state
- require recovery ceremony
- optionally destroy local vault access

## Core Principle

The duress system must be implemented as a small privileged state machine, not scattered application behavior.

No ordinary app should independently restore:

- radios
- sensors
- vault access
- transfer mode
- notification visibility
- messaging sync

## Separate Secrets

### Normal PIN

Unlocks the normal profile.

### Duress PIN

Triggers configured duress behavior.

### Recovery PIN or Passphrase

Begins recovery from duress state.

### Hardware Key

Acts as a possession factor for:

- enrollment
- recovery
- revocation
- high-risk identity changes

### Biometric

Optional local intent check.

Biometrics are not root identity. They should not replace recovery passphrases or hardware keys.

## Duress Modes

### Decoy Mode

Purpose:

- someone is looking over the user's shoulder
- the device should not obviously panic

Behavior:

- opens a limited decoy profile
- hides secure conversations
- hides contact names
- suppresses sensitive notifications
- preserves harmless notes or music if configured
- does not reveal that duress mode is active

### Lockdown Mode

Purpose:

- user may lose physical control of the device

Behavior:

- locks secure vault
- evicts vault keys from memory
- stops messaging
- disables Wi-Fi
- disables Bluetooth
- disables USB data
- disables camera
- disables microphone
- disables sensors where platform allows
- hides notifications
- requires recovery ceremony

### Burn Mode

Purpose:

- device should no longer be recoverable locally

Behavior:

- destroys local vault keys
- destroys local message database access
- marks device identity as locally compromised
- requires re-enrollment
- can later send signed revocation if network returns

Burn mode is optional and dangerous.

Burn mode should never be the default duress action.

## Recovery Flow

Recovery should be staged and narrow.

Suggested flow:

1. duress state active
2. enter recovery PIN or passphrase
3. optional biometric local confirmation
4. temporary recovery window opens
5. USB-HID/FIDO2-only path enabled
6. hardware key challenge-response
7. secure vault keys are unsealed or rewrapped
8. system restores configured normal state

## Recovery USB Policy

Allowed during recovery window:

- USB HID
- FIDO2/security-key protocol

Blocked during recovery window:

- MTP
- ADB
- mass storage
- MIDI
- tethering
- vendor accessory protocols
- file transfer
- debugging
- general USB data

## Recovery Window

Suggested behavior:

- short timeout, for example 60 to 120 seconds
- failed hardware-key attempt closes the window
- retry requires recovery PIN or passphrase again
- increasing delay after failures

## Configurable Setting

Allow local recovery after duress:

- ON
- OFF

Example policy:

- Decoy PIN: local recovery allowed
- Lockdown PIN: local recovery allowed or disabled by user
- Burn PIN: local recovery impossible by definition
