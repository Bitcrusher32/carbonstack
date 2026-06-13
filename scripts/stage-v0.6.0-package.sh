#!/usr/bin/env bash
set -euo pipefail

target_version="v0.6.0"
release_title="State/UX Boundary Validation Pre-Release"

script_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
carbonstack_dir="$(cd "$script_dir/.." && pwd)"
umbrella_root="$(cd "$carbonstack_dir/.." && pwd)"

stage_root="${1:-/tmp/carbonstack-v0.6.0-stage}"
package_name="carbonstack-v0.6.0-package"
package_root="$stage_root/$package_name"
archive_path="$stage_root/$package_name.tar.gz"

comms_dir="$umbrella_root/carbonstack-comms"
cypher_dir="$umbrella_root/carbonstack-cypher"
os_dir="$umbrella_root/carbonstack-os"

echo "===== CarbonStack v0.6.0 package staging ====="
echo "target_version: $target_version"
echo "release_title:  $release_title"
echo "umbrella_root:  $umbrella_root"
echo "stage_root:     $stage_root"
echo "package_root:   $package_root"
echo "archive_path:   $archive_path"
echo

require_repo() {
  local path="$1"
  local name="$2"
  if [ ! -d "$path/.git" ]; then
    echo "MISSING REPO: $name at $path" >&2
    exit 1
  fi
}

require_no_uncommitted_tracked_changes() {
  local path="$1"
  local name="$2"
  echo "checking tracked cleanliness: $name"
  if ! git -C "$path" diff --quiet; then
    echo "DIRTY TRACKED WORKTREE: $name" >&2
    git -C "$path" status --short
    exit 1
  fi
  if ! git -C "$path" diff --cached --quiet; then
    echo "DIRTY INDEX: $name" >&2
    git -C "$path" status --short
    exit 1
  fi
}

safe_remove_stage_root() {
  case "$stage_root" in
    /tmp/carbonstack-*|"$HOME"/carbonstack-*)
      rm -rf "$stage_root"
      ;;
    *)
      echo "Refusing to remove unsafe stage_root: $stage_root" >&2
      echo "Use /tmp/carbonstack-* or \$HOME/carbonstack-* for this script." >&2
      exit 1
      ;;
  esac
}

archive_repo() {
  local src="$1"
  local dest="$2"
  local name="$3"
  echo "archiving $name HEAD -> $dest"
  mkdir -p "$dest"
  git -C "$src" archive --format=tar HEAD | tar -x -C "$dest"
}

require_repo "$carbonstack_dir" "carbonstack"
require_repo "$comms_dir" "carbonstack-comms"
require_repo "$cypher_dir" "carbonstack-cypher"

if [ -d "$os_dir/.git" ]; then
  echo "carbonstack-os present, but intentionally excluded from v0.6.0 package."
fi

require_no_uncommitted_tracked_changes "$carbonstack_dir" "carbonstack"
require_no_uncommitted_tracked_changes "$comms_dir" "carbonstack-comms"
require_no_uncommitted_tracked_changes "$cypher_dir" "carbonstack-cypher"

carbonstack_head="$(git -C "$carbonstack_dir" rev-parse --short HEAD)"
comms_head="$(git -C "$comms_dir" rev-parse --short HEAD)"
cypher_head="$(git -C "$cypher_dir" rev-parse --short HEAD)"
os_head=""
if [ -d "$os_dir/.git" ]; then
  os_head="$(git -C "$os_dir" rev-parse --short HEAD)"
fi

safe_remove_stage_root
mkdir -p "$package_root"

archive_repo "$carbonstack_dir" "$package_root/carbonstack" "carbonstack"
archive_repo "$comms_dir" "$package_root/carbonstack-comms" "carbonstack-comms"
archive_repo "$cypher_dir" "$package_root/carbonstack-cypher" "carbonstack-cypher"

if [ -e "$package_root/carbonstack-os" ]; then
  echo "ERROR: carbonstack-os unexpectedly present in package root" >&2
  exit 1
fi

release_dir="$package_root/release"
mkdir -p "$release_dir"

