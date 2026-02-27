---
# adventofcode2015-w0v5
title: Apply perf workflow pattern to Day23
status: completed
type: task
priority: normal
created_at: 2026-02-27T11:53:19Z
updated_at: 2026-02-27T14:13:18Z
---

Performance workflow pattern:

- [x] Capture b0 baseline benchmark in benches/day23-b0.txt (trim PASS/ok lines)
- [x] Apply one targeted performance optimization
- [x] Capture b1 benchmark in benches/day23-b1.txt (trim PASS/ok lines)
- [x] Run benchstat b0 vs b1
- [x] Optionally iterate with b2..bN (one change per iteration), each compared via benchstat
- [x] Write day23.adoc with:
  - b0 baseline summary
  - b1 (or bN) benchstat delta summary
  - concise optimization notes
  - raw benchstat block
- [x] Keep README untouched (no include) unless explicitly requested
- [x] Verify day tests and benchmark command still pass

Scope: Day23 only.
Execution model: nonblocking, independent task for parallel agents.

## Summary of Changes

- Added Day23 step 3 benchmark result as `benches/day23-b3.txt`
- Replaced string-interpreted execution with parsed instruction structs and direct opcode dispatch
- Added `benches/day23-benchstat-b0-b3.txt` and updated `day23.adoc` with `b3` vs `b0`
- Verified `go test -run ^TestDay23 ./...` and `go test -run=^$ -bench=Day23Part.$ -count=1 -benchmem`
