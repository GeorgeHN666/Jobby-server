package main

import (
	"fmt"

	gomail "gopkg.in/gomail.v2"
)

func (app *application) SendEmail(email, code, reason, supportEmail string, expiry int) error {

	mail := gomail.NewMessage()

	template := ``

	mail.SetHeader("From", "jobby@enterprices.com")
	mail.SetHeader("To", email)
	mail.SetHeader("Subject", fmt.Sprintf("Your %s code is: %s", reason, code))
	mail.SetBody("text/html", template)

	send := gomail.NewDialer("smtp.mailtrap.io", 587, "3d29a653402464", "4cbe5b99763b69")
	err := send.DialAndSend(mail)
	if err != nil {
		return err
	}

	return nil
}
