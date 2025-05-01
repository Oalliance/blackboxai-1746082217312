package main

import "time"

// ParticipantType defines the type of marketplace participant
type ParticipantType string

const (
	Shipper         ParticipantType = "Shipper"
	Consignee       ParticipantType = "Consignee"
	Carrier         ParticipantType = "Carrier"
	FreightForwarder ParticipantType = "FreightForwarder"
	CustomsBroker   ParticipantType = "CustomsBroker"
)

// Participant represents a marketplace participant
type Participant struct {
	ID   string
	Name string
	Type ParticipantType
}

// ServiceCategory defines the logistics service category
type ServiceCategory string

const (
	Import        ServiceCategory = "Import"
	Export        ServiceCategory = "Export"
	Transit       ServiceCategory = "Transit"
	Transshipment ServiceCategory = "Transshipment"
)

// CargoType defines the type of cargo
type CargoType string

const (
	GeneralCargo CargoType = "GeneralCargo"
	Perishable  CargoType = "Perishable"
	Hazardous   CargoType = "Hazardous"
	Fragile     CargoType = "Fragile"
)

// PackagingMode defines the packaging mode
type PackagingMode string

const (
	Container PackagingMode = "Container"
	Loose     PackagingMode = "Loose"
	Pallet    PackagingMode = "Pallet"
)

// TransportationMode defines mode of transportation
type TransportationMode string

const (
	Sea TransportationMode = "Sea"
	Air TransportationMode = "Air"
	Land TransportationMode = "Land"
)

// FreightQuote represents a freight quotation
type FreightQuote struct {
	ID                 string
	ServiceCategory    ServiceCategory
	CargoType          CargoType
	PackagingMode      PackagingMode
	Origin             string
	Destination        string
	TransportationMode TransportationMode
	Rate               float64
	ValidUntil         time.Time
}

// FreightBid represents a bid on a freight quote
type FreightBid struct {
	ID          string
	QuoteID     string
	CarrierID   string
	BidAmount   float64
	BidTime     time.Time
	IsAccepted  bool
}

// Booking represents a confirmed cargo booking
type Booking struct {
	ID          string
	QuoteID     string
	BidID       string
	ShipperID   string
	CarrierID   string
	BookingTime time.Time
	Status      string
}
