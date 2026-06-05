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

require_clean_repo() {
  local repo_path="$1"
  local repo_name="$2"

  git -C "$repo_path" diff --quiet || fail "$repo_name has unstaged changes"
  git -C "$repo_path" diff --cached --quiet || fail "$repo_name has staged but uncommitted changes"

  if [ -n "$(git -C "$repo_path" status --short)" ]; then
    fail "$repo_name has untracked or dirty files; clean before staging release package"
  fi
}

repo_root="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
umbrella_root="$(cd "$repo_root/.." && pwd)"
stage_root="${1:-/tmp/carbonstack-v0.5.0-rehearsal/stage}"
package_root="$stage_root/package"
release_dir="$package_root/release"

require_cmd git
require_cmd tar
require_cmd python3

for repo in carbonstack carbonstack-comms carbonstack-cypher; do
  [ -d "$umbrella_root/$repo/.git" ] || fail "missing live git checkout: $umbrella_root/$repo"
  require_clean_repo "$umbrella_root/$repo" "$repo"
done

if [ -d "$umbrella_root/carbonstack-os" ]; then
  printf 'note: carbonstack-os exists in umbrella but is intentionally excluded from this runnable package\n'
fi

log "Prepare stage root"
rm -rf "$stage_root"
mkdir -p "$package_root"
mkdir -p "$release_dir"

log "Copy tracked repository contents with git archive"
for repo in carbonstack carbonstack-comms carbonstack-cypher; do
  printf 'copy_repo: %s\n' "$repo"
  git -C "$umbrella_root/$repo" archive --format=tar --prefix="$repo/" HEAD | tar -x -C "$package_root"
done

log "Write release metadata skeleton"

cp "$umbrella_root/carbonstack/LICENSE" "$release_dir/LICENSE"

python3 - "$package_root" "$umbrella_root" <<'PYMETA'
import json
import subprocess
import sys
from datetime import datetime, timezone
from pathlib import Path

package_root = Path(sys.argv[1]).resolve()
umbrella_root = Path(sys.argv[2]).resolve()
release_dir = package_root / "release"

def git(repo, *args):
    return subprocess.check_output(["git", "-C", str(umbrella_root / repo), *args], text=True).strip()

repos = {}
for repo in ["carbonstack", "carbonstack-comms", "carbonstack-cypher"]:
    repos[repo] = {
        "commit": git(repo, "rev-parse", "HEAD"),
        "short_commit": git(repo, "rev-parse", "--short", "HEAD"),
        "summary": git(repo, "log", "-1", "--oneline"),
    }

generated_at = datetime.now(timezone.utc).isoformat()

manifest = {
    "schema_version": "carbonstack-release-manifest/v0",
    "release": {
        "target_version": "v0.5.0",
        "stage": "v0.4.21 package rehearsal staging",
        "status": "staged package skeleton only; checksums and final assets are deferred to v0.4.22/v0.4.23",
        "generated_at_utc": generated_at,
        "source_of_truth": "Gitea",
        "github_mirror_status": "secondary push mirror only",
    },
    "package_shape": {
        "included": ["carbonstack", "carbonstack-comms", "carbonstack-cypher", "release"],
        "excluded": ["carbonstack-os"],
        "carbonstack_os_note": "Excluded from runnable package; future constrained appliance OS concept only.",
    },
    "repos": repos,
    "validation_intent": [
        "v0.4.21 stages package source root and release metadata skeleton",
        "v0.4.22 writes checksums, archives, fresh-extracts, verifies checksums, and runs full validation",
        "v0.4.23 finalizes LogDoc/asset generation",
        "v0.5.0 cuts minor epoch pre-release",
    ],
    "nonclaims": [
        "not production ready",
        "not production E2EE",
        "not local-backbone",
        "not deployment",
        "not general-public usable software",
        "not PQ or quantum-safe messaging",
        "not Android ready",
        "not CarbonStackOS ready",
        "not externally audited or certified",
    ],
}

(release_dir / "manifest.json").write_text(json.dumps(manifest, indent=2) + "\n", encoding="utf-8")

validation_freeze_lines = [
    "# CarbonStack v0.5.0 Validation Freeze Skeleton",
    "",
    "Status: v0.4.21 staged package skeleton",
    "Target release: v0.5.0 minor epoch pre-release",
    f"Generated UTC: {generated_at}",
    "",
    "This file is staged for release rehearsal continuity.",
    "",
    "v0.4.21 purpose:",
    "",
    "    stage carbonstack, carbonstack-comms, carbonstack-cypher, and release metadata",
    "    exclude carbonstack-os from the runnable package",
    "    preserve v0.4.0-style package shape",
    "    defer checksums, archive creation, fresh extraction validation, and final release assets",
    "",
    "Frozen heads for this staging run:",
    "",
    f"    carbonstack        {repos['carbonstack']['summary']}",
    f"    carbonstack-comms  {repos['carbonstack-comms']['summary']}",
    f"    carbonstack-cypher {repos['carbonstack-cypher']['summary']}",
    "",
    "Boundary:",
    "",
    "    This is not the final v0.5.0 release.",
    "    This is not production readiness.",
    "    This is not local-backbone.",
    "    This is not release-package validation yet because checksums and fresh extraction validation are deferred.",
    "",
]
(release_dir / "validation-freeze.md").write_text("\n".join(validation_freeze_lines), encoding="utf-8")

