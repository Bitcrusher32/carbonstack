# Current status notice

This document is a historical canonical specification draft.

It preserves project doctrine and long-range design intent. It should not be treated as the current release surface by itself.

For current public behavior and known-good validation, start with:

- `README.md`
- `docs/README.md`
- `docs/113-experimental-backbone-deployability-runbook-v0.md`

As of the v0.2.61/v0.2.62 public-surface cleanup rung, the current validated artifact is the experimental Cypher + Comms OpenMLS relay backbone. This document may contain older assumptions, future targets, and long-range design language.

---
# CarbonStack Full Specification

Status: Draft Canonical Specification  
Repository: `carbonstack`  
Applies to: CarbonStackOS, CarbonStackComms, CarbonStackCypher  
Initial appliance testbed: Pixel 4a on-hand  
Reference target: Pixel 9a  
Spec intent: preserve current project doctrine for human-in-loop engineering and future implementation work

---

## 0. Specification Authority

This document is the canonical project intent document for CarbonStack.

More specific repository documents may refine implementation details, but they SHOULD NOT contradict this specification without an explicit revision to this document.

Authoritative project naming:

- **CarbonStackOS**: the constrained Android-derived appliance operating system.
- **CarbonStackComms**: the text-first encrypted messaging client.
- **CarbonStackCypher**: the self-hostable hostile-server relay and storage stack.
- **CarbonStack**: the overall secure communications appliance stack and shared doctrine.

Repository authority model:

- `carbonstack` is the source of truth for shared doctrine, threat model, requirements language, protocol notes, text policy, file policy, security assurance, and roadmap.
- `carbonstack-comms` implements the messaging client.
- `carbonstack-cypher` implements the server/relay stack.
- `carbonstack-os` implements or documents the appliance OS target.

If an implementation repo contradicts the canonical specification, the canonical specification wins until the spec is revised.

---

## 1. Requirement Language

CarbonStack uses explicit requirement language to separate current obligations, strong preferences, optional features, future work, and rejected scope.

- **MUST**: required for the relevant phase, component, or claimed compliance level.
- **MUST NOT**: explicitly forbidden for the relevant phase, component, or claimed compliance level.
- **SHOULD**: strongly preferred; deviation requires technical justification.
- **SHOULD NOT**: strongly discouraged; deviation requires technical justification.
- **MAY**: optional.
- **DEFERRED**: valid future goal, not required for the current phase.
- **REJECTED**: intentionally out of scope for this iteration or profile.

Requirement levels are scoped. A feature MAY be rejected for CarbonStackOS MVP while being allowed in forks. A feature MAY be optional for the normal Android development app while required for future high-assurance appliance builds.

---

## 2. North Star

CarbonStackOS is a deliberately constrained, security-first Android-derived operating system for a single-purpose communications appliance.

Its goal is not to be a better smartphone.

Its goal is to stop being a smartphone.

CarbonStack exists to build open-source maximalist secure communications that can actually be deployed, understood, audited, and operated by technically competent users without hidden institutional knowledge.

The system prioritizes:

- boring text
- loud trust changes
- hostile-server assumptions
- disposable parsers
- immutable base
- minimal ambient attack surface
- explicit identity ceremonies
- small trusted groups
- refusal of convenience features that silently expand authority

CarbonStack is not a mass-market messaging app. It is not a general-purpose smartphone platform. It is not a browser platform. It is a secure communications appliance stack.

---

## 3. Core Doctrine

The central CarbonStack rule is:

> Every feature is guilty until it proves it does not add unacceptable parser, network, sensor, identity, or filesystem authority.

CarbonStack assumes:

> The server may be hostile.  
> The network may be hostile.  
> Inbound files may be hostile.  
> Other users' devices may become hostile.  
> The OS must minimize what any one compromised component can reach.

CarbonStack also accepts the core endpoint truth:

> Encrypted messaging cannot protect plaintext once it is displayed on a compromised endpoint.

CarbonStack therefore does not claim perfect security. It reduces attack surface by refusing to be a general-purpose endpoint.

A security claim without a test is a wish.

A security claim without a limitation is marketing.

---

## 4. Project Goals

CarbonStack's long-term goal is a deployable open-source secure communications stack with a path from normal Android experimentation to hardened appliance use.

The project aims to provide:

1. a canonical security doctrine and specification;
2. a text-first encrypted messaging client;
3. a self-hostable hostile-server relay;
4. strict text and parser policy;
5. hardware-key-backed identity and recovery for high-assurance profiles;
6. loud trust-state changes;
7. a constrained Android-derived appliance OS profile;
8. repeatable deployment instructions;
9. clear security limitations;
10. human-readable documentation.

CarbonStack values understandable operation. Security that only exists in scattered expert lore is not good enough for the project goal.

Documentation is part of the security model.

---

## 5. Stack Components

### 5.1 CarbonStackOS

CarbonStackOS is the hardened Android-derived appliance operating system.

It is responsible for:

- immutable or immutable-like base posture;
- locked bootloader and verified boot where supported;
- signed OS images and signed updates;
- rollback resistance where supported;
- removal of general smartphone features;
- no general app ecosystem;
- no browser platform in the appliance profile;
- no normal package installation workflow;
- interface gating;
- duress state management;
- parser compartmentalization;
- device assurance enforcement;
- local vault and domain separation support.

CarbonStackOS is the final high-assurance goalpost, not the first implementation step.

### 5.2 CarbonStackComms

CarbonStackComms is the text-first encrypted messaging client.

It is responsible for:

