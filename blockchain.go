package main

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"log"
	"sync"
	"time"
)

// Block represents each 'item' in the blockchain
type Block struct {
	Index        int
	Timestamp    time.Time
	Data         string
	PrevHash     string
	Hash         string
	Nonce        int
}

// Blockchain is a series of validated Blocks
type Blockchain struct {
	blocks []Block
	mutex  sync.RWMutex
}

// NewBlockchain creates a new Blockchain with genesis block
func NewBlockchain() *Blockchain {
	bc := &Blockchain{}
	genesisBlock := Block{
		Index:     0,
		Timestamp: time.Now(),
		Data:      "Genesis Block",
		PrevHash:  "",
		Hash:      "",
		Nonce:     0,
	}
	genesisBlock.Hash = calculateHash(genesisBlock)
	bc.blocks = append(bc.blocks, genesisBlock)
	return bc
}

// AddBlock adds a new block to the blockchain
func (bc *Blockchain) AddBlock(data string) error {
	bc.mutex.Lock()
	defer bc.mutex.Unlock()

	if len(bc.blocks) == 0 {
		return errors.New("blockchain has no genesis block")
	}

	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := Block{
		Index:     prevBlock.Index + 1,
		Timestamp: time.Now(),
		Data:      data,
		PrevHash:  prevBlock.Hash,
		Nonce:     0,
	}
	newBlock = mineBlock(newBlock)
	bc.blocks = append(bc.blocks, newBlock)
	log.Printf("Block %d added with hash %s", newBlock.Index, newBlock.Hash)
	return nil
}

// GetBlocks returns the blockchain blocks
func (bc *Blockchain) GetBlocks() []Block {
	bc.mutex.RLock()
	defer bc.mutex.RUnlock()
	return bc.blocks
}

// calculateHash calculates the hash of a block
func calculateHash(block Block) string {
	record := string(rune(block.Index)) + block.Timestamp.String() + block.Data + block.PrevHash + string(rune(block.Nonce))
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

// mineBlock performs proof-of-work to find a hash with difficulty prefix
func mineBlock(block Block) Block {
	difficulty := 3
	prefix := ""
	for i := 0; i < difficulty; i++ {
		prefix += "0"
	}

	for {
		hash := calculateHash(block)
		if hash[:difficulty] == prefix {
			block.Hash = hash
			break
		}
		block.Nonce++
	}
	return block
}
