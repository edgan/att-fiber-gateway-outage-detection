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

## Documentation
[Here](docs/) is the documentation.

You can test the behavior and output by setting the gateway flag to the ipv4 of the DNS record used.

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

## Related project
[att-fiber-gateway-info](https://github.com/edgan/att-fiber-gateway-info/)