- user identity;
- local secure vault;
- message encryption and decryption;
- strict text validation and rendering;
- QR or hardware-key verification;
- trust-state display;
- key-change warnings;
- group state display;
- local session locking;
- high-assurance hardware-key workflows where supported.

CarbonStackComms SHOULD first be developed as a normal Android-compatible app, then hardened for CarbonStackOS.

### 5.3 CarbonStackCypher

CarbonStackCypher is the self-hostable relay/server stack.

It is responsible for:

- encrypted envelope storage;
- routing;
- basic rate limiting;
- group delivery support;
- device revocation propagation;
- invite-only or private community operation;
- admin audit logging;
- deployment tooling.

CarbonStackCypher is not trusted.

It MUST NOT receive message plaintext.

It MUST NOT store user private keys.

It MUST NOT be required to be honest for message confidentiality.

---

## 6. MVP Tiers

CarbonStack uses staged MVP tiers.

### 6.1 MVP-0: Specification MVP

MVP-0 is the documentation and doctrine baseline.

Required artifacts:

- canonical full specification;
- threat model;
- non-goals;
- stack architecture;
- text and character policy;
- file transfer policy;
- encryption-domain model;
- duress-system model;
- recovery/revocation model;
- security assurance matrix;
- roadmap.

MVP-0 is complete when the project intent is preserved well enough that future human-in-loop development can continue without rediscovering the philosophy from chat logs.

### 6.2 MVP-1: Software Stack MVP

MVP-1 is the initial deployable software stack:

- CarbonStackComms as a normal Android-compatible application;
- CarbonStackCypher as a self-hostable relay/server;
- one-device-per-user initial model;
- text-only messaging;
- strict text policy;
- no rich previews;
- no inline attachments;
- no browser rendering;
- no plaintext server access;
- invite-only registration;
- QR or explicit verification workflow;
- loud key-change warnings;
- encrypted local vault;
- no plaintext notification content.

Hardware-key support is not a hard requirement for the most basic MVP-1 release, but the architecture MUST NOT preclude it.

### 6.3 MVP-2: Appliance OS MVP

MVP-2 is the CarbonStackOS appliance prototype.

Target strategy:

- Pixel 4a is the practical on-hand development testbed.
- Pixel 9a is the reference target for future appliance planning.
- The reference target MAY change if bootloader, verified boot, signing, update support, hardware-backed keystore, maintenance, or availability make it unsuitable.

MVP-2 requirements:

- Android-derived base;
- AOSP foundation;
- informed by GrapheneOS-derived hardening concepts;
- locked bootloader where supported;
- verified boot where supported;
- signed OS images;
- no general app ecosystem;
- no browser;
- no Play Services dependency in appliance mode;
- no arbitrary APK installation;
- no ADB in production;
- CarbonStackComms as primary application;
- Wi-Fi for approved functions;
- cellular rejected for this iteration;
- USB data disabled by default;
- Bluetooth disabled by default or restricted;
- initial privileged duress state machine.

### 6.4 MVP-3: Constrained Utility Layer

MVP-3 adds optional local utility features only after the core secure comms stack is working.

Potential additions:

- TXT notes;
- constrained Markdown notes;
- WAV playback;
- FLAC playback;
- internal canonical playlist format;
- signed manifests;
- custom USB transfer mode;
- optional restricted Bluetooth audio.

These features MUST remain subordinate to the communications-appliance goal.

---

## 7. Platform Base

CarbonStackOS SHOULD be based on:

- AOSP foundation;
- GrapheneOS-derived hardening concepts where feasible;
- device-specific support;
- locked bootloader where possible;
- verified boot where possible;
- rollback protection where possible;
- signed OS images only;
- reproducible or independently buildable releases where practical.

CarbonStackOS SHOULD NOT claim to be GrapheneOS unless it is actually a proper GrapheneOS fork and follows licensing, attribution, and build realities.

Preferred wording:

> CarbonStackOS is based on AOSP and informed by GrapheneOS-derived hardening concepts where compatible with licensing, build feasibility, device support, and project scope.

The base OS SHOULD be immutable from the ordinary user's perspective.

No normal user workflow should allow:

- APK sideloading;
- package installation;
- developer options;
- ADB in production;
- dynamic code loading;
- JIT where avoidable;
- user-installed fonts;
- user-installed keyboards;
- plugin systems;
- browser engines;
- general URI handling.

A system compromise should require either a platform vulnerability or a signed malicious OS image, not ordinary user interaction.

---

## 8. Device Target Policy

CarbonStackOS uses device targets, not universal Android promises.

Initial practical testbed:

- **Pixel 4a**, because it is available on-hand.

Reference target:

- **Pixel 9a**, for future appliance planning.

The reference target is not permanent. It MAY change if the device fails practical security or maintainability requirements.

Target-device selection criteria:

- bootloader behavior;
- verified boot behavior;
- ability to build or install controlled images;
- security patch lifecycle;
- hardware-backed keystore support;
- StrongBox support if available;
- Wi-Fi stability;
- USB behavior;
- YubiKey/FIDO2 path;
- camera/sensor isolation if used later;
- community documentation;
- replacement availability;
- cost.

Device assurance tiers:

- **Tier A**: tested CarbonStackOS target with documented build, boot, update, and policy behavior.
- **Tier B**: normal Android device that can run CarbonStackComms with partial device checks.
- **Tier C**: unsupported or low-assurance device.
- **Rejected for high assurance**: rooted devices, unlocked bootloader devices, stale-patch devices, unknown OS builds, untrusted app signatures.

CarbonStackComms MAY run on broader Android devices in low-assurance mode. CarbonStackOS MUST remain narrow and target-specific.

