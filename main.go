package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/miekg/dns"
)

func main() {
	// Define flags with default values
	hostname := flag.String("hostname", "google.com", "The hostname to look up (default: google.com)")
	dnsserver := flag.String("dnsserver", "8.8.8.8", "The DNS server ipv4 address to use (default: 8.8.8.8)")
	gateway := flag.String("gateway", "192.168.1.254", "The gateway ipv4 address to compare against (default: 192.168.1.254)")

	// Parse flags
	flag.Parse()

	// Prepare the DNS query
	client := new(dns.Client)
	message := new(dns.Msg)
	message.SetQuestion(dns.Fqdn(*hostname), dns.TypeA)

	// Perform the DNS query
	response, _, err := client.Exchange(message, *dnsserver+":53")
	if err != nil {
		fmt.Printf("Error performing DNS lookup: %v\n", err)
		os.Exit(1)
	}

	// Process the results
	if response.Rcode != dns.RcodeSuccess {
		fmt.Printf("DNS query failed with Rcode: %d\n", response.Rcode)
		os.Exit(1)
	}

	found := false
	for _, answer := range response.Answer {
		if aRecord, ok := answer.(*dns.A); ok {
			found = true
			ip := aRecord.A.String()
			if ip == *gateway {
				fmt.Printf("Outage detected: %s\n", ip)
				os.Exit(1)
			} else {
				fmt.Println("No outage detected:", ip)
			}
		}
	}

	if !found {
		fmt.Println("No A records found.")
	}
}
