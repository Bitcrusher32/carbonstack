# Go Validation Runner Implementation v0

Status: implementation / validation record
Phase: v0.3.12 runner hardening and docs integration
Previous checkpoint: docs/130-go-validation-runner-design-v0.md

## 1. Purpose

This document records the first hardening and documentation integration pass for the Go-based CarbonStack validation runner.

The runner lives at:

    tools/carbonstack-validate

It is intended to become the umbrella validation authority over time while still calling repo-local tests.

## 2. Current profiles

Implemented profiles:

    doctor
    core
    full

`full` currently aliases `core`.

This is intentional. The real full profile should wait until later v0.3.x work creates release/deployability validation surfaces that are real enough to test.

## 3. Current core behavior

The current core profile runs:

    doctor;
    pre-test artifact scan;
    targeted OpenMLS real-Cypher lifecycle test;
    full carbonstack-comms package tests;
    full carbonstack-cypher package tests;
    post-test artifact scan.

The runner streams stdout/stderr and returns nonzero on failure.

## 4. Root inference

The runner infers the umbrella root by walking upward from the current directory and detecting sibling repo layout:

    carbonstack/
    carbonstack-comms/
    carbonstack-cypher/

This avoids hard-coding a user-specific Windows path.

The runner also supports explicit root override:

    --root <path>

## 5. Toolchain reporting

The runner reports executable paths and versions for:

    go;
    rustc;
    cargo;
    sqlite3.

The runner records the current Rust floor lesson:

    Debian apt rustc 1.85.0 was too old for OpenMLS 0.8.1.
    rustup stable rustc/cargo 1.96.0 passed under WSL Debian.

The runner reports this state but does not install or mutate toolchains.

## 6. Artifact scan behavior

Artifact scans are non-destructive.

Pre-test hits may indicate source/copy hygiene issues.

Post-test hits are expected only when they stay in known generated roots.

Current classifications include:

    known-openmls-sidecar-generated-root;
    research-generated-root;
    local-go-cache-root;
    review.

The runner does not delete files.

## 7. Windows validation

The runner was validated from Windows PowerShell.

Validated profiles:

    doctor;
    core;
    full.

Windows result:

    PASS

Known Windows note:

    sqlite3 may not be available in PATH.
    This is currently a warning, not a failure.

## 8. WSL Debian validation

The runner was validated from WSL Debian.

Validated profiles:

    doctor;
    core;
    full.

WSL Debian result:

    PASS

Observed WSL toolchain:

    Go 1.24.4 linux/amd64;
    rustc 1.96.0 through rustup stable;
    cargo 1.96.0 through rustup stable;
    sqlite3 3.46.1.

WSL pre-test artifact scan was clean.

WSL post-test artifact scan found expected OpenMLS sidecar generated roots:

    internal/protocol/mls/openmls-sidecar/target
    internal/protocol/mls/openmls-sidecar/.carbonstack-openmls-sidecar-state

## 9. Interpretation

The Go validation runner is now implemented and has passed the current core validation path on both Windows and WSL Debian.

This does not yet replace public release validation.

It does establish the runner as the future umbrella validation authority candidate.

## 10. Nonclaims

Passing the runner does not prove:

    production readiness;
    production E2EE;
    hostile-server safety;
    metadata privacy;
    Debian deployability;
    systemd readiness;
    cloudflared readiness;
    real homelab validation;
    external audit;
    certification.

## 11. Next work

Recommended next work:

    validate runner behavior from clean source snapshot copies;
    preserve full as an alias for core until real release/deployability checks exist;
    avoid adding shell-specific validation logic unless it is a thin wrapper around the Go runner;
    defer real Debian homelab work until deployability and IRC-style setup work require it.
