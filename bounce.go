package postmark

import (
	"fmt"
	"net/url"
	"strconv"
	"time"
)

const (
	bounceDeliveryStatsAPIPath = "deliverystats"
	bounceBouncesAPIPath       = "bounces"
)

// BounceService handles communication with the bounce related methods of the
// Postmark API (http://developer.postmarkapp.com/developer-api-bounce.html)
type BounceService service

/*
 * GetDeliveryStats --------------------------------------------------------- */

// BounceType represents the type of bounce with a count.
type BounceType struct {
	Type  string
	Name  string
	Count int64
}

// BounceDeliveryStats represents a collection of bounce stats.
type BounceDeliveryStats struct {
	InactiveMails int64
	Bounces       []BounceType
}

// GetDeliveryStats will return all delivery stats aggregated by bounce type.
func (s *BounceService) GetDeliveryStats() (*BounceDeliveryStats, *Response, error) {
	request, err := s.client.NewRequest("GET", bounceDeliveryStatsAPIPath, nil)
	if err != nil {
		return nil, nil, err
	}

	deliveryStats := &BounceDeliveryStats{}
	response, err := s.client.Do(request, deliveryStats)
	if err != nil {
		return nil, response, err
	}

	return deliveryStats, response, nil
}

/*
 * GetBounces --------------------------------------------------------- */

// Bounces represents a slice of bounces and a given count.
type Bounces struct {
	TotalCount int64
	Bounces    []Bounce
}

// Bounce represents a BounceType in further detail.
type Bounce struct {
	ID            int64
	Type          string
	TypeCode      int64
	Name          string
	Tag           string
	MessageID     string
	Description   string
	Details       string
	Email         string
	BouncedAt     time.Time
	DumpAvailable bool
	Inactive      bool
	CanActivate   bool
	Subject       string
}

// GetBounces will return all bounces.
func (s *BounceService) GetBounces(bounceCount, bounceOffset int, parameters map[string]interface{}) (*Bounces, *Response, error) {

	// Construct query parameters.
	values := &url.Values{}
	values.Add("count", strconv.Itoa(bounceCount))
	values.Add("offset", strconv.Itoa(bounceOffset))
	for key, value := range parameters {
		values.Add(key, fmt.Sprintf("%v", value))
	}

	request, err := s.client.NewRequest("GET", fmt.Sprintf("%s?%s", bounceDeliveryStatsAPIPath, values.Encode()), nil)
	if err != nil {
		return nil, nil, err
	}

	bounces := Bounces{}
	response, err := s.client.Do(request, &bounces)
	if err != nil {
		return nil, nil, err
	}

	return &bounces, response, nil
}
