package tasty9

import (
	"log"
	"strconv"
	"strings"
	"sync"

	crawler "github.com/lessbutter/mealkit/cmd/crawler/utils"
	infra "github.com/lessbutter/mealkit/src"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gocolly/colly"
	"github.com/lessbutter/mealkit/src/model"
	"github.com/lessbutter/mealkit/src/utils"
)

func CrawlTasty9(conn *mongo.Client, wg *sync.WaitGroup, brand model.Brand) {
	categories := infra.ListCategories(conn)

	url := "https://tasty9.com/product/list.html?cate_no=95"
	c := colly.NewCollector(
		colly.AllowedDomains("tasty9.com", "www.tasty9.com"),
	)

	c.OnHTML(".box", func(e *colly.HTMLElement) {
		var product model.Product
		saledPrice := e.ChildText(".description ul li:first-child > span")
		saledPrice = strings.Split(saledPrice, "원")[0]
		saledPriceInt := utils.ParsePriceString(saledPrice)
		discountedPrice := e.ChildText(".description ul li:nth-child(2) > span")
		discountedPrice = strings.Split(discountedPrice, "원")[0]
		discountedPriceInt := utils.ParsePriceString(discountedPrice)

		product.Name = strings.Replace(e.ChildText(".description .name a"), "상품명 : ", "", 1)
		product.Imageurl = "https:" + e.ChildAttr(".thumbnail .prdImg a img", "src")
		product.Producturl = "https://tasty9.com" + e.ChildAttr("a", "href")
		product.Price = saledPriceInt
		product.Discountedprice = 0
		if discountedPriceInt > 0 {
			product.Discountedprice = discountedPriceInt - saledPriceInt
		}
		product.Brand = &brand
		product.Deliveryfee = ""
		product.Category = crawler.InferProductCategoryFromName(conn, categories, product.Name)
		product.Purchasecount = 0
		product.Reviewcount = 0
		product.Reviewscore = 0
		product.Mallname = "Tasty9"

		if product.Name != "" {
			infra.AddProduct(conn, product)
		}
	})

	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL.String())
	})

	for _, page := range []int{1, 2} {
		pageurl := url + "&page=" + strconv.Itoa(page)
		c.Visit(pageurl)
	}

	wg.Done()
}
