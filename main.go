package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// Extract a HTTP code from the end of a path.
// e.x. /foo/bar/200 == 200
// If the last part of the path can't convert to an int, then return
// http.StatusBadRequest
func codeFromPath(p string) int {
	pathParts := strings.Split(p, "/")
	statusCode, err := strconv.Atoi(pathParts[len(pathParts)-1])
	if err != nil {
		statusCode = http.StatusBadRequest
	}
	return statusCode
}

func main() {
	log.SetFlags(0)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		statusCode := codeFromPath(r.URL.Path)

		log.Printf("path=%q host=%q remote=%q status=%d\n", r.URL.Path, r.Host, r.RemoteAddr, statusCode)

		w.WriteHeader(statusCode)
	})

	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
