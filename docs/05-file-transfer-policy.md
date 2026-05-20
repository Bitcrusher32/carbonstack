# File Transfer Policy

CarbonStack treats inbound files as hostile.

File transfer is not a default convenience feature. It is a constrained, explicit, quarantine-first operation.

## Allowed Initial File Types

- TXT
- MD
- WAV
- FLAC

Extensions are not trusted. Files must be validated by content and parser rules.

## Inbound Flow

1. receive file into quarantine
2. validate content
3. reject ambiguous or malformed files
4. strip metadata
5. rewrite into canonical internal form
6. move accepted file into approved local library
7. deny direct open-from-transfer

## Rejected Initially

- ZIP
- PDF
- DOCX
- HTML
- SVG
- JPEG
- PNG
- MP4
- MOV
- M4A
- MP3
- archives
- executables
- scripts
- unknown containers
- playlist imports

## Principle

A file parser compromise should not become a message compromise.
