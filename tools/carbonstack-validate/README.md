# CarbonStack Validate

Status: experimental validation runner
Phase: v0.3.11 implementation scaffold

This is the first Go-based umbrella validation runner for CarbonStack.

It is intended to replace shell-specific umbrella validation over time while still calling repo-local tests.

## Profiles

### doctor

Reports environment, inferred repo layout, required paths, and toolchain versions.

    go run . --profile doctor

### core

Runs the current core validation path:

    doctor
    pre-test artifact scan
    targeted OpenMLS real-Cypher lifecycle test
    full carbonstack-comms package tests
    full carbonstack-cypher package tests
    post-test artifact scan

    go run . --profile core

### full

Currently aliases `core`.

    go run . --profile full

## Expected layout

The runner expects sibling repos:

    carbonstack/
    carbonstack-comms/
    carbonstack-cypher/

Run from:

    carbonstack/tools/carbonstack-validate

or pass an explicit umbrella root later when supported by workflow.

## Boundaries

This runner does not prove production readiness, production E2EE, hostile-server safety, metadata privacy, Debian deployability, systemd readiness, cloudflared readiness, audit, or certification.

It does not install dependencies, delete artifacts, package releases, publish releases, configure services, or deploy anything.
