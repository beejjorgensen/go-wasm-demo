package main

import (
	"log"
	"net/http"
)

func logf(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func main() {
	log.Printf("Listening on port 8000")

	fileServer := http.FileServer(http.Dir("."))
	log.Fatal(http.ListenAndServe(":8000", logf(fileServer)))
}
