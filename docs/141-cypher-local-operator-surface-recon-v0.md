# CarbonStack v0.3.22 Cypher Local Operator Surface Recon

Status: v0 recon checkpoint  
Scope: v0.3.x local-only backbone deployability line  
Primary environment: WSL Debian  
Generated: 2026-06-02 13:57:36 -0400

## 1. Purpose

This document records the Cypher local operator surface recon rung after v0.3.21.

v0.3.21 established the local-only backbone deployability model. This v0.3.22 recon inspects the current code/documentation surfaces that matter before implementing a persistent local operator profile.

This is intentionally docs/recon only. It does not implement deployability, systemd, cloudflared, public ingress, runtime Comms UX, Android, CarbonStackOS, production security, or hostile-server proof.

## 2. Current Repo Heads

```text
carbonstack        6850c58 docs: record local backbone deployability recon
carbonstack-comms  012c8bf scripts: support source snapshot artifact guard
carbonstack-cypher 6f92c34 chore: fix readme formatting
carbonstack-os     1bbbe52 docs: clarify CarbonStackOS target direction
```

## 3. Current Working Tree Status

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

## 5. Terminology Decision: CarbonStack Relay Space

Future server/conversation mechanics should use the term **CarbonStack Relay Space** instead of IRC-like server unless historical comparison is needed.

Working definition:

> A CarbonStack Relay Space is one addressable Cypher relay context. In the early model, one relay space maps to one Cypher instance or isolated Cypher profile. Future tooling may allow one Cypher install to supervise multiple isolated relay spaces, but each relay space should keep separate address/config/data/operator boundaries.

This preserves the useful part of the IRC analogy — self-hosted, addressable, text-first relay contexts — without importing IRC culture, public-channel assumptions, or moderation semantics.

## 6. Cypher Source Surface Inventory

Cypher files inspected:

```text
carbonstack-cypher/README.md
carbonstack-cypher/cmd/cypher/main.go
carbonstack-cypher/docs/01-cypher-architecture.md
carbonstack-cypher/docs/02-envelope-model.md
carbonstack-cypher/docs/03-api-surface.md
carbonstack-cypher/docs/04-storage-model.md
carbonstack-cypher/docs/05-mvp-roadmap.md
carbonstack-cypher/docs/06-phase1-vertical-slice.md
carbonstack-cypher/docs/07-data-model-v0.md
carbonstack-cypher/docs/08-api-contract-v0.md
carbonstack-cypher/go.mod
carbonstack-cypher/internal/config/config.go
carbonstack-cypher/internal/db/db.go
carbonstack-cypher/internal/httpapi/api.go
carbonstack-cypher/internal/httpapi/api_test.go
carbonstack-cypher/migrations/001_init.sql
carbonstack-cypher/migrations/002_envelope_payload_metadata.sql
```

## 7. Cypher Environment / Config Surface

Observed `CYPHER_*` and config-related hits:

- `carbonstack-cypher/internal/config/config.go:14` — `Addr:          getEnv("CYPHER_ADDR", ":8080"),`
- `carbonstack-cypher/internal/config/config.go:15` — `DBPath:        getEnv("CYPHER_DB", "cypher.db"),`
- `carbonstack-cypher/internal/config/config.go:16` — `MigrationsDir: getEnv("CYPHER_MIGRATIONS", "migrations"),`
- `carbonstack-cypher/internal/config/config.go:17` — `DevInviteCode: getEnv("CYPHER_DEV_INVITE", "dev-invite"),`

Bind/address related hits:

- `carbonstack-cypher/docs/06-phase1-vertical-slice.md:54` — `5. Client A submits an envelope addressed to Client B device.`
- `carbonstack-cypher/internal/config/config.go:6` — `Addr          string`
- `carbonstack-cypher/internal/config/config.go:14` — `Addr:          getEnv("CYPHER_ADDR", ":8080"),`
- `carbonstack-cypher/cmd/cypher/main.go:31` — `log.Printf("CarbonStackCypher listening on %s", cfg.Addr)`
- `carbonstack-cypher/cmd/cypher/main.go:35` — `if err := http.ListenAndServe(cfg.Addr, api.Routes()); err != nil {`

Database/migration related hits:

- `carbonstack-cypher/README.md:71` — `Current migrations:`
- `carbonstack-cypher/README.md:73` — `migrations/001_init.sql`
- `carbonstack-cypher/README.md:74` — `migrations/002_envelope_payload_metadata.sql`
- `carbonstack-cypher/docs/05-mvp-roadmap.md:12` — `- SQLite migrations;`
- `carbonstack-cypher/docs/06-phase1-vertical-slice.md:77` — `Initial database: SQLite`
- `carbonstack-cypher/docs/01-cypher-architecture.md:119` — `The server database should be useless for plaintext recovery.`
- `carbonstack-cypher/docs/01-cypher-architecture.md:121` — `A database dump should not reveal message contents.`
- `carbonstack-cypher/docs/07-data-model-v0.md:9` — `The schema is implemented for the current relay/server scaffold. It is not a production database contract. It may change before any stable release.`
- `carbonstack-cypher/docs/07-data-model-v0.md:25` — `Current migrations:`
- `carbonstack-cypher/docs/07-data-model-v0.md:27` — `migrations/001_init.sql`
- `carbonstack-cypher/docs/07-data-model-v0.md:28` — `migrations/002_envelope_payload_metadata.sql`
- `carbonstack-cypher/docs/04-storage-model.md:27` — `The current schema is not a production database contract.`
- `carbonstack-cypher/docs/04-storage-model.md:29` — `## Current migrations`
- `carbonstack-cypher/docs/04-storage-model.md:31` — `Current migrations:`
- `carbonstack-cypher/docs/04-storage-model.md:33` — `migrations/001_init.sql`
- `carbonstack-cypher/docs/04-storage-model.md:34` — `migrations/002_envelope_payload_metadata.sql`
- `carbonstack-cypher/internal/httpapi/api_test.go:220` — `dbPath := filepath.Join(t.TempDir(), "cypher-test.db")`
- `carbonstack-cypher/internal/httpapi/api_test.go:222` — `store, err := db.Open(dbPath)`
- `carbonstack-cypher/internal/httpapi/api_test.go:231` — `migrationsDir := filepath.Join("..", "..", "migrations")`
- `carbonstack-cypher/internal/httpapi/api_test.go:232` — `if err := store.Migrate(migrationsDir); err != nil {`
- `carbonstack-cypher/internal/httpapi/api.go:5` — `"database/sql"`
- `carbonstack-cypher/internal/httpapi/api.go:84` — `err := a.store.DB.QueryRow(`
- `carbonstack-cypher/internal/httpapi/api.go:101` — `_, err = a.store.DB.Exec(`
- `carbonstack-cypher/internal/httpapi/api.go:138` — `tx, err := a.store.DB.Begin()`
- `carbonstack-cypher/internal/httpapi/api.go:230` — `err := a.store.DB.QueryRow(`
- `carbonstack-cypher/internal/httpapi/api.go:247` — `_, err = a.store.DB.Exec(`
- `carbonstack-cypher/internal/httpapi/api.go:275` — `rows, err := a.store.DB.Query(`
- `carbonstack-cypher/internal/httpapi/api.go:374` — `_, err = a.store.DB.Exec(`
- `carbonstack-cypher/internal/httpapi/api.go:413` — `rows, err := a.store.DB.Query(`
- `carbonstack-cypher/internal/httpapi/api.go:495` — `err := a.store.DB.QueryRow(`
- `carbonstack-cypher/internal/httpapi/api.go:525` — `tx, err := a.store.DB.Begin()`
- `carbonstack-cypher/internal/httpapi/api.go:592` — `err := a.store.DB.QueryRow(`
- `carbonstack-cypher/internal/config/config.go:7` — `DBPath        string`
- `carbonstack-cypher/internal/config/config.go:15` — `DBPath:        getEnv("CYPHER_DB", "cypher.db"),`
- `carbonstack-cypher/internal/config/config.go:16` — `MigrationsDir: getEnv("CYPHER_MIGRATIONS", "migrations"),`
- `carbonstack-cypher/internal/db/db.go:5` — `"database/sql"`
- `carbonstack-cypher/internal/db/db.go:14` — `_ "modernc.org/sqlite"`
- `carbonstack-cypher/internal/db/db.go:18` — `DB *sql.DB`
- `carbonstack-cypher/internal/db/db.go:22` — `conn, err := sql.Open("sqlite", path)`
- `carbonstack-cypher/internal/db/db.go:32` — `return &Store{DB: conn}, nil`
- `carbonstack-cypher/internal/db/db.go:36` — `return s.DB.Close()`
- `carbonstack-cypher/internal/db/db.go:42` — `return fmt.Errorf("read migrations dir: %w", err)`
- `carbonstack-cypher/internal/db/db.go:63` — `if _, err := s.DB.Exec(string(body)); err != nil {`
- `carbonstack-cypher/internal/db/db.go:79` — `err := s.DB.QueryRow("SELECT invite_id FROM invites WHERE invite_code_hash = ? LIMIT 1", hash).Scan(&existing)`
- `carbonstack-cypher/internal/db/db.go:87` — `_, err = s.DB.Exec(`
- `carbonstack-cypher/cmd/cypher/main.go:15` — `store, err := db.Open(cfg.DBPath)`
- `carbonstack-cypher/cmd/cypher/main.go:17` — `log.Fatalf("open database: %v", err)`
- `carbonstack-cypher/cmd/cypher/main.go:22` — `log.Fatalf("run migrations: %v", err)`
- `carbonstack-cypher/cmd/cypher/main.go:32` — `log.Printf("database: %s", cfg.DBPath)`

Invite/bootstrap related hits:

- `carbonstack-cypher/README.md:17` — `- development invite/account/device registration;`
- `carbonstack-cypher/README.md:60` — `POST /v0/invites/claim`
- `carbonstack-cypher/docs/08-api-contract-v0.md:43` — `## Invite claim`
- `carbonstack-cypher/docs/08-api-contract-v0.md:45` — `POST /v0/invites/claim`
- `carbonstack-cypher/docs/08-api-contract-v0.md:47` — `Purpose: claim a development invite and create an account.`
- `carbonstack-cypher/docs/05-mvp-roadmap.md:13` — `- development invite/account/device records;`
- `carbonstack-cypher/docs/06-phase1-vertical-slice.md:15` — `- invite-only registration`
- `carbonstack-cypher/docs/06-phase1-vertical-slice.md:42` — `- Invite`
- `carbonstack-cypher/docs/06-phase1-vertical-slice.md:51` — `2. Operator creates or seeds invite codes.`
- `carbonstack-cypher/docs/06-phase1-vertical-slice.md:52` — `3. Client A claims invite and registers account/device.`
- `carbonstack-cypher/docs/06-phase1-vertical-slice.md:53` — `4. Client B claims invite and registers account/device.`
- `carbonstack-cypher/docs/01-cypher-architecture.md:18` — `- supporting invite-only registration`
- `carbonstack-cypher/docs/01-cypher-architecture.md:64` — `- invite-only registration`
- `carbonstack-cypher/docs/01-cypher-architecture.md:85` — `- private invite flows`
- `carbonstack-cypher/docs/01-cypher-architecture.md:97` — `- invite creation`
- `carbonstack-cypher/docs/07-data-model-v0.md:32` — `### invites`
- `carbonstack-cypher/docs/07-data-model-v0.md:34` — `Purpose: allow development invite/account creation without open signup.`
- `carbonstack-cypher/docs/07-data-model-v0.md:38` — `- `invite_id``
- `carbonstack-cypher/docs/07-data-model-v0.md:39` — `- `invite_code_hash``
- `carbonstack-cypher/docs/07-data-model-v0.md:45` — `Invite codes are not a production authentication system.`
- `carbonstack-cypher/docs/04-storage-model.md:38` — `### invites`
- `carbonstack-cypher/docs/04-storage-model.md:40` — `Development invite records.`
- `carbonstack-cypher/docs/03-api-surface.md:16` — `POST /v0/dev/invites`
- `carbonstack-cypher/docs/03-api-surface.md:17` — `POST /v0/invites/claim`
- `carbonstack-cypher/docs/03-api-surface.md:31` — `## Development invite/account/device routes`
- `carbonstack-cypher/migrations/001_init.sql:1` — `CREATE TABLE IF NOT EXISTS invites (`
- `carbonstack-cypher/migrations/001_init.sql:2` — `invite_id TEXT PRIMARY KEY,`
- `carbonstack-cypher/migrations/001_init.sql:3` — `invite_code_hash TEXT NOT NULL UNIQUE,`
- `carbonstack-cypher/internal/httpapi/api_test.go:41` — `alice := claimInvite(t, server.URL, "dev-invite", "alice-test")`
- `carbonstack-cypher/internal/httpapi/api_test.go:43` — `bobInvite := createDevInvite(t, server.URL, "bob-test-invite")`
- `carbonstack-cypher/internal/httpapi/api_test.go:44` — `if bobInvite.InviteCode != "bob-test-invite" {`
- `carbonstack-cypher/internal/httpapi/api_test.go:45` — `t.Fatalf("expected bob-test-invite, got %q", bobInvite.InviteCode)`
- `carbonstack-cypher/internal/httpapi/api_test.go:48` — `bob := claimInvite(t, server.URL, "bob-test-invite", "bob-test")`
- `carbonstack-cypher/internal/httpapi/api_test.go:119` — `alice := claimInvite(t, server.URL, "dev-invite", "alice-test")`
- `carbonstack-cypher/internal/httpapi/api_test.go:120` — `createDevInvite(t, server.URL, "bob-test-invite")`
- `carbonstack-cypher/internal/httpapi/api_test.go:121` — `bob := claimInvite(t, server.URL, "bob-test-invite", "bob-test")`
- `carbonstack-cypher/internal/httpapi/api_test.go:155` — `alice := claimInvite(t, server.URL, "dev-invite", "alice-test")`
- `carbonstack-cypher/internal/httpapi/api_test.go:185` — `alice := claimInvite(t, server.URL, "dev-invite", "alice-test")`
- `carbonstack-cypher/internal/httpapi/api_test.go:187` — `createDevInvite(t, server.URL, "bob-test-invite")`
- `carbonstack-cypher/internal/httpapi/api_test.go:188` — `bob := claimInvite(t, server.URL, "bob-test-invite", "bob-test")`
- `carbonstack-cypher/internal/httpapi/api_test.go:190` — `createDevInvite(t, server.URL, "mallory-test-invite")`
- `carbonstack-cypher/internal/httpapi/api_test.go:191` — `mallory := claimInvite(t, server.URL, "mallory-test-invite", "mallory-test")`
- `carbonstack-cypher/internal/httpapi/api_test.go:236` — `if err := store.SeedDevInvite("dev-invite"); err != nil {`
- `carbonstack-cypher/internal/httpapi/api_test.go:237` — `t.Fatalf("seed dev invite: %v", err)`
- `carbonstack-cypher/internal/httpapi/api_test.go:244` — `type claimInviteResponse struct {`
- `carbonstack-cypher/internal/httpapi/api_test.go:249` — `type devInviteResponse struct {`
- `carbonstack-cypher/internal/httpapi/api_test.go:250` — `InviteID   string `json:"invite_id"``
- `carbonstack-cypher/internal/httpapi/api_test.go:251` — `InviteCode string `json:"invite_code"``
- `carbonstack-cypher/internal/httpapi/api_test.go:306` — `func createDevInvite(t *testing.T, serverURL string, inviteCode string) devInviteResponse {`
- `carbonstack-cypher/internal/httpapi/api_test.go:309` — `var resp devInviteResponse`
- `carbonstack-cypher/internal/httpapi/api_test.go:311` — `"invite_code": inviteCode,`
- `carbonstack-cypher/internal/httpapi/api_test.go:314` — `doPost(t, serverURL+"/v0/dev/invites", body, http.StatusCreated, &resp)`
- `carbonstack-cypher/internal/httpapi/api_test.go:318` — `func claimInvite(t *testing.T, serverURL string, inviteCode string, displayName string) claimInviteResponse {`
- `carbonstack-cypher/internal/httpapi/api_test.go:321` — `var resp claimInviteResponse`
- `carbonstack-cypher/internal/httpapi/api_test.go:323` — `"invite_code":  inviteCode,`
- `carbonstack-cypher/internal/httpapi/api_test.go:327` — `doPost(t, serverURL+"/v0/invites/claim", body, http.StatusCreated, &resp)`
- `carbonstack-cypher/internal/httpapi/api_test.go:468` — `alice := claimInvite(t, server.URL, "dev-invite", "alice-test")`
- `carbonstack-cypher/internal/httpapi/api_test.go:469` — `createDevInvite(t, server.URL, "bob-test-invite")`
- `carbonstack-cypher/internal/httpapi/api_test.go:470` — `bob := claimInvite(t, server.URL, "bob-test-invite", "bob-test")`
- `carbonstack-cypher/internal/httpapi/api_test.go:504` — `func TestDevInviteCanAutogenerateCode(t *testing.T) {`
- `carbonstack-cypher/internal/httpapi/api_test.go:508` — `resp := createDevInvite(t, server.URL, "")`
- `carbonstack-cypher/internal/httpapi/api_test.go:509` — `if !strings.HasPrefix(resp.InviteCode, "dev-") {`
- `carbonstack-cypher/internal/httpapi/api_test.go:510` — `t.Fatalf("expected autogenerated dev invite, got %q", resp.InviteCode)`
- `carbonstack-cypher/internal/httpapi/api_test.go:542` — `alice := claimInvite(t, server.URL, "dev-invite", "alice-test")`
- `carbonstack-cypher/internal/httpapi/api_test.go:544` — `createDevInvite(t, server.URL, "bob-test-invite")`
- `carbonstack-cypher/internal/httpapi/api_test.go:545` — `bob := claimInvite(t, server.URL, "bob-test-invite", "bob-test")`
- `carbonstack-cypher/internal/httpapi/api_test.go:594` — `alice := claimInvite(t, server.URL, "dev-invite", "alice-test")`
- `carbonstack-cypher/internal/httpapi/api_test.go:596` — `createDevInvite(t, server.URL, "bob-test-invite")`
- `carbonstack-cypher/internal/httpapi/api_test.go:597` — `bob := claimInvite(t, server.URL, "bob-test-invite", "bob-test")`
- `carbonstack-cypher/internal/httpapi/api_test.go:623` — `alice := claimInvite(t, server.URL, "dev-invite", "alice-test")`
- `carbonstack-cypher/internal/httpapi/api_test.go:625` — `createDevInvite(t, server.URL, "bob-test-invite")`
- `carbonstack-cypher/internal/httpapi/api_test.go:626` — `bob := claimInvite(t, server.URL, "bob-test-invite", "bob-test")`
- `carbonstack-cypher/internal/httpapi/api.go:42` — `mux.HandleFunc("POST /v0/dev/invites", a.createDevInvite)`
- `carbonstack-cypher/internal/httpapi/api.go:43` — `mux.HandleFunc("POST /v0/invites/claim", a.claimInvite)`
- `carbonstack-cypher/internal/httpapi/api.go:61` — `type createDevInviteRequest struct {`
- `carbonstack-cypher/internal/httpapi/api.go:62` — `InviteCode string `json:"invite_code"``
- `carbonstack-cypher/internal/httpapi/api.go:65` — `func (a *API) createDevInvite(w http.ResponseWriter, r *http.Request) {`
- `carbonstack-cypher/internal/httpapi/api.go:71` — `var req createDevInviteRequest`
- `carbonstack-cypher/internal/httpapi/api.go:76` — `req.InviteCode = strings.TrimSpace(req.InviteCode)`
- `carbonstack-cypher/internal/httpapi/api.go:77` — `if req.InviteCode == "" {`
- `carbonstack-cypher/internal/httpapi/api.go:78` — `req.InviteCode = "dev-" + uuid.NewString()`
- `carbonstack-cypher/internal/httpapi/api.go:81` — `codeHash := db.HashInviteCode(req.InviteCode)`
- `carbonstack-cypher/internal/httpapi/api.go:85` — `"SELECT invite_id FROM invites WHERE invite_code_hash = ? LIMIT 1",`
- `carbonstack-cypher/internal/httpapi/api.go:90` — `writeError(w, http.StatusConflict, "invite_exists", "invite code already exists")`
- `carbonstack-cypher/internal/httpapi/api.go:98` — `inviteID := uuid.NewString()`
- `carbonstack-cypher/internal/httpapi/api.go:102` — `"INSERT INTO invites (invite_id, invite_code_hash, created_at) VALUES (?, ?, ?)",`
- `carbonstack-cypher/internal/httpapi/api.go:103` — `inviteID,`
- `carbonstack-cypher/internal/httpapi/api.go:113` — `"invite_id":   inviteID,`
- `carbonstack-cypher/internal/httpapi/api.go:114` — `"invite_code": req.InviteCode,`
- `carbonstack-cypher/internal/httpapi/api.go:119` — `type claimInviteRequest struct {`
- `carbonstack-cypher/internal/httpapi/api.go:120` — `InviteCode  string `json:"invite_code"``
- `carbonstack-cypher/internal/httpapi/api.go:124` — `func (a *API) claimInvite(w http.ResponseWriter, r *http.Request) {`
- `carbonstack-cypher/internal/httpapi/api.go:125` — `var req claimInviteRequest`
- `carbonstack-cypher/internal/httpapi/api.go:130` — `req.InviteCode = strings.TrimSpace(req.InviteCode)`
- `carbonstack-cypher/internal/httpapi/api.go:133` — `if req.InviteCode == "" \|\| req.DisplayName == "" {`
- `carbonstack-cypher/internal/httpapi/api.go:134` — `writeError(w, http.StatusBadRequest, "invalid_request", "invite_code and display_name are required")`
- `carbonstack-cypher/internal/httpapi/api.go:145` — `codeHash := db.HashInviteCode(req.InviteCode)`
- `carbonstack-cypher/internal/httpapi/api.go:147` — `var inviteID string`
- `carbonstack-cypher/internal/httpapi/api.go:152` — `"SELECT invite_id, claimed_at, disabled_at FROM invites WHERE invite_code_hash = ? LIMIT 1",`
- `carbonstack-cypher/internal/httpapi/api.go:154` — `).Scan(&inviteID, &claimedAt, &disabledAt)`
- `carbonstack-cypher/internal/httpapi/api.go:157` — `writeError(w, http.StatusUnauthorized, "invalid_invite", "invite not found")`
- `carbonstack-cypher/internal/httpapi/api.go:165` — `writeError(w, http.StatusConflict, "invite_claimed", "invite already claimed")`
- `carbonstack-cypher/internal/httpapi/api.go:169` — `writeError(w, http.StatusForbidden, "invite_disabled", "invite disabled")`
- `carbonstack-cypher/internal/httpapi/api.go:187` — `"UPDATE invites SET claimed_at = ?, claimed_by_account_id = ? WHERE invite_id = ?",`
- `carbonstack-cypher/internal/httpapi/api.go:190` — `inviteID,`
- `carbonstack-cypher/internal/config/config.go:9` — `DevInviteCode string`
- `carbonstack-cypher/internal/config/config.go:17` — `DevInviteCode: getEnv("CYPHER_DEV_INVITE", "dev-invite"),`
- `carbonstack-cypher/internal/db/db.go:71` — `func (s *Store) SeedDevInvite(inviteCode string) error {`
- `carbonstack-cypher/internal/db/db.go:72` — `if inviteCode == "" {`
- `carbonstack-cypher/internal/db/db.go:76` — `hash := HashInviteCode(inviteCode)`
- `carbonstack-cypher/internal/db/db.go:79` — `err := s.DB.QueryRow("SELECT invite_id FROM invites WHERE invite_code_hash = ? LIMIT 1", hash).Scan(&existing)`
- `carbonstack-cypher/internal/db/db.go:88` — `"INSERT INTO invites (invite_id, invite_code_hash, created_at) VALUES (?, ?, ?)",`
- `carbonstack-cypher/internal/db/db.go:96` — `func HashInviteCode(code string) string {`
- `carbonstack-cypher/cmd/cypher/main.go:25` — `if err := store.SeedDevInvite(cfg.DevInviteCode); err != nil {`
- `carbonstack-cypher/cmd/cypher/main.go:26` — `log.Fatalf("seed dev invite: %v", err)`
- `carbonstack-cypher/cmd/cypher/main.go:29` — `api := httpapi.New(store, cfg.DevInviteCode != "")`
- `carbonstack-cypher/cmd/cypher/main.go:33` — `log.Printf("dev invite enabled: %t", cfg.DevInviteCode != "")`

