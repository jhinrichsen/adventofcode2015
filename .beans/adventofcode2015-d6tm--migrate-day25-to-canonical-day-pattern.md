---
# adventofcode2015-d6tm
title: Migrate Day25 to canonical day pattern
status: completed
type: task
priority: normal
created_at: 2026-02-27T11:20:26Z
updated_at: 2026-02-27T11:31:00Z
---

- [x] Add NewDay25 parser + Day25(puzzle, part1 bool) uint
- [x] Keep code-sequence internals unexported
- [x] Add testdata/day25.txt and part test via helper
- [x] Run Day25 tests and benchmark

## Summary of Changes

- Added parser `NewDay25(lines)` and canonical solver `Day25(puzzle, part1 bool) uint`.
- Moved target coordinate into `testdata/day25.txt` and wired part test through `testWithParser`.
- Kept sequence internals unexported (`day25State`, `day25CodeAt`).
- Migrated benchmark to helper pattern with `BenchmarkDay25Part1`.
- Verified with Day25-focused tests and benchmark.
