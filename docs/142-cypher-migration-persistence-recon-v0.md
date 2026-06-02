# CarbonStack v0.3.23 Cypher Migration Persistence Recon

Status: v0 recon / proof checkpoint  
Scope: v0.3.x local-only backbone deployability line  
Primary environment: WSL Debian  
Generated: 2026-06-02 14:27:18 -0400

## 1. Purpose

This document records a proof-first recon of CarbonStackCypher SQLite migration behavior against a reused persistent database.

v0.3.21 recorded the local-only backbone deployability model. v0.3.22 recorded the Cypher local operator surface recon and identified persistent SQLite migration behavior as the highest-priority deployability risk. This v0.3.23 recon tests that risk directly before any schema or migration implementation work.

This is intentionally not a production deployability claim.

## 2. Current Repo Heads

```text
carbonstack        328ffc4 docs: record Cypher local operator surface recon
carbonstack-comms  012c8bf scripts: support source snapshot artifact guard
carbonstack-cypher 6f92c34 chore: fix readme formatting
carbonstack-os     1bbbe52 docs: clarify CarbonStackOS target direction
```

## 3. Working Tree Status Before Recon

```text
[carbonstack]
(clean)

[carbonstack-comms]
(clean)

[carbonstack-cypher]
(clean)

[carbonstack-os]
(clean)
```

## 4. Toolchain Snapshot

```text
git:     git version 2.47.3
go:      go version go1.24.4 linux/amd64
rustc:   rustc 1.96.0 (ac68faa20 2026-05-25)
cargo:   cargo 1.96.0 (30a34c682 2026-05-25)
sqlite3: 3.46.1 2024-08-13 09:16:08 c9c2ab54ba1f5f46360f1b4f35d849cd3f080e6fc2b6c60e91b16c63f69aalt1 (64-bit)
```

## 5. Question Tested

Can the current CarbonStackCypher migration path safely run twice against the same SQLite database?

This matters because a local operator profile implies that a Cypher SQLite database may survive process restarts. If migrations are reapplied blindly, startup can fail even though a fresh test database works.

## 6. Temporary Recon Test Method

A temporary Go test file was written under:

```text
carbonstack-cypher/internal/db/migrate_twice_recon_test.go
```

The test:

1. creates a temporary SQLite DB;
2. opens it through the current `internal/db` package;
3. runs `Store.Migrate("../../migrations")` once;
4. runs `Store.Migrate("../../migrations")` a second time against the same DB;
5. expects the second pass to fail under current behavior;
6. deletes the temporary test file after the proof run.

The temporary test file was not intended to be committed.

## 7. Recon Test Result

Result:

```text
PASS / current hazard reproduced
```

Command:

```text
cd ~/repos/carbonstack_umbrella/carbonstack-cypher
go test ./internal/db -run TestMigrateTwiceReconCurrentBehavior -count=1 -v
```

Output:

```text
=== RUN   TestMigrateTwiceReconCurrentBehavior
    migrate_twice_recon_test.go:28: second migration pass failed as expected for current recon: apply migration ../../migrations/002_envelope_payload_metadata.sql: SQL logic error: duplicate column name: payload_sha256 (1)
--- PASS: TestMigrateTwiceReconCurrentBehavior (0.00s)
PASS
ok  	git.bitcrusher32.win/bitcrusher32/carbonstack-cypher/internal/db	0.005s
```

Cypher working tree after removing the temporary recon test:

```text
(clean)
```

## 8. Current Migration Files

### 001_init.sql

