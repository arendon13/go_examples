package main

import (
	"bytes"
	"net/smtp"
	"strconv"
	"text/template"
)

// EmailMessage will organize data for an email message
type EmailMessage struct {
	From, Subject, Body string
	To                  []string
}

// EmailCredentials will organize data for email credentials
type EmailCredentials struct {
	Username, Password, Server string
	Port                       int
}

const emailTemplate = `From: {{.From}}
To: {{.To}}
Subject: {{.Subject}}

{{.Body}}
`

var t *template.Template

func init() {
	t = template.New("Email")
	t.Parse(emailTemplate)
}

func main() {
	message := &EmailMessage{
		From:    "me@example.com",
		To:      []string{"you@example.com"},
		Subject: "A test",
		Body:    "Just saying hi",
	}

	var body bytes.Buffer
	t.Execute(&body, message)

	authCreds := &EmailCredentials{
		Username: "myUsername",
		Password: "myPass",
		Server:   "smtp.example.com",
		Port:     25,
	}

	auth := smtp.PlainAuth("",
		authCreds.Username,
		authCreds.Password,
		authCreds.Server,
	)

	smtp.SendMail(authCreds.Server+":"+strconv.Itoa(authCreds.Port),
		auth,
		message.From,
		message.To,
		body.Bytes())

}
