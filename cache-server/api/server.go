package server

import (
	"encoding/json"
	"net/http"
	"wildberries-l0-task/cache/cache"
)

func handleGetByUid(rw http.ResponseWriter, r *http.Response){
	if r.Request.Method != http.MethodGet{
		rw.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	id := r.Request.URL.Query().Get("id")
	order, exists := cache.GetOrderByID(id)
	if !exists{
		rw.WriteHeader(http.StatusNotFound)
		rw.Write([]byte("Order not found"))
	}
	data, err := json.Marshal(order)
	if err != nil{
		rw.WriteHeader(http.StatusInternalServerError)
	}
	rw.WriteHeader(http.StatusOK)
	rw.Write(data)
}