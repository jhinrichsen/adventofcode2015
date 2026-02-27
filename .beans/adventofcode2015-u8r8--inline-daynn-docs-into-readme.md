---
# adventofcode2015-u8r8
title: Inline dayNN docs into README
status: completed
type: task
priority: normal
created_at: 2026-02-27T14:17:20Z
updated_at: 2026-02-27T14:19:17Z
---

Inline all `dayNN.adoc` documents into `README.adoc` sequentially.

## Plan
- [x] Find all dayNN.adoc files
- [x] Merge contents into README.adoc in order
- [x] Verify README formatting
- [x] Summarize changes

## Summary of Changes

- Inlined all `dayNN.adoc` files into `README.adoc` in ascending order (`day02` through `day25`).
- Preserved the existing README intro/benchmark content before day sections.
- Preserved the existing `== SAST (Static Application Security Testing)` section after the day block.
- Fixed an intermediate duplicated day block so each day section appears exactly once in sequence.
