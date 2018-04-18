package main

import (
	"github.com/fsnotify/fsnotify"
	"log"
)

type Klient struct {
	NazwaKlienta string
	SciezkaDoFolderuUzytkownika   string
	SerwerPlikowURL string
}


func (klient Klient) Sledz() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()
	if err = watcher.Add(klient.SciezkaDoFolderuUzytkownika); err != nil {
		log.Fatal(err)
	}

	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				klient.odpowiedzNaZdarzenie(event)
			case err := <-watcher.Errors:
				log.Println("error:", err)
			}
		}
	}()
	<-done
}

func (k Klient) odpowiedzNaZdarzenie(event fsnotify.Event){
	if event.Op&fsnotify.Create == fsnotify.Create {
		go przeslijNaSerwerPlikow(k.SerwerPlikowURL, k.SciezkaDoFolderuUzytkownika + "/" + event.Name)
		log.Println("created file file: ", event.Name)
	}else if event.Op&fsnotify.Remove == fsnotify.Remove {
		log.Println("removed file file: ", event.Name)
	}else if event.Op&fsnotify.Write == fsnotify.Write {
		log.Println("modified file: ", event.Name)
	}else if event.Op&fsnotify.Rename == fsnotify.Rename {
		log.Println("renamed file: ", event.Name)
	}
}
