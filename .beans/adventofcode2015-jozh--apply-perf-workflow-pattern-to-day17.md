---
# adventofcode2015-jozh
title: Apply perf workflow pattern to Day17
status: completed
type: task
priority: normal
created_at: 2026-02-27T11:53:19Z
updated_at: 2026-02-27T12:43:36Z
---

Performance workflow pattern:

- [x] Capture b0 baseline benchmark in benches/day17-b0.txt (trim PASS/ok lines)
- [x] Apply one targeted performance optimization
- [x] Capture b1 benchmark in benches/day17-b1.txt (trim PASS/ok lines)
- [x] Run benchstat b0 vs b1
- [x] Optionally iterate with b2..bN (one change per iteration), each compared via benchstat
- [x] Write day17.adoc with:
  - b0 baseline summary
  - b1 (or bN) benchstat delta summary
  - concise optimization notes
  - raw benchstat block
- [x] Keep README untouched (no include) unless explicitly requested
- [x] Verify day tests and benchmark command still pass

Scope: Day17 only.
Execution model: nonblocking, independent task for parallel agents.

## Summary of Changes

- Captured `benches/day17-b0.txt` and `benches/day17-b1.txt`
- Ran `benchstat` and saved output to `benches/day17-benchstat.txt`
- Added `day17.adoc` with benchmark workflow notes and raw benchstat block
- Verified `go test -run ^TestDay17 ./...` and `go test -run=^$ -bench=Day17Part.$ -count=1`
