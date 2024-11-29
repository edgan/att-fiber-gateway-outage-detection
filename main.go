package main

import (
	"fmt"
	"os"
)

func main() {
	colorMode := checkColorTerminal()
	flags, version := returnFlags(colorMode)

	if *version {
		fmt.Println(returnApplicationNameVersion())
		os.Exit(0)
	}

	checkDNS(flags)
}
