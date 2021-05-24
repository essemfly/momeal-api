package src

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"

	"github.com/lessbutter/mealkit/src/model"
	"github.com/lessbutter/mealkit/src/utils"
)

func ListBrands(conn *mongo.Client) []model.Brand {
	b := conn.Database("mealkit").Collection("brands")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := b.Find(ctx, bson.M{})
	utils.CheckErr(err)

	var brands []model.Brand
	err = cursor.All(ctx, &brands)
	utils.CheckErr(err)

	return brands
}

func ListCategories(conn *mongo.Client) []model.Category {
	c := conn.Database("mealkit").Collection("categories")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := c.Find(ctx, bson.M{})
	utils.CheckErr(err)

	var categories []model.Category
	err = cursor.All(ctx, &categories)
	utils.CheckErr(err)

	return categories
}

func FindBrandByName(conn *mongo.Client, name string) model.Brand {
	b := conn.Database("mealkit").Collection("brands")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var brand model.Brand
	b.FindOne(ctx, bson.M{"name": name}).Decode(&brand)
	return brand
}

func FindCategoryByLabel(conn *mongo.Client, label string) model.Category {
	c := conn.Database("mealkit").Collection("categories")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var category model.Category
	c.FindOne(ctx, bson.M{"label": label}).Decode(&category)
	return category
}

func AddCategories(conn *mongo.Client, categories []model.Category) {
	categoriesCollection := conn.Database("mealkit").Collection("categories")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	opts := options.Update().SetUpsert(true)

	for _, cat := range categories {
		filter := bson.M{"name": cat.Name}
		_, err := categoriesCollection.UpdateOne(ctx, filter, bson.M{"$set": cat}, opts)
		utils.CheckErr(err)
	}
}

func AddBrands(conn *mongo.Client, brands []model.Brand) {
	brandsCollection := conn.Database("mealkit").Collection("brands")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	opts := options.Update().SetUpsert(true)

	for _, brand := range brands {
		filter := bson.M{"name": brand.Name}
		_, err := brandsCollection.UpdateOne(ctx, filter, bson.M{"$set": brand}, opts)
		utils.CheckErr(err)
	}
}

func AddProduct(conn *mongo.Client, product model.Product) {
	pc := conn.Database("mealkit").Collection("products")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	opts := options.Update().SetUpsert(true)

	filter := bson.M{"name": product.Name}

	var oldProduct model.Product
	err := pc.FindOne(ctx, filter).Decode(&oldProduct)
	if err == nil {
		product.Category = oldProduct.Category
		product.Created = oldProduct.Created
		product.Removed = oldProduct.Removed
	}
	_, err = pc.UpdateOne(ctx, filter, bson.M{"$set": product}, opts)
	utils.CheckErr(err)
}

func AddProducts(conn *mongo.Client, products []model.Product) {
	pc := conn.Database("mealkit").Collection("products")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	opts := options.Update().SetUpsert(true)

	for _, product := range products {
		var oldProduct model.Product
		filter := bson.M{"name": product.Name}
		// 기존에 상품이 있었으면, 그 상품의 category는 바꾸지 않는다.
		err := pc.FindOne(ctx, filter).Decode(&oldProduct)
		if err == nil {
			product.Category = oldProduct.Category
			product.Created = oldProduct.Created
			product.Removed = oldProduct.Removed
		}
		_, err = pc.UpdateOne(ctx, filter, bson.M{"$set": product}, opts)
		utils.CheckErr(err)
	}
}