---

## 9. Connectivity Model

CarbonStackOS is not a normal phone.

Baseline connectivity for this iteration:

- no cellular;
- Wi-Fi only for approved system functions;
- Bluetooth disabled by default or restricted;
- USB data disabled by default;
- no general LAN discovery;
- no general web access;
- no browser;
- no arbitrary internet capability.

Allowed network functions:

- CarbonStackComms messaging;
- CarbonStackCypher server communication;
- strict file-transfer mode, if enabled;
- OS update checking and downloading through signed update channel.

Disallowed:

- web browsing;
- ad tech;
- tracking telemetry;
- cloud assistant services;
- general app internet access;
- automatic link opening;
- automatic link previews;
- nearby-device convenience ecosystems.

### 9.1 Cellular Policy

Cellular is rejected entirely for this iteration.

CarbonStackOS MUST NOT depend on:

- phone numbers;
- SMS;
- carrier identity;
- cellular modem connectivity;
- calling;
- mobile carrier services.

Initial appliance builds SHOULD disable or omit cellular functionality.

Forks may explore cellular if they want, but this canonical iteration rejects it. Cellular support would be a major attack-surface expansion and a philosophical departure from the initial appliance profile.

---

## 10. CarbonStackComms Overview

CarbonStackComms is the deployable text-first messaging client.

Its purpose:

- private group messaging;
- self-hostable deployment;
- hardware-key-backed identity in future high-assurance mode;
- hostile-server encryption model;
- minimal parser exposure;
- no rich-message attack surface.

CarbonStackComms should feel IRC-adjacent: simple, boring, direct text communication with modern encrypted messaging properties.

Default properties:

- text only;
- no inline attachments;
- no rich previews;
- no stickers;
- no GIFs;
- no voice messages;
- no arbitrary Unicode rendering;
- no browser-based rendering;
- no hidden linkification;
- no server-trusted identity changes.

### 10.1 Normal Android App Mode

Normal Android app mode exists for early development and broader testing.

It MAY support:

- ordinary Android installation;
- basic local vault encryption;
- QR verification;
- strict text policy;
- CarbonStackCypher connection;
- low-assurance device-state warnings.

It MUST NOT claim full appliance security.

### 10.2 High-Assurance Appliance Mode

High-assurance appliance mode is the future target.

It SHOULD require:

- approved device profile;
- locked bootloader;
- verified OS state;
- hardware-backed local key storage;
- hardware-key-backed identity enrollment and recovery;
- no unapproved apps;
- no browser;
- no general file manager;
- no third-party keyboards;
- no third-party accessibility services;
- CarbonStackOS policy enforcement.

Hardware keys are required for full future high-assurance releases.

---

## 11. CarbonStackCypher Overview

CarbonStackCypher is the deployable server stack.

Its job is routing and storage, not trust.

Server responsibilities:

- store encrypted envelopes;
- route messages;
- enforce basic rate limits;
- support group delivery;
- support device revocation propagation;
- support self-hosting;
- support small private communities first.

The server MUST NOT be able to:

- read plaintext;
- silently add group members;
- silently replace client keys;
- forge sender identity;
- rewrite group history undetectably;
- silently roll back group state.

Desirable future properties:

- append-only membership logs;
- auditable group epochs;
- hardware-key-signed device enrollment;
- hardware-key-signed device revocation;
- private invite flows;
- QR or hardware-key ceremony for contact verification.

CarbonStackCypher should be treated as compromised by design.

---

## 12. Protocol Commitment

CarbonStackComms MUST NOT invent cryptography casually.

The exact protocol remains future work.

Candidate protocol foundations to investigate:

- Signal Protocol / libsignal;
- MLS, the Messaging Layer Security standard;
- Noise framework for lower-level channel construction;
- existing audited libraries that match project requirements.

Protocol selection criteria:

- hostile-server support;
- auditable group membership;
- append-only trust history;
- replay resistance;
- rollback resistance;
- clear compromise recovery;
- device revocation;
- hardware-key identity workflows;
- licensing compatibility;
- testability;
- implementation maturity;
- mobile viability.

Until a protocol is selected, CarbonStack documentation MUST distinguish doctrine from implementation.

The north star is:

- hostile server;
- auditable group membership;
- append-only trust history;
- replay resistance;
- clear compromise recovery;
- loud safety-number or key-change UX.

---

## 13. Identity and Trust Model

CarbonStack identity MUST be user-visible and resistant to silent server manipulation.

Desired identity model:

- each client instance has a device identity;
- device identity is locally generated;
- private identity material never leaves the device unprotected;
- high-assurance releases require hardware-key-backed enrollment;
- group membership changes are explicit;
- key changes are loud;
- device replacement requires visible verification;
- server cannot silently add members;
- server cannot silently replace keys.

Trust records SHOULD be append-only or append-only-like from the user perspective.

CarbonStackComms SHOULD maintain:

- contact trust records;
- safety-number history;
- key-change history;
- group membership history;
- device enrollment history;
- revocation history.

The server MAY assist with distribution, but MUST NOT be trusted as the authority for identity truth.

---

## 14. Key Change UX

Key changes MUST be loud.

Silent key replacement is forbidden.

If a contact identity changes, CarbonStackComms MUST clearly show:

- which contact changed;
- what changed;
- when the client observed it;
- whether messages are blocked pending verification;
- what verification action is required.

For high-assurance profiles, key replacement SHOULD require:

- hardware-key approval;
- recovery passphrase or trusted-device approval;
- group-visible notice;
- explicit re-verification.

Server-initiated silent identity replacement MUST be treated as hostile.

