package main

import (
	"net/http"
	"wildberries-l0-task/cache/database"
	"wildberries-l0-task/cache/nats"
)

func main(){
	go nats.Subscribe()
	database.Connect()
	http.HandleFunc("/ping", pong)
	http.ListenAndServe(":8080", nil)
}

func pong(rw http.ResponseWriter, r *http.Request){
	rw.WriteHeader(http.StatusAccepted)
	rw.Write([]byte("pong"))
}