package main

import (
	"log"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/suneil/pricecheck/store"
	"github.com/suneil/useragent"
)

var (
	client = &http.Client{}
)

// Product product structure
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

	req.Header.Set("User-Agent", useragent.GetRandomUserAgent())

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

	var price32 float64
	price := ""
	title := product.name

	doc.Find("#priceblock_ourprice").Each(func(i int, s *goquery.Selection) {
		price = s.Text()
	})

	price = strings.Replace(price, "$", "", -1)
	price32, err = strconv.ParseFloat(price, 64)
	if err != nil {
		price32 = 0.0
	}

	doc.Find("#productTitle").Each(func(i int, s *goquery.Selection) {
		title = s.Text()
	})

	u, _ := url.Parse(product.url)
	item := store.NewItem(path.Base(u.Path), title, price32)
	store.Store(item)
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
