---
# adventofcode2015-xtw0
title: Migrate Day16 to canonical day pattern
status: completed
type: task
created_at: 2026-02-27T11:04:06Z
updated_at: 2026-02-27T11:10:00Z
---

- [x] Replace manual Day16 logic with NewDay16 + Day16(puzzle, part1 bool) uint
- [x] Fix part2 property rule handling in solver
- [x] Convert tests and add one-line benchmarks
- [x] Run Day16 tests and benchmarks

## Summary of Changes

- Migrated Day16 to parser+solver pattern: `NewDay16` and `Day16(puzzle, part1 bool) uint`.
- Replaced manual test-side scanning with one-line parser helper part tests and one-line benchmarks.
- Kept parse coverage with compact parser unit test.
- Fixed part2 property matching rules (`cats/trees` greater-than and `pomeranians/goldfish` less-than).
- Verified with Day16-focused tests and benchmarks.
