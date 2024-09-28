package mailUtils

import (
    gomail "gopkg.in/mail.v2"
)

var smtpHost = "smtp.gmail.com"
var smtpPort = 587

type Creds struct {
    Email string
    Password string
}

type Message struct {
	Subject string
	Body string
}

func SendMessage(creds Creds, recipient string, message Message) error {
	m := gomail.NewMessage()
    m.SetHeader("From", creds.Email)
    m.SetHeader("To", recipient)
    m.SetHeader("Subject", message.Subject)
    m.SetBody("text/plain", message.Body)

    d := gomail.NewDialer(smtpHost, smtpPort, creds.Email, creds.Password)

    return d.DialAndSend(m)
}

func SendMultiple(creds Creds, recipients []string, message Message) {
	for _, r := range recipients {
		err := SendMessage(creds, r, message)
		if err != nil {
			panic(err)
		}
	}
}