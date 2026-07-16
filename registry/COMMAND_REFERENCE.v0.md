# CarbonStack Command Reference v0

Status: **generated dev/operator command reference**

Generated from `registry/commands.v0.yaml` by `tools/registry/render-command-reference.py`.

Do not hand-edit this file. Update the registry source and rerun the renderer.

Boundary:

- Registry presence is classification, not promotion.
- This reference is dev/operator-facing, not general-public UX documentation.
- This is not a production security claim.
- This is not a man-page set.
- `full`, `release-snapshot`, and `integrated-runtime-dev` are distinct and must not be merged.
- Internal sidecar commands, Cypher API surfaces, and legacy scripts are documented for boundary clarity, not promoted as user-facing commands.

Registry entry count: **139**

## Release/package validation and package-helper profiles

Entries in this section: **4**

### `runner.full`

- **Command:** `go run . --profile full`
- **Repo:** `carbonstack`
- **Component:** `tools/carbonstack-validate`
- **Kind:** `runner-profile`
- **Audience:** `public`
- **Maturity:** `release_supported`
- **Introduced in:** `v0.3.34`
- **Source path:** `tools/carbonstack-validate/main.go`
- **Validation surface:** release package validation ladder
- **Front README candidate:** `true`

**What it does:** Run release-snapshot, then local-cypher, from a fresh release package root.

**Why it exists:** Current v0.4.0 release package validation ladder.

- **Example:** `cd carbonstack/tools/carbonstack-validate && go run . --profile full --root /path/to/release-package-root --clean-generated`
- **Required flags:** Not recorded in registry.
- **Optional flags:**
  - `--root` — release package root to validate. Boundary: intended for fresh extracted or throwaway staged package roots, not the live umbrella.
  - `--clean-generated` — clean known generated roots after successful validation. Boundary: does not make full an integrated runtime profile.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - not deployment
  - not local-backbone
  - not runtime Comms UX
  - not production security proof

### `runner.release-snapshot`

- **Command:** `go run . --profile release-snapshot`
- **Repo:** `carbonstack`
- **Component:** `tools/carbonstack-validate`
- **Kind:** `runner-profile`
- **Audience:** `dev`
- **Maturity:** `release_supported`
- **Introduced in:** `v0.3.34`
- **Source path:** `tools/carbonstack-validate/release_snapshot.go`
- **Validation surface:** release package layout/checksum/core validation
- **Front README candidate:** `false`

**What it does:** Validate release package layout, metadata, checksums, strict pre-test artifact hygiene, then core.

**Why it exists:** Prevents generated/private/build artifacts from leaking into release package validation roots.

- **Example:** `cd carbonstack/tools/carbonstack-validate && go run . --profile release-snapshot --root /path/to/release-package-root`
- **Required flags:** Not recorded in registry.
- **Optional flags:**
  - `--root` — release package root to validate. Boundary: package-root validation only; not live integrated runtime.
  - `--clean-generated` — clean known generated roots after successful validation. Boundary: does not add integrated-runtime-dev to release-snapshot.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - does not package or upload releases
  - not deployment
  - not production security proof

### `runner.verify-checksums`

- **Command:** `go run . --profile verify-checksums`
- **Repo:** `carbonstack`
- **Component:** `tools/carbonstack-validate`
- **Kind:** `runner-profile`
- **Audience:** `public`
- **Maturity:** `release_supported`
- **Introduced in:** `v0.3.x`
- **Source path:** `tools/carbonstack-validate/checksums.go`
- **Validation surface:** release checksum verification
- **Front README candidate:** `true`

**What it does:** Verify release/checksums.txt against the current package root.

**Why it exists:** Confirms release package files match generated checksum metadata.

- **Example:** `cd carbonstack/tools/carbonstack-validate && go run . --profile verify-checksums --root /path/to/package-root`
- **Required flags:** Not recorded in registry.
- **Optional flags:**
  - `--root` — package root whose release/checksums.txt should be verified. Boundary: checksum verification only; does not replace full validation.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - does not prove runtime security
  - does not replace full validation

### `runner.write-checksums`

- **Command:** `go run . --profile write-checksums`
- **Repo:** `carbonstack`
- **Component:** `tools/carbonstack-validate`
- **Kind:** `runner-profile`
- **Audience:** `dev`
- **Maturity:** `release_supported`
- **Introduced in:** `v0.3.x`
- **Source path:** `tools/carbonstack-validate/checksums.go`
- **Validation surface:** release checksum generation
- **Front README candidate:** `false`

**What it does:** Write SHA-256 checksums for a clean release package root.

**Why it exists:** Produces release/checksums.txt for release package verification.

- **Example:** `cd carbonstack/tools/carbonstack-validate && go run . --profile write-checksums --root /path/to/package-root`
- **Required flags:** Not recorded in registry.
- **Optional flags:**
  - `--root` — package root to write checksums for. Boundary: release helper only; does not publish release assets.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - does not publish releases
  - does not validate runtime security

## Diagnostic, core, local, and registry inspection profiles

Entries in this section: **4**

### `runner.core`

- **Command:** `go run . --profile core`
- **Repo:** `carbonstack`
- **Component:** `tools/carbonstack-validate`
- **Kind:** `runner-profile`
- **Audience:** `public`
- **Maturity:** `experimental`
- **Introduced in:** `pre-v0.4.x`
- **Source path:** `tools/carbonstack-validate/main.go`
- **Validation surface:** core implementation validation
- **Front README candidate:** `true`

**What it does:** Run doctor, OpenMLS real-Cypher lifecycle test, Comms tests, Cypher tests, and artifact scans.

**Why it exists:** Core implementation sanity path for source checkouts and release snapshot internals.

- **Example:** `cd carbonstack/tools/carbonstack-validate && go run . --profile core --clean-generated`
- **Required flags:** Not recorded in registry.
- **Optional flags:**
  - `--root` — optional umbrella/package root override. Boundary: root must contain the required carbonstack, carbonstack-comms, and carbonstack-cypher paths.
  - `--clean-generated` — clean known generated roots after successful validation. Boundary: only known OpenMLS sidecar generated roots are cleaned by this runner flag.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - not deployment
  - not production security proof

### `runner.doctor`

- **Command:** `go run . --profile doctor`
- **Repo:** `carbonstack`
- **Component:** `tools/carbonstack-validate`
- **Kind:** `runner-profile`
- **Audience:** `public`
- **Maturity:** `experimental`
- **Introduced in:** `pre-v0.4.x`
- **Source path:** `tools/carbonstack-validate/main.go`
- **Validation surface:** runner self-check
- **Front README candidate:** `true`

**What it does:** Report environment, inferred repo layout, required paths, and toolchain versions.

**Why it exists:** Baseline diagnostics before deeper validation.

- **Example:** `cd carbonstack/tools/carbonstack-validate && go run . --profile doctor`
- **Required flags:** Not recorded in registry.
- **Optional flags:**
  - `--root` — optional umbrella root override; defaults to inferred parent layout where possible. Boundary: does not install or mutate toolchains.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - does not install or mutate toolchains
  - does not prove deployment readiness

### `runner.local-cypher`

- **Command:** `go run . --profile local-cypher`
- **Repo:** `carbonstack`
- **Component:** `tools/carbonstack-validate`
- **Kind:** `runner-profile`
- **Audience:** `public`
- **Maturity:** `experimental`
- **Introduced in:** `v0.3.30`
- **Source path:** `tools/carbonstack-validate/local_cypher.go`
- **Validation surface:** local Cypher lifecycle validation
- **Front README candidate:** `true`

**What it does:** Validate a temporary local Cypher lifecycle with invite, device, envelope, ack, restart, and negative protocol checks.

**Why it exists:** Proves local Cypher relay/storage lifecycle without claiming whole-stack local-backbone.

- **Example:** `cd carbonstack/tools/carbonstack-validate && go run . --profile local-cypher`
- **Required flags:** Not recorded in registry.
- **Optional flags:**
  - `--root` — optional umbrella/package root override. Boundary: root must contain carbonstack-cypher and validation dependencies.
  - `--clean-generated` — clean known generated roots after successful validation. Boundary: does not silently delete arbitrary Cypher databases.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - not local-backbone
  - not runtime Comms UX
  - not public ingress
  - not deployment

### `runner.registry-lookup`

- **Command:** `go run . --profile registry-lookup --registry-id <id>`
- **Repo:** `carbonstack`
- **Component:** `tools/carbonstack-validate`
- **Kind:** `runner-profile`
- **Audience:** `dev`
- **Maturity:** `experimental`
- **Introduced in:** `v0.6.7`
- **Source path:** `tools/carbonstack-validate/registry_lookup.go`
- **Validation surface:** command-boundary registry lookup
- **Front README candidate:** `false`

**What it does:** Print a dev/operator command-registry entry by registry ID or literal command.

**Why it exists:** Makes command-boundary classification, maturity, lifecycle status, source path, validation surface, related surfaces, and nonclaims inspectable without manually reading the full registry YAML.

- **Example:** `cd carbonstack/tools/carbonstack-validate && go run . --profile registry-lookup --registry-id comms.message-send-dev`
- **Required flags:** Not recorded in registry.
- **Optional flags:**
  - `--registry-id` — Registry ID to inspect, such as comms.message-send-dev. Boundary: Lookup selector only; does not execute the command or validate the surface.
  - `--command` — Literal command string to inspect, such as go run ./cmd/comms message-send-dev. Boundary: Lookup selector only; may be ambiguous if multiple entries share command text.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - registry presence is not promotion
  - not generated public manual
  - not production UX
  - not security certification

## Live-dev runtime validation profiles

Entries in this section: **4**

### `runner.dev-runtime-openmls`

- **Command:** `go run . --profile dev-runtime-openmls`
- **Repo:** `carbonstack`
- **Component:** `tools/carbonstack-validate`
- **Kind:** `runner-profile`
- **Audience:** `dev`
- **Maturity:** `dev_only`
- **Lifecycle status:** `lower_level_direct_proof_transition_candidate`
- **Introduced in:** `v0.4.9`
- **Source path:** `tools/carbonstack-validate/dev_runtime_openmls.go`
- **Validation surface:** direct OpenMLS runtime smoke
- **Front README candidate:** `false`

**What it does:** Validate the direct-sidecar dev OpenMLS runtime smoke path through Comms and Cypher.

**Why it exists:** Runner-wrapped baseline proof for openmls-send-dev -> Cypher -> openmls-inbox-dev --ack.

- **Example:** `cd carbonstack/tools/carbonstack-validate && go run . --profile dev-runtime-openmls --clean-generated`
- **Required flags:** Not recorded in registry.
- **Optional flags:**
  - `--root` — optional live umbrella root override. Boundary: live-dev profile; not release-package validation.
  - `--clean-generated` — clean known generated roots after successful validation. Boundary: only known OpenMLS sidecar generated roots are cleaned.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - not local-backbone
  - not release-package validation
  - not included in full
  - not mature messaging UX
  - not production security proof

### `runner.dev-runtime-openmls-wrappers`

- **Command:** `go run . --profile dev-runtime-openmls-wrappers`
- **Repo:** `carbonstack`
- **Component:** `tools/carbonstack-validate`
- **Kind:** `runner-profile`
- **Audience:** `dev`
- **Maturity:** `dev_only`
- **Introduced in:** `v0.4.17`
- **Source path:** `tools/carbonstack-validate/dev_runtime_openmls_wrappers.go`
- **Validation surface:** Relay Space-scoped wrapper-based OpenMLS runtime smoke
- **Front README candidate:** `false`

**What it does:** Validate wrapper bootstrap plus Relay Space creation, active routing membership, scoped normal-message send/open, and ack-after-open through Comms and Cypher.

**Why it exists:** Separate maturity profile proving the current wrapper normal-message path with explicit Relay Space context before any future consolidation with the lower-level direct smoke.

- **Required flags:** Not recorded in registry.
- **Optional flags:**
  - `--root` — Explicit CarbonStack umbrella root.
  - `--clean-generated` — Remove known generated OpenMLS sidecar roots after validation.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - not local-backbone
  - does not replace dev-runtime-openmls yet
  - not release-package validation
  - not included in full
  - not mature messaging UX
  - active Relay membership is routing authority only
  - not identity verification
  - not production security proof

### `runner.integrated-runtime-dev`

- **Command:** `go run . --profile integrated-runtime-dev`
- **Repo:** `carbonstack`
- **Component:** `tools/carbonstack-validate`
- **Kind:** `runner-profile`
- **Audience:** `dev`
- **Maturity:** `dev_only`
- **Introduced in:** `v0.6.10`
- **Source path:** `tools/carbonstack-validate/integrated_runtime_dev.go`
- **Validation surface:** live-umbrella Relay onboarding plus Relay Space-scoped normal-message composition
- **Front README candidate:** `false`

**What it does:** Run Relay onboarding proof, then the scoped wrapper normal-message runtime proof, as a live-umbrella integrated dev profile.

**Why it exists:** Provides a bounded in-series validation ladder from Relay onboarding into explicit Relay Space context, active routing membership, scoped message send/open, and ack-after-open without mutating full or release-snapshot.

- **Required flags:** Not recorded in registry.
- **Optional flags:**
  - `--root` — Explicit CarbonStack umbrella root.
  - `--clean-generated` — Remove known generated OpenMLS sidecar roots after validation.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - not included in full
  - not included in release-snapshot
  - not release-package validation
  - not package-root validation
  - not production secure messaging
  - not production E2EE
  - not hostile-server safety
  - not metadata privacy
  - not identity verification
  - not secure enrollment
  - not local-backbone
  - not deployment
  - not mature messenger UX
  - not general-public UX
  - active Relay membership is routing authority only
  - does not replace relay-openmls-join-dev
  - does not replace dev-runtime-openmls-wrappers

### `runner.relay-openmls-join-dev`

- **Command:** `go run . --profile relay-openmls-join-dev`
- **Repo:** `carbonstack`
- **Component:** `tools/carbonstack-validate`
- **Kind:** `runner-profile`
- **Audience:** `dev`
- **Maturity:** `dev_only`
- **Lifecycle status:** `relay_onboarding_join_transition_candidate`
- **Introduced in:** `Not recorded in registry.`
- **Source path:** `tools/carbonstack-validate/relay_openmls_join_dev.go`
- **Validation surface:** live-umbrella-relay-openmls-join-dev-positive-path
- **Front README candidate:** `false`

**What it does:** Run the manual/dev positive-path Relay Space OpenMLS join validation profile; excluded from full/release-snapshot for v0.6.0.

**Why it exists:** Proves a bounded local/dev KeyPackage -> add-member -> Welcome -> join path with no-ack and ACK_AFTER_JOIN subruns using runner-owned temp state and strict nonclaims. As of v0.5.66, this remains manual/dev-only and excluded from full/release-snapshot until repeated clean runs, package-root rehearsal, artifact behavior review, and an explicit future inclusion decision.

- **Example:** `cd carbonstack/tools/carbonstack-validate && go run . --profile relay-openmls-join-dev --compact-summary`
- **Required flags:** Not recorded in registry.
- **Optional flags:**
  - `--compact-summary` — Print a compact evidence summary after successful profile execution. Boundary: Output/evidence convenience only; does not preserve generated artifacts or change validation scope.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - not included in full
  - not included in release-snapshot
  - not release-package validation
  - future full-profile candidate only after explicit inclusion decision
  - not local-backbone
  - not production secure messaging
  - not identity verification
  - not hostile-server safety
  - not metadata privacy
  - not audit or certification

## Recommended dev/pre-alpha normal-message wrappers

Entries in this section: **2**

### `comms.message-inbox-dev`

- **Command:** `go run ./cmd/comms message-inbox-dev`
- **Repo:** `carbonstack-comms`
- **Component:** `cmd/comms`
- **Kind:** `cli`
- **Audience:** `dev`
- **Maturity:** `dev_only`
- **Lifecycle status:** `recommended_dev_wrapper`
- **Introduced in:** `v0.6.5`
- **Source path:** `carbonstack-comms/internal/app/message_wrappers_dev.go`
- **Validation surface:** dev-runtime-openmls-wrappers, integrated-runtime-dev, same-state normal-message profiles, and Comms package tests
- **Front README candidate:** `true`

**What it does:** Fetch the Relay Space-scoped device inbox, classify ordinary-message envelopes, sidecar-open supported messages, and optionally ack only after successful open.

**Why it exists:** Recommended dev/pre-alpha ordinary-message receive/open wrapper after onboarding. It requires explicit Relay Space context, preserves unsupported content or protocol pairs for explicit no-ack classification, and never treats routing membership as identity, trust, or MLS authorization.

- **Required flags:**
  - `--relay-space` — Relay Space ID used for scoped device-inbox fetch and any subsequent scoped acknowledgement. Boundary: Required routing context; the command does not create the Relay Space or enroll members.
  - `--sidecar-device-label` — Local OpenMLS sidecar device-state label used for message open.
  - `--conversation` — Local OpenMLS conversation label used for message open.
- **Optional flags:**
  - `--state` — Local Comms state path; defaults to the normal Comms state path.
  - `--sidecar-dir` — Explicit OpenMLS sidecar project directory override.
  - `--message-label` — Optional application-message label filter.
  - `--limit` — Maximum scoped inbox records to inspect; defaults to one.
  - `--ack` — Request a scoped acknowledgement only after successful OpenMLS message-open. Boundary: Fetch, unsupported skip, and message-open failure do not acknowledge the envelope.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - dev/pre-alpha only
  - not mature inbox UX
  - not production E2EE claim
  - not hostile-server safety
  - not metadata privacy
  - not verified identity
  - not durable local receive storage
  - does not perform Relay onboarding
  - Relay membership is routing authority only
  - Relay membership is not identity, trust, or MLS authorization
  - no ack on fetch
  - no ack on unsupported content or protocol
  - no ack on message-open failure

### `comms.message-send-dev`

- **Command:** `go run ./cmd/comms message-send-dev`
- **Repo:** `carbonstack-comms`
- **Component:** `cmd/comms`
- **Kind:** `cli`
- **Audience:** `dev`
- **Maturity:** `dev_only`
- **Lifecycle status:** `recommended_dev_wrapper`
- **Introduced in:** `v0.6.5`
- **Source path:** `carbonstack-comms/internal/app/message_wrappers_dev.go`
- **Validation surface:** dev-runtime-openmls-wrappers, integrated-runtime-dev, same-state normal-message profiles, and Comms package tests
- **Front README candidate:** `true`

**What it does:** Protect an OpenMLS application message and submit it through the required Relay Space-scoped Cypher route.

**Why it exists:** Recommended dev/pre-alpha ordinary-message send wrapper after onboarding. It requires explicit Relay Space context plus active sender and recipient device routing membership; that membership is routing authority only and is not identity, trust, or MLS authorization.

- **Required flags:**
  - `--relay-space` — Relay Space ID used for the scoped application-message submission. Boundary: Required routing context; the command does not create the Relay Space or enroll members.
  - `--to-device` — Recipient Cypher device ID, which must be an active member of the selected Relay Space.
  - `--message` — Plaintext input protected by the OpenMLS sidecar before Relay submission.
  - `--sidecar-device-label` — Local OpenMLS sidecar device-state label used for message protection.
  - `--conversation` — Local OpenMLS conversation label used for message protection.
- **Optional flags:**
  - `--state` — Local Comms state path; defaults to the normal Comms state path.
  - `--strict` — Block unknown, unverified, changed, or revoked trust decisions where the current dev trust surface supports them.
  - `--sidecar-dir` — Explicit OpenMLS sidecar project directory override.
  - `--message-label` — Dev/test message label carried inside the protected application message.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - dev/pre-alpha only
  - not mature send UX
  - not production E2EE claim
  - not hostile-server safety
  - not metadata privacy
  - not verified identity
  - not secure enrollment
  - does not perform Relay onboarding
  - Relay membership is routing authority only
  - Relay membership is not identity, trust, or MLS authorization

## Lower-level direct OpenMLS message proof commands

Entries in this section: **2**

### `comms.openmls-inbox-dev`

- **Command:** `go run ./cmd/comms openmls-inbox-dev`
- **Repo:** `carbonstack-comms`
- **Component:** `cmd/comms`
- **Kind:** `cli`
- **Audience:** `dev`
- **Maturity:** `dev_only`
- **Lifecycle status:** `lower_level_direct_proof_transition_candidate`
- **Introduced in:** `v0.4.4`
- **Source path:** `carbonstack-comms/internal/app/openmls_runtime.go`
- **Validation surface:** dev-runtime-openmls and dev-runtime-openmls-wrappers
- **Front README candidate:** `true`

**What it does:** Direct dev/pre-alpha OpenMLS application-message inbox/open proof path through Cypher inbox and sidecar message-open.

**Why it exists:** Lower-level direct OpenMLS/Cypher implementation proof surface behind the opinionated message-inbox-dev wrapper; useful for debugging, tests, and implementation visibility.

- **Required flags:** Not recorded in registry.
- **Optional flags:** Not recorded in registry.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - dev/pre-alpha only
  - not mature inbox UX
  - not production E2EE claim
  - not hostile-server safety
  - not metadata privacy
  - not verified identity
  - not durable local receive storage
  - does not perform Relay onboarding
  - no ack on fetch

### `comms.openmls-send-dev`

- **Command:** `go run ./cmd/comms openmls-send-dev`
- **Repo:** `carbonstack-comms`
- **Component:** `cmd/comms`
- **Kind:** `cli`
- **Audience:** `dev`
- **Maturity:** `dev_only`
- **Lifecycle status:** `lower_level_direct_proof_transition_candidate`
- **Introduced in:** `v0.4.3`
- **Source path:** `carbonstack-comms/internal/app/openmls_runtime.go`
- **Validation surface:** dev-runtime-openmls and dev-runtime-openmls-wrappers
- **Front README candidate:** `true`

**What it does:** Direct dev/pre-alpha OpenMLS application-message send proof path through sidecar message-protect and Cypher.

