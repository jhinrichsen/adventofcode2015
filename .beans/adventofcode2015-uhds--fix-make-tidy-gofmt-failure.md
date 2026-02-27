---
# adventofcode2015-uhds
title: Fix make tidy gofmt failure
status: completed
type: task
priority: normal
created_at: 2026-02-27T14:39:21Z
updated_at: 2026-02-27T14:40:21Z
---

Identify unformatted Go files, run gofmt, and verify make tidy passes.

## Summary of Changes

- Formatted all unformatted Go files reported by gofmt.
- Removed unused functions that failed golangci-lint (, , , ).
- Verified test -z "$(gofmt -l .)"
CGO_ENABLED=0 go vet
CGO_ENABLED=0 go run github.com/golangci/golangci-lint/cmd/golangci-lint@latest run now passes (, , Smart, fast linters runner.

Usage:
  golangci-lint [flags]
  golangci-lint [command]

Available Commands:
  cache       Cache control and information.
  completion  Generate the autocompletion script for the specified shell
  config      Configuration file information and verification.
  custom      Build a version of golangci-lint with custom linters.
  fmt         Format Go source files.
  formatters  List current formatters configuration.
  help        Display extra help
  linters     List current linters configuration.
  migrate     Migrate configuration file from v1 to v2.
  run         Lint the code.
  version     Display the golangci-lint version.

Flags:
      --color string   Use color when printing; can be 'always', 'auto', or 'never' (default "auto")
  -h, --help           Help for a command
  -v, --verbose        Verbose output
      --version        Print version

Use "golangci-lint [command] --help" for more information about a command.).

## Summary Correction

- Removed unused functions: keys, day16PropID, day21ItemCombinations, benchLinesErr.
- make tidy passes end-to-end.
