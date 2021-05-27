package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	"github.com/lessbutter/momeal-api/database"
	"github.com/lessbutter/momeal-api/src/generated"
	"github.com/lessbutter/momeal-api/src/model"
	"github.com/lessbutter/momeal-api/src/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *queryResolver) Products(ctx context.Context, filter model.ProductsInput) ([]*model.Product, error) {
	collection := database.Db.Collection("products")
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
	options.SetSort(bson.D{{Key: "purchasecount", Value: -1}, {Key: "soldout", Value: 1}})
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
	collection := database.Db.Collection("categories")
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
	collection := database.Db.Collection("brands")
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