## 8. Cypher API / Relay Surface

Route/API related hits:

- `carbonstack-cypher/README.md:1` — `# CarbonStackCypher`
- `carbonstack-cypher/README.md:3` — `CarbonStackCypher is the experimental relay/storage server for CarbonStack.`
- `carbonstack-cypher/README.md:7` — `The current validated role of Cypher is to relay CarbonStackComms/OpenMLS sidecar artifacts through a simple HTTP JSON envelope API.`
- `carbonstack-cypher/README.md:9` — `_Related repositories: [carbonstack](https://git.bitcrusher32.win/bitcrusher32/carbonstack) / [carbonstack-comms](https://git.bitcrusher32.win/bitcrusher32/carbonstack-comms) / [carbonstack-os](https://git.bitcrusher32.win/bitcrusher32/carbonstack-os)_`
- `carbonstack-cypher/README.md:17` — `- development invite/account/device registration;`
- `carbonstack-cypher/README.md:19` — `- recipient inbox listing;`
- `carbonstack-cypher/README.md:20` — `- envelope acknowledgement;`
- `carbonstack-cypher/README.md:26` — `carbonstack.mls.keypackage.v0`
- `carbonstack-cypher/README.md:27` — `carbonstack.mls.welcome.v0`
- `carbonstack-cypher/README.md:28` — `carbonstack.mls.application-message.v0`
- `carbonstack-cypher/README.md:32` — `carbonstack-openmls-sidecar-v0`
- `carbonstack-cypher/README.md:36` — `carbonstack.message.text.stub.v0`
- `carbonstack-cypher/README.md:59` — `GET  /v0/health`
- `carbonstack-cypher/README.md:60` — `POST /v0/invites/claim`
- `carbonstack-cypher/README.md:61` — `POST /v0/devices/register`
- `carbonstack-cypher/README.md:62` — `GET  /v0/accounts/{account_id}/devices`
- `carbonstack-cypher/README.md:63` — `POST /v0/envelopes`
- `carbonstack-cypher/README.md:64` — `GET  /v0/devices/{device_id}/envelopes`
- `carbonstack-cypher/README.md:65` — `POST /v0/envelopes/{envelope_id}/ack`
- `carbonstack-cypher/README.md:84` — `For the full CarbonStack OpenMLS relay proof, use the runbook in the main `carbonstack` repo:`
- `carbonstack-cypher/README.md:86` — `docs/113-experimental-backbone-deployability-runbook-v0.md`
- `carbonstack-cypher/README.md:97` — `- rollback/replay safety against a malicious server;`
- `carbonstack-cypher/README.md:101` — `Cypher is a relay/storage component inside the current experimental backbone.`
- `carbonstack-cypher/docs/08-api-contract-v0.md:1` — `# CarbonStackCypher API Contract v0`
- `carbonstack-cypher/docs/08-api-contract-v0.md:4` — `Component: CarbonStackCypher`
- `carbonstack-cypher/docs/08-api-contract-v0.md:7` — `This document describes the current HTTP JSON API used by CarbonStackCypher.`
- `carbonstack-cypher/docs/08-api-contract-v0.md:9` — `The API is stable enough for the current CarbonStackComms smoke harness. It is not a stable public protocol.`
- `carbonstack-cypher/docs/08-api-contract-v0.md:31` — `GET /v0/health`
- `carbonstack-cypher/docs/08-api-contract-v0.md:39` — `"service": "carbonstack-cypher",`
- `carbonstack-cypher/docs/08-api-contract-v0.md:45` — `POST /v0/invites/claim`
- `carbonstack-cypher/docs/08-api-contract-v0.md:53` — `POST /v0/devices/register`
- `carbonstack-cypher/docs/08-api-contract-v0.md:61` — `GET /v0/accounts/{account_id}/devices`
- `carbonstack-cypher/docs/08-api-contract-v0.md:65` — `## Envelope submit`
- `carbonstack-cypher/docs/08-api-contract-v0.md:67` — `POST /v0/envelopes`
- `carbonstack-cypher/docs/08-api-contract-v0.md:69` — `Purpose: submit an opaque envelope for a recipient device.`
- `carbonstack-cypher/docs/08-api-contract-v0.md:82` — `carbonstack.mls.keypackage.v0`
- `carbonstack-cypher/docs/08-api-contract-v0.md:83` — `carbonstack.mls.welcome.v0`
- `carbonstack-cypher/docs/08-api-contract-v0.md:84` — `carbonstack.mls.application-message.v0`
- `carbonstack-cypher/docs/08-api-contract-v0.md:88` — `carbonstack-openmls-sidecar-v0`
- `carbonstack-cypher/docs/08-api-contract-v0.md:92` — `carbonstack.message.text.stub.v0`
- `carbonstack-cypher/docs/08-api-contract-v0.md:110` — `GET /v0/devices/{device_id}/envelopes`
- `carbonstack-cypher/docs/08-api-contract-v0.md:130` — `POST /v0/envelopes/{envelope_id}/ack`
- `carbonstack-cypher/docs/08-api-contract-v0.md:142` — `- rejects wrong-recipient ack;`
- `carbonstack-cypher/docs/08-api-contract-v0.md:144` — `- sets or returns `delivery_state = acknowledged`.`
- `carbonstack-cypher/docs/08-api-contract-v0.md:148` — `In the current Comms proof, ack occurs only after recipient-side OpenMLS sidecar consume succeeds.`
- `carbonstack-cypher/docs/08-api-contract-v0.md:150` — `Cypher itself does not know OpenMLS consume state. It records the ack request and delivery state.`
- `carbonstack-cypher/docs/05-mvp-roadmap.md:1` — `# CarbonStackCypher MVP Roadmap`
- `carbonstack-cypher/docs/05-mvp-roadmap.md:4` — `Component: CarbonStackCypher`
- `carbonstack-cypher/docs/05-mvp-roadmap.md:9` — `Cypher currently implements the relay/server scaffold needed by the CarbonStack experimental backbone:`
- `carbonstack-cypher/docs/05-mvp-roadmap.md:13` — `- development invite/account/device records;`
- `carbonstack-cypher/docs/05-mvp-roadmap.md:15` — `- recipient inbox listing;`
- `carbonstack-cypher/docs/05-mvp-roadmap.md:16` — `- envelope ack;`
- `carbonstack-cypher/docs/05-mvp-roadmap.md:36` — `- inbox/ack semantics cleanup;`
- `carbonstack-cypher/docs/05-mvp-roadmap.md:40` — `- release-facing docs alignment with `carbonstack`;`
- `carbonstack-cypher/docs/05-mvp-roadmap.md:50` — `- hostile-server rollback/replay harnesses;`
- `carbonstack-cypher/docs/05-mvp-roadmap.md:63` — `Cypher is one component in the current experimental CarbonStack backbone.`
- `carbonstack-cypher/docs/06-phase1-vertical-slice.md:1` — `# CarbonStackCypher Phase 1 Vertical Slice`
- `carbonstack-cypher/docs/06-phase1-vertical-slice.md:7` — `CarbonStackCypher Phase 1 is a local-first relay skeleton for storing and retrieving encrypted-envelope-shaped records.`
- `carbonstack-cypher/docs/06-phase1-vertical-slice.md:13` — `Build the smallest server that can support a CarbonStackComms CLI lifecycle:`
- `carbonstack-cypher/docs/06-phase1-vertical-slice.md:20` — `- envelope acknowledgement`
- `carbonstack-cypher/docs/06-phase1-vertical-slice.md:24` — `CarbonStackCypher routes and stores envelopes.`
- `carbonstack-cypher/docs/06-phase1-vertical-slice.md:26` — `CarbonStackCypher must not require message plaintext.`
- `carbonstack-cypher/docs/06-phase1-vertical-slice.md:46` — `- Delivery acknowledgement`
- `carbonstack-cypher/docs/06-phase1-vertical-slice.md:52` — `3. Client A claims invite and registers account/device.`
- `carbonstack-cypher/docs/06-phase1-vertical-slice.md:53` — `4. Client B claims invite and registers account/device.`
- `carbonstack-cypher/docs/06-phase1-vertical-slice.md:54` — `5. Client A submits an envelope addressed to Client B device.`
- `carbonstack-cypher/docs/06-phase1-vertical-slice.md:56` — `7. Client B polls/retrieves queued envelopes.`
- `carbonstack-cypher/docs/06-phase1-vertical-slice.md:57` — `8. Client B acknowledges envelope receipt.`
- `carbonstack-cypher/docs/06-phase1-vertical-slice.md:74` — `## Recommended Implementation Stack`
- `carbonstack-cypher/docs/01-cypher-architecture.md:1` — `# CarbonStackCypher Architecture`
- `carbonstack-cypher/docs/01-cypher-architecture.md:3` — `CarbonStackCypher is the hostile-server relay and storage stack for CarbonStack.`
- `carbonstack-cypher/docs/01-cypher-architecture.md:7` — `CarbonStackCypher must be designed as if it may be compromised, malicious, misconfigured, legally pressured, or operated by an untrusted party.`
- `carbonstack-cypher/docs/01-cypher-architecture.md:11` — `CarbonStackCypher is responsible for:`
- `carbonstack-cypher/docs/01-cypher-architecture.md:26` — `CarbonStackCypher must not be responsible for:`
- `carbonstack-cypher/docs/01-cypher-architecture.md:41` — `Clients must assume CarbonStackCypher may attempt to:`
- `carbonstack-cypher/docs/01-cypher-architecture.md:47` — `- roll back visible state`
- `carbonstack-cypher/docs/01-cypher-architecture.md:56` — `CarbonStackCypher must be architected so these attacks are either impossible, detectable by clients, or documented as limitations.`
- `carbonstack-cypher/docs/01-cypher-architecture.md:125` — `CarbonStackCypher is infrastructure, not authority.`
- `carbonstack-cypher/docs/07-data-model-v0.md:1` — `# CarbonStackCypher Data Model v0`
- `carbonstack-cypher/docs/07-data-model-v0.md:4` — `Component: CarbonStackCypher`
- `carbonstack-cypher/docs/07-data-model-v0.md:7` — `This document describes the current SQLite data model used by CarbonStackCypher.`
- `carbonstack-cypher/docs/07-data-model-v0.md:109` — `- `acknowledged``
- `carbonstack-cypher/docs/07-data-model-v0.md:111` — ``queued` means the envelope is available through the recipient inbox route.`
- `carbonstack-cypher/docs/07-data-model-v0.md:113` — ``acknowledged` means Cypher accepted a recipient-device ack for the envelope.`
- `carbonstack-cypher/docs/07-data-model-v0.md:121` — `First correct-recipient ack records an ack event and marks the envelope acknowledged.`
- `carbonstack-cypher/docs/07-data-model-v0.md:123` — `Repeated correct-recipient ack returns acknowledged without creating a new semantic state.`
- `carbonstack-cypher/docs/07-data-model-v0.md:125` — `Wrong-recipient ack is rejected.`
- `carbonstack-cypher/docs/07-data-model-v0.md:127` — `Unknown-envelope ack is rejected.`
- `carbonstack-cypher/docs/07-data-model-v0.md:129` — `Missing-recipient ack is rejected.`
- `carbonstack-cypher/docs/07-data-model-v0.md:131` — `The current consume-then-ack rule is enforced by Comms tests and harness behavior. Cypher does not know OpenMLS sidecar consume state.`
- `carbonstack-cypher/docs/04-storage-model.md:1` — `# CarbonStackCypher Storage Model`
- `carbonstack-cypher/docs/04-storage-model.md:4` — `Component: CarbonStackCypher`
- `carbonstack-cypher/docs/04-storage-model.md:15` — `## Current storage backend`
- `carbonstack-cypher/docs/04-storage-model.md:17` — `Current backend:`
- `carbonstack-cypher/docs/04-storage-model.md:25` — `- experimental backbone validation.`
- `carbonstack-cypher/docs/04-storage-model.md:79` — `### envelope_acks`
- `carbonstack-cypher/docs/04-storage-model.md:83` — `Current ack semantics:`
- `carbonstack-cypher/docs/04-storage-model.md:85` — `- same-recipient ack is idempotent at the API layer;`
- `carbonstack-cypher/docs/04-storage-model.md:86` — `- wrong-recipient ack is rejected;`
- `carbonstack-cypher/docs/04-storage-model.md:87` — `- ack sets the envelope delivery state to `acknowledged`;`
- `carbonstack-cypher/docs/04-storage-model.md:88` — `- inbox returns queued envelopes only.`
- `carbonstack-cypher/docs/04-storage-model.md:90` — `Ack records are server records of a recipient-device ack request. They are not proof that Cypher independently verified sidecar consume.`
- `carbonstack-cypher/docs/04-storage-model.md:94` — `Future retention policy may delete acknowledged envelopes or expire old queued envelopes.`
- `carbonstack-cypher/docs/03-api-surface.md:1` — `# CarbonStackCypher API Surface`
- `carbonstack-cypher/docs/03-api-surface.md:4` — `Component: CarbonStackCypher`
- `carbonstack-cypher/docs/03-api-surface.md:15` — `GET  /v0/health`
- `carbonstack-cypher/docs/03-api-surface.md:16` — `POST /v0/dev/invites`
- `carbonstack-cypher/docs/03-api-surface.md:17` — `POST /v0/invites/claim`
- `carbonstack-cypher/docs/03-api-surface.md:18` — `POST /v0/devices/register`
- `carbonstack-cypher/docs/03-api-surface.md:19` — `GET  /v0/accounts/{account_id}/devices`
- `carbonstack-cypher/docs/03-api-surface.md:20` — `POST /v0/envelopes`
- `carbonstack-cypher/docs/03-api-surface.md:21` — `GET  /v0/devices/{device_id}/envelopes`
- `carbonstack-cypher/docs/03-api-surface.md:22` — `POST /v0/envelopes/{envelope_id}/ack`
- `carbonstack-cypher/docs/03-api-surface.md:31` — `## Development invite/account/device routes`
- `carbonstack-cypher/docs/03-api-surface.md:40` — `## Envelope submit`
- `carbonstack-cypher/docs/03-api-surface.md:44` — `POST /v0/envelopes`
- `carbonstack-cypher/docs/03-api-surface.md:62` — `GET /v0/devices/{device_id}/envelopes`
- `carbonstack-cypher/docs/03-api-surface.md:68` — `Current inbox behavior:`
- `carbonstack-cypher/docs/03-api-surface.md:72` — `It does not return acknowledged envelopes.`
- `carbonstack-cypher/docs/03-api-surface.md:74` — `Inbox retrieval is not ack.`
- `carbonstack-cypher/docs/03-api-surface.md:82` — `POST /v0/envelopes/{envelope_id}/ack`
- `carbonstack-cypher/docs/03-api-surface.md:88` — `Current ack behavior:`
- `carbonstack-cypher/docs/03-api-surface.md:92` — `- rejects wrong-recipient ack;`
- `carbonstack-cypher/docs/03-api-surface.md:94` — `- sets or returns `delivery_state = acknowledged`.`
- `carbonstack-cypher/docs/03-api-surface.md:98` — `In the current CarbonStackComms proof, Comms sends ack only after the relevant sidecar consume command succeeds.`
- `carbonstack-cypher/docs/02-envelope-model.md:1` — `# CarbonStackCypher Envelope Model`
- `carbonstack-cypher/docs/02-envelope-model.md:4` — `Component: CarbonStackCypher`
- `carbonstack-cypher/docs/02-envelope-model.md:33` — `For OpenMLS relay artifacts, this name is historical and imperfect. The payload may be a KeyPackage artifact, Welcome artifact, or application-message artifact.`
- `carbonstack-cypher/docs/02-envelope-model.md:39` — `carbonstack.mls.keypackage.v0`
- `carbonstack-cypher/docs/02-envelope-model.md:40` — `carbonstack.mls.welcome.v0`
- `carbonstack-cypher/docs/02-envelope-model.md:41` — `carbonstack.mls.application-message.v0`
- `carbonstack-cypher/docs/02-envelope-model.md:45` — `carbonstack.message.text.stub.v0`
- `carbonstack-cypher/docs/02-envelope-model.md:53` — `carbonstack-openmls-sidecar-v0`
- `carbonstack-cypher/docs/02-envelope-model.md:59` — `The OpenMLS relay protocol version is a CarbonStack compatibility label. It is not a claim of generic OpenMLS standard compatibility.`
- `carbonstack-cypher/docs/02-envelope-model.md:83` — `acknowledged`
- `carbonstack-cypher/docs/02-envelope-model.md:85` — ``queued` means the envelope is available through the recipient inbox route.`
- `carbonstack-cypher/docs/02-envelope-model.md:87` — ``acknowledged` means Cypher accepted a recipient-device ack for the envelope.`
- `carbonstack-cypher/docs/02-envelope-model.md:89` — `Cypher does not know whether the sidecar consumed the artifact. Comms decides when to ack.`
- `carbonstack-cypher/migrations/001_init.sql:42` — `CREATE TABLE IF NOT EXISTS envelope_acks (`
- `carbonstack-cypher/migrations/001_init.sql:43` — `ack_id TEXT PRIMARY KEY,`
- `carbonstack-cypher/migrations/001_init.sql:46` — `acknowledged_at TEXT NOT NULL,`
- `carbonstack-cypher/internal/httpapi/api_test.go:1` — `package httpapi_test`
- `carbonstack-cypher/internal/httpapi/api_test.go:15` — `"git.bitcrusher32.win/bitcrusher32/carbonstack-cypher/internal/db"`
- `carbonstack-cypher/internal/httpapi/api_test.go:16` — `"git.bitcrusher32.win/bitcrusher32/carbonstack-cypher/internal/httpapi"`
- `carbonstack-cypher/internal/httpapi/api_test.go:24` — `doGet(t, server.URL+"/v0/health", http.StatusOK, &resp)`
- `carbonstack-cypher/internal/httpapi/api_test.go:29` — `if resp["service"] != "carbonstack-cypher" {`
- `carbonstack-cypher/internal/httpapi/api_test.go:30` — `t.Fatalf("expected service carbonstack-cypher, got %q", resp["service"])`
- `carbonstack-cypher/internal/httpapi/api_test.go:72` — `envelope := submitEnvelope(t, server.URL, aliceDevice.DeviceID, bobDevice.DeviceID, ciphertextB64)`
- `carbonstack-cypher/internal/httpapi/api_test.go:80` — `inbox := getInbox(t, server.URL, bobDevice.DeviceID)`
- `carbonstack-cypher/internal/httpapi/api_test.go:81` — `if len(inbox.Envelopes) != 1 {`
- `carbonstack-cypher/internal/httpapi/api_test.go:82` — `t.Fatalf("expected 1 queued envelope, got %d", len(inbox.Envelopes))`
- `carbonstack-cypher/internal/httpapi/api_test.go:85` — `gotEnvelope := inbox.Envelopes[0]`
- `carbonstack-cypher/internal/httpapi/api_test.go:104` — `ack := ackEnvelope(t, server.URL, envelope.EnvelopeID, bobDevice.DeviceID)`
- `carbonstack-cypher/internal/httpapi/api_test.go:105` — `if ack.DeliveryState != "acknowledged" {`
- `carbonstack-cypher/internal/httpapi/api_test.go:106` — `t.Fatalf("expected acknowledged, got %q", ack.DeliveryState)`
- `carbonstack-cypher/internal/httpapi/api_test.go:109` — `inboxAfterAck := getInbox(t, server.URL, bobDevice.DeviceID)`
- `carbonstack-cypher/internal/httpapi/api_test.go:110` — `if len(inboxAfterAck.Envelopes) != 0 {`
- `carbonstack-cypher/internal/httpapi/api_test.go:111` — `t.Fatalf("expected empty inbox after ack, got %d envelopes", len(inboxAfterAck.Envelopes))`
- `carbonstack-cypher/internal/httpapi/api_test.go:126` — `ciphertextB64 := base64.StdEncoding.EncodeToString([]byte("idempotent ack test"))`
- `carbonstack-cypher/internal/httpapi/api_test.go:127` — `envelope := submitEnvelope(t, server.URL, aliceDevice.DeviceID, bobDevice.DeviceID, ciphertextB64)`
- ... 105 more hits omitted from this generated recon doc.