**Why it exists:** Lower-level direct OpenMLS/Cypher implementation proof surface behind the opinionated message-send-dev wrapper; useful for debugging, tests, and implementation visibility.

- **Required flags:** Not recorded in registry.
- **Optional flags:** Not recorded in registry.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - dev/pre-alpha only
  - not mature send UX
  - not production E2EE claim
  - not hostile-server safety
  - not metadata privacy
  - not verified identity
  - does not perform Relay onboarding

## Relay onboarding and artifact commands

Entries in this section: **9**

### `comms.openmls-relay-add-member-dev`

- **Command:** `go run ./cmd/comms openmls-relay-add-member-dev`
- **Repo:** `carbonstack-comms`
- **Component:** `cmd/comms`
- **Kind:** `cli`
- **Audience:** `dev`
- **Maturity:** `dev_only`
- **Lifecycle status:** `relay_onboarding_artifact_bridge`
- **Introduced in:** `Not recorded in registry.`
- **Source path:** `carbonstack-comms/internal/app/openmls_relay_dev.go`
- **Validation surface:** relay-openmls-join-dev and Comms package tests
- **Front README candidate:** `false`

**What it does:** Consume a Relay Space KeyPackage, run sidecar conversation-add-member, and submit the produced Welcome through Relay Space.

**Why it exists:** Bridges Relay Space KeyPackage transport to OpenMLS sidecar add-member and Welcome transport in the dev/pre-alpha join scaffold.

- **Required flags:**
  - `--relay-space` — Relay Space ID containing a queued KeyPackage envelope.
  - `--sidecar-device-label` — OpenMLS sidecar device label for the adding member.
  - `--conversation` — OpenMLS sidecar conversation label that receives the new member.
- **Optional flags:**
  - `--state` — local Comms state path; defaults to the normal Comms state path.
  - `--sidecar-dir` — OpenMLS sidecar directory; defaults to the repo-local sidecar path.
  - `--welcome-to-device` — explicit recipient device ID for the produced Welcome; defaults to KeyPackage envelope sender_device_id.
  - `--client-created-at` — test/dev timestamp override for deterministic Welcome envelope metadata.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - not join automation
  - not identity verification
  - not local-backbone
  - not production membership UX
  - does not ack KeyPackage or Welcome

### `comms.openmls-relay-join-dev`

- **Command:** `go run ./cmd/comms openmls-relay-join-dev`
- **Repo:** `carbonstack-comms`
- **Component:** `cmd/comms`
- **Kind:** `cli`
- **Audience:** `dev`
- **Maturity:** `dev_only`
- **Lifecycle status:** `relay_onboarding_join_transition_candidate`
- **Introduced in:** `Not recorded in registry.`
- **Source path:** `carbonstack-comms/internal/app/openmls_relay_dev.go`
- **Validation surface:** relay-openmls-join-dev and Comms package tests
- **Front README candidate:** `false`

**What it does:** Consume a Relay Space Welcome, run sidecar conversation-join, and optionally ack only after join succeeds.

**Why it exists:** Completes the dev/pre-alpha Relay Space OpenMLS join scaffold while preserving explicit nonclaims.

- **Required flags:**
  - `--relay-space` — Relay Space ID containing a queued Welcome envelope.
  - `--sidecar-device-label` — OpenMLS sidecar device label for the joining member.
  - `--conversation` — OpenMLS sidecar conversation label to join/load.
- **Optional flags:**
  - `--state` — local Comms state path; defaults to the normal Comms state path.
  - `--sidecar-dir` — OpenMLS sidecar directory; defaults to the repo-local sidecar path.
  - `--ack-after-join` — ack the Welcome envelope only after sidecar conversation-join succeeds. Boundary: Acknowledges delivery/local processing only; not verification or production receive proof.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - not identity verification
  - not local-backbone
  - not production UX
  - not hostile-server safety
  - not metadata privacy

### `comms.openmls-relay-keypackage-consume-dev`

- **Command:** `go run ./cmd/comms openmls-relay-keypackage-consume-dev`
- **Repo:** `carbonstack-comms`
- **Component:** `cmd/comms`
- **Kind:** `cli`
- **Audience:** `dev`
- **Maturity:** `dev_only`
- **Lifecycle status:** `relay_onboarding_artifact_consume`
- **Introduced in:** `Not recorded in registry.`
- **Source path:** `carbonstack-comms/internal/app/openmls_keypackage_consume_dev.go`
- **Validation surface:** keypackage-consume-dev and Comms package tests
- **Front README candidate:** `false`

**What it does:** Persist one queued Relay Space KeyPackage artifact locally and ACK only after the receipt is durable.

**Why it exists:** Closes Gate B5d delivery consume/receipt semantics without starting add-member or Welcome lifecycle.

- **Required flags:**
  - `--relay-space` — Relay Space ID containing the queued KeyPackage envelope.
  - `--envelope-id` — exact KeyPackage envelope to consume into local receipt state.
- **Optional flags:**
  - `--state` — local Comms state path; defaults to the normal Comms state path.
  - `--receipt-root` — local receipt root; defaults beside the Comms state file.
  - `--expected-payload-sha256` — optional decoded payload SHA-256 hex guard.
  - `--expected-key-package-ref` — optional operator expectation recorded in the receipt; not cryptographically verified by this command.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - not add-member
  - not Welcome lifecycle
  - not identity verification
  - not trust promotion
  - not public KeyPackage directory
  - not production key distribution
  - not B6

### `comms.openmls-relay-keypackage-inbox-dev`

- **Command:** `go run ./cmd/comms openmls-relay-keypackage-inbox-dev`
- **Repo:** `carbonstack-comms`
- **Component:** `cmd/comms`
- **Kind:** `cli`
- **Audience:** `dev`
- **Maturity:** `dev_only`
- **Lifecycle status:** `relay_onboarding_artifact_transport`
- **Introduced in:** `Not recorded in registry.`
- **Source path:** `carbonstack-comms/internal/app/openmls_relay_dev.go`
- **Validation surface:** relay-openmls-join-dev and Comms package tests
- **Front README candidate:** `false`

**What it does:** Fetch Relay Space scoped OpenMLS KeyPackage envelopes and write KeyPackage artifacts without ack.

**Why it exists:** Exposes the dev KeyPackage inbox/write step used before Relay/OpenMLS add-member.

- **Required flags:**
  - `--relay-space` — Relay Space ID to inspect for queued KeyPackage envelopes.
- **Optional flags:**
  - `--state` — local Comms state path; defaults to the normal Comms state path.
  - `--limit` — maximum KeyPackage envelopes to write; defaults to one.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - no add-member
  - no ack
  - no trust mutation
  - no verification
  - not production inbox UX

### `comms.openmls-relay-keypackage-publish-dev`

- **Command:** `go run ./cmd/comms openmls-relay-keypackage-publish-dev`
- **Repo:** `carbonstack-comms`
- **Component:** `cmd/comms`
- **Kind:** `comms-command`
- **Audience:** `dev`
- **Maturity:** `dev_only`
- **Lifecycle status:** `active`
- **Introduced in:** `v0.7.6`
- **Source path:** `carbonstack-comms/internal/app/openmls_keypackage_publication_dev.go`
- **Validation surface:** explicit-b5b-generation-selection-b5a-inspection-dedicated-relay-publication-created-or-replay
- **Front README candidate:** `false`

**What it does:** Inspect and publish one existing active KeyPackage generation through the B5c route.

**Why it exists:** Expose deterministic Relay publication without silently generating, consuming, acknowledging, or promoting identity.

- **Required flags:**
  - `--relay-space` — Relay Space destination.
  - `--to-device` — Recipient Cypher device.
  - `--sidecar-device-label` — Local sidecar owner label.
  - `--generation-id` — Existing active B5b generation.
- **Optional flags:**
  - `--state` — Explicit Comms state path.
  - `--sidecar-dir` — Explicit active OpenMLS sidecar directory.
  - `--client-created-at` — Test/dev metadata override; not publication identity.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - not implicit KeyPackage generation
  - not retired-generation publication
  - not KeyPackage consumption
  - not KeyPackage ACK
  - not add-member
  - not Welcome submission or join
  - not account or device identity verification
  - not Relay Space membership verification by the client
  - not trust promotion
  - not secure enrollment
  - not public UX stability
  - not production key distribution
  - not local-backbone
  - not deployment
  - not audit or certification

### `comms.openmls-relay-keypackage-submit-dev`

- **Command:** `go run ./cmd/comms openmls-relay-keypackage-submit-dev`
- **Repo:** `carbonstack-comms`
- **Component:** `cmd/comms`
- **Kind:** `cli`
- **Audience:** `dev`
- **Maturity:** `dev_only`
- **Lifecycle status:** `relay_onboarding_artifact_transport`
- **Introduced in:** `Not recorded in registry.`
- **Source path:** `carbonstack-comms/internal/app/openmls_relay_dev.go`
- **Validation surface:** relay-openmls-join-dev and Comms package tests
- **Front README candidate:** `false`

**What it does:** Export a dev OpenMLS KeyPackage/public bundle through the sidecar and submit it as a Relay Space scoped KeyPackage envelope.

**Why it exists:** Starts the Relay Space OpenMLS join scaffold by transporting a recipient KeyPackage artifact through Cypher without claiming identity verification or mature join UX.

- **Required flags:**
  - `--relay-space` — Relay Space ID that scopes the KeyPackage envelope.
  - `--to-device` — recipient Cypher device ID for the KeyPackage envelope.
  - `--sidecar-device-label` — OpenMLS sidecar device label used for public-bundle export.
- **Optional flags:**
  - `--state` — local Comms state path; defaults to the normal Comms state path.
  - `--sidecar-dir` — OpenMLS sidecar directory; defaults to the repo-local sidecar path.
  - `--client-created-at` — test/dev timestamp override for deterministic envelope metadata.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - not join automation
  - not identity verification
  - not production key distribution UX
  - not local-backbone

### `comms.openmls-relay-welcome-consume-dev`

- **Command:** `go run ./cmd/comms openmls-relay-welcome-consume-dev`
- **Repo:** `carbonstack-comms`
- **Component:** `cmd/comms`
- **Kind:** `cli`
- **Audience:** `dev`
- **Maturity:** `dev_only`
- **Lifecycle status:** `relay_onboarding_welcome_consume`
- **Introduced in:** `Not recorded in registry.`
- **Source path:** `carbonstack-comms/internal/app/openmls_welcome_lifecycle_dev.go`
- **Validation surface:** welcome-lifecycle-dev and Comms package tests
- **Front README candidate:** `false`

**What it does:** Persist one queued Relay Space Welcome artifact locally, run sidecar join, and ACK only after persisted join evidence.

**Why it exists:** Closes Gate B6 Welcome consume/join/ACK-after-join semantics without claiming identity verification or trust promotion.

- **Required flags:**
  - `--relay-space` — Relay Space ID containing the queued Welcome envelope.
  - `--envelope-id` — exact Welcome envelope to consume and join.
  - `--sidecar-device-label` — OpenMLS sidecar device label for the joining member.
  - `--conversation` — OpenMLS sidecar conversation label to join/load.
- **Optional flags:**
  - `--state` — local Comms state path.
  - `--sidecar-dir` — OpenMLS sidecar directory.
  - `--receipt-root` — local Welcome receipt root.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - not verified identity
  - not trust promotion
  - not secure enrollment
  - not production E2EE
  - not B7 Cypher/MLS reconciliation
  - not public directory safety

### `comms.openmls-relay-welcome-inbox-dev`

- **Command:** `go run ./cmd/comms openmls-relay-welcome-inbox-dev`
- **Repo:** `carbonstack-comms`
- **Component:** `cmd/comms`
- **Kind:** `cli`
- **Audience:** `dev`
- **Maturity:** `dev_only`
- **Lifecycle status:** `relay_onboarding_artifact_transport`
- **Introduced in:** `Not recorded in registry.`
- **Source path:** `carbonstack-comms/internal/app/openmls_relay_dev.go`
- **Validation surface:** relay-openmls-join-dev and Comms package tests
- **Front README candidate:** `false`

**What it does:** Fetch Relay Space scoped OpenMLS Welcome envelopes and write Welcome artifacts without ack.

**Why it exists:** Exposes the dev Welcome inbox/write step used before Relay/OpenMLS join.

- **Required flags:**
  - `--relay-space` — Relay Space ID to inspect for queued Welcome envelopes.
- **Optional flags:**
  - `--state` — local Comms state path; defaults to the normal Comms state path.
  - `--limit` — maximum Welcome envelopes to write; defaults to one.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - no conversation-join
  - no ack
  - no trust mutation
  - no verification
  - not production inbox UX

### `comms.openmls-relay-welcome-submit-dev`

- **Command:** `go run ./cmd/comms openmls-relay-welcome-submit-dev`
- **Repo:** `carbonstack-comms`
- **Component:** `cmd/comms`
- **Kind:** `cli`
- **Audience:** `dev`
- **Maturity:** `dev_only`
- **Lifecycle status:** `relay_onboarding_artifact_transport`
- **Introduced in:** `Not recorded in registry.`
- **Source path:** `carbonstack-comms/internal/app/openmls_relay_dev.go`
- **Validation surface:** relay-openmls-join-dev and Comms package tests
- **Front README candidate:** `false`

**What it does:** Submit an existing OpenMLS Welcome artifact as a Relay Space scoped Welcome envelope.

**Why it exists:** Provides the explicit Welcome transport step used by Relay/OpenMLS join scaffolding.

- **Required flags:**
  - `--relay-space` — Relay Space ID that scopes the Welcome envelope.
  - `--to-device` — recipient Cypher device ID for the Welcome envelope.
  - `--welcome` — local OpenMLS Welcome artifact path to submit.
- **Optional flags:**
  - `--state` — local Comms state path; defaults to the normal Comms state path.
  - `--client-created-at` — test/dev timestamp override for deterministic envelope metadata.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - not join automation
  - not identity verification
  - not production membership UX
  - not local-backbone

## OpenMLS bootstrap, identity, and conversation commands

Entries in this section: **12**

### `comms.openmls-bundle-export-dev`

- **Command:** `go run ./cmd/comms openmls-bundle-export-dev`
- **Repo:** `carbonstack-comms`
- **Component:** `cmd/comms`
- **Kind:** `cli`
- **Audience:** `dev`
- **Maturity:** `dev_only`
- **Introduced in:** `v0.4.14`
- **Source path:** `carbonstack-comms/internal/app/openmls_bootstrap.go`
- **Validation surface:** dev-runtime-openmls-wrappers
- **Front README candidate:** `false`

**What it does:** Export a dev OpenMLS KeyPackage bundle and optional artifact path.

**Why it exists:** Wrapper step for member onboarding in dev runtime proofs.

- **Required flags:** Not recorded in registry.
- **Optional flags:** Not recorded in registry.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - not production identity/key distribution UX

### `comms.openmls-conversation-add-member-dev`

- **Command:** `go run ./cmd/comms openmls-conversation-add-member-dev`
- **Repo:** `carbonstack-comms`
- **Component:** `cmd/comms`
- **Kind:** `cli`
- **Audience:** `dev`
- **Maturity:** `dev_only`
- **Introduced in:** `v0.4.15`
- **Source path:** `carbonstack-comms/internal/app/openmls_bootstrap.go`
- **Validation surface:** dev-runtime-openmls-wrappers
- **Front README candidate:** `false`

**What it does:** Add a member KeyPackage to a dev OpenMLS conversation and write a Welcome artifact.

**Why it exists:** Wrapper bootstrap step for Welcome generation in wrapper-based runtime smoke.

- **Required flags:** Not recorded in registry.
- **Optional flags:** Not recorded in registry.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - not production membership UX
  - not Relay Space join UX

### `comms.openmls-conversation-create-dev`

- **Command:** `go run ./cmd/comms openmls-conversation-create-dev`
- **Repo:** `carbonstack-comms`
- **Component:** `cmd/comms`
- **Kind:** `cli`
- **Audience:** `dev`
- **Maturity:** `dev_only`
- **Introduced in:** `v0.4.14`
- **Source path:** `carbonstack-comms/internal/app/openmls_bootstrap.go`
- **Validation surface:** dev-runtime-openmls-wrappers
- **Front README candidate:** `false`

**What it does:** Create a dev-local OpenMLS conversation for an explicit sidecar label and conversation label.

**Why it exists:** Wrapper bootstrap step before add-member/join and message proof.

- **Required flags:** Not recorded in registry.
- **Optional flags:** Not recorded in registry.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - not Relay Space join UX

### `comms.openmls-conversation-join-dev`

- **Command:** `go run ./cmd/comms openmls-conversation-join-dev`
- **Repo:** `carbonstack-comms`
- **Component:** `cmd/comms`
- **Kind:** `cli`
- **Audience:** `dev`
- **Maturity:** `dev_only`
- **Introduced in:** `v0.4.15`
- **Source path:** `carbonstack-comms/internal/app/openmls_bootstrap.go`
- **Validation surface:** dev-runtime-openmls-wrappers
- **Front README candidate:** `false`

**What it does:** Consume a Welcome artifact and join a dev OpenMLS conversation.

**Why it exists:** Wrapper bootstrap step for completing Alice/Bob wrapper smoke setup.

- **Required flags:** Not recorded in registry.
- **Optional flags:** Not recorded in registry.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - not production membership UX
  - not Relay Space join UX

### `comms.openmls-conversation-load-check-dev`

- **Command:** `go run ./cmd/comms openmls-conversation-load-check-dev`
- **Repo:** `carbonstack-comms`
- **Component:** `cmd/comms`
- **Kind:** `cli`
- **Audience:** `dev`
- **Maturity:** `dev_only`
- **Introduced in:** `v0.4.14`
- **Source path:** `carbonstack-comms/internal/app/openmls_bootstrap.go`
- **Validation surface:** dev-runtime-openmls-wrappers
- **Front README candidate:** `false`

**What it does:** Check that a dev-local sidecar conversation can be loaded and reports group_reloadable.

**Why it exists:** Bootstrap state sanity check for wrapper proof chains.

- **Required flags:** Not recorded in registry.
- **Optional flags:** Not recorded in registry.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - dev-local sidecar state check only
  - not production membership UX
  - not identity verification
  - not secure recovery

### `comms.openmls-cypher-mls-mismatch-inspect-dev`

- **Command:** `go run ./cmd/comms openmls-cypher-mls-mismatch-inspect-dev`
- **Repo:** `carbonstack-comms`
- **Component:** `cmd/comms`
- **Kind:** `cli`
- **Audience:** `dev`
- **Maturity:** `dev_only`
- **Lifecycle status:** `relay_state_mismatch_inspection`
- **Introduced in:** `Not recorded in registry.`
- **Source path:** `carbonstack-comms/internal/app/openmls_cypher_mls_mismatch_dev.go`
- **Validation surface:** cypher-mls-mismatch-dev and Comms package tests
- **Front README candidate:** `false`

**What it does:** Classify Cypher Relay Space membership versus local MLS/receipt state and refuse unsafe mismatches.

**Why it exists:** Closes Gate B7 as a bounded mismatch inspection/refusal leaf without silent repair, trust promotion, or identity binding.

- **Required flags:**
  - `--relay-space` — Relay Space ID under inspection.
  - `--cypher-member-state` — explicit Cypher Relay Space member-state snapshot.
- **Optional flags:**
  - `--state` — local Comms state path.
  - `--mls-group-state` — local MLS group state snapshot or auto.
  - `--sidecar-dir` — OpenMLS sidecar directory.
  - `--sidecar-device-label` — OpenMLS sidecar device label.
  - `--conversation` — OpenMLS conversation label.
  - `--conversation-state` — explicit OpenMLS conversation state path.
  - `--keypackage-receipt` — optional B5d KeyPackage receipt manifest.
  - `--welcome-receipt` — optional B6 Welcome receipt manifest.
  - `--allow-refusal-exit-zero` — print refusal evidence but return exit 0 for reporting harnesses.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - not silent repair
  - not silent rejoin
  - not trust promotion
  - not verified identity
  - not Cypher/MLS reconciliation
  - not B8 workflow engine
  - not B9 Gate B closure
  - not production E2EE

### `comms.openmls-identity-create-dev`

- **Command:** `go run ./cmd/comms openmls-identity-create-dev`
- **Repo:** `carbonstack-comms`
- **Component:** `cmd/comms`
- **Kind:** `cli`
- **Audience:** `dev`
- **Maturity:** `dev_only`
- **Introduced in:** `v0.4.13`
- **Source path:** `carbonstack-comms/internal/app/openmls_bootstrap.go`
- **Validation surface:** dev-runtime-openmls-wrappers
- **Front README candidate:** `false`

**What it does:** Create a dev-local OpenMLS sidecar identity with an explicit sidecar device label.

**Why it exists:** Bootstrap helper for wrapper-based runtime smoke.

- **Required flags:** Not recorded in registry.
- **Optional flags:** Not recorded in registry.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - not production identity UX
  - does not mutate Comms trust state

### `comms.openmls-identity-status-dev`

- **Command:** `go run ./cmd/comms openmls-identity-status-dev`
- **Repo:** `carbonstack-comms`
- **Component:** `cmd/comms`
- **Kind:** `cli`
- **Audience:** `dev`
- **Maturity:** `dev_only`
- **Introduced in:** `v0.4.13`
- **Source path:** `carbonstack-comms/internal/app/openmls_bootstrap.go`
- **Validation surface:** dev-runtime-openmls-wrappers
- **Front README candidate:** `false`

**What it does:** Check dev-local OpenMLS sidecar identity status for an explicit sidecar device label.

**Why it exists:** Bootstrap inspection helper for wrapper-based runtime smoke.

- **Required flags:** Not recorded in registry.
- **Optional flags:** Not recorded in registry.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - not production identity UX
  - does not mutate Comms trust state

### `comms.openmls-keypackage-generate-dev`

