package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	data, err := os.ReadFile("/proc/cpuinfo")
	if err != nil {
		fmt.Print("unknown")
		return
	}

	re := regexp.MustCompile(`model name\s*:\s*(.+)`)
	matches := re.FindStringSubmatch(string(data))
	if len(matches) < 2 {
		fmt.Print("unknown")
		return
	}

	name := matches[1]

	// Remove CPU speed (e.g., "@ 2.4GHz")
	name = regexp.MustCompile(`@.*`).ReplaceAllString(name, "")

	// Remove special characters
	name = regexp.MustCompile(`[()/@]`).ReplaceAllString(name, "")

	// Convert spaces to underscores
	name = strings.ReplaceAll(name, " ", "_")

	// Collapse multiple consecutive underscores
	name = regexp.MustCompile(`_+`).ReplaceAllString(name, "_")

	// Remove trailing underscores
	name = strings.TrimRight(name, "_")

	// Convert to lowercase
	name = strings.ToLower(name)

	fmt.Print(name)
}
