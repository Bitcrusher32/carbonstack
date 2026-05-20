# Bluetooth Model

Bluetooth is a convenience cost, not a free feature.

CarbonStackOS should treat Bluetooth as a constrained, optional interface.

## Baseline

- Bluetooth disabled by default
- manual pairing only
- no always-discoverable mode
- no automatic convenience pairing
- no broad nearby-device ecosystem

## Allowed Initial Use

Optional:

- A2DP audio output only

Preferred:

- wired audio for high-security use

## Disallowed

- Fast Pair
- Nearby Share
- BLE convenience ecosystem
- file transfer profiles
- contact sharing
- message access profile
- smartwatch integration
- companion-device integration
- automatic device discovery
- background pairing prompts
- Bluetooth metadata sync beyond bare minimum

## System Modes

### Normal Appliance Mode

Bluetooth audio may be enabled if user permits it.

Bluetooth should remain narrow and explicit.

### High-Security Mode

Bluetooth fully disabled.

### Duress / Lockdown Mode

Bluetooth forcibly disabled.

Only the privileged state machine may restore it.

## Pairing Policy

Pairing should be:

- manual
- visible
- logged
- revocable
- blocked during lockdown
- blocked during recovery unless specifically required

A new Bluetooth pairing may be treated as a critical state change by CarbonStackComms or CarbonStackOS.

## Principle

Bluetooth should not become a hidden file-transfer, contact-sharing, metadata-sharing, or companion-device surface.