- **Command:** `go run ./cmd/comms openmls-keypackage-generate-dev`
- **Repo:** `carbonstack-comms`
- **Component:** `cmd/comms`
- **Kind:** `comms-command`
- **Audience:** `dev`
- **Maturity:** `dev_only`
- **Lifecycle status:** `active`
- **Introduced in:** `v0.7.5`
- **Source path:** `carbonstack-comms/internal/app/openmls_keypackage_rotation_dev.go`
- **Validation surface:** stable-dev-wrapper-for-repeatable-keypackage-generation
- **Front README candidate:** `false`

**What it does:** Generate or replay a persistent local KeyPackage generation through Comms.

**Why it exists:** Expose the B5b generation primitive without promoting it to public or production UX.

- **Required flags:**
  - `--sidecar-device-label` — Local sidecar device label.
  - `--request-id` — Device-local idempotency identity.
- **Optional flags:**
  - `--sidecar-dir` — Explicit active OpenMLS sidecar directory.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - not Relay publication
  - not KeyPackage consumption
  - not Welcome lifecycle
  - not identity verification
  - not trust promotion
  - not public UX stability
  - not production onboarding
  - not local-backbone
  - not deployment
  - not audit or certification

### `comms.openmls-keypackage-inspect-dev`

- **Command:** `go run ./cmd/comms openmls-keypackage-inspect-dev`
- **Repo:** `carbonstack-comms`
- **Component:** `cmd/comms`
- **Kind:** `comms-command`
- **Audience:** `dev`
- **Maturity:** `dev_only`
- **Lifecycle status:** `active`
- **Introduced in:** `v0.7.5`
- **Source path:** `carbonstack-comms/internal/app/openmls_keypackage_inspect_dev.go`
- **Validation surface:** stable-dev-wrapper-for-read-only-keypackage-inspection
- **Front README candidate:** `false`

**What it does:** Inspect a serialized OpenMLS KeyPackage through a stable Comms dev wrapper.

**Why it exists:** Expose the Gate B5a inspection result without promoting the internal sidecar command to production UX.

- **Required flags:**
  - `--sidecar-device-label` — Local sidecar device label used for ownership evidence.
  - `--keypackage` — Serialized OpenMLS KeyPackage artifact path to inspect.
- **Optional flags:**
  - `--sidecar-dir` — Explicit active OpenMLS sidecar directory.
  - `--generation-manifest` — Optional immutable B5b generation manifest used as local ownership evidence.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - not repeatable KeyPackage generation
  - not KeyPackage rotation
  - not Relay publication
  - not KeyPackage consumption
  - not Welcome lifecycle
  - not account identity verification
  - not device identity verification
  - not Relay Space membership verification
  - not trust promotion
  - not secure enrollment
  - not public UX stability
  - not production E2EE
  - not local-backbone
  - not deployment
  - not audit or certification

### `comms.openmls-keypackage-inventory-dev`

- **Command:** `go run ./cmd/comms openmls-keypackage-inventory-dev`
- **Repo:** `carbonstack-comms`
- **Component:** `cmd/comms`
- **Kind:** `comms-command`
- **Audience:** `dev`
- **Maturity:** `dev_only`
- **Lifecycle status:** `active`
- **Introduced in:** `v0.7.5`
- **Source path:** `carbonstack-comms/internal/app/openmls_keypackage_rotation_dev.go`
- **Validation surface:** stable-dev-wrapper-for-read-only-keypackage-inventory
- **Front README candidate:** `false`

**What it does:** Inspect persistent local KeyPackage generations through Comms.

**Why it exists:** Expose the B5b inventory without implicit repair or Relay behavior.

- **Required flags:**
  - `--sidecar-device-label` — Local sidecar device label.
- **Optional flags:**
  - `--sidecar-dir` — Explicit active OpenMLS sidecar directory.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - not mutation or implicit repair
  - not Relay publication
  - not KeyPackage consumption
  - not Welcome lifecycle
  - not identity verification
  - not trust promotion
  - not public UX stability
  - not production onboarding
  - not local-backbone
  - not deployment
  - not audit or certification

### `comms.openmls-keypackage-retire-dev`

- **Command:** `go run ./cmd/comms openmls-keypackage-retire-dev`
- **Repo:** `carbonstack-comms`
- **Component:** `cmd/comms`
- **Kind:** `comms-command`
- **Audience:** `dev`
- **Maturity:** `dev_only`
- **Lifecycle status:** `active`
- **Introduced in:** `v0.7.5`
- **Source path:** `carbonstack-comms/internal/app/openmls_keypackage_rotation_dev.go`
- **Validation surface:** stable-dev-wrapper-for-metadata-only-keypackage-retirement
- **Front README candidate:** `false`

**What it does:** Retire a non-current local KeyPackage generation through Comms.

**Why it exists:** Expose explicit B5b retirement without deletion, Relay revocation, or trust mutation.

- **Required flags:**
  - `--sidecar-device-label` — Local sidecar device label.
  - `--generation-id` — Non-current generation to retire.
- **Optional flags:**
  - `--sidecar-dir` — Explicit active OpenMLS sidecar directory.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - not deletion or garbage collection
  - not Relay revocation
  - not KeyPackage consumption
  - not Welcome lifecycle
  - not trust revocation
  - not public UX stability
  - not production onboarding
  - not local-backbone
  - not deployment
  - not audit or certification

## Comms state, account, device, and trust commands

Entries in this section: **11**

### `comms.basic-local-trust-accept-dev`

- **Command:** `go run ./cmd/comms basic-local-trust-accept-dev`
- **Repo:** `carbonstack-comms`
- **Component:** `cmd/comms`
- **Kind:** `dev-command`
- **Audience:** `dev`
- **Maturity:** `dev_only`
- **Lifecycle status:** `active`
- **Introduced in:** `Not recorded in registry.`
- **Source path:** `carbonstack-comms/internal/app/basic_local_trust_posture_dev.go`
- **Validation surface:** gate-f-f5-basic-local-trust-candidate-posture
- **Front README candidate:** `false`

**What it does:** Record an explicit local manual trust candidate acceptance event with loud nonclaims.

**Why it exists:** Provides a dev/pre-alpha local evidence path for manual candidate acceptance without verified identity, automatic trust promotion, or cryptographic identity binding.

- **Required flags:**
  - `--accept-candidate` — explicit operator confirmation that a local candidate acceptance event should be written.
  - `--reason` — operator reason for local manual candidate acceptance.
- **Optional flags:** Not recorded in registry.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - not verified identity
  - not full trust promotion
  - not secure enrollment
  - not server-hostile identity replacement proof
  - not real-world person verification
  - not cryptographic binding across Cypher Comms and OpenMLS identities
  - not automatic trust promotion
  - not trust from Relay membership
  - not trust from successful Welcome or MLS join
  - not package-runtime candidate
  - not release readiness
  - not production E2EE

### `comms.basic-local-trust-posture-dev`

- **Command:** `go run ./cmd/comms basic-local-trust-posture-dev`
- **Repo:** `carbonstack-comms`
- **Component:** `cmd/comms`
- **Kind:** `dev-command`
- **Audience:** `dev`
- **Maturity:** `dev_only`
- **Lifecycle status:** `active`
- **Introduced in:** `Not recorded in registry.`
- **Source path:** `carbonstack-comms/internal/app/basic_local_trust_posture_dev.go`
- **Validation surface:** gate-f-f5-basic-local-trust-candidate-posture
- **Front README candidate:** `false`

**What it does:** Print a basic local trust posture report across Cypher account/device, Comms local fingerprint, OpenMLS sidecar/KeyPackage, and Relay Space routing evidence.

**Why it exists:** Makes identity-domain separation explicit before v0.8.0 without claiming verified identity, trust promotion, secure enrollment, or cryptographic binding.

- **Required flags:** Not recorded in registry.
- **Optional flags:**
  - `--report` — write the posture report to a JSON file.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - not verified identity
  - not full trust promotion
  - not secure enrollment
  - not server-hostile identity replacement proof
  - not real-world person verification
  - not cryptographic binding across Cypher Comms and OpenMLS identities
  - not automatic trust promotion
  - not trust from Relay membership
  - not trust from successful Welcome or MLS join
  - not package-runtime candidate
  - not release readiness
  - not production E2EE

### `comms.claim-invite`

- **Command:** `go run ./cmd/comms claim-invite`
- **Repo:** `carbonstack-comms`
- **Component:** `cmd/comms`
- **Kind:** `cli`
- **Audience:** `dev`
- **Maturity:** `experimental`
- **Introduced in:** `Not recorded in registry.`
- **Source path:** `carbonstack-comms/internal/app/commands.go`
- **Validation surface:** runtime smoke scripts
- **Front README candidate:** `false`

**What it does:** Claim an invite and bind local Comms state to a Cypher account.

**Why it exists:** Establishes account identity in local/dev Comms flows.

- **Required flags:** Not recorded in registry.
- **Optional flags:** Not recorded in registry.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - dev/pre-alpha invite claim only
  - not secure enrollment
  - not identity verification
  - not production account recovery
  - not deployment readiness

### `comms.fingerprint`

- **Command:** `go run ./cmd/comms fingerprint`
- **Repo:** `carbonstack-comms`
- **Component:** `cmd/comms`
- **Kind:** `cli`
- **Audience:** `dev`
- **Maturity:** `experimental`
- **Introduced in:** `Not recorded in registry.`
- **Source path:** `carbonstack-comms/internal/app/commands.go`
- **Validation surface:** trust lifecycle tests
- **Front README candidate:** `false`

**What it does:** Print or inspect trust/fingerprint information for local device state.

**Why it exists:** Supports trust lifecycle development before mature UX.

- **Required flags:** Not recorded in registry.
- **Optional flags:** Not recorded in registry.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - dev/operator trust inspection helper only
  - does not verify identity by itself
  - not mature verification ceremony
  - not production trust UX

### `comms.init`

- **Command:** `go run ./cmd/comms init`
- **Repo:** `carbonstack-comms`
- **Component:** `cmd/comms`
- **Kind:** `cli`
- **Audience:** `dev`
- **Maturity:** `experimental`
- **Introduced in:** `Not recorded in registry.`
- **Source path:** `carbonstack-comms/internal/app/commands.go`
- **Validation surface:** Comms package tests and runtime smoke scripts
- **Front README candidate:** `false`

**What it does:** Initialize a local Comms state file for a server URL.

**Why it exists:** Creates local client state for dev/test Comms workflows.

- **Required flags:** Not recorded in registry.
- **Optional flags:** Not recorded in registry.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - dev/pre-alpha local state initialization only
  - not secure enrollment
  - not production account setup
  - not production key storage
  - not deployment readiness

### `comms.list-devices`

- **Command:** `go run ./cmd/comms list-devices`
- **Repo:** `carbonstack-comms`
- **Component:** `cmd/comms`
- **Kind:** `cli`
- **Audience:** `dev`
- **Maturity:** `experimental`
- **Introduced in:** `Not recorded in registry.`
- **Source path:** `carbonstack-comms/internal/app/commands.go`
- **Validation surface:** Comms tests
- **Front README candidate:** `false`

**What it does:** List devices known through the current Comms/Cypher account context.

**Why it exists:** Debug/operator helper for device visibility.

- **Required flags:** Not recorded in registry.
- **Optional flags:** Not recorded in registry.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - dev/operator visibility helper only
  - not device verification
  - not authorization proof
  - not production device-management UX

### `comms.register-device`

- **Command:** `go run ./cmd/comms register-device`
- **Repo:** `carbonstack-comms`
- **Component:** `cmd/comms`
- **Kind:** `cli`
- **Audience:** `dev`
- **Maturity:** `experimental`
- **Introduced in:** `Not recorded in registry.`
- **Source path:** `carbonstack-comms/internal/app/commands.go`
- **Validation surface:** runtime smoke scripts
- **Front README candidate:** `false`

**What it does:** Register a device for the current Comms state/account.

**Why it exists:** Provides device IDs used by Cypher envelope routing and smoke tests.

- **Required flags:** Not recorded in registry.
- **Optional flags:** Not recorded in registry.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - dev/pre-alpha device registration only
  - not verified device ceremony
  - not production device enrollment
  - not secure key storage

### `comms.revoke-device`

- **Command:** `go run ./cmd/comms revoke-device`
- **Repo:** `carbonstack-comms`
- **Component:** `cmd/comms`
- **Kind:** `cli`
- **Audience:** `dev`
- **Maturity:** `experimental`
- **Introduced in:** `Not recorded in registry.`
- **Source path:** `carbonstack-comms/internal/app/commands.go`
- **Validation surface:** trust lifecycle tests
- **Front README candidate:** `false`

**What it does:** Revoke a device in the local trust model.

**Why it exists:** Supports device lifecycle and trust event testing.

- **Required flags:** Not recorded in registry.
- **Optional flags:** Not recorded in registry.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - dev/pre-alpha trust-state mutation helper
  - not production revocation UX
  - not account recovery
  - not secure re-enrollment

### `comms.trust-history`

- **Command:** `go run ./cmd/comms trust-history`
- **Repo:** `carbonstack-comms`
- **Component:** `cmd/comms`
- **Kind:** `cli`
- **Audience:** `dev`
- **Maturity:** `experimental`
- **Introduced in:** `Not recorded in registry.`
- **Source path:** `carbonstack-comms/internal/app/commands.go`
- **Validation surface:** trust lifecycle tests
- **Front README candidate:** `false`

**What it does:** Show trust event history for debugging trust-state changes.

**Why it exists:** Makes trust lifecycle evidence inspectable.

- **Required flags:** Not recorded in registry.
- **Optional flags:** Not recorded in registry.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - dev/operator trust history inspection only
  - not identity verification
  - not production audit log
  - not secure trust ceremony

### `comms.trust-list`

- **Command:** `go run ./cmd/comms trust-list`
- **Repo:** `carbonstack-comms`
- **Component:** `cmd/comms`
- **Kind:** `cli`
- **Audience:** `dev`
- **Maturity:** `experimental`
- **Introduced in:** `Not recorded in registry.`
- **Source path:** `carbonstack-comms/internal/app/commands.go`
- **Validation surface:** trust lifecycle tests
- **Front README candidate:** `false`

**What it does:** List known trust records from local Comms state.

**Why it exists:** Debug helper for trust-state development.

- **Required flags:** Not recorded in registry.
- **Optional flags:** Not recorded in registry.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - dev/operator trust-state inspection only
  - not identity verification
  - not production trust UX
  - does not prove hostile-server safety

### `comms.verify-device`

- **Command:** `go run ./cmd/comms verify-device`
- **Repo:** `carbonstack-comms`
- **Component:** `cmd/comms`
- **Kind:** `cli`
- **Audience:** `dev`
- **Maturity:** `experimental`
- **Introduced in:** `Not recorded in registry.`
- **Source path:** `carbonstack-comms/internal/app/commands.go`
- **Validation surface:** trust lifecycle tests
- **Front README candidate:** `false`

**What it does:** Mark or evaluate a device as verified in the dev trust model.

**Why it exists:** Supports trust-state test/dev flows.

- **Required flags:** Not recorded in registry.
- **Optional flags:** Not recorded in registry.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - not production trust UX

## Legacy stub-era and continuity commands

Entries in this section: **3**

### `comms.ack`

- **Command:** `go run ./cmd/comms ack`
- **Repo:** `carbonstack-comms`
- **Component:** `cmd/comms`
- **Kind:** `cli`
- **Audience:** `legacy`
- **Maturity:** `legacy`
- **Lifecycle status:** `legacy_stub_explicit_opt_in`
- **Introduced in:** `Not recorded in registry.`
- **Source path:** `carbonstack-comms/internal/app/commands.go`
- **Validation surface:** Comms package tests and runtime smoke scripts
- **Front README candidate:** `false`

**What it does:** Explicitly gated standalone legacy ack helper for older Cypher envelope workflows.

**Why it exists:** Preserved for explicit ack workflows and continuity; current OpenMLS receive tests should prefer openmls-inbox-dev --ack or scoped Relay Space ack helpers.

- **Required flags:**
  - `--allow-legacy-stub` — explicit opt-in required to run the standalone legacy ack helper. Boundary: gates the old standalone CLI surface without changing OpenMLS inbox --ack or Relay Space scoped ack helpers.
  - `--envelope` — envelope ID to acknowledge through the older standalone CLI surface.
- **Optional flags:**
  - `--state` — local Comms state path; defaults to the normal Comms state path.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - not a standalone secure receive proof
  - requires explicit legacy-stub opt-in
  - does not affect openmls-inbox-dev --ack
  - does not affect Relay Space scoped ack helpers

### `comms.inbox`

- **Command:** `go run ./cmd/comms inbox`
- **Repo:** `carbonstack-comms`
- **Component:** `cmd/comms`
- **Kind:** `cli`
- **Audience:** `legacy`
- **Maturity:** `legacy`
- **Lifecycle status:** `legacy_stub_explicit_opt_in`
- **Introduced in:** `Not recorded in registry.`
- **Source path:** `carbonstack-comms/internal/app/commands.go`
- **Validation surface:** Comms package tests
- **Front README candidate:** `false`

**What it does:** Explicitly gated stub-era inbox command using the older mock/stub message path.

**Why it exists:** Preserved for continuity and low-level compatibility tests while OpenMLS runtime commands mature; requires explicit legacy opt-in.

- **Required flags:**
  - `--allow-legacy-stub` — explicit opt-in required to run the legacy mock/stub inbox path. Boundary: prevents accidental use of stub-era inbox as normal message UX.
- **Optional flags:**
  - `--state` — local Comms state path; defaults to the normal Comms state path.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - old inbox is not OpenMLS-backed yet
  - not mature messaging UX
  - requires explicit legacy-stub opt-in

### `comms.send`

- **Command:** `go run ./cmd/comms send`
- **Repo:** `carbonstack-comms`
- **Component:** `cmd/comms`
- **Kind:** `cli`
- **Audience:** `legacy`
- **Maturity:** `legacy`
- **Lifecycle status:** `legacy_stub_explicit_opt_in`
- **Introduced in:** `Not recorded in registry.`
- **Source path:** `carbonstack-comms/internal/app/commands.go`
- **Validation surface:** Comms package tests
- **Front README candidate:** `false`

**What it does:** Explicitly gated stub-era send command using the older mock/stub message path.

**Why it exists:** Preserved for continuity and low-level compatibility tests while OpenMLS runtime commands mature; requires explicit legacy opt-in.

- **Required flags:**
  - `--allow-legacy-stub` — explicit opt-in required to run the legacy mock/stub send path. Boundary: prevents accidental use of stub-era send as normal message UX.
  - `--to-device` — recipient Cypher device ID for the legacy stub envelope.
  - `--message` — plaintext input for the legacy mock/stub provider.
- **Optional flags:**
  - `--state` — local Comms state path; defaults to the normal Comms state path.
  - `--strict` — block sending to unknown, unverified, or changed devices.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - old send is not OpenMLS-backed yet
  - not mature messaging UX
  - requires explicit legacy-stub opt-in

## Historical scripts and smoke helpers

Entries in this section: **16**

### `carbonstack.script.rehearse-v0.5.0-package`

- **Command:** `scripts/rehearse-v0.5.0-package.sh`
- **Repo:** `carbonstack`
- **Component:** `scripts`
- **Kind:** `script`
- **Audience:** `dev`
- **Maturity:** `experimental`
- **Introduced in:** `v0.4.22`
- **Source path:** `carbonstack/scripts/rehearse-v0.5.0-package.sh`
- **Validation surface:** v0.5.0 package checksum/archive/fresh-extraction rehearsal
- **Front README candidate:** `false`

**What it does:** Stage a v0.5.0 package, write checksums, archive it, fresh-extract it, verify checksums, and run full validation.

**Why it exists:** Provides repeatable v0.5.0 release-package rehearsal before final release assets and tag.

- **Example:** `cd carbonstack && scripts/rehearse-v0.5.0-package.sh`
- **Required flags:** Not recorded in registry.
- **Optional flags:** Not recorded in registry.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - does not cut release
  - does not upload assets
  - does not include carbonstack-os
  - does not make runtime OpenMLS profiles part of full
  - does not create local-backbone

### `carbonstack.script.rehearse-v0.6.0-package`

- **Command:** `scripts/rehearse-v0.6.0-package.sh`
- **Repo:** `carbonstack`
- **Component:** `scripts`
- **Kind:** `script`
- **Audience:** `dev`
- **Maturity:** `experimental`
- **Introduced in:** `Not recorded in registry.`
- **Source path:** `carbonstack/scripts/rehearse-v0.6.0-package.sh`
- **Validation surface:** v0.6.0 staged archive fresh-extraction rehearsal
- **Front README candidate:** `false`

**What it does:** Stage, archive, fresh-extract, checksum-verify, and run full validation from a v0.6.0 rehearsal package root.

**Why it exists:** Provides the v0.6.0 G9 package-rehearsal proof path from a fresh extracted package root.

- **Required flags:** Not recorded in registry.
- **Optional flags:**
  - `positional stage_root` — Optional stage root; defaults to /tmp/carbonstack-v0.6.0-stage. Boundary: Must be /tmp/carbonstack-* or $HOME/carbonstack-* because the stage script removes the stage root.
  - `positional rehearsal_root` — Optional rehearsal/extraction root; defaults to /tmp/carbonstack-v0.6.0-rehearsal. Boundary: Must be /tmp/carbonstack-* or $HOME/carbonstack-* because the script removes the rehearsal root.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - does not cut release
  - does not upload assets
  - does not include carbonstack-os
  - not production security proof
  - not local-backbone
  - not hostile-server safety

### `carbonstack.script.rehearse-v0.7.0-package`

- **Command:** `scripts/rehearse-v0.7.0-package.sh`
- **Repo:** `carbonstack`
- **Component:** `scripts`
- **Kind:** `script`
- **Audience:** `dev`
- **Maturity:** `experimental`
- **Introduced in:** `v0.6.33`
- **Source path:** `carbonstack/scripts/rehearse-v0.7.0-package.sh`
- **Validation surface:** v0.7.0 staged archive fresh-extraction rehearsal
- **Front README candidate:** `false`

