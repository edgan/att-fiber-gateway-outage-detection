package main

import (
	"fmt"
	"os"
	"time"
)

func logFatal(msg interface{}) {
	var output string
	switch v := msg.(type) {
	case string:
		output = v
	case error:
		output = v.Error()
	default:
		output = "Unknown error"
	}

	fmt.Fprintf(os.Stderr, "Error: %s\n", output)
	os.Exit(1)
}

func logFatalf(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
	os.Exit(1)
}

func logWithTimestamp(message string) {
        timestamp := time.Now().Format("2006-01-02 15:04:05")
        fmt.Printf("[%s] %s\n", timestamp, message)
}
