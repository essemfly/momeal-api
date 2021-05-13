package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"net/url"

	"github.com/lessbutter/mealkit/cmd/naversearch/src"
	"github.com/lessbutter/mealkit/config"
	infra "github.com/lessbutter/mealkit/src"
	"github.com/lessbutter/mealkit/src/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

type NaverSearchResponseParser struct {
	ShoppingResult struct {
		Products []src.NaverProductEntity `json:"products"`
		Total    int                      `json:"total"`
	} `json:"shoppingResult"`
}

func getCategoryProductsNum(conn *mongo.Client, wg *sync.WaitGroup, query string) int {
	var client http.Client
	queryUtf := url.QueryEscape(query)
	url := "https://search.shopping.naver.com/search/all?sort=date&pagingIndex=1&pagingSize=80&viewType=list&productSet=total&query=" + queryUtf + "&catId=50000006"

	for {
		response, ok := makeRequest(&client, url)
		defer response.Body.Close()
		if !ok {
			log.Println("Retry: " + query)
			time.Sleep(10 * time.Second)
		} else {
			searchResults := &NaverSearchResponseParser{}
			json.NewDecoder(response.Body).Decode(searchResults)
			return searchResults.ShoppingResult.Total
		}
	}
}

func CrawlNaverSearchStarter(conn *mongo.Client, wg *sync.WaitGroup, query string, numTotalProducts int, pageSize int, divider int) {
	endPage := numTotalProducts / pageSize
	rem := numTotalProducts % pageSize
	if rem > 0 {
		endPage += 1
	}

	log.Println("Keyword: " + query + " total pages: " + strconv.Itoa(endPage))

	i := 1
	for {
		wg.Add(1)
		if i+divider > endPage+1 {
			go CrawlNaverSearchResult(conn, wg, query, pageSize, i, endPage+1)
			break
		} else {
			go CrawlNaverSearchResult(conn, wg, query, pageSize, i, i+divider)
			i += divider
		}
		time.Sleep(10 * time.Second)
	}
	wg.Done()
}

func CrawlNaverSearchResult(conn *mongo.Client, wg *sync.WaitGroup, query string, pageSize int, start int, end int) {
	var client http.Client

	queryUtf := url.QueryEscape(query)
	for i := start; i < end; i++ {
		log.Println("    Running keyword: " + query + " page#: " + strconv.Itoa(i))
		url := "https://search.shopping.naver.com/search/all?sort=date&pagingIndex=" + strconv.Itoa(i) + "&pagingSize=" + strconv.Itoa(pageSize) + "&viewType=list&productSet=total&query=" + queryUtf + "&catId=50000006"
		for {
			response, ok := makeRequest(&client, url)
			if !ok {
				log.Println("Retry: " + strconv.Itoa(i))
				time.Sleep(10 * time.Second)
			} else {
				products := parseResponse(conn, response, query)
				log.Println("    Page:" + query + " crawling Success: " + strconv.Itoa(i))
				if len(products) != 0 {
					src.AddNaverProducts(conn, products)
				}
				break
			}
			defer response.Body.Close()
		}

		time.Sleep(10 * time.Second)
	}
	wg.Done()
}

func makeRequest(client *http.Client, url string) (*http.Response, bool) {
	req, err := http.NewRequest("GET", url, nil)
	utils.CheckErr(err)

	req.Header.Add("user-agent", "Crawler")
	req.Header.Add("urlprefix", "/api")
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

func parseResponse(conn *mongo.Client, resp *http.Response, query string) []interface{} {
	searchResults := &NaverSearchResponseParser{}
	json.NewDecoder(resp.Body).Decode(searchResults)

	products := make([]interface{}, 0)
	for _, result := range searchResults.ShoppingResult.Products {
		if !isValidCategory(result.Category1Name) {
			// log.Println("Different Category name: " + result.Category1Name + " -> " + result.MallProductUrl)
			continue
		}

		// isAlreadyRegisteredMall := src.CheckMallExist(conn, result.MallInfo)
		// if !isAlreadyRegisteredMall {
		// 	src.AddMall(conn, result.MallInfo)
		// }
		result.MomilBrand = query
		products = append(products, result)
	}

	return products
}

func isValidCategory(category string) bool {
	switch category {
	case
		"식품",
		"출산/육아":
		return true
	}
	return false
}

func main() {
	log.Println("Main function Start")
	var wg sync.WaitGroup

	conf := config.GetConfiguration()
	conn := infra.MongoConn(conf)
	categories := []string{
		"맛수러움",
		"프레시지",
		"프레시밀",
		"푸드어셈블",
		"자연맛남 밀키트",
		"얌테이블 밀키트",
		"이츠웰 밀키트",
		"쿡솜씨",
		"피코크 밀키트",
		"마이셰프",
		"에슐리 밀키트",
		"파우즈",
		"앙트레",
		"올쿡",
		"모노키친 밀키트",
		"닥터키친 밀키트",
		"파파쿡",
		"심플리쿡",
		"아로이키친",
		"테이스티나인",
	}
	for _, val := range categories {
		numTotalProducts := getCategoryProductsNum(conn, &wg, val)
		pageSize := 80
		divider := 10
		wg.Add(1)
		go CrawlNaverSearchStarter(conn, &wg, val, numTotalProducts, pageSize, divider)
		time.Sleep(time.Second * 10)
	}

	wg.Wait()
	log.Println("Main function End")
}
