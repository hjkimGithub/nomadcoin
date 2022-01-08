package main

import (
	"github.com/hjkimGithub/nomadcoin/explorer"
	"github.com/hjkimGithub/nomadcoin/rest"
)

func main() {
	go explorer.Start(5000)
	rest.Start(4000)
}
