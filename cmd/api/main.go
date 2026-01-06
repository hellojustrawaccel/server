package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("starting api on :8080")

	http.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("ok"))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
