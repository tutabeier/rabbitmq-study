package main

import (
	"log"
	"net/http"

	"github.com/streadway/amqp"
)

func producer(w http.ResponseWriter, r *http.Request) {
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
	w.Write([]byte("producer"))
}

func main() {
	log.Print("producer started")
	http.HandleFunc("/", producer)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
