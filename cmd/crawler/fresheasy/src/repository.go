package src

import (
	"context"
	"time"

	"github.com/lessbutter/mealkit/src/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddProduct(conn *mongo.Client, product FresheasyProductEntity) (*mongo.InsertOneResult, error) {
	productsCollection := conn.Database("mealkit").Collection("fresheasy_products")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	addProductResult, err := productsCollection.InsertOne(ctx, bson.D{
		{Key: "title", Value: product.Title},
		{Key: "manufacture", Value: product.Manufacture},
		{Key: "distribution_channel", Value: product.DistributionChannel},
		{Key: "distribution_url", Value: product.DistributionUrl},
		{Key: "price", Value: product.Price},
		{Key: "servings", Value: product.Servings},
	})
	utils.CheckErr(err)
	return addProductResult, err
}
