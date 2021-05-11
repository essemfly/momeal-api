package domains

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"lessbutter.co/mealkit/external"
	"lessbutter.co/mealkit/utils"
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

// TODO: brand field, category field generation 생성 필요
// 현재는 MallInfo기준으로 brands들을 만들고 있음
type NaverProductEntity struct {
	Name           string          `json:"productName"`
	Title          string          `json:"productTitle"`
	TitleOrg       string          `json:"productTitleOrg"`
	AttributeValue string          `json:"attributeValue"`
	CharacterValue string          `json:"characterValue"`
	ImageUrl       string          `json:"imageUrl"`
	Price          string          `json:"price"`
	PriceUnit      string          `json:"priceUnit"`
	Maker          string          `json:"maker"`
	Brand          string          `json:"brand"`
	Category1Name  string          `json:"category1Name"`
	MallName       string          `json:"mallName"`
	MallNameOrg    string          `json:"mallNameOrg"`
	MallProductUrl string          `json:"mallProductUrl"`
	DeliveryFee    string          `json:"dlvryCont"`
	PurchaseCount  int             `json:"purchaseCnt"`
	ReviewCount    int             `json:"reviewCountSum"`
	MallInfo       NaverMallEntity `json:"mallInfoCache"`
}

type NaverMallEntity struct {
	Name        string            `json:"name"`
	BizAddr     string            `'json:"bizplBaseAddr"`
	BizNo       string            `json:"businessNo"`
	Description string            `json:"mallIntroduction"`
	LogoUrl     map[string]string `json:"mallLogos"`
}

func AddProduct(product ProductEntity) (*mongo.InsertOneResult, error) {
	conn := external.MongoConn()
	productsCollection := conn.Database("mealkit").Collection("products")
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

func AddNaverProducts(products []interface{}) (*mongo.InsertManyResult, error) {
	conn := external.MongoConn()
	productsCollection := conn.Database("mealkit").Collection("products")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	ret, err := productsCollection.InsertMany(ctx, products)
	utils.CheckErr(err)

	return ret, err
}

func AddMall(mall NaverMallEntity) (*mongo.InsertOneResult, error) {
	conn := external.MongoConn()
	mallsCollection := conn.Database("mealkit").Collection("malls")
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

func CheckMallExist(mall NaverMallEntity) bool {
	conn := external.MongoConn()
	mallsCollection := conn.Database("mealkit").Collection("malls")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result := mallsCollection.FindOne(ctx, bson.M{"name": mall.Name})
	return result.Err() == nil
}
