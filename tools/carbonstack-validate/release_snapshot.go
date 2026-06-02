package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func (r *Runner) ReleaseSnapshot() error {
	r.PrintHeader("release-snapshot")

	fmt.Printf("release_package_root: %s\n", r.UmbrellaRoot)
	fmt.Println("release-snapshot checks run before core validation")
	fmt.Println("run-order warning: validate only fresh extracted/staged package roots")
	fmt.Println("do not validate the package source root that will later be archived/published")
	fmt.Println("a successful validation generates artifacts, so rerun from a fresh extraction")

	if err := r.CheckReleaseSnapshotLayout(); err != nil {
		return err
	}

	if err := r.StrictPreTestArtifactScan(); err != nil {
		return err
	}

	if err := r.VerifyReleaseChecksums(); err != nil {
		return err
	}

	fmt.Println()
	fmt.Println("release-snapshot package checks and checksum verification passed; calling core validation")
	fmt.Println()

	return r.Core()
}

func (r *Runner) CheckReleaseSnapshotLayout() error {
	fmt.Println()
	fmt.Println("== Release snapshot layout checks ==")

	requiredDirs := []string{
		r.CarbonStack,
		r.Comms,
		r.Cypher,
	}

	requiredFiles := []string{
		filepath.Join(r.CarbonStack, "README.md"),
		filepath.Join(r.CarbonStack, "LICENSE"),
		filepath.Join(r.CarbonStack, "docs", "README.md"),
		filepath.Join(r.CarbonStack, "roadmap", "ROADMAP.md"),
		filepath.Join(r.CarbonStack, "tools", "carbonstack-validate", "go.mod"),
		filepath.Join(r.CarbonStack, "tools", "carbonstack-validate", "main.go"),
		filepath.Join(r.CarbonStack, "tools", "carbonstack-validate", "README.md"),

		filepath.Join(r.Comms, "README.md"),
		filepath.Join(r.Comms, "LICENSE"),
		filepath.Join(r.Comms, "go.mod"),
		filepath.Join(r.Comms, "internal", "protocol", "mls", "openmls-sidecar", "Cargo.toml"),

		filepath.Join(r.Cypher, "README.md"),
		filepath.Join(r.Cypher, "LICENSE"),
		filepath.Join(r.Cypher, "go.mod"),
		filepath.Join(r.Cypher, "go.sum"),
		filepath.Join(r.Cypher, "migrations", "001_init.sql"),
	}

	var missing []string

	for _, dir := range requiredDirs {
		info, err := os.Stat(dir)
		if err != nil || !info.IsDir() {
			fmt.Printf("MISSING DIR:  %s\n", dir)
			missing = append(missing, dir)
		} else {
			fmt.Printf("OK DIR:       %s\n", dir)
		}
	}

	for _, file := range requiredFiles {
		info, err := os.Stat(file)
		if err != nil || info.IsDir() {
			fmt.Printf("MISSING FILE: %s\n", file)
			missing = append(missing, file)
		} else {
			fmt.Printf("OK FILE:      %s\n", file)
		}
	}

	commsGoSum := filepath.Join(r.Comms, "go.sum")
	if info, err := os.Stat(commsGoSum); err == nil && !info.IsDir() {
		fmt.Printf("OPTIONAL FILE PRESENT: %s\n", commsGoSum)
	} else {
		fmt.Printf("OPTIONAL FILE ABSENT:  %s\n", commsGoSum)
	}

	if err := r.CheckReleaseMetadata(); err != nil {
		missing = append(missing, err.Error())
	}

	if len(missing) > 0 {
		return fmt.Errorf("release snapshot layout check failed with %d missing required item(s)", len(missing))
	}

	return nil
}

func (r *Runner) CheckReleaseMetadata() error {
	fmt.Println()
	fmt.Println("== Release metadata checks ==")

	releaseDir := filepath.Join(r.UmbrellaRoot, "release")

	info, err := os.Stat(releaseDir)
	if err != nil || !info.IsDir() {
		fmt.Printf("MISSING DIR:  %s\n", releaseDir)
		return fmt.Errorf("missing release metadata dir: %s", releaseDir)
	}

	fmt.Printf("OK DIR:       %s\n", releaseDir)

	requiredAny := [][]string{
		{
			filepath.Join(releaseDir, "manifest.json"),
		},
		{
			filepath.Join(releaseDir, "checksums.txt"),
		},
		{
			filepath.Join(releaseDir, "validation-freeze.md"),
			filepath.Join(releaseDir, "testing-runbook.md"),
		},
	}

	var missingGroups int

	for _, group := range requiredAny {
		found := false
		for _, path := range group {
			info, err := os.Stat(path)
			if err == nil && !info.IsDir() {
				fmt.Printf("OK RELEASE FILE: %s\n", path)
				found = true
			}
		}

		if !found {
			fmt.Printf("MISSING RELEASE FILE GROUP: %s\n", strings.Join(group, " OR "))
			missingGroups++
		}
	}

	optional := []string{
		filepath.Join(releaseDir, "release-notes.md"),
		filepath.Join(releaseDir, "LICENSE"),
	}

	for _, path := range optional {
		info, err := os.Stat(path)
		if err == nil && !info.IsDir() {
			fmt.Printf("OPTIONAL RELEASE FILE PRESENT: %s\n", path)
		} else {
			fmt.Printf("OPTIONAL RELEASE FILE ABSENT:  %s\n", path)
		}
	}

	if missingGroups > 0 {
		return fmt.Errorf("missing %d required release metadata group(s)", missingGroups)
	}

	return nil
}

func (r *Runner) StrictPreTestArtifactScan() error {
	fmt.Println()
	fmt.Println("== release-snapshot strict pre-test artifact scan ==")

	patterns := []string{
		".git",
		"target",
		".carbonstack-openmls-sidecar-state",
		"provider-storage.json",
		"signer.json",
		".go-cache",
		".go-tmp",
		"Thumbs.db",
		".DS_Store",
	}

	suffixes := []string{
		".db",
		".db-shm",
		".db-wal",
		".exe",
		".test.exe",
	}

	roots := []string{
		r.CarbonStack,
		r.Comms,
		r.Cypher,
	}

	var hits []string

	for _, root := range roots {
		_ = filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
			if err != nil {
				return nil
			}

			name := d.Name()

			for _, pattern := range patterns {
				if name == pattern {
					hits = append(hits, path)
					if d.IsDir() {
						return filepath.SkipDir
					}
					return nil
				}
			}

			if !d.IsDir() {
				for _, suffix := range suffixes {
					if strings.HasSuffix(name, suffix) {
						hits = append(hits, path)
						return nil
					}
				}
			}

			return nil
		})
	}

	if len(hits) == 0 {
		fmt.Println("strict pre-test artifact scan: PASS / no forbidden artifacts")
		return nil
	}

	fmt.Println("strict pre-test artifact scan: FAIL")
	for _, hit := range hits {
		fmt.Printf("  %s\n", hit)
	}

	return fmt.Errorf("release snapshot contains forbidden generated/private/build artifacts before tests")
}
