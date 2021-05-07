package product

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductEntity struct {
	Title               string
	ImagePath           string
	Manufacture         string
	DistributionChannel string
	DistributionUrl     string
	Price               string
	Servings            int
	Category            string
}

func AddProduct(writeDb *mongo.Collection, product ProductEntity) (*mongo.InsertOneResult, error) {
	addProductResult, err := writeDb.InsertOne(context.TODO(), bson.D{
		{Key: "title", Value: product.Title},
		{Key: "manufacture", Value: product.Manufacture},
		{Key: "distribution_channel", Value: product.DistributionChannel},
		{Key: "distribution_url", Value: product.DistributionUrl},
		{Key: "price", Value: product.Price},
		{Key: "servings", Value: product.Servings},
	})
	if err != nil {
		log.Fatal(err)
	}
	return addProductResult, err
}
