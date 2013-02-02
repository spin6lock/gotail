package main

import (
	"log"
	"github.com/howeyc/fsnotify"
)

func main(){
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	ch := make(chan int)
	go func(){
		for {
			select {
			case ev := <-watcher.Event:
				log.Println("event:", ev)
			case err := <-watcher.Error:
				log.Println("error:", err)
			}
		}
		ch <-1
	}()

	err = watcher.Watch("test.log")
	if err != nil {
		log.Fatal(err)
	}
	<-ch
	watcher.Close()
}