---

## 15. Group Membership Model

CarbonStack is designed for small trusted groups first.

Group membership changes SHOULD be:

- explicit;
- auditable;
- visible to all members;
- tied to identity state;
- resistant to server rollback;
- resistant to silent member injection.

CarbonStackComms SHOULD detect attempts to:

- add unknown members silently;
- remove members silently;
- roll back group state;
- replay stale membership epochs;
- present inconsistent group views.

CarbonStackCypher may route group messages but MUST NOT be trusted to define group truth.

Strong hostile-server group integrity may require protocol maturity and is not necessarily complete in MVP-1. It remains a core requirement for mature CarbonStackComms.

---

## 16. Text and Character Policy

Text-only does not mean automatically safe.

Unicode is a meaningful attack surface.

CarbonStackOS and CarbonStackComms should use strict text normalization and rendering.

### 16.1 MVP Chat Character Profile

MVP chat SHOULD begin with:

- printable ASCII;
- newline;
- small curated punctuation set.

Broader Unicode support MAY be added later only with:

- strict normalization;
- visible blocking behavior;
- dangerous category rejection;
- test corpus coverage;
- bundled font strategy.

### 16.2 Baseline Text Rules

For all supported text profiles:

- UTF-8 only;
- reject invalid byte sequences;
- normalize to NFC where Unicode is supported;
- reject or visibly mark control characters;
- reject bidi override/control characters;
- reject zero-width characters;
- reject private-use characters;
- reject unassigned codepoints;
- reject dangerous combining sequences;
- limit maximum message size;
- limit maximum line length.

Unexpected characters should not silently render.

Preferred behavior:

- `[U+202E blocked]`
- `[U+200B blocked]`
- `[unsupported character]`

### 16.3 Rendering Rules

Rendering rules:

- bundled fonts only;
- no downloadable fonts;
- no emoji fallback engine unless tightly curated;
- no HTML;
- no CSS;
- no Markdown in chat by default;
- no clickable links by default;
- no automatic URL detection;
- no embedded previews.

Principle:

> Characters should be boring, visible, normalized, and unsurprising.

---

## 17. Local Notes and Markdown

CarbonStackOS may include a local notes system.

Allowed note formats:

- `.txt`
- `.md`

Extensions are not trusted. Files must be validated by content and parser rules.

Markdown should be a tiny subset, not web content.

Allowed Markdown subset:

- headings;
- bold;
- italic;
- unordered lists;
- numbered lists;
- blockquotes;
- inline code;
- fenced code blocks.

Rejected Markdown features:

- HTML;
- inline SVG;
- images;
- external links as clickable objects;
- tables unless explicitly added later;
- footnotes;
- LaTeX/math;
- embedded files;
- automatic linkification;
- custom extensions;
- remote content.

Markdown should render through native constrained text rendering, not WebView.

The notes app MUST NOT have default access to:

- network;
- message vault;
- camera;
- sensors;
- USB control;
- arbitrary filesystem browsing.

The notes app MAY receive explicit exported text from CarbonStackComms only through a narrow user-approved share/export path.

---

## 18. File Transfer Philosophy

CarbonStackOS may support file transfer, but only through a constrained, explicit, quarantine-first model.

Allowed inbound user file types for the planned constrained utility layer:

- TXT;
- MD;
- WAV;
- FLAC.

Additional internal or system-only formats MAY exist:

- signed manifests;
- canonical internal playlist format;
- encrypted backup bundles;
- public diagnostics bundles.

These internal formats are not general-purpose file acceptance.

Important rule:

> Do not trust file extensions. Validate structure, normalize, strip, and only then expose.

Inbound flow:

1. receive file into quarantine;
2. validate content;
3. reject ambiguous or malformed files;
4. strip metadata;
5. rewrite into canonical internal form;
6. move accepted file into approved local library;
7. deny direct open-from-transfer.

Rejected outright for this iteration:

- ZIP;
- archives;
- PDF;
- DOCX;
- HTML;
- SVG;
- JPEG;
- PNG;
- MP4;
- MOV;
- M4A;
- MP3;
- executables;
- scripts;
- unknown containers;
- playlist imports.

MP3 and PDF are not deferred media goals in this iteration. They are rejected for this profile.

The system should reject unsupported types at every layer:

- transfer protocol;
- quarantine validator;
- canonical storage;
- file browser;
- open-with routing;
- media player;
- text editor;
- messenger attachment logic;
- backup/export tooling.

---

## 19. Audio System

CarbonStackOS may include a barebones local audio player in a future constrained utility layer.

Allowed audio formats:

- WAV;
- FLAC.

Avoid VLC. VLC is too broad a parser ecosystem for the base system.

The CarbonStackOS audio player should be purpose-built and minimal.

Allowed features:

- play;
- pause;
- seek;
- volume;
- local playlists created on-device;
- loop one;
- loop playlist;
- shuffle if local only;
- basic file browser over approved music library.

Avoid:

- album art;
- lyrics;
- online metadata lookup;
- visualizers;
- DSP plugins;
- equalizer plugins;
- casting;
- Bluetooth metadata sync beyond bare minimum;
- sharing menus;
- opening arbitrary URIs;
- third-party codecs;
- playlist import formats.

### 19.1 WAV Rules

Accepted WAV should be narrow:

- RIFF/WAVE;
- PCM only;
- 16-bit or 24-bit integer PCM;
- 1 or 2 channels;
- 44100 Hz or 48000 Hz, maybe 96000 Hz later;
- maximum duration;
- maximum file size;
- known chunks only;
- reject or strip metadata chunks;
- reject compressed WAV codecs.

