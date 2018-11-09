package main

import (
	"log"
	"net/http"
)

func consumer(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("consumer"))
}
func main() {
	log.Print("consumer started")
	http.HandleFunc("/", consumer)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
