FROM ubuntu

COPY bin/att-fiber-gateway-outage-detection_linux_amd64 /usr/local/bin/att-fiber-gateway-outage-detection
CMD ["att-fiber-gateway-outage-detection", "-datadog"]
