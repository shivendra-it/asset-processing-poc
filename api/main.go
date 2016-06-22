package main

import (
	"github.com/gorilla/mux"
	"github.com/streadway/amqp"
	"log"
	"net/http"
	"os"
)

var (
	rmqConn      *amqp.Connection = nil
	payloadQueue                  = make(chan string)
	rmqServer                     = "amqp://guest:guest@" + os.Getenv("RABBITMQ_SERVER") + ":5672"
)

func init() {
	rmqConn = connect(rmqServer)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/v1/", doWork).Methods("POST")
	log.Fatal(http.ListenAndServe(":80", router))
}
