package main

import (
	"fmt"
	"time"
)

func runWatcher() {
	fmt.Println("Run docker containers watcher ...")

	ticker := time.NewTicker(3 * time.Second)

	for range ticker.C {
		fmt.Println("")
	}
}
