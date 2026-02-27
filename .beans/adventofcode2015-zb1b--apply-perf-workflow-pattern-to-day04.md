---
# adventofcode2015-zb1b
title: Apply perf workflow pattern to Day04
status: completed
type: task
priority: normal
created_at: 2026-02-27T11:53:19Z
updated_at: 2026-02-27T12:00:57Z
---

Performance workflow pattern:

- [x] Capture b0 baseline benchmark in benches/day04-b0.txt (trim PASS/ok lines)
- [x] Apply one targeted performance optimization
- [x] Capture b1 benchmark in benches/day04-b1.txt (trim PASS/ok lines)
- [x] Run benchstat b0 vs b1
- [x] Optionally iterate with b2..bN (one change per iteration), each compared via benchstat (not needed after first significant gain)
- [x] Write day04.adoc with:
  - b0 baseline summary
  - b1 (or bN) benchstat delta summary
  - concise optimization notes
  - raw benchstat block
- [x] Keep README untouched (no include) unless explicitly requested
- [x] Verify day tests and benchmark command still pass

Scope: Day04 only.
Execution model: nonblocking, independent task for parallel agents.

## Summary of Changes

- Captured `benches/day04-b0.txt` and `benches/day04-b1.txt` using the standard benchmark workflow and trimmed trailing PASS/ok lines.
- Optimized `day04.go` by replacing per-iteration string formatting and string-prefix checks with reusable byte-buffer construction and direct MD5 leading-byte checks.
- Ran `benchstat benches/day04-b0.txt benches/day04-b1.txt` and saved output to `benches/day04-benchstat.txt`.
- Added `day04.adoc` with baseline summary, delta summary, optimization notes, and raw benchstat block.
- Verified `go test -run '^TestDay04' ./...` and `go test -run=^$ -bench=Day04Part.$ -count=1 -benchmem` both pass.
