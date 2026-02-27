---
# adventofcode2015-pv3y
title: Migrate Day08 to canonical day pattern
status: completed
type: task
priority: normal
created_at: 2026-02-27T10:27:56Z
updated_at: 2026-02-27T10:29:04Z
---

- [x] Rename day8 testdata files to zero-padded names
- [x] Replace Day8Part1/Day8Part2 with Day08(lines, part1 bool) uint
- [x] Convert tests and benchmarks to one-line helper pattern
- [x] Run Day08 tests and benchmarks

## Summary of Changes

- Renamed Day08 fixtures to zero-padded names (`testdata/day08.txt`, `testdata/day08_example.txt`).
- Replaced legacy exported part functions with canonical `Day08(lines []string, part1 bool) uint`.
- Added internal per-line calculators for memory length and encoded length.
- Rewrote Day08 tests and benchmarks to one-line shared helper pattern (`testLines` / `benchLines`).
- Verified with Day08-focused tests and benchmarks.
