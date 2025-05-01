package main

import (
	"errors"
	"sync"
	"time"
)

// Dispute represents a dispute in the system
type Dispute struct {
	ID          string
	BookingID   string
	RaiserID    string
	Reason      string
	Resolution  string
	Status      string // e.g., "Open", "Resolved", "Rejected"
	CreatedAt   time.Time
	ResolvedAt  *time.Time
}

// DisputeService manages disputes
type DisputeService struct {
	disputes map[string]Dispute
	mutex    sync.RWMutex
}

// NewDisputeService creates a new DisputeService instance
func NewDisputeService() *DisputeService {
	return &DisputeService{
		disputes: make(map[string]Dispute),
	}
}

// RaiseDispute raises a new dispute
func (ds *DisputeService) RaiseDispute(bookingID, raiserID, reason string) (Dispute, error) {
	if bookingID == "" || raiserID == "" || reason == "" {
		return Dispute{}, errors.New("bookingID, raiserID, and reason are required")
	}
	ds.mutex.Lock()
	defer ds.mutex.Unlock()

	id := generateUUID()
	dispute := Dispute{
		ID:        id,
		BookingID: bookingID,
		RaiserID:  raiserID,
		Reason:    reason,
		Status:    "Open",
		CreatedAt: time.Now(),
	}
	ds.disputes[id] = dispute
	return dispute, nil
}

// ResolveDispute resolves an existing dispute
func (ds *DisputeService) ResolveDispute(disputeID, resolution string) error {
	ds.mutex.Lock()
	defer ds.mutex.Unlock()

	dispute, exists := ds.disputes[disputeID]
	if !exists {
		return errors.New("dispute not found")
	}
	if dispute.Status != "Open" {
		return errors.New("dispute already resolved or closed")
	}

	now := time.Now()
	dispute.Resolution = resolution
	dispute.Status = "Resolved"
	dispute.ResolvedAt = &now
	ds.disputes[disputeID] = dispute
	return nil
}

// GetDispute returns a dispute by ID
func (ds *DisputeService) GetDispute(disputeID string) (Dispute, error) {
	ds.mutex.RLock()
	defer ds.mutex.RUnlock()

	dispute, exists := ds.disputes[disputeID]
	if !exists {
		return Dispute{}, errors.New("dispute not found")
	}
	return dispute, nil
}

// generateUUID generates a UUID string (placeholder)
func generateUUID() string {
	// Use github.com/google/uuid or similar in real implementation
	return "dispute-" + time.Now().Format("20060102150405")
}
