package main

import (
	"encoding/json"
	"errors"
	"log"
	"sync"
	"time"

	"github.com/google/uuid"
)

// Marketplace represents the logistics marketplace service
type Marketplace struct {
	blockchain *Blockchain

	participants map[string]Participant
	quotes       map[string]FreightQuote
	bids         map[string][]FreightBid
	bookings     map[string]Booking

	mutex sync.RWMutex

	MembershipManager   *MembershipManager
	SubscriptionService *SubscriptionService
}

// NewMarketplace creates a new Marketplace instance
func NewMarketplace(bc *Blockchain) *Marketplace {
	return &Marketplace{
		blockchain:          bc,
		participants:        make(map[string]Participant),
		quotes:              make(map[string]FreightQuote),
		bids:                make(map[string][]FreightBid),
		bookings:            make(map[string]Booking),
		MembershipManager:   NewMembershipManager(),
		SubscriptionService: NewSubscriptionService(),
	}
}

// NewMarketplace creates a new Marketplace instance
func NewMarketplace(bc *Blockchain) *Marketplace {
	return &Marketplace{
		blockchain:        bc,
		participants:      make(map[string]Participant),
		quotes:            make(map[string]FreightQuote),
		bids:              make(map[string][]FreightBid),
		bookings:          make(map[string]Booking),
		MembershipManager: NewMembershipManager(),
	}
}

// NewMarketplace creates a new Marketplace instance
func NewMarketplace(bc *Blockchain) *Marketplace {
	return &Marketplace{
		blockchain:        bc,
		participants:      make(map[string]Participant),
		quotes:            make(map[string]FreightQuote),
		bids:              make(map[string][]FreightBid),
		bookings:          make(map[string]Booking),
		MembershipManager: NewMembershipManager(),
	}
}

// RegisterParticipant registers a new participant in the marketplace
func (m *Marketplace) RegisterParticipant(name string, pType ParticipantType) Participant {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	id := uuid.New().String()
	participant := Participant{
		ID:   id,
		Name: name,
		Type: pType,
	}
	m.participants[id] = participant
	log.Printf("Participant registered: %s (%s)", name, id)
	return participant
}

// CreateFreightQuote creates a new freight quote
func (m *Marketplace) CreateFreightQuote(serviceCategory ServiceCategory, cargoType CargoType, packagingMode PackagingMode, origin, destination string, transportationMode TransportationMode, rate float64, validUntil time.Time) (FreightQuote, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if rate <= 0 {
		return FreightQuote{}, errors.New("rate must be positive")
	}
	if validUntil.Before(time.Now()) {
		return FreightQuote{}, errors.New("validUntil must be in the future")
	}

	id := uuid.New().String()
	quote := FreightQuote{
		ID:                 id,
		ServiceCategory:    serviceCategory,
		CargoType:          cargoType,
		PackagingMode:      packagingMode,
		Origin:             origin,
		Destination:        destination,
		TransportationMode: transportationMode,
		Rate:               rate,
		ValidUntil:         validUntil,
	}
	m.quotes[id] = quote

	// Add to blockchain
	data, err := json.Marshal(quote)
	if err != nil {
		log.Printf("Error marshaling quote: %v", err)
		return FreightQuote{}, err
	}
	err = m.blockchain.AddBlock(string(data))
	if err != nil {
		log.Printf("Error adding quote to blockchain: %v", err)
		return FreightQuote{}, err
	}

	log.Printf("Freight quote created: %s", id)
	return quote, nil
}

// PlaceBid places a bid on a freight quote
func (m *Marketplace) PlaceBid(quoteID, carrierID string, bidAmount float64) (FreightBid, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	quote, exists := m.quotes[quoteID]
	if !exists {
		return FreightBid{}, errors.New("quote not found")
	}

	// Check if carrier exists
	if _, ok := m.participants[carrierID]; !ok {
		return FreightBid{}, errors.New("carrier not found")
	}

	if bidAmount <= 0 {
		return FreightBid{}, errors.New("bid amount must be positive")
	}

	bid := FreightBid{
		ID:         uuid.New().String(),
		QuoteID:    quoteID,
		CarrierID:  carrierID,
		BidAmount:  bidAmount,
		BidTime:    time.Now(),
		IsAccepted: false,
	}
	m.bids[quoteID] = append(m.bids[quoteID], bid)

	// Add to blockchain
	data, err := json.Marshal(bid)
	if err != nil {
		log.Printf("Error marshaling bid: %v", err)
		return FreightBid{}, err
	}
	err = m.blockchain.AddBlock(string(data))
	if err != nil {
		log.Printf("Error adding bid to blockchain: %v", err)
		return FreightBid{}, err
	}

	log.Printf("Bid placed: %s on quote %s", bid.ID, quoteID)
	return bid, nil
}

// ConfirmBooking confirms a booking based on accepted bid
func (m *Marketplace) ConfirmBooking(quoteID, bidID, shipperID string) (Booking, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	bids, exists := m.bids[quoteID]
	if !exists {
		return Booking{}, errors.New("no bids for quote")
	}

	var acceptedBid *FreightBid
	for i, b := range bids {
		if b.ID == bidID {
			acceptedBid = &b
			m.bids[quoteID][i].IsAccepted = true
			break
		}
	}
	if acceptedBid == nil {
		return Booking{}, errors.New("bid not found")
	}

	// Check shipper exists
	if _, ok := m.participants[shipperID]; !ok {
		return Booking{}, errors.New("shipper not found")
	}

	booking := Booking{
		ID:          uuid.New().String(),
		QuoteID:     quoteID,
		BidID:       bidID,
		ShipperID:   shipperID,
		CarrierID:   acceptedBid.CarrierID,
		BookingTime: time.Now(),
		Status:      "Confirmed",
	}
	m.bookings[booking.ID] = booking

	// Add to blockchain
	data, err := json.Marshal(booking)
	if err != nil {
		log.Printf("Error marshaling booking: %v", err)
		return Booking{}, err
	}
	err = m.blockchain.AddBlock(string(data))
	if err != nil {
		log.Printf("Error adding booking to blockchain: %v", err)
		return Booking{}, err
	}

	log.Printf("Booking confirmed: %s", booking.ID)
	return booking, nil
}
