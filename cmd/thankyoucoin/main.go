package main

import (
	"fmt"
	"os"

	"github.com/kalibroida/ThankyouCoin_Node/cmd/thankyoucoin/launcher"
)

func main() {
	if err := launcher.Launch(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
