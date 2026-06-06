[META NOTE: REDACTIONS REPRESENTED AS ▓▓ CHARACTERS. Each ▓▓ marker replaces a sensitive local user/path or local host/committer identifier, regardless of the original string length.]

# CarbonStack LogDoc v0.4.23PRIME

**Last updated:** 2026-06-05 local session  
**Owner/operator:** bitcrusher32  
**Assistant/model:** GPT-5.5 Thinking  
**Current phase:** Phase 2F / **v0.5.0 minor epoch pre-release cut and v0.4.x PRIME handoff complete**. The v0.5.0 Gitea pre-release is now published at `https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.5.0`. v0.4.22 proved checksum/archive/fresh-extraction package rehearsal; v0.4.23PRIME completed the release-facing work: v0.5.0 release notes in v0.4.0 continuity style, final release package asset preparation, asset checksum generation, Gitea release publication, and preservation of the final v0.4.x handoff before compression into LogDoc v0.5.0. `carbonstack` remains at `c6aa4e3 (HEAD -> main, origin/main, origin/HEAD) test: rehearse v0.5.0 package validation`; `carbonstack-comms` remains at `cb4e59d (HEAD -> main, origin/main, origin/HEAD) test: add wrapper-based OpenMLS runtime smoke proof`; `carbonstack-cypher` remains at `9ab994c (HEAD -> main, origin/main, origin/HEAD) docs: point to local operator runbook`; `carbonstack-os` remains at `1bbbe52 (HEAD -> main, origin/main, origin/HEAD) docs: clarify CarbonStackOS target direction`.

**Version schema:** `v[scope].[timeline]`. This file is `v0.4.23PRIME`, the final PRIME ledger for the v0.4.x minor epoch after the v0.5.0 release cut. PRIME denotes the final LogDoc in a minor epoch or pre-compression handoff. Per LogDoc V2 practice, this Markdown file preserves current operational continuity and carries the final v0.4.x release-prep thread into the next compressed LogDoc v0.5.0 state; the JSON breakpoint remains the lean machine-readable handoff state.

**Compression note:** The detailed v0.3.0 through v0.3.36PRIME process ledger remains archived in PRIME/sanitized PRIME. v0.4.0 compressed that pre-release sprint. v0.4.1 through v0.4.12 preserved runtime recon, command contracts, dev OpenMLS send/inbox commands, smoke proof, runner profile, boundary checks, wrapper recon, and bootstrap command contract. v0.4.13 through v0.4.15 preserved staged bootstrap wrapper implementation. v0.4.16 preserved the first committed wrapper-based runtime smoke variant. v0.4.17 preserved the separate manual runner profile for that wrapper smoke. v0.4.18 preserved the first provisional cross-repo command registry. v0.4.19 preserved registry validation/test coverage and local-help/manual planning. v0.4.20 preserved roadmap refresh and release-surface current-state cleanup. v0.4.21 preserved v0.5.0 package rehearsal staging. v0.4.22 preserved v0.5.0 checksum/archive/fresh-extraction package validation rehearsal. v0.4.23PRIME now preserves the final v0.5.0 release cut and closes the v0.4.x release-prep line before compression into LogDoc v0.5.0.

---

## 1. Project Goal

**Active goal:** Close the v0.4.x release-prep line and transition into the v0.5.0 minor epoch after a successful Gitea-source-of-truth pre-release cut.

CarbonStack’s broader project intent remains:

- **CarbonStack** is the secure-communications appliance-stack project and shared doctrine/release authority.
- **CarbonStackComms** is the text-first messaging client and current OpenMLS sidecar/relay integration surface.
- **CarbonStackCypher** is the self-hostable relay/storage stack.
- **CarbonStackOS** is future constrained Android-derived appliance OS work and remains out of current runnable validation packages.
- The `carbonstack` Gitea repo remains the source of truth for release state, release assets, docs, registry, and public claims.
- GitHub mirrors exist as secondary push mirrors only; Gitea remains authoritative for current pre-release work.

**North star:**

> Its goal is not to be a better smartphone. Its goal is to stop being a smartphone.

**Core doctrine:**

> Every feature is guilty until it proves it does not add unacceptable parser, network, sensor, identity, or filesystem authority.

**Current reality after v0.4.23PRIME:**

    Public release is now:
      v0.5.0 Runtime and Registry Validation Minor Epoch Pre-Release
      https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.5.0

    Gitea release status:
      Pre-Release

    Release target:
      v0.5.0 at carbonstack commit c6aa4e3 / c6aa4e3183.

    Public release title:
      CarbonStack v0.5.0 Runtime and Registry Validation Minor Epoch Pre-Release

    Release body:
      v0.4.0-style Markdown continuity.
      Explicitly describes v0.5.0 as a minor epoch pre-release.
      Frames the release as accumulated v0.4.x runtime, runner, registry, and package-validation work.
      Explicitly states GitHub mirrors remain secondary push mirrors.
      Explicitly states this is not general-public usable software, not v1.0.0, not local-backbone, and not production secure.

    Release package/assets were prepared from the v0.4.22 rehearsal flow:
      stage clean package.
      write checksums.
      verify checksums.
      archive package.
      fresh-extract package.
      verify checksums from fresh extraction.
      run full --clean-generated from fresh extraction.
      copy final-named assets into a Windows-accessible upload folder.
      generate asset checksums.
      upload release assets to Gitea.

    Uploaded / public release assets:
      carbonstack-v0.5.0-minor-epoch-pre-release.tgz
      carbonstack-v0.5.0-release-manifest.json
      carbonstack-v0.5.0-package-checksums.txt
      carbonstack-v0.5.0-asset-checksums.txt
      carbonstack-v0.5.0-validation-freeze.md
      v0.5.0-testing-runbook.md
      v0.5.0-release-notes.md
      LICENSE

    Gitea also auto-generates:
      Source Code (ZIP)
      Source Code (TAR.GZ)

    Important release boundary:
      The Gitea Source Code ZIP/TAR.GZ are carbonstack-only autogenerated snapshots, not the intended multi-repo validation package.
      The intended runnable validation artifact is the attached v0.5.0 package plus manifest, checksums, validation freeze, and testing runbook.

Hard nonclaims remain: v0.5.0 does **not** prove production readiness, production E2EE, hostile-server safety, metadata privacy, Android readiness, CarbonStackOS readiness, secure vault/storage, public ingress safety, real homelab validation, audit, certification, PQ/hybrid security, mature messenger UX, release-package runtime OpenMLS validation, Relay Space join/onboarding, `full` runtime OpenMLS validation, or local-backbone readiness. This is a minor epoch pre-release of accumulated v0.4.x runtime/runner/registry/package-validation work, not v1.0.0 and not a production secure messenger.

---

## 2. Current Public Release State

Current public release:

    https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.5.0

Release title:

    CarbonStack v0.5.0 Runtime and Registry Validation Minor Epoch Pre-Release

Release tag:

    v0.5.0

Release commit:

    c6aa4e3183 / c6aa4e3 test: rehearse v0.5.0 package validation

Gitea release status:

    Pre-Release

Release published:

    2026-06-05 local session

Release page confirms:

    v0.5.0 released from commit c6aa4e3183.
    0 commits to main since this release at publication.
    Release body describes the same v0.5.0 minor epoch pre-release framing.
    Release assets are attached with v0.5.0 final filenames.

Important release-page warning:

    The default Gitea “Source Code” ZIP/TAR.GZ downloads are only auto-generated archives of the carbonstack repo at the v0.5.0 tag.
    Use the attached v0.5.0 package, manifest, checksums, validation freeze, and testing runbook for the intended multi-repo validation package.

Previous public release:

    v0.4.0 Broad Local Deployability Pre-Release
    https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.4.0

Post-release public surface:

    carbonstack README is evergreen.
    Releases page is the known-good artifact surface.
    Gitea remains source of truth.
    GitHub mirrors remain secondary push mirrors, not release authority.

Official general-public usable releases:

    Not yet.
    User explicitly intends official usable/general-public releases to happen on GitHub later, such as the v1.0.0 major epoch release.
    v0.5.0 is a Gitea-source-of-truth pre-release/minor epoch release, not a general-public official usable release.

---

## 3. Current Repo Heads

    carbonstack        c6aa4e3 (HEAD -> main, origin/main, origin/HEAD) test: rehearse v0.5.0 package validation
    carbonstack-comms  cb4e59d (HEAD -> main, origin/main, origin/HEAD) test: add wrapper-based OpenMLS runtime smoke proof
    carbonstack-cypher 9ab994c (HEAD -> main, origin/main, origin/HEAD) docs: point to local operator runbook
    carbonstack-os     1bbbe52 (HEAD -> main, origin/main, origin/HEAD) docs: clarify CarbonStackOS target direction

Current public release tag:

    v0.5.0 -> c6aa4e3183 / c6aa4e3 test: rehearse v0.5.0 package validation

Previous public release tag:

    v0.4.0 -> 24ac9fc494 / 24ac9fc docs: record v0.4.0 package rehearsal

Recent maturity / release-prep commits:

    v0.4.17 wrapper-smoke runner/profile alignment:
      a501442 feat: add wrapper OpenMLS runtime validation profile

    v0.4.18 command registry:
      a4162d7 docs: add known-good command registry

    v0.4.19 registry validation:
      568fb45 test: validate command registry coverage

    v0.4.20 release-surface cleanup:
      7d4fc0d docs: refresh release-surface current state

    v0.4.21 package staging:
      2c2554d chore: stage v0.5.0 package rehearsal

    v0.4.22 package rehearsal validation / v0.5.0 release commit:
      c6aa4e3 test: rehearse v0.5.0 package validation

Continuity note:

    `c6aa4e3` is both the final v0.4.22 release-prep commit and the v0.5.0 release target.
    v0.4.23PRIME is a LogDoc/release-state handoff, not a new code commit.

---

## 4. Current Validated / Known State

Validated / known at the v0.4.23PRIME checkpoint:

    [RELEASE] v0.5.0 Gitea pre-release is published.
    [RELEASE] v0.5.0 is attached to carbonstack commit c6aa4e3 / c6aa4e3183.
    [RELEASE] Release title is CarbonStack v0.5.0 Runtime and Registry Validation Minor Epoch Pre-Release.
    [RELEASE] Release body follows v0.4.0 continuity style while reflecting v0.5.0 actual evidence.
    [RELEASE] Release body explicitly warns that Gitea Source Code ZIP/TAR.GZ are autogenerated carbonstack-only archives.
    [RELEASE] Release assets include the final v0.5.0 package, manifest, package checksums, asset checksums, validation freeze, testing runbook, release notes, and LICENSE.
    [PACKAGE] v0.5.0 package rehearsal passed with staged checksums, archive, fresh extraction, verify-checksums, full --clean-generated, and final generated-root absence checks.
    [PACKAGE] carbonstack-os stayed excluded from the runnable package.
    [DOCS] v0.5.0 release notes were prepared in v0.4.0 release style.
    [HELPERS] v0.5.0 release-helper system remains:
      scripts/stage-v0.5.0-package.sh
      scripts/rehearse-v0.5.0-package.sh
    [REGISTRY] command registry includes both release helpers.
    [RUNNER] full remains release-snapshot followed by local-cypher.
    [RUNNER] dev-runtime-openmls and dev-runtime-openmls-wrappers remain live-umbrella-only and are not included in full.
    [MIRRORS] Gitea remains authoritative; GitHub repos are secondary push mirrors.

v0.4.23PRIME release/publication observed:

    v0.5.0 release page exists at:
      https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.5.0

    Release page title:
      CarbonStack v0.5.0 Runtime and Registry Validation Minor Epoch Pre-Release

    Release page status:
      Pre-Release

    Release page commit:
      c6aa4e3183 / c6aa4e3.

    Release page assets:
      carbonstack-v0.5.0-asset-checksums.txt
      carbonstack-v0.5.0-minor-epoch-pre-release.tgz
      carbonstack-v0.5.0-package-checksums.txt
      carbonstack-v0.5.0-release-manifest.json
      carbonstack-v0.5.0-validation-freeze.md
      LICENSE
      v0.5.0-release-notes.md
      v0.5.0-testing-runbook.md

    Autogenerated Gitea downloads:
      Source Code (ZIP)
      Source Code (TAR.GZ)

Expected live umbrella validation commands remain:

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go test ./... -count=1
    go run . --profile dev-runtime-openmls-wrappers --clean-generated
    go run . --profile dev-runtime-openmls --clean-generated
    go run . --profile local-cypher
    go run . --profile doctor
    go run . --profile core --clean-generated

Expected public package validation commands:

    cd <package-root>/carbonstack/tools/carbonstack-validate
    go run . --profile verify-checksums --root <package-root>
    go run . --profile full --root <package-root> --clean-generated

Expected v0.5.0 release-helper commands:

    cd ~/repos/carbonstack_umbrella/carbonstack
    scripts/stage-v0.5.0-package.sh
    scripts/rehearse-v0.5.0-package.sh

---

## 5. v0.4.23PRIME Work Completed

### 5.1 v0.5.0 release notes were formulated using v0.4.0 continuity style

Inputs:

    v0.4.0 release screenshot.
    v0.4.0 raw Markdown body.
    v0.4.0 release link:
      https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.4.0
    v0.4.22 LogDoc actual evidence.

The release notes preserved the v0.4.0 structure:

    title
    status / primary artifact / release type / platform / secondary validation
    opening release explanation
    most important validation path
    What changed
    Validated
    Release assets
    Testing notes
    Boundary

v0.5.0-specific substance:

    minor epoch pre-release.
    accumulated v0.4.x runtime, runner, registry, and package-validation milestone.
    dev OpenMLS runtime commands.
    direct and wrapper-based runtime smoke proofs.
    separated runner profiles.
    provisional known-good command registry.
    registry validation.
    repeatable v0.5.0 package rehearsal helpers.
    release-helper system.
    command-surface hygiene.
    explicit GitHub mirror-secondary boundary.

### 5.2 Final v0.5.0 assets were prepared for upload

Asset staging flow:

    Re-ran final rehearsal from WSL.
    Used /tmp/carbonstack-v0.5.0-rehearsal as rehearsal root.
    Copied final-named assets into /tmp/carbonstack-v0.5.0-release-assets.
    Copied those final assets into a Windows-accessible Downloads folder for Gitea browser upload.
    Replaced v0.5.0-release-notes.md with final release notes.
    Regenerated asset checksums after editing release notes.
    Uploaded the final asset set to the Gitea v0.5.0 release page.

Final asset names:

    carbonstack-v0.5.0-minor-epoch-pre-release.tgz
    carbonstack-v0.5.0-release-manifest.json
    carbonstack-v0.5.0-package-checksums.txt
    carbonstack-v0.5.0-asset-checksums.txt
    carbonstack-v0.5.0-validation-freeze.md
    v0.5.0-testing-runbook.md
    v0.5.0-release-notes.md
    LICENSE

### 5.3 Gitea release was published

Release URL:

    https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.5.0

Release title:

    CarbonStack v0.5.0 Runtime and Registry Validation Minor Epoch Pre-Release

Tag:

    v0.5.0

Status:

    Pre-Release

Target commit:

    c6aa4e3183 / c6aa4e3

Release page observed:

    0 commits to main since this release at publication time.
    Attached assets matched the intended v0.5.0 release asset list.
    Gitea autogenerated Source Code ZIP/TAR.GZ also appeared, as expected.

### 5.4 v0.4.xPRIME handoff state was defined

v0.4.23PRIME is the final handoff for the v0.4.x minor epoch.

Purpose:

    Preserve final v0.4.x process continuity before compression into LogDoc v0.5.0.
    Mark v0.5.0 release as complete.
    Preserve release asset/publication details.
    Preserve blunders and process lessons from the release cut.
    Define the next safe post-release actions for v0.5.x.

---

## 6. Critical Blunders / Lessons to Carry Forward

### 6.1 Gitea Source Code ZIP/TAR.GZ are automatic and not the intended package

Gitea automatically shows Source Code ZIP and Source Code TAR.GZ for the release tag.

Lesson:

    Always keep explicit release body wording warning that these are auto-generated carbonstack-only archives.
    The intended multi-repo validation artifact is the attached package `.tgz` plus manifest/checksums/runbook/freeze assets.
    Do not treat Gitea autogenerated source archives as the release package.

### 6.2 Release notes must track actual evidence, not version-number ambition

v0.5.0 sounds larger than v0.4.0, but the release remains pre-alpha/experimental.

Lesson:

    Keep the boundary explicit.
    Do not claim production readiness.
    Do not claim general-public usable software.
    Do not claim local-backbone.
    Do not claim PQ/quantum-safe messaging.
    Do not claim runtime OpenMLS profiles are part of `full`.

### 6.3 Release-helper root discipline remains critical

The release package should be staged clean, checksummed, archived, fresh-extracted, and validated from fresh extraction.

Lesson:

    Do not run `release-snapshot`/`full` inside the source package root intended for publication.
    Use fresh extraction for validation.
    Use `--clean-generated` where documented.
    Validate asset names and asset checksums after final release notes are edited.

### 6.4 Browser upload can drift from generated asset names if not checked

The release page must match the release body filenames exactly.

Lesson:

    Before publishing, compare the release body asset list to the final upload folder.
    Re-run asset checksums after editing release notes.
    Ensure final file names match:
      carbonstack-v0.5.0-minor-epoch-pre-release.tgz
      carbonstack-v0.5.0-release-manifest.json
      carbonstack-v0.5.0-package-checksums.txt
      carbonstack-v0.5.0-asset-checksums.txt
      carbonstack-v0.5.0-validation-freeze.md
      v0.5.0-testing-runbook.md
      v0.5.0-release-notes.md
      LICENSE

### 6.5 Prior copy/paste fragility remains relevant

The release-prep line had multiple paste/copy corruption incidents before v0.4.22 stabilized the helper flow.

Lesson:

    Continue using smaller copy blocks.
    Use `bash -n` before committing shell scripts.
    Use `git diff` and file sanity checks after any generated docs/scripts.
    For release work, treat final committed source + post-commit validation as source of truth.

### 6.6 Commit metadata warning remains unresolved

Earlier commits still showed local committer metadata:

    Committer: bitcrusher32 <▓▓>

Lesson:

    Not a v0.5.0 blocker.
    Eventually configure Git author/committer metadata cleanly for professionalism/publishing hygiene.

---

## 7. Current Blockers / Not Validated

Current blockers and nonvalidated areas after v0.4.23PRIME:

    v0.5.0 is released, but v0.4.xPRIME has not yet been compressed into LogDoc v0.5.0.
    Existing send/inbox remain stub-era.
    Runtime Comms has dev-only OpenMLS commands, direct smoke proof, wrapper-based smoke proof, and separate manual runner profiles, but no mature user-facing UX.
    Neither runtime profile is included in full.
    Neither runtime profile is part of public release package validation.
    No Relay Space join model exists yet.
    local-backbone runner profile does not exist and remains reserved.
    Generated command reference from the registry does not exist yet.
    Local help/manual generation from the registry does not exist yet.
    v0.5.x provider/trust/vault and PQ/hybrid migration readiness remain future work.
    Hostile-server harnesses do not exist yet.
    Actual deployability/ops hardening is not done yet.

Hard nonclaims remain:

    Do not claim production readiness.
    Do not claim production E2EE.
    Do not claim hostile-server safety.
    Do not claim metadata privacy.
    Do not claim Debian production deployability.
    Do not claim systemd readiness.
    Do not claim cloudflared readiness.
    Do not claim Android readiness.
    Do not claim CarbonStackOS readiness.
    Do not claim secure vault/storage.
    Do not claim real homelab validation.
    Do not claim external audit or certification.
    Do not claim CarbonStack Relay Space mechanics are implemented.
    Do not claim local-backbone exists or is justified yet.
    Do not claim local-cypher validates runtime Comms UX.
    Do not claim full validates local-backbone.
    Do not claim full is a deployment command.
    Do not claim dev-runtime-openmls is release-package validation.
    Do not claim dev-runtime-openmls-wrappers is release-package validation.
    Do not claim send/inbox are OpenMLS-backed.
    Do not claim PQ/hybrid security or quantum-safe messaging.
    Do not let GitHub mirrors supersede Gitea source-of-truth status.
    Do not treat v0.5.0 as general-public usable software.
    Do not treat v0.5.0 as v1.0.0.

Allowed claim:

    CarbonStack current public release is now v0.5.0 Runtime and Registry Validation Minor Epoch Pre-Release.
    v0.5.0 is a Gitea-source-of-truth pre-release.
    v0.5.0 packages accumulated v0.4.x runtime, runner, registry, and package-validation work.
    v0.5.0 release package assets are attached to the Gitea release page.
    The intended runnable validation artifact is the attached multi-repo package, not Gitea autogenerated source archives.
    The package root includes carbonstack, carbonstack-comms, carbonstack-cypher, and release metadata.
    carbonstack-os is explicitly excluded from the runnable package.
    The v0.5.0 rehearsal flow writes release/checksums.txt, verifies checksums in the staged package, archives package, fresh-extracts, verifies checksums again, and runs full from fresh extraction.
    Final generated/private/build roots are absent after --clean-generated in the rehearsal flow.
    Gitea remains source of truth; GitHub mirrors are secondary push mirrors.

---

## 8. Next Safest Actions

Recommended next checkpoint:

    v0.5.0 LogDoc compression / post-release baseline.

Focus:

    Compress v0.4.xPRIME into LogDoc v0.5.0.
    Treat v0.5.0 as the new current public release.
    Preserve v0.4.23PRIME as the final detailed v0.4.x ledger.
    Decide the first v0.5.x implementation/planning rung.
    Likely begin with post-release preflight rather than immediate code.
    Confirm whether roadmap should be refreshed after the v0.5.0 release cut.
    Re-evaluate v0.5.x state/trust/vault/PQ planning order before implementation.

Recommended v0.5.x first rungs:

    v0.5.1 post-release preflight and baseline.
    v0.5.2 state/trust/vault/PQ planning recon.
    v0.5.3 storage/trust/provider-state inventory.
    Later: PQ/hybrid ciphersuite migration planning only after state/trust/provider boundaries are clear.

Avoid next:

    starting PQ implementation immediately without state/trust/provider inventory.
    treating v0.5.0 as production secure.
    including runtime profiles in full without deliberate design/release decision.
    adding local-backbone naming before the system earns it.
    treating GitHub mirrors as source of truth.
    rewriting historical numbered docs merely because they are stale.

---

## 9. Critical Paths / Functions

### Source-of-truth public surfaces

    [GITEA SOURCE OF TRUTH] https://git.bitcrusher32.win/bitcrusher32/carbonstack
    [RELEASES] https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases
    [CURRENT PUBLIC RELEASE] https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.5.0
    [PREVIOUS PUBLIC RELEASE] https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.4.0

GitHub mirrors:

    secondary push mirrors for discoverability/redundancy only.
    Do not treat mirrors as release authority unless project policy changes.
    User intends official general-public releases later on GitHub, such as v1.0.0, but v0.5.0 is not that.

### v0.5.0 public release assets

    [ASSET] carbonstack-v0.5.0-minor-epoch-pre-release.tgz
    [ASSET] carbonstack-v0.5.0-release-manifest.json
    [ASSET] carbonstack-v0.5.0-package-checksums.txt
    [ASSET] carbonstack-v0.5.0-asset-checksums.txt
    [ASSET] carbonstack-v0.5.0-validation-freeze.md
    [ASSET] v0.5.0-testing-runbook.md
    [ASSET] v0.5.0-release-notes.md
    [ASSET] LICENSE

### WSL Debian active working umbrella

    [WSL-DEBIAN] umbrella: ~/repos/carbonstack_umbrella
    [WSL-DEBIAN] carbonstack: ~/repos/carbonstack_umbrella/carbonstack
    [WSL-DEBIAN] carbonstack-comms: ~/repos/carbonstack_umbrella/carbonstack-comms
    [WSL-DEBIAN] carbonstack-cypher: ~/repos/carbonstack_umbrella/carbonstack-cypher
    [WSL-DEBIAN] carbonstack-os: ~/repos/carbonstack_umbrella/carbonstack-os
    [GO-RUNNER] runner dir: ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate

### v0.5.0 package staging/rehearsal surfaces

    [STAGING SCRIPT] carbonstack/scripts/stage-v0.5.0-package.sh
    [REHEARSAL SCRIPT] carbonstack/scripts/rehearse-v0.5.0-package.sh
    [SCRIPTS README] carbonstack/scripts/README.md
    [STAGING DOC] carbonstack/docs/166-v0.5.0-package-rehearsal-plan-v0.md
    [REHEARSAL DOC] carbonstack/docs/167-v0.5.0-package-checksum-and-fresh-extraction-rehearsal-v0.md
    [REGISTRY ENTRY] carbonstack.script.stage-v0.5.0-package in carbonstack/registry/commands.v0.yaml
    [REGISTRY ENTRY] carbonstack.script.rehearse-v0.5.0-package in carbonstack/registry/commands.v0.yaml
    [DEFAULT REHEARSAL ROOT] /tmp/carbonstack-v0.5.0-rehearsal
    [DEFAULT STAGE ROOT] /tmp/carbonstack-v0.5.0-rehearsal/stage
    [DEFAULT PACKAGE ROOT] /tmp/carbonstack-v0.5.0-rehearsal/stage/package
    [DEFAULT ARCHIVE] /tmp/carbonstack-v0.5.0-rehearsal/archive/carbonstack-v0.5.0-package-rehearsal.tar.gz
    [DEFAULT FRESH EXTRACTION ROOT] /tmp/carbonstack-v0.5.0-rehearsal/extract/package

### Current live public docs surfaces

    [DOC] top README:
    carbonstack/README.md

    [DOC] docs index:
    carbonstack/docs/README.md

    [DOC] roadmap:
    carbonstack/roadmap/ROADMAP.md

    [RUNNER] runner README:
    carbonstack/tools/carbonstack-validate/README.md

    [REGISTRY] registry README:
    carbonstack/registry/README.md

### Registry surfaces

    [REGISTRY] root:
    carbonstack/registry

    [REGISTRY] provisional command registry:
    carbonstack/registry/commands.v0.yaml

    [REGISTRY] registry README:
    carbonstack/registry/README.md

    [REGISTRY-TEST] command registry validation:
    carbonstack/tools/carbonstack-validate/command_registry_test.go

### Runner surfaces

    [RUNNER] root:
    carbonstack/tools/carbonstack-validate

    [RUNNER] dispatch:
    carbonstack/tools/carbonstack-validate/main.go

    [RUNNER] local-cypher:
    carbonstack/tools/carbonstack-validate/local_cypher.go

    [RUNNER] release-snapshot:
    carbonstack/tools/carbonstack-validate/release_snapshot.go

    [RUNNER] checksums:
    carbonstack/tools/carbonstack-validate/checksums.go

### Comms runtime surfaces

    [COMMS] direct-sidecar dev runtime smoke proof:
    carbonstack-comms/scripts/dev-openmls-runtime-smoke.sh

    [COMMS] wrapper-based dev runtime smoke proof:
    carbonstack-comms/scripts/dev-openmls-runtime-smoke-wrappers.sh

### Known-good commands

    [LIVE VALIDATION]
    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go test ./... -count=1
    go run . --profile dev-runtime-openmls-wrappers --clean-generated
    go run . --profile dev-runtime-openmls --clean-generated
    go run . --profile local-cypher
    go run . --profile doctor
    go run . --profile core --clean-generated

    [PUBLIC PACKAGE VALIDATION]
    cd <package-root>/carbonstack/tools/carbonstack-validate
    go run . --profile verify-checksums --root <package-root>
    go run . --profile full --root <package-root> --clean-generated

    [PACKAGE STAGING]
    cd ~/repos/carbonstack_umbrella/carbonstack
    scripts/stage-v0.5.0-package.sh

    [PACKAGE REHEARSAL]
    cd ~/repos/carbonstack_umbrella/carbonstack
    scripts/rehearse-v0.5.0-package.sh

    [STAGED CHECKSUM SHAPE]
    cd /tmp/carbonstack-v0.5.0-rehearsal/stage/package/carbonstack/tools/carbonstack-validate
    go run . --profile write-checksums --root /tmp/carbonstack-v0.5.0-rehearsal/stage/package
    go run . --profile verify-checksums --root /tmp/carbonstack-v0.5.0-rehearsal/stage/package

    [FRESH EXTRACTION VALIDATION SHAPE]
    cd /tmp/carbonstack-v0.5.0-rehearsal/extract/package/carbonstack/tools/carbonstack-validate
    go run . --profile verify-checksums --root /tmp/carbonstack-v0.5.0-rehearsal/extract/package
    go run . --profile full --root /tmp/carbonstack-v0.5.0-rehearsal/extract/package --clean-generated

---

## 10. v0.5.x Forward Plan

### Immediate next: v0.5.0 LogDoc compression / post-release baseline

Expected:

    Compress v0.4.23PRIME into a lean LogDoc v0.5.0.
    Preserve v0.4.23PRIME as the final detailed v0.4.x release ledger.
    Treat v0.5.0 as the current public release.
    Run a post-release preflight.
    Decide whether to refresh the roadmap now that v0.5.0 is public.
    Choose the first v0.5.x implementation/planning rung.

### Recommended v0.5.1

Expected:

    post-release preflight and baseline.
    validate repo/release/public surfaces are sane after v0.5.0.
    check whether docs/README/roadmap need current-public-release updates.
    avoid jumping directly into PQ implementation.

### Recommended v0.5.2+

Expected:

    state/trust/vault/PQ planning recon.
    storage/trust/provider-state inventory.
    provider/vault boundary plan.
    PQ/hybrid ciphersuite migration planning only after state/trust/provider boundaries are clear.

---

## 11. v0.5.x and Later Carry-Forward

### v0.5.x implementation after release

Local storage, trust state, provider state, vault design, and PQ/hybrid ciphersuite migration readiness.

Do not start PQ/hybrid implementation before the runtime OpenMLS/Cypher path exists and is documented. PQ/hybrid migration touches provider state, ciphersuite metadata, trust display, artifact size, reinit/rekey/recovery policy, and claim boundaries.

### v0.6.x+

Hostile-server and abuse-resistance harnesses.

### v0.7.x+

Deployability and operations hardening.

### v0.8.x+

Documented self-pentest / adversarial validation.

### v0.9.x+

Claim-boundary review.

### v0.10.x+

Android backend/app work.

### v1.x.x

Public app/server major epoch and much later CarbonStackOS/appliance research, if justified.

---

## 12. Lean Breakpoint Summary

    CarbonStack v0.4.23PRIME is the final v0.4.x PRIME handoff after the v0.5.0 minor epoch pre-release cut. v0.5.0 is now the current public Gitea-source-of-truth pre-release at https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.5.0. carbonstack remains at c6aa4e3 test: rehearse v0.5.0 package validation, which is also the v0.5.0 release commit. carbonstack-comms remains at cb4e59d, carbonstack-cypher at 9ab994c, and carbonstack-os at 1bbbe52. The release attached the final v0.5.0 package, release manifest, package checksums, asset checksums, validation freeze, testing runbook, release notes, and LICENSE. Gitea autogenerated Source Code ZIP/TAR.GZ are present but explicitly not the intended multi-repo validation package. The v0.5.0 release preserves the v0.4.0 release style while updating substance for accumulated v0.4.x runtime/runner/registry/package-validation work. Next safe action: compress v0.4.23PRIME into LogDoc v0.5.0, then do a v0.5.1 post-release preflight before starting state/trust/vault/PQ planning.

---

## 13. Preserved Immediate Previous Handoff: v0.4.22

The following is the previous v0.4.22 handoff. Where it conflicts with the v0.4.23PRIME overlay above, v0.4.23PRIME wins for current state.



# CarbonStack LogDoc v0.4.22

**Last updated:** 2026-06-05 local session  
**Owner/operator:** bitcrusher32  
**Assistant/model:** GPT-5.5 Thinking  
**Current phase:** Phase 2F / **v0.5.0 checksum/archive/fresh-extraction package rehearsal checkpoint complete**. The v0.4.0 public broad local deployability pre-release remains the current public release. v0.4.21 added the v0.5.0 package staging helper; v0.4.22 proves the next release-package rehearsal step by staging a clean package, writing `release/checksums.txt`, verifying checksums in the staged package, creating a rehearsal archive, fresh-extracting it, verifying checksums again, running `full --clean-generated` from the fresh extraction, and confirming final generated/private/build roots are absent. `carbonstack` is now at `c6aa4e3 (HEAD -> main, origin/main, origin/HEAD) test: rehearse v0.5.0 package validation`; `carbonstack-comms` remains at `cb4e59d (HEAD -> main, origin/main, origin/HEAD) test: add wrapper-based OpenMLS runtime smoke proof`; `carbonstack-cypher` remains at `9ab994c (HEAD -> main, origin/main, origin/HEAD) docs: point to local operator runbook`; `carbonstack-os` remains at `1bbbe52 (HEAD -> main, origin/main, origin/HEAD) docs: clarify CarbonStackOS target direction`.

**Version schema:** `v[scope].[timeline]`. This file is `v0.4.22`, the v0.5.0 package checksum / archive / fresh-extraction rehearsal checkpoint after the v0.4.21 package staging checkpoint. Per LogDoc V2 practice, this Markdown file preserves current operational continuity and carries the compressed late-v0.4.x release-prep thread toward v0.4.23 final LogDoc/export/asset generation and the v0.5.0 minor epoch release; the JSON breakpoint remains the lean machine-readable handoff state.

**Compression note:** The detailed v0.3.0 through v0.3.36PRIME process ledger remains archived in PRIME/sanitized PRIME. v0.4.0 compressed that pre-release sprint; v0.4.1 through v0.4.12 preserved runtime recon, command contracts, dev OpenMLS send/inbox commands, smoke proof, runner profile, boundary checks, wrapper recon, and bootstrap command contract. v0.4.13 through v0.4.15 preserved staged bootstrap wrapper implementation. v0.4.16 preserved the first committed wrapper-based runtime smoke variant. v0.4.17 preserved the separate manual runner profile for that wrapper smoke. v0.4.18 preserved the first provisional cross-repo command registry. v0.4.19 preserved registry validation/test coverage and local-help/manual planning. v0.4.20 preserved roadmap refresh and release-surface current-state cleanup. v0.4.21 preserved v0.5.0 package rehearsal staging. v0.4.22 now preserves v0.5.0 checksum/archive/fresh-extraction package validation rehearsal.

---

## 1. Project Goal

**Active goal:** Finish the compressed late-v0.4.x release-prep runway toward a v0.5.0 minor epoch pre-release, while preserving v0.4.0 release-style continuity, package validation discipline, Gitea source-of-truth status, GitHub mirror-secondary status, and strict nonclaims.

CarbonStack’s broader project intent remains:

- **CarbonStack** is the secure-communications appliance-stack project and shared doctrine/release authority.
- **CarbonStackComms** is the text-first messaging client and current OpenMLS sidecar/relay integration surface.
- **CarbonStackCypher** is the self-hostable relay/storage stack.
- **CarbonStackOS** is future constrained Android-derived appliance OS work and remains out of current runnable validation packages.
- The `carbonstack` Gitea repo remains the source of truth for release state, release assets, docs, registry, and public claims.
- GitHub mirrors exist as secondary push mirrors only; Gitea remains authoritative for current pre-release work.

**North star:**

> Its goal is not to be a better smartphone. Its goal is to stop being a smartphone.

**Core doctrine:**

> Every feature is guilty until it proves it does not add unacceptable parser, network, sensor, identity, or filesystem authority.

**Current reality after v0.4.22:**

    carbonstack/scripts/stage-v0.5.0-package.sh exists and is executable.
    carbonstack/scripts/rehearse-v0.5.0-package.sh now exists and is executable.
    carbonstack/scripts/README.md now records the v0.5.0 release-helper split.
    carbonstack/docs/167-v0.5.0-package-checksum-and-fresh-extraction-rehearsal-v0.md records the v0.4.22 package validation rehearsal.
    carbonstack/registry/commands.v0.yaml includes:
      carbonstack.script.stage-v0.5.0-package
      carbonstack.script.rehearse-v0.5.0-package
    carbonstack/docs/README.md indexes docs/167.
    carbonstack/roadmap/ROADMAP.md records v0.4.22 package rehearsal completion and points to v0.4.23 final release-note/LogDoc/asset work.

    v0.5.0 release-helper system:
      scripts/stage-v0.5.0-package.sh stages a clean package skeleton and metadata.
      scripts/rehearse-v0.5.0-package.sh stages, writes checksums, verifies staged checksums, archives, fresh-extracts, verifies fresh checksums, and runs full from fresh extraction.

    Verified rehearsal roots:
      /tmp/carbonstack-v0.5.0-rehearsal/stage/package
      /tmp/carbonstack-v0.5.0-rehearsal/archive/carbonstack-v0.5.0-package-rehearsal.tar.gz
      /tmp/carbonstack-v0.5.0-rehearsal/extract/package

    Verified release files after fresh extraction:
      release/LICENSE
      release/manifest.json
      release/release-notes-draft.md
      release/testing-runbook.md
      release/validation-freeze.md
      release/checksums.txt

    Confirmed excluded:
      carbonstack-os

    Confirmed final generated/private/build roots absent after `full --clean-generated`:
      carbonstack-comms/internal/protocol/mls/openmls-sidecar/.carbonstack-openmls-sidecar-state
      carbonstack-comms/internal/protocol/mls/openmls-sidecar/target
      carbonstack-comms/internal/protocol/mls/openmls-sidecar/provider-storage.json
      carbonstack-comms/internal/protocol/mls/openmls-sidecar/signer.json

Hard nonclaims remain: v0.4.22 does **not** prove production readiness, production E2EE, hostile-server safety, metadata privacy, Android readiness, CarbonStackOS readiness, secure vault/storage, public ingress safety, real homelab validation, audit, certification, PQ/hybrid security, mature messenger UX, release-package runtime OpenMLS validation, Relay Space join/onboarding, `full` runtime OpenMLS validation, or local-backbone readiness. This is package rehearsal infrastructure and fresh-extraction validation, not final release cutting and not production validation.

---

## 2. Current Public Release State

Public release:

    https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.4.0

Release title:

    CarbonStack v0.4.0 Broad Local Deployability Pre-Release

Release tag:

    v0.4.0

Release commit:

    24ac9fc494 / 24ac9fc docs: record v0.4.0 package rehearsal

Gitea release status:

    Pre-Release

Important release-page warning remains:

    The default Gitea Source Code ZIP/TAR.GZ downloads are auto-generated carbonstack-only archives.
    The intended runnable validation artifact is the attached multi-repo release package plus companion assets.

Post-release public surface:

    carbonstack README is evergreen.
    It points interested parties to the Releases page for runnable / known-good artifacts.
    Gitea is source of truth.
    GitHub mirrors are secondary push mirrors, not release authority.

Official general-public usable releases:

    Not yet.
    User explicitly intends official usable/general-public releases to happen on GitHub later, such as the v1.0.0 major epoch release.
    Current v0.5.0 target is a Gitea-source-of-truth pre-release/minor epoch release, not a general-public official usable release.

---

## 3. Current Repo Heads

    carbonstack        c6aa4e3 (HEAD -> main, origin/main, origin/HEAD) test: rehearse v0.5.0 package validation
    carbonstack-comms  cb4e59d (HEAD -> main, origin/main, origin/HEAD) test: add wrapper-based OpenMLS runtime smoke proof
    carbonstack-cypher 9ab994c (HEAD -> main, origin/main, origin/HEAD) docs: point to local operator runbook
    carbonstack-os     1bbbe52 (HEAD -> main, origin/main, origin/HEAD) docs: clarify CarbonStackOS target direction

Current public release tag:

    v0.4.0 -> 24ac9fc494 / 24ac9fc docs: record v0.4.0 package rehearsal

Recent maturity / release-prep commits:

    v0.4.17 wrapper-smoke runner/profile alignment:
      a501442 feat: add wrapper OpenMLS runtime validation profile

    v0.4.18 command registry:
      a4162d7 docs: add known-good command registry

    v0.4.19 registry validation:
      568fb45 test: validate command registry coverage

    v0.4.20 release-surface cleanup:
      7d4fc0d docs: refresh release-surface current state

    v0.4.21 package staging:
      2c2554d chore: stage v0.5.0 package rehearsal

    Current v0.4.22 package rehearsal validation:
      c6aa4e3 test: rehearse v0.5.0 package validation

Continuity note:

    `c6aa4e3` is a carbonstack release-prep/rehearsal commit after the v0.4.0 public release tag.
    It does not retag v0.4.0.
    The public release remains attached to `24ac9fc494`.

---

## 4. Current Validated / Known State

Validated / known at the v0.4.22 checkpoint:

    [RELEASE] v0.4.0 Gitea release remains the current public release.
    [PACKAGE] v0.5.0 package staging helper exists at carbonstack/scripts/stage-v0.5.0-package.sh.
    [PACKAGE] v0.5.0 package rehearsal helper exists at carbonstack/scripts/rehearse-v0.5.0-package.sh.
    [PACKAGE] v0.5.0 staged package skeleton was generated from clean live tracked repos.
    [PACKAGE] release/checksums.txt was generated in the staged package root.
    [PACKAGE] staged-package checksum verification passed.
    [PACKAGE] the staged package was archived to /tmp/carbonstack-v0.5.0-rehearsal/archive/carbonstack-v0.5.0-package-rehearsal.tar.gz.
    [PACKAGE] the archive was extracted fresh to /tmp/carbonstack-v0.5.0-rehearsal/extract/package.
    [PACKAGE] fresh-extraction checksum verification passed.
    [PACKAGE] full passed from the fresh extraction with --clean-generated.
    [PACKAGE] final generated/private/build root absence checks passed.
    [PACKAGE] carbonstack-os stayed excluded from the runnable package.
    [DOCS] docs/167-v0.5.0-package-checksum-and-fresh-extraction-rehearsal-v0.md records the v0.4.22 package rehearsal.
    [DOCS] scripts/README.md records the v0.5.0 release-helper system.
    [REGISTRY] registry/commands.v0.yaml includes carbonstack.script.rehearse-v0.5.0-package.
    [REGISTRY] command registry Go test passed after adding the new helper entry.
    [RUNNER] full remains release-snapshot followed by local-cypher.
    [RUNNER] dev-runtime-openmls and dev-runtime-openmls-wrappers remain live-umbrella-only and are not included in full.
    [MIRRORS] Gitea remains authoritative; GitHub repos are secondary push mirrors.

v0.4.22 validation and staging observed from the logs:

    Recovery/recon:
      Partial failed v0.4.22 state was restored.
      docs/README.md sanity was restored and preserved.
      registry/commands.v0.yaml retained the v0.4.21 stage helper entry.
      Existing stage helper passed from clean tree.
      release/checksums.txt was intentionally absent during the stage-only check.
      carbonstack-os remained excluded.

    Temporary rehearsal proof:
      Temporary helper was run outside the repo to avoid the clean-tree guard.
      stage-v0.5.0-package.sh passed.
      write-checksums passed with 317 checksum entries.
      verify-checksums passed in staged package.
      archive was created.
      fresh extraction succeeded.
      verify-checksums passed from fresh extraction.
      full --clean-generated passed from fresh extraction.
      final generated/private/build root absence checks passed.

    Commit/push:
      c6aa4e3 test: rehearse v0.5.0 package validation.
      Six files changed.
      New docs/167-v0.5.0-package-checksum-and-fresh-extraction-rehearsal-v0.md.
      New scripts/README.md.
      New executable scripts/rehearse-v0.5.0-package.sh.
      Updated registry/commands.v0.yaml, docs/README.md, roadmap/ROADMAP.md.
      First push failed authentication.
      Second push succeeded to origin/main.

    Final post-commit rehearsal verification:
      git status --short was clean.
      bash -n scripts/rehearse-v0.5.0-package.sh passed.
      scripts/rehearse-v0.5.0-package.sh passed.
      PASS: v0.5.0 package rehearsal completed.
      package_root: /tmp/carbonstack-v0.5.0-rehearsal/stage/package.
      archive_path: /tmp/carbonstack-v0.5.0-rehearsal/archive/carbonstack-v0.5.0-package-rehearsal.tar.gz.
      extracted_package_root: /tmp/carbonstack-v0.5.0-rehearsal/extract/package.
      Final repo heads were clean across all four repos.

Expected live umbrella validation commands remain:

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go test ./... -count=1
    go run . --profile dev-runtime-openmls-wrappers --clean-generated
    go run . --profile dev-runtime-openmls --clean-generated
    go run . --profile local-cypher
    go run . --profile doctor
    go run . --profile core --clean-generated

Expected v0.5.0 package staging command:

    cd ~/repos/carbonstack_umbrella/carbonstack
    scripts/stage-v0.5.0-package.sh

Expected v0.5.0 package rehearsal command:

    cd ~/repos/carbonstack_umbrella/carbonstack
    scripts/rehearse-v0.5.0-package.sh

Expected v0.5.0 package validation/rehearsal shape:

    cd /tmp/carbonstack-v0.5.0-rehearsal/stage/package/carbonstack/tools/carbonstack-validate
    go run . --profile write-checksums --root /tmp/carbonstack-v0.5.0-rehearsal/stage/package
    go run . --profile verify-checksums --root /tmp/carbonstack-v0.5.0-rehearsal/stage/package

    cd /tmp/carbonstack-v0.5.0-rehearsal/extract/package/carbonstack/tools/carbonstack-validate
    go run . --profile verify-checksums --root /tmp/carbonstack-v0.5.0-rehearsal/extract/package
    go run . --profile full --root /tmp/carbonstack-v0.5.0-rehearsal/extract/package --clean-generated

---

## 5. v0.4.22 Work Completed

### 5.1 Failed partial v0.4.22 attempt was cleaned up

The earlier v0.4.22 attempt created untracked partial files and triggered the v0.4.21 staging helper’s clean-tree guard. This was recovered by:

    restoring tracked docs/README.md, registry/commands.v0.yaml, and roadmap/ROADMAP.md to HEAD;
    removing untracked partial scripts/rehearse-v0.5.0-package.sh;
    removing untracked partial docs/167-v0.5.0-package-checksum-and-fresh-extraction-rehearsal-v0.md;
    confirming docs/README.md was sane again;
    confirming the registry still contained the v0.4.21 stage helper entry.

The key lesson was that the staging helper correctly refuses dirty source trees, so a rehearsal helper that lives inside the repo cannot be tested before commit unless run from outside the repo or the tree is clean.

### 5.2 Existing stage helper was revalidated

From a clean tree:

    bash -n scripts/stage-v0.5.0-package.sh passed.
    scripts/stage-v0.5.0-package.sh passed.
    The staged release files existed.
    release/checksums.txt was intentionally deferred.
    carbonstack-os was excluded.
    final status remained clean.

This confirmed v0.4.21 staging infrastructure still worked before adding v0.4.22 rehearsal logic.

### 5.3 Temporary out-of-repo rehearsal helper proved the full flow

A temporary helper was written under `/tmp/rehearse-v0.5.0-package.sh` and run from outside the repo to avoid dirtying the source tree before stage.

It performed:

    stage package skeleton from clean live repos
    write release/checksums.txt in staged package
    verify release/checksums.txt in staged package
    archive staged package
    fresh extract archive
    verify release/checksums.txt from fresh extraction
    run full validation from fresh extraction with --clean-generated
    confirm release metadata exists
    confirm carbonstack-os excluded
    confirm final generated/private/build roots absent

Observed output included:

    checksum entries: 317
    checksum verification passed: 317 file(s)
    PASS: v0.5.0 package rehearsal completed

### 5.4 Rehearsal helper was committed

New file:

    carbonstack/scripts/rehearse-v0.5.0-package.sh

Purpose:

    Repeatably perform the v0.5.0 package checksum/archive/fresh-extraction/full validation rehearsal.

Default rehearsal root:

    /tmp/carbonstack-v0.5.0-rehearsal

Default stage package root:

    /tmp/carbonstack-v0.5.0-rehearsal/stage/package

Default archive:

    /tmp/carbonstack-v0.5.0-rehearsal/archive/carbonstack-v0.5.0-package-rehearsal.tar.gz

Default fresh extraction:

    /tmp/carbonstack-v0.5.0-rehearsal/extract/package

Important invariant:

    The helper runs the existing stage helper first.
    The stage helper requires clean live source repos.
    The helper writes checksums in the staged package root.
    The helper archives the staged package.
    The helper runs verify-checksums and full from the fresh extraction.
    The helper checks final absence of known OpenMLS sidecar generated/build/private roots.

### 5.5 Release helper system was documented

New file:

    carbonstack/scripts/README.md

It records the helper split:

    scripts/stage-v0.5.0-package.sh
    scripts/rehearse-v0.5.0-package.sh

It explains:

    stage-v0.5.0-package.sh stages a clean v0.5.0 package skeleton from tracked carbonstack, carbonstack-comms, and carbonstack-cypher files;
    it writes release metadata skeleton files;
    it intentionally excludes carbonstack-os;
    rehearse-v0.5.0-package.sh runs the release rehearsal flow:
      stage package skeleton;
      write release/checksums.txt;
      verify checksums in staged package;
      archive staged package;
      fresh extract archive;
      verify checksums from fresh extraction;
      run full from fresh extraction with --clean-generated.

It preserves boundaries:

    these scripts do not cut the release;
    do not upload release assets;
    do not make runtime OpenMLS profiles part of full;
    do not create local-backbone;
    do not start PQ/state/vault implementation.

### 5.6 v0.4.22 rehearsal doc was added

New file:

    carbonstack/docs/167-v0.5.0-package-checksum-and-fresh-extraction-rehearsal-v0.md

Records:

    v0.4.22 builds on v0.4.21 package staging;
    the new rehearsal helper exists;
    the rehearsal flow stages, checksums, verifies, archives, fresh-extracts, verifies again, and runs full;
    the v0.4.0/v0.3.36 continuity rule remains: do not treat the live umbrella as the release package root, use a throwaway staged package root, write checksums only in the package root, archive the package, and validate from fresh extraction;
    during full, release-snapshot runs core before local-cypher, and known generated roots may temporarily appear but must be removed by --clean-generated by the end;
    release-helper commands are part of the known-good command registry/readme surface;
    v0.4.22 does not cut release, upload assets, create final notes, export final LogDoc, include carbonstack-os, add runtime profiles to full, create local-backbone, or start v0.5.x state/trust/vault/PQ implementation.

### 5.7 Registry, docs index, and roadmap were updated

Updated files:

    carbonstack/registry/commands.v0.yaml
    carbonstack/docs/README.md
    carbonstack/roadmap/ROADMAP.md

Registry new entry:

    carbonstack.script.rehearse-v0.5.0-package

Classification:

    kind: script
    audience: dev
    maturity: experimental
    introduced_in: v0.4.22
    source_path: carbonstack/scripts/rehearse-v0.5.0-package.sh
    validation_surface: v0.5.0 package checksum/archive/fresh-extraction rehearsal
    include_in_front_readme: false

Nonclaims:

    does not cut release.
    does not upload assets.
    does not include carbonstack-os.
    does not make runtime OpenMLS profiles part of full.
    does not create local-backbone.

Docs index now includes:

    docs/167-v0.5.0-package-checksum-and-fresh-extraction-rehearsal-v0.md

Roadmap now records:

    v0.4.22 v0.5.0 checksum and fresh extraction rehearsal.
    The release-helper system has both stage and rehearse helpers.
    v0.4.23 should handle final release notes, LogDoc sanitization/export, release asset generation, and v0.5.0 release cut prep.

### 5.8 Validation passed and final post-commit rehearsal was confirmed

Before commit:

    temporary out-of-repo rehearsal passed.
    command registry Go test passed after adding the new helper registry entry.
    git commit succeeded.

Commit/push:

    c6aa4e3 test: rehearse v0.5.0 package validation.
    First push failed authentication.
    Second push succeeded.

After commit:

    git status --short was clean.
    bash -n scripts/rehearse-v0.5.0-package.sh passed.
    scripts/rehearse-v0.5.0-package.sh passed.
    PASS: v0.5.0 package rehearsal completed.
    final four-repo snapshot was clean.

---

## 6. Critical Blunders / Lessons to Carry Forward

### 6.1 Clean-tree guard must be respected during helper development

The v0.4.21 stage helper requires clean `carbonstack`, `carbonstack-comms`, and `carbonstack-cypher` repos. Any untracked new helper or doc inside `carbonstack` makes the stage helper fail.

Lesson:

    When adding scripts that call the stage helper, prove the flow with a temporary helper under /tmp first.
    Then commit the known-good helper.
    Then run the committed helper from a clean tree after commit.

### 6.2 Previous repair attempts caused copy/paste corruption

The session included visible paste corruption such as command fragments and malformed script/doc tails. A prior repair also risked collapsing docs/README.md into escaped newline text before the tracked file was restored.

Lesson:

    Use smaller blocks.
    Avoid complex nested heredocs where possible.
    Do not continue after a failed block if files may be partially corrupted.
    Always inspect `git diff` and run `bash -n` before committing scripts.
    Restore tracked files before retrying if paste corruption is suspected.

### 6.3 Temporary helper output contained visible pasted command corruption, but committed helper passed

The temporary helper creation paste showed visible corrupted text in the log, but the executed flow still reached the intended rehearsal success markers, and the committed helper passed post-commit from a clean tree.

Lesson:

    Treat the committed helper plus post-commit clean-tree run as the source of truth.
    Do not treat transient pasteback corruption as a blocker if the final committed source and post-commit run pass.

### 6.4 Release root discipline is now explicit

The v0.4.0/v0.3.36 continuity rule remains important:

    stage the package;
    write checksums in package root;
    archive the package;
    validate from fresh extraction;
    do not run release-snapshot/full against the package source root intended for publish/archive.

Lesson:

    The final v0.5.0 release package should be created from a clean staged source root and validated from fresh extraction.
    Release-snapshot/full generate artifacts, so they must not be run inside the source root that will be archived/published.

### 6.5 Commit metadata warning remains

The commit log again showed auto-generated local commit metadata:

    Committer: bitcrusher32 <▓▓>

This is not a functional validation failure, but remains a professionalism/publishing hygiene issue to address eventually.

### 6.6 First push auth failure remains a recurring non-blocker

The first push again failed authentication and a later push succeeded.

Lesson:

    Do not mark the rung failed solely on first push auth failure if a later push succeeds and final heads/status are clean.
    Still record it for blunder/timeline continuity.

---

## 7. Current Blockers / Not Validated

Current blockers and nonvalidated areas after v0.4.22:

    v0.5.0 package rehearsal has passed with checksums/archive/fresh extraction/full validation, but final release assets are not generated yet.
    v0.5.0 final release notes are not written yet.
    v0.5.0 final LogDoc sanitization/export is not done yet.
    v0.5.0 tag/release is not cut.
    v0.4.xPRIME compression into LogDoc v0.5.0 is not done yet.
    Existing send/inbox remain stub-era.
    Runtime Comms has dev-only OpenMLS commands, direct smoke proof, wrapper-based smoke proof, and separate manual runner profiles, but no mature user-facing UX.
    Neither runtime profile is included in full.
    Neither runtime profile is part of public release package validation.
    No Relay Space join model exists yet.
    local-backbone runner profile does not exist and remains reserved.
    Generated command reference from the registry does not exist yet.
    Local help/manual generation from the registry does not exist yet.
    v0.5.x provider/trust/vault and PQ/hybrid migration readiness remain future work.

Hard nonclaims remain:

    Do not claim production readiness.
    Do not claim production E2EE.
    Do not claim hostile-server safety.
    Do not claim metadata privacy.
    Do not claim Debian production deployability.
    Do not claim systemd readiness.
    Do not claim cloudflared readiness.
    Do not claim Android readiness.
    Do not claim CarbonStackOS readiness.
    Do not claim secure vault/storage.
    Do not claim real homelab validation.
    Do not claim external audit or certification.
    Do not claim CarbonStack Relay Space mechanics are implemented.
    Do not claim local-backbone exists or is justified yet.
    Do not claim local-cypher validates runtime Comms UX.
    Do not claim full validates local-backbone.
    Do not claim full is a deployment command.
    Do not claim dev-runtime-openmls is release-package validation yet.
    Do not claim dev-runtime-openmls-wrappers is release-package validation yet.
    Do not claim send/inbox are OpenMLS-backed yet.
    Do not claim PQ/hybrid security or quantum-safe messaging.
    Do not let GitHub mirrors supersede Gitea source-of-truth status.

Allowed claim:

    CarbonStack current public release remains v0.4.0 broad local deployability pre-release.
    CarbonStack mainline is now at v0.4.22 package checksum/archive/fresh-extraction rehearsal.
    The v0.5.0 package rehearsal helper exists and passes from a clean tree.
    The package rehearsal writes release/checksums.txt, verifies checksums in staged package, archives package, fresh-extracts, verifies checksums again, and runs full from fresh extraction.
    The package root includes carbonstack, carbonstack-comms, carbonstack-cypher, and release metadata.
    carbonstack-os is explicitly excluded from the runnable package.
    Final generated/private/build roots are absent after --clean-generated.
    Gitea remains source of truth; GitHub mirrors are secondary push mirrors.

---

## 8. Next Safest Actions

Recommended next rung:

    v0.4.23 final release notes, LogDoc sanitization/export, release asset generation, and v0.5.0 release cut prep.

Focus:

    Use v0.4.0 release screenshot/raw MD/link as continuity reference.
    Formulate v0.5.0 release notes in v0.4.0 style.
    Generate final release package/assets with release naming, likely using the rehearsal helper as basis.
    Generate package checksum/asset checksum instructions as appropriate.
    Sanitize/export final LogDoc as v0.4.xPRIME if needed.
    Prepare the final v0.5.0 release asset set.
    Preserve v0.5.0 nonclaims.
    Keep runtime profiles outside full.
    Keep carbonstack-os excluded from the runnable package.
    Do not cut release until final notes/assets/checks are agreed and verified.

Avoid next:

    cutting v0.5.0 before final release notes/assets are prepared.
    running release-snapshot/full inside the archive source root intended for publication.
    adding dev-runtime-openmls or dev-runtime-openmls-wrappers to full without a deliberate release-package validation decision.
    calling the package local-backbone.
    starting PQ/state/vault work before v0.5.0 release prep finishes.
    including carbonstack-os in the runnable package.
    making general-public usable claims.
    rewriting historical numbered docs merely because they are stale.

---

## 9. Critical Paths / Functions

### Source-of-truth public surfaces

    [GITEA SOURCE OF TRUTH] https://git.bitcrusher32.win/bitcrusher32/carbonstack
    [RELEASES] https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases
    [CURRENT PUBLIC RELEASE] https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.4.0

GitHub mirrors:

    secondary push mirrors for discoverability/redundancy only.
    Do not treat mirrors as release authority unless project policy changes.
    User intends official general-public releases later on GitHub, such as v1.0.0, but current v0.5.0 is not that.

### WSL Debian active working umbrella

    [WSL-DEBIAN] umbrella: ~/repos/carbonstack_umbrella
    [WSL-DEBIAN] carbonstack: ~/repos/carbonstack_umbrella/carbonstack
    [WSL-DEBIAN] carbonstack-comms: ~/repos/carbonstack_umbrella/carbonstack-comms
    [WSL-DEBIAN] carbonstack-cypher: ~/repos/carbonstack_umbrella/carbonstack-cypher
    [WSL-DEBIAN] carbonstack-os: ~/repos/carbonstack_umbrella/carbonstack-os
    [GO-RUNNER] runner dir: ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate

### v0.5.0 package staging/rehearsal surfaces

    [STAGING SCRIPT] carbonstack/scripts/stage-v0.5.0-package.sh
    [REHEARSAL SCRIPT] carbonstack/scripts/rehearse-v0.5.0-package.sh
    [SCRIPTS README] carbonstack/scripts/README.md
    [STAGING DOC] carbonstack/docs/166-v0.5.0-package-rehearsal-plan-v0.md
    [REHEARSAL DOC] carbonstack/docs/167-v0.5.0-package-checksum-and-fresh-extraction-rehearsal-v0.md
    [REGISTRY ENTRY] carbonstack.script.stage-v0.5.0-package in carbonstack/registry/commands.v0.yaml
    [REGISTRY ENTRY] carbonstack.script.rehearse-v0.5.0-package in carbonstack/registry/commands.v0.yaml
    [DEFAULT REHEARSAL ROOT] /tmp/carbonstack-v0.5.0-rehearsal
    [DEFAULT STAGE ROOT] /tmp/carbonstack-v0.5.0-rehearsal/stage
    [DEFAULT PACKAGE ROOT] /tmp/carbonstack-v0.5.0-rehearsal/stage/package
    [DEFAULT ARCHIVE] /tmp/carbonstack-v0.5.0-rehearsal/archive/carbonstack-v0.5.0-package-rehearsal.tar.gz
    [DEFAULT FRESH EXTRACTION ROOT] /tmp/carbonstack-v0.5.0-rehearsal/extract/package

### Current live public docs surfaces

    [DOC] top README:
    carbonstack/README.md

    [DOC] docs index:
    carbonstack/docs/README.md

    [DOC] roadmap:
    carbonstack/roadmap/ROADMAP.md

    [RUNNER] runner README:
    carbonstack/tools/carbonstack-validate/README.md

    [REGISTRY] registry README:
    carbonstack/registry/README.md

### Registry surfaces

    [REGISTRY] root:
    carbonstack/registry

    [REGISTRY] provisional command registry:
    carbonstack/registry/commands.v0.yaml

    [REGISTRY] registry README:
    carbonstack/registry/README.md

    [REGISTRY-TEST] command registry validation:
    carbonstack/tools/carbonstack-validate/command_registry_test.go

### Runner surfaces

    [RUNNER] root:
    carbonstack/tools/carbonstack-validate

    [RUNNER] dispatch:
    carbonstack/tools/carbonstack-validate/main.go

    [RUNNER] local-cypher:
    carbonstack/tools/carbonstack-validate/local_cypher.go

    [RUNNER] release-snapshot:
    carbonstack/tools/carbonstack-validate/release_snapshot.go

    [RUNNER] checksums:
    carbonstack/tools/carbonstack-validate/checksums.go

### Comms runtime surfaces

    [COMMS] direct-sidecar dev runtime smoke proof:
    carbonstack-comms/scripts/dev-openmls-runtime-smoke.sh

    [COMMS] wrapper-based dev runtime smoke proof:
    carbonstack-comms/scripts/dev-openmls-runtime-smoke-wrappers.sh

### Known-good commands

    [LIVE VALIDATION]
    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go test ./... -count=1
    go run . --profile dev-runtime-openmls-wrappers --clean-generated
    go run . --profile dev-runtime-openmls --clean-generated
    go run . --profile local-cypher
    go run . --profile doctor
    go run . --profile core --clean-generated

    [PACKAGE STAGING]
    cd ~/repos/carbonstack_umbrella/carbonstack
    scripts/stage-v0.5.0-package.sh

    [PACKAGE REHEARSAL]
    cd ~/repos/carbonstack_umbrella/carbonstack
    scripts/rehearse-v0.5.0-package.sh

    [STAGED CHECKSUM SHAPE]
    cd /tmp/carbonstack-v0.5.0-rehearsal/stage/package/carbonstack/tools/carbonstack-validate
    go run . --profile write-checksums --root /tmp/carbonstack-v0.5.0-rehearsal/stage/package
    go run . --profile verify-checksums --root /tmp/carbonstack-v0.5.0-rehearsal/stage/package

    [FRESH EXTRACTION VALIDATION SHAPE]
    cd /tmp/carbonstack-v0.5.0-rehearsal/extract/package/carbonstack/tools/carbonstack-validate
    go run . --profile verify-checksums --root /tmp/carbonstack-v0.5.0-rehearsal/extract/package
    go run . --profile full --root /tmp/carbonstack-v0.5.0-rehearsal/extract/package --clean-generated

---

## 10. v0.4.x Forward Plan

### Immediate next: v0.4.23 final release notes + LogDoc/export/assets

Expected:

    Use the v0.4.0 release screenshot/raw MD/link as continuity reference.
    Draft v0.5.0 release notes with v0.4.0 style and v0.5.0 actual evidence.
    Prepare final LogDoc sanitization/export, likely v0.4.xPRIME.
    Generate final v0.5.0 release package/assets.
    Run final release package validation and asset/checksum checks.
    Confirm Gitea release plan.
    Do not cut release until final template-aligned notes and final package validation are accepted.

### v0.5.0

Expected:

    minor epoch release.
    v0.4.xPRIME compression into LogDoc v0.5.0.
    Gitea official pre-release.
    GitHub remains secondary mirror.
    Not general-public usable.
    Not v1.0.0.

---

## 11. v0.5.x and Later Carry-Forward

### v0.5.0 minor epoch release

After v0.4.23 final release-prep, cut a v0.5.0 minor epoch release using the accumulated v0.4.x work.

Dedicated remaining work should include:

    final release notes using v0.4.0 continuity style.
    final LogDoc sanitization/export.
    release asset generation.
    release cut.
    v0.4.xPRIME compression into LogDoc v0.5.0.

### v0.5.x implementation after release

Local storage, trust state, provider state, vault design, and PQ/hybrid ciphersuite migration readiness.

Do not start PQ/hybrid implementation before the runtime OpenMLS/Cypher path exists and is documented. PQ/hybrid migration touches provider state, ciphersuite metadata, trust display, artifact size, reinit/rekey/recovery policy, and claim boundaries.

### v0.6.x+

Hostile-server and abuse-resistance harnesses.

### v0.7.x+

Deployability and operations hardening.

### v0.8.x+

Documented self-pentest / adversarial validation.

### v0.9.x+

Claim-boundary review.

### v0.10.x+

Android backend/app work.

### v1.x.x

Public app/server major epoch and much later CarbonStackOS/appliance research, if justified.

---

## 12. Lean Breakpoint Summary

    CarbonStack v0.4.22 is the v0.5.0 package checksum/archive/fresh-extraction rehearsal checkpoint. v0.4.0 remains the current public broad local deployability pre-release. carbonstack is now at c6aa4e3 test: rehearse v0.5.0 package validation. carbonstack-comms remains at cb4e59d test: add wrapper-based OpenMLS runtime smoke proof. v0.4.22 adds scripts/rehearse-v0.5.0-package.sh, scripts/README.md, docs/167-v0.5.0-package-checksum-and-fresh-extraction-rehearsal-v0.md, a command registry entry for carbonstack.script.rehearse-v0.5.0-package, docs index update, and roadmap update. The rehearsal helper stages a clean v0.5.0 package, writes release/checksums.txt, verifies staged checksums, archives the package, fresh-extracts it, verifies fresh checksums, runs full --clean-generated from fresh extraction, confirms carbonstack-os is excluded, and confirms known final generated/private/build roots are absent. The first push failed auth and the second push succeeded. Post-commit clean-tree helper execution passed and final repo heads were clean across all four repos. Next safe action: v0.4.23 final release notes using v0.4.0 continuity reference, final LogDoc sanitization/export, release asset generation, and v0.5.0 release cut prep.

---

## 13. Preserved Immediate Previous Handoff: v0.4.21

The following is the previous v0.4.21 handoff. Where it conflicts with the v0.4.22 overlay above, v0.4.22 wins for current state.



# CarbonStack LogDoc v0.4.21

**Last updated:** 2026-06-05 local session  
**Owner/operator:** bitcrusher32  
**Assistant/model:** GPT-5.5 Thinking  
**Current phase:** Phase 2F / **v0.5.0 package rehearsal plan and staging checkpoint complete**. The v0.4.0 public broad local deployability pre-release remains the current public release. v0.4.20 refreshed roadmap/current release surfaces; v0.4.21 starts the compressed v0.5.0 minor-epoch release runway by adding a repeatable v0.5.0 package staging helper, recording the package rehearsal plan, updating the command registry/docs/roadmap, and confirming the clean package skeleton stages successfully after commit. `carbonstack` is now at `2c2554d (HEAD -> main, origin/main, origin/HEAD) chore: stage v0.5.0 package rehearsal`; `carbonstack-comms` remains at `cb4e59d (HEAD -> main, origin/main, origin/HEAD) test: add wrapper-based OpenMLS runtime smoke proof`; `carbonstack-cypher` remains at `9ab994c (HEAD -> main, origin/main, origin/HEAD) docs: point to local operator runbook`; `carbonstack-os` remains at `1bbbe52 (HEAD -> main, origin/main, origin/HEAD) docs: clarify CarbonStackOS target direction`.

**Version schema:** `v[scope].[timeline]`. This file is `v0.4.21`, the v0.5.0 package rehearsal plan / asset inventory / staging implementation checkpoint after the v0.4.20 roadmap refresh and release-surface cleanup checkpoint. Per LogDoc V2 practice, this Markdown file preserves current operational continuity and carries the v0.4.x release-prep thread toward v0.5.0; the JSON breakpoint remains the lean machine-readable handoff state.

**Compression note:** The detailed v0.3.0 through v0.3.36PRIME process ledger remains archived in PRIME/sanitized PRIME. v0.4.0 compressed that pre-release sprint; v0.4.1 through v0.4.12 preserved runtime recon, command contracts, dev OpenMLS send/inbox commands, smoke proof, runner profile, boundary checks, wrapper recon, and bootstrap command contract. v0.4.13 through v0.4.15 preserved staged bootstrap wrapper implementation. v0.4.16 preserved the first committed wrapper-based runtime smoke variant. v0.4.17 preserved the separate manual runner profile for that wrapper smoke. v0.4.18 preserved the first provisional cross-repo command registry. v0.4.19 preserved registry validation/test coverage and local-help/manual planning. v0.4.20 preserved roadmap refresh and release-surface current-state cleanup. v0.4.21 now preserves v0.5.0 package rehearsal staging.

---

## 1. Project Goal

**Active goal:** Execute the compressed late-v0.4.x release-prep runway toward a v0.5.0 minor epoch pre-release, while preserving v0.4.0 release-style continuity, package validation discipline, Gitea source-of-truth status, and strict nonclaims.

CarbonStack’s broader project intent remains:

- **CarbonStack** is the secure-communications appliance-stack project and shared doctrine/release authority.
- **CarbonStackComms** is the text-first messaging client and current OpenMLS sidecar/relay integration surface.
- **CarbonStackCypher** is the self-hostable relay/storage stack.
- **CarbonStackOS** is future constrained Android-derived appliance OS work and remains out of current runnable validation packages.
- The `carbonstack` Gitea repo remains the source of truth for release state, release assets, docs, registry, and public claims.
- GitHub mirrors exist as secondary push mirrors only; Gitea remains authoritative for current pre-release work.

**North star:**

> Its goal is not to be a better smartphone. Its goal is to stop being a smartphone.

**Core doctrine:**

> Every feature is guilty until it proves it does not add unacceptable parser, network, sensor, identity, or filesystem authority.

**Current reality after v0.4.21:**

    carbonstack/scripts/stage-v0.5.0-package.sh exists and is executable.
    The helper stages a clean v0.5.0 package skeleton from tracked files using git archive.
    It requires clean live git checkouts for carbonstack, carbonstack-comms, and carbonstack-cypher.
    It intentionally excludes carbonstack-os from the runnable package.
    It writes release metadata skeleton files under release/.
    It checks for forbidden generated/private/build artifacts in the staged package.
    It does not generate final checksums, archive the package, validate fresh extraction, cut a tag, or create final release assets.

    carbonstack/docs/166-v0.5.0-package-rehearsal-plan-v0.md records the checkpoint.
    carbonstack/registry/commands.v0.yaml now includes carbonstack.script.stage-v0.5.0-package.
    carbonstack/docs/README.md indexes docs/166.
    carbonstack/roadmap/ROADMAP.md now uses the compressed late-v0.4.x release-prep runway:
      v0.4.21 package rehearsal plan, asset inventory, and package staging implementation.
      v0.4.22 checksum generation, fresh extraction validation, and release notes formulation using v0.4.0 continuity reference.
      v0.4.23 final LogDoc sanitization/export and release asset generation.
      v0.5.0 minor epoch release, followed by v0.4.xPRIME compression into LogDoc v0.5.0.

    Verified staged package path:
      /tmp/carbonstack-v0.5.0-rehearsal/stage/package

    Verified staged release files:
      release/LICENSE
      release/manifest.json
      release/release-notes-draft.md
      release/testing-runbook.md
      release/validation-freeze.md

    Confirmed intentionally absent:
      release/checksums.txt

    Confirmed excluded:
      carbonstack-os

Hard nonclaims remain: v0.4.21 does **not** prove production readiness, production E2EE, hostile-server safety, metadata privacy, Android readiness, CarbonStackOS readiness, secure vault/storage, public ingress safety, real homelab validation, audit, certification, PQ/hybrid security, mature messenger UX, release-package runtime OpenMLS validation, Relay Space join/onboarding, `full` runtime OpenMLS validation, or local-backbone readiness. This is package staging infrastructure, not final release validation.

---

## 2. Current Public Release State

Public release:

    https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.4.0

Release title:

    CarbonStack v0.4.0 Broad Local Deployability Pre-Release

Release tag:

    v0.4.0

Release commit:

    24ac9fc494 / 24ac9fc docs: record v0.4.0 package rehearsal

Gitea release status:

    Pre-Release

Important release-page warning remains:

    The default Gitea Source Code ZIP/TAR.GZ downloads are auto-generated carbonstack-only archives.
    The intended runnable validation artifact is the attached multi-repo release package plus companion assets.

Post-release public surface:

    carbonstack README is evergreen.
    It points interested parties to the Releases page for runnable / known-good artifacts.
    Gitea is source of truth.
    GitHub mirrors are secondary push mirrors, not release authority.

Official general-public usable releases:

    Not yet.
    User explicitly intends official usable/general-public releases to happen on GitHub later, such as the v1.0.0 major epoch release.
    Current v0.5.0 target is a Gitea-source-of-truth pre-release/minor epoch release, not a general-public official usable release.

---

## 3. Current Repo Heads

    carbonstack        2c2554d (HEAD -> main, origin/main, origin/HEAD) chore: stage v0.5.0 package rehearsal
    carbonstack-comms  cb4e59d (HEAD -> main, origin/main, origin/HEAD) test: add wrapper-based OpenMLS runtime smoke proof
    carbonstack-cypher 9ab994c (HEAD -> main, origin/main, origin/HEAD) docs: point to local operator runbook
    carbonstack-os     1bbbe52 (HEAD -> main, origin/main, origin/HEAD) docs: clarify CarbonStackOS target direction

Current public release tag:

    v0.4.0 -> 24ac9fc494 / 24ac9fc docs: record v0.4.0 package rehearsal

Recent maturity / release-prep commits:

    v0.4.17 wrapper-smoke runner/profile alignment:
      a501442 feat: add wrapper OpenMLS runtime validation profile

    v0.4.18 command registry:
      a4162d7 docs: add known-good command registry

    v0.4.19 registry validation:
      568fb45 test: validate command registry coverage

    v0.4.20 release-surface cleanup:
      7d4fc0d docs: refresh release-surface current state

    Current v0.4.21 package staging:
      2c2554d chore: stage v0.5.0 package rehearsal

Continuity note:

    `2c2554d` is a carbonstack release-prep/staging commit after the v0.4.0 public release tag.
    It does not retag v0.4.0.
    The public release remains attached to `24ac9fc494`.

---

## 4. Current Validated / Known State

Validated / known at the v0.4.21 checkpoint:

    [RELEASE] v0.4.0 Gitea release remains the current public release.
    [PACKAGE] v0.5.0 package staging helper exists at carbonstack/scripts/stage-v0.5.0-package.sh.
    [PACKAGE] v0.5.0 staged package skeleton successfully staged after commit.
    [PACKAGE] staged package includes carbonstack, carbonstack-comms, carbonstack-cypher, and release metadata skeleton.
    [PACKAGE] carbonstack-os was explicitly excluded from the runnable package.
    [PACKAGE] release/checksums.txt was intentionally absent/deferred.
    [PACKAGE] release/manifest.json, validation-freeze.md, testing-runbook.md, release-notes-draft.md, and LICENSE were present.
    [PACKAGE] forbidden generated/private/build artifact check passed.
    [DOCS] docs/166-v0.5.0-package-rehearsal-plan-v0.md records the v0.4.21 package rehearsal plan and staging implementation.
    [REGISTRY] registry/commands.v0.yaml now includes carbonstack.script.stage-v0.5.0-package.
    [ROADMAP] roadmap/ROADMAP.md records the compressed late-v0.4.x release-prep runway.
    [RUNNER] full remains release-snapshot followed by local-cypher.
    [RUNNER] dev-runtime-openmls and dev-runtime-openmls-wrappers remain live-umbrella-only and are not included in full.
    [MIRRORS] Gitea remains authoritative; GitHub repos are secondary push mirrors.

v0.4.21 validation and staging observed from the logs:

    Initial attempt:
      stage helper failed before commit because carbonstack had unstaged changes.
      failure was expected/correct behavior from the clean-tree guard.
      no staged package existed from that attempt.

    Patch/validation before commit:
      bash -n scripts/stage-v0.5.0-package.sh passed.
      live validation passed:
        go test ./... -count=1 in carbonstack/tools/carbonstack-validate.
        dev-runtime-openmls-wrappers --clean-generated.
        dev-runtime-openmls --clean-generated.
        local-cypher.
        doctor.
        core --clean-generated.
        go test ./internal/app -count=1 in carbonstack-comms.
        go test ./... -count=1 -timeout 600s in carbonstack-comms.

    Commit/push:
      2c2554d chore: stage v0.5.0 package rehearsal.
      Five files changed.
      New docs/166-v0.5.0-package-rehearsal-plan-v0.md.
      New executable scripts/stage-v0.5.0-package.sh.
      Updated docs/README.md, registry/commands.v0.yaml, roadmap/ROADMAP.md.
      Push succeeded to origin/main.

    Final post-commit staging verification:
      git status --short was clean.
      bash -n scripts/stage-v0.5.0-package.sh passed.
      scripts/stage-v0.5.0-package.sh passed.
      stage_root: /tmp/carbonstack-v0.5.0-rehearsal/stage.
      package_root: /tmp/carbonstack-v0.5.0-rehearsal/stage/package.
      release_dir: /tmp/carbonstack-v0.5.0-rehearsal/stage/package/release.
      Package skeleton file check passed for README files and release metadata.
      Forbidden generated/private/build artifact check passed.
      PASS: v0.5.0 package skeleton staged.
      release/checksums.txt intentionally deferred.
      carbonstack-os excluded.
      Final repo heads were clean across all four repos.

Expected live umbrella validation commands remain:

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go test ./... -count=1
    go run . --profile dev-runtime-openmls-wrappers --clean-generated
    go run . --profile dev-runtime-openmls --clean-generated
    go run . --profile local-cypher
    go run . --profile doctor
    go run . --profile core --clean-generated

Expected v0.5.0 package staging command:

    cd ~/repos/carbonstack_umbrella/carbonstack
    scripts/stage-v0.5.0-package.sh

Expected v0.4.0 public package validation commands remain unchanged:

    cd <package-root>/carbonstack/tools/carbonstack-validate
    go run . --profile verify-checksums --root <package-root>
    go run . --profile full --root <package-root> --clean-generated

---

## 5. v0.4.21 Work Completed

### 5.1 Package/release shape scout confirmed the v0.5.0 package direction

The scout confirmed the existing release runner still expects a release-like package root containing:

    carbonstack/
    carbonstack-comms/
    carbonstack-cypher/
    release/

The release metadata checks still require these groups:

    release/manifest.json
    release/checksums.txt
    release/validation-freeze.md OR release/testing-runbook.md

The v0.4.0 package rehearsal doc established the precedent:

    package/
      carbonstack/
      carbonstack-comms/
      carbonstack-cypher/
      release/
        LICENSE
        manifest.json
        release-notes.md
        testing-runbook.md
        validation-freeze.md
        checksums.txt

v0.4.21 kept that package shape for v0.5.0 but deferred checksums, archive creation, fresh extraction validation, and final release notes/assets to v0.4.22/v0.4.23.

### 5.2 Staging helper was added

New file:

    carbonstack/scripts/stage-v0.5.0-package.sh

Purpose:

    Stage a v0.5.0 package skeleton from tracked repo contents.

Default staging root:

    /tmp/carbonstack-v0.5.0-rehearsal/stage

Default package root:

    /tmp/carbonstack-v0.5.0-rehearsal/stage/package

Included:

    carbonstack
    carbonstack-comms
    carbonstack-cypher
    release/

Excluded:

    carbonstack-os
    .git
    untracked files
    generated/private/build artifacts

Implementation details:

    Uses git archive for tracked contents.
    Requires clean live git checkouts before staging.
    Requires git, tar, and python3.
    Copies carbonstack/LICENSE into release/LICENSE.
    Writes release metadata skeleton with an inline Python metadata writer.
    Checks for forbidden generated/private/build artifacts in the staged package.
    Prints explicit nonclaims and deferred status through release metadata.

### 5.3 Release metadata skeleton was staged

Generated under the staged package root:

    release/manifest.json
    release/validation-freeze.md
    release/testing-runbook.md
    release/release-notes-draft.md
    release/LICENSE

Intentional absence:

    release/checksums.txt

Deferred:

    final checksums
    archive
    fresh extraction validation
    final release notes
    final LogDoc export
    final release assets
    tag/release cut

### 5.4 v0.4.21 package rehearsal plan doc was added

New file:

    carbonstack/docs/166-v0.5.0-package-rehearsal-plan-v0.md

Records:

    v0.4.21 starts the compressed v0.5.0 release runway.
    package shape remains carbonstack + carbonstack-comms + carbonstack-cypher + release.
    carbonstack-os remains excluded from the runnable package.
    scripts/stage-v0.5.0-package.sh is the new staging helper.
    release metadata skeleton files are staged.
    final checksums/archive/assets are deferred.
    expected v0.4.22 flow is checksum generation, archive/fresh extraction validation, and release notes formulation.
    v0.4.21 does not cut v0.5.0 and does not start PQ/state/vault work.

### 5.5 Registry was updated

Updated file:

    carbonstack/registry/commands.v0.yaml

New entry:

    carbonstack.script.stage-v0.5.0-package

Classification:

    kind: script
    audience: dev
    maturity: experimental
    introduced_in: v0.4.21
    source_path: carbonstack/scripts/stage-v0.5.0-package.sh
    validation_surface: v0.5.0 package rehearsal staging
    include_in_front_readme: false

Nonclaims:

    does not write final checksums.
    does not archive package.
    does not cut release.
    does not validate fresh extraction.
    does not include carbonstack-os.

### 5.6 Docs index and roadmap were updated

Updated files:

    carbonstack/docs/README.md
    carbonstack/roadmap/ROADMAP.md

Docs index now includes:

    docs/166-v0.5.0-package-rehearsal-plan-v0.md

Roadmap now records:

    compressed late-v0.4.x release-prep rungs:
      v0.4.21 package rehearsal plan, asset inventory, and package staging implementation.
      v0.4.22 checksum generation, fresh extraction validation, and release notes formulation using v0.4.0 continuity reference.
      v0.4.23 final LogDoc sanitization/export and release asset generation.
      v0.5.0 minor epoch release, followed by v0.4.xPRIME compression into LogDoc v0.5.0.

    v0.4.21 completed package rehearsal plan and staging.
    next rung is v0.4.22 checksum generation / fresh extraction validation / release notes formulation.

### 5.7 Validation passed and final post-commit staging was confirmed

Pre-commit validation passed:

    go test ./... -count=1 in carbonstack/tools/carbonstack-validate.
    go run . --profile dev-runtime-openmls-wrappers --clean-generated.
    go run . --profile dev-runtime-openmls --clean-generated.
    go run . --profile local-cypher.
    go run . --profile doctor.
    go run . --profile core --clean-generated.
    go test ./internal/app -count=1 in carbonstack-comms.
    go test ./... -count=1 -timeout 600s in carbonstack-comms.

Post-commit staging verification passed:

    bash -n scripts/stage-v0.5.0-package.sh.
    scripts/stage-v0.5.0-package.sh.
    release skeleton files existed.
    release/checksums.txt was absent as intended.
    carbonstack-os was absent as intended.
    final four-repo snapshot was clean.

---

## 6. Critical Blunders / Lessons to Carry Forward

### 6.1 Nested quote script-generation bug broke the first patch block

The first script-generation block failed with a Python `SyntaxError` because the generated shell script contained nested triple-quoted Python strings inside an outer Python triple-quoted writer.

Lesson:

    Avoid nested triple-quoted string stacks when generating scripts via Python.
    Use raw single-triple script text and avoid inner triple-quoted strings.
    Prefer smaller blocks for release engineering work.
    Stop on syntax errors before trying to continue.

### 6.2 The staging helper correctly refused dirty worktrees

The first script run before commit failed with:

    FAIL: carbonstack has unstaged changes

This was not a script defect. It was the clean-tree guard behaving correctly because the staging script/doc/registry/roadmap edits were still uncommitted.

Lesson:

    Do not run the staging helper against the package source root while it has uncommitted edits.
    For final proof, commit first, then run the staging helper from a clean tree.
    This clean-tree guard is useful and should remain.

### 6.3 Initial staging verification looked falsely clean because commands continued after failure

After the pre-commit helper failure, later checks still reported no checksums and no carbonstack-os because the package directory did not exist. That was not a valid staging success.

Lesson:

    Treat `PASS: v0.5.0 package skeleton staged` as the actual staging success marker.
    Confirm release files exist with `ls -la "$STAGED/release"`.
    Confirm final clean repo status after successful post-commit staging.
    Do not rely on absence checks if the staged package root was not created.

### 6.4 Copy-box/pasteback fragility persisted

The scout/patch session had command copy issues and the first implementation block needed replacement.

Lesson:

    Continue using smaller copy blocks.
    Avoid heredoc-heavy or nested-code generation when possible.
    For release/rung work, prefer one clear action per block.

### 6.5 Commit metadata warning remains

The commit log again showed auto-generated local commit metadata:

    Committer: bitcrusher32 <▓▓>

This is not a functional validation failure, but remains a professionalism/publishing hygiene issue to address eventually.

---

## 7. Current Blockers / Not Validated

Current blockers and nonvalidated areas after v0.4.21:

    v0.5.0 package skeleton is staged successfully, but final checksums are not generated yet.
    v0.5.0 release archive is not created yet.
    v0.5.0 fresh extraction validation has not run yet.
    v0.5.0 final release notes are not written yet.
    v0.5.0 final LogDoc sanitization/export is not done yet.
    v0.5.0 release assets are not generated yet.
    v0.5.0 tag/release is not cut.
    Existing send/inbox remain stub-era.
    Runtime Comms has dev-only OpenMLS commands, direct smoke proof, wrapper-based smoke proof, and separate manual runner profiles, but no mature user-facing UX.
    Neither runtime profile is included in full.
    Neither runtime profile is part of public release package validation.
    No Relay Space join model exists yet.
    local-backbone runner profile does not exist and remains reserved.
    Generated command reference from the registry does not exist yet.
    Local help/manual generation from the registry does not exist yet.
    v0.5.x provider/trust/vault and PQ/hybrid migration readiness remain future work.

Hard nonclaims remain:

    Do not claim production readiness.
    Do not claim production E2EE.
    Do not claim hostile-server safety.
    Do not claim metadata privacy.
    Do not claim Debian production deployability.
    Do not claim systemd readiness.
    Do not claim cloudflared readiness.
    Do not claim Android readiness.
    Do not claim CarbonStackOS readiness.
    Do not claim secure vault/storage.
    Do not claim real homelab validation.
    Do not claim external audit or certification.
    Do not claim CarbonStack Relay Space mechanics are implemented.
    Do not claim local-backbone exists or is justified yet.
    Do not claim local-cypher validates runtime Comms UX.
    Do not claim full validates local-backbone.
    Do not claim full is a deployment command.
    Do not claim dev-runtime-openmls is release-package validation yet.
    Do not claim dev-runtime-openmls-wrappers is release-package validation yet.
    Do not claim send/inbox are OpenMLS-backed yet.
    Do not claim PQ/hybrid security or quantum-safe messaging.
    Do not let GitHub mirrors supersede Gitea source-of-truth status.

Allowed claim:

    CarbonStack current public release remains v0.4.0 broad local deployability pre-release.
    CarbonStack mainline is now at v0.4.21 package rehearsal staging.
    The v0.5.0 package staging helper exists and stages a clean package skeleton.
    The package skeleton includes carbonstack, carbonstack-comms, carbonstack-cypher, and release metadata skeleton.
    carbonstack-os is explicitly excluded from the runnable package.
    release/checksums.txt is intentionally deferred to v0.4.22.
    The staged release metadata skeleton contains manifest.json, validation-freeze.md, testing-runbook.md, release-notes-draft.md, and LICENSE.
    The staging helper requires clean live git checkouts and correctly refuses dirty source trees.
    Gitea remains source of truth; GitHub mirrors are secondary push mirrors.

---

## 8. Next Safest Actions

Recommended next rung:

    v0.4.22 checksum generation, fresh extraction validation, and release notes formulation using v0.4.0 continuity reference.

Focus:

    Re-run scripts/stage-v0.5.0-package.sh from clean tree.
    Generate release/checksums.txt inside the clean staged package root.
    Verify checksums in the staged package root.
    Archive the package source root without running release-snapshot inside it.
    Extract archive into a fresh throwaway validation root.
    Verify checksums from the fresh extraction.
    Run full from the fresh extraction with --clean-generated.
    Formulate release notes using v0.4.0 release raw MD/screenshot/link as continuity reference.
    Preserve v0.5.0 nonclaims.
    Keep runtime profiles outside full.
    Keep carbonstack-os excluded from the runnable package.

Avoid next:

    cutting v0.5.0 before fresh extraction validation.
    running release-snapshot inside the archive source root before archiving.
    adding dev-runtime-openmls or dev-runtime-openmls-wrappers to full without a deliberate release-package validation decision.
    calling the package local-backbone.
    starting PQ/state/vault work before v0.5.0 release prep finishes.
    including carbonstack-os in the runnable package.
    making general-public usable claims.

---

## 9. Critical Paths / Functions

### Source-of-truth public surfaces

    [GITEA SOURCE OF TRUTH] https://git.bitcrusher32.win/bitcrusher32/carbonstack
    [RELEASES] https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases
    [CURRENT PUBLIC RELEASE] https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.4.0

GitHub mirrors:

    secondary push mirrors for discoverability/redundancy only.
    Do not treat mirrors as release authority unless project policy changes.
    User intends official general-public releases later on GitHub, such as v1.0.0, but current v0.5.0 is not that.

### WSL Debian active working umbrella

    [WSL-DEBIAN] umbrella: ~/repos/carbonstack_umbrella
    [WSL-DEBIAN] carbonstack: ~/repos/carbonstack_umbrella/carbonstack
    [WSL-DEBIAN] carbonstack-comms: ~/repos/carbonstack_umbrella/carbonstack-comms
    [WSL-DEBIAN] carbonstack-cypher: ~/repos/carbonstack_umbrella/carbonstack-cypher
    [WSL-DEBIAN] carbonstack-os: ~/repos/carbonstack_umbrella/carbonstack-os
    [GO-RUNNER] runner dir: ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate

### v0.5.0 package staging surfaces

    [STAGING SCRIPT] carbonstack/scripts/stage-v0.5.0-package.sh
    [STAGING DOC] carbonstack/docs/166-v0.5.0-package-rehearsal-plan-v0.md
    [REGISTRY ENTRY] carbonstack.script.stage-v0.5.0-package in carbonstack/registry/commands.v0.yaml
    [DEFAULT STAGE ROOT] /tmp/carbonstack-v0.5.0-rehearsal/stage
    [DEFAULT PACKAGE ROOT] /tmp/carbonstack-v0.5.0-rehearsal/stage/package
    [DEFAULT RELEASE DIR] /tmp/carbonstack-v0.5.0-rehearsal/stage/package/release

### Current live public docs surfaces

    [DOC] top README:
    carbonstack/README.md

    [DOC] docs index:
    carbonstack/docs/README.md

    [DOC] roadmap:
    carbonstack/roadmap/ROADMAP.md

    [RUNNER] runner README:
    carbonstack/tools/carbonstack-validate/README.md

    [REGISTRY] registry README:
    carbonstack/registry/README.md

### Registry surfaces

    [REGISTRY] root:
    carbonstack/registry

    [REGISTRY] provisional command registry:
    carbonstack/registry/commands.v0.yaml

    [REGISTRY] registry README:
    carbonstack/registry/README.md

    [REGISTRY-TEST] command registry validation:
    carbonstack/tools/carbonstack-validate/command_registry_test.go

### Runner surfaces

    [RUNNER] root:
    carbonstack/tools/carbonstack-validate

    [RUNNER] dispatch:
    carbonstack/tools/carbonstack-validate/main.go

    [RUNNER] local-cypher:
    carbonstack/tools/carbonstack-validate/local_cypher.go

    [RUNNER] release-snapshot:
    carbonstack/tools/carbonstack-validate/release_snapshot.go

    [RUNNER] checksums:
    carbonstack/tools/carbonstack-validate/checksums.go

### Comms runtime surfaces

    [COMMS] direct-sidecar dev runtime smoke proof:
    carbonstack-comms/scripts/dev-openmls-runtime-smoke.sh

    [COMMS] wrapper-based dev runtime smoke proof:
    carbonstack-comms/scripts/dev-openmls-runtime-smoke-wrappers.sh

### Known-good commands

    [LIVE VALIDATION]
    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go test ./... -count=1
    go run . --profile dev-runtime-openmls-wrappers --clean-generated
    go run . --profile dev-runtime-openmls --clean-generated
    go run . --profile local-cypher
    go run . --profile doctor
    go run . --profile core --clean-generated

    [PACKAGE STAGING]
    cd ~/repos/carbonstack_umbrella/carbonstack
    scripts/stage-v0.5.0-package.sh

    [NEXT PACKAGE VALIDATION SHAPE]
    cd /tmp/carbonstack-v0.5.0-rehearsal/stage/package/carbonstack/tools/carbonstack-validate
    go run . --profile write-checksums --root /tmp/carbonstack-v0.5.0-rehearsal/stage/package
    go run . --profile verify-checksums --root /tmp/carbonstack-v0.5.0-rehearsal/stage/package

    [FRESH EXTRACTION VALIDATION SHAPE]
    cd <fresh-extraction>/package/carbonstack/tools/carbonstack-validate
    go run . --profile verify-checksums --root <fresh-extraction>/package
    go run . --profile full --root <fresh-extraction>/package --clean-generated

---

## 10. v0.4.x Forward Plan

### Immediate next: v0.4.22 checksum + fresh extraction validation + release notes formulation

Expected:

    Stage package skeleton from clean tree.
    Write checksums.
    Verify checksums in staged package.
    Archive package source root.
    Extract archive into fresh throwaway validation root.
    Verify checksums from fresh extraction.
    Run full from fresh extraction with --clean-generated.
    Formulate release notes using v0.4.0 release screenshot/raw MD/link as continuity reference.
    Do not cut release yet unless explicitly decided.
    Do not generate final PRIME/asset export until v0.4.23.

### v0.4.23

Expected:

    final LogDoc sanitization/export.
    release asset generation.
    final pre-release state check.
    prepare v0.4.xPRIME compression into LogDoc v0.5.0.

### v0.5.0

Expected:

    minor epoch release.
    v0.4.xPRIME compression into LogDoc v0.5.0.
    Gitea official pre-release.
    GitHub remains secondary mirror.
    Not general-public usable.
    Not v1.0.0.

---

## 11. v0.5.x and Later Carry-Forward

### v0.5.0 minor epoch release

After v0.4.22/v0.4.23 release-prep, cut a v0.5.0 minor epoch release using the accumulated v0.4.x work.

Dedicated remaining rungs should include:

    checksum generation.
    fresh extraction validation.
    release notes using v0.4.0 continuity style.
    final LogDoc sanitization/export.
    release asset generation.
    release cut.
    v0.4.xPRIME compression into LogDoc v0.5.0.

### v0.5.x implementation after release

Local storage, trust state, provider state, vault design, and PQ/hybrid ciphersuite migration readiness.

Do not start PQ/hybrid implementation before the runtime OpenMLS/Cypher path exists and is documented. PQ/hybrid migration touches provider state, ciphersuite metadata, trust display, artifact size, reinit/rekey/recovery policy, and claim boundaries.

### v0.6.x+

Hostile-server and abuse-resistance harnesses.

### v0.7.x+

Deployability and operations hardening.

### v0.8.x+

Documented self-pentest / adversarial validation.

### v0.9.x+

Claim-boundary review.

### v0.10.x+

Android backend/app work.

### v1.x.x

Public app/server major epoch and much later CarbonStackOS/appliance research, if justified.

---

## 12. Lean Breakpoint Summary

    CarbonStack v0.4.21 is the v0.5.0 package rehearsal plan and staging implementation checkpoint. v0.4.0 remains the current public broad local deployability pre-release. carbonstack is now at 2c2554d chore: stage v0.5.0 package rehearsal. carbonstack-comms remains at cb4e59d test: add wrapper-based OpenMLS runtime smoke proof. v0.4.21 adds scripts/stage-v0.5.0-package.sh, docs/166-v0.5.0-package-rehearsal-plan-v0.md, a command registry entry for carbonstack.script.stage-v0.5.0-package, docs index update, and roadmap update for the compressed late-v0.4.x release-prep runway. The staging helper uses clean tracked repo contents to stage carbonstack, carbonstack-comms, carbonstack-cypher, and release metadata skeleton while excluding carbonstack-os and forbidden generated/private/build artifacts. The first pre-commit staging run correctly failed because carbonstack had unstaged changes. After commit/push, staging succeeded: PASS: v0.5.0 package skeleton staged; release metadata existed; release/checksums.txt was intentionally absent; carbonstack-os was excluded; final repo heads were clean across all four repos. Next safe action: v0.4.22 checksum generation, fresh extraction validation, and release notes formulation using the v0.4.0 release as continuity reference.

---

## 13. Preserved Immediate Previous Handoff: v0.4.20

The following is the previous v0.4.20 handoff. Where it conflicts with the v0.4.21 overlay above, v0.4.21 wins for current state.



# CarbonStack LogDoc v0.4.20

**Last updated:** 2026-06-05 local session  
**Owner/operator:** bitcrusher32  
**Assistant/model:** GPT-5.5 Thinking  
**Current phase:** Phase 2F / **v0.4.x roadmap refresh + release-surface current-state cleanup checkpoint complete**. The v0.4.0 public broad local deployability pre-release remains the current public release. v0.4.18 added the first provisional cross-repo known-good command registry; v0.4.19 added Go-test validation for that registry and recorded the local help/manual boundary plan; v0.4.20 refreshed the long-term roadmap and patched live public/release-facing docs so they no longer present v0.3.20/v0.3.32 as current. `carbonstack` is now at `7d4fc0d (HEAD -> main, origin/main, origin/HEAD) docs: refresh release-surface current state`; `carbonstack-comms` remains at `cb4e59d (HEAD -> main, origin/main, origin/HEAD) test: add wrapper-based OpenMLS runtime smoke proof`; `carbonstack-cypher` remains at `9ab994c (HEAD -> main, origin/main, origin/HEAD) docs: point to local operator runbook`; `carbonstack-os` remains at `1bbbe52 (HEAD -> main, origin/main, origin/HEAD) docs: clarify CarbonStackOS target direction`.

**Version schema:** `v[scope].[timeline]`. This file is `v0.4.20`, the roadmap-refresh and release-surface cleanup checkpoint after the v0.4.19 command-registry validation checkpoint. Per LogDoc V2 practice, this Markdown file preserves current operational continuity and carries the v0.4.x runtime-integration / pre-v0.5.0 release-prep thread forward; the JSON breakpoint remains the lean machine-readable handoff state.

**Compression note:** The detailed v0.3.0 through v0.3.36PRIME process ledger remains archived in PRIME/sanitized PRIME. v0.4.0 compressed that pre-release sprint; v0.4.1 through v0.4.12 preserved runtime recon, command contracts, dev OpenMLS send/inbox commands, smoke proof, runner profile, boundary checks, wrapper recon, and bootstrap command contract. v0.4.13 through v0.4.15 preserved staged bootstrap wrapper implementation. v0.4.16 preserved the first committed wrapper-based runtime smoke variant. v0.4.17 preserved the separate manual runner profile for that wrapper smoke. v0.4.18 preserved the first provisional cross-repo command registry. v0.4.19 preserved registry validation/test coverage and local-help/manual planning. v0.4.20 now preserves roadmap refresh and release-surface current-state cleanup.

---

## 1. Project Goal

**Active goal:** Transition from late-v0.4.x command-surface maturity into v0.5.0 minor-epoch release preparation, while keeping the public surface current, claim-bounded, and release-style-compatible with v0.4.0. v0.4.20 refreshes the roadmap to the post-v0.4.19 state and updates live public/release-facing docs so current status points to v0.4.0 as the public release and v0.4.20 as the current mainline checkpoint.

CarbonStack’s broader project intent remains:

- **CarbonStack** is the secure-communications appliance-stack project and shared doctrine/release authority.
- **CarbonStackComms** is the text-first messaging client and current OpenMLS sidecar/relay integration surface.
- **CarbonStackCypher** is the self-hostable relay/storage stack.
- **CarbonStackOS** is future constrained Android-derived appliance OS work and remains out of current runnable validation packages.
- The `carbonstack` Gitea repo remains the source of truth for release state, release assets, docs, registry, and public claims.
- GitHub mirrors exist as secondary push mirrors only; Gitea remains authoritative.

**North star:**

> Its goal is not to be a better smartphone. Its goal is to stop being a smartphone.

**Core doctrine:**

> Every feature is guilty until it proves it does not add unacceptable parser, network, sensor, identity, or filesystem authority.

**Current reality after v0.4.20:**

    carbonstack/docs/README.md now identifies:
      current public release: v0.4.0 broad local deployability pre-release.
      current mainline checkpoint: v0.4.20 roadmap refresh after v0.4.19 registry validation/local-help planning.
      current mainline validation surface including runner Go tests, both runtime profiles, local-cypher, doctor, and core.
      current v0.5.0 direction as a minor epoch release, not PQ/state/vault implementation.

    carbonstack/roadmap/ROADMAP.md now identifies:
      current state after v0.4.20.
      Gitea as authoritative for releases/assets/current project state.
      GitHub as push mirrors only.
      current validated surfaces and runner split.
      command registry and registry validation as current command-surface hygiene.
      v0.5.0 as a minor epoch release runway after late-v0.4.x cleanup.
      v0.5.0 release-prep rungs before major v0.5.x state/trust/vault/PQ implementation.

    carbonstack/tools/carbonstack-validate/README.md now identifies:
      phase: v0.4.x validation runner / v0.5.0 release-prep surface.
      Debian / WSL Debian as the preferred current runner-backed validation target.
      updated WSL Debian example using the current live umbrella path and go test / doctor / core / local-cypher flow.
      removal of stale Windows-specific `C:\▓▓` example text from the current runner README.

    Numbered docs under carbonstack/docs remain mostly historical archive records.
    They were intentionally not mass-rewritten.
    docs/README.md remains the live archive index and current-vs-historical boundary surface.

    Existing runtime proof split remains:
      dev-runtime-openmls:
        direct sidecar bootstrap -> openmls-send-dev -> Cypher -> openmls-inbox-dev --ack
      dev-runtime-openmls-wrappers:
        openmls-*-dev bootstrap wrappers -> openmls-send-dev -> Cypher -> openmls-inbox-dev --ack

    dev-runtime-openmls remains the direct-smoke baseline.
    dev-runtime-openmls-wrappers remains the wrapper-smoke maturity surface.
    Both runtime profiles remain manual-only and live-umbrella-only.
    Neither runtime profile is included in full.
    full remains release-snapshot + local-cypher.
    Existing send/inbox remain stub-era and are explicitly included in the registry as legacy/stub-era surfaces.
    local-backbone remains reserved.

Hard nonclaims remain: v0.4.20 does **not** prove production readiness, production E2EE, hostile-server safety, metadata privacy, Android readiness, CarbonStackOS readiness, secure vault/storage, public ingress safety, real homelab validation, audit, certification, PQ/hybrid security, mature messenger UX, release-package runtime validation, Relay Space join/onboarding, `full` runtime OpenMLS validation, or local-backbone readiness. The roadmap/docs cleanup is release-surface hygiene, not a runtime or security capability.

---

## 2. Current Public Release State

Public release:

    https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.4.0

Release title:

    CarbonStack v0.4.0 Broad Local Deployability Pre-Release

Release tag:

    v0.4.0

Release commit:

    24ac9fc494 / 24ac9fc docs: record v0.4.0 package rehearsal

Gitea release status:

    Pre-Release

Important release-page warning remains:

    The default Gitea Source Code ZIP/TAR.GZ downloads are auto-generated carbonstack-only archives.
    The intended runnable validation artifact is the attached multi-repo release package plus companion assets.

Post-release public surface:

    carbonstack README is evergreen.
    It points interested parties to the Releases page for runnable / known-good artifacts.
    Gitea is source of truth.
    GitHub mirrors are secondary push mirrors, not release authority.

Official general-public usable releases:

    Not yet.
    User explicitly intends official usable/general-public releases to happen on GitHub later, such as the v1.0.0 major epoch release.
    Current v0.5.0 target is a Gitea-source-of-truth pre-release/minor epoch release, not a general-public official usable release.

---

## 3. Current Repo Heads

    carbonstack        7d4fc0d (HEAD -> main, origin/main, origin/HEAD) docs: refresh release-surface current state
    carbonstack-comms  cb4e59d (HEAD -> main, origin/main, origin/HEAD) test: add wrapper-based OpenMLS runtime smoke proof
    carbonstack-cypher 9ab994c (HEAD -> main, origin/main, origin/HEAD) docs: point to local operator runbook
    carbonstack-os     1bbbe52 (HEAD -> main, origin/main, origin/HEAD) docs: clarify CarbonStackOS target direction

Current public release tag:

    v0.4.0 -> 24ac9fc494 / 24ac9fc docs: record v0.4.0 package rehearsal

Recent runtime / maturity commits:

    v0.4.13 identity wrapper implementation:
      20c4cc9 feat: add dev OpenMLS identity bootstrap commands

    v0.4.14 bundle/conversation wrapper implementation:
      1f5c09e feat: add dev OpenMLS bundle and conversation bootstrap commands

    v0.4.15 add-member/join wrapper implementation:
      c2d3b16 feat: add dev OpenMLS add-member and join bootstrap commands

    v0.4.16 wrapper-based smoke proof:
      cb4e59d test: add wrapper-based OpenMLS runtime smoke proof

    v0.4.17 wrapper-smoke runner/profile alignment:
      a501442 feat: add wrapper OpenMLS runtime validation profile

    v0.4.18 command registry:
      a4162d7 docs: add known-good command registry

    v0.4.19 registry validation:
      568fb45 test: validate command registry coverage

    Current v0.4.20 release-surface cleanup:
      7d4fc0d docs: refresh release-surface current state

Continuity note:

    `7d4fc0d` is a carbonstack docs/roadmap/runner-README commit after the v0.4.0 public release tag.
    It does not retag v0.4.0.
    The release remains attached to `24ac9fc494`.

---

## 4. Current Validated / Known State

Validated / known at the v0.4.20 checkpoint:

    [RELEASE] v0.4.0 Gitea release exists, is marked Pre-Release, and is attached to 24ac9fc494.
    [DOCS] v0.4.20 roadmap PDF refresh was generated in-session as CarbonStack_Long_Term_Roadmap_v0.4.20_REFRESH.pdf.
    [DOCS] carbonstack/docs/README.md current-state block now points to v0.4.0 public release and v0.4.20 mainline checkpoint.
    [DOCS] carbonstack/roadmap/ROADMAP.md current-state block now points to v0.4.20 and includes a late-v0.4.x -> v0.5.0 minor epoch release runway.
    [DOCS] carbonstack/tools/carbonstack-validate/README.md phase/examples were refreshed for v0.4.x validation runner / v0.5.0 release-prep surface and WSL Debian current path.
    [DOCS] numbered docs under carbonstack/docs were intentionally not mass-rewritten because they are a historical archive.
    [RUNNER] full means release-snapshot followed by local-cypher.
    [RUNNER] dev-runtime-openmls validates the original direct-sidecar dev/pre-alpha OpenMLS application-message runtime CLI smoke proof.
    [RUNNER] dev-runtime-openmls-wrappers validates the wrapper-based dev/pre-alpha OpenMLS application-message runtime CLI smoke proof.
    [RUNNER] both runtime profiles remain manual-only and live-umbrella-only.
    [RUNNER] neither runtime profile is included in full.
    [REGISTRY] registry/commands.v0.yaml exists and remains provisional.
    [REGISTRY] command_registry_test.go validates the registry under Go test.
    [MIRRORS] all four repos have GitHub push-mirror functionality, with Gitea remaining source of truth.

v0.4.20 validation observed from the log:

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go test ./... -count=1 passed, including command_registry_test.go.
    go run . --profile dev-runtime-openmls-wrappers --clean-generated passed.
    go run . --profile dev-runtime-openmls --clean-generated passed.
    go run . --profile local-cypher passed.
    go run . --profile doctor passed.
    go run . --profile core --clean-generated passed.

    cd ~/repos/carbonstack_umbrella/carbonstack-comms
    go test ./internal/app -count=1 passed.
    go test ./... -count=1 -timeout 600s passed.

Final commit/push:

    carbonstack:
    7d4fc0d docs: refresh release-surface current state

Final repo snapshot:

    carbonstack        7d4fc0d docs: refresh release-surface current state
    carbonstack-comms  cb4e59d test: add wrapper-based OpenMLS runtime smoke proof
    carbonstack-cypher 9ab994c docs: point to local operator runbook
    carbonstack-os     1bbbe52 docs: clarify CarbonStackOS target direction

Expected live umbrella validation commands now include:

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go test ./... -count=1
    go run . --profile dev-runtime-openmls-wrappers --clean-generated
    go run . --profile dev-runtime-openmls --clean-generated
    go run . --profile local-cypher
    go run . --profile doctor
    go run . --profile core --clean-generated

Expected v0.4.0 public package validation commands remain unchanged:

    cd <package-root>/carbonstack/tools/carbonstack-validate
    go run . --profile verify-checksums --root <package-root>
    go run . --profile full --root <package-root> --clean-generated

---

## 5. v0.4.20 Work Completed

### 5.1 Long-term roadmap refresh was generated

Roadmap artifact generated in-session:

    CarbonStack_Long_Term_Roadmap_v0.4.20_REFRESH.pdf

Purpose:

    Refresh the v0.4.14 roadmap to the post-v0.4.19 state.
    Preserve the same general roadmap formatting.
    Record v0.4.15 through v0.4.19 as completed work.
    Carry forward late-v0.4.x cleanup/maturity rungs.
    Add v0.5.0 as the next minor epoch release target after v0.4.x cleanup/release-hardening.
    Preserve v0.5.x state/trust/vault/PQ as post-release implementation direction.
    Preserve PQ/hybrid caution and nonclaims.
    Preserve local-backbone reserve.
    Preserve direct vs wrapper runtime profile split.
    Preserve command registry as source-of-truth metadata.

Release-framing decision:

    v0.5.0 should be a minor epoch release similar in role/style to v0.4.0.
    Use the v0.4.0 Gitea release style as continuity reference.
    v0.5.0 is not the start of PQ/state/vault implementation.
    v0.5.0 is not a general-public official usable GitHub release.
    Official general-public releases belong later, such as the intended v1.0.0 major epoch line.

### 5.2 Release-surface / stale-claim scout identified targeted current-surface drift

The v0.4.21-style scout found that most numbered docs are intentionally stale historical archive records and should not be mass-rewritten.

Targeted live/current-surface drift:

    docs/README.md:
      Still opened with stale v0.3.20 public testing release and v0.3.32 current mainline validation state.
      Needed to identify v0.4.0 as current public release and v0.4.20 as current mainline checkpoint.

    roadmap/ROADMAP.md:
      Still opened with stale v0.3.32 current-state framing.
      Needed to identify current state after v0.4.20, current validated surfaces, current runner split, command registry status, nonclaims, and the late-v0.4.x -> v0.5.0 minor epoch release runway.

    tools/carbonstack-validate/README.md:
      Still said `Phase: v0.3.12 runner hardening / docs integration`.
      Still contained a Windows example with a user-specific `C:\▓▓\...` path.
      Needed to refresh phase and examples toward v0.4.x validation runner / v0.5.0 release-prep surface and WSL Debian as preferred current validation target.

### 5.3 docs/README.md current-state block was patched

Updated:

    carbonstack/docs/README.md

New current-state framing:

    current public release: v0.4.0 broad local deployability pre-release.
    current mainline checkpoint: v0.4.20 roadmap refresh after v0.4.19 registry validation and local-help planning.
    current mainline validation surface:
      go test ./... -count=1
      dev-runtime-openmls-wrappers --clean-generated
      dev-runtime-openmls --clean-generated
      local-cypher
      doctor
      core --clean-generated
    current release-package validation shape remains release-specific.
    v0.4.0 public release package should use release-attached runbook/assets.
    v0.5.0 direction is a minor epoch release of accumulated v0.4.x runtime/runner/wrapper/registry/validation work.

Preserved:

    historical archive note.
    current-vs-historical rule.
    numbered docs index.
    security and maturity warning.

### 5.4 roadmap/ROADMAP.md current-state and release-runway block was patched

Updated:

    carbonstack/roadmap/ROADMAP.md

New current-state framing:

    current state after v0.4.20.
    v0.4.0 remains current public release artifact.
    mainline has moved through runtime OpenMLS integration, wrapper-based smoke validation, runner profile separation, command registry, and registry validation.
    Gitea is authoritative for releases/assets/current project state.
    GitHub mirrors are push mirrors only.
    general-public usable official releases belong later, such as the intended v1.0.0 major epoch line.

New current validated surfaces:

    v0.4.0 release-package validation shape:
      verify-checksums
      full --clean-generated

    live mainline validation shape:
      go test ./... -count=1
      dev-runtime-openmls-wrappers --clean-generated
      dev-runtime-openmls --clean-generated
      local-cypher
      doctor
      core --clean-generated

New runner split:

    full:
      release-package validation ladder
      release-snapshot followed by local-cypher
      not deployment
      not local-backbone
      not runtime Comms UX

    dev-runtime-openmls:
      manual live-umbrella direct-sidecar OpenMLS runtime smoke profile
      openmls-send-dev -> Cypher -> openmls-inbox-dev --ack
      not included in full

    dev-runtime-openmls-wrappers:
      manual live-umbrella wrapper-bootstrap OpenMLS runtime smoke profile
      openmls-*-dev bootstrap wrappers -> openmls-send-dev -> Cypher -> openmls-inbox-dev --ack
      separate maturity surface
      does not replace dev-runtime-openmls yet
      not included in full

New v0.5.0 minor epoch release runway:

    v0.4.21 release-surface recon and stale-claim scan
    v0.4.22 public-surface cleanup patch
    v0.4.23 v0.5.0 package rehearsal plan
    v0.4.24 v0.5.0 package rehearsal implementation
    v0.4.25 fresh extraction validation and asset/checksum rehearsal
    v0.4.26 release notes, LogDoc, and final pre-release checkpoint
    v0.5.0 minor epoch release

### 5.5 tools/carbonstack-validate/README.md phase/examples were patched

Updated:

    carbonstack/tools/carbonstack-validate/README.md

Changes:

    `Phase: v0.3.12 runner hardening / docs integration`
    became:
    `Phase: v0.4.x validation runner / v0.5.0 release-prep surface`

    Windows example was changed into a Windows note:
      Windows validation is not the current mainline release-prep target.
      Prefer Debian / WSL Debian unless a release runbook says otherwise.

    WSL Debian example was updated to:
      ~/.cargo/env
      ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
      go test ./... -count=1
      doctor
      core --clean-generated
      local-cypher

### 5.6 Validation passed

Observed passing:

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go test ./... -count=1
    go run . --profile dev-runtime-openmls-wrappers --clean-generated
    go run . --profile dev-runtime-openmls --clean-generated
    go run . --profile local-cypher
    go run . --profile doctor
    go run . --profile core --clean-generated

    cd ~/repos/carbonstack_umbrella/carbonstack-comms
    go test ./internal/app -count=1
    go test ./... -count=1 -timeout 600s

Final repo snapshot was clean across all four repos.

---

## 6. Critical Blunders / Lessons to Carry Forward

### 6.1 Copy-box/pasteback fragility persisted

The v0.4.21 scout paste had copy-box breakage and tail-glitched command fragments around the runner section. The safe approach was to use the clearly landed public-surface excerpts and final validation/clean snapshot, then provide smaller patch blocks.

Lesson:

    Continue splitting commands into smaller blocks.
    Avoid overly massive single paste commands when possible.
    Treat final commit/push plus final four-repo clean snapshot as breakpoint authority.

### 6.2 Historical numbered docs should not be “fixed” just because they are stale

The scout found many stale-looking strings in numbered docs, but the project intentionally uses numbered docs as a chronological archive.

Lesson:

    Do not mass-rewrite older numbered docs.
    Patch only live/current surfaces:
      top README
      docs/README.md
      roadmap/ROADMAP.md
      runner README
      release runbooks
      registry README
      current release notes
    Keep historical docs as provenance unless they are actively linked as current behavior.

### 6.3 The docs index had become a current-state hazard

docs/README.md correctly explained historical/archive semantics, but its top current-state block still pointed to v0.3.20/v0.3.32.

Lesson:

    The docs index is not merely historical; it is a live navigation surface.
    Its current-state header must be kept up to date after milestone rungs.

### 6.4 The roadmap had become a current-state hazard

roadmap/ROADMAP.md still opened from a v0.3.32 worldview even though later sections recorded v0.4.17-v0.4.19.

Lesson:

    Roadmap front matter must be refreshed after grouped goalsets.
    It should not leave old current-state language at the top while newer sections exist below.

### 6.5 Runner README had stale phase and PII-ish local path exposure

The runner README still had a v0.3.12 phase line and a Windows example path containing `C:\▓▓`.

Lesson:

    Current live READMEs should avoid user-specific local paths.
    Use generic or current preferred WSL Debian examples.
    Windows can remain mentioned as secondary/best-effort, but current release-prep validation is Debian / WSL Debian first.

### 6.6 Commit metadata warning remains a recurring polish issue

The commit log again showed auto-generated local commit metadata:

    Committer: bitcrusher32 <▓▓>

This is not a functional validation failure, but remains a professionalism/publishing hygiene issue to address eventually.

---

## 7. Current Blockers / Not Validated

Current blockers and nonvalidated areas after v0.4.20:

    Existing send/inbox remain stub-era.
    Runtime Comms has dev-only OpenMLS commands, direct smoke proof, wrapper-based smoke proof, and separate manual runner profiles, but no mature user-facing UX.
    Neither runtime profile is included in full.
    Neither runtime profile is part of public release package validation.
    Fresh release-package behavior for dev-runtime-openmls or dev-runtime-openmls-wrappers has not been tested as a supported validation surface.
    No Relay Space join model exists yet.
    local-backbone runner profile does not exist and remains reserved.
    Registry validation is now a Go test, but not a runner profile.
    Generated command reference from the registry does not exist yet.
    Local help/manual generation from the registry does not exist yet.
    Short explainer strings are not yet integrated directly beside command registration in Comms.
    README/dev command sprawl is mitigated by registry + test, but not fully solved.
    v0.5.0 minor epoch release ladder is planned but not implemented.
    v0.5.0 package rehearsal has not started.
    v0.5.0 release assets/checksums/runbook/notes have not been staged.
    Post-release public download verification is still useful if not already performed.
    Sanitized PRIME placement in carbonstack/logdoc-loop-system still needs a policy decision.
    Production-safe provider storage and secure vault design remain v0.5.x future work.
    PQ/hybrid ciphersuite migration readiness remains v0.5.x future work.
    Hostile-server rollback/replay/metadata-abuse validation remains future work.
    Android/CarbonStackOS implementation remains later work.
    CarbonStack Relay Space mechanics remain a future concept, not an implementation.

Hard nonclaims remain:

    Do not claim production readiness.
    Do not claim production E2EE.
    Do not claim hostile-server safety.
    Do not claim metadata privacy.
    Do not claim Debian production deployability.
    Do not claim systemd readiness.
    Do not claim cloudflared readiness.
    Do not claim Android readiness.
    Do not claim CarbonStackOS readiness.
    Do not claim secure vault/storage.
    Do not claim real homelab validation.
    Do not claim external audit or certification.
    Do not claim CarbonStack Relay Space mechanics are implemented.
    Do not claim local-backbone exists or is justified yet.
    Do not claim local-cypher validates runtime Comms UX.
    Do not claim full validates local-backbone.
    Do not claim full is a deployment command.
    Do not claim dev-runtime-openmls is release-package validation yet.
    Do not claim dev-runtime-openmls-wrappers is release-package validation yet.
    Do not claim send/inbox are OpenMLS-backed yet.
    Do not claim PQ/hybrid security or quantum-safe messaging.
    Do not let GitHub mirrors supersede Gitea source-of-truth status.

Allowed claim:

    CarbonStack current public release remains v0.4.0 broad local deployability pre-release.
    CarbonStack mainline is now at v0.4.20 roadmap refresh / release-surface cleanup.
    The live docs index and roadmap now reflect v0.4.20 current state.
    The validation runner README no longer carries the stale v0.3.12 phase or user-specific Windows path example.
    CarbonStack has a provisional cross-repo known-good command registry at `registry/commands.v0.yaml`.
    The registry is validated by `tools/carbonstack-validate/command_registry_test.go` when run from a full umbrella/release-package-like layout.
    CarbonStackComms still has dev-only openmls-send-dev and openmls-inbox-dev commands.
    CarbonStackComms still has direct-sidecar and wrapper-based dev runtime smoke scripts.
    carbonstack/tools/carbonstack-validate still has separate manual dev-runtime-openmls and dev-runtime-openmls-wrappers profiles.
    Both runtime profiles are live-git-umbrella-only and not included in full.
    Existing send/inbox remain stub-era.
    v0.5.0 is now the planned minor epoch release target after late-v0.4.x cleanup/release-hardening.
    Gitea remains source of truth; GitHub mirrors are secondary push mirrors.

---

## 8. Next Safest Actions

Recommended next rung:

    v0.4.21 / v0.4.22 package-rehearsal planning for the v0.5.0 minor epoch release.

Focus:

    Define the exact v0.5.0 package rehearsal plan.
    Decide which repos are included in the v0.5.0 runnable validation package.
    Decide whether carbonstack-os remains excluded from runnable package assets.
    Decide release metadata names and paths.
    Decide whether registry files and generated roadmap PDF are included in the package.
    Decide whether current full still remains verify-checksums + full only, or whether an additional live-umbrella-only runtime profile is documented separately.
    Decide whether dev-runtime-openmls / dev-runtime-openmls-wrappers stay outside release-package validation for v0.5.0.
    Draft the validation freeze expectations.
    Preserve v0.4.0 release style: release notes, attached package, checksums, manifest, validation-freeze, testing runbook, explicit nonclaims.

Likely candidates after planning:

    v0.5.0 package rehearsal implementation.
    fresh extraction validation.
    asset/checksum rehearsal.
    release notes drafting using the v0.4.0/v0.3.20 continuity style.
    LogDoc/breakpoint export.
    final release cut.

Avoid next:

    replacing the direct smoke script.
    merging dev-runtime-openmls and dev-runtime-openmls-wrappers.
    adding either runtime profile to full without a deliberate release-package validation decision.
    calling either runtime profile local-backbone.
    treating wrapper bootstrap as Relay Space join UX.
    moving into PQ/vault work before v0.5.0 release preparation.
    public ingress.
    cloudflared.
    systemd.
    real homelab deployment.
    Android app.
    CarbonStackOS implementation.
    production/security claims.

---

## 9. Critical Paths / Functions

### Source-of-truth public surfaces

    [GITEA SOURCE OF TRUTH] https://git.bitcrusher32.win/bitcrusher32/carbonstack
    [RELEASES] https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases
    [CURRENT PUBLIC RELEASE] https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.4.0

GitHub mirrors:

    secondary push mirrors for discoverability/redundancy only.
    Do not treat mirrors as release authority unless project policy changes.
    User intends official general-public releases later on GitHub, such as v1.0.0, but current v0.5.0 is not that.

### WSL Debian active working umbrella

    [WSL-DEBIAN] umbrella: ~/repos/carbonstack_umbrella
    [WSL-DEBIAN] carbonstack: ~/repos/carbonstack_umbrella/carbonstack
    [WSL-DEBIAN] carbonstack-comms: ~/repos/carbonstack_umbrella/carbonstack-comms
    [WSL-DEBIAN] carbonstack-cypher: ~/repos/carbonstack_umbrella/carbonstack-cypher
    [WSL-DEBIAN] carbonstack-os: ~/repos/carbonstack_umbrella/carbonstack-os
    [GO-RUNNER] runner dir: ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate

### Current live public docs surfaces

    [DOC] top README:
    carbonstack/README.md

    [DOC] docs index:
    carbonstack/docs/README.md

    [DOC] roadmap:
    carbonstack/roadmap/ROADMAP.md

    [RUNNER] runner README:
    carbonstack/tools/carbonstack-validate/README.md

    [REGISTRY] registry README:
    carbonstack/registry/README.md

### Registry surfaces

    [REGISTRY] root:
    carbonstack/registry

    [REGISTRY] provisional command registry:
    carbonstack/registry/commands.v0.yaml

    [REGISTRY] registry README:
    carbonstack/registry/README.md

    [REGISTRY-TEST] command registry validation:
    carbonstack/tools/carbonstack-validate/command_registry_test.go

    [DOC] v0.4.18 registry checkpoint:
    carbonstack/docs/164-known-good-command-registry-recon-v0.md

    [DOC] v0.4.19 registry validation/help plan:
    carbonstack/docs/165-command-registry-validation-and-help-plan-v0.md

### Runtime integration doc surfaces

    [DOC] v0.4.1 runtime recon:
    carbonstack/docs/155-runtime-comms-openmls-cypher-recon-v0.md

    [DOC] v0.4.2 command contract:
    carbonstack/docs/156-runtime-openmls-command-contract-v0.md

    [DOC] v0.4.6 smoke proof record:
    carbonstack/docs/157-dev-runtime-openmls-smoke-proof-v0.md

    [DOC] v0.4.7 pre-local-backbone assessment:
    carbonstack/docs/158-pre-local-backbone-assessment-v0.md

    [DOC] v0.4.9 runner profile:
    carbonstack/docs/159-dev-runtime-openmls-runner-profile-v0.md

    [DOC] v0.4.10 profile boundary:
    carbonstack/docs/160-dev-runtime-openmls-profile-boundary-v0.md

    [DOC] v0.4.11 bootstrap wrapper recon:
    carbonstack/docs/161-openmls-bootstrap-wrapper-recon-v0.md

    [DOC] v0.4.12 bootstrap command contract:
    carbonstack/docs/162-openmls-bootstrap-command-contract-v0.md

    [DOC] v0.4.17 wrapper profile:
    carbonstack/docs/163-dev-runtime-openmls-wrapper-runner-profile-v0.md

Numbered docs note:

    Most numbered docs are historical archive records.
    Do not mass-rewrite old numbered docs for stale language unless they are explicitly being promoted as current behavior.

### Runner surfaces

    [RUNNER] root:
    carbonstack/tools/carbonstack-validate

    [RUNNER] dispatch:
    carbonstack/tools/carbonstack-validate/main.go

    [RUNNER] registry validation test:
    carbonstack/tools/carbonstack-validate/command_registry_test.go

    [RUNNER] dev-runtime-openmls implementation:
    carbonstack/tools/carbonstack-validate/dev_runtime_openmls.go

    [RUNNER] dev-runtime-openmls-wrappers implementation:
    carbonstack/tools/carbonstack-validate/dev_runtime_openmls_wrappers.go

    [RUNNER] runner README:
    carbonstack/tools/carbonstack-validate/README.md

    [RUNNER] local-cypher:
    carbonstack/tools/carbonstack-validate/local_cypher.go

    [RUNNER] release-snapshot:
    carbonstack/tools/carbonstack-validate/release_snapshot.go

    [RUNNER] checksums:
    carbonstack/tools/carbonstack-validate/checksums.go

### Comms runtime surfaces

    [COMMS] CLI entry:
    carbonstack-comms/cmd/comms/main.go

    [COMMS] command registry and non-OpenMLS commands:
    carbonstack-comms/internal/app/commands.go

    [COMMS] OpenMLS message runtime command helpers:
    carbonstack-comms/internal/app/openmls_runtime.go

    [COMMS] OpenMLS bootstrap wrapper helpers:
    carbonstack-comms/internal/app/openmls_bootstrap.go

    [COMMS] OpenMLS bootstrap wrapper tests:
    carbonstack-comms/internal/app/openmls_bootstrap_dev_test.go

    [COMMS] openmls-send-dev tests:
    carbonstack-comms/internal/app/openmls_send_dev_test.go

    [COMMS] openmls-inbox-dev tests:
    carbonstack-comms/internal/app/openmls_inbox_dev_test.go

    [COMMS] direct-sidecar dev runtime smoke proof:
    carbonstack-comms/scripts/dev-openmls-runtime-smoke.sh

    [COMMS] wrapper-based dev runtime smoke proof:
    carbonstack-comms/scripts/dev-openmls-runtime-smoke-wrappers.sh

    [COMMS] OpenMLS sidecar dir:
    carbonstack-comms/internal/protocol/mls/openmls-sidecar

    [COMMS] sidecar README:
    carbonstack-comms/internal/protocol/mls/openmls-sidecar/README.md

### v0.4.20 runner/test behavior

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go test ./... -count=1

    Includes command_registry_test.go:
      validates registry structure and coverage.
      skips without a sibling umbrella layout.
      should be treated as the current registry drift guard.

### Existing v0.4.0 package validation

    cd <package-root>/carbonstack/tools/carbonstack-validate
    go run . --profile verify-checksums --root <package-root>
    go run . --profile full --root <package-root> --clean-generated

---

## 10. v0.4.x Forward Plan

### Immediate next: v0.5.0 package rehearsal planning

Expected:

    Define the v0.5.0 package shape.
    Reuse v0.4.0 release presentation style.
    Confirm included repos and excluded repos.
    Confirm release metadata names.
    Confirm validation command ladder.
    Confirm whether registry and roadmap PDF are included in package assets.
    Confirm runtime profile treatment: live-umbrella evidence only vs release-package validation.
    Preserve nonclaims.

### v0.4.21+

Possible:

    v0.5.0 package rehearsal plan.
    v0.5.0 package rehearsal implementation.
    fresh extraction validation.
    asset/checksum rehearsal.
    release notes / validation freeze / testing runbook.
    final LogDoc and breakpoint export.
    v0.5.0 release cut.

---

## 11. v0.5.x and Later Carry-Forward

### v0.5.0 minor epoch release

After late v0.4.x cleanup/release-surface work, cut a v0.5.0 minor epoch release using the accumulated v0.4.x work.

Dedicated rungs should include:

    cleanup.
    deployability testing.
    release-style replication.
    release package rehearsal.
    release hardening.
    public wording / claim-boundary cleanup.
    final LogDoc and breakpoint export.
    release asset generation.
    release notes using the established continuity style.

### v0.5.x implementation after release

Local storage, trust state, provider state, vault design, and PQ/hybrid ciphersuite migration readiness.

Do not start PQ/hybrid implementation before the runtime OpenMLS/Cypher path exists and is documented. PQ/hybrid migration touches provider state, ciphersuite metadata, trust display, artifact size, reinit/rekey/recovery policy, and claim boundaries.

### v0.6.x+

Hostile-server and abuse-resistance harnesses.

### v0.7.x+

Deployability and operations hardening.

### v0.8.x+

Documented self-pentest / adversarial validation.

### v0.9.x+

Claim-boundary review.

### v0.10.x+

Android backend/app work.

### v1.x.x

Public app/server major epoch and much later CarbonStackOS/appliance research, if justified.

---

## 12. Lean Breakpoint Summary

    CarbonStack v0.4.20 is the roadmap-refresh and release-surface current-state cleanup checkpoint. v0.4.0 remains the current public broad local deployability pre-release. carbonstack is now at 7d4fc0d docs: refresh release-surface current state. carbonstack-comms remains at cb4e59d test: add wrapper-based OpenMLS runtime smoke proof. v0.4.20 refreshed the long-term roadmap PDF in-session and patched live current-surface docs: docs/README.md, roadmap/ROADMAP.md, and tools/carbonstack-validate/README.md. Numbered docs were intentionally not mass-rewritten because they are historical archive records. The docs index now points to v0.4.0 public release and v0.4.20 mainline checkpoint. The roadmap now opens from v0.4.20 current state and carries the late-v0.4.x -> v0.5.0 minor epoch release runway. The runner README now identifies v0.4.x validation runner / v0.5.0 release-prep surface, demotes Windows to a note, and updates WSL Debian examples. Validation passed for carbonstack tools go tests, dev-runtime-openmls-wrappers --clean-generated, dev-runtime-openmls --clean-generated, local-cypher, doctor, core --clean-generated, Comms app tests, and full Comms tests. Final repo heads were clean across all four repos. Next safe action: plan the v0.5.0 package rehearsal and release-ladder details while preserving v0.4.0-style release continuity and current nonclaims.

---

## 13. Preserved Immediate Previous Handoff: v0.4.19

The following is the previous v0.4.19 handoff. Where it conflicts with the v0.4.20 overlay above, v0.4.20 wins for current state.



# CarbonStack LogDoc v0.4.19

**Last updated:** 2026-06-05 local session  
**Owner/operator:** bitcrusher32  
**Assistant/model:** GPT-5.5 Thinking  
**Current phase:** Phase 2F / **v0.4.x command registry validation and local-help/manual planning checkpoint complete**. The v0.4.0 public broad local deployability pre-release remains the current public release. v0.4.18 added the first provisional cross-repo known-good command registry; v0.4.19 validates that registry against live command surfaces through Go-test coverage and records the local help/manual boundary plan. `carbonstack` is now at `568fb45 (HEAD -> main, origin/main, origin/HEAD) test: validate command registry coverage`; `carbonstack-comms` remains at `cb4e59d (HEAD -> main, origin/main, origin/HEAD) test: add wrapper-based OpenMLS runtime smoke proof`; `carbonstack-cypher` remains at `9ab994c (HEAD -> main, origin/main, origin/HEAD) docs: point to local operator runbook`; `carbonstack-os` remains at `1bbbe52 (HEAD -> main, origin/main, origin/HEAD) docs: clarify CarbonStackOS target direction`.

**Version schema:** `v[scope].[timeline]`. This file is `v0.4.19`, the command-registry validation and local-help planning checkpoint after the v0.4.18 known-good command registry checkpoint. Per LogDoc V2 practice, this Markdown file preserves current operational continuity and carries the v0.4.x runtime-integration / pre-v0.5.x maturity thread forward; the JSON breakpoint remains the lean machine-readable handoff state.

**Compression note:** The detailed v0.3.0 through v0.3.36PRIME process ledger remains archived in PRIME/sanitized PRIME. v0.4.0 compressed that pre-release sprint; v0.4.1 through v0.4.12 preserved runtime recon, command contracts, dev OpenMLS send/inbox commands, smoke proof, runner profile, boundary checks, wrapper recon, and bootstrap command contract. v0.4.13 through v0.4.15 preserved staged bootstrap wrapper implementation. v0.4.16 preserved the first committed wrapper-based runtime smoke variant. v0.4.17 preserved the separate manual runner profile for that wrapper smoke. v0.4.18 preserved the first provisional cross-repo command registry. v0.4.19 now preserves registry validation/test coverage and the local-help/manual planning boundary.

---

## 1. Project Goal

**Active goal:** Complete the late-v0.4.x grouped command-surface maturity goalset before refreshing the long-term roadmap and before entering the v0.5.0 minor-epoch release ladder. v0.4.19 makes the provisional command registry harder to rot by validating it through `carbonstack/tools/carbonstack-validate/command_registry_test.go`, while explicitly deferring generated command references, local embedded help, runner-profile promotion for registry validation, and v0.5.x state/trust/vault/PQ implementation.

CarbonStack’s broader project intent remains:

- **CarbonStack** is the secure-communications appliance-stack project and shared doctrine/release authority.
- **CarbonStackComms** is the text-first messaging client and current OpenMLS sidecar/relay integration surface.
- **CarbonStackCypher** is the self-hostable relay/storage stack.
- **CarbonStackOS** is future constrained Android-derived appliance OS work and remains out of current runnable validation packages.
- The `carbonstack` Gitea repo remains the source of truth for release state, release assets, docs, registry, and public claims.
- GitHub mirrors exist as secondary push mirrors only; Gitea remains authoritative.

**North star:**

> Its goal is not to be a better smartphone. Its goal is to stop being a smartphone.

**Core doctrine:**

> Every feature is guilty until it proves it does not add unacceptable parser, network, sensor, identity, or filesystem authority.

**Current reality after v0.4.19:**

    carbonstack/registry/commands.v0.yaml remains the provisional cross-repo command registry.
    registry/README.md now records validation status and the Go-test-backed check.
    docs/165-command-registry-validation-and-help-plan-v0.md records the v0.4.19 checkpoint.
    tools/carbonstack-validate/command_registry_test.go validates registry coverage and policy boundaries.
    docs/README.md indexes the v0.4.19 registry-validation/help-plan document.
    roadmap/ROADMAP.md records v0.4.19 and carries forward the v0.5.0 minor epoch release ladder.

    The registry validation test checks:
      entries exist and IDs are unique.
      required fields exist for each entry.
      source_path values resolve to existing files/directories.
      runner profile dispatch is covered.
      Comms command dispatch is covered.
      Comms scripts are covered, excluding scripts/README.md as documentation.
      sidecar command IDs are present.
      Cypher server/API IDs are present.
      legacy send/inbox/ack remain marked legacy.
      direct and wrapper OpenMLS runtime profiles remain distinct.
      local-backbone and Gitea source-of-truth boundaries remain present.
      include_in_front_readme count remains small enough to avoid README sprawl.

    Local-help/manual planning decision:
      B first: registry validation test/check.
      C later: generated Markdown command reference from registry.
      D later: Comms embedded help/manual strings only after command structure is ready.
      E later: registry-validation runner profile only if the check proves stable and useful.

    Existing runtime proof split remains:
      dev-runtime-openmls:
        direct sidecar bootstrap -> openmls-send-dev -> Cypher -> openmls-inbox-dev --ack
      dev-runtime-openmls-wrappers:
        openmls-*-dev bootstrap wrappers -> openmls-send-dev -> Cypher -> openmls-inbox-dev --ack

    dev-runtime-openmls remains the direct-smoke baseline.
    dev-runtime-openmls-wrappers remains the wrapper-smoke maturity surface.
    Both runtime profiles remain manual-only and live-umbrella-only.
    Neither runtime profile is included in full.
    full remains release-snapshot + local-cypher.
    Existing send/inbox remain stub-era and are explicitly included in the registry as legacy/stub-era surfaces.
    local-backbone remains reserved.

Hard nonclaims remain: v0.4.19 does **not** prove production readiness, production E2EE, hostile-server safety, metadata privacy, Android readiness, CarbonStackOS readiness, secure vault/storage, public ingress safety, real homelab validation, audit, certification, PQ/hybrid security, mature messenger UX, release-package runtime validation, Relay Space join/onboarding, `full` runtime OpenMLS validation, or local-backbone readiness. Registry validation is command-surface hygiene, not a runtime or security capability.

---

## 2. Current Public Release State

Public release:

    https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.4.0

Release title:

    CarbonStack v0.4.0 Broad Local Deployability Pre-Release

Release tag:

    v0.4.0

Release commit:

    24ac9fc494 / 24ac9fc docs: record v0.4.0 package rehearsal

Gitea release status:

    Pre-Release

Important release-page warning remains:

    The default Gitea Source Code ZIP/TAR.GZ downloads are auto-generated carbonstack-only archives.
    The intended runnable validation artifact is the attached multi-repo release package plus companion assets.

Post-release public surface:

    carbonstack README is evergreen.
    It points interested parties to the Releases page for runnable / known-good artifacts.
    Gitea is source of truth.
    GitHub mirrors are secondary push mirrors, not release authority.

---

## 3. Current Repo Heads

    carbonstack        568fb45 (HEAD -> main, origin/main, origin/HEAD) test: validate command registry coverage
    carbonstack-comms  cb4e59d (HEAD -> main, origin/main, origin/HEAD) test: add wrapper-based OpenMLS runtime smoke proof
    carbonstack-cypher 9ab994c (HEAD -> main, origin/main, origin/HEAD) docs: point to local operator runbook
    carbonstack-os     1bbbe52 (HEAD -> main, origin/main, origin/HEAD) docs: clarify CarbonStackOS target direction

Current public release tag:

    v0.4.0 -> 24ac9fc494 / 24ac9fc docs: record v0.4.0 package rehearsal

Recent runtime / maturity commits:

    v0.4.13 identity wrapper implementation:
      20c4cc9 feat: add dev OpenMLS identity bootstrap commands

    v0.4.14 bundle/conversation wrapper implementation:
      1f5c09e feat: add dev OpenMLS bundle and conversation bootstrap commands

    v0.4.15 add-member/join wrapper implementation:
      c2d3b16 feat: add dev OpenMLS add-member and join bootstrap commands

    v0.4.16 wrapper-based smoke proof:
      cb4e59d test: add wrapper-based OpenMLS runtime smoke proof

    v0.4.17 wrapper-smoke runner/profile alignment:
      a501442 feat: add wrapper OpenMLS runtime validation profile

    v0.4.18 command registry:
      a4162d7 docs: add known-good command registry

    Current v0.4.19 registry validation:
      568fb45 test: validate command registry coverage

Continuity note:

    `568fb45` is a carbonstack test/docs/registry/roadmap commit after the v0.4.0 public release tag.
    It does not retag v0.4.0.
    The release remains attached to `24ac9fc494`.

---

## 4. Current Validated / Known State

Validated / known at the v0.4.19 checkpoint:

    [RELEASE] v0.4.0 Gitea release exists, is marked Pre-Release, and is attached to 24ac9fc494.
    [RUNNER] full means release-snapshot followed by local-cypher.
    [RUNNER] dev-runtime-openmls validates the original direct-sidecar dev/pre-alpha OpenMLS application-message runtime CLI smoke proof.
    [RUNNER] dev-runtime-openmls-wrappers validates the wrapper-based dev/pre-alpha OpenMLS application-message runtime CLI smoke proof.
    [RUNNER] both runtime profiles remain manual-only and live-umbrella-only.
    [RUNNER] neither runtime profile is included in full.
    [REGISTRY] registry/commands.v0.yaml exists and remains provisional.
    [REGISTRY] registry/README.md records validation status.
    [REGISTRY] command_registry_test.go now validates the registry under Go test.
    [REGISTRY] registry structural recon found 62 entries, 62 command fields, no duplicate IDs, and no missing required markers.
    [REGISTRY] runner profile coverage had no missing and no extra runner IDs.
    [REGISTRY] Comms command coverage had no missing and no extra Comms IDs.
    [REGISTRY] scripts/README.md remains intentionally unregistered because it is documentation, not a script surface.
    [DOCS] docs/165-command-registry-validation-and-help-plan-v0.md records v0.4.19.
    [DOCS] docs/README.md indexes the v0.4.19 document.
    [DOCS] roadmap/ROADMAP.md records v0.4.19 and carries forward v0.5.0 release-ladder intent.
    [MIRRORS] all four repos have GitHub push-mirror functionality, with Gitea remaining source of truth.

v0.4.19 validation observed from the log:

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go test ./... -count=1 passed, including command_registry_test.go.
    go run . --profile dev-runtime-openmls-wrappers --clean-generated passed.
    go run . --profile dev-runtime-openmls --clean-generated passed.
    go run . --profile local-cypher passed.
    go run . --profile doctor passed.
    go run . --profile core --clean-generated passed.

    cd ~/repos/carbonstack_umbrella/carbonstack-comms
    go test ./internal/app -count=1 passed.
    go test ./... -count=1 -timeout 600s passed.

Final commit/push:

    carbonstack:
    568fb45 test: validate command registry coverage

Final repo snapshot:

    carbonstack        568fb45 test: validate command registry coverage
    carbonstack-comms  cb4e59d test: add wrapper-based OpenMLS runtime smoke proof
    carbonstack-cypher 9ab994c docs: point to local operator runbook
    carbonstack-os     1bbbe52 docs: clarify CarbonStackOS target direction

Expected live umbrella validation commands now include:

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go test ./... -count=1
    go run . --profile dev-runtime-openmls-wrappers --clean-generated
    go run . --profile dev-runtime-openmls --clean-generated
    go run . --profile local-cypher
    go run . --profile doctor
    go run . --profile core --clean-generated

Expected v0.4.0 public package validation commands remain unchanged:

    cd <package-root>/carbonstack/tools/carbonstack-validate
    go run . --profile verify-checksums --root <package-root>
    go run . --profile full --root <package-root> --clean-generated

---

## 5. v0.4.19 Work Completed

### 5.1 Registry validation Go test was added

New file:

    carbonstack/tools/carbonstack-validate/command_registry_test.go

Commit:

    568fb45 test: validate command registry coverage

The test validates:

    registry entries parse successfully.
    duplicate IDs are rejected.
    every entry has command, repo, component, kind, audience, maturity, source_path, short_help, why_exists, and include_in_front_readme.
    source_path values resolve against the live umbrella or carbonstack root.
    runner profile dispatch cases are represented in registry IDs.
    Comms command dispatch cases are represented in registry IDs.
    Comms scripts are represented in registry text, excluding scripts/README.md.
    expected sidecar command IDs are present.
    expected Cypher server/API IDs are present.
    registry retains the local-backbone nonclaim.
    registry retains the Gitea source-of-truth boundary.
    comms.send, comms.inbox, and comms.ack remain maturity: legacy.
    direct and wrapper OpenMLS runtime profiles keep distinct validation surfaces.
    include_in_front_readme count remains nonzero but small enough to prevent README sprawl.

The test skips only when run without a sibling umbrella layout containing carbonstack-comms and carbonstack-cypher.

### 5.2 Local help/manual planning was documented

New doc:

    carbonstack/docs/165-command-registry-validation-and-help-plan-v0.md

The plan records:

    B first: registry validation test/check.
    C later: generated Markdown command reference from registry.
    D later: Comms embedded help/manual strings only after command structure is ready.
    E later: runner profile for registry validation only if useful and stable.

It explicitly avoids for now:

    generated local command help.
    moving command text into commands.go prematurely.
    replacing component READMEs with generated docs.
    expanding the top README into a command encyclopedia.

### 5.3 README / registry / generated command reference boundary was recorded

Boundary decision:

    Front README:
      purpose, current release, source of truth, where to find releases, minimal validation entrypoint.
      only commands marked include_in_front_readme: true should be candidates.

    Registry:
      complete command/profile/script/API inventory.
      maturity, audience, why_exists, source_path, validation_surface, nonclaims, replacement status.
      source of truth for command-surface hygiene.

    Generated command reference, later:
      human-readable view derived from registry.
      grouped by repo/component/audience/maturity.
      should not replace release runbooks.
      should not become a claim surface without nonclaims preserved.

    Component READMEs:
      local component workflow.
      important examples.
      boundary warnings.
      not necessarily complete cross-stack inventory.

### 5.4 v0.5.0 minor epoch release carry-forward was recorded

v0.4.19 records the user preference that after late v0.4.x finishes, CarbonStack should cut a v0.5.0 minor epoch release using the accumulated v0.4.x work.

The v0.5.0 release ladder should have dedicated rungs for:

    cleanup.
    deployability testing.
    release-style replication.
    release package rehearsal.
    release hardening.
    public wording / claim-boundary cleanup.
    final LogDoc and breakpoint export.
    release asset generation.
    release notes using the established continuity style.

This should happen before major new v0.5.x state/trust/vault/PQ implementation begins unless a specific reason emerges to merge release and implementation work.

### 5.5 Docs and roadmap were updated

Updated:

    carbonstack/registry/README.md
    carbonstack/docs/README.md
    carbonstack/roadmap/ROADMAP.md

New:

    carbonstack/docs/165-command-registry-validation-and-help-plan-v0.md

The roadmap now records v0.4.19 and carries forward the post-v0.4.x v0.5.0 minor epoch release ladder.

### 5.6 Validation passed

Observed passing:

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go test ./... -count=1
    go run . --profile dev-runtime-openmls-wrappers --clean-generated
    go run . --profile dev-runtime-openmls --clean-generated
    go run . --profile local-cypher
    go run . --profile doctor
    go run . --profile core --clean-generated

    cd ~/repos/carbonstack_umbrella/carbonstack-comms
    go test ./internal/app -count=1
    go test ./... -count=1 -timeout 600s

Final repo snapshot was clean across all four repos.

---

## 6. Critical Blunders / Lessons to Carry Forward

### 6.1 Large pasteback glitch risk persists, but final evidence was sufficient

The v0.4.19 recon and implementation pastebacks were still large and some earlier lines were paste-glitched. However, the final implementation log contained the authoritative commit/push and final repo snapshot.

Lesson:

    Continue using recon -> pasteback -> implement.
    Continue splitting large work into blocks.
    Treat final commit/push plus final four-repo clean snapshot as breakpoint authority.

### 6.2 Registry completeness is now tested, but still policy-v0

The registry now has Go-test coverage, but the policy is still intentionally v0.

Lesson:

    Registry validation should evolve conservatively.
    Do not promote registry validation to a runner profile until the test is stable and useful.
    Do not assume registry v0 is a perfect map of all future command surfaces.

### 6.3 Source-path validation now creates a maintenance responsibility

Because command_registry_test.go validates source_path existence, moving files or changing repo structure now requires registry updates.

Lesson:

    Any CLI command, runner profile, smoke script, sidecar command, release helper, operator/API surface, or path movement must update the registry in the same checkpoint or next checkpoint.

### 6.4 scripts/README.md exception is intentional

The recon surfaced `scripts/README.md` as an unregistered file under Comms scripts. It is documentation, not a command surface, and the test explicitly excludes it.

Lesson:

    Do not over-register documentation files as commands just because they live in a scripts directory.

### 6.5 Local-help generation remains deferred

The project now has enough metadata for future help/manual generation, but v0.4.19 deliberately avoids generating docs or embedding help strings.

Lesson:

    Use the registry as source-of-truth metadata first.
    Add generated references only after the registry proves stable.
    Do not move help text into Comms command registration until command structure is ready.

### 6.6 Commit metadata warning remains a recurring polish issue

The commit log again showed auto-generated local commit metadata:

    Committer: bitcrusher32 <▓▓>

This is not a functional validation failure, but remains a professionalism/publishing hygiene issue to address eventually.

---

## 7. Current Blockers / Not Validated

Current blockers and nonvalidated areas after v0.4.19:

    Existing send/inbox remain stub-era.
    Runtime Comms has dev-only OpenMLS commands, direct smoke proof, wrapper-based smoke proof, and separate manual runner profiles, but no mature user-facing UX.
    Neither runtime profile is included in full.
    Neither runtime profile is part of public release package validation.
    Fresh release-package behavior for dev-runtime-openmls or dev-runtime-openmls-wrappers has not been tested as a supported validation surface.
    No Relay Space join model exists yet.
    local-backbone runner profile does not exist and remains reserved.
    Registry validation is now a Go test, but not a runner profile.
    Generated command reference from the registry does not exist yet.
    Local help/manual generation from the registry does not exist yet.
    Short explainer strings are not yet integrated directly beside command registration in Comms.
    README/dev command sprawl is mitigated by registry + test, but not fully solved.
    The long-term roadmap PDF is outdated after v0.4.19 and needs refresh next.
    v0.5.0 minor epoch release ladder is planned but not implemented.
    Post-release public download verification is still useful if not already performed.
    Sanitized PRIME placement in carbonstack/logdoc-loop-system still needs a policy decision.
    Production-safe provider storage and secure vault design remain v0.5.x future work.
    PQ/hybrid ciphersuite migration readiness remains v0.5.x future work.
    Hostile-server rollback/replay/metadata-abuse validation remains future work.
    Android/CarbonStackOS implementation remains later work.
    CarbonStack Relay Space mechanics remain a future concept, not an implementation.

Hard nonclaims remain:

    Do not claim production readiness.
    Do not claim production E2EE.
    Do not claim hostile-server safety.
    Do not claim metadata privacy.
    Do not claim Debian production deployability.
    Do not claim systemd readiness.
    Do not claim cloudflared readiness.
    Do not claim Android readiness.
    Do not claim CarbonStackOS readiness.
    Do not claim secure vault/storage.
    Do not claim real homelab validation.
    Do not claim external audit or certification.
    Do not claim CarbonStack Relay Space mechanics are implemented.
    Do not claim local-backbone exists or is justified yet.
    Do not claim local-cypher validates runtime Comms UX.
    Do not claim full validates local-backbone.
    Do not claim full is a deployment command.
    Do not claim dev-runtime-openmls is release-package validation yet.
    Do not claim dev-runtime-openmls-wrappers is release-package validation yet.
    Do not claim send/inbox are OpenMLS-backed yet.
    Do not claim PQ/hybrid security or quantum-safe messaging.
    Do not let GitHub mirrors supersede Gitea source-of-truth status.

Allowed claim:

    CarbonStack now has a provisional cross-repo known-good command registry at `registry/commands.v0.yaml`.
    The registry is validated by `tools/carbonstack-validate/command_registry_test.go` when run from a full umbrella/release-package-like layout.
    The registry tracks runner profiles, Comms CLI commands, old stub-era send/inbox/ack, OpenMLS dev runtime commands, OpenMLS bootstrap wrappers, Comms smoke scripts, OpenMLS sidecar commands, Cypher CLI/env/API surfaces, and legacy PowerShell helpers.
    The registry includes maturity, audience, validation surface, short help, why-exists, nonclaims, and front-door README inclusion metadata.
    The local help/manual plan is recorded, but generated docs and embedded help are deferred.
    CarbonStackComms still has dev-only openmls-send-dev and openmls-inbox-dev commands.
    CarbonStackComms still has direct-sidecar and wrapper-based dev runtime smoke scripts.
    carbonstack/tools/carbonstack-validate still has separate manual dev-runtime-openmls and dev-runtime-openmls-wrappers profiles.
    Both runtime profiles are live-git-umbrella-only and not included in full.
    Existing send/inbox remain stub-era.
    v0.4.0 remains the current public broad local deployability pre-release.
    Gitea remains source of truth; GitHub mirrors are secondary push mirrors.

---

## 8. Next Safest Actions

Recommended next rung:

    Refresh the long-term roadmap after the grouped registry/help-planning goalset.

Focus:

    Update the v0.4.14 roadmap to the post-v0.4.19 state.
    Record completed v0.4.15 through v0.4.19 work.
    Carry forward late-v0.4.x cleanup/maturity rungs.
    Add the v0.5.0 minor epoch release ladder before major v0.5.x state/trust/vault/PQ implementation.
    Preserve PQ/hybrid migration caution.
    Preserve local-backbone reserve.
    Preserve direct vs wrapper runtime profile split.
    Preserve command registry as source-of-truth metadata.

Likely candidates after roadmap refresh:

    v0.4.20+ sidecar/dev-state reset semantics.
    generated command reference planning/implementation if still desired.
    exact sidecar setup runbook.
    clearer dev state reset semantics.
    better command output shape.
    decision on whether stub-era send/inbox are aliased, replaced, or deprecated.
    decision on when local-backbone naming stops being premature.
    v0.5.0 release cleanup / deployability testing / release-style replication / release hardening.

Avoid next:

    replacing the direct smoke script.
    merging dev-runtime-openmls and dev-runtime-openmls-wrappers.
    adding either runtime profile to full.
    calling either runtime profile local-backbone.
    treating wrapper bootstrap as Relay Space join UX.
    moving into PQ/vault work before roadmap refresh and state/trust/storage preflight.
    public ingress.
    cloudflared.
    systemd.
    real homelab deployment.
    Android app.
    CarbonStackOS implementation.
    production/security claims.

---

## 9. Critical Paths / Functions

### Source-of-truth public surfaces

    [GITEA SOURCE OF TRUTH] https://git.bitcrusher32.win/bitcrusher32/carbonstack
    [RELEASES] https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases
    [CURRENT PUBLIC RELEASE] https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.4.0

GitHub mirrors:

    secondary push mirrors for discoverability/redundancy only.
    Do not treat mirrors as release authority unless project policy changes.

### WSL Debian active working umbrella

    [WSL-DEBIAN] umbrella: ~/repos/carbonstack_umbrella
    [WSL-DEBIAN] carbonstack: ~/repos/carbonstack_umbrella/carbonstack
    [WSL-DEBIAN] carbonstack-comms: ~/repos/carbonstack_umbrella/carbonstack-comms
    [WSL-DEBIAN] carbonstack-cypher: ~/repos/carbonstack_umbrella/carbonstack-cypher
    [WSL-DEBIAN] carbonstack-os: ~/repos/carbonstack_umbrella/carbonstack-os
    [GO-RUNNER] runner dir: ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate

### Registry surfaces

    [REGISTRY] root:
    carbonstack/registry

    [REGISTRY] provisional command registry:
    carbonstack/registry/commands.v0.yaml

    [REGISTRY] registry README:
    carbonstack/registry/README.md

    [REGISTRY-TEST] command registry validation:
    carbonstack/tools/carbonstack-validate/command_registry_test.go

    [DOC] v0.4.18 registry checkpoint:
    carbonstack/docs/164-known-good-command-registry-recon-v0.md

    [DOC] v0.4.19 registry validation/help plan:
    carbonstack/docs/165-command-registry-validation-and-help-plan-v0.md

### Runtime integration doc surfaces

    [DOC] v0.4.1 runtime recon:
    carbonstack/docs/155-runtime-comms-openmls-cypher-recon-v0.md

    [DOC] v0.4.2 command contract:
    carbonstack/docs/156-runtime-openmls-command-contract-v0.md

    [DOC] v0.4.6 smoke proof record:
    carbonstack/docs/157-dev-runtime-openmls-smoke-proof-v0.md

    [DOC] v0.4.7 pre-local-backbone assessment:
    carbonstack/docs/158-pre-local-backbone-assessment-v0.md

    [DOC] v0.4.9 runner profile:
    carbonstack/docs/159-dev-runtime-openmls-runner-profile-v0.md

    [DOC] v0.4.10 profile boundary:
    carbonstack/docs/160-dev-runtime-openmls-profile-boundary-v0.md

    [DOC] v0.4.11 bootstrap wrapper recon:
    carbonstack/docs/161-openmls-bootstrap-wrapper-recon-v0.md

    [DOC] v0.4.12 bootstrap command contract:
    carbonstack/docs/162-openmls-bootstrap-command-contract-v0.md

    [DOC] v0.4.17 wrapper profile:
    carbonstack/docs/163-dev-runtime-openmls-wrapper-runner-profile-v0.md

    [DOC] docs index:
    carbonstack/docs/README.md

    [DOC] roadmap:
    carbonstack/roadmap/ROADMAP.md

    [DOC] top README:
    carbonstack/README.md

### Runner surfaces

    [RUNNER] root:
    carbonstack/tools/carbonstack-validate

    [RUNNER] dispatch:
    carbonstack/tools/carbonstack-validate/main.go

    [RUNNER] registry validation test:
    carbonstack/tools/carbonstack-validate/command_registry_test.go

    [RUNNER] dev-runtime-openmls implementation:
    carbonstack/tools/carbonstack-validate/dev_runtime_openmls.go

    [RUNNER] dev-runtime-openmls-wrappers implementation:
    carbonstack/tools/carbonstack-validate/dev_runtime_openmls_wrappers.go

    [RUNNER] runner README:
    carbonstack/tools/carbonstack-validate/README.md

    [RUNNER] local-cypher:
    carbonstack/tools/carbonstack-validate/local_cypher.go

    [RUNNER] release-snapshot:
    carbonstack/tools/carbonstack-validate/release_snapshot.go

    [RUNNER] checksums:
    carbonstack/tools/carbonstack-validate/checksums.go

### Comms runtime surfaces

    [COMMS] CLI entry:
    carbonstack-comms/cmd/comms/main.go

    [COMMS] command registry and non-OpenMLS commands:
    carbonstack-comms/internal/app/commands.go

    [COMMS] OpenMLS message runtime command helpers:
    carbonstack-comms/internal/app/openmls_runtime.go

    [COMMS] OpenMLS bootstrap wrapper helpers:
    carbonstack-comms/internal/app/openmls_bootstrap.go

    [COMMS] OpenMLS bootstrap wrapper tests:
    carbonstack-comms/internal/app/openmls_bootstrap_dev_test.go

    [COMMS] openmls-send-dev tests:
    carbonstack-comms/internal/app/openmls_send_dev_test.go

    [COMMS] openmls-inbox-dev tests:
    carbonstack-comms/internal/app/openmls_inbox_dev_test.go

    [COMMS] direct-sidecar dev runtime smoke proof:
    carbonstack-comms/scripts/dev-openmls-runtime-smoke.sh

    [COMMS] wrapper-based dev runtime smoke proof:
    carbonstack-comms/scripts/dev-openmls-runtime-smoke-wrappers.sh

    [COMMS] OpenMLS sidecar dir:
    carbonstack-comms/internal/protocol/mls/openmls-sidecar

    [COMMS] sidecar README:
    carbonstack-comms/internal/protocol/mls/openmls-sidecar/README.md

### v0.4.19 runner/test behavior

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go test ./... -count=1

    Includes command_registry_test.go:
      validates registry structure and coverage.
      skips without a sibling umbrella layout.
      should be treated as the current registry drift guard.

### Existing v0.4.0 package validation

    cd <package-root>/carbonstack/tools/carbonstack-validate
    go run . --profile verify-checksums --root <package-root>
    go run . --profile full --root <package-root> --clean-generated

---

## 10. v0.4.x Forward Plan

### Immediate next: roadmap refresh after grouped registry/help-planning goalset

Expected:

    Update the roadmap from v0.4.14 to post-v0.4.19.
    Record completed v0.4.15 add-member/join wrappers.
    Record completed v0.4.16 wrapper-based smoke.
    Record completed v0.4.17 wrapper runtime profile.
    Record completed v0.4.18 command registry.
    Record completed v0.4.19 registry validation/local-help plan.
    Carry forward late-v0.4.x cleanup/maturity tasks.
    Carry forward v0.5.0 minor epoch release target.
    Preserve v0.5.x state/trust/vault/PQ as post-release implementation direction.

### v0.4.20+

Possible:

    generated command reference from registry.
    exact sidecar setup runbook.
    clearer dev state reset semantics.
    better command output shape.
    decision on whether stub-era send/inbox are aliased, replaced, or deprecated.
    decision on when local-backbone naming stops being premature.
    v0.5.0 release cleanup / deployability testing / release-style replication / release hardening.

---

## 11. v0.5.x and Later Carry-Forward

### v0.5.0 minor epoch release

After late v0.4.x finishes, cut a v0.5.0 minor epoch release using the accumulated v0.4.x work.

Dedicated rungs should include:

    cleanup.
    deployability testing.
    release-style replication.
    release package rehearsal.
    release hardening.
    public wording / claim-boundary cleanup.
    final LogDoc and breakpoint export.
    release asset generation.
    release notes using the established continuity style.

### v0.5.x implementation after release

Local storage, trust state, provider state, vault design, and PQ/hybrid ciphersuite migration readiness.

Do not start PQ/hybrid implementation before the runtime OpenMLS/Cypher path exists and is documented. PQ/hybrid migration touches provider state, ciphersuite metadata, trust display, artifact size, reinit/rekey/recovery policy, and claim boundaries.

### v0.6.x+

Hostile-server and abuse-resistance harnesses.

### v0.7.x+

Deployability and operations hardening.

### v0.8.x+

Documented self-pentest / adversarial validation.

### v0.9.x+

Claim-boundary review.

### v0.10.x+

Android backend/app work.

### v1.x.x

Public app/server major epoch and much later CarbonStackOS/appliance research, if justified.

---

## 12. Lean Breakpoint Summary

    CarbonStack v0.4.19 is the command registry validation and local-help/manual planning checkpoint. v0.4.0 remains the current public broad local deployability pre-release. carbonstack is now at 568fb45 test: validate command registry coverage. carbonstack-comms remains at cb4e59d test: add wrapper-based OpenMLS runtime smoke proof. v0.4.19 adds tools/carbonstack-validate/command_registry_test.go and docs/165-command-registry-validation-and-help-plan-v0.md, updates registry/README.md, docs/README.md, and roadmap/ROADMAP.md. The registry remains provisional but now has Go-test coverage for structure, required fields, source_path existence, runner coverage, Comms command coverage, Comms script coverage, sidecar IDs, Cypher IDs, legacy send/inbox/ack classification, direct-vs-wrapper OpenMLS runtime profile separation, local-backbone/Gitea boundaries, and include_in_front_readme count sanity. Local help/manual planning is recorded: validate first, generate Markdown references later, defer Comms embedded help strings, and only promote registry validation to a runner profile after it proves stable. Validation passed for carbonstack tools go tests, dev-runtime-openmls-wrappers --clean-generated, dev-runtime-openmls --clean-generated, local-cypher, doctor, core --clean-generated, Comms app tests, and full Comms tests. Final repo heads were clean across all four repos. Next safe action: refresh the long-term roadmap for post-v0.4.19 state, then continue late-v0.4.x cleanup and eventual v0.5.0 minor epoch release preparation.

---

## 13. Preserved Immediate Previous Handoff: v0.4.18

The following is the previous v0.4.18 handoff. Where it conflicts with the v0.4.19 overlay above, v0.4.19 wins for current state.


# CarbonStack LogDoc v0.4.18

**Last updated:** 2026-06-05 local session  
**Owner/operator:** bitcrusher32  
**Assistant/model:** GPT-5.5 Thinking  
**Current phase:** Phase 2F / **v0.4.x known-good command registry checkpoint complete**. The v0.4.0 public broad local deployability pre-release remains the current public release. v0.4.17 promoted the wrapper-based OpenMLS runtime smoke proof into a separate manual `dev-runtime-openmls-wrappers` runner profile; v0.4.18 now adds the first provisional cross-repo known-good command registry in `carbonstack/registry`, while preserving `dev-runtime-openmls` and `dev-runtime-openmls-wrappers` as separate manual profiles, keeping `full` unchanged, and keeping `local-backbone` reserved. `carbonstack` is now at `a4162d7 (HEAD -> main, origin/main, origin/HEAD) docs: add known-good command registry`; `carbonstack-comms` remains at `cb4e59d (HEAD -> main, origin/main, origin/HEAD) test: add wrapper-based OpenMLS runtime smoke proof`; `carbonstack-cypher` remains at `9ab994c (HEAD -> main, origin/main, origin/HEAD) docs: point to local operator runbook`; `carbonstack-os` remains at `1bbbe52 (HEAD -> main, origin/main, origin/HEAD) docs: clarify CarbonStackOS target direction`.

**Version schema:** `v[scope].[timeline]`. This file is `v0.4.18`, the known-good command registry checkpoint after the v0.4.17 wrapper-smoke runner/profile alignment checkpoint. Per LogDoc V2 practice, this Markdown file preserves current operational continuity and carries the v0.4.x runtime-integration / pre-v0.5.x maturity thread forward; the JSON breakpoint remains the lean machine-readable handoff state.

**Compression note:** The detailed v0.3.0 through v0.3.36PRIME process ledger remains archived in PRIME/sanitized PRIME. v0.4.0 compressed that pre-release sprint; v0.4.1 through v0.4.12 preserved runtime recon, command contracts, dev OpenMLS send/inbox commands, smoke proof, runner profile, boundary checks, wrapper recon, and bootstrap command contract. v0.4.13 through v0.4.15 preserved staged bootstrap wrapper implementation. v0.4.16 preserved the first committed wrapper-based runtime smoke variant. v0.4.17 preserved the separate manual runner profile for that wrapper smoke. v0.4.18 now preserves the first provisional cross-repo command registry.

---

## 1. Project Goal

**Active goal:** Continue late-v0.4.x maturity by making CarbonStack’s command/profile/script/API surface explicit before v0.5.x state/trust/vault/PQ work expands the surface area. v0.4.18 creates a provisional known-good command registry in the main `carbonstack` repo so the project can keep critical, dev-only, internal, legacy, release, runner, sidecar, Cypher, and script surfaces legible without turning front-door READMEs into command encyclopedias.

CarbonStack’s broader project intent remains:

- **CarbonStack** is the secure-communications appliance-stack project and shared doctrine/release authority.
- **CarbonStackComms** is the text-first messaging client and current OpenMLS sidecar/relay integration surface.
- **CarbonStackCypher** is the self-hostable relay/storage stack.
- **CarbonStackOS** is future constrained Android-derived appliance OS work and remains out of current runnable validation packages.
- The `carbonstack` Gitea repo remains the source of truth for release state, release assets, docs, registry, and public claims.
- GitHub mirrors exist as secondary push mirrors only; Gitea remains authoritative.

**North star:**

> Its goal is not to be a better smartphone. Its goal is to stop being a smartphone.

**Core doctrine:**

> Every feature is guilty until it proves it does not add unacceptable parser, network, sensor, identity, or filesystem authority.

**Current reality after v0.4.18:**

    carbonstack/registry now exists.
    registry/commands.v0.yaml is the first provisional cross-repo command registry.
    registry/README.md explains registry purpose, boundaries, maintenance rule, and front-door README rule.
    docs/164-known-good-command-registry-recon-v0.md records the v0.4.18 registry checkpoint.
    docs/README.md indexes the v0.4.18 registry doc.
    roadmap/ROADMAP.md records the v0.4.18 command registry checkpoint.

    The registry currently tracks 62 entries after marker validation:
      runner profiles
      Comms CLI commands
      old stub-era send/inbox/ack commands
      OpenMLS dev runtime commands
      OpenMLS bootstrap wrapper commands
      Comms smoke scripts
      OpenMLS sidecar commands
      Cypher CLI/env/API surfaces
      legacy PowerShell helper scripts

    The registry uses fields such as:
      id
      command
      repo
      component
      kind
      audience
      maturity
      introduced_in where known
      source_path
      validation_surface
      short_help
      why_exists
      example where useful
      nonclaims
      related profiles/scripts
      replacement/deprecation status
      include_in_front_readme

    Existing runtime proof split remains:
      dev-runtime-openmls:
        direct sidecar bootstrap -> openmls-send-dev -> Cypher -> openmls-inbox-dev --ack
      dev-runtime-openmls-wrappers:
        openmls-*-dev bootstrap wrappers -> openmls-send-dev -> Cypher -> openmls-inbox-dev --ack

    dev-runtime-openmls remains the direct-smoke baseline.
    dev-runtime-openmls-wrappers remains the wrapper-smoke maturity surface.
    Both runtime profiles remain manual-only and live-umbrella-only.
    Neither runtime profile is included in full.
    full remains release-snapshot + local-cypher.
    Existing send/inbox remain stub-era and are explicitly included in the registry as legacy/stub-era surfaces.
    local-backbone remains reserved.

Hard nonclaims remain: v0.4.18 does **not** prove production readiness, production E2EE, hostile-server safety, metadata privacy, Android readiness, CarbonStackOS readiness, secure vault/storage, public ingress safety, real homelab validation, audit, certification, PQ/hybrid security, mature messenger UX, release-package runtime validation, Relay Space join/onboarding, `full` runtime OpenMLS validation, or local-backbone readiness. The registry is metadata and claim-boundary hygiene, not a runtime or security capability.

---

## 2. Current Public Release State

Public release:

    https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.4.0

Release title:

    CarbonStack v0.4.0 Broad Local Deployability Pre-Release

Release tag:

    v0.4.0

Release commit:

    24ac9fc494 / 24ac9fc docs: record v0.4.0 package rehearsal

Gitea release status:

    Pre-Release

Important release-page warning remains:

    The default Gitea Source Code ZIP/TAR.GZ downloads are auto-generated carbonstack-only archives.
    The intended runnable validation artifact is the attached multi-repo release package plus companion assets.

Post-release public surface:

    carbonstack README is evergreen.
    It points interested parties to the Releases page for runnable / known-good artifacts.
    Gitea is source of truth.
    GitHub mirrors are secondary push mirrors, not release authority.

---

## 3. Current Repo Heads

    carbonstack        a4162d7 (HEAD -> main, origin/main, origin/HEAD) docs: add known-good command registry
    carbonstack-comms  cb4e59d (HEAD -> main, origin/main, origin/HEAD) test: add wrapper-based OpenMLS runtime smoke proof
    carbonstack-cypher 9ab994c (HEAD -> main, origin/main, origin/HEAD) docs: point to local operator runbook
    carbonstack-os     1bbbe52 (HEAD -> main, origin/main, origin/HEAD) docs: clarify CarbonStackOS target direction

Current public release tag:

    v0.4.0 -> 24ac9fc494 / 24ac9fc docs: record v0.4.0 package rehearsal

Recent runtime / maturity commits:

    v0.4.8 extraction:
      0a48ae1 refactor: split OpenMLS runtime command helpers

    v0.4.9 runner profile:
      8eeadb2 feat: add dev OpenMLS runtime validation profile

    v0.4.10 profile boundary:
      7384f43 docs: record dev OpenMLS runtime profile boundary

    v0.4.11 wrapper recon:
      8ee3ebb docs: plan dev OpenMLS bootstrap wrappers

    v0.4.12 command contract:
      0583683 docs: define dev OpenMLS bootstrap command contract

    v0.4.13 identity wrapper implementation:
      20c4cc9 feat: add dev OpenMLS identity bootstrap commands

    v0.4.14 bundle/conversation wrapper implementation:
      1f5c09e feat: add dev OpenMLS bundle and conversation bootstrap commands

    v0.4.15 add-member/join wrapper implementation:
      c2d3b16 feat: add dev OpenMLS add-member and join bootstrap commands

    v0.4.16 wrapper-based smoke proof:
      cb4e59d test: add wrapper-based OpenMLS runtime smoke proof

    v0.4.17 wrapper-smoke runner/profile alignment:
      a501442 feat: add wrapper OpenMLS runtime validation profile

    Current v0.4.18 command registry:
      a4162d7 docs: add known-good command registry

Continuity note:

    `a4162d7` is a carbonstack docs/registry/roadmap commit after the v0.4.0 public release tag.
    It does not retag v0.4.0.
    The release remains attached to `24ac9fc494`.

---

## 4. Current Validated / Known State

Validated / known at the v0.4.18 checkpoint:

    [RELEASE] v0.4.0 Gitea release exists, is marked Pre-Release, and is attached to 24ac9fc494.
    [RUNNER] full means release-snapshot followed by local-cypher.
    [RUNNER] dev-runtime-openmls validates the original direct-sidecar dev/pre-alpha OpenMLS application-message runtime CLI smoke proof.
    [RUNNER] dev-runtime-openmls-wrappers validates the wrapper-based dev/pre-alpha OpenMLS application-message runtime CLI smoke proof.
    [RUNNER] both runtime profiles remain manual-only and live-umbrella-only.
    [RUNNER] neither runtime profile is included in full.
    [REGISTRY] registry/commands.v0.yaml exists.
    [REGISTRY] registry/README.md exists.
    [REGISTRY] registry marker validation passed with 62 entries and no duplicate IDs.
    [REGISTRY] old stub-era send, inbox, and ack are included as legacy/stub-era surfaces.
    [REGISTRY] direct and wrapper OpenMLS runtime profiles are kept separate.
    [DOCS] docs/164-known-good-command-registry-recon-v0.md records the registry checkpoint.
    [DOCS] docs/README.md indexes the v0.4.18 registry document.
    [DOCS] roadmap/ROADMAP.md records the v0.4.18 command registry checkpoint.
    [MIRRORS] all four repos have GitHub push-mirror functionality, with Gitea remaining source of truth.

v0.4.18 validation observed from the log:

    registry marker check passed: 62 entries.
    go test ./... -count=1 passed in carbonstack/tools/carbonstack-validate.
    go run . --profile dev-runtime-openmls-wrappers --clean-generated passed.
    go run . --profile dev-runtime-openmls --clean-generated passed.
    go run . --profile local-cypher passed.
    go run . --profile doctor passed.
    go run . --profile core --clean-generated passed.
    carbonstack was committed and pushed at:
      a4162d7 docs: add known-good command registry
    final repo status showed all four repos clean.

Expected live umbrella validation commands now include:

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go run . --profile dev-runtime-openmls-wrappers --clean-generated
    go run . --profile dev-runtime-openmls --clean-generated
    go run . --profile local-cypher
    go run . --profile doctor
    go run . --profile core --clean-generated

Expected v0.4.0 public package validation commands remain unchanged:

    cd <package-root>/carbonstack/tools/carbonstack-validate
    go run . --profile verify-checksums --root <package-root>
    go run . --profile full --root <package-root> --clean-generated

---

## 5. v0.4.18 Work Completed

### 5.1 First provisional command registry was added

New directory:

    carbonstack/registry

New files:

    carbonstack/registry/commands.v0.yaml
    carbonstack/registry/README.md

Commit:

    a4162d7 docs: add known-good command registry

Registry schema version:

    carbonstack-command-registry/v0

Registry status:

    provisional

Registry purpose:

    Cross-repo registry of known command, script, profile, sidecar, operator, and API surfaces that are easy to lose track of as CarbonStack grows.

The registry is explicitly not:

    a release promise.
    a production security claim.
    local-backbone.
    a complete security model.
    release-package runtime validation.
    proof of production readiness, production E2EE, hostile-server safety, metadata privacy, PQ/hybrid security, audit, or certification.

### 5.2 Registry coverage was established

The first registry pass covers:

    carbonstack validation runner profiles:
      doctor
      core
      local-cypher
      dev-runtime-openmls
      dev-runtime-openmls-wrappers
      full
      release-snapshot
      write-checksums
      verify-checksums

    carbonstack-comms CLI:
      init
      dev-create-invite
      claim-invite
      register-device
      list-devices
      fingerprint
      verify-device
      trust-history
      trust-list
      simulate-key-change
      revoke-device
      send
      inbox
      ack
      openmls-send-dev
      openmls-inbox-dev
      openmls-identity-create-dev
      openmls-identity-status-dev
      openmls-bundle-export-dev
      openmls-conversation-create-dev
      openmls-conversation-load-check-dev
      openmls-conversation-add-member-dev
      openmls-conversation-join-dev

    carbonstack-comms scripts:
      dev-openmls-runtime-smoke.sh
      dev-openmls-runtime-smoke-wrappers.sh
      self-test-openmls-backbone.ps1
      smoke-openmls-real-cypher-relay.ps1
      check-no-rust-artifacts.ps1
      test-local-lifecycle.ps1
      test-trust-lifecycle.ps1

    OpenMLS sidecar commands:
      provider-info
      identity-create
      identity-status
      public-bundle-export
      conversation-create
      conversation-load-check
      conversation-add-member
      conversation-join
      message-protect
      message-open
      state-checkpoint as unsupported/future
      state-load-check as unsupported/future

    carbonstack-cypher:
      go run ./cmd/cypher
      key environment variables
      health/invite/device/envelope/ack API surfaces

    carbonstack legacy scripts:
      scripts/validate-local.ps1
      scripts/validate-phase1.ps1

### 5.3 Registry docs and roadmap were updated

New doc:

    carbonstack/docs/164-known-good-command-registry-recon-v0.md

Updated docs index:

    carbonstack/docs/README.md

Updated roadmap:

    carbonstack/roadmap/ROADMAP.md

The new doc records:

    why the registry exists
    why it lives in main carbonstack
    what surfaces are covered
    the v0 schema fields
    why stub-era send/inbox/ack are included as legacy
    what v0.4.18 does not change
    why this matters before v0.5.x
    suggested next rung: command registry validation / local-help planning

The registry README records:

    registry purpose
    registry boundaries
    maintenance rule
    front-door README rule

Important front-door README policy:

    Only commands marked `include_in_front_readme: true` should be considered for top-level/front-door docs.
    Dev/internal/legacy command surfaces should usually live in registry docs, component READMEs, or implementation-specific docs.

### 5.4 Validation passed

Commands observed passing:

    cd ~/repos/carbonstack_umbrella/carbonstack
    python3 registry marker check
      registry marker check passed: 62 entries

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go test ./... -count=1
    go run . --profile dev-runtime-openmls-wrappers --clean-generated
    go run . --profile dev-runtime-openmls --clean-generated
    go run . --profile local-cypher
    go run . --profile doctor
    go run . --profile core --clean-generated

Final commit/push:

    carbonstack:
    a4162d7 docs: add known-good command registry

Final repo snapshot:

    carbonstack        a4162d7 docs: add known-good command registry
    carbonstack-comms  cb4e59d test: add wrapper-based OpenMLS runtime smoke proof
    carbonstack-cypher 9ab994c docs: point to local operator runbook
    carbonstack-os     1bbbe52 docs: clarify CarbonStackOS target direction

---

## 6. Critical Blunders / Lessons to Carry Forward

### 6.1 Wrong copy endpoint / incomplete evidence blunder

The first attempted v0.4.18 handoff paste ended before commit/push and the final four-repo snapshot. That made it insufficient for a stable breakpoint even though registry generation and validation had passed.

Lesson:

    Do not cut a LogDoc breakpoint without commit/push evidence and final repo heads/status.
    The final repo snapshot remains the authority.

### 6.2 Recon pasteback remained large and partially glitched

The command-registry recon and diff output were very large. The copy/paste still had truncation/glitch risk, but the finalization log contained enough evidence: registry marker check, validation passes, commit/push, and final clean heads.

Lesson:

    Keep using recon -> pasteback -> implement.
    Keep splitting large work into smaller paste blocks.
    Do not treat a large scan as complete merely because it is long.
    Use marker checks and final repo snapshots as the stable endpoint.

### 6.3 First push failed, second push succeeded

The finalization log shows an initial authentication failure on `git push`, then a second `git push` succeeded.

Lesson:

    A failed first push is not a committed remote state.
    Confirm remote state with the final `origin/main` head in the four-repo snapshot.

### 6.4 Commit metadata warning remains a recurring polish issue

The commit log again showed auto-generated local commit metadata:

    Committer: bitcrusher32 <▓▓>

This is not a functional validation failure, but remains a professionalism/publishing hygiene issue to address eventually.

### 6.5 Registry is provisional and must not become stale

The registry is only useful if it is updated as command surfaces change.

Lesson:

    When adding a CLI command, runner profile, smoke script, sidecar command, release helper, or operator/API surface, update the registry in the same checkpoint or next checkpoint.

---

## 7. Current Blockers / Not Validated

Current blockers and nonvalidated areas after v0.4.18:

    Existing send/inbox remain stub-era.
    Runtime Comms has dev-only OpenMLS commands, direct smoke proof, wrapper-based smoke proof, and separate manual runner profiles, but no mature user-facing UX.
    Neither runtime profile is included in full.
    Neither runtime profile is part of public release package validation.
    Fresh release-package behavior for dev-runtime-openmls or dev-runtime-openmls-wrappers has not been tested as a supported validation surface.
    No Relay Space join model exists yet.
    local-backbone runner profile does not exist and remains reserved.
    Registry validation is still a basic marker/duplicate check, not a full schema/dispatch consistency validator.
    Local help/manual generation from the registry does not exist yet.
    Short explainer strings are not yet integrated directly beside command registration in Comms.
    README/dev command sprawl is mitigated by the registry, but not fully solved.
    Post-release public download verification is still useful if not already performed.
    Sanitized PRIME placement in carbonstack/logdoc-loop-system still needs a policy decision.
    Production-safe provider storage and secure vault design remain v0.5.x future work.
    PQ/hybrid ciphersuite migration readiness remains v0.5.x future work.
    Hostile-server rollback/replay/metadata-abuse validation remains future work.
    Android/CarbonStackOS implementation remains later work.
    CarbonStack Relay Space mechanics remain a future concept, not an implementation.

Hard nonclaims remain:

    Do not claim production readiness.
    Do not claim production E2EE.
    Do not claim hostile-server safety.
    Do not claim metadata privacy.
    Do not claim Debian production deployability.
    Do not claim systemd readiness.
    Do not claim cloudflared readiness.
    Do not claim Android readiness.
    Do not claim CarbonStackOS readiness.
    Do not claim secure vault/storage.
    Do not claim real homelab validation.
    Do not claim external audit or certification.
    Do not claim CarbonStack Relay Space mechanics are implemented.
    Do not claim local-backbone exists or is justified yet.
    Do not claim local-cypher validates runtime Comms UX.
    Do not claim full validates local-backbone.
    Do not claim full is a deployment command.
    Do not claim dev-runtime-openmls is release-package validation yet.
    Do not claim dev-runtime-openmls-wrappers is release-package validation yet.
    Do not claim send/inbox are OpenMLS-backed yet.
    Do not claim PQ/hybrid security or quantum-safe messaging.
    Do not let GitHub mirrors supersede Gitea source-of-truth status.

Allowed claim:

    CarbonStack now has a provisional cross-repo known-good command registry at `registry/commands.v0.yaml`.
    The registry tracks runner profiles, Comms CLI commands, old stub-era send/inbox/ack, OpenMLS dev runtime commands, OpenMLS bootstrap wrappers, Comms smoke scripts, OpenMLS sidecar commands, Cypher CLI/env/API surfaces, and legacy PowerShell helpers.
    The registry includes maturity, audience, validation surface, short help, why-exists, nonclaims, and front-door README inclusion metadata.
    CarbonStackComms still has dev-only openmls-send-dev and openmls-inbox-dev commands.
    CarbonStackComms still has direct-sidecar and wrapper-based dev runtime smoke scripts.
    carbonstack/tools/carbonstack-validate still has separate manual dev-runtime-openmls and dev-runtime-openmls-wrappers profiles.
    Both runtime profiles are live-git-umbrella-only and not included in full.
    Existing send/inbox remain stub-era.
    v0.4.0 remains the current public broad local deployability pre-release.
    Gitea remains source of truth; GitHub mirrors are secondary push mirrors.

---

## 8. Next Safest Actions

Recommended next rung:

    v0.4.19 command registry validation / local-help planning

Focus:

    decide whether to add a lightweight registry validator.
    check registry entries against current runner dispatch and Comms command dispatch.
    decide whether registry can generate component command references.
    decide whether `commands.go` should gain short explainer strings beside command registration.
    decide whether registry should become YAML-only, JSON-only, or YAML with generated JSON later.
    keep front-door README changes minimal.
    keep registry as metadata/navigation/claim-boundary layer.
    do not start v0.5.x state/trust/vault/PQ work until command registry validation/local-help planning is settled.

Likely candidates:

    carbonstack/tools/carbonstack-validate registry check profile or script.
    registry/schema notes.
    docs/165-command-registry-validation-local-help-plan-v0.md.
    maybe a generated command reference later, but not necessarily in the next rung.

Avoid next:

    replacing the direct smoke script.
    merging dev-runtime-openmls and dev-runtime-openmls-wrappers.
    adding either runtime profile to full.
    calling either runtime profile local-backbone.
    treating wrapper bootstrap as Relay Space join UX.
    moving into PQ/vault work before state/trust/storage preflight and registry validation work.
    public ingress.
    cloudflared.
    systemd.
    real homelab deployment.
    Android app.
    CarbonStackOS implementation.
    production/security claims.

---

## 9. Critical Paths / Functions

### Source-of-truth public surfaces

    [GITEA SOURCE OF TRUTH] https://git.bitcrusher32.win/bitcrusher32/carbonstack
    [RELEASES] https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases
    [CURRENT PUBLIC RELEASE] https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.4.0

GitHub mirrors:

    secondary push mirrors for discoverability/redundancy only.
    Do not treat mirrors as release authority unless project policy changes.

### WSL Debian active working umbrella

    [WSL-DEBIAN] umbrella: ~/repos/carbonstack_umbrella
    [WSL-DEBIAN] carbonstack: ~/repos/carbonstack_umbrella/carbonstack
    [WSL-DEBIAN] carbonstack-comms: ~/repos/carbonstack_umbrella/carbonstack-comms
    [WSL-DEBIAN] carbonstack-cypher: ~/repos/carbonstack_umbrella/carbonstack-cypher
    [WSL-DEBIAN] carbonstack-os: ~/repos/carbonstack_umbrella/carbonstack-os
    [GO-RUNNER] runner dir: ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate

### Registry surfaces

    [REGISTRY] root:
    carbonstack/registry

    [REGISTRY] provisional command registry:
    carbonstack/registry/commands.v0.yaml

    [REGISTRY] registry README:
    carbonstack/registry/README.md

    [DOC] v0.4.18 registry checkpoint:
    carbonstack/docs/164-known-good-command-registry-recon-v0.md

### Runtime integration doc surfaces

    [DOC] v0.4.1 runtime recon:
    carbonstack/docs/155-runtime-comms-openmls-cypher-recon-v0.md

    [DOC] v0.4.2 command contract:
    carbonstack/docs/156-runtime-openmls-command-contract-v0.md

    [DOC] v0.4.6 smoke proof record:
    carbonstack/docs/157-dev-runtime-openmls-smoke-proof-v0.md

    [DOC] v0.4.7 pre-local-backbone assessment:
    carbonstack/docs/158-pre-local-backbone-assessment-v0.md

    [DOC] v0.4.9 runner profile:
    carbonstack/docs/159-dev-runtime-openmls-runner-profile-v0.md

    [DOC] v0.4.10 profile boundary:
    carbonstack/docs/160-dev-runtime-openmls-profile-boundary-v0.md

    [DOC] v0.4.11 bootstrap wrapper recon:
    carbonstack/docs/161-openmls-bootstrap-wrapper-recon-v0.md

    [DOC] v0.4.12 bootstrap command contract:
    carbonstack/docs/162-openmls-bootstrap-command-contract-v0.md

    [DOC] v0.4.17 wrapper profile:
    carbonstack/docs/163-dev-runtime-openmls-wrapper-runner-profile-v0.md

    [DOC] docs index:
    carbonstack/docs/README.md

    [DOC] roadmap:
    carbonstack/roadmap/ROADMAP.md

    [DOC] top README:
    carbonstack/README.md

### Runner surfaces

    [RUNNER] root:
    carbonstack/tools/carbonstack-validate

    [RUNNER] dispatch:
    carbonstack/tools/carbonstack-validate/main.go

    [RUNNER] dev-runtime-openmls implementation:
    carbonstack/tools/carbonstack-validate/dev_runtime_openmls.go

    [RUNNER] dev-runtime-openmls-wrappers implementation:
    carbonstack/tools/carbonstack-validate/dev_runtime_openmls_wrappers.go

    [RUNNER] runner README:
    carbonstack/tools/carbonstack-validate/README.md

    [RUNNER] local-cypher:
    carbonstack/tools/carbonstack-validate/local_cypher.go

    [RUNNER] release-snapshot:
    carbonstack/tools/carbonstack-validate/release_snapshot.go

    [RUNNER] checksums:
    carbonstack/tools/carbonstack-validate/checksums.go

### Comms runtime surfaces

    [COMMS] CLI entry:
    carbonstack-comms/cmd/comms/main.go

    [COMMS] command registry and non-OpenMLS commands:
    carbonstack-comms/internal/app/commands.go

    [COMMS] OpenMLS message runtime command helpers:
    carbonstack-comms/internal/app/openmls_runtime.go

    [COMMS] OpenMLS bootstrap wrapper helpers:
    carbonstack-comms/internal/app/openmls_bootstrap.go

    [COMMS] OpenMLS bootstrap wrapper tests:
    carbonstack-comms/internal/app/openmls_bootstrap_dev_test.go

    [COMMS] openmls-send-dev tests:
    carbonstack-comms/internal/app/openmls_send_dev_test.go

    [COMMS] openmls-inbox-dev tests:
    carbonstack-comms/internal/app/openmls_inbox_dev_test.go

    [COMMS] direct-sidecar dev runtime smoke proof:
    carbonstack-comms/scripts/dev-openmls-runtime-smoke.sh

    [COMMS] wrapper-based dev runtime smoke proof:
    carbonstack-comms/scripts/dev-openmls-runtime-smoke-wrappers.sh

    [COMMS] OpenMLS sidecar dir:
    carbonstack-comms/internal/protocol/mls/openmls-sidecar

    [COMMS] sidecar README:
    carbonstack-comms/internal/protocol/mls/openmls-sidecar/README.md

### Current stub-era Comms functions / behaviors

    send:
      trust policy evaluation
      mock/stub provider encryption
      Cypher envelope submit
      registry maturity: legacy

    inbox:
      Cypher inbox retrieval
      mock/stub provider plaintext display as stub_plaintext
      registry maturity: legacy

    ack:
      envelope ack by envelope ID and recipient device ID
      registry maturity: legacy

### Implemented v0.4.x dev OpenMLS commands

    openmls-send-dev:
      internal/app/openmls_runtime.go
      state.RequireReadyDevice
      trust.EvaluateSend
      runOpenMLSMessageProtect
      sidecar message-protect
      internal/relay.SubmitOpenMLSArtifactEnvelope
      ArtifactKindApplicationMessage
      dev/pre-alpha warning
      registry maturity: dev_only

    openmls-inbox-dev:
      internal/app/openmls_runtime.go
      state.RequireReadyDevice
      Cypher inbox
      skip unsupported envelopes
      internal/relay.WriteOpenMLSArtifactFromEnvelope
      runOpenMLSMessageOpen
      sidecar message-open
      print plaintext after sidecar success
      optional ack only after sidecar success
      registry maturity: dev_only

    openmls-identity-create-dev:
      internal/app/openmls_bootstrap.go
      sidecar identity-create
      required --sidecar-device-label
      optional --sidecar-dir
      stable key/value output
      no Comms state/trust mutation
      registry maturity: dev_only

    openmls-identity-status-dev:
      internal/app/openmls_bootstrap.go
      sidecar identity-status
      required --sidecar-device-label
      optional --sidecar-dir
      stable key/value output
      prints identity_exists when present
      no Comms state/trust mutation
      registry maturity: dev_only

    openmls-bundle-export-dev:
      internal/app/openmls_bootstrap.go
      sidecar public-bundle-export
      required --sidecar-device-label
      optional --sidecar-dir
      optional --write-artifact
      stable key/value output
      prints key_package_artifact_path_hint when present
      prints key_package_artifact_path when a hint is present
      no Comms state/trust mutation
      registry maturity: dev_only

    openmls-conversation-create-dev:
      internal/app/openmls_bootstrap.go
      sidecar conversation-create
      required --sidecar-device-label
      required --conversation
      stable key/value output
      prints sidecar_conversation_label
      prints conversation_state_path_hint when present
      prints conversation_summary_path_hint when present
      prints provider_storage_path_hint when present
      no Comms state/trust mutation
      registry maturity: dev_only

    openmls-conversation-load-check-dev:
      internal/app/openmls_bootstrap.go
      sidecar conversation-load-check
      required --sidecar-device-label
      required --conversation
      stable key/value output
      prints group_reloadable when present
      no Comms state/trust mutation
      registry maturity: dev_only

    openmls-conversation-add-member-dev:
      internal/app/openmls_bootstrap.go
      sidecar conversation-add-member
      required --sidecar-device-label
      required --conversation
      required --member-keypackage
      resolves --member-keypackage to absolute path before sidecar call
      stable key/value output
      prints welcome_artifact_path_hint and welcome_artifact_path when present
      prints welcome_manifest_path_hint and welcome_manifest_path when present
      prints welcome_artifact_sha256 and welcome_artifact_size_bytes when present
      prints member_added, welcome_artifact_written, group_reloadable, member counts, and epochs when present
      no Comms state/trust mutation
      registry maturity: dev_only

    openmls-conversation-join-dev:
      internal/app/openmls_bootstrap.go
      sidecar conversation-join
      required --sidecar-device-label
      required --conversation
      required --welcome
      resolves --welcome to absolute path before sidecar call
      stable key/value output
      prints joined, group_reloadable, member_count, epoch, join summary hint, conversation state hint, conversation summary hint, and provider storage hint when present
      no Comms state/trust mutation
      registry maturity: dev_only

### Runtime smoke proofs after v0.4.18

    scripts/dev-openmls-runtime-smoke.sh:
      temp workspace
      temp Cypher binary
      temp Cypher DB
      dynamic loopback Cypher port
      temp Alice/Bob Comms state
      sidecar state reset
      direct sidecar identity/conversation bootstrap
      openmls-send-dev application-message send
      openmls-inbox-dev application-message open + ack
      inbox-empty-after-ack verification
      robust cleanup with INT -> TERM -> KILL escalation

    scripts/dev-openmls-runtime-smoke-wrappers.sh:
      temp workspace
      temp Cypher binary
      temp Cypher DB
      dynamic loopback Cypher port
      temp Alice/Bob Comms state
      sidecar state reset
      Comms wrapper-based identity/conversation bootstrap
      openmls-send-dev application-message send
      openmls-inbox-dev application-message open + ack
      inbox-empty-after-ack verification
      robust cleanup with INT -> TERM -> KILL escalation

### v0.4.18 runner profiles

    go run . --profile dev-runtime-openmls --clean-generated

    Runner.DevRuntimeOpenMLS:
      CheckRequiredPaths
      CheckLiveGitUmbrella
      toolchain reporting
      ArtifactScan("pre-dev-runtime-openmls")
      RunStep for carbonstack-comms/scripts/dev-openmls-runtime-smoke.sh
      ArtifactScan("post-dev-runtime-openmls")
      explicit nonclaim summary

    go run . --profile dev-runtime-openmls-wrappers --clean-generated

    Runner.DevRuntimeOpenMLSWrappers:
      CheckRequiredPaths
      CheckLiveGitUmbrella
      toolchain reporting
      ArtifactScan("pre-dev-runtime-openmls-wrappers")
      RunStep for carbonstack-comms/scripts/dev-openmls-runtime-smoke-wrappers.sh
      ArtifactScan("post-dev-runtime-openmls-wrappers")
      explicit nonclaim summary

### Existing v0.4.0 package validation

    cd <package-root>/carbonstack/tools/carbonstack-validate
    go run . --profile verify-checksums --root <package-root>
    go run . --profile full --root <package-root> --clean-generated

---

## 10. v0.4.x Forward Plan

### Immediate next: v0.4.19 command registry validation / local-help planning

Expected:

    inspect whether registry entries can be validated against runner dispatch and Comms command dispatch.
    decide whether a registry validator should live in tools/carbonstack-validate or as a registry-local helper.
    decide whether registry should be allowed to generate Markdown command references later.
    decide whether local command help should use short explainer strings stored beside command registration.
    decide if registry should remain YAML-only for now or gain generated JSON later.
    keep the front-door README from becoming a command encyclopedia.
    avoid state/trust/vault/PQ implementation until registry validation/local-help planning is done.

### v0.4.20+

Possible:

    implement registry validator or schema check.
    add short help/manual metadata near Comms command registration.
    generate or hand-maintain a command reference doc.
    exact sidecar setup runbook.
    clearer dev state reset semantics.
    better command output shape.
    decision on whether stub-era send/inbox are aliased, replaced, or deprecated.
    decision on when local-backbone naming stops being premature.

---

## 11. v0.5.x and Later Carry-Forward

### v0.5.x

Local storage, trust state, provider state, vault design, and PQ/hybrid ciphersuite migration readiness.

Do not start PQ/hybrid implementation before the runtime OpenMLS/Cypher path exists and is documented. PQ/hybrid migration touches provider state, ciphersuite metadata, trust display, artifact size, reinit/rekey/recovery policy, and claim boundaries.

### v0.6.x+

Hostile-server and abuse-resistance harnesses.

### v0.7.x+

Deployability and operations hardening.

### v0.8.x+

Documented self-pentest / adversarial validation.

### v0.9.x+

Claim-boundary review.

### v0.10.x+

Android backend/app work.

### v1.x.x

Public app/server major epoch and much later CarbonStackOS/appliance research, if justified.

---

## 12. Lean Breakpoint Summary

    CarbonStack v0.4.18 is the known-good command registry checkpoint. v0.4.0 remains the current public broad local deployability pre-release. carbonstack is now at a4162d7 docs: add known-good command registry. carbonstack-comms remains at cb4e59d test: add wrapper-based OpenMLS runtime smoke proof. v0.4.18 adds registry/commands.v0.yaml, registry/README.md, docs/164-known-good-command-registry-recon-v0.md, updates docs/README.md, and updates roadmap/ROADMAP.md. The registry is provisional cross-repo metadata covering runner profiles, Comms CLI commands, old stub-era send/inbox/ack, OpenMLS dev runtime commands, OpenMLS bootstrap wrappers, Comms smoke scripts, OpenMLS sidecar commands, Cypher CLI/env/API surfaces, and legacy PowerShell helpers. It includes maturity, audience, validation surface, short help, why-exists, nonclaims, replacement/deprecation status, and include_in_front_readme fields. Marker validation passed with 62 entries. Runner validation passed for dev-runtime-openmls-wrappers --clean-generated, dev-runtime-openmls --clean-generated, local-cypher, doctor, and core --clean-generated. Final repo heads were clean across all four repos. Next safe action: v0.4.19 command registry validation / local-help planning.

---

## 13. Preserved Immediate Previous Handoff: v0.4.17

The following is the previous v0.4.17 handoff. Where it conflicts with the v0.4.18 overlay above, v0.4.18 wins for current state.



# CarbonStack LogDoc v0.4.17

**Last updated:** 2026-06-05 local session  
**Owner/operator:** bitcrusher32  
**Assistant/model:** GPT-5.5 Thinking  
**Current phase:** Phase 2F / **v0.4.x wrapper OpenMLS runtime validation profile checkpoint complete**. The v0.4.0 public broad local deployability pre-release remains the current public release. v0.4.16 added the parallel wrapper-based OpenMLS runtime smoke script in `carbonstack-comms`; v0.4.17 now promotes that wrapper smoke into a separate manual `carbonstack/tools/carbonstack-validate` profile named `dev-runtime-openmls-wrappers`, while preserving the original `dev-runtime-openmls` direct-smoke baseline, keeping `full` unchanged, and keeping `local-backbone` reserved. `carbonstack` is now at `a501442 (HEAD -> main, origin/main, origin/HEAD) feat: add wrapper OpenMLS runtime validation profile`; `carbonstack-comms` remains at `cb4e59d (HEAD -> main, origin/main, origin/HEAD) test: add wrapper-based OpenMLS runtime smoke proof`; `carbonstack-cypher` remains at `9ab994c (HEAD -> main, origin/main, origin/HEAD) docs: point to local operator runbook`; `carbonstack-os` remains at `1bbbe52 (HEAD -> main, origin/main, origin/HEAD) docs: clarify CarbonStackOS target direction`.

**Version schema:** `v[scope].[timeline]`. This file is `v0.4.17`, the wrapper-smoke runner/profile alignment checkpoint after the v0.4.16 wrapper-based smoke proof checkpoint. Per LogDoc V2 practice, this Markdown file preserves current operational continuity and carries the v0.4.x runtime-integration thread forward; the JSON breakpoint remains the lean machine-readable handoff state.

**Compression note:** The detailed v0.3.0 through v0.3.36PRIME process ledger remains archived in PRIME/sanitized PRIME. v0.4.0 compressed that pre-release sprint; v0.4.1 through v0.4.12 preserved runtime recon, command contracts, dev OpenMLS send/inbox commands, smoke proof, runner profile, boundary checks, wrapper recon, and bootstrap command contract. v0.4.13 through v0.4.15 preserved staged bootstrap wrapper implementation. v0.4.16 preserved the first committed wrapper-based runtime smoke variant. v0.4.17 now preserves the separate manual runner profile for that wrapper smoke.

---

## 1. Project Goal

**Active goal:** Continue moving CarbonStackComms from stub-era runtime `send`/`inbox` behavior toward explicit dev-only OpenMLS/Cypher-backed runtime commands and repeatable validation, while preserving claim discipline. v0.4.17 gives the wrapper-based OpenMLS runtime proof its own separate manual runner profile, without replacing the original direct-sidecar runtime proof or promoting either proof into `full`.

CarbonStack’s broader project intent remains:

- **CarbonStack** is the secure-communications appliance-stack project and shared doctrine/release authority.
- **CarbonStackComms** is the text-first messaging client and current OpenMLS sidecar/relay integration surface.
- **CarbonStackCypher** is the self-hostable relay/storage stack.
- **CarbonStackOS** is future constrained Android-derived appliance OS work and remains out of current runnable validation packages.
- The `carbonstack` Gitea repo remains the source of truth for release state, release assets, docs, and public claims.
- GitHub mirrors exist as secondary push mirrors only; Gitea remains authoritative.

**North star:**

> Its goal is not to be a better smartphone. Its goal is to stop being a smartphone.

**Core doctrine:**

> Every feature is guilty until it proves it does not add unacceptable parser, network, sensor, identity, or filesystem authority.

**Current reality after v0.4.17:**

    openmls-send-dev is a dev-only OpenMLS send-side command.
    openmls-inbox-dev is a dev-only OpenMLS receive/open/optional-ack command.
    scripts/dev-openmls-runtime-smoke.sh remains the original direct-sidecar bootstrap smoke proof:
      direct sidecar bootstrap -> openmls-send-dev -> Cypher -> openmls-inbox-dev --ack
    scripts/dev-openmls-runtime-smoke-wrappers.sh remains the wrapper-based smoke proof:
      openmls-*-dev bootstrap wrappers -> openmls-send-dev -> Cypher -> openmls-inbox-dev --ack

    tools/carbonstack-validate now has two separate manual runtime profiles:
      dev-runtime-openmls
      dev-runtime-openmls-wrappers

    dev-runtime-openmls remains the direct-sidecar smoke baseline.
    dev-runtime-openmls-wrappers is the wrapper-smoke maturity surface.
    Both runtime profiles are manual-only and live-umbrella-only.
    Neither runtime profile is included in full.
    full remains release-snapshot + local-cypher.
    The original direct smoke script remains unchanged.
    The wrapper smoke script remains a parallel proof, not a replacement.

    The planned dev-only OpenMLS bootstrap wrapper set exists:
      openmls-identity-create-dev
      openmls-identity-status-dev
      openmls-bundle-export-dev
      openmls-conversation-create-dev
      openmls-conversation-load-check-dev
      openmls-conversation-add-member-dev
      openmls-conversation-join-dev

    Bootstrap wrappers live in carbonstack-comms/internal/app/openmls_bootstrap.go.
    Bootstrap wrapper tests live in carbonstack-comms/internal/app/openmls_bootstrap_dev_test.go.
    Wrapper-based smoke uses those wrappers for identity, bundle export, conversation create/load-check, add-member / Welcome generation, and join / Welcome consumption.
    Runtime send/open still flows through openmls-send-dev and openmls-inbox-dev.
    Existing send/inbox remain stub-era.
    local-backbone remains reserved.

Hard nonclaims remain: v0.4.17 does **not** prove production readiness, production E2EE, hostile-server safety, metadata privacy, Android readiness, CarbonStackOS readiness, secure vault/storage, public ingress safety, real homelab validation, audit, certification, PQ/hybrid security, mature messenger UX, release-package runtime validation, Relay Space join/onboarding, `full` runtime OpenMLS validation, or local-backbone readiness.

---

## 2. Current Public Release State

Public release:

    https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.4.0

Release title:

    CarbonStack v0.4.0 Broad Local Deployability Pre-Release

Release tag:

    v0.4.0

Release commit:

    24ac9fc494 / 24ac9fc docs: record v0.4.0 package rehearsal

Gitea release status:

    Pre-Release

Important release-page warning remains:

    The default Gitea Source Code ZIP/TAR.GZ downloads are auto-generated carbonstack-only archives.
    The intended runnable validation artifact is the attached multi-repo release package plus companion assets.

Post-release public surface:

    carbonstack README is evergreen.
    It points interested parties to the Releases page for runnable / known-good artifacts.
    Gitea is source of truth.
    GitHub mirrors are secondary push mirrors, not release authority.

---

## 3. Current Repo Heads

    carbonstack        a501442 (HEAD -> main, origin/main, origin/HEAD) feat: add wrapper OpenMLS runtime validation profile
    carbonstack-comms  cb4e59d (HEAD -> main, origin/main, origin/HEAD) test: add wrapper-based OpenMLS runtime smoke proof
    carbonstack-cypher 9ab994c (HEAD -> main, origin/main, origin/HEAD) docs: point to local operator runbook
    carbonstack-os     1bbbe52 (HEAD -> main, origin/main, origin/HEAD) docs: clarify CarbonStackOS target direction

Current public release tag:

    v0.4.0 -> 24ac9fc494 / 24ac9fc docs: record v0.4.0 package rehearsal

Recent runtime/thread commits:

    v0.4.8 extraction:
      0a48ae1 refactor: split OpenMLS runtime command helpers

    v0.4.9 runner profile:
      8eeadb2 feat: add dev OpenMLS runtime validation profile

    v0.4.10 profile boundary:
      7384f43 docs: record dev OpenMLS runtime profile boundary

    v0.4.11 wrapper recon:
      8ee3ebb docs: plan dev OpenMLS bootstrap wrappers

    v0.4.12 command contract:
      0583683 docs: define dev OpenMLS bootstrap command contract

    v0.4.13 identity wrapper implementation:
      20c4cc9 feat: add dev OpenMLS identity bootstrap commands

    v0.4.14 bundle/conversation wrapper implementation:
      1f5c09e feat: add dev OpenMLS bundle and conversation bootstrap commands

    v0.4.15 add-member/join wrapper implementation:
      c2d3b16 feat: add dev OpenMLS add-member and join bootstrap commands

    v0.4.16 wrapper-based smoke proof:
      cb4e59d test: add wrapper-based OpenMLS runtime smoke proof

    Current v0.4.17 wrapper-smoke runner/profile alignment:
      a501442 feat: add wrapper OpenMLS runtime validation profile

Continuity note:

    `a501442` is a carbonstack runner/docs/roadmap commit after the v0.4.0 public release tag.
    It does not retag v0.4.0.
    The release remains attached to `24ac9fc494`.

---

## 4. Current Validated / Known State

Validated / known at the v0.4.17 checkpoint:

    [RELEASE] v0.4.0 Gitea release exists, is marked Pre-Release, and is attached to 24ac9fc494.
    [RUNNER] full means release-snapshot followed by local-cypher.
    [RUNNER] dev-runtime-openmls validates the original direct-sidecar dev/pre-alpha OpenMLS application-message runtime CLI smoke proof.
    [RUNNER] dev-runtime-openmls-wrappers now validates the wrapper-based dev/pre-alpha OpenMLS application-message runtime CLI smoke proof.
    [RUNNER] both runtime profiles remain manual-only and live-umbrella-only.
    [RUNNER] neither runtime profile is included in full.
    [COMMS] openmls-send-dev exists.
    [COMMS] openmls-inbox-dev exists.
    [COMMS] all planned dev-only OpenMLS bootstrap wrappers exist.
    [COMMS] scripts/dev-openmls-runtime-smoke.sh still exists and remains the direct-sidecar smoke proof.
    [COMMS] scripts/dev-openmls-runtime-smoke-wrappers.sh exists and remains the wrapper-based smoke proof.
    [DOCS] docs/163-dev-runtime-openmls-wrapper-runner-profile-v0.md records the new profile and boundary.
    [DOCS] docs/README.md indexes the v0.4.17 profile document.
    [DOCS] roadmap/ROADMAP.md records the v0.4.17 wrapper validation profile.
    [MIRRORS] all four repos have GitHub push-mirror functionality, with Gitea remaining source of truth.

v0.4.17 validation result from the final log:

    go test ./internal/app -count=1 passed in carbonstack-comms.
    go test ./... -count=1 -timeout 600s passed in carbonstack-comms.
    scripts/dev-openmls-runtime-smoke-wrappers.sh passed.
    scripts/dev-openmls-runtime-smoke.sh passed.
    go test ./... -count=1 passed in carbonstack/tools/carbonstack-validate.
    go run . --profile dev-runtime-openmls-wrappers --clean-generated passed.
    go run . --profile dev-runtime-openmls --clean-generated passed.
    local-cypher passed.
    doctor passed.
    core --clean-generated passed.
    carbonstack was committed and pushed at:
      a501442 feat: add wrapper OpenMLS runtime validation profile
    final repo status showed all four repos clean.

Expected live umbrella validation commands now include:

    cd ~/repos/carbonstack_umbrella/carbonstack-comms
    go test ./internal/app -count=1
    go test ./... -count=1 -timeout 600s
    scripts/dev-openmls-runtime-smoke-wrappers.sh
    scripts/dev-openmls-runtime-smoke.sh

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go run . --profile dev-runtime-openmls-wrappers --clean-generated
    go run . --profile dev-runtime-openmls --clean-generated
    go run . --profile local-cypher
    go run . --profile doctor
    go run . --profile core --clean-generated

Expected v0.4.0 public package validation commands remain unchanged:

    cd <package-root>/carbonstack/tools/carbonstack-validate
    go run . --profile verify-checksums --root <package-root>
    go run . --profile full --root <package-root> --clean-generated

---

## 5. v0.4.17 Work Completed

### 5.1 Separate wrapper-smoke runner profile was added

New runner file:

    carbonstack/tools/carbonstack-validate/dev_runtime_openmls_wrappers.go

New profile:

    dev-runtime-openmls-wrappers

Command:

    go run . --profile dev-runtime-openmls-wrappers --clean-generated

Wrapped script:

    carbonstack-comms/scripts/dev-openmls-runtime-smoke-wrappers.sh

Proof shape:

    openmls-*-dev bootstrap wrappers -> openmls-send-dev -> Cypher -> openmls-inbox-dev --ack

The profile follows the same broad guard pattern as `dev-runtime-openmls`:

    required path checks
    live git umbrella guard
    toolchain reporting
    pre-profile artifact scan
    wrapper smoke script execution
    post-profile artifact scan
    explicit nonclaim summary

Profile boundary:

    manual-only.
    live-umbrella-only.
    not included in full.
    not release-package validation yet.
    not local-backbone.
    not deployment.
    not production/security proof.
    separate wrapper-smoke maturity surface.
    does not replace dev-runtime-openmls yet.

### 5.2 Existing direct-smoke runner profile was preserved

Existing profile remains:

    dev-runtime-openmls

Existing runner file remains:

    carbonstack/tools/carbonstack-validate/dev_runtime_openmls.go

It continues to wrap:

    carbonstack-comms/scripts/dev-openmls-runtime-smoke.sh

Correct interpretation:

    dev-runtime-openmls remains the direct-sidecar smoke baseline.
    dev-runtime-openmls-wrappers is the wrapper-smoke maturity surface.
    the two profiles are separate while wrappers mature.
    they may be merged later only after wrapper behavior, reset semantics, runner semantics, and docs justify that maturity step.

### 5.3 Runner dispatch and README were updated

Updated runner files:

    carbonstack/tools/carbonstack-validate/main.go
    carbonstack/tools/carbonstack-validate/README.md

`main.go` now accepts and dispatches:

    dev-runtime-openmls-wrappers

The runner README now documents:

    dev-runtime-openmls-wrappers
    its proof shape
    its wrapped script
    its manual-only and live-umbrella-only status
    its nonclaims
    the fact that it is not included in full
    the fact that it does not replace dev-runtime-openmls yet

### 5.4 Main docs and roadmap were updated

New doc:

    carbonstack/docs/163-dev-runtime-openmls-wrapper-runner-profile-v0.md

Updated docs index:

    carbonstack/docs/README.md

Updated roadmap:

    carbonstack/roadmap/ROADMAP.md

The new doc records:

    why `dev-runtime-openmls-wrappers` exists
    why it is separate from `dev-runtime-openmls`
    why it is not local-backbone
    why it is not release-package validation
    validation expectation
    suggested next rung: pre-v0.5.x known-good command registry recon/planning

### 5.5 Validation passed

Commands observed passing:

    cd ~/repos/carbonstack_umbrella/carbonstack-comms
    go test ./internal/app -count=1
    go test ./... -count=1 -timeout 600s
    scripts/dev-openmls-runtime-smoke-wrappers.sh
    scripts/dev-openmls-runtime-smoke.sh

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go test ./... -count=1
    go run . --profile dev-runtime-openmls-wrappers --clean-generated
    go run . --profile dev-runtime-openmls --clean-generated
    go run . --profile local-cypher
    go run . --profile doctor
    go run . --profile core --clean-generated

Final commit/push:

    carbonstack:
    a501442 feat: add wrapper OpenMLS runtime validation profile

Final repo snapshot:

    carbonstack        a501442 feat: add wrapper OpenMLS runtime validation profile
    carbonstack-comms  cb4e59d test: add wrapper-based OpenMLS runtime smoke proof
    carbonstack-cypher 9ab994c docs: point to local operator runbook
    carbonstack-os     1bbbe52 docs: clarify CarbonStackOS target direction

---

## 6. Critical Blunders / Lessons to Carry Forward

### 6.1 Recon pasteback was useful but still partially glitched

The v0.4.17 recon pasteback again showed command-tail corruption near the runner/docs section. It still included the useful source surfaces: current heads, both smoke scripts, runner dispatch, existing `dev-runtime-openmls` implementation, runner README boundary, and main docs/roadmap references.

Lesson:

    Keep using recon -> pasteback -> implement.
    Keep splitting large work into smaller paste blocks.
    Do not trust a glitched recon transcript as complete validation evidence.
    Final validation and final repo snapshots are the authority.

### 6.2 Separate profile avoids premature replacement

v0.4.17 deliberately does not replace `dev-runtime-openmls`.

Lesson:

    The direct-sidecar smoke remains a lower-level control/baseline.
    The wrapper-smoke profile is a maturity surface.
    Do not merge or replace profiles until wrapper behavior, reset semantics, runner semantics, and docs justify it.

### 6.3 Runtime validation remains live-umbrella-only

Both runtime profiles intentionally require sibling git checkouts.

Lesson:

    Do not treat either runtime profile as release-package validation.
    Do not add either runtime profile to full yet.
    Fresh-root behavior for runtime profiles remains future work.

### 6.4 Generated roots remain expected after smoke/profile runs

Both direct and wrapper runtime profiles can create:

    .carbonstack-openmls-sidecar-state
    target

This remains acceptable only when:

    the paths are known generated roots.
    artifact scans classify them correctly.
    --clean-generated removes them before a clean breakpoint.

### 6.5 Commit metadata warning remains a recurring polish issue

The commit log again showed auto-generated local commit metadata:

    Committer: bitcrusher32 <▓▓>

This is not a functional validation failure, but remains a professionalism/publishing hygiene issue to address eventually.

---

## 7. Current Blockers / Not Validated

Current blockers and nonvalidated areas after v0.4.17:

    Existing send/inbox remain stub-era.
    Runtime Comms has dev-only OpenMLS commands, direct smoke proof, wrapper-based smoke proof, and separate manual runner profiles, but no mature user-facing UX.
    Neither runtime profile is included in full.
    Neither runtime profile is part of public release package validation.
    Fresh release-package behavior for dev-runtime-openmls or dev-runtime-openmls-wrappers has not been tested as a supported validation surface.
    No Relay Space join model exists yet.
    local-backbone runner profile does not exist and remains reserved.
    Post-release public download verification is still useful if not already performed.
    Sanitized PRIME placement in carbonstack/logdoc-loop-system still needs a policy decision.
    Production-safe provider storage and secure vault design remain v0.5.x future work.
    PQ/hybrid ciphersuite migration readiness remains v0.5.x future work.
    Hostile-server rollback/replay/metadata-abuse validation remains future work.
    Android/CarbonStackOS implementation remains later work.
    CarbonStack Relay Space mechanics remain a future concept, not an implementation.
    Known-good command registry YAML/JSON does not exist yet.
    README/dev command sprawl remains a known future documentation risk.

Hard nonclaims remain:

    Do not claim production readiness.
    Do not claim production E2EE.
    Do not claim hostile-server safety.
    Do not claim metadata privacy.
    Do not claim Debian production deployability.
    Do not claim systemd readiness.
    Do not claim cloudflared readiness.
    Do not claim Android readiness.
    Do not claim CarbonStackOS readiness.
    Do not claim secure vault/storage.
    Do not claim real homelab validation.
    Do not claim external audit or certification.
    Do not claim CarbonStack Relay Space mechanics are implemented.
    Do not claim local-backbone exists or is justified yet.
    Do not claim local-cypher validates runtime Comms UX.
    Do not claim full validates local-backbone.
    Do not claim full is a deployment command.
    Do not claim dev-runtime-openmls is release-package validation yet.
    Do not claim dev-runtime-openmls-wrappers is release-package validation yet.
    Do not claim send/inbox are OpenMLS-backed yet.
    Do not claim PQ/hybrid security or quantum-safe messaging.
    Do not let GitHub mirrors supersede Gitea source-of-truth status.

Allowed claim:

    CarbonStackComms now has dev-only openmls-send-dev and openmls-inbox-dev commands.
    CarbonStackComms now has a direct-sidecar dev runtime smoke script proving direct sidecar bootstrap -> openmls-send-dev -> Cypher -> openmls-inbox-dev --ack.
    CarbonStackComms now has a wrapper-based dev runtime smoke script proving openmls-*-dev bootstrap wrappers -> openmls-send-dev -> Cypher -> openmls-inbox-dev --ack.
    CarbonStackComms now has OpenMLS runtime command helper glue split into internal/app/openmls_runtime.go.
    CarbonStackComms now has dev-only OpenMLS bootstrap wrappers for identity create/status, bundle export, conversation create/load-check, conversation add-member, and conversation join.
    The bootstrap wrappers parse sidecar JSON, print stable key/value output, keep sidecar labels explicit, resolve path-consuming input artifacts to absolute paths, and do not mutate Comms state/trust files.
    carbonstack/tools/carbonstack-validate has a manual dev-runtime-openmls profile wrapping the direct-sidecar smoke.
    carbonstack/tools/carbonstack-validate now has a separate manual dev-runtime-openmls-wrappers profile wrapping the wrapper-based smoke.
    Both runtime profiles are live-git-umbrella-only and not included in full.
    Existing send/inbox remain stub-era.
    v0.4.0 remains the current public broad local deployability pre-release.
    Gitea remains source of truth; GitHub mirrors are secondary push mirrors.

---

## 8. Next Safest Actions

Recommended next rung:

    v0.4.18 pre-v0.5.x known-good command registry recon/planning

Focus:

    scout existing commands across carbonstack, carbonstack-comms, carbonstack-cypher, and carbonstack-os.
    scout current runner profiles and smoke scripts.
    decide registry location and schema.
    keep the front-door README from becoming a command encyclopedia.
    distinguish user-facing, operator-facing, dev-only, internal, script, and runner commands.
    include maturity, repo, validation surface, why it exists, and nonclaims.
    do not start v0.5.x state/trust/vault/PQ work until command registry planning is done.

Likely registry location candidates:

    carbonstack/docs/command-registry.yaml
    carbonstack/docs/command-registry.json
    carbonstack/docs/164-known-good-command-registry-recon-v0.md
    carbonstack/docs/command-reference-v0.md

Registry should probably be machine-readable plus a short doc, not README-only.

Avoid next:

    replacing the direct smoke script.
    merging dev-runtime-openmls and dev-runtime-openmls-wrappers.
    adding either runtime profile to full.
    calling either runtime profile local-backbone.
    treating wrapper bootstrap as Relay Space join UX.
    moving into PQ/vault work before state/trust/storage preflight and registry work.
    public ingress.
    cloudflared.
    systemd.
    real homelab deployment.
    Android app.
    CarbonStackOS implementation.
    production/security claims.

---

## 9. Critical Paths / Functions

### Source-of-truth public surfaces

    [GITEA SOURCE OF TRUTH] https://git.bitcrusher32.win/bitcrusher32/carbonstack
    [RELEASES] https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases
    [CURRENT PUBLIC RELEASE] https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.4.0

GitHub mirrors:

    secondary push mirrors for discoverability/redundancy only.
    Do not treat mirrors as release authority unless project policy changes.

### WSL Debian active working umbrella

    [WSL-DEBIAN] umbrella: ~/repos/carbonstack_umbrella
    [WSL-DEBIAN] carbonstack: ~/repos/carbonstack_umbrella/carbonstack
    [WSL-DEBIAN] carbonstack-comms: ~/repos/carbonstack_umbrella/carbonstack-comms
    [WSL-DEBIAN] carbonstack-cypher: ~/repos/carbonstack_umbrella/carbonstack-cypher
    [WSL-DEBIAN] carbonstack-os: ~/repos/carbonstack_umbrella/carbonstack-os
    [GO-RUNNER] runner dir: ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate

### Runtime integration doc surfaces

    [DOC] v0.4.1 runtime recon:
    carbonstack/docs/155-runtime-comms-openmls-cypher-recon-v0.md

    [DOC] v0.4.2 command contract:
    carbonstack/docs/156-runtime-openmls-command-contract-v0.md

    [DOC] v0.4.6 smoke proof record:
    carbonstack/docs/157-dev-runtime-openmls-smoke-proof-v0.md

    [DOC] v0.4.7 pre-local-backbone assessment:
    carbonstack/docs/158-pre-local-backbone-assessment-v0.md

    [DOC] v0.4.9 runner profile:
    carbonstack/docs/159-dev-runtime-openmls-runner-profile-v0.md

    [DOC] v0.4.10 profile boundary:
    carbonstack/docs/160-dev-runtime-openmls-profile-boundary-v0.md

    [DOC] v0.4.11 bootstrap wrapper recon:
    carbonstack/docs/161-openmls-bootstrap-wrapper-recon-v0.md

    [DOC] v0.4.12 bootstrap command contract:
    carbonstack/docs/162-openmls-bootstrap-command-contract-v0.md

    [DOC] v0.4.17 wrapper profile:
    carbonstack/docs/163-dev-runtime-openmls-wrapper-runner-profile-v0.md

    [DOC] roadmap refresh:
    CarbonStack_Long_Term_Roadmap_v0.4.14_REFRESH.pdf

    [DOC] docs index:
    carbonstack/docs/README.md

    [DOC] roadmap:
    carbonstack/roadmap/ROADMAP.md

    [DOC] top README:
    carbonstack/README.md

### Runner surfaces

    [RUNNER] root:
    carbonstack/tools/carbonstack-validate

    [RUNNER] dispatch:
    carbonstack/tools/carbonstack-validate/main.go

    [RUNNER] dev-runtime-openmls implementation:
    carbonstack/tools/carbonstack-validate/dev_runtime_openmls.go

    [RUNNER] dev-runtime-openmls-wrappers implementation:
    carbonstack/tools/carbonstack-validate/dev_runtime_openmls_wrappers.go

    [RUNNER] runner README:
    carbonstack/tools/carbonstack-validate/README.md

    [RUNNER] local-cypher:
    carbonstack/tools/carbonstack-validate/local_cypher.go

    [RUNNER] release-snapshot:
    carbonstack/tools/carbonstack-validate/release_snapshot.go

    [RUNNER] checksums:
    carbonstack/tools/carbonstack-validate/checksums.go

### Comms runtime surfaces

    [COMMS] CLI entry:
    carbonstack-comms/cmd/comms/main.go

    [COMMS] command registry and non-OpenMLS commands:
    carbonstack-comms/internal/app/commands.go

    [COMMS] OpenMLS message runtime command helpers:
    carbonstack-comms/internal/app/openmls_runtime.go

    [COMMS] OpenMLS bootstrap wrapper helpers:
    carbonstack-comms/internal/app/openmls_bootstrap.go

    [COMMS] OpenMLS bootstrap wrapper tests:
    carbonstack-comms/internal/app/openmls_bootstrap_dev_test.go

    [COMMS] openmls-send-dev tests:
    carbonstack-comms/internal/app/openmls_send_dev_test.go

    [COMMS] openmls-inbox-dev tests:
    carbonstack-comms/internal/app/openmls_inbox_dev_test.go

    [COMMS] direct-sidecar dev runtime smoke proof:
    carbonstack-comms/scripts/dev-openmls-runtime-smoke.sh

    [COMMS] wrapper-based dev runtime smoke proof:
    carbonstack-comms/scripts/dev-openmls-runtime-smoke-wrappers.sh

    [COMMS] OpenMLS sidecar dir:
    carbonstack-comms/internal/protocol/mls/openmls-sidecar

    [COMMS] sidecar README:
    carbonstack-comms/internal/protocol/mls/openmls-sidecar/README.md

### Current stub-era Comms functions / behaviors

    send:
      trust policy evaluation
      mock/stub provider encryption
      Cypher envelope submit

    inbox:
      Cypher inbox retrieval
      mock/stub provider plaintext display as stub_plaintext

    ack:
      envelope ack by envelope ID and recipient device ID

### Implemented v0.4.x dev OpenMLS commands

    openmls-send-dev:
      internal/app/openmls_runtime.go
      state.RequireReadyDevice
      trust.EvaluateSend
      runOpenMLSMessageProtect
      sidecar message-protect
      internal/relay.SubmitOpenMLSArtifactEnvelope
      ArtifactKindApplicationMessage
      dev/pre-alpha warning

    openmls-inbox-dev:
      internal/app/openmls_runtime.go
      state.RequireReadyDevice
      Cypher inbox
      skip unsupported envelopes
      internal/relay.WriteOpenMLSArtifactFromEnvelope
      runOpenMLSMessageOpen
      sidecar message-open
      print plaintext after sidecar success
      optional ack only after sidecar success

    openmls-identity-create-dev:
      internal/app/openmls_bootstrap.go
      sidecar identity-create
      required --sidecar-device-label
      optional --sidecar-dir
      stable key/value output
      no Comms state/trust mutation

    openmls-identity-status-dev:
      internal/app/openmls_bootstrap.go
      sidecar identity-status
      required --sidecar-device-label
      optional --sidecar-dir
      stable key/value output
      prints identity_exists when present
      no Comms state/trust mutation

    openmls-bundle-export-dev:
      internal/app/openmls_bootstrap.go
      sidecar public-bundle-export
      required --sidecar-device-label
      optional --sidecar-dir
      optional --write-artifact
      stable key/value output
      prints key_package_artifact_path_hint when present
      prints key_package_artifact_path when a hint is present
      no Comms state/trust mutation

    openmls-conversation-create-dev:
      internal/app/openmls_bootstrap.go
      sidecar conversation-create
      required --sidecar-device-label
      required --conversation
      stable key/value output
      prints sidecar_conversation_label
      prints conversation_state_path_hint when present
      prints conversation_summary_path_hint when present
      prints provider_storage_path_hint when present
      no Comms state/trust mutation

    openmls-conversation-load-check-dev:
      internal/app/openmls_bootstrap.go
      sidecar conversation-load-check
      required --sidecar-device-label
      required --conversation
      stable key/value output
      prints group_reloadable when present
      no Comms state/trust mutation

    openmls-conversation-add-member-dev:
      internal/app/openmls_bootstrap.go
      sidecar conversation-add-member
      required --sidecar-device-label
      required --conversation
      required --member-keypackage
      resolves --member-keypackage to absolute path before sidecar call
      stable key/value output
      prints welcome_artifact_path_hint and welcome_artifact_path when present
      prints welcome_manifest_path_hint and welcome_manifest_path when present
      prints welcome_artifact_sha256 and welcome_artifact_size_bytes when present
      prints member_added, welcome_artifact_written, group_reloadable, member counts, and epochs when present
      no Comms state/trust mutation

    openmls-conversation-join-dev:
      internal/app/openmls_bootstrap.go
      sidecar conversation-join
      required --sidecar-device-label
      required --conversation
      required --welcome
      resolves --welcome to absolute path before sidecar call
      stable key/value output
      prints joined, group_reloadable, member_count, epoch, join summary hint, conversation state hint, conversation summary hint, and provider storage hint when present
      no Comms state/trust mutation

### Runtime smoke proofs after v0.4.17

    scripts/dev-openmls-runtime-smoke.sh:
      temp workspace
      temp Cypher binary
      temp Cypher DB
      dynamic loopback Cypher port
      temp Alice/Bob Comms state
      sidecar state reset
      direct sidecar identity/conversation bootstrap
      openmls-send-dev application-message send
      openmls-inbox-dev application-message open + ack
      inbox-empty-after-ack verification
      robust cleanup with INT -> TERM -> KILL escalation

    scripts/dev-openmls-runtime-smoke-wrappers.sh:
      temp workspace
      temp Cypher binary
      temp Cypher DB
      dynamic loopback Cypher port
      temp Alice/Bob Comms state
      sidecar state reset
      Comms wrapper-based identity/conversation bootstrap
      openmls-send-dev application-message send
      openmls-inbox-dev application-message open + ack
      inbox-empty-after-ack verification
      robust cleanup with INT -> TERM -> KILL escalation

### v0.4.17 runner profiles

    go run . --profile dev-runtime-openmls --clean-generated

    Runner.DevRuntimeOpenMLS:
      CheckRequiredPaths
      CheckLiveGitUmbrella
      toolchain reporting
      ArtifactScan("pre-dev-runtime-openmls")
      RunStep for carbonstack-comms/scripts/dev-openmls-runtime-smoke.sh
      ArtifactScan("post-dev-runtime-openmls")
      explicit nonclaim summary

    go run . --profile dev-runtime-openmls-wrappers --clean-generated

    Runner.DevRuntimeOpenMLSWrappers:
      CheckRequiredPaths
      CheckLiveGitUmbrella
      toolchain reporting
      ArtifactScan("pre-dev-runtime-openmls-wrappers")
      RunStep for carbonstack-comms/scripts/dev-openmls-runtime-smoke-wrappers.sh
      ArtifactScan("post-dev-runtime-openmls-wrappers")
      explicit nonclaim summary

### Existing v0.4.0 package validation

    cd <package-root>/carbonstack/tools/carbonstack-validate
    go run . --profile verify-checksums --root <package-root>
    go run . --profile full --root <package-root> --clean-generated

---

## 10. v0.4.x Forward Plan

### Immediate next: v0.4.18 pre-v0.5.x known-good command registry recon/planning

Expected:

    scout current implemented commands and scripts across all four repos.
    decide a registry schema and location.
    decide whether registry should be YAML, JSON, or both.
    distinguish user-facing, operator-facing, dev-only, internal, script, and runner command surfaces.
    record maturity, repo, validation surface, why it exists, nonclaims, and replacement/deprecation status.
    keep the front-door README from becoming a command encyclopedia.
    avoid state/trust/vault/PQ implementation until registry planning is done.

### v0.4.19+

Possible:

    implement initial command registry.
    add a short command reference / registry docs index.
    update front README to point at registry instead of growing command sections.
    exact sidecar setup runbook.
    clearer dev state reset semantics.
    better command output shape.
    decision on whether stub-era send/inbox are aliased, replaced, or deprecated.
    decision on when local-backbone naming stops being premature.

---

## 11. v0.5.x and Later Carry-Forward

### v0.5.x

Local storage, trust state, provider state, vault design, and PQ/hybrid ciphersuite migration readiness.

Do not start PQ/hybrid implementation before the runtime OpenMLS/Cypher path exists and is documented. PQ/hybrid migration touches provider state, ciphersuite metadata, trust display, artifact size, reinit/rekey/recovery policy, and claim boundaries.

### v0.6.x+

Hostile-server and abuse-resistance harnesses.

### v0.7.x+

Deployability and operations hardening.

### v0.8.x+

Documented self-pentest / adversarial validation.

### v0.9.x+

Claim-boundary review.

### v0.10.x+

Android backend/app work.

### v1.x.x

Public app/server major epoch and much later CarbonStackOS/appliance research, if justified.

---

## 12. Lean Breakpoint Summary

    CarbonStack v0.4.17 is the wrapper OpenMLS runtime validation profile checkpoint. v0.4.0 remains the current public broad local deployability pre-release. carbonstack is now at a501442 feat: add wrapper OpenMLS runtime validation profile. carbonstack-comms remains at cb4e59d test: add wrapper-based OpenMLS runtime smoke proof. v0.4.17 adds tools/carbonstack-validate/dev_runtime_openmls_wrappers.go, updates tools/carbonstack-validate/main.go and README.md, adds docs/163-dev-runtime-openmls-wrapper-runner-profile-v0.md, updates docs/README.md, and updates roadmap/ROADMAP.md. The new manual dev-runtime-openmls-wrappers profile wraps scripts/dev-openmls-runtime-smoke-wrappers.sh and validates openmls-*-dev bootstrap wrappers -> openmls-send-dev -> Cypher -> openmls-inbox-dev --ack. The existing dev-runtime-openmls profile remains unchanged and continues to wrap the direct-sidecar smoke script. full remains unchanged and excludes both runtime profiles. Validation passed for Comms app tests, full Comms tests, wrapper smoke, direct smoke, runner tests, dev-runtime-openmls-wrappers --clean-generated, dev-runtime-openmls --clean-generated, local-cypher, doctor, and core --clean-generated. Final repo heads were clean across all four repos. Next safe action: v0.4.18 pre-v0.5.x known-good command registry recon/planning.

---

## 13. Preserved Immediate Previous Handoff: v0.4.16

The following is the previous v0.4.16 handoff. Where it conflicts with the v0.4.17 overlay above, v0.4.17 wins for current state.



# CarbonStack LogDoc v0.4.16

**Last updated:** 2026-06-05 local session  
**Owner/operator:** bitcrusher32  
**Assistant/model:** GPT-5.5 Thinking  
**Current phase:** Phase 2F / **v0.4.x wrapper-based OpenMLS runtime smoke proof checkpoint complete**. The v0.4.0 public broad local deployability pre-release remains the current public release. v0.4.15 completed the planned dev-only OpenMLS bootstrap wrapper set; v0.4.16 now adds a parallel wrapper-based smoke proof in `carbonstack-comms` while preserving the original direct-sidecar smoke script and keeping `dev-runtime-openmls` / `full` runner semantics unchanged. `carbonstack` remains at `0583683 (HEAD -> main, origin/main, origin/HEAD) docs: define dev OpenMLS bootstrap command contract`; `carbonstack-comms` is now at `cb4e59d (HEAD -> main, origin/main, origin/HEAD) test: add wrapper-based OpenMLS runtime smoke proof`; `carbonstack-cypher` remains at `9ab994c (HEAD -> main, origin/main, origin/HEAD) docs: point to local operator runbook`; `carbonstack-os` remains at `1bbbe52 (HEAD -> main, origin/main, origin/HEAD) docs: clarify CarbonStackOS target direction`.

**Version schema:** `v[scope].[timeline]`. This file is `v0.4.16`, the wrapper-based smoke proof checkpoint after the v0.4.15 full bootstrap-wrapper coverage checkpoint. Per LogDoc V2 practice, this Markdown file preserves current operational continuity and carries the v0.4.x runtime-integration thread forward; the JSON breakpoint remains the lean machine-readable handoff state.

**Compression note:** The detailed v0.3.0 through v0.3.36PRIME process ledger remains archived in PRIME/sanitized PRIME. v0.4.0 compressed that pre-release sprint; v0.4.1 through v0.4.12 preserved runtime recon, command contracts, dev OpenMLS send/inbox commands, smoke proof, runner profile, boundary checks, wrapper recon, and bootstrap command contract. v0.4.13 through v0.4.15 preserved staged bootstrap wrapper implementation. v0.4.16 now preserves the first committed wrapper-based runtime smoke variant.

---

## 1. Project Goal

**Active goal:** Continue moving CarbonStackComms from stub-era runtime `send`/`inbox` behavior toward explicit dev-only OpenMLS/Cypher-backed runtime commands and repeatable validation, while preserving claim discipline. v0.4.16 adds a parallel smoke script that proves the complete `openmls-*-dev` bootstrap wrapper chain can feed the existing runtime message path.

CarbonStack’s broader project intent remains:

- **CarbonStack** is the secure-communications appliance-stack project and shared doctrine/release authority.
- **CarbonStackComms** is the text-first messaging client and current OpenMLS sidecar/relay integration surface.
- **CarbonStackCypher** is the self-hostable relay/storage stack.
- **CarbonStackOS** is future constrained Android-derived appliance OS work and remains out of current runnable validation packages.
- The `carbonstack` Gitea repo remains the source of truth for release state, release assets, docs, and public claims.
- GitHub mirrors exist as secondary push mirrors only; Gitea remains authoritative.

**North star:**

> Its goal is not to be a better smartphone. Its goal is to stop being a smartphone.

**Core doctrine:**

> Every feature is guilty until it proves it does not add unacceptable parser, network, sensor, identity, or filesystem authority.

**Current reality after v0.4.16:**

    openmls-send-dev is a dev-only OpenMLS send-side command.
    openmls-inbox-dev is a dev-only OpenMLS receive/open/optional-ack command.
    scripts/dev-openmls-runtime-smoke.sh remains the original direct-sidecar bootstrap smoke proof:
      direct sidecar bootstrap -> openmls-send-dev -> Cypher -> openmls-inbox-dev --ack
    scripts/dev-openmls-runtime-smoke-wrappers.sh now exists as a parallel wrapper-based smoke proof:
      openmls-*-dev bootstrap wrappers -> openmls-send-dev -> Cypher -> openmls-inbox-dev --ack
    tools/carbonstack-validate still has a manual dev-runtime-openmls profile that wraps the original direct smoke script.
    dev-runtime-openmls remains manual-only, live-git-umbrella-only, and outside full.
    full remains release-snapshot + local-cypher.
    The original direct smoke script remains unchanged.

    The planned dev-only OpenMLS bootstrap wrapper set exists:
      openmls-identity-create-dev
      openmls-identity-status-dev
      openmls-bundle-export-dev
      openmls-conversation-create-dev
      openmls-conversation-load-check-dev
      openmls-conversation-add-member-dev
      openmls-conversation-join-dev

    Bootstrap wrappers live in carbonstack-comms/internal/app/openmls_bootstrap.go.
    Bootstrap wrapper tests live in carbonstack-comms/internal/app/openmls_bootstrap_dev_test.go.
    Wrapper-based smoke uses those wrappers for identity, bundle export, conversation create/load-check, add-member / Welcome generation, and join / Welcome consumption.
    Runtime send/open still flows through openmls-send-dev and openmls-inbox-dev.
    The wrapper-based smoke verifies plaintext match, ack-after-open success, and Bob inbox empty after ack.
    Existing send/inbox remain stub-era.
    local-backbone remains reserved.

Hard nonclaims remain: v0.4.16 does **not** prove production readiness, production E2EE, hostile-server safety, metadata privacy, Android readiness, CarbonStackOS readiness, secure vault/storage, public ingress safety, real homelab validation, audit, certification, PQ/hybrid security, mature messenger UX, release-package runtime validation, Relay Space join/onboarding, `full` runtime OpenMLS validation, or local-backbone readiness.

---

## 2. Current Public Release State

Public release:

    https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.4.0

Release title:

    CarbonStack v0.4.0 Broad Local Deployability Pre-Release

Release tag:

    v0.4.0

Release commit:

    24ac9fc494 / 24ac9fc docs: record v0.4.0 package rehearsal

Gitea release status:

    Pre-Release

Important release-page warning remains:

    The default Gitea Source Code ZIP/TAR.GZ downloads are auto-generated carbonstack-only archives.
    The intended runnable validation artifact is the attached multi-repo release package plus companion assets.

Post-release public surface:

    carbonstack README is evergreen.
    It points interested parties to the Releases page for runnable / known-good artifacts.
    Gitea is source of truth.
    GitHub mirrors are secondary push mirrors, not release authority.

---

## 3. Current Repo Heads

    carbonstack        0583683 (HEAD -> main, origin/main, origin/HEAD) docs: define dev OpenMLS bootstrap command contract
    carbonstack-comms  cb4e59d (HEAD -> main, origin/main, origin/HEAD) test: add wrapper-based OpenMLS runtime smoke proof
    carbonstack-cypher 9ab994c (HEAD -> main, origin/main, origin/HEAD) docs: point to local operator runbook
    carbonstack-os     1bbbe52 (HEAD -> main, origin/main, origin/HEAD) docs: clarify CarbonStackOS target direction

Current public release tag:

    v0.4.0 -> 24ac9fc494 / 24ac9fc docs: record v0.4.0 package rehearsal

Recent runtime/thread commits:

    v0.4.8 extraction:
      0a48ae1 refactor: split OpenMLS runtime command helpers

    v0.4.9 runner profile:
      8eeadb2 feat: add dev OpenMLS runtime validation profile

    v0.4.10 profile boundary:
      7384f43 docs: record dev OpenMLS runtime profile boundary

    v0.4.11 wrapper recon:
      8ee3ebb docs: plan dev OpenMLS bootstrap wrappers

    v0.4.12 command contract:
      0583683 docs: define dev OpenMLS bootstrap command contract

    v0.4.13 identity wrapper implementation:
      20c4cc9 feat: add dev OpenMLS identity bootstrap commands

    v0.4.14 bundle/conversation wrapper implementation:
      1f5c09e feat: add dev OpenMLS bundle and conversation bootstrap commands

    v0.4.15 add-member/join wrapper implementation:
      c2d3b16 feat: add dev OpenMLS add-member and join bootstrap commands

    Current v0.4.16 wrapper-based smoke proof:
      cb4e59d test: add wrapper-based OpenMLS runtime smoke proof

Continuity note:

    `cb4e59d` is a carbonstack-comms test/smoke proof commit after the v0.4.0 public release tag.
    It does not retag v0.4.0.
    The release remains attached to `24ac9fc494`.

---

## 4. Current Validated / Known State

Validated / known at the v0.4.16 checkpoint:

    [RELEASE] v0.4.0 Gitea release exists, is marked Pre-Release, and is attached to 24ac9fc494.
    [RUNNER] full means release-snapshot followed by local-cypher.
    [RUNNER] dev-runtime-openmls validates the original direct-sidecar dev/pre-alpha OpenMLS application-message runtime CLI smoke proof.
    [RUNNER] dev-runtime-openmls remains manual-only and live-umbrella-only.
    [RUNNER] dev-runtime-openmls is not included in full.
    [COMMS] openmls-send-dev exists.
    [COMMS] openmls-inbox-dev exists.
    [COMMS] all planned dev-only OpenMLS bootstrap wrappers exist.
    [COMMS] scripts/dev-openmls-runtime-smoke.sh still exists and remains the direct-sidecar smoke proof.
    [COMMS] scripts/dev-openmls-runtime-smoke-wrappers.sh now exists and is executable.
    [COMMS] README.md now documents the wrapper-based dev runtime OpenMLS smoke proof and boundaries.
    [COMMS] existing send/inbox/ack remain in place and stub-era.
    [MIRRORS] all four repos have GitHub push-mirror functionality, with Gitea remaining source of truth.

v0.4.16 validation result from the final log:

    go test ./internal/app -count=1 passed in carbonstack-comms.
    go test ./... -count=1 -timeout 600s passed in carbonstack-comms.
    scripts/dev-openmls-runtime-smoke-wrappers.sh passed.
    scripts/dev-openmls-runtime-smoke.sh passed.
    go test ./... -count=1 passed in carbonstack/tools/carbonstack-validate.
    go run . --profile dev-runtime-openmls --clean-generated passed.
    local-cypher passed.
    doctor passed.
    core --clean-generated passed.
    carbonstack-comms was committed and pushed at:
      cb4e59d test: add wrapper-based OpenMLS runtime smoke proof
    final repo status showed all four repos clean.

Wrapper-based smoke proof observed:

    proof: openmls-*-dev bootstrap wrappers -> openmls-send-dev -> Cypher -> openmls-inbox-dev --ack
    plaintext: hello bob through wrapper bootstrap runtime smoke
    boundary: dev/pre-alpha wrapper smoke proof; not local-backbone; not production messaging UX

Direct smoke proof also remained passing:

    proof: openmls-send-dev -> Cypher -> openmls-inbox-dev --ack
    plaintext: hello bob through openmls runtime smoke
    boundary: dev/pre-alpha smoke proof; not local-backbone; not production messaging UX

Expected live umbrella validation commands now include:

    cd ~/repos/carbonstack_umbrella/carbonstack-comms
    go test ./internal/app -count=1
    go test ./... -count=1 -timeout 600s
    scripts/dev-openmls-runtime-smoke-wrappers.sh
    scripts/dev-openmls-runtime-smoke.sh

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go run . --profile dev-runtime-openmls --clean-generated
    go run . --profile local-cypher
    go run . --profile doctor
    go run . --profile core --clean-generated

Expected v0.4.0 public package validation commands remain unchanged:

    cd <package-root>/carbonstack/tools/carbonstack-validate
    go run . --profile verify-checksums --root <package-root>
    go run . --profile full --root <package-root> --clean-generated

---

## 5. v0.4.16 Work Completed

### 5.1 Wrapper-based runtime smoke variant was added

New script:

    carbonstack-comms/scripts/dev-openmls-runtime-smoke-wrappers.sh

Commit:

    cb4e59d test: add wrapper-based OpenMLS runtime smoke proof

The new script is a parallel proof, not a replacement for the original direct-sidecar smoke script.

Proof shape:

    openmls-*-dev bootstrap wrappers -> openmls-send-dev -> Cypher -> openmls-inbox-dev --ack

It creates:

    temporary workspace
    temporary Cypher binary
    temporary Cypher DB
    temporary Alice/Bob Comms state files
    temporary Cypher devices
    dev-local OpenMLS sidecar identities through Comms wrappers
    dev-local OpenMLS conversation through Comms wrappers
    wrapper-produced KeyPackage and Welcome artifacts

It validates:

    openmls-identity-create-dev for Alice/Bob.
    openmls-bundle-export-dev for Bob.
    openmls-conversation-create-dev for Alice.
    openmls-conversation-load-check-dev for Alice.
    openmls-conversation-add-member-dev for Alice adding Bob.
    openmls-conversation-join-dev for Bob consuming the Welcome.
    openmls-conversation-load-check-dev for Bob.
    openmls-send-dev from Alice to Bob through Cypher.
    openmls-inbox-dev --ack for Bob.
    plaintext match.
    ack after sidecar message-open success.
    empty Bob inbox after ack.

### 5.2 Original direct-sidecar smoke was preserved

Existing script remains:

    carbonstack-comms/scripts/dev-openmls-runtime-smoke.sh

It still proves:

    direct sidecar bootstrap -> openmls-send-dev -> Cypher -> openmls-inbox-dev --ack

No migration happened in v0.4.16.

Correct interpretation:

    wrapper-based smoke now exists as a second proof.
    direct-sidecar smoke remains the current runner-wrapped proof.
    no deletion/replacement happened.
    no runner profile change happened.

### 5.3 carbonstack-comms README was updated

Updated file:

    carbonstack-comms/README.md

New section:

    Wrapper-based dev runtime OpenMLS smoke proof

The README records:

    scripts/dev-openmls-runtime-smoke-wrappers.sh
    proof shape:
      openmls-*-dev bootstrap wrappers -> openmls-send-dev -> Cypher -> openmls-inbox-dev --ack
    boundary:
      dev/pre-alpha wrapper smoke proof
      not local-backbone
      not mature messaging UX
      not production E2EE
      does not replace the direct-sidecar smoke script yet
      manual dev-runtime-openmls runner profile still wraps scripts/dev-openmls-runtime-smoke.sh

### 5.4 Validation passed

Commands observed passing:

    cd ~/repos/carbonstack_umbrella/carbonstack-comms
    go test ./internal/app -count=1
    go test ./... -count=1 -timeout 600s
    scripts/dev-openmls-runtime-smoke-wrappers.sh
    scripts/dev-openmls-runtime-smoke.sh

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go test ./... -count=1
    go run . --profile dev-runtime-openmls --clean-generated
    go run . --profile local-cypher
    go run . --profile doctor
    go run . --profile core --clean-generated

Final commit/push:

    carbonstack-comms:
    cb4e59d test: add wrapper-based OpenMLS runtime smoke proof

Final repo snapshot:

    carbonstack        0583683 docs: define dev OpenMLS bootstrap command contract
    carbonstack-comms  cb4e59d test: add wrapper-based OpenMLS runtime smoke proof
    carbonstack-cypher 9ab994c docs: point to local operator runbook
    carbonstack-os     1bbbe52 docs: clarify CarbonStackOS target direction

---

## 6. Critical Blunders / Lessons to Carry Forward

### 6.1 Copy/paste truncation/glitching remains a workflow issue

The v0.4.16 recon pasteback was partially mangled near the `full smoke script` section and later command tail. Useful context still landed, but the scout was not perfectly clean.

Lesson:

    prefer smaller paste blocks when recon payloads are large.
    break long recon into inspect / validate / commit phases.
    keep avoiding copy box structures that have already caused breakage.
    final validation and final repo snapshots matter more than a perfect scout transcript.

### 6.2 Wrapper-based smoke should not immediately replace the direct smoke

v0.4.16 proves the wrapper-based smoke can pass, but the original direct-sidecar smoke remains valuable as the lower-level baseline.

Do not assume:

    direct-sidecar smoke can be deleted.
    dev-runtime-openmls should immediately switch to the wrapper smoke.
    full should include either runtime smoke.
    local-backbone is justified.

Correct current stance:

    keep both smoke scripts for now.
    direct-sidecar smoke remains runner-wrapped by dev-runtime-openmls.
    wrapper-based smoke is the higher-level parallel proof.
    future runner/doc alignment can decide whether to add an optional wrapper smoke profile or switch dev-runtime-openmls later.

### 6.3 Known generated roots are expected after smoke/tests

Both direct and wrapper smoke can create:

    .carbonstack-openmls-sidecar-state
    target

This remains acceptable only when:

    the paths are known generated roots.
    artifact scan classifies them correctly.
    --clean-generated removes them before a clean breakpoint.

### 6.4 Commit metadata warning remains a recurring polish issue

The commit log again showed auto-generated local commit metadata:

    Committer: bitcrusher32 <▓▓>

This is not a functional validation failure, but remains a professionalism/publishing hygiene issue to address eventually.

---

## 7. Current Blockers / Not Validated

Current blockers and nonvalidated areas after v0.4.16:

    Existing send/inbox remain stub-era.
    Runtime Comms has dev-only OpenMLS commands, direct smoke proof, wrapper-based smoke proof, a manual runner profile, and full bootstrap wrapper coverage, but no mature user-facing UX.
    The committed direct smoke script still exists and remains the dev-runtime-openmls runner target.
    The wrapper-based smoke script exists but is not yet a runner profile.
    The wrapper-based smoke script is not included in full.
    No Relay Space join model exists yet.
    local-backbone runner profile does not exist and remains reserved.
    dev-runtime-openmls remains manual-only and live-git-umbrella-only.
    dev-runtime-openmls is not included in full.
    dev-runtime-openmls is not part of public release package validation.
    Fresh release-package behavior for dev-runtime-openmls or wrapper smoke has not been tested as a supported validation surface.
    Post-release public download verification is still useful if not already performed.
    Sanitized PRIME placement in carbonstack/logdoc-loop-system still needs a policy decision.
    Production-safe provider storage and secure vault design remain v0.5.x future work.
    PQ/hybrid ciphersuite migration readiness remains v0.5.x future work.
    Hostile-server rollback/replay/metadata-abuse validation remains future work.
    Android/CarbonStackOS implementation remains later work.
    CarbonStack Relay Space mechanics remain a future concept, not an implementation.

Hard nonclaims remain:

    Do not claim production readiness.
    Do not claim production E2EE.
    Do not claim hostile-server safety.
    Do not claim metadata privacy.
    Do not claim Debian production deployability.
    Do not claim systemd readiness.
    Do not claim cloudflared readiness.
    Do not claim Android readiness.
    Do not claim CarbonStackOS readiness.
    Do not claim secure vault/storage.
    Do not claim real homelab validation.
    Do not claim external audit or certification.
    Do not claim CarbonStack Relay Space mechanics are implemented.
    Do not claim local-backbone exists or is justified yet.
    Do not claim local-cypher validates runtime Comms UX.
    Do not claim full validates local-backbone.
    Do not claim full is a deployment command.
    Do not claim dev-runtime-openmls is release-package validation yet.
    Do not claim wrapper-based smoke is release-package validation yet.
    Do not claim send/inbox are OpenMLS-backed yet.
    Do not claim PQ/hybrid security or quantum-safe messaging.
    Do not let GitHub mirrors supersede Gitea source-of-truth status.

Allowed claim:

    CarbonStackComms now has dev-only openmls-send-dev and openmls-inbox-dev commands.
    CarbonStackComms now has a direct-sidecar dev runtime smoke script proving openmls-send-dev -> Cypher -> openmls-inbox-dev --ack.
    CarbonStackComms now has a wrapper-based dev runtime smoke script proving openmls-*-dev bootstrap wrappers -> openmls-send-dev -> Cypher -> openmls-inbox-dev --ack.
    CarbonStackComms now has OpenMLS runtime command helper glue split into internal/app/openmls_runtime.go.
    CarbonStackComms now has dev-only OpenMLS bootstrap wrappers for identity create/status, bundle export, conversation create/load-check, conversation add-member, and conversation join.
    The bootstrap wrappers parse sidecar JSON, print stable key/value output, keep sidecar labels explicit, resolve path-consuming input artifacts to absolute paths, and do not mutate Comms state/trust files.
    carbonstack/tools/carbonstack-validate has a manual dev-runtime-openmls profile.
    dev-runtime-openmls wraps the direct-sidecar comms smoke script and validates the dev/pre-alpha OpenMLS application-message CLI path.
    dev-runtime-openmls is live-git-umbrella-only and not included in full.
    Existing send/inbox remain stub-era.
    v0.4.0 remains the current public broad local deployability pre-release.
    Gitea remains source of truth; GitHub mirrors are secondary push mirrors.

---

## 8. Next Safest Actions

Recommended next rung:

    preflight after v0.4.16

Questions for the preflight:

    Should `dev-runtime-openmls` stay pointed at the original direct-sidecar smoke, or should a second manual runner profile wrap the wrapper smoke?
    Should the wrapper smoke remain a Comms-only script for one more checkpoint before runner integration?
    Should a known-good command registry YAML/JSON come before or after runner alignment?
    Is local-backbone still premature after both smoke proofs pass? Likely yes, but the reason has shifted from missing proof to missing mature UX/state/trust/release-package semantics.
    Should README command sprawl be moved into a dedicated command registry / command reference doc?

Likely v0.4.17 options:

    Option A: docs-only smoke-boundary record in main carbonstack.
    Option B: add a second manual runner profile for wrapper smoke, e.g. dev-runtime-openmls-wrappers.
    Option C: keep runner unchanged and start known-good command registry planning.
    Option D: do a small README/command-reference cleanup to move lower-level dev command sprawl out of the front-door README.

Recommendation before deciding:

    perform a preflight.
    do not immediately switch dev-runtime-openmls.
    do not add wrapper smoke to full.
    do not call this local-backbone yet.

Avoid next:

    replacing the direct smoke script.
    adding dev-runtime-openmls to full.
    adding wrapper smoke to full.
    calling either smoke proof local-backbone.
    treating wrapper bootstrap as Relay Space join UX.
    moving into PQ/vault work before state/trust/storage preflight.
    public ingress.
    cloudflared.
    systemd.
    real homelab deployment.
    Android app.
    CarbonStackOS implementation.
    production/security claims.

---

## 9. Critical Paths / Functions

### Source-of-truth public surfaces

    [GITEA SOURCE OF TRUTH] https://git.bitcrusher32.win/bitcrusher32/carbonstack
    [RELEASES] https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases
    [CURRENT PUBLIC RELEASE] https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.4.0

GitHub mirrors:

    secondary push mirrors for discoverability/redundancy only.
    Do not treat mirrors as release authority unless project policy changes.

### WSL Debian active working umbrella

    [WSL-DEBIAN] umbrella: ~/repos/carbonstack_umbrella
    [WSL-DEBIAN] carbonstack: ~/repos/carbonstack_umbrella/carbonstack
    [WSL-DEBIAN] carbonstack-comms: ~/repos/carbonstack_umbrella/carbonstack-comms
    [WSL-DEBIAN] carbonstack-cypher: ~/repos/carbonstack_umbrella/carbonstack-cypher
    [WSL-DEBIAN] carbonstack-os: ~/repos/carbonstack_umbrella/carbonstack-os
    [GO-RUNNER] runner dir: ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate

### Runtime integration doc surfaces

    [DOC] v0.4.1 runtime recon:
    carbonstack/docs/155-runtime-comms-openmls-cypher-recon-v0.md

    [DOC] v0.4.2 command contract:
    carbonstack/docs/156-runtime-openmls-command-contract-v0.md

    [DOC] v0.4.6 smoke proof record:
    carbonstack/docs/157-dev-runtime-openmls-smoke-proof-v0.md

    [DOC] v0.4.7 pre-local-backbone assessment:
    carbonstack/docs/158-pre-local-backbone-assessment-v0.md

    [DOC] v0.4.9 runner profile:
    carbonstack/docs/159-dev-runtime-openmls-runner-profile-v0.md

    [DOC] v0.4.10 profile boundary:
    carbonstack/docs/160-dev-runtime-openmls-profile-boundary-v0.md

    [DOC] v0.4.11 bootstrap wrapper recon:
    carbonstack/docs/161-openmls-bootstrap-wrapper-recon-v0.md

    [DOC] v0.4.12 bootstrap command contract:
    carbonstack/docs/162-openmls-bootstrap-command-contract-v0.md

    [DOC] roadmap refresh:
    CarbonStack_Long_Term_Roadmap_v0.4.14_REFRESH.pdf

    [DOC] docs index:
    carbonstack/docs/README.md

    [DOC] roadmap:
    carbonstack/roadmap/ROADMAP.md

    [DOC] top README:
    carbonstack/README.md

### Runner surfaces

    [RUNNER] root:
    carbonstack/tools/carbonstack-validate

    [RUNNER] dispatch:
    carbonstack/tools/carbonstack-validate/main.go

    [RUNNER] dev-runtime-openmls implementation:
    carbonstack/tools/carbonstack-validate/dev_runtime_openmls.go

    [RUNNER] runner README:
    carbonstack/tools/carbonstack-validate/README.md

    [RUNNER] local-cypher:
    carbonstack/tools/carbonstack-validate/local_cypher.go

    [RUNNER] release-snapshot:
    carbonstack/tools/carbonstack-validate/release_snapshot.go

    [RUNNER] checksums:
    carbonstack/tools/carbonstack-validate/checksums.go

### Comms runtime surfaces

    [COMMS] CLI entry:
    carbonstack-comms/cmd/comms/main.go

    [COMMS] command registry and non-OpenMLS commands:
    carbonstack-comms/internal/app/commands.go

    [COMMS] OpenMLS message runtime command helpers:
    carbonstack-comms/internal/app/openmls_runtime.go

    [COMMS] OpenMLS bootstrap wrapper helpers:
    carbonstack-comms/internal/app/openmls_bootstrap.go

    [COMMS] OpenMLS bootstrap wrapper tests:
    carbonstack-comms/internal/app/openmls_bootstrap_dev_test.go

    [COMMS] openmls-send-dev tests:
    carbonstack-comms/internal/app/openmls_send_dev_test.go

    [COMMS] openmls-inbox-dev tests:
    carbonstack-comms/internal/app/openmls_inbox_dev_test.go

    [COMMS] direct-sidecar dev runtime smoke proof:
    carbonstack-comms/scripts/dev-openmls-runtime-smoke.sh

    [COMMS] wrapper-based dev runtime smoke proof:
    carbonstack-comms/scripts/dev-openmls-runtime-smoke-wrappers.sh

    [COMMS] OpenMLS sidecar dir:
    carbonstack-comms/internal/protocol/mls/openmls-sidecar

    [COMMS] sidecar README:
    carbonstack-comms/internal/protocol/mls/openmls-sidecar/README.md

### Current stub-era Comms functions / behaviors

    send:
      trust policy evaluation
      mock/stub provider encryption
      Cypher envelope submit

    inbox:
      Cypher inbox retrieval
      mock/stub provider plaintext display as stub_plaintext

    ack:
      envelope ack by envelope ID and recipient device ID

### Implemented v0.4.x dev OpenMLS commands

    openmls-send-dev:
      internal/app/openmls_runtime.go
      state.RequireReadyDevice
      trust.EvaluateSend
      runOpenMLSMessageProtect
      sidecar message-protect
      internal/relay.SubmitOpenMLSArtifactEnvelope
      ArtifactKindApplicationMessage
      dev/pre-alpha warning

    openmls-inbox-dev:
      internal/app/openmls_runtime.go
      state.RequireReadyDevice
      Cypher inbox
      skip unsupported envelopes
      internal/relay.WriteOpenMLSArtifactFromEnvelope
      runOpenMLSMessageOpen
      sidecar message-open
      print plaintext after sidecar success
      optional ack only after sidecar success

    openmls-identity-create-dev:
      internal/app/openmls_bootstrap.go
      sidecar identity-create
      required --sidecar-device-label
      optional --sidecar-dir
      stable key/value output
      no Comms state/trust mutation

    openmls-identity-status-dev:
      internal/app/openmls_bootstrap.go
      sidecar identity-status
      required --sidecar-device-label
      optional --sidecar-dir
      stable key/value output
      prints identity_exists when present
      no Comms state/trust mutation

    openmls-bundle-export-dev:
      internal/app/openmls_bootstrap.go
      sidecar public-bundle-export
      required --sidecar-device-label
      optional --sidecar-dir
      optional --write-artifact
      stable key/value output
      prints key_package_artifact_path_hint when present
      prints key_package_artifact_path when a hint is present
      no Comms state/trust mutation

    openmls-conversation-create-dev:
      internal/app/openmls_bootstrap.go
      sidecar conversation-create
      required --sidecar-device-label
      required --conversation
      stable key/value output
      prints sidecar_conversation_label
      prints conversation_state_path_hint when present
      prints conversation_summary_path_hint when present
      prints provider_storage_path_hint when present
      no Comms state/trust mutation

    openmls-conversation-load-check-dev:
      internal/app/openmls_bootstrap.go
      sidecar conversation-load-check
      required --sidecar-device-label
      required --conversation
      stable key/value output
      prints group_reloadable when present
      no Comms state/trust mutation

    openmls-conversation-add-member-dev:
      internal/app/openmls_bootstrap.go
      sidecar conversation-add-member
      required --sidecar-device-label
      required --conversation
      required --member-keypackage
      resolves --member-keypackage to absolute path before sidecar call
      stable key/value output
      prints welcome_artifact_path_hint and welcome_artifact_path when present
      prints welcome_manifest_path_hint and welcome_manifest_path when present
      prints welcome_artifact_sha256 and welcome_artifact_size_bytes when present
      prints member_added, welcome_artifact_written, group_reloadable, member counts, and epochs when present
      no Comms state/trust mutation

    openmls-conversation-join-dev:
      internal/app/openmls_bootstrap.go
      sidecar conversation-join
      required --sidecar-device-label
      required --conversation
      required --welcome
      resolves --welcome to absolute path before sidecar call
      stable key/value output
      prints joined, group_reloadable, member_count, epoch, join summary hint, conversation state hint, conversation summary hint, and provider storage hint when present
      no Comms state/trust mutation

### Runtime smoke proofs after v0.4.16

    scripts/dev-openmls-runtime-smoke.sh:
      temp workspace
      temp Cypher binary
      temp Cypher DB
      dynamic loopback Cypher port
      temp Alice/Bob Comms state
      sidecar state reset
      direct sidecar identity/conversation bootstrap
      openmls-send-dev application-message send
      openmls-inbox-dev application-message open + ack
      inbox-empty-after-ack verification
      robust cleanup with INT -> TERM -> KILL escalation

    scripts/dev-openmls-runtime-smoke-wrappers.sh:
      temp workspace
      temp Cypher binary
      temp Cypher DB
      dynamic loopback Cypher port
      temp Alice/Bob Comms state
      sidecar state reset
      Comms wrapper-based identity/conversation bootstrap
      openmls-send-dev application-message send
      openmls-inbox-dev application-message open + ack
      inbox-empty-after-ack verification
      robust cleanup with INT -> TERM -> KILL escalation

### v0.4.9/v0.4.10 runner profile

    go run . --profile dev-runtime-openmls --clean-generated

    Runner.DevRuntimeOpenMLS:
      CheckRequiredPaths
      CheckLiveGitUmbrella
      toolchain reporting
      ArtifactScan("pre-dev-runtime-openmls")
      RunStep for carbonstack-comms/scripts/dev-openmls-runtime-smoke.sh
      ArtifactScan("post-dev-runtime-openmls")
      explicit nonclaim summary

### Existing v0.4.0 package validation

    cd <package-root>/carbonstack/tools/carbonstack-validate
    go run . --profile verify-checksums --root <package-root>
    go run . --profile full --root <package-root> --clean-generated

---

## 10. v0.4.x Forward Plan

### Immediate next: preflight after v0.4.16

Expected:

    reassess relative completion after both direct and wrapper-based smoke proofs pass.
    decide whether to add a second manual runner profile for wrapper smoke.
    decide whether to record v0.4.16 smoke-boundary evidence in main carbonstack docs.
    decide whether known-good command registry planning should come next.
    decide whether README command sprawl should be moved into a command reference / registry.
    keep dev-runtime-openmls manual-only and outside full.
    keep local-backbone reserved unless a later preflight deliberately changes the claim boundary.

### v0.4.17+

Possible:

    docs-only smoke-boundary record.
    second manual runner profile for wrapper smoke.
    known-good command registry YAML/JSON outside the front-door README unless user-critical.
    exact sidecar setup runbook.
    clearer dev state reset semantics.
    better command output shape.
    decision on whether stub-era send/inbox are aliased, replaced, or deprecated.
    decision on when local-backbone naming stops being premature.

---

## 11. v0.5.x and Later Carry-Forward

### v0.5.x

Local storage, trust state, provider state, vault design, and PQ/hybrid ciphersuite migration readiness.

Do not start PQ/hybrid implementation before the runtime OpenMLS/Cypher path exists and is documented. PQ/hybrid migration touches provider state, ciphersuite metadata, trust display, artifact size, reinit/rekey/recovery policy, and claim boundaries.

### v0.6.x+

Hostile-server and abuse-resistance harnesses.

### v0.7.x+

Deployability and operations hardening.

### v0.8.x+

Documented self-pentest / adversarial validation.

### v0.9.x+

Claim-boundary review.

### v0.10.x+

Android backend/app work.

### v1.x.x

Public app/server major epoch and much later CarbonStackOS/appliance research, if justified.

---

## 12. Lean Breakpoint Summary

    CarbonStack v0.4.16 is the wrapper-based OpenMLS runtime smoke proof checkpoint. v0.4.0 remains the current public broad local deployability pre-release. carbonstack remains at 0583683 docs: define dev OpenMLS bootstrap command contract. carbonstack-comms is now at cb4e59d test: add wrapper-based OpenMLS runtime smoke proof. v0.4.16 adds scripts/dev-openmls-runtime-smoke-wrappers.sh and updates README.md with a wrapper-based dev runtime OpenMLS smoke proof section. The new script proves openmls-*-dev bootstrap wrappers -> openmls-send-dev -> Cypher -> openmls-inbox-dev --ack, while the original direct-sidecar scripts/dev-openmls-runtime-smoke.sh remains unchanged and still passes. The dev-runtime-openmls runner profile remains unchanged and still wraps the original direct-sidecar smoke script; full remains unchanged. Validation passed for app tests, full Comms tests, wrapper-based smoke, direct smoke, dev-runtime-openmls --clean-generated, local-cypher, doctor, and core --clean-generated. Final repo heads were clean across all four repos. Next safe action: preflight after v0.4.16 to decide runner/doc/registry alignment.

---

## 13. Preserved Immediate Previous Handoff: v0.4.15

The following is the previous v0.4.15 handoff. Where it conflicts with the v0.4.16 overlay above, v0.4.16 wins for current state.




# CarbonStack LogDoc v0.4.15

**Last updated:** 2026-06-05 local session  
**Owner/operator:** bitcrusher32  
**Assistant/model:** GPT-5.5 Thinking  
**Current phase:** Phase 2F / **v0.4.x dev-only OpenMLS add-member/join bootstrap wrapper implementation checkpoint complete**. The v0.4.0 public broad local deployability pre-release remains the current public release. v0.4.12 defined the strict dev-only OpenMLS bootstrap wrapper contract; v0.4.13 implemented identity create/status wrappers; v0.4.14 implemented bundle export and conversation create/load-check wrappers; v0.4.15 now completes the planned dev-only bootstrap wrapper set in `carbonstack-comms` with `openmls-conversation-add-member-dev` and `openmls-conversation-join-dev`. `carbonstack` remains at `0583683 (HEAD -> main, origin/main, origin/HEAD) docs: define dev OpenMLS bootstrap command contract`; `carbonstack-comms` is now at `c2d3b16 (HEAD -> main, origin/main, origin/HEAD) feat: add dev OpenMLS add-member and join bootstrap commands`; `carbonstack-cypher` remains at `9ab994c (HEAD -> main, origin/main, origin/HEAD) docs: point to local operator runbook`; `carbonstack-os` remains at `1bbbe52 (HEAD -> main, origin/main, origin/HEAD) docs: clarify CarbonStackOS target direction`.

**Version schema:** `v[scope].[timeline]`. This file is `v0.4.15`, the third staged dev-only OpenMLS bootstrap wrapper implementation checkpoint after the v0.4.14 bundle/conversation wrapper checkpoint. Per LogDoc V2 practice, this Markdown file preserves current operational continuity and carries the v0.4.x runtime-integration thread forward; the JSON breakpoint remains the lean machine-readable handoff state.

**Compression note:** The detailed v0.3.0 through v0.3.36PRIME process ledger remains archived in PRIME/sanitized PRIME. v0.4.0 compressed that pre-release sprint; v0.4.1 through v0.4.12 preserved runtime recon, command contracts, dev OpenMLS send/inbox commands, smoke proof, runner profile, boundary checks, wrapper recon, and bootstrap command contract. v0.4.13 preserved identity create/status wrappers. v0.4.14 preserved bundle export and conversation create/load-check wrappers. v0.4.15 now preserves add-member / Welcome generation and join / Welcome consumption wrappers.

---

## 1. Project Goal

**Active goal:** Continue moving CarbonStackComms from stub-era runtime `send`/`inbox` behavior toward explicit dev-only OpenMLS/Cypher-backed runtime commands and repeatable validation, while preserving claim discipline. v0.4.15 completes the planned dev-only bootstrap wrapper surface, but does **not** migrate the smoke script, does **not** rename anything local-backbone, and does **not** change the release/package validation boundary.

CarbonStack’s broader project intent remains:

- **CarbonStack** is the secure-communications appliance-stack project and shared doctrine/release authority.
- **CarbonStackComms** is the text-first messaging client and current OpenMLS sidecar/relay integration surface.
- **CarbonStackCypher** is the self-hostable relay/storage stack.
- **CarbonStackOS** is future constrained Android-derived appliance OS work and remains out of current runnable validation packages.
- The `carbonstack` Gitea repo remains the source of truth for release state, release assets, docs, and public claims.
- GitHub mirrors exist as secondary push mirrors only; Gitea remains authoritative.

**North star:**

> Its goal is not to be a better smartphone. Its goal is to stop being a smartphone.

**Core doctrine:**

> Every feature is guilty until it proves it does not add unacceptable parser, network, sensor, identity, or filesystem authority.

**Current reality after v0.4.15:**

    openmls-send-dev is a dev-only OpenMLS send-side command.
    openmls-inbox-dev is a dev-only OpenMLS receive/open/optional-ack command.
    scripts/dev-openmls-runtime-smoke.sh proves the dev CLI application-message path:
      openmls-send-dev -> Cypher -> openmls-inbox-dev --ack
    tools/carbonstack-validate has a manual dev-runtime-openmls profile that wraps that smoke script.
    dev-runtime-openmls remains manual-only, live-git-umbrella-only, and outside full.

    The planned dev-only OpenMLS bootstrap wrapper set now exists:
      openmls-identity-create-dev
      openmls-identity-status-dev
      openmls-bundle-export-dev
      openmls-conversation-create-dev
      openmls-conversation-load-check-dev
      openmls-conversation-add-member-dev
      openmls-conversation-join-dev

    Bootstrap wrappers live in carbonstack-comms/internal/app/openmls_bootstrap.go.
    Bootstrap wrapper tests live in carbonstack-comms/internal/app/openmls_bootstrap_dev_test.go.
    Bootstrap wrappers keep sidecar labels explicit.
    Bootstrap wrappers normalize sidecar JSON to stable key/value Comms output.
    Path-consuming wrappers resolve input artifacts to absolute paths before invoking the sidecar.
    Add-member preserves member_keypackage_path_hint, welcome_artifact_path_hint, welcome_artifact_path, welcome manifest hints, Welcome hash/size, and membership/epoch fields when present.
    Join preserves welcome_artifact_path_hint, joined, group_reloadable, member_count, epoch, and conversation/provider path hints when present.
    Bootstrap wrappers do not mutate Comms state/trust files.
    carbonstack-comms README documents the current dev-only OpenMLS bootstrap commands and boundaries.
    Existing send/inbox remain stub-era.
    The smoke script still uses direct sidecar bootstrap.
    No wrapper-based smoke variant exists yet.
    local-backbone remains reserved.

Hard nonclaims remain: v0.4.15 does **not** prove production readiness, production E2EE, hostile-server safety, metadata privacy, Android readiness, CarbonStackOS readiness, secure vault/storage, public ingress safety, real homelab validation, audit, certification, PQ/hybrid security, mature messenger UX, release-package runtime validation, Relay Space join/onboarding, smoke-script wrapper migration, or local-backbone readiness.

---

## 2. Current Public Release State

Public release:

    https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.4.0

Release title:

    CarbonStack v0.4.0 Broad Local Deployability Pre-Release

Release tag:

    v0.4.0

Release commit:

    24ac9fc494 / 24ac9fc docs: record v0.4.0 package rehearsal

Gitea release status:

    Pre-Release

Important release-page warning remains:

    The default Gitea Source Code ZIP/TAR.GZ downloads are auto-generated carbonstack-only archives.
    The intended runnable validation artifact is the attached multi-repo release package plus companion assets.

Post-release public surface:

    carbonstack README is evergreen.
    It points interested parties to the Releases page for runnable / known-good artifacts.
    Gitea is source of truth.
    GitHub mirrors are secondary push mirrors, not release authority.

---

## 3. Current Repo Heads

    carbonstack        0583683 (HEAD -> main, origin/main, origin/HEAD) docs: define dev OpenMLS bootstrap command contract
    carbonstack-comms  c2d3b16 (HEAD -> main, origin/main, origin/HEAD) feat: add dev OpenMLS add-member and join bootstrap commands
    carbonstack-cypher 9ab994c (HEAD -> main, origin/main, origin/HEAD) docs: point to local operator runbook
    carbonstack-os     1bbbe52 (HEAD -> main, origin/main, origin/HEAD) docs: clarify CarbonStackOS target direction

Current public release tag:

    v0.4.0 -> 24ac9fc494 / 24ac9fc docs: record v0.4.0 package rehearsal

Recent runtime/thread commits:

    v0.4.8 extraction:
      0a48ae1 refactor: split OpenMLS runtime command helpers

    v0.4.9 runner profile:
      8eeadb2 feat: add dev OpenMLS runtime validation profile

    v0.4.10 profile boundary:
      7384f43 docs: record dev OpenMLS runtime profile boundary

    v0.4.11 wrapper recon:
      8ee3ebb docs: plan dev OpenMLS bootstrap wrappers

    v0.4.12 command contract:
      0583683 docs: define dev OpenMLS bootstrap command contract

    v0.4.13 identity wrapper implementation:
      20c4cc9 feat: add dev OpenMLS identity bootstrap commands

    v0.4.14 bundle/conversation wrapper implementation:
      1f5c09e feat: add dev OpenMLS bundle and conversation bootstrap commands

    Current v0.4.15 add-member/join wrapper implementation:
      c2d3b16 feat: add dev OpenMLS add-member and join bootstrap commands

Continuity note:

    `c2d3b16` is a carbonstack-comms implementation commit after the v0.4.0 public release tag.
    It does not retag v0.4.0.
    The release remains attached to `24ac9fc494`.

---

## 4. Current Validated / Known State

Validated / known at the v0.4.15 checkpoint:

    [RELEASE] v0.4.0 Gitea release exists, is marked Pre-Release, and is attached to 24ac9fc494.
    [RUNNER] full means release-snapshot followed by local-cypher.
    [RUNNER] dev-runtime-openmls validates the dev/pre-alpha OpenMLS application-message runtime CLI smoke proof.
    [RUNNER] dev-runtime-openmls remains manual-only and live-umbrella-only.
    [RUNNER] dev-runtime-openmls is not included in full.
    [DOCS] docs/162-openmls-bootstrap-command-contract-v0.md records the v0.4.12 dev-only OpenMLS bootstrap command contract.
    [COMMS] openmls-send-dev exists.
    [COMMS] openmls-inbox-dev exists.
    [COMMS] all planned dev-only OpenMLS bootstrap wrappers now exist.
    [COMMS] internal/app/openmls_bootstrap.go contains identity, bundle, conversation-create/load-check, add-member, and join wrappers.
    [COMMS] internal/app/openmls_bootstrap_dev_test.go tests the current bootstrap wrapper set with injected sidecar seams.
    [COMMS] internal/app/commands.go registers the new wrapper commands and lists them in usage.
    [COMMS] README.md documents the current dev-only OpenMLS bootstrap commands and their boundary.
    [COMMS] scripts/dev-openmls-runtime-smoke.sh still uses direct sidecar bootstrap.
    [COMMS] existing send/inbox/ack remain in place and stub-era.
    [MIRRORS] all four repos have GitHub push-mirror functionality, with Gitea remaining source of truth.

v0.4.15 validation result from the final log:

    go test ./internal/app -count=1 passed in carbonstack-comms.
    go test ./... -count=1 -timeout 600s passed in carbonstack-comms.
    Manual full wrapper chain smoke passed:
      identity-create-dev for Alice.
      identity-create-dev for Bob.
      bundle-export-dev for Bob.
      conversation-create-dev for Alice.
      conversation-load-check-dev for Alice.
      conversation-add-member-dev for Alice adding Bob.
      conversation-join-dev for Bob consuming Welcome.
      conversation-load-check-dev for Bob after join.
    add-member output included:
      status: welcome_created.
      member_added: true.
      welcome_artifact_written: true.
      group_reloadable: true.
      member_count_before: 1.
      member_count_after: 2.
      epoch_before: GroupEpoch(0).
      epoch_after: GroupEpoch(1).
      welcome_artifact_path_hint.
      welcome_artifact_path.
      welcome_manifest_path_hint.
      welcome_manifest_path.
      welcome_artifact_sha256.
      welcome_artifact_size_bytes: 879.
    join output included:
      status: joined.
      joined: true.
      group_reloadable: true.
      member_count: 2.
      epoch: GroupEpoch(1).
      join_summary_path_hint.
      conversation_state_path_hint.
      conversation_summary_path_hint.
      provider_storage_path_hint.
    go test ./... -count=1 passed in carbonstack/tools/carbonstack-validate.
    go run . --profile dev-runtime-openmls --clean-generated passed.
    local-cypher passed.
    doctor passed.
    core --clean-generated passed.
    carbonstack-comms was committed and pushed at:
      c2d3b16 feat: add dev OpenMLS add-member and join bootstrap commands
    final repo status showed all four repos clean.

Expected live umbrella validation commands now include:

    cd ~/repos/carbonstack_umbrella/carbonstack-comms
    go test ./internal/app -count=1
    go test ./... -count=1 -timeout 600s

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go run . --profile dev-runtime-openmls --clean-generated
    go run . --profile local-cypher
    go run . --profile doctor
    go run . --profile core --clean-generated

New manual full wrapper-chain check:

    cd ~/repos/carbonstack_umbrella/carbonstack-comms
    rm -rf internal/protocol/mls/openmls-sidecar/.carbonstack-openmls-sidecar-state

    go run ./cmd/comms openmls-identity-create-dev --sidecar-device-label carbonstack-v0415-alice
    go run ./cmd/comms openmls-identity-create-dev --sidecar-device-label carbonstack-v0415-bob

    bob_bundle_output="$(go run ./cmd/comms openmls-bundle-export-dev --sidecar-device-label carbonstack-v0415-bob --write-artifact)"
    printf '%s
' "$bob_bundle_output"
    bob_keypackage="$(printf '%s
' "$bob_bundle_output" | awk -F': ' '/^key_package_artifact_path: / {print $2; exit}')"

    go run ./cmd/comms openmls-conversation-create-dev --sidecar-device-label carbonstack-v0415-alice --conversation carbonstack-v0415-conversation
    go run ./cmd/comms openmls-conversation-load-check-dev --sidecar-device-label carbonstack-v0415-alice --conversation carbonstack-v0415-conversation

    add_member_output="$(go run ./cmd/comms openmls-conversation-add-member-dev --sidecar-device-label carbonstack-v0415-alice --conversation carbonstack-v0415-conversation --member-keypackage "$bob_keypackage")"
    printf '%s
' "$add_member_output"
    welcome_path="$(printf '%s
' "$add_member_output" | awk -F': ' '/^welcome_artifact_path: / {print $2; exit}')"

    go run ./cmd/comms openmls-conversation-join-dev --sidecar-device-label carbonstack-v0415-bob --conversation carbonstack-v0415-conversation --welcome "$welcome_path"
    go run ./cmd/comms openmls-conversation-load-check-dev --sidecar-device-label carbonstack-v0415-bob --conversation carbonstack-v0415-conversation

Expected v0.4.0 public package validation commands remain unchanged:

    cd <package-root>/carbonstack/tools/carbonstack-validate
    go run . --profile verify-checksums --root <package-root>
    go run . --profile full --root <package-root> --clean-generated

---

## 5. v0.4.15 Work Completed

### 5.1 Third bootstrap wrapper subset was implemented

New commands implemented in carbonstack-comms:

    openmls-conversation-add-member-dev
    openmls-conversation-join-dev

Together with v0.4.13 and v0.4.14, the currently implemented dev-only bootstrap wrapper set is now:

    openmls-identity-create-dev
    openmls-identity-status-dev
    openmls-bundle-export-dev
    openmls-conversation-create-dev
    openmls-conversation-load-check-dev
    openmls-conversation-add-member-dev
    openmls-conversation-join-dev

These commands are intentionally dev-only and sidecar-bootstrap-scoped.

They do not:

    replace send/inbox.
    mutate Comms state.json.
    mutate trust.json.
    mutate trust-events.jsonl.
    call Cypher.
    implement Relay Space join UX.
    imply local-backbone.
    imply production identity, conversation, or membership UX.

### 5.2 openmls_bootstrap.go was extended

Updated file:

    carbonstack-comms/internal/app/openmls_bootstrap.go

New command functions:

    cmdOpenMLSConversationAddMemberDev
    cmdOpenMLSConversationJoinDev

New helper behavior:

    bootstrapPrintOptionalBool
    bootstrapPrintOptionalNumber

Implementation behavior:

    openmls-conversation-add-member-dev calls sidecar conversation-add-member.
    openmls-conversation-add-member-dev requires --sidecar-device-label, --conversation, and --member-keypackage.
    openmls-conversation-add-member-dev resolves --member-keypackage to an absolute path before sidecar invocation.
    openmls-conversation-add-member-dev preserves member_keypackage_path_hint, welcome_artifact_path_hint, welcome_manifest_path_hint, welcome hash/size, membership booleans/counts, and epoch fields.
    openmls-conversation-add-member-dev resolves and prints welcome_artifact_path and welcome_manifest_path when hints are present.
    openmls-conversation-join-dev calls sidecar conversation-join.
    openmls-conversation-join-dev requires --sidecar-device-label, --conversation, and --welcome.
    openmls-conversation-join-dev resolves --welcome to an absolute path before sidecar invocation.
    openmls-conversation-join-dev preserves welcome_artifact_path_hint, joined, group_reloadable, member_count, epoch, join summary hint, conversation state hint, conversation summary hint, and provider storage hint.
    all new wrappers print stable key/value output and dev/pre-alpha membership UX warnings.

### 5.3 openmls_bootstrap_dev_test.go was extended

Updated file:

    carbonstack-comms/internal/app/openmls_bootstrap_dev_test.go

New tests cover:

    openmls-conversation-add-member-dev requires --sidecar-device-label, --conversation, and --member-keypackage.
    openmls-conversation-join-dev requires --sidecar-device-label, --conversation, and --welcome.
    add-member invokes conversation-add-member with expected args.
    add-member prints stable output, path hints, resolved welcome paths, hash/size, membership booleans/counts, and epoch fields.
    join invokes conversation-join with expected args.
    join prints stable output, joined/group_reloadable/member_count/epoch, and sidecar path hints.

Test design still follows the v0.4.12 contract:

    app tests use injected seams.
    app tests do not run the Rust sidecar directly.
    real sidecar execution remains in manual wrapper smoke, protocol tests, smoke script, and runner validation.

### 5.4 commands.go was updated

Updated file:

    carbonstack-comms/internal/app/commands.go

New dispatch cases:

    openmls-conversation-add-member-dev
    openmls-conversation-join-dev

Usage output now lists:

    openmls-conversation-add-member-dev
    openmls-conversation-join-dev

The existing command registry remains in commands.go. Message runtime helpers remain in openmls_runtime.go. Bootstrap wrapper helpers remain in openmls_bootstrap.go.

### 5.5 carbonstack-comms README was updated

Updated file:

    carbonstack-comms/README.md

The `Dev-only OpenMLS bootstrap commands` section now documents:

    openmls-identity-create-dev
    openmls-identity-status-dev
    openmls-bundle-export-dev
    openmls-conversation-create-dev
    openmls-conversation-load-check-dev
    openmls-conversation-add-member-dev
    openmls-conversation-join-dev

It includes example commands for add-member and join:

    go run ./cmd/comms openmls-conversation-add-member-dev --sidecar-device-label carbonstack-dev-alice --conversation carbonstack-dev-conversation --member-keypackage <path-to-member-keypackage>
    go run ./cmd/comms openmls-conversation-join-dev --sidecar-device-label carbonstack-dev-bob --conversation carbonstack-dev-conversation --welcome <path-to-welcome>

Boundary preserved:

    sidecar labels are explicit for now.
    Comms state/trust files are not mutated by these wrappers.
    existing send/inbox remain stub-era.
    dev-runtime-openmls remains the current manual smoke-profile proof.
    these commands are not production identity UX, not Relay Space join UX, not local-backbone, and not secure vault/state management.

### 5.6 Manual full wrapper-chain smoke passed

Manual check used Alice/Bob wrapper bootstrap through:

    identity create
    bundle export
    conversation create
    conversation load-check
    add-member / Welcome export
    join / Welcome consume
    Bob conversation load-check

Observed add-member output included:

    command: openmls-conversation-add-member-dev
    status: welcome_created
    sidecar_command: conversation-add-member
    sidecar_device_label: carbonstack-v0415-alice
    sidecar_conversation_label: carbonstack-v0415-conversation
    member_keypackage_path_hint: /home/Bitcrusher32/repos/carbonstack_umbrella/carbonstack-comms/internal/protocol/mls/openmls-sidecar/.carbonstack-openmls-sidecar-state/dev/devices/carbonstack-v0415-bob/public-bundle.keypackage.bin
    welcome_artifact_path_hint: .carbonstack-openmls-sidecar-state/dev/devices/carbonstack-v0415-alice/conversations/carbonstack-v0415-conversation/welcome.bin
    welcome_artifact_path: internal/protocol/mls/openmls-sidecar/.carbonstack-openmls-sidecar-state/dev/devices/carbonstack-v0415-alice/conversations/carbonstack-v0415-conversation/welcome.bin
    welcome_manifest_path_hint: .carbonstack-openmls-sidecar-state/dev/devices/carbonstack-v0415-alice/conversations/carbonstack-v0415-conversation/welcome-manifest.json
    welcome_manifest_path: internal/protocol/mls/openmls-sidecar/.carbonstack-openmls-sidecar-state/dev/devices/carbonstack-v0415-alice/conversations/carbonstack-v0415-conversation/welcome-manifest.json
    welcome_artifact_sha256: sha256:96b90bfb7378a862f532dbd8afd14264e89c94b2490041910eb0919b3c9d8521
    welcome_artifact_size_bytes: 879
    member_added: true
    welcome_artifact_written: true
    group_reloadable: true
    member_count_before: 1
    member_count_after: 2
    epoch_before: GroupEpoch(0)
    epoch_after: GroupEpoch(1)

Observed join output included:

    command: openmls-conversation-join-dev
    status: joined
    sidecar_command: conversation-join
    sidecar_device_label: carbonstack-v0415-bob
    sidecar_conversation_label: carbonstack-v0415-conversation
    welcome_artifact_path_hint: /home/Bitcrusher32/repos/carbonstack_umbrella/carbonstack-comms/internal/protocol/mls/openmls-sidecar/.carbonstack-openmls-sidecar-state/dev/devices/carbonstack-v0415-alice/conversations/carbonstack-v0415-conversation/welcome.bin
    joined: true
    group_reloadable: true
    member_count: 2
    epoch: GroupEpoch(1)
    join_summary_path_hint: .carbonstack-openmls-sidecar-state/dev/devices/carbonstack-v0415-bob/conversations/carbonstack-v0415-conversation/join-summary.json
    conversation_state_path_hint: .carbonstack-openmls-sidecar-state/dev/devices/carbonstack-v0415-bob/conversations/carbonstack-v0415-conversation
    conversation_summary_path_hint: .carbonstack-openmls-sidecar-state/dev/devices/carbonstack-v0415-bob/conversations/carbonstack-v0415-conversation/conversation-summary.json
    provider_storage_path_hint: .carbonstack-openmls-sidecar-state/dev/devices/carbonstack-v0415-bob/conversations/carbonstack-v0415-conversation/provider-storage.json

Bob load-check after join passed with:

    command: openmls-conversation-load-check-dev
    status: loaded
    sidecar_command: conversation-load-check
    sidecar_device_label: carbonstack-v0415-bob
    sidecar_conversation_label: carbonstack-v0415-conversation
    group_reloadable: true

### 5.7 Validation passed

Commands observed passing:

    cd ~/repos/carbonstack_umbrella/carbonstack-comms
    go test ./internal/app -count=1
    go test ./... -count=1 -timeout 600s

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go test ./... -count=1
    go run . --profile dev-runtime-openmls --clean-generated
    go run . --profile local-cypher
    go run . --profile doctor
    go run . --profile core --clean-generated

Final commit/push:

    carbonstack-comms:
    c2d3b16 feat: add dev OpenMLS add-member and join bootstrap commands

Final repo snapshot:

    carbonstack        0583683 docs: define dev OpenMLS bootstrap command contract
    carbonstack-comms  c2d3b16 feat: add dev OpenMLS add-member and join bootstrap commands
    carbonstack-cypher 9ab994c docs: point to local operator runbook
    carbonstack-os     1bbbe52 docs: clarify CarbonStackOS target direction

---

## 6. Critical Blunders / Lessons to Carry Forward

### 6.1 Add-member/join recon exposed a sidecar cwd/path contract

The initial v0.4.15 scout failed in a useful way:

    direct sidecar add-member failed with conversation_or_artifact_missing / identity state is missing.
    the failure occurred because `cargo run --manifest-path` was invoked from the Comms repo root.
    the sidecar dev state is cwd-relative to the sidecar directory.
    the existing wrapper runner uses cmd.Dir = sidecarDir, which is the correct execution contract.

Lesson:

    For sidecar-backed wrappers, keep the sidecar process cwd as sidecarDir.
    For path-consuming sidecar commands, resolve input artifact paths to absolute paths before passing them into the sidecar.
    Do not assume --manifest-path also gives sidecar cwd semantics.

This lesson is important for future smoke-script migration and any runner/profile work that calls sidecar commands directly.

### 6.2 Manual wrapper smoke again left known generated roots before dev-runtime-openmls

As in v0.4.13 and v0.4.14, the manual wrapper smoke intentionally created OpenMLS sidecar state and Rust build artifacts. During the later `dev-runtime-openmls` run, the pre-profile artifact scan reported known generated roots:

    .carbonstack-openmls-sidecar-state
    target

This stayed inside known generated paths and was cleaned by `--clean-generated`.

Lesson:

    Manual wrapper smoke can leave expected sidecar state and target roots.
    This is acceptable only when they stay in known generated paths.
    Always follow with --clean-generated when preparing a clean breakpoint.

### 6.3 Commit metadata warning and first-push auth failure remain recurring polish issues

The commit log again showed auto-generated local commit metadata:

    Committer: bitcrusher32 <▓▓>

The first `git push` failed authentication, and the follow-up `git push` succeeded.

This is not a functional validation failure, but remains a professionalism/publishing hygiene issue to address eventually.

### 6.4 Wrapper coverage is complete, but wrapper-based smoke is not

v0.4.15 completes the planned wrapper surface from v0.4.12.

Do not assume:

    smoke script uses wrappers.
    direct sidecar smoke can be removed.
    wrapper-based smoke exists.
    full validates wrapper-based bootstrap.
    local-backbone is justified.

Correct interpretation:

    All planned dev-only bootstrap wrappers now exist and passed manual wrapper-chain smoke.
    The committed runtime smoke proof still uses direct sidecar bootstrap.
    The next safe question is wrapper-based smoke variant / migration recon.

---

## 7. Current Blockers / Not Validated

Current blockers and nonvalidated areas after v0.4.15:

    Existing send/inbox remain stub-era.
    Runtime Comms has dev-only OpenMLS commands, a smoke proof, a manual runner profile, and full bootstrap wrapper coverage, but no mature user-facing UX.
    The committed smoke script still uses direct sidecar bootstrap.
    No wrapper-based smoke variant exists yet.
    Smoke script has not been migrated to wrappers.
    No Relay Space join model exists yet.
    local-backbone runner profile does not exist and remains reserved.
    dev-runtime-openmls remains manual-only and live-git-umbrella-only.
    dev-runtime-openmls is not included in full.
    dev-runtime-openmls is not part of public release package validation.
    Fresh release-package behavior for dev-runtime-openmls has not been tested as a supported validation surface.
    Post-release public download verification is still useful if not already performed.
    Sanitized PRIME placement in carbonstack/logdoc-loop-system still needs a policy decision.
    Production-safe provider storage and secure vault design remain v0.5.x future work.
    PQ/hybrid ciphersuite migration readiness remains v0.5.x future work.
    Hostile-server rollback/replay/metadata-abuse validation remains future work.
    Android/CarbonStackOS implementation remains later work.
    CarbonStack Relay Space mechanics remain a future concept, not an implementation.

Hard nonclaims remain:

    Do not claim production readiness.
    Do not claim production E2EE.
    Do not claim hostile-server safety.
    Do not claim metadata privacy.
    Do not claim Debian production deployability.
    Do not claim systemd readiness.
    Do not claim cloudflared readiness.
    Do not claim Android readiness.
    Do not claim CarbonStackOS readiness.
    Do not claim secure vault/storage.
    Do not claim real homelab validation.
    Do not claim external audit or certification.
    Do not claim CarbonStack Relay Space mechanics are implemented.
    Do not claim local-backbone exists or is justified yet.
    Do not claim local-cypher validates runtime Comms UX.
    Do not claim full validates local-backbone.
    Do not claim full is a deployment command.
    Do not claim dev-runtime-openmls is release-package validation yet.
    Do not claim send/inbox are OpenMLS-backed yet.
    Do not claim PQ/hybrid security or quantum-safe messaging.
    Do not claim smoke script uses wrappers yet.
    Do not let GitHub mirrors supersede Gitea source-of-truth status.

Allowed claim:

    CarbonStackComms now has dev-only openmls-send-dev and openmls-inbox-dev commands.
    CarbonStackComms now has a dev runtime smoke script proving openmls-send-dev -> Cypher -> openmls-inbox-dev --ack.
    CarbonStackComms now has OpenMLS runtime command helper glue split into internal/app/openmls_runtime.go.
    CarbonStackComms now has dev-only OpenMLS bootstrap wrappers for identity create/status, bundle export, conversation create/load-check, conversation add-member, and conversation join.
    The bootstrap wrappers parse sidecar JSON, print stable key/value output, and keep sidecar labels explicit.
    The bootstrap wrappers resolve path-consuming input artifacts to absolute paths before calling the sidecar.
    The bootstrap wrappers do not mutate Comms state/trust files.
    carbonstack/tools/carbonstack-validate has a manual dev-runtime-openmls profile.
    dev-runtime-openmls wraps the comms smoke script and validates the dev/pre-alpha OpenMLS application-message CLI path.
    dev-runtime-openmls is live-git-umbrella-only and not included in full.
    The main carbonstack repo records the smoke proof, pre-local-backbone assessment, manual dev-runtime-openmls runner profile, profile boundary check, bootstrap wrapper recon/planning doc, and bootstrap wrapper command contract.
    Existing send/inbox remain stub-era.
    v0.4.0 remains the current public broad local deployability pre-release.
    Gitea remains source of truth; GitHub mirrors are secondary push mirrors.

---

## 8. Next Safest Actions

Recommended next rung:

    v0.4.16 smoke-script migration recon / wrapper-based smoke variant decision

Focus:

    inspect scripts/dev-openmls-runtime-smoke.sh direct sidecar bootstrap sequence.
    compare direct sidecar bootstrap to the now-complete wrapper chain.
    decide whether to:
      keep direct-sidecar smoke only,
      add wrapper-based smoke as a second script/profile,
      or migrate the existing smoke script to wrappers.
    preserve the direct smoke proof until wrapper-based smoke proves equally stable.
    avoid local-backbone naming.
    avoid dev-runtime-openmls/full changes until recon supports them.
    preserve the cwd/path lesson from v0.4.15.
    validate direct dev-runtime-openmls before/after any candidate changes.

Possible v0.4.16 implementation only if recon is clean:

    add scripts/dev-openmls-runtime-wrapper-smoke.sh as a parallel smoke variant.
    do not replace scripts/dev-openmls-runtime-smoke.sh immediately.
    maybe add docs/README note or Comms README note if a parallel wrapper smoke lands.
    do not add runner profile yet unless wrapper smoke proves stable.

Future late v0.4.x:

    known-good command registry YAML/JSON outside the front-door README unless user-critical.
    exact sidecar setup runbook.
    clearer dev state reset semantics.
    better command output shape.
    decision on whether stub-era send/inbox are aliased, replaced, or deprecated.
    decision on when local-backbone naming stops being premature.

Avoid next:

    replacing the smoke script before wrapper-based smoke is proven.
    adding dev-runtime-openmls to full.
    calling any wrapper local-backbone.
    using Relay Space join naming for dev bootstrap wrappers.
    moving dev-runtime-openmls into public release validation without fresh-root behavior checks.
    public ingress.
    cloudflared.
    systemd.
    real homelab deployment.
    Android app.
    CarbonStackOS implementation.
    CarbonStack Relay Space implementation.
    PQ/hybrid ciphersuite implementation.
    production/security claims.

---

## 9. Critical Paths / Functions

### Source-of-truth public surfaces

    [GITEA SOURCE OF TRUTH] https://git.bitcrusher32.win/bitcrusher32/carbonstack
    [RELEASES] https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases
    [CURRENT PUBLIC RELEASE] https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.4.0

GitHub mirrors:

    secondary push mirrors for discoverability/redundancy only.
    Do not treat mirrors as release authority unless project policy changes.

### WSL Debian active working umbrella

    [WSL-DEBIAN] umbrella: ~/repos/carbonstack_umbrella
    [WSL-DEBIAN] carbonstack: ~/repos/carbonstack_umbrella/carbonstack
    [WSL-DEBIAN] carbonstack-comms: ~/repos/carbonstack_umbrella/carbonstack-comms
    [WSL-DEBIAN] carbonstack-cypher: ~/repos/carbonstack_umbrella/carbonstack-cypher
    [WSL-DEBIAN] carbonstack-os: ~/repos/carbonstack_umbrella/carbonstack-os
    [GO-RUNNER] runner dir: ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate

### Runtime integration doc surfaces

    [DOC] v0.4.1 runtime recon:
    carbonstack/docs/155-runtime-comms-openmls-cypher-recon-v0.md

    [DOC] v0.4.2 command contract:
    carbonstack/docs/156-runtime-openmls-command-contract-v0.md

    [DOC] v0.4.6 smoke proof record:
    carbonstack/docs/157-dev-runtime-openmls-smoke-proof-v0.md

    [DOC] v0.4.7 pre-local-backbone assessment:
    carbonstack/docs/158-pre-local-backbone-assessment-v0.md

    [DOC] v0.4.9 runner profile:
    carbonstack/docs/159-dev-runtime-openmls-runner-profile-v0.md

    [DOC] v0.4.10 profile boundary:
    carbonstack/docs/160-dev-runtime-openmls-profile-boundary-v0.md

    [DOC] v0.4.11 bootstrap wrapper recon:
    carbonstack/docs/161-openmls-bootstrap-wrapper-recon-v0.md

    [DOC] v0.4.12 bootstrap command contract:
    carbonstack/docs/162-openmls-bootstrap-command-contract-v0.md

    [DOC] roadmap refresh:
    CarbonStack_Long_Term_Roadmap_v0.4.14_REFRESH.pdf

    [DOC] docs index:
    carbonstack/docs/README.md

    [DOC] roadmap:
    carbonstack/roadmap/ROADMAP.md

    [DOC] top README:
    carbonstack/README.md

### Runner surfaces

    [RUNNER] root:
    carbonstack/tools/carbonstack-validate

    [RUNNER] dispatch:
    carbonstack/tools/carbonstack-validate/main.go

    [RUNNER] dev-runtime-openmls implementation:
    carbonstack/tools/carbonstack-validate/dev_runtime_openmls.go

    [RUNNER] runner README:
    carbonstack/tools/carbonstack-validate/README.md

    [RUNNER] local-cypher:
    carbonstack/tools/carbonstack-validate/local_cypher.go

    [RUNNER] release-snapshot:
    carbonstack/tools/carbonstack-validate/release_snapshot.go

    [RUNNER] checksums:
    carbonstack/tools/carbonstack-validate/checksums.go

### Comms runtime surfaces

    [COMMS] CLI entry:
    carbonstack-comms/cmd/comms/main.go

    [COMMS] command registry and non-OpenMLS commands:
    carbonstack-comms/internal/app/commands.go

    [COMMS] OpenMLS message runtime command helpers:
    carbonstack-comms/internal/app/openmls_runtime.go

    [COMMS] OpenMLS bootstrap wrapper helpers:
    carbonstack-comms/internal/app/openmls_bootstrap.go

    [COMMS] OpenMLS bootstrap wrapper tests:
    carbonstack-comms/internal/app/openmls_bootstrap_dev_test.go

    [COMMS] openmls-send-dev tests:
    carbonstack-comms/internal/app/openmls_send_dev_test.go

    [COMMS] openmls-inbox-dev tests:
    carbonstack-comms/internal/app/openmls_inbox_dev_test.go

    [COMMS] dev runtime smoke proof:
    carbonstack-comms/scripts/dev-openmls-runtime-smoke.sh

    [COMMS] OpenMLS sidecar dir:
    carbonstack-comms/internal/protocol/mls/openmls-sidecar

    [COMMS] sidecar README:
    carbonstack-comms/internal/protocol/mls/openmls-sidecar/README.md

### Current stub-era Comms functions / behaviors

    send:
      trust policy evaluation
      mock/stub provider encryption
      Cypher envelope submit

    inbox:
      Cypher inbox retrieval
      mock/stub provider plaintext display as stub_plaintext

    ack:
      envelope ack by envelope ID and recipient device ID

### Implemented v0.4.x dev OpenMLS commands

    openmls-send-dev:
      internal/app/openmls_runtime.go
      state.RequireReadyDevice
      trust.EvaluateSend
      runOpenMLSMessageProtect
      sidecar message-protect
      internal/relay.SubmitOpenMLSArtifactEnvelope
      ArtifactKindApplicationMessage
      dev/pre-alpha warning

    openmls-inbox-dev:
      internal/app/openmls_runtime.go
      state.RequireReadyDevice
      Cypher inbox
      skip unsupported envelopes
      internal/relay.WriteOpenMLSArtifactFromEnvelope
      runOpenMLSMessageOpen
      sidecar message-open
      print plaintext after sidecar success
      optional ack only after sidecar success

    openmls-identity-create-dev:
      internal/app/openmls_bootstrap.go
      sidecar identity-create
      required --sidecar-device-label
      optional --sidecar-dir
      stable key/value output
      no Comms state/trust mutation

    openmls-identity-status-dev:
      internal/app/openmls_bootstrap.go
      sidecar identity-status
      required --sidecar-device-label
      optional --sidecar-dir
      stable key/value output
      prints identity_exists when present
      no Comms state/trust mutation

    openmls-bundle-export-dev:
      internal/app/openmls_bootstrap.go
      sidecar public-bundle-export
      required --sidecar-device-label
      optional --sidecar-dir
      optional --write-artifact
      stable key/value output
      prints key_package_artifact_path_hint when present
      prints key_package_artifact_path when a hint is present
      no Comms state/trust mutation

    openmls-conversation-create-dev:
      internal/app/openmls_bootstrap.go
      sidecar conversation-create
      required --sidecar-device-label
      required --conversation
      stable key/value output
      prints sidecar_conversation_label
      prints conversation_state_path_hint when present
      prints conversation_summary_path_hint when present
      prints provider_storage_path_hint when present
      no Comms state/trust mutation

    openmls-conversation-load-check-dev:
      internal/app/openmls_bootstrap.go
      sidecar conversation-load-check
      required --sidecar-device-label
      required --conversation
      stable key/value output
      prints group_reloadable when present
      no Comms state/trust mutation

    openmls-conversation-add-member-dev:
      internal/app/openmls_bootstrap.go
      sidecar conversation-add-member
      required --sidecar-device-label
      required --conversation
      required --member-keypackage
      resolves --member-keypackage to absolute path before sidecar call
      stable key/value output
      prints welcome_artifact_path_hint and welcome_artifact_path when present
      prints welcome_manifest_path_hint and welcome_manifest_path when present
      prints welcome_artifact_sha256 and welcome_artifact_size_bytes when present
      prints member_added, welcome_artifact_written, group_reloadable, member counts, and epochs when present
      no Comms state/trust mutation

    openmls-conversation-join-dev:
      internal/app/openmls_bootstrap.go
      sidecar conversation-join
      required --sidecar-device-label
      required --conversation
      required --welcome
      resolves --welcome to absolute path before sidecar call
      stable key/value output
      prints joined, group_reloadable, member_count, epoch, join summary hint, conversation state hint, conversation summary hint, and provider storage hint when present
      no Comms state/trust mutation

### v0.4.5 through v0.4.15 runtime smoke proof

    scripts/dev-openmls-runtime-smoke.sh:
      temp workspace
      temp Cypher binary
      temp Cypher DB
      dynamic loopback Cypher port
      temp Alice/Bob Comms state
      sidecar state reset
      direct sidecar identity/conversation bootstrap
      openmls-send-dev application-message send
      openmls-inbox-dev application-message open + ack
      inbox-empty-after-ack verification
      robust cleanup with INT -> TERM -> KILL escalation

### v0.4.9/v0.4.10 runner profile

    go run . --profile dev-runtime-openmls --clean-generated

    Runner.DevRuntimeOpenMLS:
      CheckRequiredPaths
      CheckLiveGitUmbrella
      toolchain reporting
      ArtifactScan("pre-dev-runtime-openmls")
      RunStep for carbonstack-comms/scripts/dev-openmls-runtime-smoke.sh
      ArtifactScan("post-dev-runtime-openmls")
      explicit nonclaim summary

### Existing v0.4.0 package validation

    cd <package-root>/carbonstack/tools/carbonstack-validate
    go run . --profile verify-checksums --root <package-root>
    go run . --profile full --root <package-root> --clean-generated

---

## 10. v0.4.x Forward Plan

### Immediate next: v0.4.16 smoke-script migration recon / wrapper-based smoke variant decision

Expected:

    inspect the current direct-sidecar bootstrap in scripts/dev-openmls-runtime-smoke.sh.
    compare against the complete wrapper chain from v0.4.15.
    decide whether the next implementation should:
      add a parallel wrapper-based smoke script,
      migrate existing smoke bootstrap to wrappers,
      or keep direct sidecar smoke as the only committed smoke for now.
    preserve the direct smoke proof unless wrapper-based smoke proves equally stable.
    keep dev-runtime-openmls manual-only and outside full.
    keep local-backbone reserved.
    document the sidecar cwd/path lesson if smoke work touches sidecar execution.

### v0.4.17+

Likely:

    wrapper-based smoke variant, if v0.4.16 recon supports it.
    maybe runner/doc alignment after wrapper smoke passes.
    known-good command registry YAML/JSON outside the front-door README unless user-critical.
    exact sidecar setup runbook.
    clearer dev state reset semantics.
    better command output shape.
    decision on whether stub-era send/inbox are aliased, replaced, or deprecated.
    decision on when local-backbone naming stops being premature.

---

## 11. v0.5.x and Later Carry-Forward

### v0.5.x

Local storage, trust state, provider state, vault design, and PQ/hybrid ciphersuite migration readiness.

Do not start PQ/hybrid implementation before the runtime OpenMLS/Cypher path exists and is documented. PQ/hybrid migration touches provider state, ciphersuite metadata, trust display, artifact size, reinit/rekey/recovery policy, and claim boundaries.

### v0.6.x+

Hostile-server and abuse-resistance harnesses.

### v0.7.x+

Deployability and operations hardening.

### v0.8.x+

Documented self-pentest / adversarial validation.

### v0.9.x+

Claim-boundary review.

### v0.10.x+

Android backend/app work.

### v1.x.x

Public app/server major epoch and much later CarbonStackOS/appliance research, if justified.

---

## 12. Lean Breakpoint Summary

    CarbonStack v0.4.15 is the completed planned dev-only OpenMLS bootstrap wrapper implementation checkpoint. v0.4.0 remains the current public broad local deployability pre-release. carbonstack remains at 0583683 docs: define dev OpenMLS bootstrap command contract. carbonstack-comms is now at c2d3b16 feat: add dev OpenMLS add-member and join bootstrap commands. v0.4.15 extends internal/app/openmls_bootstrap.go and internal/app/openmls_bootstrap_dev_test.go, registers openmls-conversation-add-member-dev and openmls-conversation-join-dev, and updates README.md command docs. The current dev-only bootstrap wrapper set now covers identity create/status, public bundle export, conversation create/load-check, conversation add-member, and conversation join. The commands keep sidecar labels explicit, normalize sidecar JSON into stable Comms key/value output, resolve path-consuming inputs to absolute paths before sidecar invocation, preserve Welcome path/hash/size/membership fields, and do not mutate Comms state/trust files. Validation passed for app tests, full Comms tests, manual full wrapper-chain smoke, dev-runtime-openmls --clean-generated, local-cypher, doctor, and core --clean-generated. Final repo heads were clean across all four repos. Next safe rung: v0.4.16 smoke-script migration recon / wrapper-based smoke variant decision.

---

## 13. Preserved Immediate Previous Handoff: v0.4.14

The following is the previous v0.4.14 handoff. Where it conflicts with the v0.4.15 overlay above, v0.4.15 wins for current state.



# CarbonStack LogDoc v0.4.14

**Last updated:** 2026-06-05 local session  
**Owner/operator:** bitcrusher32  
**Assistant/model:** GPT-5.5 Thinking  
**Current phase:** Phase 2F / **v0.4.x dev-only OpenMLS bundle/conversation bootstrap wrapper implementation checkpoint complete**. The v0.4.0 public broad local deployability pre-release remains the current public release. v0.4.12 defined the strict dev-only OpenMLS bootstrap wrapper contract; v0.4.13 implemented the first identity wrapper subset; v0.4.14 now implements the next staged bootstrap wrapper subset in `carbonstack-comms`: `openmls-bundle-export-dev`, `openmls-conversation-create-dev`, and `openmls-conversation-load-check-dev`. `carbonstack` remains at `0583683 (HEAD -> main, origin/main, origin/HEAD) docs: define dev OpenMLS bootstrap command contract`; `carbonstack-comms` is now at `1f5c09e (HEAD -> main, origin/main, origin/HEAD) feat: add dev OpenMLS bundle and conversation bootstrap commands`; `carbonstack-cypher` remains at `9ab994c (HEAD -> main, origin/main, origin/HEAD) docs: point to local operator runbook`; `carbonstack-os` remains at `1bbbe52 (HEAD -> main, origin/main, origin/HEAD) docs: clarify CarbonStackOS target direction`.

**Version schema:** `v[scope].[timeline]`. This file is `v0.4.14`, the second dev-only OpenMLS bootstrap wrapper implementation checkpoint after the v0.4.13 identity wrapper checkpoint. Per LogDoc V2 practice, this Markdown file preserves current operational continuity and carries the v0.4.x runtime-integration thread forward; the JSON breakpoint remains the lean machine-readable handoff state.

**Compression note:** The detailed v0.3.0 through v0.3.36PRIME process ledger remains archived in PRIME/sanitized PRIME. v0.4.0 compressed that pre-release sprint; v0.4.1 through v0.4.12 preserved runtime recon, command contracts, dev OpenMLS send/inbox commands, smoke proof, runner profile, boundary checks, wrapper recon, and bootstrap command contract. v0.4.13 preserved the first implementation of identity create/status wrappers. v0.4.14 now preserves the next staged implementation: bundle export and conversation create/load-check wrappers.

---

## 1. Project Goal

**Active goal:** Continue moving CarbonStackComms from stub-era runtime `send`/`inbox` behavior toward explicit dev-only OpenMLS/Cypher-backed runtime commands and repeatable validation, while preserving claim discipline. v0.4.14 implements the second bootstrap wrapper subset and proves the wrapper pattern can extend beyond identity without migrating the smoke script or overclaiming local-backbone readiness.

CarbonStack’s broader project intent remains:

- **CarbonStack** is the secure-communications appliance-stack project and shared doctrine/release authority.
- **CarbonStackComms** is the text-first messaging client and current OpenMLS sidecar/relay integration surface.
- **CarbonStackCypher** is the self-hostable relay/storage stack.
- **CarbonStackOS** is future constrained Android-derived appliance OS work and remains out of current runnable validation packages.
- The `carbonstack` Gitea repo remains the source of truth for release state, release assets, docs, and public claims.
- GitHub mirrors exist as secondary push mirrors only; Gitea remains authoritative.

**North star:**

> Its goal is not to be a better smartphone. Its goal is to stop being a smartphone.

**Core doctrine:**

> Every feature is guilty until it proves it does not add unacceptable parser, network, sensor, identity, or filesystem authority.

**Current reality after v0.4.14:**

    openmls-send-dev is a dev-only OpenMLS send-side command.
    openmls-inbox-dev is a dev-only OpenMLS receive/open/optional-ack command.
    scripts/dev-openmls-runtime-smoke.sh proves the dev CLI application-message path:
      openmls-send-dev -> Cypher -> openmls-inbox-dev --ack
    tools/carbonstack-validate has a manual dev-runtime-openmls profile that wraps that smoke script.
    dev-runtime-openmls remains manual-only, live-git-umbrella-only, and outside full.
    openmls-identity-create-dev exists.
    openmls-identity-status-dev exists.
    openmls-bundle-export-dev now exists.
    openmls-conversation-create-dev now exists.
    openmls-conversation-load-check-dev now exists.
    Bootstrap wrappers live in carbonstack-comms/internal/app/openmls_bootstrap.go.
    Bootstrap wrapper tests live in carbonstack-comms/internal/app/openmls_bootstrap_dev_test.go.
    Bootstrap wrappers keep sidecar labels explicit.
    Bootstrap wrappers normalize sidecar JSON to stable key/value Comms output.
    Bundle export preserves key_package_artifact_path_hint and prints key_package_artifact_path when present.
    Conversation create preserves conversation/provider path hints.
    Conversation load-check prints group_reloadable when present.
    Bootstrap wrappers do not mutate Comms state/trust files.
    carbonstack-comms README documents the current dev-only OpenMLS bootstrap commands and boundaries.
    Existing send/inbox remain stub-era.
    Add-member/join wrappers are not implemented yet.
    The smoke script still uses direct sidecar bootstrap.
    local-backbone remains reserved.

Hard nonclaims remain: v0.4.14 does **not** prove production readiness, production E2EE, hostile-server safety, metadata privacy, Android readiness, CarbonStackOS readiness, secure vault/storage, public ingress safety, real homelab validation, audit, certification, PQ/hybrid security, mature messenger UX, release-package runtime validation, Relay Space join/onboarding, add-member/join wrapper coverage, smoke-script wrapper migration, or local-backbone readiness.

---

## 2. Current Public Release State

Public release:

    https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.4.0

Release title:

    CarbonStack v0.4.0 Broad Local Deployability Pre-Release

Release tag:

    v0.4.0

Release commit:

    24ac9fc494 / 24ac9fc docs: record v0.4.0 package rehearsal

Gitea release status:

    Pre-Release

Important release-page warning remains:

    The default Gitea Source Code ZIP/TAR.GZ downloads are auto-generated carbonstack-only archives.
    The intended runnable validation artifact is the attached multi-repo release package plus companion assets.

Post-release public surface:

    carbonstack README is evergreen.
    It points interested parties to the Releases page for runnable / known-good artifacts.
    Gitea is source of truth.
    GitHub mirrors are secondary push mirrors, not release authority.

---

## 3. Current Repo Heads

    carbonstack        0583683 (HEAD -> main, origin/main, origin/HEAD) docs: define dev OpenMLS bootstrap command contract
    carbonstack-comms  1f5c09e (HEAD -> main, origin/main, origin/HEAD) feat: add dev OpenMLS bundle and conversation bootstrap commands
    carbonstack-cypher 9ab994c (HEAD -> main, origin/main, origin/HEAD) docs: point to local operator runbook
    carbonstack-os     1bbbe52 (HEAD -> main, origin/main, origin/HEAD) docs: clarify CarbonStackOS target direction

Current public release tag:

    v0.4.0 -> 24ac9fc494 / 24ac9fc docs: record v0.4.0 package rehearsal

Recent runtime/thread commits:

    v0.4.8 extraction:
      0a48ae1 refactor: split OpenMLS runtime command helpers

    v0.4.9 runner profile:
      8eeadb2 feat: add dev OpenMLS runtime validation profile

    v0.4.10 profile boundary:
      7384f43 docs: record dev OpenMLS runtime profile boundary

    v0.4.11 wrapper recon:
      8ee3ebb docs: plan dev OpenMLS bootstrap wrappers

    v0.4.12 command contract:
      0583683 docs: define dev OpenMLS bootstrap command contract

    v0.4.13 identity wrapper implementation:
      20c4cc9 feat: add dev OpenMLS identity bootstrap commands

    Current v0.4.14 bundle/conversation wrapper implementation:
      1f5c09e feat: add dev OpenMLS bundle and conversation bootstrap commands

Continuity note:

    `1f5c09e` is a carbonstack-comms implementation commit after the v0.4.0 public release tag.
    It does not retag v0.4.0.
    The release remains attached to `24ac9fc494`.

---

## 4. Current Validated / Known State

Validated / known at the v0.4.14 checkpoint:

    [RELEASE] v0.4.0 Gitea release exists, is marked Pre-Release, and is attached to 24ac9fc494.
    [RUNNER] full means release-snapshot followed by local-cypher.
    [RUNNER] dev-runtime-openmls validates the dev/pre-alpha OpenMLS application-message runtime CLI smoke proof.
    [RUNNER] dev-runtime-openmls remains manual-only and live-umbrella-only.
    [RUNNER] dev-runtime-openmls is not included in full.
    [DOCS] docs/162-openmls-bootstrap-command-contract-v0.md records the v0.4.12 dev-only OpenMLS bootstrap command contract.
    [COMMS] openmls-send-dev exists.
    [COMMS] openmls-inbox-dev exists.
    [COMMS] openmls-identity-create-dev exists.
    [COMMS] openmls-identity-status-dev exists.
    [COMMS] openmls-bundle-export-dev now exists.
    [COMMS] openmls-conversation-create-dev now exists.
    [COMMS] openmls-conversation-load-check-dev now exists.
    [COMMS] internal/app/openmls_bootstrap.go contains identity, bundle, conversation-create, and conversation-load-check wrappers.
    [COMMS] internal/app/openmls_bootstrap_dev_test.go tests the current bootstrap wrapper set with injected sidecar seams.
    [COMMS] internal/app/commands.go registers the new wrapper commands and lists them in usage.
    [COMMS] README.md documents the current dev-only OpenMLS bootstrap commands and their boundary.
    [COMMS] scripts/dev-openmls-runtime-smoke.sh still uses direct sidecar bootstrap.
    [COMMS] future add-member/join wrappers remain unimplemented.
    [COMMS] existing send/inbox/ack remain in place and stub-era.
    [MIRRORS] all four repos have GitHub push-mirror functionality, with Gitea remaining source of truth.

v0.4.14 validation result from the final log:

    go test ./internal/app -count=1 passed in carbonstack-comms.
    go test ./... -count=1 -timeout 600s passed in carbonstack-comms.
    Manual identity create/status wrapper checks passed.
    Manual bundle export wrapper check passed and printed key_package_artifact_path_hint plus key_package_artifact_path.
    Manual conversation create wrapper check passed and printed conversation_state_path_hint, conversation_summary_path_hint, and provider_storage_path_hint.
    Manual conversation load-check wrapper check passed and printed group_reloadable: true.
    go test ./... -count=1 passed in carbonstack/tools/carbonstack-validate.
    go run . --profile dev-runtime-openmls --clean-generated passed.
    local-cypher passed.
    doctor passed.
    core --clean-generated passed.
    carbonstack-comms was committed and pushed at:
      1f5c09e feat: add dev OpenMLS bundle and conversation bootstrap commands
    final repo status showed all four repos clean.

Expected live umbrella validation commands now include:

    cd ~/repos/carbonstack_umbrella/carbonstack-comms
    go test ./internal/app -count=1
    go test ./... -count=1 -timeout 600s

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go run . --profile dev-runtime-openmls --clean-generated
    go run . --profile local-cypher
    go run . --profile doctor
    go run . --profile core --clean-generated

New manual v0.4.14 wrapper checks:

    cd ~/repos/carbonstack_umbrella/carbonstack-comms
    rm -rf internal/protocol/mls/openmls-sidecar/.carbonstack-openmls-sidecar-state
    go run ./cmd/comms openmls-identity-create-dev --sidecar-device-label carbonstack-v0414-alice
    go run ./cmd/comms openmls-identity-status-dev --sidecar-device-label carbonstack-v0414-alice
    go run ./cmd/comms openmls-bundle-export-dev --sidecar-device-label carbonstack-v0414-alice --write-artifact
    go run ./cmd/comms openmls-conversation-create-dev --sidecar-device-label carbonstack-v0414-alice --conversation carbonstack-v0414-conversation
    go run ./cmd/comms openmls-conversation-load-check-dev --sidecar-device-label carbonstack-v0414-alice --conversation carbonstack-v0414-conversation

Expected v0.4.0 public package validation commands remain unchanged:

    cd <package-root>/carbonstack/tools/carbonstack-validate
    go run . --profile verify-checksums --root <package-root>
    go run . --profile full --root <package-root> --clean-generated

---

## 5. v0.4.14 Work Completed

### 5.1 Second bootstrap wrapper subset was implemented

New commands implemented in carbonstack-comms:

    openmls-bundle-export-dev
    openmls-conversation-create-dev
    openmls-conversation-load-check-dev

Together with the v0.4.13 identity wrappers, the currently implemented dev-only bootstrap wrapper set is now:

    openmls-identity-create-dev
    openmls-identity-status-dev
    openmls-bundle-export-dev
    openmls-conversation-create-dev
    openmls-conversation-load-check-dev

These commands are intentionally dev-only and sidecar-bootstrap-scoped.

They do not:

    replace send/inbox.
    mutate Comms state.json.
    mutate trust.json.
    mutate trust-events.jsonl.
    call Cypher.
    implement Relay Space join UX.
    imply local-backbone.
    imply production identity or conversation UX.

### 5.2 openmls_bootstrap.go was extended

Updated file:

    carbonstack-comms/internal/app/openmls_bootstrap.go

New command functions:

    cmdOpenMLSBundleExportDev
    cmdOpenMLSConversationCreateDev
    cmdOpenMLSConversationLoadCheckDev

New helper behavior:

    bootstrapPrintOptionalString
    bootstrapPathFromHint

Implementation behavior:

    openmls-bundle-export-dev calls sidecar public-bundle-export.
    openmls-bundle-export-dev accepts --write-artifact.
    openmls-bundle-export-dev prints key_package_artifact_path_hint when present.
    openmls-bundle-export-dev resolves and prints key_package_artifact_path when a hint is present.
    openmls-conversation-create-dev calls sidecar conversation-create.
    openmls-conversation-create-dev requires --sidecar-device-label and --conversation.
    openmls-conversation-create-dev preserves conversation_state_path_hint, conversation_summary_path_hint, and provider_storage_path_hint when present.
    openmls-conversation-load-check-dev calls sidecar conversation-load-check.
    openmls-conversation-load-check-dev requires --sidecar-device-label and --conversation.
    openmls-conversation-load-check-dev prints group_reloadable when present.
    all new wrappers print stable key/value output and dev/pre-alpha warnings.

### 5.3 openmls_bootstrap_dev_test.go was extended

Updated file:

    carbonstack-comms/internal/app/openmls_bootstrap_dev_test.go

New tests cover:

    openmls-bundle-export-dev requires --sidecar-device-label.
    openmls-conversation-create-dev requires both --sidecar-device-label and --conversation.
    openmls-conversation-load-check-dev requires both --sidecar-device-label and --conversation.
    bundle export invokes public-bundle-export with expected args.
    bundle export prints stable output and resolved key_package_artifact_path.
    conversation create invokes conversation-create with expected args.
    conversation create prints stable output and sidecar path hints.
    conversation load-check invokes conversation-load-check with expected args.
    conversation load-check prints stable output and group_reloadable when present.

Test design still follows the v0.4.12 contract:

    app tests use injected seams.
    app tests do not run the Rust sidecar directly.
    real sidecar execution remains in manual wrapper smoke, protocol tests, smoke script, and runner validation.

### 5.4 commands.go was updated

Updated file:

    carbonstack-comms/internal/app/commands.go

New dispatch cases:

    openmls-bundle-export-dev
    openmls-conversation-create-dev
    openmls-conversation-load-check-dev

Usage output now lists:

    openmls-bundle-export-dev
    openmls-conversation-create-dev
    openmls-conversation-load-check-dev

The existing command registry remains in commands.go. Message runtime helpers remain in openmls_runtime.go. Bootstrap wrapper helpers remain in openmls_bootstrap.go.

### 5.5 carbonstack-comms README was updated

Updated file:

    carbonstack-comms/README.md

The `Dev-only OpenMLS bootstrap commands` section now documents:

    openmls-identity-create-dev
    openmls-identity-status-dev
    openmls-bundle-export-dev
    openmls-conversation-create-dev
    openmls-conversation-load-check-dev

It includes example commands for identity create/status, bundle export, conversation create, and conversation load-check.

Boundary preserved:

    sidecar labels are explicit for now.
    Comms state/trust files are not mutated by these wrappers.
    existing send/inbox remain stub-era.
    dev-runtime-openmls remains the current manual smoke-profile proof.
    these commands are not production identity UX, not Relay Space join UX, not local-backbone, and not secure vault/state management.

### 5.6 Manual wrapper smoke passed

Manual check:

    rm -rf internal/protocol/mls/openmls-sidecar/.carbonstack-openmls-sidecar-state

    go run ./cmd/comms openmls-identity-create-dev --sidecar-device-label carbonstack-v0414-alice
    go run ./cmd/comms openmls-identity-status-dev --sidecar-device-label carbonstack-v0414-alice
    go run ./cmd/comms openmls-bundle-export-dev --sidecar-device-label carbonstack-v0414-alice --write-artifact
    go run ./cmd/comms openmls-conversation-create-dev --sidecar-device-label carbonstack-v0414-alice --conversation carbonstack-v0414-conversation
    go run ./cmd/comms openmls-conversation-load-check-dev --sidecar-device-label carbonstack-v0414-alice --conversation carbonstack-v0414-conversation

Observed output included:

    command: openmls-bundle-export-dev
    status: exported
    sidecar_command: public-bundle-export
    sidecar_device_label: carbonstack-v0414-alice
    key_package_artifact_path_hint: .carbonstack-openmls-sidecar-state/dev/devices/carbonstack-v0414-alice/public-bundle.keypackage.bin
    key_package_artifact_path: internal/protocol/mls/openmls-sidecar/.carbonstack-openmls-sidecar-state/dev/devices/carbonstack-v0414-alice/public-bundle.keypackage.bin

and:

    command: openmls-conversation-create-dev
    status: created
    sidecar_command: conversation-create
    sidecar_device_label: carbonstack-v0414-alice
    sidecar_conversation_label: carbonstack-v0414-conversation
    conversation_state_path_hint: .carbonstack-openmls-sidecar-state/dev/devices/carbonstack-v0414-alice/conversations/carbonstack-v0414-conversation
    conversation_summary_path_hint: .carbonstack-openmls-sidecar-state/dev/devices/carbonstack-v0414-alice/conversations/carbonstack-v0414-conversation/conversation-summary.json
    provider_storage_path_hint: .carbonstack-openmls-sidecar-state/dev/devices/carbonstack-v0414-alice/conversations/carbonstack-v0414-conversation/provider-storage.json

and:

    command: openmls-conversation-load-check-dev
    status: loaded
    sidecar_command: conversation-load-check
    sidecar_device_label: carbonstack-v0414-alice
    sidecar_conversation_label: carbonstack-v0414-conversation
    group_reloadable: true

### 5.7 Validation passed

Commands observed passing:

    cd ~/repos/carbonstack_umbrella/carbonstack-comms
    go test ./internal/app -count=1
    go test ./... -count=1 -timeout 600s

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go test ./... -count=1
    go run . --profile dev-runtime-openmls --clean-generated
    go run . --profile local-cypher
    go run . --profile doctor
    go run . --profile core --clean-generated

Final commit/push:

    carbonstack-comms:
    1f5c09e feat: add dev OpenMLS bundle and conversation bootstrap commands

Final repo snapshot:

    carbonstack        0583683 docs: define dev OpenMLS bootstrap command contract
    carbonstack-comms  1f5c09e feat: add dev OpenMLS bundle and conversation bootstrap commands
    carbonstack-cypher 9ab994c docs: point to local operator runbook
    carbonstack-os     1bbbe52 docs: clarify CarbonStackOS target direction

---

## 6. Critical Blunders / Lessons to Carry Forward

### 6.1 Manual wrapper smoke again left known generated roots before dev-runtime-openmls

As in v0.4.13, the manual wrapper smoke intentionally created OpenMLS sidecar state and Rust build artifacts. During the later `dev-runtime-openmls` run, the pre-profile artifact scan reported known generated roots:

    .carbonstack-openmls-sidecar-state
    target

This stayed inside known generated paths and was cleaned by `--clean-generated`.

Lesson:

    Manual wrapper smoke can leave expected sidecar state and target roots.
    This is acceptable only when they stay in known generated paths.
    Always follow with --clean-generated when preparing a clean breakpoint.

### 6.2 Commit metadata warning remains a polish issue

The commit log again showed auto-generated local commit metadata:

    Committer: bitcrusher32 <▓▓>

This is not a functional validation failure, but remains a professionalism/publishing hygiene issue to address eventually with `git config --global user.name` and `git config --global user.email` if desired.

### 6.3 Wrapper contract now holds for identity, bundle export, conversation create, and load-check — but not add-member/join yet

v0.4.14 proves the contract shape for five bootstrap commands:

    openmls-identity-create-dev
    openmls-identity-status-dev
    openmls-bundle-export-dev
    openmls-conversation-create-dev
    openmls-conversation-load-check-dev

Do not assume:

    conversation add-member wrapper exists.
    conversation join wrapper exists.
    smoke script uses wrappers.
    full bootstrap wrapper coverage exists.
    local-backbone is justified.

Correct interpretation:

    second bootstrap wrapper subset exists and passed tests.
    member/welcome wrappers still need staged implementation.

---

## 7. Current Blockers / Not Validated

Current blockers and nonvalidated areas after v0.4.14:

    Existing send/inbox remain stub-era.
    Runtime Comms has dev-only OpenMLS commands, a smoke proof, and a manual runner profile, but no mature user-facing UX.
    Sidecar KeyPackage/Welcome/bootstrap setup is still direct dev setup inside the smoke script.
    Conversation add-member/join wrappers are not implemented yet.
    Smoke script has not been migrated to wrappers.
    No wrapper-based smoke variant exists yet.
    No Relay Space join model exists yet.
    local-backbone runner profile does not exist and remains reserved.
    dev-runtime-openmls remains manual-only and live-git-umbrella-only.
    dev-runtime-openmls is not included in full.
    dev-runtime-openmls is not part of public release package validation.
    Fresh release-package behavior for dev-runtime-openmls has not been tested as a supported validation surface.
    Post-release public download verification is still useful if not already performed.
    Sanitized PRIME placement in carbonstack/logdoc-loop-system still needs a policy decision.
    Production-safe provider storage and secure vault design remain v0.5.x future work.
    PQ/hybrid ciphersuite migration readiness remains v0.5.x future work.
    Hostile-server rollback/replay/metadata-abuse validation remains future work.
    Android/CarbonStackOS implementation remains later work.
    CarbonStack Relay Space mechanics remain a future concept, not an implementation.

Hard nonclaims remain:

    Do not claim production readiness.
    Do not claim production E2EE.
    Do not claim hostile-server safety.
    Do not claim metadata privacy.
    Do not claim Debian production deployability.
    Do not claim systemd readiness.
    Do not claim cloudflared readiness.
    Do not claim Android readiness.
    Do not claim CarbonStackOS readiness.
    Do not claim secure vault/storage.
    Do not claim real homelab validation.
    Do not claim external audit or certification.
    Do not claim CarbonStack Relay Space mechanics are implemented.
    Do not claim local-backbone exists or is justified yet.
    Do not claim local-cypher validates runtime Comms UX.
    Do not claim full validates local-backbone.
    Do not claim full is a deployment command.
    Do not claim dev-runtime-openmls is release-package validation yet.
    Do not claim send/inbox are OpenMLS-backed yet.
    Do not claim PQ/hybrid security or quantum-safe messaging.
    Do not claim all bootstrap wrapper commands exist yet.
    Do not let GitHub mirrors supersede Gitea source-of-truth status.

Allowed claim:

    CarbonStackComms now has dev-only openmls-send-dev and openmls-inbox-dev commands.
    CarbonStackComms now has a dev runtime smoke script proving openmls-send-dev -> Cypher -> openmls-inbox-dev --ack.
    CarbonStackComms now has OpenMLS runtime command helper glue split into internal/app/openmls_runtime.go.
    CarbonStackComms now has dev-only OpenMLS bootstrap wrappers for identity create/status, bundle export, conversation create, and conversation load-check.
    The bootstrap wrappers parse sidecar JSON, print stable key/value output, and keep sidecar labels explicit.
    The bootstrap wrappers do not mutate Comms state/trust files.
    carbonstack/tools/carbonstack-validate has a manual dev-runtime-openmls profile.
    dev-runtime-openmls wraps the comms smoke script and validates the dev/pre-alpha OpenMLS application-message CLI path.
    dev-runtime-openmls is live-git-umbrella-only and not included in full.
    The main carbonstack repo records the smoke proof, pre-local-backbone assessment, manual dev-runtime-openmls runner profile, profile boundary check, bootstrap wrapper recon/planning doc, and bootstrap wrapper command contract.
    Existing send/inbox remain stub-era.
    v0.4.0 remains the current public broad local deployability pre-release.
    Gitea remains source of truth; GitHub mirrors are secondary push mirrors.

---

## 8. Next Safest Actions

Recommended next rung:

    refresh the long-term roadmap artifact after v0.4.14

Rationale:

    The uploaded v0.4.0 roadmap is now stale in its v0.4.x details.
    v0.4.1-v0.4.14 substantially changed the runtime integration state.
    The later v0.5.x+ structure remains broadly useful, but it should be refreshed against the real implemented v0.4.x arc.

Likely roadmap refresh should record:

    v0.4.1 runtime recon.
    v0.4.2 runtime OpenMLS command contract.
    v0.4.3 openmls-send-dev.
    v0.4.4 openmls-inbox-dev.
    v0.4.5 dev runtime smoke proof.
    v0.4.6 main repo smoke proof alignment.
    v0.4.7 pre-local-backbone assessment.
    v0.4.8 OpenMLS runtime helper extraction.
    v0.4.9 manual dev-runtime-openmls runner profile.
    v0.4.10 profile boundary evidence.
    v0.4.11 bootstrap wrapper recon.
    v0.4.12 bootstrap command contract.
    v0.4.13 identity create/status wrappers.
    v0.4.14 bundle export + conversation create/load-check wrappers.
    near next: v0.4.15 add-member/join wrappers.
    near next: v0.4.16 smoke-script migration recon / wrapper-based smoke variant decision.
    future late v0.4.x: known-good command registry YAML/JSON outside the front-door README unless user-critical.
    v0.5.x: state/trust/vault/PQ readiness remains after runtime/bootstrap wrapper path matures.

Next code rung after roadmap refresh:

    v0.4.15 implement add-member + join wrappers

Focus:

    carbonstack-comms only.
    implement openmls-conversation-add-member-dev.
    implement openmls-conversation-join-dev.
    preserve explicit sidecar labels and stable key/value output.
    preserve welcome_artifact_path_hint and welcome_artifact_path where applicable.
    accept explicit --member-keypackage and --welcome paths.
    do not mutate Comms state/trust.
    do not migrate the smoke script yet.
    validate app tests, full Comms tests, manual wrapper checks, dev-runtime-openmls, local-cypher, doctor, and core.

Avoid next:

    replacing the smoke script before add-member/join wrappers exist and pass tests.
    adding dev-runtime-openmls to full.
    calling any wrapper local-backbone.
    using Relay Space join naming for dev bootstrap wrappers.
    moving dev-runtime-openmls into public release validation without fresh-root behavior checks.
    public ingress.
    cloudflared.
    systemd.
    real homelab deployment.
    Android app.
    CarbonStackOS implementation.
    CarbonStack Relay Space implementation.
    PQ/hybrid ciphersuite implementation.
    production/security claims.

---

## 9. Critical Paths / Functions

### Source-of-truth public surfaces

    [GITEA SOURCE OF TRUTH] https://git.bitcrusher32.win/bitcrusher32/carbonstack
    [RELEASES] https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases
    [CURRENT PUBLIC RELEASE] https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.4.0

GitHub mirrors:

    secondary push mirrors for discoverability/redundancy only.
    Do not treat mirrors as release authority unless project policy changes.

### WSL Debian active working umbrella

    [WSL-DEBIAN] umbrella: ~/repos/carbonstack_umbrella
    [WSL-DEBIAN] carbonstack: ~/repos/carbonstack_umbrella/carbonstack
    [WSL-DEBIAN] carbonstack-comms: ~/repos/carbonstack_umbrella/carbonstack-comms
    [WSL-DEBIAN] carbonstack-cypher: ~/repos/carbonstack_umbrella/carbonstack-cypher
    [WSL-DEBIAN] carbonstack-os: ~/repos/carbonstack_umbrella/carbonstack-os
    [GO-RUNNER] runner dir: ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate

### Runtime integration doc surfaces

    [DOC] v0.4.1 runtime recon:
    carbonstack/docs/155-runtime-comms-openmls-cypher-recon-v0.md

    [DOC] v0.4.2 command contract:
    carbonstack/docs/156-runtime-openmls-command-contract-v0.md

    [DOC] v0.4.6 smoke proof record:
    carbonstack/docs/157-dev-runtime-openmls-smoke-proof-v0.md

    [DOC] v0.4.7 pre-local-backbone assessment:
    carbonstack/docs/158-pre-local-backbone-assessment-v0.md

    [DOC] v0.4.9 runner profile:
    carbonstack/docs/159-dev-runtime-openmls-runner-profile-v0.md

    [DOC] v0.4.10 profile boundary:
    carbonstack/docs/160-dev-runtime-openmls-profile-boundary-v0.md

    [DOC] v0.4.11 bootstrap wrapper recon:
    carbonstack/docs/161-openmls-bootstrap-wrapper-recon-v0.md

    [DOC] v0.4.12 bootstrap command contract:
    carbonstack/docs/162-openmls-bootstrap-command-contract-v0.md

    [DOC] docs index:
    carbonstack/docs/README.md

    [DOC] roadmap:
    carbonstack/roadmap/ROADMAP.md

    [DOC] top README:
    carbonstack/README.md

### Runner surfaces

    [RUNNER] root:
    carbonstack/tools/carbonstack-validate

    [RUNNER] dispatch:
    carbonstack/tools/carbonstack-validate/main.go

    [RUNNER] dev-runtime-openmls implementation:
    carbonstack/tools/carbonstack-validate/dev_runtime_openmls.go

    [RUNNER] runner README:
    carbonstack/tools/carbonstack-validate/README.md

    [RUNNER] local-cypher:
    carbonstack/tools/carbonstack-validate/local_cypher.go

    [RUNNER] release-snapshot:
    carbonstack/tools/carbonstack-validate/release_snapshot.go

    [RUNNER] checksums:
    carbonstack/tools/carbonstack-validate/checksums.go

### Comms runtime surfaces

    [COMMS] CLI entry:
    carbonstack-comms/cmd/comms/main.go

    [COMMS] command registry and non-OpenMLS commands:
    carbonstack-comms/internal/app/commands.go

    [COMMS] OpenMLS message runtime command helpers:
    carbonstack-comms/internal/app/openmls_runtime.go

    [COMMS] OpenMLS bootstrap wrapper helpers:
    carbonstack-comms/internal/app/openmls_bootstrap.go

    [COMMS] OpenMLS bootstrap wrapper tests:
    carbonstack-comms/internal/app/openmls_bootstrap_dev_test.go

    [COMMS] openmls-send-dev tests:
    carbonstack-comms/internal/app/openmls_send_dev_test.go

    [COMMS] openmls-inbox-dev tests:
    carbonstack-comms/internal/app/openmls_inbox_dev_test.go

    [COMMS] dev runtime smoke proof:
    carbonstack-comms/scripts/dev-openmls-runtime-smoke.sh

    [COMMS] OpenMLS sidecar dir:
    carbonstack-comms/internal/protocol/mls/openmls-sidecar

    [COMMS] sidecar README:
    carbonstack-comms/internal/protocol/mls/openmls-sidecar/README.md

### Current stub-era Comms functions / behaviors

    send:
      trust policy evaluation
      mock/stub provider encryption
      Cypher envelope submit

    inbox:
      Cypher inbox retrieval
      mock/stub provider plaintext display as stub_plaintext

    ack:
      envelope ack by envelope ID and recipient device ID

### Implemented v0.4.x dev OpenMLS commands

    openmls-send-dev:
      internal/app/openmls_runtime.go
      state.RequireReadyDevice
      trust.EvaluateSend
      runOpenMLSMessageProtect
      sidecar message-protect
      internal/relay.SubmitOpenMLSArtifactEnvelope
      ArtifactKindApplicationMessage
      dev/pre-alpha warning

    openmls-inbox-dev:
      internal/app/openmls_runtime.go
      state.RequireReadyDevice
      Cypher inbox
      skip unsupported envelopes
      internal/relay.WriteOpenMLSArtifactFromEnvelope
      runOpenMLSMessageOpen
      sidecar message-open
      print plaintext after sidecar success
      optional ack only after sidecar success

    openmls-identity-create-dev:
      internal/app/openmls_bootstrap.go
      sidecar identity-create
      required --sidecar-device-label
      optional --sidecar-dir
      stable key/value output
      no Comms state/trust mutation

    openmls-identity-status-dev:
      internal/app/openmls_bootstrap.go
      sidecar identity-status
      required --sidecar-device-label
      optional --sidecar-dir
      stable key/value output
      prints identity_exists when present
      no Comms state/trust mutation

    openmls-bundle-export-dev:
      internal/app/openmls_bootstrap.go
      sidecar public-bundle-export
      required --sidecar-device-label
      optional --sidecar-dir
      optional --write-artifact
      stable key/value output
      prints key_package_artifact_path_hint when present
      prints key_package_artifact_path when a hint is present
      no Comms state/trust mutation

    openmls-conversation-create-dev:
      internal/app/openmls_bootstrap.go
      sidecar conversation-create
      required --sidecar-device-label
      required --conversation
      stable key/value output
      prints sidecar_conversation_label
      prints conversation_state_path_hint when present
      prints conversation_summary_path_hint when present
      prints provider_storage_path_hint when present
      no Comms state/trust mutation

    openmls-conversation-load-check-dev:
      internal/app/openmls_bootstrap.go
      sidecar conversation-load-check
      required --sidecar-device-label
      required --conversation
      stable key/value output
      prints group_reloadable when present
      no Comms state/trust mutation

### Contracted but not yet implemented v0.4.15+ dev OpenMLS bootstrap wrappers

    openmls-conversation-add-member-dev:
      sidecar conversation-add-member
      required --sidecar-device-label, --conversation, and --member-keypackage

    openmls-conversation-join-dev:
      sidecar conversation-join
      required --sidecar-device-label, --conversation, and --welcome

### v0.4.5 through v0.4.14 runtime smoke proof

    scripts/dev-openmls-runtime-smoke.sh:
      temp workspace
      temp Cypher binary
      temp Cypher DB
      dynamic loopback Cypher port
      temp Alice/Bob Comms state
      sidecar state reset
      direct sidecar identity/conversation bootstrap
      openmls-send-dev application-message send
      openmls-inbox-dev application-message open + ack
      inbox-empty-after-ack verification
      robust cleanup with INT -> TERM -> KILL escalation

### v0.4.9/v0.4.10 runner profile

    go run . --profile dev-runtime-openmls --clean-generated

    Runner.DevRuntimeOpenMLS:
      CheckRequiredPaths
      CheckLiveGitUmbrella
      toolchain reporting
      ArtifactScan("pre-dev-runtime-openmls")
      RunStep for carbonstack-comms/scripts/dev-openmls-runtime-smoke.sh
      ArtifactScan("post-dev-runtime-openmls")
      explicit nonclaim summary

### Existing v0.4.0 package validation

    cd <package-root>/carbonstack/tools/carbonstack-validate
    go run . --profile verify-checksums --root <package-root>
    go run . --profile full --root <package-root> --clean-generated

---

## 10. v0.4.x Forward Plan

### Immediate next: roadmap refresh after v0.4.14

Expected:

    update the latest long-term roadmap artifact to reflect v0.4.1-v0.4.14 reality.
    preserve the v0.5.x state/trust/vault/PQ direction.
    preserve v0.6.x hostile-server harness direction.
    preserve v0.7.x deployability/ops-hardening direction.
    clarify that v0.4.x now includes dev runtime proof, manual runner profile, and partial bootstrap wrapper implementation.
    keep local-backbone reserved.
    keep dev-runtime-openmls outside full.
    carry forward known-good command registry YAML/JSON as late-v0.4.x or pre-v0.5.x idea.

### v0.4.15 — add-member/join wrapper implementation

Expected after roadmap refresh:

    carbonstack-comms-only implementation.
    extend internal/app/openmls_bootstrap.go.
    extend internal/app/openmls_bootstrap_dev_test.go.
    register openmls-conversation-add-member-dev.
    register openmls-conversation-join-dev.
    keep command registry in commands.go.
    keep output as stable key/value text.
    preserve welcome path hints and absolute paths where applicable.
    do not mutate Comms state/trust.
    update carbonstack-comms README if implementation lands.
    validate app tests, full Comms tests, manual wrapper checks, dev-runtime-openmls, local-cypher, doctor, and core.

### v0.4.16+

Likely:

    smoke-script migration recon / wrapper-based smoke variant decision.
    maybe wrapper-based smoke variant before replacing direct-sidecar bootstrap.
    maybe known-good command registry YAML/JSON after wrapper surface stabilizes.
    exact sidecar setup runbook.
    clearer dev state reset semantics.
    better command output shape.
    decision on whether stub-era send/inbox are aliased, replaced, or deprecated.
    decision on when local-backbone naming stops being premature.

---

## 11. v0.5.x and Later Carry-Forward

### v0.5.x

Local storage, trust state, provider state, vault design, and PQ/hybrid ciphersuite migration readiness.

Do not start PQ/hybrid implementation before the runtime OpenMLS/Cypher path exists and is documented. PQ/hybrid migration touches provider state, ciphersuite metadata, trust display, artifact size, reinit/rekey/recovery policy, and claim boundaries.

### v0.6.x+

Hostile-server and abuse-resistance harnesses.

### v0.7.x+

Deployability and operations hardening.

### v0.8.x+

Documented self-pentest / adversarial validation.

### v0.9.x+

Claim-boundary review.

### v0.10.x+

Android backend/app work.

### v1.x.x

Public app/server major epoch and much later CarbonStackOS/appliance research, if justified.

---

## 12. Lean Breakpoint Summary

    CarbonStack v0.4.14 is the second dev-only OpenMLS bootstrap wrapper implementation checkpoint. v0.4.0 remains the current public broad local deployability pre-release. carbonstack remains at 0583683 docs: define dev OpenMLS bootstrap command contract. carbonstack-comms is now at 1f5c09e feat: add dev OpenMLS bundle and conversation bootstrap commands. v0.4.14 extends internal/app/openmls_bootstrap.go and internal/app/openmls_bootstrap_dev_test.go, registers openmls-bundle-export-dev, openmls-conversation-create-dev, and openmls-conversation-load-check-dev, and updates README.md command docs. The current dev-only bootstrap wrapper set now covers identity create/status, public bundle export, conversation create, and conversation load-check. The commands keep sidecar labels explicit, normalize sidecar JSON into stable Comms key/value output, preserve key path hints, and do not mutate Comms state/trust files. Validation passed for app tests, full Comms tests, manual wrapper smoke, dev-runtime-openmls --clean-generated, local-cypher, doctor, and core --clean-generated. Final repo heads were clean across all four repos. Next safe rung: refresh the long-term roadmap artifact, then v0.4.15 add-member/join wrappers.

---

## 13. Preserved Immediate Previous Handoff: v0.4.13

The following is the previous v0.4.13 handoff. Where it conflicts with the v0.4.14 overlay above, v0.4.14 wins for current state.


# CarbonStack LogDoc v0.4.13

**Last updated:** 2026-06-05 local session  
**Owner/operator:** bitcrusher32  
**Assistant/model:** GPT-5.5 Thinking  
**Current phase:** Phase 2F / **v0.4.x dev-only OpenMLS identity bootstrap wrapper implementation checkpoint complete**. The v0.4.0 public broad local deployability pre-release remains the current public release. v0.4.12 defined the strict dev-only OpenMLS bootstrap wrapper contract; v0.4.13 now implements the first minimal wrapper subset in `carbonstack-comms`: `openmls-identity-create-dev` and `openmls-identity-status-dev`. `carbonstack` remains at `0583683 (HEAD -> main, origin/main, origin/HEAD) docs: define dev OpenMLS bootstrap command contract`; `carbonstack-comms` is now at `20c4cc9 (HEAD -> main, origin/main, origin/HEAD) feat: add dev OpenMLS identity bootstrap commands`; `carbonstack-cypher` remains at `9ab994c (HEAD -> main, origin/main, origin/HEAD) docs: point to local operator runbook`; `carbonstack-os` remains at `1bbbe52 (HEAD -> main, origin/main, origin/HEAD) docs: clarify CarbonStackOS target direction`.

**Version schema:** `v[scope].[timeline]`. This file is `v0.4.13`, the first dev-only OpenMLS bootstrap wrapper implementation checkpoint after the v0.4.12 command-contract checkpoint. Per LogDoc V2 practice, this Markdown file preserves current operational continuity and carries the v0.4.x runtime-integration thread forward; the JSON breakpoint remains the lean machine-readable handoff state.

**Compression note:** The detailed v0.3.0 through v0.3.36PRIME process ledger remains archived in PRIME/sanitized PRIME. v0.4.0 compressed that pre-release sprint; v0.4.1 through v0.4.12 preserved the runtime recon, command contracts, dev OpenMLS send/inbox commands, smoke proof, runner profile, boundary checks, wrapper recon, and bootstrap command contract. v0.4.13 now preserves the first implementation of that wrapper contract: identity create/status only.

---

## 1. Project Goal

**Active goal:** Continue moving CarbonStackComms from stub-era runtime `send`/`inbox` behavior toward explicit dev-only OpenMLS/Cypher-backed runtime commands and repeatable validation, while preserving claim discipline. v0.4.13 implements the smallest safe bootstrap wrapper subset and proves the wrapper pattern without migrating the smoke script or overclaiming local-backbone readiness.

CarbonStack’s broader project intent remains:

- **CarbonStack** is the secure-communications appliance-stack project and shared doctrine/release authority.
- **CarbonStackComms** is the text-first messaging client and current OpenMLS sidecar/relay integration surface.
- **CarbonStackCypher** is the self-hostable relay/storage stack.
- **CarbonStackOS** is future constrained Android-derived appliance OS work and remains out of current runnable validation packages.
- The `carbonstack` Gitea repo remains the source of truth for release state, release assets, docs, and public claims.
- GitHub mirrors exist as secondary push mirrors only; Gitea remains authoritative.

**North star:**

> Its goal is not to be a better smartphone. Its goal is to stop being a smartphone.

**Core doctrine:**

> Every feature is guilty until it proves it does not add unacceptable parser, network, sensor, identity, or filesystem authority.

**Current reality after v0.4.13:**

    openmls-send-dev is a dev-only OpenMLS send-side command.
    openmls-inbox-dev is a dev-only OpenMLS receive/open/optional-ack command.
    scripts/dev-openmls-runtime-smoke.sh proves the dev CLI application-message path:
      openmls-send-dev -> Cypher -> openmls-inbox-dev --ack
    tools/carbonstack-validate has a manual dev-runtime-openmls profile that wraps that smoke script.
    dev-runtime-openmls remains manual-only, live-git-umbrella-only, and outside full.
    openmls-identity-create-dev now exists.
    openmls-identity-status-dev now exists.
    identity bootstrap wrappers live in carbonstack-comms/internal/app/openmls_bootstrap.go.
    identity bootstrap wrapper tests live in carbonstack-comms/internal/app/openmls_bootstrap_dev_test.go.
    identity wrappers keep sidecar labels explicit through --sidecar-device-label.
    identity wrappers normalize sidecar JSON to stable key/value Comms output.
    identity wrappers do not mutate Comms state/trust files.
    carbonstack-comms README documents the identity bootstrap wrappers and boundaries.
    Existing send/inbox remain stub-era.
    Bundle/conversation/add-member/join wrappers are not implemented yet.
    The smoke script still uses direct sidecar bootstrap.
    local-backbone remains reserved.

Hard nonclaims remain: v0.4.13 does **not** prove production readiness, production E2EE, hostile-server safety, metadata privacy, Android readiness, CarbonStackOS readiness, secure vault/storage, public ingress safety, real homelab validation, audit, certification, PQ/hybrid security, mature messenger UX, release-package runtime validation, Relay Space join/onboarding, full bootstrap wrapper coverage, smoke-script wrapper migration, or local-backbone readiness.

---

## 2. Current Public Release State

Public release:

    https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.4.0

Release title:

    CarbonStack v0.4.0 Broad Local Deployability Pre-Release

Release tag:

    v0.4.0

Release commit:

    24ac9fc494 / 24ac9fc docs: record v0.4.0 package rehearsal

Gitea release status:

    Pre-Release

Important release-page warning remains:

    The default Gitea Source Code ZIP/TAR.GZ downloads are auto-generated carbonstack-only archives.
    The intended runnable validation artifact is the attached multi-repo release package plus companion assets.

Post-release public surface:

    carbonstack README is evergreen.
    It points interested parties to the Releases page for runnable / known-good artifacts.
    Gitea is source of truth.
    GitHub mirrors are secondary push mirrors, not release authority.

---

## 3. Current Repo Heads

    carbonstack        0583683 (HEAD -> main, origin/main, origin/HEAD) docs: define dev OpenMLS bootstrap command contract
    carbonstack-comms  20c4cc9 (HEAD -> main, origin/main, origin/HEAD) feat: add dev OpenMLS identity bootstrap commands
    carbonstack-cypher 9ab994c (HEAD -> main, origin/main, origin/HEAD) docs: point to local operator runbook
    carbonstack-os     1bbbe52 (HEAD -> main, origin/main, origin/HEAD) docs: clarify CarbonStackOS target direction

Current public release tag:

    v0.4.0 -> 24ac9fc494 / 24ac9fc docs: record v0.4.0 package rehearsal

Recent runtime/thread commits:

    v0.4.8 extraction:
      0a48ae1 refactor: split OpenMLS runtime command helpers

    v0.4.9 runner profile:
      8eeadb2 feat: add dev OpenMLS runtime validation profile

    v0.4.10 profile boundary:
      7384f43 docs: record dev OpenMLS runtime profile boundary

    v0.4.11 wrapper recon:
      8ee3ebb docs: plan dev OpenMLS bootstrap wrappers

    v0.4.12 command contract:
      0583683 docs: define dev OpenMLS bootstrap command contract

    Current v0.4.13 identity wrapper implementation:
      20c4cc9 feat: add dev OpenMLS identity bootstrap commands

Continuity note:

    `20c4cc9` is a carbonstack-comms implementation commit after the v0.4.0 public release tag.
    It does not retag v0.4.0.
    The release remains attached to `24ac9fc494`.

---

## 4. Current Validated / Known State

Validated / known at the v0.4.13 checkpoint:

    [RELEASE] v0.4.0 Gitea release exists, is marked Pre-Release, and is attached to 24ac9fc494.
    [RUNNER] full means release-snapshot followed by local-cypher.
    [RUNNER] dev-runtime-openmls validates the dev/pre-alpha OpenMLS application-message runtime CLI smoke proof.
    [RUNNER] dev-runtime-openmls remains manual-only and live-umbrella-only.
    [RUNNER] dev-runtime-openmls is not included in full.
    [DOCS] docs/162-openmls-bootstrap-command-contract-v0.md records the v0.4.12 dev-only OpenMLS bootstrap command contract.
    [COMMS] openmls-send-dev exists.
    [COMMS] openmls-inbox-dev exists.
    [COMMS] openmls-identity-create-dev now exists.
    [COMMS] openmls-identity-status-dev now exists.
    [COMMS] internal/app/openmls_bootstrap.go now exists.
    [COMMS] internal/app/openmls_bootstrap_dev_test.go now exists.
    [COMMS] internal/app/commands.go registers the two new identity bootstrap commands and lists them in usage.
    [COMMS] README.md now documents the dev-only OpenMLS bootstrap commands and their boundary.
    [COMMS] scripts/dev-openmls-runtime-smoke.sh still uses direct sidecar bootstrap.
    [COMMS] future bundle/conversation/add-member/join wrappers remain unimplemented.
    [COMMS] existing send/inbox/ack remain in place and stub-era.
    [MIRRORS] all four repos have GitHub push-mirror functionality, with Gitea remaining source of truth.

v0.4.13 validation result from the final log:

    go test ./internal/app -count=1 passed in carbonstack-comms.
    go test ./... -count=1 -timeout 600s passed in carbonstack-comms.
    Manual openmls-identity-create-dev run passed.
    Manual openmls-identity-status-dev run passed and printed identity_exists: true.
    go test ./... -count=1 passed in carbonstack/tools/carbonstack-validate.
    go run . --profile dev-runtime-openmls --clean-generated passed.
    local-cypher passed.
    doctor passed.
    core --clean-generated passed.
    carbonstack-comms was committed and pushed at:
      20c4cc9 feat: add dev OpenMLS identity bootstrap commands
    final repo status showed all four repos clean.

Expected live umbrella validation commands now include:

    cd ~/repos/carbonstack_umbrella/carbonstack-comms
    go test ./internal/app -count=1
    go test ./... -count=1 -timeout 600s

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go run . --profile dev-runtime-openmls --clean-generated
    go run . --profile local-cypher
    go run . --profile doctor
    go run . --profile core --clean-generated

New manual identity wrapper checks:

    cd ~/repos/carbonstack_umbrella/carbonstack-comms
    rm -rf internal/protocol/mls/openmls-sidecar/.carbonstack-openmls-sidecar-state
    go run ./cmd/comms openmls-identity-create-dev --sidecar-device-label carbonstack-v0413-identity-test
    go run ./cmd/comms openmls-identity-status-dev --sidecar-device-label carbonstack-v0413-identity-test

Expected v0.4.0 public package validation commands remain unchanged:

    cd <package-root>/carbonstack/tools/carbonstack-validate
    go run . --profile verify-checksums --root <package-root>
    go run . --profile full --root <package-root> --clean-generated

---

## 5. v0.4.13 Work Completed

### 5.1 First bootstrap wrapper subset was implemented

New commands implemented in carbonstack-comms:

    openmls-identity-create-dev
    openmls-identity-status-dev

These commands are intentionally dev-only and sidecar-bootstrap-scoped.

They do not:

    replace send/inbox.
    mutate Comms state.json.
    mutate trust.json.
    mutate trust-events.jsonl.
    call Cypher.
    implement Relay Space join UX.
    imply local-backbone.
    imply production identity UX.

### 5.2 openmls_bootstrap.go was added

New file:

    carbonstack-comms/internal/app/openmls_bootstrap.go

It contains:

    openMLSSidecarBootstrapEnvelope.
    runOpenMLSBootstrapSidecarForCommand injection seam.
    cmdOpenMLSIdentityCreateDev.
    cmdOpenMLSIdentityStatusDev.
    runOpenMLSBootstrapSidecar.
    bootstrapStringField.
    bootstrapBoolField.

Implementation behavior:

    calls the OpenMLS sidecar with cargo run --quiet -- <sidecar-command>.
    parses sidecar JSON envelopes.
    converts sidecar ok=false into Go errors.
    avoids printing success after sidecar failure.
    prints stable key/value output.
    keeps the sidecar device label explicit.
    prints dev/pre-alpha identity UX warning.

### 5.3 openmls_bootstrap_dev_test.go was added

New file:

    carbonstack-comms/internal/app/openmls_bootstrap_dev_test.go

Tests cover:

    openmls-identity-create-dev requires --sidecar-device-label.
    openmls-identity-status-dev requires --sidecar-device-label.
    identity-create invokes the expected sidecar command and args.
    identity-create prints stable output.
    identity-status invokes the expected sidecar command and args.
    identity-status prints stable output including identity_exists when present.
    sidecar failure prevents success output.

Test design follows the v0.4.12 contract:

    app tests use injected seams.
    app tests do not run the Rust sidecar directly.
    real sidecar execution remains in manual wrapper smoke, protocol tests, smoke script, and runner validation.

### 5.4 commands.go was updated

Updated file:

    carbonstack-comms/internal/app/commands.go

New dispatch cases:

    openmls-identity-create-dev
    openmls-identity-status-dev

Usage output now lists:

    openmls-identity-create-dev
    openmls-identity-status-dev

The existing command registry remains in commands.go. Message runtime helpers remain in openmls_runtime.go. Bootstrap wrapper helpers now live in openmls_bootstrap.go.

### 5.5 carbonstack-comms README was updated

Updated file:

    carbonstack-comms/README.md

New section:

    Dev-only OpenMLS bootstrap commands

It documents:

    openmls-identity-create-dev
    openmls-identity-status-dev

It includes examples:

    go run ./cmd/comms openmls-identity-create-dev --sidecar-device-label carbonstack-dev-alice
    go run ./cmd/comms openmls-identity-status-dev --sidecar-device-label carbonstack-dev-alice

Boundary preserved:

    sidecar labels are explicit for now.
    Comms state/trust files are not mutated by these wrappers.
    existing send/inbox remain stub-era.
    dev-runtime-openmls remains the current manual smoke-profile proof.
    these commands are not production identity UX, not Relay Space join UX, not local-backbone, and not secure vault/state management.

### 5.6 Manual wrapper smoke passed

Manual check:

    rm -rf internal/protocol/mls/openmls-sidecar/.carbonstack-openmls-sidecar-state

    go run ./cmd/comms openmls-identity-create-dev --sidecar-device-label carbonstack-v0413-identity-test
    go run ./cmd/comms openmls-identity-status-dev --sidecar-device-label carbonstack-v0413-identity-test

Observed output included:

    openmls dev bootstrap
    command: openmls-identity-create-dev
    status: created
    sidecar_command: identity-create
    sidecar_device_label: carbonstack-v0413-identity-test
    warning: dev/pre-alpha OpenMLS bootstrap path; not production identity UX

and:

    openmls dev bootstrap
    command: openmls-identity-status-dev
    status: loaded
    sidecar_command: identity-status
    sidecar_device_label: carbonstack-v0413-identity-test
    identity_exists: true
    warning: dev/pre-alpha OpenMLS bootstrap path; not production identity UX

### 5.7 Validation passed

Commands observed passing:

    cd ~/repos/carbonstack_umbrella/carbonstack-comms
    go test ./internal/app -count=1
    go test ./... -count=1 -timeout 600s

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go test ./... -count=1
    go run . --profile dev-runtime-openmls --clean-generated
    go run . --profile local-cypher
    go run . --profile doctor
    go run . --profile core --clean-generated

Final commit/push:

    carbonstack-comms:
    20c4cc9 feat: add dev OpenMLS identity bootstrap commands

Final repo snapshot:

    carbonstack        0583683 docs: define dev OpenMLS bootstrap command contract
    carbonstack-comms  20c4cc9 feat: add dev OpenMLS identity bootstrap commands
    carbonstack-cypher 9ab994c docs: point to local operator runbook
    carbonstack-os     1bbbe52 docs: clarify CarbonStackOS target direction

---

## 6. Critical Blunders / Lessons to Carry Forward

### 6.1 Pre-dev-runtime artifact scan saw known generated roots from manual wrapper smoke

During validation, `dev-runtime-openmls` reported pre-test artifact scan hits for known OpenMLS sidecar generated roots:

    .carbonstack-openmls-sidecar-state
    target

This was expected because the manual identity wrapper smoke had just created sidecar state / Rust build artifacts. The runner classified them as known OpenMLS generated roots, and `--clean-generated` removed them after the profile.

Lesson:

    Manual wrapper smoke can leave expected sidecar state and target roots.
    This is acceptable only when they stay in known generated paths.
    Always follow with --clean-generated when preparing a clean breakpoint.

### 6.2 Commit metadata warning remains a polish issue

The commit log again showed auto-generated local commit metadata:

    Committer: bitcrusher32 <▓▓>

This is not a functional validation failure, but remains a professionalism/publishing hygiene issue to address eventually with `git config --global user.name` and `git config --global user.email` if desired.

### 6.3 Wrapper contract held for first subset, but not all wrappers are proven

v0.4.13 proves the contract shape for identity create/status only.

Do not assume:

    bundle export wrapper exists.
    conversation create/load-check wrappers exist.
    add-member/join wrappers exist.
    smoke script uses wrappers.
    full bootstrap wrapper coverage exists.
    local-backbone is justified.

Correct interpretation:

    first identity bootstrap wrapper subset exists and passed tests.
    remaining wrappers still need staged implementation.

---

## 7. Current Blockers / Not Validated

Current blockers and nonvalidated areas after v0.4.13:

    Existing send/inbox remain stub-era.
    Runtime Comms has dev-only OpenMLS commands, a smoke proof, and a manual runner profile, but no mature user-facing UX.
    Sidecar KeyPackage/Welcome/bootstrap setup is still direct dev setup inside the smoke script.
    Bundle export wrapper is not implemented yet.
    Conversation create/load-check wrappers are not implemented yet.
    Conversation add-member/join wrappers are not implemented yet.
    Smoke script has not been migrated to wrappers.
    No Relay Space join model exists yet.
    local-backbone runner profile does not exist and remains reserved.
    dev-runtime-openmls remains manual-only and live-git-umbrella-only.
    dev-runtime-openmls is not included in full.
    dev-runtime-openmls is not part of public release package validation.
    Fresh release-package behavior for dev-runtime-openmls has not been tested as a supported validation surface.
    Post-release public download verification is still useful if not already performed.
    Sanitized PRIME placement in carbonstack/logdoc-loop-system still needs a policy decision.
    Production-safe provider storage and secure vault design remain v0.5.x future work.
    PQ/hybrid ciphersuite migration readiness remains v0.5.x future work.
    Hostile-server rollback/replay/metadata-abuse validation remains future work.
    Android/CarbonStackOS implementation remains later work.
    CarbonStack Relay Space mechanics remain a future concept, not an implementation.

Hard nonclaims remain:

    Do not claim production readiness.
    Do not claim production E2EE.
    Do not claim hostile-server safety.
    Do not claim metadata privacy.
    Do not claim Debian production deployability.
    Do not claim systemd readiness.
    Do not claim cloudflared readiness.
    Do not claim Android readiness.
    Do not claim CarbonStackOS readiness.
    Do not claim secure vault/storage.
    Do not claim real homelab validation.
    Do not claim external audit or certification.
    Do not claim CarbonStack Relay Space mechanics are implemented.
    Do not claim local-backbone exists or is justified yet.
    Do not claim local-cypher validates runtime Comms UX.
    Do not claim full validates local-backbone.
    Do not claim full is a deployment command.
    Do not claim dev-runtime-openmls is release-package validation yet.
    Do not claim send/inbox are OpenMLS-backed yet.
    Do not claim PQ/hybrid security or quantum-safe messaging.
    Do not claim all bootstrap wrapper commands exist yet.
    Do not let GitHub mirrors supersede Gitea source-of-truth status.

Allowed claim:

    CarbonStackComms now has dev-only openmls-send-dev and openmls-inbox-dev commands.
    CarbonStackComms now has a dev runtime smoke script proving openmls-send-dev -> Cypher -> openmls-inbox-dev --ack.
    CarbonStackComms now has OpenMLS runtime command helper glue split into internal/app/openmls_runtime.go.
    CarbonStackComms now has dev-only openmls-identity-create-dev and openmls-identity-status-dev bootstrap wrappers.
    The identity bootstrap wrappers parse sidecar JSON, print stable key/value output, and keep sidecar labels explicit.
    The identity bootstrap wrappers do not mutate Comms state/trust files.
    carbonstack/tools/carbonstack-validate has a manual dev-runtime-openmls profile.
    dev-runtime-openmls wraps the comms smoke script and validates the dev/pre-alpha OpenMLS application-message CLI path.
    dev-runtime-openmls is live-git-umbrella-only and not included in full.
    The main carbonstack repo records the smoke proof, pre-local-backbone assessment, manual dev-runtime-openmls runner profile, profile boundary check, bootstrap wrapper recon/planning doc, and bootstrap wrapper command contract.
    Existing send/inbox remain stub-era.
    v0.4.0 remains the current public broad local deployability pre-release.
    Gitea remains source of truth; GitHub mirrors are secondary push mirrors.

---

## 8. Next Safest Actions

Recommended next rung:

    v0.4.14 implement bundle export + conversation create/load-check dev wrappers

Focus:

    carbonstack-comms only.
    inspect exact sidecar JSON fields for public-bundle-export, conversation-create, and conversation-load-check if needed.
    extend internal/app/openmls_bootstrap.go.
    extend internal/app/openmls_bootstrap_dev_test.go.
    register:
      openmls-bundle-export-dev
      openmls-conversation-create-dev
      openmls-conversation-load-check-dev
    keep sidecar labels explicit.
    keep output stable key/value text.
    preserve path hints and print absolute artifact paths when applicable.
    do not mutate Comms state/trust.
    do not touch send/inbox.
    do not migrate the smoke script yet.
    update carbonstack-comms README if implementation lands.
    validate app tests, Comms tests, manual wrapper checks, dev-runtime-openmls, local-cypher, doctor, and core.

After v0.4.14:

    update the long-term roadmap PDF/roadmap artifact to reflect v0.4.1-v0.4.14 reality.
    carry forward a future known-good command registry idea for late v0.4.x or pre-v0.5.x.

Suggested future command registry idea:

    a preliminary YAML or JSON inventory of known-good project backbone commands.
    include command name, repo/component, maturity, purpose, why it exists, validation status, and nonclaims.
    place it outside front-door README unless it becomes user-critical.
    likely timing: end of v0.4.x or just before PQ work in v0.5.x.

Avoid next:

    implementing add-member/join before v0.4.14 proves bundle/conversation wrappers.
    replacing the smoke script before wrapper tests exist.
    adding dev-runtime-openmls to full.
    calling any wrapper local-backbone.
    using Relay Space join naming for dev bootstrap wrappers.
    moving dev-runtime-openmls into public release validation without fresh-root behavior checks.
    public ingress.
    cloudflared.
    systemd.
    real homelab deployment.
    Android app.
    CarbonStackOS implementation.
    CarbonStack Relay Space implementation.
    PQ/hybrid ciphersuite implementation.
    production/security claims.

---

## 9. Critical Paths / Functions

### Source-of-truth public surfaces

    [GITEA SOURCE OF TRUTH] https://git.bitcrusher32.win/bitcrusher32/carbonstack
    [RELEASES] https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases
    [CURRENT PUBLIC RELEASE] https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.4.0

GitHub mirrors:

    secondary push mirrors for discoverability/redundancy only.
    Do not treat mirrors as release authority unless project policy changes.

### WSL Debian active working umbrella

    [WSL-DEBIAN] umbrella: ~/repos/carbonstack_umbrella
    [WSL-DEBIAN] carbonstack: ~/repos/carbonstack_umbrella/carbonstack
    [WSL-DEBIAN] carbonstack-comms: ~/repos/carbonstack_umbrella/carbonstack-comms
    [WSL-DEBIAN] carbonstack-cypher: ~/repos/carbonstack_umbrella/carbonstack-cypher
    [WSL-DEBIAN] carbonstack-os: ~/repos/carbonstack_umbrella/carbonstack-os
    [GO-RUNNER] runner dir: ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate

### Runtime integration doc surfaces

    [DOC] v0.4.1 runtime recon:
    carbonstack/docs/155-runtime-comms-openmls-cypher-recon-v0.md

    [DOC] v0.4.2 command contract:
    carbonstack/docs/156-runtime-openmls-command-contract-v0.md

    [DOC] v0.4.6 smoke proof record:
    carbonstack/docs/157-dev-runtime-openmls-smoke-proof-v0.md

    [DOC] v0.4.7 pre-local-backbone assessment:
    carbonstack/docs/158-pre-local-backbone-assessment-v0.md

    [DOC] v0.4.9 runner profile:
    carbonstack/docs/159-dev-runtime-openmls-runner-profile-v0.md

    [DOC] v0.4.10 profile boundary:
    carbonstack/docs/160-dev-runtime-openmls-profile-boundary-v0.md

    [DOC] v0.4.11 bootstrap wrapper recon:
    carbonstack/docs/161-openmls-bootstrap-wrapper-recon-v0.md

    [DOC] v0.4.12 bootstrap command contract:
    carbonstack/docs/162-openmls-bootstrap-command-contract-v0.md

    [DOC] docs index:
    carbonstack/docs/README.md

    [DOC] roadmap:
    carbonstack/roadmap/ROADMAP.md

    [DOC] top README:
    carbonstack/README.md

### Runner surfaces

    [RUNNER] root:
    carbonstack/tools/carbonstack-validate

    [RUNNER] dispatch:
    carbonstack/tools/carbonstack-validate/main.go

    [RUNNER] dev-runtime-openmls implementation:
    carbonstack/tools/carbonstack-validate/dev_runtime_openmls.go

    [RUNNER] runner README:
    carbonstack/tools/carbonstack-validate/README.md

    [RUNNER] local-cypher:
    carbonstack/tools/carbonstack-validate/local_cypher.go

    [RUNNER] release-snapshot:
    carbonstack/tools/carbonstack-validate/release_snapshot.go

    [RUNNER] checksums:
    carbonstack/tools/carbonstack-validate/checksums.go

### Comms runtime surfaces

    [COMMS] CLI entry:
    carbonstack-comms/cmd/comms/main.go

    [COMMS] command registry and non-OpenMLS commands:
    carbonstack-comms/internal/app/commands.go

    [COMMS] OpenMLS message runtime command helpers:
    carbonstack-comms/internal/app/openmls_runtime.go

    [COMMS] OpenMLS bootstrap wrapper helpers:
    carbonstack-comms/internal/app/openmls_bootstrap.go

    [COMMS] OpenMLS bootstrap wrapper tests:
    carbonstack-comms/internal/app/openmls_bootstrap_dev_test.go

    [COMMS] openmls-send-dev tests:
    carbonstack-comms/internal/app/openmls_send_dev_test.go

    [COMMS] openmls-inbox-dev tests:
    carbonstack-comms/internal/app/openmls_inbox_dev_test.go

    [COMMS] dev runtime smoke proof:
    carbonstack-comms/scripts/dev-openmls-runtime-smoke.sh

    [COMMS] OpenMLS sidecar dir:
    carbonstack-comms/internal/protocol/mls/openmls-sidecar

    [COMMS] sidecar README:
    carbonstack-comms/internal/protocol/mls/openmls-sidecar/README.md

### Current stub-era Comms functions / behaviors

    send:
      trust policy evaluation
      mock/stub provider encryption
      Cypher envelope submit

    inbox:
      Cypher inbox retrieval
      mock/stub provider plaintext display as stub_plaintext

    ack:
      envelope ack by envelope ID and recipient device ID

### Implemented v0.4.x dev OpenMLS commands

    openmls-send-dev:
      internal/app/openmls_runtime.go
      state.RequireReadyDevice
      trust.EvaluateSend
      runOpenMLSMessageProtect
      sidecar message-protect
      internal/relay.SubmitOpenMLSArtifactEnvelope
      ArtifactKindApplicationMessage
      dev/pre-alpha warning

    openmls-inbox-dev:
      internal/app/openmls_runtime.go
      state.RequireReadyDevice
      Cypher inbox
      skip unsupported envelopes
      internal/relay.WriteOpenMLSArtifactFromEnvelope
      runOpenMLSMessageOpen
      sidecar message-open
      print plaintext after sidecar success
      optional ack only after sidecar success

    openmls-identity-create-dev:
      internal/app/openmls_bootstrap.go
      sidecar identity-create
      required --sidecar-device-label
      optional --sidecar-dir
      stable key/value output
      no Comms state/trust mutation

    openmls-identity-status-dev:
      internal/app/openmls_bootstrap.go
      sidecar identity-status
      required --sidecar-device-label
      optional --sidecar-dir
      stable key/value output
      prints identity_exists when present
      no Comms state/trust mutation

### Contracted but not yet implemented v0.4.14+ dev OpenMLS bootstrap wrappers

    openmls-bundle-export-dev:
      sidecar public-bundle-export
      required --sidecar-device-label
      optional --write-artifact

    openmls-conversation-create-dev:
      sidecar conversation-create
      required --sidecar-device-label and --conversation

    openmls-conversation-load-check-dev:
      sidecar conversation-load-check
      required --sidecar-device-label and --conversation

    openmls-conversation-add-member-dev:
      sidecar conversation-add-member
      required --sidecar-device-label, --conversation, and --member-keypackage

    openmls-conversation-join-dev:
      sidecar conversation-join
      required --sidecar-device-label, --conversation, and --welcome

### v0.4.5 through v0.4.13 runtime smoke proof

    scripts/dev-openmls-runtime-smoke.sh:
      temp workspace
      temp Cypher binary
      temp Cypher DB
      dynamic loopback Cypher port
      temp Alice/Bob Comms state
      sidecar state reset
      direct sidecar identity/conversation bootstrap
      openmls-send-dev application-message send
      openmls-inbox-dev application-message open + ack
      inbox-empty-after-ack verification
      robust cleanup with INT -> TERM -> KILL escalation

### v0.4.9/v0.4.10 runner profile

    go run . --profile dev-runtime-openmls --clean-generated

    Runner.DevRuntimeOpenMLS:
      CheckRequiredPaths
      CheckLiveGitUmbrella
      toolchain reporting
      ArtifactScan("pre-dev-runtime-openmls")
      RunStep for carbonstack-comms/scripts/dev-openmls-runtime-smoke.sh
      ArtifactScan("post-dev-runtime-openmls")
      explicit nonclaim summary

### Existing v0.4.0 package validation

    cd <package-root>/carbonstack/tools/carbonstack-validate
    go run . --profile verify-checksums --root <package-root>
    go run . --profile full --root <package-root> --clean-generated

---

## 10. v0.4.x Forward Plan

### v0.4.14 — bundle export + conversation create/load-check wrapper implementation

Expected:

    carbonstack-comms-only implementation.
    extend internal/app/openmls_bootstrap.go.
    extend internal/app/openmls_bootstrap_dev_test.go.
    register openmls-bundle-export-dev.
    register openmls-conversation-create-dev.
    register openmls-conversation-load-check-dev.
    keep command registry in commands.go.
    keep output as stable key/value text.
    preserve path hints and absolute paths where applicable.
    do not mutate Comms state/trust.
    update carbonstack-comms README if implementation lands.
    validate app tests, full Comms tests, manual wrapper checks, dev-runtime-openmls, local-cypher, doctor, and core.

After v0.4.14:

    update latest long-term roadmap artifact to reflect current v0.4.x reality.

### v0.4.15+

Possible staged implementation:

    v0.4.15:
      openmls-conversation-add-member-dev
      openmls-conversation-join-dev

    v0.4.16:
      smoke-script migration recon / wrapper-based smoke variant decision

### v0.4.x later polish

Likely needs:

    exact sidecar setup runbook.
    clearer dev state reset semantics.
    better command output shape.
    decision on whether stub-era send/inbox are aliased, replaced, or deprecated.
    preliminary known-good command registry YAML/JSON outside the front-door README unless user-critical.
    decision on when local-backbone naming stops being premature.

---

## 11. v0.5.x and Later Carry-Forward

### v0.5.x

Local storage, trust state, provider state, vault design, and PQ/hybrid ciphersuite migration readiness.

Do not start PQ/hybrid implementation before the runtime OpenMLS/Cypher path exists and is documented. PQ/hybrid migration touches provider state, ciphersuite metadata, trust display, artifact size, reinit/rekey/recovery policy, and claim boundaries.

### v0.6.x+

Hostile-server and abuse-resistance harnesses.

### v0.7.x+

Deployability and operations hardening.

### v0.8.x+

Documented self-pentest / adversarial validation.

### v0.9.x+

Claim-boundary review.

### v0.10.x+

Android backend/app work.

### v1.x.x

Public app/server major epoch and much later CarbonStackOS/appliance research, if justified.

---

## 12. Lean Breakpoint Summary

    CarbonStack v0.4.13 is the first dev-only OpenMLS bootstrap wrapper implementation checkpoint. v0.4.0 remains the current public broad local deployability pre-release. carbonstack remains at 0583683 docs: define dev OpenMLS bootstrap command contract. carbonstack-comms is now at 20c4cc9 feat: add dev OpenMLS identity bootstrap commands. v0.4.13 adds internal/app/openmls_bootstrap.go, internal/app/openmls_bootstrap_dev_test.go, registers openmls-identity-create-dev and openmls-identity-status-dev, and documents the new dev-only identity bootstrap commands in carbonstack-comms README.md. The commands keep sidecar labels explicit, normalize sidecar JSON into stable Comms key/value output, and do not mutate Comms state/trust files. Validation passed for app tests, full Comms tests, manual identity wrapper smoke, dev-runtime-openmls --clean-generated, local-cypher, doctor, and core --clean-generated. Final repo heads were clean across all four repos. Next safe rung: v0.4.14 bundle export + conversation create/load-check wrappers, then update the long-term roadmap artifact after v0.4.14.

---

## 13. Preserved Immediate Previous Handoff: v0.4.12

The following is the previous v0.4.12 handoff. Where it conflicts with the v0.4.13 overlay above, v0.4.13 wins for current state.


# CarbonStack LogDoc v0.4.12

**Last updated:** 2026-06-05 local session  
**Owner/operator:** bitcrusher32  
**Assistant/model:** GPT-5.5 Thinking  
**Current phase:** Phase 2F / **v0.4.x dev-only OpenMLS bootstrap command-contract checkpoint complete**. The v0.4.0 public broad local deployability pre-release remains the current public release. v0.4.5 proved the dev runtime OpenMLS CLI path; v0.4.6 recorded that proof in the main `carbonstack` repo; v0.4.7 kept `local-backbone` reserved; v0.4.8 split OpenMLS runtime command glue into `carbonstack-comms/internal/app/openmls_runtime.go`; v0.4.9 added the manual live-umbrella-only `dev-runtime-openmls` validation profile; v0.4.10 stabilized that profile boundary; v0.4.11 planned dev-only OpenMLS bootstrap wrappers; v0.4.12 now defines the strict command contract before any wrapper implementation. `carbonstack` is now at `0583683 (HEAD -> main, origin/main, origin/HEAD) docs: define dev OpenMLS bootstrap command contract`; `carbonstack-comms` remains at `0a48ae1 (HEAD -> main, origin/main, origin/HEAD) refactor: split OpenMLS runtime command helpers`; `carbonstack-cypher` remains at `9ab994c (HEAD -> main, origin/main, origin/HEAD) docs: point to local operator runbook`; `carbonstack-os` remains at `1bbbe52 (HEAD -> main, origin/main, origin/HEAD) docs: clarify CarbonStackOS target direction`.

**Version schema:** `v[scope].[timeline]`. This file is `v0.4.12`, the dev-only OpenMLS bootstrap command-contract checkpoint after the v0.4.11 wrapper recon/planning checkpoint. Per LogDoc V2 practice, this Markdown file preserves current operational continuity and carries the v0.4.x runtime-integration thread forward; the JSON breakpoint remains the lean machine-readable handoff state.

**Compression note:** The detailed v0.3.0 through v0.3.36PRIME process ledger remains archived in PRIME/sanitized PRIME. v0.4.0 compressed that pre-release sprint; v0.4.1 preserved runtime recon; v0.4.2 preserved the original runtime command-contract decision; v0.4.3 preserved the first dev-only OpenMLS send-side implementation; v0.4.4 preserved the first dev-only OpenMLS receive/open/optional-ack implementation; v0.4.5 preserved the first committed dev runtime smoke proof; v0.4.6 preserved project-level smoke-proof alignment; v0.4.7 preserved the pre-local-backbone decision boundary; v0.4.8 preserved OpenMLS helper extraction; v0.4.9 preserved the manual `dev-runtime-openmls` runner profile; v0.4.10 preserved the profile boundary evidence; v0.4.11 preserved wrapper recon/planning; v0.4.12 now preserves the formal bootstrap wrapper command contract.

---

## 1. Project Goal

**Active goal:** Continue moving CarbonStackComms from stub-era runtime `send`/`inbox` behavior toward explicit dev-only OpenMLS/Cypher-backed runtime commands and repeatable validation, while preserving claim discipline. v0.4.12 did not implement wrapper commands; it defined the contract that future wrapper implementations must follow.

CarbonStack’s broader project intent remains:

- **CarbonStack** is the secure-communications appliance-stack project and shared doctrine/release authority.
- **CarbonStackComms** is the text-first messaging client and current OpenMLS sidecar/relay integration surface.
- **CarbonStackCypher** is the self-hostable relay/storage stack.
- **CarbonStackOS** is future constrained Android-derived appliance OS work and remains out of current runnable validation packages.
- The `carbonstack` Gitea repo remains the source of truth for release state, release assets, docs, and public claims.
- GitHub mirrors exist as secondary push mirrors only; Gitea remains authoritative.

**North star:**

> Its goal is not to be a better smartphone. Its goal is to stop being a smartphone.

**Core doctrine:**

> Every feature is guilty until it proves it does not add unacceptable parser, network, sensor, identity, or filesystem authority.

**Current reality after v0.4.12:** the runtime proof remains valid, `dev-runtime-openmls` remains manual/live-umbrella-only, sidecar bootstrap wrappers are still not implemented, and the exact wrapper contract now exists in the main project docs.

Correct interpretation:

    openmls-send-dev is a dev-only OpenMLS send-side command.
    openmls-inbox-dev is a dev-only OpenMLS receive/open/optional-ack command.
    scripts/dev-openmls-runtime-smoke.sh proves the dev CLI application-message path:
      openmls-send-dev -> Cypher -> openmls-inbox-dev --ack
    tools/carbonstack-validate has a manual dev-runtime-openmls profile that wraps that smoke script.
    dev-runtime-openmls is repeatable from the live umbrella checkout.
    dev-runtime-openmls is live-git-umbrella-only for now.
    dev-runtime-openmls intentionally refuses non-git package-like roots.
    dev-runtime-openmls is not included in full.
    dev-runtime-openmls is not release-package validation yet.
    docs/162-openmls-bootstrap-command-contract-v0.md defines the future bootstrap wrapper contract.
    internal/app/openmls_runtime.go remains the OpenMLS message runtime command helper file.
    Future bootstrap wrapper code should likely live in internal/app/openmls_bootstrap.go.
    Existing send/inbox remain stub-era.
    local-backbone remains reserved.

Hard nonclaims remain: v0.4.12 does **not** prove production readiness, production E2EE, hostile-server safety, metadata privacy, Android readiness, CarbonStackOS readiness, secure vault/storage, public ingress safety, real homelab validation, audit, certification, PQ/hybrid security, mature messenger UX, release-package runtime validation, Relay Space join/onboarding, wrapper implementation, or local-backbone readiness.

---

## 2. Current Public Release State

Public release:

    https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.4.0

Release title:

    CarbonStack v0.4.0 Broad Local Deployability Pre-Release

Release tag:

    v0.4.0

Release commit:

    24ac9fc494 / 24ac9fc docs: record v0.4.0 package rehearsal

Gitea release status:

    Pre-Release

Release assets attached:

    carbonstack-v0.4.0-broad-local-deployability-pre-release.tgz
    carbonstack-v0.4.0-release-manifest.json
    carbonstack-v0.4.0-package-checksums.txt
    carbonstack-v0.4.0-asset-checksums.txt
    carbonstack-v0.4.0-validation-freeze.md
    v0.4.0-testing-runbook.md
    v0.4.0-release-notes.md
    LICENSE

Important release-page warning:

    The default Gitea Source Code ZIP/TAR.GZ downloads are auto-generated carbonstack-only archives.
    The intended runnable validation artifact is the attached multi-repo release package plus companion assets.

Post-release public surface:

    carbonstack README is evergreen.
    It points interested parties to the Releases page for runnable / known-good artifacts.
    It states Gitea is the source of truth.
    GitHub mirrors are secondary push mirrors, not release authority.

---

## 3. Current Repo Heads

    carbonstack        0583683 (HEAD -> main, origin/main, origin/HEAD) docs: define dev OpenMLS bootstrap command contract
    carbonstack-comms  0a48ae1 (HEAD -> main, origin/main, origin/HEAD) refactor: split OpenMLS runtime command helpers
    carbonstack-cypher 9ab994c (HEAD -> main, origin/main, origin/HEAD) docs: point to local operator runbook
    carbonstack-os     1bbbe52 (HEAD -> main, origin/main, origin/HEAD) docs: clarify CarbonStackOS target direction

Current public release tag:

    v0.4.0 -> 24ac9fc494 / 24ac9fc docs: record v0.4.0 package rehearsal

Post-release public-surface cleanup:

    6ba7aab docs: make front readme evergreen

v0.4.1 runtime recon commit:

    83d2f0f docs: record runtime OpenMLS Comms recon

v0.4.2 command contract commit:

    a799764 docs: define runtime OpenMLS command contract

v0.4.3 send implementation commit:

    39fb287 feat: add dev OpenMLS send command

v0.4.4 inbox implementation commit:

    65bc707 feat: add dev OpenMLS inbox command

v0.4.5 smoke proof commit:

    8e6e8b4 test: add dev OpenMLS runtime smoke proof

v0.4.6 main-doc alignment commit:

    3cd77a8 docs: record dev OpenMLS runtime smoke proof

v0.4.7 assessment commit:

    dc2d16c docs: assess pre-local-backbone runtime proof

v0.4.8 extraction commit:

    0a48ae1 refactor: split OpenMLS runtime command helpers

v0.4.9 runner-profile commit:

    8eeadb2 feat: add dev OpenMLS runtime validation profile

v0.4.10 boundary-doc commit:

    7384f43 docs: record dev OpenMLS runtime profile boundary

v0.4.11 wrapper-recon commit:

    8ee3ebb docs: plan dev OpenMLS bootstrap wrappers

Current v0.4.12 command-contract commit:

    0583683 docs: define dev OpenMLS bootstrap command contract

Continuity note:

    `0583683` is a `carbonstack` docs/contract commit after the `v0.4.0` public release tag.
    It does not retag v0.4.0.
    The release remains attached to `24ac9fc494`.

---

## 4. Current Validated / Known State

Validated / known at the v0.4.12 checkpoint:

    [RELEASE] v0.4.0 Gitea release exists, is marked Pre-Release, and is attached to 24ac9fc494.
    [PACKAGE] v0.4.0 package includes carbonstack, carbonstack-comms, carbonstack-cypher, and release metadata.
    [PACKAGE] carbonstack-os is deliberately excluded from the runnable package.
    [RUNNER] full means release-snapshot followed by local-cypher.
    [RUNNER] release-snapshot validates package layout, metadata, strict pre-test artifact scan, package checksums, and core.
    [RUNNER] core validates doctor, OpenMLS real-Cypher lifecycle, Comms package tests, Cypher package tests, and artifact scans.
    [RUNNER] local-cypher validates Cypher-only lifecycle, restart/persistence against the same temporary DB, positive opaque OpenMLS envelope/ack flow, and historical negative stub-text/OpenMLS protocol pairing.
    [RUNNER] dev-runtime-openmls validates the dev/pre-alpha OpenMLS application-message runtime CLI smoke proof.
    [RUNNER] dev-runtime-openmls is manual-only and live-umbrella-only for now.
    [RUNNER] dev-runtime-openmls is not included in full.
    [RUNNER] dev-runtime-openmls checks live git markers for carbonstack, carbonstack-comms, and carbonstack-cypher.
    [RUNNER] dev-runtime-openmls refuses non-git package-like roots as intended.
    [RUNNER] dev-runtime-openmls wraps carbonstack-comms/scripts/dev-openmls-runtime-smoke.sh.
    [DOCS] docs/159-dev-runtime-openmls-runner-profile-v0.md records the v0.4.9 manual runner profile.
    [DOCS] docs/160-dev-runtime-openmls-profile-boundary-v0.md records the v0.4.10 boundary check.
    [DOCS] docs/161-openmls-bootstrap-wrapper-recon-v0.md records the v0.4.11 bootstrap wrapper recon/planning boundary.
    [DOCS] docs/162-openmls-bootstrap-command-contract-v0.md records the v0.4.12 dev-only OpenMLS bootstrap command contract.
    [DOCS] docs/README.md indexes docs/162-openmls-bootstrap-command-contract-v0.md.
    [DOCS] roadmap/ROADMAP.md records v0.4.12 and recommends v0.4.13 identity create/status wrapper implementation next.
    [COMMS] openmls-send-dev exists in carbonstack-comms.
    [COMMS] openmls-inbox-dev exists in carbonstack-comms.
    [COMMS] scripts/dev-openmls-runtime-smoke.sh exists in carbonstack-comms.
    [COMMS] smoke script still uses direct sidecar bootstrap calls for identity, public bundle, conversation create, add-member, and join.
    [COMMS] internal/app/openmls_runtime.go contains OpenMLS message runtime command helper code.
    [COMMS] future bootstrap wrapper code should likely live in internal/app/openmls_bootstrap.go.
    [COMMS] internal/app/commands.go remains the command registry / non-OpenMLS command surface.
    [COMMS] existing send/inbox/ack remain in place and stub-era.
    [MIRRORS] all four repos have GitHub push-mirror functionality, with Gitea remaining source of truth.

v0.4.12 validation result from the final log:

    go test ./... -count=1 passed for carbonstack/tools/carbonstack-validate.
    go run . --profile dev-runtime-openmls --clean-generated passed.
    local-cypher passed.
    doctor passed.
    core --clean-generated passed.
    docs/162-openmls-bootstrap-command-contract-v0.md was added.
    docs/README.md was updated to index the v0.4.12 doc.
    roadmap/ROADMAP.md was updated to record v0.4.12 and recommend v0.4.13.
    carbonstack was committed and pushed at:
      0583683 docs: define dev OpenMLS bootstrap command contract
    final repo status showed all four repos clean.

Expected live umbrella validation commands now include:

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go run . --profile dev-runtime-openmls --clean-generated
    go run . --profile local-cypher
    go run . --profile doctor
    go run . --profile core --clean-generated

Expected v0.4.0 public package validation commands remain unchanged:

    cd <package-root>/carbonstack/tools/carbonstack-validate
    go run . --profile verify-checksums --root <package-root>
    go run . --profile full --root <package-root> --clean-generated

Current direct smoke command remains:

    cd ~/repos/carbonstack_umbrella/carbonstack-comms
    scripts/dev-openmls-runtime-smoke.sh

---

## 5. v0.4.12 Work Completed

### 5.1 Bootstrap command-contract recon was sufficient for docs, not implementation

The v0.4.12 recon inspected:

    docs/161-openmls-bootstrap-wrapper-recon-v0.md.
    roadmap v0.4.11/v0.4.12 area.
    OpenMLS sidecar README.
    OpenMLS sidecar src/main.rs command handlers.
    OpenMLS sidecar JSON output fields.
    OpenMLS sidecar src/paths.rs.
    CarbonStackComms command registry.
    carbonstack-comms/internal/app/openmls_runtime.go.
    openmls-send-dev and openmls-inbox-dev tests.
    smoke script bootstrap sequence.
    protocol lifecycle tests.
    state/trust non-mutation boundary.

Key findings:

    The contract is clear enough to write now.
    The sidecar already exposes the needed bootstrap primitives.
    Existing app tests provide an injection-seam pattern for future wrapper tests.
    Wrappers should normalize sidecar JSON into stable Comms key/value output.
    Wrappers should preserve sidecar path hints and print absolute artifact paths where useful.
    Wrappers must not mutate Comms state/trust files.
    Identity create/status is the safest first implementation subset.
    Implementation still should not happen until after this contract checkpoint.

### 5.2 docs/162-openmls-bootstrap-command-contract-v0.md was added

New file:

    carbonstack/docs/162-openmls-bootstrap-command-contract-v0.md

It defines:

    future dev-only wrapper command names.
    shared flag policy.
    command-specific required/optional flags.
    sidecar call mapping.
    success output shape.
    sidecar JSON parsing policy.
    path hint / absolute path policy.
    error policy.
    state/trust non-mutation policy.
    app unit-test seam policy.
    implementation file boundary.
    staged implementation order.
    smoke-script migration conditions.
    nonclaims.

### 5.3 Command names were formalized

Future commands should be named:

    openmls-identity-create-dev
    openmls-identity-status-dev
    openmls-bundle-export-dev
    openmls-conversation-create-dev
    openmls-conversation-load-check-dev
    openmls-conversation-add-member-dev
    openmls-conversation-join-dev

Naming rules:

    use openmls-*-dev.
    keep dev explicit.
    do not use local-backbone.
    do not use Relay Space naming.
    do not use secure/vault/production naming.
    do not imply mature onboarding.

### 5.4 Shared and command-specific flag policy was formalized

All wrappers should support:

    --sidecar-dir <path>

Default:

    internal/protocol/mls/openmls-sidecar

Device-scoped wrappers should require:

    --sidecar-device-label <label>

Conversation-scoped wrappers should require:

    --conversation <label>

Path-consuming wrappers should require explicit artifact path flags:

    --member-keypackage <path>
    --welcome <path>

Do not auto-derive sidecar labels from Comms state yet.

### 5.5 Output and parsing policy was formalized

Wrapper commands should:

    parse sidecar JSON.
    normalize success output to stable human-readable key/value text.
    preserve sidecar_command.
    preserve sidecar path hints.
    print absolute artifact paths when a relative sidecar-generated artifact hint is present.
    surface sidecar error code/message on failure.
    avoid printing private material.
    avoid claiming vault safety, verified identity, trusted membership, or Relay Space membership.

Sidecar ok=false envelopes should become Go errors.

If parsing fails, wrappers should return an error and must not print success.

### 5.6 State/trust boundary was formalized

Bootstrap wrappers must not mutate:

    Comms state.json
    trust.json
    trust-events.jsonl

Bootstrap wrappers may mutate sidecar dev-local state because that is the purpose of sidecar bootstrap commands.

Current wrapper contract:

    no state.RequireReadyDevice.
    no trust.EvaluateSend.
    no trust mutation.
    no Cypher calls.

This keeps wrappers explicitly sidecar bootstrap/dev utilities, not mature identity UX.

### 5.7 Testing and file-boundary policy was formalized

Future wrapper code should live in:

    carbonstack-comms/internal/app/openmls_bootstrap.go

Future wrapper tests should live in:

    carbonstack-comms/internal/app/openmls_bootstrap_dev_test.go

Keep:

    command registration and usage in internal/app/commands.go.
    message runtime helpers in internal/app/openmls_runtime.go.
    bootstrap wrapper helpers in internal/app/openmls_bootstrap.go.

Future app unit tests should not run Rust sidecar commands directly.

Preferred test seam:

    var runOpenMLSBootstrapSidecarForCommand = runOpenMLSBootstrapSidecar

Suggested helper signature:

    func runOpenMLSBootstrapSidecar(sidecarDir string, sidecarCommand string, args ...string) (openMLSSidecarBootstrapEnvelope, error)

Real Rust sidecar execution should remain in:

    protocol tests.
    smoke script.
    dev-runtime-openmls profile.

### 5.8 Implementation order was formalized

Preferred staged implementation:

    v0.4.13:
      openmls-identity-create-dev
      openmls-identity-status-dev

    v0.4.14:
      openmls-bundle-export-dev
      openmls-conversation-create-dev
      openmls-conversation-load-check-dev

    v0.4.15:
      openmls-conversation-add-member-dev
      openmls-conversation-join-dev

    v0.4.16:
      smoke-script migration recon
      decide whether direct sidecar bootstrap should be replaced with wrapper calls

Do not migrate the smoke script before wrapper tests and direct smoke proof both pass.

### 5.9 docs/README.md was updated

Updated file:

    carbonstack/docs/README.md

Added index entry:

    docs/162-openmls-bootstrap-command-contract-v0.md

Meaning:

    The docs archive now indexes the v0.4.12 dev-only OpenMLS bootstrap command contract.

### 5.10 roadmap/ROADMAP.md was updated

Updated file:

    carbonstack/roadmap/ROADMAP.md

Added section:

    v0.4.12 — dev-only OpenMLS bootstrap command contract

It records:

    wrapper names use openmls-*-dev.
    sidecar labels remain explicit.
    sidecar JSON should be normalized to stable Comms key/value output.
    sidecar path hints should be preserved with absolute paths when applicable.
    wrappers must not mutate Comms state/trust files.
    future implementation should live in openmls_bootstrap.go.
    next recommended rung is v0.4.13 identity create/status dev wrapper implementation.

### 5.11 Validation passed

Commands observed passing:

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go test ./... -count=1
    go run . --profile dev-runtime-openmls --clean-generated
    go run . --profile local-cypher
    go run . --profile doctor
    go run . --profile core --clean-generated

Final commit/push:

    carbonstack:
    0583683 docs: define dev OpenMLS bootstrap command contract

Final repo snapshot:

    carbonstack        0583683 docs: define dev OpenMLS bootstrap command contract
    carbonstack-comms  0a48ae1 refactor: split OpenMLS runtime command helpers
    carbonstack-cypher 9ab994c docs: point to local operator runbook
    carbonstack-os     1bbbe52 docs: clarify CarbonStackOS target direction

---

## 6. Critical Blunders / Lessons to Carry Forward

### 6.1 Recon pasteback was again partially truncated/glitched, but the important contract inputs landed

The v0.4.12 recon pasteback was visibly truncated/glitched near the end of the original command text, similar to v0.4.11. However, enough useful contract data landed:

    current repo heads.
    v0.4.11 planning context.
    roadmap v0.4.11/v0.4.12 area.
    sidecar README command list and boundary.
    sidecar main.rs command handlers / JSON output fragments.
    sidecar JSON output field grep.
    sidecar path helper contract.
    current Comms command registry.
    current OpenMLS runtime helper contract style.
    validation output.
    commit/push output.
    final clean repo snapshot.

Lesson:

    Use the recon for docs/contract work.
    Do a tighter targeted scout before implementation if exact field details are uncertain.
    Do not infer implementation behavior from truncated sections.

### 6.2 Commit metadata warning remains a polish issue

The commit log again showed auto-generated local commit metadata:

    Committer: bitcrusher32 <▓▓>

This is not a functional validation failure, but it remains a professionalism/publishing hygiene issue to address eventually with `git config --global user.name` and `git config --global user.email` if desired.

### 6.3 v0.4.12 is not implementation

v0.4.12 is a docs/contract checkpoint.

Do not assume:

    wrappers exist.
    openmls_bootstrap.go exists.
    openmls_bootstrap_dev_test.go exists.
    the smoke script uses wrappers.
    identity create/status wrappers are implemented.
    send/inbox are OpenMLS-backed.
    local-backbone is justified.
    dev-runtime-openmls belongs in full.

Correct interpretation:

    wrapper contract exists.
    wrapper implementation does not.

---

## 7. Current Blockers / Not Validated

Current blockers and nonvalidated areas after v0.4.12:

    Existing send/inbox remain stub-era.
    Runtime Comms has dev-only OpenMLS commands, a smoke proof, and a manual runner profile, but no mature user-facing UX.
    Sidecar KeyPackage/Welcome/bootstrap setup is still direct dev setup, not Comms runtime UX.
    No dev-only bootstrap wrapper commands exist yet.
    No openmls_bootstrap.go exists yet.
    No openmls_bootstrap_dev_test.go exists yet.
    No Relay Space join model exists yet.
    local-backbone runner profile does not exist and remains reserved.
    dev-runtime-openmls exists but is manual-only.
    dev-runtime-openmls is not included in full.
    dev-runtime-openmls is live-git-umbrella-only.
    dev-runtime-openmls is not part of public release package validation.
    Fresh release-package behavior for dev-runtime-openmls has not been tested as a supported validation surface.
    Post-release public download verification is still useful if not already performed.
    Sanitized PRIME placement in carbonstack/logdoc-loop-system still needs a policy decision.
    Production-safe provider storage and secure vault design remain v0.5.x future work.
    PQ/hybrid ciphersuite migration readiness remains v0.5.x future work.
    Hostile-server rollback/replay/metadata-abuse validation remains future work.
    Android/CarbonStackOS implementation remains later work.
    CarbonStack Relay Space mechanics remain a future concept, not an implementation.
    Public ingress, systemd, cloudflared, real homelab deployment, and production service paths remain out of scope.

Hard nonclaims remain:

    Do not claim production readiness.
    Do not claim production E2EE.
    Do not claim hostile-server safety.
    Do not claim metadata privacy.
    Do not claim Debian production deployability.
    Do not claim systemd readiness.
    Do not claim cloudflared readiness.
    Do not claim Android readiness.
    Do not claim CarbonStackOS readiness.
    Do not claim secure vault/storage.
    Do not claim real homelab validation.
    Do not claim external audit or certification.
    Do not claim CarbonStack Relay Space mechanics are implemented.
    Do not claim local-backbone exists or is justified yet.
    Do not claim local-cypher validates runtime Comms UX.
    Do not claim full validates local-backbone.
    Do not claim full is a deployment command.
    Do not claim dev-runtime-openmls is release-package validation yet.
    Do not claim send/inbox are OpenMLS-backed yet.
    Do not claim PQ/hybrid security or quantum-safe messaging.
    Do not claim bootstrap wrapper commands exist yet.
    Do not let GitHub mirrors supersede Gitea source-of-truth status.

Allowed claim:

    CarbonStackComms now has dev-only openmls-send-dev and openmls-inbox-dev commands.
    CarbonStackComms now has a dev runtime smoke script proving openmls-send-dev -> Cypher -> openmls-inbox-dev --ack.
    CarbonStackComms now has OpenMLS runtime command helper glue split into internal/app/openmls_runtime.go.
    The main carbonstack repo records the smoke proof, the pre-local-backbone assessment, the manual dev-runtime-openmls runner profile, the profile boundary check, the bootstrap-wrapper recon/planning doc, and the bootstrap wrapper command contract.
    carbonstack/tools/carbonstack-validate has a manual dev-runtime-openmls profile.
    dev-runtime-openmls wraps the comms smoke script and validates the dev/pre-alpha OpenMLS application-message CLI path.
    dev-runtime-openmls is live-git-umbrella-only and not included in full.
    dev-runtime-openmls intentionally refuses non-git package-like roots.
    The smoke script starts a temporary local Cypher server, uses temporary Comms state files, creates dev-local OpenMLS sidecar identities/conversation, sends an OpenMLS application message through the Comms CLI, opens it through the Comms CLI, acks after sidecar success, and verifies Bob's inbox is empty after ack.
    v0.4.12 defines the dev-only OpenMLS bootstrap wrapper contract.
    Existing send/inbox remain stub-era.
    v0.4.0 remains the current public broad local deployability pre-release.
    Gitea remains source of truth; GitHub mirrors are secondary push mirrors.

---

## 8. Next Safest Actions

Recommended next rung:

    v0.4.13 implement identity create/status dev wrappers

Focus:

    carbonstack-comms only.
    do targeted recon if exact sidecar identity JSON fields are needed before patch.
    add internal/app/openmls_bootstrap.go.
    add internal/app/openmls_bootstrap_dev_test.go.
    register openmls-identity-create-dev.
    register openmls-identity-status-dev.
    keep sidecar labels explicit.
    normalize sidecar JSON to stable key/value output.
    do not mutate Comms state/trust.
    do not touch send/inbox.
    do not migrate the smoke script yet.
    update carbonstack-comms README only if implementation lands.
    validate app tests, Comms tests, dev-runtime-openmls, local-cypher, doctor, and core.

Preferred implementation sequence after v0.4.13:

    v0.4.14:
      openmls-bundle-export-dev
      openmls-conversation-create-dev
      openmls-conversation-load-check-dev

    v0.4.15:
      openmls-conversation-add-member-dev
      openmls-conversation-join-dev

    v0.4.16:
      smoke-script migration recon / wrapper-based smoke variant decision

Avoid next:

    implementing all wrappers at once without v0.4.13 proving the shape.
    replacing send/inbox immediately.
    replacing the smoke script before wrapper tests exist.
    adding dev-runtime-openmls to full.
    calling any wrapper local-backbone.
    using Relay Space join naming for dev bootstrap wrappers.
    moving dev-runtime-openmls into public release validation without fresh-root behavior checks.
    public ingress.
    cloudflared.
    systemd.
    real homelab deployment.
    Android app.
    CarbonStackOS implementation.
    CarbonStack Relay Space implementation.
    PQ/hybrid ciphersuite implementation.
    production/security claims.

---

## 9. Critical Paths / Functions

### Source-of-truth public surfaces

    [GITEA SOURCE OF TRUTH] https://git.bitcrusher32.win/bitcrusher32/carbonstack
    [RELEASES] https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases
    [CURRENT PUBLIC RELEASE] https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.4.0

GitHub mirrors:

    secondary push mirrors for discoverability/redundancy only.
    Do not treat mirrors as release authority unless project policy changes.

### WSL Debian active working umbrella

    [WSL-DEBIAN] umbrella: ~/repos/carbonstack_umbrella
    [WSL-DEBIAN] carbonstack: ~/repos/carbonstack_umbrella/carbonstack
    [WSL-DEBIAN] carbonstack-comms: ~/repos/carbonstack_umbrella/carbonstack-comms
    [WSL-DEBIAN] carbonstack-cypher: ~/repos/carbonstack_umbrella/carbonstack-cypher
    [WSL-DEBIAN] carbonstack-os: ~/repos/carbonstack_umbrella/carbonstack-os
    [GO-RUNNER] runner dir: ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate

### Runtime integration doc surfaces

    [DOC] v0.4.1 runtime recon:
    carbonstack/docs/155-runtime-comms-openmls-cypher-recon-v0.md

    [DOC] v0.4.2 command contract:
    carbonstack/docs/156-runtime-openmls-command-contract-v0.md

    [DOC] v0.4.6 smoke proof record:
    carbonstack/docs/157-dev-runtime-openmls-smoke-proof-v0.md

    [DOC] v0.4.7 pre-local-backbone assessment:
    carbonstack/docs/158-pre-local-backbone-assessment-v0.md

    [DOC] v0.4.9 runner profile:
    carbonstack/docs/159-dev-runtime-openmls-runner-profile-v0.md

    [DOC] v0.4.10 profile boundary:
    carbonstack/docs/160-dev-runtime-openmls-profile-boundary-v0.md

    [DOC] v0.4.11 bootstrap wrapper recon:
    carbonstack/docs/161-openmls-bootstrap-wrapper-recon-v0.md

    [DOC] v0.4.12 bootstrap command contract:
    carbonstack/docs/162-openmls-bootstrap-command-contract-v0.md

    [DOC] docs index:
    carbonstack/docs/README.md

    [DOC] roadmap:
    carbonstack/roadmap/ROADMAP.md

    [DOC] top README:
    carbonstack/README.md

### Runner surfaces

    [RUNNER] root:
    carbonstack/tools/carbonstack-validate

    [RUNNER] dispatch:
    carbonstack/tools/carbonstack-validate/main.go

    [RUNNER] dev-runtime-openmls implementation:
    carbonstack/tools/carbonstack-validate/dev_runtime_openmls.go

    [RUNNER] runner README:
    carbonstack/tools/carbonstack-validate/README.md

    [RUNNER] local-cypher:
    carbonstack/tools/carbonstack-validate/local_cypher.go

    [RUNNER] release-snapshot:
    carbonstack/tools/carbonstack-validate/release_snapshot.go

    [RUNNER] checksums:
    carbonstack/tools/carbonstack-validate/checksums.go

### Comms runtime surfaces

    [COMMS] CLI entry:
    carbonstack-comms/cmd/comms/main.go

    [COMMS] command registry and non-OpenMLS commands:
    carbonstack-comms/internal/app/commands.go

    [COMMS] OpenMLS message runtime command helpers:
    carbonstack-comms/internal/app/openmls_runtime.go

    [COMMS] Future OpenMLS bootstrap wrapper file:
    carbonstack-comms/internal/app/openmls_bootstrap.go

    [COMMS] Future OpenMLS bootstrap wrapper tests:
    carbonstack-comms/internal/app/openmls_bootstrap_dev_test.go

    [COMMS] openmls-send-dev tests:
    carbonstack-comms/internal/app/openmls_send_dev_test.go

    [COMMS] openmls-inbox-dev tests:
    carbonstack-comms/internal/app/openmls_inbox_dev_test.go

    [COMMS] dev runtime smoke proof:
    carbonstack-comms/scripts/dev-openmls-runtime-smoke.sh

    [COMMS] Cypher client:
    carbonstack-comms/internal/client/cypher.go

    [COMMS] relay helpers:
    carbonstack-comms/internal/relay/cypher_bridge.go
    carbonstack-comms/internal/relay/openmls_artifacts.go

    [COMMS] OpenMLS sidecar dir:
    carbonstack-comms/internal/protocol/mls/openmls-sidecar

    [COMMS] sidecar README:
    carbonstack-comms/internal/protocol/mls/openmls-sidecar/README.md

### Current stub-era Comms functions / behaviors

    send:
      trust policy evaluation
      mock/stub provider encryption
      Cypher envelope submit

    inbox:
      Cypher inbox retrieval
      mock/stub provider plaintext display as stub_plaintext

    ack:
      envelope ack by envelope ID and recipient device ID

### Implemented v0.4.x dev OpenMLS commands

    openmls-send-dev:
      internal/app/openmls_runtime.go
      state.RequireReadyDevice
      trust.EvaluateSend
      runOpenMLSMessageProtect
      sidecar message-protect
      internal/relay.SubmitOpenMLSArtifactEnvelope
      ArtifactKindApplicationMessage
      dev/pre-alpha warning

    openmls-inbox-dev:
      internal/app/openmls_runtime.go
      state.RequireReadyDevice
      Cypher inbox
      skip unsupported envelopes
      internal/relay.WriteOpenMLSArtifactFromEnvelope
      runOpenMLSMessageOpen
      sidecar message-open
      print plaintext after sidecar success
      optional ack only after sidecar success

### Contracted but not yet implemented v0.4.13+ dev OpenMLS bootstrap wrappers

    openmls-identity-create-dev:
      sidecar identity-create
      required --sidecar-device-label

    openmls-identity-status-dev:
      sidecar identity-status
      required --sidecar-device-label

    openmls-bundle-export-dev:
      sidecar public-bundle-export
      required --sidecar-device-label
      optional --write-artifact

    openmls-conversation-create-dev:
      sidecar conversation-create
      required --sidecar-device-label and --conversation

    openmls-conversation-load-check-dev:
      sidecar conversation-load-check
      required --sidecar-device-label and --conversation

    openmls-conversation-add-member-dev:
      sidecar conversation-add-member
      required --sidecar-device-label, --conversation, and --member-keypackage

    openmls-conversation-join-dev:
      sidecar conversation-join
      required --sidecar-device-label, --conversation, and --welcome

### v0.4.5 through v0.4.12 runtime smoke proof

    scripts/dev-openmls-runtime-smoke.sh:
      temp workspace
      temp Cypher binary
      temp Cypher DB
      dynamic loopback Cypher port
      temp Alice/Bob Comms state
      sidecar state reset
      direct sidecar identity/conversation bootstrap
      openmls-send-dev application-message send
      openmls-inbox-dev application-message open + ack
      inbox-empty-after-ack verification
      robust cleanup with INT -> TERM -> KILL escalation

### v0.4.9/v0.4.10 runner profile

    go run . --profile dev-runtime-openmls --clean-generated

    Runner.DevRuntimeOpenMLS:
      CheckRequiredPaths
      CheckLiveGitUmbrella
      toolchain reporting
      ArtifactScan("pre-dev-runtime-openmls")
      RunStep for carbonstack-comms/scripts/dev-openmls-runtime-smoke.sh
      ArtifactScan("post-dev-runtime-openmls")
      explicit nonclaim summary

    Runner.CheckLiveGitUmbrella:
      requires .git markers in carbonstack, carbonstack-comms, carbonstack-cypher
      refuses non-git package-like roots

### OpenMLS sidecar commands relevant to v0.4.x

    provider-info
    identity-create
    identity-status
    public-bundle-export
    conversation-create
    conversation-load-check
    conversation-add-member
    conversation-join
    message-protect
    message-open

### Relay/API functions relevant to runtime integration

    internal/client.CypherClient.SubmitEnvelope
    internal/client.CypherClient.Inbox
    internal/client.CypherClient.AckEnvelope
    internal/relay.SubmitOpenMLSArtifactEnvelope
    internal/relay.WriteOpenMLSArtifactFromEnvelope

### Existing v0.4.0 package validation

    cd <package-root>/carbonstack/tools/carbonstack-validate
    go run . --profile verify-checksums --root <package-root>
    go run . --profile full --root <package-root> --clean-generated

### Live umbrella validation

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go run . --profile dev-runtime-openmls --clean-generated
    go run . --profile local-cypher
    go run . --profile doctor
    go run . --profile core --clean-generated

### Envelope/protocol constants

Accepted opaque OpenMLS application-message pair:

    content_type=carbonstack.mls.application-message.v0
    protocol_version=carbonstack-openmls-sidecar-v0

Historical invalid pair validated by local-cypher:

    content_type=carbonstack.message.text.stub.v0
    protocol_version=carbonstack-openmls-sidecar-v0

Expected result:

    HTTP 400
    error code unsupported_protocol_version

---

## 10. v0.4.x Forward Plan

### v0.4.13 — identity create/status wrapper implementation

Expected:

    carbonstack-comms-only implementation.
    add internal/app/openmls_bootstrap.go.
    add internal/app/openmls_bootstrap_dev_test.go.
    register openmls-identity-create-dev.
    register openmls-identity-status-dev.
    keep command registry in commands.go.
    keep output as stable key/value text.
    parse sidecar JSON through injected test seams.
    do not mutate Comms state/trust.
    update carbonstack-comms README if implementation lands.
    validate app tests, full Comms tests, dev-runtime-openmls, local-cypher, doctor, and core.

### v0.4.14+

Possible staged implementation:

    v0.4.14:
      openmls-bundle-export-dev
      openmls-conversation-create-dev
      openmls-conversation-load-check-dev

    v0.4.15:
      openmls-conversation-add-member-dev
      openmls-conversation-join-dev

    v0.4.16:
      smoke-script migration recon / wrapper-based smoke variant decision

### v0.4.x later polish

Likely needs:

    exact sidecar setup runbook.
    clearer dev state reset semantics.
    better command output shape.
    decision on whether stub-era send/inbox are aliased, replaced, or deprecated.
    decision on when local-backbone naming stops being premature.

---

## 11. v0.5.x and Later Carry-Forward

### v0.5.x

Local storage, trust state, provider state, vault design, and PQ/hybrid ciphersuite migration readiness.

Do not start PQ/hybrid implementation before the runtime OpenMLS/Cypher path exists and is documented. PQ/hybrid migration touches provider state, ciphersuite metadata, trust display, artifact size, reinit/rekey/recovery policy, and claim boundaries.

### v0.6.x+

Hostile-server and abuse-resistance harnesses.

### v0.7.x+

Deployability and operations hardening.

### v0.8.x+

Documented self-pentest / adversarial validation.

### v0.9.x+

Claim-boundary review.

### v0.10.x+

Android backend/app work.

### v1.x.x

Public app/server major epoch and much later CarbonStackOS/appliance research, if justified.

---

## 12. Lean Breakpoint Summary

    CarbonStack v0.4.12 is the dev-only OpenMLS bootstrap command-contract checkpoint. v0.4.0 remains the current public broad local deployability pre-release. carbonstack is now at 0583683 docs: define dev OpenMLS bootstrap command contract. carbonstack-comms remains at 0a48ae1 refactor: split OpenMLS runtime command helpers. v0.4.12 adds docs/162-openmls-bootstrap-command-contract-v0.md, updates docs/README.md, and updates roadmap/ROADMAP.md. The checkpoint defines future openmls-*-dev bootstrap wrapper names, flags, output shape, sidecar JSON parsing policy, path-hint/absolute-path policy, no-Comms-state/trust-mutation boundary, app-test injection seam, implementation file boundary, and staged rollout. No wrapper commands were implemented. Validation passed for go test ./... in the runner, dev-runtime-openmls --clean-generated, local-cypher, doctor, and core --clean-generated, with final clean repo heads across all four repos. Next safe rung: v0.4.13 identity create/status wrapper implementation.

---

## 13. Preserved Immediate Previous Handoff: v0.4.11

The following is the previous v0.4.11 handoff. Where it conflicts with the v0.4.12 overlay above, v0.4.12 wins for current state.


# CarbonStack LogDoc v0.4.11

**Last updated:** 2026-06-05 local session  
**Owner/operator:** bitcrusher32  
**Assistant/model:** GPT-5.5 Thinking  
**Current phase:** Phase 2F / **v0.4.x OpenMLS bootstrap wrapper recon/planning checkpoint complete**. The v0.4.0 public broad local deployability pre-release remains the current public release. v0.4.5 proved the dev runtime OpenMLS CLI path; v0.4.6 recorded that smoke proof in the main `carbonstack` repo; v0.4.7 assessed the proof and kept `local-backbone` reserved; v0.4.8 split OpenMLS runtime command glue into `carbonstack-comms/internal/app/openmls_runtime.go`; v0.4.9 added a **manual**, live-umbrella-only `dev-runtime-openmls` validation profile; v0.4.10 recorded that profile boundary as stable; v0.4.11 now records the sidecar bootstrap wrapper recon/planning decision. `carbonstack` is now at `8ee3ebb (HEAD -> main, origin/main, origin/HEAD) docs: plan dev OpenMLS bootstrap wrappers`; `carbonstack-comms` remains at `0a48ae1 (HEAD -> main, origin/main, origin/HEAD) refactor: split OpenMLS runtime command helpers`; `carbonstack-cypher` remains at `9ab994c (HEAD -> main, origin/main, origin/HEAD) docs: point to local operator runbook`; `carbonstack-os` remains at `1bbbe52 (HEAD -> main, origin/main, origin/HEAD) docs: clarify CarbonStackOS target direction`.

**Version schema:** `v[scope].[timeline]`. This file is `v0.4.11`, the OpenMLS bootstrap wrapper recon/planning checkpoint after the v0.4.10 `dev-runtime-openmls` profile-boundary evidence checkpoint. Per LogDoc V2 practice, this Markdown file preserves current operational continuity and carries the v0.4.x runtime-integration thread forward; the JSON breakpoint remains the lean machine-readable handoff state.

**Compression note:** The detailed v0.3.0 through v0.3.36PRIME process ledger remains archived in PRIME/sanitized PRIME. v0.4.0 compressed that pre-release sprint; v0.4.1 preserved runtime recon; v0.4.2 preserved the command-contract decision; v0.4.3 preserved the first dev-only OpenMLS send-side implementation; v0.4.4 preserved the first dev-only OpenMLS receive/open/optional-ack implementation; v0.4.5 preserved the first committed dev runtime smoke proof; v0.4.6 preserved project-level documentation/status alignment for that proof; v0.4.7 preserved the decision boundary before local-backbone/runner-profile promotion; v0.4.8 preserved behavior-preserving OpenMLS command-helper extraction; v0.4.9 preserved the manual `dev-runtime-openmls` runner-profile implementation; v0.4.10 preserved the profile boundary evidence; v0.4.11 now preserves the OpenMLS bootstrap wrapper recon and the decision to write a command contract before implementation.

---

## 1. Project Goal

**Active goal:** Continue moving CarbonStackComms from stub-era runtime `send`/`inbox` behavior toward explicit dev-only OpenMLS/Cypher-backed runtime commands and repeatable validation, while preserving claim discipline. v0.4.11 did not implement new commands; it recorded the next design seam: direct OpenMLS sidecar bootstrap in the smoke script should be planned as explicit dev-only Comms wrapper commands before any smoke-script migration or local-backbone naming.

CarbonStack’s broader project intent remains:

- **CarbonStack** is the secure-communications appliance-stack project and shared doctrine/release authority.
- **CarbonStackComms** is the text-first messaging client and current OpenMLS sidecar/relay integration surface.
- **CarbonStackCypher** is the self-hostable relay/storage stack.
- **CarbonStackOS** is future constrained Android-derived appliance OS work and remains out of current runnable validation packages.
- The `carbonstack` Gitea repo remains the source of truth for release state, release assets, docs, and public claims.
- GitHub mirrors exist as secondary push mirrors only; Gitea remains authoritative.

**North star:**

> Its goal is not to be a better smartphone. Its goal is to stop being a smartphone.

**Core doctrine:**

> Every feature is guilty until it proves it does not add unacceptable parser, network, sensor, identity, or filesystem authority.

**Current reality after v0.4.11:** the runtime proof remains valid, the `dev-runtime-openmls` profile remains manual/live-umbrella-only, and sidecar bootstrap wrappers are now planned but not implemented.

Correct interpretation:

    openmls-send-dev is a dev-only OpenMLS send-side command.
    openmls-inbox-dev is a dev-only OpenMLS receive/open/optional-ack command.
    scripts/dev-openmls-runtime-smoke.sh proves the dev CLI application-message path:
      openmls-send-dev -> Cypher -> openmls-inbox-dev --ack
    tools/carbonstack-validate has a manual dev-runtime-openmls profile that wraps that smoke script.
    dev-runtime-openmls is repeatable from the live umbrella checkout.
    dev-runtime-openmls is live-git-umbrella-only for now.
    dev-runtime-openmls intentionally refuses non-git package-like roots.
    dev-runtime-openmls is not included in full.
    dev-runtime-openmls is not release-package validation yet.
    docs/160-dev-runtime-openmls-profile-boundary-v0.md records the profile boundary check.
    docs/161-openmls-bootstrap-wrapper-recon-v0.md records the bootstrap wrapper recon and planning boundary.
    internal/app/openmls_runtime.go remains the OpenMLS message runtime command helper file.
    Future bootstrap wrapper code should likely live in internal/app/openmls_bootstrap.go.
    Existing send/inbox remain stub-era.
    local-backbone remains reserved.

Hard nonclaims remain: v0.4.11 does **not** prove production readiness, production E2EE, hostile-server safety, metadata privacy, Android readiness, CarbonStackOS readiness, secure vault/storage, public ingress safety, real homelab validation, audit, certification, PQ/hybrid security, mature messenger UX, release-package runtime validation, Relay Space join/onboarding, or local-backbone readiness.

---

## 2. Current Public Release State

Public release:

    https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.4.0

Release title:

    CarbonStack v0.4.0 Broad Local Deployability Pre-Release

Release tag:

    v0.4.0

Release commit:

    24ac9fc494 / 24ac9fc docs: record v0.4.0 package rehearsal

Gitea release status:

    Pre-Release

Release assets attached:

    carbonstack-v0.4.0-broad-local-deployability-pre-release.tgz
    carbonstack-v0.4.0-release-manifest.json
    carbonstack-v0.4.0-package-checksums.txt
    carbonstack-v0.4.0-asset-checksums.txt
    carbonstack-v0.4.0-validation-freeze.md
    v0.4.0-testing-runbook.md
    v0.4.0-release-notes.md
    LICENSE

Important release-page warning:

    The default Gitea Source Code ZIP/TAR.GZ downloads are auto-generated carbonstack-only archives.
    The intended runnable validation artifact is the attached multi-repo release package plus companion assets.

Post-release public surface:

    carbonstack README is evergreen.
    It points interested parties to the Releases page for runnable / known-good artifacts.
    It states Gitea is the source of truth.
    GitHub mirrors are secondary push mirrors, not release authority.

---

## 3. Current Repo Heads

    carbonstack        8ee3ebb (HEAD -> main, origin/main, origin/HEAD) docs: plan dev OpenMLS bootstrap wrappers
    carbonstack-comms  0a48ae1 (HEAD -> main, origin/main, origin/HEAD) refactor: split OpenMLS runtime command helpers
    carbonstack-cypher 9ab994c (HEAD -> main, origin/main, origin/HEAD) docs: point to local operator runbook
    carbonstack-os     1bbbe52 (HEAD -> main, origin/main, origin/HEAD) docs: clarify CarbonStackOS target direction

Current public release tag:

    v0.4.0 -> 24ac9fc494 / 24ac9fc docs: record v0.4.0 package rehearsal

Post-release public-surface cleanup:

    6ba7aab docs: make front readme evergreen

v0.4.1 runtime recon commit:

    83d2f0f docs: record runtime OpenMLS Comms recon

v0.4.2 command contract commit:

    a799764 docs: define runtime OpenMLS command contract

v0.4.3 send implementation commit:

    39fb287 feat: add dev OpenMLS send command

v0.4.4 inbox implementation commit:

    65bc707 feat: add dev OpenMLS inbox command

v0.4.5 smoke proof commit:

    8e6e8b4 test: add dev OpenMLS runtime smoke proof

v0.4.6 main-doc alignment commit:

    3cd77a8 docs: record dev OpenMLS runtime smoke proof

v0.4.7 assessment commit:

    dc2d16c docs: assess pre-local-backbone runtime proof

v0.4.8 extraction commit:

    0a48ae1 refactor: split OpenMLS runtime command helpers

v0.4.9 runner-profile commit:

    8eeadb2 feat: add dev OpenMLS runtime validation profile

v0.4.10 boundary-doc commit:

    7384f43 docs: record dev OpenMLS runtime profile boundary

Current v0.4.11 wrapper-recon commit:

    8ee3ebb docs: plan dev OpenMLS bootstrap wrappers

Continuity note:

    `8ee3ebb` is a `carbonstack` docs/planning commit after the `v0.4.0` public release tag.
    It does not retag v0.4.0.
    The release remains attached to `24ac9fc494`.

---

## 4. Current Validated / Known State

Validated / known at the v0.4.11 checkpoint:

    [RELEASE] v0.4.0 Gitea release exists, is marked Pre-Release, and is attached to 24ac9fc494.
    [PACKAGE] v0.4.0 package includes carbonstack, carbonstack-comms, carbonstack-cypher, and release metadata.
    [PACKAGE] carbonstack-os is deliberately excluded from the runnable package.
    [RUNNER] full means release-snapshot followed by local-cypher.
    [RUNNER] release-snapshot validates package layout, metadata, strict pre-test artifact scan, package checksums, and core.
    [RUNNER] core validates doctor, OpenMLS real-Cypher lifecycle, Comms package tests, Cypher package tests, and artifact scans.
    [RUNNER] local-cypher validates Cypher-only lifecycle, restart/persistence against the same temporary DB, positive opaque OpenMLS envelope/ack flow, and historical negative stub-text/OpenMLS protocol pairing.
    [RUNNER] dev-runtime-openmls validates the dev/pre-alpha OpenMLS application-message runtime CLI smoke proof.
    [RUNNER] dev-runtime-openmls is manual-only and live-umbrella-only for now.
    [RUNNER] dev-runtime-openmls is not included in full.
    [RUNNER] dev-runtime-openmls checks live git markers for carbonstack, carbonstack-comms, and carbonstack-cypher.
    [RUNNER] dev-runtime-openmls refuses non-git package-like roots as intended.
    [RUNNER] dev-runtime-openmls wraps carbonstack-comms/scripts/dev-openmls-runtime-smoke.sh.
    [DOCS] docs/159-dev-runtime-openmls-runner-profile-v0.md records the v0.4.9 manual runner profile.
    [DOCS] docs/160-dev-runtime-openmls-profile-boundary-v0.md records the v0.4.10 boundary check.
    [DOCS] docs/161-openmls-bootstrap-wrapper-recon-v0.md records the v0.4.11 bootstrap wrapper recon/planning boundary.
    [DOCS] docs/README.md indexes docs/161-openmls-bootstrap-wrapper-recon-v0.md.
    [DOCS] roadmap/ROADMAP.md records v0.4.11 and recommends v0.4.12 dev-only OpenMLS bootstrap command contract next.
    [COMMS] openmls-send-dev exists in carbonstack-comms.
    [COMMS] openmls-inbox-dev exists in carbonstack-comms.
    [COMMS] scripts/dev-openmls-runtime-smoke.sh exists in carbonstack-comms.
    [COMMS] smoke script still uses direct sidecar bootstrap calls for identity, public bundle, conversation create, add-member, and join.
    [COMMS] internal/app/openmls_runtime.go contains OpenMLS message runtime command helper code.
    [COMMS] future bootstrap wrapper code should likely live in internal/app/openmls_bootstrap.go.
    [COMMS] internal/app/commands.go remains the command registry / non-OpenMLS command surface.
    [COMMS] existing send/inbox/ack remain in place and stub-era.
    [MIRRORS] all four repos have GitHub push-mirror functionality, with Gitea remaining source of truth.

v0.4.11 validation result from the final log:

    go test ./... -count=1 passed for carbonstack/tools/carbonstack-validate.
    go run . --profile dev-runtime-openmls --clean-generated passed.
    local-cypher passed.
    doctor passed.
    core --clean-generated passed.
    docs/161-openmls-bootstrap-wrapper-recon-v0.md was added.
    docs/README.md was updated to index the v0.4.11 doc.
    roadmap/ROADMAP.md was updated to record v0.4.11 and recommend v0.4.12.
    carbonstack was committed and pushed at:
      8ee3ebb docs: plan dev OpenMLS bootstrap wrappers
    final repo status showed all four repos clean.

Expected live umbrella validation commands now include:

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go run . --profile dev-runtime-openmls --clean-generated
    go run . --profile local-cypher
    go run . --profile doctor
    go run . --profile core --clean-generated

Expected v0.4.0 public package validation commands remain unchanged:

    cd <package-root>/carbonstack/tools/carbonstack-validate
    go run . --profile verify-checksums --root <package-root>
    go run . --profile full --root <package-root> --clean-generated

Current direct smoke command remains:

    cd ~/repos/carbonstack_umbrella/carbonstack-comms
    scripts/dev-openmls-runtime-smoke.sh

---

## 5. v0.4.11 Work Completed

### 5.1 Bootstrap wrapper scout confirmed the next seam

The v0.4.11 scout inspected:

    v0.4.10 profile-boundary document.
    v0.4.x roadmap section.
    v0.5.x state/trust/vault planning boundary.
    carbonstack-comms/scripts/dev-openmls-runtime-smoke.sh.
    smoke-script sidecar command sequence.
    carbonstack-comms README dev runtime sections.
    carbonstack-comms internal/app command registry.
    carbonstack-comms internal/app/openmls_runtime.go.
    openmls-send-dev and openmls-inbox-dev tests.
    OpenMLS sidecar README command list and boundary.
    OpenMLS sidecar Rust command definitions and flags.
    existing protocol lifecycle tests as contract source.
    Comms state/trust/label surfaces.

Key findings:

    The smoke script currently bootstraps sidecar state directly.
    Direct bootstrap is acceptable for the current smoke proof, but not mature Comms runtime UX.
    Candidate wrappers should keep the explicit openmls-*-dev pattern.
    Sidecar labels should remain explicit for now instead of being auto-derived from Comms labels.
    Wrapper code should likely live in openmls_bootstrap.go rather than expanding openmls_runtime.go.
    A command-contract doc should exist before implementation.
    The smoke script should not be migrated until wrappers have contracts, tests, and stable behavior.
    local-backbone remains premature.

### 5.2 Current direct sidecar bootstrap sequence was documented

Current direct sidecar bootstrap calls in the smoke script:

    identity-create --device-label <alice-label>
    identity-create --device-label <bob-label>
    public-bundle-export --device-label <bob-label> --write-artifact
    conversation-create --device-label <alice-label> --conversation-label <conversation-label>
    conversation-add-member --device-label <alice-label> --conversation-label <conversation-label> --member-keypackage <path>
    conversation-join --device-label <bob-label> --conversation-label <conversation-label> --welcome <path>

These are setup calls.

The actual runtime proof target remains:

    openmls-send-dev -> Cypher -> openmls-inbox-dev --ack

### 5.3 Candidate wrapper names were accepted for future contract work

Recommended names:

    openmls-identity-create-dev
    openmls-identity-status-dev
    openmls-bundle-export-dev
    openmls-conversation-create-dev
    openmls-conversation-load-check-dev
    openmls-conversation-add-member-dev
    openmls-conversation-join-dev

Avoid names:

    local-backbone
    backbone
    join-space
    relay-space
    production
    secure
    vault
    setup

Reason:

    openmls-*-dev matches existing openmls-send-dev and openmls-inbox-dev.
    dev suffix keeps maturity loud.
    names avoid importing future Relay Space or local-backbone assumptions.

### 5.4 Suggested code organization was recorded

Future wrapper code should likely live at:

    carbonstack-comms/internal/app/openmls_bootstrap.go

Keep:

    command registration and usage in commands.go.
    message protect/open runtime glue in openmls_runtime.go.
    bootstrap wrapper glue in openmls_bootstrap.go.

Testing should follow the existing injection-seam style from openmls_send_dev_test.go and openmls_inbox_dev_test.go, not rely on full Rust sidecar execution in app unit tests.

### 5.5 docs/161-openmls-bootstrap-wrapper-recon-v0.md was added

New file:

    carbonstack/docs/161-openmls-bootstrap-wrapper-recon-v0.md

It records:

    v0.4.11 is recon/planning only.
    Direct sidecar bootstrap exists in the smoke script.
    Candidate dev-only wrappers are reasonable.
    Wrappers should keep openmls-*-dev naming.
    Wrapper code should likely live in openmls_bootstrap.go.
    Sidecar labels should remain explicit.
    Output should be stable key/value output.
    Tests should use injected sidecar seams.
    Do not immediately replace the smoke script.
    Do not call this local-backbone.
    Keep v0.5.x state/trust/vault/PQ work deferred.
    Recommended next rung is v0.4.12 dev-only OpenMLS bootstrap command contract.

### 5.6 docs/README.md was updated

Updated file:

    carbonstack/docs/README.md

Added index entry:

    docs/161-openmls-bootstrap-wrapper-recon-v0.md

Meaning:

    The docs archive now indexes the v0.4.11 recon/planning doc for dev-only OpenMLS bootstrap wrappers.

### 5.7 roadmap/ROADMAP.md was updated

Updated file:

    carbonstack/roadmap/ROADMAP.md

Added section:

    v0.4.11 — OpenMLS bootstrap wrapper recon

It records:

    Direct sidecar bootstrap in scripts/dev-openmls-runtime-smoke.sh was inspected.
    Candidate wrapper names should follow openmls-*-dev.
    Wrapper planning should stay separate from local-backbone and Relay Space naming.
    Future wrapper code should likely live in carbonstack-comms/internal/app/openmls_bootstrap.go.
    The smoke script should not be migrated until wrapper contracts and tests exist.
    Next recommended rung is v0.4.12 dev-only OpenMLS bootstrap command contract.

### 5.8 Validation passed

Commands observed passing:

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go test ./... -count=1
    go run . --profile dev-runtime-openmls --clean-generated
    go run . --profile local-cypher
    go run . --profile doctor
    go run . --profile core --clean-generated

Final commit/push:

    carbonstack:
    8ee3ebb docs: plan dev OpenMLS bootstrap wrappers

Final repo snapshot:

    carbonstack        8ee3ebb docs: plan dev OpenMLS bootstrap wrappers
    carbonstack-comms  0a48ae1 refactor: split OpenMLS runtime command helpers
    carbonstack-cypher 9ab994c docs: point to local operator runbook
    carbonstack-os     1bbbe52 docs: clarify CarbonStackOS target direction

---

## 6. Critical Blunders / Lessons to Carry Forward

### 6.1 Scout pasteback was partially glitched/truncated, but enough landed

The v0.4.11 scout pasteback showed a stray/glitched tail around the original recon command:

    echo "===== V0.4.11 SIDECAR BOOTSTRAP WRAPPER RECON COMPLETE ====="stack-os; do1?"d/inbox dev?"of target?"ersation\|trust\|label\|sanitizeLabel" \n\|message

Despite that, enough useful recon landed:

    current repo heads and clean status.
    v0.4.10 boundary context.
    roadmap v0.4.x/v0.5.x area.
    smoke script full body and sidecar bootstrap sequence.
    README dev runtime command sections.
    command registry / usage / current OpenMLS dispatch.
    openmls_runtime.go.
    current OpenMLS app tests.
    sidecar README command list and boundary.
    sidecar Rust command definitions / flags.
    lifecycle-test contract fragments.
    final validation log and final clean repo snapshot.

Lesson:

    If a giant scout pasteback gets mangled, do not overread missing details.
    Use it for planning only unless the exact implementation-critical details are fully present.
    For implementation, do another targeted scout if needed.

### 6.2 First push failed authentication, follow-up WSL push succeeded

As in v0.4.9, the first push after the v0.4.11 commit failed authentication from the Windows/PowerShell surface:

    remote: Failed to authenticate user
    fatal: Authentication failed for 'https://git.bitcrusher32.win/bitcrusher32/carbonstack.git/'

The user re-entered WSL Debian, confirmed the repo was one commit ahead, and pushed successfully:

    7384f43..8ee3ebb  main -> main

Final snapshot showed:

    carbonstack 8ee3ebb (HEAD -> main, origin/main, origin/HEAD)

Lesson:

    Trust the final repo-head snapshot.
    Authentication hiccups are not project-state failures if a later push succeeds and origin/main matches HEAD.

### 6.3 Commit metadata warning remains a polish issue

The commit log again showed auto-generated local commit metadata:

    Committer: bitcrusher32 <▓▓>

This is not a functional validation failure, but it remains a professionalism/publishing hygiene issue to address eventually with `git config --global user.name` and `git config --global user.email` if desired.

### 6.4 No implementation details should be inferred beyond the planning doc

v0.4.11 is a docs/planning checkpoint.

Do not assume:

    wrappers exist.
    openmls_bootstrap.go exists.
    smoke script uses wrappers.
    send/inbox are OpenMLS-backed.
    local-backbone is justified.
    dev-runtime-openmls belongs in full.

Correct interpretation:

    wrapper planning exists.
    wrapper implementation does not.

---

## 7. Current Blockers / Not Validated

Current blockers and nonvalidated areas after v0.4.11:

    Existing send/inbox remain stub-era.
    Runtime Comms has dev-only OpenMLS commands, a smoke proof, and a manual runner profile, but no mature user-facing UX.
    Sidecar KeyPackage/Welcome/bootstrap setup is still direct dev setup, not Comms runtime UX.
    No dev-only bootstrap wrapper commands exist yet.
    No openmls_bootstrap.go exists yet.
    No Relay Space join model exists yet.
    local-backbone runner profile does not exist and remains reserved.
    dev-runtime-openmls exists but is manual-only.
    dev-runtime-openmls is not included in full.
    dev-runtime-openmls is live-git-umbrella-only.
    dev-runtime-openmls is not part of public release package validation.
    Fresh release-package behavior for dev-runtime-openmls has not been tested as a supported validation surface.
    Post-release public download verification is still useful if not already performed.
    Sanitized PRIME placement in carbonstack/logdoc-loop-system still needs a policy decision.
    Production-safe provider storage and secure vault design remain v0.5.x future work.
    PQ/hybrid ciphersuite migration readiness remains v0.5.x future work.
    Hostile-server rollback/replay/metadata-abuse validation remains future work.
    Android/CarbonStackOS implementation remains later work.
    CarbonStack Relay Space mechanics remain a future concept, not an implementation.
    Public ingress, systemd, cloudflared, real homelab deployment, and production service paths remain out of scope.

Hard nonclaims remain:

    Do not claim production readiness.
    Do not claim production E2EE.
    Do not claim hostile-server safety.
    Do not claim metadata privacy.
    Do not claim Debian production deployability.
    Do not claim systemd readiness.
    Do not claim cloudflared readiness.
    Do not claim Android readiness.
    Do not claim CarbonStackOS readiness.
    Do not claim secure vault/storage.
    Do not claim real homelab validation.
    Do not claim external audit or certification.
    Do not claim CarbonStack Relay Space mechanics are implemented.
    Do not claim local-backbone exists or is justified yet.
    Do not claim local-cypher validates runtime Comms UX.
    Do not claim full validates local-backbone.
    Do not claim full is a deployment command.
    Do not claim dev-runtime-openmls is release-package validation yet.
    Do not claim send/inbox are OpenMLS-backed yet.
    Do not claim PQ/hybrid security or quantum-safe messaging.
    Do not let GitHub mirrors supersede Gitea source-of-truth status.

Allowed claim:

    CarbonStackComms now has dev-only openmls-send-dev and openmls-inbox-dev commands.
    CarbonStackComms now has a dev runtime smoke script proving openmls-send-dev -> Cypher -> openmls-inbox-dev --ack.
    CarbonStackComms now has OpenMLS runtime command helper glue split into internal/app/openmls_runtime.go.
    The main carbonstack repo records the smoke proof, the pre-local-backbone assessment, the manual dev-runtime-openmls runner profile, the profile boundary check, and the bootstrap-wrapper recon/planning doc.
    carbonstack/tools/carbonstack-validate has a manual dev-runtime-openmls profile.
    dev-runtime-openmls wraps the comms smoke script and validates the dev/pre-alpha OpenMLS application-message CLI path.
    dev-runtime-openmls is live-git-umbrella-only and not included in full.
    dev-runtime-openmls intentionally refuses non-git package-like roots.
    The smoke script starts a temporary local Cypher server, uses temporary Comms state files, creates dev-local OpenMLS sidecar identities/conversation, sends an OpenMLS application message through the Comms CLI, opens it through the Comms CLI, acks after sidecar success, and verifies Bob's inbox is empty after ack.
    v0.4.11 identifies dev-only OpenMLS bootstrap wrappers as a reasonable next planning target.
    Existing send/inbox remain stub-era.
    v0.4.0 remains the current public broad local deployability pre-release.
    Gitea remains source of truth; GitHub mirrors are secondary push mirrors.

---

## 8. Next Safest Actions

Recommended next rung:

    v0.4.12 dev-only OpenMLS bootstrap command contract

Focus:

    define exact flags.
    define exact output fields.
    define sidecar JSON parsing policy.
    define path-hint / absolute-path policy.
    define testing seams.
    define implementation order.
    define smoke-script migration conditions.
    preserve explicit openmls-*-dev naming.
    preserve explicit sidecar labels for now.

Candidate wrapper commands:

    openmls-identity-create-dev
    openmls-identity-status-dev
    openmls-bundle-export-dev
    openmls-conversation-create-dev
    openmls-conversation-load-check-dev
    openmls-conversation-add-member-dev
    openmls-conversation-join-dev

Preferred implementation sequence after contract:

    v0.4.13:
      openmls-identity-create-dev
      openmls-identity-status-dev

    v0.4.14:
      openmls-bundle-export-dev
      openmls-conversation-create-dev
      openmls-conversation-load-check-dev

    v0.4.15:
      openmls-conversation-add-member-dev
      openmls-conversation-join-dev

    v0.4.16:
      smoke-script migration recon / wrapper-based smoke variant decision

Avoid next:

    implementing all wrappers without a contract.
    replacing send/inbox immediately.
    replacing the smoke script before wrapper tests exist.
    adding dev-runtime-openmls to full.
    calling any wrapper local-backbone.
    using Relay Space join naming for dev bootstrap wrappers.
    moving dev-runtime-openmls into public release validation without fresh-root behavior checks.
    public ingress.
    cloudflared.
    systemd.
    real homelab deployment.
    Android app.
    CarbonStackOS implementation.
    CarbonStack Relay Space implementation.
    PQ/hybrid ciphersuite implementation.
    production/security claims.

---

## 9. Critical Paths / Functions

### Source-of-truth public surfaces

    [GITEA SOURCE OF TRUTH] https://git.bitcrusher32.win/bitcrusher32/carbonstack
    [RELEASES] https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases
    [CURRENT PUBLIC RELEASE] https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.4.0

GitHub mirrors:

    secondary push mirrors for discoverability/redundancy only.
    Do not treat mirrors as release authority unless project policy changes.

### WSL Debian active working umbrella

    [WSL-DEBIAN] umbrella: ~/repos/carbonstack_umbrella
    [WSL-DEBIAN] carbonstack: ~/repos/carbonstack_umbrella/carbonstack
    [WSL-DEBIAN] carbonstack-comms: ~/repos/carbonstack_umbrella/carbonstack-comms
    [WSL-DEBIAN] carbonstack-cypher: ~/repos/carbonstack_umbrella/carbonstack-cypher
    [WSL-DEBIAN] carbonstack-os: ~/repos/carbonstack_umbrella/carbonstack-os
    [GO-RUNNER] runner dir: ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate

### Runtime integration doc surfaces

    [DOC] v0.4.1 runtime recon:
    carbonstack/docs/155-runtime-comms-openmls-cypher-recon-v0.md

    [DOC] v0.4.2 command contract:
    carbonstack/docs/156-runtime-openmls-command-contract-v0.md

    [DOC] v0.4.6 smoke proof record:
    carbonstack/docs/157-dev-runtime-openmls-smoke-proof-v0.md

    [DOC] v0.4.7 pre-local-backbone assessment:
    carbonstack/docs/158-pre-local-backbone-assessment-v0.md

    [DOC] v0.4.9 runner profile:
    carbonstack/docs/159-dev-runtime-openmls-runner-profile-v0.md

    [DOC] v0.4.10 profile boundary:
    carbonstack/docs/160-dev-runtime-openmls-profile-boundary-v0.md

    [DOC] v0.4.11 bootstrap wrapper recon:
    carbonstack/docs/161-openmls-bootstrap-wrapper-recon-v0.md

    [DOC] docs index:
    carbonstack/docs/README.md

    [DOC] roadmap:
    carbonstack/roadmap/ROADMAP.md

    [DOC] top README:
    carbonstack/README.md

### Runner surfaces

    [RUNNER] root:
    carbonstack/tools/carbonstack-validate

    [RUNNER] dispatch:
    carbonstack/tools/carbonstack-validate/main.go

    [RUNNER] dev-runtime-openmls implementation:
    carbonstack/tools/carbonstack-validate/dev_runtime_openmls.go

    [RUNNER] runner README:
    carbonstack/tools/carbonstack-validate/README.md

    [RUNNER] local-cypher:
    carbonstack/tools/carbonstack-validate/local_cypher.go

    [RUNNER] release-snapshot:
    carbonstack/tools/carbonstack-validate/release_snapshot.go

    [RUNNER] checksums:
    carbonstack/tools/carbonstack-validate/checksums.go

### Comms runtime surfaces

    [COMMS] CLI entry:
    carbonstack-comms/cmd/comms/main.go

    [COMMS] command registry and non-OpenMLS commands:
    carbonstack-comms/internal/app/commands.go

    [COMMS] OpenMLS message runtime command helpers:
    carbonstack-comms/internal/app/openmls_runtime.go

    [COMMS] Future OpenMLS bootstrap wrapper candidate:
    carbonstack-comms/internal/app/openmls_bootstrap.go

    [COMMS] openmls-send-dev tests:
    carbonstack-comms/internal/app/openmls_send_dev_test.go

    [COMMS] openmls-inbox-dev tests:
    carbonstack-comms/internal/app/openmls_inbox_dev_test.go

    [COMMS] dev runtime smoke proof:
    carbonstack-comms/scripts/dev-openmls-runtime-smoke.sh

    [COMMS] Cypher client:
    carbonstack-comms/internal/client/cypher.go

    [COMMS] relay helpers:
    carbonstack-comms/internal/relay/cypher_bridge.go
    carbonstack-comms/internal/relay/openmls_artifacts.go

    [COMMS] OpenMLS sidecar dir:
    carbonstack-comms/internal/protocol/mls/openmls-sidecar

    [COMMS] sidecar README:
    carbonstack-comms/internal/protocol/mls/openmls-sidecar/README.md

### Current stub-era Comms functions / behaviors

    send:
      trust policy evaluation
      mock/stub provider encryption
      Cypher envelope submit

    inbox:
      Cypher inbox retrieval
      mock/stub provider plaintext display as stub_plaintext

    ack:
      envelope ack by envelope ID and recipient device ID

### Implemented v0.4.x dev OpenMLS commands

    openmls-send-dev:
      internal/app/openmls_runtime.go
      state.RequireReadyDevice
      trust.EvaluateSend
      runOpenMLSMessageProtect
      sidecar message-protect
      internal/relay.SubmitOpenMLSArtifactEnvelope
      ArtifactKindApplicationMessage
      dev/pre-alpha warning

    openmls-inbox-dev:
      internal/app/openmls_runtime.go
      state.RequireReadyDevice
      Cypher inbox
      skip unsupported envelopes
      internal/relay.WriteOpenMLSArtifactFromEnvelope
      runOpenMLSMessageOpen
      sidecar message-open
      print plaintext after sidecar success
      optional ack only after sidecar success

### Candidate v0.4.12+ dev OpenMLS bootstrap wrappers

    openmls-identity-create-dev:
      sidecar identity-create

    openmls-identity-status-dev:
      sidecar identity-status

    openmls-bundle-export-dev:
      sidecar public-bundle-export

    openmls-conversation-create-dev:
      sidecar conversation-create

    openmls-conversation-load-check-dev:
      sidecar conversation-load-check

    openmls-conversation-add-member-dev:
      sidecar conversation-add-member

    openmls-conversation-join-dev:
      sidecar conversation-join

### v0.4.5 through v0.4.11 runtime smoke proof

    scripts/dev-openmls-runtime-smoke.sh:
      temp workspace
      temp Cypher binary
      temp Cypher DB
      dynamic loopback Cypher port
      temp Alice/Bob Comms state
      sidecar state reset
      direct sidecar identity/conversation bootstrap
      openmls-send-dev application-message send
      openmls-inbox-dev application-message open + ack
      inbox-empty-after-ack verification
      robust cleanup with INT -> TERM -> KILL escalation

### v0.4.9/v0.4.10 runner profile

    go run . --profile dev-runtime-openmls --clean-generated

    Runner.DevRuntimeOpenMLS:
      CheckRequiredPaths
      CheckLiveGitUmbrella
      toolchain reporting
      ArtifactScan("pre-dev-runtime-openmls")
      RunStep for carbonstack-comms/scripts/dev-openmls-runtime-smoke.sh
      ArtifactScan("post-dev-runtime-openmls")
      explicit nonclaim summary

    Runner.CheckLiveGitUmbrella:
      requires .git markers in carbonstack, carbonstack-comms, carbonstack-cypher
      refuses non-git package-like roots

### OpenMLS sidecar commands relevant to v0.4.x

    provider-info
    identity-create
    identity-status
    public-bundle-export
    conversation-create
    conversation-load-check
    conversation-add-member
    conversation-join
    message-protect
    message-open

### Relay/API functions relevant to runtime integration

    internal/client.CypherClient.SubmitEnvelope
    internal/client.CypherClient.Inbox
    internal/client.CypherClient.AckEnvelope
    internal/relay.SubmitOpenMLSArtifactEnvelope
    internal/relay.WriteOpenMLSArtifactFromEnvelope

### Existing v0.4.0 package validation

    cd <package-root>/carbonstack/tools/carbonstack-validate
    go run . --profile verify-checksums --root <package-root>
    go run . --profile full --root <package-root> --clean-generated

### Live umbrella validation

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go run . --profile dev-runtime-openmls --clean-generated
    go run . --profile local-cypher
    go run . --profile doctor
    go run . --profile core --clean-generated

### Envelope/protocol constants

Accepted opaque OpenMLS application-message pair:

    content_type=carbonstack.mls.application-message.v0
    protocol_version=carbonstack-openmls-sidecar-v0

Historical invalid pair validated by local-cypher:

    content_type=carbonstack.message.text.stub.v0
    protocol_version=carbonstack-openmls-sidecar-v0

Expected result:

    HTTP 400
    error code unsupported_protocol_version

---

## 10. v0.4.x Forward Plan

### v0.4.12 — dev-only OpenMLS bootstrap command contract

Expected:

    define exact flags.
    define exact output fields.
    define sidecar JSON parsing policy.
    define path-hint / absolute-path policy.
    define testing seams.
    define implementation order.
    define smoke-script migration conditions.

Preferred style:

    commands remain explicit openmls-*-dev.
    labels remain explicit sidecar labels for now.
    output remains stable key/value text.
    wrappers do not mutate Comms trust state.
    wrappers do not claim production identity UX.
    wrappers do not replace send/inbox.
    wrappers do not import Relay Space or local-backbone naming.

### v0.4.13+

Possible staged implementation:

    v0.4.13:
      openmls-identity-create-dev
      openmls-identity-status-dev

    v0.4.14:
      openmls-bundle-export-dev
      openmls-conversation-create-dev
      openmls-conversation-load-check-dev

    v0.4.15:
      openmls-conversation-add-member-dev
      openmls-conversation-join-dev

    v0.4.16:
      smoke-script migration recon / wrapper-based smoke variant decision

### v0.4.x later polish

Likely needs:

    exact sidecar setup runbook.
    clearer dev state reset semantics.
    better command output shape.
    decision on whether stub-era send/inbox are aliased, replaced, or deprecated.
    decision on when local-backbone naming stops being premature.

---

## 11. v0.5.x and Later Carry-Forward

### v0.5.x

Local storage, trust state, provider state, vault design, and PQ/hybrid ciphersuite migration readiness.

Do not start PQ/hybrid implementation before the runtime OpenMLS/Cypher path exists and is documented. PQ/hybrid migration touches provider state, ciphersuite metadata, trust display, artifact size, reinit/rekey/recovery policy, and claim boundaries.

### v0.6.x+

Hostile-server and abuse-resistance harnesses.

### v0.7.x+

Deployability and operations hardening.

### v0.8.x+

Documented self-pentest / adversarial validation.

### v0.9.x+

Claim-boundary review.

### v0.10.x+

Android backend/app work.

### v1.x.x

Public app/server major epoch and much later CarbonStackOS/appliance research, if justified.

---

## 12. Lean Breakpoint Summary

    CarbonStack v0.4.11 is the OpenMLS bootstrap wrapper recon/planning checkpoint. v0.4.0 remains the current public broad local deployability pre-release. carbonstack is now at 8ee3ebb docs: plan dev OpenMLS bootstrap wrappers. carbonstack-comms remains at 0a48ae1 refactor: split OpenMLS runtime command helpers. v0.4.11 adds docs/161-openmls-bootstrap-wrapper-recon-v0.md, updates docs/README.md, and updates roadmap/ROADMAP.md. The checkpoint records that direct sidecar bootstrap inside scripts/dev-openmls-runtime-smoke.sh should not be blindly replaced; future wrappers should use explicit openmls-*-dev names, likely live in carbonstack-comms/internal/app/openmls_bootstrap.go, preserve explicit sidecar labels, and get a command-contract doc before implementation. Validation passed for go test ./... in the runner, dev-runtime-openmls --clean-generated, local-cypher, doctor, and core --clean-generated, with final clean repo heads across all four repos. Next safe rung: v0.4.12 dev-only OpenMLS bootstrap command contract.

---

## 13. Preserved Immediate Previous Handoff: v0.4.10

The following is the previous v0.4.10 handoff. Where it conflicts with the v0.4.11 overlay above, v0.4.11 wins for current state.


# CarbonStack LogDoc v0.4.10

**Last updated:** 2026-06-05 local session  
**Owner/operator:** bitcrusher32  
**Assistant/model:** GPT-5.5 Thinking  
**Current phase:** Phase 2F / **v0.4.x dev-runtime-openmls profile boundary checkpoint complete**. The v0.4.0 public broad local deployability pre-release remains the current public release. v0.4.5 proved the dev runtime OpenMLS CLI path; v0.4.6 recorded that smoke proof in the main `carbonstack` repo; v0.4.7 assessed the proof and kept `local-backbone` reserved; v0.4.8 split OpenMLS runtime command glue into `carbonstack-comms/internal/app/openmls_runtime.go`; v0.4.9 added a **manual**, live-umbrella-only `dev-runtime-openmls` validation profile; v0.4.10 now records that profile boundary as stable: repeatable from the live umbrella, cleanable through `--clean-generated`, refusing non-git package-like roots as intended, and still excluded from `full`. `carbonstack` is now at `7384f43 (HEAD -> main, origin/main, origin/HEAD) docs: record dev OpenMLS runtime profile boundary`; `carbonstack-comms` remains at `0a48ae1 (HEAD -> main, origin/main, origin/HEAD) refactor: split OpenMLS runtime command helpers`; `carbonstack-cypher` remains at `9ab994c (HEAD -> main, origin/main, origin/HEAD) docs: point to local operator runbook`; `carbonstack-os` remains at `1bbbe52 (HEAD -> main, origin/main, origin/HEAD) docs: clarify CarbonStackOS target direction`.

**Version schema:** `v[scope].[timeline]`. This file is `v0.4.10`, the manual `dev-runtime-openmls` profile boundary / documentation-evidence checkpoint after the v0.4.9 runner-profile implementation. Per LogDoc V2 practice, this Markdown file preserves current operational continuity and carries the v0.4.x runtime-integration thread forward; the JSON breakpoint remains the lean machine-readable handoff state.

**Compression note:** The detailed v0.3.0 through v0.3.36PRIME process ledger remains archived in PRIME/sanitized PRIME. v0.4.0 compressed that pre-release sprint; v0.4.1 preserved runtime recon; v0.4.2 preserved the command-contract decision; v0.4.3 preserved the first dev-only OpenMLS send-side implementation; v0.4.4 preserved the first dev-only OpenMLS receive/open/optional-ack implementation; v0.4.5 preserved the first committed dev runtime smoke proof; v0.4.6 preserved project-level documentation/status alignment for that proof; v0.4.7 preserved the decision boundary before local-backbone/runner-profile promotion; v0.4.8 preserved behavior-preserving OpenMLS command-helper extraction; v0.4.9 preserved the manual `dev-runtime-openmls` runner-profile implementation; v0.4.10 now preserves the boundary evidence showing the profile behaves as intended.

---

## 1. Project Goal

**Active goal:** Continue moving CarbonStackComms from stub-era runtime `send`/`inbox` behavior toward explicit dev-only OpenMLS/Cypher-backed runtime commands and repeatable validation, while preserving claim discipline. v0.4.10 confirmed and documented the boundary of the manual `dev-runtime-openmls` profile instead of changing runner behavior.

CarbonStack’s broader project intent remains:

- **CarbonStack** is the secure-communications appliance-stack project and shared doctrine/release authority.
- **CarbonStackComms** is the text-first messaging client and current OpenMLS sidecar/relay integration surface.
- **CarbonStackCypher** is the self-hostable relay/storage stack.
- **CarbonStackOS** is future constrained Android-derived appliance OS work and remains out of current runnable validation packages.
- The `carbonstack` Gitea repo remains the source of truth for release state, release assets, docs, and public claims.
- GitHub mirrors exist as secondary push mirrors only; Gitea remains authoritative.

**North star:**

> Its goal is not to be a better smartphone. Its goal is to stop being a smartphone.

**Core doctrine:**

> Every feature is guilty until it proves it does not add unacceptable parser, network, sensor, identity, or filesystem authority.

**Current reality after v0.4.10:** the dev runtime OpenMLS application-message smoke proof is callable through the main `carbonstack` validation runner, and its boundary has now been checked and recorded.

Correct interpretation:

    openmls-send-dev is a dev-only OpenMLS send-side command.
    openmls-inbox-dev is a dev-only OpenMLS receive/open/optional-ack command.
    scripts/dev-openmls-runtime-smoke.sh proves the dev CLI application-message path:
      openmls-send-dev -> Cypher -> openmls-inbox-dev --ack
    tools/carbonstack-validate has a manual dev-runtime-openmls profile that wraps that smoke script.
    dev-runtime-openmls is repeatable from the live umbrella checkout.
    dev-runtime-openmls is live-git-umbrella-only for now.
    dev-runtime-openmls intentionally refuses non-git package-like roots.
    dev-runtime-openmls is not included in full.
    dev-runtime-openmls is not release-package validation yet.
    docs/159-dev-runtime-openmls-runner-profile-v0.md records the runner profile.
    docs/160-dev-runtime-openmls-profile-boundary-v0.md records the boundary check.
    internal/app/openmls_runtime.go remains the OpenMLS runtime command helper file.
    Existing send/inbox remain stub-era.
    local-backbone remains reserved.

Hard nonclaims remain: v0.4.10 does **not** prove production readiness, production E2EE, hostile-server safety, metadata privacy, Android readiness, CarbonStackOS readiness, secure vault/storage, public ingress safety, real homelab validation, audit, certification, PQ/hybrid security, mature messenger UX, release-package runtime validation, or local-backbone readiness.

---

## 2. Current Public Release State

Public release:

    https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.4.0

Release title:

    CarbonStack v0.4.0 Broad Local Deployability Pre-Release

Release tag:

    v0.4.0

Release commit:

    24ac9fc494 / 24ac9fc docs: record v0.4.0 package rehearsal

Gitea release status:

    Pre-Release

Release assets attached:

    carbonstack-v0.4.0-broad-local-deployability-pre-release.tgz
    carbonstack-v0.4.0-release-manifest.json
    carbonstack-v0.4.0-package-checksums.txt
    carbonstack-v0.4.0-asset-checksums.txt
    carbonstack-v0.4.0-validation-freeze.md
    v0.4.0-testing-runbook.md
    v0.4.0-release-notes.md
    LICENSE

Important release-page warning:

    The default Gitea Source Code ZIP/TAR.GZ downloads are auto-generated carbonstack-only archives.
    The intended runnable validation artifact is the attached multi-repo release package plus companion assets.

Post-release public surface:

    carbonstack README is evergreen.
    It points interested parties to the Releases page for runnable / known-good artifacts.
    It states Gitea is the source of truth.
    GitHub mirrors are secondary push mirrors, not release authority.

---

## 3. Current Repo Heads

    carbonstack        7384f43 (HEAD -> main, origin/main, origin/HEAD) docs: record dev OpenMLS runtime profile boundary
    carbonstack-comms  0a48ae1 (HEAD -> main, origin/main, origin/HEAD) refactor: split OpenMLS runtime command helpers
    carbonstack-cypher 9ab994c (HEAD -> main, origin/main, origin/HEAD) docs: point to local operator runbook
    carbonstack-os     1bbbe52 (HEAD -> main, origin/main, origin/HEAD) docs: clarify CarbonStackOS target direction

Current public release tag:

    v0.4.0 -> 24ac9fc494 / 24ac9fc docs: record v0.4.0 package rehearsal

Post-release public-surface cleanup:

    6ba7aab docs: make front readme evergreen

v0.4.1 runtime recon commit:

    83d2f0f docs: record runtime OpenMLS Comms recon

v0.4.2 command contract commit:

    a799764 docs: define runtime OpenMLS command contract

v0.4.3 send implementation commit:

    39fb287 feat: add dev OpenMLS send command

v0.4.4 inbox implementation commit:

    65bc707 feat: add dev OpenMLS inbox command

v0.4.5 smoke proof commit:

    8e6e8b4 test: add dev OpenMLS runtime smoke proof

v0.4.6 main-doc alignment commit:

    3cd77a8 docs: record dev OpenMLS runtime smoke proof

v0.4.7 assessment commit:

    dc2d16c docs: assess pre-local-backbone runtime proof

v0.4.8 extraction commit:

    0a48ae1 refactor: split OpenMLS runtime command helpers

v0.4.9 runner-profile commit:

    8eeadb2 feat: add dev OpenMLS runtime validation profile

Current v0.4.10 boundary-doc commit:

    7384f43 docs: record dev OpenMLS runtime profile boundary

Continuity note:

    `7384f43` is a `carbonstack` docs/evidence commit after the `v0.4.0` public release tag.
    It does not retag v0.4.0.
    The release remains attached to `24ac9fc494`.

---

## 4. Current Validated / Known State

Validated / known at the v0.4.10 checkpoint:

    [RELEASE] v0.4.0 Gitea release exists, is marked Pre-Release, and is attached to 24ac9fc494.
    [PACKAGE] v0.4.0 package includes carbonstack, carbonstack-comms, carbonstack-cypher, and release metadata.
    [PACKAGE] carbonstack-os is deliberately excluded from the runnable package.
    [RUNNER] full means release-snapshot followed by local-cypher.
    [RUNNER] release-snapshot validates package layout, metadata, strict pre-test artifact scan, package checksums, and core.
    [RUNNER] core validates doctor, OpenMLS real-Cypher lifecycle, Comms package tests, Cypher package tests, and artifact scans.
    [RUNNER] local-cypher validates Cypher-only lifecycle, restart/persistence against the same temporary DB, positive opaque OpenMLS envelope/ack flow, and historical negative stub-text/OpenMLS protocol pairing.
    [RUNNER] dev-runtime-openmls validates the dev/pre-alpha OpenMLS application-message runtime CLI smoke proof.
    [RUNNER] dev-runtime-openmls is manual-only and live-umbrella-only for now.
    [RUNNER] dev-runtime-openmls is not included in full.
    [RUNNER] dev-runtime-openmls checks live git markers for carbonstack, carbonstack-comms, and carbonstack-cypher.
    [RUNNER] dev-runtime-openmls refuses non-git package-like roots as intended.
    [RUNNER] dev-runtime-openmls wraps carbonstack-comms/scripts/dev-openmls-runtime-smoke.sh.
    [RUNNER] dev-runtime-openmls performs pre/post artifact scans and relies on --clean-generated for known OpenMLS generated roots after successful runs.
    [DOCS] docs/159-dev-runtime-openmls-runner-profile-v0.md records the v0.4.9 manual runner profile.
    [DOCS] docs/160-dev-runtime-openmls-profile-boundary-v0.md records the v0.4.10 boundary check.
    [DOCS] docs/README.md indexes docs/160-dev-runtime-openmls-profile-boundary-v0.md.
    [DOCS] roadmap/ROADMAP.md records v0.4.10 and recommends sidecar bootstrap wrapper recon next.
    [DOCS] tools/carbonstack-validate/README.md clarifies that dev-runtime-openmls refuses non-git package-like roots and remains separate from full.
    [COMMS] openmls-send-dev exists in carbonstack-comms.
    [COMMS] openmls-inbox-dev exists in carbonstack-comms.
    [COMMS] scripts/dev-openmls-runtime-smoke.sh exists in carbonstack-comms.
    [COMMS] internal/app/openmls_runtime.go contains OpenMLS runtime command helper code.
    [COMMS] internal/app/commands.go remains the command registry / non-OpenMLS command surface.
    [COMMS] existing send/inbox/ack remain in place and stub-era.
    [MIRRORS] all four repos have GitHub push-mirror functionality, with Gitea remaining source of truth.

v0.4.10 scout and validation result:

    go test ./... -count=1 passed for carbonstack/tools/carbonstack-validate.
    go run . --profile dev-runtime-openmls passed without --clean-generated.
    The uncleaned dev-runtime-openmls run left only known OpenMLS sidecar generated roots:
      .carbonstack-openmls-sidecar-state
      target
    repo status remained clean after the uncleaned run because generated roots are ignored/known.
    core --clean-generated cleaned known generated roots after the uncleaned run.
    go run . --profile dev-runtime-openmls --clean-generated passed.
    The cleaned dev-runtime-openmls run removed known OpenMLS sidecar generated roots after success.
    non-git package-like root refusal was tested and passed:
      dev-runtime-openmls refused missing .git markers for carbonstack, carbonstack-comms, and carbonstack-cypher.
    full dispatch was inspected and still does not call dev-runtime-openmls.
    Running full from the live umbrella failed at release-snapshot metadata because live umbrella is not a release package root.
    That full failure was expected and is not a dev-runtime-openmls failure.
    local-cypher passed.
    doctor passed.
    core --clean-generated passed.
    final repo status showed carbonstack at 7384f43 and all four repos clean.

Expected live umbrella validation commands now include:

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go run . --profile dev-runtime-openmls --clean-generated
    go run . --profile local-cypher
    go run . --profile doctor
    go run . --profile core --clean-generated

Expected v0.4.0 public package validation commands remain unchanged:

    cd <package-root>/carbonstack/tools/carbonstack-validate
    go run . --profile verify-checksums --root <package-root>
    go run . --profile full --root <package-root> --clean-generated

Current direct smoke command remains:

    cd ~/repos/carbonstack_umbrella/carbonstack-comms
    scripts/dev-openmls-runtime-smoke.sh

---

## 5. v0.4.10 Work Completed

### 5.1 Boundary scout confirmed the profile behaves as intended

The v0.4.10 scout inspected:

    current dev-runtime-openmls runner implementation.
    runner dispatch around dev-runtime-openmls and full.
    runner README profile documentation.
    docs/159-dev-runtime-openmls-runner-profile-v0.md.
    docs index and roadmap references.
    live dev-runtime-openmls repeatability.
    behavior with and without --clean-generated.
    generated-root cleanup behavior.
    full profile separation.
    non-git package-like root refusal.
    final baseline validation.

Key findings:

    dev-runtime-openmls is repeatable from the live umbrella checkout.
    Running without --clean-generated leaves only known OpenMLS sidecar generated roots.
    --clean-generated cleans the known generated roots after a successful run.
    full remains release-snapshot -> local-cypher and does not call dev-runtime-openmls.
    non-git package-like roots are refused correctly through live .git marker checks.
    No runner code change was necessary.

### 5.2 docs/160-dev-runtime-openmls-profile-boundary-v0.md was added

New file:

    carbonstack/docs/160-dev-runtime-openmls-profile-boundary-v0.md

It records:

    v0.4.10 is a boundary/polish documentation checkpoint.
    dev-runtime-openmls is repeatable from the live umbrella.
    dev-runtime-openmls is manual-only.
    dev-runtime-openmls is live-git-umbrella-only.
    dev-runtime-openmls is not included in full.
    dev-runtime-openmls is not release-package validation.
    dev-runtime-openmls is not local-backbone.
    dev-runtime-openmls is not production messaging UX.
    uncleaned runs leave only known OpenMLS sidecar generated roots.
    non-git package-like root refusal works as intended.
    next recommended rung is sidecar bootstrap wrapper recon.

### 5.3 docs/README.md was updated

Updated file:

    carbonstack/docs/README.md

Added index entry:

    docs/160-dev-runtime-openmls-profile-boundary-v0.md

Meaning:

    The docs archive now indexes the v0.4.10 boundary check for the manual dev-runtime-openmls profile.

### 5.4 roadmap/ROADMAP.md was updated

Updated file:

    carbonstack/roadmap/ROADMAP.md

Added section:

    v0.4.10 — dev-runtime-openmls profile boundary check

It records:

    live umbrella repeatability.
    known generated-root behavior.
    --clean-generated cleanup behavior.
    non-git package-like root refusal.
    full remains separate and does not call dev-runtime-openmls.
    next recommended rung is v0.4.11 sidecar bootstrap wrapper recon.

### 5.5 tools/carbonstack-validate/README.md was polished

Updated file:

    carbonstack/tools/carbonstack-validate/README.md

Added clarification:

    dev-runtime-openmls intentionally refuses non-git package-like roots for now.
    use it from the live umbrella checkout, not release package roots.
    full remains the release-package validation ladder and does not include dev-runtime-openmls.

### 5.6 Validation passed

Commands observed passing or expected-failing:

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go test ./... -count=1
    go run . --profile dev-runtime-openmls
    go run . --profile doctor
    go run . --profile core --clean-generated
    go run . --profile dev-runtime-openmls --clean-generated
    go run . --profile full --clean-generated  # expected failure from live umbrella due missing release metadata
    go run . --profile dev-runtime-openmls --root <non-git-package-like-root>  # expected failure, passed refusal check
    go run . --profile local-cypher
    go run . --profile doctor
    go run . --profile core --clean-generated

Final baseline:

    go test ./... -count=1 passed in the runner.
    dev-runtime-openmls --clean-generated passed.
    local-cypher passed.
    doctor passed.
    core --clean-generated passed.
    final repo status was clean across all four repos.

Final commit/push:

    carbonstack:
    7384f43 docs: record dev OpenMLS runtime profile boundary

Final repo snapshot:

    carbonstack        7384f43 docs: record dev OpenMLS runtime profile boundary
    carbonstack-comms  0a48ae1 refactor: split OpenMLS runtime command helpers
    carbonstack-cypher 9ab994c docs: point to local operator runbook
    carbonstack-os     1bbbe52 docs: clarify CarbonStackOS target direction

---

## 6. Critical Blunders / Lessons to Carry Forward

### 6.1 full failed from live umbrella for the expected reason

During boundary scouting, `go run . --profile full --clean-generated` was run from the live umbrella.

It failed during `release-snapshot` metadata checks because the live umbrella root did not contain the required `release/` metadata directory.

This is expected.

Correct interpretation:

    full is a release-package validation ladder.
    full should be run from a fresh extracted or throwaway staged release package root.
    dev-runtime-openmls is a live-umbrella manual profile.
    The full failure was not a dev-runtime-openmls failure.
    The full dispatch still does not call dev-runtime-openmls.

Lesson:

    Do not use full from the live umbrella as proof that the runner is broken.
    Use full from package roots.
    Use dev-runtime-openmls from live git umbrella roots.

### 6.2 Non-git package-like refusal worked

The scout copied carbonstack/carbonstack-comms/carbonstack-cypher into a temp package-like root and removed `.git` directories.

Expected result:

    dev-runtime-openmls should fail.

Observed result:

    dev-runtime-openmls reported missing live checkout markers for all three repos.
    it exited nonzero.
    scout printed PASS for the refusal check.

Lesson:

    The live-git-umbrella guard is working as intended.
    Do not describe dev-runtime-openmls as a release-package profile yet.

### 6.3 Known generated roots remain expected after smoke/profile runs

`dev-runtime-openmls` runs the smoke proof, which uses OpenMLS sidecar and Rust build paths. Known generated roots may appear:

    carbonstack-comms/internal/protocol/mls/openmls-sidecar/.carbonstack-openmls-sidecar-state
    carbonstack-comms/internal/protocol/mls/openmls-sidecar/target

This is expected as long as hits remain limited to known generated roots and `--clean-generated` removes them after success.

### 6.4 Commit metadata warning remains a polish issue

The commit log again showed auto-generated local commit metadata:

    Committer: bitcrusher32 <▓▓>

This is not a functional validation failure, but it remains a professionalism/publishing hygiene issue to address eventually with `git config --global user.name` and `git config --global user.email` if desired.

### 6.5 Manual profile boundary remains strict

v0.4.10 confirms that v0.4.9's manual profile is behaving properly, but the local-backbone boundary still holds:

    direct sidecar bootstrap remains.
    sidecar identity/KeyPackage/Welcome/conversation setup is not Comms runtime UX.
    send/inbox remain stub-era.
    profile is live-git-umbrella-only.
    profile is not included in full.
    profile is not release-package validation.

Correct phrase:

    manual dev-runtime-openmls validation profile

Still reserved:

    local-backbone

---

## 7. Current Blockers / Not Validated

Current blockers and nonvalidated areas after v0.4.10:

    Existing send/inbox remain stub-era.
    Runtime Comms has dev-only OpenMLS commands, a smoke proof, and a manual runner profile, but no mature user-facing UX.
    Sidecar KeyPackage/Welcome/bootstrap setup is still direct dev setup, not Comms runtime UX.
    No Relay Space join model exists yet.
    local-backbone runner profile does not exist and remains reserved.
    dev-runtime-openmls exists but is manual-only.
    dev-runtime-openmls is not included in full.
    dev-runtime-openmls is live-git-umbrella-only.
    dev-runtime-openmls is not part of public release package validation.
    Fresh release-package behavior for dev-runtime-openmls has not been tested as a supported validation surface.
    OpenMLS runtime helpers are split out, but no new bootstrap wrappers exist yet.
    Post-release public download verification is still useful if not already performed.
    Sanitized PRIME placement in carbonstack/logdoc-loop-system still needs a policy decision.
    Production-safe provider storage and secure vault design remain v0.5.x future work.
    PQ/hybrid ciphersuite migration readiness remains v0.5.x future work.
    Hostile-server rollback/replay/metadata-abuse validation remains future work.
    Android/CarbonStackOS implementation remains later work.
    CarbonStack Relay Space mechanics remain a future concept, not an implementation.
    Public ingress, systemd, cloudflared, real homelab deployment, and production service paths remain out of scope.

Hard nonclaims remain:

    Do not claim production readiness.
    Do not claim production E2EE.
    Do not claim hostile-server safety.
    Do not claim metadata privacy.
    Do not claim Debian production deployability.
    Do not claim systemd readiness.
    Do not claim cloudflared readiness.
    Do not claim Android readiness.
    Do not claim CarbonStackOS readiness.
    Do not claim secure vault/storage.
    Do not claim real homelab validation.
    Do not claim external audit or certification.
    Do not claim CarbonStack Relay Space mechanics are implemented.
    Do not claim local-backbone exists or is justified yet.
    Do not claim local-cypher validates runtime Comms UX.
    Do not claim full validates local-backbone.
    Do not claim full is a deployment command.
    Do not claim dev-runtime-openmls is release-package validation yet.
    Do not claim send/inbox are OpenMLS-backed yet.
    Do not claim PQ/hybrid security or quantum-safe messaging.
    Do not let GitHub mirrors supersede Gitea source-of-truth status.

Allowed claim:

    CarbonStackComms now has dev-only openmls-send-dev and openmls-inbox-dev commands.
    CarbonStackComms now has a dev runtime smoke script proving openmls-send-dev -> Cypher -> openmls-inbox-dev --ack.
    CarbonStackComms now has OpenMLS runtime command helper glue split into internal/app/openmls_runtime.go.
    The main carbonstack repo records the smoke proof, the pre-local-backbone assessment, the manual dev-runtime-openmls runner profile, and the profile boundary check.
    carbonstack/tools/carbonstack-validate has a manual dev-runtime-openmls profile.
    dev-runtime-openmls wraps the comms smoke script and validates the dev/pre-alpha OpenMLS application-message CLI path.
    dev-runtime-openmls is live-git-umbrella-only and not included in full.
    dev-runtime-openmls intentionally refuses non-git package-like roots.
    The smoke script starts a temporary local Cypher server, uses temporary Comms state files, creates dev-local OpenMLS sidecar identities/conversation, sends an OpenMLS application message through the Comms CLI, opens it through the Comms CLI, acks after sidecar success, and verifies Bob's inbox is empty after ack.
    Existing send/inbox remain stub-era.
    v0.4.0 remains the current public broad local deployability pre-release.
    Gitea remains source of truth; GitHub mirrors are secondary push mirrors.

---

## 8. Next Safest Actions

Recommended next rung:

    v0.4.11 sidecar bootstrap wrapper recon

Focus:

    inspect whether Comms should add dev-only wrappers for direct sidecar setup steps currently embedded in scripts/dev-openmls-runtime-smoke.sh.
    identify exact sidecar CLI flags and JSON output shapes for bootstrap commands.
    decide whether wrappers should be separate dev commands or kept script-only.
    keep all wrapper naming dev-only.
    avoid replacing send/inbox yet.
    avoid local-backbone naming.
    avoid putting dev-runtime-openmls into full.
    keep v0.5.x state/trust/vault/PQ work deferred.

Candidate wrapper areas:

    OpenMLS identity create/status
    public bundle export
    conversation create
    conversation add member / Welcome generation
    conversation join
    conversation load-check

Alternative next rung:

    v0.4.11 command migration/deprecation planning for stub-era send/inbox

Preferred order:

    sidecar bootstrap wrapper recon first.
    then decide whether wrappers are worth implementing before changing send/inbox.

Avoid next:

    adding dev-runtime-openmls to full immediately.
    replacing send/inbox immediately.
    adding local-backbone prematurely.
    adding a runner profile named local-backbone.
    moving dev-runtime-openmls into public release validation without fresh-root behavior checks.
    public ingress.
    cloudflared.
    systemd.
    real homelab deployment.
    Android app.
    CarbonStackOS implementation.
    CarbonStack Relay Space implementation.
    PQ/hybrid ciphersuite implementation.
    production/security claims.

---

## 9. Critical Paths / Functions

### Source-of-truth public surfaces

    [GITEA SOURCE OF TRUTH] https://git.bitcrusher32.win/bitcrusher32/carbonstack
    [RELEASES] https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases
    [CURRENT PUBLIC RELEASE] https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.4.0

GitHub mirrors:

    secondary push mirrors for discoverability/redundancy only.
    Do not treat mirrors as release authority unless project policy changes.

### WSL Debian active working umbrella

    [WSL-DEBIAN] umbrella: ~/repos/carbonstack_umbrella
    [WSL-DEBIAN] carbonstack: ~/repos/carbonstack_umbrella/carbonstack
    [WSL-DEBIAN] carbonstack-comms: ~/repos/carbonstack_umbrella/carbonstack-comms
    [WSL-DEBIAN] carbonstack-cypher: ~/repos/carbonstack_umbrella/carbonstack-cypher
    [WSL-DEBIAN] carbonstack-os: ~/repos/carbonstack_umbrella/carbonstack-os
    [GO-RUNNER] runner dir: ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate

### Runtime integration doc surfaces

    [DOC] v0.4.1 runtime recon:
    carbonstack/docs/155-runtime-comms-openmls-cypher-recon-v0.md

    [DOC] v0.4.2 command contract:
    carbonstack/docs/156-runtime-openmls-command-contract-v0.md

    [DOC] v0.4.6 smoke proof record:
    carbonstack/docs/157-dev-runtime-openmls-smoke-proof-v0.md

    [DOC] v0.4.7 pre-local-backbone assessment:
    carbonstack/docs/158-pre-local-backbone-assessment-v0.md

    [DOC] v0.4.9 runner profile:
    carbonstack/docs/159-dev-runtime-openmls-runner-profile-v0.md

    [DOC] v0.4.10 profile boundary:
    carbonstack/docs/160-dev-runtime-openmls-profile-boundary-v0.md

    [DOC] docs index:
    carbonstack/docs/README.md

    [DOC] roadmap:
    carbonstack/roadmap/ROADMAP.md

    [DOC] top README:
    carbonstack/README.md

### Runner surfaces

    [RUNNER] root:
    carbonstack/tools/carbonstack-validate

    [RUNNER] dispatch:
    carbonstack/tools/carbonstack-validate/main.go

    [RUNNER] dev-runtime-openmls implementation:
    carbonstack/tools/carbonstack-validate/dev_runtime_openmls.go

    [RUNNER] runner README:
    carbonstack/tools/carbonstack-validate/README.md

    [RUNNER] local-cypher:
    carbonstack/tools/carbonstack-validate/local_cypher.go

    [RUNNER] release-snapshot:
    carbonstack/tools/carbonstack-validate/release_snapshot.go

    [RUNNER] checksums:
    carbonstack/tools/carbonstack-validate/checksums.go

### Comms runtime surfaces

    [COMMS] CLI entry:
    carbonstack-comms/cmd/comms/main.go

    [COMMS] command registry and non-OpenMLS commands:
    carbonstack-comms/internal/app/commands.go

    [COMMS] OpenMLS runtime command helpers:
    carbonstack-comms/internal/app/openmls_runtime.go

    [COMMS] openmls-send-dev tests:
    carbonstack-comms/internal/app/openmls_send_dev_test.go

    [COMMS] openmls-inbox-dev tests:
    carbonstack-comms/internal/app/openmls_inbox_dev_test.go

    [COMMS] dev runtime smoke proof:
    carbonstack-comms/scripts/dev-openmls-runtime-smoke.sh

    [COMMS] Cypher client:
    carbonstack-comms/internal/client/cypher.go

    [COMMS] relay helpers:
    carbonstack-comms/internal/relay/cypher_bridge.go
    carbonstack-comms/internal/relay/openmls_artifacts.go

    [COMMS] OpenMLS sidecar dir:
    carbonstack-comms/internal/protocol/mls/openmls-sidecar

    [COMMS] sidecar README:
    carbonstack-comms/internal/protocol/mls/openmls-sidecar/README.md

### Current stub-era Comms functions / behaviors

    send:
      trust policy evaluation
      mock/stub provider encryption
      Cypher envelope submit

    inbox:
      Cypher inbox retrieval
      mock/stub provider plaintext display as stub_plaintext

    ack:
      envelope ack by envelope ID and recipient device ID

### Implemented v0.4.x dev OpenMLS commands

    openmls-send-dev:
      internal/app/openmls_runtime.go
      state.RequireReadyDevice
      trust.EvaluateSend
      runOpenMLSMessageProtect
      sidecar message-protect
      internal/relay.SubmitOpenMLSArtifactEnvelope
      ArtifactKindApplicationMessage
      dev/pre-alpha warning

    openmls-inbox-dev:
      internal/app/openmls_runtime.go
      state.RequireReadyDevice
      Cypher inbox
      skip unsupported envelopes
      internal/relay.WriteOpenMLSArtifactFromEnvelope
      runOpenMLSMessageOpen
      sidecar message-open
      print plaintext after sidecar success
      optional ack only after sidecar success

### v0.4.5/v0.4.8/v0.4.9/v0.4.10 runtime smoke proof

    scripts/dev-openmls-runtime-smoke.sh:
      temp workspace
      temp Cypher binary
      temp Cypher DB
      dynamic loopback Cypher port
      temp Alice/Bob Comms state
      sidecar state reset
      sidecar identity/conversation bootstrap
      openmls-send-dev application-message send
      openmls-inbox-dev application-message open + ack
      inbox-empty-after-ack verification
      robust cleanup with INT -> TERM -> KILL escalation

### v0.4.9/v0.4.10 runner profile

    go run . --profile dev-runtime-openmls --clean-generated

    Runner.DevRuntimeOpenMLS:
      CheckRequiredPaths
      CheckLiveGitUmbrella
      toolchain reporting
      ArtifactScan("pre-dev-runtime-openmls")
      RunStep for carbonstack-comms/scripts/dev-openmls-runtime-smoke.sh
      ArtifactScan("post-dev-runtime-openmls")
      explicit nonclaim summary

    Runner.CheckLiveGitUmbrella:
      requires .git markers in carbonstack, carbonstack-comms, carbonstack-cypher
      refuses non-git package-like roots

### OpenMLS sidecar commands relevant to v0.4.x

    provider-info
    identity-create
    identity-status
    public-bundle-export
    conversation-create
    conversation-load-check
    conversation-add-member
    conversation-join
    message-protect
    message-open

### Relay/API functions relevant to runtime integration

    internal/client.CypherClient.SubmitEnvelope
    internal/client.CypherClient.Inbox
    internal/client.CypherClient.AckEnvelope
    internal/relay.SubmitOpenMLSArtifactEnvelope
    internal/relay.WriteOpenMLSArtifactFromEnvelope

### Existing v0.4.0 package validation

    cd <package-root>/carbonstack/tools/carbonstack-validate
    go run . --profile verify-checksums --root <package-root>
    go run . --profile full --root <package-root> --clean-generated

### Live umbrella validation

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go run . --profile dev-runtime-openmls --clean-generated
    go run . --profile local-cypher
    go run . --profile doctor
    go run . --profile core --clean-generated

### Envelope/protocol constants

Accepted opaque OpenMLS application-message pair:

    content_type=carbonstack.mls.application-message.v0
    protocol_version=carbonstack-openmls-sidecar-v0

Historical invalid pair validated by local-cypher:

    content_type=carbonstack.message.text.stub.v0
    protocol_version=carbonstack-openmls-sidecar-v0

Expected result:

    HTTP 400
    error code unsupported_protocol_version

---

## 10. v0.4.x Forward Plan

### v0.4.11 — sidecar bootstrap wrapper recon

Expected:

    inspect direct sidecar setup steps currently used by the smoke script.
    decide whether Comms should add dev-only wrapper commands for those steps.
    preserve strict dev-only naming.
    preserve claim discipline.
    avoid calling wrappers Relay Space join UX.
    avoid replacing send/inbox yet.

Candidate dev-only wrapper commands / functions:

    identity-create / identity-status
    public-bundle-export
    conversation-create
    conversation-add-member
    conversation-join
    conversation-load-check

### v0.4.12+ — possible bootstrap wrapper implementation or command migration planning

Possible paths:

    add dev-only wrapper commands around sidecar bootstrap flows.
    improve sidecar setup/reset ergonomics.
    document command migration/deprecation options for stub-era send/inbox.
    test dev-runtime-openmls from cleaner live checkouts.
    later consider release-root behavior, but do not add to full yet.

### v0.4.x later polish

Likely needs:

    exact sidecar setup runbook.
    clearer dev state reset semantics.
    better command output shape.
    decision on whether stub-era send/inbox are aliased, replaced, or deprecated.
    decision on when local-backbone naming stops being premature.

---

## 11. v0.5.x and Later Carry-Forward

### v0.5.x

Local storage, trust state, provider state, vault design, and PQ/hybrid ciphersuite migration readiness.

Do not start PQ/hybrid implementation before the runtime OpenMLS/Cypher path exists and is documented. PQ/hybrid migration touches provider state, ciphersuite metadata, trust display, artifact size, reinit/rekey/recovery policy, and claim boundaries.

### v0.6.x+

Hostile-server and abuse-resistance harnesses.

### v0.7.x+

Deployability and operations hardening.

### v0.8.x+

Documented self-pentest / adversarial validation.

### v0.9.x+

Claim-boundary review.

### v0.10.x+

Android backend/app work.

### v1.x.x

Public app/server major epoch and much later CarbonStackOS/appliance research, if justified.

---

## 12. Lean Breakpoint Summary

    CarbonStack v0.4.10 is the dev-runtime-openmls profile boundary checkpoint. v0.4.0 remains the current public broad local deployability pre-release. carbonstack is now at 7384f43 docs: record dev OpenMLS runtime profile boundary. carbonstack-comms remains at 0a48ae1 refactor: split OpenMLS runtime command helpers. v0.4.10 adds docs/160-dev-runtime-openmls-profile-boundary-v0.md, updates docs/README.md, roadmap/ROADMAP.md, and tools/carbonstack-validate/README.md. The manual dev-runtime-openmls profile was repeatability checked from live umbrella, left only known generated roots without cleanup, cleaned known roots with --clean-generated, refused non-git package-like roots, and remained outside full. Validation passed for go test ./... in the runner, dev-runtime-openmls --clean-generated, local-cypher, doctor, and core --clean-generated, with final clean repo heads across all four repos. Next safe rung: v0.4.11 sidecar bootstrap wrapper recon.

---

## 13. Preserved Immediate Previous Handoff: v0.4.9

The following is the previous v0.4.9 handoff. Where it conflicts with the v0.4.10 overlay above, v0.4.10 wins for current state.


# CarbonStack LogDoc v0.4.9

**Last updated:** 2026-06-05 local session  
**Owner/operator:** bitcrusher32  
**Assistant/model:** GPT-5.5 Thinking  
**Current phase:** Phase 2F / **v0.4.x manual dev-runtime-openmls validation profile checkpoint complete**. The v0.4.0 public broad local deployability pre-release remains the current public release. v0.4.5 proved the dev runtime OpenMLS CLI path; v0.4.6 recorded that smoke proof in the main `carbonstack` repo; v0.4.7 assessed the proof and kept `local-backbone` reserved; v0.4.8 split OpenMLS runtime command glue into `carbonstack-comms/internal/app/openmls_runtime.go`; v0.4.9 now promotes the existing comms-local smoke proof into a **manual**, live-umbrella-only `carbonstack/tools/carbonstack-validate` profile named `dev-runtime-openmls`. `carbonstack` is now at `8eeadb2 (HEAD -> main, origin/main, origin/HEAD) feat: add dev OpenMLS runtime validation profile`; `carbonstack-comms` remains at `0a48ae1 (HEAD -> main, origin/main, origin/HEAD) refactor: split OpenMLS runtime command helpers`; `carbonstack-cypher` remains at `9ab994c (HEAD -> main, origin/main, origin/HEAD) docs: point to local operator runbook`; `carbonstack-os` remains at `1bbbe52 (HEAD -> main, origin/main, origin/HEAD) docs: clarify CarbonStackOS target direction`.

**Version schema:** `v[scope].[timeline]`. This file is `v0.4.9`, the manual `dev-runtime-openmls` runner-profile checkpoint after the v0.4.8 behavior-preserving command/helper extraction. Per LogDoc V2 practice, this Markdown file preserves current operational continuity and carries the v0.4.x runtime-integration thread forward; the JSON breakpoint remains the lean machine-readable handoff state.

**Compression note:** The detailed v0.3.0 through v0.3.36PRIME process ledger remains archived in PRIME/sanitized PRIME. v0.4.0 compressed that pre-release sprint; v0.4.1 preserved runtime recon; v0.4.2 preserved the command-contract decision; v0.4.3 preserved the first dev-only OpenMLS send-side implementation; v0.4.4 preserved the first dev-only OpenMLS receive/open/optional-ack implementation; v0.4.5 preserved the first committed dev runtime smoke proof; v0.4.6 preserved project-level documentation/status alignment for that proof; v0.4.7 preserved the decision boundary before any local-backbone/runner-profile promotion; v0.4.8 preserved behavior-preserving OpenMLS command-helper extraction; v0.4.9 now preserves the manual `dev-runtime-openmls` runner-profile implementation and its claim boundaries.

---

## 1. Project Goal

**Active goal:** Continue moving CarbonStackComms from stub-era runtime `send`/`inbox` behavior toward explicit dev-only OpenMLS/Cypher-backed runtime commands and repeatable validation, while preserving claim discipline. v0.4.9 completed the first main `carbonstack` validation-runner profile for the dev runtime OpenMLS CLI smoke proof.

CarbonStack’s broader project intent remains:

- **CarbonStack** is the secure-communications appliance-stack project and shared doctrine/release authority.
- **CarbonStackComms** is the text-first messaging client and current OpenMLS sidecar/relay integration surface.
- **CarbonStackCypher** is the self-hostable relay/storage stack.
- **CarbonStackOS** is future constrained Android-derived appliance OS work and remains out of current runnable validation packages.
- The `carbonstack` Gitea repo remains the source of truth for release state, release assets, docs, and public claims.
- GitHub mirrors exist as secondary push mirrors only; Gitea remains authoritative.

**North star:**

> Its goal is not to be a better smartphone. Its goal is to stop being a smartphone.

**Core doctrine:**

> Every feature is guilty until it proves it does not add unacceptable parser, network, sensor, identity, or filesystem authority.

**Current reality after v0.4.9:** the dev runtime OpenMLS application-message smoke proof is still dev/pre-alpha, but it is now callable through a main project runner profile:

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go run . --profile dev-runtime-openmls --clean-generated

Correct interpretation:

    openmls-send-dev is a dev-only OpenMLS send-side command.
    openmls-inbox-dev is a dev-only OpenMLS receive/open/optional-ack command.
    scripts/dev-openmls-runtime-smoke.sh proves the dev CLI application-message path:
      openmls-send-dev -> Cypher -> openmls-inbox-dev --ack
    tools/carbonstack-validate now has a manual dev-runtime-openmls profile that wraps that smoke script.
    dev-runtime-openmls is live-umbrella-only for now.
    dev-runtime-openmls is not included in full.
    dev-runtime-openmls is not release-package validation yet.
    docs/159-dev-runtime-openmls-runner-profile-v0.md records the runner profile.
    internal/app/openmls_runtime.go remains the OpenMLS runtime command helper file.
    Existing send/inbox remain stub-era.
    local-backbone remains reserved.

Hard nonclaims remain: v0.4.9 does **not** prove production readiness, production E2EE, hostile-server safety, metadata privacy, Android readiness, CarbonStackOS readiness, secure vault/storage, public ingress safety, real homelab validation, audit, certification, PQ/hybrid security, mature messenger UX, release-package runtime validation, or local-backbone readiness.

---

## 2. Current Public Release State

Public release:

    https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.4.0

Release title:

    CarbonStack v0.4.0 Broad Local Deployability Pre-Release

Release tag:

    v0.4.0

Release commit:

    24ac9fc494 / 24ac9fc docs: record v0.4.0 package rehearsal

Gitea release status:

    Pre-Release

Release assets attached:

    carbonstack-v0.4.0-broad-local-deployability-pre-release.tgz
    carbonstack-v0.4.0-release-manifest.json
    carbonstack-v0.4.0-package-checksums.txt
    carbonstack-v0.4.0-asset-checksums.txt
    carbonstack-v0.4.0-validation-freeze.md
    v0.4.0-testing-runbook.md
    v0.4.0-release-notes.md
    LICENSE

Important release-page warning:

    The default Gitea Source Code ZIP/TAR.GZ downloads are auto-generated carbonstack-only archives.
    The intended runnable validation artifact is the attached multi-repo release package plus companion assets.

Post-release public surface:

    carbonstack README is evergreen.
    It points interested parties to the Releases page for runnable / known-good artifacts.
    It states Gitea is the source of truth.
    GitHub mirrors are secondary push mirrors, not release authority.

---

## 3. Current Repo Heads

    carbonstack        8eeadb2 (HEAD -> main, origin/main, origin/HEAD) feat: add dev OpenMLS runtime validation profile
    carbonstack-comms  0a48ae1 (HEAD -> main, origin/main, origin/HEAD) refactor: split OpenMLS runtime command helpers
    carbonstack-cypher 9ab994c (HEAD -> main, origin/main, origin/HEAD) docs: point to local operator runbook
    carbonstack-os     1bbbe52 (HEAD -> main, origin/main, origin/HEAD) docs: clarify CarbonStackOS target direction

Current public release tag:

    v0.4.0 -> 24ac9fc494 / 24ac9fc docs: record v0.4.0 package rehearsal

Post-release public-surface cleanup:

    6ba7aab docs: make front readme evergreen

v0.4.1 runtime recon commit:

    83d2f0f docs: record runtime OpenMLS Comms recon

v0.4.2 command contract commit:

    a799764 docs: define runtime OpenMLS command contract

v0.4.3 send implementation commit:

    39fb287 feat: add dev OpenMLS send command

v0.4.4 inbox implementation commit:

    65bc707 feat: add dev OpenMLS inbox command

v0.4.5 smoke proof commit:

    8e6e8b4 test: add dev OpenMLS runtime smoke proof

v0.4.6 main-doc alignment commit:

    3cd77a8 docs: record dev OpenMLS runtime smoke proof

v0.4.7 assessment commit:

    dc2d16c docs: assess pre-local-backbone runtime proof

v0.4.8 extraction commit:

    0a48ae1 refactor: split OpenMLS runtime command helpers

Current v0.4.9 runner-profile commit:

    8eeadb2 feat: add dev OpenMLS runtime validation profile

Continuity note:

    `8eeadb2` is a `carbonstack` runner/docs commit after the `v0.4.0` public release tag.
    It does not retag v0.4.0.
    The release remains attached to `24ac9fc494`.

---

## 4. Current Validated / Known State

Validated / known at the v0.4.9 checkpoint:

    [RELEASE] v0.4.0 Gitea release exists, is marked Pre-Release, and is attached to 24ac9fc494.
    [PACKAGE] v0.4.0 package includes carbonstack, carbonstack-comms, carbonstack-cypher, and release metadata.
    [PACKAGE] carbonstack-os is deliberately excluded from the runnable package.
    [RUNNER] full means release-snapshot followed by local-cypher.
    [RUNNER] release-snapshot validates package layout, metadata, strict pre-test artifact scan, package checksums, and core.
    [RUNNER] core validates doctor, OpenMLS real-Cypher lifecycle, Comms package tests, Cypher package tests, and artifact scans.
    [RUNNER] local-cypher validates Cypher-only lifecycle, restart/persistence against the same temporary DB, positive opaque OpenMLS envelope/ack flow, and historical negative stub-text/OpenMLS protocol pairing.
    [RUNNER] dev-runtime-openmls now validates the dev/pre-alpha OpenMLS application-message runtime CLI smoke proof.
    [RUNNER] dev-runtime-openmls is manual-only and live-umbrella-only for now.
    [RUNNER] dev-runtime-openmls is not included in full.
    [RUNNER] dev-runtime-openmls checks live git markers for carbonstack, carbonstack-comms, and carbonstack-cypher.
    [RUNNER] dev-runtime-openmls wraps carbonstack-comms/scripts/dev-openmls-runtime-smoke.sh.
    [RUNNER] dev-runtime-openmls performs pre/post artifact scans and relies on --clean-generated for known OpenMLS generated roots after successful runs.
    [DOCS] docs/159-dev-runtime-openmls-runner-profile-v0.md records the v0.4.9 manual runner profile.
    [DOCS] docs/README.md indexes docs/159-dev-runtime-openmls-runner-profile-v0.md.
    [DOCS] roadmap/ROADMAP.md records v0.4.9 and the boundary that this is not full/release-package/local-backbone validation.
    [DOCS] tools/carbonstack-validate/README.md documents the dev-runtime-openmls profile.
    [COMMS] openmls-send-dev exists in carbonstack-comms.
    [COMMS] openmls-inbox-dev exists in carbonstack-comms.
    [COMMS] scripts/dev-openmls-runtime-smoke.sh exists in carbonstack-comms.
    [COMMS] internal/app/openmls_runtime.go contains OpenMLS runtime command helper code.
    [COMMS] internal/app/commands.go remains the command registry / non-OpenMLS command surface.
    [COMMS] existing send/inbox/ack remain in place and stub-era.
    [MIRRORS] all four repos have GitHub push-mirror functionality, with Gitea remaining source of truth.

v0.4.9 validation result from the final log:

    go test ./... -count=1 passed for carbonstack/tools/carbonstack-validate.
    go run . --profile dev-runtime-openmls --clean-generated passed.
    dev-runtime-openmls printed the explicit boundary:
      not local-backbone, not mature messaging UX, not deployment, and not a production/security claim.
    dev-runtime-openmls enforced live umbrella checkout markers for:
      carbonstack/.git
      carbonstack-comms/.git
      carbonstack-cypher/.git
    dev-runtime-openmls ran the comms smoke script through Runner.RunStep.
    The smoke script proved:
      openmls-send-dev -> Cypher -> openmls-inbox-dev --ack
    The smoke script verified:
      plaintext: hello bob through openmls runtime smoke
      message_opened: true
      acked: true
      recipient inbox empty after ack
    post-dev-runtime-openmls artifact scan found only known OpenMLS sidecar generated roots.
    --clean-generated removed:
      .carbonstack-openmls-sidecar-state
      target
    local-cypher passed.
    doctor passed.
    core --clean-generated passed.
    OpenMLS real-Cypher lifecycle passed during core.
    carbonstack-comms package tests passed during core.
    carbonstack-cypher package tests passed during core.
    full was inspected and remained unchanged:
      release-snapshot, then local-cypher
    final repo status showed carbonstack at 8eeadb2 and all four repos clean.

Expected live umbrella validation commands now include:

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go run . --profile dev-runtime-openmls --clean-generated
    go run . --profile local-cypher
    go run . --profile doctor
    go run . --profile core --clean-generated

Expected v0.4.0 public package validation commands remain unchanged:

    cd <package-root>/carbonstack/tools/carbonstack-validate
    go run . --profile verify-checksums --root <package-root>
    go run . --profile full --root <package-root> --clean-generated

Current direct smoke command remains:

    cd ~/repos/carbonstack_umbrella/carbonstack-comms
    scripts/dev-openmls-runtime-smoke.sh

---

## 5. v0.4.9 Work Completed

### 5.1 Runner-profile recon confirmed implementation was small enough

The v0.4.9 recon inspected:

    v0.4.7/v0.4.8 runner-profile guidance.
    roadmap v0.4.7/v0.4.8/v0.4.9 area.
    carbonstack validation runner architecture.
    tools/carbonstack-validate profile dispatch.
    runner README/profile documentation.
    all runner Go files.
    carbonstack-comms dev runtime smoke script.
    README smoke docs.
    artifact scan / cleanup implementation.
    live direct smoke behavior.
    current validation ladder.

Key findings:

    Runner dispatch is centralized in tools/carbonstack-validate/main.go.
    Runner already has reusable RunStep.
    Runner already has non-destructive ArtifactScan.
    Runner already has CleanGeneratedArtifacts triggered globally by --clean-generated after successful profiles.
    The smoke script is deterministic and self-contained enough to wrap.
    The smoke script still directly bootstraps sidecar identity/KeyPackage/Welcome/conversation setup.
    The correct profile name is dev-runtime-openmls, not local-backbone.
    The first profile should be manual-only and live-umbrella-only.
    full should remain unchanged until dev-runtime-openmls maturity testing is complete.

### 5.2 dev_runtime_openmls.go was added

New file:

    carbonstack/tools/carbonstack-validate/dev_runtime_openmls.go

New functions:

    Runner.DevRuntimeOpenMLS()
    Runner.CheckLiveGitUmbrella(profileName string)

The profile:

    prints dev/pre-alpha/manual status.
    prints scope and proof:
      openmls-send-dev -> Cypher -> openmls-inbox-dev --ack
    prints explicit boundaries:
      not local-backbone
      not mature messaging UX
      not deployment
      not production/security claim
      live umbrella checkout only for now
      not included in full
    checks required sibling paths.
    checks live git markers for carbonstack, carbonstack-comms, and carbonstack-cypher.
    reports go/rustc/cargo/curl/python3/bash versions.
    runs pre-dev-runtime-openmls artifact scan.
    runs carbonstack-comms/scripts/dev-openmls-runtime-smoke.sh through RunStep.
    runs post-dev-runtime-openmls artifact scan.
    recommends --clean-generated after successful runs.

### 5.3 main.go profile dispatch was updated

Updated file:

    carbonstack/tools/carbonstack-validate/main.go

Changes:

    Added dev-runtime-openmls to the profile flag description.
    Added dispatch case:
      case "dev-runtime-openmls":
          runErr = r.DevRuntimeOpenMLS()
    Updated unknown-profile error text.

Important continuity:

    full remains unchanged.
    full still runs:
      release-snapshot
      local-cypher
    full does not call dev-runtime-openmls.

### 5.4 Runner README was updated

Updated file:

    carbonstack/tools/carbonstack-validate/README.md

Added section:

    dev-runtime-openmls

It documents:

    command:
      go run . --profile dev-runtime-openmls --clean-generated
    manual-only.
    live-umbrella-only.
    sibling git checkout requirement.
    wraps:
      carbonstack-comms/scripts/dev-openmls-runtime-smoke.sh
    boundaries:
      not local-backbone
      not mature messaging UX
      not deployment
      not release-package validation yet
      not production/security proof
      not included in full
    --clean-generated recommendation.

### 5.5 Main docs and roadmap were updated

New file:

    carbonstack/docs/159-dev-runtime-openmls-runner-profile-v0.md

Updated files:

    carbonstack/docs/README.md
    carbonstack/roadmap/ROADMAP.md

Docs record:

    v0.4.9 adds the manual dev-runtime-openmls profile.
    The profile wraps the Comms smoke proof.
    It remains live-umbrella-only.
    It is not included in full.
    It is not local-backbone.
    It is not release-package validation yet.
    Existing send/inbox remain stub-era.
    The suggested next rung is v0.4.10 validation/profile polish or sidecar bootstrap wrapper recon.

### 5.6 Validation passed

Commands observed passing:

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go test ./... -count=1
    go run . --profile dev-runtime-openmls --clean-generated
    go run . --profile local-cypher
    go run . --profile doctor
    go run . --profile core --clean-generated

Additional verification:

    grep -n 'case "full"' -A12 main.go

Confirmed full remains:

    release-snapshot
    local-cypher

It does not include dev-runtime-openmls.

Final commit/push:

    carbonstack:
    8eeadb2 feat: add dev OpenMLS runtime validation profile

Final repo snapshot:

    carbonstack        8eeadb2 feat: add dev OpenMLS runtime validation profile
    carbonstack-comms  0a48ae1 refactor: split OpenMLS runtime command helpers
    carbonstack-cypher 9ab994c docs: point to local operator runbook
    carbonstack-os     1bbbe52 docs: clarify CarbonStackOS target direction

---

## 6. Critical Blunders / Lessons to Carry Forward

### 6.1 First push failed authentication, follow-up push succeeded

The first `git push` after the v0.4.9 commit failed:

    remote: Failed to authenticate user
    fatal: Authentication failed for 'https://git.bitcrusher32.win/bitcrusher32/carbonstack.git/'

A follow-up `git push` succeeded and pushed:

    dc2d16c..8eeadb2  main -> main

Final snapshot showed:

    carbonstack 8eeadb2 (HEAD -> main, origin/main, origin/HEAD)

Lesson:

    Always trust the final repo-head snapshot over the first push attempt.
    Authentication hiccups are not project-state failures if a later push succeeds and origin/main matches HEAD.

### 6.2 Commit metadata warning remains a polish issue

The commit log again showed auto-generated local commit metadata:

    Committer: bitcrusher32 <▓▓>

This is not a functional validation failure, but it remains a professionalism/publishing hygiene issue to address eventually with `git config --global user.name` and `git config --global user.email` if desired.

### 6.3 One grep path typo was harmless

The final verification command included:

    grep -RIn "dev-runtime-openmls" main.go README.md docs/../../roadmap/ROADMAP.md

from inside `tools/carbonstack-validate`, which produced:

    grep: docs/../../roadmap/ROADMAP.md: No such file or directory

This was harmless because the actual diffs and repo status were already inspected, validation passed, and the final committed files were correct.

Lesson:

    When grepping from tools/carbonstack-validate, prefer absolute or repo-root anchored paths.

### 6.4 Generated roots are expected after dev-runtime-openmls

`dev-runtime-openmls` runs the smoke proof, which uses the OpenMLS sidecar and Rust build path. Known generated roots may appear:

    carbonstack-comms/internal/protocol/mls/openmls-sidecar/.carbonstack-openmls-sidecar-state
    carbonstack-comms/internal/protocol/mls/openmls-sidecar/target

This is expected as long as hits remain limited to known generated roots and `--clean-generated` removes them after success.

### 6.5 Manual profile does not mean local-backbone

v0.4.9 is a meaningful repeatability improvement, but the local-backbone boundary still holds:

    direct sidecar bootstrap remains.
    sidecar identity/KeyPackage/Welcome/conversation setup is not Comms runtime UX.
    send/inbox remain stub-era.
    profile is live-umbrella-only.
    profile is not included in full.
    profile is not release-package validation.

Correct phrase:

    manual dev-runtime-openmls validation profile

Still reserved:

    local-backbone

---

## 7. Current Blockers / Not Validated

Current blockers and nonvalidated areas after v0.4.9:

    Existing send/inbox remain stub-era.
    Runtime Comms has dev-only OpenMLS commands, a smoke proof, and now a manual runner profile, but no mature user-facing UX.
    Sidecar KeyPackage/Welcome/bootstrap setup is still direct dev setup, not Comms runtime UX.
    No Relay Space join model exists yet.
    local-backbone runner profile does not exist and remains reserved.
    dev-runtime-openmls exists but is manual-only.
    dev-runtime-openmls is not included in full.
    dev-runtime-openmls is live-umbrella-only.
    dev-runtime-openmls is not part of public release package validation.
    Fresh release-package behavior for dev-runtime-openmls has not been tested as a supported validation surface.
    OpenMLS runtime helpers are split out, but no new bootstrap wrappers exist yet.
    Post-release public download verification is still useful if not already performed.
    Sanitized PRIME placement in carbonstack/logdoc-loop-system still needs a policy decision.
    Production-safe provider storage and secure vault design remain v0.5.x future work.
    PQ/hybrid ciphersuite migration readiness remains v0.5.x future work.
    Hostile-server rollback/replay/metadata-abuse validation remains future work.
    Android/CarbonStackOS implementation remains later work.
    CarbonStack Relay Space mechanics remain a future concept, not an implementation.
    Public ingress, systemd, cloudflared, real homelab deployment, and production service paths remain out of scope.

Hard nonclaims remain:

    Do not claim production readiness.
    Do not claim production E2EE.
    Do not claim hostile-server safety.
    Do not claim metadata privacy.
    Do not claim Debian production deployability.
    Do not claim systemd readiness.
    Do not claim cloudflared readiness.
    Do not claim Android readiness.
    Do not claim CarbonStackOS readiness.
    Do not claim secure vault/storage.
    Do not claim real homelab validation.
    Do not claim external audit or certification.
    Do not claim CarbonStack Relay Space mechanics are implemented.
    Do not claim local-backbone exists or is justified yet.
    Do not claim local-cypher validates runtime Comms UX.
    Do not claim full validates local-backbone.
    Do not claim full is a deployment command.
    Do not claim dev-runtime-openmls is release-package validation yet.
    Do not claim send/inbox are OpenMLS-backed yet.
    Do not claim PQ/hybrid security or quantum-safe messaging.
    Do not let GitHub mirrors supersede Gitea source-of-truth status.

Allowed claim:

    CarbonStackComms now has dev-only openmls-send-dev and openmls-inbox-dev commands.
    CarbonStackComms now has a dev runtime smoke script proving openmls-send-dev -> Cypher -> openmls-inbox-dev --ack.
    CarbonStackComms now has OpenMLS runtime command helper glue split into internal/app/openmls_runtime.go.
    The main carbonstack repo records the smoke proof, the pre-local-backbone assessment, and the manual dev-runtime-openmls runner profile.
    carbonstack/tools/carbonstack-validate now has a manual dev-runtime-openmls profile.
    dev-runtime-openmls wraps the comms smoke script and validates the dev/pre-alpha OpenMLS application-message CLI path.
    dev-runtime-openmls is live-umbrella-only and not included in full.
    The smoke script starts a temporary local Cypher server, uses temporary Comms state files, creates dev-local OpenMLS sidecar identities/conversation, sends an OpenMLS application message through the Comms CLI, opens it through the Comms CLI, acks after sidecar success, and verifies Bob's inbox is empty after ack.
    Existing send/inbox remain stub-era.
    v0.4.0 remains the current public broad local deployability pre-release.
    Gitea remains source of truth; GitHub mirrors are secondary push mirrors.

---

## 8. Next Safest Actions

Recommended next rung:

    v0.4.10 validation/profile polish or sidecar bootstrap wrapper recon

Two viable paths:

    A. validation/profile polish:
       inspect whether dev-runtime-openmls should pre-clean known generated roots before running.
       inspect fresh live-checkout behavior after prior generated roots exist.
       verify profile output remains clear and non-overclaiming.
       decide whether docs should state manual-only even more strongly.
       keep full unchanged.

    B. sidecar bootstrap wrapper recon:
       inspect whether Comms should add dev-only wrappers for:
         identity-create
         public-bundle-export
         conversation-create
         conversation-add-member
         conversation-join
         conversation-load-check
       keep them explicitly dev-only.
       do not conflate them with final Relay Space join UX.

Preferred lean path:

    do a short preflight first.
    If no urgent cleanup issue appears, choose between:
      v0.4.10 dev-runtime-openmls polish
      or
      v0.4.10 sidecar bootstrap wrapper recon

Avoid next:

    adding dev-runtime-openmls to full immediately.
    replacing send/inbox immediately.
    adding local-backbone prematurely.
    adding a runner profile named local-backbone.
    moving dev-runtime-openmls into public release validation without fresh-root behavior checks.
    public ingress.
    cloudflared.
    systemd.
    real homelab deployment.
    Android app.
    CarbonStackOS implementation.
    CarbonStack Relay Space implementation.
    PQ/hybrid ciphersuite implementation.
    production/security claims.

---

## 9. Critical Paths / Functions

### Source-of-truth public surfaces

    [GITEA SOURCE OF TRUTH] https://git.bitcrusher32.win/bitcrusher32/carbonstack
    [RELEASES] https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases
    [CURRENT PUBLIC RELEASE] https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.4.0

GitHub mirrors:

    secondary push mirrors for discoverability/redundancy only.
    Do not treat mirrors as release authority unless project policy changes.

### WSL Debian active working umbrella

    [WSL-DEBIAN] umbrella: ~/repos/carbonstack_umbrella
    [WSL-DEBIAN] carbonstack: ~/repos/carbonstack_umbrella/carbonstack
    [WSL-DEBIAN] carbonstack-comms: ~/repos/carbonstack_umbrella/carbonstack-comms
    [WSL-DEBIAN] carbonstack-cypher: ~/repos/carbonstack_umbrella/carbonstack-cypher
    [WSL-DEBIAN] carbonstack-os: ~/repos/carbonstack_umbrella/carbonstack-os
    [GO-RUNNER] runner dir: ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate

### Runtime integration doc surfaces

    [DOC] v0.4.1 runtime recon:
    carbonstack/docs/155-runtime-comms-openmls-cypher-recon-v0.md

    [DOC] v0.4.2 command contract:
    carbonstack/docs/156-runtime-openmls-command-contract-v0.md

    [DOC] v0.4.6 smoke proof record:
    carbonstack/docs/157-dev-runtime-openmls-smoke-proof-v0.md

    [DOC] v0.4.7 pre-local-backbone assessment:
    carbonstack/docs/158-pre-local-backbone-assessment-v0.md

    [DOC] v0.4.9 runner profile:
    carbonstack/docs/159-dev-runtime-openmls-runner-profile-v0.md

    [DOC] docs index:
    carbonstack/docs/README.md

    [DOC] roadmap:
    carbonstack/roadmap/ROADMAP.md

    [DOC] top README:
    carbonstack/README.md

### Runner surfaces

    [RUNNER] root:
    carbonstack/tools/carbonstack-validate

    [RUNNER] dispatch:
    carbonstack/tools/carbonstack-validate/main.go

    [RUNNER] dev-runtime-openmls implementation:
    carbonstack/tools/carbonstack-validate/dev_runtime_openmls.go

    [RUNNER] runner README:
    carbonstack/tools/carbonstack-validate/README.md

    [RUNNER] local-cypher:
    carbonstack/tools/carbonstack-validate/local_cypher.go

    [RUNNER] release-snapshot:
    carbonstack/tools/carbonstack-validate/release_snapshot.go

    [RUNNER] checksums:
    carbonstack/tools/carbonstack-validate/checksums.go

### Comms runtime surfaces

    [COMMS] CLI entry:
    carbonstack-comms/cmd/comms/main.go

    [COMMS] command registry and non-OpenMLS commands:
    carbonstack-comms/internal/app/commands.go

    [COMMS] OpenMLS runtime command helpers:
    carbonstack-comms/internal/app/openmls_runtime.go

    [COMMS] openmls-send-dev tests:
    carbonstack-comms/internal/app/openmls_send_dev_test.go

    [COMMS] openmls-inbox-dev tests:
    carbonstack-comms/internal/app/openmls_inbox_dev_test.go

    [COMMS] dev runtime smoke proof:
    carbonstack-comms/scripts/dev-openmls-runtime-smoke.sh

    [COMMS] Cypher client:
    carbonstack-comms/internal/client/cypher.go

    [COMMS] relay helpers:
    carbonstack-comms/internal/relay/cypher_bridge.go
    carbonstack-comms/internal/relay/openmls_artifacts.go

    [COMMS] OpenMLS sidecar dir:
    carbonstack-comms/internal/protocol/mls/openmls-sidecar

    [COMMS] sidecar README:
    carbonstack-comms/internal/protocol/mls/openmls-sidecar/README.md

### Current stub-era Comms functions / behaviors

    send:
      trust policy evaluation
      mock/stub provider encryption
      Cypher envelope submit

    inbox:
      Cypher inbox retrieval
      mock/stub provider plaintext display as stub_plaintext

    ack:
      envelope ack by envelope ID and recipient device ID

### Implemented v0.4.x dev OpenMLS commands

    openmls-send-dev:
      internal/app/openmls_runtime.go
      state.RequireReadyDevice
      trust.EvaluateSend
      runOpenMLSMessageProtect
      sidecar message-protect
      internal/relay.SubmitOpenMLSArtifactEnvelope
      ArtifactKindApplicationMessage
      dev/pre-alpha warning

    openmls-inbox-dev:
      internal/app/openmls_runtime.go
      state.RequireReadyDevice
      Cypher inbox
      skip unsupported envelopes
      internal/relay.WriteOpenMLSArtifactFromEnvelope
      runOpenMLSMessageOpen
      sidecar message-open
      print plaintext after sidecar success
      optional ack only after sidecar success

### v0.4.5/v0.4.8/v0.4.9 runtime smoke proof

    scripts/dev-openmls-runtime-smoke.sh:
      temp workspace
      temp Cypher binary
      temp Cypher DB
      dynamic loopback Cypher port
      temp Alice/Bob Comms state
      sidecar state reset
      sidecar identity/conversation bootstrap
      openmls-send-dev application-message send
      openmls-inbox-dev application-message open + ack
      inbox-empty-after-ack verification
      robust cleanup with INT -> TERM -> KILL escalation

### v0.4.9 runner profile

    go run . --profile dev-runtime-openmls --clean-generated

    Runner.DevRuntimeOpenMLS:
      CheckRequiredPaths
      CheckLiveGitUmbrella
      toolchain reporting
      ArtifactScan("pre-dev-runtime-openmls")
      RunStep for carbonstack-comms/scripts/dev-openmls-runtime-smoke.sh
      ArtifactScan("post-dev-runtime-openmls")
      explicit nonclaim summary

    Runner.CheckLiveGitUmbrella:
      requires .git markers in carbonstack, carbonstack-comms, carbonstack-cypher

### OpenMLS sidecar commands relevant to v0.4.x

    provider-info
    identity-create
    identity-status
    public-bundle-export
    conversation-create
    conversation-load-check
    conversation-add-member
    conversation-join
    message-protect
    message-open

### Relay/API functions relevant to runtime integration

    internal/client.CypherClient.SubmitEnvelope
    internal/client.CypherClient.Inbox
    internal/client.CypherClient.AckEnvelope
    internal/relay.SubmitOpenMLSArtifactEnvelope
    internal/relay.WriteOpenMLSArtifactFromEnvelope

### Existing v0.4.0 package validation

    cd <package-root>/carbonstack/tools/carbonstack-validate
    go run . --profile verify-checksums --root <package-root>
    go run . --profile full --root <package-root> --clean-generated

### Live umbrella validation

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go run . --profile dev-runtime-openmls --clean-generated
    go run . --profile local-cypher
    go run . --profile doctor
    go run . --profile core --clean-generated

### Envelope/protocol constants

Accepted opaque OpenMLS application-message pair:

    content_type=carbonstack.mls.application-message.v0
    protocol_version=carbonstack-openmls-sidecar-v0

Historical invalid pair validated by local-cypher:

    content_type=carbonstack.message.text.stub.v0
    protocol_version=carbonstack-openmls-sidecar-v0

Expected result:

    HTTP 400
    error code unsupported_protocol_version

---

## 10. v0.4.x Forward Plan

### v0.4.10 — validation/profile polish or sidecar bootstrap wrapper recon

Option A:

    Dev-runtime-openmls profile polish:
      inspect generated-root pre-clean behavior.
      inspect repeated profile run behavior.
      decide if profile should pre-warn more strongly when known roots already exist.
      keep full unchanged.
      keep release-package validation unchanged.

Option B:

    Dev-only OpenMLS bootstrap wrapper recon:
      identity-create
      public-bundle-export
      conversation-create
      conversation-add-member
      conversation-join
      conversation-load-check

Preferred sequence:

    short preflight first.
    choose based on current risk:
      if validation profile feels brittle, polish it.
      if profile feels stable, recon bootstrap wrappers.

### v0.4.11+

Possible paths:

    add dev-only bootstrap wrappers.
    improve sidecar setup/reset ergonomics.
    document command migration/deprecation options for stub-era send/inbox.
    test dev-runtime-openmls from cleaner live checkouts.
    later consider release-root behavior, but do not add to full yet.

### v0.4.x later polish

Likely needs:

    exact sidecar setup runbook.
    clearer dev state reset semantics.
    better command output shape.
    decision on whether stub-era send/inbox are aliased, replaced, or deprecated.
    decision on when local-backbone naming stops being premature.

---

## 11. v0.5.x and Later Carry-Forward

### v0.5.x

Local storage, trust state, provider state, vault design, and PQ/hybrid ciphersuite migration readiness.

Do not start PQ/hybrid implementation before the runtime OpenMLS/Cypher path exists and is documented. PQ/hybrid migration touches provider state, ciphersuite metadata, trust display, artifact size, reinit/rekey/recovery policy, and claim boundaries.

### v0.6.x+

Hostile-server and abuse-resistance harnesses.

### v0.7.x+

Deployability and operations hardening.

### v0.8.x+

Documented self-pentest / adversarial validation.

### v0.9.x+

Claim-boundary review.

### v0.10.x+

Android backend/app work.

### v1.x.x

Public app/server major epoch and much later CarbonStackOS/appliance research, if justified.

---

## 12. Lean Breakpoint Summary

    CarbonStack v0.4.9 is the manual dev-runtime-openmls runner-profile checkpoint. v0.4.0 remains the current public broad local deployability pre-release. carbonstack is now at 8eeadb2 feat: add dev OpenMLS runtime validation profile. carbonstack-comms remains at 0a48ae1 refactor: split OpenMLS runtime command helpers. v0.4.9 adds carbonstack/tools/carbonstack-validate/dev_runtime_openmls.go, updates runner profile dispatch in main.go, updates tools/carbonstack-validate/README.md, adds docs/159-dev-runtime-openmls-runner-profile-v0.md, updates docs/README.md, and updates roadmap/ROADMAP.md. The new manual profile wraps the Comms smoke script and validates openmls-send-dev -> Cypher -> openmls-inbox-dev --ack from a live umbrella checkout. It is not included in full, not release-package validation yet, not local-backbone, not mature messaging UX, and not production/security proof. Validation passed for go test ./... in the runner, dev-runtime-openmls --clean-generated, local-cypher, doctor, and core --clean-generated, with final clean repo heads across all four repos. Next safe rung: v0.4.10 validation/profile polish or sidecar bootstrap wrapper recon.

---

## 13. Preserved Immediate Previous Handoff: v0.4.8

The following is the previous v0.4.8 handoff. Where it conflicts with the v0.4.9 overlay above, v0.4.9 wins for current state.

# CarbonStack LogDoc v0.4.8

**Last updated:** 2026-06-04 local session  
**Owner/operator:** bitcrusher32  
**Assistant/model:** GPT-5.5 Thinking  
**Current phase:** Phase 2F / **v0.4.x behavior-preserving OpenMLS command/helper extraction checkpoint complete**. The v0.4.0 public broad local deployability pre-release remains the current public release. v0.4.5 proved the dev runtime OpenMLS CLI path; v0.4.6 recorded that smoke proof in the main `carbonstack` repo; v0.4.7 assessed the proof and kept `local-backbone` reserved; v0.4.8 now performs the planned behavior-preserving extraction in `carbonstack-comms`, moving OpenMLS runtime command glue out of `internal/app/commands.go` into `internal/app/openmls_runtime.go` without changing command behavior. `carbonstack` remains at `dc2d16c (HEAD -> main, origin/main, origin/HEAD) docs: assess pre-local-backbone runtime proof`; `carbonstack-comms` is now at `0a48ae1 (HEAD -> main, origin/main, origin/HEAD) refactor: split OpenMLS runtime command helpers`; `carbonstack-cypher` remains at `9ab994c (HEAD -> main, origin/main, origin/HEAD) docs: point to local operator runbook`; `carbonstack-os` remains at `1bbbe52 (HEAD -> main, origin/main, origin/HEAD) docs: clarify CarbonStackOS target direction`.

**Version schema:** `v[scope].[timeline]`. This file is `v0.4.8`, the behavior-preserving OpenMLS command/helper extraction checkpoint after the v0.4.7 pre-local-backbone assessment. Per LogDoc V2 practice, this Markdown file preserves current operational continuity and carries the v0.4.x runtime-integration thread forward; the JSON breakpoint remains the lean machine-readable handoff state.

**Compression note:** The detailed v0.3.0 through v0.3.36PRIME process ledger remains archived in PRIME/sanitized PRIME. v0.4.0 compressed that pre-release sprint; v0.4.1 preserved the runtime recon; v0.4.2 preserved the command-contract decision; v0.4.3 preserved the first dev-only OpenMLS send-side implementation; v0.4.4 preserved the first dev-only OpenMLS receive/open/optional-ack implementation; v0.4.5 preserved the first committed dev runtime smoke proof; v0.4.6 preserved project-level documentation/status alignment for that proof; v0.4.7 preserved the decision boundary before any local-backbone/runner-profile promotion; v0.4.8 now preserves the behavior-preserving code organization step that makes future OpenMLS runtime growth safer.

---

## 1. Project Goal

**Active goal:** Continue moving CarbonStackComms from stub-era runtime `send`/`inbox` behavior toward explicit dev-only OpenMLS/Cypher-backed runtime commands while preserving claim discipline and keeping command implementation maintainable. v0.4.8 completed a behavior-preserving extraction: OpenMLS runtime command glue now lives in a dedicated app file instead of bloating `commands.go`.

CarbonStack’s broader project intent remains:

- **CarbonStack** is the secure-communications appliance-stack project and shared doctrine/release authority.
- **CarbonStackComms** is the text-first messaging client and current OpenMLS sidecar/relay integration surface.
- **CarbonStackCypher** is the self-hostable relay/storage stack.
- **CarbonStackOS** is future constrained Android-derived appliance OS work and remains out of current runnable validation packages.
- The `carbonstack` Gitea repo remains the source of truth for release state, release assets, docs, and public claims.
- GitHub mirrors exist as secondary push mirrors only; Gitea remains authoritative.

**North star:**

> Its goal is not to be a better smartphone. Its goal is to stop being a smartphone.

**Core doctrine:**

> Every feature is guilty until it proves it does not add unacceptable parser, network, sensor, identity, or filesystem authority.

**Current reality after v0.4.8:** the dev runtime OpenMLS application-message smoke proof remains implemented and documented; the v0.4.7 assessment remains valid; and the OpenMLS runtime command implementation has now been split into a dedicated file with behavior preserved.

Correct interpretation:

    openmls-send-dev is a dev-only OpenMLS send-side command.
    openmls-inbox-dev is a dev-only OpenMLS receive/open/optional-ack command.
    scripts/dev-openmls-runtime-smoke.sh proves the dev CLI application-message path:
      openmls-send-dev -> Cypher -> openmls-inbox-dev --ack
    docs/157-dev-runtime-openmls-smoke-proof-v0.md records that proof.
    docs/158-pre-local-backbone-assessment-v0.md records that the proof should remain a comms-local dev helper for now.
    internal/app/openmls_runtime.go now holds the OpenMLS runtime command helpers.
    internal/app/commands.go still owns command registration, non-OpenMLS app commands, and stub-era send/inbox/ack.
    A future runner profile may be named dev-runtime-openmls, not local-backbone.
    The smoke script directly bootstraps sidecar identities, KeyPackage/Welcome, and conversation setup for dev purposes.
    Existing send/inbox remain stub-era.
    local-backbone remains reserved.

Hard nonclaims remain: v0.4.8 does **not** prove production readiness, production E2EE, hostile-server safety, metadata privacy, Android readiness, CarbonStackOS readiness, secure vault/storage, public ingress safety, real homelab validation, audit, certification, PQ/hybrid security, mature messenger UX, or local-backbone readiness.

---

## 2. Current Public Release State

Public release:

    https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.4.0

Release title:

    CarbonStack v0.4.0 Broad Local Deployability Pre-Release

Release tag:

    v0.4.0

Release commit:

    24ac9fc494 / 24ac9fc docs: record v0.4.0 package rehearsal

Gitea release status:

    Pre-Release

Release assets attached:

    carbonstack-v0.4.0-broad-local-deployability-pre-release.tgz
    carbonstack-v0.4.0-release-manifest.json
    carbonstack-v0.4.0-package-checksums.txt
    carbonstack-v0.4.0-asset-checksums.txt
    carbonstack-v0.4.0-validation-freeze.md
    v0.4.0-testing-runbook.md
    v0.4.0-release-notes.md
    LICENSE

Important release-page warning:

    The default Gitea Source Code ZIP/TAR.GZ downloads are auto-generated carbonstack-only archives.
    The intended runnable validation artifact is the attached multi-repo release package plus companion assets.

Post-release public surface:

    carbonstack README is evergreen.
    It points interested parties to the Releases page for runnable / known-good artifacts.
    It states Gitea is the source of truth.
    GitHub mirrors are secondary push mirrors, not release authority.

---

## 3. Current Repo Heads

    carbonstack        dc2d16c (HEAD -> main, origin/main, origin/HEAD) docs: assess pre-local-backbone runtime proof
    carbonstack-comms  0a48ae1 (HEAD -> main, origin/main, origin/HEAD) refactor: split OpenMLS runtime command helpers
    carbonstack-cypher 9ab994c (HEAD -> main, origin/main, origin/HEAD) docs: point to local operator runbook
    carbonstack-os     1bbbe52 (HEAD -> main, origin/main, origin/HEAD) docs: clarify CarbonStackOS target direction

Current public release tag:

    v0.4.0 -> 24ac9fc494 / 24ac9fc docs: record v0.4.0 package rehearsal

Post-release public-surface cleanup:

    6ba7aab docs: make front readme evergreen

v0.4.1 runtime recon commit:

    83d2f0f docs: record runtime OpenMLS Comms recon

v0.4.2 command contract commit:

    a799764 docs: define runtime OpenMLS command contract

v0.4.3 send implementation commit:

    39fb287 feat: add dev OpenMLS send command

v0.4.4 inbox implementation commit:

    65bc707 feat: add dev OpenMLS inbox command

v0.4.5 smoke proof commit:

    8e6e8b4 test: add dev OpenMLS runtime smoke proof

v0.4.6 main-doc alignment commit:

    3cd77a8 docs: record dev OpenMLS runtime smoke proof

v0.4.7 assessment commit:

    dc2d16c docs: assess pre-local-backbone runtime proof

Current v0.4.8 extraction commit:

    0a48ae1 refactor: split OpenMLS runtime command helpers

Continuity note:

    `0a48ae1` is a `carbonstack-comms` behavior-preserving refactor commit after the `v0.4.0` public release tag.
    It does not retag v0.4.0.
    The release remains attached to `24ac9fc494`.

---

## 4. Current Validated / Known State

Validated / known at the v0.4.8 checkpoint:

    [RELEASE] v0.4.0 Gitea release exists, is marked Pre-Release, and is attached to 24ac9fc494.
    [PACKAGE] v0.4.0 package includes carbonstack, carbonstack-comms, carbonstack-cypher, and release metadata.
    [PACKAGE] carbonstack-os is deliberately excluded from the runnable package.
    [RUNNER] full means release-snapshot followed by local-cypher.
    [RUNNER] release-snapshot validates package layout, metadata, strict pre-test artifact scan, package checksums, and core.
    [RUNNER] core validates doctor, OpenMLS real-Cypher lifecycle, Comms package tests, Cypher package tests, and artifact scans.
    [RUNNER] local-cypher validates Cypher-only lifecycle, restart/persistence against the same temporary DB, positive opaque OpenMLS envelope/ack flow, and historical negative stub-text/OpenMLS protocol pairing.
    [DOCS] main carbonstack README is evergreen and lists the dev runtime OpenMLS CLI smoke proof as demonstrated experimental validation coverage.
    [DOCS] docs/155-runtime-comms-openmls-cypher-recon-v0.md records the v0.4.1 runtime recon.
    [DOCS] docs/156-runtime-openmls-command-contract-v0.md records the v0.4.2 command contract.
    [DOCS] docs/157-dev-runtime-openmls-smoke-proof-v0.md records the v0.4.5 dev runtime smoke proof.
    [DOCS] docs/158-pre-local-backbone-assessment-v0.md records the v0.4.7 assessment and plan.
    [COMMS] openmls-send-dev exists in carbonstack-comms.
    [COMMS] openmls-inbox-dev exists in carbonstack-comms.
    [COMMS] scripts/dev-openmls-runtime-smoke.sh exists in carbonstack-comms.
    [COMMS] internal/app/openmls_runtime.go now contains OpenMLS runtime command helper code.
    [COMMS] internal/app/commands.go now has less runtime-integration glue and remains the command registry / non-OpenMLS command surface.
    [COMMS] existing send/inbox/ack remain in place and stub-era.
    [MIRRORS] all four repos have GitHub push-mirror functionality, with Gitea remaining source of truth.

v0.4.8 extraction result:

    The OpenMLS block was moved out of internal/app/commands.go.
    New file internal/app/openmls_runtime.go was created.
    commands.go imports were reduced by removing OpenMLS-only imports.
    Run(...) and usage() remained in commands.go.
    Stub-era send/inbox/ack remained in commands.go.
    openmls-send-dev behavior remained unchanged.
    openmls-inbox-dev behavior remained unchanged.
    Test seams remained package-level and continued working.
    No runner profile was added.
    No local-backbone claim was added.

v0.4.8 validation result from the final log:

    go test ./internal/app -count=1 passed.
    go test ./... -count=1 -timeout 600s passed in carbonstack-comms.
    scripts/dev-openmls-runtime-smoke.sh passed.
    The smoke script proved openmls-send-dev -> Cypher -> openmls-inbox-dev --ack.
    The smoke script verified plaintext: hello bob through openmls runtime smoke.
    The smoke script verified acked: true and Bob inbox empty after ack.
    local-cypher passed.
    doctor passed.
    core --clean-generated passed.
    OpenMLS real-Cypher lifecycle passed during core.
    carbonstack-comms package tests passed during core.
    carbonstack-cypher package tests passed during core.
    Known OpenMLS generated roots were removed by --clean-generated.
    Final repo status showed carbonstack-comms at 0a48ae1 and all four repos clean.

Expected v0.4.0 public package validation commands remain:

    cd <package-root>/carbonstack/tools/carbonstack-validate
    go run . --profile verify-checksums --root <package-root>
    go run . --profile full --root <package-root> --clean-generated

Live umbrella validation commands remain:

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go run . --profile local-cypher
    go run . --profile doctor
    go run . --profile core --clean-generated

Current v0.4.5/v0.4.8 smoke command remains:

    cd ~/repos/carbonstack_umbrella/carbonstack-comms
    scripts/dev-openmls-runtime-smoke.sh

---

## 5. v0.4.8 Work Completed

### 5.1 Helper-extraction recon was performed

The v0.4.8 recon inspected:

    v0.4.7 assessment doc and roadmap next-rung guidance.
    carbonstack-comms internal/app file list.
    internal/app/commands.go size, imports, and symbol index.
    the full commands.go contents.
    openmls_send_dev_test.go.
    openmls_inbox_dev_test.go.
    OpenMLS command references.
    non-OpenMLS command references and shared helpers.
    import dependency pressure.
    app tests, full Comms tests, smoke proof, and main validation baseline.

Key finding:

    OpenMLS runtime glue was already a contiguous block in commands.go:
      defaultOpenMLSSidecarDir through parseOpenMLSOpenEnvelope

Decision:

    Move that contiguous OpenMLS block into a dedicated file.
    Keep command registration in commands.go.
    Keep non-OpenMLS commands in commands.go.
    Preserve behavior and test seams.

### 5.2 internal/app/openmls_runtime.go was added

New file:

    carbonstack-comms/internal/app/openmls_runtime.go

Moved into this file:

    defaultOpenMLSSidecarDir
    openMLSMessageProtectResult
    openMLSSidecarErrorEnvelope
    openMLSSidecarProtectEnvelope
    openMLSSidecarProtectData
    runOpenMLSMessageProtectForCommand
    submitOpenMLSArtifactEnvelopeForCommand
    cmdOpenMLSSendDev
    runOpenMLSMessageProtect
    parseOpenMLSProtectEnvelope
    openMLSMessageOpenResult
    openMLSSidecarOpenEnvelope
    openMLSSidecarOpenData
    inboxForCommand
    ackEnvelopeForCommand
    writeOpenMLSArtifactFromEnvelopeForCommand
    runOpenMLSMessageOpenForCommand
    cmdOpenMLSInboxDev
    writeOpenMLSInboxDevArtifact
    runOpenMLSMessageOpen
    parseOpenMLSOpenEnvelope

This split keeps the OpenMLS command implementation together and reduces pressure on commands.go before future sidecar bootstrap/runtime wrappers are considered.

### 5.3 internal/app/commands.go was simplified

Updated file:

    carbonstack-comms/internal/app/commands.go

Removed from imports because they are now OpenMLS-only and live in openmls_runtime.go:

    encoding/json
    git.bitcrusher32.win/bitcrusher32/carbonstack-comms/internal/relay
    os
    os/exec
    path/filepath
    strconv

Kept in commands.go:

    Run(args []string)
    usage()
    init
    dev-create-invite
    claim-invite
    register-device
    list-devices
    fingerprint
    verify-device
    trust-history
    trust-list
    simulate-key-change
    revoke-device
    send
    inbox
    ack
    sanitizeLabel

Important continuity:

    Run(...) still dispatches openmls-send-dev and openmls-inbox-dev.
    usage() still lists openmls-send-dev and openmls-inbox-dev.
    Existing stub-era send/inbox/ack behavior was not changed.

### 5.4 Validation passed

Commands observed passing:

    cd ~/repos/carbonstack_umbrella/carbonstack-comms
    go test ./internal/app -count=1
    go test ./... -count=1 -timeout 600s
    scripts/dev-openmls-runtime-smoke.sh

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go run . --profile local-cypher
    go run . --profile doctor
    go run . --profile core --clean-generated

Smoke proof result remained valid:

    PASS: dev runtime OpenMLS CLI smoke proof
    proof: openmls-send-dev -> Cypher -> openmls-inbox-dev --ack
    plaintext: hello bob through openmls runtime smoke
    boundary: dev/pre-alpha smoke proof; not local-backbone; not production messaging UX

Final commit/push:

    carbonstack-comms:
    0a48ae1 refactor: split OpenMLS runtime command helpers

Final repo snapshot:

    carbonstack        dc2d16c docs: assess pre-local-backbone runtime proof
    carbonstack-comms  0a48ae1 refactor: split OpenMLS runtime command helpers
    carbonstack-cypher 9ab994c docs: point to local operator runbook
    carbonstack-os     1bbbe52 docs: clarify CarbonStackOS target direction

---

## 6. Critical Blunders / Lessons to Carry Forward

### 6.1 Commit metadata warning remains a polish issue

The commit log again showed auto-generated local commit metadata:

    Committer: bitcrusher32 <▓▓>

This is not a functional validation failure, but it remains a professionalism/publishing hygiene issue to address eventually with `git config --global user.name` and `git config --global user.email` if desired.

### 6.2 Known generated roots appeared before cleanup

Because the smoke script and OpenMLS tests run sidecar flows, validation saw known generated roots:

    carbonstack-comms/internal/protocol/mls/openmls-sidecar/.carbonstack-openmls-sidecar-state
    carbonstack-comms/internal/protocol/mls/openmls-sidecar/target

The runner classified these as known OpenMLS generated roots and `core --clean-generated` removed them.

Lesson:

    This is expected after smoke/tests as long as hits remain limited to known generated roots.
    Use --clean-generated when ending a rung that runs OpenMLS sidecar tests or smoke proof.
    Do not treat these known roots as source hygiene failures unless they appear in an unexpected phase or outside known paths.

### 6.3 Behavior-preserving extraction worked because the block was contiguous

The extraction was safe because the OpenMLS command block was contiguous and had clear markers:

    defaultOpenMLSSidecarDir
    through
    parseOpenMLSOpenEnvelope

Lesson:

    Future extraction should continue to prefer small, contiguous, behavior-preserving moves.
    Avoid broad rewrites while the runtime path is still pre-alpha.

### 6.4 local-backbone remains reserved

Even after extraction, the state from v0.4.7 still holds:

    dev runtime smoke proof exists.
    OpenMLS command glue is cleaner.
    But direct sidecar bootstrap remains.
    No runner profile exists yet.
    No release-package proof exists yet.
    Existing send/inbox remain stub-era.

Correct current phrase remains:

    dev runtime OpenMLS smoke proof

Possible future runner profile remains:

    dev-runtime-openmls

Still reserved:

    local-backbone

---

## 7. Current Blockers / Not Validated

Current blockers and nonvalidated areas after v0.4.8:

    Existing send/inbox remain stub-era.
    Runtime Comms has dev-only OpenMLS commands and a smoke proof, but no mature user-facing UX.
    Sidecar KeyPackage/Welcome/bootstrap setup is still direct dev setup, not Comms runtime UX.
    No Relay Space join model exists yet.
    local-backbone runner profile does not exist and remains reserved.
    dev-runtime-openmls runner profile does not exist yet.
    The dev runtime smoke script is not yet integrated into the main carbonstack validation runner.
    The smoke proof is not yet part of a public release package validation surface.
    Fresh release-package behavior for the smoke script has not been tested as a runner target.
    OpenMLS runtime helpers are split out, but no new bootstrap wrappers exist yet.
    Post-release public download verification is still useful if not already performed.
    Sanitized PRIME placement in carbonstack/logdoc-loop-system still needs a policy decision.
    Production-safe provider storage and secure vault design remain v0.5.x future work.
    PQ/hybrid ciphersuite migration readiness remains v0.5.x future work.
    Hostile-server rollback/replay/metadata-abuse validation remains future work.
    Android/CarbonStackOS implementation remains later work.
    CarbonStack Relay Space mechanics remain a future concept, not an implementation.
    Public ingress, systemd, cloudflared, real homelab deployment, and production service paths remain out of scope.

Hard nonclaims remain:

    Do not claim production readiness.
    Do not claim production E2EE.
    Do not claim hostile-server safety.
    Do not claim metadata privacy.
    Do not claim Debian production deployability.
    Do not claim systemd readiness.
    Do not claim cloudflared readiness.
    Do not claim Android readiness.
    Do not claim CarbonStackOS readiness.
    Do not claim secure vault/storage.
    Do not claim real homelab validation.
    Do not claim external audit or certification.
    Do not claim CarbonStack Relay Space mechanics are implemented.
    Do not claim local-backbone exists or is justified yet.
    Do not claim local-cypher validates runtime Comms UX.
    Do not claim full validates local-backbone.
    Do not claim full is a deployment command.
    Do not claim send/inbox are OpenMLS-backed yet.
    Do not claim PQ/hybrid security or quantum-safe messaging.
    Do not let GitHub mirrors supersede Gitea source-of-truth status.

Allowed claim:

    CarbonStackComms now has dev-only openmls-send-dev and openmls-inbox-dev commands.
    CarbonStackComms now has a dev runtime smoke script proving openmls-send-dev -> Cypher -> openmls-inbox-dev --ack.
    CarbonStackComms now has OpenMLS runtime command helper glue split into internal/app/openmls_runtime.go.
    The main carbonstack repo records the smoke proof and the pre-local-backbone assessment.
    A future dev-runtime-openmls runner profile is plausible, but not implemented.
    The smoke script starts a temporary local Cypher server, uses temporary Comms state files, creates dev-local OpenMLS sidecar identities/conversation, sends an OpenMLS application message through the Comms CLI, opens it through the Comms CLI, acks after sidecar success, and verifies Bob's inbox is empty after ack.
    Existing send/inbox remain stub-era.
    v0.4.0 remains the current public broad local deployability pre-release.
    Gitea remains source of truth; GitHub mirrors are secondary push mirrors.

---

## 8. Next Safest Actions

Recommended next rung:

    v0.4.9 dev-runtime-openmls runner-profile recon from fresh package-like root

Focus:

    inspect whether scripts/dev-openmls-runtime-smoke.sh can be safely called by carbonstack/tools/carbonstack-validate.
    verify sibling-root assumptions from live umbrella and package-like roots.
    define exact profile name and nonclaim language.
    decide whether generated-root cleanup should be part of the profile or only driven by --clean-generated.
    confirm the profile would not overclaim local-backbone.
    avoid implementing runner profile until recon confirms fresh-root behavior.

Possible alternative next rung:

    v0.4.9 dev OpenMLS bootstrap-wrapper recon

This would inspect whether direct sidecar identity/public-bundle/conversation-add-member/conversation-join setup should get Comms dev wrappers before runner profile promotion.

Preferred order:

    runner-profile recon from fresh package-like root first.
    then decide between implementing dev-runtime-openmls profile or adding bootstrap wrappers.

Avoid next:

    replacing send/inbox immediately.
    adding local-backbone prematurely.
    adding a runner profile named local-backbone.
    moving smoke proof into public release validation without fresh-root behavior checks.
    public ingress.
    cloudflared.
    systemd.
    real homelab deployment.
    Android app.
    CarbonStackOS implementation.
    CarbonStack Relay Space implementation.
    PQ/hybrid ciphersuite implementation.
    production/security claims.

---

## 9. Critical Paths / Functions

### Source-of-truth public surfaces

    [GITEA SOURCE OF TRUTH] https://git.bitcrusher32.win/bitcrusher32/carbonstack
    [RELEASES] https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases
    [CURRENT PUBLIC RELEASE] https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.4.0

GitHub mirrors:

    secondary push mirrors for discoverability/redundancy only.
    Do not treat mirrors as release authority unless project policy changes.

### WSL Debian active working umbrella

    [WSL-DEBIAN] umbrella: ~/repos/carbonstack_umbrella
    [WSL-DEBIAN] carbonstack: ~/repos/carbonstack_umbrella/carbonstack
    [WSL-DEBIAN] carbonstack-comms: ~/repos/carbonstack_umbrella/carbonstack-comms
    [WSL-DEBIAN] carbonstack-cypher: ~/repos/carbonstack_umbrella/carbonstack-cypher
    [WSL-DEBIAN] carbonstack-os: ~/repos/carbonstack_umbrella/carbonstack-os
    [GO-RUNNER] runner dir: ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate

### Runtime integration doc surfaces

    [DOC] v0.4.1 runtime recon:
    carbonstack/docs/155-runtime-comms-openmls-cypher-recon-v0.md

    [DOC] v0.4.2 command contract:
    carbonstack/docs/156-runtime-openmls-command-contract-v0.md

    [DOC] v0.4.6 smoke proof record:
    carbonstack/docs/157-dev-runtime-openmls-smoke-proof-v0.md

    [DOC] v0.4.7 pre-local-backbone assessment:
    carbonstack/docs/158-pre-local-backbone-assessment-v0.md

    [DOC] docs index:
    carbonstack/docs/README.md

    [DOC] roadmap:
    carbonstack/roadmap/ROADMAP.md

    [DOC] top README:
    carbonstack/README.md

### Comms runtime surfaces

    [COMMS] CLI entry:
    carbonstack-comms/cmd/comms/main.go

    [COMMS] command registry and non-OpenMLS commands:
    carbonstack-comms/internal/app/commands.go

    [COMMS] OpenMLS runtime command helpers:
    carbonstack-comms/internal/app/openmls_runtime.go

    [COMMS] openmls-send-dev tests:
    carbonstack-comms/internal/app/openmls_send_dev_test.go

    [COMMS] openmls-inbox-dev tests:
    carbonstack-comms/internal/app/openmls_inbox_dev_test.go

    [COMMS] dev runtime smoke proof:
    carbonstack-comms/scripts/dev-openmls-runtime-smoke.sh

    [COMMS] Cypher client:
    carbonstack-comms/internal/client/cypher.go

    [COMMS] relay helpers:
    carbonstack-comms/internal/relay/cypher_bridge.go
    carbonstack-comms/internal/relay/openmls_artifacts.go

    [COMMS] OpenMLS sidecar dir:
    carbonstack-comms/internal/protocol/mls/openmls-sidecar

    [COMMS] sidecar README:
    carbonstack-comms/internal/protocol/mls/openmls-sidecar/README.md

### Current stub-era Comms functions / behaviors

    send:
      trust policy evaluation
      mock/stub provider encryption
      Cypher envelope submit

    inbox:
      Cypher inbox retrieval
      mock/stub provider plaintext display as stub_plaintext

    ack:
      envelope ack by envelope ID and recipient device ID

### Implemented v0.4.x dev OpenMLS commands

    openmls-send-dev:
      internal/app/openmls_runtime.go
      state.RequireReadyDevice
      trust.EvaluateSend
      runOpenMLSMessageProtect
      sidecar message-protect
      internal/relay.SubmitOpenMLSArtifactEnvelope
      ArtifactKindApplicationMessage
      dev/pre-alpha warning

    openmls-inbox-dev:
      internal/app/openmls_runtime.go
      state.RequireReadyDevice
      Cypher inbox
      skip unsupported envelopes
      internal/relay.WriteOpenMLSArtifactFromEnvelope
      runOpenMLSMessageOpen
      sidecar message-open
      print plaintext after sidecar success
      optional ack only after sidecar success

### v0.4.5/v0.4.8 runtime smoke proof function

    scripts/dev-openmls-runtime-smoke.sh:
      temp workspace
      temp Cypher binary
      temp Cypher DB
      dynamic loopback Cypher port
      temp Alice/Bob Comms state
      sidecar state reset
      sidecar identity/conversation bootstrap
      openmls-send-dev application-message send
      openmls-inbox-dev application-message open + ack
      inbox-empty-after-ack verification
      robust cleanup with INT -> TERM -> KILL escalation

### OpenMLS sidecar commands relevant to v0.4.x

    provider-info
    identity-create
    identity-status
    public-bundle-export
    conversation-create
    conversation-load-check
    conversation-add-member
    conversation-join
    message-protect
    message-open

### Relay/API functions relevant to runtime integration

    internal/client.CypherClient.SubmitEnvelope
    internal/client.CypherClient.Inbox
    internal/client.CypherClient.AckEnvelope
    internal/relay.SubmitOpenMLSArtifactEnvelope
    internal/relay.WriteOpenMLSArtifactFromEnvelope

### Existing v0.4.0 package validation

    cd <package-root>/carbonstack/tools/carbonstack-validate
    go run . --profile verify-checksums --root <package-root>
    go run . --profile full --root <package-root> --clean-generated

### Live umbrella validation

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go run . --profile local-cypher
    go run . --profile doctor
    go run . --profile core --clean-generated

### Envelope/protocol constants

Accepted opaque OpenMLS application-message pair:

    content_type=carbonstack.mls.application-message.v0
    protocol_version=carbonstack-openmls-sidecar-v0

Historical invalid pair validated by local-cypher:

    content_type=carbonstack.message.text.stub.v0
    protocol_version=carbonstack-openmls-sidecar-v0

Expected result:

    HTTP 400
    error code unsupported_protocol_version

---

## 10. v0.4.x Forward Plan

### v0.4.9 — dev-runtime-openmls runner-profile recon from fresh package-like root

Expected:

    inspect live umbrella and fresh package-like root behavior.
    confirm whether carbonstack-validate can call the comms-local smoke script safely.
    define exact profile name and nonclaim language.
    verify generated-root cleanup expectations.
    decide whether implementation is safe or whether bootstrap wrappers should come first.

### v0.4.10+ — possible dev-runtime-openmls profile implementation or bootstrap wrapper recon

Possible paths:

    add a clearly named dev-runtime-openmls profile, not local-backbone.
    or defer runner integration and improve sidecar setup/reset ergonomics first.
    or add dev-only bootstrap wrappers for identity/public-bundle/conversation-add-member/join.
    or document command migration/deprecation options for stub-era send/inbox.

### v0.4.x later polish

Likely needs:

    exact sidecar setup runbook.
    clearer dev state reset semantics.
    better command output shape.
    decision on whether stub-era send/inbox are aliased, replaced, or deprecated.
    decision on when local-backbone naming stops being premature.

---

## 11. v0.5.x and Later Carry-Forward

### v0.5.x

Local storage, trust state, provider state, vault design, and PQ/hybrid ciphersuite migration readiness.

Do not start PQ/hybrid implementation before the runtime OpenMLS/Cypher path exists and is documented. PQ/hybrid migration touches provider state, ciphersuite metadata, trust display, artifact size, reinit/rekey/recovery policy, and claim boundaries.

### v0.6.x+

Hostile-server and abuse-resistance harnesses.

### v0.7.x+

Deployability and operations hardening.

### v0.8.x+

Documented self-pentest / adversarial validation.

### v0.9.x+

Claim-boundary review.

### v0.10.x+

Android backend/app work.

### v1.x.x

Public app/server major epoch and much later CarbonStackOS/appliance research, if justified.

---

## 12. Lean Breakpoint Summary

    CarbonStack v0.4.8 is the behavior-preserving OpenMLS command/helper extraction checkpoint. v0.4.0 remains the current public broad local deployability pre-release. carbonstack remains at dc2d16c docs: assess pre-local-backbone runtime proof. carbonstack-comms is now at 0a48ae1 refactor: split OpenMLS runtime command helpers. v0.4.8 creates carbonstack-comms/internal/app/openmls_runtime.go and moves OpenMLS runtime command glue out of commands.go. openmls-send-dev and openmls-inbox-dev behavior remains unchanged. Validation passed for app tests, full Comms package tests, dev runtime smoke proof, local-cypher, doctor, and core --clean-generated, with final clean repo heads across all four repos. Existing send/inbox remain stub-era. This is not local-backbone, not mature messaging UX, and not production security proof. Next safe rung: v0.4.9 dev-runtime-openmls runner-profile recon from fresh package-like root.

---

## 13. Preserved Immediate Previous Handoff: v0.4.7

The following is the previous v0.4.7 handoff. Where it conflicts with the v0.4.8 overlay above, v0.4.8 wins for current state.


# CarbonStack LogDoc v0.4.7

**Last updated:** 2026-06-04 local session  
**Owner/operator:** bitcrusher32  
**Assistant/model:** GPT-5.5 Thinking  
**Current phase:** Phase 2F / **v0.4.x pre-local-backbone assessment checkpoint complete**. The v0.4.0 public broad local deployability pre-release remains the current public release. v0.4.5 proved the dev runtime OpenMLS CLI path; v0.4.6 recorded that smoke proof in the main `carbonstack` repo; v0.4.7 now records the assessment decision: do **not** promote the smoke proof to `local-backbone` yet, do **not** add a runner profile yet, keep the smoke script as a `carbonstack-comms` dev helper for now, prefer future `dev-runtime-openmls` naming if the proof becomes a runner profile, and plan behavior-preserving OpenMLS command/helper extraction before adding more runtime wrappers. `carbonstack` is now at `dc2d16c (HEAD -> main, origin/main, origin/HEAD) docs: assess pre-local-backbone runtime proof`; `carbonstack-comms` remains at `8e6e8b4 (HEAD -> main, origin/main, origin/HEAD) test: add dev OpenMLS runtime smoke proof`; `carbonstack-cypher` remains at `9ab994c (HEAD -> main, origin/main, origin/HEAD) docs: point to local operator runbook`; `carbonstack-os` remains at `1bbbe52 (HEAD -> main, origin/main, origin/HEAD) docs: clarify CarbonStackOS target direction`.

**Version schema:** `v[scope].[timeline]`. This file is `v0.4.7`, the pre-local-backbone assessment / validation-profile decision checkpoint after the v0.4.6 main-repo smoke-proof documentation alignment. Per LogDoc V2 practice, this Markdown file preserves current operational continuity and carries the v0.4.x runtime-integration thread forward; the JSON breakpoint remains the lean machine-readable handoff state.

**Compression note:** The detailed v0.3.0 through v0.3.36PRIME process ledger remains archived in PRIME/sanitized PRIME. v0.4.0 compressed that pre-release sprint; v0.4.1 preserved the runtime recon; v0.4.2 preserved the command-contract decision; v0.4.3 preserved the first dev-only OpenMLS send-side implementation; v0.4.4 preserved the first dev-only OpenMLS receive/open/optional-ack implementation; v0.4.5 preserved the first committed dev runtime smoke proof; v0.4.6 preserved project-level documentation/status alignment for that proof; v0.4.7 now preserves the decision boundary before any local-backbone/runner-profile promotion.

---

## 1. Project Goal

**Active goal:** Continue moving CarbonStackComms from stub-era runtime `send`/`inbox` behavior toward explicit dev-only OpenMLS/Cypher-backed runtime commands while preserving claim discipline. v0.4.7 completed the pre-local-backbone assessment and decided not to promote the v0.4.5 smoke proof to local-backbone or a main validation profile yet.

CarbonStack’s broader project intent remains:

- **CarbonStack** is the secure-communications appliance-stack project and shared doctrine/release authority.
- **CarbonStackComms** is the text-first messaging client and current OpenMLS sidecar/relay integration surface.
- **CarbonStackCypher** is the self-hostable relay/storage stack.
- **CarbonStackOS** is future constrained Android-derived appliance OS work and remains out of current runnable validation packages.
- The `carbonstack` Gitea repo remains the source of truth for release state, release assets, docs, and public claims.
- GitHub mirrors exist as secondary push mirrors only; Gitea remains authoritative.

**North star:**

> Its goal is not to be a better smartphone. Its goal is to stop being a smartphone.

**Core doctrine:**

> Every feature is guilty until it proves it does not add unacceptable parser, network, sensor, identity, or filesystem authority.

**Current reality after v0.4.7:** the dev runtime OpenMLS application-message smoke proof is implemented, recorded in main docs, and now assessed. The correct interpretation is still cautious:

    openmls-send-dev is a dev-only OpenMLS send-side command.
    openmls-inbox-dev is a dev-only OpenMLS receive/open/optional-ack command.
    scripts/dev-openmls-runtime-smoke.sh proves the dev CLI application-message path:
      openmls-send-dev -> Cypher -> openmls-inbox-dev --ack
    docs/157-dev-runtime-openmls-smoke-proof-v0.md records that proof.
    docs/158-pre-local-backbone-assessment-v0.md records that the proof should remain a comms-local dev helper for now.
    A future runner profile may be named dev-runtime-openmls, not local-backbone.
    The smoke script directly bootstraps sidecar identities, KeyPackage/Welcome, and conversation setup for dev purposes.
    Existing send/inbox remain stub-era.
    local-backbone remains reserved.

Hard nonclaims remain: v0.4.7 does **not** prove production readiness, production E2EE, hostile-server safety, metadata privacy, Android readiness, CarbonStackOS readiness, secure vault/storage, public ingress safety, real homelab validation, audit, certification, PQ/hybrid security, mature messenger UX, or local-backbone readiness.

---

## 2. Current Public Release State

Public release:

    https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.4.0

Release title:

    CarbonStack v0.4.0 Broad Local Deployability Pre-Release

Release tag:

    v0.4.0

Release commit:

    24ac9fc494 / 24ac9fc docs: record v0.4.0 package rehearsal

Gitea release status:

    Pre-Release

Release assets attached:

    carbonstack-v0.4.0-broad-local-deployability-pre-release.tgz
    carbonstack-v0.4.0-release-manifest.json
    carbonstack-v0.4.0-package-checksums.txt
    carbonstack-v0.4.0-asset-checksums.txt
    carbonstack-v0.4.0-validation-freeze.md
    v0.4.0-testing-runbook.md
    v0.4.0-release-notes.md
    LICENSE

Important release-page warning:

    The default Gitea Source Code ZIP/TAR.GZ downloads are auto-generated carbonstack-only archives.
    The intended runnable validation artifact is the attached multi-repo release package plus companion assets.

Post-release public surface:

    carbonstack README is evergreen.
    It points interested parties to the Releases page for runnable / known-good artifacts.
    It states Gitea is the source of truth.
    GitHub mirrors are secondary push mirrors, not release authority.

---

## 3. Current Repo Heads

    carbonstack        dc2d16c (HEAD -> main, origin/main, origin/HEAD) docs: assess pre-local-backbone runtime proof
    carbonstack-comms  8e6e8b4 (HEAD -> main, origin/main, origin/HEAD) test: add dev OpenMLS runtime smoke proof
    carbonstack-cypher 9ab994c (HEAD -> main, origin/main, origin/HEAD) docs: point to local operator runbook
    carbonstack-os     1bbbe52 (HEAD -> main, origin/main, origin/HEAD) docs: clarify CarbonStackOS target direction

Current public release tag:

    v0.4.0 -> 24ac9fc494 / 24ac9fc docs: record v0.4.0 package rehearsal

Post-release public-surface cleanup:

    6ba7aab docs: make front readme evergreen

v0.4.1 runtime recon commit:

    83d2f0f docs: record runtime OpenMLS Comms recon

v0.4.2 command contract commit:

    a799764 docs: define runtime OpenMLS command contract

v0.4.3 send implementation commit:

    39fb287 feat: add dev OpenMLS send command

v0.4.4 inbox implementation commit:

    65bc707 feat: add dev OpenMLS inbox command

v0.4.5 smoke proof commit:

    8e6e8b4 test: add dev OpenMLS runtime smoke proof

v0.4.6 main-doc alignment commit:

    3cd77a8 docs: record dev OpenMLS runtime smoke proof

Current v0.4.7 assessment commit:

    dc2d16c docs: assess pre-local-backbone runtime proof

Continuity note:

    `dc2d16c` is a `carbonstack` documentation/planning commit after the `v0.4.0` public release tag.
    It does not retag v0.4.0.
    The release remains attached to `24ac9fc494`.

---

## 4. Current Validated / Known State

Validated / known at the v0.4.7 checkpoint:

    [RELEASE] v0.4.0 Gitea release exists, is marked Pre-Release, and is attached to 24ac9fc494.
    [PACKAGE] v0.4.0 package includes carbonstack, carbonstack-comms, carbonstack-cypher, and release metadata.
    [PACKAGE] carbonstack-os is deliberately excluded from the runnable package.
    [RUNNER] full means release-snapshot followed by local-cypher.
    [RUNNER] release-snapshot validates package layout, metadata, strict pre-test artifact scan, package checksums, and core.
    [RUNNER] core validates doctor, OpenMLS real-Cypher lifecycle, Comms package tests, Cypher package tests, and artifact scans.
    [RUNNER] local-cypher validates Cypher-only lifecycle, restart/persistence against the same temporary DB, positive opaque OpenMLS envelope/ack flow, and historical negative stub-text/OpenMLS protocol pairing.
    [DOCS] main carbonstack README is evergreen and lists the dev runtime OpenMLS CLI smoke proof as demonstrated experimental validation coverage.
    [DOCS] docs/155-runtime-comms-openmls-cypher-recon-v0.md records the v0.4.1 runtime recon.
    [DOCS] docs/156-runtime-openmls-command-contract-v0.md records the v0.4.2 command contract.
    [DOCS] docs/157-dev-runtime-openmls-smoke-proof-v0.md records the v0.4.5 dev runtime smoke proof.
    [DOCS] docs/158-pre-local-backbone-assessment-v0.md records the v0.4.7 assessment and plan.
    [DOCS] docs/README.md indexes the v0.4.7 assessment doc and fixed the v0.4.5 bullet formatting.
    [DOCS] roadmap/ROADMAP.md records the v0.4.7 assessment and changes the stale v0.4.6 next-rung label to v0.4.7.
    [COMMS] openmls-send-dev exists in carbonstack-comms.
    [COMMS] openmls-inbox-dev exists in carbonstack-comms.
    [COMMS] scripts/dev-openmls-runtime-smoke.sh exists in carbonstack-comms.
    [COMMS] README documents the dev runtime OpenMLS smoke proof.
    [COMMS] existing send/inbox/ack remain in place and stub-era.
    [MIRRORS] all four repos have GitHub push-mirror functionality, with Gitea remaining source of truth.

v0.4.7 assessment result:

    local-backbone remains reserved.
    No runner profile was added in this rung.
    The smoke script should remain a carbonstack-comms dev helper for now.
    A future runner profile is plausible, with dev-runtime-openmls preferred as the non-overclaiming name.
    Runner promotion should wait until boundary docs, generated-state cleanup, and release-root behavior are explicit/tested.
    Helper extraction should be considered before adding more OpenMLS runtime wrappers.
    The next preferred rung is behavior-preserving OpenMLS command/helper extraction recon.
    Alternative next path is dev-runtime-openmls runner-profile recon from a fresh package-like root.

v0.4.7 validation result from the final log:

    local-cypher passed.
    doctor passed.
    core --clean-generated passed.
    pre-test artifact scan was clean.
    OpenMLS real-Cypher lifecycle passed during core.
    carbonstack-comms package tests passed during core.
    carbonstack-cypher package tests passed during core.
    post-test artifact scan found known OpenMLS generated roots.
    --clean-generated removed .carbonstack-openmls-sidecar-state and target.
    final repo status showed carbonstack at dc2d16c and all four repos clean.

Expected v0.4.0 public package validation commands remain:

    cd <package-root>/carbonstack/tools/carbonstack-validate
    go run . --profile verify-checksums --root <package-root>
    go run . --profile full --root <package-root> --clean-generated

Live umbrella validation commands remain:

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go run . --profile local-cypher
    go run . --profile doctor
    go run . --profile core --clean-generated

Current v0.4.5 smoke command remains:

    cd ~/repos/carbonstack_umbrella/carbonstack-comms
    scripts/dev-openmls-runtime-smoke.sh

---

## 5. v0.4.7 Work Completed

### 5.1 Pre-local-backbone assessment recon was performed

The v0.4.7 recon inspected:

    main carbonstack README claim/nonclaim surfaces
    docs/README.md current release and validation docs index
    docs/157-dev-runtime-openmls-smoke-proof-v0.md
    roadmap/ROADMAP.md v0.4.x section
    carbonstack validation runner architecture
    tools/carbonstack-validate profile dispatch
    tools/carbonstack-validate README profile boundaries
    local-cypher runner structure
    release-snapshot runner structure
    carbonstack-comms dev OpenMLS runtime smoke script
    carbonstack-comms README smoke/runtime command sections
    generated artifact ignore/cleanup policy
    carbonstack-comms commands.go size/function index
    app command tests
    state/trust path surfaces
    OpenMLS sidecar state/reset surfaces
    Cypher runtime/temp DB surfaces
    live smoke script behavior
    live validation baseline

Key findings:

    runner architecture is simple enough to add a future profile.
    current runner README is claim-careful and already warns that local-cypher is not local-backbone/runtime UX.
    full is already described as a validation ladder, not deployment.
    the smoke script is deterministic and self-contained enough to be a future runner target.
    the smoke script still depends on direct sidecar bootstrap for identity/KeyPackage/Welcome/conversation setup.
    local-backbone naming remains premature.
    docs needed a v0.4.6 -> v0.4.7 next-rung correction.
    commands.go has real extraction pressure before more OpenMLS runtime wrappers are added.

### 5.2 Assessment doc was added

New file:

    carbonstack/docs/158-pre-local-backbone-assessment-v0.md

This doc records:

    the v0.4.7 assessment after smoke proof documentation alignment.
    the decision not to add a runner profile yet.
    the decision not to use local-backbone naming yet.
    the recommended future runner profile name:
      dev-runtime-openmls
    the need to keep the smoke script as a comms-local dev helper for now.
    the state/reset cleanup considerations around sidecar generated roots.
    the recommendation to plan helper extraction before more runtime growth.
    the fact that KeyPackage/Welcome/bootstrap remain direct sidecar setup.
    the roadmap correction from v0.4.6 to v0.4.7 assessment.
    the recommended next rung:
      v0.4.8 behavior-preserving command/helper extraction recon

### 5.3 Docs index was updated

Updated file:

    carbonstack/docs/README.md

Changes:

    Fixed formatting for the v0.4.5 smoke-proof doc bullet so the path is backticked like adjacent entries.
    Added the v0.4.7 assessment doc:
      docs/158-pre-local-backbone-assessment-v0.md

### 5.4 Roadmap was updated

Updated file:

    carbonstack/roadmap/ROADMAP.md

Changes:

    Replaced stale next-rung wording:
      v0.4.6 pre-local-backbone assessment and validation-profile decision
    with:
      v0.4.7 pre-local-backbone assessment and validation-profile decision

    Added v0.4.7 result section:
      local-backbone remains reserved.
      dev-runtime-openmls is the preferred future runner-profile name if the smoke proof is promoted.
      the v0.4.5 smoke script should remain a carbonstack-comms dev helper until runner boundaries, generated-state cleanup, and release-root behavior are documented/tested.
      commands.go/helper extraction should be considered before adding more OpenMLS runtime wrappers.

    Recommended next rung:
      v0.4.8 behavior-preserving OpenMLS command/helper extraction recon, or dev-runtime-openmls runner-profile recon if runner promotion becomes the priority.

### 5.5 Validation passed

Commands observed passing:

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go run . --profile local-cypher
    go run . --profile doctor
    go run . --profile core --clean-generated

Final commit/push:

    carbonstack:
    dc2d16c docs: assess pre-local-backbone runtime proof

Final repo snapshot:

    carbonstack        dc2d16c docs: assess pre-local-backbone runtime proof
    carbonstack-comms  8e6e8b4 test: add dev OpenMLS runtime smoke proof
    carbonstack-cypher 9ab994c docs: point to local operator runbook
    carbonstack-os     1bbbe52 docs: clarify CarbonStackOS target direction

### 5.6 Blunder: directory typo / uncommitted state moment

During commit/push, the user accidentally ran:

    cd carcd ~/repos/carbonstack_umbrella/carbonstack

from `carbonstack-os`, causing:

    -bash: cd: too many arguments

Then `git add` ran from the wrong repo and failed:

    fatal: pathspec 'docs/158-pre-local-backbone-assessment-v0.md' did not match any files

A second typo occurred:

    cd carbonstacak

causing:

    No such file or directory

The user corrected into:

    cd ~/repos/carbonstack_umbrella/carbonstack

and successfully committed/pushed.

Lesson:

    The final repo snapshot is what matters.
    When shell context drifts into `carbonstack-os` or umbrella root, re-anchor with the full absolute path before staging/committing.

---

## 6. Critical Blunders / Lessons to Carry Forward

### 6.1 Wrong-directory commit attempt

The first commit attempt was run from the wrong place because of a malformed `cd` command. It produced harmless errors and no commit.

Correct practice:

    before staging docs in carbonstack, run:
      cd ~/repos/carbonstack_umbrella/carbonstack
      pwd
      git status --short

This is especially important after long multi-repo recon blocks that end in `carbonstack-os`.

### 6.2 Commit metadata warning remains a polish issue

The commit log again showed auto-generated local commit metadata:

    Committer: bitcrusher32 <▓▓>

This is not a functional validation failure, but it remains a professionalism/publishing hygiene issue to address eventually with `git config --global user.name` and `git config --global user.email` if desired.

### 6.3 Local-backbone remains tempting but still too strong

The smoke proof is real. The docs now acknowledge it. The assessment confirms that naming it local-backbone would still overclaim because sidecar bootstrap is direct dev setup and mature runtime UX does not exist.

Correct current phrase:

    dev runtime OpenMLS smoke proof

Possible future runner profile:

    dev-runtime-openmls

Still reserved:

    local-backbone

### 6.4 Helper extraction should precede more runtime wrappers

`commands.go` is still workable, but the direction is clear:

    old stub-era commands are still present.
    openmls-send-dev added sidecar protect glue.
    openmls-inbox-dev added sidecar open glue.
    future identity/public-bundle/conversation-add-member/join wrappers would make commands.go much larger.

Do not add more OpenMLS runtime wrappers blind. Recon and likely behavior-preserving extraction should happen first.

### 6.5 Runner profile can wait

A future `dev-runtime-openmls` profile is plausible, but premature for this exact checkpoint.

Reasons:

    smoke proof is still comms-local.
    direct sidecar bootstrap remains.
    release-root/fresh-package behavior has not been tested.
    generated-state cleanup needs explicit runner contract.
    profile naming must avoid overclaiming.

---

## 7. Current Blockers / Not Validated

Current blockers and nonvalidated areas after v0.4.7:

    Existing send/inbox remain stub-era.
    Runtime Comms has dev-only OpenMLS commands and a smoke proof, but no mature user-facing UX.
    Sidecar KeyPackage/Welcome/bootstrap setup is still direct dev setup, not Comms runtime UX.
    No Relay Space join model exists yet.
    local-backbone runner profile does not exist and remains reserved.
    dev-runtime-openmls runner profile does not exist yet.
    The dev runtime smoke script is not yet integrated into the main carbonstack validation runner.
    The smoke proof is not yet part of a public release package validation surface.
    Behavior-preserving helper extraction has not been performed.
    Fresh release-package behavior for the smoke script has not been tested as a runner target.
    Post-release public download verification is still useful if not already performed.
    Sanitized PRIME placement in carbonstack/logdoc-loop-system still needs a policy decision.
    Production-safe provider storage and secure vault design remain v0.5.x future work.
    PQ/hybrid ciphersuite migration readiness remains v0.5.x future work.
    Hostile-server rollback/replay/metadata-abuse validation remains future work.
    Android/CarbonStackOS implementation remains later work.
    CarbonStack Relay Space mechanics remain a future concept, not an implementation.
    Public ingress, systemd, cloudflared, real homelab deployment, and production service paths remain out of scope.

Hard nonclaims remain:

    Do not claim production readiness.
    Do not claim production E2EE.
    Do not claim hostile-server safety.
    Do not claim metadata privacy.
    Do not claim Debian production deployability.
    Do not claim systemd readiness.
    Do not claim cloudflared readiness.
    Do not claim Android readiness.
    Do not claim CarbonStackOS readiness.
    Do not claim secure vault/storage.
    Do not claim real homelab validation.
    Do not claim external audit or certification.
    Do not claim CarbonStack Relay Space mechanics are implemented.
    Do not claim local-backbone exists or is justified yet.
    Do not claim local-cypher validates runtime Comms UX.
    Do not claim full validates local-backbone.
    Do not claim full is a deployment command.
    Do not claim send/inbox are OpenMLS-backed yet.
    Do not claim PQ/hybrid security or quantum-safe messaging.
    Do not let GitHub mirrors supersede Gitea source-of-truth status.

Allowed claim:

    CarbonStackComms now has dev-only openmls-send-dev and openmls-inbox-dev commands.
    CarbonStackComms now has a dev runtime smoke script proving openmls-send-dev -> Cypher -> openmls-inbox-dev --ack.
    The main carbonstack repo now records that smoke proof and the pre-local-backbone assessment.
    A future dev-runtime-openmls runner profile is plausible, but not implemented.
    The smoke script starts a temporary local Cypher server, uses temporary Comms state files, creates dev-local OpenMLS sidecar identities/conversation, sends an OpenMLS application message through the Comms CLI, opens it through the Comms CLI, acks after sidecar success, and verifies Bob's inbox is empty after ack.
    Existing send/inbox remain stub-era.
    v0.4.0 remains the current public broad local deployability pre-release.
    Gitea remains source of truth; GitHub mirrors are secondary push mirrors.

---

## 8. Next Safest Actions

Recommended next rung:

    v0.4.8 behavior-preserving OpenMLS command/helper extraction recon

Focus:

    inspect internal/app command boundaries.
    identify exactly what can move out of commands.go without behavior changes.
    keep CLI registration and old commands stable.
    keep openmls-send-dev and openmls-inbox-dev outputs stable.
    keep package-level test seams or replace them with equally testable structure.
    validate app tests, full comms tests, smoke script, and runner core afterward.

Likely implementation target after recon:

    carbonstack-comms/internal/app/openmls_runtime.go
    or carbonstack-comms/internal/app/openmls_commands.go

Possible alternative next rung:

    v0.4.8 dev-runtime-openmls runner-profile recon from a fresh package-like root

Preferred order:

    extraction recon first.
    behavior-preserving extraction second.
    runner-profile implementation later.

Rationale:

    more runtime work and future sidecar bootstrap wrappers will become safer if OpenMLS command glue is organized before profile promotion.

Avoid next:

    replacing send/inbox immediately.
    adding local-backbone prematurely.
    adding a runner profile named local-backbone.
    moving smoke proof into public release validation without fresh-root behavior checks.
    public ingress.
    cloudflared.
    systemd.
    real homelab deployment.
    Android app.
    CarbonStackOS implementation.
    CarbonStack Relay Space implementation.
    PQ/hybrid ciphersuite implementation.
    production/security claims.

---

## 9. Critical Paths / Functions

### Source-of-truth public surfaces

    [GITEA SOURCE OF TRUTH] https://git.bitcrusher32.win/bitcrusher32/carbonstack
    [RELEASES] https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases
    [CURRENT PUBLIC RELEASE] https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.4.0

GitHub mirrors:

    secondary push mirrors for discoverability/redundancy only.
    Do not treat mirrors as release authority unless project policy changes.

### WSL Debian active working umbrella

    [WSL-DEBIAN] umbrella: ~/repos/carbonstack_umbrella
    [WSL-DEBIAN] carbonstack: ~/repos/carbonstack_umbrella/carbonstack
    [WSL-DEBIAN] carbonstack-comms: ~/repos/carbonstack_umbrella/carbonstack-comms
    [WSL-DEBIAN] carbonstack-cypher: ~/repos/carbonstack_umbrella/carbonstack-cypher
    [WSL-DEBIAN] carbonstack-os: ~/repos/carbonstack_umbrella/carbonstack-os
    [GO-RUNNER] runner dir: ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate

### Runtime integration doc surfaces

    [DOC] v0.4.1 runtime recon:
    carbonstack/docs/155-runtime-comms-openmls-cypher-recon-v0.md

    [DOC] v0.4.2 command contract:
    carbonstack/docs/156-runtime-openmls-command-contract-v0.md

    [DOC] v0.4.6 smoke proof record:
    carbonstack/docs/157-dev-runtime-openmls-smoke-proof-v0.md

    [DOC] v0.4.7 pre-local-backbone assessment:
    carbonstack/docs/158-pre-local-backbone-assessment-v0.md

    [DOC] docs index:
    carbonstack/docs/README.md

    [DOC] roadmap:
    carbonstack/roadmap/ROADMAP.md

    [DOC] top README:
    carbonstack/README.md

### Comms runtime surfaces

    [COMMS] CLI entry:
    carbonstack-comms/cmd/comms/main.go

    [COMMS] command implementation:
    carbonstack-comms/internal/app/commands.go

    [COMMS] likely future extraction target:
    carbonstack-comms/internal/app/openmls_runtime.go
    or
    carbonstack-comms/internal/app/openmls_commands.go

    [COMMS] openmls-send-dev tests:
    carbonstack-comms/internal/app/openmls_send_dev_test.go

    [COMMS] openmls-inbox-dev tests:
    carbonstack-comms/internal/app/openmls_inbox_dev_test.go

    [COMMS] dev runtime smoke proof:
    carbonstack-comms/scripts/dev-openmls-runtime-smoke.sh

    [COMMS] Cypher client:
    carbonstack-comms/internal/client/cypher.go

    [COMMS] relay helpers:
    carbonstack-comms/internal/relay/cypher_bridge.go
    carbonstack-comms/internal/relay/openmls_artifacts.go

    [COMMS] OpenMLS sidecar dir:
    carbonstack-comms/internal/protocol/mls/openmls-sidecar

    [COMMS] sidecar README:
    carbonstack-comms/internal/protocol/mls/openmls-sidecar/README.md

### Current stub-era Comms functions / behaviors

    send:
      trust policy evaluation
      mock/stub provider encryption
      Cypher envelope submit

    inbox:
      Cypher inbox retrieval
      mock/stub provider plaintext display as stub_plaintext

    ack:
      envelope ack by envelope ID and recipient device ID

### Implemented v0.4.x dev OpenMLS commands

    openmls-send-dev:
      state.RequireReadyDevice
      trust.EvaluateSend
      runOpenMLSMessageProtect
      sidecar message-protect
      internal/relay.SubmitOpenMLSArtifactEnvelope
      ArtifactKindApplicationMessage
      dev/pre-alpha warning

    openmls-inbox-dev:
      state.RequireReadyDevice
      Cypher inbox
      skip unsupported envelopes
      internal/relay.WriteOpenMLSArtifactFromEnvelope
      runOpenMLSMessageOpen
      sidecar message-open
      print plaintext after sidecar success
      optional ack only after sidecar success

### v0.4.5 runtime smoke proof function

    scripts/dev-openmls-runtime-smoke.sh:
      temp workspace
      temp Cypher binary
      temp Cypher DB
      dynamic loopback Cypher port
      temp Alice/Bob Comms state
      sidecar state reset
      sidecar identity/conversation bootstrap
      openmls-send-dev application-message send
      openmls-inbox-dev application-message open + ack
      inbox-empty-after-ack verification
      robust cleanup with INT -> TERM -> KILL escalation

### OpenMLS sidecar commands relevant to v0.4.x

    provider-info
    identity-create
    identity-status
    public-bundle-export
    conversation-create
    conversation-load-check
    conversation-add-member
    conversation-join
    message-protect
    message-open

### Relay/API functions relevant to runtime integration

    internal/client.CypherClient.SubmitEnvelope
    internal/client.CypherClient.Inbox
    internal/client.CypherClient.AckEnvelope
    internal/relay.SubmitOpenMLSArtifactEnvelope
    internal/relay.WriteOpenMLSArtifactFromEnvelope

### Existing v0.4.0 package validation

    cd <package-root>/carbonstack/tools/carbonstack-validate
    go run . --profile verify-checksums --root <package-root>
    go run . --profile full --root <package-root> --clean-generated

### Live umbrella validation

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go run . --profile local-cypher
    go run . --profile doctor
    go run . --profile core --clean-generated

### Envelope/protocol constants

Accepted opaque OpenMLS application-message pair:

    content_type=carbonstack.mls.application-message.v0
    protocol_version=carbonstack-openmls-sidecar-v0

Historical invalid pair validated by local-cypher:

    content_type=carbonstack.message.text.stub.v0
    protocol_version=carbonstack-openmls-sidecar-v0

Expected result:

    HTTP 400
    error code unsupported_protocol_version

---

## 10. v0.4.x Forward Plan

### v0.4.8 — behavior-preserving OpenMLS command/helper extraction recon

Expected:

    inspect commands.go and app tests.
    choose a safe extraction target.
    avoid behavior changes.
    preserve existing command output.
    preserve openmls-send-dev and openmls-inbox-dev tests.
    run smoke + validation after extraction if implemented.

### v0.4.9+ — possible extraction implementation / runner-profile recon

Possible paths:

    perform behavior-preserving extraction.
    then recon a dev-runtime-openmls runner profile from fresh package-like roots.
    or improve sidecar setup/reset ergonomics first.
    or document command migration/deprecation options for stub-era send/inbox.

### v0.4.x later polish

Likely needs:

    exact sidecar setup runbook.
    clearer dev state reset semantics.
    better command output shape.
    possible helper extraction if commands.go becomes too large.
    decision on whether stub-era send/inbox are aliased, replaced, or deprecated.
    decision on when local-backbone naming stops being premature.

---

## 11. v0.5.x and Later Carry-Forward

### v0.5.x

Local storage, trust state, provider state, vault design, and PQ/hybrid ciphersuite migration readiness.

Do not start PQ/hybrid implementation before the runtime OpenMLS/Cypher path exists and is documented. PQ/hybrid migration touches provider state, ciphersuite metadata, trust display, artifact size, reinit/rekey/recovery policy, and claim boundaries.

### v0.6.x+

Hostile-server and abuse-resistance harnesses.

### v0.7.x+

Deployability and operations hardening.

### v0.8.x+

Documented self-pentest / adversarial validation.

### v0.9.x+

Claim-boundary review.

### v0.10.x+

Android backend/app work.

### v1.x.x

Public app/server major epoch and much later CarbonStackOS/appliance research, if justified.

---

## 12. Lean Breakpoint Summary

    CarbonStack v0.4.7 is the pre-local-backbone assessment / validation-profile decision checkpoint. v0.4.0 remains the current public broad local deployability pre-release. carbonstack is now at dc2d16c docs: assess pre-local-backbone runtime proof. carbonstack-comms remains at 8e6e8b4 test: add dev OpenMLS runtime smoke proof. v0.4.7 adds carbonstack/docs/158-pre-local-backbone-assessment-v0.md, updates docs/README.md, and updates roadmap/ROADMAP.md. The assessment keeps local-backbone reserved, keeps the smoke proof comms-local for now, names dev-runtime-openmls as a plausible future runner profile, and recommends helper extraction before more runtime growth. Validation passed for local-cypher, doctor, and core --clean-generated, with final clean repo heads across all four repos. Existing send/inbox remain stub-era. This is not local-backbone, not mature messaging UX, and not production security proof. Next safe rung: v0.4.8 behavior-preserving OpenMLS command/helper extraction recon.

---

## 13. Preserved Immediate Previous Handoff: v0.4.6

The following is the previous v0.4.6 handoff. Where it conflicts with the v0.4.7 overlay above, v0.4.7 wins for current state.


# CarbonStack LogDoc v0.4.6

**Last updated:** 2026-06-04 local session  
**Owner/operator:** bitcrusher32  
**Assistant/model:** GPT-5.5 Thinking  
**Current phase:** Phase 2F / **v0.4.x smoke proof documentation and main-repo status alignment checkpoint complete**. The v0.4.0 public broad local deployability pre-release remains the current public release. v0.4.3 implemented `openmls-send-dev`; v0.4.4 implemented `openmls-inbox-dev`; v0.4.5 committed and validated `scripts/dev-openmls-runtime-smoke.sh` in `carbonstack-comms`; v0.4.6 now records that smoke proof in the main `carbonstack` repo through `docs/157-dev-runtime-openmls-smoke-proof-v0.md`, updates the docs index, updates the roadmap, and lightly aligns the top README. `carbonstack` is now at `3cd77a8 (HEAD -> main, origin/main, origin/HEAD) docs: record dev OpenMLS runtime smoke proof`; `carbonstack-comms` remains at `8e6e8b4 (HEAD -> main, origin/main, origin/HEAD) test: add dev OpenMLS runtime smoke proof`; `carbonstack-cypher` remains at `9ab994c (HEAD -> main, origin/main, origin/HEAD) docs: point to local operator runbook`; `carbonstack-os` remains at `1bbbe52 (HEAD -> main, origin/main, origin/HEAD) docs: clarify CarbonStackOS target direction`.

**Version schema:** `v[scope].[timeline]`. This file is `v0.4.6`, the main-repo documentation/status alignment checkpoint after the v0.4.5 dev runtime smoke proof. Per LogDoc V2 practice, this Markdown file preserves current operational continuity and carries the v0.4.x runtime-integration thread forward; the JSON breakpoint remains the lean machine-readable handoff state.

**Compression note:** The detailed v0.3.0 through v0.3.36PRIME process ledger remains archived in PRIME/sanitized PRIME. v0.4.0 compressed that pre-release sprint; v0.4.1 preserved the runtime recon; v0.4.2 preserved the command-contract decision; v0.4.3 preserved the first dev-only OpenMLS send-side implementation; v0.4.4 preserved the first dev-only OpenMLS receive/open/optional-ack implementation; v0.4.5 preserved the first committed dev runtime smoke proof; v0.4.6 now preserves project-level documentation/status alignment for that proof. Future work should continue from this v0.4.x runtime thread.

---

## 1. Project Goal

**Active goal:** Continue moving CarbonStackComms from stub-era runtime `send`/`inbox` behavior toward explicit dev-only OpenMLS/Cypher-backed runtime commands and project-level validation/documentation maturity. v0.4.6 completed the main `carbonstack` documentation/status alignment after the first committed dev runtime smoke proof.

CarbonStack’s broader project intent remains:

- **CarbonStack** is the secure-communications appliance-stack project and shared doctrine/release authority.
- **CarbonStackComms** is the text-first messaging client and current OpenMLS sidecar/relay integration surface.
- **CarbonStackCypher** is the self-hostable relay/storage stack.
- **CarbonStackOS** is future constrained Android-derived appliance OS work and remains out of current runnable validation packages.
- The `carbonstack` Gitea repo remains the source of truth for release state, release assets, docs, and public claims.
- GitHub mirrors exist as secondary push mirrors only; Gitea remains authoritative.

**North star:**

> Its goal is not to be a better smartphone. Its goal is to stop being a smartphone.

**Core doctrine:**

> Every feature is guilty until it proves it does not add unacceptable parser, network, sensor, identity, or filesystem authority.

**Current reality after v0.4.6:** the dev runtime OpenMLS application-message smoke proof is now both implemented in `carbonstack-comms` and recorded in the main `carbonstack` project docs. The old stub-era `send`, `inbox`, and `ack` commands still remain in place and must not be described as the mature OpenMLS-backed runtime UX.

Correct interpretation:

    openmls-send-dev is a dev-only OpenMLS send-side command.
    openmls-inbox-dev is a dev-only OpenMLS receive/open/optional-ack command.
    scripts/dev-openmls-runtime-smoke.sh proves the dev CLI application-message path:
      openmls-send-dev -> Cypher -> openmls-inbox-dev --ack
    docs/157-dev-runtime-openmls-smoke-proof-v0.md records that proof at the main carbonstack level.
    The smoke script directly bootstraps sidecar identities, KeyPackage/Welcome, and conversation setup for dev purposes.
    The application-message path is the runtime CLI proof target.
    Existing send/inbox remain stub-era.
    local-backbone remains reserved until whole-path validation is integrated/documented enough to deserve that name.

Hard nonclaims remain: v0.4.6 does **not** prove production readiness, production E2EE, hostile-server safety, metadata privacy, Android readiness, CarbonStackOS readiness, secure vault/storage, public ingress safety, real homelab validation, audit, certification, PQ/hybrid security, mature messenger UX, or local-backbone readiness.

---

## 2. Current Public Release State

Public release:

    https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.4.0

Release title:

    CarbonStack v0.4.0 Broad Local Deployability Pre-Release

Release tag:

    v0.4.0

Release commit:

    24ac9fc494 / 24ac9fc docs: record v0.4.0 package rehearsal

Gitea release status:

    Pre-Release

Release assets attached:

    carbonstack-v0.4.0-broad-local-deployability-pre-release.tgz
    carbonstack-v0.4.0-release-manifest.json
    carbonstack-v0.4.0-package-checksums.txt
    carbonstack-v0.4.0-asset-checksums.txt
    carbonstack-v0.4.0-validation-freeze.md
    v0.4.0-testing-runbook.md
    v0.4.0-release-notes.md
    LICENSE

Important release-page warning:

    The default Gitea Source Code ZIP/TAR.GZ downloads are auto-generated carbonstack-only archives.
    The intended runnable validation artifact is the attached multi-repo release package plus companion assets.

Post-release public surface:

    carbonstack README is evergreen.
    It points interested parties to the Releases page for runnable / known-good artifacts.
    It states Gitea is the source of truth.
    GitHub mirrors are secondary push mirrors, not release authority.

---

## 3. Current Repo Heads

    carbonstack        3cd77a8 (HEAD -> main, origin/main, origin/HEAD) docs: record dev OpenMLS runtime smoke proof
    carbonstack-comms  8e6e8b4 (HEAD -> main, origin/main, origin/HEAD) test: add dev OpenMLS runtime smoke proof
    carbonstack-cypher 9ab994c (HEAD -> main, origin/main, origin/HEAD) docs: point to local operator runbook
    carbonstack-os     1bbbe52 (HEAD -> main, origin/main, origin/HEAD) docs: clarify CarbonStackOS target direction

Current public release tag:

    v0.4.0 -> 24ac9fc494 / 24ac9fc docs: record v0.4.0 package rehearsal

Post-release public-surface cleanup:

    6ba7aab docs: make front readme evergreen

v0.4.1 runtime recon commit:

    83d2f0f docs: record runtime OpenMLS Comms recon

v0.4.2 command contract commit:

    a799764 docs: define runtime OpenMLS command contract

v0.4.3 send implementation commit:

    39fb287 feat: add dev OpenMLS send command

v0.4.4 inbox implementation commit:

    65bc707 feat: add dev OpenMLS inbox command

v0.4.5 smoke proof commit:

    8e6e8b4 test: add dev OpenMLS runtime smoke proof

Current v0.4.6 main-doc alignment commit:

    3cd77a8 docs: record dev OpenMLS runtime smoke proof

Continuity note:

    `3cd77a8` is a `carbonstack` documentation/status-alignment commit after the `v0.4.0` public release tag.
    It does not retag v0.4.0.
    The release remains attached to `24ac9fc494`.

---

## 4. Current Validated / Known State

Validated / known at the v0.4.6 checkpoint:

    [RELEASE] v0.4.0 Gitea release exists, is marked Pre-Release, and is attached to 24ac9fc494.
    [PACKAGE] v0.4.0 package includes carbonstack, carbonstack-comms, carbonstack-cypher, and release metadata.
    [PACKAGE] carbonstack-os is deliberately excluded from the runnable package.
    [RUNNER] full means release-snapshot followed by local-cypher.
    [RUNNER] release-snapshot validates package layout, metadata, strict pre-test artifact scan, package checksums, and core.
    [RUNNER] core validates doctor, OpenMLS real-Cypher lifecycle, Comms package tests, Cypher package tests, and artifact scans.
    [RUNNER] local-cypher validates Cypher-only lifecycle, restart/persistence against the same temporary DB, positive opaque OpenMLS envelope/ack flow, and historical negative stub-text/OpenMLS protocol pairing.
    [DOCS] main carbonstack README is evergreen and now lists the dev runtime OpenMLS CLI smoke proof as demonstrated experimental validation coverage.
    [DOCS] docs/155-runtime-comms-openmls-cypher-recon-v0.md records the v0.4.1 runtime recon.
    [DOCS] docs/156-runtime-openmls-command-contract-v0.md records the v0.4.2 command contract.
    [DOCS] docs/157-dev-runtime-openmls-smoke-proof-v0.md records the v0.4.5 dev runtime smoke proof.
    [DOCS] docs/README.md indexes the v0.4.5 smoke proof doc.
    [DOCS] roadmap/ROADMAP.md records the v0.4.5 result and marks v0.4.6+ as pre-local-backbone assessment / validation-profile decision territory.
    [COMMS] openmls-send-dev exists in carbonstack-comms.
    [COMMS] openmls-inbox-dev exists in carbonstack-comms.
    [COMMS] scripts/dev-openmls-runtime-smoke.sh exists in carbonstack-comms.
    [COMMS] README documents the dev runtime OpenMLS smoke proof.
    [COMMS] existing send/inbox/ack remain in place and stub-era.
    [MIRRORS] all four repos have GitHub push-mirror functionality, with Gitea remaining source of truth.

v0.4.6 documentation/status-alignment result from the final log:

    README.md now includes a demonstrated validation bullet for:
      a dev runtime OpenMLS CLI smoke proof for openmls-send-dev -> Cypher -> openmls-inbox-dev --ack
    README.md changed "polished Comms runtime send/inbox UX" to "mature Comms runtime send/inbox UX" in the nonclaims list.
    docs/157-dev-runtime-openmls-smoke-proof-v0.md was created.
    docs/README.md now indexes docs/157-dev-runtime-openmls-smoke-proof-v0.md.
    roadmap/ROADMAP.md now records the v0.4.5 smoke proof result and recommends v0.4.6 pre-local-backbone assessment / validation-profile decision.
    local-cypher passed.
    doctor passed.
    core --clean-generated passed.
    pre/post local-cypher artifact scans were clean in this rung.
    final repo status showed carbonstack at 3cd77a8 and all four repos clean.

Expected v0.4.0 public package validation commands remain:

    cd <package-root>/carbonstack/tools/carbonstack-validate
    go run . --profile verify-checksums --root <package-root>
    go run . --profile full --root <package-root> --clean-generated

Live umbrella validation commands remain:

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go run . --profile local-cypher
    go run . --profile doctor
    go run . --profile core --clean-generated

Current v0.4.5 smoke command remains:

    cd ~/repos/carbonstack_umbrella/carbonstack-comms
    scripts/dev-openmls-runtime-smoke.sh

---

## 5. v0.4.6 Work Completed

### 5.1 Main carbonstack smoke-proof doc was added

New file:

    carbonstack/docs/157-dev-runtime-openmls-smoke-proof-v0.md

This doc records:

    the v0.4.5 dev runtime smoke proof
    the command path:
      openmls-send-dev -> Cypher -> openmls-inbox-dev --ack
    the temporary Cypher server / temporary DB / temporary Comms state design
    direct dev sidecar setup for identity, KeyPackage, Welcome, and conversation bootstrap
    runtime CLI proof target as the OpenMLS application-message path
    plaintext match
    ack only after sidecar success
    Bob inbox empty after ack
    the initial cleanup/self-termination blunder and fix
    validation result
    current allowed claims
    why this is not local-backbone yet
    suggested next rung: pre-local-backbone assessment

### 5.2 Docs index was updated

Updated file:

    carbonstack/docs/README.md

Added index entry:

    docs/157-dev-runtime-openmls-smoke-proof-v0.md — records the v0.4.5 dev runtime OpenMLS CLI smoke proof using openmls-send-dev -> Cypher -> openmls-inbox-dev --ack.

Minor formatting note:

    The added bullet does not use backticks around the path, unlike adjacent entries. This is not functionally harmful, but may be a small docs polish item later.

### 5.3 Roadmap was updated

Updated file:

    carbonstack/roadmap/ROADMAP.md

Added v0.4.5 result section:

    carbonstack-comms now includes scripts/dev-openmls-runtime-smoke.sh.
    The script proves openmls-send-dev -> Cypher -> openmls-inbox-dev --ack for the OpenMLS application-message path.
    It verifies plaintext and confirms the recipient inbox is empty after ack.

Boundary preserved:

    This is a dev/pre-alpha smoke proof.
    This is not local-backbone.
    This is not production messaging UX.
    Sidecar KeyPackage/Welcome/bootstrap setup remains direct dev setup.
    Existing send/inbox remain stub-era.

Recommended next rung in roadmap:

    v0.4.6 pre-local-backbone assessment and validation-profile decision.

Note:

    Because the current breakpoint is v0.4.6 and the roadmap entry still phrases the next rung as v0.4.6, the next actual work should treat that as "v0.4.7 pre-local-backbone assessment / validation-profile decision" or update the roadmap wording in the next docs polish.

### 5.4 Top README was lightly aligned

Updated file:

    carbonstack/README.md

Changes:

    The demonstrated validation coverage list now includes:
      a dev runtime OpenMLS CLI smoke proof for openmls-send-dev -> Cypher -> openmls-inbox-dev --ack

    The nonclaims list now says:
      mature Comms runtime send/inbox UX

    instead of:
      polished Comms runtime send/inbox UX

Rationale:

    "mature" better matches the fact that dev runtime commands and a smoke proof now exist, while the old send/inbox UX remains stub-era and no mature user-facing runtime UX exists yet.

### 5.5 Validation passed

Commands observed passing:

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go run . --profile local-cypher
    go run . --profile doctor
    go run . --profile core --clean-generated

Final commit/push:

    carbonstack:
    3cd77a8 docs: record dev OpenMLS runtime smoke proof

Final repo snapshot:

    carbonstack        3cd77a8 docs: record dev OpenMLS runtime smoke proof
    carbonstack-comms  8e6e8b4 test: add dev OpenMLS runtime smoke proof
    carbonstack-cypher 9ab994c docs: point to local operator runbook
    carbonstack-os     1bbbe52 docs: clarify CarbonStackOS target direction

---

## 6. Critical Blunders / Lessons to Carry Forward

### 6.1 Push authentication failed once, then succeeded

The first `git push` after the `3cd77a8` commit failed with:

    remote: Failed to authenticate user
    fatal: Authentication failed for 'https://git.bitcrusher32.win/bitcrusher32/carbonstack.git/'

A follow-up `git push` succeeded and advanced `origin/main` from `a799764` to `3cd77a8`.

Lesson:

    Keep treating a failed first push as a transport/auth issue unless git state suggests otherwise.
    Always verify with the final repo-head/status snapshot.

### 6.2 Commit metadata warning remains a polish issue

The commit log again showed auto-generated local commit metadata:

    Committer: bitcrusher32 <▓▓>

This is not a functional validation failure, but it remains a professionalism/publishing hygiene issue to address eventually with `git config --global user.name` and `git config --global user.email` if desired.

### 6.3 Roadmap next-rung wording is now slightly stale

The roadmap entry created during this rung says:

    Recommended next rung:
      v0.4.6 pre-local-backbone assessment and validation-profile decision.

Because this LogDoc now marks the docs/status-alignment rung as v0.4.6, the next actual rung should either:

    interpret that roadmap line as v0.4.7 pre-local-backbone assessment / validation-profile decision, or
    do a small roadmap polish later to update the version label.

This is minor and not worth hotfixing unless the next rung touches the roadmap anyway.

### 6.4 Main docs now allow the smoke-proof claim, but not local-backbone

The top README now truthfully says CarbonStack has demonstrated a dev runtime OpenMLS CLI smoke proof.

Do not let that drift into stronger claims:

    not local-backbone
    not mature runtime send/inbox UX
    not production messaging
    not release-package public validation yet
    not Relay Space join mechanics

### 6.5 The docs index bullet formatting is slightly inconsistent

`docs/README.md` added the `docs/157...` bullet without backticks around the path.

This is not a functional issue. It is only a polish issue and can be cleaned up if the docs index is touched again.

---

## 7. Current Blockers / Not Validated

Current blockers and nonvalidated areas after v0.4.6:

    Existing send/inbox remain stub-era.
    Runtime Comms has dev-only OpenMLS commands and a smoke proof, but no mature user-facing UX.
    Sidecar KeyPackage/Welcome/bootstrap setup is still direct dev setup, not Comms runtime UX.
    No Relay Space join model exists yet.
    local-backbone runner profile does not exist and remains reserved.
    The dev runtime smoke script is not yet integrated into the main carbonstack validation runner.
    The smoke proof is not yet part of a public release package validation surface.
    Post-release public download verification is still useful if not already performed.
    Sanitized PRIME placement in carbonstack/logdoc-loop-system still needs a policy decision.
    Production-safe provider storage and secure vault design remain v0.5.x future work.
    PQ/hybrid ciphersuite migration readiness remains v0.5.x future work.
    Hostile-server rollback/replay/metadata-abuse validation remains future work.
    Android/CarbonStackOS implementation remains later work.
    CarbonStack Relay Space mechanics remain a future concept, not an implementation.
    Public ingress, systemd, cloudflared, real homelab deployment, and production service paths remain out of scope.

Hard nonclaims remain:

    Do not claim production readiness.
    Do not claim production E2EE.
    Do not claim hostile-server safety.
    Do not claim metadata privacy.
    Do not claim Debian production deployability.
    Do not claim systemd readiness.
    Do not claim cloudflared readiness.
    Do not claim Android readiness.
    Do not claim CarbonStackOS readiness.
    Do not claim secure vault/storage.
    Do not claim real homelab validation.
    Do not claim external audit or certification.
    Do not claim CarbonStack Relay Space mechanics are implemented.
    Do not claim local-backbone exists or is justified yet.
    Do not claim local-cypher validates runtime Comms UX.
    Do not claim full validates local-backbone.
    Do not claim full is a deployment command.
    Do not claim send/inbox are OpenMLS-backed yet.
    Do not claim PQ/hybrid security or quantum-safe messaging.
    Do not let GitHub mirrors supersede Gitea source-of-truth status.

Allowed claim:

    CarbonStackComms now has dev-only openmls-send-dev and openmls-inbox-dev commands.
    CarbonStackComms now has a dev runtime smoke script proving openmls-send-dev -> Cypher -> openmls-inbox-dev --ack.
    The main carbonstack repo now records that smoke proof in docs/157-dev-runtime-openmls-smoke-proof-v0.md, docs/README.md, roadmap/ROADMAP.md, and README.md.
    The smoke script starts a temporary local Cypher server, uses temporary Comms state files, creates dev-local OpenMLS sidecar identities/conversation, sends an OpenMLS application message through the Comms CLI, opens it through the Comms CLI, acks after sidecar success, and verifies Bob's inbox is empty after ack.
    Existing send/inbox remain stub-era.
    v0.4.0 remains the current public broad local deployability pre-release.
    Gitea remains source of truth; GitHub mirrors are secondary push mirrors.

---

## 8. Next Safest Actions

Recommended next rung:

    v0.4.7 pre-local-backbone assessment / validation-profile decision

Focus:

    decide whether the dev runtime smoke proof should remain a comms-local helper for now or become a carbonstack runner profile later.
    assess whether local-backbone naming is still premature.
    inspect whether sidecar bootstrap should be wrapped before runner integration.
    decide whether a runner profile would overclaim maturity.
    document reset/state/generated-root semantics.
    clean minor docs wording if already touching roadmap/docs index.

Likely outcome:

    Probably recon/docs first, not immediate runner integration.
    Runner integration may be useful later, but only after boundary language and cleanup semantics are extremely clear.

Possible follow-up after assessment:

    v0.4.8 runner-profile implementation if justified.
    v0.4.8 or v0.4.9 command ergonomics/setup simplification if runner integration is premature.
    v0.4.x late decision on whether stub-era send/inbox are aliased, replaced, or explicitly deprecated.
    v0.5.x state/trust/vault/PQ planning once v0.4.x runtime proof is sufficiently documented and stable.

Avoid next:

    replacing send/inbox immediately.
    adding local-backbone prematurely.
    moving smoke proof into public release validation without a boundary doc.
    public ingress.
    cloudflared.
    systemd.
    real homelab deployment.
    Android app.
    CarbonStackOS implementation.
    CarbonStack Relay Space implementation.
    PQ/hybrid ciphersuite implementation.
    production/security claims.
    broad negative-path suite unless the smoke proof exposes a direct blocker.

---

## 9. Critical Paths / Functions

### Source-of-truth public surfaces

    [GITEA SOURCE OF TRUTH] https://git.bitcrusher32.win/bitcrusher32/carbonstack
    [RELEASES] https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases
    [CURRENT PUBLIC RELEASE] https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.4.0

GitHub mirrors:

    secondary push mirrors for discoverability/redundancy only.
    Do not treat mirrors as release authority unless project policy changes.

### WSL Debian active working umbrella

    [WSL-DEBIAN] umbrella: ~/repos/carbonstack_umbrella
    [WSL-DEBIAN] carbonstack: ~/repos/carbonstack_umbrella/carbonstack
    [WSL-DEBIAN] carbonstack-comms: ~/repos/carbonstack_umbrella/carbonstack-comms
    [WSL-DEBIAN] carbonstack-cypher: ~/repos/carbonstack_umbrella/carbonstack-cypher
    [WSL-DEBIAN] carbonstack-os: ~/repos/carbonstack_umbrella/carbonstack-os
    [GO-RUNNER] runner dir: ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate

### Runtime integration doc surfaces

    [DOC] v0.4.1 runtime recon:
    carbonstack/docs/155-runtime-comms-openmls-cypher-recon-v0.md

    [DOC] v0.4.2 command contract:
    carbonstack/docs/156-runtime-openmls-command-contract-v0.md

    [DOC] v0.4.6 smoke proof record:
    carbonstack/docs/157-dev-runtime-openmls-smoke-proof-v0.md

    [DOC] docs index:
    carbonstack/docs/README.md

    [DOC] roadmap:
    carbonstack/roadmap/ROADMAP.md

    [DOC] top README:
    carbonstack/README.md

### Comms runtime surfaces

    [COMMS] CLI entry:
    carbonstack-comms/cmd/comms/main.go

    [COMMS] command implementation:
    carbonstack-comms/internal/app/commands.go

    [COMMS] openmls-send-dev tests:
    carbonstack-comms/internal/app/openmls_send_dev_test.go

    [COMMS] openmls-inbox-dev tests:
    carbonstack-comms/internal/app/openmls_inbox_dev_test.go

    [COMMS] dev runtime smoke proof:
    carbonstack-comms/scripts/dev-openmls-runtime-smoke.sh

    [COMMS] Cypher client:
    carbonstack-comms/internal/client/cypher.go

    [COMMS] relay helpers:
    carbonstack-comms/internal/relay/cypher_bridge.go
    carbonstack-comms/internal/relay/openmls_artifacts.go

    [COMMS] OpenMLS sidecar dir:
    carbonstack-comms/internal/protocol/mls/openmls-sidecar

    [COMMS] sidecar README:
    carbonstack-comms/internal/protocol/mls/openmls-sidecar/README.md

### Current stub-era Comms functions / behaviors

    send:
      trust policy evaluation
      mock/stub provider encryption
      Cypher envelope submit

    inbox:
      Cypher inbox retrieval
      mock/stub provider plaintext display as stub_plaintext

    ack:
      envelope ack by envelope ID and recipient device ID

### Implemented v0.4.x dev OpenMLS commands

    openmls-send-dev:
      state.RequireReadyDevice
      trust.EvaluateSend
      runOpenMLSMessageProtect
      sidecar message-protect
      internal/relay.SubmitOpenMLSArtifactEnvelope
      ArtifactKindApplicationMessage
      dev/pre-alpha warning

    openmls-inbox-dev:
      state.RequireReadyDevice
      Cypher inbox
      skip unsupported envelopes
      internal/relay.WriteOpenMLSArtifactFromEnvelope
      runOpenMLSMessageOpen
      sidecar message-open
      print plaintext after sidecar success
      optional ack only after sidecar success

### v0.4.5 runtime smoke proof function

    scripts/dev-openmls-runtime-smoke.sh:
      temp workspace
      temp Cypher binary
      temp Cypher DB
      dynamic loopback Cypher port
      temp Alice/Bob Comms state
      sidecar state reset
      sidecar identity/conversation bootstrap
      openmls-send-dev application-message send
      openmls-inbox-dev application-message open + ack
      inbox-empty-after-ack verification
      robust cleanup with INT -> TERM -> KILL escalation

### OpenMLS sidecar commands relevant to v0.4.x

    provider-info
    identity-create
    identity-status
    public-bundle-export
    conversation-create
    conversation-load-check
    conversation-add-member
    conversation-join
    message-protect
    message-open

### Relay/API functions relevant to runtime integration

    internal/client.CypherClient.SubmitEnvelope
    internal/client.CypherClient.Inbox
    internal/client.CypherClient.AckEnvelope
    internal/relay.SubmitOpenMLSArtifactEnvelope
    internal/relay.WriteOpenMLSArtifactFromEnvelope

### Existing v0.4.0 package validation

    cd <package-root>/carbonstack/tools/carbonstack-validate
    go run . --profile verify-checksums --root <package-root>
    go run . --profile full --root <package-root> --clean-generated

### Live umbrella validation

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go run . --profile local-cypher
    go run . --profile doctor
    go run . --profile core --clean-generated

### Envelope/protocol constants

Accepted opaque OpenMLS application-message pair:

    content_type=carbonstack.mls.application-message.v0
    protocol_version=carbonstack-openmls-sidecar-v0

Historical invalid pair validated by local-cypher:

    content_type=carbonstack.message.text.stub.v0
    protocol_version=carbonstack-openmls-sidecar-v0

Expected result:

    HTTP 400
    error code unsupported_protocol_version

---

## 10. v0.4.x Forward Plan

### v0.4.7 — pre-local-backbone assessment / validation-profile decision

Expected:

    assess whether smoke proof should stay a comms-local helper for now.
    assess whether to wrap it in carbonstack/tools/carbonstack-validate.
    assess whether local-backbone naming is still premature.
    decide what setup/reset semantics must exist before any runner profile.
    avoid overclaiming public release/package readiness.

### v0.4.8+ — possible runner-profile or ergonomics work

Possible paths:

    add a clearly named dev-runtime-openmls-smoke runner profile, not local-backbone.
    or defer runner integration and improve sidecar setup/reset ergonomics first.
    or document command migration/deprecation options for stub-era send/inbox.
    or split commands.go helpers if runtime command complexity keeps growing.

### v0.4.x later polish

Likely needs:

    exact sidecar setup runbook.
    clearer dev state reset semantics.
    better command output shape.
    possible helper extraction if commands.go becomes too large.
    decision on whether stub-era send/inbox are aliased, replaced, or deprecated.
    decision on whether local-backbone naming is still premature.

---

## 11. v0.5.x and Later Carry-Forward

### v0.5.x

Local storage, trust state, provider state, vault design, and PQ/hybrid ciphersuite migration readiness.

Do not start PQ/hybrid implementation before the runtime OpenMLS/Cypher path exists and is documented. PQ/hybrid migration touches provider state, ciphersuite metadata, trust display, artifact size, reinit/rekey/recovery policy, and claim boundaries.

### v0.6.x+

Hostile-server and abuse-resistance harnesses.

### v0.7.x+

Deployability and operations hardening.

### v0.8.x+

Documented self-pentest / adversarial validation.

### v0.9.x+

Claim-boundary review.

### v0.10.x+

Android backend/app work.

### v1.x.x

Public app/server major epoch and much later CarbonStackOS/appliance research, if justified.

---

## 12. Lean Breakpoint Summary

    CarbonStack v0.4.6 is the smoke-proof documentation / main-repo status alignment checkpoint. v0.4.0 remains the current public broad local deployability pre-release. carbonstack is now at 3cd77a8 docs: record dev OpenMLS runtime smoke proof. carbonstack-comms remains at 8e6e8b4 test: add dev OpenMLS runtime smoke proof. v0.4.6 adds carbonstack/docs/157-dev-runtime-openmls-smoke-proof-v0.md, updates docs/README.md, updates roadmap/ROADMAP.md, and lightly aligns README.md. The main repo now records that CarbonStackComms has a dev runtime OpenMLS CLI smoke proof for openmls-send-dev -> Cypher -> openmls-inbox-dev --ack. Validation passed for local-cypher, doctor, and core --clean-generated, with final clean repo heads across all four repos. Existing send/inbox remain stub-era. This is not local-backbone, not mature messaging UX, and not production security proof. Next safe rung: v0.4.7 pre-local-backbone assessment / validation-profile decision.

---

## 13. Preserved Immediate Previous Handoff: v0.4.5

The following is the previous v0.4.5 handoff. Where it conflicts with the v0.4.6 overlay above, v0.4.6 wins for current state.


# CarbonStack LogDoc v0.4.5

**Last updated:** 2026-06-04 local session  
**Owner/operator:** bitcrusher32  
**Assistant/model:** GPT-5.5 Thinking  
**Current phase:** Phase 2F / **v0.4.x dev runtime OpenMLS CLI smoke proof checkpoint complete**. The v0.4.0 public broad local deployability pre-release remains the current public release. v0.4.1 identified the runtime seam; v0.4.2 defined the runtime OpenMLS command contract; v0.4.3 implemented `openmls-send-dev`; v0.4.4 implemented `openmls-inbox-dev`; v0.4.5 now adds and validates `scripts/dev-openmls-runtime-smoke.sh`, proving the dev CLI runtime message path `openmls-send-dev -> Cypher -> openmls-inbox-dev --ack` against a temporary local Cypher server and dev-local OpenMLS sidecar setup. `carbonstack` remains at `a799764 (HEAD -> main, origin/main, origin/HEAD) docs: define runtime OpenMLS command contract`; `carbonstack-comms` is now at `8e6e8b4 (HEAD -> main, origin/main, origin/HEAD) test: add dev OpenMLS runtime smoke proof`; `carbonstack-cypher` remains at `9ab994c (HEAD -> main, origin/main, origin/HEAD) docs: point to local operator runbook`; `carbonstack-os` remains at `1bbbe52 (HEAD -> main, origin/main, origin/HEAD) docs: clarify CarbonStackOS target direction`.

**Version schema:** `v[scope].[timeline]`. This file is `v0.4.5`, the first dev runtime smoke proof checkpoint after the v0.4.3/v0.4.4 dev command implementation rungs. Per LogDoc V2 practice, this Markdown file preserves current operational continuity and carries the v0.4.x runtime-integration thread forward; the JSON breakpoint remains the lean machine-readable handoff state.

**Compression note:** The detailed v0.3.0 through v0.3.36PRIME process ledger remains archived in PRIME/sanitized PRIME. v0.4.0 compressed that pre-release sprint; v0.4.1 preserved the runtime recon; v0.4.2 preserved the command-contract decision; v0.4.3 preserved the first dev-only OpenMLS send-side implementation; v0.4.4 preserved the first dev-only OpenMLS receive/open/optional-ack implementation; v0.4.5 now preserves the first committed dev runtime smoke proof tying the two dev commands together. Future work should continue from this v0.4.x runtime thread.

---

## 1. Project Goal

**Active goal:** Continue moving CarbonStackComms from stub-era runtime `send`/`inbox` behavior toward explicit dev-only OpenMLS/Cypher-backed runtime commands and smoke validation. v0.4.5 completed the first committed dev runtime smoke proof: `openmls-send-dev -> Cypher -> openmls-inbox-dev --ack`.

CarbonStack’s broader project intent remains:

- **CarbonStack** is the secure-communications appliance-stack project and shared doctrine/release authority.
- **CarbonStackComms** is the text-first messaging client and current OpenMLS sidecar/relay integration surface.
- **CarbonStackCypher** is the self-hostable relay/storage stack.
- **CarbonStackOS** is future constrained Android-derived appliance OS work and remains out of current runnable validation packages.
- The `carbonstack` Gitea repo remains the source of truth for release state, release assets, docs, and public claims.
- GitHub mirrors exist as secondary push mirrors only; Gitea remains authoritative.

**North star:**

> Its goal is not to be a better smartphone. Its goal is to stop being a smartphone.

**Core doctrine:**

> Every feature is guilty until it proves it does not add unacceptable parser, network, sensor, identity, or filesystem authority.

**Current reality after v0.4.5:** `openmls-send-dev` and `openmls-inbox-dev` are implemented, tested, documented, and now exercised together by a committed dev runtime smoke script. The existing stub-era `send`, `inbox`, and `ack` commands remain in place and must not be described as OpenMLS-backed.

Correct interpretation:

    openmls-send-dev is a dev-only OpenMLS send-side command.
    openmls-inbox-dev is a dev-only OpenMLS receive/open/optional-ack command.
    scripts/dev-openmls-runtime-smoke.sh proves the dev CLI application-message path:
      openmls-send-dev -> Cypher -> openmls-inbox-dev --ack
    The smoke script directly bootstraps sidecar identities, KeyPackage/Welcome, and conversation setup for dev purposes.
    The application-message path is the runtime CLI proof target.
    Existing send/inbox remain stub-era.
    local-backbone remains reserved until whole-path validation is integrated/documented enough to deserve that name.

Hard nonclaims remain: v0.4.5 does **not** prove production readiness, production E2EE, hostile-server safety, metadata privacy, Android readiness, CarbonStackOS readiness, secure vault/storage, public ingress safety, real homelab validation, audit, certification, PQ/hybrid security, mature messenger UX, or local-backbone readiness.

---

## 2. Current Public Release State

Public release:

    https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.4.0

Release title:

    CarbonStack v0.4.0 Broad Local Deployability Pre-Release

Release tag:

    v0.4.0

Release commit:

    24ac9fc494 / 24ac9fc docs: record v0.4.0 package rehearsal

Gitea release status:

    Pre-Release

Release assets attached:

    carbonstack-v0.4.0-broad-local-deployability-pre-release.tgz
    carbonstack-v0.4.0-release-manifest.json
    carbonstack-v0.4.0-package-checksums.txt
    carbonstack-v0.4.0-asset-checksums.txt
    carbonstack-v0.4.0-validation-freeze.md
    v0.4.0-testing-runbook.md
    v0.4.0-release-notes.md
    LICENSE

Important release-page warning:

    The default Gitea Source Code ZIP/TAR.GZ downloads are auto-generated carbonstack-only archives.
    The intended runnable validation artifact is the attached multi-repo release package plus companion assets.

Post-release public surface:

    carbonstack README is evergreen.
    It points interested parties to the Releases page for runnable / known-good artifacts.
    It states Gitea is the source of truth.
    GitHub mirrors are secondary push mirrors, not release authority.

---

## 3. Current Repo Heads

    carbonstack        a799764 (HEAD -> main, origin/main, origin/HEAD) docs: define runtime OpenMLS command contract
    carbonstack-comms  8e6e8b4 (HEAD -> main, origin/main, origin/HEAD) test: add dev OpenMLS runtime smoke proof
    carbonstack-cypher 9ab994c (HEAD -> main, origin/main, origin/HEAD) docs: point to local operator runbook
    carbonstack-os     1bbbe52 (HEAD -> main, origin/main, origin/HEAD) docs: clarify CarbonStackOS target direction

Current public release tag:

    v0.4.0 -> 24ac9fc494 / 24ac9fc docs: record v0.4.0 package rehearsal

Post-release public-surface cleanup:

    6ba7aab docs: make front readme evergreen

v0.4.1 runtime recon commit:

    83d2f0f docs: record runtime OpenMLS Comms recon

v0.4.2 command contract commit:

    a799764 docs: define runtime OpenMLS command contract

v0.4.3 send implementation commit:

    39fb287 feat: add dev OpenMLS send command

v0.4.4 inbox implementation commit:

    65bc707 feat: add dev OpenMLS inbox command

Current v0.4.5 smoke proof commit:

    8e6e8b4 test: add dev OpenMLS runtime smoke proof

Continuity note:

    `8e6e8b4` is a `carbonstack-comms` test/script/docs commit after the `v0.4.0` public release tag.
    It does not retag v0.4.0.
    The release remains attached to `24ac9fc494`.

---

## 4. Current Validated / Known State

Validated / known at the v0.4.5 checkpoint:

    [RELEASE] v0.4.0 Gitea release exists, is marked Pre-Release, and is attached to 24ac9fc494.
    [PACKAGE] v0.4.0 package includes carbonstack, carbonstack-comms, carbonstack-cypher, and release metadata.
    [PACKAGE] carbonstack-os is deliberately excluded from the runnable package.
    [RUNNER] full means release-snapshot followed by local-cypher.
    [RUNNER] release-snapshot validates package layout, metadata, strict pre-test artifact scan, package checksums, and core.
    [RUNNER] core validates doctor, OpenMLS real-Cypher lifecycle, Comms package tests, Cypher package tests, and artifact scans.
    [RUNNER] local-cypher validates Cypher-only lifecycle, restart/persistence against the same temporary DB, positive opaque OpenMLS envelope/ack flow, and historical negative stub-text/OpenMLS protocol pairing.
    [DOCS] main carbonstack README is evergreen.
    [DOCS] docs/155-runtime-comms-openmls-cypher-recon-v0.md records the v0.4.1 runtime recon.
    [DOCS] docs/156-runtime-openmls-command-contract-v0.md records the v0.4.2 command contract.
    [COMMS] openmls-send-dev exists in carbonstack-comms.
    [COMMS] openmls-inbox-dev exists in carbonstack-comms.
    [COMMS] scripts/dev-openmls-runtime-smoke.sh exists in carbonstack-comms.
    [COMMS] README documents the dev runtime OpenMLS smoke proof.
    [COMMS] existing send/inbox/ack remain in place and stub-era.
    [MIRRORS] all four repos have GitHub push-mirror functionality, with Gitea remaining source of truth.

v0.4.5 smoke proof result from the final log:

    scripts/dev-openmls-runtime-smoke.sh passed and self-exited.
    The timed run completed in about 5.452 seconds after warm builds/tooling.
    The script created a temporary local Cypher server and temporary Comms state files.
    The script created dev-local OpenMLS sidecar identities and a dev-local conversation.
    The script used openmls-send-dev to submit an application-message envelope through Cypher.
    The script used openmls-inbox-dev --ack to retrieve, open, print, and ack the message.
    The opened plaintext matched:
      hello bob through openmls runtime smoke
    The message was acked after sidecar success.
    Bob's inbox was empty after ack.
    The script printed:
      PASS: dev runtime OpenMLS CLI smoke proof
      proof: openmls-send-dev -> Cypher -> openmls-inbox-dev --ack
      boundary: dev/pre-alpha smoke proof; not local-backbone; not production messaging UX

v0.4.5 validation result from the final log:

    scripts/dev-openmls-runtime-smoke.sh passed more than once after cleanup fix.
    go test ./internal/app -count=1 passed.
    go test ./... -count=1 -timeout 600s passed in carbonstack-comms.
    local-cypher passed.
    doctor passed.
    core --clean-generated passed.
    OpenMLS real-Cypher lifecycle passed during core.
    carbonstack-comms package tests passed during core.
    carbonstack-cypher package tests passed during core.
    --clean-generated removed known OpenMLS sidecar generated roots.
    final repo status showed carbonstack-comms at 8e6e8b4 and all four repos clean.

Observed validation nuance:

    local-cypher and core saw known OpenMLS sidecar generated roots in pre/post artifact scans because OpenMLS sidecar tests and the new runtime smoke script had just been run.
    The runner classified them as known OpenMLS generated roots.
    core --clean-generated removed .carbonstack-openmls-sidecar-state and target.
    This is not a release failure or source hygiene failure as long as hits remain limited to known generated roots and cleanup is explicit.

Expected v0.4.0 public package validation commands remain:

    cd <package-root>/carbonstack/tools/carbonstack-validate
    go run . --profile verify-checksums --root <package-root>
    go run . --profile full --root <package-root> --clean-generated

Live umbrella validation commands remain:

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go run . --profile local-cypher
    go run . --profile doctor
    go run . --profile core --clean-generated

New v0.4.5 smoke command:

    cd ~/repos/carbonstack_umbrella/carbonstack-comms
    scripts/dev-openmls-runtime-smoke.sh

---

## 5. v0.4.5 Work Completed

### 5.1 Runtime smoke proof recon was performed

The v0.4.5 recon inspected:

    carbonstack-comms dev command surfaces
    openmls-send-dev implementation
    openmls-inbox-dev implementation
    carbonstack-comms README dev command note
    OpenMLS sidecar command/flag surface
    sidecar state path helpers
    existing full lifecycle real-Cypher test
    Cypher command/config surface
    current validation baseline

Useful findings:

    openmls-send-dev and openmls-inbox-dev are both registered in internal/app/commands.go.
    The sidecar supports identity-create, public-bundle-export, conversation-create, conversation-add-member, conversation-join, message-protect, and message-open.
    The real-Cypher lifecycle test already provides a deterministic sidecar bootstrap sequence.
    Cypher can be built and run temporarily from the umbrella sibling repo.
    A script can bootstrap sidecar identities/conversation directly, then use Comms CLI for the application-message runtime proof.
    This is enough for Option B: a scripted smoke proof rather than docs-only planning.

### 5.2 dev-openmls-runtime-smoke.sh was added

New script:

    carbonstack-comms/scripts/dev-openmls-runtime-smoke.sh

Script behavior:

    resolves carbonstack-comms repo root
    resolves umbrella root
    resolves carbonstack-cypher sibling path
    resolves OpenMLS sidecar path
    requires go, cargo, curl, and python3
    creates a temporary smoke workspace
    builds a temporary Cypher binary
    starts a temporary Cypher server on a dynamic loopback port
    creates temporary Alice and Bob Comms state files
    claims/registers Alice and Bob through Comms CLI
    creates dev-local OpenMLS sidecar identities
    bootstraps sidecar conversation state directly:
      Bob public-bundle-export --write-artifact
      Alice conversation-create
      Alice conversation-add-member with Bob KeyPackage
      Bob conversation-join with Welcome
    sends application message through:
      go run ./cmd/comms openmls-send-dev
    opens and acks application message through:
      go run ./cmd/comms openmls-inbox-dev --ack
    checks send output for status/content_type/protocol_version
    checks inbox output for open success, plaintext match, and ack success
    verifies Bob inbox is empty after ack
    prints explicit pass/proof/boundary lines

Critical boundary:

    sidecar KeyPackage/Welcome/bootstrap setup is still direct dev setup.
    the application-message path is the runtime CLI proof target.
    this is not local-backbone.
    this is not mature messaging UX.
    this is not production security proof.

### 5.3 Cleanup/self-termination was fixed

Initial manual run proved the runtime path, but the user had to manually interrupt after the PASS block.

The likely issue was cleanup waiting on the temporary Cypher process after sending INT.

Fix applied:

    cleanup trap now handles EXIT, INT, and TERM.
    cleanup sends INT first.
    cleanup waits briefly.
    cleanup escalates to TERM.
    cleanup waits briefly.
    cleanup escalates to KILL if necessary.
    cleanup waits and suppresses expected process-exit noise.

Result:

    timed rerun printed PASS and returned to the shell prompt by itself.
    observed time:
      real 0m5.452s
      user 0m1.676s
      sys 0m0.957s

Lesson:

    runtime smoke scripts should be self-terminating and automation/logging friendly.
    a proof that only finishes after Ctrl+C is not acceptable as a committed smoke harness.

### 5.4 README was updated

`carbonstack-comms/README.md` now includes:

    ## Dev runtime OpenMLS smoke proof

The README explains:

    scripts/dev-openmls-runtime-smoke.sh is the smoke command.
    It creates temporary local Cypher server/state.
    It creates dev-local sidecar identities/conversation.
    It proves:
      openmls-send-dev -> Cypher -> openmls-inbox-dev --ack
    Boundary:
      dev/pre-alpha only
      not local-backbone
      not production messaging UX
      does not replace old stub-era send/inbox
      sidecar KeyPackage/Welcome/bootstrap setup is still direct dev setup
      application-message path is runtime CLI proof target

### 5.5 Validation passed

Commands observed passing:

    cd ~/repos/carbonstack_umbrella/carbonstack-comms
    time scripts/dev-openmls-runtime-smoke.sh
    scripts/dev-openmls-runtime-smoke.sh
    go test ./internal/app -count=1
    go test ./... -count=1 -timeout 600s

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go run . --profile local-cypher
    go run . --profile doctor
    go run . --profile core --clean-generated

Final commit/push:

    carbonstack-comms:
    8e6e8b4 test: add dev OpenMLS runtime smoke proof

Final repo snapshot:

    carbonstack        a799764 docs: define runtime OpenMLS command contract
    carbonstack-comms  8e6e8b4 test: add dev OpenMLS runtime smoke proof
    carbonstack-cypher 9ab994c docs: point to local operator runbook
    carbonstack-os     1bbbe52 docs: clarify CarbonStackOS target direction

---

## 6. Critical Blunders / Lessons to Carry Forward

### 6.1 A passing smoke proof still needs clean exit behavior

The first smoke run printed the PASS block and proved the message path, but the user had to press Ctrl+C.

This was not a cryptographic/protocol failure, but it was a harness quality issue.

Correct fix:

    make cleanup robust.
    handle EXIT/INT/TERM.
    attempt graceful interrupt first.
    escalate to TERM/KILL if needed.
    verify timed run returns to shell by itself.

### 6.2 v0.4.5 proves the application-message CLI path, not onboarding

The script uses direct sidecar setup for:

    identity-create
    public-bundle-export
    conversation-create
    conversation-add-member
    conversation-join

The actual runtime CLI proof is:

    openmls-send-dev -> Cypher -> openmls-inbox-dev --ack

Do not overclaim:

    full onboarding is not runtime UX yet.
    Relay Space join model is not implemented.
    old send/inbox are not OpenMLS-backed.
    local-backbone is not yet claimed.

### 6.3 Dev trust warning is expected

The smoke output includes:

    WARNING: recipient device is unknown; dev mode allows sending but mature mode should block until verification
    WARNING: openmls-send-dev is running in dev trust mode; use --strict to block unknown or unverified recipients

This is expected at this stage.

Do not hide it; it is useful evidence that trust-state maturity remains future work.

### 6.4 Scripted smoke proof is useful but not yet a runner profile

The smoke script is now repeatable, but it is still a carbonstack-comms dev script.

Do not immediately move it into `carbonstack/tools/carbonstack-validate` or call it `local-backbone` until:

    boundaries are documented in main carbonstack docs.
    generated state cleanup semantics are well understood.
    the command is stable enough for public release package validation.
    the name does not overclaim deployment maturity.

### 6.5 Known OpenMLS generated roots remain expected around sidecar tests/smoke

OpenMLS tests and the smoke script may produce:

    .carbonstack-openmls-sidecar-state
    target

The runner may see these in pre/post scans if tests/smoke ran immediately beforehand. This remains acceptable only because:

    they are classified as known OpenMLS generated roots.
    artifact scan is non-destructive.
    --clean-generated explicitly removes them.

Do not generalize cleanup beyond known generated roots.

### 6.6 Commit metadata warning remains a polish issue

The commit log again showed auto-generated local commit metadata:

    Committer: bitcrusher32 <▓▓>

This is not a functional validation failure, but it remains a professionalism/publishing hygiene issue to address eventually with `git config --global user.name` and `git config --global user.email` if desired.

---

## 7. Current Blockers / Not Validated

Current blockers and nonvalidated areas after v0.4.5:

    Existing send/inbox remain stub-era.
    Runtime Comms has dev-only OpenMLS commands, but no mature user-facing UX.
    Sidecar KeyPackage/Welcome/bootstrap setup is still direct dev setup, not Comms runtime UX.
    No Relay Space join model exists yet.
    local-backbone runner profile does not exist and remains reserved.
    The dev runtime smoke script is not yet integrated into the main carbonstack validation runner.
    Post-release public download verification is still useful if not already performed.
    Sanitized PRIME placement in carbonstack/logdoc-loop-system still needs a policy decision.
    Production-safe provider storage and secure vault design remain v0.5.x future work.
    PQ/hybrid ciphersuite migration readiness remains v0.5.x future work.
    Hostile-server rollback/replay/metadata-abuse validation remains future work.
    Android/CarbonStackOS implementation remains later work.
    CarbonStack Relay Space mechanics remain a future concept, not an implementation.
    Public ingress, systemd, cloudflared, real homelab deployment, and production service paths remain out of scope.

Hard nonclaims remain:

    Do not claim production readiness.
    Do not claim production E2EE.
    Do not claim hostile-server safety.
    Do not claim metadata privacy.
    Do not claim Debian production deployability.
    Do not claim systemd readiness.
    Do not claim cloudflared readiness.
    Do not claim Android readiness.
    Do not claim CarbonStackOS readiness.
    Do not claim secure vault/storage.
    Do not claim real homelab validation.
    Do not claim external audit or certification.
    Do not claim CarbonStack Relay Space mechanics are implemented.
    Do not claim local-backbone exists or is justified yet.
    Do not claim local-cypher validates runtime Comms UX.
    Do not claim full validates local-backbone.
    Do not claim full is a deployment command.
    Do not claim send/inbox are OpenMLS-backed yet.
    Do not claim PQ/hybrid security or quantum-safe messaging.
    Do not let GitHub mirrors supersede Gitea source-of-truth status.

Allowed claim:

    CarbonStackComms now has dev-only openmls-send-dev and openmls-inbox-dev commands.
    CarbonStackComms now has a dev runtime smoke script proving openmls-send-dev -> Cypher -> openmls-inbox-dev --ack.
    The smoke script starts a temporary local Cypher server, uses temporary Comms state files, creates dev-local OpenMLS sidecar identities/conversation, sends an OpenMLS application message through the Comms CLI, opens it through the Comms CLI, acks after sidecar success, and verifies Bob's inbox is empty after ack.
    Existing send/inbox remain stub-era.
    v0.4.0 remains the current public broad local deployability pre-release.
    Gitea remains source of truth; GitHub mirrors are secondary push mirrors.

---

## 8. Next Safest Actions

Recommended next rung:

    v0.4.6 smoke proof documentation / main-repo status alignment

Focus:

    record the v0.4.5 runtime smoke proof in main carbonstack docs/roadmap.
    decide whether the smoke script remains a comms-local dev helper or should later become a runner profile.
    keep local-backbone reserved unless the project explicitly decides the smoke proof is stable enough and non-overclaiming.
    document the boundary:
      direct sidecar bootstrap still required.
      application-message path is runtime CLI-proven.
      existing send/inbox are still stub-era.
      not production UX/security.

Possible follow-up after docs alignment:

    v0.4.7 runner-profile recon or scripted validation planning.
    v0.4.8 command ergonomics / setup simplification.
    v0.4.9 decision on local-backbone naming or further deferral.
    v0.5.x state/trust/vault/PQ planning once v0.4.x runtime proof is sufficiently documented and stable.

Expected implementation repo next:

    carbonstack for docs/roadmap alignment.
    carbonstack-comms only if small smoke-script refinements are needed.

Avoid next:

    replacing send/inbox immediately.
    adding local-backbone prematurely.
    moving smoke proof into public release validation without a boundary doc.
    public ingress.
    cloudflared.
    systemd.
    real homelab deployment.
    Android app.
    CarbonStackOS implementation.
    CarbonStack Relay Space implementation.
    PQ/hybrid ciphersuite implementation.
    production/security claims.
    broad negative-path suite unless the smoke proof exposes a direct blocker.

---

## 9. Critical Paths / Functions

### Source-of-truth public surfaces

    [GITEA SOURCE OF TRUTH] https://git.bitcrusher32.win/bitcrusher32/carbonstack
    [RELEASES] https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases
    [CURRENT PUBLIC RELEASE] https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.4.0

GitHub mirrors:

    secondary push mirrors for discoverability/redundancy only.
    Do not treat mirrors as release authority unless project policy changes.

### WSL Debian active working umbrella

    [WSL-DEBIAN] umbrella: ~/repos/carbonstack_umbrella
    [WSL-DEBIAN] carbonstack: ~/repos/carbonstack_umbrella/carbonstack
    [WSL-DEBIAN] carbonstack-comms: ~/repos/carbonstack_umbrella/carbonstack-comms
    [WSL-DEBIAN] carbonstack-cypher: ~/repos/carbonstack_umbrella/carbonstack-cypher
    [WSL-DEBIAN] carbonstack-os: ~/repos/carbonstack_umbrella/carbonstack-os
    [GO-RUNNER] runner dir: ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate

### Runtime integration doc surfaces

    [DOC] v0.4.1 runtime recon:
    carbonstack/docs/155-runtime-comms-openmls-cypher-recon-v0.md

    [DOC] v0.4.2 command contract:
    carbonstack/docs/156-runtime-openmls-command-contract-v0.md

    [DOC] docs index:
    carbonstack/docs/README.md

    [DOC] roadmap:
    carbonstack/roadmap/ROADMAP.md

### Comms runtime surfaces

    [COMMS] CLI entry:
    carbonstack-comms/cmd/comms/main.go

    [COMMS] command implementation:
    carbonstack-comms/internal/app/commands.go

    [COMMS] openmls-send-dev tests:
    carbonstack-comms/internal/app/openmls_send_dev_test.go

    [COMMS] openmls-inbox-dev tests:
    carbonstack-comms/internal/app/openmls_inbox_dev_test.go

    [COMMS] dev runtime smoke proof:
    carbonstack-comms/scripts/dev-openmls-runtime-smoke.sh

    [COMMS] Cypher client:
    carbonstack-comms/internal/client/cypher.go

    [COMMS] relay helpers:
    carbonstack-comms/internal/relay/cypher_bridge.go
    carbonstack-comms/internal/relay/openmls_artifacts.go

    [COMMS] OpenMLS sidecar dir:
    carbonstack-comms/internal/protocol/mls/openmls-sidecar

    [COMMS] sidecar README:
    carbonstack-comms/internal/protocol/mls/openmls-sidecar/README.md

### Current stub-era Comms functions / behaviors

    send:
      trust policy evaluation
      mock/stub provider encryption
      Cypher envelope submit

    inbox:
      Cypher inbox retrieval
      mock/stub provider plaintext display as stub_plaintext

    ack:
      envelope ack by envelope ID and recipient device ID

### Implemented v0.4.x dev OpenMLS commands

    openmls-send-dev:
      state.RequireReadyDevice
      trust.EvaluateSend
      runOpenMLSMessageProtect
      sidecar message-protect
      internal/relay.SubmitOpenMLSArtifactEnvelope
      ArtifactKindApplicationMessage
      dev/pre-alpha warning

    openmls-inbox-dev:
      state.RequireReadyDevice
      Cypher inbox
      skip unsupported envelopes
      internal/relay.WriteOpenMLSArtifactFromEnvelope
      runOpenMLSMessageOpen
      sidecar message-open
      print plaintext after sidecar success
      optional ack only after sidecar success

### v0.4.5 runtime smoke proof function

    scripts/dev-openmls-runtime-smoke.sh:
      temp workspace
      temp Cypher binary
      temp Cypher DB
      dynamic loopback Cypher port
      temp Alice/Bob Comms state
      sidecar state reset
      sidecar identity/conversation bootstrap
      openmls-send-dev application-message send
      openmls-inbox-dev application-message open + ack
      inbox-empty-after-ack verification
      robust cleanup with INT -> TERM -> KILL escalation

### OpenMLS sidecar commands relevant to v0.4.x

    provider-info
    identity-create
    identity-status
    public-bundle-export
    conversation-create
    conversation-load-check
    conversation-add-member
    conversation-join
    message-protect
    message-open

### Relay/API functions relevant to runtime integration

    internal/client.CypherClient.SubmitEnvelope
    internal/client.CypherClient.Inbox
    internal/client.CypherClient.AckEnvelope
    internal/relay.SubmitOpenMLSArtifactEnvelope
    internal/relay.WriteOpenMLSArtifactFromEnvelope

### Existing v0.4.0 package validation

    cd <package-root>/carbonstack/tools/carbonstack-validate
    go run . --profile verify-checksums --root <package-root>
    go run . --profile full --root <package-root> --clean-generated

### Live umbrella validation

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go run . --profile local-cypher
    go run . --profile doctor
    go run . --profile core --clean-generated

### Envelope/protocol constants

Accepted opaque OpenMLS application-message pair:

    content_type=carbonstack.mls.application-message.v0
    protocol_version=carbonstack-openmls-sidecar-v0

Historical invalid pair validated by local-cypher:

    content_type=carbonstack.message.text.stub.v0
    protocol_version=carbonstack-openmls-sidecar-v0

Expected result:

    HTTP 400
    error code unsupported_protocol_version

---

## 10. v0.4.x Forward Plan

### v0.4.6 — smoke proof documentation / main-repo status alignment

Expected:

    main carbonstack doc records the v0.4.5 dev runtime smoke proof.
    roadmap v0.4.x status updates from "attempt smoke proof" to "smoke proof exists in carbonstack-comms".
    docs clarify not local-backbone yet.
    docs clarify direct sidecar bootstrap remains.

### v0.4.7+ — runner-profile recon or deferral

Only consider runner/profile movement when:

    the smoke script is stable enough.
    generated state cleanup is predictable.
    boundary docs are clear.
    public release/package validation will not overclaim the command.

### v0.4.x later polish

Likely needs:

    exact sidecar setup runbook.
    clearer dev state reset semantics.
    better command output shape.
    possible helper extraction if commands.go becomes too large.
    decision on whether stub-era send/inbox are aliased, replaced, or deprecated.
    decision on whether local-backbone naming is still premature.

---

## 11. v0.5.x and Later Carry-Forward

### v0.5.x

Local storage, trust state, provider state, vault design, and PQ/hybrid ciphersuite migration readiness.

Do not start PQ/hybrid implementation before the runtime OpenMLS/Cypher path exists and is documented. PQ/hybrid migration touches provider state, ciphersuite metadata, trust display, artifact size, reinit/rekey/recovery policy, and claim boundaries.

### v0.6.x+

Hostile-server and abuse-resistance harnesses.

### v0.7.x+

Deployability and operations hardening.

### v0.8.x+

Documented self-pentest / adversarial validation.

### v0.9.x+

Claim-boundary review.

### v0.10.x+

Android backend/app work.

### v1.x.x

Public app/server major epoch and much later CarbonStackOS/appliance research, if justified.

---

## 12. Lean Breakpoint Summary

    CarbonStack v0.4.5 is the dev runtime OpenMLS CLI smoke proof checkpoint. v0.4.0 remains the current public broad local deployability pre-release. carbonstack remains at a799764 docs: define runtime OpenMLS command contract. carbonstack-comms is now at 8e6e8b4 test: add dev OpenMLS runtime smoke proof. v0.4.5 adds scripts/dev-openmls-runtime-smoke.sh and README documentation in carbonstack-comms. The script starts a temporary local Cypher server, creates temporary Comms state, bootstraps dev-local OpenMLS sidecar identity/conversation setup directly, then proves openmls-send-dev -> Cypher -> openmls-inbox-dev --ack for an application-message. The smoke output proves plaintext match, ack after sidecar success, and Bob inbox empty after ack. An initial manual run required Ctrl+C after PASS, so cleanup was hardened with INT -> TERM -> KILL fallback; a timed rerun self-exited in about 5.452s. Validation passed for the smoke script, app tests, comms package tests, local-cypher, doctor, and core --clean-generated. Existing send/inbox remain stub-era. This is not local-backbone, not mature messaging UX, and not production security proof. Next safe rung: v0.4.6 document the smoke proof in main carbonstack docs/roadmap and decide whether runner-profile movement is justified or premature.


---

## 13. Preserved Immediate Previous Handoff: v0.4.4

The following is the previous v0.4.4 handoff. Where it conflicts with the v0.4.5 overlay above, v0.4.5 wins for current state.


# CarbonStack LogDoc v0.4.4

**Last updated:** 2026-06-03 local session  
**Owner/operator:** bitcrusher32  
**Assistant/model:** GPT-5.5 Thinking  
**Current phase:** Phase 2F / **v0.4.x first dev-only OpenMLS inbox/open/ack command checkpoint complete**. The v0.4.0 public broad local deployability pre-release remains the current public release. v0.4.1 identified the runtime seam; v0.4.2 defined the runtime OpenMLS command contract; v0.4.3 implemented `openmls-send-dev`; v0.4.4 now implements the receive/open/optional-ack half through `openmls-inbox-dev` in `carbonstack-comms`. `carbonstack` remains at `a799764 (HEAD -> main, origin/main, origin/HEAD) docs: define runtime OpenMLS command contract`; `carbonstack-comms` is now at `65bc707 (HEAD -> main, origin/main, origin/HEAD) feat: add dev OpenMLS inbox command`; `carbonstack-cypher` remains at `9ab994c (HEAD -> main, origin/main, origin/HEAD) docs: point to local operator runbook`; `carbonstack-os` remains at `1bbbe52 (HEAD -> main, origin/main, origin/HEAD) docs: clarify CarbonStackOS target direction`.

**Version schema:** `v[scope].[timeline]`. This file is `v0.4.4`, the second implementation checkpoint after the v0.4.2 runtime OpenMLS command contract. Per LogDoc V2 practice, this Markdown file preserves current operational continuity and carries the v0.4.x runtime-integration thread forward; the JSON breakpoint remains the lean machine-readable handoff state.

**Compression note:** The detailed v0.3.0 through v0.3.36PRIME process ledger remains archived in PRIME/sanitized PRIME. v0.4.0 compressed that pre-release sprint; v0.4.1 preserved the runtime recon; v0.4.2 preserved the command-contract decision; v0.4.3 preserved the first dev-only OpenMLS send-side implementation; v0.4.4 now preserves the first dev-only OpenMLS receive/open/optional-ack implementation. Future work should continue from this v0.4.x runtime thread.

---

## 1. Project Goal

**Active goal:** Move CarbonStackComms from stub-era runtime `send`/`inbox` behavior toward explicit dev-only OpenMLS/Cypher-backed runtime commands. v0.4.4 completed the second implementation rung: `openmls-inbox-dev` now exists in `carbonstack-comms`.

CarbonStack’s broader project intent remains:

- **CarbonStack** is the secure-communications appliance-stack project and shared doctrine/release authority.
- **CarbonStackComms** is the text-first messaging client and current OpenMLS sidecar/relay integration surface.
- **CarbonStackCypher** is the self-hostable relay/storage stack.
- **CarbonStackOS** is future constrained Android-derived appliance OS work and remains out of current runnable validation packages.
- The `carbonstack` Gitea repo remains the source of truth for release state, release assets, docs, and public claims.
- GitHub mirrors exist as secondary push mirrors only; Gitea remains authoritative.

**North star:**

> Its goal is not to be a better smartphone. Its goal is to stop being a smartphone.

**Core doctrine:**

> Every feature is guilty until it proves it does not add unacceptable parser, network, sensor, identity, or filesystem authority.

**Current reality after v0.4.4:** `openmls-send-dev` and `openmls-inbox-dev` are implemented and tested as explicit dev-only OpenMLS runtime commands. The existing stub-era `send`, `inbox`, and `ack` commands remain in place and must not be described as OpenMLS-backed.

Correct interpretation:

    openmls-send-dev is a dev-only OpenMLS send-side command.
    openmls-inbox-dev is a dev-only OpenMLS receive/open/optional-ack command.
    openmls-inbox-dev fetches the current device inbox through Cypher, skips unsupported envelopes, writes OpenMLS application-message artifacts through the relay helper, calls sidecar message-open, prints plaintext only after sidecar success, and acks only after sidecar success when --ack is explicitly set.
    Existing send/inbox remain stub-era.
    There is now enough CLI surface to attempt a dev runtime smoke proof, but that proof has not been cut yet as a committed runner/documented milestone.
    local-backbone remains reserved until a repeatable whole-path runtime validation exists.

Hard nonclaims remain: v0.4.4 does **not** prove production readiness, production E2EE, hostile-server safety, metadata privacy, Android readiness, CarbonStackOS readiness, secure vault/storage, public ingress safety, real homelab validation, audit, certification, PQ/hybrid security, mature messenger UX, or local-backbone readiness.

---

## 2. Current Public Release State

Public release:

    https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.4.0

Release title:

    CarbonStack v0.4.0 Broad Local Deployability Pre-Release

Release tag:

    v0.4.0

Release commit:

    24ac9fc494 / 24ac9fc docs: record v0.4.0 package rehearsal

Gitea release status:

    Pre-Release

Release assets attached:

    carbonstack-v0.4.0-broad-local-deployability-pre-release.tgz
    carbonstack-v0.4.0-release-manifest.json
    carbonstack-v0.4.0-package-checksums.txt
    carbonstack-v0.4.0-asset-checksums.txt
    carbonstack-v0.4.0-validation-freeze.md
    v0.4.0-testing-runbook.md
    v0.4.0-release-notes.md
    LICENSE

Important release-page warning:

    The default Gitea Source Code ZIP/TAR.GZ downloads are auto-generated carbonstack-only archives.
    The intended runnable validation artifact is the attached multi-repo release package plus companion assets.

Post-release public surface:

    carbonstack README is evergreen.
    It points interested parties to the Releases page for runnable / known-good artifacts.
    It states Gitea is the source of truth.
    GitHub mirrors are secondary push mirrors, not release authority.

---

## 3. Current Repo Heads

    carbonstack        a799764 (HEAD -> main, origin/main, origin/HEAD) docs: define runtime OpenMLS command contract
    carbonstack-comms  65bc707 (HEAD -> main, origin/main, origin/HEAD) feat: add dev OpenMLS inbox command
    carbonstack-cypher 9ab994c (HEAD -> main, origin/main, origin/HEAD) docs: point to local operator runbook
    carbonstack-os     1bbbe52 (HEAD -> main, origin/main, origin/HEAD) docs: clarify CarbonStackOS target direction

Current public release tag:

    v0.4.0 -> 24ac9fc494 / 24ac9fc docs: record v0.4.0 package rehearsal

Post-release public-surface cleanup:

    6ba7aab docs: make front readme evergreen

v0.4.1 runtime recon commit:

    83d2f0f docs: record runtime OpenMLS Comms recon

v0.4.2 command contract commit:

    a799764 docs: define runtime OpenMLS command contract

v0.4.3 send implementation commit:

    39fb287 feat: add dev OpenMLS send command

Current v0.4.4 inbox implementation commit:

    65bc707 feat: add dev OpenMLS inbox command

Continuity note:

    `65bc707` is a `carbonstack-comms` implementation commit after the `v0.4.0` public release tag. It does not retag v0.4.0. The release remains attached to `24ac9fc494`.

---

## 4. Current Validated / Known State

Validated / known at the v0.4.4 checkpoint:

    [RELEASE] v0.4.0 Gitea release exists, is marked Pre-Release, and is attached to 24ac9fc494.
    [PACKAGE] v0.4.0 package includes carbonstack, carbonstack-comms, carbonstack-cypher, and release metadata.
    [PACKAGE] carbonstack-os is deliberately excluded from the runnable package.
    [RUNNER] full means release-snapshot followed by local-cypher.
    [RUNNER] release-snapshot validates package layout, metadata, strict pre-test artifact scan, package checksums, and core.
    [RUNNER] core validates doctor, OpenMLS real-Cypher lifecycle, Comms package tests, Cypher package tests, and artifact scans.
    [RUNNER] local-cypher validates Cypher-only lifecycle, restart/persistence against the same temporary DB, positive opaque OpenMLS envelope/ack flow, and historical negative stub-text/OpenMLS protocol pairing.
    [DOCS] main carbonstack README is evergreen.
    [DOCS] docs/155-runtime-comms-openmls-cypher-recon-v0.md records the v0.4.1 runtime recon.
    [DOCS] docs/156-runtime-openmls-command-contract-v0.md records the v0.4.2 command contract.
    [COMMS] openmls-send-dev exists in carbonstack-comms.
    [COMMS] openmls-inbox-dev exists in carbonstack-comms.
    [COMMS] existing send/inbox/ack remain in place and stub-era.
    [MIRRORS] all four repos have GitHub push-mirror functionality, with Gitea remaining source of truth.

v0.4.4 validation result from the final log:

    go test ./internal/app -count=1 passed.
    go test ./... -count=1 -timeout 600s passed in carbonstack-comms.
    local-cypher passed.
    doctor passed.
    core --clean-generated passed.
    OpenMLS real-Cypher lifecycle passed during core.
    carbonstack-comms package tests passed during core.
    carbonstack-cypher package tests passed during core.
    --clean-generated removed known OpenMLS sidecar generated roots.
    final repo status showed carbonstack-comms at 65bc707 and all four repos clean.

Observed validation nuance:

    local-cypher and core saw known OpenMLS sidecar generated roots in pre/post artifact scans because OpenMLS sidecar tests had just been run.
    The runner classified them as known OpenMLS generated roots.
    core --clean-generated removed .carbonstack-openmls-sidecar-state and target.
    This is not a release failure or source hygiene failure as long as hits remain limited to known generated roots and cleanup is explicit.

Expected v0.4.0 public package validation commands remain:

    cd <package-root>/carbonstack/tools/carbonstack-validate
    go run . --profile verify-checksums --root <package-root>
    go run . --profile full --root <package-root> --clean-generated

Live umbrella validation commands remain:

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go run . --profile local-cypher
    go run . --profile doctor
    go run . --profile core --clean-generated

---

## 5. v0.4.4 Work Completed

### 5.1 Focused implementation recon was performed

The v0.4.4 recon inspected `carbonstack-comms` command, client, relay, and OpenMLS sidecar message-open surfaces before implementation.

Useful findings:

    openmls-send-dev was already registered in internal/app/commands.go.
    client.CypherClient already exposed Inbox and AckEnvelope.
    client.EnvelopeRecord already carried content_type, protocol_version, ciphertext_b64, payload metadata, server timestamp, and delivery state.
    internal/relay.WriteOpenMLSArtifactFromEnvelope already validated content_type/protocol_version, decoded ciphertext_b64, checked payload_size_bytes/payload_sha256 when present, created output directories, and wrote artifact bytes.
    OpenMLS sidecar message-open used --device-label, --conversation-label, optional --message-label, and --message <path>.
    OpenMLS sidecar message-open success JSON included plaintext_utf8, plaintext_len, message_opened, message_artifact_path_hint, and message_open_summary_path_hint.
    Existing lower-level tests already proved sidecar message-open behavior through the real-Cypher lifecycle.

### 5.2 openmls-inbox-dev was implemented

`carbonstack-comms/internal/app/commands.go` now includes:

    openmls-inbox-dev in the app.Run switch.
    openmls-inbox-dev in usage output.
    cmdOpenMLSInboxDev.
    writeOpenMLSInboxDevArtifact.
    runOpenMLSMessageOpen.
    parseOpenMLSOpenEnvelope.
    package-level function variables for inbox, ack, artifact write, and sidecar open injection in tests.

Command shape:

    go run ./cmd/comms openmls-inbox-dev \
      --sidecar-device-label <recipient-sidecar-device-label> \
      --conversation <sidecar-conversation-label> \
      [--message-label <label>] \
      [--limit 1] \
      [--ack]

Implemented behavior:

    requires --sidecar-device-label and --conversation.
    defaults --limit to 1 and rejects --limit < 1.
    defaults --ack to false.
    loads ready Comms state and current device ID.
    fetches the current device inbox through Cypher.
    skips unsupported non-OpenMLS application-message envelopes with a clear skipped_unsupported_envelope output block.
    writes OpenMLS application-message artifacts through internal/relay.WriteOpenMLSArtifactFromEnvelope.
    calls sidecar message-open through cargo.
    prints plaintext only after sidecar message-open succeeds.
    does not ack unless --ack is set.
    if --ack is set, ack happens only after message-open succeeds.
    if message-open fails, the envelope remains unacked and the command reports failure without treating the whole command as a process failure.
    prints dev/pre-alpha warning and summary counters.

### 5.3 Tests were added

New test file:

    carbonstack-comms/internal/app/openmls_inbox_dev_test.go

Test coverage added:

    TestOpenMLSInboxDevRejectsMissingRequiredFlags
    TestOpenMLSInboxDevSkipsUnsupportedEnvelopes
    TestOpenMLSInboxDevDoesNotAckWhenMessageOpenFails
    TestOpenMLSInboxDevAcksOnlyAfterMessageOpenSuccess

Important behavior proven:

    missing required flags produce an error.
    unsupported envelopes are skipped and do not call message-open.
    message-open failure prevents ack even when --ack is set.
    ack occurs only after message-open success when --ack is explicitly set.
    tests stub inbox, ack, artifact write, and sidecar open hooks to test command glue without requiring Cargo or live Cypher.

### 5.4 README was updated

`carbonstack-comms/README.md` now documents both dev-only OpenMLS runtime commands:

    openmls-send-dev
    openmls-inbox-dev

The README note clarifies:

    commands are dev-only and pre-alpha.
    they do not replace existing stub-era send/inbox.
    openmls-send-dev calls sidecar message-protect and submits the resulting application-message artifact through Cypher.
    openmls-inbox-dev fetches the inbox, skips unsupported envelopes, writes OpenMLS artifacts, calls sidecar message-open, prints plaintext only after success, and acks only after success when --ack is set.
    this is not mature messaging UX, not local-backbone, and not a production security claim.

### 5.5 Validation passed

Commands observed passing:

    cd ~/repos/carbonstack_umbrella/carbonstack-comms
    go test ./internal/app -count=1
    go test ./... -count=1 -timeout 600s

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go run . --profile local-cypher
    go run . --profile doctor
    go run . --profile core --clean-generated

Final commit/push:

    carbonstack-comms:
    65bc707 feat: add dev OpenMLS inbox command

Final repo snapshot:

    carbonstack        a799764 docs: define runtime OpenMLS command contract
    carbonstack-comms  65bc707 feat: add dev OpenMLS inbox command
    carbonstack-cypher 9ab994c docs: point to local operator runbook
    carbonstack-os     1bbbe52 docs: clarify CarbonStackOS target direction

---

## 6. Critical Blunders / Lessons to Carry Forward

### 6.1 The runtime path is now two commands, but still not a smoke proof

`openmls-send-dev` and `openmls-inbox-dev` both exist.

However, this does not yet prove:

    a full dev CLI roundtrip from one device to another through both commands.
    a repeatable runtime smoke test.
    local-backbone.
    production messaging UX.

Correct next path:

    do a focused runtime smoke proof rung before naming anything local-backbone.

### 6.2 Ack discipline was preserved

The most important v0.4.4 invariant is:

    ack only after sidecar message-open succeeds, and only when --ack is explicitly set.

This preserves consume-success-gated ack semantics.

### 6.3 Unsupported envelopes must not poison dev inbox tests

`openmls-inbox-dev` skips unsupported envelopes instead of failing the entire command.

Reason:

    dev inboxes may still contain old stub-era messages.
    this keeps runtime OpenMLS work from breaking mixed-development environments.
    the command still clearly reports skipped unsupported envelopes.

### 6.4 Test injection seams remain useful

The v0.4.4 command followed the v0.4.3 pattern:

    add runtime glue.
    expose package-level function variables only where useful for tests.
    avoid invoking Cargo/live Cypher in app-level command tests.
    rely on lower-level protocol/relay tests for real sidecar/Cypher behavior.

### 6.5 Known OpenMLS generated roots are expected around sidecar tests

OpenMLS tests may produce:

    .carbonstack-openmls-sidecar-state
    target

The runner may see these in pre/post scans if tests were run immediately beforehand. This remains acceptable only because:

    they are classified as known OpenMLS generated roots.
    artifact scan is non-destructive.
    --clean-generated explicitly removes them.

Do not generalize cleanup beyond known generated roots.

### 6.6 Commit metadata warning remains a polish issue

The commit log again showed auto-generated local commit metadata:

    Committer: bitcrusher32 <▓▓>

This is not a functional validation failure, but it is a professionalism/publishing hygiene issue to address eventually with `git config --global user.name` and `git config --global user.email` if desired.

---

## 7. Current Blockers / Not Validated

Current blockers and nonvalidated areas after v0.4.4:

    No end-to-end dev CLI OpenMLS send/inbox roundtrip has been proven yet.
    No runtime smoke proof exists yet for openmls-send-dev -> Cypher -> openmls-inbox-dev.
    Existing send/inbox remain stub-era.
    Runtime Comms has dev-only OpenMLS commands, but no mature user-facing UX.
    local-backbone runner profile does not exist and remains reserved.
    Post-release public download verification is still useful if not already performed.
    Sanitized PRIME placement in carbonstack/logdoc-loop-system still needs a policy decision.
    Production-safe provider storage and secure vault design remain v0.5.x future work.
    PQ/hybrid ciphersuite migration readiness remains v0.5.x future work.
    Hostile-server rollback/replay/metadata-abuse validation remains future work.
    Android/CarbonStackOS implementation remains later work.
    CarbonStack Relay Space mechanics remain a future concept, not an implementation.
    Public ingress, systemd, cloudflared, real homelab deployment, and production service paths remain out of scope.

Hard nonclaims remain:

    Do not claim production readiness.
    Do not claim production E2EE.
    Do not claim hostile-server safety.
    Do not claim metadata privacy.
    Do not claim Debian production deployability.
    Do not claim systemd readiness.
    Do not claim cloudflared readiness.
    Do not claim Android readiness.
    Do not claim CarbonStackOS readiness.
    Do not claim secure vault/storage.
    Do not claim real homelab validation.
    Do not claim external audit or certification.
    Do not claim CarbonStack Relay Space mechanics are implemented.
    Do not claim local-backbone exists or is justified yet.
    Do not claim local-cypher validates runtime Comms UX.
    Do not claim full validates local-backbone.
    Do not claim full is a deployment command.
    Do not claim send/inbox are OpenMLS-backed yet.
    Do not claim PQ/hybrid security or quantum-safe messaging.
    Do not let GitHub mirrors supersede Gitea source-of-truth status.

Allowed claim:

    CarbonStackComms now has dev-only openmls-send-dev and openmls-inbox-dev commands.
    openmls-send-dev calls sidecar message-protect and submits the resulting application-message artifact through internal/relay.SubmitOpenMLSArtifactEnvelope.
    openmls-inbox-dev fetches the current device inbox through Cypher, skips unsupported envelopes, writes OpenMLS application-message artifacts through internal/relay.WriteOpenMLSArtifactFromEnvelope, calls sidecar message-open, prints plaintext after sidecar success, and only acks after sidecar success when --ack is explicitly set.
    Existing send/inbox remain stub-era.
    v0.4.0 remains the current public broad local deployability pre-release.
    Gitea remains source of truth; GitHub mirrors are secondary push mirrors.

---

## 8. Next Safest Actions

Recommended next rung:

    v0.4.5 dev runtime smoke proof

Focus:

    prove a local dev runtime path using the new Comms CLI commands:
      openmls-send-dev
      openmls-inbox-dev

Likely approach:

    recon first to determine exact setup prerequisites.
    use current OpenMLS sidecar dev setup commands for identities/conversation/member join.
    run a local Cypher server.
    create/claim/register devices as needed.
    send an OpenMLS application-message through openmls-send-dev.
    retrieve/open it through openmls-inbox-dev.
    optionally ack only after successful open.
    document exactly what worked and what was still manual/dev-only.

Expected implementation repo:

    maybe carbonstack-comms only, if smoke proof exposes small helper/docs needs.
    maybe carbonstack docs only, if it is purely a recorded smoke-run.
    avoid runner profile unless the smoke proof becomes repeatable enough.

Avoid next:

    replacing send/inbox
    adding local-backbone prematurely
    claiming whole-path maturity before smoke proof
    public ingress
    cloudflared
    systemd
    real homelab deployment
    Android app
    CarbonStackOS implementation
    CarbonStack Relay Space implementation
    PQ/hybrid ciphersuite implementation
    production/security claims
    broad negative-path suite unless the runtime smoke exposes a direct blocker

---

## 9. Critical Paths / Functions

### Source-of-truth public surfaces

    [GITEA SOURCE OF TRUTH] https://git.bitcrusher32.win/bitcrusher32/carbonstack
    [RELEASES] https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases
    [CURRENT PUBLIC RELEASE] https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.4.0

GitHub mirrors:

    secondary push mirrors for discoverability/redundancy only.
    Do not treat mirrors as release authority unless project policy changes.

### WSL Debian active working umbrella

    [WSL-DEBIAN] umbrella: ~/repos/carbonstack_umbrella
    [WSL-DEBIAN] carbonstack: ~/repos/carbonstack_umbrella/carbonstack
    [WSL-DEBIAN] carbonstack-comms: ~/repos/carbonstack_umbrella/carbonstack-comms
    [WSL-DEBIAN] carbonstack-cypher: ~/repos/carbonstack_umbrella/carbonstack-cypher
    [WSL-DEBIAN] carbonstack-os: ~/repos/carbonstack_umbrella/carbonstack-os
    [GO-RUNNER] runner dir: ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate

### Runtime integration doc surfaces

    [DOC] v0.4.1 runtime recon:
    carbonstack/docs/155-runtime-comms-openmls-cypher-recon-v0.md

    [DOC] v0.4.2 command contract:
    carbonstack/docs/156-runtime-openmls-command-contract-v0.md

    [DOC] docs index:
    carbonstack/docs/README.md

    [DOC] roadmap:
    carbonstack/roadmap/ROADMAP.md

### Comms runtime surfaces

    [COMMS] CLI entry:
    carbonstack-comms/cmd/comms/main.go

    [COMMS] command implementation:
    carbonstack-comms/internal/app/commands.go

    [COMMS] openmls-send-dev tests:
    carbonstack-comms/internal/app/openmls_send_dev_test.go

    [COMMS] openmls-inbox-dev tests:
    carbonstack-comms/internal/app/openmls_inbox_dev_test.go

    [COMMS] Cypher client:
    carbonstack-comms/internal/client/cypher.go

    [COMMS] relay helpers:
    carbonstack-comms/internal/relay/cypher_bridge.go
    carbonstack-comms/internal/relay/openmls_artifacts.go

    [COMMS] OpenMLS sidecar dir:
    carbonstack-comms/internal/protocol/mls/openmls-sidecar

    [COMMS] sidecar README:
    carbonstack-comms/internal/protocol/mls/openmls-sidecar/README.md

### Current stub-era Comms functions / behaviors

    send:
      trust policy evaluation
      mock/stub provider encryption
      Cypher envelope submit

    inbox:
      Cypher inbox retrieval
      mock/stub provider plaintext display as stub_plaintext

    ack:
      envelope ack by envelope ID and recipient device ID

### Implemented v0.4.x dev OpenMLS commands

    openmls-send-dev:
      state.RequireReadyDevice
      trust.EvaluateSend
      runOpenMLSMessageProtect
      sidecar message-protect
      internal/relay.SubmitOpenMLSArtifactEnvelope
      ArtifactKindApplicationMessage
      dev/pre-alpha warning

    openmls-inbox-dev:
      state.RequireReadyDevice
      Cypher inbox
      skip unsupported envelopes
      internal/relay.WriteOpenMLSArtifactFromEnvelope
      runOpenMLSMessageOpen
      sidecar message-open
      print plaintext after sidecar success
      optional ack only after sidecar success

### OpenMLS sidecar commands relevant to v0.4.x

    provider-info
    identity-create
    identity-status
    public-bundle-export
    conversation-create
    conversation-load-check
    conversation-add-member
    conversation-join
    message-protect
    message-open

### Relay/API functions relevant to runtime integration

    internal/client.CypherClient.SubmitEnvelope
    internal/client.CypherClient.Inbox
    internal/client.CypherClient.AckEnvelope
    internal/relay.SubmitOpenMLSArtifactEnvelope
    internal/relay.WriteOpenMLSArtifactFromEnvelope

### Existing v0.4.0 package validation

    cd <package-root>/carbonstack/tools/carbonstack-validate
    go run . --profile verify-checksums --root <package-root>
    go run . --profile full --root <package-root> --clean-generated

### Live umbrella validation

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go run . --profile local-cypher
    go run . --profile doctor
    go run . --profile core --clean-generated

### Envelope/protocol constants

Accepted opaque OpenMLS application-message pair:

    content_type=carbonstack.mls.application-message.v0
    protocol_version=carbonstack-openmls-sidecar-v0

Historical invalid pair validated by local-cypher:

    content_type=carbonstack.message.text.stub.v0
    protocol_version=carbonstack-openmls-sidecar-v0

Expected result:

    HTTP 400
    error code unsupported_protocol_version

---

## 10. v0.4.x Forward Plan

### v0.4.5 — dev runtime smoke proof

Expected:

    local dev runtime proof that uses Comms CLI commands against Cypher/OpenMLS sidecar.
    still dev/pre-alpha.
    likely manual or docs-first at first.
    not local-backbone unless whole-path behavior is meaningful and repeatable enough.

### v0.4.6+ — local-backbone decision point

Only consider `local-backbone` when:

    runtime Comms send/inbox path is actually OpenMLS/Cypher-backed.
    state expectations are documented.
    validation can run the whole path repeatably.
    name does not overclaim production/deployment maturity.

### v0.4.x later polish

Likely needs:

    exact sidecar setup runbook.
    clearer dev state reset semantics.
    better command output shape.
    possible helper extraction if commands.go becomes too large.
    eventual decision on whether stub-era send/inbox are aliased, replaced, or deprecated.

---

## 11. v0.5.x and Later Carry-Forward

### v0.5.x

Local storage, trust state, provider state, vault design, and PQ/hybrid ciphersuite migration readiness.

Do not start PQ/hybrid implementation before the runtime OpenMLS/Cypher path exists. PQ/hybrid migration touches provider state, ciphersuite metadata, trust display, artifact size, reinit/rekey/recovery policy, and claim boundaries.

### v0.6.x+

Hostile-server and abuse-resistance harnesses.

### v0.7.x+

Deployability and operations hardening.

### v0.8.x+

Documented self-pentest / adversarial validation.

### v0.9.x+

Claim-boundary review.

### v0.10.x+

Android backend/app work.

### v1.x.x

Public app/server major epoch and much later CarbonStackOS/appliance research, if justified.

---

## 12. Lean Breakpoint Summary

    CarbonStack v0.4.4 is the first dev-only OpenMLS inbox/open/optional-ack command checkpoint. v0.4.0 remains the current public broad local deployability pre-release. carbonstack remains at a799764 docs: define runtime OpenMLS command contract. carbonstack-comms is now at 65bc707 feat: add dev OpenMLS inbox command. v0.4.4 implements openmls-inbox-dev in carbonstack-comms while leaving existing send/inbox/ack intact. openmls-inbox-dev loads ready state, fetches Cypher inbox for the current device, skips unsupported envelopes, writes OpenMLS application-message artifacts with relay.WriteOpenMLSArtifactFromEnvelope, calls sidecar message-open, prints plaintext only after sidecar success, and only acks after sidecar success when --ack is explicitly set. Tests prove missing required flags, unsupported envelope skip, no ack when message-open fails, and ack only after message-open success. Validation passed for app tests, comms package tests, local-cypher, doctor, and core --clean-generated. Next safe rung: v0.4.5 dev runtime smoke proof using openmls-send-dev -> Cypher -> openmls-inbox-dev.


---

## 13. Preserved Immediate Previous Handoff: v0.4.3

The following is the previous v0.4.3 handoff. Where it conflicts with the v0.4.4 overlay above, v0.4.4 wins for current state.


# CarbonStack LogDoc v0.4.3

**Last updated:** 2026-06-03 local session  
**Owner/operator:** bitcrusher32  
**Assistant/model:** GPT-5.5 Thinking  
**Current phase:** Phase 2F / **v0.4.x first dev-only OpenMLS send command checkpoint complete**. The v0.4.0 public broad local deployability pre-release remains the current public release. v0.4.1 identified the runtime seam; v0.4.2 defined the runtime OpenMLS command contract; v0.4.3 now implements the first dev-only OpenMLS send-side command in `carbonstack-comms`. `carbonstack` remains at `a799764 (HEAD -> main, origin/main, origin/HEAD) docs: define runtime OpenMLS command contract`; `carbonstack-comms` is now at `39fb287 (HEAD -> main, origin/main, origin/HEAD) feat: add dev OpenMLS send command`; `carbonstack-cypher` remains at `9ab994c (HEAD -> main, origin/main, origin/HEAD) docs: point to local operator runbook`; `carbonstack-os` remains at `1bbbe52 (HEAD -> main, origin/main, origin/HEAD) docs: clarify CarbonStackOS target direction`.

**Version schema:** `v[scope].[timeline]`. This file is `v0.4.3`, the first implementation checkpoint after the v0.4.2 runtime OpenMLS command contract. Per LogDoc V2 practice, this Markdown file preserves current operational continuity and carries the v0.4.x runtime-integration thread forward; the JSON breakpoint remains the lean machine-readable handoff state.

**Compression note:** The detailed v0.3.0 through v0.3.36PRIME process ledger remains archived in PRIME/sanitized PRIME. v0.4.0 compressed that pre-release sprint; v0.4.1 preserved the runtime recon; v0.4.2 preserved the command-contract decision; v0.4.3 now preserves the first dev-only OpenMLS send-side implementation. Future work should continue from this v0.4.x runtime thread.

---

## 1. Project Goal

**Active goal:** Move CarbonStackComms from stub-era runtime `send`/`inbox` behavior toward explicit dev-only OpenMLS/Cypher-backed runtime commands. v0.4.3 completed the first implementation rung: `openmls-send-dev` now exists in `carbonstack-comms`.

CarbonStack’s broader project intent remains:

- **CarbonStack** is the secure-communications appliance-stack project and shared doctrine/release authority.
- **CarbonStackComms** is the text-first messaging client and current OpenMLS sidecar/relay integration surface.
- **CarbonStackCypher** is the self-hostable relay/storage stack.
- **CarbonStackOS** is future constrained Android-derived appliance OS work and remains out of current runnable validation packages.
- The `carbonstack` Gitea repo remains the source of truth for release state, release assets, docs, and public claims.
- GitHub mirrors exist as secondary push mirrors only; Gitea remains authoritative.

**North star:**

> Its goal is not to be a better smartphone. Its goal is to stop being a smartphone.

**Core doctrine:**

> Every feature is guilty until it proves it does not add unacceptable parser, network, sensor, identity, or filesystem authority.

**Current reality after v0.4.3:** `openmls-send-dev` is implemented and tested, but `openmls-inbox-dev` is not implemented yet. The existing stub-era `send`, `inbox`, and `ack` commands remain in place.

Correct interpretation:

    openmls-send-dev is a dev-only OpenMLS send-side command.
    It calls the OpenMLS sidecar message-protect path and submits the resulting application-message artifact through the Cypher relay helper.
    Existing send/inbox remain stub-era and must not be claimed OpenMLS-backed yet.
    There is still no full runtime OpenMLS send/inbox/ack proof because the receive/open/ack command is not implemented yet.
    local-backbone remains reserved until whole-path runtime validation exists.

Hard nonclaims remain: v0.4.3 does **not** prove production readiness, production E2EE, hostile-server safety, metadata privacy, Android readiness, CarbonStackOS readiness, secure vault/storage, public ingress safety, real homelab validation, audit, certification, PQ/hybrid security, mature messenger UX, or local-backbone readiness.

---

## 2. Current Public Release State

Public release:

    https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.4.0

Release title:

    CarbonStack v0.4.0 Broad Local Deployability Pre-Release

Release tag:

    v0.4.0

Release commit:

    24ac9fc494 / 24ac9fc docs: record v0.4.0 package rehearsal

Gitea release status:

    Pre-Release

Release assets attached:

    carbonstack-v0.4.0-broad-local-deployability-pre-release.tgz
    carbonstack-v0.4.0-release-manifest.json
    carbonstack-v0.4.0-package-checksums.txt
    carbonstack-v0.4.0-asset-checksums.txt
    carbonstack-v0.4.0-validation-freeze.md
    v0.4.0-testing-runbook.md
    v0.4.0-release-notes.md
    LICENSE

Important release-page warning:

    The default Gitea Source Code ZIP/TAR.GZ downloads are auto-generated carbonstack-only archives.
    The intended runnable validation artifact is the attached multi-repo release package plus companion assets.

Post-release public surface:

    carbonstack README is evergreen.
    It points interested parties to the Releases page for runnable / known-good artifacts.
    It states Gitea is the source of truth.
    GitHub mirrors are secondary push mirrors, not release authority.

---

## 3. Current Repo Heads

    carbonstack        a799764 (HEAD -> main, origin/main, origin/HEAD) docs: define runtime OpenMLS command contract
    carbonstack-comms  39fb287 (HEAD -> main, origin/main, origin/HEAD) feat: add dev OpenMLS send command
    carbonstack-cypher 9ab994c (HEAD -> main, origin/main, origin/HEAD) docs: point to local operator runbook
    carbonstack-os     1bbbe52 (HEAD -> main, origin/main, origin/HEAD) docs: clarify CarbonStackOS target direction

Current public release tag:

    v0.4.0 -> 24ac9fc494 / 24ac9fc docs: record v0.4.0 package rehearsal

Post-release public-surface cleanup:

    6ba7aab docs: make front readme evergreen

v0.4.1 runtime recon commit:

    83d2f0f docs: record runtime OpenMLS Comms recon

v0.4.2 command contract commit:

    a799764 docs: define runtime OpenMLS command contract

Current v0.4.3 implementation commit:

    39fb287 feat: add dev OpenMLS send command

Continuity note:

    `39fb287` is a `carbonstack-comms` implementation commit after the `v0.4.0` public release tag. It does not retag v0.4.0. The release remains attached to `24ac9fc494`.

---

## 4. Current Validated / Known State

Validated / known at the v0.4.3 checkpoint:

    [RELEASE] v0.4.0 Gitea release exists, is marked Pre-Release, and is attached to 24ac9fc494.
    [PACKAGE] v0.4.0 package includes carbonstack, carbonstack-comms, carbonstack-cypher, and release metadata.
    [PACKAGE] carbonstack-os is deliberately excluded from the runnable package.
    [RUNNER] full means release-snapshot followed by local-cypher.
    [RUNNER] release-snapshot validates package layout, metadata, strict pre-test artifact scan, package checksums, and core.
    [RUNNER] core validates doctor, OpenMLS real-Cypher lifecycle, Comms package tests, Cypher package tests, and artifact scans.
    [RUNNER] local-cypher validates Cypher-only lifecycle, restart/persistence against the same temporary DB, positive opaque OpenMLS envelope/ack flow, and historical negative stub-text/OpenMLS protocol pairing.
    [DOCS] main carbonstack README is evergreen.
    [DOCS] docs/155-runtime-comms-openmls-cypher-recon-v0.md records the v0.4.1 runtime recon.
    [DOCS] docs/156-runtime-openmls-command-contract-v0.md records the v0.4.2 command contract.
    [COMMS] openmls-send-dev exists in carbonstack-comms.
    [COMMS] existing send/inbox/ack remain in place.
    [COMMS] openmls-inbox-dev is not implemented yet.
    [MIRRORS] all four repos have GitHub push-mirror functionality, with Gitea remaining source of truth.

v0.4.3 validation result from the final log:

    go test ./internal/app -count=1 passed.
    go test ./... -count=1 -timeout 600s passed in carbonstack-comms.
    local-cypher passed.
    doctor passed.
    core --clean-generated passed.
    OpenMLS real-Cypher lifecycle passed during core.
    carbonstack-comms package tests passed during core.
    carbonstack-cypher package tests passed during core.
    --clean-generated removed known OpenMLS sidecar generated roots.
    final repo status showed carbonstack-comms at 39fb287 and all four repos clean.

Observed validation nuance:

    local-cypher and core saw known OpenMLS sidecar generated roots in pre/post artifact scans because OpenMLS sidecar tests had just been run.
    The runner classified them as known OpenMLS generated roots.
    core --clean-generated removed .carbonstack-openmls-sidecar-state and target.
    This is not a release failure or source hygiene failure as long as hits remain limited to known generated roots and cleanup is explicit.

Expected v0.4.0 public package validation commands remain:

    cd <package-root>/carbonstack/tools/carbonstack-validate
    go run . --profile verify-checksums --root <package-root>
    go run . --profile full --root <package-root> --clean-generated

Live umbrella validation commands remain:

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go run . --profile local-cypher
    go run . --profile doctor
    go run . --profile core --clean-generated

---

## 5. v0.4.3 Work Completed

### 5.1 Focused implementation recon was performed

The v0.4.3 recon inspected `carbonstack-comms` command, state, trust, client, relay, and sidecar surfaces before implementation.

Useful findings:

    cmd/comms/main.go simply delegates to app.Run.
    internal/app/commands.go is simple and easy to extend.
    existing send uses trust.EvaluateSend, mock/stub crypto, and client.SubmitEnvelope.
    state.RequireReadyDevice supplies the current server URL and sender device ID.
    trust.EvaluateSend already supports dev-mode default behavior with optional strict mode.
    internal/relay.SubmitOpenMLSArtifactEnvelope already maps application-message artifacts to the correct content_type/protocol_version pair.
    the OpenMLS sidecar message-protect path returns message_artifact_path_hint.
    lower-level tests already prove OpenMLS artifact submission through Cypher.

The first patch attempt failed because a marker block did not match the exact local file formatting. The work stopped correctly, then resumed with a safer targeted patch that did not depend on a full import-block match.

### 5.2 openmls-send-dev was implemented

`carbonstack-comms/internal/app/commands.go` now includes:

    openmls-send-dev in the app.Run switch.
    openmls-send-dev in usage output.
    cmdOpenMLSSendDev.
    runOpenMLSMessageProtect.
    parseOpenMLSProtectEnvelope.
    package-level function variables for sidecar protect and relay submit injection in tests.

Command shape:

    go run ./cmd/comms openmls-send-dev \
      --to-device <recipient-cypher-device-id> \
      --sidecar-device-label <sender-sidecar-device-label> \
      --conversation <sidecar-conversation-label> \
      --message <plaintext> \
      [--message-label <label>] \
      [--strict]

Implemented behavior:

    requires --to-device, --message, --sidecar-device-label, and --conversation.
    uses --message only; no --message-file yet.
    uses dev-mode trust behavior by default.
    supports optional --strict.
    warns when running in dev trust mode.
    blocks if trust policy disallows send.
    calls OpenMLS sidecar message-protect through cargo.
    requires message_artifact_path_hint from the sidecar.
    resolves relative artifact hints against the sidecar dir.
    submits the artifact through internal/relay.SubmitOpenMLSArtifactEnvelope as ArtifactKindApplicationMessage.
    prints envelope and OpenMLS metadata.
    labels output as dev/pre-alpha runtime path.

### 5.3 Tests were added

New test file:

    carbonstack-comms/internal/app/openmls_send_dev_test.go

Test coverage added:

    TestOpenMLSSendDevRejectsMissingRequiredFlags
    TestOpenMLSSendDevDoesNotSubmitWhenProtectFails
    TestOpenMLSSendDevSubmitsApplicationMessageArtifact

Important behavior proven:

    missing required flags produce an error.
    sidecar protect failure prevents Cypher submit.
    on sidecar success, the command submits ArtifactKindApplicationMessage through the relay helper with sender/recipient/server fields wired from state and flags.

Implementation note:

    tests stub `runOpenMLSMessageProtectForCommand` and `submitOpenMLSArtifactEnvelopeForCommand`, so command glue is tested without invoking Cargo or a live Cypher server.

### 5.4 README was updated

`carbonstack-comms/README.md` now documents the dev-only OpenMLS runtime command.

The README note clarifies:

    openmls-send-dev is dev-only and pre-alpha.
    it does not replace existing stub-era send.
    it calls sidecar message-protect.
    it submits the resulting application-message artifact through the Cypher relay helper.
    it preserves dev-mode trust by default with optional --strict.
    it is not mature messaging UX, local-backbone, or a production security claim.

### 5.5 Validation passed

Commands observed passing:

    cd ~/repos/carbonstack_umbrella/carbonstack-comms
    go test ./internal/app -count=1
    go test ./... -count=1 -timeout 600s

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go run . --profile local-cypher
    go run . --profile doctor
    go run . --profile core --clean-generated

Final commit/push:

    carbonstack-comms:
    39fb287 feat: add dev OpenMLS send command

Commit nuance:

    The first push failed due authentication after commit metadata warnings. A follow-up `git push` succeeded and advanced origin/main from 012c8bf to 39fb287.

Final repo snapshot:

    carbonstack        a799764 docs: define runtime OpenMLS command contract
    carbonstack-comms  39fb287 feat: add dev OpenMLS send command
    carbonstack-cypher 9ab994c docs: point to local operator runbook
    carbonstack-os     1bbbe52 docs: clarify CarbonStackOS target direction

---

## 6. Critical Blunders / Lessons to Carry Forward

### 6.1 Marker-based patching can fail on formatting

The first implementation script failed with:

    expected marker block not found

This was a patch fragility issue, not a project issue.

Lesson:

    for live code edits, prefer smaller targeted replacements or AST-ish helpers over exact full-block imports.
    stop immediately when markers fail.
    inspect or use safer patch strategy before continuing.

### 6.2 Testable runtime glue beats importing test helpers

The sidecar helper logic already existed in tests, but the implementation correctly added production runtime glue instead of importing test-only helpers.

Lesson:

    add small production seams.
    expose function variables only where useful for test injection.
    do not pull test-only helpers into runtime code.

### 6.3 Send-side implementation is not full messaging

`openmls-send-dev` only proves the send side:

    plaintext -> sidecar message-protect -> artifact -> Cypher relay helper submit

It does not prove:

    inbox retrieval
    message-open
    consume-success-gated ack
    full CLI runtime roundtrip
    local-backbone

### 6.4 Existing send/inbox remain stub-era

Do not let the presence of `openmls-send-dev` imply existing `send` or `inbox` are upgraded.

Correct claim:

    openmls-send-dev exists as a dev-only OpenMLS send-side command.
    send/inbox are still legacy/stub-era.
    openmls-inbox-dev is next.

### 6.5 Known OpenMLS generated roots are expected around sidecar tests

OpenMLS tests may produce:

    .carbonstack-openmls-sidecar-state
    target

The runner may see these in pre/post scans if tests were run immediately beforehand. This remains acceptable only because:

    they are classified as known OpenMLS generated roots.
    artifact scan is non-destructive.
    --clean-generated explicitly removes them.

Do not generalize cleanup beyond known generated roots.

---

## 7. Current Blockers / Not Validated

Current blockers and nonvalidated areas after v0.4.3:

    openmls-inbox-dev is not implemented yet.
    No runtime receive/open/ack command exists yet.
    No end-to-end dev CLI OpenMLS send/inbox roundtrip exists yet.
    Existing send/inbox remain stub-era.
    Runtime Comms still does not have a whole-path OpenMLS CLI proof.
    Ack consume-success behavior is specified but not implemented in runtime command code yet.
    local-backbone runner profile does not exist and remains reserved.
    Post-release public download verification is still useful if not already performed.
    Sanitized PRIME placement in carbonstack/logdoc-loop-system still needs a policy decision.
    Production-safe provider storage and secure vault design remain v0.5.x future work.
    PQ/hybrid ciphersuite migration readiness remains v0.5.x future work.
    Hostile-server rollback/replay/metadata-abuse validation remains future work.
    Android/CarbonStackOS implementation remains later work.
    CarbonStack Relay Space mechanics remain a future concept, not an implementation.
    Public ingress, systemd, cloudflared, real homelab deployment, and production service paths remain out of scope.

Hard nonclaims remain:

    Do not claim production readiness.
    Do not claim production E2EE.
    Do not claim hostile-server safety.
    Do not claim metadata privacy.
    Do not claim Debian production deployability.
    Do not claim systemd readiness.
    Do not claim cloudflared readiness.
    Do not claim Android readiness.
    Do not claim CarbonStackOS readiness.
    Do not claim secure vault/storage.
    Do not claim real homelab validation.
    Do not claim external audit or certification.
    Do not claim CarbonStack Relay Space mechanics are implemented.
    Do not claim local-backbone exists or is justified yet.
    Do not claim local-cypher validates runtime Comms UX.
    Do not claim full validates local-backbone.
    Do not claim full is a deployment command.
    Do not claim send/inbox are OpenMLS-backed yet.
    Do not claim openmls-inbox-dev is implemented yet.
    Do not claim PQ/hybrid security or quantum-safe messaging.
    Do not let GitHub mirrors supersede Gitea source-of-truth status.

Allowed claim:

    CarbonStackComms now has a dev-only openmls-send-dev command.
    openmls-send-dev calls sidecar message-protect and submits the resulting application-message artifact through internal/relay.SubmitOpenMLSArtifactEnvelope.
    openmls-send-dev is tested for missing required flags, no-submit-on-protect-failure, and successful application-message artifact submit through the injected relay seam.
    Existing send/inbox remain stub-era.
    v0.4.0 remains the current public broad local deployability pre-release.
    Gitea remains source of truth; GitHub mirrors are secondary push mirrors.

---

## 8. Next Safest Actions

Recommended next rung:

    v0.4.4 first dev-only OpenMLS inbox/open/ack command

Focus:

    implement openmls-inbox-dev in carbonstack-comms.
    keep it explicit and dev-only.
    fetch Cypher inbox for current device ID.
    filter/select OpenMLS application-message envelopes.
    validate content_type and protocol_version through relay helper.
    write artifact using internal/relay.WriteOpenMLSArtifactFromEnvelope.
    call sidecar message-open.
    print plaintext only after sidecar success.
    if --ack is set, ack only after sidecar message-open succeeds.
    add minimal tests proving no ack before sidecar success and ack after success.
    update carbonstack-comms README/docs as needed.

Expected implementation repo:

    carbonstack-comms

Expected docs/release authority repo updates if needed:

    carbonstack

Avoid next:

    replacing send/inbox
    adding local-backbone
    implementing broad runtime smoke before inbox/open/ack exists
    public ingress
    cloudflared
    systemd
    real homelab deployment
    Android app
    CarbonStackOS implementation
    CarbonStack Relay Space implementation
    PQ/hybrid ciphersuite implementation
    production/security claims
    broad negative-path suite unless the runtime seam exposes a direct blocker

---

## 9. Critical Paths / Functions

### Source-of-truth public surfaces

    [GITEA SOURCE OF TRUTH] https://git.bitcrusher32.win/bitcrusher32/carbonstack
    [RELEASES] https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases
    [CURRENT PUBLIC RELEASE] https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.4.0

GitHub mirrors:

    secondary push mirrors for discoverability/redundancy only.
    Do not treat mirrors as release authority unless project policy changes.

### WSL Debian active working umbrella

    [WSL-DEBIAN] umbrella: ~/repos/carbonstack_umbrella
    [WSL-DEBIAN] carbonstack: ~/repos/carbonstack_umbrella/carbonstack
    [WSL-DEBIAN] carbonstack-comms: ~/repos/carbonstack_umbrella/carbonstack-comms
    [WSL-DEBIAN] carbonstack-cypher: ~/repos/carbonstack_umbrella/carbonstack-cypher
    [WSL-DEBIAN] carbonstack-os: ~/repos/carbonstack_umbrella/carbonstack-os
    [GO-RUNNER] runner dir: ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate

### Runtime integration doc surfaces

    [DOC] v0.4.1 runtime recon:
    carbonstack/docs/155-runtime-comms-openmls-cypher-recon-v0.md

    [DOC] v0.4.2 command contract:
    carbonstack/docs/156-runtime-openmls-command-contract-v0.md

    [DOC] docs index:
    carbonstack/docs/README.md

    [DOC] roadmap:
    carbonstack/roadmap/ROADMAP.md

### Comms runtime surfaces

    [COMMS] CLI entry:
    carbonstack-comms/cmd/comms/main.go

    [COMMS] command implementation:
    carbonstack-comms/internal/app/commands.go

    [COMMS] openmls-send-dev tests:
    carbonstack-comms/internal/app/openmls_send_dev_test.go

    [COMMS] Cypher client:
    carbonstack-comms/internal/client/cypher.go

    [COMMS] relay helpers:
    carbonstack-comms/internal/relay/cypher_bridge.go
    carbonstack-comms/internal/relay/openmls_artifacts.go

    [COMMS] OpenMLS sidecar dir:
    carbonstack-comms/internal/protocol/mls/openmls-sidecar

    [COMMS] sidecar README:
    carbonstack-comms/internal/protocol/mls/openmls-sidecar/README.md

### Current stub-era Comms functions / behaviors

    send:
      trust policy evaluation
      mock/stub provider encryption
      Cypher envelope submit

    inbox:
      Cypher inbox retrieval
      mock/stub provider plaintext display as stub_plaintext

    ack:
      envelope ack by envelope ID and recipient device ID

### Implemented v0.4.3 dev send function

    openmls-send-dev:
      state.RequireReadyDevice
      trust.EvaluateSend
      runOpenMLSMessageProtect
      sidecar message-protect
      internal/relay.SubmitOpenMLSArtifactEnvelope
      ArtifactKindApplicationMessage
      dev/pre-alpha warning

### Planned v0.4.4 dev inbox/open/ack function

    openmls-inbox-dev:
      Cypher inbox
      filter/select OpenMLS application-message envelopes
      internal/relay.WriteOpenMLSArtifactFromEnvelope
      sidecar message-open
      optional ack after message-open success

### OpenMLS sidecar commands relevant to v0.4.x

    provider-info
    identity-create
    identity-status
    public-bundle-export
    conversation-create
    conversation-load-check
    conversation-add-member
    conversation-join
    message-protect
    message-open

### Relay/API functions relevant to runtime integration

    internal/client.CypherClient.SubmitEnvelope
    internal/client.CypherClient.Inbox
    internal/client.CypherClient.AckEnvelope
    internal/relay.SubmitOpenMLSArtifactEnvelope
    internal/relay.WriteOpenMLSArtifactFromEnvelope

### Existing v0.4.0 package validation

    cd <package-root>/carbonstack/tools/carbonstack-validate
    go run . --profile verify-checksums --root <package-root>
    go run . --profile full --root <package-root> --clean-generated

### Live umbrella validation

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go run . --profile local-cypher
    go run . --profile doctor
    go run . --profile core --clean-generated

### Envelope/protocol constants

Accepted opaque OpenMLS application-message pair:

    content_type=carbonstack.mls.application-message.v0
    protocol_version=carbonstack-openmls-sidecar-v0

Historical invalid pair validated by local-cypher:

    content_type=carbonstack.message.text.stub.v0
    protocol_version=carbonstack-openmls-sidecar-v0

Expected result:

    HTTP 400
    error code unsupported_protocol_version

---

## 10. v0.4.x Forward Plan

### v0.4.4 — first dev-only OpenMLS inbox/open/ack command

Implementation target:

    openmls-inbox-dev

Expected behavior:

    query Cypher inbox
    write OpenMLS artifact from envelope
    call sidecar message-open
    print plaintext only after sidecar success
    ack only after sidecar success when --ack is set

### v0.4.5 — runtime smoke proof

Expected:

    local dev runtime proof that uses Comms CLI commands against Cypher/OpenMLS sidecar
    still dev/pre-alpha
    not local-backbone unless whole-path behavior is meaningful enough

### v0.4.6+ — local-backbone decision point

Only consider `local-backbone` when:

    runtime Comms send/inbox path is actually OpenMLS/Cypher-backed
    state expectations are documented
    validation can run the whole path repeatably
    name does not overclaim production/deployment maturity

---

## 11. v0.5.x and Later Carry-Forward

### v0.5.x

Local storage, trust state, provider state, vault design, and PQ/hybrid ciphersuite migration readiness.

Do not start PQ/hybrid implementation before the runtime OpenMLS/Cypher path exists. PQ/hybrid migration touches provider state, ciphersuite metadata, trust display, artifact size, reinit/rekey/recovery policy, and claim boundaries.

### v0.6.x+

Hostile-server and abuse-resistance harnesses.

### v0.7.x+

Deployability and operations hardening.

### v0.8.x+

Documented self-pentest / adversarial validation.

### v0.9.x+

Claim-boundary review.

### v0.10.x+

Android backend/app work.

### v1.x.x

Public app/server major epoch and much later CarbonStackOS/appliance research, if justified.

---

## 12. Lean Breakpoint Summary

    CarbonStack v0.4.3 is the first dev-only OpenMLS send command checkpoint. v0.4.0 remains the current public broad local deployability pre-release. carbonstack remains at a799764 docs: define runtime OpenMLS command contract. carbonstack-comms is now at 39fb287 feat: add dev OpenMLS send command. v0.4.3 implements openmls-send-dev in carbonstack-comms while leaving existing send/inbox/ack intact. openmls-send-dev uses dev-mode trust by default with optional --strict, calls sidecar message-protect, requires a message_artifact_path_hint, resolves the artifact path, and submits it as an OpenMLS application-message artifact through internal/relay.SubmitOpenMLSArtifactEnvelope. Tests prove missing required flags, no submit when protect fails, and successful application-message artifact submit through injected seams. Validation passed for app tests, comms package tests, local-cypher, doctor, and core --clean-generated. Next safe rung: v0.4.4 implement openmls-inbox-dev with sidecar message-open and ack only after consume success.


---

## 13. Preserved Immediate Previous Handoff: v0.4.2

The following is the previous v0.4.2 handoff. Where it conflicts with the v0.4.3 overlay above, v0.4.3 wins for current state.


# CarbonStack LogDoc v0.4.2

**Last updated:** 2026-06-03 local session  
**Owner/operator:** bitcrusher32  
**Assistant/model:** GPT-5.5 Thinking  
**Current phase:** Phase 2F / **v0.4.x runtime OpenMLS command contract checkpoint complete**. The v0.4.0 public broad local deployability pre-release remains the current public release. v0.4.1 identified the runtime seam: CarbonStackComms has `send` / `inbox` / `ack`, but `send` and `inbox` remain stub-era while OpenMLS/Cypher relay is validated underneath through lower-level seams. v0.4.2 now locks in the first runtime OpenMLS command contract. `carbonstack` is now at `a799764 (HEAD -> main, origin/main, origin/HEAD) docs: define runtime OpenMLS command contract`. `carbonstack-comms` remains at `012c8bf (HEAD -> main, origin/main, origin/HEAD) scripts: support source snapshot artifact guard`; `carbonstack-cypher` remains at `9ab994c (HEAD -> main, origin/main, origin/HEAD) docs: point to local operator runbook`; `carbonstack-os` remains at `1bbbe52 (HEAD -> main, origin/main, origin/HEAD) docs: clarify CarbonStackOS target direction`.

**Version schema:** `v[scope].[timeline]`. This file is `v0.4.2`, the runtime OpenMLS command-contract handoff after `docs/156-runtime-openmls-command-contract-v0.md` landed. Per LogDoc V2 practice, this Markdown file preserves current operational continuity and carries the v0.4.x runtime-integration thread forward; the JSON breakpoint remains the lean machine-readable handoff state.

**Compression note:** The detailed v0.3.0 through v0.3.36PRIME process ledger remains archived in PRIME/sanitized PRIME. v0.4.0 compressed that pre-release sprint; v0.4.1 preserved the runtime recon; v0.4.2 now preserves the command-contract decision. Future work should continue from this v0.4.x thread, not from the old v0.3.x release-packaging frame.

---

## 1. Project Goal

**Active goal:** Move CarbonStackComms from stub-era runtime `send`/`inbox` behavior toward explicit dev-only OpenMLS/Cypher-backed runtime commands. v0.4.2 completed the contract-first rung. The next safe work is implementation of the first dev-only OpenMLS send command in `carbonstack-comms`.

CarbonStack’s broader project intent remains:

- **CarbonStack** is the secure-communications appliance-stack project and shared doctrine/release authority.
- **CarbonStackComms** is the text-first messaging client and current OpenMLS sidecar/relay integration surface.
- **CarbonStackCypher** is the self-hostable relay/storage stack.
- **CarbonStackOS** is future constrained Android-derived appliance OS work and remains out of current runnable validation packages.
- The `carbonstack` Gitea repo remains the source of truth for release state, release assets, docs, and public claims.
- GitHub mirrors exist as secondary push mirrors only; Gitea remains authoritative.

**North star:**

> Its goal is not to be a better smartphone. Its goal is to stop being a smartphone.

**Core doctrine:**

> Every feature is guilty until it proves it does not add unacceptable parser, network, sensor, identity, or filesystem authority.

**Current reality after v0.4.2:** v0.4.2 does not implement runtime OpenMLS messaging yet. It defines the contract that will guide implementation.

The chosen dev command names are:

    openmls-send-dev
    openmls-inbox-dev

Existing stub-era commands remain for now:

    send
    inbox
    ack

Correct interpretation:

    openmls-send-dev / openmls-inbox-dev are future explicit dev-only OpenMLS runtime commands.
    send / inbox must not be silently replaced yet.
    ack must remain consume-success gated for OpenMLS runtime inbox.
    local-backbone remains reserved until whole-path runtime validation exists.
    v0.4.2 is documentation/contract work, not production or runtime proof.

Hard nonclaims remain: v0.4.2 does **not** prove production readiness, production E2EE, hostile-server safety, metadata privacy, Android readiness, CarbonStackOS readiness, secure vault/storage, public ingress safety, real homelab validation, audit, certification, PQ/hybrid security, or mature messenger UX.

---

## 2. Current Public Release State

Public release:

    https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.4.0

Release title:

    CarbonStack v0.4.0 Broad Local Deployability Pre-Release

Release tag:

    v0.4.0

Release commit:

    24ac9fc494 / 24ac9fc docs: record v0.4.0 package rehearsal

Gitea release status:

    Pre-Release

Release assets attached:

    carbonstack-v0.4.0-broad-local-deployability-pre-release.tgz
    carbonstack-v0.4.0-release-manifest.json
    carbonstack-v0.4.0-package-checksums.txt
    carbonstack-v0.4.0-asset-checksums.txt
    carbonstack-v0.4.0-validation-freeze.md
    v0.4.0-testing-runbook.md
    v0.4.0-release-notes.md
    LICENSE

Important release-page warning:

    The default Gitea Source Code ZIP/TAR.GZ downloads are auto-generated carbonstack-only archives.
    The intended runnable validation artifact is the attached multi-repo release package plus companion assets.

Post-release public surface:

    carbonstack README is evergreen.
    It points interested parties to the Releases page for runnable / known-good artifacts.
    It states Gitea is the source of truth.
    GitHub mirrors are secondary push mirrors, not release authority.

---

## 3. Current Repo Heads

    carbonstack        a799764 (HEAD -> main, origin/main, origin/HEAD) docs: define runtime OpenMLS command contract
    carbonstack-comms  012c8bf (HEAD -> main, origin/main, origin/HEAD) scripts: support source snapshot artifact guard
    carbonstack-cypher 9ab994c (HEAD -> main, origin/main, origin/HEAD) docs: point to local operator runbook
    carbonstack-os     1bbbe52 (HEAD -> main, origin/main, origin/HEAD) docs: clarify CarbonStackOS target direction

Current public release tag:

    v0.4.0 -> 24ac9fc494 / 24ac9fc docs: record v0.4.0 package rehearsal

Post-release public-surface cleanup:

    6ba7aab docs: make front readme evergreen

v0.4.1 runtime recon commit:

    83d2f0f docs: record runtime OpenMLS Comms recon

Current v0.4.2 docs commit:

    a799764 docs: define runtime OpenMLS command contract

Continuity note:

    `a799764` is post-release mainline documentation/contract work. It does not retag v0.4.0. The v0.4.0 release remains attached to `24ac9fc494`.

---

## 4. Current Validated / Known State

Validated / known at the v0.4.2 checkpoint:

    [RELEASE] v0.4.0 Gitea release exists, is marked Pre-Release, and is attached to 24ac9fc494.
    [PACKAGE] v0.4.0 package includes carbonstack, carbonstack-comms, carbonstack-cypher, and release metadata.
    [PACKAGE] carbonstack-os is deliberately excluded from the runnable package.
    [RUNNER] full means release-snapshot followed by local-cypher.
    [RUNNER] release-snapshot validates package layout, metadata, strict pre-test artifact scan, package checksums, and core.
    [RUNNER] core validates doctor, OpenMLS real-Cypher lifecycle, Comms package tests, Cypher package tests, and artifact scans.
    [RUNNER] local-cypher validates Cypher-only lifecycle, restart/persistence against the same temporary DB, positive opaque OpenMLS envelope/ack flow, and historical negative stub-text/OpenMLS protocol pairing.
    [DOCS] main carbonstack README is evergreen.
    [DOCS] docs/155-runtime-comms-openmls-cypher-recon-v0.md records the v0.4.1 runtime recon.
    [DOCS] docs/156-runtime-openmls-command-contract-v0.md records the v0.4.2 command contract.
    [DOCS] docs/README.md and roadmap/ROADMAP.md now point to the v0.4.2 contract.
    [MIRRORS] all four repos have GitHub push-mirror functionality, with Gitea remaining source of truth.

v0.4.2 validation result from the final log:

    local-cypher passed.
    doctor passed.
    core --clean-generated passed.
    OpenMLS real-Cypher lifecycle test passed.
    carbonstack-comms package tests passed.
    carbonstack-cypher package tests passed.
    post-test artifact scan saw only known OpenMLS sidecar generated roots.
    --clean-generated removed .carbonstack-openmls-sidecar-state and target.
    final clean status showed no dirty files across the four repos.

Current contract result:

    openmls-send-dev is the planned first dev-only OpenMLS send command.
    openmls-inbox-dev is the planned first dev-only OpenMLS inbox/open/ack command.
    existing send/inbox/ack remain untouched for now.
    ack must remain sidecar-consume-success gated.
    v0.5.x state/trust/vault and PQ/hybrid ciphersuite migration remain future work.

Expected v0.4.0 public package validation commands remain:

    cd <package-root>/carbonstack/tools/carbonstack-validate
    go run . --profile verify-checksums --root <package-root>
    go run . --profile full --root <package-root> --clean-generated

Live umbrella validation commands remain:

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go run . --profile local-cypher
    go run . --profile doctor
    go run . --profile core --clean-generated

---

## 5. v0.4.2 Work Completed

### 5.1 Runtime OpenMLS command contract was written

New contract doc:

    carbonstack/docs/156-runtime-openmls-command-contract-v0.md

Updated docs:

    carbonstack/docs/README.md
    carbonstack/roadmap/ROADMAP.md

Mainline docs commit:

    a799764 docs: define runtime OpenMLS command contract

The contract is deliberately docs-first. It defines behavior before implementation to avoid accidentally mutating the old stub-era CLI into an overclaimed runtime path.

### 5.2 Dev command names were chosen

Chosen names:

    openmls-send-dev
    openmls-inbox-dev

Reasoning:

    explicit dev-only naming
    avoids silent replacement of send/inbox
    front-loads experimental status
    keeps CLI-only runtime proof separate from UX/product claims
    avoids Android/GUI/Relay Space assumptions

### 5.3 openmls-send-dev contract was defined

Contract purpose:

    protect plaintext through OpenMLS sidecar message-protect
    submit resulting application-message artifact through Cypher

Candidate command shape:

    comms openmls-send-dev \
      --to-device <recipient-cypher-device-id> \
      --sidecar-device-label <sender-sidecar-device-label> \
      --conversation <sidecar-conversation-label> \
      --message <plaintext> \
      [--message-label <label>] \
      [--strict]

Required behavior:

    load local Comms state
    evaluate trust policy where applicable
    call sidecar message-protect
    require sidecar success before relay submit
    obtain application-message artifact path from sidecar output
    submit artifact using internal/relay.SubmitOpenMLSArtifactEnvelope
    print envelope ID and metadata
    do not print private sidecar state
    do not claim production-send success

Failure rule:

    if sidecar message-protect fails, do not submit to Cypher.

### 5.4 openmls-inbox-dev contract was defined

Contract purpose:

    retrieve OpenMLS application-message envelopes from Cypher
    write artifacts to sidecar-compatible location
    call sidecar message-open
    ack only after sidecar consume succeeds

Candidate command shape:

    comms openmls-inbox-dev \
      --sidecar-device-label <recipient-sidecar-device-label> \
      --conversation <sidecar-conversation-label> \
      [--limit <n>] \
      [--ack]

Required behavior:

    load local Comms state
    fetch inbox through Cypher client
    filter/select application-message envelopes
    validate content_type and protocol_version
    write artifact using internal/relay.WriteOpenMLSArtifactFromEnvelope
    call sidecar message-open
    print plaintext only after sidecar success
    if --ack is set, ack only after message-open succeeds
    if message-open fails, leave envelope unacked and report failure

Critical rule:

    no ack before sidecar message-open success.

### 5.5 Test expectations were defined before implementation

Future v0.4.3/v0.4.4 tests should prove:

    openmls-send-dev rejects missing required flags
    openmls-send-dev does not submit if sidecar protect fails
    openmls-send-dev submits application-message artifact through Cypher on sidecar success
    openmls-inbox-dev does not ack before message-open success
    openmls-inbox-dev writes and opens application-message artifact
    openmls-inbox-dev acks only after successful message-open when --ack is set

### 5.6 Validation passed after docs commit

Final validation commands run:

    go run . --profile local-cypher
    go run . --profile doctor
    go run . --profile core --clean-generated

Observed result:

    all validation passed.
    OpenMLS real-Cypher lifecycle passed.
    carbonstack-comms package tests passed.
    carbonstack-cypher package tests passed.
    cleanup removed known generated roots.
    final repo status was clean.

---

## 6. Critical Blunders / Lessons to Carry Forward

### 6.1 Contract before implementation prevents overclaim drift

The biggest risk after v0.4.1 was jumping directly from “seam identified” into code and accidentally making `send`/`inbox` appear production-like.

Correct path:

    define dev-only command contract first
    implement narrow send command next
    implement inbox/open/ack later
    only then consider broader runtime proof

### 6.2 Do not silently replace send/inbox

Existing `send` and `inbox` are stub-era. Replacing them immediately would blur history and might break older tests or docs.

Correct path:

    add openmls-send-dev and openmls-inbox-dev
    prove them
    decide later whether send/inbox should alias, migrate, or be removed

### 6.3 Ack policy remains sacred

For OpenMLS runtime inbox:

    retrieval is not enough
    artifact write is not enough
    sidecar message-open success is required
    only then may ack happen

This preserves the already-established consume-then-ack discipline.

### 6.4 local-backbone remains premature

v0.4.2 defines commands, but no runtime path has been implemented yet.

Therefore:

    do not add local-backbone
    do not claim local-backbone
    do not claim full/local-cypher validate runtime Comms UX

### 6.5 v0.5.x concerns must not invade v0.4.3 implementation

The first dev-only send command may use dev-local state and explicit labels.

Do not solve all of this in v0.4.3:

    secure vault
    production-safe provider storage
    backup/export/recovery
    PQ/hybrid ciphersuite migration
    trust-state UX polish
    hostile-server validation

Those are later roadmap streams.

---

## 7. Current Blockers / Not Validated

Current blockers and nonvalidated areas after v0.4.2:

    Runtime OpenMLS commands are not implemented yet.
    openmls-send-dev does not exist yet.
    openmls-inbox-dev does not exist yet.
    Runtime Comms send/inbox remain stub-era.
    Runtime Comms does not yet call sidecar message-protect/message-open through user-facing commands.
    Runtime Comms does not yet use internal/relay helpers from user-facing commands.
    No runtime smoke proof exists yet for dev-only OpenMLS commands.
    local-backbone runner profile does not exist and remains reserved.
    Post-release public download verification is still useful if not already performed.
    Sanitized PRIME placement in carbonstack/logdoc-loop-system still needs a policy decision.
    Production-safe provider storage and secure vault design remain v0.5.x future work.
    PQ/hybrid ciphersuite migration readiness remains v0.5.x future work.
    Hostile-server rollback/replay/metadata-abuse validation remains future work.
    Android/CarbonStackOS implementation remains later work.
    CarbonStack Relay Space mechanics remain a future concept, not an implementation.
    Public ingress, systemd, cloudflared, real homelab deployment, and production service paths remain out of scope.

Hard nonclaims remain:

    Do not claim production readiness.
    Do not claim production E2EE.
    Do not claim hostile-server safety.
    Do not claim metadata privacy.
    Do not claim Debian production deployability.
    Do not claim systemd readiness.
    Do not claim cloudflared readiness.
    Do not claim Android readiness.
    Do not claim CarbonStackOS readiness.
    Do not claim secure vault/storage.
    Do not claim real homelab validation.
    Do not claim external audit or certification.
    Do not claim CarbonStack Relay Space mechanics are implemented.
    Do not claim local-backbone exists or is justified yet.
    Do not claim local-cypher validates runtime Comms UX.
    Do not claim full validates local-backbone.
    Do not claim full is a deployment command.
    Do not claim send/inbox are OpenMLS-backed yet.
    Do not claim openmls-send-dev/openmls-inbox-dev are implemented yet.
    Do not claim PQ/hybrid security or quantum-safe messaging.
    Do not let GitHub mirrors supersede Gitea source-of-truth status.

Allowed claim:

    CarbonStack v0.4.2 mainline has documented the runtime OpenMLS command contract.
    The planned explicit dev-only commands are openmls-send-dev and openmls-inbox-dev.
    v0.4.2 preserves the decision not to silently replace send/inbox yet.
    v0.4.2 preserves consume-success-gated ack policy.
    v0.4.0 remains the current public broad local deployability pre-release.
    Gitea remains source of truth; GitHub mirrors are secondary push mirrors.

---

## 8. Next Safest Actions

Recommended next rung:

    v0.4.3 first dev-only OpenMLS send command

Focus:

    implement openmls-send-dev in carbonstack-comms
    keep it explicit and dev-only
    call sidecar message-protect
    submit application-message artifact through internal/relay.SubmitOpenMLSArtifactEnvelope
    print envelope ID and metadata
    add minimal tests
    update carbonstack-comms README/docs as needed
    avoid touching existing send/inbox unless necessary for shared helper extraction

Expected implementation repo:

    carbonstack-comms

Expected docs/release authority repo updates if needed:

    carbonstack

Avoid next:

    implementing inbox/open/ack in the same rung unless send is trivially complete and validated
    replacing send/inbox
    adding local-backbone
    public ingress
    cloudflared
    systemd
    real homelab deployment
    Android app
    CarbonStackOS implementation
    CarbonStack Relay Space implementation
    PQ/hybrid ciphersuite implementation
    production/security claims
    broad negative-path suite unless the runtime seam exposes a direct blocker

---

## 9. Critical Paths / Functions

### Source-of-truth public surfaces

    [GITEA SOURCE OF TRUTH] https://git.bitcrusher32.win/bitcrusher32/carbonstack
    [RELEASES] https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases
    [CURRENT PUBLIC RELEASE] https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.4.0

GitHub mirrors:

    secondary push mirrors for discoverability/redundancy only.
    Do not treat mirrors as release authority unless project policy changes.

### WSL Debian active working umbrella

    [WSL-DEBIAN] umbrella: ~/repos/carbonstack_umbrella
    [WSL-DEBIAN] carbonstack: ~/repos/carbonstack_umbrella/carbonstack
    [WSL-DEBIAN] carbonstack-comms: ~/repos/carbonstack_umbrella/carbonstack-comms
    [WSL-DEBIAN] carbonstack-cypher: ~/repos/carbonstack_umbrella/carbonstack-cypher
    [WSL-DEBIAN] carbonstack-os: ~/repos/carbonstack_umbrella/carbonstack-os
    [GO-RUNNER] runner dir: ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate

### Runtime integration doc surfaces

    [DOC] v0.4.1 runtime recon:
    carbonstack/docs/155-runtime-comms-openmls-cypher-recon-v0.md

    [DOC] v0.4.2 command contract:
    carbonstack/docs/156-runtime-openmls-command-contract-v0.md

    [DOC] docs index:
    carbonstack/docs/README.md

    [DOC] roadmap:
    carbonstack/roadmap/ROADMAP.md

### Comms runtime surfaces

    [COMMS] CLI entry:
    carbonstack-comms/cmd/comms/main.go

    [COMMS] command implementation:
    carbonstack-comms/internal/app/commands.go

    [COMMS] Cypher client:
    carbonstack-comms/internal/client/cypher.go

    [COMMS] relay helpers:
    carbonstack-comms/internal/relay/cypher_bridge.go
    carbonstack-comms/internal/relay/openmls_artifacts.go

    [COMMS] OpenMLS sidecar dir:
    carbonstack-comms/internal/protocol/mls/openmls-sidecar

    [COMMS] sidecar README:
    carbonstack-comms/internal/protocol/mls/openmls-sidecar/README.md

### Current stub-era Comms functions / behaviors

    send:
      trust policy evaluation
      mock/stub provider encryption
      Cypher envelope submit

    inbox:
      Cypher inbox retrieval
      mock/stub provider plaintext display as stub_plaintext

    ack:
      envelope ack by envelope ID and recipient device ID

### Planned v0.4.x dev runtime commands

    openmls-send-dev:
      sidecar message-protect
      internal/relay.SubmitOpenMLSArtifactEnvelope
      print envelope metadata

    openmls-inbox-dev:
      Cypher inbox
      internal/relay.WriteOpenMLSArtifactFromEnvelope
      sidecar message-open
      optional ack only after successful message-open

### OpenMLS sidecar commands relevant to v0.4.x

    provider-info
    identity-create
    identity-status
    public-bundle-export
    conversation-create
    conversation-load-check
    conversation-add-member
    conversation-join
    message-protect
    message-open

### Relay/API functions relevant to runtime integration

    internal/client.CypherClient.SubmitEnvelope
    internal/client.CypherClient.Inbox
    internal/client.CypherClient.AckEnvelope
    internal/relay.SubmitOpenMLSArtifactEnvelope
    internal/relay.WriteOpenMLSArtifactFromEnvelope

### Existing v0.4.0 package validation

    cd <package-root>/carbonstack/tools/carbonstack-validate
    go run . --profile verify-checksums --root <package-root>
    go run . --profile full --root <package-root> --clean-generated

### Live umbrella validation

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go run . --profile local-cypher
    go run . --profile doctor
    go run . --profile core --clean-generated

### Envelope/protocol constants

Accepted opaque OpenMLS application-message pair:

    content_type=carbonstack.mls.application-message.v0
    protocol_version=carbonstack-openmls-sidecar-v0

Historical invalid pair validated by local-cypher:

    content_type=carbonstack.message.text.stub.v0
    protocol_version=carbonstack-openmls-sidecar-v0

Expected result:

    HTTP 400
    error code unsupported_protocol_version

---

## 10. v0.4.x Forward Plan

### v0.4.3 — first dev-only OpenMLS send command

Likely implementation target:

    openmls-send-dev

Expected behavior:

    call sidecar message-protect
    submit OpenMLS application-message artifact through relay helper
    print envelope ID and metadata
    no mature UX claims

### v0.4.4 — first dev-only OpenMLS inbox/open/ack command

Implementation target:

    openmls-inbox-dev

Expected behavior:

    query Cypher inbox
    write OpenMLS artifact from envelope
    call sidecar message-open
    print plaintext only after sidecar success
    ack only after sidecar success

### v0.4.5 — runtime smoke proof

Expected:

    local dev runtime proof that uses Comms CLI commands against Cypher/OpenMLS sidecar
    still dev/pre-alpha
    not local-backbone unless whole-path behavior is meaningful enough

### v0.4.6+ — local-backbone decision point

Only consider `local-backbone` when:

    runtime Comms send/inbox path is actually OpenMLS/Cypher-backed
    state expectations are documented
    validation can run the whole path repeatably
    name does not overclaim production/deployment maturity

---

## 11. v0.5.x and Later Carry-Forward

### v0.5.x

Local storage, trust state, provider state, vault design, and PQ/hybrid ciphersuite migration readiness.

Do not start PQ/hybrid implementation before the runtime OpenMLS/Cypher path exists. PQ/hybrid migration touches provider state, ciphersuite metadata, trust display, artifact size, reinit/rekey/recovery policy, and claim boundaries.

### v0.6.x+

Hostile-server and abuse-resistance harnesses.

### v0.7.x+

Deployability and operations hardening.

### v0.8.x+

Documented self-pentest / adversarial validation.

### v0.9.x+

Claim-boundary review.

### v0.10.x+

Android backend/app work.

### v1.x.x

Public app/server major epoch and much later CarbonStackOS/appliance research, if justified.

---

## 12. Lean Breakpoint Summary

    CarbonStack v0.4.2 is the runtime OpenMLS command contract checkpoint. v0.4.0 remains the current public broad local deployability pre-release. carbonstack main is now at a799764 docs: define runtime OpenMLS command contract. v0.4.2 locks in the dev-only command names openmls-send-dev and openmls-inbox-dev, while preserving the decision not to silently replace existing send/inbox yet. The contract says openmls-send-dev should call sidecar message-protect and submit application-message artifacts through internal/relay.SubmitOpenMLSArtifactEnvelope; openmls-inbox-dev should retrieve Cypher envelopes, write artifacts, call sidecar message-open, and ack only after sidecar consume succeeds. Runtime OpenMLS commands are not implemented yet. The next safe rung is v0.4.3: implement the first dev-only OpenMLS send command in carbonstack-comms, with minimal tests and no production/security/local-backbone claims.


---

## 13. Preserved Immediate Previous Handoff: v0.4.1

The following is the previous v0.4.1 handoff. Where it conflicts with the v0.4.2 overlay above, v0.4.2 wins for current state.


# CarbonStack LogDoc v0.4.1

**Last updated:** 2026-06-03 local session  
**Owner/operator:** bitcrusher32  
**Assistant/model:** GPT-5.5 Thinking  
**Current phase:** Phase 2F / **v0.4.x runtime Comms/OpenMLS integration recon has begun**. The v0.4.0 public broad local deployability pre-release is complete, the main `carbonstack` README is evergreen, and the first post-release runtime recon has landed. `carbonstack` is now at `83d2f0f docs: record runtime OpenMLS Comms recon`. `carbonstack-comms` remains at `012c8bf scripts: support source snapshot artifact guard`; `carbonstack-cypher` remains at `9ab994c docs: point to local operator runbook`; `carbonstack-os` remains at `1bbbe52 docs: clarify CarbonStackOS target direction`.

**Version schema:** `v[scope].[timeline]`. This file is `v0.4.1`, the first post-v0.4.0 runtime-integration checkpoint. Per LogDoc V2 practice, this Markdown file preserves current operational continuity and the key recon findings; the JSON breakpoint remains the lean machine-readable handoff state.

**Compression note:** The detailed v0.3.0 through v0.3.36PRIME process ledger remains archived in PRIME/sanitized PRIME. The prior v0.4.0 handoff compressed that history and remains preserved below only as immediate predecessor context. This v0.4.1 file should now carry the active v0.4.x runtime-integration thread forward.

---

## 1. Project Goal

**Active goal:** Move CarbonStack from v0.4.0 release-package/local-deployability validation into v0.4.x runtime Comms/OpenMLS integration. The next safe work is to define a runtime OpenMLS command contract before implementation.

CarbonStack’s broader project intent remains:

- **CarbonStack** is the secure-communications appliance-stack project and shared doctrine/release authority.
- **CarbonStackComms** is the text-first messaging client and current OpenMLS sidecar/relay integration surface.
- **CarbonStackCypher** is the self-hostable relay/storage stack.
- **CarbonStackOS** is future constrained Android-derived appliance OS work and remains out of current runnable validation packages.
- The `carbonstack` Gitea repo remains the source of truth for release state, release assets, docs, and public claims.
- GitHub mirrors exist as secondary push mirrors only; Gitea remains authoritative.

**North star:**

> Its goal is not to be a better smartphone. Its goal is to stop being a smartphone.

**Core doctrine:**

> Every feature is guilty until it proves it does not add unacceptable parser, network, sensor, identity, or filesystem authority.

**Current reality after v0.4.1:** v0.4.0 proved the package/release validation ladder. v0.4.1 proves the project has now correctly pivoted into runtime integration recon rather than more packaging work.

The most important v0.4.1 finding:

    CarbonStackComms already has user-facing CLI commands named send, inbox, and ack.
    However, send/inbox are still stub-era runtime behavior.
    The real OpenMLS/Cypher relay path exists and is validated through lower-level protocol/relay tests, not through the current user-facing runtime CLI.

Correct next direction:

    do not silently replace send/inbox yet.
    add explicit dev-only OpenMLS runtime command contracts first.
    keep local-backbone reserved until whole-path runtime validation exists.

Candidate dev command names carried forward:

    openmls-send-dev
    openmls-inbox-dev

Alternative names:

    send-openmls-dev
    inbox-openmls-dev

Hard nonclaim remains: v0.4.1 does **not** prove production readiness, production E2EE, hostile-server safety, metadata privacy, Android readiness, CarbonStackOS readiness, secure vault/storage, public ingress safety, real homelab validation, audit, certification, or mature messenger UX.

---

## 2. Current Public Release State

Public release:

    https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.4.0

Release title:

    CarbonStack v0.4.0 Broad Local Deployability Pre-Release

Release tag:

    v0.4.0

Release commit:

    24ac9fc494 / 24ac9fc docs: record v0.4.0 package rehearsal

Gitea release status:

    Pre-Release

Release assets attached:

    carbonstack-v0.4.0-broad-local-deployability-pre-release.tgz
    carbonstack-v0.4.0-release-manifest.json
    carbonstack-v0.4.0-package-checksums.txt
    carbonstack-v0.4.0-asset-checksums.txt
    carbonstack-v0.4.0-validation-freeze.md
    v0.4.0-testing-runbook.md
    v0.4.0-release-notes.md
    LICENSE

Important release-page warning:

    The default Gitea Source Code ZIP/TAR.GZ downloads are auto-generated carbonstack-only archives.
    The intended runnable validation artifact is the attached multi-repo release package plus companion assets.

Post-release public surface:

    carbonstack README is evergreen.
    It points interested parties to the Releases page for runnable / known-good artifacts.
    It states Gitea is the source of truth.
    GitHub mirrors are secondary push mirrors, not release authority.

---

## 3. Current Repo Heads

    carbonstack        83d2f0f (HEAD -> main, origin/main, origin/HEAD) docs: record runtime OpenMLS Comms recon
    carbonstack-comms  012c8bf (HEAD -> main, origin/main, origin/HEAD) scripts: support source snapshot artifact guard
    carbonstack-cypher 9ab994c (HEAD -> main, origin/main, origin/HEAD) docs: point to local operator runbook
    carbonstack-os     1bbbe52 (HEAD -> main, origin/main, origin/HEAD) docs: clarify CarbonStackOS target direction

Current public release tag:

    v0.4.0 -> 24ac9fc494 / 24ac9fc docs: record v0.4.0 package rehearsal

Post-release public-surface cleanup:

    6ba7aab docs: make front readme evergreen

Current v0.4.1 docs commit:

    83d2f0f docs: record runtime OpenMLS Comms recon

Continuity note:

    `83d2f0f` is post-release mainline documentation/recon work. It does not retag v0.4.0. The v0.4.0 release remains attached to `24ac9fc494`.

---

## 4. Current Validated / Known State

Validated / known at the v0.4.1 checkpoint:

    [RELEASE] v0.4.0 Gitea release exists, is marked Pre-Release, and is attached to 24ac9fc494.
    [PACKAGE] v0.4.0 package includes carbonstack, carbonstack-comms, carbonstack-cypher, and release metadata.
    [PACKAGE] carbonstack-os is deliberately excluded from the runnable package.
    [RUNNER] full means release-snapshot followed by local-cypher.
    [RUNNER] release-snapshot validates package layout, metadata, strict pre-test artifact scan, package checksums, and core.
    [RUNNER] core validates doctor, OpenMLS real-Cypher lifecycle, Comms package tests, Cypher package tests, and artifact scans.
    [RUNNER] local-cypher validates Cypher-only lifecycle, restart/persistence against the same temporary DB, positive opaque OpenMLS envelope/ack flow, and historical negative stub-text/OpenMLS protocol pairing.
    [DOCS] main carbonstack README is evergreen.
    [DOCS] docs/155-runtime-comms-openmls-cypher-recon-v0.md exists and records the first v0.4.x runtime integration recon.
    [MIRRORS] all four repos have GitHub push-mirror functionality, with Gitea remaining source of truth.

v0.4.1 recon result:

    Comms runtime CLI exists but remains stub-era for send/inbox.
    Lower-level OpenMLS/Cypher relay seams already exist and are validated.
    The gap is wiring the runtime CLI to the OpenMLS sidecar and relay helpers.

Observed Comms CLI command surface from recon:

    dev-create-invite
    claim-invite
    register-device
    list-devices
    verify-device
    trust-events
    mark-device-changed
    revoke-device
    send
    inbox
    ack

Current stub-era runtime behavior:

    send evaluates trust policy, uses mock/stub provider encryption, and submits a Cypher envelope.
    inbox retrieves queued Cypher envelopes and prints stub_plaintext from mock/stub provider decryption.
    ack acknowledges by envelope ID and recipient device, but current runtime inbox does not consume through OpenMLS first.

Existing lower-level Comms/Cypher pieces:

    internal/client/cypher.go exposes create/claim invite, register/list devices, submit envelope, inbox, and ack.
    internal/relay.SubmitOpenMLSArtifactEnvelope exists.
    internal/relay.WriteOpenMLSArtifactFromEnvelope exists.
    OpenMLS sidecar commands include provider-info, identity-create, identity-status, public-bundle-export, conversation-create, conversation-load-check, conversation-add-member, conversation-join, message-protect, and message-open.
    Tests already prove OpenMLS KeyPackage, Welcome, and application-message artifacts can be relayed through real local Cypher and consumed by the sidecar.

Expected v0.4.0 public package validation commands remain:

    cd <package-root>/carbonstack/tools/carbonstack-validate
    go run . --profile verify-checksums --root <package-root>
    go run . --profile full --root <package-root> --clean-generated

Live umbrella validation commands remain:

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go run . --profile local-cypher
    go run . --profile doctor
    go run . --profile core --clean-generated

---

## 5. v0.4.1 Work Completed

### 5.1 Runtime Comms/OpenMLS/Cypher recon was performed

Recon inspected all three active repos, with `carbonstack-comms` as the primary target:

    carbonstack
    carbonstack-comms
    carbonstack-cypher

`carbonstack-os` was included only in final status checks, not treated as part of runtime wiring.

The recon collected:

    repo heads/status
    top-level tree structure
    Comms CLI command surface
    Comms README/docs references to OpenMLS/Cypher/send/inbox
    Go package/test lists
    command entrypoints
    OpenMLS sidecar files/Cargo surface
    sidecar command/API references
    Cypher API/docs/handler surface
    runner profile and local-cypher/core references
    cross-repo content_type/protocol_version names
    live validation baseline
    final clean status

### 5.2 Main seam was identified

The main seam is clear:

    runtime Comms CLI send/inbox remain old stub-era behavior
    OpenMLS/Cypher relay is already validated beneath the runtime CLI
    v0.4.x should join them carefully rather than replacing send/inbox in one jump

The first implementation should not be broad UX, Android, OS, public deployment, or local-backbone.

### 5.3 v0.4.1 docs were added

New/updated docs:

    carbonstack/docs/155-runtime-comms-openmls-cypher-recon-v0.md
    carbonstack/docs/README.md
    carbonstack/roadmap/ROADMAP.md

Mainline docs commit:

    83d2f0f docs: record runtime OpenMLS Comms recon

The doc records:

    send/inbox are currently stub-era
    the OpenMLS/Cypher relay seam exists below runtime CLI
    dev-only OpenMLS runtime commands should be introduced before replacing send/inbox
    ack must remain consume-success gated
    local-backbone remains reserved
    v0.4.2 should be a runtime OpenMLS command contract

### 5.4 Final status after v0.4.1 rung

Final repo snapshot:

    carbonstack        83d2f0f docs: record runtime OpenMLS Comms recon
    carbonstack-comms  012c8bf scripts: support source snapshot artifact guard
    carbonstack-cypher 9ab994c docs: point to local operator runbook
    carbonstack-os     1bbbe52 docs: clarify CarbonStackOS target direction

Final status showed no dirty files in the four repos.

---

## 6. Critical Blunders / Lessons to Carry Forward

### 6.1 Runtime CLI existence is not runtime OpenMLS integration

The presence of `send`, `inbox`, and `ack` commands does not mean runtime OpenMLS messaging exists.

Current reality:

    send/inbox exist
    send/inbox are stub-era
    OpenMLS/Cypher is validated in lower-level tests

Lesson:

    always distinguish command surface from validated cryptographic/runtime behavior.

### 6.2 Do not replace send/inbox too early

Directly replacing `send` and `inbox` risks breaking older stub-era behavior and creating a false impression of mature messaging UX.

Correct path:

    add explicit dev-only OpenMLS runtime commands first
    prove behavior
    document warnings
    decide later whether send/inbox become aliases or replacements

### 6.3 Ack must remain sidecar-consume gated

Current doctrine:

    ack after successful sidecar consume, not merely after envelope retrieval.

For OpenMLS runtime inbox:

    message-open must succeed before ack.
    failed sidecar consume should leave envelope queued or explicitly reported as unacked.

### 6.4 local-backbone is still premature

Even after v0.4.1, the project should not add or claim `local-backbone`.

Reason:

    runtime Comms CLI still does not use OpenMLS/Cypher.
    local-cypher is Cypher-only.
    full is a package validation ladder.
    whole-path runtime behavior is not implemented yet.

### 6.5 v0.4.x should stay CLI/dev-first

The right path is CLI-only dev/runtime proof first.

Do not jump to:

    Android
    GUI
    CarbonStackOS
    public ingress
    systemd/cloudflared
    real homelab deployability
    polished user UX

### 6.6 PQ/hybrid ciphersuite migration remains later

The refreshed roadmap now includes post-quantum / hybrid ciphersuite migration readiness in v0.5.x, not as a v0.4.1 blocker.

Reason:

    v0.4.x first needs runtime OpenMLS/Cypher send/inbox wiring.
    PQ/hybrid migration touches ciphersuites, provider state, trust state, artifact size, protocol metadata, migration/reinit/rekey policy, and claim boundaries.

---

## 7. Current Blockers / Not Validated

Current blockers and nonvalidated areas after v0.4.1:

    Runtime Comms send/inbox remain stub-era.
    No dev-only OpenMLS runtime commands exist yet.
    No runtime OpenMLS command contract has been written yet.
    No CLI flag/state/output schema exists for OpenMLS runtime send/inbox.
    Runtime Comms does not yet call sidecar message-protect/message-open through user-facing commands.
    Runtime Comms does not yet use internal/relay helpers from user-facing commands.
    local-backbone runner profile does not exist and remains reserved.
    Post-release public download verification is still useful if not already performed.
    Sanitized PRIME placement in carbonstack/logdoc-loop-system still needs a policy decision.
    Production-safe provider storage and secure vault design remain v0.5.x future work.
    PQ/hybrid ciphersuite migration readiness remains v0.5.x future work.
    Hostile-server rollback/replay/metadata-abuse validation remains future work.
    Android/CarbonStackOS implementation remains later work.
    CarbonStack Relay Space mechanics remain a future concept, not an implementation.
    Public ingress, systemd, cloudflared, real homelab deployment, and production service paths remain out of scope.

Hard nonclaims remain:

    Do not claim production readiness.
    Do not claim production E2EE.
    Do not claim hostile-server safety.
    Do not claim metadata privacy.
    Do not claim Debian production deployability.
    Do not claim systemd readiness.
    Do not claim cloudflared readiness.
    Do not claim Android readiness.
    Do not claim CarbonStackOS readiness.
    Do not claim secure vault/storage.
    Do not claim real homelab validation.
    Do not claim external audit or certification.
    Do not claim CarbonStack Relay Space mechanics are implemented.
    Do not claim local-backbone exists or is justified yet.
    Do not claim local-cypher validates runtime Comms UX.
    Do not claim full validates local-backbone.
    Do not claim full is a deployment command.
    Do not claim send/inbox are OpenMLS-backed yet.
    Do not claim PQ/hybrid security or quantum-safe messaging.
    Do not let GitHub mirrors supersede Gitea source-of-truth status.

Allowed claim:

    CarbonStack v0.4.1 mainline has documented the first runtime Comms/OpenMLS/Cypher recon.
    v0.4.1 found that send/inbox are still stub-era while OpenMLS/Cypher relay is already validated below the runtime CLI.
    The next recommended rung is a runtime OpenMLS command contract.
    v0.4.0 remains the current public broad local deployability pre-release.
    Gitea remains source of truth; GitHub mirrors are secondary push mirrors.

---

## 8. Next Safest Actions

Recommended next rung:

    v0.4.2 runtime OpenMLS command contract

Focus:

    choose dev command names
    define CLI flags
    define required local state
    define sidecar invocation boundaries
    define output schema and warnings
    define ack-on-success behavior
    define tests before implementation
    define what remains stub-era
    define what must wait for v0.5.x provider/trust/vault work

Likely contract candidates:

    openmls-send-dev
    openmls-inbox-dev

or:

    send-openmls-dev
    inbox-openmls-dev

Recommended v0.4.2 shape:

    docs/contract first
    no large code changes
    implementation only after command/state/output boundaries are explicit

Avoid next:

    replacing send/inbox immediately
    adding local-backbone
    public ingress
    cloudflared
    systemd
    real homelab deployment
    Android app
    CarbonStackOS implementation
    CarbonStack Relay Space implementation
    PQ/hybrid ciphersuite implementation
    production/security claims
    broad negative-path suite unless the runtime seam exposes a direct blocker

---

## 9. Critical Paths / Functions

### Source-of-truth public surfaces

    [GITEA SOURCE OF TRUTH] https://git.bitcrusher32.win/bitcrusher32/carbonstack
    [RELEASES] https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases
    [CURRENT PUBLIC RELEASE] https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.4.0

GitHub mirrors:

    secondary push mirrors for discoverability/redundancy only.
    Do not treat mirrors as release authority unless project policy changes.

### WSL Debian active working umbrella

    [WSL-DEBIAN] umbrella: ~/repos/carbonstack_umbrella
    [WSL-DEBIAN] carbonstack: ~/repos/carbonstack_umbrella/carbonstack
    [WSL-DEBIAN] carbonstack-comms: ~/repos/carbonstack_umbrella/carbonstack-comms
    [WSL-DEBIAN] carbonstack-cypher: ~/repos/carbonstack_umbrella/carbonstack-cypher
    [WSL-DEBIAN] carbonstack-os: ~/repos/carbonstack_umbrella/carbonstack-os
    [GO-RUNNER] runner dir: ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate

### Runtime integration doc surfaces

    [DOC] v0.4.1 runtime recon:
    carbonstack/docs/155-runtime-comms-openmls-cypher-recon-v0.md

    [DOC] docs index:
    carbonstack/docs/README.md

    [DOC] roadmap:
    carbonstack/roadmap/ROADMAP.md

### Comms runtime surfaces

    [COMMS] CLI entry:
    carbonstack-comms/cmd/comms/main.go

    [COMMS] command implementation:
    carbonstack-comms/internal/app/commands.go

    [COMMS] Cypher client:
    carbonstack-comms/internal/client/cypher.go

    [COMMS] relay helpers:
    carbonstack-comms/internal/relay/cypher_bridge.go
    carbonstack-comms/internal/relay/openmls_artifacts.go

    [COMMS] OpenMLS sidecar dir:
    carbonstack-comms/internal/protocol/mls/openmls-sidecar

    [COMMS] sidecar README:
    carbonstack-comms/internal/protocol/mls/openmls-sidecar/README.md

### Current stub-era Comms functions / behaviors

    send:
      trust policy evaluation
      mock/stub provider encryption
      Cypher envelope submit

    inbox:
      Cypher inbox retrieval
      mock/stub provider plaintext display as stub_plaintext

    ack:
      envelope ack by envelope ID and recipient device ID

### OpenMLS sidecar commands relevant to v0.4.x

    provider-info
    identity-create
    identity-status
    public-bundle-export
    conversation-create
    conversation-load-check
    conversation-add-member
    conversation-join
    message-protect
    message-open

### Relay/API functions relevant to runtime integration

    internal/client.CypherClient.SubmitEnvelope
    internal/client.CypherClient.Inbox
    internal/client.CypherClient.AckEnvelope
    internal/relay.SubmitOpenMLSArtifactEnvelope
    internal/relay.WriteOpenMLSArtifactFromEnvelope

### Existing v0.4.0 package validation

    cd <package-root>/carbonstack/tools/carbonstack-validate
    go run . --profile verify-checksums --root <package-root>
    go run . --profile full --root <package-root> --clean-generated

### Live umbrella validation

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go run . --profile local-cypher
    go run . --profile doctor
    go run . --profile core --clean-generated

### Envelope/protocol constants

Accepted opaque OpenMLS application-message pair:

    content_type=carbonstack.mls.application-message.v0
    protocol_version=carbonstack-openmls-sidecar-v0

Historical invalid pair validated by local-cypher:

    content_type=carbonstack.message.text.stub.v0
    protocol_version=carbonstack-openmls-sidecar-v0

Expected result:

    HTTP 400
    error code unsupported_protocol_version

---

## 10. v0.4.x Forward Plan

### v0.4.2 — runtime OpenMLS command contract

Define before implementation:

    command names
    CLI flags
    state dependencies
    sidecar command invocation boundaries
    relay helper usage
    output schema
    warnings/nonclaims
    ack-on-success behavior
    tests

### v0.4.3 — first dev-only OpenMLS send command

Likely implementation target:

    openmls-send-dev

or:

    send-openmls-dev

Expected behavior:

    call sidecar message-protect
    submit OpenMLS application-message artifact through relay helper
    print envelope ID and metadata
    no mature UX claims

### v0.4.4 — first dev-only OpenMLS inbox/open/ack command

Likely implementation target:

    openmls-inbox-dev

or:

    inbox-openmls-dev

Expected behavior:

    query Cypher inbox
    write OpenMLS artifact from envelope
    call sidecar message-open
    print plaintext only after sidecar success
    ack only after sidecar success

### v0.4.5 — runtime smoke proof

Expected:

    local dev runtime proof that uses Comms CLI commands against Cypher/OpenMLS sidecar
    still dev/pre-alpha
    not local-backbone unless whole-path behavior is meaningful enough

### v0.4.6+ — local-backbone decision point

Only consider `local-backbone` when:

    runtime Comms send/inbox path is actually OpenMLS/Cypher-backed
    state expectations are documented
    validation can run the whole path repeatably
    name does not overclaim production/deployment maturity

---

## 11. v0.5.x and Later Carry-Forward

### v0.5.x

Local storage, trust state, provider state, vault design, and PQ/hybrid ciphersuite migration readiness.

Do not start PQ/hybrid implementation before the runtime OpenMLS/Cypher path exists. PQ/hybrid migration touches provider state, ciphersuite metadata, trust display, artifact size, reinit/rekey/recovery policy, and claim boundaries.

### v0.6.x+

Hostile-server and abuse-resistance harnesses.

### v0.7.x+

Deployability and operations hardening.

### v0.8.x+

Documented self-pentest / adversarial validation.

### v0.9.x+

Claim-boundary review.

### v0.10.x+

Android backend/app work.

### v1.x.x

Public app/server major epoch and much later CarbonStackOS/appliance research, if justified.

---

## 12. Lean Breakpoint Summary

    CarbonStack v0.4.1 is the first post-v0.4.0 runtime integration recon checkpoint. v0.4.0 remains the current public broad local deployability pre-release. carbonstack main is now at 83d2f0f docs: record runtime OpenMLS Comms recon. The v0.4.1 recon found that CarbonStackComms has send/inbox/ack CLI commands, but send/inbox remain stub-era and use mock/stub provider behavior. The lower-level OpenMLS/Cypher relay path already exists and is validated through internal relay/protocol tests. The next safe rung is v0.4.2 runtime OpenMLS command contract: choose explicit dev-only command names, define flags/state/output/ack behavior, then implement carefully. Do not replace send/inbox yet, do not claim local-backbone, and do not make production/security/PQ claims.


---

## 13. Preserved Immediate Previous Handoff: v0.4.0

The following is the previous compressed v0.4.0 handoff. Where it conflicts with the v0.4.1 overlay above, v0.4.1 wins for current state.

# CarbonStack LogDoc v0.4.0

**Last updated:** 2026-06-03 local session  
**Owner/operator:** bitcrusher32  
**Assistant/model:** GPT-5.5 Thinking  
**Current phase:** Phase 2E complete / **v0.4.0 minor epoch release complete + post-release README cleanup**. The public CarbonStack `v0.4.0` release is live on the Gitea source-of-truth repo as **CarbonStack v0.4.0 Broad Local Deployability Pre-Release**, marked pre-release, attached to `24ac9fc494` / `24ac9fc docs: record v0.4.0 package rehearsal`. After release, the main `carbonstack` README was rewritten to be evergreen and now points interested parties to the Releases page for runnable / known-good artifacts rather than hardcoding stale “current release” language. `carbonstack` main is now at `6ba7aab docs: make front readme evergreen`; `carbonstack-comms` remains at `012c8bf scripts: support source snapshot artifact guard`; `carbonstack-cypher` remains at `9ab994c docs: point to local operator runbook`; `carbonstack-os` remains at `1bbbe52 docs: clarify CarbonStackOS target direction`.

**Version schema:** `v[scope].[timeline]`. This file is `v0.4.0`, the compressed post-minor-epoch handoff after the v0.4.0 public release cut, PRIME ledger generation/sanitization pass, and evergreen front-README cleanup. Per LogDoc V2 practice, this Markdown file preserves the current operational handoff and high-level continuity, while the JSON breakpoint remains the lean machine-readable state.

**Compression note:** The detailed v0.3.0 through v0.3.36PRIME process ledger is intentionally **not** preserved in full here. That continuity remains archived in the v0.4.0PRIME / sanitized PRIME LogDoc. This v0.4.0 LogDoc compresses the older timeline to top-level history, lessons, and current working state so future branches can proceed without dragging the entire pre-release sprint ledger forward.

---

## 1. Project Goal

**Active goal:** Preserve CarbonStack state after completing the v0.4.0 minor epoch release and cleaning the main public README into an evergreen front door. The next safe work is either post-upload public-asset verification, committing/publishing the sanitized PRIME ledger as appropriate, or beginning v0.4.x runtime Comms/OpenMLS integration recon.

CarbonStack’s broader project intent remains:

- **CarbonStack** is the secure-communications appliance-stack project and shared doctrine/release authority.
- **CarbonStackComms** is the text-first messaging client and current OpenMLS sidecar/relay integration surface.
- **CarbonStackCypher** is the self-hostable relay/storage stack.
- **CarbonStackOS** is the future constrained Android-derived appliance OS and is not part of the runnable v0.4.0 package.
- The `carbonstack` Gitea repo remains the source of truth for release state, release assets, docs, and public claims.
- GitHub mirrors now exist for all four repos as push mirrors, but Gitea remains authoritative.

**North star:**

> Its goal is not to be a better smartphone. Its goal is to stop being a smartphone.

**Core doctrine:**

> Every feature is guilty until it proves it does not add unacceptable parser, network, sensor, identity, or filesystem authority.

**Current release reality:** v0.4.0 is the current public broad local deployability pre-release. It validates a Debian / WSL Debian-first multi-repo package path using:

    cd <package-root>/carbonstack/tools/carbonstack-validate
    go run . --profile verify-checksums --root <package-root>
    go run . --profile full --root <package-root> --clean-generated

Where:

    full = release-snapshot + local-cypher
    release-snapshot already calls core
    full does not call core twice
    local-cypher validates the Cypher-only local lifecycle afterward
    --clean-generated removes known OpenMLS sidecar generated roots after successful validation

This release is still **not** production secure. It does **not** prove production readiness, production E2EE, hostile-server safety, metadata privacy, Android readiness, CarbonStackOS readiness, secure vault/storage, Debian production deployability, systemd readiness, cloudflared readiness, public ingress safety, real homelab validation, external audit, or certification.

---

## 2. Current Public Release State

Public release:

    https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.4.0

Release title:

    CarbonStack v0.4.0 Broad Local Deployability Pre-Release

Release tag:

    v0.4.0

Release commit:

    24ac9fc494 / 24ac9fc docs: record v0.4.0 package rehearsal

Gitea release status:

    Pre-Release

Release timestamp observed from Gitea:

    2026-06-03 09:45:20 -04:00

Release assets attached:

    carbonstack-v0.4.0-broad-local-deployability-pre-release.tgz
    carbonstack-v0.4.0-release-manifest.json
    carbonstack-v0.4.0-package-checksums.txt
    carbonstack-v0.4.0-asset-checksums.txt
    carbonstack-v0.4.0-validation-freeze.md
    v0.4.0-testing-runbook.md
    v0.4.0-release-notes.md
    LICENSE

Important release-page warning:

    The default Gitea Source Code ZIP/TAR.GZ downloads are auto-generated carbonstack-only archives.
    The intended runnable validation artifact is the attached multi-repo release package plus companion assets.

Important post-release README change:

    carbonstack README is now evergreen.
    It no longer hardcodes v0.3.20 as current or v0.4.0 as future.
    It tells interested parties to check the Releases page for runnable / known-good artifacts.
    It states that Gitea is the canonical source of truth and GitHub mirrors are secondary push mirrors.

---

## 3. Current Repo Heads

    carbonstack        6ba7aab (HEAD -> main, origin/main, origin/HEAD) docs: make front readme evergreen
    carbonstack-comms  012c8bf (HEAD -> main, origin/main, origin/HEAD) scripts: support source snapshot artifact guard
    carbonstack-cypher 9ab994c (HEAD -> main, origin/main, origin/HEAD) docs: point to local operator runbook
    carbonstack-os     1bbbe52 (HEAD -> main, origin/main, origin/HEAD) docs: clarify CarbonStackOS target direction

Current public release tag:

    v0.4.0 -> 24ac9fc494 / 24ac9fc docs: record v0.4.0 package rehearsal

Important prior public release tag:

    v0.3.20 -> 3eeb1a1e1a / 3eeb1a1 docs: clean public release surface

Important historical testbed:

    v0.3.4 -> Windows 11 / PowerShell legacy testbed

Continuity note:

    `6ba7aab` is a post-release public-surface cleanup commit after the `v0.4.0` tag. The release remains attached to `24ac9fc494`, while `main` has advanced to the evergreen README state.

---

## 4. Current Validated State

Validated / known at the v0.4.0 breakpoint:

    [RELEASE] v0.4.0 Gitea release exists, is marked Pre-Release, and is attached to 24ac9fc494.
    [RELEASE] v0.4.0 includes the intended multi-repo package and seven companion release assets plus LICENSE.
    [PACKAGE] v0.4.0 package includes carbonstack, carbonstack-comms, carbonstack-cypher, and release metadata.
    [PACKAGE] carbonstack-os is deliberately excluded from the runnable package.
    [PACKAGE VALIDATION] final package generation/fresh extraction/verify-checksums/full --clean-generated passed before upload.
    [RUNNER] full means release-snapshot followed by local-cypher.
    [RUNNER] release-snapshot validates package layout, metadata, strict pre-test artifact scan, package checksums, and core.
    [RUNNER] core validates doctor, OpenMLS real-Cypher lifecycle, Comms package tests, Cypher package tests, and artifact scans.
    [RUNNER] local-cypher validates Cypher-only lifecycle, restart/persistence against the same temporary DB, positive opaque OpenMLS envelope/ack flow, and the historical negative stub-text/OpenMLS protocol pairing.
    [CLEANUP] --clean-generated removes known OpenMLS sidecar generated roots after successful validation.
    [DOCS] sanitized PRIME LogDoc exists as the detailed v0.3.0 -> v0.4.0 release sprint ledger candidate.
    [DOCS] main carbonstack README is now evergreen at 6ba7aab.
    [MIRRORS] all four repos have GitHub push-mirror functionality, with Gitea remaining source of truth.

Observed WSL Debian toolchain baseline from the v0.4.0 release-prep cycle:

    git version 2.47.3
    go version go1.24.4 linux/amd64
    rustc 1.96.0
    cargo 1.96.0
    sqlite3 3.46.1

Expected v0.4.0 public package validation commands:

    cd <package-root>/carbonstack/tools/carbonstack-validate
    go run . --profile verify-checksums --root <package-root>
    go run . --profile full --root <package-root> --clean-generated

Known generated roots handled by cleanup:

    carbonstack-comms/internal/protocol/mls/openmls-sidecar/.carbonstack-openmls-sidecar-state
    carbonstack-comms/internal/protocol/mls/openmls-sidecar/target

---

## 5. Condensed Continuity Timeline: v0.3.0 -> v0.4.0

### v0.3.0 line: release packaging foundation

The v0.3.0 line established the early release-packaging plan and began treating CarbonStack as a multi-repo release surface rather than a single-repo toy. It set the direction for using `carbonstack` as the public release/doctrine authority while leaving implementation details in component repos.

Top-level lesson:

    Release claims must follow packaged, testable artifacts, not just repo state.

### v0.3.4 line: Windows / PowerShell legacy testbed

The v0.3.4 line became the older Windows 11 / PowerShell validation testbed. It remains historically useful but is not the mainline direction after v0.3.20.

Top-level lesson:

    Windows validation helped expose packaging and extraction issues, but Windows path/build fragility made Debian / WSL Debian a better mainline release-validation target.

### v0.3.20 line: runner-backed testing release

The v0.3.20 release published the first major runner-backed testing package using the Go validation runner, release-snapshot, package checksums, release metadata, and Debian / WSL Debian as the primary direction. It also preserved Windows 11 short-path validation as the final explicit Windows dev/test validation for that phase.

Top-level lesson:

    The release package must include its own runbook, checksums, metadata, and clear warning that Gitea Source Code archives are not the intended multi-repo validation package.

### v0.3.21 -> v0.3.28: local deployability model and Cypher operator groundwork

This span clarified the local-only deployability path, the CarbonStack Relay Space terminology, Cypher operator config/data conventions, schema migration discipline, and the local Cypher API lifecycle proof. It also preserved the decision that `local-backbone` should remain reserved until whole-stack runtime Comms validation exists.

Top-level lessons:

    CarbonStack Relay Space is the preferred future term.
    SQLite migrations need explicit tracking/idempotence.
    Local operator paths and DB conventions must be documented.
    `local-backbone` should not be claimed before runtime Comms wiring justifies it.

### v0.3.29 -> v0.3.32: local-cypher runner and first negative path

This span defined and implemented the `local-cypher` runner profile. It added positive Cypher lifecycle validation, restart/persistence against the same temp DB, explicit cleanup with `--clean-generated`, and one historically grounded negative-path test for the invalid stub-text/OpenMLS protocol pairing.

Top-level lessons:

    One real historical negative path is more valuable than a speculative broad suite before release.
    Cleanup must remain explicit and narrow.
    Cypher-only validation is not runtime Comms UX and must be labeled accordingly.

### v0.3.33 -> v0.3.35: public-surface cleanup and full profile semantics

This span cleaned the public README/docs/roadmap surface for the pre-v0.4.0 runway and changed `full` from a stale `core` alias into the release-package validation ladder:

    full = release-snapshot + local-cypher

Top-level lessons:

    Public docs rot quickly when they name “current” versions.
    `release-snapshot` already calls core, so `full` must not duplicate core.
    `full` is validation, not deployment.

### v0.3.36 / v0.4.0PRIME: package rehearsal, release cut, PRIME ledger

The v0.3.36 package rehearsal proved a throwaway v0.4.0-style package root could be staged, checksum-covered, archived, fresh-extracted, checksum-verified, and validated through `full --clean-generated`. v0.4.0PRIME then preserved the release cut, asset generation/upload mechanics, and future v0.4.x plan as a full dev ledger/case-study candidate.

Top-level lessons:

    Live-umbrella `full` failure is expected without release metadata; package-root discipline matters.
    Build under `/tmp`, validate from fresh extraction, and upload only the final asset folder.
    The release-level `LICENSE` is worth keeping for multi-repo asset clarity.
    `carbonstack-os` should remain related but excluded from runnable packages until OS work is real.
    PRIME logs should preserve full continuity first, then be sanitized before public upload.

### v0.4.0 post-release README cleanup

After the release, the main `carbonstack` README was rewritten to be evergreen. It now points to Releases for known-good artifacts and avoids hardcoded “current release” wording. It also records that Gitea is source of truth while GitHub mirrors are secondary push mirrors.

Top-level lesson:

    Front READMEs should be stable orientation documents; release-specific truth belongs in release pages and runbooks.

---

## 6. Critical Blunders / Lessons to Carry Forward

1. **Do not let public docs hardcode stale release status.**  
   Use the Releases page as the known-good artifact surface.

2. **Do not confuse package validation with deployment.**  
   `full` is a release-package validation ladder, not a service runner, deploy command, or local-backbone proof.

3. **Do not run package-root profiles casually against live umbrellas.**  
   `release-snapshot` and `full` expect release metadata/checksums and are best used against fresh extracted or throwaway staged package roots.

4. **Do not upload default Gitea Source Code archives as intended release assets.**  
   They are carbonstack-only autogenerated archives, not the multi-repo validation package.

5. **Do not package CarbonStackOS until it becomes a runnable artifact.**  
   It remains a future OS vision, not a v0.4.0 release component.

6. **Keep cleanup explicit.**  
   `--clean-generated` should remain opt-in and narrowly scoped to known OpenMLS generated roots.

7. **Preserve LogDoc continuity, then compress after epoch closure.**  
   PRIME captured the detailed v0.3.x sprint. This v0.4.0 handoff intentionally compresses it.

8. **Sanitize before public case-study upload.**  
   The sanitized PRIME LogDoc exists for public/dev-ledger use; raw PRIME should not be uploaded.

---

## 7. Current Blockers / Not Validated

Current blockers and nonvalidated areas after v0.4.0 release and README cleanup:

    Post-release validation from a fresh independent public download is still useful if not already performed.
    Sanitized PRIME Markdown/JSON placement in carbonstack and/or logdoc-loop-system still needs to be decided.
    v0.4.x runtime Comms OpenMLS integration has not begun.
    Runtime Comms send/inbox remains stub-era and not wired as user-facing OpenMLS messenger UX.
    local-cypher has one negative path but not a broad adversarial suite.
    local-backbone runner profile does not exist and remains reserved.
    No helper command exists yet.
    Production-safe provider storage and secure vault design are unsolved.
    Hostile-server rollback/replay/metadata-abuse validation remains future work.
    Android/CarbonStackOS implementation remains later work.
    CarbonStack Relay Space mechanics remain a future concept, not an implementation.
    Public ingress, systemd, cloudflared, real homelab deployment, and production service paths remain out of scope.

Hard nonclaims remain:

    Do not claim production readiness.
    Do not claim production E2EE.
    Do not claim hostile-server safety.
    Do not claim metadata privacy.
    Do not claim Debian production deployability.
    Do not claim systemd readiness.
    Do not claim cloudflared readiness.
    Do not claim Android readiness.
    Do not claim CarbonStackOS readiness.
    Do not claim secure vault/storage.
    Do not claim real homelab validation.
    Do not claim external audit or certification.
    Do not claim CarbonStack Relay Space mechanics are implemented.
    Do not claim local-backbone exists or is justified yet.
    Do not claim local-cypher validates runtime Comms UX.
    Do not claim full validates local-backbone.
    Do not claim full is a deployment command.
    Do not claim --clean-generated is a general cleanup tool.
    Do not let GitHub mirrors supersede Gitea source-of-truth status.

Allowed claim:

    CarbonStack v0.4.0 is a public broad local deployability pre-release.
    v0.4.0 publishes a multi-repo validation package containing carbonstack, carbonstack-comms, carbonstack-cypher, and release metadata.
    v0.4.0 validates the intended package path with verify-checksums and full.
    full means release-snapshot followed by local-cypher.
    release-snapshot already calls core.
    v0.4.0 is Debian / WSL Debian-first.
    v0.4.0 has no Windows validation as part of this release surface.
    v0.4.0 remains pre-alpha/experimental and not production secure.
    carbonstack main has an evergreen public README that points testers to Releases for known-good artifacts.
    Gitea remains source of truth; GitHub mirrors are secondary push mirrors.

---

## 8. Next Safest Actions

Recommended immediate post-release sequence:

    1. Optionally verify v0.4.0 from a fresh public download:
       - download release package/assets from Gitea
       - verify asset checksums
       - extract package fresh
       - run verify-checksums
       - run full --clean-generated
       - record result

    2. Decide where the sanitized PRIME dev ledger belongs:
       - carbonstack docs/archive/case-study surface
       - logdoc-loop-system case-study/examples surface
       - both, with pointers

    3. Begin v0.4.x runtime integration recon:
       - inspect current Comms send/inbox command path
       - find remaining stub-era messaging
       - identify how Comms should deliberately use the OpenMLS sidecar/backbone path
       - decide what state must remain dev-only until v0.5.x storage/trust/vault work

Avoid next:

    local-backbone runner profile before runtime Comms wiring exists
    helper implementation unless it clearly reduces repeated operator mistakes
    public ingress
    cloudflared
    systemd
    real homelab deployment
    Android app
    CarbonStackOS implementation
    CarbonStack Relay Space implementation
    production claims
    broad negative-path suite unless v0.4.x integration exposes a direct blocker

---

## 9. Critical Paths / Functions

### Source-of-truth public surfaces

    [GITEA SOURCE OF TRUTH] https://git.bitcrusher32.win/bitcrusher32/carbonstack
    [RELEASES] https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases
    [CURRENT PUBLIC RELEASE] https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.4.0

GitHub mirrors:

    secondary push mirrors for discoverability/redundancy only.
    Do not treat mirrors as release authority unless project policy changes.

### WSL Debian active working umbrella

    [WSL-DEBIAN] umbrella: ~/repos/carbonstack_umbrella
    [WSL-DEBIAN] carbonstack: ~/repos/carbonstack_umbrella/carbonstack
    [WSL-DEBIAN] carbonstack-comms: ~/repos/carbonstack_umbrella/carbonstack-comms
    [WSL-DEBIAN] carbonstack-cypher: ~/repos/carbonstack_umbrella/carbonstack-cypher
    [WSL-DEBIAN] carbonstack-os: ~/repos/carbonstack_umbrella/carbonstack-os
    [GO-RUNNER] runner dir: ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate

### Release-generation paths used for v0.4.0

    [WSL TEMP BUILD] /tmp/carbonstack-v0.4.0-release-build
    [WSL TEMP PACKAGE SOURCE] /tmp/carbonstack-v0.4.0-release-build/stage/package
    [WSL TEMP FRESH EXTRACTION] /tmp/carbonstack-v0.4.0-release-build/extract/package
    [WSL TEMP UPLOAD ASSETS] /tmp/carbonstack-v0.4.0-release-build/upload-assets
    [WINDOWS UPLOAD STAGING] Downloads\CarbonStack-v0.4.0-upload-assets

### Live umbrella validation

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go run . --profile local-cypher
    go run . --profile doctor
    go run . --profile core --clean-generated

### Package-root validation

    cd <package-root>/carbonstack/tools/carbonstack-validate
    go run . --profile verify-checksums --root <package-root>
    go run . --profile full --root <package-root> --clean-generated

### v0.4.0 package shape

    package/
      carbonstack/
      carbonstack-comms/
      carbonstack-cypher/
      release/
        LICENSE
        manifest.json
        release-notes.md
        testing-runbook.md
        validation-freeze.md
        checksums.txt

### Go runner surfaces after v0.4.0

    carbonstack/tools/carbonstack-validate/main.go
    carbonstack/tools/carbonstack-validate/local_cypher.go
    carbonstack/tools/carbonstack-validate/checksums.go
    carbonstack/tools/carbonstack-validate/release_snapshot.go
    carbonstack/tools/carbonstack-validate/README.md
    carbonstack/tools/carbonstack-validate/go.mod

Runner profiles:

    doctor
    core
    local-cypher
    full
    release-snapshot
    write-checksums
    verify-checksums

### full profile behavior

    full:
      ReleaseSnapshot()
      LocalCypher()

    ReleaseSnapshot:
      CheckReleaseSnapshotLayout()
      StrictPreTestArtifactScan()
      VerifyReleaseChecksums()
      Core()

    Core:
      Doctor()
      ArtifactScan("pre-test")
      OpenMLS real-Cypher lifecycle
      carbonstack-comms package tests
      carbonstack-cypher package tests
      ArtifactScan("post-test")

    LocalCypher:
      Cypher-only local lifecycle
      negative stub-text/OpenMLS protocol pairing rejection
      positive opaque OpenMLS envelope/ack lifecycle
      restart against same temporary DB

### Envelope/protocol constants

Accepted opaque envelope pair:

    content_type=carbonstack.mls.application-message.v0
    protocol_version=carbonstack-openmls-sidecar-v0

Negative protocol pair validated:

    content_type=carbonstack.message.text.stub.v0
    protocol_version=carbonstack-openmls-sidecar-v0

Expected result:

    HTTP 400
    error code unsupported_protocol_version

---

## 10. v0.4.x Forward Plan

The v0.4.0 release completes the local deployability pre-release runway. The next minor-epoch work should move from release-package validation into runtime integration.

### v0.4.x primary direction

Runtime Comms OpenMLS integration:

    send/inbox stops being stub-era
    Comms deliberately uses the OpenMLS sidecar/backbone path
    Cypher integration becomes part of actual client flow
    UX remains dev/pre-alpha
    no production security claims

### v0.4.x early recon questions

    What is the current Comms send/inbox command path?
    Where does stub-era messaging still exist?
    What state is needed for OpenMLS send/inbox to be operator-usable?
    How should local-cypher validation evolve once runtime Comms uses Cypher/OpenMLS?
    When is local-backbone justified?
    What persistent state can remain dev-only?
    What should be explicitly deferred to v0.5.x vault/trust/provider-state work?

### v0.5.x likely direction

Local storage, trust state, provider state, and vault design:

    device identity state
    provider storage
    trust records
    vault boundaries
    local encryption
    backup/export policy
    recovery assumptions
    state migration

### v0.6.x+ still later

    hostile-server and abuse-resistance harnesses
    deployability/operations hardening
    documented self-pentest/adversarial validation
    claim-boundary review
    Android app
    CarbonStackOS/appliance work

---

## 11. Lean Breakpoint Summary

    CarbonStack v0.4.0 is the compressed post-minor-epoch handoff after the v0.4.0 broad local deployability pre-release and evergreen README cleanup. The v0.4.0 Gitea release is live, marked Pre-Release, attached to 24ac9fc494, and includes the intended multi-repo package plus manifest, package checksums, asset checksums, validation freeze, testing runbook, release notes, and LICENSE. The release validates the Debian/WSL Debian-first package path using verify-checksums and full; full means release-snapshot followed by local-cypher, and release-snapshot already calls core. After the release, carbonstack main advanced to 6ba7aab docs: make front readme evergreen, so the front README now points users to the Releases page for known-good artifacts and states Gitea is source of truth while GitHub mirrors are secondary. v0.3.0-v0.3.36PRIME continuity has been compressed here because the sanitized PRIME ledger remains archived. Next safest work: optional public-download verification, decide sanitized PRIME placement, then begin v0.4.x runtime Comms OpenMLS integration recon.

