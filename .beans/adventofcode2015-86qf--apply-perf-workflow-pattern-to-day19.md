---
# adventofcode2015-86qf
title: Apply perf workflow pattern to Day19
status: completed
type: task
priority: normal
created_at: 2026-02-27T11:53:19Z
updated_at: 2026-02-27T13:36:16Z
---

Performance workflow pattern:

- [x] Capture b0 baseline benchmark in benches/day19-b0.txt (trim PASS/ok lines)
- [x] Apply one targeted performance optimization
- [x] Capture b1 benchmark in benches/day19-b1.txt (trim PASS/ok lines)
- [x] Run benchstat b0 vs b1
- [x] Optionally iterate with b2..bN (one change per iteration), each compared via benchstat
- [x] Write day19.adoc with:
  - b0 baseline summary
  - b1 (or bN) benchstat delta summary
  - concise optimization notes
  - raw benchstat block
- [x] Keep README untouched (no include) unless explicitly requested
- [x] Verify day tests and benchmark command still pass

Scope: Day19 only.
Execution model: nonblocking, independent task for parallel agents.

## Summary of Changes

- Captured fresh `benches/day19-b0.txt` and `benches/day19-b1.txt`
- Reworked part 2 reduction to use deterministic randomized greedy attempts with iterative bounded-search fallback
- Removed no-op replacement generation in fallback search
- Updated `benches/day19-benchstat.txt` and `day19.adoc`
- Verified `go test -run ^TestDay19 ./...` and `go test -run=^$ -bench=Day19Part.$ -count=1`
