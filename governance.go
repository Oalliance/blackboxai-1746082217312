package main

import (
	"encoding/json"
	"errors"
	"sync"
	"time"
	"github.com/google/uuid"
)

// Governance represents blockchain governance module
type Governance struct {
	blockchain          *Blockchain
	proposals           map[string]Proposal
	mutex               sync.Mutex
	membershipManager   *MembershipManager
	subscriptionService *SubscriptionService
}

// NewGovernance creates a new Governance instance
func NewGovernance(bc *Blockchain, mm *MembershipManager, ss *SubscriptionService) *Governance {
	return &Governance{
		blockchain:          bc,
		proposals:           make(map[string]Proposal),
		membershipManager:   mm,
		subscriptionService: ss,
	}
}

// CreateProposal creates a new governance proposal
func (g *Governance) CreateProposal(title, description, proposerID string) (Proposal, error) {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	// Check subscription active
	if g.subscriptionService != nil {
		if !g.subscriptionService.CheckActive(proposerID) {
			return Proposal{}, errors.New("proposer does not have an active subscription")
		}
	}

	id := uuid.New().String()
	proposal := Proposal{
		ID:          id,
		Title:       title,
		Description: description,
		ProposerID:  proposerID,
		CreatedAt:   time.Now(),
		Status:      ProposalPending,
		Votes:       make(map[string]bool),
	}
	g.proposals[id] = proposal

	// Add to blockchain
	data, _ := json.Marshal(proposal)
	g.blockchain.AddBlock(string(data))

	return proposal, nil
}

// VoteProposal casts a vote on a proposal
func (g *Governance) VoteProposal(proposalID, participantID string, approve bool) error {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	// Check subscription active
	if g.subscriptionService != nil {
		if !g.subscriptionService.CheckActive(participantID) {
			return errors.New("participant does not have an active subscription")
		}
	}

	proposal, exists := g.proposals[proposalID]
	if !exists {
		return errors.New("proposal not found")
	}

	// Check if participant already voted
	if _, voted := proposal.Votes[participantID]; voted {
		return errors.New("participant already voted")
	}

	proposal.Votes[participantID] = approve

	// Update proposal status if majority reached (simple majority)
	approveCount := 0
	rejectCount := 0
	for _, v := range proposal.Votes {
		if v {
			approveCount++
		} else {
			rejectCount++
		}
	}

	// For simplicity, if votes >= 3, decide
	if len(proposal.Votes) >= 3 {
		if approveCount > rejectCount {
			proposal.Status = ProposalApproved
		} else {
			proposal.Status = ProposalRejected
		}
	}

	g.proposals[proposalID] = proposal

	// Add vote to blockchain
	voteData := struct {
		ProposalID    string
		ParticipantID string
		Approve       bool
	}{
		ProposalID:    proposalID,
		ParticipantID: participantID,
		Approve:       approve,
	}
	data, _ := json.Marshal(voteData)
	g.blockchain.AddBlock(string(data))

	return nil
}

// ProposalStatus defines status of a governance proposal
type ProposalStatus string

const (
	ProposalPending  ProposalStatus = "Pending"
	ProposalApproved ProposalStatus = "Approved"
	ProposalRejected ProposalStatus = "Rejected"
)

// Proposal represents a governance proposal
type Proposal struct {
	ID          string
	Title       string
	Description string
	ProposerID  string
	CreatedAt   time.Time
	Status      ProposalStatus
	Votes       map[string]bool // participantID -> vote (true=approve, false=reject)
}

// NewGovernance creates a new Governance instance
func NewGovernance(bc *Blockchain) *Governance {
	return &Governance{
		blockchain: bc,
		proposals:  make(map[string]Proposal),
	}
}

// CreateProposal creates a new governance proposal
func (g *Governance) CreateProposal(title, description, proposerID string) Proposal {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	id := uuid.New().String()
	proposal := Proposal{
		ID:          id,
		Title:       title,
		Description: description,
		ProposerID:  proposerID,
		CreatedAt:   time.Now(),
		Status:      ProposalPending,
		Votes:       make(map[string]bool),
	}
	g.proposals[id] = proposal

	// Add to blockchain
	data, _ := json.Marshal(proposal)
	g.blockchain.AddBlock(string(data))

	return proposal
}

// VoteProposal casts a vote on a proposal
func (g *Governance) VoteProposal(proposalID, participantID string, approve bool) error {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	proposal, exists := g.proposals[proposalID]
	if !exists {
		return errors.New("proposal not found")
	}

	// Check if participant already voted
	if _, voted := proposal.Votes[participantID]; voted {
		return errors.New("participant already voted")
	}

	proposal.Votes[participantID] = approve

	// Update proposal status if majority reached (simple majority)
	approveCount := 0
	rejectCount := 0
	for _, v := range proposal.Votes {
		if v {
			approveCount++
		} else {
			rejectCount++
		}
	}

	// For simplicity, if votes >= 3, decide
	if len(proposal.Votes) >= 3 {
		if approveCount > rejectCount {
			proposal.Status = ProposalApproved
		} else {
			proposal.Status = ProposalRejected
		}
	}

	g.proposals[proposalID] = proposal

	// Add vote to blockchain
	voteData := struct {
		ProposalID    string
		ParticipantID string
		Approve       bool
	}{
		ProposalID:    proposalID,
		ParticipantID: participantID,
		Approve:       approve,
	}
	data, _ := json.Marshal(voteData)
	g.blockchain.AddBlock(string(data))

	return nil
}
