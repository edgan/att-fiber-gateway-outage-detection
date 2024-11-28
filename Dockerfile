FROM ubuntu:24.04

COPY docker/att-fiber-gateway-outage-detection.sh /usr/local/bin/att-fiber-gateway-outage-detection.sh
COPY bin/att-fiber-gateway-outage-detection_linux_amd64 /usr/local/bin/att-fiber-gateway-outage-detection

CMD ["att-fiber-gateway-outage-detection.sh"]
