---
# adventofcode2015-0ur6
title: Migrate Day07 to parser+solver one-liner tests
status: completed
type: task
priority: normal
created_at: 2026-02-27T10:25:38Z
updated_at: 2026-02-27T10:26:24Z
---

- [x] Rename day07 testdata files to zero-padded names
- [x] Replace Day07Part1/Day07Part2 with canonical Day07(puzzle, part1 bool) uint
- [x] Keep/adjust NewDay07 parser and internal evaluator helpers
- [x] Convert day07 tests/benchmarks to one-line helper pattern
- [x] Run Day07 tests and benchmarks

## Summary of Changes

- Renamed Day07 input and example fixtures to zero-padded filenames (`testdata/day07*.txt`).
- Replaced legacy `Day07Part1`/`Day07Part2` API with canonical `Day07(puzzle, part1 bool) uint`.
- Kept parser/evaluator internals and switched wire evaluation to ID-based lookup for tests/internal solver calls.
- Converted Day07 part tests and benchmarks to shared one-line helper pattern (`testWithParser` / `benchWithParser`).
- Verified with `go test -run ^TestDay07 .` and `go test -run ^$ -bench ^BenchmarkDay07 .`.
