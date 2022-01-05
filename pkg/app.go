package main

import (
	"log"
	"net/http"
	"text/template"
)

const port = ":8000"

type PageData struct {
	Title    string
	Products []Product
}

type Product struct {
	Name  string
	Price int
}

func main() {

	var products []Product

	products = append(products, Product{
		Name:  "Apple",
		Price: 20,
	})
	products = append(products, Product{
		Name:  "Banana",
		Price: 24,
	})
	products = append(products, Product{
		Name:  "Orange",
		Price: 23,
	})

	pageData := PageData{
		Title:    "Golang Templates Demo",
		Products: products,
	}

	http.HandleFunc("/test", func(response http.ResponseWriter, request *http.Request) {
		parsed, err := template.ParseFiles("./templates/test.tmpl")

		if err != nil {
			log.Fatal(err)
		}
		parsed.Execute(response, pageData)
	})

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
