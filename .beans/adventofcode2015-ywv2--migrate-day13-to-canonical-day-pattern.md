---
# adventofcode2015-ywv2
title: Migrate Day13 to canonical day pattern
status: completed
type: task
created_at: 2026-02-27T11:00:36Z
updated_at: 2026-02-27T11:02:00Z
---

- [x] Replace Day13Part1/Day13Part2 with NewDay13 + Day13(puzzle, part1 bool) uint
- [x] Convert tests and add one-line benchmarks
- [x] Run Day13 tests and benchmarks

## Summary of Changes

- Replaced legacy Day13 part APIs with parser+solver pattern: `NewDay13` and `Day13(puzzle, part1 bool) uint`.
- Migrated tests to one-line helper calls and added one-line benchmarks.
- Kept permutation-based iterative solver for both parts (part2 includes extra neutral attendee).
- Verified with Day13-focused tests and benchmarks.
