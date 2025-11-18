# Build-GoProject.ps1

$projectRoot = Resolve-Path .
$outputDir = Join-Path $projectRoot "packages\umbr-download\bin"

# Ensure output directory exists
if (-not (Test-Path $outputDir)) {
    New-Item -ItemType Directory -Path $outputDir -Force | Out-Null
}

# Platforms to build for
$platforms = @(
    @{GOOS="windows"; GOARCH="amd64"},
    @{GOOS="linux"; GOARCH="amd64"},
    @{GOOS="darwin"; GOARCH="amd64"}
)

# Project name
$projectName = "go-download"

# Path to CLI project
$cliPath = Join-Path $projectRoot "cli"

foreach ($platform in $platforms) {
    $os = $platform.GOOS
    $arch = $platform.GOARCH
    $outputName = "$projectName-$os"
    
    if ($os -eq "windows") {
        $outputName += ".exe"
    }

    $outputPath = Join-Path $outputDir $outputName

    Write-Host "Building for $os/$arch -> $outputPath"

    # Set environment variables for cross-compilation
    $env:GOOS = $os
    $env:GOARCH = $arch

    # Build command
    Push-Location $cliPath
    & go build -o $outputPath -ldflags "-s -w" .
    Pop-Location

    if ($LASTEXITCODE -ne 0) {
        Write-Host "Build failed for $os/$arch" -ForegroundColor Red
    } else {
        Write-Host "Build succeeded for $os/$arch" -ForegroundColor Green
    }
}

Write-Host "All builds completed!"
