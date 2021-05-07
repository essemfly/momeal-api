package main

import (
	"fmt"

	"lessbutter.co/mealkit/product"
	"lessbutter.co/mealkit/storage"
)

func init() {
	fmt.Println("Initialization..")
}

func main() {
	fmt.Println("Main function Start")
	conn, _ := storage.MongoConn()
	mongo := conn.Database("mealkit").Collection("products")

	sample_product := product.ProductEntity{
		Title:               "seokmin",
		Manufacture:         "LessButter co.",
		DistributionChannel: "Softbank",
		DistributionUrl:     "https://lessbutter.co",
		Price:               100000000,
		Servings:            3,
	}
	ret, _ := product.AddProduct(mongo, sample_product)
	fmt.Printf("------------ID: %v is uploaded------------", ret.InsertedID)
	fmt.Println("Main function End")
}
