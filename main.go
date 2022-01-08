package main

import (
	"github.com/hjkimGithub/nomadcoin/explorer"
	"github.com/hjkimGithub/nomadcoin/rest"
)

func main() {
	explorer.Start(3000)
	rest.Start(4000)
}
