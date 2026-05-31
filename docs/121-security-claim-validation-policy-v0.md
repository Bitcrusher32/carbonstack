# Security Claim Validation Policy v0

Status: policy / future gate
Phase: v0.3.x governance, before future adversarial-validation epoch

## 1. Purpose

CarbonStack must not make safety or security claims without a validation record.

This policy defines the minimum documentation standard for future security, safety, deployment-hardening, hostile-server, or penetration-testing claims.

This policy does not certify CarbonStack.

This policy does not replace an external audit.

## 2. Core rule

No security claim without a validation record.

No validation record without a stated non-scope.

No passing test may be generalized beyond what it actually tested.

## 3. Required validation record fields

Each claim-supporting validation record must include:

1. claim being evaluated;
2. tested version, commit, release, or deployment;
3. tested environment;
4. methodology;
5. tools or manual techniques used;
6. assumptions;
7. result;
8. what the result validates;
9. what the result does not validate;
10. residual risks;
11. claim decision:
    - allowed;
    - narrowed;
    - deferred;
    - rejected.

## 4. Methodology requirement

Methodology can be simple, but it must exist.

A record may reference:

    a CarbonStack docs methodology section;
    a test matrix;
    a pentest checklist;
    a threat-model document;
    tool output;
    manual reproduction steps.

The methodology must be specific enough to understand what was actually tested.

## 5. Pen-test / adversarial-validation gate

Before any meaningful security/safety release claim, CarbonStack should complete a dedicated adversarial-validation minor epoch.

This future epoch may include:

    self-directed penetration testing;
    hostile-server behavior tests;
    abuse-resistance harnesses;
    deployment exposure tests;
    AI-assisted adversarial probing;
    limited external attempts against scoped temporary deployments.

This is not an external audit unless a qualified external audit actually happens.

## 6. Forbidden unsupported claims

Do not claim:

    secure messenger;
    production E2EE;
    production readiness;
    hostile-server safety;
    metadata privacy;
    trustless operation;
    drop-in replacement;
    certified security;
    external audit;
    audit-ready status.

Unless each exact claim has direct validation evidence and the claim wording is narrowed to the tested context.

## 7. Allowed narrow claim shape

Good claim shape:

    In tested build X, under deployment shape Y, test method Z validated behavior A for cases B and C.

Bad claim shape:

    CarbonStack is secure.

Good claim shape:

    In the tested local v0.x harness, Comms acknowledged relay envelopes only after the documented OpenMLS sidecar consume command succeeded.

Bad claim shape:

    CarbonStack guarantees safe delivery semantics.

## 8. Relationship to future external audit

Self-testing can improve quality and narrow claims.

Self-testing is not an external audit.

AI-assisted testing is not an external audit.

Community attempts are not an external audit unless they are scoped, documented, performed by identifiable qualified reviewers, and reported as an audit.

## 9. Current v0.3.0 status

The v0.3.0 release is an experimental backbone release.

It does not make production security claims.

It does not make hostile-server safety claims.

It does not make metadata privacy claims.

It does not claim external audit or certification.
