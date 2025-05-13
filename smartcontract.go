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

	// Reentrancy guard mutex
	reentrancyLock sync.Mutex
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
	sc.reentrancyLock.Lock()
	defer sc.reentrancyLock.Unlock()

	if err := validateAddress(participantID); err != nil {
		return err
	}
	if amount <= 0 {
		return errors.New("amount must be greater than zero")
	}
	// Use safe math for overflow/underflow protection
	safeAmount, err := safeAddFloat64(0, amount)
	if err != nil {
		return err
	}
	logger := getLogger()
	logger.LogEvent("MintToken called for participant: " + participantID)
	// State changes happen before external calls inside MintTokens
	return sc.TokenLedger.MintTokens(participantID, safeAmount)
}

// safeAddFloat64 safely adds two float64 numbers and checks for overflow
func safeAddFloat64(a, b float64) (float64, error) {
	result := a + b
	if (result < a) != (b < 0) {
		return 0, errors.New("float64 addition overflow or underflow")
	}
	return result, nil
}

func (sc *SmartContract) TransferToken(fromID, toID string, amount float64) error {
	sc.reentrancyLock.Lock()
	defer sc.reentrancyLock.Unlock()

	if err := validateAddress(fromID); err != nil {
		return err
	}
	if err := validateAddress(toID); err != nil {
		return err
	}
	if amount <= 0 {
		return errors.New("amount must be greater than zero")
	}
	// Use safe math for overflow/underflow protection
	safeAmount, err := safeAddFloat64(0, amount)
	if err != nil {
		return err
	}
	logger := getLogger()
	logger.LogEvent("TransferToken called from " + fromID + " to " + toID)
	// State changes happen before external calls inside TransferTokens
	success, err := sc.TokenLedger.TransferTokensWithCheck(fromID, toID, safeAmount)
	if err != nil {
		return err
	}
	if !success {
		return errors.New("token transfer failed")
	}
	return nil
}

// safeAddFloat64 safely adds two float64 numbers and checks for overflow
func safeAddFloat64(a, b float64) (float64, error) {
	result := a + b
	if (result < a) != (b < 0) {
		return 0, errors.New("float64 addition overflow or underflow")
	}
	return result, nil
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

// GetRandomness simulates getting randomness securely using an oracle or commitment scheme
func (sc *SmartContract) GetRandomness(seed string) (int64, error) {
	// Placeholder for oracle-based or commitment scheme randomness
	// In production, integrate Chainlink VRF or similar secure randomness source
	randomValue := sc.pseudoRandom(seed)
	return randomValue, nil
}

// pseudoRandom is a simple deterministic pseudo-random generator for demonstration only
func (sc *SmartContract) pseudoRandom(seed string) int64 {
	hash := int64(0)
	for _, c := range seed {
		hash = (hash*31 + int64(c)) % 1000000007
	}
	return hash
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
	// Restrict function to operate only within validated and predictable conditions to avoid flash loan reliance
	if !sc.isValidQuoteRequest(origin, destination, rate) {
		return FreightQuote{}, errors.New("invalid quote request parameters")
	}
	// Add business logic, validations, and emit events if needed
	if rate <= 0 {
		return FreightQuote{}, errors.New("rate must be positive")
	}
	quote := sc.Marketplace.CreateFreightQuote(serviceCategory, cargoType, packagingMode, origin, destination, transportationMode, rate, validUntil)
	// Emit event or add to blockchain handled by marketplace
	return quote, nil
}

func (sc *SmartContract) isValidQuoteRequest(origin, destination string, rate float64) bool {
	// Implement validation logic to ensure request is legitimate and not flash loan manipulation
	if origin == "" || destination == "" || rate <= 0 {
		return false
	}
	// Additional checks can be added here
	return true
}

// PlaceBid places a bid on a freight quote via smart contract logic
func (sc *SmartContract) PlaceBid(quoteID, carrierID string, bidAmount float64) (FreightBid, error) {
	// Access control: restrict to authorized participants only
	if !sc.isAuthorizedParticipant(carrierID) {
		return FreightBid{}, errors.New("unauthorized participant")
	}
	if bidAmount <= 0 {
		return FreightBid{}, errors.New("bid amount must be positive")
	}
	bid, err := sc.Marketplace.PlaceBid(quoteID, carrierID, bidAmount)
	return bid, err
}

func (sc *SmartContract) isAuthorizedParticipant(participantID string) bool {
	// Implement access control logic here
	// For example, check if participant has active membership
	active, err := sc.CheckMembershipActive(participantID)
	if err != nil {
		return false
	}
	return active
}

// ConfirmBooking confirms a booking via smart contract logic
func (sc *SmartContract) ConfirmBooking(quoteID, bidID, shipperID string) (Booking, error) {
	// Access control: restrict to authorized participants only
	if !sc.isAuthorizedParticipant(shipperID) {
		return Booking{}, errors.New("unauthorized participant")
	}
	booking, err := sc.Marketplace.ConfirmBooking(quoteID, bidID, shipperID)
	return booking, err
}

// CreateProposal creates a governance proposal via smart contract logic
func (sc *SmartContract) CreateProposal(title, description, proposerID string) (Proposal, error) {
	// Access control: restrict to authorized participants only
	if !sc.isAuthorizedParticipant(proposerID) {
		return Proposal{}, errors.New("unauthorized participant")
	}
	if title == "" || description == "" {
		return Proposal{}, errors.New("title and description cannot be empty")
	}
	proposal := sc.Marketplace.blockchain.governance.CreateProposal(title, description, proposerID)
	return proposal, nil
}

// VoteProposal casts a vote on a proposal via smart contract logic
func (sc *SmartContract) VoteProposal(proposalID, participantID string, approve bool) error {
	// Access control: restrict to authorized participants only
	if !sc.isAuthorizedParticipant(participantID) {
		return errors.New("unauthorized participant")
	}
	err := sc.Marketplace.blockchain.governance.VoteProposal(proposalID, participantID, approve)
	return err
}
