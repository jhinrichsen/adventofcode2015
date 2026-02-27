---
# adventofcode2015-bxs0
title: Migrate Day18 to canonical day pattern
status: completed
type: task
created_at: 2026-02-27T11:06:21Z
updated_at: 2026-02-27T11:17:00Z
---

- [x] Replace Day18Part1/Day18Part2 with NewDay18 + Day18(puzzle, part1 bool) uint
- [x] Remove legacy file-I/O test helpers and migrate to one-line part tests/benchmarks
- [x] Keep concise sample evolution coverage
- [x] Run Day18 tests and benchmarks

## Summary of Changes

- Replaced legacy Day18 part APIs with `NewDay18` parser and canonical `Day18(puzzle, part1 bool) uint`.
- Implemented iterative grid stepping with explicit neighbor counting and optional stuck corners for part2.
- Removed legacy file-I/O test helpers and migrated part tests/benchmarks to shared one-line parser helpers.
- Kept concise sample evolution and part2 sample coverage.
- Verified with Day18-focused tests and benchmarks.
