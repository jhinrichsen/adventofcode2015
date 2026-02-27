---
# adventofcode2015-af5w
title: Apply perf workflow pattern to Day12
status: completed
type: task
priority: normal
created_at: 2026-02-27T11:53:19Z
updated_at: 2026-02-27T12:07:36Z
---

Performance workflow pattern:

- [x] Capture b0 baseline benchmark in benches/day12-b0.txt (trim PASS/ok lines)
- [x] Apply one targeted performance optimization
- [x] Capture b1 benchmark in benches/day12-b1.txt (trim PASS/ok lines)
- [x] Run benchstat b0 vs b1
- [x] Optionally iterate with b2..bN (one change per iteration), each compared via benchstat
- [x] Write day12.adoc with:
  - b0 baseline summary
  - b1 (or bN) benchstat delta summary
  - concise optimization notes
  - raw benchstat block
- [x] Keep README untouched (no include) unless explicitly requested
- [x] Verify day tests and benchmark command still pass

Scope: Day12 only.
Execution model: nonblocking, independent task for parallel agents.

## Summary of Changes

- Captured `benches/day12-b0.txt` and `benches/day12-b1.txt`, then iterated once more to `benches/day12-b2.txt`
- Replaced Day12 part 2 generic JSON decode traversal with an iterative byte parser
- Added `benches/day12-benchstat-b0-b2.txt` and `benches/day12-benchstat-b1-b2.txt`
- Added `day12.adoc` with baseline, delta summary, optimization notes, and raw benchstat block
- Verified `go test -run ^TestDay12 ./...` and `go test -run=^$ -bench=Day12Part.$ -count=1`
