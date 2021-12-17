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

# Check if docker-watchdog is already installed on system
# Application name and bin directory
usr_bin_dir="/usr/bin/" 
application="docker-watchdog"

echo ""

if [ -f "$usr_bin_dir$application" ]; then
	echo "[!] The docker-watchdog is already installed"
	exit 0
else
	echo "[*] The docker-watchdog is not found, continue installation"
fi

## GO BUILD
# Build the docker watchdog application from source
echo "[*] Installing docker-watchdog for linux"
echo "[*] Building go application from source"

# You also can use go build command
#/usr/local/go/bin/go build

# With docker:
docker run -it --rm \
	--name docker_watchdog_build \
	-v "$(pwd):/go/src/docker-watchdog" \
	-w /go/src/docker-watchdog \
	golang \
	go build

echo "[*] Moving $application to $usr_bin_dir"
mv $application $usr_bin_dir

## DOCKER WATCHDOG SERVICE

# Check if docker-watchdog.service is exist
# If not exist, create copy to systemd directory:
# /etc/systemd/system/
docker_watchdog_service_file="$application.service"
systemd_dir="/etc/systemd/system/"

if [ ! -f "$systemd_dir/$docker_watchdog_service_file" ]; then	
	echo "[*] Copiying ./linux/systemd/$docker_watchdog_service_file to $systemd_dir"
	cp ./linux/systemd/$docker_watchdog_service_file $systemd_dir
fi

# DOCKER WATCHDOG CONFIGURATION
default_configuration_path="/opt"
configuration_file=".docker-watchdog.yaml"
config_path=$default_configuration_path/$configuration_file
example_config="./example/conf/.docker-watchdog.example.yaml"

if [ ! -f "$config_path" ]; then
	echo "[*] Copiying example configuration to $config_path"
	cp $example_config $config_path
else
	echo "[*] Configuration is already exist"
	echo "[*] Skip copiying ..."
fi

# ENABLE SYSTEMCTL
echo "[*] Reload systemd"
systemctl daemon-reload

echo "[*] Enabling docker-watchdog service"
systemctl enable $application
systemctl restart $application

echo "[*] Please setup your docker-watchdog configuration on /opt/.docker-watchdog.yaml"
echo "[*] Finish :)"

exit 0
