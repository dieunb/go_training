package models

type Pricing struct {
	Title    string    `json:"title"`
	Amount   int       `json:"amount"`
	Features []Feature `json:"features"`
}

type Feature struct {
	Name string `json:"name"`
}
