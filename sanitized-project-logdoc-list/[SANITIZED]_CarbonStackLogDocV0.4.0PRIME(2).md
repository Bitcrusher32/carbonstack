# CarbonStack LogDoc v0.4.0PRIME

**Last updated:** 2026-06-03 local session  
**Owner/operator:** bitcrusher32  
**Assistant/model:** GPT-5.5 Thinking  
**Current phase:** Phase 2E / **v0.4.0 PRIME broad local deployability pre-release cut**. `carbonstack` release tag `v0.4.0` is now live on Gitea as **CarbonStack v0.4.0 Broad Local Deployability Pre-Release**, marked as a pre-release and attached to `24ac9fc494` / `24ac9fc docs: record v0.4.0 package rehearsal`. The release publishes the multi-repo v0.4.0 validation package and seven companion assets: package tarball, manifest, package checksums, asset checksums, validation freeze, testing runbook, release notes, plus the release-level `LICENSE`. The release body uses the v0.3.20 continuity style while updating the boundary to the v0.4.0 `full` validation ladder (`release-snapshot + local-cypher`) and Debian/WSL Debian-first scope. `carbonstack-comms` remains at `012c8bf scripts: support source snapshot artifact guard`; `carbonstack-cypher` remains at `9ab994c docs: point to local operator runbook`; `carbonstack-os` remains at `1bbbe52 docs: clarify CarbonStackOS target direction`.

**Version schema:** `v[scope].[timeline]` with PRIME suffix for public/case-study candidate master state. This file is `v0.4.0PRIME`, the post-release master ledger after the v0.4.0 broad local deployability pre-release. Per LogDoc V2 practice, this Markdown file intentionally preserves verbose operational continuity; the JSON breakpoint remains the lean handoff state.

**PRIME status note:** This file is intentionally labeled **PRIME** because it is intended to become a public dev ledger / LogDoc case-study source after a later sanitization pass. **Do not treat this file as sanitized.** It currently preserves exact operational process, paths, blunders, release mechanics, and continuity context. A later pass should sanitize PII and critical local topology before upload to public CarbonStack / LogDoc Loop System surfaces.

**Naming correction carried forward:** Earlier v0.3.34 filename/internal-heading mismatch remains historical provenance only. Treat the current v0.4.0PRIME overlay as authoritative current state. Any stale internal version headings preserved below are historical process logs, not current naming truth.

**LogDoc V2 usage note:** LogDoc V2 uses a split between a Markdown LogDoc for working-session continuity, a JSON state file for end-of-session handoff, Git for history preservation, and README/docs/release pages for public claims. This v0.4.0PRIME Markdown file intentionally stays “bloated” where operational process, blunders, validation ladders, package/release mechanics, and historical context matter. The v0.4.0PRIME JSON should remain lean and should agree with this file on latest current state.

---

## 1. Project Goal

**Active goal:** Preserve CarbonStack state after cutting the v0.4.0 broad local deployability pre-release and prepare for post-release cleanup: LogDoc PRIME generation, later PII/security sanitization, release-readiness continuity, and the v0.4.x runtime-integration epoch. The immediate next work should be post-release validation/review, sanitization of this PRIME ledger, and then planning the v0.4.x runtime Comms/OpenMLS integration path.

CarbonStack’s broader project intent remains:

- **CarbonStack** is the secure-communications appliance-stack project and shared doctrine/release authority.
- **CarbonStackOS** is the future constrained Android-derived appliance OS.
- **CarbonStackComms** is the text-first encrypted messaging client and current OpenMLS sidecar/relay integration surface.
- **CarbonStackCypher** is the self-hostable relay/storage stack.
- The `carbonstack` repo remains the doctrine/spec/docs/release authority; implementation repos must not contradict it without revising the spec.

**North star:**

> Its goal is not to be a better smartphone. Its goal is to stop being a smartphone.

**Core doctrine:**

> Every feature is guilty until it proves it does not add unacceptable parser, network, sensor, identity, or filesystem authority.

**Current reality after v0.4.0PRIME:** v0.4.0 is now the current public broad local deployability pre-release. v0.3.20 remains historically important as the prior runner-backed testing release and presentation/style continuity template. v0.4.0 advances the release surface from v0.3.20’s `verify-checksums` + `release-snapshot` path into the v0.4.0 `verify-checksums` + `full` package validation path:

    full = release-snapshot + local-cypher

Important meaning:

    release-snapshot validates package layout, release metadata, strict pre-test artifact scan, checksums, and core.
    release-snapshot already calls core.
    full does not call core twice.
    local-cypher validates the Cypher-only local lifecycle afterward.
    --clean-generated removes known OpenMLS sidecar generated roots after successful validation.

Recommended v0.4.0 package-validation command shape:

    cd <package-root>/carbonstack/tools/carbonstack-validate
    go run . --profile verify-checksums --root <package-root>
    go run . --profile full --root <package-root> --clean-generated

This release is still **not** a production deployability release. It does **not** prove production readiness, production E2EE, hostile-server safety, metadata privacy, Android readiness, CarbonStackOS readiness, secure vault/storage, Debian production deployability, systemd readiness, cloudflared readiness, public ingress safety, real homelab validation, external audit, or certification.

---

## 2. Current Release State

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

Release body headline:

    CarbonStack v0.4.0 broad local deployability pre-release

Release body core claims:

    Status: Pre-alpha / Experimental
    Primary artifact: Cypher + Comms OpenMLS relay backbone validation package
    Release type: Broad local deployability pre-release / research-and-development milestone
    Primary platform: Debian 13 / WSL2 Debian, linux/amd64
    Secondary validation: Not provided as part of this release surface

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

    The default Gitea Source Code ZIP/TAR.GZ downloads are only auto-generated archives of the carbonstack repo at the v0.4.0 tag.
    Use the attached v0.4.0 package, manifest, checksums, validation freeze, and testing runbook for the intended multi-repo validation package.

Testing notes in the release page:

    Start with v0.4.0-testing-runbook.md.
    Debian / WSL Debian first.
    From a fresh extracted package root:
      go run . --profile verify-checksums --root <package-root>
      go run . --profile full --root <package-root> --clean-generated
    full runs release-snapshot and local-cypher.
    release-snapshot already calls core.
    full is a validation ladder, not a deployment command.

---

## 3. Current Repo Heads

    carbonstack        24ac9fc (HEAD -> main, origin/main, origin/HEAD) docs: record v0.4.0 package rehearsal
    carbonstack-comms  012c8bf (HEAD -> main, origin/main, origin/HEAD) scripts: support source snapshot artifact guard
    carbonstack-cypher 9ab994c (HEAD -> main, origin/main, origin/HEAD) docs: point to local operator runbook
    carbonstack-os     1bbbe52 (HEAD -> main, origin/main, origin/HEAD) docs: clarify CarbonStackOS target direction

Current public release tag:

    v0.4.0 -> 24ac9fc494 / 24ac9fc docs: record v0.4.0 package rehearsal

Important previous public release tag:

    v0.3.20 -> 3eeb1a1e1a / 3eeb1a1 docs: clean public release surface

Previous mainline deployability checkpoints:

    v0.4.0 release -> 24ac9fc494 / 24ac9fc docs: record v0.4.0 package rehearsal
    v0.3.36 mainline -> 24ac9fc docs: record v0.4.0 package rehearsal
    v0.3.35 mainline -> 63ffabf docs: record full profile release validation ladder
    v0.3.34/v0.3.33 naming-correction source -> d83d2ef docs: prepare pre-v0.4.0 release surface
    v0.3.32 mainline -> cff3ab4 docs: record local-cypher negative protocol validation
    v0.3.31 mainline -> 7894278 docs: record local-cypher polish and cleanup control
    v0.3.30 mainline -> 15a3758 docs: record local-cypher runner implementation
    v0.3.30 runner  -> a98ee9f runner: add local-cypher
    v0.3.29 mainline -> 37a6d2a docs: record local-cypher validation contract
    v0.3.28 mainline -> af33139 docs: record local operator helper runner decision recon
    v0.3.27 mainline -> e516fc7 docs: record local Cypher API lifecycle proof
    v0.3.26 mainline -> 6e14fe0 docs: record local operator config data convention
    v0.3.25 mainline -> a5c6351 docs: update readme, roadmap, and historical doc
    v0.3.25 cypher  -> 9ab994c docs: point to local operator runbook
    v0.3.24 mainline -> 76f0258 docs: record Cypher schema migrations implementation
    v0.3.24 cypher  -> 42f838f db: track applied migrations

Important continuity note: v0.4.0 supersedes v0.3.20 as the current public local deployability validation release, but v0.3.20 remains historically important as the prior runner-backed testing release and presentation continuity template. v0.4.0 remains pre-alpha/experimental and must not be reframed as production.

---

## 4. Current Validated State

Validated at the v0.4.0PRIME checkpoint:

    [WSL-DEBIAN] working umbrella exists at ~/repos/carbonstack_umbrella
    [WSL-DEBIAN] carbonstack is at 24ac9fc
    [WSL-DEBIAN] carbonstack-comms is at 012c8bf
    [WSL-DEBIAN] carbonstack-cypher is at 9ab994c
    [WSL-DEBIAN] carbonstack-os is at 1bbbe52
    [PACKAGE GENERATION] final v0.4.0 release package was staged under /tmp/carbonstack-v0.4.0-release-build
    [PACKAGE GENERATION] final upload assets were copied to Windows Downloads/CarbonStack-v0.4.0-upload-assets
    [PACKAGE GENERATION] release package included carbonstack, carbonstack-comms, carbonstack-cypher, and release metadata
    [PACKAGE GENERATION] carbonstack-os was deliberately excluded from the runnable package
    [PACKAGE GENERATION] release/manifest.json, release/validation-freeze.md, release/testing-runbook.md, release/release-notes.md, release/LICENSE, and release/checksums.txt existed in the staged/extracted package
    [CHECKSUMS] write-checksums passed in the staged package and produced package checksums
    [CHECKSUMS] verify-checksums passed in the staged package
    [ASSETS] asset checksums were generated for all release upload assets
    [PACKAGE] archive creation passed
    [PACKAGE] fresh extraction passed
    [CHECKSUMS] verify-checksums passed from fresh extraction
    [FULL] full passed from the fresh extraction with --clean-generated
    [FULL] release-snapshot passed package layout, metadata, strict pre-test artifact scan, checksum verification, and core
    [FULL] local-cypher passed after release-snapshot and still covered the invalid stub-text/OpenMLS protocol pairing
    [ARTIFACT CLEANUP] known OpenMLS generated roots were absent after full --clean-generated
    [RELEASE PAGE] Gitea v0.4.0 release exists and is marked Pre-Release
    [RELEASE PAGE] release body follows v0.3.20 continuity style while using v0.4.0 claims/boundaries
    [RELEASE PAGE] eight intended release-level assets are attached
    [RELEASE PAGE] release-level LICENSE asset was kept for continuity and multi-repo package clarity

Observed WSL Debian toolchain baseline from immediately prior validation cycle:

    git version 2.47.3
    go version go1.24.4 linux/amd64
    rustc 1.96.0 (ac68faa20 2026-05-25)
    cargo 1.96.0 (30a34c682 2026-05-25)
    sqlite3 3.46.1

Observed package rehearsal / release-candidate result inherited into release:

    write-checksums: PASS / 290 package entries during rehearsal
    staged verify-checksums: PASS
    archive staged package: PASS
    fresh extraction: PASS
    fresh extraction verify-checksums: PASS
    fresh extraction full --clean-generated: PASS
    final artifact check: .carbonstack-openmls-sidecar-state absent
    final artifact check: target absent
    final artifact check: provider-storage.json absent
    final artifact check: signer.json absent

Release-page validated list:

    v0.4.0-style package source root staging
    package release metadata generation
    package-internal checksum generation
    package-internal checksum verification
    archive creation
    fresh package extraction
    fresh extracted package root layout
    fresh extracted package root release metadata checks
    fresh extracted package root strict pre-test artifact scan
    verify-checksums from fresh package extraction
    full from fresh package extraction
    release-snapshot through full
    core through release-snapshot
    OpenMLS real-Cypher lifecycle validation
    CarbonStackComms package tests
    CarbonStackCypher package tests
    local-cypher Cypher-only local lifecycle validation
    local-cypher restart/persistence check against the same temporary DB
    local-cypher rejection of historical invalid stub-text/OpenMLS protocol pairing
    post-test artifact scoping under known OpenMLS sidecar generated roots
    explicit --clean-generated cleanup of known OpenMLS generated roots
    Debian 13 / WSL2 Debian, linux/amd64 validation

---

## 5. v0.4.0PRIME Work Completed

### 5.1 v0.3.20 release was used as presentation continuity template

Before cutting v0.4.0, the old v0.3.20 release page, screenshot, link, and Markdown paste were used as the continuity template.

Kept from v0.3.20 style:

    title casing and release heading structure
    Status / Primary artifact / Release type / Primary platform / Secondary validation block
    What changed section
    Validated section
    Release assets section
    Testing notes section
    Boundary section
    explicit warning that Gitea Source Code archives are not the intended multi-repo validation package
    standalone release-level LICENSE asset

Updated for v0.4.0 reality:

    release type became broad local deployability pre-release / research-and-development milestone
    secondary validation became not provided as part of this release surface
    current validation path became verify-checksums + full
    full was explained as release-snapshot + local-cypher
    release-snapshot already calls core
    Windows validation was removed from the v0.4.0 release surface
    carbonstack-os remained excluded from runnable package
    v0.4.0-specific nonclaims were preserved

### 5.2 Release-level LICENSE asset was deliberately kept

Decision:

    keep LICENSE as a standalone release asset.

Reasoning:

    v0.3.20 shipped a release-level LICENSE asset.
    each individual repo has the MIT license, but the release package is a multi-repo artifact with its own release metadata surface.
    users may download only attached release assets/package.
    a top-level release LICENSE avoids ambiguity and preserves continuity.

This decision should not be treated as a claim of new licensing complexity. It is a clarity/continuity measure.

### 5.3 Final release assets were generated

Asset generation used WSL Debian and a WSL -> Windows Downloads staging path.

Important paths used:

    WSL build area:
    /tmp/carbonstack-v0.4.0-release-build

    WSL upload-assets folder:
    /tmp/carbonstack-v0.4.0-release-build/upload-assets

    Windows upload folder:
    Downloads\CarbonStack-v0.4.0-upload-assets

The final upload asset list:

    carbonstack-v0.4.0-broad-local-deployability-pre-release.tgz
    carbonstack-v0.4.0-release-manifest.json
    carbonstack-v0.4.0-package-checksums.txt
    carbonstack-v0.4.0-asset-checksums.txt
    carbonstack-v0.4.0-validation-freeze.md
    v0.4.0-testing-runbook.md
    v0.4.0-release-notes.md
    LICENSE

Explicit non-upload guidance preserved:

    do not upload the Gitea default Source Code ZIP/TAR.GZ as intended package assets
    do not upload the stage/package folder itself
    do not upload the extract/package folder
    do not upload carbonstack-os
    do not upload target
    do not upload .carbonstack-openmls-sidecar-state
    do not upload provider-storage.json
    do not upload signer.json

### 5.4 Final package validation passed before upload

The final package generation script:

    staged package root
    generated release metadata
    wrote package checksums
    verified package checksums in staged package
    archived the package
    copied individual release assets
    wrote asset checksums
    copied assets to Windows Downloads staging folder
    fresh-extracted the final package
    verified checksums from fresh extraction
    ran full from fresh extraction with --clean-generated
    checked known generated roots were absent afterward
    checked live umbrella final status

The package was then uploaded to Gitea as the v0.4.0 release.

### 5.5 v0.4.0 Gitea release was cut

The Gitea release exists at:

    https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.4.0

Observed from the release page:

    Release title: CarbonStack v0.4.0 Broad Local Deployability Pre-Release
    Tag: v0.4.0
    Commit: 24ac9fc494
    Pre-Release label present
    release body contains the intended v0.4.0 boundary
    release page lists all intended upload assets
    asset sizes include package tarball around 629 KiB, package checksums around 36 KiB, manifest around 2.5 KiB, validation freeze around 1.6 KiB, runbook around 1.5 KiB, release notes around 907 B, LICENSE around 1.1 KiB, and asset checksums around 703 B.

### 5.6 v0.4.0PRIME LogDoc work began after release

After release, the first cleanup task is this LogDoc update:

    v0.3.36 -> v0.4.0PRIME

Purpose:

    preserve release timeline
    preserve asset generation/upload mechanics
    preserve blunders/lessons
    preserve future to-do path
    preserve critical paths/functions
    mark as PRIME for later sanitization and upload as dev ledger / LogDoc case study

Important instruction:

    do not sanitize this PRIME file yet.
    sanitize later before public upload.

---

## 6. Critical Blunders / Lessons Preserved

### 6.1 Live package rehearsals prevented release panic

The live-umbrella `full` failure during recon was expected because release metadata did not exist. The throwaway package rehearsal turned that into a known path rather than a release-day surprise.

Correct lesson:

    release-snapshot/full are package-root disciplines, not generic live-repo commands.
    create package metadata in a staged/fresh package root.
    verify from fresh extraction.

### 6.2 Keep `full` as validation ladder, not deployment command

The v0.4.0 release page correctly says:

    full = release-snapshot + local-cypher
    release-snapshot already calls core
    full is not a deployment command

Correct lesson:

    release-package validation is not the same as deployment.
    do not let v0.4.0 local deployability language drift into production deployability claims.

### 6.3 The release-level LICENSE choice was correct

The initial question was whether to omit the MIT license file because each repo already contains the same license.

Decision:

    keep it.

Reason:

    release is a multi-repo artifact.
    release-level metadata exists separately.
    v0.3.20 included LICENSE.
    attached assets are sometimes consumed independently.

### 6.4 Windows upload staging path avoided asset mixups

The WSL build and Windows upload path split was important.

Correct lesson:

    build under /tmp
    validate from fresh extraction
    copy only upload-assets into Windows Downloads staging
    upload only the final eight files

This avoided uploading stage/extract trees, generated roots, or unrelated repo files.

### 6.5 carbonstack-os stayed out of the runnable package

`carbonstack-os` remains related future work, not part of the runnable v0.4.0 package.

Correct lesson:

    keep OS vision visible but not falsely packaged.
    do not let CarbonStackOS inflate the v0.4.0 release claim.

### 6.6 Pre-release flag matters

The release is properly marked Pre-Release.

Correct lesson:

    all v0.4.0 public presentation should keep pre-alpha/experimental/pre-release framing.
    do not let the minor epoch number imply maturity.

### 6.7 PRIME is not sanitized

This LogDoc is intended as a future public case-study source, but not yet safe to publish.

Correct lesson:

    first preserve full continuity.
    then run a separate sanitization pass.
    do not accidentally upload raw PRIME if it contains local usernames, critical paths, or other PII/operational topology.

---

## 7. Current Blockers / Not Validated

Current blockers and nonvalidated areas after v0.4.0 release:

    Post-release validation from a third/fresh independent download has not yet been performed unless done manually after upload.
    The v0.4.0 release page has been checked through Gitea rendered HTML, but a full downloaded-asset verification from the public release page would be a useful follow-up.
    This v0.4.0PRIME LogDoc is not sanitized yet.
    The v0.4.0PRIME JSON and Markdown have not yet been committed/uploaded to carbonstack/logdoc-loop-system.
    Public case-study upload must wait for sanitization.
    Runtime Comms send/inbox remains stub-era and not wired as user-facing OpenMLS messenger UX.
    v0.4.x runtime Comms OpenMLS integration has not begun.
    local-cypher implements positive Cypher-only lifecycle validation plus one negative protocol-pairing check, but no broad negative-path suite exists yet.
    local-backbone runner profile does not exist and remains reserved for later whole-stack validation.
    No helper command exists yet.
    Production-safe provider storage and secure vault design are unsolved.
    Hostile-server rollback/replay/metadata-abuse validation remains future work.
    Android/CarbonStackOS implementation remains later work.
    CarbonStack Relay Space mechanics remain a defined future concept, not an implementation.
    Public ingress, systemd, cloudflared, real homelab deployment, and production service paths remain out of scope.
    Temporary DB validation remains dev/test-only and does not replace later persistent user/operator-state validation for public Comms maturity.

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
    Do not claim v0.4.0 proves production deployability.
    Do not publish raw PRIME before sanitization.

Updated allowed claim:

    CarbonStack v0.4.0 is a public broad local deployability pre-release.
    v0.4.0 publishes a multi-repo validation package containing carbonstack, carbonstack-comms, carbonstack-cypher, and release metadata.
    carbonstack-os is related future work but excluded from the runnable package.
    v0.4.0 validates the intended package path with verify-checksums and full.
    full means release-snapshot followed by local-cypher.
    release-snapshot already calls core.
    v0.4.0 is Debian / WSL Debian-first.
    v0.4.0 has no Windows validation as part of this release surface.
    v0.4.0 remains pre-alpha/experimental and not production secure.

---

## 8. Next Safest Actions

Recommended immediate post-release sequence:

    v0.4.0PRIME:
      preserve unsanitized full release continuity in LogDoc and lean JSON

    v0.4.0PRIME-sanitized:
      sanitize PII/local topology/security-sensitive paths before public upload

    post-release verification:
      download the attached v0.4.0 package/assets from Gitea
      verify asset checksums
      extract package fresh
      run verify-checksums
      run full --clean-generated
      record result

    repo/docs cleanup:
      decide where sanitized v0.4.0PRIME belongs in carbonstack
      decide where sanitized case study belongs in logdoc-loop-system
      add pointers without overloading current release docs

Then begin v0.4.x planning:

    v0.4.1 or v0.4.x early:
      runtime Comms OpenMLS integration recon

    v0.4.x:
      send/inbox stops being stub-era
      Comms deliberately uses the OpenMLS sidecar/backbone path
      Cypher integration becomes part of actual client flow
      still pre-alpha/dev UX
      no production security claims

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
    broad negative-path suite unless v0.4.x runtime integration exposes a direct blocker

---

## 9. Critical Paths / Functions

### WSL Debian active working umbrella

    [WSL-DEBIAN] umbrella: ~/repos/carbonstack_umbrella
    [WSL-DEBIAN] carbonstack: ~/repos/carbonstack_umbrella/carbonstack
    [WSL-DEBIAN] carbonstack-comms: ~/repos/carbonstack_umbrella/carbonstack-comms
    [WSL-DEBIAN] carbonstack-cypher: ~/repos/carbonstack_umbrella/carbonstack-cypher
    [WSL-DEBIAN] carbonstack-os: ~/repos/carbonstack_umbrella/carbonstack-os
    [GO-RUNNER] runner dir: ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate

### Release generation paths used for v0.4.0

    [WSL TEMP BUILD] /tmp/carbonstack-v0.4.0-release-build
    [WSL TEMP PACKAGE SOURCE] /tmp/carbonstack-v0.4.0-release-build/stage/package
    [WSL TEMP FRESH EXTRACTION] /tmp/carbonstack-v0.4.0-release-build/extract/package
    [WSL TEMP UPLOAD ASSETS] /tmp/carbonstack-v0.4.0-release-build/upload-assets
    [WINDOWS UPLOAD STAGING] Downloads\CarbonStack-v0.4.0-upload-assets

### v0.4.0 release assets

    [ASSET] carbonstack-v0.4.0-broad-local-deployability-pre-release.tgz
    [ASSET] carbonstack-v0.4.0-release-manifest.json
    [ASSET] carbonstack-v0.4.0-package-checksums.txt
    [ASSET] carbonstack-v0.4.0-asset-checksums.txt
    [ASSET] carbonstack-v0.4.0-validation-freeze.md
    [ASSET] v0.4.0-testing-runbook.md
    [ASSET] v0.4.0-release-notes.md
    [ASSET] LICENSE

### WSL Debian known-good validation commands

Live umbrella validation:

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go run . --profile local-cypher
    go run . --profile doctor
    go run . --profile core --clean-generated

Package-root validation:

    cd <package-root>/carbonstack/tools/carbonstack-validate
    go run . --profile verify-checksums --root <package-root>
    go run . --profile full --root <package-root> --clean-generated

Post-upload public-asset validation pattern:

    download carbonstack-v0.4.0-broad-local-deployability-pre-release.tgz
    download carbonstack-v0.4.0-asset-checksums.txt
    optionally download companion metadata assets
    verify asset checksums
    extract package to a fresh root
    cd <package-root>/carbonstack/tools/carbonstack-validate
    go run . --profile verify-checksums --root <package-root>
    go run . --profile full --root <package-root> --clean-generated

### v0.4.0 docs / release surfaces

    [RELEASE] v0.4.0 Gitea release:
    https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.4.0

    [DOC] v0.3.36 v0.4.0 package rehearsal:
    carbonstack/docs/154-v0.4.0-package-rehearsal-v0.md

    [DOC] v0.3.35 full profile release validation ladder:
    carbonstack/docs/153-full-profile-release-validation-ladder-v0.md

    [DOC] v0.3.34/v0.3.33 naming-correction source / pre-v0.4.0 release-surface cleanup:
    carbonstack/docs/152-pre-v0.4.0-release-surface-cleanup-v0.md

    [DOC] v0.3.32 negative protocol validation result:
    carbonstack/docs/151-local-cypher-negative-protocol-validation-v0.md

    [DOC] docs index:
    carbonstack/docs/README.md

    [DOC] roadmap:
    carbonstack/roadmap/ROADMAP.md

    [DOC] top-level README:
    carbonstack/README.md

### Go runner surfaces after v0.4.0

    [RUNNER] directory:
    carbonstack/tools/carbonstack-validate

    [RUNNER] implementation files:
    carbonstack/tools/carbonstack-validate/main.go
    carbonstack/tools/carbonstack-validate/local_cypher.go
    carbonstack/tools/carbonstack-validate/checksums.go
    carbonstack/tools/carbonstack-validate/release_snapshot.go
    carbonstack/tools/carbonstack-validate/README.md
    carbonstack/tools/carbonstack-validate/go.mod

    [RUNNER] profiles:
    doctor
    core
    local-cypher
    full
    release-snapshot
    write-checksums
    verify-checksums

    [RUNNER] cleanup flag:
    --clean-generated

### full profile behavior after v0.4.0

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

### --clean-generated behavior

    Runs only after successful profile execution.
    Deletes only:
      carbonstack-comms/internal/protocol/mls/openmls-sidecar/.carbonstack-openmls-sidecar-state
      carbonstack-comms/internal/protocol/mls/openmls-sidecar/target

    Does not delete:
      arbitrary untracked files
      manual local operator DBs
      $HOME/.local/share/carbonstack/cypher/cypher.db
      provider-storage.json unless added by a future deliberate expansion
      signer.json unless added by a future deliberate expansion

    Does not replace:
      ArtifactScan
      StrictPreTestArtifactScan
      release-snapshot pre-test hygiene checks

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

### Accepted envelope pair

    content_type=carbonstack.mls.application-message.v0
    protocol_version=carbonstack-openmls-sidecar-v0

### Negative protocol pair validated

    content_type=carbonstack.message.text.stub.v0
    protocol_version=carbonstack-openmls-sidecar-v0

Expected result:

    HTTP 400
    error code unsupported_protocol_version

---

## 10. v0.4.x Forward Plan

The v0.4.0 release completes the local deployability pre-release runway. The next minor-epoch work should move from release-package validation into runtime integration planning.

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

    CarbonStack v0.4.0PRIME is the post-release master ledger after the v0.4.0 broad local deployability pre-release cut. The v0.4.0 Gitea release is live at the carbonstack repo, marked Pre-Release, attached to 24ac9fc494 / 24ac9fc docs: record v0.4.0 package rehearsal, and includes the intended multi-repo package plus manifest, package checksums, asset checksums, validation freeze, testing runbook, release notes, and LICENSE. The release validates the Debian/WSL Debian-first package path using verify-checksums and full; full means release-snapshot followed by local-cypher, and release-snapshot already calls core. The release remains pre-alpha/experimental and does not prove production readiness, production E2EE, hostile-server safety, metadata privacy, Android/CarbonStackOS readiness, secure vault/storage, deployability, public ingress, audit, or certification. This PRIME file is intentionally unsanitized and should be sanitized before public upload as a dev ledger / LogDoc case study. Next safest work: sanitize v0.4.0PRIME, optionally perform post-upload fresh download verification, then begin v0.4.x runtime Comms OpenMLS integration recon.

---

## 12. Preserved v0.3.36 Operational Process Log

The following section preserves the previous v0.3.36 LogDoc body as operational continuity. It is intentionally retained rather than compressed away. Where preserved text conflicts with the v0.4.0PRIME current-state overlay above, v0.4.0PRIME wins for current state; the preserved body remains the provenance/process ledger for the v0.4.0 package rehearsal work and earlier history.



# CarbonStack LogDoc v0.3.36

**Last updated:** 2026-06-03 local session  
**Owner/operator:** bitcrusher32  
**Assistant/model:** GPT-5.5 Thinking  
**Current phase:** Phase 2E-prep / v0.3.x **v0.4.0-style package rehearsal checkpoint**. `carbonstack` is now at `24ac9fc docs: record v0.4.0 package rehearsal`, after the v0.3.36 rung staged a throwaway v0.4.0-style package root, generated release metadata/checksums, archived it, fresh-extracted it, verified checksums from the extraction, and validated the package through the `full` profile with `--clean-generated`. `carbonstack-cypher` remains at `9ab994c docs: point to local operator runbook`; `carbonstack-comms` remains at `012c8bf scripts: support source snapshot artifact guard`; `carbonstack-os` remains at `1bbbe52 docs: clarify CarbonStackOS target direction`. The live WSL Debian working umbrella at `~/repos/carbonstack_umbrella` passed `local-cypher`, `doctor`, and `core --clean-generated` after the v0.3.36 docs commit. Final clean status showed no dirty files across the four repos.

**Version schema:** `v[scope].[timeline]`. This file is `v0.3.36`, the v0.4.0-style package rehearsal handoff after the `v0.3.35` full-profile release validation ladder checkpoint. Per LogDoc V2 practice, this Markdown file intentionally preserves verbose operational continuity; the JSON breakpoint remains the lean handoff state.

**Naming correction carried forward:** The earlier `v0.3.34` filename/internal-heading mismatch remains historical provenance only. Treat the current overlay in this v0.3.36 file as authoritative current state. Any stale internal version headings preserved below are historical process logs, not current naming truth.

**LogDoc V2 usage note:** LogDoc V2 uses a split between a Markdown LogDoc for working-session continuity, a JSON state file for end-of-session handoff, Git for history preservation, and README/docs for public claims. This v0.3.36 Markdown file intentionally stays “bloated” where operational process, blunders, validation ladders, package rehearsal mechanics, and historical context matter. The v0.3.36 JSON should remain lean and should agree with this file on the latest current state.

---

## 1. Project Goal

**Active goal:** Preserve CarbonStack state after proving the v0.4.0-style package path in a throwaway rehearsal. The immediate next line of work should be v0.3.37 release-note and asset-plan prep, using the old v0.3.20 release page/style as continuity template once provided. The release should not be cut until final template-aligned notes, final package naming, asset expectations, and a final validation pass are complete.

CarbonStack’s broader project intent remains:

- **CarbonStack** is the secure-communications appliance-stack project and shared doctrine/release authority.
- **CarbonStackOS** is the future constrained Android-derived appliance OS.
- **CarbonStackComms** is the text-first encrypted messaging client and current OpenMLS sidecar/relay integration surface.
- **CarbonStackCypher** is the self-hostable relay/storage stack.
- The `carbonstack` repo remains the doctrine/spec/docs/release authority; implementation repos must not contradict it without revising the spec.

**North star:**

> Its goal is not to be a better smartphone. Its goal is to stop being a smartphone.

**Core doctrine:**

> Every feature is guilty until it proves it does not add unacceptable parser, network, sensor, identity, or filesystem authority.

**Current reality after v0.3.36:** `v0.3.20` remains the current public runner-backed testing release. `v0.3.21` through `v0.3.35` built the local-only deployability runway, implemented `local-cypher`, added explicit cleanup, added one negative protocol validation, cleaned the public release surface, and converted `full` from a stale `core` alias into the release-package validation ladder. v0.3.36 now proves the package mechanics in a throwaway rehearsal:

    staged package root under /tmp/carbonstack-v0.4.0-rehearsal/stage/package
    included carbonstack, carbonstack-comms, carbonstack-cypher
    excluded carbonstack-os from the runnable package
    created release metadata
    wrote release/checksums.txt
    verified checksums in staged package
    archived package
    fresh-extracted archive
    verified checksums from fresh extraction
    ran full from fresh extraction with --clean-generated
    confirmed known generated roots were absent after cleanup

Important meaning:

    full = release-snapshot + local-cypher
    release-snapshot already calls core
    full does not call core twice
    release-snapshot validates package layout, release metadata, strict pre-test artifact scan, checksums, and core
    local-cypher validates the Cypher-only lifecycle afterward
    --clean-generated removes known OpenMLS sidecar generated roots after the successful profile

Recommended v0.4.0-style package-validation command shape:

    cd <package-root>/carbonstack/tools/carbonstack-validate
    go run . --profile verify-checksums --root <package-root>
    go run . --profile full --root <package-root> --clean-generated

This is still **not** a production deployability release. It does **not** prove production readiness, production E2EE, hostile-server safety, metadata privacy, Android readiness, secure vault/storage, Debian production deployability, systemd readiness, cloudflared readiness, public ingress safety, real homelab validation, external audit, or certification.

---

## 2. Current Repo Heads

    carbonstack        24ac9fc (HEAD -> main, origin/main, origin/HEAD) docs: record v0.4.0 package rehearsal
    carbonstack-comms  012c8bf (HEAD -> main, origin/main, origin/HEAD) scripts: support source snapshot artifact guard
    carbonstack-cypher 9ab994c (HEAD -> main, origin/main, origin/HEAD) docs: point to local operator runbook
    carbonstack-os     1bbbe52 (HEAD -> main, origin/main, origin/HEAD) docs: clarify CarbonStackOS target direction

Current `carbonstack` public runner-backed testing release tag:

    v0.3.20 -> 3eeb1a1e1a / 3eeb1a1 docs: clean public release surface

Previous mainline deployability checkpoints:

    v0.3.36 mainline -> 24ac9fc docs: record v0.4.0 package rehearsal
    v0.3.35 mainline -> 63ffabf docs: record full profile release validation ladder
    v0.3.34/v0.3.33 naming-correction source -> d83d2ef docs: prepare pre-v0.4.0 release surface
    v0.3.32 mainline -> cff3ab4 docs: record local-cypher negative protocol validation
    v0.3.31 mainline -> 7894278 docs: record local-cypher polish and cleanup control
    v0.3.30 mainline -> 15a3758 docs: record local-cypher runner implementation
    v0.3.30 runner  -> a98ee9f runner: add local-cypher
    v0.3.29 mainline -> 37a6d2a docs: record local-cypher validation contract
    v0.3.28 mainline -> af33139 docs: record local operator helper runner decision recon
    v0.3.27 mainline -> e516fc7 docs: record local Cypher API lifecycle proof
    v0.3.26 mainline -> 6e14fe0 docs: record local operator config data convention
    v0.3.25 mainline -> a5c6351 docs: update readme, roadmap, and historical doc
    v0.3.25 cypher  -> 9ab994c docs: point to local operator runbook
    v0.3.24 mainline -> 76f0258 docs: record Cypher schema migrations implementation
    v0.3.24 cypher  -> 42f838f db: track applied migrations

Important continuity note: v0.3.36 does not supersede v0.3.20 as a public release. It advances mainline release readiness by proving the package rehearsal mechanics and documenting `docs/154-v0.4.0-package-rehearsal-v0.md`. It is safe to pause here because the package rehearsal passed, the docs/result record was pushed, WSL Debian `local-cypher` / `doctor` / `core --clean-generated` passed after the docs commit, and final status was clean across all four repos.

---

## 3. Current Validated State

Validated at the v0.3.36 checkpoint:

    [WSL-DEBIAN] working umbrella exists at ~/repos/carbonstack_umbrella
    [WSL-DEBIAN] carbonstack is at 24ac9fc
    [WSL-DEBIAN] carbonstack-comms is at 012c8bf
    [WSL-DEBIAN] carbonstack-cypher is at 9ab994c
    [WSL-DEBIAN] carbonstack-os is at 1bbbe52
    [PACKAGE REHEARSAL] throwaway v0.4.0-style package root was staged under /tmp/carbonstack-v0.4.0-rehearsal
    [PACKAGE REHEARSAL] package included carbonstack, carbonstack-comms, carbonstack-cypher, and release metadata
    [PACKAGE REHEARSAL] carbonstack-os was deliberately excluded from the runnable package
    [PACKAGE REHEARSAL] release/manifest.json, release/validation-freeze.md, release/testing-runbook.md, release/release-notes.md, release/LICENSE, and release/checksums.txt existed in the staged/extracted package
    [CHECKSUMS] write-checksums passed in the staged package and wrote 290 entries
    [CHECKSUMS] verify-checksums passed in the staged package
    [PACKAGE] archive creation passed
    [PACKAGE] fresh extraction passed
    [CHECKSUMS] verify-checksums passed from fresh extraction
    [FULL] full passed from the fresh extraction with --clean-generated
    [FULL] release-snapshot passed package layout, metadata, strict pre-test artifact scan, checksum verification, and core
    [FULL] local-cypher passed after release-snapshot and still printed PASS: reject invalid stub-text/OpenMLS protocol pairing
    [ARTIFACT CLEANUP] known OpenMLS generated roots were absent after full --clean-generated
    [CARBONSTACK] docs/154-v0.4.0-package-rehearsal-v0.md exists and records the v0.3.36 result
    [GO-RUNNER] local-cypher passed after the v0.3.36 docs commit
    [GO-RUNNER] doctor passed after the v0.3.36 docs commit
    [GO-RUNNER] core --clean-generated passed after the v0.3.36 docs commit
    [FINAL STATUS] all four repos were clean after validation without a separate manual cleanup step

Observed WSL Debian toolchain baseline:

    git version 2.47.3
    go version go1.24.4 linux/amd64
    rustc 1.96.0 (ac68faa20 2026-05-25)
    cargo 1.96.0 (30a34c682 2026-05-25)
    sqlite3 3.46.1

Observed package rehearsal result:

    write-checksums: PASS / 290 entries
    staged verify-checksums: PASS
    archive staged package: PASS
    fresh extraction: PASS
    fresh extraction verify-checksums: PASS
    fresh extraction full --clean-generated: PASS
    final artifact check: .carbonstack-openmls-sidecar-state absent
    final artifact check: target absent
    final artifact check: provider-storage.json absent
    final artifact check: signer.json absent

Observed live `local-cypher` validation result during the v0.3.36 final flow:

    required paths: PASS
    pre-local-cypher artifact scan: PASS / no generated/private/build artifact hits
    temporary Cypher binary build: PASS
    first health check: PASS
    seeded Alice invite claim: PASS
    Bob dev invite creation: PASS
    Bob invite claim: PASS
    Alice device registration: PASS
    Bob device registration: PASS
    Alice device listing: PASS
    invalid stub-text/OpenMLS protocol pairing rejection: PASS
    opaque OpenMLS application-message envelope submit: PASS
    Bob inbox retrieval and payload metadata verification: PASS
    envelope ack: PASS
    Bob inbox after ack empty: PASS
    restart against same temporary DB: PASS
    restart health check: PASS
    persisted Alice device state after restart: PASS
    Bob's acked inbox remains empty after restart: PASS
    post-local-cypher artifact scan: PASS / no generated/private/build artifact hits
    VALIDATION PASSED

Observed live `core --clean-generated` validation result during the v0.3.36 final flow:

    doctor nested inside core: PASS
    pre-test artifact scan: PASS / no generated/private/build artifact hits
    OpenMLS real-Cypher lifecycle: PASS
    carbonstack-comms package tests: PASS
    carbonstack-cypher package tests: PASS
    post-test artifact scan: expected OpenMLS sidecar generated roots only
    clean generated artifacts step: PASS
    removed .carbonstack-openmls-sidecar-state
    removed target
    VALIDATION PASSED

Expected/generated roots handled by `--clean-generated`:

    carbonstack-comms/internal/protocol/mls/openmls-sidecar/.carbonstack-openmls-sidecar-state
    carbonstack-comms/internal/protocol/mls/openmls-sidecar/target

---

## 4. v0.3.36 Work Completed

### 4.1 Package rehearsal recon identified expected metadata gap

Initial package/release recon showed that the live umbrella had a valid repo layout and passing live validation, but was not a true release package root yet.

Observed missing release metadata in the live umbrella before staging:

    release/manifest.json
    release/validation-freeze.md
    release/testing-runbook.md
    release/release-notes.md
    release/LICENSE

This was not catastrophic. It confirmed that the live umbrella itself was not the release package. The next correct move was a throwaway staged package rehearsal.

### 4.2 Throwaway v0.4.0-style package root was created

A temporary package rehearsal root was created under:

    /tmp/carbonstack-v0.4.0-rehearsal

Package shape:

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

Included repos:

    carbonstack
    carbonstack-comms
    carbonstack-cypher

Excluded repo:

    carbonstack-os

Reason for excluding carbonstack-os:

    related future OS repo, not part of the runnable v0.4.0 validation package.

### 4.3 Checksums, archive, fresh extraction, and full validation passed

The rehearsal flow succeeded:

    go run . --profile write-checksums --root /tmp/carbonstack-v0.4.0-rehearsal/stage/package
    go run . --profile verify-checksums --root /tmp/carbonstack-v0.4.0-rehearsal/stage/package
    tar -czf /tmp/carbonstack-v0.4.0-rehearsal/carbonstack-v0.4.0-rehearsal-package.tar.gz package
    tar -xzf /tmp/carbonstack-v0.4.0-rehearsal/carbonstack-v0.4.0-rehearsal-package.tar.gz
    go run . --profile verify-checksums --root /tmp/carbonstack-v0.4.0-rehearsal/extract/package
    go run . --profile full --root /tmp/carbonstack-v0.4.0-rehearsal/extract/package --clean-generated

The fresh extracted package passed:

    release-snapshot layout checks
    release metadata checks
    strict pre-test artifact scan
    release checksum verification
    core validation
    local-cypher validation
    final generated-root cleanup

### 4.4 Important observation: local-cypher can see known generated roots after release-snapshot/core

During `full`, `release-snapshot` runs `core` before `local-cypher`. `core` generates OpenMLS sidecar state and Rust build artifacts.

Result:

    local-cypher's pre-local-cypher artifact scan can see known generated roots after release-snapshot/core.

This was not a release blocker because:

    artifact scan classified them as known OpenMLS sidecar generated roots
    validation continued
    --clean-generated removed the known roots after the successful full profile
    final artifact check showed those roots absent

Lesson:

    do not treat this as a failed package rehearsal.
    do not patch cleanup boundaries unless a later release candidate validation shows this creates a claim or usability problem.
    current behavior is acceptable for v0.4.0 pre-release if documented honestly.

### 4.5 v0.3.36 result docs landed

New/updated docs at v0.3.36:

    carbonstack/docs/154-v0.4.0-package-rehearsal-v0.md
    carbonstack/docs/README.md
    carbonstack/roadmap/ROADMAP.md

Final mainline docs commit:

    24ac9fc docs: record v0.4.0 package rehearsal

The docs record:

    throwaway package shape
    staging/fresh extraction flow
    checksum generation and verification
    full --root <package-root> --clean-generated result
    release-snapshot/core/local-cypher interaction
    known generated-root observation
    allowed claim and forbidden claims
    next rung as release-note / asset-plan prep

### 4.6 Final validation passed with cleanup

Final WSL Debian validation sequence included:

    go run . --profile local-cypher
    go run . --profile doctor
    go run . --profile core --clean-generated

Final WSL Debian repo snapshot:

    carbonstack        24ac9fc docs: record v0.4.0 package rehearsal
    carbonstack-comms  012c8bf scripts: support source snapshot artifact guard
    carbonstack-cypher 9ab994c docs: point to local operator runbook
    carbonstack-os     1bbbe52 docs: clarify CarbonStackOS target direction

Final validation passed under WSL Debian. `local-cypher` passed and still included the negative-path line. `doctor` passed. `core --clean-generated` passed OpenMLS real-Cypher lifecycle, Comms package tests, and Cypher package tests, then removed the expected OpenMLS generated roots through the explicit cleanup flag. Final clean status showed no dirty files across all four repos.

---

## 5. Critical Blunders / Lessons Preserved

### 5.1 Missing live-umbrella release metadata was expected, not catastrophic

The first recon reported missing release metadata under `~/repos/carbonstack_umbrella/release`.

Correct interpretation:

    the live umbrella is not itself the release package root.
    release metadata should be created in a staged package/release directory for the release package.
    failure of full against a live umbrella missing release metadata is a useful guard, not a disaster.

### 5.2 `write-checksums` can mutate release/checksums.txt

The recon proved `write-checksums` can create or rewrite `release/checksums.txt` in the root it is pointed at.

Correct lesson:

    use throwaway staged package roots for package rehearsal.
    avoid pointing write-checksums at the live umbrella unless intentionally staging release metadata there.
    clean accidental live-umbrella release metadata if it appears.

### 5.3 Fresh extraction validation passed

The most important v0.3.36 result is not just that the live repo tests passed. It is that a tarred package could be fresh-extracted and validated using the intended package command.

Allowed claim:

    a throwaway v0.4.0-style package rehearsal passed checksum verification and full runner validation from a fresh extraction under WSL Debian.

### 5.4 `full` generated known artifacts before local-cypher, but cleanup handled them

Observed behavior:

    release-snapshot -> core generated known OpenMLS sidecar roots.
    local-cypher ran afterward and observed those known roots in artifact scans.
    --clean-generated removed them at the end.

Correct lesson:

    this is acceptable for now because the roots are known/classified and cleanup succeeded.
    do not rush a cleanup-boundary patch without a real release-candidate blocker.

### 5.5 carbonstack-os remains related, not packaged

The rehearsal package deliberately included only:

    carbonstack
    carbonstack-comms
    carbonstack-cypher

Correct lesson:

    carbonstack-os should remain linked/related, but not part of the runnable v0.4.0 validation package until OS work becomes real.

---

## 6. Current Blockers / Not Validated

Current blockers and nonvalidated areas:

    v0.4.0 release notes are not drafted in final release-page style yet.
    The old v0.3.20 release format/template has not yet been provided/applied.
    Final v0.4.0 release asset names are not finalized.
    Final v0.4.0 package artifact has not yet been generated as the actual release asset.
    No Gitea release has been cut for v0.4.0.
    No final release-page Markdown has been posted or validated visually.
    local-cypher implements positive Cypher-only lifecycle validation plus one negative protocol-pairing check, but no broad negative-path suite exists yet.
    local-backbone runner profile does not exist and remains reserved for later whole-stack validation.
    No helper command exists yet.
    Runtime Comms send/inbox remains stub-era and not wired as user-facing OpenMLS messenger UX.
    Production-safe provider storage and secure vault design are unsolved.
    Hostile-server rollback/replay/metadata-abuse validation remains future work.
    Android/CarbonStackOS implementation remains later work.
    CarbonStack Relay Space mechanics remain a defined future concept, not an implementation.
    Public ingress, systemd, cloudflared, real homelab deployment, and production service paths remain out of scope.
    Temporary DB validation remains dev/test-only and does not replace later persistent user/operator-state validation for public Comms maturity.

Hard nonclaims remain:

    Do not claim production readiness.
    Do not claim production E2EE.
    Do not claim hostile-server safety.
    Do not claim metadata privacy.
    Do not claim Debian production deployability.
    Do not claim systemd readiness.
    Do not claim cloudflared readiness.
    Do not claim Android readiness.
    Do not claim secure vault/storage.
    Do not claim external audit or certification.
    Do not claim CarbonStack Relay Space mechanics are implemented.
    Do not claim local-backbone exists or is justified yet.
    Do not claim local-cypher validates runtime Comms UX.
    Do not claim full validates local-backbone.
    Do not claim full is a deployment command.
    Do not claim --clean-generated is a general cleanup tool.
    Do not claim v0.3.36 is itself the v0.4.0 release.

Updated allowed claim:

    CarbonStack v0.3.36 records a throwaway v0.4.0-style package rehearsal.
    The package rehearsal staged carbonstack, carbonstack-comms, carbonstack-cypher, and release metadata.
    carbonstack-os was excluded from the runnable package.
    write-checksums and verify-checksums passed.
    A package archive was created and fresh-extracted.
    verify-checksums passed from the fresh extraction.
    full --root <fresh package> --clean-generated passed from the fresh extraction.
    local-cypher, doctor, and core --clean-generated passed under WSL Debian after the v0.3.36 docs commit.
    Final clean status showed no dirty files across the four repos.

---

## 7. Next Safest Actions

Recommended next sequence after this v0.3.36 breakpoint:

    v0.3.37 preferred:
      release-note and asset-plan prep using the old v0.3.20 release as presentation/continuity template

Recommended concrete next rung:

    1. User provides old v0.3.20 Gitea release link, screenshot, and Markdown paste.
    2. Compare v0.3.20 release format against current v0.4.0 intended claims.
    3. Draft v0.4.0 release notes in matching style.
    4. Define final release asset names.
    5. Define final package generation commands.
    6. Define final validation instructions for users/testers.
    7. Define exact known-good toolchain block.
    8. Preserve all nonclaims.
    9. Do not cut v0.4.0 until final release package and release text are both ready.

Likely late-v0.3.x path:

    v0.3.37:
      v0.4.0 release-note and asset-plan prep

    v0.3.38:
      final release candidate package generation and validation

    v0.3.39 / v0.4.0:
      release cut if final package, final notes, checksums, and validation are clean

Avoid next:

    local-backbone runner profile
    helper implementation
    public ingress
    cloudflared
    systemd
    real homelab deployment
    runtime Comms send/inbox UX
    Android app
    CarbonStackOS
    CarbonStack Relay Space implementation
    production release
    broad negative-path suite unless final release validation exposes a direct blocker

---

## 8. Critical Paths / Functions

### WSL Debian active working umbrella

    [WSL-DEBIAN] umbrella: ~/repos/carbonstack_umbrella
    [WSL-DEBIAN] carbonstack: ~/repos/carbonstack_umbrella/carbonstack
    [WSL-DEBIAN] carbonstack-comms: ~/repos/carbonstack_umbrella/carbonstack-comms
    [WSL-DEBIAN] carbonstack-cypher: ~/repos/carbonstack_umbrella/carbonstack-cypher
    [WSL-DEBIAN] carbonstack-os: ~/repos/carbonstack_umbrella/carbonstack-os
    [GO-RUNNER] runner dir: ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate

### WSL Debian known-good validation commands

Live umbrella validation:

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go run . --profile local-cypher
    go run . --profile doctor
    go run . --profile core --clean-generated

Package-root validation:

    cd <package-root>/carbonstack/tools/carbonstack-validate
    go run . --profile verify-checksums --root <package-root>
    go run . --profile full --root <package-root> --clean-generated

### v0.3.36 docs

    [DOC] v0.3.36 v0.4.0 package rehearsal:
    carbonstack/docs/154-v0.4.0-package-rehearsal-v0.md

    [DOC] v0.3.35 full profile release validation ladder:
    carbonstack/docs/153-full-profile-release-validation-ladder-v0.md

    [DOC] v0.3.34/v0.3.33 naming-correction source / pre-v0.4.0 release-surface cleanup:
    carbonstack/docs/152-pre-v0.4.0-release-surface-cleanup-v0.md

    [DOC] v0.3.32 negative protocol validation result:
    carbonstack/docs/151-local-cypher-negative-protocol-validation-v0.md

    [DOC] docs index:
    carbonstack/docs/README.md

    [DOC] roadmap:
    carbonstack/roadmap/ROADMAP.md

    [DOC] top-level README:
    carbonstack/README.md

### Go runner surfaces after v0.3.36

    [RUNNER] directory:
    carbonstack/tools/carbonstack-validate

    [RUNNER] implementation files:
    carbonstack/tools/carbonstack-validate/main.go
    carbonstack/tools/carbonstack-validate/local_cypher.go
    carbonstack/tools/carbonstack-validate/checksums.go
    carbonstack/tools/carbonstack-validate/release_snapshot.go
    carbonstack/tools/carbonstack-validate/README.md
    carbonstack/tools/carbonstack-validate/go.mod

    [RUNNER] profiles:
    doctor
    core
    local-cypher
    full
    release-snapshot
    write-checksums
    verify-checksums

    [RUNNER] cleanup flag:
    --clean-generated

### full profile behavior after v0.3.36

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

### --clean-generated behavior

    Runs only after successful profile execution.
    Deletes only:
      carbonstack-comms/internal/protocol/mls/openmls-sidecar/.carbonstack-openmls-sidecar-state
      carbonstack-comms/internal/protocol/mls/openmls-sidecar/target

    Does not delete:
      arbitrary untracked files
      manual local operator DBs
      $HOME/.local/share/carbonstack/cypher/cypher.db
      provider-storage.json unless added by a future deliberate expansion
      signer.json unless added by a future deliberate expansion

    Does not replace:
      ArtifactScan
      StrictPreTestArtifactScan
      release-snapshot pre-test hygiene checks

### Rehearsed v0.4.0-style package shape

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

### Current mainline public-surface state

    Top-level README:
      v0.3.20 remains current public testing release.
      v0.3.32/v0.3.36 is current mainline validation state.
      v0.4.0 is framed as a broad local deployability pre-release.

    Roadmap:
      late v0.3.x now points toward release-note/asset-plan prep and then final release candidate package validation.

    docs/README:
      current release/validation docs list now includes docs/154.

### Accepted envelope pair

    content_type=carbonstack.mls.application-message.v0
    protocol_version=carbonstack-openmls-sidecar-v0

### Negative protocol pair now validated

    content_type=carbonstack.message.text.stub.v0
    protocol_version=carbonstack-openmls-sidecar-v0

Expected result:

    HTTP 400
    error code unsupported_protocol_version

---

## 9. Lean Breakpoint Summary

    CarbonStack v0.3.36 is a stable mainline v0.4.0-style package rehearsal breakpoint after the v0.3.35 full-profile release validation ladder checkpoint. It proves that a throwaway package containing carbonstack, carbonstack-comms, carbonstack-cypher, and release metadata can be staged, checksum-covered, archived, fresh-extracted, checksum-verified, and validated through full --root <package-root> --clean-generated under WSL Debian. carbonstack-os remains related but excluded from the runnable package. carbonstack is now at 24ac9fc docs: record v0.4.0 package rehearsal. local-cypher, doctor, and core --clean-generated passed after the docs commit, and final clean status showed no dirty files across all four repos. The next safest rung is v0.3.37 release-note and asset-plan prep using the old v0.3.20 release as presentation template, not a release cut yet.

---

## 10. Preserved v0.3.35 Operational Process Log

The following section preserves the previous v0.3.35 LogDoc body as operational continuity. It is intentionally retained rather than compressed away. Where preserved text conflicts with the v0.3.36 current-state overlay above, v0.3.36 wins for current state; the preserved body remains the provenance/process ledger for the full-profile release validation ladder work and earlier history.



# CarbonStack LogDoc v0.3.35

**Last updated:** 2026-06-03 local session  
**Owner/operator:** bitcrusher32  
**Assistant/model:** GPT-5.5 Thinking  
**Current phase:** Phase 2E-prep / v0.3.x **full-profile release validation ladder checkpoint**. `carbonstack` is now at `63ffabf docs: record full profile release validation ladder`, after the v0.3.35 rung changed the Go runner semantics so `full` is no longer a stale `core` alias. The `full` profile now means `release-snapshot` followed by `local-cypher`; `release-snapshot` already calls `core`, so `full` deliberately does **not** call `core` a second time. `carbonstack-cypher` remains at `9ab994c docs: point to local operator runbook`; `carbonstack-comms` remains at `012c8bf scripts: support source snapshot artifact guard`; `carbonstack-os` remains at `1bbbe52 docs: clarify CarbonStackOS target direction`. The live WSL Debian working umbrella at `~/repos/carbonstack_umbrella` passed `local-cypher`, `doctor`, and `core --clean-generated` after the full-profile release validation ladder docs commit. Final clean status showed no dirty files across the four repos.

**Version schema:** `v[scope].[timeline]`. This file is `v0.3.35`, the full-profile release validation ladder handoff after the prior pre-v0.4.0 release-surface cleanup checkpoint. Per LogDoc V2 practice, this Markdown file intentionally preserves verbose operational continuity; the JSON breakpoint remains the lean handoff state.

**Naming correction note:** The immediately preceding LogDoc file was provided as `CarbonStackLogDocV0.3.34.md`, but its internal body still says `# CarbonStack LogDoc v0.3.33` and describes the pre-v0.4.0 release-surface cleanup checkpoint. Treat that prior document as the **v0.3.34 continuity source with stale internal v0.3.33 naming artifacts**. This v0.3.35 overlay is the current authoritative state. Any older internal heading/version text preserved below remains provenance, not current naming truth.

**LogDoc V2 usage note:** LogDoc V2 uses a split between a Markdown LogDoc for working-session continuity, a JSON state file for end-of-session handoff, Git for history preservation, and README/docs for public claims. This v0.3.35 Markdown file intentionally stays “bloated” where operational process, blunders, validation ladders, and historical context matter. The v0.3.35 JSON should remain lean and should agree with this file on the latest current state.

---

## 1. Project Goal

**Active goal:** Preserve CarbonStack state after converting the `full` validation profile from a stale `core` alias into the intended v0.4.0 release-package validation ladder. The immediate next line of work should be v0.3.36 package rehearsal / release-candidate packaging validation: create or stage a release-like package root, generate/verify checksums and release metadata, validate from a fresh extraction or throwaway package root, and only then prepare v0.4.0 release notes/assets if the evidence is clean.

CarbonStack’s broader project intent remains:

- **CarbonStack** is the secure-communications appliance-stack project and shared doctrine/release authority.
- **CarbonStackOS** is the future constrained Android-derived appliance OS.
- **CarbonStackComms** is the text-first encrypted messaging client and current OpenMLS sidecar/relay integration surface.
- **CarbonStackCypher** is the self-hostable relay/storage stack.
- The `carbonstack` repo remains the doctrine/spec/docs/release authority; implementation repos must not contradict it without revising the spec.

**North star:**

> Its goal is not to be a better smartphone. Its goal is to stop being a smartphone.

**Core doctrine:**

> Every feature is guilty until it proves it does not add unacceptable parser, network, sensor, identity, or filesystem authority.

**Current reality after v0.3.35:** `v0.3.20` remains the current public runner-backed testing release. `v0.3.21` through `v0.3.33/v0.3.34` built the local-only deployability runway, implemented `local-cypher`, added explicit cleanup, added one negative protocol validation, and cleaned the pre-v0.4.0 public release surface. v0.3.35 now makes `full` the release-package validation ladder:

    full = release-snapshot + local-cypher

Important meaning:

    release-snapshot performs release package/layout checks, strict pre-test artifact scan, checksum verification, and then calls core.
    local-cypher then performs the Cypher-only local lifecycle validation.
    core is not called twice.
    --clean-generated can be used with full to remove known generated roots after the whole successful profile.

Recommended future package-validation command shape:

    go run . --profile full --root /path/to/release-package-root --clean-generated

This is still **not** a production deployability release. It does **not** prove production readiness, production E2EE, hostile-server safety, metadata privacy, Android readiness, secure vault/storage, Debian production deployability, systemd readiness, cloudflared readiness, public ingress safety, real homelab validation, external audit, or certification.

---

## 2. Current Repo Heads

    carbonstack        63ffabf (HEAD -> main, origin/main, origin/HEAD) docs: record full profile release validation ladder
    carbonstack-comms  012c8bf (HEAD -> main, origin/main, origin/HEAD) scripts: support source snapshot artifact guard
    carbonstack-cypher 9ab994c (HEAD -> main, origin/main, origin/HEAD) docs: point to local operator runbook
    carbonstack-os     1bbbe52 (HEAD -> main, origin/main, origin/HEAD) docs: clarify CarbonStackOS target direction

Current `carbonstack` public runner-backed testing release tag:

    v0.3.20 -> 3eeb1a1e1a / 3eeb1a1 docs: clean public release surface

Previous mainline deployability checkpoints:

    v0.3.35 mainline -> 63ffabf docs: record full profile release validation ladder
    v0.3.34/v0.3.33 naming-correction source -> d83d2ef docs: prepare pre-v0.4.0 release surface
    v0.3.32 mainline -> cff3ab4 docs: record local-cypher negative protocol validation
    v0.3.31 mainline -> 7894278 docs: record local-cypher polish and cleanup control
    v0.3.30 mainline -> 15a3758 docs: record local-cypher runner implementation
    v0.3.30 runner  -> a98ee9f runner: add local-cypher
    v0.3.29 mainline -> 37a6d2a docs: record local-cypher validation contract
    v0.3.28 mainline -> af33139 docs: record local operator helper runner decision recon
    v0.3.27 mainline -> e516fc7 docs: record local Cypher API lifecycle proof
    v0.3.26 mainline -> 6e14fe0 docs: record local operator config data convention
    v0.3.25 mainline -> a5c6351 docs: update readme, roadmap, and historical doc
    v0.3.25 cypher  -> 9ab994c docs: point to local operator runbook
    v0.3.24 mainline -> 76f0258 docs: record Cypher schema migrations implementation
    v0.3.24 cypher  -> 42f838f db: track applied migrations

Important continuity note: v0.3.35 does not supersede v0.3.20 as a public release. It advances mainline release readiness by making the runner’s `full` profile match the emerging v0.4.0 release-package validation model. It is safe to pause here because the runner semantics change and docs/result record were pushed, WSL Debian `local-cypher` / `doctor` / `core --clean-generated` passed after the docs commit, and final status was clean across all four repos.

---

## 3. Current Validated State

Validated at the v0.3.35 checkpoint:

    [WSL-DEBIAN] working umbrella exists at ~/repos/carbonstack_umbrella
    [WSL-DEBIAN] carbonstack is at 63ffabf
    [WSL-DEBIAN] carbonstack-comms is at 012c8bf
    [WSL-DEBIAN] carbonstack-cypher is at 9ab994c
    [WSL-DEBIAN] carbonstack-os is at 1bbbe52
    [CARBONSTACK] full profile release validation ladder semantics landed
    [CARBONSTACK] docs/153 full profile release validation ladder record committed and pushed
    [GO-RUNNER] local-cypher passed after the v0.3.35 docs commit
    [GO-RUNNER] doctor passed after the v0.3.35 docs commit
    [GO-RUNNER] core --clean-generated passed after the v0.3.35 docs commit
    [ARTIFACT CLEANUP] known OpenMLS generated roots were removed by the runner's explicit --clean-generated flag
    [FINAL STATUS] all four repos were clean after validation without a separate manual cleanup step

Observed WSL Debian toolchain baseline:

    git version 2.47.3
    go version go1.24.4 linux/amd64
    rustc 1.96.0 (ac68faa20 2026-05-25)
    cargo 1.96.0 (30a34c682 2026-05-25)
    sqlite3 3.46.1

Observed `local-cypher` validation result during the v0.3.35 final flow:

    required paths: PASS
    pre-local-cypher artifact scan: PASS / no generated/private/build artifact hits
    temporary Cypher binary build: PASS
    first health check: PASS
    seeded Alice invite claim: PASS
    Bob dev invite creation: PASS
    Bob invite claim: PASS
    Alice device registration: PASS
    Bob device registration: PASS
    Alice device listing: PASS
    invalid stub-text/OpenMLS protocol pairing rejection: PASS
    opaque OpenMLS application-message envelope submit: PASS
    Bob inbox retrieval and payload metadata verification: PASS
    envelope ack: PASS
    Bob inbox after ack empty: PASS
    restart against same temporary DB: PASS
    restart health check: PASS
    persisted Alice device state after restart: PASS
    Bob's acked inbox remains empty after restart: PASS
    post-local-cypher artifact scan: PASS / no generated/private/build artifact hits
    VALIDATION PASSED

Observed `core --clean-generated` validation result during the v0.3.35 final flow:

    doctor nested inside core: PASS
    pre-test artifact scan: PASS / no generated/private/build artifact hits
    OpenMLS real-Cypher lifecycle: PASS
    carbonstack-comms package tests: PASS
    carbonstack-cypher package tests: PASS
    post-test artifact scan: expected OpenMLS sidecar generated roots only
    clean generated artifacts step: PASS
    removed .carbonstack-openmls-sidecar-state
    removed target
    VALIDATION PASSED

Expected/generated roots handled by `--clean-generated`:

    carbonstack-comms/internal/protocol/mls/openmls-sidecar/.carbonstack-openmls-sidecar-state
    carbonstack-comms/internal/protocol/mls/openmls-sidecar/target

---

## 4. v0.3.35 Work Completed

### 4.1 Release validation ladder recon happened before patching

Before patching, WSL Debian recon inspected:

    runner help/profile surface
    main.go profile dispatch and flag list
    release_snapshot.go
    checksums.go
    local_cypher.go key surface
    top-level README
    roadmap
    docs/README
    v0.3.20 release docs
    v0.3.32/v0.3.33 recent result docs
    release-snapshot/package claim scan
    live validation behavior

The recon confirmed:

    full still aliased core.
    release-snapshot already performs release layout checks, strict pre-test artifact scan, checksum verification, and then calls core.
    local-cypher exists as a separate Cypher-only local lifecycle profile.
    release-snapshot is intended for a fresh extracted or throwaway staged package root.
    release-snapshot run-order warnings remain important because core generates OpenMLS sidecar state and Rust build artifacts.
    the live working repo is not the same as a final release package root.
    v0.4.0 needs a release validation ladder, not more feature work.

### 4.2 `full` stopped being a stale `core` alias

Before v0.3.35, the runner README and dispatch said `full` currently aliases `core`.

That was acceptable earlier, but it became stale once the project had:

    release-snapshot
    local-cypher
    core
    --clean-generated
    pre-v0.4.0 release-surface cleanup

New behavior:

    full runs release-snapshot, then local-cypher.

Important detail:

    release-snapshot already calls core, so full does not call core a second time.

This makes `full` the intended future v0.4.0 release-package validation ladder without overloading `core` or inventing `local-backbone`.

### 4.3 `full` remains package-validation, not live-repo validation

`full` is now intended for fresh extracted or throwaway staged release package roots.

Recommended shape:

    go run . --profile full --root /path/to/release-package-root --clean-generated

It should not be treated as a casual live working-tree profile unless the working tree is intentionally staged like a release package with release metadata/checksums.

Known behavior:

    release-snapshot requires package metadata under release/
    release-snapshot verifies checksums
    release-snapshot runs strict pre-test artifact scan before tests
    release-snapshot calls core
    local-cypher runs afterward under full

### 4.4 README/docs/roadmap update landed

New/updated docs at v0.3.35:

    carbonstack/docs/153-full-profile-release-validation-ladder-v0.md
    carbonstack/docs/README.md
    carbonstack/roadmap/ROADMAP.md
    carbonstack/tools/carbonstack-validate/README.md

Final mainline docs commit:

    63ffabf docs: record full profile release validation ladder

The docs record:

    the recon finding that full was stale
    the reason not to use release-snapshot -> local-cypher -> core
    the correct sequence full = release-snapshot + local-cypher
    the recommended v0.4.0 validation command shape
    current scope boundaries
    next rung as release package rehearsal

### 4.5 Final validation passed with cleanup

Final WSL Debian validation sequence included:

    go run . --profile local-cypher
    go run . --profile doctor
    go run . --profile core --clean-generated

Final WSL Debian repo snapshot:

    carbonstack        63ffabf docs: record full profile release validation ladder
    carbonstack-comms  012c8bf scripts: support source snapshot artifact guard
    carbonstack-cypher 9ab994c docs: point to local operator runbook
    carbonstack-os     1bbbe52 docs: clarify CarbonStackOS target direction

Final validation passed under WSL Debian. `local-cypher` passed and still included the negative-path line. `doctor` passed. `core --clean-generated` passed OpenMLS real-Cypher lifecycle, Comms package tests, and Cypher package tests, then removed the expected OpenMLS generated roots through the explicit cleanup flag. Final clean status showed no dirty files across all four repos.

---

## 5. Critical Blunders / Lessons Preserved

### 5.1 The prior LogDoc naming offset must not propagate

The provided prior LogDoc was renamed to v0.3.34 by filename, but the internal document still said v0.3.33 and described the pre-v0.4.0 release-surface cleanup checkpoint.

Correct continuity handling:

    treat the prior file as the v0.3.34 continuity source with stale internal v0.3.33 naming artifacts.
    do not rewrite the preserved historical body except by adding this correction note.
    this v0.3.35 overlay is authoritative for current state.

### 5.2 User intuition about `full` was directionally correct

The suggested idea was that `full` should become a validation ladder instead of a `core` alias.

Corrected implementation:

    not release-snapshot -> local-cypher -> core
    yes release-snapshot -> local-cypher

Reason:

    release-snapshot already calls core.

Lesson:

    good intuition, but inspect runner call graph before composing validation profiles.
    avoid duplicating expensive/side-effecting validation steps.

### 5.3 `full` should not become `local-backbone`

The `full` profile now coordinates release-package validation surfaces. It is not a whole-stack runtime deployability proof.

Correct lesson:

    `full` is a release validation ladder.
    `local-backbone` remains reserved for later whole-stack validation after Comms runtime UX is actually wired through the backbone.

### 5.4 `release-snapshot` still requires package-root discipline

`release-snapshot` is not a generic live working-tree command.

Correct lesson:

    run it from a fresh extracted or throwaway staged package root.
    do not archive a package root after running release-snapshot inside it unless expected generated artifacts were cleaned and package rules still pass.
    the next rung must rehearse the actual package path, not only live repo validation.

### 5.5 Validation cadence remains healthy

v0.3.35 still validated:

    local-cypher
    doctor
    core --clean-generated

Correct lesson:

    this is the live umbrella validation ladder.
    `full` is now reserved for package validation once a package root exists.

---

## 6. Current Blockers / Not Validated

Current blockers and nonvalidated areas:

    v0.4.0 release package/checklist validation has not yet been rehearsed from a staged or fresh extracted package root.
    full has been wired as the release validation ladder, but full itself has not yet been proven against a real release package root in the final v0.4.0 package shape.
    No v0.4.0 release candidate package has been staged, checksum-verified, extracted, and validated.
    local-cypher implements positive Cypher-only lifecycle validation plus one negative protocol-pairing check, but no broad negative-path suite exists yet.
    local-backbone runner profile does not exist and remains reserved for later whole-stack validation.
    No helper command exists yet.
    Runtime Comms send/inbox remains stub-era and not wired as user-facing OpenMLS messenger UX.
    Production-safe provider storage and secure vault design are unsolved.
    Hostile-server rollback/replay/metadata-abuse validation remains future work.
    Android/CarbonStackOS implementation remains later work.
    CarbonStack Relay Space mechanics remain a defined future concept, not an implementation.
    Public ingress, systemd, cloudflared, real homelab deployment, and production service paths remain out of scope.
    Temporary DB validation remains dev/test-only and does not replace later persistent user/operator-state validation for public Comms maturity.

Hard nonclaims remain:

    Do not claim production readiness.
    Do not claim production E2EE.
    Do not claim hostile-server safety.
    Do not claim metadata privacy.
    Do not claim Debian production deployability.
    Do not claim systemd readiness.
    Do not claim cloudflared readiness.
    Do not claim Android readiness.
    Do not claim secure vault/storage.
    Do not claim external audit or certification.
    Do not claim CarbonStack Relay Space mechanics are implemented.
    Do not claim local-backbone exists or is justified yet.
    Do not claim local-cypher validates runtime Comms UX.
    Do not claim full validates local-backbone.
    Do not claim full is a deployment command.
    Do not claim --clean-generated is a general cleanup tool.
    Do not claim v0.3.35 is itself the v0.4.0 release.

Updated allowed claim:

    CarbonStack v0.3.35 makes the Go runner's full profile the release-package validation ladder.
    full now runs release-snapshot followed by local-cypher.
    release-snapshot already calls core, so full does not call core twice.
    The recommended future release-package validation command shape is `go run . --profile full --root /path/to/release-package-root --clean-generated`.
    local-cypher, doctor, and core --clean-generated passed under WSL Debian after the v0.3.35 docs commit.
    Final clean status showed no dirty files across the four repos.

---

## 7. Next Safest Actions

Recommended next sequence after this v0.3.35 breakpoint:

    v0.3.36 preferred:
      release package rehearsal / fresh package-root validation

Recommended concrete next rung:

    1. Recon existing package-generation/release asset scripts or absence thereof.
    2. Decide the v0.4.0 candidate package shape.
    3. Create a throwaway staged release package root.
    4. Ensure release metadata exists:
       release/manifest.json
       release/checksums.txt
       release/validation-freeze.md
       release/testing-runbook.md
       release/release-notes.md if ready
       release/LICENSE if needed
    5. Run write-checksums against the staged package source root where appropriate.
    6. Create a fresh extraction / throwaway validation root.
    7. Run:
       go run . --profile full --root <package-root> --clean-generated
    8. Document failures or success.
    9. If clean, move toward v0.4.0 release candidate wording and asset plan.
    10. Do not cut v0.4.0 until package rehearsal evidence is clean and release text is ready.

Avoid next:

    local-backbone runner profile
    helper implementation
    public ingress
    cloudflared
    systemd
    real homelab deployment
    runtime Comms send/inbox UX
    Android app
    CarbonStackOS
    CarbonStack Relay Space implementation
    production release
    broad negative-path suite unless release validation exposes a direct blocker

---

## 8. Critical Paths / Functions

### WSL Debian active working umbrella

    [WSL-DEBIAN] umbrella: ~/repos/carbonstack_umbrella
    [WSL-DEBIAN] carbonstack: ~/repos/carbonstack_umbrella/carbonstack
    [WSL-DEBIAN] carbonstack-comms: ~/repos/carbonstack_umbrella/carbonstack-comms
    [WSL-DEBIAN] carbonstack-cypher: ~/repos/carbonstack_umbrella/carbonstack-cypher
    [WSL-DEBIAN] carbonstack-os: ~/repos/carbonstack_umbrella/carbonstack-os
    [GO-RUNNER] runner dir: ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate

### WSL Debian known-good validation commands

Live umbrella validation:

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go run . --profile local-cypher
    go run . --profile doctor
    go run . --profile core --clean-generated

Future package-root validation:

    cd <package-root>/carbonstack/tools/carbonstack-validate
    go run . --profile full --root <package-root> --clean-generated

### v0.3.35 docs

    [DOC] v0.3.35 full profile release validation ladder:
    carbonstack/docs/153-full-profile-release-validation-ladder-v0.md

    [DOC] v0.3.34/v0.3.33 naming-correction source / pre-v0.4.0 release-surface cleanup:
    carbonstack/docs/152-pre-v0.4.0-release-surface-cleanup-v0.md

    [DOC] v0.3.32 negative protocol validation result:
    carbonstack/docs/151-local-cypher-negative-protocol-validation-v0.md

    [DOC] docs index:
    carbonstack/docs/README.md

    [DOC] roadmap:
    carbonstack/roadmap/ROADMAP.md

    [DOC] top-level README:
    carbonstack/README.md

### Go runner surfaces after v0.3.35

    [RUNNER] directory:
    carbonstack/tools/carbonstack-validate

    [RUNNER] implementation files:
    carbonstack/tools/carbonstack-validate/main.go
    carbonstack/tools/carbonstack-validate/local_cypher.go
    carbonstack/tools/carbonstack-validate/checksums.go
    carbonstack/tools/carbonstack-validate/release_snapshot.go
    carbonstack/tools/carbonstack-validate/README.md
    carbonstack/tools/carbonstack-validate/go.mod

    [RUNNER] profiles:
    doctor
    core
    local-cypher
    full
    release-snapshot
    write-checksums
    verify-checksums

    [RUNNER] cleanup flag:
    --clean-generated

### full profile behavior after v0.3.35

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

### --clean-generated behavior

    Runs only after successful profile execution.
    Deletes only:
      carbonstack-comms/internal/protocol/mls/openmls-sidecar/.carbonstack-openmls-sidecar-state
      carbonstack-comms/internal/protocol/mls/openmls-sidecar/target

    Does not delete:
      arbitrary untracked files
      manual local operator DBs
      $HOME/.local/share/carbonstack/cypher/cypher.db
      provider-storage.json unless added by a future deliberate expansion
      signer.json unless added by a future deliberate expansion

    Does not replace:
      ArtifactScan
      StrictPreTestArtifactScan
      release-snapshot pre-test hygiene checks

### Current mainline public-surface state

    Top-level README:
      v0.3.20 remains current public testing release.
      v0.3.32/v0.3.35 is current mainline validation state.
      v0.4.0 is framed as a broad local deployability pre-release.

    Roadmap:
      late v0.3.x now points toward release package validation/rehearsal and then v0.4.0 release candidate work.

    docs/README:
      current release/validation docs list now includes docs/153.

### Accepted envelope pair

    content_type=carbonstack.mls.application-message.v0
    protocol_version=carbonstack-openmls-sidecar-v0

### Negative protocol pair now validated

    content_type=carbonstack.message.text.stub.v0
    protocol_version=carbonstack-openmls-sidecar-v0

Expected result:

    HTTP 400
    error code unsupported_protocol_version

---

## 9. Lean Breakpoint Summary

    CarbonStack v0.3.35 is a stable mainline full-profile release validation ladder breakpoint after the pre-v0.4.0 release-surface cleanup checkpoint. It corrects the runner semantics so full is no longer a stale core alias. full now runs release-snapshot followed by local-cypher; release-snapshot already calls core, so full does not duplicate core. The recommended future v0.4.0 package-validation command is go run . --profile full --root <package-root> --clean-generated from a fresh extracted or throwaway staged package root. local-cypher, doctor, and core --clean-generated passed under WSL Debian after the docs commit, and final clean status showed no dirty files across all four repos. The next safest rung is release package rehearsal / fresh package-root validation, not new feature work or a release cut.

---

## 10. Preserved v0.3.34/v0.3.33 Operational Process Log

The following section preserves the previous LogDoc body as operational continuity. It is intentionally retained rather than compressed away. The prior file was named `CarbonStackLogDocV0.3.34.md` by the user but still contains internal `v0.3.33` headings and text. Treat that as a naming artifact. Where preserved text conflicts with the v0.3.35 current-state overlay above, v0.3.35 wins for current state; the preserved body remains the provenance/process ledger for the pre-v0.4.0 release-surface cleanup work and earlier history.



# CarbonStack LogDoc v0.3.33

**Last updated:** 2026-06-03 local session  
**Owner/operator:** bitcrusher32  
**Assistant/model:** GPT-5.5 Thinking  
**Current phase:** Phase 2E-prep / v0.3.x **pre-v0.4.0 release-surface cleanup checkpoint**. `carbonstack` is now at `d83d2ef docs: prepare pre-v0.4.0 release surface`, after v0.3.33 refreshed the top-level README, roadmap, docs index, and added `docs/152-pre-v0.4.0-release-surface-cleanup-v0.md`. `carbonstack-cypher` remains at `9ab994c docs: point to local operator runbook`; `carbonstack-comms` remains at `012c8bf scripts: support source snapshot artifact guard`; `carbonstack-os` remains at `1bbbe52 docs: clarify CarbonStackOS target direction`. The live WSL Debian working umbrella at `~/repos/carbonstack_umbrella` passed `local-cypher`, `doctor`, and `core --clean-generated` after the v0.3.33 docs/public-surface cleanup commit. Final clean status showed no dirty files across the four repos.

**Version schema:** `v[scope].[timeline]`. This file is `v0.3.33`, the pre-v0.4.0 release-surface cleanup handoff after the `v0.3.32` local-cypher negative protocol validation checkpoint. Per LogDoc V2 practice, this Markdown file intentionally preserves verbose operational continuity; the JSON breakpoint remains the lean handoff state. The dense v0.3.32, v0.3.31, v0.3.30, v0.3.29, v0.3.28, v0.3.27, v0.3.26, v0.3.25, v0.3.24, v0.3.23, v0.3.22, v0.3.21, and v0.3.20 operational process logs are preserved later in this file rather than compressed away.

**LogDoc V2 usage note:** LogDoc V2 uses a split between a Markdown LogDoc for working-session continuity, a JSON state file for end-of-session handoff, Git for history preservation, and README/docs for public claims. This v0.3.33 Markdown file therefore intentionally stays “bloated” where operational process, blunders, validation ladders, and historical context matter. The v0.3.33 JSON should remain lean and should agree with this file on the latest current state.

---

## 1. Project Goal

**Active goal:** Preserve CarbonStack state after cleaning the public/release-facing documentation surface for the pre-v0.4.0 runway. v0.3.33 did not add new validation features; it aligned the top-level README, roadmap, docs index, and a new result doc around the current reality: v0.3.20 remains the current public runner-backed testing release, while mainline has advanced to v0.3.32/v0.3.33 with WSL Debian `local-cypher`, `doctor`, and `core --clean-generated` validation. The immediate next line of work should be v0.3.34 release checklist and package validation rehearsal, not more feature work unless the rehearsal exposes a release-blocking issue.

CarbonStack’s broader project intent remains:

- **CarbonStack** is the secure-communications appliance-stack project and shared doctrine/release authority.
- **CarbonStackOS** is the future constrained Android-derived appliance OS.
- **CarbonStackComms** is the text-first encrypted messaging client and current OpenMLS sidecar/relay integration surface.
- **CarbonStackCypher** is the self-hostable relay/storage stack.
- The `carbonstack` repo remains the doctrine/spec/docs/release authority; implementation repos must not contradict it without revising the spec.

**North star:**

> Its goal is not to be a better smartphone. Its goal is to stop being a smartphone.

**Core doctrine:**

> Every feature is guilty until it proves it does not add unacceptable parser, network, sensor, identity, or filesystem authority.

**Current reality after v0.3.33:** `v0.3.20` remains the current public runner-backed testing release. `v0.3.21` defined the local-only deployability model. `v0.3.22` narrowed the local deployability surface around Cypher operator config/API/database assumptions and adopted **CarbonStack Relay Space** terminology. `v0.3.23` proved the SQLite repeated-migration persistence hazard. `v0.3.24` implemented `schema_migrations` tracking in CarbonStackCypher. `v0.3.25` recorded the first local-only Cypher operator runbook skeleton. `v0.3.26` recorded the local operator config/data convention. `v0.3.27` proved a real local-only Cypher explicit-env API lifecycle. `v0.3.28` recorded the helper/runner decision recon and reserved `local-backbone` for later whole-stack validation. `v0.3.29` defined the future `local-cypher` validation contract. `v0.3.30` implemented that contract as the first `local-cypher` profile in the Go runner. `v0.3.31` polished `local-cypher` output and added explicit opt-in `--clean-generated`. `v0.3.32` added the first negative-path validation to `local-cypher`. `v0.3.33` cleaned the public surface so the repo no longer leads with duplicated/stale v0.3.20-era current-state wording and now clearly frames the v0.4.0 target as a broad local deployability pre-release / research-and-development milestone.

The current public release page remains:

    https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.3.20

The current public testing interpretation remains:

    v0.3.20 is the preferred public runner-backed testing release.
    v0.3.4 remains available as the older Windows 11 / PowerShell legacy testbed.
    Debian / WSL Debian is now the mainline public dev/test validation direction.
    Windows 11 short-path validation was the final explicit Windows dev/test validation for the v0.3.20 phase.

v0.4.0 release framing direction is now surfaced in the top-level public docs:

    v0.4.0 should be framed as a broad local deployability pre-release.
    It should be positioned as a milestone / research-and-development release.
    It should clearly state that it is not intended for public-user use.
    It should clearly state that it is not intended for application use.
    It should clearly state that it is not production secure.
    It should clearly state that it is not hostile-server-certified.
    Its concrete validated artifact should remain the WSL Debian runner-backed validation surface, including local-cypher/core.

This is still **not** a production deployability release. It does **not** prove production readiness, production E2EE, hostile-server safety, metadata privacy, Android readiness, secure vault/storage, Debian production deployability, systemd readiness, cloudflared readiness, public ingress safety, real homelab validation, external audit, or certification.

---

## 2. Current Repo Heads

    carbonstack        d83d2ef (HEAD -> main, origin/main, origin/HEAD) docs: prepare pre-v0.4.0 release surface
    carbonstack-comms  012c8bf (HEAD -> main, origin/main, origin/HEAD) scripts: support source snapshot artifact guard
    carbonstack-cypher 9ab994c (HEAD -> main, origin/main, origin/HEAD) docs: point to local operator runbook
    carbonstack-os     1bbbe52 (HEAD -> main, origin/main, origin/HEAD) docs: clarify CarbonStackOS target direction

Current `carbonstack` public runner-backed testing release tag:

    v0.3.20 -> 3eeb1a1e1a / 3eeb1a1 docs: clean public release surface

Previous mainline deployability checkpoints:

    v0.3.33 mainline -> d83d2ef docs: prepare pre-v0.4.0 release surface
    v0.3.32 mainline -> cff3ab4 docs: record local-cypher negative protocol validation
    v0.3.31 mainline -> 7894278 docs: record local-cypher polish and cleanup control
    v0.3.30 mainline -> 15a3758 docs: record local-cypher runner implementation
    v0.3.30 runner  -> a98ee9f runner: add local-cypher
    v0.3.29 mainline -> 37a6d2a docs: record local-cypher validation contract
    v0.3.28 mainline -> af33139 docs: record local operator helper runner decision recon
    v0.3.27 mainline -> e516fc7 docs: record local Cypher API lifecycle proof
    v0.3.26 mainline -> 6e14fe0 docs: record local operator config data convention
    v0.3.25 mainline -> a5c6351 docs: update readme, roadmap, and historical doc
    v0.3.25 cypher  -> 9ab994c docs: point to local operator runbook
    v0.3.24 mainline -> 76f0258 docs: record Cypher schema migrations implementation
    v0.3.24 cypher  -> 42f838f db: track applied migrations

Legacy `carbonstack` Windows/PowerShell testbed release tag:

    v0.3.4 -> 7c5b7e6 docs: record release snapshot self-test fix

Historical `carbonstack` release tag:

    v0.3.0 -> 59245fc docs: freeze v0.3.0 release packaging plan

Important continuity note: `v0.3.33` does not supersede `v0.3.20` as a public release. It advances mainline release readiness by cleaning public-surface wording and establishing a current v0.4.0 runway. It is safe to pause here because the docs cleanup was pushed, WSL Debian `local-cypher` / `doctor` / `core --clean-generated` passed after the docs commit, and final status was clean across all four repos.

---

## 3. Current Validated State

Validated at the v0.3.33 checkpoint:

    [WSL-DEBIAN] working umbrella exists at ~/repos/carbonstack_umbrella
    [WSL-DEBIAN] carbonstack is at d83d2ef
    [WSL-DEBIAN] carbonstack-comms is at 012c8bf
    [WSL-DEBIAN] carbonstack-cypher is at 9ab994c
    [WSL-DEBIAN] carbonstack-os is at 1bbbe52
    [CARBONSTACK] README.md was rewritten/cleaned for current public testing and current mainline validation status
    [CARBONSTACK] roadmap/ROADMAP.md was rewritten/cleaned for the late-v0.3.x -> v0.4.0 release runway
    [CARBONSTACK] docs/README.md was rewritten/cleaned to surface current public testing status and v0.3.32/v0.3.33 mainline validation state
    [CARBONSTACK] docs/152-pre-v0.4.0-release-surface-cleanup-v0.md exists and records the v0.3.33 release-surface cleanup
    [GO-RUNNER] local-cypher passed after the v0.3.33 docs commit
    [GO-RUNNER] doctor passed after the v0.3.33 docs commit
    [GO-RUNNER] core --clean-generated passed after the v0.3.33 docs commit
    [ARTIFACT CLEANUP] known OpenMLS generated roots were removed by the runner's explicit --clean-generated flag
    [FINAL STATUS] all four repos were clean after validation without a separate manual cleanup step

Observed WSL Debian toolchain baseline:

    git version 2.47.3
    go version go1.24.4 linux/amd64
    rustc 1.96.0 (ac68faa20 2026-05-25)
    cargo 1.96.0 (30a34c682 2026-05-25)
    sqlite3 3.46.1

Observed `local-cypher` validation result during the v0.3.33 final flow:

    required paths: PASS
    pre-local-cypher artifact scan: PASS / no generated/private/build artifact hits
    temporary Cypher binary build: PASS
    first health check: PASS
    seeded Alice invite claim: PASS
    Bob dev invite creation: PASS
    Bob invite claim: PASS
    Alice device registration: PASS
    Bob device registration: PASS
    Alice device listing: PASS
    invalid stub-text/OpenMLS protocol pairing rejection: PASS
    opaque OpenMLS application-message envelope submit: PASS
    Bob inbox retrieval and payload metadata verification: PASS
    envelope ack: PASS
    Bob inbox after ack empty: PASS
    first Cypher process stopped with expected SIGINT termination
    restart against same temporary DB: PASS
    restart health check: PASS
    persisted Alice device state after restart: PASS
    Bob's acked inbox remains empty after restart: PASS
    second Cypher process stopped with expected SIGINT termination
    post-local-cypher artifact scan: PASS / no generated/private/build artifact hits
    VALIDATION PASSED

Observed `core --clean-generated` validation result during the v0.3.33 final flow:

    doctor nested inside core: PASS
    pre-test artifact scan: PASS / no generated/private/build artifact hits
    OpenMLS real-Cypher lifecycle: PASS
    carbonstack-comms package tests: PASS
    carbonstack-cypher package tests: PASS
    post-test artifact scan: expected OpenMLS sidecar generated roots only
    clean generated artifacts step: PASS
    removed .carbonstack-openmls-sidecar-state
    removed target
    VALIDATION PASSED

Expected/generated roots handled by `--clean-generated`:

    carbonstack-comms/internal/protocol/mls/openmls-sidecar/.carbonstack-openmls-sidecar-state
    carbonstack-comms/internal/protocol/mls/openmls-sidecar/target

Still absent unless created by future behavior:

    carbonstack-comms/internal/protocol/mls/openmls-sidecar/provider-storage.json
    carbonstack-comms/internal/protocol/mls/openmls-sidecar/signer.json

---

## 4. v0.3.33 Work Completed

### 4.1 Release-surface recon happened before cleanup

Before patching, WSL Debian recon inspected top-level `README.md`, `roadmap/ROADMAP.md`, `docs/README.md`, recent docs, release/validation doc inventory, release-snapshot/checksum surfaces, runner help, and validation behavior.

The recon found:

    top-level README had duplicated "Current testing path" sections.
    README still framed near-term v0.3.x using stale v0.3.20-era language.
    roadmap still carried older v0.2.x/v0.3.0 scaffolding before newer current-state material.
    docs/README still pointed to v0.3.20 as current release direction but did not surface v0.3.32/local-cypher as the current mainline validation state.
    validation itself remained healthy.
    local-cypher passed.
    doctor passed.
    core --clean-generated passed.
    final repo status was clean before cleanup patching.

This preserved the project method:

    recon -> targeted public-surface cleanup -> validation -> breakpoint

### 4.2 Python path bug happened and was fixed

The first v0.3.33 cleanup script rewrote the intended files but crashed while generating the result doc because it computed sibling repo paths incorrectly.

Observed error:

    FileNotFoundError: [Errno 2] No such file or directory: 'carbonstack-comms'

Root cause:

    root = Path(".")
    repo_head(root.parent / "carbonstack-comms")

Because the script ran from `~/repos/carbonstack_umbrella/carbonstack`, the relative `root.parent` was not the actual umbrella path. Python looked for a sibling repo in the wrong place.

Fix:

    use Path.cwd().resolve() for carbonstack
    use carbonstack.parent as the real umbrella path
    derive comms/cypher/osrepo from that resolved umbrella path

Lesson preserved:

    when scripts run inside repo roots but need sibling repos, resolve the current path first.
    do not rely on relative Path(".").parent for sibling repo discovery.
    public-surface docs were rewritten before the crash, so the fix safely rewrote them again and added the missing result doc.

### 4.3 Top-level README cleanup landed

`README.md` was rewritten to:

    keep v0.3.20 as current public testing release
    surface v0.3.32 as current mainline validation state
    describe local-cypher/doctor/core --clean-generated as the current mainline validation commands
    remove duplicated Current testing path sections
    clarify that local-cypher is Cypher-only
    list what is currently proven
    list what is not proven
    define v0.4.0 as broad local deployability pre-release
    add CarbonStack Relay Space terminology guidance

### 4.4 Roadmap cleanup landed

`roadmap/ROADMAP.md` was rewritten to:

    remove stale v0.2.x/v0.3.0 current-target scaffolding from the top-level current roadmap
    define current state after v0.3.32
    define v0.3.33 as pre-v0.4.0 release-surface cleanup
    define v0.3.34 as release checklist and package validation rehearsal
    define v0.3.35 as release candidate wording and asset plan
    define v0.3.36+ as the possible v0.4.0 release cut if clean
    expand v0.4.x, v0.5.x, v0.6.x, v0.7.x, v0.8.x, v0.9.x, and later Android/CarbonStackOS direction
    preserve the governing principle that every major claim must trail validation

### 4.5 Docs index cleanup landed

`docs/README.md` was rewritten to:

    keep the docs archive model
    point users to top-level README, roadmap, and v0.3.20 release docs for public status
    surface v0.3.32 local-cypher negative protocol validation as current mainline state
    list current validation commands
    describe v0.4.0 release direction
    list the current release/validation docs from v0.3.20 through v0.3.33
    preserve strict historical-doc warning and security/maturity warning

### 4.6 New v0.3.33 result doc landed

New result doc:

    carbonstack/docs/152-pre-v0.4.0-release-surface-cleanup-v0.md

Docs commit:

    d83d2ef docs: prepare pre-v0.4.0 release surface

The doc records:

    purpose of v0.3.33
    repo heads before docs commit
    recon findings
    cleanup performed
    v0.4.0 framing
    current concrete validation surface
    current nonclaims
    recommended next rung: v0.3.34 release checklist and package validation rehearsal

### 4.7 Final validation passed with cleanup

Final WSL Debian validation sequence included:

    go run . --profile local-cypher
    go run . --profile doctor
    go run . --profile core --clean-generated

Final WSL Debian repo snapshot:

    carbonstack        d83d2ef docs: prepare pre-v0.4.0 release surface
    carbonstack-comms  012c8bf scripts: support source snapshot artifact guard
    carbonstack-cypher 9ab994c docs: point to local operator runbook
    carbonstack-os     1bbbe52 docs: clarify CarbonStackOS target direction

Final validation passed under WSL Debian. `local-cypher` passed and still included the negative-path line. `doctor` passed. `core --clean-generated` passed OpenMLS real-Cypher lifecycle, Comms package tests, and Cypher package tests, then removed the expected OpenMLS generated roots through the explicit cleanup flag. Final clean status showed no dirty files across all four repos.

---

## 5. Critical Blunders / Lessons Preserved

### 5.1 Public-surface drift was real

The recon showed the repo’s public front door had drifted behind mainline:

    README duplicated Current testing path sections.
    README still implied older v0.3.20-era near-term work.
    roadmap retained old v0.2.x/v0.3.0 scaffolding at the top.
    docs/README did not surface the current v0.3.32 mainline validation state.

Correct lesson:

    release readiness is not only tests passing.
    public wording can become stale enough to mislead users even when implementation/validation is healthy.
    v0.4.0 release prep must keep public claims aligned with current validation evidence.

### 5.2 Python sibling-path bug happened during docs generation

Observed error:

    FileNotFoundError: [Errno 2] No such file or directory: 'carbonstack-comms'

Root cause:

    the script tried to access sibling repos via a relative parent path while running inside the carbonstack repo.

Correct lesson:

    use resolved absolute paths for repo-umbrella scripts.
    when generating docs that mention sibling repo heads, derive the umbrella from `Path.cwd().resolve().parent`.
    if a script rewrites files before failing, rerun a fixed idempotent script rather than manually patching fragments.

### 5.3 v0.4.0 release wording is now a formal project surface

v0.3.33 moved v0.4.0 framing out of chat-only context and into README/roadmap/docs.

Correct lesson:

    release names are claims.
    v0.4.0 should be broad enough to describe local deployability milestone work, but cautious enough to avoid implying public user readiness.

### 5.4 No extra negative paths were added

The project deliberately stopped at the single historically grounded negative path before v0.4.0 prep.

Correct lesson:

    the current release-prep risk is public-surface accuracy, not lack of another negative-path test.
    broader negative-path suites belong later unless release validation exposes a direct blocker.

### 5.5 The validation cadence remains healthy

v0.3.33 validated:

    local-cypher
    doctor
    core --clean-generated

Correct lesson:

    this is now the standard late-v0.3.x WSL Debian live umbrella validation ladder.
    `core --clean-generated` avoids manual cleanup and keeps final status clean.

---

## 6. Current Blockers / Not Validated

Current blockers and nonvalidated areas:

    v0.4.0 release package/checklist validation has not been rehearsed yet.
    A v0.4.0 release candidate has not been staged, packaged, checksum-verified, downloaded/extracted, or validated.
    local-cypher implements positive Cypher-only lifecycle validation plus one negative protocol-pairing check, but no broad negative-path suite exists yet.
    local-backbone runner profile does not exist and remains reserved for later whole-stack validation.
    No helper command exists yet.
    Local-only operator convention is documented and proof-tested, but not enforced by code outside local-cypher.
    Cypher default bind remains :8080; local-only docs recommend explicit 127.0.0.1:8080.
    Runtime Comms send/inbox remains stub-era and not wired as user-facing OpenMLS messenger UX.
    Production-safe provider storage and secure vault design are unsolved.
    Hostile-server rollback/replay/metadata-abuse validation remains future work.
    Android/CarbonStackOS implementation remains later work.
    CarbonStack Relay Space mechanics remain a defined future concept, not an implementation.
    Public ingress, systemd, cloudflared, real homelab deployment, and production service paths remain out of scope.
    Temporary DB validation remains dev/test-only and does not replace later persistent user/operator-state validation for public Comms maturity.
    --clean-generated is intentionally narrow and does not replace release-snapshot strict pre-test validation.

Hard nonclaims remain:

    Do not claim production readiness.
    Do not claim production E2EE.
    Do not claim hostile-server safety.
    Do not claim metadata privacy.
    Do not claim Debian production deployability.
    Do not claim systemd readiness.
    Do not claim cloudflared readiness.
    Do not claim Android readiness.
    Do not claim secure vault/storage.
    Do not claim external audit or certification.
    Do not claim CarbonStack Relay Space mechanics are implemented.
    Do not claim local-backbone exists or is justified yet.
    Do not claim local-cypher validates runtime Comms UX.
    Do not claim local-cypher is sufficient for public user testing maturity.
    Do not claim --clean-generated is a general cleanup tool.
    Do not claim v0.3.33 is itself the v0.4.0 release.

Updated allowed claim:

    CarbonStack v0.3.33 prepares the public documentation surface for the pre-v0.4.0 release runway.
    v0.3.20 remains the current public runner-backed testing release.
    Current mainline validation state is v0.3.32/v0.3.33 with local-cypher, doctor, and core --clean-generated under WSL Debian.
    v0.4.0 is framed as a broad local deployability pre-release / research-and-development milestone, not a public-user-ready or production-security release.
    local-cypher, doctor, and core --clean-generated passed under WSL Debian after the v0.3.33 docs commit.
    Final clean status showed no dirty files across the four repos.

---

## 7. Next Safest Actions

Recommended next sequence after this v0.3.33 breakpoint:

    v0.3.34 preferred:
      release checklist and package validation rehearsal

Recommended concrete next rung:

    1. Recon release-snapshot/checksum/package generation state.
    2. Decide the v0.4.0 candidate package shape.
    3. Generate or rehearse a release-like package locally.
    4. Validate checksums.
    5. Validate release-snapshot from a clean/fresh extraction context.
    6. Validate local-cypher, doctor, and core --clean-generated from the relevant runner surface.
    7. Document any release-surface or package-shape gaps.
    8. Do not cut v0.4.0 yet unless the package rehearsal is clean enough and release notes are ready.

Likely late-v0.3.x path:

    v0.3.34:
      release checklist and package validation rehearsal

    v0.3.35:
      v0.4.0 release candidate wording and asset plan

    v0.3.36+:
      release candidate package validation / release cut if clean

Avoid next:

    local-backbone runner profile
    helper implementation
    public ingress
    cloudflared
    systemd
    real homelab deployment
    runtime Comms send/inbox UX
    Android app
    CarbonStackOS
    CarbonStack Relay Space implementation
    production release
    broad negative-path suite unless release validation exposes a direct blocker

---

## 8. Critical Paths / Functions

### WSL Debian active working umbrella

    [WSL-DEBIAN] umbrella: ~/repos/carbonstack_umbrella
    [WSL-DEBIAN] carbonstack: ~/repos/carbonstack_umbrella/carbonstack
    [WSL-DEBIAN] carbonstack-comms: ~/repos/carbonstack_umbrella/carbonstack-comms
    [WSL-DEBIAN] carbonstack-cypher: ~/repos/carbonstack_umbrella/carbonstack-cypher
    [WSL-DEBIAN] carbonstack-os: ~/repos/carbonstack_umbrella/carbonstack-os
    [GO-RUNNER] runner dir: ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate

### WSL Debian known-good validation commands

From runner dir:

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go run . --profile local-cypher
    go run . --profile doctor
    go run . --profile core --clean-generated

From carbonstack repo root:

    cd ~/repos/carbonstack_umbrella/carbonstack
    go run ./tools/carbonstack-validate --profile local-cypher
    go run ./tools/carbonstack-validate --profile doctor
    go run ./tools/carbonstack-validate --profile core --clean-generated

### v0.3.33 docs

    [DOC] v0.3.33 pre-v0.4.0 release-surface cleanup:
    carbonstack/docs/152-pre-v0.4.0-release-surface-cleanup-v0.md

    [DOC] v0.3.32 negative protocol validation result:
    carbonstack/docs/151-local-cypher-negative-protocol-validation-v0.md

    [DOC] v0.3.31 polish/cleanup result:
    carbonstack/docs/150-local-cypher-polish-generated-cleanup-v0.md

    [DOC] v0.3.30 implementation result:
    carbonstack/docs/149-local-cypher-runner-implementation-v0.md

    [DOC] v0.3.29 contract:
    carbonstack/docs/148-local-cypher-validation-contract-v0.md

    [DOC] docs index:
    carbonstack/docs/README.md

    [DOC] roadmap:
    carbonstack/roadmap/ROADMAP.md

    [DOC] top-level README:
    carbonstack/README.md

### Go runner surfaces after v0.3.33

    [RUNNER] directory:
    carbonstack/tools/carbonstack-validate

    [RUNNER] implementation files:
    carbonstack/tools/carbonstack-validate/main.go
    carbonstack/tools/carbonstack-validate/local_cypher.go
    carbonstack/tools/carbonstack-validate/checksums.go
    carbonstack/tools/carbonstack-validate/release_snapshot.go
    carbonstack/tools/carbonstack-validate/README.md
    carbonstack/tools/carbonstack-validate/go.mod

    [RUNNER] profiles:
    doctor
    core
    local-cypher
    full
    release-snapshot
    write-checksums
    verify-checksums

    [RUNNER] cleanup flag:
    --clean-generated

### --clean-generated behavior

    Runs only after successful profile execution.
    Deletes only:
      carbonstack-comms/internal/protocol/mls/openmls-sidecar/.carbonstack-openmls-sidecar-state
      carbonstack-comms/internal/protocol/mls/openmls-sidecar/target

    Does not delete:
      arbitrary untracked files
      manual local operator DBs
      $HOME/.local/share/carbonstack/cypher/cypher.db
      provider-storage.json unless added by a future deliberate expansion
      signer.json unless added by a future deliberate expansion

    Does not replace:
      ArtifactScan
      StrictPreTestArtifactScan
      release-snapshot pre-test hygiene checks

### local-cypher lifecycle after v0.3.33

    start:
      build temporary Cypher binary
      create temporary DB directory
      choose dynamic loopback port
      set CYPHER_ADDR to 127.0.0.1:<dynamic_port>
      set CYPHER_DB to temporary SQLite DB path
      set CYPHER_MIGRATIONS to source-tree migrations
      set CYPHER_DEV_INVITE to known temporary dev invite
      start Cypher
      wait for GET /v0/health

    account/device:
      POST /v0/invites/claim
      POST /v0/dev/invites
      POST /v0/invites/claim
      POST /v0/devices/register
      GET /v0/accounts/<account_id>/devices

    negative envelope check:
      POST /v0/envelopes with:
        content_type=carbonstack.message.text.stub.v0
        protocol_version=carbonstack-openmls-sidecar-v0
      expected:
        HTTP 400
        unsupported_protocol_version

    accepted envelope:
      POST /v0/envelopes with:
        content_type=carbonstack.mls.application-message.v0
        protocol_version=carbonstack-openmls-sidecar-v0
      GET /v0/devices/<device_id>/envelopes
      POST /v0/envelopes/<envelope_id>/ack
      GET /v0/devices/<device_id>/envelopes

    restart:
      stop Cypher
      restart against same temp DB
      GET /v0/health
      GET /v0/accounts/<account_id>/devices
      GET /v0/devices/<device_id>/envelopes

    cleanup:
      stop process
      remove temp binary
      remove temp DB / proof dir
      leave no source-tree artifacts

### Current mainline public-surface state

    Top-level README:
      v0.3.20 remains current public testing release.
      v0.3.32/v0.3.33 is current mainline validation state.
      v0.4.0 is framed as a broad local deployability pre-release.

    Roadmap:
      late v0.3.x now points toward v0.3.34 release checklist/package validation rehearsal, v0.3.35 RC wording/asset plan, and v0.3.36+ release cut if clean.

    docs/README:
      current release/validation docs list now includes docs/152.

### Accepted envelope pair

    content_type=carbonstack.mls.application-message.v0
    protocol_version=carbonstack-openmls-sidecar-v0

### Negative protocol pair now validated

    content_type=carbonstack.message.text.stub.v0
    protocol_version=carbonstack-openmls-sidecar-v0

Expected result:

    HTTP 400
    error code unsupported_protocol_version

### Other future negative path candidates

    invalid base64 envelope
    ack unknown envelope
    ack with wrong recipient
    missing recipient_device_id on ack
    malformed account device path
    duplicate invite code
    already-claimed invite

### Critical functions / behavior preserved from v0.3.24

    Store.Migrate:
      ensures schema_migrations exists;
      reads migration .sql files in sorted order;
      computes sha256 checksum for each migration;
      skips matching already-applied migrations;
      fails on checksum mismatch for already-applied migration;
      applies new migrations in a transaction;
      records migration after successful application.

    TestMigrateRecordsAppliedMigrations:
      proves applied migrations are recorded with non-empty checksum/applied_at.

    TestMigrateCanRunTwiceAgainstSameDB:
      proves repeated migration against the same DB now succeeds.

    TestMigrateDetectsAppliedMigrationChecksumMismatch:
      proves already-applied migration drift hard-fails.

---

## 9. Lean Breakpoint Summary

    CarbonStack v0.3.33 is a stable mainline pre-v0.4.0 release-surface cleanup breakpoint after v0.3.32. It does not add new validation features. It cleans the top-level README, roadmap, and docs index so they reflect the current split between v0.3.20 as the current public runner-backed testing release and v0.3.32/v0.3.33 as the current mainline WSL Debian validation state. It records v0.4.0 as a broad local deployability pre-release / research-and-development milestone, not public-user readiness, application-use readiness, production security, hostile-server certification, mature Comms, Android, or CarbonStackOS. local-cypher, doctor, and core --clean-generated passed under WSL Debian after the docs commit, and final clean status showed no dirty files across all four repos. Next safest rung is v0.3.34 release checklist and package validation rehearsal.

---

## 10. Preserved v0.3.32 Operational Process Log

The following section preserves the previous v0.3.32 LogDoc body as operational continuity. It is intentionally retained rather than compressed away. Where v0.3.32 text conflicts with the v0.3.33 current-state overlay above, v0.3.33 wins for current state; v0.3.32 remains the provenance/process ledger for the local-cypher negative protocol validation work and the preserved v0.3.31 / v0.3.30 / v0.3.29 / v0.3.28 / v0.3.27 / v0.3.26 / v0.3.25 / v0.3.24 / v0.3.23 / v0.3.22 / v0.3.21 / v0.3.20 history.



# CarbonStack LogDoc v0.3.32

**Last updated:** 2026-06-03 local session  
**Owner/operator:** bitcrusher32  
**Assistant/model:** GPT-5.5 Thinking  
**Current phase:** Phase 2E-prep / v0.3.x **local-cypher negative protocol validation checkpoint**. `carbonstack` is now at `cff3ab4 docs: record local-cypher negative protocol validation`, after the v0.3.32 runner patch added the first negative-path check to the `local-cypher` profile. `carbonstack-cypher` remains at `9ab994c docs: point to local operator runbook`; `carbonstack-comms` remains at `012c8bf scripts: support source snapshot artifact guard`; `carbonstack-os` remains at `1bbbe52 docs: clarify CarbonStackOS target direction`. The live WSL Debian working umbrella at `~/repos/carbonstack_umbrella` passed `local-cypher`, `doctor`, and `core --clean-generated` after the v0.3.32 docs commit. `local-cypher` now explicitly proves the historical stub-text/OpenMLS protocol mismatch is rejected before the positive lifecycle continues, and final clean status showed no dirty files across the four repos.

**Version schema:** `v[scope].[timeline]`. This file is `v0.3.32`, the local-cypher negative protocol validation handoff after the `v0.3.31` local-cypher polish / generated-artifact cleanup-control checkpoint. Per LogDoc V2 practice, this Markdown file intentionally preserves verbose operational continuity; the JSON breakpoint remains the lean handoff state. The dense v0.3.31, v0.3.30, v0.3.29, v0.3.28, v0.3.27, v0.3.26, v0.3.25, v0.3.24, v0.3.23, v0.3.22, v0.3.21, and v0.3.20 operational process logs are preserved later in this file rather than compressed away.

**LogDoc V2 usage note:** LogDoc V2 uses a split between a Markdown LogDoc for working-session continuity, a JSON state file for end-of-session handoff, Git for history preservation, and README/docs for public claims. This v0.3.32 Markdown file therefore intentionally stays “bloated” where operational process, blunders, validation ladders, and historical context matter. The v0.3.32 JSON should remain lean and should agree with this file on the latest current state.

---

## 1. Project Goal

**Active goal:** Preserve CarbonStack state after adding and documenting the first negative-path validation inside the implemented `local-cypher` runner profile. The immediate next line of work should be release-surface readiness preflight and/or a review of whether this single negative path is sufficient for the pre-v0.4.0 checkpoint. This is not `local-backbone`, not runtime Comms UX, not public ingress, not systemd/cloudflared, not real homelab deployment, not Android/CarbonStackOS work, and not the pre-v0.4.0 release cut yet.

CarbonStack’s broader project intent remains:

- **CarbonStack** is the secure-communications appliance-stack project and shared doctrine/release authority.
- **CarbonStackOS** is the future constrained Android-derived appliance OS.
- **CarbonStackComms** is the text-first encrypted messaging client and current OpenMLS sidecar/relay integration surface.
- **CarbonStackCypher** is the self-hostable relay/storage stack.
- The `carbonstack` repo remains the doctrine/spec/docs/release authority; implementation repos must not contradict it without revising the spec.

**North star:**

> Its goal is not to be a better smartphone. Its goal is to stop being a smartphone.

**Core doctrine:**

> Every feature is guilty until it proves it does not add unacceptable parser, network, sensor, identity, or filesystem authority.

**Current reality after v0.3.32:** `v0.3.20` remains the current public runner-backed testing release. `v0.3.21` defined the local-only deployability model. `v0.3.22` narrowed the local deployability surface around Cypher operator config/API/database assumptions and adopted **CarbonStack Relay Space** terminology. `v0.3.23` proved the SQLite repeated-migration persistence hazard. `v0.3.24` implemented `schema_migrations` tracking in CarbonStackCypher. `v0.3.25` recorded the first local-only Cypher operator runbook skeleton. `v0.3.26` recorded the local operator config/data convention. `v0.3.27` proved a real local-only Cypher explicit-env API lifecycle. `v0.3.28` recorded the helper/runner decision recon and reserved `local-backbone` for later whole-stack validation. `v0.3.29` defined the future `local-cypher` validation contract. `v0.3.30` implemented that contract as the first `local-cypher` profile in the Go runner. `v0.3.31` polished `local-cypher` output and added explicit opt-in `--clean-generated`. `v0.3.32` now adds the first negative-path validation to `local-cypher`: the historical stub-text/OpenMLS protocol mismatch must fail with `unsupported_protocol_version` before the valid positive lifecycle is allowed to continue.

The current public release page remains:

    https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.3.20

The current public testing interpretation remains:

    v0.3.20 is the preferred public runner-backed testing release.
    v0.3.4 remains available as the older Windows 11 / PowerShell legacy testbed.
    Debian / WSL Debian is now the mainline public dev/test validation direction.
    Windows 11 short-path validation was the final explicit Windows dev/test validation for this phase.

v0.4.0 release framing direction remains:

    v0.4.0 should be framed as a broad local deployability pre-release.
    It should be positioned as a milestone / research-and-development release.
    It should clearly state that it is not intended for public-user use.
    It should clearly state that it is not intended for application use.
    It should clearly state that it still requires testing, integration, and maturity before real use.
    Its concrete validated artifact should remain the WSL Debian runner-backed validation surface, including local-cypher/core.

This is still **not** a production deployability release. It does **not** prove production readiness, production E2EE, hostile-server safety, metadata privacy, Android readiness, secure vault/storage, Debian production deployability, systemd readiness, cloudflared readiness, public ingress safety, real homelab validation, external audit, or certification.

---

## 2. Current Repo Heads

    carbonstack        cff3ab4 (HEAD -> main, origin/main, origin/HEAD) docs: record local-cypher negative protocol validation
    carbonstack-comms  012c8bf (HEAD -> main, origin/main, origin/HEAD) scripts: support source snapshot artifact guard
    carbonstack-cypher 9ab994c (HEAD -> main, origin/main, origin/HEAD) docs: point to local operator runbook
    carbonstack-os     1bbbe52 (HEAD -> main, origin/main, origin/HEAD) docs: clarify CarbonStackOS target direction

Current `carbonstack` public runner-backed testing release tag:

    v0.3.20 -> 3eeb1a1e1a / 3eeb1a1 docs: clean public release surface

Previous mainline deployability checkpoints:

    v0.3.32 mainline -> cff3ab4 docs: record local-cypher negative protocol validation
    v0.3.31 mainline -> 7894278 docs: record local-cypher polish and cleanup control
    v0.3.30 mainline -> 15a3758 docs: record local-cypher runner implementation
    v0.3.30 runner  -> a98ee9f runner: add local-cypher
    v0.3.29 mainline -> 37a6d2a docs: record local-cypher validation contract
    v0.3.28 mainline -> af33139 docs: record local operator helper runner decision recon
    v0.3.27 mainline -> e516fc7 docs: record local Cypher API lifecycle proof
    v0.3.26 mainline -> 6e14fe0 docs: record local operator config data convention
    v0.3.25 mainline -> a5c6351 docs: update readme, roadmap, and historical doc
    v0.3.25 cypher  -> 9ab994c docs: point to local operator runbook
    v0.3.24 mainline -> 76f0258 docs: record Cypher schema migrations implementation
    v0.3.24 cypher  -> 42f838f db: track applied migrations
    v0.3.23 mainline -> 1fef7fe docs: record Cypher migration persistence recon
    v0.3.22 mainline -> 328ffc4 docs: record Cypher local operator surface recon
    v0.3.21 mainline -> 6850c58 docs: record local backbone deployability recon

Legacy `carbonstack` Windows/PowerShell testbed release tag:

    v0.3.4 -> 7c5b7e6 docs: record release snapshot self-test fix

Historical `carbonstack` release tag:

    v0.3.0 -> 59245fc docs: freeze v0.3.0 release packaging plan

Important continuity note: `v0.3.32` does not supersede `v0.3.20` as a public release. It advances mainline local-only deployability validation by adding one negative protocol check to the Cypher-only `local-cypher` runner profile. It is safe to pause here because the runner negative-path implementation and docs/result record were pushed, WSL Debian `local-cypher` / `doctor` / `core --clean-generated` passed after the docs commit, and final status was clean across all four repos.

---

## 3. Current Validated State

Validated at the v0.3.32 checkpoint:

    [WSL-DEBIAN] working umbrella exists at ~/repos/carbonstack_umbrella
    [WSL-DEBIAN] carbonstack is at cff3ab4
    [WSL-DEBIAN] carbonstack-comms is at 012c8bf
    [WSL-DEBIAN] carbonstack-cypher is at 9ab994c
    [WSL-DEBIAN] carbonstack-os is at 1bbbe52
    [CARBONSTACK] runner negative-path patch added invalid stub-text/OpenMLS protocol rejection to local-cypher
    [CARBONSTACK] docs/151 local-cypher negative protocol validation record committed and pushed
    [GO-RUNNER] local-cypher passed after the v0.3.32 docs commit
    [GO-RUNNER] doctor passed after the v0.3.32 docs commit
    [GO-RUNNER] core --clean-generated passed after the v0.3.32 docs commit
    [ARTIFACT CLEANUP] known OpenMLS generated roots were removed by the runner's explicit --clean-generated flag
    [FINAL STATUS] all four repos were clean after validation without a separate manual cleanup step

Observed WSL Debian toolchain baseline:

    git version 2.47.3
    go version go1.24.4 linux/amd64
    rustc 1.96.0 (ac68faa20 2026-05-25)
    cargo 1.96.0 (30a34c682 2026-05-25)
    sqlite3 3.46.1

Observed `local-cypher` validation result during the v0.3.32 final flow:

    required paths: PASS
    pre-local-cypher artifact scan: PASS / no generated/private/build artifact hits
    temporary Cypher binary build: PASS
    first health check: PASS
    seeded Alice invite claim: PASS
    Bob dev invite creation: PASS
    Bob invite claim: PASS
    Alice device registration: PASS
    Bob device registration: PASS
    Alice device listing: PASS
    invalid stub-text/OpenMLS protocol pairing rejection: PASS
    opaque OpenMLS application-message envelope submit: PASS
    Bob inbox retrieval and payload metadata verification: PASS
    envelope ack: PASS
    Bob inbox after ack empty: PASS
    first Cypher process stopped with expected SIGINT termination
    restart against same temporary DB: PASS
    restart health check: PASS
    persisted Alice device state after restart: PASS
    Bob's acked inbox remains empty after restart: PASS
    second Cypher process stopped with expected SIGINT termination
    post-local-cypher artifact scan: PASS / no generated/private/build artifact hits
    VALIDATION PASSED

Observed `core --clean-generated` validation result during the v0.3.32 final flow:

    doctor nested inside core: PASS
    pre-test artifact scan: PASS / no generated/private/build artifact hits
    OpenMLS real-Cypher lifecycle: PASS
    carbonstack-comms package tests: PASS
    carbonstack-cypher package tests: PASS
    post-test artifact scan: expected OpenMLS sidecar generated roots only
    clean generated artifacts step: PASS
    removed .carbonstack-openmls-sidecar-state
    removed target
    VALIDATION PASSED

Expected/generated roots handled by `--clean-generated`:

    carbonstack-comms/internal/protocol/mls/openmls-sidecar/.carbonstack-openmls-sidecar-state
    carbonstack-comms/internal/protocol/mls/openmls-sidecar/target

Still absent unless created by future behavior:

    carbonstack-comms/internal/protocol/mls/openmls-sidecar/provider-storage.json
    carbonstack-comms/internal/protocol/mls/openmls-sidecar/signer.json

---

## 4. v0.3.32 Work Completed

### 4.1 Recon confirmed the exact error contract before patching

Before patching, WSL Debian recon inspected the runner local-cypher surface and Cypher HTTP API tests/handlers.

The recon confirmed:

    local-cypher already had the positive lifecycle in local_cypher.go.
    Cypher already had the exact negative test in internal/httpapi/api_test.go.
    The historical bad pair is:
      content_type=carbonstack.message.text.stub.v0
      protocol_version=carbonstack-openmls-sidecar-v0
    The expected HTTP status is 400 / http.StatusBadRequest.
    The expected error code is unsupported_protocol_version.
    The response error shape is a JSON object with error.code.
    local-cypher and core --clean-generated both passed before patching.
    final repo status was clean before patching.

This preserved the project method:

    recon -> targeted patch -> validation -> docs/result -> final validation -> breakpoint

### 4.2 Initial insertion-target issue was avoided/fixed

An attempted patch style that depended on a brittle exact insertion target failed with:

    target insertion point not found in local_cypher.go

The durable fix was to patch by searching for the positive envelope submit marker:

    envelope, err := localCypherPOST(baseURL+"/v0/envelopes", map[string]any{

and insert the negative-path block immediately before it.

Lesson preserved:

    local_cypher.go is evolving; avoid overly brittle exact multi-line insertion anchors.
    Prefer short stable semantic markers around stable behavior.
    When patching runner lifecycle code, compile and run local-cypher before committing.

### 4.3 local-cypher negative-path validation landed

Updated runner files at v0.3.32:

    carbonstack/tools/carbonstack-validate/local_cypher.go
    carbonstack/tools/carbonstack-validate/README.md

New helper added:

    localCypherExpectErrorCode

New local-cypher validation step:

    PASS: reject invalid stub-text/OpenMLS protocol pairing

Negative request:

    content_type=carbonstack.message.text.stub.v0
    protocol_version=carbonstack-openmls-sidecar-v0

Expected response:

    HTTP 400
    unsupported_protocol_version

The check runs before the accepted opaque OpenMLS application-message envelope is submitted. The accepted pair remains:

    content_type=carbonstack.mls.application-message.v0
    protocol_version=carbonstack-openmls-sidecar-v0

### 4.4 README/docs/roadmap update landed

New/updated docs at v0.3.32:

    carbonstack/docs/151-local-cypher-negative-protocol-validation-v0.md
    carbonstack/docs/README.md
    carbonstack/roadmap/ROADMAP.md

Mainline docs commit:

    cff3ab4 docs: record local-cypher negative protocol validation

The docs record the single negative-path scope, the historical v0.3.27 blunder provenance, the accepted envelope pair, validation output, and the ongoing nonclaims.

### 4.5 Final validation passed with cleanup

Final WSL Debian validation sequence included:

    go run . --profile local-cypher
    go run . --profile doctor
    go run . --profile core --clean-generated

Final WSL Debian repo snapshot:

    carbonstack        cff3ab4 docs: record local-cypher negative protocol validation
    carbonstack-comms  012c8bf scripts: support source snapshot artifact guard
    carbonstack-cypher 9ab994c docs: point to local operator runbook
    carbonstack-os     1bbbe52 docs: clarify CarbonStackOS target direction

Final validation passed under WSL Debian. `local-cypher` passed and included the new negative-path line. `doctor` passed. `core --clean-generated` passed OpenMLS real-Cypher lifecycle, Comms package tests, and Cypher package tests, then removed the expected OpenMLS generated roots through the explicit cleanup flag. Final clean status showed no dirty files across all four repos.

---

## 5. Critical Blunders / Lessons Preserved

### 5.1 The v0.3.27 protocol-pairing blunder is now a validation asset

Original bad assumption from v0.3.27:

    content_type=carbonstack.message.text.stub.v0
    protocol_version=carbonstack-openmls-sidecar-v0

Cypher correctly rejected it as:

    unsupported_protocol_version

v0.3.32 now turns that preserved mistake into an explicit local-cypher validation step.

Correct lesson:

    blunders preserved in LogDoc can become regression tests.
    the runner now proves Cypher does not blindly accept an arbitrary content/protocol pairing.

### 5.2 Keep negative-path expansion narrow

v0.3.32 adds one negative path only. It does not add a broad adversarial suite.

Correct lesson:

    one historically grounded negative check is useful and low-risk.
    broader negative-path validation should be deliberate, not bundled casually.
    pre-v0.4.0 may not need a wide negative-path suite if release claims remain careful.

### 5.3 Brittle patch anchors caused an insertion failure

The target insertion point failure showed that lifecycle files are not stable enough for large exact multi-line replacement scripts.

Correct lesson:

    use short stable semantic markers.
    inspect and compile immediately after lifecycle runner patches.
    avoid mixing multiple unrelated changes in one patch when insertion is fragile.

### 5.4 local-cypher still remains Cypher-only

Even with one negative path, `local-cypher` is still not a full CarbonStack backbone proof.

Correct lesson:

    local-cypher validates Cypher API lifecycle behavior under dev/test constraints.
    it does not validate runtime Comms UX, hostile-server safety, metadata privacy, or production security.

### 5.5 --clean-generated continues to reduce workflow drag safely

The final validation used `core --clean-generated` and reached clean status without manual cleanup.

Correct lesson:

    explicit opt-in cleanup has become the preferred validation cadence.
    keep cleanup narrow and non-default.

---

## 6. Current Blockers / Not Validated

Current blockers and nonvalidated areas:

    local-cypher now implements positive Cypher-only lifecycle validation plus one negative protocol-pairing check.
    No broad negative-path suite exists yet.
    local-backbone runner profile does not exist and remains reserved for later whole-stack validation.
    No helper command exists yet.
    Local-only operator convention is documented and proof-tested, but not enforced by code outside local-cypher.
    Cypher default bind remains :8080; local-only docs recommend explicit 127.0.0.1:8080.
    Runtime Comms send/inbox remains stub-era and not wired as user-facing OpenMLS messenger UX.
    Production-safe provider storage and secure vault design are unsolved.
    Hostile-server rollback/replay/metadata-abuse validation remains future work.
    Android/CarbonStackOS implementation remains later work.
    CarbonStack Relay Space mechanics remain a defined future concept, not an implementation.
    Public ingress, systemd, cloudflared, real homelab deployment, and production service paths remain out of scope.
    Temporary DB validation remains dev/test-only and does not replace later persistent user/operator-state validation for public Comms maturity.
    --clean-generated is intentionally narrow and does not replace release-snapshot strict pre-test validation.

Hard nonclaims remain:

    Do not claim production readiness.
    Do not claim production E2EE.
    Do not claim hostile-server safety.
    Do not claim metadata privacy.
    Do not claim Debian production deployability.
    Do not claim systemd readiness.
    Do not claim cloudflared readiness.
    Do not claim Android readiness.
    Do not claim secure vault/storage.
    Do not claim external audit or certification.
    Do not claim CarbonStack Relay Space mechanics are implemented.
    Do not claim local-backbone exists or is justified yet.
    Do not claim local-cypher validates runtime Comms UX.
    Do not claim local-cypher is sufficient for public user testing maturity.
    Do not claim --clean-generated is a general cleanup tool.
    Do not claim v0.3.32 is a public release.

Updated allowed claim:

    CarbonStack v0.3.32 adds the first negative-path validation to local-cypher.
    local-cypher now rejects the historical stub-text/OpenMLS protocol mismatch with unsupported_protocol_version before continuing the positive lifecycle.
    local-cypher, doctor, and core --clean-generated passed under WSL Debian after the v0.3.32 docs commit.
    Final clean status showed no dirty files across the four repos without manual rm -rf cleanup.

---

## 7. Next Safest Actions

Recommended next sequence after this v0.3.32 breakpoint:

    v0.3.33 preferred:
      local-cypher negative-path result review and pre-v0.4.0 release-surface readiness preflight

Recommended concrete next rung:

    1. Review whether the single negative path is sufficient for the pre-v0.4.0 broad local deployability pre-release.
    2. Decide whether additional negative paths should be deferred until after v0.4.0 or added before release.
    3. Begin pre-v0.4.0 README/docs/roadmap claim-cleanup planning.
    4. Confirm release framing: milestone / research-and-development release, not public-user/application-use readiness.
    5. Keep local-backbone deferred.
    6. Keep runtime Comms UX deferred to v0.4.x.
    7. Validate local-cypher, doctor, and core --clean-generated after any runner/doc changes.
    8. Continue using --clean-generated for core validation rather than manual rm -rf cleanup.

Avoid next:

    local-backbone runner profile
    helper implementation
    public ingress
    cloudflared
    systemd
    real homelab deployment
    runtime Comms send/inbox UX
    Android app
    CarbonStackOS
    CarbonStack Relay Space implementation
    production release

---

## 8. Critical Paths / Functions

### WSL Debian active working umbrella

    [WSL-DEBIAN] umbrella: ~/repos/carbonstack_umbrella
    [WSL-DEBIAN] carbonstack: ~/repos/carbonstack_umbrella/carbonstack
    [WSL-DEBIAN] carbonstack-comms: ~/repos/carbonstack_umbrella/carbonstack-comms
    [WSL-DEBIAN] carbonstack-cypher: ~/repos/carbonstack_umbrella/carbonstack-cypher
    [WSL-DEBIAN] carbonstack-os: ~/repos/carbonstack_umbrella/carbonstack-os
    [GO-RUNNER] runner dir: ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate

### WSL Debian known-good validation commands

From runner dir:

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go run . --profile local-cypher
    go run . --profile doctor
    go run . --profile core --clean-generated

From carbonstack repo root:

    cd ~/repos/carbonstack_umbrella/carbonstack
    go run ./tools/carbonstack-validate --profile local-cypher
    go run ./tools/carbonstack-validate --profile doctor
    go run ./tools/carbonstack-validate --profile core --clean-generated

### v0.3.32 docs

    [DOC] v0.3.32 negative protocol validation result:
    carbonstack/docs/151-local-cypher-negative-protocol-validation-v0.md

    [DOC] v0.3.31 polish/cleanup result:
    carbonstack/docs/150-local-cypher-polish-generated-cleanup-v0.md

    [DOC] v0.3.30 implementation result:
    carbonstack/docs/149-local-cypher-runner-implementation-v0.md

    [DOC] v0.3.29 contract:
    carbonstack/docs/148-local-cypher-validation-contract-v0.md

    [DOC] docs index:
    carbonstack/docs/README.md

    [DOC] roadmap:
    carbonstack/roadmap/ROADMAP.md

### Go runner surfaces after v0.3.32

    [RUNNER] directory:
    carbonstack/tools/carbonstack-validate

    [RUNNER] implementation files:
    carbonstack/tools/carbonstack-validate/main.go
    carbonstack/tools/carbonstack-validate/local_cypher.go
    carbonstack/tools/carbonstack-validate/checksums.go
    carbonstack/tools/carbonstack-validate/release_snapshot.go
    carbonstack/tools/carbonstack-validate/README.md
    carbonstack/tools/carbonstack-validate/go.mod

    [RUNNER] profiles:
    doctor
    core
    local-cypher
    full
    release-snapshot
    write-checksums
    verify-checksums

    [RUNNER] cleanup flag:
    --clean-generated

### --clean-generated behavior

    Runs only after successful profile execution.
    Deletes only:
      carbonstack-comms/internal/protocol/mls/openmls-sidecar/.carbonstack-openmls-sidecar-state
      carbonstack-comms/internal/protocol/mls/openmls-sidecar/target

    Does not delete:
      arbitrary untracked files
      manual local operator DBs
      $HOME/.local/share/carbonstack/cypher/cypher.db
      provider-storage.json unless added by a future deliberate expansion
      signer.json unless added by a future deliberate expansion

    Does not replace:
      ArtifactScan
      StrictPreTestArtifactScan
      release-snapshot pre-test hygiene checks

### local-cypher lifecycle after v0.3.32

    start:
      build temporary Cypher binary
      create temporary DB directory
      choose dynamic loopback port
      set CYPHER_ADDR to 127.0.0.1:<dynamic_port>
      set CYPHER_DB to temporary SQLite DB path
      set CYPHER_MIGRATIONS to source-tree migrations
      set CYPHER_DEV_INVITE to known temporary dev invite
      start Cypher
      wait for GET /v0/health

    account/device:
      POST /v0/invites/claim
      POST /v0/dev/invites
      POST /v0/invites/claim
      POST /v0/devices/register
      GET /v0/accounts/<account_id>/devices

    negative envelope check:
      POST /v0/envelopes with:
        content_type=carbonstack.message.text.stub.v0
        protocol_version=carbonstack-openmls-sidecar-v0
      expected:
        HTTP 400
        unsupported_protocol_version

    accepted envelope:
      POST /v0/envelopes with:
        content_type=carbonstack.mls.application-message.v0
        protocol_version=carbonstack-openmls-sidecar-v0
      GET /v0/devices/<device_id>/envelopes
      POST /v0/envelopes/<envelope_id>/ack
      GET /v0/devices/<device_id>/envelopes

    restart:
      stop Cypher
      restart against same temp DB
      GET /v0/health
      GET /v0/accounts/<account_id>/devices
      GET /v0/devices/<device_id>/envelopes

    cleanup:
      stop process
      remove temp binary
      remove temp DB / proof dir
      leave no source-tree artifacts

### Accepted envelope pair

    content_type=carbonstack.mls.application-message.v0
    protocol_version=carbonstack-openmls-sidecar-v0

### Negative protocol pair now validated

    content_type=carbonstack.message.text.stub.v0
    protocol_version=carbonstack-openmls-sidecar-v0

Expected result:

    HTTP 400
    error code unsupported_protocol_version

### Other future negative path candidates

    invalid base64 envelope
    ack unknown envelope
    ack with wrong recipient
    missing recipient_device_id on ack
    malformed account device path
    duplicate invite code
    already-claimed invite

### Critical functions / behavior preserved from v0.3.24

    Store.Migrate:
      ensures schema_migrations exists;
      reads migration .sql files in sorted order;
      computes sha256 checksum for each migration;
      skips matching already-applied migrations;
      fails on checksum mismatch for already-applied migration;
      applies new migrations in a transaction;
      records migration after successful application.

    TestMigrateRecordsAppliedMigrations:
      proves applied migrations are recorded with non-empty checksum/applied_at.

    TestMigrateCanRunTwiceAgainstSameDB:
      proves repeated migration against the same DB now succeeds.

    TestMigrateDetectsAppliedMigrationChecksumMismatch:
      proves already-applied migration drift hard-fails.

---

## 9. Lean Breakpoint Summary

    CarbonStack v0.3.32 is a stable mainline negative-path validation breakpoint after v0.3.31. It adds the first local-cypher negative protocol validation by turning the preserved v0.3.27 stub-text/OpenMLS pairing blunder into a runner check: content_type=carbonstack.message.text.stub.v0 with protocol_version=carbonstack-openmls-sidecar-v0 must return HTTP 400 / unsupported_protocol_version. local-cypher still performs the positive invite/device/envelope/ack lifecycle, restart against the same temporary DB, persisted-state checks, and artifact scans. doctor and core --clean-generated passed under WSL Debian after the docs commit, --clean-generated removed known OpenMLS generated roots, and final clean status showed no dirty files across all four repos. local-backbone, helper tooling, broader negative-path suite, runtime Comms UX, CarbonStack Relay Space mechanics, public ingress, systemd, cloudflared, homelab, Android, CarbonStackOS, hostile-server proof, secure vault, persistent user/operator-state maturity, and production/security claims remain deferred.

---

## 10. Preserved v0.3.31 Operational Process Log

The following section preserves the previous v0.3.31 LogDoc body as operational continuity. It is intentionally retained rather than compressed away. Where v0.3.31 text conflicts with the v0.3.32 current-state overlay above, v0.3.32 wins for current state; v0.3.31 remains the provenance/process ledger for the local-cypher polish / generated-artifact cleanup-control work and the preserved v0.3.30 / v0.3.29 / v0.3.28 / v0.3.27 / v0.3.26 / v0.3.25 / v0.3.24 / v0.3.23 / v0.3.22 / v0.3.21 / v0.3.20 history.



# CarbonStack LogDoc v0.3.31

**Last updated:** 2026-06-03 local session  
**Owner/operator:** bitcrusher32  
**Assistant/model:** GPT-5.5 Thinking  
**Current phase:** Phase 2E-prep / v0.3.x **local-cypher polish and explicit generated-artifact cleanup checkpoint**. `carbonstack` is now at `7894278 docs: record local-cypher polish and cleanup control`, after `v0.3.31` first landed the runner cleanup/polish implementation at the preceding runner commit for the explicit `--clean-generated` flag. `carbonstack-cypher` remains at `9ab994c docs: point to local operator runbook`; `carbonstack-comms` remains at `012c8bf scripts: support source snapshot artifact guard`; `carbonstack-os` remains at `1bbbe52 docs: clarify CarbonStackOS target direction`. The live WSL Debian working umbrella at `~/repos/carbonstack_umbrella` passed `local-cypher`, `doctor`, and `core --clean-generated` after the v0.3.31 docs commit, and final clean status showed no dirty files across the four repos without a separate manual `rm -rf` cleanup step.

**Version schema:** `v[scope].[timeline]`. This file is `v0.3.31`, the local-cypher polish / generated-artifact cleanup-control handoff after the `v0.3.30` local-cypher runner implementation checkpoint. Per LogDoc V2 practice, this Markdown file intentionally preserves verbose operational continuity; the JSON breakpoint remains the lean handoff state. The dense v0.3.30, v0.3.29, v0.3.28, v0.3.27, v0.3.26, v0.3.25, v0.3.24, v0.3.23, v0.3.22, v0.3.21, and v0.3.20 operational process logs are preserved later in this file rather than compressed away.

**LogDoc V2 usage note:** LogDoc V2 uses a split between a Markdown LogDoc for working-session continuity, a JSON state file for end-of-session handoff, Git for history preservation, and README/docs for public claims. This v0.3.31 Markdown file therefore intentionally stays “bloated” where operational process, blunders, validation ladders, and historical context matter. The v0.3.31 JSON should remain lean and should agree with this file on the latest current state.

---

## 1. Project Goal

**Active goal:** Preserve CarbonStack state after polishing the implemented `local-cypher` runner profile and adding an explicit generated-artifact cleanup flag. The immediate next line of work should be a narrow v0.3.32 negative-path validation rung for `local-cypher`, likely starting with the historically grounded `unsupported_protocol_version` case from the v0.3.27 blunder. This is not `local-backbone`, not runtime Comms UX, not public ingress, not systemd/cloudflared, not real homelab deployment, not Android/CarbonStackOS work, and not the pre-v0.4.0 release cut yet.

CarbonStack’s broader project intent remains:

- **CarbonStack** is the secure-communications appliance-stack project and shared doctrine/release authority.
- **CarbonStackOS** is the future constrained Android-derived appliance OS.
- **CarbonStackComms** is the text-first encrypted messaging client and current OpenMLS sidecar/relay integration surface.
- **CarbonStackCypher** is the self-hostable relay/storage stack.
- The `carbonstack` repo remains the doctrine/spec/docs/release authority; implementation repos must not contradict it without revising the spec.

**North star:**

> Its goal is not to be a better smartphone. Its goal is to stop being a smartphone.

**Core doctrine:**

> Every feature is guilty until it proves it does not add unacceptable parser, network, sensor, identity, or filesystem authority.

**Current reality after v0.3.31:** `v0.3.20` remains the current public runner-backed testing release. `v0.3.21` defined the local-only deployability model. `v0.3.22` narrowed the local deployability surface around Cypher operator config/API/database assumptions and adopted **CarbonStack Relay Space** terminology. `v0.3.23` proved the SQLite repeated-migration persistence hazard. `v0.3.24` implemented `schema_migrations` tracking in CarbonStackCypher. `v0.3.25` recorded the first local-only Cypher operator runbook skeleton. `v0.3.26` recorded the local operator config/data convention. `v0.3.27` proved a real local-only Cypher explicit-env API lifecycle. `v0.3.28` recorded the helper/runner decision recon and reserved `local-backbone` for later whole-stack validation. `v0.3.29` defined the future `local-cypher` validation contract. `v0.3.30` implemented that contract as the first `local-cypher` profile in the Go runner. `v0.3.31` now polishes `local-cypher` toolchain output and adds the explicit opt-in `--clean-generated` flag, allowing known OpenMLS sidecar generated roots to be removed after successful validation without manual shell cleanup.

The current public release page remains:

    https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.3.20

The current public testing interpretation remains:

    v0.3.20 is the preferred public runner-backed testing release.
    v0.3.4 remains available as the older Windows 11 / PowerShell legacy testbed.
    Debian / WSL Debian is now the mainline public dev/test validation direction.
    Windows 11 short-path validation was the final explicit Windows dev/test validation for this phase.

v0.4.0 release framing direction:

    v0.4.0 should be framed as a broad local deployability pre-release.
    It should be positioned as a milestone / research-and-development release.
    It should clearly state that it is not intended for public-user use.
    It should clearly state that it is not intended for application use.
    It should clearly state that it still requires testing, integration, and maturity before real use.
    Its concrete validated artifact should remain the WSL Debian runner-backed validation surface, including local-cypher/core.

This is still **not** a production deployability release. It does **not** prove production readiness, production E2EE, hostile-server safety, metadata privacy, Android readiness, secure vault/storage, Debian production deployability, systemd readiness, cloudflared readiness, public ingress safety, real homelab validation, external audit, or certification.

---

## 2. Current Repo Heads

    carbonstack        7894278 (HEAD -> main, origin/main, origin/HEAD) docs: record local-cypher polish and cleanup control
    carbonstack-comms  012c8bf (HEAD -> main, origin/main, origin/HEAD) scripts: support source snapshot artifact guard
    carbonstack-cypher 9ab994c (HEAD -> main, origin/main, origin/HEAD) docs: point to local operator runbook
    carbonstack-os     1bbbe52 (HEAD -> main, origin/main, origin/HEAD) docs: clarify CarbonStackOS target direction

Current `carbonstack` public runner-backed testing release tag:

    v0.3.20 -> 3eeb1a1e1a / 3eeb1a1 docs: clean public release surface

Previous mainline deployability checkpoints:

    v0.3.31 mainline -> 7894278 docs: record local-cypher polish and cleanup control
    v0.3.30 mainline -> 15a3758 docs: record local-cypher runner implementation
    v0.3.30 runner  -> a98ee9f runner: add local-cypher
    v0.3.29 mainline -> 37a6d2a docs: record local-cypher validation contract
    v0.3.28 mainline -> af33139 docs: record local operator helper runner decision recon
    v0.3.27 mainline -> e516fc7 docs: record local Cypher API lifecycle proof
    v0.3.26 mainline -> 6e14fe0 docs: record local operator config data convention
    v0.3.25 mainline -> a5c6351 docs: update readme, roadmap, and historical doc
    v0.3.25 cypher  -> 9ab994c docs: point to local operator runbook
    v0.3.24 mainline -> 76f0258 docs: record Cypher schema migrations implementation
    v0.3.24 cypher  -> 42f838f db: track applied migrations
    v0.3.23 mainline -> 1fef7fe docs: record Cypher migration persistence recon
    v0.3.22 mainline -> 328ffc4 docs: record Cypher local operator surface recon
    v0.3.21 mainline -> 6850c58 docs: record local backbone deployability recon

Legacy `carbonstack` Windows/PowerShell testbed release tag:

    v0.3.4 -> 7c5b7e6 docs: record release snapshot self-test fix

Historical `carbonstack` release tag:

    v0.3.0 -> 59245fc docs: freeze v0.3.0 release packaging plan

Important continuity note: `v0.3.31` does not supersede `v0.3.20` as a public release. It advances mainline local-only deployability validation by polishing the Cypher-only `local-cypher` runner profile and adding explicit cleanup control. It is safe to pause here because the runner polish/cleanup implementation and docs/result record were pushed, WSL Debian `local-cypher` / `doctor` / `core --clean-generated` passed after the docs commit, and final status was clean across all four repos.

---

## 3. Current Validated State

Validated at the v0.3.31 checkpoint:

    [WSL-DEBIAN] working umbrella exists at ~/repos/carbonstack_umbrella
    [WSL-DEBIAN] carbonstack is at 7894278
    [WSL-DEBIAN] carbonstack-comms is at 012c8bf
    [WSL-DEBIAN] carbonstack-cypher is at 9ab994c
    [WSL-DEBIAN] carbonstack-os is at 1bbbe52
    [CARBONSTACK] runner cleanup/polish commit added the explicit --clean-generated flag and local-cypher output polish
    [CARBONSTACK] docs/150 local-cypher polish / generated cleanup control record committed and pushed
    [GO-RUNNER] local-cypher passed after the v0.3.31 docs commit
    [GO-RUNNER] doctor passed after the v0.3.31 docs commit
    [GO-RUNNER] core --clean-generated passed after the v0.3.31 docs commit
    [ARTIFACT CLEANUP] known OpenMLS generated roots were removed by the runner's explicit --clean-generated flag
    [FINAL STATUS] all four repos were clean after validation without a separate manual cleanup step

Observed WSL Debian toolchain baseline:

    git version 2.47.3
    go version go1.24.4 linux/amd64
    rustc 1.96.0 (ac68faa20 2026-05-25)
    cargo 1.96.0 (30a34c682 2026-05-25)
    sqlite3 3.46.1

Observed `local-cypher` validation result during the v0.3.31 final flow:

    required paths: PASS
    local-cypher toolchain block is now cleaner and no longer prints the prior duplicated wrapper labels
    pre-local-cypher artifact scan: PASS / no generated/private/build artifact hits
    temporary Cypher binary build: PASS
    first health check: PASS
    seeded Alice invite claim: PASS
    Bob dev invite creation: PASS
    Bob invite claim: PASS
    Alice device registration: PASS
    Bob device registration: PASS
    Alice device listing: PASS
    opaque OpenMLS application-message envelope submit: PASS
    Bob inbox retrieval and payload metadata verification: PASS
    envelope ack: PASS
    Bob inbox after ack empty: PASS
    first Cypher process stopped with expected SIGINT termination
    restart against same temporary DB: PASS
    restart health check: PASS
    persisted Alice device state after restart: PASS
    Bob's acked inbox remains empty after restart: PASS
    second Cypher process stopped with expected SIGINT termination
    post-local-cypher artifact scan: PASS / no generated/private/build artifact hits
    VALIDATION PASSED

Observed `core --clean-generated` validation result during the v0.3.31 final flow:

    doctor nested inside core: PASS
    pre-test artifact scan: PASS / no generated/private/build artifact hits
    OpenMLS real-Cypher lifecycle: PASS
    carbonstack-comms package tests: PASS
    carbonstack-cypher package tests: PASS
    post-test artifact scan: expected OpenMLS sidecar generated roots only
    clean generated artifacts step: PASS
    removed .carbonstack-openmls-sidecar-state
    removed target
    VALIDATION PASSED

Expected/generated roots handled by `--clean-generated`:

    carbonstack-comms/internal/protocol/mls/openmls-sidecar/.carbonstack-openmls-sidecar-state
    carbonstack-comms/internal/protocol/mls/openmls-sidecar/target

Still absent during/after this flow unless created by future behavior:

    carbonstack-comms/internal/protocol/mls/openmls-sidecar/provider-storage.json
    carbonstack-comms/internal/protocol/mls/openmls-sidecar/signer.json

---

## 4. v0.3.31 Work Completed

### 4.1 Full recon happened before polish

Before patching, a full WSL Debian recon inspected the runner file inventory, help/profile surface, profile dispatch, flags, local-cypher implementation surface, toolchain reporting surface, artifact scan surface, strict pre-test/checksum exclusions, generated artifact real paths, local-cypher runtime behavior, core runtime behavior, and docs continuity pointers.

The recon established:

    local-cypher itself leaves no dirty files.
    local-cypher does not create OpenMLS sidecar generated roots.
    core creates expected OpenMLS generated roots.
    provider-storage.json and signer.json were absent.
    manual cleanup after every core run had become workflow drag.
    local-cypher toolchain output had duplicate/noisy reporting.
    runner artifact classification already recognizes the known generated roots.
    docs already point through v0.3.30.
    v0.3.31 naturally fits as a polish/cleanup-control rung.

This preserved the project method:

    recon -> docs/decision -> implementation patch -> validation -> cleanup -> breakpoint

### 4.2 Initial partial patch blunder happened and was fixed

The first v0.3.31 cleanup-flag patch partially applied but failed to wire the new flag into the `Runner`.

Observed failure:

    ./main.go:49:2: declared and not used: cleanGenerated

The partial patch had added the `cleanGenerated` flag variable and `CleanGenerated` field, but it had not inserted:

    r.CleanGenerated = *cleanGenerated

and the post-success cleanup call / `CleanGeneratedArtifacts` method had not landed.

The fix was to explicitly patch `main.go` so that:

    `cleanGenerated` is wired into the runner after `NewRunner`.
    cleanup runs only after successful profile execution.
    cleanup failure exits as validation failure.
    `CleanGeneratedArtifacts` exists and only removes known generated roots.

Lesson preserved:

    partial scripted patches can silently miss insertion points.
    after patch scripts, always run `go test ./... -count=1` before moving forward.
    the failure was a compile-time runner patch failure, not a CarbonStack validation/runtime failure.

### 4.3 Explicit `--clean-generated` flag landed

Updated runner files at v0.3.31:

    carbonstack/tools/carbonstack-validate/main.go
    carbonstack/tools/carbonstack-validate/local_cypher.go
    carbonstack/tools/carbonstack-validate/README.md

New flag:

    --clean-generated

Example:

    go run . --profile core --clean-generated

Behavior:

    disabled by default
    runs only after successful profile execution
    prints every removed/skipped path
    only removes known generated roots
    does not remove arbitrary untracked files
    does not touch manual local operator DBs
    does not touch $HOME/.local/share/carbonstack/cypher/cypher.db
    does not replace artifact scans
    cleanup failure fails validation

Current cleanup scope:

    carbonstack-comms/internal/protocol/mls/openmls-sidecar/.carbonstack-openmls-sidecar-state
    carbonstack-comms/internal/protocol/mls/openmls-sidecar/target

The cleanup step is intentionally explicit and opt-in. It should not become default behavior without another decision rung.

### 4.4 local-cypher output polish landed

The local-cypher toolchain output was simplified.

Before v0.3.31, local-cypher printed confusing wrapper labels around a helper that already prints labels, causing duplicated-ish lines such as:

    go path: /usr/lib/go-1.24/bin/go
    go version: /usr/lib/go-1.24
    go path: /usr/lib/go-1.24
    go path: /usr/lib/go-1.24/bin/go
    go version: go version go1.24.4 linux/amd64
    go version: go version go1.24.4 linux/amd64

After v0.3.31, local-cypher reports cleanly:

    go path: /usr/lib/go-1.24/bin/go
    go version: go version go1.24.4 linux/amd64
    sqlite3 path: /usr/bin/sqlite3
    sqlite3 version: 3.46.1 ...

The profile remains functionally the same; this was output polish only.

### 4.5 README/docs/roadmap update landed

New/updated docs at v0.3.31:

    carbonstack/docs/150-local-cypher-polish-generated-cleanup-v0.md
    carbonstack/docs/README.md
    carbonstack/roadmap/ROADMAP.md

Mainline docs commit:

    7894278 docs: record local-cypher polish and cleanup control

The docs record the recon findings, cleanup flag behavior, safety boundary, validation results, and v0.4.0 release framing direction.

### 4.6 Final validation passed with no manual cleanup

Final WSL Debian validation sequence included:

    go run . --profile local-cypher
    go run . --profile doctor
    go run . --profile core --clean-generated

Final WSL Debian repo snapshot:

    carbonstack        7894278 docs: record local-cypher polish and cleanup control
    carbonstack-comms  012c8bf scripts: support source snapshot artifact guard
    carbonstack-cypher 9ab994c docs: point to local operator runbook
    carbonstack-os     1bbbe52 docs: clarify CarbonStackOS target direction

Final validation passed under WSL Debian. `local-cypher` passed without generating source-tree artifacts. `doctor` passed. `core --clean-generated` passed OpenMLS real-Cypher lifecycle, Comms package tests, and Cypher package tests, then removed the expected OpenMLS generated roots through the explicit cleanup flag. Final clean status showed no dirty files across all four repos.

---

## 5. Critical Blunders / Lessons Preserved

### 5.1 The first v0.3.31 patch partially applied and failed compile

The important blunder:

    cleanGenerated was declared but not used.

Observed error:

    ./main.go:49:2: declared and not used: cleanGenerated

Root cause:

    patch insertion hit the flag declaration and Runner field but missed the wiring line and cleanup method/call insertion.

Correct lesson:

    always compile the runner immediately after scripted patching.
    prefer narrower deterministic string replacements for known file contents.
    inspect `main.go` after patch scripts when adding flags, dispatch, or post-run behavior.
    this was a compile-time patching blunder, not a validation failure.

### 5.2 Cleanup must remain explicit and narrow

The cleanup flag solves workflow drag, but it is deliberately not general cleanup.

Correct lesson:

    --clean-generated is allowed because it only deletes known generated roots.
    it must not become a broad untracked-file remover.
    it must not touch manual local operator DBs.
    it must not replace artifact scans.
    it should not become default behavior without another decision rung.

### 5.3 local-cypher remains clean without cleanup

Recon confirmed `local-cypher` itself leaves no OpenMLS sidecar generated roots and no dirty repo state.

Correct lesson:

    local-cypher remains a small Cypher-only validation profile.
    core is the profile that creates OpenMLS sidecar generated roots.
    cleanup flag mainly helps core / release-adjacent validation workflow.

### 5.4 Output polish was appropriate here

The local-cypher toolchain output had become distracting and was easy to fix. It was safe to include with cleanup-control work because it touched the same runner polish surface and did not alter validation semantics.

Correct lesson:

    small output polish is acceptable when it reduces confusion and does not change claims.
    do not bundle broader semantic changes into such a polish rung.

### 5.5 v0.4.0 framing direction is now captured

v0.3.31 explicitly captures the user-approved release framing direction:

    v0.4.0 should be a broad local deployability pre-release.
    It is a milestone / research-and-development release.
    It is not for public-user or application use until testing, integration, and maturity.

Correct lesson:

    release names are claims.
    v0.4.0 should be broad enough to describe the local deployability milestone, but cautious enough to avoid implying real secure app readiness.

---

## 6. Current Blockers / Not Validated

Current blockers and nonvalidated areas:

    local-cypher now exists and is polished, but only implements positive Cypher-only lifecycle validation.
    Negative-path local-cypher validation is not implemented yet.
    local-backbone runner profile does not exist and remains reserved for later whole-stack validation.
    No helper command exists yet.
    Local-only operator convention is documented and proof-tested, but not enforced by code outside local-cypher.
    Cypher default bind remains :8080; local-only docs recommend explicit 127.0.0.1:8080.
    Runtime Comms send/inbox remains stub-era and not wired as user-facing OpenMLS messenger UX.
    Production-safe provider storage and secure vault design are unsolved.
    Hostile-server rollback/replay/metadata-abuse validation remains future work.
    Android/CarbonStackOS implementation remains later work.
    CarbonStack Relay Space mechanics remain a defined future concept, not an implementation.
    Public ingress, systemd, cloudflared, real homelab deployment, and production service paths remain out of scope.
    Temporary DB validation remains dev/test-only and does not replace later persistent user/operator-state validation for public Comms maturity.
    --clean-generated is intentionally narrow and does not replace release-snapshot strict pre-test validation.

Hard nonclaims remain:

    Do not claim production readiness.
    Do not claim production E2EE.
    Do not claim hostile-server safety.
    Do not claim metadata privacy.
    Do not claim Debian production deployability.
    Do not claim systemd readiness.
    Do not claim cloudflared readiness.
    Do not claim Android readiness.
    Do not claim secure vault/storage.
    Do not claim external audit or certification.
    Do not claim CarbonStack Relay Space mechanics are implemented.
    Do not claim local-backbone exists or is justified yet.
    Do not claim local-cypher validates runtime Comms UX.
    Do not claim local-cypher is sufficient for public user testing maturity.
    Do not claim --clean-generated is a general cleanup tool.
    Do not claim v0.3.31 is a public release.

Updated allowed claim:

    CarbonStack v0.3.31 polishes the local-cypher Go runner profile and adds an explicit opt-in --clean-generated flag.
    The cleanup flag runs after successful profile execution and removes only known OpenMLS sidecar generated roots.
    local-cypher, doctor, and core --clean-generated passed under WSL Debian after the v0.3.31 docs commit.
    Final clean status showed no dirty files across the four repos without manual rm -rf cleanup.

---

## 7. Next Safest Actions

Recommended next sequence after this v0.3.31 breakpoint:

    v0.3.32 preferred:
      single negative-path local-cypher validation, starting with the historically grounded unsupported_protocol_version case

Recommended concrete next rung:

    1. Recon current local-cypher code and Cypher HTTP tests for the invalid content/protocol pair.
    2. Add one negative-path lifecycle check to local-cypher after the positive envelope path or before it with isolated payload.
    3. Use content_type=carbonstack.message.text.stub.v0 with protocol_version=carbonstack-openmls-sidecar-v0.
    4. Expect HTTP 400.
    5. Verify the response contains unsupported_protocol_version if the current API exposes that code reliably.
    6. Validate go test ./... for the runner.
    7. Validate local-cypher.
    8. Validate doctor.
    9. Validate core --clean-generated.
    10. Document result as v0.3.32.
    11. Keep local-backbone deferred.
    12. Keep runtime Comms UX deferred to v0.4.x.

Alternative if implementation looks risky:

    v0.3.32:
      local-cypher negative-path validation contract/recon doc

    v0.3.33:
      implement single negative-path validation

Avoid next:

    local-backbone runner profile
    helper implementation
    public ingress
    cloudflared
    systemd
    real homelab deployment
    runtime Comms send/inbox UX
    Android app
    CarbonStackOS
    CarbonStack Relay Space implementation
    production release

---

## 8. Critical Paths / Functions

### WSL Debian active working umbrella

    [WSL-DEBIAN] umbrella: ~/repos/carbonstack_umbrella
    [WSL-DEBIAN] carbonstack: ~/repos/carbonstack_umbrella/carbonstack
    [WSL-DEBIAN] carbonstack-comms: ~/repos/carbonstack_umbrella/carbonstack-comms
    [WSL-DEBIAN] carbonstack-cypher: ~/repos/carbonstack_umbrella/carbonstack-cypher
    [WSL-DEBIAN] carbonstack-os: ~/repos/carbonstack_umbrella/carbonstack-os
    [GO-RUNNER] runner dir: ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate

### WSL Debian known-good validation commands

From runner dir:

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go run . --profile local-cypher
    go run . --profile doctor
    go run . --profile core --clean-generated

From carbonstack repo root:

    cd ~/repos/carbonstack_umbrella/carbonstack
    go run ./tools/carbonstack-validate --profile local-cypher
    go run ./tools/carbonstack-validate --profile doctor
    go run ./tools/carbonstack-validate --profile core --clean-generated

### v0.3.31 docs

    [DOC] v0.3.31 polish/cleanup result:
    carbonstack/docs/150-local-cypher-polish-generated-cleanup-v0.md

    [DOC] v0.3.30 implementation result:
    carbonstack/docs/149-local-cypher-runner-implementation-v0.md

    [DOC] v0.3.29 contract:
    carbonstack/docs/148-local-cypher-validation-contract-v0.md

    [DOC] docs index:
    carbonstack/docs/README.md

    [DOC] roadmap:
    carbonstack/roadmap/ROADMAP.md

### Go runner surfaces after v0.3.31

    [RUNNER] directory:
    carbonstack/tools/carbonstack-validate

    [RUNNER] implementation files:
    carbonstack/tools/carbonstack-validate/main.go
    carbonstack/tools/carbonstack-validate/local_cypher.go
    carbonstack/tools/carbonstack-validate/checksums.go
    carbonstack/tools/carbonstack-validate/release_snapshot.go
    carbonstack/tools/carbonstack-validate/README.md
    carbonstack/tools/carbonstack-validate/go.mod

    [RUNNER] profiles:
    doctor
    core
    local-cypher
    full
    release-snapshot
    write-checksums
    verify-checksums

    [RUNNER] cleanup flag:
    --clean-generated

### --clean-generated behavior

    Runs only after successful profile execution.
    Deletes only:
      carbonstack-comms/internal/protocol/mls/openmls-sidecar/.carbonstack-openmls-sidecar-state
      carbonstack-comms/internal/protocol/mls/openmls-sidecar/target

    Does not delete:
      arbitrary untracked files
      manual local operator DBs
      $HOME/.local/share/carbonstack/cypher/cypher.db
      provider-storage.json unless added by a future deliberate expansion
      signer.json unless added by a future deliberate expansion

    Does not replace:
      ArtifactScan
      StrictPreTestArtifactScan
      release-snapshot pre-test hygiene checks

### local-cypher lifecycle

    start:
      build temporary Cypher binary
      create temporary DB directory
      choose dynamic loopback port
      set CYPHER_ADDR to 127.0.0.1:<dynamic_port>
      set CYPHER_DB to temporary SQLite DB path
      set CYPHER_MIGRATIONS to source-tree migrations
      set CYPHER_DEV_INVITE to known temporary dev invite
      start Cypher
      wait for GET /v0/health

    account/device:
      POST /v0/invites/claim
      POST /v0/dev/invites
      POST /v0/invites/claim
      POST /v0/devices/register
      GET /v0/accounts/<account_id>/devices

    envelope:
      POST /v0/envelopes
      GET /v0/devices/<device_id>/envelopes
      POST /v0/envelopes/<envelope_id>/ack
      GET /v0/devices/<device_id>/envelopes

    restart:
      stop Cypher
      restart against same temp DB
      GET /v0/health
      GET /v0/accounts/<account_id>/devices
      GET /v0/devices/<device_id>/envelopes

    cleanup:
      stop process
      remove temp binary
      remove temp DB / proof dir
      leave no source-tree artifacts

### Accepted envelope pair

    content_type=carbonstack.mls.application-message.v0
    protocol_version=carbonstack-openmls-sidecar-v0

### Future negative path candidate

    content_type=carbonstack.message.text.stub.v0
    protocol_version=carbonstack-openmls-sidecar-v0

Expected result:

    HTTP 400
    error code unsupported_protocol_version

### Critical functions / behavior preserved from v0.3.24

    Store.Migrate:
      ensures schema_migrations exists;
      reads migration .sql files in sorted order;
      computes sha256 checksum for each migration;
      skips matching already-applied migrations;
      fails on checksum mismatch for already-applied migration;
      applies new migrations in a transaction;
      records migration after successful application.

    TestMigrateRecordsAppliedMigrations:
      proves applied migrations are recorded with non-empty checksum/applied_at.

    TestMigrateCanRunTwiceAgainstSameDB:
      proves repeated migration against the same DB now succeeds.

    TestMigrateDetectsAppliedMigrationChecksumMismatch:
      proves already-applied migration drift hard-fails.

---

## 9. Lean Breakpoint Summary

    CarbonStack v0.3.31 is a stable mainline runner-polish and cleanup-control breakpoint after v0.3.30. It polishes the implemented local-cypher Go runner profile, adds the explicit opt-in --clean-generated flag, documents the cleanup safety boundary, records v0.4.0 as a broad local deployability pre-release / milestone / research-and-development release rather than public-user or application-use readiness, validates local-cypher, doctor, and core --clean-generated under WSL Debian, and reaches clean final status across all four repos without manual rm -rf cleanup. local-backbone, helper tooling, negative-path validation, runtime Comms UX, CarbonStack Relay Space mechanics, public ingress, systemd, cloudflared, homelab, Android, CarbonStackOS, hostile-server proof, secure vault, persistent user/operator-state maturity, and production/security claims remain deferred.

---

## 10. Preserved v0.3.30 Operational Process Log

The following section preserves the previous v0.3.30 LogDoc body as operational continuity. It is intentionally retained rather than compressed away. Where v0.3.30 text conflicts with the v0.3.31 current-state overlay above, v0.3.31 wins for current state; v0.3.30 remains the provenance/process ledger for the local-cypher runner implementation and the preserved v0.3.29 / v0.3.28 / v0.3.27 / v0.3.26 / v0.3.25 / v0.3.24 / v0.3.23 / v0.3.22 / v0.3.21 / v0.3.20 history.



# CarbonStack LogDoc v0.3.30

**Last updated:** 2026-06-02 local session  
**Owner/operator:** bitcrusher32  
**Assistant/model:** GPT-5.5 Thinking  
**Current phase:** Phase 2E-prep / v0.3.x **local-cypher runner implementation checkpoint**. `carbonstack` is now at `15a3758 docs: record local-cypher runner implementation`, after `v0.3.30` first landed the runner implementation at `a98ee9f runner: add local-cypher`. `carbonstack-cypher` remains at `9ab994c docs: point to local operator runbook`; `carbonstack-comms` remains at `012c8bf scripts: support source snapshot artifact guard`; `carbonstack-os` remains at `1bbbe52 docs: clarify CarbonStackOS target direction`. The live WSL Debian working umbrella at `~/repos/carbonstack_umbrella` passed `local-cypher`, `doctor`, and `core` after the v0.3.30 docs commit, expected OpenMLS sidecar artifacts were cleaned afterward, and final clean status showed no dirty files across the four repos.

**Version schema:** `v[scope].[timeline]`. This file is `v0.3.30`, the local-cypher Go runner implementation handoff after the `v0.3.29` local-cypher validation contract recon checkpoint. Per LogDoc V2 practice, this Markdown file intentionally preserves verbose operational continuity; the JSON breakpoint remains the lean handoff state. The dense v0.3.29, v0.3.28, v0.3.27, v0.3.26, v0.3.25, v0.3.24, v0.3.23, v0.3.22, v0.3.21, and v0.3.20 operational process logs are preserved later in this file rather than compressed away.

**LogDoc V2 usage note:** LogDoc V2 uses a split between a Markdown LogDoc for working-session continuity, a JSON state file for end-of-session handoff, Git for history preservation, and README/docs for public claims. This v0.3.30 Markdown file therefore intentionally stays “bloated” where operational process, blunders, validation ladders, and historical context matter. The v0.3.30 JSON should remain lean and should agree with this file on the latest current state.

---

## 1. Project Goal

**Active goal:** Preserve CarbonStack state after implementing the first `local-cypher` Go runner profile and recording the implementation result. The immediate next line of work should be proof/result hardening, review, and possibly negative-path planning, not `local-backbone`, not helper tooling, not runtime Comms UX, not public ingress, not systemd/cloudflared, not real homelab deployment, and not Android/CarbonStackOS work.

CarbonStack’s broader project intent remains:

- **CarbonStack** is the secure-communications appliance-stack project and shared doctrine/release authority.
- **CarbonStackOS** is the future constrained Android-derived appliance OS.
- **CarbonStackComms** is the text-first encrypted messaging client and current OpenMLS sidecar/relay integration surface.
- **CarbonStackCypher** is the self-hostable relay/storage stack.
- The `carbonstack` repo remains the doctrine/spec/docs/release authority; implementation repos must not contradict it without revising the spec.

**North star:**

> Its goal is not to be a better smartphone. Its goal is to stop being a smartphone.

**Core doctrine:**

> Every feature is guilty until it proves it does not add unacceptable parser, network, sensor, identity, or filesystem authority.

**Current reality after v0.3.30:** `v0.3.20` remains the current public runner-backed testing release. `v0.3.21` defined the local-only deployability model. `v0.3.22` narrowed the local deployability surface around Cypher operator config/API/database assumptions and adopted **CarbonStack Relay Space** terminology. `v0.3.23` proved the SQLite repeated-migration persistence hazard. `v0.3.24` implemented `schema_migrations` tracking in CarbonStackCypher. `v0.3.25` recorded the first local-only Cypher operator runbook skeleton. `v0.3.26` recorded the local operator config/data convention. `v0.3.27` proved a real local-only Cypher explicit-env API lifecycle. `v0.3.28` recorded the helper/runner decision recon and reserved `local-backbone` for later whole-stack validation. `v0.3.29` defined the future `local-cypher` validation contract. `v0.3.30` now implements that contract as the first `local-cypher` profile in the Go runner, with a runner-owned temporary Cypher binary, temporary isolated SQLite DB, dynamic loopback port, positive invite/device/envelope/ack lifecycle, restart against the same DB, persisted-state checks, cleanup, and artifact scans.

The current public release page remains:

    https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.3.20

The current public testing interpretation remains:

    v0.3.20 is the preferred public runner-backed testing release.
    v0.3.4 remains available as the older Windows 11 / PowerShell legacy testbed.
    Debian / WSL Debian is now the mainline public dev/test validation direction.
    Windows 11 short-path validation was the final explicit Windows dev/test validation for this phase.

This is still **not** a deployability release. It does **not** prove production readiness, production E2EE, hostile-server safety, metadata privacy, Android readiness, secure vault/storage, Debian production deployability, systemd readiness, cloudflared readiness, public ingress safety, real homelab validation, external audit, or certification.

---

## 2. Current Repo Heads

    carbonstack        15a3758 (HEAD -> main, origin/main, origin/HEAD) docs: record local-cypher runner implementation
    carbonstack-comms  012c8bf (HEAD -> main, origin/main, origin/HEAD) scripts: support source snapshot artifact guard
    carbonstack-cypher 9ab994c (HEAD -> main, origin/main, origin/HEAD) docs: point to local operator runbook
    carbonstack-os     1bbbe52 (HEAD -> main, origin/main, origin/HEAD) docs: clarify CarbonStackOS target direction

Current `carbonstack` public runner-backed testing release tag:

    v0.3.20 -> 3eeb1a1e1a / 3eeb1a1 docs: clean public release surface

Previous mainline deployability checkpoints:

    v0.3.30 runner  -> a98ee9f runner: add local-cypher
    v0.3.29 mainline -> 37a6d2a docs: record local-cypher validation contract
    v0.3.28 mainline -> af33139 docs: record local operator helper runner decision recon
    v0.3.27 mainline -> e516fc7 docs: record local Cypher API lifecycle proof
    v0.3.26 mainline -> 6e14fe0 docs: record local operator config data convention
    v0.3.25 mainline -> a5c6351 docs: update readme, roadmap, and historical doc
    v0.3.25 cypher  -> 9ab994c docs: point to local operator runbook
    v0.3.24 mainline -> 76f0258 docs: record Cypher schema migrations implementation
    v0.3.24 cypher  -> 42f838f db: track applied migrations
    v0.3.23 mainline -> 1fef7fe docs: record Cypher migration persistence recon
    v0.3.22 mainline -> 328ffc4 docs: record Cypher local operator surface recon
    v0.3.21 mainline -> 6850c58 docs: record local backbone deployability recon

Legacy `carbonstack` Windows/PowerShell testbed release tag:

    v0.3.4 -> 7c5b7e6 docs: record release snapshot self-test fix

Historical `carbonstack` release tag:

    v0.3.0 -> 59245fc docs: freeze v0.3.0 release packaging plan

Important continuity note: `v0.3.30` does not supersede `v0.3.20` as a public release. It advances mainline local-only deployability validation by implementing the Cypher-only `local-cypher` runner profile. It is safe to pause here because the runner implementation and docs/result record were pushed, WSL Debian `local-cypher` / `doctor` / `core` passed after the docs commit, generated artifacts were cleaned, and final status was clean across all four repos.

---

## 3. Current Validated State

Validated at the v0.3.30 checkpoint:

    [WSL-DEBIAN] working umbrella exists at ~/repos/carbonstack_umbrella
    [WSL-DEBIAN] carbonstack is at 15a3758
    [WSL-DEBIAN] carbonstack-comms is at 012c8bf
    [WSL-DEBIAN] carbonstack-cypher is at 9ab994c
    [WSL-DEBIAN] carbonstack-os is at 1bbbe52
    [CARBONSTACK] runner commit a98ee9f added the local-cypher profile
    [CARBONSTACK] docs/149 local-cypher runner implementation result committed and pushed
    [GO-RUNNER] local-cypher passed after the v0.3.30 docs commit
    [GO-RUNNER] doctor passed after the v0.3.30 docs commit
    [GO-RUNNER] core passed after the v0.3.30 docs commit
    [ARTIFACT CLEANUP] expected OpenMLS generated roots were removed after validation
    [FINAL STATUS] all four repos were clean after cleanup and commit

Observed WSL Debian toolchain baseline:

    git version 2.47.3
    go version go1.24.4 linux/amd64
    rustc 1.96.0 (ac68faa20 2026-05-25)
    cargo 1.96.0 (30a34c682 2026-05-25)
    sqlite3 3.46.1

Observed `local-cypher` validation result during the v0.3.30 flow:

    required paths: PASS
    pre-local-cypher artifact scan: PASS / no generated/private/build artifact hits
    temporary Cypher binary build: PASS
    first health check: PASS
    seeded Alice invite claim: PASS
    Bob dev invite creation: PASS
    Bob invite claim: PASS
    Alice device registration: PASS
    Bob device registration: PASS
    Alice device listing: PASS
    opaque OpenMLS application-message envelope submit: PASS
    Bob inbox retrieval and payload metadata verification: PASS
    envelope ack: PASS
    Bob inbox after ack empty: PASS
    first Cypher process stopped with expected SIGINT termination
    restart against same temporary DB: PASS
    restart health check: PASS
    persisted Alice device state after restart: PASS
    Bob's acked inbox remains empty after restart: PASS
    second Cypher process stopped with expected SIGINT termination
    post-local-cypher artifact scan: PASS / no generated/private/build artifact hits
    VALIDATION PASSED

Observed `core` validation result during the v0.3.30 flow:

    pre-test artifact scan: PASS / no generated/private/build artifact hits
    OpenMLS real-Cypher lifecycle: PASS
    carbonstack-comms package tests: PASS
    carbonstack-cypher package tests: PASS
    post-test artifact scan: expected OpenMLS sidecar generated roots only
    VALIDATION PASSED

Expected post-test generated roots from `core`:

    carbonstack-comms/internal/protocol/mls/openmls-sidecar/.carbonstack-openmls-sidecar-state
    carbonstack-comms/internal/protocol/mls/openmls-sidecar/target

These roots are expected only after tests and must not be committed or treated as source artifacts. They were cleaned before the final clean status snapshot.

---

## 4. v0.3.30 Work Completed

### 4.1 local-cypher runner implementation landed

New/updated runner files at v0.3.30:

    carbonstack/tools/carbonstack-validate/local_cypher.go
    carbonstack/tools/carbonstack-validate/main.go
    carbonstack/tools/carbonstack-validate/README.md

Implementation commit:

    a98ee9f runner: add local-cypher

The runner now accepts:

    go run . --profile local-cypher

The implementation intentionally lives in a new `local_cypher.go` file rather than bloating `main.go`. `main.go` was only patched for profile dispatch/help text.

### 4.2 local-cypher behavior implemented

`local-cypher` now:

    checks required paths
    reports basic toolchain information
    runs a pre-local-cypher artifact scan
    creates a temporary directory under /tmp
    builds a temporary Cypher binary
    creates a temporary isolated SQLite DB
    chooses a dynamic loopback port
    starts Cypher on 127.0.0.1 only
    waits for GET /v0/health
    claims the seeded dev invite
    creates a second dev invite
    claims the second invite
    registers Alice/Bob devices
    lists Alice devices
    submits an opaque OpenMLS application-message envelope
    retrieves Bob's inbox and verifies payload metadata
    acks the envelope
    verifies Bob's inbox is empty after ack
    stops Cypher
    restarts Cypher against the same temporary DB
    verifies health after restart
    verifies persisted Alice device state after restart
    verifies Bob's acked inbox remains empty after restart
    stops Cypher again
    removes temporary state
    runs a post-local-cypher artifact scan

Accepted envelope pair used by the profile:

    content_type=carbonstack.mls.application-message.v0
    protocol_version=carbonstack-openmls-sidecar-v0

The profile remains Cypher-only. It does not call itself `local-backbone`, does not validate runtime Comms UX, does not expose public ingress, and does not make production/security claims.

### 4.3 local-cypher docs/result record landed

New/updated docs at v0.3.30:

    carbonstack/docs/149-local-cypher-runner-implementation-v0.md
    carbonstack/docs/README.md
    carbonstack/roadmap/ROADMAP.md

Docs commit:

    15a3758 docs: record local-cypher runner implementation

The result doc records the implementation files, validation output, and scope boundary. It explicitly preserves the nonclaim that this is not `local-backbone`, not runtime Comms UX, not public ingress, not production deployability, and not negative-path validation yet.

### 4.4 Final validation and cleanup completed

Final WSL Debian validation sequence included:

    go run . --profile local-cypher
    go run . --profile doctor
    go run . --profile core

Final WSL Debian repo snapshot:

    carbonstack        15a3758 docs: record local-cypher runner implementation
    carbonstack-comms  012c8bf scripts: support source snapshot artifact guard
    carbonstack-cypher 9ab994c docs: point to local operator runbook
    carbonstack-os     1bbbe52 docs: clarify CarbonStackOS target direction

Final validation passed under WSL Debian. `local-cypher` passed without generating source-tree artifacts. `core` passed OpenMLS real-Cypher lifecycle, Comms package tests, and Cypher package tests, including `internal/db` and `internal/httpapi` package tests. After validation, expected OpenMLS generated roots were removed from `carbonstack-comms`, and final clean status showed no dirty files across the four repos.

---

## 5. Critical Blunders / Lessons Preserved

### 5.1 v0.3.30 correctly followed contract-before-code

The implementation followed the v0.3.29 contract closely instead of inventing a new runner shape in-place.

Correct lesson:

    recon -> contract docs -> implementation -> proof/result docs -> breakpoint

This is exactly the workflow LogDoc is meant to preserve.

### 5.2 Dynamic loopback port decision worked

The runner chose a dynamic `127.0.0.1` port during validation rather than using a fixed high port. The observed run used:

    127.0.0.1:44373

Correct lesson:

    dynamic loopback port selection is viable for local-cypher.
    fixed port conflicts are avoidable without adding much UX surface.
    local-cypher must remain loopback-only.

### 5.3 Temporary state cleanup worked

The runner used a temporary path:

    /tmp/carbonstack-local-cypher-87725335

and `local-cypher`'s pre/post artifact scans reported no generated/private/build artifact hits.

Correct lesson:

    local-cypher can validate Cypher-only lifecycle without touching the source tree or manual operator DB.
    this keeps dev/test validation repeatable and non-destructive.

### 5.4 Expected process termination messages are acceptable

`local-cypher` stops the Cypher process with an interrupt signal and reports:

    INFO: first Cypher process stopped with expected termination: signal: interrupt
    INFO: second Cypher process stopped with expected termination: signal: interrupt

Correct lesson:

    expected SIGINT/SIGTERM-style shutdown should be treated as normal runner behavior, not a failure.
    if future cleanup hardening changes process management, preserve this distinction.

### 5.5 Cosmetic toolchain-reporting rough edge observed

The `local-cypher` toolchain block currently prints duplicated/confusing Go path/version lines, including a `go version:` label paired with `go env GOROOT` output.

Observed rough shape:

    go path: /usr/lib/go-1.24/bin/go
    go version: /usr/lib/go-1.24
    go path: /usr/lib/go-1.24
    go path: /usr/lib/go-1.24/bin/go
    go version: go version go1.24.4 linux/amd64
    go version: go version go1.24.4 linux/amd64

This is not a validation failure, but it is worth cleaning later if it becomes distracting.

Correct lesson:

    local-cypher output is functionally valid, but its toolchain-reporting section could be tightened in a later polish rung.
    do not mix cosmetic output cleanup with negative-path or release-surface changes unless deliberately scoped.

### 5.6 Artifact cleanup cadence held correctly

The final validation produced expected sidecar generated roots from `core`, then cleanup was run. Final repo status snapshot showed clean status after cleanup. Continue this pattern after every `core` run.

---

## 6. Current Blockers / Not Validated

Current blockers and nonvalidated areas:

    local-cypher now exists, but only implements positive Cypher-only lifecycle validation.
    Negative-path local-cypher validation is not implemented yet.
    local-backbone runner profile does not exist and remains reserved for later whole-stack validation.
    No helper command exists yet.
    Local-only operator convention is documented and proof-tested, but not enforced by code outside local-cypher.
    Cypher default bind remains :8080; local-only docs recommend explicit 127.0.0.1:8080.
    Runtime Comms send/inbox remains stub-era and not wired as user-facing OpenMLS messenger UX.
    Production-safe provider storage and secure vault design are unsolved.
    Hostile-server rollback/replay/metadata-abuse validation remains future work.
    Android/CarbonStackOS implementation remains later work.
    CarbonStack Relay Space mechanics remain a defined future concept, not an implementation.
    Public ingress, systemd, cloudflared, real homelab deployment, and production service paths remain out of scope.
    Temporary DB validation remains dev/test-only and does not replace later persistent user/operator-state validation for public Comms maturity.
    local-cypher toolchain output is slightly noisy/cosmetic and can be cleaned later.

Hard nonclaims remain:

    Do not claim production readiness.
    Do not claim production E2EE.
    Do not claim hostile-server safety.
    Do not claim metadata privacy.
    Do not claim Debian production deployability.
    Do not claim systemd readiness.
    Do not claim cloudflared readiness.
    Do not claim Android readiness.
    Do not claim secure vault/storage.
    Do not claim external audit or certification.
    Do not claim CarbonStack Relay Space mechanics are implemented.
    Do not claim local-backbone exists or is justified yet.
    Do not claim local-cypher validates runtime Comms UX.
    Do not claim local-cypher is sufficient for public user testing maturity.
    Do not claim v0.3.30 is a public release.

Updated allowed claim:

    CarbonStack v0.3.30 implements the first local-cypher Go runner profile.
    The profile is Cypher-only and runner-owned.
    It builds a temporary Cypher binary, uses a temporary isolated SQLite DB, binds to a dynamic loopback-only port, runs the positive invite/device/envelope/ack lifecycle, restarts against the same DB, checks persisted device state and acked inbox behavior, and cleans temporary state.
    WSL Debian local-cypher, doctor, and core validation passed after the v0.3.30 docs commit.
    Final clean status showed no dirty files across the four repos after expected artifact cleanup.

---

## 7. Next Safest Actions

Recommended next sequence after this v0.3.30 breakpoint:

    v0.3.31 preferred:
      local-cypher proof/result hardening checkpoint or negative-path planning doc

Recommended concrete next rung:

    1. Review local-cypher output and decide whether toolchain-reporting cleanup should be immediate or deferred.
    2. Decide whether negative-path validation should be a docs/recon rung first or an implementation patch.
    3. Candidate negative-path test: invalid content/protocol pair should return unsupported_protocol_version.
    4. Keep helper tooling deferred.
    5. Keep local-backbone deferred.
    6. Keep runtime Comms UX deferred to v0.4.x.
    7. Continue validating local-cypher, doctor, and core after runner changes.
    8. Continue cleaning OpenMLS generated roots after core.

Alternative next sequence if polishing before negative paths:

    v0.3.31:
      local-cypher output polish / toolchain report cleanup

    v0.3.32:
      local-cypher negative-path validation contract/recon

    v0.3.33:
      implement local-cypher negative-path validation

Avoid next:

    local-backbone runner profile
    helper implementation
    public ingress
    cloudflared
    systemd
    real homelab deployment
    runtime Comms send/inbox UX
    Android app
    CarbonStackOS
    CarbonStack Relay Space implementation
    production release

---

## 8. Critical Paths / Functions

### WSL Debian active working umbrella

    [WSL-DEBIAN] umbrella: ~/repos/carbonstack_umbrella
    [WSL-DEBIAN] carbonstack: ~/repos/carbonstack_umbrella/carbonstack
    [WSL-DEBIAN] carbonstack-comms: ~/repos/carbonstack_umbrella/carbonstack-comms
    [WSL-DEBIAN] carbonstack-cypher: ~/repos/carbonstack_umbrella/carbonstack-cypher
    [WSL-DEBIAN] carbonstack-os: ~/repos/carbonstack_umbrella/carbonstack-os
    [GO-RUNNER] runner dir: ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate

### WSL Debian known-good validation commands

From runner dir:

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go run . --profile local-cypher
    go run . --profile doctor
    go run . --profile core

From carbonstack repo root:

    cd ~/repos/carbonstack_umbrella/carbonstack
    go run ./tools/carbonstack-validate --profile local-cypher
    go run ./tools/carbonstack-validate --profile doctor
    go run ./tools/carbonstack-validate --profile core

### Artifact cleanup after validation

    cd ~/repos/carbonstack_umbrella/carbonstack-comms
    rm -rf internal/protocol/mls/openmls-sidecar/.carbonstack-openmls-sidecar-state
    rm -rf internal/protocol/mls/openmls-sidecar/target

### v0.3.30 docs

    [DOC] v0.3.30 implementation result:
    carbonstack/docs/149-local-cypher-runner-implementation-v0.md

    [DOC] v0.3.29 contract:
    carbonstack/docs/148-local-cypher-validation-contract-v0.md

    [DOC] docs index:
    carbonstack/docs/README.md

    [DOC] roadmap:
    carbonstack/roadmap/ROADMAP.md

### Go runner surfaces after v0.3.30

    [RUNNER] directory:
    carbonstack/tools/carbonstack-validate

    [RUNNER] implementation files:
    carbonstack/tools/carbonstack-validate/main.go
    carbonstack/tools/carbonstack-validate/local_cypher.go
    carbonstack/tools/carbonstack-validate/checksums.go
    carbonstack/tools/carbonstack-validate/release_snapshot.go
    carbonstack/tools/carbonstack-validate/README.md
    carbonstack/tools/carbonstack-validate/go.mod

    [RUNNER] profiles:
    doctor
    core
    local-cypher
    full
    release-snapshot
    write-checksums
    verify-checksums

### local-cypher lifecycle

    start:
      build temporary Cypher binary
      create temporary DB directory
      choose dynamic loopback port
      set CYPHER_ADDR to 127.0.0.1:<dynamic_port>
      set CYPHER_DB to temporary SQLite DB path
      set CYPHER_MIGRATIONS to source-tree migrations
      set CYPHER_DEV_INVITE to known temporary dev invite
      start Cypher
      wait for GET /v0/health

    account/device:
      POST /v0/invites/claim
      POST /v0/dev/invites
      POST /v0/invites/claim
      POST /v0/devices/register
      GET /v0/accounts/<account_id>/devices

    envelope:
      POST /v0/envelopes
      GET /v0/devices/<device_id>/envelopes
      POST /v0/envelopes/<envelope_id>/ack
      GET /v0/devices/<device_id>/envelopes

    restart:
      stop Cypher
      restart against same temp DB
      GET /v0/health
      GET /v0/accounts/<account_id>/devices
      GET /v0/devices/<device_id>/envelopes

    cleanup:
      stop process
      remove temp binary
      remove temp DB / proof dir
      leave no source-tree artifacts

### Accepted envelope pair

    content_type=carbonstack.mls.application-message.v0
    protocol_version=carbonstack-openmls-sidecar-v0

### Future negative path candidate

    content_type=carbonstack.message.text.stub.v0
    protocol_version=carbonstack-openmls-sidecar-v0

Expected result:

    HTTP 400
    error code unsupported_protocol_version

### Critical functions / behavior preserved from v0.3.24

    Store.Migrate:
      ensures schema_migrations exists;
      reads migration .sql files in sorted order;
      computes sha256 checksum for each migration;
      skips matching already-applied migrations;
      fails on checksum mismatch for already-applied migration;
      applies new migrations in a transaction;
      records migration after successful application.

    TestMigrateRecordsAppliedMigrations:
      proves applied migrations are recorded with non-empty checksum/applied_at.

    TestMigrateCanRunTwiceAgainstSameDB:
      proves repeated migration against the same DB now succeeds.

    TestMigrateDetectsAppliedMigrationChecksumMismatch:
      proves already-applied migration drift hard-fails.

---

## 9. Lean Breakpoint Summary

    CarbonStack v0.3.30 is a stable mainline implementation breakpoint after v0.3.29. It implements the first local-cypher Go runner profile in carbonstack/tools/carbonstack-validate via runner commit a98ee9f and docs/result commit 15a3758. local-cypher is Cypher-only, runner-owned, temporary-DB, dynamic-loopback, positive-lifecycle validation. It builds a temporary Cypher binary, starts Cypher on 127.0.0.1, claims invites, registers devices, submits/retrieves/acks an opaque OpenMLS application-message envelope, restarts against the same DB, verifies persisted device state and acked inbox behavior, cleans temporary state, and performs artifact scans. WSL Debian local-cypher, doctor, and core passed after the docs commit, expected OpenMLS generated roots were cleaned, and final status was clean across all four repos. local-backbone, helper tooling, negative-path validation, runtime Comms UX, CarbonStack Relay Space mechanics, public ingress, systemd, cloudflared, homelab, Android, CarbonStackOS, hostile-server proof, secure vault, and production/security claims remain deferred.

---

## 10. Preserved v0.3.29 Operational Process Log

The following section preserves the previous v0.3.29 LogDoc body as operational continuity. It is intentionally retained rather than compressed away. Where v0.3.29 text conflicts with the v0.3.30 current-state overlay above, v0.3.30 wins for current state; v0.3.29 remains the provenance/process ledger for the local-cypher validation contract and the preserved v0.3.28 / v0.3.27 / v0.3.26 / v0.3.25 / v0.3.24 / v0.3.23 / v0.3.22 / v0.3.21 / v0.3.20 history.



# CarbonStack LogDoc v0.3.29

**Last updated:** 2026-06-02 local session  
**Owner/operator:** bitcrusher32  
**Assistant/model:** GPT-5.5 Thinking  
**Current phase:** Phase 2E-prep / v0.3.x **local-cypher validation contract recon checkpoint**. `carbonstack` is now at `37a6d2a docs: record local-cypher validation contract`, after `v0.3.28` landed the local operator helper/runner decision recon at `af33139`. `carbonstack-cypher` remains at `9ab994c docs: point to local operator runbook`; `carbonstack-comms` remains at `012c8bf scripts: support source snapshot artifact guard`; `carbonstack-os` remains at `1bbbe52 docs: clarify CarbonStackOS target direction`. The live WSL Debian working umbrella at `~/repos/carbonstack_umbrella` passed `doctor` and `core` after the v0.3.29 docs commit, expected OpenMLS sidecar artifacts were cleaned afterward, and final clean status showed no dirty files across the four repos.

**Version schema:** `v[scope].[timeline]`. This file is `v0.3.29`, the local-cypher validation contract recon handoff after the `v0.3.28` local operator helper/runner decision recon checkpoint. Per LogDoc V2 practice, this Markdown file intentionally preserves verbose operational continuity; the JSON breakpoint remains the lean handoff state. The dense v0.3.28, v0.3.27, v0.3.26, v0.3.25, v0.3.24, v0.3.23, v0.3.22, v0.3.21, and v0.3.20 operational process logs are preserved later in this file rather than compressed away.

**LogDoc V2 usage note:** LogDoc V2 uses a split between a Markdown LogDoc for working-session continuity, a JSON state file for end-of-session handoff, Git for history preservation, and README/docs for public claims. This v0.3.29 Markdown file therefore intentionally stays “bloated” where operational process, blunders, validation ladders, and historical context matter. The v0.3.29 JSON should remain lean and should agree with this file on the latest current state.

---

## 1. Project Goal

**Active goal:** Preserve CarbonStack state after recording the future `local-cypher` validation contract. The immediate next line of work may be `v0.3.30` implementation of the `local-cypher` runner profile, but only if the v0.3.29 contract still looks correct after review. This is not `local-backbone`, not helper tooling, not runtime Comms UX, not public ingress, not systemd/cloudflared, not real homelab deployment, and not Android/CarbonStackOS work.

CarbonStack’s broader project intent remains:

- **CarbonStack** is the secure-communications appliance-stack project and shared doctrine/release authority.
- **CarbonStackOS** is the future constrained Android-derived appliance OS.
- **CarbonStackComms** is the text-first encrypted messaging client and current OpenMLS sidecar/relay integration surface.
- **CarbonStackCypher** is the self-hostable relay/storage stack.
- The `carbonstack` repo remains the doctrine/spec/docs/release authority; implementation repos must not contradict it without revising the spec.

**North star:**

> Its goal is not to be a better smartphone. Its goal is to stop being a smartphone.

**Core doctrine:**

> Every feature is guilty until it proves it does not add unacceptable parser, network, sensor, identity, or filesystem authority.

**Current reality after v0.3.29:** `v0.3.20` remains the current public runner-backed testing release. `v0.3.21` defined the local-only deployability model. `v0.3.22` narrowed the local deployability surface around Cypher operator config/API/database assumptions and adopted **CarbonStack Relay Space** terminology. `v0.3.23` proved the SQLite repeated-migration persistence hazard. `v0.3.24` implemented `schema_migrations` tracking in CarbonStackCypher. `v0.3.25` recorded the first local-only Cypher operator runbook skeleton. `v0.3.26` recorded the local operator config/data convention. `v0.3.27` proved a real local-only Cypher explicit-env API lifecycle. `v0.3.28` recorded the helper/runner decision recon and reserved `local-backbone` for later whole-stack validation. `v0.3.29` now records the future `local-cypher` validation contract: a Cypher-only, runner-owned, temporary-DB, explicit-loopback, positive invite/device/envelope/ack lifecycle profile with restart and cleanup requirements.

The current public release page remains:

    https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.3.20

The current public testing interpretation remains:

    v0.3.20 is the preferred public runner-backed testing release.
    v0.3.4 remains available as the older Windows 11 / PowerShell legacy testbed.
    Debian / WSL Debian is now the mainline public dev/test validation direction.
    Windows 11 short-path validation was the final explicit Windows dev/test validation for this phase.

This is still **not** a deployability release. It does **not** prove production readiness, production E2EE, hostile-server safety, metadata privacy, Android readiness, secure vault/storage, Debian production deployability, systemd readiness, cloudflared readiness, public ingress safety, real homelab validation, external audit, or certification.

---

## 2. Current Repo Heads

    carbonstack        37a6d2a (HEAD -> main, origin/main, origin/HEAD) docs: record local-cypher validation contract
    carbonstack-comms  012c8bf (HEAD -> main, origin/main, origin/HEAD) scripts: support source snapshot artifact guard
    carbonstack-cypher 9ab994c (HEAD -> main, origin/main, origin/HEAD) docs: point to local operator runbook
    carbonstack-os     1bbbe52 (HEAD -> main, origin/main, origin/HEAD) docs: clarify CarbonStackOS target direction

Current `carbonstack` public runner-backed testing release tag:

    v0.3.20 -> 3eeb1a1e1a / 3eeb1a1 docs: clean public release surface

Previous mainline deployability checkpoints:

    v0.3.28 mainline -> af33139 docs: record local operator helper runner decision recon
    v0.3.27 mainline -> e516fc7 docs: record local Cypher API lifecycle proof
    v0.3.26 mainline -> 6e14fe0 docs: record local operator config data convention
    v0.3.25 mainline -> a5c6351 docs: update readme, roadmap, and historical doc
    v0.3.25 cypher  -> 9ab994c docs: point to local operator runbook
    v0.3.24 mainline -> 76f0258 docs: record Cypher schema migrations implementation
    v0.3.24 cypher  -> 42f838f db: track applied migrations
    v0.3.23 mainline -> 1fef7fe docs: record Cypher migration persistence recon
    v0.3.22 mainline -> 328ffc4 docs: record Cypher local operator surface recon
    v0.3.21 mainline -> 6850c58 docs: record local backbone deployability recon

Legacy `carbonstack` Windows/PowerShell testbed release tag:

    v0.3.4 -> 7c5b7e6 docs: record release snapshot self-test fix

Historical `carbonstack` release tag:

    v0.3.0 -> 59245fc docs: freeze v0.3.0 release packaging plan

Important continuity note: `v0.3.29` does not supersede `v0.3.20` as a public release. It advances mainline local-only deployability reasoning after the v0.3.28 decision recon. It is safe to pause here because the `local-cypher` validation contract doc was pushed, WSL Debian `doctor`/`core` passed after the commit, generated artifacts were cleaned, and final status was clean across all four repos.

---

## 3. Current Validated State

Validated at the v0.3.29 checkpoint:

    [WSL-DEBIAN] working umbrella exists at ~/repos/carbonstack_umbrella
    [WSL-DEBIAN] carbonstack is at 37a6d2a
    [WSL-DEBIAN] carbonstack-comms is at 012c8bf
    [WSL-DEBIAN] carbonstack-cypher is at 9ab994c
    [WSL-DEBIAN] carbonstack-os is at 1bbbe52
    [CARBONSTACK] docs/148 local-cypher validation contract committed and pushed
    [GO-RUNNER] doctor passed after the v0.3.29 docs commit
    [GO-RUNNER] core passed after the v0.3.29 docs commit
    [ARTIFACT CLEANUP] expected OpenMLS generated roots were removed after validation
    [FINAL STATUS] all four repos were clean after cleanup and commit

Observed WSL Debian toolchain baseline:

    git version 2.47.3
    go version go1.24.4 linux/amd64
    rustc 1.96.0 (ac68faa20 2026-05-25)
    cargo 1.96.0 (30a34c682 2026-05-25)
    sqlite3 3.46.1

Observed `core` validation result during the v0.3.29 flow:

    pre-test artifact scan: PASS / no generated/private/build artifact hits
    OpenMLS real-Cypher lifecycle: PASS
    carbonstack-comms package tests: PASS
    carbonstack-cypher package tests: PASS
    post-test artifact scan: expected OpenMLS sidecar generated roots only
    VALIDATION PASSED

Expected post-test generated roots:

    carbonstack-comms/internal/protocol/mls/openmls-sidecar/.carbonstack-openmls-sidecar-state
    carbonstack-comms/internal/protocol/mls/openmls-sidecar/target

These roots are expected only after tests and must not be committed or treated as source artifacts. They were cleaned before the final clean status snapshot.

---

## 4. v0.3.29 Work Completed

### 4.1 WSL Debian runner/API scout completed before docs

Before writing the v0.3.29 doc, a read-only WSL Debian scout captured the actual runner and Cypher API surfaces rather than guessing.

The scout confirmed:

    carbonstack/tools/carbonstack-validate currently contains:
      main.go
      checksums.go
      release_snapshot.go
      README.md
      go.mod

    existing runner profiles are:
      doctor
      core
      full
      release-snapshot
      write-checksums
      verify-checksums

    profile dispatch is centralized in main.go
    the runner can infer umbrella roots or accept --root
    artifact scanning already exists
    core currently runs pre-test scan, OpenMLS real-Cypher lifecycle, Comms package tests, Cypher package tests, and post-test scan
    Cypher API tests already contain positive lifecycle and useful negative-path contract material
    current Cypher config defaults remain CYPHER_ADDR=:8080, CYPHER_DB=cypher.db, CYPHER_MIGRATIONS=migrations, and CYPHER_DEV_INVITE=dev-invite
    local operator docs still recommend explicit CYPHER_ADDR=127.0.0.1:8080

This preserved the project method:

    recon -> docs
    implementation -> patch
    validation -> cleanup -> breakpoint

### 4.2 local-cypher validation contract landed

New/updated docs at v0.3.29:

    carbonstack/docs/148-local-cypher-validation-contract-v0.md
    carbonstack/docs/README.md
    carbonstack/roadmap/ROADMAP.md

Commit:

    37a6d2a docs: record local-cypher validation contract

The doc records the future `local-cypher` profile contract. It is intentionally docs-only and does not add `local-cypher` to the runner yet.

### 4.3 Future profile name and command shape recorded

Future profile name:

    local-cypher

Intended future command shape:

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go run . --profile local-cypher

Optional explicit-root form:

    go run . --profile local-cypher --root /path/to/carbonstack_umbrella

Important naming boundary:

    Do not call this profile local-backbone.

`local-backbone` remains reserved for a later whole-stack validation surface where CarbonStackComms, CarbonStackCypher, OpenMLS runtime wiring, and backbone-level lifecycle semantics are meaningfully integrated.

### 4.4 Future positive lifecycle contract recorded

A future `local-cypher` profile should validate the Cypher-only local API lifecycle under an explicit local operator environment.

Required positive lifecycle:

    build a temporary Cypher binary
    create a temporary isolated data directory
    create a temporary isolated SQLite DB
    set CYPHER_ADDR to an available 127.0.0.1 port
    set CYPHER_DB to the temporary DB path
    set CYPHER_MIGRATIONS to the source-tree migrations path
    set CYPHER_DEV_INVITE to a known temporary dev invite
    start Cypher
    wait for GET /v0/health
    claim seeded invite
    create second dev invite
    claim second invite
    register two devices
    list account devices
    submit opaque OpenMLS application-message envelope
    retrieve recipient inbox
    verify payload_sha256
    verify payload_size_bytes
    ack envelope
    verify recipient inbox is empty after ack
    stop Cypher
    restart Cypher against the same temporary DB
    verify /v0/health after restart
    verify persisted device state remains visible
    verify acked recipient inbox remains empty after restart
    stop Cypher
    remove temporary binary and temporary DB directory

Required accepted envelope pair:

    content_type=carbonstack.mls.application-message.v0
    protocol_version=carbonstack-openmls-sidecar-v0

### 4.5 Future scope exclusions recorded

`local-cypher` should not validate:

    runtime CarbonStackComms send/inbox UX
    OpenMLS user-facing message flow through Comms CLI
    CarbonStack Relay Space mechanics
    local-backbone behavior
    production E2EE
    hostile-server safety
    metadata privacy
    public ingress
    LAN exposure
    cloudflared
    systemd
    real homelab deployment
    Android app behavior
    CarbonStackOS
    external audit
    certification

The key boundary:

    local-cypher is a Cypher API lifecycle validation profile.
    local-cypher is not a CarbonStack secure messaging claim.

### 4.6 Temporary DB and process lifecycle contract recorded

`local-cypher` should use temporary isolated DBs by default and should not mutate:

    $HOME/.local/share/carbonstack/cypher/cypher.db

Reason:

    validation should be repeatable
    validation should not depend on private local operator state
    validation should not destroy a user's manual operator DB
    validation should avoid stale state coupling
    validation should clean up after itself

`local-cypher` should own the process lifecycle:

    build the temporary binary
    start the binary itself
    wait for health
    run lifecycle calls
    stop the process
    restart the process against the same temp DB
    stop the second process
    clean temporary files

It should not require a pre-existing manually started Cypher server.

### 4.7 Port, artifact, and cleanup contracts recorded

`local-cypher` should bind only to loopback.

Preferred future behavior:

    choose a temporary localhost port dynamically if practical

Acceptable first implementation:

    use a fixed high localhost port with clear error on conflict

Non-negotiable:

    do not bind 0.0.0.0
    do not bind LAN interfaces
    do not expose public ingress

`local-cypher` should leave no source-tree artifacts and should not leave behind:

    temporary binary
    temporary SQLite DB
    temp proof directory
    generated sidecar state
    Rust target directory
    provider-storage.json
    signer.json
    .go-cache
    .go-tmp

### 4.8 Pass/fail and negative-path contract recorded

`local-cypher` should fail if required paths are missing, Cypher build fails, health never succeeds, lifecycle calls fail, payload hash/size mismatches occur, ack behavior is wrong, restart fails, persisted state is missing after restart, or cleanup fails badly enough to leave an active server or suspicious artifacts.

Negative-path coverage is not required for the first implementation, but should include the v0.3.27 blunder later:

    content_type=carbonstack.message.text.stub.v0
    protocol_version=carbonstack-openmls-sidecar-v0

Expected result:

    HTTP 400
    error code unsupported_protocol_version

Recommendation preserved:

    implement positive local-cypher profile first
    add negative-path coverage in a follow-up rung unless implementation is very small and obvious

### 4.9 Final validation and cleanup completed

Final WSL Debian repo snapshot:

    carbonstack        37a6d2a docs: record local-cypher validation contract
    carbonstack-comms  012c8bf scripts: support source snapshot artifact guard
    carbonstack-cypher 9ab994c docs: point to local operator runbook
    carbonstack-os     1bbbe52 docs: clarify CarbonStackOS target direction

Final validation:

    go run . --profile doctor
    go run . --profile core

Final validation passed under WSL Debian. `core` passed OpenMLS real-Cypher lifecycle, Comms package tests, and Cypher package tests, including `internal/db` and `internal/httpapi` package tests. After validation, expected OpenMLS generated roots were removed from `carbonstack-comms`, and final clean status showed no dirty files across the four repos.

---

## 5. Critical Blunders / Lessons Preserved

### 5.1 v0.3.29 correctly did recon before writing the contract

The rung started with a WSL Debian command scout before docs were written. This prevented invented runner mechanics.

Correct lesson:

    Inspect live code and tests first.
    Write contract docs second.
    Implement only after the contract is stable.

### 5.2 v0.3.29 correctly stayed docs-only

The contract could have been turned directly into a runner patch, but that would have skipped the review point.

Correct lesson:

    Contract first.
    Implementation second.
    Patch only after no obvious contract changes are needed.

### 5.3 `local-cypher` was kept narrow

The contract defines Cypher-only validation and preserves the `local-backbone` name for later whole-stack validation.

Correct lesson:

    Names are claims.
    Do not let a Cypher-only lifecycle profile imply full CarbonStack backbone maturity.

### 5.4 Temporary DB policy remains bounded

Temporary DBs are correct for dev/test validation and pre-v0.4.x release-adjacent validation. They are still not sufficient for mature public Comms or real-user testing.

Correct lesson:

    Temporary DB validation is repeatable and safe for development.
    Mature user-facing validation later needs persistent operator/user-state evidence.

### 5.5 Artifact cleanup cadence held correctly

The final validation produced expected sidecar generated roots, then cleanup was run. Final repo status snapshot showed clean status after cleanup. Continue this pattern after every `core` run.

---

## 6. Current Blockers / Not Validated

Current blockers and nonvalidated areas:

    v0.3.29 is contract/recon only and intentionally does not add a runner profile.
    No local-cypher runner profile exists yet.
    No local-backbone runner profile exists yet.
    No helper command exists yet.
    Local-only operator convention is documented and proof-tested, but not enforced by code.
    Cypher default bind remains :8080; local-only docs recommend explicit 127.0.0.1:8080.
    Runtime Comms send/inbox remains stub-era and not wired as user-facing OpenMLS messenger UX.
    Production-safe provider storage and secure vault design are unsolved.
    Hostile-server rollback/replay/metadata-abuse validation remains future work.
    Android/CarbonStackOS implementation remains later work.
    CarbonStack Relay Space mechanics remain a defined future concept, not an implementation.
    Public ingress, systemd, cloudflared, real homelab deployment, and production service paths remain out of scope.
    Temporary DB validation remains dev/test-only and does not replace later persistent user/operator-state validation for public Comms maturity.

Hard nonclaims remain:

    Do not claim production readiness.
    Do not claim production E2EE.
    Do not claim hostile-server safety.
    Do not claim metadata privacy.
    Do not claim Debian production deployability.
    Do not claim systemd readiness.
    Do not claim cloudflared readiness.
    Do not claim Android readiness.
    Do not claim secure vault/storage.
    Do not claim external audit or certification.
    Do not claim CarbonStack Relay Space mechanics are implemented.
    Do not claim local operator profile is complete.
    Do not claim v0.3.29 is a public release.
    Do not claim `local-cypher` exists yet.
    Do not claim `local-backbone` exists or is justified yet.

Updated allowed claim:

    CarbonStack v0.3.29 records the future local-cypher validation contract after a WSL Debian command scout.
    It defines local-cypher as a future Cypher-only, runner-owned, temporary-DB lifecycle validation profile.
    It reserves local-backbone for later whole-stack validation.
    It does not implement the runner profile yet.
    WSL Debian doctor/core validation passed after the v0.3.29 docs commit.
    Final clean status showed no dirty files across the four repos after expected artifact cleanup.

---

## 7. Next Safest Actions

Recommended next sequence after this v0.3.29 breakpoint:

    v0.3.30 preferred:
      implement local-cypher runner profile, if the contract still looks correct

Recommended concrete next rung:

    1. Patch carbonstack/tools/carbonstack-validate to accept --profile local-cypher.
    2. Add runner-owned local Cypher lifecycle implementation.
    3. Build a temporary Cypher binary rather than using go run.
    4. Use loopback only.
    5. Use a temporary isolated DB.
    6. Use source-tree migrations path.
    7. Run the positive invite/device/envelope/ack lifecycle.
    8. Restart Cypher against the same DB.
    9. Verify persisted device state and acked inbox behavior.
    10. Clean temporary files/processes on success and failure where possible.
    11. Update runner README and help text.
    12. Validate doctor/core and local-cypher.
    13. Document result in a new carbonstack docs file.

Alternative next sequence if implementation feels risky:

    v0.3.30:
      local-cypher implementation scout/recon
    v0.3.31:
      implementation patch

Avoid next:

    local-backbone runner profile
    helper implementation
    public ingress
    cloudflared
    systemd
    real homelab deployment
    runtime Comms send/inbox UX
    Android app
    CarbonStackOS
    CarbonStack Relay Space implementation
    production release

---

## 8. Critical Paths / Functions

### WSL Debian active working umbrella

    [WSL-DEBIAN] umbrella: ~/repos/carbonstack_umbrella
    [WSL-DEBIAN] carbonstack: ~/repos/carbonstack_umbrella/carbonstack
    [WSL-DEBIAN] carbonstack-comms: ~/repos/carbonstack_umbrella/carbonstack-comms
    [WSL-DEBIAN] carbonstack-cypher: ~/repos/carbonstack_umbrella/carbonstack-cypher
    [WSL-DEBIAN] carbonstack-os: ~/repos/carbonstack_umbrella/carbonstack-os
    [GO-RUNNER] runner dir: ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate

### WSL Debian known-good validation commands

From runner dir:

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go run . --profile doctor
    go run . --profile core

From carbonstack repo root:

    cd ~/repos/carbonstack_umbrella/carbonstack
    go run ./tools/carbonstack-validate --profile doctor
    go run ./tools/carbonstack-validate --profile core

### Artifact cleanup after validation

    cd ~/repos/carbonstack_umbrella/carbonstack-comms
    rm -rf internal/protocol/mls/openmls-sidecar/.carbonstack-openmls-sidecar-state
    rm -rf internal/protocol/mls/openmls-sidecar/target

### v0.3.29 docs

    [DOC] v0.3.29 contract:
    carbonstack/docs/148-local-cypher-validation-contract-v0.md

    [DOC] v0.3.28 decision recon:
    carbonstack/docs/147-local-operator-helper-runner-decision-recon-v0.md

    [DOC] v0.3.27 proof:
    carbonstack/docs/146-local-cypher-explicit-env-api-lifecycle-proof-v0.md

    [DOC] docs index:
    carbonstack/docs/README.md

    [DOC] roadmap:
    carbonstack/roadmap/ROADMAP.md

### Go runner surfaces

    [RUNNER] directory:
    carbonstack/tools/carbonstack-validate

    [RUNNER] current implementation files:
    carbonstack/tools/carbonstack-validate/main.go
    carbonstack/tools/carbonstack-validate/checksums.go
    carbonstack/tools/carbonstack-validate/release_snapshot.go
    carbonstack/tools/carbonstack-validate/README.md
    carbonstack/tools/carbonstack-validate/go.mod

    [RUNNER] current profiles:
    doctor
    core
    full
    release-snapshot
    write-checksums
    verify-checksums

    [RUNNER] future candidate profile:
    local-cypher

### Future local-cypher lifecycle contract

    start:
      build temporary Cypher binary
      create temporary DB directory
      set CYPHER_ADDR to loopback port
      set CYPHER_DB to temporary SQLite DB path
      set CYPHER_MIGRATIONS to source-tree migrations
      set CYPHER_DEV_INVITE to known temporary dev invite
      start Cypher
      wait for GET /v0/health

    account/device:
      POST /v0/invites/claim
      POST /v0/dev/invites
      POST /v0/invites/claim
      POST /v0/devices/register
      GET /v0/accounts/<account_id>/devices

    envelope:
      POST /v0/envelopes
      GET /v0/devices/<device_id>/envelopes
      POST /v0/envelopes/<envelope_id>/ack
      GET /v0/devices/<device_id>/envelopes

    restart:
      stop Cypher
      restart against same temp DB
      GET /v0/health
      GET /v0/accounts/<account_id>/devices
      GET /v0/devices/<device_id>/envelopes

    cleanup:
      stop process
      remove temp binary
      remove temp DB / proof dir
      leave no source-tree artifacts

### Accepted envelope pair

    content_type=carbonstack.mls.application-message.v0
    protocol_version=carbonstack-openmls-sidecar-v0

### Future negative path candidate

    content_type=carbonstack.message.text.stub.v0
    protocol_version=carbonstack-openmls-sidecar-v0

Expected result:

    HTTP 400
    error code unsupported_protocol_version

### Critical functions / behavior preserved from v0.3.24

    Store.Migrate:
      ensures schema_migrations exists;
      reads migration .sql files in sorted order;
      computes sha256 checksum for each migration;
      skips matching already-applied migrations;
      fails on checksum mismatch for already-applied migration;
      applies new migrations in a transaction;
      records migration after successful application.

    TestMigrateRecordsAppliedMigrations:
      proves applied migrations are recorded with non-empty checksum/applied_at.

    TestMigrateCanRunTwiceAgainstSameDB:
      proves repeated migration against the same DB now succeeds.

    TestMigrateDetectsAppliedMigrationChecksumMismatch:
      proves already-applied migration drift hard-fails.

---

## 9. Lean Breakpoint Summary

    CarbonStack v0.3.29 is a stable mainline docs/contract breakpoint after v0.3.28. It records the future local-cypher validation contract in carbonstack/docs/148 after a WSL Debian read-only command scout, defines local-cypher as a Cypher-only runner-owned lifecycle validation profile using temporary isolated DBs, explicit loopback bind, positive invite/device/envelope/ack lifecycle, restart against the same DB, and cleanup requirements; reserves local-backbone for later whole-stack validation; validates WSL Debian doctor/core after the docs commit; cleans generated OpenMLS artifacts; and leaves local-cypher implementation, helper tooling, local-backbone runner profile, runtime Comms OpenMLS UX, CarbonStack Relay Space mechanics, public ingress, systemd, cloudflared, homelab, Android, CarbonStackOS, hostile-server proof, secure vault, and production/security claims deferred.

---

## 10. Preserved v0.3.28 Operational Process Log

The following section preserves the previous v0.3.28 LogDoc body as operational continuity. It is intentionally retained rather than compressed away. Where v0.3.28 text conflicts with the v0.3.29 current-state overlay above, v0.3.29 wins for current state; v0.3.28 remains the provenance/process ledger for the local operator helper/runner decision recon and the preserved v0.3.27 / v0.3.26 / v0.3.25 / v0.3.24 / v0.3.23 / v0.3.22 / v0.3.21 / v0.3.20 history.



# CarbonStack LogDoc v0.3.28

**Last updated:** 2026-06-02 local session  
**Owner/operator:** bitcrusher32  
**Assistant/model:** GPT-5.5 Thinking  
**Current phase:** Phase 2E-prep / v0.3.x **local operator helper/runner decision recon checkpoint**. `carbonstack` is now at `af33139 docs: record local operator helper runner decision recon`, after `v0.3.27` landed the local Cypher explicit-env API lifecycle proof at `e516fc7`. `carbonstack-cypher` remains at `9ab994c docs: point to local operator runbook`; `carbonstack-comms` remains at `012c8bf scripts: support source snapshot artifact guard`; `carbonstack-os` remains at `1bbbe52 docs: clarify CarbonStackOS target direction`. The live WSL Debian working umbrella at `~/repos/carbonstack_umbrella` passed `doctor` and `core` after the v0.3.28 docs commit, expected OpenMLS sidecar artifacts were cleaned afterward, and final clean status showed no dirty files across the four repos.

**Version schema:** `v[scope].[timeline]`. This file is `v0.3.28`, the local operator helper/runner decision recon handoff after the `v0.3.27` local Cypher explicit-env API lifecycle proof checkpoint. Per LogDoc V2 practice, this Markdown file intentionally preserves verbose operational continuity; the JSON breakpoint remains the lean handoff state. The dense v0.3.27, v0.3.26, v0.3.25, v0.3.24, v0.3.23, v0.3.22, v0.3.21, and v0.3.20 operational process logs are preserved later in this file rather than compressed away.

**LogDoc V2 usage note:** LogDoc V2 uses a split between a Markdown LogDoc for working-session continuity, a JSON state file for end-of-session handoff, Git for history preservation, and README/docs for public claims. This v0.3.28 Markdown file therefore intentionally stays “bloated” where operational process, blunders, validation ladders, and historical context matter. The v0.3.28 JSON should remain lean and should agree with this file on the latest current state.

---

## 1. Project Goal

**Active goal:** Preserve CarbonStack state after recording the local operator helper/runner decision recon. The immediate next line of work should be a `local-cypher` validation contract doc or one more lifecycle proof review, not helper implementation, not `local-backbone`, not runtime Comms UX, not public ingress, not systemd/cloudflared, not real homelab deployment, and not Android/CarbonStackOS work.

CarbonStack’s broader project intent remains:

- **CarbonStack** is the secure-communications appliance-stack project and shared doctrine/release authority.
- **CarbonStackOS** is the future constrained Android-derived appliance OS.
- **CarbonStackComms** is the text-first encrypted messaging client and current OpenMLS sidecar/relay integration surface.
- **CarbonStackCypher** is the self-hostable relay/storage stack.
- The `carbonstack` repo remains the doctrine/spec/docs/release authority; implementation repos must not contradict it without revising the spec.

**North star:**

> Its goal is not to be a better smartphone. Its goal is to stop being a smartphone.

**Core doctrine:**

> Every feature is guilty until it proves it does not add unacceptable parser, network, sensor, identity, or filesystem authority.

**Current reality after v0.3.28:** `v0.3.20` remains the current public runner-backed testing release. `v0.3.21` defined the local-only deployability model. `v0.3.22` narrowed the local deployability surface around Cypher operator config/API/database assumptions and adopted **CarbonStack Relay Space** terminology. `v0.3.23` proved the SQLite repeated-migration persistence hazard. `v0.3.24` implemented `schema_migrations` tracking in CarbonStackCypher. `v0.3.25` recorded the first local-only Cypher operator runbook skeleton. `v0.3.26` recorded the local operator config/data convention. `v0.3.27` proved a real local-only Cypher explicit-env API lifecycle. `v0.3.28` now records the helper/runner decision recon: do not add helper tooling yet, do not add a runner profile yet, reserve `local-backbone` for later whole-stack validation, use `local-cypher` as the future Cypher-only profile name if/when justified, and keep temporary isolated DB validation as a dev/test-only mechanism rather than a mature public Comms or real-user-testing substitute.

The current public release page remains:

    https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.3.20

The current public testing interpretation remains:

    v0.3.20 is the preferred public runner-backed testing release.
    v0.3.4 remains available as the older Windows 11 / PowerShell legacy testbed.
    Debian / WSL Debian is now the mainline public dev/test validation direction.
    Windows 11 short-path validation was the final explicit Windows dev/test validation for this phase.

This is still **not** a deployability release. It does **not** prove production readiness, production E2EE, hostile-server safety, metadata privacy, Android readiness, secure vault/storage, Debian production deployability, systemd readiness, cloudflared readiness, public ingress safety, real homelab validation, external audit, or certification.

---

## 2. Current Repo Heads

    carbonstack        af33139 (HEAD -> main, origin/main, origin/HEAD) docs: record local operator helper runner decision recon
    carbonstack-comms  012c8bf (HEAD -> main, origin/main, origin/HEAD) scripts: support source snapshot artifact guard
    carbonstack-cypher 9ab994c (HEAD -> main, origin/main, origin/HEAD) docs: point to local operator runbook
    carbonstack-os     1bbbe52 (HEAD -> main, origin/main, origin/HEAD) docs: clarify CarbonStackOS target direction

Current `carbonstack` public runner-backed testing release tag:

    v0.3.20 -> 3eeb1a1e1a / 3eeb1a1 docs: clean public release surface

Previous mainline deployability checkpoints:

    v0.3.27 mainline -> e516fc7 docs: record local Cypher API lifecycle proof
    v0.3.26 mainline -> 6e14fe0 docs: record local operator config data convention
    v0.3.25 mainline -> a5c6351 docs: update readme, roadmap, and historical doc
    v0.3.25 cypher  -> 9ab994c docs: point to local operator runbook
    v0.3.24 mainline -> 76f0258 docs: record Cypher schema migrations implementation
    v0.3.24 cypher  -> 42f838f db: track applied migrations
    v0.3.23 mainline -> 1fef7fe docs: record Cypher migration persistence recon
    v0.3.22 mainline -> 328ffc4 docs: record Cypher local operator surface recon
    v0.3.21 mainline -> 6850c58 docs: record local backbone deployability recon

Legacy `carbonstack` Windows/PowerShell testbed release tag:

    v0.3.4 -> 7c5b7e6 docs: record release snapshot self-test fix

Historical `carbonstack` release tag:

    v0.3.0 -> 59245fc docs: freeze v0.3.0 release packaging plan

Important continuity note: `v0.3.28` does not supersede `v0.3.20` as a public release. It advances mainline local-only deployability reasoning after the v0.3.27 explicit-env lifecycle proof. It is safe to pause here because the decision recon doc was pushed, WSL Debian `doctor`/`core` passed after the commit, generated artifacts were cleaned, and final status was clean across all four repos.

---

## 3. Current Validated State

Validated at the v0.3.28 checkpoint:

    [WSL-DEBIAN] working umbrella exists at ~/repos/carbonstack_umbrella
    [WSL-DEBIAN] carbonstack is at af33139
    [WSL-DEBIAN] carbonstack-comms is at 012c8bf
    [WSL-DEBIAN] carbonstack-cypher is at 9ab994c
    [WSL-DEBIAN] carbonstack-os is at 1bbbe52
    [CARBONSTACK] docs/147 local operator helper/runner decision recon committed and pushed
    [GO-RUNNER] doctor passed after the v0.3.28 docs commit
    [GO-RUNNER] core passed after the v0.3.28 docs commit
    [ARTIFACT CLEANUP] expected OpenMLS generated roots were removed after validation
    [FINAL STATUS] all four repos were clean after cleanup and commit

Observed WSL Debian toolchain baseline:

    git version 2.47.3
    go version go1.24.4 linux/amd64
    rustc 1.96.0 (ac68faa20 2026-05-25)
    cargo 1.96.0 (30a34c682 2026-05-25)
    sqlite3 3.46.1

Observed `core` validation result during the v0.3.28 flow:

    pre-test artifact scan: PASS / no generated/private/build artifact hits
    OpenMLS real-Cypher lifecycle: PASS
    carbonstack-comms package tests: PASS
    carbonstack-cypher package tests: PASS
    post-test artifact scan: expected OpenMLS sidecar generated roots only
    VALIDATION PASSED

Expected post-test generated roots:

    carbonstack-comms/internal/protocol/mls/openmls-sidecar/.carbonstack-openmls-sidecar-state
    carbonstack-comms/internal/protocol/mls/openmls-sidecar/target

These roots are expected only after tests and must not be committed or treated as source artifacts. They were cleaned before the final clean status snapshot.

---

## 4. v0.3.28 Work Completed

### 4.1 Local operator helper/runner decision recon landed

New/updated docs at v0.3.28:

    carbonstack/docs/147-local-operator-helper-runner-decision-recon-v0.md
    carbonstack/docs/README.md
    carbonstack/roadmap/ROADMAP.md

Commit:

    af33139 docs: record local operator helper runner decision recon

The doc records a pause-and-decide checkpoint after the v0.3.27 local Cypher explicit-env API lifecycle proof. It is intentionally docs-only and does not add helper tooling, runner profiles, config parsers, Cypher behavior changes, Comms runtime wiring, public ingress, systemd, cloudflared, or homelab guidance.

### 4.2 Helper tooling intentionally deferred

v0.3.28 records that helper tooling is plausible later, but not now.

Possible future helper shape remains:

    carbonstack local start
    carbonstack local health
    carbonstack local stop
    carbonstack local reset

Decision:

    no helper in v0.3.28
    no helper before the lifecycle contract is clearer
    prefer explicit manual commands and docs for now

Reasoning:

    v0.3.27 proved a lifecycle manually, but did not define mature operator UX.
    helper commands would freeze start/stop/reset semantics too early.
    reset behavior is intentionally blunt and should remain user-visible.
    helper tooling risks hiding important experimental state boundaries.
    config parser/helper decisions should follow a stable manual convention.

### 4.3 Runner profile intentionally deferred

v0.3.28 records that runner automation is plausible later, but premature.

Decision:

    no local-cypher runner profile in v0.3.28
    no local-backbone runner profile in v0.3.28
    revisit after at least one more lifecycle proof/review or contract-definition rung

Open questions preserved before adding a profile:

    Should the runner build and start Cypher itself?
    Should the runner target an already-running Cypher process?
    Should the runner own cleanup?
    Should it always use a temporary DB?
    Should it ever touch the documented manual operator DB path?
    Should negative API paths be included?
    Should the profile be Cypher-only or include Comms?
    Should it be release-blocking or informational?

### 4.4 Naming decision recorded: `local-cypher` vs `local-backbone`

v0.3.28 records the naming split:

    local-cypher:
      future candidate name for Cypher-only lifecycle validation

    local-backbone:
      reserved for a later top-level runner/wrapper that validates the actual CarbonStack backbone when Comms/Cypher/OpenMLS are wired meaningfully enough to deserve that name

Current rule:

    Do not call a Cypher-only lifecycle proof local-backbone.

This is important because `v0.3.27` proved a Cypher API lifecycle only. It did not validate runtime Comms UX, end-to-end user messaging, a full CarbonStack backbone, or CarbonStack Relay Space mechanics.

### 4.5 Temporary DB policy recorded

v0.3.28 records temporary isolated DBs as the correct default for dev/test validation and possible future `local-cypher` validation.

This applies to:

    local lifecycle proof scripts
    future local-cypher runner profile, if added
    release-adjacent dev/test validation
    CI-like validation paths

Reasoning:

    avoids mutating a user's manual operator DB
    makes runs repeatable
    makes cleanup straightforward
    avoids stale state coupling
    avoids accidental reliance on local private state
    fits source/test separation established in v0.3.26

Important maturity boundary:

    Temporary DB validation is acceptable for dev/test and pre-v0.4.x release-adjacent validation.
    It should not be treated as the final model for public Comms use or real-user testing.
    Later maturity must address real persistent operator/user state, backup/restore expectations, upgrade semantics, and failure recovery.

### 4.6 Negative-path validation deferred but preserved

v0.3.28 preserves the v0.3.27 accidental negative-path finding:

    content_type=carbonstack.message.text.stub.v0
    protocol_version=carbonstack-openmls-sidecar-v0

Cypher correctly rejected the pair with:

    unsupported_protocol_version

Decision:

    do not add negative-path runner coverage in v0.3.28
    record it as future validation maturity work

Candidate future negative paths:

    invalid content/protocol pairing
    invalid base64 envelope
    ack unknown envelope
    ack with wrong recipient
    malformed device lookup
    duplicate invite code
    already-claimed invite

### 4.7 `127.0.0.1` code default change deferred

Current code default still allows:

    CYPHER_ADDR=:8080

Current local operator docs recommend:

    CYPHER_ADDR=127.0.0.1:8080

Decision:

    do not change the code default in v0.3.28

Reasoning:

    default bind change is real behavior, not docs cleanup
    it should be its own focused code/config rung if chosen
    current docs already require explicit local operator env
    v0.3.x should avoid mixing docs/recon with unrelated behavior changes

### 4.8 Final validation and cleanup completed

Final WSL Debian repo snapshot:

    carbonstack        af33139 docs: record local operator helper runner decision recon
    carbonstack-comms  012c8bf scripts: support source snapshot artifact guard
    carbonstack-cypher 9ab994c docs: point to local operator runbook
    carbonstack-os     1bbbe52 docs: clarify CarbonStackOS target direction

Final validation:

    go run . --profile doctor
    go run . --profile core

Final validation passed under WSL Debian. `core` passed OpenMLS real-Cypher lifecycle, Comms package tests, and Cypher package tests, including `internal/db` and `internal/httpapi` package tests. After validation, expected OpenMLS generated roots were removed from `carbonstack-comms`, and final clean status showed no dirty files across the four repos.

---

## 5. Critical Blunders / Lessons Preserved

### 5.1 v0.3.28 correctly stayed docs-only

After a successful lifecycle proof, the tempting next move was to add a helper or runner profile. v0.3.28 avoided that and preserved the design decision first.

Correct lesson:

    Evidence first.
    Decision recon second.
    Tooling third, only once the contract is concrete.

### 5.2 `local-backbone` naming was correctly reserved

Calling a Cypher-only lifecycle validation profile `local-backbone` would overstate maturity. The name should be reserved for a real whole-CarbonStack backbone validation surface once Comms/Cypher/OpenMLS are wired meaningfully enough.

Correct lesson:

    `local-cypher` can describe Cypher-only validation.
    `local-backbone` should describe whole-stack backbone validation later.
    Naming is a claim boundary.

### 5.3 Temporary DB validation is useful but bounded

Temporary DB validation is right for dev/test because it is repeatable and safe. But it must not be used forever as the only evidence for public Comms or real-user testing maturity.

Correct lesson:

    temp DBs are good for dev/test and release-adjacent pre-v0.4.x validation.
    mature public/user testing must later address persistent real operator/user state, upgrade, backup/restore, and failure recovery.

### 5.4 Helper scripts can hide experimental state

A helper script too early would risk hiding the fact that reset is destructive, state is experimental, and lifecycle semantics are still being defined.

Correct lesson:

    Manual commands remain useful because they keep state boundaries visible.
    Helper tooling should clarify mature semantics, not conceal immature ones.

### 5.5 Artifact cleanup cadence held correctly

The final validation produced expected sidecar generated roots, then cleanup was run. Final repo status snapshot showed clean status after cleanup. Continue this pattern after every `core` run.

---

## 6. Current Blockers / Not Validated

Current blockers and nonvalidated areas:

    v0.3.28 decides not to add helper/runner tooling yet.
    Local-only operator convention is documented and proof-tested, but not enforced by code.
    Cypher default bind remains :8080; local-only docs recommend explicit 127.0.0.1:8080.
    No local-cypher runner profile exists yet.
    No local-backbone runner profile exists yet.
    No helper command exists yet.
    No config file parser exists yet.
    Runtime Comms send/inbox remains stub-era and not wired as user-facing OpenMLS messenger UX.
    Production-safe provider storage and secure vault design are unsolved.
    Hostile-server rollback/replay/metadata-abuse validation remains future work.
    Android/CarbonStackOS implementation remains later work.
    CarbonStack Relay Space mechanics remain a defined future concept, not an implementation.
    Public ingress, systemd, cloudflared, real homelab deployment, and production service paths remain out of scope.
    Temporary DB validation remains dev/test-only and does not replace later persistent user/operator-state validation for public Comms maturity.

Hard nonclaims remain:

    Do not claim production readiness.
    Do not claim production E2EE.
    Do not claim hostile-server safety.
    Do not claim metadata privacy.
    Do not claim Debian production deployability.
    Do not claim systemd readiness.
    Do not claim cloudflared readiness.
    Do not claim Android readiness.
    Do not claim secure vault/storage.
    Do not claim external audit or certification.
    Do not claim CarbonStack Relay Space mechanics are implemented.
    Do not claim local operator profile is complete.
    Do not claim v0.3.28 is a public release.
    Do not claim `local-backbone` exists or is justified yet.

Updated allowed claim:

    CarbonStack v0.3.28 records a local operator helper/runner decision recon after the v0.3.27 local Cypher explicit-env API lifecycle proof.
    It defers helper tooling and runner profile implementation.
    It reserves `local-backbone` for later whole-stack validation.
    It treats `local-cypher` as the future candidate name for Cypher-only validation if/when justified.
    It confirms temporary isolated DBs as the correct dev/test validation default while noting that they are not sufficient for later public Comms or real-user-testing maturity.
    WSL Debian doctor/core validation passed after the v0.3.28 docs commit.
    Final clean status showed no dirty files across the four repos after expected artifact cleanup.

---

## 7. Next Safest Actions

Recommended next sequence after this v0.3.28 breakpoint:

    v0.3.29 preferred:
      local-cypher validation contract doc, not implementation yet

Recommended concrete next rung:

    1. Define what a future `local-cypher` runner profile would own.
    2. Define what it would not own.
    3. Decide whether it starts/stops Cypher itself.
    4. Decide temporary DB lifecycle details.
    5. Decide whether negative paths are in-scope for the first profile.
    6. Decide whether it is release-blocking, informational, or dev-only.
    7. Keep `local-backbone` reserved for later whole-stack validation.
    8. Keep helper tooling deferred unless the contract becomes obvious.

Alternative next sequence if more proof is desired before runner-contract docs:

    v0.3.29:
      negative-path local Cypher validation recon

    v0.3.30:
      local-cypher validation contract doc

Avoid next:

    helper implementation
    local-backbone runner profile
    public ingress
    cloudflared
    systemd
    real homelab deployment
    runtime Comms send/inbox UX
    Android app
    CarbonStackOS
    CarbonStack Relay Space implementation
    production release

---

## 8. Critical Paths / Functions

### WSL Debian active working umbrella

    [WSL-DEBIAN] umbrella: ~/repos/carbonstack_umbrella
    [WSL-DEBIAN] carbonstack: ~/repos/carbonstack_umbrella/carbonstack
    [WSL-DEBIAN] carbonstack-comms: ~/repos/carbonstack_umbrella/carbonstack-comms
    [WSL-DEBIAN] carbonstack-cypher: ~/repos/carbonstack_umbrella/carbonstack-cypher
    [WSL-DEBIAN] carbonstack-os: ~/repos/carbonstack_umbrella/carbonstack-os
    [GO-RUNNER] runner dir: ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate

### WSL Debian known-good validation commands

From runner dir:

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go run . --profile doctor
    go run . --profile core

From carbonstack repo root:

    cd ~/repos/carbonstack_umbrella/carbonstack
    go run ./tools/carbonstack-validate --profile doctor
    go run ./tools/carbonstack-validate --profile core

### Artifact cleanup after validation

    cd ~/repos/carbonstack_umbrella/carbonstack-comms
    rm -rf internal/protocol/mls/openmls-sidecar/.carbonstack-openmls-sidecar-state
    rm -rf internal/protocol/mls/openmls-sidecar/target

### v0.3.28 docs

    [DOC] v0.3.28 decision recon:
    carbonstack/docs/147-local-operator-helper-runner-decision-recon-v0.md

    [DOC] v0.3.27 proof:
    carbonstack/docs/146-local-cypher-explicit-env-api-lifecycle-proof-v0.md

    [DOC] v0.3.26 convention:
    carbonstack/docs/145-local-operator-config-data-convention-v0.md

    [DOC] docs index:
    carbonstack/docs/README.md

    [DOC] roadmap:
    carbonstack/roadmap/ROADMAP.md

### Local operator convention paths

    [LOCAL OPERATOR] data directory:
    ~/.local/share/carbonstack/cypher

    [LOCAL OPERATOR] DB path:
    ~/.local/share/carbonstack/cypher/cypher.db

    [LOCAL OPERATOR] config directory:
    ~/.config/carbonstack/cypher

    [LOCAL OPERATOR] source-tree migrations path:
    ~/repos/carbonstack_umbrella/carbonstack-cypher/migrations

### v0.3.27 API lifecycle contract preserved for future `local-cypher`

    health:
      GET /v0/health

    claim seeded invite:
      POST /v0/invites/claim
      invite_code=dev-invite
      display_name=<name>

    create dev invite:
      POST /v0/dev/invites
      invite_code=<code>

    register device:
      POST /v0/devices/register
      account_id=<account_id>
      device_label=<label>
      public_identity_key=<stub or real future key>
      public_prekey_bundle=<stub or real future prekey bundle>

    list account devices:
      GET /v0/accounts/<account_id>/devices

    submit opaque OpenMLS application-message envelope:
      POST /v0/envelopes
      sender_device_id=<device_id>
      recipient_device_id=<device_id>
      content_type=carbonstack.mls.application-message.v0
      protocol_version=carbonstack-openmls-sidecar-v0
      ciphertext_b64=<base64 opaque payload>

    retrieve recipient inbox:
      GET /v0/devices/<device_id>/envelopes

    ack envelope:
      POST /v0/envelopes/<envelope_id>/ack
      recipient_device_id=<recipient_device_id>

### Critical functions / behavior preserved from v0.3.24

    Store.Migrate:
      ensures schema_migrations exists;
      reads migration .sql files in sorted order;
      computes sha256 checksum for each migration;
      skips matching already-applied migrations;
      fails on checksum mismatch for already-applied migration;
      applies new migrations in a transaction;
      records migration after successful application.

    TestMigrateRecordsAppliedMigrations:
      proves applied migrations are recorded with non-empty checksum/applied_at.

    TestMigrateCanRunTwiceAgainstSameDB:
      proves repeated migration against the same DB now succeeds.

    TestMigrateDetectsAppliedMigrationChecksumMismatch:
      proves already-applied migration drift hard-fails.

---

## 9. Lean Breakpoint Summary

    CarbonStack v0.3.28 is a stable mainline docs/decision breakpoint after v0.3.27. It records a local operator helper/runner decision recon in carbonstack/docs/147, keeps helper tooling and runner profile implementation deferred, reserves local-backbone for later whole-stack validation, treats local-cypher as the future candidate name for Cypher-only lifecycle validation, confirms temporary isolated DBs as the correct dev/test validation default while noting they are not sufficient for later public Comms/real-user testing maturity, validates WSL Debian doctor/core after the docs commit, cleans generated OpenMLS artifacts, and leaves local-cypher validation contract, helper tooling, local-backbone runner profile, runtime Comms OpenMLS UX, CarbonStack Relay Space mechanics, public ingress, systemd, cloudflared, homelab, Android, CarbonStackOS, hostile-server proof, secure vault, and production/security claims deferred.

---

## 10. Preserved v0.3.27 Operational Process Log

The following section preserves the previous v0.3.27 LogDoc body as operational continuity. It is intentionally retained rather than compressed away. Where v0.3.27 text conflicts with the v0.3.28 current-state overlay above, v0.3.28 wins for current state; v0.3.27 remains the provenance/process ledger for the local Cypher explicit-env API lifecycle proof and the preserved v0.3.26 / v0.3.25 / v0.3.24 / v0.3.23 / v0.3.22 / v0.3.21 / v0.3.20 history.



# CarbonStack LogDoc v0.3.27

**Last updated:** 2026-06-02 local session  
**Owner/operator:** bitcrusher32  
**Assistant/model:** GPT-5.5 Thinking  
**Current phase:** Phase 2E-prep / v0.3.x **local Cypher explicit-env API lifecycle proof checkpoint**. `carbonstack` is now at `e516fc7 docs: record local Cypher API lifecycle proof`, after `v0.3.26` landed the local operator config/data convention at `6e14fe0`. `carbonstack-cypher` remains at `9ab994c docs: point to local operator runbook`; `carbonstack-comms` remains at `012c8bf scripts: support source snapshot artifact guard`; `carbonstack-os` remains at `1bbbe52 docs: clarify CarbonStackOS target direction`. The live WSL Debian working umbrella at `~/repos/carbonstack_umbrella` passed `doctor` and `core` after the v0.3.27 docs commit, expected OpenMLS sidecar artifacts were cleaned afterward, and final clean status showed no dirty files across the four repos.

**Version schema:** `v[scope].[timeline]`. This file is `v0.3.27`, the local Cypher explicit-env API lifecycle proof handoff after the `v0.3.26` local operator config/data convention checkpoint. Per LogDoc V2 practice, this Markdown file intentionally preserves verbose operational continuity; the JSON breakpoint remains the lean handoff state. The dense v0.3.26, v0.3.25, v0.3.24, v0.3.23, v0.3.22, v0.3.21, and v0.3.20 operational process logs are preserved later in this file rather than compressed away.

**LogDoc V2 usage note:** LogDoc V2 uses a split between a Markdown LogDoc for working-session continuity, a JSON state file for end-of-session handoff, Git for history preservation, and README/docs for public claims. This v0.3.27 Markdown file therefore intentionally stays “bloated” where operational process, blunders, validation ladders, and historical context matter. The v0.3.27 JSON should remain lean and should agree with this file on the latest current state.

---

## 1. Project Goal

**Active goal:** Preserve CarbonStack state after proving the local-only Cypher API lifecycle under the explicit local operator environment recorded in v0.3.26. The immediate next line of work is a local operator lifecycle proof review and helper/runner decision point, not runtime Comms UX, not public ingress, not systemd/cloudflared, not real homelab deployment, and not Android/CarbonStackOS work.

CarbonStack’s broader project intent remains:

- **CarbonStack** is the secure-communications appliance-stack project and shared doctrine/release authority.
- **CarbonStackOS** is the future constrained Android-derived appliance OS.
- **CarbonStackComms** is the text-first encrypted messaging client and current OpenMLS sidecar/relay integration surface.
- **CarbonStackCypher** is the self-hostable relay/storage stack.
- The `carbonstack` repo remains the doctrine/spec/docs/release authority; implementation repos must not contradict it without revising the spec.

**North star:**

> Its goal is not to be a better smartphone. Its goal is to stop being a smartphone.

**Core doctrine:**

> Every feature is guilty until it proves it does not add unacceptable parser, network, sensor, identity, or filesystem authority.

**Current reality after v0.3.27:** `v0.3.20` remains the current public runner-backed testing release. `v0.3.21` defined the local-only deployability model. `v0.3.22` narrowed the local deployability surface around Cypher operator config/API/database assumptions and adopted **CarbonStack Relay Space** terminology. `v0.3.23` proved the SQLite repeated-migration persistence hazard. `v0.3.24` implemented `schema_migrations` tracking in CarbonStackCypher. `v0.3.25` recorded the first local-only Cypher operator runbook skeleton. `v0.3.26` recorded the local operator config/data convention. `v0.3.27` now records a real local-only Cypher API lifecycle proof under explicit env vars: health, invite claim, dev invite creation, second invite claim, device registration, account device listing, opaque OpenMLS application-message envelope submit/retrieve/ack, restart against the same DB, and persistence checks after restart.

The current public release page remains:

    https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.3.20

The current public testing interpretation remains:

    v0.3.20 is the preferred public runner-backed testing release.
    v0.3.4 remains available as the older Windows 11 / PowerShell legacy testbed.
    Debian / WSL Debian is now the mainline public dev/test validation direction.
    Windows 11 short-path validation was the final explicit Windows dev/test validation for this phase.

This is still **not** a deployability release. It does **not** prove production readiness, production E2EE, hostile-server safety, metadata privacy, Android readiness, secure vault/storage, Debian production deployability, systemd readiness, cloudflared readiness, public ingress safety, real homelab validation, external audit, or certification.

---

## 2. Current Repo Heads

    carbonstack        e516fc7 (HEAD -> main, origin/main, origin/HEAD) docs: record local Cypher API lifecycle proof
    carbonstack-comms  012c8bf (HEAD -> main, origin/main, origin/HEAD) scripts: support source snapshot artifact guard
    carbonstack-cypher 9ab994c (HEAD -> main, origin/main, origin/HEAD) docs: point to local operator runbook
    carbonstack-os     1bbbe52 (HEAD -> main, origin/main, origin/HEAD) docs: clarify CarbonStackOS target direction

Current `carbonstack` public runner-backed testing release tag:

    v0.3.20 -> 3eeb1a1e1a / 3eeb1a1 docs: clean public release surface

Previous mainline deployability checkpoints:

    v0.3.26 mainline -> 6e14fe0 docs: record local operator config data convention
    v0.3.25 mainline -> a5c6351 docs: update readme, roadmap, and historical doc
    v0.3.25 cypher  -> 9ab994c docs: point to local operator runbook
    v0.3.24 mainline -> 76f0258 docs: record Cypher schema migrations implementation
    v0.3.24 cypher  -> 42f838f db: track applied migrations
    v0.3.23 mainline -> 1fef7fe docs: record Cypher migration persistence recon
    v0.3.22 mainline -> 328ffc4 docs: record Cypher local operator surface recon
    v0.3.21 mainline -> 6850c58 docs: record local backbone deployability recon

Legacy `carbonstack` Windows/PowerShell testbed release tag:

    v0.3.4 -> 7c5b7e6 docs: record release snapshot self-test fix

Historical `carbonstack` release tag:

    v0.3.0 -> 59245fc docs: freeze v0.3.0 release packaging plan

Important continuity note: `v0.3.27` does not supersede `v0.3.20` as a public release. It advances mainline local-only deployability evidence after the v0.3.26 convention doc. It is safe to pause here because the lifecycle proof doc was pushed, WSL Debian `doctor`/`core` passed after the commit, generated artifacts were cleaned, and final status was clean across all four repos.

---

## 3. Current Validated State

Validated at the v0.3.27 checkpoint:

    [WSL-DEBIAN] working umbrella exists at ~/repos/carbonstack_umbrella
    [WSL-DEBIAN] carbonstack is at e516fc7
    [WSL-DEBIAN] carbonstack-comms is at 012c8bf
    [WSL-DEBIAN] carbonstack-cypher is at 9ab994c
    [WSL-DEBIAN] carbonstack-os is at 1bbbe52
    [CARBONSTACK] docs/146 local Cypher explicit-env API lifecycle proof committed and pushed
    [GO-RUNNER] doctor passed after the v0.3.27 docs commit
    [GO-RUNNER] core passed after the v0.3.27 docs commit
    [ARTIFACT CLEANUP] expected OpenMLS generated roots were removed after validation
    [FINAL STATUS] all four repos were clean after cleanup and commit

Observed WSL Debian toolchain baseline:

    git version 2.47.3
    go version go1.24.4 linux/amd64
    rustc 1.96.0 (ac68faa20 2026-05-25)
    cargo 1.96.0 (30a34c682 2026-05-25)
    sqlite3 3.46.1

Observed `core` validation result during the v0.3.27 flow:

    pre-test artifact scan: PASS / no generated/private/build artifact hits
    OpenMLS real-Cypher lifecycle: PASS
    carbonstack-comms package tests: PASS
    carbonstack-cypher package tests: PASS
    post-test artifact scan: expected OpenMLS sidecar generated roots only
    VALIDATION PASSED

Expected post-test generated roots:

    carbonstack-comms/internal/protocol/mls/openmls-sidecar/.carbonstack-openmls-sidecar-state
    carbonstack-comms/internal/protocol/mls/openmls-sidecar/target

These roots are expected only after tests and must not be committed or treated as source artifacts. They were cleaned before the final clean status snapshot.

---

## 4. v0.3.27 Work Completed

### 4.1 Local Cypher explicit-env API lifecycle proof landed

New/updated docs at v0.3.27:

    carbonstack/docs/146-local-cypher-explicit-env-api-lifecycle-proof-v0.md
    carbonstack/docs/README.md
    carbonstack/roadmap/ROADMAP.md

Commit:

    e516fc7 docs: record local Cypher API lifecycle proof

The doc records a real local-only API lifecycle proof under explicit env vars and a temporary local SQLite DB. This moves the v0.3.x line beyond a health-check-only proof.

### 4.2 Proof environment

The proof used an isolated local proof setup:

    CYPHER_ADDR=127.0.0.1:18081
    CYPHER_DB=$HOME/.local/share/carbonstack/cypher-v0327-lifecycle-proof/cypher.db
    CYPHER_MIGRATIONS=$HOME/repos/carbonstack_umbrella/carbonstack-cypher/migrations
    CYPHER_DEV_INVITE=dev-invite

The proof built/launched a temporary Cypher binary from `carbonstack-cypher` rather than using `go run`, following the v0.3.25 lesson that process-lifecycle proofs should avoid `go run` wrapper cleanup ambiguity. The temporary DB and binary were intended to be deleted after the proof. Final clean status showed no dirty project repo files.

### 4.3 Lifecycle steps proven

The v0.3.27 proof completed the local-only Cypher lifecycle:

    first health check
    claim seeded alice invite
    create bob dev invite
    claim bob invite
    register alice device
    register bob device
    list alice devices
    submit opaque OpenMLS application-message envelope
    retrieve bob inbox before ack
    verify payload hash/size metadata
    ack envelope
    retrieve bob inbox after ack
    restart Cypher against the same DB
    restart health check against same DB
    list persisted alice devices after restart
    verify acked bob inbox remains empty after restart

The corrected envelope pair was:

    content_type=carbonstack.mls.application-message.v0
    protocol_version=carbonstack-openmls-sidecar-v0

This keeps the payload opaque and matches the current Cypher-accepted OpenMLS sidecar envelope pattern.

### 4.4 Persistence-after-restart proof added

The proof did not stop at the first process lifecycle. It restarted Cypher against the same temporary SQLite DB and verified:

    health endpoint responded after restart
    already-applied migrations did not block restart
    persisted Alice device state remained visible
    Bob's acked inbox remained empty after restart

This gives the local-only operator line stronger evidence that the v0.3.24 `schema_migrations` work and v0.3.26 operator convention can support a basic restart path under explicit env vars.

### 4.5 Final validation and cleanup completed

Final WSL Debian repo snapshot:

    carbonstack        e516fc7 docs: record local Cypher API lifecycle proof
    carbonstack-comms  012c8bf scripts: support source snapshot artifact guard
    carbonstack-cypher 9ab994c docs: point to local operator runbook
    carbonstack-os     1bbbe52 docs: clarify CarbonStackOS target direction

Final validation:

    go run . --profile doctor
    go run . --profile core

Final validation passed under WSL Debian. `core` passed OpenMLS real-Cypher lifecycle, Comms package tests, and Cypher package tests, including `internal/db` and `internal/httpapi` package tests. After validation, expected OpenMLS generated roots were removed from `carbonstack-comms`, and final clean status showed no dirty files across the four repos.

---

## 5. Critical Blunders / Lessons Preserved

### 5.1 Python variable-shadowing blunder: `status` function overwritten by HTTP status int

The first v0.3.27 script had:

    def status(repo):
        ...

and later reused:

    status, alice, raw = http_json(...)

This overwrote the `status()` function with an integer HTTP status code. When the doc template later tried to call `status(carbonstack)`, Python threw:

    TypeError: 'int' object is not callable

Correct lesson:

    Avoid generic function names like status in long scripts.
    Use explicit names like repo_status and http_status.
    Treat script-generation helpers as production-ish enough to avoid obvious shadowing traps.

### 5.2 Wrong protocol/content-type pairing blunder

The first lifecycle proof attempt used:

    content_type=carbonstack.message.text.stub.v0
    protocol_version=carbonstack-openmls-sidecar-v0

Cypher correctly rejected this with:

    unsupported_protocol_version

This was a proof-script blunder, not a Cypher failure. The corrected proof used:

    content_type=carbonstack.mls.application-message.v0
    protocol_version=carbonstack-openmls-sidecar-v0

Correct lesson:

    Do not casually mix stub-era text content types with OpenMLS sidecar protocol versions.
    Use the API test contract as the authority before scripting lifecycle proofs.
    Preserve rejected proof attempts when they reveal contract boundaries.

### 5.3 Broad grep / terminal scroll blunder

During the failed proof investigation, broad output pushed the actual failure out of the visible terminal scrollback. The failure was recovered by reading the generated doc with `cat`.

Correct lesson:

    For proof scripts, write the exact error into the generated doc.
    Use targeted grep after failure.
    Prefer `cat docs/<proof-doc>` or `grep -n "Proof errors"` over broad-spectrum greps when debugging.

### 5.4 v0.3.27 correctly remained docs/result-only

The proof did not require changing `carbonstack-cypher` code. The current Cypher API was capable of completing the lifecycle once the script used the correct content/protocol pair.

Correct lesson:

    Do not modify code to compensate for a bad proof script.
    First confirm whether the script or the system is wrong.
    Only patch implementation repos when the repo behavior is actually deficient.

### 5.5 Artifact cleanup cadence held correctly

The final validation produced expected sidecar generated roots, then cleanup was run. Final repo status snapshot showed clean status after cleanup. Continue this pattern after every `core` run.

---

## 6. Current Blockers / Not Validated

Current blockers and nonvalidated areas:

    v0.3.27 proves a local explicit-env API lifecycle, but does not implement helper tooling.
    Local-only operator convention is documented and proof-tested, but not enforced by code.
    Cypher default bind remains :8080; local-only docs recommend explicit 127.0.0.1:8080.
    No local-backbone runner profile exists yet.
    No helper command exists yet.
    No config file parser exists yet.
    Runtime Comms send/inbox remains stub-era and not wired as user-facing OpenMLS messenger UX.
    Production-safe provider storage and secure vault design are unsolved.
    Hostile-server rollback/replay/metadata-abuse validation remains future work.
    Android/CarbonStackOS implementation remains later work.
    CarbonStack Relay Space mechanics remain a defined future concept, not an implementation.
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
    Do not claim secure vault/storage.
    Do not claim external audit or certification.
    Do not claim CarbonStack Relay Space mechanics are implemented.
    Do not claim local operator profile is complete.
    Do not claim v0.3.27 is a public release.

Updated allowed claim:

    CarbonStack v0.3.27 records a local-only Cypher explicit-env API lifecycle proof.
    The proof starts Cypher on loopback with a temporary SQLite DB, claims invites, registers devices, submits/retrieves/acks an opaque OpenMLS application-message envelope, restarts against the same DB, and verifies persisted state remains available.
    WSL Debian doctor/core validation passed after the v0.3.27 docs commit.
    Final clean status showed no dirty files across the four repos after expected artifact cleanup.

---

## 7. Next Safest Actions

Recommended next sequence after this v0.3.27 breakpoint:

    v0.3.28 preferred:
      local operator lifecycle proof review and helper/runner decision point

Recommended concrete next rung:

    1. Review v0.3.27 proof and decide whether helper tooling is justified yet.
    2. Decide whether a local-backbone runner profile has a concrete enough success/failure contract.
    3. Decide whether the runner should start/stop Cypher or target an already-running process.
    4. Decide what DB lifecycle a future runner profile should use.
    5. Decide whether additional negative-path API lifecycle proof is needed first.
    6. Keep public ingress/systemd/cloudflared/homelab deferred.
    7. Keep runtime Comms UX deferred to v0.4.x.

Do **not** jump directly to:

    public ingress
    cloudflared
    systemd
    real homelab
    runtime Comms send/inbox OpenMLS UX
    Android app
    CarbonStackOS
    remote admin plane
    CarbonStack Relay Space implementation
    production release
    public deployment claims

---

## 8. Critical Paths / Functions

### WSL Debian active working umbrella

    [WSL-DEBIAN] umbrella: ~/repos/carbonstack_umbrella
    [WSL-DEBIAN] carbonstack: ~/repos/carbonstack_umbrella/carbonstack
    [WSL-DEBIAN] carbonstack-comms: ~/repos/carbonstack_umbrella/carbonstack-comms
    [WSL-DEBIAN] carbonstack-cypher: ~/repos/carbonstack_umbrella/carbonstack-cypher
    [WSL-DEBIAN] carbonstack-os: ~/repos/carbonstack_umbrella/carbonstack-os
    [GO-RUNNER] runner dir: ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate

### WSL Debian known-good validation commands

From runner dir:

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go run . --profile doctor
    go run . --profile core

From carbonstack repo root:

    cd ~/repos/carbonstack_umbrella/carbonstack
    go run ./tools/carbonstack-validate --profile doctor
    go run ./tools/carbonstack-validate --profile core

### Artifact cleanup after validation

    cd ~/repos/carbonstack_umbrella/carbonstack-comms
    rm -rf internal/protocol/mls/openmls-sidecar/.carbonstack-openmls-sidecar-state
    rm -rf internal/protocol/mls/openmls-sidecar/target

### v0.3.27 docs

    [DOC] v0.3.27 proof:
    carbonstack/docs/146-local-cypher-explicit-env-api-lifecycle-proof-v0.md

    [DOC] docs index:
    carbonstack/docs/README.md

    [DOC] roadmap:
    carbonstack/roadmap/ROADMAP.md

### Local operator convention paths

    [LOCAL OPERATOR] data directory:
    ~/.local/share/carbonstack/cypher

    [LOCAL OPERATOR] DB path:
    ~/.local/share/carbonstack/cypher/cypher.db

    [LOCAL OPERATOR] config directory:
    ~/.config/carbonstack/cypher

    [LOCAL OPERATOR] source-tree migrations path:
    ~/repos/carbonstack_umbrella/carbonstack-cypher/migrations

### v0.3.27 temporary proof paths

    [LOCAL PROOF] temporary data dir:
    ~/.local/share/carbonstack/cypher-v0327-lifecycle-proof

    [LOCAL PROOF] temporary DB:
    ~/.local/share/carbonstack/cypher-v0327-lifecycle-proof/cypher.db

    [LOCAL PROOF] temporary binary:
    ~/.local/share/carbonstack/cypher-v0327-lifecycle-proof/carbonstack-cypher-lifecycle-proof

### Recommended local-only operator commands

    cd ~/repos/carbonstack_umbrella/carbonstack-cypher
    mkdir -p ~/.local/share/carbonstack/cypher
    mkdir -p ~/.config/carbonstack/cypher

    CYPHER_ADDR=127.0.0.1:8080 \
    CYPHER_DB=$HOME/.local/share/carbonstack/cypher/cypher.db \
    CYPHER_MIGRATIONS=$HOME/repos/carbonstack_umbrella/carbonstack-cypher/migrations \
    CYPHER_DEV_INVITE=dev-invite \
    go run ./cmd/cypher

Health check:

    curl -i http://127.0.0.1:8080/v0/health

Reset:

    rm -f ~/.local/share/carbonstack/cypher/cypher.db

### v0.3.27 API lifecycle contract

    health:
      GET /v0/health

    claim seeded invite:
      POST /v0/invites/claim
      invite_code=dev-invite
      display_name=<name>

    create dev invite:
      POST /v0/dev/invites
      invite_code=<code>

    register device:
      POST /v0/devices/register
      account_id=<account_id>
      device_label=<label>
      public_identity_key=<stub or real future key>
      public_prekey_bundle=<stub or real future prekey bundle>

    list account devices:
      GET /v0/accounts/<account_id>/devices

    submit opaque OpenMLS application-message envelope:
      POST /v0/envelopes
      sender_device_id=<device_id>
      recipient_device_id=<device_id>
      content_type=carbonstack.mls.application-message.v0
      protocol_version=carbonstack-openmls-sidecar-v0
      ciphertext_b64=<base64 opaque payload>

    retrieve recipient inbox:
      GET /v0/devices/<device_id>/envelopes

    ack envelope:
      POST /v0/envelopes/<envelope_id>/ack
      recipient_device_id=<recipient_device_id>

### Critical functions / behavior preserved from v0.3.24

    Store.Migrate:
      ensures schema_migrations exists;
      reads migration .sql files in sorted order;
      computes sha256 checksum for each migration;
      skips matching already-applied migrations;
      fails on checksum mismatch for already-applied migration;
      applies new migrations in a transaction;
      records migration after successful application.

    TestMigrateRecordsAppliedMigrations:
      proves applied migrations are recorded with non-empty checksum/applied_at.

    TestMigrateCanRunTwiceAgainstSameDB:
      proves repeated migration against the same DB now succeeds.

    TestMigrateDetectsAppliedMigrationChecksumMismatch:
      proves already-applied migration drift hard-fails.

---

## 9. Lean Breakpoint Summary

    CarbonStack v0.3.27 is a stable mainline docs/proof breakpoint after v0.3.26. It records a local-only Cypher explicit-env API lifecycle proof in carbonstack/docs/146, using loopback bind, temporary SQLite DB, source-tree migrations path, seeded dev invite, invite claim, second dev invite creation/claim, device registration/listing, opaque OpenMLS application-message envelope submit/retrieve/ack, restart against the same DB, and post-restart persistence checks. It preserves the wrong protocol/content-type blunder from the first failed proof attempt, validates WSL Debian doctor/core after the docs commit, cleans generated OpenMLS artifacts, and leaves helper tooling, local-backbone runner profile, runtime Comms OpenMLS UX, CarbonStack Relay Space mechanics, public ingress, systemd, cloudflared, homelab, Android, CarbonStackOS, hostile-server proof, secure vault, and production/security claims deferred.

---

## 10. Preserved v0.3.26 Operational Process Log

The following section preserves the previous v0.3.26 LogDoc body as operational continuity. It is intentionally retained rather than compressed away. Where v0.3.26 text conflicts with the v0.3.27 current-state overlay above, v0.3.27 wins for current state; v0.3.26 remains the provenance/process ledger for the local operator config/data convention and the preserved v0.3.25 / v0.3.24 / v0.3.23 / v0.3.22 / v0.3.21 / v0.3.20 history.



# CarbonStack LogDoc v0.3.26

**Last updated:** 2026-06-02 local session  
**Owner/operator:** bitcrusher32  
**Assistant/model:** GPT-5.5 Thinking  
**Current phase:** Phase 2E-prep / v0.3.x **local operator config/data convention checkpoint**. `carbonstack` is now at `6e14fe0 docs: record local operator config data convention`, after `v0.3.25` landed the local Cypher operator runbook skeleton at `a5c6351` and `carbonstack-cypher` landed the README pointer at `9ab994c`. `carbonstack-comms` remains at `012c8bf scripts: support source snapshot artifact guard`; `carbonstack-cypher` remains at `9ab994c docs: point to local operator runbook`; `carbonstack-os` remains at `1bbbe52 docs: clarify CarbonStackOS target direction`. The live WSL Debian working umbrella at `~/repos/carbonstack_umbrella` passed `doctor` and `core` after the v0.3.26 docs commit, expected OpenMLS sidecar artifacts were cleaned afterward, and final clean status showed no dirty files across the four repos.

**Version schema:** `v[scope].[timeline]`. This file is `v0.3.26`, the local operator config/data convention handoff after the `v0.3.25` local Cypher operator runbook skeleton checkpoint. Per LogDoc V2 practice, this Markdown file intentionally preserves verbose operational continuity; the JSON breakpoint remains the lean handoff state. The dense v0.3.25, v0.3.24, v0.3.23, v0.3.22, v0.3.21, and v0.3.20 operational process logs are preserved later in this file rather than compressed away.

**LogDoc V2 usage note:** LogDoc V2 uses a split between a Markdown LogDoc for working-session continuity, a JSON state file for end-of-session handoff, Git for history preservation, and README/docs for public claims. This v0.3.26 Markdown file therefore intentionally stays “bloated” where operational process, blunders, validation ladders, and historical context matter. The v0.3.26 JSON should remain lean and should agree with this file on the latest current state.

---

## 1. Project Goal

**Active goal:** Preserve CarbonStack state after recording the local-only Cypher operator config/data convention. The immediate next line of work is an explicit local Cypher API lifecycle proof under the documented operator environment, not runtime Comms UX, not public ingress, not systemd/cloudflared, not real homelab deployment, and not Android/CarbonStackOS work.

CarbonStack’s broader project intent remains:

- **CarbonStack** is the secure-communications appliance-stack project and shared doctrine/release authority.
- **CarbonStackOS** is the future constrained Android-derived appliance OS.
- **CarbonStackComms** is the text-first encrypted messaging client and current OpenMLS sidecar/relay integration surface.
- **CarbonStackCypher** is the self-hostable relay/storage stack.
- The `carbonstack` repo remains the doctrine/spec/docs/release authority; implementation repos must not contradict it without revising the spec.

**North star:**

> Its goal is not to be a better smartphone. Its goal is to stop being a smartphone.

**Core doctrine:**

> Every feature is guilty until it proves it does not add unacceptable parser, network, sensor, identity, or filesystem authority.

**Current reality after v0.3.26:** `v0.3.20` remains the current public runner-backed testing release. `v0.3.21` defined the local-only deployability model. `v0.3.22` narrowed the local deployability surface around Cypher operator config/API/database assumptions and adopted **CarbonStack Relay Space** terminology. `v0.3.23` proved the SQLite repeated-migration persistence hazard. `v0.3.24` implemented `schema_migrations` tracking in CarbonStackCypher. `v0.3.25` recorded the first local-only Cypher operator runbook skeleton. `v0.3.26` now records the local operator config/data convention: explicit loopback bind, DB path outside the source tree, source-tree migrations path for WSL Debian development, blunt reset semantics, and explicit deferral of config parsers, helper commands, and local-backbone runner profiles.

The current public release page remains:

    https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.3.20

The current public testing interpretation remains:

    v0.3.20 is the preferred public runner-backed testing release.
    v0.3.4 remains available as the older Windows 11 / PowerShell legacy testbed.
    Debian / WSL Debian is now the mainline public dev/test validation direction.
    Windows 11 short-path validation was the final explicit Windows dev/test validation for this phase.

This is still **not** a deployability release. It does **not** prove production readiness, production E2EE, hostile-server safety, metadata privacy, Android readiness, secure vault/storage, Debian production deployability, systemd readiness, cloudflared readiness, public ingress safety, real homelab validation, external audit, or certification.

---

## 2. Current Repo Heads

    carbonstack        6e14fe0 (HEAD -> main, origin/main, origin/HEAD) docs: record local operator config data convention
    carbonstack-comms  012c8bf (HEAD -> main, origin/main, origin/HEAD) scripts: support source snapshot artifact guard
    carbonstack-cypher 9ab994c (HEAD -> main, origin/main, origin/HEAD) docs: point to local operator runbook
    carbonstack-os     1bbbe52 (HEAD -> main, origin/main, origin/HEAD) docs: clarify CarbonStackOS target direction

Current `carbonstack` public runner-backed testing release tag:

    v0.3.20 -> 3eeb1a1e1a / 3eeb1a1 docs: clean public release surface

Previous mainline deployability checkpoints:

    v0.3.25 mainline -> a5c6351 docs: update readme, roadmap, and historical doc
    v0.3.25 cypher  -> 9ab994c docs: point to local operator runbook
    v0.3.24 mainline -> 76f0258 docs: record Cypher schema migrations implementation
    v0.3.24 cypher  -> 42f838f db: track applied migrations
    v0.3.23 mainline -> 1fef7fe docs: record Cypher migration persistence recon
    v0.3.22 mainline -> 328ffc4 docs: record Cypher local operator surface recon
    v0.3.21 mainline -> 6850c58 docs: record local backbone deployability recon

Legacy `carbonstack` Windows/PowerShell testbed release tag:

    v0.3.4 -> 7c5b7e6 docs: record release snapshot self-test fix

Historical `carbonstack` release tag:

    v0.3.0 -> 59245fc docs: freeze v0.3.0 release packaging plan

Important continuity note: `v0.3.26` does not supersede `v0.3.20` as a public release. It advances mainline local-only deployability documentation after the v0.3.25 runbook skeleton. It is safe to pause here because the convention doc was pushed, WSL Debian `doctor`/`core` passed after the commit, generated artifacts were cleaned, and final status was clean across all four repos.

---

## 3. Current Validated State

Validated at the v0.3.26 checkpoint:

    [WSL-DEBIAN] working umbrella exists at ~/repos/carbonstack_umbrella
    [WSL-DEBIAN] carbonstack is at 6e14fe0
    [WSL-DEBIAN] carbonstack-comms is at 012c8bf
    [WSL-DEBIAN] carbonstack-cypher is at 9ab994c
    [WSL-DEBIAN] carbonstack-os is at 1bbbe52
    [CARBONSTACK] docs/145 local operator config/data convention committed and pushed
    [GO-RUNNER] doctor passed after the v0.3.26 docs commit
    [GO-RUNNER] core passed after the v0.3.26 docs commit
    [ARTIFACT CLEANUP] expected OpenMLS generated roots were removed after validation
    [FINAL STATUS] all four repos were clean after cleanup and commit

Observed WSL Debian toolchain baseline:

    git version 2.47.3
    go version go1.24.4 linux/amd64
    rustc 1.96.0 (ac68faa20 2026-05-25)
    cargo 1.96.0 (30a34c682 2026-05-25)
    sqlite3 3.46.1

Observed `core` validation result during the v0.3.26 flow:

    pre-test artifact scan: PASS / no generated/private/build artifact hits
    OpenMLS real-Cypher lifecycle: PASS
    carbonstack-comms package tests: PASS
    carbonstack-cypher package tests: PASS
    post-test artifact scan: expected OpenMLS sidecar generated roots only
    VALIDATION PASSED

Expected post-test generated roots:

    carbonstack-comms/internal/protocol/mls/openmls-sidecar/.carbonstack-openmls-sidecar-state
    carbonstack-comms/internal/protocol/mls/openmls-sidecar/target

These roots are expected only after tests and must not be committed or treated as source artifacts. They were cleaned before the final clean status snapshot.

---

## 4. v0.3.26 Work Completed

### 4.1 Local operator config/data convention landed

New/updated docs at v0.3.26:

    carbonstack/docs/145-local-operator-config-data-convention-v0.md
    carbonstack/docs/README.md
    carbonstack/roadmap/ROADMAP.md

Commit:

    6e14fe0 docs: record local operator config data convention

The convention doc records the next refinement after the v0.3.25 runbook skeleton: how local operator state should be placed and reasoned about before helper tooling or runner profiles exist.

### 4.2 Local-only data/config convention is now explicit

v0.3.26 standardizes the current WSL Debian local-operator convention:

    Recommended local-only address:
      127.0.0.1:8080

    Recommended local operator data directory:
      $HOME/.local/share/carbonstack/cypher

    Recommended local operator DB path:
      $HOME/.local/share/carbonstack/cypher/cypher.db

    Recommended local operator config directory:
      $HOME/.config/carbonstack/cypher

    Recommended source-tree migrations path during WSL Debian development:
      $HOME/repos/carbonstack_umbrella/carbonstack-cypher/migrations

    Recommended current config style:
      explicit environment variables

The current explicit operator env remains:

    CYPHER_ADDR=127.0.0.1:8080
    CYPHER_DB=$HOME/.local/share/carbonstack/cypher/cypher.db
    CYPHER_MIGRATIONS=$HOME/repos/carbonstack_umbrella/carbonstack-cypher/migrations
    CYPHER_DEV_INVITE=dev-invite

### 4.3 Source tree state and operator state are now separated by doctrine

v0.3.26 records the rule that source tree state and operator state should stay separate.

Source tree should contain:

    code
    tests
    migrations
    docs
    validation runner inputs

Operator state should contain:

    local SQLite DB
    local operator config notes/env
    temporary logs if added later
    local reset target

The DB should not live at:

    carbonstack-cypher/cypher.db

Preferred current path:

    $HOME/.local/share/carbonstack/cypher/cypher.db

Reason:

    Keep live local operator DB state outside Git repos.
    Keep source tree inspectable, commit-safe, package-safe, and validation-safe.
    Keep operator reset simple and explicit.

### 4.4 Config parser, helper command, and runner profile intentionally deferred

v0.3.26 explicitly defers a config parser.

Reasoning recorded:

    env vars are already supported
    local operator semantics are still maturing
    TOML/YAML/JSON config introduces parser and UX surface
    v0.3.x is local-only deployability groundwork, not mature packaging

v0.3.26 also defers helper tooling.

Possible future helper shape remains:

    carbonstack local start
    carbonstack local health
    carbonstack local stop
    carbonstack local reset

But helper tooling should follow a validated manual convention rather than invent one.

v0.3.26 also defers a local-backbone runner profile because its contract is not yet concrete enough. The missing contract pieces remain:

    stable local operator startup convention
    stable health check expectation
    stable reset semantics
    stable generated artifact boundaries
    decision on whether runner starts/stops Cypher or targets an already-running process
    decision on DB lifecycle during validation

### 4.5 Next rung narrowed to explicit-env API lifecycle proof

v0.3.26 recommends:

    v0.3.27 local Cypher explicit-env API lifecycle proof

Suggested proof scope:

    start Cypher with explicit local operator env
    verify /v0/health
    create dev invite
    claim invite
    register device
    submit envelope
    retrieve envelope
    ack envelope
    stop Cypher
    restart Cypher against the same DB
    verify persisted state where appropriate

This should still be local-only and pre-alpha. It should not become a public deployment guide.

### 4.6 Final validation and cleanup completed

Final WSL Debian repo snapshot:

    carbonstack        6e14fe0 docs: record local operator config data convention
    carbonstack-comms  012c8bf scripts: support source snapshot artifact guard
    carbonstack-cypher 9ab994c docs: point to local operator runbook
    carbonstack-os     1bbbe52 docs: clarify CarbonStackOS target direction

Final validation:

    go run . --profile doctor
    go run . --profile core

Final validation passed under WSL Debian. `core` passed OpenMLS real-Cypher lifecycle, Comms package tests, and Cypher package tests, including `internal/db` package tests. After validation, expected OpenMLS generated roots were removed from `carbonstack-comms`, and final clean status showed no dirty files across the four repos.

---

## 5. Critical Blunders / Lessons Preserved

### 5.1 v0.3.26 was correctly docs-only

No code changes were needed for this rung. The convention needed to exist before helper tooling or runner profiles. This avoided turning a documentation/rationale rung into premature tooling.

Correct lesson:

    Manual convention first.
    Helper tooling second.
    Runner profile only after the success/failure contract is concrete.

### 5.2 Source-tree DB state remains a trap to avoid

The convention makes explicit that `carbonstack-cypher/cypher.db` is not the desired operator DB location. Keeping DB state in the repo would increase risk of accidental commits, stale state confusion, release artifact contamination, and unclear reset semantics.

Correct lesson:

    Operator state belongs outside the source tree.
    Source repos should stay source/package/test surfaces.
    Local operator DBs should be obvious reset targets.

### 5.3 `CYPHER_ADDR=:8080` is still a code default, not the local-operator recommendation

The code default still exists, but the local operator convention continues to use:

    CYPHER_ADDR=127.0.0.1:8080

Correct lesson:

    Do not describe the current default as local-only.
    Use explicit env vars for operator experiments.
    Consider changing defaults later only as its own focused code/config rung.

### 5.4 Config parser deferral is security-consistent

Deferring TOML/YAML/JSON config parsing is not laziness; it is consistent with CarbonStack doctrine. New parser surfaces should not be added until the semantics are stable enough to justify them.

Correct lesson:

    Environment variables are boring and sufficient for now.
    Parser surface should follow stable operator semantics, not precede them.

### 5.5 Validation and cleanup cadence held correctly

The v0.3.26 flow validated after the commit, then cleaned expected generated OpenMLS roots, then captured final clean status. This is the stricter habit that should be preserved for future breakpoints.

---

## 6. Current Blockers / Not Validated

Current blockers and nonvalidated areas:

    Local-only operator convention is documented, not enforced by code.
    Local operator full API lifecycle proof is not done yet.
    Cypher default bind remains :8080; local-only docs recommend explicit 127.0.0.1:8080.
    No local-backbone runner profile exists yet.
    No helper command exists yet.
    No config file parser exists yet.
    Runtime Comms send/inbox remains stub-era and not wired as user-facing OpenMLS messenger UX.
    Production-safe provider storage and secure vault design are unsolved.
    Hostile-server rollback/replay/metadata-abuse validation remains future work.
    Android/CarbonStackOS implementation remains later work.
    CarbonStack Relay Space mechanics remain a defined future concept, not an implementation.
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
    Do not claim secure vault/storage.
    Do not claim external audit or certification.
    Do not claim CarbonStack Relay Space mechanics are implemented.
    Do not claim local operator profile is complete.
    Do not claim v0.3.26 is a public release.

Updated allowed claim:

    CarbonStack v0.3.26 records the local operator config/data convention.
    The convention recommends explicit 127.0.0.1 loopback binding, a DB path outside the source tree, explicit source-tree migrations path during WSL Debian development, blunt DB reset semantics, and continued deferral of config parser/helper/runner-profile work.
    WSL Debian doctor/core validation passed after the v0.3.26 docs commit.

---

## 7. Next Safest Actions

Recommended next sequence after this v0.3.26 breakpoint:

    v0.3.27 preferred:
      local Cypher explicit-env API lifecycle proof

Recommended concrete next rung:

    1. Start Cypher with explicit local operator env.
    2. Verify /v0/health.
    3. Create a dev invite or use explicit dev invite.
    4. Claim invite.
    5. Register a device.
    6. Submit an envelope.
    7. Retrieve envelope.
    8. Ack envelope.
    9. Stop Cypher.
    10. Restart Cypher against the same DB.
    11. Verify persisted DB state where appropriate.
    12. Document the proof in carbonstack/docs.
    13. Keep it local-only and pre-alpha.

Do **not** jump directly to:

    public ingress
    cloudflared
    systemd
    real homelab
    runtime Comms send/inbox OpenMLS UX
    Android app
    CarbonStackOS
    remote admin plane
    CarbonStack Relay Space implementation
    local-backbone runner profile

---

## 8. Critical Paths / Functions

### WSL Debian active working umbrella

    [WSL-DEBIAN] umbrella: ~/repos/carbonstack_umbrella
    [WSL-DEBIAN] carbonstack: ~/repos/carbonstack_umbrella/carbonstack
    [WSL-DEBIAN] carbonstack-comms: ~/repos/carbonstack_umbrella/carbonstack-comms
    [WSL-DEBIAN] carbonstack-cypher: ~/repos/carbonstack_umbrella/carbonstack-cypher
    [WSL-DEBIAN] carbonstack-os: ~/repos/carbonstack_umbrella/carbonstack-os
    [GO-RUNNER] runner dir: ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate

### WSL Debian known-good validation commands

From runner dir:

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go run . --profile doctor
    go run . --profile core

From carbonstack repo root:

    cd ~/repos/carbonstack_umbrella/carbonstack
    go run ./tools/carbonstack-validate --profile doctor
    go run ./tools/carbonstack-validate --profile core

### Artifact cleanup after validation

    cd ~/repos/carbonstack_umbrella/carbonstack-comms
    rm -rf internal/protocol/mls/openmls-sidecar/.carbonstack-openmls-sidecar-state
    rm -rf internal/protocol/mls/openmls-sidecar/target

### v0.3.26 docs

    [DOC] v0.3.26 convention:
    carbonstack/docs/145-local-operator-config-data-convention-v0.md

    [DOC] docs index:
    carbonstack/docs/README.md

    [DOC] roadmap:
    carbonstack/roadmap/ROADMAP.md

### Local operator convention paths

    [LOCAL OPERATOR] data directory:
    ~/.local/share/carbonstack/cypher

    [LOCAL OPERATOR] DB path:
    ~/.local/share/carbonstack/cypher/cypher.db

    [LOCAL OPERATOR] config directory:
    ~/.config/carbonstack/cypher

    [LOCAL OPERATOR] source-tree migrations path:
    ~/repos/carbonstack_umbrella/carbonstack-cypher/migrations

### Recommended local-only operator commands

    cd ~/repos/carbonstack_umbrella/carbonstack-cypher
    mkdir -p ~/.local/share/carbonstack/cypher
    mkdir -p ~/.config/carbonstack/cypher

    CYPHER_ADDR=127.0.0.1:8080     CYPHER_DB=$HOME/.local/share/carbonstack/cypher/cypher.db     CYPHER_MIGRATIONS=$HOME/repos/carbonstack_umbrella/carbonstack-cypher/migrations     CYPHER_DEV_INVITE=dev-invite     go run ./cmd/cypher

Health check:

    curl -i http://127.0.0.1:8080/v0/health

Reset:

    rm -f ~/.local/share/carbonstack/cypher/cypher.db

### Critical functions / behavior preserved from v0.3.24

    Store.Migrate:
      ensures schema_migrations exists;
      reads migration .sql files in sorted order;
      computes sha256 checksum for each migration;
      skips matching already-applied migrations;
      fails on checksum mismatch for already-applied migration;
      applies new migrations in a transaction;
      records migration after successful application.

    TestMigrateRecordsAppliedMigrations:
      proves applied migrations are recorded with non-empty checksum/applied_at.

    TestMigrateCanRunTwiceAgainstSameDB:
      proves repeated migration against the same DB now succeeds.

    TestMigrateDetectsAppliedMigrationChecksumMismatch:
      proves already-applied migration drift hard-fails.

---

## 9. Lean Breakpoint Summary

    CarbonStack v0.3.26 is a stable mainline docs/convention breakpoint after v0.3.25. It records the local-only operator config/data convention in carbonstack/docs/145, standardizes explicit 127.0.0.1 loopback bind, DB path outside the source tree, source-tree migrations path for WSL Debian development, config directory, blunt DB reset semantics, and continued deferral of config parser/helper/local-backbone runner profile work; validates WSL Debian doctor/core after the docs commit; cleans generated OpenMLS artifacts; and leaves explicit-env API lifecycle proof, local-backbone runner profile, runtime Comms OpenMLS UX, CarbonStack Relay Space mechanics, public ingress, systemd, cloudflared, homelab, Android, CarbonStackOS, hostile-server proof, secure vault, and production/security claims deferred.

---

## 10. Preserved v0.3.25 Operational Process Log

The following section preserves the previous v0.3.25 LogDoc body as operational continuity. It is intentionally retained rather than compressed away. Where v0.3.25 text conflicts with the v0.3.26 current-state overlay above, v0.3.26 wins for current state; v0.3.25 remains the provenance/process ledger for the local Cypher operator runbook skeleton and the preserved v0.3.24 / v0.3.23 / v0.3.22 / v0.3.21 / v0.3.20 history.




# CarbonStack LogDoc v0.3.25

**Last updated:** 2026-06-02 local session  
**Owner/operator:** bitcrusher32  
**Assistant/model:** GPT-5.5 Thinking  
**Current phase:** Phase 2E-prep / v0.3.x **local Cypher operator runbook skeleton checkpoint**. `carbonstack` is now at `a5c6351 docs: update readme, roadmap, and historical doc`, after `v0.3.24` landed Cypher schema migration tracking docs at `76f0258`. `carbonstack-cypher` is now at `9ab994c docs: point to local operator runbook`, after `42f838f db: track applied migrations` implemented `schema_migrations`. `carbonstack-comms` remains at `012c8bf scripts: support source snapshot artifact guard`; `carbonstack-os` remains at `1bbbe52 docs: clarify CarbonStackOS target direction`. The live WSL Debian working umbrella at `~/repos/carbonstack_umbrella` passed `doctor` and `core` during the v0.3.25 runbook work, expected OpenMLS sidecar artifacts were cleaned afterward, and final clean status showed no dirty files across the four repos after the CarbonStack docs commit.

**Version schema:** `v[scope].[timeline]`. This file is `v0.3.25`, the local Cypher operator runbook skeleton handoff after the `v0.3.24` Cypher schema migration tracking implementation checkpoint. Per LogDoc V2 practice, this Markdown file intentionally preserves verbose operational continuity; the JSON breakpoint remains the lean handoff state. The dense v0.3.24, v0.3.23, v0.3.22, v0.3.21, and v0.3.20 operational process logs are preserved later in this file rather than compressed away.

**LogDoc V2 usage note:** LogDoc V2 uses a split between a Markdown LogDoc for working-session continuity, a JSON state file for end-of-session handoff, Git for history preservation, and README/docs for public claims. This v0.3.25 Markdown file therefore intentionally stays “bloated” where operational process, blunders, validation ladders, and historical context matter. The v0.3.25 JSON should remain lean and should agree with this file on the latest current state.

---

## 1. Project Goal

**Active goal:** Preserve CarbonStack state after recording the first local-only Cypher operator runbook skeleton and pointing `carbonstack-cypher` at that runbook. The immediate next line of work is to refine the local operator config/data-path convention and/or run a more explicit local operator lifecycle proof, not runtime Comms UX, not public ingress, not systemd/cloudflared, not real homelab deployment, and not Android/CarbonStackOS work.

CarbonStack’s broader project intent remains:

- **CarbonStack** is the secure-communications appliance-stack project and shared doctrine/release authority.
- **CarbonStackOS** is the future constrained Android-derived appliance OS.
- **CarbonStackComms** is the text-first encrypted messaging client and current OpenMLS sidecar/relay integration surface.
- **CarbonStackCypher** is the self-hostable relay/storage stack.
- The `carbonstack` repo remains the doctrine/spec/docs/release authority; implementation repos must not contradict it without revising the spec.

**North star:**

> Its goal is not to be a better smartphone. Its goal is to stop being a smartphone.

**Core doctrine:**

> Every feature is guilty until it proves it does not add unacceptable parser, network, sensor, identity, or filesystem authority.

**Current reality after v0.3.25:** `v0.3.20` remains the current public runner-backed testing release. `v0.3.21` defined the local-only deployability model. `v0.3.22` narrowed the local deployability surface around Cypher operator config/API/database assumptions and adopted **CarbonStack Relay Space** terminology. `v0.3.23` proved the SQLite repeated-migration persistence hazard. `v0.3.24` implemented `schema_migrations` tracking in CarbonStackCypher. `v0.3.25` records the first local-only Cypher operator runbook skeleton, with explicit loopback startup, explicit SQLite DB path, explicit migrations path, reset semantics, and a health-check proof captured in docs.

The current public release page remains:

    https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.3.20

The current public testing interpretation remains:

    v0.3.20 is the preferred public runner-backed testing release.
    v0.3.4 remains available as the older Windows 11 / PowerShell legacy testbed.
    Debian / WSL Debian is now the mainline public dev/test validation direction.
    Windows 11 short-path validation was the final explicit Windows dev/test validation for this phase.

This is still **not** a deployability release. It does **not** prove production readiness, production E2EE, hostile-server safety, metadata privacy, Android readiness, secure vault/storage, Debian production deployability, systemd readiness, cloudflared readiness, public ingress safety, real homelab validation, external audit, or certification.

---

## 2. Current Repo Heads

    carbonstack        a5c6351 docs: update readme, roadmap, and historical doc
    carbonstack-comms  012c8bf scripts: support source snapshot artifact guard
    carbonstack-cypher 9ab994c docs: point to local operator runbook
    carbonstack-os     1bbbe52 docs: clarify CarbonStackOS target direction

Current `carbonstack` public runner-backed testing release tag:

    v0.3.20 -> 3eeb1a1e1a / 3eeb1a1 docs: clean public release surface

Previous mainline deployability checkpoints:

    v0.3.24 mainline -> 76f0258 docs: record Cypher schema migrations implementation
    v0.3.24 cypher  -> 42f838f db: track applied migrations
    v0.3.23 mainline -> 1fef7fe docs: record Cypher migration persistence recon
    v0.3.22 mainline -> 328ffc4 docs: record Cypher local operator surface recon
    v0.3.21 mainline -> 6850c58 docs: record local backbone deployability recon

Legacy `carbonstack` Windows/PowerShell testbed release tag:

    v0.3.4 -> 7c5b7e6 docs: record release snapshot self-test fix

Historical `carbonstack` release tag:

    v0.3.0 -> 59245fc docs: freeze v0.3.0 release packaging plan

Important continuity note: `v0.3.25` does not supersede `v0.3.20` as a public release. It advances mainline local-only deployability documentation after the v0.3.24 schema migration tracking implementation. It is safe to pause here because the README/doc pointers were pushed, WSL Debian `doctor`/`core` passed during the rung, generated artifacts were cleaned, and final status was clean across all four repos.

---

## 3. Current Validated State

Validated at the v0.3.25 checkpoint:

    [WSL-DEBIAN] working umbrella exists at ~/repos/carbonstack_umbrella
    [WSL-DEBIAN] carbonstack is at a5c6351
    [WSL-DEBIAN] carbonstack-comms is at 012c8bf
    [WSL-DEBIAN] carbonstack-cypher is at 9ab994c
    [WSL-DEBIAN] carbonstack-os is at 1bbbe52
    [CYPHER] README pointer to local operator runbook committed and pushed
    [CARBONSTACK] docs/144 local Cypher operator runbook skeleton committed and pushed
    [GO-RUNNER] doctor passed during v0.3.25 runbook work
    [GO-RUNNER] core passed during v0.3.25 runbook work
    [ARTIFACT CLEANUP] expected OpenMLS generated roots were removed after validation
    [FINAL STATUS] all four repos were clean after cleanup and commits

Observed WSL Debian toolchain baseline:

    git version 2.47.3
    go version go1.24.4 linux/amd64
    rustc 1.96.0 (ac68faa20 2026-05-25)
    cargo 1.96.0 (30a34c682 2026-05-25)
    sqlite3 3.46.1

Observed `core` validation result during the v0.3.25 flow:

    pre-test artifact scan: PASS / no generated/private/build artifact hits
    OpenMLS real-Cypher lifecycle: PASS
    carbonstack-comms package tests: PASS
    carbonstack-cypher package tests: PASS
    post-test artifact scan: expected OpenMLS sidecar generated roots only
    VALIDATION PASSED

Expected post-test generated roots:

    carbonstack-comms/internal/protocol/mls/openmls-sidecar/.carbonstack-openmls-sidecar-state
    carbonstack-comms/internal/protocol/mls/openmls-sidecar/target

These roots are expected only after tests and must not be committed or treated as source artifacts. They were cleaned before the final clean status snapshot.

Note: the final uploaded snapshot shows validation run while `carbonstack` docs were still dirty, then generated artifacts were cleaned, then the `carbonstack` docs were committed and pushed. Since the final commit was docs-only and the same docs changes were already present during the validation run, this remains a valid local docs checkpoint. A future release-grade validation should still rerun validation after final commits as a stricter habit.

---

## 4. v0.3.25 Work Completed

### 4.1 Local Cypher operator runbook skeleton landed

New/updated docs at v0.3.25:

    carbonstack/docs/144-local-cypher-operator-runbook-skeleton-v0.md
    carbonstack/docs/README.md
    carbonstack/roadmap/ROADMAP.md

Commit:

    a5c6351 docs: update readme, roadmap, and historical doc

The commit message is generic, but the content is the v0.3.25 local-only Cypher operator runbook skeleton.

The runbook records:

    explicit loopback-only Cypher startup;
    experimental persistent local SQLite DB path;
    explicit migrations path;
    explicit dev invite/bootstrap setting;
    health-check command;
    reset/cleanup command;
    current persistence boundary after schema_migrations;
    non-goals and nonclaims;
    next rungs.

Recommended local-only startup in docs:

    cd ~/repos/carbonstack_umbrella/carbonstack-cypher
    mkdir -p ~/.local/share/carbonstack/cypher

    CYPHER_ADDR=127.0.0.1:8080 \
    CYPHER_DB=$HOME/.local/share/carbonstack/cypher/cypher.db \
    CYPHER_MIGRATIONS=$HOME/repos/carbonstack_umbrella/carbonstack-cypher/migrations \
    CYPHER_DEV_INVITE=dev-invite \
    go run ./cmd/cypher

Recommended health check:

    curl -i http://127.0.0.1:8080/v0/health

Recommended blunt reset command:

    rm -f ~/.local/share/carbonstack/cypher/cypher.db

### 4.2 Minimal health-check proof captured in docs

A temporary proof built/launched Cypher as a local binary with isolated v0.3.25 health-check state, using:

    CYPHER_ADDR=127.0.0.1:18080
    CYPHER_DB=$HOME/.local/share/carbonstack/cypher-v0325-healthcheck/cypher.db
    CYPHER_MIGRATIONS=$HOME/repos/carbonstack_umbrella/carbonstack-cypher/migrations
    CYPHER_DEV_INVITE=v0.3.25-healthcheck-dev-invite

The proof checked:

    http://127.0.0.1:18080/v0/health

The exact health response text is recorded inside `docs/144-local-cypher-operator-runbook-skeleton-v0.md`. The temporary proof DB and binary were intended to be deleted after the proof. The final repo status did not show generated health-check artifacts inside project repos.

### 4.3 carbonstack-cypher README pointer landed

Cypher docs commit:

    9ab994c docs: point to local operator runbook

Updated file:

    carbonstack-cypher/README.md

The README now points readers to:

    carbonstack/docs/144-local-cypher-operator-runbook-skeleton-v0.md

It also warns that local operator usage should use explicit settings such as `CYPHER_ADDR=127.0.0.1:8080`, explicit `CYPHER_DB`, and explicit `CYPHER_MIGRATIONS`, and that this remains pre-alpha local deployability work rather than a production deployment guide.

### 4.4 Final validation and cleanup completed

During the v0.3.25 flow, WSL Debian runner validation passed:

    go run . --profile doctor
    go run . --profile core

`core` passed:

    OpenMLS real-Cypher lifecycle
    carbonstack-comms package tests
    carbonstack-cypher package tests, including internal/db tests

Expected OpenMLS sidecar generated roots were removed after validation. Final clean status showed:

    carbonstack        a5c6351 docs: update readme, roadmap, and historical doc
    carbonstack-comms  012c8bf scripts: support source snapshot artifact guard
    carbonstack-cypher 9ab994c docs: point to local operator runbook
    carbonstack-os     1bbbe52 docs: clarify CarbonStackOS target direction

No dirty files were shown in the final clean status snapshot.

---

## 5. Critical Blunders / Lessons Preserved

### 5.1 Python process cleanup blunder: `go run` wrapper did not terminate cleanly

The first v0.3.25 script attempted to launch:

    go run ./cmd/cypher

and then terminate it during a health-check proof. Python timed out while waiting for the `go run` wrapper/child process to terminate:

    subprocess.TimeoutExpired: Command '['go', 'run', './cmd/cypher']' timed out after 5 seconds

Correct lesson:

    For process-lifecycle proofs, prefer building a temporary binary and launching that binary in its own process group.
    Kill the process group, not just the wrapper process.
    Use longer cleanup timeouts and robust SIGTERM/SIGKILL fallback.

The retry used a temporary `carbonstack-cypher-healthcheck` binary under the temporary health-check data directory, which is the better pattern for future local operator lifecycle proofs.

### 5.2 `git push` before commit is not meaningful

During the v0.3.25 flow, `git push` was run in `carbonstack` before the dirty docs were staged and committed. Git correctly reported:

    Everything up-to-date

Then `git status` still showed:

    modified: docs/README.md
    modified: roadmap/ROADMAP.md
    untracked: docs/144-local-cypher-operator-runbook-skeleton-v0.md

Correct lesson:

    Always check `git status --short` before assuming push changed anything.
    Push only after staging and committing the intended files.

### 5.3 Path typo blunder: `roadmap/ROADMAP.m`

During staging, the operator tried:

    git add roadmap/ROADMAP.m

which failed because the correct file is:

    roadmap/ROADMAP.md

This was corrected before commit.

### 5.4 Generic commit message blunder

The `carbonstack` commit message was:

    docs: update readme, roadmap, and historical doc

This is accurate but less semantically useful than the intended message:

    docs: record local Cypher operator runbook skeleton

Correct lesson:

    Commit messages matter for timeline readability.
    The docs content still preserves the actual rung meaning.
    Do not rewrite history just to improve this commit message unless there is a concrete reason.

### 5.5 Validation occurred before the final docs commit, but the final commit was docs-only

The uploaded snapshot shows `doctor`/`core` validation occurred while `carbonstack` docs were dirty, then the docs were committed afterward. Because the final commit was docs-only and the same content existed during validation, this is acceptable for the v0.3.25 docs checkpoint.

Stricter future habit:

    Run doctor/core after final commits when a breakpoint is about to be cut.
    Then clean OpenMLS artifacts.
    Then capture final clean status.

### 5.6 Artifact cleanup remains required after validation

The final validation produced expected sidecar generated roots, then cleanup was run. Final repo status snapshot showed clean status after cleanup. Continue this pattern after every `core` run.

---

## 6. Current Blockers / Not Validated

Current blockers and nonvalidated areas:

    Local-only operator runbook is a skeleton, not mature deployment documentation.
    Local operator config/data path convention is documented but not enforced by code.
    Cypher default bind remains :8080; local-only runbook recommends explicit 127.0.0.1:8080.
    Health-check proof is minimal; it does not validate full API lifecycle under a persistent operator DB.
    No local-backbone runner profile exists yet.
    Runtime Comms send/inbox remains stub-era and not wired as user-facing OpenMLS messenger UX.
    Production-safe provider storage and secure vault design are unsolved.
    Hostile-server rollback/replay/metadata-abuse validation remains future work.
    Android/CarbonStackOS implementation remains later work.
    CarbonStack Relay Space mechanics remain a defined future concept, not an implementation.
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
    Do not claim secure vault/storage.
    Do not claim external audit or certification.
    Do not claim CarbonStack Relay Space mechanics are implemented.
    Do not claim local operator profile is complete.
    Do not claim v0.3.25 is a public release.

Updated allowed claim:

    CarbonStack v0.3.25 records the first local-only Cypher operator runbook skeleton.
    carbonstack-cypher README points to the CarbonStack docs runbook skeleton.
    The runbook recommends explicit 127.0.0.1 loopback binding, explicit DB path, explicit migrations path, explicit dev invite, health-check command, and blunt reset command.
    WSL Debian doctor/core validation passed during the v0.3.25 flow.

---

## 7. Next Safest Actions

Recommended next sequence after this v0.3.25 breakpoint:

    v0.3.26 preferred:
      local operator config/data path convention refinement and/or exact local API lifecycle proof

Recommended concrete next rung options:

    Option A: docs refinement / convention
      tighten the runbook and local data/config path rationale;
      document why ~/.local/share/carbonstack/cypher/cypher.db is used;
      document source-tree migrations path versus installed migrations path;
      document dev-invite warning more clearly.

    Option B: lifecycle proof
      start Cypher with explicit local operator env;
      call /v0/health;
      create a dev invite or use explicit dev invite;
      claim invite;
      register device;
      submit envelope;
      retrieve envelope;
      ack envelope;
      stop process;
      restart with same DB;
      verify schema_migrations skip behavior still works;
      document the proof in carbonstack/docs.

Recommendation:

    Do Option B soon, but keep it local-only and avoid public ingress/systemd/cloudflared. This would provide much stronger evidence that the local operator skeleton is not just a health endpoint.

Do **not** jump directly to:

    public ingress
    cloudflared
    systemd
    real homelab
    runtime Comms send/inbox OpenMLS UX
    Android app
    CarbonStackOS
    remote admin plane
    CarbonStack Relay Space implementation
    local-backbone runner profile

---

## 8. Critical Paths / Functions

### WSL Debian active working umbrella

    [WSL-DEBIAN] umbrella: ~/repos/carbonstack_umbrella
    [WSL-DEBIAN] carbonstack: ~/repos/carbonstack_umbrella/carbonstack
    [WSL-DEBIAN] carbonstack-comms: ~/repos/carbonstack_umbrella/carbonstack-comms
    [WSL-DEBIAN] carbonstack-cypher: ~/repos/carbonstack_umbrella/carbonstack-cypher
    [WSL-DEBIAN] carbonstack-os: ~/repos/carbonstack_umbrella/carbonstack-os
    [GO-RUNNER] runner dir: ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate

### WSL Debian known-good validation commands

From runner dir:

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go run . --profile doctor
    go run . --profile core

From carbonstack repo root:

    cd ~/repos/carbonstack_umbrella/carbonstack
    go run ./tools/carbonstack-validate --profile doctor
    go run ./tools/carbonstack-validate --profile core

### Artifact cleanup after validation

    cd ~/repos/carbonstack_umbrella/carbonstack-comms
    rm -rf internal/protocol/mls/openmls-sidecar/.carbonstack-openmls-sidecar-state
    rm -rf internal/protocol/mls/openmls-sidecar/target

### v0.3.25 docs and README pointer

    [DOC] v0.3.25 runbook skeleton:
    carbonstack/docs/144-local-cypher-operator-runbook-skeleton-v0.md

    [DOC] docs index:
    carbonstack/docs/README.md

    [DOC] roadmap:
    carbonstack/roadmap/ROADMAP.md

    [CYPHER DOC] README pointer:
    carbonstack-cypher/README.md

### Recommended local-only operator commands from docs

    cd ~/repos/carbonstack_umbrella/carbonstack-cypher
    mkdir -p ~/.local/share/carbonstack/cypher

    CYPHER_ADDR=127.0.0.1:8080 \
    CYPHER_DB=$HOME/.local/share/carbonstack/cypher/cypher.db \
    CYPHER_MIGRATIONS=$HOME/repos/carbonstack_umbrella/carbonstack-cypher/migrations \
    CYPHER_DEV_INVITE=dev-invite \
    go run ./cmd/cypher

Health check:

    curl -i http://127.0.0.1:8080/v0/health

Reset:

    rm -f ~/.local/share/carbonstack/cypher/cypher.db

### Critical functions / behavior preserved from v0.3.24

    Store.Migrate:
      ensures schema_migrations exists;
      reads migration .sql files in sorted order;
      computes sha256 checksum for each migration;
      skips matching already-applied migrations;
      fails on checksum mismatch for already-applied migration;
      applies new migrations in a transaction;
      records migration after successful application.

    TestMigrateRecordsAppliedMigrations:
      proves applied migrations are recorded with non-empty checksum/applied_at.

    TestMigrateCanRunTwiceAgainstSameDB:
      proves repeated migration against the same DB now succeeds.

    TestMigrateDetectsAppliedMigrationChecksumMismatch:
      proves already-applied migration drift hard-fails.

---

## 9. Lean Breakpoint Summary

    CarbonStack v0.3.25 is a stable mainline docs/runbook breakpoint after v0.3.24. It records the first local-only Cypher operator runbook skeleton in carbonstack/docs/144, points carbonstack-cypher README to that runbook, documents explicit loopback startup, explicit SQLite DB path, explicit migrations path, explicit dev invite, health-check command, and blunt reset command, preserves WSL Debian doctor/core validation during the rung, cleans generated OpenMLS artifacts, and leaves mature local operator lifecycle proof, config/data convention refinement, local-backbone runner profile, runtime Comms OpenMLS UX, CarbonStack Relay Space mechanics, public ingress, systemd, cloudflared, homelab, Android, CarbonStackOS, hostile-server proof, secure vault, and production/security claims deferred.

---

## 10. Preserved v0.3.24 Operational Process Log

The following section preserves the previous v0.3.24 LogDoc body as operational continuity. It is intentionally retained rather than compressed away. Where v0.3.24 text conflicts with the v0.3.25 current-state overlay above, v0.3.25 wins for current state; v0.3.24 remains the provenance/process ledger for the Cypher schema migration tracking implementation and the preserved v0.3.23 / v0.3.22 / v0.3.21 / v0.3.20 history.




# CarbonStack LogDoc v0.3.24

**Last updated:** 2026-06-02 local session  
**Owner/operator:** bitcrusher32  
**Assistant/model:** GPT-5.5 Thinking  
**Current phase:** Phase 2E-prep / v0.3.x **Cypher schema_migrations implementation checkpoint**. `carbonstack` is now at `76f0258 docs: record Cypher schema migrations implementation`, after `v0.3.23` proved the Cypher repeated-migration persistence hazard at `1fef7fe`. `carbonstack-cypher` is now at `42f838f db: track applied migrations`, implementing migration tracking after the v0.3.23 proof-first recon. `carbonstack-comms` remains at `012c8bf scripts: support source snapshot artifact guard`; `carbonstack-os` remains at `1bbbe52 docs: clarify CarbonStackOS target direction`. The live WSL Debian working umbrella at `~/repos/carbonstack_umbrella` passed `doctor` and `core` after both the Cypher code commit and CarbonStack docs commit, and generated sidecar artifacts were cleaned afterward.

**Version schema:** `v[scope].[timeline]`. This file is `v0.3.24`, the Cypher schema migration tracking implementation handoff after the `v0.3.23` Cypher migration persistence recon checkpoint. Per LogDoc V2 practice, this Markdown file intentionally preserves verbose operational continuity; the JSON breakpoint remains the lean handoff state. The dense v0.3.23, v0.3.22, v0.3.21, and v0.3.20 operational process logs are preserved later in this file rather than compressed away.

**LogDoc V2 usage note:** LogDoc V2 uses a split between a Markdown LogDoc for working-session continuity, a JSON state file for end-of-session handoff, Git for history preservation, and README/docs for public claims. This v0.3.24 Markdown file therefore intentionally stays “bloated” where operational process, blunders, validation ladders, and historical context matter. The v0.3.24 JSON should remain lean and should agree with this file on the latest current state.

---

## 1. Project Goal

**Active goal:** Preserve CarbonStack state after implementing and validating minimal CarbonStackCypher migration tracking. The immediate next line of work is local-only operator runbook/config-data convention work, not runtime Comms UX, not public ingress, not systemd/cloudflared, not real homelab deployment, and not Android/CarbonStackOS work.

CarbonStack’s broader project intent remains:

- **CarbonStack** is the secure-communications appliance-stack project and shared doctrine/release authority.
- **CarbonStackOS** is the future constrained Android-derived appliance OS.
- **CarbonStackComms** is the text-first encrypted messaging client and current OpenMLS sidecar/relay integration surface.
- **CarbonStackCypher** is the self-hostable relay/storage stack.
- The `carbonstack` repo remains the doctrine/spec/docs/release authority; implementation repos must not contradict it without revising the spec.

**North star:**

> Its goal is not to be a better smartphone. Its goal is to stop being a smartphone.

**Core doctrine:**

> Every feature is guilty until it proves it does not add unacceptable parser, network, sensor, identity, or filesystem authority.

**Current reality after v0.3.24:** `v0.3.20` remains the current public runner-backed testing release. `v0.3.21` defined the local-only deployability model. `v0.3.22` narrowed the local deployability surface around Cypher operator config/API/database assumptions and adopted **CarbonStack Relay Space** terminology. `v0.3.23` proved the concrete SQLite migration persistence hazard. `v0.3.24` implements the first fix for that hazard: CarbonStackCypher now tracks applied migrations and skips already-applied matching migrations, while hard-failing if an already-applied migration file drifts by checksum.

The current public release page remains:

    https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.3.20

The current public testing interpretation remains:

    v0.3.20 is the preferred public runner-backed testing release.
    v0.3.4 remains available as the older Windows 11 / PowerShell legacy testbed.
    Debian / WSL Debian is now the mainline public dev/test validation direction.
    Windows 11 short-path validation was the final explicit Windows dev/test validation for this phase.

This is still **not** a deployability release. It does **not** prove production readiness, production E2EE, hostile-server safety, metadata privacy, Android readiness, secure vault/storage, Debian deployability, systemd readiness, cloudflared readiness, public ingress safety, real homelab validation, external audit, or certification.

---

## 2. Current Repo Heads

    carbonstack        76f0258 docs: record Cypher schema migrations implementation
    carbonstack-comms  012c8bf scripts: support source snapshot artifact guard
    carbonstack-cypher 42f838f db: track applied migrations
    carbonstack-os     1bbbe52 docs: clarify CarbonStackOS target direction

Current `carbonstack` public runner-backed testing release tag:

    v0.3.20 -> 3eeb1a1e1a / 3eeb1a1 docs: clean public release surface

Previous mainline deployability checkpoints:

    v0.3.23 mainline -> 1fef7fe docs: record Cypher migration persistence recon
    v0.3.22 mainline -> 328ffc4 docs: record Cypher local operator surface recon
    v0.3.21 mainline -> 6850c58 docs: record local backbone deployability recon

Legacy `carbonstack` Windows/PowerShell testbed release tag:

    v0.3.4 -> 7c5b7e6 docs: record release snapshot self-test fix

Historical `carbonstack` release tag:

    v0.3.0 -> 59245fc docs: freeze v0.3.0 release packaging plan

Important continuity note: `v0.3.24` does not supersede `v0.3.20` as a public release. It advances mainline implementation after the v0.3.20 public runner-backed testing release, the v0.3.21 local-only deployability model checkpoint, the v0.3.22 Cypher local operator surface recon, and the v0.3.23 proof-first migration persistence recon. It is safe to pause here because the code/docs were pushed, WSL Debian `doctor`/`core` validated afterward, and generated artifacts were cleaned.

---

## 3. Current Validated State

Validated at the v0.3.24 checkpoint:

    [WSL-DEBIAN] working umbrella exists at ~/repos/carbonstack_umbrella
    [WSL-DEBIAN] carbonstack is at 76f0258
    [WSL-DEBIAN] carbonstack-comms is at 012c8bf
    [WSL-DEBIAN] carbonstack-cypher is at 42f838f
    [WSL-DEBIAN] carbonstack-os is at 1bbbe52
    [CYPHER] schema_migrations implementation committed and pushed
    [CYPHER] go test ./internal/db -count=1 -v passed during docs-generation/validation flow
    [CYPHER] go test ./... -count=1 passed during docs-generation/validation flow
    [GO-RUNNER] doctor passed after v0.3.24 commits
    [GO-RUNNER] core passed after v0.3.24 commits
    [ARTIFACT CLEANUP] expected OpenMLS generated roots were removed after final validation
    [FINAL STATUS] all four repos were clean after cleanup

Observed WSL Debian toolchain baseline:

    git version 2.47.3
    go version go1.24.4 linux/amd64
    rustc 1.96.0 (ac68faa20 2026-05-25)
    cargo 1.96.0 (30a34c682 2026-05-25)
    sqlite3 3.46.1

Observed `core` validation result:

    pre-test artifact scan: PASS / no generated/private/build artifact hits
    OpenMLS real-Cypher lifecycle: PASS
    carbonstack-comms package tests: PASS
    carbonstack-cypher package tests: PASS
    post-test artifact scan: expected OpenMLS sidecar generated roots only
    VALIDATION PASSED

Expected post-test generated roots:

    carbonstack-comms/internal/protocol/mls/openmls-sidecar/.carbonstack-openmls-sidecar-state
    carbonstack-comms/internal/protocol/mls/openmls-sidecar/target

These roots are expected only after tests and must not be committed or treated as source artifacts. They were cleaned before the final clean status snapshot.

---

## 4. v0.3.24 Work Completed

### 4.1 CarbonStackCypher migration tracking implemented

Code commit:

    carbonstack-cypher 42f838f db: track applied migrations

Primary implementation file:

    carbonstack-cypher/internal/db/db.go

New permanent test file:

    carbonstack-cypher/internal/db/migrate_test.go

Implementation behavior:

    Store.Migrate now ensures a schema_migrations table exists.
    Migration files are still read in sorted filename order.
    Each .sql migration file gets a SHA-256 checksum.
    If migration_name + same sha256 already exists, the migration is skipped.
    If migration_name exists with a different sha256, migration fails with checksum mismatch.
    New migrations are applied inside a transaction.
    Migration state is recorded only after successful application.
    Applied migration records include migration_name, sha256, and applied_at.

schema_migrations shape:

    migration_name TEXT PRIMARY KEY
    sha256 TEXT NOT NULL
    applied_at TEXT NOT NULL

This directly addresses the v0.3.23 proof result: migration 002 is no longer blindly reapplied against an already-upgraded DB if it was previously recorded with the same checksum.

### 4.2 Permanent migration tests added

New test file:

    carbonstack-cypher/internal/db/migrate_test.go

Tests added:

    TestMigrateRecordsAppliedMigrations
    TestMigrateCanRunTwiceAgainstSameDB
    TestMigrateDetectsAppliedMigrationChecksumMismatch

Validation intent:

    Fresh DB migration records applied migrations.
    Running Migrate twice against the same DB now succeeds.
    Already-applied migration file drift hard-fails by checksum mismatch.

This turns the v0.3.23 temporary proof/recon test into permanent desired-behavior tests after the implementation.

### 4.3 CarbonStack docs/result record landed

Docs commit:

    carbonstack 76f0258 docs: record Cypher schema migrations implementation

New/updated docs:

    carbonstack/docs/143-cypher-schema-migrations-implementation-v0.md
    carbonstack/docs/README.md
    carbonstack/roadmap/ROADMAP.md

The docs result record captures:

    why v0.3.24 exists;
    what v0.3.23 proved;
    schema_migrations implementation behavior;
    tests added;
    validation results;
    doctrine/security interpretation;
    remaining limits;
    next recommended rung.

### 4.4 Migration-file immutability policy clarified

v0.3.24 strengthens the migration policy:

    Applied migration files should be treated as immutable.
    If a migration needs correction, add a new migration.
    Do not silently edit an already-applied migration.
    If an already-applied migration file changes, checksum mismatch should hard-fail.

This fits CarbonStack’s evidence-led and attack-surface-minimizing style. It favors loud failure over silent schema drift.

### 4.5 Final validation and cleanup completed

Final WSL Debian repo snapshot:

    carbonstack        76f0258 docs: record Cypher schema migrations implementation
    carbonstack-comms  012c8bf scripts: support source snapshot artifact guard
    carbonstack-cypher 42f838f db: track applied migrations
    carbonstack-os     1bbbe52 docs: clarify CarbonStackOS target direction

Final validation:

    go run . --profile doctor
    go run . --profile core

Final validation passed under WSL Debian. `core` passed OpenMLS real-Cypher lifecycle, Comms package tests, and Cypher package tests, including the new `internal/db` package tests. After validation, expected OpenMLS generated roots were removed from `carbonstack-comms`, and final clean status showed no dirty files across the four repos.

---

## 5. Critical Blunders / Lessons Preserved

### 5.1 v0.3.23 proof-first pattern paid off

v0.3.23 created a temporary proof test, documented the repeated-migration failure, and removed the test. v0.3.24 then implemented the desired behavior and added permanent regression tests.

Correct pattern preserved:

    Proof/recon test:
      OK to create, run, document, and remove.

    Permanent regression test:
      Add after implementation changes desired behavior.

This pattern should be reused when uncertainty is high.

### 5.2 Migration files are now part of an evidence boundary

Because migration checksums are recorded, historical migration files become operationally significant. Editing an already-applied migration is no longer a harmless docs/code tweak; it can intentionally break startup/migration with a checksum mismatch.

Correct policy:

    Treat applied migrations as immutable.
    Fix forward with a new migration.
    Prefer loud failure over silent schema drift.

### 5.3 schema_migrations is deployability hardening, not deployability completion

v0.3.24 removes the known repeated-migration blocker, but it does not implement the local operator profile. It also does not validate public service behavior, systemd, cloudflared, a real homelab, or a multi-relay-space operator model.

Correct current claim:

    Persistent local Cypher migration behavior is now materially safer for already-applied migrations.
    Full local operator deployability remains future work.

### 5.4 Keep the rung narrow

v0.3.24 deliberately did not change:

    CYPHER_ADDR default
    local operator data path defaults
    Comms runtime send/inbox UX
    runner profiles
    public release surface
    systemd/cloudflared/homelab paths

That was the right call. The migration fix is enough for one validated rung.

### 5.5 Artifact cleanup remains required after validation

The final validation produced expected sidecar generated roots, then cleanup was run. Final repo status snapshot showed clean status after cleanup. Continue this pattern after every `core` run.

---

## 6. Current Blockers / Not Validated

Current blockers and nonvalidated areas:

    Local-only operator profile is still not implemented.
    Local operator runbook is not implemented beyond recon/planning.
    Cypher default bind is still broader than the local-only docs recommendation if left as :8080.
    Local config/data path convention is not implemented.
    No local-backbone runner profile exists yet.
    Runtime Comms send/inbox remains stub-era and not wired as user-facing OpenMLS messenger UX.
    Production-safe provider storage and secure vault design are unsolved.
    Hostile-server rollback/replay/metadata-abuse validation remains future work.
    Android/CarbonStackOS implementation remains later work.
    CarbonStack Relay Space mechanics remain a defined future concept, not an implementation.
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
    Do not claim secure vault/storage.
    Do not claim external audit or certification.
    Do not claim CarbonStack Relay Space mechanics are implemented.
    Do not claim local operator profile is complete.

Updated allowed claim:

    CarbonStackCypher now tracks applied migrations by filename, SHA-256 checksum, and applied_at timestamp.
    Re-running migrations against the same DB now skips matching already-applied migrations.
    Editing an already-applied migration file causes checksum mismatch failure.
    This removes the v0.3.23 repeated-migration blocker for experimental local operator persistence, but does not make CarbonStack production-deployable.

---

## 7. Next Safest Actions

Recommended next sequence after this v0.3.24 breakpoint:

    v0.3.25 preferred:
      local operator runbook skeleton and config/data path convention planning

Recommended concrete next rung:

    1. Draft local-only Cypher operator runbook skeleton in carbonstack/docs.
    2. Document 127.0.0.1-only startup.
    3. Document explicit CYPHER_DB path.
    4. Document explicit CYPHER_MIGRATIONS path.
    5. Document explicit CYPHER_DEV_INVITE behavior and warning.
    6. Document reset/cleanup semantics.
    7. Document what survives restart.
    8. Document that persistence is still experimental.
    9. Avoid systemd/cloudflared/public ingress.
    10. Decide whether a later code rung should change CYPHER_ADDR default to 127.0.0.1:8080 or keep defaults and require explicit local-operator env.

Do **not** jump directly to:

    public ingress
    cloudflared
    systemd
    real homelab
    runtime Comms send/inbox OpenMLS UX
    Android app
    CarbonStackOS
    remote admin plane
    CarbonStack Relay Space implementation
    local-backbone runner profile

---

## 8. Critical Paths / Functions

### WSL Debian active working umbrella

    [WSL-DEBIAN] umbrella: ~/repos/carbonstack_umbrella
    [WSL-DEBIAN] carbonstack: ~/repos/carbonstack_umbrella/carbonstack
    [WSL-DEBIAN] carbonstack-comms: ~/repos/carbonstack_umbrella/carbonstack-comms
    [WSL-DEBIAN] carbonstack-cypher: ~/repos/carbonstack_umbrella/carbonstack-cypher
    [WSL-DEBIAN] carbonstack-os: ~/repos/carbonstack_umbrella/carbonstack-os
    [GO-RUNNER] runner dir: ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate

### WSL Debian known-good validation commands

From runner dir:

    cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
    go run . --profile doctor
    go run . --profile core

From carbonstack repo root:

    cd ~/repos/carbonstack_umbrella/carbonstack
    go run ./tools/carbonstack-validate --profile doctor
    go run ./tools/carbonstack-validate --profile core

### Artifact cleanup after validation

    cd ~/repos/carbonstack_umbrella/carbonstack-comms
    rm -rf internal/protocol/mls/openmls-sidecar/.carbonstack-openmls-sidecar-state
    rm -rf internal/protocol/mls/openmls-sidecar/target

### v0.3.24 docs

    [DOC] v0.3.24 implementation record:
    carbonstack/docs/143-cypher-schema-migrations-implementation-v0.md

    [DOC] docs index:
    carbonstack/docs/README.md

    [DOC] roadmap:
    carbonstack/roadmap/ROADMAP.md

### Cypher migration implementation surfaces

    [CODE] cypher DB:
    carbonstack-cypher/internal/db/db.go

    [TEST] cypher migration tests:
    carbonstack-cypher/internal/db/migrate_test.go

    [SQL] migration 001:
    carbonstack-cypher/migrations/001_init.sql

    [SQL] migration 002:
    carbonstack-cypher/migrations/002_envelope_payload_metadata.sql

### Critical functions / behavior

    Store.Migrate:
      ensures schema_migrations exists;
      reads migration .sql files in sorted order;
      computes sha256 checksum for each migration;
      skips matching already-applied migrations;
      fails on checksum mismatch for already-applied migration;
      applies new migrations in a transaction;
      records migration after successful application.

    TestMigrateRecordsAppliedMigrations:
      proves applied migrations are recorded with non-empty checksum/applied_at.

    TestMigrateCanRunTwiceAgainstSameDB:
      proves repeated migration against the same DB now succeeds.

    TestMigrateDetectsAppliedMigrationChecksumMismatch:
      proves already-applied migration drift hard-fails.

---

## 9. Lean Breakpoint Summary

    CarbonStack v0.3.24 is a stable mainline implementation breakpoint after v0.3.23. It implements CarbonStackCypher schema_migrations tracking in carbonstack-cypher, records migration filename, SHA-256 checksum, and applied_at timestamp, skips matching already-applied migrations, hard-fails on checksum drift, adds permanent migration tests, records the result in carbonstack docs, validates WSL Debian doctor/core after both commits, cleans generated OpenMLS artifacts, and leaves local operator runbook/config-data convention, local-backbone runner profile, runtime Comms OpenMLS UX, CarbonStack Relay Space mechanics, public ingress, systemd, cloudflared, homelab, Android, CarbonStackOS, hostile-server proof, secure vault, and production/security claims deferred.

---

## 10. Preserved v0.3.23 Operational Process Log

The following section preserves the previous v0.3.23 LogDoc body as operational continuity. It is intentionally retained rather than compressed away. Where v0.3.23 text conflicts with the v0.3.24 current-state overlay above, v0.3.24 wins for current state; v0.3.23 remains the provenance/process ledger for the Cypher migration persistence recon and the preserved v0.3.22 / v0.3.21 / v0.3.20 history.




# CarbonStack LogDoc v0.3.23

**Last updated:** 2026-06-02 local session  
**Owner/operator:** bitcrusher32  
**Assistant/model:** GPT-5.5 Thinking  
**Current phase:** Phase 2E-prep / v0.3.x **Cypher migration persistence recon checkpoint**. `carbonstack` is now at `1fef7fe docs: record Cypher migration persistence recon`, after `v0.3.22` landed Cypher local operator surface recon at `328ffc4`, `v0.3.21` landed local-only backbone deployability recon at `6850c58`, and `v0.3.20` remains the current public runner-backed testing release at `3eeb1a1e1a / 3eeb1a1 docs: clean public release surface`. `carbonstack-comms` remains at `012c8bf scripts: support source snapshot artifact guard`; `carbonstack-cypher` remains at `6f92c34 chore: fix readme formatting`; `carbonstack-os` remains at `1bbbe52 docs: clarify CarbonStackOS target direction`. The live WSL Debian working umbrella at `~/repos/carbonstack_umbrella` passed `doctor` and `core` after the v0.3.23 docs/recon commit, and generated sidecar artifacts were cleaned afterward.

**Version schema:** `v[scope].[timeline]`. This file is `v0.3.23`, the Cypher migration persistence recon handoff after the `v0.3.22` Cypher local operator surface recon checkpoint. Per LogDoc V2 practice, this Markdown file intentionally preserves verbose operational continuity; the JSON breakpoint remains the lean handoff state. The dense v0.3.22, v0.3.21, and v0.3.20 operational process logs are preserved later in this file rather than compressed away.

**LogDoc V2 usage note:** LogDoc V2 uses a split between a Markdown LogDoc for working-session continuity, a JSON state file for end-of-session handoff, Git for history preservation, and README/docs for public claims. This v0.3.23 Markdown file therefore intentionally stays “bloated” where operational process, blunders, validation ladders, and historical context matter. The v0.3.23 JSON should remain lean and should agree with this file on the latest current state.

---

## 1. Project Goal

**Active goal:** Preserve CarbonStack state after completing the v0.3.23 WSL Debian Cypher migration persistence recon checkpoint. The immediate next line of work is to implement or design the migration-hardening fix that this recon proved necessary: `schema_migrations` tracking or an equivalent idempotent migration mechanism for CarbonStackCypher before stronger persistent local operator DB claims.

CarbonStack’s broader project intent remains:

- **CarbonStack** is the secure-communications appliance-stack project and shared doctrine/release authority.
- **CarbonStackOS** is the future constrained Android-derived appliance OS.
- **CarbonStackComms** is the text-first encrypted messaging client and current OpenMLS sidecar/relay integration surface.
- **CarbonStackCypher** is the self-hostable relay/storage stack.
- The `carbonstack` repo remains the doctrine/spec/docs/release authority; implementation repos must not contradict it without revising the spec.

**North star:**

> Its goal is not to be a better smartphone. Its goal is to stop being a smartphone.

**Core doctrine:**

> Every feature is guilty until it proves it does not add unacceptable parser, network, sensor, identity, or filesystem authority.

**Current reality after v0.3.23:** `v0.3.20` remains the current public runner-backed testing release. `v0.3.21` defined the local-only deployability model. `v0.3.22` narrowed the local deployability surface around Cypher operator config/API/database assumptions and adopted **CarbonStack Relay Space** terminology. `v0.3.23` proves the concrete SQLite migration persistence hazard: fresh DB migration remains known-good, but the current migration path cannot be claimed safe for persistent local operator DB restart/upgrade behavior until migration tracking or idempotency is implemented.

The current public release page remains:

```text
https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.3.20
```

The current public testing interpretation remains:

```text
v0.3.20 is the preferred public runner-backed testing release.
v0.3.4 remains available as the older Windows 11 / PowerShell legacy testbed.
Debian / WSL Debian is now the mainline public dev/test validation direction.
Windows 11 short-path validation was the final explicit Windows dev/test validation for this phase.
```

This is still **not** a deployability release. It does **not** prove production readiness, production E2EE, hostile-server safety, metadata privacy, Android readiness, secure vault/storage, Debian deployability, systemd readiness, cloudflared readiness, public ingress safety, real homelab validation, external audit, or certification.

---

## 2. Current Repo Heads

```text
carbonstack        1fef7fe docs: record Cypher migration persistence recon
carbonstack-comms  012c8bf scripts: support source snapshot artifact guard
carbonstack-cypher 6f92c34 chore: fix readme formatting
carbonstack-os     1bbbe52 docs: clarify CarbonStackOS target direction
```

Current `carbonstack` public runner-backed testing release tag:

```text
v0.3.20 -> 3eeb1a1e1a / 3eeb1a1 docs: clean public release surface
```

Previous mainline deployability recon checkpoints:

```text
v0.3.22 mainline -> 328ffc4 docs: record Cypher local operator surface recon
v0.3.21 mainline -> 6850c58 docs: record local backbone deployability recon
```

Legacy `carbonstack` Windows/PowerShell testbed release tag:

```text
v0.3.4 -> 7c5b7e6 docs: record release snapshot self-test fix
```

Historical `carbonstack` release tag:

```text
v0.3.0 -> 59245fc docs: freeze v0.3.0 release packaging plan
```

Important continuity note: `v0.3.23` does not supersede `v0.3.20` as a public release. It advances mainline documentation/recon after the v0.3.20 public runner-backed testing release, the v0.3.21 local-only deployability model checkpoint, and the v0.3.22 Cypher local operator surface recon. It is safe to pause here because the recon doc was pushed, the WSL Debian working umbrella validated afterward, and generated artifacts were cleaned.

---

## 3. Current Validated State

Validated at the v0.3.23 checkpoint:

```text
[WSL-DEBIAN] working umbrella exists at ~/repos/carbonstack_umbrella
[WSL-DEBIAN] carbonstack is at 1fef7fe
[WSL-DEBIAN] carbonstack-comms is at 012c8bf
[WSL-DEBIAN] carbonstack-cypher is at 6f92c34
[WSL-DEBIAN] carbonstack-os is at 1bbbe52
[GO-RUNNER] doctor passed after v0.3.23 docs/recon commit
[GO-RUNNER] core passed after v0.3.23 docs/recon commit
[ARTIFACT CLEANUP] expected OpenMLS generated roots were removed after final validation
[FINAL STATUS] all four repos were clean after cleanup
```

Observed WSL Debian toolchain baseline:

```text
git version 2.47.3
go version go1.24.4 linux/amd64
rustc 1.96.0 (ac68faa20 2026-05-25)
cargo 1.96.0 (30a34c682 2026-05-25)
sqlite3 3.46.1
```

Observed `core` validation result:

```text
pre-test artifact scan: PASS / no generated/private/build artifact hits
OpenMLS real-Cypher lifecycle: PASS
carbonstack-comms package tests: PASS
carbonstack-cypher package tests: PASS
post-test artifact scan: expected OpenMLS sidecar generated roots only
VALIDATION PASSED
```

Expected post-test generated roots:

```text
carbonstack-comms/internal/protocol/mls/openmls-sidecar/.carbonstack-openmls-sidecar-state
carbonstack-comms/internal/protocol/mls/openmls-sidecar/target
```

These roots are expected only after tests and must not be committed or treated as source artifacts. They were cleaned before the final clean status snapshot.

---

## 4. v0.3.23 Work Completed

### 4.1 Cypher migration persistence recon doc landed

New/updated docs at v0.3.23:

```text
carbonstack/docs/142-cypher-migration-persistence-recon-v0.md
carbonstack/docs/README.md
carbonstack/roadmap/ROADMAP.md
```

Commit:

```text
1fef7fe docs: record Cypher migration persistence recon
```

The recon doc records a proof-first test of the current CarbonStackCypher migration path against a reused persistent SQLite database. It intentionally documents the behavior before implementation so the next code rung is evidence-led rather than vibes-led.

### 4.2 Temporary recon test proved the current migration hazard

A temporary Go test file was created under:

```text
carbonstack-cypher/internal/db/migrate_twice_recon_test.go
```

The temporary test:

```text
1. created a temporary SQLite DB;
2. opened it through the current internal/db package;
3. ran Store.Migrate("../../migrations") once;
4. ran Store.Migrate("../../migrations") a second time against the same DB;
5. expected the second pass to fail under current behavior;
6. deleted the temporary test file after the proof run.
```

The temporary test was not committed and `carbonstack-cypher` remains unchanged at:

```text
6f92c34 chore: fix readme formatting
```

Result:

```text
PASS / current hazard reproduced
```

Interpretation:

```text
The first migration pass succeeds against a fresh DB.
The second migration pass against the same already-migrated DB fails under current behavior.
The failure matches the expected repeated-ALTER migration hazard.
```

### 4.3 Current migration behavior is now proven, not guessed

Current migration files remain:

```text
carbonstack-cypher/migrations/001_init.sql
carbonstack-cypher/migrations/002_envelope_payload_metadata.sql
```

Current migration characteristics:

```text
001_init.sql uses CREATE TABLE IF NOT EXISTS.
002_envelope_payload_metadata.sql uses raw ALTER TABLE envelopes ADD COLUMN payload_sha256 TEXT;
002_envelope_payload_metadata.sql uses raw ALTER TABLE envelopes ADD COLUMN payload_size_bytes INTEGER;
No schema_migrations tracking exists in the current code.
Store.Migrate reads migration files and executes them in sorted order.
```

The important result:

```text
Fresh DB behavior is known-good.
Persistent local DB behavior is experimental.
Repeated migration against an already-upgraded DB is not safe under current behavior.
Schema reset may be required between versions until migration tracking/idempotency is implemented.
```

### 4.4 v0.3.23 changed the next implementation priority

Before v0.3.23, migration behavior was the highest-priority recon concern from v0.3.22. After v0.3.23, it is no longer merely a concern; it is a validated blocker for stronger local operator persistence claims.

Next safest implementation rung:

```text
v0.3.24:
  implement minimal schema_migrations tracking or equivalent migration idempotency in carbonstack-cypher
  add tests for fresh DB migration and repeated migration against the same DB
```

### 4.5 Final validation and cleanup completed

Final WSL Debian repo snapshot:

```text
carbonstack        1fef7fe docs: record Cypher migration persistence recon
carbonstack-comms  012c8bf scripts: support source snapshot artifact guard
carbonstack-cypher 6f92c34 chore: fix readme formatting
carbonstack-os     1bbbe52 docs: clarify CarbonStackOS target direction
```

Final validation:

```text
go run . --profile doctor
go run . --profile core
```

Final validation passed under WSL Debian. After validation, expected OpenMLS generated roots were removed from `carbonstack-comms`, and final clean status showed no dirty files across the four repos.

---

## 5. Critical Blunders / Lessons Preserved

### 5.1 Do not commit proof-only failure/recon tests unless they become permanent regression tests

The v0.3.23 temporary test was intended to prove the current behavior, not to permanently add a test that encodes a bad state. It was created, run, captured in the docs, and removed.

Future equivalent pattern:

```text
Temporary proof/recon test:
  OK if removed after result is documented.

Permanent regression test:
  OK after implementation changes the desired behavior.
```

For v0.3.24, the permanent test should likely invert this: repeated migration against the same DB should pass after `schema_migrations` or idempotent migration handling is implemented.

### 5.2 The DB hazard is restart/upgrade deployability, not normal fresh-test behavior

Existing tests and runner validation remain green because they use fresh/test-owned DB paths. That does not contradict the v0.3.23 finding. A local operator profile has a different requirement: the database may persist across process restarts and future schema changes.

Correct interpretation:

```text
Fresh DB tests passing does not prove persistent DB restart/upgrade safety.
```

### 5.3 Documentation-first was useful here

The proof-first recon prevented premature implementation by showing the exact class of failure. This now gives a clear implementation target:

```text
Track which migrations have been applied.
Do not reapply migration 002.
Add tests that prove repeated migrate works.
```

### 5.4 SQLite remains the right current substrate

This finding does not mean SQLite should be replaced. It means migration behavior needs to be tracked/hardened. PostgreSQL or another OSS DB can remain a future option after the data/operator model stabilizes.

### 5.5 Artifact cleanup remains required after validation

The final validation produced expected sidecar generated roots, then cleanup was run. Final repo status snapshot showed clean status after cleanup. Continue this pattern after every `core` run.

---

## 6. Current Blockers / Not Validated

Current blockers and nonvalidated areas:

```text
Local-only operator profile is not implemented.
Persistent local Cypher DB restart/upgrade behavior is known unsafe under current repeated-migration behavior.
Cypher migration tracking/idempotency is not implemented yet.
Local operator runbook is not implemented beyond recon/planning.
No local-backbone runner profile exists yet.
Runtime Comms send/inbox remains stub-era and not wired as user-facing OpenMLS messenger UX.
Production-safe provider storage and secure vault design are unsolved.
Hostile-server rollback/replay/metadata-abuse validation remains future work.
Android/CarbonStackOS implementation remains later work.
CarbonStack Relay Space mechanics remain a defined future concept, not an implementation.
Public ingress, systemd, cloudflared, real homelab deployment, and production service paths remain out of scope.
```

Hard nonclaims remain:

```text
Do not claim production readiness.
Do not claim production E2EE.
Do not claim hostile-server safety.
Do not claim metadata privacy.
Do not claim Debian deployability.
Do not claim systemd readiness.
Do not claim cloudflared readiness.
Do not claim Android readiness.
Do not claim secure vault/storage.
Do not claim external audit or certification.
Do not claim SQLite persistence/migrations are stable yet.
```

---

## 7. Next Safest Actions

Recommended next implementation sequence after this v0.3.23 breakpoint:

```text
v0.3.24 preferred:
  Cypher schema_migrations implementation in carbonstack-cypher.
```

Recommended concrete next rung:

```text
1. Add a schema_migrations table.
2. Make Store.Migrate apply each migration file exactly once.
3. Record migration filename/version and applied_at timestamp.
4. Preserve sorted migration application order.
5. Add tests for:
   - fresh DB migration;
   - repeated migration on the same DB;
   - migration table records applied versions;
   - existing envelope payload metadata columns remain available.
6. Update carbonstack-cypher docs if needed.
7. Add carbonstack/docs result record after implementation and validation.
8. Run full WSL Debian doctor/core.
9. Clean generated OpenMLS artifacts after validation.
```

Alternative, if implementation is deferred:

```text
Document wipe-only experimental DB behavior explicitly and defer schema_migrations.
```

Recommendation:

```text
Implement schema_migrations before claiming persistent local operator DB behavior.
```

Do **not** jump directly to:

```text
public ingress
cloudflared
systemd
real homelab
runtime Comms send/inbox OpenMLS UX
Android app
CarbonStackOS
remote admin plane
CarbonStack Relay Space implementation
local-backbone runner profile
```

---

## 8. Critical Paths / Commands

### WSL Debian active working umbrella

```text
[WSL-DEBIAN] umbrella: ~/repos/carbonstack_umbrella
[WSL-DEBIAN] carbonstack: ~/repos/carbonstack_umbrella/carbonstack
[WSL-DEBIAN] carbonstack-comms: ~/repos/carbonstack_umbrella/carbonstack-comms
[WSL-DEBIAN] carbonstack-cypher: ~/repos/carbonstack_umbrella/carbonstack-cypher
[WSL-DEBIAN] carbonstack-os: ~/repos/carbonstack_umbrella/carbonstack-os
[GO-RUNNER] runner dir: ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
```

### WSL Debian known-good validation commands

From runner dir:

```text
cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
go run . --profile doctor
go run . --profile core
```

From carbonstack repo root:

```text
cd ~/repos/carbonstack_umbrella/carbonstack
go run ./tools/carbonstack-validate --profile doctor
go run ./tools/carbonstack-validate --profile core
```

### Artifact cleanup after validation

```text
cd ~/repos/carbonstack_umbrella/carbonstack-comms
rm -rf internal/protocol/mls/openmls-sidecar/.carbonstack-openmls-sidecar-state
rm -rf internal/protocol/mls/openmls-sidecar/target
```

### v0.3.23 docs

```text
[DOC] v0.3.23 recon:
carbonstack/docs/142-cypher-migration-persistence-recon-v0.md

[DOC] docs index:
carbonstack/docs/README.md

[DOC] roadmap:
carbonstack/roadmap/ROADMAP.md
```

### Cypher migration implementation surfaces for v0.3.24

```text
[CODE] cypher DB:
carbonstack-cypher/internal/db/db.go

[SQL] migration 001:
carbonstack-cypher/migrations/001_init.sql

[SQL] migration 002:
carbonstack-cypher/migrations/002_envelope_payload_metadata.sql

[TEST] likely new DB migration tests:
carbonstack-cypher/internal/db/*_test.go
```

---

## 9. Lean Breakpoint Summary

```text
CarbonStack v0.3.23 is a stable mainline docs/recon breakpoint after v0.3.22. It records a WSL Debian proof-first Cypher migration persistence recon: a temporary Go test showed the current migration path succeeds on a fresh SQLite DB but fails when migrations are run a second time against the same already-migrated DB. The temporary test was removed, carbonstack-cypher remained unchanged, carbonstack docs recorded the finding, doctor/core validation passed afterward, generated OpenMLS artifacts were cleaned, and the next safest implementation rung is schema_migrations/idempotent migration handling in carbonstack-cypher before stronger persistent local operator DB claims.
```

---

## 10. Preserved v0.3.22 Operational Process Log

The following section preserves the previous v0.3.22 LogDoc body as operational continuity. It is intentionally retained rather than compressed away. Where v0.3.22 text conflicts with the v0.3.23 current-state overlay above, v0.3.23 wins for current state; v0.3.22 remains the provenance/process ledger for the Cypher local operator surface recon and the preserved v0.3.21 / v0.3.20 history.



# CarbonStack LogDoc v0.3.22

**Last updated:** 2026-06-02 local session  
**Owner/operator:** bitcrusher32  
**Assistant/model:** GPT-5.5 Thinking  
**Current phase:** Phase 2E-prep / v0.3.x **Cypher local operator surface recon checkpoint**. `carbonstack` is now at `328ffc4 docs: record Cypher local operator surface recon`, after `v0.3.21` landed local-only backbone deployability recon at `6850c58` and `v0.3.20` was cut as the current public runner-backed testing release at `3eeb1a1e1a / 3eeb1a1 docs: clean public release surface`. `carbonstack-comms` remains at `012c8bf scripts: support source snapshot artifact guard`; `carbonstack-cypher` remains at `6f92c34 chore: fix readme formatting`; `carbonstack-os` remains at `1bbbe52 docs: clarify CarbonStackOS target direction`. The live WSL Debian working umbrella at `~/repos/carbonstack_umbrella` passed `doctor` and `core` after the v0.3.22 docs/recon commit, and generated sidecar artifacts were cleaned afterward.

**Version schema:** `v[scope].[timeline]`. This file is `v0.3.22`, the Cypher local operator surface recon handoff after the `v0.3.21` local-only backbone deployability recon checkpoint. Per LogDoc V2 practice, this Markdown file intentionally preserves verbose operational continuity; the JSON breakpoint remains the lean handoff state. The dense v0.3.21 and v0.3.20 operational process logs are preserved later in this file rather than compressed away.

**LogDoc V2 usage note:** LogDoc V2 uses a split between a Markdown LogDoc for working-session continuity, a JSON state file for end-of-session handoff, Git for history preservation, and README/docs for public claims. This v0.3.22 Markdown file therefore intentionally stays “bloated” where operational process, blunders, validation ladders, and historical context matter. The v0.3.22 JSON should remain lean and should agree with this file on the latest current state.

---

## 1. Project Goal

**Active goal:** Preserve CarbonStack state after completing the v0.3.22 WSL Debian Cypher local operator surface recon checkpoint. The immediate next line of work is still local-only deployability maturation, not runtime Comms UX, not public ingress, not systemd/cloudflared, not real homelab deployment, and not Android/CarbonStackOS work.

CarbonStack’s broader project intent remains:

- **CarbonStack** is the secure-communications appliance-stack project and shared doctrine/release authority.
- **CarbonStackOS** is the future constrained Android-derived appliance OS.
- **CarbonStackComms** is the text-first encrypted messaging client and current OpenMLS sidecar/relay integration surface.
- **CarbonStackCypher** is the self-hostable relay/storage stack.
- The `carbonstack` repo remains the doctrine/spec/docs/release authority; implementation repos must not contradict it without revising the spec.

**North star:**

> Its goal is not to be a better smartphone. Its goal is to stop being a smartphone.

**Core doctrine:**

> Every feature is guilty until it proves it does not add unacceptable parser, network, sensor, identity, or filesystem authority.

**Current reality after v0.3.22:** `v0.3.20` remains the current public runner-backed testing release. `v0.3.21` defined the local-only deployability model. `v0.3.22` is a mainline docs/recon checkpoint that narrows the next technical risk to Cypher local operator surfaces: config defaults, bind address, SQLite database path, migration behavior, dev invite bootstrap behavior, HTTP API assumptions, Comms-to-Cypher seam, OpenMLS sidecar generated-state boundaries, and future runner/operator-helper direction.

The current public release page remains:

```text
https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.3.20
```

The current public testing interpretation remains:

```text
v0.3.20 is the preferred public runner-backed testing release.
v0.3.4 remains available as the older Windows 11 / PowerShell legacy testbed.
Debian / WSL Debian is now the mainline public dev/test validation direction.
Windows 11 short-path validation was the final explicit Windows dev/test validation for this phase.
```

This is still **not** a deployability release. It does **not** prove production readiness, production E2EE, hostile-server safety, metadata privacy, Android readiness, secure vault/storage, Debian deployability, systemd readiness, cloudflared readiness, public ingress safety, real homelab validation, external audit, or certification.

---

## 2. Current Repo Heads

```text
carbonstack        328ffc4 docs: record Cypher local operator surface recon
carbonstack-comms  012c8bf scripts: support source snapshot artifact guard
carbonstack-cypher 6f92c34 chore: fix readme formatting
carbonstack-os     1bbbe52 docs: clarify CarbonStackOS target direction
```

Current `carbonstack` public runner-backed testing release tag:

```text
v0.3.20 -> 3eeb1a1e1a / 3eeb1a1 docs: clean public release surface
```

Previous mainline deployability recon checkpoint:

```text
v0.3.21 mainline -> 6850c58 docs: record local backbone deployability recon
```

Legacy `carbonstack` Windows/PowerShell testbed release tag:

```text
v0.3.4 -> 7c5b7e6 docs: record release snapshot self-test fix
```

Historical `carbonstack` release tag:

```text
v0.3.0 -> 59245fc docs: freeze v0.3.0 release packaging plan
```

Important continuity note: `v0.3.22` does not supersede `v0.3.20` as a public release. It advances mainline documentation/recon after the v0.3.20 public runner-backed testing release and after the v0.3.21 local-only deployability model checkpoint. It is safe to pause here because the recon doc was pushed and the WSL Debian working umbrella validated afterward.

---

## 3. Current Validated State

Validated at the v0.3.22 checkpoint:

```text
[WSL-DEBIAN] working umbrella exists at ~/repos/carbonstack_umbrella
[WSL-DEBIAN] carbonstack is at 328ffc4
[WSL-DEBIAN] carbonstack-comms is at 012c8bf
[WSL-DEBIAN] carbonstack-cypher is at 6f92c34
[WSL-DEBIAN] carbonstack-os is at 1bbbe52
[GO-RUNNER] doctor passed after v0.3.22 docs/recon commit
[GO-RUNNER] core passed after v0.3.22 docs/recon commit
[ARTIFACT CLEANUP] expected OpenMLS generated roots were removed after final validation
```

Observed WSL Debian toolchain baseline:

```text
git version 2.47.3
go version go1.24.4 linux/amd64
rustc 1.96.0 (ac68faa20 2026-05-25)
cargo 1.96.0 (30a34c682 2026-05-25)
sqlite3 3.46.1
```

Observed `core` validation result:

```text
pre-test artifact scan: PASS / no generated/private/build artifact hits
OpenMLS real-Cypher lifecycle: PASS
carbonstack-comms package tests: PASS
carbonstack-cypher package tests: PASS
post-test artifact scan: expected OpenMLS sidecar generated roots only
VALIDATION PASSED
```

Expected post-test generated roots:

```text
carbonstack-comms/internal/protocol/mls/openmls-sidecar/.carbonstack-openmls-sidecar-state
carbonstack-comms/internal/protocol/mls/openmls-sidecar/target
```

These roots are expected only after tests and must not be committed or treated as source artifacts.

---

## 4. v0.3.22 Work Completed

### 4.1 Cypher local operator surface recon doc landed

New/updated docs at v0.3.22:

```text
carbonstack/docs/141-cypher-local-operator-surface-recon-v0.md
carbonstack/docs/README.md
carbonstack/roadmap/ROADMAP.md
```

Commit:

```text
328ffc4 docs: record Cypher local operator surface recon
```

The recon doc is intentionally verbose/generated. It inspects current Cypher, Comms, OpenMLS sidecar, and runner surfaces before code changes. It is not a polished operator manual and not an implementation result.

Doc presence confirmed:

```text
docs/141-cypher-local-operator-surface-recon-v0.md
size observed: 133274 bytes
```

Docs pointers confirmed:

```text
docs/README.md: docs/141-cypher-local-operator-surface-recon-v0.md
roadmap/ROADMAP.md: v0.3.22 — Cypher local operator surface recon
```

### 4.2 CarbonStack Relay Space terminology adopted

v0.3.22 records the terminology decision:

```text
CarbonStack Relay Space
```

Working definition:

```text
A CarbonStack Relay Space is one addressable Cypher relay context.
In the early model, one relay space maps to one Cypher instance or isolated Cypher profile.
Future tooling may allow one Cypher install to supervise multiple isolated relay spaces, but each relay space should keep separate address/config/data/operator boundaries.
```

This replaces loose “IRC-like” wording except when historical comparison is useful. The intended meaning is self-hosted, addressable, blind relay context — not IRC culture, moderation semantics, public channels, or a many-room community-server model.

### 4.3 Cypher config surface confirmed

Current Cypher config defaults observed:

```text
CYPHER_ADDR default: :8080
CYPHER_DB default: cypher.db
CYPHER_MIGRATIONS default: migrations
CYPHER_DEV_INVITE default: dev-invite
```

Relevant code surface:

```text
carbonstack-cypher/internal/config/config.go
carbonstack-cypher/cmd/cypher/main.go
```

Important interpretation:

```text
:8080 is the current code default and may bind more broadly than the local-only profile should imply.
v0.3.21/v0.3.22 docs recommend 127.0.0.1:8080 for local-only operator use.
Do not expose 0.0.0.0/LAN/public ingress/cloudflared/systemd yet.
```

### 4.4 Cypher DB/migration surface confirmed

Current migration files:

```text
carbonstack-cypher/migrations/001_init.sql
carbonstack-cypher/migrations/002_envelope_payload_metadata.sql
```

Current schema categories:

```text
invites
accounts
devices
envelopes
envelope_acks
```

Current payload metadata migration:

```text
ALTER TABLE envelopes ADD COLUMN payload_sha256 TEXT;
ALTER TABLE envelopes ADD COLUMN payload_size_bytes INTEGER;
```

Critical recon result:

```text
No schema_migrations tracking was observed.
001_init.sql uses CREATE TABLE IF NOT EXISTS patterns.
002_envelope_payload_metadata.sql uses ALTER TABLE ADD COLUMN without an observed idempotency guard.
The Store.Migrate path reads migration files and applies them, so reapplying an already-applied ALTER TABLE migration may become a persistent DB restart/upgrade hazard.
```

Current safe claim:

```text
Fresh DB behavior is known-good.
Persistent local DB behavior remains experimental.
Schema reset may be required between versions.
Before stronger local operator persistence claims, Cypher likely needs schema_migrations tracking or explicit wipe-only experimental DB documentation.
```

### 4.5 Cypher API/relay surface confirmed

Current HTTP route surface observed in `carbonstack-cypher/internal/httpapi/api.go`:

```text
GET  /v0/health
POST /v0/dev/invites
POST /v0/invites/claim
POST /v0/devices/register
GET  /v0/accounts/
POST /v0/envelopes
GET  /v0/devices/
POST /v0/envelopes/
```

Current API behavior surfaces include:

```text
dev invite creation
invite claim
device registration
account device listing
envelope submit
device inbox/retrieve
envelope ack
supported content type checks
supported protocol checks
payload hash/size metadata
idempotent recipient ack behavior
```

Interpretation:

```text
Cypher remains a relay/storage component, not a plaintext trust root.
Cypher should not decide client trust, MLS group validity, or message truth.
Cypher currently stores relay/server state and opaque encrypted payloads.
```

### 4.6 Comms-to-Cypher seam confirmed

Current Comms runtime flags/defaults observed:

```text
internal/app/commands.go: --server flag
internal/state/state.go: DefaultServerURL = http://localhost:8080
```

Current client surface:

```text
internal/client/cypher.go
```

Current real-Cypher backbone proof:

```text
internal/protocol/openmls_sidecar_real_cypher_relay_test.go
```

Interpretation:

```text
Comms already has a server URL seam via --server and state/default server URL.
v0.3.x should not wire runtime send/inbox to OpenMLS yet.
CARBONSTACK_CYPHER_URL remains a plausible future env seam, but current code already has a CLI/state server URL path.
Runtime Comms OpenMLS send/inbox integration remains v0.4.x.
```

### 4.7 OpenMLS sidecar critical surface confirmed

The recon grep confirmed continued relevance of:

```text
provider-storage
signer.json
KeyPackage
Welcome
application-message
message-protect
message-open
conversation-create
conversation-join
identity-create
```

Interpretation:

```text
OpenMLS sidecar remains dev/pre-alpha infrastructure.
Generated sidecar state must remain scoped to known generated roots.
Secret-bearing signer/provider material must never be committed, pasted, or normalized as safe.
The OpenMLS real-Cypher lifecycle is backbone validation, not production storage/vault proof.
```

### 4.8 Runner/validation surface confirmed

Relevant runner profiles remain:

```text
doctor
core
full
release-snapshot
write-checksums
verify-checksums
```

Interpretation:

```text
Do not add a local-backbone runner profile yet.
The local operator success/failure contract is not concrete enough.
A future local-backbone profile may make sense after DB persistence, config/data layout, start/stop, and runbook behavior are defined.
```

---

## 5. Critical Blunders / Lessons Preserved

### 5.1 Generated recon docs can be huge; that is acceptable for this rung

`docs/141-cypher-local-operator-surface-recon-v0.md` is large because it captured grep-driven recon across Cypher, Comms, OpenMLS sidecar, and runner surfaces. This is acceptable because v0.3.22 is a recon rung, not a polished public guide.

Do not treat the size as a reason to compress away operational evidence in LogDoc. Polish can happen later in an operator runbook.

### 5.2 Current Cypher default bind is not the desired local-only operator bind

The code default is currently:

```text
CYPHER_ADDR=:8080
```

The local-only operator recommendation is:

```text
CYPHER_ADDR=127.0.0.1:8080
```

This distinction matters. `:8080` should not be casually described as local-only. v0.3.22 should not patch it yet unless the next implementation rung decides to harden the local operator default. For now it is documented as a deployability surface.

### 5.3 Migration risk is now concrete, not theoretical

The recon confirmed `002_envelope_payload_metadata.sql` uses raw `ALTER TABLE ... ADD COLUMN` statements and no observed `schema_migrations` tracking. This is the strongest candidate for the next code rung.

The project should not claim persistent local operator DB stability until repeated startup / already-migrated DB behavior is tested or hardened.

### 5.4 Current dev invites are still not final invites

Cypher dev invite behavior exists and is useful for local/test bootstrapping, but it must remain separated from future CarbonStack concepts:

```text
relay access invite
conversation membership invite
device enrollment invite
operator/admin invite
QR verification ceremony
hardware-key-backed enrollment/recovery ceremony
```

### 5.5 Existing Comms server URL seam reduces future work

The recon confirmed Comms already has a `--server` flag / state default path. The future `CARBONSTACK_CYPHER_URL` idea remains plausible but is not strictly the only way to connect Comms to Cypher. The eventual v0.4 runtime integration should build on the existing seam where sensible.

### 5.6 Artifact cleanup remains required after validation

The final validation produced expected sidecar generated roots, then cleanup was run. Final repo status snapshot showed clean status after cleanup. Continue this pattern.

---

## 6. Current Blockers / Not Validated

Current blockers and nonvalidated areas:

```text
Local-only operator profile is not implemented.
Persistent local Cypher DB restart/upgrade behavior is not validated.
Cypher migration tracking/idempotency is not hardened.
Local operator runbook is not implemented beyond recon/planning.
No local-backbone runner profile exists yet.
Runtime Comms send/inbox remains stub-era and not wired as user-facing OpenMLS messenger UX.
Production-safe provider storage and secure vault design are unsolved.
Hostile-server rollback/replay/metadata-abuse validation remains future work.
Android/CarbonStackOS implementation remains later work.
CarbonStack Relay Space mechanics remain a defined future concept, not an implementation.
Public ingress, systemd, cloudflared, real homelab deployment, and production service paths remain out of scope.
```

Hard nonclaims remain:

```text
Do not claim production readiness.
Do not claim production E2EE.
Do not claim hostile-server safety.
Do not claim metadata privacy.
Do not claim Debian deployability.
Do not claim systemd readiness.
Do not claim cloudflared readiness.
Do not claim Android readiness.
Do not claim secure vault/storage.
Do not claim external audit or certification.
```

---

## 7. Next Safest Actions

Recommended next implementation sequence after this v0.3.22 breakpoint:

```text
v0.3.23 preferred:
  Cypher migration/persistence hardening recon-to-code rung.
```

Recommended concrete next rung:

```text
1. Add or plan a schema_migrations table.
2. Make migration application one-time/version-tracked.
3. Add tests for:
   - fresh DB migration;
   - reopening the same already-migrated DB;
   - applying migration 002 only once;
   - preserving existing envelope metadata columns.
4. Keep SQLite.
5. Keep local-only operator profile claims experimental until tests pass.
```

Alternative if not coding yet:

```text
v0.3.23 docs-only:
  write a focused Cypher migration hardening design doc before code.
```

Do **not** jump directly to:

```text
public ingress
cloudflared
systemd
real homelab
runtime Comms send/inbox OpenMLS UX
Android app
CarbonStackOS
remote admin plane
CarbonStack Relay Space implementation
local-backbone runner profile
```

---

## 8. Critical Paths / Commands

### WSL Debian active working umbrella

```text
[WSL-DEBIAN] umbrella: ~/repos/carbonstack_umbrella
[WSL-DEBIAN] carbonstack: ~/repos/carbonstack_umbrella/carbonstack
[WSL-DEBIAN] carbonstack-comms: ~/repos/carbonstack_umbrella/carbonstack-comms
[WSL-DEBIAN] carbonstack-cypher: ~/repos/carbonstack_umbrella/carbonstack-cypher
[WSL-DEBIAN] carbonstack-os: ~/repos/carbonstack_umbrella/carbonstack-os
[GO-RUNNER] runner dir: ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
```

### WSL Debian known-good validation commands

From runner dir:

```text
cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
go run . --profile doctor
go run . --profile core
```

From carbonstack repo root:

```text
cd ~/repos/carbonstack_umbrella/carbonstack
go run ./tools/carbonstack-validate --profile doctor
go run ./tools/carbonstack-validate --profile core
```

### Artifact cleanup after validation

```text
cd ~/repos/carbonstack_umbrella/carbonstack-comms
rm -rf internal/protocol/mls/openmls-sidecar/.carbonstack-openmls-sidecar-state
rm -rf internal/protocol/mls/openmls-sidecar/target
```

### v0.3.22 docs

```text
[DOC] v0.3.22 recon:
carbonstack/docs/141-cypher-local-operator-surface-recon-v0.md

[DOC] docs index:
carbonstack/docs/README.md

[DOC] roadmap:
carbonstack/roadmap/ROADMAP.md
```

### Cypher deployability surfaces

```text
[CODE] cypher config:
carbonstack-cypher/internal/config/config.go

[CODE] cypher main:
carbonstack-cypher/cmd/cypher/main.go

[CODE] cypher DB:
carbonstack-cypher/internal/db/db.go

[CODE] cypher API:
carbonstack-cypher/internal/httpapi/api.go

[CODE] cypher tests:
carbonstack-cypher/internal/httpapi/api_test.go

[SQL] migrations:
carbonstack-cypher/migrations/001_init.sql
carbonstack-cypher/migrations/002_envelope_payload_metadata.sql
```

### Comms deployability surfaces

```text
[CODE] Comms commands:
carbonstack-comms/internal/app/commands.go

[CODE] Comms state/default server:
carbonstack-comms/internal/state/state.go

[CODE] Cypher client:
carbonstack-comms/internal/client/cypher.go

[TEST] real Cypher OpenMLS lifecycle:
carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go
```

---

## 9. Lean Breakpoint Summary

```text
CarbonStack v0.3.22 is a stable mainline docs/recon breakpoint after v0.3.21. It records a WSL Debian-first Cypher local operator surface recon, adopts CarbonStack Relay Space terminology, confirms Cypher config defaults/API/migration surfaces, identifies raw SQLite migration reapplication as the highest-priority deployability risk, confirms Comms already has a server URL seam while runtime OpenMLS UX remains v0.4.x, validates doctor/core after the docs commit, and leaves persistent local operator profile implementation, migration hardening, local runbook, local-backbone runner profile, Relay Space mechanics, public ingress, systemd, cloudflared, homelab, Android, CarbonStackOS, hostile-server proof, secure vault, and production/security claims deferred.
```

---

## 10. Preserved v0.3.21 Operational Process Log

The following section preserves the previous v0.3.21 LogDoc body as operational continuity. It is intentionally retained rather than compressed away. Where v0.3.21 text conflicts with the v0.3.22 current-state overlay above, v0.3.22 wins for current state; v0.3.21 remains the provenance/process ledger for the local-only backbone deployability recon and the preserved v0.3.20 public runner-backed testing release history.


# CarbonStack LogDoc v0.3.21

**Last updated:** 2026-06-02 local session  
**Owner/operator:** bitcrusher32  
**Assistant/model:** GPT-5.5 Thinking  
**Current phase:** Phase 2E-prep / v0.3.x **local-only backbone deployability recon checkpoint**. `carbonstack` is now at `6850c58 docs: record local backbone deployability recon`, after `v0.3.20` was cut at `3eeb1a1e1a / 3eeb1a1 docs: clean public release surface`. `carbonstack-comms` remains at `012c8bf scripts: support source snapshot artifact guard`; `carbonstack-cypher` remains at `6f92c34 chore: fix readme formatting`; `carbonstack-os` remains at `1bbbe52 docs: clarify CarbonStackOS target direction`. The live WSL Debian working umbrella at `~/repos/carbonstack_umbrella` passed `doctor` and `core` through the Go validation runner after the v0.3.21 docs/recon commit.

**Version schema:** `v[scope].[timeline]`. This file is `v0.3.21`, the local-only backbone deployability recon handoff after the `v0.3.20` public runner-backed testing release checkpoint. Per LogDoc V2 practice, this Markdown file intentionally preserves verbose operational continuity; the JSON breakpoint remains the lean handoff state. The dense v0.3.20 operational process log is preserved later in this file rather than compressed away.

**LogDoc V2 usage note:** LogDoc V2 uses a split between a Markdown LogDoc for working-session continuity, a JSON state file for end-of-session handoff, Git for history preservation, and README/docs for public claims. This v0.3.21 Markdown file therefore intentionally stays “bloated” where operational process, blunders, validation ladders, and historical context matter. The v0.3.21 JSON should remain lean and should agree with this file on the latest current state.

---

## 1. Project Goal

**Active goal:** Preserve CarbonStack state after completing the v0.3.21 WSL Debian local-only backbone deployability recon checkpoint. The immediate next line of work is not another public release, not runtime Comms UX, and not OS/Android work. The next line is to turn the v0.3.21 recon into concrete local-only backbone deployability planning/implementation, beginning with Cypher persistence/config/migration behavior and a server-like local operator model.

CarbonStack’s broader project intent remains:

- **CarbonStack** is the secure-communications appliance-stack project and shared doctrine/release authority.
- **CarbonStackOS** is the future constrained Android-derived appliance OS.
- **CarbonStackComms** is the text-first encrypted messaging client and current OpenMLS sidecar/relay integration surface.
- **CarbonStackCypher** is the self-hostable relay/storage stack.
- The `carbonstack` repo remains the doctrine/spec/docs/release authority; implementation repos must not contradict it without revising the spec.

**North star:**

> Its goal is not to be a better smartphone. Its goal is to stop being a smartphone.

**Core doctrine:**

> Every feature is guilty until it proves it does not add unacceptable parser, network, sensor, identity, or filesystem authority.

**Current reality after v0.3.21:** `v0.3.20` remains the current public runner-backed testing release. `v0.3.21` is a mainline docs/recon checkpoint after that release. It records the WSL Debian working development environment, local-only deployability model, Cypher config/data/migration surfaces, Comms-to-Cypher addressing seam, OpenMLS sidecar generated-state boundaries, future relay-space direction, and nonclaims before local deployability implementation begins.

The current public release page remains:

```text
https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.3.20
```

The current public testing interpretation remains:

```text
v0.3.20 is the preferred public runner-backed testing release.
v0.3.4 remains available as the older Windows 11 / PowerShell legacy testbed.
Debian / WSL Debian is now the mainline public dev/test validation direction.
Windows 11 short-path validation was the final explicit Windows dev/test validation for this phase.
```

This is still **not** a deployability release. It does **not** prove production readiness, production E2EE, hostile-server safety, metadata privacy, Android readiness, secure vault/storage, Debian deployability, systemd readiness, cloudflared readiness, public ingress safety, real homelab validation, external audit, or certification.

---

## 2. Current Repo Heads

```text
carbonstack        6850c58 docs: record local backbone deployability recon
carbonstack-comms  012c8bf scripts: support source snapshot artifact guard
carbonstack-cypher 6f92c34 chore: fix readme formatting
carbonstack-os     1bbbe52 docs: clarify CarbonStackOS target direction
```

Current `carbonstack` public runner-backed testing release tag:

```text
v0.3.20 -> 3eeb1a1e1a / 3eeb1a1 docs: clean public release surface
```

Legacy `carbonstack` Windows/PowerShell testbed release tag:

```text
v0.3.4 -> 7c5b7e6 docs: record release snapshot self-test fix
```

Historical `carbonstack` release tag:

```text
v0.3.0 -> 59245fc docs: freeze v0.3.0 release packaging plan
```

Important continuity note: `v0.3.21` does not supersede `v0.3.20` as a public release. It advances mainline documentation/recon after the v0.3.20 public runner-backed testing release. It is safe to pause here because the recon doc was pushed and the WSL Debian working umbrella validated afterward.

---

## 3. Current Validated State

Validated at the v0.3.21 checkpoint:

```text
[WSL-DEBIAN] working umbrella exists at ~/repos/carbonstack_umbrella
[WSL-DEBIAN] carbonstack is at 6850c58
[WSL-DEBIAN] carbonstack-comms is at 012c8bf
[WSL-DEBIAN] carbonstack-cypher is at 6f92c34
[WSL-DEBIAN] carbonstack-os is at 1bbbe52
[GO-RUNNER] doctor passed
[GO-RUNNER] core passed
```

Observed WSL Debian toolchain baseline:

```text
git version 2.47.3
go version go1.24.4 linux/amd64
rustc 1.96.0 (ac68faa20 2026-05-25)
cargo 1.96.0 (30a34c682 2026-05-25)
sqlite3 3.46.1
```

Observed `core` validation result:

```text
pre-test artifact scan: PASS / no generated/private/build artifact hits
OpenMLS real-Cypher lifecycle: PASS
carbonstack-comms package tests: PASS
carbonstack-cypher package tests: PASS
post-test artifact scan: expected OpenMLS sidecar generated roots only
VALIDATION PASSED
```

Expected post-test generated roots:

```text
carbonstack-comms/internal/protocol/mls/openmls-sidecar/.carbonstack-openmls-sidecar-state
carbonstack-comms/internal/protocol/mls/openmls-sidecar/target
```

These roots are expected only after tests and must not be committed or treated as source artifacts.

---

## 4. v0.3.21 Work Completed

### 4.1 WSL Debian working umbrella restored as live dev area

The operator escaped the extracted v0.3.20 validation package and created/validated a proper live WSL Debian sibling-repo working area:

```text
~/repos/carbonstack_umbrella
```

This matters because earlier commands were accidentally run from:

```text
~/carbonstack-v0.3.20-test/package/carbonstack/tools/carbonstack-validate
```

That path is a public release validation extraction, not a live Git working umbrella. From inside `tools/carbonstack-validate`, commands shaped like `go run ./tools/carbonstack-validate` resolve incorrectly to a nested non-existent path. Correct commands from the runner directory are:

```text
go run . --profile doctor
go run . --profile core
```

Correct commands from the `carbonstack` repo root are:

```text
go run ./tools/carbonstack-validate --profile doctor
go run ./tools/carbonstack-validate --profile core
```

### 4.2 v0.3.21 local deployability recon doc landed

New/updated docs at v0.3.21:

```text
carbonstack/docs/140-local-backbone-deployability-recon-v0.md
carbonstack/docs/README.md
carbonstack/roadmap/ROADMAP.md
```

Commit:

```text
6850c58 docs: record local backbone deployability recon
```

The recon doc records:

- current validated WSL Debian baseline;
- local-only deployability goal;
- explicit non-goals;
- local-only operator model;
- Cypher config surface;
- SQLite/database/migration position;
- Comms-to-Cypher addressing seam;
- OpenMLS sidecar generated-state boundaries;
- Go runner/validation surface;
- future relay-space / IRC-like mechanics;
- future invite/join category separation;
- future admin direction;
- recommended next implementation rungs;
- current claim boundary.

### 4.3 v0.3.21 design decisions captured

Current decisions:

```text
v0.3.21 is docs + recon first.
Initial deployability is local-only.
WSL Debian is the primary active dev/test environment.
Windows is now secondary/legacy confirmation, not the mainline deployability target.
Runtime Comms OpenMLS send/inbox wiring remains v0.4.x.
Cloudflared/public ingress/systemd/real homelab remain deferred.
Cypher should use simple explicit configuration first.
SQLite remains appropriate for v0.3.21 local-only single-node relay testing.
Persistent DB behavior is experimental until migration tracking/idempotency is hardened.
```

Recommended local-only Cypher operator shape from the recon:

```text
CYPHER_ADDR=127.0.0.1:8080
CYPHER_DB=$HOME/.local/share/carbonstack/cypher/cypher.db
CYPHER_MIGRATIONS=./migrations
CYPHER_DEV_INVITE=dev-invite
```

Recommended Comms/Cypher addressing seam:

```text
CARBONSTACK_CYPHER_URL=http://127.0.0.1:8080
```

Future possible CLI flag:

```text
--cypher-url http://127.0.0.1:8080
```

### 4.4 Relay-space / IRC-like mechanics refined but deferred

The project’s later “IRC-like” meaning is **not** IRC culture/moderation. It means self-hosted, addressable relay spaces.

Current future direction:

```text
One addressable Cypher relay space maps to one server/conversation context.
One relay/IP/link/address thing corresponds to one isolated relay space.
Multiple relay spaces may later be implemented as multiple Cypher instances or isolated profiles under one install.
```

This is intentionally simpler than a many-room general-purpose server model. It may increase deployment overhead, but reduces early attack surface, metadata coupling, cross-space administration complexity, and trust-scope confusion.

The recon explicitly defers implementation of relay-space mechanics until after local deployability is grounded.

### 4.5 Invite and admin concepts separated for future work

Current dev invites remain temporary bootstrap plumbing only.

Future invite categories should remain separate:

```text
relay access invite
conversation membership invite
device enrollment invite
operator/admin invite
QR verification ceremony
hardware-key-backed enrollment/recovery ceremony
```

Long-term admin direction may be app-mediated and signed by an operator/admin identity or hardware-backed key. For v0.3.21, there is no remote admin plane, web dashboard, public operator workflow, or admin UX.

---

## 5. Critical Blunders / Lessons Preserved

### 5.1 Release validation extraction is not a working clone

The operator initially attempted to use the v0.3.20 extracted validation package as if it were the live WSL working clone. This produced expected failures:

```text
cd: /home/Bitcrusher32/repos/carbonstack_umbrella/carbonstack: No such file or directory
fatal: not a git repository
stat .../tools/carbonstack-validate/tools/carbonstack-validate: directory not found
```

Correct interpretation:

```text
~/carbonstack-v0.3.20-test/package/... = extracted public validation package
~/repos/carbonstack_umbrella/... = live WSL Debian working clone area
```

### 5.2 Runner invocation depends on current directory

From `carbonstack` repo root:

```text
go run ./tools/carbonstack-validate --profile doctor
go run ./tools/carbonstack-validate --profile core
```

From `carbonstack/tools/carbonstack-validate`:

```text
go run . --profile doctor
go run . --profile core
```

Do not run `go run ./tools/carbonstack-validate` from inside `tools/carbonstack-validate`.

### 5.3 Persistent DB deployability is not the same as test DB success

Cypher test flows can use fresh or temporary DBs. A local operator profile raises a different question: can a persistent SQLite DB safely survive restart and schema evolution?

Current safe stance:

```text
Fresh DB behavior is known-good.
Persistent local DB behavior is experimental.
Schema reset may be required between versions.
Migration tracking or idempotent migration guards are likely needed before stronger persistence claims.
```

### 5.4 SQLite remains correct for v0.3.21

SQLite is acceptable for v0.3.21 local-only single-node relay testing because it is simple, inspectable, easy to package, and avoids introducing a separate server dependency before the data/operator model stabilizes.

PostgreSQL or another OSS DB may be considered later if the operator/deployability model matures enough to justify it. Do not switch DBs before the Cypher data model stabilizes.

### 5.5 Generated OpenMLS sidecar roots remain expected after tests only

The v0.3.21 validation run produced expected post-test artifacts under known generated roots:

```text
.carbonstack-openmls-sidecar-state
target
```

These are acceptable after tests only when scoped under the known OpenMLS sidecar generated root. They remain forbidden as source/package pre-test artifacts.

---

## 6. Current Blockers / Not Validated

Current blockers and nonvalidated areas:

```text
Local-only deployability is planned/reconned but not implemented.
Persistent local Cypher DB restart/upgrade behavior is not validated.
Cypher migration tracking/idempotency is not hardened.
Local operator runbook is not implemented beyond recon/planning.
No local-backbone runner profile exists yet.
Runtime Comms send/inbox remains stub-era and not wired as user-facing OpenMLS messenger UX.
Production-safe provider storage and secure vault design are unsolved.
Hostile-server rollback/replay/metadata-abuse validation remains future work.
Android/CarbonStackOS implementation remains later work.
IRC-like relay-space/server-system mechanics remain deferred until local deployability is grounded.
Public ingress, systemd, cloudflared, real homelab deployment, and production service paths remain out of scope.
```

Hard nonclaims remain:

```text
Do not claim production readiness.
Do not claim production E2EE.
Do not claim hostile-server safety.
Do not claim metadata privacy.
Do not claim Debian deployability.
Do not claim systemd readiness.
Do not claim cloudflared readiness.
Do not claim Android readiness.
Do not claim secure vault/storage.
Do not claim external audit or certification.
```

---

## 7. Next Safest Actions

Recommended next implementation sequence after this v0.3.21 breakpoint:

```text
v0.3.22 candidate direction A:
  Cypher migration/persistence hardening plan and/or first implementation.

v0.3.22 candidate direction B:
  Local-only Cypher operator config/data-dir runbook skeleton.
```

My recommendation after v0.3.21:

```text
1. Recon current Cypher migration code and tests.
2. Decide schema_migrations vs idempotent migration guards.
3. Add tests for fresh DB migration and reopening an already-migrated DB.
4. Only then make persistent local operator DB claims stronger.
5. Build the local-only runbook/config convention around that behavior.
```

Do **not** jump directly to:

```text
public ingress
cloudflared
systemd
real homelab
runtime Comms send/inbox OpenMLS UX
Android app
CarbonStackOS
remote admin plane
IRC-like relay-space implementation
```

---

## 8. Critical Paths / Commands

### WSL Debian active working umbrella

```text
[WSL-DEBIAN] umbrella: ~/repos/carbonstack_umbrella
[WSL-DEBIAN] carbonstack: ~/repos/carbonstack_umbrella/carbonstack
[WSL-DEBIAN] carbonstack-comms: ~/repos/carbonstack_umbrella/carbonstack-comms
[WSL-DEBIAN] carbonstack-cypher: ~/repos/carbonstack_umbrella/carbonstack-cypher
[WSL-DEBIAN] carbonstack-os: ~/repos/carbonstack_umbrella/carbonstack-os
[GO-RUNNER] runner dir: ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
```

### WSL Debian known-good validation commands

From runner dir:

```text
cd ~/repos/carbonstack_umbrella/carbonstack/tools/carbonstack-validate
go run . --profile doctor
go run . --profile core
```

From carbonstack repo root:

```text
cd ~/repos/carbonstack_umbrella/carbonstack
go run ./tools/carbonstack-validate --profile doctor
go run ./tools/carbonstack-validate --profile core
```

### Artifact cleanup after validation

```text
cd ~/repos/carbonstack_umbrella/carbonstack-comms
rm -rf internal/protocol/mls/openmls-sidecar/.carbonstack-openmls-sidecar-state
rm -rf internal/protocol/mls/openmls-sidecar/target
```

---

## 9. Lean Breakpoint Summary

```text
CarbonStack v0.3.21 is a stable mainline docs/recon breakpoint after the v0.3.20 public runner-backed testing release. It records the WSL Debian live sibling-repo working umbrella, validates doctor/core after the local deployability recon commit, defines the local-only backbone deployability model, identifies Cypher persistence/migration as the highest-value next technical risk, and defers runtime Comms UX, public ingress, systemd, cloudflared, homelab, Android, CarbonStackOS, hostile-server proof, secure vault, and production/security claims.
```

---

## 10. Preserved v0.3.20 Operational Process Log

The following section preserves the previous v0.3.20 LogDoc body as operational continuity. It is intentionally retained rather than compressed away. Where v0.3.20 text conflicts with the v0.3.21 current-state overlay above, v0.3.21 wins for current state; v0.3.20 remains the provenance/process ledger for the public runner-backed testing release and the preceding v0.3.x validation ladder.

# CarbonStack LogDoc v0.3.20

**Last updated:** 2026-06-02 local session  
**Owner/operator:** bitcrusher32  
**Assistant/model:** GPT-5.5 Thinking  
**Current phase:** Phase 2E-prep / v0.3.x **public runner-backed testing release checkpoint**. `carbonstack` is now at `3eeb1a1 docs: clean public release surface` and tagged `v0.3.20` at `3eeb1a1e1a`. `v0.3.20` is the first public runner-backed testing release and supersedes `v0.3.4` as the preferred public testing path. `v0.3.4` remains available as the older Windows 11 / PowerShell legacy testbed at `7c5b7e6`. `carbonstack-comms` remains at `012c8bf scripts: support source snapshot artifact guard`; `carbonstack-cypher` remains at `6f92c34 chore: fix readme formatting`; `carbonstack-os` remains at `1bbbe52 docs: clarify CarbonStackOS target direction`.

**Version schema:** `v[scope].[timeline]`. This file is `v0.3.20`, the public runner-backed testing release handoff after the `v0.3.19` local downloaded-asset simulation checkpoint. The dense v0.3.0[PRIME] timeline remains compressed; repo roles, critical paths, validation lessons, blunders, nonclaims, and future direction are preserved.

---

## 1. Project Goal

**Active goal:** Preserve CarbonStack state after cutting the first public runner-backed testing release, `v0.3.20`, while explicitly recording the Debian-first release transition, the final explicit Windows dev/test validation for this phase, and the remaining nonclaims before real deployability work begins.

CarbonStack’s broader project intent remains:

- **CarbonStack** is the secure-communications appliance-stack project and shared doctrine/release authority.
- **CarbonStackOS** is the future constrained Android-derived appliance OS.
- **CarbonStackComms** is the text-first encrypted messaging client and current OpenMLS sidecar/relay integration surface.
- **CarbonStackCypher** is the self-hostable relay/storage stack.
- The `carbonstack` repo remains the doctrine/spec/docs/release authority; implementation repos must not contradict it without revising the spec.

**North star:**

> Its goal is not to be a better smartphone. Its goal is to stop being a smartphone.

**Core doctrine:**

> Every feature is guilty until it proves it does not add unacceptable parser, network, sensor, identity, or filesystem authority.

**Current reality after v0.3.20:** `v0.3.20` is now the public runner-backed testing release. The release page is live at:

```text
https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.3.20
```

The tag/release points at `3eeb1a1e1a` / `3eeb1a1 docs: clean public release surface`. The release page is titled **CarbonStack v0.3.20 Runner-Backed Testing Release** and is marked pre-release. It publishes the intended multi-repo runner-backed package and metadata assets:

```text
carbonstack-v0.3.20-runner-backed-testing-release.tgz
carbonstack-v0.3.20-release-manifest.json
carbonstack-v0.3.20-package-checksums.txt
carbonstack-v0.3.20-asset-checksums.txt
carbonstack-v0.3.20-validation-freeze.md
v0.3.20-testing-runbook.md
v0.3.20-release-notes.md
LICENSE
```

The v0.3.20 release validates the current Cypher + Comms OpenMLS relay backbone through the Go runner (`verify-checksums` and `release-snapshot`), real package-internal checksums, release metadata, fresh package extraction validation, OpenMLS real-Cypher lifecycle validation, Comms package tests, Cypher package tests, and post-test artifact scoping.

Primary release target:

```text
Debian 13 / WSL2 Debian, linux/amd64
```

Secondary/final explicit Windows dev/test validation for this phase:

```text
Windows 11, windows/amd64
short extraction root: C:\cs-v0320\package
CARGO_BUILD_JOBS=1
```

Important platform interpretation:

```text
v0.3.20 supersedes v0.3.4 as the preferred public testing path.
v0.3.4 remains available as the older Windows 11 / PowerShell legacy testbed.
v0.3.20 is expected to be the final release in this phase where Windows dev/test validation is explicitly provided as part of the runner-backed release surface.
After v0.3.20, CarbonStack mainline public dev/test releases should migrate to Debian / WSL Debian first.
This does not mean Windows is permanently unsupported; later Linux-family, BSD, or Windows ports may be reconsidered after the server/backbone stack is mature.
```

This is still **not** a deployability release. It does **not** prove production readiness, production E2EE, hostile-server safety, metadata privacy, Android readiness, secure vault/storage, Debian deployability, systemd readiness, cloudflared readiness, public ingress safety, real homelab validation, external audit, or certification.

## 2. Current Repo Heads

```text
carbonstack        3eeb1a1 docs: clean public release surface
carbonstack-comms  012c8bf scripts: support source snapshot artifact guard
carbonstack-cypher 6f92c34 chore: fix readme formatting
carbonstack-os     1bbbe52 docs: clarify CarbonStackOS target direction
```

Current `carbonstack` public runner-backed testing release tag:

```text
v0.3.20 -> 3eeb1a1e1a / 3eeb1a1 docs: clean public release surface
```

Legacy `carbonstack` Windows/PowerShell testbed release tag:

```text
v0.3.4 -> 7c5b7e6 docs: record release snapshot self-test fix
```

Historical `carbonstack` release tag:

```text
v0.3.0 -> 59245fc docs: freeze v0.3.0 release packaging plan
```

Important continuity note: `v0.3.20` is the public runner-backed testing release checkpoint after the `v0.3.19` local downloaded-asset simulation / release-readiness planning checkpoint. It moved the public preferred testing path from the older v0.3.4 Windows 11 / PowerShell flow to the v0.3.20 Debian-first runner-backed package flow. It does not prove deployability or production security.

## 3. Compressed Timeline Context

### v0.3.0

- Tagged `carbonstack` at `59245fc`.
- Published the first public experimental backbone epoch release.
- Packaged the known-good Cypher + Comms OpenMLS relay backbone proof.
- Preserved hard nonclaims and the Gitea default source-archive caveat.

### v0.3.1

- Downloaded and verified v0.3.0 public release assets.
- Corrected `carbonstack-comms/go.sum` from required to optional/include-if-present because it never existed.
- Added security claim validation policy and future adversarial-validation/self-pentest gate.

### v0.3.2

- Clarified public identity: end goal is the secure-communications appliance stack; current artifact is the v0.3.x experimental backbone.
- Renamed `[SANITIZED]LogdocList` to `sanitized-project-logdoc-list`.
- Clarified CarbonStackOS target direction and fixed CarbonStackOS license continuity.

### v0.3.3

- Tested clean v0.3.0 source snapshots from extraction.
- Confirmed core validation worked from extraction.
- Found the public Comms self-test wrapper failed because `check-no-rust-artifacts.ps1` assumed Git metadata in source snapshots without `.git/`.

### v0.3.4

- Fixed `carbonstack-comms/scripts/check-no-rust-artifacts.ps1`.
- Added Git working-tree mode and no-Git source snapshot filesystem fallback mode.
- Validated no-Git public self-test, no-Git public self-test `-Full`, no-Git CarbonStack `validate-local.ps1`, and live CarbonStack `validate-local.ps1`.
- Tagged and published `v0.3.4` as the recommended public testing bug hotfix release.
- Later release-surface update added explicit platform target language: `Platform: Windows 11`.

### v0.3.5

- Downloaded the public `v0.3.4` assets from the Gitea release page.
- Verified checksum-covered assets.
- Extracted source snapshots.
- Confirmed no obvious forbidden generated/private/build artifacts were present before tests.
- Ran the public no-Git self-test path from the downloaded/extracted v0.3.4 release snapshots.
- Added `docs/125-v0.3.4-public-release-asset-verification-v0.md`.

### v0.3.6

- Added `docs/126-v0.3.4-public-tester-runbook-v0.md`.
- Added/adjusted release/testing pointers.
- Preserved version-specific runbook policy and generated artifact policy.
- Added release README direction: main README should point to latest release for testing/dev, not carry every release-specific runbook forever.

### v0.3.7

- Added `docs/127-windows-reliance-and-debian-homelab-recon-v0.md`.
- Linked the recon from `docs/README.md`.
- Updated `roadmap/ROADMAP.md` with WSL Debian bridge sequencing before real homelab validation.
- Recorded that current Windows reliance is mostly script/runbook/validation-layer, not obviously core Go/Rust implementation.
- Recorded the Debian homelab as a plausible future target but not yet safe to deploy against.
- Decided WSL Debian is a good quick-portability bridge before touching the real Debian host.

### v0.3.8

- Added `docs/128-wsl-debian-quick-portability-setup-scout-v0.md`.
- Installed/confirmed Debian WSL target rather than relying on default Ubuntu WSL.
- Created WSL workspace at `~/carbonstack-wsl`.
- Copied current working repo snapshots into WSL.
- Confirmed WSL Debian setup baseline:
  - Debian GNU/Linux 13 trixie / WSL2 / linux amd64;
  - `git version 2.47.3`;
  - `go version go1.24.4 linux/amd64`;
  - Debian apt `rustc 1.85.0`;
  - Debian apt `cargo 1.85.0`;
  - `sqlite3 3.46.1`.
- Found artifact hygiene warning: `carbonstack-cypher/cypher.db` was copied into WSL workspace and must be excluded/removed before WSL tests.

### v0.3.9

- Added `docs/129-wsl-debian-quick-portability-test-v0.md`.
- Initial WSL Comms test failed because Debian apt Rust `1.85.0` was too old for `openmls 0.8.1`.
- Direct Rust sidecar probe revealed `error[E0658]: use of unstable library feature unsigned_is_multiple_of`.
- rustup stable installed Rust/Cargo `1.96.0`; shell PATH had to load `$HOME/.cargo/env`.
- Direct sidecar identity-create passed after Rust correction.
- Targeted OpenMLS real-Cypher lifecycle test passed on WSL Debian.
- Full `carbonstack-comms/internal/protocol` package passed on WSL Debian.
- Full `carbonstack-comms` package suite passed on WSL Debian.
- Full `carbonstack-cypher` package suite passed on WSL Debian.
- Post-test generated artifacts stayed in expected OpenMLS sidecar state and Rust `target/` roots.

### v0.3.10

- Added `docs/130-go-validation-runner-design-v0.md`.
- Recorded that validation should converge on an umbrella Go runner under `carbonstack/tools/carbonstack-validate`.
- Confirmed the umbrella runner should orchestrate repo-local tests rather than replacing them.
- Defined initial profile direction: `doctor`, `core`, and future `full`.
- Defined design non-goals: no deployment, no cloudflared/systemd, no package installation, no release publication, no runtime Comms UX, no security claims.
- Preserved intended authority transition:
  - `[WIN-PWRSHL]` remains current public release compatibility / legacy validation path.
  - `[WSL-DEBIAN]` is the preferred fast core-dev/test candidate.
  - `[GO-RUNNER]` is the future validation authority.
- Deferred real Debian homelab tests until deployability / IRC-style setup work actually needs them.

### v0.3.11

- Implemented the first Go validation runner in `carbonstack/tools/carbonstack-validate`.
- Added a standalone Go module for the runner.
- Added runner README.
- Implemented `--profile doctor`.
- Implemented `--profile core`.
- Implemented `--profile full` as an alias for `core` for now.
- Implemented explicit `--root` support.
- Fixed initial root inference so the runner can infer the umbrella root from nested paths such as `carbonstack/tools/carbonstack-validate` rather than assuming a hard-coded local Windows path.
- Fixed a PowerShell/BOM blunder: `Set-Content -Encoding UTF8` on Windows PowerShell 5.1 wrote a BOM to `go.mod`, which Go rejected as `unexpected input character '\ufeff'`; files were rewritten UTF-8 without BOM.
- Validated `[WIN-PWRSHL]` runner doctor/core path with explicit `--root`.
- Validated `[WSL-DEBIAN]` runner doctor/core path without explicit root from `~/carbonstack-wsl/carbonstack/tools/carbonstack-validate`.

### v0.3.12

- Hardened the Go validation runner at `carbonstack/tools/carbonstack-validate`.
- Updated runner README with clearer profile, root inference, Windows, WSL Debian, Rust toolchain, artifact scan, and boundary notes.
- Added `docs/131-go-validation-runner-implementation-v0.md`.
- Linked the v0.3.12 runner implementation/hardening record from `docs/README.md`.
- Updated `roadmap/ROADMAP.md` with the v0.3.12 Go validation runner hardening direction.
- Improved runner UX/durability:
  - prints executable paths for `go`, `rustc`, `cargo`, and `sqlite3`;
  - preserves root inference and explicit `--root`;
  - keeps `full` as an alias for `core`;
  - prints known Rust floor note;
  - warns if Rust appears older than the known-good floor;
  - improves artifact scan phase labels (`pre-test`, `post-test`);
  - classifies artifact hits as `known-openmls-sidecar-generated-root`, `research-generated-root`, `local-go-cache-root`, or `review`;
  - remains non-destructive.
- Preserved the Windows PowerShell 5.1 encoding lesson: `UTF8NoBOM` is not supported as a `Set-Content -Encoding` value in Windows PowerShell 5.1; use `Set-Content -Encoding UTF8` and then rewrite with `.NET UTF8Encoding($false)` for Go/module/doc files where BOM-free output matters.
- `carbonstack` advanced to `0f06da4 tools: harden CarbonStack validation runner`.

### v0.3.13

- Added `docs/132-clean-working-snapshot-runner-validation-v0.md`.
- Linked the clean working snapshot runner validation record from `docs/README.md`.
- Updated `roadmap/ROADMAP.md` with v0.3.13 clean working snapshot runner validation and the next direction toward release-snapshot validation, runner-backed public testing docs, and later actual backbone deployability work within late v0.3.x.
- Created a clean working snapshot archive at `clean-snapshot-validation/v0.3.13/carbonstack-v0.3.13-clean-working-snapshot.tgz`.
- Snapshot excluded `.git`, generated Rust/OpenMLS state, Go caches/temp roots, DBs, exes, and OS junk.
- Validated the clean snapshot on `[WSL-DEBIAN]` first from `~/carbonstack-v0.3.13-clean`.
- WSL preflight artifact scan was clean.
- WSL runner `go test ./...`, `doctor`, `core`, and `full` passed.
- WSL `full` correctly remained an alias for `core`; because `full` was run after `core`, its pre-test scan observed already-generated OpenMLS sidecar state/target roots from the prior `core` run.
- WSL post-run external artifact scan showed expected OpenMLS sidecar generated roots plus nested `provider-storage.json` and `signer.json` under `.carbonstack-openmls-sidecar-state`.
- Validated the clean snapshot on `[WIN-PWRSHL]` second from `clean-snapshot-validation/v0.3.13/windows-clean`.
- Windows preflight artifact scan was clean.
- Windows runner `go test ./...`, `doctor`, `core`, and `full` passed.
- Windows `sqlite3` was missing from PATH and remained a warning, not a failure.
- Windows `rustc 1.95.0` triggered the runner warning that it is below the WSL known-good Rust floor of `1.96.0`, but the clean snapshot still passed.
- Windows post-run external artifact scan showed expected OpenMLS sidecar generated roots; many `.exe` files appeared under the sidecar `target/` tree because Cargo emits Windows build helpers and binaries.
- Refined artifact interpretation: `provider-storage.json`, `signer.json`, and Windows `.exe` build artifacts are acceptable only when generated after tests and scoped under known OpenMLS sidecar generated roots, not when present before tests or scattered elsewhere.
- `carbonstack` advanced to `585d899 docs: record clean snapshot runner validation`.

### v0.3.14

- Added `docs/133-release-snapshot-profile-design-v0.md`.
- Linked the release-snapshot profile design record from `docs/README.md`.
- Updated `roadmap/ROADMAP.md` with v0.3.14 release-snapshot profile design and the intended late-v0.3.x sequence.
- Kept v0.3.14 design-only: no runner code implementation, no public release, no release upload, and no deployability work.
- Designed future `go run . --profile release-snapshot --root <release-package-root>` behavior.
- Chose a formal release package root as the design target while keeping compatibility with sibling repo layout where practical.
- Defined required repo folders: `carbonstack`, `carbonstack-comms`, and `carbonstack-cypher`; `carbonstack-os` remains optional for current v0.3.x backbone releases.
- Defined required repo files and preserved `carbonstack-comms/go.sum` as optional/include-if-present.
- Defined preferred release metadata files under a future `release/` folder: `manifest.json`, `checksums.txt`, `validation-freeze.md`, `testing-runbook.md`, `release-notes.md`, and `LICENSE`.
- Defined stricter pre-test artifact policy for release-snapshot validation: forbidden generated/private/build artifacts should fail the profile before tests unless explicitly allowlisted.
- Defined post-test artifact policy: OpenMLS sidecar `target/` and `.carbonstack-openmls-sidecar-state/` are expected generated roots; `provider-storage.json`, `signer.json`, and Windows Cargo `.exe` files are acceptable only under known generated roots after tests.
- Defined validation flow: package root detection, required folder/file checks, release metadata checks, doctor/toolchain checks, strict pre-test artifact scan, call `core`, post-test artifact scan, summarize result and nonclaims.
- Preserved `full` as a `core` alias until real release/deployability surfaces exist.
- Recorded the nested Go module validation blunder: `go test .\tools\carbonstack-validate\...` from the non-Go-module `carbonstack` root fails because `tools/carbonstack-validate` is a standalone nested Go module; the correct command is to `cd carbonstack/tools/carbonstack-validate` and run `go test ./...`.
- Added late-v0.3.x continuity: after runner-backed validation/release-snapshot/public testing release readiness is stable, v0.3.x should include actual backbone deployability work; after local deployability is ironed out, later planning should cover deployable server-system mechanics, IRC-like CarbonStack standard server systems, and server/admin workflows under the hostile-endpoint assumption.
- `carbonstack` advanced to `52785a2 docs: design release snapshot validation profile`.

### v0.3.15

- Implemented the Go runner `release-snapshot` profile in `carbonstack/tools/carbonstack-validate`.
- Added `release_snapshot.go`.
- Patched `main.go` so `--profile release-snapshot` dispatches to `r.ReleaseSnapshot()`.
- Updated the runner README with `release-snapshot` usage.
- Added `docs/134-release-snapshot-profile-implementation-v0.md`.
- Linked the implementation record from `docs/README.md`.
- Updated `roadmap/ROADMAP.md` with v0.3.15 release-snapshot profile implementation and next sequencing.
- Implemented package checks for a formal release-like root containing `carbonstack/`, `carbonstack-comms/`, `carbonstack-cypher/`, and `release/`.
- Implemented required repo folder/file checks and minimal release metadata checks.
- Preserved `carbonstack-comms/go.sum` as optional/include-if-present.
- Implemented strict pre-test artifact scan for `release-snapshot`. Unlike `core`, `release-snapshot` fails if generated/private/build artifacts are present before tests.
- Preserved layered validation: `release-snapshot` runs package checks and strict pre-test artifact scan, then calls `core`.
- Built a local formal release-like package at `release-snapshot-validation/v0.3.15`.
- First WSL run showed the profile was missing from the package because the tarball contained a stale runner or `main.go` had not been patched to dispatch the new profile.
- Live-root sanity probe then confirmed the profile was recognized and correctly failed against the live umbrella root because the live root does not contain the required `release/` metadata directory.
- A later package attempt failed strict pre-test artifact scanning because the packaging source root had been validated before tar creation. Running `release-snapshot` on the packaging source root generated OpenMLS sidecar `target/` and `.carbonstack-openmls-sidecar-state/`, contaminating the tarball. Correct rule: validate throwaway extractions, not the packaging source root intended for tar/publishing.
- Rebuilt the package root cleanly, sanity-checked only non-generating facts before tar, then created a fresh tarball.
- WSL Debian validation passed from a fresh extraction: package checks passed, strict pre-test scan passed, `core` ran, targeted OpenMLS real-Cypher lifecycle passed, Comms tests passed, Cypher tests passed, and post-test artifacts stayed under expected OpenMLS sidecar generated roots.
- Windows validation also passed from a fresh extraction before commit.
- `carbonstack` advanced to `6cdbe5f tools: implement release snapshot validation profile`.

### v0.3.16

- Hardened the Go runner `release-snapshot` profile output.
- Added explicit run-order warnings to `release_snapshot.go`:
  - validate only fresh extracted/staged package roots;
  - do not validate the package source root that will later be archived/published;
  - a successful validation generates artifacts, so rerun from a fresh extraction.
- Updated the runner README with the `release-snapshot` run-order warning.
- Added `docs/135-release-snapshot-validation-hardening-v0.md`.
- Linked the v0.3.16 release-snapshot validation hardening record from `docs/README.md`.
- Updated `roadmap/ROADMAP.md` with the v0.3.16 validation-hardening rule and next direction.
- Built a staged release-like package under `release-snapshot-validation/v0.3.16`.
- Package source root was treated as non-generating: only runner code presence, `release_snapshot.go`, `go test ./...`, and forbidden artifact scans were allowed before archiving.
- Created archive `release-snapshot-validation/v0.3.16/carbonstack-v0.3.16-staged-release-like-package.tgz`.
- WSL Debian validation passed from a fresh extraction at `~/carbonstack-v0.3.16-release-snapshot`.
- WSL Debian runner output confirmed the new run-order warnings were printed, layout/metadata checks passed, strict pre-test artifact scan passed, `core` ran, OpenMLS real-Cypher lifecycle passed, Comms package tests passed, Cypher package tests passed, and post-test artifacts stayed under expected OpenMLS sidecar generated roots.
- Windows validation passed from a fresh extraction at `release-snapshot-validation/v0.3.16/windows-release-snapshot`.
- Windows runner output confirmed the new run-order warnings were printed, layout/metadata checks passed, strict pre-test artifact scan passed, `core` ran, OpenMLS real-Cypher lifecycle passed, Comms package tests passed, Cypher package tests passed, and post-test artifacts stayed under expected OpenMLS sidecar generated roots.
- Windows still warned that `sqlite3` was unavailable in PATH and that Rust `1.95.0` is below the conservative WSL known-good floor of `1.96.0`; validation still passed, so these remain warnings, not blockers.
- Preserved checksum policy: the v0.3.16 staged package still uses placeholder/policy-refinement checksum text; real checksum generation/verification is future release-staging work, not yet implemented.
- Preserved the regex patch blunder: exact PowerShell string replacement failed against `release_snapshot.go` because `gofmt` indentation did not match `$Old`; use regex or inspect with `cat`/`Select-String` when patching formatted Go source.
- `carbonstack` advanced to `d159042 tools: harden release snapshot validation profile`.

### v0.3.17

- Added `docs/136-runner-backed-public-testing-path-v0.md`.
- Linked the runner-backed public testing path draft from `docs/README.md`.
- Updated `roadmap/ROADMAP.md` with v0.3.17 runner-backed public testing docs and the next direction toward real manifest/checksum semantics, release staging rehearsal, uploaded/downloaded asset verification, and only then an optional public runner-backed testing release.
- Kept v0.3.17 documentation-only: no new public release, no tag movement, no uploaded/downloaded release verification, no deployability work, no public README switch that would imply v0.3.4 is runner-backed.
- Documented that `v0.3.4` remains the current recommended public testing release and Windows 11 / PowerShell remains the current public testing boundary for that release.
- Documented the future runner-backed testing path using the Go runner under `carbonstack/tools/carbonstack-validate`.
- Reconfirmed validation authority direction: WSL Debian first, Windows second before breakpoints/releases, Go runner as preferred authority, PowerShell as legacy/current public v0.3.4 path until a runner-backed public release exists.
- Documented implemented runner profiles: `doctor`, `core`, `full`, and `release-snapshot`; `full` still aliases `core`.
- Documented expected future package shape: `carbonstack/`, `carbonstack-comms/`, `carbonstack-cypher/`, and `release/` metadata.
- Documented the `release-snapshot` fresh extraction rule and the package-source-root contamination hazard.
- Documented forbidden pre-test artifacts and expected post-test OpenMLS sidecar generated roots.
- Documented WSL Debian known-good path/toolchain context and Windows known-good path/toolchain context.
- Preserved Rust policy: Debian apt Rust 1.85.0 is known-bad for the OpenMLS 0.8.1 path; WSL Debian Rust/Cargo 1.96.0 and Windows Rust/Cargo 1.95.0 are known-good in current tests; runner Rust warning remains conservative and tests remain the authority.
- Documented what passing `release-snapshot` validates and what it explicitly does not validate.
- `carbonstack` advanced to `a8ed825 docs: describe runner-backed public testing path`.


### v0.3.18

- Added `docs/137-runner-backed-release-staging-rehearsal-v0.md`.
- Implemented Go runner checksum helper profile work around `write-checksums` and `verify-checksums`.
- Added `carbonstack/tools/carbonstack-validate/checksums.go`.
- Patched `release-snapshot` so it verifies real package checksums before calling `core`.
- Rehearsed a local runner-backed release-candidate package with `carbonstack/`, `carbonstack-comms/`, `carbonstack-cypher/`, and `release/`.
- Recorded `carbonstack-os` as related but not included in the current runnable v0.3.x package.
- Fixed checksum parser blunder caused by paths with spaces: parse first 64 characters as SHA-256 and the remainder as the path, rather than using `strings.Fields`.
- Fixed carrier archive contamination lesson: keep `.tgz`/transport archives outside the package root using a `downloads/` and `package/` split.
- Validated the local release-candidate rehearsal package from fresh WSL Debian and Windows extractions.
- Did not publish a public release and did not touch the Gitea release surface.

### v0.3.19

- Added `docs/138-local-downloaded-asset-simulation-v0.md`.
- Linked the local downloaded-asset simulation record from `docs/README.md`.
- Updated `roadmap/ROADMAP.md` with the v0.3.19 local simulation direction and short-path Windows validation caveat.
- Kept v0.3.19 local-only: no Gitea upload, no tag, no public release, and no deployability work.
- Simulated a public-user asset flow with `published-assets/`, `downloads/`, and a fresh extracted `package/` root.
- Preserved v0.3.4 as the current recommended public testing release until a future runner-backed public release is fully staged, verified, documented, and intentionally published.
- WSL Debian local downloaded-asset simulation passed from a fresh package extraction.
- Windows long nested extraction initially failed during Rust/MSVC linking with `LNK1104` while creating build-script `.exe` outputs under the OpenMLS sidecar `target/` tree.
- Windows short-path validation then passed from `C:\cs-v0319-win2\package` with `CARGO_BUILD_JOBS=1`.
- Recorded the interpretation that Windows runner-backed validation is path/toolchain/file-lock sensitive and should use a short extraction root until better characterized.
- Deferred a separate `release-verify` runner profile until after v0.3.21+ maturity, when deployable backbone and IRC-style join/admin mechanics justify a dedicated public release-verification surface.
- `carbonstack` advanced to `0f448f2 docs: record local downloaded asset simulation`.

---

### v0.3.20

- Added `docs/139-runner-backed-testing-release-cleanup-v0.md`.
- Updated public README/docs/roadmap/front-door wording for the runner-backed release transition.
- Cleaned stale public-surface wording without rewriting historical numbered docs.
- Built the final v0.3.20 runner-backed release package under `release-staging/v0.3.20`.
- Generated real package-internal checksums with `write-checksums`.
- Created release assets and outer asset checksums.
- Revalidated the final package from fresh WSL Debian extraction.
- Revalidated the final package from Windows short extraction root `C:\cs-v0320\package` with `CARGO_BUILD_JOBS=1`.
- Published `v0.3.20` as a Gitea pre-release: **CarbonStack v0.3.20 Runner-Backed Testing Release**.
- Attached the intended multi-repo runner-backed package and metadata assets.
- Preserved the warning that Gitea default Source Code ZIP/TAR.GZ archives are only auto-generated archives of `carbonstack`, not the intended multi-repo validation package.
- Marked Debian 13 / WSL2 Debian as the primary validated target.
- Marked Windows 11 short-path validation as secondary and the final explicit Windows dev/test validation for this phase.
- Preserved `v0.3.4` as the older Windows 11 / PowerShell legacy testbed.
- Did not claim deployability, production security, production E2EE, hostile-server safety, metadata privacy, audit, or certification.

## 4. New at v0.3.15

### 4.1 Release-snapshot profile implemented

v0.3.15 implements the profile designed in v0.3.14:

```text
go run . --profile release-snapshot --root <release-package-root>
```

The implementation lives under:

```text
carbonstack/tools/carbonstack-validate
```

New/changed runner surfaces:

```text
carbonstack/tools/carbonstack-validate/main.go
carbonstack/tools/carbonstack-validate/release_snapshot.go
carbonstack/tools/carbonstack-validate/README.md
```

The profile is now recognized by the runner and dispatches to `r.ReleaseSnapshot()`.

### 4.2 Implemented release package checks

The implemented profile expects a formal release-like package root containing:

```text
carbonstack/
carbonstack-comms/
carbonstack-cypher/
release/
```

Current required repo checks include the core docs/source/test surfaces needed to validate the v0.3.x experimental backbone. `carbonstack-os` remains optional for current v0.3.x release-snapshot validation because it is not part of the current runnable backbone proof.

`carbonstack-comms/go.sum` remains optional/include-if-present.

### 4.3 Implemented release metadata checks

Current minimal release metadata checks include:

```text
release/manifest.json
release/checksums.txt
release/validation-freeze.md OR release/testing-runbook.md
```

Optional release metadata includes:

```text
release/release-notes.md
release/LICENSE
```

The v0.3.15 local validation package used placeholder checksum text for implementation validation. Future public release packages must include real checksum coverage.

### 4.4 Strict pre-test artifact scan implemented

Unlike `core`, `release-snapshot` treats forbidden pre-test generated/private/build artifacts as a validation failure.

Forbidden before tests include:

```text
.git/
target/
.carbonstack-openmls-sidecar-state/
.go-cache/
.go-tmp/
provider-storage.json
signer.json
*.db
*.db-shm
*.db-wal
*.exe
*.test.exe
Thumbs.db
.DS_Store
```

This is intentionally stricter than the non-destructive reporting behavior used by `core`.

### 4.5 Layered validation preserved

The implemented flow is:

```text
release-snapshot package checks
strict pre-test artifact scan
core validation
post-test artifact scan
```

This preserves the design model:

```text
release-snapshot checks -> core
```

and leaves:

```text
full -> core alias
```

unchanged.

### 4.6 Local formal release-like package validation

A local formal release-like package was built under:

```text
C:\▓▓\repos\carbonstack_umbrella\release-snapshot-validation\v0.3.15
```

The package root shape was:

```text
package-root/
  carbonstack/
  carbonstack-comms/
  carbonstack-cypher/
  release/
```

Archive:

```text
release-snapshot-validation/v0.3.15/carbonstack-v0.3.15-release-snapshot-package.tgz
```

WSL extraction/validation root:

```text
~/carbonstack-v0.3.15-release-snapshot
```

Windows extraction/validation root:

```text
C:\▓▓\repos\carbonstack_umbrella\release-snapshot-validation\v0.3.15\windows-release-snapshot
```

### 4.7 WSL Debian release-snapshot validation passed

WSL Debian validation passed from a fresh extraction.

Validated:

```text
go test ./...
go run . --profile release-snapshot --root "$HOME/carbonstack-v0.3.15-release-snapshot"
```

Observed WSL result:

```text
release layout checks: PASS
release metadata checks: PASS
strict pre-test artifact scan: PASS
OpenMLS real-Cypher lifecycle: PASS
carbonstack-comms package tests: PASS
carbonstack-cypher package tests: PASS
post-test artifacts: expected OpenMLS sidecar generated roots only
```

Observed WSL toolchain in the runner output:

```text
go version go1.24.4 linux/amd64
rustc 1.96.0
cargo 1.96.0
sqlite3 3.46.1
```

### 4.8 Windows release-snapshot validation passed

Windows validation passed from a fresh extraction after WSL validation.

Validated:

```powershell
go test ./...
go run . --profile release-snapshot --root C:\▓▓\repos\carbonstack_umbrella\release-snapshot-validation\v0.3.15\windows-release-snapshot
```

Windows remains the second confirmation environment before breakpoints/releases. Windows `sqlite3` may still be missing from PATH and remain a warning, not a blocker. Windows Rust `1.95.0` is known-good for the current tested path despite the conservative Rust floor warning.

### 4.9 Blunders / repair lessons from v0.3.15

#### Stale package runner blunder

The first WSL package validation reported:

```text
unknown profile "release-snapshot"; expected doctor, core, or full
```

This meant the extracted package contained a stale runner or `main.go` had not been patched to dispatch the new profile. The repair was to inspect `main.go`, confirm `release_snapshot.go` existed, patch the switch robustly, rebuild the package root, and recreate the tarball.

#### Live-root failure was expected

Running:

```powershell
go run . --profile release-snapshot --root C:\▓▓\repos\carbonstack_umbrella
```

against the live umbrella root correctly failed because the live root does not contain the formal `release/` metadata directory. This confirmed the profile was recognized and correctly differentiates a live dev root from a formal release package root.

#### Package-root contamination blunder

Running `release-snapshot` on the package root before tar creation generated OpenMLS sidecar artifacts inside the package root:

```text
carbonstack-comms/internal/protocol/mls/openmls-sidecar/target
carbonstack-comms/internal/protocol/mls/openmls-sidecar/.carbonstack-openmls-sidecar-state
```

Those artifacts were then included in the tarball, and WSL correctly failed strict pre-test artifact scanning.

Correct rule:

```text
Never validate the package root intended for tar/publishing.
Validate a throwaway extraction.
Only run non-generating sanity checks on the packaging source root before tar creation.
```

### 4.10 Interpretation

v0.3.15 implements and validates the release-snapshot profile against a local formal release-like package. v0.3.16 hardens the release-snapshot profile output/run-order documentation and validates a staged release-like package from fresh extractions on WSL Debian and Windows.

It does not publish a public runner-backed release.

It does not update public testing docs to the runner path.

It does not prove public Linux support, Debian deployability, systemd readiness, cloudflared readiness, real homelab readiness, production security, hostile-server safety, metadata privacy, audit, or certification.

---


## 4B. New at v0.3.16

### 4B.1 Release-snapshot output hardened

v0.3.16 adds explicit run-order warnings to the `release-snapshot` profile output. The runner now prints:

```text
run-order warning: validate only fresh extracted/staged package roots
do not validate the package source root that will later be archived/published
a successful validation generates artifacts, so rerun from a fresh extraction
```

This makes the v0.3.15 package-root contamination lesson visible directly in the validation output.

### 4B.2 Runner README updated

The runner README now documents the `release-snapshot` run-order warning:

```text
create clean package source root
archive it without running release-snapshot inside it
extract archive into a throwaway validation root
run release-snapshot from the throwaway extraction
discard or preserve the throwaway extraction only as validation evidence
```

Do not run `release-snapshot` twice in the same extraction unless the second run is intentionally expected to fail strict pre-test artifact scanning.

### 4B.3 Staged release-like package validation recorded

v0.3.16 adds:

```text
docs/135-release-snapshot-validation-hardening-v0.md
```

This records staged release-like package validation from fresh extractions on WSL Debian and Windows.

### 4B.4 v0.3.16 staged package paths

[WIN-PWRSHL] package staging root:

```text
C:\▓▓\repos\carbonstack_umbrella\release-snapshot-validation\v0.3.16
```

Package source root:

```text
C:\▓▓\repos\carbonstack_umbrella\release-snapshot-validation\v0.3.16\package-source-root
```

Archive:

```text
C:\▓▓\repos\carbonstack_umbrella\release-snapshot-validation\v0.3.16\carbonstack-v0.3.16-staged-release-like-package.tgz
```

[WSL-DEBIAN] validation root:

```text
~/carbonstack-v0.3.16-release-snapshot
```

[WIN-PWRSHL] validation root:

```text
C:\▓▓\repos\carbonstack_umbrella\release-snapshot-validation\v0.3.16\windows-release-snapshot
```

### 4B.5 WSL Debian v0.3.16 validation passed

WSL Debian validation passed from a fresh extraction.

Validated:

```bash
. "$HOME/.cargo/env"
cd "$HOME/carbonstack-v0.3.16-release-snapshot/carbonstack/tools/carbonstack-validate"
go test ./...
go run . --profile release-snapshot --root "$HOME/carbonstack-v0.3.16-release-snapshot"
```

Observed WSL result:

```text
release-snapshot run-order warnings printed
release layout checks: PASS
release metadata checks: PASS
strict pre-test artifact scan: PASS
OpenMLS real-Cypher lifecycle: PASS
carbonstack-comms package tests: PASS
carbonstack-cypher package tests: PASS
post-test artifacts: expected OpenMLS sidecar generated roots only
VALIDATION PASSED
```

Observed WSL toolchain in runner output:

```text
go version go1.24.4 linux/amd64
rustc 1.96.0
cargo 1.96.0
sqlite3 3.46.1
```

### 4B.6 Windows v0.3.16 validation passed

Windows validation passed from a fresh extraction after WSL validation.

Validated:

```powershell
cd C:\▓▓\repos\carbonstack_umbrella\release-snapshot-validation\v0.3.16\windows-release-snapshot\carbonstack\tools\carbonstack-validate

go test ./...
go run . --profile release-snapshot --root C:\▓▓\repos\carbonstack_umbrella\release-snapshot-validation\v0.3.16\windows-release-snapshot
```

Observed Windows result:

```text
release-snapshot run-order warnings printed
release layout checks: PASS
release metadata checks: PASS
strict pre-test artifact scan: PASS
OpenMLS real-Cypher lifecycle: PASS
carbonstack-comms package tests: PASS
carbonstack-cypher package tests: PASS
post-test artifacts: expected OpenMLS sidecar generated roots only
VALIDATION PASSED
```

Windows still warned:

```text
sqlite3 path: WARN not found in PATH
rustc appears older than known-good floor 1.96
```

These remain warnings, not blockers, because current Windows validation passed with Rust/Cargo `1.95.0`.

### 4B.7 Checksum policy preserved

The v0.3.16 staged package still uses placeholder checksum text for local profile hardening. This is acceptable for v0.3.16 validation-hardening work. It is **not** acceptable for a public runner-backed release.

Real checksum generation/verification remains future release-staging work, likely around the runner-backed release staging / uploaded-download verification rungs.

### 4B.8 PowerShell exact replacement blunder

The first attempt to add run-order warnings used `$Text.Contains($Old)` and `.Replace($Old, $New)`. It failed because the exact string did not match `gofmt` indentation in `release_snapshot.go`.

Repair:

```text
inspect with cat / Select-String
use regex replacement for formatted Go source
run gofmt afterward
confirm with Select-String
```

### 4B.9 Interpretation

v0.3.16 validates the hardened `release-snapshot` profile against a staged release-like package from fresh extractions on WSL Debian and Windows.

It does not publish a runner-backed public release.

It does not verify uploaded/downloaded release assets.

It does not implement real checksum generation/verification.

It does not prove deployability, production security, hostile-server safety, metadata privacy, audit, or certification.



## 4C. New at v0.3.17

### 4C.1 Runner-backed public testing path documented

v0.3.17 adds:

```text
carbonstack/docs/136-runner-backed-public-testing-path-v0.md
```

This document prepares a future runner-backed public testing path without superseding the current public v0.3.4 Windows/PowerShell release.

The document records:

```text
current public release: v0.3.4
current v0.3.4 public testing boundary: Windows 11 / PowerShell
future validation authority: Go runner
future validation order: WSL Debian first, Windows second
future release package shape: carbonstack + carbonstack-comms + carbonstack-cypher + release metadata
```

### 4C.2 v0.3.4 explicitly preserved

v0.3.17 does not claim v0.3.4 is runner-backed.

It preserves the distinction:

```text
v0.3.4:
  current recommended public testing release
  Windows 11 / PowerShell public testing path

v0.3.17:
  mainline docs checkpoint
  future runner-backed public testing path draft
  no new public release
```

### 4C.3 Future release-snapshot tester path recorded

The documented future command shape remains:

```text
cd carbonstack/tools/carbonstack-validate
go test ./...
go run . --profile release-snapshot --root <fresh-extracted-package-root>
```

The document records that `release-snapshot` must be run from a fresh extracted/throwaway root and must not be run against the package source root that will later be archived or published.

### 4C.4 Toolchain and artifact policy recorded for testers

The v0.3.17 document records known-good/current validation context:

```text
WSL Debian:
  Debian GNU/Linux 13 trixie under WSL2
  linux/amd64
  Go 1.24.4
  rustup stable Rust/Cargo 1.96.0
  sqlite3 3.46.1

Windows:
  Windows 11
  Go 1.26.3 windows/amd64
  Rust/Cargo 1.95.0
  sqlite3 may be unavailable in PATH and remains a warning for current tests
```

It also records that Debian apt Rust 1.85.0 is known-bad for the current OpenMLS 0.8.1 path, while WSL Rust 1.96.0 and Windows Rust 1.95.0 are known-good in current tests.

### 4C.5 Scope and nonclaims documented

The future runner-backed public testing path document records that passing `release-snapshot` validates only package layout, release metadata presence, clean pre-test artifact state, current core backbone validation, and post-test artifact scoping.

It explicitly does not validate production readiness, production E2EE, hostile-server safety, metadata privacy, Android readiness, secure vault/storage, public Linux release support, Debian deployability, systemd readiness, cloudflared readiness, public ingress safety, real homelab validation, external audit, or certification.

### 4C.6 Interpretation

v0.3.17 prepares the documentation surface needed before a future runner-backed public testing release.

It does not publish that release.

It does not verify uploaded/downloaded runner-backed assets.

It does not implement real checksum generation or verification.

It does not prove deployability.


## 4D. New at v0.3.18

### 4D.1 Go checksum helper profiles implemented

v0.3.18 adds real checksum support to the Go validation runner.

New runner file:

```text
carbonstack/tools/carbonstack-validate/checksums.go
```

New profiles:

```text
go run . --profile write-checksums --root <package-root>
go run . --profile verify-checksums --root <package-root>
```

`write-checksums` writes:

```text
release/checksums.txt
```

`verify-checksums` verifies the checksum file against the current package root.

The checksum helper excludes generated/private/build artifacts, carrier archives, and `release/checksums.txt` itself.

### 4D.2 release-snapshot now verifies real checksums

`release-snapshot` now runs checksum verification before calling `core`.

The updated flow is:

```text
release-snapshot package/layout checks
strict pre-test artifact scan
release checksum verification
core validation
post-test artifact scan
```

This moves the runner-backed release path from placeholder checksum semantics toward real release-candidate rehearsal semantics.

### 4D.3 v0.3.18 local release-candidate rehearsal package

[WIN-PWRSHL] staging root:

```text
C:\▓▓\repos\carbonstack_umbrella\release-staging\v0.3.18
```

[WIN-PWRSHL] package source root:

```text
C:\▓▓\repos\carbonstack_umbrella\release-staging\v0.3.18\local-release-candidate-rehearsal
```

Carrier archive:

```text
C:\▓▓\repos\carbonstack_umbrella\release-staging\v0.3.18\carbonstack-v0.3.18-local-release-candidate-rehearsal.tgz
```

Package shape:

```text
<package-root>/
  carbonstack/
  carbonstack-comms/
  carbonstack-cypher/
  release/
    manifest.json
    checksums.txt
    validation-freeze.md
    testing-runbook.md
    release-notes.md
    LICENSE
```

### 4D.4 carbonstack-os status

`carbonstack-os` was recorded in the manifest as related but not included.

Reason:

```text
carbonstack-os is part of the long-term appliance stack direction,
but it is not part of the current runnable v0.3.x backbone validation package.
```

This keeps the full-stack direction visible without pretending that the OS repo is part of the current release-candidate validation surface.

### 4D.5 WSL Debian validation passed

[WSL-DEBIAN] validation root:

```text
~/carbonstack-v0.3.18-release-rehearsal
```

The validation used a safer split:

```text
~/carbonstack-v0.3.18-release-rehearsal/downloads
~/carbonstack-v0.3.18-release-rehearsal/package
```

The carrier archive stayed under `downloads/`.

The package root passed validation under `package/`.

Validated from:

```text
~/carbonstack-v0.3.18-release-rehearsal/package/carbonstack/tools/carbonstack-validate
```

Commands:

```bash
go test ./...
go run . --profile verify-checksums --root "$HOME/carbonstack-v0.3.18-release-rehearsal/package"
go run . --profile release-snapshot --root "$HOME/carbonstack-v0.3.18-release-rehearsal/package"
```

Observed WSL result:

```text
package-root top-level file scan: clean
verify-checksums: PASS
release layout checks: PASS
release metadata checks: PASS
strict pre-test artifact scan: PASS
release checksum verification: PASS
OpenMLS real-Cypher lifecycle: PASS
carbonstack-comms package tests: PASS
carbonstack-cypher package tests: PASS
post-test artifacts: expected OpenMLS sidecar generated roots only
VALIDATION PASSED
```

### 4D.6 Windows validation passed

[WIN-PWRSHL] validation root:

```text
C:\▓▓\repos\carbonstack_umbrella\release-staging\v0.3.18\windows-release-rehearsal
```

The validation used the same split:

```text
windows-release-rehearsal\downloads
windows-release-rehearsal\package
```

Validated from:

```text
C:\▓▓\repos\carbonstack_umbrella\release-staging\v0.3.18\windows-release-rehearsal\package\carbonstack\tools\carbonstack-validate
```

Commands:

```powershell
go test ./...
go run . --profile verify-checksums --root C:\▓▓\repos\carbonstack_umbrella\release-staging\v0.3.18\windows-release-rehearsal\package
go run . --profile release-snapshot --root C:\▓▓\repos\carbonstack_umbrella\release-staging\v0.3.18\windows-release-rehearsal\package
```

Observed Windows result:

```text
package-root top-level file scan: clean
verify-checksums: PASS
release layout checks: PASS
release metadata checks: PASS
strict pre-test artifact scan: PASS
release checksum verification: PASS
OpenMLS real-Cypher lifecycle: PASS
carbonstack-comms package tests: PASS
carbonstack-cypher package tests: PASS
post-test artifacts: expected OpenMLS sidecar generated roots only
VALIDATION PASSED
```

### 4D.7 Checksum parsing blunder

The first checksum verification failed because paths with spaces were parsed using whitespace splitting.

Bad parser assumption:

```text
strings.Fields(line)
```

This failed on paths like sanitized LogDoc filenames or docs with spaces.

Repair:

```text
first 64 characters = SHA-256 hash
remaining trimmed text = relative path
```

This allows checksum entries with spaces in paths.

### 4D.8 Carrier archive blunder

A WSL validation attempt placed the carrier archive directly inside the package root:

```text
~/carbonstack-v0.3.18-release-rehearsal/
  carbonstack-v0.3.18-local-release-candidate-rehearsal.tgz
  carbonstack/
  carbonstack-comms/
  carbonstack-cypher/
  release/
```

`verify-checksums` correctly saw the `.tgz` as an unexpected package-root file and failed because the archive was not listed in `release/checksums.txt`.

Repair:

```text
downloads/ contains the carrier archive
package/ contains only extracted package contents
--root points at package/
```

Defensive archive suffix exclusions were also added:

```text
.tgz
.tar.gz
.tar
.zip
```

The workflow rule remains stronger than the code-level exclusion: carrier archives should stay outside the package root.

### 4D.9 Interpretation

v0.3.18 rehearses a local runner-backed release-candidate package with real Go-generated checksum coverage.

Allowed claim:

```text
A local runner-backed release-candidate rehearsal package with real Go-generated checksums passed verification and release-snapshot validation from fresh extractions on WSL Debian and Windows.
```

This still does not prove:

```text
public release readiness
uploaded/downloaded asset integrity
public Linux release support
Debian deployability
systemd readiness
cloudflared readiness
real homelab validation
production readiness
production E2EE
hostile-server safety
metadata privacy
external audit
certification
```

### 4D.10 Next direction

Next safest work:

```text
v0.3.19 uploaded/downloaded release asset verification workflow
```

That should verify future release assets as a public user would:

```text
download uploaded assets
verify checksums
extract fresh package root
run verify-checksums
run release-snapshot
validate WSL Debian first
validate Windows second
record Gitea/release-surface quirks
```

Do not begin actual deployability until runner-backed release readiness is clean.



## 4E. New at v0.3.19

### 4E.1 Local downloaded-asset simulation recorded

v0.3.19 adds:

```text
carbonstack/docs/138-local-downloaded-asset-simulation-v0.md
```

This records a local public-user-style asset flow before touching the Gitea release surface. The simulated shape is:

```text
release-verification/v0.3.19-local-download-simulation/
  published-assets/
  downloads/
  extracted/package/
```

Interpretation:

```text
published-assets/ simulates the release-page asset store.
downloads/ simulates what a tester receives.
extracted/package/ is the fresh package root used for validation.
```

The carrier archive stays outside the package root. The package root must contain only:

```text
carbonstack/
carbonstack-comms/
carbonstack-cypher/
release/
```

### 4E.2 No Gitea release surface yet

v0.3.19 intentionally avoids the Gitea release surface. It is a local simulation / planning checkpoint, not a public release.

Reasoning:

```text
Hotfix releases like v0.3.4 are inevitable sometimes,
but the project should avoid hotfix flooding by fully proving packaging,
platform tests, docs, and release text before publication.
```

The eventual public runner-backed release may be `v0.3.19` or `v0.3.20` depending on which boundary is cleaner. Use the cleaner release boundary, not the most compressed one.

### 4E.3 WSL Debian local simulation passed

[WSL-DEBIAN] local downloaded-asset simulation used a fresh package extraction and passed:

```text
verify-checksums: PASS
release layout checks: PASS
release metadata checks: PASS
strict pre-test artifact scan: PASS
release checksum verification: PASS
OpenMLS real-Cypher lifecycle: PASS
carbonstack-comms package tests: PASS
carbonstack-cypher package tests: PASS
post-test artifacts: expected OpenMLS sidecar generated roots only
VALIDATION PASSED
```

### 4E.4 Windows long-path failure

Initial Windows validation from the long nested release-verification extraction path reached checksum verification and release-snapshot pre-core checks, but failed inside `core` during Rust/OpenMLS sidecar compilation.

Failure pattern:

```text
linking with link.exe failed: exit code: 1104
LINK : fatal error LNK1104: cannot open file ... build_script_build-<hash>.exe
```

This occurred across many unrelated Rust crates/build scripts, including examples such as `generic-array`, `quote`, `serde`, `serde_core`, `proc-macro2`, `hax-lib`, `zerocopy`, `crossbeam-utils`, `curve25519-dalek`, and `libc`.

Interpretation:

```text
This was not a deterministic CarbonStack protocol failure.
It was a Windows Rust/MSVC linker/path/file-lock/environment failure while creating build-script executables.
```

Likely contributing factors remain:

```text
long nested extraction path;
Windows file locking / AV / indexing timing;
MSVC linker behavior around newly generated Rust build-script executables;
parallel Cargo build sensitivity.
```

### 4E.5 Windows short-path validation passed

Windows validation then passed from a short extraction root:

```text
C:\cs-v0319-win2\package
```

with:

```powershell
$env:CARGO_BUILD_JOBS = "1"
```

Observed Windows short-path result:

```text
verify-checksums: PASS for 271 files
release layout checks: PASS
release metadata checks: PASS
strict pre-test artifact scan: PASS
release checksum verification: PASS
OpenMLS real-Cypher lifecycle: PASS
carbonstack-comms package tests: PASS
carbonstack-cypher package tests: PASS
post-test artifacts: expected OpenMLS sidecar generated roots only
VALIDATION PASSED
```

Windows toolchain context from the pass:

```text
go version go1.26.3 windows/amd64
rustc 1.95.0
cargo 1.95.0
sqlite3 unavailable in PATH: warning only
```

Windows still triggers the conservative Rust floor warning because the WSL known-good floor is Rust 1.96.0, but Windows Rust 1.95.0 remains known-good for the current tested path when the short-path/root and single-job Cargo mitigation are used.

### 4E.6 v0.3.19 Windows validation caveat

Future Windows known-good runner-backed validation should use a short extraction root until the Windows path/toolchain/file-lock behavior is better characterized.

Known-good current shape:

```text
C:\cs-v0319-win2\package
CARGO_BUILD_JOBS=1
```

Do not treat a pass in this shape as proof that arbitrary long nested Windows release-verification paths are stable for Rust/OpenMLS compilation.

### 4E.7 release-verify profile deferred

A separate `release-verify` runner profile is deferred. It may make sense after v0.3.21+ maturity, once deployable backbone work and IRC-style join/admin mechanics create enough mature public release surface to justify a dedicated verification profile.

For now, the correct local verification path remains:

```text
verify-checksums
release-snapshot
```

### 4E.8 Interpretation

v0.3.19 proves that a local downloaded-asset simulation can pass checksum verification and release-snapshot validation from fresh extractions on WSL Debian and Windows, with the Windows caveat that the known-good pass used a short extraction root and `CARGO_BUILD_JOBS=1`.

It does not prove Gitea upload/download integrity.

It does not publish a public runner-backed release.

It does not supersede v0.3.4.

It does not prove public Linux support, Debian deployability, systemd readiness, cloudflared readiness, production security, hostile-server safety, metadata privacy, audit, or certification.

### 4E.9 Next direction

Recommended next work:

```text
full public-release readiness docs cleanup/check;
final release page text draft;
final in-release testing runbook review;
WSL Debian known-good platform test from fresh package extraction;
Windows known-good platform test from short fresh package extraction;
then choose whether the public runner-backed release boundary is v0.3.19 or v0.3.20.
```


## 4F. New at v0.3.20

### 4F.1 Public runner-backed testing release cut

v0.3.20 cuts the first public runner-backed CarbonStack testing release.

Release page:

```text
https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.3.20
```

Release title:

```text
CarbonStack v0.3.20 Runner-Backed Testing Release
```

Release tag/commit:

```text
v0.3.20 -> 3eeb1a1e1a / 3eeb1a1 docs: clean public release surface
```

The release is marked pre-release and uses the same broad visual/release-note style as the v0.3.4 testing bug hotfix for continuity.

### 4F.2 Public-surface cleanup before release

Before the release cut, the public-facing surface was cleaned:

```text
carbonstack/README.md
carbonstack/docs/README.md
carbonstack/roadmap/ROADMAP.md
carbonstack/sanitized-project-logdoc-list/README.md
carbonstack/docs/139-runner-backed-testing-release-cleanup-v0.md
```

The cleanup intentionally did **not** rewrite historical numbered docs. Instead, public/front-door surfaces now explain:

```text
historical docs may contain stale assumptions;
current status starts from the top-level README and current release/runbook docs;
v0.3.20 is Debian/WSL-Debian-first;
v0.3.4 remains the older Windows 11 / PowerShell testbed;
Windows validation is secondary/final explicit dev-test validation for this phase;
nonclaims remain active.
```

A stale public-surface phrasing blunder was cleaned: older text still implied “latest release page” or “uploaded/downloaded verification” as the next step. That wording was updated to match the v0.3.20 release reality and local-first verification path.

### 4F.3 v0.3.20 release package and assets

Final package staging root:

```text
C:\▓▓\repos\carbonstack_umbrella\release-staging\v0.3.20
```

Final package source root:

```text
C:\▓▓\repos\carbonstack_umbrella\release-staging\v0.3.20\runner-backed-testing-release-candidate
```

Final package archive:

```text
C:\▓▓\repos\carbonstack_umbrella\release-staging\v0.3.20\carbonstack-v0.3.20-runner-backed-testing-release.tgz
```

Final asset folder:

```text
C:\▓▓\repos\carbonstack_umbrella\release-staging\v0.3.20\assets
```

Published intended assets:

```text
carbonstack-v0.3.20-runner-backed-testing-release.tgz
carbonstack-v0.3.20-release-manifest.json
carbonstack-v0.3.20-package-checksums.txt
carbonstack-v0.3.20-asset-checksums.txt
carbonstack-v0.3.20-validation-freeze.md
v0.3.20-testing-runbook.md
v0.3.20-release-notes.md
LICENSE
```

The Gitea default Source Code ZIP/TAR.GZ archives remain explicitly **not** the intended validation package because they are auto-generated archives of only the `carbonstack` repo at the tag.

### 4F.4 v0.3.20 WSL Debian validation passed

WSL Debian validation root:

```text
~/carbonstack-v0.3.20-test/package
```

Validation command shape:

```bash
. "$HOME/.cargo/env"
cd "$HOME/carbonstack-v0.3.20-test/package/carbonstack/tools/carbonstack-validate"
go test ./...
go run . --profile verify-checksums --root "$HOME/carbonstack-v0.3.20-test/package"
go run . --profile release-snapshot --root "$HOME/carbonstack-v0.3.20-test/package"
```

Observed WSL validation behavior:

```text
fresh extracted package root used;
top-level directories: carbonstack, carbonstack-comms, carbonstack-cypher, release;
top-level files: empty;
pre-run forbidden artifact scan clean;
verify-checksums passed with 274 file(s);
release-snapshot passed;
core validation passed;
post-test generated artifacts stayed under known OpenMLS sidecar generated roots.
```

### 4F.5 v0.3.20 Windows short-path validation passed

Windows validation root:

```text
C:\cs-v0320\package
```

Windows validation environment:

```powershell
$env:RUST_BACKTRACE = "full"
$env:CARGO_BUILD_JOBS = "1"
```

Validation command shape:

```powershell
cd C:\cs-v0320\package\carbonstack\tools\carbonstack-validate
go test ./...
go run . --profile verify-checksums --root "C:\cs-v0320\package"
go run . --profile release-snapshot --root "C:\cs-v0320\package"
```

Observed Windows validation behavior:

```text
fresh extracted package root used;
top-level directories: carbonstack, carbonstack-comms, carbonstack-cypher, release;
top-level files: empty;
pre-run forbidden artifact scan clean;
verify-checksums passed with 274 file(s);
release-snapshot package/layout checks passed;
release metadata checks passed;
strict pre-test artifact scan passed;
release checksum verification passed;
OpenMLS real-Cypher lifecycle passed in ~1m17s;
carbonstack-comms package tests passed in ~1m06s;
carbonstack-cypher package tests passed;
post-test artifacts were limited to known OpenMLS sidecar generated roots;
VALIDATION PASSED.
```

Windows still warned:

```text
sqlite3 path: WARN not found in PATH
rustc appears older than known-good floor 1.96
```

These remain warnings, not blockers, because the Windows short-path v0.3.20 validation passed.

### 4F.6 v0.3.20 release page formatting continuity

The v0.3.20 release description was reformatted to follow the older v0.3.4 release style for continuity:

```text
Status / Primary artifact / Release type / Platform lines;
short explanatory body;
What changed;
Validated;
Release assets;
Gitea default Source Code archive caveat;
Testing notes;
Boundary.
```

The final release page preserves hard nonclaims and platform migration language.

### 4F.7 Interpretation

v0.3.20 is the first public runner-backed testing release.

Allowed interpretation:

```text
CarbonStack v0.3.20 publishes a pre-alpha / experimental runner-backed testing package for the current Cypher + Comms OpenMLS relay backbone.
The intended multi-repo package passed checksum verification and release-snapshot validation from fresh WSL Debian and Windows short-path extractions.
Debian / WSL Debian is the primary public dev/test validation path.
Windows 11 short-path validation is the final explicit Windows dev/test validation for this phase.
```

Not allowed interpretation:

```text
production-ready
production E2EE
hostile-server safe
metadata-private
Android-ready
secure vault/storage ready
Debian deployable
systemd-ready
cloudflared-ready
public ingress safe
real homelab validated
externally audited
certified
```

### 4F.8 Next direction

Next v0.3.x work should move away from release packaging and toward actual **local-only backbone deployability** planning/implementation:

```text
local server-like layout;
Cypher bind/config/data-dir policy;
Comms-to-Cypher config;
start/stop lifecycle;
logs/errors;
runner validation for local deployment profile;
Debian/server-like layout first.
```

Still defer:

```text
cloudflared;
public ingress;
systemd service;
real homelab deployment;
shared production service paths;
MariaDB;
production claims;
separate release-verify runner profile;
IRC-style join/admin mechanics until local deployability is ironed out.
```

## 5. Current Public-Surface Status

### carbonstack

Current public role:

```text
doctrine/spec/docs/release authority
public release/front-door repo
sanitized project LogDoc archive host
v0.3.20 public runner-backed testing release authority
v0.3.4 legacy Windows 11 / PowerShell testbed authority
public release package/checksum/runbook host
future Debian-first dev/test release authority
```

Current artifact language:

```text
experimental secure-communications appliance-stack project
current validated artifact is the v0.3.x experimental Cypher + Comms OpenMLS relay backbone
```

Current main:

```text
3eeb1a1 docs: clean public release surface
```

Current preferred public testing release:

```text
v0.3.20 at 3eeb1a1e1a / 3eeb1a1
```

Legacy public Windows/PowerShell testbed release:

```text
v0.3.4 at 7c5b7e6
```

### carbonstack-comms

Current public role:

```text
OpenMLS sidecar bridge
relay helpers
real-Cypher lifecycle tests
OpenMLS backbone self-test wrapper
source-snapshot-safe artifact guard
future Comms runtime implementation surface
```

Current head:

```text
012c8bf scripts: support source snapshot artifact guard
```

Runtime `cmd/comms send` / `inbox` remains stub-era and is not the OpenMLS messenger UX.

### carbonstack-cypher

Current public role:

```text
real local relay/server
envelope API
OpenMLS content types
payload metadata
queued-only inbox
idempotent same-recipient ack
SQLite migrations
```

Current head:

```text
6f92c34 chore: fix readme formatting
```

Cypher does not handle plaintext and does not decide trust.

### carbonstack-os

Current public role:

```text
future appliance OS direction
not included in current runnable v0.3.x backbone validation package
related but explicitly not part of the v0.3.20 release package
```

Current head:

```text
1bbbe52 docs: clarify CarbonStackOS target direction
```

## 6. Critical Paths

### [WIN-PWRSHL] Local repos

```text
umbrella:             C:\▓▓\repos\carbonstack_umbrella
carbonstack:          C:\▓▓\repos\carbonstack_umbrella\carbonstack
carbonstack-comms:    C:\▓▓\repos\carbonstack_umbrella\carbonstack-comms
carbonstack-cypher:   C:\▓▓\repos\carbonstack_umbrella\carbonstack-cypher
carbonstack-os:       C:\▓▓\repos\carbonstack_umbrella\carbonstack-os
```

### [WSL-DEBIAN] WSL workspace

```text
wsl_workspace:        ~/carbonstack-wsl
carbonstack:          ~/carbonstack-wsl/carbonstack
carbonstack-comms:    ~/carbonstack-wsl/carbonstack-comms
carbonstack-cypher:   ~/carbonstack-wsl/carbonstack-cypher
openmls-sidecar:      ~/carbonstack-wsl/carbonstack-comms/internal/protocol/mls/openmls-sidecar
wsl_tar_bridge:       /mnt/c/Users/udaiv/repos/carbonstack_umbrella/carbonstack-wsl-current.tgz
rustup_env:           ~/.cargo/env
rustup_bin:           ~/.cargo/bin
```

### [GO-RUNNER] Implemented validation runner

```text
runner_dir:           carbonstack/tools/carbonstack-validate
runner_go_mod:        carbonstack/tools/carbonstack-validate/go.mod
runner_main:          carbonstack/tools/carbonstack-validate/main.go
runner_readme:        carbonstack/tools/carbonstack-validate/README.md
entrypoint_doctor:    go run . --profile doctor
entrypoint_core:      go run . --profile core
entrypoint_full:      go run . --profile full
root_override:        go run . --profile core --root <umbrella-root>
release_snapshot:     go run . --profile release-snapshot --root <release-package-root>
```

Runner-backed public testing doc:

```text
carbonstack/docs/136-runner-backed-public-testing-path-v0.md
```

### CarbonStack current docs/release surfaces

```text
carbonstack/README.md
carbonstack/roadmap/ROADMAP.md
carbonstack/docs/README.md
carbonstack/docs/v0.3.0-minor-epoch-release.md
carbonstack/docs/119-v0.3.0-release-packaging-freeze-v0.md
carbonstack/docs/120-v0.3.0-post-release-verification-v0.md
carbonstack/docs/121-security-claim-validation-policy-v0.md
carbonstack/docs/122-logdoc-case-study-sanitization-plan-v0.md
carbonstack/docs/123-v0.3.0-clean-snapshot-self-test-recon-v0.md
carbonstack/docs/124-v0.3.4-release-snapshot-self-test-fix-v0.md
carbonstack/docs/125-v0.3.4-public-release-asset-verification-v0.md
carbonstack/docs/126-v0.3.4-public-tester-runbook-v0.md
carbonstack/docs/127-windows-reliance-and-debian-homelab-recon-v0.md
carbonstack/docs/128-wsl-debian-quick-portability-setup-scout-v0.md
carbonstack/docs/129-wsl-debian-quick-portability-test-v0.md
carbonstack/docs/130-go-validation-runner-design-v0.md
carbonstack/docs/131-go-validation-runner-implementation-v0.md
carbonstack/docs/132-clean-working-snapshot-runner-validation-v0.md
carbonstack/docs/133-release-snapshot-profile-design-v0.md
carbonstack/sanitized-project-logdoc-list/
carbonstack/scripts/validate-local.ps1
carbonstack/tools/carbonstack-validate/
```

### [WIN-PWRSHL] v0.3.7 recon artifacts

```text
C:\▓▓\repos\carbonstack_umbrella\recon\v0.3.7\windows-reliance-scout.txt
C:\▓▓\repos\carbonstack_umbrella\recon\v0.3.7\script-surface-scout.txt
C:\▓▓\repos\carbonstack_umbrella\recon\v0.3.7\go-rust-path-assumption-scout.txt
C:\▓▓\repos\carbonstack_umbrella\recon\v0.3.7\release-platform-wording-scout.txt
C:\▓▓\repos\carbonstack_umbrella\recon\v0.3.7\local-toolchain-context.txt
```

The broad Windows reliance scout is noisy and should not drive implementation by itself.

### [WIN-PWRSHL] Clean snapshot validation

```text
C:\▓▓\repos\carbonstack_umbrella\clean-snapshot-validation\v0.3.13
C:\▓▓\repos\carbonstack_umbrella\clean-snapshot-validation\v0.3.13\carbonstack-v0.3.13-clean-working-snapshot.tgz
C:\▓▓\repos\carbonstack_umbrella\clean-snapshot-validation\v0.3.13\windows-clean
```

### [WSL-DEBIAN] Clean snapshot validation

```text
~/carbonstack-v0.3.13-clean
```


### [WIN-PWRSHL] Release-snapshot validation v0.3.15

```text
C:\▓▓\repos\carbonstack_umbrella\release-snapshot-validation\v0.3.15
C:\▓▓\repos\carbonstack_umbrella\release-snapshot-validation\v0.3.15\package-root
C:\▓▓\repos\carbonstack_umbrella\release-snapshot-validation\v0.3.15\carbonstack-v0.3.15-release-snapshot-package.tgz
C:\▓▓\repos\carbonstack_umbrella\release-snapshot-validation\v0.3.15\windows-release-snapshot
```

### [WSL-DEBIAN] Release-snapshot validation v0.3.15

```text
~/carbonstack-v0.3.15-release-snapshot
```


### [WIN-PWRSHL] Release-snapshot validation v0.3.16

```text
C:\▓▓\repos\carbonstack_umbrella\release-snapshot-validation\v0.3.16
C:\▓▓\repos\carbonstack_umbrella\release-snapshot-validation\v0.3.16\package-source-root
C:\▓▓\repos\carbonstack_umbrella\release-snapshot-validation\v0.3.16\carbonstack-v0.3.16-staged-release-like-package.tgz
C:\▓▓\repos\carbonstack_umbrella\release-snapshot-validation\v0.3.16\windows-release-snapshot
```

### [WSL-DEBIAN] Release-snapshot validation v0.3.16

```text
~/carbonstack-v0.3.16-release-snapshot
```

### [WIN-PWRSHL] Release verification / staging

```text
C:\▓▓\repos\carbonstack_umbrella\release-verification\v0.3.0
C:\▓▓\repos\carbonstack_umbrella\release-verification\v0.3.4
C:\▓▓\repos\carbonstack_umbrella\release-staging\v0.3.4
C:\▓▓\repos\carbonstack_umbrella\release-verification\v0.3.4-public
```

### CarbonStackComms critical surfaces

```text
carbonstack-comms/README.md
carbonstack-comms/scripts/self-test-openmls-backbone.ps1
carbonstack-comms/scripts/smoke-openmls-real-cypher-relay.ps1
carbonstack-comms/scripts/check-no-rust-artifacts.ps1
carbonstack-comms/scripts/README.md
carbonstack-comms/internal/protocol/
carbonstack-comms/internal/protocol/mls/openmls-sidecar/
carbonstack-comms/internal/relay/
carbonstack-comms/internal/client/
carbonstack-comms/go.mod
```

`carbonstack-comms/go.sum` does not currently exist. Treat as optional/include-if-present.

### CarbonStackCypher critical surfaces

```text
carbonstack-cypher/README.md
carbonstack-cypher/cmd/cypher/
carbonstack-cypher/internal/httpapi/
carbonstack-cypher/internal/db/
carbonstack-cypher/internal/config/
carbonstack-cypher/migrations/
carbonstack-cypher/docs/
carbonstack-cypher/go.mod
carbonstack-cypher/go.sum
```

### CarbonStackOS critical surfaces

```text
carbonstack-os/README.md
carbonstack-os/LICENSE
```

---

## 7. Generated / Private Validation Outputs

Do not ship or commit:

```text
target/
.carbonstack-openmls-sidecar-state/
provider-storage.json
signer.json
*.db
*.db-shm
*.db-wal
*.exe
*.test.exe
.go-cache/
.go-tmp/
```

After validation, these may be present in clean-run roots. Their presence after tests is expected, but they must remain excluded from release snapshots.

v0.3.8 added an artifact-hygiene lesson: the WSL tar bridge copied `carbonstack-cypher/cypher.db`. Remove DBs before tests and update future tar bridge excludes to include `*.db`, `*.db-shm`, `*.db-wal`, `*.exe`, and `*.test.exe`.

v0.3.9 post-test WSL artifact scan found generated OpenMLS dev state and Rust `target/` under the sidecar tree. This is expected after tests and did not indicate a shipped-source artifact issue.

v0.3.11 runner artifact scanning was broad and non-destructive. v0.3.12 improves the scan labels/classification while keeping the scan non-destructive.

No cleanup script is required yet as long as test artifacts remain in known test/build roots and are documented. If artifacts begin scattering into ambiguous paths, revisit cleanup tooling.

---

## 8. Known-Good Validation

### [WIN-PWRSHL] Go runner doctor/core/full

```powershell
cd C:\▓▓\repos\carbonstack_umbrella\carbonstack\tools\carbonstack-validate

go test ./...
go run . --profile doctor
go run . --profile core
go run . --profile full
```

Important nested-module note: run `go test ./...` from inside `carbonstack/tools/carbonstack-validate`. Do not run `go test .\tools\carbonstack-validate\...` from the parent `carbonstack` root unless a future `go.work` workspace is intentionally added.

If root inference behaves unexpectedly, use explicit root override:

```powershell
go run . --profile core --root C:\▓▓\repos\carbonstack_umbrella
```

### [WSL-DEBIAN] Go runner doctor/core/full

```bash
. "$HOME/.cargo/env"

cd "$HOME/carbonstack-wsl/carbonstack/tools/carbonstack-validate"

go test ./...
go run . --profile doctor
go run . --profile core
go run . --profile full
```

### [WSL-DEBIAN] v0.3.13 clean working snapshot runner validation

```bash
. "$HOME/.cargo/env"
cd "$HOME/carbonstack-v0.3.13-clean/carbonstack/tools/carbonstack-validate"

go test ./...
go run . --profile doctor
go run . --profile core
go run . --profile full
```

Passed at v0.3.13.

### [WIN-PWRSHL] v0.3.13 clean working snapshot runner validation

```powershell
cd C:\▓▓\repos\carbonstack_umbrella\clean-snapshot-validation\v0.3.13\windows-clean\carbonstack\tools\carbonstack-validate

go test ./...
go run . --profile doctor
go run . --profile core
go run . --profile full
```

Passed at v0.3.13.


### [WSL-DEBIAN] v0.3.15 release-snapshot validation

```bash
. "$HOME/.cargo/env"
cd "$HOME/carbonstack-v0.3.15-release-snapshot/carbonstack/tools/carbonstack-validate"

go test ./...
go run . --profile release-snapshot --root "$HOME/carbonstack-v0.3.15-release-snapshot"
```

Passed at v0.3.15 from a fresh extraction.

### [WIN-PWRSHL] v0.3.15 release-snapshot validation

```powershell
cd C:\▓▓\repos\carbonstack_umbrella\release-snapshot-validation\v0.3.15\windows-release-snapshot\carbonstack\tools\carbonstack-validate

go test ./...
go run . --profile release-snapshot --root C:\▓▓\repos\carbonstack_umbrella\release-snapshot-validation\v0.3.15\windows-release-snapshot
```

Passed at v0.3.15 from a fresh extraction.

### [WIN-PWRSHL] v0.3.15 package-root rule

Do not run `release-snapshot` against the package root that will be tarred/published. It generates artifacts. Use a throwaway extraction for validation.

Allowed pre-tar sanity checks on the package root:

```powershell
cd C:\▓▓\repos\carbonstack_umbrella\release-snapshot-validation\v0.3.15\package-root\carbonstack\tools\carbonstack-validate

Select-String -Path .\main.go -Pattern 'release-snapshot|expected doctor'
Test-Path .\release_snapshot.go
go test ./...
```


### [GO-RUNNER] future runner-backed public testing path

Documented at v0.3.17:

```text
carbonstack/docs/136-runner-backed-public-testing-path-v0.md
```

Core future command shape from a fresh extracted package root:

```text
cd carbonstack/tools/carbonstack-validate
go test ./...
go run . --profile release-snapshot --root <fresh-extracted-package-root>
```

This is documentation for a future runner-backed public testing release. It does not supersede the current public v0.3.4 Windows/PowerShell release path yet.

### [WIN-PWRSHL] Legacy validation still available

```powershell
powershell -ExecutionPolicy Bypass -File .\scripts\validate-local.ps1
```

### [WIN-PWRSHL] From public downloaded/extracted v0.3.4 source snapshots

```powershell
cd .\carbonstack-comms
powershell -ExecutionPolicy Bypass -File .\scripts\self-test-openmls-backbone.ps1
powershell -ExecutionPolicy Bypass -File .\scripts\self-test-openmls-backbone.ps1 -Full

cd .\carbonstack
powershell -ExecutionPolicy Bypass -File .\scripts\validate-local.ps1
```

All passed in the v0.3.5 verification on the Windows public release path.

### [WSL-DEBIAN] Rust toolchain setup for OpenMLS sidecar

```bash
. "$HOME/.cargo/env"
command -v rustup
command -v rustc
command -v cargo
rustup default stable
rustc --version
cargo --version
```

Expected working Rust after v0.3.9+:

```text
rustc 1.96.0
cargo 1.96.0
```

Expected Rust/Cargo paths:

```text
~/.cargo/bin/rustc
~/.cargo/bin/cargo
```

### [WSL-DEBIAN] Direct package tests still available

```bash
cd "$HOME/carbonstack-wsl/carbonstack-comms"

go test ./internal/protocol -run TestOpenMLSSidecarFullLifecycleRelayThroughRealCypherServer -count=1 -timeout 300s -v
go test ./internal/protocol -count=1 -timeout 600s -v
go test ./... -count=1 -timeout 600s

cd "$HOME/carbonstack-wsl/carbonstack-cypher"

go test ./... -count=1
```

Passed at v0.3.9 before the Go runner existed and is now wrapped by the runner.


### [WSL-DEBIAN] v0.3.16 staged release-like validation

```bash
. "$HOME/.cargo/env"
cd "$HOME/carbonstack-v0.3.16-release-snapshot/carbonstack/tools/carbonstack-validate"

go test ./...
go run . --profile release-snapshot --root "$HOME/carbonstack-v0.3.16-release-snapshot"
```

Passed at v0.3.16 from a fresh extraction.

### [WIN-PWRSHL] v0.3.16 staged release-like validation

```powershell
cd C:\▓▓\repos\carbonstack_umbrella\release-snapshot-validation\v0.3.16\windows-release-snapshot\carbonstack\tools\carbonstack-validate

go test ./...
go run . --profile release-snapshot --root C:\▓▓\repos\carbonstack_umbrella\release-snapshot-validation\v0.3.16\windows-release-snapshot
```

Passed at v0.3.16 from a fresh extraction.

---

## 9. Allowed Claims

Allowed:

```text
CarbonStack is an experimental secure-communications appliance-stack project.
The current verified public testing artifact is the v0.3.x experimental Cypher + Comms OpenMLS relay backbone.
CarbonStack v0.3.20 is a pre-alpha / experimental runner-backed testing release.
v0.3.20 is the first public runner-backed testing release for the current backbone.
v0.3.20 supersedes v0.3.4 as the preferred public testing path.
v0.3.4 remains available as the older Windows 11 / PowerShell legacy testbed.
The v0.3.20 intended multi-repo package is carbonstack-v0.3.20-runner-backed-testing-release.tgz, not the default Gitea Source Code archives.
The v0.3.20 package includes carbonstack, carbonstack-comms, carbonstack-cypher, and release metadata.
carbonstack-os is related future appliance-stack work but is not included in the current runnable v0.3.x validation package.
v0.3.20 includes real package-internal SHA-256 checksum coverage.
v0.3.20 includes asset checksums for the uploaded release files.
v0.3.20 includes manifest, validation freeze, testing runbook, release notes, package checksums, asset checksums, and LICENSE.
v0.3.20 passed WSL Debian fresh-extraction validation.
v0.3.20 passed Windows 11 short-path fresh-extraction validation at C:\cs-v0320\package with CARGO_BUILD_JOBS=1.
v0.3.20 validates release package layout, Go-generated checksums, release-snapshot validation, OpenMLS real-Cypher lifecycle behavior, CarbonStackComms tests, CarbonStackCypher tests, and post-test artifact scoping.
Debian 13 / WSL2 Debian, linux/amd64 is the primary validated target for v0.3.20.
Windows 11, windows/amd64 is secondary/final explicit Windows dev/test validation for this phase.
After v0.3.20, CarbonStack mainline public dev/test releases should migrate to Debian / WSL Debian first.
```

Also still allowed:

```text
The uploaded public v0.3.4 release assets were downloaded, checksum-verified, extracted, and validated from no-Git source snapshots at v0.3.5.
The v0.3.4 release improves public testability of the experimental backbone from source snapshots.
The v0.3.7 recon found that Windows reliance is mostly in scripts/runbooks/validation, not obviously in the Go/Rust core.
v0.3.9 validated the targeted OpenMLS real-Cypher lifecycle and package tests under WSL Debian after rustup stable correction.
v0.3.11 implemented the first umbrella Go validation runner.
v0.3.18 implemented real checksum helper profiles and release-snapshot checksum verification.
v0.3.19 local downloaded-asset simulation passed on WSL Debian and passed on Windows with a short-root caveat.
```

No safety/security claim may be made without a validation record.

No validation record may exist without a stated non-scope.

No passing test may be generalized beyond what it actually tested.

## 10. Not Allowed Claims

Do not claim:

```text
finished secure messenger
production E2E
production-ready
hostile-server safe
metadata-private
trustless
drop-in replacement
finished product
Android-ready
externally audited
certified
audit-ready
Debian deployable
deployment-ready
systemd-ready
cloudflared-ready
public ingress safe
real homelab validated
universal release ready
Windows remains mainline after v0.3.20
arbitrary long nested Windows extraction roots are known-good for Rust/OpenMLS validation
Go runner proves security
Go runner proves deployment readiness
release-snapshot proves production security
v0.3.20 proves deployability
v0.3.20 proves public server readiness
v0.3.20 proves production E2EE
v0.3.20 proves hostile-server safety
v0.3.20 proves metadata privacy
v0.3.20 proves audit/certification
```

Clarified after v0.3.20:

```text
v0.3.20 is public and runner-backed, but it is still pre-alpha / experimental.
v0.3.20 validates the release/testing package, not production deployment.
v0.3.20 marks Debian / WSL Debian as primary for the runner-backed path, but not as a production deployment target.
Windows short-path validation passed, but long nested Windows paths remain suspect due to the v0.3.19 Rust/MSVC LNK1104 failure.
```

No safety/security claim may be made without a validation record.

No validation record may exist without a stated non-scope.

No passing test may be generalized beyond what it actually tested.

## 11. Phase Status

```text
Phase 1:
  local relay/client skeleton stable and test-protected.

Phase 2A:
  trust-state scaffold and trust UX polish validated at dev-scaffold level.

Phase 2B:
  provider-boundary doc and provider-neutral code skeleton validated; no real provider in CLI.

Phase 2C:
  closed for mainline work; OpenMLS feasibility/provider-contract/trust-decision shaping validated.

Phase 2D:
  mainline closed for research continuity; promoted sidecar and Comms/Cypher backbone proof are stable enough for v0.3.x maturation.

Phase 2E:
  active v0.3.x maturation/release substrate phase. v0.3.0 released the experimental backbone proof. v0.3.4 fixed the older Windows/PowerShell source-snapshot testing path. v0.3.5-v0.3.19 progressively verified assets, built the public tester runbook, established WSL Debian validation, designed/implemented/hardened the Go runner, implemented release-snapshot, implemented real checksum support, rehearsed local release-candidate and downloaded-asset flows, and captured the Windows short-path caveat. v0.3.20 cut the first public runner-backed testing release with Debian / WSL Debian as the primary validation target and Windows short-path validation as secondary/final explicit Windows dev/test validation for this phase.

v0.3.x:
  active line. Public runner-backed testing release now exists at v0.3.20. Next direction is local-only backbone deployability planning and implementation, Debian/server-like layout first. Do not jump to production, cloudflared, systemd, public ingress, real homelab, or IRC-style server/admin mechanics until local deployability basics are ironed out.
```

## 12. Blunders / Repair Lessons Preserved

- Public runbooks should target broad testers early; this prevents later semantic drift and makes release behavior easier to stabilize.
- Version-specific runbooks should be included in release surfaces because defaults, setup, and known-good paths may change.
- Main README should point to the latest release for testing/dev instead of carrying every release-specific runbook forever.
- Current public release surfaces should state platform target clearly; v0.3.4 is currently Windows 11 / PowerShell validated.
- v0.3.7’s broad Windows scout was too noisy to drive implementation. Use focused scouts and actual validation results.
- Windows reliance appears concentrated in PowerShell scripts, release/runbook commands, and validation UX rather than obvious Go/Rust core implementation.
- WSL Debian is useful as a quick portability bridge, but it is not a deployment proof and not a substitute for real Debian homelab validation.
- v0.3.8 confirmed that WSL Debian setup/toolchains are available enough for the next quick-portability test rung.
- The WSL tar bridge copied `carbonstack-cypher/cypher.db`; this is a generated artifact hygiene issue, not protocol failure. Remove it before tests and add DB/exe excludes to future tar bridge commands.
- v0.3.9 initially looked like a massive Comms protocol failure, but it was a Rust toolchain floor issue.
- Debian apt Rust 1.85.0 was too old for OpenMLS 0.8.1 because OpenMLS used `is_multiple_of` APIs unavailable on that Rust version.
- Go tests hid useful Rust/Cargo stderr too much when sidecar compilation failed; future Go test harnesses or the Go validation runner should surface sidecar build/stderr output clearly.
- rustup stable Rust 1.96.0 fixed the WSL OpenMLS sidecar build/test path.
- WSL Debian tests were much faster and avoided Windows antivirus interruption seen in the Windows Go test executable path.
- WSL Debian is now plausible as a future preferred core-validation environment, but only after a cross-platform validation runner exists and public release docs are updated honestly.
- v0.3.10 correctly kept the runner rung design-only; this avoided implementing shell/path assumptions before agreeing on authority, profiles, and non-scope.
- v0.3.11 found a Windows PowerShell encoding blunder: `Set-Content -Encoding UTF8` in Windows PowerShell 5.1 wrote a BOM to `go.mod`; Go rejected this with `unexpected input character '\ufeff'`. Use BOM-free UTF-8 writes for Go module/source files.
- v0.3.12 reinforced that Windows PowerShell 5.1 does not support `Set-Content -Encoding UTF8NoBOM`; use .NET `System.Text.UTF8Encoding($false)` after writing content when BOM-free output matters.
- v0.3.11 found and fixed an initial root-inference flaw: launching from `carbonstack/tools/carbonstack-validate` must infer the umbrella root by walking upward and checking sibling repo layout, not by assuming the current working directory is the umbrella root.
- Explicit `--root` remains useful and worked on Windows before inference was fixed.
- The Go runner should remain boring and explicit: print environment, run known commands, stream output, fail clearly, and avoid hidden magic.
- The runner should not install dependencies, configure deployment, expose cloudflared, create systemd services, delete artifacts, or make security claims.
- Real Debian homelab tests are deferred much later, closer to deployability and IRC-style setup needs; WSL Debian is enough for current Linux-like quick-portability work.
- The real Debian homelab already hosts services; do not experiment recklessly on it.
- Initial real Debian validation, when eventually needed, must be local-only, separate-directory, manual, and non-public.
- Do not expose Cypher through cloudflared until local validation and operator model work justify it.
- v0.3.3 correctly found that public source snapshots intentionally lack `.git/`; validation scripts that rely on Git need non-Git fallback behavior.
- v0.3.5 verified the actual uploaded v0.3.4 release assets, not just local release-staging outputs.
- A release is not fully trustworthy just because local staging passed; downloaded public assets need their own verification pass.
- Pre-test artifact scan and post-test artifact scan have different meanings. Pre-test hits would indicate shipped/copied artifacts. Post-test hits are expected generated validation outputs.
- Test artifacts are acceptable without a cleanup script for now as long as they remain in specific known roots and are clearly documented.
- Gitea default source archives are not the intended multi-repo release package; attach explicit multi-repo source snapshots.
- Keep v0.3.0 historical; do not move the tag. v0.3.4 is the recommended public testing release.
- Tailscale and Cloudflare WARP can interfere with localhost/ephemeral port real-server tests.
- Avast may interfere with Go-generated test executables. Keep any workaround narrow to the CarbonStack Go temp/cache root, not all of `%TEMP%`.
- `carbonstack-comms/go.sum` never existed; do not require it unless dependency state creates it.
- Runtime `cmd/comms send/inbox` remains stub-era; do not present it as OpenMLS messenger UX.

- v0.3.13 refined artifact interpretation: generated `provider-storage.json`, `signer.json`, and Windows Cargo `.exe` files are acceptable only after validation and only under known OpenMLS sidecar generated roots. Filename-only scans look scary on Windows because Cargo emits many `.exe` build helpers under `target/`; root scoping matters.
- Running `full` immediately after `core` in the same clean snapshot will make `full`'s pre-test scan report artifacts produced by the prior `core` run. This is expected in a reused validation root and should be interpreted by run order, not as shipped-source contamination.
- Windows Rust `1.95.0` currently passes the clean snapshot validation even though the runner warns it is below the WSL known-good `1.96.0` floor. Treat the Rust floor warning as conservative until a formal minimum supported Rust policy is defined.
- v0.3.14 preserved a nested Go module lesson: `tools/carbonstack-validate` is a standalone Go module. Run `go test ./...` from inside that directory; running `go test .\tools\carbonstack-validate\...` from the non-Go-module `carbonstack` root fails unless a future `go.work` workspace is intentionally added.
- v0.3.14 kept release-snapshot design-only, preventing implementation before the release package layout, strict pre-test artifact policy, metadata expectations, and nonclaims were agreed.
- v0.3.15 found a stale package runner failure: WSL extracted package validation reported `unknown profile "release-snapshot"` because the tarball/package contained an old runner or `main.go` lacked the dispatch case. Always sanity-check the package-root runner contains `release-snapshot` before creating the tarball.
- v0.3.15 confirmed that running `release-snapshot` against the live umbrella root should fail unless the live root has a formal `release/` metadata folder. This is expected and distinguishes dev roots from formal release-like roots.
- v0.3.15 found a package-root contamination blunder: running `release-snapshot` on the package root before tar creation generated OpenMLS `target/` and `.carbonstack-openmls-sidecar-state/`, contaminating the tarball. The strict pre-test scan correctly rejected the contaminated package.
- Correct v0.3.15 rule: never validate the package root intended for tar/publishing. Validate throwaway extractions. Only run non-generating sanity checks on the packaging source root before tar creation.
- v0.3.16 exact PowerShell string replacement failed against `release_snapshot.go` because the expected `$Old` block did not match `gofmt` indentation. Repair was to inspect with `cat`/`Select-String`, use regex replacement, run `gofmt`, and confirm with `Select-String`.
- v0.3.16 preserves the package-source-root rule directly in runner output: validate only fresh extracted/staged package roots; do not validate the package source root that will later be archived/published; successful validation generates artifacts, so rerun from a fresh extraction.
- v0.3.16 checksum policy remains intentionally incomplete: placeholder checksum text is acceptable for local staged validation hardening, but public runner-backed releases require real checksum generation and verification before uploaded/downloaded asset verification.

- Payload metadata is not a trust root, not an authenticity proof, and not a hostile-server guarantee.
- Ack is not delivery/read proof; Cypher records recipient-device ack, while Comms decides when to ack after sidecar consume.

---


- v0.3.17 preserved a critical public-surface distinction: documenting a future runner-backed public testing path does not supersede the existing v0.3.4 public release or imply v0.3.4 is runner-backed.
- Tester-facing docs must state both the positive validation path and the nonclaims; otherwise a passing `release-snapshot` run can be misread as public Linux support, deployability, or security certification.

## 13. Next TODO

Immediate next safest work after v0.3.20:

```text
1. Mark this v0.3.20 breakpoint and preserve release evidence.
2. Treat v0.3.20 as the preferred public runner-backed testing release.
3. Treat v0.3.4 as the older Windows 11 / PowerShell legacy testbed.
4. Do not continue release-packaging churn unless a real post-release bug is found.
5. Begin v0.3.21+ local-only backbone deployability planning.
6. Keep Debian / WSL Debian as the mainline validation and development environment.
7. Keep Windows out of mainline release/deployability work unless a specific compatibility check is intentionally needed.
8. Define local server-like layout before touching real homelab or public ingress.
9. Define Cypher bind address, port, config, and data-dir policy.
10. Define Comms-to-Cypher config / endpoint policy.
11. Define start/stop lifecycle, logs, errors, and validation profile needs for local deployability.
12. Keep `release-verify` deferred until deployable backbone and IRC-style join/admin mechanics mature enough to justify a dedicated public release verification surface.
```

Do not start yet:

```text
runtime OpenMLS send/inbox UX as a user product
Android / CarbonStackOS implementation
secure vault implementation
hostile-server-complete claims
metadata privacy claims
Debian deployability claims
systemd service work
cloudflared exposure
public ingress
real homelab deployment
shared production service paths
MariaDB integration
external audit/certification language
Windows-first release mode after v0.3.20
```

Remaining v0.3.x minor epoch workflow:

```text
v0.3.21+:
  actual local-only backbone deployability planning and implementation;
  Debian/server-like layout first;
  Cypher local service/config/data-dir semantics;
  Comms-to-Cypher endpoint configuration;
  lifecycle/log/error/runbook shape;
  runner validation for local deployability when concrete enough.

After local deployability is ironed out:
  deployable server-system mechanics planning;
  IRC-like CarbonStack standard server systems;
  join/discovery/operator/admin workflows;
  hostile-endpoint doctrine versus normal server administration;
  metadata visibility boundaries and client-side verification duties.
```

## 13B. v0.3.18 Critical Path / Command Addendum

### [GO-RUNNER] checksum profiles

```text
write_checksums:      go run . --profile write-checksums --root <package-root>
verify_checksums:     go run . --profile verify-checksums --root <package-root>
release_snapshot:     go run . --profile release-snapshot --root <package-root>
```

`release-snapshot` now verifies real checksums before calling `core`.

### [WIN-PWRSHL] v0.3.18 staging paths

```text
release_staging_root: C:\▓▓\repos\carbonstack_umbrella\release-staging\v0.3.18
package_source_root:  C:\▓▓\repos\carbonstack_umbrella\release-staging\v0.3.18\local-release-candidate-rehearsal
carrier_archive:      C:\▓▓\repos\carbonstack_umbrella\release-staging\v0.3.18\carbonstack-v0.3.18-local-release-candidate-rehearsal.tgz
windows_validation:   C:\▓▓\repos\carbonstack_umbrella\release-staging\v0.3.18\windows-release-rehearsal\package
```

### [WSL-DEBIAN] v0.3.18 validation paths

```text
validation_root:      ~/carbonstack-v0.3.18-release-rehearsal
downloads_dir:        ~/carbonstack-v0.3.18-release-rehearsal/downloads
package_root:         ~/carbonstack-v0.3.18-release-rehearsal/package
runner_dir:           ~/carbonstack-v0.3.18-release-rehearsal/package/carbonstack/tools/carbonstack-validate
```

### v0.3.18 known-good validation pattern

```text
carrier archive stays in downloads/
package root contains only extracted package contents
go test ./...
go run . --profile verify-checksums --root <package-root>
go run . --profile release-snapshot --root <package-root>
```

Do not place the carrier `.tgz` inside the package root. The runner defensively skips archive suffixes, but the workflow still requires a clean package root.


## 13C. v0.3.19 Critical Path / Command Addendum

### [WIN-PWRSHL] v0.3.19 local simulation paths

```text
verification_root:  C:\▓▓\repos\carbonstack_umbrella\release-verification\v0.3.19-local-download-simulation
published_assets:   C:\▓▓\repos\carbonstack_umbrella\release-verification\v0.3.19-local-download-simulation\published-assets
downloads:          C:\▓▓\repos\carbonstack_umbrella\release-verification\v0.3.19-local-download-simulation\downloads
long_package_root:  C:\▓▓\repos\carbonstack_umbrella\release-verification\v0.3.19-local-download-simulation\extracted-windows-clean\package
short_package_root: C:\cs-v0319-win2\package
```

### [WSL-DEBIAN] v0.3.19 local simulation paths

```text
validation_root: ~/carbonstack-v0.3.19-local-download-simulation
downloads_dir:   ~/carbonstack-v0.3.19-local-download-simulation/downloads
package_root:    ~/carbonstack-v0.3.19-local-download-simulation/package
runner_dir:      ~/carbonstack-v0.3.19-local-download-simulation/package/carbonstack/tools/carbonstack-validate
```

### [WIN-PWRSHL] known-good Windows short-path validation pattern

```powershell
$PackageRoot = "C:\cs-v0319-win2\package"
$env:RUST_BACKTRACE = "full"
$env:CARGO_BUILD_JOBS = "1"
cd "$PackageRoot\carbonstack\tools\carbonstack-validate"
go test ./...
go run . --profile verify-checksums --root "$PackageRoot"
go run . --profile release-snapshot --root "$PackageRoot"
```

### [WSL-DEBIAN] known-good local downloaded-asset validation pattern

```bash
. "$HOME/.cargo/env"
cd "$HOME/carbonstack-v0.3.19-local-download-simulation/package/carbonstack/tools/carbonstack-validate"
go test ./...
go run . --profile verify-checksums --root "$HOME/carbonstack-v0.3.19-local-download-simulation/package"
go run . --profile release-snapshot --root "$HOME/carbonstack-v0.3.19-local-download-simulation/package"
```

### v0.3.19 stable pause interpretation

v0.3.19 is a stable pause point after local downloaded-asset simulation. The package/release-snapshot/checksum path passed on WSL Debian and on Windows using a short extraction path. The long nested Windows extraction path exposed a Rust/MSVC linker/file-lock/path sensitivity and should be recorded as a validation caveat, not buried as a random transient failure.


## 13D. v0.3.20 Critical Path / Command Addendum

### [WIN-PWRSHL] v0.3.20 release staging paths

```text
C:\▓▓\repos\carbonstack_umbrella\release-staging\v0.3.20
C:\▓▓\repos\carbonstack_umbrella\release-staging\v0.3.20\runner-backed-testing-release-candidate
C:\▓▓\repos\carbonstack_umbrella\release-staging\v0.3.20\carbonstack-v0.3.20-runner-backed-testing-release.tgz
C:\▓▓\repos\carbonstack_umbrella\release-staging\v0.3.20\assets
```

### [WSL-DEBIAN] v0.3.20 validation path

```text
~/carbonstack-v0.3.20-test/downloads
~/carbonstack-v0.3.20-test/package
```

Known-good command shape:

```bash
. "$HOME/.cargo/env"
cd "$HOME/carbonstack-v0.3.20-test/package/carbonstack/tools/carbonstack-validate"
go test ./...
go run . --profile verify-checksums --root "$HOME/carbonstack-v0.3.20-test/package"
go run . --profile release-snapshot --root "$HOME/carbonstack-v0.3.20-test/package"
```

### [WIN-PWRSHL] v0.3.20 short-path Windows validation

Known-good root:

```text
C:\cs-v0320\package
```

Known-good environment:

```powershell
$env:RUST_BACKTRACE = "full"
$env:CARGO_BUILD_JOBS = "1"
```

Known-good command shape:

```powershell
cd C:\cs-v0320\package\carbonstack\tools\carbonstack-validate
go test ./...
go run . --profile verify-checksums --root "C:\cs-v0320\package"
go run . --profile release-snapshot --root "C:\cs-v0320\package"
```

### [PUBLIC RELEASE] v0.3.20 release URL and assets

```text
https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.3.20
```

Assets:

```text
carbonstack-v0.3.20-runner-backed-testing-release.tgz
carbonstack-v0.3.20-release-manifest.json
carbonstack-v0.3.20-package-checksums.txt
carbonstack-v0.3.20-asset-checksums.txt
carbonstack-v0.3.20-validation-freeze.md
v0.3.20-testing-runbook.md
v0.3.20-release-notes.md
LICENSE
```

### v0.3.20 stable pause interpretation

```text
Public runner-backed testing release exists.
Debian / WSL Debian is primary.
Windows short-path validation is secondary/final explicit Windows dev/test validation for this phase.
v0.3.4 remains a legacy Windows/PowerShell testbed.
Next work should be local-only deployability planning, not more release packaging, unless a real post-release bug appears.
```

## 14. Handoff Summary

v0.3.20 is a stable breakpoint.

The project now has a public runner-backed testing release at:

```text
https://git.bitcrusher32.win/bitcrusher32/carbonstack/releases/tag/v0.3.20
```

The release is tagged at `3eeb1a1e1a` / `3eeb1a1 docs: clean public release surface`.

The v0.3.20 release package and release page preserve the intended boundary:

```text
pre-alpha / experimental;
Cypher + Comms OpenMLS relay backbone;
Debian 13 / WSL2 Debian primary validation target;
Windows 11 short-path validation secondary/final explicit Windows dev-test validation for this phase;
v0.3.4 remains the older Windows 11 / PowerShell legacy testbed.
```

The v0.3.20 work before release included:

```text
public-surface cleanup across README/docs/roadmap/LogDoc archive pointer;
stale release wording cleanup;
final v0.3.20 package rebuild;
real package checksum regeneration;
asset checksum creation;
WSL Debian fresh extraction validation;
Windows short-path fresh extraction validation;
release page formatting to match v0.3.4 continuity;
Gitea release asset upload/publication.
```

The key v0.3.20 continuity point is release-readiness honesty: the public runner-backed testing release is real, but it validates the package/test backbone only. It does not prove deployability, production security, hostile-server safety, metadata privacy, Android readiness, audit, or certification.

The next safest work is v0.3.21+ local-only backbone deployability planning/implementation with Debian/server-like layout first. Do not touch real homelab, cloudflared, systemd, public ingress, shared production service paths, MariaDB, or IRC-style server/admin mechanics until local deployability basics are ironed out.

