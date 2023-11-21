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

	// Get delivery stats.
	stats, response, err := client.Bounce.GetDeliveryStats()
	if err != nil {
		fmt.Printf("ERR: \n%v\n%v\n", response, err)
	}

	// Output the results.
	fmt.Printf("Delivery Stats: \n%v\n\n", stats)

	// Get bounces (with filters)
	params := map[string]interface{}{
		"type":     "HardBounce",
		"fromdate": "2015-01-01",
		"todate":   "2016-11-30",
	}
	bounces, response, err := client.Bounce.GetBounces(500, 0, params)
	if err != nil {
		fmt.Printf("ERR: \n%v\n%v\n", response, err)
	}

	// Output the results.
	fmt.Printf("Bounces: \n%v\n\n", bounces)
}
