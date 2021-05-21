package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	"github.com/lessbutter/mealkit/config"
	"github.com/lessbutter/mealkit/src"
	"github.com/lessbutter/mealkit/src/generated"
	"github.com/lessbutter/mealkit/src/model"
	"github.com/lessbutter/mealkit/src/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *queryResolver) Products(ctx context.Context, filter model.ProductsInput) ([]*model.Product, error) {
	collection := conn.Database("mealkit").Collection("products")
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	dbfilter := bson.M{
		"$or": bson.A{
			bson.M{"removed": false},
			bson.M{"removed": bson.M{"$exists": false}},
		},
	}

	if filter.Category != nil {
		dbfilter["category.name"] = filter.Category
	} else if filter.Brand != nil {
		dbfilter["brand._id"] = *filter.Brand
	}

	if filter.Search != nil {
		dbfilter["name"] = primitive.Regex{
			Pattern: *filter.Search,
			Options: "i",
		}
	}

	options := options.Find()
	options.SetSort(bson.M{"purchasecount": -1})
	options.SetLimit(int64(filter.Limit))
	options.SetSkip(int64(filter.Offset))

	cur, err := collection.Find(ctx, dbfilter, options)
	utils.CheckErr(err)

	var products []*model.Product
	for cur.Next(ctx) {
		var product *model.Product
		err := cur.Decode(&product)
		utils.CheckErr(err)
		products = append(products, product)
	}
	return products, nil
}

func (r *queryResolver) Categories(ctx context.Context) ([]*model.Category, error) {
	collection := conn.Database("mealkit").Collection("categories")
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	options := options.Find()
	options.SetSort(bson.M{"order": 1})

	cur, _ := collection.Find(ctx, bson.M{}, options)

	var categories []*model.Category
	for cur.Next(ctx) {
		var category *model.Category
		err := cur.Decode(&category)
		utils.CheckErr(err)
		categories = append(categories, category)
	}
	return categories, nil
}

func (r *queryResolver) Brands(ctx context.Context) ([]*model.Brand, error) {
	collection := conn.Database("mealkit").Collection("brands")
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	options := options.Find()
	options.SetSort(bson.M{"order": 1})

	cur, _ := collection.Find(ctx, bson.M{}, options)

	var brands []*model.Brand
	for cur.Next(ctx) {
		var brand *model.Brand
		err := cur.Decode(&brand)
		utils.CheckErr(err)
		brands = append(brands, brand)
	}
	return brands, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
var conn = src.MongoConn(config.GetConfiguration())

func (r *queryResolver) Category(ctx context.Context, name model.Category) (*model.Category, error) {
	panic(fmt.Errorf("not implemented"))
}
