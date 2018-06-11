# Benchmarks

In this document you will find benchmark results for different endpoints of ORY Hydra. All benchmarks are executed
using [rakyll/hey](https://github.com/rakyll/hey). Please note that these benchmarks run against the in-memory storage
adapter of ORY Hydra. These benchmarks represent what performance you would get with a zero-overhead database implementation.

We do not include benchmarks against databases (e.g. MySQL or PostgreSQL) as the performance greatly differs between
deployments (e.g. request latency, database configuration) and tweaking individual things may greatly improve performance.
We believe, for that reason, that benchmark results for these database adapters are difficult to generalize and potentially
deceiving. They are thus not included.

This file is updated on every push to master. It thus represents the benchmark data for the latest version.

All benchmarks run 10.000 requests in total, with 100 concurrent requests. All benchmarks run on Circle-CI with a
["2 CPU cores and 4GB RAM"](https://support.circleci.com/hc/en-us/articles/360000489307-Why-do-my-tests-take-longer-to-run-on-CircleCI-than-locally-)
configuration.

## BCrypt

ORY Hydra uses BCrypt to obfuscate secrets of OAuth 2.0 Clients. When using flows such as the OAuth 2.0 Client Credentials
Grant, ORY Hydra validates the client credentials using BCrypt which causes (by design) CPU load. CPU load and performance
depend on the BCrypt cost which can be set using the environment variable `BCRYPT_COST`. For these benchmarks,
we have set `BCRYPT_COST=8`.

## OAuth 2.0

This section contains various benchmarks against OAuth 2.0 endpoints

### Token Introspection

```

Summary:
  Total:	1.2732 secs
  Slowest:	0.0829 secs
  Fastest:	0.0002 secs
  Average:	0.0122 secs
  Requests/sec:	7854.0238
  
  Total data:	1550000 bytes
  Size/request:	155 bytes

Response time histogram:
  0.000 [1]	|
  0.009 [3395]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.017 [4353]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.025 [1850]	|■■■■■■■■■■■■■■■■■
  0.033 [278]	|■■■
  0.042 [53]	|
  0.050 [36]	|
  0.058 [17]	|
  0.066 [7]	|
  0.075 [2]	|
  0.083 [8]	|


Latency distribution:
  10% in 0.0018 secs
  25% in 0.0062 secs
  50% in 0.0130 secs
  75% in 0.0164 secs
  90% in 0.0206 secs
  95% in 0.0241 secs
  99% in 0.0358 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0002 secs, 0.0829 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0124 secs
  req write:	0.0001 secs, 0.0000 secs, 0.0142 secs
  resp wait:	0.0115 secs, 0.0002 secs, 0.0673 secs
  resp read:	0.0003 secs, 0.0000 secs, 0.0142 secs

Status code distribution:
  [200]	10000 responses



```

### Client Credentials Grant

This endpoint uses [BCrypt](#bcrypt).

```

Summary:
  Total:	24.7131 secs
  Slowest:	1.0859 secs
  Fastest:	0.0229 secs
  Average:	0.2405 secs
  Requests/sec:	404.6444
  
  Total data:	1570000 bytes
  Size/request:	157 bytes

Response time histogram:
  0.023 [1]	|
  0.129 [1532]	|■■■■■■■■■■■■■■■
  0.235 [4129]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.342 [2837]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.448 [991]	|■■■■■■■■■■
  0.554 [347]	|■■■
  0.661 [110]	|■
  0.767 [42]	|
  0.873 [6]	|
  0.980 [3]	|
  1.086 [2]	|


Latency distribution:
  10% in 0.1094 secs
  25% in 0.1721 secs
  50% in 0.2177 secs
  75% in 0.2966 secs
  90% in 0.3841 secs
  95% in 0.4513 secs
  99% in 0.5998 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0229 secs, 1.0859 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0101 secs
  req write:	0.0002 secs, 0.0000 secs, 0.0709 secs
  resp wait:	0.2396 secs, 0.0227 secs, 1.0858 secs
  resp read:	0.0005 secs, 0.0000 secs, 0.1273 secs

Status code distribution:
  [200]	10000 responses



```

## OAuth 2.0 Client Management

### Creating OAuth 2.0 Clients

This endpoint uses [BCrypt](#bcrypt) and generates IDs and secrets by reading from  which negatively impacts
performance. Performance will be better if IDs and secrets are set in the request as opposed to generated by ORY Hydra.

```
This test is currently disabled due to issues with /dev/urandom being inaccessible in the CI.
```

### Listing OAuth 2.0 Clients

```

Summary:
  Total:	0.7328 secs
  Slowest:	0.0499 secs
  Fastest:	0.0002 secs
  Average:	0.0071 secs
  Requests/sec:	13645.5871
  
  Total data:	2670000 bytes
  Size/request:	267 bytes

Response time histogram:
  0.000 [1]	|
  0.005 [3655]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.010 [4299]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.015 [1519]	|■■■■■■■■■■■■■■
  0.020 [286]	|■■■
  0.025 [89]	|■
  0.030 [102]	|■
  0.035 [20]	|
  0.040 [10]	|
  0.045 [13]	|
  0.050 [6]	|


Latency distribution:
  10% in 0.0012 secs
  25% in 0.0034 secs
  50% in 0.0060 secs
  75% in 0.0096 secs
  90% in 0.0125 secs
  95% in 0.0154 secs
  99% in 0.0279 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0002 secs, 0.0499 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0106 secs
  req write:	0.0001 secs, 0.0000 secs, 0.0139 secs
  resp wait:	0.0033 secs, 0.0001 secs, 0.0341 secs
  resp read:	0.0020 secs, 0.0000 secs, 0.0183 secs

Status code distribution:
  [200]	10000 responses



```

### Fetching a specific OAuth 2.0 Client

```

Summary:
  Total:	0.7515 secs
  Slowest:	0.0376 secs
  Fastest:	0.0002 secs
  Average:	0.0073 secs
  Requests/sec:	13306.7166
  
  Total data:	2650000 bytes
  Size/request:	265 bytes

Response time histogram:
  0.000 [1]	|
  0.004 [2950]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.008 [2622]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.011 [2906]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.015 [1005]	|■■■■■■■■■■■■■■
  0.019 [246]	|■■■
  0.023 [124]	|■■
  0.026 [59]	|■
  0.030 [35]	|
  0.034 [45]	|■
  0.038 [7]	|


Latency distribution:
  10% in 0.0009 secs
  25% in 0.0032 secs
  50% in 0.0067 secs
  75% in 0.0104 secs
  90% in 0.0125 secs
  95% in 0.0153 secs
  99% in 0.0249 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0002 secs, 0.0376 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0070 secs
  req write:	0.0001 secs, 0.0000 secs, 0.0122 secs
  resp wait:	0.0042 secs, 0.0001 secs, 0.0348 secs
  resp read:	0.0016 secs, 0.0000 secs, 0.0213 secs

Status code distribution:
  [200]	10000 responses



```