Recon interpretation:

- Cypher remains the relay/storage component, not a plaintext trust root.
- Deployability planning should treat Cypher as a local operator process first.
- The first local profile should prefer `127.0.0.1` binding and avoid `0.0.0.0`, LAN exposure, systemd, public ingress, and cloudflared.
- Current dev invite behavior should remain bootstrap plumbing, not the final relay-space invite, membership invite, device enrollment, or operator authority model.

## 9. SQLite and Migration Surface

Migration files / summary:

- `carbonstack-cypher/migrations/001_init.sql` — uses CREATE TABLE IF NOT EXISTS
- `carbonstack-cypher/migrations/002_envelope_payload_metadata.sql` — contains ALTER TABLE

Migration-related hits:

- `carbonstack-cypher/README.md:71` — `Current migrations:`
- `carbonstack-cypher/README.md:73` — `migrations/001_init.sql`
- `carbonstack-cypher/README.md:74` — `migrations/002_envelope_payload_metadata.sql`
- `carbonstack-cypher/docs/05-mvp-roadmap.md:12` — `- SQLite migrations;`
- `carbonstack-cypher/docs/05-mvp-roadmap.md:53` — `- production migration strategy.`
- `carbonstack-cypher/docs/07-data-model-v0.md:25` — `Current migrations:`
- `carbonstack-cypher/docs/07-data-model-v0.md:27` — `migrations/001_init.sql`
- `carbonstack-cypher/docs/07-data-model-v0.md:28` — `migrations/002_envelope_payload_metadata.sql`
- `carbonstack-cypher/docs/04-storage-model.md:29` — `## Current migrations`
- `carbonstack-cypher/docs/04-storage-model.md:31` — `Current migrations:`
- `carbonstack-cypher/docs/04-storage-model.md:33` — `migrations/001_init.sql`
- `carbonstack-cypher/docs/04-storage-model.md:34` — `migrations/002_envelope_payload_metadata.sql`
- `carbonstack-cypher/migrations/001_init.sql:1` — `CREATE TABLE IF NOT EXISTS invites (`
- `carbonstack-cypher/migrations/001_init.sql:10` — `CREATE TABLE IF NOT EXISTS accounts (`
- `carbonstack-cypher/migrations/001_init.sql:17` — `CREATE TABLE IF NOT EXISTS devices (`
- `carbonstack-cypher/migrations/001_init.sql:28` — `CREATE TABLE IF NOT EXISTS envelopes (`
- `carbonstack-cypher/migrations/001_init.sql:42` — `CREATE TABLE IF NOT EXISTS envelope_acks (`
- `carbonstack-cypher/migrations/002_envelope_payload_metadata.sql:1` — `ALTER TABLE envelopes ADD COLUMN payload_sha256 TEXT;`
- `carbonstack-cypher/migrations/002_envelope_payload_metadata.sql:2` — `ALTER TABLE envelopes ADD COLUMN payload_size_bytes INTEGER;`
- `carbonstack-cypher/internal/httpapi/api_test.go:231` — `migrationsDir := filepath.Join("..", "..", "migrations")`
- `carbonstack-cypher/internal/httpapi/api_test.go:232` — `if err := store.Migrate(migrationsDir); err != nil {`
- `carbonstack-cypher/internal/config/config.go:16` — `MigrationsDir: getEnv("CYPHER_MIGRATIONS", "migrations"),`
- `carbonstack-cypher/internal/db/db.go:42` — `return fmt.Errorf("read migrations dir: %w", err)`
- `carbonstack-cypher/internal/db/db.go:60` — `return fmt.Errorf("read migration %s: %w", file, err)`
- `carbonstack-cypher/internal/db/db.go:64` — `return fmt.Errorf("apply migration %s: %w", file, err)`
- `carbonstack-cypher/cmd/cypher/main.go:22` — `log.Fatalf("run migrations: %v", err)`

Recon interpretation:

- SQLite remains appropriate for the current local-only v0.3.x backbone deployability line.
- SQLite stores relay/server state such as invites, accounts, devices, opaque envelopes, envelope metadata, and ack/delivery state.
- SQLite must not be described as plaintext message storage or as a client trust oracle.
- Persistent DB behavior remains experimental until migration behavior is hardened.
- If migrations are reapplied blindly, `ALTER TABLE ADD COLUMN` style migrations can become restart/upgrade hazards.
- Before stronger persistence claims, Cypher likely needs either `schema_migrations` tracking or explicit wipe-only experimental DB documentation.

Recommended next technical risk to inspect or fix:

> Can Cypher start repeatedly against the same persistent SQLite DB without migration errors?

## 10. Comms-to-Cypher Addressing Surface

Comms relay/Cypher URL related hits:

