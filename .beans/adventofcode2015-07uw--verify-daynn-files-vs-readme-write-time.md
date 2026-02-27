---
# adventofcode2015-07uw
title: Verify dayNN files vs README write time
status: completed
type: task
priority: normal
created_at: 2026-02-27T14:24:22Z
updated_at: 2026-02-27T14:24:59Z
---

Check whether any `dayNN.adoc` file has been modified after the latest write time of `README.adoc`.

## Plan
- [x] Capture README timestamp
- [x] Compare all dayNN timestamps
- [x] Report any newer day files

## Summary of Changes

- Compared `README.adoc` modification time against all `dayNN.adoc` files.
- Found one file newer than README: `day06.adoc` (README mtime: 2026-02-27 15:18:39 +0100, day06 mtime: 2026-02-27 15:19:13 +0100).
- All other `dayNN.adoc` files are not newer than README.
