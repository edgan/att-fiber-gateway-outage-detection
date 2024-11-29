package main

import (
	"fmt"

	"github.com/miekg/dns"
)

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
