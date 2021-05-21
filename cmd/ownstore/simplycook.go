package ownstore

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"
	"sync"

	crawler "github.com/lessbutter/mealkit/cmd/crawler/utils"
	infra "github.com/lessbutter/mealkit/src"
	"github.com/lessbutter/mealkit/src/model"

	"go.mongodb.org/mongo-driver/mongo"
)

type SimplycookResponseParser struct {
	Success bool `json:"success"`
	Data    struct {
		Fields []struct {
			Products []SimplyCookProductEntity `json:"products"`
		} `json:"fields"`
	} `json:"data"`
}

type SimplyCookProductEntity struct {
	ItemName         string `json:"itemNm"`
	ItemImg          string `json:"itemImg"`
	ItemId           string `json:"itemId"`
	SellPosbQty      string `json:"sellPosbQty"`
	NormalPrice      int    `json:"normSprc"`
	BeforeSellPrice  int    `json:"bfrSellAmt"`
	DiscountedPrice  int    `json:"sellAmt"`
	DiscountedPrice2 int    `json:"mbsDcAmt"`
}

func CrawlSimplycook(conn *mongo.Client, wg *sync.WaitGroup, brand model.Brand) {
	resp, ok := crawler.MakeRequest(brand.CrawlingUrl)
	defer resp.Body.Close()
	if !ok {
		log.Println("Retry: " + brand.Name)
	}

	crawlResults := &SimplycookResponseParser{}
	json.NewDecoder(resp.Body).Decode(crawlResults)

	categories := infra.ListCategories(conn)
	var newProducts []model.Product
	for _, field := range crawlResults.Data.Fields {
		newProducts = append(newProducts, MapCrawlResultsToModels(conn, brand, field.Products, categories)...)
	}

	infra.AddProducts(conn, newProducts)

	log.Println(brand.Name + " NUM: " + strconv.Itoa(len(newProducts)))
	wg.Done()
}

func BuildProductUrl(brandname string, productid int) string {
	return "https://smartstore.naver.com/" + brandname + "/products/" + strconv.Itoa(productid)
}

func BuildImageurl(imgUrl string) string {
	preUrl := "https://image.gsecretail.com" + imgUrl
	preUrl = strings.Replace(preUrl, "{SIZE}", "300", 1)
	return preUrl
}

func MapCrawlResultsToModels(conn *mongo.Client, brand model.Brand, products []SimplyCookProductEntity, categories []model.Category) []model.Product {
	var newProducts []model.Product
	for _, product := range products {
		possibleQty, _ := strconv.Atoi(product.SellPosbQty)
		if possibleQty > 0 {
			newProduct := model.Product{
				Name:            product.ItemName,
				Imageurl:        BuildImageurl(product.ItemImg),
				Price:           product.DiscountedPrice,
				Discountedprice: product.NormalPrice - product.DiscountedPrice,
				Producturl:      "https://www.gsfresh.com/md/product_detail?itemId=" + product.ItemId,
				Deliveryfee:     "",
				Brand:           &brand,
				Category:        crawler.InferProductCategoryFromName(conn, categories, product.ItemName),
				Purchasecount:   0,
				Reviewcount:     0,
				Reviewscore:     0.0,
				Mallname:        brand.CrawlFrom,
			}
			newProducts = append(newProducts, newProduct)
		} else {
			log.Println(product.ItemName)
		}
	}
	return newProducts
}
