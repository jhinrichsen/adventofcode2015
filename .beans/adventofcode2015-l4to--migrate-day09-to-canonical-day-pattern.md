---
# adventofcode2015-l4to
title: Migrate Day09 to canonical day pattern
status: completed
type: task
priority: normal
created_at: 2026-02-27T10:43:32Z
updated_at: 2026-02-27T10:44:36Z
---

- [x] Rename day9 testdata files to zero-padded names
- [x] Replace legacy Day9 API with NewDay09 + Day09(puzzle, part1 bool) uint
- [x] Convert tests and benchmarks to one-line helper pattern
- [x] Run Day09 tests and benchmarks

## Summary of Changes

- Renamed Day09 fixtures to zero-padded names (`testdata/day09.txt`, `testdata/day09_example.txt`).
- Replaced legacy `Day9(lines) (min,max,error)` API with `NewDay09(lines)` parser and canonical `Day09(puzzle, part1 bool) uint` solver.
- Updated tests to one-line parser-helper pattern and added one-line Day09 benchmarks.
- Restored shared `keys(map[string]bool)` helper used by Day13 after refactor extraction.
- Verified with `go test -run ^TestDay09 .` and `go test -run ^$ -bench ^BenchmarkDay09 .`.
