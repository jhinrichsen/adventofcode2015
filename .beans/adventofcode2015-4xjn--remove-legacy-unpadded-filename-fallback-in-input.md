---
# adventofcode2015-4xjn
title: Remove legacy unpadded filename fallback in input_test helpers
status: completed
type: task
priority: high
created_at: 2026-02-27T09:26:05Z
updated_at: 2026-02-27T11:57:25Z
---

Drop fallback logic in input_test.go filename helpers and use strict zero-padded day file naming only (dayXX/dayXX_example...). Update remaining testdata filenames and verify tests.


## Summary of Changes

- Removed unpadded filename fallback logic from `input_test.go` helpers (`filename`, `exampleFilename`, `exampleNFilename`).
- Renamed remaining unpadded testdata files to strict zero-padded naming: `day01.txt`, `day02.txt`, `day03.txt`, `day05.txt`, `day06.txt`.
- Verified with `go test -short ./...`.
