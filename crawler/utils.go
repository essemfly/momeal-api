package crawler

import (
	"github.com/gocolly/colly"
	"lessbutter.co/mealkit/product"
)

func MainCrawler() []product.ProductEntity {
	c := colly.NewCollector(
		colly.AllowedDomains("https://essemfly.com"),
	)

	c.OnHTML(".products li", func(element *colly.HTMLElement) {

	})

	var ret []product.ProductEntity

	sampleProduct := product.ProductEntity{
		Title:               "seokmin",
		Manufacture:         "LessButter co.",
		DistributionChannel: "Softbank",
		DistributionUrl:     "https://lessbutter.co",
		Price:               "100000000",
		Servings:            3,
	}

	ret = append(ret, sampleProduct)
	return ret
}
