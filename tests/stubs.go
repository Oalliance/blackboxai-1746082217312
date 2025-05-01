package tests

import (
	"errors"
	"time"
)

// Stub types and functions to resolve undefined errors in marketplace_test.go

type ParticipantType string
type ServiceCategory string
type CargoType string
type PackagingMode string
type TransportationMode string

const (
	Shipper ParticipantType = "Shipper"
	Carrier ParticipantType = "Carrier"

	Import ServiceCategory = "Import"
	Export ServiceCategory = "Export"

	GeneralCargo CargoType = "GeneralCargo"
	Container   PackagingMode = "Container"
	Sea         TransportationMode = "Sea"
)

type Blockchain struct{}

func NewBlockchain() *Blockchain {
	return &Blockchain{}
}

type Marketplace struct{}

func NewMarketplace(bc *Blockchain) *Marketplace {
	return &Marketplace{}
}

type Participant struct {
	Name string
	Type ParticipantType
}

func (m *Marketplace) RegisterParticipant(name string, pType ParticipantType) Participant {
	return Participant{Name: name, Type: pType}
}

type FreightQuote struct {
	ID          string
	Origin      string
	Destination string
}

func (m *Marketplace) CreateFreightQuote(serviceCategory ServiceCategory, cargoType CargoType, packagingMode PackagingMode, origin, destination string, transportationMode TransportationMode, rate float64, validUntil time.Time) (FreightQuote, error) {
	return FreightQuote{Origin: origin, Destination: destination}, nil
}

type FreightBid struct {
	BidAmount float64
}

func (m *Marketplace) PlaceBid(quoteID, carrierID string, bidAmount float64) (FreightBid, error) {
	return FreightBid{BidAmount: bidAmount}, nil
}

type Booking struct {
	Status string
}

func (m *Marketplace) ConfirmBooking(quoteID, bidID, shipperID string) (Booking, error) {
	return Booking{Status: "Confirmed"}, nil
}
