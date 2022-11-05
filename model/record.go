package model

type Record struct {
	Id             string  `json:"id"`
	Price          float32 `json:"price"`
	ExpirationDate string  `json:"expirationDate"`
}
