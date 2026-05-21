$ErrorActionPreference = "Stop"

$ScriptDir = Split-Path -Parent $MyInvocation.MyCommand.Path
$LocalRunner = Join-Path $ScriptDir "validate-local.ps1"

powershell -NoProfile -ExecutionPolicy Bypass -File $LocalRunner
exit $LASTEXITCODE
