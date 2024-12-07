// att-fiber-gateway-outage-detection
// A cross platform golang program to check for Internet outages via DNS for devices behind AT&T Fiber gateways.
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
