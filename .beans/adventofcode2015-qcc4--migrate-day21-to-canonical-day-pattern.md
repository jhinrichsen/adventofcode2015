---
# adventofcode2015-qcc4
title: Migrate Day21 to canonical day pattern
status: completed
type: task
created_at: 2026-02-27T11:10:20Z
updated_at: 2026-02-27T11:34:00Z
---

- [x] Replace Day21Part1/Day21Part2 with NewDay21 + Day21(puzzle, part1 bool) uint
- [x] Add testdata/day21.txt and parse boss stats
- [x] Convert part tests/benchmarks to one-line helper pattern
- [x] Run Day21 tests and benchmarks

## Summary of Changes

- Replaced constant-driven Day21 logic with parser+solver pattern: `NewDay21` and `Day21(puzzle, part1 bool) uint`.
- Added `testdata/day21.txt` and parser for boss hit points, damage, and armor.
- Implemented canonical solver with full legal gear combinations (0/1/2 rings), minimal win cost and maximal loss cost.
- Converted part tests and benchmarks to one-line parser helper pattern.
- Kept compact duel sanity test for naked-player loss.
- Verified with Day21-focused tests and benchmarks.
