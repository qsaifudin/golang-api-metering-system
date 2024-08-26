package tests

import (
	"api_metering_system/metering"
	"testing"
)

func TestApiMetering(t *testing.T) {
	apiMetrics := metering.NewApiMetrics()

	// Simulate 1000 requests
	for i := 1; i <= 1000; i++ {
		apiMetrics.Increment("/api/endpoint1")
		t.Logf("Simulate requests to endpoint1: %d",i)
	}
	t.Logf("Retrieve request counts for all endpoints : %v",apiMetrics.GetMetrics()["/api/endpoint1"] )


	// Simulate the 1001st request
	err := apiMetrics.Increment("/api/endpoint1")
	t.Logf("Simulate the 1001st request : %v",err)
	if err == nil {
		t.Fatal("Expected an error after 1000 requests, but got none")
	}
}
