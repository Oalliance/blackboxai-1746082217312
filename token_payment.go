package main

import (
	"errors"
	"sync"
	"time"
)

// Token represents a token with ERC-20/ERC-1155 standard features
type Token struct {
	Name        string
	Symbol      string
	Decimals    int
	TotalSupply float64

	// For ERC-1155 multi-token support
	TokenID string
}

// TokenLedger manages token balances, allowances, and transfers
type TokenLedger struct {
	balances      map[string]map[string]float64 // participantID -> tokenID -> balance
	escrowed      map[string]map[string]float64 // participantID -> tokenID -> escrowed amount
	allowances    map[string]map[string]map[string]float64 // owner -> spender -> tokenID -> allowance
	mutex         sync.Mutex
}

// LockTokensInEscrow locks tokens in escrow for a participant
func (tl *TokenLedger) LockTokensInEscrow(participantID, tokenID string, amount float64) error {
	if amount <= 0 {
		return errors.New("amount must be positive")
	}
	tl.mutex.Lock()
	defer tl.mutex.Unlock()

	if tl.balances[participantID] == nil || tl.balances[participantID][tokenID] < amount {
		return errors.New("insufficient balance to lock in escrow")
	}

	if tl.escrowed[participantID] == nil {
		tl.escrowed[participantID] = make(map[string]float64)
	}

	tl.balances[participantID][tokenID] -= amount
	tl.escrowed[participantID][tokenID] += amount
	return nil
}

// ReleaseEscrowTokens releases escrowed tokens back to participant's balance
func (tl *TokenLedger) ReleaseEscrowTokens(participantID, tokenID string, amount float64) error {
	if amount <= 0 {
		return errors.New("amount must be positive")
	}
	tl.mutex.Lock()
	defer tl.mutex.Unlock()

	if tl.escrowed[participantID] == nil || tl.escrowed[participantID][tokenID] < amount {
		return errors.New("insufficient escrowed tokens to release")
	}

	tl.escrowed[participantID][tokenID] -= amount
	tl.balances[participantID][tokenID] += amount
	return nil
}

// RefundEscrowTokens refunds escrowed tokens to participant's balance (similar to release)
func (tl *TokenLedger) RefundEscrowTokens(participantID, tokenID string, amount float64) error {
	// For now, same as ReleaseEscrowTokens
	return tl.ReleaseEscrowTokens(participantID, tokenID, amount)
}

// NewTokenLedger creates a new TokenLedger instance
func NewTokenLedger() *TokenLedger {
	return &TokenLedger{
		balances:   make(map[string]map[string]float64),
		allowances: make(map[string]map[string]map[string]float64),
	}
}

// MintTokens mints tokens to a participant for a specific tokenID
func (tl *TokenLedger) MintTokens(participantID, tokenID string, amount float64) error {
	if amount <= 0 {
		return errors.New("amount must be positive")
	}
	tl.mutex.Lock()
	defer tl.mutex.Unlock()

	if tl.balances[participantID] == nil {
		tl.balances[participantID] = make(map[string]float64)
	}
	tl.balances[participantID][tokenID] += amount
	return nil
}

// GetBalance returns the token balance of a participant for a specific tokenID
func (tl *TokenLedger) GetBalance(participantID, tokenID string) float64 {
	tl.mutex.Lock()
	defer tl.mutex.Unlock()

	if tl.balances[participantID] == nil {
		return 0
	}
	return tl.balances[participantID][tokenID]
}

// Approve allows a spender to spend tokens on behalf of the owner for a specific tokenID
func (tl *TokenLedger) Approve(ownerID, spenderID, tokenID string, amount float64) error {
	if amount < 0 {
		return errors.New("amount cannot be negative")
	}
	tl.mutex.Lock()
	defer tl.mutex.Unlock()

	if tl.allowances[ownerID] == nil {
		tl.allowances[ownerID] = make(map[string]map[string]float64)
	}
	if tl.allowances[ownerID][spenderID] == nil {
		tl.allowances[ownerID][spenderID] = make(map[string]float64)
	}
	tl.allowances[ownerID][spenderID][tokenID] = amount
	return nil
}

// Allowance returns the remaining allowance a spender has from an owner for a specific tokenID
func (tl *TokenLedger) Allowance(ownerID, spenderID, tokenID string) float64 {
	tl.mutex.Lock()
	defer tl.mutex.Unlock()

	if tl.allowances[ownerID] == nil || tl.allowances[ownerID][spenderID] == nil {
		return 0
	}
	return tl.allowances[ownerID][spenderID][tokenID]
}

// TransferTokens transfers tokens from one participant to another for a specific tokenID
func (tl *TokenLedger) TransferTokens(fromID, toID, tokenID string, amount float64) error {
	if amount <= 0 {
		return errors.New("amount must be positive")
	}
	tl.mutex.Lock()
	defer tl.mutex.Unlock()

	if tl.balances[fromID] == nil || tl.balances[fromID][tokenID] < amount {
		return errors.New("insufficient balance")
	}

	if tl.balances[toID] == nil {
		tl.balances[toID] = make(map[string]float64)
	}

	tl.balances[fromID][tokenID] -= amount
	tl.balances[toID][tokenID] += amount
	return nil
}

