package ownstore

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"

	crawler "github.com/lessbutter/momeal-api/cmd/crawler/utils"
	infra "github.com/lessbutter/momeal-api/src"
	"github.com/lessbutter/momeal-api/src/model"
)

type SimplyCookResponseParser struct {
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
	SuppleFirmCd     string `json:"supplFirmCd"`
	StorId           string `json:"storId"`
	MallId           string `json:"mallId"`
}

type SimplyCookReviewEntity struct {
	Success bool `json:"success"`
	Data    struct {
		EvlAvgScor string `json:"evlAvgScor"`
		TotEvlCnt  string `json:"totEvlCnt"`
	} `json:"data"`
}

func CrawlSimplycook(wg *sync.WaitGroup, brand model.Brand) {
	resp, ok := crawler.MakeRequest(brand.CrawlingUrl)
	defer resp.Body.Close()
	if !ok {
		log.Println("Retry: " + brand.Name)
	}

	crawlResults := &SimplyCookResponseParser{}
	json.NewDecoder(resp.Body).Decode(crawlResults)

	categories := infra.ListCategories()
	var newProducts []*model.Product
	for _, field := range crawlResults.Data.Fields {
		newProducts = append(newProducts, MapCrawlResultsToModels(brand, field.Products, categories)...)
	}

	newProducts = infra.UpdateProductsFieldExcept(newProducts)
	infra.AddProducts(newProducts)

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

func MapCrawlResultsToModels(brand model.Brand, products []SimplyCookProductEntity, categories []model.Category) []*model.Product {
	var newProducts []*model.Product
	for _, product := range products {
		possibleQty, _ := strconv.Atoi(product.SellPosbQty)

		reviewUrl := "https://api.gsecretail.com/fo/md/itmcmgnt/item-basis-comment-evl-score?itemId=" + product.ItemId
		resp, ok := crawler.MakeRequest(reviewUrl)
		defer resp.Body.Close()
		if !ok {
			log.Println("Failed in getting review: " + brand.Name + " - " + product.ItemName)
		}
		reviewCrawlResults := &SimplyCookReviewEntity{}
		json.NewDecoder(resp.Body).Decode(reviewCrawlResults)
		reviewCount, _ := strconv.Atoi(reviewCrawlResults.Data.TotEvlCnt)
		reviewScore, _ := strconv.ParseFloat(reviewCrawlResults.Data.EvlAvgScor, 32)

		newProduct := model.Product{
			Name:            product.ItemName,
			Imageurl:        BuildImageurl(product.ItemImg),
			Price:           product.DiscountedPrice,
			Discountedprice: product.NormalPrice - product.DiscountedPrice,
			Producturl:      "https://m.gsfresh.com/md/product_detail?itemId=" + product.ItemId + "&storId=" + product.StorId + "&supplFirmCd=" + product.SuppleFirmCd + "&mallId=" + product.MallId,
			Deliveryfee:     "",
			Brand:           &brand,
			Category:        crawler.InferProductCategoryFromName(categories, product.ItemName),
			Purchasecount:   0,
			Reviewcount:     reviewCount,
			Reviewscore:     reviewScore,
			Mallname:        brand.CrawlFrom,
			Originalid:      product.ItemId,
			Soldout:         true,
			Removed:         false,
			Created:         time.Now(),
			Updated:         time.Now(),
		}

		if possibleQty > 0 {
			newProduct.Soldout = false
		}
		newProducts = append(newProducts, &newProduct)
	}
	return newProducts
}
