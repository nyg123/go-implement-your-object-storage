package main

import (
	"log"
	"net/http"
	"os"

	"objectstorage/chapter2/apiServer/heartbeat"
	"objectstorage/chapter2/apiServer/locate"
	"objectstorage/chapter2/apiServer/objects"
)

func main() {
	go heartbeat.ListenHeartbeat()
	http.HandleFunc("/objects/", objects.Handler)
	http.HandleFunc("/locate/", locate.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