**What it does:** Stage, archive, fresh-extract, asset-checksum-verify, package-checksum-verify, run full-validate-release, and scan final artifacts for the v0.7.0 rehearsal package.

**Why it exists:** Provides the v0.7.0 fresh package-root rehearsal path before any public release tag or asset upload, preserving package artifact hygiene and current release-validation naming boundaries.

- **Required flags:** Not recorded in registry.
- **Optional flags:**
  - `positional stage_root` — Optional stage root; defaults to /tmp/carbonstack-v0.7.0-stage. Boundary: Must be /tmp/carbonstack-* or $HOME/carbonstack-* because the stage script removes the stage root.
  - `positional rehearsal_root` — Optional rehearsal/extraction root; defaults to /tmp/carbonstack-v0.7.0-rehearsal. Boundary: Must be /tmp/carbonstack-* or $HOME/carbonstack-* because the script removes the rehearsal root.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - does not cut release
  - does not upload assets
  - does not include carbonstack-os
  - not production security proof
  - not deployment
  - not local-backbone
  - not hostile-server safety
  - package rehearsal evidence only

### `carbonstack.script.stage-v0.5.0-package`

- **Command:** `scripts/stage-v0.5.0-package.sh`
- **Repo:** `carbonstack`
- **Component:** `scripts`
- **Kind:** `script`
- **Audience:** `dev`
- **Maturity:** `experimental`
- **Introduced in:** `v0.4.21`
- **Source path:** `carbonstack/scripts/stage-v0.5.0-package.sh`
- **Validation surface:** v0.5.0 package rehearsal staging
- **Front README candidate:** `false`

**What it does:** Stage a clean v0.5.0 package skeleton from tracked carbonstack, carbonstack-comms, and carbonstack-cypher files.

**Why it exists:** Provides repeatable package staging before checksum, archive, fresh extraction validation, and final release asset generation.

- **Example:** `cd carbonstack && scripts/stage-v0.5.0-package.sh`
- **Required flags:** Not recorded in registry.
- **Optional flags:** Not recorded in registry.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - does not write final checksums
  - does not archive package
  - does not cut release
  - does not validate fresh extraction
  - does not include carbonstack-os

### `carbonstack.script.stage-v0.6.0-package`

- **Command:** `scripts/stage-v0.6.0-package.sh`
- **Repo:** `carbonstack`
- **Component:** `scripts`
- **Kind:** `script`
- **Audience:** `dev`
- **Maturity:** `experimental`
- **Introduced in:** `Not recorded in registry.`
- **Source path:** `carbonstack/scripts/stage-v0.6.0-package.sh`
- **Validation surface:** v0.6.0 package staging and checksum preparation
- **Front README candidate:** `false`

**What it does:** Stage a tracked-source v0.6.0 rehearsal package with release metadata, registry artifacts, checksums, and archive output.

**Why it exists:** Provides the v0.6.0-specific staging path so historical v0.5.0 package scripts are not reused as v0.6.0 evidence unchanged.

- **Required flags:** Not recorded in registry.
- **Optional flags:**
  - `positional stage_root` — Optional stage root; defaults to /tmp/carbonstack-v0.6.0-stage. Boundary: Must be /tmp/carbonstack-* or $HOME/carbonstack-* because the script removes the stage root.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - does not cut release
  - does not upload assets
  - does not include carbonstack-os
  - not production security validation
  - not local-backbone

### `carbonstack.script.stage-v0.7.0-package`

- **Command:** `scripts/stage-v0.7.0-package.sh`
- **Repo:** `carbonstack`
- **Component:** `scripts`
- **Kind:** `script`
- **Audience:** `dev`
- **Maturity:** `experimental`
- **Introduced in:** `v0.6.33`
- **Source path:** `carbonstack/scripts/stage-v0.7.0-package.sh`
- **Validation surface:** v0.7.0 package staging and release metadata generation
- **Front README candidate:** `false`

**What it does:** Stage a tracked-source v0.7.0 rehearsal package with release metadata, generated command reference, checksums, staged release assets, and archive output.

**Why it exists:** Provides the v0.7.0-specific package staging path so v0.6.0 package scripts are treated as prior art rather than reused as v0.7.0 evidence unchanged.

- **Required flags:** Not recorded in registry.
- **Optional flags:**
  - `positional stage_root` — Optional stage root; defaults to /tmp/carbonstack-v0.7.0-stage. Boundary: Must be /tmp/carbonstack-* or $HOME/carbonstack-* because the script removes the stage root.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - does not cut release
  - does not upload assets
  - does not include carbonstack-os
  - not production security validation
  - not deployment
  - not local-backbone
  - not hostile-server safety

### `carbonstack.script.validate-local`

- **Command:** `powershell -ExecutionPolicy Bypass -File ./scripts/validate-local.ps1`
- **Repo:** `carbonstack`
- **Component:** `scripts`
- **Kind:** `script`
- **Audience:** `legacy`
- **Maturity:** `legacy`
- **Introduced in:** `Not recorded in registry.`
- **Source path:** `carbonstack/scripts/validate-local.ps1`
- **Validation surface:** historical continuity only; superseded by current Go runner profiles where applicable
- **Front README candidate:** `false`

**What it does:** Historical/local PowerShell validation helper.

**Why it exists:** Preserved for historical continuity; Go runner is current main validation surface.

- **Required flags:** Not recorded in registry.
- **Optional flags:** Not recorded in registry.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - legacy continuity surface
  - not current primary validation path
  - not release-package validation
  - not production security proof
  - not local-backbone

### `carbonstack.script.validate-phase1`

- **Command:** `powershell -ExecutionPolicy Bypass -File ./scripts/validate-phase1.ps1`
- **Repo:** `carbonstack`
- **Component:** `scripts`
- **Kind:** `script`
- **Audience:** `legacy`
- **Maturity:** `legacy`
- **Introduced in:** `Not recorded in registry.`
- **Source path:** `carbonstack/scripts/validate-phase1.ps1`
- **Validation surface:** historical continuity only; superseded by current Go runner profiles where applicable
- **Front README candidate:** `false`

**What it does:** Historical Phase 1 PowerShell validation helper.

**Why it exists:** Preserved for historical continuity; Go runner is current main validation surface.

- **Required flags:** Not recorded in registry.
- **Optional flags:** Not recorded in registry.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - legacy continuity surface
  - not current primary validation path
  - not release-package validation
  - not production security proof
  - not local-backbone

### `comms.script.check-no-rust-artifacts`

- **Command:** `powershell -ExecutionPolicy Bypass -File ./scripts/check-no-rust-artifacts.ps1`
- **Repo:** `carbonstack-comms`
- **Component:** `scripts`
- **Kind:** `script`
- **Audience:** `legacy`
- **Maturity:** `legacy`
- **Introduced in:** `Not recorded in registry.`
- **Source path:** `carbonstack-comms/scripts/check-no-rust-artifacts.ps1`
- **Validation surface:** historical continuity only; superseded by current Go runner profiles where applicable
- **Front README candidate:** `false`

**What it does:** Check for Rust build/generated artifacts in the repo.

**Why it exists:** Historical hygiene guard before runner artifact scans matured.

- **Required flags:** Not recorded in registry.
- **Optional flags:** Not recorded in registry.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - legacy continuity surface
  - not current primary validation path
  - not release-package validation
  - not production security proof
  - not local-backbone

### `comms.script.direct-openmls-runtime-smoke`

- **Command:** `scripts/dev-openmls-runtime-smoke.sh`
- **Repo:** `carbonstack-comms`
- **Component:** `scripts`
- **Kind:** `script`
- **Audience:** `dev`
- **Maturity:** `dev_only`
- **Introduced in:** `v0.4.5`
- **Source path:** `carbonstack-comms/scripts/dev-openmls-runtime-smoke.sh`
- **Validation surface:** dev-runtime-openmls runner profile and direct OpenMLS runtime smoke
- **Front README candidate:** `false`

**What it does:** Direct-sidecar bootstrap smoke proving openmls-send-dev -> Cypher -> openmls-inbox-dev --ack.

**Why it exists:** Known-good lower-level OpenMLS runtime baseline.

- **Required flags:** Not recorded in registry.
- **Optional flags:** Not recorded in registry.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - not local-backbone
  - not production messaging UX

### `comms.script.openmls-relay-narrow-join-smoke-dev`

- **Command:** `scripts/openmls-relay-narrow-join-smoke-dev.sh`
- **Repo:** `carbonstack-comms`
- **Component:** `scripts`
- **Kind:** `script`
- **Audience:** `dev`
- **Maturity:** `dev_only`
- **Introduced in:** `Not recorded in registry.`
- **Source path:** `carbonstack-comms/scripts/openmls-relay-narrow-join-smoke-dev.sh`
- **Validation surface:** relay-openmls-join-dev
- **Front README candidate:** `false`

**What it does:** Scripted Relay Space OpenMLS KeyPackage -> add-member -> Welcome -> join smoke path used by the runner profile.

**Why it exists:** Provides the shell-level proof path wrapped by runner.relay-openmls-join-dev.

- **Required flags:** Not recorded in registry.
- **Optional flags:** Not recorded in registry.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - not standalone release-package validation
  - not local-backbone
  - not production messaging UX

### `comms.script.self-test-openmls-backbone`

- **Command:** `powershell -ExecutionPolicy Bypass -File ./scripts/self-test-openmls-backbone.ps1`
- **Repo:** `carbonstack-comms`
- **Component:** `scripts`
- **Kind:** `script`
- **Audience:** `legacy`
- **Maturity:** `legacy`
- **Introduced in:** `Not recorded in registry.`
- **Source path:** `carbonstack-comms/scripts/self-test-openmls-backbone.ps1`
- **Validation surface:** historical continuity only; superseded by current Go runner profiles where applicable
- **Front README candidate:** `false`

**What it does:** Legacy/Windows-oriented OpenMLS backbone self-test harness retained for continuity.

**Why it exists:** Historical validation route before newer Go runner/runtime profiles matured.

- **Required flags:** Not recorded in registry.
- **Optional flags:** Not recorded in registry.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - not current primary WSL Debian validation path

### `comms.script.smoke-openmls-real-cypher-relay`

- **Command:** `powershell -ExecutionPolicy Bypass -File ./scripts/smoke-openmls-real-cypher-relay.ps1`
- **Repo:** `carbonstack-comms`
- **Component:** `scripts`
- **Kind:** `script`
- **Audience:** `legacy`
- **Maturity:** `legacy`
- **Introduced in:** `Not recorded in registry.`
- **Source path:** `carbonstack-comms/scripts/smoke-openmls-real-cypher-relay.ps1`
- **Validation surface:** historical continuity only; superseded by current Go runner profiles where applicable
- **Front README candidate:** `false`

**What it does:** Lower-level historical real-Cypher/OpenMLS relay smoke harness.

**Why it exists:** Debug and continuity path from earlier relay proof work.

- **Required flags:** Not recorded in registry.
- **Optional flags:** Not recorded in registry.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - legacy continuity surface
  - not current primary validation path
  - not release-package validation
  - not production security proof
  - not local-backbone

### `comms.script.test-local-lifecycle`

- **Command:** `powershell -ExecutionPolicy Bypass -File ./scripts/test-local-lifecycle.ps1`
- **Repo:** `carbonstack-comms`
- **Component:** `scripts`
- **Kind:** `script`
- **Audience:** `legacy`
- **Maturity:** `legacy`
- **Introduced in:** `Not recorded in registry.`
- **Source path:** `carbonstack-comms/scripts/test-local-lifecycle.ps1`
- **Validation surface:** historical continuity only; superseded by current Go runner profiles where applicable
- **Front README candidate:** `false`

**What it does:** Historical local lifecycle PowerShell validation helper.

**Why it exists:** Preserved for continuity; superseded by Go runner profiles where applicable.

- **Required flags:** Not recorded in registry.
- **Optional flags:** Not recorded in registry.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - legacy continuity surface
  - not current primary validation path
  - not release-package validation
  - not production security proof
  - not local-backbone

### `comms.script.test-trust-lifecycle`

- **Command:** `powershell -ExecutionPolicy Bypass -File ./scripts/test-trust-lifecycle.ps1`
- **Repo:** `carbonstack-comms`
- **Component:** `scripts`
- **Kind:** `script`
- **Audience:** `legacy`
- **Maturity:** `legacy`
- **Introduced in:** `Not recorded in registry.`
- **Source path:** `carbonstack-comms/scripts/test-trust-lifecycle.ps1`
- **Validation surface:** historical continuity only; superseded by current Go runner profiles where applicable
- **Front README candidate:** `false`

**What it does:** Historical trust lifecycle PowerShell validation helper.

**Why it exists:** Preserved for trust lifecycle continuity.

- **Required flags:** Not recorded in registry.
- **Optional flags:** Not recorded in registry.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - legacy continuity surface
  - not current primary validation path
  - not release-package validation
  - not production security proof
  - not local-backbone

### `comms.script.wrapper-openmls-runtime-smoke`

- **Command:** `scripts/dev-openmls-runtime-smoke-wrappers.sh`
- **Repo:** `carbonstack-comms`
- **Component:** `scripts`
- **Kind:** `script`
- **Audience:** `dev`
- **Maturity:** `dev_only`
- **Introduced in:** `v0.4.16`
- **Source path:** `carbonstack-comms/scripts/dev-openmls-runtime-smoke-wrappers.sh`
- **Validation surface:** dev-runtime-openmls-wrappers runner profile and scoped wrapper OpenMLS runtime smoke
- **Front README candidate:** `false`

**What it does:** Create a Relay Space, register active Alice/Bob routing members, then prove scoped message-send-dev -> Cypher -> message-inbox-dev --ack.

**Why it exists:** Higher-level wrapper maturity proof for the Relay Space-scoped normal-message path while the lower-level direct-sidecar baseline remains separately callable.

- **Required flags:** Not recorded in registry.
- **Optional flags:** Not recorded in registry.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - not local-backbone
  - does not replace direct smoke yet
  - not Relay onboarding UX
  - active membership is routing authority only
  - not identity verification
  - not production messaging UX
  - not production security proof

## Internal OpenMLS sidecar provider commands

Entries in this section: **14**

### `sidecar.conversation-add-member`

- **Command:** `cargo run -- conversation-add-member`
- **Repo:** `carbonstack-comms`
- **Component:** `openmls-sidecar`
- **Kind:** `sidecar-cli`
- **Audience:** `internal`
- **Maturity:** `internal`
- **Introduced in:** `Not recorded in registry.`
- **Source path:** `carbonstack-comms/internal/protocol/mls/openmls-sidecar/src/main.rs`
- **Validation surface:** dev-runtime-openmls-wrappers, relay-openmls-join-dev, and Comms OpenMLS bootstrap tests
- **Front README candidate:** `false`

**What it does:** Consume a KeyPackage, add member, and write a Welcome artifact.

**Why it exists:** Internal provider command wrapped by openmls-conversation-add-member-dev.

- **Wrapped by:** openmls-conversation-add-member-dev
- **Required flags:** Not recorded in registry.
- **Optional flags:** Not recorded in registry.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - internal provider command
  - dev-local membership artifact generation only
  - not production membership UX
  - not identity verification
  - not secure enrollment

### `sidecar.conversation-create`

- **Command:** `cargo run -- conversation-create`
- **Repo:** `carbonstack-comms`
- **Component:** `openmls-sidecar`
- **Kind:** `sidecar-cli`
- **Audience:** `internal`
- **Maturity:** `internal`
- **Introduced in:** `Not recorded in registry.`
- **Source path:** `carbonstack-comms/internal/protocol/mls/openmls-sidecar/src/main.rs`
- **Validation surface:** dev-runtime-openmls-wrappers and Comms OpenMLS bootstrap tests
- **Front README candidate:** `false`

**What it does:** Create sidecar OpenMLS conversation state.

**Why it exists:** Internal provider command wrapped by openmls-conversation-create-dev.

- **Wrapped by:** openmls-conversation-create-dev
- **Required flags:** Not recorded in registry.
- **Optional flags:** Not recorded in registry.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - internal provider command
  - dev-local conversation state only
  - not production group creation UX
  - not secure enrollment

### `sidecar.conversation-join`

- **Command:** `cargo run -- conversation-join`
- **Repo:** `carbonstack-comms`
- **Component:** `openmls-sidecar`
- **Kind:** `sidecar-cli`
- **Audience:** `internal`
- **Maturity:** `internal`
- **Introduced in:** `Not recorded in registry.`
- **Source path:** `carbonstack-comms/internal/protocol/mls/openmls-sidecar/src/main.rs`
- **Validation surface:** dev-runtime-openmls-wrappers, relay-openmls-join-dev, and Comms OpenMLS bootstrap tests
- **Front README candidate:** `false`

**What it does:** Consume a Welcome artifact and join a sidecar conversation.

**Why it exists:** Internal provider command wrapped by openmls-conversation-join-dev.

- **Wrapped by:** openmls-conversation-join-dev
- **Required flags:** Not recorded in registry.
- **Optional flags:** Not recorded in registry.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - internal provider command
  - dev-local Welcome consumption only
  - not production join UX
  - not identity verification
  - not secure enrollment

### `sidecar.conversation-load-check`

- **Command:** `cargo run -- conversation-load-check`
- **Repo:** `carbonstack-comms`
- **Component:** `openmls-sidecar`
- **Kind:** `sidecar-cli`
- **Audience:** `internal`
- **Maturity:** `internal`
- **Introduced in:** `Not recorded in registry.`
- **Source path:** `carbonstack-comms/internal/protocol/mls/openmls-sidecar/src/main.rs`
- **Validation surface:** dev-runtime-openmls-wrappers and Comms OpenMLS bootstrap tests
- **Front README candidate:** `false`

**What it does:** Load-check a sidecar conversation.

**Why it exists:** Internal provider sanity check wrapped by openmls-conversation-load-check-dev.

- **Wrapped by:** openmls-conversation-load-check-dev
- **Required flags:** Not recorded in registry.
- **Optional flags:** Not recorded in registry.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - internal provider command
  - dev-local state load check only
  - not recovery
  - not production durability proof

### `sidecar.identity-create`

- **Command:** `cargo run -- identity-create`
- **Repo:** `carbonstack-comms`
- **Component:** `openmls-sidecar`
- **Kind:** `sidecar-cli`
- **Audience:** `internal`
- **Maturity:** `internal`
- **Introduced in:** `Not recorded in registry.`
- **Source path:** `carbonstack-comms/internal/protocol/mls/openmls-sidecar/src/main.rs`
- **Validation surface:** dev-runtime-openmls-wrappers and Comms OpenMLS bootstrap tests
- **Front README candidate:** `false`

**What it does:** Create dev-local OpenMLS sidecar identity state.

**Why it exists:** Internal provider command wrapped by openmls-identity-create-dev.

- **Wrapped by:** openmls-identity-create-dev
- **Required flags:** Not recorded in registry.
- **Optional flags:** Not recorded in registry.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - internal provider command
  - dev-local identity state only
  - not production identity enrollment
  - not secure key storage
  - not verified identity

### `sidecar.identity-status`

- **Command:** `cargo run -- identity-status`
- **Repo:** `carbonstack-comms`
- **Component:** `openmls-sidecar`
- **Kind:** `sidecar-cli`
- **Audience:** `internal`
- **Maturity:** `internal`
- **Introduced in:** `Not recorded in registry.`
- **Source path:** `carbonstack-comms/internal/protocol/mls/openmls-sidecar/src/main.rs`
- **Validation surface:** dev-runtime-openmls-wrappers and Comms OpenMLS bootstrap tests
- **Front README candidate:** `false`

**What it does:** Inspect dev-local OpenMLS sidecar identity state.

**Why it exists:** Internal provider command wrapped by openmls-identity-status-dev.

- **Wrapped by:** openmls-identity-status-dev
- **Required flags:** Not recorded in registry.
- **Optional flags:** Not recorded in registry.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - internal provider command
  - dev-local status inspection only
  - not identity verification
  - not production trust UX

### `sidecar.keypackage-generate`

- **Command:** `cargo run -- keypackage-generate --device-label <label> --request-id <safe-id>`
- **Repo:** `carbonstack-comms`
- **Component:** `internal/protocol/mls/openmls-sidecar`
- **Kind:** `sidecar-command`
- **Audience:** `internal`
- **Maturity:** `internal`
- **Lifecycle status:** `active`
- **Introduced in:** `v0.7.5`
- **Source path:** `carbonstack-comms/internal/protocol/mls/openmls-sidecar/src/keypackage_rotation.rs`
- **Validation surface:** persistent-repeatable-keypackage-generation-idempotent-request-legacy-adoption-provider-state
- **Front README candidate:** `false`

**What it does:** Generate or replay one device-local persistent OpenMLS KeyPackage generation.

**Why it exists:** Establish Gate B5b repeatable generation while preserving provider state and every prior generation.

- **Required flags:**
  - `--device-label` — Local sidecar device label that owns the KeyPackage inventory and provider state.
  - `--request-id` — Permanent device-local idempotency identity for the generation request.
- **Optional flags:** Not recorded in registry.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - not Relay publication
  - not KeyPackage consumption
  - not KeyPackage ACK
  - not Welcome lifecycle
  - not account or Relay Space identity binding
  - not trust promotion
  - not secure enrollment
  - not production onboarding
  - not local-backbone
  - not deployment
  - not audit or certification

### `sidecar.keypackage-inspect`

- **Command:** `cargo run -- keypackage-inspect --device-label <label> --keypackage <path>`
- **Repo:** `carbonstack-comms`
- **Component:** `internal/protocol/mls/openmls-sidecar`
- **Kind:** `sidecar-command`
- **Audience:** `internal`
- **Maturity:** `internal`
- **Lifecycle status:** `active`
- **Introduced in:** `v0.7.5`
- **Source path:** `carbonstack-comms/internal/protocol/mls/openmls-sidecar/src/keypackage_inspect.rs`
- **Validation surface:** read-only-openmls-keypackage-ref-lifetime-artifact-integrity-local-owner-evidence
- **Front README candidate:** `false`

