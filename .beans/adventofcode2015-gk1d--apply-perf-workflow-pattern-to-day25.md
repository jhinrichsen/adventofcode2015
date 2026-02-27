---
# adventofcode2015-gk1d
title: Apply perf workflow pattern to Day25
status: completed
type: task
priority: normal
created_at: 2026-02-27T11:53:19Z
updated_at: 2026-02-27T12:11:37Z
---

Performance workflow pattern:

- [x] Capture b0 baseline benchmark in benches/day25-b0.txt (trim PASS/ok lines)
- [x] Apply one targeted performance optimization
- [x] Capture b1 benchmark in benches/day25-b1.txt (trim PASS/ok lines)
- [x] Run benchstat b0 vs b1
- [x] Optionally iterate with b2..bN (one change per iteration), each compared via benchstat
- [x] Write day25.adoc with:
  - b0 baseline summary
  - b1 (or bN) benchstat delta summary
  - concise optimization notes
  - raw benchstat block
- [x] Keep README untouched (no include) unless explicitly requested
- [x] Verify day tests and benchmark command still pass

Scope: Day25 only.
Execution model: nonblocking, independent task for parallel agents.

## Summary of Changes

- Captured `benches/day25-b0.txt` and `benches/day25-b1.txt`
- Replaced Day25 coordinate stepping loop with direct diagonal index math and iterative modular exponentiation
- Added `benches/day25-benchstat.txt` and `day25.adoc`
- Verified `go test -run ^TestDay25 ./...` and `go test -run=^$ -bench=Day25Part.$ -count=1`
