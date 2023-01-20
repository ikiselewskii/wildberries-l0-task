package nats

import (
	"github.com/nats-io/stan.go"
	"log"
)


func Subscribe(){
	sc, err := stan.Connect("test-cluster", "client-123", stan.NatsURL("nats://localhost:4222"))
    if err != nil {
        log.Fatalf("Error connecting to NATS Streaming: %v", err)
    }
    defer sc.Close()

    // Subscribe to a channel
    _, err = sc.Subscribe("my-channel", handleMessage, stan.DurableName("my-durable"))
    if err != nil {
        log.Fatalf("Error subscribing to channel: %v", err)
    }

    // Wait for messages
    select {}

}

func handleMessage(msg *stan.Msg){
	log.Printf("Received message: %s", string(msg.Data))
	msg.Ack()
}