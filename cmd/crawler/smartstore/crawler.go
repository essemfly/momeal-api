package smartstore

import (
	"encoding/json"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/lessbutter/mealkit/cmd/crawler/smartstore/src"
	crawler "github.com/lessbutter/mealkit/cmd/crawler/utils"
	infra "github.com/lessbutter/mealkit/src"
	"github.com/lessbutter/mealkit/src/model"
)

func CrawlSmartStore(wg *sync.WaitGroup, brand model.Brand) {
	pageSize := 80
	pageNum := 1
	pageurl := AddUrlPageQuery(brand.CrawlingUrl, pageNum, pageSize)

	resp, ok := crawler.MakeRequest(pageurl)
	defer resp.Body.Close()
	if !ok {
		log.Println("Retry: " + brand.Name + strconv.Itoa(pageNum))
	}

	crawlResults := &src.SmartstoreResponseParser{}
	json.NewDecoder(resp.Body).Decode(crawlResults)

	categories := infra.ListCategories()
	products := MapCrawlResultsToModels(brand, crawlResults.Products, categories)

	for crawlResults.TotalCount > pageSize*pageNum {
		pageNum += 1
		pageurl := AddUrlPageQuery(brand.CrawlingUrl, pageNum, pageSize)
		resp, ok := crawler.MakeRequest(pageurl)
		defer resp.Body.Close()
		if !ok {
			log.Println("Retry: " + brand.Name + strconv.Itoa(pageNum))
			pageNum -= 1
		} else {
			results := &src.SmartstoreResponseParser{}
			json.NewDecoder(resp.Body).Decode(results)
			products = append(products, MapCrawlResultsToModels(brand, results.Products, categories)...)
		}
	}
	products = infra.UpdateProductsFieldExcept(products)
	infra.AddProducts(products)

	log.Println(brand.Name + " NUM: " + strconv.Itoa(len(products)))
	wg.Done()
}

func AddUrlPageQuery(url string, page int, pageSize int) string {
	return url + "&page=" + strconv.Itoa(page) + "&pageSize=" + strconv.Itoa(pageSize)
}

func BuildProductUrl(brandname string, productid int) string {
	return "https://smartstore.naver.com/" + brandname + "/products/" + strconv.Itoa(productid)
}

func CheckOutofStock(status string) bool {
	return status == "OUTOFSTOCK"
}

func MapCrawlResultsToModels(brand model.Brand, products []src.SmartstoreProductEntity, categories []model.Category) []*model.Product {
	var newProducts []*model.Product
	for _, product := range products {
		newProduct := model.Product{
			Name:            product.Name,
			Imageurl:        product.Imageurl,
			Price:           product.Benefits.MobileDiscountedSalePrice,
			Discountedprice: product.Benefits.MobileSellerImmediateDiscountAmount,
			Brand:           &brand,
			Producturl:      BuildProductUrl(brand.SmartstoreBrandName, product.NaverProductId),
			Deliveryfee:     strconv.Itoa(product.DeliveryInfo.BaseFee),
			Category:        crawler.InferProductCategoryFromName(categories, product.Name),
			Purchasecount:   product.SaleAmount.CumulationSaleCount,
			Reviewcount:     product.ReviewAmount.TotalReviewCount,
			Reviewscore:     float64(product.ReviewAmount.AverageReviewScore),
			Mallname:        brand.CrawlFrom,
			Originalid:      strconv.Itoa(product.NaverProductId),
			Soldout:         CheckOutofStock(product.ProductStatus),
			Removed:         false,
			Created:         time.Now(),
			Updated:         time.Now(),
		}
		newProducts = append(newProducts, &newProduct)
	}
	return newProducts
}
