# Metrics-API

[![codecov](https://codecov.io/gh/mic90/metrics-api/branch/master/graph/badge.svg?token=8RS6X1XORO)](https://codecov.io/gh/mic90/metrics-api)
[![Build Status](https://travis-ci.com/mic90/metrics-api.svg?token=P2rqpx8Ukgp6jqhbA9Bu&branch=master)](https://travis-ci.com/mic90/metrics-api)

# Build
```bash
docker build -t metrics-api .
```
Should produce an app image:
```
REPOSITORY    TAG       IMAGE ID       CREATED         SIZE
metrics-api   latest    d49373b9ddb8   2 seconds ago   28.5MB
```

# Run
```bash
docker run -p 8080:8080 metrics-api:latest
```

# Available metrics

### Gauge

A gauge is a metric that represents a single numerical value that can arbitrarily go up and down.

Gauges are typically used for measured values like temperatures or current memory usage, but also "counts" that can go up and down, like the number of concurrent requests.

### Counter
A counter is a cumulative metric that represents a single monotonically increasing counter whose value can only increase or be reset to zero on restart. 
For example, you can use a counter to represent the number of requests served, tasks completed, or errors.

# Storage model

To increase data retrieval performance, metric data is split into smaller chunks, grouped by time ranges.

The chunks are kept in sorted array, so when one index is found all consecutive data chunks can be retrieved by iterating that array.

Additional lookup map is used to find chunks indexes, based on provided time ranges.

# API

There is self hosted swagger viewer at `/swagger/index.html` [endpoint](http://127.0.0.1:8080/swagger/index.html)
