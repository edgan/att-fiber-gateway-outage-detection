package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/miekg/dns"
)

func main() {
	debug := flag.Bool("debug", false, "Enable debug mode to log all results (default false)")
	dnsserver := flag.String("dnsserver", "8.8.8.8", "The DNS server's IPv4 address to use")
	gateway := flag.String("gateway", "192.168.1.254", "The gateway IPv4 address to compare against)")
	hostname := flag.String("hostname", "google.com", "The hostname to look up")
	noloop := flag.Bool("noloop", false, "Disable the loop and run the check only once (default false)")
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
			outageDetected := processDNSResponse(response, *gateway, *debug)
			if outageDetected && *noloop {
				os.Exit(1) // Exit with return code 1 only if noloop is set
			}
		}

		// Exit if noloop is set
		if *noloop {
			break
		}

		// Adjust sleep duration to ensure consistent intervals
		elapsed := time.Since(start)
		sleepDuration := time.Duration(*sleep)*time.Second - elapsed
		if sleepDuration > 0 {
			time.Sleep(sleepDuration)
		}
	}
}

func processDNSResponse(response *dns.Msg, gateway string, debug bool) bool {
	found := false
	for _, answer := range response.Answer {
		if aRecord, ok := answer.(*dns.A); ok {
			found = true
			ip := aRecord.A.String()
			if ip == gateway {
				logWithTimestamp(fmt.Sprintf("Outage detected: %s", ip))
				return true
			} else if debug {
				logWithTimestamp(fmt.Sprintf("No outage detected: %s", ip))
			}
		}
	}

	if !found {
		logWithTimestamp("No A records found.")
	}
	return false
}

func logWithTimestamp(message string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("[%s] %s\n", timestamp, message)
}
