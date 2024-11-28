# Docker
[Here](/Dockerfile) is a Dockerfile for running this in a Linux container.

## Building the image
```
docker build . -t att-fiber-gateway-outage-detection
```

# DockHub reepository
https://hub.docker.com/repository/docker/edgan/att-fiber-gateway-outage-detection/general

## Styles
There are two styles of usage, daemon and command line.

Daemon uses `docker/att-fiber-gateway-outage-detection.sh` to pass in values for
flags. It also assumes you want `-datadog` as a default flag.

Command line lets you just directly specify all the flags directly.

### Daemon environment variables for flag values
```
DNSHOST
DNSSERVER
GATEWAY
MODEL
SLEEP
STATSDIPPORT
```

### Daemon usage
```
docker run -it -e MODEL='<model>' edgan/att-fiber-gateway-outage-detection:1.0.6
docker run -it -e GATEWAY='<ip>' edgan/att-fiber-gateway-outage-detection:1.0.6
docker run -it -e STATSDIPPORT='<ip>:<port>' edgan/att-fiber-gateway-outage-detection:1.0.6
```

### Daemon examples
```
docker run -it -e MODEL='bgw210700' edgan/att-fiber-gateway-outage-detection:1.0.6
docker run -it -e GATEWAY='192.168.10.1' edgan/att-fiber-gateway-outage-detection:1.0.6
docker run -it -e STATSDIPPORT=192.168.10.10:8125' edgan/att-fiber-gateway-outage-detection:1.0.6
```


### Command line usage
```
docker run -it edgan/att-fiber-gateway-outage-detection:1.0.6 att-fiber-gateway-outage-detection -model <model>
docker run -it edgan/att-fiber-gateway-outage-detection:1.0.6 att-fiber-gateway-outage-detection -gateway <ip>
docker run -it edgan/att-fiber-gateway-outage-detection:1.0.6 att-fiber-gateway-outage-detection -statsdipport <ip:port>
```

### Command line examples
```
docker run -it edgan/att-fiber-gateway-outage-detection:1.0.6 att-fiber-gateway-outage-detection -model bgw210700
docker run -it edgan/att-fiber-gateway-outage-detection:1.0.6 att-fiber-gateway-outage-detection -gateway 192.168.10.1
docker run -it edgan/att-fiber-gateway-outage-detection:1.0.6 att-fiber-gateway-outage-detection -statsdipport 192.168.10.10:8125
```
