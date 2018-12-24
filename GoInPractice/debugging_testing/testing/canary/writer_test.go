package customWriter

import (
	"testing"
)

func TestWriter(t *testing.T) {
	// The following will throw a compiler error
	// var _ io.Writer = &MyWriter{}
}
