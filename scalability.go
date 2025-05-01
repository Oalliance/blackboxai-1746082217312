package main

import (
	"fmt"
	"sync"
	"time"
)

// Sidechain represents a simple sidechain structure
type Sidechain struct {
	blocks []Block
	mutex  sync.Mutex
	name   string
}

// NewSidechain creates a new sidechain instance
func NewSidechain(name string) *Sidechain {
	return &Sidechain{
		blocks: []Block{},
		name:   name,
	}
}

// AddBlock adds a block to the sidechain
func (sc *Sidechain) AddBlock(data string) {
	sc.mutex.Lock()
	defer sc.mutex.Unlock()

	var prevHash string
	if len(sc.blocks) > 0 {
		prevHash = sc.blocks[len(sc.blocks)-1].Hash
	}

	newBlock := Block{
		Index:     len(sc.blocks),
		Timestamp: time.Now(),
		Data:      data,
		PrevHash:  prevHash,
		Nonce:     0,
	}
	newBlock = mineBlock(newBlock)
	sc.blocks = append(sc.blocks, newBlock)
	fmt.Printf("Sidechain %s: Added block %d with hash %s\n", sc.name, newBlock.Index, newBlock.Hash)
}

// SyncWithMainChain simulates syncing sidechain state with main blockchain
func (sc *Sidechain) SyncWithMainChain(mainChain *Blockchain) {
	sc.mutex.Lock()
	defer sc.mutex.Unlock()

	// For demo, just print syncing info
	fmt.Printf("Sidechain %s syncing with main chain. Sidechain blocks: %d, Main chain blocks: %d\n", sc.name, len(sc.blocks), len(mainChain.GetBlocks()))
	// Real implementation would include cross-chain communication and state verification
}
