package main

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

type registryEntry struct {
	ID     string
	Fields map[string]string
}

func TestCommandRegistryV0Coverage(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	carbonstackRoot := filepath.Clean(filepath.Join(cwd, "..", ".."))
	umbrellaRoot := filepath.Dir(carbonstackRoot)

	if !dirExists(filepath.Join(umbrellaRoot, "carbonstack-comms")) || !dirExists(filepath.Join(umbrellaRoot, "carbonstack-cypher")) {
		t.Skip("command registry coverage test requires a CarbonStack umbrella layout with carbonstack-comms and carbonstack-cypher siblings")
	}

	registryPath := filepath.Join(carbonstackRoot, "registry", "commands.v0.yaml")
	textBytes, err := os.ReadFile(registryPath)
	if err != nil {
		t.Fatalf("read registry: %v", err)
	}
	registryText := string(textBytes)

	entries := parseRegistryEntries(registryText)
	if len(entries) == 0 {
		t.Fatalf("no registry entries parsed from %s", registryPath)
	}

	ids := map[string]registryEntry{}
	for _, entry := range entries {
		if prior, ok := ids[entry.ID]; ok {
			t.Fatalf("duplicate registry id %q found twice: %#v / %#v", entry.ID, prior, entry)
		}
		ids[entry.ID] = entry
	}

	requiredFields := []string{
		"command",
		"repo",
		"component",
		"kind",
		"audience",
		"maturity",
		"source_path",
		"short_help",
		"why_exists",
		"include_in_front_readme",
	}

	for _, entry := range entries {
		for _, field := range requiredFields {
			if strings.TrimSpace(entry.Fields[field]) == "" {
				t.Fatalf("registry entry %s missing required field %s", entry.ID, field)
			}
		}

		sourcePath := resolveRegistrySourcePath(umbrellaRoot, carbonstackRoot, entry.Fields["source_path"])
		if _, err := os.Stat(sourcePath); err != nil {
			t.Fatalf("registry entry %s has missing source_path %q resolved to %s: %v", entry.ID, entry.Fields["source_path"], sourcePath, err)
		}
	}

	assertRegistryContainsIDs(t, ids, []string{
		"runner.doctor",
		"runner.core",
		"runner.local-cypher",
		"runner.dev-runtime-openmls",
		"runner.dev-runtime-openmls-wrappers",
		"runner.full",
		"runner.release-snapshot",
		"runner.write-checksums",
		"runner.verify-checksums",
		"comms.send",
		"comms.inbox",
		"comms.ack",
		"comms.openmls-send-dev",
		"comms.openmls-inbox-dev",
		"comms.openmls-conversation-join-dev",
		"sidecar.message-protect",
		"sidecar.message-open",
		"cypher.server",
		"cypher.api.envelope-ack",
	})

	assertRunnerProfileCoverage(t, ids, carbonstackRoot)
	assertCommsCommandCoverage(t, ids, umbrellaRoot)
	assertCommsScriptCoverage(t, registryText, umbrellaRoot)
	assertSidecarCommandCoverage(t, ids)
	assertCypherSurfaceCoverage(t, ids)
	assertRegistryPolicyBoundaries(t, entries, registryText)
}

func parseRegistryEntries(text string) []registryEntry {
	entries := []registryEntry{}
	var current *registryEntry

	for _, line := range strings.Split(text, "\n") {
		trimmed := strings.TrimSpace(line)

		if strings.HasPrefix(trimmed, "- id: ") {
			entry := registryEntry{
				ID:     strings.TrimSpace(strings.TrimPrefix(trimmed, "- id: ")),
				Fields: map[string]string{},
			}
			entries = append(entries, entry)
			current = &entries[len(entries)-1]
			continue
		}

		if current == nil || strings.HasPrefix(trimmed, "- ") || !strings.Contains(trimmed, ":") {
			continue
		}

		parts := strings.SplitN(trimmed, ":", 2)
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		if key != "" && current.Fields[key] == "" {
			current.Fields[key] = value
		}
	}

	return entries
}

func resolveRegistrySourcePath(umbrellaRoot string, carbonstackRoot string, sourcePath string) string {
	sourcePath = strings.TrimSpace(sourcePath)

	switch {
	case strings.HasPrefix(sourcePath, "carbonstack/"):
		return filepath.Join(umbrellaRoot, filepath.FromSlash(sourcePath))
	case strings.HasPrefix(sourcePath, "carbonstack-comms/"):
		return filepath.Join(umbrellaRoot, filepath.FromSlash(sourcePath))
	case strings.HasPrefix(sourcePath, "carbonstack-cypher/"):
		return filepath.Join(umbrellaRoot, filepath.FromSlash(sourcePath))
	case strings.HasPrefix(sourcePath, "carbonstack-os/"):
		return filepath.Join(umbrellaRoot, filepath.FromSlash(sourcePath))
	default:
		return filepath.Join(carbonstackRoot, filepath.FromSlash(sourcePath))
	}
}

func assertRunnerProfileCoverage(t *testing.T, ids map[string]registryEntry, carbonstackRoot string) {
	t.Helper()

	mainBytes, err := os.ReadFile(filepath.Join(carbonstackRoot, "tools", "carbonstack-validate", "main.go"))
	if err != nil {
		t.Fatalf("read runner main.go: %v", err)
	}

	re := regexp.MustCompile(`case "([^"]+)":`)
	matches := re.FindAllStringSubmatch(string(mainBytes), -1)

	if len(matches) == 0 {
		t.Fatalf("no runner profile cases found")
	}

	for _, match := range matches {
		id := "runner." + match[1]
		if _, ok := ids[id]; !ok {
			t.Fatalf("runner profile %q missing registry id %q", match[1], id)
		}
	}
}

