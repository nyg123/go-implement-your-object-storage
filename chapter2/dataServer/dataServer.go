package main

import (
	"log"
	"net/http"
	"os"

	"objectstorage/chapter2/apiServer/objects"
	"objectstorage/chapter2/dataServer/heartbeat"
	"objectstorage/chapter2/dataServer/locate"
)

func main() {
	go heartbeat.StartHeartbeat()
	go locate.StartLocate()
	http.HandleFunc("/objects/", objects.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
