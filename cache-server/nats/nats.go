package nats

import (
	"encoding/json"
	"log"
	"wildberries-l0-task/cache/cache"
	"wildberries-l0-task/cache/database"
	"wildberries-l0-task/cache/models"

	"github.com/nats-io/stan.go"
)

func Subscribe() {
	sc, err := stan.Connect("test-cluster", "client-cacher", stan.NatsURL("nats://nats:4222"))
	if err != nil {
		log.Fatalf("ERROR connecting to NATS Streaming: %v", err)
	}
	defer sc.Close()

	_, err = sc.Subscribe("wb-channel", handleMessage, stan.DurableName("my-durable"), stan.SetManualAckMode())
	if err != nil {
		log.Fatalf("ERROR subscribing to channel: %v", err)
	}

	select {}

}

func handleMessage(msg *stan.Msg) {
	log.Printf("Received message: %s", string(msg.Data))
	var order models.Order
	if err := json.Unmarshal(msg.Data, &order); err != nil {
		log.Printf("ERROR Unmarshalling Went Wrong: %s", err)
		msg.Ack()
		return
	}
	cache.Add(&order)
	err := database.AddOrder(order.Order_uid, msg.Data)
	if err != nil{
		log.Printf("ERROR Can`t add into a database")
		msg.Ack()
		return
	}
	msg.Ack()
}
