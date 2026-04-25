# wolfx_cli

A lightweight Cobra-based CLI for passive web scanning and basic security header checks.

## Current stage

This project is in the **initial working stage**. It currently supports a basic `scan` command that:

* Accepts a target URL
* Sends a passive HTTP/HTTPS request
* Prints response status
* Detects missing common security headers
* Extracts page title
* Shows basic TLS certificate details when available
* Supports JSON output

## Project structure

```text
wolfx_cli/
├── cmd/
│   ├── root.go
│   ├── scan.go
│   └── version.go
├── internal/
│   ├── models/
│   │   └── result.go
│   ├── output/
│   │   └── output.go
│   └── scanner/
│       ├── checks.go
│       └── scanner.go
├── main.go
└── go.mod
```

## Requirements

* Go 1.22 or newer
* Internet access for dependency download on first run

## Install Go on Kali Linux

If `go` is not installed, install it first.

### Option 1: Install from package manager

```bash
sudo apt update
sudo apt install golang -y
```

Check the installation:

```bash
go version
```

### Option 2: Manual install from Go official release

If the package manager version fails, install Go manually:

```bash
wget https://go.dev/dl/go1.22.5.linux-amd64.tar.gz
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf go1.22.5.linux-amd64.tar.gz
```

Add Go to PATH:

```bash
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc
```

Verify:

```bash
go version
```

## Setup the project

Go to the project folder:

```bash
cd ~/Desktop/wolfx_cli
```

Download dependencies:

```bash
go mod tidy
```

If you get import errors, make sure the `module` name in `go.mod` matches your import paths.

Example:

```go
module github.com/yourorg/wolfx_cli
```

Then imports should look like:

```go
import "github.com/yourorg/wolfx_cli/internal/scanner"
```

## Run the CLI

### Run directly

```bash
go run . scan --url https://example.com
```

### Run with JSON output

```bash
go run . scan --url https://example.com --json
```

### Run with verbose output

```bash
go run . scan --url https://example.com --verbose
```

### Build a binary

```bash
go build -o wolfx
./wolfx scan --url https://example.com
```

## Example output

```text
== WebScan Result ==
Target:          https://example.com
Final URL:       https://example.com/
Status Code:     200
Response Time:   123 ms
Server:          example-server
Title:           Example Domain
TLS:             enabled
Missing Headers:
  - Content-Security-Policy
  - X-Frame-Options
```

## Command reference

### `scan`

Run a passive scan against a target URL.

Flags:

* `-u, --url` target URL
* `--timeout` request timeout, example `15s`
* `--json` output results in JSON
* `--insecure` skip TLS certificate verification
* `--follow-redirects` follow HTTP redirects
* `--user-agent` custom user agent
* `--max-body-bytes` maximum response body size to read
* `-v, --verbose` show extra details

Example:

```bash
go run . scan --url https://example.com --timeout 15s --verbose
```

### `version`

Print the current version.

```bash
go run . version
```

## Notes

This tool is designed for **authorized testing only**.

It currently performs passive checks only and does not run exploit payloads.

## Next planned features

* Scan multiple URLs from a file
* Concurrency support
* Export reports to JSON, CSV, and HTML
* Directory discovery module
* Proxy support
* Config file support
* Plugin-style checks

## Troubleshooting

### `go: command not found`

Install Go and make sure it is in your PATH.

### `module not found` or import errors

Check the module path in `go.mod` and update all internal imports to match it.

### `404` while installing Go from apt

Update the package list or use the manual Go install method.

## License

Cyber wolfx Team...
