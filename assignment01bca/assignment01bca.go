// assignment01bca.go
package assignment01bca

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
)

// Block represents each 'item' in the blockchain
type Block struct {
	Transaction  string
	Nonce        int
	PreviousHash string
	Hash         string
}

// Blockchain is a series of validated Blocks
var Blockchain []*Block

// CalculateHash calculates the SHA256 hash of a string
func CalculateHash(stringToHash string) string {
	hash := sha256.Sum256([]byte(stringToHash))
	return hex.EncodeToString(hash[:])
}

// CreateHash generates the hash for a block
func (b *Block) CreateHash() {
	data := b.Transaction + strconv.Itoa(b.Nonce) + b.PreviousHash
	b.Hash = CalculateHash(data)
}

// NewBlock creates a new block with the provided transaction, nonce, and previous hash
func NewBlock(transaction string, nonce int, previousHash string) *Block {
	block := &Block{
		Transaction:  transaction,
		Nonce:        nonce,
		PreviousHash: previousHash,
	}
	block.CreateHash()
	Blockchain = append(Blockchain, block)
	fmt.Println("New block added:")
	fmt.Printf("Transaction: %s, Nonce: %d, Previous Hash: %s, Hash: %s\n\n",
		block.Transaction, block.Nonce, block.PreviousHash, block.Hash)
	return block
}

// ListBlocks prints all the blocks in the blockchain
func ListBlocks() {
	fmt.Println("Listing all blocks in the blockchain:")
	for i, block := range Blockchain {
		fmt.Printf("Block %d:\n", i)
		fmt.Printf("Transaction: %s\n", block.Transaction)
		fmt.Printf("Nonce: %d\n", block.Nonce)
		fmt.Printf("Previous Hash: %s\n", block.PreviousHash)
		fmt.Printf("Hash: %s\n\n", block.Hash)
	}
}

// ChangeBlock changes the transaction data of the block at the given index
func ChangeBlock(index int, newTransaction string) {
	if index < len(Blockchain) && index >= 0 {
		fmt.Printf("Changing transaction of Block %d to '%s'\n", index, newTransaction)
		Blockchain[index].Transaction = newTransaction
		Blockchain[index].CreateHash()
		fmt.Println("Block hash updated due to transaction change.\n")
	} else {
		fmt.Println("Invalid block index\n")
	}
}

// VerifyChain verifies the integrity of the blockchain
func VerifyChain() {
	fmt.Println("Verifying the blockchain:")
	for i := 1; i < len(Blockchain); i++ {
		previousBlock := Blockchain[i-1]
		currentBlock := Blockchain[i]

		if currentBlock.PreviousHash != previousBlock.Hash {
			fmt.Printf("Blockchain invalid at block %d: Previous hash does not match.\n\n", i)
			return
		}

		data := currentBlock.Transaction + strconv.Itoa(currentBlock.Nonce) + currentBlock.PreviousHash
		recalculatedHash := CalculateHash(data)
		if currentBlock.Hash != recalculatedHash {
			fmt.Printf("Blockchain invalid at block %d: Hash does not match.\n\n", i)
			return
		}
	}
	fmt.Println("Blockchain is valid.\n")
}
