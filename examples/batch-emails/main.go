package main

import (
	"fmt"
	"net/http"

	"github.com/mattevans/postmark-go"
)

func main() {
	// Init client with round tripper adding auth fields.
	client := postmark.NewClient(
		postmark.WithClient(&http.Client{
			Transport: &postmark.AuthTransport{Token: "SERVER_API_TOKEN"},
		}),
	)

	// Slice of recievers
	receivers := []string{
		"jack@sparrow.com",
		"fiona@shrek.com",
	}

	// Build emails
	emailRequests := []*postmark.Email{}
	for _, receiver := range receivers {
		emailRequests = append(emailRequests, &postmark.Email{
			From:       "mail@company.com",
			To:         receiver,
			Subject:    "My Test Email",
			HTMLBody:   "<html><body><strong>Hello</strong> dear Postmark user.</body></html>",
			TextBody:   "Hello dear Postmark user",
			Tag:        "onboarding",
			TrackOpens: true,
		})
	}

	// Send them!
	_, response, err := client.Email.SendBatch(emailRequests)
	if err != nil {
		fmt.Printf("ERR: \n%v\n%v\n", response, err)
	}
}
