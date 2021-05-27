package main

import (
	"github.com/lessbutter/momeal-api/cmd/seeds/seeds"
	"github.com/lessbutter/momeal-api/config"
	"github.com/lessbutter/momeal-api/database"
)

func main() {
	conf := config.GetConfiguration()
	database.InitDB(conf)
	seeds.AddCategories()
	seeds.AddBrands()
}
