package main

import (
	"embed"
	"strings"
)

func returnApplicationName() string {
	applicationName := "att-fiber-gateway-outage-detection"

	return applicationName
}

func returnApplicationNameVersion() string {
	version := returnVersion()
	applicationName := returnApplicationName()
	applicationNameVersion := applicationName + " " + version

	return applicationNameVersion
}

//go:embed .version
var versionFile embed.FS

func returnVersion() string {
	versionBytes, _ := versionFile.ReadFile(".version")
	version := strings.TrimSpace(string(versionBytes))

	return version
}
