# WSL Debian Quick-Portability Setup Scout v0

Status: setup scout / planning record
Phase: v0.3.8 backbone portability triage
Previous checkpoint: docs/127-windows-reliance-and-debian-homelab-recon-v0.md

## 1. Purpose

This document records the v0.3.8 WSL Debian setup scout.

The goal was to establish a Debian WSL environment for quick portability triage before touching the real Debian homelab.

This is not a Linux support claim.

This is not Debian deployability validation.

This is not a public release.

## 2. Target

The selected quick-portability target is Debian under WSL.

Reason:

    Debian matches the future homelab target more closely than the default Ubuntu WSL environment.
    WSL lets the project catch Linux path/toolchain/script assumptions before touching the real homelab.
    WSL is safer than testing directly on a machine that already runs Gitea, WordPress, cloudflared, MariaDB, nginx, and other services.

## 3. WSL environment observed

Observed WSL environment:

    Debian GNU/Linux 13 trixie
    VERSION_ID=13
    VERSION_CODENAME=trixie
    DEBIAN_VERSION_FULL=13.5
    WSL2 Linux kernel
    x86_64 / amd64

Observed toolchains:

    git version 2.47.3
    go version go1.24.4 linux/amd64
    rustc 1.85.0
    cargo 1.85.0
    sqlite3 3.46.1

This is sufficient for the next quick-portability test rung.

## 4. Workspace

Workspace used:

    ~/carbonstack-wsl

Copied working repositories:

    carbonstack
    carbonstack-comms
    carbonstack-cypher

The initial copy used a tar bridge from the Windows working tree into the WSL Debian home filesystem.

The WSL home filesystem is preferred for this triage over working directly under /mnt/c because it more closely reflects normal Linux filesystem behavior.

## 5. Important artifact hygiene finding

The initial tar bridge copied:

    carbonstack-cypher/cypher.db

This is a generated/local database artifact and should not be part of future WSL triage snapshots.

This is not a protocol failure.

It is a packaging/copy hygiene finding.

Future tar bridge commands should exclude:

    *.db
    *.db-shm
    *.db-wal
    *.exe
    *.test.exe

in addition to the existing generated/private/build artifact excludes.

## 6. Current classification

WSL Debian setup:

    PASS

Toolchain availability:

    PASS

Workspace layout:

    PASS

Current working repo copy:

    PASS with artifact hygiene warning

Ready for full WSL tests:

    not yet in this checkpoint

The next rung should remove the copied DB artifact, re-check pre-test artifact hygiene, then run WSL package-level tests.

## 7. What WSL can validate

WSL Debian quick-portability can validate:

    Go tests on Linux paths;
    Rust sidecar build/test behavior on Linux;
    Unix path separator behavior;
    local 127.0.0.1 Cypher lifecycle in Linux userspace;
    which Windows PowerShell scripts need Linux equivalents;
    whether a future cross-platform validation runner is needed.

## 8. What WSL does not validate

WSL Debian does not validate:

    real Debian homelab deployability;
    systemd service behavior;
    cloudflared ingress;
    firewall behavior;
    long-running operations;
    storage layout on the real host;
    coexistence with Gitea, WordPress, MariaDB, nginx, and existing services;
    public deployment readiness.

## 9. Next test sequence

Recommended v0.3.9 WSL quick-portability test order:

    1. Remove copied generated artifacts from the WSL workspace.
    2. Run a pre-test artifact scan.
    3. Run carbonstack-cypher package tests:
           go test ./... -count=1
    4. Run carbonstack-comms package tests:
           go test ./... -count=1
    5. If Go tests expose Rust sidecar build/test needs, run cargo tests directly from the sidecar directory.
    6. Do not rely on PowerShell wrappers for the first Linux pass.
    7. Classify the wrapper gap:
           bash bridge;
           PowerShell Core;
           future Go validation runner.

## 10. Long-term validation runner decision

Long-term preference:

    small Go validation runner

Reason:

    CarbonStack already depends on Go;
    a Go runner can reduce divergent PowerShell and bash validation logic;
    one cross-platform validation entrypoint would reduce release friction;
    Windows and Linux release paths should converge before shell-specific tech debt compounds.

Short-term bridge:

    direct Go/Rust commands first;
    bash only where needed;
    avoid claiming Linux public testing support until validation records exist.

## 11. Current nonclaims

This checkpoint does not validate:

    Linux portability;
    Debian readiness;
    Debian deployability;
    WSL compatibility beyond setup/toolchain availability;
    systemd behavior;
    cloudflared routing;
    public ingress;
    production readiness;
    production E2EE;
    hostile-server safety;
    metadata privacy;
    secure vault/storage;
    external audit;
    certification.

## 12. Next work

Proceed to v0.3.9 WSL Debian quick-portability testing.

Do not touch the real Debian homelab yet.

Do not expose Cypher through cloudflared.

Do not create systemd services.

Do not move to runtime Comms UX until the validation substrate is less platform-fragile.
