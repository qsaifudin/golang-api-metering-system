package tests

import (
	"api_metering_system/metering"
	"testing"
)

func TestStorageMetering(t *testing.T) {
	storageMetrics := metering.NewStorageMetrics()

	err := storageMetrics.AddFileSize(750 * 1024 * 1024)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}	
	t.Logf("Simulate upload a file and track its size: OK",)
	
	t.Logf("Retrieve total storage used: %v", storageMetrics.GetTotalStorage())
	
	err = storageMetrics.AddFileSize(1 << 30) // Exceed limit
	t.Logf("Simulate exceed limit: %v", err)
	if err == nil {
		t.Fatal("Expected an error, got nil")
	}
}
