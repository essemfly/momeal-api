package src

import (
	"context"
	"time"

	"github.com/lessbutter/mealkit/src/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddNaverProducts(conn *mongo.Client, products []interface{}) (*mongo.InsertManyResult, error) {
	productsCollection := conn.Database("mealkit").Collection("naver_products")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	ret, err := productsCollection.InsertMany(ctx, products)
	utils.CheckErr(err)

	return ret, err
}

func CheckMallExist(conn *mongo.Client, mall NaverMallEntity) bool {
	mallsCollection := conn.Database("mealkit").Collection("naver_malls")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result := mallsCollection.FindOne(ctx, bson.M{"name": mall.Name})
	return result.Err() == nil
}

func AddMall(conn *mongo.Client, mall NaverMallEntity) (*mongo.InsertOneResult, error) {
	mallsCollection := conn.Database("mealkit").Collection("naver_malls")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	addMallResult, err := mallsCollection.InsertOne(ctx, bson.D{
		{Key: "name", Value: mall.Name},
		{Key: "bizplBaseAddr", Value: mall.BizAddr},
		{Key: "businessNo", Value: mall.BizNo},
		{Key: "mallLogos", Value: mall.LogoUrl},
		{Key: "mallIntroduction", Value: mall.Description},
	})
	utils.CheckErr(err)
	return addMallResult, err
}
