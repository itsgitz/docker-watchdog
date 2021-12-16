#!/usr/bin/env bash

# # Create user for docker-watchdog application
# username="docker-watchdog"

# if ! id -u $username; then
# 	echo "[*] Creating user ..."
# 	useradd $username -g sudo
# else
# 	echo "[*] User already exist"
# 	echo "[*] Continue ..."
# fi

## GO BUILD
# Build the docker watchdog application from source
echo "[*] Installing docker-watchdog for linux"
echo "[*] Building go application from source"
/usr/local/go/bin/go build

# output
usr_bin_dir="/usr/bin/" 
application="docker-watchdog"

echo "[*] Moving $application to $usr_bin_dir"
mv $application $usr_bin_dir

## DOCKER WATCHDOG SERVICE

# Check if docker-watchdog.service is exist
# If not exist, create copy to systemd directory:
# /etc/systemd/system/
docker_watchdog_service_file="$application.service"
systemd_dir="/etc/systemd/system/"

if [ ! -f "$systemd_dir/$docker_watchdog_service_file" ]; then	
	echo "[*] Copiying $docker_watchdog_service_file to $systemd_dir"
	cp ./linux/$docker_watchdog_service_file $systemd_dir
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
