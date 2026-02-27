---
# adventofcode2015-eqr4
title: Migrate Day15 to canonical day pattern
status: completed
type: task
created_at: 2026-02-27T11:03:00Z
updated_at: 2026-02-27T11:07:00Z
---

- [x] Replace Day15Part1/Day15Part2 with NewDay15 + Day15(puzzle, part1 bool) uint
- [x] Unexport day15 internals and simplify solver loop
- [x] Convert tests and benchmarks to one-line helper pattern
- [x] Run Day15 tests and benchmarks

## Summary of Changes

- Replaced legacy Day15 part APIs with `NewDay15` parser and canonical `Day15(puzzle, part1 bool) uint`.
- Unexported day15 internals and simplified scoring/calorie evaluation with direct totals.
- Kept composition generator workflow and iterated candidate mixes directly.
- Converted tests/benchmarks to one-line parser helper style for parts.
- Verified with Day15-focused tests and benchmarks.