// TransferFrom allows a spender to transfer tokens on behalf of the owner for a specific tokenID
func (tl *TokenLedger) TransferFrom(ownerID, spenderID, toID, tokenID string, amount float64) error {
	if amount <= 0 {
		return errors.New("amount must be positive")
	}
	tl.mutex.Lock()
	defer tl.mutex.Unlock()

	if tl.balances[ownerID] == nil || tl.balances[ownerID][tokenID] < amount {
		return errors.New("insufficient balance")
	}
	if tl.allowances[ownerID] == nil || tl.allowances[ownerID][spenderID] == nil || tl.allowances[ownerID][spenderID][tokenID] < amount {
		return errors.New("allowance exceeded")
	}

	if tl.balances[toID] == nil {
		tl.balances[toID] = make(map[string]float64)
	}

	tl.balances[ownerID][tokenID] -= amount
	tl.balances[toID][tokenID] += amount
	tl.allowances[ownerID][spenderID][tokenID] -= amount
	return nil
}

// BatchTransferTokens transfers multiple token amounts for different tokenIDs from one participant to another (ERC-1155)
func (tl *TokenLedger) BatchTransferTokens(fromID, toID string, tokenAmounts map[string]float64) error {
	tl.mutex.Lock()
	defer tl.mutex.Unlock()

	for tokenID, amount := range tokenAmounts {
		if amount <= 0 {
			return errors.New("amount must be positive")
		}
		if tl.balances[fromID] == nil || tl.balances[fromID][tokenID] < amount {
			return errors.New("insufficient balance for token " + tokenID)
		}
	}

	if tl.balances[toID] == nil {
		tl.balances[toID] = make(map[string]float64)
	}

	for tokenID, amount := range tokenAmounts {
		tl.balances[fromID][tokenID] -= amount
		tl.balances[toID][tokenID] += amount
	}
	return nil
}

// TokenPaymentSystem integrates token payments with marketplace and blockchain
type TokenPaymentSystem struct {
	tokenLedger *TokenLedger
	blockchain  *Blockchain
}

// NewTokenPaymentSystem creates a new TokenPaymentSystem instance
func NewTokenPaymentSystem(blockchain *Blockchain) *TokenPaymentSystem {
	return &TokenPaymentSystem{
		tokenLedger: NewTokenLedger(),
		blockchain:  blockchain,
	}
}

// PayFreightBooking processes payment for a booking using tokens
func (tps *TokenPaymentSystem) PayFreightBooking(payerID, payeeID, tokenID string, amount float64, bookingID string) error {
	if amount <= 0 {
		return errors.New("amount must be positive")
	}

	// Transfer tokens
	err := tps.tokenLedger.TransferTokens(payerID, payeeID, tokenID, amount)
	if err != nil {
		return err
	}

	// Record payment on blockchain
	paymentRecord := struct {
		PayerID   string
		PayeeID   string
		TokenID   string
		Amount    float64
		BookingID string
		Timestamp time.Time
	}{
		PayerID:   payerID,
		PayeeID:   payeeID,
		TokenID:   tokenID,
		Amount:    amount,
		BookingID: bookingID,
		Timestamp: time.Now(),
	}
	data, _ := json.Marshal(paymentRecord)
	tps.blockchain.AddBlock(string(data))

	return nil
}

// TokenLedger manages token balances and transfers
type TokenLedger struct {
	balances map[string]float64
	mutex    sync.Mutex
}

// NewTokenLedger creates a new TokenLedger instance
func NewTokenLedger() *TokenLedger {
	return &TokenLedger{
		balances: make(map[string]float64),
	}
}

// MintTokens mints tokens to a participant
func (tl *TokenLedger) MintTokens(participantID string, amount float64) error {
	if amount <= 0 {
		return errors.New("amount must be positive")
	}
	tl.mutex.Lock()
	defer tl.mutex.Unlock()

	tl.balances[participantID] += amount
	return nil
}

// GetBalance returns the token balance of a participant
func (tl *TokenLedger) GetBalance(participantID string) float64 {
	tl.mutex.Lock()
	defer tl.mutex.Unlock()

	return tl.balances[participantID]
}

// TransferTokens transfers tokens from one participant to another
func (tl *TokenLedger) TransferTokens(fromID, toID string, amount float64) error {
	if amount <= 0 {
		return errors.New("amount must be positive")
	}
	tl.mutex.Lock()
	defer tl.mutex.Unlock()

	if tl.balances[fromID] < amount {
		return errors.New("insufficient balance")
	}

	tl.balances[fromID] -= amount
	tl.balances[toID] += amount
	return nil
}

// TokenPaymentSystem integrates token payments with marketplace and blockchain
type TokenPaymentSystem struct {
	tokenLedger *TokenLedger
	blockchain  *Blockchain
}

// NewTokenPaymentSystem creates a new TokenPaymentSystem instance
func NewTokenPaymentSystem(blockchain *Blockchain) *TokenPaymentSystem {
	return &TokenPaymentSystem{
		tokenLedger: NewTokenLedger(),
		blockchain:  blockchain,
	}
}

// PayFreightBooking processes payment for a booking using tokens
func (tps *TokenPaymentSystem) PayFreightBooking(payerID, payeeID string, amount float64, bookingID string) error {
	if amount <= 0 {
		return errors.New("amount must be positive")
	}

	// Transfer tokens
	err := tps.tokenLedger.TransferTokens(payerID, payeeID, amount)
	if err != nil {
		return err
	}

	// Record payment on blockchain
	paymentRecord := struct {
		PayerID   string
		PayeeID   string
		Amount    float64
		BookingID string
		Timestamp time.Time
	}{
		PayerID:   payerID,
		PayeeID:   payeeID,
		Amount:    amount,
		BookingID: bookingID,
		Timestamp: time.Now(),
	}
	data, _ := json.Marshal(paymentRecord)
	tps.blockchain.AddBlock(string(data))

	return nil
}
