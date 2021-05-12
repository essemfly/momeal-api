package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	"github.com/lessbutter/mealkit/config"
	"github.com/lessbutter/mealkit/src"
	generated1 "github.com/lessbutter/mealkit/src/generated"
	"github.com/lessbutter/mealkit/src/model"
	"github.com/lessbutter/mealkit/src/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *queryResolver) Products(ctx context.Context, filter model.ProductsInput) ([]*model.Product, error) {
	collection := conn.Database("mealkit").Collection("products")
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	dbfilter := bson.M{}
	if filter.Category != nil {
		dbfilter = bson.M{"category": filter.Category}
	} else if filter.Brand != nil {
		dbfilter = bson.M{"brand": filter.Brand}
	}

	options := options.Find()
	options.SetSort(bson.D{{"_id", -1}})
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

	cur, _ := collection.Find(ctx, bson.M{})

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
	collection := conn.Database("mealkit").Collection("malls")
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	cur, _ := collection.Find(ctx, bson.M{})

	var brands []*model.Brand
	for cur.Next(ctx) {
		var brand *model.Brand
		err := cur.Decode(&brand)
		utils.CheckErr(err)
		brands = append(brands, brand)
	}
	return brands, nil
}

// Query returns generated1.QueryResolver implementation.
func (r *Resolver) Query() generated1.QueryResolver { return &queryResolver{r} }

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
