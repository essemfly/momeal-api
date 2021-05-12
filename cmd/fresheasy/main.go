package main

import (
	"log"
	"strconv"
	"strings"
	"sync"

	"github.com/gocolly/colly"
	"go.mongodb.org/mongo-driver/mongo"
	"lessbutter.co/mealkit/cmd/fresheasy/src"
	"lessbutter.co/mealkit/config"
	infra "lessbutter.co/mealkit/src"
)

func CrawlFreshEasy(conn *mongo.Client, wg *sync.WaitGroup, pageNum string) {
	url := "https://fresheasy.co.kr/goods/search_list?per=200&sorting=regist&page=" + pageNum
	counts := 0

	c := colly.NewCollector(
		colly.AllowedDomains("fresheasy.co.kr", "www.fresheasy.co.kr"),
	)

	c.OnHTML(".horizontal-list-item", func(e *colly.HTMLElement) {
		temp := src.FresheasyProductEntity{}
		temp.DistributionUrl = "https://fresheasy.co.kr" + e.ChildAttr(".item_img_area a", "href")
		temp.Title = e.ChildText(".goods_name a")
		temp.DistributionChannel = "FreshEasy"
		priceInString := e.ChildText(".horizontal-list-item-info__sale-price")
		temp.Price = strings.ReplaceAll(priceInString, "\"", "")

		src.AddProduct(conn, temp)
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
	config := config.GetConfiguration()
	conn := infra.MongoConn(config)
	go CrawlFreshEasy(conn, &wg, "1")

	wg.Wait()
	log.Println("Main function End")
}
