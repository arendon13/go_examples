package mystringutil

import (
	"log"
	"strings"
)

// Pad will truncate any string that is greater than a max num
func Pad(s string, max uint) string {
	log.Printf("Testing Len: %d, Str: %s\n", max, s)
	ln := uint(len(s))
	if ln > max {
		return s[:max-1] // should be s[:max] not s[:max-1]
	}
	s += strings.Repeat(" ", int(max-ln))
	return s
}
