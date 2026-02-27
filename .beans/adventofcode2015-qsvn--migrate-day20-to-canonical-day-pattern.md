---
# adventofcode2015-qsvn
title: Migrate Day20 to canonical day pattern
status: completed
type: task
created_at: 2026-02-27T11:09:02Z
updated_at: 2026-02-27T11:27:00Z
---

- [x] Replace constant-driven Day20Part1/Day20Part2 with NewDay20 + Day20(puzzle, part1 bool) uint
- [x] Add testdata/day20.txt and parser for target value
- [x] Convert part tests/benchmarks to one-line helper pattern
- [x] Run Day20 tests and benchmarks

## Summary of Changes

- Replaced constant-driven Day20 APIs with parser+solver pattern: `NewDay20` and `Day20(puzzle, part1 bool) uint`.
- Added `testdata/day20.txt` as single-line puzzle input target.
- Implemented sieve-based part1 and part2 solvers in canonical path.
- Converted part tests and benchmarks to one-line parser helper pattern.
- Kept concise presents sample checks via internal helper.
- Verified with Day20-focused tests and benchmarks.
