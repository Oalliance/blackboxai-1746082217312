package main

import (
	"encoding/json"
	"errors"
	"sync"
	"time"
)

// FreightQuotationSystem manages freight quotes and bidding with business logic
type FreightQuotationSystem struct {
	marketplace *Marketplace
	mutex       sync.Mutex
}

// NewFreightQuotationSystem creates a new FreightQuotationSystem instance
func NewFreightQuotationSystem(marketplace *Marketplace) *FreightQuotationSystem {
	return &FreightQuotationSystem{
		marketplace: marketplace,
	}
}

func (fqs *FreightQuotationSystem) CreateQuote(serviceCategory ServiceCategory, cargoType CargoType, packagingMode PackagingMode, originCode, destinationCode string, transportationMode TransportationMode, rate float64, validUntil time.Time) (FreightQuote, error) {
	fqs.mutex.Lock()
	defer fqs.mutex.Unlock()

	if rate <= 0 {
		return FreightQuote{}, errors.New("rate must be positive")
	}
	if validUntil.Before(time.Now()) {
		return FreightQuote{}, errors.New("validUntil must be in the future")
	}

	// Validate Transit service category restrictions
	if serviceCategory == Transit {
		if transportationMode != Air && transportationMode != Sea {
			return FreightQuote{}, errors.New("Transit service category can only be Air or Sea transportation mode")
		}
	}

	// Validate origin and destination codes based on transportation mode
	if transportationMode == Air {
		if !validIATAAirportCodes[originCode] {
			return FreightQuote{}, errors.New("origin code is not a valid IATA airport code")
		}
		if !validIATAAirportCodes[destinationCode] {
			return FreightQuote{}, errors.New("destination code is not a valid IATA airport code")
		}
	} else if transportationMode == Sea {
		if !validIMOSeraportCodes[originCode] {
			return FreightQuote{}, errors.New("origin code is not a valid IMO seaport code")
		}
		if !validIMOSeraportCodes[destinationCode] {
			return FreightQuote{}, errors.New("destination code is not a valid IMO seaport code")
		}
	}

	quote := fqs.marketplace.CreateFreightQuote(serviceCategory, cargoType, packagingMode, originCode, destinationCode, transportationMode, rate, validUntil)
	return quote, nil
}

// PlaceBid places a bid on a freight quote with validations
func (fqs *FreightQuotationSystem) PlaceBid(quoteID, carrierID string, bidAmount float64) (FreightBid, error) {
	fqs.mutex.Lock()
	defer fqs.mutex.Unlock()

	if bidAmount <= 0 {
		return FreightBid{}, errors.New("bid amount must be positive")
	}

	quote, exists := fqs.marketplace.quotes[quoteID]
	if !exists {
		return FreightBid{}, errors.New("quote not found")
	}

	if time.Now().After(quote.ValidUntil) {
		return FreightBid{}, errors.New("quote has expired")
	}

	bid, err := fqs.marketplace.PlaceBid(quoteID, carrierID, bidAmount)
	return bid, err
}

// AcceptBid accepts a bid and confirms booking
func (fqs *FreightQuotationSystem) AcceptBid(quoteID, bidID, shipperID string) (Booking, error) {
	fqs.mutex.Lock()
	defer fqs.mutex.Unlock()

	bids, exists := fqs.marketplace.bids[quoteID]
	if !exists {
		return Booking{}, errors.New("no bids for quote")
	}

	var acceptedBid *FreightBid
	for i, b := range bids {
		if b.ID == bidID {
			acceptedBid = &b
			fqs.marketplace.bids[quoteID][i].IsAccepted = true
			break
		}
	}
	if acceptedBid == nil {
		return Booking{}, errors.New("bid not found")
	}

	booking, err := fqs.marketplace.ConfirmBooking(quoteID, bidID, shipperID)
	return booking, err
}

// ListBids lists all bids for a given quote
func (fqs *FreightQuotationSystem) ListBids(quoteID string) ([]FreightBid, error) {
	fqs.mutex.Lock()
	defer fqs.mutex.Unlock()

	bids, exists := fqs.marketplace.bids[quoteID]
	if !exists {
		return nil, errors.New("no bids for quote")
	}
	return bids, nil
}
