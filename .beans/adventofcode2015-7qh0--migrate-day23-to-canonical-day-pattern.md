---
# adventofcode2015-7qh0
title: Migrate Day23 to canonical day pattern
status: completed
type: task
priority: normal
created_at: 2026-02-27T11:18:05Z
updated_at: 2026-02-27T11:24:00Z
---

- [x] Add NewDay23 + Day23(puzzle, part1 bool) uint
- [x] Keep VM execution internal with parser-normalized instructions
- [x] Convert part tests/benchmarks to helper one-liners
- [x] Run Day23 tests and benchmarks

## Summary of Changes

- Added `NewDay23(lines)` parser and canonical solver `Day23(puzzle, part1 bool) uint`.
- Kept VM interpreter internal via `day23Run(...)` with normalized comma-free instructions.
- Removed panic behavior for unknown registers; execution now exits safely.
- Migrated part tests and benchmarks to one-line shared helpers and retained concise example coverage.
- Verified with Day23-focused tests and benchmarks.
