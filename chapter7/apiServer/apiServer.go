package main

import (
	"log"
	"net/http"
	"os"

	"objectstorage/chapter7/apiServer/heartbeat"
	"objectstorage/chapter7/apiServer/locate"
	"objectstorage/chapter7/apiServer/objects"
	"objectstorage/chapter7/apiServer/temp"
	"objectstorage/chapter7/apiServer/versions"
)

func main() {
	go heartbeat.ListenHeartbeat()
	http.HandleFunc("/objects/", objects.Handler)
	http.HandleFunc("/temp/", temp.Handler)
	http.HandleFunc("/locate/", locate.Handler)
	http.HandleFunc("/versions/", versions.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
