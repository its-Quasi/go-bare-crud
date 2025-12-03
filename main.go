package main

import (
	"fmt"
)

type Book struct {
	Title string
	Pages uint
}

func do(book *Book) {
	book.Pages++
}

func main() {
	ddia := Book{
		Title: "DDIA",
		Pages: 101,
	}

	fmt.Println(ddia)

	do(&ddia)

	fmt.Println(ddia)
}
