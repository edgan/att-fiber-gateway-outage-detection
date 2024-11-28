# Helm
`att-fiber-gatway-outage-detection` can be run in Kubernetes via the helm chart.

## Values
* flags.dnshost
  * If you want to use a hostname to be checked other than the default of
google.com
* flags.dnsserver
  * If you want to use a dns server ip address other than 8.8.8.8
* flags.gateway
  * If your AT&T Fiber gateway's ip address is something other than the default
of 192.168.1.254
* flags.model
  * If your model is not bgw320505
  * Alternatives are bgw210700 and bgw320500
* flags.sleep
  * if you want to change the default of reporting every 10 seconds
* flags.statsdipport
  * If your statsd ip:port is something other than the default of 127.0.0.1:8125
  * Note, 127.0.0.1 in the container is not the same as 127.0.0.1 on the host

## Usage
```
cd helm
helm package .
helm install att-fiber-gateway-outage-detection att-fiber-gateway-outage-detection-0.1.0.tgz -f values.yaml
```

## Checking on it
```
kubectl get pods -l app=att-fiber-gateway-outage-detection
kubectl get pods -l app=att-fiber-gateway-outage-detection -o yaml
kubectl logs -l app=att-fiber-gateway-outage-detection
```
