package models

type Order struct{
	Order_uid string `json:"order_uid"`
	Track_number string `json:"track_number"`
	Entry string `json:"entry"`
	Delivery delivery `json:"delivery"`
	Payment payment `json:"payment"`
	Items []Item `json:"items"`
	Locale string `json:"locale"`
	Internal_signature string `json:"internal_signature"`
	Customer_id string `json:"customer_id"`
	Delivery_service string `json:"delivery_service"`
	Shardkey string `json:"shardkey"`
	Sm_id int `json:"sm_id"`
	Date_created string `json:"date_created"`
	Oof_shard string `json:"oof_shard"`

}