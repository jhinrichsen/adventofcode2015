---
# adventofcode2015-rzet
title: Apply perf workflow pattern to Day14
status: completed
type: task
priority: normal
created_at: 2026-02-27T11:53:19Z
updated_at: 2026-02-27T12:51:17Z
---

Performance workflow pattern:

- [x] Capture b0 baseline benchmark in benches/day14-b0.txt (trim PASS/ok lines)
- [x] Apply one targeted performance optimization
- [x] Capture b1 benchmark in benches/day14-b1.txt (trim PASS/ok lines)
- [x] Run benchstat b0 vs b1
- [x] Optionally iterate with b2..bN (one change per iteration), each compared via benchstat
- [x] Write day14.adoc with:
  - b0 baseline summary
  - b1 (or bN) benchstat delta summary
  - concise optimization notes
  - raw benchstat block
- [x] Keep README untouched (no include) unless explicitly requested
- [x] Verify day tests and benchmark command still pass

Scope: Day14 only.
Execution model: nonblocking, independent task for parallel agents.

## Summary of Changes

- Verified Day14 tests and benchmark command execution
- Generated missing `benches/day14-benchstat.txt` from `day14-b0.txt` and `day14-b1.txt`
- Refreshed `day14.adoc` to include baseline/follow-up references and raw benchstat block
- Verified `go test -run ^TestDay14 ./...` and `go test -run=^$ -bench=Day14Part.$ -count=1`
