#!/usr/bin/env bash
set -euo pipefail

target_version="v0.7.0"
release_title="Cumulative Pre-Alpha Engineering Boundary Pre-Release"
release_slug="cumulative-pre-alpha-engineering-boundary-pre-release"

script_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
carbonstack_dir="$(cd "$script_dir/.." && pwd)"
umbrella_root="$(cd "$carbonstack_dir/.." && pwd)"

stage_root="${1:-/tmp/carbonstack-v0.7.0-stage}"
package_name="carbonstack-v0.7.0-package"
archive_name="carbonstack-v0.7.0-${release_slug}.tgz"
archive_path="$stage_root/$archive_name"
package_root="$stage_root/$package_name"

asset_manifest_name="carbonstack-v0.7.0-release-manifest.json"
package_checksums_name="carbonstack-v0.7.0-package-checksums.txt"
asset_checksums_name="carbonstack-v0.7.0-asset-checksums.txt"
validation_freeze_name="carbonstack-v0.7.0-validation-freeze.md"
testing_runbook_name="v0.7.0-testing-runbook.md"
release_notes_name="v0.7.0-release-notes.md"

comms_dir="$umbrella_root/carbonstack-comms"
cypher_dir="$umbrella_root/carbonstack-cypher"
os_dir="$umbrella_root/carbonstack-os"

echo "===== CarbonStack v0.7.0 package staging ====="
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

command_or_unknown() {
  local cmd="$1"
  shift
  if command -v "$cmd" >/dev/null 2>&1; then
    "$cmd" "$@" 2>/dev/null || true
  else
    echo "not-found"
  fi
}

require_repo "$carbonstack_dir" "carbonstack"
require_repo "$comms_dir" "carbonstack-comms"
require_repo "$cypher_dir" "carbonstack-cypher"

if [ -d "$os_dir/.git" ]; then
  echo "carbonstack-os present, but intentionally excluded from v0.7.0 runnable package."
fi

require_no_uncommitted_tracked_changes "$carbonstack_dir" "carbonstack"
require_no_uncommitted_tracked_changes "$comms_dir" "carbonstack-comms"
require_no_uncommitted_tracked_changes "$cypher_dir" "carbonstack-cypher"

carbonstack_head="$(git -C "$carbonstack_dir" rev-parse --short HEAD)"
carbonstack_head_full="$(git -C "$carbonstack_dir" rev-parse HEAD)"
comms_head="$(git -C "$comms_dir" rev-parse --short HEAD)"
comms_head_full="$(git -C "$comms_dir" rev-parse HEAD)"
cypher_head="$(git -C "$cypher_dir" rev-parse --short HEAD)"
cypher_head_full="$(git -C "$cypher_dir" rev-parse HEAD)"
os_head=""
os_head_full=""
if [ -d "$os_dir/.git" ]; then
  os_head="$(git -C "$os_dir" rev-parse --short HEAD)"
  os_head_full="$(git -C "$os_dir" rev-parse HEAD)"
fi

go_version="$(command_or_unknown go version)"
rustc_version="$(command_or_unknown rustc --version)"
cargo_version="$(command_or_unknown cargo --version)"
sqlite_version="$(command_or_unknown sqlite3 --version)"
created_utc="$(date -u +"%Y-%m-%dT%H:%M:%SZ")"

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

cat > "$release_dir/manifest.json" <<EOF
{
  "schema_version": "carbonstack-release-manifest/v0.3",
  "target_version": "$target_version",
  "release_title": "$release_title",
  "package_role": "staged rehearsal package",
  "created_utc": "$created_utc",
  "public_release_status": "pre-release candidate rehearsal; not tagged unless explicitly released",
  "included_repositories": {
    "carbonstack": {"short": "$carbonstack_head", "full": "$carbonstack_head_full"},
    "carbonstack-comms": {"short": "$comms_head", "full": "$comms_head_full"},
    "carbonstack-cypher": {"short": "$cypher_head", "full": "$cypher_head_full"}
  },
  "excluded_repositories": {
    "carbonstack-os": {
      "short": "$os_head",
      "full": "$os_head_full",
      "reason": "future constrained-appliance OS work; excluded from runnable package"
    }
  },
  "validation_entrypoint": "carbonstack/tools/carbonstack-validate",
  "primary_validation": [
    "go run . --profile verify-checksums --root <package-root>",
    "go run . --profile full-validate-release --root <package-root> --clean-generated"
  ],
  "compatibility_validation": [
    "go run . --profile full --root <package-root> --clean-generated"
  ],
  "observed_toolchains_at_staging": {
    "go": "$go_version",
    "rustc": "$rustc_version",
    "cargo": "$cargo_version",
    "sqlite3": "$sqlite_version"
  },
  "explicit_nonclaims": [
    "not production secure messaging",
    "not production E2EE",
    "not local-backbone",
    "not hostile-server safety",
    "not metadata privacy",
    "not verified identity",
    "not secure enrollment",
    "not production vault or key storage",
    "not backup or restore",
    "not PQ or hybrid security",
    "not quantum-safe messaging",
    "not Android or CarbonStackOS readiness",
    "not deployment readiness",
    "not public ingress safety",
    "not audit or certification",
    "not general-public usability"
  ]
}
EOF

cat > "$release_dir/release-notes.md" <<EOF
# CarbonStack $target_version - $release_title

Status: staged rehearsal notes. These notes are package-rehearsal material until an explicit public release is cut.

## Scope

This package validates the candidate v0.7.0 cumulative pre-alpha engineering boundary after the v0.6.x OpenMLS message-flow integration and boundary-hardening lane.

