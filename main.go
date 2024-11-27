package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/miekg/dns"
)

func main() {
	apiKey := flag.String("apikey", "", "Datadog api key")
	apiURL := flag.String("apiurl", "https://api.datadoghq.com/api/v1/series", "Datadog api url")
	datadog := flag.Bool("datadog", false, "Send metric data to datadog (default false)")
	dnsserver := flag.String("dnsserver", "8.8.8.8", "The DNS server's IPv4 address to use")
	gateway := flag.String("gateway", "192.168.1.254", "The gateway IPv4 address to compare against)")
	hostname := flag.String("hostname", "google.com", "The hostname to look up")
	logLevel := flag.String("loglevel", "error", "Log level")
	metric := flag.Bool("metric", false, "Output as a metric instead of as a message (default false)")
	model := flag.String("model", "bgw320505", "The model name of the gateway")
	noloop := flag.Bool("noloop", false, "Disable the loop and run the check only once (default false)")
	retryDelay := flag.Int("retrydelay", 60, "Number of seconds between retries") // Goal of 1 minute given that outages tend to be 1-2 minutes
	retryLimit := flag.Int("retrylimit", 480, "Number of retries")                // Goal of 8 hours given retryDelay * 60 * 8
	sleep := flag.Int("sleep", 10, "The time in seconds to sleep between each check")

	// Parse flags
	flag.Parse()

	log := true

        if *datadog {
                *metric = true
                log = false
        } else {
		if *apiKey != "" {
			logFatal("-apikey only required for -datadog")
		}
	}

	isLowercaseAlpha := regexp.MustCompile(`^[a-z]+$`)

	if len(*apiKey) != 32 && !isLowercaseAlpha.MatchString(*apiKey) {
		logFatal("-apikey must valid")
	}

	logValue := 1

	*logLevel = strings.ToLower(*logLevel)

	if *logLevel == "debug" {
		logValue = 2
	} else if *logLevel == "trace" {
		logValue = 3
	}

	retryTime := time.Duration(*retryDelay) * time.Second

	for {
		start := time.Now() // Track the start time of the iteration

		// Prepare the DNS query
		client := new(dns.Client)
		message := new(dns.Msg)
		message.SetQuestion(dns.Fqdn(*hostname), dns.TypeA)

		// Perform the DNS query
		response, _, err := client.Exchange(message, *dnsserver+":53")
		if err == nil {
			// Process the results
			outageDetected := processDNSResponse(response, *apiKey, *apiURL, *datadog, logValue, *gateway, log, *metric, *model, retryTime, *retryLimit)
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

func processDNSResponse(response *dns.Msg, apiKey string, apiURL string, datadog bool, logValue int, gateway string, log bool, metric bool, model string, retryTime time.Duration, retryLimit int) bool {
	found := false

	metricName := model + ".outage"
	metricValue := "0.0"

	metricString := metricName + "=" + metricValue
	metricData := makeMetric(model, metricName, metricValue)

	metrics := []Metric{metricData}
	for _, answer := range response.Answer {
		if aRecord, ok := answer.(*dns.A); ok {
			found = true
			ip := aRecord.A.String()
			if ip == gateway {
				if datadog {
					metricValue = "1.0"

					metricString = metricName + "=" + metricValue
					metricData = makeMetric(model, metricName, metricValue)

					metrics = append(metrics, metricData)
					sendMetrics(apiKey, apiURL, logValue, metrics, retryTime, retryLimit)
				}

				if log || logValue >= 2 {
					logWithTimestamp(fmt.Sprintf("Outage detected: %s", ip))
				}

				return true
			} else if logValue >= 2 && !metric {
				logWithTimestamp(fmt.Sprintf("No outage detected: %s", ip))
			}

			if metric && !datadog {
				fmt.Println(metricString)
			} else if datadog {
				metrics = append(metrics, metricData)
				sendMetrics(apiKey, apiURL, logValue, metrics, retryTime, retryLimit)
			}
		}
	}

	if !found {
		logWithTimestamp("No A records found.")
	}
	return false
}
