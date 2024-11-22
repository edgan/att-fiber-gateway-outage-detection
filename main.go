package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/miekg/dns"
)

func main() {
	// Define flags with default values
	debug := flag.Bool("debug", false, "Enable debug mode to log all results (default false)")
	dnsserver := flag.String("dnsserver", "8.8.8.8", "The DNS server's IPv4 address to use")
	gateway := flag.String("gateway", "192.168.1.254", "The gateway IPv4 address to compare against")
	hostname := flag.String("hostname", "google.com", "The hostname to look up")
	sleep := flag.Int("sleep", 10, "The time in seconds to sleep between each check")

	// Parse flags
	flag.Parse()

	for {
		start := time.Now() // Track the start time of the iteration

		// Prepare the DNS query
		client := new(dns.Client)
		message := new(dns.Msg)
		message.SetQuestion(dns.Fqdn(*hostname), dns.TypeA)

		// Perform the DNS query
		response, _, err := client.Exchange(message, *dnsserver+":53")
		if err != nil {
			logWithTimestamp(fmt.Sprintf("Error performing DNS lookup: %v", err))
		} else {
			// Process the results
			if response.Rcode != dns.RcodeSuccess {
				logWithTimestamp(fmt.Sprintf("DNS query failed with Rcode: %d", response.Rcode))
			} else {
				processDNSResponse(response, *gateway, *debug)
			}
		}

		// Adjust sleep duration to ensure consistent intervals
		elapsed := time.Since(start)
		sleepDuration := time.Duration(*sleep)*time.Second - elapsed
		if sleepDuration > 0 {
			time.Sleep(sleepDuration)
		}
	}
}

// processDNSResponse handles the processing of the DNS response
func processDNSResponse(response *dns.Msg, gateway string, debug bool) {
	found := false
	for _, answer := range response.Answer {
		if aRecord, ok := answer.(*dns.A); ok {
			found = true
			ip := aRecord.A.String()
			if ip == gateway {
				logWithTimestamp(fmt.Sprintf("Outage detected: %s", ip))
			} else if debug {
				logWithTimestamp(fmt.Sprintf("No outage detected: %s", ip))
			}
		}
	}

	if !found {
		logWithTimestamp("No A records found.")
	}
}

// logWithTimestamp prints a message prefixed with the current timestamp
func logWithTimestamp(message string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("[%s] %s\n", timestamp, message)
}
