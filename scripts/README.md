# CarbonStack scripts

## v0.5.0 release helper system

Current helper split:

    scripts/stage-v0.5.0-package.sh
    scripts/rehearse-v0.5.0-package.sh

`stage-v0.5.0-package.sh` stages a clean v0.5.0 package skeleton from tracked `carbonstack`, `carbonstack-comms`, and `carbonstack-cypher` files, then writes release metadata skeleton files. It intentionally excludes `carbonstack-os`.

`rehearse-v0.5.0-package.sh` runs the release rehearsal flow from outside the package root discipline:

    stage package skeleton
    write release/checksums.txt
    verify checksums in staged package
    archive staged package
    fresh extract archive
    verify checksums from fresh extraction
    run full from fresh extraction with --clean-generated

Boundary:

    These scripts do not cut the release.
    These scripts do not upload release assets.
    These scripts do not make runtime OpenMLS profiles part of full.
    These scripts do not create local-backbone.
    These scripts do not start PQ/state/vault implementation.
