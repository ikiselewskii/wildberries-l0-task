package main

import (
	"log"
	"github.com/nats-io/stan.go"
    "encoding/json"
    "github.com/google/uuid"
)

type delivery struct{
	Name string `json:"name"`
	Phone string `json:"phone"`
	Zip string `json:"zip"`
	City string `json:"city"`
	Address string `json:"address"`
	Region string `json:"region"`
	Email string `json:"email"`
}

type item struct{
	Chrt_id int `json:"chrt_id"`
	Track_number string `json:"track_number"`
	Price int `json:"price"`
    Rid string `json:"rid"`
	Name string `json:"name"`
	Sale int `json:"sale"`
	Size string `json:"size"`
	Total_price int `json:"total_price"`
	Nm_id int `json:"nm_id"`
	Brand string `json:"brand"`
	Status int `json:"status"` 
}

type payment struct{
	Transaction string `json:"transaction"`
	Request_id string `json:"request_id"`
	Currency string `json:"currency"`
	Provider string `json:"provider"`
	Amount int `json:"amount"`
	Payment_dt int `json:"payment_dt"`
	Bank string `json:"bank"`
	Delivery_cost int `json:"delivery_cost"`
	Goods_total int `json:"goods_total"`
	Custom_fee int `json:"custom_fee"`
}

type order struct{
	Order_uid string `json:"order_uid"`
	Track_number string `json:"track_number"`
	Entry string `json:"entry"`
	Delivery delivery `json:"delivery"`
	Payment payment `json:"payment"`
	Items []item `json:"items"`
	Locale string `json:"locale"`
	Internal_signature string `json:"internal_signature"`
	Customer_id string `json:"customer_id"`
	Delivery_service string `json:"delivery_service"`
	Shardkey string `json:"shardkey"`
	Sm_id int `json:"sm_id"`
	Date_created string `json:"date_created"`
	Oof_shard string `json:"oof_shard"`

}


func main(){
    log.Println("LOG Hello I`m up and running")
    nc, err := stan.Connect("test-cluster", "client-publisher", stan.NatsURL("nats://nats:4222"))
    if err != nil {
        log.Fatalf("Error connecting to NATS Streaming: %v", err)
    }
    defer nc.Close()

    for i := 0; i < 10; i++{
        data,err := json.Marshal( &order{
        Order_uid: uuid.NewString(),
        Track_number: "WBILMTESTTRACK",
        Entry: "WBIL",
        Delivery: delivery{
            Name: "Test Testov",
            Phone: "+9720000000",
            Zip: "2639809",
            City: "Kiryat Mozkin",
            Address: "Ploshad Mira 15",
            Region: "Kraiot",
            Email: "test@gmail.com",
        },
        Payment: payment{
            Transaction: "b563feb7b2b84b6test",
            Request_id: "",
            Currency: "USD",
            Provider: "wbpay",
            Amount: 1817,
            Payment_dt: 1637907727,
            Bank: "alpha",
            Delivery_cost: 1500,
            Goods_total: 317,
            Custom_fee: 0,
        },
        Items: []item{{
            Chrt_id: 9934930,
            Track_number: "WBILMTESTTRACK",
            Price: 453,
            Rid: "ab4219087a764ae0btest",
            Name: "Mascaras",
            Sale: 30,
            Size: "0",
            Total_price: 317,
            Nm_id: 2389212,
            Brand: "Vivienne Sabo",
            Status: 202,
        },
        },
        Locale: "en",
        Internal_signature: "",
        Customer_id: "test",
        Delivery_service: "meest",
        Shardkey: "9",
        Sm_id: 99,
        Date_created: "2021-11-26T06:22:19Z",
        Oof_shard: "1",
        })
        if err != nil{
            panic(err)
        }
        if err := nc.Publish("wb-channel", data); err != nil {
            log.Fatalf("Error sending message: %v", err)
        }

    log.Println("Message sent to channel.")
    }
}