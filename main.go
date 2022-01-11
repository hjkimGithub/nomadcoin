package main

import (
	"github.com/hjkimGithub/nomadcoin/cli"
	"github.com/hjkimGithub/nomadcoin/db"
)

func main() {
	defer db.Close()
	cli.Start()
}
