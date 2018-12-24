package main

import (
	"bytes"
	"testing"
	"text/template"
)

// technique 30
// can be run `go test -bench .`
// can also be run `go test -bench . -cpu=1,2,4` to test
// using a different number of CPUs

// b.RunParallel() will execute the t.Execute() repeatedly
// but as separate goroutines
func BenchmarkParallelTemplates(b *testing.B) {
	tpl := "Hello {{.Name}}"
	t, _ := template.New("test").Parse(tpl)
	data := &map[string]string{
		"Name": "World",
	}
	b.RunParallel(func(pb *testing.PB) {
		var buf bytes.Buffer
		for pb.Next() {
			t.Execute(&buf, data)
			buf.Reset()
		}
	})
}
