package main

import (
	"testing"
	"time"
)

func TestMarketplace_RegisterParticipant(t *testing.T) {
	bc := NewBlockchain()
	marketplace := NewMarketplace(bc)

	participant := marketplace.RegisterParticipant("Test Shipper", ParticipantType("Shipper"))
	if participant.Name != "Test Shipper" {
		t.Errorf("Expected participant name 'Test Shipper', got '%s'", participant.Name)
	}
	if participant.Type != ParticipantType("Shipper") {
		t.Errorf("Expected participant type 'Shipper', got '%s'", participant.Type)
	}
}

func TestMarketplace_CreateFreightQuote(t *testing.T) {
	bc := NewBlockchain()
	marketplace := NewMarketplace(bc)

	validUntil := time.Now().Add(24 * time.Hour)
	quote, err := marketplace.CreateFreightQuote(ServiceCategory("Import"), CargoType("GeneralCargo"), PackagingMode("Container"), "NYC", "LON", TransportationMode("Sea"), 1000.0, validUntil)
	if err != nil {
		t.Fatalf("CreateFreightQuote failed: %v", err)
	}
	if quote.Origin != "NYC" || quote.Destination != "LON" {
		t.Errorf("Quote origin or destination mismatch")
	}
}

func TestMarketplace_PlaceBid(t *testing.T) {
	bc := NewBlockchain()
	marketplace := NewMarketplace(bc)

	shipper := marketplace.RegisterParticipant("Shipper1", ParticipantType("Shipper"))
	carrier := marketplace.RegisterParticipant("Carrier1", ParticipantType("Carrier"))

	validUntil := time.Now().Add(24 * time.Hour)
	quote, err := marketplace.CreateFreightQuote(ServiceCategory("Import"), CargoType("GeneralCargo"), PackagingMode("Container"), "NYC", "LON", TransportationMode("Sea"), 1000.0, validUntil)
	if err != nil {
		t.Fatalf("CreateFreightQuote failed: %v", err)
	}

	bid, err := marketplace.PlaceBid(quote.ID, carrier.ID, 900.0)
	if err != nil {
		t.Fatalf("PlaceBid failed: %v", err)
	}
	if bid.BidAmount != 900.0 {
		t.Errorf("Bid amount mismatch")
	}
}

func TestMarketplace_ConfirmBooking(t *testing.T) {
	bc := NewBlockchain()
	marketplace := NewMarketplace(bc)

	shipper := marketplace.RegisterParticipant("Shipper1", ParticipantType("Shipper"))
	carrier := marketplace.RegisterParticipant("Carrier1", ParticipantType("Carrier"))

	validUntil := time.Now().Add(24 * time.Hour)
	quote, err := marketplace.CreateFreightQuote(ServiceCategory("Import"), CargoType("GeneralCargo"), PackagingMode("Container"), "NYC", "LON", TransportationMode("Sea"), 1000.0, validUntil)
	if err != nil {
		t.Fatalf("CreateFreightQuote failed: %v", err)
	}

	bid, err := marketplace.PlaceBid(quote.ID, carrier.ID, 900.0)
	if err != nil {
		t.Fatalf("PlaceBid failed: %v", err)
	}

	booking, err := marketplace.ConfirmBooking(quote.ID, bid.ID, shipper.ID)
	if err != nil {
		t.Fatalf("ConfirmBooking failed: %v", err)
	}
	if booking.Status != "Confirmed" {
		t.Errorf("Booking status mismatch")
	}
}
