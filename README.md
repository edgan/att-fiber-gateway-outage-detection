# att-fiber-gateway-outage-detection
![Screenshot1](/screenshots/datadog-outage-dashboard.png)

## Description
A cross platform golang program to check for Internet outages via DNS for
devices behind a AT&T Fiber gateway.

## Features
* Detects outages
* Can generate and report outage metric to
[statsd](https://github.com/statsd/statsd)([Datadog](https://www.datadoghq.com/))
* Custom [Datadog](https://www.datadoghq.com/) dashboards using the metric
* Custom [Datadog](https://www.datadoghq.com/) monitors using the metric for
alerting

## Note
This is intended to detect outages where the AT&T Fiber gateway does not reboot,
because it has to be booted to return DNS queries.

## Supported hardware
* [BGW320-505 gateway](https://help.sonic.com/hc/en-us/articles/1500000066642-BGW320)
* BGW320-500 gateway, expected to work, but untested 
* BGW210-700 gateway, expected to work, but untested

## How it works
By default the AT&T Fiber gateway intercepts A record DNS queries during an
outage. It returns the ipv4 address of the gateway as the answer. So if the ipv4
address returned matches the ipv4 address of the gateway an outage has been
detected. By default it loops to keep detecting outages.

Example:
```
nslookup google.com 8.8.8.8
Server:     8.8.8.8
Address:    8.8.8.8#53

Non-authoritative answer:
Name:    google.com
Address: 192.168.1.254
```

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

You can test the behavior and output by setting the gateway flag to the ipv4 of the DNS record used.

Testing:
```
./att-fiber-gateway-outage-detection -gateway 172.217.14.78
./att-fiber-gateway-outage-detection -gateway 172.217.14.78 -debug
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
## Metric
The `-metric` flag outputs the metric name and value. The metric name includes
the model name.

```
bgw320505.outage=0
bgw320505.outage=1
```

### Model names
The `-model` flag will take any string, but here are some suggested values for
the different models.

```
bgw210700
bgw320500
bgw320505
```

## Datadog
You can use the `-datadog` flag to report the results to statsd. The repeated
default value is `0.0`. When it detects an outage it returns `1.0`.

If you run `att-fiber-gateway-outage-detection` on more than one system you can
see if any of them saw an outage.

### Dashboards
In `datadog/dashboards` there are `json` files broken out by `model`, and
`multi` and `single`.
The `multi` version lets you specify different hostnames
for different hosts running `att-fiber-gateway-outage-detection`. The advantage
of `multi` is that you can set different colors per host. Then you can tell
their values apart.
:8125
The `single` version doesn't specify a host. It can be used with one or multiple
hosts, but if it is multiple the values will be aggregated together.

Dashboard files:
```
BGW210-700-multi.json
BGW210-700-single.json
BGW320-500-multi.json
BGW320-500-single.json
BGW320-505-multi.json
BGW320-505-single.json
```

#### Monitor
In `datadog/monitors` there is are example `json` files. They have an example
`monitor` that takes the metrics and alerts if the sum over five minutes is
greater than zero.

Note the default alerting method is set to `@pagerduty-Home`. This should be
changed to match your configured alerting method.

Monitor files:
```
BGW210-700-outage.json
BGW320-500-outage.json
BGW320-505-outage.json
```

## Compiling
To build for only system's platform:
```
go build
```

To binaries for all supported combinations of operating systems and architectures:
```
scripts/builds.sh
```

## Builds
See the `.go_builds` file for the list of supported combinations of operating
systems and architectures.

## Story
I have been experiencing intermittent outages with my AT&T Fiber. After having
AT&T upgrade my gateway to firmware version `6.30.5` it has been more stable.
Especially with the issue of the gateway rebooting itself. Though that still
isn't perfect.

Yet I continue to see at least one 1-2 minute outage daily. I have monitoring
the network traffic on the outside interface of my personal router, setup in
passthrough mode, with tcpdump. During an outage I could clearly see the
gateway intercepting DNS queries for A records.

I had been testing a script and program to detect outages by detecting it
intercepting web traffic. This worked, but was more error prone and complex.
This method is much simpler, and should always work.

## Proof
[Here](PROOF.md) is a link to another file with examples of the output from
tcpdump showing the interception during an outage.
