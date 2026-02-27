---
# adventofcode2015-38d4
title: Migrate Day11 to canonical day pattern
status: completed
type: task
created_at: 2026-02-27T10:50:28Z
updated_at: 2026-02-27T10:52:00Z
---

- [x] Add testdata/day11.txt input fixture
- [x] Add NewDay11 parser + Day11(puzzle, part1 bool) solver
- [x] Convert part tests to one-line helper pattern and add one-line benchmarks
- [x] Run Day11 tests and benchmarks

## Summary of Changes

- Added `testdata/day11.txt` and moved part tests to shared file-based helper calls.
- Added parser `NewDay11(lines []string) (Day11Puzzle, error)`.
- Added canonical solver `Day11(puzzle Day11Puzzle, part1 bool) string`.
- Added one-line Day11 part benchmarks via `benchWithParser`.
- Updated `req2` to byte-based ASCII scanning.
- Verified with Day11-focused tests and benchmarks.
