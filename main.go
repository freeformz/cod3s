package main

import (
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		pathParts := strings.Split(r.URL.Path, "/")
		statusCode, err := strconv.Atoi(pathParts[len(pathParts)-1])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(statusCode)
	})

	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
