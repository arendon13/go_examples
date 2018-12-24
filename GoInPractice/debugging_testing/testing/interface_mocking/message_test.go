package msg

import (
	"testing"
)

// technique 27

// MockMessage mocks the message struct from message.go
type MockMessage struct {
	email, subject string
	body           []byte
}

// Send is mocking the function that will be used to comply
// with the Messager Interface
func (m *MockMessage) Send(email, subject string, body []byte) error {
	m.email = email
	m.subject = subject
	m.body = body
	return nil
}

func TestAlert(t *testing.T) {
	msgr := new(MockMessage)
	body := []byte("Critical Error")

	Alert(msgr, body)

	if msgr.subject != "Critical Error" {
		t.Errorf("Expected 'Critical Error', Got '%s'", msgr.subject)
	}
	// continue on to verify results
}
