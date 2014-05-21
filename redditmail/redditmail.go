package main

import (
	"github.com/mandolyte/redditnews"
	"log"
	"net/smtp"
)

func main() {
	to := "cecil.new@gmail.com"
	subject := "Go Articles on Reddit"
	message := redditnews.Email()
	body := "To: " + to + "\r\nSubject: " +
		subject + "\r\n\r\n" + message

	auth := smtp.PlainAuth("", "cecil.new",
		"nxmbxtezqraahtal",
		"smtp.gmail.com")

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"cecil.new@gmail.com",
		[]string{to},
		[]byte(body))

	if err != nil {
		log.Fatal("SendMail: ", err)
		return
	}
}
