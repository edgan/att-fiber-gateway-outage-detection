package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func makeMetric(model string, name string, value string) Metric {
	metric := Metric{
		Metric: name,
		Points: [][]interface{}{
			{float64(time.Now().Unix()), value}, // Current timestamp and value
		},
		Type: "gauge",
		Tags: []string{"gateway:" + model},
	}

	return metric
}

func sendMetrics(apiKey string, apiURL string, logValue int, metrics []Metric, retryTime time.Duration, retryLimit int) error {
	payload := Payload{Series: metrics}

	// Serialize the payload to JSON
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	for attempt := 0; attempt < retryLimit; attempt++ {
		// Create the HTTP request
		req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(data))
		if err != nil {
			return err
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("DD-API-KEY", apiKey)

		// Send the request
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Printf("Attempt %d: Error sending metrics: %v", attempt+1, err)
		} else {
			defer resp.Body.Close()
			if resp.StatusCode == http.StatusOK {
				if logValue >= 2 {
					logWithTimestamp("Metrics sent successfully")
				}

				return nil
			} else {
				log.Printf("Attempt %d: Received status code %d", attempt+1, resp.StatusCode)
			}
		}

		// Retry delay
		time.Sleep(retryTime)
	}

	return err
}
