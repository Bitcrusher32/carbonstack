#!/usr/bin/env bash
set -euo pipefail

fail() {
  printf 'FAIL: %s\n' "$1" >&2
  exit 1
}

log() {
  printf '\n==> %s\n' "$1"
}

require_cmd() {
  command -v "$1" >/dev/null 2>&1 || fail "missing required command: $1"
}

repo_root="${1:-$HOME/repos/carbonstack_umbrella/carbonstack}"
rehearsal_root="${2:-/tmp/carbonstack-v0.5.0-rehearsal}"
stage_root="$rehearsal_root/stage"
package_root="$stage_root/package"
archive_dir="$rehearsal_root/archive"
extract_root="$rehearsal_root/extract"
archive_path="$archive_dir/carbonstack-v0.5.0-package-rehearsal.tar.gz"
extracted_package_root="$extract_root/package"

require_cmd git
require_cmd tar
require_cmd python3
require_cmd go
require_cmd rustc
require_cmd cargo

[ -d "$repo_root/.git" ] || fail "repo_root is not carbonstack git checkout: $repo_root"

log "Stage package skeleton from clean live repos"
"$repo_root/scripts/stage-v0.5.0-package.sh" "$stage_root"

log "Write release checksums in staged package"
cd "$package_root/carbonstack/tools/carbonstack-validate"
go run . --profile write-checksums --root "$package_root"

log "Verify release checksums in staged package"
go run . --profile verify-checksums --root "$package_root"

log "Archive staged package"
rm -rf "$archive_dir" "$extract_root"
mkdir -p "$archive_dir" "$extract_root"
tar -C "$stage_root" -czf "$archive_path" package
printf 'archive_path: %s\n' "$archive_path"

log "Fresh extract package archive"
tar -C "$extract_root" -xzf "$archive_path"

if [ ! -d "$extracted_package_root" ]; then
  fail "fresh extraction did not create expected package root: $extracted_package_root"
fi

log "Verify release checksums from fresh extraction"
cd "$extracted_package_root/carbonstack/tools/carbonstack-validate"
go run . --profile verify-checksums --root "$extracted_package_root"

log "Run full validation from fresh extraction"
go run . --profile full --root "$extracted_package_root" --clean-generated

log "Confirm expected release files"
for path in \
  "$extracted_package_root/release/manifest.json" \
  "$extracted_package_root/release/checksums.txt" \
  "$extracted_package_root/release/validation-freeze.md" \
  "$extracted_package_root/release/testing-runbook.md" \
  "$extracted_package_root/release/release-notes-draft.md" \
  "$extracted_package_root/release/LICENSE"
do
  if [ -f "$path" ]; then
    printf 'OK: %s\n' "$path"
  else
    fail "missing expected release file after extraction: $path"
  fi
done

log "Confirm carbonstack-os excluded"
if [ -e "$extracted_package_root/carbonstack-os" ]; then
  fail "carbonstack-os unexpectedly present in runnable package"
fi
printf 'OK: carbonstack-os excluded\n'

log "Confirm final generated/private/build roots are absent"
forbidden=(
  "$extracted_package_root/carbonstack-comms/internal/protocol/mls/openmls-sidecar/.carbonstack-openmls-sidecar-state"
  "$extracted_package_root/carbonstack-comms/internal/protocol/mls/openmls-sidecar/target"
  "$extracted_package_root/carbonstack-comms/internal/protocol/mls/openmls-sidecar/provider-storage.json"
  "$extracted_package_root/carbonstack-comms/internal/protocol/mls/openmls-sidecar/signer.json"
)

found=0
for path in "${forbidden[@]}"; do
  if [ -e "$path" ]; then
    printf 'FORBIDDEN: %s\n' "$path"
    found=1
  else
    printf 'OK absent: %s\n' "$path"
  fi
done

if [ "$found" -ne 0 ]; then
  fail "forbidden generated/private/build root remained after full --clean-generated"
fi

printf '\nPASS: v0.5.0 package rehearsal completed\n'
printf 'package_root: %s\n' "$package_root"
printf 'archive_path: %s\n' "$archive_path"
printf 'extracted_package_root: %s\n' "$extracted_package_root"
