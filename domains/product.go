package domains

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"lessbutter.co/mealkit/storage"
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

type NaverProductEntity struct {
	Name           string `json:"productName"`
	Title          string `json:"productTitle"`
	TitleOrg       string `json:"productTitleOrg"`
	AttributeValue string `json:"attributeValue"`
	CharacterValue string `json:"characterValue"`
	ImageUrl       string `json:"imageUrl"`
	Price          string `json:"price"`
	PriceUnit      string `json:"priceUnit"`
	Maker          string `json:"maker"`
	Brand          string `json:"brand"`
	MallName       string `json:"mallName"`
	MallNameOrg    string `json:"mallNameOrg"`
	MallProductUrl string `json:"mallProductUrl"`
	DeliveryFee    string `json:"dlvryCont"`
}

func AddProduct(product ProductEntity) (*mongo.InsertOneResult, error) {
	conn, _ := storage.MongoConn()
	productsCollection := conn.Database("mealkit").Collection("products")
	addProductResult, err := productsCollection.InsertOne(context.TODO(), bson.D{
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

func AddNaverProducts(products []interface{}) (*mongo.InsertManyResult, error) {
	conn, _ := storage.MongoConn()
	productsCollection := conn.Database("mealkit").Collection("products")
	ret, err := productsCollection.InsertMany(context.TODO(), products)
	if err != nil {
		log.Fatal(err)
	}

	return ret, err
}
