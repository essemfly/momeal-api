package smartstore

import (
	"encoding/json"
	"log"
	"strconv"
	"sync"

	crawler "github.com/lessbutter/mealkit/cmd/crawler/utils"
	"github.com/lessbutter/mealkit/cmd/smartstore/src"
	infra "github.com/lessbutter/mealkit/src"
	"github.com/lessbutter/mealkit/src/model"

	"go.mongodb.org/mongo-driver/mongo"
)

func CrawlSmartStore(conn *mongo.Client, wg *sync.WaitGroup, brand model.Brand) {
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

	categories := infra.ListCategories(conn)
	products := MapCrawlResultsToModels(conn, brand, crawlResults.Products, categories)

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
			products = append(products, MapCrawlResultsToModels(conn, brand, results.Products, categories)...)
		}
	}

	infra.AddProducts(conn, products)

	log.Println(brand.Name + " NUM: " + strconv.Itoa(len(products)))
	wg.Done()
}

func AddUrlPageQuery(url string, page int, pageSize int) string {
	return url + "&page=" + strconv.Itoa(page) + "&pageSize=" + strconv.Itoa(pageSize)
}

func BuildProductUrl(brandname string, productid int) string {
	return "https://smartstore.naver.com/" + brandname + "/products/" + strconv.Itoa(productid)
}

func MapCrawlResultsToModels(conn *mongo.Client, brand model.Brand, products []src.SmartstoreProductEntity, categories []model.Category) []model.Product {
	var newProducts []model.Product
	for _, product := range products {
		newProduct := model.Product{
			Name:            product.Name,
			Imageurl:        product.Imageurl,
			Price:           product.Benefits.MobileDiscountedSalePrice,
			Discountedprice: product.Benefits.MobileSellerImmediateDiscountAmount,
			Brand:           &brand,
			Producturl:      BuildProductUrl(brand.SmartstoreBrandName, product.NaverProductId),
			Deliveryfee:     strconv.Itoa(product.DeliveryInfo.BaseFee),
			Category:        crawler.InferProductCategoryFromName(conn, categories, product.Name),
			Purchasecount:   product.SaleAmount.CumulationSaleCount,
			Reviewcount:     product.ReviewAmount.TotalReviewCount,
			Reviewscore:     float64(product.ReviewAmount.AverageReviewScore),
			Mallname:        brand.CrawlFrom,
		}
		newProducts = append(newProducts, newProduct)
	}
	return newProducts
}
