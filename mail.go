package utils

import (
	"errors"
	"fmt"
	"net/smtp"
	"strconv"
	"strings"
)


type Mail struct {
	Name  string
	Email string
}

type SendMailRequest struct {
	// From email address and name we are sending from
	From Mail

	// To list email addresses we are sending email to
	To []Mail

	// Host of SMTP server
	Host string

	// Port of SMTP server
	Port int

	Message string

	// Password of email address we are sending from
	Password string

	// Subject of the email
	Subject string
}

// SendMail sends an email via SMTP server
func SendMail(request SendMailRequest) error {

	// check if mail to targets have been provided
	if len(request.To) == 0 {
		return errors.New("no mail to specified")
	}

	var sb strings.Builder
	var to []string

	for _, mailTo := range request.To {
		sb.WriteString(fmt.Sprintf("To: \"%v\" <%v>", mailTo.Name, mailTo.Email) + "\n")
		to = append(to, mailTo.Email)
	}

	sb.WriteString(fmt.Sprintf("From: \"%v\" <%v>", request.From.Name, request.From.Email) + "\n")
	sb.WriteString(fmt.Sprintf("Subject: %v", request.Subject) + "\n")
	sb.WriteString("\n")
	sb.WriteString(request.Message)

	auth := smtp.PlainAuth("", request.From.Email, request.Password, request.Host)
	if err := smtp.SendMail(request.Host+":"+strconv.Itoa(request.Port), auth, request.From.Email, to, []byte(sb.String())); err != nil {
		return err
	}
	return nil
}