**What it does:** Inspect and validate a serialized OpenMLS KeyPackage without mutating state.

**Why it exists:** Establish the Gate B5a material KeyPackage identity, lifetime, integrity, and local ownership foundation.

- **Required flags:**
  - `--device-label` — Local sidecar device label whose public-bundle metadata provides ownership evidence.
  - `--keypackage` — Serialized OpenMLS KeyPackage artifact path to inspect.
- **Optional flags:**
  - `--generation-manifest` — Optional immutable B5b generation manifest used as local ownership evidence.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - not repeatable KeyPackage generation
  - not KeyPackage rotation
  - not Relay publication
  - not KeyPackage consumption
  - not Welcome lifecycle
  - not account identity verification
  - not device identity verification
  - not Relay Space membership verification
  - not trust promotion
  - not secure enrollment
  - not production E2EE
  - not local-backbone
  - not deployment
  - not audit or certification

### `sidecar.keypackage-inventory`

- **Command:** `cargo run -- keypackage-inventory --device-label <label>`
- **Repo:** `carbonstack-comms`
- **Component:** `internal/protocol/mls/openmls-sidecar`
- **Kind:** `sidecar-command`
- **Audience:** `internal`
- **Maturity:** `internal`
- **Lifecycle status:** `active`
- **Introduced in:** `v0.7.5`
- **Source path:** `carbonstack-comms/internal/protocol/mls/openmls-sidecar/src/keypackage_rotation.rs`
- **Validation surface:** read-only-persistent-keypackage-generation-inventory
- **Front README candidate:** `false`

**What it does:** Inspect the persistent device-local KeyPackage generation inventory.

**Why it exists:** Expose current, active, and retired generations across restart without implicit repair.

- **Required flags:**
  - `--device-label` — Local sidecar device label whose inventory is inspected.
- **Optional flags:** Not recorded in registry.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - not mutation or implicit repair
  - not Relay publication
  - not KeyPackage consumption
  - not KeyPackage ACK
  - not Welcome lifecycle
  - not identity verification
  - not trust promotion
  - not production onboarding
  - not local-backbone
  - not deployment
  - not audit or certification

### `sidecar.keypackage-retire`

- **Command:** `cargo run -- keypackage-retire --device-label <label> --generation-id <generation-id>`
- **Repo:** `carbonstack-comms`
- **Component:** `internal/protocol/mls/openmls-sidecar`
- **Kind:** `sidecar-command`
- **Audience:** `internal`
- **Maturity:** `internal`
- **Lifecycle status:** `active`
- **Introduced in:** `v0.7.5`
- **Source path:** `carbonstack-comms/internal/protocol/mls/openmls-sidecar/src/keypackage_rotation.rs`
- **Validation surface:** explicit-noncurrent-keypackage-retirement-metadata-only
- **Front README candidate:** `false`

**What it does:** Retire a non-current KeyPackage generation without deleting its material.

**Why it exists:** Establish explicit overlap retirement while preserving artifacts and private provider state for later Welcome handling.

- **Required flags:**
  - `--device-label` — Local sidecar device label that owns the generation.
  - `--generation-id` — Non-current persistent generation to mark retired.
- **Optional flags:** Not recorded in registry.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - not deletion or garbage collection
  - not Relay revocation
  - not KeyPackage consumption
  - not KeyPackage ACK
  - not Welcome lifecycle
  - not trust revocation
  - not production onboarding
  - not local-backbone
  - not deployment
  - not audit or certification

### `sidecar.message-open`

- **Command:** `cargo run -- message-open`
- **Repo:** `carbonstack-comms`
- **Component:** `openmls-sidecar`
- **Kind:** `sidecar-cli`
- **Audience:** `internal`
- **Maturity:** `internal`
- **Introduced in:** `Not recorded in registry.`
- **Source path:** `carbonstack-comms/internal/protocol/mls/openmls-sidecar/src/main.rs`
- **Validation surface:** dev-runtime-openmls, dev-runtime-openmls-wrappers, and Comms OpenMLS runtime tests
- **Front README candidate:** `false`

**What it does:** Open an OpenMLS application-message artifact and print plaintext after sidecar success.

**Why it exists:** Internal provider command used by openmls-inbox-dev.

- **Wrapped by:** openmls-inbox-dev
- **Required flags:** Not recorded in registry.
- **Optional flags:** Not recorded in registry.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - internal provider command
  - dev/pre-alpha application-message open primitive
  - not production E2EE claim
  - not hostile-server safety
  - not metadata privacy
  - not mature receive UX

### `sidecar.message-protect`

- **Command:** `cargo run -- message-protect`
- **Repo:** `carbonstack-comms`
- **Component:** `openmls-sidecar`
- **Kind:** `sidecar-cli`
- **Audience:** `internal`
- **Maturity:** `internal`
- **Introduced in:** `Not recorded in registry.`
- **Source path:** `carbonstack-comms/internal/protocol/mls/openmls-sidecar/src/main.rs`
- **Validation surface:** dev-runtime-openmls, dev-runtime-openmls-wrappers, and Comms OpenMLS runtime tests
- **Front README candidate:** `false`

**What it does:** Protect plaintext into an OpenMLS application-message artifact.

**Why it exists:** Internal provider command used by openmls-send-dev.

- **Wrapped by:** openmls-send-dev
- **Required flags:** Not recorded in registry.
- **Optional flags:** Not recorded in registry.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - internal provider command
  - dev/pre-alpha application-message protect primitive
  - not production E2EE claim
  - not hostile-server safety
  - not metadata privacy
  - not mature send UX

### `sidecar.provider-info`

- **Command:** `cargo run -- provider-info`
- **Repo:** `carbonstack-comms`
- **Component:** `openmls-sidecar`
- **Kind:** `sidecar-cli`
- **Audience:** `internal`
- **Maturity:** `internal`
- **Introduced in:** `Not recorded in registry.`
- **Source path:** `carbonstack-comms/internal/protocol/mls/openmls-sidecar/src/main.rs`
- **Validation surface:** OpenMLS sidecar build/contract sanity through Comms tests and runner doctor visibility
- **Front README candidate:** `false`

**What it does:** Print provider/sidecar information.

**Why it exists:** Sidecar diagnostics and contract sanity.

- **Required flags:** Not recorded in registry.
- **Optional flags:** Not recorded in registry.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - internal diagnostic surface
  - not user-facing UX
  - not production provider attestation
  - not security audit

### `sidecar.public-bundle-export`

- **Command:** `cargo run -- public-bundle-export`
- **Repo:** `carbonstack-comms`
- **Component:** `openmls-sidecar`
- **Kind:** `sidecar-cli`
- **Audience:** `internal`
- **Maturity:** `internal`
- **Introduced in:** `Not recorded in registry.`
- **Source path:** `carbonstack-comms/internal/protocol/mls/openmls-sidecar/src/main.rs`
- **Validation surface:** dev-runtime-openmls-wrappers, relay-openmls-join-dev, and Comms OpenMLS bootstrap tests
- **Front README candidate:** `false`

**What it does:** Export a KeyPackage/public bundle and optionally write an artifact.

**Why it exists:** Internal provider command wrapped by openmls-bundle-export-dev.

- **Wrapped by:** openmls-bundle-export-dev
- **Required flags:** Not recorded in registry.
- **Optional flags:** Not recorded in registry.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - internal provider command
  - exports dev KeyPackage/public bundle artifacts
  - not production key distribution UX
  - not secure enrollment
  - not identity verification

## Cypher server and HTTP API surfaces

Entries in this section: **11**

### `cypher.api.accounts-devices`

- **Command:** `GET /v0/accounts/{account_id}/devices`
- **Repo:** `carbonstack-cypher`
- **Component:** `httpapi`
- **Kind:** `api-surface`
- **Audience:** `dev`
- **Maturity:** `experimental`
- **Introduced in:** `Not recorded in registry.`
- **Source path:** `carbonstack-cypher/internal/httpapi`
- **Validation surface:** local-cypher and Comms device listing smoke paths
- **Front README candidate:** `false`

**What it does:** List devices under an account.

**Why it exists:** Validates persisted local Cypher device state.

- **Required flags:** Not recorded in registry.
- **Optional flags:** Not recorded in registry.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - dev/pre-alpha device listing API
  - not authorization model proof
  - not production device management UX
  - not metadata privacy claim

### `cypher.api.dev-invites`

- **Command:** `POST /v0/dev/invites`
- **Repo:** `carbonstack-cypher`
- **Component:** `httpapi`
- **Kind:** `api-surface`
- **Audience:** `dev`
- **Maturity:** `dev_only`
- **Introduced in:** `Not recorded in registry.`
- **Source path:** `carbonstack-cypher/internal/httpapi`
- **Validation surface:** local-cypher and Comms dev invite smoke paths
- **Front README candidate:** `false`

**What it does:** Create development invite codes.

**Why it exists:** Dev-only test/onboarding helper.

- **Required flags:** Not recorded in registry.
- **Optional flags:** Not recorded in registry.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - dev-only invite creation API
  - not production invite issuance
  - not abuse-resistant onboarding
  - not public API stability claim

### `cypher.api.device-envelopes`

- **Command:** `GET /v0/devices/{device_id}/envelopes`
- **Repo:** `carbonstack-cypher`
- **Component:** `httpapi`
- **Kind:** `api-surface`
- **Audience:** `dev`
- **Maturity:** `experimental`
- **Introduced in:** `Not recorded in registry.`
- **Source path:** `carbonstack-cypher/internal/httpapi`
- **Validation surface:** local-cypher, dev-runtime-openmls, dev-runtime-openmls-wrappers, and integrated-runtime-dev
- **Front README candidate:** `false`

**What it does:** Retrieve queued envelopes for a device.

**Why it exists:** Inbox fetch primitive for Comms runtime commands and smoke scripts.

- **Required flags:** Not recorded in registry.
- **Optional flags:** Not recorded in registry.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - dev/pre-alpha envelope fetch API
  - fetch alone is not ack eligibility
  - not durable local receive storage
  - not hostile-server safety
  - not metadata privacy

### `cypher.api.devices-register`

- **Command:** `POST /v0/devices/register`
- **Repo:** `carbonstack-cypher`
- **Component:** `httpapi`
- **Kind:** `api-surface`
- **Audience:** `dev`
- **Maturity:** `experimental`
- **Introduced in:** `Not recorded in registry.`
- **Source path:** `carbonstack-cypher/internal/httpapi`
- **Validation surface:** local-cypher and Comms device registration smoke paths
- **Front README candidate:** `false`

**What it does:** Register a device for an account.

**Why it exists:** Provides recipient/sender device IDs for envelope routing.

- **Required flags:** Not recorded in registry.
- **Optional flags:** Not recorded in registry.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - dev/pre-alpha device registration API
  - not verified device ceremony
  - not production device management
  - not secure enrollment

### `cypher.api.envelope-ack`

- **Command:** `POST /v0/envelopes/{envelope_id}/ack`
- **Repo:** `carbonstack-cypher`
- **Component:** `httpapi`
- **Kind:** `api-surface`
- **Audience:** `dev`
- **Maturity:** `experimental`
- **Introduced in:** `Not recorded in registry.`
- **Source path:** `carbonstack-cypher/internal/httpapi`
- **Validation surface:** local-cypher, dev-runtime-openmls, dev-runtime-openmls-wrappers, and integrated-runtime-dev
- **Front README candidate:** `false`

**What it does:** Acknowledge an envelope after successful receive/open.

**Why it exists:** Supports delivery-state transition and inbox-empty-after-ack proofs.

- **Required flags:** Not recorded in registry.
- **Optional flags:** Not recorded in registry.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - dev/pre-alpha envelope delivery-state API
  - ack semantics depend on caller terminal processing rules
  - not a standalone receive/open proof
  - not hostile-server safety
  - not metadata privacy

### `cypher.api.envelopes-submit`

- **Command:** `POST /v0/envelopes`
- **Repo:** `carbonstack-cypher`
- **Component:** `httpapi`
- **Kind:** `api-surface`
- **Audience:** `dev`
- **Maturity:** `experimental`
- **Introduced in:** `Not recorded in registry.`
- **Source path:** `carbonstack-cypher/internal/httpapi`
- **Validation surface:** local-cypher and OpenMLS runtime profiles
- **Front README candidate:** `false`

**What it does:** Submit an opaque envelope payload for a recipient device.

**Why it exists:** Core relay/storage primitive for Comms message delivery.

- **Required flags:** Not recorded in registry.
- **Optional flags:** Not recorded in registry.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - dev/pre-alpha opaque envelope transport API
  - does not inspect or prove message security
  - not hostile-server safety
  - not metadata privacy
  - not public API stability claim

### `cypher.api.health`

- **Command:** `GET /v0/health`
- **Repo:** `carbonstack-cypher`
- **Component:** `httpapi`
- **Kind:** `api-surface`
- **Audience:** `dev`
- **Maturity:** `experimental`
- **Introduced in:** `Not recorded in registry.`
- **Source path:** `carbonstack-cypher/internal/httpapi`
- **Validation surface:** local-cypher and runtime profile readiness checks
- **Front README candidate:** `false`

**What it does:** Health check endpoint used by validation and smoke scripts.

**Why it exists:** Allows runner/scripts to wait for temporary Cypher readiness.

- **Required flags:** Not recorded in registry.
- **Optional flags:** Not recorded in registry.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - dev/local health endpoint only
  - not deployment readiness
  - not public ingress hardening
  - not security proof

### `cypher.api.invites-claim`

- **Command:** `POST /v0/invites/claim`
- **Repo:** `carbonstack-cypher`
- **Component:** `httpapi`
- **Kind:** `api-surface`
- **Audience:** `dev`
- **Maturity:** `experimental`
- **Introduced in:** `Not recorded in registry.`
- **Source path:** `carbonstack-cypher/internal/httpapi`
- **Validation surface:** local-cypher and Comms account onboarding smoke paths
- **Front README candidate:** `false`

**What it does:** Claim an invite and create/account-bind local test identity.

**Why it exists:** Supports local-cypher and Comms smoke onboarding.

- **Required flags:** Not recorded in registry.
- **Optional flags:** Not recorded in registry.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - dev/pre-alpha account bootstrap API
  - not secure enrollment
  - not identity verification
  - not production account recovery

### `cypher.api.relay-space-keypackage-publication`

- **Command:** `POST /v0/relay-spaces/{relay_space_id}/keypackage-publications`
- **Repo:** `carbonstack-cypher`
- **Component:** `internal/httpapi`
- **Kind:** `cypher-api-route`
- **Audience:** `dev`
- **Maturity:** `experimental`
- **Lifecycle status:** `active`
- **Introduced in:** `v0.7.6`
- **Source path:** `carbonstack-cypher/internal/httpapi/keypackage_publication.go`
- **Validation surface:** sender-scoped-keypackage-publication-created-replay-reuse-identity-conflict-concurrency-restart
- **Front README candidate:** `false`

**What it does:** Publish one opaque KeyPackage artifact with persistent exact-replay and reuse classification.

**Why it exists:** Establish the Gate B5c Relay publication authority without creating a public KeyPackage directory or consuming the artifact.

- **Required flags:**
  - `relay_space_id` — Relay Space path identity.
  - `sender_device_id` — Active sender routing device.
  - `recipient_device_id` — Active recipient routing device.
  - `key_package_ref` — Claimed OpenMLS KeyPackage reference in sha256 form.
  - `ciphertext_b64` — Exact serialized KeyPackage bytes encoded as base64.
- **Optional flags:**
  - `client_created_at` — Client metadata only; differences do not create another publication.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - not a public KeyPackage directory
  - not KeyPackage generation
  - not KeyPackage inspection
  - not KeyPackage consumption
  - not KeyPackage ACK
  - not add-member
  - not Welcome lifecycle
  - not identity verification
  - not trust promotion
  - not secure enrollment
  - not authenticated public API safety
  - not production key distribution
  - not deployment
  - not audit or certification

### `cypher.config-inspection`

- **Command:** `go run ./cmd/cypher --print-config`
- **Repo:** `carbonstack-cypher`
- **Component:** `cmd/cypher`
- **Kind:** `server-config-inspection`
- **Audience:** `dev`
- **Maturity:** `dev_only`
- **Lifecycle status:** `active`
- **Introduced in:** `Not recorded in registry.`
- **Source path:** `carbonstack-cypher/cmd/cypher/main.go`
- **Validation surface:** Gate E E2 Cypher config inspection
- **Front README candidate:** `false`

**What it does:** Print effective Cypher config as JSON and exit without starting the server.

**Why it exists:** Provides a terminating operator inspection surface before Gate E runbook/helper/service decisions, preventing --help/config probes from accidentally starting the blocking server with defaults.

- **Required flags:** Not recorded in registry.
- **Optional flags:**
  - `--print-config` — print effective config JSON and exit without starting the server.
  - `--check-config` — validate effective config and exit without starting the server.
  - `--help` — print Cypher usage and exit without starting the server.
- **Environment:** Not recorded in registry.
- **Related registry rows:**
  - runner.gate-e-native-deployment-dev
  - runner.local-cypher
  - cypher.server
- **Not claims:**
  - not Gate E final closure by itself
  - not deployment implementation
  - not semi-persistent service
  - not systemd
  - not helper install
  - not public ingress
  - not container readiness
  - not TUI
  - not full-runtime-dev
  - not production readiness
  - not production E2EE
  - not verified identity
  - not trust promotion
  - not vault security
  - not backup restore
  - not PQ or hybrid migration
  - not Android
  - not CarbonStackOS implementation
  - not v0.8.0 release readiness

### `cypher.server`

- **Command:** `go run ./cmd/cypher`
- **Repo:** `carbonstack-cypher`
- **Component:** `cmd/cypher`
- **Kind:** `cli`
- **Audience:** `dev`
- **Maturity:** `experimental`
- **Introduced in:** `Not recorded in registry.`
- **Source path:** `carbonstack-cypher/cmd/cypher/main.go`
- **Validation surface:** local-cypher, dev-runtime-openmls, dev-runtime-openmls-wrappers
- **Front README candidate:** `true`

**What it does:** Start the Cypher relay/storage HTTP server with explicit environment configuration.

**Why it exists:** Provides the local relay/storage server for Comms and validation profiles.

- **Required flags:** Not recorded in registry.
- **Optional flags:** Not recorded in registry.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - not deployment hardening
  - not public ingress
  - not systemd/cloudflared readiness

## Future or unsupported placeholders

Entries in this section: **2**

### `sidecar.state-checkpoint`

- **Command:** `cargo run -- state-checkpoint`
- **Repo:** `carbonstack-comms`
- **Component:** `openmls-sidecar`
- **Kind:** `sidecar-cli`
- **Audience:** `internal`
- **Maturity:** `future`
- **Introduced in:** `Not recorded in registry.`
- **Source path:** `carbonstack-comms/internal/protocol/mls/openmls-sidecar/README.md`
- **Validation surface:** unsupported placeholder; documentation/registry presence only
- **Front README candidate:** `false`

**What it does:** Unsupported placeholder mentioned by the sidecar as not currently implemented.

**Why it exists:** Future state/provider checkpoint concept.

- **Required flags:** Not recorded in registry.
- **Optional flags:** Not recorded in registry.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - unsupported

### `sidecar.state-load-check`

- **Command:** `cargo run -- state-load-check`
- **Repo:** `carbonstack-comms`
- **Component:** `openmls-sidecar`
- **Kind:** `sidecar-cli`
- **Audience:** `internal`
- **Maturity:** `future`
- **Introduced in:** `Not recorded in registry.`
- **Source path:** `carbonstack-comms/internal/protocol/mls/openmls-sidecar/README.md`
- **Validation surface:** unsupported placeholder; documentation/registry presence only
- **Front README candidate:** `false`

**What it does:** Unsupported placeholder mentioned by the sidecar as not currently implemented.

**Why it exists:** Future state/provider load-check concept.

- **Required flags:** Not recorded in registry.
- **Optional flags:** Not recorded in registry.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - unsupported

## Other registered surfaces

Entries in this section: **45**

### `comms.dev-create-invite`

- **Command:** `go run ./cmd/comms dev-create-invite`
- **Repo:** `carbonstack-comms`
- **Component:** `cmd/comms`
- **Kind:** `cli`
- **Audience:** `dev`
- **Maturity:** `dev_only`
- **Introduced in:** `Not recorded in registry.`
- **Source path:** `carbonstack-comms/internal/app/commands.go`
- **Validation surface:** runtime smoke scripts
- **Front README candidate:** `false`

**What it does:** Create a development invite through the configured Cypher server.

**Why it exists:** Supports local/dev account onboarding before mature invite UX exists.

- **Required flags:** Not recorded in registry.
- **Optional flags:** Not recorded in registry.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - dev helper only

### `comms.relay-space-invite-claim-dev`

- **Command:** `go run ./cmd/comms relay-space-invite-claim-dev`
- **Repo:** `carbonstack-comms`
- **Component:** `cmd/comms`
- **Kind:** `cli`
- **Audience:** `dev`
- **Maturity:** `dev_only`
- **Lifecycle status:** `active`
- **Introduced in:** `v0.7.4`
- **Source path:** `carbonstack-comms/internal/app/relay_space_invite_claim_dev.go`
- **Validation surface:** Comms client/app tests and runner.relay-space-invite-claim-dev
- **Front README candidate:** `false`

**What it does:** Claim a Relay Space invite for the existing account and device in an explicit local Comms state file.

**Why it exists:** Provides the participant-side dev/operator leaf for the atomic Cypher Relay Space invite claim route without creating accounts, registering devices, mutating trust, or implying OpenMLS membership.

- **Required flags:**
  - `--state` — Explicit local Comms state path containing the existing account_id, device_id, and server context.
  - `--invite-token` — Full Relay Space invite token; display_code and word_code are not accepted as lookup credentials.
