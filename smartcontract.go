package main

import (
	"encoding/json"
	"errors"
	"log"
	"sync"
	"time"
)

// SmartContract simulates a smart contract for the logistics marketplace
type SmartContract struct {
	// State can be stored here or in blockchain data
	Marketplace *Marketplace

	// Token ledger for ERC-20/ERC-1155 tokenization
	TokenLedger *TokenLedger

	// Dispute service for dispute resolution
	disputeService *DisputeService

	// Transportation validator for transport mode logic
	transportValidator *TransportationValidator

	// Membership manager for subscription management
	membershipManager *MembershipManager

	// Benefit registry for subscription benefits
	benefitRegistry map[string]Benefit

	eventListeners      map[string][]func(data interface{})
	eventListenersMutex sync.RWMutex
}

// Benefit represents a subscription benefit
type Benefit struct {
	ID          string
	Title       string
	Description string
	Active      bool
}

// AddBenefit adds a new benefit to the registry
func (sc *SmartContract) AddBenefit(id, title, description string) error {
	if sc.benefitRegistry == nil {
		sc.benefitRegistry = make(map[string]Benefit)
	}
	if _, exists := sc.benefitRegistry[id]; exists {
		return errors.New("benefit already exists")
	}
	sc.benefitRegistry[id] = Benefit{
		ID:          id,
		Title:       title,
		Description: description,
		Active:      true,
	}
	return nil
}

// RemoveBenefit deactivates a benefit in the registry
func (sc *SmartContract) RemoveBenefit(id string) error {
	if sc.benefitRegistry == nil {
		return errors.New("benefit registry not initialized")
	}
	benefit, exists := sc.benefitRegistry[id]
	if !exists {
		return errors.New("benefit not found")
	}
	benefit.Active = false
	sc.benefitRegistry[id] = benefit
	return nil
}

// ListBenefits lists all active benefits
func (sc *SmartContract) ListBenefits() []Benefit {
	benefits := []Benefit{}
	for _, b := range sc.benefitRegistry {
		if b.Active {
			benefits = append(benefits, b)
		}
	}
	return benefits
}

// InitializeServices initializes auxiliary services for the smart contract
func (sc *SmartContract) InitializeServices() {
	sc.disputeService = NewDisputeService()
	sc.transportValidator = NewTransportationValidator()
	sc.membershipManager = NewMembershipManager()
	sc.benefitRegistry = make(map[string]Benefit)
	sc.eventListeners = make(map[string][]func(data interface{}))
	sc.eventListenersMutex = sync.RWMutex{}
}

// SubscribeMembership subscribes a participant to a membership
func (sc *SmartContract) SubscribeMembership(participantID string, mType MembershipType) error {
	if sc.membershipManager == nil {
		return errors.New("membership manager not initialized")
	}
	_, err := sc.membershipManager.Subscribe(participantID, mType)
	return err
}

// CheckMembershipActive checks if a participant has an active membership
func (sc *SmartContract) CheckMembershipActive(participantID string) (bool, error) {
	if sc.membershipManager == nil {
		return false, errors.New("membership manager not initialized")
	}
	active := sc.membershipManager.CheckActive(participantID)
	return active, nil
}

func (sc *SmartContract) InitializeServices() {
	sc.disputeService = NewDisputeService()
	sc.transportValidator = NewTransportationValidator()
	sc.membershipManager = NewMembershipManager()
}

// TransportModeSpecificLogic applies logic based on transport mode
func (sc *SmartContract) TransportModeSpecificLogic(transportMode string) error {
	if sc.transportValidator == nil {
		return errors.New("transportation validator not initialized")
	}
	err := sc.transportValidator.ValidateMode(transportMode)
	if err != nil {
		return err
	}
	return sc.transportValidator.ApplyModeSpecificLogic(transportMode)
}

// LockTokensInEscrow locks tokens in escrow for a participant
func (sc *SmartContract) LockTokensInEscrow(participantID, tokenID string, amount float64) error {
	return sc.TokenLedger.LockTokensInEscrow(participantID, tokenID, amount)
}

