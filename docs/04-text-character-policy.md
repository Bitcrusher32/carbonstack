# Text and Character Policy

Text-only does not mean automatically safe.

CarbonStack text handling should be boring, visible, normalized, and unsurprising.

## Baseline Rules

- UTF-8 only
- reject invalid byte sequences
- normalize to NFC
- reject or visibly mark control characters
- reject bidi override/control characters
- reject zero-width characters
- reject private-use characters
- reject unassigned codepoints
- reject dangerous combining sequences
- limit maximum message size
- limit maximum line length

## Rendering Rules

- bundled fonts only
- no downloadable fonts
- no HTML
- no CSS
- no markdown in chat by default
- no clickable links by default
- no automatic URL detection
- no embedded previews

## Preferred Failure Mode

Unexpected characters should render visibly, for example:

- [U+202E blocked]
- [U+200B blocked]
- [unsupported character]
