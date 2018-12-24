package main

import (
	"bytes"
	"testing"
	"text/template"
)

// technique 29

// BechmarkTemplates is designed to zero in on the average
// time it takes to compile and run a simple text template
func BenchmarkTemplates(b *testing.B) {
	b.Logf("b.N is %d\n", b.N)
	tpl := "Hello {{.Name}}"
	data := &map[string]string{
		"Name": "World",
	}
	var buf bytes.Buffer
	for i := 0; i < b.N; i++ {
		t, _ := template.New("test").Parse(tpl)
		t.Execute(&buf, data)
		buf.Reset()
	}
}

// Running same test only this time we are parsing the
// template outside of the for loop
func BenchmarkCompiledTemplates(b *testing.B) {
	b.Logf("b.N is %d\n", b.N)
	tpl := "Hello {{.Name}}"
	t, _ := template.New("test").Parse(tpl)
	data := &map[string]string{
		"Name": "World",
	}
	var buf bytes.Buffer
	for i := 0; i < b.N; i++ {
		t.Execute(&buf, data)
		buf.Reset()
	}
}
