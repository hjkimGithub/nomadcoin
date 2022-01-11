package main

import (
	"github.com/hjkimGithub/nomadcoin/blockchain"
	"github.com/hjkimGithub/nomadcoin/cli"
)

func main() {
	blockchain.BlockChain()
	cli.Start()
}
