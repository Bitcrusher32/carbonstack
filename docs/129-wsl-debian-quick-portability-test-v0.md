# WSL Debian Quick-Portability Test v0

Status: validation record / portability triage
Phase: v0.3.9 backbone portability triage
Previous checkpoint: docs/128-wsl-debian-quick-portability-setup-scout-v0.md

## 1. Purpose

This document records the v0.3.9 WSL Debian quick-portability test.

The goal was to test whether the current CarbonStack Comms/Cypher backbone can run core package-level and OpenMLS real-Cypher lifecycle validation under Debian Linux userspace through WSL.

This is not a public Linux release.

This is not Debian deployability validation.

This is not homelab validation.

## 2. Test environment

Environment:

    Debian GNU/Linux 13 trixie
    WSL2
    linux/amd64
    workspace: ~/carbonstack-wsl

Initial toolchain from Debian packages:

    git 2.47.3
    go 1.24.4 linux/amd64
    rustc 1.85.0
    cargo 1.85.0
    sqlite3 3.46.1

Corrected Rust toolchain:

    rustup stable
    rustc 1.96.0
    cargo 1.96.0

The corrected Rust toolchain must be loaded through:

    ~/.cargo/bin

or:

    . "$HOME/.cargo/env"

## 3. Initial failure

The initial WSL Comms test failed heavily in:

    carbonstack-comms/internal/protocol

Most failing tests returned:

    exit status 101

Direct sidecar execution revealed the real root cause:

    error[E0658]: use of unstable library feature `unsigned_is_multiple_of`

This occurred while compiling:

    openmls 0.8.1

The failing Rust compiler was:

    rustc 1.85.0

## 4. Root cause

The root cause was not a CarbonStack protocol failure.

The root cause was not a Cypher server failure.

The root cause was not WSL networking.

The root cause was:

    Debian apt Rust 1.85.0 was too old for OpenMLS 0.8.1.

After installing rustup stable and ensuring the shell resolved Rust/Cargo from `~/.cargo/bin`, the OpenMLS sidecar built and ran.

## 5. Toolchain repair

The fix path was:

    install rustup stable;
    load ~/.cargo/env;
    verify rustc/cargo resolve to ~/.cargo/bin;
    use rustc 1.96.0 / cargo 1.96.0;
    rerun sidecar and Go tests.

Direct sidecar probe after repair:

    cargo run -- identity-create --state-dir /tmp/carbonstack-wsl-sidecar-probe --device-label alice

Result:

    PASS

The sidecar returned valid JSON for identity creation and wrote local dev-only identity state.

## 6. Passing validation

After rustup stable was active, the following WSL Debian tests passed.

Targeted OpenMLS real-Cypher lifecycle:

    go test ./internal/protocol -run TestOpenMLSSidecarFullLifecycleRelayThroughRealCypherServer -count=1 -timeout 300s -v

Result:

    PASS

Full Comms protocol package:

    go test ./internal/protocol -count=1 -timeout 600s -v

Result:

    PASS

Full Comms package suite:

    go test ./... -count=1 -timeout 600s

Result:

    PASS

Cypher package suite:

    go test ./... -count=1

Result:

    PASS

## 7. Post-test artifacts

Post-test generated artifacts appeared in expected locations.

Observed categories:

    carbonstack-comms/internal/protocol/mls/openmls-sidecar/target
    carbonstack-comms/internal/protocol/mls/openmls-sidecar/.carbonstack-openmls-sidecar-state
    provider-storage.json under OpenMLS sidecar dev state
    signer.json under OpenMLS sidecar dev state

These are expected post-validation artifacts.

No unexpected scattered DBs or executables were reported in the final post-test scan.

## 8. Interpretation

This is a strong WSL Debian quick-portability pass.

It validates that the current Comms/Cypher OpenMLS backbone core can run under Debian WSL after correcting the Rust toolchain floor.

The main remaining portability gap is not the Go/Rust core.

The main remaining portability gap is validation orchestration:

    Windows PowerShell wrappers exist;
    Linux direct commands work;
    a cross-platform validation runner is not implemented yet.

## 9. Toolchain floor finding

The project should record a practical Rust floor for OpenMLS sidecar work.

Current finding:

    rustc 1.85.0 is too old for OpenMLS 0.8.1.
    rustc 1.96.0 works in WSL Debian.

Future docs should avoid saying only "install rustc from apt" for Debian.

Future Debian/WSL setup should prefer:

    rustup stable

unless a distribution package is known to satisfy the required Rust version.

## 10. Go-based full-test wrapper direction

Long-term direction:

    implement a Go-based full-test wrapper

Reason:

    CarbonStack already requires Go;
    Go can provide a cross-platform validation entrypoint;
    Windows and Linux release validation should converge;
    PowerShell and bash should become thin convenience wrappers, not the validation authority.

Potential future command shape:

    go run ./tools/carbonstack-validate --profile full
    go run ./tools/carbonstack-validate --profile openmls-backbone
    go run ./tools/carbonstack-validate --profile release-snapshot

This should be planned before platform-specific release scripts become compounding tech debt.

## 11. Authoritative environment consideration

WSL Debian is now a plausible future authoritative core-validation environment.

Reasons:

    tests are fast;
    no Avast/Windows Defender-style interruption was observed;
    Linux filesystem/toolchain behavior is closer to future Debian CLI deployments;
    the current homelab also runs CLI-only Debian.

Do not switch authority yet.

Recommended sequence:

    record WSL Debian pass;
    design Go-based full-test wrapper;
    implement cross-platform validation runner;
    then consider making WSL Debian the preferred core-validation environment.

Windows remains the current public release platform boundary until a later release explicitly validates another path.

## 12. What this validates

This validates:

    WSL Debian can run the targeted OpenMLS real-Cypher lifecycle test after rustup stable;
    WSL Debian can run full carbonstack-comms/internal/protocol tests after rustup stable;
    WSL Debian can run full carbonstack-comms package tests after rustup stable;
    WSL Debian can run full carbonstack-cypher package tests;
    current core Go/Rust backbone behavior is not obviously Windows-only.

## 13. What this does not validate

This does not validate:

    public Linux release support;
    Debian homelab deployability;
    real Debian host operations;
    systemd behavior;
    cloudflared routing;
    firewall posture;
    long-running service behavior;
    public ingress;
    production readiness;
    production E2EE;
    hostile-server safety;
    metadata privacy;
    secure vault/storage;
    external audit;
    certification.

## 14. Next work

Recommended next work:

    v0.3.10:
        design the Go-based cross-platform full-test wrapper.

    v0.3.10+:
        implement the smallest useful runner profile.

    later:
        validate current source snapshots through the runner on Windows and WSL Debian.

Do not touch the real Debian homelab yet.

Do not expose Cypher through cloudflared.

Do not create a systemd service.

Do not move to runtime Comms UX until the validation substrate is less platform-fragile.
