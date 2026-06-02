# CarbonStack v0.3.24 Cypher Schema Migrations Implementation

Status: v0 implementation checkpoint  
Scope: v0.3.x local-only backbone deployability line  
Primary environment: WSL Debian  
Generated: 2026-06-02 15:05:46 -0400

## 1. Purpose

This document records the v0.3.24 implementation rung for CarbonStackCypher migration tracking.

v0.3.23 proved that the previous migration path was fresh-DB safe but not persistent-DB restart/upgrade safe, because migrations could be reapplied against the same SQLite database. v0.3.24 implements minimal migration tracking so already-applied migrations are skipped and old migration files become immutable by checksum.

## 2. Current Repo Heads

    carbonstack        1fef7fe docs: record Cypher migration persistence recon
    carbonstack-comms  012c8bf scripts: support source snapshot artifact guard
    carbonstack-cypher 42f838f db: track applied migrations
    carbonstack-os     1bbbe52 docs: clarify CarbonStackOS target direction

## 3. Working Tree Status

    [carbonstack]
    (clean)

    [carbonstack-comms]
    (clean)

    [carbonstack-cypher]
    (clean)

    [carbonstack-os]
    (clean)

## 4. Toolchain Snapshot

    git:     git version 2.47.3
    go:      go version go1.24.4 linux/amd64
    rustc:   rustc 1.96.0 (ac68faa20 2026-05-25)
    cargo:   cargo 1.96.0 (30a34c682 2026-05-25)
    sqlite3: 3.46.1 2024-08-13 09:16:08 c9c2ab54ba1f5f46360f1b4f35d849cd3f080e6fc2b6c60e91b16c63f69aalt1 (64-bit)

## 5. Implementation Summary

carbonstack-cypher now creates and uses:

    schema_migrations

The table records:

    migration_name
    sha256
    applied_at

Migration behavior is now:

    create schema_migrations if missing
    read SQL migrations in sorted filename order
    compute SHA-256 over each migration file
    skip a migration if the same migration_name and sha256 were already applied
    hard-fail if a migration_name was already applied with a different sha256
    apply new migrations inside a transaction
    record the migration only after successful application

## 6. Tests Added

New test file:

    carbonstack-cypher/internal/db/migrate_test.go

Tests cover:

    TestMigrateRecordsAppliedMigrations
    TestMigrateCanRunTwiceAgainstSameDB
    TestMigrateDetectsAppliedMigrationChecksumMismatch

## 7. Validation Results

Command:

    cd ~/repos/carbonstack_umbrella/carbonstack-cypher
    go test ./internal/db -count=1 -v

Exit code:

    0

Output:

    === RUN   TestMigrateRecordsAppliedMigrations
--- PASS: TestMigrateRecordsAppliedMigrations (0.00s)
=== RUN   TestMigrateCanRunTwiceAgainstSameDB
--- PASS: TestMigrateCanRunTwiceAgainstSameDB (0.00s)
=== RUN   TestMigrateDetectsAppliedMigrationChecksumMismatch
--- PASS: TestMigrateDetectsAppliedMigrationChecksumMismatch (0.00s)
PASS
ok  	git.bitcrusher32.win/bitcrusher32/carbonstack-cypher/internal/db	0.007s

Command:

    cd ~/repos/carbonstack_umbrella/carbonstack-cypher
    go test ./... -count=1

Exit code:

    0

Output:

    ?   	git.bitcrusher32.win/bitcrusher32/carbonstack-cypher/cmd/cypher	[no test files]
?   	git.bitcrusher32.win/bitcrusher32/carbonstack-cypher/internal/config	[no test files]
ok  	git.bitcrusher32.win/bitcrusher32/carbonstack-cypher/internal/db	0.007s
ok  	git.bitcrusher32.win/bitcrusher32/carbonstack-cypher/internal/httpapi	0.051s

## 8. Security / Doctrine Interpretation

This implementation supports CarbonStack's evidence-led and attack-surface-minimizing style.

Migration files are now effectively immutable once applied to a database. If an already-applied migration file changes, startup/migration fails instead of silently accepting the drift.

This does not make Cypher production-ready. It only removes one blocker for experimental local operator persistence.

## 9. Remaining Limits

Still not validated:

    production deployability
    production E2EE
    hostile-server safety
    metadata privacy
    public ingress
    cloudflared
    systemd
    real homelab deployment
    runtime Comms OpenMLS UX
    secure local vault
    Android app
    CarbonStackOS
    external audit
    certification

Still future work:

    local operator runbook
    local config/data-path convention
    local-only operator profile validation
    local-backbone runner profile design, if justified later
    CarbonStack Relay Space mechanics
    runtime Comms send/inbox OpenMLS integration in v0.4.x

## 10. Recommended Next Rung

Recommended next rung:

    v0.3.25 local operator runbook skeleton

Suggested focus:

    document 127.0.0.1-only Cypher startup
    document explicit CYPHER_DB path
    document explicit CYPHER_MIGRATIONS path
    document reset/cleanup semantics
    document what survives restart
    document that persistence is still experimental
    avoid systemd/cloudflared/public ingress

## 11. Implementation Reference: Migrate Function

        	"crypto/sha256"
    func (s *Store) Migrate(dir string) error {
    CREATE TABLE IF NOT EXISTS schema_migrations (
        sha256 TEXT NOT NULL,
        applied_at TEXT NOT NULL
    		return fmt.Errorf("ensure schema_migrations table: %w", err)
    		sum := sha256.Sum256(content)
    			"SELECT sha256 FROM schema_migrations WHERE migration_name = ? LIMIT 1",
    					"migration %s checksum mismatch: applied %s, current %s",
    			"INSERT INTO schema_migrations (migration_name, sha256, applied_at) VALUES (?, ?, ?)",
    	sum := sha256.Sum256([]byte(code))

## 12. Implementation Reference: Migration Tests

        func TestMigrateRecordsAppliedMigrations(t *testing.T) {
    	if err := store.Migrate(migrationsDir); err != nil {
    	rows, err := store.DB.Query("SELECT migration_name, sha256, applied_at FROM schema_migrations ORDER BY migration_name ASC")
    		t.Fatalf("query schema_migrations: %v", err)
    			t.Fatalf("scan schema_migrations row: %v", err)
    		t.Fatalf("iterate schema_migrations: %v", err)
    func TestMigrateCanRunTwiceAgainstSameDB(t *testing.T) {
    	if err := store.Migrate(migrationsDir); err != nil {
    	if err := store.Migrate(migrationsDir); err != nil {
    func TestMigrateDetectsAppliedMigrationChecksumMismatch(t *testing.T) {
    	if err := store.Migrate(migrationsCopy); err != nil {
    	if _, err := f.WriteString("\n-- checksum mismatch recon mutation\n"); err != nil {
    	err = store.Migrate(migrationsCopy)
    		t.Fatalf("expected checksum mismatch after mutating an already-applied migration")
    	if !strings.Contains(err.Error(), "checksum mismatch") {
    		t.Fatalf("error = %v, want checksum mismatch", err)
