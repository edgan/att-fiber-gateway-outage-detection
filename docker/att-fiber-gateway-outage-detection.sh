#!/bin/bash

ARGS='-datadog'

if [ ! -z "${DNSHOST}" ]; then
 ARGS="${ARGS} -dnshost ${DNSHOST}"
fi

if [ ! -z "${DNSSERVER}" ]; then
 ARGS="${ARGS} -dnsserver ${DNSSERVER}"
fi

if [ ! -z "${GATEWAY}" ]; then
 ARGS="${ARGS} -gateway ${GATEWAY}"
fi

if [ ! -z "${MODEL}" ]; then
 ARGS="${ARGS} -model ${MODEL}"
fi

if [ ! -z "${SLEEP}" ]; then
 ARGS="${ARGS} -sleep ${SLEEP}"
fi

if [ ! -z "${STATSDIPPORT}" ]; then
 ARGS="${ARGS} -statsdipport ${STATSDIPPORT}"
fi

att-fiber-gateway-outage-detection ${ARGS}
