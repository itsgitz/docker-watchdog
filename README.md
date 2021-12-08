# Docker Watchdog

Docker Watchdog is `Go` program that used for detects all stopped containers that have `exited`, `paused`, or `dead` state/status.
It uses [list containers endpoints](https://docs.docker.com/engine/api/v1.41/#operation/ContainerList) from
[Docker Engine API] (https://docs.docker.com/engine/api/v1.41/#).
The watcher will gather all docker containers information repeatedly every 3 seconds using `go Ticker` function.
If there are stopped containers detected, `Docker Watchdog` will send an alert email to Developer or System Administrator.

# Getting Started

* Go Version: `go1.17.2 linux/amd64`

# Usage
```shell
$ cd docker-watchdog
$ mv conf/config_example.yaml conf/config.yaml
$ go build
$ ./docker-watchdog
```

# Todo

* Use redis for store email cache. Email cache is used for checking containers ID of the latest email after sended to recipient.
If containers ID is same with the latest email, don't send alert email again. It's prevent email spaming.

# Contributor

Developer Team - PT Infinys System Indonesia 2022
