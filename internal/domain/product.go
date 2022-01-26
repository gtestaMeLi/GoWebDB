package domain

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Type  string  `json:"type"`
	Count int     `json:"cunt"`
	Price float64 `json:"price"`
}
