package main

import (
	"flag"
	"fmt"
	"reflect"
	"strings"

	"github.com/fatih/color"
)

// Flags holds all the configuration options
type Flags struct {
	datadog      bool
	debug        bool
	dnshost      string
	dnsserver    string
	gateway      string
	log          bool
	metric       bool
	model        string
	noloop       bool
	sleep        int
	statsdipport string
}

func returnFlags(colorMode bool) (Flags, *bool) {
	datadog := flag.Bool("datadog", false, "Send metric data to datadog (default false)")
	debug := flag.Bool("debug", false, "Enable debug mode to log all results (default false)")
	dnshost := flag.String("dnshost", "google.com", "The hostname to look up")
	dnsserver := flag.String("dnsserver", "8.8.8.8", "The DNS server's IPv4 address to use")
	gateway := flag.String("gateway", "192.168.1.254", "The gateway IPv4 address to compare against)")
	metric := flag.Bool("metric", false, "Output as a metric instead of as a message (default false)")
	model := flag.String("model", "bgw320505", "The model name of the gateway")
	noloop := flag.Bool("noloop", false, "Disable the loop and run the check only once (default false)")
	sleep := flag.Int("sleep", 10, "The time in seconds to sleep between each check")
	statsdipport := flag.String("statsdipport", "127.0.0.1:8125", "Statsd IP port")
	version := flag.Bool("version", false, "Show version")

	flag.Usage = func() {
		usage(colorMode)
	}

	// Parse flags
	flag.Parse()

	log := true

	if *datadog {
		*metric = true
		log = false
	}

	// Create a Flags instance
	flags := Flags{
		datadog:      *datadog,
		debug:        *debug,
		dnshost:      *dnshost,
		dnsserver:    *dnsserver,
		gateway:      *gateway,
		log:          log,
		metric:       *metric,
		model:        *model,
		noloop:       *noloop,
		sleep:        *sleep,
		statsdipport: *statsdipport,
	}

	return flags, version
}

func usage(colorMode bool) {
	// Define Sprintf functions based on colorMode
	var blueSprintf, boldGreenSprintf, cyanSprintf, greenSprintf func(format string, a ...interface{}) string

	if colorMode {
		blueSprintf = color.New(color.FgBlue).Sprintf
		boldGreenSprintf = color.New(color.FgGreen, color.Bold).Sprintf
		cyanSprintf = color.New(color.FgCyan).Sprintf
		greenSprintf = color.New(color.FgGreen).Sprintf
	} else {
		blueSprintf = fmt.Sprintf
		boldGreenSprintf = fmt.Sprintf
		cyanSprintf = fmt.Sprintf
		greenSprintf = fmt.Sprintf
	}

	applicationNameVersion := returnApplicationNameVersion()
	fmt.Println(greenSprintf(applicationNameVersion))

	fmt.Print(greenSprintf("\nUsage:\n"))

	flag.VisitAll(func(f *flag.Flag) {
		// Format flag name with color
		s := "  " + boldGreenSprintf("-%s", f.Name)

		// Get the type of the flag's value
		flagType := reflect.TypeOf(f.Value).Elem().Name()
		flagType = strings.TrimSuffix(flagType, "Value")
		s += blueSprintf(" %s", flagType)

		// Add default value if it exists
		if f.DefValue != "" {
			s += blueSprintf(" (default: %v)", f.DefValue)
		}

		// Add the usage description
		if f.Usage != "" {
			s += "\n    \t" + cyanSprintf(f.Usage)
		}

		fmt.Println(s)
	})
}