Inbound WAV files SHOULD be rewritten into canonical WAV form.

### 19.2 FLAC Rules

Accepted FLAC should be narrow:

- native FLAC stream only;
- no Ogg-FLAC initially;
- no embedded pictures;
- no exposed Vorbis comments by default;
- no cue sheets;
- seek tables validated or regenerated;
- sane sample rates;
- sane channel count;
- sane bit depth;
- maximum duration;
- maximum file size.

Inbound FLAC SHOULD be stripped and canonicalized before playback.

---

## 20. Sandboxing Layout

CarbonStackOS should separate authority aggressively.

Suggested compartments:

### 20.1 Transfer Service

Receives files only.

Must not have:

- message database access;
- media playback authority;
- broad network access;
- key access.

### 20.2 Quarantine Validator

Parses inbound files.

Must not have:

- network access;
- message keys;
- contacts;
- long-lived storage except quarantine and canonical output.

### 20.3 Text Editor

Opens only canonical TXT/MD.

Must not have:

- network access;
- message keys;
- broad filesystem access;
- automatic share authority.

### 20.4 Audio Player

Opens only canonical WAV/FLAC.

Must not have:

- network access;
- message keys;
- USB control;
- arbitrary URI opening.

### 20.5 Messenger

Text-only by default.

Must not have:

- inline attachment rendering;
- automatic file opening;
- rich previews;
- browser rendering.

### 20.6 Secure Vault

Stores:

- message database;
- identity keys;
- group state;
- trust records.

Secure Vault is the highest-value app data domain.

Principle:

> A FLAC bug should not become a message compromise. A Markdown bug should not become a key compromise.

---

## 21. USB Transfer Model

USB data should be off by default.

No normal MTP-style general filesystem access.

### 21.1 Device Locked

When device is locked:

- USB data fully disabled;
- charge-only where hardware permits.

### 21.2 Device Unlocked

When device is unlocked:

- charge-only by default;
- USB data does not automatically become available.

### 21.3 Transfer Mode

Transfer mode requires:

- explicit local approval;
- optional hardware-key confirmation;
- narrow custom protocol;
- short session timeout.

Transfer mode forbids:

- MTP;
- ADB;
- mass storage;
- accessory protocols;
- arbitrary filesystem browsing.

### 21.4 Inbound Transfer

Inbound transfer:

1. host sends files through narrow protocol;
2. files enter quarantine;
3. files are validated;
4. files are canonicalized;
5. only accepted files enter approved library.

### 21.5 Export

Export should be explicit.

Allowed export types:

- selected notes export;
- selected audio export;
- encrypted backup bundle;
- signed manifest;
- public diagnostics bundle.

The host should never get general browsing access to the device filesystem.

---

## 22. Bluetooth Model

Bluetooth is a convenience cost, not a free feature.

Allowed baseline:

- Bluetooth disabled by default;
- optional A2DP audio output only;
- manual pairing only;
- no always-discoverable mode.

Disallowed:

- Fast Pair;
- Nearby Share;
- BLE convenience ecosystem;
- file transfer profiles;
- contact sharing;
- message access profile;
- smartwatch integrations;
- companion-device integrations;
- automatic device discovery.

Recommended system modes:

### Normal Appliance Mode

Bluetooth audio MAY be enabled if the user explicitly permits it.

### High-Security Mode

Bluetooth fully disabled.

### Duress / Lockdown Mode

Bluetooth forcibly disabled.

Only the privileged state machine may restore Bluetooth after lockdown.

Wired audio remains the high-security recommendation.

---

## 23. Camera Model

Camera is not a core MVP requirement.

Initial recommendation:

- Phase 1: no camera;
- Phase 2: still images only, if necessary;
- Phase 3: video only if absolutely necessary.

If camera is added later:

- no EXIF;
- no GPS metadata;
- no timestamp metadata exposed to apps;
- no face detection;
- no OCR;
- no object recognition;
- no cloud backup;
- no Google Lens-style features;
- no automatic gallery intelligence.

Files should be named by relative sequence:

- `IMG-000001`;
- `IMG-000002`;
- `VID-000001`.

Not by wall-clock time.

Camera pipelines are complex. Camera support is a major attack-surface expansion.

---

## 24. Sensors

Default posture:

- deny sensor access unless required;
- no background sensor access;
- no motion analytics;
- no ambient audio intelligence;
- no assistant listening;
- no location stack unless explicitly added.

In lockdown or duress mode:

- camera disabled;
- microphone disabled;
- location disabled;
- motion sensors disabled where platform allows;
- Bluetooth disabled;
- Wi-Fi disabled;
- USB data disabled.

Sensor gating should be enforced below ordinary app permission prompts wherever possible.

---

## 25. Duress System

CarbonStackOS includes configurable duress behavior.

Key principle:

> The device should already be encrypted. A duress PIN does not "encrypt the device"; it evicts keys, locks compartments, disables interfaces, and changes system state.

Separate secrets:

- **Normal PIN**: unlocks normal profile.
- **Duress PIN**: triggers configured duress behavior.
- **Recovery PIN/passphrase**: begins recovery from duress.
- **Hardware key**: possession factor for recovery, enrollment, and revocation.
- **Biometric**: optional local intent check, not root identity.

### 25.1 Decoy Mode

Purpose:

- someone is looking over the user's shoulder;
- device should not obviously panic.

Behavior:

- opens limited decoy profile;
- hides secure conversations;
- hides contact names;
- suppresses sensitive notifications;
- keeps harmless notes/music if configured;
- does not reveal duress state.

