package main

import (
	"log"
	"net/http"
	"os"

	"objectstorage/chapter8/apiServer/heartbeat"
	"objectstorage/chapter8/apiServer/locate"
	"objectstorage/chapter8/apiServer/objects"
	"objectstorage/chapter8/apiServer/temp"
	"objectstorage/chapter8/apiServer/versions"
)

func main() {
	go heartbeat.ListenHeartbeat()
	http.HandleFunc("/objects/", objects.Handler)
	http.HandleFunc("/temp/", temp.Handler)
	http.HandleFunc("/locate/", locate.Handler)
	http.HandleFunc("/versions/", versions.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
