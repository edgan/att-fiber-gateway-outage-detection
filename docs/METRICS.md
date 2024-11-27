# Metric
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
In [datadog/dashboards/](datadog/dashboards) there are `json` files broken out by `model`, and
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
In [datadog/monitors/](datadog/monitors) there is are example `json` files. They have an example
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