It should demonstrate current public surfaces, command-boundary registry artifacts, generated command reference, release metadata, checksums, clean tracked source roots, and fresh-extraction validation behavior.

## Validation shape

    cd <package-root>/carbonstack/tools/carbonstack-validate
    go run . --profile verify-checksums --root <package-root>
    go run . --profile full-validate-release --root <package-root> --clean-generated

Compatibility:

    go run . --profile full --root <package-root> --clean-generated

full-validate-release is the preferred release-facing name. It is currently an exact alias of full.

Current behavior:

    full-validate-release = release-snapshot + local-cypher

release-snapshot already calls core, so full-validate-release does not call core a second time.

## Included repositories

- carbonstack: $carbonstack_head
- carbonstack-comms: $comms_head
- carbonstack-cypher: $cypher_head

## Excluded repositories

- carbonstack-os: excluded from runnable package; future constrained-appliance OS work only.

## Boundary summary

v0.7.0 candidate framing should cover the normal OpenMLS application-message path, message-send-dev and message-inbox-dev wrappers, deterministic same-state evidence, Welcome join partial-state safety, state-audit-dev, full-validate-release naming, adversarial harness contract docs, public/manual docs cleanup, and package rehearsal work.

## Nonclaims

This package does not claim production secure messaging, production E2EE, local-backbone, hostile-server safety, metadata privacy, verified identity, secure enrollment, production vault/key storage, backup/restore, PQ/hybrid security, quantum-safe messaging, Android readiness, CarbonStackOS readiness, deployment readiness, public ingress safety, audit, certification, or general-public usability.

## Gitea default source archive warning

Do not use Gitea default Source Code ZIP/TAR.GZ archives as the multi-repo validation package.

Those default archives are auto-generated single-repo snapshots of carbonstack.

Use the attached v0.7.0 package archive, manifest, checksums, validation freeze, release notes, testing runbook, and LICENSE for the intended multi-repo validation package if this rehearsal becomes a public release.
EOF

cat > "$release_dir/testing-runbook.md" <<EOF
# CarbonStack $target_version Testing Runbook

Use this runbook from a freshly extracted package root.

Primary validation:

    cd <package-root>/carbonstack/tools/carbonstack-validate
    go run . --profile verify-checksums --root <package-root>
    go run . --profile full-validate-release --root <package-root> --clean-generated

Compatibility validation:

    cd <package-root>/carbonstack/tools/carbonstack-validate
    go run . --profile full --root <package-root> --clean-generated

Expected boundary:

- verify-checksums verifies release/checksums.txt against the extracted package root.
- full-validate-release runs release-snapshot, then local-cypher.
- release-snapshot validates package layout, release metadata, registry/generated-reference presence, checksum verification, strict pre-test artifact hygiene, and core.
- core runs doctor, OpenMLS real-Cypher lifecycle validation, Comms tests, Cypher tests, and artifact scans.
- local-cypher validates temporary local Cypher lifecycle behavior.

Do not treat full-validate-release as a deployment command.

Do not run validation repeatedly in the same package root if generated artifacts from a prior run are still present. Use a fresh extraction for release verification, or use --clean-generated where documented.

Validated platform should be recorded from rehearsal output. Current expected target is Debian / WSL Debian linux/amd64.

Gitea default Source Code ZIP/TAR.GZ archives are not the intended multi-repo validation package.
EOF

cat > "$release_dir/validation-freeze.md" <<EOF
# CarbonStack $target_version Validation Freeze

This staged package was assembled from tracked HEAD snapshots:

- carbonstack: $carbonstack_head / $carbonstack_head_full
- carbonstack-comms: $comms_head / $comms_head_full
- carbonstack-cypher: $cypher_head / $cypher_head_full
- carbonstack-os: excluded

Observed staging toolchains:

- Go: $go_version
- Rust: $rustc_version
- Cargo: $cargo_version
- SQLite: $sqlite_version

Generated/local/private artifacts must not be present in this package.

Raw private LogDoc and breakpoint files are not release artifacts.

full-validate-release is the preferred release-validation name for this boundary. full remains compatibility wording.
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
echo "===== copying release assets to stage root ====="
cp "$release_dir/manifest.json" "$stage_root/$asset_manifest_name"
cp "$release_dir/checksums.txt" "$stage_root/$package_checksums_name"
cp "$release_dir/validation-freeze.md" "$stage_root/$validation_freeze_name"
cp "$release_dir/testing-runbook.md" "$stage_root/$testing_runbook_name"
cp "$release_dir/release-notes.md" "$stage_root/$release_notes_name"
cp "$release_dir/LICENSE" "$stage_root/LICENSE"

echo
echo "===== writing asset checksums ====="
cd "$stage_root"
sha256sum \
  "$archive_name" \
  "$asset_manifest_name" \
  "$package_checksums_name" \
  "$validation_freeze_name" \
  "$testing_runbook_name" \
  "$release_notes_name" \
  "LICENSE" > "$asset_checksums_name"

echo
echo "STAGED PACKAGE ROOT: $package_root"
echo "STAGED PACKAGE ARCHIVE: $archive_path"
echo "STAGED RELEASE ASSETS:"
echo "  $stage_root/$archive_name"
echo "  $stage_root/$asset_manifest_name"
echo "  $stage_root/$package_checksums_name"
echo "  $stage_root/$asset_checksums_name"
echo "  $stage_root/$validation_freeze_name"
echo "  $stage_root/$testing_runbook_name"
echo "  $stage_root/$release_notes_name"
echo "  $stage_root/LICENSE"
echo "CarbonStackOS excluded by design."
