package main

// Metric represents a single Datadog metric point
type Metric struct {
	Metric string          `json:"metric"`
	Points [][]interface{} `json:"points"`
	Type   string          `json:"type"`
	Tags   []string        `json:"tags"`
}

// Payload represents the payload to send to Datadog
type Payload struct {
	Series []Metric `json:"series"`
}
