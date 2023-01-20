package models

type Item struct{
	Chrt_id int `json:"chrt_id"`
	Track_number string `json:"track_number"`
	Price int `json:"price"`
	Name string `json:"name"`
	Sale int `json:"sale"`
	Size string `json:"size"`
	Total_price int `json:"total_price"`
	Nm_id int `json:"nm_id"`
	Brand string `json:"brand"`
	Status int `json:"status"` 
}