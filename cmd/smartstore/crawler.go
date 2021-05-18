package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/lessbutter/mealkit/cmd/smartstore/src"
	"github.com/lessbutter/mealkit/config"
	infra "github.com/lessbutter/mealkit/src"
	"github.com/lessbutter/mealkit/src/model"
	"github.com/lessbutter/mealkit/src/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	conf := config.GetConfiguration()
	conn := infra.MongoConn(conf)

	brands := infra.ListBrands(conn)
	CrawlBrands(conn, brands)
}

func CrawlBrands(conn *mongo.Client, brands []model.Brand) {
	var wg sync.WaitGroup
	for _, brand := range brands {
		if brand.CrawlFrom != "smartstore" {
			continue
		}
		wg.Add(1)
		go CrawlSmartStore(conn, &wg, brand)
	}
	wg.Wait()
}

func CrawlSmartStore(conn *mongo.Client, wg *sync.WaitGroup, brand model.Brand) {
	pageSize := 80
	pageNum := 1
	pageurl := AddUrlPageQuery(brand.CrawlingUrl, pageNum, pageSize)

	resp, ok := makeRequest(pageurl)
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
		resp, ok := makeRequest(pageurl)
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
			Category:        InferProductCategoryFromName(conn, categories, product.Name),
			Purchasecount:   product.SaleAmount.CumulationSaleCount,
			Reviewcount:     product.ReviewAmount.TotalReviewCount,
			Reviewscore:     float64(product.ReviewAmount.AverageReviewScore),
			Mallname:        brand.CrawlFrom,
		}
		newProducts = append(newProducts, newProduct)
	}
	return newProducts
}

