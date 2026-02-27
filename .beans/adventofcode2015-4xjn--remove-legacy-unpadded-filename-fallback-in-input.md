---
# adventofcode2015-4xjn
title: Remove legacy unpadded filename fallback in input_test helpers
status: todo
type: task
priority: normal
created_at: 2026-02-27T09:26:05Z
updated_at: 2026-02-27T09:26:05Z
---

Drop fallback logic in input_test.go filename helpers and use strict zero-padded day file naming only (dayXX/dayXX_example...). Update remaining testdata filenames and verify tests.
