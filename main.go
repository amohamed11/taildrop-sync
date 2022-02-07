package main

import (
	"fmt"
	"log"
	"net"

	"github.com/fsnotify/fsnotify"
)

func lookUpDeviceName(deviceName string) (string, error) {
	if addrs, err := net.LookupHost(deviceName); err != nil {
		return "", fmt.Errorf("error looking up IP of %q: %v", deviceName, err)
	} else if len(addrs) == 0 {
		return "", fmt.Errorf("no IPs found for %q", deviceName)
	} else {
		return addrs[0], nil
	}
}

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add("/root/Sync/kin")
	if err != nil {
		log.Fatal(err)
	}
	<-done
}
