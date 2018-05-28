package main

import (
	"log"
	"fmt"
	"net/smtp"
	"time"
)

func main() {
	for i := 0; i <= 30; i++ {
		time.Sleep(1 * time.Second)
		SendEmail()
	}
}

func SendEmail() {
	// Connect to the remote SMTP server.
	c, err := smtp.Dial("mailserver:25")
	if err != nil {
		log.Fatal(err)
	}

	// Set the sender and recipient first
	if err := c.Mail("sender@app.com"); err != nil {
		log.Fatal(err)
	}
	if err := c.Rcpt("johnmcdnl@mailserver"); err != nil {
		log.Fatal(err)
	}

	// Send the email body.
	wc, err := c.Data()
	if err != nil {
		log.Fatal(err)
	}
	_, err = fmt.Fprintf(wc, "This is the email body and it defo worked")
	if err != nil {
		log.Fatal(err)
	}
	err = wc.Close()
	if err != nil {
		log.Fatal(err)
	}

	// Send the QUIT command and close the connection.
	err = c.Quit()
	if err != nil {
		log.Fatal(err)
	}
}