func InferProductCategoryFromName(conn *mongo.Client, categories []model.Category, name string) *model.Category {
	category_mapper := map[string]string{
		"지리탕":    "해물탕",
		"연포탕":    "해물탕",
		"매운탕":    "해물탕",
		"동태탕":    "해물탕",
		"맑은탕":    "해물탕",
		"뽈탕":     "해물탕",
		"바지락탕":   "해물탕",
		"꽃게탕":    "해물탕",
		"대구탕":    "해물탕",
		"쭈꾸미탕":   "해물탕",
		"새우탕":    "해물탕",
		"알고니탕":   "해물탕",
		"짬뽕탕":    "해물탕",
		"조개탕":    "해물탕",
		"해물탕":    "해물탕",
		"육개장":    "육개장",
		"마라탕":    "마라탕",
		"된장찌개":   "된장찌개",
		"된장 찌개":  "된장찌개",
		"김치찌개":   "김치찌개",
		"김치 찌개":  "김치찌개",
		"감바스":    "감바스",
		"어묵탕":    "어묵탕",
		"덮밥":     "덮밥/비빔밥",
		"비빔밥":    "덮밥/비빔밥",
		"규동":     "덮밥/비빔밥",
		"새우장밥":   "덮밥/비빔밥",
		"짜장밥":    "덮밥/비빔밥",
		"커리":     "덮밥/비빔밥",
		"카레":     "덮밥/비빔밥",
		"곱창전골":   "곱창전골",
		"청국장":    "청국장",
		"부대찌개":   "부대찌개",
		"알탕":     "알탕",
		"밀푀유나베":  "밀푀유나베",
		"우거지탕":   "우거지탕",
		"스파게티":   "파스타",
		"파스타":    "파스타",
		"순두부찌개":  "순두부찌개",
		"순두부 찌개": "순두부찌개",
		"해물순두부":  "순두부찌개",
		"차돌박이":   "고기",
		"너바아니":   "고기",
		"제육":     "고기",
		"삼겹":     "고기",
		"오겹":     "고기",
		"곱창":     "고기",
		"막창":     "고기",
		"대창":     "고기",
		"보쌈":     "고기",
		"산적":     "고기",
		"닭갈비":    "고기",
		"바베큐":    "고기",
		"탄두리":    "고기",
		"회과육":    "고기",
		"빼당":     "고기",
		"갈비":     "고기",
		"목살":     "고기",
		"불고기":    "고기",
		"폭립":     "고기",
		"블랙라벨":   "스테이크",
		"채끝":     "스테이크",
		"스테이크":   "스테이크",
		"분식":     "분식",
		"라면":     "분식",
		"쫄면":     "분식",
		"김밥":     "분식",
		"떡볶이":    "분식",
		"소바":     "면",
		"우동":     "면",
		"라멘":     "면",
		"팟타이":    "면",
		"국수":     "면",
		"탄탄면":    "면",
		"야끼누들":   "면",
		"밀면":     "면",
		"초마짬뽕":   "면",
		"짜장면":    "면",
		"빠네":     "면",
		"봉골레":    "면",
		"리가토니":   "면",
		"나베":     "기타 국/탕/전골",
		"전골":     "기타 국/탕/전골",
		"샤브":     "기타 국/탕/전골",
		"훠궈":     "기타 국/탕/전골",
		"닭도리탕":   "기타 국/탕/전골",
		"삼계탕":    "기타 국/탕/전골",
		"순도리탕":   "기타 국/탕/전골",
		"황태진국":   "기타 국/탕/전골",
		"순대국":    "기타 국/탕/전골",
		"미역국":    "기타 국/탕/전골",
		"된장국":    "기타 국/탕/전골",
		"젓국":     "기타 국/탕/전골",
		"뭇국":     "기타 국/탕/전골",
		"들깨탕":    "기타 국/탕/전골",
		"북엇국":    "기타 국/탕/전골",
		"시래기국":   "기타 국/탕/전골",
		"불백":     "기타 국/탕/전골",
		"스키야키":   "기타 국/탕/전골",
		"찌개":     "기타 국/탕/전골",
		"두루치기":   "볶음, 찜요리",
		"볶음":     "볶음, 찜요리",
		"칠리새우":   "볶음, 찜요리",
		"고추잡채":   "볶음, 찜요리",
		"깐쇼새우":   "볶음, 찜요리",
		"어향가지":   "볶음, 찜요리",
		"버섯잡채":   "볶음, 찜요리",
		"마파두부":   "볶음, 찜요리",
		"쏭타이":    "볶음, 찜요리",
		"난자완스":   "볶음, 찜요리",
		"마라샹궈":   "볶음, 찜요리",
		"스키야끼":   "볶음, 찜요리",
		"찜":      "볶음, 찜요리",
		"튀김":     "기타 요리",
		"퀘사디아":   "기타 요리",
		"짜글":     "기타 요리",
		"수제비":    "기타 요리",
		"월남쌈":    "기타 요리",
		"부추전":    "기타 요리",
		"맥앤치즈":   "기타 요리",
		"비프스튜":   "기타 요리",
		"샌드위치":   "기타 요리",
		"굴전":     "기타 요리",
		"오꼬노미야끼": "기타 요리",
		"솥밥":     "기타 요리",
		"파히타":    "기타 요리",
		"멕시칸":    "기타 요리",
		"유린기":    "기타 요리",
		"플래터":    "기타 요리",
		"양꿍":     "기타 요리",
		"포케":     "기타 요리",
		"치즈구이":   "기타 요리",
		"유산슬":    "기타 요리",
	}

	for k, v := range category_mapper {
		if strings.Contains(name, k) {
			cat := infra.FindCategoryByLabel(conn, v)
			return &cat
		}
	}
	return nil
}

func makeRequest(url string) (*http.Response, bool) {
	var client http.Client
	req, err := http.NewRequest("GET", url, nil)
	utils.CheckErr(err)

	req.Header.Add("user-agent", "Crawler")
	req.Header.Add("accept", "application/json")

	resp, err := client.Do(req)
	utils.CheckErr(err)

	if resp.StatusCode == http.StatusOK {
		return resp, true
	} else {
		b, _ := ioutil.ReadAll(resp.Body)
		log.Println(string(b))
		return resp, false
	}
}
