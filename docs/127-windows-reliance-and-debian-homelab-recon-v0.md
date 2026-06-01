# Windows Reliance and Debian Homelab Recon v0

Status: recon / planning record
Phase: v0.3.7 backbone maturation
Previous checkpoint: docs/126-v0.3.4-public-tester-runbook-v0.md

## 1. Purpose

This document records the v0.3.7 Windows reliance and Debian homelab recon.

The goal is to classify portability constraints before implementing Linux/Debian support.

This is a recon checkpoint, not a portability fix.

## 2. Current public platform boundary

The current recommended public testing release is:

    v0.3.4 testing bug hotfix

Current public testing path:

    attached v0.3.4 release source snapshots
    Windows / PowerShell
    local Go toolchain
    local Rust/Cargo toolchain
    no-Git source snapshot self-test path

The v0.3.4 release header was updated to mark the platform target as:

    Platform: Windows 11

This is correct public hygiene. The validated release path is Windows/PowerShell. Linux, Debian, WSL, and deployment readiness are not yet validated.

## 3. Recon inputs

Recon inputs included:

    Windows reliance scout
    script surface scout
    Go/Rust path assumption scout
    release/platform wording scout
    local Windows toolchain context
    Debian homelab scout

The broad Windows reliance scout was noisy because it searched too broadly and returned a large amount of historical docs and generated/platform-symbol noise. It is useful only as a rough warning. Focused scouts were higher signal.

## 4. Windows reliance classification

The active portability reliance is mostly at the script/runbook/validation layer.

Active script surfaces include:

    carbonstack/scripts/validate-local.ps1
    carbonstack/scripts/validate-phase1.ps1
    carbonstack-comms/scripts/check-no-rust-artifacts.ps1
    carbonstack-comms/scripts/self-test-openmls-backbone.ps1
    carbonstack-comms/scripts/smoke-openmls-real-cypher-relay.ps1
    carbonstack-comms/scripts/test-local-lifecycle.ps1
    carbonstack-comms/scripts/test-trust-lifecycle.ps1

Current release and validation docs also describe PowerShell commands.

This means the public testing UX is Windows-first.

## 5. Core implementation portability classification

The Go/Rust code does not appear obviously Windows-only from the recon.

Current implementation uses common portable primitives such as:

    filepath.Join
    t.TempDir
    local HTTP listeners
    SQLite-backed test DB paths
    Rust path joins for OpenMLS sidecar state

Known state files include:

    provider-storage.json
    signer.json

These are not inherently Windows-only.

However, Linux/Debian portability is not validated until the tests actually run on a Linux environment.

## 6. Current local validation environment

The local validation environment recorded by recon was:

    Windows / PowerShell
    PowerShell 5.1
    Go windows/amd64
    Rust/Cargo installed
    dedicated CarbonStack Go temp/cache path

The Windows product/build output was not perfectly semantically clean, but the practical release target remains Windows/PowerShell.

## 7. Debian homelab target

The Debian homelab scout identified a plausible future validation/deploy target:

    Debian GNU/Linux 13 trixie
    x86_64 / amd64
    systemd
    4 CPU cores
    23 GiB RAM
    large ext4 root disk
    separate /mnt/ssd mount

The machine already runs real services, including:

    Gitea
    cloudflared tunnels
    WordPress/web stack components
    MariaDB
    nginx/PHP-related services

The machine is therefore a useful future deploy target, but CarbonStack testing must not interfere with existing services.

## 8. Debian cautions

Debian recon did not confirm the full CarbonStack build/test toolchain.

Confirmed:

    git

Not confirmed in the captured scout output:

    go
    rustc
    cargo
    sqlite3

Before running CarbonStack validation on Debian, confirm or install the missing toolchain pieces.

The scout also showed an input/output error on /mnt/usb. Treat this as machine hygiene outside CarbonStack unless that mount is intentionally used for CarbonStack state.

## 9. Debian safety posture

Initial Debian validation must be local-only.

Do not expose Cypher publicly during first Debian tests.

Do not route Cypher through cloudflared during first Debian tests.

Do not create a systemd service during first Debian tests.

Do not share MariaDB or existing production service paths.

Initial target:

    run from a separate test directory
    bind Cypher to localhost only
    use temp/local SQLite paths
    run tests manually
    record results
    delete generated test outputs manually if needed

## 10. WSL Debian bridge

WSL Debian is a good quick-portability bridge.

It can answer early questions before touching the real Debian homelab:

    do Go tests run under Linux paths?
    does the Rust sidecar build under Linux?
    do path assumptions survive Unix separators?
    can the self-test be expressed without Windows PowerShell?
    which scripts need bash or cross-platform replacements?

WSL Debian is not equivalent to the real Debian homelab.

It does not prove:

    systemd behavior
    cloudflared interaction
    firewall behavior
    long-running service behavior
    real deployability
    production ops readiness

Use WSL as portability triage, then validate on the real Debian host later.

## 11. Recommended next sequence

Recommended next v0.3.x sequence:

    v0.3.8:
        WSL Debian quick-portability triage plan

    v0.3.9:
        WSL Debian quick-portability test

    v0.3.10:
        Real Debian homelab local-only validation plan

    v0.3.11:
        Real Debian homelab local-only validation attempt

    later:
        low-level Cypher deploy config and operator model

This sequence avoids jumping from Windows release validation directly to public Debian deployment.

## 12. What this recon validates

This recon validates that:

    Windows/PowerShell reliance has been identified as mostly script/runbook-level.
    Core Go/Rust code is not obviously Windows-only from static scouting.
    Debian homelab target exists and is plausible.
    Debian homelab must be treated carefully because it already hosts real services.
    WSL Debian is useful as a quick triage layer before real homelab validation.

## 13. What this recon does not validate

This recon does not validate:

    Linux portability
    Debian readiness
    WSL compatibility
    Debian deployability
    systemd service behavior
    cloudflared routing
    public ingress
    production readiness
    production E2EE
    hostile-server safety
    metadata privacy
    secure vault/storage
    external audit
    certification

## 14. Next work

Next work should plan WSL Debian quick-portability triage.

Do not implement Debian deployment yet.

Do not expose Cypher through cloudflared.

Do not make Linux/Debian support claims until validation records exist.

Keep the current public release boundary:

    Platform: Windows 11
    v0.3.4 is the recommended public testing release
    Debian/Linux is future recon and validation work
