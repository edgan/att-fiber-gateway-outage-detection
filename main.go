package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/miekg/dns"
)

func main() {
        datadog := flag.Bool("datadog", false, "Send metric data to datadog (default false)")
	debug := flag.Bool("debug", false, "Enable debug mode to log all results (default false)")
	dnshost := flag.String("dnshost", "google.com", "The hostname to look up")
	dnsserver := flag.String("dnsserver", "8.8.8.8", "The DNS server's IPv4 address to use")
	gateway := flag.String("gateway", "192.168.1.254", "The gateway IPv4 address to compare against)")
	metric := flag.Bool("metric", false, "Output as a metric instead of as a message (default false)")
	model := flag.String("model", "bgw320505", "The model name of the gateway")
	noloop := flag.Bool("noloop", false, "Disable the loop and run the check only once (default false)")
	sleep := flag.Int("sleep", 10, "The time in seconds to sleep between each check")
	statsdipport := flag.String("statsdipport", "127.0.0.1:8125", "Statsd ip port")

	// Parse flags
	flag.Parse()

	log := true

	if *datadog {
		*metric = true
		log = false
	}

	for {
		start := time.Now() // Track the start time of the iteration

		// Prepare the DNS query
		client := new(dns.Client)
		message := new(dns.Msg)
		message.SetQuestion(dns.Fqdn(*dnshost), dns.TypeA)

		// Perform the DNS query
		response, _, err := client.Exchange(message, *dnsserver+":53")
		if err == nil {
			// Process the results
			outageDetected := processDNSResponse(response, *datadog, *debug, *gateway, log, *metric, *model, *statsdipport)
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

func processDNSResponse(response *dns.Msg, datadog bool, debug bool, gateway string, log bool, metric bool, model string, statsdipport string) bool {
	found := false

	metrics := []string{}

        outageMetricName := model + ".outage"
	outageMetricValue := "0.0"

	outageMetric := outageMetricName + "=" + outageMetricValue

	for _, answer := range response.Answer {
		if aRecord, ok := answer.(*dns.A); ok {
			found = true
			ip := aRecord.A.String()
			if ip == gateway {
				outageMetricValue = "1.0"
				outageMetric = outageMetricName + "=" + outageMetricValue

				metrics = append(metrics, outageMetric)

				if datadog {
					giveMetricsToDatadogStatsd(debug, metrics, model, statsdipport)
				}
				if debug || log {
					logWithTimestamp(fmt.Sprintf("Outage detected: %s", ip))
				}

				return true
			} else if debug && !metric {
				logWithTimestamp(fmt.Sprintf("No outage detected: %s", ip))
			}

			if metric && !datadog {
				fmt.Println(outageMetric)
			} else if datadog {
				metrics = append(metrics, outageMetric)
				giveMetricsToDatadogStatsd(debug, metrics, model, statsdipport)
			}
		}
	}

	if !found {
		logWithTimestamp("No A records found.")
	}
	return false
}
