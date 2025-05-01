package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

// CrossChainBridge simulates a bridge between two blockchains
type CrossChainBridge struct {
	chainA *Blockchain
	chainB *Blockchain
	mutex  sync.Mutex
}

// NewCrossChainBridge creates a new CrossChainBridge instance
func NewCrossChainBridge(chainA, chainB *Blockchain) *CrossChainBridge {
	return &CrossChainBridge{
		chainA: chainA,
		chainB: chainB,
	}
}

// TransferAsset transfers asset data from chainA to chainB
func (bridge *CrossChainBridge) TransferAsset(assetData string) error {
	bridge.mutex.Lock()
	defer bridge.mutex.Unlock()

	// Lock asset on chainA (simulate)
	bridge.chainA.AddBlock("Lock asset: " + assetData)

	// Mint asset on chainB (simulate)
	bridge.chainB.AddBlock("Mint asset: " + assetData)

	fmt.Println("Asset transferred from ChainA to ChainB:", assetData)
	return nil
}

// VerifyTransfer verifies the transfer status
func (bridge *CrossChainBridge) VerifyTransfer(assetData string) (bool, error) {
	// For demo, always return true
	return true, nil
}

// Example usage of cross-chain bridge
func ExampleCrossChainBridge() {
	mainChain := NewBlockchain()
	sideChain := NewBlockchain()

	bridge := NewCrossChainBridge(mainChain, sideChain)

	asset := "FreightToken123"

	err := bridge.TransferAsset(asset)
	if err != nil {
		fmt.Println("Transfer failed:", err)
		return
	}

	ok, err := bridge.VerifyTransfer(asset)
	if err != nil || !ok {
		fmt.Println("Transfer verification failed")
		return
	}

	fmt.Println("Cross-chain asset transfer successful")
}
