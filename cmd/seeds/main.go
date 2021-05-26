package main

import (
	"github.com/lessbutter/mealkit/cmd/seeds/seeds"
	"github.com/lessbutter/mealkit/config"
	"github.com/lessbutter/mealkit/database"
)

func main() {
	conf := config.GetConfiguration()
	database.InitDB(conf)
	seeds.AddCategories()
	seeds.AddBrands()
}
