package src

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"

	"lessbutter.co/mealkit/src/utils"
)

func AddCategories(conn *mongo.Client, categories []interface{}) (*mongo.InsertManyResult, error) {
	categoriesCollection := conn.Database("mealkit").Collection("categories")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	ret, err := categoriesCollection.InsertMany(ctx, categories)
	utils.CheckErr(err)

	return ret, err
}

func AddBrands(conn *mongo.Client, brands []interface{}) (*mongo.InsertManyResult, error) {
	brandsCollection := conn.Database("mealkit").Collection("brands")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	ret, err := brandsCollection.InsertMany(ctx, brands)
	utils.CheckErr(err)

	return ret, err
}
