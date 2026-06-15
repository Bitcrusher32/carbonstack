#!/usr/bin/env python3
"""
Render CarbonStack's generated dev/operator command reference.

Input:
  registry/commands.v0.yaml

Output:
  registry/COMMAND_REFERENCE.v0.md

This renderer is intentionally dependency-free. It parses the constrained registry
shape CarbonStack currently uses rather than requiring PyYAML on release/tester
machines.
"""

from __future__ import annotations

import argparse
import re
import sys
from collections import defaultdict
from pathlib import Path


SECTION_ORDER = [
    "01 Release/package validation and package-helper profiles",
    "02 Diagnostic, core, local, and registry inspection profiles",
    "03 Live-dev runtime validation profiles",
    "04 Recommended dev/pre-alpha normal-message wrappers",
    "05 Lower-level direct OpenMLS message proof commands",
    "06 Relay onboarding and artifact commands",
    "07 OpenMLS bootstrap, identity, and conversation commands",
    "08 Comms state, account, device, and trust commands",
    "09 Legacy stub-era and continuity commands",
    "10 Historical scripts and smoke helpers",
    "11 Internal OpenMLS sidecar provider commands",
    "12 Cypher server and HTTP API surfaces",
    "13 Future or unsupported placeholders",
    "14 Other registered surfaces",
]


def clean_scalar(value: str) -> str:
    value = value.strip()
    if len(value) >= 2 and ((value[0] == value[-1] == '"') or (value[0] == value[-1] == "'")):
        return value[1:-1]
    return value


def split_entries(text: str) -> list[str]:
    blocks = re.split(r"(?m)^(?=  - id: )", text)
    return [
        block.rstrip() + "\n"
        for block in blocks
        if re.search(r"(?m)^  - id:\s*([^\s]+)", block)
    ]


def parse_block(block: str) -> dict[str, str]:
    entry: dict[str, str] = {"raw": block}
    match = re.search(r"(?m)^  - id:\s*([^\s]+)", block)
    entry["id"] = match.group(1) if match else "<missing-id>"

    lines = block.splitlines()
    index = 0

    while index < len(lines):
        line = lines[index]
        match = re.match(r"^    ([A-Za-z0-9_]+):(?:\s*(.*))?$", line)
        if not match:
            index += 1
            continue

        key, value = match.group(1), match.group(2) or ""
        if value.strip():
            entry[key] = clean_scalar(value)
            index += 1
            continue

        nested: list[str] = []
        index += 1
        while index < len(lines):
            next_line = lines[index]
            if re.match(r"^    [A-Za-z0-9_]+:", next_line) or re.match(r"^  - id:", next_line):
                break
            nested.append(next_line)
            index += 1

        entry[f"raw_{key}"] = "\n".join(nested).rstrip()

    return entry


def simple_list(entry: dict[str, str], key: str) -> list[str]:
    raw = entry.get(f"raw_{key}", "")
    values: list[str] = []
    for line in raw.splitlines():
        match = re.match(r"^\s*-\s*(.+?)\s*$", line)
        if match and not re.match(r"^[A-Za-z0-9_]+:", match.group(1)):
            values.append(clean_scalar(match.group(1)))
    return values


def map_list(entry: dict[str, str], key: str) -> list[dict[str, str]]:
    raw = entry.get(f"raw_{key}", "")
    items: list[dict[str, str]] = []
    current: dict[str, str] | None = None

    for line in raw.splitlines():
        match = re.match(r"^\s*-\s*([A-Za-z0-9_]+):\s*(.*?)\s*$", line)
        if match:
            if current:
                items.append(current)
            current = {match.group(1): clean_scalar(match.group(2))}
            continue

        match = re.match(r"^\s+([A-Za-z0-9_]+):\s*(.*?)\s*$", line)
        if match and current is not None:
            current[match.group(1)] = clean_scalar(match.group(2))

    if current:
        items.append(current)

    return items


