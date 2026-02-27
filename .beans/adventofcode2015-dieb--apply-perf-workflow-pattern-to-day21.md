---
# adventofcode2015-dieb
title: Apply perf workflow pattern to Day21
status: completed
type: task
priority: normal
created_at: 2026-02-27T11:53:19Z
updated_at: 2026-02-27T12:47:25Z
---

Performance workflow pattern:

- [x] Capture b0 baseline benchmark in benches/day21-b0.txt (trim PASS/ok lines)
- [x] Apply one targeted performance optimization
- [x] Capture b1 benchmark in benches/day21-b1.txt (trim PASS/ok lines)
- [x] Run benchstat b0 vs b1
- [x] Optionally iterate with b2..bN (one change per iteration), each compared via benchstat (not needed after b1)
- [x] Write day21.adoc with:
  - b0 baseline summary
  - b1 (or bN) benchstat delta summary
  - concise optimization notes
  - raw benchstat block
- [x] Keep README untouched (no include) unless explicitly requested
- [x] Verify day tests and benchmark command still pass

Scope: Day21 only.
Execution model: nonblocking, independent task for parallel agents.

## Summary of Changes

- Captured baseline benchmark in `benches/day21-b0.txt`.
- Applied one targeted optimization in `day21.go`:
  - removed gear-combination materialization and evaluated combinations directly
  - replaced turn-by-turn combat simulation with arithmetic turn-count comparison
- Captured post-change benchmark in `benches/day21-b1.txt` and compared with `benchstat` in `benches/day21-benchstat.txt`.
- Added `day21.adoc` with baseline summary, delta summary, optimization notes, and raw benchstat block.
- Verified correctness and benchmark command with:
  - `go test day21.go day21_test.go pattern_test.go input_test.go`
  - `go test -run=^$ -bench=Day21Part.$ -count=1 -benchmem day21.go day21_test.go pattern_test.go input_test.go`
