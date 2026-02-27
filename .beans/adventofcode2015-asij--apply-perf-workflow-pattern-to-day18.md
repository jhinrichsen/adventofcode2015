---
# adventofcode2015-asij
title: Apply perf workflow pattern to Day18
status: completed
type: task
priority: normal
created_at: 2026-02-27T11:53:19Z
updated_at: 2026-02-27T14:16:53Z
---

Performance workflow pattern:

- [x] Capture b0 baseline benchmark in benches/day18-b0.txt (trim PASS/ok lines)
- [x] Apply one targeted performance optimization
- [x] Capture b1 benchmark in benches/day18-b1.txt (trim PASS/ok lines)
- [x] Run benchstat b0 vs b1
- [x] Optionally iterate with b2..bN (one change per iteration), each compared via benchstat
- [x] Write day18.adoc with:
  - b0 baseline summary
  - b1 (or bN) benchstat delta summary
  - concise optimization notes
  - raw benchstat block
- [x] Keep README untouched (no include) unless explicitly requested
- [x] Verify day tests and benchmark command still pass

Scope: Day18 only.
Execution model: nonblocking, independent task for parallel agents.

## Summary of Changes

- Added Day18 step 2 benchmark result as `benches/day18-b2.txt`
- Reworked simulation to use preallocated double buffers with step-time swapping
- Inlined neighbor counting inside the update loop
- Added `benches/day18-benchstat-b0-b2.txt` and updated `day18.adoc` with `b2` vs `b0`
- Verified `go test -run ^TestDay18 ./...` and `go test -run=^$ -bench=Day18Part.$ -count=1 -benchmem`
