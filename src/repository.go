package src

import (
	"context"
	"log"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
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

func AddProduct(conn *mongo.Client, product *model.Product) {
	pc := conn.Database("mealkit").Collection("products")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	opts := options.Update().SetUpsert(true)

	filter := bson.M{"name": &product.Name}
	_, err := pc.UpdateOne(ctx, filter, bson.M{"$set": &product}, opts)
	utils.CheckErr(err)
}

func UpdateProductsFieldExcept(conn *mongo.Client, products []*model.Product) []*model.Product {
	pc := conn.Database("mealkit").Collection("products")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	for _, product := range products {
		var oldProduct model.Product
		filter := bson.M{"name": &product.Name}
		err := pc.FindOne(ctx, filter).Decode(&oldProduct)
		if err == nil {
			product.Category = oldProduct.Category
			product.Created = oldProduct.Created
			product.Removed = oldProduct.Removed
			product.IsNew = false
		} else {
			product.IsNew = true
		}
	}
	return products
}

func AddProducts(conn *mongo.Client, products []*model.Product) {
	pc := conn.Database("mealkit").Collection("products")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	opts := options.Update().SetUpsert(true)

	for _, product := range products {
		filter := bson.M{"name": product.Name}
		_, err := pc.UpdateOne(ctx, filter, bson.M{"$set": &product}, opts)
		utils.CheckErr(err)
	}
}

func WriteCrawlingUpdateRecord(conn *mongo.Client) {
	c := conn.Database("mealkit").Collection("crawling_records")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	options := options.FindOne()
	options.SetSort(bson.M{"_id": -1})

	var lastRecord model.CrawlingRecord
	err := c.FindOne(ctx, bson.M{}, options).Decode(&lastRecord)

	lastUpdatedDate := lastRecord.Date
	if err != nil {
		lastUpdatedDate = time.Now().UTC()
	}

	pc := conn.Database("mealkit").Collection("products")
	filter := bson.M{"isnew": true, "removed": false}
	newProducts, err := pc.CountDocuments(ctx, filter)
	utils.CheckErr(err)

	outProducts, err := pc.CountDocuments(
		ctx,
		bson.M{
			"updated": bson.M{
				"$lte": primitive.NewDateTimeFromTime(lastUpdatedDate),
			},
		})
	utils.CheckErr(err)

	c.InsertOne(ctx, bson.M{"date": time.Now(), "newproducts": newProducts, "outproducts": outProducts})
	log.Println("Last Update:" + lastUpdatedDate.String())
	log.Println("NEW_PRODUCTS:" + strconv.Itoa(int(newProducts)) + "  OUT_PRODUCTS: " + strconv.Itoa(int(outProducts)))
}
