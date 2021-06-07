package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/lessbutter/momeal-api/src"
	"github.com/lessbutter/momeal-api/src/generated"
	"github.com/lessbutter/momeal-api/src/model"
)

func (r *queryResolver) Products(ctx context.Context, filter model.ProductsInput) ([]*model.Product, error) {
	categoryQuery := ""
	brandQuery := ""
	searchQuery := ""

	if filter.Category != nil {
		categoryQuery = *filter.Category
	} else if filter.Brand != nil {
		brandQuery = *filter.Brand
	}

	if filter.Search != nil {
		searchQuery = *filter.Search
	}

	products := src.ListProducts(brandQuery, categoryQuery, searchQuery, filter.Limit, filter.Offset)
	return products, nil
}

func (r *queryResolver) Categories(ctx context.Context) ([]*model.Category, error) {
	categories := src.ListCategories()
	return categories, nil
}

func (r *queryResolver) Brands(ctx context.Context) ([]*model.Brand, error) {
	brands := src.ListBrands()
	return brands, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
