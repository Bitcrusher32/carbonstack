# Recovery, Revocation, and Device Loss

CarbonStack must distinguish between locking, wiping, and revoking.

These are separate operations.

## Lock

Locking means:

- local vault keys are evicted from memory
- secure state becomes inaccessible
- data remains recoverable with the correct recovery ceremony

Locking is appropriate for:

- normal session timeout
- lockdown mode
- temporary loss of confidence
- suspicious state change

## Wipe

Wiping means:

- local encrypted data is deleted
- local vault material may be destroyed
- recovery may or may not be possible depending on backup policy

Wiping is appropriate for:

- user-requested reset
- device disposal
- severe compromise response
- optional configured duress behavior

## Revoke

Revocation means:

- other users or groups are told that a device identity is no longer trusted

Revocation is a network-visible trust action.

A seized offline device can lock or wipe itself.

It cannot notify others until some trusted channel reaches the network.

## Device Loss

When a device is lost:

1. mark local device as lost where possible
2. use trusted recovery device or hardware key
3. generate signed revocation
4. propagate revocation to CarbonStackCypher
5. notify groups
6. require re-verification for replacement device

## Device Replacement

Device replacement must be visible.

The system must not allow silent replacement of identity keys.

Replacement should require:

- hardware-key approval
- recovery passphrase where configured
- visible contact warning
- group-visible device change
- explicit re-verification where policy requires it

## Recovery After Duress

Recovery should be staged.

Suggested flow:

1. recovery PIN or passphrase
2. optional biometric local confirmation
3. hardware-key challenge-response
4. vault unsealing or rewrapping
5. trust-state review
6. normal state restoration

## Burn Mode

Burn mode means the local device is no longer recoverable locally.

Burn mode should:

- destroy local vault keys
- destroy local message database access
- mark device identity locally compromised
- require re-enrollment

Burn mode may later send signed revocation if network returns.

Burn mode should be optional and dangerous.

## Principle

Recovery restores local access.

Revocation restores group trust.

They are not the same thing.
