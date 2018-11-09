package main

import (
	"log"
	"net/http"
)

func producer(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("producer"))
}
func main() {
	log.Print("producer started")
	http.HandleFunc("/", producer)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
