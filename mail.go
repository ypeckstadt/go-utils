package utils

import (
	"net/smtp"
	"strconv"
)

type SendMailRequest struct {
	// From email address we are sending from
	From     string

	// To list email addresses we are sending email to
	To       []string

	// Host of SMTP server
	Host     string

	// Port of SMTP server
	Port     int

	Message  string

	// Password of email address we are sending from
	Password string
}

// SendMail sends an email via SMTP server
func SendMail(request SendMailRequest) error {
	auth := smtp.PlainAuth("", request.From, request.Password, request.Host)
	if err := smtp.SendMail(request.Host +":" + strconv.Itoa(request.Port), auth, request.From, request.To, []byte(request.Message)); err != nil {
		return err
	}
	return nil
}
