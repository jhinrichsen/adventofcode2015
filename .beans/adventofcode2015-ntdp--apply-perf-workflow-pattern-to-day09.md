---
# adventofcode2015-ntdp
title: Apply perf workflow pattern to Day09
status: completed
type: task
priority: normal
created_at: 2026-02-27T11:53:19Z
updated_at: 2026-02-27T12:12:34Z
---

Performance workflow pattern:

- [x] Capture b0 baseline benchmark in benches/day09-b0.txt (trim PASS/ok lines)
- [x] Apply one targeted performance optimization
- [x] Capture b1 benchmark in benches/day09-b1.txt (trim PASS/ok lines)
- [x] Run benchstat b0 vs b1
- [x] Optionally iterate with b2..bN (one change per iteration), each compared via benchstat
- [x] Write day09.adoc with:
  - b0 baseline summary
  - b1 (or bN) benchstat delta summary
  - concise optimization notes
  - raw benchstat block
- [x] Keep README untouched (no include) unless explicitly requested
- [x] Verify day tests and benchmark command still pass

Scope: Day09 only.
Execution model: nonblocking, independent task for parallel agents.

## Summary of Changes

- Processed Day09 performance-workflow task and marked checklist items complete.
- Recorded completion for queue progression under parallel-agent execution.


## Summary of Changes

- Captured baseline benchmark in `benches/day09-b0.txt`.
- Optimized Day09 by replacing channel/copy-based permutation enumeration with in-place iterative Heap traversal and direct evaluation.
- Captured post-change benchmark in `benches/day09-b1.txt` and compared with benchstat.
- Added `day09.adoc` with b0 baseline summary, b1 benchstat delta summary, optimization notes, and raw benchstat block.
- Verified Day09 tests and benchmark command still pass.
