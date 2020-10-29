# HTTP latency monitor

This is a minimalist Prometheus exporter, exposing the latency of a
single HTTP target.

Example:

```bash
httplat http://k8s.io/
```

Then open http://localhost:9080/metrics to see the latency.

It will make an HTTP request to the target at regular intervals, and
record the latency using a Prometheus histogram. The interval is
currently hard-coded to 10 seconds. The request timeout is set to
the interval, meaning that any request taking longer than 10 seconds
will be counted as taking 10 seconds.

It doesn't care about the HTTP status code, which means that if the
monitored target serves 404 or 500 errors *very fast*, the exporter
will treat these responses just like `200 OK`. You're warned!

By default, it serves on port 9080, but that can be changed by
setting the `PORT` environment variable.
