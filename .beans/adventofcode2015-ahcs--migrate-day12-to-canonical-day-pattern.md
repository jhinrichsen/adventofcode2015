---
# adventofcode2015-ahcs
title: Migrate Day12 to canonical day pattern
status: completed
type: task
created_at: 2026-02-27T10:51:35Z
updated_at: 2026-02-27T10:54:00Z
---

- [x] Replace legacy day12 helpers with canonical Day12(buf, part1 bool) (uint, error)
- [x] Convert part tests and benchmarks to one-line input_test.go helpers
- [x] Keep/trim sample tests to concise table tests without manual file I/O
- [x] Run Day12 tests and benchmarks

## Summary of Changes

- Replaced legacy helpers with canonical `Day12(buf []byte, part1 bool) (uint, error)`.
- Implemented iterative part1 byte-scan number summation and iterative part2 JSON walk with red-object exclusion.
- Migrated part tests/benchmarks to one-line shared helpers (`testSolver` / `bench`).
- Kept concise table-driven sample coverage for both parts without manual file I/O.
- Verified with Day12-focused tests and benchmarks.
