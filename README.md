# att-fiber-gateway-outage-detection

## Description
A cross platform golang program to check for Internet outages via DNS for
devices behind AT&T Fiber gateways.

## Note
This is intended to detect outages where the AT&T Fiber gateway does not reboot,
because it has to be booted to return DNS queries.

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
  -debug
        Enable debug mode to log all results (default: false)
  -dnsserver string
        The DNS server's IPv4 address to use (default "8.8.8.8")
  -gateway string
        The gateway's IPv4 address to compare against (default "192.168.1.254")
  -hostname string
        The hostname to look up (default "google.com")
  -noloop
        Disable the loop and run the check only once (default: false)
  -sleep int
        The time in seconds to sleep between each check (default 10)
```

## Example of usage
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

# Proof
[Here](PROOF.md) is a link to another file with examples of the output from
tcpdump showing the interception during an outage.
