package main

import (
	"bytes"
	"testing"
	"text/template"
)

// technique 31
// can be run `go test -bench .`
// can also be run `go test -bench . -cpu=1,2,4` to test
// using a different number of CPUs

// goroutines in b.RunParallel() will attempt to wrtie
// using the same bytes.Buffer
// Using the -race flag will help to identify the error
// e.g. `go test -bench Oops -race -cpu=1,2,4`
func BenchmarkParallelOops(b *testing.B) {
	tpl := "Hello {{.Name}}"
	t, _ := template.New("test").Parse(tpl)
	data := &map[string]string{
		"Name": "World",
	}
	var buf bytes.Buffer
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			t.Execute(&buf, data)
			buf.Reset()
		}
	})
}
