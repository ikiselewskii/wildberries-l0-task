package main

import (
	"wildberries-l0-task/cache/api"
	"wildberries-l0-task/cache/cache"
	"wildberries-l0-task/cache/database"
	"wildberries-l0-task/cache/nats"
	"log"
)

func main(){
	log.Println("LOG Hello I`m up and running")
	cache.Initialize()
	database.Connect()
	err := database.GetAllOrders(); if err != nil{
		log.Panicf("ERROR can`t fill cache from db due to %s", err)
	}
	go nats.Subscribe()
	handler.Initialize()
}

