#!/usr/bin/env bash

echo "[*] Uninstalling docker-watchdog"
echo "[*] Disabling the docker-watchdog service ..."

# Disable docker-watchdog.service
systemctl stop docker-watchdog
systemctl disable docker-watchdog

# Remove docker-watchdog executable file from /usr/bin
echo "[*] Removing /usr/bin/docker-watchdog ..."
rm /usr/bin/docker-watchdog

# Remove docker-watchdog.service from /etc/systemd/system/docker-watchdog.service
echo "[*] Removing /etc/systemd/system/docker-watchdog.service"
rm /etc/systemd/system/docker-watchdog.service

echo "[*] Removing configuration file /opt/.docker-watchdog.yaml"
rm /opt/.docker-watchdog.yaml

echo "[*] Finish :)"
