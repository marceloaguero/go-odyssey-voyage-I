package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"context"
	"fmt"

	"github.com/marceloaguero/go-odyssey-voyage-I/reviews/graph/model"
)

// Reviews is the resolver for the reviews field.
func (r *locationResolver) Reviews(ctx context.Context, obj *model.Location) ([]*model.Review, error) {
	panic(fmt.Errorf("not implemented: Reviews - reviews"))
}

// SubmitReview is the resolver for the submitReview field.
func (r *mutationResolver) SubmitReview(ctx context.Context, locationReview *model.LocationReviewInput) (*model.SubmitReviewResponse, error) {
	review, err := r.usecase.Create(locationReview)

	reviewResponse := &model.SubmitReviewResponse{
		Code:           200,
		Success:        true,
		Message:        "Review added!",
		LocationReview: review,
	}

	return reviewResponse, err
}

// LatestReviews is the resolver for the latestReviews field.
func (r *queryResolver) LatestReviews(ctx context.Context) ([]*model.Review, error) {
	return r.usecase.GetAll()
}

// Location returns LocationResolver implementation.
func (r *Resolver) Location() LocationResolver { return &locationResolver{r} }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type locationResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
