package postmark

import (
	"fmt"
	"time"
)

// EmailHeader holds values of a email header.
type EmailHeader struct {
	Name  string
	Value string
}

// Email is the set of parameters that can be used when sending an email.
type Email struct {
	From          string                 `json:",omitempty"`
	To            string                 `json:",omitempty"`
	Cc            string                 `json:",omitempty"`
	Bcc           string                 `json:",omitempty"`
	Subject       string                 `json:",omitempty"`
	Tag           string                 `json:",omitempty"`
	TemplateID    int                    `json:",omitempty"`
	TemplateModel map[string]interface{} `json:",omitempty"`
	HTMLBody      string                 `json:",omitempty"`
	TextBody      string                 `json:",omitempty"`
	ReplyTo       string                 `json:",omitempty"`
	Headers       []EmailHeader          `json:",omitempty"`
	TrackOpens    bool                   `json:",omitempty"`
}

// EmailResponse is the set of parameters that is used in response to a send
// request
type EmailResponse struct {
	To          string
	SubmittedAt time.Time
	MessageID   string
	ErrorCode   int64
	Message     string
}

// SendEmail will send the email via Postmark.
func (client *Client) SendEmail(email Email) (EmailResponse, error) {

	// Postmark API endpoints change if using a template for the email.
	endpoint := "email"
	if email.TemplateID != 0 {
		endpoint = "email/withTemplate"
	}

	response := EmailResponse{}
	err := client.Request("POST", endpoint, email, &response)

	// If our response code is not successful, handle appropriately.
	if response.ErrorCode != 0 {
		return response, fmt.Errorf(`%v %s`, response.ErrorCode, response.Message)
	}

	return response, err
}
