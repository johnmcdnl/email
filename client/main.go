package main

import (
	"time"
	"github.com/emersion/go-imap"
	"log"
	"github.com/emersion/go-imap/client"
)

func main(){
	for i := 0; i <= 30; i++ {
		time.Sleep(1 * time.Second)
		ReadMail()
	}
}

func ReadMail(){
		log.Println("Connecting to server...")
		// Connect to server
		c, err := client.Dial("mailserver:143")
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Connected")

		// Don't forget to logout
		defer c.Logout()

		// Login
		if err := c.Login("johnmcdnl", "password"); err != nil {
			log.Fatal(err)
		}
		log.Println("Logged in")

		// List mailboxes
		mailboxes := make(chan *imap.MailboxInfo, 10)
		done := make(chan error, 1)
		go func () {
			done <- c.List("", "*", mailboxes)
		}()

		log.Println("Mailboxes:")
		for m := range mailboxes {
			log.Println("* " + m.Name)
		}

		if err := <-done; err != nil {
			log.Fatal(err)
		}

		// Select INBOX
		mbox, err := c.Select("INBOX", false)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Flags for INBOX:", mbox.Flags)

		// Get the last 4 messages
		from := uint32(1)
		to := mbox.Messages
		if mbox.Messages > 3 {
			// We're using unsigned integers here, only substract if the result is > 0
			from = mbox.Messages - 3
		}
		seqset := new(imap.SeqSet)
		seqset.AddRange(from, to)

		messages := make(chan *imap.Message, 10)
		done = make(chan error, 1)
		go func() {
			done <- c.Fetch(seqset, []string{imap.EnvelopeMsgAttr}, messages)
		}()

		log.Println("Last 4 messages:")
		for msg := range messages {
			log.Println("* " + msg.Envelope.Subject)
		}

		if err := <-done; err != nil {
			log.Fatal(err)
		}

		log.Println("Done!")

}