### 25.2 Lockdown Mode

Purpose:

- user may lose physical control of device.

Behavior:

- locks secure vault;
- evicts vault keys from memory;
- stops messaging;
- disables Wi-Fi;
- disables Bluetooth;
- disables USB data;
- disables camera;
- disables microphone;
- disables sensors where possible;
- hides notifications;
- requires recovery ceremony.

### 25.3 Burn Mode

Purpose:

- device should no longer be recoverable locally.

Behavior:

- destroys local vault keys;
- destroys local message database access;
- marks device identity as locally compromised;
- requires re-enrollment;
- can later send signed revocation if network returns.

Burn mode is optional and dangerous.

Burn mode MUST NOT be the default duress action.

Decoy, lockdown, and burn remain in-scope even if only lockdown is implemented first.

---

## 26. Duress Recovery Flow

Recovery should be staged and narrow.

Suggested flow:

1. duress state active;
2. enter dedicated recovery PIN/passphrase;
3. biometric local confirmation, if configured;
4. temporary recovery window opens;
5. USB-HID/FIDO2-only path enabled;
6. hardware key challenge-response;
7. secure vault keys are unsealed or rewrapped;
8. system restores configured normal state.

USB should not generally return during recovery.

Allowed during recovery window:

- USB HID;
- FIDO2/security-key protocol.

Still blocked:

- MTP;
- ADB;
- mass storage;
- MIDI;
- tethering;
- vendor accessory protocols;
- file transfer;
- debugging;
- general USB data.

Recovery window:

- short timeout, e.g. 60 to 120 seconds;
- failed hardware-key attempt closes window;
- retry requires recovery PIN/passphrase again;
- increasing delay after failures.

Configurable setting:

- Allow local recovery after duress: ON/OFF.

Per-mode example:

- Decoy PIN: local recovery allowed.
- Lockdown PIN: local recovery allowed or disabled by user.
- Burn PIN: local recovery impossible by definition.

Biometric should be optional. Hardware key plus recovery passphrase should remain the stronger root recovery path.

---

## 27. Duress State Machine

The duress system should be implemented as a small privileged state machine, not scattered app behavior.

Possible states:

- NORMAL;
- DECOY;
- DURESS_LOCKDOWN;
- RECOVERY_ARMED;
- RECOVERY_USB_KEY_ONLY;
- RECOVERED;
- BURNED;
- REENROLLMENT_ONLY.

Allowed transitions:

- NORMAL â†’ DECOY;
- NORMAL â†’ DURESS_LOCKDOWN;
- NORMAL â†’ BURNED;
- DECOY â†’ RECOVERY_ARMED;
- DURESS_LOCKDOWN â†’ RECOVERY_ARMED;
- RECOVERY_ARMED â†’ RECOVERY_USB_KEY_ONLY;
- RECOVERY_USB_KEY_ONLY â†’ RECOVERED;
- BURNED â†’ REENROLLMENT_ONLY.

This prevents components from independently re-enabling dangerous features.

Example bad failure:

- Bluetooth comes back because the audio player asked for it.

Example desired behavior:

- only the privileged state machine can restore radios, sensors, vault access, or transfer mode.

---

## 28. Encryption Domains

CarbonStackOS should use separated encryption domains.

Suggested domains:

### 28.1 OS Base

- read-only;
- verified;
- immutable from ordinary user perspective.

### 28.2 User Profile

- local non-sensitive settings;
- optional notes;
- optional music if configured.

### 28.3 Secure Messaging Vault

Contains:

- message database;
- identity keys;
- group state;
- contact trust records;
- revocation state.

### 28.4 Quarantine / Import Vault

Contains inbound files before validation.

Must not have access to message keys.

### 28.5 Recovery Vault

Contains minimal state required for recovery ceremony.

Should not contain message plaintext.

Lockdown should:

- evict keys;
- unmount or cryptographically lock sensitive vaults;
- deny app access;
- disable transfer endpoints;
- disable media indexing;
- make quarantine inaccessible.

Lockdown should not merely hide files.

---

## 29. Notifications

Notifications should avoid leaking content.

Default:

- no message body on lock screen;
- no sender names on lock screen unless configured;
- no rich previews;
- no images;
- no action buttons that expose state.

In duress/decoy:

- sensitive notifications suppressed;
- messaging sync stopped or hidden;
- no obvious "duress active" banner.

Disallowed:

- inline replies from lock screen;
- rich message previews;
- decrypted notification payloads;
- sender avatar display in high-security mode;
- automatic link previews.

Notifications should announce activity without becoming a side channel for message content, identity state, or duress state.

---

## 30. Updates

OS updates should be:

- signed;
- verified;
- rollback-protected where supported;
- delivered through minimal trusted updater;
- not dependent on general browser/web stack.

The device should not accept:

- unsigned OS images;
- older vulnerable rollback images;
- user-modified system partitions;
- untrusted update channels.

For project credibility, aim toward:

- published hashes;
- signed release manifests;
- reproducible builds where practical;
- independent build instructions;
- clear security changelogs.

A malicious signed update remains a catastrophic risk.

Signing infrastructure should be treated as high-value security infrastructure.

Recommended future controls:

- offline signing where practical;
- hardware-backed signing keys where practical;
- documented key ceremony;
- release signatures;
- emergency key revocation;
- build environment isolation.

---

## 31. Recovery, Revocation, and Device Loss

CarbonStack must distinguish:

### 31.1 Lock

Local keys are evicted. Data remains recoverable through normal recovery ceremony.

### 31.2 Wipe

Local encrypted data or vault material is deleted.