// ReleaseEscrowTokens releases escrowed tokens back to participant's balance
func (sc *SmartContract) ReleaseEscrowTokens(participantID, tokenID string, amount float64) error {
	return sc.TokenLedger.ReleaseEscrowTokens(participantID, tokenID, amount)
}

// RefundEscrowTokens refunds escrowed tokens to participant's balance
import (
	"log"
	"sync"
	"time"
	"errors"
)

var logger *Logger
var once sync.Once

func getLogger() *Logger {
	once.Do(func() {
		var err error
		logger, err = NewLogger("marketplace.log")
		if err != nil {
			log.Fatalf("Failed to initialize logger: %v", err)
		}
	})
	return logger
}

func NewSmartContract(marketplace *Marketplace) *SmartContract {
	logger := getLogger()
	logger.LogEvent("SmartContract instance created")
	return &SmartContract{
		Marketplace: marketplace,
		TokenLedger: NewTokenLedger(),
	}
}

func (sc *SmartContract) MintToken(participantID string, amount float64) error {
	logger := getLogger()
	logger.LogEvent("MintToken called for participant: " + participantID)
	return sc.TokenLedger.MintTokens(participantID, amount)
}

func (sc *SmartContract) TransferToken(fromID, toID string, amount float64) error {
	logger := getLogger()
	logger.LogEvent("TransferToken called from " + fromID + " to " + toID)
	return sc.TokenLedger.TransferTokens(fromID, toID, amount)
}

func (sc *SmartContract) RefundEscrowTokens(participantID, tokenID string, amount float64) error {
	return sc.TokenLedger.RefundEscrowTokens(participantID, tokenID, amount)
}


import (
	"log"
	"sync"
	"time"
	"errors"
)

var logger *Logger
var once sync.Once

func getLogger() *Logger {
	once.Do(func() {
		var err error
		logger, err = NewLogger("marketplace.log")
		if err != nil {
			log.Fatalf("Failed to initialize logger: %v", err)
		}
	})
	return logger
}

// NewSmartContract creates a new SmartContract instance
func NewSmartContract(marketplace *Marketplace) *SmartContract {
	logger := getLogger()
	logger.LogEvent("SmartContract instance created")
	return &SmartContract{
		Marketplace: marketplace,
		TokenLedger: NewTokenLedger(),
	}
}

func (sc *SmartContract) MintToken(participantID string, amount float64) error {
	logger := getLogger()
	logger.LogEvent("MintToken called for participant: " + participantID)
	return sc.TokenLedger.MintTokens(participantID, amount)
}

func (sc *SmartContract) TransferToken(fromID, toID string, amount float64) error {
	logger := getLogger()
	logger.LogEvent("TransferToken called from " + fromID + " to " + toID)
	return sc.TokenLedger.TransferTokens(fromID, toID, amount)
}

import "sync"

// ListenToEvent simulates adding an event listener for contract events
func (sc *SmartContract) ListenToEvent(eventName string, callback func(data interface{})) {
	sc.eventListenersMutex.Lock()
	defer sc.eventListenersMutex.Unlock()

	if sc.eventListeners == nil {
		sc.eventListeners = make(map[string][]func(data interface{}))
	}
	sc.eventListeners[eventName] = append(sc.eventListeners[eventName], callback)
}

// EmitEvent emits an event to all registered listeners
func (sc *SmartContract) EmitEvent(eventName string, data interface{}) {
	sc.eventListenersMutex.RLock()
	defer sc.eventListenersMutex.RUnlock()

	if listeners, exists := sc.eventListeners[eventName]; exists {
		for _, listener := range listeners {
			go listener(data)
		}
	}
}

// RaiseDispute allows raising a dispute for a booking
func (sc *SmartContract) RaiseDispute(bookingID, raiserID, reason string) error {
	if sc.disputeService == nil {
		return errors.New("dispute service not initialized")
	}
	_, err := sc.disputeService.RaiseDispute(bookingID, raiserID, reason)
	return err
}

// ResolveDispute resolves a dispute by an authorized participant
func (sc *SmartContract) ResolveDispute(disputeID, resolverID string, resolution string) error {
	if sc.disputeService == nil {
		return errors.New("dispute service not initialized")
	}
	err := sc.disputeService.ResolveDispute(disputeID, resolution)
	return err
}

