package main

import (
    "fmt"
    "log"
    "net/smtp"
    "strings"
)

const CONFIG_SMTP_HOST = "smtp.gmail.com"
const CONFIG_SMTP_PORT = 587
const CONFIG_SENDER_NAME = "Okta gilang Aljaffarsyah <aljaffarsyah2@gmail.com>"
const CONFIG_AUTH_EMAIL = "aljaffarsyah2@gmail.com"
const CONFIG_AUTH_PASSWORD = "pybpcsvrjxgwoowc"

func main() {
    to := []string{"aljaffarsyah10@gmail.com", "aljaffarsyah1@gmail.com"}
    cc := []string{"aljaffarsyah10@gmail.com"}
    subject := "Test mail"
    message := "Hello"

    err := sendMail(to, cc, subject, message)
    if err != nil {
        log.Fatal(err.Error())
    }

    log.Println("Mail sent!")
}

func sendMail(to []string, cc []string, subject, message string) error {
    body := "From: " + CONFIG_SENDER_NAME + "\n" +
        "To: " + strings.Join(to, ",") + "\n" +
        "Cc: " + strings.Join(cc, ",") + "\n" +
        "Subject: " + subject + "\n\n" +
        message

    auth := smtp.PlainAuth("", CONFIG_AUTH_EMAIL, CONFIG_AUTH_PASSWORD, CONFIG_SMTP_HOST)
    smtpAddr := fmt.Sprintf("%s:%d", CONFIG_SMTP_HOST, CONFIG_SMTP_PORT)

    err := smtp.SendMail(smtpAddr, auth, CONFIG_AUTH_EMAIL, append(to, cc...), []byte(body))
    if err != nil {
        return err
    }

    return nil
}