### 31.3 Revoke

Other users/groups are told a device identity is no longer trusted.

Important limitation:

An offline seized device can lock or wipe itself. It cannot notify others until some trusted channel reaches the network.

CarbonStackComms should eventually support:

- hardware-key-signed device revocation;
- group-visible device removal;
- forced re-verification after device recovery;
- new-device enrollment ceremony;
- compromise notices.

Device replacement must be visible. Silent replacement is forbidden.

---

## 32. CarbonStackOS Non-Goals

CarbonStackOS should avoid:

- browser engines;
- WebView-dependent apps;
- general APK installation;
- arbitrary code execution;
- Google Play Services dependency;
- rich notifications;
- link previews;
- cloud sync;
- assistant features;
- general file managers;
- cellular modem dependence;
- third-party keyboards;
- third-party accessibility services;
- general media ecosystems.

CarbonStackOS is not:

- a general smartphone OS;
- a privacy-themed Android ROM for daily-driver use;
- an app platform;
- a browser platform;
- a media-first device;
- a cloud-sync device.

It is a communications appliance.

---

## 33. CarbonStackComms Non-Goals

CarbonStackComms MVP is not:

- a Signal clone;
- a Discord clone;
- a WhatsApp clone;
- a Matrix replacement;
- a social network;
- a rich-media chat app;
- a file-sharing platform.

CarbonStackComms MVP rejects:

- inline attachments;
- rich previews;
- stickers;
- GIFs;
- voice messages;
- automatic link detection;
- clickable links by default;
- arbitrary Unicode rendering;
- browser rendering;
- Markdown chat rendering;
- hidden identity changes.

---

## 34. CarbonStackCypher Non-Goals

CarbonStackCypher is not:

- a trusted identity authority;
- a plaintext message store;
- a cloud account provider;
- a social discovery service;
- a general push notification content service;
- a centralized trust root.

CarbonStackCypher MUST NOT be required to be honest for message confidentiality.

It may still be able to observe metadata in MVP-1.

CarbonStack does not initially claim strong metadata privacy against the relay.

Metadata resistance is a future research and design area.

---

## 35. Security Assurance Matrix

The Security Assurance Matrix tracks how CarbonStack claims are tested, limited, and validated.

Every major claim should map to:

- intended protection;
- mechanism;
- test method;
- failure mode;
- known limitation;
- project phase.

### 35.1 Protocol Assurance

Questions:

- Can the server read plaintext?
- Can the server silently replace identity keys?
- Can the server silently add group members?
- Can messages be replayed?
- Can group state be rolled back?
- Are key changes visible?

Tests:

- hostile-server simulation;
- replay tests;
- rollback tests;
- identity-replacement tests;
- group-membership mutation tests;
- client warning tests.

### 35.2 Parser Assurance

Questions:

- Are invalid byte sequences rejected?
- Are dangerous Unicode controls blocked or marked?
- Are unsupported files rejected?
- Are accepted files canonicalized?
- Can parser compromise reach message keys?

Tests:

- Unicode rejection corpus;
- malformed TXT/MD corpus;
- malformed WAV/FLAC corpus;
- quarantine validation tests;
- compartment permission tests.

### 35.3 Device Assurance

Questions:

- Is boot state verified?
- Is the bootloader locked?
- Is the OS signed?
- Is rollback blocked?
- Is the app running on an approved device profile?
- Are unsafe device states refused?

Tests:

- unlocked bootloader test;
- stale patch test;
- developer-options test;
- USB-debugging test;
- root-detection test;
- app-signature mismatch test;
- device-profile mismatch test.

### 35.4 Local Vault Assurance

Questions:

- Is the vault cryptographically locked?
- Are keys evicted on lock?
- Are keys evicted on duress?
- Can notes or media compartments access message keys?
- Can backups expose plaintext?

Tests:

- session timeout tests;
- lockdown tests;
- memory/key lifecycle tests;
- compartment access tests;
- backup export tests.

### 35.5 Server Assurance

Questions:

- Does CarbonStackCypher store only encrypted envelopes?
- Are admin actions logged?
- Can admin access plaintext?
- Are registrations invite-only?
- Can revocation propagate?
- Can the server forge trust events?

Tests:

- database inspection tests;
- admin permission tests;
- audit log tests;
- invite-flow tests;
- revocation-flow tests;
- malicious relay tests.

### 35.6 OS Appliance Assurance

Questions:

- Is the system free of browser engines?
- Is WebView absent or unusable by normal apps?
- Is APK sideloading blocked?
- Is ADB disabled in production?
- Is USB data disabled by default?
- Is Bluetooth restricted?
- Are updates signed?

Tests:

- installed package audit;
- forbidden component audit;
- USB mode tests;
- Bluetooth profile tests;
- update signature tests;
- rollback tests;
- production-build policy tests.

Example row format:

| Claim | Mechanism | Test | Failure Mode | Phase |
|---|---|---|---|---|
| Server cannot read messages | client-side E2EE | inspect database and relay logs | client leaks plaintext before encryption | MVP-1 |
| Key changes are loud | trust-state UI | forced key replacement test | user ignores warning | MVP-1 |
| FLAC parser cannot access message vault | compartment separation | permission and IPC test | OS sandbox escape | MVP-3 |
| Lockdown evicts vault keys | privileged state machine | trigger lockdown and attempt vault read | compromised kernel | MVP-2 |

---

## 36. Deployment and Accessibility Goals

CarbonStack is open source and should be deployable.

The project MUST avoid security-through-inaccessibility.

Documentation should allow a technically competent user to:

