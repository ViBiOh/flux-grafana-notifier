# flux-notifier

[![Build](https://github.com/ViBiOh/flux-notifier/workflows/Build/badge.svg)](https://github.com/ViBiOh/flux-notifier/actions)
[![codecov](https://codecov.io/gh/ViBiOh/flux-notifier/branch/main/graph/badge.svg)](https://codecov.io/gh/ViBiOh/flux-notifier)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=ViBiOh_flux-notifier&metric=alert_status)](https://sonarcloud.io/dashboard?id=ViBiOh_flux-notifier)

## Getting started

Golang binary is built with static link. You can download it directly from the [Github Release page](https://github.com/ViBiOh/flux-notifier/releases) or build it by yourself by cloning this repo and running `make`.

A Docker image is available for `amd64`, `arm` and `arm64` platforms on Docker Hub: [vibioh/flux-notifier](https://hub.docker.com/r/vibioh/flux-notifier/tags).

You can configure app by passing CLI args or environment variables (cf. [Usage](#usage) section). CLI override environment variables.

You'll find a Kubernetes exemple in the [`infra/`](infra/) folder, using my [`app chart`](https://github.com/ViBiOh/charts/tree/main/app)

## CI

Following variables are required for CI:

|            Name            |           Purpose           |
| :------------------------: | :-------------------------: |
|      **DOCKER_USER**       | for publishing Docker image |
|      **DOCKER_PASS**       | for publishing Docker image |
| **SCRIPTS_NO_INTERACTIVE** |  for running scripts in CI  |

## Usage

```bash
Usage of flux-notifier:
  -address string
        [server] Listen address {FLUX_NOTIFIER_ADDRESS}
  -cert string
        [server] Certificate file {FLUX_NOTIFIER_CERT}
  -corsCredentials
        [cors] Access-Control-Allow-Credentials {FLUX_NOTIFIER_CORS_CREDENTIALS}
  -corsExpose string
        [cors] Access-Control-Expose-Headers {FLUX_NOTIFIER_CORS_EXPOSE}
  -corsHeaders string
        [cors] Access-Control-Allow-Headers {FLUX_NOTIFIER_CORS_HEADERS} (default "Content-Type")
  -corsMethods string
        [cors] Access-Control-Allow-Methods {FLUX_NOTIFIER_CORS_METHODS} (default "GET")
  -corsOrigin string
        [cors] Access-Control-Allow-Origin {FLUX_NOTIFIER_CORS_ORIGIN} (default "*")
  -csp string
        [owasp] Content-Security-Policy {FLUX_NOTIFIER_CSP} (default "default-src 'self'; base-uri 'self'")
  -frameOptions string
        [owasp] X-Frame-Options {FLUX_NOTIFIER_FRAME_OPTIONS} (default "deny")
  -graceDuration string
        [http] Grace duration when SIGTERM received {FLUX_NOTIFIER_GRACE_DURATION} (default "30s")
  -grafanaAddress string
        [grafana] Address {FLUX_NOTIFIER_GRAFANA_ADDRESS} (default "http://grafana")
  -grafanaPassword string
        [grafana] Password for auth {FLUX_NOTIFIER_GRAFANA_PASSWORD}
  -grafanaUsername string
        [grafana] Username for auth {FLUX_NOTIFIER_GRAFANA_USERNAME}
  -hsts
        [owasp] Indicate Strict Transport Security {FLUX_NOTIFIER_HSTS} (default true)
  -idleTimeout string
        [server] Idle Timeout {FLUX_NOTIFIER_IDLE_TIMEOUT} (default "2m")
  -key string
        [server] Key file {FLUX_NOTIFIER_KEY}
  -loggerJson
        [logger] Log format as JSON {FLUX_NOTIFIER_LOGGER_JSON}
  -loggerLevel string
        [logger] Logger level {FLUX_NOTIFIER_LOGGER_LEVEL} (default "INFO")
  -loggerLevelKey string
        [logger] Key for level in JSON {FLUX_NOTIFIER_LOGGER_LEVEL_KEY} (default "level")
  -loggerMessageKey string
        [logger] Key for message in JSON {FLUX_NOTIFIER_LOGGER_MESSAGE_KEY} (default "message")
  -loggerTimeKey string
        [logger] Key for timestamp in JSON {FLUX_NOTIFIER_LOGGER_TIME_KEY} (default "time")
  -okStatus int
        [http] Healthy HTTP Status code {FLUX_NOTIFIER_OK_STATUS} (default 204)
  -port uint
        [server] Listen port {FLUX_NOTIFIER_PORT} (default 1080)
  -prometheusAddress string
        [prometheus] Listen address {FLUX_NOTIFIER_PROMETHEUS_ADDRESS}
  -prometheusCert string
        [prometheus] Certificate file {FLUX_NOTIFIER_PROMETHEUS_CERT}
  -prometheusIdleTimeout string
        [prometheus] Idle Timeout {FLUX_NOTIFIER_PROMETHEUS_IDLE_TIMEOUT} (default "10s")
  -prometheusIgnore string
        [prometheus] Ignored path prefixes for metrics, comma separated {FLUX_NOTIFIER_PROMETHEUS_IGNORE}
  -prometheusKey string
        [prometheus] Key file {FLUX_NOTIFIER_PROMETHEUS_KEY}
  -prometheusPort uint
        [prometheus] Listen port {FLUX_NOTIFIER_PROMETHEUS_PORT} (default 9090)
  -prometheusReadTimeout string
        [prometheus] Read Timeout {FLUX_NOTIFIER_PROMETHEUS_READ_TIMEOUT} (default "5s")
  -prometheusShutdownTimeout string
        [prometheus] Shutdown Timeout {FLUX_NOTIFIER_PROMETHEUS_SHUTDOWN_TIMEOUT} (default "5s")
  -prometheusWriteTimeout string
        [prometheus] Write Timeout {FLUX_NOTIFIER_PROMETHEUS_WRITE_TIMEOUT} (default "10s")
  -readTimeout string
        [server] Read Timeout {FLUX_NOTIFIER_READ_TIMEOUT} (default "5s")
  -shutdownTimeout string
        [server] Shutdown Timeout {FLUX_NOTIFIER_SHUTDOWN_TIMEOUT} (default "10s")
  -url string
        [alcotest] URL to check {FLUX_NOTIFIER_URL}
  -userAgent string
        [alcotest] User-Agent for check {FLUX_NOTIFIER_USER_AGENT} (default "Alcotest")
  -writeTimeout string
        [server] Write Timeout {FLUX_NOTIFIER_WRITE_TIMEOUT} (default "10s")
```
