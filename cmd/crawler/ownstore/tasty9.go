package ownstore

import (
	"log"
	"strconv"
	"strings"
	"sync"
	"time"

	crawler "github.com/lessbutter/momeal-api/cmd/crawler/utils"
	infra "github.com/lessbutter/momeal-api/src"

	"github.com/gocolly/colly"
	"github.com/lessbutter/momeal-api/src/model"
	"github.com/lessbutter/momeal-api/src/utils"
)

func CrawlTasty9(wg *sync.WaitGroup, brand model.Brand) {
	categories := infra.ListCategories()

	url := "https://tasty9.com/product/list.html?cate_no=95"
	c := colly.NewCollector(
		colly.AllowedDomains("tasty9.com", "www.tasty9.com"),
	)
	num := 0

	c.OnHTML(".xans-record-", func(e *colly.HTMLElement) {
		productName := strings.Replace(e.ChildText(".description .name a"), "상품명 : ", "", 1)
		if productName != "" {
			num += 1

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
			product.Category = crawler.InferProductCategoryFromName(categories, product.Name)
			product.Purchasecount = 0
			product.Reviewcount = 0
			product.Reviewscore = 0
			product.Mallname = "tasty9"
			product.Originalid = strings.Split(e.Attr("id"), "_")[1]
			product.Soldout = false
			product.Removed = false
			product.Created = time.Now()
			product.Updated = time.Now()

			products := infra.UpdateProductsFieldExcept([]*model.Product{&product})
			infra.AddProducts(products)
		}
	})

	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL.String())
	})

	for _, page := range []int{1, 2} {
		pageurl := url + "&page=" + strconv.Itoa(page)
		c.Visit(pageurl)
	}

	log.Println(brand.Name + " NUM: " + strconv.Itoa(num))

	wg.Done()
}
