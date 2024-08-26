package metering

import "errors"

// MaxRequests is the maximum allowed requests per endpoint.
const MaxRequests = 1000

// ApiMetrics tracks the number of requests for each endpoint.
type ApiMetrics struct {
	metrics map[string]int // Stores request counts for endpoints.
}

// NewApiMetrics creates a new ApiMetrics instance.
func NewApiMetrics() *ApiMetrics {
	return &ApiMetrics{
		metrics: make(map[string]int),
	}
}

// Increment increases the request count for an endpoint.
// Returns an error if the maximum limit is reached.
func (am *ApiMetrics) Increment(endpoint string) error {
	if am.metrics[endpoint] >= MaxRequests {
		return errors.New("request limit exceeded")
	}
	am.metrics[endpoint]++
	return nil
}

// GetMetrics returns the current request counts.
func (am *ApiMetrics) GetMetrics() map[string]int {
	return am.metrics
}
