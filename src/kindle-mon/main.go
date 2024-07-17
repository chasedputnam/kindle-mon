package main

import (
	"log"
	"os"

	"github.com/chasedputnam/kindle-mon/classifier"
	"github.com/chasedputnam/kindle-mon/queue"
	"github.com/fsnotify/fsnotify"
)

var (
	monitoredPath = "/Users/chase.putnam/Downloads/kindle"
	timeout       = 60
)

func initLog() {
	f, err := os.OpenFile("kindle-mon.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatalf("error closing file: %v", err)
		}
	}(f)
	log.SetOutput(f)
}

func main() {
	initLog()
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer func(watcher *fsnotify.Watcher) {
		err := watcher.Close()
		if err != nil {
			log.Fatal("error: ", err)
		}
	}(watcher)

	done := make(chan bool)

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				//Check the type of event and process new ebook files
				log.Printf("event: %v\n", event)
				if (event.Op & fsnotify.Create) == fsnotify.Create {
					log.Printf("created file: %s\n", event.Name)
					requests := classifier.Classify(event.Name)
					queuedRequests := queue.Queue(requests)
					queue.SendMail(queuedRequests, timeout)
				}

			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error: ", err)
			}
		}
	}()

	// Watch the folder for changes
	err = watcher.Add(monitoredPath)
	if err != nil {
		log.Fatal(err)
	}
	<-done
}
