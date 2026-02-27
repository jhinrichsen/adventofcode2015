---
# adventofcode2015-ifvj
title: Migrate Day17 to canonical day pattern
status: completed
type: task
created_at: 2026-02-27T11:05:12Z
updated_at: 2026-02-27T11:13:00Z
---

- [x] Replace Day17Part1/Day17Part2 with NewDay17 + Day17(puzzle, part1 bool) uint
- [x] Remove PowerSet allocation pattern in favor of direct combination scan
- [x] Convert tests and add one-line benchmarks
- [x] Run Day17 tests and benchmarks

## Summary of Changes

- Replaced legacy Day17 part APIs with parser+solver pattern: `NewDay17` and `Day17(puzzle, part1 bool) uint`.
- Removed `PowerSet` allocation approach and switched to direct bitmask combination scanning.
- Added one-line part tests and benchmarks using shared parser helpers.
- Kept example coverage through direct helper calls at 25-liter target.
- Verified with Day17-focused tests and benchmarks.
