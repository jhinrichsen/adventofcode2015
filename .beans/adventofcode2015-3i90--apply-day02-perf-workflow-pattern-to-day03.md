---
# adventofcode2015-3i90
title: Apply perf workflow pattern to Day03
status: completed
type: task
priority: normal
created_at: 2026-02-27T11:48:58Z
updated_at: 2026-02-27T11:59:30Z
---

Performance workflow pattern:

- [x] Capture b0 baseline benchmark in benches/dayXX-b0.txt (trim PASS/ok lines)
- [x] Apply one targeted performance optimization
- [x] Capture b1 benchmark in benches/dayXX-b1.txt (trim PASS/ok lines)
- [x] Run benchstat b0 vs b1
- [x] Optionally iterate with b2..bN (one change per iteration), each compared via benchstat
- [x] Write dayXX.adoc with:
  - b0 baseline summary
  - b1 (or bN) benchstat delta summary
  - concise optimization notes
  - raw benchstat block
- [x] Keep README untouched (no include) unless explicitly requested
- [x] Verify day tests and benchmark command still pass

Scope: Day03 only.


## Summary of Changes

- Captured Day03 baseline in `benches/day03-b0.txt`.
- Optimized `Day03` by switching visited-house tracking from `map[image.Point]bool` to packed coordinate keys in `map[uint64]struct{}` and inlining moves.
- Captured post-change benchmark in `benches/day03-b1.txt` and compared with benchstat.
- Added `day03.adoc` with b0 baseline summary, b1 benchstat delta summary, optimization notes, and raw benchstat block.
- Kept README untouched and reverified Day03 tests and benchmark command.
