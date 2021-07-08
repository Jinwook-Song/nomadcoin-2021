package main

import (
	"fmt"

	"github.com/Jinwook-Song/nomadcoin-2021/blockchain"
)


func main() {
	chain := blockchain.GetBlockchain()
	fmt.Println(chain)
}

