module github.com/go_examples/modules/hello

go 1.23.6

replace github.com/go_examples/modules/greetings => ../greetings

require github.com/go_examples/modules/greetings v0.0.0-00010101000000-000000000000
