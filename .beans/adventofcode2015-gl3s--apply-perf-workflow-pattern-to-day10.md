---
# adventofcode2015-gl3s
title: Apply perf workflow pattern to Day10
status: completed
type: task
priority: normal
created_at: 2026-02-27T11:53:19Z
updated_at: 2026-02-27T14:26:03Z
---

Performance workflow pattern:

- [x] Capture b0 baseline benchmark in benches/day10-b0.txt (trim PASS/ok lines)
- [x] Apply one targeted performance optimization
- [x] Capture b1 benchmark in benches/day10-b1.txt (trim PASS/ok lines)
- [x] Run benchstat b0 vs b1
- [x] Optionally iterate with b2..bN (one change per iteration), each compared via benchstat
- [x] Write day10.adoc with:
  - b0 baseline summary
  - b1 (or bN) benchstat delta summary
  - concise optimization notes
  - raw benchstat block
- [x] Keep README untouched (no include) unless explicitly requested
- [x] Verify day tests and benchmark command still pass

Scope: Day10 only.
Execution model: nonblocking, independent task for parallel agents.

## Summary of Changes

- Existing step-1 optimization remains the best Day10 implementation
- Attempted step-2 buffer-capacity synchronization and measured `b2`
- Confirmed regression versus `b1` (`+21.55%` part1 time, `+3.42%` part2 time; higher allocs/B/op)
- Reverted Day10 code to prior step-1 implementation and re-verified tests
