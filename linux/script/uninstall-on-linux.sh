#!/usr/bin/env bash

# Copyright © 2021 PT Infinys System Indonesia <anggit@isi.co.id>
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

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
