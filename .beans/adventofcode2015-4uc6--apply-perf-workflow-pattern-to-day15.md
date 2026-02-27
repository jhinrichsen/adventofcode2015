---
# adventofcode2015-4uc6
title: Apply perf workflow pattern to Day15
status: completed
type: task
priority: normal
created_at: 2026-02-27T11:53:19Z
updated_at: 2026-02-27T12:39:47Z
---

Performance workflow pattern:

- [x] Capture b0 baseline benchmark in benches/day15-b0.txt (trim PASS/ok lines)
- [x] Apply one targeted performance optimization
- [x] Capture b1 benchmark in benches/day15-b1.txt (trim PASS/ok lines)
- [x] Run benchstat b0 vs b1
- [x] Optionally iterate with b2..bN (one change per iteration), each compared via benchstat
- [x] Write day15.adoc with:
  - b0 baseline summary
  - b1 (or bN) benchstat delta summary
  - concise optimization notes
  - raw benchstat block
- [x] Keep README untouched (no include) unless explicitly requested
- [x] Verify day tests and benchmark command still pass

Scope: Day15 only.
Execution model: nonblocking, independent task for parallel agents.

## Summary of Changes

- Captured `benches/day15-b0.txt` and `benches/day15-b1.txt`
- Replaced channel-based composition generation with iterative in-process composition traversal
- Fused calorie and score evaluation into a single pass
- Added `benches/day15-benchstat.txt` and `day15.adoc`
- Verified `go test -run ^TestDay15 ./...` and `go test -run=^$ -bench=Day15Part.$ -count=1`
