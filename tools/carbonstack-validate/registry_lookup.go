package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type registryLookupEntry struct {
	ID     string
	Fields map[string]string
	Lists  map[string][]string
}

func (r *Runner) RegistryLookup(registryID string, literalCommand string) error {
	r.PrintHeader("registry-lookup")

	fmt.Println("status: dev/operator registry lookup")
	fmt.Println("scope: command-boundary registry inspection")
	fmt.Println("boundary: registry presence is classification, not promotion; output is not a public UX stability claim")
	fmt.Println()

	if strings.TrimSpace(registryID) == "" && strings.TrimSpace(literalCommand) == "" {
		return fmt.Errorf("registry-lookup requires --registry-id or --command")
	}
	if strings.TrimSpace(registryID) != "" && strings.TrimSpace(literalCommand) != "" {
		return fmt.Errorf("registry-lookup accepts either --registry-id or --command, not both")
	}

	registryPath := filepath.Join(r.CarbonStack, "registry", "commands.v0.yaml")
	raw, err := os.ReadFile(registryPath)
	if err != nil {
		return fmt.Errorf("read registry %s: %w", registryPath, err)
	}

	entries := parseRegistryLookupEntries(string(raw))
	var matches []registryLookupEntry

	for _, entry := range entries {
		switch {
		case strings.TrimSpace(registryID) != "":
			if entry.ID == strings.TrimSpace(registryID) {
				matches = append(matches, entry)
			}
		case strings.TrimSpace(literalCommand) != "":
			if entry.Fields["command"] == strings.TrimSpace(literalCommand) {
				matches = append(matches, entry)
			}
		}
	}

	if len(matches) == 0 {
		switch {
		case strings.TrimSpace(registryID) != "":
			return fmt.Errorf("registry entry not found for id %q", registryID)
		default:
			return fmt.Errorf("registry entry not found for command %q", literalCommand)
		}
	}
	if len(matches) > 1 {
		fmt.Printf("matches: %d\n", len(matches))
		for _, entry := range matches {
			fmt.Printf("- %s\n", entry.ID)
		}
		return fmt.Errorf("registry lookup was ambiguous")
	}

	printRegistryLookupEntry(matches[0])
	return nil
}

func parseRegistryLookupEntries(text string) []registryLookupEntry {
	var entries []registryLookupEntry
	var current *registryLookupEntry
	currentList := ""

	for _, line := range strings.Split(text, "\n") {
		trimmed := strings.TrimSpace(line)

		if strings.HasPrefix(trimmed, "- id: ") {
			entry := registryLookupEntry{
				ID:     strings.TrimSpace(strings.TrimPrefix(trimmed, "- id: ")),
				Fields: map[string]string{},
				Lists:  map[string][]string{},
			}
			entries = append(entries, entry)
			current = &entries[len(entries)-1]
			currentList = ""
			continue
		}

		if current == nil {
			continue
		}

		if strings.HasPrefix(trimmed, "- ") {
			if currentList != "" {
				current.Lists[currentList] = append(current.Lists[currentList], strings.TrimSpace(strings.TrimPrefix(trimmed, "- ")))
			}
			continue
		}

		if !strings.Contains(trimmed, ":") {
			continue
		}

		parts := strings.SplitN(trimmed, ":", 2)
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		if key == "" {
			continue
		}

		if value == "" {
			currentList = key
			if _, ok := current.Lists[currentList]; !ok {
				current.Lists[currentList] = []string{}
			}
			continue
		}

		currentList = ""
		if _, exists := current.Fields[key]; !exists {
			current.Fields[key] = value
		}
	}

	return entries
}

func printRegistryLookupEntry(entry registryLookupEntry) {
	fmt.Println("registry entry")
	fmt.Printf("id: %s\n", entry.ID)
	printRegistryField(entry, "command")
	printRegistryField(entry, "repo")
	printRegistryField(entry, "component")
	printRegistryField(entry, "kind")
	printRegistryField(entry, "audience")
	printRegistryField(entry, "maturity")
	printRegistryField(entry, "lifecycle_status")
	printRegistryField(entry, "introduced_in")
	printRegistryField(entry, "source_path")
	printRegistryField(entry, "validation_surface")
	printRegistryField(entry, "short_help")
	printRegistryField(entry, "why_exists")
	printRegistryField(entry, "replaced_by")
	printRegistryField(entry, "wrapped_by")
	printRegistryField(entry, "include_in_front_readme")

	printRegistryList(entry, "required_flags")
	printRegistryList(entry, "optional_flags")
	printRegistryList(entry, "related_profiles")
	printRegistryList(entry, "related_scripts")
	printRegistryList(entry, "nonclaims")

	fmt.Println("boundary: registry presence is classification, not promotion")
}

func printRegistryField(entry registryLookupEntry, key string) {
	if value := strings.TrimSpace(entry.Fields[key]); value != "" {
		fmt.Printf("%s: %s\n", key, value)
	}
}

func printRegistryList(entry registryLookupEntry, key string) {
	values := entry.Lists[key]
	if len(values) == 0 {
		return
	}
	fmt.Printf("%s:\n", key)
	for _, value := range values {
		fmt.Printf("  - %s\n", value)
	}
}
