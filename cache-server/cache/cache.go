package cache
import (
	"wildberries-l0-task/cache/models"
	"sync"
)


type storage struct {
	cache map[string]*models.Order
	mx sync.Mutex
}

var cache storage

func Initialize() *storage{
	cache := make(map[string]*models.Order)
	return &storage{
		cache : cache,
	}
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