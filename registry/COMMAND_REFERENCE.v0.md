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

Registry entry count: **79**

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
- **Validation surface:** wrapper-based OpenMLS runtime smoke
- **Front README candidate:** `false`

**What it does:** Validate the bootstrap-wrapper plus message-wrapper dev OpenMLS runtime path through Comms and Cypher.

**Why it exists:** Separate maturity profile for openmls-*-dev bootstrap wrappers plus message-send-dev -> Cypher -> message-inbox-dev --ack before any future merge with direct smoke.

- **Example:** `cd carbonstack/tools/carbonstack-validate && go run . --profile dev-runtime-openmls-wrappers --clean-generated`
- **Required flags:** Not recorded in registry.
- **Optional flags:**
  - `--root` — optional live umbrella root override. Boundary: live-dev profile; not release-package validation.
  - `--clean-generated` — clean known generated roots after successful validation. Boundary: only known OpenMLS sidecar generated roots are cleaned.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - not local-backbone
  - does not replace dev-runtime-openmls yet
  - not release-package validation
  - not included in full
  - not mature messaging UX
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
- **Validation surface:** live-umbrella-integrated-runtime-dev-composition
- **Front README candidate:** `false`

**What it does:** Run Relay onboarding proof, then wrapper normal-message runtime proof, as a live-umbrella integrated dev profile.

**Why it exists:** Provides the first bounded in-series integrated runtime validation ladder without mutating full or release-snapshot.

- **Example:** `cd carbonstack/tools/carbonstack-validate && go run . --profile integrated-runtime-dev --root ~/repos/carbonstack_umbrella --clean-generated`
- **Required flags:** Not recorded in registry.
- **Optional flags:**
  - `--root` — optional live umbrella root override. Boundary: live-dev profile; not release-package validation.
  - `--clean-generated` — clean known generated roots after successful validation. Boundary: does not make the profile package-root validation.
- **Environment:** Not recorded in registry.
- **Related registry rows:**
  - runner.relay-openmls-join-dev
  - runner.dev-runtime-openmls-wrappers
  - comms.openmls-relay-join-dev
  - comms.message-send-dev
  - comms.message-inbox-dev
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
  - does not replace relay-openmls-join-dev
  - does not replace dev-runtime-openmls-wrappers
  - first implementation composes existing profile proofs in sequence; does not claim same-state package-root release proof

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
- **Validation surface:** dev-runtime-openmls-wrappers and Comms package tests
- **Front README candidate:** `true`

**What it does:** Opinionated dev/pre-alpha normal-message inbox/open wrapper over the OpenMLS application-message path for public testing.

**Why it exists:** Current opinionated dev/pre-alpha message inbox/open wrapper over the OpenMLS application-message receive/open path for the OpenMLS/Cypher proof surface; not mature product UX.

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
- **Validation surface:** dev-runtime-openmls-wrappers and Comms package tests
- **Front README candidate:** `true`

**What it does:** Opinionated dev/pre-alpha normal-message send wrapper over the OpenMLS application-message path for public testing.

**Why it exists:** Current opinionated dev/pre-alpha message send wrapper over the OpenMLS application-message send path for the OpenMLS/Cypher proof surface; not mature product UX.

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
  - not secure enrollment
  - does not perform Relay onboarding

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

Entries in this section: **6**

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

Entries in this section: **7**

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

## Comms state, account, device, and trust commands

Entries in this section: **9**

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

Entries in this section: **14**

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
- **Validation surface:** dev-runtime-openmls-wrappers runner profile and wrapper OpenMLS runtime smoke
- **Front README candidate:** `false`

**What it does:** Wrapper-bootstrap smoke proving openmls-*-dev wrappers -> message-send-dev -> Cypher -> message-inbox-dev --ack.

**Why it exists:** Higher-level wrapper maturity proof while direct-sidecar baseline remains.

- **Required flags:** Not recorded in registry.
- **Optional flags:** Not recorded in registry.
- **Environment:** Not recorded in registry.
- **Related registry rows:** Not recorded in registry.
- **Not claims:**
  - not local-backbone
  - does not replace direct smoke yet

## Internal OpenMLS sidecar provider commands

Entries in this section: **10**

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

Entries in this section: **9**

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

Entries in this section: **3**

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