- understand the threat model;
- deploy CarbonStackCypher;
- build or install CarbonStackComms;
- understand trust warnings;
- perform contact verification;
- understand recovery limitations;
- understand server compromise assumptions;
- understand endpoint compromise assumptions;
- understand what CarbonStack does not protect against.

Deployment docs SHOULD include:

- minimal server deployment;
- update process;
- backup process;
- admin authentication setup;
- invite-only registration setup;
- test instance setup;
- production hardening notes;
- known insecure configurations.

Complexity is allowed only where it is visible and justified.

---

## 37. Fork Policy

CarbonStack is open source.

Forks are encouraged to explore different tradeoffs.

However, forks that add rejected features should not claim the same security profile without clearly documenting the change.

Examples:

- a fork with cellular support should state that it expands the radio/baseband/carrier attack surface;
- a fork with browser support should state that it is no longer CarbonStackOS appliance-profile compatible;
- a fork with PDF/MP3 support should state that it expands parser exposure;
- a fork with Play Services should state the dependency and trust implications.

CarbonStack should be strict in its canonical profile without trying to prevent experimentation elsewhere.

---

## 38. Remaining Hard Risks

CarbonStack reduces attack surface. It does not create perfect security.

Remaining hard risks:

- kernel vulnerabilities;
- Wi-Fi firmware/driver bugs;
- Bluetooth stack bugs if enabled;
- USB controller bugs;
- GPU/display driver bugs;
- camera/vendor HAL bugs;
- supply-chain compromise;
- malicious signed update;
- compromised build system;
- physical coercion;
- hardware implants;
- evil-maid attacks;
- compromised recipient devices;
- metadata leakage;
- group member betrayal;
- server traffic analysis;
- user error;
- social engineering.

Core truth:

> Encrypted messaging cannot protect plaintext once it is displayed on a compromised endpoint.

CarbonStackOS is about making endpoint compromise much harder by refusing to be a general-purpose endpoint.

---

## 39. Open Research Questions

CarbonStack must still research:

- final messaging protocol selection;
- libsignal viability and licensing implications;
- MLS viability for small groups;
- Noise-based channel viability;
- Pixel 4a testbed practicality;
- Pixel 9a reference target practicality;
- verified boot and signing workflow;
- custom OS update workflow;
- hardware-key identity binding;
- Android hardware-backed keystore behavior;
- StrongBox availability by target;
- device attestation strategy;
- how to implement CarbonStackOS without becoming an unmaintainable ROM project;
- how to safely implement duress state machine privileges;
- how to quarantine and canonicalize WAV/FLAC safely;
- how to avoid WebView and browser engine dependencies;
- how to test parser compartment boundaries;
- how to minimize metadata leakage without making the system unusable.

Research findings should be documented in the relevant repo and summarized in the canonical spec when they affect doctrine.

---

## 40. Suggested Initial File Map

Recommended current repository layout:

```text
carbonstack/
  README.md
  docs/
    00-project-charter.md
    01-threat-model.md
    02-non-goals.md
    03-stack-architecture.md
    04-text-character-policy.md
    05-file-transfer-policy.md
    06-duress-system.md
    07-encryption-domains.md
    08-usb-transfer-model.md
    09-bluetooth-model.md
    10-notifications.md
    11-updates.md
    12-recovery-revocation-device-loss.md
    13-security-assurance-matrix.md
    14-carbonstack-full-specification.md
  roadmap/
    ROADMAP.md
```

Implementation repositories:

```text
carbonstack-comms/
  docs/
  src/
  tests/
  security/

carbonstack-cypher/
  docs/
  server/
  tests/
  deploy/
  security/

carbonstack-os/
  docs/
  device-profiles/
  build/
  security/
  state-machine/
```

---

## 41. Glossary

### Appliance OS

An operating system constrained for a narrow purpose, not a general-purpose user endpoint.

### CarbonStack

The overall secure communications appliance project.

### CarbonStackOS

The hardened Android-derived appliance operating system.

### CarbonStackComms

The text-first encrypted messaging client.

### CarbonStackCypher

The hostile-server relay and storage stack.

### Hostile Server

A server assumed to be malicious, compromised, coerced, or unreliable.

### Secure Vault

The cryptographic domain containing message database, identity keys, group state, and trust records.

### Quarantine

Temporary containment area for inbound files before validation and canonicalization.

### Canonicalization

Rewriting accepted input into a narrow internal form before exposing it to users or apps.

### Loud Trust Change

A visible, unavoidable warning when identity, key, membership, or device state changes.

### Duress PIN

A PIN that triggers decoy, lockdown, or burn behavior instead of normal unlock.

### Burn Mode

Optional dangerous mode that destroys local vault access and requires re-enrollment.

### Recovery Ceremony

A staged process using recovery secret and hardware key to restore access after duress or lockout.

### High-Assurance Mode

CarbonStack profile requiring controlled hardware, verified OS state, and hardware-key-backed trust workflows.

### Low-Assurance Mode

Normal Android app mode used for development, testing, and broader compatibility without full appliance guarantees.

---

## 42. Final Doctrine Summary

CarbonStack is not trying to be convenient by default.

CarbonStack is trying to be narrow enough to reason about.

The mature system should be:

- open source;
- deployable;
- text-first;
- hostile-server aware;
- explicit about trust;
- strict about files;
- strict about radios;
- strict about parsers;
- honest about endpoints;
- documented enough for human operators;
- constrained enough that compromise of one part does not automatically compromise everything.

The purpose of CarbonStackOS is not to make Android nicer.

The purpose of CarbonStackOS is to stop being a smartphone.

