# notifier

[![Build](https://github.com/ViBiOh/notifier/workflows/Build/badge.svg)](https://github.com/ViBiOh/notifier/actions)
[![codecov](https://codecov.io/gh/ViBiOh/notifier/branch/main/graph/badge.svg)](https://codecov.io/gh/ViBiOh/notifier)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=ViBiOh_notifier&metric=alert_status)](https://sonarcloud.io/dashboard?id=ViBiOh_notifier)

## Getting started

Golang binary is built with static link. You can download it directly from the [Github Release page](https://github.com/ViBiOh/notifier/releases) or build it by yourself by cloning this repo and running `make`.

A Docker image is available for `amd64`, `arm` and `arm64` platforms on Docker Hub: [vibioh/notifier](https://hub.docker.com/r/vibioh/notifier/tags).

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

The application can be configured by passing CLI args described below or their equivalent as environment variable. CLI values take precedence over environments variables.

Be careful when using the CLI values, if someone list the processes on the system, they will appear in plain-text. Pass secrets by environment variables: it's less easily visible.

```bash
Usage of notifier:
  -address string
        [server] Listen address {NOTIFIER_ADDRESS}
  -cert string
        [server] Certificate file {NOTIFIER_CERT}
  -corsCredentials
        [cors] Access-Control-Allow-Credentials {NOTIFIER_CORS_CREDENTIALS}
  -corsExpose string
        [cors] Access-Control-Expose-Headers {NOTIFIER_CORS_EXPOSE}
  -corsHeaders string
        [cors] Access-Control-Allow-Headers {NOTIFIER_CORS_HEADERS} (default "Content-Type")
  -corsMethods string
        [cors] Access-Control-Allow-Methods {NOTIFIER_CORS_METHODS} (default "GET")
  -corsOrigin string
        [cors] Access-Control-Allow-Origin {NOTIFIER_CORS_ORIGIN} (default "*")
  -csp string
        [owasp] Content-Security-Policy {NOTIFIER_CSP} (default "default-src 'self'; base-uri 'self'")
  -discordWebhookURL string
        [discord] Discord Webhook WebhookURL {NOTIFIER_DISCORD_WEBHOOK_URL} (default "WebhookURL")
  -fibrSecret string
        [fibr] Webhook Secret {NOTIFIER_FIBR_SECRET}
  -frameOptions string
        [owasp] X-Frame-Options {NOTIFIER_FRAME_OPTIONS} (default "deny")
  -graceDuration string
        [http] Grace duration when SIGTERM received {NOTIFIER_GRACE_DURATION} (default "30s")
  -grafanaAddress string
        [grafana] Grafana Address {NOTIFIER_GRAFANA_ADDRESS} (default "http://grafana")
  -grafanaPassword string
        [grafana] Grafana Basic Auth Password {NOTIFIER_GRAFANA_PASSWORD}
  -grafanaUsername string
        [grafana] Grafana Basic Auth Username {NOTIFIER_GRAFANA_USERNAME}
  -hsts
        [owasp] Indicate Strict Transport Security {NOTIFIER_HSTS} (default true)
  -idleTimeout string
        [server] Idle Timeout {NOTIFIER_IDLE_TIMEOUT} (default "2m")
  -key string
        [server] Key file {NOTIFIER_KEY}
  -loggerJson
        [logger] Log format as JSON {NOTIFIER_LOGGER_JSON}
  -loggerLevel string
        [logger] Logger level {NOTIFIER_LOGGER_LEVEL} (default "INFO")
  -loggerLevelKey string
        [logger] Key for level in JSON {NOTIFIER_LOGGER_LEVEL_KEY} (default "level")
  -loggerMessageKey string
        [logger] Key for message in JSON {NOTIFIER_LOGGER_MESSAGE_KEY} (default "message")
  -loggerTimeKey string
        [logger] Key for timestamp in JSON {NOTIFIER_LOGGER_TIME_KEY} (default "time")
  -mailerName string
        [mailer] HTTP Username or AMQP Exchange name {NOTIFIER_MAILER_NAME} (default "mailer")
  -mailerPassword string
        [mailer] HTTP Pass {NOTIFIER_MAILER_PASSWORD}
  -mailerURL string
        [mailer] URL (https?:// or amqps?://) {NOTIFIER_MAILER_URL}
  -okStatus int
        [http] Healthy HTTP Status code {NOTIFIER_OK_STATUS} (default 204)
  -port uint
        [server] Listen port (0 to disable) {NOTIFIER_PORT} (default 1080)
  -prometheusAddress string
        [prometheus] Listen address {NOTIFIER_PROMETHEUS_ADDRESS}
  -prometheusCert string
        [prometheus] Certificate file {NOTIFIER_PROMETHEUS_CERT}
  -prometheusGzip
        [prometheus] Enable gzip compression of metrics output {NOTIFIER_PROMETHEUS_GZIP}
  -prometheusIdleTimeout string
        [prometheus] Idle Timeout {NOTIFIER_PROMETHEUS_IDLE_TIMEOUT} (default "10s")
  -prometheusIgnore string
        [prometheus] Ignored path prefixes for metrics, comma separated {NOTIFIER_PROMETHEUS_IGNORE}
  -prometheusKey string
        [prometheus] Key file {NOTIFIER_PROMETHEUS_KEY}
  -prometheusPort uint
        [prometheus] Listen port (0 to disable) {NOTIFIER_PROMETHEUS_PORT} (default 9090)
  -prometheusReadTimeout string
        [prometheus] Read Timeout {NOTIFIER_PROMETHEUS_READ_TIMEOUT} (default "5s")
  -prometheusShutdownTimeout string
        [prometheus] Shutdown Timeout {NOTIFIER_PROMETHEUS_SHUTDOWN_TIMEOUT} (default "5s")
  -prometheusWriteTimeout string
        [prometheus] Write Timeout {NOTIFIER_PROMETHEUS_WRITE_TIMEOUT} (default "10s")
  -readTimeout string
        [server] Read Timeout {NOTIFIER_READ_TIMEOUT} (default "5s")
  -shutdownTimeout string
        [server] Shutdown Timeout {NOTIFIER_SHUTDOWN_TIMEOUT} (default "10s")
  -sshRecipient string
        [ssh] SSH Notification recipient {NOTIFIER_SSH_RECIPIENT}
  -sshSender string
        [ssh] SSH Notification sender {NOTIFIER_SSH_SENDER}
  -url string
        [alcotest] URL to check {NOTIFIER_URL}
  -userAgent string
        [alcotest] User-Agent for check {NOTIFIER_USER_AGENT} (default "Alcotest")
  -writeTimeout string
        [server] Write Timeout {NOTIFIER_WRITE_TIMEOUT} (default "10s")
```
