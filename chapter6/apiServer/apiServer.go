package main

import (
	"log"
	"net/http"
	"os"

	"objectstorage/chapter6/apiServer/heartbeat"
	"objectstorage/chapter6/apiServer/locate"
	"objectstorage/chapter6/apiServer/objects"
	"objectstorage/chapter6/apiServer/temp"
	"objectstorage/chapter6/apiServer/versions"
)

func main() {
	go heartbeat.ListenHeartbeat()
	http.HandleFunc("/objects/", objects.Handler)
	http.HandleFunc("/temp/", temp.Handler)
	http.HandleFunc("/locate/", locate.Handler)
	http.HandleFunc("/versions/", versions.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
