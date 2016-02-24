package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// run go generate whenever useragents.txt is changed
//go:generate go run scripts/includetxt.go

var (
	client = &http.Client{}
)

// Product asdf
type Product struct {
	url      string
	name     string
	respChan chan *http.Response
}

func fetch(product *Product) {
	req, err := http.NewRequest("GET", product.url, nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("User-Agent", GetRandomUA())

	response, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	product.respChan <- response
}

func price(product *Product) {
	response := <-product.respChan

	doc, err := goquery.NewDocumentFromResponse(response)
	if err != nil {
		log.Fatal(err)
	}

	price := ""
	title := product.name

	doc.Find("#priceblock_ourprice").Each(func(i int, s *goquery.Selection) {
		price = s.Text()
	})

	doc.Find("#productTitle").Each(func(i int, s *goquery.Selection) {
		title = s.Text()
	})

	fmt.Printf("%s => %s\n", price, title)

}

func main() {
	products := [2]Product{
		Product{"http://www.amazon.com/dp/B00PXYRMPE", "Dell 34 in curved monitor", make(chan *http.Response, 1)},
		Product{"http://www.amazon.com/dp/B00OKSEWL6", "LG 34 in curved monitor", make(chan *http.Response, 1)}}

	for idx := range products {
		go fetch(&products[idx])
	}

	for _, product := range products {
		price(&product)
	}
}
