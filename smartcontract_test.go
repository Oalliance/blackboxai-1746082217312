package main

import (
	"testing"
	"time"
)

func TestMintAndTransferToken(t *testing.T) {
	marketplace := NewMarketplace(NewBlockchain())
	sc := NewSmartContract(marketplace)

	participantA := "participantA"
	participantB := "participantB"
	tokenID := "TOKEN1"

	// Mint tokens to participantA
	err := sc.MintToken(participantA, 1000)
	if err != nil {
		t.Fatalf("MintToken failed: %v", err)
	}

	// Transfer tokens from participantA to participantB
	err = sc.TransferToken(participantA, participantB, 200)
	if err != nil {
		t.Fatalf("TransferToken failed: %v", err)
	}

	balanceA := sc.TokenLedger.GetBalance(participantA, tokenID)
	balanceB := sc.TokenLedger.GetBalance(participantB, tokenID)

	if balanceA != 800 {
		t.Errorf("Expected balanceA 800, got %f", balanceA)
	}
	if balanceB != 200 {
		t.Errorf("Expected balanceB 200, got %f", balanceB)
	}
}

func TestEscrowLockReleaseRefund(t *testing.T) {
	marketplace := NewMarketplace(NewBlockchain())
	sc := NewSmartContract(marketplace)

	participant := "participant"
	tokenID := "TOKEN1"

	// Mint tokens
	err := sc.MintToken(participant, 500)
	if err != nil {
		t.Fatalf("MintToken failed: %v", err)
	}

	// Lock tokens in escrow
	err = sc.LockTokensInEscrow(participant, tokenID, 300)
	if err != nil {
		t.Fatalf("LockTokensInEscrow failed: %v", err)
	}

	escrowed := sc.TokenLedger.escrowed[participant][tokenID]
	if escrowed != 300 {
		t.Errorf("Expected escrowed 300, got %f", escrowed)
	}

	// Release escrow tokens
	err = sc.ReleaseEscrowTokens(participant, tokenID, 100)
	if err != nil {
		t.Fatalf("ReleaseEscrowTokens failed: %v", err)
	}

	escrowed = sc.TokenLedger.escrowed[participant][tokenID]
	if escrowed != 200 {
		t.Errorf("Expected escrowed 200 after release, got %f", escrowed)
	}

	// Refund escrow tokens
	err = sc.RefundEscrowTokens(participant, tokenID, 200)
	if err != nil {
		t.Fatalf("RefundEscrowTokens failed: %v", err)
	}

	escrowed = sc.TokenLedger.escrowed[participant][tokenID]
	if escrowed != 0 {
		t.Errorf("Expected escrowed 0 after refund, got %f", escrowed)
	}
}

func TestDisputeRaiseResolve(t *testing.T) {
	ds := NewDisputeService()

	bookingID := "booking1"
	raiserID := "user1"
	reason := "Package damaged"

	dispute, err := ds.RaiseDispute(bookingID, raiserID, reason)
	if err != nil {
		t.Fatalf("RaiseDispute failed: %v", err)
	}

	if dispute.Status != "Open" {
		t.Errorf("Expected dispute status Open, got %s", dispute.Status)
	}

	err = ds.ResolveDispute(dispute.ID, "Resolved in favor of shipper")
	if err != nil {
		t.Fatalf("ResolveDispute failed: %v", err)
	}

	resolvedDispute, err := ds.GetDispute(dispute.ID)
	if err != nil {
		t.Fatalf("GetDispute failed: %v", err)
	}

	if resolvedDispute.Status != "Resolved" {
		t.Errorf("Expected dispute status Resolved, got %s", resolvedDispute.Status)
	}
}

func TestTransportModeValidation(t *testing.T) {
	tv := NewTransportationValidator()

	err := tv.ValidateMode("air")
	if err != nil {
		t.Errorf("Expected air mode valid, got error: %v", err)
	}

	err = tv.ValidateMode("invalid_mode")
	if err == nil {
		t.Errorf("Expected error for invalid mode, got nil")
	}
}
