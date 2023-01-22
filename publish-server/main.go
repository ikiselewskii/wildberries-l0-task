package main

import (
	"log"
	"github.com/nats-io/stan.go"
)

func main(){
	  // Connect to NATS Streaming server
    nc, err := stan.Connect("test-cluster", "client-publisher", stan.NatsURL("nats://nats:4222"))
    if err != nil {
        log.Fatalf("Error connecting to NATS Streaming: %v", err)
    }
    defer nc.Close()

    // Send message to channel
    if err := nc.Publish("wb-channel", []byte("Hello, NATS Streaming!")); err != nil {
        log.Fatalf("Error sending message: %v", err)
    }

    log.Println("Message sent to channel.")
}