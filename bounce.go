package postmark

const (
	bounceDeliveryStatsAPIPath = "deliverystats"
)

// BounceService handles communication with the bounce related methods of the
// Postmark API (http://developer.postmarkapp.com/developer-api-bounce.html)
type BounceService service

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

// GetDeliveryStats will build and execute request to send an email via the API.
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
