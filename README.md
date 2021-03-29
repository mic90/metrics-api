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

### Counter

# Storage model

# API

App provides self hosted swagger docs at `/swagger/index.html` endpoint

[http://127.0.0.1:8080/swagger/index.html](http://127.0.0.1:8080/swagger/index.html)
