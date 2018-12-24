package main

import (
	"fmt"
)

// ParseError created as test struct
type ParseError struct {
	Message    string
	Line, Char int
}

// technique 17
func main() {

}

// Go error type has one Error() that can easily be overwritten
func (p *ParseError) Error() string {
	format := "%s on Line %d, Char %d"
	return fmt.Sprintf(format, p.Message, p.Line, p.Char)
}
