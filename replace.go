package adventofcode2015

import "strings"

// ReplaceNth replaces the nth occurrence of old in s with new. Similar to
// Replace() in the standard library, but Replace() will replace n times, while
// ReplaceNth will replace the nth occurrence 1 time.
func ReplaceNth(s, old, new string, n int) string {
	i := 0
	for j := 1; j <= n; j++ {
		x := strings.Index(s[i:], old)
		if x < 0 {
			break
		}
		i += x
		if j == n {
			return s[:i] + new + s[i+len(old):]
		}
		i += len(old)
	}
	return s
}
