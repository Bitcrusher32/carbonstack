# LogDoc Case Study Sanitization Plan v0

Status: planning / future documentation task
Phase: v0.3.x post-release documentation

## 1. Purpose

The v0.3.0[PRIME] LogDoc is an intentionally uncompressed project handoff.

A sanitized version should be produced as a LogDoc case study and workflow document for the `carbonstack` repo.

The sanitized case study should show how the LogDoc workflow preserved project continuity across a long research, implementation, testing, release-hardening, and release-packaging ladder.

## 2. Source material

Primary source:

    CarbonStackLogDocV0.3.0_PRIME.md

The PRIME file is intentionally uncompressed.

Do not compress or rewrite it in place.

## 3. Target output

Possible target path:

    docs/logdoc-case-study-carbonstack-v0.3.0.md

The case study should be safe to publish.

It should be written as a workflow/process document, not as a private chat transcript.

## 4. Sanitization rules

Remove or replace:

    personal machine-specific sensitive paths;
    secrets;
    private keys;
    generated OpenMLS provider state;
    local DB names if sensitive;
    anything that would weaken project security if copied;
    private emotional/process details not needed for the workflow case study.

Preserve:

    project timeline;
    mistakes and recovery loops;
    validation failures and fixes;
    release-hardening decisions;
    public nonclaim discipline;
    LogDoc method lessons;
    version ladder;
    repo role boundaries;
    why the release was scoped as an experimental backbone.

## 5. Security rule

Do not publish raw `provider-storage.json`.

Do not publish raw `signer.json`.

Do not publish generated OpenMLS private/dev state.

Do not publish local runtime databases.

Do not publish anything that provides an attack path against the project or local machine.

## 6. Case study structure

Recommended structure:

1. what LogDoc is;
2. what CarbonStack needed from LogDoc;
3. project timeline from v0.0.0 to v0.3.0;
4. how checkpoints worked;
5. how blunders were preserved instead of hidden;
6. how validation gates were used;
7. how release wording/nonclaims were controlled;
8. what worked;
9. what failed;
10. what should improve in future LogDoc use.

## 7. Timing

Produce this after the v0.3.0 public release assets are verified.

Acceptable timing:

    end of v0.3.1

or:

    v0.3.2 if release verification finds issues that should be fixed first.
