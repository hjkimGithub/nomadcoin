package main

import (
	"fmt"

	"github.com/hjkimGithub/nomadcoin/blockchain"
)

func main() {
	chain := blockchain.GetBlockChain()
	chain.AddBlock("Second Block")
	chain.AddBlock("Third Block")
	chain.AddBlock("Fourth Block")
	for _, block := range chain.AllBlocks() {
		fmt.Println(block.Data)
		fmt.Println(block.Hash)
		fmt.Println(block.Prevhash)
	}
}
