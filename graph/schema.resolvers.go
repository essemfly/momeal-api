package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"lessbutter.co/mealkit/graph/generated"
	"lessbutter.co/mealkit/graph/model"
)

func (r *queryResolver) Products(ctx context.Context, input model.QueryInput) ([]*model.Product, error) {

	// filter := bson.M{"category": input.Category}

	// // 데이터 읽기
	// res, err := GetCollection(client, collection).Find(ctx, filter)
	// U.CheckErr(err)

	// input.Limit
	// input.Offset
	// input.Category
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Category(ctx context.Context, name model.Category) (*model.Category, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
