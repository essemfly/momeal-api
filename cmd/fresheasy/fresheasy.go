package main

import (
	"log"
	"strconv"
	"strings"
	"sync"

	"github.com/gocolly/colly"
	product "lessbutter.co/mealkit/domains"
)

func CrawlFreshEasy(wg *sync.WaitGroup, pageNum string) {
	url := "https://fresheasy.co.kr/goods/search_list?per=200&sorting=regist&page=" + pageNum
	counts := 0

	c := colly.NewCollector(
		colly.AllowedDomains("fresheasy.co.kr", "www.fresheasy.co.kr"),
	)

	c.OnHTML(".horizontal-list-item", func(e *colly.HTMLElement) {
		temp := product.ProductEntity{}
		temp.DistributionUrl = "https://fresheasy.co.kr" + e.ChildAttr(".item_img_area a", "href")
		temp.Title = e.ChildText(".goods_name a")
		temp.DistributionChannel = "FreshEasy"
		priceInString := e.ChildText(".horizontal-list-item-info__sale-price")
		temp.Price = strings.ReplaceAll(priceInString, "\"", "")

		product.AddProduct(temp)
		counts += 1

	})
	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL.String())
	})

	c.Visit(url)
	log.Print("Page: " + pageNum + ", Total: " + strconv.Itoa(counts) + "products added")
	wg.Done()
}

func main() {
	log.Println("Main function Start")
	var wg sync.WaitGroup

	wg.Add(1)
	go CrawlFreshEasy(&wg, "1")

	wg.Wait()
	log.Println("Main function End")
}
