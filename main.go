package main

import (
	"github.com/hjkimGithub/nomadcoin/cli"
	"github.com/hjkimGithub/nomadcoin/db"
)

func main() {
	// wallet.Wallet()
	defer db.Close()
	cli.Start()
}
