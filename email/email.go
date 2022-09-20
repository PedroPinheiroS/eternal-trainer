package email

import (
	"errors"
	"io/ioutil"
	"log"

	"github.com/PedroPinheiroS/eternal-trainer/env"
	"github.com/PedroPinheiroS/eternal-trainer/utils"
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"github.com/emersion/go-message/mail"
)

func connect() (*client.Client, error) {
	// Connect to server
	c, err := client.DialTLS("imap.mail.yahoo.com:993", nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected")

	return c, err
}

func RetrieveToken() (string, error) {
	log.Println("Connecting to server...")

	c, err := connect()
	if err != nil {
		log.Println(err)
		return "", err
	}

	// Don't forget to logout
	defer c.Logout()

	// Login
	err = c.Login(env.Email, env.Password)
	if err != nil {
		log.Println(err)
	}

	log.Println("Logged in")

	mbox, err := c.Select("INBOX", false)
	if err != nil {
		log.Println(err)
	}

	// Get the last message
	if mbox.Messages == 0 {
		log.Println("No message in mailbox")
		return "", errors.New("no message in mailbox")
	}
	seqSet := new(imap.SeqSet)
	seqSet.AddNum(mbox.Messages)

	// Get the whole message body
	var section imap.BodySectionName
	items := []imap.FetchItem{section.FetchItem()}

	messages := make(chan *imap.Message, 1)
	go func() {
		if err := c.Fetch(seqSet, items, messages); err != nil {
			log.Fatal(err)
		}
	}()

	msg := <-messages
	if msg == nil {
		log.Println("Server didn't returned message")
		return "", errors.New("server didn't returned message")
	}

	r := msg.GetBody(&section)
	if r == nil {
		log.Println("Server didn't returned message body")
		return "", errors.New("server didn't returned message body")
	}

	// Create a new mail reader
	mr, err := mail.CreateReader(r)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	p, _ := mr.NextPart()

	// This is the message's text (can be plain-text or HTML)
	b, _ := ioutil.ReadAll(p.Body)

	body := string(b)

	token := utils.GetToken(body)

	tokenClean := utils.CleanToken(token)

	//trimbString := strings.Trim("", bString)
	return tokenClean, nil
}
