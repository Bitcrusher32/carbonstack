# USB Transfer Model

USB data should be off by default.

CarbonStackOS should not expose a normal MTP-style filesystem workflow.

The host should never receive general browsing access to the device filesystem.

## Baseline USB States

### Device Locked

USB data fully disabled.

Allowed:

- charging only, where hardware permits

Blocked:

- MTP
- ADB
- mass storage
- file transfer
- tethering
- accessory protocols
- debugging

### Device Unlocked

USB remains charge-only by default.

USB data does not become available merely because the user unlocked the device.

### Transfer Mode

Transfer mode requires explicit local approval.

Transfer mode may require hardware-key confirmation.

Transfer mode should use a narrow custom protocol.

## Preferred Transfer Mode Properties

- explicit local approval
- optional hardware-key confirmation
- short-lived session
- visible peer identity where possible
- no general filesystem browsing
- no MTP
- no ADB
- no mass storage
- no vendor accessory protocols
- no silent background transfer

## Inbound Transfer Flow

1. host requests transfer
2. user approves transfer mode locally
3. optional hardware-key confirmation
4. host sends file through narrow protocol
5. file enters quarantine
6. file is validated by content, not extension
7. metadata is stripped where applicable
8. file is rewritten into canonical internal form
9. accepted file moves into approved local library
10. rejected file is deleted or retained only in quarantine diagnostics

## Export Flow

Export must be explicit.

Possible export types:

- selected notes export
- selected audio export
- encrypted backup bundle
- signed manifest
- public diagnostics bundle

Export should not provide broad filesystem access.

## Recovery Mode USB

During recovery, USB should remain restricted.

Allowed:

- USB HID
- FIDO2/security-key protocol

Blocked:

- MTP
- ADB
- mass storage
- MIDI
- tethering
- vendor accessory protocols
- file transfer
- debugging
- general USB data

## Principle

USB is a local physical attack surface.

It should never become a general convenience pathway by default.
