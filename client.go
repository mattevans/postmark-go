package postmark

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// The default backend for the postmark API.
const (
	PostmarkURI = `https://api.postmarkapp.com`
)

// The different types of connections available to the Postmark API.
const (
	PostmarkConnectionTypeAccount = `account`
	PostmarkConnectionTypeServer  = `server`
)

// Client holds a connection to the Postmark API.
type Client struct {
	HTTPClient     *http.Client
	Token          string
	ConnectionType string
	BackendURI     string
}

// NewPostmarkClient creates a new Client with the appropriate API token and the
// ability to override the backend.
func NewPostmarkClient(token, connectionType string, newBackendURI *string) *Client {
	// Are we setting a different backend?
	backendURI := PostmarkURI
	if newBackendURI != nil {
		backendURI = *newBackendURI
	}

	return &Client{
		HTTPClient:     &http.Client{},
		Token:          token,
		ConnectionType: connectionType,
		BackendURI:     backendURI,
	}
}

// Request builds and executes a request to the Postmark API marshalling the
// response body.
func (client *Client) Request(method, path string, payload, destination interface{}) error {
	// Avoid shadowing
	var err error

	// Build the request
	url := fmt.Sprintf("%s/%s", client.BackendURI, path)
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return err
	}

	// If we have a payload, write to the request body.
	if payload != nil {
		var payloadData []byte
		payloadData, err = json.Marshal(payload)
		if err != nil {
			return err
		}
		req.Body = ioutil.NopCloser(bytes.NewBuffer(payloadData))
	}

	// Append the appropriate headers to the request.
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	// Ensure we pass the correct header depending on the connectiont type.
	if client.ConnectionType == PostmarkConnectionTypeServer {
		req.Header.Add("X-Postmark-Server-Token", client.Token)
	}
	if client.ConnectionType == PostmarkConnectionTypeAccount {
		req.Header.Add("X-Postmark-Account-Token", client.Token)
	}

	// Execute the request.
	res, err := client.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	// Read the request.
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	// Unmarshal reaquest body into destination interface.
	err = json.Unmarshal(body, destination)
	return err
}
