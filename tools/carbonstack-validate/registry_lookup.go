package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type registryLookupEntry struct {
	ID     string
	Fields map[string]string
	Lists  map[string][]string
}

type registryLookupOptions struct {
	RegistryID       string
	LiteralCommand   string
	List             bool
	Audience         string
	Maturity         string
	LifecycleStatus  string
	Kind             string
	FrontReadmeOnly  bool
	MissingNonclaims bool
}

func (r *Runner) RegistryLookup(registryID string, literalCommand string) error {
	return r.RegistryLookupWithOptions(registryLookupOptions{
		RegistryID:     registryID,
		LiteralCommand: literalCommand,
	})
}

func (r *Runner) RegistryLookupWithOptions(opts registryLookupOptions) error {
	r.PrintHeader("registry-lookup")

	fmt.Println("status: dev/operator registry lookup")
	fmt.Println("scope: command-boundary registry inspection")
	fmt.Println("boundary: registry presence is classification, not promotion; output is not a public UX stability claim")
	fmt.Println()

	registryPath := filepath.Join(r.CarbonStack, "registry", "commands.v0.yaml")
	raw, err := os.ReadFile(registryPath)
	if err != nil {
		return fmt.Errorf("read registry %s: %w", registryPath, err)
	}

	entries := parseRegistryLookupEntries(string(raw))
	matches := filterRegistryLookupEntries(entries, opts)

	if opts.List {
		printRegistryLookupList(matches, opts)
		return nil
	}

	if strings.TrimSpace(opts.RegistryID) == "" && strings.TrimSpace(opts.LiteralCommand) == "" {
		return fmt.Errorf("registry-lookup requires --registry-id or --command, or --list with optional filters")
	}
	if strings.TrimSpace(opts.RegistryID) != "" && strings.TrimSpace(opts.LiteralCommand) != "" {
		return fmt.Errorf("registry-lookup accepts either --registry-id or --command, not both")
	}

	if len(matches) == 0 {
		switch {
		case strings.TrimSpace(opts.RegistryID) != "":
			return fmt.Errorf("registry entry not found for id %q", opts.RegistryID)
		default:
			return fmt.Errorf("registry entry not found for command %q", opts.LiteralCommand)
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

func filterRegistryLookupEntries(entries []registryLookupEntry, opts registryLookupOptions) []registryLookupEntry {
	var matches []registryLookupEntry
	for _, entry := range entries {
		if opts.RegistryID != "" && entry.ID != strings.TrimSpace(opts.RegistryID) {
			continue
		}
		if opts.LiteralCommand != "" && entry.Fields["command"] != strings.TrimSpace(opts.LiteralCommand) {
			continue
		}
		if opts.Audience != "" && entry.Fields["audience"] != strings.TrimSpace(opts.Audience) {
			continue
		}
		if opts.Maturity != "" && entry.Fields["maturity"] != strings.TrimSpace(opts.Maturity) {
			continue
		}
		if opts.LifecycleStatus != "" && entry.Fields["lifecycle_status"] != strings.TrimSpace(opts.LifecycleStatus) {
			continue
		}
		if opts.Kind != "" && entry.Fields["kind"] != strings.TrimSpace(opts.Kind) {
			continue
		}
		if opts.FrontReadmeOnly && entry.Fields["include_in_front_readme"] != "true" {
			continue
		}
		if opts.MissingNonclaims && len(entry.Lists["nonclaims"]) > 0 {
			continue
		}
		matches = append(matches, entry)
	}
	sort.Slice(matches, func(i, j int) bool {
		return matches[i].ID < matches[j].ID
	})
	return matches
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

func printRegistryLookupList(entries []registryLookupEntry, opts registryLookupOptions) {
	fmt.Println("registry entries")
	fmt.Printf("matches: %d\n", len(entries))
	if opts.Audience != "" {
		fmt.Printf("filter_audience: %s\n", opts.Audience)
	}
	if opts.Maturity != "" {
		fmt.Printf("filter_maturity: %s\n", opts.Maturity)
	}
	if opts.LifecycleStatus != "" {
		fmt.Printf("filter_lifecycle_status: %s\n", opts.LifecycleStatus)
	}
	if opts.Kind != "" {
		fmt.Printf("filter_kind: %s\n", opts.Kind)
	}
	if opts.FrontReadmeOnly {
		fmt.Println("filter_front_readme_only: true")
	}
	if opts.MissingNonclaims {
		fmt.Println("filter_missing_nonclaims: true")
	}
	for _, entry := range entries {
		fmt.Printf("- id: %s\n", entry.ID)
		if command := entry.Fields["command"]; command != "" {
			fmt.Printf("  command: %s\n", command)
		}
		if audience := entry.Fields["audience"]; audience != "" {
			fmt.Printf("  audience: %s\n", audience)
		}
		if maturity := entry.Fields["maturity"]; maturity != "" {
			fmt.Printf("  maturity: %s\n", maturity)
		}
		if lifecycle := entry.Fields["lifecycle_status"]; lifecycle != "" {
			fmt.Printf("  lifecycle_status: %s\n", lifecycle)
		}
		if help := entry.Fields["short_help"]; help != "" {
			fmt.Printf("  short_help: %s\n", help)
		}
	}
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
