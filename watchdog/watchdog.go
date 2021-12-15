package watchdog

import (
	"time"
)

//Watchdog data collections (struct)
type Watchdog struct {
	Config *Config
}

//Run the Docker Watchdog application
func (w *Watchdog) Run() {
	//Run in the background with goroutine
	go w.watcher()

	select {}
}

func (w *Watchdog) watcher() {
	w.initWatchDog()

	ticker := time.NewTicker(3 * time.Second)

	for range ticker.C {
		//Get all stopped containers information
		w.getStoppedContainers()
	}
}

func (w *Watchdog) initWatchDog() {

	//Before run the watcher, remove the cache file first
	removeCache()

	//Check config...
	//If config is not exist, create new config
	w.checkConfig()

	SuccessText.Printf("[*] Run the Docker Watchdog \n")
}