created_utc="$(date -u +"%Y-%m-%dT%H:%M:%SZ")"

cat > "$release_dir/manifest.json" <<EOF
{
  "schema_version": "carbonstack-release-manifest/v0.2",
  "target_version": "$target_version",
  "release_title": "$release_title",
  "package_role": "staged rehearsal package",
  "created_utc": "$created_utc",
  "public_release_status": "pre-release candidate rehearsal; not tagged unless explicitly released",
  "included_repositories": {
    "carbonstack": "$carbonstack_head",
    "carbonstack-comms": "$comms_head",
    "carbonstack-cypher": "$cypher_head"
  },
  "excluded_repositories": {
    "carbonstack-os": "$os_head"
  },
  "validation_entrypoint": "carbonstack/tools/carbonstack-validate",
  "intended_validation": [
    "go run . --profile verify-checksums --root <package-root>",
    "go run . --profile full --root <package-root> --clean-generated"
  ],
  "explicit_nonclaims": [
    "not production secure messaging",
    "not production E2EE",
    "not local-backbone",
    "not hostile-server safety",
    "not metadata privacy",
    "not verified identity",
    "not secure enrollment",
    "not production vault or key storage",
    "not Android or CarbonStackOS readiness",
    "not audit or certification",
    "not general-public usability"
  ]
}
EOF

cat > "$release_dir/release-notes.md" <<EOF
# CarbonStack $target_version — $release_title

Status: staged rehearsal notes. These notes are package-rehearsal material until an explicit public release is cut.

## Scope

This package is intended to validate the v0.6.0 State/UX Boundary runway after the v0.5.63-v0.5.68B cleanup sequence.

It should demonstrate that the package contains current public surfaces, registry command-boundary artifacts, release metadata, checksums, and clean staged source roots that can run the release validation ladder from a fresh extraction.

## Validation shape

    cd <package-root>/carbonstack/tools/carbonstack-validate
    go run . --profile verify-checksums --root <package-root>
    go run . --profile full --root <package-root> --clean-generated

## Nonclaims

This package does not claim production secure messaging, production E2EE, local-backbone, hostile-server safety, metadata privacy, verified identity, secure enrollment, production vault/key storage, Android readiness, CarbonStackOS readiness, audit, certification, or general-public usability.

## Relay/OpenMLS profile boundary

relay-openmls-join-dev remains manual/dev-only for v0.6.0 and is excluded from full and release-snapshot.
EOF

cat > "$release_dir/testing-runbook.md" <<EOF
# CarbonStack $target_version Testing Runbook

Use this runbook from a freshly extracted package root.

    cd <package-root>/carbonstack/tools/carbonstack-validate
    go run . --profile verify-checksums --root <package-root>
    go run . --profile full --root <package-root> --clean-generated

Expected boundary:

- release-snapshot validates package layout, release metadata, registry/table presence, checksum verification, and strict pre-test artifact hygiene;
- full then runs local-cypher after release-snapshot;
- relay-openmls-join-dev is not part of full/release-snapshot for v0.6.0.
EOF

cat > "$release_dir/validation-freeze.md" <<EOF
# CarbonStack $target_version Validation Freeze

This staged package was assembled from tracked HEAD snapshots:

- carbonstack: $carbonstack_head
- carbonstack-comms: $comms_head
- carbonstack-cypher: $cypher_head
- carbonstack-os: excluded

Generated/local/private artifacts must not be present in this package.

Raw private LogDoc and breakpoint files are not release artifacts.
EOF

cp "$package_root/carbonstack/LICENSE" "$release_dir/LICENSE"

echo
echo "===== writing package checksums ====="
cd "$package_root/carbonstack/tools/carbonstack-validate"
go run . --profile write-checksums --root "$package_root"
go run . --profile verify-checksums --root "$package_root"

echo
echo "===== creating archive ====="
cd "$stage_root"
tar -czf "$archive_path" "$package_name"

echo
echo "STAGED PACKAGE ROOT: $package_root"
echo "STAGED PACKAGE ARCHIVE: $archive_path"
echo "CarbonStackOS excluded by design."
