---
# adventofcode2015-mkvv
title: Apply perf workflow pattern to Day16
status: completed
type: task
priority: normal
created_at: 2026-02-27T11:53:19Z
updated_at: 2026-02-27T13:46:51Z
---

Performance workflow pattern:

- [x] Capture b0 baseline benchmark in benches/day16-b0.txt (trim PASS/ok lines)
- [x] Apply one targeted performance optimization
- [x] Capture b1 benchmark in benches/day16-b1.txt (trim PASS/ok lines)
- [x] Run benchstat b0 vs b1
- [x] Optionally iterate with b2..bN (one change per iteration), each compared via benchstat
- [x] Write day16.adoc with:
  - b0 baseline summary
  - b1 (or bN) benchstat delta summary
  - concise optimization notes
  - raw benchstat block
- [x] Keep README untouched (no include) unless explicitly requested
- [x] Verify day tests and benchmark command still pass

Scope: Day16 only.
Execution model: nonblocking, independent task for parallel agents.

## Summary of Changes

- Added Day16 step 2 benchmark result as `benches/day16-b2.txt`
- Reworked parser to single-pass scanning with manual numeric parsing and zero-allocation property dispatch
- Added `benches/day16-benchstat-b0-b2.txt` and updated `day16.adoc` with `b2` vs `b0`
- Verified `go test -run ^TestDay16 ./...` and `go test -run=^$ -bench=Day16Part.$ -count=1 -benchmem`
