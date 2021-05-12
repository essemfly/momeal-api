package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/lessbutter/mealkit/cmd/naversearch/src"
	"github.com/lessbutter/mealkit/config"
	infra "github.com/lessbutter/mealkit/src"
	"github.com/lessbutter/mealkit/src/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

type NaverSearchResponseParser struct {
	ShoppingResult struct {
		Products []src.NaverProductEntity `json:"products"`
	} `json:"shoppingResult"`
}

func CrawlNaverSearchResult(conn *mongo.Client, wg *sync.WaitGroup, start, divide int) {
	endNumber := 679
	var client http.Client
	for i := start; i < start+divide; i++ {
		if i > endNumber {
			break
		}
		url := "https://search.shopping.naver.com/search/all?sort=date&pagingIndex=" + strconv.Itoa(i) + "&pagingSize=80&viewType=list&productSet=total&query=%EB%B0%80%ED%82%A4%ED%8A%B8"
		for {
			response, ok := makeRequest(&client, url)
			if !ok {
				log.Println("Retry: " + strconv.Itoa(i) + "")
				time.Sleep(5 * time.Second)
			} else {
				products := parseResponse(conn, response)
				log.Println("Page crawling Success: " + strconv.Itoa(i))
				if len(products) != 0 {
					src.AddNaverProducts(conn, products)
				}
				if i == start {
					time.Sleep(time.Second)
					wg.Add(1)
					go CrawlNaverSearchResult(conn, wg, start+divide, divide)
				}

				break
			}
			defer response.Body.Close()
		}

		time.Sleep(5 * time.Second)
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

func parseResponse(conn *mongo.Client, resp *http.Response) []interface{} {
	searchResults := &NaverSearchResponseParser{}
	json.NewDecoder(resp.Body).Decode(searchResults)

	products := make([]interface{}, 0)
	for _, result := range searchResults.ShoppingResult.Products {
		if !isValidCategory(result.Category1Name) {
			log.Println("Different Category name: " + result.Category1Name + " -> " + result.MallProductUrl)
			continue
		}

		isAlreadyRegisteredMall := src.CheckMallExist(conn, result.MallInfo)
		if !isAlreadyRegisteredMall {
			src.AddMall(conn, result.MallInfo)
		}

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

	wg.Add(1)
	conf := config.GetConfiguration()
	conn := infra.MongoConn(conf)
	go CrawlNaverSearchResult(conn, &wg, 1, 10)

	wg.Wait()
	log.Println("Main function End")
}
