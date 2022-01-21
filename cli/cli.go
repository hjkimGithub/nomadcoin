package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/hjkimGithub/nomadcoin/explorer"
	"github.com/hjkimGithub/nomadcoin/rest"
)

func usage() {
	fmt.Printf("Welcome to NOMADCOIN\n")
	fmt.Printf("Please use the following commands\n")
	fmt.Printf("-port:	Set Port(default: 4000)\n")
	fmt.Printf("-mode:	Set mode(Option: 'html' or 'rest' or 'both'\n\n")
	os.Exit(0)
}

func Start() {
	if len(os.Args) == 1 {
		usage()
	}

	port := flag.Int("port", 4000, "Set Port(Default: 4000)")
	// port2 := flag.Int("port", 9000, "Set Port(Default: 9000)")
	mode := flag.String("mode", "rest", "Set mode(Option: 'rest', 'html', 'both'")

	flag.Parse()

	switch *mode {
	case "rest":
		// start rest api
		rest.Start(*port)
	case "html":
		// start html explorer
		explorer.Start(*port)
	default:
		usage()
	}
}
