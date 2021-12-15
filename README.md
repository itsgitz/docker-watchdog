# Docker Watchdog

Docker Watchdog is `Go` application that used for detects all stopped containers that have `exited`, `paused`, or `dead` state/status.
It uses [list containers endpoints](https://docs.docker.com/engine/api/v1.41/#operation/ContainerList) from
[Docker Engine API] (https://docs.docker.com/engine/api/v1.41/#).
The watcher will gather all docker containers information repeatedly every 3 seconds using `go Ticker` function.
If there are stopped containers detected, `Docker Watchdog` will send an email alert to Developer or System Administrator.

# Getting Started

## Prerequisite
* Go Version: `go1.17.2 linux/amd64`

## Usage

1. Docker Watchdog creates the default configuration file on `/opt/` directory called `.docker-watchdog.yaml`
```shell
$ cd docker-watchdog
$ go build
$ ./docker-watchdog run
```

2. Also, you can specify your own configuration file path with option `--config my-docker-watchdog.yaml`
```shell
$ ./docker-watchdog run --config my-docker-watchdog.yaml
```

# Todo

* Use redis for store email cache. Email cache is used for checking containers ID of the latest email after sended to recipient.
If containers ID is same with the latest email, don't send alert email again. It's prevent email spaming.

# Contributor

Developer Team - PT Infinys System Indonesia 2022
