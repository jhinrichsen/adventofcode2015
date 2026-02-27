---
# adventofcode2015-lrh6
title: Migrate Day19 to canonical day pattern
status: completed
type: task
created_at: 2026-02-27T11:07:42Z
updated_at: 2026-02-27T11:22:00Z
---

- [x] Replace Day19Part1/Day19Part2 with NewDay19 + Day19(puzzle, part1 bool) uint
- [x] Keep iterative reducer strategy and modernize internals
- [x] Convert part tests and add one-line benchmarks
- [x] Run Day19 tests and benchmarks

## Summary of Changes

- Replaced legacy Day19 APIs with parser+solver pattern: `NewDay19` and `Day19(puzzle, part1 bool) uint`.
- Preserved iterative reduction strategy for part2 and modernized sorting via `slices.SortFunc`.
- Converted part tests to one-line parser helper calls and added one-line benchmarks.
- Kept focused unit coverage for `replaceNth` and both sample scenarios.
- Verified with Day19-focused tests and benchmarks.
