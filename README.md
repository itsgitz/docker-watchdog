# Docker Watchdog

Docker Watchdog is `Go` program that used for get all stopped containers that have `exited`, `paused`, or `dead` state/status.
It uses [list containers endpoints](https://docs.docker.com/engine/api/v1.41/#operation/ContainerList) from
[Docker Engine API] (https://docs.docker.com/engine/api/v1.41/#).
The watcher will gather all docker containers information repeatedly every 3 seconds using `go Ticker` function.
If there are stopped containers detected, `Docker Watchdog` will send an alert email to Developer or System Administrator.

# Getting Started

* Go Version: `go1.17.2 linux/amd64`

# Usage
```shell
$ cd docker-watchdog
$ go build
$ ./docker-watchdog
```

# Todo

* Send alert email

# Contributor

Developer Team - PT Infinys System Indonesia 2022
