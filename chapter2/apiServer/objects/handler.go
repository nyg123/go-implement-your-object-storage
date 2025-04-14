package objects

import (
	"net/http"

	objects2 "objectstorage/chapter2/objects"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	m := r.Method
	if m == http.MethodPut {
		objects2.put(w, r)
		return
	}
	if m == http.MethodGet {
		objects2.get(w, r)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}
