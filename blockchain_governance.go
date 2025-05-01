package main

import (
	"encoding/json"
	"errors"
	"sync"
	"time"
)

// EnhancedGovernance represents an enhanced blockchain governance module
type EnhancedGovernance struct {
	blockchain *Blockchain
	proposals  map[string]Proposal
	mutex      sync.Mutex
	quorum     int
}

// NewEnhancedGovernance creates a new EnhancedGovernance instance with quorum
func NewEnhancedGovernance(bc *Blockchain, quorum int) *EnhancedGovernance {
	return &EnhancedGovernance{
		blockchain: bc,
		proposals:  make(map[string]Proposal),
		quorum:     quorum,
	}
}

// CreateProposal creates a new governance proposal
func (g *EnhancedGovernance) CreateProposal(title, description, proposerID string) Proposal {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	id := generateID()
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
func (g *EnhancedGovernance) VoteProposal(proposalID, participantID string, approve bool) error {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	proposal, exists := g.proposals[proposalID]
	if !exists {
		return errors.New("proposal not found")
	}

	if _, voted := proposal.Votes[participantID]; voted {
		return errors.New("participant already voted")
	}

	proposal.Votes[participantID] = approve

	approveCount := 0
	rejectCount := 0
	for _, v := range proposal.Votes {
		if v {
			approveCount++
		} else {
			rejectCount++
		}
	}

	// Check if quorum reached
	if len(proposal.Votes) >= g.quorum {
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
		Timestamp     time.Time
	}{
		ProposalID:    proposalID,
		ParticipantID: participantID,
		Approve:       approve,
		Timestamp:     time.Now(),
	}
	data, _ := json.Marshal(voteData)
	g.blockchain.AddBlock(string(data))

	return nil
}

// generateID generates a unique ID (placeholder)
func generateID() string {
	return time.Now().Format("20060102150405.000000000")
}