def group_for(entry: dict[str, str]) -> str:
    ident = entry["id"]
    kind = entry.get("kind", "")
    maturity = entry.get("maturity", "")

    # Scripts are classified before generic legacy and Comms state/trust rows.
    # This prevents legacy PowerShell and smoke helpers from being promoted as
    # ordinary legacy commands or Comms trust commands.
    if ".script." in ident or kind == "script":
        return "10 Historical scripts and smoke helpers"

    if ident in {
        "runner.full",
        "runner.release-snapshot",
        "runner.verify-checksums",
        "runner.write-checksums",
    }:
        return "01 Release/package validation and package-helper profiles"

    if ident in {"runner.doctor", "runner.core", "runner.local-cypher", "runner.registry-lookup"}:
        return "02 Diagnostic, core, local, and registry inspection profiles"

    if ident in {
        "runner.integrated-runtime-dev",
        "runner.dev-runtime-openmls",
        "runner.dev-runtime-openmls-wrappers",
        "runner.relay-openmls-join-dev",
    }:
        return "03 Live-dev runtime validation profiles"

    if ident in {"comms.message-send-dev", "comms.message-inbox-dev"}:
        return "04 Recommended dev/pre-alpha normal-message wrappers"

    if ident in {"comms.openmls-send-dev", "comms.openmls-inbox-dev"}:
        return "05 Lower-level direct OpenMLS message proof commands"

    if ident.startswith("comms.openmls-relay-"):
        return "06 Relay onboarding and artifact commands"

    if ident.startswith("comms.openmls-"):
        return "07 OpenMLS bootstrap, identity, and conversation commands"

    if ident.startswith("comms.") and any(
        token in ident
        for token in [
            "init",
            "claim-invite",
            "register-device",
            "list-devices",
            "fingerprint",
            "trust",
            "revoke",
            "account",
            "device",
        ]
    ):
        return "08 Comms state, account, device, and trust commands"

    if ident in {"comms.send", "comms.inbox", "comms.ack"} or maturity == "legacy":
        return "09 Legacy stub-era and continuity commands"

    if ident.startswith("sidecar.") or kind == "sidecar-cli":
        if "state-checkpoint" in ident or "state-load-check" in ident or maturity in {"future", "unsupported"}:
            return "13 Future or unsupported placeholders"
        return "11 Internal OpenMLS sidecar provider commands"

    if ident.startswith("cypher.") or kind == "api-surface":
        return "12 Cypher server and HTTP API surfaces"

    return "14 Other registered surfaces"


def field(entry: dict[str, str], key: str) -> str:
    value = str(entry.get(key, "")).replace("\n", " ").strip()
    return value if value else "Not recorded in registry."


def render_flag_list(title: str, items: list[dict[str, str]]) -> list[str]:
    if not items:
        return [f"- **{title}:** Not recorded in registry."]

    output = [f"- **{title}:**"]
    for item in items:
        flag = item.get("flag") or item.get("name") or item.get("env") or item.get("key") or "<unknown>"
        meaning = item.get("meaning", "").strip()
        boundary = item.get("boundary", "").strip()
        line = f"  - `{flag}`"
        if meaning:
            line += f" — {meaning}"
        if boundary:
            line += f" Boundary: {boundary}"
        output.append(line)

    return output


def render_simple_list(title: str, values: list[str]) -> list[str]:
    if not values:
        return [f"- **{title}:** Not recorded in registry."]

    output = [f"- **{title}:**"]
    for value in values:
        output.append(f"  - {value}")
    return output


def validate(entries: list[dict[str, str]], included: list[str]) -> list[str]:
    id_set = {entry["id"] for entry in entries}
    included_counts: dict[str, int] = defaultdict(int)

    for ident in included:
        included_counts[ident] += 1

    missing_from_render = sorted(id_set - set(included))
    duplicates = sorted(ident for ident, count in included_counts.items() if count != 1)

    risk_issues: list[str] = []

    for entry in entries:
        ident = entry["id"]
        kind = entry.get("kind", "")
        section = group_for(entry)
        nonclaims = simple_list(entry, "nonclaims")

        if not nonclaims:
            risk_issues.append(f"{ident}: missing nonclaims")
        if kind in {"sidecar-cli", "api-surface", "script"} and not entry.get("validation_surface"):
            risk_issues.append(f"{ident}: missing validation_surface for {kind}")
        if (kind == "script" or ".script." in ident) and "Historical scripts" not in section:
            risk_issues.append(f"{ident}: script row not in historical scripts section")
        if kind == "sidecar-cli" and "Internal OpenMLS sidecar" not in section and "Future" not in section:
            risk_issues.append(f"{ident}: sidecar row not in internal/future section")
        if kind == "api-surface" and "Cypher server and HTTP API" not in section:
            risk_issues.append(f"{ident}: API row not in Cypher/API section")
        if ident == "runner.full" and "Release/package" not in section:
            risk_issues.append("runner.full not in release/package section")
        if ident == "runner.release-snapshot" and "Release/package" not in section:
            risk_issues.append("runner.release-snapshot not in release/package section")
        if ident == "runner.integrated-runtime-dev" and "Live-dev runtime" not in section:
            risk_issues.append("runner.integrated-runtime-dev not in live-dev runtime section")

    issues: list[str] = []
    issues.extend(f"missing from render: {ident}" for ident in missing_from_render)
    issues.extend(f"duplicate render: {ident}" for ident in duplicates)
    issues.extend(risk_issues)
    return issues


