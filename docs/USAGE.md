## Usage
```
Usage of ./att-fiber-gateway-outage-detection:
  -datadog
        Send metric data to datadog (default false)
  -debug
        Enable debug mode to log all results (default false)
  -dnsserver string
        The DNS server's IPv4 address to use (default "8.8.8.8")
  -gateway string
        The gateway IPv4 address to compare against) (default "192.168.1.254")
  -hostname string
        The hostname to look up (default "google.com")
  -metric
        Output as a metric instead of as a message (default false)
  -model string
        The model name of the gateway (default "bgw320505")
  -noloop
        Disable the loop and run the check only once (default false)
  -sleep int
        The time in seconds to sleep between each check (default 10)
  -statsdipport string
        Statsd ip port (default "127.0.0.1:8125")
``` 

## Examples of usage
```
./att-fiber-gateway-outage-detection
./att-fiber-gateway-outage-detection -debug
./att-fiber-gateway-outage-detection -debug -noloop
./att-fiber-gateway-outage-detection -dnsserver 1.1.1.1
./att-fiber-gateway-outage-detection -gateway 192.168.1.1
./att-fiber-gateway-outage-detection -hostname facebook.com
./att-fiber-gateway-outage-detection -noloop
./att-fiber-gateway-outage-detection -sleep 20
./att-fiber-gateway-outage-detection -dnsserver 8.8.4.4 -gateway 192.168.10.254
./att-fiber-gateway-outage-detection -dnsserver 1.1.1.1 -hostname apple.com
./att-fiber-gateway-outage-detection -metric
./att-fiber-gateway-outage-detection -metric -model bgw320500
./att-fiber-gateway-outage-detection -datadog
./att-fiber-gateway-outage-detection -datadog -model bgw210700
./att-fiber-gateway-outage-detection -datadog -statsdipport 192.168.1.125:8125
```

## Example output
In case of an outage:
```
[2024-11-21 18:55:23] Outage detected: 192.168.1.254
```

Debug mode:
```
[2024-11-21 18:55:30] No outage detected: 172.217.14.78
[2024-11-21 18:55:40] No outage detected: 172.217.14.78
```

You can test the behavior and output by setting the gateway flag to the ipv4 of the DNS record used.
