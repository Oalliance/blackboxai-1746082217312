package tests

import (
	"net/http"
	"testing"
	"time"
)

func TestEndToEnd(t *testing.T) {
	client := &http.Client{Timeout: 10 * time.Second}

	// Register participant
	resp, err := client.Post("http://localhost:8080/api/participants", "application/json", 
		strings.NewReader(`{"name":"Test Shipper","type":"Shipper"}`))
	if err != nil || resp.StatusCode != http.StatusOK {
		t.Fatalf("Failed to register participant: %v", err)
	}

	// Create freight quote
	resp, err = client.Post("http://localhost:8080/api/quotes", "application/json",
		strings.NewReader(`{
			"service_category":"Import",
			"cargo_type":"GeneralCargo",
			"packaging_mode":"Container",
			"origin":"NYC",
			"destination":"LON",
			"transportation_mode":"Sea",
			"rate":1000,
			"valid_until":"2024-12-31T23:59:59Z"
		}`))
	if err != nil || resp.StatusCode != http.StatusOK {
		t.Fatalf("Failed to create freight quote: %v", err)
	}

	// Additional end-to-end steps can be added here
}
