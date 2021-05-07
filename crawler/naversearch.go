package crawler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type NaverProduct struct {
	Name           string `json:"productName"`
	Title          string `json:"productTitle"`
	TitleOrg       string `json:"productTitleOrg"`
	AttributeValue string `json:"attributeValue"`
	CharacterValue string `json:"characterValue"`
	ImageUrl       string `json:"imageUrl"`
	Price          string `json:"price"`
	PriceUnit      string `json:"priceUnit"`
	Maker          string `json:"maker"`
	Brand          string `json:"brand"`
	MallName       string `json:"mallName"`
	MallNameOrg    string `json:"mallNameOrg"`
	MallProductUrl string `json:"mallProductUrl"`
	DeliveryFee    string `json:"dlvryCont"`
}

type Mall struct {
	Name             string `json:"name"`
	Address          string `json:"bizplBaseAddr"`
	BusinessNo       string `json:"businessNo"`
	MallIntroduction string `json:"mallIntroduction"`
}

func CrawlNaverSearch(wg *sync.WaitGroup, pageNum string) []NaverProduct {
	// conn, _ := storage.MongoConn()
	// productsCollection := conn.Database("mealkit").Collection("products")
	// mallsCollection := conn.Database("mealkit").Collection("malls")

	// Category ID
	// url := fmt.Sprintf("https://search.shopping.naver.com/search/category?sort=rel&pagingIndex=%v&pagingSize=80&viewType=list&productSet=total&catId=50006808&deliveryFee=&deliveryTypeValue=&iq=&eq=&xq=", pageNum)
	// Search Result
	url := "https://search.shopping.naver.com/search/all?sort=date&pagingIndex=" + pageNum + "&pagingSize=80&viewType=list&productSet=total&query=%EB%B0%80%ED%82%A4%ED%8A%B8"
	counts := 0

	d := GetData(url)

	var data map[string]interface{}
	json.Unmarshal([]byte(d), &data)
	shoppingResult := data["shoppingResult"]
	productResults := shoppingResult["products"]
	log.Print(data["shoppingResult"])

	var pdSlice []NaverProduct
	var pd NaverProduct
	err := json.Unmarshal(d, &pd)

	if err != nil {
		panic(err)
	}

	log.Print("Page: " + pageNum + ", Total: " + strconv.Itoa(counts) + "products added")
	wg.Done()

	return pdSlice
}

func GetData(url string) []byte {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	var client = http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	jsonByte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	return jsonByte
}