```sql
CREATE TABLE IF NOT EXISTS invites (
    invite_id TEXT PRIMARY KEY,
    invite_code_hash TEXT NOT NULL UNIQUE,
    created_at TEXT NOT NULL,
    claimed_at TEXT,
    claimed_by_account_id TEXT,
    disabled_at TEXT
);

CREATE TABLE IF NOT EXISTS accounts (
    account_id TEXT PRIMARY KEY,
    display_name TEXT NOT NULL,
    created_at TEXT NOT NULL,
    disabled_at TEXT
);

CREATE TABLE IF NOT EXISTS devices (
    device_id TEXT PRIMARY KEY,
    account_id TEXT NOT NULL,
    device_label TEXT NOT NULL,
    public_identity_key TEXT NOT NULL,
    public_prekey_bundle TEXT,
    created_at TEXT NOT NULL,
    revoked_at TEXT,
    FOREIGN KEY(account_id) REFERENCES accounts(account_id)
);

CREATE TABLE IF NOT EXISTS envelopes (
    envelope_id TEXT PRIMARY KEY,
    sender_device_id TEXT NOT NULL,
    recipient_device_id TEXT NOT NULL,
    content_type TEXT NOT NULL,
    protocol_version TEXT NOT NULL,
    ciphertext_b64 TEXT NOT NULL,
    client_created_at TEXT,
    server_received_at TEXT NOT NULL,
    delivery_state TEXT NOT NULL,
    FOREIGN KEY(sender_device_id) REFERENCES devices(device_id),
    FOREIGN KEY(recipient_device_id) REFERENCES devices(device_id)
);

CREATE TABLE IF NOT EXISTS envelope_acks (
    ack_id TEXT PRIMARY KEY,
    envelope_id TEXT NOT NULL,
    recipient_device_id TEXT NOT NULL,
    acknowledged_at TEXT NOT NULL,
    FOREIGN KEY(envelope_id) REFERENCES envelopes(envelope_id),
    FOREIGN KEY(recipient_device_id) REFERENCES devices(device_id)
);
```

### 002_envelope_payload_metadata.sql

```sql
ALTER TABLE envelopes ADD COLUMN payload_sha256 TEXT;
ALTER TABLE envelopes ADD COLUMN payload_size_bytes INTEGER;
```

## 9. Current Migrate Function Excerpt

```go
func (s *Store) Migrate(dir string) error {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("read migrations dir: %w", err)
	}

	var files []string
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		if filepath.Ext(entry.Name()) == ".sql" {
			files = append(files, filepath.Join(dir, entry.Name()))
		}
	}

	sort.Strings(files)

	for _, file := range files {
		body, err := os.ReadFile(file)
		if err != nil {
			return fmt.Errorf("read migration %s: %w", file, err)
		}

		if _, err := s.DB.Exec(string(body)); err != nil {
			return fmt.Errorf("apply migration %s: %w", file, err)
		}
	}

	return nil
}
```

## 10. Interpretation

The current migration path is good enough for fresh test databases, but it is not yet safe to claim persistent local operator DB restart/upgrade behavior.

The current runner reads migration files and executes them in sorted order. `001_init.sql` is mostly restart-friendly because it uses `CREATE TABLE IF NOT EXISTS`, but `002_envelope_payload_metadata.sql` uses raw `ALTER TABLE ... ADD COLUMN`. When migrations are applied a second time against an already-upgraded DB, SQLite reports duplicate-column behavior.

This confirms the v0.3.22 recon concern: persistent SQLite state needs migration tracking, idempotent migration guards, or explicit wipe-only experimental documentation before stronger local operator claims.

## 11. Deployability Impact

This affects the local operator profile directly.

Current safe claim:

```text
Fresh Cypher DB behavior is the known-good path.
Persistent local Cypher DB behavior is experimental.
Schema reset may be required between versions.
```

Not safe yet:

```text
Cypher can safely run persistent local operator DB migrations across restarts/upgrades.
```

## 12. Recommended Next Implementation Rung

Recommended next rung:

```text
v0.3.24:
  implement minimal schema_migrations tracking in carbonstack-cypher
  add tests for:
    fresh DB migration
    repeated migration on same DB
    migration table records applied versions
```

Alternative:

```text
Document wipe-only experimental DB behavior and defer schema_migrations.
```

Recommendation: implement `schema_migrations` before claiming persistent local operator DB behavior.

## 13. Explicit Nonclaims

This recon does not validate:

- production deployability
- production E2EE
- hostile-server safety
- metadata privacy
- public ingress
- cloudflared
- systemd
- real homelab deployment
- runtime Comms OpenMLS UX
- secure local vault
- Android app
- CarbonStackOS
- external audit
- certification

## 14. Summary

v0.3.23 proves the current persistent SQLite migration hazard rather than assuming it.

The next safest technical step is to add migration tracking or idempotent migration behavior before building a stronger persistent local operator profile.
