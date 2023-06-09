package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"context"

	"github.com/marceloaguero/go-odyssey-voyage-I/locations/graph/model"
)

// NewLocation is the resolver for the newLocation field.
func (r *mutationResolver) NewLocation(ctx context.Context, input model.LocationInput) (*model.Location, error) {
	return r.usecase.Create(&input)
}

// Locations is the resolver for the locations field.
func (r *queryResolver) Locations(ctx context.Context) ([]*model.Location, error) {
	return r.usecase.GetAll()
}

// Location is the resolver for the location field.
func (r *queryResolver) Location(ctx context.Context, id string) (*model.Location, error) {
	return r.usecase.GetByID(id)
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
