package main

import (
	"time"
)

func runWatcher() {
	successText.Printf("[*] --- Docker Watchdog --- \n")
	successText.Printf("[*] Run docker containers watcher ... \n")

	ticker := time.NewTicker(3 * time.Second)

	for range ticker.C {
		//Get all stopped containers information
		getStoppedContainers()
	}
}
