package msg

// technique 27

// Message struct contains elements that would be
// necessary in sending a message
type Message struct {
}

// Send will handle logic to send a message
func (m *Message) Send(email, subject string, body []byte) error {
	// Senf logic goes here...
	return nil
}

// Messager adds a layer of abstraction for sending message
type Messager interface {
	Send(email, subject string, body []byte) error
}

// Alert invokes any valid messager
func Alert(m Messager, problem []byte) error {
	return m.Send("noc@example.com", "Critical Error", problem)
}
