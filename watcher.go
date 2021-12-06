package main

import (
	"fmt"
	"time"
)

func runWatcher() {
	fmt.Printf("Run docker containers watcher ... \n\n")

	ticker := time.NewTicker(3 * time.Second)

	for range ticker.C {
		//Get all stopped containers information
		getStoppedContainers()
	}
}