- **Optional flags:**
  - `--server` — Explicit Cypher server URL override; otherwise the selected state file controls the server.
  - `--display-label` — Routing-member display label; defaults to the selected state device label.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - not account creation
  - not device registration
  - not Relay Space administration
  - not identity verification
  - not trust promotion
  - not OpenMLS group membership
  - not secure enrollment
  - not production membership UX
  - not authenticated public API safety
  - not deployment
  - not local-backbone
  - not audit or certification

### `comms.relay-space-member-state-dev`

- **Command:** `go run ./cmd/comms relay-space-member-state-dev`
- **Repo:** `carbonstack-comms`
- **Component:** `cmd/comms`
- **Kind:** `cli`
- **Audience:** `dev`
- **Maturity:** `dev_only`
- **Lifecycle status:** `active`
- **Introduced in:** `v0.7.4`
- **Source path:** `carbonstack-comms/internal/app/relay_space_member_state_dev.go`
- **Validation surface:** Comms client/app tests and runner.relay-space-member-state-dev
- **Front README candidate:** `false`

**What it does:** Explicitly transition one Relay Space routing member between active, disabled, and left states.

**Why it exists:** Provides the bounded dev/operator surface for the Cypher routing-member state authority while preserving idempotence, routing enforcement, left-member rejoin refusal, and strict identity/trust/OpenMLS nonclaims.

- **Required flags:**
  - `--state` — Explicit local Comms state path controlling the Cypher server context; the command does not rewrite the file.
  - `--relay-space-id` — Authoritative Relay Space ID containing the routing member.
  - `--routing-member-id` — Routing member ID to transition.
  - `--target-state` — Target state active, disabled, or left.
- **Optional flags:**
  - `--server` — Explicit Cypher server URL override; otherwise the selected state file controls the server.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - not authenticated Relay Space administration
  - not production authorization
  - not caller identity proof
  - not identity verification
  - not trust promotion
  - not OpenMLS group membership mutation
  - not member deletion
  - not an explicit rejoin workflow
  - not secure enrollment
  - not production membership UX
  - not public API safety
  - not deployment
  - not local-backbone
  - not audit or certification

### `comms.simulate-key-change`

- **Command:** `go run ./cmd/comms simulate-key-change`
- **Repo:** `carbonstack-comms`
- **Component:** `cmd/comms`
- **Kind:** `cli`
- **Audience:** `dev`
- **Maturity:** `dev_only`
- **Introduced in:** `Not recorded in registry.`
- **Source path:** `carbonstack-comms/internal/app/commands.go`
- **Validation surface:** trust lifecycle tests
- **Front README candidate:** `false`

**What it does:** Simulate a key-change event for trust lifecycle testing.

**Why it exists:** Negative/test helper for trust policy behavior.

- **Required flags:** Not recorded in registry.
- **Optional flags:** Not recorded in registry.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - dev/test helper only

### `comms.state-audit-dev`

- **Command:** `go run ./cmd/comms state-audit-dev`
- **Repo:** `carbonstack-comms`
- **Component:** `cmd/comms`
- **Kind:** `cli`
- **Audience:** `dev`
- **Maturity:** `dev_only`
- **Introduced in:** `v0.5.60`
- **Source path:** `carbonstack-comms/internal/app/commands.go`
- **Validation surface:** Comms package tests and command registry coverage
- **Front README candidate:** `false`

**What it does:** Report non-mutating local state-domain inventory for Comms, trust/candidate state, OpenMLS sidecar generated state, build output, and local Cypher DB boundaries.

**Why it exists:** First executable state/vault-spine primitive; makes local state domains visible without printing secrets, mutating trust, deleting artifacts, or claiming production vault readiness.

- **Required flags:** Not recorded in registry.
- **Optional flags:** Not recorded in registry.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - not vault encryption
  - not production key storage
  - not recovery
  - not deletion or cleanup
  - not trust verification
  - not local-backbone
  - not production secure messaging

### `comms.state-path-policy-dev`

- **Command:** `go run ./cmd/comms state-path-policy-dev`
- **Repo:** `carbonstack-comms`
- **Component:** `cmd/comms`
- **Kind:** `cli`
- **Audience:** `dev`
- **Maturity:** `dev_only`
- **Lifecycle status:** `gate_c_path_policy`
- **Introduced in:** `Not recorded in registry.`
- **Source path:** `carbonstack-comms/internal/app/state_path_policy_dev.go`
- **Validation surface:** state-path-policy-dev and Comms package tests
- **Front README candidate:** `false`

**What it does:** Classify Comms state-root/path policy, supported overrides, derived roots, and external state authority boundaries without migration or relocation.

**Why it exists:** Provides the Gate C3 path policy/refusal surface before C4 atomic write, lock, partial-state, and replay behavior.

- **Required flags:** Not recorded in registry.
- **Optional flags:**
  - `--state` — local Comms state path; explicit compatibility is preserved.
  - `--state-root` — optional Comms-owned root override for path-policy classification.
  - `--sidecar-dir` — OpenMLS sidecar directory classified as sidecar-owned state.
  - `--cypher-db` — optional Cypher DB path classified as Cypher-owned state.
  - `--validator-temp-root` — optional validator temp root classified as generated validation state.
  - `--evidence-root` — optional evidence root classified as evidence-only.
  - `--output` — optional generated evidence report path.
  - `--allow-refusal-exit-zero` — print refusal classification but exit zero for validator/profile use.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - not migration
  - not silent repair
  - not state relocation
  - not cleanup implementation
  - not C4 atomicity or lock closure
  - not C5 Gate C closure
  - not vault security
  - not backup restore
  - not trust promotion
  - not verified identity
  - not Cypher/MLS reconciliation
  - not deployment
  - not full-runtime-dev
  - not Gate D runtime aggregate
  - not production E2EE

### `comms.state-schema-compat-dev`

- **Command:** `go run ./cmd/comms state-schema-compat-dev`
- **Repo:** `carbonstack-comms`
- **Component:** `cmd/comms`
- **Kind:** `cli`
- **Audience:** `dev`
- **Maturity:** `dev_only`
- **Lifecycle status:** `gate_c_schema_compatibility`
- **Introduced in:** `Not recorded in registry.`
- **Source path:** `carbonstack-comms/internal/app/state_schema_compat_dev.go`
- **Validation surface:** state-schema-compat-dev and Comms package tests
- **Front README candidate:** `false`

**What it does:** Classify and refuse unsupported schema versions for Comms-owned JSON state/report/receipt artifacts without migration or repair.

**Why it exists:** Provides the Gate C2 compatibility/refusal surface before deeper path policy, atomic write, migration, or runtime aggregate work.

- **Required flags:** Not recorded in registry.
- **Optional flags:**
  - `--kind` — schema kind to classify, such as keypackage-receipt, welcome-receipt, workflow-report, state-substrate-inventory, cypher-mls-mismatch-report, or comms-state.
  - `--path` — JSON file path to inspect.
  - `--output` — optional generated evidence report path.
  - `--allow-refusal-exit-zero` — print refusal classification but exit zero for validator/profile use.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - not migration
  - not silent repair
  - not state relocation
  - not deletion or cleanup
  - not vault security
  - not backup restore
  - not trust promotion
  - not verified identity
  - not Cypher/MLS reconciliation
  - not deployment
  - not full-runtime-dev
  - not Gate D runtime aggregate
  - not production E2EE

### `comms.state-substrate-inventory-dev`

- **Command:** `go run ./cmd/comms state-substrate-inventory-dev`
- **Repo:** `carbonstack-comms`
- **Component:** `cmd/comms`
- **Kind:** `cli`
- **Audience:** `dev`
- **Maturity:** `dev_only`
- **Lifecycle status:** `gate_c_state_substrate_inventory`
- **Introduced in:** `Not recorded in registry.`
- **Source path:** `carbonstack-comms/internal/app/state_substrate_inventory_dev.go`
- **Validation surface:** state-substrate-inventory-dev and Comms package tests
- **Front README candidate:** `false`

**What it does:** Produce a read-only-by-default inventory of Comms, sidecar, Cypher, evidence, receipt, workflow, trust, and candidate state boundaries.

**Why it exists:** Opens Gate C implementation with a machine-readable authority map before schema enforcement, migration, repair, vault, backup, deployment, or runtime aggregation work.

- **Required flags:** Not recorded in registry.
- **Optional flags:**
  - `--state` — local Comms state path; remains supported and does not require canonical-root migration.
  - `--state-root` — optional canonical Comms-owned state root override for inventory classification.
  - `--sidecar-dir` — OpenMLS sidecar directory to classify as sidecar-owned state.
  - `--cypher-db` — optional Cypher DB path to classify as Cypher-owned server state.
  - `--evidence-root` — optional evidence root to classify as generated evidence, not runtime state.
  - `--output` — optional machine-readable inventory report path; generated evidence only.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - not schema enforcement
  - not migration
  - not silent repair
  - not deletion or cleanup
  - not vault security
  - not backup restore
  - not trust promotion
  - not verified identity
  - not Cypher/MLS reconciliation
  - not deployment
  - not full-runtime-dev
  - not Gate D runtime aggregate
  - not production E2EE

### `comms.state-write-policy-dev`

- **Command:** `go run ./cmd/comms state-write-policy-dev`
- **Repo:** `carbonstack-comms`
- **Component:** `cmd/comms`
- **Kind:** `cli`
- **Audience:** `dev`
- **Maturity:** `dev_only`
- **Lifecycle status:** `gate_c_write_policy`
- **Introduced in:** `Not recorded in registry.`
- **Source path:** `carbonstack-comms/internal/app/state_write_policy_dev.go`
- **Validation surface:** state-write-policy-dev and Comms package tests
- **Front README candidate:** `false`

**What it does:** Classify atomic write, lock, partial-state, replay, and cleanup-boundary policy for Comms-owned state surfaces without migration or writer rewiring.

**Why it exists:** Provides the Gate C4 write-policy classification surface before C5 Gate C closure.

- **Required flags:** Not recorded in registry.
- **Optional flags:**
  - `--state-root` — Comms-owned state root to use when describing write policy.
  - `--sidecar-dir` — OpenMLS sidecar directory classified as sidecar-owned state.
  - `--cypher-db` — optional Cypher DB path classified as Cypher-owned state.
  - `--validator-temp-root` — optional validator temp root classified as generated validation state.
  - `--evidence-root` — optional evidence root classified as evidence-only.
  - `--output` — optional generated evidence report path.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - not migration
  - not silent repair
  - not state relocation
  - not runtime writer rewiring
  - not cleanup implementation
  - not destructive cleanup
  - not C5 Gate C closure
  - not vault security
  - not backup restore
  - not trust promotion
  - not verified identity
  - not Cypher/MLS reconciliation
  - not deployment
  - not full-runtime-dev
  - not Gate D runtime aggregate
  - not production E2EE

### `comms.workflow-relay-onboarding-dev`

- **Command:** `go run ./cmd/comms workflow-relay-onboarding-dev`
- **Repo:** `carbonstack-comms`
- **Component:** `cmd/comms`
- **Kind:** `cli`
- **Audience:** `dev`
- **Maturity:** `dev_only`
- **Lifecycle status:** `relay_onboarding_workflow_evaluator`
- **Introduced in:** `Not recorded in registry.`
- **Source path:** `carbonstack-comms/internal/app/workflow_relay_onboarding_dev.go`
- **Validation surface:** workflow-relay-onboarding-dev and Comms package tests
- **Front README candidate:** `false`

**What it does:** Evaluate Relay onboarding workflow readiness while preserving B4/B5/B6/B7 leaf boundaries.

**Why it exists:** Closes Gate B8 as a reusable noninteractive workflow report/evaluator without hiding receipts, ACK discipline, or mismatch refusal.

- **Required flags:**
  - `--relay-space` — Relay Space ID under workflow evaluation.
  - `--cypher-member-state` — explicit Cypher Relay Space member-state snapshot.
- **Optional flags:**
  - `--state` — local Comms state path.
  - `--workflow-id` — stable workflow report ID.
  - `--mls-group-state` — local MLS group state snapshot or auto.
  - `--sidecar-dir` — OpenMLS sidecar directory.
  - `--sidecar-device-label` — OpenMLS sidecar device label.
  - `--conversation` — OpenMLS conversation label.
  - `--conversation-state` — explicit OpenMLS conversation-state path.
  - `--keypackage-receipt` — optional B5d KeyPackage receipt manifest.
  - `--welcome-receipt` — optional B6 Welcome receipt manifest.
  - `--report-root` — local workflow report root.
  - `--allow-refusal-exit-zero` — print refusal report but return exit 0 for reporting harnesses.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - not silent repair
  - not silent rejoin
  - not trust promotion
  - not verified identity
  - not Cypher/MLS reconciliation
  - not B9 Gate B closure
  - not full-runtime-dev
  - not production E2EE

### `runner.cypher-mls-mismatch-dev`

- **Command:** `go run . --profile cypher-mls-mismatch-dev`
- **Repo:** `carbonstack`
- **Component:** `tools/carbonstack-validate`
- **Kind:** `runner-profile`
- **Audience:** `dev`
- **Maturity:** `dev_only`
- **Lifecycle status:** `active`
- **Introduced in:** `Not recorded in registry.`
- **Source path:** `carbonstack/tools/carbonstack-validate/cypher_mls_mismatch_dev.go`
- **Validation surface:** cypher-mls-mismatch-dev-profile
- **Front README candidate:** `false`

**What it does:** Validate Gate B7 Cypher/MLS mismatch classification and refusal behavior.

**Why it exists:** Proves B7 mismatch inspection/refusal while preserving B5/B6 lifecycle semantics and nonclaims.

- **Required flags:** Not recorded in registry.
- **Optional flags:**
  - `--compact-summary` — print compact evidence where supported.
  - `--clean-generated` — remove known generated/build artifacts after successful profile execution.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - not silent repair
  - not silent rejoin
  - not trust promotion
  - not verified identity
  - not B8 workflow engine
  - not B9 Gate B closure
  - not production E2EE

### `runner.full-operational-spine-dev`

- **Command:** `go run . --profile full-operational-spine-dev`
- **Repo:** `carbonstack`
- **Component:** `tools/carbonstack-validate`
- **Kind:** `runner-profile`
- **Audience:** `dev`
- **Maturity:** `future`
- **Introduced in:** `v0.7.1`
- **Source path:** `docs/234-v0.7.1-gate-a-operational-workflow-contract-v0.md`
- **Validation surface:** documentation and registry placeholder only; not implemented
- **Front README candidate:** `false`

**What it does:** Future preferred full operational-spine lifecycle aggregate after the required workflow leaves are stable.

**Why it exists:** Preserves the accepted Gate A requirement that the preferred lifecycle has a command-table-visible full-prefixed classification while leaf and internal workflows remain documented, callable, and separately testable.

- **Required flags:** Not recorded in registry.
- **Optional flags:** Not recorded in registry.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - future placeholder only
  - command is not implemented
  - not included in full
  - not included in full-validate-release
  - not release-package validation
  - does not replace documented or callable leaf profiles
  - not deployment
  - not production secure messaging
  - not hostile-server safety
  - not adversarial campaign evidence

### `runner.full-validate-release`

- **Command:** `go run . --profile full-validate-release`
- **Repo:** `carbonstack`
- **Component:** `tools/carbonstack-validate`
- **Kind:** `runner-profile`
- **Audience:** `public`
- **Maturity:** `release_supported`
- **Introduced in:** `v0.6.30`
- **Source path:** `tools/carbonstack-validate/main.go`
- **Validation surface:** release package validation ladder
- **Front README candidate:** `true`

**What it does:** Run the explicit release-package validation profile; exact alias to current full behavior.

**Why it exists:** Preferred explicit name for the release/package-root validation ladder previously exposed as full; prevents broad full naming from implying live-dev, deployment, or production security validation.

- **Required flags:** Not recorded in registry.
- **Optional flags:**
  - `--root`
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - not deployment
  - not local-backbone
  - not runtime Comms UX
  - not production security proof
  - not full-runtime-dev
  - not live-dev aggregation
  - not adversarial harness
  - not package publisher

### `runner.gate-b-relay-lifecycle-closure-dev`

- **Command:** `go run . --profile gate-b-relay-lifecycle-closure-dev`
- **Repo:** `carbonstack`
- **Component:** `tools/carbonstack-validate`
- **Kind:** `runner-profile`
- **Audience:** `dev`
- **Maturity:** `dev_only`
- **Lifecycle status:** `gate_b_closure`
- **Introduced in:** `Not recorded in registry.`
- **Source path:** `carbonstack/tools/carbonstack-validate/gate_b_relay_lifecycle_closure_dev.go`
- **Validation surface:** gate-b-relay-lifecycle-closure-dev-profile
- **Front README candidate:** `false`

**What it does:** Validate Gate B Relay lifecycle closure across B1-B8 without promoting production, trust, identity, deployment, or Gate C claims.

**Why it exists:** Closes Gate B as a coherent dev/pre-alpha Relay lifecycle integration lane while preserving every leaf boundary and nonclaim.

- **Required flags:** Not recorded in registry.
- **Optional flags:**
  - `--compact-summary` — print compact evidence where supported.
  - `--clean-generated` — remove known generated/build artifacts after successful profile execution.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - not production E2EE
  - not verified identity
  - not trust promotion
  - not secure enrollment
  - not hostile-server safety
  - not Gate C state enforcement
  - not Gate D runtime aggregate
  - not Gate E native deployment
  - not Gate F v0.8.0 maturity
  - not full-runtime-dev
  - not deployment readiness
  - not container packaging
  - not PQ or hybrid support
  - not vault backup restore
  - not Android
  - not CarbonStackOS implementation

### `runner.gate-c-state-substrate-closure-dev`

- **Command:** `go run . --profile gate-c-state-substrate-closure-dev`
- **Repo:** `carbonstack`
- **Component:** `tools/carbonstack-validate`
- **Kind:** `runner-profile`
- **Audience:** `dev`
- **Maturity:** `dev_only`
- **Lifecycle status:** `active`
- **Introduced in:** `Not recorded in registry.`
- **Source path:** `carbonstack/tools/carbonstack-validate/gate_c_state_substrate_closure_dev.go`
- **Validation surface:** gate-c-state-substrate-closure-dev-profile
- **Front README candidate:** `false`

**What it does:** Validate Gate C state-substrate closure across C1 inventory, C2 schema compatibility, C3 path policy, and C4 write policy without opening Gate D.

**Why it exists:** Closes Gate C as a bounded dev/pre-alpha state-substrate enforcement lane and blocks Gate D until a fresh Gate D contract is accepted.

- **Required flags:** Not recorded in registry.
- **Optional flags:**
  - `--compact-summary` — print compact evidence where supported.
  - `--clean-generated` — remove known generated/build artifacts after successful profile execution.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - not Gate D runtime aggregate
  - not full-runtime-dev
  - not migration
  - not silent repair
  - not state relocation
  - not runtime writer rewiring
  - not cleanup implementation
  - not destructive cleanup
  - not vault security
  - not backup restore
  - not trust promotion
  - not verified identity
  - not Cypher/MLS reconciliation
  - not deployment
  - not production E2EE
  - not PQ or hybrid migration
  - not Android
  - not CarbonStackOS implementation

### `runner.gate-d-runtime-aggregate-dev`

- **Command:** `go run . --profile gate-d-runtime-aggregate-dev`
- **Repo:** `carbonstack`
- **Component:** `tools/carbonstack-validate`
- **Kind:** `runner-profile`
- **Audience:** `dev`
- **Maturity:** `dev_only`
- **Lifecycle status:** `active`
- **Introduced in:** `Not recorded in registry.`
- **Source path:** `carbonstack/tools/carbonstack-validate/gate_d_runtime_aggregate_dev.go`
- **Validation surface:** gate-d-runtime-aggregate-dev-profile
- **Front README candidate:** `false`

**What it does:** Validate the Gate D mechanical runtime aggregate across Gate C state preflight, workflow onboarding, restart/resume inspection, and minimal normal-message send/inbox proof without promoting full-runtime-dev.

**Why it exists:** Closes Gate D as a bounded dev/pre-alpha runtime aggregate while reserving full-runtime-dev until a later explicit promotion decision.

- **Required flags:** Not recorded in registry.
- **Optional flags:**
  - `--compact-summary` — print compact evidence where supported.
  - `--clean-generated` — remove known generated/build artifacts after successful profile execution.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - not full-runtime-dev
  - not Gate E native deployment
  - not production readiness
  - not production E2EE
  - not verified identity
  - not trust promotion
  - not vault security
  - not backup restore
  - not public ingress
  - not TUI
  - not Android
  - not CarbonStackOS implementation
  - not PQ or hybrid migration
  - not v0.8.0 release readiness

### `runner.gate-e-native-deployment-closure-dev`

- **Command:** `go run . --profile gate-e-native-deployment-closure-dev`
- **Repo:** `carbonstack`
- **Component:** `tools/carbonstack-validate`
- **Kind:** `runner-profile`
- **Audience:** `dev`
- **Maturity:** `dev_only`
- **Lifecycle status:** `active`
- **Introduced in:** `Not recorded in registry.`
- **Source path:** `carbonstack/tools/carbonstack-validate/gate_e_native_deployment_closure_dev.go`
- **Validation surface:** gate-e-native-deployment-closure-profile
- **Front README candidate:** `false`

**What it does:** Close Gate E manual-private native deployment by validating Gate E E1, Cypher terminating config inspection, registry references, and nonclaims.

**Why it exists:** Establishes the bounded Gate E closure surface before Gate F preflight while preventing service/systemd/helper, public ingress, container, TUI, full-runtime-dev, trust, identity, vault, backup, PQ, Android, and OS scope creep.

- **Required flags:** Not recorded in registry.
- **Optional flags:**
  - `--compact-summary` — print compact evidence where supported.
  - `--clean-generated` — remove known generated/build artifacts after successful profile execution.
- **Environment:** Not recorded in registry.
- **Related registry rows:**
  - runner.gate-e-native-deployment-dev
  - cypher.config-inspection
  - runner.gate-d-runtime-aggregate-dev