func assertCommsCommandCoverage(t *testing.T, ids map[string]registryEntry, umbrellaRoot string) {
	t.Helper()

	commandBytes, err := os.ReadFile(filepath.Join(umbrellaRoot, "carbonstack-comms", "internal", "app", "commands.go"))
	if err != nil {
		t.Fatalf("read Comms commands.go: %v", err)
	}

	text := string(commandBytes)
	start := strings.Index(text, "func Run(args []string) error")
	end := strings.Index(text, "func usage()")
	if start < 0 || end < 0 || end <= start {
		t.Fatalf("could not isolate Comms Run command dispatch")
	}

	dispatch := text[start:end]
	re := regexp.MustCompile(`case "([^"]+)":`)
	matches := re.FindAllStringSubmatch(dispatch, -1)

	if len(matches) == 0 {
		t.Fatalf("no Comms command cases found")
	}

	for _, match := range matches {
		id := "comms." + match[1]
		if _, ok := ids[id]; !ok {
			t.Fatalf("Comms command %q missing registry id %q", match[1], id)
		}
	}
}

func assertCommsScriptCoverage(t *testing.T, registryText string, umbrellaRoot string) {
	t.Helper()

	scriptsDir := filepath.Join(umbrellaRoot, "carbonstack-comms", "scripts")
	entries, err := os.ReadDir(scriptsDir)
	if err != nil {
		t.Fatalf("read Comms scripts dir: %v", err)
	}

	for _, entry := range entries {
		if entry.IsDir() || entry.Name() == "README.md" {
			continue
		}

		scriptPath := "scripts/" + entry.Name()
		if !strings.Contains(registryText, scriptPath) {
			t.Fatalf("Comms script %s is missing from registry", scriptPath)
		}
	}
}

func assertSidecarCommandCoverage(t *testing.T, ids map[string]registryEntry) {
	t.Helper()

	assertRegistryContainsIDs(t, ids, []string{
		"sidecar.provider-info",
		"sidecar.identity-create",
		"sidecar.identity-status",
		"sidecar.public-bundle-export",
		"sidecar.conversation-create",
		"sidecar.conversation-load-check",
		"sidecar.conversation-add-member",
		"sidecar.conversation-join",
		"sidecar.message-protect",
		"sidecar.message-open",
		"sidecar.state-checkpoint",
		"sidecar.state-load-check",
	})
}

func assertCypherSurfaceCoverage(t *testing.T, ids map[string]registryEntry) {
	t.Helper()

	assertRegistryContainsIDs(t, ids, []string{
		"cypher.server",
		"cypher.api.health",
		"cypher.api.invites-claim",
		"cypher.api.dev-invites",
		"cypher.api.devices-register",
		"cypher.api.accounts-devices",
		"cypher.api.envelopes-submit",
		"cypher.api.device-envelopes",
		"cypher.api.envelope-ack",
	})
}

func assertRegistryPolicyBoundaries(t *testing.T, entries []registryEntry, registryText string) {
	t.Helper()

	if !strings.Contains(registryText, "This registry is not local-backbone.") {
		t.Fatalf("registry missing local-backbone nonclaim")
	}

	if !strings.Contains(registryText, "Gitea remains source of truth") {
		t.Fatalf("registry missing Gitea source-of-truth boundary")
	}

	for _, id := range []string{"comms.send", "comms.inbox", "comms.ack"} {
		entry := findEntry(entries, id)
		if entry.ID == "" {
			t.Fatalf("missing legacy entry %s", id)
		}
		if entry.Fields["maturity"] != "legacy" {
			t.Fatalf("entry %s should remain maturity legacy, got %q", id, entry.Fields["maturity"])
		}
	}

	direct := findEntry(entries, "runner.dev-runtime-openmls")
	wrappers := findEntry(entries, "runner.dev-runtime-openmls-wrappers")

	if direct.ID == "" || wrappers.ID == "" {
		t.Fatalf("missing direct or wrapper OpenMLS runtime profile entries")
	}

	if direct.Fields["validation_surface"] == wrappers.Fields["validation_surface"] {
		t.Fatalf("direct and wrapper OpenMLS runtime profiles should keep distinct validation surfaces")
	}

	frontReadmeCount := 0
	for _, entry := range entries {
		if entry.Fields["include_in_front_readme"] == "true" {
			frontReadmeCount++
		}
	}

	if frontReadmeCount == 0 || frontReadmeCount > 12 {
		t.Fatalf("unexpected include_in_front_readme count: %d", frontReadmeCount)
	}
}

func assertRegistryContainsIDs(t *testing.T, ids map[string]registryEntry, required []string) {
	t.Helper()

	for _, id := range required {
		if _, ok := ids[id]; !ok {
			t.Fatalf("missing registry id %s", id)
		}
	}
}

func findEntry(entries []registryEntry, id string) registryEntry {
	for _, entry := range entries {
		if entry.ID == id {
			return entry
		}
	}
	return registryEntry{}
}

func dirExists(path string) bool {
	info, err := os.Stat(path)
	return err == nil && info.IsDir()
}
