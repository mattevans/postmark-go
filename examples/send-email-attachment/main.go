package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"

	postmark "github.com/mattevans/postmark-go"
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
	}

	// Build and append our attachment(s)
	emailReq.Attachments = []postmark.EmailAttachment{
		buildAttachment("./attachment_example.txt"),
	}

	// Send it!
	_, response, err := client.Email.Send(emailReq)
	if err != nil {
		fmt.Printf("ERR: \n%v\n%v\n", response, err)
	}
}

// buildAttachment receives a given path, reads the file and returns a
// postmark.EmailAttachment
func buildAttachment(path string) postmark.EmailAttachment {
	// Read our attachment
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Print(err)
	}

	// Determine the ContentType.
	contentType := http.DetectContentType(bytes)

	// Build and return our attachments
	return postmark.EmailAttachment{
		Name:        fmt.Sprintf("example%s", filepath.Ext(path)),
		Content:     bytes,
		ContentType: &contentType,
	}
}
