package main

import (
	"sync"

	"github.com/lessbutter/momeal-api/cmd/crawler/ownstore"
	"github.com/lessbutter/momeal-api/cmd/crawler/smartstore"
	"github.com/lessbutter/momeal-api/config"
	"github.com/lessbutter/momeal-api/database"
	infra "github.com/lessbutter/momeal-api/src"
	"github.com/lessbutter/momeal-api/src/model"
)

func main() {
	conf := config.GetConfiguration()
	database.InitDB(conf)

	brands := infra.ListBrands()
	CrawlBrands(brands)
}

func CrawlBrands(brands []model.Brand) {
	var wg sync.WaitGroup
	for _, brand := range brands {
		switch brand.CrawlFrom {
		case "smartstore":
			wg.Add(1)
			go smartstore.CrawlSmartStore(&wg, brand)
		case "tasty9":
			wg.Add(1)
			go ownstore.CrawlTasty9(&wg, brand)
		case "emart":
			wg.Add(1)
			go ownstore.CrawlPeacock(&wg, brand)
		case "gsshop":
			wg.Add(1)
			go ownstore.CrawlSimplycook(&wg, brand)
		case "monokitchen":
			wg.Add(1)
			go ownstore.CrawlMonokitchen(&wg, brand)
		case "cookit":
			wg.Add(1)
			go ownstore.CrawlCookit(&wg, brand)
		default:

		}
	}
	wg.Wait()
	infra.WriteCrawlingUpdateRecord()
}
