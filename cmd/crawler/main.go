package main

import (
	"sync"

	"github.com/lessbutter/mealkit/cmd/smartstore"
	"github.com/lessbutter/mealkit/cmd/tasty9"
	"github.com/lessbutter/mealkit/config"
	infra "github.com/lessbutter/mealkit/src"
	"github.com/lessbutter/mealkit/src/model"
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
		switch brand.CrawlFrom {
		case "smartstore":
			wg.Add(1)
			go smartstore.CrawlSmartStore(conn, &wg, brand)
		case "tasty9":
			wg.Add(1)
			go tasty9.CrawlTasty9(conn, &wg, brand)
		default:

		}
	}
	wg.Wait()
}