- `carbonstack-comms/README.md:13` — `The current validated artifact is a development proof that CarbonStackComms can use an OpenMLS sidecar and CarbonStackCypher relay storage to complete a local OpenMLS relay lifecycle.`
- `carbonstack-comms/README.md:16` — `_Related repositories: [carbonstack](https://git.bitcrusher32.win/bitcrusher32/carbonstack) / [carbonstack-cypher](https://git.bitcrusher32.win/bitcrusher32/carbonstack-cypher) / [carbonstack-os](https://git.bitcrusher32.win/bitcrusher32/carbonstack-os)_`
- `carbonstack-comms/README.md:25` — `- an internal relay helper for Cypher/OpenMLS artifact transport;`
- `carbonstack-comms/README.md:26` — `- a real-Cypher smoke harness;`
- `carbonstack-comms/README.md:30` — `## Current validated relay path`
- `carbonstack-comms/README.md:35` — `2. Comms submits it to Cypher as an opaque envelope.`
- `carbonstack-comms/README.md:38` — `5. Comms submits the Welcome to Cypher.`
- `carbonstack-comms/README.md:42` — `9. Comms submits it to Cypher.`
- `carbonstack-comms/README.md:68` — `It is used by protocol tests and relay smoke proofs.`
- `carbonstack-comms/README.md:113` — `powershell -ExecutionPolicy Bypass -File .\scripts\smoke-openmls-real-cypher-relay.ps1`
- `carbonstack-comms/docs/03-message-lifecycle.md:18` — `10. Encrypted envelope is submitted to CarbonStackCypher.`
- `carbonstack-comms/docs/03-message-lifecycle.md:23` — `1. Client receives encrypted envelope from CarbonStackCypher.`
- `carbonstack-comms/docs/04-phase1-client-lifecycle.md:11` — `The current validated proof is the OpenMLS sidecar + Cypher relay smoke path, not a polished runtime CLI messenger.`
- `carbonstack-comms/docs/04-phase1-client-lifecycle.md:26` — `6. Send it to CarbonStackCypher.`
- `carbonstack-comms/docs/04-phase1-client-lifecycle.md:33` — `It is not the current OpenMLS relay proof.`
- `carbonstack-comms/docs/04-phase1-client-lifecycle.md:40` — `2. Comms submits it to Cypher.`
- `carbonstack-comms/docs/04-phase1-client-lifecycle.md:43` — `5. Comms submits the Welcome to Cypher.`
- `carbonstack-comms/docs/04-phase1-client-lifecycle.md:47` — `9. Comms submits it to Cypher.`
- `carbonstack-comms/docs/04-phase1-client-lifecycle.md:63` — `The CLI/dev harness path remains useful, but it must be updated around the OpenMLS sidecar relay model rather than the old stub-only lifecycle.`
- `carbonstack-comms/docs/02-client-protocol-foundation.md:24` — `The current validated artifact is not a production protocol. It is a local development proof that OpenMLS artifacts can be generated, relayed through Cypher, consumed, and acknowledged after sidecar success.`
- `carbonstack-comms/docs/02-client-protocol-foundation.md:31` — `- Cypher opaque envelope relay;`
- `carbonstack-comms/docs/02-client-protocol-foundation.md:49` — `## Current OpenMLS relay content types`
- `carbonstack-comms/docs/02-client-protocol-foundation.md:51` — `The current Cypher relay path uses:`
- `carbonstack-comms/docs/requirements.md:16` — `- Cypher opaque envelope relay;`
- `carbonstack-comms/docs/requirements.md:19` — `- real local Cypher smoke harness.`
- `carbonstack-comms/docs/05-local-state-model-v0.md:59` — `The current validated relay proof focuses on:`
- `carbonstack-comms/docs/05-local-state-model-v0.md:62` — `- Cypher envelope relay;`
- `carbonstack-comms/scripts/README.md:17` — `The wrapper delegates to the lower-level real-Cypher smoke harness:`
- `carbonstack-comms/scripts/README.md:19` — `powershell -ExecutionPolicy Bypass -File .\scripts\smoke-openmls-real-cypher-relay.ps1`
- `carbonstack-comms/scripts/README.md:21` — `It proves the local OpenMLS sidecar + real Cypher server relay lifecycle:`
- `carbonstack-comms/scripts/README.md:23` — `1. start a real local Cypher server;`
- `carbonstack-comms/scripts/README.md:26` — `4. relay the KeyPackage through Cypher;`
- `carbonstack-comms/scripts/README.md:28` — `6. relay the Welcome through Cypher;`
- `carbonstack-comms/scripts/README.md:31` — `9. relay the application-message through Cypher;`
- `carbonstack-comms/scripts/README.md:54` — `The `-Full` path delegates to the lower-level smoke harness and runs the targeted real-server proof, relay tests, protocol tests, broader Go tests, and generated Rust/OpenMLS artifact guard.`
- `carbonstack-comms/scripts/README.md:58` — `powershell -ExecutionPolicy Bypass -File .\scripts\smoke-openmls-real-cypher-relay.ps1 -Full`
- `carbonstack-comms/scripts/README.md:60` — `Use `-Full` before pushing changes that affect relay, protocol, sidecar, client, or script behavior.`
- `carbonstack-comms/scripts/README.md:72` — `## Stale Cypher process warning`
- `carbonstack-comms/scripts/README.md:74` — `The smoke harness refuses to run if a `cypher` process is already active.`
- `carbonstack-comms/scripts/README.md:78` — `Get-Process cypher -ErrorAction SilentlyContinue \| Select-Object Id, ProcessName, Path`
- `carbonstack-comms/scripts/README.md:80` — `Stop stale test processes when no intentional Cypher server is running:`
- `carbonstack-comms/scripts/README.md:82` — `Get-Process cypher -ErrorAction SilentlyContinue \| Stop-Process -Force`
- `carbonstack-comms/scripts/README.md:84` — `This matters on Windows because stale `cypher.exe` processes can hold temp SQLite files open.`
- `carbonstack-comms/scripts/README.md:90` — `They are not the current OpenMLS + Cypher backbone proof.`
- `carbonstack-comms/scripts/README.md:103` — `It calls the existing real-Cypher smoke harness instead of duplicating the proof logic.`
- `carbonstack-comms/scripts/smoke-openmls-real-cypher-relay.ps1:25` — `Write-Host "CarbonStackComms OpenMLS real-Cypher relay smoke harness"`
- `carbonstack-comms/scripts/smoke-openmls-real-cypher-relay.ps1:34` — `Write-Host "  TestOpenMLSSidecarFullLifecycleRelayThroughRealCypherServer"`
- `carbonstack-comms/scripts/smoke-openmls-real-cypher-relay.ps1:37` — `Write-Host "  build a temp carbonstack-cypher test binary"`
- `carbonstack-comms/scripts/smoke-openmls-real-cypher-relay.ps1:38` — `Write-Host "  start a real Cypher server on localhost"`
- `carbonstack-comms/scripts/smoke-openmls-real-cypher-relay.ps1:40` — `Write-Host "  run OpenMLS KeyPackage -> Welcome -> application-message relay"`
- `carbonstack-comms/scripts/smoke-openmls-real-cypher-relay.ps1:46` — `$StaleCypher = Get-Process cypher -ErrorAction SilentlyContinue`
- `carbonstack-comms/scripts/smoke-openmls-real-cypher-relay.ps1:47` — `if ($StaleCypher) {`
- `carbonstack-comms/scripts/smoke-openmls-real-cypher-relay.ps1:48` — `Write-Host "WARNING: existing cypher processes detected:"`
- `carbonstack-comms/scripts/smoke-openmls-real-cypher-relay.ps1:49` — `$StaleCypher \| Select-Object Id, ProcessName, Path \| Format-Table \| Out-String \| Write-Host`
- `carbonstack-comms/scripts/smoke-openmls-real-cypher-relay.ps1:50` — `Write-Host "Refusing to continue while cypher processes are already running."`
- `carbonstack-comms/scripts/smoke-openmls-real-cypher-relay.ps1:52` — `Write-Host "  Get-Process cypher -ErrorAction SilentlyContinue \| Stop-Process -Force"`
- `carbonstack-comms/scripts/smoke-openmls-real-cypher-relay.ps1:56` — `Write-Host "Running targeted real-Cypher relay lifecycle smoke test..."`
- `carbonstack-comms/scripts/smoke-openmls-real-cypher-relay.ps1:57` — `Invoke-NativeCommand go test -p 1 ./internal/protocol -run "TestOpenMLSSidecarFullLifecycleRelayThroughRealCypherServer" -count=1 -timeout 240s`
- `carbonstack-comms/scripts/smoke-openmls-real-cypher-relay.ps1:65` — `Write-Host "Running broader protocol/relay validation because -Full was provided..."`
- `carbonstack-comms/scripts/smoke-openmls-real-cypher-relay.ps1:67` — `Invoke-NativeCommand go test -p 1 ./internal/protocol -run "TestOpenMLSSidecarFullLifecycleRelayThroughCypherEnvelope\|TestOpenMLSSidecarFullLifecycleRelayThroughRealCypherServer" -count=1 -timeout 300s`
- `carbonstack-comms/scripts/smoke-openmls-real-cypher-relay.ps1:68` — `Invoke-NativeCommand go test -p 1 ./internal/relay`
- `carbonstack-comms/scripts/test-local-lifecycle.ps1:2` — `[string]$Server = "http://localhost:8080"`
- `carbonstack-comms/scripts/test-trust-lifecycle.ps1:2` — `[string]$Server = "http://localhost:8080"`
- `carbonstack-comms/scripts/self-test-openmls-backbone.ps1:24` — `Write-Host "  + CarbonStackCypher real local server"`
- `carbonstack-comms/scripts/self-test-openmls-backbone.ps1:25` — `Write-Host "  + opaque OpenMLS artifact envelope relay"`
- `carbonstack-comms/scripts/self-test-openmls-backbone.ps1:30` — `Write-Host "  scripts/smoke-openmls-real-cypher-relay.ps1"`
- `carbonstack-comms/scripts/self-test-openmls-backbone.ps1:35` — `"-File", ".\scripts\smoke-openmls-real-cypher-relay.ps1"`
- `carbonstack-comms/internal/app/commands.go:79` — `serverURL := fs.String("server", state.DefaultServerURL, "CarbonStackCypher server URL")`
- `carbonstack-comms/internal/app/commands.go:85` — `ServerURL:       strings.TrimRight(*serverURL, "/"),`
- `carbonstack-comms/internal/app/commands.go:101` — `serverURL := fs.String("server", "", "CarbonStackCypher server URL override")`
- `carbonstack-comms/internal/app/commands.go:107` — `server := state.ServerFromStateOrFlag(*statePath, *serverURL)`
- `carbonstack-comms/internal/app/commands.go:125` — `serverURL := fs.String("server", "", "CarbonStackCypher server URL override")`
- `carbonstack-comms/internal/app/commands.go:137` — `server := state.ServerFromStateOrFlag(*statePath, *serverURL)`
- `carbonstack-comms/internal/state/state.go:12` — `DefaultServerURL = "http://localhost:8080"`
- `carbonstack-comms/internal/state/state_test.go:12` — `ServerURL:   "http://localhost:8080/",`
- `carbonstack-comms/internal/state/state_test.go:26` — `if got.ServerURL != "http://localhost:8080" {`
- `carbonstack-comms/internal/state/state_test.go:43` — `ServerURL: "http://localhost:8080",`
- `carbonstack-comms/internal/state/state_test.go:61` — `ServerURL: "http://localhost:8080",`
- `carbonstack-comms/internal/state/state_test.go:78` — `ServerURL: "http://localhost:9090/",`
- `carbonstack-comms/internal/state/state_test.go:86` — `if got != "http://localhost:9090" {`
- `carbonstack-comms/internal/client/cypher_test.go:11` — `func TestCypherClientLifecycleMethods(t *testing.T) {`
- `carbonstack-comms/internal/client/cypher_test.go:203` — `func TestCypherClientErrorResponse(t *testing.T) {`
- `carbonstack-comms/internal/client/cypher.go:81` — `type CypherClient struct {`
- `carbonstack-comms/internal/client/cypher.go:85` — `func New(serverURL string) CypherClient {`
- `carbonstack-comms/internal/client/cypher.go:86` — `return CypherClient{ServerURL: serverURL}`
- `carbonstack-comms/internal/client/cypher.go:89` — `func (c CypherClient) CreateDevInvite(inviteCode string) (DevInviteResponse, error) {`
- `carbonstack-comms/internal/client/cypher.go:96` — `func (c CypherClient) ClaimInvite(inviteCode string, displayName string) (ClaimInviteResponse, error) {`
- `carbonstack-comms/internal/client/cypher.go:106` — `func (c CypherClient) RegisterDevice(accountID string, label string, publicIdentityKey string, publicPrekeyBundle string) (RegisterDeviceResponse, error) {`
- `carbonstack-comms/internal/client/cypher.go:118` — `func (c CypherClient) ListDevices(accountID string) (ListDevicesResponse, error) {`
- `carbonstack-comms/internal/client/cypher.go:124` — `func (c CypherClient) SubmitEnvelope(senderDeviceID string, recipientDeviceID string, contentType string, protocolVersion string, ciphertextB64 string, clientCreatedAt string) (SubmitEnvelopeResponse, error) {`
- `carbonstack-comms/internal/client/cypher.go:138` — `func (c CypherClient) Inbox(deviceID string) (InboxResponse, error) {`
- `carbonstack-comms/internal/client/cypher.go:144` — `func (c CypherClient) AckEnvelope(envelopeID string, recipientDeviceID string) (AckEnvelopeResponse, error) {`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:11` — `"git.bitcrusher32.win/bitcrusher32/carbonstack-comms/internal/relay"`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:14` — `func TestOpenMLSSidecarFullLifecycleRelayThroughRealCypherServer(t *testing.T) {`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:17` — `server := startRealCypherTestServer(t)`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:18` — `cypherClient := client.New(server.URL())`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:20` — `aliceAccount, err := cypherClient.ClaimInvite("dev-invite", "alice-real-cypher-smoke")`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:22` — `t.Fatalf("claim Alice invite against real Cypher server: %v", err)`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:25` — `bobInvite, err := cypherClient.CreateDevInvite("bob-real-cypher-smoke-invite")`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:27` — `t.Fatalf("create Bob invite against real Cypher server: %v", err)`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:30` — `bobAccount, err := cypherClient.ClaimInvite(bobInvite.InviteCode, "bob-real-cypher-smoke")`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:32` — `t.Fatalf("claim Bob invite against real Cypher server: %v", err)`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:35` — `aliceCypherDevice, err := cypherClient.RegisterDevice(`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:37` — `"alice-real-cypher-device",`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:38` — `"stub-alice-real-cypher-public-identity-key",`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:39` — `"stub-alice-real-cypher-prekey-bundle",`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:42` — `t.Fatalf("register Alice Cypher device: %v", err)`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:45` — `bobCypherDevice, err := cypherClient.RegisterDevice(`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:47` — `"bob-real-cypher-device",`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:48` — `"stub-bob-real-cypher-public-identity-key",`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:49` — `"stub-bob-real-cypher-prekey-bundle",`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:52` — `t.Fatalf("register Bob Cypher device: %v", err)`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:83` — `if _, err := relay.SubmitOpenMLSArtifactEnvelope(`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:84` — `cypherClient,`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:85` — `bobCypherDevice.DeviceID,`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:86` — `aliceCypherDevice.DeviceID,`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:87` — `relay.ArtifactKindKeyPackage,`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:91` — `t.Fatalf("SubmitOpenMLSArtifactEnvelope keypackage through real Cypher failed: %v", err)`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:94` — `aliceInbox, err := cypherClient.Inbox(aliceCypherDevice.DeviceID)`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:96` — `t.Fatalf("alice real Cypher inbox failed: %v", err)`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:99` — `t.Fatalf("expected 1 keypackage relay envelope for Alice, got %d", len(aliceInbox.Envelopes))`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:103` — `if keyPackageEnvelope.ContentType != relay.ContentTypeOpenMLSKeyPackage {`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:104` — `t.Fatalf("keypackage content type = %q, want %q", keyPackageEnvelope.ContentType, relay.ContentTypeOpenMLSKeyPackage)`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:108` — `if err := relay.WriteOpenMLSArtifactFromEnvelope(downloadedKeyPackagePath, keyPackageEnvelope); err != nil {`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:134` — `t.Fatalf("conversation-add-member with relayed keypackage failed: %v\n%s", addMemberErr, string(addMemberOutput))`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:152` — `if _, err := cypherClient.AckEnvelope(keyPackageEnvelope.EnvelopeID, aliceCypherDevice.DeviceID); err != nil {`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:156` — `aliceInboxAfterAck, err := cypherClient.Inbox(aliceCypherDevice.DeviceID)`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:158` — `t.Fatalf("alice real Cypher inbox after keypackage ack failed: %v", err)`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:166` — `if _, err := relay.SubmitOpenMLSArtifactEnvelope(`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:167` — `cypherClient,`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:168` — `aliceCypherDevice.DeviceID,`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:169` — `bobCypherDevice.DeviceID,`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:170` — `relay.ArtifactKindWelcome,`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:174` — `t.Fatalf("SubmitOpenMLSArtifactEnvelope welcome through real Cypher failed: %v", err)`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:177` — `bobWelcomeInbox, err := cypherClient.Inbox(bobCypherDevice.DeviceID)`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:179` — `t.Fatalf("bob real Cypher inbox for welcome failed: %v", err)`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:182` — `t.Fatalf("expected 1 welcome relay envelope for Bob, got %d", len(bobWelcomeInbox.Envelopes))`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:186` — `if welcomeEnvelope.ContentType != relay.ContentTypeOpenMLSWelcome {`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:187` — `t.Fatalf("welcome content type = %q, want %q", welcomeEnvelope.ContentType, relay.ContentTypeOpenMLSWelcome)`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:191` — `if err := relay.WriteOpenMLSArtifactFromEnvelope(downloadedWelcomePath, welcomeEnvelope); err != nil {`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:202` — `t.Fatalf("conversation-join with relayed welcome failed: %v\n%s", joinErr, string(joinOutput))`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:217` — `if _, err := cypherClient.AckEnvelope(welcomeEnvelope.EnvelopeID, bobCypherDevice.DeviceID); err != nil {`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:221` — `bobInboxAfterWelcomeAck, err := cypherClient.Inbox(bobCypherDevice.DeviceID)`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:223` — `t.Fatalf("bob real Cypher inbox after welcome ack failed: %v", err)`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:229` — `messageLabel := "real-cypher-full-lifecycle-message-0001"`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:230` — `plaintext := "hello bob through real cypher server lifecycle"`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:242` — `if _, err := relay.SubmitOpenMLSArtifactEnvelope(`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:243` — `cypherClient,`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:244` — `aliceCypherDevice.DeviceID,`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:245` — `bobCypherDevice.DeviceID,`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:246` — `relay.ArtifactKindApplicationMessage,`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:250` — `t.Fatalf("SubmitOpenMLSArtifactEnvelope application message through real Cypher failed: %v", err)`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:253` — `bobMessageInbox, err := cypherClient.Inbox(bobCypherDevice.DeviceID)`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:255` — `t.Fatalf("bob real Cypher inbox for application message failed: %v", err)`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:264` — `if envelope.ContentType == relay.ContentTypeOpenMLSApplicationMessage {`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:271` — `t.Fatal("expected application-message relay envelope in Bob inbox")`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:275` — `if err := relay.WriteOpenMLSArtifactFromEnvelope(downloadedMessagePath, messageEnvelope); err != nil {`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:283` — `if _, err := cypherClient.AckEnvelope(messageEnvelope.EnvelopeID, bobCypherDevice.DeviceID); err != nil {`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:287` — `bobInboxAfterMessageAck, err := cypherClient.Inbox(bobCypherDevice.DeviceID)`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:289` — `t.Fatalf("bob real Cypher inbox after application-message ack failed: %v", err)`
- `carbonstack-comms/internal/protocol/openmls_sidecar_relay_test.go:8` — `"git.bitcrusher32.win/bitcrusher32/carbonstack-comms/internal/relay"`
- `carbonstack-comms/internal/protocol/openmls_sidecar_relay_test.go:11` — `func TestOpenMLSSidecarApplicationMessageRelayThroughCypherEnvelope(t *testing.T) {`
- `carbonstack-comms/internal/protocol/openmls_sidecar_relay_test.go:14` — `tc := newProtocolTestCypherServer(t)`
- `carbonstack-comms/internal/protocol/openmls_sidecar_relay_test.go:15` — `cypherClient := client.New(tc.URL())`
- `carbonstack-comms/internal/protocol/openmls_sidecar_relay_test.go:19` — `messageLabel := "relay-message-0001"`
- `carbonstack-comms/internal/protocol/openmls_sidecar_relay_test.go:20` — `plaintext := "hello bob through cypher relay"`
- `carbonstack-comms/internal/protocol/openmls_sidecar_relay_test.go:30` — `submitResp, err := relay.SubmitOpenMLSArtifactEnvelope(`
- `carbonstack-comms/internal/protocol/openmls_sidecar_relay_test.go:31` — `cypherClient,`
- `carbonstack-comms/internal/protocol/openmls_sidecar_relay_test.go:32` — `"alice-cypher-device-id",`
- `carbonstack-comms/internal/protocol/openmls_sidecar_relay_test.go:33` — `"bob-cypher-device-id",`
- `carbonstack-comms/internal/protocol/openmls_sidecar_relay_test.go:34` — `relay.ArtifactKindApplicationMessage,`
- `carbonstack-comms/internal/protocol/openmls_sidecar_relay_test.go:46` — `inbox, err := cypherClient.Inbox("bob-cypher-device-id")`
- `carbonstack-comms/internal/protocol/openmls_sidecar_relay_test.go:48` — `t.Fatalf("Cypher inbox failed: %v", err)`
- `carbonstack-comms/internal/protocol/openmls_sidecar_relay_test.go:52` — `t.Fatalf("expected 1 relay envelope, got %d", len(inbox.Envelopes))`
- `carbonstack-comms/internal/protocol/openmls_sidecar_relay_test.go:56` — `if envelope.ContentType != relay.ContentTypeOpenMLSApplicationMessage {`
- `carbonstack-comms/internal/protocol/openmls_sidecar_relay_test.go:57` — `t.Fatalf("content type = %q, want %q", envelope.ContentType, relay.ContentTypeOpenMLSApplicationMessage)`
- `carbonstack-comms/internal/protocol/openmls_sidecar_relay_test.go:60` — `if envelope.ProtocolVersion != relay.ProtocolVersionOpenMLSSidecar {`
- `carbonstack-comms/internal/protocol/openmls_sidecar_relay_test.go:61` — `t.Fatalf("protocol version = %q, want %q", envelope.ProtocolVersion, relay.ProtocolVersionOpenMLSSidecar)`
- `carbonstack-comms/internal/protocol/openmls_sidecar_relay_test.go:65` — `if err := relay.WriteOpenMLSArtifactFromEnvelope(downloadedArtifactPath, envelope); err != nil {`
- `carbonstack-comms/internal/protocol/openmls_sidecar_relay_test.go:73` — `func TestOpenMLSSidecarKeyPackageWelcomeRelayThroughCypherEnvelope(t *testing.T) {`
- ... 159 more hits omitted from this generated recon doc.

Recon interpretation:

- v0.3.x should identify the Comms-to-Cypher seam without wiring runtime `comms send` / `inbox` to OpenMLS yet.
- A future simple environment seam remains plausible:
  - `CARBONSTACK_CYPHER_URL=http://127.0.0.1:8080`
- A future CLI flag remains plausible:
  - `--cypher-url http://127.0.0.1:8080`
- Runtime Comms UX wiring remains v0.4.x work.

## 11. OpenMLS / Sidecar Deployability Surface

OpenMLS/sidecar related hits:

- `carbonstack-comms/README.md:13` — `The current validated artifact is a development proof that CarbonStackComms can use an OpenMLS sidecar and CarbonStackCypher relay storage to complete a local OpenMLS relay lifecycle.`
- `carbonstack-comms/README.md:23` — `- a promoted OpenMLS development sidecar;`
- `carbonstack-comms/README.md:24` — `- protocol tests for the OpenMLS sidecar lifecycle;`
- `carbonstack-comms/README.md:25` — `- an internal relay helper for Cypher/OpenMLS artifact transport;`
- `carbonstack-comms/README.md:27` — `- metadata validation before writing downloaded sidecar artifacts;`
- `carbonstack-comms/README.md:34` — `1. Bob exports an OpenMLS KeyPackage artifact.`
- `carbonstack-comms/README.md:37` — `4. Alice consumes it through the OpenMLS sidecar and creates a Welcome.`
- `carbonstack-comms/README.md:38` — `5. Comms submits the Welcome to Cypher.`
- `carbonstack-comms/README.md:39` — `6. Bob retrieves and writes the Welcome.`
- `carbonstack-comms/README.md:40` — `7. Bob consumes it through the sidecar.`
- `carbonstack-comms/README.md:41` — `8. Alice creates an application-message artifact.`
- `carbonstack-comms/README.md:43` — `10. Bob retrieves, validates metadata, writes, and consumes the application-message.`
- `carbonstack-comms/README.md:45` — `12. Envelopes are acked only after sidecar consume succeeds.`
- `carbonstack-comms/README.md:62` — `## OpenMLS sidecar`
- `carbonstack-comms/README.md:64` — `The promoted OpenMLS sidecar lives at:`
- `carbonstack-comms/README.md:66` — `internal/protocol/mls/openmls-sidecar`
- `carbonstack-comms/README.md:74` — `Generated signer/provider state must not be committed.`
- `carbonstack-comms/README.md:78` — `Run the current OpenMLS backbone self-test:`
- `carbonstack-comms/README.md:80` — `powershell -ExecutionPolicy Bypass -File .\scripts\self-test-openmls-backbone.ps1`
- `carbonstack-comms/README.md:84` — `powershell -ExecutionPolicy Bypass -File .\scripts\self-test-openmls-backbone.ps1 -Full`
- `carbonstack-comms/README.md:97` — `- production runtime send/inbox OpenMLS UX;`
- `carbonstack-comms/README.md:109` — `powershell -ExecutionPolicy Bypass -File .\scripts\self-test-openmls-backbone.ps1`
- `carbonstack-comms/README.md:113` — `powershell -ExecutionPolicy Bypass -File .\scripts\smoke-openmls-real-cypher-relay.ps1`
- `carbonstack-comms/docs/04-phase1-client-lifecycle.md:11` — `The current validated proof is the OpenMLS sidecar + Cypher relay smoke path, not a polished runtime CLI messenger.`
- `carbonstack-comms/docs/04-phase1-client-lifecycle.md:33` — `It is not the current OpenMLS relay proof.`
- `carbonstack-comms/docs/04-phase1-client-lifecycle.md:39` — `1. Bob exports an OpenMLS KeyPackage artifact.`
- `carbonstack-comms/docs/04-phase1-client-lifecycle.md:42` — `4. Alice consumes it through the OpenMLS sidecar and creates a Welcome.`
- `carbonstack-comms/docs/04-phase1-client-lifecycle.md:43` — `5. Comms submits the Welcome to Cypher.`
- `carbonstack-comms/docs/04-phase1-client-lifecycle.md:45` — `7. Bob consumes it through the sidecar.`
- `carbonstack-comms/docs/04-phase1-client-lifecycle.md:46` — `8. Alice creates an application-message artifact.`
- `carbonstack-comms/docs/04-phase1-client-lifecycle.md:50` — `12. Envelopes are acked only after sidecar consume succeeds.`
- `carbonstack-comms/docs/04-phase1-client-lifecycle.md:63` — `The CLI/dev harness path remains useful, but it must be updated around the OpenMLS sidecar relay model rather than the old stub-only lifecycle.`
- `carbonstack-comms/docs/02-client-protocol-foundation.md:22` — `The current mainline experimental proof uses an OpenMLS sidecar.`
- `carbonstack-comms/docs/02-client-protocol-foundation.md:24` — `The current validated artifact is not a production protocol. It is a local development proof that OpenMLS artifacts can be generated, relayed through Cypher, consumed, and acknowledged after sidecar success.`
- `carbonstack-comms/docs/02-client-protocol-foundation.md:28` — `- KeyPackage artifact;`
- `carbonstack-comms/docs/02-client-protocol-foundation.md:29` — `- Welcome artifact;`
- `carbonstack-comms/docs/02-client-protocol-foundation.md:30` — `- application-message artifact;`
- `carbonstack-comms/docs/02-client-protocol-foundation.md:49` — `## Current OpenMLS relay content types`
- `carbonstack-comms/docs/02-client-protocol-foundation.md:55` — `carbonstack.mls.application-message.v0`
- `carbonstack-comms/docs/02-client-protocol-foundation.md:59` — `carbonstack-openmls-sidecar-v0`
- `carbonstack-comms/docs/local-vault.md:9` — `It is not implemented by the current OpenMLS sidecar development state.`
- `carbonstack-comms/docs/local-vault.md:17` — `- OpenMLS group state;`
- `carbonstack-comms/docs/local-vault.md:32` — `Current development state may include generated signer/provider files and sidecar artifacts.`
- `carbonstack-comms/docs/requirements.md:15` — `- OpenMLS sidecar artifact lifecycle;`
- `carbonstack-comms/docs/requirements.md:53` — `The current development sidecar state is not that vault.`
- `carbonstack-comms/docs/05-local-state-model-v0.md:13` — `The current OpenMLS sidecar proof uses dev-local sidecar state and generated provider/signer files. Those files are sensitive development artifacts and must not be committed.`
- `carbonstack-comms/docs/05-local-state-model-v0.md:21` — `- `.carbonstack-openmls-sidecar-state/``
- `carbonstack-comms/docs/05-local-state-model-v0.md:22` — `- `signer.json``
- `carbonstack-comms/docs/05-local-state-model-v0.md:23` — `- `provider-storage.json``
- `carbonstack-comms/docs/05-local-state-model-v0.md:24` — `- raw OpenMLS group/provider state`
- `carbonstack-comms/docs/05-local-state-model-v0.md:61` — `- OpenMLS sidecar artifacts;`
- `carbonstack-comms/scripts/README.md:13` — `powershell -ExecutionPolicy Bypass -File .\scripts\self-test-openmls-backbone.ps1`
- `carbonstack-comms/scripts/README.md:15` — `This is the current OpenMLS backbone self-test path.`
- `carbonstack-comms/scripts/README.md:19` — `powershell -ExecutionPolicy Bypass -File .\scripts\smoke-openmls-real-cypher-relay.ps1`
- `carbonstack-comms/scripts/README.md:21` — `It proves the local OpenMLS sidecar + real Cypher server relay lifecycle:`
- `carbonstack-comms/scripts/README.md:25` — `3. export an OpenMLS KeyPackage;`
- `carbonstack-comms/scripts/README.md:26` — `4. relay the KeyPackage through Cypher;`
- `carbonstack-comms/scripts/README.md:27` — `5. consume the KeyPackage and produce a Welcome;`
- `carbonstack-comms/scripts/README.md:28` — `6. relay the Welcome through Cypher;`
- `carbonstack-comms/scripts/README.md:29` — `7. consume the Welcome;`
- `carbonstack-comms/scripts/README.md:30` — `8. protect an application-message;`
- `carbonstack-comms/scripts/README.md:31` — `9. relay the application-message through Cypher;`
- `carbonstack-comms/scripts/README.md:33` — `11. consume the application-message through the sidecar;`
- `carbonstack-comms/scripts/README.md:34` — `12. ack envelopes only after sidecar consume succeeds.`
- `carbonstack-comms/scripts/README.md:52` — `powershell -ExecutionPolicy Bypass -File .\scripts\self-test-openmls-backbone.ps1 -Full`
- `carbonstack-comms/scripts/README.md:54` — `The `-Full` path delegates to the lower-level smoke harness and runs the targeted real-server proof, relay tests, protocol tests, broader Go tests, and generated Rust/OpenMLS artifact guard.`
- `carbonstack-comms/scripts/README.md:58` — `powershell -ExecutionPolicy Bypass -File .\scripts\smoke-openmls-real-cypher-relay.ps1 -Full`
- `carbonstack-comms/scripts/README.md:60` — `Use `-Full` before pushing changes that affect relay, protocol, sidecar, client, or script behavior.`
- `carbonstack-comms/scripts/README.md:70` — `It does not check every sensitive generated sidecar file. It is a build-artifact guard, not a complete secret scanner.`
- `carbonstack-comms/scripts/README.md:90` — `They are not the current OpenMLS + Cypher backbone proof.`
- `carbonstack-comms/scripts/README.md:92` — `Use the OpenMLS backbone self-test path for the current known-good backbone validation.`
- `carbonstack-comms/scripts/README.md:97` — `OpenMLS backbone self-test harness`
- `carbonstack-comms/scripts/README.md:101` — `scripts/self-test-openmls-backbone.ps1`
- `carbonstack-comms/scripts/smoke-openmls-real-cypher-relay.ps1:25` — `Write-Host "CarbonStackComms OpenMLS real-Cypher relay smoke harness"`
- `carbonstack-comms/scripts/smoke-openmls-real-cypher-relay.ps1:33` — `Write-Host "This harness runs the current known-good OpenMLS backbone proof:"`
- `carbonstack-comms/scripts/smoke-openmls-real-cypher-relay.ps1:34` — `Write-Host "  TestOpenMLSSidecarFullLifecycleRelayThroughRealCypherServer"`
- `carbonstack-comms/scripts/smoke-openmls-real-cypher-relay.ps1:40` — `Write-Host "  run OpenMLS KeyPackage -> Welcome -> application-message relay"`
- `carbonstack-comms/scripts/smoke-openmls-real-cypher-relay.ps1:42` — `Write-Host "  ack envelopes only after successful sidecar consume"`
- `carbonstack-comms/scripts/smoke-openmls-real-cypher-relay.ps1:43` — `Write-Host "  verify final sidecar message-open plaintext recovery"`
- `carbonstack-comms/scripts/smoke-openmls-real-cypher-relay.ps1:57` — `Invoke-NativeCommand go test -p 1 ./internal/protocol -run "TestOpenMLSSidecarFullLifecycleRelayThroughRealCypherServer" -count=1 -timeout 240s`
- `carbonstack-comms/scripts/smoke-openmls-real-cypher-relay.ps1:60` — `Write-Host "Running generated Rust/OpenMLS artifact guard..."`
- `carbonstack-comms/scripts/smoke-openmls-real-cypher-relay.ps1:67` — `Invoke-NativeCommand go test -p 1 ./internal/protocol -run "TestOpenMLSSidecarFullLifecycleRelayThroughCypherEnvelope\|TestOpenMLSSidecarFullLifecycleRelayThroughRealCypherServer" -count=1 -timeout 300s`
- `carbonstack-comms/scripts/smoke-openmls-real-cypher-relay.ps1:73` — `Write-Host "Running generated Rust/OpenMLS artifact guard again after full validation..."`
- `carbonstack-comms/scripts/self-test-openmls-backbone.ps1:10` — `Write-Host "CarbonStack OpenMLS backbone self-test harness"`
- `carbonstack-comms/scripts/self-test-openmls-backbone.ps1:23` — `Write-Host "  CarbonStackComms OpenMLS sidecar"`
- `carbonstack-comms/scripts/self-test-openmls-backbone.ps1:25` — `Write-Host "  + opaque OpenMLS artifact envelope relay"`
- `carbonstack-comms/scripts/self-test-openmls-backbone.ps1:30` — `Write-Host "  scripts/smoke-openmls-real-cypher-relay.ps1"`
- `carbonstack-comms/scripts/self-test-openmls-backbone.ps1:35` — `"-File", ".\scripts\smoke-openmls-real-cypher-relay.ps1"`
- `carbonstack-comms/scripts/self-test-openmls-backbone.ps1:53` — `throw "OpenMLS backbone self-test failed with exit code $LASTEXITCODE"`
- `carbonstack-comms/scripts/self-test-openmls-backbone.ps1:57` — `Write-Host "OpenMLS backbone self-test completed successfully."`
- `carbonstack-comms/scripts/check-no-rust-artifacts.ps1:14` — `".carbonstack-openmls-sidecar-state",`
- `carbonstack-comms/scripts/check-no-rust-artifacts.ps1:20` — `"provider-storage.json",`
- `carbonstack-comms/scripts/check-no-rust-artifacts.ps1:21` — `"signer.json"`
- `carbonstack-comms/scripts/check-no-rust-artifacts.ps1:151` — `Write-Host "Note: tests may generate target/, OpenMLS dev state, temp DBs, and executables after validation."`
- `carbonstack-comms/internal/protocol/openmls_sidecar_message_negative_test.go:10` — `func TestOpenMLSSidecarMessageOpenWrongDeviceRejected(t *testing.T) {`
- `carbonstack-comms/internal/protocol/openmls_sidecar_message_negative_test.go:11` — `removeOpenMLSSidecarState(t)`
- `carbonstack-comms/internal/protocol/openmls_sidecar_message_negative_test.go:13` — `addMemberEnvelope := setupOpenMLSTwoMemberConversation(t)`
- `carbonstack-comms/internal/protocol/openmls_sidecar_message_negative_test.go:15` — `eveIdentityOutput, eveIdentityErr := runOpenMLSSidecar("identity-create", "--device-label", "carbonstack-eve-device")`
- `carbonstack-comms/internal/protocol/openmls_sidecar_message_negative_test.go:20` — `message1ProtectEnvelope := protectOpenMLSSidecarMessage(t, "message-0001", "hello bob wrong device probe")`
- `carbonstack-comms/internal/protocol/openmls_sidecar_message_negative_test.go:22` — `wrongDeviceOutput, wrongDeviceErr := runOpenMLSSidecar(`
- `carbonstack-comms/internal/protocol/openmls_sidecar_message_negative_test.go:57` — `func TestOpenMLSSidecarMessageOpenWrongConversationRejected(t *testing.T) {`
- `carbonstack-comms/internal/protocol/openmls_sidecar_message_negative_test.go:58` — `removeOpenMLSSidecarState(t)`
- `carbonstack-comms/internal/protocol/openmls_sidecar_message_negative_test.go:60` — `setupOpenMLSTwoMemberConversation(t)`
- `carbonstack-comms/internal/protocol/openmls_sidecar_message_negative_test.go:62` — `message1ProtectEnvelope := protectOpenMLSSidecarMessage(t, "message-0001", "hello bob wrong conversation probe")`
- `carbonstack-comms/internal/protocol/openmls_sidecar_message_negative_test.go:64` — `wrongConversationOutput, wrongConversationErr := runOpenMLSSidecar(`
- `carbonstack-comms/internal/protocol/openmls_sidecar_message_negative_test.go:99` — `func TestOpenMLSSidecarMessageOpenDuplicateRejected(t *testing.T) {`
- `carbonstack-comms/internal/protocol/openmls_sidecar_message_negative_test.go:100` — `removeOpenMLSSidecarState(t)`
- `carbonstack-comms/internal/protocol/openmls_sidecar_message_negative_test.go:102` — `addMemberEnvelope := setupOpenMLSTwoMemberConversation(t)`
- `carbonstack-comms/internal/protocol/openmls_sidecar_message_negative_test.go:104` — `message1ProtectEnvelope := protectOpenMLSSidecarMessage(t, "message-0001", "hello bob 1")`
- `carbonstack-comms/internal/protocol/openmls_sidecar_message_negative_test.go:106` — `message1OpenEnvelope, message1OpenOutput := openOpenMLSSidecarMessage(t, "message-0001", message1ProtectEnvelope.Data.MessageArtifactPathHint)`
- `carbonstack-comms/internal/protocol/openmls_sidecar_message_negative_test.go:110` — `duplicateOpenOutput, duplicateOpenErr := runOpenMLSSidecar(`
- `carbonstack-comms/internal/protocol/openmls_sidecar_message_negative_test.go:137` — `func TestOpenMLSSidecarMessageOpenCorruptArtifactRejected(t *testing.T) {`
- `carbonstack-comms/internal/protocol/openmls_sidecar_message_negative_test.go:138` — `removeOpenMLSSidecarState(t)`
- `carbonstack-comms/internal/protocol/openmls_sidecar_message_negative_test.go:140` — `setupOpenMLSTwoMemberConversation(t)`
- `carbonstack-comms/internal/protocol/openmls_sidecar_message_negative_test.go:142` — `message1ProtectEnvelope := protectOpenMLSSidecarMessage(t, "message-0001", "hello bob 1")`
- `carbonstack-comms/internal/protocol/openmls_sidecar_message_negative_test.go:145` — `badPath := filepath.Join(filepath.Dir(goodPath), "corrupt-application-message.bin")`
- `carbonstack-comms/internal/protocol/openmls_sidecar_message_negative_test.go:166` — `corruptOpenOutput, corruptOpenErr := runOpenMLSSidecar(`
- `carbonstack-comms/internal/protocol/provider_trust.go:55` — `ProviderEventConversationWelcomeCreated,`
- `carbonstack-comms/internal/protocol/provider_trust.go:56` — `ProviderEventConversationWelcomeStaged,`
- `carbonstack-comms/internal/protocol/openmls_sidecar_provider_info_test.go:7` — `func TestOpenMLSSidecarProviderInfoCommand(t *testing.T) {`
- `carbonstack-comms/internal/protocol/openmls_sidecar_provider_info_test.go:8` — `output, err := runOpenMLSSidecar("provider-info")`
- `carbonstack-comms/internal/protocol/openmls_sidecar_provider_info_test.go:10` — `t.Fatalf("run OpenMLS sidecar provider-info: %v", err)`
- `carbonstack-comms/internal/protocol/openmls_sidecar_provider_info_test.go:69` — `func TestOpenMLSSidecarUnsupportedCommandEnvelope(t *testing.T) {`
- `carbonstack-comms/internal/protocol/openmls_sidecar_provider_info_test.go:70` — `output, err := runOpenMLSSidecar("state-checkpoint")`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:14` — `func TestOpenMLSSidecarFullLifecycleRelayThroughRealCypherServer(t *testing.T) {`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:15` — `removeOpenMLSSidecarState(t)`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:55` — `aliceIdentityOutput, aliceIdentityErr := runOpenMLSSidecar("identity-create", "--device-label", "carbonstack-alice-device")`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:61` — `bobIdentityOutput, bobIdentityErr := runOpenMLSSidecar("identity-create", "--device-label", "carbonstack-bob-device")`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:67` — `bobBundleOutput, bobBundleErr := runOpenMLSSidecar("public-bundle-export", "--device-label", "carbonstack-bob-device", "--write-artifact")`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:76` — `if bobBundleEnvelope.Data.KeyPackageArtifactPathHint == "" {`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:81` — `keyPackageArtifactPath := filepath.Join(openMLSSidecarDir, bobBundleEnvelope.Data.KeyPackageArtifactPathHint)`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:83` — `if _, err := relay.SubmitOpenMLSArtifactEnvelope(`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:87` — `relay.ArtifactKindKeyPackage,`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:91` — `t.Fatalf("SubmitOpenMLSArtifactEnvelope keypackage through real Cypher failed: %v", err)`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:103` — `if keyPackageEnvelope.ContentType != relay.ContentTypeOpenMLSKeyPackage {`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:104` — `t.Fatalf("keypackage content type = %q, want %q", keyPackageEnvelope.ContentType, relay.ContentTypeOpenMLSKeyPackage)`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:107` — `downloadedKeyPackagePath := filepath.Join(t.TempDir(), "downloaded-public-bundle.keypackage.bin")`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:108` — `if err := relay.WriteOpenMLSArtifactFromEnvelope(downloadedKeyPackagePath, keyPackageEnvelope); err != nil {`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:109` — `t.Fatalf("WriteOpenMLSArtifactFromEnvelope keypackage failed: %v", err)`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:112` — `aliceConversationOutput, aliceConversationErr := runOpenMLSSidecar(`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:127` — `addMemberOutput, addMemberErr := runOpenMLSSidecar(`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:131` — `"--member-keypackage", downloadedKeyPackagePath,`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:144` — `if addMemberEnvelope.Data.WelcomeArtifactPathHint == "" {`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:164` — `welcomeArtifactPath := filepath.Join(openMLSSidecarDir, addMemberEnvelope.Data.WelcomeArtifactPathHint)`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:166` — `if _, err := relay.SubmitOpenMLSArtifactEnvelope(`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:170` — `relay.ArtifactKindWelcome,`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:174` — `t.Fatalf("SubmitOpenMLSArtifactEnvelope welcome through real Cypher failed: %v", err)`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:177` — `bobWelcomeInbox, err := cypherClient.Inbox(bobCypherDevice.DeviceID)`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:181` — `if len(bobWelcomeInbox.Envelopes) != 1 {`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:182` — `t.Fatalf("expected 1 welcome relay envelope for Bob, got %d", len(bobWelcomeInbox.Envelopes))`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:185` — `welcomeEnvelope := bobWelcomeInbox.Envelopes[0]`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:186` — `if welcomeEnvelope.ContentType != relay.ContentTypeOpenMLSWelcome {`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:187` — `t.Fatalf("welcome content type = %q, want %q", welcomeEnvelope.ContentType, relay.ContentTypeOpenMLSWelcome)`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:190` — `downloadedWelcomePath := filepath.Join(t.TempDir(), "downloaded-welcome.bin")`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:191` — `if err := relay.WriteOpenMLSArtifactFromEnvelope(downloadedWelcomePath, welcomeEnvelope); err != nil {`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:192` — `t.Fatalf("WriteOpenMLSArtifactFromEnvelope welcome failed: %v", err)`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:195` — `joinOutput, joinErr := runOpenMLSSidecar(`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:199` — `"--welcome", downloadedWelcomePath,`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:221` — `bobInboxAfterWelcomeAck, err := cypherClient.Inbox(bobCypherDevice.DeviceID)`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:225` — `if len(bobInboxAfterWelcomeAck.Envelopes) != 0 {`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:226` — `t.Fatalf("expected Bob inbox to be empty after welcome ack, got %d envelopes", len(bobInboxAfterWelcomeAck.Envelopes))`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:232` — `protectEnvelope := protectOpenMLSSidecarMessage(t, messageLabel, plaintext)`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:242` — `if _, err := relay.SubmitOpenMLSArtifactEnvelope(`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:250` — `t.Fatalf("SubmitOpenMLSArtifactEnvelope application message through real Cypher failed: %v", err)`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:258` — `t.Fatalf("expected 1 queued application-message envelope for Bob after welcome ack, got %d", len(bobMessageInbox.Envelopes))`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:264` — `if envelope.ContentType == relay.ContentTypeOpenMLSApplicationMessage {`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:271` — `t.Fatal("expected application-message relay envelope in Bob inbox")`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:274` — `downloadedMessagePath := filepath.Join(t.TempDir(), "downloaded-application-message.bin")`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:275` — `if err := relay.WriteOpenMLSArtifactFromEnvelope(downloadedMessagePath, messageEnvelope); err != nil {`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:276` — `t.Fatalf("WriteOpenMLSArtifactFromEnvelope application message failed: %v", err)`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:279` — `openEnvelope, openOutput := openOpenMLSSidecarMessage(t, messageLabel, downloadedMessagePath)`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:284` — `t.Fatalf("ack application-message envelope after successful message-open failed: %v", err)`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:289` — `t.Fatalf("bob real Cypher inbox after application-message ack failed: %v", err)`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:292` — `t.Fatalf("expected Bob inbox to be empty after application-message ack, got %d envelopes", len(bobInboxAfterMessageAck.Envelopes))`
- `carbonstack-comms/internal/protocol/openmls_fixture_test.go:11` — `const openMLSFixtureDir = "mls/research/openmls-minimal/fixtures/dev"`
- `carbonstack-comms/internal/protocol/openmls_fixture_test.go:27` — `PublicKeyPackageHashRefLength int    `json:"public_key_package_hash_ref_length"``
- `carbonstack-comms/internal/protocol/openmls_fixture_test.go:34` — `OpenMLSError                string   `json:"openmls_error"``
- `carbonstack-comms/internal/protocol/openmls_fixture_test.go:44` — `func TestOpenMLSProviderSummaryFixture(t *testing.T) {`
- `carbonstack-comms/internal/protocol/openmls_fixture_test.go:52` — `if summary.ProviderName != "openmls" {`
- `carbonstack-comms/internal/protocol/openmls_fixture_test.go:73` — `func TestOpenMLSDeviceSummaryFixtures(t *testing.T) {`
- `carbonstack-comms/internal/protocol/openmls_fixture_test.go:84` — `func TestOpenMLSInvalidSignatureFixture(t *testing.T) {`
- `carbonstack-comms/internal/protocol/openmls_fixture_test.go:88` — `if fixture.OpenMLSError != "ValidationError(InvalidSignature)" {`
- `carbonstack-comms/internal/protocol/openmls_fixture_test.go:89` — `t.Fatalf("unexpected OpenMLS error: %q", fixture.OpenMLSError)`
- `carbonstack-comms/internal/protocol/openmls_fixture_test.go:105` — `func TestOpenMLSProviderEventFixtureStream(t *testing.T) {`
- `carbonstack-comms/internal/protocol/openmls_fixture_test.go:152` — `if fixture.PublicKeyPackageHashRefLength <= 0 {`
- `carbonstack-comms/internal/protocol/openmls_fixture_test.go:153` — `t.Fatalf("expected positive KeyPackage hash ref length for %s fixture", role)`
- `carbonstack-comms/internal/protocol/provider_trust_test.go:9` — `ProviderEventConversationWelcomeCreated,`
- `carbonstack-comms/internal/protocol/provider_trust_test.go:10` — `ProviderEventConversationWelcomeStaged,`
- `carbonstack-comms/internal/protocol/provider_trust_test.go:88` — `t.Fatal("missing signer should block send")`
- `carbonstack-comms/internal/protocol/provider_trust_test.go:92` — `t.Fatal("missing signer should be user visible")`
- `carbonstack-comms/internal/protocol/provider_trust_test.go:96` — `t.Fatal("missing signer should be history relevant")`
- `carbonstack-comms/internal/protocol/openmls_sidecar_public_bundle_test.go:10` — `func TestOpenMLSSidecarPublicBundleExportMissingIdentity(t *testing.T) {`
- `carbonstack-comms/internal/protocol/openmls_sidecar_public_bundle_test.go:11` — `removeOpenMLSSidecarState(t)`
- `carbonstack-comms/internal/protocol/openmls_sidecar_public_bundle_test.go:13` — `output, err := runOpenMLSSidecar("public-bundle-export", "--device-label", "carbonstack-alice-device")`
- `carbonstack-comms/internal/protocol/openmls_sidecar_public_bundle_test.go:36` — `func TestOpenMLSSidecarPublicBundleExportCreatesSummary(t *testing.T) {`
- `carbonstack-comms/internal/protocol/openmls_sidecar_public_bundle_test.go:37` — `removeOpenMLSSidecarState(t)`
- `carbonstack-comms/internal/protocol/openmls_sidecar_public_bundle_test.go:39` — `createOutput, createErr := runOpenMLSSidecar("identity-create", "--device-label", "carbonstack-alice-device")`
- `carbonstack-comms/internal/protocol/openmls_sidecar_public_bundle_test.go:46` — `exportOutput, exportErr := runOpenMLSSidecar("public-bundle-export", "--device-label", "carbonstack-alice-device")`
- `carbonstack-comms/internal/protocol/openmls_sidecar_public_bundle_test.go:83` — `if !exportEnvelope.Data.KeyPackageCreated {`
- `carbonstack-comms/internal/protocol/openmls_sidecar_public_bundle_test.go:87` — `if exportEnvelope.Data.KeyPackageArtifactWritten {`
- `carbonstack-comms/internal/protocol/openmls_sidecar_public_bundle_test.go:88` — `t.Fatal("public-bundle-export must not claim full KeyPackage artifact was written in this rung")`
- `carbonstack-comms/internal/protocol/openmls_sidecar_public_bundle_test.go:92` — `t.Fatal("public-bundle-export should report provider_storage_written=true because KeyPackage private provider state is needed for later Welcome consumption")`
- `carbonstack-comms/internal/protocol/openmls_sidecar_public_bundle_test.go:103` — `if exportEnvelope.Data.KeyPackageRef == "" {`
- `carbonstack-comms/internal/protocol/openmls_sidecar_public_bundle_test.go:107` — `if !strings.HasPrefix(exportEnvelope.Data.KeyPackageRef, "sha256:") {`
- `carbonstack-comms/internal/protocol/openmls_sidecar_public_bundle_test.go:108` — `t.Fatalf("key package ref = %q, want sha256 prefix", exportEnvelope.Data.KeyPackageRef)`
- `carbonstack-comms/internal/protocol/openmls_sidecar_public_bundle_test.go:111` — `if exportEnvelope.Data.KeyPackageHashLen != 32 {`
- `carbonstack-comms/internal/protocol/openmls_sidecar_public_bundle_test.go:112` — `t.Fatalf("key package hash len = %d, want 32", exportEnvelope.Data.KeyPackageHashLen)`
- `carbonstack-comms/internal/protocol/openmls_sidecar_public_bundle_test.go:129` — `stateDir := filepath.Join(openMLSSidecarDir, ".carbonstack-openmls-sidecar-state", "dev", "devices", "carbonstack-alice-device")`
- `carbonstack-comms/internal/protocol/openmls_sidecar_public_bundle_test.go:131` — `signerPath := filepath.Join(stateDir, "signer.json")`
- `carbonstack-comms/internal/protocol/openmls_sidecar_public_bundle_test.go:134` — `assertFileExists(t, signerPath)`
- `carbonstack-comms/internal/protocol/openmls_sidecar_public_bundle_test.go:141` — `KeyPackageCreated            bool   `json:"key_package_created"``
- `carbonstack-comms/internal/protocol/openmls_sidecar_public_bundle_test.go:142` — `KeyPackageRef                string `json:"key_package_ref"``
- `carbonstack-comms/internal/protocol/openmls_sidecar_public_bundle_test.go:143` — `KeyPackageHashLen            int    `json:"key_package_hash_len"``
- `carbonstack-comms/internal/protocol/openmls_sidecar_public_bundle_test.go:144` — `KeyPackageArtifactWritten    bool   `json:"key_package_artifact_written"``
- `carbonstack-comms/internal/protocol/openmls_sidecar_public_bundle_test.go:145` — `KeyPackageArtifactPathHint   string `json:"key_package_artifact_path_hint"``
- `carbonstack-comms/internal/protocol/openmls_sidecar_public_bundle_test.go:146` — `KeyPackageArtifactSHA256     string `json:"key_package_artifact_sha256"``
- `carbonstack-comms/internal/protocol/openmls_sidecar_public_bundle_test.go:147` — `KeyPackageArtifactSizeBytes  int    `json:"key_package_artifact_size_bytes"``
- `carbonstack-comms/internal/protocol/openmls_sidecar_public_bundle_test.go:156` — `WelcomeArtifactWritten       bool   `json:"welcome_artifact_written"``
- `carbonstack-comms/internal/protocol/openmls_sidecar_public_bundle_test.go:163` — `WelcomeArtifactPathHint      string `json:"welcome_artifact_path_hint"``
- `carbonstack-comms/internal/protocol/openmls_sidecar_public_bundle_test.go:164` — `WelcomeManifestPathHint      string `json:"welcome_manifest_path_hint"``
- ... 536 more hits omitted from this generated recon doc.

Recon interpretation:

- The OpenMLS sidecar remains dev/pre-alpha infrastructure.
- Generated sidecar state, `provider-storage.json`, `signer.json`, KeyPackage artifacts, Welcome artifacts, and application-message artifacts must stay scoped to known generated roots.
- Secret-bearing generated state must not be committed, pasted, or normalized as safe.
- The current sidecar lifecycle proof is valuable as backbone validation, not production secure storage.

## 12. Runtime Comms UX Surface

Runtime command related hits:

- `carbonstack-comms/README.md:97` — `- production runtime send/inbox OpenMLS UX;`
- `carbonstack-comms/docs/03-message-lifecycle.md:19` — `11. Local send record is stored in the secure vault.`
- `carbonstack-comms/docs/03-message-lifecycle.md:40` — `- sender device is revoked`
- `carbonstack-comms/docs/04-phase1-client-lifecycle.md:71` — `It does not describe a finished runtime send/inbox UX.`
- `carbonstack-comms/docs/02-client-protocol-foundation.md:43` — `- forge sender identity;`
- `carbonstack-comms/docs/requirements.md:43` — `- forge sender identity;`
- `carbonstack-comms/docs/requirements.md:61` — `- polished runtime send/inbox UX;`
- `carbonstack-comms/docs/01-comms-architecture.md:66` — `- sender device identifier`
- `carbonstack-comms/docs/01-comms-architecture.md:111` — `The sending path and receiving path should both enforce text policy.`
- `carbonstack-comms/scripts/README.md:11` — `Run:`
- `carbonstack-comms/scripts/README.md:50` — `Run:`
- `carbonstack-comms/scripts/README.md:64` — `Run:`
- `carbonstack-comms/scripts/test-local-lifecycle.ps1:20` — `[string[]]$Args`
- `carbonstack-comms/scripts/test-local-lifecycle.ps1:24` — `Write-Host "> go run .\cmd\comms $($Args -join ' ')"`
- `carbonstack-comms/scripts/test-local-lifecycle.ps1:25` — `go run .\cmd\comms @Args`
- `carbonstack-comms/scripts/test-local-lifecycle.ps1:28` — `throw "command failed: go run .\cmd\comms $($Args -join ' ')"`
- `carbonstack-comms/scripts/test-local-lifecycle.ps1:52` — `$SendOutput = go run .\cmd\comms send --state $AliceState --to-device $BobDeviceId --message $Message`
- `carbonstack-comms/scripts/test-local-lifecycle.ps1:56` — `throw "send command failed"`
- `carbonstack-comms/scripts/test-local-lifecycle.ps1:59` — `$InboxOutput = go run .\cmd\comms inbox --state $BobState`
- `carbonstack-comms/scripts/test-local-lifecycle.ps1:63` — `throw "inbox command failed"`
- `carbonstack-comms/scripts/test-local-lifecycle.ps1:68` — `throw "no envelope_id found in Bob inbox output"`
- `carbonstack-comms/scripts/test-local-lifecycle.ps1:78` — `$InboxAfterAck = go run .\cmd\comms inbox --state $BobState`
- `carbonstack-comms/scripts/test-local-lifecycle.ps1:82` — `throw "post-ack inbox command failed"`
- `carbonstack-comms/scripts/test-trust-lifecycle.ps1:15` — `$VerifiedMessage = "phase2a verified send $RunId"`
- `carbonstack-comms/scripts/test-trust-lifecycle.ps1:16` — `$ReverifiedMessage = "phase2a reverified send $RunId"`
- `carbonstack-comms/scripts/test-trust-lifecycle.ps1:22` — `[string[]]$Args`
- `carbonstack-comms/scripts/test-trust-lifecycle.ps1:26` — `Write-Host "> go run .\cmd\comms $($Args -join ' ')"`
- `carbonstack-comms/scripts/test-trust-lifecycle.ps1:27` — `go run .\cmd\comms @Args`
- `carbonstack-comms/scripts/test-trust-lifecycle.ps1:30` — `throw "command failed: go run .\cmd\comms $($Args -join ' ')"`
- `carbonstack-comms/scripts/test-trust-lifecycle.ps1:37` — `[string[]]$Args`
- `carbonstack-comms/scripts/test-trust-lifecycle.ps1:41` — `Write-Host "> go run .\cmd\comms $($Args -join ' ')"`
- `carbonstack-comms/scripts/test-trust-lifecycle.ps1:42` — `$Output = go run .\cmd\comms @Args`
- `carbonstack-comms/scripts/test-trust-lifecycle.ps1:48` — `throw "command failed: go run .\cmd\comms $($Args -join ' ')"`
- `carbonstack-comms/scripts/test-trust-lifecycle.ps1:57` — `[string[]]$Args`
- `carbonstack-comms/scripts/test-trust-lifecycle.ps1:61` — `Write-Host "> go run .\cmd\comms $($Args -join ' ')"`
- `carbonstack-comms/scripts/test-trust-lifecycle.ps1:67` — `$Output = & go run .\cmd\comms @Args 2>&1`
- `carbonstack-comms/scripts/test-trust-lifecycle.ps1:119` — `Run-Comms send --state $AliceState --to-device $BobDeviceId --message $VerifiedMessage --strict`
- `carbonstack-comms/scripts/test-trust-lifecycle.ps1:129` — `$StrictChanged = Run-Comms-AllowFailure send --state $AliceState --to-device $BobDeviceId --message "this should block" --strict`
- `carbonstack-comms/scripts/test-trust-lifecycle.ps1:132` — `throw "strict send unexpectedly succeeded after simulated key change"`
- `carbonstack-comms/scripts/test-trust-lifecycle.ps1:139` — `Run-Comms send --state $AliceState --to-device $BobDeviceId --message $ChangedDevMessage`
- `carbonstack-comms/scripts/test-trust-lifecycle.ps1:148` — `Run-Comms send --state $AliceState --to-device $BobDeviceId --message $ReverifiedMessage --strict`
- `carbonstack-comms/scripts/test-trust-lifecycle.ps1:157` — `$RevokedDev = Run-Comms-AllowFailure send --state $AliceState --to-device $BobDeviceId --message "this should block even in dev"`
- `carbonstack-comms/scripts/test-trust-lifecycle.ps1:160` — `throw "dev send unexpectedly succeeded after revocation"`
- `carbonstack-comms/scripts/self-test-openmls-backbone.ps1:33` — `$SmokeArgs = @(`
- `carbonstack-comms/scripts/self-test-openmls-backbone.ps1:42` — `$SmokeArgs += "-Full"`
- `carbonstack-comms/scripts/self-test-openmls-backbone.ps1:50` — `& powershell @SmokeArgs`
- `carbonstack-comms/internal/app/commands.go:5` — `"flag"`
- `carbonstack-comms/internal/app/commands.go:45` — `case "send":`
- `carbonstack-comms/internal/app/commands.go:47` — `case "inbox":`
- `carbonstack-comms/internal/app/commands.go:71` — `fmt.Println("  send")`
- `carbonstack-comms/internal/app/commands.go:72` — `fmt.Println("  inbox")`
- `carbonstack-comms/internal/app/commands.go:77` — `fs := flag.NewFlagSet("init", flag.ExitOnError)`
- `carbonstack-comms/internal/app/commands.go:99` — `fs := flag.NewFlagSet("dev-create-invite", flag.ExitOnError)`
- `carbonstack-comms/internal/app/commands.go:123` — `fs := flag.NewFlagSet("claim-invite", flag.ExitOnError)`
- `carbonstack-comms/internal/app/commands.go:161` — `fs := flag.NewFlagSet("register-device", flag.ExitOnError)`
- `carbonstack-comms/internal/app/commands.go:208` — `fs := flag.NewFlagSet("list-devices", flag.ExitOnError)`
- `carbonstack-comms/internal/app/commands.go:255` — `fs := flag.NewFlagSet("fingerprint", flag.ExitOnError)`
- `carbonstack-comms/internal/app/commands.go:280` — `fs := flag.NewFlagSet("verify-device", flag.ExitOnError)`
- `carbonstack-comms/internal/app/commands.go:308` — `fs := flag.NewFlagSet("trust-history", flag.ExitOnError)`
- `carbonstack-comms/internal/app/commands.go:337` — `fs := flag.NewFlagSet("trust-list", flag.ExitOnError)`
- `carbonstack-comms/internal/app/commands.go:365` — `fs := flag.NewFlagSet("simulate-key-change", flag.ExitOnError)`
- `carbonstack-comms/internal/app/commands.go:387` — `fs := flag.NewFlagSet("revoke-device", flag.ExitOnError)`
- `carbonstack-comms/internal/app/commands.go:407` — `fs := flag.NewFlagSet("send", flag.ExitOnError)`
- `carbonstack-comms/internal/app/commands.go:411` — `strict := fs.Bool("strict", false, "block sending to unknown, unverified, or changed devices")`
- `carbonstack-comms/internal/app/commands.go:436` — `return fmt.Errorf("send blocked by trust policy: recipient trust_state=%s", decision.TrustState)`
- `carbonstack-comms/internal/app/commands.go:463` — `fs := flag.NewFlagSet("inbox", flag.ExitOnError)`
- `carbonstack-comms/internal/app/commands.go:498` — `fs := flag.NewFlagSet("ack", flag.ExitOnError)`
- `carbonstack-comms/internal/trust/trust.go:337` — `Warning:    "recipient device is unknown; dev mode allows sending but mature mode should block until verification",`
- `carbonstack-comms/internal/trust/trust.go:377` — `Warning:    "recipient device is not verified; dev mode allows sending but mature mode should block",`
- `carbonstack-comms/internal/trust/trust_test.go:70` — `t.Fatalf("evaluate dev send: %v", err)`
- `carbonstack-comms/internal/trust/trust_test.go:78` — `t.Fatalf("evaluate strict send: %v", err)`
- `carbonstack-comms/internal/trust/trust_test.go:99` — `t.Fatalf("evaluate send: %v", err)`
- `carbonstack-comms/internal/trust/trust_test.go:125` — `t.Fatalf("evaluate send: %v", err)`
- `carbonstack-comms/internal/trust/trust_test.go:151` — `t.Fatalf("evaluate send: %v", err)`
- `carbonstack-comms/internal/state/state_test.go:70` — `t.Fatalf("expected flag server URL, got %q", got)`
- `carbonstack-comms/internal/client/cypher_test.go:90` — `if req["sender_device_id"] != "device-1" {`
- `carbonstack-comms/internal/client/cypher_test.go:91` — `t.Fatalf("expected sender device-1, got %q", req["sender_device_id"])`
- `carbonstack-comms/internal/client/cypher_test.go:186` — `inbox, err := c.Inbox("device-2")`
- `carbonstack-comms/internal/client/cypher_test.go:188` — `t.Fatalf("inbox: %v", err)`
- `carbonstack-comms/internal/client/cypher_test.go:190` — `if len(inbox.Envelopes) != 1 {`
- `carbonstack-comms/internal/client/cypher_test.go:191` — `t.Fatalf("expected 1 envelope, got %d", len(inbox.Envelopes))`
- `carbonstack-comms/internal/client/cypher.go:58` — `SenderDeviceID    string `json:"sender_device_id"``
- `carbonstack-comms/internal/client/cypher.go:124` — `func (c CypherClient) SubmitEnvelope(senderDeviceID string, recipientDeviceID string, contentType string, protocolVersion string, ciphertextB64 string, clientCreatedAt string) (SubmitEnvelopeResponse, error) {`
- `carbonstack-comms/internal/client/cypher.go:127` — `"sender_device_id":    senderDeviceID,`
- `carbonstack-comms/internal/protocol/provider_trust.go:12` — `ProviderTrustActionBlockSend           ProviderTrustAction = "block_send"`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:96` — `t.Fatalf("alice real Cypher inbox failed: %v", err)`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:158` — `t.Fatalf("alice real Cypher inbox after keypackage ack failed: %v", err)`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:161` — `t.Fatalf("expected Alice inbox to be empty after keypackage ack, got %d envelopes", len(aliceInboxAfterAck.Envelopes))`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:179` — `t.Fatalf("bob real Cypher inbox for welcome failed: %v", err)`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:223` — `t.Fatalf("bob real Cypher inbox after welcome ack failed: %v", err)`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:226` — `t.Fatalf("expected Bob inbox to be empty after welcome ack, got %d envelopes", len(bobInboxAfterWelcomeAck.Envelopes))`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:255` — `t.Fatalf("bob real Cypher inbox for application message failed: %v", err)`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:271` — `t.Fatal("expected application-message relay envelope in Bob inbox")`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:289` — `t.Fatalf("bob real Cypher inbox after application-message ack failed: %v", err)`
- `carbonstack-comms/internal/protocol/openmls_sidecar_real_cypher_relay_test.go:292` — `t.Fatalf("expected Bob inbox to be empty after application-message ack, got %d envelopes", len(bobInboxAfterMessageAck.Envelopes))`
- `carbonstack-comms/internal/protocol/mock.go:98` — `return ProtectedMessage{}, state, errors.New("sender device_id is required")`
- `carbonstack-comms/internal/protocol/provider_trust_test.go:26` — `t.Fatalf("happy-path event %q should not block send/receive/open", event)`
- `carbonstack-comms/internal/protocol/provider_trust_test.go:68` — `t.Fatal("missing storage should block send")`
- `carbonstack-comms/internal/protocol/provider_trust_test.go:88` — `t.Fatal("missing signer should block send")`
- `carbonstack-comms/internal/protocol/provider_trust_test.go:129` — `t.Fatal("group unrecoverable should block send, receive, and open")`
- `carbonstack-comms/internal/protocol/provider_trust_test.go:167` — `t.Fatal("unsupported command should not block send, receive, or open")`
- `carbonstack-comms/internal/protocol/provider_trust_test.go:214` — `t.Fatalf("%q should not block send, receive, or open", event)`
- `carbonstack-comms/internal/protocol/provider_trust_test.go:244` — `t.Fatal("identity prep state written should not block send, receive, or open")`
- `carbonstack-comms/internal/protocol/provider_trust_test.go:271` — `t.Fatal("identity exists should not block send, receive, or open")`
- `carbonstack-comms/internal/protocol/provider_trust_test.go:298` — `t.Fatal("checkpoint failed should block send/current outgoing state mutation")`
- `carbonstack-comms/internal/protocol/provider_trust_test.go:319` — `t.Fatal("identity created should not block send, receive, or open")`
- `carbonstack-comms/internal/protocol/provider_trust_test.go:345` — `t.Fatal("identity loaded should not block send, receive, or open")`
- `carbonstack-comms/internal/protocol/provider_trust_test.go:373` — `t.Fatal("identity missing should block current identity-dependent send/outgoing operation")`
- `carbonstack-comms/internal/protocol/provider_trust_test.go:395` — `t.Fatal("public bundle exported should not block send, receive, or open")`
- `carbonstack-comms/internal/protocol/openmls_sidecar_relay_test.go:46` — `inbox, err := cypherClient.Inbox("bob-cypher-device-id")`
- `carbonstack-comms/internal/protocol/openmls_sidecar_relay_test.go:48` — `t.Fatalf("Cypher inbox failed: %v", err)`
- `carbonstack-comms/internal/protocol/openmls_sidecar_relay_test.go:51` — `if len(inbox.Envelopes) != 1 {`
- `carbonstack-comms/internal/protocol/openmls_sidecar_relay_test.go:52` — `t.Fatalf("expected 1 relay envelope, got %d", len(inbox.Envelopes))`
- `carbonstack-comms/internal/protocol/openmls_sidecar_relay_test.go:55` — `envelope := inbox.Envelopes[0]`
- `carbonstack-comms/internal/protocol/openmls_sidecar_relay_test.go:124` — `t.Fatalf("alice Cypher inbox failed: %v", err)`
- `carbonstack-comms/internal/protocol/openmls_sidecar_relay_test.go:192` — `t.Fatalf("alice Cypher inbox after keypackage ack failed: %v", err)`
- `carbonstack-comms/internal/protocol/openmls_sidecar_relay_test.go:195` — `t.Fatalf("expected Alice inbox to be empty after keypackage ack, got %d envelopes", len(aliceInboxAfterAck.Envelopes))`
- `carbonstack-comms/internal/protocol/openmls_sidecar_relay_test.go:217` — `t.Fatalf("bob Cypher inbox failed: %v", err)`
- `carbonstack-comms/internal/protocol/openmls_sidecar_relay_test.go:309` — `t.Fatalf("alice Cypher inbox failed: %v", err)`
- `carbonstack-comms/internal/protocol/openmls_sidecar_relay_test.go:371` — `t.Fatalf("alice Cypher inbox after keypackage ack failed: %v", err)`
- `carbonstack-comms/internal/protocol/openmls_sidecar_relay_test.go:374` — `t.Fatalf("expected Alice inbox to be empty after keypackage ack, got %d envelopes", len(aliceInboxAfterAck.Envelopes))`
- `carbonstack-comms/internal/protocol/openmls_sidecar_relay_test.go:392` — `t.Fatalf("bob Cypher inbox for welcome failed: %v", err)`
- `carbonstack-comms/internal/protocol/openmls_sidecar_relay_test.go:436` — `t.Fatalf("bob Cypher inbox after welcome ack failed: %v", err)`
- `carbonstack-comms/internal/protocol/openmls_sidecar_relay_test.go:439` — `t.Fatalf("expected Bob inbox to be empty after welcome ack, got %d envelopes", len(bobInboxAfterWelcomeAck.Envelopes))`
- `carbonstack-comms/internal/protocol/openmls_sidecar_relay_test.go:468` — `t.Fatalf("bob Cypher inbox for application message failed: %v", err)`
- `carbonstack-comms/internal/protocol/openmls_sidecar_relay_test.go:484` — `t.Fatal("expected application-message relay envelope in Bob inbox")`
- `carbonstack-comms/internal/protocol/openmls_sidecar_relay_test.go:504` — `t.Fatalf("bob Cypher inbox after application-message ack failed: %v", err)`
- `carbonstack-comms/internal/protocol/openmls_sidecar_relay_test.go:507` — `t.Fatalf("expected Bob inbox to be empty after application-message ack, got %d envelopes", len(bobInboxAfterMessageAck.Envelopes))`
- `carbonstack-comms/internal/protocol/types.go:66` — `SenderDeviceID string         `json:"sender_device_id"``
- `carbonstack-comms/internal/protocol/types.go:74` — `SenderDeviceID string            `json:"sender_device_id"``
- `carbonstack-comms/internal/protocol/types.go:90` — `SenderDeviceID string            `json:"sender_device_id"``
- `carbonstack-comms/internal/protocol/openmls_sidecar_test_cypher_server_test.go:20` — `SenderDeviceID    string `json:"sender_device_id"``
- `carbonstack-comms/internal/protocol/openmls_sidecar_test_cypher_server_test.go:66` — `SenderDeviceID    string `json:"sender_device_id"``
- `carbonstack-comms/internal/protocol/openmls_sidecar_helpers_test.go:265` — `cmdArgs := append([]string{"run", "--quiet", "--"}, args...)`
- `carbonstack-comms/internal/protocol/openmls_sidecar_helpers_test.go:266` — `cmd := exec.Command("cargo", cmdArgs...)`
- `carbonstack-comms/internal/relay/cypher_bridge_test.go:104` — `if captured["sender_device_id"] != "alice-device-id" {`
- `carbonstack-comms/internal/relay/cypher_bridge_test.go:105` — `t.Fatalf("sender_device_id = %q", captured["sender_device_id"])`
- `carbonstack-comms/internal/relay/cypher_bridge.go:17` — `senderDeviceID string,`
- `carbonstack-comms/internal/relay/cypher_bridge.go:38` — `senderDeviceID,`
- `carbonstack-comms/internal/relay/test_cypher_server_test.go:17` — `SenderDeviceID    string `json:"sender_device_id"``
- `carbonstack-comms/internal/relay/test_cypher_server_test.go:63` — `SenderDeviceID    string `json:"sender_device_id"``
- `carbonstack-comms/internal/protocol/mls/README.md:25` — `It is not wired into polished runtime `send` / `inbox` UX.`
- `carbonstack-comms/internal/protocol/mls/openmls-sidecar/README.md:27` — `It is not wired into polished CarbonStackComms runtime `send` / `inbox` UX.`
- `carbonstack-comms/internal/protocol/mls/research/openmls-minimal/README.md:64` — `- sends local Alice application message`
- `carbonstack-comms/internal/protocol/mls/research/openmls-minimal/README.md:86` — `- comms send`
- `carbonstack-comms/internal/protocol/mls/research/openmls-minimal/README.md:87` — `- comms inbox`
- `carbonstack-comms/internal/protocol/mls/research/openmls-minimal/README.md:149` — `The scratch crate now sends two sequential Alice-to-Bob application messages inside one process.`
- `carbonstack-comms/cmd/comms/main.go:11` — `if err := app.Run(os.Args[1:]); err != nil {`

Recon interpretation:

- Runtime `send` / `inbox` should remain stub-era until the v0.4.x runtime Comms OpenMLS integration phase.
- v0.3.22 should not try to mature user-facing messaging UX.
- The deployability line should focus on local operator backbone surfaces first.

## 13. Runner / Validation Surface

Runner profile hits:

- `carbonstack/tools/carbonstack-validate/README.md:12` — `### doctor`
- `carbonstack/tools/carbonstack-validate/README.md:16` — `go run . --profile doctor`
- `carbonstack/tools/carbonstack-validate/README.md:18` — `### core`
- `carbonstack/tools/carbonstack-validate/README.md:20` — `Runs the current core validation path:`
- `carbonstack/tools/carbonstack-validate/README.md:22` — `doctor`
- `carbonstack/tools/carbonstack-validate/README.md:25` — `full carbonstack-comms package tests`
- `carbonstack/tools/carbonstack-validate/README.md:26` — `full carbonstack-cypher package tests`
- `carbonstack/tools/carbonstack-validate/README.md:29` — `go run . --profile core`
- `carbonstack/tools/carbonstack-validate/README.md:31` — `### full`
- `carbonstack/tools/carbonstack-validate/README.md:33` — `Currently aliases `core`.`
- `carbonstack/tools/carbonstack-validate/README.md:35` — `go run . --profile full`
- `carbonstack/tools/carbonstack-validate/README.md:37` — ``full` should remain a simple alias until later v0.3.x work creates real release/deployability validation surfaces.`
- `carbonstack/tools/carbonstack-validate/README.md:53` — `go run . --profile core --root /path/to/carbonstack_umbrella`
- `carbonstack/tools/carbonstack-validate/README.md:58` — `go run . --profile doctor`
- `carbonstack/tools/carbonstack-validate/README.md:59` — `go run . --profile core`
- `carbonstack/tools/carbonstack-validate/README.md:65` — `go run . --profile doctor`
- `carbonstack/tools/carbonstack-validate/README.md:66` — `go run . --profile core`
- `carbonstack/tools/carbonstack-validate/README.md:89` — `### release-snapshot`
- `carbonstack/tools/carbonstack-validate/README.md:91` — `Validates a formal release-like package root before calling `core`.`
- `carbonstack/tools/carbonstack-validate/README.md:93` — `go run . --profile release-snapshot --root /path/to/release-package-root`
- `carbonstack/tools/carbonstack-validate/README.md:102` — `The profile checks required repo files, release metadata, and fails if forbidden generated/private/build artifacts are present before tests.`
- `carbonstack/tools/carbonstack-validate/README.md:104` — `After package checks pass, it calls `core`.`
- `carbonstack/tools/carbonstack-validate/README.md:106` — ``release-snapshot` does not package, upload, deploy, clean, install dependencies, or make security claims.`
- `carbonstack/tools/carbonstack-validate/README.md:107` — `## release-snapshot run-order warning`
- `carbonstack/tools/carbonstack-validate/README.md:109` — ``release-snapshot` must be run from a fresh extracted or throwaway staged package root.`
- `carbonstack/tools/carbonstack-validate/README.md:113` — `A successful `release-snapshot` run calls `core`, and `core` generates OpenMLS sidecar state and Rust build artifacts. If that same package root is archived afterward, the archive will contain forbidden generated/private/build artifacts and should fail strict pre-test validation later.`
- `carbonstack/tools/carbonstack-validate/README.md:118` — `archive it without running release-snapshot inside it`
- `carbonstack/tools/carbonstack-validate/README.md:120` — `run release-snapshot from the throwaway extraction`
- `carbonstack/tools/carbonstack-validate/README.md:123` — `Do not run `release-snapshot` twice in the same extraction unless you intentionally expect the second run to fail strict pre-test artifact scanning.`
- `carbonstack/tools/carbonstack-validate/README.md:124` — `## Release checksum helper profiles`
- `carbonstack/tools/carbonstack-validate/README.md:126` — `### write-checksums`
- `carbonstack/tools/carbonstack-validate/README.md:130` — `go run . --profile write-checksums --root /path/to/package-root`
- `carbonstack/tools/carbonstack-validate/README.md:138` — `### verify-checksums`
- `carbonstack/tools/carbonstack-validate/README.md:142` — `go run . --profile verify-checksums --root /path/to/package-root`
- `carbonstack/tools/carbonstack-validate/README.md:144` — `### release-snapshot relationship`
- `carbonstack/tools/carbonstack-validate/README.md:146` — ``release-snapshot` now verifies real checksums before calling `core`.`
- `carbonstack/tools/carbonstack-validate/README.md:152` — `run write-checksums against the package source root`
- `carbonstack/tools/carbonstack-validate/README.md:155` — `validate from a fresh extraction with release-snapshot`
- `carbonstack/tools/carbonstack-validate/README.md:157` — `Do not run `release-snapshot` against the package source root intended for archive/publish.`
- `carbonstack/tools/carbonstack-validate/checksums.go:21` — `r.PrintHeader("write-checksums")`
- `carbonstack/tools/carbonstack-validate/main.go:46` — `profile := flag.String("profile", "doctor", "validation profile: doctor, core, full, release-snapshot, write-checksums, verify-checksums")`
- `carbonstack/tools/carbonstack-validate/main.go:50` — `r, err := NewRunner(*profile, *rootOverride)`
- `carbonstack/tools/carbonstack-validate/main.go:59` — `case "doctor":`
- `carbonstack/tools/carbonstack-validate/main.go:61` — `case "core":`
- `carbonstack/tools/carbonstack-validate/main.go:63` — `case "full":`
- `carbonstack/tools/carbonstack-validate/main.go:64` — `fmt.Println("profile full currently aliases core")`
- `carbonstack/tools/carbonstack-validate/main.go:66` — `case "release-snapshot":`
- `carbonstack/tools/carbonstack-validate/main.go:68` — `case "write-checksums":`
- `carbonstack/tools/carbonstack-validate/main.go:70` — `case "verify-checksums":`
- `carbonstack/tools/carbonstack-validate/main.go:73` — `runErr = fmt.Errorf("unknown profile %q; expected doctor, core, full, release-snapshot, write-checksums, or verify-checksums", r.Profile)`
- `carbonstack/tools/carbonstack-validate/main.go:84` — `func NewRunner(profile string, rootOverride string) (*Runner, error) {`
- `carbonstack/tools/carbonstack-validate/main.go:112` — `Profile:      profile,`
- `carbonstack/tools/carbonstack-validate/main.go:164` — `r.PrintHeader("doctor")`
- `carbonstack/tools/carbonstack-validate/main.go:206` — `r.PrintHeader("core")`
- `carbonstack/tools/carbonstack-validate/main.go:512` — `fmt.Printf("CarbonStack validation profile: %s\n", name)`
- `carbonstack/tools/carbonstack-validate/release_snapshot.go:11` — `r.PrintHeader("release-snapshot")`
- `carbonstack/tools/carbonstack-validate/release_snapshot.go:14` — `fmt.Println("release-snapshot checks run before core validation")`
- `carbonstack/tools/carbonstack-validate/release_snapshot.go:32` — `fmt.Println("release-snapshot package checks and checksum verification passed; calling core validation")`
- `carbonstack/tools/carbonstack-validate/release_snapshot.go:177` — `fmt.Println("== release-snapshot strict pre-test artifact scan ==")`

Runner artifact-scan hits:

- `carbonstack/tools/carbonstack-validate/README.md:23` — `pre-test artifact scan`
- `carbonstack/tools/carbonstack-validate/README.md:24` — `targeted OpenMLS real-Cypher lifecycle test`
- `carbonstack/tools/carbonstack-validate/README.md:27` — `post-test artifact scan`
- `carbonstack/tools/carbonstack-validate/README.md:82` — `Post-test hits are expected only when they stay in known generated roots such as the OpenMLS sidecar `target/` and `.carbonstack-openmls-sidecar-state/`.`
- `carbonstack/tools/carbonstack-validate/README.md:88` — `It does not install dependencies, delete artifacts, package releases, publish releases, configure services, or deploy anything.`
- `carbonstack/tools/carbonstack-validate/README.md:102` — `The profile checks required repo files, release metadata, and fails if forbidden generated/private/build artifacts are present before tests.`
- `carbonstack/tools/carbonstack-validate/README.md:113` — `A successful `release-snapshot` run calls `core`, and `core` generates OpenMLS sidecar state and Rust build artifacts. If that same package root is archived afterward, the archive will contain forbidden generated/private/build artifacts and should fail strict pre-test validation later.`
- `carbonstack/tools/carbonstack-validate/README.md:123` — `Do not run `release-snapshot` twice in the same extraction unless you intentionally expect the second run to fail strict pre-test artifact scanning.`
- `carbonstack/tools/carbonstack-validate/README.md:136` — `The helper excludes generated/private/build artifacts and excludes `release/checksums.txt` itself.`
- `carbonstack/tools/carbonstack-validate/checksums.go:229` — `case ".git", "target", ".carbonstack-openmls-sidecar-state", ".go-cache", ".go-tmp":`
- `carbonstack/tools/carbonstack-validate/checksums.go:246` — `case "provider-storage.json", "signer.json", "Thumbs.db", ".DS_Store":`
- `carbonstack/tools/carbonstack-validate/main.go:213` — `r.ArtifactScan("pre-test")`
- `carbonstack/tools/carbonstack-validate/main.go:264` — `r.ArtifactScan("post-test")`
- `carbonstack/tools/carbonstack-validate/main.go:381` — `fmt.Printf("== %s artifact scan ==\n", phase)`
- `carbonstack/tools/carbonstack-validate/main.go:384` — `"target",`
- `carbonstack/tools/carbonstack-validate/main.go:386` — `"provider-storage.json",`
- `carbonstack/tools/carbonstack-validate/main.go:387` — `"signer.json",`
- `carbonstack/tools/carbonstack-validate/main.go:447` — `fmt.Printf("%s artifact scan: no generated/private/build artifact hits\n", phase)`
- `carbonstack/tools/carbonstack-validate/main.go:451` — `fmt.Printf("%s artifact scan hits:\n", phase)`
- `carbonstack/tools/carbonstack-validate/main.go:457` — `fmt.Println("artifact scan is non-destructive")`
- `carbonstack/tools/carbonstack-validate/main.go:458` — `fmt.Println("pre-test hits are potential source/copy hygiene issues")`
- `carbonstack/tools/carbonstack-validate/main.go:459` — `fmt.Println("post-test hits are expected only when they stay in known generated roots")`
- `carbonstack/tools/carbonstack-validate/main.go:465` — `if strings.Contains(normalized, "internal/protocol/mls/openmls-sidecar/target") \|\|`
- `carbonstack/tools/carbonstack-validate/main.go:467` — `return "known-openmls-sidecar-generated-root"`
- `carbonstack/tools/carbonstack-validate/main.go:471` — `return "research-generated-root"`
- `carbonstack/tools/carbonstack-validate/release_snapshot.go:17` — `fmt.Println("a successful validation generates artifacts, so rerun from a fresh extraction")`
- `carbonstack/tools/carbonstack-validate/release_snapshot.go:177` — `fmt.Println("== release-snapshot strict pre-test artifact scan ==")`
- `carbonstack/tools/carbonstack-validate/release_snapshot.go:181` — `"target",`
- `carbonstack/tools/carbonstack-validate/release_snapshot.go:183` — `"provider-storage.json",`
- `carbonstack/tools/carbonstack-validate/release_snapshot.go:184` — `"signer.json",`
- `carbonstack/tools/carbonstack-validate/release_snapshot.go:239` — `fmt.Println("strict pre-test artifact scan: PASS / no forbidden artifacts")`
- `carbonstack/tools/carbonstack-validate/release_snapshot.go:243` — `fmt.Println("strict pre-test artifact scan: FAIL")`
- `carbonstack/tools/carbonstack-validate/release_snapshot.go:248` — `return fmt.Errorf("release snapshot contains forbidden generated/private/build artifacts before tests")`

Recon interpretation:

- The Go runner remains the preferred validation authority candidate.
- `doctor`, `core`, `full`, `release-snapshot`, `write-checksums`, and `verify-checksums` are the known validation/release-adjacent profiles.
- A future `local-backbone` or similarly named profile may make sense only after the local operator model is concrete enough to validate.
- Do not add a deployability runner profile before its success/failure contract is clear.

## 14. Local Operator Model Recommendation

Recommended current local-only operator model:

```text
CYPHER_ADDR=127.0.0.1:8080
CYPHER_DB=$HOME/.local/share/carbonstack/cypher/cypher.db
CYPHER_MIGRATIONS=<repo-or-installed-migrations-path>
CYPHER_DEV_INVITE=<explicit dev/bootstrap invite>
```

This is a planning target, not yet a validated deployability guarantee.

Recommended state policy:

- `dev-test` profile: temp DB, disposable, runner/test-owned.
- `local-operator` profile: explicit persistent SQLite path, experimental persistence, reset allowed between versions.
- no production migration guarantee yet.
- no public ingress.
- no systemd.
- no cloudflared.
- no real homelab deployment.

## 15. Future Umbrella Helper Direction

A future umbrella helper would be useful during maturity, but should not be rushed.

Possible future shape:

```text
carbonstack local start
carbonstack local validate
carbonstack local stop
carbonstack local reset
```

This should remain a future direction until Cypher config/data/migration behavior is sufficiently stable.

## 16. Recommended Next Rungs

Recommended order:

1. Manually/automatically test repeated Cypher startup against the same SQLite DB.
2. Decide whether v0.3.23 should implement `schema_migrations` tracking or document wipe-only experimental DB behavior.
3. Write a local operator runbook skeleton after the DB persistence stance is clear.
4. Define a stable local data/config directory convention.
5. Only then consider a runner-backed local deployability profile.

## 17. Explicit Nonclaims

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

## 18. Summary

v0.3.22 should be treated as a Cypher local operator surface recon rung.

The highest-priority deployability risk is persistent SQLite migration behavior. The strongest current deployability direction is still local-only, WSL Debian first, simple environment-driven config, `127.0.0.1` bind, explicit data paths, and strict nonclaims.
