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

# Build the docker watchdog application from source
echo "[*] Installing docker-watchdog for linux"
echo "[*] Building go application from source"
/usr/local/go/bin/go build

# output
usr_bin_dir="/usr/bin/" 
go_output="docker-watchdog"

echo "[*] Moving $go_output to $usr_bin_dir"
mv $go_output $usr_bin_dir

# Check if docker-watchdog.service is exist
# If not exist, create copy to systemd directory:
# /etc/systemd/system/
docker_watchdog_service_file="$go_output.service"
systemd_dir="/etc/systemd/system/"

if [ ! -f "$systemd_dir/$docker_watchdog_service_file" ]; then	
	echo "[*] Copiying $docker_watchdog_service_file to $systemd_dir"
	cp ./linux/$docker_watchdog_service_file $systemd_dir
fi

echo "[*] Reload systemd"
systemctl daemon-reload

echo "[*] Finish :)"
