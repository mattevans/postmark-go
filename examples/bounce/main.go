package main

import (
	"fmt"
	"net/http"

	postmark "github.com/mattevans/postmark-go"
)

func main() {
	// Authenticate.
	auth := &http.Client{
		Transport: &postmark.AuthTransport{Token: "SERVER_API_TOKEN"},
	}
	client := postmark.NewClient(auth)

	// Get delivery stats.
	stats, response, err := client.Bounce.GetDeliveryStats()
	if err != nil {
		fmt.Printf("ERR: \n%v\n%v\n", response, err)
	}

	// Output the results.
	fmt.Printf("Delivery Stats: \n%v\n\n", stats)

	// Get bounces
	params := map[string]interface{}{}
	bounces, response, err := client.Bounce.GetBounces(500, 0, params)
	if err != nil {
		fmt.Printf("ERR: \n%v\n%v\n", response, err)
	}

	// Output the results.
	fmt.Printf("Bounces: \n%v\n\n", bounces)
}