- **Not claims:**
  - not Gate F
  - not v0.8.0 release readiness
  - not semi-persistent service
  - not systemd
  - not helper install
  - not public ingress
  - not container readiness
  - not TUI
  - not full-runtime-dev
  - not production readiness
  - not production E2EE
  - not verified identity
  - not trust promotion
  - not vault security
  - not backup restore
  - not PQ or hybrid migration
  - not Android
  - not CarbonStackOS implementation

### `runner.gate-e-native-deployment-dev`

- **Command:** `go run . --profile gate-e-native-deployment-dev`
- **Repo:** `carbonstack`
- **Component:** `tools/carbonstack-validate`
- **Kind:** `runner-profile`
- **Audience:** `dev`
- **Maturity:** `dev_only`
- **Lifecycle status:** `active`
- **Introduced in:** `Not recorded in registry.`
- **Source path:** `carbonstack/tools/carbonstack-validate/gate_e_native_deployment_dev.go`
- **Validation surface:** gate-e-e1-native-deployment-profile
- **Front README candidate:** `false`

**What it does:** Validate Gate E E1 manual-private native deployment context with explicit-env Cypher start/stop/restart and Gate C state-policy inspection over deployment roots.

**Why it exists:** Establishes the first bounded native deployment proof after Gate D while keeping service/systemd, public ingress, containers, TUI, full-runtime-dev, trust, identity, vault, backup, PQ, Android, and OS work out of scope.

- **Required flags:** Not recorded in registry.
- **Optional flags:**
  - `--compact-summary` — print compact evidence where supported.
  - `--clean-generated` — remove known generated/build artifacts after successful profile execution.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - not Gate E final closure by itself
  - not semi-persistent service
  - not systemd
  - not helper install
  - not public ingress
  - not container readiness
  - not TUI
  - not full-runtime-dev
  - not production readiness
  - not production E2EE
  - not verified identity
  - not trust promotion
  - not vault security
  - not backup restore
  - not PQ or hybrid migration
  - not Android
  - not CarbonStackOS implementation
  - not v0.8.0 release readiness

### `runner.gate-f-basic-local-trust-posture-dev`

- **Command:** `go run . --profile gate-f-basic-local-trust-posture-dev`
- **Repo:** `carbonstack`
- **Component:** `tools/carbonstack-validate`
- **Kind:** `runner-profile`
- **Audience:** `dev`
- **Maturity:** `dev_only`
- **Lifecycle status:** `active`
- **Introduced in:** `Not recorded in registry.`
- **Source path:** `carbonstack/tools/carbonstack-validate/gate_f_basic_local_trust_posture_dev.go`
- **Validation surface:** gate-f-f5-basic-local-trust-candidate-posture
- **Front README candidate:** `false`

**What it does:** Validate Gate F F5 basic local manual trust candidate posture without verified identity, trust promotion, secure enrollment, cryptographic binding, package/runtime candidate, or release work.

**Why it exists:** Closes the minimal v0.7.x local trust posture so v0.8.0 can honestly state verified identity remains a nonclaim while a basic local acceptance model exists for future expansion.

- **Required flags:** Not recorded in registry.
- **Optional flags:**
  - `--compact-summary` — print compact evidence where supported.
- **Environment:** Not recorded in registry.
- **Related registry rows:**
  - comms.basic-local-trust-posture-dev
  - comms.basic-local-trust-accept-dev
  - comms.message-send-dev
  - comms.openmls-relay-join-dev
  - runner.gate-f-code-health-source-hygiene-dev
- **Not claims:**
  - not verified identity
  - not full trust promotion
  - not secure enrollment
  - not server-hostile identity replacement proof
  - not real-world person verification
  - not cryptographic binding across Cypher Comms and OpenMLS identities
  - not automatic trust promotion
  - not trust from Relay membership
  - not trust from successful Welcome or MLS join
  - not package-runtime candidate
  - not v0.8.0 release readiness
  - not release creation
  - not release upload
  - not package publication
  - not package staging execution
  - not full-runtime-dev
  - not migration implementation
  - not repair implementation
  - not destructive cleanup
  - not state relocation
  - not service or systemd
  - not helper install
  - not public ingress
  - not container readiness
  - not TUI
  - not production E2EE
  - not vault security
  - not backup restore
  - not PQ or hybrid migration
  - not Android
  - not CarbonStackOS implementation

### `runner.gate-f-code-health-source-hygiene-dev`

- **Command:** `go run . --profile gate-f-code-health-source-hygiene-dev`
- **Repo:** `carbonstack`
- **Component:** `tools/carbonstack-validate`
- **Kind:** `runner-profile`
- **Audience:** `dev`
- **Maturity:** `dev_only`
- **Lifecycle status:** `active`
- **Introduced in:** `Not recorded in registry.`
- **Source path:** `carbonstack/tools/carbonstack-validate/gate_f_code_health_source_hygiene_dev.go`
- **Validation surface:** gate-f-f4-code-health-source-hygiene-classification
- **Front README candidate:** `false`

**What it does:** Validate Gate F F4 code-health and source-hygiene classification without package/runtime candidate, release creation, migration, repair, cleanup, or full-runtime-dev.

**Why it exists:** Closes a bounded source-hygiene layer after F3 observability so later package/runtime candidate work starts from clean generated-cache policy, non-destructive cypher.db classification, and static helper safety classification.

- **Required flags:** Not recorded in registry.
- **Optional flags:**
  - `--compact-summary` — print compact evidence where supported.
- **Environment:** Not recorded in registry.
- **Related registry rows:**
  - runner.gate-f-compat-rollback-observability-dev
  - runner.gate-f-release-package-surface-dev
  - runner.gate-f-operator-docs-runbook-dev
  - runner.release-snapshot
  - runner.full-validate-release
- **Not claims:**
  - not v0.8.0 release readiness
  - not release creation
  - not release upload
  - not package publication
  - not package staging execution
  - not package-runtime candidate
  - not full-runtime-dev
  - not migration implementation
  - not repair implementation
  - not destructive cleanup
  - not state relocation
  - not service or systemd
  - not helper install
  - not public ingress
  - not container readiness
  - not TUI
  - not production readiness
  - not production E2EE
  - not verified identity
  - not trust promotion
  - not vault security
  - not backup restore
  - not PQ or hybrid migration
  - not Android
  - not CarbonStackOS implementation

### `runner.gate-f-compat-rollback-observability-dev`

- **Command:** `go run . --profile gate-f-compat-rollback-observability-dev`
- **Repo:** `carbonstack`
- **Component:** `tools/carbonstack-validate`
- **Kind:** `runner-profile`
- **Audience:** `dev`
- **Maturity:** `dev_only`
- **Lifecycle status:** `active`
- **Introduced in:** `Not recorded in registry.`
- **Source path:** `carbonstack/tools/carbonstack-validate/gate_f_compat_rollback_observability_dev.go`
- **Validation surface:** gate-f-f3-compatibility-rollback-observability-refusal-posture
- **Front README candidate:** `false`

**What it does:** Validate Gate F F3 aggregate compatibility, stale-state, rollback observability, and refusal posture without implementing migration, repair, cleanup, package/runtime candidate, or full-runtime-dev.

**Why it exists:** Reuses Gate C compatibility/path/write-policy authorities, Gate F F1/F2 release/runbook authorities, and Cypher config inspection to prove current observability/refusal posture before later code-health and package/runtime work.

- **Required flags:** Not recorded in registry.
- **Optional flags:**
  - `--compact-summary` — print compact evidence where supported.
- **Environment:** Not recorded in registry.
- **Related registry rows:**
  - runner.state-schema-compat-dev
  - runner.state-path-policy-dev
  - runner.state-write-policy-dev
  - runner.gate-f-release-package-surface-dev
  - runner.gate-f-operator-docs-runbook-dev
  - cypher.config-inspection
- **Not claims:**
  - not v0.8.0 release readiness
  - not release creation
  - not release upload
  - not package publication
  - not package staging execution
  - not package-runtime candidate
  - not full-runtime-dev
  - not migration implementation
  - not repair implementation
  - not destructive cleanup
  - not state relocation
  - not service or systemd
  - not helper install
  - not public ingress
  - not container readiness
  - not TUI
  - not production readiness
  - not production E2EE
  - not verified identity
  - not trust promotion
  - not vault security
  - not backup restore
  - not PQ or hybrid migration
  - not Android
  - not CarbonStackOS implementation

### `runner.gate-f-operator-docs-runbook-dev`

- **Command:** `go run . --profile gate-f-operator-docs-runbook-dev`
- **Repo:** `carbonstack`
- **Component:** `tools/carbonstack-validate`
- **Kind:** `runner-profile`
- **Audience:** `dev`
- **Maturity:** `dev_only`
- **Lifecycle status:** `active`
- **Introduced in:** `Not recorded in registry.`
- **Source path:** `carbonstack/tools/carbonstack-validate/gate_f_operator_docs_runbook_dev.go`
- **Validation surface:** gate-f-f2-operator-docs-runbook-closure
- **Front README candidate:** `false`

**What it does:** Validate Gate F F2 manual-private lifecycle, config/env, release/package authority, and failure/refusal/hygiene operator documentation.

**Why it exists:** Closes the operator documentation and runbook layer after Gate F F1 surface classification and before later compatibility, package/runtime, or release-candidate work.

- **Required flags:** Not recorded in registry.
- **Optional flags:**
  - `--compact-summary` — print compact evidence where supported.
- **Environment:** Not recorded in registry.
- **Related registry rows:**
  - runner.gate-f-release-package-surface-dev
  - runner.gate-e-native-deployment-closure-dev
  - cypher.config-inspection
  - runner.full-validate-release
  - runner.release-snapshot
- **Not claims:**
  - not v0.8.0 release readiness
  - not release creation
  - not release upload
  - not package publication
  - not package staging execution
  - not package-runtime candidate
  - not full-runtime-dev
  - not migration implementation
  - not service or systemd
  - not helper install
  - not public ingress
  - not container readiness
  - not TUI
  - not production readiness
  - not production E2EE
  - not verified identity
  - not trust promotion
  - not vault security
  - not backup restore
  - not PQ or hybrid migration
  - not Android
  - not CarbonStackOS implementation

### `runner.gate-f-package-runtime-candidate-dev`

- **Command:** `go run . --profile gate-f-package-runtime-candidate-dev`
- **Repo:** `carbonstack`
- **Component:** `tools/carbonstack-validate`
- **Kind:** `runner-profile`
- **Audience:** `dev`
- **Maturity:** `dev_only`
- **Lifecycle status:** `active`
- **Introduced in:** `Not recorded in registry.`
- **Source path:** `carbonstack/tools/carbonstack-validate/gate_f_package_runtime_candidate_dev.go`
- **Validation surface:** gate-f-f6-package-runtime-candidate-validation
- **Front README candidate:** `false`

**What it does:** Validate a disposable package/runtime candidate root shape without release creation, release upload, package publication, legacy package staging execution, or full-runtime-dev promotion.

**Why it exists:** Closes the v0.7.x package/runtime candidate validation layer so v0.8.0 release-candidate closure can use a coherent root model and manual release handoff without pretending a public release has already been created.

- **Required flags:** Not recorded in registry.
- **Optional flags:**
  - `--compact-summary` — print compact evidence where supported.
- **Environment:** Not recorded in registry.
- **Related registry rows:**
  - runner.gate-f-release-package-surface-dev
  - runner.gate-f-operator-docs-runbook-dev
  - runner.gate-f-compat-rollback-observability-dev
  - runner.gate-f-code-health-source-hygiene-dev
  - runner.gate-f-basic-local-trust-posture-dev
  - comms.basic-local-trust-posture-dev
  - cypher.config-inspection
  - runner.release-snapshot
  - runner.full-validate-release
- **Not claims:**
  - not v0.8.0 release readiness
  - not release creation
  - not release upload
  - not package publication
  - not package staging execution
  - not public package artifact creation
  - not full-runtime-dev
  - not service or systemd
  - not helper install
  - not public ingress
  - not container readiness
  - not TUI
  - not migration implementation
  - not repair implementation
  - not destructive cleanup
  - not state relocation
  - not verified identity
  - not full trust promotion
  - not secure enrollment
  - not cryptographic identity binding
  - not vault security
  - not backup restore
  - not production E2EE
  - not PQ or hybrid migration
  - not Android
  - not CarbonStackOS implementation

### `runner.gate-f-release-package-surface-dev`

- **Command:** `go run . --profile gate-f-release-package-surface-dev`
- **Repo:** `carbonstack`
- **Component:** `tools/carbonstack-validate`
- **Kind:** `runner-profile`
- **Audience:** `dev`
- **Maturity:** `dev_only`
- **Lifecycle status:** `active`
- **Introduced in:** `Not recorded in registry.`
- **Source path:** `carbonstack/tools/carbonstack-validate/gate_f_release_package_surface_dev.go`
- **Validation surface:** gate-f-f1-release-package-runtime-surface-classification
- **Front README candidate:** `false`

**What it does:** Classify release validation, package staging, package rehearsal, runtime validation, helper, and hygiene surfaces before v0.8.0 package/runtime candidate work.

**Why it exists:** Prevents historical package scripts and validation profiles from being mistaken for v0.8.0 release authority while preserving manual release creation as an operator process.

- **Required flags:** Not recorded in registry.
- **Optional flags:**
  - `--compact-summary` — print compact evidence where supported.
- **Environment:** Not recorded in registry.
- **Related registry rows:**
  - runner.full-validate-release
  - runner.release-snapshot
  - runner.local-cypher
  - runner.integrated-runtime-dev
  - runner.gate-d-runtime-aggregate-dev
  - runner.gate-e-native-deployment-closure-dev
- **Not claims:**
  - not v0.8.0 release readiness
  - not release creation
  - not release upload
  - not package publication
  - not package staging execution
  - not package-runtime candidate
  - not full-runtime-dev
  - not migration implementation
  - not service or systemd
  - not helper install
  - not public ingress
  - not container readiness
  - not TUI
  - not production readiness
  - not production E2EE
  - not verified identity
  - not trust promotion
  - not vault security
  - not backup restore
  - not PQ or hybrid migration
  - not Android
  - not CarbonStackOS implementation

### `runner.keypackage-consume-dev`

- **Command:** `go run . --profile keypackage-consume-dev`
- **Repo:** `carbonstack`
- **Component:** `tools/carbonstack-validate`
- **Kind:** `runner-profile`
- **Audience:** `dev`
- **Maturity:** `dev_only`
- **Lifecycle status:** `active`
- **Introduced in:** `Not recorded in registry.`
- **Source path:** `carbonstack/tools/carbonstack-validate/keypackage_consume_dev.go`
- **Validation surface:** keypackage-consume-dev-profile
- **Front README candidate:** `false`

**What it does:** Validate Gate B5d KeyPackage consume/receipt and ACK-after-persist semantics.

**Why it exists:** Proves the B5d bounded delivery-consume leaf and supports B5e Gate B5 closure.

- **Required flags:** Not recorded in registry.
- **Optional flags:**
  - `--compact-summary` — print compact evidence where supported.
  - `--clean-generated` — remove known generated/build artifacts after successful profile execution.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - not add-member
  - not Welcome lifecycle
  - not identity verification
  - not trust promotion
  - not production E2EE
  - not B6

### `runner.keypackage-inspect-dev`

- **Command:** `go run . --profile keypackage-inspect-dev`
- **Repo:** `carbonstack`
- **Component:** `tools/carbonstack-validate`
- **Kind:** `runner-profile`
- **Audience:** `dev`
- **Maturity:** `dev_only`
- **Lifecycle status:** `active`
- **Introduced in:** `v0.7.5`
- **Source path:** `tools/carbonstack-validate/keypackage_inspect_dev.go`
- **Validation surface:** live-keypackage-generation-read-only-inspection-lifetime-local-owner-cross-owner-tamper-refusal
- **Front README candidate:** `false`

**What it does:** Prove material KeyPackage inspection, lifetime metadata, local ownership evidence, and read-only refusal behavior.

**Why it exists:** Permanently lock the Gate B5a foundation before repeatable generation and rotation begin in B5b.

- **Required flags:** Not recorded in registry.
- **Optional flags:**
  - `--root` — Explicit live umbrella root containing the CarbonStack repositories.
  - `--clean-generated` — Remove known generated OpenMLS build/state roots after successful validation.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - not included in full
  - not included in full-validate-release
  - not included in release-snapshot
  - not release-package validation
  - not repeatable KeyPackage generation
  - not KeyPackage rotation
  - not Relay publication
  - not KeyPackage consumption
  - not Welcome lifecycle
  - not account identity verification
  - not device identity verification
  - not Relay Space membership verification
  - not trust promotion
  - not secure enrollment
  - not local-backbone
  - not deployment
  - not public ingress safety
  - not mature messenger UX
  - not audit or certification

### `runner.keypackage-publication-dev`

- **Command:** `go run . --profile keypackage-publication-dev`
- **Repo:** `carbonstack`
- **Component:** `tools/carbonstack-validate`
- **Kind:** `runner-profile`
- **Audience:** `dev`
- **Maturity:** `dev_only`
- **Lifecycle status:** `active`
- **Introduced in:** `v0.7.6`
- **Source path:** `carbonstack/tools/carbonstack-validate/keypackage_publication_dev.go`
- **Validation surface:** b5b-generation-b5a-inspection-cypher-created-replay-conflicts-concurrency-restart-comms-selection
- **Front README candidate:** `false`

**What it does:** Prove the bounded B5c KeyPackage publication contract.

**Why it exists:** Permanently lock publication identity, duplicate/reuse classification, persistence, and no-consume/no-ACK boundaries before B5d.

- **Required flags:** Not recorded in registry.
- **Optional flags:**
  - `--root` — Explicit live umbrella root containing the CarbonStack repositories.
  - `--clean-generated` — Remove known generated OpenMLS build and state roots after successful validation.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - not included in full
  - not included in full-validate-release
  - not included in release-snapshot
  - not release-package validation
  - not KeyPackage consumption
  - not KeyPackage ACK
  - not add-member
  - not Welcome lifecycle
  - not identity verification
  - not trust promotion
  - not secure enrollment
  - not production secure messaging
  - not production E2EE
  - not local-backbone
  - not deployment
  - not public ingress safety
  - not mature messenger UX
  - not audit or certification

### `runner.keypackage-rotation-dev`

- **Command:** `go run . --profile keypackage-rotation-dev`
- **Repo:** `carbonstack`
- **Component:** `tools/carbonstack-validate`
- **Kind:** `runner-profile`
- **Audience:** `dev`
- **Maturity:** `dev_only`
- **Lifecycle status:** `active`
- **Introduced in:** `v0.7.5`
- **Source path:** `tools/carbonstack-validate/keypackage_rotation_dev.go`
- **Validation surface:** live-legacy-adoption-repeatable-generation-idempotence-retirement-concurrency-restart-inspection
- **Front README candidate:** `false`

**What it does:** Prove the persistent B5b KeyPackage generation and rotation lifecycle.

**Why it exists:** Permanently lock B5b before Relay publication begins in B5c.

- **Required flags:** Not recorded in registry.
- **Optional flags:**
  - `--root` — Explicit live umbrella root containing the CarbonStack repositories.
  - `--clean-generated` — Remove known generated OpenMLS build and state roots after successful validation.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - not included in full
  - not included in full-validate-release
  - not included in release-snapshot
  - not release-package validation
  - not Relay publication
  - not KeyPackage consumption
  - not KeyPackage ACK
  - not Welcome lifecycle
  - not account identity verification
  - not Relay Space membership verification
  - not trust promotion
  - not secure enrollment
  - not local-backbone
  - not deployment
  - not public ingress safety
  - not mature messenger UX
  - not audit or certification

### `runner.relay-space-delivery-authority-dev`

- **Command:** `go run . --profile relay-space-delivery-authority-dev`
- **Repo:** `carbonstack`
- **Component:** `tools/carbonstack-validate`
- **Kind:** `runner-profile`
- **Audience:** `dev`
- **Maturity:** `dev_only`
- **Lifecycle status:** `active`
- **Introduced in:** `v0.7.5`
- **Source path:** `tools/carbonstack-validate/relay_space_delivery_authority_dev.go`
- **Validation surface:** live-umbrella-active-membership-scoped-inbox-ack-disable-left-restart-reactivation
- **Front README candidate:** `false`

**What it does:** Prove scoped Relay Space inbox and ACK require current active recipient membership while queued envelopes remain persisted.

**Why it exists:** Locks the Gate B5 prerequisite that disable and leave revoke fetch and ACK authority without deleting queued delivery state, and that explicit reactivation restores access.

- **Required flags:** Not recorded in registry.
- **Optional flags:**
  - `--root` — Explicit live umbrella root containing the CarbonStack repositories.
  - `--clean-generated` — Remove known generated OpenMLS build/state roots after successful validation where present.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - not included in full
  - not included in full-validate-release
  - not included in release-snapshot
  - not release-package validation
  - not KeyPackage generation
  - not KeyPackage inspection
  - not KeyPackage consumption
  - not Welcome lifecycle
  - not OpenMLS group membership mutation
  - not identity verification
  - not trust promotion
  - not authenticated administration
  - not envelope deletion
  - not secure enrollment
  - not local-backbone
  - not deployment
  - not public ingress safety
  - not mature messenger UX
  - not audit or certification

### `runner.relay-space-invite-claim-dev`

- **Command:** `go run . --profile relay-space-invite-claim-dev`
- **Repo:** `carbonstack`
- **Component:** `tools/carbonstack-validate`
- **Kind:** `runner-profile`
- **Audience:** `dev`
- **Maturity:** `dev_only`
- **Lifecycle status:** `active`
- **Introduced in:** `v0.7.4`
- **Source path:** `tools/carbonstack-validate/relay_space_invite_claim_dev.go`
- **Validation surface:** live-umbrella-relay-space-invite-claim-created-idempotent-no-local-state-mutation
- **Front README candidate:** `false`

