package models

type Book struct {
	Title string
}

type BookIndex struct {
	PageTitle string
	Books     []Book
}
