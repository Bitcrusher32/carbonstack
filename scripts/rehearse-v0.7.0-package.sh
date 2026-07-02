#!/usr/bin/env bash
set -euo pipefail

script_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

stage_root="${1:-/tmp/carbonstack-v0.7.0-stage}"
rehearsal_root="${2:-/tmp/carbonstack-v0.7.0-rehearsal}"
package_name="carbonstack-v0.7.0-package"
archive_name="carbonstack-v0.7.0-cumulative-pre-alpha-engineering-boundary-pre-release.tgz"
archive_path="$stage_root/$archive_name"
extract_root="$rehearsal_root/extract"
package_root="$extract_root/$package_name"
asset_checksums_name="carbonstack-v0.7.0-asset-checksums.txt"

echo "===== CarbonStack v0.7.0 package rehearsal ====="
echo "stage_root:     $stage_root"
echo "rehearsal_root: $rehearsal_root"
echo "archive_path:   $archive_path"
echo "package_root:   $package_root"
echo

"$script_dir/stage-v0.7.0-package.sh" "$stage_root"

echo
echo "===== verifying staged release asset checksums ====="
cd "$stage_root"
sha256sum -c "$asset_checksums_name"

case "$rehearsal_root" in
  /tmp/carbonstack-*|"$HOME"/carbonstack-*)
    rm -rf "$rehearsal_root"
    ;;
  *)
    echo "Refusing to remove unsafe rehearsal_root: $rehearsal_root" >&2
    echo "Use /tmp/carbonstack-* or \$HOME/carbonstack-* for this script." >&2
    exit 1
    ;;
esac

mkdir -p "$extract_root"

echo
echo "===== extracting staged archive ====="
tar -xzf "$archive_path" -C "$extract_root"

if [ ! -d "$package_root" ]; then
  echo "MISSING extracted package root: $package_root" >&2
  exit 1
fi

if [ -e "$package_root/carbonstack-os" ]; then
  echo "ERROR: carbonstack-os unexpectedly present in extracted package" >&2
  exit 1
fi

echo
echo "===== verifying package checksums from fresh extraction ====="
cd "$package_root/carbonstack/tools/carbonstack-validate"
go run . --profile verify-checksums --root "$package_root"

echo
echo "===== running full-validate-release from fresh extraction ====="
go run . --profile full-validate-release --root "$package_root" --clean-generated

echo
echo "===== final package artifact scan ====="
cd "$package_root"

bad=0

check_absent_name() {
  local pattern="$1"
  local description="$2"
  if find . -name "$pattern" -print -quit | grep -q .; then
    echo "FORBIDDEN PRESENT: $description ($pattern)"
    find . -name "$pattern" -print
    bad=1
  else
    echo "OK ABSENT: $description ($pattern)"
  fi
}

check_absent_path_fragment() {
  local fragment="$1"
  local description="$2"
  if find . -path "*$fragment*" -print -quit | grep -q .; then
    echo "FORBIDDEN PRESENT: $description ($fragment)"
    find . -path "*$fragment*" -print
    bad=1
  else
    echo "OK ABSENT: $description ($fragment)"
  fi
}

check_absent_git_metadata() {
  if find . -name ".git" -print -quit | grep -q .; then
    echo "FORBIDDEN PRESENT: git metadata (.git)"
    find . -name ".git" -print
    bad=1
  else
    echo "OK ABSENT: git metadata (.git)"
  fi
}

check_required_file() {
  local path="$1"
  if [ -f "$path" ]; then
    echo "OK FILE: $path"
  else
    echo "MISSING FILE: $path"
    bad=1
  fi
}

check_optional_file() {
  local path="$1"
  if [ -f "$path" ]; then
    echo "OK OPTIONAL FILE PRESENT: $path"
  else
    echo "OPTIONAL FILE ABSENT: $path"
  fi
}

check_required_file "./release/manifest.json"
check_required_file "./release/checksums.txt"
check_required_file "./release/release-notes.md"
check_required_file "./release/testing-runbook.md"
check_required_file "./release/validation-freeze.md"
check_required_file "./release/LICENSE"
check_required_file "./carbonstack/registry/README.md"
check_required_file "./carbonstack/registry/commands.v0.yaml"
check_required_file "./carbonstack/registry/COMMAND_REFERENCE.v0.md"
check_optional_file "./carbonstack/registry/COMMAND_BOUNDARY_TABLE.v0.md"

check_absent_git_metadata
check_absent_path_fragment "/target" "Rust target output"
check_absent_path_fragment "/.carbonstack-openmls-sidecar-state" "OpenMLS sidecar generated state"
check_absent_name "cypher.db" "local Cypher DB"
check_absent_name "*.db-shm" "SQLite shm"
check_absent_name "*.db-wal" "SQLite wal"
check_absent_name "provider-storage.json" "provider storage"
check_absent_name "signer.json" "signer file"
check_absent_name "CarbonStackLogDoc*" "raw private LogDoc"
check_absent_name "CarbonStack_Breakpoint*" "raw private breakpoint"

if [ "$bad" -ne 0 ]; then
  echo "PACKAGE REHEARSAL FAILED"
  exit 1
fi

echo
echo "PACKAGE REHEARSAL PASSED"
echo "fresh_package_root: $package_root"
echo "stage_root: $stage_root"
echo "archive_path: $archive_path"
