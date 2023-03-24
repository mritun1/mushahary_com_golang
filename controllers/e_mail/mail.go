package e_mail

import (
	"os"

	"gopkg.in/gomail.v2"
)

func SendMail(fromMail string, toMail string, sub string, mgs string) (string, error) {
	// Set up email message
	m := gomail.NewMessage()
	m.SetHeader("From", fromMail)
	m.SetHeader("To", toMail)
	m.SetHeader("Subject", sub)
	m.SetBody("text/plain", mgs)

	// Set up SMTP authentication information
	d := gomail.NewDialer("smtp.gmail.com", 587, os.Getenv("LOCAL_SMTP_EMAIL"), os.Getenv("LOCAL_SMTP_PASS"))

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		return "Failed", err
	}

	return "Email send Success, We will check your email, Thank you:)", nil

}
