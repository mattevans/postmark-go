package main

import (
	"fmt"
	"net/http"

	"github.com/mattevans/postmark-go"
)

func main() {
	// Authenticate.
	auth := &http.Client{
		Transport: &postmark.AuthTransport{Token: "SERVER_API_TOKEN"},
	}
	client := postmark.NewClient(auth)

	// Build the email.
	emailReq := &postmark.Email{
		From:       "mail@company.com",
		To:         "jack@sparrow.com",
		Subject:    "My Test Email",
		HTMLBody:   "<html><body><strong>Hello</strong> dear Postmark user.</body></html>",
		TextBody:   "Hello dear Postmark user",
		Tag:        "onboarding",
		TrackOpens: true,
		Metadata: map[string]string{
			"client-id": "123456",
			"client-ip": "127.0.0.1",
		},
	}

	// Send it!
	_, response, err := client.Email.Send(emailReq)
	if err != nil {
		fmt.Printf("ERR: \n%v\n%v\n", response, err)
	}
}
