package main

import (
	"log"
	"net/http"
	"os"

	"objectstorage/chapter6/dataServer/heartbeat"
	"objectstorage/chapter6/dataServer/locate"
	"objectstorage/chapter6/dataServer/objects"
	"objectstorage/chapter6/dataServer/temp"
)

func main() {
	locate.CollectObjects()
	go heartbeat.StartHeartbeat()
	go locate.StartLocate()
	http.HandleFunc("/objects/", objects.Handler)
	http.HandleFunc("/temp/", temp.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
