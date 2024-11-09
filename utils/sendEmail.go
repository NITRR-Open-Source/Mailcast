package utils

import (
	"fmt"
	"log"
	"net/smtp"
)

func SendEmail(subject string, body string, recipient string, sender string, password string) error {
	// Set up SMTP server configuration
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	auth := smtp.PlainAuth("", sender, password, smtpHost)

	// Prepare the email content
	msg := []byte("From: " + sender + "\r\n" +
		"To: " + recipient + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"MIME-Version: 1.0\r\n" +
		"Content-Type: text/html; charset=\"UTF-8\"\r\n" +
		"\r\n" +
		body + "\r\n")

	// Establish TLS-encrypted connection and send the email
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, sender, []string{recipient}, msg)
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	log.Println("Mail sent successfully to", recipient)
	return nil
}
