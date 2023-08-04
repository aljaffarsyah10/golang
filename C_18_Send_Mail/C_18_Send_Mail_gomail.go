package main

import (
    "gopkg.in/gomail.v2"
    "log"
)

const CONFIG_SMTP_HOST = "smtp.gmail.com"
const CONFIG_SMTP_PORT = 587
const CONFIG_SENDER_NAME = "Okta gilang Aljaffarsyah <aljaffarsyah2@gmail.com>"
const CONFIG_AUTH_EMAIL = "aljaffarsyah2@gmail.com"
const CONFIG_AUTH_PASSWORD = "pybpcsvrjxgwoowc"

func main() {
    mailer := gomail.NewMessage()
    mailer.SetHeader("From", CONFIG_SENDER_NAME)
    mailer.SetHeader("To", "aljaffarsyah10@gmail.com", "aljaffarsyah1@gmail.com")
    mailer.SetAddressHeader("Cc", "oktagilang12@gmail.com", "Testing")
    mailer.SetHeader("Subject", "Test mail")
    mailer.SetBody("text/html", "Hello, <b>have a nice day</b>")
    mailer.Attach("./sample.png")

    dialer := gomail.NewDialer(
        CONFIG_SMTP_HOST,
        CONFIG_SMTP_PORT,
        CONFIG_AUTH_EMAIL,
        CONFIG_AUTH_PASSWORD,
    )

	// Tanpa Otentikasi (sekarang harus pakai otentikasi, jadi gabisa lagi)
	// dialer := &gomail.Dialer{Host: CONFIG_SMTP_HOST, Port: CONFIG_SMTP_PORT}

    err := dialer.DialAndSend(mailer)
    if err != nil {
        log.Fatal(err.Error())
    }

    log.Println("Mail sent!")
}