package crawler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"sync"

	product "lessbutter.co/mealkit/domains"
)

type Mall struct {
	Name             string `json:"name"`
	Address          string `json:"bizplBaseAddr"`
	BusinessNo       string `json:"businessNo"`
	MallIntroduction string `json:"mallIntroduction"`
}

type ResponseParser struct {
	ShoppingResult struct {
		Products []product.NaverProductEntity `json:"products"`
	} `json:"shoppingResult"`
}

func CrawlNaverSearch(wg *sync.WaitGroup, pageNum int) {

	url := "https://search.shopping.naver.com/search/all?sort=date&pagingIndex=" + strconv.Itoa(pageNum) + "&pagingSize=80&viewType=list&productSet=total&query=%EB%B0%80%ED%82%A4%ED%8A%B8"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Page crawling: " + strconv.Itoa(pageNum))

	req.Header.Add("user-agent", "Crawler")
	req.Header.Add("urlprefix", "/api")
	req.Header.Add("accept", "application/json")

	var client http.Client
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		if err != nil {
			log.Fatal(err)
		}

		response := &ResponseParser{}
		json.NewDecoder(resp.Body).Decode(response)

		products := make([]interface{}, 0)
		for _, result := range response.ShoppingResult.Products {
			products = append(products, result)
		}
		product.AddNaverProducts(products)

		if len(products) == 80 {
			wg.Add(1)
			go CrawlNaverSearch(wg, pageNum+1)
		}
		wg.Done()
	} else {
		log.Fatal("http Status Not OK")
		wg.Done()
	}
}
