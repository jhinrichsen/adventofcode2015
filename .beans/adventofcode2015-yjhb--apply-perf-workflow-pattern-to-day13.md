---
# adventofcode2015-yjhb
title: Apply perf workflow pattern to Day13
status: completed
type: task
priority: normal
created_at: 2026-02-27T11:53:19Z
updated_at: 2026-02-27T13:22:12Z
---

Performance workflow pattern:

- [x] Capture b0 baseline benchmark in benches/day13-b0.txt (trim PASS/ok lines)
- [x] Apply one targeted performance optimization
- [x] Capture b1 benchmark in benches/day13-b1.txt (trim PASS/ok lines)
- [x] Run benchstat b0 vs b1
- [x] Optionally iterate with b2..bN (one change per iteration), each compared via benchstat
- [x] Write day13.adoc with:
  - b0 baseline summary
  - b1 (or bN) benchstat delta summary
  - concise optimization notes
  - raw benchstat block
- [x] Keep README untouched (no include) unless explicitly requested
- [x] Verify day tests and benchmark command still pass

Scope: Day13 only.
Execution model: nonblocking, independent task for parallel agents.

## Summary of Changes

- Captured fresh `benches/day13-b0.txt` and `benches/day13-b1.txt`
- Replaced channel-based permutation generation with iterative next-permutation and fixed-seat canonicalization
- Added `benches/day13-benchstat.txt`
- Updated `day13.adoc` with baseline, delta summary, optimization notes, and raw benchstat block
- Verified `go test -run ^TestDay13 ./...` and `go test -run=^$ -bench=Day13Part.$ -count=1`
