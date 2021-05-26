package ownstore

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	crawler "github.com/lessbutter/mealkit/cmd/crawler/utils"
	infra "github.com/lessbutter/mealkit/src"
	"github.com/lessbutter/mealkit/src/model"
	"github.com/lessbutter/mealkit/src/utils"

	"golang.org/x/net/html"
)

type CookitResponseParser struct {
	TotalCount int    `json:"totalCount"`
	Html       string `json:"html"`
}

type Node struct {
	Type  string
	Token html.Token
	Doc   *html.Tokenizer
}

func CrawlCookit(wg *sync.WaitGroup, brand model.Brand) {
	log.Println("Cookit gogosing")

	categories := infra.ListCategories()
	resp, ok := crawler.MakeRequest(brand.CrawlingUrl)
	defer resp.Body.Close()
	if !ok {
		log.Println("Retry: " + brand.Name)
	}

	crawlResults := &CookitResponseParser{}
	json.NewDecoder(resp.Body).Decode(crawlResults)

	r := strings.NewReader(crawlResults.Html)
	doc, err := goquery.NewDocumentFromReader(r)
	utils.CheckErr(err)
	num := 0

	doc.Find(".item").Each(func(i int, s *goquery.Selection) {
		num += 1
		producturl, _ := s.Find("div.txt_wrap a.conts").Attr("href")
		originalid := strings.Split(producturl, "=")[1]
		imageurl, _ := s.Find("div.img_wrap img").Attr("src")
		name := s.Find("div.txt_wrap a.conts div.tit_info").Text()
		name = strings.TrimSpace(name)
		price := s.Find("div.txt_wrap a.conts div.price_info p.sale span.price").Text()
		priceInt := utils.ParsePriceString(price)
		discountdPrice := s.Find("div.txt_wrap a div.price_info p.cost").Text()
		discountedPriceInt := utils.ParsePriceString(discountdPrice)
		reviewcountString := s.Find("div.txt_wrap a.conts div.etc_info div.review_num").Text()
		reviewcount := utils.ParsePriceString(reviewcountString)
		reviewscore := s.Find("div.txt_wrap a.conts div.etc_info div.rating_wrap span.rating_star span.star").Text()
		reviewscore = strings.TrimSpace(reviewscore)
		reviewscore = strings.TrimLeft(reviewscore, "별점 ")
		reviewscore = strings.TrimRight(reviewscore, "점")
		reviewscoreFloat, _ := strconv.ParseFloat(reviewscore, 64)

		if discountedPriceInt > 0 {
			discountedPriceInt = discountedPriceInt - priceInt
		}

		product := model.Product{
			Name:            name,
			Producturl:      "https://www.cjcookit.com/" + producturl,
			Imageurl:        "https:" + imageurl,
			Price:           priceInt,
			Discountedprice: discountedPriceInt,
			Brand:           &brand,
			Deliveryfee:     "",
			Category:        crawler.InferProductCategoryFromName(categories, name),
			Reviewcount:     reviewcount,
			Reviewscore:     reviewscoreFloat,
			Mallname:        brand.CrawlFrom,
			Originalid:      originalid,
			Soldout:         false,
			Removed:         false,
			Created:         time.Now(),
			Updated:         time.Now(),
		}

		products := infra.UpdateProductsFieldExcept([]*model.Product{&product})
		infra.AddProducts(products)
	})

	log.Println(brand.Name + " NUM: " + strconv.Itoa(num))

	wg.Done()
}
