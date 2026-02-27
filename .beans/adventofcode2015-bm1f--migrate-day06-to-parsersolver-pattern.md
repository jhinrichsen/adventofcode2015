---
# adventofcode2015-bm1f
title: Migrate Day06 to parser+solver pattern
status: completed
type: task
priority: normal
created_at: 2026-02-27T10:05:55Z
updated_at: 2026-02-27T11:13:00Z
---

- [x] Refactor day06.go to NewDay06 parser + Day06 solver
- [x] Remove Day6Part1/Day6Part2 legacy entrypoints
- [x] Update day06_test.go to TestDay06*/BenchmarkDay06* one-liners via testWithParser/benchWithParser
- [x] Run Day06-focused tests and benchmarks

Summary:
- Added typed parser model (`Day06Puzzle`) and parser entrypoint (`NewDay06`) for instruction parsing.
- Unified both parts under canonical `Day06(puzzle, part1 bool) uint` solver.
- Migrated tests/benchmarks to two-digit names and shared one-line parser helpers.