**What it does:** Prove the Comms Relay Space invite-claim command creates one routing member, preserves idempotence, and does not rewrite local state.

**Why it exists:** Locks the Gate B4c participant/operator claim leaf to the atomic B4b Cypher route with deterministic created and already_active classifications, exact claim accounting, and strict identity/trust/OpenMLS nonclaims.

- **Required flags:** Not recorded in registry.
- **Optional flags:**
  - `--root` — Explicit live umbrella root containing the CarbonStack repositories.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - not included in full
  - not included in full-validate-release
  - not included in release-snapshot
  - not release-package validation
  - not production secure messaging
  - not production E2EE
  - not identity verification
  - not trust promotion
  - not OpenMLS group membership
  - not secure enrollment
  - not member disable, leave, or removal lifecycle
  - not local-backbone
  - not deployment
  - not public ingress safety
  - not mature messenger UX
  - not audit or certification

### `runner.relay-space-member-restart-dev`

- **Command:** `go run . --profile relay-space-member-restart-dev`
- **Repo:** `carbonstack`
- **Component:** `tools/carbonstack-validate`
- **Kind:** `runner-profile`
- **Audience:** `dev`
- **Maturity:** `dev_only`
- **Lifecycle status:** `active`
- **Introduced in:** `v0.7.4`
- **Source path:** `tools/carbonstack-validate/relay_space_member_restart_dev.go`
- **Validation surface:** live-umbrella-member-list-inspection-and-same-database-restart-persistence
- **Front README candidate:** `false`

**What it does:** Prove disabled and left Relay Space routing-member states survive Cypher restart and remain inspectable.

**Why it exists:** Permanently locks Gate B4 restart persistence, member-list inspection, routing refusal/restoration, left-member rejoin refusal, and local-state immutability before formal B4 closure.

- **Required flags:** Not recorded in registry.
- **Optional flags:**
  - `--root` — Explicit live umbrella root containing the CarbonStack repositories.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - not included in full
  - not included in full-validate-release
  - not included in release-snapshot
  - not release-package validation
  - not authenticated Relay Space administration
  - not production authorization
  - not production backup or restore
  - not rollback safety
  - not hostile-server safety
  - not identity verification
  - not trust promotion
  - not OpenMLS group membership mutation
  - not member deletion
  - not an explicit rejoin workflow
  - not secure enrollment
  - not local-backbone
  - not deployment
  - not public ingress safety
  - not mature messenger UX
  - not audit or certification

### `runner.relay-space-member-state-dev`

- **Command:** `go run . --profile relay-space-member-state-dev`
- **Repo:** `carbonstack`
- **Component:** `tools/carbonstack-validate`
- **Kind:** `runner-profile`
- **Audience:** `dev`
- **Maturity:** `dev_only`
- **Lifecycle status:** `active`
- **Introduced in:** `v0.7.4`
- **Source path:** `tools/carbonstack-validate/relay_space_member_state_dev.go`
- **Validation surface:** live-umbrella-disable-idempotence-reactivate-leave-rejoin-refusal-routing-enforcement
- **Front README candidate:** `false`

**What it does:** Prove the Comms member-state operator command controls routing authority and preserves the explicit rejoin boundary.

**Why it exists:** Locks the Gate B4d state-transition leaf to deterministic disable, reactivate, leave, idempotence, routing refusal/restoration, local-state immutability, and identity/trust/OpenMLS nonclaims.

- **Required flags:** Not recorded in registry.
- **Optional flags:**
  - `--root` — Explicit live umbrella root containing the CarbonStack repositories.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - not included in full
  - not included in full-validate-release
  - not included in release-snapshot
  - not release-package validation
  - not authenticated Relay Space administration
  - not production authorization
  - not production secure messaging
  - not production E2EE
  - not identity verification
  - not trust promotion
  - not OpenMLS group membership mutation
  - not member deletion
  - not an explicit rejoin workflow
  - not secure enrollment
  - not local-backbone
  - not deployment
  - not public ingress safety
  - not mature messenger UX
  - not audit or certification

### `runner.same-state-integrated-dev`

- **Command:** `go run . --profile same-state-integrated-dev`
- **Repo:** `carbonstack`
- **Component:** `tools/carbonstack-validate`
- **Kind:** `runner-profile`
- **Audience:** `dev`
- **Maturity:** `dev_only`
- **Introduced in:** `v0.6.16`
- **Source path:** `tools/carbonstack-validate/same_state_integrated_dev.go`
- **Validation surface:** live-umbrella-same-state-integrated-dev-positive-path
- **Front README candidate:** `false`

**What it does:** Prove Relay onboarding plus normal message send/open/ack in one coherent same-state live-dev universe.

**Why it exists:** Converts the v0.6.16A temporary same-state proof probe into a bounded runner profile without mutating full, release-snapshot, or the existing sequential integrated-runtime-dev profile.

- **Required flags:** Not recorded in registry.
- **Optional flags:**
  - `--root`
  - `--clean-generated`
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - not included in full
  - not included in release-snapshot
  - not release-package validation
  - not package-root validation
  - not production secure messaging
  - not production E2EE
  - not hostile-server safety
  - not metadata privacy
  - not identity verification
  - not secure enrollment
  - not local-backbone
  - not deployment
  - not mature messenger UX
  - not general-public UX
  - positive-path same-state proof only
  - not adversarial relay safety
  - not vault/key-storage safety
  - not PQ or hybrid security

### `runner.same-state-message-failure-dev`

- **Command:** `go run . --profile same-state-message-failure-dev`
- **Repo:** `carbonstack`
- **Component:** `tools/carbonstack-validate`
- **Kind:** `runner-profile`
- **Audience:** `dev`
- **Maturity:** `dev_only`
- **Introduced in:** `v0.6.17`
- **Source path:** `tools/carbonstack-validate/same_state_message_failure_dev.go`
- **Validation surface:** live-umbrella-same-state-normal-message-open-failure-noack
- **Front README candidate:** `false`

**What it does:** Prove wrong-conversation normal message open does not ack or drain the inbox after same-state Relay join.

**Why it exists:** First same-state failure-hardening companion profile after same-state-integrated-dev; locks the application-message ack rule for message-open failure without adding adversarial harness claims or changing full/release-snapshot.

- **Required flags:** Not recorded in registry.
- **Optional flags:**
  - `--root`
  - `--clean-generated`
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - not included in full
  - not included in release-snapshot
  - not release-package validation
  - not package-root validation
  - not adversarial relay harness
  - not hostile-server safety
  - not metadata privacy
  - not production secure messaging
  - not production E2EE
  - not identity verification
  - not secure enrollment
  - not local-backbone
  - not deployment
  - not mature messenger UX
  - not general-public UX
  - currently covers wrong-conversation message-open failure only
  - not vault/key-storage safety
  - not PQ or hybrid security

### `runner.same-state-message-malformed-payload-dev`

- **Command:** `go run . --profile same-state-message-malformed-payload-dev`
- **Repo:** `carbonstack`
- **Component:** `tools/carbonstack-validate`
- **Kind:** `runner-profile`
- **Audience:** `dev`
- **Maturity:** `dev_only`
- **Introduced in:** `v0.6.23`
- **Source path:** `tools/carbonstack-validate/same_state_message_malformed_payload_dev.go`
- **Validation surface:** live-umbrella-same-state-malformed-normal-message-payload-noopen-noack-nodrain
- **Front README candidate:** `false`

**What it does:** Prove malformed normal application-message payloads do not falsely open, ack, drain, mutate provider state, or rewrite envelope state.

**Why it exists:** Locks malformed normal-message payload failure behavior into a deterministic same-state live-dev profile before replay/duplicate classification and adversarial harness planning.

- **Required flags:** Not recorded in registry.
- **Optional flags:**
  - `--root`
  - `--clean-generated`
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - not included in full
  - not included in release-snapshot
  - not release-package validation
  - not package-root validation
  - not adversarial relay harness
  - not hostile-server safety
  - not metadata privacy
  - not production secure messaging
  - not production E2EE
  - not identity verification
  - not secure enrollment
  - not replay or duplicate classification
  - not local-backbone
  - not deployment
  - not mature messenger UX
  - not general-public UX
  - currently covers malformed normal application-message payload mutations only
  - does not cover unsupported content_type, which is handled by same-state-message-unsupported-dev
  - does not cover stale provider state, which is documented in v0.6.22
  - not vault/key-storage safety
  - not PQ or hybrid security

### `runner.same-state-message-recipient-failure-dev`

- **Command:** `go run . --profile same-state-message-recipient-failure-dev`
- **Repo:** `carbonstack`
- **Component:** `tools/carbonstack-validate`
- **Kind:** `runner-profile`
- **Audience:** `dev`
- **Maturity:** `dev_only`
- **Introduced in:** `v0.6.19`
- **Source path:** `tools/carbonstack-validate/same_state_message_recipient_failure_dev.go`
- **Validation surface:** live-umbrella-same-state-normal-message-wrong-recipient-device-no-false-success
- **Front README candidate:** `false`

**What it does:** Prove wrong recipient/device/sidecar attempts do not falsely open, ack, or drain Bob's inbox after same-state Relay join.

**Why it exists:** Third same-state failure-hardening companion profile after same-state-integrated-dev; locks no-false-success/no-ack/no-drain behavior for wrong recipient/device/sidecar combinations without adding adversarial harness claims or changing full/release-snapshot.

- **Required flags:** Not recorded in registry.
- **Optional flags:**
  - `--root`
  - `--clean-generated`
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - not included in full
  - not included in release-snapshot
  - not release-package validation
  - not package-root validation
  - not adversarial relay harness
  - not hostile-server safety
  - not metadata privacy
  - not production secure messaging
  - not production E2EE
  - not identity verification
  - not secure enrollment
  - not local-backbone
  - not deployment
  - not mature messenger UX
  - not general-public UX
  - currently covers wrong recipient/device/sidecar no-false-success only
  - not vault/key-storage safety
  - not PQ or hybrid security

### `runner.same-state-message-replay-classification-dev`

- **Command:** `go run . --profile same-state-message-replay-classification-dev`
- **Repo:** `carbonstack`
- **Component:** `tools/carbonstack-validate`
- **Kind:** `runner-profile`
- **Audience:** `dev`
- **Maturity:** `dev_only`
- **Introduced in:** `v0.6.24`
- **Source path:** `tools/carbonstack-validate/same_state_message_replay_classification_dev.go`
- **Validation surface:** live-umbrella-same-state-normal-message-replay-classification
- **Front README candidate:** `false`

**What it does:** Classify normal application-message duplicate/replay behavior without claiming replay safety.

**Why it exists:** Locks duplicate/replay classification into a deterministic same-state live-dev profile before adversarial harness and vault/PQ modeling work.

- **Required flags:** Not recorded in registry.
- **Optional flags:**
  - `--root`
  - `--clean-generated`
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - not included in full
  - not included in release-snapshot
  - not release-package validation
  - not package-root validation
  - not adversarial relay harness
  - not hostile-server safety
  - not replay safety
  - not metadata privacy
  - not production secure messaging
  - not production E2EE
  - not identity verification
  - not secure enrollment
  - not local-backbone
  - not deployment
  - not mature messenger UX
  - not general-public UX
  - currently covers normal application-message duplicate/replay cases only
  - does not cover Welcome replay
  - does not cover KeyPackage replay
  - does not cover server equivocation
  - does not cover network drop/delay
  - does not cover malicious relay harness behavior
  - not vault/key-storage safety
  - not PQ or hybrid security

### `runner.same-state-message-unsupported-dev`

- **Command:** `go run . --profile same-state-message-unsupported-dev`
- **Repo:** `carbonstack`
- **Component:** `tools/carbonstack-validate`
- **Kind:** `runner-profile`
- **Audience:** `dev`
- **Maturity:** `dev_only`
- **Introduced in:** `v0.6.18`
- **Source path:** `tools/carbonstack-validate/same_state_message_unsupported_dev.go`
- **Validation surface:** live-umbrella-same-state-normal-message-unsupported-content-type-noack
- **Front README candidate:** `false`

**What it does:** Prove unsupported normal application-message content_type does not ack or drain the inbox after same-state Relay join.

**Why it exists:** Second same-state failure-hardening companion profile after same-state-integrated-dev; locks the application-message ack rule for unsupported normal-message envelopes without adding adversarial harness claims or changing full/release-snapshot.

- **Required flags:** Not recorded in registry.
- **Optional flags:**
  - `--root`
  - `--clean-generated`
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - not included in full
  - not included in release-snapshot
  - not release-package validation
  - not package-root validation
  - not adversarial relay harness
  - not hostile-server safety
  - not metadata privacy
  - not production secure messaging
  - not production E2EE
  - not identity verification
  - not secure enrollment
  - not local-backbone
  - not deployment
  - not mature messenger UX
  - not general-public UX
  - currently covers unsupported normal application-message content_type only
  - not vault/key-storage safety
  - not PQ or hybrid security

### `runner.same-state-welcome-join-failure-dev`

- **Command:** `go run . --profile same-state-welcome-join-failure-dev`
- **Repo:** `carbonstack`
- **Component:** `tools/carbonstack-validate`
- **Kind:** `runner-profile`
- **Audience:** `dev`
- **Maturity:** `dev_only`
- **Introduced in:** `v0.6.21`
- **Source path:** `tools/carbonstack-validate/same_state_welcome_join_failure_dev.go`
- **Validation surface:** live-umbrella-same-state-corrupt-welcome-join-noack-nodrain-no-state-poison
- **Front README candidate:** `false`

**What it does:** Prove corrupt Welcome join fails without ack, Relay inbox drain, or final/staging sidecar state poison, then restored Welcome joins and acks.

**Why it exists:** Locks the v0.6.20 atomic Welcome join state-write fix into a live-umbrella failure-path profile before stale-state modeling and later full-runtime aggregation.

- **Required flags:** Not recorded in registry.
- **Optional flags:**
  - `--root`
  - `--clean-generated`
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - not included in full
  - not included in release-snapshot
  - not release-package validation
  - not package-root validation
  - not adversarial relay harness
  - not hostile-server safety
  - not metadata privacy
  - not production secure messaging
  - not production E2EE
  - not identity verification
  - not secure enrollment
  - not local-backbone
  - not deployment
  - not mature messenger UX
  - not general-public UX
  - currently covers corrupt Welcome join failure with restored Welcome recovery only
  - not stale provider state modeling
  - not vault/key-storage safety
  - not PQ or hybrid security

### `runner.state-path-policy-dev`

- **Command:** `go run . --profile state-path-policy-dev`
- **Repo:** `carbonstack`
- **Component:** `tools/carbonstack-validate`
- **Kind:** `runner-profile`
- **Audience:** `dev`
- **Maturity:** `dev_only`
- **Lifecycle status:** `active`
- **Introduced in:** `Not recorded in registry.`
- **Source path:** `carbonstack/tools/carbonstack-validate/state_path_policy_dev.go`
- **Validation surface:** state-path-policy-dev-profile
- **Front README candidate:** `false`

**What it does:** Validate Gate C3 path policy, explicit state-root classification, unsafe path refusal, and registry/reference freshness.

**Why it exists:** Closes Gate C3 as the path policy and explicit state-root semantics subgate before C4 atomicity work.

- **Required flags:** Not recorded in registry.
- **Optional flags:**
  - `--compact-summary` — print compact evidence where supported.
  - `--clean-generated` — remove known generated/build artifacts after successful profile execution.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - not migration
  - not silent repair
  - not state relocation
  - not cleanup implementation
  - not C4 atomicity or lock closure
  - not C5 Gate C closure
  - not vault security
  - not backup restore
  - not trust promotion
  - not verified identity
  - not deployment
  - not full-runtime-dev
  - not Gate D runtime aggregate
  - not production E2EE

### `runner.state-schema-compat-dev`

- **Command:** `go run . --profile state-schema-compat-dev`
- **Repo:** `carbonstack`
- **Component:** `tools/carbonstack-validate`
- **Kind:** `runner-profile`
- **Audience:** `dev`
- **Maturity:** `dev_only`
- **Lifecycle status:** `active`
- **Introduced in:** `Not recorded in registry.`
- **Source path:** `carbonstack/tools/carbonstack-validate/state_schema_compat_dev.go`
- **Validation surface:** state-schema-compat-dev-profile
- **Front README candidate:** `false`

**What it does:** Validate Gate C2 schema/version compatibility classification and refusal for Comms-owned JSON artifacts.

**Why it exists:** Closes Gate C2 as an explicit schema compatibility/refusal subgate before C3 path policy and C4 atomicity work.

- **Required flags:** Not recorded in registry.
- **Optional flags:**
  - `--compact-summary` — print compact evidence where supported.
  - `--clean-generated` — remove known generated/build artifacts after successful profile execution.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - not migration
  - not silent repair
  - not state relocation
  - not vault security
  - not backup restore
  - not trust promotion
  - not verified identity
  - not deployment
  - not full-runtime-dev
  - not Gate D runtime aggregate
  - not production E2EE

### `runner.state-substrate-inventory-dev`

- **Command:** `go run . --profile state-substrate-inventory-dev`
- **Repo:** `carbonstack`
- **Component:** `tools/carbonstack-validate`
- **Kind:** `runner-profile`
- **Audience:** `dev`
- **Maturity:** `dev_only`
- **Lifecycle status:** `active`
- **Introduced in:** `Not recorded in registry.`
- **Source path:** `carbonstack/tools/carbonstack-validate/state_substrate_inventory_dev.go`
- **Validation surface:** state-substrate-inventory-dev-profile
- **Front README candidate:** `false`

**What it does:** Validate Gate C1 state substrate inventory, generated report output, canonical root policy, and Gate B regression compatibility.

**Why it exists:** Closes Gate C1 as an inventory/authority-map subgate before C2 schema/version enforcement.

- **Required flags:** Not recorded in registry.
- **Optional flags:**
  - `--compact-summary` — print compact evidence where supported.
  - `--clean-generated` — remove known generated/build artifacts after successful profile execution.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - not schema enforcement
  - not migration
  - not silent repair
  - not vault security
  - not backup restore
  - not trust promotion
  - not verified identity
  - not deployment
  - not full-runtime-dev
  - not Gate D runtime aggregate
  - not production E2EE

### `runner.state-write-policy-dev`

- **Command:** `go run . --profile state-write-policy-dev`
- **Repo:** `carbonstack`
- **Component:** `tools/carbonstack-validate`
- **Kind:** `runner-profile`
- **Audience:** `dev`
- **Maturity:** `dev_only`
- **Lifecycle status:** `active`
- **Introduced in:** `Not recorded in registry.`
- **Source path:** `carbonstack/tools/carbonstack-validate/state_write_policy_dev.go`
- **Validation surface:** state-write-policy-dev-profile
- **Front README candidate:** `false`

**What it does:** Validate Gate C4 write-policy classification, report generation, C3/C2/C1 smoke, and registry/reference freshness.

**Why it exists:** Closes Gate C4 as the write policy and partial-state classification subgate before C5 Gate C closure.

- **Required flags:** Not recorded in registry.
- **Optional flags:**
  - `--compact-summary` — print compact evidence where supported.
  - `--clean-generated` — remove known generated/build artifacts after successful profile execution.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - not migration
  - not silent repair
  - not state relocation
  - not runtime writer rewiring
  - not cleanup implementation
  - not destructive cleanup
  - not C5 Gate C closure
  - not vault security
  - not backup restore
  - not trust promotion
  - not verified identity
  - not deployment
  - not full-runtime-dev
  - not Gate D runtime aggregate
  - not production E2EE

### `runner.welcome-lifecycle-dev`

- **Command:** `go run . --profile welcome-lifecycle-dev`
- **Repo:** `carbonstack`
- **Component:** `tools/carbonstack-validate`
- **Kind:** `runner-profile`
- **Audience:** `dev`
- **Maturity:** `dev_only`
- **Lifecycle status:** `active`
- **Introduced in:** `Not recorded in registry.`
- **Source path:** `carbonstack/tools/carbonstack-validate/welcome_lifecycle_dev.go`
- **Validation surface:** welcome-lifecycle-dev-profile
- **Front README candidate:** `false`

**What it does:** Validate Gate B6 Welcome persistence, join, ACK-after-join, and failure non-ACK behavior.

**Why it exists:** Proves the bounded B6 Welcome lifecycle leaf and supports v0.7.9 B6 closure.

- **Required flags:** Not recorded in registry.
- **Optional flags:**
  - `--compact-summary` — print compact evidence where supported.
  - `--clean-generated` — remove known generated/build artifacts after successful profile execution.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - not verified identity
  - not trust promotion
  - not production E2EE
  - not B7 Cypher/MLS reconciliation
  - not full operational aggregate

### `runner.workflow-relay-onboarding-dev`

- **Command:** `go run . --profile workflow-relay-onboarding-dev`
- **Repo:** `carbonstack`
- **Component:** `tools/carbonstack-validate`
- **Kind:** `runner-profile`
- **Audience:** `dev`
- **Maturity:** `dev_only`
- **Lifecycle status:** `active`
- **Introduced in:** `Not recorded in registry.`
- **Source path:** `carbonstack/tools/carbonstack-validate/workflow_relay_onboarding_dev.go`
- **Validation surface:** workflow-relay-onboarding-dev-profile
- **Front README candidate:** `false`

**What it does:** Validate Gate B8 workflow readiness, partial-state refusal, durable report, replay, and boundary preservation.

**Why it exists:** Proves B8 reusable workflow report/evaluator while preserving B4/B5/B6/B7 lifecycle semantics and nonclaims.

- **Required flags:** Not recorded in registry.
- **Optional flags:**
  - `--compact-summary` — print compact evidence where supported.
  - `--clean-generated` — remove known generated/build artifacts after successful profile execution.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - not silent repair
  - not silent rejoin
  - not trust promotion
  - not verified identity
  - not B9 Gate B closure
  - not full-runtime-dev
  - not production E2EE
