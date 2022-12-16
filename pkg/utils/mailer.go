package utils

import (
	"net/smtp"
	"os"
)

var (
	from = ""
	pass = ""
	host = ""
	port = ""
)

func setMailEnv() {
	from = os.Getenv("MAIL_FROM")
	pass = os.Getenv("MAIL_PWD")
	host = os.Getenv("MAIL_HOST")
	port = os.Getenv("MAIL_PORT")
}

func SendInfoMail(email, password string) error {
	setMailEnv()
	to := []string{email}

	msg := []byte("From: " + from + "\r\n" +
		"To: " + to[0] + "\r\n" +
		"Subject: Info\r\n\r\n" +
		"Hi," + "\r\n\r\n" +
		"You have an account now" + "\r\n\r\n" +
		"team superb")

	auth := smtp.PlainAuth("", from, pass, host)

	err := smtp.SendMail(host+":"+port, auth, from, to, msg)
	if err != nil {
		return err
	}
	return nil
}
