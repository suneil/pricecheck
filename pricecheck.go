package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

// Product asdf
type Product struct {
	url  string
	name string
}

func price(product Product, wg *sync.WaitGroup) {
	defer wg.Done()

	doc, err := goquery.NewDocument(product.url)
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
	var wg sync.WaitGroup

	products := [2]Product{
		Product{"http://www.amazon.com/dp/B00PXYRMPE", "Dell 34 in curved monitor"},
		Product{"http://www.amazon.com/dp/B00OKSEWL6", "LG 34 in curved monitor"}}

	for _, product := range products {
		wg.Add(1)
		go price(product, &wg)
	}

	wg.Wait()
}
