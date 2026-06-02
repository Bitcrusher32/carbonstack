# Runner-Backed Public Testing Path v0

Status: public testing documentation draft
Phase: v0.3.17 runner-backed public testing docs
Previous checkpoint: docs/135-release-snapshot-validation-hardening-v0.md

## 1. Purpose

This document records the first runner-backed public testing path for CarbonStack.

It does not replace the current public v0.3.4 release yet.

It does not publish a new release.

It does not verify uploaded/downloaded release assets.

It documents how future runner-backed testing releases should be validated once release staging and uploaded/downloaded verification are ready.

## 2. Current public release status

Current recommended public testing release:

    v0.3.4

Current v0.3.4 public testing path:

    Windows 11 / PowerShell

v0.3.17 does not supersede v0.3.4.

v0.3.17 documents the future runner-backed validation path that should be used by a later runner-backed public testing release.

## 3. Validation authority model

Preferred validation order:

    WSL Debian first;
    Windows second before breakpoint/release confirmation.

Preferred validation authority:

    Go runner.

Legacy/current public release authority:

    PowerShell path documented by v0.3.4/v0.3.6 runbooks.

The goal is not to replace PowerShell with WSL-specific bash authority.

The goal is to replace shell-specific umbrella validation with the Go runner wherever possible.

## 4. Runner location

Runner module:

    carbonstack/tools/carbonstack-validate

The runner is a standalone nested Go module.

Correct command shape:

    cd carbonstack/tools/carbonstack-validate
    go test ./...

Do not run:

    go test ./tools/carbonstack-validate/...

from the parent carbonstack repo unless a future go.work workspace is intentionally added.

## 5. Current runner profiles

Implemented profiles:

    doctor
    core
    full
    release-snapshot

Current behavior:

    doctor reports environment, repo paths, required paths, and toolchains.
    core runs doctor, artifact scans, OpenMLS real-Cypher lifecycle, Comms tests, and Cypher tests.
    full currently aliases core.
    release-snapshot validates a formal release-like package root, performs strict pre-test artifact checks, then calls core.

## 6. Expected future release package shape

Future runner-backed testing packages should use a formal package root:

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

carbonstack-os is not required for current v0.3.x backbone validation packages.

## 7. release-snapshot command

From a fresh extracted package root:

    cd carbonstack/tools/carbonstack-validate
    go test ./...
    go run . --profile release-snapshot --root <fresh-extracted-package-root>

The release-snapshot profile checks package layout and release metadata before running core.

## 8. Fresh extraction rule

release-snapshot must be run from a fresh extracted or throwaway staged package root.

Do not validate the package source root that will later be archived or published.

A successful release-snapshot run calls core, and core generates OpenMLS sidecar state and Rust build artifacts.

If that same package root is archived afterward, the archive will contain forbidden generated/private/build artifacts and should fail strict pre-test validation later.

Correct flow:

    create clean package source root;
    perform only non-generating sanity checks;
    archive the clean package source root;
    extract archive into a fresh throwaway validation root;
    run release-snapshot from the throwaway extraction;
    discard or preserve the throwaway extraction only as validation evidence.

Do not run release-snapshot twice in the same extraction unless you intentionally expect the second run to fail strict pre-test artifact scanning.

## 9. Forbidden pre-test artifacts

release-snapshot fails when these are present before tests:

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

A public release package should not ship generated/private/build artifacts.

## 10. Expected post-test artifacts

After validation, generated artifacts are expected only under known OpenMLS sidecar generated roots:

    carbonstack-comms/internal/protocol/mls/openmls-sidecar/target
    carbonstack-comms/internal/protocol/mls/openmls-sidecar/.carbonstack-openmls-sidecar-state

provider-storage.json and signer.json are acceptable only under the generated sidecar state root after tests.

Windows Cargo .exe files are acceptable only under the sidecar target root after tests.

Generated artifacts outside known roots should be treated as a validation problem.

## 11. WSL Debian path

WSL Debian is the preferred first validation environment.

Known-good WSL baseline from current v0.3.x validation:

    Debian GNU/Linux 13 trixie under WSL2;
    linux/amd64;
    Go 1.24.4;
    rustup stable Rust/Cargo 1.96.0;
    sqlite3 3.46.1.

Before running Rust/OpenMLS tests:

    . "$HOME/.cargo/env"

Example:

    rm -rf "$HOME/carbonstack-release-test"
    mkdir -p "$HOME/carbonstack-release-test"
    cd "$HOME/carbonstack-release-test"

    tar -xzf <release-package>.tgz

    . "$HOME/.cargo/env"

    cd "$HOME/carbonstack-release-test/carbonstack/tools/carbonstack-validate"

    go test ./...
    go run . --profile release-snapshot --root "$HOME/carbonstack-release-test"

## 12. Windows path

Windows remains the second confirmation environment before breakpoints/releases.

Known-good Windows baseline from current v0.3.x validation:

    Windows 11;
    Go 1.26.3 windows/amd64;
    Rust/Cargo 1.95.0;
    sqlite3 may be unavailable in PATH and remains a warning for current tests.

Example:

    cd C:\path\to\fresh\release-extraction\carbonstack\tools\carbonstack-validate

    go test ./...
    go run . --profile release-snapshot --root C:\path\to\fresh\release-extraction

Windows Rust 1.95.0 currently passes the tested path despite the conservative runner warning that it is below the WSL known-good 1.96.0 floor.

## 13. Known Rust policy

Known bad:

    Debian apt rustc 1.85.0 failed with OpenMLS 0.8.1.

Known good in current tests:

    WSL Debian rustup stable rustc/cargo 1.96.0;
    Windows rustc/cargo 1.95.0.

The current Rust warning is conservative.

Tests remain the authority.

## 14. What passing release-snapshot validates

Passing release-snapshot validates only:

    expected package layout exists;
    expected release metadata exists;
    forbidden generated/private/build artifacts were absent before tests;
    current core backbone validation passed;
    post-test generated artifacts stayed under expected OpenMLS sidecar roots.

## 15. What passing release-snapshot does not validate

Passing release-snapshot does not prove:

    production readiness;
    production E2EE;
    hostile-server safety;
    metadata privacy;
    Android readiness;
    secure vault/storage;
    public Linux release support;
    Debian deployability;
    systemd readiness;
    cloudflared readiness;
    public ingress safety;
    real homelab validation;
    external audit;
    certification.

## 16. Relationship to v0.3.4

v0.3.4 remains the current recommended public testing release until a later runner-backed public release is staged, uploaded, downloaded, and verified.

This document prepares the future runner-backed public testing path.

It should not be used to claim that v0.3.4 itself is runner-backed.

## 17. Next work

Recommended next work:

    define real manifest/checksum semantics;
    rehearse runner-backed release staging;
    verify uploaded/downloaded release assets;
    only then consider a public runner-backed testing release.