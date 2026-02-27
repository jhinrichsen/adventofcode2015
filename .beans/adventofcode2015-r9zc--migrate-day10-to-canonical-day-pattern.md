---
# adventofcode2015-r9zc
title: Migrate Day10 to canonical day pattern
status: completed
type: task
created_at: 2026-02-27T10:48:20Z
updated_at: 2026-02-27T10:51:00Z
---

- [x] Add testdata/day10.txt input fixture
- [x] Implement NewDay10 parser + Day10(puzzle, part1 bool) uint
- [x] Convert part tests and benchmarks to one-line parser helper pattern
- [x] Run Day10 tests and benchmarks

## Summary of Changes

- Added `testdata/day10.txt` and moved part tests to load input via shared file helpers.
- Added parser `NewDay10(lines []string) (Day10Puzzle, error)`.
- Added canonical solver `Day10(puzzle Day10Puzzle, part1 bool) uint`.
- Migrated part tests and benchmarks to one-line `testWithParser` / `benchWithParser` pattern.
- Verified with Day10-focused tests and benchmarks.
