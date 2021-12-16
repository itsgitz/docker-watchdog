package watchdog

import "time"

func (w *Watchdog) watcher() {
	w.initWatcher()

	ticker := time.NewTicker(3 * time.Second)

	for range ticker.C {
		//Get all stopped containers information
		w.getStoppedContainers()
	}
}

func (w *Watchdog) initWatcher() {

	//Before run the watcher, remove the cache file first
	removeCache()

	//Check config...
	//If config is not exist, create new config
	w.checkConfig()

	SuccessText.Printf("[*] Run the Docker Watchdog \n")
}
