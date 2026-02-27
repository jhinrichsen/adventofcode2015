---
# adventofcode2015-v8rn
title: Migrate Day14 to canonical day pattern
status: completed
type: task
created_at: 2026-02-27T11:01:46Z
updated_at: 2026-02-27T11:05:00Z
---

- [x] Replace Day14Part1/Day14Part2 with NewDay14 + Day14(puzzle, part1 bool) uint
- [x] Unexport day14 internal structs/helpers
- [x] Convert part tests and add one-line benchmarks
- [x] Run Day14 tests and benchmarks

## Summary of Changes

- Replaced legacy Day14 part APIs with `NewDay14` parser and canonical `Day14(puzzle, part1 bool) uint`.
- Unexported Day14 internals (`day14Reindeer`, parse/distance/score helpers).
- Converted part tests to one-line helper calls and added one-line benchmarks.
- Kept concise example coverage using helper functions at 1000-second sample duration.
- Verified with Day14-focused tests and benchmarks.
