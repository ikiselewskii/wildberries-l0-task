package models


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