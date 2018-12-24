package customWriter

// MyWriter is a custom struct
type MyWriter struct {
	// ...
}

// Write is a function that should implement io.Writer.Send()
func (m *MyWriter) Write([]byte) error {
	// logic for MyWriter
	return nil
}