def render(entries: list[dict[str, str]]) -> tuple[str, list[str]]:
    groups: dict[str, list[dict[str, str]]] = defaultdict(list)
    for entry in entries:
        groups[group_for(entry)].append(entry)

    lines: list[str] = []
    lines.append("# CarbonStack Command Reference v0")
    lines.append("")
    lines.append("Status: **generated dev/operator command reference**")
    lines.append("")
    lines.append("Generated from `registry/commands.v0.yaml` by `tools/registry/render-command-reference.py`.")
    lines.append("")
    lines.append("Do not hand-edit this file. Update the registry source and rerun the renderer.")
    lines.append("")
    lines.append("Boundary:")
    lines.append("")
    lines.append("- Registry presence is classification, not promotion.")
    lines.append("- This reference is dev/operator-facing, not general-public UX documentation.")
    lines.append("- This is not a production security claim.")
    lines.append("- This is not a man-page set.")
    lines.append("- `full`, `release-snapshot`, and `integrated-runtime-dev` are distinct and must not be merged.")
    lines.append("- Internal sidecar commands, Cypher API surfaces, and legacy scripts are documented for boundary clarity, not promoted as user-facing commands.")
    lines.append("")
    lines.append(f"Registry entry count: **{len(entries)}**")
    lines.append("")

    included: list[str] = []

    for section in SECTION_ORDER:
        group_entries = sorted(groups.get(section, []), key=lambda item: item["id"])
        if not group_entries:
            continue

        display_section = re.sub(r"^\d+\s+", "", section)
        lines.append(f"## {display_section}")
        lines.append("")
        lines.append(f"Entries in this section: **{len(group_entries)}**")
        lines.append("")

        for entry in group_entries:
            included.append(entry["id"])
            lines.append(f"### `{entry['id']}`")
            lines.append("")
            lines.append(f"- **Command:** `{field(entry, 'command')}`")
            lines.append(f"- **Repo:** `{field(entry, 'repo')}`")
            lines.append(f"- **Component:** `{field(entry, 'component')}`")
            lines.append(f"- **Kind:** `{field(entry, 'kind')}`")
            lines.append(f"- **Audience:** `{field(entry, 'audience')}`")
            lines.append(f"- **Maturity:** `{field(entry, 'maturity')}`")
            if entry.get("lifecycle_status"):
                lines.append(f"- **Lifecycle status:** `{field(entry, 'lifecycle_status')}`")
            lines.append(f"- **Introduced in:** `{field(entry, 'introduced_in')}`")
            lines.append(f"- **Source path:** `{field(entry, 'source_path')}`")
            lines.append(f"- **Validation surface:** {field(entry, 'validation_surface')}")
            lines.append(f"- **Front README candidate:** `{field(entry, 'include_in_front_readme')}`")
            lines.append("")
            lines.append(f"**What it does:** {field(entry, 'short_help')}")
            lines.append("")
            lines.append(f"**Why it exists:** {field(entry, 'why_exists')}")
            lines.append("")

            for scalar_key, label in [
                ("wrapped_by", "Wrapped by"),
                ("replaces", "Replaces"),
                ("replacement", "Replacement"),
                ("boundary", "Boundary note"),
                ("example", "Example"),
            ]:
                if entry.get(scalar_key):
                    value = field(entry, scalar_key)
                    if scalar_key == "example":
                        lines.append(f"- **{label}:** `{value}`")
                    else:
                        lines.append(f"- **{label}:** {value}")

            for line in render_flag_list("Required flags", map_list(entry, "required_flags")):
                lines.append(line)
            for line in render_flag_list("Optional flags", map_list(entry, "optional_flags")):
                lines.append(line)
            for line in render_flag_list("Environment", map_list(entry, "environment")):
                lines.append(line)
            for line in render_simple_list("Related registry rows", simple_list(entry, "related")):
                lines.append(line)

            lines.append("- **Not claims:**")
            nonclaims = simple_list(entry, "nonclaims")
            if nonclaims:
                for nonclaim in nonclaims:
                    lines.append(f"  - {nonclaim}")
            else:
                lines.append("  - WARNING: no nonclaims recorded in registry.")

            lines.append("")

    return "\n".join(lines).rstrip() + "\n", included


def load_entries(registry_path: Path) -> list[dict[str, str]]:
    return [parse_block(block) for block in split_entries(registry_path.read_text(encoding="utf-8", errors="replace"))]


def main(argv: list[str]) -> int:
    parser = argparse.ArgumentParser(description="Render CarbonStack generated command reference.")
    parser.add_argument("--registry", default="registry/commands.v0.yaml", help="input registry YAML path")
    parser.add_argument("--output", default="registry/COMMAND_REFERENCE.v0.md", help="output Markdown path")
    parser.add_argument("--check", action="store_true", help="verify output is up to date without writing")
    args = parser.parse_args(argv)

    registry_path = Path(args.registry)
    output_path = Path(args.output)

    entries = load_entries(registry_path)
    rendered, included = render(entries)
    issues = validate(entries, included)

    if issues:
        print("render validation failed:", file=sys.stderr)
        for issue in issues:
            print(f"- {issue}", file=sys.stderr)
        return 2

    if args.check:
        if not output_path.exists():
            print(f"missing generated reference: {output_path}", file=sys.stderr)
            return 3
        current = output_path.read_text(encoding="utf-8", errors="replace")
        if current != rendered:
            print(f"generated reference is stale: {output_path}", file=sys.stderr)
            return 4
        print(f"OK: generated reference is current: {output_path}")
        print(f"entries={len(entries)}")
        return 0

    output_path.parent.mkdir(parents=True, exist_ok=True)
    output_path.write_text(rendered, encoding="utf-8")
    print(f"WROTE {output_path}")
    print(f"entries={len(entries)}")
    return 0


if __name__ == "__main__":
    raise SystemExit(main(sys.argv[1:]))
