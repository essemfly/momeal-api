package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"lessbutter.co/mealkit/external"
	"lessbutter.co/mealkit/graph/generated"
	"lessbutter.co/mealkit/graph/model"
	"lessbutter.co/mealkit/utils"
)

func (r *queryResolver) Products(ctx context.Context, input model.ProductsInput) ([]*model.Product, error) {
	collection := conn.Database("mealkit").Collection("products")
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	filter := bson.M{"brand": "프레시지"}
	cur, err := collection.Find(ctx, filter)
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

func (r *queryResolver) Categories(ctx context.Context) ([]model.Category, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Brands(ctx context.Context) ([]*model.Brand, error) {
	panic(fmt.Errorf("not implemented"))
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
var conn = external.MongoConn()

func (r *queryResolver) Category(ctx context.Context, name model.Category) (*model.Category, error) {
	panic(fmt.Errorf("not implemented"))
}
