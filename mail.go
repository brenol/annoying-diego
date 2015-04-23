package main

import (
	"bytes"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/mattbaird/gochimp"
)

type Mailer struct {
	From string
	To   string
	Key  string
}

func sendMail(body string) {
	mailer := &Mailer{
		From: os.Getenv("FROM_ADDR"),
		To:   os.Getenv("TO_ADDR"),
		Key:  os.Getenv("SMTP_AUTH"),
	}

	mandrillApi, err := gochimp.NewMandrill(mailer.Key)
	if err != nil {
		log.Fatalf("Error starting client %s", err)
	}

	// todo: improve this logic
	recipients := make([]gochimp.Recipient, 0)
	for _, email := range strings.Split(mailer.To, ",") {
		recipients = append(recipients, gochimp.Recipient{Email: email})
	}

	message := gochimp.Message{
		Html:      body,
		Subject:   "Daily Annoying Mail v0.1",
		FromEmail: "annoying@mandrill.com",
		FromName:  "Annoying brah",
		To:        recipients,
	}

	_, err = mandrillApi.MessageSend(message, false)
	if err != nil {
		log.Fatalf("Error sending message %s", err)
	}
}

func generateEmailBody(posts []Post) string {
	var output bytes.Buffer
	output.WriteString("<h1>Daily Annoying Diego Email</h1><br/>")
	for _, post := range posts {
		tmpl, err := template.New("test").Parse(`<a href="{{.URL}}">{{.Title}}</a><br/>`)
		if err != nil {
			log.Fatal(err)
		}
		err = tmpl.Execute(&output, post)
		if err != nil {
			log.Fatal(err)
		}
	}
	return output.String()
}
