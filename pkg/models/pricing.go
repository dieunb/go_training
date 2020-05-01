package models

type Pricing struct {
	Title    string
	Amount   int
	Features []Feature
}

type Feature struct {
	Name string
}
