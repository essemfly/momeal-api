package ownstore

import (
	"log"
	"strconv"
	"strings"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gocolly/colly"
	crawler "github.com/lessbutter/mealkit/cmd/crawler/utils"
	infra "github.com/lessbutter/mealkit/src"
	"github.com/lessbutter/mealkit/src/model"
	"github.com/lessbutter/mealkit/src/utils"
)

func CrawlPeacock(conn *mongo.Client, wg *sync.WaitGroup, brand model.Brand) {
	categories := infra.ListCategories(conn)
	url := "http://emart.ssg.com/specialStore/ssgpeacock/ajaxSubItemList.ssg?aplSiteNo=6001&pageSize=100&ctgId=6000073847"
	c := colly.NewCollector(
		colly.AllowedDomains("emart.ssg.com"),
	)

	c.OnHTML(".cunit_t232", func(e *colly.HTMLElement) {
		var product model.Product

		product.Name = e.ChildText(".cunit_info .cunit_md .title a > em:first-child")
		product.Imageurl = "http:" + e.ChildAttr(".cunit_prod .thmb a img:first-child", "src")
		product.Producturl = e.ChildAttr(".cunit_prod .thmb a", "href")
		product.Price = utils.ParsePriceString(e.ChildText(".ssg_price"))
		product.Discountedprice = 0
		product.Brand = &brand
		product.Deliveryfee = ""
		product.Category = crawler.InferProductCategoryFromName(conn, categories, product.Name)
		product.Purchasecount = 0
		product.Reviewcount = utils.ParsePriceString(e.ChildText(".rate_tx em"))
		reviewscore := strings.Replace(e.ChildText(".rate_bg .blind"), "별점 ", "", 1)
		reviewscore = strings.Replace(reviewscore, "점", "", 1)
		reviewscoreFloat, _ := strconv.ParseFloat(reviewscore, 32)
		product.Reviewscore = reviewscoreFloat
		product.Mallname = "emart"
		product.Originalid = e.ChildAttr(".cunit_prod .thmb a", "data-info")
		product.Soldout = false
		soldout := e.ChildText(".cunit_soldout")
		if soldout != "" {
			product.Soldout = true
		}
		product.Removed = false
		product.Created = time.Now()
		product.Updated = time.Now()

		if product.Name != "" {
			infra.AddProduct(conn, product)
		}
	})

	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL.String())
	})

	c.Visit(url)

	log.Println(brand.Name + ": Finished")
	wg.Done()
}
