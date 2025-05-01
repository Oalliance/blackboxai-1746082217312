package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

// Role defines user roles for access control
type Role string

const (
	AdminRole     Role = "Admin"
	ShipperRole   Role = "Shipper"
	CarrierRole   Role = "Carrier"
	BrokerRole    Role = "Broker"
	ForwarderRole Role = "Forwarder"
)

// AccessControl manages role-based permissions
type AccessControl struct {
	userRoles map[string]Role
	mutex     sync.Mutex
}

// NewAccessControl creates a new AccessControl instance
func NewAccessControl() *AccessControl {
	return &AccessControl{
		userRoles: make(map[string]Role),
	}
}

// AssignRole assigns a role to a user
func (ac *AccessControl) AssignRole(userID string, role Role) {
	ac.mutex.Lock()
	defer ac.mutex.Unlock()
	ac.userRoles[userID] = role
}

// CheckRole checks if a user has a specific role
func (ac *AccessControl) CheckRole(userID string, role Role) bool {
	ac.mutex.Lock()
	defer ac.mutex.Unlock()
	return ac.userRoles[userID] == role
}

// MultiSigAuthorization simulates multi-signature authorization
type MultiSigAuthorization struct {
	requiredSigs int
	signers     map[string]bool
	mutex       sync.Mutex
}

// NewMultiSigAuthorization creates a new MultiSigAuthorization instance
func NewMultiSigAuthorization(required int) *MultiSigAuthorization {
	return &MultiSigAuthorization{
		requiredSigs: required,
		signers:     make(map[string]bool),
	}
}

// Sign adds a signature from a user
func (msa *MultiSigAuthorization) Sign(userID string) {
	msa.mutex.Lock()
	defer msa.mutex.Unlock()
	msa.signers[userID] = true
}

// IsAuthorized checks if required signatures are met
func (msa *MultiSigAuthorization) IsAuthorized() bool {
	msa.mutex.Lock()
	defer msa.mutex.Unlock()
	return len(msa.signers) >= msa.requiredSigs
}

// Escrow represents an escrow with time-locked release
type Escrow struct {
	amount   float64
	payer    string
	payee    string
	lockTime time.Time
	released bool
	mutex    sync.Mutex
}

// NewEscrow creates a new Escrow instance
func NewEscrow(amount float64, payer, payee string, lockDuration time.Duration) *Escrow {
	return &Escrow{
		amount:   amount,
		payer:    payer,
		payee:    payee,
		lockTime: time.Now().Add(lockDuration),
		released: false,
	}
}

// Release releases funds if lock time passed
func (e *Escrow) Release() error {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	if e.released {
		return errors.New("funds already released")
	}
	if time.Now().Before(e.lockTime) {
		return errors.New("lock time not reached")
	}
	e.released = true
	fmt.Printf("Released %.2f from %s to %s\n", e.amount, e.payer, e.payee)
	return nil
}

// ProxyContract simulates upgradeable contract via proxy pattern
type ProxyContract struct {
	implementation interface{}
	mutex          sync.Mutex

	// Compliance tracking data
	complianceRecords map[string]ComplianceRecord
}

// ComplianceRecord represents a compliance event or status
type ComplianceRecord struct {
	ParticipantID string
	Event        string
	Timestamp    time.Time
	Details      string
}

// NewProxyContract creates a new ProxyContract instance
func NewProxyContract(impl interface{}) *ProxyContract {
	return &ProxyContract{
		implementation:    impl,
		complianceRecords: make(map[string]ComplianceRecord),
	}
}

// Upgrade upgrades the implementation with authorization
func (p *ProxyContract) Upgrade(newImpl interface{}, msa *MultiSigAuthorization) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if !msa.IsAuthorized() {
		return errors.New("not authorized to upgrade")
	}
	p.implementation = newImpl
	fmt.Println("Contract upgraded successfully")
	return nil
}

// AddComplianceRecord adds a compliance record for a participant
func (p *ProxyContract) AddComplianceRecord(participantID, event, details string) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.complianceRecords[participantID] = ComplianceRecord{
		ParticipantID: participantID,
		Event:         event,
		Timestamp:     time.Now(),
		Details:       details,
	}
}

// GetComplianceRecord retrieves a compliance record for a participant
func (p *ProxyContract) GetComplianceRecord(participantID string) (ComplianceRecord, bool) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	record, exists := p.complianceRecords[participantID]
	return record, exists
}
