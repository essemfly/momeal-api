package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	product "lessbutter.co/mealkit/domains"
	"lessbutter.co/mealkit/utils"
)

type Mall struct {
	Name             string `json:"name"`
	Address          string `json:"bizplBaseAddr"`
	BusinessNo       string `json:"businessNo"`
	MallIntroduction string `json:"mallIntroduction"`
}

type NaverSearchResponseParser struct {
	ShoppingResult struct {
		Products []product.NaverProductEntity `json:"products"`
	} `json:"shoppingResult"`
}

func CrawlNaverSearchResult(wg *sync.WaitGroup, start, divide int) {
	endNumber := 6767
	var client http.Client
	for i := start; i < start+divide; i++ {
		if i > endNumber {
			break
		}
		url := "https://search.shopping.naver.com/search/all?sort=date&pagingIndex=" + strconv.Itoa(i) + "&pagingSize=80&viewType=list&productSet=total&query=%EB%B0%80%ED%82%A4%ED%8A%B8"
		for {
			products, ok := makeRequest(&client, url)
			if !ok {
				log.Println("Retry: " + strconv.Itoa(i) + "")
				time.Sleep(5 * time.Second)
			} else {
				log.Println("Page crawling Success: " + strconv.Itoa(i))
				if len(products) != 0 {
					product.AddNaverProducts(products)
				}
				if i == start && len(products) == 80 {
					wg.Add(1)
					go CrawlNaverSearchResult(wg, start+divide, divide)
				}

				break
			}
		}

		time.Sleep(5 * time.Second)
	}
	wg.Done()
}

func makeRequest(client *http.Client, url string) ([]interface{}, bool) {
	req, err := http.NewRequest("GET", url, nil)
	utils.CheckErr(err)

	req.Header.Add("user-agent", "Crawler")
	req.Header.Add("urlprefix", "/api")
	req.Header.Add("accept", "application/json")

	resp, err := client.Do(req)
	utils.CheckErr(err)
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		products := parseResponse(resp)
		return products, true
	} else {
		b, _ := ioutil.ReadAll(resp.Body)
		log.Println(string(b))
		return nil, false
	}
}

func parseResponse(resp *http.Response) []interface{} {
	searchResults := &NaverSearchResponseParser{}
	json.NewDecoder(resp.Body).Decode(searchResults)

	products := make([]interface{}, 0)
	for _, result := range searchResults.ShoppingResult.Products {
		products = append(products, result)
	}
	return products
}

func main() {
	log.Println("Main function Start")
	var wg sync.WaitGroup

	wg.Add(1)
	go CrawlNaverSearchResult(&wg, 1, 10)

	wg.Wait()
	log.Println("Main function End")
}
