package models

type Company struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	PriceTrack []int `json:"priceTrack"`
	CurrentPrice int `json:"currentPrice"`
}