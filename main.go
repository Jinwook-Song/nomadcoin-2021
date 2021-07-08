package main

import (
	"fmt"

	"github.com/Jinwook-Song/nomadcoin-2021/blockchain"
)


func main() {
	chain := blockchain.GetBlockchain()
	chain.AddBlock("Second Block")
	chain.AddBlock("Third Block")
	chain.AddBlock("Fourth Block")
	for _, block := range chain.AllBlocks() {
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Printf("Prev: %s\n", block.PrevHash)
	}
}

