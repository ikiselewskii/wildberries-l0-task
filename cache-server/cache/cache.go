package cache
import (
	"wildberries-l0-task/cache/models"
	"sync"
)


type storage struct {
	cache map[int]*models.Order
	mx sync.Mutex
}

var cache storage

func New() *storage{
	cache := make(map[int]*models.Order)
	return &storage{
		cache : cache,
	}
}

func GetOrderByID(id int) *models.Order{
	cache.mx.Lock()
	order := cache.cache[id]
	cache.mx.Unlock()
	return order
}

func Add (order *models.Order) {
	cache.mx.Lock()
	cache.cache[order.ID] = order
	cache.mx.Unlock()
}