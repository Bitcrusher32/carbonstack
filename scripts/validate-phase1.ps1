$ErrorActionPreference = "Stop"

Write-Host "CarbonStack Phase 1 Local Validation"
Write-Host "===================================="

$CarbonStackRepo = Split-Path -Parent $PSScriptRoot
$Root = Split-Path -Parent $CarbonStackRepo

$Cypher = Join-Path $Root "carbonstack-cypher"
$Comms = Join-Path $Root "carbonstack-comms"
$CarbonStack = Join-Path $Root "carbonstack"

function Step {
    param([string]$Name)
    Write-Host ""
    Write-Host "==> $Name"
}

function Run-In {
    param(
        [string]$Path,
        [string]$Command
    )

    Push-Location $Path
    try {
        Write-Host "> $Command"
        powershell -NoProfile -ExecutionPolicy Bypass -Command $Command
        if ($LASTEXITCODE -ne 0) {
            throw "command failed with exit code $LASTEXITCODE"
        }
    }
    finally {
        Pop-Location
    }
}

Step "Repo status summary"
Get-ChildItem -Path $Root -Directory | Where-Object {
    Test-Path (Join-Path $_.FullName ".git")
} | ForEach-Object {
    Write-Host ""
    Write-Host "=== $($_.Name) ==="
    git -C $_.FullName status --short
    git -C $_.FullName log --oneline -2
}

Step "Check required Phase 1 files"

$RequiredFiles = @(
    "carbonstack\docs\16-phase1-integration-plan.md",
    "carbonstack\docs\17-phase1-test-matrix.md",
    "carbonstack-cypher\cmd\cypher\main.go",
    "carbonstack-cypher\internal\httpapi\api.go",
    "carbonstack-cypher\internal\httpapi\api_test.go",
    "carbonstack-cypher\migrations\001_init.sql",
    "carbonstack-comms\cmd\comms\main.go",
    "carbonstack-comms\internal\app\commands.go",
    "carbonstack-comms\internal\client\cypher.go",
    "carbonstack-comms\internal\client\cypher_test.go",
    "carbonstack-comms\internal\crypto\mock.go",
    "carbonstack-comms\internal\crypto\mock_test.go",
    "carbonstack-comms\internal\state\state.go",
    "carbonstack-comms\internal\state\state_test.go",
    "carbonstack-comms\scripts\test-local-lifecycle.ps1"
)

foreach ($File in $RequiredFiles) {
    $FullPath = Join-Path $Root $File
    if (-not (Test-Path $FullPath)) {
        throw "missing required file: $File"
    }
    Write-Host "OK $File"
}

Step "Run Cypher tests"
Run-In $Cypher "go test ./..."

Step "Run Comms package tests"
Run-In $Comms "go test ./..."

Step "Start temporary Cypher server for lifecycle smoke test"

$TempDir = Join-Path $Root ".phase1-validation-temp"
$TempDb = Join-Path $TempDir "cypher-validation.db"

Remove-Item -Recurse -Force $TempDir -ErrorAction SilentlyContinue
New-Item -ItemType Directory -Force -Path $TempDir | Out-Null

$OldCypherDb = $Env:CYPHER_DB
$OldCypherAddr = $Env:CYPHER_ADDR
$OldCypherDevInvite = $Env:CYPHER_DEV_INVITE

$Env:CYPHER_DB = $TempDb
$Env:CYPHER_ADDR = ":8080"
$Env:CYPHER_DEV_INVITE = "phase1-validation-root"

$CypherProcess = Start-Process `
    -FilePath "go" `
    -ArgumentList "run .\cmd\cypher" `
    -WorkingDirectory $Cypher `
    -PassThru `
    -WindowStyle Hidden

try {
    $Healthy = $false

    for ($i = 0; $i -lt 30; $i++) {
        try {
            $Health = Invoke-RestMethod -Method GET -Uri "http://localhost:8080/v0/health"
            if ($Health.status -eq "ok") {
                $Healthy = $true
                break
            }
        }
        catch {
            Start-Sleep -Milliseconds 500
        }
    }

    if (-not $Healthy) {
        throw "Cypher server did not become healthy"
    }

    Step "Run Comms lifecycle smoke test"
Run-In $Comms "powershell -ExecutionPolicy Bypass -File .\scripts\test-local-lifecycle.ps1"

Step "Run Comms Phase 2A trust lifecycle smoke test"
Run-In $Comms "powershell -ExecutionPolicy Bypass -File .\scripts\test-trust-lifecycle.ps1"
}
finally {
    if ($CypherProcess -and -not $CypherProcess.HasExited) {
        Stop-Process -Id $CypherProcess.Id -Force
    }

    if ($null -ne $OldCypherDb) { $Env:CYPHER_DB = $OldCypherDb } else { Remove-Item Env:\CYPHER_DB -ErrorAction SilentlyContinue }
    if ($null -ne $OldCypherAddr) { $Env:CYPHER_ADDR = $OldCypherAddr } else { Remove-Item Env:\CYPHER_ADDR -ErrorAction SilentlyContinue }
    if ($null -ne $OldCypherDevInvite) { $Env:CYPHER_DEV_INVITE = $OldCypherDevInvite } else { Remove-Item Env:\CYPHER_DEV_INVITE -ErrorAction SilentlyContinue }

    Remove-Item -Recurse -Force $TempDir -ErrorAction SilentlyContinue
}

Step "Validation complete"
Write-Host "PASS: CarbonStack Phase 1 local validation passed"

