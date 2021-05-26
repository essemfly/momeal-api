package ownstore

import (
	"log"
	"strconv"
	"strings"
	"sync"
	"time"

	crawler "github.com/lessbutter/mealkit/cmd/crawler/utils"
	infra "github.com/lessbutter/mealkit/src"

	"github.com/gocolly/colly"
	"github.com/lessbutter/mealkit/src/model"
	"github.com/lessbutter/mealkit/src/utils"
)

func CrawlMonokitchen(wg *sync.WaitGroup, brand model.Brand) {
	categories := infra.ListCategories()

	url := "http://mono-kitchen.co.kr/shop/shopbrand.html?xcode=015"
	c := colly.NewCollector(
		colly.AllowedDomains("mono-kitchen.co.kr"),
	)
	num := 0

	c.OnHTML(".list_kitchen li", func(e *colly.HTMLElement) {
		var product model.Product
		product.Name = e.ChildText("a .info_box .tit")
		if len(product.Name) > 0 {
			num += 1
			product.Price = utils.ParsePriceString(e.ChildText("a .info_box .pay .cost span"))
			product.Discountedprice = 0
			originalPrice := utils.ParsePriceString(e.ChildText("a .info_box .pay .discount span"))
			if originalPrice > 0 {
				product.Discountedprice = originalPrice - product.Price
			}
			product.Producturl = "http://mono-kitchen.co.kr" + e.ChildAttr("a", "href")
			product.Imageurl = "http://mono-kitchen.co.kr" + e.ChildAttr("a .img_box .thumbnail .centered img", "src")
			product.Brand = &brand
			product.Deliveryfee = ""
			product.Category = crawler.InferProductCategoryFromName(categories, product.Name)
			product.Purchasecount = 0
			product.Reviewcount = 0
			product.Reviewscore = 0
			product.Mallname = brand.CrawlFrom
			product.Originalid = strings.Split(strings.Split(e.ChildAttr("a", "href"), "&")[0], "=")[1]
			product.Soldout = false
			product.Removed = false
			product.Created = time.Now()
			product.Updated = time.Now()

			svgDiv := e.ChildAttr("a .img_box .thumbnail .centered svg filter", "id")
			if len(svgDiv) > 0 {
				product.Soldout = true
				product.Imageurl = "http://mono-kitchen.co.kr" + e.ChildAttr("a .img_box .thumbnail .centered svg image", "href")
			}
			products := infra.UpdateProductsFieldExcept([]*model.Product{&product})
			infra.AddProducts(products)
		}
	})

	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL.String())
		r.ResponseCharacterEncoding = "euc-kr"
	})

	for _, page := range []int{1, 2, 3} {
		pageurl := url + "&page=" + strconv.Itoa(page)
		c.Visit(pageurl)
	}

	log.Println(brand.Name + " NUM: " + strconv.Itoa(num))

	wg.Done()
}
