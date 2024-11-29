package main

import (
	"fmt"
	"os"
	"time"

	"github.com/miekg/dns"
)

func checkDNS(flags Flags) {
	for {
		start := time.Now() // Track the start time of the iteration

		// Prepare the DNS query
		client := new(dns.Client)
		message := new(dns.Msg)
		message.SetQuestion(dns.Fqdn(flags.dnshost), dns.TypeA)

		// Perform the DNS query
		response, _, err := client.Exchange(message, flags.dnsserver+":53")
		if err == nil {
			// Process the results
			outageDetected := processDNSResponse(response, flags.datadog, flags.debug, flags.gateway, flags.log, flags.metric, flags.model, flags.statsdipport)
			if outageDetected && flags.noloop {
				os.Exit(1) // Exit with return code 1 only if noloop is set
			}
		} else {
			// Handle DNS query error (optional)
			if flags.debug {
				// Log the error if debug is enabled
				// Replace with your logging mechanism
				fmt.Printf("DNS query error: %v\n", err)
			}
		}

		// Exit if noloop is set
		if flags.noloop {
			break
		}

		// Adjust sleep duration to ensure consistent intervals
		elapsed := time.Since(start)
		sleepDuration := time.Duration(flags.sleep)*time.Second - elapsed
		if sleepDuration > 0 {
			time.Sleep(sleepDuration)
		}
	}
}