// TransportModeSpecificLogic applies logic based on transport mode
func (sc *SmartContract) TransportModeSpecificLogic(transportMode string) error {
	if sc.transportValidator == nil {
		return errors.New("transportation validator not initialized")
	}
	err := sc.transportValidator.ValidateMode(transportMode)
	if err != nil {
		return err
	}
	return sc.transportValidator.ApplyModeSpecificLogic(transportMode)
}

// Add eventListeners map and mutex to SmartContract struct
type SmartContract struct {
	// existing fields...

	eventListeners      map[string][]func(data interface{})
	eventListenersMutex sync.RWMutex
}

// RaiseDispute allows raising a dispute for a booking
func (sc *SmartContract) RaiseDispute(bookingID, raiserID, reason string) error {
	if sc.disputeService == nil {
		return errors.New("dispute service not initialized")
	}
	_, err := sc.disputeService.RaiseDispute(bookingID, raiserID, reason)
	return err
}

// ResolveDispute resolves a dispute by an authorized participant
func (sc *SmartContract) ResolveDispute(disputeID, resolverID string, resolution string) error {
	if sc.disputeService == nil {
		return errors.New("dispute service not initialized")
	}
	err := sc.disputeService.ResolveDispute(disputeID, resolution)
	return err
}

// TransportModeSpecificLogic applies logic based on transport mode
func (sc *SmartContract) TransportModeSpecificLogic(transportMode string) error {
	// Add transport mode specific logic here
	return nil
}

import (
	"log"
	"sync"
	"time"
	"errors"
)

var logger *Logger
var once sync.Once

func getLogger() *Logger {
	once.Do(func() {
		var err error
		logger, err = NewLogger("marketplace.log")
		if err != nil {
			log.Fatalf("Failed to initialize logger: %v", err)
		}
	})
	return logger
}

func NewSmartContract(marketplace *Marketplace) *SmartContract {
	logger := getLogger()
	logger.LogEvent("SmartContract instance created")
	return &SmartContract{
		Marketplace: marketplace,
		TokenLedger: NewTokenLedger(),
	}
}

// CreateFreightQuote creates a freight quote via smart contract logic
func (sc *SmartContract) CreateFreightQuote(serviceCategory ServiceCategory, cargoType CargoType, packagingMode PackagingMode, origin, destination string, transportationMode TransportationMode, rate float64, validUntil time.Time) (FreightQuote, error) {
	// Add business logic, validations, and emit events if needed
	if rate <= 0 {
		return FreightQuote{}, errors.New("rate must be positive")
	}
	quote := sc.Marketplace.CreateFreightQuote(serviceCategory, cargoType, packagingMode, origin, destination, transportationMode, rate, validUntil)
	// Emit event or add to blockchain handled by marketplace
	return quote, nil
}

// PlaceBid places a bid on a freight quote via smart contract logic
func (sc *SmartContract) PlaceBid(quoteID, carrierID string, bidAmount float64) (FreightBid, error) {
	if bidAmount <= 0 {
		return FreightBid{}, errors.New("bid amount must be positive")
	}
	bid, err := sc.Marketplace.PlaceBid(quoteID, carrierID, bidAmount)
	return bid, err
}

// ConfirmBooking confirms a booking via smart contract logic
func (sc *SmartContract) ConfirmBooking(quoteID, bidID, shipperID string) (Booking, error) {
	booking, err := sc.Marketplace.ConfirmBooking(quoteID, bidID, shipperID)
	return booking, err
}

// CreateProposal creates a governance proposal via smart contract logic
func (sc *SmartContract) CreateProposal(title, description, proposerID string) (Proposal, error) {
	if title == "" || description == "" {
		return Proposal{}, errors.New("title and description cannot be empty")
	}
	proposal := sc.Marketplace.blockchain.governance.CreateProposal(title, description, proposerID)
	return proposal, nil
}

// VoteProposal casts a vote on a proposal via smart contract logic
func (sc *SmartContract) VoteProposal(proposalID, participantID string, approve bool) error {
	err := sc.Marketplace.blockchain.governance.VoteProposal(proposalID, participantID, approve)
	return err
}
