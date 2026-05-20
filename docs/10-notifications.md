# Notifications

CarbonStack notifications should avoid leaking sensitive content.

Notifications are an attack surface and a privacy surface.

## Default Policy

Default notifications should not expose:

- message body
- sender names on lock screen
- group names on lock screen
- message previews
- attachments
- images
- action buttons that expose state

## Allowed Minimal Notification

A safe default notification may say:

- New message
- CarbonStackComms activity
- Secure message pending

No plaintext content should appear.

## Disallowed

- rich previews
- inline replies from lock screen
- sender avatar display
- message body preview
- image preview
- link preview
- notification action buttons that mark read or reply
- hidden linkification
- notification-carried decrypted payloads

## Lock Screen Behavior

On lock screen:

- no message body
- no sender names unless explicitly configured
- no group names unless explicitly configured
- no conversation previews
- no sensitive counts if configured high-security

## Decoy Mode

In decoy mode:

- sensitive notifications suppressed
- secure contacts hidden
- secure conversations hidden
- notification behavior should not reveal duress state

## Lockdown Mode

In lockdown mode:

- messaging sync stopped or hidden
- sensitive notifications suppressed
- no obvious "duress active" banner
- no state-revealing prompts

## Principle

Notifications should announce activity without becoming a side channel for message content, identity state, or duress state.
