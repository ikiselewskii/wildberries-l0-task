package cache

import (
	"log"
	"sync"
	"wildberries-l0-task/cache/models"
)


type storage struct {
	cache map[string]*models.Order
	mx sync.Mutex
}

var cache storage

func Initialize() *storage{
	allocation := make(map[string]*models.Order)
	cache = storage{cache : allocation,}
	log.Print("LOG Cache initialized")
	return &cache
}

func GetOrderByID(id string) (*models.Order, bool){
	cache.mx.Lock()
	order,exists := cache.cache[id]
	cache.mx.Unlock()
	return order, exists
}

func Add (order *models.Order) {
	cache.mx.Lock()
	cache.cache[order.Order_uid] = order
	cache.mx.Unlock()
}