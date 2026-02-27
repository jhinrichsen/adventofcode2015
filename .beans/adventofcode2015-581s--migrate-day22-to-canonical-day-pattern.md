---
# adventofcode2015-581s
title: Migrate Day22 to canonical day pattern
status: completed
type: task
priority: normal
created_at: 2026-02-27T11:15:42Z
updated_at: 2026-02-27T11:20:00Z
---

- [x] Add NewDay22 parser + Day22(puzzle, part1 bool) uint
- [x] Wire day22 solver to parsed boss stats instead of constants
- [x] Add testdata/day22.txt and migrate part tests to parser helper
- [x] Run Day22 tests

## Summary of Changes

- Added parser `NewDay22(lines []string) (Day22Puzzle, error)` and canonical solver `Day22(puzzle, part1 bool) uint`.
- Wired brute-force solver to parsed boss stats instead of hardcoded values.
- Added `testdata/day22.txt` for file-based part tests.
- Migrated part tests to one-line `testWithParser` helpers; kept example/simulator tests unchanged.
- Verified with `go test -short -run '^TestDay22' .`.
