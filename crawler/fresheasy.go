package crawler

import (
	"log"
	"strconv"
	"strings"
	"sync"

	"github.com/gocolly/colly"
	"lessbutter.co/mealkit/product"
	"lessbutter.co/mealkit/storage"
)

func CrawlFreshEasy(wg *sync.WaitGroup, pageNum string) {
	conn, _ := storage.MongoConn()
	mongo := conn.Database("mealkit").Collection("products")
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

		product.AddProduct(mongo, temp)
		counts += 1

	})
	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL.String())
	})

	c.Visit(url)
	log.Print("Page: " + pageNum + ", Total: " + strconv.Itoa(counts) + "products added")
	wg.Done()
}
