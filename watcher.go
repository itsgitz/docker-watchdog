package main

import (
	"time"
)

func runWatcher() {
	successText.Printf("[*] --- Docker Watchdog --- \n")
	successText.Printf("[*] Run docker containers watcher ... \n")

	//Run as background with goroutine
	go watcher()

	select {}
}

func watcher() {
	ticker := time.NewTicker(3 * time.Second)

	for range ticker.C {
		//Get all stopped containers information
		getStoppedContainers()
	}
}
