---
# adventofcode2015-x7dv
title: Migrate Day24 to canonical day pattern
status: completed
type: task
priority: normal
created_at: 2026-02-27T11:19:17Z
updated_at: 2026-02-27T11:27:00Z
---

- [x] Replace Day24Part1/Day24Part2 with NewDay24 + Day24(puzzle, part1 bool) uint
- [x] Remove panic path for unsplittable totals
- [x] Convert part tests to parser helper one-liners
- [x] Run Day24 tests (short mode for heavy cases)

## Summary of Changes

- Added parser `NewDay24(lines)` and canonical solver `Day24(puzzle, part1 bool) uint`.
- Replaced legacy part entrypoints and removed panic on invalid group split (returns `0` instead).
- Migrated all Day24 part tests to one-line parser helper pattern.
- Fixed part2 test path to call part2 solver (previously pointed at part1 logic).
- Verified with `go test -short -run '^TestDay24' .`.
