package main

import (
	"log"
	"net/http"
	"os"

	"objectstorage/chapter3/apiServer/objects"
	"objectstorage/chapter3/dataServer/heartbeat"
	"objectstorage/chapter3/dataServer/locate"
)

func main() {
	go heartbeat.StartHeartbeat()
	go locate.StartLocate()
	http.HandleFunc("/objects/", objects.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
