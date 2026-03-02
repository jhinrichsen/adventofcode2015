---
# adventofcode2015-rbzr
title: Audit Day22 bean consistency
status: completed
type: task
priority: normal
created_at: 2026-02-28T13:29:44Z
updated_at: 2026-02-28T13:29:49Z
---

Review Day22-related beans for inconsistencies or misleading verification claims.

## Summary of Changes

- Reviewed all Day22-related beans and cross-checked verification claims.
- Found one misleading validation command in perf bean (`go test -run ^TestDay22$`) that runs zero tests.
- Confirmed runtime increase is explained by separate change removing `testing.Short()` skips for Day22/Day24 tests.
