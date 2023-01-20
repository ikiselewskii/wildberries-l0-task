package cache
import "wildberries-l0-task/cache/models"


type storage map[string]models.Order

var cache storage

func New() *storage{
	cache = make(storage)
	return &cache
}

func GetOrderByUID(uid string) models.Order{
	return cache[uid]
}

func Add (order models.Order) {
	cache[order.Order_uid] = order
}