package mystringutil

import (
	"testing"
	"testing/quick"
)

// Simple pad unit test
func TestPad(t *testing.T) {
	if r := Pad("test", 6); len(r) != 6 {
		t.Errorf("Expected 6, got %d", len(r))
	}
}

// Generative test for pad
func TestPadGenerative(t *testing.T) {
	fn := func(s string, max uint8) bool {
		p := Pad(s, uint(max))
		return len(p) == int(max)
	}

	if err := quick.Check(fn, &quick.Config{MaxCount: 200}); err != nil {
		t.Error(err)
	}
}
