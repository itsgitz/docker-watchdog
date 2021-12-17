# Docker Watchdog

Docker Watchdog is `Go` application that used for detects all stopped containers that have `exited`, `paused`, or `dead` state/status.
It uses [list containers endpoints](https://docs.docker.com/engine/api/v1.41/#operation/ContainerList) from
[Docker Engine API] (https://docs.docker.com/engine/api/v1.41/#).
The watcher will gather all docker containers information repeatedly every 3 seconds using `go Ticker` function.
If there are stopped containers detected, `Docker Watchdog` will send an email alert to Developer or System Administrator.

# Getting Started

## Prerequisite
* Go Version: `go1.17.2 linux/amd64`
* Docker Engine: `20.10+`

## Installation

**Note**: The installation only tested on linux-like operating system (linux distro)

### A. With `make` command
1. Clone this repository
2. In the root directory, type:

```shell
$ sudo make install
```
3. After installation, the `docker-watchdog` application will run as linux service. But, we have to specify configuration value
on `/opt/.docker-watchdog.yaml`
4. For configuration example you can see on [example/conf/.docker-watchdog.yaml](./example/conf/.docker-watchdog.example.yaml)

### B. Build from source
1. Use `go build` command

```shell
$ go build
```

2. The default executable output named `docker-watchdog`

## Usage

### Run with systemctl
1. Docker Watchdog creates the default configuration file on `/opt/` directory called `.docker-watchdog.yaml`.
If you installing `docker-watchdog` with `make` command, you just simply run:

```shell
$ sudo systemctl start docker-watchdog
```

2. Make sure that the configuration value is correct. See [example configuration](./example/conf/.docker-watchdog.example.yaml)

### A. Run with go executable

1. If you build the application from source code, you can simply run:

```shell
$ ./docker-watchdog run
```

2. Default configuration path is `/opt/.docker-watchdog.yaml`
3. Or you can specify your own configuration file with option `--config` 

```shell
$ ./docker-watchdog run --config my-docker-watchdog.yaml
```


# Todo

* Installation script on linux

# Contributor

Developer Team - PT Infinys System Indonesia 2022
