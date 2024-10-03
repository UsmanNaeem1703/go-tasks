package main

import (
	"assignment01bca/assignment01bca"
)

func main() {
	// Create the genesis block
	genesisBlock := assignment01bca.NewBlock("Genesis Block", 0, "0")

	// Add more blocks
	block1 := assignment01bca.NewBlock("Asim pays Ahsan 10 coins", 1, genesisBlock.Hash)
	block2 := assignment01bca.NewBlock("Ahsan pays Usman 5 coins", 2, block1.Hash)
	assignment01bca.NewBlock("Usman pays Faheem 5 coins", 2, block2.Hash)

	// List all blocks
	assignment01bca.ListBlocks()

	// Verify the blockchain
	assignment01bca.VerifyChain()

	// Change a block's transaction
	assignment01bca.ChangeBlock(2, "Ahsan pays Asim 10 coins")

	// List all blocks after the change
	assignment01bca.ListBlocks()

	// Verify the blockchain after the change
	assignment01bca.VerifyChain()
}
