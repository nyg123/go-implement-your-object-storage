package main

import (
	"log"
	"net/http"
	"os"

	"objectstorage/chapter5/apiServer/heartbeat"
	"objectstorage/chapter5/apiServer/locate"
	"objectstorage/chapter5/apiServer/objects"
	"objectstorage/chapter5/apiServer/versions"
)

func main() {
	go heartbeat.ListenHeartbeat()
	http.HandleFunc("/objects/", objects.Handler)
	http.HandleFunc("/locate/", locate.Handler)
	http.HandleFunc("/versions/", versions.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
