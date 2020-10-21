package main

import (
	"fmt"
	"log"
	"net/smtp"
	"strings"
)

const CONFIG_SMTP_HOST = "smtp.gmail.com"
const CONFIG_SMTP_PORT = 587
const CONFIG_SENDER_NAME = "PT. Makmur Subur Jaya <emailanda@gmail.com>"

func sendMail(to []string, cc []string, subject, message string) error {
	body := "From: " + CONFIG_SENDER_NAME + "\n" +
		"To: " + strings.Join(to, ",") + "\n" +
		"Cc: " + strings.Join(cc, ",") + "\n" +
		"Subject: " + subject + "\n\n" +
		message

	smtpAddr := fmt.Sprintf("%s:%d", CONFIG_SMTP_HOST, CONFIG_SMTP_PORT)
	c, err := smtp.Dial(smtpAddr)
	if c != nil {
		defer c.Quit()
	}
	if err != nil {
		return err
	}
	defer c.Close()

	err = c.Mail(CONFIG_SENDER_NAME)
	if err != nil {
		return err
	}

	w, err := c.Data()
	if err != nil {
		return err
	}

	_, err = w.Write([]byte(body))
	if err != nil {
		return err
	}

	err = w.Close()
	if err != nil {
		return err
	}

	return nil
}

func main() {
	to := []string{"recipient1@gmail.com", "emaillain@gmail.com"}
	cc := []string{"tralalala@gmail.com"}
	subject := "Test mail"
	message := "Hello"

	err := sendMail(to, cc, subject, message)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Mail sent!")
}
