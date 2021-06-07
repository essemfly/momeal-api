package src

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/lessbutter/momeal-api/database"
	"github.com/lessbutter/momeal-api/src/model"
	"github.com/lessbutter/momeal-api/src/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ListBrands() []*model.Brand {
	b := database.Db.Collection("brands")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	options := options.Find()
	options.SetSort(bson.M{"order": 1})

	cursor, err := b.Find(ctx, bson.M{}, options)
	utils.CheckErr(err)

	var brands []*model.Brand
	err = cursor.All(ctx, &brands)
	utils.CheckErr(err)

	return brands
}

func ListCategories() []*model.Category {
	c := database.Db.Collection("categories")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	options := options.Find()
	options.SetSort(bson.M{"order": 1})

	cursor, err := c.Find(ctx, bson.M{}, options)
	utils.CheckErr(err)

	var categories []*model.Category
	err = cursor.All(ctx, &categories)
	utils.CheckErr(err)

	return categories
}

func FindBrandByName(name string) model.Brand {
	b := database.Db.Collection("brands")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var brand model.Brand
	b.FindOne(ctx, bson.M{"name": name}).Decode(&brand)
	return brand
}

func FindCategoryByLabel(label string) model.Category {
	c := database.Db.Collection("categories")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var category model.Category
	c.FindOne(ctx, bson.M{"label": label}).Decode(&category)
	return category
}

func AddCategories(categories []model.Category) {
	categoriesCollection := database.Db.Collection("categories")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	opts := options.Update().SetUpsert(true)

	for _, cat := range categories {
		filter := bson.M{"label": cat.Label}
		_, err := categoriesCollection.UpdateOne(ctx, filter, bson.M{"$set": cat}, opts)
		utils.CheckErr(err)
	}
}

func AddBrands(brands []model.Brand) {
	brandsCollection := database.Db.Collection("brands")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	opts := options.Update().SetUpsert(true)

	for _, brand := range brands {
		filter := bson.M{"name": brand.Name}
		_, err := brandsCollection.UpdateOne(ctx, filter, bson.M{"$set": brand}, opts)
		utils.CheckErr(err)
	}
}

func AddProduct(product *model.Product) {
	pc := database.Db.Collection("products")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	opts := options.Update().SetUpsert(true)

	filter := bson.M{"name": &product.Name}
	_, err := pc.UpdateOne(ctx, filter, bson.M{"$set": &product}, opts)
	utils.CheckErr(err)
}

func UpdateProductsFieldExcept(products []*model.Product) []*model.Product {
	pc := database.Db.Collection("products")
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

func AddProducts(products []*model.Product) {
	pc := database.Db.Collection("products")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	opts := options.Update().SetUpsert(true)

	for _, product := range products {
		filter := bson.M{"name": product.Name}
		_, err := pc.UpdateOne(ctx, filter, bson.M{"$set": &product}, opts)
		utils.CheckErr(err)
	}
}

func ListProducts(brand, category, search string, limit, offset int) []*model.Product {
	collection := database.Db.Collection("products")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	dbfilter := bson.M{
		"$or": bson.A{
			bson.M{"removed": false},
			bson.M{"removed": bson.M{"$exists": false}},
		},
	}

	if category != "" {
		dbfilter["category._id"] = category
	} else if brand != "" {
		dbfilter["brand._id"] = brand
	}

	if search != "" {
		dbfilter["name"] = primitive.Regex{
			Pattern: search,
			Options: "i",
		}
	}

	options := options.Find()
	options.SetSort(bson.D{{Key: "reviewscore", Value: -1}, {Key: "soldout", Value: 1}, {Key: "_id", Value: -1}})
	options.SetLimit(int64(limit))
	options.SetSkip(int64(offset))

	cur, err := collection.Find(ctx, dbfilter, options)
	utils.CheckErr(err)

	var products []*model.Product
	for cur.Next(ctx) {
		var product *model.Product
		err := cur.Decode(&product)
		utils.CheckErr(err)
		products = append(products, product)
	}
	return products
}

func WriteCrawlingUpdateRecord() {
	c := database.Db.Collection("crawling_records")
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

	pc := database.Db.Collection("products")
	filter := bson.M{"isnew": true, "removed": false}
	newProducts, err := pc.CountDocuments(ctx, filter)
	utils.CheckErr(err)

	outProducts, err := pc.CountDocuments(
		ctx,
		bson.M{
			"removed": false,
			"updated": bson.M{
				"$lte": primitive.NewDateTimeFromTime(lastUpdatedDate),
			},
		})
	utils.CheckErr(err)

	c.InsertOne(ctx, bson.M{"date": time.Now(), "newproducts": newProducts, "outproducts": outProducts})
	log.Println("Last Update:" + lastUpdatedDate.String())
	log.Println("NEW_PRODUCTS:" + strconv.Itoa(int(newProducts)) + "  OUT_PRODUCTS: " + strconv.Itoa(int(outProducts)))
}