testing_runbook_lines = [
    "# CarbonStack v0.5.0 Testing Runbook Skeleton",
    "",
    "Status: v0.4.21 staged package skeleton",
    "Target release: v0.5.0 minor epoch pre-release",
    "",
    "This runbook is a skeleton for the v0.5.0 release rehearsal.",
    "",
    "Current v0.4.21 staging command:",
    "",
    "    carbonstack/scripts/stage-v0.5.0-package.sh",
    "",
    "Expected v0.4.22 rehearsal flow:",
    "",
    "    cd <staged-package-root>/carbonstack/tools/carbonstack-validate",
    "    go run . --profile write-checksums --root <staged-package-root>",
    "    go run . --profile verify-checksums --root <staged-package-root>",
    "",
    "    archive the staged package without running release-snapshot inside the archive source root",
    "",
    "    extract archive into a fresh throwaway validation root",
    "",
    "    cd <fresh-extraction>/package/carbonstack/tools/carbonstack-validate",
    "    go run . --profile verify-checksums --root <fresh-extraction>/package",
    "    go run . --profile full --root <fresh-extraction>/package --clean-generated",
    "",
    "Boundary:",
    "",
    "    full remains release-snapshot followed by local-cypher.",
    "    dev-runtime-openmls and dev-runtime-openmls-wrappers remain live-umbrella-only and are not included in full.",
    "    This is not deployment.",
    "    This is not local-backbone.",
    "    This is not production security proof.",
    "",
]
(release_dir / "testing-runbook.md").write_text("\n".join(testing_runbook_lines), encoding="utf-8")

release_notes_lines = [
    "# CarbonStack v0.5.0 Release Notes Draft Placeholder",
    "",
    "This file is a staging placeholder for v0.4.21.",
    "",
    "Final release notes should be formulated in v0.4.22 using the v0.4.0 release as the continuity reference.",
    "",
    "Expected framing:",
    "",
    "    CarbonStack v0.5.0 minor epoch pre-release",
    "    accumulated v0.4.x runtime, wrapper, runner, registry, and validation work",
    "    Gitea official pre-release",
    "    GitHub mirrors remain secondary push mirrors",
    "    not general-public usable software",
    "    not production secure",
    "    not local-backbone",
    "    not v1.0.0",
    "    not PQ/state/vault implementation",
    "",
]
(release_dir / "release-notes-draft.md").write_text("\n".join(release_notes_lines), encoding="utf-8")
PYMETA

log "Stage summary"
printf 'stage_root: %s\n' "$stage_root"
printf 'package_root: %s\n' "$package_root"
printf 'release_dir: %s\n' "$release_dir"
printf 'included: carbonstack carbonstack-comms carbonstack-cypher release\n'
printf 'excluded: carbonstack-os\n'
printf 'checksums: deferred to v0.4.22\n'
printf 'archive/assets: deferred to v0.4.22/v0.4.23\n'

log "Package skeleton file check"
for path in \
  "$package_root/carbonstack/README.md" \
  "$package_root/carbonstack-comms/README.md" \
  "$package_root/carbonstack-cypher/README.md" \
  "$release_dir/manifest.json" \
  "$release_dir/validation-freeze.md" \
  "$release_dir/testing-runbook.md" \
  "$release_dir/release-notes-draft.md" \
  "$release_dir/LICENSE"
do
  if [ -f "$path" ]; then
    printf 'OK: %s\n' "$path"
  else
    fail "missing staged file: $path"
  fi
done

log "Forbidden generated/private/build artifact check"
hits="$(find "$package_root" \
  \( -path "*/.git" -o \
     -path "*/target" -o \
     -path "*/.carbonstack-openmls-sidecar-state" -o \
     -name "*.db" -o \
     -name "*.db-shm" -o \
     -name "*.db-wal" -o \
     -name "*.exe" -o \
     -name "*.test.exe" \) \
  -print)"

if [ -n "$hits" ]; then
  printf '%s\n' "$hits"
  fail "forbidden generated/private/build artifact found in staged package"
fi

printf '\nPASS: v0.5.0 package skeleton staged\n'
