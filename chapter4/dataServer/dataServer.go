package main

import (
	"log"
	"net/http"
	"os"

	"objectstorage/chapter4/dataServer/heartbeat"
	"objectstorage/chapter4/dataServer/locate"
	"objectstorage/chapter4/dataServer/objects"
	"objectstorage/chapter4/dataServer/temp"
)

func main() {
	locate.CollectObjects()
	go heartbeat.StartHeartbeat()
	go locate.StartLocate()
	http.HandleFunc("/objects/", objects.Handler)
	http.HandleFunc("/temp/", temp.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
