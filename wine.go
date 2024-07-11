package main

type Wine struct {
	Name   string  `json:"name"`
	Origin string  `json:"origin"`
	Type   string  `json:"type"`
	Review string  `json:"review"`
	Points int     `json:"points"`
	Price  float32 `json:"price"`
}

// NewWine creates a new Wine instance
func NewWine(name, origin, typeWine, review string, points int, price float32) *Wine {
	return &Wine{
		Name:   name,
		Origin: origin,
		Type:   typeWine,
		Review: review,
		Points: points,
		Price:  price,
	}
}
